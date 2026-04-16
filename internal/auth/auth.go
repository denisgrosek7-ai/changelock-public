package auth

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	ModeDisabled    = "disabled"
	ModeStaticToken = "static-token"
	ModeOIDCJWT     = "oidc-jwt"

	RoleViewer        = "viewer"
	RoleOperator      = "operator"
	RoleSecurityAdmin = "security_admin"
	RoleService       = "service_internal"

	IdentityTypeHuman   = "human"
	IdentityTypeService = "service"
)

type TokenConfig struct {
	Token   string `json:"token"`
	Subject string `json:"subject"`
	Role    string `json:"role"`
	TokenID string `json:"token_id,omitempty"`
}

type Config struct {
	Mode   string
	tokens []TokenConfig
	oidc   *oidcVerifier
}

type Principal struct {
	Authenticated bool   `json:"authenticated"`
	AuthMode      string `json:"auth_mode"`
	Subject       string `json:"subject,omitempty"`
	Role          string `json:"role,omitempty"`
	TokenID       string `json:"token_id,omitempty"`
	IdentityType  string `json:"identity_type,omitempty"`
	Email         string `json:"email,omitempty"`
	TenantID      string `json:"tenant_id,omitempty"`
	GlobalScope   bool   `json:"global_scope,omitempty"`
}

type AccessError struct {
	Status  int
	Message string
}

func (e *AccessError) Error() string {
	return e.Message
}

var (
	ErrMissingBearerToken      = &AccessError{Status: http.StatusUnauthorized, Message: "bearer token required"}
	ErrMalformedAuthorization  = &AccessError{Status: http.StatusUnauthorized, Message: "malformed authorization header"}
	ErrInvalidBearerToken      = &AccessError{Status: http.StatusUnauthorized, Message: "invalid bearer token"}
	ErrNoMappedRole            = &AccessError{Status: http.StatusForbidden, Message: "no ChangeLock role mapping for token"}
	ErrTenantScopeRequired     = &AccessError{Status: http.StatusForbidden, Message: "tenant claim is required for this token"}
	ErrInsufficientPermissions = &AccessError{Status: http.StatusForbidden, Message: "insufficient role"}
)

type principalContextKey struct{}

func DisabledConfig() Config {
	return Config{Mode: ModeDisabled}
}

func ParseConfig(mode string, rawTokens string) (Config, error) {
	normalizedMode := strings.ToLower(strings.TrimSpace(mode))
	if normalizedMode == "" {
		normalizedMode = ModeDisabled
	}

	switch normalizedMode {
	case ModeDisabled:
		return DisabledConfig(), nil
	case ModeStaticToken:
	default:
		return Config{}, errors.New("unsupported CHANGELOCK_AUTH_MODE: " + normalizedMode)
	}

	rawTokens = strings.TrimSpace(rawTokens)
	if rawTokens == "" {
		return Config{}, errors.New("CHANGELOCK_AUTH_TOKENS_JSON is required when CHANGELOCK_AUTH_MODE=static-token")
	}

	var entries []TokenConfig
	if err := json.Unmarshal([]byte(rawTokens), &entries); err != nil {
		return Config{}, errors.New("invalid CHANGELOCK_AUTH_TOKENS_JSON: " + err.Error())
	}
	if len(entries) == 0 {
		return Config{}, errors.New("CHANGELOCK_AUTH_TOKENS_JSON must contain at least one token")
	}

	seenTokens := map[string]struct{}{}
	seenTokenIDs := map[string]struct{}{}
	validated := make([]TokenConfig, 0, len(entries))
	for _, entry := range entries {
		entry.Token = strings.TrimSpace(entry.Token)
		entry.Subject = strings.TrimSpace(entry.Subject)
		entry.Role = strings.TrimSpace(entry.Role)
		entry.TokenID = strings.TrimSpace(entry.TokenID)

		switch {
		case entry.Token == "":
			return Config{}, errors.New("auth token entries require token")
		case entry.Subject == "":
			return Config{}, errors.New("auth token entries require subject")
		case !validRole(entry.Role):
			return Config{}, errors.New("auth token entry has unsupported role: " + entry.Role)
		}

		if _, exists := seenTokens[entry.Token]; exists {
			return Config{}, errors.New("duplicate auth token configured")
		}
		seenTokens[entry.Token] = struct{}{}

		if entry.TokenID != "" {
			if _, exists := seenTokenIDs[entry.TokenID]; exists {
				return Config{}, errors.New("duplicate auth token_id configured: " + entry.TokenID)
			}
			seenTokenIDs[entry.TokenID] = struct{}{}
		}

		validated = append(validated, entry)
	}

	return Config{
		Mode:   normalizedMode,
		tokens: validated,
	}, nil
}

func (c Config) IsEnabled() bool {
	return c.Mode != ModeDisabled
}

func (c Config) Self() Principal {
	if c.Mode == ModeDisabled {
		return Principal{
			Authenticated: true,
			AuthMode:      ModeDisabled,
			Subject:       "auth-disabled",
			Role:          RoleSecurityAdmin,
			TokenID:       "disabled-mode",
			IdentityType:  IdentityTypeHuman,
			GlobalScope:   true,
		}
	}
	return Principal{
		Authenticated: false,
		AuthMode:      c.Mode,
	}
}

func (c Config) AuthenticateRequest(r *http.Request) (Principal, error) {
	if c.Mode == ModeDisabled {
		return c.Self(), nil
	}

	token, err := bearerTokenFromHeader(strings.TrimSpace(r.Header.Get("Authorization")))
	if err != nil {
		return Principal{}, err
	}

	switch c.Mode {
	case ModeStaticToken:
		entry, ok := c.lookupToken(token)
		if !ok {
			return Principal{}, ErrInvalidBearerToken
		}

		identityType := IdentityTypeHuman
		globalScope := false
		if entry.Role == RoleService {
			identityType = IdentityTypeService
			globalScope = true
		}

		return Principal{
			Authenticated: true,
			AuthMode:      c.Mode,
			Subject:       entry.Subject,
			Role:          entry.Role,
			TokenID:       entry.TokenID,
			IdentityType:  identityType,
			GlobalScope:   globalScope,
		}, nil
	case ModeOIDCJWT:
		if c.oidc == nil {
			return Principal{}, ErrInvalidBearerToken
		}
		return c.oidc.authenticate(r.Context(), token)
	default:
		return Principal{}, ErrInvalidBearerToken
	}
}

func (c Config) Require(r *http.Request, allowedRoles ...string) (Principal, error) {
	principal, err := c.AuthenticateRequest(r)
	if err != nil {
		return Principal{}, err
	}
	if c.Mode == ModeDisabled {
		return principal, nil
	}
	if len(allowedRoles) == 0 {
		return principal, nil
	}

	for _, role := range allowedRoles {
		if principal.Role == role {
			return principal, nil
		}
	}
	return Principal{}, ErrInsufficientPermissions
}

func WithPrincipal(ctx context.Context, principal Principal) context.Context {
	return context.WithValue(ctx, principalContextKey{}, principal)
}

func PrincipalFromContext(ctx context.Context) (Principal, bool) {
	principal, ok := ctx.Value(principalContextKey{}).(Principal)
	return principal, ok
}

func StatusCode(err error) int {
	var accessErr *AccessError
	if errors.As(err, &accessErr) {
		return accessErr.Status
	}
	return http.StatusInternalServerError
}

func validRole(role string) bool {
	switch role {
	case RoleViewer, RoleOperator, RoleSecurityAdmin, RoleService:
		return true
	default:
		return false
	}
}

func bearerTokenFromHeader(value string) (string, error) {
	if value == "" {
		return "", ErrMissingBearerToken
	}

	parts := strings.Fields(value)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
		return "", ErrMalformedAuthorization
	}
	return parts[1], nil
}

func (c Config) lookupToken(token string) (TokenConfig, bool) {
	var match TokenConfig
	found := false
	for _, entry := range c.tokens {
		if len(entry.Token) != len(token) {
			continue
		}
		if subtle.ConstantTimeCompare([]byte(entry.Token), []byte(token)) == 1 {
			match = entry
			found = true
		}
	}
	return match, found
}

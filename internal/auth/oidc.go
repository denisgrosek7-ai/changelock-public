package auth

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultOIDCRoleClaim    = "groups"
	defaultOIDCSubjectClaim = "sub"
	defaultOIDCEmailClaim   = "email"
	defaultOIDCJWKSTTL      = 5 * time.Minute
	defaultOIDCFetchTimeout = 5 * time.Second
	defaultOIDCClockSkew    = time.Minute
)

type OIDCOptions struct {
	Issuer                   string
	Audiences                []string
	JWKSURL                  string
	RoleClaim                string
	TenantClaim              string
	EmailClaim               string
	SubjectClaim             string
	ClockSkew                time.Duration
	RequireTenantScope       bool
	AllowGlobalSecurityAdmin bool
	RoleBindings             map[string][]string
	HTTPClient               *http.Client
	Now                      func() time.Time
}

type oidcVerifier struct {
	issuer                   string
	audiences                []string
	roleClaim                string
	tenantClaim              string
	emailClaim               string
	subjectClaim             string
	clockSkew                time.Duration
	requireTenantScope       bool
	allowGlobalSecurityAdmin bool
	roleBindings             map[string]map[string]struct{}
	jwks                     *jwksCache
	now                      func() time.Time
}

type jwksCache struct {
	url       string
	client    *http.Client
	cacheTTL  time.Duration
	mu        sync.RWMutex
	fetchedAt time.Time
	keys      map[string]*rsa.PublicKey
}

type jwtHeader struct {
	Algorithm string `json:"alg"`
	KeyID     string `json:"kid"`
	Type      string `json:"typ,omitempty"`
}

type jwksDocument struct {
	Keys []jsonWebKey `json:"keys"`
}

type jsonWebKey struct {
	KeyType string `json:"kty"`
	KeyID   string `json:"kid"`
	Use     string `json:"use,omitempty"`
	N       string `json:"n,omitempty"`
	E       string `json:"e,omitempty"`
}

func ParseEnvConfig(getenv func(string) string) (Config, error) {
	if getenv == nil {
		return Config{}, errors.New("getenv is required")
	}

	mode := strings.TrimSpace(getenv("CHANGELOCK_AUTH_MODE"))
	if mode == "" {
		mode = ModeDisabled
	}

	if strings.EqualFold(mode, ModeOIDCJWT) {
		options, err := oidcOptionsFromEnv(getenv)
		if err != nil {
			return Config{}, err
		}
		return ParseOIDCConfig(options)
	}

	return ParseConfig(mode, getenv("CHANGELOCK_AUTH_TOKENS_JSON"))
}

func ParseOIDCConfig(options OIDCOptions) (Config, error) {
	issuer := strings.TrimSpace(options.Issuer)
	jwksURL := strings.TrimSpace(options.JWKSURL)
	if issuer == "" {
		return Config{}, errors.New("CHANGELOCK_OIDC_ISSUER is required when CHANGELOCK_AUTH_MODE=oidc-jwt")
	}
	if jwksURL == "" {
		return Config{}, errors.New("CHANGELOCK_OIDC_JWKS_URL is required when CHANGELOCK_AUTH_MODE=oidc-jwt")
	}

	audiences := make([]string, 0, len(options.Audiences))
	seenAudiences := map[string]struct{}{}
	for _, audience := range options.Audiences {
		audience = strings.TrimSpace(audience)
		if audience == "" {
			continue
		}
		if _, exists := seenAudiences[audience]; exists {
			continue
		}
		seenAudiences[audience] = struct{}{}
		audiences = append(audiences, audience)
	}
	if len(audiences) == 0 {
		return Config{}, errors.New("CHANGELOCK_OIDC_AUDIENCES is required when CHANGELOCK_AUTH_MODE=oidc-jwt")
	}

	roleClaim := strings.TrimSpace(options.RoleClaim)
	if roleClaim == "" {
		roleClaim = defaultOIDCRoleClaim
	}
	subjectClaim := strings.TrimSpace(options.SubjectClaim)
	if subjectClaim == "" {
		subjectClaim = defaultOIDCSubjectClaim
	}
	emailClaim := strings.TrimSpace(options.EmailClaim)
	if emailClaim == "" {
		emailClaim = defaultOIDCEmailClaim
	}
	tenantClaim := strings.TrimSpace(options.TenantClaim)
	if options.RequireTenantScope && tenantClaim == "" {
		return Config{}, errors.New("CHANGELOCK_AUTH_TENANT_CLAIM is required when CHANGELOCK_AUTH_REQUIRE_TENANT_SCOPE=true")
	}

	clockSkew := options.ClockSkew
	if clockSkew < 0 {
		return Config{}, errors.New("CHANGELOCK_OIDC_CLOCK_SKEW must be >= 0")
	}
	if clockSkew == 0 {
		clockSkew = defaultOIDCClockSkew
	}

	roleBindings := map[string]map[string]struct{}{}
	seenBindings := map[string]string{}
	for role, values := range options.RoleBindings {
		role = strings.TrimSpace(role)
		if !validRole(role) {
			return Config{}, errors.New("auth role binding has unsupported role: " + role)
		}
		if _, exists := roleBindings[role]; !exists {
			roleBindings[role] = map[string]struct{}{}
		}
		for _, value := range values {
			value = strings.TrimSpace(value)
			if value == "" {
				return Config{}, errors.New("auth role bindings require non-empty values")
			}
			if existing, exists := seenBindings[value]; exists && existing != role {
				return Config{}, errors.New("duplicate auth role binding configured: " + value)
			}
			seenBindings[value] = role
			roleBindings[role][value] = struct{}{}
		}
	}
	if len(seenBindings) == 0 {
		return Config{}, errors.New("CHANGELOCK_AUTH_ROLE_BINDINGS_JSON must contain at least one role binding when CHANGELOCK_AUTH_MODE=oidc-jwt")
	}

	httpClient := options.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultOIDCFetchTimeout}
	}
	now := options.Now
	if now == nil {
		now = time.Now
	}

	return Config{
		Mode: ModeOIDCJWT,
		oidc: &oidcVerifier{
			issuer:                   issuer,
			audiences:                audiences,
			roleClaim:                roleClaim,
			tenantClaim:              tenantClaim,
			emailClaim:               emailClaim,
			subjectClaim:             subjectClaim,
			clockSkew:                clockSkew,
			requireTenantScope:       options.RequireTenantScope,
			allowGlobalSecurityAdmin: options.AllowGlobalSecurityAdmin,
			roleBindings:             roleBindings,
			now:                      now,
			jwks: &jwksCache{
				url:      jwksURL,
				client:   httpClient,
				cacheTTL: defaultOIDCJWKSTTL,
				keys:     map[string]*rsa.PublicKey{},
			},
		},
	}, nil
}

func oidcOptionsFromEnv(getenv func(string) string) (OIDCOptions, error) {
	options := OIDCOptions{
		Issuer:       getenv("CHANGELOCK_OIDC_ISSUER"),
		JWKSURL:      getenv("CHANGELOCK_OIDC_JWKS_URL"),
		RoleClaim:    getenv("CHANGELOCK_AUTH_ROLE_CLAIM"),
		TenantClaim:  getenv("CHANGELOCK_AUTH_TENANT_CLAIM"),
		EmailClaim:   getenv("CHANGELOCK_AUTH_EMAIL_CLAIM"),
		SubjectClaim: getenv("CHANGELOCK_AUTH_SUBJECT_CLAIM"),
		Audiences:    splitCSV(getenv("CHANGELOCK_OIDC_AUDIENCES")),
	}

	var err error
	options.ClockSkew, err = durationEnv(getenv("CHANGELOCK_OIDC_CLOCK_SKEW"), defaultOIDCClockSkew)
	if err != nil {
		return OIDCOptions{}, errors.New("invalid CHANGELOCK_OIDC_CLOCK_SKEW: " + err.Error())
	}
	options.RequireTenantScope, err = boolEnv(getenv("CHANGELOCK_AUTH_REQUIRE_TENANT_SCOPE"))
	if err != nil {
		return OIDCOptions{}, errors.New("invalid CHANGELOCK_AUTH_REQUIRE_TENANT_SCOPE: " + err.Error())
	}
	options.AllowGlobalSecurityAdmin, err = boolEnv(getenv("CHANGELOCK_AUTH_ALLOW_GLOBAL_SECURITY_ADMIN"))
	if err != nil {
		return OIDCOptions{}, errors.New("invalid CHANGELOCK_AUTH_ALLOW_GLOBAL_SECURITY_ADMIN: " + err.Error())
	}

	bindings := strings.TrimSpace(getenv("CHANGELOCK_AUTH_ROLE_BINDINGS_JSON"))
	if bindings != "" {
		if err := json.Unmarshal([]byte(bindings), &options.RoleBindings); err != nil {
			return OIDCOptions{}, errors.New("invalid CHANGELOCK_AUTH_ROLE_BINDINGS_JSON: " + err.Error())
		}
	}

	return options, nil
}

func (v *oidcVerifier) authenticate(ctx context.Context, token string) (Principal, error) {
	claims, err := v.validateToken(ctx, token)
	if err != nil {
		return Principal{}, err
	}

	subject, ok := readRequiredStringClaim(claims, v.subjectClaim)
	if !ok {
		return Principal{}, ErrInvalidBearerToken
	}
	email, emailErr := readOptionalStringClaim(claims, v.emailClaim)
	if emailErr != nil {
		return Principal{}, ErrInvalidBearerToken
	}

	roleValues, roleErr := readStringClaims(claims, v.roleClaim)
	if roleErr != nil {
		return Principal{}, ErrInvalidBearerToken
	}
	role, identityType, err := v.resolveRole(roleValues)
	if err != nil {
		return Principal{}, err
	}

	tenantID := ""
	if v.tenantClaim != "" {
		var tenantErr error
		tenantID, tenantErr = readOptionalStringClaim(claims, v.tenantClaim)
		if tenantErr != nil {
			return Principal{}, ErrInvalidBearerToken
		}
	}

	globalScope := false
	if role == RoleService {
		globalScope = true
		tenantID = ""
	} else if v.requireTenantScope && tenantID == "" {
		if role == RoleSecurityAdmin && v.allowGlobalSecurityAdmin {
			globalScope = true
		} else {
			return Principal{}, ErrTenantScopeRequired
		}
	}

	return Principal{
		Authenticated: true,
		AuthMode:      ModeOIDCJWT,
		Subject:       subject,
		Role:          role,
		IdentityType:  identityType,
		Email:         email,
		TenantID:      tenantID,
		GlobalScope:   globalScope,
	}, nil
}

func (v *oidcVerifier) validateToken(ctx context.Context, token string) (map[string]any, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, ErrInvalidBearerToken
	}

	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, ErrInvalidBearerToken
	}
	var header jwtHeader
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		return nil, ErrInvalidBearerToken
	}
	header.Algorithm = strings.TrimSpace(header.Algorithm)
	header.KeyID = strings.TrimSpace(header.KeyID)
	hash, err := jwtHash(header.Algorithm)
	if err != nil {
		return nil, ErrInvalidBearerToken
	}

	key, err := v.jwks.key(ctx, header.KeyID)
	if err != nil {
		return nil, ErrInvalidBearerToken
	}
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, ErrInvalidBearerToken
	}
	digest := hashSignatureInput(hash, []byte(parts[0]+"."+parts[1]))
	if err := rsa.VerifyPKCS1v15(key, hash, digest, signature); err != nil {
		return nil, ErrInvalidBearerToken
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, ErrInvalidBearerToken
	}
	claims := map[string]any{}
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return nil, ErrInvalidBearerToken
	}
	if !v.validClaims(claims) {
		return nil, ErrInvalidBearerToken
	}
	return claims, nil
}

func (v *oidcVerifier) validClaims(claims map[string]any) bool {
	now := v.now().UTC()
	if issuer, ok := readRequiredStringClaim(claims, "iss"); !ok || issuer != v.issuer {
		return false
	}

	audiences, err := readStringClaims(claims, "aud")
	if err != nil || len(audiences) == 0 {
		return false
	}
	if !containsAny(audiences, v.audiences) {
		return false
	}

	expiry, ok := readNumericDateClaim(claims, "exp")
	if !ok || now.After(expiry.Add(v.clockSkew)) {
		return false
	}
	if notBefore, ok := readNumericDateClaim(claims, "nbf"); ok && now.Add(v.clockSkew).Before(notBefore) {
		return false
	}
	if issuedAt, ok := readNumericDateClaim(claims, "iat"); ok && now.Add(v.clockSkew).Before(issuedAt) {
		return false
	}
	return true
}

func (v *oidcVerifier) resolveRole(values []string) (string, string, error) {
	matched := map[string]struct{}{}
	for _, value := range values {
		for role, bindings := range v.roleBindings {
			if _, ok := bindings[value]; ok {
				matched[role] = struct{}{}
			}
		}
	}
	if len(matched) == 0 {
		return "", "", ErrNoMappedRole
	}
	if _, serviceMatched := matched[RoleService]; serviceMatched && len(matched) > 1 {
		return "", "", ErrNoMappedRole
	}
	if _, ok := matched[RoleService]; ok {
		return RoleService, IdentityTypeService, nil
	}
	for _, role := range []string{RoleSecurityAdmin, RoleOperator, RoleViewer} {
		if _, ok := matched[role]; ok {
			return role, IdentityTypeHuman, nil
		}
	}
	return "", "", ErrNoMappedRole
}

func (c *jwksCache) key(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	if key, ok := c.cachedKey(kid); ok {
		return key, nil
	}
	if err := c.refresh(ctx); err != nil {
		return nil, err
	}
	if key, ok := c.cachedKey(kid); ok {
		return key, nil
	}
	return nil, ErrInvalidBearerToken
}

func (c *jwksCache) cachedKey(kid string) (*rsa.PublicKey, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.keys) == 0 || time.Since(c.fetchedAt) > c.cacheTTL {
		return nil, false
	}
	if kid == "" && len(c.keys) == 1 {
		for _, key := range c.keys {
			return key, true
		}
	}
	key, ok := c.keys[kid]
	return key, ok
}

func (c *jwksCache) refresh(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url, nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("jwks endpoint returned %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return err
	}
	var document jwksDocument
	if err := json.Unmarshal(body, &document); err != nil {
		return err
	}
	keys := map[string]*rsa.PublicKey{}
	for _, key := range document.Keys {
		publicKey, ok := parseRSAPublicKey(key)
		if !ok {
			continue
		}
		keys[key.KeyID] = publicKey
	}
	if len(keys) == 0 {
		return errors.New("jwks contains no supported RSA signing keys")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.keys = keys
	c.fetchedAt = time.Now()
	return nil
}

func parseRSAPublicKey(key jsonWebKey) (*rsa.PublicKey, bool) {
	if key.KeyType != "RSA" {
		return nil, false
	}
	if strings.TrimSpace(key.Use) != "" && key.Use != "sig" {
		return nil, false
	}
	nBytes, err := base64.RawURLEncoding.DecodeString(strings.TrimSpace(key.N))
	if err != nil || len(nBytes) == 0 {
		return nil, false
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(strings.TrimSpace(key.E))
	if err != nil || len(eBytes) == 0 {
		return nil, false
	}

	n := new(big.Int).SetBytes(nBytes)
	e := 0
	for _, b := range eBytes {
		e = e<<8 + int(b)
	}
	if e == 0 {
		return nil, false
	}
	return &rsa.PublicKey{N: n, E: e}, true
}

func jwtHash(algorithm string) (crypto.Hash, error) {
	switch algorithm {
	case "RS256":
		return crypto.SHA256, nil
	case "RS384":
		return crypto.SHA384, nil
	case "RS512":
		return crypto.SHA512, nil
	default:
		return 0, errors.New("unsupported jwt algorithm")
	}
}

func hashSignatureInput(hash crypto.Hash, input []byte) []byte {
	switch hash {
	case crypto.SHA256:
		sum := sha256.Sum256(input)
		return sum[:]
	case crypto.SHA384:
		sum := sha512.Sum384(input)
		return sum[:]
	case crypto.SHA512:
		sum := sha512.Sum512(input)
		return sum[:]
	default:
		return nil
	}
}

func readRequiredStringClaim(claims map[string]any, field string) (string, bool) {
	value, err := readOptionalStringClaim(claims, field)
	if err != nil || value == "" {
		return "", false
	}
	return value, true
}

func readOptionalStringClaim(claims map[string]any, field string) (string, error) {
	field = strings.TrimSpace(field)
	if field == "" {
		return "", nil
	}
	value, ok := claims[field]
	if !ok || value == nil {
		return "", nil
	}
	text, ok := value.(string)
	if !ok {
		return "", errors.New("claim must be a string")
	}
	return strings.TrimSpace(text), nil
}

func readStringClaims(claims map[string]any, field string) ([]string, error) {
	field = strings.TrimSpace(field)
	if field == "" {
		return nil, nil
	}
	value, ok := claims[field]
	if !ok || value == nil {
		return nil, nil
	}
	switch typed := value.(type) {
	case string:
		typed = strings.TrimSpace(typed)
		if typed == "" {
			return nil, nil
		}
		return []string{typed}, nil
	case []any:
		values := make([]string, 0, len(typed))
		for _, item := range typed {
			text, ok := item.(string)
			if !ok {
				return nil, errors.New("claim array must contain strings")
			}
			text = strings.TrimSpace(text)
			if text != "" {
				values = append(values, text)
			}
		}
		return values, nil
	default:
		return nil, errors.New("claim must be a string or string array")
	}
}

func readNumericDateClaim(claims map[string]any, field string) (time.Time, bool) {
	value, ok := claims[field]
	if !ok || value == nil {
		return time.Time{}, false
	}
	switch typed := value.(type) {
	case float64:
		return time.Unix(int64(typed), 0).UTC(), true
	case json.Number:
		seconds, err := typed.Int64()
		if err != nil {
			return time.Time{}, false
		}
		return time.Unix(seconds, 0).UTC(), true
	case int64:
		return time.Unix(typed, 0).UTC(), true
	case int:
		return time.Unix(int64(typed), 0).UTC(), true
	case string:
		seconds, err := strconv.ParseInt(strings.TrimSpace(typed), 10, 64)
		if err != nil {
			return time.Time{}, false
		}
		return time.Unix(seconds, 0).UTC(), true
	default:
		return time.Time{}, false
	}
}

func splitCSV(raw string) []string {
	parts := []string{}
	for _, part := range strings.Split(raw, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			parts = append(parts, part)
		}
	}
	return parts
}

func durationEnv(raw string, fallback time.Duration) (time.Duration, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return fallback, nil
	}
	value, err := time.ParseDuration(raw)
	if err != nil {
		return 0, err
	}
	if value < 0 {
		return 0, errors.New("must be >= 0")
	}
	return value, nil
}

func boolEnv(raw string) (bool, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return false, nil
	}
	return strconv.ParseBool(raw)
}

func containsAny(values []string, expected []string) bool {
	allowed := map[string]struct{}{}
	for _, item := range expected {
		allowed[item] = struct{}{}
	}
	for _, item := range values {
		if _, ok := allowed[item]; ok {
			return true
		}
	}
	return false
}

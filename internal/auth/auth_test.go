package auth

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestParseConfigRejectsInvalidValues(t *testing.T) {
	if _, err := ParseConfig("bogus", "[]"); err == nil {
		t.Fatal("expected invalid mode error")
	}
	if _, err := ParseConfig(ModeStaticToken, `not-json`); err == nil {
		t.Fatal("expected invalid JSON error")
	}
	if _, err := ParseConfig(ModeStaticToken, `[{"token":"a","subject":"u","role":"viewer"},{"token":"a","subject":"v","role":"operator"}]`); err == nil {
		t.Fatal("expected duplicate token error")
	}
	if _, err := ParseConfig(ModeStaticToken, `[{"token":"a","subject":"u","role":"viewer","token_id":"dup"},{"token":"b","subject":"v","role":"operator","token_id":"dup"}]`); err == nil {
		t.Fatal("expected duplicate token_id error")
	}
	if _, err := ParseConfig(ModeStaticToken, `[{"token":"a","subject":"u","role":"bogus"}]`); err == nil {
		t.Fatal("expected invalid role error")
	}
}

func TestDisabledModeAllowsProtectedRoute(t *testing.T) {
	cfg := DisabledConfig()
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/events", nil)

	principal, err := cfg.Require(req, RoleViewer)
	if err != nil {
		t.Fatalf("Require() error = %v", err)
	}
	if !principal.Authenticated || principal.AuthMode != ModeDisabled || principal.Role != RoleSecurityAdmin {
		t.Fatalf("unexpected principal %#v", principal)
	}
}

func TestStaticTokenAuthentication(t *testing.T) {
	cfg, err := ParseConfig(ModeStaticToken, `[{"token":"viewer-demo-token","subject":"demo-viewer","role":"viewer","token_id":"viewer-demo"}]`)
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}

	tests := []struct {
		name   string
		header string
		err    error
	}{
		{name: "missing", err: ErrMissingBearerToken},
		{name: "malformed", header: "Token abc", err: ErrMalformedAuthorization},
		{name: "invalid", header: "Bearer wrong", err: ErrInvalidBearerToken},
		{name: "valid", header: "Bearer viewer-demo-token"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/v1/reports/events", nil)
			if tc.header != "" {
				req.Header.Set("Authorization", tc.header)
			}

			principal, requireErr := cfg.Require(req, RoleViewer)
			if tc.err != nil {
				if requireErr != tc.err {
					t.Fatalf("expected %v, got %v", tc.err, requireErr)
				}
				return
			}
			if requireErr != nil {
				t.Fatalf("Require() error = %v", requireErr)
			}
			if principal.Subject != "demo-viewer" || principal.Role != RoleViewer || principal.TokenID != "viewer-demo" {
				t.Fatalf("unexpected principal %#v", principal)
			}
		})
	}
}

func TestRequireRejectsInsufficientRole(t *testing.T) {
	cfg, err := ParseConfig(ModeStaticToken, `[{"token":"viewer-demo-token","subject":"demo-viewer","role":"viewer"}]`)
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/v1/exceptions", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")

	if _, err := cfg.Require(req, RoleSecurityAdmin); err != ErrInsufficientPermissions {
		t.Fatalf("expected forbidden error, got %v", err)
	}
}

func TestParseEnvConfigRejectsInvalidOIDCValues(t *testing.T) {
	base := map[string]string{
		"CHANGELOCK_AUTH_MODE":               ModeOIDCJWT,
		"CHANGELOCK_OIDC_ISSUER":             "https://issuer.example.com",
		"CHANGELOCK_OIDC_JWKS_URL":           "https://issuer.example.com/jwks.json",
		"CHANGELOCK_OIDC_AUDIENCES":          "changelock-ui",
		"CHANGELOCK_AUTH_ROLE_BINDINGS_JSON": `{"viewer":["changelock-viewers"]}`,
	}

	tests := []struct {
		name      string
		override  map[string]string
		wantError string
	}{
		{
			name:      "missing issuer",
			override:  map[string]string{"CHANGELOCK_OIDC_ISSUER": ""},
			wantError: "CHANGELOCK_OIDC_ISSUER",
		},
		{
			name:      "missing audiences",
			override:  map[string]string{"CHANGELOCK_OIDC_AUDIENCES": ""},
			wantError: "CHANGELOCK_OIDC_AUDIENCES",
		},
		{
			name:      "invalid role bindings json",
			override:  map[string]string{"CHANGELOCK_AUTH_ROLE_BINDINGS_JSON": `{`},
			wantError: "CHANGELOCK_AUTH_ROLE_BINDINGS_JSON",
		},
		{
			name:      "duplicate binding value",
			override:  map[string]string{"CHANGELOCK_AUTH_ROLE_BINDINGS_JSON": `{"viewer":["same"],"operator":["same"]}`},
			wantError: "duplicate auth role binding",
		},
		{
			name:      "tenant scope requires tenant claim",
			override:  map[string]string{"CHANGELOCK_AUTH_REQUIRE_TENANT_SCOPE": "true"},
			wantError: "CHANGELOCK_AUTH_TENANT_CLAIM",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			env := map[string]string{}
			for key, value := range base {
				env[key] = value
			}
			for key, value := range tc.override {
				env[key] = value
			}
			if _, err := ParseEnvConfig(func(key string) string { return env[key] }); err == nil || !containsError(err, tc.wantError) {
				t.Fatalf("expected error containing %q, got %v", tc.wantError, err)
			}
		})
	}
}

func TestOIDCJWTAuthenticationAndTenantMapping(t *testing.T) {
	server, signer := newOIDCTestJWKS(t)
	cfg, err := ParseOIDCConfig(OIDCOptions{
		Issuer:                   server.URL,
		JWKSURL:                  server.URL + "/jwks.json",
		Audiences:                []string{"changelock-ui"},
		RoleClaim:                "groups",
		TenantClaim:              "tenant_id",
		RequireTenantScope:       true,
		AllowGlobalSecurityAdmin: true,
		RoleBindings: map[string][]string{
			RoleViewer:        {"changelock-viewers"},
			RoleOperator:      {"changelock-operators"},
			RoleSecurityAdmin: {"changelock-admins"},
			RoleService:       {"changelock-services"},
		},
		HTTPClient: server.Client(),
		Now:        time.Now,
	})
	if err != nil {
		t.Fatalf("ParseOIDCConfig() error = %v", err)
	}

	tests := []struct {
		name        string
		claims      map[string]any
		allowedRole string
		wantRole    string
		wantTenant  string
		wantType    string
		wantGlobal  bool
		wantErr     error
	}{
		{
			name:        "viewer",
			claims:      map[string]any{"sub": "viewer@example.com", "groups": []string{"changelock-viewers"}, "tenant_id": "acme"},
			allowedRole: RoleViewer,
			wantRole:    RoleViewer,
			wantTenant:  "acme",
			wantType:    IdentityTypeHuman,
		},
		{
			name:        "operator",
			claims:      map[string]any{"sub": "operator@example.com", "groups": []string{"changelock-operators", "changelock-viewers"}, "tenant_id": "acme"},
			allowedRole: RoleOperator,
			wantRole:    RoleOperator,
			wantTenant:  "acme",
			wantType:    IdentityTypeHuman,
		},
		{
			name:        "global security admin",
			claims:      map[string]any{"sub": "admin@example.com", "groups": []string{"changelock-admins"}},
			allowedRole: RoleSecurityAdmin,
			wantRole:    RoleSecurityAdmin,
			wantType:    IdentityTypeHuman,
			wantGlobal:  true,
		},
		{
			name:        "service identity",
			claims:      map[string]any{"sub": "policy-engine", "groups": []string{"changelock-services"}},
			allowedRole: RoleService,
			wantRole:    RoleService,
			wantType:    IdentityTypeService,
			wantGlobal:  true,
		},
		{
			name:        "missing mapped role",
			claims:      map[string]any{"sub": "unknown@example.com", "groups": []string{"other-group"}, "tenant_id": "acme"},
			allowedRole: RoleViewer,
			wantErr:     ErrNoMappedRole,
		},
		{
			name:        "tenant required for non-admin",
			claims:      map[string]any{"sub": "viewer@example.com", "groups": []string{"changelock-viewers"}},
			allowedRole: RoleViewer,
			wantErr:     ErrTenantScopeRequired,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			token := signer.token(t, tc.claims)
			req := httptest.NewRequest(http.MethodGet, "/v1/reports/events", nil)
			req.Header.Set("Authorization", "Bearer "+token)

			principal, err := cfg.Require(req, tc.allowedRole)
			if tc.wantErr != nil {
				if err != tc.wantErr {
					t.Fatalf("expected %v, got %v", tc.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Require() error = %v", err)
			}
			if principal.Role != tc.wantRole || principal.TenantID != tc.wantTenant || principal.IdentityType != tc.wantType || principal.GlobalScope != tc.wantGlobal {
				t.Fatalf("unexpected principal %#v", principal)
			}
		})
	}
}

func TestOIDCJWTRejectsInvalidTokenAndDoesNotFallBack(t *testing.T) {
	server, signer := newOIDCTestJWKS(t)
	cfg, err := ParseOIDCConfig(OIDCOptions{
		Issuer:       server.URL,
		JWKSURL:      server.URL + "/jwks.json",
		Audiences:    []string{"changelock-ui"},
		RoleClaim:    "groups",
		RoleBindings: map[string][]string{RoleViewer: {"changelock-viewers"}},
		HTTPClient:   server.Client(),
		Now:          time.Now,
	})
	if err != nil {
		t.Fatalf("ParseOIDCConfig() error = %v", err)
	}

	validToken := signer.token(t, map[string]any{"sub": "viewer@example.com", "groups": []string{"changelock-viewers"}})
	invalidAudience := signer.token(t, map[string]any{"sub": "viewer@example.com", "aud": []string{"wrong-audience"}, "groups": []string{"changelock-viewers"}})

	tests := []struct {
		name   string
		header string
		err    error
	}{
		{name: "opaque static token", header: "Bearer viewer-demo-token", err: ErrInvalidBearerToken},
		{name: "wrong audience", header: "Bearer " + invalidAudience, err: ErrInvalidBearerToken},
		{name: "valid jwt", header: "Bearer " + validToken},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/v1/reports/events", nil)
			req.Header.Set("Authorization", tc.header)
			_, err := cfg.Require(req, RoleViewer)
			if tc.err != nil {
				if err != tc.err {
					t.Fatalf("expected %v, got %v", tc.err, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Require() error = %v", err)
			}
		})
	}
}

type oidcTestSigner struct {
	privateKey *rsa.PrivateKey
	issuer     string
	keyID      string
}

func newOIDCTestJWKS(t *testing.T) (*httptest.Server, oidcTestSigner) {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("GenerateKey() error = %v", err)
	}
	signer := oidcTestSigner{
		privateKey: privateKey,
		keyID:      "test-key",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]any{
			"keys": []map[string]any{{
				"kty": "RSA",
				"kid": signer.keyID,
				"use": "sig",
				"n":   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.N.Bytes()),
				"e":   base64.RawURLEncoding.EncodeToString(bigEndianBytes(privateKey.PublicKey.E)),
			}},
		}
		_ = json.NewEncoder(w).Encode(payload)
	}))
	signer.issuer = server.URL
	t.Cleanup(server.Close)
	return server, signer
}

func (s oidcTestSigner) token(t *testing.T, claims map[string]any) string {
	t.Helper()

	headerBytes, err := json.Marshal(map[string]any{
		"alg": "RS256",
		"kid": s.keyID,
		"typ": "JWT",
	})
	if err != nil {
		t.Fatalf("marshal header: %v", err)
	}

	payload := map[string]any{
		"iss": s.issuer,
		"aud": []string{"changelock-ui"},
		"exp": time.Now().Add(5 * time.Minute).Unix(),
		"iat": time.Now().Add(-time.Minute).Unix(),
	}
	for key, value := range claims {
		payload[key] = value
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}

	encodedHeader := base64.RawURLEncoding.EncodeToString(headerBytes)
	encodedPayload := base64.RawURLEncoding.EncodeToString(payloadBytes)
	signingInput := encodedHeader + "." + encodedPayload
	sum := sha256.Sum256([]byte(signingInput))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, sum[:])
	if err != nil {
		t.Fatalf("SignPKCS1v15() error = %v", err)
	}
	return signingInput + "." + base64.RawURLEncoding.EncodeToString(signature)
}

func bigEndianBytes(value int) []byte {
	if value == 0 {
		return []byte{0}
	}
	buf := []byte{}
	for value > 0 {
		buf = append([]byte{byte(value & 0xff)}, buf...)
		value >>= 8
	}
	return buf
}

func containsError(err error, want string) bool {
	return err != nil && want != "" && strings.Contains(err.Error(), want)
}

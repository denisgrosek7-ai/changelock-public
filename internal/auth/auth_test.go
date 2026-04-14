package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
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

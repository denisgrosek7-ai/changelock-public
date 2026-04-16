package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
)

func TestAuthMeReturnsOIDCPrincipal(t *testing.T) {
	cfg, signer := newOIDCHandlerConfig(t, true, true)
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", cfg)

	req := httptest.NewRequest(http.MethodGet, "/v1/auth/me", nil)
	req.Header.Set("Authorization", "Bearer "+signer.token(t, map[string]any{
		"sub":       "viewer@example.com",
		"email":     "viewer@example.com",
		"groups":    []string{"changelock-viewers"},
		"tenant_id": "acme",
	}))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response authInfoResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode auth me response: %v", err)
	}
	if response.AuthMode != auth.ModeOIDCJWT || response.Role != auth.RoleViewer || response.TenantID != "acme" || response.IdentityType != auth.IdentityTypeHuman || response.Email != "viewer@example.com" {
		t.Fatalf("unexpected auth me response %#v", response)
	}
}

func TestTenantScopedJWTFiltersReportsAndAnalytics(t *testing.T) {
	store := audit.NewMemoryStore()
	mustIngestEnterpriseEvent(t, store, audit.Event{
		Component:   "deploy-gate",
		EventType:   audit.EventTypeDeployGateDecision,
		Decision:    audit.DecisionDeny,
		TenantID:    "acme",
		Environment: "prod",
		Repo:        "my-org/acme-app",
		Reasons:     []string{"workflow mismatch"},
	})
	mustIngestEnterpriseEvent(t, store, audit.Event{
		Component:   "deploy-gate",
		EventType:   audit.EventTypeDeployGateDecision,
		Decision:    audit.DecisionDeny,
		TenantID:    "globex",
		Environment: "prod",
		Repo:        "my-org/globex-app",
		Reasons:     []string{"workflow mismatch"},
	})

	cfg, signer := newOIDCHandlerConfig(t, true, true)
	handler := newHandlerWithAuth(store, "memory", cfg)

	eventsReq := httptest.NewRequest(http.MethodGet, "/v1/reports/events?limit=10", nil)
	eventsReq.Header.Set("Authorization", "Bearer "+signer.token(t, map[string]any{
		"sub":       "viewer@example.com",
		"groups":    []string{"changelock-viewers"},
		"tenant_id": "acme",
	}))
	eventsRec := httptest.NewRecorder()
	handler.ServeHTTP(eventsRec, eventsReq)
	if eventsRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", eventsRec.Code, eventsRec.Body.String())
	}

	var events eventsResponse
	if err := json.NewDecoder(eventsRec.Body).Decode(&events); err != nil {
		t.Fatalf("decode events response: %v", err)
	}
	if len(events.Events) != 1 || events.Events[0].TenantID != "acme" {
		t.Fatalf("unexpected tenant-scoped events %#v", events)
	}

	rejectReq := httptest.NewRequest(http.MethodGet, "/v1/reports/events?tenant_id=globex", nil)
	rejectReq.Header.Set("Authorization", eventsReq.Header.Get("Authorization"))
	rejectRec := httptest.NewRecorder()
	handler.ServeHTTP(rejectRec, rejectReq)
	if rejectRec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", rejectRec.Code, rejectRec.Body.String())
	}

	analyticsReq := httptest.NewRequest(http.MethodGet, "/v1/analytics/trends?window_days=30&granularity=day", nil)
	analyticsReq.Header.Set("Authorization", eventsReq.Header.Get("Authorization"))
	analyticsRec := httptest.NewRecorder()
	handler.ServeHTTP(analyticsRec, analyticsReq)
	if analyticsRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", analyticsRec.Code, analyticsRec.Body.String())
	}

	var trends audit.TrendsResponse
	if err := json.NewDecoder(analyticsRec.Body).Decode(&trends); err != nil {
		t.Fatalf("decode trends response: %v", err)
	}
	if trends.AppliedFilters["tenant_id"] != "acme" {
		t.Fatalf("expected tenant_id filter injection, got %#v", trends.AppliedFilters)
	}
}

func TestTenantScopedSecurityAdminEnforcesExceptionAndVulnerabilityScope(t *testing.T) {
	store := audit.NewMemoryStore()
	if _, err := store.CreateException(t.Context(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-GLOBEX",
		ExceptionType: audit.ExceptionTypeBreakGlass,
		TenantID:      "globex",
		Environment:   "prod",
		Reason:        "globex only",
		TicketID:      "INC-GLOBEX",
		ApprovedBy:    "security@example.com",
		TTLHours:      1,
	}); err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}
	mustIngestEnterpriseEvent(t, store, audit.Event{
		Component:   "runtime-agent",
		EventType:   audit.EventTypeRuntimeDriftResult,
		Decision:    audit.DecisionAllow,
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
		Workload:    "checkout",
		Repo:        "my-org/acme-app",
		Image:       "ghcr.io/example/acme:1.0.0",
		Digest:      "sha256:acme-digest",
	})
	mustIngestEnterpriseEvent(t, store, audit.Event{
		Component:   "runtime-agent",
		EventType:   audit.EventTypeRuntimeDriftResult,
		Decision:    audit.DecisionAllow,
		TenantID:    "globex",
		Environment: "prod",
		Namespace:   "globex-prod",
		Workload:    "checkout",
		Repo:        "my-org/globex-app",
		Image:       "ghcr.io/example/globex:1.0.0",
		Digest:      "sha256:globex-digest",
	})
	if _, err := store.RecordVulnerabilityScan(t.Context(), audit.VulnerabilityScanRequest{
		ImageDigest: "sha256:acme-digest",
		ImageRef:    "ghcr.io/example/acme:1.0.0",
		Scanner:     "trivy",
		StartedAt:   time.Now().UTC(),
		CompletedAt: ptrTimeOIDC(time.Now().UTC()),
		Status:      audit.VulnerabilityScanStatusCompleted,
		Findings: []audit.VulnerabilityFindingInput{{
			CVEID:       "CVE-2026-7001",
			Severity:    "HIGH",
			PackageName: "openssl",
		}},
	}); err != nil {
		t.Fatalf("RecordVulnerabilityScan() error = %v", err)
	}

	cfg, signer := newOIDCHandlerConfig(t, true, true)
	handler := newHandlerWithAuth(store, "memory", cfg)
	adminHeader := "Bearer " + signer.token(t, map[string]any{
		"sub":       "admin@example.com",
		"groups":    []string{"changelock-admins"},
		"tenant_id": "acme",
	})

	createReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions", bytes.NewBufferString(`{
	  "exception_id":"EX-ACME",
	  "exception_type":"BREAK_GLASS",
	  "environment":"prod",
	  "reason":"tenant injected",
	  "ticket_id":"INC-ACME",
	  "approved_by":"security@example.com",
	  "ttl_hours":1
	}`))
	createReq.Header.Set("Authorization", adminHeader)
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()
	handler.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", createRec.Code, createRec.Body.String())
	}

	var created exceptionResponse
	if err := json.NewDecoder(createRec.Body).Decode(&created); err != nil {
		t.Fatalf("decode exception response: %v", err)
	}
	if created.Exception.TenantID != "acme" {
		t.Fatalf("expected tenant injection, got %#v", created.Exception)
	}

	mismatchReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions", bytes.NewBufferString(`{
	  "exception_id":"EX-GLOBEX-NEW",
	  "exception_type":"BREAK_GLASS",
	  "tenant_id":"globex",
	  "environment":"prod",
	  "reason":"tenant mismatch",
	  "ticket_id":"INC-GLOBEX",
	  "approved_by":"security@example.com",
	  "ttl_hours":1
	}`))
	mismatchReq.Header.Set("Authorization", adminHeader)
	mismatchReq.Header.Set("Content-Type", "application/json")
	mismatchRec := httptest.NewRecorder()
	handler.ServeHTTP(mismatchRec, mismatchReq)
	if mismatchRec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", mismatchRec.Code, mismatchRec.Body.String())
	}

	revokeReq := httptest.NewRequest(http.MethodDelete, "/v1/exceptions/EX-GLOBEX", nil)
	revokeReq.Header.Set("Authorization", adminHeader)
	revokeRec := httptest.NewRecorder()
	handler.ServeHTTP(revokeRec, revokeReq)
	if revokeRec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", revokeRec.Code, revokeRec.Body.String())
	}

	vulnReq := httptest.NewRequest(http.MethodGet, "/v1/vulnerabilities/active?limit=10", nil)
	vulnReq.Header.Set("Authorization", adminHeader)
	vulnRec := httptest.NewRecorder()
	handler.ServeHTTP(vulnRec, vulnReq)
	if vulnRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", vulnRec.Code, vulnRec.Body.String())
	}

	var findings vulnerabilitiesResponse
	if err := json.NewDecoder(vulnRec.Body).Decode(&findings); err != nil {
		t.Fatalf("decode vulnerabilities response: %v", err)
	}
	if len(findings.Findings) != 1 || findings.Findings[0].ImageDigest != "sha256:acme-digest" {
		t.Fatalf("unexpected tenant-scoped vulnerabilities %#v", findings)
	}
}

func TestGlobalSecurityAdminCanReadAcrossTenants(t *testing.T) {
	store := audit.NewMemoryStore()
	mustIngestEnterpriseEvent(t, store, audit.Event{
		Component: "deploy-gate",
		EventType: audit.EventTypeDeployGateDecision,
		Decision:  audit.DecisionAllow,
		TenantID:  "globex",
	})

	cfg, signer := newOIDCHandlerConfig(t, true, true)
	handler := newHandlerWithAuth(store, "memory", cfg)

	req := httptest.NewRequest(http.MethodGet, "/v1/reports/events?tenant_id=globex", nil)
	req.Header.Set("Authorization", "Bearer "+signer.token(t, map[string]any{
		"sub":    "admin@example.com",
		"groups": []string{"changelock-admins"},
	}))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
}

type oidcHandlerSigner struct {
	privateKey *rsa.PrivateKey
	issuer     string
	keyID      string
}

func newOIDCHandlerConfig(t *testing.T, requireTenantScope bool, allowGlobalAdmin bool) (auth.Config, oidcHandlerSigner) {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("GenerateKey() error = %v", err)
	}
	signer := oidcHandlerSigner{
		privateKey: privateKey,
		keyID:      "handler-test-key",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"keys": []map[string]any{{
				"kty": "RSA",
				"kid": signer.keyID,
				"use": "sig",
				"n":   base64.RawURLEncoding.EncodeToString(privateKey.PublicKey.N.Bytes()),
				"e":   base64.RawURLEncoding.EncodeToString(bigEndianOIDC(privateKey.PublicKey.E)),
			}},
		})
	}))
	t.Cleanup(server.Close)
	signer.issuer = server.URL

	cfg, err := auth.ParseOIDCConfig(auth.OIDCOptions{
		Issuer:                   server.URL,
		JWKSURL:                  server.URL + "/jwks.json",
		Audiences:                []string{"changelock-ui"},
		RoleClaim:                "groups",
		TenantClaim:              "tenant_id",
		EmailClaim:               "email",
		SubjectClaim:             "sub",
		RequireTenantScope:       requireTenantScope,
		AllowGlobalSecurityAdmin: allowGlobalAdmin,
		RoleBindings: map[string][]string{
			auth.RoleViewer:        {"changelock-viewers"},
			auth.RoleOperator:      {"changelock-operators"},
			auth.RoleSecurityAdmin: {"changelock-admins"},
			auth.RoleService:       {"changelock-services"},
		},
		HTTPClient: server.Client(),
		Now:        time.Now,
	})
	if err != nil {
		t.Fatalf("ParseOIDCConfig() error = %v", err)
	}
	return cfg, signer
}

func (s oidcHandlerSigner) token(t *testing.T, claims map[string]any) string {
	t.Helper()

	header, err := json.Marshal(map[string]any{"alg": "RS256", "kid": s.keyID, "typ": "JWT"})
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
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}
	encodedHeader := base64.RawURLEncoding.EncodeToString(header)
	encodedBody := base64.RawURLEncoding.EncodeToString(body)
	signingInput := encodedHeader + "." + encodedBody
	sum := sha256.Sum256([]byte(signingInput))
	signature, err := rsa.SignPKCS1v15(rand.Reader, s.privateKey, crypto.SHA256, sum[:])
	if err != nil {
		t.Fatalf("SignPKCS1v15() error = %v", err)
	}
	return signingInput + "." + base64.RawURLEncoding.EncodeToString(signature)
}

func bigEndianOIDC(value int) []byte {
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

func mustIngestEnterpriseEvent(t *testing.T, store audit.Store, event audit.Event) {
	t.Helper()
	if _, err := store.Ingest(t.Context(), event); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}
}

func ptrTimeOIDC(value time.Time) *time.Time {
	return &value
}

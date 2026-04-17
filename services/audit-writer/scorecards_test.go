package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

func TestScorecardEndpointReturnsDerivedPosture(t *testing.T) {
	store := audit.NewMemoryStore()
	seedTrustScorecardStore(t, store)
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/scorecards?tenant_id=acme&repo=demo/app", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var scorecard audit.TrustScorecard
	if err := json.NewDecoder(rec.Body).Decode(&scorecard); err != nil {
		t.Fatalf("decode scorecard: %v", err)
	}
	if scorecard.ScopeType != "repository" || scorecard.ScopeRef == "" {
		t.Fatalf("unexpected scope %#v", scorecard)
	}
	if len(scorecard.Metrics) == 0 {
		t.Fatalf("expected metrics, got %#v", scorecard)
	}
}

func TestAuditReportsEndpointSupportsHTML(t *testing.T) {
	store := audit.NewMemoryStore()
	seedTrustScorecardStore(t, store)
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodPost, "/v1/audit/reports", bytes.NewBufferString(`{"tenant_id":"acme","repo":"demo/app","format":"html"}`))
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	if got := rec.Header().Get("Content-Type"); !strings.Contains(got, "text/html") {
		t.Fatalf("expected html content type, got %q", got)
	}
	if !strings.Contains(rec.Body.String(), "ChangeLock Hardening Audit") {
		t.Fatalf("expected html report body, got %s", rec.Body.String())
	}
}

func TestPublishedTrustViewHonorsPublicationMode(t *testing.T) {
	store := audit.NewMemoryStore()
	seedTrustScorecardStore(t, store)
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/trust/published?tenant_id=acme&repo=demo/app", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404 when publication disabled, got %d: %s", rec.Code, rec.Body.String())
	}

	t.Setenv("CHANGELOCK_TRUST_PUBLICATION_MODE", "preview")
	handler = newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 when publication enabled, got %d: %s", rec.Code, rec.Body.String())
	}
	var view audit.PublishedTrustView
	if err := json.NewDecoder(rec.Body).Decode(&view); err != nil {
		t.Fatalf("decode published view: %v", err)
	}
	if view.OverallGrade == "" || len(view.Badges) == 0 {
		t.Fatalf("expected sanitized public view, got %#v", view)
	}
}

func TestAuditExportRejectsPublicViewWhenDisabled(t *testing.T) {
	store := audit.NewMemoryStore()
	seedTrustScorecardStore(t, store)
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodPost, "/v1/audit/exports", bytes.NewBufferString(`{"tenant_id":"acme","repo":"demo/app","include_public_view":true}`))
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusConflict {
		t.Fatalf("expected 409, got %d: %s", rec.Code, rec.Body.String())
	}
}

func seedTrustScorecardStore(t *testing.T, store audit.Store) {
	t.Helper()
	ctx := context.Background()
	now := time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC)

	if _, err := store.Ingest(ctx, audit.Event{
		Component: "attestation-verifier",
		EventType: audit.EventTypeArtifactVerificationResult,
		Decision:  audit.DecisionAllow,
		TenantID:  "acme",
		Repo:      "demo/app",
		Timestamp: now,
		Digest:    "sha256:demo",
		VerifierSummary: &audit.VerifierSummary{
			SignatureValid:   true,
			AttestationValid: true,
		},
		Evidence: &audit.Evidence{
			VerificationState: "verified",
			Artifact: &audit.ArtifactEvidence{
				SignerIdentity:  "https://github.com/example/demo/.github/workflows/release.yml@refs/heads/main",
				Repository:      "demo/app",
				Digest:          "sha256:demo",
				SBOMArtifactRef: "oci://ghcr.io/example/demo:sbom",
			},
		},
	}); err != nil {
		t.Fatalf("seed artifact event: %v", err)
	}

	if _, err := store.Ingest(ctx, audit.Event{
		Component:        "policy-engine",
		EventType:        audit.EventTypePolicyDecision,
		Decision:         audit.DecisionAllow,
		TenantID:         "acme",
		Repo:             "demo/app",
		Timestamp:        now,
		PolicyVersion:    "2026.04.17",
		PolicyBundleID:   "bundle-a",
		PolicyBundleHash: "sha256:bundle",
	}); err != nil {
		t.Fatalf("seed policy event: %v", err)
	}

	if _, err := store.CreateException(ctx, audit.ExceptionCreateRequest{
		ExceptionID:   "EX-2026-001",
		ExceptionType: audit.ExceptionTypeBreakGlass,
		TenantID:      "acme",
		Repo:          "demo/app",
		Reason:        "incident mitigation",
		TicketID:      "INC-123",
		ApprovedBy:    "admin@example.com",
		ExpiresAt:     trustTimePointer(now.Add(24 * time.Hour)),
	}); err != nil {
		t.Fatalf("seed exception: %v", err)
	}

	if _, err := store.CreateVEXStatement(ctx, internalvex.CreateRequest{
		SourceFormat:    internalvex.SourceFormatAPI,
		VulnerabilityID: "CVE-2026-0001",
		Scope: internalvex.Scope{
			ImageDigest: "sha256:demo",
			Repo:        "demo/app",
			TenantID:    "acme",
		},
		Status:        internalvex.StatusNotAffected,
		Justification: "runtime path not reachable",
	}, "security-admin"); err != nil {
		t.Fatalf("seed vex statement: %v", err)
	}
}

func trustTimePointer(value time.Time) *time.Time {
	return &value
}

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/guidance"
)

func TestAIGuidanceEndpointsReturnGroupedReadOnlyGuidance(t *testing.T) {
	t.Setenv("CHANGELOCK_AI_GUIDANCE_MODE", guidance.ModeLocalTemplate)
	store := audit.NewMemoryStore()
	seedTrustScorecardStore(t, store)
	seedAIGuidanceVulnerabilityStore(t, store)
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/ai/guidance?tenant_id=acme&repo=demo/app", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response guidance.Response
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode guidance response: %v", err)
	}
	if len(response.Items) == 0 || response.Summary.TotalItems == 0 {
		t.Fatalf("expected guidance items, got %#v", response)
	}
	itemID := response.Items[0].ID

	itemReq := httptest.NewRequest(http.MethodGet, "/v1/ai/guidance/"+itemID+"?tenant_id=acme&repo=demo/app", nil)
	itemReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	itemRec := httptest.NewRecorder()
	handler.ServeHTTP(itemRec, itemReq)
	if itemRec.Code != http.StatusOK {
		t.Fatalf("expected 200 for guidance item, got %d: %s", itemRec.Code, itemRec.Body.String())
	}
}

func TestAIVEXDraftEndpointRequiresOperatorAndReturnsReviewOnlyDraft(t *testing.T) {
	t.Setenv("CHANGELOCK_AI_GUIDANCE_MODE", guidance.ModeLocalTemplate)
	store := audit.NewMemoryStore()
	seedTrustScorecardStore(t, store)
	seedAIGuidanceVulnerabilityStore(t, store)
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	viewerReq := httptest.NewRequest(http.MethodPost, "/v1/ai/vex-drafts", bytes.NewBufferString(`{"tenant_id":"acme","repo":"demo/app","image_digest":"sha256:demo","cve_id":"CVE-2026-9001"}`))
	viewerReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	viewerReq.Header.Set("Content-Type", "application/json")
	viewerRec := httptest.NewRecorder()
	handler.ServeHTTP(viewerRec, viewerReq)
	if viewerRec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", viewerRec.Code, viewerRec.Body.String())
	}

	operatorReq := httptest.NewRequest(http.MethodPost, "/v1/ai/vex-drafts", bytes.NewBufferString(`{"tenant_id":"acme","repo":"demo/app","image_digest":"sha256:demo","cve_id":"CVE-2026-9001"}`))
	operatorReq.Header.Set("Authorization", "Bearer operator-demo-token")
	operatorReq.Header.Set("Content-Type", "application/json")
	operatorRec := httptest.NewRecorder()
	handler.ServeHTTP(operatorRec, operatorReq)
	if operatorRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", operatorRec.Code, operatorRec.Body.String())
	}

	var response aiVEXDraftResponse
	if err := json.NewDecoder(operatorRec.Body).Decode(&response); err != nil {
		t.Fatalf("decode VEX draft response: %v", err)
	}
	if response.Draft == nil || response.Draft.CandidateStatus != "under_investigation" {
		t.Fatalf("expected review-only VEX draft, got %#v", response)
	}
	if !response.Draft.AdvisoryOnly || !response.Draft.RequiresHumanReview {
		t.Fatalf("expected advisory-only draft, got %#v", response.Draft)
	}
}

func TestAIBreakGlassGuidanceEndpointIsScopedAndReadOnly(t *testing.T) {
	t.Setenv("CHANGELOCK_AI_GUIDANCE_MODE", guidance.ModeLocalTemplate)
	store := audit.NewMemoryStore()
	seedTrustScorecardStore(t, store)
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodPost, "/v1/ai/break-glass-guidance", bytes.NewBufferString(`{"tenant_id":"acme","exception_id":"EX-2026-001"}`))
	req.Header.Set("Authorization", "Bearer operator-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response aiBreakGlassGuidanceResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode break-glass response: %v", err)
	}
	if response.Guidance == nil || response.Guidance.ScopeExplanation == "" {
		t.Fatalf("expected break-glass guidance, got %#v", response)
	}

	existing, err := store.GetException(context.Background(), "EX-2026-001")
	if err != nil {
		t.Fatalf("GetException() error = %v", err)
	}
	if existing.Status != audit.ExceptionStatusApproved {
		t.Fatalf("expected authoritative exception state unchanged, got %#v", existing)
	}
}

func seedAIGuidanceVulnerabilityStore(t *testing.T, store audit.Store) {
	t.Helper()
	now := time.Date(2026, 4, 17, 13, 0, 0, 0, time.UTC)
	if _, err := store.RecordVulnerabilityScan(context.Background(), audit.VulnerabilityScanRequest{
		ImageDigest: "sha256:demo",
		ImageRef:    "ghcr.io/example/demo@sha256:demo",
		Scanner:     "trivy",
		ScanMode:    audit.VulnerabilityScanModeOnDemand,
		StartedAt:   now,
		CompletedAt: trustTimePointer(now.Add(2 * time.Minute)),
		Status:      audit.VulnerabilityScanStatusCompleted,
		Findings: []audit.VulnerabilityFindingInput{
			{
				CVEID:          "CVE-2026-9001",
				Severity:       "HIGH",
				PackageName:    "openssl",
				PackageVersion: "3.0.14-r0",
				FixedVersion:   "3.0.15-r0",
			},
		},
	}); err != nil {
		t.Fatalf("RecordVulnerabilityScan() error = %v", err)
	}
}

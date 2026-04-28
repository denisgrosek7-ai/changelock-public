package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestDeveloperEcosystemValAHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/developer-ecosystem/vala/status?tenant_id=acme&environment=prod",
		"/v1/developer-ecosystem/vala/proofs?tenant_id=acme&environment=prod",
	}

	for _, path := range paths {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200 for %s, got %d: %s", path, rec.Code, rec.Body.String())
		}
	}

	req := httptest.NewRequest(http.MethodPost, "/v1/developer-ecosystem/vala/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestDeveloperEcosystemValAProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/vala/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.DeveloperEcosystemValAStateActive ||
		response.DependencyState != operability.DeveloperEcosystemValADependencyStateActive ||
		response.Val0CurrentState != operability.DeveloperEcosystemVal0StateActive ||
		response.Val0Point8State != operability.DeveloperEcosystemPoint8StateNotComplete {
		t.Fatalf("expected healthy Val 0 dependency for developer Val A, got %#v", response)
	}
	if response.Point8State != operability.DeveloperEcosystemPoint8StateNotComplete ||
		response.IDEBaselineState != operability.DeveloperEcosystemValAIDEBaselineStateActive ||
		response.TrustFeedbackState != operability.DeveloperEcosystemValATrustFeedbackStateActive ||
		response.CAVIVEXContextState != operability.DeveloperEcosystemValACAVIVEXStateActive ||
		response.LocalAdvisoryState != operability.DeveloperEcosystemValALocalAdvisoryStateActive ||
		response.ValidationHarnessState != operability.DeveloperEcosystemValAValidationHarnessStateActive ||
		response.MockVerificationState != operability.DeveloperEcosystemValAMockVerificationStateActive ||
		response.InspectExplainState != operability.DeveloperEcosystemValAInspectExplainStateActive ||
		response.DegradedModeState != operability.DeveloperEcosystemValADegradedModeStateActive ||
		response.NoOverclaimState != operability.DeveloperEcosystemValANoOverclaimStateActive {
		t.Fatalf("expected active developer Val A contract states with point 8 not complete, got %#v", response)
	}
	if len(response.SupportedEditors) != 2 ||
		len(response.TrustSignalClasses) != 7 ||
		len(response.ValidationClasses) != 4 ||
		len(response.SurfaceRefs) != len(operability.DeveloperEcosystemValAProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.DeveloperEcosystemValAProofEvidenceRefs()) ||
		len(response.WhyPoint8NotPass) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and read-only summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "developer_ecosystem_vala") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

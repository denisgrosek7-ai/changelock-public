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

func TestDeveloperEcosystemVal0Handlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/developer-ecosystem/val0/status?tenant_id=acme&environment=prod",
		"/v1/developer-ecosystem/val0/proofs?tenant_id=acme&environment=prod",
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

	req := httptest.NewRequest(http.MethodPost, "/v1/developer-ecosystem/val0/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status route to remain read-only, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestDeveloperEcosystemVal0ProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/val0/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.DeveloperEcosystemVal0StateActive ||
		response.DependencyState != operability.DeveloperEcosystemVal0DependencyStateActive ||
		response.Point6State != operability.ReferenceArchitecturePoint6StatePass ||
		response.Point7State != operability.VerifierEcosystemPoint7StatePass ||
		response.Point7ClosureState != operability.VerifierEcosystemValEStatePass ||
		!response.Point7PassAllowed {
		t.Fatalf("expected healthy dependency closure for developer Val 0, got %#v", response)
	}
	if response.Point8State != operability.DeveloperEcosystemPoint8StateNotComplete ||
		response.OutputClassificationState != operability.DeveloperEcosystemVal0OutputClassificationStateActive ||
		response.IDEAdvisoryState != operability.DeveloperEcosystemVal0IDEAdvisoryStateActive ||
		response.LocalProductionState != operability.DeveloperEcosystemVal0LocalProductionStateActive ||
		response.RepoPolicyBoundaryState != operability.DeveloperEcosystemVal0RepoPolicyStateActive ||
		response.PluginSafetyState != operability.DeveloperEcosystemVal0PluginSafetyStateActive ||
		response.PerformanceBudgetState != operability.DeveloperEcosystemVal0PerformanceBudgetStateActive ||
		response.DXMetricsState != operability.DeveloperEcosystemVal0DXMetricsStateActive ||
		response.NoOverclaimState != operability.DeveloperEcosystemVal0NoOverclaimStateActive {
		t.Fatalf("expected active developer discipline states with point 8 not complete, got %#v", response)
	}
	if len(response.ClassifiedSurfaces) != 4 ||
		len(response.OutputClasses) != 7 ||
		len(response.DXMetricNames) != 9 ||
		len(response.SurfaceRefs) != len(operability.DeveloperEcosystemVal0ProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.DeveloperEcosystemVal0ProofEvidenceRefs()) ||
		len(response.WhyPoint8NotPass) == 0 ||
		len(response.Limitations) == 0 ||
		len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact proof/evidence refs and read-only summary fields, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestDeveloperEcosystemVal0StatusExposesCanonicalPluginBudgetRef(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/developer-ecosystem/val0/status?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response developerEcosystemVal0StatusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode status: %v", err)
	}
	if response.Model.PluginSafety.PerformanceBudgetRef != operability.DeveloperEcosystemVal0PerformanceBudgetDisciplineID {
		t.Fatalf("expected canonical performance budget ref %q, got %#v", operability.DeveloperEcosystemVal0PerformanceBudgetDisciplineID, response)
	}
	if response.Model.PluginSafety.PerformanceBudgetRef == "developer-performance-budget" {
		t.Fatalf("expected old dangling performance budget ref to stay absent, got %#v", response)
	}
}

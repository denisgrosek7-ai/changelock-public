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

func TestVerifierEcosystemValBHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/verifier-ecosystem/valb/compatibility-matrix?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valb/schema-proof-compatibility?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valb/mixed-version-diagnostics?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valb/diagnostic-precedence?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valb/fixture-descriptors?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valb/conformance-cases?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valb/conformance-suite?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valb/output-classes?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valb/proofs?tenant_id=acme&environment=prod",
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
}

func TestVerifierEcosystemValBProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/verifier-ecosystem/valb/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response verifierEcosystemValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.VerifierEcosystemValBStateActive ||
		response.ValBState != operability.VerifierEcosystemValBStateActive ||
		response.Point5State != operability.IntelligenceCalibrationPoint5StatePass ||
		response.Point5DependencyState != operability.IntelligenceCalibrationValEStateActive ||
		response.Point6State != operability.ReferenceArchitecturePoint6StatePass ||
		response.Point6ClosureState != operability.ReferenceArchitectureValEStateActive ||
		!response.Point6PassAllowed ||
		response.Val0CurrentState != operability.VerifierEcosystemVal0StateActive ||
		response.Val0State != operability.VerifierEcosystemVal0StateActive ||
		response.ValACurrentState != operability.VerifierEcosystemValAStateActive ||
		response.ValAState != operability.VerifierEcosystemValAStateActive {
		t.Fatalf("expected active Val B proofs with healthy Točka 6, Val 0, and Val A dependencies, got %#v", response)
	}
	if response.Point7State != operability.VerifierEcosystemPoint7StateNotComplete ||
		response.CompatibilityMatrixState != operability.VerifierEcosystemValBCompatibilityMatrixStateActive ||
		response.SchemaProofCompatibilityState != operability.VerifierEcosystemValBSchemaProofCompatibilityStateActive ||
		response.MixedVersionDiagnosticState != operability.VerifierEcosystemValBMixedVersionStateActive ||
		response.DiagnosticPrecedenceState != operability.VerifierEcosystemValBDiagnosticPrecedenceStateActive ||
		response.FixtureDescriptorState != operability.VerifierEcosystemValBFixtureDescriptorStateActive ||
		response.ConformanceCaseState != operability.VerifierEcosystemValBConformanceCaseStateActive ||
		response.ConformanceSuiteState != operability.VerifierEcosystemValBConformanceSuiteStateActive ||
		response.OutputClassState != operability.VerifierEcosystemValBOutputClassStateActive ||
		response.CompatibilityState != operability.ReferenceArchitectureCompatibilityCompatible ||
		response.DerivedDiagnosticClass != operability.VerifierEcosystemDiagnosticVerified ||
		response.DerivedOutputClass != operability.VerifierEcosystemValBOutputClassVerified {
		t.Fatalf("expected active Val B component states, got %#v", response)
	}
	if len(response.SurfaceRefs) != 11 || len(response.EvidenceRefs) != 13 || len(response.WhyPoint7NotPass) == 0 || len(response.Limitations) == 0 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact surface refs, evidence refs, limitations, and summary, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

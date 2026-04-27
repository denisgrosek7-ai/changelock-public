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

func TestVerifierEcosystemVal0Handlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/verifier-ecosystem/val0/contract?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/val0/proof-envelope?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/val0/verification-scope?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/val0/schema-compatibility?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/val0/trust-root-issuer?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/val0/diagnostics?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/val0/output-boundaries?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/val0/proofs?tenant_id=acme&environment=prod",
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

func TestVerifierEcosystemVal0ProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/verifier-ecosystem/val0/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response verifierEcosystemVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.VerifierEcosystemVal0StateActive ||
		response.Val0State != operability.VerifierEcosystemVal0StateActive ||
		response.Point5State != operability.IntelligenceCalibrationPoint5StatePass ||
		response.Point5DependencyState != operability.IntelligenceCalibrationValEStateActive ||
		response.Point6State != operability.ReferenceArchitecturePoint6StatePass ||
		response.Point6ClosureState != operability.ReferenceArchitectureValEStateActive ||
		!response.Point6PassAllowed {
		t.Fatalf("expected active verifier discipline proofs with healthy Točka 6 closure dependency, got %#v", response)
	}
	if response.Point7State != operability.VerifierEcosystemPoint7StateNotComplete ||
		response.VerifierContractState != operability.VerifierEcosystemVal0ContractStateActive ||
		response.ProofEnvelopeState != operability.VerifierEcosystemVal0EnvelopeStateActive ||
		response.VerificationScopeState != operability.VerifierEcosystemVal0ScopeStateActive ||
		response.SchemaCompatibilityState != operability.VerifierEcosystemVal0CompatibilityStateActive ||
		response.TrustRootIssuerState != operability.VerifierEcosystemVal0TrustStateActive ||
		response.DiagnosticsState != operability.VerifierEcosystemVal0DiagnosticsStateActive ||
		response.OutputBoundaryState != operability.VerifierEcosystemVal0OutputBoundaryStateActive {
		t.Fatalf("expected active verifier component states with point 7 not complete, got %#v", response)
	}
	if len(response.SupportedProfiles) != 5 || len(response.SupportedModes) != 5 || len(response.SupportedScopeClasses) != 5 ||
		len(response.SurfaceRefs) != 8 || len(response.EvidenceRefs) < 5 || len(response.WhyPoint7NotPass) == 0 || len(response.Limitations) == 0 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected supported enums, exact surface refs, evidence refs, limitations, and summary, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

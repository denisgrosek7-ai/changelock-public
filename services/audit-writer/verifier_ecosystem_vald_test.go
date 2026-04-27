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

func TestVerifierEcosystemValDHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/verifier-ecosystem/vald/correctness-gate?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vald/tooling-gate?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vald/schema-compatibility-gate?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vald/diagnostics-conformance-gate?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vald/trust-key-rotation-gate?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vald/negative-diagnostics-gate?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vald/redaction-gate?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vald/publisher-artifact-gate?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vald/no-overclaim-gate?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vald/proofs?tenant_id=acme&environment=prod",
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

func TestVerifierEcosystemValDProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/verifier-ecosystem/vald/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response verifierEcosystemValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.VerifierEcosystemValDStateActive ||
		response.ValDState != operability.VerifierEcosystemValDStateActive ||
		response.Point5State != operability.IntelligenceCalibrationPoint5StatePass ||
		response.Point5DependencyState != operability.IntelligenceCalibrationValEStateActive ||
		response.Point6State != operability.ReferenceArchitecturePoint6StatePass ||
		response.Point6ClosureState != operability.ReferenceArchitectureValEStateActive ||
		!response.Point6PassAllowed ||
		response.Val0CurrentState != operability.VerifierEcosystemVal0StateActive ||
		response.Val0State != operability.VerifierEcosystemVal0StateActive ||
		response.ValACurrentState != operability.VerifierEcosystemValAStateActive ||
		response.ValAState != operability.VerifierEcosystemValAStateActive ||
		response.ValBCurrentState != operability.VerifierEcosystemValBStateActive ||
		response.ValBState != operability.VerifierEcosystemValBStateActive ||
		response.ValCCurrentState != operability.VerifierEcosystemValCStateActive ||
		response.ValCState != operability.VerifierEcosystemValCStateActive {
		t.Fatalf("expected active Val D proofs with healthy Točka 6 and Val 0 through Val C dependencies, got %#v", response)
	}
	if response.Point7State != operability.VerifierEcosystemPoint7StateNotComplete ||
		response.CorrectnessGateState != operability.VerifierEcosystemValDCorrectnessGateStateActive ||
		response.ToolingGateState != operability.VerifierEcosystemValDToolingGateStateActive ||
		response.SchemaCompatibilityGateState != operability.VerifierEcosystemValDSchemaCompatibilityGateStateActive ||
		response.DiagnosticsConformanceGateState != operability.VerifierEcosystemValDDiagnosticsConformanceGateStateActive ||
		response.TrustKeyRotationGateState != operability.VerifierEcosystemValDTrustKeyRotationGateStateActive ||
		response.NegativeDiagnosticsGateState != operability.VerifierEcosystemValDNegativeDiagnosticsGateStateActive ||
		response.RedactionGateState != operability.VerifierEcosystemValDRedactionGateStateActive ||
		response.PublisherArtifactGateState != operability.VerifierEcosystemValDPublisherArtifactGateStateActive ||
		response.NoOverclaimGateState != operability.VerifierEcosystemValDNoOverclaimGateStateActive ||
		response.DerivedPublicOutputClass != operability.VerifierEcosystemValBOutputClassVerified ||
		response.DerivedPartnerOutputClass != operability.VerifierEcosystemValBOutputClassVerified {
		t.Fatalf("expected active Val D component states, got %#v", response)
	}
	if response.TrustDistributionMode != operability.VerifierEcosystemValCDistributionModePartnerScopedDir ||
		response.OfflineDistributionScope != operability.VerifierEcosystemScopePartnerSafe ||
		response.TrustDistributionMode == response.OfflineDistributionScope {
		t.Fatalf("expected trust distribution mode to remain a Val C mode and scope to remain separate, got %#v", response)
	}
	if len(response.SurfaceRefs) != 14 || len(response.EvidenceRefs) != 16 || len(response.WhyPoint7NotPass) == 0 || len(response.Limitations) == 0 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact surface refs, evidence refs, limitations, and summary, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

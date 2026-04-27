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

func TestVerifierEcosystemValCHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/verifier-ecosystem/valc/audience-surfaces?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valc/public-output?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valc/partner-output?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valc/auditor-flow?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valc/request-contract?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valc/publisher-profile?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valc/artifact-rules?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valc/trust-distribution?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/valc/proofs?tenant_id=acme&environment=prod",
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

func TestVerifierEcosystemValCProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/verifier-ecosystem/valc/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response verifierEcosystemValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.VerifierEcosystemValCStateActive ||
		response.ValCState != operability.VerifierEcosystemValCStateActive ||
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
		response.ValBState != operability.VerifierEcosystemValBStateActive {
		t.Fatalf("expected active Val C proofs with healthy Točka 6, Val 0, Val A, and Val B dependencies, got %#v", response)
	}
	if response.Point7State != operability.VerifierEcosystemPoint7StateNotComplete ||
		response.AudienceSurfaceState != operability.VerifierEcosystemValCAudienceSurfaceStateActive ||
		response.PublicOutputState != operability.VerifierEcosystemValCPublicOutputStateActive ||
		response.PartnerOutputState != operability.VerifierEcosystemValCPartnerOutputStateActive ||
		response.AuditorFlowState != operability.VerifierEcosystemValCAuditorFlowStateActive ||
		response.RequestContractState != operability.VerifierEcosystemValCRequestContractStateActive ||
		response.PublisherProfileState != operability.VerifierEcosystemValCPublisherProfileStateActive ||
		response.ArtifactRuleState != operability.VerifierEcosystemValCArtifactRuleStateActive ||
		response.TrustDistributionState != operability.VerifierEcosystemValCTrustDistributionStateActive ||
		response.PublicOutputClass != operability.VerifierEcosystemValBOutputClassVerified ||
		response.PartnerOutputClass != operability.VerifierEcosystemValBOutputClassVerified {
		t.Fatalf("expected active Val C component states, got %#v", response)
	}
	if len(response.SurfaceRefs) != 12 || len(response.EvidenceRefs) != 14 || len(response.WhyPoint7NotPass) == 0 || len(response.Limitations) == 0 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact surface refs, evidence refs, limitations, and summary, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

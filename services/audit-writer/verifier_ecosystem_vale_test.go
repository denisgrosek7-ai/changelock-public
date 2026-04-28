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

func TestVerifierEcosystemValEHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	paths := []string{
		"/v1/verifier-ecosystem/vale/closure?tenant_id=acme&environment=prod",
		"/v1/verifier-ecosystem/vale/proofs?tenant_id=acme&environment=prod",
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

func TestVerifierEcosystemValEProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/verifier-ecosystem/vale/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response verifierEcosystemValEProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.VerifierEcosystemValEStatePass ||
		response.ValEState != operability.VerifierEcosystemValEStatePass ||
		response.Point6State != operability.ReferenceArchitecturePoint6StatePass ||
		!response.Point6PassAllowed ||
		response.Val0CurrentState != operability.VerifierEcosystemVal0StateActive ||
		response.Val0State != operability.VerifierEcosystemVal0StateActive ||
		response.ValACurrentState != operability.VerifierEcosystemValAStateActive ||
		response.ValAState != operability.VerifierEcosystemValAStateActive ||
		response.ValBCurrentState != operability.VerifierEcosystemValBStateActive ||
		response.ValBState != operability.VerifierEcosystemValBStateActive ||
		response.ValCCurrentState != operability.VerifierEcosystemValCStateActive ||
		response.ValCState != operability.VerifierEcosystemValCStateActive ||
		response.ValDCurrentState != operability.VerifierEcosystemValDStateActive ||
		response.ValDState != operability.VerifierEcosystemValDStateActive ||
		response.ValDFinalGateState != operability.VerifierEcosystemValDStateActive {
		t.Fatalf("expected healthy Val E dependencies and pass state, got %#v", response)
	}
	if response.ClosurePrerequisiteState != operability.VerifierEcosystemValEPrerequisiteStateActive ||
		response.ClosureInvariantState != operability.VerifierEcosystemValEInvariantStateActive ||
		response.ProofSurfaceState != operability.VerifierEcosystemValEProofSurfaceStateActive ||
		response.EvidenceQualityState != operability.VerifierEcosystemValEEvidenceQualityStateActive ||
		response.NoOverclaimState != operability.VerifierEcosystemValENoOverclaimStateActive ||
		response.PassRuleState != operability.VerifierEcosystemValEPassRuleStateActive ||
		response.Point7State != operability.VerifierEcosystemPoint7StatePass ||
		!response.Point7PassAllowed {
		t.Fatalf("expected active Val E closure components and point_7_pass, got %#v", response)
	}
	if len(response.SurfaceRefs) != len(operability.VerifierEcosystemValEProofSurfaceRefs()) ||
		len(response.EvidenceRefs) != len(operability.VerifierEcosystemValEProofEvidenceRefs()) ||
		len(response.ClosureInvariants) == 0 ||
		len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected exact surfaces, evidence refs, invariants, and summary, got %#v", response)
	}
	if response.Point7PassReason != operability.VerifierEcosystemValEPoint7PassReasonAllowed {
		t.Fatalf("expected Val E-only pass reason, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection disclaimer, got %#v", response)
	}
}

func TestVerifierEcosystemValEClosureModelRecomputePreservesPartialState(t *testing.T) {
	model := buildVerifierEcosystemValEClosureModel()
	model.SourceCurrentStates.Val0CurrentState = operability.VerifierEcosystemVal0StatePartial
	model.SourceValStates.Val0State = operability.VerifierEcosystemVal0StatePartial
	model.Point7PassAllowed = false
	model.Point7PassReason = operability.VerifierEcosystemValEPoint7PassReasonBlocked

	recomputed := operability.ComputeVerifierEcosystemValEClosure(model)
	if recomputed.PassRuleState != operability.VerifierEcosystemValEPassRuleStatePartial ||
		recomputed.CurrentState != operability.VerifierEcosystemValEStatePartial {
		t.Fatalf("expected canonical blocked reason to preserve partial service recompute state, got %#v", recomputed)
	}
}

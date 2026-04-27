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

func TestReferenceArchitectureValEHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/reference-architecture/vale/closure?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValEClosureResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode closure: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValEStateActive ||
					response.Model.Point6State != operability.ReferenceArchitecturePoint6StatePass ||
					!response.Model.Point6PassAllowed {
					t.Fatalf("unexpected closure response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/vale/proofs?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValEProofsResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode proofs: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValEStateActive ||
					response.ValEState != operability.ReferenceArchitectureValEStateActive ||
					response.Point6State != operability.ReferenceArchitecturePoint6StatePass ||
					!response.Point6PassAllowed {
					t.Fatalf("unexpected proofs response %#v", response)
				}
			},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, tc.path, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200 for %s, got %d: %s", tc.path, rec.Code, rec.Body.String())
		}
		tc.decode(t, rec)
	}
}

func TestReferenceArchitectureValEProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/reference-architecture/vale/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response referenceArchitectureValEProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StatePass ||
		response.Point5DependencyState != operability.IntelligenceCalibrationValEStateActive ||
		response.Val0DependencyState != operability.ReferenceArchitectureVal0StateActive ||
		response.Val0State != operability.ReferenceArchitectureVal0StateActive ||
		response.ValADependencyState != operability.ReferenceArchitectureValAStateActive ||
		response.ValAState != operability.ReferenceArchitectureValAStateActive ||
		response.ValBDependencyState != operability.ReferenceArchitectureValBStateActive ||
		response.ValBState != operability.ReferenceArchitectureValBStateActive ||
		response.ValCDependencyState != operability.ReferenceArchitectureValCStateActive ||
		response.ValCState != operability.ReferenceArchitectureValCStateActive ||
		response.ValDDependencyState != operability.ReferenceArchitectureValDStateActive ||
		response.ValDState != operability.ReferenceArchitectureValDStateActive ||
		response.ValDFinalGateState != operability.ReferenceArchitectureValDFinalGateStateActive {
		t.Fatalf("expected active dependency states, got %#v", response)
	}
	if response.ClosurePrerequisiteState != operability.ReferenceArchitectureValEPrerequisiteStateActive ||
		response.ClosureInvariantState != operability.ReferenceArchitectureValEInvariantStateActive ||
		response.ProofSurfaceState != operability.ReferenceArchitectureValEProofSurfaceStateActive ||
		response.PassRuleState != operability.ReferenceArchitectureValEPassRuleStateActive {
		t.Fatalf("expected active closure component states, got %#v", response)
	}
	if len(response.ClosureInvariants) != 6 || len(response.SurfaceRefs) != 8 || len(response.EvidenceRefs) == 0 || len(response.Limitations) == 0 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected closure invariants, exact surfaces, evidence refs, limitations, and summary, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection-only disclaimer, got %#v", response)
	}
}

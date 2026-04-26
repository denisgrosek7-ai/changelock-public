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

func TestReferenceArchitectureValCFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/reference-architecture/valc/scenario-packs?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValCScenarioPackResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode scenario packs: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValCScenarioPackStateActive || len(response.Model.ScenarioPacks) != 6 || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected scenario pack response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/valc/scaling-scenarios?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValCCollectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode scaling scenarios: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValCScalingScenarioStateActive || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected scaling scenario response %#v", response)
				}
			},
		},
		{
			path: "/v1/reference-architecture/valc/control-plane-safety?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response referenceArchitectureValCCollectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode control-plane safety: %v", err)
				}
				if response.CurrentState != operability.ReferenceArchitectureValCControlPlaneStateActive || len(response.FamilyStates) != 6 {
					t.Fatalf("unexpected control-plane response %#v", response)
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

func TestReferenceArchitectureValCProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/reference-architecture/valc/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response referenceArchitectureValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.ReferenceArchitectureValCStateActive {
		t.Fatalf("expected active proofs state, got %#v", response)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StatePass ||
		response.Point5DependencyState != operability.IntelligenceCalibrationValEStateActive ||
		response.Val0DependencyState != operability.ReferenceArchitectureVal0StateActive ||
		response.Val0State != operability.ReferenceArchitectureVal0StateActive ||
		response.ValADependencyState != operability.ReferenceArchitectureValAStateActive ||
		response.ValAState != operability.ReferenceArchitectureValAStateActive ||
		response.ValBDependencyState != operability.ReferenceArchitectureValBStateActive ||
		response.ValBState != operability.ReferenceArchitectureValBStateActive {
		t.Fatalf("expected active dependency states, got %#v", response)
	}
	if response.ValCState != operability.ReferenceArchitectureValCStateActive ||
		response.ScenarioPackState != operability.ReferenceArchitectureValCScenarioPackStateActive ||
		response.FailureTaxonomyState != operability.ReferenceArchitectureValCFailureTaxonomyStateActive ||
		response.ScenarioDescriptorState != operability.ReferenceArchitectureValCScenarioDescriptorStateActive ||
		response.DegradedModeState != operability.ReferenceArchitectureValCDegradedModeStateActive ||
		response.RecoveryExpectationState != operability.ReferenceArchitectureValCRecoveryExpectationStateActive ||
		response.ScalingScenarioState != operability.ReferenceArchitectureValCScalingScenarioStateActive ||
		response.TrustPathState != operability.ReferenceArchitectureValCTrustPathStateActive ||
		response.AuditPathState != operability.ReferenceArchitectureValCAuditPathStateActive ||
		response.ControlPlaneSafetyState != operability.ReferenceArchitectureValCControlPlaneStateActive {
		t.Fatalf("expected active Val C component states, got %#v", response)
	}
	if response.Point6State != operability.ReferenceArchitecturePoint6StateNotComplete {
		t.Fatalf("expected point 6 to remain not complete, got %#v", response)
	}
	if len(response.SupportedFamilies) != 6 || len(response.FamilyStates) != 6 || len(response.WhyPoint6NotPass) == 0 || len(response.SurfaceRefs) < 13 || len(response.EvidenceRefs) < 12 || len(response.Limitations) == 0 {
		t.Fatalf("expected supported families, refs, and limitations in proofs, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection-only disclaimer, got %#v", response)
	}
}

func TestReferenceArchitectureValCProofsPreservePoint5DependencyStateDivergence(t *testing.T) {
	response := buildReferenceArchitectureValCProofs()
	if response.Point5State != operability.IntelligenceCalibrationPoint5StatePass {
		t.Fatalf("expected point_5_state to report actual point 5 pass state, got %#v", response)
	}
	if response.Point5DependencyState != operability.IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected point_5_dependency_state to report actual dependency health, got %#v", response)
	}
	if response.Point5DependencyState == response.Point5State {
		t.Fatalf("expected point_5_state and point_5_dependency_state to be allowed to diverge, got %#v", response)
	}
}

func TestReferenceArchitectureValCProofStateRequiresHealthyPoint5Dependency(t *testing.T) {
	supportedFamilies := operability.ReferenceArchitectureValBPackRegistry().SupportedFamilies
	surfaceRefs := referenceArchitectureValCAllSurfaceRefs()
	evidenceRefs := []string{"point5", "val0", "vala", "valb", "registry", "taxonomy", "descriptors", "degraded", "recovery", "scaling", "trust", "audit"}
	limitations := []string{"Val C keeps point 6 not complete."}
	disclaimer := referenceArchitectureValCProjectionDisclaimer()

	if got := referenceArchitectureValCProofCurrentState(
		operability.ReferenceArchitectureValCStateActive,
		operability.IntelligenceCalibrationValEStateActive,
		operability.ReferenceArchitecturePoint6StateNotComplete,
		supportedFamilies,
		surfaceRefs,
		evidenceRefs,
		limitations,
		disclaimer,
	); got != operability.ReferenceArchitectureValCStateActive {
		t.Fatalf("expected active proofs state with healthy point 5 dependency, got %q", got)
	}

	for _, dependencyState := range []string{
		operability.IntelligenceCalibrationValEStateSubstantial,
		operability.IntelligenceCalibrationValEStateIncomplete,
		operability.IntelligenceCalibrationValEReviewUnsupported,
		"",
	} {
		if got := referenceArchitectureValCProofCurrentState(
			operability.ReferenceArchitectureValCStateActive,
			dependencyState,
			operability.ReferenceArchitecturePoint6StateNotComplete,
			supportedFamilies,
			surfaceRefs,
			evidenceRefs,
			limitations,
			disclaimer,
		); got == operability.ReferenceArchitectureValCStateActive {
			t.Fatalf("expected non-active proofs state for point_5_dependency_state %q, got %q", dependencyState, got)
		}
	}
}

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

func TestIntelligenceCalibrationValEFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/intelligence/calibration/vale/dependency-closure?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValEDependencyClosureResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode dependency closure: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValEDependencyClosureStateActive || response.Model.DependencyStatus != operability.IntelligenceCalibrationValEDependencyPass {
					t.Fatalf("unexpected dependency closure response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vale/coherence-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValECoherenceReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode coherence review: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValECoherenceReviewStateActive || response.Model.CoherenceState != operability.IntelligenceCalibrationValEReviewPass {
					t.Fatalf("unexpected coherence review response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vale/pass-rule?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValEPassRuleResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode pass rule: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValEPassRuleStateActive || response.Model.Point5State != operability.IntelligenceCalibrationPoint5StatePass {
					t.Fatalf("unexpected pass rule response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vale/advisory-boundary?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValEBoundaryReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode advisory boundary: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValEBoundaryReviewStateActive || response.Model.BoundaryState != operability.IntelligenceCalibrationValEReviewPass {
					t.Fatalf("unexpected advisory boundary response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vale/reachability-vex-safety?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValEReachabilityVEXSafetyResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode reachability/vex safety: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValEReachabilityVEXSafetyStateActive || response.Model.ReachabilityVEXState != operability.IntelligenceCalibrationValEReviewPass {
					t.Fatalf("unexpected reachability/vex safety response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vale/behavioral-learning-safety?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValEBehavioralLearningSafetyResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode behavioral/learning safety: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValEBehavioralLearningSafetyStateActive || response.Model.BehavioralLearningState != operability.IntelligenceCalibrationValEReviewPass {
					t.Fatalf("unexpected behavioral/learning safety response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vale/feedback-federated-safety?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValEFeedbackFederatedSafetyResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode feedback/federated safety: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValEFeedbackFederatedSafetyStateActive || response.Model.FeedbackFederatedState != operability.IntelligenceCalibrationValEReviewPass {
					t.Fatalf("unexpected feedback/federated safety response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vale/simulation-quality-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValESimulationQualityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode simulation/quality review: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValESimulationQualityStateActive || response.Model.SimulationQualityState != operability.IntelligenceCalibrationValEReviewPass {
					t.Fatalf("unexpected simulation/quality review response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vale/regression-closure?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValERegressionClosureResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode regression closure: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValERegressionClosureStateActive || response.Model.RegressionState != operability.IntelligenceCalibrationValEReviewPass {
					t.Fatalf("unexpected regression closure response %#v", response)
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

func TestIntelligenceCalibrationValEProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/intelligence/calibration/vale/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response intelligenceCalibrationValEProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val E proofs: %v", err)
	}
	if response.CurrentState != operability.IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected active Val E proofs state, got %#v", response)
	}
	if response.ValEState != operability.IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected active Val E state in proofs, got %#v", response)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StatePass || !response.PassCriteriaMet {
		t.Fatalf("expected point 5 pass with criteria met in Val E proofs, got %#v", response)
	}
	if len(response.PassBlockers) != 0 {
		t.Fatalf("expected no pass blockers in active Val E proofs, got %#v", response)
	}
	if len(response.SurfaceRefs) == 0 || len(response.EvidenceRefs) == 0 || len(response.Limitations) == 0 {
		t.Fatalf("expected Val E proofs to expose surface refs, evidence refs, and limitations, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected Val E proofs disclaimer to stay projection-only, got %#v", response)
	}
}

func TestIntelligenceCalibrationValEPassRuleModelUsesPrerequisiteValEState(t *testing.T) {
	val0 := buildIntelligenceCalibrationVal0Proofs()
	valA := buildIntelligenceCalibrationValAProofs()
	valB := buildIntelligenceCalibrationValBProofs()
	valC := buildIntelligenceCalibrationValCProofs()
	valD := buildIntelligenceCalibrationValDProofs()

	dependencyClosure := buildIntelligenceCalibrationValEDependencyClosureModel(val0, valA, valB, valC, valD)
	coherenceReview := buildIntelligenceCalibrationValECoherenceReviewModel(val0, valA, valB, valC, valD)
	boundaryReview := buildIntelligenceCalibrationValEBoundaryReviewModel(val0, valA, valB, valC, valD)
	reachabilityVEXSafety := buildIntelligenceCalibrationValEReachabilityVEXSafetyReviewModel()
	behavioralLearningSafety := buildIntelligenceCalibrationValEBehavioralLearningSafetyReviewModel()
	feedbackFederatedSafety := buildIntelligenceCalibrationValEFeedbackFederatedSafetyReviewModel()
	simulationQuality := buildIntelligenceCalibrationValESimulationQualityReviewModel()
	regressionClosure := buildIntelligenceCalibrationValERegressionClosureModel()

	passRule := buildIntelligenceCalibrationValEPassRuleModel(
		val0,
		valA,
		valB,
		valC,
		valD,
		operability.IntelligenceCalibrationValEStateActive,
		dependencyClosure,
		coherenceReview,
		boundaryReview,
		reachabilityVEXSafety,
		behavioralLearningSafety,
		feedbackFederatedSafety,
		simulationQuality,
		regressionClosure,
	)
	if passRule.ValEState != operability.IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected pass-rule model to retain prerequisite active Val E state, got %#v", passRule)
	}
	if got := operability.EvaluateIntelligenceCalibrationValEPassRuleState(passRule); got != operability.IntelligenceCalibrationValEPassRuleStateActive {
		t.Fatalf("expected active pass-rule state with active prerequisite Val E state and no blockers, got %q", got)
	}

	boundaryReview.ViolationSurfaces = []string{"/v1/intelligence/calibration/vale/proofs"}
	boundaryReview.BoundaryState = operability.IntelligenceCalibrationValEReviewBlocked
	passRule = buildIntelligenceCalibrationValEPassRuleModel(
		val0,
		valA,
		valB,
		valC,
		valD,
		operability.IntelligenceCalibrationValEStateActive,
		dependencyClosure,
		coherenceReview,
		boundaryReview,
		reachabilityVEXSafety,
		behavioralLearningSafety,
		feedbackFederatedSafety,
		simulationQuality,
		regressionClosure,
	)
	if passRule.ValEState != operability.IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected blocked pass-rule model to still retain prerequisite active Val E state, got %#v", passRule)
	}
	if got := operability.EvaluateIntelligenceCalibrationValEPassRuleState(passRule); got == operability.IntelligenceCalibrationValEPassRuleStateActive {
		t.Fatalf("expected non-active pass-rule state when blockers exist despite active prerequisite Val E state, got %q", got)
	}
}

func TestIntelligenceCalibrationValEPassRuleModelRemainsPartialWithSubstantialPrerequisiteState(t *testing.T) {
	val0 := buildIntelligenceCalibrationVal0Proofs()
	valA := buildIntelligenceCalibrationValAProofs()
	valB := buildIntelligenceCalibrationValBProofs()
	valC := buildIntelligenceCalibrationValCProofs()
	valD := buildIntelligenceCalibrationValDProofs()

	passRule := buildIntelligenceCalibrationValEPassRuleModel(
		val0,
		valA,
		valB,
		valC,
		valD,
		operability.IntelligenceCalibrationValEStateSubstantial,
		buildIntelligenceCalibrationValEDependencyClosureModel(val0, valA, valB, valC, valD),
		buildIntelligenceCalibrationValECoherenceReviewModel(val0, valA, valB, valC, valD),
		buildIntelligenceCalibrationValEBoundaryReviewModel(val0, valA, valB, valC, valD),
		buildIntelligenceCalibrationValEReachabilityVEXSafetyReviewModel(),
		buildIntelligenceCalibrationValEBehavioralLearningSafetyReviewModel(),
		buildIntelligenceCalibrationValEFeedbackFederatedSafetyReviewModel(),
		buildIntelligenceCalibrationValESimulationQualityReviewModel(),
		buildIntelligenceCalibrationValERegressionClosureModel(),
	)
	if passRule.ValEState != operability.IntelligenceCalibrationValEStateSubstantial {
		t.Fatalf("expected pass-rule model to retain substantial prerequisite Val E state, got %#v", passRule)
	}
	if got := operability.EvaluateIntelligenceCalibrationValEPassRuleState(passRule); got == operability.IntelligenceCalibrationValEPassRuleStateActive {
		t.Fatalf("expected non-active pass-rule state with substantial prerequisite Val E state, got %q", got)
	}
}

func TestIntelligenceCalibrationValEReachabilityVEXSafetyAllowsSufficientNotAffectedCandidate(t *testing.T) {
	aggregation := operability.IntelligenceCalibrationValAReachabilityAggregationContract()
	exploitability := operability.IntelligenceCalibrationValAExploitabilityCalibrationContract()
	decision := operability.IntelligenceCalibrationValADowngradeEscalationContract()
	vexCandidate := operability.IntelligenceCalibrationValAVEXCandidateContract()
	vexCandidate.SuggestedVEXOutcome = operability.IntelligenceCalibrationVEXOutcomeNotAffectedCandidate
	vexCandidate.EvidenceSufficiencyState = operability.IntelligenceCalibrationValAVEXSufficiencySufficient
	vexSufficiency := operability.IntelligenceCalibrationValAVEXSufficiencyContract()
	vexSufficiency.SufficiencyState = operability.IntelligenceCalibrationValAVEXSufficiencySufficient
	explanation := operability.IntelligenceCalibrationValAExplanationContract()
	publicationGuardrail := operability.IntelligenceCalibrationValAPublicationGuardrailContract()
	dReachability := operability.IntelligenceCalibrationValDReachabilityQualityReviewContract()
	dVEX := operability.IntelligenceCalibrationValDVEXQualityReviewContract()

	model := buildIntelligenceCalibrationValEReachabilityVEXSafetyReviewModelFromContracts(
		aggregation,
		exploitability,
		decision,
		vexCandidate,
		vexSufficiency,
		explanation,
		publicationGuardrail,
		dReachability,
		dVEX,
	)
	if !model.InsufficientEvidenceBlocksNotAffected {
		t.Fatalf("expected sufficient-evidence not_affected_candidate to remain allowed as a candidate without integrated insufficiency blocker, got %#v", model)
	}
	if got := operability.EvaluateIntelligenceCalibrationValEReachabilityVEXSafetyState(model); got != operability.IntelligenceCalibrationValEReachabilityVEXSafetyStateActive {
		t.Fatalf("expected active Val E reachability/VEX safety review with sufficient-evidence not_affected_candidate, got %q (%#v)", got, model)
	}
}

func TestIntelligenceCalibrationValEReachabilityVEXSafetyBlocksNotAffectedCandidateWithoutSufficientEvidence(t *testing.T) {
	buildModel := func(candidateState string, sufficiencyState string, staleRefs []string, unsupportedRefs []string, missingClasses []string) operability.IntelligenceCalibrationIntegratedReachabilityVEXSafetyReview {
		vexCandidate := operability.IntelligenceCalibrationValAVEXCandidateContract()
		vexCandidate.SuggestedVEXOutcome = operability.IntelligenceCalibrationVEXOutcomeNotAffectedCandidate
		vexCandidate.EvidenceSufficiencyState = candidateState
		vexSufficiency := operability.IntelligenceCalibrationValAVEXSufficiencyContract()
		vexSufficiency.SufficiencyState = sufficiencyState
		vexSufficiency.StaleEvidenceRefs = staleRefs
		vexSufficiency.UnsupportedEvidenceRefs = unsupportedRefs
		vexSufficiency.MissingEvidenceClasses = missingClasses
		if len(missingClasses) > 0 {
			vexSufficiency.PresentEvidenceClasses = []string{operability.IntelligenceCalibrationEvidenceDirectlyEvidenced}
		}
		return buildIntelligenceCalibrationValEReachabilityVEXSafetyReviewModelFromContracts(
			operability.IntelligenceCalibrationValAReachabilityAggregationContract(),
			operability.IntelligenceCalibrationValAExploitabilityCalibrationContract(),
			operability.IntelligenceCalibrationValADowngradeEscalationContract(),
			vexCandidate,
			vexSufficiency,
			operability.IntelligenceCalibrationValAExplanationContract(),
			operability.IntelligenceCalibrationValAPublicationGuardrailContract(),
			operability.IntelligenceCalibrationValDReachabilityQualityReviewContract(),
			operability.IntelligenceCalibrationValDVEXQualityReviewContract(),
		)
	}

	for _, tc := range []struct {
		name            string
		candidateState  string
		sufficiency     string
		staleRefs       []string
		unsupportedRefs []string
		missingClasses  []string
	}{
		{
			name:           "insufficient",
			candidateState: operability.IntelligenceCalibrationValAVEXSufficiencyInsufficient,
			sufficiency:    operability.IntelligenceCalibrationValAVEXSufficiencyInsufficient,
			missingClasses: []string{operability.IntelligenceCalibrationEvidenceStronglyInferred},
		},
		{
			name:           "stale",
			candidateState: operability.IntelligenceCalibrationValAVEXSufficiencyStale,
			sufficiency:    operability.IntelligenceCalibrationValAVEXSufficiencyStale,
			staleRefs:      []string{"evidence:stale"},
		},
		{
			name:            "unsupported",
			candidateState:  operability.IntelligenceCalibrationValAVEXSufficiencyUnsupported,
			sufficiency:     operability.IntelligenceCalibrationValAVEXSufficiencyUnsupported,
			unsupportedRefs: []string{"evidence:unsupported"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			model := buildModel(tc.candidateState, tc.sufficiency, tc.staleRefs, tc.unsupportedRefs, tc.missingClasses)
			if model.InsufficientEvidenceBlocksNotAffected {
				t.Fatalf("expected integrated Val E review to mark not_affected_candidate as blocked when sufficiency is %s, got %#v", tc.sufficiency, model)
			}
			if got := operability.EvaluateIntelligenceCalibrationValEReachabilityVEXSafetyState(model); got == operability.IntelligenceCalibrationValEReachabilityVEXSafetyStateActive {
				t.Fatalf("expected non-active Val E reachability/VEX safety review for %s-evidence not_affected_candidate, got %q (%#v)", tc.name, got, model)
			}
		})
	}
}

func TestIntelligenceCalibrationValEReachabilityVEXSafetyStillBlocksFinalVEXClaim(t *testing.T) {
	vexCandidate := operability.IntelligenceCalibrationValAVEXCandidateContract()
	vexCandidate.FinalVEXClaim = true
	model := buildIntelligenceCalibrationValEReachabilityVEXSafetyReviewModelFromContracts(
		operability.IntelligenceCalibrationValAReachabilityAggregationContract(),
		operability.IntelligenceCalibrationValAExploitabilityCalibrationContract(),
		operability.IntelligenceCalibrationValADowngradeEscalationContract(),
		vexCandidate,
		operability.IntelligenceCalibrationValAVEXSufficiencyContract(),
		operability.IntelligenceCalibrationValAExplanationContract(),
		operability.IntelligenceCalibrationValAPublicationGuardrailContract(),
		operability.IntelligenceCalibrationValDReachabilityQualityReviewContract(),
		operability.IntelligenceCalibrationValDVEXQualityReviewContract(),
	)
	if got := operability.EvaluateIntelligenceCalibrationValEReachabilityVEXSafetyState(model); got == operability.IntelligenceCalibrationValEReachabilityVEXSafetyStateActive {
		t.Fatalf("expected final_vex_claim to remain blocked in Val E reachability/VEX safety review, got %q (%#v)", got, model)
	}
}

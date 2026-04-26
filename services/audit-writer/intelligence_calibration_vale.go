package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	intelligenceCalibrationValEDependencyClosureSchema  = "point5.intelligence_calibration.vale.dependency_closure.v1"
	intelligenceCalibrationValECoherenceReviewSchema    = "point5.intelligence_calibration.vale.coherence_review.v1"
	intelligenceCalibrationValEPassRuleSchema           = "point5.intelligence_calibration.vale.pass_rule.v1"
	intelligenceCalibrationValEBoundaryReviewSchema     = "point5.intelligence_calibration.vale.advisory_boundary.v1"
	intelligenceCalibrationValEReachabilityVEXSchema    = "point5.intelligence_calibration.vale.reachability_vex_safety.v1"
	intelligenceCalibrationValEBehavioralLearningSchema = "point5.intelligence_calibration.vale.behavioral_learning_safety.v1"
	intelligenceCalibrationValEFeedbackFederatedSchema  = "point5.intelligence_calibration.vale.feedback_federated_safety.v1"
	intelligenceCalibrationValESimulationQualitySchema  = "point5.intelligence_calibration.vale.simulation_quality_review.v1"
	intelligenceCalibrationValERegressionClosureSchema  = "point5.intelligence_calibration.vale.regression_closure.v1"
	intelligenceCalibrationValEProofsSchema             = "point5.intelligence_calibration.vale.proofs.v1"
)

type intelligenceCalibrationValEDependencyClosureResponse struct {
	SchemaVersion string                                                         `json:"schema_version"`
	GeneratedAt   time.Time                                                      `json:"generated_at"`
	CurrentState  string                                                         `json:"current_state"`
	Model         operability.IntelligenceCalibrationIntegratedDependencyClosure `json:"model"`
	RouteRefs     []string                                                       `json:"route_refs,omitempty"`
	Limitations   []string                                                       `json:"limitations,omitempty"`
}

type intelligenceCalibrationValECoherenceReviewResponse struct {
	SchemaVersion string                                                     `json:"schema_version"`
	GeneratedAt   time.Time                                                  `json:"generated_at"`
	CurrentState  string                                                     `json:"current_state"`
	Model         operability.IntelligenceCalibrationCrossValCoherenceReview `json:"model"`
	RouteRefs     []string                                                   `json:"route_refs,omitempty"`
	Limitations   []string                                                   `json:"limitations,omitempty"`
}

type intelligenceCalibrationValEPassRuleResponse struct {
	SchemaVersion string                                                      `json:"schema_version"`
	GeneratedAt   time.Time                                                   `json:"generated_at"`
	CurrentState  string                                                      `json:"current_state"`
	Model         operability.IntelligenceCalibrationPoint5IntegratedPassRule `json:"model"`
	RouteRefs     []string                                                    `json:"route_refs,omitempty"`
	Limitations   []string                                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationValEBoundaryReviewResponse struct {
	SchemaVersion string                                                              `json:"schema_version"`
	GeneratedAt   time.Time                                                           `json:"generated_at"`
	CurrentState  string                                                              `json:"current_state"`
	Model         operability.IntelligenceCalibrationIntegratedAdvisoryBoundaryReview `json:"model"`
	RouteRefs     []string                                                            `json:"route_refs,omitempty"`
	Limitations   []string                                                            `json:"limitations,omitempty"`
}

type intelligenceCalibrationValEReachabilityVEXSafetyResponse struct {
	SchemaVersion string                                                                   `json:"schema_version"`
	GeneratedAt   time.Time                                                                `json:"generated_at"`
	CurrentState  string                                                                   `json:"current_state"`
	Model         operability.IntelligenceCalibrationIntegratedReachabilityVEXSafetyReview `json:"model"`
	RouteRefs     []string                                                                 `json:"route_refs,omitempty"`
	Limitations   []string                                                                 `json:"limitations,omitempty"`
}

type intelligenceCalibrationValEBehavioralLearningSafetyResponse struct {
	SchemaVersion string                                                                      `json:"schema_version"`
	GeneratedAt   time.Time                                                                   `json:"generated_at"`
	CurrentState  string                                                                      `json:"current_state"`
	Model         operability.IntelligenceCalibrationIntegratedBehavioralLearningSafetyReview `json:"model"`
	RouteRefs     []string                                                                    `json:"route_refs,omitempty"`
	Limitations   []string                                                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationValEFeedbackFederatedSafetyResponse struct {
	SchemaVersion string                                                                     `json:"schema_version"`
	GeneratedAt   time.Time                                                                  `json:"generated_at"`
	CurrentState  string                                                                     `json:"current_state"`
	Model         operability.IntelligenceCalibrationIntegratedFeedbackFederatedSafetyReview `json:"model"`
	RouteRefs     []string                                                                   `json:"route_refs,omitempty"`
	Limitations   []string                                                                   `json:"limitations,omitempty"`
}

type intelligenceCalibrationValESimulationQualityResponse struct {
	SchemaVersion string                                                               `json:"schema_version"`
	GeneratedAt   time.Time                                                            `json:"generated_at"`
	CurrentState  string                                                               `json:"current_state"`
	Model         operability.IntelligenceCalibrationIntegratedSimulationQualityReview `json:"model"`
	RouteRefs     []string                                                             `json:"route_refs,omitempty"`
	Limitations   []string                                                             `json:"limitations,omitempty"`
}

type intelligenceCalibrationValERegressionClosureResponse struct {
	SchemaVersion string                                                         `json:"schema_version"`
	GeneratedAt   time.Time                                                      `json:"generated_at"`
	CurrentState  string                                                         `json:"current_state"`
	Model         operability.IntelligenceCalibrationIntegratedRegressionClosure `json:"model"`
	RouteRefs     []string                                                       `json:"route_refs,omitempty"`
	Limitations   []string                                                       `json:"limitations,omitempty"`
}

type intelligenceCalibrationValEProofsResponse struct {
	SchemaVersion                  string    `json:"schema_version"`
	GeneratedAt                    time.Time `json:"generated_at"`
	CurrentState                   string    `json:"current_state"`
	Val0DependencyState            string    `json:"val_0_dependency_state"`
	Val0FoundationState            string    `json:"val_0_foundation_state"`
	ValADependencyState            string    `json:"val_a_dependency_state"`
	ValAReachabilityVEXState       string    `json:"val_a_reachability_vex_state"`
	ValBDependencyState            string    `json:"val_b_dependency_state"`
	ValBBehavioralLearningState    string    `json:"val_b_behavioral_learning_state"`
	ValCDependencyState            string    `json:"val_c_dependency_state"`
	ValCFeedbackFederatedState     string    `json:"val_c_feedback_suppression_federated_state"`
	ValDDependencyState            string    `json:"val_d_dependency_state"`
	ValDSimulationQualityGateState string    `json:"val_d_simulation_quality_gate_state"`
	ValEState                      string    `json:"val_e_state"`
	DependencyClosureState         string    `json:"dependency_closure_state"`
	CoherenceReviewState           string    `json:"coherence_review_state"`
	Point5State                    string    `json:"point_5_state"`
	PassCriteriaMet                bool      `json:"pass_criteria_met"`
	PassBlockers                   []string  `json:"pass_blockers,omitempty"`
	PassWarnings                   []string  `json:"pass_warnings,omitempty"`
	PassLimitations                []string  `json:"pass_limitations,omitempty"`
	AdvisoryBoundaryState          string    `json:"advisory_canonical_boundary_state"`
	ReachabilityVEXSafetyState     string    `json:"reachability_vex_safety_state"`
	BehavioralLearningSafetyState  string    `json:"behavioral_learning_safety_state"`
	FeedbackFederatedSafetyState   string    `json:"feedback_federated_safety_state"`
	SimulationQualityReviewState   string    `json:"simulation_quality_measurement_state"`
	RegressionClosureState         string    `json:"regression_closure_state"`
	EvidenceRefs                   []string  `json:"evidence_refs,omitempty"`
	SurfaceRefs                    []string  `json:"surface_refs,omitempty"`
	Limitations                    []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer           string    `json:"projection_disclaimer"`
	IntegrationSummary             []string  `json:"integration_summary,omitempty"`
}

type intelligenceCalibrationValEModels struct {
	val0 intelligenceCalibrationVal0ProofsResponse
	valA intelligenceCalibrationValAProofsResponse
	valB intelligenceCalibrationValBProofsResponse
	valC intelligenceCalibrationValCProofsResponse
	valD intelligenceCalibrationValDProofsResponse

	dependencyClosure        operability.IntelligenceCalibrationIntegratedDependencyClosure
	coherenceReview          operability.IntelligenceCalibrationCrossValCoherenceReview
	passRule                 operability.IntelligenceCalibrationPoint5IntegratedPassRule
	boundaryReview           operability.IntelligenceCalibrationIntegratedAdvisoryBoundaryReview
	reachabilityVEXSafety    operability.IntelligenceCalibrationIntegratedReachabilityVEXSafetyReview
	behavioralLearningSafety operability.IntelligenceCalibrationIntegratedBehavioralLearningSafetyReview
	feedbackFederatedSafety  operability.IntelligenceCalibrationIntegratedFeedbackFederatedSafetyReview
	simulationQualityReview  operability.IntelligenceCalibrationIntegratedSimulationQualityReview
	regressionClosure        operability.IntelligenceCalibrationIntegratedRegressionClosure

	dependencyClosureState        string
	coherenceReviewState          string
	passRuleState                 string
	boundaryReviewState           string
	reachabilityVEXSafetyState    string
	behavioralLearningSafetyState string
	feedbackFederatedSafetyState  string
	simulationQualityState        string
	regressionClosureState        string
	valEState                     string
	point5State                   string

	surfaceRefs  []string
	evidenceRefs []string
	limitations  []string
}

func intelligenceCalibrationValERequiredVals() []string {
	return []string{"val_0", "val_a", "val_b", "val_c", "val_d", "val_e"}
}

func intelligenceCalibrationValEAllSurfaceRefs() []string {
	return []string{
		"/v1/intelligence/calibration/val0/proofs",
		"/v1/intelligence/calibration/vala/proofs",
		"/v1/intelligence/calibration/valb/proofs",
		"/v1/intelligence/calibration/valc/proofs",
		"/v1/intelligence/calibration/vald/proofs",
		"/v1/intelligence/calibration/vale/dependency-closure",
		"/v1/intelligence/calibration/vale/coherence-review",
		"/v1/intelligence/calibration/vale/pass-rule",
		"/v1/intelligence/calibration/vale/advisory-boundary",
		"/v1/intelligence/calibration/vale/reachability-vex-safety",
		"/v1/intelligence/calibration/vale/behavioral-learning-safety",
		"/v1/intelligence/calibration/vale/feedback-federated-safety",
		"/v1/intelligence/calibration/vale/simulation-quality-review",
		"/v1/intelligence/calibration/vale/regression-closure",
		"/v1/intelligence/calibration/vale/proofs",
	}
}

func intelligenceCalibrationValEEvidenceRefs() []string {
	return []string{
		"val0_proofs",
		"vala_proofs",
		"valb_proofs",
		"valc_proofs",
		"vald_proofs",
		"integrated_dependency_closure",
		"cross_val_coherence_review",
		"integrated_point5_pass_rule",
		"integrated_advisory_boundary_review",
		"integrated_reachability_vex_safety_review",
		"integrated_behavioral_learning_safety_review",
		"integrated_feedback_federated_safety_review",
		"integrated_simulation_quality_review",
		"integrated_regression_closure",
		"evidence_spine",
	}
}

func intelligenceCalibrationValEProjectionDisclaimer() string {
	return "projection_only not_canonical_truth integrated_intelligence_calibration_closure"
}

func intelligenceCalibrationValEHasProjectionDisclaimer(disclaimer string) bool {
	trimmed := strings.TrimSpace(disclaimer)
	return strings.Contains(trimmed, "projection_only") && strings.Contains(trimmed, "not_canonical_truth")
}

func intelligenceCalibrationValEContainsTrimmedString(values []string, expected string) bool {
	needle := strings.TrimSpace(expected)
	if needle == "" {
		return false
	}
	for _, value := range values {
		if strings.TrimSpace(value) == needle {
			return true
		}
	}
	return false
}

func intelligenceCalibrationValEContainsAllTrimmedStrings(values []string, expected ...string) bool {
	for _, item := range expected {
		if !intelligenceCalibrationValEContainsTrimmedString(values, item) {
			return false
		}
	}
	return true
}

func intelligenceCalibrationValEContainsSubstringInTrimmedStrings(values []string, expected string) bool {
	needle := strings.TrimSpace(expected)
	if needle == "" {
		return false
	}
	for _, value := range values {
		if strings.Contains(strings.TrimSpace(value), needle) {
			return true
		}
	}
	return false
}

func intelligenceCalibrationValEParseTimestamp(value string) (time.Time, bool) {
	parsed, err := time.Parse(time.RFC3339, strings.TrimSpace(value))
	if err != nil {
		return time.Time{}, false
	}
	return parsed, true
}

func collectIntelligenceCalibrationValELimitations(val0 intelligenceCalibrationVal0ProofsResponse, valA intelligenceCalibrationValAProofsResponse, valB intelligenceCalibrationValBProofsResponse, valC intelligenceCalibrationValCProofsResponse, valD intelligenceCalibrationValDProofsResponse) []string {
	limitations := []string{
		"Integrated closure remains a projection-only summary and does not replace canonical evidence truth or governance.",
		"Točka 5 can become pass only through the active Val E integrated closure; any missing, partial, inconsistent, or unsupported prior val keeps point_5_state not_complete.",
	}
	limitations = appendPrefixedLimitations(limitations, "val0", val0.Limitations)
	limitations = appendPrefixedLimitations(limitations, "vala", valA.Limitations)
	limitations = appendPrefixedLimitations(limitations, "valb", valB.Limitations)
	limitations = appendPrefixedLimitations(limitations, "valc", valC.Limitations)
	limitations = appendPrefixedLimitations(limitations, "vald", valD.Limitations)
	return limitations
}

func intelligenceCalibrationValEActiveValNames(val0State, valAState, valBState, valCState, valDState, valEState string) []string {
	active := []string{}
	if val0State == operability.IntelligenceCalibrationVal0StateActive {
		active = append(active, "val_0")
	}
	if valAState == operability.IntelligenceCalibrationValAStateActive {
		active = append(active, "val_a")
	}
	if valBState == operability.IntelligenceCalibrationValBStateActive {
		active = append(active, "val_b")
	}
	if valCState == operability.IntelligenceCalibrationValCStateActive {
		active = append(active, "val_c")
	}
	if valDState == operability.IntelligenceCalibrationValDStateActive {
		active = append(active, "val_d")
	}
	if valEState == operability.IntelligenceCalibrationValEStateActive {
		active = append(active, "val_e")
	}
	return active
}

func intelligenceCalibrationValEClassifyRequiredVals(val0State, valAState, valBState, valCState, valDState, valEState string) (missingVals, partialVals, unsupportedVals []string) {
	for _, item := range []struct {
		name  string
		state string
		want  string
	}{
		{name: "val_0", state: val0State, want: operability.IntelligenceCalibrationVal0StateActive},
		{name: "val_a", state: valAState, want: operability.IntelligenceCalibrationValAStateActive},
		{name: "val_b", state: valBState, want: operability.IntelligenceCalibrationValBStateActive},
		{name: "val_c", state: valCState, want: operability.IntelligenceCalibrationValCStateActive},
		{name: "val_d", state: valDState, want: operability.IntelligenceCalibrationValDStateActive},
		{name: "val_e", state: valEState, want: operability.IntelligenceCalibrationValEStateActive},
	} {
		switch {
		case item.state == "":
			missingVals = append(missingVals, item.name)
		case item.state == item.want:
		case item.state == operability.IntelligenceCalibrationValEDependencyUnsupported:
			unsupportedVals = append(unsupportedVals, item.name)
		default:
			partialVals = append(partialVals, item.name)
		}
	}
	return missingVals, partialVals, unsupportedVals
}

func intelligenceCalibrationValDAllQualityStatesActive(valD intelligenceCalibrationValDProofsResponse) bool {
	return valD.SimulationHarnessState == operability.IntelligenceCalibrationValDSimulationHarnessStateActive &&
		valD.ScenarioLibraryState == operability.IntelligenceCalibrationValDScenarioLibraryStateActive &&
		valD.MissedDetectionAnalysisState == operability.IntelligenceCalibrationValDMissedDetectionStateActive &&
		valD.FPFNBalanceReviewState == operability.IntelligenceCalibrationValDFPFNBalanceStateActive &&
		valD.ConfidenceCalibrationReviewState == operability.IntelligenceCalibrationValDConfidenceReviewStateActive &&
		valD.VEXQualityReviewState == operability.IntelligenceCalibrationValDVEXQualityStateActive &&
		valD.ReachabilityQualityReviewState == operability.IntelligenceCalibrationValDReachabilityQualityStateActive &&
		valD.BehavioralQualityReviewState == operability.IntelligenceCalibrationValDBehavioralQualityStateActive &&
		valD.FederatedQualityReviewState == operability.IntelligenceCalibrationValDFederatedQualityStateActive &&
		valD.SimulationCoverageReviewState == operability.IntelligenceCalibrationValDSimulationCoverageStateActive &&
		valD.QualityScoreboardState == operability.IntelligenceCalibrationValDQualityScoreboardStateActive
}

func buildIntelligenceCalibrationValEDependencyClosureModel(val0 intelligenceCalibrationVal0ProofsResponse, valA intelligenceCalibrationValAProofsResponse, valB intelligenceCalibrationValBProofsResponse, valC intelligenceCalibrationValCProofsResponse, valD intelligenceCalibrationValDProofsResponse) operability.IntelligenceCalibrationIntegratedDependencyClosure {
	model := operability.IntelligenceCalibrationValEDependencyClosureContract()
	model.Val0State = val0.Val0State
	model.ValAState = valA.ValAState
	model.ValBState = valB.ValBState
	model.ValCState = valC.ValCState
	model.ValDState = valD.ValDState
	model.DependencyEvidenceRefs = []string{"val0_proofs", "vala_proofs", "valb_proofs", "valc_proofs", "vald_proofs"}
	model.DependencySurfaceRefs = []string{
		"/v1/intelligence/calibration/val0/proofs",
		"/v1/intelligence/calibration/vala/proofs",
		"/v1/intelligence/calibration/valb/proofs",
		"/v1/intelligence/calibration/valc/proofs",
		"/v1/intelligence/calibration/vald/proofs",
		"/v1/intelligence/calibration/vale/dependency-closure",
		"/v1/intelligence/calibration/vale/proofs",
	}
	model.ProofStatesObserved = val0.CurrentState != "" && valA.CurrentState != "" && valB.CurrentState != "" && valC.CurrentState != "" && valD.CurrentState != ""

	if val0.CurrentState == "" || val0.Val0State == "" {
		model.MissingVals = append(model.MissingVals, "val_0")
	}
	if valA.CurrentState == "" || valA.ValAState == "" {
		model.MissingVals = append(model.MissingVals, "val_a")
	}
	if valB.CurrentState == "" || valB.ValBState == "" {
		model.MissingVals = append(model.MissingVals, "val_b")
	}
	if valC.CurrentState == "" || valC.ValCState == "" {
		model.MissingVals = append(model.MissingVals, "val_c")
	}
	if valD.CurrentState == "" || valD.ValDState == "" {
		model.MissingVals = append(model.MissingVals, "val_d")
	}

	if val0.Val0State != operability.IntelligenceCalibrationVal0StateActive && val0.Val0State != "" {
		model.InactiveVals = append(model.InactiveVals, "val_0")
	}
	if valA.ValAState != operability.IntelligenceCalibrationValAStateActive && valA.ValAState != "" {
		model.InactiveVals = append(model.InactiveVals, "val_a")
	}
	if valB.ValBState != operability.IntelligenceCalibrationValBStateActive && valB.ValBState != "" {
		model.InactiveVals = append(model.InactiveVals, "val_b")
	}
	if valC.ValCState != operability.IntelligenceCalibrationValCStateActive && valC.ValCState != "" {
		model.InactiveVals = append(model.InactiveVals, "val_c")
	}
	if valD.ValDState != operability.IntelligenceCalibrationValDStateActive && valD.ValDState != "" {
		model.InactiveVals = append(model.InactiveVals, "val_d")
	}

	if val0.CurrentState != "" && val0.CurrentState != operability.IntelligenceCalibrationVal0StateActive {
		model.InconsistentVals = append(model.InconsistentVals, "val_0.proofs_state")
	}
	if valA.CurrentState != "" && valA.CurrentState != operability.IntelligenceCalibrationValAStateActive {
		model.InconsistentVals = append(model.InconsistentVals, "val_a.proofs_state")
	}
	if valB.CurrentState != "" && valB.CurrentState != operability.IntelligenceCalibrationValBStateActive {
		model.InconsistentVals = append(model.InconsistentVals, "val_b.proofs_state")
	}
	if valC.CurrentState != "" && valC.CurrentState != operability.IntelligenceCalibrationValCStateActive {
		model.InconsistentVals = append(model.InconsistentVals, "val_c.proofs_state")
	}
	if valD.CurrentState != "" && valD.CurrentState != operability.IntelligenceCalibrationValDStateActive {
		model.InconsistentVals = append(model.InconsistentVals, "val_d.proofs_state")
	}
	if val0.Point5State == operability.IntelligenceCalibrationPoint5StatePass {
		model.InconsistentVals = append(model.InconsistentVals, "val_0.claims_point_5_pass")
	}
	if valA.Point5State == operability.IntelligenceCalibrationPoint5StatePass {
		model.InconsistentVals = append(model.InconsistentVals, "val_a.claims_point_5_pass")
	}
	if valB.Point5State == operability.IntelligenceCalibrationPoint5StatePass {
		model.InconsistentVals = append(model.InconsistentVals, "val_b.claims_point_5_pass")
	}
	if valC.Point5State == operability.IntelligenceCalibrationPoint5StatePass {
		model.InconsistentVals = append(model.InconsistentVals, "val_c.claims_point_5_pass")
	}
	if valD.Point5State == operability.IntelligenceCalibrationPoint5StatePass {
		model.InconsistentVals = append(model.InconsistentVals, "val_d.claims_point_5_pass")
	}

	switch {
	case len(model.MissingVals) > 0:
		model.DependencyStatus = operability.IntelligenceCalibrationValEDependencyIncomplete
	case len(model.InactiveVals) > 0:
		model.DependencyStatus = operability.IntelligenceCalibrationValEDependencyFail
	case len(model.InconsistentVals) > 0:
		model.DependencyStatus = operability.IntelligenceCalibrationValEDependencyPartial
	default:
		model.DependencyStatus = operability.IntelligenceCalibrationValEDependencyPass
	}
	return model
}

func buildIntelligenceCalibrationValECoherenceReviewModel(val0 intelligenceCalibrationVal0ProofsResponse, valA intelligenceCalibrationValAProofsResponse, valB intelligenceCalibrationValBProofsResponse, valC intelligenceCalibrationValCProofsResponse, valD intelligenceCalibrationValDProofsResponse) operability.IntelligenceCalibrationCrossValCoherenceReview {
	model := operability.IntelligenceCalibrationValECoherenceReviewContract()
	model.CarriedForwardLimitations = collectIntelligenceCalibrationValELimitations(val0, valA, valB, valC, valD)
	model.EvidenceRefs = []string{"val0_proofs", "vala_proofs", "valb_proofs", "valc_proofs", "vald_proofs"}
	model.SurfaceRefs = []string{
		"/v1/intelligence/calibration/val0/proofs",
		"/v1/intelligence/calibration/vala/proofs",
		"/v1/intelligence/calibration/valb/proofs",
		"/v1/intelligence/calibration/valc/proofs",
		"/v1/intelligence/calibration/vald/proofs",
		"/v1/intelligence/calibration/vale/coherence-review",
		"/v1/intelligence/calibration/vale/proofs",
	}
	model.Val0ContractsUsedByLaterVals = valA.Val0DependencyState == operability.IntelligenceCalibrationVal0StateActive &&
		valB.Val0DependencyState == operability.IntelligenceCalibrationVal0StateActive &&
		valC.Val0DependencyState == operability.IntelligenceCalibrationVal0StateActive &&
		valD.Val0DependencyState == operability.IntelligenceCalibrationVal0StateActive
	model.ValAReachabilityVEXGuardrailsRespectedByValD = valD.ValADependencyState == operability.IntelligenceCalibrationValAStateActive &&
		valD.ReachabilityQualityReviewState == operability.IntelligenceCalibrationValDReachabilityQualityStateActive &&
		valD.VEXQualityReviewState == operability.IntelligenceCalibrationValDVEXQualityStateActive
	model.ValBBehavioralLearningGuardrailsRespectedByValD = valD.ValBDependencyState == operability.IntelligenceCalibrationValBStateActive &&
		valD.BehavioralQualityReviewState == operability.IntelligenceCalibrationValDBehavioralQualityStateActive
	model.ValCFeedbackSuppressionFederatedRespectedByValD = valD.ValCDependencyState == operability.IntelligenceCalibrationValCStateActive &&
		valD.FederatedQualityReviewState == operability.IntelligenceCalibrationValDFederatedQualityStateActive
	model.ValDSimulationQualityCoversRequiredPreviousSlices = intelligenceCalibrationValDAllQualityStatesActive(valD)
	model.NoPriorValClaimsPoint5Pass = val0.Point5State != operability.IntelligenceCalibrationPoint5StatePass &&
		valA.Point5State != operability.IntelligenceCalibrationPoint5StatePass &&
		valB.Point5State != operability.IntelligenceCalibrationPoint5StatePass &&
		valC.Point5State != operability.IntelligenceCalibrationPoint5StatePass &&
		valD.Point5State != operability.IntelligenceCalibrationPoint5StatePass
	model.LimitationsCarriedForward = len(model.CarriedForwardLimitations) > 0
	model.AdvisoryProjectionBoundariesPreserved = intelligenceCalibrationValEHasProjectionDisclaimer(val0.ProjectionDisclaimer) &&
		intelligenceCalibrationValEHasProjectionDisclaimer(valA.ProjectionDisclaimer) &&
		intelligenceCalibrationValEHasProjectionDisclaimer(valB.ProjectionDisclaimer) &&
		intelligenceCalibrationValEHasProjectionDisclaimer(valC.ProjectionDisclaimer) &&
		intelligenceCalibrationValEHasProjectionDisclaimer(valD.ProjectionDisclaimer)

	if !model.Val0ContractsUsedByLaterVals {
		model.MissingLinks = append(model.MissingLinks,
			"val0.contracts->vala.reachability_vex_calibration",
			"val0.contracts->valb.behavioral_learning_mode",
			"val0.contracts->valc.feedback_suppression_federated_tuning",
			"val0.contracts->vald.simulation_quality_gate",
		)
	}
	if !model.ValAReachabilityVEXGuardrailsRespectedByValD {
		model.MissingLinks = append(model.MissingLinks, "vala.guardrails->vald.reachability_vex_quality_review")
	}
	if !model.ValBBehavioralLearningGuardrailsRespectedByValD {
		model.MissingLinks = append(model.MissingLinks, "valb.guardrails->vald.behavioral_learning_quality_review")
	}
	if !model.ValCFeedbackSuppressionFederatedRespectedByValD {
		model.MissingLinks = append(model.MissingLinks, "valc.guardrails->vald.feedback_federated_quality_review")
	}
	if !model.ValDSimulationQualityCoversRequiredPreviousSlices {
		model.MissingLinks = append(model.MissingLinks, "vald.simulation_quality_gate->vale.integrated_closure")
	}
	if !model.LimitationsCarriedForward {
		model.MissingLinks = append(model.MissingLinks, "limitations->vale.integrated_closure")
	}
	if !model.NoPriorValClaimsPoint5Pass {
		model.InconsistentLinks = append(model.InconsistentLinks, "prior_vals->point5_not_complete_until_vale")
	}
	if !model.AdvisoryProjectionBoundariesPreserved {
		model.InconsistentLinks = append(model.InconsistentLinks, "prior_vals->projection_only_boundary")
	}

	if len(model.MissingLinks) > 0 || len(model.InconsistentLinks) > 0 {
		model.CoherenceState = operability.IntelligenceCalibrationValEReviewBlocked
	} else {
		model.CoherenceState = operability.IntelligenceCalibrationValEReviewPass
	}
	return model
}

func buildIntelligenceCalibrationValEBoundaryReviewModel(val0 intelligenceCalibrationVal0ProofsResponse, valA intelligenceCalibrationValAProofsResponse, valB intelligenceCalibrationValBProofsResponse, valC intelligenceCalibrationValCProofsResponse, valD intelligenceCalibrationValDProofsResponse) operability.IntelligenceCalibrationIntegratedAdvisoryBoundaryReview {
	model := operability.IntelligenceCalibrationValEBoundaryReviewContract()
	aGuard := operability.IntelligenceCalibrationValAPublicationGuardrailContract()
	aVEX := operability.IntelligenceCalibrationValAVEXCandidateContract()
	bLearning := operability.IntelligenceCalibrationValBLearningModeRuntimeContract()
	bAdoption := operability.IntelligenceCalibrationValBBaselineAdoptionContract()
	bGuard := operability.IntelligenceCalibrationValBGuardrailContract()
	cFeedback := operability.IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	cSuppression := operability.IntelligenceCalibrationValCSuppressionSafetyContract()
	cProposal := operability.IntelligenceCalibrationValCTuningProposalContract()
	cLocal := operability.IntelligenceCalibrationValCLocalChangeReviewContract()
	cFederated := operability.IntelligenceCalibrationValCFederatedSignalWeightingContract()
	cSimilarity := operability.IntelligenceCalibrationValCSimilarityGatingContract()
	cOverride := operability.IntelligenceCalibrationValCLocalOverrideDisciplineContract()
	cPropagation := operability.IntelligenceCalibrationValCPropagationPolicyContract()
	dScoreboard := operability.IntelligenceCalibrationValDQualityScoreboardContract()

	model.CalibrationOutputsRemainProjections = intelligenceCalibrationValEHasProjectionDisclaimer(val0.ProjectionDisclaimer)
	model.ConfidenceScoresRemainAdvisory = val0.ConfidenceState == operability.IntelligenceCalibrationVal0ConfidenceStateActive
	model.ReachabilityInferenceRemainsAdvisory = intelligenceCalibrationValEHasProjectionDisclaimer(valA.ProjectionDisclaimer)
	model.VEXCandidateOutputsRemainCandidates = !aVEX.FinalVEXClaim && !aVEX.PublicationAllowed && valA.PublicationGuardrailState == operability.IntelligenceCalibrationValAPublicationGuardrailStateActive
	model.BehavioralBaselinesRemainAdvisory = valB.GuardrailState == operability.IntelligenceCalibrationValBGuardrailStateActive
	model.LearningModeRemainsBoundedReviewRequired = bLearning.OutputReviewRequired && !bLearning.CanRelaxEnforcement && !bLearning.CanSuppressCriticalAlerts && !bLearning.CanAutoPromoteBaseline
	model.FeedbackRemainsAdvisory = !cFeedback.MutatesIntelligence && !cFeedback.SuppressesSignals && !cFeedback.LowersPriority
	model.SuppressionRemainsCandidateReviewBound = cSuppression.ReviewerRef != "" && !cSuppression.DeletesEvidence && !cSuppression.SuppressesFalseNegativePath && cSuppression.ReopenOnNewEvidence && cSuppression.StrongerEvidenceReopens
	model.FederatedSignalsRemainAdvisory = !cFederated.ProducesLocalSafeState && !cSimilarity.OverridesLocalEvidence && !cPropagation.PropagationAllowed
	model.SimulationMetricsRemainOperationalIndicators = !dScoreboard.ClaimsUniversalQuality && intelligenceCalibrationValEHasProjectionDisclaimer(valD.ProjectionDisclaimer)
	model.IntegratedClosureSummaryRemainsProjectionOnly = true
	model.NoMutationWithoutGovernance = aGuard.GovernanceRequiredForPublication && bAdoption.GovernanceRequired && !cProposal.MutatesActiveCalibration && !cLocal.MutatesActiveCalibration && !cPropagation.MutatesRemoteCalibration
	model.FinalVEXPublicationBlocked = aGuard.FinalClaimBlocked && !aVEX.FinalVEXClaim && valD.VEXQualityReviewState == operability.IntelligenceCalibrationValDVEXQualityStateActive
	model.FederatedOverrideLocalEvidenceBlocked = cOverride.LocalEvidenceWins && !cSimilarity.OverridesLocalEvidence
	model.LearningModeCriticalControlRelaxationBlocked = bGuard.CriticalControlRelaxationBlocked && !bLearning.CanRelaxEnforcement
	model.EvidenceSpineRemainsCanonical = intelligenceCalibrationValEContainsTrimmedString(val0.EvidenceRefs, "evidence_spine")
	model.CheckedSurfaces = []string{
		"/v1/intelligence/calibration/val0/proofs",
		"/v1/intelligence/calibration/vala/publication-guardrail",
		"/v1/intelligence/calibration/valb/safety-guardrails",
		"/v1/intelligence/calibration/valc/feedback-intake",
		"/v1/intelligence/calibration/valc/local-override",
		"/v1/intelligence/calibration/valc/propagation-policy",
		"/v1/intelligence/calibration/vald/quality-scoreboard",
		"/v1/intelligence/calibration/vale/proofs",
	}
	model.EvidenceRefs = []string{"evidence_spine", "val0_proofs", "vala_proofs", "valb_proofs", "valc_proofs", "vald_proofs"}
	if !model.EvidenceSpineRemainsCanonical {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/intelligence/calibration/val0/proofs")
	}
	if !model.VEXCandidateOutputsRemainCandidates || !aGuard.FinalClaimBlocked || aGuard.PublicationAllowed {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/intelligence/calibration/vala/publication-guardrail")
	}
	if !model.BehavioralBaselinesRemainAdvisory || !bAdoption.GovernanceRequired {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/intelligence/calibration/valb/safety-guardrails")
	}
	if !model.FeedbackRemainsAdvisory {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/intelligence/calibration/valc/feedback-intake")
	}
	if !model.FederatedOverrideLocalEvidenceBlocked {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/intelligence/calibration/valc/local-override")
	}
	if !model.FederatedSignalsRemainAdvisory || cPropagation.PropagatesRawLocalEvidence || !cPropagation.RequiresReview {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/intelligence/calibration/valc/propagation-policy")
	}
	if !model.SimulationMetricsRemainOperationalIndicators {
		model.ViolationSurfaces = append(model.ViolationSurfaces, "/v1/intelligence/calibration/vald/quality-scoreboard")
	}
	if len(model.ViolationSurfaces) > 0 {
		model.BoundaryState = operability.IntelligenceCalibrationValEReviewBlocked
	} else {
		model.BoundaryState = operability.IntelligenceCalibrationValEReviewPass
	}
	return model
}

func buildIntelligenceCalibrationValEReachabilityVEXSafetyReviewModel() operability.IntelligenceCalibrationIntegratedReachabilityVEXSafetyReview {
	return buildIntelligenceCalibrationValEReachabilityVEXSafetyReviewModelFromContracts(
		operability.IntelligenceCalibrationValAReachabilityAggregationContract(),
		operability.IntelligenceCalibrationValAExploitabilityCalibrationContract(),
		operability.IntelligenceCalibrationValADowngradeEscalationContract(),
		operability.IntelligenceCalibrationValAVEXCandidateContract(),
		operability.IntelligenceCalibrationValAVEXSufficiencyContract(),
		operability.IntelligenceCalibrationValAExplanationContract(),
		operability.IntelligenceCalibrationValAPublicationGuardrailContract(),
		operability.IntelligenceCalibrationValDReachabilityQualityReviewContract(),
		operability.IntelligenceCalibrationValDVEXQualityReviewContract(),
	)
}

func buildIntelligenceCalibrationValEReachabilityVEXSafetyReviewModelFromContracts(aggregation operability.ReachabilitySignalAggregationContract, exploitability operability.ContextualExploitabilityCalibrationContract, decision operability.DowngradeEscalationDisciplineContract, vexCandidate operability.VEXCandidateCalibrationContract, vexSufficiency operability.VEXEvidenceSufficiencyContract, explanation operability.ReachabilityVEXExplanationContract, publicationGuardrail operability.NoFinalPublicationVEXGuardrailContract, dReachability operability.ReachabilityCalibrationQualityReviewContract, dVEX operability.VEXCandidateQualityReviewContract) operability.IntelligenceCalibrationIntegratedReachabilityVEXSafetyReview {
	model := operability.IntelligenceCalibrationValEReachabilityVEXSafetyReviewContract()
	notAffectedCandidateRequiresBlock := vexCandidate.SuggestedVEXOutcome == operability.IntelligenceCalibrationVEXOutcomeNotAffectedCandidate &&
		(vexCandidate.EvidenceSufficiencyState != operability.IntelligenceCalibrationValAVEXSufficiencySufficient ||
			vexSufficiency.SufficiencyState != operability.IntelligenceCalibrationValAVEXSufficiencySufficient ||
			len(vexSufficiency.MissingEvidenceClasses) > 0 ||
			len(vexSufficiency.StaleEvidenceRefs) > 0 ||
			len(vexSufficiency.UnsupportedEvidenceRefs) > 0)

	model.PackagePresenceOnlyBlocked = !aggregation.PackagePresenceImpliesExploit && dReachability.PackagePresenceOnlyBlocked
	model.RuntimeLoadedOnlyBlocked = !aggregation.RuntimeLoadedImpliesExecution && dReachability.RuntimeLoadedOnlyBlocked
	model.UnsupportedReachabilityRemainsExplicit = dReachability.UnsupportedSignalExplicit
	model.DowngradeRequiresEvidenceExplanationExpiryRollback = decision.EvidenceClass != operability.IntelligenceCalibrationEvidenceUnsupported &&
		len(decision.ReasonCodes) > 0 &&
		strings.TrimSpace(decision.Explanation) != "" &&
		strings.TrimSpace(decision.ExpiresAt) != "" &&
		strings.TrimSpace(decision.RollbackRef) != ""
	model.CriticalDowngradeReviewBound = !decision.AppliesToExcludedCritical && decision.ReviewerRequired && dReachability.ExcludedCriticalDowngradeBlocked
	model.VEXCandidateNotFinalVEX = !vexCandidate.FinalVEXClaim && !vexCandidate.PublicationAllowed
	model.FinalVEXClaimBlocked = !vexCandidate.FinalVEXClaim && publicationGuardrail.FinalClaimBlocked && dVEX.FinalVEXPublicationBlocked
	model.PublicationAllowedBlocked = !vexCandidate.PublicationAllowed && !publicationGuardrail.PublicationAllowed
	model.InsufficientEvidenceBlocksNotAffected = dVEX.NotAffectedRequiresSufficientEvidence && !notAffectedCandidateRequiresBlock
	model.StaleUnsupportedEvidenceBlocksReviewedVEXCandidate = dVEX.StaleOrUnsupportedReviewedBlocked &&
		len(vexSufficiency.StaleEvidenceRefs) == 0 &&
		len(vexSufficiency.UnsupportedEvidenceRefs) == 0
	model.ExplanationDistinguishesNotEvidencedFromSafe = explanation.DistinguishesNotEvidencedFromSafe &&
		!exploitability.CurrentlyNotEvidencedIsSafe &&
		!exploitability.LowEvidenceBecomesSafe

	if !model.PackagePresenceOnlyBlocked {
		model.Blockers = append(model.Blockers, "package presence alone can still imply exploitable reachability")
	}
	if !model.RuntimeLoadedOnlyBlocked {
		model.Blockers = append(model.Blockers, "runtime loaded alone can still imply vulnerable execution")
	}
	if !model.DowngradeRequiresEvidenceExplanationExpiryRollback {
		model.Blockers = append(model.Blockers, "downgrade path lacks evidence, explanation, expiry/limitation, or rollback discipline")
	}
	if !model.FinalVEXClaimBlocked || publicationGuardrail.PublicationAllowed {
		model.Blockers = append(model.Blockers, "final vex publication path remains open inside intelligence calibration")
	}
	if !model.InsufficientEvidenceBlocksNotAffected {
		model.Blockers = append(model.Blockers, "insufficient evidence can still produce not_affected")
	}
	if !model.ExplanationDistinguishesNotEvidencedFromSafe {
		model.Blockers = append(model.Blockers, "reachability/vex explanation no longer distinguishes not-evidenced from safe")
	}
	if len(model.Blockers) > 0 {
		model.ReachabilityVEXState = operability.IntelligenceCalibrationValEReviewBlocked
	} else {
		model.ReachabilityVEXState = operability.IntelligenceCalibrationValEReviewPass
	}
	return model
}

func buildIntelligenceCalibrationValEBehavioralLearningSafetyReviewModel() operability.IntelligenceCalibrationIntegratedBehavioralLearningSafetyReview {
	model := operability.IntelligenceCalibrationValEBehavioralLearningSafetyReviewContract()
	baseline := operability.IntelligenceCalibrationValBBehavioralBaselineContract()
	learning := operability.IntelligenceCalibrationValBLearningModeRuntimeContract()
	threshold := operability.IntelligenceCalibrationValBAnomalyThresholdContract()
	weighting := operability.IntelligenceCalibrationValBCriticalityWeightingContract()
	freshness := operability.IntelligenceCalibrationValBBaselineFreshnessContract()
	adoption := operability.IntelligenceCalibrationValBBaselineAdoptionContract()
	guardrail := operability.IntelligenceCalibrationValBGuardrailContract()
	behavioralQuality := operability.IntelligenceCalibrationValDBehavioralQualityReviewContract()

	model.ActiveBaselineFreshnessFresh = baseline.FreshnessState == operability.IntelligenceCalibrationFreshnessFresh && freshness.FreshnessState == operability.IntelligenceCalibrationFreshnessFresh
	model.StaleUnknownUnsupportedBaselineBlocked = behavioralQuality.StaleOrUnknownBaselineBlocked &&
		baseline.FreshnessState == operability.IntelligenceCalibrationFreshnessFresh &&
		freshness.FreshnessState == operability.IntelligenceCalibrationFreshnessFresh
	startAt, startOK := intelligenceCalibrationValEParseTimestamp(baseline.ObservationWindowStart)
	endAt, endOK := intelligenceCalibrationValEParseTimestamp(baseline.ObservationWindowEnd)
	model.BaselineObservationWindowBoundedAndTimestampValidated = startOK && endOK && endAt.After(startAt)
	learningStart, learningStartOK := intelligenceCalibrationValEParseTimestamp(learning.StartedAt)
	learningEnd, learningEndOK := intelligenceCalibrationValEParseTimestamp(learning.ExpiresAt)
	model.LearningModeTimestampsRFC3339AndBounded = learningStartOK && learningEndOK && learningEnd.After(learningStart) && learning.BoundedDurationConfirmed
	model.LearningModeCriticalControlRelaxationBlocked = !learning.CanRelaxEnforcement && guardrail.CriticalControlRelaxationBlocked
	model.LearningModeCriticalAlertSuppressionBlocked = !learning.CanSuppressCriticalAlerts && guardrail.AutoSuppressionBlocked
	model.LearningModeAutoBaselinePromotionBlocked = !learning.CanAutoPromoteBaseline && guardrail.AutoBaselinePromotionBlocked
	model.ThresholdDecreaseRequiresFNRiskNoteAndReview = threshold.ThresholdChangeDirection != operability.IntelligenceCalibrationValBThresholdDecreaseSensitivity ||
		(threshold.ReviewRequired && strings.TrimSpace(threshold.FalseNegativeRiskNote) != "" && behavioralQuality.ThresholdDecreaseReviewed && behavioralQuality.FNRiskNoteRequired)
	model.LowerPriorityCandidateRequiresReviewerGate = weighting.WeightingAction != operability.IntelligenceCalibrationValBWeightingLowerPriorityCandidate || weighting.ReviewerRequired
	model.BaselineAdoptionRequiresReviewerApprovalRollbackGovernance = adoption.ReviewerRequired &&
		strings.TrimSpace(adoption.ApprovalRef) != "" &&
		strings.TrimSpace(adoption.RollbackRef) != "" &&
		adoption.GovernanceRequired

	if !model.ActiveBaselineFreshnessFresh {
		model.Blockers = append(model.Blockers, "active behavioral baseline is not backed by fresh baseline freshness metadata")
	}
	if !model.BaselineObservationWindowBoundedAndTimestampValidated {
		model.Blockers = append(model.Blockers, "behavioral baseline observation window is not RFC3339-valid and chronologically bounded")
	}
	if !model.LearningModeTimestampsRFC3339AndBounded {
		model.Blockers = append(model.Blockers, "learning mode timestamps are not RFC3339-valid and bounded")
	}
	if !model.LearningModeCriticalControlRelaxationBlocked {
		model.Blockers = append(model.Blockers, "learning mode can still relax critical controls")
	}
	if !model.LearningModeCriticalAlertSuppressionBlocked {
		model.Blockers = append(model.Blockers, "learning mode can still suppress critical alerts")
	}
	if !model.LearningModeAutoBaselinePromotionBlocked {
		model.Blockers = append(model.Blockers, "learning mode can still auto-promote baseline")
	}
	if !model.ThresholdDecreaseRequiresFNRiskNoteAndReview {
		model.Blockers = append(model.Blockers, "threshold decrease can still proceed without FN risk note and review")
	}
	if !model.LowerPriorityCandidateRequiresReviewerGate {
		model.Blockers = append(model.Blockers, "lower-priority candidate can still proceed without reviewer gate")
	}
	if !model.BaselineAdoptionRequiresReviewerApprovalRollbackGovernance {
		model.Blockers = append(model.Blockers, "baseline adoption can still proceed without reviewer, approval ref, rollback ref, and governance requirement")
	}
	if len(model.Blockers) > 0 {
		model.BehavioralLearningState = operability.IntelligenceCalibrationValEReviewBlocked
	} else {
		model.BehavioralLearningState = operability.IntelligenceCalibrationValEReviewPass
	}
	return model
}

func buildIntelligenceCalibrationValEFeedbackFederatedSafetyReviewModel() operability.IntelligenceCalibrationIntegratedFeedbackFederatedSafetyReview {
	model := operability.IntelligenceCalibrationValEFeedbackFederatedSafetyReviewContract()
	feedback := operability.IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	review := operability.IntelligenceCalibrationValCFeedbackReviewCockpitContract()
	proposal := operability.IntelligenceCalibrationValCTuningProposalContract()
	suppression := operability.IntelligenceCalibrationValCSuppressionSafetyContract()
	rollback := operability.IntelligenceCalibrationValCSuppressionRollbackContract()
	local := operability.IntelligenceCalibrationValCLocalChangeReviewContract()
	federated := operability.IntelligenceCalibrationValCFederatedSignalWeightingContract()
	similarity := operability.IntelligenceCalibrationValCSimilarityGatingContract()
	override := operability.IntelligenceCalibrationValCLocalOverrideDisciplineContract()
	propagation := operability.IntelligenceCalibrationValCPropagationPolicyContract()

	model.FeedbackDoesNotMutateIntelligence = !feedback.MutatesIntelligence
	model.FeedbackDoesNotDirectlySuppressSignals = !feedback.SuppressesSignals
	model.FeedbackDoesNotDirectlyLowerPriority = !feedback.LowersPriority
	model.FalseNegativeAndMissedSeverityRemainVisible = review.FalseNegativeVisible && !feedback.RoutedAsNoiseReduction
	model.TuningProposalsRemainAdvisoryReviewBound = !proposal.MutatesActiveCalibration && proposal.ReviewerRequired
	model.SuppressionCandidatesRequireExpiryScopeReviewerRollbackReopen = strings.TrimSpace(suppression.SuppressionScope) != "" &&
		strings.TrimSpace(suppression.ExpiresAt) != "" &&
		strings.TrimSpace(suppression.ReviewerRef) != "" &&
		suppression.ReopenOnNewEvidence &&
		suppression.StrongerEvidenceReopens &&
		rollback.RollbackAvailable &&
		len(rollback.RollbackTriggerConditions) > 0 &&
		strings.TrimSpace(rollback.RollbackSafetyCheck) != ""
	model.SuppressionDoesNotDeleteEvidence = !suppression.DeletesEvidence
	model.SuppressionDoesNotHideFalseNegativePath = !suppression.SuppressesFalseNegativePath
	model.LocalCalibrationReviewDoesNotMutateActiveCalibration = !local.MutatesActiveCalibration
	model.FederatedSignalsDoNotProduceLocalSafeState = !federated.ProducesLocalSafeState
	model.SimilarityGatingDoesNotOverrideLocalEvidence = !similarity.OverridesLocalEvidence
	model.LocalEvidenceWinsOverFederatedSignal = override.LocalEvidenceWins
	model.PropagationDisabledOrAdvisoryOnlyByDefault = !propagation.PropagationAllowed &&
		(propagation.DefaultState == operability.IntelligenceCalibrationValCPropagationDisabled || propagation.DefaultState == operability.IntelligenceCalibrationValCPropagationAdvisoryOnly) &&
		propagation.RequiresRedaction &&
		propagation.RequiresReview &&
		!propagation.MutatesRemoteCalibration
	model.RawLocalEvidenceDoesNotPropagate = !propagation.PropagatesRawLocalEvidence && intelligenceCalibrationValEContainsTrimmedString(propagation.BlockedPayloadClasses, "raw_local_evidence")

	if !model.FeedbackDoesNotMutateIntelligence {
		model.Blockers = append(model.Blockers, "feedback can still mutate intelligence directly")
	}
	if !model.FeedbackDoesNotDirectlySuppressSignals {
		model.Blockers = append(model.Blockers, "feedback can still suppress signals directly")
	}
	if !model.FeedbackDoesNotDirectlyLowerPriority {
		model.Blockers = append(model.Blockers, "feedback can still lower priority directly")
	}
	if !model.FalseNegativeAndMissedSeverityRemainVisible {
		model.Blockers = append(model.Blockers, "false-negative or missed-severity feedback can still disappear into noise reduction")
	}
	if !model.SuppressionCandidatesRequireExpiryScopeReviewerRollbackReopen {
		model.Blockers = append(model.Blockers, "suppression candidate lacks expiry, scope, reviewer, rollback, or reopen-on-new-evidence discipline")
	}
	if !model.SuppressionDoesNotDeleteEvidence {
		model.Blockers = append(model.Blockers, "suppression candidate can still delete evidence")
	}
	if !model.SuppressionDoesNotHideFalseNegativePath {
		model.Blockers = append(model.Blockers, "suppression candidate can still hide false-negative path")
	}
	if !model.FederatedSignalsDoNotProduceLocalSafeState {
		model.Blockers = append(model.Blockers, "federated signal can still produce local safe/not_affected state")
	}
	if !model.SimilarityGatingDoesNotOverrideLocalEvidence || !model.LocalEvidenceWinsOverFederatedSignal {
		model.Blockers = append(model.Blockers, "federated similarity path can still override local evidence")
	}
	if !model.RawLocalEvidenceDoesNotPropagate {
		model.Blockers = append(model.Blockers, "raw local evidence can still propagate")
	}
	if len(model.Blockers) > 0 {
		model.FeedbackFederatedState = operability.IntelligenceCalibrationValEReviewBlocked
	} else {
		model.FeedbackFederatedState = operability.IntelligenceCalibrationValEReviewPass
	}
	return model
}

func buildIntelligenceCalibrationValESimulationQualityReviewModel() operability.IntelligenceCalibrationIntegratedSimulationQualityReview {
	model := operability.IntelligenceCalibrationValESimulationQualityReviewContract()
	harness := operability.IntelligenceCalibrationValDDefensiveSimulationHarnessContract()
	library := operability.IntelligenceCalibrationValDScenarioLibraryContract()
	missed := operability.IntelligenceCalibrationValDMissedDetectionAnalysisContract()
	balance := operability.IntelligenceCalibrationValDFPFNBalanceReviewContract()
	confidence := operability.IntelligenceCalibrationValDConfidenceCalibrationReviewContract()
	vex := operability.IntelligenceCalibrationValDVEXQualityReviewContract()
	reachability := operability.IntelligenceCalibrationValDReachabilityQualityReviewContract()
	behavioral := operability.IntelligenceCalibrationValDBehavioralQualityReviewContract()
	federated := operability.IntelligenceCalibrationValDFederatedQualityReviewContract()
	coverage := operability.IntelligenceCalibrationValDSimulationCoverageReviewContract()
	scoreboard := operability.IntelligenceCalibrationValDQualityScoreboardContract()

	model.HarnessPresentReplayableOrExplicitlyLimited = harness.Replayable || strings.Contains(strings.TrimSpace(harness.LimitationMessage), "replay")
	model.ExpectedAndActualOutcomesRepresented = harness.ExpectedOutcomesPresent && harness.ActualOutcomesPresent
	model.LowSignalAndAdversarialScenariosRepresented = len(library.LowSignalScenarios) > 0 && len(library.AdversarialScenarios) > 0
	model.MissedDetectionIncludesFalseNegativeVisibility = len(missed.FalseNegativeCases) > 0
	model.SuppressionCausedMissAnalysisRepresented = len(missed.SuppressionCausedMissCases) > 0
	model.FPReductionCannotHideFNRisk = !balance.FPReductionClaimed || balance.FNRiskReviewed
	model.ConfidenceReviewBlocksUnsupportedWeakInferenceHighConfidence = confidence.UnsupportedEvidenceHighConfidenceBlocked && confidence.WeaklyInferredHighConfidenceBlocked
	model.VEXQualityBlocksFinalPublicationAndInsufficientNotAffected = vex.FinalVEXPublicationBlocked && vex.NotAffectedRequiresSufficientEvidence
	model.ReachabilityQualityPreservesPackageRuntimeGuardrails = reachability.PackagePresenceOnlyBlocked && reachability.RuntimeLoadedOnlyBlocked
	model.BehavioralQualityPreservesLearningBaselineGuardrails = behavioral.StaleOrUnknownBaselineBlocked && behavioral.LearningModeRelaxationBlocked && behavioral.FNRiskNoteRequired
	model.FederatedQualityPreservesLocalEvidenceBoundary = federated.LocalEvidenceOverrideReviewed && federated.LowSimilarityConfidenceNotIncreased && federated.RawLocalEvidenceNotPropagated
	model.CoverageDoesNotClaimExhaustiveDetection = !coverage.ClaimsExhaustiveDetection && intelligenceCalibrationValEContainsSubstringInTrimmedStrings(coverage.CoverageLimitations, "not exhaustive")
	model.ScoreboardIncludesFPAndFNMetrics = strings.TrimSpace(scoreboard.FalsePositiveRateRef) != "" &&
		strings.TrimSpace(scoreboard.FalseNegativeReviewRateRef) != "" &&
		intelligenceCalibrationValEContainsAllTrimmedStrings(scoreboard.MetricRefs, operability.IntelligenceCalibrationMetricFalsePositiveRate, operability.IntelligenceCalibrationMetricFalseNegativeReviewRate)
	model.ScoreboardDoesNotClaimUniversalIntelligenceQuality = !scoreboard.ClaimsUniversalQuality
	model.MissingCoverage = append(model.MissingCoverage, coverage.MissingScenarioClasses...)
	model.MissingCoverage = append(model.MissingCoverage, coverage.CriticalMissingClasses...)

	if len(coverage.CriticalMissingClasses) > 0 {
		model.Blockers = append(model.Blockers, "simulation coverage is missing critical scenario classes")
	}
	if !model.FPReductionCannotHideFNRisk {
		model.Blockers = append(model.Blockers, "false-positive reduction still hides false-negative risk review")
	}
	if !model.ConfidenceReviewBlocksUnsupportedWeakInferenceHighConfidence {
		model.Blockers = append(model.Blockers, "unsupported or weakly inferred evidence can still become high confidence")
	}
	if !model.VEXQualityBlocksFinalPublicationAndInsufficientNotAffected {
		model.Blockers = append(model.Blockers, "vex quality review still allows final publication or insufficient-evidence not_affected")
	}
	if !model.CoverageDoesNotClaimExhaustiveDetection {
		model.Blockers = append(model.Blockers, "simulation coverage still claims exhaustive detection")
	}
	if !model.ScoreboardIncludesFPAndFNMetrics {
		model.Blockers = append(model.Blockers, "quality scoreboard is missing FP/FN metric coverage")
	}
	if !model.ScoreboardDoesNotClaimUniversalIntelligenceQuality {
		model.Blockers = append(model.Blockers, "quality scoreboard still claims universal intelligence quality")
	}
	if len(model.Blockers) > 0 {
		model.SimulationQualityState = operability.IntelligenceCalibrationValEReviewBlocked
	} else {
		model.SimulationQualityState = operability.IntelligenceCalibrationValEReviewPass
	}
	return model
}

func buildIntelligenceCalibrationValERegressionClosureModel() operability.IntelligenceCalibrationIntegratedRegressionClosure {
	return operability.IntelligenceCalibrationValERegressionClosureContract()
}

func buildIntelligenceCalibrationValEPassRuleModel(val0 intelligenceCalibrationVal0ProofsResponse, valA intelligenceCalibrationValAProofsResponse, valB intelligenceCalibrationValBProofsResponse, valC intelligenceCalibrationValCProofsResponse, valD intelligenceCalibrationValDProofsResponse, prereqValEState string, dependencyClosure operability.IntelligenceCalibrationIntegratedDependencyClosure, coherenceReview operability.IntelligenceCalibrationCrossValCoherenceReview, boundaryReview operability.IntelligenceCalibrationIntegratedAdvisoryBoundaryReview, reachabilityVEXSafety operability.IntelligenceCalibrationIntegratedReachabilityVEXSafetyReview, behavioralLearningSafety operability.IntelligenceCalibrationIntegratedBehavioralLearningSafetyReview, feedbackFederatedSafety operability.IntelligenceCalibrationIntegratedFeedbackFederatedSafetyReview, simulationQuality operability.IntelligenceCalibrationIntegratedSimulationQualityReview, regressionClosure operability.IntelligenceCalibrationIntegratedRegressionClosure) operability.IntelligenceCalibrationPoint5IntegratedPassRule {
	model := operability.IntelligenceCalibrationValEPassRuleContract()
	model.RequiredVals = intelligenceCalibrationValERequiredVals()
	model.PassLimitations = collectIntelligenceCalibrationValELimitations(val0, valA, valB, valC, valD)
	// Pass-rule evaluation is bootstrapped from the prerequisite closure state.
	// Final integrated Val E state is computed separately and gates point_5_state later.
	model.ValEState = prereqValEState
	model.ActiveVals = intelligenceCalibrationValEActiveValNames(val0.Val0State, valA.ValAState, valB.ValBState, valC.ValCState, valD.ValDState, prereqValEState)
	model.MissingVals, model.PartialVals, model.UnsupportedVals = intelligenceCalibrationValEClassifyRequiredVals(val0.Val0State, valA.ValAState, valB.ValBState, valC.ValCState, valD.ValDState, prereqValEState)

	if len(dependencyClosure.MissingVals) > 0 {
		model.PassBlockers = append(model.PassBlockers, "dependency closure missing required vals")
	}
	if len(dependencyClosure.InactiveVals) > 0 {
		model.PassBlockers = append(model.PassBlockers, "dependency closure has inactive vals")
	}
	if len(dependencyClosure.InconsistentVals) > 0 {
		model.PassBlockers = append(model.PassBlockers, "dependency closure has inconsistent vals")
	}
	if len(coherenceReview.MissingLinks) > 0 {
		model.PassBlockers = append(model.PassBlockers, "cross-val coherence review has missing critical links")
	}
	if len(coherenceReview.InconsistentLinks) > 0 {
		model.PassBlockers = append(model.PassBlockers, "cross-val coherence review has inconsistent links")
	}
	if len(boundaryReview.ViolationSurfaces) > 0 {
		model.PassBlockers = append(model.PassBlockers, "advisory/canonical boundary review has violations")
	}
	if len(reachabilityVEXSafety.Blockers) > 0 {
		model.PassBlockers = append(model.PassBlockers, "integrated reachability/vex safety review has blockers")
	}
	if len(behavioralLearningSafety.Blockers) > 0 {
		model.PassBlockers = append(model.PassBlockers, "integrated behavioral/learning safety review has blockers")
	}
	if len(feedbackFederatedSafety.Blockers) > 0 {
		model.PassBlockers = append(model.PassBlockers, "integrated feedback/federated safety review has blockers")
	}
	if len(simulationQuality.Blockers) > 0 {
		model.PassBlockers = append(model.PassBlockers, "integrated simulation/quality review has blockers")
	}
	if len(regressionClosure.CriticalMissingCategories) > 0 {
		model.PassBlockers = append(model.PassBlockers, "regression closure has critical missing categories")
	}
	if prereqValEState != operability.IntelligenceCalibrationValEStateActive {
		model.PassBlockers = append(model.PassBlockers, "val_e prerequisites are not fully active")
	}
	model.PassWarnings = append(model.PassWarnings,
		"Integrated closure remains projection-only and does not replace canonical truth or later governance.",
	)
	model.Point5State = operability.IntelligenceCalibrationPoint5StateNotComplete
	model.PassCriteriaMet = false
	if len(model.PassBlockers) == 0 &&
		len(model.MissingVals) == 0 &&
		len(model.PartialVals) == 0 &&
		len(model.UnsupportedVals) == 0 &&
		prereqValEState == operability.IntelligenceCalibrationValEStateActive {
		model.Point5State = operability.IntelligenceCalibrationPoint5StatePass
		model.PassCriteriaMet = true
	}
	return model
}

func buildIntelligenceCalibrationValEModelsCurrentState(models intelligenceCalibrationValEModels) string {
	return operability.EvaluateIntelligenceCalibrationValEProofsState(
		models.val0.Val0State,
		models.valA.ValAState,
		models.valB.ValBState,
		models.valC.ValCState,
		models.valD.ValDState,
		models.dependencyClosureState,
		models.coherenceReviewState,
		models.passRuleState,
		models.boundaryReviewState,
		models.reachabilityVEXSafetyState,
		models.behavioralLearningSafetyState,
		models.feedbackFederatedSafetyState,
		models.simulationQualityState,
		models.regressionClosureState,
		models.point5State,
		models.surfaceRefs,
		models.evidenceRefs,
		models.limitations,
		intelligenceCalibrationValEProjectionDisclaimer(),
	)
}

func buildIntelligenceCalibrationValEModels() intelligenceCalibrationValEModels {
	val0 := buildIntelligenceCalibrationVal0Proofs()
	valA := buildIntelligenceCalibrationValAProofs()
	valB := buildIntelligenceCalibrationValBProofs()
	valC := buildIntelligenceCalibrationValCProofs()
	valD := buildIntelligenceCalibrationValDProofs()

	dependencyClosure := buildIntelligenceCalibrationValEDependencyClosureModel(val0, valA, valB, valC, valD)
	dependencyClosureState := operability.EvaluateIntelligenceCalibrationValEDependencyClosureState(dependencyClosure)
	coherenceReview := buildIntelligenceCalibrationValECoherenceReviewModel(val0, valA, valB, valC, valD)
	coherenceReviewState := operability.EvaluateIntelligenceCalibrationValECoherenceReviewState(coherenceReview)
	boundaryReview := buildIntelligenceCalibrationValEBoundaryReviewModel(val0, valA, valB, valC, valD)
	boundaryReviewState := operability.EvaluateIntelligenceCalibrationValEBoundaryReviewState(boundaryReview)
	reachabilityVEXSafety := buildIntelligenceCalibrationValEReachabilityVEXSafetyReviewModel()
	reachabilityVEXSafetyState := operability.EvaluateIntelligenceCalibrationValEReachabilityVEXSafetyState(reachabilityVEXSafety)
	behavioralLearningSafety := buildIntelligenceCalibrationValEBehavioralLearningSafetyReviewModel()
	behavioralLearningSafetyState := operability.EvaluateIntelligenceCalibrationValEBehavioralLearningSafetyState(behavioralLearningSafety)
	feedbackFederatedSafety := buildIntelligenceCalibrationValEFeedbackFederatedSafetyReviewModel()
	feedbackFederatedSafetyState := operability.EvaluateIntelligenceCalibrationValEFeedbackFederatedSafetyState(feedbackFederatedSafety)
	simulationQualityReview := buildIntelligenceCalibrationValESimulationQualityReviewModel()
	simulationQualityState := operability.EvaluateIntelligenceCalibrationValESimulationQualityState(simulationQualityReview)
	regressionClosure := buildIntelligenceCalibrationValERegressionClosureModel()
	regressionClosureState := operability.EvaluateIntelligenceCalibrationValERegressionClosureState(regressionClosure)

	prereqValEState := operability.EvaluateIntelligenceCalibrationValEPrerequisiteState(
		val0.Val0State,
		valA.ValAState,
		valB.ValBState,
		valC.ValCState,
		valD.ValDState,
		dependencyClosureState,
		coherenceReviewState,
		boundaryReviewState,
		reachabilityVEXSafetyState,
		behavioralLearningSafetyState,
		feedbackFederatedSafetyState,
		simulationQualityState,
		regressionClosureState,
	)

	passRule := buildIntelligenceCalibrationValEPassRuleModel(
		val0,
		valA,
		valB,
		valC,
		valD,
		prereqValEState,
		dependencyClosure,
		coherenceReview,
		boundaryReview,
		reachabilityVEXSafety,
		behavioralLearningSafety,
		feedbackFederatedSafety,
		simulationQualityReview,
		regressionClosure,
	)
	passRuleState := operability.EvaluateIntelligenceCalibrationValEPassRuleState(passRule)
	valEState := operability.EvaluateIntelligenceCalibrationValEState(
		val0.Val0State,
		valA.ValAState,
		valB.ValBState,
		valC.ValCState,
		valD.ValDState,
		dependencyClosureState,
		coherenceReviewState,
		passRuleState,
		boundaryReviewState,
		reachabilityVEXSafetyState,
		behavioralLearningSafetyState,
		feedbackFederatedSafetyState,
		simulationQualityState,
		regressionClosureState,
	)

	dependencyClosure.ValEState = valEState
	passRule.ValEState = prereqValEState
	passRule.ActiveVals = intelligenceCalibrationValEActiveValNames(val0.Val0State, valA.ValAState, valB.ValBState, valC.ValCState, valD.ValDState, prereqValEState)
	passRule.MissingVals, passRule.PartialVals, passRule.UnsupportedVals = intelligenceCalibrationValEClassifyRequiredVals(val0.Val0State, valA.ValAState, valB.ValBState, valC.ValCState, valD.ValDState, prereqValEState)
	passRule.Point5State = operability.IntelligenceCalibrationPoint5StateNotComplete
	passRule.PassCriteriaMet = false
	if len(passRule.PassBlockers) == 0 &&
		len(passRule.MissingVals) == 0 &&
		len(passRule.PartialVals) == 0 &&
		len(passRule.UnsupportedVals) == 0 &&
		valEState == operability.IntelligenceCalibrationValEStateActive {
		passRule.Point5State = operability.IntelligenceCalibrationPoint5StatePass
		passRule.PassCriteriaMet = true
	}
	passRuleState = operability.EvaluateIntelligenceCalibrationValEPassRuleState(passRule)
	valEState = operability.EvaluateIntelligenceCalibrationValEState(
		val0.Val0State,
		valA.ValAState,
		valB.ValBState,
		valC.ValCState,
		valD.ValDState,
		dependencyClosureState,
		coherenceReviewState,
		passRuleState,
		boundaryReviewState,
		reachabilityVEXSafetyState,
		behavioralLearningSafetyState,
		feedbackFederatedSafetyState,
		simulationQualityState,
		regressionClosureState,
	)

	dependencyClosure.ValEState = valEState
	passRule.ValEState = prereqValEState
	point5State := operability.IntelligenceCalibrationPoint5StateNotComplete
	if valEState == operability.IntelligenceCalibrationValEStateActive && passRuleState == operability.IntelligenceCalibrationValEPassRuleStateActive {
		point5State = operability.IntelligenceCalibrationPoint5StatePass
	}
	passRule.Point5State = point5State
	passRule.PassCriteriaMet = point5State == operability.IntelligenceCalibrationPoint5StatePass
	passRuleState = operability.EvaluateIntelligenceCalibrationValEPassRuleState(passRule)

	return intelligenceCalibrationValEModels{
		val0:                          val0,
		valA:                          valA,
		valB:                          valB,
		valC:                          valC,
		valD:                          valD,
		dependencyClosure:             dependencyClosure,
		coherenceReview:               coherenceReview,
		passRule:                      passRule,
		boundaryReview:                boundaryReview,
		reachabilityVEXSafety:         reachabilityVEXSafety,
		behavioralLearningSafety:      behavioralLearningSafety,
		feedbackFederatedSafety:       feedbackFederatedSafety,
		simulationQualityReview:       simulationQualityReview,
		regressionClosure:             regressionClosure,
		dependencyClosureState:        dependencyClosureState,
		coherenceReviewState:          coherenceReviewState,
		passRuleState:                 passRuleState,
		boundaryReviewState:           boundaryReviewState,
		reachabilityVEXSafetyState:    reachabilityVEXSafetyState,
		behavioralLearningSafetyState: behavioralLearningSafetyState,
		feedbackFederatedSafetyState:  feedbackFederatedSafetyState,
		simulationQualityState:        simulationQualityState,
		regressionClosureState:        regressionClosureState,
		valEState:                     valEState,
		point5State:                   point5State,
		surfaceRefs:                   intelligenceCalibrationValEAllSurfaceRefs(),
		evidenceRefs:                  intelligenceCalibrationValEEvidenceRefs(),
		limitations:                   collectIntelligenceCalibrationValELimitations(val0, valA, valB, valC, valD),
	}
}

func (s server) intelligenceCalibrationValEDependencyClosureHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	models := buildIntelligenceCalibrationValEModels()
	httpjson.Write(w, http.StatusOK, intelligenceCalibrationValEDependencyClosureResponse{
		SchemaVersion: intelligenceCalibrationValEDependencyClosureSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.dependencyClosureState,
		Model:         models.dependencyClosure,
		RouteRefs:     []string{"/v1/intelligence/calibration/vale/dependency-closure"},
		Limitations:   models.limitations,
	})
}

func (s server) intelligenceCalibrationValECoherenceReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	models := buildIntelligenceCalibrationValEModels()
	httpjson.Write(w, http.StatusOK, intelligenceCalibrationValECoherenceReviewResponse{
		SchemaVersion: intelligenceCalibrationValECoherenceReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.coherenceReviewState,
		Model:         models.coherenceReview,
		RouteRefs:     []string{"/v1/intelligence/calibration/vale/coherence-review"},
		Limitations:   models.coherenceReview.CarriedForwardLimitations,
	})
}

func (s server) intelligenceCalibrationValEPassRuleHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	models := buildIntelligenceCalibrationValEModels()
	httpjson.Write(w, http.StatusOK, intelligenceCalibrationValEPassRuleResponse{
		SchemaVersion: intelligenceCalibrationValEPassRuleSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.passRuleState,
		Model:         models.passRule,
		RouteRefs:     []string{"/v1/intelligence/calibration/vale/pass-rule"},
		Limitations:   models.passRule.PassLimitations,
	})
}

func (s server) intelligenceCalibrationValEBoundaryReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	models := buildIntelligenceCalibrationValEModels()
	httpjson.Write(w, http.StatusOK, intelligenceCalibrationValEBoundaryReviewResponse{
		SchemaVersion: intelligenceCalibrationValEBoundaryReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.boundaryReviewState,
		Model:         models.boundaryReview,
		RouteRefs:     []string{"/v1/intelligence/calibration/vale/advisory-boundary"},
		Limitations:   []string{models.boundaryReview.LimitationMessage},
	})
}

func (s server) intelligenceCalibrationValEReachabilityVEXSafetyHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	models := buildIntelligenceCalibrationValEModels()
	httpjson.Write(w, http.StatusOK, intelligenceCalibrationValEReachabilityVEXSafetyResponse{
		SchemaVersion: intelligenceCalibrationValEReachabilityVEXSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.reachabilityVEXSafetyState,
		Model:         models.reachabilityVEXSafety,
		RouteRefs:     []string{"/v1/intelligence/calibration/vale/reachability-vex-safety"},
		Limitations:   models.reachabilityVEXSafety.Limitations,
	})
}

func (s server) intelligenceCalibrationValEBehavioralLearningSafetyHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	models := buildIntelligenceCalibrationValEModels()
	httpjson.Write(w, http.StatusOK, intelligenceCalibrationValEBehavioralLearningSafetyResponse{
		SchemaVersion: intelligenceCalibrationValEBehavioralLearningSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.behavioralLearningSafetyState,
		Model:         models.behavioralLearningSafety,
		RouteRefs:     []string{"/v1/intelligence/calibration/vale/behavioral-learning-safety"},
		Limitations:   models.behavioralLearningSafety.Limitations,
	})
}

func (s server) intelligenceCalibrationValEFeedbackFederatedSafetyHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	models := buildIntelligenceCalibrationValEModels()
	httpjson.Write(w, http.StatusOK, intelligenceCalibrationValEFeedbackFederatedSafetyResponse{
		SchemaVersion: intelligenceCalibrationValEFeedbackFederatedSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.feedbackFederatedSafetyState,
		Model:         models.feedbackFederatedSafety,
		RouteRefs:     []string{"/v1/intelligence/calibration/vale/feedback-federated-safety"},
		Limitations:   models.feedbackFederatedSafety.Limitations,
	})
}

func (s server) intelligenceCalibrationValESimulationQualityReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	models := buildIntelligenceCalibrationValEModels()
	httpjson.Write(w, http.StatusOK, intelligenceCalibrationValESimulationQualityResponse{
		SchemaVersion: intelligenceCalibrationValESimulationQualitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.simulationQualityState,
		Model:         models.simulationQualityReview,
		RouteRefs:     []string{"/v1/intelligence/calibration/vale/simulation-quality-review"},
		Limitations:   models.simulationQualityReview.Limitations,
	})
}

func (s server) intelligenceCalibrationValERegressionClosureHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	models := buildIntelligenceCalibrationValEModels()
	httpjson.Write(w, http.StatusOK, intelligenceCalibrationValERegressionClosureResponse{
		SchemaVersion: intelligenceCalibrationValERegressionClosureSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  models.regressionClosureState,
		Model:         models.regressionClosure,
		RouteRefs:     []string{"/v1/intelligence/calibration/vale/regression-closure"},
		Limitations:   []string{models.regressionClosure.LimitationMessage},
	})
}

func (s server) intelligenceCalibrationValEProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValEProofs())
}

func buildIntelligenceCalibrationValEProofs() intelligenceCalibrationValEProofsResponse {
	models := buildIntelligenceCalibrationValEModels()
	currentState := buildIntelligenceCalibrationValEModelsCurrentState(models)

	return intelligenceCalibrationValEProofsResponse{
		SchemaVersion:                  intelligenceCalibrationValEProofsSchema,
		GeneratedAt:                    publicSampleTime(),
		CurrentState:                   currentState,
		Val0DependencyState:            models.val0.CurrentState,
		Val0FoundationState:            models.val0.Val0State,
		ValADependencyState:            models.valA.CurrentState,
		ValAReachabilityVEXState:       models.valA.ValAState,
		ValBDependencyState:            models.valB.CurrentState,
		ValBBehavioralLearningState:    models.valB.ValBState,
		ValCDependencyState:            models.valC.CurrentState,
		ValCFeedbackFederatedState:     models.valC.ValCState,
		ValDDependencyState:            models.valD.CurrentState,
		ValDSimulationQualityGateState: models.valD.ValDState,
		ValEState:                      models.valEState,
		DependencyClosureState:         models.dependencyClosureState,
		CoherenceReviewState:           models.coherenceReviewState,
		Point5State:                    models.point5State,
		PassCriteriaMet:                models.passRule.PassCriteriaMet,
		PassBlockers:                   models.passRule.PassBlockers,
		PassWarnings:                   models.passRule.PassWarnings,
		PassLimitations:                models.passRule.PassLimitations,
		AdvisoryBoundaryState:          models.boundaryReviewState,
		ReachabilityVEXSafetyState:     models.reachabilityVEXSafetyState,
		BehavioralLearningSafetyState:  models.behavioralLearningSafetyState,
		FeedbackFederatedSafetyState:   models.feedbackFederatedSafetyState,
		SimulationQualityReviewState:   models.simulationQualityState,
		RegressionClosureState:         models.regressionClosureState,
		EvidenceRefs:                   models.evidenceRefs,
		SurfaceRefs:                    models.surfaceRefs,
		Limitations:                    models.limitations,
		ProjectionDisclaimer:           intelligenceCalibrationValEProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val E integrates Val 0 through Val D into a single fail-closed intelligence calibration closure and is the only Točka 5 slice allowed to mark point_5_state pass.",
			"Integrated closure stays bounded to dependency closure, coherence, advisory/canonical boundary review, reachability/VEX safety, behavioral/learning safety, feedback/federated safety, simulation/quality review, and regression closure over existing proofs.",
			"Even at pass, the integrated summary remains projection-only and not canonical truth; no intelligence calibration surface gains mutation, suppression, priority, VEX publication, or governance authority on its own.",
		},
	}
}

package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	intelligenceCalibrationValBBaselineSchema    = "point5.intelligence_calibration.valb.behavioral_baseline.v1"
	intelligenceCalibrationValBLearningSchema    = "point5.intelligence_calibration.valb.learning_mode_runtime.v1"
	intelligenceCalibrationValBThresholdSchema   = "point5.intelligence_calibration.valb.anomaly_threshold.v1"
	intelligenceCalibrationValBDriftSchema       = "point5.intelligence_calibration.valb.drift_sensitivity.v1"
	intelligenceCalibrationValBWeightingSchema   = "point5.intelligence_calibration.valb.criticality_weighting.v1"
	intelligenceCalibrationValBFreshnessSchema   = "point5.intelligence_calibration.valb.baseline_freshness.v1"
	intelligenceCalibrationValBAdoptionSchema    = "point5.intelligence_calibration.valb.baseline_adoption.v1"
	intelligenceCalibrationValBExplanationSchema = "point5.intelligence_calibration.valb.behavioral_explanations.v1"
	intelligenceCalibrationValBGuardrailSchema   = "point5.intelligence_calibration.valb.behavioral_guardrails.v1"
	intelligenceCalibrationValBProofsSchema      = "point5.intelligence_calibration.valb.proofs.v1"
)

type intelligenceCalibrationValBBaselineResponse struct {
	SchemaVersion string                                        `json:"schema_version"`
	GeneratedAt   time.Time                                     `json:"generated_at"`
	CurrentState  string                                        `json:"current_state"`
	Model         operability.BehavioralBaselineProfileContract `json:"model"`
	RouteRefs     []string                                      `json:"route_refs,omitempty"`
	Limitations   []string                                      `json:"limitations,omitempty"`
}

type intelligenceCalibrationValBLearningResponse struct {
	SchemaVersion string                                            `json:"schema_version"`
	GeneratedAt   time.Time                                         `json:"generated_at"`
	CurrentState  string                                            `json:"current_state"`
	Model         operability.LearningModeRuntimeDisciplineContract `json:"model"`
	RouteRefs     []string                                          `json:"route_refs,omitempty"`
	Limitations   []string                                          `json:"limitations,omitempty"`
}

type intelligenceCalibrationValBThresholdResponse struct {
	SchemaVersion string                                          `json:"schema_version"`
	GeneratedAt   time.Time                                       `json:"generated_at"`
	CurrentState  string                                          `json:"current_state"`
	Model         operability.AnomalyThresholdCalibrationContract `json:"model"`
	RouteRefs     []string                                        `json:"route_refs,omitempty"`
	Limitations   []string                                        `json:"limitations,omitempty"`
}

type intelligenceCalibrationValBDriftResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         operability.DriftSensitivityScalingContract `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationValBWeightingResponse struct {
	SchemaVersion string                                        `json:"schema_version"`
	GeneratedAt   time.Time                                     `json:"generated_at"`
	CurrentState  string                                        `json:"current_state"`
	Model         operability.CriticalityAwareWeightingContract `json:"model"`
	RouteRefs     []string                                      `json:"route_refs,omitempty"`
	Limitations   []string                                      `json:"limitations,omitempty"`
}

type intelligenceCalibrationValBFreshnessResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         operability.BaselineFreshnessExpiryContract `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationValBAdoptionResponse struct {
	SchemaVersion string                                     `json:"schema_version"`
	GeneratedAt   time.Time                                  `json:"generated_at"`
	CurrentState  string                                     `json:"current_state"`
	Model         operability.BaselineAdoptionReviewContract `json:"model"`
	RouteRefs     []string                                   `json:"route_refs,omitempty"`
	Limitations   []string                                   `json:"limitations,omitempty"`
}

type intelligenceCalibrationValBExplanationResponse struct {
	SchemaVersion string                                               `json:"schema_version"`
	GeneratedAt   time.Time                                            `json:"generated_at"`
	CurrentState  string                                               `json:"current_state"`
	Model         operability.BehavioralCalibrationExplanationContract `json:"model"`
	RouteRefs     []string                                             `json:"route_refs,omitempty"`
	Limitations   []string                                             `json:"limitations,omitempty"`
}

type intelligenceCalibrationValBGuardrailResponse struct {
	SchemaVersion string                                                   `json:"schema_version"`
	GeneratedAt   time.Time                                                `json:"generated_at"`
	CurrentState  string                                                   `json:"current_state"`
	Model         operability.BehavioralCalibrationSafetyGuardrailContract `json:"model"`
	RouteRefs     []string                                                 `json:"route_refs,omitempty"`
	Limitations   []string                                                 `json:"limitations,omitempty"`
}

type intelligenceCalibrationValBProofsResponse struct {
	SchemaVersion            string    `json:"schema_version"`
	GeneratedAt              time.Time `json:"generated_at"`
	CurrentState             string    `json:"current_state"`
	Val0DependencyState      string    `json:"val_0_dependency_state"`
	Val0FoundationState      string    `json:"val_0_foundation_state"`
	ValADependencyState      string    `json:"val_a_dependency_state"`
	ValAReachabilityVEXState string    `json:"val_a_reachability_vex_state"`
	ValBState                string    `json:"val_b_state"`
	Point5State              string    `json:"point_5_state"`
	BehavioralBaselineState  string    `json:"behavioral_baseline_profile_state"`
	LearningRuntimeState     string    `json:"learning_mode_runtime_state"`
	ThresholdState           string    `json:"anomaly_threshold_calibration_state"`
	DriftState               string    `json:"drift_sensitivity_scaling_state"`
	WeightingState           string    `json:"criticality_aware_weighting_state"`
	BaselineFreshnessState   string    `json:"baseline_freshness_expiry_state"`
	BaselineAdoptionState    string    `json:"baseline_adoption_review_state"`
	ExplanationState         string    `json:"behavioral_explanation_state"`
	GuardrailState           string    `json:"behavioral_safety_guardrail_state"`
	WhyPoint5NotPass         []string  `json:"why_point_5_not_pass,omitempty"`
	SurfaceRefs              []string  `json:"surface_refs,omitempty"`
	EvidenceRefs             []string  `json:"evidence_refs,omitempty"`
	Limitations              []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer     string    `json:"projection_disclaimer"`
	IntegrationSummary       []string  `json:"integration_summary,omitempty"`
}

func intelligenceCalibrationValBAllSurfaceRefs() []string {
	return []string{
		"/v1/intelligence/calibration/valb/behavioral-baseline",
		"/v1/intelligence/calibration/valb/learning-mode-runtime",
		"/v1/intelligence/calibration/valb/anomaly-thresholds",
		"/v1/intelligence/calibration/valb/drift-sensitivity",
		"/v1/intelligence/calibration/valb/criticality-weighting",
		"/v1/intelligence/calibration/valb/baseline-freshness",
		"/v1/intelligence/calibration/valb/baseline-adoption",
		"/v1/intelligence/calibration/valb/explanations",
		"/v1/intelligence/calibration/valb/safety-guardrails",
		"/v1/intelligence/calibration/valb/proofs",
	}
}

func intelligenceCalibrationValBEvidenceRefs() []string {
	return []string{
		"val0_proofs",
		"vala_proofs",
		"behavioral_baseline_profile",
		"learning_mode_runtime_discipline",
		"anomaly_threshold_calibration",
		"drift_sensitivity_scaling",
		"criticality_aware_weighting",
		"baseline_freshness_expiry",
		"baseline_adoption_review",
		"behavioral_explanation_payload",
		"behavioral_safety_guardrail",
		"evidence_spine",
	}
}

func intelligenceCalibrationValBProjectionDisclaimer() string {
	return "projection_only not_canonical_truth advisory_behavioral_baseline_learning_mode"
}

func (s server) intelligenceCalibrationValBBaselineHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValBBaseline())
}

func (s server) intelligenceCalibrationValBLearningHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValBLearning())
}

func (s server) intelligenceCalibrationValBThresholdHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValBThreshold())
}

func (s server) intelligenceCalibrationValBDriftHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValBDrift())
}

func (s server) intelligenceCalibrationValBWeightingHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValBWeighting())
}

func (s server) intelligenceCalibrationValBFreshnessHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValBFreshness())
}

func (s server) intelligenceCalibrationValBAdoptionHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValBAdoption())
}

func (s server) intelligenceCalibrationValBExplanationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValBExplanation())
}

func (s server) intelligenceCalibrationValBGuardrailHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValBGuardrail())
}

func (s server) intelligenceCalibrationValBProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValBProofs())
}

func buildIntelligenceCalibrationValBBaseline() intelligenceCalibrationValBBaselineResponse {
	model := operability.IntelligenceCalibrationValBBehavioralBaselineContract()
	return intelligenceCalibrationValBBaselineResponse{
		SchemaVersion: intelligenceCalibrationValBBaselineSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valb/behavioral-baseline"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValBLearning() intelligenceCalibrationValBLearningResponse {
	model := operability.IntelligenceCalibrationValBLearningModeRuntimeContract()
	return intelligenceCalibrationValBLearningResponse{
		SchemaVersion: intelligenceCalibrationValBLearningSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValBLearningRuntimeState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valb/learning-mode-runtime"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValBThreshold() intelligenceCalibrationValBThresholdResponse {
	model := operability.IntelligenceCalibrationValBAnomalyThresholdContract()
	return intelligenceCalibrationValBThresholdResponse{
		SchemaVersion: intelligenceCalibrationValBThresholdSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValBThresholdState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valb/anomaly-thresholds"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValBDrift() intelligenceCalibrationValBDriftResponse {
	model := operability.IntelligenceCalibrationValBDriftSensitivityContract()
	return intelligenceCalibrationValBDriftResponse{
		SchemaVersion: intelligenceCalibrationValBDriftSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValBDriftState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valb/drift-sensitivity"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValBWeighting() intelligenceCalibrationValBWeightingResponse {
	model := operability.IntelligenceCalibrationValBCriticalityWeightingContract()
	return intelligenceCalibrationValBWeightingResponse{
		SchemaVersion: intelligenceCalibrationValBWeightingSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValBWeightingState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valb/criticality-weighting"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValBFreshness() intelligenceCalibrationValBFreshnessResponse {
	model := operability.IntelligenceCalibrationValBBaselineFreshnessContract()
	return intelligenceCalibrationValBFreshnessResponse{
		SchemaVersion: intelligenceCalibrationValBFreshnessSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valb/baseline-freshness"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValBAdoption() intelligenceCalibrationValBAdoptionResponse {
	model := operability.IntelligenceCalibrationValBBaselineAdoptionContract()
	return intelligenceCalibrationValBAdoptionResponse{
		SchemaVersion: intelligenceCalibrationValBAdoptionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValBBaselineAdoptionState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valb/baseline-adoption"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValBExplanation() intelligenceCalibrationValBExplanationResponse {
	model := operability.IntelligenceCalibrationValBExplanationContract()
	return intelligenceCalibrationValBExplanationResponse{
		SchemaVersion: intelligenceCalibrationValBExplanationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValBExplanationState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valb/explanations"},
		Limitations:   []string{model.UncertaintyNote, model.NextStep},
	}
}

func buildIntelligenceCalibrationValBGuardrail() intelligenceCalibrationValBGuardrailResponse {
	model := operability.IntelligenceCalibrationValBGuardrailContract()
	return intelligenceCalibrationValBGuardrailResponse{
		SchemaVersion: intelligenceCalibrationValBGuardrailSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValBGuardrailState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valb/safety-guardrails"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValBProofs() intelligenceCalibrationValBProofsResponse {
	val0 := buildIntelligenceCalibrationVal0Proofs()
	valA := buildIntelligenceCalibrationValAProofs()
	baseline := operability.IntelligenceCalibrationValBBehavioralBaselineContract()
	learning := operability.IntelligenceCalibrationValBLearningModeRuntimeContract()
	threshold := operability.IntelligenceCalibrationValBAnomalyThresholdContract()
	drift := operability.IntelligenceCalibrationValBDriftSensitivityContract()
	weighting := operability.IntelligenceCalibrationValBCriticalityWeightingContract()
	freshness := operability.IntelligenceCalibrationValBBaselineFreshnessContract()
	adoption := operability.IntelligenceCalibrationValBBaselineAdoptionContract()
	explanation := operability.IntelligenceCalibrationValBExplanationContract()
	guardrail := operability.IntelligenceCalibrationValBGuardrailContract()

	baselineState := operability.EvaluateIntelligenceCalibrationValBBehavioralBaselineState(baseline)
	learningState := operability.EvaluateIntelligenceCalibrationValBLearningRuntimeState(learning)
	thresholdState := operability.EvaluateIntelligenceCalibrationValBThresholdState(threshold)
	driftState := operability.EvaluateIntelligenceCalibrationValBDriftState(drift)
	weightingState := operability.EvaluateIntelligenceCalibrationValBWeightingState(weighting)
	freshnessState := operability.EvaluateIntelligenceCalibrationValBBaselineFreshnessState(freshness)
	adoptionState := operability.EvaluateIntelligenceCalibrationValBBaselineAdoptionState(adoption)
	explanationState := operability.EvaluateIntelligenceCalibrationValBExplanationState(explanation)
	guardrailState := operability.EvaluateIntelligenceCalibrationValBGuardrailState(guardrail)

	surfaceRefs := intelligenceCalibrationValBAllSurfaceRefs()
	evidenceRefs := intelligenceCalibrationValBEvidenceRefs()
	limitations := []string{
		"Val B proves only Behavioral Baseline & Learning Mode readiness and does not claim complete intelligence calibration, feedback tuning, or integrated closure.",
		"Behavioral baselines, learning mode runtime, threshold calibration, drift scaling, and criticality weighting remain advisory projections over canonical evidence and cannot auto-suppress, auto-promote, mutate enforcement, or mutate canonical priority.",
	}
	whyPoint5NotPass := []string{
		"Točka 5 remains not complete because later waves still need feedback/suppression/federated tuning, simulation, final calibration gate, and integrated closure.",
		"Val B adds bounded behavioral baseline and learning mode runtime discipline only; it does not add automatic suppression, enforcement mutation, or final authority.",
	}

	valBState := operability.EvaluateIntelligenceCalibrationValBState(
		val0.CurrentState,
		val0.Val0State,
		valA.CurrentState,
		valA.ValAState,
		baselineState,
		learningState,
		thresholdState,
		driftState,
		weightingState,
		freshnessState,
		adoptionState,
		explanationState,
		guardrailState,
	)
	currentState := operability.EvaluateIntelligenceCalibrationValBProofsState(
		val0.CurrentState,
		val0.Val0State,
		valA.CurrentState,
		valA.ValAState,
		baselineState,
		learningState,
		thresholdState,
		driftState,
		weightingState,
		freshnessState,
		adoptionState,
		explanationState,
		guardrailState,
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint5NotPass,
		intelligenceCalibrationValBProjectionDisclaimer(),
	)

	return intelligenceCalibrationValBProofsResponse{
		SchemaVersion:            intelligenceCalibrationValBProofsSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             currentState,
		Val0DependencyState:      val0.CurrentState,
		Val0FoundationState:      val0.Val0State,
		ValADependencyState:      valA.CurrentState,
		ValAReachabilityVEXState: valA.ValAState,
		ValBState:                valBState,
		Point5State:              operability.IntelligenceCalibrationPoint5StateNotComplete,
		BehavioralBaselineState:  baselineState,
		LearningRuntimeState:     learningState,
		ThresholdState:           thresholdState,
		DriftState:               driftState,
		WeightingState:           weightingState,
		BaselineFreshnessState:   freshnessState,
		BaselineAdoptionState:    adoptionState,
		ExplanationState:         explanationState,
		GuardrailState:           guardrailState,
		WhyPoint5NotPass:         whyPoint5NotPass,
		SurfaceRefs:              surfaceRefs,
		EvidenceRefs:             evidenceRefs,
		Limitations:              limitations,
		ProjectionDisclaimer:     intelligenceCalibrationValBProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val B adds bounded behavioral baseline profiling, learning mode runtime discipline, threshold calibration, drift sensitivity scaling, criticality-aware weighting, baseline freshness/adoption review, explanations, and safety guardrails.",
			"Val B is fail-closed on active Val 0 and Val A proofs and keeps all behavioral calibration outputs advisory rather than canonical or mutating.",
			"Točka 5 remains not complete until later waves add feedback/federated tuning, simulation, final calibration gate, and integrated closure.",
		},
	}
}

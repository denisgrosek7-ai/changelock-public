package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	intelligenceCalibrationValDSimulationHarnessSchema  = "point5.intelligence_calibration.vald.simulation_harness.v1"
	intelligenceCalibrationValDScenarioLibrarySchema    = "point5.intelligence_calibration.vald.scenario_library.v1"
	intelligenceCalibrationValDMissedDetectionSchema    = "point5.intelligence_calibration.vald.missed_detection_analysis.v1"
	intelligenceCalibrationValDFPFNBalanceSchema        = "point5.intelligence_calibration.vald.fp_fn_balance.v1"
	intelligenceCalibrationValDConfidenceReviewSchema   = "point5.intelligence_calibration.vald.confidence_review.v1"
	intelligenceCalibrationValDVEXQualitySchema         = "point5.intelligence_calibration.vald.vex_quality.v1"
	intelligenceCalibrationValDReachabilitySchema       = "point5.intelligence_calibration.vald.reachability_quality.v1"
	intelligenceCalibrationValDBehavioralQualitySchema  = "point5.intelligence_calibration.vald.behavioral_quality.v1"
	intelligenceCalibrationValDFederatedQualitySchema   = "point5.intelligence_calibration.vald.federated_quality.v1"
	intelligenceCalibrationValDSimulationCoverageSchema = "point5.intelligence_calibration.vald.simulation_coverage.v1"
	intelligenceCalibrationValDQualityScoreboardSchema  = "point5.intelligence_calibration.vald.quality_scoreboard.v1"
	intelligenceCalibrationValDProofsSchema             = "point5.intelligence_calibration.vald.proofs.v1"
)

type intelligenceCalibrationValDSimulationHarnessResponse struct {
	SchemaVersion string                                                 `json:"schema_version"`
	GeneratedAt   time.Time                                              `json:"generated_at"`
	CurrentState  string                                                 `json:"current_state"`
	Model         operability.DefensiveSimulationScenarioHarnessContract `json:"model"`
	RouteRefs     []string                                               `json:"route_refs,omitempty"`
	Limitations   []string                                               `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDScenarioLibraryResponse struct {
	SchemaVersion string                                                  `json:"schema_version"`
	GeneratedAt   time.Time                                               `json:"generated_at"`
	CurrentState  string                                                  `json:"current_state"`
	Model         operability.AdversarialLowSignalScenarioLibraryContract `json:"model"`
	RouteRefs     []string                                                `json:"route_refs,omitempty"`
	Limitations   []string                                                `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDMissedDetectionResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         operability.MissedDetectionAnalysisContract `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDFPFNBalanceResponse struct {
	SchemaVersion string                                                      `json:"schema_version"`
	GeneratedAt   time.Time                                                   `json:"generated_at"`
	CurrentState  string                                                      `json:"current_state"`
	Model         operability.FalsePositiveFalseNegativeBalanceReviewContract `json:"model"`
	RouteRefs     []string                                                    `json:"route_refs,omitempty"`
	Limitations   []string                                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDConfidenceReviewResponse struct {
	SchemaVersion string                                          `json:"schema_version"`
	GeneratedAt   time.Time                                       `json:"generated_at"`
	CurrentState  string                                          `json:"current_state"`
	Model         operability.ConfidenceCalibrationReviewContract `json:"model"`
	RouteRefs     []string                                        `json:"route_refs,omitempty"`
	Limitations   []string                                        `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDVEXQualityResponse struct {
	SchemaVersion string                                        `json:"schema_version"`
	GeneratedAt   time.Time                                     `json:"generated_at"`
	CurrentState  string                                        `json:"current_state"`
	Model         operability.VEXCandidateQualityReviewContract `json:"model"`
	RouteRefs     []string                                      `json:"route_refs,omitempty"`
	Limitations   []string                                      `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDReachabilityQualityResponse struct {
	SchemaVersion string                                                   `json:"schema_version"`
	GeneratedAt   time.Time                                                `json:"generated_at"`
	CurrentState  string                                                   `json:"current_state"`
	Model         operability.ReachabilityCalibrationQualityReviewContract `json:"model"`
	RouteRefs     []string                                                 `json:"route_refs,omitempty"`
	Limitations   []string                                                 `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDBehavioralQualityResponse struct {
	SchemaVersion string                                                 `json:"schema_version"`
	GeneratedAt   time.Time                                              `json:"generated_at"`
	CurrentState  string                                                 `json:"current_state"`
	Model         operability.BehavioralCalibrationQualityReviewContract `json:"model"`
	RouteRefs     []string                                               `json:"route_refs,omitempty"`
	Limitations   []string                                               `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDFederatedQualityResponse struct {
	SchemaVersion string                                              `json:"schema_version"`
	GeneratedAt   time.Time                                           `json:"generated_at"`
	CurrentState  string                                              `json:"current_state"`
	Model         operability.FederatedWeightingQualityReviewContract `json:"model"`
	RouteRefs     []string                                            `json:"route_refs,omitempty"`
	Limitations   []string                                            `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDSimulationCoverageResponse struct {
	SchemaVersion string                                       `json:"schema_version"`
	GeneratedAt   time.Time                                    `json:"generated_at"`
	CurrentState  string                                       `json:"current_state"`
	Model         operability.SimulationCoverageReviewContract `json:"model"`
	RouteRefs     []string                                     `json:"route_refs,omitempty"`
	Limitations   []string                                     `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDQualityScoreboardResponse struct {
	SchemaVersion string                                            `json:"schema_version"`
	GeneratedAt   time.Time                                         `json:"generated_at"`
	CurrentState  string                                            `json:"current_state"`
	Model         operability.IntelligenceQualityScoreboardContract `json:"model"`
	RouteRefs     []string                                          `json:"route_refs,omitempty"`
	Limitations   []string                                          `json:"limitations,omitempty"`
}

type intelligenceCalibrationValDProofsResponse struct {
	SchemaVersion                         string    `json:"schema_version"`
	GeneratedAt                           time.Time `json:"generated_at"`
	CurrentState                          string    `json:"current_state"`
	Val0DependencyState                   string    `json:"val_0_dependency_state"`
	Val0FoundationState                   string    `json:"val_0_foundation_state"`
	ValADependencyState                   string    `json:"val_a_dependency_state"`
	ValAReachabilityVEXState              string    `json:"val_a_reachability_vex_state"`
	ValBDependencyState                   string    `json:"val_b_dependency_state"`
	ValBBehavioralLearningState           string    `json:"val_b_behavioral_learning_state"`
	ValCDependencyState                   string    `json:"val_c_dependency_state"`
	ValCFeedbackSuppressionFederatedState string    `json:"val_c_feedback_suppression_federated_state"`
	ValDState                             string    `json:"val_d_state"`
	Point5State                           string    `json:"point_5_state"`
	SimulationHarnessState                string    `json:"simulation_harness_state"`
	ScenarioLibraryState                  string    `json:"scenario_library_state"`
	MissedDetectionAnalysisState          string    `json:"missed_detection_analysis_state"`
	FPFNBalanceReviewState                string    `json:"fp_fn_balance_review_state"`
	ConfidenceCalibrationReviewState      string    `json:"confidence_calibration_review_state"`
	VEXQualityReviewState                 string    `json:"vex_quality_review_state"`
	ReachabilityQualityReviewState        string    `json:"reachability_quality_review_state"`
	BehavioralQualityReviewState          string    `json:"behavioral_quality_review_state"`
	FederatedQualityReviewState           string    `json:"federated_weighting_quality_review_state"`
	SimulationCoverageReviewState         string    `json:"simulation_coverage_review_state"`
	QualityScoreboardState                string    `json:"intelligence_quality_scoreboard_state"`
	WhyPoint5NotPass                      []string  `json:"why_point_5_not_pass,omitempty"`
	SurfaceRefs                           []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                          []string  `json:"evidence_refs,omitempty"`
	Limitations                           []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer                  string    `json:"projection_disclaimer"`
	IntegrationSummary                    []string  `json:"integration_summary,omitempty"`
}

func intelligenceCalibrationValDAllSurfaceRefs() []string {
	return []string{
		"/v1/intelligence/calibration/vald/simulation-harness",
		"/v1/intelligence/calibration/vald/scenario-library",
		"/v1/intelligence/calibration/vald/missed-detection-analysis",
		"/v1/intelligence/calibration/vald/fp-fn-balance",
		"/v1/intelligence/calibration/vald/confidence-review",
		"/v1/intelligence/calibration/vald/vex-quality",
		"/v1/intelligence/calibration/vald/reachability-quality",
		"/v1/intelligence/calibration/vald/behavioral-quality",
		"/v1/intelligence/calibration/vald/federated-quality",
		"/v1/intelligence/calibration/vald/simulation-coverage",
		"/v1/intelligence/calibration/vald/quality-scoreboard",
		"/v1/intelligence/calibration/vald/proofs",
	}
}

func intelligenceCalibrationValDEvidenceRefs() []string {
	return []string{
		"val0_proofs",
		"vala_proofs",
		"valb_proofs",
		"valc_proofs",
		"defensive_simulation_harness",
		"adversarial_scenario_library",
		"missed_detection_analysis",
		"false_positive_false_negative_balance_review",
		"confidence_calibration_review",
		"vex_candidate_quality_review",
		"reachability_calibration_quality_review",
		"behavioral_calibration_quality_review",
		"federated_weighting_quality_review",
		"simulation_coverage_review",
		"intelligence_quality_scoreboard",
		"evidence_spine",
	}
}

func intelligenceCalibrationValDProjectionDisclaimer() string {
	return "projection_only not_canonical_truth advisory_defensive_simulation_quality_measurement_gate"
}

func (s server) intelligenceCalibrationValDSimulationHarnessHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDSimulationHarness())
}

func (s server) intelligenceCalibrationValDScenarioLibraryHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDScenarioLibrary())
}

func (s server) intelligenceCalibrationValDMissedDetectionHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDMissedDetection())
}

func (s server) intelligenceCalibrationValDFPFNBalanceHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDFPFNBalance())
}

func (s server) intelligenceCalibrationValDConfidenceReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDConfidenceReview())
}

func (s server) intelligenceCalibrationValDVEXQualityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDVEXQuality())
}

func (s server) intelligenceCalibrationValDReachabilityQualityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDReachabilityQuality())
}

func (s server) intelligenceCalibrationValDBehavioralQualityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDBehavioralQuality())
}

func (s server) intelligenceCalibrationValDFederatedQualityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDFederatedQuality())
}

func (s server) intelligenceCalibrationValDSimulationCoverageHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDSimulationCoverage())
}

func (s server) intelligenceCalibrationValDQualityScoreboardHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDQualityScoreboard())
}

func (s server) intelligenceCalibrationValDProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValDProofs())
}

func buildIntelligenceCalibrationValDSimulationHarness() intelligenceCalibrationValDSimulationHarnessResponse {
	model := operability.IntelligenceCalibrationValDDefensiveSimulationHarnessContract()
	return intelligenceCalibrationValDSimulationHarnessResponse{
		SchemaVersion: intelligenceCalibrationValDSimulationHarnessSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDSimulationHarnessState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/simulation-harness"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValDScenarioLibrary() intelligenceCalibrationValDScenarioLibraryResponse {
	model := operability.IntelligenceCalibrationValDScenarioLibraryContract()
	return intelligenceCalibrationValDScenarioLibraryResponse{
		SchemaVersion: intelligenceCalibrationValDScenarioLibrarySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDScenarioLibraryState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/scenario-library"},
		Limitations:   model.CoverageLimitations,
	}
}

func buildIntelligenceCalibrationValDMissedDetection() intelligenceCalibrationValDMissedDetectionResponse {
	model := operability.IntelligenceCalibrationValDMissedDetectionAnalysisContract()
	return intelligenceCalibrationValDMissedDetectionResponse{
		SchemaVersion: intelligenceCalibrationValDMissedDetectionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDMissedDetectionState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/missed-detection-analysis"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValDFPFNBalance() intelligenceCalibrationValDFPFNBalanceResponse {
	model := operability.IntelligenceCalibrationValDFPFNBalanceReviewContract()
	return intelligenceCalibrationValDFPFNBalanceResponse{
		SchemaVersion: intelligenceCalibrationValDFPFNBalanceSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDFPFNBalanceState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/fp-fn-balance"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValDConfidenceReview() intelligenceCalibrationValDConfidenceReviewResponse {
	model := operability.IntelligenceCalibrationValDConfidenceCalibrationReviewContract()
	return intelligenceCalibrationValDConfidenceReviewResponse{
		SchemaVersion: intelligenceCalibrationValDConfidenceReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDConfidenceReviewState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/confidence-review"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValDVEXQuality() intelligenceCalibrationValDVEXQualityResponse {
	model := operability.IntelligenceCalibrationValDVEXQualityReviewContract()
	return intelligenceCalibrationValDVEXQualityResponse{
		SchemaVersion: intelligenceCalibrationValDVEXQualitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDVEXQualityState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/vex-quality"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValDReachabilityQuality() intelligenceCalibrationValDReachabilityQualityResponse {
	model := operability.IntelligenceCalibrationValDReachabilityQualityReviewContract()
	return intelligenceCalibrationValDReachabilityQualityResponse{
		SchemaVersion: intelligenceCalibrationValDReachabilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDReachabilityQualityState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/reachability-quality"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValDBehavioralQuality() intelligenceCalibrationValDBehavioralQualityResponse {
	model := operability.IntelligenceCalibrationValDBehavioralQualityReviewContract()
	return intelligenceCalibrationValDBehavioralQualityResponse{
		SchemaVersion: intelligenceCalibrationValDBehavioralQualitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDBehavioralQualityState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/behavioral-quality"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValDFederatedQuality() intelligenceCalibrationValDFederatedQualityResponse {
	model := operability.IntelligenceCalibrationValDFederatedQualityReviewContract()
	return intelligenceCalibrationValDFederatedQualityResponse{
		SchemaVersion: intelligenceCalibrationValDFederatedQualitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDFederatedQualityState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/federated-quality"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValDSimulationCoverage() intelligenceCalibrationValDSimulationCoverageResponse {
	model := operability.IntelligenceCalibrationValDSimulationCoverageReviewContract()
	return intelligenceCalibrationValDSimulationCoverageResponse{
		SchemaVersion: intelligenceCalibrationValDSimulationCoverageSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDSimulationCoverageState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/simulation-coverage"},
		Limitations:   model.CoverageLimitations,
	}
}

func buildIntelligenceCalibrationValDQualityScoreboard() intelligenceCalibrationValDQualityScoreboardResponse {
	model := operability.IntelligenceCalibrationValDQualityScoreboardContract()
	return intelligenceCalibrationValDQualityScoreboardResponse{
		SchemaVersion: intelligenceCalibrationValDQualityScoreboardSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValDQualityScoreboardState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vald/quality-scoreboard"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValDProofs() intelligenceCalibrationValDProofsResponse {
	val0 := buildIntelligenceCalibrationVal0Proofs()
	valA := buildIntelligenceCalibrationValAProofs()
	valB := buildIntelligenceCalibrationValBProofs()
	valC := buildIntelligenceCalibrationValCProofs()

	simulationHarness := operability.IntelligenceCalibrationValDDefensiveSimulationHarnessContract()
	scenarioLibrary := operability.IntelligenceCalibrationValDScenarioLibraryContract()
	missedDetection := operability.IntelligenceCalibrationValDMissedDetectionAnalysisContract()
	fpfnBalance := operability.IntelligenceCalibrationValDFPFNBalanceReviewContract()
	confidenceReview := operability.IntelligenceCalibrationValDConfidenceCalibrationReviewContract()
	vexQuality := operability.IntelligenceCalibrationValDVEXQualityReviewContract()
	reachabilityQuality := operability.IntelligenceCalibrationValDReachabilityQualityReviewContract()
	behavioralQuality := operability.IntelligenceCalibrationValDBehavioralQualityReviewContract()
	federatedQuality := operability.IntelligenceCalibrationValDFederatedQualityReviewContract()
	simulationCoverage := operability.IntelligenceCalibrationValDSimulationCoverageReviewContract()
	qualityScoreboard := operability.IntelligenceCalibrationValDQualityScoreboardContract()

	simulationHarnessState := operability.EvaluateIntelligenceCalibrationValDSimulationHarnessState(simulationHarness)
	scenarioLibraryState := operability.EvaluateIntelligenceCalibrationValDScenarioLibraryState(scenarioLibrary)
	missedDetectionState := operability.EvaluateIntelligenceCalibrationValDMissedDetectionState(missedDetection)
	fpfnBalanceState := operability.EvaluateIntelligenceCalibrationValDFPFNBalanceState(fpfnBalance)
	confidenceReviewState := operability.EvaluateIntelligenceCalibrationValDConfidenceReviewState(confidenceReview)
	vexQualityState := operability.EvaluateIntelligenceCalibrationValDVEXQualityState(vexQuality)
	reachabilityQualityState := operability.EvaluateIntelligenceCalibrationValDReachabilityQualityState(reachabilityQuality)
	behavioralQualityState := operability.EvaluateIntelligenceCalibrationValDBehavioralQualityState(behavioralQuality)
	federatedQualityState := operability.EvaluateIntelligenceCalibrationValDFederatedQualityState(federatedQuality)
	simulationCoverageState := operability.EvaluateIntelligenceCalibrationValDSimulationCoverageState(simulationCoverage)
	qualityScoreboardState := operability.EvaluateIntelligenceCalibrationValDQualityScoreboardState(qualityScoreboard)

	surfaceRefs := intelligenceCalibrationValDAllSurfaceRefs()
	evidenceRefs := intelligenceCalibrationValDEvidenceRefs()
	limitations := []string{
		"Val D proves only Defensive Simulation & Quality Measurement Gate readiness and does not claim complete intelligence calibration or integrated closure.",
		"Simulation harnesses, adversarial libraries, missed-detection analyses, reviews, coverage summaries, and scoreboards remain scoped advisory proofs over canonical evidence and cannot mutate enforcement, priority, suppression, VEX publication, or governance.",
	}
	whyPoint5NotPass := []string{
		"Točka 5 remains not complete because Val E integrated closure is still required before any final calibration pass can be claimed.",
		"Val D adds defensive simulation and quality measurement review only; it does not create automatic suppression, active mutations, public claims, or final authority.",
	}

	valDState := operability.EvaluateIntelligenceCalibrationValDState(
		val0.CurrentState,
		val0.Val0State,
		valA.CurrentState,
		valA.ValAState,
		valB.CurrentState,
		valB.ValBState,
		valC.CurrentState,
		valC.ValCState,
		simulationHarnessState,
		scenarioLibraryState,
		missedDetectionState,
		fpfnBalanceState,
		confidenceReviewState,
		vexQualityState,
		reachabilityQualityState,
		behavioralQualityState,
		federatedQualityState,
		simulationCoverageState,
		qualityScoreboardState,
	)
	currentState := operability.EvaluateIntelligenceCalibrationValDProofsState(
		val0.CurrentState,
		val0.Val0State,
		valA.CurrentState,
		valA.ValAState,
		valB.CurrentState,
		valB.ValBState,
		valC.CurrentState,
		valC.ValCState,
		simulationHarnessState,
		scenarioLibraryState,
		missedDetectionState,
		fpfnBalanceState,
		confidenceReviewState,
		vexQualityState,
		reachabilityQualityState,
		behavioralQualityState,
		federatedQualityState,
		simulationCoverageState,
		qualityScoreboardState,
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint5NotPass,
		intelligenceCalibrationValDProjectionDisclaimer(),
	)

	return intelligenceCalibrationValDProofsResponse{
		SchemaVersion:                         intelligenceCalibrationValDProofsSchema,
		GeneratedAt:                           publicSampleTime(),
		CurrentState:                          currentState,
		Val0DependencyState:                   val0.CurrentState,
		Val0FoundationState:                   val0.Val0State,
		ValADependencyState:                   valA.CurrentState,
		ValAReachabilityVEXState:              valA.ValAState,
		ValBDependencyState:                   valB.CurrentState,
		ValBBehavioralLearningState:           valB.ValBState,
		ValCDependencyState:                   valC.CurrentState,
		ValCFeedbackSuppressionFederatedState: valC.ValCState,
		ValDState:                             valDState,
		Point5State:                           operability.IntelligenceCalibrationPoint5StateNotComplete,
		SimulationHarnessState:                simulationHarnessState,
		ScenarioLibraryState:                  scenarioLibraryState,
		MissedDetectionAnalysisState:          missedDetectionState,
		FPFNBalanceReviewState:                fpfnBalanceState,
		ConfidenceCalibrationReviewState:      confidenceReviewState,
		VEXQualityReviewState:                 vexQualityState,
		ReachabilityQualityReviewState:        reachabilityQualityState,
		BehavioralQualityReviewState:          behavioralQualityState,
		FederatedQualityReviewState:           federatedQualityState,
		SimulationCoverageReviewState:         simulationCoverageState,
		QualityScoreboardState:                qualityScoreboardState,
		WhyPoint5NotPass:                      whyPoint5NotPass,
		SurfaceRefs:                           surfaceRefs,
		EvidenceRefs:                          evidenceRefs,
		Limitations:                           limitations,
		ProjectionDisclaimer:                  intelligenceCalibrationValDProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val D adds a replayable defensive simulation harness, adversarial scenario library, missed-detection analysis, FP/FN balance review, confidence/VEX/reachability/behavioral/federated quality reviews, coverage review, and a bounded intelligence quality scoreboard.",
			"Val D is fail-closed on active Val 0, Val A, Val B, and Val C proofs and keeps all simulation and quality outputs advisory rather than canonical or mutating.",
			"Točka 5 remains not complete until Val E integrated closure is implemented.",
		},
	}
}

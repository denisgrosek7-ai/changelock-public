package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	intelligenceCalibrationVal0DatasetSchema      = "point5.intelligence_calibration.val0.datasets.v1"
	intelligenceCalibrationVal0ConfidenceSchema   = "point5.intelligence_calibration.val0.confidence_model.v1"
	intelligenceCalibrationVal0LifecycleSchema    = "point5.intelligence_calibration.val0.output_lifecycle.v1"
	intelligenceCalibrationVal0ReachabilitySchema = "point5.intelligence_calibration.val0.reachability_taxonomy.v1"
	intelligenceCalibrationVal0VEXSchema          = "point5.intelligence_calibration.val0.vex_candidates.v1"
	intelligenceCalibrationVal0FeedbackSchema     = "point5.intelligence_calibration.val0.feedback.v1"
	intelligenceCalibrationVal0LearningModeSchema = "point5.intelligence_calibration.val0.learning_mode.v1"
	intelligenceCalibrationVal0SuppressionSchema  = "point5.intelligence_calibration.val0.suppression_safety.v1"
	intelligenceCalibrationVal0FederatedSchema    = "point5.intelligence_calibration.val0.federated_boundary.v1"
	intelligenceCalibrationVal0ProvenanceSchema   = "point5.intelligence_calibration.val0.provenance.v1"
	intelligenceCalibrationVal0FreshnessSchema    = "point5.intelligence_calibration.val0.freshness_expiry.v1"
	intelligenceCalibrationVal0MetricsSchema      = "point5.intelligence_calibration.val0.metrics.v1"
	intelligenceCalibrationVal0FPFNSchema         = "point5.intelligence_calibration.val0.fp_fn_discipline.v1"
	intelligenceCalibrationVal0RollbackSchema     = "point5.intelligence_calibration.val0.rollback.v1"
	intelligenceCalibrationVal0ProofsSchema       = "point5.intelligence_calibration.val0.proofs.v1"
)

type intelligenceCalibrationVal0DatasetResponse struct {
	SchemaVersion string                                 `json:"schema_version"`
	GeneratedAt   time.Time                              `json:"generated_at"`
	CurrentState  string                                 `json:"current_state"`
	Model         operability.CalibrationDatasetContract `json:"model"`
	RouteRefs     []string                               `json:"route_refs,omitempty"`
	Limitations   []string                               `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0ConfidenceResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         operability.ConfidenceEvidenceClassContract `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0LifecycleResponse struct {
	SchemaVersion string                                          `json:"schema_version"`
	GeneratedAt   time.Time                                       `json:"generated_at"`
	CurrentState  string                                          `json:"current_state"`
	Model         operability.IntelligenceOutputLifecycleContract `json:"model"`
	RouteRefs     []string                                        `json:"route_refs,omitempty"`
	Limitations   []string                                        `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0ReachabilityResponse struct {
	SchemaVersion string                                   `json:"schema_version"`
	GeneratedAt   time.Time                                `json:"generated_at"`
	CurrentState  string                                   `json:"current_state"`
	Model         operability.ReachabilityTaxonomyContract `json:"model"`
	RouteRefs     []string                                 `json:"route_refs,omitempty"`
	Limitations   []string                                 `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0VEXResponse struct {
	SchemaVersion string                                     `json:"schema_version"`
	GeneratedAt   time.Time                                  `json:"generated_at"`
	CurrentState  string                                     `json:"current_state"`
	Model         operability.VEXCandidateGovernanceContract `json:"model"`
	RouteRefs     []string                                   `json:"route_refs,omitempty"`
	Limitations   []string                                   `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0FeedbackResponse struct {
	SchemaVersion string                                     `json:"schema_version"`
	GeneratedAt   time.Time                                  `json:"generated_at"`
	CurrentState  string                                     `json:"current_state"`
	Model         operability.FeedbackClassificationContract `json:"model"`
	RouteRefs     []string                                   `json:"route_refs,omitempty"`
	Limitations   []string                                   `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0LearningModeResponse struct {
	SchemaVersion string                                    `json:"schema_version"`
	GeneratedAt   time.Time                                 `json:"generated_at"`
	CurrentState  string                                    `json:"current_state"`
	Model         operability.LearningModeGuardrailContract `json:"model"`
	RouteRefs     []string                                  `json:"route_refs,omitempty"`
	Limitations   []string                                  `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0SuppressionResponse struct {
	SchemaVersion string                                `json:"schema_version"`
	GeneratedAt   time.Time                             `json:"generated_at"`
	CurrentState  string                                `json:"current_state"`
	Model         operability.SuppressionSafetyContract `json:"model"`
	RouteRefs     []string                              `json:"route_refs,omitempty"`
	Limitations   []string                              `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0FederatedResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         operability.FederatedSignalBoundaryContract `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0ProvenanceResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Model         operability.CalibrationProvenanceChainContract `json:"model"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0FreshnessResponse struct {
	SchemaVersion string                                          `json:"schema_version"`
	GeneratedAt   time.Time                                       `json:"generated_at"`
	CurrentState  string                                          `json:"current_state"`
	Model         operability.IntelligenceFreshnessExpiryContract `json:"model"`
	RouteRefs     []string                                        `json:"route_refs,omitempty"`
	Limitations   []string                                        `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0MetricsResponse struct {
	SchemaVersion string                                           `json:"schema_version"`
	GeneratedAt   time.Time                                        `json:"generated_at"`
	CurrentState  string                                           `json:"current_state"`
	Model         operability.CalibrationMetricsDefinitionContract `json:"model"`
	RouteRefs     []string                                         `json:"route_refs,omitempty"`
	Limitations   []string                                         `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0FPFNResponse struct {
	SchemaVersion string                                              `json:"schema_version"`
	GeneratedAt   time.Time                                           `json:"generated_at"`
	CurrentState  string                                              `json:"current_state"`
	Model         operability.FalsePositiveNegativeDisciplineContract `json:"model"`
	RouteRefs     []string                                            `json:"route_refs,omitempty"`
	Limitations   []string                                            `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0RollbackResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         operability.CalibrationRollbackUndoContract `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationVal0ProofsResponse struct {
	SchemaVersion          string    `json:"schema_version"`
	GeneratedAt            time.Time `json:"generated_at"`
	CurrentState           string    `json:"current_state"`
	Val0State              string    `json:"val_0_state"`
	Point5State            string    `json:"point_5_state"`
	DatasetState           string    `json:"dataset_discipline_state"`
	ConfidenceState        string    `json:"confidence_evidence_class_state"`
	LifecycleState         string    `json:"intelligence_lifecycle_state"`
	ReachabilityState      string    `json:"reachability_taxonomy_state"`
	VEXState               string    `json:"vex_candidate_lifecycle_state"`
	FeedbackState          string    `json:"feedback_classification_state"`
	LearningModeState      string    `json:"learning_mode_guardrail_state"`
	SuppressionState       string    `json:"suppression_safety_state"`
	FederatedBoundaryState string    `json:"federated_boundary_state"`
	ProvenanceState        string    `json:"provenance_chain_state"`
	FreshnessState         string    `json:"freshness_expiry_state"`
	MetricsState           string    `json:"metrics_definition_state"`
	FPFNState              string    `json:"fp_fn_discipline_state"`
	RollbackState          string    `json:"rollback_undo_state"`
	WhyPoint5NotPass       []string  `json:"why_point_5_not_pass,omitempty"`
	SurfaceRefs            []string  `json:"surface_refs,omitempty"`
	EvidenceRefs           []string  `json:"evidence_refs,omitempty"`
	Limitations            []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer   string    `json:"projection_disclaimer"`
	IntegrationSummary     []string  `json:"integration_summary,omitempty"`
}

func intelligenceCalibrationVal0AllSurfaceRefs() []string {
	return []string{
		"/v1/intelligence/calibration/val0/datasets",
		"/v1/intelligence/calibration/val0/confidence-model",
		"/v1/intelligence/calibration/val0/output-lifecycle",
		"/v1/intelligence/calibration/val0/reachability-taxonomy",
		"/v1/intelligence/calibration/val0/vex-candidates",
		"/v1/intelligence/calibration/val0/feedback",
		"/v1/intelligence/calibration/val0/learning-mode",
		"/v1/intelligence/calibration/val0/suppression-safety",
		"/v1/intelligence/calibration/val0/federated-boundary",
		"/v1/intelligence/calibration/val0/provenance",
		"/v1/intelligence/calibration/val0/freshness-expiry",
		"/v1/intelligence/calibration/val0/metrics",
		"/v1/intelligence/calibration/val0/fp-fn-discipline",
		"/v1/intelligence/calibration/val0/rollback",
		"/v1/intelligence/calibration/val0/proofs",
	}
}

func intelligenceCalibrationVal0EvidenceRefs() []string {
	return []string{
		"dataset_manifest",
		"confidence_reason_trace",
		"lifecycle_state_contract",
		"reachability_taxonomy_contract",
		"vex_candidate_contract",
		"feedback_review_contract",
		"learning_mode_guardrails",
		"suppression_safety_contract",
		"federated_boundary_contract",
		"calibration_provenance_contract",
		"freshness_expiry_contract",
		"metrics_definition_contract",
		"fp_fn_discipline_contract",
		"rollback_undo_contract",
		"evidence_spine",
	}
}

func intelligenceCalibrationVal0ProjectionDisclaimer() string {
	return "projection_only not_canonical_truth advisory_intelligence_calibration_foundation"
}

func (s server) intelligenceCalibrationVal0DatasetHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Dataset())
}

func (s server) intelligenceCalibrationVal0ConfidenceHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Confidence())
}

func (s server) intelligenceCalibrationVal0LifecycleHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Lifecycle())
}

func (s server) intelligenceCalibrationVal0ReachabilityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Reachability())
}

func (s server) intelligenceCalibrationVal0VEXHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0VEX())
}

func (s server) intelligenceCalibrationVal0FeedbackHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Feedback())
}

func (s server) intelligenceCalibrationVal0LearningModeHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0LearningMode())
}

func (s server) intelligenceCalibrationVal0SuppressionHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Suppression())
}

func (s server) intelligenceCalibrationVal0FederatedHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0FederatedBoundary())
}

func (s server) intelligenceCalibrationVal0ProvenanceHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Provenance())
}

func (s server) intelligenceCalibrationVal0FreshnessHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Freshness())
}

func (s server) intelligenceCalibrationVal0MetricsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Metrics())
}

func (s server) intelligenceCalibrationVal0FPFNHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0FPFN())
}

func (s server) intelligenceCalibrationVal0RollbackHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Rollback())
}

func (s server) intelligenceCalibrationVal0ProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationVal0Proofs())
}

func buildIntelligenceCalibrationVal0Dataset() intelligenceCalibrationVal0DatasetResponse {
	model := operability.IntelligenceCalibrationVal0DatasetContract()
	return intelligenceCalibrationVal0DatasetResponse{
		SchemaVersion: intelligenceCalibrationVal0DatasetSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0DatasetState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/metrics",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildIntelligenceCalibrationVal0Confidence() intelligenceCalibrationVal0ConfidenceResponse {
	model := operability.IntelligenceCalibrationVal0ConfidenceContract()
	return intelligenceCalibrationVal0ConfidenceResponse{
		SchemaVersion: intelligenceCalibrationVal0ConfidenceSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0ConfidenceState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/freshness-expiry",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationVal0Lifecycle() intelligenceCalibrationVal0LifecycleResponse {
	model := operability.IntelligenceCalibrationVal0OutputLifecycleContract()
	return intelligenceCalibrationVal0LifecycleResponse{
		SchemaVersion: intelligenceCalibrationVal0LifecycleSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0LifecycleState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/vex-candidates",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildIntelligenceCalibrationVal0Reachability() intelligenceCalibrationVal0ReachabilityResponse {
	model := operability.IntelligenceCalibrationVal0ReachabilityTaxonomyContract()
	return intelligenceCalibrationVal0ReachabilityResponse{
		SchemaVersion: intelligenceCalibrationVal0ReachabilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0ReachabilityState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/vex-candidates",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildIntelligenceCalibrationVal0VEX() intelligenceCalibrationVal0VEXResponse {
	model := operability.IntelligenceCalibrationVal0VEXCandidateContract()
	return intelligenceCalibrationVal0VEXResponse{
		SchemaVersion: intelligenceCalibrationVal0VEXSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0VEXState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/reachability-taxonomy",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildIntelligenceCalibrationVal0Feedback() intelligenceCalibrationVal0FeedbackResponse {
	model := operability.IntelligenceCalibrationVal0FeedbackContract()
	return intelligenceCalibrationVal0FeedbackResponse{
		SchemaVersion: intelligenceCalibrationVal0FeedbackSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0FeedbackState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/provenance",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildIntelligenceCalibrationVal0LearningMode() intelligenceCalibrationVal0LearningModeResponse {
	model := operability.IntelligenceCalibrationVal0LearningModeContract()
	return intelligenceCalibrationVal0LearningModeResponse{
		SchemaVersion: intelligenceCalibrationVal0LearningModeSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0LearningModeState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/fp-fn-discipline",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationVal0Suppression() intelligenceCalibrationVal0SuppressionResponse {
	model := operability.IntelligenceCalibrationVal0SuppressionSafetyContract()
	return intelligenceCalibrationVal0SuppressionResponse{
		SchemaVersion: intelligenceCalibrationVal0SuppressionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0SuppressionState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/fp-fn-discipline",
			"/v1/intelligence/calibration/val0/rollback",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationVal0FederatedBoundary() intelligenceCalibrationVal0FederatedResponse {
	model := operability.IntelligenceCalibrationVal0FederatedBoundaryContract()
	return intelligenceCalibrationVal0FederatedResponse{
		SchemaVersion: intelligenceCalibrationVal0FederatedSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0FederatedBoundaryState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/confidence-model",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationVal0Provenance() intelligenceCalibrationVal0ProvenanceResponse {
	model := operability.IntelligenceCalibrationVal0ProvenanceContract()
	return intelligenceCalibrationVal0ProvenanceResponse{
		SchemaVersion: intelligenceCalibrationVal0ProvenanceSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0ProvenanceState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/feedback",
			"/v1/intelligence/calibration/val0/rollback",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: []string{"Calibration provenance is traceable and still bounded by later review and governance."},
	}
}

func buildIntelligenceCalibrationVal0Freshness() intelligenceCalibrationVal0FreshnessResponse {
	model := operability.IntelligenceCalibrationVal0FreshnessContract()
	return intelligenceCalibrationVal0FreshnessResponse{
		SchemaVersion: intelligenceCalibrationVal0FreshnessSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0FreshnessState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/confidence-model",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildIntelligenceCalibrationVal0Metrics() intelligenceCalibrationVal0MetricsResponse {
	model := operability.IntelligenceCalibrationVal0MetricsContract()
	return intelligenceCalibrationVal0MetricsResponse{
		SchemaVersion: intelligenceCalibrationVal0MetricsSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0MetricsState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/datasets",
			"/v1/intelligence/calibration/val0/rollback",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: model.Limitations,
	}
}

func buildIntelligenceCalibrationVal0FPFN() intelligenceCalibrationVal0FPFNResponse {
	model := operability.IntelligenceCalibrationVal0FPFNContract()
	return intelligenceCalibrationVal0FPFNResponse{
		SchemaVersion: intelligenceCalibrationVal0FPFNSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0FPFNState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/suppression-safety",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationVal0Rollback() intelligenceCalibrationVal0RollbackResponse {
	model := operability.IntelligenceCalibrationVal0RollbackContract()
	return intelligenceCalibrationVal0RollbackResponse{
		SchemaVersion: intelligenceCalibrationVal0RollbackSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationVal0RollbackState(model),
		Model:         model,
		RouteRefs: []string{
			"/v1/intelligence/calibration/val0/metrics",
			"/v1/intelligence/calibration/val0/provenance",
			"/v1/intelligence/calibration/val0/proofs",
		},
		Limitations: []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationVal0Proofs() intelligenceCalibrationVal0ProofsResponse {
	dataset := operability.IntelligenceCalibrationVal0DatasetContract()
	confidence := operability.IntelligenceCalibrationVal0ConfidenceContract()
	lifecycle := operability.IntelligenceCalibrationVal0OutputLifecycleContract()
	reachability := operability.IntelligenceCalibrationVal0ReachabilityTaxonomyContract()
	vex := operability.IntelligenceCalibrationVal0VEXCandidateContract()
	feedback := operability.IntelligenceCalibrationVal0FeedbackContract()
	learningMode := operability.IntelligenceCalibrationVal0LearningModeContract()
	suppression := operability.IntelligenceCalibrationVal0SuppressionSafetyContract()
	federated := operability.IntelligenceCalibrationVal0FederatedBoundaryContract()
	provenance := operability.IntelligenceCalibrationVal0ProvenanceContract()
	freshness := operability.IntelligenceCalibrationVal0FreshnessContract()
	metrics := operability.IntelligenceCalibrationVal0MetricsContract()
	fpfn := operability.IntelligenceCalibrationVal0FPFNContract()
	rollback := operability.IntelligenceCalibrationVal0RollbackContract()

	datasetState := operability.EvaluateIntelligenceCalibrationVal0DatasetState(dataset)
	confidenceState := operability.EvaluateIntelligenceCalibrationVal0ConfidenceState(confidence)
	lifecycleState := operability.EvaluateIntelligenceCalibrationVal0LifecycleState(lifecycle)
	reachabilityState := operability.EvaluateIntelligenceCalibrationVal0ReachabilityState(reachability)
	vexState := operability.EvaluateIntelligenceCalibrationVal0VEXState(vex)
	feedbackState := operability.EvaluateIntelligenceCalibrationVal0FeedbackState(feedback)
	learningModeState := operability.EvaluateIntelligenceCalibrationVal0LearningModeState(learningMode)
	suppressionState := operability.EvaluateIntelligenceCalibrationVal0SuppressionState(suppression)
	federatedState := operability.EvaluateIntelligenceCalibrationVal0FederatedBoundaryState(federated)
	provenanceState := operability.EvaluateIntelligenceCalibrationVal0ProvenanceState(provenance)
	freshnessState := operability.EvaluateIntelligenceCalibrationVal0FreshnessState(freshness)
	metricsState := operability.EvaluateIntelligenceCalibrationVal0MetricsState(metrics)
	fpfnState := operability.EvaluateIntelligenceCalibrationVal0FPFNState(fpfn)
	rollbackState := operability.EvaluateIntelligenceCalibrationVal0RollbackState(rollback)

	surfaceRefs := intelligenceCalibrationVal0AllSurfaceRefs()
	evidenceRefs := intelligenceCalibrationVal0EvidenceRefs()
	limitations := []string{
		"Val 0 proves only calibration discipline foundation and does not claim Točka 5 calibrated intelligence completion.",
		"Intelligence remains advisory projection_only not_canonical_truth and cannot approve, mutate, suppress critical evidence, or replace governance.",
	}
	whyPoint5NotPass := []string{
		"Točka 5 remains not complete because later waves still need the actual reachability, behavioral baseline, feedback tuning, simulation, final gate, and integrated closure layers.",
		"Val 0 defines fail-closed semantic foundations only and does not introduce learning, suppression, federated propagation, or calibration mutation authority.",
	}

	val0State := operability.EvaluateIntelligenceCalibrationVal0State(
		datasetState,
		confidenceState,
		lifecycleState,
		reachabilityState,
		vexState,
		feedbackState,
		learningModeState,
		suppressionState,
		federatedState,
		provenanceState,
		freshnessState,
		metricsState,
		fpfnState,
		rollbackState,
	)
	currentState := operability.EvaluateIntelligenceCalibrationVal0ProofsState(
		datasetState,
		confidenceState,
		lifecycleState,
		reachabilityState,
		vexState,
		feedbackState,
		learningModeState,
		suppressionState,
		federatedState,
		provenanceState,
		freshnessState,
		metricsState,
		fpfnState,
		rollbackState,
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint5NotPass,
		intelligenceCalibrationVal0ProjectionDisclaimer(),
	)

	return intelligenceCalibrationVal0ProofsResponse{
		SchemaVersion:          intelligenceCalibrationVal0ProofsSchema,
		GeneratedAt:            publicSampleTime(),
		CurrentState:           currentState,
		Val0State:              val0State,
		Point5State:            operability.IntelligenceCalibrationPoint5StateNotComplete,
		DatasetState:           datasetState,
		ConfidenceState:        confidenceState,
		LifecycleState:         lifecycleState,
		ReachabilityState:      reachabilityState,
		VEXState:               vexState,
		FeedbackState:          feedbackState,
		LearningModeState:      learningModeState,
		SuppressionState:       suppressionState,
		FederatedBoundaryState: federatedState,
		ProvenanceState:        provenanceState,
		FreshnessState:         freshnessState,
		MetricsState:           metricsState,
		FPFNState:              fpfnState,
		RollbackState:          rollbackState,
		WhyPoint5NotPass:       whyPoint5NotPass,
		SurfaceRefs:            surfaceRefs,
		EvidenceRefs:           evidenceRefs,
		Limitations:            limitations,
		ProjectionDisclaimer:   intelligenceCalibrationVal0ProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val 0 establishes fail-closed calibration dataset, confidence, lifecycle, reachability, VEX candidate, feedback, suppression, federated, freshness, metrics, FP/FN, and rollback contracts.",
			"All intelligence outputs remain advisory projections over canonical evidence and do not gain mutation, approval, or suppression authority in Val 0.",
			"Točka 5 remains not complete until later waves add bounded engines, review gates, and integrated closure.",
		},
	}
}

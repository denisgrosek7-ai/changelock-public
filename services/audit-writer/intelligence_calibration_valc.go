package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	intelligenceCalibrationValCFeedbackIntakeSchema      = "point5.intelligence_calibration.valc.feedback_intake.v1"
	intelligenceCalibrationValCReviewCockpitSchema       = "point5.intelligence_calibration.valc.feedback_review_cockpit.v1"
	intelligenceCalibrationValCTuningProposalSchema      = "point5.intelligence_calibration.valc.tuning_proposals.v1"
	intelligenceCalibrationValCSuppressionSafetySchema   = "point5.intelligence_calibration.valc.suppression_safety.v1"
	intelligenceCalibrationValCSuppressionRollbackSchema = "point5.intelligence_calibration.valc.suppression_rollback.v1"
	intelligenceCalibrationValCLocalChangeReviewSchema   = "point5.intelligence_calibration.valc.local_change_review.v1"
	intelligenceCalibrationValCFederatedWeightingSchema  = "point5.intelligence_calibration.valc.federated_weighting.v1"
	intelligenceCalibrationValCSimilarityGatingSchema    = "point5.intelligence_calibration.valc.similarity_gating.v1"
	intelligenceCalibrationValCLocalOverrideSchema       = "point5.intelligence_calibration.valc.local_override.v1"
	intelligenceCalibrationValCPropagationPolicySchema   = "point5.intelligence_calibration.valc.propagation_policy.v1"
	intelligenceCalibrationValCExplanationSchema         = "point5.intelligence_calibration.valc.explanations.v1"
	intelligenceCalibrationValCProofsSchema              = "point5.intelligence_calibration.valc.proofs.v1"
)

type intelligenceCalibrationValCFeedbackIntakeResponse struct {
	SchemaVersion string                                       `json:"schema_version"`
	GeneratedAt   time.Time                                    `json:"generated_at"`
	CurrentState  string                                       `json:"current_state"`
	Model         operability.StructuredFeedbackIntakeContract `json:"model"`
	RouteRefs     []string                                     `json:"route_refs,omitempty"`
	Limitations   []string                                     `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCReviewCockpitResponse struct {
	SchemaVersion string                                    `json:"schema_version"`
	GeneratedAt   time.Time                                 `json:"generated_at"`
	CurrentState  string                                    `json:"current_state"`
	Model         operability.FeedbackReviewCockpitContract `json:"model"`
	RouteRefs     []string                                  `json:"route_refs,omitempty"`
	Limitations   []string                                  `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCTuningProposalResponse struct {
	SchemaVersion string                             `json:"schema_version"`
	GeneratedAt   time.Time                          `json:"generated_at"`
	CurrentState  string                             `json:"current_state"`
	Model         operability.TuningProposalContract `json:"model"`
	RouteRefs     []string                           `json:"route_refs,omitempty"`
	Limitations   []string                           `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCSuppressionSafetyResponse struct {
	SchemaVersion string                                           `json:"schema_version"`
	GeneratedAt   time.Time                                        `json:"generated_at"`
	CurrentState  string                                           `json:"current_state"`
	Model         operability.SuppressionSafetyApplicationContract `json:"model"`
	RouteRefs     []string                                         `json:"route_refs,omitempty"`
	Limitations   []string                                         `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCSuppressionRollbackResponse struct {
	SchemaVersion string                                  `json:"schema_version"`
	GeneratedAt   time.Time                               `json:"generated_at"`
	CurrentState  string                                  `json:"current_state"`
	Model         operability.SuppressionRollbackContract `json:"model"`
	RouteRefs     []string                                `json:"route_refs,omitempty"`
	Limitations   []string                                `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCLocalChangeReviewResponse struct {
	SchemaVersion string                                           `json:"schema_version"`
	GeneratedAt   time.Time                                        `json:"generated_at"`
	CurrentState  string                                           `json:"current_state"`
	Model         operability.LocalCalibrationChangeReviewContract `json:"model"`
	RouteRefs     []string                                         `json:"route_refs,omitempty"`
	Limitations   []string                                         `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCFederatedWeightingResponse struct {
	SchemaVersion string                                       `json:"schema_version"`
	GeneratedAt   time.Time                                    `json:"generated_at"`
	CurrentState  string                                       `json:"current_state"`
	Model         operability.FederatedSignalWeightingContract `json:"model"`
	RouteRefs     []string                                     `json:"route_refs,omitempty"`
	Limitations   []string                                     `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCSimilarityGatingResponse struct {
	SchemaVersion string                                          `json:"schema_version"`
	GeneratedAt   time.Time                                       `json:"generated_at"`
	CurrentState  string                                          `json:"current_state"`
	Model         operability.EnvironmentSimilarityGatingContract `json:"model"`
	RouteRefs     []string                                        `json:"route_refs,omitempty"`
	Limitations   []string                                        `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCLocalOverrideResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         operability.LocalOverrideDisciplineContract `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCPropagationPolicyResponse struct {
	SchemaVersion string                                       `json:"schema_version"`
	GeneratedAt   time.Time                                    `json:"generated_at"`
	CurrentState  string                                       `json:"current_state"`
	Model         operability.BoundedPropagationPolicyContract `json:"model"`
	RouteRefs     []string                                     `json:"route_refs,omitempty"`
	Limitations   []string                                     `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCExplanationResponse struct {
	SchemaVersion string                                           `json:"schema_version"`
	GeneratedAt   time.Time                                        `json:"generated_at"`
	CurrentState  string                                           `json:"current_state"`
	Model         operability.FeedbackFederatedExplanationContract `json:"model"`
	RouteRefs     []string                                         `json:"route_refs,omitempty"`
	Limitations   []string                                         `json:"limitations,omitempty"`
}

type intelligenceCalibrationValCProofsResponse struct {
	SchemaVersion               string    `json:"schema_version"`
	GeneratedAt                 time.Time `json:"generated_at"`
	CurrentState                string    `json:"current_state"`
	Val0DependencyState         string    `json:"val_0_dependency_state"`
	Val0FoundationState         string    `json:"val_0_foundation_state"`
	ValADependencyState         string    `json:"val_a_dependency_state"`
	ValAReachabilityVEXState    string    `json:"val_a_reachability_vex_state"`
	ValBDependencyState         string    `json:"val_b_dependency_state"`
	ValBBehavioralLearningState string    `json:"val_b_behavioral_learning_state"`
	ValCState                   string    `json:"val_c_state"`
	Point5State                 string    `json:"point_5_state"`
	FeedbackIntakeState         string    `json:"structured_feedback_intake_state"`
	ReviewCockpitState          string    `json:"feedback_review_cockpit_state"`
	TuningProposalState         string    `json:"tuning_proposal_state"`
	SuppressionSafetyState      string    `json:"suppression_safety_application_state"`
	SuppressionRollbackState    string    `json:"suppression_rollback_state"`
	LocalChangeReviewState      string    `json:"local_calibration_change_review_state"`
	FederatedWeightingState     string    `json:"federated_signal_weighting_state"`
	SimilarityGatingState       string    `json:"environment_similarity_gating_state"`
	LocalOverrideState          string    `json:"local_override_discipline_state"`
	PropagationPolicyState      string    `json:"bounded_propagation_policy_state"`
	ExplanationState            string    `json:"feedback_federated_explanation_state"`
	WhyPoint5NotPass            []string  `json:"why_point_5_not_pass,omitempty"`
	SurfaceRefs                 []string  `json:"surface_refs,omitempty"`
	EvidenceRefs                []string  `json:"evidence_refs,omitempty"`
	Limitations                 []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer        string    `json:"projection_disclaimer"`
	IntegrationSummary          []string  `json:"integration_summary,omitempty"`
}

func intelligenceCalibrationValCAllSurfaceRefs() []string {
	return []string{
		"/v1/intelligence/calibration/valc/feedback-intake",
		"/v1/intelligence/calibration/valc/feedback-review",
		"/v1/intelligence/calibration/valc/tuning-proposals",
		"/v1/intelligence/calibration/valc/suppression-safety",
		"/v1/intelligence/calibration/valc/suppression-rollback",
		"/v1/intelligence/calibration/valc/local-change-review",
		"/v1/intelligence/calibration/valc/federated-weighting",
		"/v1/intelligence/calibration/valc/similarity-gating",
		"/v1/intelligence/calibration/valc/local-override",
		"/v1/intelligence/calibration/valc/propagation-policy",
		"/v1/intelligence/calibration/valc/explanations",
		"/v1/intelligence/calibration/valc/proofs",
	}
}

func intelligenceCalibrationValCEvidenceRefs() []string {
	return []string{
		"val0_proofs",
		"vala_proofs",
		"valb_proofs",
		"structured_feedback_intake",
		"feedback_review_cockpit",
		"tuning_proposal_generation",
		"suppression_safety_application",
		"suppression_rollback_path",
		"local_calibration_change_review",
		"federated_signal_weighting",
		"environment_similarity_gating",
		"local_override_discipline",
		"bounded_propagation_policy",
		"feedback_federated_explanation",
		"evidence_spine",
	}
}

func intelligenceCalibrationValCProjectionDisclaimer() string {
	return "projection_only not_canonical_truth advisory_feedback_suppression_federated_tuning"
}

func (s server) intelligenceCalibrationValCFeedbackIntakeHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCFeedbackIntake())
}

func (s server) intelligenceCalibrationValCReviewCockpitHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCReviewCockpit())
}

func (s server) intelligenceCalibrationValCTuningProposalHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCTuningProposal())
}

func (s server) intelligenceCalibrationValCSuppressionSafetyHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCSuppressionSafety())
}

func (s server) intelligenceCalibrationValCSuppressionRollbackHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCSuppressionRollback())
}

func (s server) intelligenceCalibrationValCLocalChangeReviewHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCLocalChangeReview())
}

func (s server) intelligenceCalibrationValCFederatedWeightingHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCFederatedWeighting())
}

func (s server) intelligenceCalibrationValCSimilarityGatingHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCSimilarityGating())
}

func (s server) intelligenceCalibrationValCLocalOverrideHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCLocalOverride())
}

func (s server) intelligenceCalibrationValCPropagationPolicyHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCPropagationPolicy())
}

func (s server) intelligenceCalibrationValCExplanationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCExplanation())
}

func (s server) intelligenceCalibrationValCProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValCProofs())
}

func buildIntelligenceCalibrationValCFeedbackIntake() intelligenceCalibrationValCFeedbackIntakeResponse {
	model := operability.IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	return intelligenceCalibrationValCFeedbackIntakeResponse{
		SchemaVersion: intelligenceCalibrationValCFeedbackIntakeSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCFeedbackIntakeState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/feedback-intake"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValCReviewCockpit() intelligenceCalibrationValCReviewCockpitResponse {
	model := operability.IntelligenceCalibrationValCFeedbackReviewCockpitContract()
	return intelligenceCalibrationValCReviewCockpitResponse{
		SchemaVersion: intelligenceCalibrationValCReviewCockpitSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCReviewCockpitState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/feedback-review"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValCTuningProposal() intelligenceCalibrationValCTuningProposalResponse {
	model := operability.IntelligenceCalibrationValCTuningProposalContract()
	return intelligenceCalibrationValCTuningProposalResponse{
		SchemaVersion: intelligenceCalibrationValCTuningProposalSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCTuningProposalState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/tuning-proposals"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValCSuppressionSafety() intelligenceCalibrationValCSuppressionSafetyResponse {
	model := operability.IntelligenceCalibrationValCSuppressionSafetyContract()
	return intelligenceCalibrationValCSuppressionSafetyResponse{
		SchemaVersion: intelligenceCalibrationValCSuppressionSafetySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCSuppressionSafetyState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/suppression-safety"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValCSuppressionRollback() intelligenceCalibrationValCSuppressionRollbackResponse {
	model := operability.IntelligenceCalibrationValCSuppressionRollbackContract()
	return intelligenceCalibrationValCSuppressionRollbackResponse{
		SchemaVersion: intelligenceCalibrationValCSuppressionRollbackSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCSuppressionRollbackState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/suppression-rollback"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValCLocalChangeReview() intelligenceCalibrationValCLocalChangeReviewResponse {
	model := operability.IntelligenceCalibrationValCLocalChangeReviewContract()
	return intelligenceCalibrationValCLocalChangeReviewResponse{
		SchemaVersion: intelligenceCalibrationValCLocalChangeReviewSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCLocalChangeReviewState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/local-change-review"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValCFederatedWeighting() intelligenceCalibrationValCFederatedWeightingResponse {
	model := operability.IntelligenceCalibrationValCFederatedSignalWeightingContract()
	return intelligenceCalibrationValCFederatedWeightingResponse{
		SchemaVersion: intelligenceCalibrationValCFederatedWeightingSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCFederatedWeightingState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/federated-weighting"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValCSimilarityGating() intelligenceCalibrationValCSimilarityGatingResponse {
	model := operability.IntelligenceCalibrationValCSimilarityGatingContract()
	return intelligenceCalibrationValCSimilarityGatingResponse{
		SchemaVersion: intelligenceCalibrationValCSimilarityGatingSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCSimilarityGatingState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/similarity-gating"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValCLocalOverride() intelligenceCalibrationValCLocalOverrideResponse {
	model := operability.IntelligenceCalibrationValCLocalOverrideDisciplineContract()
	return intelligenceCalibrationValCLocalOverrideResponse{
		SchemaVersion: intelligenceCalibrationValCLocalOverrideSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCLocalOverrideState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/local-override"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValCPropagationPolicy() intelligenceCalibrationValCPropagationPolicyResponse {
	model := operability.IntelligenceCalibrationValCPropagationPolicyContract()
	return intelligenceCalibrationValCPropagationPolicyResponse{
		SchemaVersion: intelligenceCalibrationValCPropagationPolicySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCPropagationPolicyState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/propagation-policy"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValCExplanation() intelligenceCalibrationValCExplanationResponse {
	model := operability.IntelligenceCalibrationValCExplanationContract()
	return intelligenceCalibrationValCExplanationResponse{
		SchemaVersion: intelligenceCalibrationValCExplanationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValCExplanationState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/valc/explanations"},
		Limitations:   []string{model.LocalOverrideNote, model.PropagationNote, model.NextStep},
	}
}

func buildIntelligenceCalibrationValCProofs() intelligenceCalibrationValCProofsResponse {
	val0 := buildIntelligenceCalibrationVal0Proofs()
	valA := buildIntelligenceCalibrationValAProofs()
	valB := buildIntelligenceCalibrationValBProofs()
	feedbackIntake := operability.IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	reviewCockpit := operability.IntelligenceCalibrationValCFeedbackReviewCockpitContract()
	tuningProposal := operability.IntelligenceCalibrationValCTuningProposalContract()
	suppressionSafety := operability.IntelligenceCalibrationValCSuppressionSafetyContract()
	suppressionRollback := operability.IntelligenceCalibrationValCSuppressionRollbackContract()
	localChangeReview := operability.IntelligenceCalibrationValCLocalChangeReviewContract()
	federatedWeighting := operability.IntelligenceCalibrationValCFederatedSignalWeightingContract()
	similarityGating := operability.IntelligenceCalibrationValCSimilarityGatingContract()
	localOverride := operability.IntelligenceCalibrationValCLocalOverrideDisciplineContract()
	propagationPolicy := operability.IntelligenceCalibrationValCPropagationPolicyContract()
	explanation := operability.IntelligenceCalibrationValCExplanationContract()

	feedbackIntakeState := operability.EvaluateIntelligenceCalibrationValCFeedbackIntakeState(feedbackIntake)
	reviewCockpitState := operability.EvaluateIntelligenceCalibrationValCReviewCockpitState(reviewCockpit)
	tuningProposalState := operability.EvaluateIntelligenceCalibrationValCTuningProposalState(tuningProposal)
	suppressionSafetyState := operability.EvaluateIntelligenceCalibrationValCSuppressionSafetyState(suppressionSafety)
	suppressionRollbackState := operability.EvaluateIntelligenceCalibrationValCSuppressionRollbackState(suppressionRollback)
	localChangeReviewState := operability.EvaluateIntelligenceCalibrationValCLocalChangeReviewState(localChangeReview)
	federatedWeightingState := operability.EvaluateIntelligenceCalibrationValCFederatedWeightingState(federatedWeighting)
	similarityGatingState := operability.EvaluateIntelligenceCalibrationValCSimilarityGatingState(similarityGating)
	localOverrideState := operability.EvaluateIntelligenceCalibrationValCLocalOverrideState(localOverride)
	propagationPolicyState := operability.EvaluateIntelligenceCalibrationValCPropagationPolicyState(propagationPolicy)
	explanationState := operability.EvaluateIntelligenceCalibrationValCExplanationState(explanation)

	surfaceRefs := intelligenceCalibrationValCAllSurfaceRefs()
	evidenceRefs := intelligenceCalibrationValCEvidenceRefs()
	limitations := []string{
		"Val C proves only Feedback, Suppression & Federated Tuning readiness and does not claim complete intelligence calibration, simulation, or integrated closure.",
		"Feedback intake, tuning proposals, suppression candidates, federated hints, similarity gating, overrides, and propagation policy remain advisory projections over canonical evidence and cannot mutate active calibration, enforcement, priority, or final VEX state in Val C.",
	}
	whyPoint5NotPass := []string{
		"Točka 5 remains not complete because later waves still need defensive simulation, final calibration gate, and integrated closure.",
		"Val C adds bounded feedback, suppression-candidate, and federated tuning discipline only; it does not add active suppression, remote mutation, automatic propagation, or final authority.",
	}

	valCState := operability.EvaluateIntelligenceCalibrationValCState(
		val0.CurrentState,
		val0.Val0State,
		valA.CurrentState,
		valA.ValAState,
		valB.CurrentState,
		valB.ValBState,
		feedbackIntakeState,
		reviewCockpitState,
		tuningProposalState,
		suppressionSafetyState,
		suppressionRollbackState,
		localChangeReviewState,
		federatedWeightingState,
		similarityGatingState,
		localOverrideState,
		propagationPolicyState,
		explanationState,
	)
	currentState := operability.EvaluateIntelligenceCalibrationValCProofsState(
		val0.CurrentState,
		val0.Val0State,
		valA.CurrentState,
		valA.ValAState,
		valB.CurrentState,
		valB.ValBState,
		feedbackIntakeState,
		reviewCockpitState,
		tuningProposalState,
		suppressionSafetyState,
		suppressionRollbackState,
		localChangeReviewState,
		federatedWeightingState,
		similarityGatingState,
		localOverrideState,
		propagationPolicyState,
		explanationState,
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint5NotPass,
		intelligenceCalibrationValCProjectionDisclaimer(),
	)

	return intelligenceCalibrationValCProofsResponse{
		SchemaVersion:               intelligenceCalibrationValCProofsSchema,
		GeneratedAt:                 publicSampleTime(),
		CurrentState:                currentState,
		Val0DependencyState:         val0.CurrentState,
		Val0FoundationState:         val0.Val0State,
		ValADependencyState:         valA.CurrentState,
		ValAReachabilityVEXState:    valA.ValAState,
		ValBDependencyState:         valB.CurrentState,
		ValBBehavioralLearningState: valB.ValBState,
		ValCState:                   valCState,
		Point5State:                 operability.IntelligenceCalibrationPoint5StateNotComplete,
		FeedbackIntakeState:         feedbackIntakeState,
		ReviewCockpitState:          reviewCockpitState,
		TuningProposalState:         tuningProposalState,
		SuppressionSafetyState:      suppressionSafetyState,
		SuppressionRollbackState:    suppressionRollbackState,
		LocalChangeReviewState:      localChangeReviewState,
		FederatedWeightingState:     federatedWeightingState,
		SimilarityGatingState:       similarityGatingState,
		LocalOverrideState:          localOverrideState,
		PropagationPolicyState:      propagationPolicyState,
		ExplanationState:            explanationState,
		WhyPoint5NotPass:            whyPoint5NotPass,
		SurfaceRefs:                 surfaceRefs,
		EvidenceRefs:                evidenceRefs,
		Limitations:                 limitations,
		ProjectionDisclaimer:        intelligenceCalibrationValCProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val C adds bounded structured feedback intake, review cockpit, tuning proposals, suppression candidate safety/rollback, local change review, federated weighting, similarity gating, override discipline, propagation policy, and explanations.",
			"Val C is fail-closed on active Val 0, Val A, and Val B proofs and keeps all feedback, suppression, and federated outputs advisory rather than canonical or mutating.",
			"Točka 5 remains not complete until later waves add defensive simulation, final calibration gate, and integrated closure.",
		},
	}
}

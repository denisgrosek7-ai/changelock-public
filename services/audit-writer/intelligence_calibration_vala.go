package main

import (
	"net/http"
	"time"

	"github.com/denisgrosek/changelock/internal/httpjson"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

const (
	intelligenceCalibrationValAAggregationSchema    = "point5.intelligence_calibration.vala.reachability_aggregation.v1"
	intelligenceCalibrationValAExploitabilitySchema = "point5.intelligence_calibration.vala.exploitability_calibration.v1"
	intelligenceCalibrationValADecisionSchema       = "point5.intelligence_calibration.vala.downgrade_escalation.v1"
	intelligenceCalibrationValACAVISchema           = "point5.intelligence_calibration.vala.cavi_tuning.v1"
	intelligenceCalibrationValAVEXSchema            = "point5.intelligence_calibration.vala.vex_candidate_calibration.v1"
	intelligenceCalibrationValASufficiencySchema    = "point5.intelligence_calibration.vala.vex_sufficiency.v1"
	intelligenceCalibrationValAExplanationSchema    = "point5.intelligence_calibration.vala.explanations.v1"
	intelligenceCalibrationValAOutcomeSchema        = "point5.intelligence_calibration.vala.confidence_outcomes.v1"
	intelligenceCalibrationValAGuardrailSchema      = "point5.intelligence_calibration.vala.publication_guardrail.v1"
	intelligenceCalibrationValAProofsSchema         = "point5.intelligence_calibration.vala.proofs.v1"
)

type intelligenceCalibrationValAAggregationResponse struct {
	SchemaVersion string                                            `json:"schema_version"`
	GeneratedAt   time.Time                                         `json:"generated_at"`
	CurrentState  string                                            `json:"current_state"`
	Model         operability.ReachabilitySignalAggregationContract `json:"model"`
	RouteRefs     []string                                          `json:"route_refs,omitempty"`
	Limitations   []string                                          `json:"limitations,omitempty"`
}

type intelligenceCalibrationValAExploitabilityResponse struct {
	SchemaVersion string                                                  `json:"schema_version"`
	GeneratedAt   time.Time                                               `json:"generated_at"`
	CurrentState  string                                                  `json:"current_state"`
	Model         operability.ContextualExploitabilityCalibrationContract `json:"model"`
	RouteRefs     []string                                                `json:"route_refs,omitempty"`
	Limitations   []string                                                `json:"limitations,omitempty"`
}

type intelligenceCalibrationValADecisionResponse struct {
	SchemaVersion string                                            `json:"schema_version"`
	GeneratedAt   time.Time                                         `json:"generated_at"`
	CurrentState  string                                            `json:"current_state"`
	Model         operability.DowngradeEscalationDisciplineContract `json:"model"`
	RouteRefs     []string                                          `json:"route_refs,omitempty"`
	Limitations   []string                                          `json:"limitations,omitempty"`
}

type intelligenceCalibrationValACAVIResponse struct {
	SchemaVersion string                                     `json:"schema_version"`
	GeneratedAt   time.Time                                  `json:"generated_at"`
	CurrentState  string                                     `json:"current_state"`
	Model         operability.CAVIReachabilityTuningContract `json:"model"`
	RouteRefs     []string                                   `json:"route_refs,omitempty"`
	Limitations   []string                                   `json:"limitations,omitempty"`
}

type intelligenceCalibrationValAVEXResponse struct {
	SchemaVersion string                                      `json:"schema_version"`
	GeneratedAt   time.Time                                   `json:"generated_at"`
	CurrentState  string                                      `json:"current_state"`
	Model         operability.VEXCandidateCalibrationContract `json:"model"`
	RouteRefs     []string                                    `json:"route_refs,omitempty"`
	Limitations   []string                                    `json:"limitations,omitempty"`
}

type intelligenceCalibrationValASufficiencyResponse struct {
	SchemaVersion string                                     `json:"schema_version"`
	GeneratedAt   time.Time                                  `json:"generated_at"`
	CurrentState  string                                     `json:"current_state"`
	Model         operability.VEXEvidenceSufficiencyContract `json:"model"`
	RouteRefs     []string                                   `json:"route_refs,omitempty"`
	Limitations   []string                                   `json:"limitations,omitempty"`
}

type intelligenceCalibrationValAExplanationResponse struct {
	SchemaVersion string                                         `json:"schema_version"`
	GeneratedAt   time.Time                                      `json:"generated_at"`
	CurrentState  string                                         `json:"current_state"`
	Model         operability.ReachabilityVEXExplanationContract `json:"model"`
	RouteRefs     []string                                       `json:"route_refs,omitempty"`
	Limitations   []string                                       `json:"limitations,omitempty"`
}

type intelligenceCalibrationValAOutcomeResponse struct {
	SchemaVersion string                                                 `json:"schema_version"`
	GeneratedAt   time.Time                                              `json:"generated_at"`
	CurrentState  string                                                 `json:"current_state"`
	Model         operability.ConfidenceBoundReachabilityOutcomeContract `json:"model"`
	RouteRefs     []string                                               `json:"route_refs,omitempty"`
	Limitations   []string                                               `json:"limitations,omitempty"`
}

type intelligenceCalibrationValAGuardrailResponse struct {
	SchemaVersion string                                             `json:"schema_version"`
	GeneratedAt   time.Time                                          `json:"generated_at"`
	CurrentState  string                                             `json:"current_state"`
	Model         operability.NoFinalPublicationVEXGuardrailContract `json:"model"`
	RouteRefs     []string                                           `json:"route_refs,omitempty"`
	Limitations   []string                                           `json:"limitations,omitempty"`
}

type intelligenceCalibrationValAProofsResponse struct {
	SchemaVersion             string    `json:"schema_version"`
	GeneratedAt               time.Time `json:"generated_at"`
	CurrentState              string    `json:"current_state"`
	Val0DependencyState       string    `json:"val_0_dependency_state"`
	Val0FoundationState       string    `json:"val_0_foundation_state"`
	ValAState                 string    `json:"val_a_state"`
	Point5State               string    `json:"point_5_state"`
	AggregationState          string    `json:"reachability_signal_aggregation_state"`
	ExploitabilityState       string    `json:"contextual_exploitability_state"`
	DecisionState             string    `json:"downgrade_escalation_discipline_state"`
	CAVIState                 string    `json:"cavi_reachability_tuning_state"`
	VEXCandidateState         string    `json:"vex_candidate_calibration_state"`
	VEXSufficiencyState       string    `json:"vex_evidence_sufficiency_state"`
	ExplanationState          string    `json:"reachability_vex_explanation_state"`
	ConfidenceOutcomeState    string    `json:"confidence_bound_outcome_state"`
	PublicationGuardrailState string    `json:"no_final_publication_guardrail_state"`
	WhyPoint5NotPass          []string  `json:"why_point_5_not_pass,omitempty"`
	SurfaceRefs               []string  `json:"surface_refs,omitempty"`
	EvidenceRefs              []string  `json:"evidence_refs,omitempty"`
	Limitations               []string  `json:"limitations,omitempty"`
	ProjectionDisclaimer      string    `json:"projection_disclaimer"`
	IntegrationSummary        []string  `json:"integration_summary,omitempty"`
}

func intelligenceCalibrationValAAllSurfaceRefs() []string {
	return []string{
		"/v1/intelligence/calibration/vala/reachability-aggregation",
		"/v1/intelligence/calibration/vala/exploitability-calibration",
		"/v1/intelligence/calibration/vala/downgrade-escalation",
		"/v1/intelligence/calibration/vala/cavi-tuning",
		"/v1/intelligence/calibration/vala/vex-candidates",
		"/v1/intelligence/calibration/vala/vex-sufficiency",
		"/v1/intelligence/calibration/vala/explanations",
		"/v1/intelligence/calibration/vala/confidence-outcomes",
		"/v1/intelligence/calibration/vala/publication-guardrail",
		"/v1/intelligence/calibration/vala/proofs",
	}
}

func intelligenceCalibrationValAEvidenceRefs() []string {
	return []string{
		"val0_proofs",
		"reachability_signal_aggregation",
		"contextual_exploitability_calibration",
		"downgrade_escalation_guardrail",
		"cavi_tuning_contract",
		"vex_candidate_calibration",
		"vex_sufficiency_check",
		"reachability_vex_explanation",
		"confidence_bound_outcome",
		"no_final_publication_guardrail",
		"evidence_spine",
	}
}

func intelligenceCalibrationValAProjectionDisclaimer() string {
	return "projection_only not_canonical_truth advisory_reachability_vex_calibration"
}

func (s server) intelligenceCalibrationValAAggregationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValAAggregation())
}

func (s server) intelligenceCalibrationValAExploitabilityHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValAExploitability())
}

func (s server) intelligenceCalibrationValADecisionHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValADecision())
}

func (s server) intelligenceCalibrationValACAVIHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValACAVI())
}

func (s server) intelligenceCalibrationValAVEXHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValAVEX())
}

func (s server) intelligenceCalibrationValASufficiencyHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValASufficiency())
}

func (s server) intelligenceCalibrationValAExplanationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValAExplanation())
}

func (s server) intelligenceCalibrationValAOutcomeHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValAOutcome())
}

func (s server) intelligenceCalibrationValAGuardrailHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValAGuardrail())
}

func (s server) intelligenceCalibrationValAProofsHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := s.enterpriseWorkflowAuthorityVal0AuthorizeRead(w, r); !ok {
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildIntelligenceCalibrationValAProofs())
}

func buildIntelligenceCalibrationValAAggregation() intelligenceCalibrationValAAggregationResponse {
	model := operability.IntelligenceCalibrationValAReachabilityAggregationContract()
	return intelligenceCalibrationValAAggregationResponse{
		SchemaVersion: intelligenceCalibrationValAAggregationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValAAggregationState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vala/reachability-aggregation"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValAExploitability() intelligenceCalibrationValAExploitabilityResponse {
	model := operability.IntelligenceCalibrationValAExploitabilityCalibrationContract()
	return intelligenceCalibrationValAExploitabilityResponse{
		SchemaVersion: intelligenceCalibrationValAExploitabilitySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValAExploitabilityState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vala/exploitability-calibration"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValADecision() intelligenceCalibrationValADecisionResponse {
	model := operability.IntelligenceCalibrationValADowngradeEscalationContract()
	return intelligenceCalibrationValADecisionResponse{
		SchemaVersion: intelligenceCalibrationValADecisionSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValADecisionState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vala/downgrade-escalation"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValACAVI() intelligenceCalibrationValACAVIResponse {
	model := operability.IntelligenceCalibrationValACAVITuningContract()
	return intelligenceCalibrationValACAVIResponse{
		SchemaVersion: intelligenceCalibrationValACAVISchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValACAVIState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vala/cavi-tuning"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValAVEX() intelligenceCalibrationValAVEXResponse {
	model := operability.IntelligenceCalibrationValAVEXCandidateContract()
	return intelligenceCalibrationValAVEXResponse{
		SchemaVersion: intelligenceCalibrationValAVEXSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValAVEXCandidateState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vala/vex-candidates"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValASufficiency() intelligenceCalibrationValASufficiencyResponse {
	model := operability.IntelligenceCalibrationValAVEXSufficiencyContract()
	return intelligenceCalibrationValASufficiencyResponse{
		SchemaVersion: intelligenceCalibrationValASufficiencySchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValAVEXSufficiencyState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vala/vex-sufficiency"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValAExplanation() intelligenceCalibrationValAExplanationResponse {
	model := operability.IntelligenceCalibrationValAExplanationContract()
	return intelligenceCalibrationValAExplanationResponse{
		SchemaVersion: intelligenceCalibrationValAExplanationSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValAExplanationState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vala/explanations"},
		Limitations:   []string{model.UncertaintyNote, model.NextStep},
	}
}

func buildIntelligenceCalibrationValAOutcome() intelligenceCalibrationValAOutcomeResponse {
	model := operability.IntelligenceCalibrationValAConfidenceOutcomeContract()
	return intelligenceCalibrationValAOutcomeResponse{
		SchemaVersion: intelligenceCalibrationValAOutcomeSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValAConfidenceOutcomeState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vala/confidence-outcomes"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValAGuardrail() intelligenceCalibrationValAGuardrailResponse {
	model := operability.IntelligenceCalibrationValAPublicationGuardrailContract()
	return intelligenceCalibrationValAGuardrailResponse{
		SchemaVersion: intelligenceCalibrationValAGuardrailSchema,
		GeneratedAt:   publicSampleTime(),
		CurrentState:  operability.EvaluateIntelligenceCalibrationValAPublicationGuardrailState(model),
		Model:         model,
		RouteRefs:     []string{"/v1/intelligence/calibration/vala/publication-guardrail"},
		Limitations:   []string{model.LimitationMessage},
	}
}

func buildIntelligenceCalibrationValAProofs() intelligenceCalibrationValAProofsResponse {
	val0 := buildIntelligenceCalibrationVal0Proofs()
	aggregation := operability.IntelligenceCalibrationValAReachabilityAggregationContract()
	exploitability := operability.IntelligenceCalibrationValAExploitabilityCalibrationContract()
	decision := operability.IntelligenceCalibrationValADowngradeEscalationContract()
	cavi := operability.IntelligenceCalibrationValACAVITuningContract()
	vex := operability.IntelligenceCalibrationValAVEXCandidateContract()
	sufficiency := operability.IntelligenceCalibrationValAVEXSufficiencyContract()
	explanation := operability.IntelligenceCalibrationValAExplanationContract()
	outcome := operability.IntelligenceCalibrationValAConfidenceOutcomeContract()
	guardrail := operability.IntelligenceCalibrationValAPublicationGuardrailContract()

	aggregationState := operability.EvaluateIntelligenceCalibrationValAAggregationState(aggregation)
	exploitabilityState := operability.EvaluateIntelligenceCalibrationValAExploitabilityState(exploitability)
	decisionState := operability.EvaluateIntelligenceCalibrationValADecisionState(decision)
	caviState := operability.EvaluateIntelligenceCalibrationValACAVIState(cavi)
	vexState := operability.EvaluateIntelligenceCalibrationValAVEXCandidateState(vex)
	sufficiencyState := operability.EvaluateIntelligenceCalibrationValAVEXSufficiencyState(sufficiency)
	explanationState := operability.EvaluateIntelligenceCalibrationValAExplanationState(explanation)
	outcomeState := operability.EvaluateIntelligenceCalibrationValAConfidenceOutcomeState(outcome)
	guardrailState := operability.EvaluateIntelligenceCalibrationValAPublicationGuardrailState(guardrail)

	surfaceRefs := intelligenceCalibrationValAAllSurfaceRefs()
	evidenceRefs := intelligenceCalibrationValAEvidenceRefs()
	limitations := []string{
		"Val A proves only Reachability & VEX Calibration readiness and does not claim complete intelligence calibration or final VEX publication.",
		"Reachability, exploitability, and VEX candidate outputs remain advisory projections over canonical evidence and cannot auto-mutate priority, suppress evidence, or publish final VEX.",
	}
	whyPoint5NotPass := []string{
		"Točka 5 remains not complete because later waves still need behavioral baselines, feedback/federated tuning, simulation, final calibration gate, and integrated closure.",
		"Val A adds bounded reachability and VEX candidate calibration only; it does not add final publication, suppression, or integrated authority.",
	}

	valAState := operability.EvaluateIntelligenceCalibrationValAState(
		val0.CurrentState,
		val0.Val0State,
		aggregationState,
		exploitabilityState,
		decisionState,
		caviState,
		vexState,
		sufficiencyState,
		explanationState,
		outcomeState,
		guardrailState,
	)
	currentState := operability.EvaluateIntelligenceCalibrationValAProofsState(
		val0.CurrentState,
		val0.Val0State,
		aggregationState,
		exploitabilityState,
		decisionState,
		caviState,
		vexState,
		sufficiencyState,
		explanationState,
		outcomeState,
		guardrailState,
		surfaceRefs,
		evidenceRefs,
		limitations,
		whyPoint5NotPass,
		intelligenceCalibrationValAProjectionDisclaimer(),
	)

	return intelligenceCalibrationValAProofsResponse{
		SchemaVersion:             intelligenceCalibrationValAProofsSchema,
		GeneratedAt:               publicSampleTime(),
		CurrentState:              currentState,
		Val0DependencyState:       val0.CurrentState,
		Val0FoundationState:       val0.Val0State,
		ValAState:                 valAState,
		Point5State:               operability.IntelligenceCalibrationPoint5StateNotComplete,
		AggregationState:          aggregationState,
		ExploitabilityState:       exploitabilityState,
		DecisionState:             decisionState,
		CAVIState:                 caviState,
		VEXCandidateState:         vexState,
		VEXSufficiencyState:       sufficiencyState,
		ExplanationState:          explanationState,
		ConfidenceOutcomeState:    outcomeState,
		PublicationGuardrailState: guardrailState,
		WhyPoint5NotPass:          whyPoint5NotPass,
		SurfaceRefs:               surfaceRefs,
		EvidenceRefs:              evidenceRefs,
		Limitations:               limitations,
		ProjectionDisclaimer:      intelligenceCalibrationValAProjectionDisclaimer(),
		IntegrationSummary: []string{
			"Val A adds bounded reachability aggregation, contextual exploitability, CAVI-style tuning, VEX candidate calibration, sufficiency checks, explanations, confidence outcomes, and explicit no-final-publication guardrails.",
			"Val A is fail-closed on Val 0 calibration foundation proofs and keeps all reachability and VEX outputs advisory rather than canonical or final.",
			"Točka 5 remains not complete until later waves add additional calibration engines, final gate behavior, and integrated closure.",
		},
	}
}

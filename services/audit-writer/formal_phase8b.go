package main

import (
	"net/http"
	"strings"
	"time"

	formalcore "github.com/denisgrosek/changelock/internal/formal"
	"github.com/denisgrosek/changelock/internal/httpjson"
)

const (
	phase8ConsensusReviewSchema            = "8b.formal_governance_consensus_review.v1"
	phase8PolicySuggestionsSchema          = "8b.formal_governance_policy_suggestions.v1"
	phase8AuthorityRoutingSchema           = "8b.formal_governance_authority_routing.v1"
	phase8AIGuardrailsSchema               = "8b.formal_governance_ai_guardrails.v1"
	phase8ModelRiskSchema                  = "8b.formal_governance_model_risk.v1"
	phase8ConsensusReviewStateActive       = "phase8b_consensus_review_active"
	phase8ConsensusReviewStatePartial      = "phase8b_consensus_review_partial"
	phase8ConsensusReviewStateIncomplete   = "phase8b_consensus_review_incomplete"
	phase8PolicySuggestionsStateActive     = "phase8b_policy_suggestions_active"
	phase8PolicySuggestionsStatePartial    = "phase8b_policy_suggestions_partial"
	phase8PolicySuggestionsStateIncomplete = "phase8b_policy_suggestions_incomplete"
	phase8AuthorityRoutingStateActive      = "phase8b_authority_routing_active"
	phase8AuthorityRoutingStatePartial     = "phase8b_authority_routing_partial"
	phase8AuthorityRoutingStateIncomplete  = "phase8b_authority_routing_incomplete"
	phase8AIGuardrailsStateActive          = "phase8b_ai_guardrails_active"
	phase8AIGuardrailsStatePartial         = "phase8b_ai_guardrails_partial"
	phase8AIGuardrailsStateIncomplete      = "phase8b_ai_guardrails_incomplete"
	phase8ModelRiskStateActive             = "phase8b_model_risk_active"
	phase8ModelRiskStatePartial            = "phase8b_model_risk_partial"
	phase8ModelRiskStateIncomplete         = "phase8b_model_risk_incomplete"
)

type phase8ConsensusReviewItem struct {
	ReviewID                  string   `json:"review_id"`
	CurrentState              string   `json:"current_state"`
	Subject                   string   `json:"subject"`
	QuorumThreshold           string   `json:"quorum_threshold"`
	AbstainState              string   `json:"abstain_state"`
	WeightedDisagreementState string   `json:"weighted_disagreement_state"`
	MinorityReportSupport     string   `json:"minority_report_support"`
	ReviewRequired            bool     `json:"review_required"`
	ReasonCodes               []string `json:"reason_codes,omitempty"`
	RouteRefs                 []string `json:"route_refs,omitempty"`
	Limitations               []string `json:"limitations,omitempty"`
}

type phase8ConsensusReviewResponse struct {
	SchemaVersion  string                      `json:"schema_version"`
	GeneratedAt    time.Time                   `json:"generated_at"`
	CurrentState   string                      `json:"current_state"`
	ConsensusState string                      `json:"consensus_state"`
	QuorumState    string                      `json:"quorum_state"`
	MinorityState  string                      `json:"minority_state"`
	Items          []phase8ConsensusReviewItem `json:"items,omitempty"`
	RouteRefs      []string                    `json:"route_refs,omitempty"`
	Limitations    []string                    `json:"limitations,omitempty"`
}

type phase8PolicySuggestion struct {
	SuggestionID         string   `json:"suggestion_id"`
	CurrentState         string   `json:"current_state"`
	Summary              string   `json:"summary"`
	BlastRadiusEstimate  string   `json:"blast_radius_estimate"`
	RollbackFeasibility  string   `json:"rollback_feasibility"`
	CompatibilityWarning string   `json:"compatibility_warning"`
	ApprovalState        string   `json:"approval_state"`
	ForbiddenActions     []string `json:"forbidden_actions,omitempty"`
	ReasonCodes          []string `json:"reason_codes,omitempty"`
	RouteRefs            []string `json:"route_refs,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
}

type phase8PolicySuggestionsResponse struct {
	SchemaVersion    string                   `json:"schema_version"`
	GeneratedAt      time.Time                `json:"generated_at"`
	CurrentState     string                   `json:"current_state"`
	SuggestionState  string                   `json:"suggestion_state"`
	ApprovalBoundary string                   `json:"approval_boundary"`
	Suggestions      []phase8PolicySuggestion `json:"suggestions,omitempty"`
	RouteRefs        []string                 `json:"route_refs,omitempty"`
	Limitations      []string                 `json:"limitations,omitempty"`
}

type phase8AuthorityRoutingResponse struct {
	SchemaVersion            string                       `json:"schema_version"`
	GeneratedAt              time.Time                    `json:"generated_at"`
	CurrentState             string                       `json:"current_state"`
	RoutingState             string                       `json:"routing_state"`
	AlternateApproverPath    string                       `json:"alternate_approver_path"`
	DeadlockResolutionPath   string                       `json:"deadlock_resolution_path"`
	EmergencySuspensionPath  string                       `json:"emergency_suspension_path"`
	NonDelegableActions      []string                     `json:"non_delegable_actions,omitempty"`
	SeparationOfDutiesRules  []string                     `json:"separation_of_duties_rules,omitempty"`
	QuorumRules              []string                     `json:"quorum_rules,omitempty"`
	ConstitutionalBoundaries []string                     `json:"constitutional_boundaries,omitempty"`
	ChallengeWorkflow        formalcore.ChallengeWorkflow `json:"challenge_workflow"`
	RouteRefs                []string                     `json:"route_refs,omitempty"`
	Limitations              []string                     `json:"limitations,omitempty"`
}

type phase8AIGuardrailsResponse struct {
	SchemaVersion            string                   `json:"schema_version"`
	GeneratedAt              time.Time                `json:"generated_at"`
	CurrentState             string                   `json:"current_state"`
	GuardrailState           string                   `json:"guardrail_state"`
	EscalationState          string                   `json:"escalation_state"`
	ConfidenceState          string                   `json:"confidence_state"`
	ConstitutionalBoundaries []string                 `json:"constitutional_boundaries,omitempty"`
	Guardrails               []formalcore.AIGuardrail `json:"guardrails,omitempty"`
	RouteRefs                []string                 `json:"route_refs,omitempty"`
	Limitations              []string                 `json:"limitations,omitempty"`
}

type phase8ModelRiskResponse struct {
	SchemaVersion           string                               `json:"schema_version"`
	GeneratedAt             time.Time                            `json:"generated_at"`
	CurrentState            string                               `json:"current_state"`
	ModelRiskState          string                               `json:"model_risk_state"`
	DependencyRegistryState string                               `json:"dependency_registry_state"`
	ReviewState             string                               `json:"review_state"`
	ModelRiskContracts      []formalcore.ModelRiskContract       `json:"model_risk_contracts,omitempty"`
	Dependencies            []formalcore.InstitutionalDependency `json:"dependencies,omitempty"`
	RouteRefs               []string                             `json:"route_refs,omitempty"`
	Limitations             []string                             `json:"limitations,omitempty"`
}

func (s server) phase8ConsensusReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8ConsensusReview())
}

func (s server) phase8PolicySuggestionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8PolicySuggestions())
}

func (s server) phase8AuthorityRoutingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8AuthorityRouting())
}

func (s server) phase8AIGuardrailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8AIGuardrails())
}

func (s server) phase8ModelRiskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	httpjson.Write(w, http.StatusOK, buildPhase8ModelRisk())
}

func buildPhase8ConsensusReview() phase8ConsensusReviewResponse {
	return phase8ConsensusReviewResponse{
		SchemaVersion:  phase8ConsensusReviewSchema,
		GeneratedAt:    publicSampleTime(),
		CurrentState:   phase8GovernanceSliceState(phase8ConsensusReviewStateActive, phase8ConsensusReviewStatePartial, phase8ConsensusReviewStateIncomplete),
		ConsensusState: "consensus_assisted_review_visible",
		QuorumState:    "quorum_threshold_and_abstain_visible",
		MinorityState:  "weighted_disagreement_and_minority_report_visible",
		Items:          phase8ConsensusReviewItems(),
		RouteRefs: []string{
			"/v1/formal/phase8/governed-autonomy",
			"/v1/formal/phase8/governance/authority-routing",
			"/v1/formal/phase8/proofs",
		},
		Limitations: []string{
			"Consensus-assisted review remains bounded review support and never becomes hidden final authority or automatic external release.",
			"Consensus output remains review-required and cannot override non-delegable actions through silent majority tone.",
		},
	}
}

func buildPhase8PolicySuggestions() phase8PolicySuggestionsResponse {
	return phase8PolicySuggestionsResponse{
		SchemaVersion:    phase8PolicySuggestionsSchema,
		GeneratedAt:      publicSampleTime(),
		CurrentState:     phase8GovernanceSliceState(phase8PolicySuggestionsStateActive, phase8PolicySuggestionsStatePartial, phase8PolicySuggestionsStateIncomplete),
		SuggestionState:  "autonomous_policy_suggestions_visible",
		ApprovalBoundary: "advisory_until_formally_approved",
		Suggestions:      phase8PolicySuggestions(),
		RouteRefs: []string{
			"/v1/formal/phase8/governance/consensus-review",
			"/v1/formal/phase8/governance/authority-routing",
			"/v1/formal/phase8/governance/ai-guardrails",
		},
		Limitations: []string{
			"Policy evolution suggestions remain advisory only until approved through formal authority routing.",
			"Suggestions cannot trigger non-delegable actions, widen audience, or silently mutate active formal profiles.",
		},
	}
}

func buildPhase8AuthorityRouting() phase8AuthorityRoutingResponse {
	control := phase8AuthorityControl()
	return phase8AuthorityRoutingResponse{
		SchemaVersion:            phase8AuthorityRoutingSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             phase8GovernanceSliceState(phase8AuthorityRoutingStateActive, phase8AuthorityRoutingStatePartial, phase8AuthorityRoutingStateIncomplete),
		RoutingState:             "formal_authority_routing_active",
		AlternateApproverPath:    control.AlternateApproverPath,
		DeadlockResolutionPath:   control.DeadlockResolutionPath,
		EmergencySuspensionPath:  control.EmergencySuspensionPath,
		NonDelegableActions:      control.NonDelegableActions,
		SeparationOfDutiesRules:  control.SeparationOfDutiesRules,
		QuorumRules:              control.QuorumRules,
		ConstitutionalBoundaries: phase8ConstitutionalBoundaries(),
		ChallengeWorkflow:        phase8GovernanceChallengeWorkflow(),
		RouteRefs: []string{
			"/v1/formal/phase8/governance/policy-suggestions",
			"/v1/formal/phase8/governance/ai-guardrails",
			"/v1/formal/phase8/proofs",
		},
		Limitations: []string{
			"Authority routing remains explicit, reviewable, and non-hidden; no single model or proposer can silently override the routing path.",
			"Emergency suspension remains a boundary control and not a shortcut for hidden release or unilateral authority expansion.",
		},
	}
}

func buildPhase8AIGuardrails() phase8AIGuardrailsResponse {
	return phase8AIGuardrailsResponse{
		SchemaVersion:            phase8AIGuardrailsSchema,
		GeneratedAt:              publicSampleTime(),
		CurrentState:             phase8GovernanceSliceState(phase8AIGuardrailsStateActive, phase8AIGuardrailsStatePartial, phase8AIGuardrailsStateIncomplete),
		GuardrailState:           "formal_ai_guardrails_surface_active",
		EscalationState:          "human_escalation_threshold_visible",
		ConfidenceState:          "confidence_floor_visible",
		ConstitutionalBoundaries: phase8ConstitutionalBoundaries(),
		Guardrails:               formalcore.AIGuardrails(),
		RouteRefs: []string{
			"/v1/formal/phase8/governed-autonomy",
			"/v1/formal/phase8/governance/model-risk",
			"/v1/formal/phase8/proofs",
		},
		Limitations: []string{
			"AI guardrails remain explicit enforcement boundaries and do not permit legal verdicts, non-delegable action execution, or silent disclosure widening.",
			"Guardrails constrain suggestions and routing but do not themselves become a replacement authority body.",
		},
	}
}

func buildPhase8ModelRisk() phase8ModelRiskResponse {
	return phase8ModelRiskResponse{
		SchemaVersion:           phase8ModelRiskSchema,
		GeneratedAt:             publicSampleTime(),
		CurrentState:            phase8GovernanceSliceState(phase8ModelRiskStateActive, phase8ModelRiskStatePartial, phase8ModelRiskStateIncomplete),
		ModelRiskState:          "formal_model_risk_baseline_active",
		DependencyRegistryState: "institutional_dependency_registry_visible",
		ReviewState:             "challenger_review_and_rollback_path_visible",
		ModelRiskContracts:      formalcore.ModelRiskContracts(),
		Dependencies:            formalcore.InstitutionalDependencies(),
		RouteRefs: []string{
			"/v1/formal/phase8/governance/ai-guardrails",
			"/v1/formal/phase8/contracts",
			"/v1/formal/phase8/proofs",
		},
		Limitations: []string{
			"Model risk remains bounded to review assistance and governance support; it does not claim autonomous institutional judgment.",
			"Dependency visibility supports traceability and change triggers but does not create hidden external authority dependencies.",
		},
	}
}

func phase8GovernanceSliceState(active, partial, incomplete string) string {
	switch formalcore.EvaluateGovernedAutonomyState() {
	case formalcore.GovernedAutonomyStateActive:
		return active
	case formalcore.GovernedAutonomyStatePartial:
		return partial
	default:
		return incomplete
	}
}

func phase8AuthorityControl() formalcore.AuthorityControl {
	controls := formalcore.AuthorityControls()
	if len(controls) == 0 {
		return formalcore.AuthorityControl{}
	}
	return controls[0]
}

func phase8GovernanceChallengeWorkflow() formalcore.ChallengeWorkflow {
	workflows := formalcore.ChallengeWorkflows()
	if len(workflows) == 0 {
		return formalcore.ChallengeWorkflow{}
	}
	return workflows[0]
}

func phase8ConsensusReviewItems() []phase8ConsensusReviewItem {
	return []phase8ConsensusReviewItem{
		{
			ReviewID:                  "formal_claim_release_consensus",
			CurrentState:              "review_required",
			Subject:                   "formal claim release boundary review",
			QuorumThreshold:           "two_person_rule_plus_cross_function_review",
			AbstainState:              "abstain_visible",
			WeightedDisagreementState: "weighted_disagreement_visible",
			MinorityReportSupport:     "minority_report_supported",
			ReviewRequired:            true,
			ReasonCodes:               []string{"consensus_assist_only", "no_hidden_majority_override"},
			RouteRefs:                 []string{"/v1/formal/phase8/governance/authority-routing"},
			Limitations: []string{
				"Consensus review remains advisory input into formal routing rather than a final release authority.",
			},
		},
		{
			ReviewID:                  "policy_change_consensus",
			CurrentState:              "review_required",
			Subject:                   "policy evolution suggestion review",
			QuorumThreshold:           "cross_function_quorum_required",
			AbstainState:              "abstain_visible",
			WeightedDisagreementState: "weighted_disagreement_visible",
			MinorityReportSupport:     "minority_report_supported",
			ReviewRequired:            true,
			ReasonCodes:               []string{"policy_change_review_required", "minority_report_preserved"},
			RouteRefs:                 []string{"/v1/formal/phase8/governance/policy-suggestions"},
			Limitations: []string{
				"Consensus support cannot silently activate policy changes or widen claim scope without formal approval.",
			},
		},
	}
}

func phase8PolicySuggestions() []phase8PolicySuggestion {
	return []phase8PolicySuggestion{
		{
			SuggestionID:         "tighten_regulator_release_redaction_review",
			CurrentState:         "advisory_pending_review",
			Summary:              "Suggest stricter redaction review before regulator-safe disclosure reuse when jurisdiction overlays conflict.",
			BlastRadiusEstimate:  "low_medium_bounded_to_regulator_release_paths",
			RollbackFeasibility:  "high_reversible_via_formal_profile_rollback",
			CompatibilityWarning: "may require updated manual interpretation notes for overlapping jurisdiction profiles",
			ApprovalState:        "formal_approval_required",
			ForbiddenActions:     []string{"automatic_profile_activation", "direct_external_release"},
			ReasonCodes:          []string{"blast_radius_visible", "rollback_feasible", "compatibility_warning_visible"},
			RouteRefs:            []string{"/v1/formal/phase8/governance/authority-routing"},
			Limitations: []string{
				"Suggestion remains advisory and cannot mutate live formal profiles without explicit approval.",
			},
		},
		{
			SuggestionID:         "strengthen_certification_snapshot_hold",
			CurrentState:         "advisory_pending_review",
			Summary:              "Suggest stronger challenge hold behavior for certification snapshots with unresolved disagreement or ambiguity markers.",
			BlastRadiusEstimate:  "medium_bounded_to_certification_support_workflow",
			RollbackFeasibility:  "medium_requires_challenge_workflow_review",
			CompatibilityWarning: "may reduce snapshot reuse until assessor review timing is updated",
			ApprovalState:        "formal_approval_required",
			ForbiddenActions:     []string{"automatic_certification_release", "hidden_issue_suppression"},
			ReasonCodes:          []string{"support_artifact_only", "challenge_hold_visible"},
			RouteRefs:            []string{"/v1/formal/phase8/governance/consensus-review"},
			Limitations: []string{
				"Suggestion can influence review routing only; it cannot create certification issuance authority.",
			},
		},
	}
}

func phase8ConstitutionalBoundaries() []string {
	return []string{
		"competition_law_caution",
		"anti_capture_rule",
		"no_single_hidden_model_override",
	}
}

func hasForbiddenRecommendationClass(guardrails []formalcore.AIGuardrail, target string) bool {
	for _, guardrail := range guardrails {
		if containsString(guardrail.ProhibitedRecommendationClasses, target) {
			return true
		}
	}
	return false
}

func hasDependencyClass(dependencies []formalcore.InstitutionalDependency, target string) bool {
	for _, dependency := range dependencies {
		if strings.TrimSpace(dependency.DependencyClass) == strings.TrimSpace(target) {
			return true
		}
	}
	return false
}

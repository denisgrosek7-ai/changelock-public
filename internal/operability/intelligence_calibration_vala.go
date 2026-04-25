package operability

import "strings"

const (
	IntelligenceCalibrationValAAggregationStateActive     = "intelligence_calibration_vala_reachability_aggregation_active"
	IntelligenceCalibrationValAAggregationStatePartial    = "intelligence_calibration_vala_reachability_aggregation_partial"
	IntelligenceCalibrationValAAggregationStateIncomplete = "intelligence_calibration_vala_reachability_aggregation_incomplete"

	IntelligenceCalibrationValAExploitabilityStateActive     = "intelligence_calibration_vala_exploitability_calibration_active"
	IntelligenceCalibrationValAExploitabilityStatePartial    = "intelligence_calibration_vala_exploitability_calibration_partial"
	IntelligenceCalibrationValAExploitabilityStateIncomplete = "intelligence_calibration_vala_exploitability_calibration_incomplete"

	IntelligenceCalibrationValADecisionStateActive     = "intelligence_calibration_vala_downgrade_escalation_active"
	IntelligenceCalibrationValADecisionStatePartial    = "intelligence_calibration_vala_downgrade_escalation_partial"
	IntelligenceCalibrationValADecisionStateIncomplete = "intelligence_calibration_vala_downgrade_escalation_incomplete"

	IntelligenceCalibrationValACAVIStateActive     = "intelligence_calibration_vala_cavi_tuning_active"
	IntelligenceCalibrationValACAVIStatePartial    = "intelligence_calibration_vala_cavi_tuning_partial"
	IntelligenceCalibrationValACAVIStateIncomplete = "intelligence_calibration_vala_cavi_tuning_incomplete"

	IntelligenceCalibrationValAVEXCandidateStateActive     = "intelligence_calibration_vala_vex_candidate_active"
	IntelligenceCalibrationValAVEXCandidateStatePartial    = "intelligence_calibration_vala_vex_candidate_partial"
	IntelligenceCalibrationValAVEXCandidateStateIncomplete = "intelligence_calibration_vala_vex_candidate_incomplete"

	IntelligenceCalibrationValAVEXSufficiencyStateActive     = "intelligence_calibration_vala_vex_sufficiency_active"
	IntelligenceCalibrationValAVEXSufficiencyStatePartial    = "intelligence_calibration_vala_vex_sufficiency_partial"
	IntelligenceCalibrationValAVEXSufficiencyStateIncomplete = "intelligence_calibration_vala_vex_sufficiency_incomplete"

	IntelligenceCalibrationValAExplanationStateActive     = "intelligence_calibration_vala_explanation_active"
	IntelligenceCalibrationValAExplanationStatePartial    = "intelligence_calibration_vala_explanation_partial"
	IntelligenceCalibrationValAExplanationStateIncomplete = "intelligence_calibration_vala_explanation_incomplete"

	IntelligenceCalibrationValAConfidenceOutcomeStateActive     = "intelligence_calibration_vala_confidence_outcome_active"
	IntelligenceCalibrationValAConfidenceOutcomeStatePartial    = "intelligence_calibration_vala_confidence_outcome_partial"
	IntelligenceCalibrationValAConfidenceOutcomeStateIncomplete = "intelligence_calibration_vala_confidence_outcome_incomplete"

	IntelligenceCalibrationValAPublicationGuardrailStateActive     = "intelligence_calibration_vala_publication_guardrail_active"
	IntelligenceCalibrationValAPublicationGuardrailStatePartial    = "intelligence_calibration_vala_publication_guardrail_partial"
	IntelligenceCalibrationValAPublicationGuardrailStateIncomplete = "intelligence_calibration_vala_publication_guardrail_incomplete"

	IntelligenceCalibrationValAStateIncomplete  = "intelligence_calibration_vala_incomplete"
	IntelligenceCalibrationValAStateSubstantial = "intelligence_calibration_vala_substantially_ready"
	IntelligenceCalibrationValAStateActive      = "intelligence_calibration_vala_active"

	IntelligenceCalibrationValAAggregationComplete          = "complete"
	IntelligenceCalibrationValAAggregationPartial           = "partial"
	IntelligenceCalibrationValAAggregationStale             = "stale"
	IntelligenceCalibrationValAAggregationUnsupported       = "unsupported"
	IntelligenceCalibrationValAAggregationInsufficient      = "insufficient_evidence"
	IntelligenceCalibrationValAExploitabilityHighRisk       = "high_risk_relevant"
	IntelligenceCalibrationValAExploitabilityPotential      = "potentially_relevant"
	IntelligenceCalibrationValAExploitabilityLowEvidence    = "low_evidence_relevant"
	IntelligenceCalibrationValAExploitabilityNotEvidenced   = "currently_not_evidenced"
	IntelligenceCalibrationValAExploitabilityUnsupported    = "unsupported"
	IntelligenceCalibrationValAExploitabilityRequiresReview = "requires_review"
	IntelligenceCalibrationValADecisionDowngrade            = "downgrade"
	IntelligenceCalibrationValADecisionEscalation           = "escalation"
	IntelligenceCalibrationValADecisionNoChange             = "no_change"
	IntelligenceCalibrationValADecisionRequiresReview       = "requires_review"
	IntelligenceCalibrationValAExecutionRuntimeObserved     = "runtime_observed"
	IntelligenceCalibrationValAExecutionStaticallyReachable = "statically_reachable"
	IntelligenceCalibrationValAExecutionPresentOnly         = "present_only"
	IntelligenceCalibrationValAExecutionUnsupported         = "unsupported"
	IntelligenceCalibrationValAContextAssetSpecific         = "asset_specific"
	IntelligenceCalibrationValAContextWorkloadSpecific      = "workload_specific"
	IntelligenceCalibrationValAContextEnvironmentSpecific   = "environment_specific"
	IntelligenceCalibrationValAContextGeneric               = "generic"
	IntelligenceCalibrationValATuningEscalate               = "escalate"
	IntelligenceCalibrationValATuningDowngradeCandidate     = "downgrade_candidate"
	IntelligenceCalibrationValATuningKeepPriority           = "keep_priority"
	IntelligenceCalibrationValATuningRequiresReview         = "requires_review"
	IntelligenceCalibrationValATuningUnsupported            = "unsupported"
	IntelligenceCalibrationValAVEXSufficiencySufficient     = "sufficient_for_candidate"
	IntelligenceCalibrationValAVEXSufficiencyInsufficient   = "insufficient"
	IntelligenceCalibrationValAVEXSufficiencyStale          = "stale"
	IntelligenceCalibrationValAVEXSufficiencyUnsupported    = "unsupported"
	IntelligenceCalibrationValAVEXSufficiencyRequiresReview = "requires_review"
	IntelligenceCalibrationValAOutcomeHighConfidence        = "high_confidence_relevant"
	IntelligenceCalibrationValAOutcomeMediumConfidence      = "medium_confidence_relevant"
	IntelligenceCalibrationValAOutcomeLowConfidence         = "low_confidence_relevant"
	IntelligenceCalibrationValAOutcomeInsufficient          = "insufficient_evidence"
	IntelligenceCalibrationValAOutcomeUnsupported           = "unsupported"
	IntelligenceCalibrationValAOutcomeRequiresReview        = "requires_review"
)

type ReachabilitySignalAggregationContract struct {
	CurrentState                  string   `json:"current_state"`
	SupportedAggregationStates    []string `json:"supported_aggregation_states,omitempty"`
	AggregationID                 string   `json:"aggregation_id"`
	VulnerabilityRef              string   `json:"vulnerability_ref"`
	AssetOrProductRef             string   `json:"asset_or_product_ref"`
	PackageRef                    string   `json:"package_ref"`
	ComponentOrFunctionRef        string   `json:"component_or_function_ref"`
	WorkloadContextRef            string   `json:"workload_context_ref"`
	SignalClasses                 []string `json:"signal_classes,omitempty"`
	StaticSignalRefs              []string `json:"static_signal_refs,omitempty"`
	RuntimeSignalRefs             []string `json:"runtime_signal_refs,omitempty"`
	CallPathSignalRefs            []string `json:"call_path_signal_refs,omitempty"`
	ConfigContextRefs             []string `json:"config_context_refs,omitempty"`
	EvidenceRefs                  []string `json:"evidence_refs,omitempty"`
	FreshnessState                string   `json:"freshness_state"`
	AggregationState              string   `json:"aggregation_state"`
	LimitationMessage             string   `json:"limitation_message"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	PackagePresenceImpliesExploit bool     `json:"package_presence_implies_exploitable_reachability"`
	RuntimeLoadedImpliesExecution bool     `json:"runtime_loaded_implies_vulnerable_execution"`
	PartialTreatedComplete        bool     `json:"partial_treated_as_complete"`
	AdvisoryOnly                  bool     `json:"advisory_only"`
}

type ContextualExploitabilityCalibrationContract struct {
	CurrentState                string   `json:"current_state"`
	SupportedExploitability     []string `json:"supported_exploitability_states,omitempty"`
	CalibrationID               string   `json:"calibration_id"`
	VulnerabilityRef            string   `json:"vulnerability_ref"`
	AssetOrProductRef           string   `json:"asset_or_product_ref"`
	ReachabilityRef             string   `json:"reachability_ref"`
	ExploitabilityState         string   `json:"exploitability_state"`
	ConfidenceBand              string   `json:"confidence_band"`
	EvidenceClass               string   `json:"evidence_class"`
	DowngradeReasonCodes        []string `json:"downgrade_reason_codes,omitempty"`
	EscalationReasonCodes       []string `json:"escalation_reason_codes,omitempty"`
	Explanation                 string   `json:"explanation"`
	UncertaintyNote             string   `json:"uncertainty_note"`
	LocalContextRefs            []string `json:"local_context_refs,omitempty"`
	LimitationMessage           string   `json:"limitation_message"`
	ReviewerRequired            bool     `json:"reviewer_required"`
	AdvisoryOnly                bool     `json:"advisory_only"`
	CurrentlyNotEvidencedIsSafe bool     `json:"currently_not_evidenced_means_safe"`
	LowEvidenceBecomesSafe      bool     `json:"low_evidence_becomes_not_affected"`
	UnsupportedBecomesLowRisk   bool     `json:"unsupported_becomes_low_risk"`
}

type DowngradeEscalationDisciplineContract struct {
	CurrentState                  string   `json:"current_state"`
	SupportedDecisionTypes        []string `json:"supported_decision_types,omitempty"`
	DecisionID                    string   `json:"decision_id"`
	DecisionType                  string   `json:"decision_type"`
	EvidenceClass                 string   `json:"evidence_class"`
	ConfidenceBand                string   `json:"confidence_band"`
	ReasonCodes                   []string `json:"reason_codes,omitempty"`
	Explanation                   string   `json:"explanation"`
	AffectedSubjects              []string `json:"affected_subjects,omitempty"`
	ExcludedCriticalClasses       []string `json:"excluded_critical_classes,omitempty"`
	ReviewerRequired              bool     `json:"reviewer_required"`
	ExpiresAt                     string   `json:"expires_at,omitempty"`
	RollbackRef                   string   `json:"rollback_ref"`
	EvidenceRefs                  []string `json:"evidence_refs,omitempty"`
	LimitationMessage             string   `json:"limitation_message"`
	AppliesToExcludedCritical     bool     `json:"applies_to_excluded_critical_class"`
	MutatesCanonicalPriority      bool     `json:"mutates_canonical_priority"`
	RequiresReviewTreatedApproved bool     `json:"requires_review_treated_as_approved"`
}

type CAVIReachabilityTuningContract struct {
	CurrentState                string   `json:"current_state"`
	SupportedExecutionContexts  []string `json:"supported_execution_contexts,omitempty"`
	SupportedContextSensitivity []string `json:"supported_context_sensitivity,omitempty"`
	SupportedRecommendations    []string `json:"supported_tuning_recommendations,omitempty"`
	CAVIProfileID               string   `json:"cavi_profile_id"`
	VulnerabilityRef            string   `json:"vulnerability_ref"`
	CallPathEvidenceState       string   `json:"call_path_evidence_state"`
	ExecutionContextState       string   `json:"execution_context_state"`
	ContextSensitivity          string   `json:"context_sensitivity"`
	PackageToFunctionLinkage    bool     `json:"package_to_function_linkage_present"`
	ExploitPreconditionsKnown   bool     `json:"exploit_preconditions_known"`
	ConfidenceBand              string   `json:"confidence_band"`
	TuningRecommendation        string   `json:"tuning_recommendation"`
	ExplanationRequired         bool     `json:"explanation_required"`
	Explanation                 string   `json:"explanation"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	LimitationMessage           string   `json:"limitation_message"`
	AdvisoryOnly                bool     `json:"advisory_only"`
}

type VEXCandidateCalibrationContract struct {
	CurrentState             string   `json:"current_state"`
	SupportedOutcomes        []string `json:"supported_vex_outcomes,omitempty"`
	SupportedStates          []string `json:"supported_candidate_states,omitempty"`
	SupportedSufficiency     []string `json:"supported_sufficiency_states,omitempty"`
	CandidateID              string   `json:"candidate_id"`
	VulnerabilityRef         string   `json:"vulnerability_ref"`
	ProductOrAssetRef        string   `json:"product_or_asset_ref"`
	ReachabilityRef          string   `json:"reachability_ref"`
	ExploitabilityRef        string   `json:"exploitability_ref"`
	SuggestedVEXOutcome      string   `json:"suggested_vex_outcome"`
	CandidateState           string   `json:"candidate_state"`
	ConfidenceBand           string   `json:"confidence_band"`
	EvidenceSufficiencyState string   `json:"evidence_sufficiency_state"`
	ReviewerRequired         bool     `json:"reviewer_required"`
	PublicationAllowed       bool     `json:"publication_allowed"`
	FinalVEXClaim            bool     `json:"final_vex_claim"`
	Expiry                   string   `json:"expiry"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	ReasonCode               string   `json:"reason_code"`
	LimitationMessage        string   `json:"limitation_message"`
	AdvisoryOnly             bool     `json:"advisory_only"`
}

type VEXEvidenceSufficiencyContract struct {
	CurrentState            string   `json:"current_state"`
	SupportedSufficiency    []string `json:"supported_sufficiency_states,omitempty"`
	SufficiencyCheckID      string   `json:"sufficiency_check_id"`
	CandidateRef            string   `json:"candidate_ref"`
	RequiredEvidenceClasses []string `json:"required_evidence_classes,omitempty"`
	PresentEvidenceClasses  []string `json:"present_evidence_classes,omitempty"`
	MissingEvidenceClasses  []string `json:"missing_evidence_classes,omitempty"`
	StaleEvidenceRefs       []string `json:"stale_evidence_refs,omitempty"`
	UnsupportedEvidenceRefs []string `json:"unsupported_evidence_refs,omitempty"`
	SufficiencyState        string   `json:"sufficiency_state"`
	LimitationMessage       string   `json:"limitation_message"`
	FinalPublicationImplied bool     `json:"final_publication_implied"`
}

type ReachabilityVEXExplanationContract struct {
	CurrentState                         string   `json:"current_state"`
	ReasonCode                           string   `json:"reason_code"`
	HumanMessage                         string   `json:"human_message"`
	TechnicalDetail                      string   `json:"technical_detail"`
	VulnerabilityRef                     string   `json:"vulnerability_ref"`
	AssetOrProductRef                    string   `json:"asset_or_product_ref"`
	EvidenceRefs                         []string `json:"evidence_refs,omitempty"`
	ConfidenceBand                       string   `json:"confidence_band"`
	EvidenceClass                        string   `json:"evidence_class"`
	UncertaintyNote                      string   `json:"uncertainty_note"`
	NextStep                             string   `json:"next_step"`
	ReviewerRequired                     bool     `json:"reviewer_required"`
	VisibilityScope                      string   `json:"visibility_scope"`
	RedactionTier                        string   `json:"redaction_tier"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
	DistinguishesNotEvidencedFromSafe    bool     `json:"distinguishes_not_evidenced_from_safe"`
	RedactionTurnsInsufficientSufficient bool     `json:"redaction_turns_insufficient_into_sufficient"`
	LeaksInternalEvidence                bool     `json:"leaks_internal_evidence"`
}

type ConfidenceBoundReachabilityOutcomeContract struct {
	CurrentState        string   `json:"current_state"`
	SupportedOutcomes   []string `json:"supported_outcome_states,omitempty"`
	OutcomeID           string   `json:"outcome_id"`
	ReachabilityRef     string   `json:"reachability_ref"`
	ExploitabilityRef   string   `json:"exploitability_ref"`
	ConfidenceBand      string   `json:"confidence_band"`
	EvidenceClass       string   `json:"evidence_class"`
	OutcomeState        string   `json:"outcome_state"`
	ConfidenceCapReason string   `json:"confidence_cap_reason"`
	FreshnessState      string   `json:"freshness_state"`
	LimitationMessage   string   `json:"limitation_message"`
	AdvisoryOnly        bool     `json:"advisory_only"`
}

type NoFinalPublicationVEXGuardrailContract struct {
	CurrentState                     string   `json:"current_state"`
	GuardrailID                      string   `json:"guardrail_id"`
	PublicationAllowed               bool     `json:"publication_allowed"`
	FinalClaimBlocked                bool     `json:"final_claim_blocked"`
	AllowedOutputs                   []string `json:"allowed_outputs,omitempty"`
	BlockedOutputs                   []string `json:"blocked_outputs,omitempty"`
	GovernanceRequiredForPublication bool     `json:"governance_required_for_publication"`
	LimitationMessage                string   `json:"limitation_message"`
}

func intelligenceCalibrationValAAggregationSignalClasses() []string {
	return []string{
		"package_present",
		"static_call_path",
		"runtime_loaded",
		"runtime_executed",
		"config_enabled",
		"exploit_precondition_met",
		"compensating_control_present",
		"unsupported_signal",
	}
}

func intelligenceCalibrationValAAggregationStates() []string {
	return []string{
		IntelligenceCalibrationValAAggregationComplete,
		IntelligenceCalibrationValAAggregationPartial,
		IntelligenceCalibrationValAAggregationStale,
		IntelligenceCalibrationValAAggregationUnsupported,
		IntelligenceCalibrationValAAggregationInsufficient,
	}
}

func intelligenceCalibrationValAExploitabilityStates() []string {
	return []string{
		IntelligenceCalibrationValAExploitabilityHighRisk,
		IntelligenceCalibrationValAExploitabilityPotential,
		IntelligenceCalibrationValAExploitabilityLowEvidence,
		IntelligenceCalibrationValAExploitabilityNotEvidenced,
		IntelligenceCalibrationValAExploitabilityUnsupported,
		IntelligenceCalibrationValAExploitabilityRequiresReview,
	}
}

func intelligenceCalibrationValADecisionTypes() []string {
	return []string{
		IntelligenceCalibrationValADecisionDowngrade,
		IntelligenceCalibrationValADecisionEscalation,
		IntelligenceCalibrationValADecisionNoChange,
		IntelligenceCalibrationValADecisionRequiresReview,
	}
}

func intelligenceCalibrationValAExecutionContexts() []string {
	return []string{
		IntelligenceCalibrationValAExecutionRuntimeObserved,
		IntelligenceCalibrationValAExecutionStaticallyReachable,
		IntelligenceCalibrationValAExecutionPresentOnly,
		IntelligenceCalibrationValAExecutionUnsupported,
	}
}

func intelligenceCalibrationValAContextSensitivity() []string {
	return []string{
		IntelligenceCalibrationValAContextAssetSpecific,
		IntelligenceCalibrationValAContextWorkloadSpecific,
		IntelligenceCalibrationValAContextEnvironmentSpecific,
		IntelligenceCalibrationValAContextGeneric,
	}
}

func intelligenceCalibrationValATuningRecommendations() []string {
	return []string{
		IntelligenceCalibrationValATuningEscalate,
		IntelligenceCalibrationValATuningDowngradeCandidate,
		IntelligenceCalibrationValATuningKeepPriority,
		IntelligenceCalibrationValATuningRequiresReview,
		IntelligenceCalibrationValATuningUnsupported,
	}
}

func intelligenceCalibrationValAVEXSufficiencyStates() []string {
	return []string{
		IntelligenceCalibrationValAVEXSufficiencySufficient,
		IntelligenceCalibrationValAVEXSufficiencyInsufficient,
		IntelligenceCalibrationValAVEXSufficiencyStale,
		IntelligenceCalibrationValAVEXSufficiencyUnsupported,
		IntelligenceCalibrationValAVEXSufficiencyRequiresReview,
	}
}

func intelligenceCalibrationValAOutcomeStates() []string {
	return []string{
		IntelligenceCalibrationValAOutcomeHighConfidence,
		IntelligenceCalibrationValAOutcomeMediumConfidence,
		IntelligenceCalibrationValAOutcomeLowConfidence,
		IntelligenceCalibrationValAOutcomeInsufficient,
		IntelligenceCalibrationValAOutcomeUnsupported,
		IntelligenceCalibrationValAOutcomeRequiresReview,
	}
}

func IntelligenceCalibrationValAReachabilityAggregationContract() ReachabilitySignalAggregationContract {
	return ReachabilitySignalAggregationContract{
		CurrentState:                  "reachability_signal_aggregation_ready",
		SupportedAggregationStates:    intelligenceCalibrationValAAggregationStates(),
		AggregationID:                 "reachability-aggregation-1",
		VulnerabilityRef:              "CVE-2026-1001",
		AssetOrProductRef:             "cluster-a/acme-prod/Deployment/api",
		PackageRef:                    "pkg:golang/github.com/acme/api",
		ComponentOrFunctionRef:        "github.com/acme/api/internal/handler.Login",
		WorkloadContextRef:            "cluster-a/acme-prod/Deployment/api",
		SignalClasses:                 intelligenceCalibrationValAAggregationSignalClasses(),
		StaticSignalRefs:              []string{"static:callpath/login"},
		RuntimeSignalRefs:             []string{"runtime:loaded/api", "runtime:executed/login"},
		CallPathSignalRefs:            []string{"callpath:login_to_parser"},
		ConfigContextRefs:             []string{"config:auth_enabled"},
		EvidenceRefs:                  []string{"evidence:reachability-aggregation", "evidence_spine"},
		FreshnessState:                IntelligenceCalibrationFreshnessFresh,
		AggregationState:              IntelligenceCalibrationValAAggregationComplete,
		LimitationMessage:             "Reachability aggregation remains projection_only and bounded by observed signal coverage.",
		ProjectionDisclaimer:          "projection_only not_canonical_truth reachability_signal_aggregation",
		PackagePresenceImpliesExploit: false,
		RuntimeLoadedImpliesExecution: false,
		PartialTreatedComplete:        false,
		AdvisoryOnly:                  true,
	}
}

func IntelligenceCalibrationValAExploitabilityCalibrationContract() ContextualExploitabilityCalibrationContract {
	return ContextualExploitabilityCalibrationContract{
		CurrentState:                "contextual_exploitability_calibration_ready",
		SupportedExploitability:     intelligenceCalibrationValAExploitabilityStates(),
		CalibrationID:               "exploitability-calibration-1",
		VulnerabilityRef:            "CVE-2026-1001",
		AssetOrProductRef:           "cluster-a/acme-prod/Deployment/api",
		ReachabilityRef:             "reachability-aggregation-1",
		ExploitabilityState:         IntelligenceCalibrationValAExploitabilityPotential,
		ConfidenceBand:              IntelligenceCalibrationConfidenceMedium,
		EvidenceClass:               IntelligenceCalibrationEvidenceStronglyInferred,
		DowngradeReasonCodes:        []string{"compensating_control_present"},
		EscalationReasonCodes:       []string{"runtime_execution_observed"},
		Explanation:                 "Current context is potentially relevant, but not presently safe, because compensating controls and execution evidence still require reviewed interpretation.",
		UncertaintyNote:             "currently_not_evidenced never means safe and low-evidence posture stays review-bound",
		LocalContextRefs:            []string{"context:prod-auth", "control:waf-active"},
		LimitationMessage:           "Exploitability calibration remains advisory and bounded to current local context.",
		ReviewerRequired:            true,
		AdvisoryOnly:                true,
		CurrentlyNotEvidencedIsSafe: false,
		LowEvidenceBecomesSafe:      false,
		UnsupportedBecomesLowRisk:   false,
	}
}

func IntelligenceCalibrationValADowngradeEscalationContract() DowngradeEscalationDisciplineContract {
	return DowngradeEscalationDisciplineContract{
		CurrentState:                  "downgrade_escalation_guardrail_ready",
		SupportedDecisionTypes:        intelligenceCalibrationValADecisionTypes(),
		DecisionID:                    "decision-1",
		DecisionType:                  IntelligenceCalibrationValADecisionDowngrade,
		EvidenceClass:                 IntelligenceCalibrationEvidenceStronglyInferred,
		ConfidenceBand:                IntelligenceCalibrationConfidenceMedium,
		ReasonCodes:                   []string{"compensating_control_present"},
		Explanation:                   "Downgrade remains a bounded candidate because compensating controls reduce current exposure without creating a final safe claim.",
		AffectedSubjects:              []string{"cluster-a/acme-prod/Deployment/api"},
		ExcludedCriticalClasses:       []string{"active_exploitation", "critical_runtime_blockers"},
		ReviewerRequired:              true,
		ExpiresAt:                     "2026-04-30T08:00:00Z",
		RollbackRef:                   "rollback:decision-1",
		EvidenceRefs:                  []string{"control:waf-active", "review:reachability-1"},
		LimitationMessage:             "Decision output is advisory only and does not mutate canonical priority without later governance.",
		AppliesToExcludedCritical:     false,
		MutatesCanonicalPriority:      false,
		RequiresReviewTreatedApproved: false,
	}
}

func IntelligenceCalibrationValACAVITuningContract() CAVIReachabilityTuningContract {
	return CAVIReachabilityTuningContract{
		CurrentState:                "cavi_reachability_tuning_ready",
		SupportedExecutionContexts:  intelligenceCalibrationValAExecutionContexts(),
		SupportedContextSensitivity: intelligenceCalibrationValAContextSensitivity(),
		SupportedRecommendations:    intelligenceCalibrationValATuningRecommendations(),
		CAVIProfileID:               "cavi-profile-1",
		VulnerabilityRef:            "CVE-2026-1001",
		CallPathEvidenceState:       IntelligenceCalibrationEvidenceStronglyInferred,
		ExecutionContextState:       IntelligenceCalibrationValAExecutionStaticallyReachable,
		ContextSensitivity:          IntelligenceCalibrationValAContextWorkloadSpecific,
		PackageToFunctionLinkage:    true,
		ExploitPreconditionsKnown:   true,
		ConfidenceBand:              IntelligenceCalibrationConfidenceMedium,
		TuningRecommendation:        IntelligenceCalibrationValATuningKeepPriority,
		ExplanationRequired:         true,
		Explanation:                 "CAVI tuning remains advisory and reflects bounded linkage, call-path, and execution-context evidence.",
		EvidenceRefs:                []string{"callpath:login_to_parser", "runtime:loaded/api"},
		LimitationMessage:           "Tuning output stays projection_only and cannot mutate priority directly.",
		AdvisoryOnly:                true,
	}
}

func IntelligenceCalibrationValAVEXCandidateContract() VEXCandidateCalibrationContract {
	return VEXCandidateCalibrationContract{
		CurrentState:             "vex_candidate_calibration_ready",
		SupportedOutcomes:        intelligenceCalibrationVal0VEXOutcomes(),
		SupportedStates:          intelligenceCalibrationVal0VEXStates(),
		SupportedSufficiency:     intelligenceCalibrationValAVEXSufficiencyStates(),
		CandidateID:              "vex-candidate-a1",
		VulnerabilityRef:         "CVE-2026-1001",
		ProductOrAssetRef:        "cluster-a/acme-prod/Deployment/api",
		ReachabilityRef:          "reachability-aggregation-1",
		ExploitabilityRef:        "exploitability-calibration-1",
		SuggestedVEXOutcome:      IntelligenceCalibrationVEXOutcomeRequiresReview,
		CandidateState:           IntelligenceCalibrationVEXStateRequiresReview,
		ConfidenceBand:           IntelligenceCalibrationConfidenceMedium,
		EvidenceSufficiencyState: IntelligenceCalibrationValAVEXSufficiencySufficient,
		ReviewerRequired:         true,
		PublicationAllowed:       false,
		FinalVEXClaim:            false,
		Expiry:                   "2026-04-30T08:00:00Z",
		EvidenceRefs:             []string{"evidence:vex-candidate-a1", "evidence:reachability-aggregation"},
		ReasonCode:               "candidate_only_requires_governed_vex_review",
		LimitationMessage:        "Val A VEX output stays candidate-only and cannot publish final VEX truth.",
		AdvisoryOnly:             true,
	}
}

func IntelligenceCalibrationValAVEXSufficiencyContract() VEXEvidenceSufficiencyContract {
	return VEXEvidenceSufficiencyContract{
		CurrentState:            "vex_evidence_sufficiency_ready",
		SupportedSufficiency:    intelligenceCalibrationValAVEXSufficiencyStates(),
		SufficiencyCheckID:      "vex-sufficiency-1",
		CandidateRef:            "vex-candidate-a1",
		RequiredEvidenceClasses: []string{IntelligenceCalibrationEvidenceDirectlyEvidenced, IntelligenceCalibrationEvidenceStronglyInferred},
		PresentEvidenceClasses:  []string{IntelligenceCalibrationEvidenceDirectlyEvidenced, IntelligenceCalibrationEvidenceStronglyInferred},
		MissingEvidenceClasses:  nil,
		StaleEvidenceRefs:       nil,
		UnsupportedEvidenceRefs: nil,
		SufficiencyState:        IntelligenceCalibrationValAVEXSufficiencySufficient,
		LimitationMessage:       "Evidence sufficiency is bounded to candidate viability only and does not authorize final publication.",
		FinalPublicationImplied: false,
	}
}

func IntelligenceCalibrationValAExplanationContract() ReachabilityVEXExplanationContract {
	return ReachabilityVEXExplanationContract{
		CurrentState:                         "reachability_vex_explanation_ready",
		ReasonCode:                           "runtime_execution_observed_requires_review",
		HumanMessage:                         "Reachability is not currently evidenced as safe; the current output remains a reviewed calibration candidate only.",
		TechnicalDetail:                      "Observed runtime load and bounded call-path evidence support relevance review, but they do not create final VEX truth or a safe claim.",
		VulnerabilityRef:                     "CVE-2026-1001",
		AssetOrProductRef:                    "cluster-a/acme-prod/Deployment/api",
		EvidenceRefs:                         []string{"evidence:reachability-aggregation", "evidence:vex-candidate-a1"},
		ConfidenceBand:                       IntelligenceCalibrationConfidenceMedium,
		EvidenceClass:                        IntelligenceCalibrationEvidenceStronglyInferred,
		UncertaintyNote:                      "not_evidenced is distinct from safe and uncertainty remains explicit",
		NextStep:                             "request reviewed VEX candidate evaluation before any downstream publication or priority change",
		ReviewerRequired:                     true,
		VisibilityScope:                      ProductionUsabilityVisibilityOperator,
		RedactionTier:                        ProductionUsabilityRedactionLow,
		ProjectionDisclaimer:                 "projection_only not_canonical_truth reachability_vex_explanation",
		DistinguishesNotEvidencedFromSafe:    true,
		RedactionTurnsInsufficientSufficient: false,
		LeaksInternalEvidence:                false,
	}
}

func IntelligenceCalibrationValAConfidenceOutcomeContract() ConfidenceBoundReachabilityOutcomeContract {
	return ConfidenceBoundReachabilityOutcomeContract{
		CurrentState:        "confidence_bound_outcome_ready",
		SupportedOutcomes:   intelligenceCalibrationValAOutcomeStates(),
		OutcomeID:           "confidence-outcome-1",
		ReachabilityRef:     "reachability-aggregation-1",
		ExploitabilityRef:   "exploitability-calibration-1",
		ConfidenceBand:      IntelligenceCalibrationConfidenceMedium,
		EvidenceClass:       IntelligenceCalibrationEvidenceStronglyInferred,
		OutcomeState:        IntelligenceCalibrationValAOutcomeMediumConfidence,
		ConfidenceCapReason: "bounded_linkage_and_execution_context_limit_high_confidence_claims",
		FreshnessState:      IntelligenceCalibrationFreshnessFresh,
		LimitationMessage:   "Confidence-bound outcome remains advisory and freshness-bound.",
		AdvisoryOnly:        true,
	}
}

func IntelligenceCalibrationValAPublicationGuardrailContract() NoFinalPublicationVEXGuardrailContract {
	return NoFinalPublicationVEXGuardrailContract{
		CurrentState:                     "no_final_publication_guardrail_ready",
		GuardrailID:                      "vex-publication-guardrail-a1",
		PublicationAllowed:               false,
		FinalClaimBlocked:                true,
		AllowedOutputs:                   []string{"candidate_vex_output", "reviewed_candidate_projection", "bounded_internal_review_summary"},
		BlockedOutputs:                   []string{"final_vex_publication", "public_safe_final_vex_claim", "automatic_safe_claim"},
		GovernanceRequiredForPublication: true,
		LimitationMessage:                "Val A blocks final VEX publication and allows candidate-only outputs until later governed capability exists.",
	}
}

func EvaluateIntelligenceCalibrationValAAggregationState(model ReachabilitySignalAggregationContract) string {
	if strings.TrimSpace(model.AggregationID) == "" || strings.TrimSpace(model.VulnerabilityRef) == "" || strings.TrimSpace(model.AssetOrProductRef) == "" || strings.TrimSpace(model.PackageRef) == "" || strings.TrimSpace(model.ComponentOrFunctionRef) == "" || strings.TrimSpace(model.WorkloadContextRef) == "" || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.AggregationState) == "" || strings.TrimSpace(model.LimitationMessage) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValAAggregationStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedAggregationStates, intelligenceCalibrationValAAggregationStates()...) || !containsExactTrimmedStringSet(model.SignalClasses, intelligenceCalibrationValAAggregationSignalClasses()...) || len(model.StaticSignalRefs) == 0 || len(model.RuntimeSignalRefs) == 0 || len(model.CallPathSignalRefs) == 0 || len(model.ConfigContextRefs) == 0 || len(model.EvidenceRefs) == 0 {
		return IntelligenceCalibrationValAAggregationStatePartial
	}
	if !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || !containsTrimmedString(model.SupportedAggregationStates, model.AggregationState) || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") || !model.AdvisoryOnly {
		return IntelligenceCalibrationValAAggregationStatePartial
	}
	if model.PackagePresenceImpliesExploit || model.RuntimeLoadedImpliesExecution || model.PartialTreatedComplete {
		return IntelligenceCalibrationValAAggregationStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessStale && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "stale") {
		return IntelligenceCalibrationValAAggregationStatePartial
	}
	if model.AggregationState != IntelligenceCalibrationValAAggregationComplete || model.FreshnessState != IntelligenceCalibrationFreshnessFresh {
		return IntelligenceCalibrationValAAggregationStatePartial
	}
	return IntelligenceCalibrationValAAggregationStateActive
}

func EvaluateIntelligenceCalibrationValAExploitabilityState(model ContextualExploitabilityCalibrationContract) string {
	if strings.TrimSpace(model.CalibrationID) == "" || strings.TrimSpace(model.VulnerabilityRef) == "" || strings.TrimSpace(model.AssetOrProductRef) == "" || strings.TrimSpace(model.ReachabilityRef) == "" || strings.TrimSpace(model.ExploitabilityState) == "" || strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.EvidenceClass) == "" || strings.TrimSpace(model.Explanation) == "" || strings.TrimSpace(model.UncertaintyNote) == "" || strings.TrimSpace(model.LimitationMessage) == "" || len(model.LocalContextRefs) == 0 {
		return IntelligenceCalibrationValAExploitabilityStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedExploitability, intelligenceCalibrationValAExploitabilityStates()...) || !containsTrimmedString(model.SupportedExploitability, model.ExploitabilityState) || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceBand) || !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), model.EvidenceClass) || !model.AdvisoryOnly {
		return IntelligenceCalibrationValAExploitabilityStatePartial
	}
	if model.CurrentlyNotEvidencedIsSafe || model.LowEvidenceBecomesSafe || model.UnsupportedBecomesLowRisk || !model.ReviewerRequired {
		return IntelligenceCalibrationValAExploitabilityStatePartial
	}
	if model.ExploitabilityState == IntelligenceCalibrationValAExploitabilityHighRisk && (len(model.EscalationReasonCodes) == 0 || len(model.LocalContextRefs) == 0) {
		return IntelligenceCalibrationValAExploitabilityStatePartial
	}
	if model.ExploitabilityState == IntelligenceCalibrationValAExploitabilityUnsupported && model.ConfidenceBand == IntelligenceCalibrationConfidenceHigh {
		return IntelligenceCalibrationValAExploitabilityStatePartial
	}
	return IntelligenceCalibrationValAExploitabilityStateActive
}

func EvaluateIntelligenceCalibrationValADecisionState(model DowngradeEscalationDisciplineContract) string {
	if strings.TrimSpace(model.DecisionID) == "" || strings.TrimSpace(model.DecisionType) == "" || strings.TrimSpace(model.EvidenceClass) == "" || strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.Explanation) == "" || strings.TrimSpace(model.RollbackRef) == "" || strings.TrimSpace(model.LimitationMessage) == "" || len(model.AffectedSubjects) == 0 {
		return IntelligenceCalibrationValADecisionStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedDecisionTypes, intelligenceCalibrationValADecisionTypes()...) || !containsTrimmedString(model.SupportedDecisionTypes, model.DecisionType) || !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), model.EvidenceClass) || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceBand) {
		return IntelligenceCalibrationValADecisionStatePartial
	}
	if model.MutatesCanonicalPriority || model.RequiresReviewTreatedApproved {
		return IntelligenceCalibrationValADecisionStatePartial
	}
	switch model.DecisionType {
	case IntelligenceCalibrationValADecisionDowngrade:
		if model.EvidenceClass == IntelligenceCalibrationEvidenceUnsupported || len(model.ReasonCodes) == 0 || strings.TrimSpace(model.Explanation) == "" || strings.TrimSpace(model.RollbackRef) == "" || model.AppliesToExcludedCritical || (strings.TrimSpace(model.ExpiresAt) == "" && strings.TrimSpace(model.LimitationMessage) == "") {
			return IntelligenceCalibrationValADecisionStatePartial
		}
	case IntelligenceCalibrationValADecisionEscalation:
		if len(model.ReasonCodes) == 0 || len(model.EvidenceRefs) == 0 {
			return IntelligenceCalibrationValADecisionStatePartial
		}
	case IntelligenceCalibrationValADecisionRequiresReview:
		if !model.ReviewerRequired {
			return IntelligenceCalibrationValADecisionStatePartial
		}
	}
	return IntelligenceCalibrationValADecisionStateActive
}

func EvaluateIntelligenceCalibrationValACAVIState(model CAVIReachabilityTuningContract) string {
	if strings.TrimSpace(model.CAVIProfileID) == "" || strings.TrimSpace(model.VulnerabilityRef) == "" || strings.TrimSpace(model.CallPathEvidenceState) == "" || strings.TrimSpace(model.ExecutionContextState) == "" || strings.TrimSpace(model.ContextSensitivity) == "" || strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.TuningRecommendation) == "" || strings.TrimSpace(model.Explanation) == "" || strings.TrimSpace(model.LimitationMessage) == "" || len(model.EvidenceRefs) == 0 {
		return IntelligenceCalibrationValACAVIStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedExecutionContexts, intelligenceCalibrationValAExecutionContexts()...) || !containsExactTrimmedStringSet(model.SupportedContextSensitivity, intelligenceCalibrationValAContextSensitivity()...) || !containsExactTrimmedStringSet(model.SupportedRecommendations, intelligenceCalibrationValATuningRecommendations()...) || !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), model.CallPathEvidenceState) || !containsTrimmedString(model.SupportedExecutionContexts, model.ExecutionContextState) || !containsTrimmedString(model.SupportedContextSensitivity, model.ContextSensitivity) || !containsTrimmedString(model.SupportedRecommendations, model.TuningRecommendation) || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceBand) || !model.AdvisoryOnly {
		return IntelligenceCalibrationValACAVIStatePartial
	}
	if !model.ExplanationRequired {
		return IntelligenceCalibrationValACAVIStatePartial
	}
	if model.ExecutionContextState == IntelligenceCalibrationValAExecutionPresentOnly && model.TuningRecommendation == IntelligenceCalibrationValATuningDowngradeCandidate {
		return IntelligenceCalibrationValACAVIStatePartial
	}
	if model.CallPathEvidenceState == IntelligenceCalibrationEvidenceUnsupported && model.ConfidenceBand == IntelligenceCalibrationConfidenceHigh {
		return IntelligenceCalibrationValACAVIStatePartial
	}
	if !model.PackageToFunctionLinkage && model.ConfidenceBand == IntelligenceCalibrationConfidenceHigh {
		return IntelligenceCalibrationValACAVIStatePartial
	}
	if !model.ExploitPreconditionsKnown && model.TuningRecommendation != IntelligenceCalibrationValATuningRequiresReview {
		return IntelligenceCalibrationValACAVIStatePartial
	}
	return IntelligenceCalibrationValACAVIStateActive
}

func EvaluateIntelligenceCalibrationValAVEXCandidateState(model VEXCandidateCalibrationContract) string {
	if strings.TrimSpace(model.CandidateID) == "" || strings.TrimSpace(model.VulnerabilityRef) == "" || strings.TrimSpace(model.ProductOrAssetRef) == "" || strings.TrimSpace(model.ReachabilityRef) == "" || strings.TrimSpace(model.ExploitabilityRef) == "" || strings.TrimSpace(model.SuggestedVEXOutcome) == "" || strings.TrimSpace(model.CandidateState) == "" || strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.EvidenceSufficiencyState) == "" || strings.TrimSpace(model.Expiry) == "" || strings.TrimSpace(model.ReasonCode) == "" || strings.TrimSpace(model.LimitationMessage) == "" || len(model.EvidenceRefs) == 0 {
		return IntelligenceCalibrationValAVEXCandidateStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedOutcomes, intelligenceCalibrationVal0VEXOutcomes()...) || !containsExactTrimmedStringSet(model.SupportedStates, intelligenceCalibrationVal0VEXStates()...) || !containsExactTrimmedStringSet(model.SupportedSufficiency, intelligenceCalibrationValAVEXSufficiencyStates()...) || !containsTrimmedString(model.SupportedOutcomes, model.SuggestedVEXOutcome) || !containsTrimmedString(model.SupportedStates, model.CandidateState) || !containsTrimmedString(model.SupportedSufficiency, model.EvidenceSufficiencyState) || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceBand) || !model.AdvisoryOnly {
		return IntelligenceCalibrationValAVEXCandidateStatePartial
	}
	if model.FinalVEXClaim || model.PublicationAllowed || !model.ReviewerRequired || model.CandidateState == IntelligenceCalibrationVEXStateExpired {
		return IntelligenceCalibrationValAVEXCandidateStatePartial
	}
	if model.SuggestedVEXOutcome == IntelligenceCalibrationVEXOutcomeNotAffectedCandidate && model.EvidenceSufficiencyState != IntelligenceCalibrationValAVEXSufficiencySufficient {
		return IntelligenceCalibrationValAVEXCandidateStatePartial
	}
	if (model.EvidenceSufficiencyState == IntelligenceCalibrationValAVEXSufficiencyStale || model.EvidenceSufficiencyState == IntelligenceCalibrationValAVEXSufficiencyUnsupported) && model.CandidateState == IntelligenceCalibrationVEXStateReviewed {
		return IntelligenceCalibrationValAVEXCandidateStatePartial
	}
	return IntelligenceCalibrationValAVEXCandidateStateActive
}

func EvaluateIntelligenceCalibrationValAVEXSufficiencyState(model VEXEvidenceSufficiencyContract) string {
	if strings.TrimSpace(model.SufficiencyCheckID) == "" || strings.TrimSpace(model.CandidateRef) == "" || strings.TrimSpace(model.SufficiencyState) == "" || strings.TrimSpace(model.LimitationMessage) == "" || len(model.RequiredEvidenceClasses) == 0 || len(model.PresentEvidenceClasses) == 0 {
		return IntelligenceCalibrationValAVEXSufficiencyStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedSufficiency, intelligenceCalibrationValAVEXSufficiencyStates()...) || !containsTrimmedString(model.SupportedSufficiency, model.SufficiencyState) || model.FinalPublicationImplied {
		return IntelligenceCalibrationValAVEXSufficiencyStatePartial
	}
	for _, required := range model.RequiredEvidenceClasses {
		if !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), required) || !containsTrimmedString(model.PresentEvidenceClasses, required) {
			return IntelligenceCalibrationValAVEXSufficiencyStatePartial
		}
	}
	if len(model.MissingEvidenceClasses) > 0 {
		return IntelligenceCalibrationValAVEXSufficiencyStatePartial
	}
	if model.SufficiencyState == IntelligenceCalibrationValAVEXSufficiencySufficient && (len(model.StaleEvidenceRefs) > 0 || len(model.UnsupportedEvidenceRefs) > 0) {
		return IntelligenceCalibrationValAVEXSufficiencyStatePartial
	}
	if model.SufficiencyState == IntelligenceCalibrationValAVEXSufficiencyStale && len(model.StaleEvidenceRefs) == 0 {
		return IntelligenceCalibrationValAVEXSufficiencyStatePartial
	}
	if model.SufficiencyState == IntelligenceCalibrationValAVEXSufficiencyUnsupported && len(model.UnsupportedEvidenceRefs) == 0 {
		return IntelligenceCalibrationValAVEXSufficiencyStatePartial
	}
	if model.SufficiencyState != IntelligenceCalibrationValAVEXSufficiencySufficient {
		return IntelligenceCalibrationValAVEXSufficiencyStatePartial
	}
	return IntelligenceCalibrationValAVEXSufficiencyStateActive
}

func EvaluateIntelligenceCalibrationValAExplanationState(model ReachabilityVEXExplanationContract) string {
	if strings.TrimSpace(model.ReasonCode) == "" || strings.TrimSpace(model.HumanMessage) == "" || strings.TrimSpace(model.TechnicalDetail) == "" || strings.TrimSpace(model.VulnerabilityRef) == "" || strings.TrimSpace(model.AssetOrProductRef) == "" || len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.EvidenceClass) == "" || strings.TrimSpace(model.UncertaintyNote) == "" || strings.TrimSpace(model.NextStep) == "" || strings.TrimSpace(model.VisibilityScope) == "" || strings.TrimSpace(model.RedactionTier) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValAExplanationStateIncomplete
	}
	if !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceBand) || !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), model.EvidenceClass) || !containsTrimmedString(productionUsabilityValAExplainScopes(), model.VisibilityScope) || !containsTrimmedString(ProductionUsabilityVal0ExplainabilityContract().SupportedRedactionTiers, model.RedactionTier) || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValAExplanationStatePartial
	}
	if !model.ReviewerRequired || !model.DistinguishesNotEvidencedFromSafe || model.RedactionTurnsInsufficientSufficient || model.LeaksInternalEvidence {
		return IntelligenceCalibrationValAExplanationStatePartial
	}
	if model.VisibilityScope == ProductionUsabilityVisibilityPublicSafe && model.RedactionTier != ProductionUsabilityRedactionPublicSafe {
		return IntelligenceCalibrationValAExplanationStatePartial
	}
	if model.VisibilityScope == ProductionUsabilityVisibilityPartner && (model.RedactionTier == ProductionUsabilityRedactionNone || model.RedactionTier == ProductionUsabilityRedactionLow) {
		return IntelligenceCalibrationValAExplanationStatePartial
	}
	return IntelligenceCalibrationValAExplanationStateActive
}

func EvaluateIntelligenceCalibrationValAConfidenceOutcomeState(model ConfidenceBoundReachabilityOutcomeContract) string {
	if strings.TrimSpace(model.OutcomeID) == "" || strings.TrimSpace(model.ReachabilityRef) == "" || strings.TrimSpace(model.ExploitabilityRef) == "" || strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.EvidenceClass) == "" || strings.TrimSpace(model.OutcomeState) == "" || strings.TrimSpace(model.ConfidenceCapReason) == "" || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValAConfidenceOutcomeStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedOutcomes, intelligenceCalibrationValAOutcomeStates()...) || !containsTrimmedString(model.SupportedOutcomes, model.OutcomeState) || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceBand) || !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), model.EvidenceClass) || !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || !model.AdvisoryOnly {
		return IntelligenceCalibrationValAConfidenceOutcomeStatePartial
	}
	if model.EvidenceClass == IntelligenceCalibrationEvidenceUnsupported && model.OutcomeState != IntelligenceCalibrationValAOutcomeUnsupported {
		return IntelligenceCalibrationValAConfidenceOutcomeStatePartial
	}
	if model.EvidenceClass == IntelligenceCalibrationEvidenceWeaklyInferred && model.OutcomeState == IntelligenceCalibrationValAOutcomeHighConfidence {
		return IntelligenceCalibrationValAConfidenceOutcomeStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessStale && (model.ConfidenceBand == IntelligenceCalibrationConfidenceHigh || !strings.Contains(strings.TrimSpace(model.LimitationMessage), "stale")) {
		return IntelligenceCalibrationValAConfidenceOutcomeStatePartial
	}
	if model.OutcomeState == IntelligenceCalibrationValAOutcomeInsufficient && model.ConfidenceBand == IntelligenceCalibrationConfidenceHigh {
		return IntelligenceCalibrationValAConfidenceOutcomeStatePartial
	}
	return IntelligenceCalibrationValAConfidenceOutcomeStateActive
}

func EvaluateIntelligenceCalibrationValAPublicationGuardrailState(model NoFinalPublicationVEXGuardrailContract) string {
	if strings.TrimSpace(model.GuardrailID) == "" || strings.TrimSpace(model.LimitationMessage) == "" || len(model.AllowedOutputs) == 0 || len(model.BlockedOutputs) == 0 {
		return IntelligenceCalibrationValAPublicationGuardrailStateIncomplete
	}
	if model.PublicationAllowed || !model.FinalClaimBlocked || !model.GovernanceRequiredForPublication {
		return IntelligenceCalibrationValAPublicationGuardrailStatePartial
	}
	if !containsTrimmedString(model.AllowedOutputs, "candidate_vex_output") || !containsTrimmedString(model.BlockedOutputs, "final_vex_publication") {
		return IntelligenceCalibrationValAPublicationGuardrailStatePartial
	}
	return IntelligenceCalibrationValAPublicationGuardrailStateActive
}

func EvaluateIntelligenceCalibrationValAState(val0DependencyState, val0FoundationState, aggregationState, exploitabilityState, decisionState, caviState, vexCandidateState, vexSufficiencyState, explanationState, confidenceOutcomeState, publicationGuardrailState string) string {
	if strings.TrimSpace(val0DependencyState) != IntelligenceCalibrationVal0StateActive || strings.TrimSpace(val0FoundationState) != IntelligenceCalibrationVal0StateActive {
		return IntelligenceCalibrationValAStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(aggregationState),
		strings.TrimSpace(exploitabilityState),
		strings.TrimSpace(decisionState),
		strings.TrimSpace(caviState),
		strings.TrimSpace(vexCandidateState),
		strings.TrimSpace(vexSufficiencyState),
		strings.TrimSpace(explanationState),
		strings.TrimSpace(confidenceOutcomeState),
		strings.TrimSpace(publicationGuardrailState),
	} {
		switch state {
		case IntelligenceCalibrationValAAggregationStateActive,
			IntelligenceCalibrationValAExploitabilityStateActive,
			IntelligenceCalibrationValADecisionStateActive,
			IntelligenceCalibrationValACAVIStateActive,
			IntelligenceCalibrationValAVEXCandidateStateActive,
			IntelligenceCalibrationValAVEXSufficiencyStateActive,
			IntelligenceCalibrationValAExplanationStateActive,
			IntelligenceCalibrationValAConfidenceOutcomeStateActive,
			IntelligenceCalibrationValAPublicationGuardrailStateActive:
		case IntelligenceCalibrationValAAggregationStatePartial,
			IntelligenceCalibrationValAExploitabilityStatePartial,
			IntelligenceCalibrationValADecisionStatePartial,
			IntelligenceCalibrationValACAVIStatePartial,
			IntelligenceCalibrationValAVEXCandidateStatePartial,
			IntelligenceCalibrationValAVEXSufficiencyStatePartial,
			IntelligenceCalibrationValAExplanationStatePartial,
			IntelligenceCalibrationValAConfidenceOutcomeStatePartial,
			IntelligenceCalibrationValAPublicationGuardrailStatePartial:
			hasPartial = true
		default:
			return IntelligenceCalibrationValAStateIncomplete
		}
	}
	if hasPartial {
		return IntelligenceCalibrationValAStateSubstantial
	}
	return IntelligenceCalibrationValAStateActive
}

func EvaluateIntelligenceCalibrationValAProofsState(val0DependencyState, val0FoundationState, aggregationState, exploitabilityState, decisionState, caviState, vexCandidateState, vexSufficiencyState, explanationState, confidenceOutcomeState, publicationGuardrailState string, surfaceRefs, evidenceRefs, limitations, whyPoint5NotPass []string, projectionDisclaimer string) string {
	baseState := EvaluateIntelligenceCalibrationValAState(
		val0DependencyState,
		val0FoundationState,
		aggregationState,
		exploitabilityState,
		decisionState,
		caviState,
		vexCandidateState,
		vexSufficiencyState,
		explanationState,
		confidenceOutcomeState,
		publicationGuardrailState,
	)
	if len(surfaceRefs) < 10 || len(evidenceRefs) < 8 || len(limitations) == 0 || len(whyPoint5NotPass) == 0 || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "not_canonical_truth") {
		if baseState == IntelligenceCalibrationValAStateActive {
			return IntelligenceCalibrationValAStateSubstantial
		}
		return baseState
	}
	return baseState
}

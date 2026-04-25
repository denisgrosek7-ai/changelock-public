package operability

import "strings"

const (
	IntelligenceCalibrationValCFeedbackIntakeStateActive     = "intelligence_calibration_valc_feedback_intake_active"
	IntelligenceCalibrationValCFeedbackIntakeStatePartial    = "intelligence_calibration_valc_feedback_intake_partial"
	IntelligenceCalibrationValCFeedbackIntakeStateIncomplete = "intelligence_calibration_valc_feedback_intake_incomplete"

	IntelligenceCalibrationValCReviewCockpitStateActive     = "intelligence_calibration_valc_feedback_review_cockpit_active"
	IntelligenceCalibrationValCReviewCockpitStatePartial    = "intelligence_calibration_valc_feedback_review_cockpit_partial"
	IntelligenceCalibrationValCReviewCockpitStateIncomplete = "intelligence_calibration_valc_feedback_review_cockpit_incomplete"

	IntelligenceCalibrationValCTuningProposalStateActive     = "intelligence_calibration_valc_tuning_proposal_active"
	IntelligenceCalibrationValCTuningProposalStatePartial    = "intelligence_calibration_valc_tuning_proposal_partial"
	IntelligenceCalibrationValCTuningProposalStateIncomplete = "intelligence_calibration_valc_tuning_proposal_incomplete"

	IntelligenceCalibrationValCSuppressionSafetyStateActive     = "intelligence_calibration_valc_suppression_safety_active"
	IntelligenceCalibrationValCSuppressionSafetyStatePartial    = "intelligence_calibration_valc_suppression_safety_partial"
	IntelligenceCalibrationValCSuppressionSafetyStateIncomplete = "intelligence_calibration_valc_suppression_safety_incomplete"

	IntelligenceCalibrationValCSuppressionRollbackStateActive     = "intelligence_calibration_valc_suppression_rollback_active"
	IntelligenceCalibrationValCSuppressionRollbackStatePartial    = "intelligence_calibration_valc_suppression_rollback_partial"
	IntelligenceCalibrationValCSuppressionRollbackStateIncomplete = "intelligence_calibration_valc_suppression_rollback_incomplete"

	IntelligenceCalibrationValCLocalChangeReviewStateActive     = "intelligence_calibration_valc_local_change_review_active"
	IntelligenceCalibrationValCLocalChangeReviewStatePartial    = "intelligence_calibration_valc_local_change_review_partial"
	IntelligenceCalibrationValCLocalChangeReviewStateIncomplete = "intelligence_calibration_valc_local_change_review_incomplete"

	IntelligenceCalibrationValCFederatedWeightingStateActive     = "intelligence_calibration_valc_federated_weighting_active"
	IntelligenceCalibrationValCFederatedWeightingStatePartial    = "intelligence_calibration_valc_federated_weighting_partial"
	IntelligenceCalibrationValCFederatedWeightingStateIncomplete = "intelligence_calibration_valc_federated_weighting_incomplete"

	IntelligenceCalibrationValCSimilarityGatingStateActive     = "intelligence_calibration_valc_similarity_gating_active"
	IntelligenceCalibrationValCSimilarityGatingStatePartial    = "intelligence_calibration_valc_similarity_gating_partial"
	IntelligenceCalibrationValCSimilarityGatingStateIncomplete = "intelligence_calibration_valc_similarity_gating_incomplete"

	IntelligenceCalibrationValCLocalOverrideStateActive     = "intelligence_calibration_valc_local_override_active"
	IntelligenceCalibrationValCLocalOverrideStatePartial    = "intelligence_calibration_valc_local_override_partial"
	IntelligenceCalibrationValCLocalOverrideStateIncomplete = "intelligence_calibration_valc_local_override_incomplete"

	IntelligenceCalibrationValCPropagationPolicyStateActive     = "intelligence_calibration_valc_propagation_policy_active"
	IntelligenceCalibrationValCPropagationPolicyStatePartial    = "intelligence_calibration_valc_propagation_policy_partial"
	IntelligenceCalibrationValCPropagationPolicyStateIncomplete = "intelligence_calibration_valc_propagation_policy_incomplete"

	IntelligenceCalibrationValCExplanationStateActive     = "intelligence_calibration_valc_feedback_federated_explanation_active"
	IntelligenceCalibrationValCExplanationStatePartial    = "intelligence_calibration_valc_feedback_federated_explanation_partial"
	IntelligenceCalibrationValCExplanationStateIncomplete = "intelligence_calibration_valc_feedback_federated_explanation_incomplete"

	IntelligenceCalibrationValCStateIncomplete  = "intelligence_calibration_valc_incomplete"
	IntelligenceCalibrationValCStateSubstantial = "intelligence_calibration_valc_substantially_ready"
	IntelligenceCalibrationValCStateActive      = "intelligence_calibration_valc_active"

	IntelligenceCalibrationValCLocalApplicabilityOnly             = "local_only"
	IntelligenceCalibrationValCLocalApplicabilityCandidate        = "local_candidate"
	IntelligenceCalibrationValCFederatedApplicabilityCandidate    = "federated_candidate"
	IntelligenceCalibrationValCApplicabilityUnsupported           = "unsupported"
	IntelligenceCalibrationValCTriagePending                      = "pending"
	IntelligenceCalibrationValCTriageInReview                     = "in_review"
	IntelligenceCalibrationValCTriageBlocked                      = "blocked"
	IntelligenceCalibrationValCTriageReviewed                     = "reviewed"
	IntelligenceCalibrationValCTriageUnsupported                  = "unsupported"
	IntelligenceCalibrationValCProposalThresholdAdjustment        = "threshold_adjustment"
	IntelligenceCalibrationValCProposalSuppressionCandidate       = "suppression_candidate"
	IntelligenceCalibrationValCProposalConfidenceRecalibration    = "confidence_recalibration"
	IntelligenceCalibrationValCProposalWeightingAdjustment        = "weighting_adjustment"
	IntelligenceCalibrationValCProposalExplanationUpdate          = "explanation_update"
	IntelligenceCalibrationValCProposalRequiresReview             = "requires_review"
	IntelligenceCalibrationValCFederatedSourceQualityHigh         = "high"
	IntelligenceCalibrationValCFederatedSourceQualityMedium       = "medium"
	IntelligenceCalibrationValCFederatedSourceQualityLow          = "low"
	IntelligenceCalibrationValCFederatedSourceQualityUnknown      = "unknown"
	IntelligenceCalibrationValCFederatedSourceQualityUnsupported  = "unsupported"
	IntelligenceCalibrationValCSimilarityBandHigh                 = "high"
	IntelligenceCalibrationValCSimilarityBandMedium               = "medium"
	IntelligenceCalibrationValCSimilarityBandLow                  = "low"
	IntelligenceCalibrationValCSimilarityBandUnknown              = "unknown"
	IntelligenceCalibrationValCSimilarityBandUnsupported          = "unsupported"
	IntelligenceCalibrationValCSimilarityDecisionAllowAdvisoryUse = "allow_advisory_use"
	IntelligenceCalibrationValCSimilarityDecisionCapConfidence    = "cap_confidence"
	IntelligenceCalibrationValCSimilarityDecisionRequireReview    = "require_review"
	IntelligenceCalibrationValCSimilarityDecisionReject           = "reject"
	IntelligenceCalibrationValCPropagationDisabled                = "disabled"
	IntelligenceCalibrationValCPropagationAdvisoryOnly            = "advisory_only"
	IntelligenceCalibrationValCPropagationReviewRequired          = "review_required"
	IntelligenceCalibrationValCPropagationUnsupported             = "unsupported"
)

type StructuredFeedbackIntakeContract struct {
	CurrentState             string   `json:"current_state"`
	SupportedFeedbackClasses []string `json:"supported_feedback_classes,omitempty"`
	SupportedApplicability   []string `json:"supported_local_applicability,omitempty"`
	FeedbackID               string   `json:"feedback_id"`
	ActorRef                 string   `json:"actor_ref"`
	SignalRef                string   `json:"signal_ref"`
	SignalType               string   `json:"signal_type"`
	FeedbackClass            string   `json:"feedback_class"`
	ReasonCode               string   `json:"reason_code"`
	HumanReason              string   `json:"human_reason"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	LocalApplicability       string   `json:"local_applicability"`
	ReviewRequired           bool     `json:"review_required"`
	CreatedAt                string   `json:"created_at"`
	FreshnessState           string   `json:"freshness_state"`
	LimitationMessage        string   `json:"limitation_message"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
	MutatesIntelligence      bool     `json:"mutates_intelligence"`
	SuppressesSignals        bool     `json:"suppresses_signals"`
	LowersPriority           bool     `json:"lowers_priority"`
	RoutedAsNoiseReduction   bool     `json:"routed_as_noise_reduction"`
}

type FeedbackReviewCockpitContract struct {
	CurrentState               string   `json:"current_state"`
	SupportedTriageStates      []string `json:"supported_triage_states,omitempty"`
	ReviewQueueID              string   `json:"review_queue_id"`
	FeedbackRefs               []string `json:"feedback_refs,omitempty"`
	PendingReviewCount         int      `json:"pending_review_count"`
	HighRiskFeedbackCount      int      `json:"high_risk_feedback_count"`
	FalseNegativeCount         int      `json:"false_negative_count"`
	FalsePositiveCount         int      `json:"false_positive_count"`
	ReviewerRequired           bool     `json:"reviewer_required"`
	TriageState                string   `json:"triage_state"`
	EscalationRequired         bool     `json:"escalation_required"`
	ReviewPolicyRef            string   `json:"review_policy_ref"`
	LimitationMessage          string   `json:"limitation_message"`
	ReviewedTreatedAsApplied   bool     `json:"reviewed_treated_as_applied"`
	FalseNegativeVisible       bool     `json:"false_negative_visible"`
	UnsupportedFeedbackVisible bool     `json:"unsupported_feedback_visible"`
}

type TuningProposalContract struct {
	CurrentState             string   `json:"current_state"`
	SupportedChangeTypes     []string `json:"supported_change_types,omitempty"`
	SupportedApprovalStates  []string `json:"supported_approval_states,omitempty"`
	ProposalID               string   `json:"proposal_id"`
	SourceFeedbackRefs       []string `json:"source_feedback_refs,omitempty"`
	AffectedSignalClasses    []string `json:"affected_signal_classes,omitempty"`
	ProposedChangeType       string   `json:"proposed_change_type"`
	ProposedScope            string   `json:"proposed_scope"`
	ExpectedEffect           string   `json:"expected_effect"`
	FalsePositiveRiskNote    string   `json:"fp_risk_note"`
	FalseNegativeRiskNote    string   `json:"fn_risk_note"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	ReviewerRequired         bool     `json:"reviewer_required"`
	ApprovalState            string   `json:"approval_state"`
	MutatesActiveCalibration bool     `json:"mutates_active_calibration"`
	RollbackRef              string   `json:"rollback_ref"`
	LimitationMessage        string   `json:"limitation_message"`
	MayDecreaseSensitivity   bool     `json:"may_decrease_sensitivity"`
}

type SuppressionSafetyApplicationContract struct {
	CurrentState                string   `json:"current_state"`
	SuppressionCandidateID      string   `json:"suppression_candidate_id"`
	SourceFeedbackRefs          []string `json:"source_feedback_refs,omitempty"`
	SignalClass                 string   `json:"signal_class"`
	SuppressionScope            string   `json:"suppression_scope"`
	AffectedSubjects            []string `json:"affected_subjects,omitempty"`
	ExcludedCriticalClasses     []string `json:"excluded_critical_classes,omitempty"`
	ExpiresAt                   string   `json:"expires_at"`
	ReviewerRef                 string   `json:"reviewer_ref"`
	ReopenOnNewEvidence         bool     `json:"reopen_on_new_evidence"`
	StrongerEvidenceReopens     bool     `json:"stronger_evidence_reopens"`
	DeletesEvidence             bool     `json:"deletes_evidence"`
	SuppressesCriticalClass     bool     `json:"suppresses_critical_class"`
	SuppressesFalseNegativePath bool     `json:"suppresses_false_negative_path"`
	LimitationMessage           string   `json:"limitation_message"`
}

type SuppressionRollbackContract struct {
	CurrentState              string   `json:"current_state"`
	RollbackID                string   `json:"rollback_id"`
	SuppressionCandidateRef   string   `json:"suppression_candidate_ref"`
	RollbackAvailable         bool     `json:"rollback_available"`
	RollbackTriggerConditions []string `json:"rollback_trigger_conditions,omitempty"`
	RollbackSafetyCheck       string   `json:"rollback_safety_check"`
	BeforeStateRef            string   `json:"before_state_ref"`
	AfterStateRef             string   `json:"after_state_ref"`
	ReviewerRequired          bool     `json:"reviewer_required"`
	LimitationMessage         string   `json:"limitation_message"`
}

type LocalCalibrationChangeReviewContract struct {
	CurrentState             string   `json:"current_state"`
	SupportedApprovalStates  []string `json:"supported_approval_states,omitempty"`
	ChangeReviewID           string   `json:"change_review_id"`
	ProposalRef              string   `json:"proposal_ref"`
	LocalScope               string   `json:"local_scope"`
	AffectedAssets           []string `json:"affected_assets,omitempty"`
	AffectedSignalClasses    []string `json:"affected_signal_classes,omitempty"`
	ReviewerRequired         bool     `json:"reviewer_required"`
	ApprovalState            string   `json:"approval_state"`
	PreviewAvailable         bool     `json:"preview_available"`
	StagedRolloutSupported   bool     `json:"staged_rollout_supported"`
	RollbackRef              string   `json:"rollback_ref"`
	BeforeMetricSnapshotRef  string   `json:"before_metric_snapshot_ref"`
	AfterMetricSnapshotRef   string   `json:"after_metric_snapshot_ref"`
	MutatesActiveCalibration bool     `json:"mutates_active_calibration"`
	LimitationMessage        string   `json:"limitation_message"`
}

type FederatedSignalWeightingContract struct {
	CurrentState           string   `json:"current_state"`
	FederatedSignalID      string   `json:"federated_signal_id"`
	SourceRef              string   `json:"source_ref"`
	SourceTrustWeight      float64  `json:"source_trust_weight"`
	SourceQualityState     string   `json:"source_quality_state"`
	SignalClass            string   `json:"signal_class"`
	ConfidenceCap          string   `json:"confidence_cap"`
	ReasonTrace            string   `json:"reason_trace"`
	EvidenceRefs           []string `json:"evidence_refs,omitempty"`
	FreshnessState         string   `json:"freshness_state"`
	LimitationMessage      string   `json:"limitation_message"`
	ProducesLocalSafeState bool     `json:"produces_local_safe_or_not_affected_state"`
}

type EnvironmentSimilarityGatingContract struct {
	CurrentState           string   `json:"current_state"`
	SimilarityProfileID    string   `json:"similarity_profile_id"`
	LocalEnvironmentRef    string   `json:"local_environment_ref"`
	RemoteEnvironmentRef   string   `json:"remote_environment_ref"`
	SimilarityScore        float64  `json:"similarity_score"`
	SimilarityBand         string   `json:"similarity_band"`
	ComparedDimensions     []string `json:"compared_dimensions,omitempty"`
	MissingDimensions      []string `json:"missing_dimensions,omitempty"`
	GatingDecision         string   `json:"gating_decision"`
	LimitationMessage      string   `json:"limitation_message"`
	IncreasesConfidence    bool     `json:"increases_confidence"`
	OverridesLocalEvidence bool     `json:"overrides_local_evidence"`
}

type LocalOverrideDisciplineContract struct {
	CurrentState           string `json:"current_state"`
	OverridePolicyID       string `json:"override_policy_id"`
	LocalEvidenceWins      bool   `json:"local_evidence_wins"`
	LocalOverrideAllowed   bool   `json:"local_override_allowed"`
	OverrideReasonRequired bool   `json:"override_reason_required"`
	OverrideAuditRequired  bool   `json:"override_audit_required"`
	FederatedConfidenceCap string `json:"federated_confidence_cap"`
	LocalPolicyBoundaryRef string `json:"local_policy_boundary_ref"`
	LimitationMessage      string `json:"limitation_message"`
}

type BoundedPropagationPolicyContract struct {
	CurrentState               string   `json:"current_state"`
	PropagationPolicyID        string   `json:"propagation_policy_id"`
	PropagationAllowed         bool     `json:"propagation_allowed"`
	DefaultState               string   `json:"default_state"`
	AllowedPayloadClasses      []string `json:"allowed_payload_classes,omitempty"`
	BlockedPayloadClasses      []string `json:"blocked_payload_classes,omitempty"`
	RequiresRedaction          bool     `json:"requires_redaction"`
	RequiresReview             bool     `json:"requires_review"`
	MutatesRemoteCalibration   bool     `json:"mutates_remote_calibration"`
	LimitationMessage          string   `json:"limitation_message"`
	PropagatesRawLocalEvidence bool     `json:"propagates_raw_local_evidence"`
}

type FeedbackFederatedExplanationContract struct {
	CurrentState                    string   `json:"current_state"`
	ReasonCode                      string   `json:"reason_code"`
	HumanMessage                    string   `json:"human_message"`
	TechnicalDetail                 string   `json:"technical_detail"`
	FeedbackRefs                    []string `json:"feedback_refs,omitempty"`
	FederatedSignalRefs             []string `json:"federated_signal_refs,omitempty"`
	ConfidenceBand                  string   `json:"confidence_band"`
	EvidenceClass                   string   `json:"evidence_class"`
	FalsePositiveRiskNote           string   `json:"fp_risk_note"`
	FalseNegativeRiskNote           string   `json:"fn_risk_note"`
	LocalOverrideNote               string   `json:"local_override_note"`
	PropagationNote                 string   `json:"propagation_note"`
	NextStep                        string   `json:"next_step"`
	ReviewerRequired                bool     `json:"reviewer_required"`
	VisibilityScope                 string   `json:"visibility_scope"`
	RedactionTier                   string   `json:"redaction_tier"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
	LeaksInternalEvidence           bool     `json:"leaks_internal_evidence"`
	ReviewRequiredPresentedApproved bool     `json:"review_required_presented_as_approved"`
	FederatedAdvisoryOnlyExplicit   bool     `json:"federated_advisory_only_explicit"`
}

func intelligenceCalibrationValCApplicabilityStates() []string {
	return []string{
		IntelligenceCalibrationValCLocalApplicabilityOnly,
		IntelligenceCalibrationValCLocalApplicabilityCandidate,
		IntelligenceCalibrationValCFederatedApplicabilityCandidate,
		IntelligenceCalibrationValCApplicabilityUnsupported,
	}
}

func intelligenceCalibrationValCTriageStates() []string {
	return []string{
		IntelligenceCalibrationValCTriagePending,
		IntelligenceCalibrationValCTriageInReview,
		IntelligenceCalibrationValCTriageBlocked,
		IntelligenceCalibrationValCTriageReviewed,
		IntelligenceCalibrationValCTriageUnsupported,
	}
}

func intelligenceCalibrationValCProposalChangeTypes() []string {
	return []string{
		IntelligenceCalibrationValCProposalThresholdAdjustment,
		IntelligenceCalibrationValCProposalSuppressionCandidate,
		IntelligenceCalibrationValCProposalConfidenceRecalibration,
		IntelligenceCalibrationValCProposalWeightingAdjustment,
		IntelligenceCalibrationValCProposalExplanationUpdate,
		IntelligenceCalibrationValCProposalRequiresReview,
	}
}

func intelligenceCalibrationValCFederatedSourceQualityStates() []string {
	return []string{
		IntelligenceCalibrationValCFederatedSourceQualityHigh,
		IntelligenceCalibrationValCFederatedSourceQualityMedium,
		IntelligenceCalibrationValCFederatedSourceQualityLow,
		IntelligenceCalibrationValCFederatedSourceQualityUnknown,
		IntelligenceCalibrationValCFederatedSourceQualityUnsupported,
	}
}

func intelligenceCalibrationValCSimilarityBands() []string {
	return []string{
		IntelligenceCalibrationValCSimilarityBandHigh,
		IntelligenceCalibrationValCSimilarityBandMedium,
		IntelligenceCalibrationValCSimilarityBandLow,
		IntelligenceCalibrationValCSimilarityBandUnknown,
		IntelligenceCalibrationValCSimilarityBandUnsupported,
	}
}

func intelligenceCalibrationValCSimilarityDimensions() []string {
	return []string{"workload", "runtime", "dependency_graph", "config", "exposure", "controls"}
}

func intelligenceCalibrationValCSimilarityCriticalDimensions() []string {
	return []string{"workload", "runtime", "config", "exposure", "controls"}
}

func intelligenceCalibrationValCSimilarityDecisions() []string {
	return []string{
		IntelligenceCalibrationValCSimilarityDecisionAllowAdvisoryUse,
		IntelligenceCalibrationValCSimilarityDecisionCapConfidence,
		IntelligenceCalibrationValCSimilarityDecisionRequireReview,
		IntelligenceCalibrationValCSimilarityDecisionReject,
	}
}

func intelligenceCalibrationValCPropagationDefaultStates() []string {
	return []string{
		IntelligenceCalibrationValCPropagationDisabled,
		IntelligenceCalibrationValCPropagationAdvisoryOnly,
		IntelligenceCalibrationValCPropagationReviewRequired,
		IntelligenceCalibrationValCPropagationUnsupported,
	}
}

func intelligenceCalibrationValCSignalClasses() []string {
	return []string{
		"process_behavior",
		"network_behavior",
		"file_behavior",
		"runtime_behavior",
		"workload_behavior",
		"identity_or_actor_behavior",
		"package_present",
		"static_call_path",
		"runtime_loaded",
		"runtime_executed",
		"config_enabled",
		"exploit_precondition_met",
		"compensating_control_present",
		IntelligenceCalibrationValBUnsupportedSignal,
	}
}

func IntelligenceCalibrationValCStructuredFeedbackIntakeContract() StructuredFeedbackIntakeContract {
	return StructuredFeedbackIntakeContract{
		CurrentState:             IntelligenceCalibrationValCFeedbackIntakeStateActive,
		SupportedFeedbackClasses: intelligenceCalibrationVal0FeedbackClasses(),
		SupportedApplicability:   intelligenceCalibrationValCApplicabilityStates(),
		FeedbackID:               "feedback-001",
		ActorRef:                 "operator/secops-01",
		SignalRef:                "signal/runtime-behavior-001",
		SignalType:               "runtime_behavior",
		FeedbackClass:            IntelligenceCalibrationFeedbackNoisyButUseful,
		ReasonCode:               "operator_review_noise_bounded",
		HumanReason:              "Signal is locally noisy but still useful after bounded review.",
		EvidenceRefs:             []string{"evidence/feedback-001", "evidence/runtime-window-001"},
		LocalApplicability:       IntelligenceCalibrationValCLocalApplicabilityOnly,
		ReviewRequired:           true,
		CreatedAt:                "2026-04-25T09:00:00Z",
		FreshnessState:           IntelligenceCalibrationFreshnessFresh,
		LimitationMessage:        "feedback intake remains advisory and review-bounded before any local calibration proposal is considered",
		ProjectionDisclaimer:     "projection_only not_canonical_truth advisory_feedback_intake",
	}
}

func IntelligenceCalibrationValCFeedbackReviewCockpitContract() FeedbackReviewCockpitContract {
	return FeedbackReviewCockpitContract{
		CurrentState:               IntelligenceCalibrationValCReviewCockpitStateActive,
		SupportedTriageStates:      intelligenceCalibrationValCTriageStates(),
		ReviewQueueID:              "feedback-review-queue-001",
		FeedbackRefs:               []string{"feedback-001", "feedback-002", "feedback-003"},
		PendingReviewCount:         3,
		HighRiskFeedbackCount:      1,
		FalseNegativeCount:         1,
		FalsePositiveCount:         1,
		ReviewerRequired:           true,
		TriageState:                IntelligenceCalibrationValCTriageInReview,
		EscalationRequired:         true,
		ReviewPolicyRef:            "policy/feedback-review-001",
		LimitationMessage:          "review cockpit is bounded to review triage and does not apply tuning or suppression on its own",
		FalseNegativeVisible:       true,
		UnsupportedFeedbackVisible: true,
	}
}

func IntelligenceCalibrationValCTuningProposalContract() TuningProposalContract {
	return TuningProposalContract{
		CurrentState:             IntelligenceCalibrationValCTuningProposalStateActive,
		SupportedChangeTypes:     intelligenceCalibrationValCProposalChangeTypes(),
		SupportedApprovalStates:  intelligenceCalibrationVal0ApprovalStates(),
		ProposalID:               "proposal-001",
		SourceFeedbackRefs:       []string{"feedback-001"},
		AffectedSignalClasses:    []string{"runtime_behavior"},
		ProposedChangeType:       IntelligenceCalibrationValCProposalThresholdAdjustment,
		ProposedScope:            "tenant_local runtime_behavior",
		ExpectedEffect:           "Reduce locally noisy runtime-behavior false positives while preserving false-negative review visibility.",
		FalsePositiveRiskNote:    "FP volume should decrease only after review; no active threshold mutation occurs in Val C.",
		FalseNegativeRiskNote:    "FN risk remains explicit because any sensitivity decrease still requires later governance.",
		EvidenceRefs:             []string{"evidence/proposal-001"},
		ReviewerRequired:         true,
		ApprovalState:            IntelligenceCalibrationApprovalApproved,
		MutatesActiveCalibration: false,
		RollbackRef:              "rollback/proposal-001",
		LimitationMessage:        "tuning proposals remain advisory and do not mutate active calibration in Val C",
		MayDecreaseSensitivity:   true,
	}
}

func IntelligenceCalibrationValCSuppressionSafetyContract() SuppressionSafetyApplicationContract {
	return SuppressionSafetyApplicationContract{
		CurrentState:            IntelligenceCalibrationValCSuppressionSafetyStateActive,
		SuppressionCandidateID:  "suppression-candidate-001",
		SourceFeedbackRefs:      []string{"feedback-001"},
		SignalClass:             "runtime_behavior",
		SuppressionScope:        "tenant_local signal_class/runtime_behavior",
		AffectedSubjects:        []string{"tenant/acme", "workload/payments-api"},
		ExcludedCriticalClasses: []string{IntelligenceCalibrationValBCriticalityCritical, IntelligenceCalibrationValBCriticalityHigh},
		ExpiresAt:               "2026-05-02T09:00:00Z",
		ReviewerRef:             "reviewer/secops-approver-01",
		ReopenOnNewEvidence:     true,
		StrongerEvidenceReopens: true,
		LimitationMessage:       "suppression remains a candidate only in Val C and cannot hide evidence or false-negative review paths",
	}
}

func IntelligenceCalibrationValCSuppressionRollbackContract() SuppressionRollbackContract {
	return SuppressionRollbackContract{
		CurrentState:              IntelligenceCalibrationValCSuppressionRollbackStateActive,
		RollbackID:                "suppression-rollback-001",
		SuppressionCandidateRef:   "suppression-candidate-001",
		RollbackAvailable:         true,
		RollbackTriggerConditions: []string{"new_evidence", "false_negative_discovery", "governance_rejection"},
		RollbackSafetyCheck:       "verify rollback preserves evidence visibility and reopens prior candidate review context",
		BeforeStateRef:            "suppression-state/before-001",
		AfterStateRef:             "suppression-state/after-001",
		ReviewerRequired:          true,
		LimitationMessage:         "rollback path is advisory and review-bounded in Val C",
	}
}

func IntelligenceCalibrationValCLocalChangeReviewContract() LocalCalibrationChangeReviewContract {
	return LocalCalibrationChangeReviewContract{
		CurrentState:             IntelligenceCalibrationValCLocalChangeReviewStateActive,
		SupportedApprovalStates:  intelligenceCalibrationVal0ApprovalStates(),
		ChangeReviewID:           "local-change-review-001",
		ProposalRef:              "proposal-001",
		LocalScope:               "tenant_local runtime_behavior",
		AffectedAssets:           []string{"payments-api"},
		AffectedSignalClasses:    []string{"runtime_behavior"},
		ReviewerRequired:         true,
		ApprovalState:            IntelligenceCalibrationApprovalApproved,
		PreviewAvailable:         true,
		StagedRolloutSupported:   true,
		RollbackRef:              "rollback/local-change-001",
		BeforeMetricSnapshotRef:  "metrics/before-local-change-001",
		AfterMetricSnapshotRef:   "metrics/after-local-change-001",
		MutatesActiveCalibration: false,
		LimitationMessage:        "local calibration review stays advisory in Val C and does not mutate active calibration or baseline state",
	}
}

func IntelligenceCalibrationValCFederatedSignalWeightingContract() FederatedSignalWeightingContract {
	return FederatedSignalWeightingContract{
		CurrentState:       IntelligenceCalibrationValCFederatedWeightingStateActive,
		FederatedSignalID:  "federated-signal-001",
		SourceRef:          "remote/peer-tenant-01",
		SourceTrustWeight:  0.72,
		SourceQualityState: IntelligenceCalibrationValCFederatedSourceQualityMedium,
		SignalClass:        "runtime_behavior",
		ConfidenceCap:      IntelligenceCalibrationConfidenceMedium,
		ReasonTrace:        "bounded remote runtime signal is informative only after local similarity and override policy checks",
		EvidenceRefs:       []string{"evidence/federated-001"},
		FreshnessState:     IntelligenceCalibrationFreshnessFresh,
		LimitationMessage:  "federated weighting is advisory only and cannot declare local safe or not_affected state on its own",
	}
}

func IntelligenceCalibrationValCSimilarityGatingContract() EnvironmentSimilarityGatingContract {
	return EnvironmentSimilarityGatingContract{
		CurrentState:         IntelligenceCalibrationValCSimilarityGatingStateActive,
		SimilarityProfileID:  "similarity-profile-001",
		LocalEnvironmentRef:  "env/local-prod",
		RemoteEnvironmentRef: "env/peer-prod",
		SimilarityScore:      0.78,
		SimilarityBand:       IntelligenceCalibrationValCSimilarityBandMedium,
		ComparedDimensions:   intelligenceCalibrationValCSimilarityDimensions(),
		GatingDecision:       IntelligenceCalibrationValCSimilarityDecisionAllowAdvisoryUse,
		LimitationMessage:    "environment similarity only gates advisory reuse and cannot override local evidence",
	}
}

func IntelligenceCalibrationValCLocalOverrideDisciplineContract() LocalOverrideDisciplineContract {
	return LocalOverrideDisciplineContract{
		CurrentState:           IntelligenceCalibrationValCLocalOverrideStateActive,
		OverridePolicyID:       "override-policy-001",
		LocalEvidenceWins:      true,
		LocalOverrideAllowed:   true,
		OverrideReasonRequired: true,
		OverrideAuditRequired:  true,
		FederatedConfidenceCap: IntelligenceCalibrationConfidenceMedium,
		LocalPolicyBoundaryRef: "policy/local-override-boundary-001",
		LimitationMessage:      "local override discipline remains review-bounded and prevents federated bypass of local evidence or policy boundaries",
	}
}

func IntelligenceCalibrationValCPropagationPolicyContract() BoundedPropagationPolicyContract {
	return BoundedPropagationPolicyContract{
		CurrentState:          IntelligenceCalibrationValCPropagationPolicyStateActive,
		PropagationPolicyID:   "propagation-policy-001",
		DefaultState:          IntelligenceCalibrationValCPropagationAdvisoryOnly,
		AllowedPayloadClasses: []string{"redacted_feedback_summary", "redacted_similarity_hint"},
		BlockedPayloadClasses: []string{"raw_local_evidence", "unsupported_payload", "remote_calibration_mutation"},
		RequiresRedaction:     true,
		RequiresReview:        true,
		LimitationMessage:     "propagation remains advisory only in Val C and cannot mutate remote calibration or propagate raw local evidence",
	}
}

func IntelligenceCalibrationValCExplanationContract() FeedbackFederatedExplanationContract {
	return FeedbackFederatedExplanationContract{
		CurrentState:                  IntelligenceCalibrationValCExplanationStateActive,
		ReasonCode:                    "federated_feedback_review_required",
		HumanMessage:                  "Feedback and federated hints remain advisory and require local review before any governed calibration change.",
		TechnicalDetail:               "Local runtime_behavior noise reduction hints are bounded by false-negative review, local override policy, and propagation redaction rules.",
		FeedbackRefs:                  []string{"feedback-001"},
		FederatedSignalRefs:           []string{"federated-signal-001"},
		ConfidenceBand:                IntelligenceCalibrationConfidenceMedium,
		EvidenceClass:                 IntelligenceCalibrationEvidenceStronglyInferred,
		FalsePositiveRiskNote:         "FP reduction is bounded and cannot bypass false-negative review.",
		FalseNegativeRiskNote:         "FN risk remains explicit and blocks silent suppression or active threshold mutation.",
		LocalOverrideNote:             "Local evidence wins and override reason plus audit remain required.",
		PropagationNote:               "Federated signal is advisory only and any propagated payload must remain redacted and reviewed.",
		NextStep:                      "Route proposal to governed review before any active calibration change.",
		ReviewerRequired:              true,
		VisibilityScope:               ProductionUsabilityVisibilityOperator,
		RedactionTier:                 ProductionUsabilityRedactionMedium,
		ProjectionDisclaimer:          "projection_only not_canonical_truth advisory_feedback_federated_explanation",
		FederatedAdvisoryOnlyExplicit: true,
	}
}

func EvaluateIntelligenceCalibrationValCFeedbackIntakeState(model StructuredFeedbackIntakeContract) string {
	if strings.TrimSpace(model.FeedbackID) == "" || strings.TrimSpace(model.ActorRef) == "" || strings.TrimSpace(model.SignalRef) == "" || strings.TrimSpace(model.SignalType) == "" || strings.TrimSpace(model.FeedbackClass) == "" || strings.TrimSpace(model.ReasonCode) == "" || strings.TrimSpace(model.HumanReason) == "" || strings.TrimSpace(model.LocalApplicability) == "" || strings.TrimSpace(model.CreatedAt) == "" || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.LimitationMessage) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValCFeedbackIntakeStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedFeedbackClasses, intelligenceCalibrationVal0FeedbackClasses()...) || !containsExactTrimmedStringSet(model.SupportedApplicability, intelligenceCalibrationValCApplicabilityStates()...) || !containsTrimmedString(model.SupportedFeedbackClasses, model.FeedbackClass) || !containsTrimmedString(model.SupportedApplicability, model.LocalApplicability) || !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValCFeedbackIntakeStatePartial
	}
	if _, ok := parseIntelligenceCalibrationValBTimestamp(model.CreatedAt); !ok {
		return IntelligenceCalibrationValCFeedbackIntakeStatePartial
	}
	if model.MutatesIntelligence || model.SuppressesSignals || model.LowersPriority || !model.ReviewRequired {
		return IntelligenceCalibrationValCFeedbackIntakeStatePartial
	}
	if (model.FeedbackClass == IntelligenceCalibrationFeedbackFalseNegative || model.FeedbackClass == IntelligenceCalibrationFeedbackMissedSeverity) && model.RoutedAsNoiseReduction {
		return IntelligenceCalibrationValCFeedbackIntakeStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessStale && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "stale") {
		return IntelligenceCalibrationValCFeedbackIntakeStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessUnknown && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "unknown") {
		return IntelligenceCalibrationValCFeedbackIntakeStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessUnsupported && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "unsupported") {
		return IntelligenceCalibrationValCFeedbackIntakeStatePartial
	}
	return IntelligenceCalibrationValCFeedbackIntakeStateActive
}

func EvaluateIntelligenceCalibrationValCReviewCockpitState(model FeedbackReviewCockpitContract) string {
	if strings.TrimSpace(model.ReviewQueueID) == "" || len(model.FeedbackRefs) == 0 || strings.TrimSpace(model.TriageState) == "" || strings.TrimSpace(model.ReviewPolicyRef) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValCReviewCockpitStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedTriageStates, intelligenceCalibrationValCTriageStates()...) || !containsTrimmedString(model.SupportedTriageStates, model.TriageState) || !model.ReviewerRequired {
		return IntelligenceCalibrationValCReviewCockpitStatePartial
	}
	if model.PendingReviewCount < 0 || model.HighRiskFeedbackCount < 0 || model.FalseNegativeCount < 0 || model.FalsePositiveCount < 0 {
		return IntelligenceCalibrationValCReviewCockpitStatePartial
	}
	if model.ReviewedTreatedAsApplied || !model.UnsupportedFeedbackVisible {
		return IntelligenceCalibrationValCReviewCockpitStatePartial
	}
	if model.FalseNegativeCount > 0 && !model.FalseNegativeVisible {
		return IntelligenceCalibrationValCReviewCockpitStatePartial
	}
	if model.HighRiskFeedbackCount > 0 && !model.EscalationRequired {
		return IntelligenceCalibrationValCReviewCockpitStatePartial
	}
	return IntelligenceCalibrationValCReviewCockpitStateActive
}

func EvaluateIntelligenceCalibrationValCTuningProposalState(model TuningProposalContract) string {
	if strings.TrimSpace(model.ProposalID) == "" || len(model.SourceFeedbackRefs) == 0 || len(model.AffectedSignalClasses) == 0 || strings.TrimSpace(model.ProposedChangeType) == "" || strings.TrimSpace(model.ProposedScope) == "" || strings.TrimSpace(model.ExpectedEffect) == "" || strings.TrimSpace(model.FalsePositiveRiskNote) == "" || strings.TrimSpace(model.ApprovalState) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValCTuningProposalStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedChangeTypes, intelligenceCalibrationValCProposalChangeTypes()...) || !containsExactTrimmedStringSet(model.SupportedApprovalStates, intelligenceCalibrationVal0ApprovalStates()...) || !containsTrimmedString(model.SupportedChangeTypes, model.ProposedChangeType) || !containsTrimmedString(model.SupportedApprovalStates, model.ApprovalState) {
		return IntelligenceCalibrationValCTuningProposalStatePartial
	}
	for _, signalClass := range model.AffectedSignalClasses {
		if !containsTrimmedString(intelligenceCalibrationValCSignalClasses(), signalClass) {
			return IntelligenceCalibrationValCTuningProposalStatePartial
		}
	}
	if model.MutatesActiveCalibration {
		return IntelligenceCalibrationValCTuningProposalStatePartial
	}
	if model.MayDecreaseSensitivity && strings.TrimSpace(model.FalseNegativeRiskNote) == "" {
		return IntelligenceCalibrationValCTuningProposalStatePartial
	}
	if model.ProposedChangeType == IntelligenceCalibrationValCProposalSuppressionCandidate && !model.ReviewerRequired {
		return IntelligenceCalibrationValCTuningProposalStatePartial
	}
	if model.ApprovalState == IntelligenceCalibrationApprovalProposed || model.ApprovalState == IntelligenceCalibrationApprovalReviewRequired {
		return IntelligenceCalibrationValCTuningProposalStatePartial
	}
	if model.ApprovalState == IntelligenceCalibrationApprovalApproved && strings.TrimSpace(model.RollbackRef) == "" {
		return IntelligenceCalibrationValCTuningProposalStatePartial
	}
	return IntelligenceCalibrationValCTuningProposalStateActive
}

func EvaluateIntelligenceCalibrationValCSuppressionSafetyState(model SuppressionSafetyApplicationContract) string {
	if strings.TrimSpace(model.SuppressionCandidateID) == "" || len(model.SourceFeedbackRefs) == 0 || strings.TrimSpace(model.SignalClass) == "" || strings.TrimSpace(model.SuppressionScope) == "" || len(model.AffectedSubjects) == 0 || len(model.ExcludedCriticalClasses) == 0 || strings.TrimSpace(model.ExpiresAt) == "" || strings.TrimSpace(model.ReviewerRef) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValCSuppressionSafetyStateIncomplete
	}
	if !containsTrimmedString(intelligenceCalibrationValCSignalClasses(), model.SignalClass) {
		return IntelligenceCalibrationValCSuppressionSafetyStatePartial
	}
	if _, ok := parseIntelligenceCalibrationValBTimestamp(model.ExpiresAt); !ok {
		return IntelligenceCalibrationValCSuppressionSafetyStatePartial
	}
	if model.DeletesEvidence || model.SuppressesCriticalClass || model.SuppressesFalseNegativePath || !model.ReopenOnNewEvidence || !model.StrongerEvidenceReopens {
		return IntelligenceCalibrationValCSuppressionSafetyStatePartial
	}
	return IntelligenceCalibrationValCSuppressionSafetyStateActive
}

func EvaluateIntelligenceCalibrationValCSuppressionRollbackState(model SuppressionRollbackContract) string {
	if strings.TrimSpace(model.RollbackID) == "" || strings.TrimSpace(model.SuppressionCandidateRef) == "" || len(model.RollbackTriggerConditions) == 0 || strings.TrimSpace(model.RollbackSafetyCheck) == "" || strings.TrimSpace(model.BeforeStateRef) == "" || strings.TrimSpace(model.AfterStateRef) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValCSuppressionRollbackStateIncomplete
	}
	if !model.RollbackAvailable || !model.ReviewerRequired {
		return IntelligenceCalibrationValCSuppressionRollbackStatePartial
	}
	if !containsAllTrimmedStrings(model.RollbackTriggerConditions, "new_evidence", "false_negative_discovery") {
		return IntelligenceCalibrationValCSuppressionRollbackStatePartial
	}
	return IntelligenceCalibrationValCSuppressionRollbackStateActive
}

func EvaluateIntelligenceCalibrationValCLocalChangeReviewState(model LocalCalibrationChangeReviewContract) string {
	if strings.TrimSpace(model.ChangeReviewID) == "" || strings.TrimSpace(model.ProposalRef) == "" || strings.TrimSpace(model.LocalScope) == "" || len(model.AffectedAssets) == 0 || len(model.AffectedSignalClasses) == 0 || strings.TrimSpace(model.ApprovalState) == "" || strings.TrimSpace(model.BeforeMetricSnapshotRef) == "" || strings.TrimSpace(model.AfterMetricSnapshotRef) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValCLocalChangeReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedApprovalStates, intelligenceCalibrationVal0ApprovalStates()...) || !containsTrimmedString(model.SupportedApprovalStates, model.ApprovalState) || !model.ReviewerRequired || !model.PreviewAvailable {
		return IntelligenceCalibrationValCLocalChangeReviewStatePartial
	}
	for _, signalClass := range model.AffectedSignalClasses {
		if !containsTrimmedString(intelligenceCalibrationValCSignalClasses(), signalClass) {
			return IntelligenceCalibrationValCLocalChangeReviewStatePartial
		}
	}
	if model.MutatesActiveCalibration {
		return IntelligenceCalibrationValCLocalChangeReviewStatePartial
	}
	if model.ApprovalState == IntelligenceCalibrationApprovalProposed || model.ApprovalState == IntelligenceCalibrationApprovalReviewRequired {
		return IntelligenceCalibrationValCLocalChangeReviewStatePartial
	}
	if model.ApprovalState == IntelligenceCalibrationApprovalApproved && strings.TrimSpace(model.RollbackRef) == "" {
		return IntelligenceCalibrationValCLocalChangeReviewStatePartial
	}
	if !model.StagedRolloutSupported && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "staged") {
		return IntelligenceCalibrationValCLocalChangeReviewStatePartial
	}
	return IntelligenceCalibrationValCLocalChangeReviewStateActive
}

func EvaluateIntelligenceCalibrationValCFederatedWeightingState(model FederatedSignalWeightingContract) string {
	if strings.TrimSpace(model.FederatedSignalID) == "" || strings.TrimSpace(model.SourceRef) == "" || strings.TrimSpace(model.SourceQualityState) == "" || strings.TrimSpace(model.SignalClass) == "" || strings.TrimSpace(model.ConfidenceCap) == "" || strings.TrimSpace(model.ReasonTrace) == "" || len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValCFederatedWeightingStateIncomplete
	}
	if model.SourceTrustWeight <= 0 || model.SourceTrustWeight > 1 || !containsTrimmedString(intelligenceCalibrationValCFederatedSourceQualityStates(), model.SourceQualityState) || !containsTrimmedString(intelligenceCalibrationValCSignalClasses(), model.SignalClass) || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceCap) || !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) {
		return IntelligenceCalibrationValCFederatedWeightingStatePartial
	}
	if model.ProducesLocalSafeState {
		return IntelligenceCalibrationValCFederatedWeightingStatePartial
	}
	if (model.SourceQualityState == IntelligenceCalibrationValCFederatedSourceQualityUnknown || model.SourceQualityState == IntelligenceCalibrationValCFederatedSourceQualityUnsupported) && model.ConfidenceCap != IntelligenceCalibrationConfidenceLow && model.ConfidenceCap != IntelligenceCalibrationConfidenceUnknown {
		return IntelligenceCalibrationValCFederatedWeightingStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessStale && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "stale") {
		return IntelligenceCalibrationValCFederatedWeightingStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessUnknown && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "unknown") {
		return IntelligenceCalibrationValCFederatedWeightingStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessUnsupported && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "unsupported") {
		return IntelligenceCalibrationValCFederatedWeightingStatePartial
	}
	return IntelligenceCalibrationValCFederatedWeightingStateActive
}

func EvaluateIntelligenceCalibrationValCSimilarityGatingState(model EnvironmentSimilarityGatingContract) string {
	if strings.TrimSpace(model.SimilarityProfileID) == "" || strings.TrimSpace(model.LocalEnvironmentRef) == "" || strings.TrimSpace(model.RemoteEnvironmentRef) == "" || strings.TrimSpace(model.SimilarityBand) == "" || len(model.ComparedDimensions) == 0 || strings.TrimSpace(model.GatingDecision) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValCSimilarityGatingStateIncomplete
	}
	if model.SimilarityScore < 0 || model.SimilarityScore > 1 || !containsTrimmedString(intelligenceCalibrationValCSimilarityBands(), model.SimilarityBand) || !containsTrimmedString(intelligenceCalibrationValCSimilarityDecisions(), model.GatingDecision) {
		return IntelligenceCalibrationValCSimilarityGatingStatePartial
	}
	for _, dimension := range model.ComparedDimensions {
		if !containsTrimmedString(intelligenceCalibrationValCSimilarityDimensions(), dimension) {
			return IntelligenceCalibrationValCSimilarityGatingStatePartial
		}
	}
	for _, dimension := range model.MissingDimensions {
		if !containsTrimmedString(intelligenceCalibrationValCSimilarityDimensions(), dimension) {
			return IntelligenceCalibrationValCSimilarityGatingStatePartial
		}
	}
	if model.GatingDecision == IntelligenceCalibrationValCSimilarityDecisionAllowAdvisoryUse {
		for _, criticalDimension := range intelligenceCalibrationValCSimilarityCriticalDimensions() {
			if containsTrimmedString(model.MissingDimensions, criticalDimension) {
				return IntelligenceCalibrationValCSimilarityGatingStatePartial
			}
		}
	}
	if model.SimilarityBand == IntelligenceCalibrationValCSimilarityBandUnknown && model.GatingDecision != IntelligenceCalibrationValCSimilarityDecisionCapConfidence && model.GatingDecision != IntelligenceCalibrationValCSimilarityDecisionRequireReview {
		return IntelligenceCalibrationValCSimilarityGatingStatePartial
	}
	if model.SimilarityBand == IntelligenceCalibrationValCSimilarityBandLow && model.IncreasesConfidence {
		return IntelligenceCalibrationValCSimilarityGatingStatePartial
	}
	if model.OverridesLocalEvidence {
		return IntelligenceCalibrationValCSimilarityGatingStatePartial
	}
	return IntelligenceCalibrationValCSimilarityGatingStateActive
}

func EvaluateIntelligenceCalibrationValCLocalOverrideState(model LocalOverrideDisciplineContract) string {
	if strings.TrimSpace(model.OverridePolicyID) == "" || strings.TrimSpace(model.FederatedConfidenceCap) == "" || strings.TrimSpace(model.LocalPolicyBoundaryRef) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValCLocalOverrideStateIncomplete
	}
	if !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.FederatedConfidenceCap) {
		return IntelligenceCalibrationValCLocalOverrideStatePartial
	}
	if !model.LocalEvidenceWins || !model.LocalOverrideAllowed || !model.OverrideReasonRequired || !model.OverrideAuditRequired {
		return IntelligenceCalibrationValCLocalOverrideStatePartial
	}
	return IntelligenceCalibrationValCLocalOverrideStateActive
}

func EvaluateIntelligenceCalibrationValCPropagationPolicyState(model BoundedPropagationPolicyContract) string {
	if strings.TrimSpace(model.PropagationPolicyID) == "" || strings.TrimSpace(model.DefaultState) == "" || len(model.AllowedPayloadClasses) == 0 || len(model.BlockedPayloadClasses) == 0 || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValCPropagationPolicyStateIncomplete
	}
	if !containsTrimmedString(intelligenceCalibrationValCPropagationDefaultStates(), model.DefaultState) || model.PropagationAllowed || !model.RequiresRedaction || !model.RequiresReview || model.MutatesRemoteCalibration || model.PropagatesRawLocalEvidence {
		return IntelligenceCalibrationValCPropagationPolicyStatePartial
	}
	if model.DefaultState != IntelligenceCalibrationValCPropagationDisabled && model.DefaultState != IntelligenceCalibrationValCPropagationAdvisoryOnly {
		return IntelligenceCalibrationValCPropagationPolicyStatePartial
	}
	if !containsTrimmedString(model.BlockedPayloadClasses, "raw_local_evidence") || !containsTrimmedString(model.BlockedPayloadClasses, "unsupported_payload") {
		return IntelligenceCalibrationValCPropagationPolicyStatePartial
	}
	return IntelligenceCalibrationValCPropagationPolicyStateActive
}

func EvaluateIntelligenceCalibrationValCExplanationState(model FeedbackFederatedExplanationContract) string {
	if strings.TrimSpace(model.ReasonCode) == "" || strings.TrimSpace(model.HumanMessage) == "" || strings.TrimSpace(model.TechnicalDetail) == "" || (len(model.FeedbackRefs) == 0 && len(model.FederatedSignalRefs) == 0) || strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.EvidenceClass) == "" || strings.TrimSpace(model.FalsePositiveRiskNote) == "" || strings.TrimSpace(model.FalseNegativeRiskNote) == "" || strings.TrimSpace(model.LocalOverrideNote) == "" || strings.TrimSpace(model.PropagationNote) == "" || strings.TrimSpace(model.NextStep) == "" || strings.TrimSpace(model.VisibilityScope) == "" || strings.TrimSpace(model.RedactionTier) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValCExplanationStateIncomplete
	}
	if !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceBand) || !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), model.EvidenceClass) || !containsTrimmedString(productionUsabilityValAExplainScopes(), model.VisibilityScope) || !containsTrimmedString(ProductionUsabilityVal0ExplainabilityContract().SupportedRedactionTiers, model.RedactionTier) || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") || !model.ReviewerRequired {
		return IntelligenceCalibrationValCExplanationStatePartial
	}
	if model.LeaksInternalEvidence || model.ReviewRequiredPresentedApproved || !model.FederatedAdvisoryOnlyExplicit {
		return IntelligenceCalibrationValCExplanationStatePartial
	}
	return IntelligenceCalibrationValCExplanationStateActive
}

func EvaluateIntelligenceCalibrationValCState(val0DependencyState, val0FoundationState, valADependencyState, valAState, valBDependencyState, valBState, feedbackIntakeState, reviewCockpitState, tuningProposalState, suppressionSafetyState, suppressionRollbackState, localChangeReviewState, federatedWeightingState, similarityGatingState, localOverrideState, propagationPolicyState, explanationState string) string {
	if strings.TrimSpace(val0DependencyState) != IntelligenceCalibrationVal0StateActive || strings.TrimSpace(val0FoundationState) != IntelligenceCalibrationVal0StateActive || strings.TrimSpace(valADependencyState) != IntelligenceCalibrationValAStateActive || strings.TrimSpace(valAState) != IntelligenceCalibrationValAStateActive || strings.TrimSpace(valBDependencyState) != IntelligenceCalibrationValBStateActive || strings.TrimSpace(valBState) != IntelligenceCalibrationValBStateActive {
		return IntelligenceCalibrationValCStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(feedbackIntakeState),
		strings.TrimSpace(reviewCockpitState),
		strings.TrimSpace(tuningProposalState),
		strings.TrimSpace(suppressionSafetyState),
		strings.TrimSpace(suppressionRollbackState),
		strings.TrimSpace(localChangeReviewState),
		strings.TrimSpace(federatedWeightingState),
		strings.TrimSpace(similarityGatingState),
		strings.TrimSpace(localOverrideState),
		strings.TrimSpace(propagationPolicyState),
		strings.TrimSpace(explanationState),
	} {
		switch state {
		case IntelligenceCalibrationValCFeedbackIntakeStateActive,
			IntelligenceCalibrationValCReviewCockpitStateActive,
			IntelligenceCalibrationValCTuningProposalStateActive,
			IntelligenceCalibrationValCSuppressionSafetyStateActive,
			IntelligenceCalibrationValCSuppressionRollbackStateActive,
			IntelligenceCalibrationValCLocalChangeReviewStateActive,
			IntelligenceCalibrationValCFederatedWeightingStateActive,
			IntelligenceCalibrationValCSimilarityGatingStateActive,
			IntelligenceCalibrationValCLocalOverrideStateActive,
			IntelligenceCalibrationValCPropagationPolicyStateActive,
			IntelligenceCalibrationValCExplanationStateActive:
		case IntelligenceCalibrationValCFeedbackIntakeStatePartial,
			IntelligenceCalibrationValCReviewCockpitStatePartial,
			IntelligenceCalibrationValCTuningProposalStatePartial,
			IntelligenceCalibrationValCSuppressionSafetyStatePartial,
			IntelligenceCalibrationValCSuppressionRollbackStatePartial,
			IntelligenceCalibrationValCLocalChangeReviewStatePartial,
			IntelligenceCalibrationValCFederatedWeightingStatePartial,
			IntelligenceCalibrationValCSimilarityGatingStatePartial,
			IntelligenceCalibrationValCLocalOverrideStatePartial,
			IntelligenceCalibrationValCPropagationPolicyStatePartial,
			IntelligenceCalibrationValCExplanationStatePartial:
			hasPartial = true
		default:
			return IntelligenceCalibrationValCStateIncomplete
		}
	}
	if hasPartial {
		return IntelligenceCalibrationValCStateSubstantial
	}
	return IntelligenceCalibrationValCStateActive
}

func EvaluateIntelligenceCalibrationValCProofsState(val0DependencyState, val0FoundationState, valADependencyState, valAState, valBDependencyState, valBState, feedbackIntakeState, reviewCockpitState, tuningProposalState, suppressionSafetyState, suppressionRollbackState, localChangeReviewState, federatedWeightingState, similarityGatingState, localOverrideState, propagationPolicyState, explanationState string, surfaceRefs, evidenceRefs, limitations, whyPoint5NotPass []string, projectionDisclaimer string) string {
	baseState := EvaluateIntelligenceCalibrationValCState(
		val0DependencyState,
		val0FoundationState,
		valADependencyState,
		valAState,
		valBDependencyState,
		valBState,
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
	if len(surfaceRefs) < 12 || len(evidenceRefs) < 12 || len(limitations) == 0 || len(whyPoint5NotPass) == 0 || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "not_canonical_truth") {
		if baseState == IntelligenceCalibrationValCStateActive {
			return IntelligenceCalibrationValCStateSubstantial
		}
		return baseState
	}
	return baseState
}

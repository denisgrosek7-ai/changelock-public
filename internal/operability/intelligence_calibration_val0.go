package operability

import (
	"strings"
	"time"
)

const (
	IntelligenceCalibrationVal0DatasetStateActive     = "intelligence_calibration_val0_dataset_active"
	IntelligenceCalibrationVal0DatasetStatePartial    = "intelligence_calibration_val0_dataset_partial"
	IntelligenceCalibrationVal0DatasetStateIncomplete = "intelligence_calibration_val0_dataset_incomplete"

	IntelligenceCalibrationVal0ConfidenceStateActive     = "intelligence_calibration_val0_confidence_active"
	IntelligenceCalibrationVal0ConfidenceStatePartial    = "intelligence_calibration_val0_confidence_partial"
	IntelligenceCalibrationVal0ConfidenceStateIncomplete = "intelligence_calibration_val0_confidence_incomplete"

	IntelligenceCalibrationVal0LifecycleStateActive     = "intelligence_calibration_val0_lifecycle_active"
	IntelligenceCalibrationVal0LifecycleStatePartial    = "intelligence_calibration_val0_lifecycle_partial"
	IntelligenceCalibrationVal0LifecycleStateIncomplete = "intelligence_calibration_val0_lifecycle_incomplete"

	IntelligenceCalibrationVal0ReachabilityStateActive     = "intelligence_calibration_val0_reachability_active"
	IntelligenceCalibrationVal0ReachabilityStatePartial    = "intelligence_calibration_val0_reachability_partial"
	IntelligenceCalibrationVal0ReachabilityStateIncomplete = "intelligence_calibration_val0_reachability_incomplete"

	IntelligenceCalibrationVal0VEXStateActive     = "intelligence_calibration_val0_vex_active"
	IntelligenceCalibrationVal0VEXStatePartial    = "intelligence_calibration_val0_vex_partial"
	IntelligenceCalibrationVal0VEXStateIncomplete = "intelligence_calibration_val0_vex_incomplete"

	IntelligenceCalibrationVal0FeedbackStateActive     = "intelligence_calibration_val0_feedback_active"
	IntelligenceCalibrationVal0FeedbackStatePartial    = "intelligence_calibration_val0_feedback_partial"
	IntelligenceCalibrationVal0FeedbackStateIncomplete = "intelligence_calibration_val0_feedback_incomplete"

	IntelligenceCalibrationVal0LearningModeStateActive     = "intelligence_calibration_val0_learning_mode_active"
	IntelligenceCalibrationVal0LearningModeStatePartial    = "intelligence_calibration_val0_learning_mode_partial"
	IntelligenceCalibrationVal0LearningModeStateIncomplete = "intelligence_calibration_val0_learning_mode_incomplete"

	IntelligenceCalibrationVal0SuppressionStateActive     = "intelligence_calibration_val0_suppression_active"
	IntelligenceCalibrationVal0SuppressionStatePartial    = "intelligence_calibration_val0_suppression_partial"
	IntelligenceCalibrationVal0SuppressionStateIncomplete = "intelligence_calibration_val0_suppression_incomplete"

	IntelligenceCalibrationVal0FederatedBoundaryStateActive     = "intelligence_calibration_val0_federated_boundary_active"
	IntelligenceCalibrationVal0FederatedBoundaryStatePartial    = "intelligence_calibration_val0_federated_boundary_partial"
	IntelligenceCalibrationVal0FederatedBoundaryStateIncomplete = "intelligence_calibration_val0_federated_boundary_incomplete"

	IntelligenceCalibrationVal0ProvenanceStateActive     = "intelligence_calibration_val0_provenance_active"
	IntelligenceCalibrationVal0ProvenanceStatePartial    = "intelligence_calibration_val0_provenance_partial"
	IntelligenceCalibrationVal0ProvenanceStateIncomplete = "intelligence_calibration_val0_provenance_incomplete"

	IntelligenceCalibrationVal0FreshnessStateActive     = "intelligence_calibration_val0_freshness_active"
	IntelligenceCalibrationVal0FreshnessStatePartial    = "intelligence_calibration_val0_freshness_partial"
	IntelligenceCalibrationVal0FreshnessStateIncomplete = "intelligence_calibration_val0_freshness_incomplete"

	IntelligenceCalibrationVal0MetricsStateActive     = "intelligence_calibration_val0_metrics_active"
	IntelligenceCalibrationVal0MetricsStatePartial    = "intelligence_calibration_val0_metrics_partial"
	IntelligenceCalibrationVal0MetricsStateIncomplete = "intelligence_calibration_val0_metrics_incomplete"

	IntelligenceCalibrationVal0FPFNStateActive     = "intelligence_calibration_val0_fp_fn_active"
	IntelligenceCalibrationVal0FPFNStatePartial    = "intelligence_calibration_val0_fp_fn_partial"
	IntelligenceCalibrationVal0FPFNStateIncomplete = "intelligence_calibration_val0_fp_fn_incomplete"

	IntelligenceCalibrationVal0RollbackStateActive     = "intelligence_calibration_val0_rollback_active"
	IntelligenceCalibrationVal0RollbackStatePartial    = "intelligence_calibration_val0_rollback_partial"
	IntelligenceCalibrationVal0RollbackStateIncomplete = "intelligence_calibration_val0_rollback_incomplete"

	IntelligenceCalibrationVal0StateIncomplete  = "intelligence_calibration_val0_incomplete"
	IntelligenceCalibrationVal0StateSubstantial = "intelligence_calibration_val0_substantially_ready"
	IntelligenceCalibrationVal0StateActive      = "intelligence_calibration_val0_active"

	IntelligenceCalibrationPoint5StateNotComplete = "intelligence_calibration_point_5_not_complete"

	IntelligenceCalibrationFreshnessFresh       = "fresh"
	IntelligenceCalibrationFreshnessStale       = "stale"
	IntelligenceCalibrationFreshnessExpired     = "expired"
	IntelligenceCalibrationFreshnessUnknown     = "unknown"
	IntelligenceCalibrationFreshnessUnsupported = "unsupported"

	IntelligenceCalibrationConfidenceHigh    = "high"
	IntelligenceCalibrationConfidenceMedium  = "medium"
	IntelligenceCalibrationConfidenceLow     = "low"
	IntelligenceCalibrationConfidenceUnknown = "unknown"

	IntelligenceCalibrationEvidenceDirectlyEvidenced = "directly_evidenced"
	IntelligenceCalibrationEvidenceStronglyInferred  = "strongly_inferred"
	IntelligenceCalibrationEvidenceWeaklyInferred    = "weakly_inferred"
	IntelligenceCalibrationEvidenceUnsupported       = "unsupported"

	IntelligenceCalibrationLifecycleCandidate      = "candidate"
	IntelligenceCalibrationLifecycleObserved       = "observed"
	IntelligenceCalibrationLifecycleCalibrated     = "calibrated"
	IntelligenceCalibrationLifecycleReviewRequired = "review_required"
	IntelligenceCalibrationLifecycleAccepted       = "accepted"
	IntelligenceCalibrationLifecycleRejected       = "rejected"
	IntelligenceCalibrationLifecycleSuperseded     = "superseded"
	IntelligenceCalibrationLifecycleExpired        = "expired"

	IntelligenceCalibrationReachabilityPresentOnly         = "present_only"
	IntelligenceCalibrationReachabilityStaticallyReachable = "statically_reachable"
	IntelligenceCalibrationReachabilityRuntimeObserved     = "runtime_observed"
	IntelligenceCalibrationReachabilityExecutionPath       = "execution_path_evidenced"
	IntelligenceCalibrationReachabilityStronglyInferred    = "strongly_inferred"
	IntelligenceCalibrationReachabilityWeaklyInferred      = "weakly_inferred"
	IntelligenceCalibrationReachabilityUnsupported         = "unsupported"

	IntelligenceCalibrationVEXStateCandidate      = "candidate"
	IntelligenceCalibrationVEXStateRequiresReview = "requires_review"
	IntelligenceCalibrationVEXStateReviewed       = "reviewed"
	IntelligenceCalibrationVEXStateRejected       = "rejected"
	IntelligenceCalibrationVEXStateSuperseded     = "superseded"
	IntelligenceCalibrationVEXStateExpired        = "expired"

	IntelligenceCalibrationVEXOutcomeNotAffectedCandidate             = "not_affected_candidate"
	IntelligenceCalibrationVEXOutcomeAffectedNotExploitableCurrentCtx = "affected_but_not_exploitable_in_current_context"
	IntelligenceCalibrationVEXOutcomeHighRiskRelevant                 = "high_risk_relevant"
	IntelligenceCalibrationVEXOutcomeInsufficientEvidence             = "insufficient_evidence"
	IntelligenceCalibrationVEXOutcomeRequiresReview                   = "requires_review"

	IntelligenceCalibrationFeedbackFalsePositive    = "false_positive"
	IntelligenceCalibrationFeedbackNoisyButUseful   = "noisy_but_useful"
	IntelligenceCalibrationFeedbackCorrectLowPrio   = "correct_low_priority"
	IntelligenceCalibrationFeedbackMissedSeverity   = "missed_severity"
	IntelligenceCalibrationFeedbackFalseNegative    = "false_negative"
	IntelligenceCalibrationFeedbackNeedsMoreContext = "needs_more_context"
	IntelligenceCalibrationFeedbackAcceptedSignal   = "accepted_signal"
	IntelligenceCalibrationFeedbackRejectedSignal   = "rejected_signal"

	IntelligenceCalibrationLearningDisabled       = "disabled"
	IntelligenceCalibrationLearningEnabled        = "enabled"
	IntelligenceCalibrationLearningPaused         = "paused"
	IntelligenceCalibrationLearningExpired        = "expired"
	IntelligenceCalibrationLearningReviewRequired = "review_required"

	IntelligenceCalibrationApprovalProposed       = "proposed"
	IntelligenceCalibrationApprovalReviewRequired = "review_required"
	IntelligenceCalibrationApprovalApproved       = "approved"
	IntelligenceCalibrationApprovalRejected       = "rejected"
	IntelligenceCalibrationApprovalSuperseded     = "superseded"

	IntelligenceCalibrationMetricFalsePositiveRate         = "false_positive_rate"
	IntelligenceCalibrationMetricFalseNegativeReviewRate   = "false_negative_review_rate"
	IntelligenceCalibrationMetricOperatorAcceptanceRate    = "operator_acceptance_rate"
	IntelligenceCalibrationMetricSuppressionCorrectness    = "suppression_correctness_rate"
	IntelligenceCalibrationMetricMissedDetectionReviewRate = "missed_detection_review_rate"
	IntelligenceCalibrationMetricMeanTimeToTriage          = "mean_time_to_triage"
	IntelligenceCalibrationMetricTimeToDecisionDelta       = "time_to_decision_delta"
	IntelligenceCalibrationMetricVEXApprovalRate           = "vex_candidate_approval_rate"
	IntelligenceCalibrationMetricConfidenceError           = "confidence_calibration_error"
	IntelligenceCalibrationMetricLearningCompletion        = "learning_mode_completion_quality"
	IntelligenceCalibrationMetricFederatedUsefulness       = "federated_signal_usefulness_rate"
)

type CalibrationDatasetContract struct {
	CurrentState         string   `json:"current_state"`
	DatasetID            string   `json:"dataset_id"`
	DatasetVersion       string   `json:"dataset_version"`
	DatasetScope         string   `json:"dataset_scope"`
	ScenarioClasses      []string `json:"scenario_classes,omitempty"`
	LabeledCaseCount     int      `json:"labeled_case_count"`
	ReviewedCaseCount    int      `json:"reviewed_case_count"`
	SourceRefs           []string `json:"source_refs,omitempty"`
	EvidenceRefs         []string `json:"evidence_refs,omitempty"`
	FreshnessState       string   `json:"freshness_state"`
	SamplingPolicy       string   `json:"sampling_policy"`
	KnownBiases          []string `json:"known_biases,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
	GeneratedAt          string   `json:"generated_at"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type ConfidenceEvidenceClassContract struct {
	CurrentState             string   `json:"current_state"`
	SupportedConfidenceBands []string `json:"supported_confidence_bands,omitempty"`
	SupportedEvidenceClasses []string `json:"supported_evidence_classes,omitempty"`
	ConfidenceBand           string   `json:"confidence_band"`
	ConfidenceScore          int      `json:"confidence_score"`
	EvidenceClass            string   `json:"evidence_class"`
	UncertaintyNote          string   `json:"uncertainty_note"`
	FreshnessState           string   `json:"freshness_state"`
	ModelVersion             string   `json:"model_version"`
	RuleVersion              string   `json:"rule_version"`
	DatasetVersion           string   `json:"dataset_version"`
	ReasonTrace              []string `json:"reason_trace,omitempty"`
	LimitationMessage        string   `json:"limitation_message"`
	AdvisoryOnly             bool     `json:"advisory_only"`
}

type IntelligenceOutputLifecycleEntry struct {
	OutputID          string   `json:"output_id"`
	OutputType        string   `json:"output_type"`
	LifecycleState    string   `json:"lifecycle_state"`
	CreatedAt         string   `json:"created_at"`
	UpdatedAt         string   `json:"updated_at"`
	ExpiresAt         string   `json:"expires_at,omitempty"`
	SupersedesRef     string   `json:"supersedes_ref,omitempty"`
	ReviewerRef       string   `json:"reviewer_ref,omitempty"`
	EvidenceRefs      []string `json:"evidence_refs,omitempty"`
	ReasonCode        string   `json:"reason_code"`
	LimitationMessage string   `json:"limitation_message"`
}

type IntelligenceOutputLifecycleContract struct {
	CurrentState                  string                             `json:"current_state"`
	SupportedLifecycleStates      []string                           `json:"supported_lifecycle_states,omitempty"`
	Items                         []IntelligenceOutputLifecycleEntry `json:"items,omitempty"`
	CandidateTreatedAsAccepted    bool                               `json:"candidate_treated_as_accepted"`
	ReviewRequiredTreatedApproved bool                               `json:"review_required_treated_as_approved"`
	ExpiredTreatedAsActive        bool                               `json:"expired_treated_as_active"`
	AdvisoryOnly                  bool                               `json:"advisory_only"`
	Limitations                   []string                           `json:"limitations,omitempty"`
}

type ReachabilityTaxonomyEntry struct {
	ReachabilityClass      string   `json:"reachability_class"`
	PackageRef             string   `json:"package_ref"`
	FunctionOrComponentRef string   `json:"function_or_component_ref"`
	WorkloadContextRef     string   `json:"workload_context_ref"`
	RuntimeSignalRefs      []string `json:"runtime_signal_refs,omitempty"`
	StaticSignalRefs       []string `json:"static_signal_refs,omitempty"`
	ConfidenceBand         string   `json:"confidence_band"`
	EvidenceClass          string   `json:"evidence_class"`
	DowngradeAllowed       bool     `json:"downgrade_allowed"`
	EscalationAllowed      bool     `json:"escalation_allowed"`
	ExplanationRequired    bool     `json:"explanation_required"`
	Explanation            string   `json:"explanation"`
}

type ReachabilityTaxonomyContract struct {
	CurrentState                   string                    `json:"current_state"`
	SupportedReachabilityClasses   []string                  `json:"supported_reachability_classes,omitempty"`
	SupportedConfidenceBands       []string                  `json:"supported_confidence_bands,omitempty"`
	Example                        ReachabilityTaxonomyEntry `json:"example"`
	PackagePresenceImpliesExploit  bool                      `json:"package_presence_alone_implies_exploitable_reachability"`
	NoObservedExecutionImpliesSafe bool                      `json:"lack_of_observed_execution_implies_safe"`
	AdvisoryOnly                   bool                      `json:"advisory_only"`
	Limitations                    []string                  `json:"limitations,omitempty"`
}

type VEXCandidateEntry struct {
	CandidateID       string   `json:"candidate_id"`
	VulnerabilityRef  string   `json:"vulnerability_ref"`
	ProductOrAssetRef string   `json:"product_or_asset_ref"`
	SuggestedOutcome  string   `json:"suggested_outcome"`
	CandidateState    string   `json:"candidate_state"`
	EvidenceRefs      []string `json:"evidence_refs,omitempty"`
	ReachabilityRef   string   `json:"reachability_ref"`
	ConfidenceBand    string   `json:"confidence_band"`
	ReviewerRef       string   `json:"reviewer_ref,omitempty"`
	Expiry            string   `json:"expiry"`
	ReasonCode        string   `json:"reason_code"`
	LimitationMessage string   `json:"limitation_message"`
}

type VEXCandidateGovernanceContract struct {
	CurrentState                    string            `json:"current_state"`
	SupportedStates                 []string          `json:"supported_states,omitempty"`
	SupportedSuggestedOutcomes      []string          `json:"supported_suggested_outcomes,omitempty"`
	Example                         VEXCandidateEntry `json:"example"`
	CandidateIsFinalVEX             bool              `json:"candidate_is_final_vex"`
	InsufficientEvidenceBecomesSafe bool              `json:"insufficient_evidence_becomes_not_affected"`
	AdvisoryOnly                    bool              `json:"advisory_only"`
	Limitations                     []string          `json:"limitations,omitempty"`
}

type FeedbackClassificationEntry struct {
	FeedbackID         string   `json:"feedback_id"`
	ActorRef           string   `json:"actor_ref"`
	SignalRef          string   `json:"signal_ref"`
	FeedbackClass      string   `json:"feedback_class"`
	Reason             string   `json:"reason"`
	EvidenceRefs       []string `json:"evidence_refs,omitempty"`
	LocalApplicability string   `json:"local_applicability"`
	ReviewRequired     bool     `json:"review_required"`
	CreatedAt          string   `json:"created_at"`
	ProposedTuningRef  string   `json:"proposed_tuning_ref,omitempty"`
}

type FeedbackClassificationContract struct {
	CurrentState                string                        `json:"current_state"`
	SupportedFeedbackClasses    []string                      `json:"supported_feedback_classes,omitempty"`
	Items                       []FeedbackClassificationEntry `json:"items,omitempty"`
	DirectMutationAllowed       bool                          `json:"direct_mutation_allowed"`
	DirectSuppressionAllowed    bool                          `json:"direct_suppression_allowed"`
	FederatedPropagationAllowed bool                          `json:"federated_propagation_allowed"`
	AdvisoryOnly                bool                          `json:"advisory_only"`
	Limitations                 []string                      `json:"limitations,omitempty"`
}

type LearningModeGuardrailContract struct {
	CurrentState             string   `json:"current_state"`
	SupportedStates          []string `json:"supported_states,omitempty"`
	LearningModeState        string   `json:"learning_mode_state"`
	Scope                    string   `json:"scope"`
	StartedAt                string   `json:"started_at"`
	ExpiresAt                string   `json:"expires_at"`
	ExcludedCriticalControls []string `json:"excluded_critical_controls,omitempty"`
	AllowedSignalClasses     []string `json:"allowed_signal_classes,omitempty"`
	DisallowedSignalClasses  []string `json:"disallowed_signal_classes,omitempty"`
	OutputReviewRequired     bool     `json:"output_review_required"`
	CanRelaxEnforcement      bool     `json:"can_relax_enforcement"`
	LimitationMessage        string   `json:"limitation_message"`
}

type SuppressionSafetyContract struct {
	CurrentState            string   `json:"current_state"`
	SuppressionID           string   `json:"suppression_id"`
	SignalClass             string   `json:"signal_class"`
	SuppressionScope        string   `json:"suppression_scope"`
	ReasonCode              string   `json:"reason_code"`
	CreatedBy               string   `json:"created_by"`
	ReviewerRef             string   `json:"reviewer_ref"`
	ExpiresAt               string   `json:"expires_at"`
	AffectedSubjects        []string `json:"affected_subjects,omitempty"`
	ExcludedCriticalClasses []string `json:"excluded_critical_classes,omitempty"`
	ReopenOnNewEvidence     bool     `json:"reopen_on_new_evidence"`
	RollbackRef             string   `json:"rollback_ref"`
	EvidenceRefs            []string `json:"evidence_refs,omitempty"`
	LimitationMessage       string   `json:"limitation_message"`
	DeletesEvidence         bool     `json:"deletes_evidence"`
}

type FederatedSignalBoundaryContract struct {
	CurrentState           string   `json:"current_state"`
	FederatedSignalID      string   `json:"federated_signal_id"`
	SourceRef              string   `json:"source_ref"`
	SourceTrustWeight      int      `json:"source_trust_weight"`
	EnvironmentSimilarity  int      `json:"environment_similarity_score"`
	LocalOverrideAllowed   bool     `json:"local_override_allowed"`
	LocalEvidenceWins      bool     `json:"local_evidence_wins"`
	PropagationAllowed     bool     `json:"propagation_allowed"`
	PropagationNonMutating bool     `json:"propagation_non_mutating"`
	ConfidenceCap          string   `json:"confidence_cap"`
	ReasonTrace            []string `json:"reason_trace,omitempty"`
	LimitationMessage      string   `json:"limitation_message"`
	AutoMarksLocalSafe     bool     `json:"auto_marks_local_asset_safe"`
}

type CalibrationProvenanceChainContract struct {
	CurrentState             string   `json:"current_state"`
	SupportedApprovalStates  []string `json:"supported_approval_states,omitempty"`
	CalibrationChangeID      string   `json:"calibration_change_id"`
	SourceSignalRefs         []string `json:"source_signal_refs,omitempty"`
	DatasetRefs              []string `json:"dataset_refs,omitempty"`
	FeedbackRefs             []string `json:"feedback_refs,omitempty"`
	ModelVersion             string   `json:"model_version"`
	RuleVersion              string   `json:"rule_version"`
	ActorOrReviewerRef       string   `json:"actor_or_reviewer_ref"`
	CreatedAt                string   `json:"created_at"`
	ProposedChangeSummary    string   `json:"proposed_change_summary"`
	ExpectedEffect           string   `json:"expected_effect"`
	RiskNote                 string   `json:"risk_note"`
	ApprovalState            string   `json:"approval_state"`
	RollbackRef              string   `json:"rollback_ref"`
	MutatesActiveCalibration bool     `json:"mutates_active_calibration"`
}

type IntelligenceFreshnessEntry struct {
	IntelligenceType  string `json:"intelligence_type"`
	FreshnessState    string `json:"freshness_state"`
	ExpiresAt         string `json:"expires_at,omitempty"`
	LimitationMessage string `json:"limitation_message"`
}

type IntelligenceFreshnessExpiryContract struct {
	CurrentState                      string                       `json:"current_state"`
	SupportedFreshnessStates          []string                     `json:"supported_freshness_states,omitempty"`
	RequiredIntelligenceTypes         []string                     `json:"required_intelligence_types,omitempty"`
	Items                             []IntelligenceFreshnessEntry `json:"items,omitempty"`
	StaleTreatedAsFresh               bool                         `json:"stale_treated_as_fresh"`
	ExpiredTreatedAsActive            bool                         `json:"expired_treated_as_active"`
	UnknownFreshnessTreatedCalibrated bool                         `json:"unknown_freshness_treated_as_calibrated"`
	FreshnessAffectsConfidence        bool                         `json:"freshness_affects_confidence"`
	Limitations                       []string                     `json:"limitations,omitempty"`
}

type CalibrationMetricDefinition struct {
	MetricName       string   `json:"metric_name"`
	Definition       string   `json:"definition"`
	MeasurementScope string   `json:"measurement_scope"`
	Numerator        string   `json:"numerator"`
	Denominator      string   `json:"denominator"`
	SamplingPolicy   string   `json:"sampling_policy"`
	Freshness        string   `json:"freshness"`
	Limitations      []string `json:"limitations,omitempty"`
}

type CalibrationMetricsDefinitionContract struct {
	CurrentState              string                        `json:"current_state"`
	SupportedSamplingPolicies []string                      `json:"supported_sampling_policies,omitempty"`
	SupportedFreshnessStates  []string                      `json:"supported_freshness_states,omitempty"`
	Items                     []CalibrationMetricDefinition `json:"items,omitempty"`
	Limitations               []string                      `json:"limitations,omitempty"`
}

type FalsePositiveNegativeDisciplineContract struct {
	CurrentState                                string   `json:"current_state"`
	FalsePositiveQueueState                     string   `json:"false_positive_queue_state"`
	FalseNegativeQueueState                     string   `json:"false_negative_queue_state"`
	MissedDetectionScenarios                    []string `json:"missed_detection_scenarios,omitempty"`
	SuppressionCausedMissCheck                  bool     `json:"suppression_caused_miss_check"`
	CriticalLowSignalReviewRequired             bool     `json:"critical_low_signal_review_required"`
	ReviewCadence                               string   `json:"review_cadence"`
	EvidenceRefs                                []string `json:"evidence_refs,omitempty"`
	LimitationMessage                           string   `json:"limitation_message"`
	FalsePositiveReductionIgnoresFalseNegatives bool     `json:"false_positive_reduction_ignores_false_negative_risk"`
}

type CalibrationRollbackUndoContract struct {
	CurrentState            string `json:"current_state"`
	CalibrationChangeRef    string `json:"calibration_change_ref"`
	PreviewAvailable        bool   `json:"preview_available"`
	StagedRolloutSupported  bool   `json:"staged_rollout_supported"`
	RollbackAvailable       bool   `json:"rollback_available"`
	RollbackReasonRequired  bool   `json:"rollback_reason_required"`
	RollbackSafetyCheck     string `json:"rollback_safety_check"`
	BeforeMetricSnapshotRef string `json:"before_metric_snapshot_ref"`
	AfterMetricSnapshotRef  string `json:"after_metric_snapshot_ref"`
	LimitationMessage       string `json:"limitation_message"`
}

func intelligenceCalibrationVal0ScenarioClasses() []string {
	return []string{
		"vulnerability_relevance",
		"reachability",
		"vex_candidate",
		"anomaly",
		"drift",
		"behavioral_baseline",
		"false_positive",
		"false_negative",
		"federated_signal",
		"suppression",
		"learning_mode",
	}
}

func intelligenceCalibrationVal0ConfidenceBands() []string {
	return []string{
		IntelligenceCalibrationConfidenceHigh,
		IntelligenceCalibrationConfidenceMedium,
		IntelligenceCalibrationConfidenceLow,
		IntelligenceCalibrationConfidenceUnknown,
	}
}

func intelligenceCalibrationVal0EvidenceClasses() []string {
	return []string{
		IntelligenceCalibrationEvidenceDirectlyEvidenced,
		IntelligenceCalibrationEvidenceStronglyInferred,
		IntelligenceCalibrationEvidenceWeaklyInferred,
		IntelligenceCalibrationEvidenceUnsupported,
	}
}

func intelligenceCalibrationVal0LifecycleStates() []string {
	return []string{
		IntelligenceCalibrationLifecycleCandidate,
		IntelligenceCalibrationLifecycleObserved,
		IntelligenceCalibrationLifecycleCalibrated,
		IntelligenceCalibrationLifecycleReviewRequired,
		IntelligenceCalibrationLifecycleAccepted,
		IntelligenceCalibrationLifecycleRejected,
		IntelligenceCalibrationLifecycleSuperseded,
		IntelligenceCalibrationLifecycleExpired,
	}
}

func intelligenceCalibrationVal0ReachabilityClasses() []string {
	return []string{
		IntelligenceCalibrationReachabilityPresentOnly,
		IntelligenceCalibrationReachabilityStaticallyReachable,
		IntelligenceCalibrationReachabilityRuntimeObserved,
		IntelligenceCalibrationReachabilityExecutionPath,
		IntelligenceCalibrationReachabilityStronglyInferred,
		IntelligenceCalibrationReachabilityWeaklyInferred,
		IntelligenceCalibrationReachabilityUnsupported,
	}
}

func intelligenceCalibrationVal0VEXStates() []string {
	return []string{
		IntelligenceCalibrationVEXStateCandidate,
		IntelligenceCalibrationVEXStateRequiresReview,
		IntelligenceCalibrationVEXStateReviewed,
		IntelligenceCalibrationVEXStateRejected,
		IntelligenceCalibrationVEXStateSuperseded,
		IntelligenceCalibrationVEXStateExpired,
	}
}

func intelligenceCalibrationVal0VEXOutcomes() []string {
	return []string{
		IntelligenceCalibrationVEXOutcomeNotAffectedCandidate,
		IntelligenceCalibrationVEXOutcomeAffectedNotExploitableCurrentCtx,
		IntelligenceCalibrationVEXOutcomeHighRiskRelevant,
		IntelligenceCalibrationVEXOutcomeInsufficientEvidence,
		IntelligenceCalibrationVEXOutcomeRequiresReview,
	}
}

func intelligenceCalibrationVal0FeedbackClasses() []string {
	return []string{
		IntelligenceCalibrationFeedbackFalsePositive,
		IntelligenceCalibrationFeedbackNoisyButUseful,
		IntelligenceCalibrationFeedbackCorrectLowPrio,
		IntelligenceCalibrationFeedbackMissedSeverity,
		IntelligenceCalibrationFeedbackFalseNegative,
		IntelligenceCalibrationFeedbackNeedsMoreContext,
		IntelligenceCalibrationFeedbackAcceptedSignal,
		IntelligenceCalibrationFeedbackRejectedSignal,
	}
}

func intelligenceCalibrationVal0LearningStates() []string {
	return []string{
		IntelligenceCalibrationLearningDisabled,
		IntelligenceCalibrationLearningEnabled,
		IntelligenceCalibrationLearningPaused,
		IntelligenceCalibrationLearningExpired,
		IntelligenceCalibrationLearningReviewRequired,
	}
}

func intelligenceCalibrationVal0FreshnessStates() []string {
	return []string{
		IntelligenceCalibrationFreshnessFresh,
		IntelligenceCalibrationFreshnessStale,
		IntelligenceCalibrationFreshnessExpired,
		IntelligenceCalibrationFreshnessUnknown,
		IntelligenceCalibrationFreshnessUnsupported,
	}
}

func intelligenceCalibrationVal0ApprovalStates() []string {
	return []string{
		IntelligenceCalibrationApprovalProposed,
		IntelligenceCalibrationApprovalReviewRequired,
		IntelligenceCalibrationApprovalApproved,
		IntelligenceCalibrationApprovalRejected,
		IntelligenceCalibrationApprovalSuperseded,
	}
}

func intelligenceCalibrationVal0RequiredIntelligenceTypes() []string {
	return []string{
		"reachability_inference",
		"vex_candidate",
		"behavioral_baseline",
		"feedback_derived_tuning",
		"federated_signal",
		"anomaly_threshold",
		"calibration_dataset",
		"confidence_output",
	}
}

func intelligenceCalibrationVal0RequiredMetrics() []string {
	return []string{
		IntelligenceCalibrationMetricFalsePositiveRate,
		IntelligenceCalibrationMetricFalseNegativeReviewRate,
		IntelligenceCalibrationMetricOperatorAcceptanceRate,
		IntelligenceCalibrationMetricSuppressionCorrectness,
		IntelligenceCalibrationMetricMissedDetectionReviewRate,
		IntelligenceCalibrationMetricMeanTimeToTriage,
		IntelligenceCalibrationMetricTimeToDecisionDelta,
		IntelligenceCalibrationMetricVEXApprovalRate,
		IntelligenceCalibrationMetricConfidenceError,
		IntelligenceCalibrationMetricLearningCompletion,
		IntelligenceCalibrationMetricFederatedUsefulness,
	}
}

func intelligenceCalibrationVal0SamplingPolicies() []string {
	return []string{
		"event_sampled",
		"review_window",
		"bounded_snapshot",
		"manual_review_window",
		"unknown_explicit",
	}
}

func containsSubstringInTrimmedStrings(values []string, expected string) bool {
	expected = strings.TrimSpace(expected)
	if expected == "" {
		return false
	}
	for _, value := range values {
		if strings.Contains(strings.TrimSpace(value), expected) {
			return true
		}
	}
	return false
}

func IntelligenceCalibrationVal0DatasetContract() CalibrationDatasetContract {
	return CalibrationDatasetContract{
		CurrentState:      "calibration_dataset_contract_ready",
		DatasetID:         "intelligence-calibration-val0-dataset",
		DatasetVersion:    "point5.intelligence_calibration.val0.dataset.v1",
		DatasetScope:      "bounded_reviewed_calibration_seed",
		ScenarioClasses:   intelligenceCalibrationVal0ScenarioClasses(),
		LabeledCaseCount:  168,
		ReviewedCaseCount: 168,
		SourceRefs: []string{
			"local_review_set",
			"bounded_operator_feedback_snapshot",
		},
		EvidenceRefs: []string{
			"dataset_manifest",
			"review_annotations",
		},
		FreshnessState: IntelligenceCalibrationFreshnessFresh,
		SamplingPolicy: "review_window",
		KnownBiases: []string{
			"low_frequency_false_negative_examples_remain bounded and may underrepresent rare exploit chains",
			"federated inputs are represented only as advisory examples and not as live propagation truth",
		},
		Limitations: []string{
			"Dataset is projection_only not_canonical_truth and seeds later calibration reviews rather than proving universal intelligence quality.",
			"Stale or bounded samples must remain explicitly limited before later tuning or reachability claims are made.",
		},
		GeneratedAt:          "2026-04-25T08:00:00Z",
		ProjectionDisclaimer: "projection_only not_canonical_truth calibration_dataset",
	}
}

func IntelligenceCalibrationVal0ConfidenceContract() ConfidenceEvidenceClassContract {
	return ConfidenceEvidenceClassContract{
		CurrentState:             "confidence_contract_ready",
		SupportedConfidenceBands: intelligenceCalibrationVal0ConfidenceBands(),
		SupportedEvidenceClasses: intelligenceCalibrationVal0EvidenceClasses(),
		ConfidenceBand:           IntelligenceCalibrationConfidenceMedium,
		ConfidenceScore:          72,
		EvidenceClass:            IntelligenceCalibrationEvidenceStronglyInferred,
		UncertaintyNote:          "bounded runtime and static signals still require operator review before policy-relevant action",
		FreshnessState:           IntelligenceCalibrationFreshnessFresh,
		ModelVersion:             "point5.intelligence_calibration.val0.model.v1",
		RuleVersion:              "point5.intelligence_calibration.val0.rules.v1",
		DatasetVersion:           "point5.intelligence_calibration.val0.dataset.v1",
		ReasonTrace: []string{
			"confidence is derived from explicit evidence class plus bounded reasoning trace",
			"confidence remains advisory and cannot approve or mutate governance decisions",
		},
		LimitationMessage: "Confidence is projection_only, reason traced, and bounded by evidence freshness and dataset version.",
		AdvisoryOnly:      true,
	}
}

func IntelligenceCalibrationVal0OutputLifecycleContract() IntelligenceOutputLifecycleContract {
	return IntelligenceOutputLifecycleContract{
		CurrentState:             "intelligence_output_lifecycle_contract_ready",
		SupportedLifecycleStates: intelligenceCalibrationVal0LifecycleStates(),
		Items: []IntelligenceOutputLifecycleEntry{
			{OutputID: "out-candidate", OutputType: "reachability_candidate", LifecycleState: IntelligenceCalibrationLifecycleCandidate, CreatedAt: "2026-04-25T08:00:00Z", UpdatedAt: "2026-04-25T08:00:00Z", EvidenceRefs: []string{"reachability_static_signal"}, ReasonCode: "candidate_only_until_review", LimitationMessage: "candidate output is not accepted or final"},
			{OutputID: "out-observed", OutputType: "runtime_signal", LifecycleState: IntelligenceCalibrationLifecycleObserved, CreatedAt: "2026-04-25T08:00:00Z", UpdatedAt: "2026-04-25T08:05:00Z", EvidenceRefs: []string{"runtime_signal_ref"}, ReasonCode: "observed_runtime_signal", LimitationMessage: "observed output still requires bounded interpretation"},
			{OutputID: "out-calibrated", OutputType: "confidence_projection", LifecycleState: IntelligenceCalibrationLifecycleCalibrated, CreatedAt: "2026-04-25T08:00:00Z", UpdatedAt: "2026-04-25T08:10:00Z", EvidenceRefs: []string{"confidence_contract"}, ReasonCode: "calibrated_projection_only", LimitationMessage: "calibrated output remains advisory"},
			{OutputID: "out-review", OutputType: "vex_candidate", LifecycleState: IntelligenceCalibrationLifecycleReviewRequired, CreatedAt: "2026-04-25T08:00:00Z", UpdatedAt: "2026-04-25T08:15:00Z", EvidenceRefs: []string{"vex_candidate_contract"}, ReasonCode: "requires_human_review", LimitationMessage: "review_required is not approved"},
			{OutputID: "out-accepted", OutputType: "feedback_signal", LifecycleState: IntelligenceCalibrationLifecycleAccepted, CreatedAt: "2026-04-25T08:00:00Z", UpdatedAt: "2026-04-25T08:20:00Z", ReviewerRef: "reviewer:security-admin", EvidenceRefs: []string{"feedback_item"}, ReasonCode: "accepted_for_later_governed_tuning", LimitationMessage: "accepted still does not mutate policy or evidence directly"},
			{OutputID: "out-rejected", OutputType: "federated_hint", LifecycleState: IntelligenceCalibrationLifecycleRejected, CreatedAt: "2026-04-25T08:00:00Z", UpdatedAt: "2026-04-25T08:25:00Z", ReviewerRef: "reviewer:security-admin", EvidenceRefs: []string{"federated_hint"}, ReasonCode: "rejected_due_to_local_evidence_override", LimitationMessage: "rejected output remains traceable"},
			{OutputID: "out-superseded", OutputType: "anomaly_candidate", LifecycleState: IntelligenceCalibrationLifecycleSuperseded, CreatedAt: "2026-04-25T08:00:00Z", UpdatedAt: "2026-04-25T08:30:00Z", SupersedesRef: "out-observed", ReviewerRef: "reviewer:operator", EvidenceRefs: []string{"anomaly_candidate_ref"}, ReasonCode: "superseded_by_newer_context", LimitationMessage: "superseded output points to later context"},
			{OutputID: "out-expired", OutputType: "learning_mode_observation", LifecycleState: IntelligenceCalibrationLifecycleExpired, CreatedAt: "2026-04-25T08:00:00Z", UpdatedAt: "2026-04-25T08:35:00Z", ExpiresAt: "2026-04-26T08:35:00Z", EvidenceRefs: []string{"learning_mode_window"}, ReasonCode: "expired_observation_window", LimitationMessage: "expired output is no longer active"},
		},
		CandidateTreatedAsAccepted:    false,
		ReviewRequiredTreatedApproved: false,
		ExpiredTreatedAsActive:        false,
		AdvisoryOnly:                  true,
		Limitations: []string{
			"Lifecycle states describe intelligence advisory status only and never create final governance approval.",
		},
	}
}

func IntelligenceCalibrationVal0ReachabilityTaxonomyContract() ReachabilityTaxonomyContract {
	return ReachabilityTaxonomyContract{
		CurrentState:                 "reachability_taxonomy_contract_ready",
		SupportedReachabilityClasses: intelligenceCalibrationVal0ReachabilityClasses(),
		SupportedConfidenceBands:     intelligenceCalibrationVal0ConfidenceBands(),
		Example: ReachabilityTaxonomyEntry{
			ReachabilityClass:      IntelligenceCalibrationReachabilityRuntimeObserved,
			PackageRef:             "pkg:golang/github.com/acme/api",
			FunctionOrComponentRef: "github.com/acme/api/internal/handler.Login",
			WorkloadContextRef:     "cluster-a/acme-prod/Deployment/api",
			RuntimeSignalRefs:      []string{"runtime:callgraph/login"},
			StaticSignalRefs:       []string{"static:import_graph/login"},
			ConfidenceBand:         IntelligenceCalibrationConfidenceMedium,
			EvidenceClass:          IntelligenceCalibrationEvidenceDirectlyEvidenced,
			DowngradeAllowed:       true,
			EscalationAllowed:      true,
			ExplanationRequired:    true,
			Explanation:            "Reachability remains evidence-classed and does not turn package presence alone into exploitable context.",
		},
		PackagePresenceImpliesExploit:  false,
		NoObservedExecutionImpliesSafe: false,
		AdvisoryOnly:                   true,
		Limitations: []string{
			"Reachability output is bounded advisory context and cannot replace exploitability governance or final VEX review.",
		},
	}
}

func IntelligenceCalibrationVal0VEXCandidateContract() VEXCandidateGovernanceContract {
	return VEXCandidateGovernanceContract{
		CurrentState:               "vex_candidate_contract_ready",
		SupportedStates:            intelligenceCalibrationVal0VEXStates(),
		SupportedSuggestedOutcomes: intelligenceCalibrationVal0VEXOutcomes(),
		Example: VEXCandidateEntry{
			CandidateID:       "vex-candidate-1",
			VulnerabilityRef:  "CVE-2026-1001",
			ProductOrAssetRef: "cluster-a/acme-prod/Deployment/api",
			SuggestedOutcome:  IntelligenceCalibrationVEXOutcomeRequiresReview,
			CandidateState:    IntelligenceCalibrationVEXStateRequiresReview,
			EvidenceRefs:      []string{"reachability_runtime_signal", "package_context"},
			ReachabilityRef:   "reachability:login",
			ConfidenceBand:    IntelligenceCalibrationConfidenceMedium,
			Expiry:            "2026-04-30T08:00:00Z",
			ReasonCode:        "candidate_only_requires_governed_review",
			LimitationMessage: "VEX candidate is not final VEX and remains advisory until later governance.",
		},
		CandidateIsFinalVEX:             false,
		InsufficientEvidenceBecomesSafe: false,
		AdvisoryOnly:                    true,
		Limitations: []string{
			"Candidate VEX posture remains bounded suggestion only and does not publish final VEX truth in Val 0.",
		},
	}
}

func IntelligenceCalibrationVal0FeedbackContract() FeedbackClassificationContract {
	return FeedbackClassificationContract{
		CurrentState:             "feedback_classification_contract_ready",
		SupportedFeedbackClasses: intelligenceCalibrationVal0FeedbackClasses(),
		Items: []FeedbackClassificationEntry{
			{FeedbackID: "feedback-fp", ActorRef: "actor:operator", SignalRef: "signal:1", FeedbackClass: IntelligenceCalibrationFeedbackFalsePositive, Reason: "noise exceeded current threshold", EvidenceRefs: []string{"alert:1"}, LocalApplicability: "tenant_local", ReviewRequired: true, CreatedAt: "2026-04-25T08:00:00Z", ProposedTuningRef: "tuning:1"},
			{FeedbackID: "feedback-noisy", ActorRef: "actor:operator", SignalRef: "signal:2", FeedbackClass: IntelligenceCalibrationFeedbackNoisyButUseful, Reason: "signal is useful but should be grouped better", EvidenceRefs: []string{"alert:2"}, LocalApplicability: "tenant_local", ReviewRequired: true, CreatedAt: "2026-04-25T08:01:00Z", ProposedTuningRef: "tuning:2"},
			{FeedbackID: "feedback-low", ActorRef: "actor:operator", SignalRef: "signal:3", FeedbackClass: IntelligenceCalibrationFeedbackCorrectLowPrio, Reason: "signal is valid but lower urgency", EvidenceRefs: []string{"alert:3"}, LocalApplicability: "tenant_local", ReviewRequired: true, CreatedAt: "2026-04-25T08:02:00Z"},
			{FeedbackID: "feedback-missed", ActorRef: "actor:operator", SignalRef: "signal:4", FeedbackClass: IntelligenceCalibrationFeedbackMissedSeverity, Reason: "severity under-classified current blast radius", EvidenceRefs: []string{"alert:4"}, LocalApplicability: "tenant_local", ReviewRequired: true, CreatedAt: "2026-04-25T08:03:00Z", ProposedTuningRef: "tuning:4"},
			{FeedbackID: "feedback-fn", ActorRef: "actor:operator", SignalRef: "signal:5", FeedbackClass: IntelligenceCalibrationFeedbackFalseNegative, Reason: "expected signal was absent", EvidenceRefs: []string{"evidence:missing"}, LocalApplicability: "tenant_local", ReviewRequired: true, CreatedAt: "2026-04-25T08:04:00Z", ProposedTuningRef: "tuning:5"},
			{FeedbackID: "feedback-context", ActorRef: "actor:operator", SignalRef: "signal:6", FeedbackClass: IntelligenceCalibrationFeedbackNeedsMoreContext, Reason: "context missing to calibrate safely", EvidenceRefs: []string{"signal:6"}, LocalApplicability: "tenant_local", ReviewRequired: true, CreatedAt: "2026-04-25T08:05:00Z"},
			{FeedbackID: "feedback-accepted", ActorRef: "actor:operator", SignalRef: "signal:7", FeedbackClass: IntelligenceCalibrationFeedbackAcceptedSignal, Reason: "signal classification is useful as-is", EvidenceRefs: []string{"signal:7"}, LocalApplicability: "tenant_local", ReviewRequired: true, CreatedAt: "2026-04-25T08:06:00Z"},
			{FeedbackID: "feedback-rejected", ActorRef: "actor:operator", SignalRef: "signal:8", FeedbackClass: IntelligenceCalibrationFeedbackRejectedSignal, Reason: "signal is not locally useful", EvidenceRefs: []string{"signal:8"}, LocalApplicability: "tenant_local", ReviewRequired: true, CreatedAt: "2026-04-25T08:07:00Z"},
		},
		DirectMutationAllowed:       false,
		DirectSuppressionAllowed:    false,
		FederatedPropagationAllowed: false,
		AdvisoryOnly:                true,
		Limitations: []string{
			"Feedback is advisory input only and requires later reviewed tuning before any stable calibration effect.",
		},
	}
}

func IntelligenceCalibrationVal0LearningModeContract() LearningModeGuardrailContract {
	return LearningModeGuardrailContract{
		CurrentState:             "learning_mode_guardrail_contract_ready",
		SupportedStates:          intelligenceCalibrationVal0LearningStates(),
		LearningModeState:        IntelligenceCalibrationLearningEnabled,
		Scope:                    "tenant_local bounded observation window",
		StartedAt:                "2026-04-25T08:00:00Z",
		ExpiresAt:                "2026-04-26T08:00:00Z",
		ExcludedCriticalControls: []string{"critical_policy_denials", "critical_runtime_blockers"},
		AllowedSignalClasses:     []string{"anomaly", "behavioral_baseline", "false_positive"},
		DisallowedSignalClasses:  []string{"suppression", "federated_signal", "critical_runtime_blockers"},
		OutputReviewRequired:     true,
		CanRelaxEnforcement:      false,
		LimitationMessage:        "Learning mode is bounded observation only and cannot relax critical controls or approve new stable baselines automatically.",
	}
}

func IntelligenceCalibrationVal0SuppressionSafetyContract() SuppressionSafetyContract {
	return SuppressionSafetyContract{
		CurrentState:            "suppression_safety_contract_ready",
		SuppressionID:           "suppression-1",
		SignalClass:             "anomaly",
		SuppressionScope:        "tenant_local bounded workload scope",
		ReasonCode:              "operator_noise_window_reviewed",
		CreatedBy:               "actor:operator",
		ReviewerRef:             "reviewer:security-admin",
		ExpiresAt:               "2026-04-26T08:00:00Z",
		AffectedSubjects:        []string{"cluster-a/acme-prod/Deployment/api"},
		ExcludedCriticalClasses: []string{"critical_runtime_blockers", "false_negative"},
		ReopenOnNewEvidence:     true,
		RollbackRef:             "rollback:suppression-1",
		EvidenceRefs:            []string{"signal:anomaly-1", "review:noise-window"},
		LimitationMessage:       "Suppression is reversible, bounded, and never deletes canonical evidence.",
		DeletesEvidence:         false,
	}
}

func IntelligenceCalibrationVal0FederatedBoundaryContract() FederatedSignalBoundaryContract {
	return FederatedSignalBoundaryContract{
		CurrentState:           "federated_boundary_contract_ready",
		FederatedSignalID:      "federated-signal-1",
		SourceRef:              "partner-a/staging-cluster",
		SourceTrustWeight:      60,
		EnvironmentSimilarity:  70,
		LocalOverrideAllowed:   true,
		LocalEvidenceWins:      true,
		PropagationAllowed:     false,
		PropagationNonMutating: true,
		ConfidenceCap:          IntelligenceCalibrationConfidenceMedium,
		ReasonTrace: []string{
			"federated hint remains bounded by local evidence and similarity gating",
			"propagation is disabled by default in Val 0 and cannot mutate local truth",
		},
		LimitationMessage:  "Federated signals are advisory only and cannot override local evidence or mark local assets safe automatically.",
		AutoMarksLocalSafe: false,
	}
}

func IntelligenceCalibrationVal0ProvenanceContract() CalibrationProvenanceChainContract {
	return CalibrationProvenanceChainContract{
		CurrentState:             "calibration_provenance_contract_ready",
		SupportedApprovalStates:  intelligenceCalibrationVal0ApprovalStates(),
		CalibrationChangeID:      "calibration-change-1",
		SourceSignalRefs:         []string{"signal:1", "signal:2"},
		DatasetRefs:              []string{"dataset:intelligence-calibration-val0"},
		FeedbackRefs:             []string{"feedback-fp"},
		ModelVersion:             "point5.intelligence_calibration.val0.model.v1",
		RuleVersion:              "point5.intelligence_calibration.val0.rules.v1",
		ActorOrReviewerRef:       "reviewer:security-admin",
		CreatedAt:                "2026-04-25T08:00:00Z",
		ProposedChangeSummary:    "tighten anomaly grouping threshold under reviewed bounded scope",
		ExpectedEffect:           "lower noisy false positives without hiding critical blockers",
		RiskNote:                 "bounded change could still under-surface rare negative cases and remains review constrained",
		ApprovalState:            IntelligenceCalibrationApprovalApproved,
		RollbackRef:              "rollback:calibration-change-1",
		MutatesActiveCalibration: false,
	}
}

func IntelligenceCalibrationVal0FreshnessContract() IntelligenceFreshnessExpiryContract {
	return IntelligenceFreshnessExpiryContract{
		CurrentState:              "freshness_expiry_contract_ready",
		SupportedFreshnessStates:  intelligenceCalibrationVal0FreshnessStates(),
		RequiredIntelligenceTypes: intelligenceCalibrationVal0RequiredIntelligenceTypes(),
		Items: []IntelligenceFreshnessEntry{
			{IntelligenceType: "reachability_inference", FreshnessState: IntelligenceCalibrationFreshnessFresh, LimitationMessage: "fresh reachability remains bounded by current evidence window"},
			{IntelligenceType: "vex_candidate", FreshnessState: IntelligenceCalibrationFreshnessStale, LimitationMessage: "stale VEX candidate requires explicit review before reuse"},
			{IntelligenceType: "behavioral_baseline", FreshnessState: IntelligenceCalibrationFreshnessUnknown, LimitationMessage: "unknown freshness cannot be presented as a stable baseline"},
			{IntelligenceType: "feedback_derived_tuning", FreshnessState: IntelligenceCalibrationFreshnessFresh, LimitationMessage: "feedback-derived tuning remains bounded and review constrained"},
			{IntelligenceType: "federated_signal", FreshnessState: IntelligenceCalibrationFreshnessUnsupported, LimitationMessage: "unsupported federated freshness must remain explicit"},
			{IntelligenceType: "anomaly_threshold", FreshnessState: IntelligenceCalibrationFreshnessFresh, LimitationMessage: "anomaly threshold freshness remains scoped to current bounded window"},
			{IntelligenceType: "calibration_dataset", FreshnessState: IntelligenceCalibrationFreshnessFresh, LimitationMessage: "dataset freshness is bounded to reviewed sample scope"},
			{IntelligenceType: "confidence_output", FreshnessState: IntelligenceCalibrationFreshnessExpired, ExpiresAt: "2026-04-26T08:00:00Z", LimitationMessage: "expired confidence output must not remain active"},
		},
		StaleTreatedAsFresh:               false,
		ExpiredTreatedAsActive:            false,
		UnknownFreshnessTreatedCalibrated: false,
		FreshnessAffectsConfidence:        true,
		Limitations: []string{
			"Freshness remains explicit calibration metadata and does not create a second evidence truth layer.",
		},
	}
}

func IntelligenceCalibrationVal0MetricsContract() CalibrationMetricsDefinitionContract {
	return CalibrationMetricsDefinitionContract{
		CurrentState:              "metrics_definition_contract_ready",
		SupportedSamplingPolicies: intelligenceCalibrationVal0SamplingPolicies(),
		SupportedFreshnessStates:  intelligenceCalibrationVal0FreshnessStates(),
		Items: []CalibrationMetricDefinition{
			{MetricName: IntelligenceCalibrationMetricFalsePositiveRate, Definition: "Reviewed false positives divided by reviewed signal volume in scope.", MeasurementScope: "tenant_local reviewed signals", Numerator: "reviewed_false_positive_count", Denominator: "reviewed_signal_count", SamplingPolicy: "review_window", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Metric is bounded to reviewed local signals and not global intelligence truth."}},
			{MetricName: IntelligenceCalibrationMetricFalseNegativeReviewRate, Definition: "Reviewed false-negative findings divided by reviewed missed-detection investigations.", MeasurementScope: "tenant_local missed detection reviews", Numerator: "reviewed_false_negative_count", Denominator: "missed_detection_review_count", SamplingPolicy: "manual_review_window", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Metric excludes unreviewed missed detections."}},
			{MetricName: IntelligenceCalibrationMetricOperatorAcceptanceRate, Definition: "Accepted intelligence signals divided by reviewed operator actions.", MeasurementScope: "operator-reviewed signal set", Numerator: "accepted_signal_count", Denominator: "reviewed_operator_signal_count", SamplingPolicy: "review_window", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Acceptance remains local and role-scoped."}},
			{MetricName: IntelligenceCalibrationMetricSuppressionCorrectness, Definition: "Correct suppressions divided by reviewed suppression decisions.", MeasurementScope: "reviewed suppression decisions", Numerator: "reviewed_correct_suppression_count", Denominator: "reviewed_suppression_count", SamplingPolicy: "manual_review_window", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Correctness remains bounded to reviewed suppression outcomes."}},
			{MetricName: IntelligenceCalibrationMetricMissedDetectionReviewRate, Definition: "Missed-detection investigations divided by reviewed critical low-signal cases.", MeasurementScope: "critical low-signal review queue", Numerator: "missed_detection_investigation_count", Denominator: "critical_low_signal_review_count", SamplingPolicy: "manual_review_window", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Metric requires review queue completeness to be meaningful."}},
			{MetricName: IntelligenceCalibrationMetricMeanTimeToTriage, Definition: "Average time from signal creation to first reviewed triage decision.", MeasurementScope: "reviewed triage lifecycle", Numerator: "sum_triage_durations", Denominator: "reviewed_triage_count", SamplingPolicy: "bounded_snapshot", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Time metric is operational only and not a guarantee of decision quality."}},
			{MetricName: IntelligenceCalibrationMetricTimeToDecisionDelta, Definition: "Delta between baseline and current reviewed decision time.", MeasurementScope: "bounded decision comparison window", Numerator: "decision_time_delta_sum", Denominator: "reviewed_decision_count", SamplingPolicy: "bounded_snapshot", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Delta remains bounded to compared review windows."}},
			{MetricName: IntelligenceCalibrationMetricVEXApprovalRate, Definition: "Reviewed VEX candidates approved by later governance divided by reviewed candidates.", MeasurementScope: "reviewed VEX candidate window", Numerator: "governed_vex_approval_count", Denominator: "reviewed_vex_candidate_count", SamplingPolicy: "review_window", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Rate does not make candidate output final by itself."}},
			{MetricName: IntelligenceCalibrationMetricConfidenceError, Definition: "Observed mismatch between projected confidence bands and reviewed outcomes.", MeasurementScope: "reviewed confidence outcome set", Numerator: "confidence_mismatch_count", Denominator: "reviewed_confidence_outcome_count", SamplingPolicy: "review_window", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Confidence error is bounded by reviewed outcome set only."}},
			{MetricName: IntelligenceCalibrationMetricLearningCompletion, Definition: "Reviewed learning windows completed without critical guardrail breach divided by started windows.", MeasurementScope: "learning mode review windows", Numerator: "completed_learning_window_count", Denominator: "started_learning_window_count", SamplingPolicy: "bounded_snapshot", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Metric does not authorize learning-mode enforcement relaxation."}},
			{MetricName: IntelligenceCalibrationMetricFederatedUsefulness, Definition: "Reviewed federated hints found locally useful divided by reviewed federated hints.", MeasurementScope: "reviewed federated hint set", Numerator: "locally_useful_federated_hint_count", Denominator: "reviewed_federated_hint_count", SamplingPolicy: "review_window", Freshness: IntelligenceCalibrationFreshnessFresh, Limitations: []string{"Usefulness remains environment-scoped and cannot override local evidence."}},
		},
		Limitations: []string{
			"Calibration metrics are bounded operational indicators rather than universal intelligence quality proof.",
		},
	}
}

func IntelligenceCalibrationVal0FPFNContract() FalsePositiveNegativeDisciplineContract {
	return FalsePositiveNegativeDisciplineContract{
		CurrentState:                                "false_positive_false_negative_discipline_ready",
		FalsePositiveQueueState:                     "review_open",
		FalseNegativeQueueState:                     "review_open",
		MissedDetectionScenarios:                    []string{"silent_runtime_path", "suppression_hidden_signal"},
		SuppressionCausedMissCheck:                  true,
		CriticalLowSignalReviewRequired:             true,
		ReviewCadence:                               "daily_review_window",
		EvidenceRefs:                                []string{"review_queue:fp", "review_queue:fn"},
		LimitationMessage:                           "False-positive reduction remains bounded by explicit false-negative review and missed-detection checks.",
		FalsePositiveReductionIgnoresFalseNegatives: false,
	}
}

func IntelligenceCalibrationVal0RollbackContract() CalibrationRollbackUndoContract {
	return CalibrationRollbackUndoContract{
		CurrentState:            "rollback_undo_contract_ready",
		CalibrationChangeRef:    "calibration-change-1",
		PreviewAvailable:        true,
		StagedRolloutSupported:  true,
		RollbackAvailable:       true,
		RollbackReasonRequired:  true,
		RollbackSafetyCheck:     "bounded_metric_regression_and_guardrail_recheck_required",
		BeforeMetricSnapshotRef: "metrics:before-1",
		AfterMetricSnapshotRef:  "metrics:after-1",
		LimitationMessage:       "Rollback and undo remain bounded to reviewed calibration changes and do not replace later governed rollout authority.",
	}
}

func EvaluateIntelligenceCalibrationVal0DatasetState(model CalibrationDatasetContract) string {
	if strings.TrimSpace(model.DatasetID) == "" || strings.TrimSpace(model.DatasetVersion) == "" || strings.TrimSpace(model.DatasetScope) == "" || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.SamplingPolicy) == "" || strings.TrimSpace(model.GeneratedAt) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationVal0DatasetStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.ScenarioClasses, intelligenceCalibrationVal0ScenarioClasses()...) || len(model.SourceRefs) == 0 || len(model.EvidenceRefs) == 0 || len(model.Limitations) == 0 {
		return IntelligenceCalibrationVal0DatasetStatePartial
	}
	if !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || model.LabeledCaseCount < 0 || model.ReviewedCaseCount < 0 || model.ReviewedCaseCount > model.LabeledCaseCount {
		return IntelligenceCalibrationVal0DatasetStatePartial
	}
	if !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationVal0DatasetStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessStale && !containsSubstringInTrimmedStrings(model.Limitations, "stale") {
		return IntelligenceCalibrationVal0DatasetStatePartial
	}
	return IntelligenceCalibrationVal0DatasetStateActive
}

func EvaluateIntelligenceCalibrationVal0ConfidenceState(model ConfidenceEvidenceClassContract) string {
	if strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.EvidenceClass) == "" || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.ModelVersion) == "" || strings.TrimSpace(model.RuleVersion) == "" || strings.TrimSpace(model.DatasetVersion) == "" || len(model.ReasonTrace) == 0 || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationVal0ConfidenceStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedConfidenceBands, intelligenceCalibrationVal0ConfidenceBands()...) || !containsExactTrimmedStringSet(model.SupportedEvidenceClasses, intelligenceCalibrationVal0EvidenceClasses()...) || !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || !model.AdvisoryOnly {
		return IntelligenceCalibrationVal0ConfidenceStatePartial
	}
	if !containsTrimmedString(model.SupportedConfidenceBands, model.ConfidenceBand) || !containsTrimmedString(model.SupportedEvidenceClasses, model.EvidenceClass) || model.ConfidenceScore < 0 || model.ConfidenceScore > 100 {
		return IntelligenceCalibrationVal0ConfidenceStatePartial
	}
	if model.EvidenceClass == IntelligenceCalibrationEvidenceUnsupported && model.ConfidenceBand == IntelligenceCalibrationConfidenceHigh {
		return IntelligenceCalibrationVal0ConfidenceStatePartial
	}
	if model.ConfidenceBand == IntelligenceCalibrationConfidenceUnknown {
		return IntelligenceCalibrationVal0ConfidenceStatePartial
	}
	return IntelligenceCalibrationVal0ConfidenceStateActive
}

func EvaluateIntelligenceCalibrationVal0LifecycleState(model IntelligenceOutputLifecycleContract) string {
	if len(model.Items) == 0 {
		return IntelligenceCalibrationVal0LifecycleStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedLifecycleStates, intelligenceCalibrationVal0LifecycleStates()...) || model.CandidateTreatedAsAccepted || model.ReviewRequiredTreatedApproved || model.ExpiredTreatedAsActive || !model.AdvisoryOnly {
		return IntelligenceCalibrationVal0LifecycleStatePartial
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		state := strings.TrimSpace(item.LifecycleState)
		if strings.TrimSpace(item.OutputID) == "" || strings.TrimSpace(item.OutputType) == "" || state == "" || strings.TrimSpace(item.CreatedAt) == "" || strings.TrimSpace(item.UpdatedAt) == "" || len(item.EvidenceRefs) == 0 || strings.TrimSpace(item.ReasonCode) == "" || strings.TrimSpace(item.LimitationMessage) == "" {
			return IntelligenceCalibrationVal0LifecycleStateIncomplete
		}
		if !containsTrimmedString(model.SupportedLifecycleStates, state) {
			return IntelligenceCalibrationVal0LifecycleStatePartial
		}
		if _, duplicate := seen[state]; duplicate {
			return IntelligenceCalibrationVal0LifecycleStatePartial
		}
		seen[state] = struct{}{}
		if state == IntelligenceCalibrationLifecycleExpired && strings.TrimSpace(item.ExpiresAt) == "" {
			return IntelligenceCalibrationVal0LifecycleStatePartial
		}
		if state == IntelligenceCalibrationLifecycleSuperseded && strings.TrimSpace(item.SupersedesRef) == "" {
			return IntelligenceCalibrationVal0LifecycleStatePartial
		}
		if (state == IntelligenceCalibrationLifecycleAccepted || state == IntelligenceCalibrationLifecycleRejected) && strings.TrimSpace(item.ReviewerRef) == "" {
			return IntelligenceCalibrationVal0LifecycleStatePartial
		}
	}
	if len(seen) != len(model.SupportedLifecycleStates) {
		return IntelligenceCalibrationVal0LifecycleStatePartial
	}
	return IntelligenceCalibrationVal0LifecycleStateActive
}

func EvaluateIntelligenceCalibrationVal0ReachabilityState(model ReachabilityTaxonomyContract) string {
	if strings.TrimSpace(model.Example.ReachabilityClass) == "" || strings.TrimSpace(model.Example.PackageRef) == "" || strings.TrimSpace(model.Example.FunctionOrComponentRef) == "" || strings.TrimSpace(model.Example.WorkloadContextRef) == "" || strings.TrimSpace(model.Example.ConfidenceBand) == "" || strings.TrimSpace(model.Example.EvidenceClass) == "" || strings.TrimSpace(model.Example.Explanation) == "" {
		return IntelligenceCalibrationVal0ReachabilityStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReachabilityClasses, intelligenceCalibrationVal0ReachabilityClasses()...) || !containsExactTrimmedStringSet(model.SupportedConfidenceBands, intelligenceCalibrationVal0ConfidenceBands()...) || !model.AdvisoryOnly {
		return IntelligenceCalibrationVal0ReachabilityStatePartial
	}
	if !containsTrimmedString(model.SupportedReachabilityClasses, model.Example.ReachabilityClass) || !containsTrimmedString(model.SupportedConfidenceBands, model.Example.ConfidenceBand) || !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), model.Example.EvidenceClass) {
		return IntelligenceCalibrationVal0ReachabilityStatePartial
	}
	if model.PackagePresenceImpliesExploit || model.NoObservedExecutionImpliesSafe {
		return IntelligenceCalibrationVal0ReachabilityStatePartial
	}
	if model.Example.DowngradeAllowed && (!model.Example.ExplanationRequired || strings.TrimSpace(model.Example.EvidenceClass) == "" || strings.TrimSpace(model.Example.Explanation) == "") {
		return IntelligenceCalibrationVal0ReachabilityStatePartial
	}
	return IntelligenceCalibrationVal0ReachabilityStateActive
}

func EvaluateIntelligenceCalibrationVal0VEXState(model VEXCandidateGovernanceContract) string {
	if strings.TrimSpace(model.Example.CandidateID) == "" || strings.TrimSpace(model.Example.VulnerabilityRef) == "" || strings.TrimSpace(model.Example.ProductOrAssetRef) == "" || strings.TrimSpace(model.Example.SuggestedOutcome) == "" || strings.TrimSpace(model.Example.CandidateState) == "" || len(model.Example.EvidenceRefs) == 0 || strings.TrimSpace(model.Example.ReachabilityRef) == "" || strings.TrimSpace(model.Example.ConfidenceBand) == "" || strings.TrimSpace(model.Example.Expiry) == "" || strings.TrimSpace(model.Example.ReasonCode) == "" || strings.TrimSpace(model.Example.LimitationMessage) == "" {
		return IntelligenceCalibrationVal0VEXStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedStates, intelligenceCalibrationVal0VEXStates()...) || !containsExactTrimmedStringSet(model.SupportedSuggestedOutcomes, intelligenceCalibrationVal0VEXOutcomes()...) || model.CandidateIsFinalVEX || model.InsufficientEvidenceBecomesSafe || !model.AdvisoryOnly {
		return IntelligenceCalibrationVal0VEXStatePartial
	}
	if !containsTrimmedString(model.SupportedStates, model.Example.CandidateState) || !containsTrimmedString(model.SupportedSuggestedOutcomes, model.Example.SuggestedOutcome) || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.Example.ConfidenceBand) {
		return IntelligenceCalibrationVal0VEXStatePartial
	}
	return IntelligenceCalibrationVal0VEXStateActive
}

func EvaluateIntelligenceCalibrationVal0FeedbackState(model FeedbackClassificationContract) string {
	if len(model.Items) == 0 {
		return IntelligenceCalibrationVal0FeedbackStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedFeedbackClasses, intelligenceCalibrationVal0FeedbackClasses()...) || model.DirectMutationAllowed || model.DirectSuppressionAllowed || model.FederatedPropagationAllowed || !model.AdvisoryOnly {
		return IntelligenceCalibrationVal0FeedbackStatePartial
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		feedbackClass := strings.TrimSpace(item.FeedbackClass)
		if strings.TrimSpace(item.FeedbackID) == "" || strings.TrimSpace(item.ActorRef) == "" || strings.TrimSpace(item.SignalRef) == "" || feedbackClass == "" || strings.TrimSpace(item.Reason) == "" || strings.TrimSpace(item.LocalApplicability) == "" || strings.TrimSpace(item.CreatedAt) == "" {
			return IntelligenceCalibrationVal0FeedbackStateIncomplete
		}
		if !containsTrimmedString(model.SupportedFeedbackClasses, feedbackClass) {
			return IntelligenceCalibrationVal0FeedbackStatePartial
		}
		if _, duplicate := seen[feedbackClass]; duplicate {
			return IntelligenceCalibrationVal0FeedbackStatePartial
		}
		seen[feedbackClass] = struct{}{}
		if !item.ReviewRequired {
			return IntelligenceCalibrationVal0FeedbackStatePartial
		}
	}
	if len(seen) != len(model.SupportedFeedbackClasses) {
		return IntelligenceCalibrationVal0FeedbackStatePartial
	}
	return IntelligenceCalibrationVal0FeedbackStateActive
}

func EvaluateIntelligenceCalibrationVal0LearningModeState(model LearningModeGuardrailContract) string {
	if strings.TrimSpace(model.LearningModeState) == "" || strings.TrimSpace(model.Scope) == "" || strings.TrimSpace(model.StartedAt) == "" || strings.TrimSpace(model.ExpiresAt) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationVal0LearningModeStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedStates, intelligenceCalibrationVal0LearningStates()...) || !containsTrimmedString(model.SupportedStates, model.LearningModeState) || len(model.ExcludedCriticalControls) == 0 || len(model.AllowedSignalClasses) == 0 || len(model.DisallowedSignalClasses) == 0 {
		return IntelligenceCalibrationVal0LearningModeStatePartial
	}
	startedAt, err := time.Parse(time.RFC3339, strings.TrimSpace(model.StartedAt))
	if err != nil {
		return IntelligenceCalibrationVal0LearningModeStatePartial
	}
	expiresAt, err := time.Parse(time.RFC3339, strings.TrimSpace(model.ExpiresAt))
	if err != nil {
		return IntelligenceCalibrationVal0LearningModeStatePartial
	}
	if !model.OutputReviewRequired || model.CanRelaxEnforcement || model.LearningModeState == IntelligenceCalibrationLearningExpired || !expiresAt.After(startedAt) {
		return IntelligenceCalibrationVal0LearningModeStatePartial
	}
	return IntelligenceCalibrationVal0LearningModeStateActive
}

func EvaluateIntelligenceCalibrationVal0SuppressionState(model SuppressionSafetyContract) string {
	if strings.TrimSpace(model.SuppressionID) == "" || strings.TrimSpace(model.SignalClass) == "" || strings.TrimSpace(model.SuppressionScope) == "" || strings.TrimSpace(model.ReasonCode) == "" || strings.TrimSpace(model.CreatedBy) == "" || strings.TrimSpace(model.ReviewerRef) == "" || strings.TrimSpace(model.ExpiresAt) == "" || strings.TrimSpace(model.RollbackRef) == "" || strings.TrimSpace(model.LimitationMessage) == "" || len(model.AffectedSubjects) == 0 || len(model.EvidenceRefs) == 0 {
		return IntelligenceCalibrationVal0SuppressionStateIncomplete
	}
	if !containsTrimmedString(intelligenceCalibrationVal0ScenarioClasses(), model.SignalClass) || !model.ReopenOnNewEvidence || model.DeletesEvidence {
		return IntelligenceCalibrationVal0SuppressionStatePartial
	}
	if containsTrimmedString(model.ExcludedCriticalClasses, model.SignalClass) {
		return IntelligenceCalibrationVal0SuppressionStatePartial
	}
	return IntelligenceCalibrationVal0SuppressionStateActive
}

func EvaluateIntelligenceCalibrationVal0FederatedBoundaryState(model FederatedSignalBoundaryContract) string {
	if strings.TrimSpace(model.FederatedSignalID) == "" || strings.TrimSpace(model.SourceRef) == "" || strings.TrimSpace(model.ConfidenceCap) == "" || len(model.ReasonTrace) == 0 || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationVal0FederatedBoundaryStateIncomplete
	}
	if model.SourceTrustWeight <= 0 || model.EnvironmentSimilarity <= 0 || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceCap) {
		return IntelligenceCalibrationVal0FederatedBoundaryStatePartial
	}
	if !model.LocalOverrideAllowed || !model.LocalEvidenceWins || model.AutoMarksLocalSafe {
		return IntelligenceCalibrationVal0FederatedBoundaryStatePartial
	}
	if model.PropagationAllowed && !model.PropagationNonMutating {
		return IntelligenceCalibrationVal0FederatedBoundaryStatePartial
	}
	return IntelligenceCalibrationVal0FederatedBoundaryStateActive
}

func EvaluateIntelligenceCalibrationVal0ProvenanceState(model CalibrationProvenanceChainContract) string {
	if strings.TrimSpace(model.CalibrationChangeID) == "" || len(model.SourceSignalRefs) == 0 || len(model.DatasetRefs) == 0 || strings.TrimSpace(model.ModelVersion) == "" || strings.TrimSpace(model.RuleVersion) == "" || strings.TrimSpace(model.ActorOrReviewerRef) == "" || strings.TrimSpace(model.CreatedAt) == "" || strings.TrimSpace(model.ProposedChangeSummary) == "" || strings.TrimSpace(model.ExpectedEffect) == "" || strings.TrimSpace(model.RiskNote) == "" || strings.TrimSpace(model.ApprovalState) == "" {
		return IntelligenceCalibrationVal0ProvenanceStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedApprovalStates, intelligenceCalibrationVal0ApprovalStates()...) || !containsTrimmedString(model.SupportedApprovalStates, model.ApprovalState) {
		return IntelligenceCalibrationVal0ProvenanceStatePartial
	}
	if model.ApprovalState == IntelligenceCalibrationApprovalProposed || model.ApprovalState == IntelligenceCalibrationApprovalReviewRequired || model.MutatesActiveCalibration {
		return IntelligenceCalibrationVal0ProvenanceStatePartial
	}
	if model.ApprovalState == IntelligenceCalibrationApprovalApproved && strings.TrimSpace(model.RollbackRef) == "" {
		return IntelligenceCalibrationVal0ProvenanceStatePartial
	}
	return IntelligenceCalibrationVal0ProvenanceStateActive
}

func EvaluateIntelligenceCalibrationVal0FreshnessState(model IntelligenceFreshnessExpiryContract) string {
	if len(model.Items) == 0 {
		return IntelligenceCalibrationVal0FreshnessStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedFreshnessStates, intelligenceCalibrationVal0FreshnessStates()...) || !containsExactTrimmedStringSet(model.RequiredIntelligenceTypes, intelligenceCalibrationVal0RequiredIntelligenceTypes()...) || model.StaleTreatedAsFresh || model.ExpiredTreatedAsActive || model.UnknownFreshnessTreatedCalibrated || !model.FreshnessAffectsConfidence {
		return IntelligenceCalibrationVal0FreshnessStatePartial
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		kind := strings.TrimSpace(item.IntelligenceType)
		state := strings.TrimSpace(item.FreshnessState)
		if kind == "" || state == "" {
			return IntelligenceCalibrationVal0FreshnessStateIncomplete
		}
		if !containsTrimmedString(model.RequiredIntelligenceTypes, kind) || !containsTrimmedString(model.SupportedFreshnessStates, state) {
			return IntelligenceCalibrationVal0FreshnessStatePartial
		}
		if _, duplicate := seen[kind]; duplicate {
			return IntelligenceCalibrationVal0FreshnessStatePartial
		}
		seen[kind] = struct{}{}
		if (state == IntelligenceCalibrationFreshnessStale || state == IntelligenceCalibrationFreshnessUnknown || state == IntelligenceCalibrationFreshnessUnsupported) && strings.TrimSpace(item.LimitationMessage) == "" {
			return IntelligenceCalibrationVal0FreshnessStatePartial
		}
		if state == IntelligenceCalibrationFreshnessExpired && (strings.TrimSpace(item.ExpiresAt) == "" || strings.TrimSpace(item.LimitationMessage) == "") {
			return IntelligenceCalibrationVal0FreshnessStatePartial
		}
	}
	if len(seen) != len(model.RequiredIntelligenceTypes) {
		return IntelligenceCalibrationVal0FreshnessStatePartial
	}
	return IntelligenceCalibrationVal0FreshnessStateActive
}

func EvaluateIntelligenceCalibrationVal0MetricsState(model CalibrationMetricsDefinitionContract) string {
	if len(model.Items) == 0 {
		return IntelligenceCalibrationVal0MetricsStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedSamplingPolicies, intelligenceCalibrationVal0SamplingPolicies()...) || !containsExactTrimmedStringSet(model.SupportedFreshnessStates, intelligenceCalibrationVal0FreshnessStates()...) {
		return IntelligenceCalibrationVal0MetricsStatePartial
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		metricName := strings.TrimSpace(item.MetricName)
		if metricName == "" || strings.TrimSpace(item.Definition) == "" || strings.TrimSpace(item.MeasurementScope) == "" || strings.TrimSpace(item.Numerator) == "" || strings.TrimSpace(item.Denominator) == "" || strings.TrimSpace(item.SamplingPolicy) == "" || strings.TrimSpace(item.Freshness) == "" {
			return IntelligenceCalibrationVal0MetricsStateIncomplete
		}
		if !containsTrimmedString(intelligenceCalibrationVal0RequiredMetrics(), metricName) || !containsTrimmedString(model.SupportedSamplingPolicies, item.SamplingPolicy) || !containsTrimmedString(model.SupportedFreshnessStates, item.Freshness) || len(item.Limitations) == 0 {
			return IntelligenceCalibrationVal0MetricsStatePartial
		}
		if _, duplicate := seen[metricName]; duplicate {
			return IntelligenceCalibrationVal0MetricsStatePartial
		}
		seen[metricName] = struct{}{}
	}
	if len(seen) != len(intelligenceCalibrationVal0RequiredMetrics()) {
		return IntelligenceCalibrationVal0MetricsStatePartial
	}
	return IntelligenceCalibrationVal0MetricsStateActive
}

func EvaluateIntelligenceCalibrationVal0FPFNState(model FalsePositiveNegativeDisciplineContract) string {
	if strings.TrimSpace(model.FalsePositiveQueueState) == "" || strings.TrimSpace(model.FalseNegativeQueueState) == "" || strings.TrimSpace(model.ReviewCadence) == "" || strings.TrimSpace(model.LimitationMessage) == "" || len(model.EvidenceRefs) == 0 {
		return IntelligenceCalibrationVal0FPFNStateIncomplete
	}
	if len(model.MissedDetectionScenarios) == 0 || !model.SuppressionCausedMissCheck || !model.CriticalLowSignalReviewRequired || model.FalsePositiveReductionIgnoresFalseNegatives {
		return IntelligenceCalibrationVal0FPFNStatePartial
	}
	return IntelligenceCalibrationVal0FPFNStateActive
}

func EvaluateIntelligenceCalibrationVal0RollbackState(model CalibrationRollbackUndoContract) string {
	if strings.TrimSpace(model.CalibrationChangeRef) == "" || strings.TrimSpace(model.RollbackSafetyCheck) == "" {
		return IntelligenceCalibrationVal0RollbackStateIncomplete
	}
	if !model.RollbackReasonRequired {
		return IntelligenceCalibrationVal0RollbackStatePartial
	}
	if !model.RollbackAvailable && strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationVal0RollbackStatePartial
	}
	if (strings.TrimSpace(model.BeforeMetricSnapshotRef) == "" || strings.TrimSpace(model.AfterMetricSnapshotRef) == "") && strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationVal0RollbackStatePartial
	}
	return IntelligenceCalibrationVal0RollbackStateActive
}

func EvaluateIntelligenceCalibrationVal0State(datasetState, confidenceState, lifecycleState, reachabilityState, vexState, feedbackState, learningModeState, suppressionState, federatedBoundaryState, provenanceState, freshnessState, metricsState, fpfnState, rollbackState string) string {
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(datasetState),
		strings.TrimSpace(confidenceState),
		strings.TrimSpace(lifecycleState),
		strings.TrimSpace(reachabilityState),
		strings.TrimSpace(vexState),
		strings.TrimSpace(feedbackState),
		strings.TrimSpace(learningModeState),
		strings.TrimSpace(suppressionState),
		strings.TrimSpace(federatedBoundaryState),
		strings.TrimSpace(provenanceState),
		strings.TrimSpace(freshnessState),
		strings.TrimSpace(metricsState),
		strings.TrimSpace(fpfnState),
		strings.TrimSpace(rollbackState),
	} {
		switch state {
		case IntelligenceCalibrationVal0DatasetStateActive,
			IntelligenceCalibrationVal0ConfidenceStateActive,
			IntelligenceCalibrationVal0LifecycleStateActive,
			IntelligenceCalibrationVal0ReachabilityStateActive,
			IntelligenceCalibrationVal0VEXStateActive,
			IntelligenceCalibrationVal0FeedbackStateActive,
			IntelligenceCalibrationVal0LearningModeStateActive,
			IntelligenceCalibrationVal0SuppressionStateActive,
			IntelligenceCalibrationVal0FederatedBoundaryStateActive,
			IntelligenceCalibrationVal0ProvenanceStateActive,
			IntelligenceCalibrationVal0FreshnessStateActive,
			IntelligenceCalibrationVal0MetricsStateActive,
			IntelligenceCalibrationVal0FPFNStateActive,
			IntelligenceCalibrationVal0RollbackStateActive:
		case IntelligenceCalibrationVal0DatasetStatePartial,
			IntelligenceCalibrationVal0ConfidenceStatePartial,
			IntelligenceCalibrationVal0LifecycleStatePartial,
			IntelligenceCalibrationVal0ReachabilityStatePartial,
			IntelligenceCalibrationVal0VEXStatePartial,
			IntelligenceCalibrationVal0FeedbackStatePartial,
			IntelligenceCalibrationVal0LearningModeStatePartial,
			IntelligenceCalibrationVal0SuppressionStatePartial,
			IntelligenceCalibrationVal0FederatedBoundaryStatePartial,
			IntelligenceCalibrationVal0ProvenanceStatePartial,
			IntelligenceCalibrationVal0FreshnessStatePartial,
			IntelligenceCalibrationVal0MetricsStatePartial,
			IntelligenceCalibrationVal0FPFNStatePartial,
			IntelligenceCalibrationVal0RollbackStatePartial:
			hasPartial = true
		default:
			return IntelligenceCalibrationVal0StateIncomplete
		}
	}
	if hasPartial {
		return IntelligenceCalibrationVal0StateSubstantial
	}
	return IntelligenceCalibrationVal0StateActive
}

func EvaluateIntelligenceCalibrationVal0ProofsState(datasetState, confidenceState, lifecycleState, reachabilityState, vexState, feedbackState, learningModeState, suppressionState, federatedBoundaryState, provenanceState, freshnessState, metricsState, fpfnState, rollbackState string, surfaceRefs, evidenceRefs, limitations, whyPoint5NotPass []string, projectionDisclaimer string) string {
	baseState := EvaluateIntelligenceCalibrationVal0State(datasetState, confidenceState, lifecycleState, reachabilityState, vexState, feedbackState, learningModeState, suppressionState, federatedBoundaryState, provenanceState, freshnessState, metricsState, fpfnState, rollbackState)
	if len(surfaceRefs) < 15 || len(evidenceRefs) < 8 || len(limitations) == 0 || len(whyPoint5NotPass) == 0 || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "not_canonical_truth") {
		if baseState == IntelligenceCalibrationVal0StateActive {
			return IntelligenceCalibrationVal0StateSubstantial
		}
		return baseState
	}
	return baseState
}

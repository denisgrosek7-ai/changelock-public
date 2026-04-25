package operability

import (
	"strings"
	"time"
)

const (
	IntelligenceCalibrationValBBehavioralBaselineStateActive     = "intelligence_calibration_valb_behavioral_baseline_active"
	IntelligenceCalibrationValBBehavioralBaselineStatePartial    = "intelligence_calibration_valb_behavioral_baseline_partial"
	IntelligenceCalibrationValBBehavioralBaselineStateIncomplete = "intelligence_calibration_valb_behavioral_baseline_incomplete"

	IntelligenceCalibrationValBLearningRuntimeStateActive     = "intelligence_calibration_valb_learning_mode_runtime_active"
	IntelligenceCalibrationValBLearningRuntimeStatePartial    = "intelligence_calibration_valb_learning_mode_runtime_partial"
	IntelligenceCalibrationValBLearningRuntimeStateIncomplete = "intelligence_calibration_valb_learning_mode_runtime_incomplete"

	IntelligenceCalibrationValBThresholdStateActive     = "intelligence_calibration_valb_anomaly_threshold_active"
	IntelligenceCalibrationValBThresholdStatePartial    = "intelligence_calibration_valb_anomaly_threshold_partial"
	IntelligenceCalibrationValBThresholdStateIncomplete = "intelligence_calibration_valb_anomaly_threshold_incomplete"

	IntelligenceCalibrationValBDriftStateActive     = "intelligence_calibration_valb_drift_sensitivity_active"
	IntelligenceCalibrationValBDriftStatePartial    = "intelligence_calibration_valb_drift_sensitivity_partial"
	IntelligenceCalibrationValBDriftStateIncomplete = "intelligence_calibration_valb_drift_sensitivity_incomplete"

	IntelligenceCalibrationValBWeightingStateActive     = "intelligence_calibration_valb_criticality_weighting_active"
	IntelligenceCalibrationValBWeightingStatePartial    = "intelligence_calibration_valb_criticality_weighting_partial"
	IntelligenceCalibrationValBWeightingStateIncomplete = "intelligence_calibration_valb_criticality_weighting_incomplete"

	IntelligenceCalibrationValBBaselineFreshnessStateActive     = "intelligence_calibration_valb_baseline_freshness_active"
	IntelligenceCalibrationValBBaselineFreshnessStatePartial    = "intelligence_calibration_valb_baseline_freshness_partial"
	IntelligenceCalibrationValBBaselineFreshnessStateIncomplete = "intelligence_calibration_valb_baseline_freshness_incomplete"

	IntelligenceCalibrationValBBaselineAdoptionStateActive     = "intelligence_calibration_valb_baseline_adoption_active"
	IntelligenceCalibrationValBBaselineAdoptionStatePartial    = "intelligence_calibration_valb_baseline_adoption_partial"
	IntelligenceCalibrationValBBaselineAdoptionStateIncomplete = "intelligence_calibration_valb_baseline_adoption_incomplete"

	IntelligenceCalibrationValBExplanationStateActive     = "intelligence_calibration_valb_behavioral_explanation_active"
	IntelligenceCalibrationValBExplanationStatePartial    = "intelligence_calibration_valb_behavioral_explanation_partial"
	IntelligenceCalibrationValBExplanationStateIncomplete = "intelligence_calibration_valb_behavioral_explanation_incomplete"

	IntelligenceCalibrationValBGuardrailStateActive     = "intelligence_calibration_valb_behavioral_guardrail_active"
	IntelligenceCalibrationValBGuardrailStatePartial    = "intelligence_calibration_valb_behavioral_guardrail_partial"
	IntelligenceCalibrationValBGuardrailStateIncomplete = "intelligence_calibration_valb_behavioral_guardrail_incomplete"

	IntelligenceCalibrationValBStateIncomplete  = "intelligence_calibration_valb_incomplete"
	IntelligenceCalibrationValBStateSubstantial = "intelligence_calibration_valb_substantially_ready"
	IntelligenceCalibrationValBStateActive      = "intelligence_calibration_valb_active"

	IntelligenceCalibrationValBThresholdIncreaseSensitivity = "increase_sensitivity"
	IntelligenceCalibrationValBThresholdDecreaseSensitivity = "decrease_sensitivity"
	IntelligenceCalibrationValBThresholdKeep                = "keep"
	IntelligenceCalibrationValBThresholdRequiresReview      = "requires_review"

	IntelligenceCalibrationValBDriftAdjustmentIncrease       = "increase"
	IntelligenceCalibrationValBDriftAdjustmentDecrease       = "decrease"
	IntelligenceCalibrationValBDriftAdjustmentKeep           = "keep"
	IntelligenceCalibrationValBDriftAdjustmentRequiresReview = "requires_review"

	IntelligenceCalibrationValBCriticalityCritical = "critical"
	IntelligenceCalibrationValBCriticalityHigh     = "high"
	IntelligenceCalibrationValBCriticalityNormal   = "normal"
	IntelligenceCalibrationValBCriticalityLow      = "low"
	IntelligenceCalibrationValBCriticalityUnknown  = "unknown"

	IntelligenceCalibrationValBWeightingRaisePriority          = "raise_priority"
	IntelligenceCalibrationValBWeightingLowerPriorityCandidate = "lower_priority_candidate"
	IntelligenceCalibrationValBWeightingKeep                   = "keep"
	IntelligenceCalibrationValBWeightingRequiresReview         = "requires_review"

	IntelligenceCalibrationValBUnsupportedSignal = "unsupported_signal"
)

type BehavioralBaselineProfileContract struct {
	CurrentState                 string   `json:"current_state"`
	SupportedSignalClasses       []string `json:"supported_signal_classes,omitempty"`
	BaselineID                   string   `json:"baseline_id"`
	BaselineVersion              string   `json:"baseline_version"`
	BaselineScope                string   `json:"baseline_scope"`
	WorkloadContextRefs          []string `json:"workload_context_refs,omitempty"`
	AssetOrServiceRefs           []string `json:"asset_or_service_refs,omitempty"`
	ObservedSignalClasses        []string `json:"observed_signal_classes,omitempty"`
	ObservedProcessClasses       []string `json:"observed_process_classes,omitempty"`
	ObservedNetworkPaths         []string `json:"observed_network_paths,omitempty"`
	ObservedFileActivityClasses  []string `json:"observed_file_activity_classes,omitempty"`
	ObservedRuntimeSignalClasses []string `json:"observed_runtime_signal_classes,omitempty"`
	ObservationWindowStart       string   `json:"observation_window_start"`
	ObservationWindowEnd         string   `json:"observation_window_end"`
	SampleCount                  int      `json:"sample_count"`
	OutlierHandlingPolicy        string   `json:"outlier_handling_policy"`
	FreshnessState               string   `json:"freshness_state"`
	EvidenceRefs                 []string `json:"evidence_refs,omitempty"`
	LimitationMessage            string   `json:"limitation_message"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
	CanRelaxEnforcement          bool     `json:"can_relax_enforcement"`
	CanSuppressAlerts            bool     `json:"can_suppress_alerts"`
	AdvisoryOnly                 bool     `json:"advisory_only"`
}

type LearningModeRuntimeDisciplineContract struct {
	CurrentState              string   `json:"current_state"`
	SupportedStates           []string `json:"supported_states,omitempty"`
	SupportedSignalClasses    []string `json:"supported_signal_classes,omitempty"`
	LearningSessionID         string   `json:"learning_session_id"`
	LearningModeState         string   `json:"learning_mode_state"`
	Scope                     string   `json:"scope"`
	StartedAt                 string   `json:"started_at"`
	ExpiresAt                 string   `json:"expires_at"`
	BoundedDurationConfirmed  bool     `json:"bounded_duration_confirmed"`
	ExcludedCriticalControls  []string `json:"excluded_critical_controls,omitempty"`
	ObservedSignalClasses     []string `json:"observed_signal_classes,omitempty"`
	DisallowedSignalClasses   []string `json:"disallowed_signal_classes,omitempty"`
	OutputReviewRequired      bool     `json:"output_review_required"`
	CanRelaxEnforcement       bool     `json:"can_relax_enforcement"`
	CanAutoPromoteBaseline    bool     `json:"can_auto_promote_baseline"`
	CanSuppressCriticalAlerts bool     `json:"can_suppress_critical_alerts"`
	ReviewSummaryRequired     bool     `json:"review_summary_required"`
	LimitationMessage         string   `json:"limitation_message"`
}

type AnomalyThresholdCalibrationContract struct {
	CurrentState             string   `json:"current_state"`
	SupportedDirections      []string `json:"supported_threshold_change_directions,omitempty"`
	ThresholdProfileID       string   `json:"threshold_profile_id"`
	SignalClass              string   `json:"signal_class"`
	BaselineRef              string   `json:"baseline_ref"`
	CurrentThreshold         float64  `json:"current_threshold"`
	ProposedThreshold        float64  `json:"proposed_threshold"`
	ThresholdChangeDirection string   `json:"threshold_change_direction"`
	ConfidenceBand           string   `json:"confidence_band"`
	EvidenceClass            string   `json:"evidence_class"`
	FalsePositiveRiskNote    string   `json:"false_positive_risk_note"`
	FalseNegativeRiskNote    string   `json:"false_negative_risk_note"`
	ReviewRequired           bool     `json:"review_required"`
	AppliesToCriticalClass   bool     `json:"applies_to_critical_class"`
	LimitationMessage        string   `json:"limitation_message"`
	AdvisoryOnly             bool     `json:"advisory_only"`
	MutatesActiveDetection   bool     `json:"mutates_active_detection_state"`
}

type DriftSensitivityScalingContract struct {
	CurrentState           string   `json:"current_state"`
	SupportedDriftBands    []string `json:"supported_drift_score_bands,omitempty"`
	SupportedAdjustments   []string `json:"supported_sensitivity_adjustments,omitempty"`
	SupportedCriticalities []string `json:"supported_criticality_contexts,omitempty"`
	DriftProfileID         string   `json:"drift_profile_id"`
	BaselineRef            string   `json:"baseline_ref"`
	SignalClass            string   `json:"signal_class"`
	DriftScoreBand         string   `json:"drift_score_band"`
	SensitivityAdjustment  string   `json:"sensitivity_adjustment"`
	CriticalityContext     string   `json:"criticality_context"`
	EvidenceRefs           []string `json:"evidence_refs,omitempty"`
	FreshnessState         string   `json:"freshness_state"`
	ReviewRequired         bool     `json:"review_required"`
	LimitationMessage      string   `json:"limitation_message"`
	AdvisoryOnly           bool     `json:"advisory_only"`
	MutatesEnforcement     bool     `json:"mutates_enforcement_state"`
}

type CriticalityAwareWeightingContract struct {
	CurrentState             string   `json:"current_state"`
	SupportedCriticalities   []string `json:"supported_criticality_classes,omitempty"`
	SupportedWeighting       []string `json:"supported_weighting_actions,omitempty"`
	WeightingProfileID       string   `json:"weighting_profile_id"`
	AssetOrServiceRef        string   `json:"asset_or_service_ref"`
	CriticalityClass         string   `json:"criticality_class"`
	WeightingAction          string   `json:"weighting_action"`
	ConfidenceBand           string   `json:"confidence_band"`
	EvidenceClass            string   `json:"evidence_class"`
	LocalContextRefs         []string `json:"local_context_refs,omitempty"`
	BlastRadiusHint          string   `json:"blast_radius_hint"`
	ReviewerRequired         bool     `json:"reviewer_required"`
	LimitationMessage        string   `json:"limitation_message"`
	ReasonCode               string   `json:"reason_code"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	AdvisoryOnly             bool     `json:"advisory_only"`
	MutatesCanonicalPriority bool     `json:"mutates_canonical_priority"`
}

type BaselineFreshnessExpiryContract struct {
	CurrentState      string `json:"current_state"`
	BaselineRef       string `json:"baseline_ref"`
	FreshnessState    string `json:"freshness_state"`
	LastObservedAt    string `json:"last_observed_at"`
	ExpiresAt         string `json:"expires_at"`
	FreshnessLimit    string `json:"freshness_limit"`
	StaleReason       string `json:"stale_reason,omitempty"`
	ExpiryReason      string `json:"expiry_reason,omitempty"`
	LimitationMessage string `json:"limitation_message"`
	AdvisoryOnly      bool   `json:"advisory_only"`
}

type BaselineAdoptionReviewContract struct {
	CurrentState            string   `json:"current_state"`
	SupportedAdoptionStates []string `json:"supported_adoption_states,omitempty"`
	AdoptionID              string   `json:"adoption_id"`
	BaselineRef             string   `json:"baseline_ref"`
	ProposedAdoptionState   string   `json:"proposed_adoption_state"`
	ReviewerRequired        bool     `json:"reviewer_required"`
	ApprovalRef             string   `json:"approval_ref,omitempty"`
	OutputSummaryRef        string   `json:"output_summary_ref"`
	RiskNote                string   `json:"risk_note"`
	RollbackRef             string   `json:"rollback_ref"`
	BeforeMetricSnapshotRef string   `json:"before_metric_snapshot_ref"`
	AfterMetricSnapshotRef  string   `json:"after_metric_snapshot_ref"`
	LimitationMessage       string   `json:"limitation_message"`
	MutatesActiveBaseline   bool     `json:"mutates_active_baseline"`
	GovernanceRequired      bool     `json:"governance_required_for_mutation"`
	AdvisoryOnly            bool     `json:"advisory_only"`
}

type BehavioralCalibrationExplanationContract struct {
	CurrentState                    string   `json:"current_state"`
	ReasonCode                      string   `json:"reason_code"`
	HumanMessage                    string   `json:"human_message"`
	TechnicalDetail                 string   `json:"technical_detail"`
	BaselineRef                     string   `json:"baseline_ref"`
	SignalClass                     string   `json:"signal_class"`
	ConfidenceBand                  string   `json:"confidence_band"`
	EvidenceClass                   string   `json:"evidence_class"`
	FreshnessState                  string   `json:"freshness_state"`
	FalsePositiveRiskNote           string   `json:"false_positive_risk_note"`
	FalseNegativeRiskNote           string   `json:"false_negative_risk_note"`
	UncertaintyNote                 string   `json:"uncertainty_note"`
	NextStep                        string   `json:"next_step"`
	ReviewerRequired                bool     `json:"reviewer_required"`
	VisibilityScope                 string   `json:"visibility_scope"`
	RedactionTier                   string   `json:"redaction_tier"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
	EvidenceRefs                    []string `json:"evidence_refs,omitempty"`
	ReviewRequiredPresentedApproved bool     `json:"review_required_presented_as_approved"`
	LeaksInternalEvidence           bool     `json:"leaks_internal_evidence"`
}

type BehavioralCalibrationSafetyGuardrailContract struct {
	CurrentState                     string `json:"current_state"`
	GuardrailID                      string `json:"guardrail_id"`
	AutoSuppressionBlocked           bool   `json:"auto_suppression_blocked"`
	CriticalControlRelaxationBlocked bool   `json:"critical_control_relaxation_blocked"`
	AutoBaselinePromotionBlocked     bool   `json:"auto_baseline_promotion_blocked"`
	PriorityMutationBlocked          bool   `json:"priority_mutation_blocked"`
	EnforcementMutationBlocked       bool   `json:"enforcement_mutation_blocked"`
	RequiredReviewForCriticalChanges bool   `json:"required_review_for_critical_changes"`
	RollbackRequiredForAdoption      bool   `json:"rollback_required_for_adoption"`
	LimitationMessage                string `json:"limitation_message"`
}

func intelligenceCalibrationValBObservedSignalClasses() []string {
	return []string{
		"process_behavior",
		"network_behavior",
		"file_behavior",
		"runtime_behavior",
		"workload_behavior",
		"identity_or_actor_behavior",
		IntelligenceCalibrationValBUnsupportedSignal,
	}
}

func intelligenceCalibrationValBThresholdDirections() []string {
	return []string{
		IntelligenceCalibrationValBThresholdIncreaseSensitivity,
		IntelligenceCalibrationValBThresholdDecreaseSensitivity,
		IntelligenceCalibrationValBThresholdKeep,
		IntelligenceCalibrationValBThresholdRequiresReview,
	}
}

func intelligenceCalibrationValBDriftAdjustments() []string {
	return []string{
		IntelligenceCalibrationValBDriftAdjustmentIncrease,
		IntelligenceCalibrationValBDriftAdjustmentDecrease,
		IntelligenceCalibrationValBDriftAdjustmentKeep,
		IntelligenceCalibrationValBDriftAdjustmentRequiresReview,
	}
}

func intelligenceCalibrationValBCriticalityClasses() []string {
	return []string{
		IntelligenceCalibrationValBCriticalityCritical,
		IntelligenceCalibrationValBCriticalityHigh,
		IntelligenceCalibrationValBCriticalityNormal,
		IntelligenceCalibrationValBCriticalityLow,
		IntelligenceCalibrationValBCriticalityUnknown,
	}
}

func intelligenceCalibrationValBWeightingActions() []string {
	return []string{
		IntelligenceCalibrationValBWeightingRaisePriority,
		IntelligenceCalibrationValBWeightingLowerPriorityCandidate,
		IntelligenceCalibrationValBWeightingKeep,
		IntelligenceCalibrationValBWeightingRequiresReview,
	}
}

func intelligenceCalibrationValBAdoptionStates() []string {
	return []string{
		IntelligenceCalibrationApprovalProposed,
		IntelligenceCalibrationApprovalReviewRequired,
		IntelligenceCalibrationApprovalApproved,
		IntelligenceCalibrationApprovalRejected,
		IntelligenceCalibrationApprovalSuperseded,
	}
}

func parseIntelligenceCalibrationValBTimestamp(value string) (time.Time, bool) {
	parsed, err := time.Parse(time.RFC3339, strings.TrimSpace(value))
	if err != nil {
		return time.Time{}, false
	}
	return parsed, true
}

func IntelligenceCalibrationValBBehavioralBaselineContract() BehavioralBaselineProfileContract {
	return BehavioralBaselineProfileContract{
		CurrentState:                 IntelligenceCalibrationValBBehavioralBaselineStateActive,
		SupportedSignalClasses:       intelligenceCalibrationValBObservedSignalClasses(),
		BaselineID:                   "baseline-profile-001",
		BaselineVersion:              "2026.04.25",
		BaselineScope:                "service_behavior_scope",
		WorkloadContextRefs:          []string{"workload/prod/web"},
		AssetOrServiceRefs:           []string{"service/web-api"},
		ObservedSignalClasses:        intelligenceCalibrationValBObservedSignalClasses(),
		ObservedProcessClasses:       []string{"web_process", "worker_process"},
		ObservedNetworkPaths:         []string{"svc:web-api->postgres", "svc:web-api->redis"},
		ObservedFileActivityClasses:  []string{"config_read", "cache_write"},
		ObservedRuntimeSignalClasses: []string{"container_start", "module_load"},
		ObservationWindowStart:       "2026-04-18T00:00:00Z",
		ObservationWindowEnd:         "2026-04-25T00:00:00Z",
		SampleCount:                  240,
		OutlierHandlingPolicy:        "bounded_outlier_review_required_before_adoption",
		FreshnessState:               IntelligenceCalibrationFreshnessFresh,
		EvidenceRefs:                 []string{"behavioral_baseline_profile", "evidence_spine"},
		LimitationMessage:            "behavioral_baseline_profile_is_projection_only_and_requires_review_before_any_adoption",
		ProjectionDisclaimer:         "projection_only not_canonical_truth advisory_behavioral_baseline_learning_mode",
		CanRelaxEnforcement:          false,
		CanSuppressAlerts:            false,
		AdvisoryOnly:                 true,
	}
}

func IntelligenceCalibrationValBLearningModeRuntimeContract() LearningModeRuntimeDisciplineContract {
	return LearningModeRuntimeDisciplineContract{
		CurrentState:              IntelligenceCalibrationValBLearningRuntimeStateActive,
		SupportedStates:           intelligenceCalibrationVal0LearningStates(),
		SupportedSignalClasses:    intelligenceCalibrationValBObservedSignalClasses(),
		LearningSessionID:         "learning-session-001",
		LearningModeState:         IntelligenceCalibrationLearningEnabled,
		Scope:                     "service/web-api:behavioral_learning",
		StartedAt:                 "2026-04-25T08:00:00Z",
		ExpiresAt:                 "2026-04-25T12:00:00Z",
		BoundedDurationConfirmed:  true,
		ExcludedCriticalControls:  []string{"critical_authz_control", "critical_egress_block"},
		ObservedSignalClasses:     []string{"process_behavior", "network_behavior", "runtime_behavior"},
		DisallowedSignalClasses:   []string{"identity_or_actor_behavior", IntelligenceCalibrationValBUnsupportedSignal},
		OutputReviewRequired:      true,
		CanRelaxEnforcement:       false,
		CanAutoPromoteBaseline:    false,
		CanSuppressCriticalAlerts: false,
		ReviewSummaryRequired:     true,
		LimitationMessage:         "learning_mode_runtime_is_bounded_observation_only_and_needs_review_before_any_adoption",
	}
}

func IntelligenceCalibrationValBAnomalyThresholdContract() AnomalyThresholdCalibrationContract {
	return AnomalyThresholdCalibrationContract{
		CurrentState:             IntelligenceCalibrationValBThresholdStateActive,
		SupportedDirections:      intelligenceCalibrationValBThresholdDirections(),
		ThresholdProfileID:       "threshold-profile-001",
		SignalClass:              "runtime_behavior",
		BaselineRef:              "baseline-profile-001",
		CurrentThreshold:         0.80,
		ProposedThreshold:        0.80,
		ThresholdChangeDirection: IntelligenceCalibrationValBThresholdKeep,
		ConfidenceBand:           IntelligenceCalibrationConfidenceMedium,
		EvidenceClass:            IntelligenceCalibrationEvidenceDirectlyEvidenced,
		FalsePositiveRiskNote:    "lowering_threshold_without review can amplify noisy behavior",
		FalseNegativeRiskNote:    "raising_threshold_without evidence can hide low-signal malicious activity",
		ReviewRequired:           false,
		AppliesToCriticalClass:   false,
		LimitationMessage:        "threshold_calibration_is_advisory_only_and_does_not_mutate_active_detection_state",
		AdvisoryOnly:             true,
		MutatesActiveDetection:   false,
	}
}

func IntelligenceCalibrationValBDriftSensitivityContract() DriftSensitivityScalingContract {
	return DriftSensitivityScalingContract{
		CurrentState:           IntelligenceCalibrationValBDriftStateActive,
		SupportedDriftBands:    intelligenceCalibrationVal0ConfidenceBands(),
		SupportedAdjustments:   intelligenceCalibrationValBDriftAdjustments(),
		SupportedCriticalities: intelligenceCalibrationValBCriticalityClasses(),
		DriftProfileID:         "drift-profile-001",
		BaselineRef:            "baseline-profile-001",
		SignalClass:            "workload_behavior",
		DriftScoreBand:         IntelligenceCalibrationConfidenceMedium,
		SensitivityAdjustment:  IntelligenceCalibrationValBDriftAdjustmentKeep,
		CriticalityContext:     IntelligenceCalibrationValBCriticalityNormal,
		EvidenceRefs:           []string{"drift_signal_observation"},
		FreshnessState:         IntelligenceCalibrationFreshnessFresh,
		ReviewRequired:         false,
		LimitationMessage:      "drift_sensitivity_scaling_is_advisory_and_requires_review_before_any_enforcement_change",
		AdvisoryOnly:           true,
		MutatesEnforcement:     false,
	}
}

func IntelligenceCalibrationValBCriticalityWeightingContract() CriticalityAwareWeightingContract {
	return CriticalityAwareWeightingContract{
		CurrentState:             IntelligenceCalibrationValBWeightingStateActive,
		SupportedCriticalities:   intelligenceCalibrationValBCriticalityClasses(),
		SupportedWeighting:       intelligenceCalibrationValBWeightingActions(),
		WeightingProfileID:       "weighting-profile-001",
		AssetOrServiceRef:        "service/web-api",
		CriticalityClass:         IntelligenceCalibrationValBCriticalityNormal,
		WeightingAction:          IntelligenceCalibrationValBWeightingKeep,
		ConfidenceBand:           IntelligenceCalibrationConfidenceMedium,
		EvidenceClass:            IntelligenceCalibrationEvidenceDirectlyEvidenced,
		LocalContextRefs:         []string{"workload/prod/web"},
		BlastRadiusHint:          "service_scope",
		ReviewerRequired:         false,
		LimitationMessage:        "criticality_weighting_is_advisory_and_cannot_mutate_canonical_priority",
		ReasonCode:               "local_behavioral_context",
		EvidenceRefs:             []string{"criticality_context_observation"},
		AdvisoryOnly:             true,
		MutatesCanonicalPriority: false,
	}
}

func IntelligenceCalibrationValBBaselineFreshnessContract() BaselineFreshnessExpiryContract {
	return BaselineFreshnessExpiryContract{
		CurrentState:      IntelligenceCalibrationValBBaselineFreshnessStateActive,
		BaselineRef:       "baseline-profile-001",
		FreshnessState:    IntelligenceCalibrationFreshnessFresh,
		LastObservedAt:    "2026-04-25T00:00:00Z",
		ExpiresAt:         "2026-05-02T00:00:00Z",
		FreshnessLimit:    "168h",
		LimitationMessage: "freshness_metadata_is_projection_only_and_guides_review_before_any_behavioral_adoption",
		AdvisoryOnly:      true,
	}
}

func IntelligenceCalibrationValBBaselineAdoptionContract() BaselineAdoptionReviewContract {
	return BaselineAdoptionReviewContract{
		CurrentState:            IntelligenceCalibrationValBBaselineAdoptionStateActive,
		SupportedAdoptionStates: intelligenceCalibrationValBAdoptionStates(),
		AdoptionID:              "adoption-001",
		BaselineRef:             "baseline-profile-001",
		ProposedAdoptionState:   IntelligenceCalibrationApprovalApproved,
		ReviewerRequired:        true,
		ApprovalRef:             "approval/baseline-001",
		OutputSummaryRef:        "summary/baseline-adoption-001",
		RiskNote:                "approved_review_still_requires_later_governed_mutation_outside_valb",
		RollbackRef:             "rollback/baseline-001",
		BeforeMetricSnapshotRef: "metrics/before-baseline-adoption-001",
		AfterMetricSnapshotRef:  "metrics/after-baseline-adoption-001",
		LimitationMessage:       "baseline_adoption_review_is_bounded_and_does_not_mutate_active_baseline_in_valb",
		MutatesActiveBaseline:   false,
		GovernanceRequired:      true,
		AdvisoryOnly:            true,
	}
}

func IntelligenceCalibrationValBExplanationContract() BehavioralCalibrationExplanationContract {
	return BehavioralCalibrationExplanationContract{
		CurrentState:                    IntelligenceCalibrationValBExplanationStateActive,
		ReasonCode:                      "behavioral_threshold_review_required",
		HumanMessage:                    "Behavioral threshold recommendation is bounded and requires review before any active change.",
		TechnicalDetail:                 "Observed runtime_behavior remained within baseline variance, but FP/FN tradeoffs remain explicit and bounded.",
		BaselineRef:                     "baseline-profile-001",
		SignalClass:                     "runtime_behavior",
		ConfidenceBand:                  IntelligenceCalibrationConfidenceMedium,
		EvidenceClass:                   IntelligenceCalibrationEvidenceDirectlyEvidenced,
		FreshnessState:                  IntelligenceCalibrationFreshnessFresh,
		FalsePositiveRiskNote:           "too much sensitivity can over-alert on benign periodic runtime behavior",
		FalseNegativeRiskNote:           "too little sensitivity can miss low-signal adversarial runtime drift",
		UncertaintyNote:                 "behavioral variance remains workload-sensitive and is not final proof",
		NextStep:                        "review baseline and threshold recommendation before any governed adoption",
		ReviewerRequired:                true,
		VisibilityScope:                 ProductionUsabilityVisibilityOperator,
		RedactionTier:                   ProductionUsabilityRedactionLow,
		ProjectionDisclaimer:            "projection_only not_canonical_truth advisory_behavioral_baseline_learning_mode",
		EvidenceRefs:                    []string{"behavioral_explanation_payload"},
		ReviewRequiredPresentedApproved: false,
		LeaksInternalEvidence:           false,
	}
}

func IntelligenceCalibrationValBGuardrailContract() BehavioralCalibrationSafetyGuardrailContract {
	return BehavioralCalibrationSafetyGuardrailContract{
		CurrentState:                     IntelligenceCalibrationValBGuardrailStateActive,
		GuardrailID:                      "behavioral-guardrail-001",
		AutoSuppressionBlocked:           true,
		CriticalControlRelaxationBlocked: true,
		AutoBaselinePromotionBlocked:     true,
		PriorityMutationBlocked:          true,
		EnforcementMutationBlocked:       true,
		RequiredReviewForCriticalChanges: true,
		RollbackRequiredForAdoption:      true,
		LimitationMessage:                "behavioral_calibration_guardrails_block_auto_suppression_promotion_and_mutation_in_valb",
	}
}

func EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model BehavioralBaselineProfileContract) string {
	if strings.TrimSpace(model.BaselineID) == "" || strings.TrimSpace(model.BaselineVersion) == "" || strings.TrimSpace(model.BaselineScope) == "" || len(model.WorkloadContextRefs) == 0 || len(model.AssetOrServiceRefs) == 0 || len(model.ObservedProcessClasses) == 0 || len(model.ObservedNetworkPaths) == 0 || len(model.ObservedFileActivityClasses) == 0 || len(model.ObservedRuntimeSignalClasses) == 0 || strings.TrimSpace(model.ObservationWindowStart) == "" || strings.TrimSpace(model.ObservationWindowEnd) == "" || model.SampleCount <= 0 || strings.TrimSpace(model.OutlierHandlingPolicy) == "" || strings.TrimSpace(model.FreshnessState) == "" || len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.LimitationMessage) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValBBehavioralBaselineStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedSignalClasses, intelligenceCalibrationValBObservedSignalClasses()...) || !containsExactTrimmedStringSet(model.ObservedSignalClasses, intelligenceCalibrationValBObservedSignalClasses()...) || !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") || !model.AdvisoryOnly {
		return IntelligenceCalibrationValBBehavioralBaselineStatePartial
	}
	windowStart, ok := parseIntelligenceCalibrationValBTimestamp(model.ObservationWindowStart)
	if !ok {
		return IntelligenceCalibrationValBBehavioralBaselineStatePartial
	}
	windowEnd, ok := parseIntelligenceCalibrationValBTimestamp(model.ObservationWindowEnd)
	if !ok || !windowEnd.After(windowStart) {
		return IntelligenceCalibrationValBBehavioralBaselineStatePartial
	}
	if model.CanRelaxEnforcement || model.CanSuppressAlerts {
		return IntelligenceCalibrationValBBehavioralBaselineStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessExpired {
		return IntelligenceCalibrationValBBehavioralBaselineStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessStale && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "stale") {
		return IntelligenceCalibrationValBBehavioralBaselineStatePartial
	}
	return IntelligenceCalibrationValBBehavioralBaselineStateActive
}

func EvaluateIntelligenceCalibrationValBLearningRuntimeState(model LearningModeRuntimeDisciplineContract) string {
	if strings.TrimSpace(model.LearningSessionID) == "" || strings.TrimSpace(model.LearningModeState) == "" || strings.TrimSpace(model.Scope) == "" || strings.TrimSpace(model.StartedAt) == "" || strings.TrimSpace(model.ExpiresAt) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValBLearningRuntimeStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedStates, intelligenceCalibrationVal0LearningStates()...) || !containsTrimmedString(model.SupportedStates, model.LearningModeState) || len(model.ExcludedCriticalControls) == 0 || len(model.ObservedSignalClasses) == 0 || len(model.DisallowedSignalClasses) == 0 || !containsAllTrimmedStrings(model.SupportedSignalClasses, intelligenceCalibrationValBObservedSignalClasses()...) {
		return IntelligenceCalibrationValBLearningRuntimeStatePartial
	}
	for _, signalClass := range append(append([]string{}, model.ObservedSignalClasses...), model.DisallowedSignalClasses...) {
		if !containsTrimmedString(model.SupportedSignalClasses, signalClass) {
			return IntelligenceCalibrationValBLearningRuntimeStatePartial
		}
	}
	startedAt, ok := parseIntelligenceCalibrationValBTimestamp(model.StartedAt)
	if !ok {
		return IntelligenceCalibrationValBLearningRuntimeStatePartial
	}
	expiresAt, ok := parseIntelligenceCalibrationValBTimestamp(model.ExpiresAt)
	if !ok || !expiresAt.After(startedAt) {
		return IntelligenceCalibrationValBLearningRuntimeStatePartial
	}
	if !model.BoundedDurationConfirmed || !model.OutputReviewRequired || !model.ReviewSummaryRequired || model.CanRelaxEnforcement || model.CanAutoPromoteBaseline || model.CanSuppressCriticalAlerts || model.LearningModeState == IntelligenceCalibrationLearningExpired {
		return IntelligenceCalibrationValBLearningRuntimeStatePartial
	}
	return IntelligenceCalibrationValBLearningRuntimeStateActive
}

func EvaluateIntelligenceCalibrationValBThresholdState(model AnomalyThresholdCalibrationContract) string {
	if strings.TrimSpace(model.ThresholdProfileID) == "" || strings.TrimSpace(model.SignalClass) == "" || strings.TrimSpace(model.BaselineRef) == "" || model.CurrentThreshold <= 0 || model.ProposedThreshold <= 0 || strings.TrimSpace(model.ThresholdChangeDirection) == "" || strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.EvidenceClass) == "" || strings.TrimSpace(model.FalsePositiveRiskNote) == "" || strings.TrimSpace(model.FalseNegativeRiskNote) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValBThresholdStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedDirections, intelligenceCalibrationValBThresholdDirections()...) || !containsTrimmedString(model.SupportedDirections, model.ThresholdChangeDirection) || !containsTrimmedString(intelligenceCalibrationValBObservedSignalClasses(), model.SignalClass) || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceBand) || !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), model.EvidenceClass) || !model.AdvisoryOnly || model.MutatesActiveDetection {
		return IntelligenceCalibrationValBThresholdStatePartial
	}
	if model.AppliesToCriticalClass && model.ThresholdChangeDirection == IntelligenceCalibrationValBThresholdDecreaseSensitivity && !model.ReviewRequired {
		return IntelligenceCalibrationValBThresholdStatePartial
	}
	if model.EvidenceClass == IntelligenceCalibrationEvidenceUnsupported && model.ThresholdChangeDirection != IntelligenceCalibrationValBThresholdRequiresReview {
		return IntelligenceCalibrationValBThresholdStatePartial
	}
	return IntelligenceCalibrationValBThresholdStateActive
}

func EvaluateIntelligenceCalibrationValBDriftState(model DriftSensitivityScalingContract) string {
	if strings.TrimSpace(model.DriftProfileID) == "" || strings.TrimSpace(model.BaselineRef) == "" || strings.TrimSpace(model.SignalClass) == "" || strings.TrimSpace(model.DriftScoreBand) == "" || strings.TrimSpace(model.SensitivityAdjustment) == "" || strings.TrimSpace(model.CriticalityContext) == "" || len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValBDriftStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedDriftBands, intelligenceCalibrationVal0ConfidenceBands()...) || !containsExactTrimmedStringSet(model.SupportedAdjustments, intelligenceCalibrationValBDriftAdjustments()...) || !containsExactTrimmedStringSet(model.SupportedCriticalities, intelligenceCalibrationValBCriticalityClasses()...) || !containsTrimmedString(model.SupportedDriftBands, model.DriftScoreBand) || !containsTrimmedString(model.SupportedAdjustments, model.SensitivityAdjustment) || !containsTrimmedString(model.SupportedCriticalities, model.CriticalityContext) || !containsTrimmedString(intelligenceCalibrationValBObservedSignalClasses(), model.SignalClass) || !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || !model.AdvisoryOnly || model.MutatesEnforcement {
		return IntelligenceCalibrationValBDriftStatePartial
	}
	if model.DriftScoreBand == IntelligenceCalibrationConfidenceUnknown && model.SensitivityAdjustment != IntelligenceCalibrationValBDriftAdjustmentRequiresReview {
		return IntelligenceCalibrationValBDriftStatePartial
	}
	if (model.CriticalityContext == IntelligenceCalibrationValBCriticalityCritical || model.CriticalityContext == IntelligenceCalibrationValBCriticalityHigh) && model.SensitivityAdjustment == IntelligenceCalibrationValBDriftAdjustmentDecrease && !model.ReviewRequired {
		return IntelligenceCalibrationValBDriftStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessExpired {
		return IntelligenceCalibrationValBDriftStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessStale && !model.ReviewRequired && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "stale") {
		return IntelligenceCalibrationValBDriftStatePartial
	}
	return IntelligenceCalibrationValBDriftStateActive
}

func EvaluateIntelligenceCalibrationValBWeightingState(model CriticalityAwareWeightingContract) string {
	if strings.TrimSpace(model.WeightingProfileID) == "" || strings.TrimSpace(model.AssetOrServiceRef) == "" || strings.TrimSpace(model.CriticalityClass) == "" || strings.TrimSpace(model.WeightingAction) == "" || strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.EvidenceClass) == "" || len(model.LocalContextRefs) == 0 || strings.TrimSpace(model.BlastRadiusHint) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValBWeightingStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedCriticalities, intelligenceCalibrationValBCriticalityClasses()...) || !containsExactTrimmedStringSet(model.SupportedWeighting, intelligenceCalibrationValBWeightingActions()...) || !containsTrimmedString(model.SupportedCriticalities, model.CriticalityClass) || !containsTrimmedString(model.SupportedWeighting, model.WeightingAction) || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceBand) || !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), model.EvidenceClass) || !model.AdvisoryOnly || model.MutatesCanonicalPriority {
		return IntelligenceCalibrationValBWeightingStatePartial
	}
	if model.CriticalityClass == IntelligenceCalibrationValBCriticalityUnknown && model.WeightingAction == IntelligenceCalibrationValBWeightingLowerPriorityCandidate {
		return IntelligenceCalibrationValBWeightingStatePartial
	}
	if (model.CriticalityClass == IntelligenceCalibrationValBCriticalityCritical || model.CriticalityClass == IntelligenceCalibrationValBCriticalityHigh) && model.WeightingAction == IntelligenceCalibrationValBWeightingLowerPriorityCandidate && !model.ReviewerRequired {
		return IntelligenceCalibrationValBWeightingStatePartial
	}
	if model.WeightingAction == IntelligenceCalibrationValBWeightingRaisePriority && (strings.TrimSpace(model.ReasonCode) == "" || len(model.EvidenceRefs) == 0) {
		return IntelligenceCalibrationValBWeightingStatePartial
	}
	return IntelligenceCalibrationValBWeightingStateActive
}

func EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model BaselineFreshnessExpiryContract) string {
	if strings.TrimSpace(model.BaselineRef) == "" || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.FreshnessLimit) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValBBaselineFreshnessStateIncomplete
	}
	if !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || !model.AdvisoryOnly {
		return IntelligenceCalibrationValBBaselineFreshnessStatePartial
	}
	if strings.TrimSpace(model.LastObservedAt) == "" || strings.TrimSpace(model.ExpiresAt) == "" {
		return IntelligenceCalibrationValBBaselineFreshnessStatePartial
	}
	lastObservedAt, ok := parseIntelligenceCalibrationValBTimestamp(model.LastObservedAt)
	if !ok {
		return IntelligenceCalibrationValBBaselineFreshnessStatePartial
	}
	expiresAt, ok := parseIntelligenceCalibrationValBTimestamp(model.ExpiresAt)
	if !ok {
		return IntelligenceCalibrationValBBaselineFreshnessStatePartial
	}
	if !expiresAt.After(lastObservedAt) {
		return IntelligenceCalibrationValBBaselineFreshnessStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessExpired {
		return IntelligenceCalibrationValBBaselineFreshnessStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessUnknown && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "unknown") {
		return IntelligenceCalibrationValBBaselineFreshnessStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessUnsupported && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "unsupported") {
		return IntelligenceCalibrationValBBaselineFreshnessStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessStale && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "stale") {
		return IntelligenceCalibrationValBBaselineFreshnessStatePartial
	}
	return IntelligenceCalibrationValBBaselineFreshnessStateActive
}

func EvaluateIntelligenceCalibrationValBBaselineAdoptionState(model BaselineAdoptionReviewContract) string {
	if strings.TrimSpace(model.AdoptionID) == "" || strings.TrimSpace(model.BaselineRef) == "" || strings.TrimSpace(model.ProposedAdoptionState) == "" || strings.TrimSpace(model.OutputSummaryRef) == "" || strings.TrimSpace(model.RiskNote) == "" || strings.TrimSpace(model.BeforeMetricSnapshotRef) == "" || strings.TrimSpace(model.AfterMetricSnapshotRef) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValBBaselineAdoptionStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedAdoptionStates, intelligenceCalibrationValBAdoptionStates()...) || !containsTrimmedString(model.SupportedAdoptionStates, model.ProposedAdoptionState) || !model.AdvisoryOnly {
		return IntelligenceCalibrationValBBaselineAdoptionStatePartial
	}
	if model.ProposedAdoptionState == IntelligenceCalibrationApprovalProposed || model.ProposedAdoptionState == IntelligenceCalibrationApprovalReviewRequired {
		return IntelligenceCalibrationValBBaselineAdoptionStatePartial
	}
	if model.ProposedAdoptionState == IntelligenceCalibrationApprovalApproved && (strings.TrimSpace(model.RollbackRef) == "" || strings.TrimSpace(model.ApprovalRef) == "" || !model.GovernanceRequired || !model.ReviewerRequired) {
		return IntelligenceCalibrationValBBaselineAdoptionStatePartial
	}
	if model.MutatesActiveBaseline {
		return IntelligenceCalibrationValBBaselineAdoptionStatePartial
	}
	return IntelligenceCalibrationValBBaselineAdoptionStateActive
}

func EvaluateIntelligenceCalibrationValBExplanationState(model BehavioralCalibrationExplanationContract) string {
	if strings.TrimSpace(model.ReasonCode) == "" || strings.TrimSpace(model.HumanMessage) == "" || strings.TrimSpace(model.TechnicalDetail) == "" || strings.TrimSpace(model.BaselineRef) == "" || strings.TrimSpace(model.SignalClass) == "" || strings.TrimSpace(model.ConfidenceBand) == "" || strings.TrimSpace(model.EvidenceClass) == "" || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.FalsePositiveRiskNote) == "" || strings.TrimSpace(model.FalseNegativeRiskNote) == "" || strings.TrimSpace(model.UncertaintyNote) == "" || strings.TrimSpace(model.NextStep) == "" || strings.TrimSpace(model.VisibilityScope) == "" || strings.TrimSpace(model.RedactionTier) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" || len(model.EvidenceRefs) == 0 {
		return IntelligenceCalibrationValBExplanationStateIncomplete
	}
	if !containsTrimmedString(intelligenceCalibrationValBObservedSignalClasses(), model.SignalClass) || !containsTrimmedString(intelligenceCalibrationVal0ConfidenceBands(), model.ConfidenceBand) || !containsTrimmedString(intelligenceCalibrationVal0EvidenceClasses(), model.EvidenceClass) || !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || !containsTrimmedString(productionUsabilityValAExplainScopes(), model.VisibilityScope) || !containsTrimmedString(ProductionUsabilityVal0ExplainabilityContract().SupportedRedactionTiers, model.RedactionTier) || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValBExplanationStatePartial
	}
	if model.ReviewRequiredPresentedApproved || model.LeaksInternalEvidence {
		return IntelligenceCalibrationValBExplanationStatePartial
	}
	if model.VisibilityScope == ProductionUsabilityVisibilityPublicSafe && model.RedactionTier != ProductionUsabilityRedactionPublicSafe {
		return IntelligenceCalibrationValBExplanationStatePartial
	}
	if model.VisibilityScope == ProductionUsabilityVisibilityPartner && (model.RedactionTier == ProductionUsabilityRedactionNone || model.RedactionTier == ProductionUsabilityRedactionLow) {
		return IntelligenceCalibrationValBExplanationStatePartial
	}
	return IntelligenceCalibrationValBExplanationStateActive
}

func EvaluateIntelligenceCalibrationValBGuardrailState(model BehavioralCalibrationSafetyGuardrailContract) string {
	if strings.TrimSpace(model.GuardrailID) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValBGuardrailStateIncomplete
	}
	if !model.AutoSuppressionBlocked || !model.CriticalControlRelaxationBlocked || !model.AutoBaselinePromotionBlocked || !model.PriorityMutationBlocked || !model.EnforcementMutationBlocked || !model.RequiredReviewForCriticalChanges || !model.RollbackRequiredForAdoption {
		return IntelligenceCalibrationValBGuardrailStatePartial
	}
	return IntelligenceCalibrationValBGuardrailStateActive
}

func EvaluateIntelligenceCalibrationValBState(val0DependencyState, val0FoundationState, valADependencyState, valAState, baselineState, learningRuntimeState, thresholdState, driftState, weightingState, baselineFreshnessState, baselineAdoptionState, explanationState, guardrailState string) string {
	if strings.TrimSpace(val0DependencyState) != IntelligenceCalibrationVal0StateActive || strings.TrimSpace(val0FoundationState) != IntelligenceCalibrationVal0StateActive || strings.TrimSpace(valADependencyState) != IntelligenceCalibrationValAStateActive || strings.TrimSpace(valAState) != IntelligenceCalibrationValAStateActive {
		return IntelligenceCalibrationValBStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(baselineState),
		strings.TrimSpace(learningRuntimeState),
		strings.TrimSpace(thresholdState),
		strings.TrimSpace(driftState),
		strings.TrimSpace(weightingState),
		strings.TrimSpace(baselineFreshnessState),
		strings.TrimSpace(baselineAdoptionState),
		strings.TrimSpace(explanationState),
		strings.TrimSpace(guardrailState),
	} {
		switch state {
		case IntelligenceCalibrationValBBehavioralBaselineStateActive,
			IntelligenceCalibrationValBLearningRuntimeStateActive,
			IntelligenceCalibrationValBThresholdStateActive,
			IntelligenceCalibrationValBDriftStateActive,
			IntelligenceCalibrationValBWeightingStateActive,
			IntelligenceCalibrationValBBaselineFreshnessStateActive,
			IntelligenceCalibrationValBBaselineAdoptionStateActive,
			IntelligenceCalibrationValBExplanationStateActive,
			IntelligenceCalibrationValBGuardrailStateActive:
		case IntelligenceCalibrationValBBehavioralBaselineStatePartial,
			IntelligenceCalibrationValBLearningRuntimeStatePartial,
			IntelligenceCalibrationValBThresholdStatePartial,
			IntelligenceCalibrationValBDriftStatePartial,
			IntelligenceCalibrationValBWeightingStatePartial,
			IntelligenceCalibrationValBBaselineFreshnessStatePartial,
			IntelligenceCalibrationValBBaselineAdoptionStatePartial,
			IntelligenceCalibrationValBExplanationStatePartial,
			IntelligenceCalibrationValBGuardrailStatePartial:
			hasPartial = true
		default:
			return IntelligenceCalibrationValBStateIncomplete
		}
	}
	if hasPartial {
		return IntelligenceCalibrationValBStateSubstantial
	}
	return IntelligenceCalibrationValBStateActive
}

func EvaluateIntelligenceCalibrationValBProofsState(val0DependencyState, val0FoundationState, valADependencyState, valAState, baselineState, learningRuntimeState, thresholdState, driftState, weightingState, baselineFreshnessState, baselineAdoptionState, explanationState, guardrailState string, surfaceRefs, evidenceRefs, limitations, whyPoint5NotPass []string, projectionDisclaimer string) string {
	baseState := EvaluateIntelligenceCalibrationValBState(
		val0DependencyState,
		val0FoundationState,
		valADependencyState,
		valAState,
		baselineState,
		learningRuntimeState,
		thresholdState,
		driftState,
		weightingState,
		baselineFreshnessState,
		baselineAdoptionState,
		explanationState,
		guardrailState,
	)
	if len(surfaceRefs) < 10 || len(evidenceRefs) < 9 || len(limitations) == 0 || len(whyPoint5NotPass) == 0 || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "not_canonical_truth") {
		if baseState == IntelligenceCalibrationValBStateActive {
			return IntelligenceCalibrationValBStateSubstantial
		}
		return baseState
	}
	return baseState
}

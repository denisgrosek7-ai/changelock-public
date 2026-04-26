package operability

import "strings"

const (
	IntelligenceCalibrationValDSimulationHarnessStateActive     = "intelligence_calibration_vald_simulation_harness_active"
	IntelligenceCalibrationValDSimulationHarnessStatePartial    = "intelligence_calibration_vald_simulation_harness_partial"
	IntelligenceCalibrationValDSimulationHarnessStateIncomplete = "intelligence_calibration_vald_simulation_harness_incomplete"

	IntelligenceCalibrationValDScenarioLibraryStateActive     = "intelligence_calibration_vald_scenario_library_active"
	IntelligenceCalibrationValDScenarioLibraryStatePartial    = "intelligence_calibration_vald_scenario_library_partial"
	IntelligenceCalibrationValDScenarioLibraryStateIncomplete = "intelligence_calibration_vald_scenario_library_incomplete"

	IntelligenceCalibrationValDMissedDetectionStateActive     = "intelligence_calibration_vald_missed_detection_active"
	IntelligenceCalibrationValDMissedDetectionStatePartial    = "intelligence_calibration_vald_missed_detection_partial"
	IntelligenceCalibrationValDMissedDetectionStateIncomplete = "intelligence_calibration_vald_missed_detection_incomplete"

	IntelligenceCalibrationValDFPFNBalanceStateActive     = "intelligence_calibration_vald_fp_fn_balance_active"
	IntelligenceCalibrationValDFPFNBalanceStatePartial    = "intelligence_calibration_vald_fp_fn_balance_partial"
	IntelligenceCalibrationValDFPFNBalanceStateIncomplete = "intelligence_calibration_vald_fp_fn_balance_incomplete"

	IntelligenceCalibrationValDConfidenceReviewStateActive     = "intelligence_calibration_vald_confidence_review_active"
	IntelligenceCalibrationValDConfidenceReviewStatePartial    = "intelligence_calibration_vald_confidence_review_partial"
	IntelligenceCalibrationValDConfidenceReviewStateIncomplete = "intelligence_calibration_vald_confidence_review_incomplete"

	IntelligenceCalibrationValDVEXQualityStateActive     = "intelligence_calibration_vald_vex_quality_active"
	IntelligenceCalibrationValDVEXQualityStatePartial    = "intelligence_calibration_vald_vex_quality_partial"
	IntelligenceCalibrationValDVEXQualityStateIncomplete = "intelligence_calibration_vald_vex_quality_incomplete"

	IntelligenceCalibrationValDReachabilityQualityStateActive     = "intelligence_calibration_vald_reachability_quality_active"
	IntelligenceCalibrationValDReachabilityQualityStatePartial    = "intelligence_calibration_vald_reachability_quality_partial"
	IntelligenceCalibrationValDReachabilityQualityStateIncomplete = "intelligence_calibration_vald_reachability_quality_incomplete"

	IntelligenceCalibrationValDBehavioralQualityStateActive     = "intelligence_calibration_vald_behavioral_quality_active"
	IntelligenceCalibrationValDBehavioralQualityStatePartial    = "intelligence_calibration_vald_behavioral_quality_partial"
	IntelligenceCalibrationValDBehavioralQualityStateIncomplete = "intelligence_calibration_vald_behavioral_quality_incomplete"

	IntelligenceCalibrationValDFederatedQualityStateActive     = "intelligence_calibration_vald_federated_quality_active"
	IntelligenceCalibrationValDFederatedQualityStatePartial    = "intelligence_calibration_vald_federated_quality_partial"
	IntelligenceCalibrationValDFederatedQualityStateIncomplete = "intelligence_calibration_vald_federated_quality_incomplete"

	IntelligenceCalibrationValDSimulationCoverageStateActive     = "intelligence_calibration_vald_simulation_coverage_active"
	IntelligenceCalibrationValDSimulationCoverageStatePartial    = "intelligence_calibration_vald_simulation_coverage_partial"
	IntelligenceCalibrationValDSimulationCoverageStateIncomplete = "intelligence_calibration_vald_simulation_coverage_incomplete"

	IntelligenceCalibrationValDQualityScoreboardStateActive     = "intelligence_calibration_vald_quality_scoreboard_active"
	IntelligenceCalibrationValDQualityScoreboardStatePartial    = "intelligence_calibration_vald_quality_scoreboard_partial"
	IntelligenceCalibrationValDQualityScoreboardStateIncomplete = "intelligence_calibration_vald_quality_scoreboard_incomplete"

	IntelligenceCalibrationValDStateIncomplete  = "intelligence_calibration_vald_incomplete"
	IntelligenceCalibrationValDStateSubstantial = "intelligence_calibration_vald_substantially_ready"
	IntelligenceCalibrationValDStateActive      = "intelligence_calibration_vald_active"

	IntelligenceCalibrationValDCoveragePass        = "pass"
	IntelligenceCalibrationValDCoveragePartial     = "partial"
	IntelligenceCalibrationValDCoverageFail        = "fail"
	IntelligenceCalibrationValDCoverageUnsupported = "unsupported"
)

type DefensiveSimulationScenarioHarnessContract struct {
	CurrentState             string   `json:"current_state"`
	SupportedScenarioClasses []string `json:"supported_scenario_classes,omitempty"`
	HarnessID                string   `json:"harness_id"`
	HarnessVersion           string   `json:"harness_version"`
	ScenarioClasses          []string `json:"scenario_classes,omitempty"`
	ScenarioRefs             []string `json:"scenario_refs,omitempty"`
	ScenarioCount            int      `json:"scenario_count"`
	Replayable               bool     `json:"replayable"`
	ExpectedOutcomesPresent  bool     `json:"expected_outcomes_present"`
	ActualOutcomesPresent    bool     `json:"actual_outcomes_present"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	GeneratedAt              string   `json:"generated_at"`
	FreshnessState           string   `json:"freshness_state"`
	LimitationMessage        string   `json:"limitation_message"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
}

type AdversarialLowSignalScenarioLibraryContract struct {
	CurrentState                   string   `json:"current_state"`
	LibraryID                      string   `json:"library_id"`
	LibraryVersion                 string   `json:"library_version"`
	Scenarios                      []string `json:"scenarios,omitempty"`
	LowSignalScenarios             []string `json:"low_signal_scenarios,omitempty"`
	AdversarialScenarios           []string `json:"adversarial_scenarios,omitempty"`
	CriticalAssetScenarios         []string `json:"critical_asset_scenarios,omitempty"`
	ExpectedDetectionOutcomes      []string `json:"expected_detection_outcomes,omitempty"`
	ExpectedNonSuppressionOutcomes []string `json:"expected_non_suppression_outcomes,omitempty"`
	BlindSpotTags                  []string `json:"blind_spot_tags,omitempty"`
	CoverageLimitations            []string `json:"coverage_limitations,omitempty"`
	EvidenceRefs                   []string `json:"evidence_refs,omitempty"`
}

type MissedDetectionAnalysisContract struct {
	CurrentState               string   `json:"current_state"`
	AnalysisID                 string   `json:"analysis_id"`
	FalseNegativeCases         []string `json:"false_negative_cases,omitempty"`
	MissedDetectionScenarios   []string `json:"missed_detection_scenarios,omitempty"`
	SuppressionCausedMissCases []string `json:"suppression_caused_miss_cases,omitempty"`
	LowSignalMissCases         []string `json:"low_signal_miss_cases,omitempty"`
	CriticalAssetMissCases     []string `json:"critical_asset_miss_cases,omitempty"`
	RootCauseCategories        []string `json:"root_cause_categories,omitempty"`
	RemediationRecommendations []string `json:"remediation_recommendations,omitempty"`
	ReviewerRequired           bool     `json:"reviewer_required"`
	LimitationMessage          string   `json:"limitation_message"`
}

type FalsePositiveFalseNegativeBalanceReviewContract struct {
	CurrentState                    string `json:"current_state"`
	ReviewID                        string `json:"review_id"`
	FalsePositiveRateMetricRef      string `json:"false_positive_rate_metric_ref"`
	FalseNegativeReviewMetricRef    string `json:"false_negative_review_metric_ref"`
	SuppressionCorrectnessMetricRef string `json:"suppression_correctness_metric_ref"`
	MissedDetectionMetricRef        string `json:"missed_detection_metric_ref"`
	FPReductionClaimed              bool   `json:"fp_reduction_claimed"`
	FNRiskReviewed                  bool   `json:"fn_risk_reviewed"`
	SuppressionCausedMissChecked    bool   `json:"suppression_caused_miss_checked"`
	CriticalLowSignalReviewed       bool   `json:"critical_low_signal_reviewed"`
	LimitationMessage               string `json:"limitation_message"`
}

type ConfidenceCalibrationReviewContract struct {
	CurrentState                             string   `json:"current_state"`
	ReviewID                                 string   `json:"review_id"`
	ConfidenceBandCoverage                   []string `json:"confidence_band_coverage,omitempty"`
	ConfidenceCalibrationErrorMetricRef      string   `json:"confidence_calibration_error_metric_ref"`
	HighConfidenceCaseReviewed               bool     `json:"high_confidence_case_reviewed"`
	LowConfidenceCaseReviewed                bool     `json:"low_confidence_case_reviewed"`
	UnknownConfidenceCaseReviewed            bool     `json:"unknown_confidence_case_reviewed"`
	UnsupportedEvidenceHighConfidenceBlocked bool     `json:"unsupported_evidence_high_confidence_blocked"`
	WeaklyInferredHighConfidenceBlocked      bool     `json:"weakly_inferred_high_confidence_blocked"`
	StaleConfidenceLimited                   bool     `json:"stale_confidence_limited"`
	LimitationMessage                        string   `json:"limitation_message"`
}

type VEXCandidateQualityReviewContract struct {
	CurrentState                          string `json:"current_state"`
	ReviewID                              string `json:"review_id"`
	VEXCandidateCount                     int    `json:"vex_candidate_count"`
	ReviewedCandidateCount                int    `json:"reviewed_candidate_count"`
	RejectedCandidateCount                int    `json:"rejected_candidate_count"`
	SupersededCandidateCount              int    `json:"superseded_candidate_count"`
	InsufficientEvidenceCount             int    `json:"insufficient_evidence_count"`
	FinalVEXPublicationBlocked            bool   `json:"final_vex_publication_blocked"`
	NotAffectedRequiresSufficientEvidence bool   `json:"not_affected_requires_sufficient_evidence"`
	StaleOrUnsupportedReviewedBlocked     bool   `json:"stale_or_unsupported_reviewed_blocked"`
	ReviewerRequiredForPublication        bool   `json:"reviewer_required_for_publication"`
	LimitationMessage                     string `json:"limitation_message"`
}

type ReachabilityCalibrationQualityReviewContract struct {
	CurrentState                     string `json:"current_state"`
	ReviewID                         string `json:"review_id"`
	ReachabilityCaseCount            int    `json:"reachability_case_count"`
	PackagePresenceOnlyBlocked       bool   `json:"package_presence_only_blocked"`
	RuntimeLoadedOnlyBlocked         bool   `json:"runtime_loaded_only_blocked"`
	UnsupportedSignalExplicit        bool   `json:"unsupported_signal_explicit"`
	DowngradeEvidenceRequired        bool   `json:"downgrade_evidence_required"`
	ExcludedCriticalDowngradeBlocked bool   `json:"excluded_critical_downgrade_blocked"`
	StaleReachabilityLimited         bool   `json:"stale_reachability_limited"`
	LimitationMessage                string `json:"limitation_message"`
}

type BehavioralCalibrationQualityReviewContract struct {
	CurrentState                  string `json:"current_state"`
	ReviewID                      string `json:"review_id"`
	BaselineCaseCount             int    `json:"baseline_case_count"`
	LearningModeCaseCount         int    `json:"learning_mode_case_count"`
	BaselineFreshnessReviewed     bool   `json:"baseline_freshness_reviewed"`
	StaleOrUnknownBaselineBlocked bool   `json:"stale_or_unknown_baseline_blocked"`
	LearningModeRelaxationBlocked bool   `json:"learning_mode_relaxation_blocked"`
	AutoBaselinePromotionBlocked  bool   `json:"auto_baseline_promotion_blocked"`
	ThresholdDecreaseReviewed     bool   `json:"threshold_decrease_reviewed"`
	FNRiskNoteRequired            bool   `json:"fn_risk_note_required"`
	LimitationMessage             string `json:"limitation_message"`
}

type FederatedWeightingQualityReviewContract struct {
	CurrentState                        string `json:"current_state"`
	ReviewID                            string `json:"review_id"`
	FederatedCaseCount                  int    `json:"federated_case_count"`
	SourceWeightingReviewed             bool   `json:"source_weighting_reviewed"`
	SimilarityGatingReviewed            bool   `json:"similarity_gating_reviewed"`
	LocalEvidenceOverrideReviewed       bool   `json:"local_evidence_override_reviewed"`
	PropagationRedactionReviewed        bool   `json:"propagation_redaction_reviewed"`
	UnknownSourceQualityCapped          bool   `json:"unknown_source_quality_capped"`
	LowSimilarityConfidenceNotIncreased bool   `json:"low_similarity_confidence_not_increased"`
	RawLocalEvidenceNotPropagated       bool   `json:"raw_local_evidence_not_propagated"`
	LimitationMessage                   string `json:"limitation_message"`
}

type SimulationCoverageReviewContract struct {
	CurrentState              string   `json:"current_state"`
	CoverageID                string   `json:"coverage_id"`
	RequiredScenarioClasses   []string `json:"required_scenario_classes,omitempty"`
	CoveredScenarioClasses    []string `json:"covered_scenario_classes,omitempty"`
	MissingScenarioClasses    []string `json:"missing_scenario_classes,omitempty"`
	CriticalMissingClasses    []string `json:"critical_missing_classes,omitempty"`
	CoverageState             string   `json:"coverage_state"`
	CoverageLimitations       []string `json:"coverage_limitations,omitempty"`
	ReplayRefs                []string `json:"replay_refs,omitempty"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	ClaimsExhaustiveDetection bool     `json:"claims_exhaustive_detection"`
}

type IntelligenceQualityScoreboardContract struct {
	CurrentState                     string   `json:"current_state"`
	ScoreboardID                     string   `json:"scoreboard_id"`
	MetricRefs                       []string `json:"metric_refs,omitempty"`
	PrecisionLikeMetricRef           string   `json:"precision_like_metric_ref"`
	FalsePositiveRateRef             string   `json:"false_positive_rate_ref"`
	FalseNegativeReviewRateRef       string   `json:"false_negative_review_rate_ref"`
	OperatorAcceptanceRateRef        string   `json:"operator_acceptance_rate_ref"`
	SuppressionCorrectnessRateRef    string   `json:"suppression_correctness_rate_ref"`
	MissedDetectionReviewRateRef     string   `json:"missed_detection_review_rate_ref"`
	ConfidenceCalibrationErrorRef    string   `json:"confidence_calibration_error_ref"`
	VEXCandidateApprovalRateRef      string   `json:"vex_candidate_approval_rate_ref"`
	LearningModeCompletionQualityRef string   `json:"learning_mode_completion_quality_ref"`
	FederatedSignalUsefulnessRateRef string   `json:"federated_signal_usefulness_rate_ref"`
	FreshnessState                   string   `json:"freshness_state"`
	LimitationMessage                string   `json:"limitation_message"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
	ClaimsUniversalQuality           bool     `json:"claims_universal_quality"`
}

func intelligenceCalibrationValDScenarioClasses() []string {
	return []string{
		"reachability_false_positive",
		"reachability_false_negative",
		"vex_candidate_quality",
		"anomaly_low_signal",
		"drift_behavior_change",
		"learning_mode_boundary",
		"suppression_safety",
		"federated_misweighting",
		"critical_asset_escalation",
		"unsupported_signal_handling",
	}
}

func intelligenceCalibrationValDCoverageStates() []string {
	return []string{
		IntelligenceCalibrationValDCoveragePass,
		IntelligenceCalibrationValDCoveragePartial,
		IntelligenceCalibrationValDCoverageFail,
		IntelligenceCalibrationValDCoverageUnsupported,
	}
}

func IntelligenceCalibrationValDDefensiveSimulationHarnessContract() DefensiveSimulationScenarioHarnessContract {
	return DefensiveSimulationScenarioHarnessContract{
		CurrentState:             IntelligenceCalibrationValDSimulationHarnessStateActive,
		SupportedScenarioClasses: intelligenceCalibrationValDScenarioClasses(),
		HarnessID:                "simulation-harness-001",
		HarnessVersion:           "vald-2026-04-26",
		ScenarioClasses:          intelligenceCalibrationValDScenarioClasses(),
		ScenarioRefs: []string{
			"scenario/reachability-fp-001",
			"scenario/reachability-fn-001",
			"scenario/vex-quality-001",
			"scenario/anomaly-low-signal-001",
			"scenario/drift-change-001",
			"scenario/learning-boundary-001",
			"scenario/suppression-safety-001",
			"scenario/federated-misweighting-001",
			"scenario/critical-asset-escalation-001",
			"scenario/unsupported-signal-001",
		},
		ScenarioCount:           10,
		Replayable:              true,
		ExpectedOutcomesPresent: true,
		ActualOutcomesPresent:   true,
		EvidenceRefs:            []string{"evidence/simulation-harness-001", "evidence/replay-index-001"},
		GeneratedAt:             "2026-04-26T08:00:00Z",
		FreshnessState:          IntelligenceCalibrationFreshnessFresh,
		LimitationMessage:       "simulation harness remains a scoped replayable proof surface and not a universal detection guarantee",
		ProjectionDisclaimer:    "projection_only not_canonical_truth advisory_defensive_simulation_harness",
	}
}

func IntelligenceCalibrationValDScenarioLibraryContract() AdversarialLowSignalScenarioLibraryContract {
	return AdversarialLowSignalScenarioLibraryContract{
		CurrentState:                   IntelligenceCalibrationValDScenarioLibraryStateActive,
		LibraryID:                      "scenario-library-001",
		LibraryVersion:                 "vald-2026-04-26",
		Scenarios:                      []string{"scenario/reachability-fn-001", "scenario/anomaly-low-signal-001", "scenario/suppression-safety-001", "scenario/federated-misweighting-001", "scenario/critical-asset-escalation-001"},
		LowSignalScenarios:             []string{"scenario/anomaly-low-signal-001"},
		AdversarialScenarios:           []string{"scenario/federated-misweighting-001", "scenario/suppression-safety-001"},
		CriticalAssetScenarios:         []string{"scenario/critical-asset-escalation-001"},
		ExpectedDetectionOutcomes:      []string{"detect-low-signal-path", "preserve-critical-escalation"},
		ExpectedNonSuppressionOutcomes: []string{"do-not-suppress-critical-low-signal", "do-not-hide-false-negative-path"},
		BlindSpotTags:                  []string{"runtime_gap", "config-drift-window"},
		CoverageLimitations:            []string{"Library is bounded to reviewed adversarial and low-signal scenarios and is not universal detection proof."},
		EvidenceRefs:                   []string{"evidence/scenario-library-001"},
	}
}

func IntelligenceCalibrationValDMissedDetectionAnalysisContract() MissedDetectionAnalysisContract {
	return MissedDetectionAnalysisContract{
		CurrentState:               IntelligenceCalibrationValDMissedDetectionStateActive,
		AnalysisID:                 "missed-detection-analysis-001",
		FalseNegativeCases:         []string{"case/fn-001"},
		MissedDetectionScenarios:   []string{"scenario/reachability-fn-001", "scenario/anomaly-low-signal-001"},
		SuppressionCausedMissCases: []string{"case/suppression-miss-001"},
		LowSignalMissCases:         []string{"case/low-signal-miss-001"},
		CriticalAssetMissCases:     []string{"case/critical-asset-miss-001"},
		RootCauseCategories:        []string{"insufficient_runtime_context", "overweighted_noise_reduction"},
		RemediationRecommendations: []string{"Require review on sensitivity decreases for critical low-signal paths.", "Replay suppression-candidate scenarios before any governed change."},
		ReviewerRequired:           true,
		LimitationMessage:          "missed-detection analysis remains a bounded operational review surface and is not hidden by false-positive improvements",
	}
}

func IntelligenceCalibrationValDFPFNBalanceReviewContract() FalsePositiveFalseNegativeBalanceReviewContract {
	return FalsePositiveFalseNegativeBalanceReviewContract{
		CurrentState:                    IntelligenceCalibrationValDFPFNBalanceStateActive,
		ReviewID:                        "fp-fn-balance-review-001",
		FalsePositiveRateMetricRef:      "metric/false_positive_rate",
		FalseNegativeReviewMetricRef:    "metric/false_negative_review_rate",
		SuppressionCorrectnessMetricRef: "metric/suppression_correctness_rate",
		MissedDetectionMetricRef:        "metric/missed_detection_review_rate",
		FPReductionClaimed:              true,
		FNRiskReviewed:                  true,
		SuppressionCausedMissChecked:    true,
		CriticalLowSignalReviewed:       true,
		LimitationMessage:               "FP/FN balance review stays scoped to reviewed operational metrics and cannot prove universal intelligence quality",
	}
}

func IntelligenceCalibrationValDConfidenceCalibrationReviewContract() ConfidenceCalibrationReviewContract {
	return ConfidenceCalibrationReviewContract{
		CurrentState:                             IntelligenceCalibrationValDConfidenceReviewStateActive,
		ReviewID:                                 "confidence-review-001",
		ConfidenceBandCoverage:                   intelligenceCalibrationVal0ConfidenceBands(),
		ConfidenceCalibrationErrorMetricRef:      "metric/confidence_calibration_error",
		HighConfidenceCaseReviewed:               true,
		LowConfidenceCaseReviewed:                true,
		UnknownConfidenceCaseReviewed:            true,
		UnsupportedEvidenceHighConfidenceBlocked: true,
		WeaklyInferredHighConfidenceBlocked:      true,
		StaleConfidenceLimited:                   true,
		LimitationMessage:                        "confidence calibration review remains bounded to reviewed calibration outcomes and cannot approve intelligence authority",
	}
}

func IntelligenceCalibrationValDVEXQualityReviewContract() VEXCandidateQualityReviewContract {
	return VEXCandidateQualityReviewContract{
		CurrentState:                          IntelligenceCalibrationValDVEXQualityStateActive,
		ReviewID:                              "vex-quality-review-001",
		VEXCandidateCount:                     12,
		ReviewedCandidateCount:                7,
		RejectedCandidateCount:                2,
		SupersededCandidateCount:              1,
		InsufficientEvidenceCount:             2,
		FinalVEXPublicationBlocked:            true,
		NotAffectedRequiresSufficientEvidence: true,
		StaleOrUnsupportedReviewedBlocked:     true,
		ReviewerRequiredForPublication:        true,
		LimitationMessage:                     "VEX quality review remains advisory and does not publish final VEX or bypass later governance",
	}
}

func IntelligenceCalibrationValDReachabilityQualityReviewContract() ReachabilityCalibrationQualityReviewContract {
	return ReachabilityCalibrationQualityReviewContract{
		CurrentState:                     IntelligenceCalibrationValDReachabilityQualityStateActive,
		ReviewID:                         "reachability-quality-review-001",
		ReachabilityCaseCount:            15,
		PackagePresenceOnlyBlocked:       true,
		RuntimeLoadedOnlyBlocked:         true,
		UnsupportedSignalExplicit:        true,
		DowngradeEvidenceRequired:        true,
		ExcludedCriticalDowngradeBlocked: true,
		StaleReachabilityLimited:         true,
		LimitationMessage:                "reachability quality review remains scoped to reviewed cases and does not mutate canonical priority",
	}
}

func IntelligenceCalibrationValDBehavioralQualityReviewContract() BehavioralCalibrationQualityReviewContract {
	return BehavioralCalibrationQualityReviewContract{
		CurrentState:                  IntelligenceCalibrationValDBehavioralQualityStateActive,
		ReviewID:                      "behavioral-quality-review-001",
		BaselineCaseCount:             8,
		LearningModeCaseCount:         5,
		BaselineFreshnessReviewed:     true,
		StaleOrUnknownBaselineBlocked: true,
		LearningModeRelaxationBlocked: true,
		AutoBaselinePromotionBlocked:  true,
		ThresholdDecreaseReviewed:     true,
		FNRiskNoteRequired:            true,
		LimitationMessage:             "behavioral quality review remains advisory and does not mutate thresholds, enforcement, or baseline state",
	}
}

func IntelligenceCalibrationValDFederatedQualityReviewContract() FederatedWeightingQualityReviewContract {
	return FederatedWeightingQualityReviewContract{
		CurrentState:                        IntelligenceCalibrationValDFederatedQualityStateActive,
		ReviewID:                            "federated-quality-review-001",
		FederatedCaseCount:                  6,
		SourceWeightingReviewed:             true,
		SimilarityGatingReviewed:            true,
		LocalEvidenceOverrideReviewed:       true,
		PropagationRedactionReviewed:        true,
		UnknownSourceQualityCapped:          true,
		LowSimilarityConfidenceNotIncreased: true,
		RawLocalEvidenceNotPropagated:       true,
		LimitationMessage:                   "federated quality review remains bounded to reviewed local reuse discipline and cannot override local evidence",
	}
}

func IntelligenceCalibrationValDSimulationCoverageReviewContract() SimulationCoverageReviewContract {
	return SimulationCoverageReviewContract{
		CurrentState:            IntelligenceCalibrationValDSimulationCoverageStateActive,
		CoverageID:              "simulation-coverage-001",
		RequiredScenarioClasses: intelligenceCalibrationValDScenarioClasses(),
		CoveredScenarioClasses:  intelligenceCalibrationValDScenarioClasses(),
		MissingScenarioClasses:  []string{},
		CriticalMissingClasses:  []string{},
		CoverageState:           IntelligenceCalibrationValDCoveragePass,
		CoverageLimitations:     []string{"Coverage remains bounded to reviewed replay classes and is not exhaustive detection proof."},
		ReplayRefs:              []string{"replay/simulation-harness-001"},
		EvidenceRefs:            []string{"evidence/simulation-coverage-001"},
	}
}

func IntelligenceCalibrationValDQualityScoreboardContract() IntelligenceQualityScoreboardContract {
	return IntelligenceQualityScoreboardContract{
		CurrentState: IntelligenceCalibrationValDQualityScoreboardStateActive,
		ScoreboardID: "quality-scoreboard-001",
		MetricRefs: []string{
			IntelligenceCalibrationMetricFalsePositiveRate,
			IntelligenceCalibrationMetricFalseNegativeReviewRate,
			IntelligenceCalibrationMetricOperatorAcceptanceRate,
			IntelligenceCalibrationMetricSuppressionCorrectness,
			IntelligenceCalibrationMetricMissedDetectionReviewRate,
			IntelligenceCalibrationMetricConfidenceError,
			IntelligenceCalibrationMetricVEXApprovalRate,
			IntelligenceCalibrationMetricLearningCompletion,
			IntelligenceCalibrationMetricFederatedUsefulness,
		},
		PrecisionLikeMetricRef:           "metric/operator_acceptance_rate",
		FalsePositiveRateRef:             "metric/false_positive_rate",
		FalseNegativeReviewRateRef:       "metric/false_negative_review_rate",
		OperatorAcceptanceRateRef:        "metric/operator_acceptance_rate",
		SuppressionCorrectnessRateRef:    "metric/suppression_correctness_rate",
		MissedDetectionReviewRateRef:     "metric/missed_detection_review_rate",
		ConfidenceCalibrationErrorRef:    "metric/confidence_calibration_error",
		VEXCandidateApprovalRateRef:      "metric/vex_candidate_approval_rate",
		LearningModeCompletionQualityRef: "metric/learning_mode_completion_quality",
		FederatedSignalUsefulnessRateRef: "metric/federated_signal_usefulness_rate",
		FreshnessState:                   IntelligenceCalibrationFreshnessFresh,
		LimitationMessage:                "scoreboard metrics are scoped operational indicators and not universal intelligence quality claims",
		ProjectionDisclaimer:             "projection_only not_canonical_truth scoped_operational_indicator",
	}
}

func EvaluateIntelligenceCalibrationValDSimulationHarnessState(model DefensiveSimulationScenarioHarnessContract) string {
	if strings.TrimSpace(model.HarnessID) == "" || strings.TrimSpace(model.HarnessVersion) == "" || len(model.ScenarioClasses) == 0 || len(model.ScenarioRefs) == 0 || model.ScenarioCount <= 0 || len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.GeneratedAt) == "" || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.LimitationMessage) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValDSimulationHarnessStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedScenarioClasses, intelligenceCalibrationValDScenarioClasses()...) || !containsExactTrimmedStringSet(model.ScenarioClasses, intelligenceCalibrationValDScenarioClasses()...) || !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValDSimulationHarnessStatePartial
	}
	if model.ScenarioCount != len(model.ScenarioRefs) || !model.ExpectedOutcomesPresent || !model.ActualOutcomesPresent {
		return IntelligenceCalibrationValDSimulationHarnessStatePartial
	}
	if _, ok := parseIntelligenceCalibrationValBTimestamp(model.GeneratedAt); !ok {
		return IntelligenceCalibrationValDSimulationHarnessStatePartial
	}
	if !model.Replayable && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "replay") {
		return IntelligenceCalibrationValDSimulationHarnessStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessStale && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "stale") {
		return IntelligenceCalibrationValDSimulationHarnessStatePartial
	}
	if model.FreshnessState != IntelligenceCalibrationFreshnessFresh && model.FreshnessState != IntelligenceCalibrationFreshnessStale {
		return IntelligenceCalibrationValDSimulationHarnessStatePartial
	}
	return IntelligenceCalibrationValDSimulationHarnessStateActive
}

func EvaluateIntelligenceCalibrationValDScenarioLibraryState(model AdversarialLowSignalScenarioLibraryContract) string {
	if strings.TrimSpace(model.LibraryID) == "" || strings.TrimSpace(model.LibraryVersion) == "" || len(model.Scenarios) == 0 || len(model.ExpectedDetectionOutcomes) == 0 || len(model.ExpectedNonSuppressionOutcomes) == 0 || len(model.BlindSpotTags) == 0 || len(model.CoverageLimitations) == 0 || len(model.EvidenceRefs) == 0 {
		return IntelligenceCalibrationValDScenarioLibraryStateIncomplete
	}
	if len(model.LowSignalScenarios) == 0 || len(model.AdversarialScenarios) == 0 || len(model.CriticalAssetScenarios) == 0 {
		return IntelligenceCalibrationValDScenarioLibraryStatePartial
	}
	for _, scenarioRef := range model.LowSignalScenarios {
		if !containsTrimmedString(model.Scenarios, scenarioRef) {
			return IntelligenceCalibrationValDScenarioLibraryStatePartial
		}
	}
	for _, scenarioRef := range model.AdversarialScenarios {
		if !containsTrimmedString(model.Scenarios, scenarioRef) {
			return IntelligenceCalibrationValDScenarioLibraryStatePartial
		}
	}
	for _, scenarioRef := range model.CriticalAssetScenarios {
		if !containsTrimmedString(model.Scenarios, scenarioRef) {
			return IntelligenceCalibrationValDScenarioLibraryStatePartial
		}
	}
	return IntelligenceCalibrationValDScenarioLibraryStateActive
}

func EvaluateIntelligenceCalibrationValDMissedDetectionState(model MissedDetectionAnalysisContract) string {
	if strings.TrimSpace(model.AnalysisID) == "" || len(model.FalseNegativeCases) == 0 || len(model.MissedDetectionScenarios) == 0 || len(model.SuppressionCausedMissCases) == 0 || len(model.LowSignalMissCases) == 0 || len(model.RootCauseCategories) == 0 || len(model.RemediationRecommendations) == 0 || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValDMissedDetectionStateIncomplete
	}
	if len(model.CriticalAssetMissCases) > 0 && !model.ReviewerRequired {
		return IntelligenceCalibrationValDMissedDetectionStatePartial
	}
	return IntelligenceCalibrationValDMissedDetectionStateActive
}

func EvaluateIntelligenceCalibrationValDFPFNBalanceState(model FalsePositiveFalseNegativeBalanceReviewContract) string {
	if strings.TrimSpace(model.ReviewID) == "" || strings.TrimSpace(model.FalsePositiveRateMetricRef) == "" || strings.TrimSpace(model.FalseNegativeReviewMetricRef) == "" || strings.TrimSpace(model.SuppressionCorrectnessMetricRef) == "" || strings.TrimSpace(model.MissedDetectionMetricRef) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValDFPFNBalanceStateIncomplete
	}
	if model.FPReductionClaimed && !model.FNRiskReviewed {
		return IntelligenceCalibrationValDFPFNBalanceStatePartial
	}
	if !model.SuppressionCausedMissChecked || !model.CriticalLowSignalReviewed {
		return IntelligenceCalibrationValDFPFNBalanceStatePartial
	}
	return IntelligenceCalibrationValDFPFNBalanceStateActive
}

func EvaluateIntelligenceCalibrationValDConfidenceReviewState(model ConfidenceCalibrationReviewContract) string {
	if strings.TrimSpace(model.ReviewID) == "" || len(model.ConfidenceBandCoverage) == 0 || strings.TrimSpace(model.ConfidenceCalibrationErrorMetricRef) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValDConfidenceReviewStateIncomplete
	}
	if !containsAllTrimmedStrings(model.ConfidenceBandCoverage, IntelligenceCalibrationConfidenceHigh, IntelligenceCalibrationConfidenceLow, IntelligenceCalibrationConfidenceUnknown) {
		return IntelligenceCalibrationValDConfidenceReviewStatePartial
	}
	if !model.HighConfidenceCaseReviewed || !model.LowConfidenceCaseReviewed || !model.UnknownConfidenceCaseReviewed || !model.UnsupportedEvidenceHighConfidenceBlocked || !model.WeaklyInferredHighConfidenceBlocked || !model.StaleConfidenceLimited {
		return IntelligenceCalibrationValDConfidenceReviewStatePartial
	}
	return IntelligenceCalibrationValDConfidenceReviewStateActive
}

func EvaluateIntelligenceCalibrationValDVEXQualityState(model VEXCandidateQualityReviewContract) string {
	if strings.TrimSpace(model.ReviewID) == "" || model.VEXCandidateCount <= 0 || model.ReviewedCandidateCount < 0 || model.RejectedCandidateCount < 0 || model.SupersededCandidateCount < 0 || model.InsufficientEvidenceCount < 0 || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValDVEXQualityStateIncomplete
	}
	if !model.FinalVEXPublicationBlocked || !model.NotAffectedRequiresSufficientEvidence || !model.StaleOrUnsupportedReviewedBlocked || !model.ReviewerRequiredForPublication {
		return IntelligenceCalibrationValDVEXQualityStatePartial
	}
	return IntelligenceCalibrationValDVEXQualityStateActive
}

func EvaluateIntelligenceCalibrationValDReachabilityQualityState(model ReachabilityCalibrationQualityReviewContract) string {
	if strings.TrimSpace(model.ReviewID) == "" || model.ReachabilityCaseCount <= 0 || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValDReachabilityQualityStateIncomplete
	}
	if !model.PackagePresenceOnlyBlocked || !model.RuntimeLoadedOnlyBlocked || !model.UnsupportedSignalExplicit || !model.DowngradeEvidenceRequired || !model.ExcludedCriticalDowngradeBlocked || !model.StaleReachabilityLimited {
		return IntelligenceCalibrationValDReachabilityQualityStatePartial
	}
	return IntelligenceCalibrationValDReachabilityQualityStateActive
}

func EvaluateIntelligenceCalibrationValDBehavioralQualityState(model BehavioralCalibrationQualityReviewContract) string {
	if strings.TrimSpace(model.ReviewID) == "" || model.BaselineCaseCount <= 0 || model.LearningModeCaseCount <= 0 || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValDBehavioralQualityStateIncomplete
	}
	if !model.BaselineFreshnessReviewed || !model.StaleOrUnknownBaselineBlocked || !model.LearningModeRelaxationBlocked || !model.AutoBaselinePromotionBlocked || !model.ThresholdDecreaseReviewed || !model.FNRiskNoteRequired {
		return IntelligenceCalibrationValDBehavioralQualityStatePartial
	}
	return IntelligenceCalibrationValDBehavioralQualityStateActive
}

func EvaluateIntelligenceCalibrationValDFederatedQualityState(model FederatedWeightingQualityReviewContract) string {
	if strings.TrimSpace(model.ReviewID) == "" || model.FederatedCaseCount <= 0 || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValDFederatedQualityStateIncomplete
	}
	if !model.SourceWeightingReviewed || !model.SimilarityGatingReviewed || !model.LocalEvidenceOverrideReviewed || !model.PropagationRedactionReviewed || !model.UnknownSourceQualityCapped || !model.LowSimilarityConfidenceNotIncreased || !model.RawLocalEvidenceNotPropagated {
		return IntelligenceCalibrationValDFederatedQualityStatePartial
	}
	return IntelligenceCalibrationValDFederatedQualityStateActive
}

func EvaluateIntelligenceCalibrationValDSimulationCoverageState(model SimulationCoverageReviewContract) string {
	if strings.TrimSpace(model.CoverageID) == "" || len(model.RequiredScenarioClasses) == 0 || len(model.CoveredScenarioClasses) == 0 || strings.TrimSpace(model.CoverageState) == "" || len(model.EvidenceRefs) == 0 {
		return IntelligenceCalibrationValDSimulationCoverageStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredScenarioClasses, intelligenceCalibrationValDScenarioClasses()...) || !containsTrimmedString(intelligenceCalibrationValDCoverageStates(), model.CoverageState) {
		return IntelligenceCalibrationValDSimulationCoverageStatePartial
	}
	if model.ClaimsExhaustiveDetection {
		return IntelligenceCalibrationValDSimulationCoverageStatePartial
	}
	if len(model.ReplayRefs) == 0 && len(model.CoverageLimitations) == 0 {
		return IntelligenceCalibrationValDSimulationCoverageStatePartial
	}
	if !containsSubstringInTrimmedStrings(model.CoverageLimitations, "not exhaustive") {
		return IntelligenceCalibrationValDSimulationCoverageStatePartial
	}
	if model.CoverageState != IntelligenceCalibrationValDCoveragePass {
		return IntelligenceCalibrationValDSimulationCoverageStatePartial
	}
	if model.CoverageState == IntelligenceCalibrationValDCoveragePass && (len(model.CriticalMissingClasses) > 0 || len(model.MissingScenarioClasses) > 0) {
		return IntelligenceCalibrationValDSimulationCoverageStatePartial
	}
	for _, scenarioClass := range model.RequiredScenarioClasses {
		if !containsTrimmedString(model.CoveredScenarioClasses, scenarioClass) {
			return IntelligenceCalibrationValDSimulationCoverageStatePartial
		}
	}
	for _, scenarioClass := range model.MissingScenarioClasses {
		if !containsTrimmedString(model.RequiredScenarioClasses, scenarioClass) {
			return IntelligenceCalibrationValDSimulationCoverageStatePartial
		}
	}
	for _, scenarioClass := range model.CriticalMissingClasses {
		if !containsTrimmedString(model.RequiredScenarioClasses, scenarioClass) {
			return IntelligenceCalibrationValDSimulationCoverageStatePartial
		}
	}
	return IntelligenceCalibrationValDSimulationCoverageStateActive
}

func EvaluateIntelligenceCalibrationValDQualityScoreboardState(model IntelligenceQualityScoreboardContract) string {
	if strings.TrimSpace(model.ScoreboardID) == "" || len(model.MetricRefs) == 0 || strings.TrimSpace(model.PrecisionLikeMetricRef) == "" || strings.TrimSpace(model.FalsePositiveRateRef) == "" || strings.TrimSpace(model.FalseNegativeReviewRateRef) == "" || strings.TrimSpace(model.OperatorAcceptanceRateRef) == "" || strings.TrimSpace(model.SuppressionCorrectnessRateRef) == "" || strings.TrimSpace(model.MissedDetectionReviewRateRef) == "" || strings.TrimSpace(model.ConfidenceCalibrationErrorRef) == "" || strings.TrimSpace(model.VEXCandidateApprovalRateRef) == "" || strings.TrimSpace(model.LearningModeCompletionQualityRef) == "" || strings.TrimSpace(model.FederatedSignalUsefulnessRateRef) == "" || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.LimitationMessage) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValDQualityScoreboardStateIncomplete
	}
	if !containsTrimmedString(intelligenceCalibrationVal0FreshnessStates(), model.FreshnessState) || !containsAllTrimmedStrings(model.MetricRefs, IntelligenceCalibrationMetricFalsePositiveRate, IntelligenceCalibrationMetricFalseNegativeReviewRate, IntelligenceCalibrationMetricConfidenceError) || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValDQualityScoreboardStatePartial
	}
	if model.ClaimsUniversalQuality || !strings.Contains(strings.TrimSpace(model.LimitationMessage), "scoped operational") {
		return IntelligenceCalibrationValDQualityScoreboardStatePartial
	}
	if model.FreshnessState == IntelligenceCalibrationFreshnessStale && !strings.Contains(strings.TrimSpace(model.LimitationMessage), "stale") {
		return IntelligenceCalibrationValDQualityScoreboardStatePartial
	}
	if model.FreshnessState != IntelligenceCalibrationFreshnessFresh && model.FreshnessState != IntelligenceCalibrationFreshnessStale {
		return IntelligenceCalibrationValDQualityScoreboardStatePartial
	}
	return IntelligenceCalibrationValDQualityScoreboardStateActive
}

func EvaluateIntelligenceCalibrationValDState(val0DependencyState, val0FoundationState, valADependencyState, valAState, valBDependencyState, valBState, valCDependencyState, valCState, simulationHarnessState, scenarioLibraryState, missedDetectionState, fpfnBalanceState, confidenceReviewState, vexQualityState, reachabilityQualityState, behavioralQualityState, federatedQualityState, simulationCoverageState, qualityScoreboardState string) string {
	if strings.TrimSpace(val0DependencyState) != IntelligenceCalibrationVal0StateActive || strings.TrimSpace(val0FoundationState) != IntelligenceCalibrationVal0StateActive || strings.TrimSpace(valADependencyState) != IntelligenceCalibrationValAStateActive || strings.TrimSpace(valAState) != IntelligenceCalibrationValAStateActive || strings.TrimSpace(valBDependencyState) != IntelligenceCalibrationValBStateActive || strings.TrimSpace(valBState) != IntelligenceCalibrationValBStateActive || strings.TrimSpace(valCDependencyState) != IntelligenceCalibrationValCStateActive || strings.TrimSpace(valCState) != IntelligenceCalibrationValCStateActive {
		return IntelligenceCalibrationValDStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(simulationHarnessState),
		strings.TrimSpace(scenarioLibraryState),
		strings.TrimSpace(missedDetectionState),
		strings.TrimSpace(fpfnBalanceState),
		strings.TrimSpace(confidenceReviewState),
		strings.TrimSpace(vexQualityState),
		strings.TrimSpace(reachabilityQualityState),
		strings.TrimSpace(behavioralQualityState),
		strings.TrimSpace(federatedQualityState),
		strings.TrimSpace(simulationCoverageState),
		strings.TrimSpace(qualityScoreboardState),
	} {
		switch state {
		case IntelligenceCalibrationValDSimulationHarnessStateActive,
			IntelligenceCalibrationValDScenarioLibraryStateActive,
			IntelligenceCalibrationValDMissedDetectionStateActive,
			IntelligenceCalibrationValDFPFNBalanceStateActive,
			IntelligenceCalibrationValDConfidenceReviewStateActive,
			IntelligenceCalibrationValDVEXQualityStateActive,
			IntelligenceCalibrationValDReachabilityQualityStateActive,
			IntelligenceCalibrationValDBehavioralQualityStateActive,
			IntelligenceCalibrationValDFederatedQualityStateActive,
			IntelligenceCalibrationValDSimulationCoverageStateActive,
			IntelligenceCalibrationValDQualityScoreboardStateActive:
		case IntelligenceCalibrationValDSimulationHarnessStatePartial,
			IntelligenceCalibrationValDScenarioLibraryStatePartial,
			IntelligenceCalibrationValDMissedDetectionStatePartial,
			IntelligenceCalibrationValDFPFNBalanceStatePartial,
			IntelligenceCalibrationValDConfidenceReviewStatePartial,
			IntelligenceCalibrationValDVEXQualityStatePartial,
			IntelligenceCalibrationValDReachabilityQualityStatePartial,
			IntelligenceCalibrationValDBehavioralQualityStatePartial,
			IntelligenceCalibrationValDFederatedQualityStatePartial,
			IntelligenceCalibrationValDSimulationCoverageStatePartial,
			IntelligenceCalibrationValDQualityScoreboardStatePartial:
			hasPartial = true
		default:
			return IntelligenceCalibrationValDStateIncomplete
		}
	}
	if hasPartial {
		return IntelligenceCalibrationValDStateSubstantial
	}
	return IntelligenceCalibrationValDStateActive
}

func EvaluateIntelligenceCalibrationValDProofsState(val0DependencyState, val0FoundationState, valADependencyState, valAState, valBDependencyState, valBState, valCDependencyState, valCState, simulationHarnessState, scenarioLibraryState, missedDetectionState, fpfnBalanceState, confidenceReviewState, vexQualityState, reachabilityQualityState, behavioralQualityState, federatedQualityState, simulationCoverageState, qualityScoreboardState string, surfaceRefs, evidenceRefs, limitations, whyPoint5NotPass []string, projectionDisclaimer string) string {
	baseState := EvaluateIntelligenceCalibrationValDState(
		val0DependencyState,
		val0FoundationState,
		valADependencyState,
		valAState,
		valBDependencyState,
		valBState,
		valCDependencyState,
		valCState,
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
	if len(surfaceRefs) < 12 || len(evidenceRefs) < 12 || len(limitations) == 0 || len(whyPoint5NotPass) == 0 || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "not_canonical_truth") {
		if baseState == IntelligenceCalibrationValDStateActive {
			return IntelligenceCalibrationValDStateSubstantial
		}
		return baseState
	}
	return baseState
}

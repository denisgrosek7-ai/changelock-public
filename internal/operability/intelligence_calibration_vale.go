package operability

import "strings"

const (
	IntelligenceCalibrationValEDependencyClosureStateActive     = "intelligence_calibration_vale_dependency_closure_active"
	IntelligenceCalibrationValEDependencyClosureStatePartial    = "intelligence_calibration_vale_dependency_closure_partial"
	IntelligenceCalibrationValEDependencyClosureStateIncomplete = "intelligence_calibration_vale_dependency_closure_incomplete"

	IntelligenceCalibrationValECoherenceReviewStateActive     = "intelligence_calibration_vale_coherence_review_active"
	IntelligenceCalibrationValECoherenceReviewStatePartial    = "intelligence_calibration_vale_coherence_review_partial"
	IntelligenceCalibrationValECoherenceReviewStateIncomplete = "intelligence_calibration_vale_coherence_review_incomplete"

	IntelligenceCalibrationValEPassRuleStateActive     = "intelligence_calibration_vale_pass_rule_active"
	IntelligenceCalibrationValEPassRuleStatePartial    = "intelligence_calibration_vale_pass_rule_partial"
	IntelligenceCalibrationValEPassRuleStateIncomplete = "intelligence_calibration_vale_pass_rule_incomplete"

	IntelligenceCalibrationValEBoundaryReviewStateActive     = "intelligence_calibration_vale_advisory_boundary_active"
	IntelligenceCalibrationValEBoundaryReviewStatePartial    = "intelligence_calibration_vale_advisory_boundary_partial"
	IntelligenceCalibrationValEBoundaryReviewStateIncomplete = "intelligence_calibration_vale_advisory_boundary_incomplete"

	IntelligenceCalibrationValEReachabilityVEXSafetyStateActive     = "intelligence_calibration_vale_reachability_vex_safety_active"
	IntelligenceCalibrationValEReachabilityVEXSafetyStatePartial    = "intelligence_calibration_vale_reachability_vex_safety_partial"
	IntelligenceCalibrationValEReachabilityVEXSafetyStateIncomplete = "intelligence_calibration_vale_reachability_vex_safety_incomplete"

	IntelligenceCalibrationValEBehavioralLearningSafetyStateActive     = "intelligence_calibration_vale_behavioral_learning_safety_active"
	IntelligenceCalibrationValEBehavioralLearningSafetyStatePartial    = "intelligence_calibration_vale_behavioral_learning_safety_partial"
	IntelligenceCalibrationValEBehavioralLearningSafetyStateIncomplete = "intelligence_calibration_vale_behavioral_learning_safety_incomplete"

	IntelligenceCalibrationValEFeedbackFederatedSafetyStateActive     = "intelligence_calibration_vale_feedback_federated_safety_active"
	IntelligenceCalibrationValEFeedbackFederatedSafetyStatePartial    = "intelligence_calibration_vale_feedback_federated_safety_partial"
	IntelligenceCalibrationValEFeedbackFederatedSafetyStateIncomplete = "intelligence_calibration_vale_feedback_federated_safety_incomplete"

	IntelligenceCalibrationValESimulationQualityStateActive     = "intelligence_calibration_vale_simulation_quality_review_active"
	IntelligenceCalibrationValESimulationQualityStatePartial    = "intelligence_calibration_vale_simulation_quality_review_partial"
	IntelligenceCalibrationValESimulationQualityStateIncomplete = "intelligence_calibration_vale_simulation_quality_review_incomplete"

	IntelligenceCalibrationValERegressionClosureStateActive     = "intelligence_calibration_vale_regression_closure_active"
	IntelligenceCalibrationValERegressionClosureStatePartial    = "intelligence_calibration_vale_regression_closure_partial"
	IntelligenceCalibrationValERegressionClosureStateIncomplete = "intelligence_calibration_vale_regression_closure_incomplete"

	IntelligenceCalibrationValEStateIncomplete  = "intelligence_calibration_vale_incomplete"
	IntelligenceCalibrationValEStateSubstantial = "intelligence_calibration_vale_substantially_ready"
	IntelligenceCalibrationValEStateActive      = "intelligence_calibration_vale_active"

	IntelligenceCalibrationValEDependencyPass        = "pass"
	IntelligenceCalibrationValEDependencyFail        = "fail"
	IntelligenceCalibrationValEDependencyIncomplete  = "incomplete"
	IntelligenceCalibrationValEDependencyPartial     = "partial"
	IntelligenceCalibrationValEDependencyUnsupported = "unsupported"

	IntelligenceCalibrationValEReviewPass        = "pass"
	IntelligenceCalibrationValEReviewFail        = "fail"
	IntelligenceCalibrationValEReviewWarning     = "warning"
	IntelligenceCalibrationValEReviewBlocked     = "blocked"
	IntelligenceCalibrationValEReviewUnsupported = "unsupported"
	IntelligenceCalibrationValEReviewNotRun      = "not_run"
)

type IntelligenceCalibrationIntegratedDependencyClosure struct {
	CurrentState           string   `json:"current_state"`
	Val0State              string   `json:"val_0_state"`
	ValAState              string   `json:"val_a_state"`
	ValBState              string   `json:"val_b_state"`
	ValCState              string   `json:"val_c_state"`
	ValDState              string   `json:"val_d_state"`
	ValEState              string   `json:"val_e_state"`
	DependencyStatus       string   `json:"dependency_status"`
	MissingVals            []string `json:"missing_vals,omitempty"`
	InactiveVals           []string `json:"inactive_vals,omitempty"`
	InconsistentVals       []string `json:"inconsistent_vals,omitempty"`
	DependencyEvidenceRefs []string `json:"dependency_evidence_refs,omitempty"`
	DependencySurfaceRefs  []string `json:"dependency_surface_refs,omitempty"`
	ClosureGeneratedAt     string   `json:"closure_generated_at"`
	ProofStatesObserved    bool     `json:"proof_states_observed"`
	ProjectionDisclaimer   string   `json:"projection_disclaimer"`
}

type IntelligenceCalibrationCrossValCoherenceReview struct {
	CurrentState                                      string   `json:"current_state"`
	CoherenceState                                    string   `json:"coherence_state"`
	SupportedCoherenceStates                          []string `json:"supported_coherence_states,omitempty"`
	CheckedLinks                                      []string `json:"checked_links,omitempty"`
	MissingLinks                                      []string `json:"missing_links,omitempty"`
	InconsistentLinks                                 []string `json:"inconsistent_links,omitempty"`
	CarriedForwardLimitations                         []string `json:"carried_forward_limitations,omitempty"`
	EvidenceRefs                                      []string `json:"evidence_refs,omitempty"`
	SurfaceRefs                                       []string `json:"surface_refs,omitempty"`
	Val0ContractsUsedByLaterVals                      bool     `json:"val_0_contracts_used_by_later_vals"`
	ValAReachabilityVEXGuardrailsRespectedByValD      bool     `json:"val_a_reachability_vex_guardrails_respected_by_val_d"`
	ValBBehavioralLearningGuardrailsRespectedByValD   bool     `json:"val_b_behavioral_learning_guardrails_respected_by_val_d"`
	ValCFeedbackSuppressionFederatedRespectedByValD   bool     `json:"val_c_feedback_suppression_federated_guardrails_respected_by_val_d"`
	ValDSimulationQualityCoversRequiredPreviousSlices bool     `json:"val_d_simulation_quality_covers_required_previous_slices"`
	NoPriorValClaimsPoint5Pass                        bool     `json:"no_prior_val_claims_point_5_completion"`
	LimitationsCarriedForward                         bool     `json:"limitations_carried_forward"`
	AdvisoryProjectionBoundariesPreserved             bool     `json:"advisory_projection_boundaries_preserved"`
	ProjectionDisclaimer                              string   `json:"projection_disclaimer"`
}

type IntelligenceCalibrationPoint5IntegratedPassRule struct {
	CurrentState         string   `json:"current_state"`
	Point5State          string   `json:"point_5_state"`
	PassCriteriaMet      bool     `json:"pass_criteria_met"`
	PassBlockers         []string `json:"pass_blockers,omitempty"`
	PassWarnings         []string `json:"pass_warnings,omitempty"`
	PassLimitations      []string `json:"pass_limitations,omitempty"`
	RequiredVals         []string `json:"required_vals,omitempty"`
	ActiveVals           []string `json:"active_vals,omitempty"`
	MissingVals          []string `json:"missing_vals,omitempty"`
	PartialVals          []string `json:"partial_vals,omitempty"`
	UnsupportedVals      []string `json:"unsupported_vals,omitempty"`
	ValEState            string   `json:"val_e_state"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type IntelligenceCalibrationIntegratedAdvisoryBoundaryReview struct {
	CurrentState                                  string   `json:"current_state"`
	BoundaryState                                 string   `json:"boundary_state"`
	SupportedBoundaryStates                       []string `json:"supported_boundary_states,omitempty"`
	CheckedSurfaces                               []string `json:"checked_surfaces,omitempty"`
	ProjectionSurfaces                            []string `json:"projection_surfaces,omitempty"`
	ViolationSurfaces                             []string `json:"violation_surfaces,omitempty"`
	GovernanceRefs                                []string `json:"governance_refs,omitempty"`
	EvidenceRefs                                  []string `json:"evidence_refs,omitempty"`
	LimitationMessage                             string   `json:"limitation_message"`
	ProjectionDisclaimer                          string   `json:"projection_disclaimer"`
	EvidenceSpineRemainsCanonical                 bool     `json:"evidence_spine_remains_canonical"`
	CalibrationOutputsRemainProjections           bool     `json:"calibration_outputs_remain_projections"`
	ConfidenceScoresRemainAdvisory                bool     `json:"confidence_scores_remain_advisory"`
	ReachabilityInferenceRemainsAdvisory          bool     `json:"reachability_inference_remains_advisory"`
	VEXCandidateOutputsRemainCandidates           bool     `json:"vex_candidate_outputs_remain_candidates"`
	BehavioralBaselinesRemainAdvisory             bool     `json:"behavioral_baselines_remain_advisory"`
	LearningModeRemainsBoundedReviewRequired      bool     `json:"learning_mode_remains_bounded_review_required"`
	FeedbackRemainsAdvisory                       bool     `json:"feedback_remains_advisory"`
	SuppressionRemainsCandidateReviewBound        bool     `json:"suppression_remains_candidate_review_bound"`
	FederatedSignalsRemainAdvisory                bool     `json:"federated_signals_remain_advisory"`
	SimulationMetricsRemainOperationalIndicators  bool     `json:"simulation_metrics_remain_operational_indicators"`
	IntegratedClosureSummaryRemainsProjectionOnly bool     `json:"integrated_closure_summary_remains_projection_only"`
	NoMutationWithoutGovernance                   bool     `json:"no_mutation_without_governance"`
	FinalVEXPublicationBlocked                    bool     `json:"final_vex_publication_blocked"`
	FederatedOverrideLocalEvidenceBlocked         bool     `json:"federated_override_local_evidence_blocked"`
	LearningModeCriticalControlRelaxationBlocked  bool     `json:"learning_mode_critical_control_relaxation_blocked"`
}

type IntelligenceCalibrationIntegratedReachabilityVEXSafetyReview struct {
	CurrentState                                       string   `json:"current_state"`
	ReachabilityVEXState                               string   `json:"reachability_vex_state"`
	SupportedReviewStates                              []string `json:"supported_review_states,omitempty"`
	CheckedControls                                    []string `json:"checked_controls,omitempty"`
	Blockers                                           []string `json:"blockers,omitempty"`
	Warnings                                           []string `json:"warnings,omitempty"`
	Limitations                                        []string `json:"limitations,omitempty"`
	EvidenceRefs                                       []string `json:"evidence_refs,omitempty"`
	SurfaceRefs                                        []string `json:"surface_refs,omitempty"`
	ProjectionDisclaimer                               string   `json:"projection_disclaimer"`
	PackagePresenceOnlyBlocked                         bool     `json:"package_presence_only_blocked"`
	RuntimeLoadedOnlyBlocked                           bool     `json:"runtime_loaded_only_blocked"`
	UnsupportedReachabilityRemainsExplicit             bool     `json:"unsupported_reachability_remains_explicit"`
	DowngradeRequiresEvidenceExplanationExpiryRollback bool     `json:"downgrade_requires_evidence_explanation_expiry_rollback"`
	CriticalDowngradeReviewBound                       bool     `json:"critical_downgrade_review_bound"`
	VEXCandidateNotFinalVEX                            bool     `json:"vex_candidate_not_final_vex"`
	FinalVEXClaimBlocked                               bool     `json:"final_vex_claim_blocked"`
	PublicationAllowedBlocked                          bool     `json:"publication_allowed_blocked"`
	InsufficientEvidenceBlocksNotAffected              bool     `json:"insufficient_evidence_blocks_not_affected"`
	StaleUnsupportedEvidenceBlocksReviewedVEXCandidate bool     `json:"stale_unsupported_evidence_blocks_reviewed_vex_candidate"`
	ExplanationDistinguishesNotEvidencedFromSafe       bool     `json:"explanation_distinguishes_not_evidenced_from_safe"`
}

type IntelligenceCalibrationIntegratedBehavioralLearningSafetyReview struct {
	CurrentState                                               string   `json:"current_state"`
	BehavioralLearningState                                    string   `json:"behavioral_learning_state"`
	SupportedReviewStates                                      []string `json:"supported_review_states,omitempty"`
	CheckedControls                                            []string `json:"checked_controls,omitempty"`
	Blockers                                                   []string `json:"blockers,omitempty"`
	Warnings                                                   []string `json:"warnings,omitempty"`
	Limitations                                                []string `json:"limitations,omitempty"`
	EvidenceRefs                                               []string `json:"evidence_refs,omitempty"`
	SurfaceRefs                                                []string `json:"surface_refs,omitempty"`
	ProjectionDisclaimer                                       string   `json:"projection_disclaimer"`
	ActiveBaselineFreshnessFresh                               bool     `json:"active_baseline_freshness_fresh"`
	StaleUnknownUnsupportedBaselineBlocked                     bool     `json:"stale_unknown_unsupported_baseline_blocked"`
	BaselineObservationWindowBoundedAndTimestampValidated      bool     `json:"baseline_observation_window_bounded_and_timestamp_validated"`
	LearningModeTimestampsRFC3339AndBounded                    bool     `json:"learning_mode_timestamps_rfc3339_and_bounded"`
	LearningModeCriticalControlRelaxationBlocked               bool     `json:"learning_mode_critical_control_relaxation_blocked"`
	LearningModeCriticalAlertSuppressionBlocked                bool     `json:"learning_mode_critical_alert_suppression_blocked"`
	LearningModeAutoBaselinePromotionBlocked                   bool     `json:"learning_mode_auto_baseline_promotion_blocked"`
	ThresholdDecreaseRequiresFNRiskNoteAndReview               bool     `json:"threshold_decrease_requires_fn_risk_note_and_review"`
	LowerPriorityCandidateRequiresReviewerGate                 bool     `json:"lower_priority_candidate_requires_reviewer_gate"`
	BaselineAdoptionRequiresReviewerApprovalRollbackGovernance bool     `json:"baseline_adoption_requires_reviewer_approval_rollback_governance"`
}

type IntelligenceCalibrationIntegratedFeedbackFederatedSafetyReview struct {
	CurrentState                                                  string   `json:"current_state"`
	FeedbackFederatedState                                        string   `json:"feedback_federated_state"`
	SupportedReviewStates                                         []string `json:"supported_review_states,omitempty"`
	CheckedControls                                               []string `json:"checked_controls,omitempty"`
	Blockers                                                      []string `json:"blockers,omitempty"`
	Warnings                                                      []string `json:"warnings,omitempty"`
	Limitations                                                   []string `json:"limitations,omitempty"`
	EvidenceRefs                                                  []string `json:"evidence_refs,omitempty"`
	SurfaceRefs                                                   []string `json:"surface_refs,omitempty"`
	ProjectionDisclaimer                                          string   `json:"projection_disclaimer"`
	FeedbackDoesNotMutateIntelligence                             bool     `json:"feedback_does_not_mutate_intelligence"`
	FeedbackDoesNotDirectlySuppressSignals                        bool     `json:"feedback_does_not_directly_suppress_signals"`
	FeedbackDoesNotDirectlyLowerPriority                          bool     `json:"feedback_does_not_directly_lower_priority"`
	FalseNegativeAndMissedSeverityRemainVisible                   bool     `json:"false_negative_and_missed_severity_remain_visible"`
	TuningProposalsRemainAdvisoryReviewBound                      bool     `json:"tuning_proposals_remain_advisory_review_bound"`
	SuppressionCandidatesRequireExpiryScopeReviewerRollbackReopen bool     `json:"suppression_candidates_require_expiry_scope_reviewer_rollback_reopen"`
	SuppressionDoesNotDeleteEvidence                              bool     `json:"suppression_does_not_delete_evidence"`
	SuppressionDoesNotHideFalseNegativePath                       bool     `json:"suppression_does_not_hide_false_negative_path"`
	LocalCalibrationReviewDoesNotMutateActiveCalibration          bool     `json:"local_calibration_review_does_not_mutate_active_calibration"`
	FederatedSignalsDoNotProduceLocalSafeState                    bool     `json:"federated_signals_do_not_produce_local_safe_state"`
	SimilarityGatingDoesNotOverrideLocalEvidence                  bool     `json:"similarity_gating_does_not_override_local_evidence"`
	LocalEvidenceWinsOverFederatedSignal                          bool     `json:"local_evidence_wins_over_federated_signal"`
	PropagationDisabledOrAdvisoryOnlyByDefault                    bool     `json:"propagation_disabled_or_advisory_only_by_default"`
	RawLocalEvidenceDoesNotPropagate                              bool     `json:"raw_local_evidence_does_not_propagate"`
}

type IntelligenceCalibrationIntegratedSimulationQualityReview struct {
	CurrentState                                                 string   `json:"current_state"`
	SimulationQualityState                                       string   `json:"simulation_quality_state"`
	SupportedReviewStates                                        []string `json:"supported_review_states,omitempty"`
	CheckedControls                                              []string `json:"checked_controls,omitempty"`
	MissingCoverage                                              []string `json:"missing_coverage,omitempty"`
	Blockers                                                     []string `json:"blockers,omitempty"`
	Warnings                                                     []string `json:"warnings,omitempty"`
	Limitations                                                  []string `json:"limitations,omitempty"`
	EvidenceRefs                                                 []string `json:"evidence_refs,omitempty"`
	SurfaceRefs                                                  []string `json:"surface_refs,omitempty"`
	ProjectionDisclaimer                                         string   `json:"projection_disclaimer"`
	HarnessPresentReplayableOrExplicitlyLimited                  bool     `json:"harness_present_replayable_or_explicitly_limited"`
	ExpectedAndActualOutcomesRepresented                         bool     `json:"expected_and_actual_outcomes_represented"`
	LowSignalAndAdversarialScenariosRepresented                  bool     `json:"low_signal_and_adversarial_scenarios_represented"`
	MissedDetectionIncludesFalseNegativeVisibility               bool     `json:"missed_detection_includes_false_negative_visibility"`
	SuppressionCausedMissAnalysisRepresented                     bool     `json:"suppression_caused_miss_analysis_represented"`
	FPReductionCannotHideFNRisk                                  bool     `json:"fp_reduction_cannot_hide_fn_risk"`
	ConfidenceReviewBlocksUnsupportedWeakInferenceHighConfidence bool     `json:"confidence_review_blocks_unsupported_weak_inference_high_confidence"`
	VEXQualityBlocksFinalPublicationAndInsufficientNotAffected   bool     `json:"vex_quality_blocks_final_publication_and_insufficient_not_affected"`
	ReachabilityQualityPreservesPackageRuntimeGuardrails         bool     `json:"reachability_quality_preserves_package_runtime_guardrails"`
	BehavioralQualityPreservesLearningBaselineGuardrails         bool     `json:"behavioral_quality_preserves_learning_baseline_guardrails"`
	FederatedQualityPreservesLocalEvidenceBoundary               bool     `json:"federated_quality_preserves_local_evidence_boundary"`
	CoverageDoesNotClaimExhaustiveDetection                      bool     `json:"coverage_does_not_claim_exhaustive_detection"`
	ScoreboardIncludesFPAndFNMetrics                             bool     `json:"scoreboard_includes_fp_and_fn_metrics"`
	ScoreboardDoesNotClaimUniversalIntelligenceQuality           bool     `json:"scoreboard_does_not_claim_universal_intelligence_quality"`
}

type IntelligenceCalibrationIntegratedRegressionClosure struct {
	CurrentState                        string   `json:"current_state"`
	RegressionState                     string   `json:"regression_state"`
	CoveredCategories                   []string `json:"covered_categories,omitempty"`
	MissingCategories                   []string `json:"missing_categories,omitempty"`
	CriticalMissingCategories           []string `json:"critical_missing_categories,omitempty"`
	LimitationMessage                   string   `json:"limitation_message"`
	Val0Coverage                        bool     `json:"val_0_coverage"`
	ValACoverage                        bool     `json:"val_a_coverage"`
	ValBCoverage                        bool     `json:"val_b_coverage"`
	ValCCoverage                        bool     `json:"val_c_coverage"`
	ValDCoverage                        bool     `json:"val_d_coverage"`
	DependencyFailClosedCoverage        bool     `json:"dependency_fail_closed_coverage"`
	CanonicalAdvisoryBoundaryCoverage   bool     `json:"canonical_advisory_boundary_coverage"`
	VEXNoFinalPublicationCoverage       bool     `json:"vex_no_final_publication_guardrail_coverage"`
	LearningModeCriticalControlCoverage bool     `json:"learning_mode_critical_control_guardrail_coverage"`
	SuppressionFalseNegativeCoverage    bool     `json:"suppression_false_negative_guardrail_coverage"`
	FederatedLocalEvidenceWinsCoverage  bool     `json:"federated_local_evidence_wins_guardrail_coverage"`
	FPFNBalanceCoverage                 bool     `json:"fp_fn_balance_coverage"`
	IntegratedPassRuleCoverage          bool     `json:"integrated_pass_rule_coverage"`
}

func intelligenceCalibrationValERequiredVals() []string {
	return []string{"val_0", "val_a", "val_b", "val_c", "val_d", "val_e"}
}

func intelligenceCalibrationValEDependencyStatuses() []string {
	return []string{
		IntelligenceCalibrationValEDependencyPass,
		IntelligenceCalibrationValEDependencyFail,
		IntelligenceCalibrationValEDependencyIncomplete,
		IntelligenceCalibrationValEDependencyPartial,
		IntelligenceCalibrationValEDependencyUnsupported,
	}
}

func intelligenceCalibrationValEReviewStates() []string {
	return []string{
		IntelligenceCalibrationValEReviewPass,
		IntelligenceCalibrationValEReviewFail,
		IntelligenceCalibrationValEReviewWarning,
		IntelligenceCalibrationValEReviewBlocked,
		IntelligenceCalibrationValEReviewUnsupported,
		IntelligenceCalibrationValEReviewNotRun,
	}
}

func intelligenceCalibrationValECoherenceLinks() []string {
	return []string{
		"val0.contracts->vala.reachability_vex_calibration",
		"val0.contracts->valb.behavioral_learning_mode",
		"val0.contracts->valc.feedback_suppression_federated_tuning",
		"val0.contracts->vald.simulation_quality_gate",
		"vala.guardrails->vald.reachability_vex_quality_review",
		"valb.guardrails->vald.behavioral_learning_quality_review",
		"valc.guardrails->vald.feedback_federated_quality_review",
		"vald.simulation_quality_gate->vale.integrated_closure",
		"prior_vals->point5_not_complete_until_vale",
		"limitations->vale.integrated_closure",
	}
}

func intelligenceCalibrationValERegressionCategories() []string {
	return []string{
		"val0_foundation_contracts",
		"vala_reachability_vex_calibration",
		"valb_behavioral_baseline_learning_mode",
		"valc_feedback_suppression_federated_tuning",
		"vald_defensive_simulation_quality_measurement",
		"dependency_fail_closed_behavior",
		"canonical_advisory_boundary",
		"vex_no_final_publication_guardrail",
		"learning_mode_critical_control_guardrail",
		"suppression_false_negative_guardrail",
		"federated_local_evidence_wins_guardrail",
		"fp_fn_balance",
		"integrated_pass_rule",
	}
}

func IntelligenceCalibrationValEDependencyClosureContract() IntelligenceCalibrationIntegratedDependencyClosure {
	return IntelligenceCalibrationIntegratedDependencyClosure{
		CurrentState:     "intelligence_calibration_vale_dependency_closure_ready",
		Val0State:        IntelligenceCalibrationVal0StateActive,
		ValAState:        IntelligenceCalibrationValAStateActive,
		ValBState:        IntelligenceCalibrationValBStateActive,
		ValCState:        IntelligenceCalibrationValCStateActive,
		ValDState:        IntelligenceCalibrationValDStateActive,
		ValEState:        IntelligenceCalibrationValEStateActive,
		DependencyStatus: IntelligenceCalibrationValEDependencyPass,
		DependencyEvidenceRefs: []string{
			"val0_proofs",
			"vala_proofs",
			"valb_proofs",
			"valc_proofs",
			"vald_proofs",
		},
		DependencySurfaceRefs: []string{
			"/v1/intelligence/calibration/val0/proofs",
			"/v1/intelligence/calibration/vala/proofs",
			"/v1/intelligence/calibration/valb/proofs",
			"/v1/intelligence/calibration/valc/proofs",
			"/v1/intelligence/calibration/vald/proofs",
			"/v1/intelligence/calibration/vale/dependency-closure",
			"/v1/intelligence/calibration/vale/proofs",
		},
		ClosureGeneratedAt:   "2026-04-26T09:00:00Z",
		ProofStatesObserved:  true,
		ProjectionDisclaimer: "projection_only not_canonical_truth integrated_dependency_closure",
	}
}

func IntelligenceCalibrationValECoherenceReviewContract() IntelligenceCalibrationCrossValCoherenceReview {
	return IntelligenceCalibrationCrossValCoherenceReview{
		CurrentState:             "intelligence_calibration_vale_coherence_review_ready",
		CoherenceState:           IntelligenceCalibrationValEReviewPass,
		SupportedCoherenceStates: intelligenceCalibrationValEReviewStates(),
		CheckedLinks:             intelligenceCalibrationValECoherenceLinks(),
		CarriedForwardLimitations: []string{
			"val0: foundation contracts remain fail-closed and projection-only.",
			"vala: reachability and VEX outputs remain advisory and cannot publish final VEX.",
			"valb: behavioral baseline and learning mode remain bounded and review-required.",
			"valc: feedback, suppression candidates, and federated signals remain advisory and non-mutating.",
			"vald: simulation and quality outputs remain scoped operational indicators only.",
		},
		EvidenceRefs: []string{
			"val0_proofs",
			"vala_proofs",
			"valb_proofs",
			"valc_proofs",
			"vald_proofs",
		},
		SurfaceRefs: []string{
			"/v1/intelligence/calibration/val0/proofs",
			"/v1/intelligence/calibration/vala/proofs",
			"/v1/intelligence/calibration/valb/proofs",
			"/v1/intelligence/calibration/valc/proofs",
			"/v1/intelligence/calibration/vald/proofs",
			"/v1/intelligence/calibration/vale/coherence-review",
			"/v1/intelligence/calibration/vale/proofs",
		},
		Val0ContractsUsedByLaterVals:                      true,
		ValAReachabilityVEXGuardrailsRespectedByValD:      true,
		ValBBehavioralLearningGuardrailsRespectedByValD:   true,
		ValCFeedbackSuppressionFederatedRespectedByValD:   true,
		ValDSimulationQualityCoversRequiredPreviousSlices: true,
		NoPriorValClaimsPoint5Pass:                        true,
		LimitationsCarriedForward:                         true,
		AdvisoryProjectionBoundariesPreserved:             true,
		ProjectionDisclaimer:                              "projection_only not_canonical_truth cross_val_coherence_review",
	}
}

func IntelligenceCalibrationValEPassRuleContract() IntelligenceCalibrationPoint5IntegratedPassRule {
	return IntelligenceCalibrationPoint5IntegratedPassRule{
		CurrentState:    "intelligence_calibration_vale_pass_rule_ready",
		Point5State:     IntelligenceCalibrationPoint5StatePass,
		PassCriteriaMet: true,
		PassWarnings: []string{
			"Integrated closure remains projection-only and does not replace canonical truth or later governance.",
		},
		PassLimitations: []string{
			"Point 5 pass remains an integrated closure summary over accepted Val 0 through Val E slices.",
		},
		RequiredVals:         intelligenceCalibrationValERequiredVals(),
		ActiveVals:           intelligenceCalibrationValERequiredVals(),
		ValEState:            IntelligenceCalibrationValEStateActive,
		ProjectionDisclaimer: "projection_only not_canonical_truth integrated_point5_pass_rule",
	}
}

func IntelligenceCalibrationValEBoundaryReviewContract() IntelligenceCalibrationIntegratedAdvisoryBoundaryReview {
	return IntelligenceCalibrationIntegratedAdvisoryBoundaryReview{
		CurrentState:            "intelligence_calibration_vale_advisory_boundary_ready",
		BoundaryState:           IntelligenceCalibrationValEReviewPass,
		SupportedBoundaryStates: intelligenceCalibrationValEReviewStates(),
		CheckedSurfaces: []string{
			"/v1/intelligence/calibration/val0/proofs",
			"/v1/intelligence/calibration/vala/publication-guardrail",
			"/v1/intelligence/calibration/valb/safety-guardrails",
			"/v1/intelligence/calibration/valc/local-override",
			"/v1/intelligence/calibration/valc/propagation-policy",
			"/v1/intelligence/calibration/vald/quality-scoreboard",
			"/v1/intelligence/calibration/vale/proofs",
		},
		ProjectionSurfaces: []string{
			"calibration_dataset_outputs",
			"confidence_scores",
			"reachability_inference",
			"vex_candidates",
			"behavioral_baselines",
			"learning_mode",
			"feedback",
			"suppression_candidates",
			"federated_signals",
			"simulation_metrics",
			"integrated_closure_summary",
		},
		GovernanceRefs: []string{
			"canonical_truth_rule",
			"governance_boundary_review_contract",
			"evidence_spine",
		},
		EvidenceRefs: []string{
			"evidence_spine",
			"val0_proofs",
			"vala_proofs",
			"valb_proofs",
			"valc_proofs",
			"vald_proofs",
		},
		LimitationMessage:                             "projection_only not_canonical_truth integrated advisory boundary review",
		ProjectionDisclaimer:                          "projection_only not_canonical_truth integrated advisory boundary review",
		EvidenceSpineRemainsCanonical:                 true,
		CalibrationOutputsRemainProjections:           true,
		ConfidenceScoresRemainAdvisory:                true,
		ReachabilityInferenceRemainsAdvisory:          true,
		VEXCandidateOutputsRemainCandidates:           true,
		BehavioralBaselinesRemainAdvisory:             true,
		LearningModeRemainsBoundedReviewRequired:      true,
		FeedbackRemainsAdvisory:                       true,
		SuppressionRemainsCandidateReviewBound:        true,
		FederatedSignalsRemainAdvisory:                true,
		SimulationMetricsRemainOperationalIndicators:  true,
		IntegratedClosureSummaryRemainsProjectionOnly: true,
		NoMutationWithoutGovernance:                   true,
		FinalVEXPublicationBlocked:                    true,
		FederatedOverrideLocalEvidenceBlocked:         true,
		LearningModeCriticalControlRelaxationBlocked:  true,
	}
}

func IntelligenceCalibrationValEReachabilityVEXSafetyReviewContract() IntelligenceCalibrationIntegratedReachabilityVEXSafetyReview {
	return IntelligenceCalibrationIntegratedReachabilityVEXSafetyReview{
		CurrentState:                           "intelligence_calibration_vale_reachability_vex_safety_ready",
		ReachabilityVEXState:                   IntelligenceCalibrationValEReviewPass,
		SupportedReviewStates:                  intelligenceCalibrationValEReviewStates(),
		CheckedControls:                        []string{"package_presence_only", "runtime_loaded_only", "unsupported_reachability_signal", "downgrade_evidence_explanation_expiry_rollback", "critical_class_downgrade_review_gate", "candidate_not_final_vex", "final_vex_claim_blocked", "publication_allowed_blocked", "insufficient_evidence_blocks_not_affected", "stale_unsupported_evidence_blocks_reviewed_candidate", "not_evidenced_distinct_from_safe"},
		Warnings:                               []string{"Integrated reachability/VEX closure remains advisory and review-bound."},
		Limitations:                            []string{"Reachability and VEX closure remains projection-only and does not publish final VEX."},
		EvidenceRefs:                           []string{"vala_proofs", "vald_proofs", "evidence_spine"},
		SurfaceRefs:                            []string{"/v1/intelligence/calibration/vala/proofs", "/v1/intelligence/calibration/vald/proofs", "/v1/intelligence/calibration/vale/reachability-vex-safety"},
		ProjectionDisclaimer:                   "projection_only not_canonical_truth integrated reachability vex safety review",
		PackagePresenceOnlyBlocked:             true,
		RuntimeLoadedOnlyBlocked:               true,
		UnsupportedReachabilityRemainsExplicit: true,
		DowngradeRequiresEvidenceExplanationExpiryRollback: true,
		CriticalDowngradeReviewBound:                       true,
		VEXCandidateNotFinalVEX:                            true,
		FinalVEXClaimBlocked:                               true,
		PublicationAllowedBlocked:                          true,
		InsufficientEvidenceBlocksNotAffected:              true,
		StaleUnsupportedEvidenceBlocksReviewedVEXCandidate: true,
		ExplanationDistinguishesNotEvidencedFromSafe:       true,
	}
}

func IntelligenceCalibrationValEBehavioralLearningSafetyReviewContract() IntelligenceCalibrationIntegratedBehavioralLearningSafetyReview {
	return IntelligenceCalibrationIntegratedBehavioralLearningSafetyReview{
		CurrentState:                           "intelligence_calibration_vale_behavioral_learning_safety_ready",
		BehavioralLearningState:                IntelligenceCalibrationValEReviewPass,
		SupportedReviewStates:                  intelligenceCalibrationValEReviewStates(),
		CheckedControls:                        []string{"baseline_freshness", "baseline_observation_window", "learning_mode_timestamps", "learning_mode_critical_control_relaxation", "learning_mode_critical_alert_suppression", "learning_mode_auto_baseline_promotion", "threshold_decrease_fn_risk_review", "lower_priority_reviewer_gate", "baseline_adoption_reviewer_approval_rollback_governance"},
		Warnings:                               []string{"Behavioral and learning closure remains review-bound and non-mutating."},
		Limitations:                            []string{"Behavioral and learning closure remains projection-only and does not authorize enforcement or priority mutation."},
		EvidenceRefs:                           []string{"valb_proofs", "vald_proofs", "evidence_spine"},
		SurfaceRefs:                            []string{"/v1/intelligence/calibration/valb/proofs", "/v1/intelligence/calibration/vald/proofs", "/v1/intelligence/calibration/vale/behavioral-learning-safety"},
		ProjectionDisclaimer:                   "projection_only not_canonical_truth integrated behavioral learning safety review",
		ActiveBaselineFreshnessFresh:           true,
		StaleUnknownUnsupportedBaselineBlocked: true,
		BaselineObservationWindowBoundedAndTimestampValidated:      true,
		LearningModeTimestampsRFC3339AndBounded:                    true,
		LearningModeCriticalControlRelaxationBlocked:               true,
		LearningModeCriticalAlertSuppressionBlocked:                true,
		LearningModeAutoBaselinePromotionBlocked:                   true,
		ThresholdDecreaseRequiresFNRiskNoteAndReview:               true,
		LowerPriorityCandidateRequiresReviewerGate:                 true,
		BaselineAdoptionRequiresReviewerApprovalRollbackGovernance: true,
	}
}

func IntelligenceCalibrationValEFeedbackFederatedSafetyReviewContract() IntelligenceCalibrationIntegratedFeedbackFederatedSafetyReview {
	return IntelligenceCalibrationIntegratedFeedbackFederatedSafetyReview{
		CurrentState:                                "intelligence_calibration_vale_feedback_federated_safety_ready",
		FeedbackFederatedState:                      IntelligenceCalibrationValEReviewPass,
		SupportedReviewStates:                       intelligenceCalibrationValEReviewStates(),
		CheckedControls:                             []string{"feedback_non_mutating", "feedback_non_suppressing", "feedback_non_lowering", "false_negative_visibility", "tuning_review_bound", "suppression_candidate_expiry_scope_reviewer_rollback_reopen", "suppression_non_deleting", "suppression_false_negative_visibility", "local_change_non_mutating", "federated_non_safe", "similarity_non_override", "local_evidence_wins", "propagation_default_bounded", "raw_local_evidence_not_propagated"},
		Warnings:                                    []string{"Feedback, suppression, and federated closure remains advisory and review-bound."},
		Limitations:                                 []string{"Feedback, suppression, and federated closure remains projection-only and cannot mutate active calibration or evidence."},
		EvidenceRefs:                                []string{"valc_proofs", "evidence_spine"},
		SurfaceRefs:                                 []string{"/v1/intelligence/calibration/valc/proofs", "/v1/intelligence/calibration/vale/feedback-federated-safety"},
		ProjectionDisclaimer:                        "projection_only not_canonical_truth integrated feedback federated safety review",
		FeedbackDoesNotMutateIntelligence:           true,
		FeedbackDoesNotDirectlySuppressSignals:      true,
		FeedbackDoesNotDirectlyLowerPriority:        true,
		FalseNegativeAndMissedSeverityRemainVisible: true,
		TuningProposalsRemainAdvisoryReviewBound:    true,
		SuppressionCandidatesRequireExpiryScopeReviewerRollbackReopen: true,
		SuppressionDoesNotDeleteEvidence:                              true,
		SuppressionDoesNotHideFalseNegativePath:                       true,
		LocalCalibrationReviewDoesNotMutateActiveCalibration:          true,
		FederatedSignalsDoNotProduceLocalSafeState:                    true,
		SimilarityGatingDoesNotOverrideLocalEvidence:                  true,
		LocalEvidenceWinsOverFederatedSignal:                          true,
		PropagationDisabledOrAdvisoryOnlyByDefault:                    true,
		RawLocalEvidenceDoesNotPropagate:                              true,
	}
}

func IntelligenceCalibrationValESimulationQualityReviewContract() IntelligenceCalibrationIntegratedSimulationQualityReview {
	return IntelligenceCalibrationIntegratedSimulationQualityReview{
		CurrentState:           "intelligence_calibration_vale_simulation_quality_review_ready",
		SimulationQualityState: IntelligenceCalibrationValEReviewPass,
		SupportedReviewStates:  intelligenceCalibrationValEReviewStates(),
		CheckedControls:        []string{"simulation_harness_replayability", "expected_actual_outcomes", "low_signal_adversarial_scenarios", "false_negative_visibility", "suppression_caused_miss_analysis", "fp_fn_balance", "confidence_high_confidence_blockers", "vex_quality_blockers", "reachability_quality_guardrails", "behavioral_quality_guardrails", "federated_quality_guardrails", "coverage_non_exhaustive", "scoreboard_fp_fn_metrics", "scoreboard_non_universal_quality"},
		Warnings:               []string{"Simulation and quality closure remains scoped operational proof and not universal intelligence truth."},
		Limitations:            []string{"Simulation and quality closure remains projection-only and cannot mutate production behavior or governance."},
		EvidenceRefs:           []string{"vald_proofs", "evidence_spine"},
		SurfaceRefs:            []string{"/v1/intelligence/calibration/vald/proofs", "/v1/intelligence/calibration/vale/simulation-quality-review"},
		ProjectionDisclaimer:   "projection_only not_canonical_truth integrated simulation quality review",
		HarnessPresentReplayableOrExplicitlyLimited:                  true,
		ExpectedAndActualOutcomesRepresented:                         true,
		LowSignalAndAdversarialScenariosRepresented:                  true,
		MissedDetectionIncludesFalseNegativeVisibility:               true,
		SuppressionCausedMissAnalysisRepresented:                     true,
		FPReductionCannotHideFNRisk:                                  true,
		ConfidenceReviewBlocksUnsupportedWeakInferenceHighConfidence: true,
		VEXQualityBlocksFinalPublicationAndInsufficientNotAffected:   true,
		ReachabilityQualityPreservesPackageRuntimeGuardrails:         true,
		BehavioralQualityPreservesLearningBaselineGuardrails:         true,
		FederatedQualityPreservesLocalEvidenceBoundary:               true,
		CoverageDoesNotClaimExhaustiveDetection:                      true,
		ScoreboardIncludesFPAndFNMetrics:                             true,
		ScoreboardDoesNotClaimUniversalIntelligenceQuality:           true,
	}
}

func IntelligenceCalibrationValERegressionClosureContract() IntelligenceCalibrationIntegratedRegressionClosure {
	return IntelligenceCalibrationIntegratedRegressionClosure{
		CurrentState:                        "intelligence_calibration_vale_regression_closure_ready",
		RegressionState:                     IntelligenceCalibrationValEReviewPass,
		CoveredCategories:                   intelligenceCalibrationValERegressionCategories(),
		LimitationMessage:                   "Regression closure is bounded proof metadata and does not claim exhaustive testing.",
		Val0Coverage:                        true,
		ValACoverage:                        true,
		ValBCoverage:                        true,
		ValCCoverage:                        true,
		ValDCoverage:                        true,
		DependencyFailClosedCoverage:        true,
		CanonicalAdvisoryBoundaryCoverage:   true,
		VEXNoFinalPublicationCoverage:       true,
		LearningModeCriticalControlCoverage: true,
		SuppressionFalseNegativeCoverage:    true,
		FederatedLocalEvidenceWinsCoverage:  true,
		FPFNBalanceCoverage:                 true,
		IntegratedPassRuleCoverage:          true,
	}
}

func EvaluateIntelligenceCalibrationValEDependencyClosureState(model IntelligenceCalibrationIntegratedDependencyClosure) string {
	if strings.TrimSpace(model.CurrentState) == "" ||
		strings.TrimSpace(model.Val0State) == "" ||
		strings.TrimSpace(model.ValAState) == "" ||
		strings.TrimSpace(model.ValBState) == "" ||
		strings.TrimSpace(model.ValCState) == "" ||
		strings.TrimSpace(model.ValDState) == "" ||
		strings.TrimSpace(model.ValEState) == "" ||
		strings.TrimSpace(model.DependencyStatus) == "" ||
		strings.TrimSpace(model.ClosureGeneratedAt) == "" ||
		strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValEDependencyClosureStateIncomplete
	}
	if !containsTrimmedString(intelligenceCalibrationValEDependencyStatuses(), model.DependencyStatus) ||
		len(model.DependencyEvidenceRefs) < 5 ||
		len(model.DependencySurfaceRefs) < 7 ||
		!model.ProofStatesObserved ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValEDependencyClosureStatePartial
	}
	if _, ok := parseIntelligenceCalibrationValBTimestamp(model.ClosureGeneratedAt); !ok {
		return IntelligenceCalibrationValEDependencyClosureStatePartial
	}
	if strings.TrimSpace(model.Val0State) != IntelligenceCalibrationVal0StateActive ||
		strings.TrimSpace(model.ValAState) != IntelligenceCalibrationValAStateActive ||
		strings.TrimSpace(model.ValBState) != IntelligenceCalibrationValBStateActive ||
		strings.TrimSpace(model.ValCState) != IntelligenceCalibrationValCStateActive ||
		strings.TrimSpace(model.ValDState) != IntelligenceCalibrationValDStateActive ||
		strings.TrimSpace(model.DependencyStatus) != IntelligenceCalibrationValEDependencyPass ||
		len(model.MissingVals) > 0 ||
		len(model.InactiveVals) > 0 ||
		len(model.InconsistentVals) > 0 {
		return IntelligenceCalibrationValEDependencyClosureStatePartial
	}
	return IntelligenceCalibrationValEDependencyClosureStateActive
}

func EvaluateIntelligenceCalibrationValECoherenceReviewState(model IntelligenceCalibrationCrossValCoherenceReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.CoherenceState) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValECoherenceReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedCoherenceStates, intelligenceCalibrationValEReviewStates()...) ||
		len(model.CheckedLinks) == 0 ||
		len(model.CarriedForwardLimitations) == 0 ||
		len(model.EvidenceRefs) == 0 ||
		len(model.SurfaceRefs) == 0 ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValECoherenceReviewStatePartial
	}
	if !containsAllTrimmedStrings(model.CheckedLinks, intelligenceCalibrationValECoherenceLinks()...) ||
		len(model.MissingLinks) > 0 ||
		len(model.InconsistentLinks) > 0 ||
		!model.Val0ContractsUsedByLaterVals ||
		!model.ValAReachabilityVEXGuardrailsRespectedByValD ||
		!model.ValBBehavioralLearningGuardrailsRespectedByValD ||
		!model.ValCFeedbackSuppressionFederatedRespectedByValD ||
		!model.ValDSimulationQualityCoversRequiredPreviousSlices ||
		!model.NoPriorValClaimsPoint5Pass ||
		!model.LimitationsCarriedForward ||
		!model.AdvisoryProjectionBoundariesPreserved ||
		strings.TrimSpace(model.CoherenceState) != IntelligenceCalibrationValEReviewPass {
		return IntelligenceCalibrationValECoherenceReviewStatePartial
	}
	return IntelligenceCalibrationValECoherenceReviewStateActive
}

func EvaluateIntelligenceCalibrationValEPassRuleState(model IntelligenceCalibrationPoint5IntegratedPassRule) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.Point5State) == "" || strings.TrimSpace(model.ValEState) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValEPassRuleStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredVals, intelligenceCalibrationValERequiredVals()...) ||
		len(model.PassLimitations) == 0 ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValEPassRuleStatePartial
	}
	if strings.TrimSpace(model.Point5State) != IntelligenceCalibrationPoint5StatePass ||
		!model.PassCriteriaMet ||
		len(model.PassBlockers) > 0 ||
		len(model.MissingVals) > 0 ||
		len(model.PartialVals) > 0 ||
		len(model.UnsupportedVals) > 0 ||
		!containsExactTrimmedStringSet(model.ActiveVals, intelligenceCalibrationValERequiredVals()...) ||
		strings.TrimSpace(model.ValEState) != IntelligenceCalibrationValEStateActive {
		return IntelligenceCalibrationValEPassRuleStatePartial
	}
	return IntelligenceCalibrationValEPassRuleStateActive
}

func EvaluateIntelligenceCalibrationValEBoundaryReviewState(model IntelligenceCalibrationIntegratedAdvisoryBoundaryReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.BoundaryState) == "" || strings.TrimSpace(model.LimitationMessage) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValEBoundaryReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedBoundaryStates, intelligenceCalibrationValEReviewStates()...) ||
		len(model.CheckedSurfaces) == 0 ||
		len(model.ProjectionSurfaces) == 0 ||
		len(model.GovernanceRefs) == 0 ||
		len(model.EvidenceRefs) == 0 ||
		!strings.Contains(strings.TrimSpace(model.LimitationMessage), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.LimitationMessage), "not_canonical_truth") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValEBoundaryReviewStatePartial
	}
	if !model.EvidenceSpineRemainsCanonical ||
		!model.CalibrationOutputsRemainProjections ||
		!model.ConfidenceScoresRemainAdvisory ||
		!model.ReachabilityInferenceRemainsAdvisory ||
		!model.VEXCandidateOutputsRemainCandidates ||
		!model.BehavioralBaselinesRemainAdvisory ||
		!model.LearningModeRemainsBoundedReviewRequired ||
		!model.FeedbackRemainsAdvisory ||
		!model.SuppressionRemainsCandidateReviewBound ||
		!model.FederatedSignalsRemainAdvisory ||
		!model.SimulationMetricsRemainOperationalIndicators ||
		!model.IntegratedClosureSummaryRemainsProjectionOnly ||
		!model.NoMutationWithoutGovernance ||
		!model.FinalVEXPublicationBlocked ||
		!model.FederatedOverrideLocalEvidenceBlocked ||
		!model.LearningModeCriticalControlRelaxationBlocked ||
		len(model.ViolationSurfaces) > 0 ||
		strings.TrimSpace(model.BoundaryState) != IntelligenceCalibrationValEReviewPass {
		return IntelligenceCalibrationValEBoundaryReviewStatePartial
	}
	return IntelligenceCalibrationValEBoundaryReviewStateActive
}

func EvaluateIntelligenceCalibrationValEReachabilityVEXSafetyState(model IntelligenceCalibrationIntegratedReachabilityVEXSafetyReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReachabilityVEXState) == "" || len(model.CheckedControls) == 0 || len(model.Limitations) == 0 || len(model.EvidenceRefs) == 0 || len(model.SurfaceRefs) == 0 || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValEReachabilityVEXSafetyStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, intelligenceCalibrationValEReviewStates()...) ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValEReachabilityVEXSafetyStatePartial
	}
	if len(model.Blockers) > 0 ||
		!model.PackagePresenceOnlyBlocked ||
		!model.RuntimeLoadedOnlyBlocked ||
		!model.UnsupportedReachabilityRemainsExplicit ||
		!model.DowngradeRequiresEvidenceExplanationExpiryRollback ||
		!model.CriticalDowngradeReviewBound ||
		!model.VEXCandidateNotFinalVEX ||
		!model.FinalVEXClaimBlocked ||
		!model.PublicationAllowedBlocked ||
		!model.InsufficientEvidenceBlocksNotAffected ||
		!model.StaleUnsupportedEvidenceBlocksReviewedVEXCandidate ||
		!model.ExplanationDistinguishesNotEvidencedFromSafe ||
		strings.TrimSpace(model.ReachabilityVEXState) != IntelligenceCalibrationValEReviewPass {
		return IntelligenceCalibrationValEReachabilityVEXSafetyStatePartial
	}
	return IntelligenceCalibrationValEReachabilityVEXSafetyStateActive
}

func EvaluateIntelligenceCalibrationValEBehavioralLearningSafetyState(model IntelligenceCalibrationIntegratedBehavioralLearningSafetyReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.BehavioralLearningState) == "" || len(model.CheckedControls) == 0 || len(model.Limitations) == 0 || len(model.EvidenceRefs) == 0 || len(model.SurfaceRefs) == 0 || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValEBehavioralLearningSafetyStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, intelligenceCalibrationValEReviewStates()...) ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValEBehavioralLearningSafetyStatePartial
	}
	if len(model.Blockers) > 0 ||
		!model.ActiveBaselineFreshnessFresh ||
		!model.StaleUnknownUnsupportedBaselineBlocked ||
		!model.BaselineObservationWindowBoundedAndTimestampValidated ||
		!model.LearningModeTimestampsRFC3339AndBounded ||
		!model.LearningModeCriticalControlRelaxationBlocked ||
		!model.LearningModeCriticalAlertSuppressionBlocked ||
		!model.LearningModeAutoBaselinePromotionBlocked ||
		!model.ThresholdDecreaseRequiresFNRiskNoteAndReview ||
		!model.LowerPriorityCandidateRequiresReviewerGate ||
		!model.BaselineAdoptionRequiresReviewerApprovalRollbackGovernance ||
		strings.TrimSpace(model.BehavioralLearningState) != IntelligenceCalibrationValEReviewPass {
		return IntelligenceCalibrationValEBehavioralLearningSafetyStatePartial
	}
	return IntelligenceCalibrationValEBehavioralLearningSafetyStateActive
}

func EvaluateIntelligenceCalibrationValEFeedbackFederatedSafetyState(model IntelligenceCalibrationIntegratedFeedbackFederatedSafetyReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.FeedbackFederatedState) == "" || len(model.CheckedControls) == 0 || len(model.Limitations) == 0 || len(model.EvidenceRefs) == 0 || len(model.SurfaceRefs) == 0 || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValEFeedbackFederatedSafetyStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, intelligenceCalibrationValEReviewStates()...) ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValEFeedbackFederatedSafetyStatePartial
	}
	if len(model.Blockers) > 0 ||
		!model.FeedbackDoesNotMutateIntelligence ||
		!model.FeedbackDoesNotDirectlySuppressSignals ||
		!model.FeedbackDoesNotDirectlyLowerPriority ||
		!model.FalseNegativeAndMissedSeverityRemainVisible ||
		!model.TuningProposalsRemainAdvisoryReviewBound ||
		!model.SuppressionCandidatesRequireExpiryScopeReviewerRollbackReopen ||
		!model.SuppressionDoesNotDeleteEvidence ||
		!model.SuppressionDoesNotHideFalseNegativePath ||
		!model.LocalCalibrationReviewDoesNotMutateActiveCalibration ||
		!model.FederatedSignalsDoNotProduceLocalSafeState ||
		!model.SimilarityGatingDoesNotOverrideLocalEvidence ||
		!model.LocalEvidenceWinsOverFederatedSignal ||
		!model.PropagationDisabledOrAdvisoryOnlyByDefault ||
		!model.RawLocalEvidenceDoesNotPropagate ||
		strings.TrimSpace(model.FeedbackFederatedState) != IntelligenceCalibrationValEReviewPass {
		return IntelligenceCalibrationValEFeedbackFederatedSafetyStatePartial
	}
	return IntelligenceCalibrationValEFeedbackFederatedSafetyStateActive
}

func EvaluateIntelligenceCalibrationValESimulationQualityState(model IntelligenceCalibrationIntegratedSimulationQualityReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.SimulationQualityState) == "" || len(model.CheckedControls) == 0 || len(model.Limitations) == 0 || len(model.EvidenceRefs) == 0 || len(model.SurfaceRefs) == 0 || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return IntelligenceCalibrationValESimulationQualityStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, intelligenceCalibrationValEReviewStates()...) ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return IntelligenceCalibrationValESimulationQualityStatePartial
	}
	if len(model.Blockers) > 0 ||
		!model.HarnessPresentReplayableOrExplicitlyLimited ||
		!model.ExpectedAndActualOutcomesRepresented ||
		!model.LowSignalAndAdversarialScenariosRepresented ||
		!model.MissedDetectionIncludesFalseNegativeVisibility ||
		!model.SuppressionCausedMissAnalysisRepresented ||
		!model.FPReductionCannotHideFNRisk ||
		!model.ConfidenceReviewBlocksUnsupportedWeakInferenceHighConfidence ||
		!model.VEXQualityBlocksFinalPublicationAndInsufficientNotAffected ||
		!model.ReachabilityQualityPreservesPackageRuntimeGuardrails ||
		!model.BehavioralQualityPreservesLearningBaselineGuardrails ||
		!model.FederatedQualityPreservesLocalEvidenceBoundary ||
		!model.CoverageDoesNotClaimExhaustiveDetection ||
		!model.ScoreboardIncludesFPAndFNMetrics ||
		!model.ScoreboardDoesNotClaimUniversalIntelligenceQuality ||
		strings.TrimSpace(model.SimulationQualityState) != IntelligenceCalibrationValEReviewPass {
		return IntelligenceCalibrationValESimulationQualityStatePartial
	}
	return IntelligenceCalibrationValESimulationQualityStateActive
}

func EvaluateIntelligenceCalibrationValERegressionClosureState(model IntelligenceCalibrationIntegratedRegressionClosure) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.RegressionState) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return IntelligenceCalibrationValERegressionClosureStateIncomplete
	}
	if !containsAllTrimmedStrings(model.CoveredCategories, intelligenceCalibrationValERegressionCategories()...) ||
		len(model.CriticalMissingCategories) > 0 ||
		!model.Val0Coverage ||
		!model.ValACoverage ||
		!model.ValBCoverage ||
		!model.ValCCoverage ||
		!model.ValDCoverage ||
		!model.DependencyFailClosedCoverage ||
		!model.CanonicalAdvisoryBoundaryCoverage ||
		!model.VEXNoFinalPublicationCoverage ||
		!model.LearningModeCriticalControlCoverage ||
		!model.SuppressionFalseNegativeCoverage ||
		!model.FederatedLocalEvidenceWinsCoverage ||
		!model.FPFNBalanceCoverage ||
		!model.IntegratedPassRuleCoverage ||
		strings.TrimSpace(model.RegressionState) != IntelligenceCalibrationValEReviewPass {
		return IntelligenceCalibrationValERegressionClosureStatePartial
	}
	return IntelligenceCalibrationValERegressionClosureStateActive
}

func EvaluateIntelligenceCalibrationValEPrerequisiteState(val0State, valAState, valBState, valCState, valDState, dependencyClosureState, coherenceReviewState, boundaryReviewState, reachabilityVEXSafetyState, behavioralLearningSafetyState, feedbackFederatedSafetyState, simulationQualityState, regressionClosureState string) string {
	if strings.TrimSpace(val0State) != IntelligenceCalibrationVal0StateActive ||
		strings.TrimSpace(valAState) != IntelligenceCalibrationValAStateActive ||
		strings.TrimSpace(valBState) != IntelligenceCalibrationValBStateActive ||
		strings.TrimSpace(valCState) != IntelligenceCalibrationValCStateActive ||
		strings.TrimSpace(valDState) != IntelligenceCalibrationValDStateActive {
		return IntelligenceCalibrationValEStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		dependencyClosureState,
		coherenceReviewState,
		boundaryReviewState,
		reachabilityVEXSafetyState,
		behavioralLearningSafetyState,
		feedbackFederatedSafetyState,
		simulationQualityState,
		regressionClosureState,
	} {
		switch strings.TrimSpace(state) {
		case IntelligenceCalibrationValEDependencyClosureStateActive,
			IntelligenceCalibrationValECoherenceReviewStateActive,
			IntelligenceCalibrationValEBoundaryReviewStateActive,
			IntelligenceCalibrationValEReachabilityVEXSafetyStateActive,
			IntelligenceCalibrationValEBehavioralLearningSafetyStateActive,
			IntelligenceCalibrationValEFeedbackFederatedSafetyStateActive,
			IntelligenceCalibrationValESimulationQualityStateActive,
			IntelligenceCalibrationValERegressionClosureStateActive:
		case IntelligenceCalibrationValEDependencyClosureStatePartial,
			IntelligenceCalibrationValECoherenceReviewStatePartial,
			IntelligenceCalibrationValEBoundaryReviewStatePartial,
			IntelligenceCalibrationValEReachabilityVEXSafetyStatePartial,
			IntelligenceCalibrationValEBehavioralLearningSafetyStatePartial,
			IntelligenceCalibrationValEFeedbackFederatedSafetyStatePartial,
			IntelligenceCalibrationValESimulationQualityStatePartial,
			IntelligenceCalibrationValERegressionClosureStatePartial:
			hasPartial = true
		default:
			return IntelligenceCalibrationValEStateIncomplete
		}
	}
	if hasPartial {
		return IntelligenceCalibrationValEStateSubstantial
	}
	return IntelligenceCalibrationValEStateActive
}

func EvaluateIntelligenceCalibrationValEState(val0State, valAState, valBState, valCState, valDState, dependencyClosureState, coherenceReviewState, passRuleState, boundaryReviewState, reachabilityVEXSafetyState, behavioralLearningSafetyState, feedbackFederatedSafetyState, simulationQualityState, regressionClosureState string) string {
	baseState := EvaluateIntelligenceCalibrationValEPrerequisiteState(val0State, valAState, valBState, valCState, valDState, dependencyClosureState, coherenceReviewState, boundaryReviewState, reachabilityVEXSafetyState, behavioralLearningSafetyState, feedbackFederatedSafetyState, simulationQualityState, regressionClosureState)
	if baseState != IntelligenceCalibrationValEStateActive {
		return baseState
	}
	switch strings.TrimSpace(passRuleState) {
	case IntelligenceCalibrationValEPassRuleStateActive:
		return IntelligenceCalibrationValEStateActive
	case IntelligenceCalibrationValEPassRuleStatePartial:
		return IntelligenceCalibrationValEStateSubstantial
	default:
		return IntelligenceCalibrationValEStateIncomplete
	}
}

func EvaluateIntelligenceCalibrationValEProofsState(val0State, valAState, valBState, valCState, valDState, dependencyClosureState, coherenceReviewState, passRuleState, boundaryReviewState, reachabilityVEXSafetyState, behavioralLearningSafetyState, feedbackFederatedSafetyState, simulationQualityState, regressionClosureState, point5State string, surfaceRefs, evidenceRefs, limitations []string, projectionDisclaimer string) string {
	baseState := EvaluateIntelligenceCalibrationValEState(val0State, valAState, valBState, valCState, valDState, dependencyClosureState, coherenceReviewState, passRuleState, boundaryReviewState, reachabilityVEXSafetyState, behavioralLearningSafetyState, feedbackFederatedSafetyState, simulationQualityState, regressionClosureState)
	if len(surfaceRefs) < 15 || len(evidenceRefs) < 11 || len(limitations) == 0 || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "projection_only") || !strings.Contains(strings.TrimSpace(projectionDisclaimer), "not_canonical_truth") {
		if baseState == IntelligenceCalibrationValEStateActive {
			return IntelligenceCalibrationValEStateSubstantial
		}
		return baseState
	}
	if baseState == IntelligenceCalibrationValEStateActive && strings.TrimSpace(point5State) != IntelligenceCalibrationPoint5StatePass {
		return IntelligenceCalibrationValEStateSubstantial
	}
	return baseState
}

package operability

import "testing"

func TestIntelligenceCalibrationVal0DatasetFailsClosedWithoutVersion(t *testing.T) {
	model := IntelligenceCalibrationVal0DatasetContract()
	model.DatasetVersion = ""
	if got := EvaluateIntelligenceCalibrationVal0DatasetState(model); got == IntelligenceCalibrationVal0DatasetStateActive {
		t.Fatalf("expected non-active dataset state without version, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0DatasetFailsClosedForUnknownScenarioClass(t *testing.T) {
	model := IntelligenceCalibrationVal0DatasetContract()
	model.ScenarioClasses[0] = "unknown_scenario"
	if got := EvaluateIntelligenceCalibrationVal0DatasetState(model); got == IntelligenceCalibrationVal0DatasetStateActive {
		t.Fatalf("expected non-active dataset state for unknown scenario class, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0StaleDatasetRequiresLimitation(t *testing.T) {
	model := IntelligenceCalibrationVal0DatasetContract()
	model.FreshnessState = IntelligenceCalibrationFreshnessStale
	model.Limitations = []string{"bounded calibration sample"}
	if got := EvaluateIntelligenceCalibrationVal0DatasetState(model); got == IntelligenceCalibrationVal0DatasetStateActive {
		t.Fatalf("expected non-active dataset state when stale dataset lacks stale limitation, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0UnsupportedEvidenceCannotProduceHighConfidence(t *testing.T) {
	model := IntelligenceCalibrationVal0ConfidenceContract()
	model.EvidenceClass = IntelligenceCalibrationEvidenceUnsupported
	model.ConfidenceBand = IntelligenceCalibrationConfidenceHigh
	if got := EvaluateIntelligenceCalibrationVal0ConfidenceState(model); got == IntelligenceCalibrationVal0ConfidenceStateActive {
		t.Fatalf("expected non-active confidence state for unsupported evidence with high confidence, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0UnknownConfidenceIsNotCalibrated(t *testing.T) {
	model := IntelligenceCalibrationVal0ConfidenceContract()
	model.ConfidenceBand = IntelligenceCalibrationConfidenceUnknown
	if got := EvaluateIntelligenceCalibrationVal0ConfidenceState(model); got == IntelligenceCalibrationVal0ConfidenceStateActive {
		t.Fatalf("expected non-active confidence state for unknown confidence band, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0ExpiredLifecycleOutputIsNotActive(t *testing.T) {
	model := IntelligenceCalibrationVal0OutputLifecycleContract()
	model.ExpiredTreatedAsActive = true
	if got := EvaluateIntelligenceCalibrationVal0LifecycleState(model); got == IntelligenceCalibrationVal0LifecycleStateActive {
		t.Fatalf("expected non-active lifecycle state when expired output is treated as active, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0CandidateLifecycleOutputIsNotAccepted(t *testing.T) {
	model := IntelligenceCalibrationVal0OutputLifecycleContract()
	model.CandidateTreatedAsAccepted = true
	if got := EvaluateIntelligenceCalibrationVal0LifecycleState(model); got == IntelligenceCalibrationVal0LifecycleStateActive {
		t.Fatalf("expected non-active lifecycle state when candidate is treated as accepted, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0ReviewRequiredLifecycleOutputIsNotApproved(t *testing.T) {
	model := IntelligenceCalibrationVal0OutputLifecycleContract()
	model.ReviewRequiredTreatedApproved = true
	if got := EvaluateIntelligenceCalibrationVal0LifecycleState(model); got == IntelligenceCalibrationVal0LifecycleStateActive {
		t.Fatalf("expected non-active lifecycle state when review_required is treated as approved, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0LibraryPresenceAloneDoesNotImplyReachability(t *testing.T) {
	model := IntelligenceCalibrationVal0ReachabilityTaxonomyContract()
	model.PackagePresenceImpliesExploit = true
	if got := EvaluateIntelligenceCalibrationVal0ReachabilityState(model); got == IntelligenceCalibrationVal0ReachabilityStateActive {
		t.Fatalf("expected non-active reachability state when package presence alone implies exploitability, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0ReachabilityDowngradeRequiresEvidenceAndExplanation(t *testing.T) {
	model := IntelligenceCalibrationVal0ReachabilityTaxonomyContract()
	model.Example.DowngradeAllowed = true
	model.Example.EvidenceClass = ""
	model.Example.Explanation = ""
	if got := EvaluateIntelligenceCalibrationVal0ReachabilityState(model); got == IntelligenceCalibrationVal0ReachabilityStateActive {
		t.Fatalf("expected non-active reachability state without downgrade evidence/explanation, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0VEXCandidateIsNotFinalVEX(t *testing.T) {
	model := IntelligenceCalibrationVal0VEXCandidateContract()
	model.CandidateIsFinalVEX = true
	if got := EvaluateIntelligenceCalibrationVal0VEXState(model); got == IntelligenceCalibrationVal0VEXStateActive {
		t.Fatalf("expected non-active VEX state when candidate is treated as final VEX, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0InsufficientEvidenceCannotBecomeNotAffected(t *testing.T) {
	model := IntelligenceCalibrationVal0VEXCandidateContract()
	model.InsufficientEvidenceBecomesSafe = true
	if got := EvaluateIntelligenceCalibrationVal0VEXState(model); got == IntelligenceCalibrationVal0VEXStateActive {
		t.Fatalf("expected non-active VEX state when insufficient evidence becomes not_affected, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0FeedbackDoesNotDirectlyMutateIntelligence(t *testing.T) {
	model := IntelligenceCalibrationVal0FeedbackContract()
	model.DirectMutationAllowed = true
	if got := EvaluateIntelligenceCalibrationVal0FeedbackState(model); got == IntelligenceCalibrationVal0FeedbackStateActive {
		t.Fatalf("expected non-active feedback state when direct mutation is allowed, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0FeedbackRequiresReviewBeforeTuning(t *testing.T) {
	model := IntelligenceCalibrationVal0FeedbackContract()
	model.Items[0].ReviewRequired = false
	if got := EvaluateIntelligenceCalibrationVal0FeedbackState(model); got == IntelligenceCalibrationVal0FeedbackStateActive {
		t.Fatalf("expected non-active feedback state when review is not required, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0LearningModeCannotRelaxCriticalControls(t *testing.T) {
	model := IntelligenceCalibrationVal0LearningModeContract()
	model.CanRelaxEnforcement = true
	if got := EvaluateIntelligenceCalibrationVal0LearningModeState(model); got == IntelligenceCalibrationVal0LearningModeStateActive {
		t.Fatalf("expected non-active learning mode when critical controls can relax, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0LearningModePassesWithValidChronologicalRFC3339Timestamps(t *testing.T) {
	model := IntelligenceCalibrationVal0LearningModeContract()
	model.StartedAt = "2026-04-25T08:00:00Z"
	model.ExpiresAt = "2026-04-25T09:00:00Z"
	if got := EvaluateIntelligenceCalibrationVal0LearningModeState(model); got != IntelligenceCalibrationVal0LearningModeStateActive {
		t.Fatalf("expected active learning mode with valid chronological RFC3339 timestamps, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0LearningModeRequiresBoundedDurationAndScope(t *testing.T) {
	model := IntelligenceCalibrationVal0LearningModeContract()
	model.Scope = ""
	model.ExpiresAt = ""
	if got := EvaluateIntelligenceCalibrationVal0LearningModeState(model); got == IntelligenceCalibrationVal0LearningModeStateActive {
		t.Fatalf("expected non-active learning mode without bounded duration and scope, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0LearningModeFailsWhenExpiresAtEqualsStartedAt(t *testing.T) {
	model := IntelligenceCalibrationVal0LearningModeContract()
	model.StartedAt = "2026-04-25T08:00:00Z"
	model.ExpiresAt = "2026-04-25T08:00:00Z"
	if got := EvaluateIntelligenceCalibrationVal0LearningModeState(model); got == IntelligenceCalibrationVal0LearningModeStateActive {
		t.Fatalf("expected non-active learning mode when expires_at equals started_at, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0LearningModeFailsWhenExpiresAtIsBeforeStartedAt(t *testing.T) {
	model := IntelligenceCalibrationVal0LearningModeContract()
	model.StartedAt = "2026-04-25T09:00:00Z"
	model.ExpiresAt = "2026-04-25T08:00:00Z"
	if got := EvaluateIntelligenceCalibrationVal0LearningModeState(model); got == IntelligenceCalibrationVal0LearningModeStateActive {
		t.Fatalf("expected non-active learning mode when expires_at is before started_at, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0LearningModeFailsForMalformedStartedAt(t *testing.T) {
	model := IntelligenceCalibrationVal0LearningModeContract()
	model.StartedAt = "2026-04-25 08:00:00Z"
	if got := EvaluateIntelligenceCalibrationVal0LearningModeState(model); got == IntelligenceCalibrationVal0LearningModeStateActive {
		t.Fatalf("expected non-active learning mode for malformed started_at, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0LearningModeFailsForMalformedExpiresAt(t *testing.T) {
	model := IntelligenceCalibrationVal0LearningModeContract()
	model.ExpiresAt = "2026-04-25 09:00:00Z"
	if got := EvaluateIntelligenceCalibrationVal0LearningModeState(model); got == IntelligenceCalibrationVal0LearningModeStateActive {
		t.Fatalf("expected non-active learning mode for malformed expires_at, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0LearningModeUsesChronologicalOrderingAcrossTimezoneOffsets(t *testing.T) {
	model := IntelligenceCalibrationVal0LearningModeContract()
	model.StartedAt = "2026-04-25T10:00:00+02:00"
	model.ExpiresAt = "2026-04-25T09:30:00+01:00"
	if got := EvaluateIntelligenceCalibrationVal0LearningModeState(model); got != IntelligenceCalibrationVal0LearningModeStateActive {
		t.Fatalf("expected active learning mode when timezone-offset timestamps are chronologically ordered, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0SuppressionRequiresExpiryAndScope(t *testing.T) {
	model := IntelligenceCalibrationVal0SuppressionSafetyContract()
	model.ExpiresAt = ""
	model.SuppressionScope = ""
	if got := EvaluateIntelligenceCalibrationVal0SuppressionState(model); got == IntelligenceCalibrationVal0SuppressionStateActive {
		t.Fatalf("expected non-active suppression state without expiry and scope, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0SuppressionCannotApplyToExcludedCriticalClasses(t *testing.T) {
	model := IntelligenceCalibrationVal0SuppressionSafetyContract()
	model.SignalClass = IntelligenceCalibrationFeedbackFalseNegative
	model.ExcludedCriticalClasses = []string{IntelligenceCalibrationFeedbackFalseNegative}
	if got := EvaluateIntelligenceCalibrationVal0SuppressionState(model); got == IntelligenceCalibrationVal0SuppressionStateActive {
		t.Fatalf("expected non-active suppression state for excluded critical class, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0SuppressionReopensOnNewEvidence(t *testing.T) {
	model := IntelligenceCalibrationVal0SuppressionSafetyContract()
	model.ReopenOnNewEvidence = false
	if got := EvaluateIntelligenceCalibrationVal0SuppressionState(model); got == IntelligenceCalibrationVal0SuppressionStateActive {
		t.Fatalf("expected non-active suppression state when reopen_on_new_evidence is false, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0FederatedSignalCannotOverrideLocalEvidence(t *testing.T) {
	model := IntelligenceCalibrationVal0FederatedBoundaryContract()
	model.LocalEvidenceWins = false
	if got := EvaluateIntelligenceCalibrationVal0FederatedBoundaryState(model); got == IntelligenceCalibrationVal0FederatedBoundaryStateActive {
		t.Fatalf("expected non-active federated boundary state when local evidence does not win, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0FederatedSignalRequiresWeightingAndSimilarity(t *testing.T) {
	model := IntelligenceCalibrationVal0FederatedBoundaryContract()
	model.SourceTrustWeight = 0
	model.EnvironmentSimilarity = 0
	if got := EvaluateIntelligenceCalibrationVal0FederatedBoundaryState(model); got == IntelligenceCalibrationVal0FederatedBoundaryStateActive {
		t.Fatalf("expected non-active federated boundary state without weighting/similarity, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0ProposedCalibrationChangeIsNotApproved(t *testing.T) {
	model := IntelligenceCalibrationVal0ProvenanceContract()
	model.ApprovalState = IntelligenceCalibrationApprovalProposed
	if got := EvaluateIntelligenceCalibrationVal0ProvenanceState(model); got == IntelligenceCalibrationVal0ProvenanceStateActive {
		t.Fatalf("expected non-active provenance state for proposed change, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0ReviewRequiredCalibrationChangeDoesNotMutateActiveCalibration(t *testing.T) {
	model := IntelligenceCalibrationVal0ProvenanceContract()
	model.ApprovalState = IntelligenceCalibrationApprovalReviewRequired
	model.MutatesActiveCalibration = true
	if got := EvaluateIntelligenceCalibrationVal0ProvenanceState(model); got == IntelligenceCalibrationVal0ProvenanceStateActive {
		t.Fatalf("expected non-active provenance state for review_required mutating change, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0FreshnessHandlesStaleExpiredAndUnknownDistinctly(t *testing.T) {
	model := IntelligenceCalibrationVal0FreshnessContract()
	model.StaleTreatedAsFresh = true
	if got := EvaluateIntelligenceCalibrationVal0FreshnessState(model); got == IntelligenceCalibrationVal0FreshnessStateActive {
		t.Fatalf("expected non-active freshness state when stale is treated as fresh, got %q", got)
	}

	model = IntelligenceCalibrationVal0FreshnessContract()
	model.ExpiredTreatedAsActive = true
	if got := EvaluateIntelligenceCalibrationVal0FreshnessState(model); got == IntelligenceCalibrationVal0FreshnessStateActive {
		t.Fatalf("expected non-active freshness state when expired is treated as active, got %q", got)
	}

	model = IntelligenceCalibrationVal0FreshnessContract()
	model.UnknownFreshnessTreatedCalibrated = true
	if got := EvaluateIntelligenceCalibrationVal0FreshnessState(model); got == IntelligenceCalibrationVal0FreshnessStateActive {
		t.Fatalf("expected non-active freshness state when unknown freshness is treated as calibrated, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0MetricWithoutDefinitionOrScopeFailsClosed(t *testing.T) {
	model := IntelligenceCalibrationVal0MetricsContract()
	model.Items[0].Definition = ""
	if got := EvaluateIntelligenceCalibrationVal0MetricsState(model); got == IntelligenceCalibrationVal0MetricsStateActive {
		t.Fatalf("expected non-active metrics state without definition, got %q", got)
	}

	model = IntelligenceCalibrationVal0MetricsContract()
	model.Items[0].MeasurementScope = ""
	if got := EvaluateIntelligenceCalibrationVal0MetricsState(model); got == IntelligenceCalibrationVal0MetricsStateActive {
		t.Fatalf("expected non-active metrics state without scope, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0FalsePositiveReductionRequiresFalseNegativeDiscipline(t *testing.T) {
	model := IntelligenceCalibrationVal0FPFNContract()
	model.FalsePositiveReductionIgnoresFalseNegatives = true
	if got := EvaluateIntelligenceCalibrationVal0FPFNState(model); got == IntelligenceCalibrationVal0FPFNStateActive {
		t.Fatalf("expected non-active fp/fn discipline state when false-negative risk is ignored, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0SuppressionCausedMissCheckIsRequired(t *testing.T) {
	model := IntelligenceCalibrationVal0FPFNContract()
	model.SuppressionCausedMissCheck = false
	if got := EvaluateIntelligenceCalibrationVal0FPFNState(model); got == IntelligenceCalibrationVal0FPFNStateActive {
		t.Fatalf("expected non-active fp/fn discipline state without suppression-caused-miss check, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0RollbackRequiresRollbackOrExplicitLimitation(t *testing.T) {
	model := IntelligenceCalibrationVal0RollbackContract()
	model.RollbackAvailable = false
	model.LimitationMessage = ""
	if got := EvaluateIntelligenceCalibrationVal0RollbackState(model); got == IntelligenceCalibrationVal0RollbackStateActive {
		t.Fatalf("expected non-active rollback state without rollback or explicit limitation, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0ProofsCanBecomeActiveOnlyAsFoundationWhilePoint5RemainsNotComplete(t *testing.T) {
	datasetState := EvaluateIntelligenceCalibrationVal0DatasetState(IntelligenceCalibrationVal0DatasetContract())
	confidenceState := EvaluateIntelligenceCalibrationVal0ConfidenceState(IntelligenceCalibrationVal0ConfidenceContract())
	lifecycleState := EvaluateIntelligenceCalibrationVal0LifecycleState(IntelligenceCalibrationVal0OutputLifecycleContract())
	reachabilityState := EvaluateIntelligenceCalibrationVal0ReachabilityState(IntelligenceCalibrationVal0ReachabilityTaxonomyContract())
	vexState := EvaluateIntelligenceCalibrationVal0VEXState(IntelligenceCalibrationVal0VEXCandidateContract())
	feedbackState := EvaluateIntelligenceCalibrationVal0FeedbackState(IntelligenceCalibrationVal0FeedbackContract())
	learningModeState := EvaluateIntelligenceCalibrationVal0LearningModeState(IntelligenceCalibrationVal0LearningModeContract())
	suppressionState := EvaluateIntelligenceCalibrationVal0SuppressionState(IntelligenceCalibrationVal0SuppressionSafetyContract())
	federatedState := EvaluateIntelligenceCalibrationVal0FederatedBoundaryState(IntelligenceCalibrationVal0FederatedBoundaryContract())
	provenanceState := EvaluateIntelligenceCalibrationVal0ProvenanceState(IntelligenceCalibrationVal0ProvenanceContract())
	freshnessState := EvaluateIntelligenceCalibrationVal0FreshnessState(IntelligenceCalibrationVal0FreshnessContract())
	metricsState := EvaluateIntelligenceCalibrationVal0MetricsState(IntelligenceCalibrationVal0MetricsContract())
	fpfnState := EvaluateIntelligenceCalibrationVal0FPFNState(IntelligenceCalibrationVal0FPFNContract())
	rollbackState := EvaluateIntelligenceCalibrationVal0RollbackState(IntelligenceCalibrationVal0RollbackContract())

	val0State := EvaluateIntelligenceCalibrationVal0State(datasetState, confidenceState, lifecycleState, reachabilityState, vexState, feedbackState, learningModeState, suppressionState, federatedState, provenanceState, freshnessState, metricsState, fpfnState, rollbackState)
	if val0State != IntelligenceCalibrationVal0StateActive {
		t.Fatalf("expected active Val 0 state, got %q", val0State)
	}

	proofsState := EvaluateIntelligenceCalibrationVal0ProofsState(
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
		[]string{
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
		},
		[]string{
			"dataset_manifest",
			"confidence_reason_trace",
			"lifecycle_state_contract",
			"reachability_taxonomy_contract",
			"vex_candidate_contract",
			"feedback_review_contract",
			"learning_mode_guardrails",
			"suppression_safety_contract",
		},
		[]string{"Val 0 does not complete Točka 5."},
		[]string{"Later waves remain required."},
		"projection_only not_canonical_truth advisory_intelligence_calibration_foundation",
	)
	if proofsState != IntelligenceCalibrationVal0StateActive {
		t.Fatalf("expected active proofs state for complete foundation metadata, got %q", proofsState)
	}
	if IntelligenceCalibrationPoint5StateNotComplete != "intelligence_calibration_point_5_not_complete" {
		t.Fatalf("expected point 5 to remain not_complete in Val 0 foundation")
	}
}

func TestIntelligenceCalibrationVal0MissingRequiredComponentKeepsVal0Inactive(t *testing.T) {
	got := EvaluateIntelligenceCalibrationVal0State(
		IntelligenceCalibrationVal0DatasetStateActive,
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
		IntelligenceCalibrationVal0FPFNStateIncomplete,
		IntelligenceCalibrationVal0RollbackStateActive,
	)
	if got == IntelligenceCalibrationVal0StateActive {
		t.Fatalf("expected non-active Val 0 state when a required component is incomplete, got %q", got)
	}
}

func TestIntelligenceCalibrationVal0FoundationIsActive(t *testing.T) {
	got := EvaluateIntelligenceCalibrationVal0State(
		EvaluateIntelligenceCalibrationVal0DatasetState(IntelligenceCalibrationVal0DatasetContract()),
		EvaluateIntelligenceCalibrationVal0ConfidenceState(IntelligenceCalibrationVal0ConfidenceContract()),
		EvaluateIntelligenceCalibrationVal0LifecycleState(IntelligenceCalibrationVal0OutputLifecycleContract()),
		EvaluateIntelligenceCalibrationVal0ReachabilityState(IntelligenceCalibrationVal0ReachabilityTaxonomyContract()),
		EvaluateIntelligenceCalibrationVal0VEXState(IntelligenceCalibrationVal0VEXCandidateContract()),
		EvaluateIntelligenceCalibrationVal0FeedbackState(IntelligenceCalibrationVal0FeedbackContract()),
		EvaluateIntelligenceCalibrationVal0LearningModeState(IntelligenceCalibrationVal0LearningModeContract()),
		EvaluateIntelligenceCalibrationVal0SuppressionState(IntelligenceCalibrationVal0SuppressionSafetyContract()),
		EvaluateIntelligenceCalibrationVal0FederatedBoundaryState(IntelligenceCalibrationVal0FederatedBoundaryContract()),
		EvaluateIntelligenceCalibrationVal0ProvenanceState(IntelligenceCalibrationVal0ProvenanceContract()),
		EvaluateIntelligenceCalibrationVal0FreshnessState(IntelligenceCalibrationVal0FreshnessContract()),
		EvaluateIntelligenceCalibrationVal0MetricsState(IntelligenceCalibrationVal0MetricsContract()),
		EvaluateIntelligenceCalibrationVal0FPFNState(IntelligenceCalibrationVal0FPFNContract()),
		EvaluateIntelligenceCalibrationVal0RollbackState(IntelligenceCalibrationVal0RollbackContract()),
	)
	if got != IntelligenceCalibrationVal0StateActive {
		t.Fatalf("expected active intelligence calibration Val 0 foundation, got %q", got)
	}
}

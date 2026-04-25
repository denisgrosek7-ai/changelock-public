package operability

import "testing"

func activeIntelligenceCalibrationValBStates() (string, string, string, string, string, string, string, string, string) {
	return EvaluateIntelligenceCalibrationValBBehavioralBaselineState(IntelligenceCalibrationValBBehavioralBaselineContract()),
		EvaluateIntelligenceCalibrationValBLearningRuntimeState(IntelligenceCalibrationValBLearningModeRuntimeContract()),
		EvaluateIntelligenceCalibrationValBThresholdState(IntelligenceCalibrationValBAnomalyThresholdContract()),
		EvaluateIntelligenceCalibrationValBDriftState(IntelligenceCalibrationValBDriftSensitivityContract()),
		EvaluateIntelligenceCalibrationValBWeightingState(IntelligenceCalibrationValBCriticalityWeightingContract()),
		EvaluateIntelligenceCalibrationValBBaselineFreshnessState(IntelligenceCalibrationValBBaselineFreshnessContract()),
		EvaluateIntelligenceCalibrationValBBaselineAdoptionState(IntelligenceCalibrationValBBaselineAdoptionContract()),
		EvaluateIntelligenceCalibrationValBExplanationState(IntelligenceCalibrationValBExplanationContract()),
		EvaluateIntelligenceCalibrationValBGuardrailState(IntelligenceCalibrationValBGuardrailContract())
}

func TestIntelligenceCalibrationValBFailsClosedWithoutActiveVal0(t *testing.T) {
	baselineState, learningState, thresholdState, driftState, weightingState, freshnessState, adoptionState, explanationState, guardrailState := activeIntelligenceCalibrationValBStates()
	if got := EvaluateIntelligenceCalibrationValBState(IntelligenceCalibrationVal0StateIncomplete, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValAStateActive, baselineState, learningState, thresholdState, driftState, weightingState, freshnessState, adoptionState, explanationState, guardrailState); got == IntelligenceCalibrationValBStateActive {
		t.Fatalf("expected non-active Val B state without active Val 0 dependency, got %q", got)
	}
}

func TestIntelligenceCalibrationValBFailsClosedWithoutActiveValA(t *testing.T) {
	baselineState, learningState, thresholdState, driftState, weightingState, freshnessState, adoptionState, explanationState, guardrailState := activeIntelligenceCalibrationValBStates()
	if got := EvaluateIntelligenceCalibrationValBState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateIncomplete, IntelligenceCalibrationValAStateActive, baselineState, learningState, thresholdState, driftState, weightingState, freshnessState, adoptionState, explanationState, guardrailState); got == IntelligenceCalibrationValBStateActive {
		t.Fatalf("expected non-active Val B state without active Val A dependency, got %q", got)
	}
}

func TestIntelligenceCalibrationValBMissingBaselineVersionFailsClosed(t *testing.T) {
	model := IntelligenceCalibrationValBBehavioralBaselineContract()
	model.BaselineVersion = ""
	if got := EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model); got == IntelligenceCalibrationValBBehavioralBaselineStateActive {
		t.Fatalf("expected non-active behavioral baseline state without baseline version, got %q", got)
	}
}

func TestIntelligenceCalibrationValBObservationWindowRequiresValidRFC3339Ordering(t *testing.T) {
	model := IntelligenceCalibrationValBBehavioralBaselineContract()
	model.ObservationWindowStart = "2026-04-25T10:00:00Z"
	model.ObservationWindowEnd = "2026-04-25T09:00:00Z"
	if got := EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model); got == IntelligenceCalibrationValBBehavioralBaselineStateActive {
		t.Fatalf("expected non-active behavioral baseline state for invalid observation ordering, got %q", got)
	}
}

func TestIntelligenceCalibrationValBMalformedObservationTimestampsFailClosed(t *testing.T) {
	model := IntelligenceCalibrationValBBehavioralBaselineContract()
	model.ObservationWindowStart = "2026/04/25 10:00:00"
	if got := EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model); got == IntelligenceCalibrationValBBehavioralBaselineStateActive {
		t.Fatalf("expected non-active behavioral baseline state for malformed observation start, got %q", got)
	}
	model = IntelligenceCalibrationValBBehavioralBaselineContract()
	model.ObservationWindowEnd = "not-a-timestamp"
	if got := EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model); got == IntelligenceCalibrationValBBehavioralBaselineStateActive {
		t.Fatalf("expected non-active behavioral baseline state for malformed observation end, got %q", got)
	}
}

func TestIntelligenceCalibrationValBSampleCountMustBePositive(t *testing.T) {
	model := IntelligenceCalibrationValBBehavioralBaselineContract()
	model.SampleCount = 0
	if got := EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model); got == IntelligenceCalibrationValBBehavioralBaselineStateActive {
		t.Fatalf("expected non-active behavioral baseline state with non-positive sample count, got %q", got)
	}
}

func TestIntelligenceCalibrationValBUnsupportedBaselineSignalRemainsExplicit(t *testing.T) {
	model := IntelligenceCalibrationValBBehavioralBaselineContract()
	model.ObservedSignalClasses = model.ObservedSignalClasses[:len(model.ObservedSignalClasses)-1]
	if got := EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model); got == IntelligenceCalibrationValBBehavioralBaselineStateActive {
		t.Fatalf("expected non-active behavioral baseline state when unsupported signal disappears, got %q", got)
	}
}

func TestIntelligenceCalibrationValBStaleBaselineRequiresLimitation(t *testing.T) {
	model := IntelligenceCalibrationValBBehavioralBaselineContract()
	model.FreshnessState = IntelligenceCalibrationFreshnessStale
	model.LimitationMessage = "bounded baseline review"
	if got := EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model); got == IntelligenceCalibrationValBBehavioralBaselineStateActive {
		t.Fatalf("expected non-active behavioral baseline state without explicit stale limitation, got %q", got)
	}
}

func TestIntelligenceCalibrationValBBaselineCannotRelaxEnforcementOrSuppressAlerts(t *testing.T) {
	model := IntelligenceCalibrationValBBehavioralBaselineContract()
	model.CanRelaxEnforcement = true
	if got := EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model); got == IntelligenceCalibrationValBBehavioralBaselineStateActive {
		t.Fatalf("expected non-active behavioral baseline state when enforcement can relax, got %q", got)
	}
	model = IntelligenceCalibrationValBBehavioralBaselineContract()
	model.CanSuppressAlerts = true
	if got := EvaluateIntelligenceCalibrationValBBehavioralBaselineState(model); got == IntelligenceCalibrationValBBehavioralBaselineStateActive {
		t.Fatalf("expected non-active behavioral baseline state when alerts can be suppressed, got %q", got)
	}
}

func TestIntelligenceCalibrationValBLearningModeRuntimeTimestampsAreParsedChronologically(t *testing.T) {
	model := IntelligenceCalibrationValBLearningModeRuntimeContract()
	model.StartedAt = "2026-04-25T10:00:00+02:00"
	model.ExpiresAt = "2026-04-25T09:30:00+01:00"
	if got := EvaluateIntelligenceCalibrationValBLearningRuntimeState(model); got != IntelligenceCalibrationValBLearningRuntimeStateActive {
		t.Fatalf("expected active learning runtime state for chronologically valid timezone-offset timestamps, got %q", got)
	}
}

func TestIntelligenceCalibrationValBLearningModeCannotRelaxCriticalControls(t *testing.T) {
	model := IntelligenceCalibrationValBLearningModeRuntimeContract()
	model.CanRelaxEnforcement = true
	if got := EvaluateIntelligenceCalibrationValBLearningRuntimeState(model); got == IntelligenceCalibrationValBLearningRuntimeStateActive {
		t.Fatalf("expected non-active learning runtime state when critical controls can relax, got %q", got)
	}
}

func TestIntelligenceCalibrationValBLearningModeCannotSuppressCriticalAlerts(t *testing.T) {
	model := IntelligenceCalibrationValBLearningModeRuntimeContract()
	model.CanSuppressCriticalAlerts = true
	if got := EvaluateIntelligenceCalibrationValBLearningRuntimeState(model); got == IntelligenceCalibrationValBLearningRuntimeStateActive {
		t.Fatalf("expected non-active learning runtime state when critical alerts can be suppressed, got %q", got)
	}
}

func TestIntelligenceCalibrationValBLearningModeCannotAutoPromoteBaseline(t *testing.T) {
	model := IntelligenceCalibrationValBLearningModeRuntimeContract()
	model.CanAutoPromoteBaseline = true
	if got := EvaluateIntelligenceCalibrationValBLearningRuntimeState(model); got == IntelligenceCalibrationValBLearningRuntimeStateActive {
		t.Fatalf("expected non-active learning runtime state when baseline can auto-promote, got %q", got)
	}
}

func TestIntelligenceCalibrationValBExpiredLearningModeIsNotActive(t *testing.T) {
	model := IntelligenceCalibrationValBLearningModeRuntimeContract()
	model.LearningModeState = IntelligenceCalibrationLearningExpired
	if got := EvaluateIntelligenceCalibrationValBLearningRuntimeState(model); got == IntelligenceCalibrationValBLearningRuntimeStateActive {
		t.Fatalf("expected non-active expired learning runtime state, got %q", got)
	}
}

func TestIntelligenceCalibrationValBThresholdDecreaseSensitivityForCriticalClassRequiresReview(t *testing.T) {
	model := IntelligenceCalibrationValBAnomalyThresholdContract()
	model.ThresholdChangeDirection = IntelligenceCalibrationValBThresholdDecreaseSensitivity
	model.AppliesToCriticalClass = true
	model.ReviewRequired = false
	if got := EvaluateIntelligenceCalibrationValBThresholdState(model); got == IntelligenceCalibrationValBThresholdStateActive {
		t.Fatalf("expected non-active threshold state when critical decrease sensitivity lacks review, got %q", got)
	}
}

func TestIntelligenceCalibrationValBThresholdCalibrationRequiresBothFPRiskAndFNRiskNotes(t *testing.T) {
	model := IntelligenceCalibrationValBAnomalyThresholdContract()
	model.FalsePositiveRiskNote = ""
	if got := EvaluateIntelligenceCalibrationValBThresholdState(model); got == IntelligenceCalibrationValBThresholdStateActive {
		t.Fatalf("expected non-active threshold state without FP risk note, got %q", got)
	}
	model = IntelligenceCalibrationValBAnomalyThresholdContract()
	model.FalseNegativeRiskNote = ""
	if got := EvaluateIntelligenceCalibrationValBThresholdState(model); got == IntelligenceCalibrationValBThresholdStateActive {
		t.Fatalf("expected non-active threshold state without FN risk note, got %q", got)
	}
}

func TestIntelligenceCalibrationValBUnsupportedEvidenceCannotChangeThresholdWithoutReview(t *testing.T) {
	model := IntelligenceCalibrationValBAnomalyThresholdContract()
	model.EvidenceClass = IntelligenceCalibrationEvidenceUnsupported
	model.ThresholdChangeDirection = IntelligenceCalibrationValBThresholdDecreaseSensitivity
	if got := EvaluateIntelligenceCalibrationValBThresholdState(model); got == IntelligenceCalibrationValBThresholdStateActive {
		t.Fatalf("expected non-active threshold state when unsupported evidence changes threshold without review, got %q", got)
	}
}

func TestIntelligenceCalibrationValBThresholdChangeDoesNotMutateActiveDetectionState(t *testing.T) {
	model := IntelligenceCalibrationValBAnomalyThresholdContract()
	model.MutatesActiveDetection = true
	if got := EvaluateIntelligenceCalibrationValBThresholdState(model); got == IntelligenceCalibrationValBThresholdStateActive {
		t.Fatalf("expected non-active threshold state when active detection mutates, got %q", got)
	}
}

func TestIntelligenceCalibrationValBUnknownDriftScoreCannotDriveAutomaticSensitivityAdjustment(t *testing.T) {
	model := IntelligenceCalibrationValBDriftSensitivityContract()
	model.DriftScoreBand = IntelligenceCalibrationConfidenceUnknown
	model.SensitivityAdjustment = IntelligenceCalibrationValBDriftAdjustmentIncrease
	if got := EvaluateIntelligenceCalibrationValBDriftState(model); got == IntelligenceCalibrationValBDriftStateActive {
		t.Fatalf("expected non-active drift state when unknown drift drives automatic adjustment, got %q", got)
	}
}

func TestIntelligenceCalibrationValBDecreaseSensitivityForCriticalHighContextRequiresReview(t *testing.T) {
	model := IntelligenceCalibrationValBDriftSensitivityContract()
	model.CriticalityContext = IntelligenceCalibrationValBCriticalityCritical
	model.SensitivityAdjustment = IntelligenceCalibrationValBDriftAdjustmentDecrease
	model.ReviewRequired = false
	if got := EvaluateIntelligenceCalibrationValBDriftState(model); got == IntelligenceCalibrationValBDriftStateActive {
		t.Fatalf("expected non-active drift state when critical decrease sensitivity lacks review, got %q", got)
	}
}

func TestIntelligenceCalibrationValBStaleDriftInputsRequireLimitationOrReview(t *testing.T) {
	model := IntelligenceCalibrationValBDriftSensitivityContract()
	model.FreshnessState = IntelligenceCalibrationFreshnessStale
	model.LimitationMessage = "behavioral drift window bounded"
	model.ReviewRequired = false
	if got := EvaluateIntelligenceCalibrationValBDriftState(model); got == IntelligenceCalibrationValBDriftStateActive {
		t.Fatalf("expected non-active drift state without explicit stale limitation or review, got %q", got)
	}
}

func TestIntelligenceCalibrationValBUnknownCriticalityCannotProduceLowerPriorityCandidate(t *testing.T) {
	model := IntelligenceCalibrationValBCriticalityWeightingContract()
	model.CriticalityClass = IntelligenceCalibrationValBCriticalityUnknown
	model.WeightingAction = IntelligenceCalibrationValBWeightingLowerPriorityCandidate
	if got := EvaluateIntelligenceCalibrationValBWeightingState(model); got == IntelligenceCalibrationValBWeightingStateActive {
		t.Fatalf("expected non-active weighting state when unknown criticality lowers priority, got %q", got)
	}
}

func TestIntelligenceCalibrationValBCriticalHighAssetsCannotBeAutoLoweredWithoutReview(t *testing.T) {
	model := IntelligenceCalibrationValBCriticalityWeightingContract()
	model.CriticalityClass = IntelligenceCalibrationValBCriticalityHigh
	model.WeightingAction = IntelligenceCalibrationValBWeightingLowerPriorityCandidate
	model.ReviewerRequired = false
	if got := EvaluateIntelligenceCalibrationValBWeightingState(model); got == IntelligenceCalibrationValBWeightingStateActive {
		t.Fatalf("expected non-active weighting state when high criticality lowers priority without review, got %q", got)
	}
}

func TestIntelligenceCalibrationValBRaisePriorityRequiresReasonEvidence(t *testing.T) {
	model := IntelligenceCalibrationValBCriticalityWeightingContract()
	model.WeightingAction = IntelligenceCalibrationValBWeightingRaisePriority
	model.ReasonCode = ""
	model.EvidenceRefs = nil
	if got := EvaluateIntelligenceCalibrationValBWeightingState(model); got == IntelligenceCalibrationValBWeightingStateActive {
		t.Fatalf("expected non-active weighting state when raise priority lacks reason/evidence, got %q", got)
	}
}

func TestIntelligenceCalibrationValBExpiredBaselineFreshnessIsNotActive(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineFreshnessContract()
	model.FreshnessState = IntelligenceCalibrationFreshnessExpired
	if got := EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model); got == IntelligenceCalibrationValBBaselineFreshnessStateActive {
		t.Fatalf("expected non-active expired baseline freshness state, got %q", got)
	}
}

func TestIntelligenceCalibrationValBBaselineFreshnessFailsClosedWhenLastObservedAtIsMissing(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineFreshnessContract()
	model.LastObservedAt = ""
	if got := EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model); got == IntelligenceCalibrationValBBaselineFreshnessStateActive {
		t.Fatalf("expected non-active baseline freshness state when last_observed_at is missing, got %q", got)
	}
}

func TestIntelligenceCalibrationValBBaselineFreshnessFailsClosedWhenExpiresAtIsMissing(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineFreshnessContract()
	model.ExpiresAt = ""
	if got := EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model); got == IntelligenceCalibrationValBBaselineFreshnessStateActive {
		t.Fatalf("expected non-active baseline freshness state when expires_at is missing, got %q", got)
	}
}

func TestIntelligenceCalibrationValBBaselineFreshnessFailsClosedWhenLastObservedAtIsMalformed(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineFreshnessContract()
	model.LastObservedAt = "2026/04/25 00:00:00"
	if got := EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model); got == IntelligenceCalibrationValBBaselineFreshnessStateActive {
		t.Fatalf("expected non-active baseline freshness state when last_observed_at is malformed, got %q", got)
	}
}

func TestIntelligenceCalibrationValBBaselineFreshnessFailsClosedWhenExpiresAtIsMalformed(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineFreshnessContract()
	model.ExpiresAt = "not-a-timestamp"
	if got := EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model); got == IntelligenceCalibrationValBBaselineFreshnessStateActive {
		t.Fatalf("expected non-active baseline freshness state when expires_at is malformed, got %q", got)
	}
}

func TestIntelligenceCalibrationValBBaselineFreshnessFailsClosedWhenExpiresAtIsNotAfterLastObservedAt(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineFreshnessContract()
	model.LastObservedAt = "2026-04-25T09:00:00Z"
	model.ExpiresAt = "2026-04-25T09:00:00Z"
	if got := EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model); got == IntelligenceCalibrationValBBaselineFreshnessStateActive {
		t.Fatalf("expected non-active baseline freshness state when expires_at equals last_observed_at, got %q", got)
	}

	model = IntelligenceCalibrationValBBaselineFreshnessContract()
	model.LastObservedAt = "2026-04-25T10:00:00Z"
	model.ExpiresAt = "2026-04-25T09:00:00Z"
	if got := EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model); got == IntelligenceCalibrationValBBaselineFreshnessStateActive {
		t.Fatalf("expected non-active baseline freshness state when expires_at is before last_observed_at, got %q", got)
	}
}

func TestIntelligenceCalibrationValBBaselineFreshnessPassesWithValidRFC3339Timestamps(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineFreshnessContract()
	model.LastObservedAt = "2026-04-25T08:00:00Z"
	model.ExpiresAt = "2026-04-25T09:00:00Z"
	if got := EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model); got != IntelligenceCalibrationValBBaselineFreshnessStateActive {
		t.Fatalf("expected active baseline freshness state with valid RFC3339 timestamps, got %q", got)
	}
}

func TestIntelligenceCalibrationValBBaselineFreshnessUsesChronologicalTimezoneOrdering(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineFreshnessContract()
	model.LastObservedAt = "2026-04-25T10:00:00+02:00"
	model.ExpiresAt = "2026-04-25T09:30:00+01:00"
	if got := EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model); got != IntelligenceCalibrationValBBaselineFreshnessStateActive {
		t.Fatalf("expected active baseline freshness state for chronologically valid timezone-offset timestamps, got %q", got)
	}
}

func TestIntelligenceCalibrationValBUnknownFreshnessRequiresLimitation(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineFreshnessContract()
	model.FreshnessState = IntelligenceCalibrationFreshnessUnknown
	model.LimitationMessage = "bounded freshness review"
	if got := EvaluateIntelligenceCalibrationValBBaselineFreshnessState(model); got == IntelligenceCalibrationValBBaselineFreshnessStateActive {
		t.Fatalf("expected non-active baseline freshness state when unknown freshness lacks explicit unknown limitation, got %q", got)
	}
}

func TestIntelligenceCalibrationValBProposedBaselineAdoptionIsNotStableApproval(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineAdoptionContract()
	model.ProposedAdoptionState = IntelligenceCalibrationApprovalProposed
	if got := EvaluateIntelligenceCalibrationValBBaselineAdoptionState(model); got == IntelligenceCalibrationValBBaselineAdoptionStateActive {
		t.Fatalf("expected non-active adoption state for proposed adoption, got %q", got)
	}
}

func TestIntelligenceCalibrationValBReviewRequiredAdoptionIsNotApproval(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineAdoptionContract()
	model.ProposedAdoptionState = IntelligenceCalibrationApprovalReviewRequired
	if got := EvaluateIntelligenceCalibrationValBBaselineAdoptionState(model); got == IntelligenceCalibrationValBBaselineAdoptionStateActive {
		t.Fatalf("expected non-active adoption state for review-required adoption, got %q", got)
	}
}

func TestIntelligenceCalibrationValBApprovedAdoptionRequiresRollbackRef(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineAdoptionContract()
	model.RollbackRef = ""
	if got := EvaluateIntelligenceCalibrationValBBaselineAdoptionState(model); got == IntelligenceCalibrationValBBaselineAdoptionStateActive {
		t.Fatalf("expected non-active adoption state without rollback ref, got %q", got)
	}
}

func TestIntelligenceCalibrationValBApprovedAdoptionRequiresReviewerGate(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineAdoptionContract()
	model.ReviewerRequired = false
	if got := EvaluateIntelligenceCalibrationValBBaselineAdoptionState(model); got == IntelligenceCalibrationValBBaselineAdoptionStateActive {
		t.Fatalf("expected non-active adoption state without reviewer gate, got %q", got)
	}
}

func TestIntelligenceCalibrationValBApprovedAdoptionRequiresApprovalRef(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineAdoptionContract()
	model.ApprovalRef = ""
	if got := EvaluateIntelligenceCalibrationValBBaselineAdoptionState(model); got == IntelligenceCalibrationValBBaselineAdoptionStateActive {
		t.Fatalf("expected non-active adoption state without approval ref, got %q", got)
	}
}

func TestIntelligenceCalibrationValBApprovedAdoptionPassesWithReviewerApprovalRollbackAndGovernance(t *testing.T) {
	model := IntelligenceCalibrationValBBaselineAdoptionContract()
	if got := EvaluateIntelligenceCalibrationValBBaselineAdoptionState(model); got != IntelligenceCalibrationValBBaselineAdoptionStateActive {
		t.Fatalf("expected active adoption state with reviewer gate, approval ref, rollback ref, and governance requirement, got %q", got)
	}
}

func TestIntelligenceCalibrationValBExplanationIncludesFPRiskAndFNRiskNotes(t *testing.T) {
	model := IntelligenceCalibrationValBExplanationContract()
	model.FalsePositiveRiskNote = ""
	if got := EvaluateIntelligenceCalibrationValBExplanationState(model); got == IntelligenceCalibrationValBExplanationStateActive {
		t.Fatalf("expected non-active explanation state without FP risk note, got %q", got)
	}
	model = IntelligenceCalibrationValBExplanationContract()
	model.FalseNegativeRiskNote = ""
	if got := EvaluateIntelligenceCalibrationValBExplanationState(model); got == IntelligenceCalibrationValBExplanationStateActive {
		t.Fatalf("expected non-active explanation state without FN risk note, got %q", got)
	}
}

func TestIntelligenceCalibrationValBRedactedExplanationDoesNotConvertReviewRequiredIntoApproved(t *testing.T) {
	model := IntelligenceCalibrationValBExplanationContract()
	model.ReviewRequiredPresentedApproved = true
	if got := EvaluateIntelligenceCalibrationValBExplanationState(model); got == IntelligenceCalibrationValBExplanationStateActive {
		t.Fatalf("expected non-active explanation state when redaction converts review_required into approved, got %q", got)
	}
}

func TestIntelligenceCalibrationValBSafetyGuardrailsBlockAutoSuppressionPromotionPriorityAndEnforcementMutation(t *testing.T) {
	model := IntelligenceCalibrationValBGuardrailContract()
	model.AutoSuppressionBlocked = false
	if got := EvaluateIntelligenceCalibrationValBGuardrailState(model); got == IntelligenceCalibrationValBGuardrailStateActive {
		t.Fatalf("expected non-active guardrail state when auto suppression is not blocked, got %q", got)
	}
	model = IntelligenceCalibrationValBGuardrailContract()
	model.CriticalControlRelaxationBlocked = false
	if got := EvaluateIntelligenceCalibrationValBGuardrailState(model); got == IntelligenceCalibrationValBGuardrailStateActive {
		t.Fatalf("expected non-active guardrail state when critical control relaxation is not blocked, got %q", got)
	}
	model = IntelligenceCalibrationValBGuardrailContract()
	model.AutoBaselinePromotionBlocked = false
	if got := EvaluateIntelligenceCalibrationValBGuardrailState(model); got == IntelligenceCalibrationValBGuardrailStateActive {
		t.Fatalf("expected non-active guardrail state when auto baseline promotion is not blocked, got %q", got)
	}
	model = IntelligenceCalibrationValBGuardrailContract()
	model.PriorityMutationBlocked = false
	if got := EvaluateIntelligenceCalibrationValBGuardrailState(model); got == IntelligenceCalibrationValBGuardrailStateActive {
		t.Fatalf("expected non-active guardrail state when priority mutation is not blocked, got %q", got)
	}
	model = IntelligenceCalibrationValBGuardrailContract()
	model.EnforcementMutationBlocked = false
	if got := EvaluateIntelligenceCalibrationValBGuardrailState(model); got == IntelligenceCalibrationValBGuardrailStateActive {
		t.Fatalf("expected non-active guardrail state when enforcement mutation is not blocked, got %q", got)
	}
}

func TestIntelligenceCalibrationValBProofsCanBecomeActiveOnlyAsBehavioralBaselineAndLearningModeWhilePoint5RemainsNotComplete(t *testing.T) {
	baselineState, learningState, thresholdState, driftState, weightingState, freshnessState, adoptionState, explanationState, guardrailState := activeIntelligenceCalibrationValBStates()
	if got := EvaluateIntelligenceCalibrationValBState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValAStateActive, baselineState, learningState, thresholdState, driftState, weightingState, freshnessState, adoptionState, explanationState, guardrailState); got != IntelligenceCalibrationValBStateActive {
		t.Fatalf("expected active Val B state for valid bounded behavioral calibration, got %q", got)
	}
	proofsState := EvaluateIntelligenceCalibrationValBProofsState(
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationValAStateActive,
		IntelligenceCalibrationValAStateActive,
		baselineState,
		learningState,
		thresholdState,
		driftState,
		weightingState,
		freshnessState,
		adoptionState,
		explanationState,
		guardrailState,
		[]string{
			"/v1/intelligence/calibration/valb/behavioral-baseline",
			"/v1/intelligence/calibration/valb/learning-mode-runtime",
			"/v1/intelligence/calibration/valb/anomaly-thresholds",
			"/v1/intelligence/calibration/valb/drift-sensitivity",
			"/v1/intelligence/calibration/valb/criticality-weighting",
			"/v1/intelligence/calibration/valb/baseline-freshness",
			"/v1/intelligence/calibration/valb/baseline-adoption",
			"/v1/intelligence/calibration/valb/explanations",
			"/v1/intelligence/calibration/valb/safety-guardrails",
			"/v1/intelligence/calibration/valb/proofs",
		},
		[]string{"val0_proofs", "vala_proofs", "behavioral_baseline_profile", "learning_mode_runtime_discipline", "anomaly_threshold_calibration", "drift_sensitivity_scaling", "criticality_weighting", "baseline_freshness", "baseline_adoption", "behavioral_explanation", "behavioral_guardrail"},
		[]string{"Val B remains advisory and does not complete Točka 5."},
		[]string{"Later waves still need feedback/federated tuning, simulation, final gate, and integrated closure."},
		"projection_only not_canonical_truth advisory_behavioral_baseline_learning_mode",
	)
	if proofsState != IntelligenceCalibrationValBStateActive {
		t.Fatalf("expected active Val B proofs state for valid behavioral calibration slice, got %q", proofsState)
	}
	if IntelligenceCalibrationPoint5StateNotComplete != "intelligence_calibration_point_5_not_complete" {
		t.Fatalf("expected point 5 to remain not complete in Val B")
	}
}

func TestIntelligenceCalibrationValBMissingRequiredComponentKeepsValBInactive(t *testing.T) {
	_, learningState, thresholdState, driftState, weightingState, freshnessState, adoptionState, explanationState, guardrailState := activeIntelligenceCalibrationValBStates()
	if got := EvaluateIntelligenceCalibrationValBState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBBehavioralBaselineStateIncomplete, learningState, thresholdState, driftState, weightingState, freshnessState, adoptionState, explanationState, guardrailState); got == IntelligenceCalibrationValBStateActive {
		t.Fatalf("expected non-active Val B state when a required component is incomplete, got %q", got)
	}
}

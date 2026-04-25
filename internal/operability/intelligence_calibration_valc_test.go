package operability

import "testing"

func activeIntelligenceCalibrationValCStates() (string, string, string, string, string, string, string, string, string, string, string) {
	return EvaluateIntelligenceCalibrationValCFeedbackIntakeState(IntelligenceCalibrationValCStructuredFeedbackIntakeContract()),
		EvaluateIntelligenceCalibrationValCReviewCockpitState(IntelligenceCalibrationValCFeedbackReviewCockpitContract()),
		EvaluateIntelligenceCalibrationValCTuningProposalState(IntelligenceCalibrationValCTuningProposalContract()),
		EvaluateIntelligenceCalibrationValCSuppressionSafetyState(IntelligenceCalibrationValCSuppressionSafetyContract()),
		EvaluateIntelligenceCalibrationValCSuppressionRollbackState(IntelligenceCalibrationValCSuppressionRollbackContract()),
		EvaluateIntelligenceCalibrationValCLocalChangeReviewState(IntelligenceCalibrationValCLocalChangeReviewContract()),
		EvaluateIntelligenceCalibrationValCFederatedWeightingState(IntelligenceCalibrationValCFederatedSignalWeightingContract()),
		EvaluateIntelligenceCalibrationValCSimilarityGatingState(IntelligenceCalibrationValCSimilarityGatingContract()),
		EvaluateIntelligenceCalibrationValCLocalOverrideState(IntelligenceCalibrationValCLocalOverrideDisciplineContract()),
		EvaluateIntelligenceCalibrationValCPropagationPolicyState(IntelligenceCalibrationValCPropagationPolicyContract()),
		EvaluateIntelligenceCalibrationValCExplanationState(IntelligenceCalibrationValCExplanationContract())
}

func TestIntelligenceCalibrationValCFailsClosedWithoutActiveVal0(t *testing.T) {
	feedbackState, reviewState, proposalState, suppressionState, rollbackState, localReviewState, federatedState, similarityState, overrideState, propagationState, explanationState := activeIntelligenceCalibrationValCStates()
	if got := EvaluateIntelligenceCalibrationValCState(IntelligenceCalibrationVal0StateIncomplete, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValBStateActive, feedbackState, reviewState, proposalState, suppressionState, rollbackState, localReviewState, federatedState, similarityState, overrideState, propagationState, explanationState); got == IntelligenceCalibrationValCStateActive {
		t.Fatalf("expected non-active Val C state without active Val 0 dependency, got %q", got)
	}
}

func TestIntelligenceCalibrationValCFailsClosedWithoutActiveValA(t *testing.T) {
	feedbackState, reviewState, proposalState, suppressionState, rollbackState, localReviewState, federatedState, similarityState, overrideState, propagationState, explanationState := activeIntelligenceCalibrationValCStates()
	if got := EvaluateIntelligenceCalibrationValCState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateIncomplete, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValBStateActive, feedbackState, reviewState, proposalState, suppressionState, rollbackState, localReviewState, federatedState, similarityState, overrideState, propagationState, explanationState); got == IntelligenceCalibrationValCStateActive {
		t.Fatalf("expected non-active Val C state without active Val A dependency, got %q", got)
	}
}

func TestIntelligenceCalibrationValCFailsClosedWithoutActiveValB(t *testing.T) {
	feedbackState, reviewState, proposalState, suppressionState, rollbackState, localReviewState, federatedState, similarityState, overrideState, propagationState, explanationState := activeIntelligenceCalibrationValCStates()
	if got := EvaluateIntelligenceCalibrationValCState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateIncomplete, IntelligenceCalibrationValBStateActive, feedbackState, reviewState, proposalState, suppressionState, rollbackState, localReviewState, federatedState, similarityState, overrideState, propagationState, explanationState); got == IntelligenceCalibrationValCStateActive {
		t.Fatalf("expected non-active Val C state without active Val B dependency, got %q", got)
	}
}

func TestIntelligenceCalibrationValCFeedbackDoesNotMutateSuppressOrLowerPriority(t *testing.T) {
	model := IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	model.MutatesIntelligence = true
	if got := EvaluateIntelligenceCalibrationValCFeedbackIntakeState(model); got == IntelligenceCalibrationValCFeedbackIntakeStateActive {
		t.Fatalf("expected non-active feedback intake state when feedback mutates intelligence, got %q", got)
	}

	model = IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	model.SuppressesSignals = true
	if got := EvaluateIntelligenceCalibrationValCFeedbackIntakeState(model); got == IntelligenceCalibrationValCFeedbackIntakeStateActive {
		t.Fatalf("expected non-active feedback intake state when feedback suppresses signals, got %q", got)
	}

	model = IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	model.LowersPriority = true
	if got := EvaluateIntelligenceCalibrationValCFeedbackIntakeState(model); got == IntelligenceCalibrationValCFeedbackIntakeStateActive {
		t.Fatalf("expected non-active feedback intake state when feedback lowers priority, got %q", got)
	}
}

func TestIntelligenceCalibrationValCFalseNegativeOrMissedSeverityFeedbackIsNotNoiseReduction(t *testing.T) {
	model := IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	model.FeedbackClass = IntelligenceCalibrationFeedbackFalseNegative
	model.RoutedAsNoiseReduction = true
	if got := EvaluateIntelligenceCalibrationValCFeedbackIntakeState(model); got == IntelligenceCalibrationValCFeedbackIntakeStateActive {
		t.Fatalf("expected non-active feedback intake state when false-negative feedback is routed as noise reduction, got %q", got)
	}

	model = IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	model.FeedbackClass = IntelligenceCalibrationFeedbackMissedSeverity
	model.RoutedAsNoiseReduction = true
	if got := EvaluateIntelligenceCalibrationValCFeedbackIntakeState(model); got == IntelligenceCalibrationValCFeedbackIntakeStateActive {
		t.Fatalf("expected non-active feedback intake state when missed-severity feedback is routed as noise reduction, got %q", got)
	}
}

func TestIntelligenceCalibrationValCFeedbackMissingActorSignalOrReasonFailsClosed(t *testing.T) {
	model := IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	model.ActorRef = ""
	if got := EvaluateIntelligenceCalibrationValCFeedbackIntakeState(model); got == IntelligenceCalibrationValCFeedbackIntakeStateActive {
		t.Fatalf("expected non-active feedback intake state without actor ref, got %q", got)
	}

	model = IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	model.SignalRef = ""
	if got := EvaluateIntelligenceCalibrationValCFeedbackIntakeState(model); got == IntelligenceCalibrationValCFeedbackIntakeStateActive {
		t.Fatalf("expected non-active feedback intake state without signal ref, got %q", got)
	}

	model = IntelligenceCalibrationValCStructuredFeedbackIntakeContract()
	model.ReasonCode = ""
	if got := EvaluateIntelligenceCalibrationValCFeedbackIntakeState(model); got == IntelligenceCalibrationValCFeedbackIntakeStateActive {
		t.Fatalf("expected non-active feedback intake state without reason code, got %q", got)
	}
}

func TestIntelligenceCalibrationValCReviewCockpitDoesNotTreatReviewedAsAppliedAndKeepsRiskVisible(t *testing.T) {
	model := IntelligenceCalibrationValCFeedbackReviewCockpitContract()
	model.ReviewedTreatedAsApplied = true
	if got := EvaluateIntelligenceCalibrationValCReviewCockpitState(model); got == IntelligenceCalibrationValCReviewCockpitStateActive {
		t.Fatalf("expected non-active review cockpit state when reviewed is treated as applied tuning, got %q", got)
	}

	model = IntelligenceCalibrationValCFeedbackReviewCockpitContract()
	model.FalseNegativeVisible = false
	if got := EvaluateIntelligenceCalibrationValCReviewCockpitState(model); got == IntelligenceCalibrationValCReviewCockpitStateActive {
		t.Fatalf("expected non-active review cockpit state when false-negative feedback is not visible, got %q", got)
	}

	model = IntelligenceCalibrationValCFeedbackReviewCockpitContract()
	model.EscalationRequired = false
	if got := EvaluateIntelligenceCalibrationValCReviewCockpitState(model); got == IntelligenceCalibrationValCReviewCockpitStateActive {
		t.Fatalf("expected non-active review cockpit state when high-risk feedback lacks escalation, got %q", got)
	}
}

func TestIntelligenceCalibrationValCTuningProposalRemainsAdvisoryAndReviewBound(t *testing.T) {
	model := IntelligenceCalibrationValCTuningProposalContract()
	model.MutatesActiveCalibration = true
	if got := EvaluateIntelligenceCalibrationValCTuningProposalState(model); got == IntelligenceCalibrationValCTuningProposalStateActive {
		t.Fatalf("expected non-active tuning proposal state when proposal mutates active calibration, got %q", got)
	}

	model = IntelligenceCalibrationValCTuningProposalContract()
	model.ApprovalState = IntelligenceCalibrationApprovalProposed
	if got := EvaluateIntelligenceCalibrationValCTuningProposalState(model); got == IntelligenceCalibrationValCTuningProposalStateActive {
		t.Fatalf("expected non-active tuning proposal state for proposed proposal, got %q", got)
	}

	model = IntelligenceCalibrationValCTuningProposalContract()
	model.ApprovalState = IntelligenceCalibrationApprovalReviewRequired
	if got := EvaluateIntelligenceCalibrationValCTuningProposalState(model); got == IntelligenceCalibrationValCTuningProposalStateActive {
		t.Fatalf("expected non-active tuning proposal state for review-required proposal, got %q", got)
	}

	model = IntelligenceCalibrationValCTuningProposalContract()
	model.FalseNegativeRiskNote = ""
	if got := EvaluateIntelligenceCalibrationValCTuningProposalState(model); got == IntelligenceCalibrationValCTuningProposalStateActive {
		t.Fatalf("expected non-active tuning proposal state when sensitivity decrease lacks FN risk note, got %q", got)
	}

	model = IntelligenceCalibrationValCTuningProposalContract()
	model.ProposedChangeType = IntelligenceCalibrationValCProposalSuppressionCandidate
	model.ReviewerRequired = false
	if got := EvaluateIntelligenceCalibrationValCTuningProposalState(model); got == IntelligenceCalibrationValCTuningProposalStateActive {
		t.Fatalf("expected non-active tuning proposal state when suppression candidate lacks reviewer gate, got %q", got)
	}

	model = IntelligenceCalibrationValCTuningProposalContract()
	model.RollbackRef = ""
	if got := EvaluateIntelligenceCalibrationValCTuningProposalState(model); got == IntelligenceCalibrationValCTuningProposalStateActive {
		t.Fatalf("expected non-active tuning proposal state when approved proposal lacks rollback ref, got %q", got)
	}
}

func TestIntelligenceCalibrationValCSuppressionCandidateRequiresExpiryScopeAndReviewer(t *testing.T) {
	model := IntelligenceCalibrationValCSuppressionSafetyContract()
	model.ExpiresAt = ""
	if got := EvaluateIntelligenceCalibrationValCSuppressionSafetyState(model); got == IntelligenceCalibrationValCSuppressionSafetyStateActive {
		t.Fatalf("expected non-active suppression safety state without expiry, got %q", got)
	}

	model = IntelligenceCalibrationValCSuppressionSafetyContract()
	model.SuppressionScope = ""
	if got := EvaluateIntelligenceCalibrationValCSuppressionSafetyState(model); got == IntelligenceCalibrationValCSuppressionSafetyStateActive {
		t.Fatalf("expected non-active suppression safety state without scope, got %q", got)
	}

	model = IntelligenceCalibrationValCSuppressionSafetyContract()
	model.ReviewerRef = ""
	if got := EvaluateIntelligenceCalibrationValCSuppressionSafetyState(model); got == IntelligenceCalibrationValCSuppressionSafetyStateActive {
		t.Fatalf("expected non-active suppression safety state without reviewer ref, got %q", got)
	}
}

func TestIntelligenceCalibrationValCSuppressionCandidatePreservesEvidenceAndFalseNegativePath(t *testing.T) {
	model := IntelligenceCalibrationValCSuppressionSafetyContract()
	model.DeletesEvidence = true
	if got := EvaluateIntelligenceCalibrationValCSuppressionSafetyState(model); got == IntelligenceCalibrationValCSuppressionSafetyStateActive {
		t.Fatalf("expected non-active suppression safety state when evidence is deleted, got %q", got)
	}

	model = IntelligenceCalibrationValCSuppressionSafetyContract()
	model.SuppressesCriticalClass = true
	if got := EvaluateIntelligenceCalibrationValCSuppressionSafetyState(model); got == IntelligenceCalibrationValCSuppressionSafetyStateActive {
		t.Fatalf("expected non-active suppression safety state when critical class is suppressed, got %q", got)
	}

	model = IntelligenceCalibrationValCSuppressionSafetyContract()
	model.StrongerEvidenceReopens = false
	if got := EvaluateIntelligenceCalibrationValCSuppressionSafetyState(model); got == IntelligenceCalibrationValCSuppressionSafetyStateActive {
		t.Fatalf("expected non-active suppression safety state when stronger evidence does not reopen candidate, got %q", got)
	}

	model = IntelligenceCalibrationValCSuppressionSafetyContract()
	model.SuppressesFalseNegativePath = true
	if got := EvaluateIntelligenceCalibrationValCSuppressionSafetyState(model); got == IntelligenceCalibrationValCSuppressionSafetyStateActive {
		t.Fatalf("expected non-active suppression safety state when false-negative path is hidden, got %q", got)
	}
}

func TestIntelligenceCalibrationValCSuppressionRollbackRequiresTriggersAndSafetyCheck(t *testing.T) {
	model := IntelligenceCalibrationValCSuppressionRollbackContract()
	model.RollbackTriggerConditions = []string{"new_evidence"}
	if got := EvaluateIntelligenceCalibrationValCSuppressionRollbackState(model); got == IntelligenceCalibrationValCSuppressionRollbackStateActive {
		t.Fatalf("expected non-active suppression rollback state without false-negative discovery trigger, got %q", got)
	}

	model = IntelligenceCalibrationValCSuppressionRollbackContract()
	model.RollbackSafetyCheck = ""
	if got := EvaluateIntelligenceCalibrationValCSuppressionRollbackState(model); got == IntelligenceCalibrationValCSuppressionRollbackStateActive {
		t.Fatalf("expected non-active suppression rollback state without safety check, got %q", got)
	}
}

func TestIntelligenceCalibrationValCLocalChangeReviewRequiresExplicitLocalScopeAndNoMutation(t *testing.T) {
	model := IntelligenceCalibrationValCLocalChangeReviewContract()
	model.MutatesActiveCalibration = true
	if got := EvaluateIntelligenceCalibrationValCLocalChangeReviewState(model); got == IntelligenceCalibrationValCLocalChangeReviewStateActive {
		t.Fatalf("expected non-active local change review state when it mutates active calibration, got %q", got)
	}

	model = IntelligenceCalibrationValCLocalChangeReviewContract()
	model.LocalScope = ""
	if got := EvaluateIntelligenceCalibrationValCLocalChangeReviewState(model); got == IntelligenceCalibrationValCLocalChangeReviewStateActive {
		t.Fatalf("expected non-active local change review state without explicit local scope, got %q", got)
	}

	model = IntelligenceCalibrationValCLocalChangeReviewContract()
	model.RollbackRef = ""
	if got := EvaluateIntelligenceCalibrationValCLocalChangeReviewState(model); got == IntelligenceCalibrationValCLocalChangeReviewStateActive {
		t.Fatalf("expected non-active local change review state without rollback ref on approved change, got %q", got)
	}
}

func TestIntelligenceCalibrationValCFederatedWeightingCapsConfidenceAndCannotDeclareLocalSafe(t *testing.T) {
	model := IntelligenceCalibrationValCFederatedSignalWeightingContract()
	model.SourceQualityState = IntelligenceCalibrationValCFederatedSourceQualityUnknown
	model.ConfidenceCap = IntelligenceCalibrationConfidenceHigh
	if got := EvaluateIntelligenceCalibrationValCFederatedWeightingState(model); got == IntelligenceCalibrationValCFederatedWeightingStateActive {
		t.Fatalf("expected non-active federated weighting state when unknown source quality does not cap confidence, got %q", got)
	}

	model = IntelligenceCalibrationValCFederatedSignalWeightingContract()
	model.ProducesLocalSafeState = true
	if got := EvaluateIntelligenceCalibrationValCFederatedWeightingState(model); got == IntelligenceCalibrationValCFederatedWeightingStateActive {
		t.Fatalf("expected non-active federated weighting state when federated signal declares local safe state, got %q", got)
	}

	model = IntelligenceCalibrationValCFederatedSignalWeightingContract()
	model.FreshnessState = IntelligenceCalibrationFreshnessStale
	model.LimitationMessage = "bounded federated hint"
	if got := EvaluateIntelligenceCalibrationValCFederatedWeightingState(model); got == IntelligenceCalibrationValCFederatedWeightingStateActive {
		t.Fatalf("expected non-active federated weighting state when stale freshness lacks limitation, got %q", got)
	}
}

func TestIntelligenceCalibrationValCSimilarityGatingRequiresCriticalDimensionsAndNoConfidenceInflation(t *testing.T) {
	model := IntelligenceCalibrationValCSimilarityGatingContract()
	model.MissingDimensions = []string{"runtime"}
	if got := EvaluateIntelligenceCalibrationValCSimilarityGatingState(model); got == IntelligenceCalibrationValCSimilarityGatingStateActive {
		t.Fatalf("expected non-active similarity gating state when allow_advisory_use lacks critical runtime dimension, got %q", got)
	}

	model = IntelligenceCalibrationValCSimilarityGatingContract()
	model.SimilarityBand = IntelligenceCalibrationValCSimilarityBandUnknown
	model.GatingDecision = IntelligenceCalibrationValCSimilarityDecisionAllowAdvisoryUse
	if got := EvaluateIntelligenceCalibrationValCSimilarityGatingState(model); got == IntelligenceCalibrationValCSimilarityGatingStateActive {
		t.Fatalf("expected non-active similarity gating state when unknown similarity still allows advisory use, got %q", got)
	}

	model = IntelligenceCalibrationValCSimilarityGatingContract()
	model.SimilarityBand = IntelligenceCalibrationValCSimilarityBandLow
	model.IncreasesConfidence = true
	if got := EvaluateIntelligenceCalibrationValCSimilarityGatingState(model); got == IntelligenceCalibrationValCSimilarityGatingStateActive {
		t.Fatalf("expected non-active similarity gating state when low similarity increases confidence, got %q", got)
	}

	model = IntelligenceCalibrationValCSimilarityGatingContract()
	model.OverridesLocalEvidence = true
	if got := EvaluateIntelligenceCalibrationValCSimilarityGatingState(model); got == IntelligenceCalibrationValCSimilarityGatingStateActive {
		t.Fatalf("expected non-active similarity gating state when it overrides local evidence, got %q", got)
	}
}

func TestIntelligenceCalibrationValCLocalOverrideRequiresReasonAndAuditWhileLocalEvidenceWins(t *testing.T) {
	model := IntelligenceCalibrationValCLocalOverrideDisciplineContract()
	model.LocalEvidenceWins = false
	if got := EvaluateIntelligenceCalibrationValCLocalOverrideState(model); got == IntelligenceCalibrationValCLocalOverrideStateActive {
		t.Fatalf("expected non-active local override state when local evidence does not win, got %q", got)
	}

	model = IntelligenceCalibrationValCLocalOverrideDisciplineContract()
	model.OverrideReasonRequired = false
	if got := EvaluateIntelligenceCalibrationValCLocalOverrideState(model); got == IntelligenceCalibrationValCLocalOverrideStateActive {
		t.Fatalf("expected non-active local override state without reason requirement, got %q", got)
	}

	model = IntelligenceCalibrationValCLocalOverrideDisciplineContract()
	model.OverrideAuditRequired = false
	if got := EvaluateIntelligenceCalibrationValCLocalOverrideState(model); got == IntelligenceCalibrationValCLocalOverrideStateActive {
		t.Fatalf("expected non-active local override state without audit requirement, got %q", got)
	}
}

func TestIntelligenceCalibrationValCPropagationPolicyRemainsDisabledOrAdvisoryOnly(t *testing.T) {
	model := IntelligenceCalibrationValCPropagationPolicyContract()
	model.PropagationAllowed = true
	if got := EvaluateIntelligenceCalibrationValCPropagationPolicyState(model); got == IntelligenceCalibrationValCPropagationPolicyStateActive {
		t.Fatalf("expected non-active propagation policy state when propagation is enabled, got %q", got)
	}

	model = IntelligenceCalibrationValCPropagationPolicyContract()
	model.MutatesRemoteCalibration = true
	if got := EvaluateIntelligenceCalibrationValCPropagationPolicyState(model); got == IntelligenceCalibrationValCPropagationPolicyStateActive {
		t.Fatalf("expected non-active propagation policy state when remote calibration mutates, got %q", got)
	}

	model = IntelligenceCalibrationValCPropagationPolicyContract()
	model.RequiresRedaction = false
	if got := EvaluateIntelligenceCalibrationValCPropagationPolicyState(model); got == IntelligenceCalibrationValCPropagationPolicyStateActive {
		t.Fatalf("expected non-active propagation policy state when redaction is not required, got %q", got)
	}

	model = IntelligenceCalibrationValCPropagationPolicyContract()
	model.RequiresReview = false
	if got := EvaluateIntelligenceCalibrationValCPropagationPolicyState(model); got == IntelligenceCalibrationValCPropagationPolicyStateActive {
		t.Fatalf("expected non-active propagation policy state when review is not required, got %q", got)
	}

	model = IntelligenceCalibrationValCPropagationPolicyContract()
	model.PropagatesRawLocalEvidence = true
	if got := EvaluateIntelligenceCalibrationValCPropagationPolicyState(model); got == IntelligenceCalibrationValCPropagationPolicyStateActive {
		t.Fatalf("expected non-active propagation policy state when raw local evidence propagates, got %q", got)
	}
}

func TestIntelligenceCalibrationValCExplanationIncludesRisksAndOverridePropagationNotes(t *testing.T) {
	model := IntelligenceCalibrationValCExplanationContract()
	model.FalsePositiveRiskNote = ""
	if got := EvaluateIntelligenceCalibrationValCExplanationState(model); got == IntelligenceCalibrationValCExplanationStateActive {
		t.Fatalf("expected non-active explanation state without FP risk note, got %q", got)
	}

	model = IntelligenceCalibrationValCExplanationContract()
	model.FalseNegativeRiskNote = ""
	if got := EvaluateIntelligenceCalibrationValCExplanationState(model); got == IntelligenceCalibrationValCExplanationStateActive {
		t.Fatalf("expected non-active explanation state without FN risk note, got %q", got)
	}

	model = IntelligenceCalibrationValCExplanationContract()
	model.LocalOverrideNote = ""
	if got := EvaluateIntelligenceCalibrationValCExplanationState(model); got == IntelligenceCalibrationValCExplanationStateActive {
		t.Fatalf("expected non-active explanation state without local override note, got %q", got)
	}

	model = IntelligenceCalibrationValCExplanationContract()
	model.PropagationNote = ""
	if got := EvaluateIntelligenceCalibrationValCExplanationState(model); got == IntelligenceCalibrationValCExplanationStateActive {
		t.Fatalf("expected non-active explanation state without propagation note, got %q", got)
	}
}

func TestIntelligenceCalibrationValCRedactedExplanationDoesNotConvertReviewRequiredIntoApproved(t *testing.T) {
	model := IntelligenceCalibrationValCExplanationContract()
	model.ReviewRequiredPresentedApproved = true
	if got := EvaluateIntelligenceCalibrationValCExplanationState(model); got == IntelligenceCalibrationValCExplanationStateActive {
		t.Fatalf("expected non-active explanation state when review_required is presented as approved, got %q", got)
	}
}

func TestIntelligenceCalibrationValCProofsCanBecomeActiveOnlyAsFeedbackSuppressionAndFederatedTuningWhilePoint5RemainsNotComplete(t *testing.T) {
	feedbackState, reviewState, proposalState, suppressionState, rollbackState, localReviewState, federatedState, similarityState, overrideState, propagationState, explanationState := activeIntelligenceCalibrationValCStates()
	got := EvaluateIntelligenceCalibrationValCProofsState(
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationValAStateActive,
		IntelligenceCalibrationValAStateActive,
		IntelligenceCalibrationValBStateActive,
		IntelligenceCalibrationValBStateActive,
		feedbackState,
		reviewState,
		proposalState,
		suppressionState,
		rollbackState,
		localReviewState,
		federatedState,
		similarityState,
		overrideState,
		propagationState,
		explanationState,
		[]string{
			"/v1/intelligence/calibration/valc/feedback-intake",
			"/v1/intelligence/calibration/valc/feedback-review",
			"/v1/intelligence/calibration/valc/tuning-proposals",
			"/v1/intelligence/calibration/valc/suppression-safety",
			"/v1/intelligence/calibration/valc/suppression-rollback",
			"/v1/intelligence/calibration/valc/local-change-review",
			"/v1/intelligence/calibration/valc/federated-weighting",
			"/v1/intelligence/calibration/valc/similarity-gating",
			"/v1/intelligence/calibration/valc/local-override",
			"/v1/intelligence/calibration/valc/propagation-policy",
			"/v1/intelligence/calibration/valc/explanations",
			"/v1/intelligence/calibration/valc/proofs",
		},
		[]string{
			"val0_proofs",
			"vala_proofs",
			"valb_proofs",
			"structured_feedback_intake",
			"feedback_review_cockpit",
			"tuning_proposal_generation",
			"suppression_safety_application",
			"suppression_rollback_path",
			"local_calibration_change_review",
			"federated_signal_weighting",
			"environment_similarity_gating",
			"local_override_discipline",
			"bounded_propagation_policy",
			"feedback_federated_explanation",
			"evidence_spine",
		},
		[]string{"Val C remains advisory only."},
		[]string{"Točka 5 remains not complete."},
		"projection_only not_canonical_truth advisory_feedback_suppression_federated_tuning",
	)
	if got != IntelligenceCalibrationValCStateActive {
		t.Fatalf("expected active Val C proofs state, got %q", got)
	}
}

func TestIntelligenceCalibrationValCMissingRequiredComponentKeepsValCInactive(t *testing.T) {
	_, reviewState, proposalState, suppressionState, rollbackState, localReviewState, federatedState, similarityState, overrideState, propagationState, explanationState := activeIntelligenceCalibrationValCStates()
	if got := EvaluateIntelligenceCalibrationValCState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValCFeedbackIntakeStateIncomplete, reviewState, proposalState, suppressionState, rollbackState, localReviewState, federatedState, similarityState, overrideState, propagationState, explanationState); got == IntelligenceCalibrationValCStateActive {
		t.Fatalf("expected non-active Val C state with missing feedback intake component, got %q", got)
	}
}

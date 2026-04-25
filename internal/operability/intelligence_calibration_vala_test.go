package operability

import "testing"

func activeIntelligenceCalibrationValAStates() (string, string, string, string, string, string, string, string, string) {
	return EvaluateIntelligenceCalibrationValAAggregationState(IntelligenceCalibrationValAReachabilityAggregationContract()),
		EvaluateIntelligenceCalibrationValAExploitabilityState(IntelligenceCalibrationValAExploitabilityCalibrationContract()),
		EvaluateIntelligenceCalibrationValADecisionState(IntelligenceCalibrationValADowngradeEscalationContract()),
		EvaluateIntelligenceCalibrationValACAVIState(IntelligenceCalibrationValACAVITuningContract()),
		EvaluateIntelligenceCalibrationValAVEXCandidateState(IntelligenceCalibrationValAVEXCandidateContract()),
		EvaluateIntelligenceCalibrationValAVEXSufficiencyState(IntelligenceCalibrationValAVEXSufficiencyContract()),
		EvaluateIntelligenceCalibrationValAExplanationState(IntelligenceCalibrationValAExplanationContract()),
		EvaluateIntelligenceCalibrationValAConfidenceOutcomeState(IntelligenceCalibrationValAConfidenceOutcomeContract()),
		EvaluateIntelligenceCalibrationValAPublicationGuardrailState(IntelligenceCalibrationValAPublicationGuardrailContract())
}

func TestIntelligenceCalibrationValAFailsClosedWithoutActiveVal0(t *testing.T) {
	aggregationState, exploitabilityState, decisionState, caviState, vexState, sufficiencyState, explanationState, outcomeState, guardrailState := activeIntelligenceCalibrationValAStates()
	if got := EvaluateIntelligenceCalibrationValAState(IntelligenceCalibrationVal0StateIncomplete, IntelligenceCalibrationVal0StateActive, aggregationState, exploitabilityState, decisionState, caviState, vexState, sufficiencyState, explanationState, outcomeState, guardrailState); got == IntelligenceCalibrationValAStateActive {
		t.Fatalf("expected non-active Val A state without active Val 0 dependency, got %q", got)
	}
}

func TestIntelligenceCalibrationValAPackagePresenceAloneDoesNotImplyExploitability(t *testing.T) {
	model := IntelligenceCalibrationValAReachabilityAggregationContract()
	model.PackagePresenceImpliesExploit = true
	if got := EvaluateIntelligenceCalibrationValAAggregationState(model); got == IntelligenceCalibrationValAAggregationStateActive {
		t.Fatalf("expected non-active aggregation state when package presence alone implies exploitability, got %q", got)
	}
}

func TestIntelligenceCalibrationValARuntimeLoadedAloneDoesNotImplyExecution(t *testing.T) {
	model := IntelligenceCalibrationValAReachabilityAggregationContract()
	model.RuntimeLoadedImpliesExecution = true
	if got := EvaluateIntelligenceCalibrationValAAggregationState(model); got == IntelligenceCalibrationValAAggregationStateActive {
		t.Fatalf("expected non-active aggregation state when runtime_loaded alone implies vulnerable execution, got %q", got)
	}
}

func TestIntelligenceCalibrationValAUnsupportedSignalIsExplicit(t *testing.T) {
	model := IntelligenceCalibrationValAReachabilityAggregationContract()
	model.SignalClasses = model.SignalClasses[:len(model.SignalClasses)-1]
	if got := EvaluateIntelligenceCalibrationValAAggregationState(model); got == IntelligenceCalibrationValAAggregationStateActive {
		t.Fatalf("expected non-active aggregation state when unsupported signal disappears, got %q", got)
	}
}

func TestIntelligenceCalibrationValAPartialAggregationIsNotComplete(t *testing.T) {
	model := IntelligenceCalibrationValAReachabilityAggregationContract()
	model.AggregationState = IntelligenceCalibrationValAAggregationPartial
	if got := EvaluateIntelligenceCalibrationValAAggregationState(model); got == IntelligenceCalibrationValAAggregationStateActive {
		t.Fatalf("expected non-active aggregation state for partial aggregation, got %q", got)
	}
}

func TestIntelligenceCalibrationValAStaleAggregationRequiresLimitation(t *testing.T) {
	model := IntelligenceCalibrationValAReachabilityAggregationContract()
	model.FreshnessState = IntelligenceCalibrationFreshnessStale
	model.LimitationMessage = "bounded aggregation window"
	if got := EvaluateIntelligenceCalibrationValAAggregationState(model); got == IntelligenceCalibrationValAAggregationStateActive {
		t.Fatalf("expected non-active aggregation state when stale aggregation lacks explicit stale limitation, got %q", got)
	}
}

func TestIntelligenceCalibrationValANotEvidencedDoesNotMeanSafe(t *testing.T) {
	model := IntelligenceCalibrationValAExploitabilityCalibrationContract()
	model.CurrentlyNotEvidencedIsSafe = true
	if got := EvaluateIntelligenceCalibrationValAExploitabilityState(model); got == IntelligenceCalibrationValAExploitabilityStateActive {
		t.Fatalf("expected non-active exploitability state when not evidenced is treated as safe, got %q", got)
	}
}

func TestIntelligenceCalibrationValALowEvidenceCannotBecomeNotAffected(t *testing.T) {
	model := IntelligenceCalibrationValAExploitabilityCalibrationContract()
	model.LowEvidenceBecomesSafe = true
	if got := EvaluateIntelligenceCalibrationValAExploitabilityState(model); got == IntelligenceCalibrationValAExploitabilityStateActive {
		t.Fatalf("expected non-active exploitability state when low evidence becomes not affected, got %q", got)
	}
}

func TestIntelligenceCalibrationValAUnsupportedExploitabilityIsNotLowRisk(t *testing.T) {
	model := IntelligenceCalibrationValAExploitabilityCalibrationContract()
	model.UnsupportedBecomesLowRisk = true
	if got := EvaluateIntelligenceCalibrationValAExploitabilityState(model); got == IntelligenceCalibrationValAExploitabilityStateActive {
		t.Fatalf("expected non-active exploitability state when unsupported becomes low risk, got %q", got)
	}
}

func TestIntelligenceCalibrationValADowngradeRequiresEvidenceReasonAndExplanation(t *testing.T) {
	model := IntelligenceCalibrationValADowngradeEscalationContract()
	model.ReasonCodes = nil
	model.Explanation = ""
	if got := EvaluateIntelligenceCalibrationValADecisionState(model); got == IntelligenceCalibrationValADecisionStateActive {
		t.Fatalf("expected non-active downgrade decision without reason code and explanation, got %q", got)
	}
}

func TestIntelligenceCalibrationValADowngradeWithUnsupportedEvidenceFailsClosed(t *testing.T) {
	model := IntelligenceCalibrationValADowngradeEscalationContract()
	model.EvidenceClass = IntelligenceCalibrationEvidenceUnsupported
	if got := EvaluateIntelligenceCalibrationValADecisionState(model); got == IntelligenceCalibrationValADecisionStateActive {
		t.Fatalf("expected non-active downgrade decision with unsupported evidence, got %q", got)
	}
}

func TestIntelligenceCalibrationValADowngradeCannotApplyToExcludedCriticalClasses(t *testing.T) {
	model := IntelligenceCalibrationValADowngradeEscalationContract()
	model.AppliesToExcludedCritical = true
	if got := EvaluateIntelligenceCalibrationValADecisionState(model); got == IntelligenceCalibrationValADecisionStateActive {
		t.Fatalf("expected non-active downgrade decision for excluded critical class, got %q", got)
	}
}

func TestIntelligenceCalibrationValADowngradeRequiresExpiryOrExplicitLimitation(t *testing.T) {
	model := IntelligenceCalibrationValADowngradeEscalationContract()
	model.ExpiresAt = ""
	model.LimitationMessage = ""
	if got := EvaluateIntelligenceCalibrationValADecisionState(model); got == IntelligenceCalibrationValADecisionStateActive {
		t.Fatalf("expected non-active downgrade decision without expiry or explicit limitation, got %q", got)
	}
}

func TestIntelligenceCalibrationValAEscalationRequiresReasonAndEvidence(t *testing.T) {
	model := IntelligenceCalibrationValADowngradeEscalationContract()
	model.DecisionType = IntelligenceCalibrationValADecisionEscalation
	model.ReasonCodes = nil
	model.EvidenceRefs = nil
	if got := EvaluateIntelligenceCalibrationValADecisionState(model); got == IntelligenceCalibrationValADecisionStateActive {
		t.Fatalf("expected non-active escalation decision without reason or evidence, got %q", got)
	}
}

func TestIntelligenceCalibrationValAPresentOnlyCannotProduceDowngradeCandidate(t *testing.T) {
	model := IntelligenceCalibrationValACAVITuningContract()
	model.ExecutionContextState = IntelligenceCalibrationValAExecutionPresentOnly
	model.TuningRecommendation = IntelligenceCalibrationValATuningDowngradeCandidate
	if got := EvaluateIntelligenceCalibrationValACAVIState(model); got == IntelligenceCalibrationValACAVIStateActive {
		t.Fatalf("expected non-active CAVI state when present_only produces downgrade candidate, got %q", got)
	}
}

func TestIntelligenceCalibrationValAUnsupportedCallPathCannotProduceHighConfidence(t *testing.T) {
	model := IntelligenceCalibrationValACAVITuningContract()
	model.CallPathEvidenceState = IntelligenceCalibrationEvidenceUnsupported
	model.ConfidenceBand = IntelligenceCalibrationConfidenceHigh
	if got := EvaluateIntelligenceCalibrationValACAVIState(model); got == IntelligenceCalibrationValACAVIStateActive {
		t.Fatalf("expected non-active CAVI state when unsupported call path produces high confidence, got %q", got)
	}
}

func TestIntelligenceCalibrationValAMissingPackageToFunctionLinkageLimitsConfidence(t *testing.T) {
	model := IntelligenceCalibrationValACAVITuningContract()
	model.PackageToFunctionLinkage = false
	model.ConfidenceBand = IntelligenceCalibrationConfidenceHigh
	if got := EvaluateIntelligenceCalibrationValACAVIState(model); got == IntelligenceCalibrationValACAVIStateActive {
		t.Fatalf("expected non-active CAVI state without package-to-function linkage at high confidence, got %q", got)
	}
}

func TestIntelligenceCalibrationValAUnknownExploitPreconditionsRequireReview(t *testing.T) {
	model := IntelligenceCalibrationValACAVITuningContract()
	model.ExploitPreconditionsKnown = false
	model.TuningRecommendation = IntelligenceCalibrationValATuningKeepPriority
	if got := EvaluateIntelligenceCalibrationValACAVIState(model); got == IntelligenceCalibrationValACAVIStateActive {
		t.Fatalf("expected non-active CAVI state when exploit preconditions are unknown without review recommendation, got %q", got)
	}
}

func TestIntelligenceCalibrationValAVEXCandidateIsNotFinalVEX(t *testing.T) {
	model := IntelligenceCalibrationValAVEXCandidateContract()
	if model.FinalVEXClaim {
		t.Fatalf("expected default Val A VEX candidate to remain non-final")
	}
	if got := EvaluateIntelligenceCalibrationValAVEXCandidateState(model); got != IntelligenceCalibrationValAVEXCandidateStateActive {
		t.Fatalf("expected active candidate state for bounded non-final VEX candidate, got %q", got)
	}
}

func TestIntelligenceCalibrationValAFinalVEXClaimBlocksActiveState(t *testing.T) {
	model := IntelligenceCalibrationValAVEXCandidateContract()
	model.FinalVEXClaim = true
	if got := EvaluateIntelligenceCalibrationValAVEXCandidateState(model); got == IntelligenceCalibrationValAVEXCandidateStateActive {
		t.Fatalf("expected non-active VEX candidate state when final_vex_claim is true, got %q", got)
	}
}

func TestIntelligenceCalibrationValAPublicationAllowedIsFalseByDefault(t *testing.T) {
	model := IntelligenceCalibrationValAVEXCandidateContract()
	if model.PublicationAllowed {
		t.Fatalf("expected publication_allowed to be false by default in Val A")
	}
}

func TestIntelligenceCalibrationValAInsufficientEvidenceCannotProduceNotAffectedCandidate(t *testing.T) {
	model := IntelligenceCalibrationValAVEXCandidateContract()
	model.SuggestedVEXOutcome = IntelligenceCalibrationVEXOutcomeNotAffectedCandidate
	model.EvidenceSufficiencyState = IntelligenceCalibrationValAVEXSufficiencyInsufficient
	if got := EvaluateIntelligenceCalibrationValAVEXCandidateState(model); got == IntelligenceCalibrationValAVEXCandidateStateActive {
		t.Fatalf("expected non-active VEX candidate when insufficient evidence produces not_affected_candidate, got %q", got)
	}
}

func TestIntelligenceCalibrationValAStaleOrUnsupportedEvidenceCannotProduceReviewedCandidate(t *testing.T) {
	model := IntelligenceCalibrationValAVEXCandidateContract()
	model.CandidateState = IntelligenceCalibrationVEXStateReviewed
	model.EvidenceSufficiencyState = IntelligenceCalibrationValAVEXSufficiencyStale
	if got := EvaluateIntelligenceCalibrationValAVEXCandidateState(model); got == IntelligenceCalibrationValAVEXCandidateStateActive {
		t.Fatalf("expected non-active VEX candidate when stale evidence produces reviewed candidate, got %q", got)
	}
	model = IntelligenceCalibrationValAVEXCandidateContract()
	model.CandidateState = IntelligenceCalibrationVEXStateReviewed
	model.EvidenceSufficiencyState = IntelligenceCalibrationValAVEXSufficiencyUnsupported
	if got := EvaluateIntelligenceCalibrationValAVEXCandidateState(model); got == IntelligenceCalibrationValAVEXCandidateStateActive {
		t.Fatalf("expected non-active VEX candidate when unsupported evidence produces reviewed candidate, got %q", got)
	}
}

func TestIntelligenceCalibrationValAExpiredVEXCandidateIsNotActive(t *testing.T) {
	model := IntelligenceCalibrationValAVEXCandidateContract()
	model.CandidateState = IntelligenceCalibrationVEXStateExpired
	if got := EvaluateIntelligenceCalibrationValAVEXCandidateState(model); got == IntelligenceCalibrationValAVEXCandidateStateActive {
		t.Fatalf("expected non-active expired VEX candidate, got %q", got)
	}
}

func TestIntelligenceCalibrationValAMissingRequiredVEXEvidenceBlocksSufficiency(t *testing.T) {
	model := IntelligenceCalibrationValAVEXSufficiencyContract()
	model.PresentEvidenceClasses = []string{IntelligenceCalibrationEvidenceDirectlyEvidenced}
	if got := EvaluateIntelligenceCalibrationValAVEXSufficiencyState(model); got == IntelligenceCalibrationValAVEXSufficiencyStateActive {
		t.Fatalf("expected non-active sufficiency state when required evidence is missing, got %q", got)
	}
}

func TestIntelligenceCalibrationValAStaleOrUnsupportedEvidenceCannotSatisfySufficiency(t *testing.T) {
	model := IntelligenceCalibrationValAVEXSufficiencyContract()
	model.StaleEvidenceRefs = []string{"evidence:stale"}
	if got := EvaluateIntelligenceCalibrationValAVEXSufficiencyState(model); got == IntelligenceCalibrationValAVEXSufficiencyStateActive {
		t.Fatalf("expected non-active sufficiency state with stale evidence, got %q", got)
	}
	model = IntelligenceCalibrationValAVEXSufficiencyContract()
	model.UnsupportedEvidenceRefs = []string{"evidence:unsupported"}
	if got := EvaluateIntelligenceCalibrationValAVEXSufficiencyState(model); got == IntelligenceCalibrationValAVEXSufficiencyStateActive {
		t.Fatalf("expected non-active sufficiency state with unsupported evidence, got %q", got)
	}
}

func TestIntelligenceCalibrationValARedactedExplanationDoesNotConvertInsufficientIntoSufficient(t *testing.T) {
	model := IntelligenceCalibrationValAExplanationContract()
	model.RedactionTurnsInsufficientSufficient = true
	if got := EvaluateIntelligenceCalibrationValAExplanationState(model); got == IntelligenceCalibrationValAExplanationStateActive {
		t.Fatalf("expected non-active explanation state when redaction converts insufficient into sufficient, got %q", got)
	}
}

func TestIntelligenceCalibrationValAExplanationDistinguishesNotEvidencedFromSafe(t *testing.T) {
	model := IntelligenceCalibrationValAExplanationContract()
	model.DistinguishesNotEvidencedFromSafe = false
	if got := EvaluateIntelligenceCalibrationValAExplanationState(model); got == IntelligenceCalibrationValAExplanationStateActive {
		t.Fatalf("expected non-active explanation state when not-evidenced is not distinguished from safe, got %q", got)
	}
}

func TestIntelligenceCalibrationValAWeaklyInferredEvidenceCannotProduceHighConfidenceRelevantOutcome(t *testing.T) {
	model := IntelligenceCalibrationValAConfidenceOutcomeContract()
	model.EvidenceClass = IntelligenceCalibrationEvidenceWeaklyInferred
	model.OutcomeState = IntelligenceCalibrationValAOutcomeHighConfidence
	if got := EvaluateIntelligenceCalibrationValAConfidenceOutcomeState(model); got == IntelligenceCalibrationValAConfidenceOutcomeStateActive {
		t.Fatalf("expected non-active confidence outcome when weakly inferred evidence produces high confidence relevance, got %q", got)
	}
}

func TestIntelligenceCalibrationValAStaleFreshnessCapsConfidenceOrRequiresLimitation(t *testing.T) {
	model := IntelligenceCalibrationValAConfidenceOutcomeContract()
	model.FreshnessState = IntelligenceCalibrationFreshnessStale
	model.ConfidenceBand = IntelligenceCalibrationConfidenceHigh
	model.LimitationMessage = "bounded freshness window"
	if got := EvaluateIntelligenceCalibrationValAConfidenceOutcomeState(model); got == IntelligenceCalibrationValAConfidenceOutcomeStateActive {
		t.Fatalf("expected non-active confidence outcome when stale freshness is not properly capped, got %q", got)
	}
}

func TestIntelligenceCalibrationValANoFinalPublicationGuardrailBlocksFinalPublication(t *testing.T) {
	model := IntelligenceCalibrationValAPublicationGuardrailContract()
	model.PublicationAllowed = true
	if got := EvaluateIntelligenceCalibrationValAPublicationGuardrailState(model); got == IntelligenceCalibrationValAPublicationGuardrailStateActive {
		t.Fatalf("expected non-active publication guardrail when final publication is allowed, got %q", got)
	}
}

func TestIntelligenceCalibrationValAProofsCanBecomeActiveOnlyAsReachabilityAndVEXCalibrationWhilePoint5RemainsNotComplete(t *testing.T) {
	aggregationState, exploitabilityState, decisionState, caviState, vexState, sufficiencyState, explanationState, outcomeState, guardrailState := activeIntelligenceCalibrationValAStates()
	state := EvaluateIntelligenceCalibrationValAProofsState(
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationVal0StateActive,
		aggregationState,
		exploitabilityState,
		decisionState,
		caviState,
		vexState,
		sufficiencyState,
		explanationState,
		outcomeState,
		guardrailState,
		[]string{
			"/v1/intelligence/calibration/vala/reachability-aggregation",
			"/v1/intelligence/calibration/vala/exploitability-calibration",
			"/v1/intelligence/calibration/vala/downgrade-escalation",
			"/v1/intelligence/calibration/vala/cavi-tuning",
			"/v1/intelligence/calibration/vala/vex-candidates",
			"/v1/intelligence/calibration/vala/vex-sufficiency",
			"/v1/intelligence/calibration/vala/explanations",
			"/v1/intelligence/calibration/vala/confidence-outcomes",
			"/v1/intelligence/calibration/vala/publication-guardrail",
			"/v1/intelligence/calibration/vala/proofs",
		},
		[]string{"val0_proofs", "reachability_signal_aggregation", "contextual_exploitability_calibration", "downgrade_escalation_guardrail", "cavi_tuning_contract", "vex_candidate_calibration", "vex_sufficiency_check", "reachability_vex_explanation"},
		[]string{"Val A remains bounded and advisory only."},
		[]string{"Točka 5 still needs later vals and integrated closure."},
		"projection_only not_canonical_truth advisory_reachability_vex_calibration",
	)
	if state != IntelligenceCalibrationValAStateActive {
		t.Fatalf("expected active Val A proofs state, got %q", state)
	}
}

func TestIntelligenceCalibrationValAMissingRequiredComponentKeepsValAInactive(t *testing.T) {
	aggregationState, exploitabilityState, decisionState, caviState, vexState, sufficiencyState, explanationState, outcomeState, guardrailState := activeIntelligenceCalibrationValAStates()
	if got := EvaluateIntelligenceCalibrationValAState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAAggregationStateIncomplete, exploitabilityState, decisionState, caviState, vexState, sufficiencyState, explanationState, outcomeState, guardrailState); got == IntelligenceCalibrationValAStateActive {
		t.Fatalf("expected non-active Val A state when a required component is missing, got %q", got)
	}
	if got := EvaluateIntelligenceCalibrationValAState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, aggregationState, exploitabilityState, decisionState, caviState, vexState, sufficiencyState, explanationState, outcomeState, guardrailState); got != IntelligenceCalibrationValAStateActive {
		t.Fatalf("expected active Val A state for default contract set, got %q", got)
	}
}

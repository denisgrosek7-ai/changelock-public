package operability

import "testing"

func activeIntelligenceCalibrationValDStates() (string, string, string, string, string, string, string, string, string, string, string) {
	return EvaluateIntelligenceCalibrationValDSimulationHarnessState(IntelligenceCalibrationValDDefensiveSimulationHarnessContract()),
		EvaluateIntelligenceCalibrationValDScenarioLibraryState(IntelligenceCalibrationValDScenarioLibraryContract()),
		EvaluateIntelligenceCalibrationValDMissedDetectionState(IntelligenceCalibrationValDMissedDetectionAnalysisContract()),
		EvaluateIntelligenceCalibrationValDFPFNBalanceState(IntelligenceCalibrationValDFPFNBalanceReviewContract()),
		EvaluateIntelligenceCalibrationValDConfidenceReviewState(IntelligenceCalibrationValDConfidenceCalibrationReviewContract()),
		EvaluateIntelligenceCalibrationValDVEXQualityState(IntelligenceCalibrationValDVEXQualityReviewContract()),
		EvaluateIntelligenceCalibrationValDReachabilityQualityState(IntelligenceCalibrationValDReachabilityQualityReviewContract()),
		EvaluateIntelligenceCalibrationValDBehavioralQualityState(IntelligenceCalibrationValDBehavioralQualityReviewContract()),
		EvaluateIntelligenceCalibrationValDFederatedQualityState(IntelligenceCalibrationValDFederatedQualityReviewContract()),
		EvaluateIntelligenceCalibrationValDSimulationCoverageState(IntelligenceCalibrationValDSimulationCoverageReviewContract()),
		EvaluateIntelligenceCalibrationValDQualityScoreboardState(IntelligenceCalibrationValDQualityScoreboardContract())
}

func TestIntelligenceCalibrationValDFailsClosedWithoutActiveDependencies(t *testing.T) {
	harnessState, libraryState, missedState, balanceState, confidenceState, vexState, reachabilityState, behavioralState, federatedState, coverageState, scoreboardState := activeIntelligenceCalibrationValDStates()

	if got := EvaluateIntelligenceCalibrationValDState(IntelligenceCalibrationVal0StateIncomplete, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValCStateActive, IntelligenceCalibrationValCStateActive, harnessState, libraryState, missedState, balanceState, confidenceState, vexState, reachabilityState, behavioralState, federatedState, coverageState, scoreboardState); got == IntelligenceCalibrationValDStateActive {
		t.Fatalf("expected non-active Val D state without active Val 0 dependency, got %q", got)
	}
	if got := EvaluateIntelligenceCalibrationValDState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateIncomplete, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValCStateActive, IntelligenceCalibrationValCStateActive, harnessState, libraryState, missedState, balanceState, confidenceState, vexState, reachabilityState, behavioralState, federatedState, coverageState, scoreboardState); got == IntelligenceCalibrationValDStateActive {
		t.Fatalf("expected non-active Val D state without active Val A dependency, got %q", got)
	}
	if got := EvaluateIntelligenceCalibrationValDState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateIncomplete, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValCStateActive, IntelligenceCalibrationValCStateActive, harnessState, libraryState, missedState, balanceState, confidenceState, vexState, reachabilityState, behavioralState, federatedState, coverageState, scoreboardState); got == IntelligenceCalibrationValDStateActive {
		t.Fatalf("expected non-active Val D state without active Val B dependency, got %q", got)
	}
	if got := EvaluateIntelligenceCalibrationValDState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValCStateIncomplete, IntelligenceCalibrationValCStateActive, harnessState, libraryState, missedState, balanceState, confidenceState, vexState, reachabilityState, behavioralState, federatedState, coverageState, scoreboardState); got == IntelligenceCalibrationValDStateActive {
		t.Fatalf("expected non-active Val D state without active Val C dependency, got %q", got)
	}
}

func TestIntelligenceCalibrationValDSimulationHarnessRequiresRefsCountsOutcomesAndStaleLimitation(t *testing.T) {
	model := IntelligenceCalibrationValDDefensiveSimulationHarnessContract()
	model.ScenarioRefs = nil
	if got := EvaluateIntelligenceCalibrationValDSimulationHarnessState(model); got == IntelligenceCalibrationValDSimulationHarnessStateActive {
		t.Fatalf("expected non-active simulation harness state without scenario refs, got %q", got)
	}

	model = IntelligenceCalibrationValDDefensiveSimulationHarnessContract()
	model.ScenarioCount = 0
	if got := EvaluateIntelligenceCalibrationValDSimulationHarnessState(model); got == IntelligenceCalibrationValDSimulationHarnessStateActive {
		t.Fatalf("expected non-active simulation harness state without scenario count, got %q", got)
	}

	model = IntelligenceCalibrationValDDefensiveSimulationHarnessContract()
	model.ExpectedOutcomesPresent = false
	if got := EvaluateIntelligenceCalibrationValDSimulationHarnessState(model); got == IntelligenceCalibrationValDSimulationHarnessStateActive {
		t.Fatalf("expected non-active simulation harness state without expected outcomes, got %q", got)
	}

	model = IntelligenceCalibrationValDDefensiveSimulationHarnessContract()
	model.ActualOutcomesPresent = false
	if got := EvaluateIntelligenceCalibrationValDSimulationHarnessState(model); got == IntelligenceCalibrationValDSimulationHarnessStateActive {
		t.Fatalf("expected non-active simulation harness state without actual outcomes, got %q", got)
	}

	model = IntelligenceCalibrationValDDefensiveSimulationHarnessContract()
	model.FreshnessState = IntelligenceCalibrationFreshnessStale
	model.LimitationMessage = "bounded simulation harness"
	if got := EvaluateIntelligenceCalibrationValDSimulationHarnessState(model); got == IntelligenceCalibrationValDSimulationHarnessStateActive {
		t.Fatalf("expected non-active simulation harness state when stale freshness lacks limitation, got %q", got)
	}
}

func TestIntelligenceCalibrationValDScenarioLibraryRequiresLowSignalAdversarialCriticalBlindSpotsAndLimitations(t *testing.T) {
	model := IntelligenceCalibrationValDScenarioLibraryContract()
	model.LowSignalScenarios = nil
	if got := EvaluateIntelligenceCalibrationValDScenarioLibraryState(model); got == IntelligenceCalibrationValDScenarioLibraryStateActive {
		t.Fatalf("expected non-active scenario library state without low-signal scenarios, got %q", got)
	}

	model = IntelligenceCalibrationValDScenarioLibraryContract()
	model.AdversarialScenarios = nil
	if got := EvaluateIntelligenceCalibrationValDScenarioLibraryState(model); got == IntelligenceCalibrationValDScenarioLibraryStateActive {
		t.Fatalf("expected non-active scenario library state without adversarial scenarios, got %q", got)
	}

	model = IntelligenceCalibrationValDScenarioLibraryContract()
	model.CriticalAssetScenarios = nil
	if got := EvaluateIntelligenceCalibrationValDScenarioLibraryState(model); got == IntelligenceCalibrationValDScenarioLibraryStateActive {
		t.Fatalf("expected non-active scenario library state without critical asset scenarios, got %q", got)
	}

	model = IntelligenceCalibrationValDScenarioLibraryContract()
	model.BlindSpotTags = nil
	if got := EvaluateIntelligenceCalibrationValDScenarioLibraryState(model); got == IntelligenceCalibrationValDScenarioLibraryStateActive {
		t.Fatalf("expected non-active scenario library state without blind-spot tags, got %q", got)
	}

	model = IntelligenceCalibrationValDScenarioLibraryContract()
	model.CoverageLimitations = nil
	if got := EvaluateIntelligenceCalibrationValDScenarioLibraryState(model); got == IntelligenceCalibrationValDScenarioLibraryStateActive {
		t.Fatalf("expected non-active scenario library state without coverage limitations, got %q", got)
	}

	model = IntelligenceCalibrationValDScenarioLibraryContract()
	model.LowSignalScenarios = []string{"scenario/not-in-main-list"}
	if got := EvaluateIntelligenceCalibrationValDScenarioLibraryState(model); got == IntelligenceCalibrationValDScenarioLibraryStateActive {
		t.Fatalf("expected non-active scenario library state when low-signal scenario ref is not in scenarios, got %q", got)
	}

	model = IntelligenceCalibrationValDScenarioLibraryContract()
	model.AdversarialScenarios = []string{"scenario/not-in-main-list"}
	if got := EvaluateIntelligenceCalibrationValDScenarioLibraryState(model); got == IntelligenceCalibrationValDScenarioLibraryStateActive {
		t.Fatalf("expected non-active scenario library state when adversarial scenario ref is not in scenarios, got %q", got)
	}

	model = IntelligenceCalibrationValDScenarioLibraryContract()
	model.CriticalAssetScenarios = []string{"scenario/not-in-main-list"}
	if got := EvaluateIntelligenceCalibrationValDScenarioLibraryState(model); got == IntelligenceCalibrationValDScenarioLibraryStateActive {
		t.Fatalf("expected non-active scenario library state when critical-asset scenario ref is not in scenarios, got %q", got)
	}

	model = IntelligenceCalibrationValDScenarioLibraryContract()
	if got := EvaluateIntelligenceCalibrationValDScenarioLibraryState(model); got != IntelligenceCalibrationValDScenarioLibraryStateActive {
		t.Fatalf("expected active scenario library state when categorized scenario refs are present in scenarios, got %q", got)
	}
}

func TestIntelligenceCalibrationValDMissedDetectionAnalysisKeepsFNRiskVisibleAndReviewBound(t *testing.T) {
	model := IntelligenceCalibrationValDMissedDetectionAnalysisContract()
	model.FalseNegativeCases = nil
	if got := EvaluateIntelligenceCalibrationValDMissedDetectionState(model); got == IntelligenceCalibrationValDMissedDetectionStateActive {
		t.Fatalf("expected non-active missed-detection state without false-negative cases, got %q", got)
	}

	model = IntelligenceCalibrationValDMissedDetectionAnalysisContract()
	model.SuppressionCausedMissCases = nil
	if got := EvaluateIntelligenceCalibrationValDMissedDetectionState(model); got == IntelligenceCalibrationValDMissedDetectionStateActive {
		t.Fatalf("expected non-active missed-detection state without suppression-caused miss analysis, got %q", got)
	}

	model = IntelligenceCalibrationValDMissedDetectionAnalysisContract()
	model.ReviewerRequired = false
	if got := EvaluateIntelligenceCalibrationValDMissedDetectionState(model); got == IntelligenceCalibrationValDMissedDetectionStateActive {
		t.Fatalf("expected non-active missed-detection state when critical asset miss cases lack reviewer gate, got %q", got)
	}

	model = IntelligenceCalibrationValDMissedDetectionAnalysisContract()
	model.RootCauseCategories = nil
	if got := EvaluateIntelligenceCalibrationValDMissedDetectionState(model); got == IntelligenceCalibrationValDMissedDetectionStateActive {
		t.Fatalf("expected non-active missed-detection state without root-cause categories, got %q", got)
	}
}

func TestIntelligenceCalibrationValDFPFNBalanceRequiresFNReviewAndMissChecks(t *testing.T) {
	model := IntelligenceCalibrationValDFPFNBalanceReviewContract()
	model.FNRiskReviewed = false
	if got := EvaluateIntelligenceCalibrationValDFPFNBalanceState(model); got == IntelligenceCalibrationValDFPFNBalanceStateActive {
		t.Fatalf("expected non-active FP/FN balance state when FP reduction lacks FN risk review, got %q", got)
	}

	model = IntelligenceCalibrationValDFPFNBalanceReviewContract()
	model.SuppressionCausedMissChecked = false
	if got := EvaluateIntelligenceCalibrationValDFPFNBalanceState(model); got == IntelligenceCalibrationValDFPFNBalanceStateActive {
		t.Fatalf("expected non-active FP/FN balance state without suppression-caused miss check, got %q", got)
	}

	model = IntelligenceCalibrationValDFPFNBalanceReviewContract()
	model.CriticalLowSignalReviewed = false
	if got := EvaluateIntelligenceCalibrationValDFPFNBalanceState(model); got == IntelligenceCalibrationValDFPFNBalanceStateActive {
		t.Fatalf("expected non-active FP/FN balance state without critical low-signal review, got %q", got)
	}

	model = IntelligenceCalibrationValDFPFNBalanceReviewContract()
	model.FalseNegativeReviewMetricRef = ""
	if got := EvaluateIntelligenceCalibrationValDFPFNBalanceState(model); got == IntelligenceCalibrationValDFPFNBalanceStateActive {
		t.Fatalf("expected non-active FP/FN balance state without false-negative metric ref, got %q", got)
	}
}

func TestIntelligenceCalibrationValDConfidenceReviewBlocksUnsupportedAndWeakInferenceHighConfidence(t *testing.T) {
	model := IntelligenceCalibrationValDConfidenceCalibrationReviewContract()
	model.UnsupportedEvidenceHighConfidenceBlocked = false
	if got := EvaluateIntelligenceCalibrationValDConfidenceReviewState(model); got == IntelligenceCalibrationValDConfidenceReviewStateActive {
		t.Fatalf("expected non-active confidence review state when unsupported evidence can become high confidence, got %q", got)
	}

	model = IntelligenceCalibrationValDConfidenceCalibrationReviewContract()
	model.WeaklyInferredHighConfidenceBlocked = false
	if got := EvaluateIntelligenceCalibrationValDConfidenceReviewState(model); got == IntelligenceCalibrationValDConfidenceReviewStateActive {
		t.Fatalf("expected non-active confidence review state when weakly inferred evidence can become high confidence, got %q", got)
	}

	model = IntelligenceCalibrationValDConfidenceCalibrationReviewContract()
	model.ConfidenceCalibrationErrorMetricRef = ""
	if got := EvaluateIntelligenceCalibrationValDConfidenceReviewState(model); got == IntelligenceCalibrationValDConfidenceReviewStateActive {
		t.Fatalf("expected non-active confidence review state without calibration error metric, got %q", got)
	}
}

func TestIntelligenceCalibrationValDVEXQualityKeepsFinalPublicationBlocked(t *testing.T) {
	model := IntelligenceCalibrationValDVEXQualityReviewContract()
	model.FinalVEXPublicationBlocked = false
	if got := EvaluateIntelligenceCalibrationValDVEXQualityState(model); got == IntelligenceCalibrationValDVEXQualityStateActive {
		t.Fatalf("expected non-active VEX quality state when final publication is not blocked, got %q", got)
	}

	model = IntelligenceCalibrationValDVEXQualityReviewContract()
	model.NotAffectedRequiresSufficientEvidence = false
	if got := EvaluateIntelligenceCalibrationValDVEXQualityState(model); got == IntelligenceCalibrationValDVEXQualityStateActive {
		t.Fatalf("expected non-active VEX quality state when not_affected does not require sufficient evidence, got %q", got)
	}

	model = IntelligenceCalibrationValDVEXQualityReviewContract()
	model.StaleOrUnsupportedReviewedBlocked = false
	if got := EvaluateIntelligenceCalibrationValDVEXQualityState(model); got == IntelligenceCalibrationValDVEXQualityStateActive {
		t.Fatalf("expected non-active VEX quality state when stale/unsupported reviewed candidates are allowed, got %q", got)
	}
}

func TestIntelligenceCalibrationValDReachabilityQualityPreservesEvidenceBoundaries(t *testing.T) {
	model := IntelligenceCalibrationValDReachabilityQualityReviewContract()
	model.PackagePresenceOnlyBlocked = false
	if got := EvaluateIntelligenceCalibrationValDReachabilityQualityState(model); got == IntelligenceCalibrationValDReachabilityQualityStateActive {
		t.Fatalf("expected non-active reachability quality state when package presence alone is allowed, got %q", got)
	}

	model = IntelligenceCalibrationValDReachabilityQualityReviewContract()
	model.RuntimeLoadedOnlyBlocked = false
	if got := EvaluateIntelligenceCalibrationValDReachabilityQualityState(model); got == IntelligenceCalibrationValDReachabilityQualityStateActive {
		t.Fatalf("expected non-active reachability quality state when runtime-loaded-only path is allowed, got %q", got)
	}

	model = IntelligenceCalibrationValDReachabilityQualityReviewContract()
	model.DowngradeEvidenceRequired = false
	if got := EvaluateIntelligenceCalibrationValDReachabilityQualityState(model); got == IntelligenceCalibrationValDReachabilityQualityStateActive {
		t.Fatalf("expected non-active reachability quality state when downgrade lacks evidence/explanation requirement, got %q", got)
	}
}

func TestIntelligenceCalibrationValDBehavioralQualityBlocksUnknownOrRelaxedBehavioralPaths(t *testing.T) {
	model := IntelligenceCalibrationValDBehavioralQualityReviewContract()
	model.StaleOrUnknownBaselineBlocked = false
	if got := EvaluateIntelligenceCalibrationValDBehavioralQualityState(model); got == IntelligenceCalibrationValDBehavioralQualityStateActive {
		t.Fatalf("expected non-active behavioral quality state when stale/unknown baseline can pass, got %q", got)
	}

	model = IntelligenceCalibrationValDBehavioralQualityReviewContract()
	model.LearningModeRelaxationBlocked = false
	if got := EvaluateIntelligenceCalibrationValDBehavioralQualityState(model); got == IntelligenceCalibrationValDBehavioralQualityStateActive {
		t.Fatalf("expected non-active behavioral quality state when learning mode relaxation is allowed, got %q", got)
	}

	model = IntelligenceCalibrationValDBehavioralQualityReviewContract()
	model.FNRiskNoteRequired = false
	if got := EvaluateIntelligenceCalibrationValDBehavioralQualityState(model); got == IntelligenceCalibrationValDBehavioralQualityStateActive {
		t.Fatalf("expected non-active behavioral quality state when threshold decrease lacks FN risk note requirement, got %q", got)
	}
}

func TestIntelligenceCalibrationValDFederatedQualityPreservesLocalEvidenceBoundaries(t *testing.T) {
	model := IntelligenceCalibrationValDFederatedQualityReviewContract()
	model.LocalEvidenceOverrideReviewed = false
	if got := EvaluateIntelligenceCalibrationValDFederatedQualityState(model); got == IntelligenceCalibrationValDFederatedQualityStateActive {
		t.Fatalf("expected non-active federated quality state when local evidence override is not reviewed, got %q", got)
	}

	model = IntelligenceCalibrationValDFederatedQualityReviewContract()
	model.LowSimilarityConfidenceNotIncreased = false
	if got := EvaluateIntelligenceCalibrationValDFederatedQualityState(model); got == IntelligenceCalibrationValDFederatedQualityStateActive {
		t.Fatalf("expected non-active federated quality state when low similarity can increase confidence, got %q", got)
	}

	model = IntelligenceCalibrationValDFederatedQualityReviewContract()
	model.RawLocalEvidenceNotPropagated = false
	if got := EvaluateIntelligenceCalibrationValDFederatedQualityState(model); got == IntelligenceCalibrationValDFederatedQualityStateActive {
		t.Fatalf("expected non-active federated quality state when raw local evidence can propagate, got %q", got)
	}
}

func TestIntelligenceCalibrationValDSimulationCoverageRequiresCompleteVisibleBoundedCoverage(t *testing.T) {
	model := IntelligenceCalibrationValDSimulationCoverageReviewContract()
	model.CriticalMissingClasses = []string{"suppression_safety"}
	if got := EvaluateIntelligenceCalibrationValDSimulationCoverageState(model); got == IntelligenceCalibrationValDSimulationCoverageStateActive {
		t.Fatalf("expected non-active simulation coverage state with missing critical scenario class, got %q", got)
	}

	model = IntelligenceCalibrationValDSimulationCoverageReviewContract()
	model.ClaimsExhaustiveDetection = true
	if got := EvaluateIntelligenceCalibrationValDSimulationCoverageState(model); got == IntelligenceCalibrationValDSimulationCoverageStateActive {
		t.Fatalf("expected non-active simulation coverage state when coverage claims exhaustive detection, got %q", got)
	}

	model = IntelligenceCalibrationValDSimulationCoverageReviewContract()
	model.CoveredScenarioClasses = model.CoveredScenarioClasses[:len(model.CoveredScenarioClasses)-1]
	model.MissingScenarioClasses = []string{}
	if got := EvaluateIntelligenceCalibrationValDSimulationCoverageState(model); got == IntelligenceCalibrationValDSimulationCoverageStateActive {
		t.Fatalf("expected non-active simulation coverage state when a required scenario class is not actually covered, got %q", got)
	}

	model = IntelligenceCalibrationValDSimulationCoverageReviewContract()
	model.CoveredScenarioClasses = []string{}
	model.MissingScenarioClasses = []string{}
	if got := EvaluateIntelligenceCalibrationValDSimulationCoverageState(model); got == IntelligenceCalibrationValDSimulationCoverageStateActive {
		t.Fatalf("expected non-active simulation coverage state when covered scenario classes are empty, got %q", got)
	}

	model = IntelligenceCalibrationValDSimulationCoverageReviewContract()
	if got := EvaluateIntelligenceCalibrationValDSimulationCoverageState(model); got != IntelligenceCalibrationValDSimulationCoverageStateActive {
		t.Fatalf("expected active simulation coverage state when covered scenario classes include all required classes, got %q", got)
	}
}

func TestIntelligenceCalibrationValDQualityScoreboardRequiresFPFNMetricsAndScopedClaims(t *testing.T) {
	model := IntelligenceCalibrationValDQualityScoreboardContract()
	model.FalsePositiveRateRef = ""
	if got := EvaluateIntelligenceCalibrationValDQualityScoreboardState(model); got == IntelligenceCalibrationValDQualityScoreboardStateActive {
		t.Fatalf("expected non-active quality scoreboard state without false-positive metric ref, got %q", got)
	}

	model = IntelligenceCalibrationValDQualityScoreboardContract()
	model.FalseNegativeReviewRateRef = ""
	if got := EvaluateIntelligenceCalibrationValDQualityScoreboardState(model); got == IntelligenceCalibrationValDQualityScoreboardStateActive {
		t.Fatalf("expected non-active quality scoreboard state without false-negative metric ref, got %q", got)
	}

	model = IntelligenceCalibrationValDQualityScoreboardContract()
	model.ClaimsUniversalQuality = true
	if got := EvaluateIntelligenceCalibrationValDQualityScoreboardState(model); got == IntelligenceCalibrationValDQualityScoreboardStateActive {
		t.Fatalf("expected non-active quality scoreboard state when it claims universal intelligence quality, got %q", got)
	}
}

func TestIntelligenceCalibrationValDStateRequiresAllComponents(t *testing.T) {
	_, libraryState, missedState, balanceState, confidenceState, vexState, reachabilityState, behavioralState, federatedState, coverageState, scoreboardState := activeIntelligenceCalibrationValDStates()
	if got := EvaluateIntelligenceCalibrationValDState(
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationValAStateActive,
		IntelligenceCalibrationValAStateActive,
		IntelligenceCalibrationValBStateActive,
		IntelligenceCalibrationValBStateActive,
		IntelligenceCalibrationValCStateActive,
		IntelligenceCalibrationValCStateActive,
		IntelligenceCalibrationValDSimulationHarnessStateIncomplete,
		libraryState,
		missedState,
		balanceState,
		confidenceState,
		vexState,
		reachabilityState,
		behavioralState,
		federatedState,
		coverageState,
		scoreboardState,
	); got == IntelligenceCalibrationValDStateActive {
		t.Fatalf("expected non-active Val D state when a required component is incomplete, got %q", got)
	}
}

func TestIntelligenceCalibrationValDProofsRemainScopedAndPoint5NotComplete(t *testing.T) {
	harnessState, libraryState, missedState, balanceState, confidenceState, vexState, reachabilityState, behavioralState, federatedState, coverageState, scoreboardState := activeIntelligenceCalibrationValDStates()
	if got := EvaluateIntelligenceCalibrationValDProofsState(
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationValAStateActive,
		IntelligenceCalibrationValAStateActive,
		IntelligenceCalibrationValBStateActive,
		IntelligenceCalibrationValBStateActive,
		IntelligenceCalibrationValCStateActive,
		IntelligenceCalibrationValCStateActive,
		harnessState,
		libraryState,
		missedState,
		balanceState,
		confidenceState,
		vexState,
		reachabilityState,
		behavioralState,
		federatedState,
		coverageState,
		scoreboardState,
		[]string{"/v1/intelligence/calibration/vald/simulation-harness", "/v1/intelligence/calibration/vald/scenario-library", "/v1/intelligence/calibration/vald/missed-detection-analysis", "/v1/intelligence/calibration/vald/fp-fn-balance", "/v1/intelligence/calibration/vald/confidence-review", "/v1/intelligence/calibration/vald/vex-quality", "/v1/intelligence/calibration/vald/reachability-quality", "/v1/intelligence/calibration/vald/behavioral-quality", "/v1/intelligence/calibration/vald/federated-quality", "/v1/intelligence/calibration/vald/simulation-coverage", "/v1/intelligence/calibration/vald/quality-scoreboard", "/v1/intelligence/calibration/vald/proofs"},
		[]string{"val0_proofs", "vala_proofs", "valb_proofs", "valc_proofs", "defensive_simulation_harness", "adversarial_scenario_library", "missed_detection_analysis", "false_positive_false_negative_balance_review", "confidence_calibration_review", "vex_candidate_quality_review", "reachability_calibration_quality_review", "behavioral_calibration_quality_review", "federated_weighting_quality_review", "simulation_coverage_review", "intelligence_quality_scoreboard"},
		[]string{"Val D is advisory only."},
		[]string{"Točka 5 remains not complete until Val E integrated closure."},
		"projection_only not_canonical_truth advisory_defensive_simulation_quality_measurement_gate",
	); got != IntelligenceCalibrationValDStateActive {
		t.Fatalf("expected active Val D proofs state with active dependencies and complete surfaces, got %q", got)
	}
}

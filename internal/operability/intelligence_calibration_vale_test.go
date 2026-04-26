package operability

import "testing"

func activeIntelligenceCalibrationValEIntegratedStates() (string, string, string, string, string, string, string, string) {
	return EvaluateIntelligenceCalibrationValEDependencyClosureState(IntelligenceCalibrationValEDependencyClosureContract()),
		EvaluateIntelligenceCalibrationValECoherenceReviewState(IntelligenceCalibrationValECoherenceReviewContract()),
		EvaluateIntelligenceCalibrationValEBoundaryReviewState(IntelligenceCalibrationValEBoundaryReviewContract()),
		EvaluateIntelligenceCalibrationValEReachabilityVEXSafetyState(IntelligenceCalibrationValEReachabilityVEXSafetyReviewContract()),
		EvaluateIntelligenceCalibrationValEBehavioralLearningSafetyState(IntelligenceCalibrationValEBehavioralLearningSafetyReviewContract()),
		EvaluateIntelligenceCalibrationValEFeedbackFederatedSafetyState(IntelligenceCalibrationValEFeedbackFederatedSafetyReviewContract()),
		EvaluateIntelligenceCalibrationValESimulationQualityState(IntelligenceCalibrationValESimulationQualityReviewContract()),
		EvaluateIntelligenceCalibrationValERegressionClosureState(IntelligenceCalibrationValERegressionClosureContract())
}

func activeIntelligenceCalibrationValEProofRefs() ([]string, []string, []string) {
	return []string{
			"/v1/intelligence/calibration/val0/proofs",
			"/v1/intelligence/calibration/vala/proofs",
			"/v1/intelligence/calibration/valb/proofs",
			"/v1/intelligence/calibration/valc/proofs",
			"/v1/intelligence/calibration/vald/proofs",
			"/v1/intelligence/calibration/vale/dependency-closure",
			"/v1/intelligence/calibration/vale/coherence-review",
			"/v1/intelligence/calibration/vale/pass-rule",
			"/v1/intelligence/calibration/vale/advisory-boundary",
			"/v1/intelligence/calibration/vale/reachability-vex-safety",
			"/v1/intelligence/calibration/vale/behavioral-learning-safety",
			"/v1/intelligence/calibration/vale/feedback-federated-safety",
			"/v1/intelligence/calibration/vale/simulation-quality-review",
			"/v1/intelligence/calibration/vale/regression-closure",
			"/v1/intelligence/calibration/vale/proofs",
		}, []string{
			"val0_proofs",
			"vala_proofs",
			"valb_proofs",
			"valc_proofs",
			"vald_proofs",
			"integrated_dependency_closure",
			"cross_val_coherence_review",
			"integrated_point5_pass_rule",
			"integrated_advisory_boundary_review",
			"integrated_reachability_vex_safety_review",
			"integrated_behavioral_learning_safety_review",
			"integrated_feedback_federated_safety_review",
			"integrated_simulation_quality_review",
			"integrated_regression_closure",
			"evidence_spine",
		}, []string{
			"integrated closure remains projection-only",
			"point 5 pass remains fail-closed on prior val proofs",
		}
}

func TestIntelligenceCalibrationValEFailsClosedWithoutActiveDependencies(t *testing.T) {
	dependencyState, coherenceState, boundaryState, reachabilityState, behavioralState, feedbackState, simulationState, regressionState := activeIntelligenceCalibrationValEIntegratedStates()

	if got := EvaluateIntelligenceCalibrationValEState(IntelligenceCalibrationVal0StateIncomplete, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValCStateActive, IntelligenceCalibrationValDStateActive, dependencyState, coherenceState, IntelligenceCalibrationValEPassRuleStateActive, boundaryState, reachabilityState, behavioralState, feedbackState, simulationState, regressionState); got == IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected non-active Val E state without active Val 0 dependency, got %q", got)
	}
	if got := EvaluateIntelligenceCalibrationValEState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateIncomplete, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValCStateActive, IntelligenceCalibrationValDStateActive, dependencyState, coherenceState, IntelligenceCalibrationValEPassRuleStateActive, boundaryState, reachabilityState, behavioralState, feedbackState, simulationState, regressionState); got == IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected non-active Val E state without active Val A dependency, got %q", got)
	}
	if got := EvaluateIntelligenceCalibrationValEState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateIncomplete, IntelligenceCalibrationValCStateActive, IntelligenceCalibrationValDStateActive, dependencyState, coherenceState, IntelligenceCalibrationValEPassRuleStateActive, boundaryState, reachabilityState, behavioralState, feedbackState, simulationState, regressionState); got == IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected non-active Val E state without active Val B dependency, got %q", got)
	}
	if got := EvaluateIntelligenceCalibrationValEState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValCStateIncomplete, IntelligenceCalibrationValDStateActive, dependencyState, coherenceState, IntelligenceCalibrationValEPassRuleStateActive, boundaryState, reachabilityState, behavioralState, feedbackState, simulationState, regressionState); got == IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected non-active Val E state without active Val C dependency, got %q", got)
	}
	if got := EvaluateIntelligenceCalibrationValEState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValCStateActive, IntelligenceCalibrationValDStateIncomplete, dependencyState, coherenceState, IntelligenceCalibrationValEPassRuleStateActive, boundaryState, reachabilityState, behavioralState, feedbackState, simulationState, regressionState); got == IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected non-active Val E state without active Val D dependency, got %q", got)
	}
	if got := EvaluateIntelligenceCalibrationValEState(IntelligenceCalibrationVal0StateActive, IntelligenceCalibrationValAStateActive, IntelligenceCalibrationValBStateActive, IntelligenceCalibrationValCStateActive, IntelligenceCalibrationValDStateActive, IntelligenceCalibrationValEDependencyClosureStatePartial, coherenceState, IntelligenceCalibrationValEPassRuleStateActive, boundaryState, reachabilityState, behavioralState, feedbackState, simulationState, regressionState); got == IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected non-active Val E state when dependency closure is partial, got %q", got)
	}
}

func TestIntelligenceCalibrationValEDependencyClosureUsesProofStatesAndBlocksInconsistency(t *testing.T) {
	model := IntelligenceCalibrationValEDependencyClosureContract()
	model.ProofStatesObserved = false
	if got := EvaluateIntelligenceCalibrationValEDependencyClosureState(model); got == IntelligenceCalibrationValEDependencyClosureStateActive {
		t.Fatalf("expected non-active dependency closure without observed proof states, got %q", got)
	}

	model = IntelligenceCalibrationValEDependencyClosureContract()
	model.ValBState = ""
	if got := EvaluateIntelligenceCalibrationValEDependencyClosureState(model); got == IntelligenceCalibrationValEDependencyClosureStateActive {
		t.Fatalf("expected non-active dependency closure with missing val state, got %q", got)
	}

	model = IntelligenceCalibrationValEDependencyClosureContract()
	model.InconsistentVals = []string{"val_a.proofs_state"}
	model.DependencyStatus = IntelligenceCalibrationValEDependencyPartial
	if got := EvaluateIntelligenceCalibrationValEDependencyClosureState(model); got == IntelligenceCalibrationValEDependencyClosureStateActive {
		t.Fatalf("expected non-active dependency closure with inconsistent val proofs, got %q", got)
	}

	if got := EvaluateIntelligenceCalibrationValEDependencyClosureState(IntelligenceCalibrationValEDependencyClosureContract()); got != IntelligenceCalibrationValEDependencyClosureStateActive {
		t.Fatalf("expected active dependency closure with complete proof-state-backed inputs, got %q", got)
	}
}

func TestIntelligenceCalibrationValECoherenceReviewBlocksMissingLinksAndPreservesLimitations(t *testing.T) {
	model := IntelligenceCalibrationValECoherenceReviewContract()
	model.MissingLinks = []string{"val0.contracts->vala.reachability_vex_calibration"}
	model.CoherenceState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValECoherenceReviewState(model); got == IntelligenceCalibrationValECoherenceReviewStateActive {
		t.Fatalf("expected non-active coherence review with missing critical links, got %q", got)
	}

	model = IntelligenceCalibrationValECoherenceReviewContract()
	model.InconsistentLinks = []string{"prior_vals->point5_not_complete_until_vale"}
	model.CoherenceState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValECoherenceReviewState(model); got == IntelligenceCalibrationValECoherenceReviewStateActive {
		t.Fatalf("expected non-active coherence review with inconsistent links, got %q", got)
	}

	model = IntelligenceCalibrationValECoherenceReviewContract()
	model.CarriedForwardLimitations = nil
	model.LimitationsCarriedForward = false
	model.CoherenceState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValECoherenceReviewState(model); got == IntelligenceCalibrationValECoherenceReviewStateActive {
		t.Fatalf("expected non-active coherence review without carried-forward limitations, got %q", got)
	}
}

func TestIntelligenceCalibrationValEPassRuleRequiresActiveValEAndNoBlockers(t *testing.T) {
	model := IntelligenceCalibrationValEPassRuleContract()
	if got := EvaluateIntelligenceCalibrationValEPassRuleState(model); got != IntelligenceCalibrationValEPassRuleStateActive {
		t.Fatalf("expected active pass-rule state with complete active inputs, got %q", got)
	}

	model = IntelligenceCalibrationValEPassRuleContract()
	model.ValEState = IntelligenceCalibrationValEStateSubstantial
	model.Point5State = IntelligenceCalibrationPoint5StateNotComplete
	model.PassCriteriaMet = false
	model.ActiveVals = []string{"val_0", "val_a", "val_b", "val_c", "val_d"}
	model.PartialVals = []string{"val_e"}
	if got := EvaluateIntelligenceCalibrationValEPassRuleState(model); got == IntelligenceCalibrationValEPassRuleStateActive {
		t.Fatalf("expected non-active pass-rule state without active Val E closure, got %q", got)
	}

	model = IntelligenceCalibrationValEPassRuleContract()
	model.PassBlockers = []string{"dependency closure has inconsistent vals"}
	model.Point5State = IntelligenceCalibrationPoint5StateNotComplete
	model.PassCriteriaMet = false
	if got := EvaluateIntelligenceCalibrationValEPassRuleState(model); got == IntelligenceCalibrationValEPassRuleStateActive {
		t.Fatalf("expected non-active pass-rule state with blockers, got %q", got)
	}
}

func TestIntelligenceCalibrationValEBoundaryReviewBlocksCanonicalTruthMutationAndFinalVEX(t *testing.T) {
	model := IntelligenceCalibrationValEBoundaryReviewContract()
	model.VEXCandidateOutputsRemainCandidates = false
	model.FinalVEXPublicationBlocked = false
	model.ViolationSurfaces = []string{"/v1/intelligence/calibration/vala/publication-guardrail"}
	model.BoundaryState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEBoundaryReviewState(model); got == IntelligenceCalibrationValEBoundaryReviewStateActive {
		t.Fatalf("expected non-active boundary review when final VEX path is open, got %q", got)
	}

	model = IntelligenceCalibrationValEBoundaryReviewContract()
	model.NoMutationWithoutGovernance = false
	model.ViolationSurfaces = []string{"/v1/intelligence/calibration/vale/proofs"}
	model.BoundaryState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEBoundaryReviewState(model); got == IntelligenceCalibrationValEBoundaryReviewStateActive {
		t.Fatalf("expected non-active boundary review when advisory mutation without governance is possible, got %q", got)
	}
}

func TestIntelligenceCalibrationValEReachabilityVEXSafetyBlocksUnsafePaths(t *testing.T) {
	model := IntelligenceCalibrationValEReachabilityVEXSafetyReviewContract()
	model.PackagePresenceOnlyBlocked = false
	model.Blockers = []string{"package presence alone can imply exploitability"}
	model.ReachabilityVEXState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEReachabilityVEXSafetyState(model); got == IntelligenceCalibrationValEReachabilityVEXSafetyStateActive {
		t.Fatalf("expected non-active reachability/vex safety review when package-presence-only path is allowed, got %q", got)
	}

	model = IntelligenceCalibrationValEReachabilityVEXSafetyReviewContract()
	model.InsufficientEvidenceBlocksNotAffected = false
	model.Blockers = []string{"insufficient evidence can produce not_affected"}
	model.ReachabilityVEXState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEReachabilityVEXSafetyState(model); got == IntelligenceCalibrationValEReachabilityVEXSafetyStateActive {
		t.Fatalf("expected non-active reachability/vex safety review when insufficient evidence can produce not_affected, got %q", got)
	}

	model = IntelligenceCalibrationValEReachabilityVEXSafetyReviewContract()
	model.FinalVEXClaimBlocked = false
	model.Blockers = []string{"final vex claim is not blocked"}
	model.ReachabilityVEXState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEReachabilityVEXSafetyState(model); got == IntelligenceCalibrationValEReachabilityVEXSafetyStateActive {
		t.Fatalf("expected non-active reachability/vex safety review when final_vex_claim is not blocked, got %q", got)
	}
}

func TestIntelligenceCalibrationValEBehavioralLearningSafetyBlocksStaleBaselineRelaxationAndPriorityLowering(t *testing.T) {
	model := IntelligenceCalibrationValEBehavioralLearningSafetyReviewContract()
	model.ActiveBaselineFreshnessFresh = false
	model.Blockers = []string{"active behavioral baseline is not fresh"}
	model.BehavioralLearningState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEBehavioralLearningSafetyState(model); got == IntelligenceCalibrationValEBehavioralLearningSafetyStateActive {
		t.Fatalf("expected non-active behavioral/learning safety review when active baseline freshness is stale/unknown, got %q", got)
	}

	model = IntelligenceCalibrationValEBehavioralLearningSafetyReviewContract()
	model.LearningModeCriticalControlRelaxationBlocked = false
	model.Blockers = []string{"learning mode can relax critical controls"}
	model.BehavioralLearningState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEBehavioralLearningSafetyState(model); got == IntelligenceCalibrationValEBehavioralLearningSafetyStateActive {
		t.Fatalf("expected non-active behavioral/learning safety review when learning mode can relax critical controls, got %q", got)
	}

	model = IntelligenceCalibrationValEBehavioralLearningSafetyReviewContract()
	model.LowerPriorityCandidateRequiresReviewerGate = false
	model.Blockers = []string{"lower-priority candidate can proceed without reviewer gate"}
	model.BehavioralLearningState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEBehavioralLearningSafetyState(model); got == IntelligenceCalibrationValEBehavioralLearningSafetyStateActive {
		t.Fatalf("expected non-active behavioral/learning safety review when priority lowering lacks reviewer gate, got %q", got)
	}
}

func TestIntelligenceCalibrationValEFeedbackFederatedSafetyBlocksMutationDeletionOverrideAndPropagation(t *testing.T) {
	model := IntelligenceCalibrationValEFeedbackFederatedSafetyReviewContract()
	model.FeedbackDoesNotMutateIntelligence = false
	model.Blockers = []string{"feedback can mutate intelligence"}
	model.FeedbackFederatedState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEFeedbackFederatedSafetyState(model); got == IntelligenceCalibrationValEFeedbackFederatedSafetyStateActive {
		t.Fatalf("expected non-active feedback/federated safety review when feedback can mutate intelligence, got %q", got)
	}

	model = IntelligenceCalibrationValEFeedbackFederatedSafetyReviewContract()
	model.SuppressionDoesNotDeleteEvidence = false
	model.Blockers = []string{"suppression can delete evidence"}
	model.FeedbackFederatedState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEFeedbackFederatedSafetyState(model); got == IntelligenceCalibrationValEFeedbackFederatedSafetyStateActive {
		t.Fatalf("expected non-active feedback/federated safety review when suppression deletes evidence, got %q", got)
	}

	model = IntelligenceCalibrationValEFeedbackFederatedSafetyReviewContract()
	model.SimilarityGatingDoesNotOverrideLocalEvidence = false
	model.Blockers = []string{"federated path overrides local evidence"}
	model.FeedbackFederatedState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEFeedbackFederatedSafetyState(model); got == IntelligenceCalibrationValEFeedbackFederatedSafetyStateActive {
		t.Fatalf("expected non-active feedback/federated safety review when federated path overrides local evidence, got %q", got)
	}

	model = IntelligenceCalibrationValEFeedbackFederatedSafetyReviewContract()
	model.RawLocalEvidenceDoesNotPropagate = false
	model.Blockers = []string{"raw local evidence propagates"}
	model.FeedbackFederatedState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValEFeedbackFederatedSafetyState(model); got == IntelligenceCalibrationValEFeedbackFederatedSafetyStateActive {
		t.Fatalf("expected non-active feedback/federated safety review when raw local evidence propagates, got %q", got)
	}
}

func TestIntelligenceCalibrationValESimulationQualityBlocksCoverageUniversalClaimsAndMissingFPFN(t *testing.T) {
	model := IntelligenceCalibrationValESimulationQualityReviewContract()
	model.MissingCoverage = []string{"critical_asset_escalation"}
	model.Blockers = []string{"missing critical scenario class"}
	model.SimulationQualityState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValESimulationQualityState(model); got == IntelligenceCalibrationValESimulationQualityStateActive {
		t.Fatalf("expected non-active simulation/quality review when critical scenario class coverage is missing, got %q", got)
	}

	model = IntelligenceCalibrationValESimulationQualityReviewContract()
	model.ScoreboardDoesNotClaimUniversalIntelligenceQuality = false
	model.Blockers = []string{"scoreboard claims universal intelligence quality"}
	model.SimulationQualityState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValESimulationQualityState(model); got == IntelligenceCalibrationValESimulationQualityStateActive {
		t.Fatalf("expected non-active simulation/quality review when universal intelligence quality is claimed, got %q", got)
	}

	model = IntelligenceCalibrationValESimulationQualityReviewContract()
	model.ScoreboardIncludesFPAndFNMetrics = false
	model.Blockers = []string{"scoreboard missing fp/fn metrics"}
	model.SimulationQualityState = IntelligenceCalibrationValEReviewBlocked
	if got := EvaluateIntelligenceCalibrationValESimulationQualityState(model); got == IntelligenceCalibrationValESimulationQualityStateActive {
		t.Fatalf("expected non-active simulation/quality review when FP/FN metrics are missing, got %q", got)
	}
}

func TestIntelligenceCalibrationValERegressionClosureBlocksMissingCriticalRegressionCategories(t *testing.T) {
	model := IntelligenceCalibrationValERegressionClosureContract()
	model.CriticalMissingCategories = []string{"integrated_pass_rule"}
	if got := EvaluateIntelligenceCalibrationValERegressionClosureState(model); got == IntelligenceCalibrationValERegressionClosureStateActive {
		t.Fatalf("expected non-active regression closure with critical missing categories, got %q", got)
	}
}

func TestIntelligenceCalibrationValEProofsStateRequiresPoint5PassForActive(t *testing.T) {
	dependencyState, coherenceState, boundaryState, reachabilityState, behavioralState, feedbackState, simulationState, regressionState := activeIntelligenceCalibrationValEIntegratedStates()
	surfaceRefs, evidenceRefs, limitations := activeIntelligenceCalibrationValEProofRefs()

	activePassRule := EvaluateIntelligenceCalibrationValEPassRuleState(IntelligenceCalibrationValEPassRuleContract())
	if got := EvaluateIntelligenceCalibrationValEProofsState(
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationValAStateActive,
		IntelligenceCalibrationValBStateActive,
		IntelligenceCalibrationValCStateActive,
		IntelligenceCalibrationValDStateActive,
		dependencyState,
		coherenceState,
		activePassRule,
		boundaryState,
		reachabilityState,
		behavioralState,
		feedbackState,
		simulationState,
		regressionState,
		IntelligenceCalibrationPoint5StatePass,
		surfaceRefs,
		evidenceRefs,
		limitations,
		"projection_only not_canonical_truth integrated_intelligence_calibration_closure",
	); got != IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected active Val E proofs state with point5 pass, got %q", got)
	}

	if got := EvaluateIntelligenceCalibrationValEProofsState(
		IntelligenceCalibrationVal0StateActive,
		IntelligenceCalibrationValAStateActive,
		IntelligenceCalibrationValBStateActive,
		IntelligenceCalibrationValCStateActive,
		IntelligenceCalibrationValDStateActive,
		dependencyState,
		coherenceState,
		activePassRule,
		boundaryState,
		reachabilityState,
		behavioralState,
		feedbackState,
		simulationState,
		regressionState,
		IntelligenceCalibrationPoint5StateNotComplete,
		surfaceRefs,
		evidenceRefs,
		limitations,
		"projection_only not_canonical_truth integrated_intelligence_calibration_closure",
	); got == IntelligenceCalibrationValEStateActive {
		t.Fatalf("expected non-active Val E proofs state when point5_state is not_complete, got %q", got)
	}
}

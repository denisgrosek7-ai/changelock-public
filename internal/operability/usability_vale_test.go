package operability

import (
	"strings"
	"testing"
)

func testProductionUsabilityValESurfaceRefs() []string {
	return []string{
		"/v1/production/usability-operability-recovery/val0/proofs",
		"/v1/production/usability-operability-recovery/vala/proofs",
		"/v1/production/usability-operability-recovery/valb/proofs",
		"/v1/production/usability-operability-recovery/valc/proofs",
		"/v1/production/usability-operability-recovery/vald/proofs",
		"/v1/production/usability-operability-recovery/vale/dependency-closure",
		"/v1/production/usability-operability-recovery/vale/coherence-review",
		"/v1/production/usability-operability-recovery/vale/pass-rule",
		"/v1/production/usability-operability-recovery/vale/canonical-truth-boundary",
		"/v1/production/usability-operability-recovery/vale/redaction-export-review",
		"/v1/production/usability-operability-recovery/vale/supportability-recovery-review",
		"/v1/production/usability-operability-recovery/vale/regression-closure",
		"/v1/production/usability-operability-recovery/vale/proofs",
	}
}

func testProductionUsabilityValEEvidenceRefs() []string {
	return []string{
		"val0_proofs",
		"vala_proofs",
		"valb_proofs",
		"valc_proofs",
		"vald_proofs",
		"dependency_closure_contract",
		"coherence_review_contract",
		"point4_pass_rule_contract",
		"canonical_truth_boundary_review_contract",
		"redaction_export_review_contract",
		"supportability_recovery_review_contract",
		"regression_closure_contract",
		"evidence_spine",
	}
}

func TestProductionUsabilityValEStateRequiresActiveVal0(t *testing.T) {
	if got := EvaluateProductionUsabilityValEState(
		ProductionUsabilityVal0StateSubstantial,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDStateActive,
		ProductionUsabilityValEDependencyClosureStateActive,
		ProductionUsabilityValECoherenceReviewStateActive,
		ProductionUsabilityValEPassRuleStateActive,
		ProductionUsabilityValECanonicalBoundaryReviewStateActive,
		ProductionUsabilityValERedactionExportReviewStateActive,
		ProductionUsabilityValESupportabilityRecoveryReviewStateActive,
		ProductionUsabilityValERegressionClosureStateActive,
	); got == ProductionUsabilityValEStateActive {
		t.Fatalf("expected non-active Val E state without active Val 0, got %q", got)
	}
}

func TestProductionUsabilityValEStateRequiresActiveValA(t *testing.T) {
	if got := EvaluateProductionUsabilityValEState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateSubstantial,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDStateActive,
		ProductionUsabilityValEDependencyClosureStateActive,
		ProductionUsabilityValECoherenceReviewStateActive,
		ProductionUsabilityValEPassRuleStateActive,
		ProductionUsabilityValECanonicalBoundaryReviewStateActive,
		ProductionUsabilityValERedactionExportReviewStateActive,
		ProductionUsabilityValESupportabilityRecoveryReviewStateActive,
		ProductionUsabilityValERegressionClosureStateActive,
	); got == ProductionUsabilityValEStateActive {
		t.Fatalf("expected non-active Val E state without active Val A, got %q", got)
	}
}

func TestProductionUsabilityValEStateRequiresActiveValB(t *testing.T) {
	if got := EvaluateProductionUsabilityValEState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateSubstantial,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDStateActive,
		ProductionUsabilityValEDependencyClosureStateActive,
		ProductionUsabilityValECoherenceReviewStateActive,
		ProductionUsabilityValEPassRuleStateActive,
		ProductionUsabilityValECanonicalBoundaryReviewStateActive,
		ProductionUsabilityValERedactionExportReviewStateActive,
		ProductionUsabilityValESupportabilityRecoveryReviewStateActive,
		ProductionUsabilityValERegressionClosureStateActive,
	); got == ProductionUsabilityValEStateActive {
		t.Fatalf("expected non-active Val E state without active Val B, got %q", got)
	}
}

func TestProductionUsabilityValEStateRequiresActiveValC(t *testing.T) {
	if got := EvaluateProductionUsabilityValEState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateSubstantial,
		ProductionUsabilityValDStateActive,
		ProductionUsabilityValEDependencyClosureStateActive,
		ProductionUsabilityValECoherenceReviewStateActive,
		ProductionUsabilityValEPassRuleStateActive,
		ProductionUsabilityValECanonicalBoundaryReviewStateActive,
		ProductionUsabilityValERedactionExportReviewStateActive,
		ProductionUsabilityValESupportabilityRecoveryReviewStateActive,
		ProductionUsabilityValERegressionClosureStateActive,
	); got == ProductionUsabilityValEStateActive {
		t.Fatalf("expected non-active Val E state without active Val C, got %q", got)
	}
}

func TestProductionUsabilityValEStateRequiresActiveValD(t *testing.T) {
	if got := EvaluateProductionUsabilityValEState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDStateSubstantial,
		ProductionUsabilityValEDependencyClosureStateActive,
		ProductionUsabilityValECoherenceReviewStateActive,
		ProductionUsabilityValEPassRuleStateActive,
		ProductionUsabilityValECanonicalBoundaryReviewStateActive,
		ProductionUsabilityValERedactionExportReviewStateActive,
		ProductionUsabilityValESupportabilityRecoveryReviewStateActive,
		ProductionUsabilityValERegressionClosureStateActive,
	); got == ProductionUsabilityValEStateActive {
		t.Fatalf("expected non-active Val E state without active Val D, got %q", got)
	}
}

func TestProductionUsabilityValEStateFailsClosedForPartialDependencyState(t *testing.T) {
	if got := EvaluateProductionUsabilityValEState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDStateActive,
		ProductionUsabilityValEDependencyClosureStatePartial,
		ProductionUsabilityValECoherenceReviewStateActive,
		ProductionUsabilityValEPassRuleStateActive,
		ProductionUsabilityValECanonicalBoundaryReviewStateActive,
		ProductionUsabilityValERedactionExportReviewStateActive,
		ProductionUsabilityValESupportabilityRecoveryReviewStateActive,
		ProductionUsabilityValERegressionClosureStateActive,
	); got == ProductionUsabilityValEStateActive {
		t.Fatalf("expected non-active Val E state with partial dependency closure, got %q", got)
	}
}

func TestProductionUsabilityValEDependencyClosureDoesNotInferPassFromRoutePresenceAlone(t *testing.T) {
	model := ProductionUsabilityValEDependencyClosure()
	model.ProofStatesObserved = false
	if got := EvaluateProductionUsabilityValEDependencyClosureState(model); got == ProductionUsabilityValEDependencyClosureStateActive {
		t.Fatalf("expected non-active dependency closure without observed proof states, got %q", got)
	}
}

func TestProductionUsabilityValECoherenceReviewBlocksMissingCriticalLinks(t *testing.T) {
	model := ProductionUsabilityValECrossValCoherenceReview()
	model.MissingLinks = []string{"valc.supportability->vald.final_usability_gate"}
	if got := EvaluateProductionUsabilityValECoherenceReviewState(model); got == ProductionUsabilityValECoherenceReviewStateActive {
		t.Fatalf("expected non-active coherence review with missing critical links, got %q", got)
	}
}

func TestProductionUsabilityValECoherenceReviewBlocksInconsistentLinks(t *testing.T) {
	model := ProductionUsabilityValECrossValCoherenceReview()
	model.InconsistentLinks = []string{"prior_vals->point4_not_complete_until_vale"}
	if got := EvaluateProductionUsabilityValECoherenceReviewState(model); got == ProductionUsabilityValECoherenceReviewStateActive {
		t.Fatalf("expected non-active coherence review with inconsistent links, got %q", got)
	}
}

func TestProductionUsabilityValECoherenceReviewRequiresCarriedForwardLimitations(t *testing.T) {
	model := ProductionUsabilityValECrossValCoherenceReview()
	model.CarriedForwardLimitations = nil
	model.LimitationsCarriedForward = false
	if got := EvaluateProductionUsabilityValECoherenceReviewState(model); got == ProductionUsabilityValECoherenceReviewStateActive {
		t.Fatalf("expected non-active coherence review without carried-forward limitations, got %q", got)
	}
}

func TestProductionUsabilityValEPassRuleReturnsNotCompleteUnlessValEIsActive(t *testing.T) {
	model := ProductionUsabilityValEPoint4PassRule()
	model.ValEState = ProductionUsabilityValEStateSubstantial
	model.Point4State = ProductionUsabilityPoint4StateNotComplete
	model.PassCriteriaMet = false
	model.ActiveVals = []string{"val_0", "val_a", "val_b", "val_c", "val_d"}
	if got := EvaluateProductionUsabilityValEPassRuleState(model); got == ProductionUsabilityValEPassRuleStateActive {
		t.Fatalf("expected non-active pass rule when Val E is not active, got %q", got)
	}
}

func TestProductionUsabilityValEPassRuleBlocksPassWhenRequiredValIsMissing(t *testing.T) {
	model := ProductionUsabilityValEPoint4PassRule()
	model.ActiveVals = []string{"val_0", "val_a", "val_b", "val_c", "val_d"}
	model.MissingVals = []string{"val_e"}
	model.Point4State = ProductionUsabilityPoint4StateNotComplete
	model.PassCriteriaMet = false
	if got := EvaluateProductionUsabilityValEPassRuleState(model); got == ProductionUsabilityValEPassRuleStateActive {
		t.Fatalf("expected non-active pass rule when required val is missing, got %q", got)
	}
}

func TestProductionUsabilityValEPassRuleBlocksPassWhenIntegratedBlockerExists(t *testing.T) {
	model := ProductionUsabilityValEPoint4PassRule()
	model.PassBlockers = []string{"dependency closure has inactive vals"}
	model.Point4State = ProductionUsabilityPoint4StateNotComplete
	model.PassCriteriaMet = false
	if got := EvaluateProductionUsabilityValEPassRuleState(model); got == ProductionUsabilityValEPassRuleStateActive {
		t.Fatalf("expected non-active pass rule when blockers exist, got %q", got)
	}
}

func TestProductionUsabilityValEPassRuleCanReturnPassOnlyWhenAllDepsAndValEAreActive(t *testing.T) {
	model := ProductionUsabilityValEPoint4PassRule()
	if got := EvaluateProductionUsabilityValEPassRuleState(model); got != ProductionUsabilityValEPassRuleStateActive {
		t.Fatalf("expected active pass rule when all deps and Val E are active, got %q", got)
	}
}

func TestProductionUsabilityValECanonicalBoundaryReviewBlocksProjectionClaimingCanonicalTruth(t *testing.T) {
	model := ProductionUsabilityValECanonicalTruthBoundaryReview()
	model.ProjectionClaimsCanonicalTruth = true
	if got := EvaluateProductionUsabilityValECanonicalBoundaryReviewState(model); got == ProductionUsabilityValECanonicalBoundaryReviewStateActive {
		t.Fatalf("expected non-active canonical boundary review when projection claims canonical truth, got %q", got)
	}
}

func TestProductionUsabilityValECanonicalBoundaryReviewBlocksAdvisoryMutationWithoutGovernance(t *testing.T) {
	model := ProductionUsabilityValECanonicalTruthBoundaryReview()
	model.AdvisoryMutationWithoutGovernance = true
	if got := EvaluateProductionUsabilityValECanonicalBoundaryReviewState(model); got == ProductionUsabilityValECanonicalBoundaryReviewStateActive {
		t.Fatalf("expected non-active canonical boundary review when advisory mutates without governance, got %q", got)
	}
}

func TestProductionUsabilityValECanonicalBoundaryReviewBlocksPublicPartnerRawEvidenceExposure(t *testing.T) {
	model := ProductionUsabilityValECanonicalTruthBoundaryReview()
	model.PublicOrPartnerRawEvidenceLeaked = true
	if got := EvaluateProductionUsabilityValECanonicalBoundaryReviewState(model); got == ProductionUsabilityValECanonicalBoundaryReviewStateActive {
		t.Fatalf("expected non-active canonical boundary review when raw partner/public evidence leaks, got %q", got)
	}
}

func TestProductionUsabilityValERedactionExportReviewBlocksMissingRequiredScopeCoverage(t *testing.T) {
	model := ProductionUsabilityValEIntegratedRedactionExportReview()
	model.MissingScopes = []string{ProductionUsabilityVisibilityPartner}
	if got := EvaluateProductionUsabilityValERedactionExportReviewState(model); got == ProductionUsabilityValERedactionExportReviewStateActive {
		t.Fatalf("expected non-active redaction/export review when scope coverage is missing, got %q", got)
	}
}

func TestProductionUsabilityValERedactionExportReviewBlocksPartnerPublicFullEvidence(t *testing.T) {
	model := ProductionUsabilityValEIntegratedRedactionExportReview()
	model.PartnerOrPublicExposeFullEvidence = true
	if got := EvaluateProductionUsabilityValERedactionExportReviewState(model); got == ProductionUsabilityValERedactionExportReviewStateActive {
		t.Fatalf("expected non-active redaction/export review when partner/public exposes full evidence, got %q", got)
	}
}

func TestProductionUsabilityValERedactionExportReviewPreservesHiddenRedactedMetadataRepresentation(t *testing.T) {
	model := ProductionUsabilityValEIntegratedRedactionExportReview()
	model.HiddenMetadataRepresented = false
	if got := EvaluateProductionUsabilityValERedactionExportReviewState(model); got == ProductionUsabilityValERedactionExportReviewStateActive {
		t.Fatalf("expected non-active redaction/export review when hidden metadata is omitted, got %q", got)
	}
}

func TestProductionUsabilityValESupportabilityRecoveryReviewBlocksOverrideOfFailedProofState(t *testing.T) {
	model := ProductionUsabilityValESupportabilityRecoveryReview()
	model.SupportabilityOverridesFailedProof = true
	if got := EvaluateProductionUsabilityValESupportabilityRecoveryReviewState(model); got == ProductionUsabilityValESupportabilityRecoveryReviewStateActive {
		t.Fatalf("expected non-active supportability/recovery review when it overrides failed proof state, got %q", got)
	}
}

func TestProductionUsabilityValESupportabilityRecoveryReviewBlocksPolicyBypassRecommendation(t *testing.T) {
	model := ProductionUsabilityValESupportabilityRecoveryReview()
	model.RecoveryPolicyBypassRecommended = true
	if got := EvaluateProductionUsabilityValESupportabilityRecoveryReviewState(model); got == ProductionUsabilityValESupportabilityRecoveryReviewStateActive {
		t.Fatalf("expected non-active supportability/recovery review when policy bypass is recommended, got %q", got)
	}
}

func TestProductionUsabilityValESupportabilityRecoveryReviewBlocksMutatingUpgradeAdvisory(t *testing.T) {
	model := ProductionUsabilityValESupportabilityRecoveryReview()
	model.UpgradeAdvisoryMutatesState = true
	if got := EvaluateProductionUsabilityValESupportabilityRecoveryReviewState(model); got == ProductionUsabilityValESupportabilityRecoveryReviewStateActive {
		t.Fatalf("expected non-active supportability/recovery review when upgrade advisory mutates, got %q", got)
	}
}

func TestProductionUsabilityValERegressionClosureBlocksMissingCriticalCategory(t *testing.T) {
	model := ProductionUsabilityValERegressionClosure()
	model.CriticalMissingCategories = []string{"integrated_dependency_closure"}
	if got := EvaluateProductionUsabilityValERegressionClosureState(model); got == ProductionUsabilityValERegressionClosureStateActive {
		t.Fatalf("expected non-active regression closure with critical missing categories, got %q", got)
	}
}

func TestProductionUsabilityValEProofsStateCanBeActiveOnlyWhenPoint4Passes(t *testing.T) {
	if got := EvaluateProductionUsabilityValEProofsState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDStateActive,
		ProductionUsabilityValEDependencyClosureStateActive,
		ProductionUsabilityValECoherenceReviewStateActive,
		ProductionUsabilityValEPassRuleStateActive,
		ProductionUsabilityValECanonicalBoundaryReviewStateActive,
		ProductionUsabilityValERedactionExportReviewStateActive,
		ProductionUsabilityValESupportabilityRecoveryReviewStateActive,
		ProductionUsabilityValERegressionClosureStateActive,
		ProductionUsabilityPoint4StatePass,
		testProductionUsabilityValESurfaceRefs(),
		testProductionUsabilityValEEvidenceRefs(),
		[]string{"projection_only not_canonical_truth integrated closure"},
	); got != ProductionUsabilityValEStateActive {
		t.Fatalf("expected active proofs state when all integrated checks and point 4 pass are active, got %q", got)
	}
}

func TestProductionUsabilityValEProofsStateKeepsPoint4NotCompleteWhenValEIsIncomplete(t *testing.T) {
	if got := EvaluateProductionUsabilityValEProofsState(
		ProductionUsabilityVal0StateActive,
		ProductionUsabilityValAStateActive,
		ProductionUsabilityValBStateActive,
		ProductionUsabilityValCStateActive,
		ProductionUsabilityValDStateActive,
		ProductionUsabilityValEDependencyClosureStateActive,
		ProductionUsabilityValECoherenceReviewStateActive,
		ProductionUsabilityValEPassRuleStatePartial,
		ProductionUsabilityValECanonicalBoundaryReviewStateActive,
		ProductionUsabilityValERedactionExportReviewStateActive,
		ProductionUsabilityValESupportabilityRecoveryReviewStateActive,
		ProductionUsabilityValERegressionClosureStateActive,
		ProductionUsabilityPoint4StateNotComplete,
		testProductionUsabilityValESurfaceRefs(),
		testProductionUsabilityValEEvidenceRefs(),
		[]string{"projection_only not_canonical_truth integrated closure"},
	); got == ProductionUsabilityValEStateActive {
		t.Fatalf("expected non-active proofs state when Val E pass rule is partial, got %q", got)
	}
}

func TestProductionUsabilityValEIntegratedClosureOutputRemainsProjectionOnly(t *testing.T) {
	model := ProductionUsabilityValEDependencyClosure()
	if model.ProjectionDisclaimer == "" ||
		!strings.Contains(model.ProjectionDisclaimer, "projection_only") ||
		!strings.Contains(model.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected integrated closure disclaimer to remain projection-only and non-canonical, got %#v", model)
	}
}

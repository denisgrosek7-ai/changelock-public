package operability

import "testing"

func activeDeveloperEcosystemValAModel() DeveloperEcosystemValACore {
	model := DeveloperEcosystemValACoreModel()
	model.Dependency = DeveloperEcosystemValADependencySnapshot{
		Val0CurrentState:           DeveloperEcosystemVal0StateActive,
		Val0Point8State:            DeveloperEcosystemPoint8StateNotComplete,
		Val0OutputClassification:   DeveloperEcosystemVal0OutputClassificationStateActive,
		Val0IDEAdvisoryState:       DeveloperEcosystemVal0IDEAdvisoryStateActive,
		Val0LocalProductionState:   DeveloperEcosystemVal0LocalProductionStateActive,
		Val0RepoPolicyBoundary:     DeveloperEcosystemVal0RepoPolicyStateActive,
		Val0PluginSafetyState:      DeveloperEcosystemVal0PluginSafetyStateActive,
		Val0PerformanceBudgetState: DeveloperEcosystemVal0PerformanceBudgetStateActive,
		Val0DXMetricsState:         DeveloperEcosystemVal0DXMetricsStateActive,
		Val0NoOverclaimState:       DeveloperEcosystemVal0NoOverclaimStateActive,
		Val0ProofSurfaceRefs:       DeveloperEcosystemVal0ProofSurfaceRefs(),
		Val0EvidenceRefs:           DeveloperEcosystemVal0ProofEvidenceRefs(),
		Val0ProjectionDisclaimer:   developerEcosystemVal0ProjectionDisclaimer(),
	}
	return ComputeDeveloperEcosystemValACore(model)
}

func activeDeveloperEcosystemValALimitations() []string {
	return []string{
		"Val A defines IDE and local tooling core contracts only and does not implement production IDE marketplace publishing, SDK runtime, repo config runtime, plugin runtime, or Točka 9 work.",
		"Točka 8 remains not complete because later developer ecosystem waves are still required before any integrated closure can exist.",
		"IDE signals, local advisory, validation harness, mock verification, and inspect/explain outputs remain advisory only and cannot approve deployment, certify trust, or create canonical evidence.",
	}
}

func TestDeveloperEcosystemValAHappyPathAndPoint8NotComplete(t *testing.T) {
	model := activeDeveloperEcosystemValAModel()
	if model.CurrentState != DeveloperEcosystemValAStateActive {
		t.Fatalf("expected active developer Val A state, got %#v", model)
	}
	if model.Point8State != DeveloperEcosystemPoint8StateNotComplete {
		t.Fatalf("expected point 8 to remain not complete in Val A, got %#v", model)
	}
	if got := EvaluateDeveloperEcosystemValAProofsState(model, activeDeveloperEcosystemValALimitations()); got != DeveloperEcosystemValAStateActive {
		t.Fatalf("expected active developer Val A proofs state, got %q", got)
	}
}

func TestDeveloperEcosystemValADependencyAndContractBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValACore)
	}{
		{name: "val0 dependency partial blocks", mutate: func(model *DeveloperEcosystemValACore) {
			model.Dependency.Val0CurrentState = DeveloperEcosystemVal0StatePartial
		}},
		{name: "ide canonical truth blocked", mutate: func(model *DeveloperEcosystemValACore) {
			model.IDEBaseline.CanonicalTruthClaim = true
		}},
		{name: "ide deployment approval blocked", mutate: func(model *DeveloperEcosystemValACore) {
			model.IDEBaseline.DeploymentApprovalClaim = true
		}},
		{name: "ide missing freshness display prevents active", mutate: func(model *DeveloperEcosystemValACore) {
			model.IDEBaseline.FreshnessDisplay = false
		}},
		{name: "trust feedback suppresses failures", mutate: func(model *DeveloperEcosystemValACore) {
			model.TrustFeedback.RecommendationsSuppress = true
		}},
		{name: "cavi vex candidate cannot become reviewed", mutate: func(model *DeveloperEcosystemValACore) {
			model.CAVIVEXContext.CandidatePromotedToReviewed = true
		}},
		{name: "cavi vex cannot approve deployment", mutate: func(model *DeveloperEcosystemValACore) {
			model.CAVIVEXContext.DeploymentApprovalClaim = true
		}},
		{name: "local advisory cannot claim production equivalence", mutate: func(model *DeveloperEcosystemValACore) {
			model.LocalAdvisory.ProductionEquivalenceClaim = true
		}},
		{name: "local advisory cannot mutate canonical evidence", mutate: func(model *DeveloperEcosystemValACore) {
			model.LocalAdvisory.MutatesCanonicalEvidence = true
		}},
		{name: "validation harness unknown class fails closed", mutate: func(model *DeveloperEcosystemValACore) {
			model.ValidationHarness.UnknownValidationClass = true
		}},
		{name: "validation harness cannot claim production equivalence", mutate: func(model *DeveloperEcosystemValACore) {
			model.ValidationHarness.ProductionEquivalenceClaim = true
		}},
		{name: "mock verification cannot create canonical proof", mutate: func(model *DeveloperEcosystemValACore) {
			model.MockVerificationServer.CanonicalProofClaim = true
		}},
		{name: "mock verification stale fixture blocks", mutate: func(model *DeveloperEcosystemValACore) {
			model.MockVerificationServer.StaleFixtureDetected = true
		}},
		{name: "inspect explain cannot hide failures", mutate: func(model *DeveloperEcosystemValACore) {
			model.InspectExplain.FailureReasonsVisible = false
		}},
		{name: "inspect explain cannot hide production unknown", mutate: func(model *DeveloperEcosystemValACore) {
			model.InspectExplain.ProductionOnlyUnknownVisible = false
		}},
		{name: "degraded mode cannot silently bypass", mutate: func(model *DeveloperEcosystemValACore) {
			model.DegradedMode.SilentBypassAllowed = true
		}},
		{name: "degraded mode cannot hide failures", mutate: func(model *DeveloperEcosystemValACore) {
			model.DegradedMode.HiddenFailureSuppression = true
		}},
		{name: "no overclaim blocks point8 pass claim", mutate: func(model *DeveloperEcosystemValACore) {
			model.NoOverclaim.Point8PassClaim = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeveloperEcosystemValAModel()
		tc.mutate(&model)
		model = ComputeDeveloperEcosystemValACore(model)
		if model.CurrentState == DeveloperEcosystemValAStateActive {
			t.Fatalf("expected %s to prevent active state, got %#v", tc.name, model)
		}
		if model.Point8State != DeveloperEcosystemPoint8StateNotComplete {
			t.Fatalf("expected %s to keep point 8 not complete, got %#v", tc.name, model)
		}
	}
}

func TestDeveloperEcosystemValAProofSurfaceExactSet(t *testing.T) {
	model := activeDeveloperEcosystemValAModel()
	if got := EvaluateDeveloperEcosystemValAProofsState(model, activeDeveloperEcosystemValALimitations()); got != DeveloperEcosystemValAStateActive {
		t.Fatalf("expected exact Val A proof surface set to be active, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValACore)
	}{
		{name: "missing val0 proofs fails", mutate: func(model *DeveloperEcosystemValACore) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/val0/proofs")
		}},
		{name: "missing vala status fails", mutate: func(model *DeveloperEcosystemValACore) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/vala/status")
		}},
		{name: "duplicate proof ref fails", mutate: func(model *DeveloperEcosystemValACore) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/vala/proofs")
			model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/developer-ecosystem/val0/proofs")
		}},
		{name: "unknown extra proof ref fails", mutate: func(model *DeveloperEcosystemValACore) {
			model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/developer-ecosystem/vala/extra")
		}},
		{name: "whitespace proof ref fails", mutate: func(model *DeveloperEcosystemValACore) {
			model.ProofSurfaceRefs[0] = " "
		}},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemValAModel()
		tc.mutate(&mutated)
		if got := EvaluateDeveloperEcosystemValAProofsState(mutated, activeDeveloperEcosystemValALimitations()); got == DeveloperEcosystemValAStateActive {
			t.Fatalf("expected %s to fail exact proof validation, got %q", tc.name, got)
		}
	}
}

func TestDeveloperEcosystemValAEvidenceExactSet(t *testing.T) {
	model := activeDeveloperEcosystemValAModel()
	if !DeveloperEcosystemValAProofEvidenceQualityValid(developerEcosystemValAEvidence(), model.EvidenceRefs) {
		t.Fatalf("expected exact developer Val A evidence refs to be valid")
	}

	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValACore)
	}{
		{name: "missing ide baseline evidence fails", mutate: func(model *DeveloperEcosystemValACore) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "evidence:developer-ide-baseline-001")
		}},
		{name: "missing mock server evidence fails", mutate: func(model *DeveloperEcosystemValACore) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "evidence:developer-mock-verification-server-001")
		}},
		{name: "duplicate evidence ref fails", mutate: func(model *DeveloperEcosystemValACore) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "evidence:point8-vala-governance-001")
			model.EvidenceRefs = append(model.EvidenceRefs, "evidence:developer-ide-baseline-001")
		}},
		{name: "unknown extra evidence ref fails", mutate: func(model *DeveloperEcosystemValACore) {
			model.EvidenceRefs = append(model.EvidenceRefs, "evidence:developer-vala-extra-001")
		}},
		{name: "whitespace evidence ref fails", mutate: func(model *DeveloperEcosystemValACore) {
			model.EvidenceRefs[0] = " "
		}},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemValAModel()
		tc.mutate(&mutated)
		if DeveloperEcosystemValAProofEvidenceQualityValid(developerEcosystemValAEvidence(), mutated.EvidenceRefs) {
			t.Fatalf("expected %s to fail exact evidence validation", tc.name)
		}
	}
}

func TestDeveloperEcosystemValAVisibilityAndDegradedBehavior(t *testing.T) {
	model := activeDeveloperEcosystemValAModel()
	if len(model.TrustFeedback.SignalClasses) != 7 || len(model.InspectExplain.OutputClasses) != len(developerEcosystemVal0OutputClasses()) {
		t.Fatalf("expected trust feedback and inspect/explain to preserve classified signal/output surfaces, got %#v", model)
	}

	model = activeDeveloperEcosystemValAModel()
	model.TrustFeedback.ProductionOnlyUnknownShown = false
	model = ComputeDeveloperEcosystemValACore(model)
	if model.TrustFeedbackState == DeveloperEcosystemValATrustFeedbackStateActive {
		t.Fatalf("expected hidden production-only unknown trust feedback to prevent active state, got %#v", model)
	}

	model = activeDeveloperEcosystemValAModel()
	model.DegradedMode.FalseActiveClaim = true
	model = ComputeDeveloperEcosystemValACore(model)
	if model.DegradedModeState != DeveloperEcosystemValADegradedModeStateBlocked {
		t.Fatalf("expected false active degraded mode claim to block, got %#v", model)
	}
}

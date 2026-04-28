package operability

import "testing"

func activeDeveloperEcosystemValDModel() DeveloperEcosystemValDFinalGate {
	return ComputeDeveloperEcosystemValDFinalGate(DeveloperEcosystemValDFinalGateModel())
}

func activeDeveloperEcosystemValDLimitations() []string {
	return []string{
		"Val D is the final developer ecosystem gate only and cannot return point_8_pass or make Točka 8 complete.",
		"Integrated closure still requires Val E; Val D active is readiness consistency, not deployment approval, certification, production approval, or canonical truth.",
		"Clean-room and IP guardrail evidence is a static bounded repo check and does not claim legal certification, patent clearance, regulator approval, or formal legal opinion.",
	}
}

func TestDeveloperEcosystemValDHappyPathAndPoint8NotComplete(t *testing.T) {
	model := activeDeveloperEcosystemValDModel()
	if model.CurrentState != DeveloperEcosystemValDStateActive {
		t.Fatalf("expected active developer Val D state, got %#v", model)
	}
	if model.Point8State != DeveloperEcosystemPoint8StateNotComplete {
		t.Fatalf("expected point 8 to remain not complete in Val D, got %#v", model)
	}
	if model.FinalDeveloperEcosystemGateState != DeveloperEcosystemValDFinalGateStateActive {
		t.Fatalf("expected active final developer ecosystem gate, got %#v", model)
	}
	if model.VerifyPolicyCICompatibility.KyvernoVersion != DeveloperEcosystemValDVerifyPolicyKyvernoVersion {
		t.Fatalf("expected pinned Kyverno version, got %#v", model.VerifyPolicyCICompatibility)
	}
	if model.Val0Foundation.PluginSafetyBudgetRef != DeveloperEcosystemVal0PerformanceBudgetDisciplineID ||
		model.ValCReadiness.SandboxDisciplineID != DeveloperEcosystemValCSandboxIsolationDisciplineID ||
		model.ValCReadiness.SandboxVersion != DeveloperEcosystemValCSandboxIsolationVersion {
		t.Fatalf("expected canonical prior-wave references, got %#v", model)
	}
	if got := EvaluateDeveloperEcosystemValDProofsState(model, activeDeveloperEcosystemValDLimitations()); got != DeveloperEcosystemValDStateActive {
		t.Fatalf("expected active developer Val D proofs state, got %q", got)
	}
}

func TestDeveloperEcosystemValDPriorWaveAndCIBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValDFinalGate)
	}{
		{name: "vale point7 pass overclaim blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.ValECompatibility.Point7PassReason = "point_7_pass production approved"
		}},
		{name: "val0 partial blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.Val0Foundation.CurrentState = DeveloperEcosystemVal0StatePartial
		}},
		{name: "vala partial blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.ValAReadiness.CurrentState = DeveloperEcosystemValAStatePartial
		}},
		{name: "valb compatibility weakened blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.ValBReadiness.RepoConfigCompatibilityBehavior = "permissive"
		}},
		{name: "valc sandbox identity weakened blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.ValCReadiness.SandboxDisciplineID = "developer-ecosystem-plugin-sandbox-next"
		}},
		{name: "verify policy workflow exclusion weakened blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.VerifyPolicyCICompatibility.WorkflowFilesExcluded = false
		}},
		{name: "verify policy no input skip weakened blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.VerifyPolicyCICompatibility.EmptyManifestInputSkips = false
		}},
		{name: "verify policy actual input required weakened blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.VerifyPolicyCICompatibility.ActualManifestOrImageRequired = false
		}},
	}

	for _, tc := range testCases {
		model := activeDeveloperEcosystemValDModel()
		tc.mutate(&model)
		model = ComputeDeveloperEcosystemValDFinalGate(model)
		if model.CurrentState == DeveloperEcosystemValDStateActive {
			t.Fatalf("expected %s to prevent active state, got %#v", tc.name, model)
		}
		if model.Point8State != DeveloperEcosystemPoint8StateNotComplete {
			t.Fatalf("expected %s to keep point 8 not complete, got %#v", tc.name, model)
		}
	}
}

func TestDeveloperEcosystemValDReadinessAndBoundaryBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValDFinalGate)
	}{
		{name: "ide canonical truth claim blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.IDELocalReadiness.CanonicalTruthClaim = true
		}},
		{name: "local production equivalence blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.LocalMockNonEquivalence.ProductionEquivalenceClaim = true
		}},
		{name: "governance hidden mutation blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.GovernanceNoBypass.HiddenMutationPath = true
		}},
		{name: "governance hidden approval blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.GovernanceNoBypass.HiddenApprovalPath = true
		}},
		{name: "old dangling budget ref blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.PerformanceVisibility.ValCPluginExecutionBudgetRef = "developer-performance-budget"
		}},
		{name: "performance hidden failure blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.PerformanceVisibility.HiddenFailureSuppression = true
		}},
		{name: "examples certification claim blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.ExamplesNoCertification.SamplePluginCertificationClaim = true
		}},
		{name: "clean room legal certification claim blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.CleanRoomIPGuardrail.LegalCertificationClaim = true
		}},
		{name: "point8 pass claim blocks", mutate: func(model *DeveloperEcosystemValDFinalGate) {
			model.NoOverclaim.Point8PassClaim = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeveloperEcosystemValDModel()
		tc.mutate(&model)
		model = ComputeDeveloperEcosystemValDFinalGate(model)
		if model.CurrentState == DeveloperEcosystemValDStateActive {
			t.Fatalf("expected %s to prevent active state, got %#v", tc.name, model)
		}
	}
}

func TestDeveloperEcosystemValDVerifyPolicyExactContracts(t *testing.T) {
	model := activeDeveloperEcosystemValDModel()
	if model.VerifyPolicyCICompatibilityState != DeveloperEcosystemValDVerifyPolicyCICompatibilityStateActive {
		t.Fatalf("expected active verify-policy gate, got %#v", model.VerifyPolicyCICompatibility)
	}
	if !containsExactTrimmedStringSet(model.VerifyPolicyCICompatibility.TriggerOnlyPrefixes, developerEcosystemValDVerifyPolicyTriggerOnlyPrefixes()...) {
		t.Fatalf("expected exact trigger-only prefixes, got %#v", model.VerifyPolicyCICompatibility)
	}
	if !containsExactTrimmedStringSet(model.VerifyPolicyCICompatibility.ManifestResourcePrefixes, developerEcosystemValDVerifyPolicyManifestPrefixes()...) {
		t.Fatalf("expected exact manifest resource prefixes, got %#v", model.VerifyPolicyCICompatibility)
	}
	if !containsExactTrimmedStringSet(model.VerifyPolicyCICompatibility.OptionOnlyArgs, developerEcosystemValDVerifyPolicyOptionOnlyArgs()...) {
		t.Fatalf("expected exact option-only args, got %#v", model.VerifyPolicyCICompatibility)
	}

	mutated := activeDeveloperEcosystemValDModel()
	mutated.VerifyPolicyCICompatibility.ManifestResourcePrefixes = []string{".github/workflows", "deploy/k8s"}
	mutated = ComputeDeveloperEcosystemValDFinalGate(mutated)
	if mutated.VerifyPolicyCICompatibilityState == DeveloperEcosystemValDVerifyPolicyCICompatibilityStateActive {
		t.Fatalf("expected mixed trigger/resource prefixes to block, got %#v", mutated.VerifyPolicyCICompatibility)
	}
}

func TestDeveloperEcosystemValDProofSurfaceAndEvidenceExactSet(t *testing.T) {
	model := activeDeveloperEcosystemValDModel()
	if !DeveloperEcosystemValDProofEvidenceQualityValid(developerEcosystemValDEvidence(), model.EvidenceRefs) {
		t.Fatalf("expected active developer Val D evidence quality, got %#v", model.EvidenceRefs)
	}
	if !containsExactTrimmedStringSet(model.ProofSurfaceRefs, DeveloperEcosystemValDProofSurfaceRefs()...) {
		t.Fatalf("expected exact Val D proof surface refs, got %#v", model.ProofSurfaceRefs)
	}

	mutated := activeDeveloperEcosystemValDModel()
	mutated.ProofSurfaceRefs = append(mutated.ProofSurfaceRefs, "/v1/developer-ecosystem/vald/proofs")
	if got := EvaluateDeveloperEcosystemValDProofsState(mutated, activeDeveloperEcosystemValDLimitations()); got == DeveloperEcosystemValDStateActive {
		t.Fatalf("expected duplicate proof surface refs to prevent active proofs state, got %q", got)
	}

	mutated = activeDeveloperEcosystemValDModel()
	mutated.EvidenceRefs = append(mutated.EvidenceRefs, " ")
	if DeveloperEcosystemValDProofEvidenceQualityValid(developerEcosystemValDEvidence(), mutated.EvidenceRefs) {
		t.Fatalf("expected whitespace evidence ref to fail exact validation, got %#v", mutated.EvidenceRefs)
	}
}

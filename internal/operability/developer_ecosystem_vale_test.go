package operability

import "testing"

func activeDeveloperEcosystemValEModel() DeveloperEcosystemValEIntegratedClosure {
	return ComputeDeveloperEcosystemValEClosure(DeveloperEcosystemValEIntegratedClosureModel())
}

func TestDeveloperEcosystemValEHappyPathPassAndPoint8PassAllowed(t *testing.T) {
	model := activeDeveloperEcosystemValEModel()
	if model.CurrentState != DeveloperEcosystemValEStatePass {
		t.Fatalf("expected Val E pass state, got %#v", model)
	}
	if model.Point8State != DeveloperEcosystemPoint8StatePass || !model.Point8PassAllowed {
		t.Fatalf("expected point_8_pass only in Val E, got %#v", model)
	}
	if model.Point8PassReason != DeveloperEcosystemValEPoint8PassReasonAllowed {
		t.Fatalf("expected canonical allowed pass reason, got %#v", model)
	}
	if model.FinalPassRuleState != DeveloperEcosystemValEFinalPassRuleStateActive ||
		model.ClosureState != DeveloperEcosystemValEClosureStateActive {
		t.Fatalf("expected active closure and final pass rule, got %#v", model)
	}
	if model.Val0Source.Point8State != DeveloperEcosystemPoint8StateNotComplete ||
		model.ValASource.Point8State != DeveloperEcosystemPoint8StateNotComplete ||
		model.ValBSource.Point8State != DeveloperEcosystemPoint8StateNotComplete ||
		model.ValCSource.Point8State != DeveloperEcosystemPoint8StateNotComplete ||
		model.ValDSource.Point8State != DeveloperEcosystemPoint8StateNotComplete ||
		model.ValDSource.Point8PassAvailable {
		t.Fatalf("expected Val 0-D to remain non-pass prerequisites, got %#v", model)
	}
}

func TestDeveloperEcosystemValEDependencyAndGateBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemValEIntegratedClosure)
	}{
		{name: "vald final gate inactive", mutate: func(model *DeveloperEcosystemValEIntegratedClosure) {
			model.ValDSource.FinalDeveloperEcosystemGateState = DeveloperEcosystemValDFinalGateStatePartial
		}},
		{name: "vald point8 pass exposed", mutate: func(model *DeveloperEcosystemValEIntegratedClosure) {
			model.ValDSource.Point8PassAvailable = true
		}},
		{name: "valb compatibility weakened", mutate: func(model *DeveloperEcosystemValEIntegratedClosure) {
			model.ValBSource.RepoConfigCompatibilityBehavior = "permissive"
		}},
		{name: "valc sandbox identity weakened", mutate: func(model *DeveloperEcosystemValEIntegratedClosure) {
			model.ValCSource.SandboxDisciplineID = "developer-ecosystem-plugin-sandbox-next"
		}},
		{name: "verify policy classifier weakened", mutate: func(model *DeveloperEcosystemValEIntegratedClosure) {
			model.VerifyPolicyCICompatibility.WorkflowFilesExcluded = false
		}},
		{name: "clean room legal certification claim", mutate: func(model *DeveloperEcosystemValEIntegratedClosure) {
			model.CleanRoomIPGuardrail.LegalCertificationClaim = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeveloperEcosystemValEModel()
		tc.mutate(&model)
		model = ComputeDeveloperEcosystemValEClosure(model)
		if model.CurrentState == DeveloperEcosystemValEStatePass || model.Point8PassAllowed {
			t.Fatalf("expected %s to prevent pass, got %#v", tc.name, model)
		}
		if model.Point8State != DeveloperEcosystemPoint8StateNotComplete {
			t.Fatalf("expected %s to keep point 8 not complete, got %#v", tc.name, model)
		}
	}
}

func TestDeveloperEcosystemValEPoint8PassReasonDisciplineAndNoOverclaim(t *testing.T) {
	testCases := []struct {
		name       string
		reason     string
		mutate     func(*DeveloperEcosystemValEIntegratedClosure)
		wantNoOver string
		wantPass   string
	}{
		{name: "canonical blocked reason preserves fail closed without overclaim", reason: DeveloperEcosystemValEPoint8PassReasonBlocked, mutate: func(model *DeveloperEcosystemValEIntegratedClosure) {
			model.ValASource.CurrentState = DeveloperEcosystemValAStatePartial
		}, wantNoOver: DeveloperEcosystemValENoOverclaimStateActive, wantPass: DeveloperEcosystemValEFinalPassRuleStatePartial},
		{name: "production approved overclaim blocks", reason: "point_8_pass production approved", wantNoOver: DeveloperEcosystemValENoOverclaimStateBlocked, wantPass: DeveloperEcosystemValEFinalPassRuleStateBlocked},
		{name: "extended allowed reason plus certification blocks", reason: DeveloperEcosystemValEPoint8PassReasonAllowed + " certified", wantNoOver: DeveloperEcosystemValENoOverclaimStateBlocked, wantPass: DeveloperEcosystemValEFinalPassRuleStateBlocked},
		{name: "extended blocked reason plus approval blocks", reason: DeveloperEcosystemValEPoint8PassReasonBlocked + " production approved", wantNoOver: DeveloperEcosystemValENoOverclaimStateBlocked, wantPass: DeveloperEcosystemValEFinalPassRuleStateBlocked},
		{name: "substring bypass blocks", reason: "safe point_8_pass-ish maybe later", wantNoOver: DeveloperEcosystemValENoOverclaimStateBlocked, wantPass: DeveloperEcosystemValEFinalPassRuleStateBlocked},
	}

	for _, tc := range testCases {
		model := activeDeveloperEcosystemValEModel()
		model.Point8PassReason = tc.reason
		if tc.mutate != nil {
			tc.mutate(&model)
		}
		model = ComputeDeveloperEcosystemValEClosure(model)
		if model.NoOverclaimState != tc.wantNoOver || model.FinalPassRuleState != tc.wantPass {
			t.Fatalf("%s: got no-overclaim=%q pass=%q model=%#v", tc.name, model.NoOverclaimState, model.FinalPassRuleState, model)
		}
	}
}

func TestDeveloperEcosystemValEStateFidelityAndExactSets(t *testing.T) {
	model := activeDeveloperEcosystemValEModel()
	if !DeveloperEcosystemValEProofEvidenceQualityValid(developerEcosystemValEEvidence(), model.EvidenceRefs) {
		t.Fatalf("expected exact active evidence quality, got %#v", model.EvidenceRefs)
	}
	if !containsExactTrimmedStringSet(model.ProofSurfaceRefs, DeveloperEcosystemValEProofSurfaceRefs()...) {
		t.Fatalf("expected exact proof surface refs, got %#v", model.ProofSurfaceRefs)
	}

	partial := activeDeveloperEcosystemValEModel()
	partial.Point8PassReason = DeveloperEcosystemValEPoint8PassReasonBlocked
	partial.ValASource.CurrentState = DeveloperEcosystemValAStatePartial
	partial = ComputeDeveloperEcosystemValEClosure(partial)
	if partial.DependencyClosureState != DeveloperEcosystemValEDependencyClosureStatePartial ||
		partial.CurrentState != DeveloperEcosystemValEStatePartial ||
		partial.NoOverclaimState != DeveloperEcosystemValENoOverclaimStateActive {
		t.Fatalf("expected partial state fidelity with canonical blocked reason, got %#v", partial)
	}

	unknown := activeDeveloperEcosystemValEModel()
	unknown.Point8PassReason = DeveloperEcosystemValEPoint8PassReasonBlocked
	unknown.ValBSource.CurrentState = DeveloperEcosystemValBStateUnknown
	unknown = ComputeDeveloperEcosystemValEClosure(unknown)
	if unknown.CurrentState != DeveloperEcosystemValEStateUnknown {
		t.Fatalf("expected unknown state fidelity, got %#v", unknown)
	}

	mutated := activeDeveloperEcosystemValEModel()
	mutated.ProofSurfaceRefs = append(mutated.ProofSurfaceRefs, "/v1/developer-ecosystem/vale/proofs")
	mutated = ComputeDeveloperEcosystemValEClosure(mutated)
	if mutated.ProofSurfaceState == DeveloperEcosystemValEProofSurfaceStateActive {
		t.Fatalf("expected duplicate proof refs to fail exact validation, got %#v", mutated)
	}

	mutated = activeDeveloperEcosystemValEModel()
	mutated.EvidenceRefs = append(mutated.EvidenceRefs, " ")
	if DeveloperEcosystemValEProofEvidenceQualityValid(developerEcosystemValEEvidence(), mutated.EvidenceRefs) {
		t.Fatalf("expected whitespace evidence ref to fail exact validation, got %#v", mutated.EvidenceRefs)
	}
}

func TestDeveloperEcosystemValEVerifyPolicyAndBoundaryBlockers(t *testing.T) {
	model := activeDeveloperEcosystemValEModel()
	model.LocalMockNonEquivalence.ProductionEquivalenceClaim = true
	model = ComputeDeveloperEcosystemValEClosure(model)
	if model.LocalMockNonEquivalenceState != DeveloperEcosystemValELocalMockNonEquivalenceStateBlocked || model.Point8PassAllowed {
		t.Fatalf("expected production equivalence claim to block pass, got %#v", model)
	}

	model = activeDeveloperEcosystemValEModel()
	model.RepoSDKGovernanceBoundary.LocalPassBecomesCIPass = true
	model = ComputeDeveloperEcosystemValEClosure(model)
	if model.RepoSDKGovernanceBoundaryState != DeveloperEcosystemValERepoSDKGovernanceBoundaryStateBlocked {
		t.Fatalf("expected local pass to CI pass shortcut to block, got %#v", model)
	}

	model = activeDeveloperEcosystemValEModel()
	model.PluginExtensibilityBoundary.CustomChecksEmitPointPass = true
	model = ComputeDeveloperEcosystemValEClosure(model)
	if model.PluginExtensibilityBoundaryState != DeveloperEcosystemValEPluginExtensibilityBoundaryStateBlocked {
		t.Fatalf("expected plugin custom point pass claim to block, got %#v", model)
	}
}

package operability

import "testing"

func activeReferenceArchitectureValEModel() ReferenceArchitectureIntegratedClosure {
	return ReferenceArchitectureValEIntegratedClosureModel()
}

func TestReferenceArchitectureValEDependencyGates(t *testing.T) {
	model := activeReferenceArchitectureValEModel()
	if got := EvaluateReferenceArchitectureValEPrerequisiteState(model); got != ReferenceArchitectureValEPrerequisiteStateActive {
		t.Fatalf("expected active prerequisite state, got %q", got)
	}
	if model.CurrentState != ReferenceArchitectureValEStateActive || model.Point6State != ReferenceArchitecturePoint6StatePass || !model.Point6PassAllowed {
		t.Fatalf("expected active integrated closure with point_6_pass, got %#v", model)
	}

	testCases := []struct {
		name   string
		mutate func(*ReferenceArchitectureIntegratedClosure)
	}{
		{name: "missing val0", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.SourceCurrentStates.Val0CurrentState = ReferenceArchitectureVal0StateIncomplete
		}},
		{name: "missing vala", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.SourceCurrentStates.ValACurrentState = ReferenceArchitectureValAStatePartial
		}},
		{name: "missing valb", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.SourceCurrentStates.ValBCurrentState = ReferenceArchitectureValBStatePartial
		}},
		{name: "missing valc", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.SourceCurrentStates.ValCCurrentState = ReferenceArchitectureValCStatePartial
		}},
		{name: "missing vald", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.SourceCurrentStates.ValDCurrentState = ReferenceArchitectureValDStatePartial
		}},
		{name: "missing vald final gate", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.DependencyStates.ValDFinalGateState = ReferenceArchitectureValDFinalGateStatePartial
		}},
		{name: "point5 not pass", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.DependencyStates.Point5State = IntelligenceCalibrationPoint5StateNotComplete
		}},
		{name: "point5 dependency unhealthy", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.DependencyStates.Point5Dependency = IntelligenceCalibrationValEStateSubstantial
		}},
		{name: "route presence alone cannot pass", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.SourceCurrentStates.Val0CurrentState = ReferenceArchitectureVal0StateIncomplete
			model.ProofSurfaceRefs = referenceArchitectureValEProofSurfaceRefs()
		}},
	}

	for _, tc := range testCases {
		model := activeReferenceArchitectureValEModel()
		tc.mutate(&model)
		model = ComputeReferenceArchitectureValEClosure(model)
		if got := EvaluateReferenceArchitectureValEPrerequisiteState(model); got == ReferenceArchitectureValEPrerequisiteStateActive {
			t.Fatalf("expected non-active prerequisite state for %s, got %q", tc.name, got)
		}
		if model.Point6PassAllowed || model.Point6State == ReferenceArchitecturePoint6StatePass || model.CurrentState == ReferenceArchitectureValEStateActive {
			t.Fatalf("expected dependency regression to block point_6_pass for %s, got %#v", tc.name, model)
		}
	}
}

func TestReferenceArchitectureValEPoint6PassRule(t *testing.T) {
	model := activeReferenceArchitectureValEModel()
	if got := EvaluateReferenceArchitectureValEPassRuleState(model); got != ReferenceArchitectureValEPassRuleStateActive {
		t.Fatalf("expected active pass rule state, got %q", got)
	}
	if got := EvaluateReferenceArchitecturePoint6FinalState(model); got != ReferenceArchitecturePoint6StatePass {
		t.Fatalf("expected point_6_pass from Val E only, got %q", got)
	}

	model = activeReferenceArchitectureValEModel()
	model.ValB.ConformanceKitState = ReferenceArchitectureValBConformanceKitStatePartial
	model = ComputeReferenceArchitectureValEClosure(model)
	if model.Point6PassAllowed || model.Point6State == ReferenceArchitecturePoint6StatePass {
		t.Fatalf("expected partial inputs to block point_6_pass, got %#v", model)
	}

	model = activeReferenceArchitectureValEModel()
	model.StaleEvidenceDetected = true
	model.EvidenceFresh = false
	model = ComputeReferenceArchitectureValEClosure(model)
	if model.Point6PassAllowed || model.Point6State == ReferenceArchitecturePoint6StatePass {
		t.Fatalf("expected stale evidence to block point_6_pass, got %#v", model)
	}

	model = activeReferenceArchitectureValEModel()
	model.Point6State = ReferenceArchitecturePoint6StatePass
	model.SourceCurrentStates.Val0CurrentState = ReferenceArchitectureVal0StateIncomplete
	model = ComputeReferenceArchitectureValEClosure(model)
	if model.Point6State == ReferenceArchitecturePoint6StatePass {
		t.Fatalf("expected point_6_pass rule not to be self-referential, got %#v", model)
	}
}

func TestReferenceArchitectureValECrossValInvariants(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*ReferenceArchitectureIntegratedClosure)
	}{
		{name: "val0 unsupported conformance blocks closure", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.Val0.ConformanceState = ReferenceArchitectureConformanceUnsupported
		}},
		{name: "vala missing family blocks closure", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.ValA.SupportedFamilies = model.ValA.SupportedFamilies[:5]
		}},
		{name: "vala proof surface missing blocks closure", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.ValA.SurfaceRefs = model.ValA.SurfaceRefs[:3]
		}},
		{name: "valb invalid dependency collection blocks closure", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.ValB.ConformanceKitState = ReferenceArchitectureValBConformanceKitStatePartial
		}},
		{name: "valb invalid hook pack ref blocks closure", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.ValB.ValidationHookState = ReferenceArchitectureValBHookStatePartial
		}},
		{name: "valc unhealthy point5 dependency reporting blocks closure", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.ValC.Point5DependencyState = model.ValC.Point5State
		}},
		{name: "vald fake active-like component state blocks closure", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.ValD.OperationalVisibilityState = "not_a_real_state_active"
		}},
		{name: "vald whitespace duplicate family blocks closure", mutate: func(model *ReferenceArchitectureIntegratedClosure) {
			model.ValD.SupportedFamilies = []string{
				ReferenceArchitectureFamilyEnterpriseDefault,
				" enterprise_default ",
				ReferenceArchitectureFamilyHighAssurance,
				ReferenceArchitectureFamilyRegulatedPrivacyFirst,
				ReferenceArchitectureFamilySovereignAirGapped,
				ReferenceArchitectureFamilyPerformanceSensitive,
			}
		}},
	}

	for _, tc := range testCases {
		model := activeReferenceArchitectureValEModel()
		tc.mutate(&model)
		model = ComputeReferenceArchitectureValEClosure(model)
		if model.ClosureInvariantState == ReferenceArchitectureValEInvariantStateActive {
			t.Fatalf("expected invariant regression to fail closed for %s, got %#v", tc.name, model)
		}
		if model.Point6PassAllowed || model.Point6State == ReferenceArchitecturePoint6StatePass {
			t.Fatalf("expected invariant regression to block point_6_pass for %s, got %#v", tc.name, model)
		}
	}
}

func TestReferenceArchitectureValEProofSurfaceExactSet(t *testing.T) {
	model := activeReferenceArchitectureValEModel()
	if got := EvaluateReferenceArchitectureValEProofSurfaceState(model); got != ReferenceArchitectureValEProofSurfaceStateActive {
		t.Fatalf("expected active proof surface state, got %q", got)
	}

	for i, ref := range referenceArchitectureValEProofSurfaceRefs() {
		model := activeReferenceArchitectureValEModel()
		model.ProofSurfaceRefs = append([]string{}, referenceArchitectureValEProofSurfaceRefs()[:i]...)
		model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, referenceArchitectureValEProofSurfaceRefs()[i+1:]...)
		model = ComputeReferenceArchitectureValEClosure(model)
		if got := EvaluateReferenceArchitectureValEProofSurfaceState(model); got == ReferenceArchitectureValEProofSurfaceStateActive {
			t.Fatalf("expected missing surface %q to fail closed, got %q", ref, got)
		}
	}

	model = activeReferenceArchitectureValEModel()
	model.ProofSurfaceRefs = []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/proofs",
		"/v1/reference-architecture/valc/proofs",
		"/v1/reference-architecture/vald/proofs",
		"/v1/reference-architecture/vald/final-gate",
		"/v1/reference-architecture/vale/closure",
		"/v1/reference-architecture/vale/closure",
	}
	model = ComputeReferenceArchitectureValEClosure(model)
	if got := EvaluateReferenceArchitectureValEProofSurfaceState(model); got == ReferenceArchitectureValEProofSurfaceStateActive {
		t.Fatalf("expected duplicate surface refs not to compensate for missing required surface, got %q", got)
	}

	model = activeReferenceArchitectureValEModel()
	model.ProofSurfaceRefs = []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/proofs",
		"/v1/reference-architecture/valc/proofs",
		"/v1/reference-architecture/vald/proofs",
		"/v1/reference-architecture/vald/final-gate",
		"/v1/reference-architecture/vale/closure",
		"/v1/reference-architecture/vale/unknown",
	}
	model = ComputeReferenceArchitectureValEClosure(model)
	if got := EvaluateReferenceArchitectureValEProofSurfaceState(model); got == ReferenceArchitectureValEProofSurfaceStateActive {
		t.Fatalf("expected unknown extra surface not to compensate for missing required surface, got %q", got)
	}

	model = activeReferenceArchitectureValEModel()
	model.ProofSurfaceRefs[0] = "   "
	model = ComputeReferenceArchitectureValEClosure(model)
	if got := EvaluateReferenceArchitectureValEProofSurfaceState(model); got == ReferenceArchitectureValEProofSurfaceStateActive {
		t.Fatalf("expected whitespace surface ref to fail closed, got %q", got)
	}
}

func TestReferenceArchitectureValENoOverclaimAndFinalStateBehavior(t *testing.T) {
	model := activeReferenceArchitectureValEModel()
	model.Point6PassReason = "certified architecture"
	model = ComputeReferenceArchitectureValEClosure(model)
	if model.Point6PassAllowed || model.CurrentState == ReferenceArchitectureValEStateActive {
		t.Fatalf("expected certified architecture language to block point_6_pass, got %#v", model)
	}

	model = activeReferenceArchitectureValEModel()
	model.Limitations = []string{"deployment approved"}
	model = ComputeReferenceArchitectureValEClosure(model)
	if model.Point6PassAllowed || model.CurrentState == ReferenceArchitectureValEStateActive {
		t.Fatalf("expected deployment approval language to block point_6_pass, got %#v", model)
	}

	model = activeReferenceArchitectureValEModel()
	model.RedactionKeepsFailuresVisible = false
	model = ComputeReferenceArchitectureValEClosure(model)
	if model.Point6PassAllowed || model.Point6State == ReferenceArchitecturePoint6StatePass {
		t.Fatalf("expected redaction omission not to convert failure to pass, got %#v", model)
	}
	if model.Point6State != ReferenceArchitecturePoint6StateNotComplete || len(model.BlockingReasons) == 0 {
		t.Fatalf("expected non-pass state to preserve blocking reasons, got %#v", model)
	}
}

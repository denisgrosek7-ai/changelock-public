package operability

import "testing"

func activeDeveloperEcosystemVal0Model() DeveloperEcosystemVal0Foundation {
	model := DeveloperEcosystemVal0FoundationModel()
	model.Dependency = DeveloperEcosystemVal0DependencySnapshot{
		Point6State:                ReferenceArchitecturePoint6StatePass,
		Point7State:                VerifierEcosystemPoint7StatePass,
		Point7ClosureState:         VerifierEcosystemValEStatePass,
		Point7PrerequisiteState:    VerifierEcosystemValEPrerequisiteStateActive,
		Point7InvariantState:       VerifierEcosystemValEInvariantStateActive,
		Point7ProofSurfaceState:    VerifierEcosystemValEProofSurfaceStateActive,
		Point7EvidenceQualityState: VerifierEcosystemValEEvidenceQualityStateActive,
		Point7NoOverclaimState:     VerifierEcosystemValENoOverclaimStateActive,
		Point7PassRuleState:        VerifierEcosystemValEPassRuleStateActive,
		Point7PassAllowed:          true,
	}
	return ComputeDeveloperEcosystemVal0Foundation(model)
}

func activeDeveloperEcosystemVal0Limitations() []string {
	return []string{
		"Val 0 is only the developer discipline foundation and does not implement actual IDE, SDK, repo config runtime, mock runtime, or plugin runtime execution.",
		"Točka 8 remains not complete because later waves are still required before any integrated developer ecosystem closure can exist.",
		"Developer tooling outputs remain advisory and cannot approve deployment, certify trust, override enterprise governance, or mutate canonical evidence.",
	}
}

func TestDeveloperEcosystemVal0HappyPathAndPoint8NotComplete(t *testing.T) {
	model := activeDeveloperEcosystemVal0Model()
	if model.CurrentState != DeveloperEcosystemVal0StateActive {
		t.Fatalf("expected active developer Val 0 state, got %#v", model)
	}
	if model.Point8State != DeveloperEcosystemPoint8StateNotComplete {
		t.Fatalf("expected point 8 to remain not complete in Val 0, got %#v", model)
	}
	if got := EvaluateDeveloperEcosystemVal0ProofsState(model, activeDeveloperEcosystemVal0Limitations()); got != DeveloperEcosystemVal0StateActive {
		t.Fatalf("expected active developer proofs state, got %q", got)
	}
}

func TestDeveloperEcosystemVal0DependencyAndDisciplineBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemVal0Foundation)
	}{
		{name: "missing dependency blocks", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.Dependency.Point7State = VerifierEcosystemPoint7StateNotComplete
		}},
		{name: "advisory output cannot become approval", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.OutputClassification.AdvisoryTreatedAsApproval = true
		}},
		{name: "ide signal cannot become canonical truth", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.IDEAdvisory.CanonicalTruthClaim = true
		}},
		{name: "local mock cannot claim production equivalence", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.LocalProduction.ProductionEquivalenceClaim = true
		}},
		{name: "repo config cannot override enterprise policy", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.RepoPolicyBoundary.OverrideEnterprisePolicy = true
		}},
		{name: "plugin cannot mutate canonical evidence", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.PluginSafety.HiddenCanonicalMutation = true
		}},
		{name: "plugin cannot approve deployment", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.PluginSafety.HiddenApprovalPath = true
		}},
		{name: "unknown output class fails closed", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.OutputClassification.AllowedOutputClasses[0] = "unknownish"
		}},
		{name: "unsupported budget fails closed", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.PerformanceBudget.Budgets[0].BudgetState = DeveloperEcosystemBudgetStateUnsupported
		}},
	}

	for _, tc := range testCases {
		model := activeDeveloperEcosystemVal0Model()
		tc.mutate(&model)
		model = ComputeDeveloperEcosystemVal0Foundation(model)
		if model.CurrentState == DeveloperEcosystemVal0StateActive {
			t.Fatalf("expected %s to fail closed, got %#v", tc.name, model)
		}
	}
}

func TestDeveloperEcosystemVal0ProofSurfaceExactSet(t *testing.T) {
	model := activeDeveloperEcosystemVal0Model()
	if got := EvaluateDeveloperEcosystemVal0ProofsState(model, activeDeveloperEcosystemVal0Limitations()); got != DeveloperEcosystemVal0StateActive {
		t.Fatalf("expected exact proof surface set to be active, got %q", got)
	}

	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemVal0Foundation)
	}{
		{name: "missing vale proofs fails", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vale/proofs")
		}},
		{name: "missing status fails", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/val0/status")
		}},
		{name: "duplicate proof ref fails", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.ProofSurfaceRefs = removeTrimmedString(model.ProofSurfaceRefs, "/v1/developer-ecosystem/val0/proofs")
			model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/verifier-ecosystem/vale/proofs")
		}},
		{name: "unknown extra proof ref fails", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.ProofSurfaceRefs = append(model.ProofSurfaceRefs, "/v1/developer-ecosystem/val0/extra")
		}},
		{name: "whitespace proof ref fails", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.ProofSurfaceRefs[0] = " "
		}},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemVal0Model()
		tc.mutate(&mutated)
		if got := EvaluateDeveloperEcosystemVal0ProofsState(mutated, activeDeveloperEcosystemVal0Limitations()); got == DeveloperEcosystemVal0StateActive {
			t.Fatalf("expected %s to fail exact proof validation, got %q", tc.name, got)
		}
	}
}

func TestDeveloperEcosystemVal0EvidenceExactSet(t *testing.T) {
	model := activeDeveloperEcosystemVal0Model()
	if !DeveloperEcosystemVal0ProofEvidenceQualityValid(developerEcosystemVal0Evidence(), model.EvidenceRefs) {
		t.Fatalf("expected exact developer evidence refs to be valid")
	}

	testCases := []struct {
		name   string
		mutate func(*DeveloperEcosystemVal0Foundation)
	}{
		{name: "missing evidence ref fails", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "evidence:developer-no-overclaim-001")
		}},
		{name: "duplicate evidence ref fails", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.EvidenceRefs = removeTrimmedString(model.EvidenceRefs, "evidence:point8-governance-001")
			model.EvidenceRefs = append(model.EvidenceRefs, "evidence:developer-output-classification-001")
		}},
		{name: "unknown extra evidence ref fails", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.EvidenceRefs = append(model.EvidenceRefs, "evidence:developer-extra-001")
		}},
		{name: "whitespace evidence ref fails", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.EvidenceRefs[0] = " "
		}},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemVal0Model()
		tc.mutate(&mutated)
		if DeveloperEcosystemVal0ProofEvidenceQualityValid(developerEcosystemVal0Evidence(), mutated.EvidenceRefs) {
			t.Fatalf("expected %s to fail exact evidence validation", tc.name)
		}
	}
}

func TestDeveloperEcosystemVal0PerformanceBudgetAndDXMetrics(t *testing.T) {
	model := activeDeveloperEcosystemVal0Model()
	model.PerformanceBudget.Budgets[0].BudgetState = DeveloperEcosystemBudgetStateUnknown
	model = ComputeDeveloperEcosystemVal0Foundation(model)
	if model.PerformanceBudgetState != DeveloperEcosystemVal0PerformanceBudgetStateUnknown {
		t.Fatalf("expected unknown budget state to remain visible, got %#v", model)
	}

	model = activeDeveloperEcosystemVal0Model()
	model.PerformanceBudget.Budgets[0].BudgetState = DeveloperEcosystemBudgetStateDegraded
	model = ComputeDeveloperEcosystemVal0Foundation(model)
	if model.PerformanceBudgetState != DeveloperEcosystemVal0PerformanceBudgetStatePartial {
		t.Fatalf("expected degraded budget state to remain visible, got %#v", model)
	}

	model = activeDeveloperEcosystemVal0Model()
	model.DXMetrics.DeveloperTrustScore = true
	model = ComputeDeveloperEcosystemVal0Foundation(model)
	if model.DXMetricsState != DeveloperEcosystemVal0DXMetricsStateBlocked {
		t.Fatalf("expected developer trust score to block DX metrics discipline, got %#v", model)
	}

	model = activeDeveloperEcosystemVal0Model()
	model.DXMetrics.CertificationUse = true
	model = ComputeDeveloperEcosystemVal0Foundation(model)
	if model.DXMetricsState != DeveloperEcosystemVal0DXMetricsStateBlocked {
		t.Fatalf("expected certification use to block DX metrics discipline, got %#v", model)
	}

	model = activeDeveloperEcosystemVal0Model()
	model.DXMetrics.FastTrackApproval = true
	model = ComputeDeveloperEcosystemVal0Foundation(model)
	if model.DXMetricsState != DeveloperEcosystemVal0DXMetricsStateBlocked {
		t.Fatalf("expected fast-track approval to block DX metrics discipline, got %#v", model)
	}
}

func TestDeveloperEcosystemVal0PluginSafetyPerformanceBudgetReference(t *testing.T) {
	model := DeveloperEcosystemVal0PluginSafetyDisciplineModel()
	if model.PerformanceBudgetRef != DeveloperEcosystemVal0PerformanceBudgetDisciplineID {
		t.Fatalf("expected canonical performance budget ref %q, got %#v", DeveloperEcosystemVal0PerformanceBudgetDisciplineID, model)
	}

	active := activeDeveloperEcosystemVal0Model()
	if active.PluginSafety.PerformanceBudgetRef != DeveloperEcosystemVal0PerformanceBudgetDisciplineID ||
		active.PluginSafetyState != DeveloperEcosystemVal0PluginSafetyStateActive {
		t.Fatalf("expected active plugin safety with canonical performance budget ref, got %#v", active)
	}

	testCases := []struct {
		name     string
		ref      string
		expected string
	}{
		{name: "old dangling ref blocked", ref: "developer-performance-budget", expected: DeveloperEcosystemVal0PluginSafetyStateBlocked},
		{name: "empty ref incomplete", ref: "", expected: DeveloperEcosystemVal0PluginSafetyStateIncomplete},
		{name: "unknown mismatched ref blocked", ref: "developer-ecosystem-performance-budget-v2", expected: DeveloperEcosystemVal0PluginSafetyStateBlocked},
	}

	for _, tc := range testCases {
		mutated := activeDeveloperEcosystemVal0Model()
		mutated.PluginSafety.PerformanceBudgetRef = tc.ref
		mutated = ComputeDeveloperEcosystemVal0Foundation(mutated)
		if mutated.PluginSafetyState != tc.expected {
			t.Fatalf("expected %s to produce %s, got %#v", tc.name, tc.expected, mutated)
		}
		if mutated.CurrentState == DeveloperEcosystemVal0StateActive {
			t.Fatalf("expected %s to prevent active foundation state, got %#v", tc.name, mutated)
		}
	}

	blockers := []struct {
		name   string
		mutate func(*DeveloperEcosystemVal0Foundation)
	}{
		{name: "hidden policy override", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.PluginSafety.HiddenPolicyOverride = true
		}},
		{name: "governance bypass", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.PluginSafety.GovernanceBypass = true
		}},
		{name: "canonical truth claim", mutate: func(model *DeveloperEcosystemVal0Foundation) {
			model.PluginSafety.CanonicalTruthClaim = true
		}},
	}

	for _, tc := range blockers {
		mutated := activeDeveloperEcosystemVal0Model()
		tc.mutate(&mutated)
		mutated = ComputeDeveloperEcosystemVal0Foundation(mutated)
		if mutated.PluginSafetyState != DeveloperEcosystemVal0PluginSafetyStateBlocked {
			t.Fatalf("expected %s to block plugin safety state, got %#v", tc.name, mutated)
		}
	}
}

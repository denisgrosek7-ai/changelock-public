package operability

import (
	"encoding/json"
	"strings"
	"testing"
)

func activeDeploymentMultiTenantValBModel() DeploymentMultiTenantValBFoundation {
	model := DeploymentMultiTenantValBFoundationModel()
	return ComputeDeploymentMultiTenantValBFoundation(model)
}

func deploymentMultiTenantValBHasFinding(findings []DeploymentMultiTenantValBClosureBlockerFinding, level, surface, reason string) bool {
	for _, finding := range findings {
		if finding.BlockerLevel == level &&
			finding.Surface == surface &&
			finding.Reason == reason {
			return true
		}
	}
	return false
}

func assertDeploymentMultiTenantValBNoPoint10Pass(t *testing.T, model DeploymentMultiTenantValBFoundation) {
	t.Helper()
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal model: %v", err)
	}
	if strings.Contains(string(payload), "point_"+"10_pass") {
		t.Fatalf("expected Val B to never emit point 10 pass, got %s", string(payload))
	}
}

func TestDeploymentMultiTenantValBHappyPathAndPoint10NotComplete(t *testing.T) {
	model := activeDeploymentMultiTenantValBModel()
	if model.CurrentState != DeploymentMultiTenantValBStateActive {
		t.Fatalf("expected active Val B state, got %#v", model)
	}
	if model.ClosureBlockerState != DeploymentMultiTenantValBClosureBlockerStateActive {
		t.Fatalf("expected clean closure blocker overlay, got %#v", model)
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		t.Fatalf("expected point 10 to remain not complete, got %#v", model)
	}
	assertDeploymentMultiTenantValBNoPoint10Pass(t, model)
}

func TestDeploymentMultiTenantValBAggregateProjectionDisclaimerBlocks(t *testing.T) {
	model := activeDeploymentMultiTenantValBModel()
	model.ProjectionDisclaimer = "canonical_truth"
	model = ComputeDeploymentMultiTenantValBFoundation(model)
	if model.CurrentState != DeploymentMultiTenantValBStateBlocked {
		t.Fatalf("expected malformed aggregate projection disclaimer to block ValB state, got %#v", model)
	}
	if !containsTrimmedString(model.BlockingReasons, "aggregate_projection_disclaimer_blocked") {
		t.Fatalf("expected aggregate projection disclaimer blocking reason, got %#v", model.BlockingReasons)
	}
}

func TestDeploymentMultiTenantValBProjectionDisclaimerExactBoundedBlockers(t *testing.T) {
	testCases := []struct {
		name                string
		mutate              func(*DeploymentMultiTenantValBFoundation)
		wantDisciplineState string
	}{
		{
			name: "aggregate disclaimer suffix drift blocks exact state",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.ProjectionDisclaimer = deploymentMultiTenantValBProjectionDisclaimer() + " extra_suffix"
			},
		},
		{
			name: "aggregate disclaimer leading whitespace blocks exact state",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.ProjectionDisclaimer = " " + deploymentMultiTenantValBProjectionDisclaimer()
			},
		},
		{
			name: "aggregate disclaimer uppercase retagging blocks exact state",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.ProjectionDisclaimer = strings.ToUpper(deploymentMultiTenantValBProjectionDisclaimer())
			},
		},
		{
			name: "tenant isolation disclaimer prefix drift blocks exact discipline state",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.ProjectionDisclaimer = "prefix " + deploymentMultiTenantValBProjectionDisclaimer()
			},
			wantDisciplineState: DeploymentMultiTenantValBTenantIsolationStateBlocked,
		},
		{
			name: "no overclaim disclaimer suffix drift blocks exact discipline state",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.NoOverclaim.ProjectionDisclaimer = deploymentMultiTenantValBProjectionDisclaimer() + " extra_suffix"
			},
			wantDisciplineState: DeploymentMultiTenantValBNoOverclaimStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked ValB state, got %#v", tc.name, model)
		}
		switch tc.wantDisciplineState {
		case DeploymentMultiTenantValBTenantIsolationStateBlocked:
			if model.TenantIsolationState != tc.wantDisciplineState {
				t.Fatalf("%s: expected blocked tenant isolation state, got %#v", tc.name, model)
			}
		case DeploymentMultiTenantValBNoOverclaimStateBlocked:
			if model.NoOverclaimState != tc.wantDisciplineState {
				t.Fatalf("%s: expected blocked no-overclaim state, got %#v", tc.name, model)
			}
		}
		assertDeploymentMultiTenantValBNoPoint10Pass(t, model)
	}
}

func TestDeploymentMultiTenantValBDependencyCompatibleProjectionDisclaimersDoNotActivateFoundation(t *testing.T) {
	testCases := []struct {
		name       string
		disclaimer string
	}{
		{
			name:       "canonical disclaimer with aggregate suffix still blocks live foundation",
			disclaimer: deploymentMultiTenantValBProjectionDisclaimer() + " aggregate_dependency_snapshot",
		},
		{
			name:       "short aggregate disclaimer still blocks live foundation",
			disclaimer: "projection_only not_canonical_truth deployment_multi_tenant_valb aggregate_dependency_snapshot",
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		model.ProjectionDisclaimer = tc.disclaimer
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked ValB state, got %#v", tc.name, model)
		}
		if !containsTrimmedString(model.BlockingReasons, "aggregate_projection_disclaimer_blocked") {
			t.Fatalf("%s: expected aggregate projection disclaimer blocking reason, got %#v", tc.name, model.BlockingReasons)
		}
		assertDeploymentMultiTenantValBNoPoint10Pass(t, model)
	}
}

func TestDeploymentMultiTenantValBDependencyBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValBFoundation)
	}{
		{name: "vala current state partial blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValACurrentState = "partial"
		}},
		{name: "vala dependency state blocked blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValADependencyState = DeploymentMultiTenantValADependencyStateBlocked
		}},
		{name: "vala deployment profile matrix state blocked blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValADeploymentProfileMatrixState = DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
		}},
		{name: "vala preflight gate state blocked blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValAPreflightGateState = DeploymentMultiTenantValAPreflightGateStateBlocked
		}},
		{name: "vala identity bootstrap state blocked blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValAIdentityBootstrapState = DeploymentMultiTenantValAIdentityBootstrapStateBlocked
		}},
		{name: "vala air gapped evidence bundle state blocked blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValAAirGappedEvidenceBundleState = DeploymentMultiTenantValAAirGappedEvidenceBundleStateBlocked
		}},
		{name: "vala no overclaim state blocked blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValANoOverclaimState = DeploymentMultiTenantValANoOverclaimStateBlocked
		}},
		{name: "vala pass blocker state blocked blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValAPassBlockerState = DeploymentMultiTenantValAPassBlockerStateBlocked
		}},
		{name: "vala pass blocker state cleanup blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValAPassBlockerState = DeploymentMultiTenantValAPassBlockerStateCleanup
		}},
		{name: "point10 state complete blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.Point10State = "deployment_multi_tenant_point_10_complete"
		}},
		{name: "whitespace retagged aggregate dependency disclaimer blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ProjectionDisclaimer = " " + deploymentMultiTenantValAProjectionDisclaimer() + " aggregate_dependency_snapshot"
		}},
		{name: "uppercase aggregate dependency disclaimer blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ProjectionDisclaimer = strings.ToUpper(deploymentMultiTenantValAProjectionDisclaimer() + " aggregate_dependency_snapshot")
		}},
		{name: "malformed projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ProjectionDisclaimer = "canonical_truth"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.DependencyState != DeploymentMultiTenantValBDependencyStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked dependency, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValBDependencySnapshotCopiesAggregateComputedValAProjectionDisclaimer(t *testing.T) {
	valA := ComputeDeploymentMultiTenantValAFoundation(DeploymentMultiTenantValAFoundationModel())
	valA.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_vala aggregate_dependency_snapshot"
	valA.PassBlockerOverlay.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_vala component_pass_blocker"
	snapshot := DeploymentMultiTenantValBDependencySnapshot{
		ValACurrentState:                 valA.CurrentState,
		ValADependencyState:              valA.DependencyState,
		ValADeploymentProfileMatrixState: valA.DeploymentProfileMatrixState,
		ValAPreflightGateState:           valA.PreflightGateState,
		ValAIdentityBootstrapState:       valA.IdentityBootstrapState,
		ValAAirGappedEvidenceBundleState: valA.AirGappedEvidenceBundleState,
		ValANoOverclaimState:             valA.NoOverclaimState,
		ValAPassBlockerState:             valA.PassBlockerState,
		Point10State:                     valA.Point10State,
		ProjectionDisclaimer:             valA.ProjectionDisclaimer,
	}
	if snapshot.ProjectionDisclaimer != valA.ProjectionDisclaimer {
		t.Fatalf("expected aggregate ValA disclaimer to propagate exactly, got snapshot=%q vala=%q", snapshot.ProjectionDisclaimer, valA.ProjectionDisclaimer)
	}
	if snapshot.ProjectionDisclaimer == valA.PassBlockerOverlay.ProjectionDisclaimer {
		t.Fatalf("expected dependency snapshot not to fallback to component disclaimer, got snapshot=%q component=%q", snapshot.ProjectionDisclaimer, valA.PassBlockerOverlay.ProjectionDisclaimer)
	}
	if EvaluateDeploymentMultiTenantValBDependencyState(snapshot) != DeploymentMultiTenantValBDependencyStateActive {
		t.Fatalf("expected copied aggregate disclaimer to keep dependency active, got %#v", snapshot)
	}

	valA.ProjectionDisclaimer = "canonical_truth"
	snapshot.ProjectionDisclaimer = valA.ProjectionDisclaimer
	if EvaluateDeploymentMultiTenantValBDependencyState(snapshot) != DeploymentMultiTenantValBDependencyStateBlocked {
		t.Fatalf("expected malformed aggregate disclaimer to block dependency without component fallback, got %#v", snapshot)
	}
}

func TestDeploymentMultiTenantValBWhitespaceRetaggedDependencySnapshotBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValBFoundation)
	}{
		{name: "whitespace retagged vala current state blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValACurrentState = " " + DeploymentMultiTenantValAStateActive + " "
		}},
		{name: "tab retagged vala dependency state blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.ValADependencyState = "\t" + DeploymentMultiTenantValADependencyStateActive
		}},
		{name: "newline retagged point10 state blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.Dependency.Point10State = DeploymentMultiTenantPoint10StateNotComplete + "\n"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.DependencyState != DeploymentMultiTenantValBDependencyStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked dependency and ValB state, got %#v", tc.name, model)
		}
		assertDeploymentMultiTenantValBNoPoint10Pass(t, model)
	}
}

func TestDeploymentMultiTenantValBEvidenceRefsRequireRawExactBinding(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValBFoundation)
		assert func(*testing.T, DeploymentMultiTenantValBFoundation)
	}{
		{
			name: "tenant isolation leading whitespace evidence ref blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.EvidenceRefs[0] = " " + deploymentMultiTenantValBTenantIsolationEvidenceRefs()[0]
			},
			assert: func(t *testing.T, model DeploymentMultiTenantValBFoundation) {
				t.Helper()
				if model.TenantIsolationState != DeploymentMultiTenantValBTenantIsolationStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
					t.Fatalf("expected blocked tenant isolation state, got %#v", model)
				}
			},
		},
		{
			name: "data residency trailing whitespace evidence ref blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.EvidenceRefs[0] = deploymentMultiTenantValBDataResidencyEvidenceRefs()[0] + " "
			},
			assert: func(t *testing.T, model DeploymentMultiTenantValBFoundation) {
				t.Helper()
				if model.DataResidencyState != DeploymentMultiTenantValBDataResidencyStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
					t.Fatalf("expected blocked data residency state, got %#v", model)
				}
			},
		},
		{
			name: "tenant lifecycle tab padded evidence ref blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantLifecycle.EvidenceRefs[0] = "\t" + deploymentMultiTenantValBTenantLifecycleEvidenceRefs()[0]
			},
			assert: func(t *testing.T, model DeploymentMultiTenantValBFoundation) {
				t.Helper()
				if model.TenantLifecycleState != DeploymentMultiTenantValBTenantLifecycleStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
					t.Fatalf("expected blocked tenant lifecycle state, got %#v", model)
				}
			},
		},
		{
			name: "fair share newline padded evidence ref blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.FairShareQuota.EvidenceRefs[0] = deploymentMultiTenantValBFairShareEvidenceRefs()[0] + "\n"
			},
			assert: func(t *testing.T, model DeploymentMultiTenantValBFoundation) {
				t.Helper()
				if model.FairShareQuotaState != DeploymentMultiTenantValBFairShareQuotaStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
					t.Fatalf("expected blocked fair-share state, got %#v", model)
				}
			},
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		tc.assert(t, model)
		assertDeploymentMultiTenantValBNoPoint10Pass(t, model)
	}
}

func TestDeploymentMultiTenantValBTenantScopeRequiresRawExactBinding(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValBFoundation)
		assert func(*testing.T, DeploymentMultiTenantValBFoundation)
	}{
		{
			name: "tenant isolation bare lower-case tenant scope blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.TenantScope = "ops_scope"
			},
			assert: func(t *testing.T, model DeploymentMultiTenantValBFoundation) {
				t.Helper()
				if model.TenantIsolationState != DeploymentMultiTenantValBTenantIsolationStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
					t.Fatalf("expected blocked tenant isolation state, got %#v", model)
				}
			},
		},
		{
			name: "tenant isolation prefixed tenant scope blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.TenantScope = "ops " + deploymentMultiTenantVal0TenantScope()
			},
			assert: func(t *testing.T, model DeploymentMultiTenantValBFoundation) {
				t.Helper()
				if model.TenantIsolationState != DeploymentMultiTenantValBTenantIsolationStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
					t.Fatalf("expected blocked tenant isolation state, got %#v", model)
				}
			},
		},
		{
			name: "data residency suffixed tenant scope blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.TenantScope = deploymentMultiTenantVal0TenantScope() + " ops"
			},
			assert: func(t *testing.T, model DeploymentMultiTenantValBFoundation) {
				t.Helper()
				if model.DataResidencyState != DeploymentMultiTenantValBDataResidencyStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
					t.Fatalf("expected blocked data residency state, got %#v", model)
				}
			},
		},
		{
			name: "tenant lifecycle tab padded tenant scope blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantLifecycle.TenantScope = "\t" + deploymentMultiTenantVal0TenantScope()
			},
			assert: func(t *testing.T, model DeploymentMultiTenantValBFoundation) {
				t.Helper()
				if model.TenantLifecycleState != DeploymentMultiTenantValBTenantLifecycleStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
					t.Fatalf("expected blocked tenant lifecycle state, got %#v", model)
				}
			},
		},
		{
			name: "fair share newline padded tenant scope blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.FairShareQuota.TenantScope = deploymentMultiTenantVal0TenantScope() + "\n"
			},
			assert: func(t *testing.T, model DeploymentMultiTenantValBFoundation) {
				t.Helper()
				if model.FairShareQuotaState != DeploymentMultiTenantValBFairShareQuotaStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
					t.Fatalf("expected blocked fair-share state, got %#v", model)
				}
			},
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		tc.assert(t, model)
		assertDeploymentMultiTenantValBNoPoint10Pass(t, model)
	}
}

func TestDeploymentMultiTenantValBDataResidencyRegionConflictStateRequiresRawExactBinding(t *testing.T) {
	testCases := []struct {
		name  string
		value string
	}{
		{name: "leading whitespace conflict state blocks", value: " " + DeploymentMultiTenantConflictStateNoConflict},
		{name: "trailing whitespace conflict state blocks", value: DeploymentMultiTenantConflictStateNoConflict + " "},
		{name: "tab padded conflict state blocks", value: "\t" + DeploymentMultiTenantConflictStateNoConflict},
		{name: "newline padded conflict state blocks", value: DeploymentMultiTenantConflictStateNoConflict + "\n"},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		model.DataResidency.RegionConflictState = tc.value
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.DataResidencyState != DeploymentMultiTenantValBDataResidencyStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked data residency state, got %#v", tc.name, model)
		}
		assertDeploymentMultiTenantValBNoPoint10Pass(t, model)
	}
}

func TestDeploymentMultiTenantValBTenantIsolationBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValBFoundation)
	}{
		{name: "cross tenant audit leakage blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.CrossTenantAuditLeakagePresent = true
		}},
		{name: "cross tenant evidence leakage blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.CrossTenantEvidenceLeakagePresent = true
		}},
		{name: "cross tenant export leakage blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.CrossTenantExportLeakagePresent = true
		}},
		{name: "cross tenant credential leakage blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.CrossTenantCredentialLeakagePresent = true
		}},
		{name: "support operator access leakage blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.SupportOperatorAccessLeakagePresent = true
		}},
		{name: "tenant isolation config only blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantIsolationConfigOnly = true
		}},
		{name: "dashboard summary only blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.DashboardSummaryOnly = true
		}},
		{name: "fleet summary only blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.FleetSummaryOnly = true
		}},
		{name: "raw cross tenant evidence sharing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.RawCrossTenantEvidenceSharingPresent = true
		}},
		{name: "tenant private metadata side channel leakage blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantPrivateMetadataSideChannelLeakage = true
		}},
		{name: "missing tenant namespace isolation evidence blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantNamespaceIsolationEvidence = ""
		}},
		{name: "missing audit namespace evidence blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantScopedAuditNamespaceEvidence = ""
		}},
		{name: "missing evidence namespace evidence blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantScopedEvidenceNamespaceEvidence = ""
		}},
		{name: "missing export boundary evidence blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantScopedExportBoundaryEvidence = ""
		}},
		{name: "missing credential boundary evidence blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantScopedCredentialBoundaryEvidence = ""
		}},
		{name: "missing support operator boundary evidence blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantScopedSupportOperatorBoundaryEvidence = ""
		}},
		{name: "malformed tenant scope blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantScope = "malformed"
		}},
		{name: "unknown tenant scope blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantScope = "unknown"
		}},
		{name: "stale tenant scope blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantScope = "stale"
		}},
		{name: "wrong tenant scope blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantScope = "tenant:beta"
		}},
		{name: "duplicate tenant scope refs block", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantIsolation.TenantScope = "tenant:alpha tenant:beta"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.TenantIsolationState != DeploymentMultiTenantValBTenantIsolationStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked tenant isolation state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValBDataResidencyBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValBFoundation)
	}{
		{name: "data residency bypass blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.DataResidencyBypassPresent = true
		}},
		{name: "tenant region missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.TenantRegion = ""
		}},
		{name: "evidence region missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.EvidenceRegion = ""
		}},
		{name: "export region missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.ExportRegion = ""
		}},
		{name: "backup region missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.BackupRegionReference = ""
		}},
		{name: "support access region missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.SupportAccessRegion = ""
		}},
		{name: "malformed region blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.TenantRegion = "GLOBAL REGION"
		}},
		{name: "unknown region blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.TenantRegion = "unknown"
		}},
		{name: "stale region blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.TenantRegion = "stale"
		}},
		{name: "region export boundary missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.RegionExportBoundaryValidation = false
		}},
		{name: "cross region exception missing blocks when flow exists", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.CrossRegionFlowExists = true
			model.DataResidency.CrossRegionExceptionPath = ""
		}},
		{name: "cross region exception silently allowed blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.CrossRegionFlowExists = true
			model.DataResidency.CrossRegionExceptionPath = "bounded_cross_region_exception"
			model.DataResidency.CrossRegionExceptionSilentlyOpen = true
		}},
		{name: "backup path bypasses residency blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.BackupPathBypassesResidency = true
		}},
		{name: "export path bypasses residency blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.ExportPathBypassesResidency = true
		}},
		{name: "support path bypasses residency blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.SupportPathBypassesResidency = true
		}},
		{name: "region summary treated as canonical truth blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.RegionSummaryCanonicalTruth = true
		}},
		{name: "wrong tenant scope blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.DataResidency.TenantScope = "tenant:beta"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.DataResidencyState != DeploymentMultiTenantValBDataResidencyStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked data residency state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValBDataResidencyCrossRegionSemantics(t *testing.T) {
	testCases := []struct {
		name              string
		mutate            func(*DeploymentMultiTenantValBFoundation)
		wantResidency     string
		wantCurrent       string
		wantClosure       string
		wantFindingReason string
	}{
		{
			name: "same region happy path stays active",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.TenantRegion = "eu_central_1"
				model.DataResidency.EvidenceRegion = "eu_central_1"
				model.DataResidency.ExportRegion = "eu_central_1"
				model.DataResidency.BackupRegionReference = "eu_central_1"
				model.DataResidency.SupportAccessRegion = "eu_central_1"
				model.DataResidency.CrossRegionFlowExists = false
			},
			wantResidency: DeploymentMultiTenantValBDataResidencyStateActive,
			wantCurrent:   DeploymentMultiTenantValBStateActive,
			wantClosure:   DeploymentMultiTenantValBClosureBlockerStateActive,
		},
		{
			name: "evidence region mismatch without exception blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.TenantRegion = "tenant_region_eu"
				model.DataResidency.EvidenceRegion = "tenant_region_us"
				model.DataResidency.ExportRegion = "tenant_region_eu"
				model.DataResidency.BackupRegionReference = "tenant_region_eu"
				model.DataResidency.SupportAccessRegion = "tenant_region_eu"
				model.DataResidency.CrossRegionFlowExists = false
			},
			wantResidency:     DeploymentMultiTenantValBDataResidencyStateBlocked,
			wantCurrent:       DeploymentMultiTenantValBStateBlocked,
			wantClosure:       DeploymentMultiTenantValBClosureBlockerStateBlocked,
			wantFindingReason: "inferred cross-region residency flow without scoped audited exception",
		},
		{
			name: "export region mismatch without exception blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.TenantRegion = "tenant_region_eu"
				model.DataResidency.ExportRegion = "tenant_region_us"
				model.DataResidency.CrossRegionFlowExists = false
			},
			wantResidency:     DeploymentMultiTenantValBDataResidencyStateBlocked,
			wantCurrent:       DeploymentMultiTenantValBStateBlocked,
			wantClosure:       DeploymentMultiTenantValBClosureBlockerStateBlocked,
			wantFindingReason: "inferred cross-region residency flow without scoped audited exception",
		},
		{
			name: "backup region mismatch without exception blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.TenantRegion = "tenant_region_eu"
				model.DataResidency.BackupRegionReference = "tenant_region_us"
				model.DataResidency.CrossRegionFlowExists = false
			},
			wantResidency:     DeploymentMultiTenantValBDataResidencyStateBlocked,
			wantCurrent:       DeploymentMultiTenantValBStateBlocked,
			wantClosure:       DeploymentMultiTenantValBClosureBlockerStateBlocked,
			wantFindingReason: "inferred cross-region residency flow without scoped audited exception",
		},
		{
			name: "support region mismatch without exception blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.TenantRegion = "tenant_region_eu"
				model.DataResidency.SupportAccessRegion = "tenant_region_us"
				model.DataResidency.CrossRegionFlowExists = false
			},
			wantResidency:     DeploymentMultiTenantValBDataResidencyStateBlocked,
			wantCurrent:       DeploymentMultiTenantValBStateBlocked,
			wantClosure:       DeploymentMultiTenantValBClosureBlockerStateBlocked,
			wantFindingReason: "inferred cross-region residency flow without scoped audited exception",
		},
		{
			name: "region mismatch with valid exception may pass",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.TenantRegion = "tenant_region_eu"
				model.DataResidency.EvidenceRegion = "tenant_region_us"
				model.DataResidency.CrossRegionFlowExists = false
				model.DataResidency.CrossRegionExceptionPath = "bounded_cross_region_exception"
				model.DataResidency.CrossRegionExceptionScoped = true
				model.DataResidency.CrossRegionExceptionAudited = true
				model.DataResidency.CrossRegionExceptionSilentlyOpen = false
			},
			wantResidency: DeploymentMultiTenantValBDataResidencyStateActive,
			wantCurrent:   DeploymentMultiTenantValBStateActive,
			wantClosure:   DeploymentMultiTenantValBClosureBlockerStateActive,
		},
		{
			name: "explicit cross region flow still requires exception",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.CrossRegionFlowExists = true
				model.DataResidency.CrossRegionExceptionPath = ""
			},
			wantResidency: DeploymentMultiTenantValBDataResidencyStateBlocked,
			wantCurrent:   DeploymentMultiTenantValBStateBlocked,
			wantClosure:   DeploymentMultiTenantValBClosureBlockerStateBlocked,
		},
		{
			name: "explicit cross region flow with valid exception passes",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.CrossRegionFlowExists = true
				model.DataResidency.CrossRegionExceptionPath = "bounded_cross_region_exception"
				model.DataResidency.CrossRegionExceptionScoped = true
				model.DataResidency.CrossRegionExceptionAudited = true
				model.DataResidency.CrossRegionExceptionSilentlyOpen = false
			},
			wantResidency: DeploymentMultiTenantValBDataResidencyStateActive,
			wantCurrent:   DeploymentMultiTenantValBStateActive,
			wantClosure:   DeploymentMultiTenantValBClosureBlockerStateActive,
		},
		{
			name: "silently open exception still blocks",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.TenantRegion = "tenant_region_eu"
				model.DataResidency.EvidenceRegion = "tenant_region_us"
				model.DataResidency.CrossRegionExceptionPath = "bounded_cross_region_exception"
				model.DataResidency.CrossRegionExceptionScoped = true
				model.DataResidency.CrossRegionExceptionAudited = true
				model.DataResidency.CrossRegionExceptionSilentlyOpen = true
			},
			wantResidency:     DeploymentMultiTenantValBDataResidencyStateBlocked,
			wantCurrent:       DeploymentMultiTenantValBStateBlocked,
			wantClosure:       DeploymentMultiTenantValBClosureBlockerStateBlocked,
			wantFindingReason: "cross-region exception silently allowed",
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.DataResidencyState != tc.wantResidency {
			t.Fatalf("%s: expected data residency state %q, got %#v", tc.name, tc.wantResidency, model)
		}
		if model.CurrentState != tc.wantCurrent {
			t.Fatalf("%s: expected current state %q, got %#v", tc.name, tc.wantCurrent, model)
		}
		if model.ClosureBlockerState != tc.wantClosure {
			t.Fatalf("%s: expected closure blocker state %q, got %#v", tc.name, tc.wantClosure, model)
		}
		if tc.wantFindingReason != "" && !deploymentMultiTenantValBHasFinding(model.ClosureBlockerOverlay.Findings, DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceDataResidency, tc.wantFindingReason) {
			t.Fatalf("%s: expected CL-B0 data residency finding, got %#v", tc.name, model.ClosureBlockerOverlay)
		}
	}
}

func TestDeploymentMultiTenantValBDataResidencyInvalidCrossRegionExceptionPathsBlock(t *testing.T) {
	testCases := []string{"unknown", "partial", "incomplete", "stale", "malformed", "blocked"}
	for _, path := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		model.DataResidency.TenantRegion = "tenant_region_eu"
		model.DataResidency.EvidenceRegion = "tenant_region_us"
		model.DataResidency.CrossRegionExceptionPath = path
		model.DataResidency.CrossRegionExceptionScoped = true
		model.DataResidency.CrossRegionExceptionAudited = true
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.DataResidencyState != DeploymentMultiTenantValBDataResidencyStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked states, got %#v", path, model)
		}
		if !deploymentMultiTenantValBHasFinding(model.ClosureBlockerOverlay.Findings, DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceDataResidency, "inferred cross-region residency flow without scoped audited exception") {
			t.Fatalf("%s: expected inferred cross-region CL-B0 finding, got %#v", path, model.ClosureBlockerOverlay)
		}
	}
}

func TestDeploymentMultiTenantValBTenantLifecycleBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValBFoundation)
	}{
		{name: "tenant create missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.TenantCreatePresent = false
		}},
		{name: "tenant configure missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.TenantConfigurePresent = false
		}},
		{name: "tenant suspend missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.TenantSuspendPresent = false
		}},
		{name: "tenant transfer missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.TenantTransferPresent = false
		}},
		{name: "tenant offboard missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.TenantOffboardPresent = false
		}},
		{name: "tenant data export missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.TenantDataExportPresent = false
		}},
		{name: "tenant evidence retention missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.TenantEvidenceRetentionPresent = false
		}},
		{name: "tenant deletion missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.TenantDeletionPresent = false
		}},
		{name: "support access revoke missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.SupportAccessRevokePresent = false
		}},
		{name: "key custody rotation missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.KeyCustodyRotationPresent = false
		}},
		{name: "offboarding without support revoke blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.OffboardingRevokesSupportAccess = false
		}},
		{name: "transfer weakens boundaries blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.TransferPreservesBoundaries = false
		}},
		{name: "deletion export retention ambiguous blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.DeletionExportRetentionSemanticsExplicit = false
		}},
		{name: "lifecycle action without tenant scope blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.LifecycleActionTenantScoped = false
		}},
		{name: "lifecycle inferred from dashboard summary only blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.DashboardSummaryOnly = true
		}},
		{name: "wrong tenant scope blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.TenantLifecycle.TenantScope = "tenant:beta"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.TenantLifecycleState != DeploymentMultiTenantValBTenantLifecycleStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked tenant lifecycle state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValBFairShareQuotaBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValBFoundation)
	}{
		{name: "event budget missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.EventBudgetPerTenantPresent = false
		}},
		{name: "queue isolation missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.QueueIsolationPresent = false
		}},
		{name: "noisy tenant containment missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.NoisyTenantContainmentPresent = false
		}},
		{name: "alert flood throttling missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.AlertFloodThrottlingPresent = false
		}},
		{name: "no starvation rule missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.NoStarvationRulePresent = false
		}},
		{name: "overload downgrade semantics missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.OverloadDowngradeSemanticsPresent = false
		}},
		{name: "rate limit evidence missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.PerTenantRateLimitEvidencePresent = false
		}},
		{name: "backpressure semantics missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.PerTenantBackpressureSemanticsPresent = false
		}},
		{name: "tenant aware negative test missing blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.TenantAwareNegativeTestPresent = false
		}},
		{name: "one tenant starves another blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.OneTenantStarvesAnother = true
		}},
		{name: "noisy tenant degrades another tenant blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.NoisyTenantDegradesAnotherTenant = true
		}},
		{name: "alert flood spills across tenants blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.AlertFloodSpillsAcrossTenants = true
		}},
		{name: "overload silently ready blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.OverloadSilentlyTreatedAsReady = true
		}},
		{name: "global queue starvation without bounded degradation blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.GlobalQueueStarvationWithoutBoundedDegradation = true
		}},
		{name: "wrong tenant scope blocks", mutate: func(model *DeploymentMultiTenantValBFoundation) {
			model.FairShareQuota.TenantScope = "tenant:beta"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.FairShareQuotaState != DeploymentMultiTenantValBFairShareQuotaStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked fair-share quota state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValBClosureBlockerOverlayCLB0AndCLB1Blockers(t *testing.T) {
	testCases := []struct {
		name    string
		mutate  func(*DeploymentMultiTenantValBFoundation)
		level   string
		surface string
		reason  string
	}{
		{
			name: "cross tenant audit leakage produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.CrossTenantAuditLeakagePresent = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "cross-tenant audit leakage",
		},
		{
			name: "cross tenant evidence leakage produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.CrossTenantEvidenceLeakagePresent = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "cross-tenant evidence leakage",
		},
		{
			name: "cross tenant export leakage produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.CrossTenantExportLeakagePresent = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "cross-tenant export leakage",
		},
		{
			name: "cross tenant credential leakage produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.CrossTenantCredentialLeakagePresent = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "cross-tenant credential leakage",
		},
		{
			name: "support operator access leakage produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.SupportOperatorAccessLeakagePresent = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "support or operator access leakage",
		},
		{
			name: "raw cross tenant evidence sharing produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.RawCrossTenantEvidenceSharingPresent = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "raw cross-tenant evidence sharing",
		},
		{
			name: "tenant private metadata side channel marked safe produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.TenantPrivateMetadataSideChannelMarkedSafe = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "tenant-private metadata side-channel leakage marked safe",
		},
		{
			name: "tenant isolation config only produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.TenantIsolationConfigOnly = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "tenant isolation treated as config-only",
		},
		{
			name: "dashboard summary canonical isolation evidence produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.DashboardSummaryOnly = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "dashboard summary treated as canonical isolation evidence",
		},
		{
			name: "fleet summary canonical isolation evidence produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.FleetSummaryOnly = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "fleet summary treated as canonical isolation evidence",
		},
		{
			name: "data residency bypass produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.DataResidencyBypassPresent = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceDataResidency, reason: "data residency bypass",
		},
		{
			name: "backup residency bypass produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.BackupPathBypassesResidency = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceDataResidency, reason: "backup path bypasses data residency",
		},
		{
			name: "export residency bypass produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.ExportPathBypassesResidency = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceDataResidency, reason: "export path bypasses data residency",
		},
		{
			name: "support residency bypass produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.SupportPathBypassesResidency = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceDataResidency, reason: "support path bypasses data residency",
		},
		{
			name: "region summary canonical truth produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.RegionSummaryCanonicalTruth = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceDataResidency, reason: "region summary treated as canonical truth",
		},
		{
			name: "cross region exception silently allowed produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.CrossRegionFlowExists = true
				model.DataResidency.CrossRegionExceptionSilentlyOpen = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceDataResidency, reason: "cross-region exception silently allowed",
		},
		{
			name: "canonical evidence spine bypass produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.CanonicalEvidenceSpineBypass = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "canonical evidence spine bypass",
		},
		{
			name: "clean room violation produces cl b0 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.NoOverclaim.CleanRoomIPViolationDetected = true
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB0, surface: DeploymentMultiTenantValBClosureSurfaceCleanRoomIP, reason: "copied competitor deployment or tenant isolation artifact detected",
		},
		{
			name: "fair share negative test missing produces cl b1 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.FairShareQuota.TenantAwareNegativeTestPresent = false
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB1, surface: DeploymentMultiTenantValBClosureSurfaceFairShare, reason: "fair-share or quota policy lacks tenant-aware negative test",
		},
		{
			name: "tenant scope negative test missing produces cl b1 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.TenantScopeNegativeTestPresent = false
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB1, surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation, reason: "malformed unknown or stale tenant scope not tested",
		},
		{
			name: "tenant lifecycle revoke semantics missing produces cl b1 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantLifecycle.TenantOffboardPresent = false
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB1, surface: DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, reason: "tenant lifecycle lacks offboarding or revoke semantics",
		},
		{
			name: "support access revoke missing produces cl b1 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantLifecycle.SupportAccessRevokePresent = false
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB1, surface: DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, reason: "tenant lifecycle lacks offboarding or revoke semantics",
		},
		{
			name: "offboarding without revoke produces cl b1 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantLifecycle.OffboardingRevokesSupportAccess = false
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB1, surface: DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, reason: "offboarding does not revoke support or operator access",
		},
		{
			name: "tenant transfer weakens boundaries produces cl b1 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantLifecycle.TransferPreservesBoundaries = false
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB1, surface: DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, reason: "tenant transfer weakens audit evidence or export boundaries",
		},
		{
			name: "deletion export retention ambiguity produces cl b1 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantLifecycle.DeletionExportRetentionSemanticsExplicit = false
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB1, surface: DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, reason: "deletion export or retention semantics ambiguous",
		},
		{
			name: "cross region exception path missing produces cl b1 blocker",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.CrossRegionFlowExists = true
				model.DataResidency.CrossRegionExceptionPath = ""
			},
			level: DeploymentMultiTenantValBBlockerLevelCLB1, surface: DeploymentMultiTenantValBClosureSurfaceDataResidency, reason: "data residency exception path missing while cross-region flow exists",
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.ClosureBlockerState != DeploymentMultiTenantValBClosureBlockerStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected blocked closure blocker state, got %#v", tc.name, model)
		}
		if !deploymentMultiTenantValBHasFinding(model.ClosureBlockerOverlay.Findings, tc.level, tc.surface, tc.reason) {
			t.Fatalf("%s: expected closure blocker finding, got %#v", tc.name, model.ClosureBlockerOverlay)
		}
	}
}

func TestDeploymentMultiTenantValBClosureBlockerOverlayCLB2Cleanup(t *testing.T) {
	testCases := []struct {
		name    string
		mutate  func(*DeploymentMultiTenantValBFoundation)
		surface string
		reason  string
	}{
		{
			name: "ambiguous region naming is cleanup",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.RegionNamingExact = false
			},
			surface: DeploymentMultiTenantValBClosureSurfaceDataResidency,
			reason:  "ambiguous region naming",
		},
		{
			name: "ambiguous deployment tenant profile naming is cleanup",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.TenantProfileNamingExact = false
			},
			surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation,
			reason:  "ambiguous deployment or tenant profile naming",
		},
		{
			name: "missing safe wording example for tenant isolation is cleanup",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.SafeIsolationWordingExamplePresent = false
			},
			surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation,
			reason:  "missing safe wording example for tenant isolation",
		},
		{
			name: "missing safe wording example for data residency is cleanup",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.SafeResidencyWordingExample = false
			},
			surface: DeploymentMultiTenantValBClosureSurfaceDataResidency,
			reason:  "missing safe wording example for data residency",
		},
		{
			name: "incomplete diagnostic output for tenant isolation is cleanup",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantIsolation.DiagnosticOutputComplete = false
			},
			surface: DeploymentMultiTenantValBClosureSurfaceTenantIsolation,
			reason:  "incomplete diagnostic output for tenant isolation blockers",
		},
		{
			name: "incomplete diagnostic output for data residency is cleanup",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.DataResidency.DiagnosticOutputComplete = false
			},
			surface: DeploymentMultiTenantValBClosureSurfaceDataResidency,
			reason:  "incomplete diagnostic output for data residency blockers",
		},
		{
			name: "incomplete diagnostic output for lifecycle is cleanup",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.TenantLifecycle.DiagnosticOutputComplete = false
			},
			surface: DeploymentMultiTenantValBClosureSurfaceTenantLifecycle,
			reason:  "incomplete diagnostic output for tenant lifecycle blockers",
		},
		{
			name: "incomplete diagnostic output for fair share is cleanup",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.FairShareQuota.DiagnosticOutputComplete = false
			},
			surface: DeploymentMultiTenantValBClosureSurfaceFairShare,
			reason:  "incomplete diagnostic output for fair-share blockers",
		},
		{
			name: "incomplete runbook wording is cleanup",
			mutate: func(model *DeploymentMultiTenantValBFoundation) {
				model.FairShareQuota.RunbookWordingComplete = false
			},
			surface: DeploymentMultiTenantValBClosureSurfaceFairShare,
			reason:  "incomplete runbook wording without direct closure bypass",
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValBModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.ClosureBlockerState != DeploymentMultiTenantValBClosureBlockerStateCleanup {
			t.Fatalf("%s: expected cleanup closure blocker state, got %#v", tc.name, model)
		}
		if model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("%s: expected cleanup to block final Val B active state, got %#v", tc.name, model)
		}
		if !deploymentMultiTenantValBHasFinding(model.ClosureBlockerOverlay.Findings, DeploymentMultiTenantValBBlockerLevelCLB2, tc.surface, tc.reason) {
			t.Fatalf("%s: expected cleanup finding, got %#v", tc.name, model.ClosureBlockerOverlay)
		}
	}
}

func TestDeploymentMultiTenantValBClosureBlockerFindingReasonMustMatchExactly(t *testing.T) {
	finding := DeploymentMultiTenantValBClosureBlockerFinding{
		BlockerLevel:      DeploymentMultiTenantValBBlockerLevelCLB2,
		Surface:           DeploymentMultiTenantValBClosureSurfaceDataResidency,
		Reason:            "ambiguous region naming / extra context",
		BlocksCurrentWave: true,
		RequiredFollowup:  "normalize region naming",
	}
	if deploymentMultiTenantValBHasFinding([]DeploymentMultiTenantValBClosureBlockerFinding{finding}, DeploymentMultiTenantValBBlockerLevelCLB2, DeploymentMultiTenantValBClosureSurfaceDataResidency, "ambiguous region naming") {
		t.Fatalf("expected exact reason match only, got %#v", finding)
	}
}

func TestDeploymentMultiTenantValBClosureBlockerOverlayRejectsLegacyAndUnknownLevels(t *testing.T) {
	testCases := []struct {
		name    string
		finding DeploymentMultiTenantValBClosureBlockerFinding
	}{
		{
			name: "legacy priority zero is rejected",
			finding: DeploymentMultiTenantValBClosureBlockerFinding{
				BlockerLevel:      deploymentMultiTenantLegacyPriority("0"),
				Surface:           DeploymentMultiTenantValBClosureSurfaceTenantIsolation,
				Reason:            "legacy severity rejected",
				BlocksCurrentWave: true,
			},
		},
		{
			name: "legacy priority one is rejected",
			finding: DeploymentMultiTenantValBClosureBlockerFinding{
				BlockerLevel:      deploymentMultiTenantLegacyPriority("1"),
				Surface:           DeploymentMultiTenantValBClosureSurfaceTenantIsolation,
				Reason:            "legacy severity rejected",
				BlocksCurrentWave: true,
				RequiredFollowup:  "use cl b blocker terminology",
			},
		},
		{
			name: "legacy priority two is rejected",
			finding: DeploymentMultiTenantValBClosureBlockerFinding{
				BlockerLevel:      deploymentMultiTenantLegacyPriority("2"),
				Surface:           DeploymentMultiTenantValBClosureSurfaceTenantIsolation,
				Reason:            "legacy severity rejected",
				BlocksCurrentWave: true,
				RequiredFollowup:  "use cl b blocker terminology",
			},
		},
		{
			name: "unknown blocker level is rejected",
			finding: DeploymentMultiTenantValBClosureBlockerFinding{
				BlockerLevel:      "CL-B9",
				Surface:           DeploymentMultiTenantValBClosureSurfaceTenantIsolation,
				Reason:            "unknown blocker level",
				BlocksCurrentWave: true,
			},
		},
		{
			name: "unknown surface is rejected",
			finding: DeploymentMultiTenantValBClosureBlockerFinding{
				BlockerLevel:      DeploymentMultiTenantValBBlockerLevelCLB0,
				Surface:           "unknown_surface",
				Reason:            "unknown surface",
				BlocksCurrentWave: true,
			},
		},
	}

	for _, tc := range testCases {
		state := EvaluateDeploymentMultiTenantValBClosureBlockerState(DeploymentMultiTenantValBClosureBlockerOverlay{
			ProjectionDisclaimer: deploymentMultiTenantValBProjectionDisclaimer(),
			Findings:             []DeploymentMultiTenantValBClosureBlockerFinding{tc.finding},
		})
		if state != DeploymentMultiTenantValBClosureBlockerStateBlocked {
			t.Fatalf("%s: expected blocked state, got %q", tc.name, state)
		}
	}
}

func TestDeploymentMultiTenantValBClosureBlockerOverlayCLB3Advisory(t *testing.T) {
	advisory := DeploymentMultiTenantValBClosureBlockerOverlay{
		ProjectionDisclaimer: deploymentMultiTenantValBProjectionDisclaimer(),
		Findings: []DeploymentMultiTenantValBClosureBlockerFinding{
			{
				BlockerLevel:      DeploymentMultiTenantValBBlockerLevelCLB3,
				Surface:           DeploymentMultiTenantValBClosureSurfaceFairShare,
				Reason:            "advisory cleanup carried forward",
				BlocksCurrentWave: false,
				RequiredFollowup:  "record advisory cleanup if it is carried forward",
			},
		},
	}
	if state := EvaluateDeploymentMultiTenantValBClosureBlockerState(advisory); state != DeploymentMultiTenantValBClosureBlockerStateAdvisory {
		t.Fatalf("expected advisory closure blocker state, got %q", state)
	}

	mixed := DeploymentMultiTenantValBClosureBlockerOverlay{
		ProjectionDisclaimer: deploymentMultiTenantValBProjectionDisclaimer(),
		Findings: []DeploymentMultiTenantValBClosureBlockerFinding{
			{
				BlockerLevel:      DeploymentMultiTenantValBBlockerLevelCLB3,
				Surface:           DeploymentMultiTenantValBClosureSurfaceFairShare,
				Reason:            "advisory cleanup carried forward",
				BlocksCurrentWave: false,
				RequiredFollowup:  "record advisory cleanup if it is carried forward",
			},
			{
				BlockerLevel:      DeploymentMultiTenantValBBlockerLevelCLB0,
				Surface:           DeploymentMultiTenantValBClosureSurfaceTenantIsolation,
				Reason:            "cross-tenant audit leakage",
				BlocksCurrentWave: true,
			},
		},
	}
	if state := EvaluateDeploymentMultiTenantValBClosureBlockerState(mixed); state != DeploymentMultiTenantValBClosureBlockerStateBlocked {
		t.Fatalf("expected advisory finding to not mask stronger blocker, got %q", state)
	}
}

func TestDeploymentMultiTenantValBNoOverclaimBlockers(t *testing.T) {
	blockedClaims := []string{
		"tenant isolation guaranteed",
		"tenant isolation evidence guaranteed",
		"zero cross-tenant leakage",
		"no leakage guaranteed",
		"data residency certified",
		"data residency guaranteed",
		"region compliant by default",
		"sovereign compliant by default",
		"all tenants isolated by default",
		"tenant safe by default",
		"cross-tenant safe by default",
		"support access cannot leak",
		"backup residency guaranteed",
		"export residency guaranteed",
		"fair-share guarantees no outages",
		"quotas guarantee tenant performance",
		"noisy tenant cannot impact anyone",
		"lifecycle complete means compliant",
		"offboarding guarantees deletion",
		"deletion guaranteed",
		"transfer safe by default",
		"dashboard proves tenant isolation",
		"fleet view proves data residency",
		"fleet view evidence proves data residency",
		"support summary is canonical truth",
		"support summary evidence is canonical truth",
		"region summary is canonical truth",
		"evidence-linked tenant isolation test tenant isolation guaranteed",
		"data residency evidence data residency certified",
		"fair-share quota evidence quotas guarantee tenant performance",
		"advisory fleet visibility fleet view proves data residency",
		"support access revoke evidence support access cannot leak",
		"region/export boundary validation region compliant by default",
	}

	for _, claim := range blockedClaims {
		model := activeDeploymentMultiTenantValBModel()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValBNoOverclaimStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("expected forbidden claim %q to block, got %#v", claim, model)
		}
	}

	splitBlockedClaims := [][]string{
		{"tenant isolation", "guaranteed"},
		{"fleet view", "proves data residency"},
		{"support summary", "is canonical truth"},
	}
	for _, claims := range splitBlockedClaims {
		model := activeDeploymentMultiTenantValBModel()
		model.NoOverclaim.ObservedClaims = claims
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValBNoOverclaimStateBlocked || model.CurrentState != DeploymentMultiTenantValBStateBlocked {
			t.Fatalf("expected split forbidden claims %q to block, got %#v", claims, model)
		}
	}

	allowedClaims := []string{
		"evidence-linked tenant isolation test",
		"tenant-scoped audit boundary",
		"tenant-scoped evidence boundary",
		"tenant-scoped export boundary",
		"tenant-scoped credential boundary",
		"data residency evidence",
		"region/export boundary validation",
		"bounded cross-region exception path",
		"tenant lifecycle evidence",
		"support access revoke evidence",
		"key/custody rotation evidence",
		"fair-share quota evidence",
		"tenant-aware negative test",
		"noisy tenant containment evidence",
		"bounded degradation semantics",
		"not production approval",
		"not compliance certification",
		"not canonical truth",
		"advisory fleet visibility",
	}

	for _, claim := range allowedClaims {
		model := activeDeploymentMultiTenantValBModel()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantValBFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValBNoOverclaimStateActive || model.CurrentState != DeploymentMultiTenantValBStateActive {
			t.Fatalf("expected allowed bounded wording for %q, got %#v", claim, model)
		}
	}
}

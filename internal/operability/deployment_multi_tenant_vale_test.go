package operability

import (
	"encoding/json"
	"strings"
	"testing"
)

func activeDeploymentMultiTenantValEModel() DeploymentMultiTenantValEFoundation {
	return ComputeDeploymentMultiTenantValEFoundation(DeploymentMultiTenantValEFoundationModel())
}

func mustMarshalDeploymentMultiTenantValEJSON(t *testing.T, value any) string {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	return string(data)
}

func deploymentMultiTenantValEExactReasonPresent(reasons []string, expected string) bool {
	for _, reason := range reasons {
		if reason == expected {
			return true
		}
	}
	return false
}

func TestDeploymentMultiTenantValEHappyPathFinalPass(t *testing.T) {
	model := activeDeploymentMultiTenantValEModel()
	if model.CurrentState != DeploymentMultiTenantValEStatePass {
		t.Fatalf("expected Val E pass state, got %#v", model)
	}
	if model.Point10State != DeploymentMultiTenantPoint10StatePass {
		t.Fatalf("expected final point_10_pass state, got %#v", model)
	}
	if model.PassClosureManifest.ReviewerResult != DeploymentMultiTenantValEReviewerResultPassConfirmed {
		t.Fatalf("expected PASS_CONFIRMED reviewer result, got %#v", model.PassClosureManifest)
	}
	if len(model.CLBClosureLedger.CLB0OpenFindings) != 0 || len(model.CLBClosureLedger.CLB1OpenFindings) != 0 || len(model.CLBClosureLedger.CLB2OpenFindings) != 0 {
		t.Fatalf("expected no open CL-B0/1/2 findings, got %#v", model.CLBClosureLedger)
	}
	jsonBody := mustMarshalDeploymentMultiTenantValEJSON(t, model)
	if !strings.Contains(jsonBody, DeploymentMultiTenantPoint10StatePass) {
		t.Fatalf("expected final JSON to contain point_10_pass, got %s", jsonBody)
	}
}

func TestDeploymentMultiTenantValEDependencySnapshotCopiesComputedUpstreamOutput(t *testing.T) {
	val0 := ComputeDeploymentMultiTenantVal0Foundation(DeploymentMultiTenantVal0FoundationModel())
	val0.ProjectionDisclaimer = "canonical_truth"
	val0.NoOverclaim.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_val0 component_no_overclaim"
	val0.CurrentState = "unknown"
	val0.NoOverclaimState = "blocked"
	val0.Point10State = "deployment_multi_tenant_point_10_complete"
	val0Snapshot := deploymentMultiTenantValEVal0DependencySnapshotFromComputed(val0)
	if val0Snapshot.ProjectionDisclaimer != "canonical_truth" || val0Snapshot.CurrentState != "unknown" || val0Snapshot.NoOverclaimState != "blocked" || val0Snapshot.Point10State != "deployment_multi_tenant_point_10_complete" {
		t.Fatalf("expected val0 snapshot to copy computed output, got %#v", val0Snapshot)
	}
	if val0Snapshot.ProjectionDisclaimer == val0.NoOverclaim.ProjectionDisclaimer {
		t.Fatalf("expected val0 snapshot not to fallback to component disclaimer, got snapshot=%q component=%q", val0Snapshot.ProjectionDisclaimer, val0.NoOverclaim.ProjectionDisclaimer)
	}

	valA := ComputeDeploymentMultiTenantValAFoundation(DeploymentMultiTenantValAFoundationModel())
	valA.ProjectionDisclaimer = ""
	valA.PassBlockerOverlay.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_vala component_pass_blocker"
	valA.PassBlockerState = "blocked"
	valASnapshot := deploymentMultiTenantValEValADependencySnapshotFromComputed(valA)
	if valASnapshot.ProjectionDisclaimer != "" || valASnapshot.PassBlockerState != "blocked" {
		t.Fatalf("expected vala snapshot to copy computed output, got %#v", valASnapshot)
	}
	if valASnapshot.ProjectionDisclaimer == valA.PassBlockerOverlay.ProjectionDisclaimer {
		t.Fatalf("expected vala snapshot not to fallback to component disclaimer, got snapshot=%q component=%q", valASnapshot.ProjectionDisclaimer, valA.PassBlockerOverlay.ProjectionDisclaimer)
	}

	valB := ComputeDeploymentMultiTenantValBFoundation(DeploymentMultiTenantValBFoundationModel())
	valB.ProjectionDisclaimer = "unsupported"
	valB.ClosureBlockerOverlay.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_valb component_closure_blocker"
	valB.ClosureBlockerState = DeploymentMultiTenantValBClosureBlockerStateCleanup
	valBSnapshot := deploymentMultiTenantValEValBDependencySnapshotFromComputed(valB)
	if valBSnapshot.ProjectionDisclaimer != "unsupported" || valBSnapshot.ClosureBlockerState != DeploymentMultiTenantValBClosureBlockerStateCleanup {
		t.Fatalf("expected valb snapshot to copy computed output, got %#v", valBSnapshot)
	}
	if valBSnapshot.ProjectionDisclaimer == valB.ClosureBlockerOverlay.ProjectionDisclaimer {
		t.Fatalf("expected valb snapshot not to fallback to component disclaimer, got snapshot=%q component=%q", valBSnapshot.ProjectionDisclaimer, valB.ClosureBlockerOverlay.ProjectionDisclaimer)
	}

	valC := ComputeDeploymentMultiTenantValCFoundation(DeploymentMultiTenantValCFoundationModel())
	valC.ProjectionDisclaimer = "blocked"
	valC.ClosureBlockerOverlay.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_valc component_closure_blocker"
	valC.HAReadinessState = "unknown"
	valCSnapshot := deploymentMultiTenantValEValCDependencySnapshotFromComputed(valC)
	if valCSnapshot.ProjectionDisclaimer != "blocked" || valCSnapshot.HAReadinessState != "unknown" {
		t.Fatalf("expected valc snapshot to copy computed output, got %#v", valCSnapshot)
	}
	if valCSnapshot.ProjectionDisclaimer == valC.ClosureBlockerOverlay.ProjectionDisclaimer {
		t.Fatalf("expected valc snapshot not to fallback to component disclaimer, got snapshot=%q component=%q", valCSnapshot.ProjectionDisclaimer, valC.ClosureBlockerOverlay.ProjectionDisclaimer)
	}

	valD := ComputeDeploymentMultiTenantValDFoundation(DeploymentMultiTenantValDFoundationModel())
	valD.ProjectionDisclaimer = "production_approval"
	valD.ClosureBlockerOverlay.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_vald component_closure_blocker"
	valD.NoOverclaimState = "blocked"
	valDSnapshot := deploymentMultiTenantValEValDDependencySnapshotFromComputed(valD)
	if valDSnapshot.ProjectionDisclaimer != "production_approval" || valDSnapshot.NoOverclaimState != "blocked" {
		t.Fatalf("expected vald snapshot to copy computed output, got %#v", valDSnapshot)
	}
	if valDSnapshot.ProjectionDisclaimer == valD.ClosureBlockerOverlay.ProjectionDisclaimer {
		t.Fatalf("expected vald snapshot not to fallback to component disclaimer, got snapshot=%q component=%q", valDSnapshot.ProjectionDisclaimer, valD.ClosureBlockerOverlay.ProjectionDisclaimer)
	}
}

func TestDeploymentMultiTenantValEDependencyBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValEFoundation)
	}{
		{name: "val0 current state blocked", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.Val0.CurrentState = "blocked"
		}},
		{name: "val0 dependency state partial", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.Val0.DependencyState = "partial"
		}},
		{name: "val0 deployment validation unknown", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.Val0.DeploymentValidationState = "unknown"
		}},
		{name: "val0 point10 pass blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.Val0.Point10State = DeploymentMultiTenantPoint10StatePass
		}},
		{name: "vala current state blocked", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValA.CurrentState = "blocked"
		}},
		{name: "vala pass blocker unknown", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValA.PassBlockerState = "unknown"
		}},
		{name: "valb closure cleanup blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValB.ClosureBlockerState = DeploymentMultiTenantValBClosureBlockerStateCleanup
		}},
		{name: "valb closure advisory blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValB.ClosureBlockerState = DeploymentMultiTenantValBClosureBlockerStateAdvisory
		}},
		{name: "valc closure cleanup blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValC.ClosureBlockerState = DeploymentMultiTenantValCClosureBlockerStateCleanup
		}},
		{name: "valc privacy guard blocked", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValC.PrivacyGuardState = "blocked"
		}},
		{name: "vald closure advisory blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValD.ClosureBlockerState = DeploymentMultiTenantValDClosureBlockerStateAdvisory
		}},
		{name: "vald operator action blocked", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValD.OperatorActionState = "blocked"
		}},
		{name: "invalid val0 disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.Val0.ProjectionDisclaimer = ""
		}},
		{name: "invalid vala disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValA.ProjectionDisclaimer = "canonical_truth"
		}},
		{name: "invalid valb disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValB.ProjectionDisclaimer = "blocked"
		}},
		{name: "invalid valc disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValC.ProjectionDisclaimer = "unknown"
		}},
		{name: "invalid vald disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValD.ProjectionDisclaimer = "production_approval"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.DependencyState != DeploymentMultiTenantValEDependencyStateBlocked {
			t.Fatalf("%s: expected blocked dependency state, got %#v", tc.name, model)
		}
		if model.CurrentState != DeploymentMultiTenantValEStateBlocked {
			t.Fatalf("%s: expected blocked final state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValEWhitespaceRetaggedDependencySnapshotBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValEFoundation)
	}{
		{name: "whitespace retagged val0 current state blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.Val0.CurrentState = " " + DeploymentMultiTenantVal0StateActive + " "
		}},
		{name: "tab retagged vala point10 state blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValA.Point10State = "\t" + DeploymentMultiTenantPoint10StateNotComplete
		}},
		{name: "newline retagged valb closure blocker state blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValB.ClosureBlockerState = DeploymentMultiTenantValBClosureBlockerStateActive + "\n"
		}},
		{name: "whitespace retagged valc privacy guard state blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValC.PrivacyGuardState = " " + DeploymentMultiTenantValCPrivacyGuardStateActive + " "
		}},
		{name: "newline retagged vald no-overclaim state blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValD.NoOverclaimState = DeploymentMultiTenantValDNoOverclaimStateActive + "\n"
		}},
		{name: "whitespace retagged val0 projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.Val0.ProjectionDisclaimer = " " + deploymentMultiTenantVal0ProjectionDisclaimer() + " "
		}},
		{name: "tab retagged vala projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValA.ProjectionDisclaimer = "\t" + deploymentMultiTenantValAProjectionDisclaimer()
		}},
		{name: "newline retagged valb projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValB.ProjectionDisclaimer = deploymentMultiTenantValBProjectionDisclaimer() + "\n"
		}},
		{name: "whitespace retagged valc projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValC.ProjectionDisclaimer = " " + deploymentMultiTenantValCProjectionDisclaimer()
		}},
		{name: "tab newline retagged vald projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.Dependency.ValD.ProjectionDisclaimer = "\t" + deploymentMultiTenantValDProjectionDisclaimer() + "\n"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.DependencyState != DeploymentMultiTenantValEDependencyStateBlocked {
			t.Fatalf("%s: expected blocked dependency state, got %#v", tc.name, model)
		}
		if model.CurrentState != DeploymentMultiTenantValEStateBlocked {
			t.Fatalf("%s: expected blocked final state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValEFoundationProjectionDisclaimerExactBoundedBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*DeploymentMultiTenantValEFoundation)
		stateName string
		want      string
		state     func(DeploymentMultiTenantValEFoundation) string
	}{
		{
			name: "integrated invariant leading whitespace disclaimer blocks",
			mutate: func(model *DeploymentMultiTenantValEFoundation) {
				model.IntegratedInvariantReview.ProjectionDisclaimer = " " + deploymentMultiTenantValEProjectionDisclaimer()
			},
			stateName: "integrated invariant",
			want:      DeploymentMultiTenantValEIntegratedInvariantStateBlocked,
			state: func(model DeploymentMultiTenantValEFoundation) string {
				return model.IntegratedInvariantState
			},
		},
		{
			name: "evidence quality aggregate disclaimer blocks",
			mutate: func(model *DeploymentMultiTenantValEFoundation) {
				model.EvidenceQualityMap.ProjectionDisclaimer = deploymentMultiTenantValEProjectionDisclaimer() + " aggregate_dependency_snapshot"
			},
			stateName: "evidence quality",
			want:      DeploymentMultiTenantValEEvidenceQualityStateBlocked,
			state: func(model DeploymentMultiTenantValEFoundation) string {
				return model.EvidenceQualityState
			},
		},
		{
			name: "clb closure uppercase disclaimer blocks",
			mutate: func(model *DeploymentMultiTenantValEFoundation) {
				model.CLBClosureLedger.ProjectionDisclaimer = strings.ToUpper(deploymentMultiTenantValEProjectionDisclaimer())
			},
			stateName: "clb closure",
			want:      DeploymentMultiTenantValECLBClosureStateBlocked,
			state: func(model DeploymentMultiTenantValEFoundation) string {
				return model.CLBClosureState
			},
		},
		{
			name: "manifest tab-padded disclaimer blocks",
			mutate: func(model *DeploymentMultiTenantValEFoundation) {
				model.PassClosureManifest.ProjectionDisclaimer = "\t" + deploymentMultiTenantValEProjectionDisclaimer()
			},
			stateName: "pass closure manifest",
			want:      DeploymentMultiTenantValEPassClosureManifestStateBlocked,
			state: func(model DeploymentMultiTenantValEFoundation) string {
				return model.PassClosureManifestState
			},
		},
		{
			name: "no overclaim aggregate disclaimer blocks",
			mutate: func(model *DeploymentMultiTenantValEFoundation) {
				model.NoOverclaim.ProjectionDisclaimer = deploymentMultiTenantValEProjectionDisclaimer() + " aggregate_dependency_snapshot"
			},
			stateName: "no overclaim",
			want:      DeploymentMultiTenantValENoOverclaimStateBlocked,
			state: func(model DeploymentMultiTenantValEFoundation) string {
				return model.NoOverclaimState
			},
		},
		{
			name: "projection boundary leading whitespace disclaimer blocks",
			mutate: func(model *DeploymentMultiTenantValEFoundation) {
				model.ProjectionBoundaryReview.ProjectionDisclaimer = " " + deploymentMultiTenantValEProjectionDisclaimer()
			},
			stateName: "projection boundary",
			want:      DeploymentMultiTenantValEProjectionBoundaryStateBlocked,
			state: func(model DeploymentMultiTenantValEFoundation) string {
				return model.ProjectionBoundaryState
			},
		},
		{
			name: "projection boundary surface uppercase disclaimer blocks",
			mutate: func(model *DeploymentMultiTenantValEFoundation) {
				model.ProjectionBoundaryReview.Surfaces[0].Disclaimer = strings.ToUpper(deploymentMultiTenantValEProjectionDisclaimer())
			},
			stateName: "projection boundary",
			want:      DeploymentMultiTenantValEProjectionBoundaryStateBlocked,
			state: func(model DeploymentMultiTenantValEFoundation) string {
				return model.ProjectionBoundaryState
			},
		},
		{
			name: "clean room trailing whitespace disclaimer blocks",
			mutate: func(model *DeploymentMultiTenantValEFoundation) {
				model.CleanRoomIPReview.ProjectionDisclaimer = deploymentMultiTenantValEProjectionDisclaimer() + " "
			},
			stateName: "clean room ip",
			want:      DeploymentMultiTenantValECleanRoomIPStateBlocked,
			state: func(model DeploymentMultiTenantValEFoundation) string {
				return model.CleanRoomIPState
			},
		},
		{
			name: "point10 pass rule aggregate disclaimer blocks",
			mutate: func(model *DeploymentMultiTenantValEFoundation) {
				model.Point10PassRule.ProjectionDisclaimer = deploymentMultiTenantValEProjectionDisclaimer() + " aggregate_dependency_snapshot"
			},
			stateName: "point10 pass rule",
			want:      DeploymentMultiTenantValEPoint10PassRuleStateBlocked,
			state: func(model DeploymentMultiTenantValEFoundation) string {
				return model.Point10PassRuleState
			},
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if got := tc.state(model); got != tc.want {
			t.Fatalf("%s: expected exact %s state %q, got %#v", tc.name, tc.stateName, tc.want, model)
		}
		if model.CurrentState != DeploymentMultiTenantValEStateBlocked || model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("%s: expected exact blocked top-level state and no point_10_pass, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValEIntegratedInvariantBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValEFoundation)
	}{
		{name: "install success treated as readiness blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.InstallSuccessTreatedAsReadiness = true
		}},
		{name: "marketplace install production approval blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.MarketplaceInstallTreatedAsProductionApproval = true
		}},
		{name: "sso configured secure blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.SSOConfiguredTreatedAsSecure = true
		}},
		{name: "rbac abac bypass blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.RBACABACBypass = true
		}},
		{name: "cross tenant leakage blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.CrossTenantLeakageDetected = true
		}},
		{name: "data residency bypass blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.DataResidencyBypassDetected = true
		}},
		{name: "cross region flow without exception blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.CrossRegionFlowUnscoped = true
		}},
		{name: "ha readiness uptime guarantee blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.HAReadinessTreatedAsUptimeGuarantee = true
		}},
		{name: "backup exists restore ready blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.BackupExistsTreatedAsRestoreReady = true
		}},
		{name: "restore evidence missing blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.RestoreEvidenceMissing = true
		}},
		{name: "sla readiness uptime guarantee blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.SLAReadinessTreatedAsUptimeGuarantee = true
		}},
		{name: "connector source of truth blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.ConnectorTreatedAsSourceOfTruth = true
		}},
		{name: "connector mutation without capability blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.ConnectorMutationWithoutExplicitCapability = true
		}},
		{name: "operator support action without authority basis blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.OperatorSupportActionWithoutAuthorityBasis = true
		}},
		{name: "break glass persistent access blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.BreakGlassPersistentGlobalAccess = true
		}},
		{name: "agent production mutation blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.AgentProductionMutation = true
		}},
		{name: "agent canonical mutation blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.AgentCanonicalMutation = true
		}},
		{name: "agent self promotes blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.AgentSelfPromotes = true
		}},
		{name: "learned output canonical truth blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.LearnedOutputCanonicalTruth = true
		}},
		{name: "msp partner pass authority blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.MSPPartnerPassAuthority = true
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.IntegratedInvariantState != DeploymentMultiTenantValEIntegratedInvariantStateBlocked {
			t.Fatalf("%s: expected blocked integrated invariant state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValEEvidenceQualityBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValEFoundation)
	}{
		{name: "missing evidence id blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceID = ""
		}},
		{name: "missing scope blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Scope = ""
		}},
		{name: "global evidence scope blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Scope = "global"
		}},
		{name: "compact all tenant evidence scope blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Scope = "alltenant"
		}},
		{name: "camel compact all tenants evidence scope blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Scope = "allTenants"
		}},
		{name: "split all tenants evidence scope blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Scope = "all tenants"
		}},
		{name: "underscore all tenants evidence scope blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Scope = "all_tenants_scope"
		}},
		{name: "standalone cross evidence scope blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Scope = "cross"
		}},
		{name: "cross scope evidence scope blocks even with scoped exception", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Scope = "cross_scope"
			model.EvidenceQualityMap.Entries[0].CrossTenant = true
			model.EvidenceQualityMap.Entries[0].ScopedAuditedException = "evidence:scoped-audited-exception-1"
		}},
		{name: "obfuscated standalone cross evidence scope blocks even with scoped exception", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Scope = "c-r-o-s-s"
			model.EvidenceQualityMap.Entries[0].CrossTenant = true
			model.EvidenceQualityMap.Entries[0].ScopedAuditedException = "evidence:scoped-audited-exception-1"
		}},
		{name: "missing deployment profile blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].DeploymentProfile = ""
		}},
		{name: "missing surface blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Surface = ""
		}},
		{name: "missing policy version blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].PolicyVersion = ""
		}},
		{name: "missing engine version blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EngineVersion = ""
		}},
		{name: "missing schema version blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].SchemaVersion = ""
		}},
		{name: "missing all hash identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceHash = ""
			model.EvidenceQualityMap.Entries[0].ArtifactHash = ""
		}},
		{name: "missing timestamp blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Timestamp = ""
		}},
		{name: "missing freshness state blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].FreshnessState = ""
		}},
		{name: "unknown evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceType = "unknown"
		}},
		{name: "partial evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Source = "partial"
		}},
		{name: "stale evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].FreshnessState = "stale"
		}},
		{name: "malformed evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].PolicyVersion = "malformed"
		}},
		{name: "raw exact policy version mismatch blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].PolicyVersion = "policy_v2"
		}},
		{name: "raw exact engine version padded with whitespace blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EngineVersion = " engine_v1 "
		}},
		{name: "timestamp padded with whitespace blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Timestamp = " " + deploymentMultiTenantValEManifestTimestampActive + " "
		}},
		{name: "non-canonical timestamp offset blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Timestamp = "2026-05-08T16:45:30+00:00"
		}},
		{name: "tenant scope padded with whitespace blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].TenantScope = " " + deploymentMultiTenantValEExpectedTenantScope() + " "
		}},
		{name: "deployment profile padded with whitespace blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].DeploymentProfile = "\t" + deploymentMultiTenantValEExpectedDeploymentProfile()
		}},
		{name: "unsupported evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].SchemaVersion = "unsupported"
		}},
		{name: "validation state padded with whitespace blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].ValidationState = " " + deploymentMultiTenantValEEvidenceValidationExact + " "
		}},
		{name: "projection boundary padded with tab newline blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].ProjectionBoundary = "\t" + deploymentMultiTenantValEEvidenceProjectionBoundary + "\n"
		}},
		{name: "blocked evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].ValidationState = "blocked"
		}},
		{name: "revoked evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceHash = "revoked_hash"
		}},
		{name: "expired evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceHash = "expired_hash"
		}},
		{name: "duplicate evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceHash = "duplicate_hash"
		}},
		{name: "unrelated evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceHash = "unrelated_hash"
		}},
		{name: "cross tenant marker in source with cross tenant flag false blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Source = "computed_val0_output_cross_tenant"
			model.EvidenceQualityMap.Entries[0].CrossTenant = false
		}},
		{name: "obfuscated standalone cross marker in source blocks identity boundary", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Source = "computed_val0_output_c-r-o-s-s"
		}},
		{name: "obfuscated standalone cross marker in evidence id blocks identity boundary", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceID = "evidence:vale-c-r-o-s-s-foundation"
		}},
		{name: "obfuscated standalone cross marker in surface blocks identity boundary", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Surface = "val_c-r-o-s-s_surface"
		}},
		{name: "tenant beta marker in evidence id blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceID = "evidence:vale-tenant-beta-foundation"
		}},
		{name: "sibling boundary marker in surface blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].Surface = "deployment_sibling_boundary_surface"
		}},
		{name: "profile-like substitution marker in evidence hash blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceHash = "company_profile_hash_v1"
		}},
		{name: "dashboard summary inferred identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].DashboardSummaryOnly = true
		}},
		{name: "fleet summary inferred identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].FleetSummaryOnly = true
		}},
		{name: "portal summary inferred identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].PortalSummaryOnly = true
		}},
		{name: "agent summary inferred identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].AgentSummaryOnly = true
		}},
		{name: "connector summary inferred identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].ConnectorSummaryOnly = true
		}},
		{name: "same name inferred identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].SameNameInferredIdentity = true
		}},
		{name: "matching path inferred identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].MatchingPathIdentity = true
		}},
		{name: "same package inferred identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].SamePackageNameIdentity = true
		}},
		{name: "summary only evidence blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].SummaryOnly = true
		}},
		{name: "cross tenant evidence without exception blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].CrossTenant = true
		}},
		{name: "padded evidence id blocks raw exact identity", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceID = " " + model.EvidenceQualityMap.Entries[0].EvidenceID + " "
		}},
		{name: "tenant gamma evidence id blocks profile-like boundary laundering", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].EvidenceID = "evidence:vale-tenant-gamma-foundation"
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.EvidenceQualityState != DeploymentMultiTenantValEEvidenceQualityStateBlocked {
			t.Fatalf("%s: expected blocked evidence quality state, got %#v", tc.name, model)
		}
		if model.CurrentState != DeploymentMultiTenantValEStateBlocked || model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("%s: expected exact blocked top-level state and no point_10_pass, got %#v", tc.name, model)
		}
		for _, expected := range []string{"evidence_quality_blocked", "point_10_not_passed"} {
			if !deploymentMultiTenantValEExactReasonPresent(model.BlockingReasons, expected) {
				t.Fatalf("%s: expected exact blocking reason %q, got %#v", tc.name, expected, model.BlockingReasons)
			}
		}
	}
}

func TestDeploymentMultiTenantValEEvidenceQualityDuplicateBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValEFoundation)
	}{
		{name: "duplicate evidence id blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			duplicate := model.EvidenceQualityMap.Entries[1]
			duplicate.EvidenceHash = "distinct_hash_v2"
			duplicate.ArtifactHash = "distinct_artifact_v2"
			duplicate.EvidenceID = model.EvidenceQualityMap.Entries[0].EvidenceID
			model.EvidenceQualityMap.Entries = append(model.EvidenceQualityMap.Entries, duplicate)
		}},
		{name: "duplicate evidence hash blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			duplicate := model.EvidenceQualityMap.Entries[1]
			duplicate.EvidenceHash = model.EvidenceQualityMap.Entries[0].EvidenceHash
			model.EvidenceQualityMap.Entries = append(model.EvidenceQualityMap.Entries, duplicate)
		}},
		{name: "duplicate artifact hash blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			duplicate := model.EvidenceQualityMap.Entries[1]
			duplicate.ArtifactHash = model.EvidenceQualityMap.Entries[0].ArtifactHash
			model.EvidenceQualityMap.Entries = append(model.EvidenceQualityMap.Entries, duplicate)
		}},
		{name: "duplicate compound identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			duplicate := model.EvidenceQualityMap.Entries[0]
			duplicate.EvidenceID = "evidence:vale-duplicate-compound"
			duplicate.EvidenceHash = "duplicate_compound_hash"
			duplicate.ArtifactHash = "duplicate_compound_artifact"
			model.EvidenceQualityMap.Entries = append(model.EvidenceQualityMap.Entries, duplicate)
		}},
		{name: "same evidence id with different hash blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			duplicate := model.EvidenceQualityMap.Entries[0]
			duplicate.EvidenceHash = "conflicting_hash_v2"
			duplicate.ArtifactHash = "conflicting_artifact_v2"
			model.EvidenceQualityMap.Entries = append(model.EvidenceQualityMap.Entries, duplicate)
		}},
		{name: "same hash with unrelated evidence id blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			duplicate := model.EvidenceQualityMap.Entries[1]
			duplicate.EvidenceHash = model.EvidenceQualityMap.Entries[0].EvidenceHash
			duplicate.EvidenceID = "evidence:vale-unrelated-id"
			model.EvidenceQualityMap.Entries = append(model.EvidenceQualityMap.Entries, duplicate)
		}},
		{name: "same artifact hash reused across unrelated tenant scope blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			duplicate := model.EvidenceQualityMap.Entries[1]
			duplicate.ArtifactHash = model.EvidenceQualityMap.Entries[0].ArtifactHash
			duplicate.TenantScope = "tenant:beta"
			duplicate.ScopedAuditedException = "evidence:scoped-exception-1"
			duplicate.CrossTenant = true
			model.EvidenceQualityMap.Entries = append(model.EvidenceQualityMap.Entries, duplicate)
		}},
		{name: "cross tenant evidence without scoped exception blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].CrossTenant = true
			model.EvidenceQualityMap.Entries[0].ScopedAuditedException = ""
		}},
		{name: "cross tenant evidence with padded scoped exception blocks raw exact boundary", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].CrossTenant = true
			model.EvidenceQualityMap.Entries[0].ScopedAuditedException = " evidence:scoped-audited-exception-1 "
		}},
		{name: "cross tenant evidence with tab newline scoped exception blocks raw exact boundary", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.EvidenceQualityMap.Entries[0].CrossTenant = true
			model.EvidenceQualityMap.Entries[0].ScopedAuditedException = "\tevidence:scoped-audited-exception-1\n"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.EvidenceQualityState != DeploymentMultiTenantValEEvidenceQualityStateBlocked {
			t.Fatalf("%s: expected blocked evidence quality state, got %#v", tc.name, model)
		}
		if model.CurrentState != DeploymentMultiTenantValEStateBlocked || model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("%s: expected exact blocked top-level state and no point_10_pass, got %#v", tc.name, model)
		}
		for _, expected := range []string{"evidence_quality_blocked", "point_10_not_passed"} {
			if !deploymentMultiTenantValEExactReasonPresent(model.BlockingReasons, expected) {
				t.Fatalf("%s: expected exact blocking reason %q, got %#v", tc.name, expected, model.BlockingReasons)
			}
		}
	}

	model := activeDeploymentMultiTenantValEModel()
	model.EvidenceQualityMap.Entries[0].CrossTenant = true
	model.EvidenceQualityMap.Entries[0].ScopedAuditedException = "evidence:scoped-audited-exception-1"
	model = ComputeDeploymentMultiTenantValEFoundation(model)
	if model.EvidenceQualityState != DeploymentMultiTenantValEEvidenceQualityStateActive {
		t.Fatalf("expected exact cross-tenant evidence with scoped audited exception to remain active, got %#v", model)
	}
	if model.CurrentState != DeploymentMultiTenantValEStatePass || model.Point10State != DeploymentMultiTenantPoint10StatePass {
		t.Fatalf("expected exact scoped audited exception happy path to preserve final pass, got %#v", model)
	}
	for _, forbidden := range []string{"evidence_quality_blocked", "point_10_not_passed"} {
		if deploymentMultiTenantValEExactReasonPresent(model.BlockingReasons, forbidden) {
			t.Fatalf("expected exact scoped audited exception happy path to exclude reason %q, got %#v", forbidden, model.BlockingReasons)
		}
	}
}

func TestDeploymentMultiTenantValEEvidenceScopeCrosscheckFalsePositiveGuard(t *testing.T) {
	model := activeDeploymentMultiTenantValEModel()
	model.EvidenceQualityMap.Entries[0].Scope = "tenant_crosscheck_scope"
	model = ComputeDeploymentMultiTenantValEFoundation(model)
	if model.EvidenceQualityState != DeploymentMultiTenantValEEvidenceQualityStateActive || model.CurrentState != DeploymentMultiTenantValEStatePass {
		t.Fatalf("expected crosscheck scope to avoid false-positive cross boundary block, got %#v", model)
	}
}

func TestDeploymentMultiTenantValEEvidenceIdentityBoundaryFalsePositiveGuards(t *testing.T) {
	model := activeDeploymentMultiTenantValEModel()
	model.EvidenceQualityMap.Entries[0].EvidenceID = "evidence:vale-crosscheck-foundation"
	model.EvidenceQualityMap.Entries[0].Source = "computed_crosscheck_output"
	model.EvidenceQualityMap.Entries[0].Scope = "tenant_smalltenant_scope"
	model.EvidenceQualityMap.Entries[0].Surface = "crosscheck_surface"
	model.EvidenceQualityMap.Entries[0].EvidenceHash = "smalltenant_crosscheck_hash_v1"
	model.EvidenceQualityMap.Entries[0].ArtifactHash = "smalltenant_crosscheck_artifact_hash_v1"
	model = ComputeDeploymentMultiTenantValEFoundation(model)
	if model.EvidenceQualityState != DeploymentMultiTenantValEEvidenceQualityStateActive || model.CurrentState != DeploymentMultiTenantValEStatePass {
		t.Fatalf("expected smalltenant/crosscheck identity values to avoid false-positive boundary block, got %#v", model)
	}
}

func TestDeploymentMultiTenantValECLBClosureLedgerBlockers(t *testing.T) {
	validAdvisory := DeploymentMultiTenantValECLBFinding{
		BlockerLevel:      DeploymentMultiTenantValEBlockerLevelCLB3,
		Surface:           DeploymentMultiTenantValEClosureSurfaceProjectionBoundary,
		Reason:            "non_blocking_advisory_note",
		BlocksCurrentWave: false,
		RequiredFollowup:  "record advisory note without strengthening pass",
	}
	testCases := []struct {
		name          string
		mutate        func(*DeploymentMultiTenantValEFoundation)
		expectedState string
	}{
		{name: "clb0 open blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB0OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: DeploymentMultiTenantValEBlockerLevelCLB0, Surface: DeploymentMultiTenantValEClosureSurfaceDependencyGate, Reason: "clb0_open", BlocksCurrentWave: true}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "clb1 open blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB1OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: DeploymentMultiTenantValEBlockerLevelCLB1, Surface: DeploymentMultiTenantValEClosureSurfaceEvidenceQuality, Reason: "clb1_open", BlocksCurrentWave: true, RequiredFollowup: "close finding"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "clb2 open blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB2OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: DeploymentMultiTenantValEBlockerLevelCLB2, Surface: DeploymentMultiTenantValEClosureSurfacePassClosureManifest, Reason: "clb2_open", BlocksCurrentWave: true, RequiredFollowup: "close finding"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "clb3 advisory alone remains active", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB3AdvisoryFindings = []DeploymentMultiTenantValECLBFinding{validAdvisory}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateActive},
		{name: "unknown blocker level blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB0OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: "CL-B9", Surface: DeploymentMultiTenantValEClosureSurfaceDependencyGate, Reason: "unknown", BlocksCurrentWave: true}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "legacy priority zero blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB0OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: deploymentMultiTenantLegacyPriority("0"), Surface: DeploymentMultiTenantValEClosureSurfaceDependencyGate, Reason: "legacy", BlocksCurrentWave: true}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "legacy priority one blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB1OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: deploymentMultiTenantLegacyPriority("1"), Surface: DeploymentMultiTenantValEClosureSurfaceDependencyGate, Reason: "legacy", BlocksCurrentWave: true, RequiredFollowup: "remove"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "legacy priority two blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB2OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: deploymentMultiTenantLegacyPriority("2"), Surface: DeploymentMultiTenantValEClosureSurfaceDependencyGate, Reason: "legacy", BlocksCurrentWave: true, RequiredFollowup: "remove"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "padded CLB1 level blocks raw exact taxonomy", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB1OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: " " + DeploymentMultiTenantValEBlockerLevelCLB1 + " ", Surface: DeploymentMultiTenantValEClosureSurfaceEvidenceQuality, Reason: "clb1_open", BlocksCurrentWave: true, RequiredFollowup: "close_finding"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "tab newline CLB2 surface blocks raw exact taxonomy", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB2OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: DeploymentMultiTenantValEBlockerLevelCLB2, Surface: "\t" + DeploymentMultiTenantValEClosureSurfacePassClosureManifest + "\n", Reason: "clb2_open", BlocksCurrentWave: true, RequiredFollowup: "close_finding"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "padded CLB3 followup blocks raw exact advisory metadata", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB3AdvisoryFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: DeploymentMultiTenantValEBlockerLevelCLB3, Surface: DeploymentMultiTenantValEClosureSurfaceNoOverclaim, Reason: "advisory_only", BlocksCurrentWave: false, RequiredFollowup: " record_advisory "}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "whitespace only CLB0 reason blocks raw exact closure finding", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CLB0OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: DeploymentMultiTenantValEBlockerLevelCLB0, Surface: DeploymentMultiTenantValEClosureSurfaceDependencyGate, Reason: " \t\n", BlocksCurrentWave: true}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "exact risk exception ref remains active", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptionRefs = []string{"risk_exception_ref"}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateActive},
		{name: "padded risk exception ref blocks raw exact ledger metadata", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptionRefs = []string{" risk_exception_ref "}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "exact followup ref remains active", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RequiredFollowupRefs = []string{"followup_ref"}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateActive},
		{name: "tab newline followup ref blocks raw exact ledger metadata", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RequiredFollowupRefs = []string{"\tfollowup_ref\n"}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "exact temporary risk exception remains active", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "tenant_scope", Reason: "need_review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateActive},
		{name: "padded risk exception id blocks raw exact metadata", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: " risk_exception_1 ", Owner: "owner_a", Scope: "tenant_scope", Reason: "need_review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "padded risk exception owner blocks raw exact metadata", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: " owner_a ", Scope: "tenant_scope", Reason: "need_review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "tab newline risk exception scope blocks raw exact metadata", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "\ttenant_scope\n", Reason: "need_review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "padded risk exception reason blocks raw exact metadata", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "tenant_scope", Reason: " need_review ", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "tab newline risk exception followup ref blocks raw exact metadata", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "tenant_scope", Reason: "need_review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "\tfollowup_ref\n"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "exact permanent risk exception governance event remains active", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "tenant_scope", Reason: "need_review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref", Permanent: true, GovernanceEvent: "governance_event_1"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateActive},
		{name: "padded permanent risk exception governance event blocks raw exact metadata", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "tenant_scope", Reason: "need_review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref", Permanent: true, GovernanceEvent: " governance_event_1 "}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "exact ip legal risk exception external review plan remains active", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "tenant_scope", Reason: "need_review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref", IPLegalException: true, ExternalReviewPlan: "external_review_plan_1"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateActive},
		{name: "tab newline ip legal risk exception external review plan blocks raw exact metadata", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "tenant_scope", Reason: "need_review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref", IPLegalException: true, ExternalReviewPlan: "\texternal_review_plan_1\n"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "risk exception missing owner blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Scope: "tenant_scope", Reason: "need review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "expired exception blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "tenant_scope", Reason: "need review", Expiry: "expired", RequiredFollowupRef: "followup_ref"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "whitespace padded exception expiry blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "tenant_scope", Reason: "need review", Expiry: " " + deploymentMultiTenantValEManifestTimestampActive + " ", RequiredFollowupRef: "followup_ref"}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "permanent exception without governance event blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.RiskExceptions = []DeploymentMultiTenantValERiskException{{ExceptionID: "risk_exception_1", Owner: "owner_a", Scope: "tenant_scope", Reason: "need review", Expiry: deploymentMultiTenantValEManifestTimestampActive, RequiredFollowupRef: "followup_ref", Permanent: true}}
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "projection boundary result laundering blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.ProjectionBoundaryResult = "projection_boundary advisory_only reviewed"
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "clean room ip result laundering blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.CleanRoomIPResult = "clean_room_ip active reviewed"
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
		{name: "no overclaim result laundering blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CLBClosureLedger.NoOverclaimResult = "no_overclaim active reviewed"
		}, expectedState: DeploymentMultiTenantValECLBClosureStateBlocked},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.CLBClosureState != tc.expectedState {
			t.Fatalf("%s: expected %q, got %#v", tc.name, tc.expectedState, model)
		}
		if tc.expectedState == DeploymentMultiTenantValECLBClosureStateBlocked {
			if model.CurrentState != DeploymentMultiTenantValEStateBlocked || model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
				t.Fatalf("%s: expected top-level fail closed, got state=%q point10=%q model=%#v", tc.name, model.CurrentState, model.Point10State, model)
			}
			for _, expected := range []string{"clb_closure_blocked", "point_10_not_passed"} {
				if !deploymentMultiTenantValEExactReasonPresent(model.BlockingReasons, expected) {
					t.Fatalf("%s: expected exact blocking reason %q, got %#v", tc.name, expected, model.BlockingReasons)
				}
			}
			continue
		}
		if model.CurrentState != DeploymentMultiTenantValEStatePass || model.Point10State != DeploymentMultiTenantPoint10StatePass {
			t.Fatalf("%s: expected CL-B3 advisory happy path to preserve final pass, got state=%q point10=%q model=%#v", tc.name, model.CurrentState, model.Point10State, model)
		}
		if deploymentMultiTenantValEExactReasonPresent(model.BlockingReasons, "clb_closure_blocked") {
			t.Fatalf("%s: expected CL-B3 advisory happy path to exclude clb_closure_blocked, got %#v", tc.name, model.BlockingReasons)
		}
	}

	model := activeDeploymentMultiTenantValEModel()
	model.CLBClosureLedger.CLB3AdvisoryFindings = []DeploymentMultiTenantValECLBFinding{validAdvisory}
	model.CLBClosureLedger.CLB1OpenFindings = []DeploymentMultiTenantValECLBFinding{{BlockerLevel: DeploymentMultiTenantValEBlockerLevelCLB1, Surface: DeploymentMultiTenantValEClosureSurfaceNoOverclaim, Reason: "clb1_open", BlocksCurrentWave: true, RequiredFollowup: "close finding"}}
	model = ComputeDeploymentMultiTenantValEFoundation(model)
	if model.CLBClosureState != DeploymentMultiTenantValECLBClosureStateBlocked {
		t.Fatalf("expected CL-B3 not to mask CL-B1, got %#v", model)
	}
}

func TestDeploymentMultiTenantValEPassClosureManifestBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValEFoundation)
	}{
		{name: "missing point id blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.PointID = ""
		}},
		{name: "wrong point id blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.PointID = "point_11"
		}},
		{name: "wrong wave id blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.WaveID = "val_d"
		}},
		{name: "missing dependency gate result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.DependencyGateResult = ""
		}},
		{name: "missing evidence identity blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.EvidenceIdentity = ""
		}},
		{name: "missing commands run blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CommandsRun = nil
		}},
		{name: "missing tests run blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.TestsRun = nil
		}},
		{name: "missing negative fixtures blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.NegativeFixturesRun = nil
		}},
		{name: "missing projection boundary result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.ProjectionBoundaryResult = ""
		}},
		{name: "missing no overclaim grep blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.NoOverclaimGrepResult = ""
		}},
		{name: "missing clean room result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CleanRoomIPResult = ""
		}},
		{name: "missing clb closure result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CLBClosureResult = ""
		}},
		{name: "missing evidence quality result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.EvidenceQualityResult = ""
		}},
		{name: "projection boundary exact token missing blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.ProjectionBoundaryResult = "projection_active_but_export_bypass"
		}},
		{name: "no overclaim unreviewed blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.NoOverclaimGrepResult = "forbidden_claims_absent_but_unreviewed"
		}},
		{name: "clean room ip failed blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CleanRoomIPResult = "clean_room_ip_failed"
		}},
		{name: "clb closure missing clb1 and clb2 blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CLBClosureResult = deploymentMultiTenantValECLBToken0None
		}},
		{name: "clb closure mixed open token blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CLBClosureResult = "clb0_none_but_clb1_open"
		}},
		{name: "pass confirmed before gates active blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.IntegratedInvariantReview.AgentSelfPromotes = true
		}},
		{name: "commit sha before pass confirmed blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.ReviewerResult = "pending_review"
			model.PassClosureManifest.CommitSHAIfAvailable = "sha123"
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.PassClosureManifestState != DeploymentMultiTenantValEPassClosureManifestStateBlocked {
			t.Fatalf("%s: expected blocked manifest state, got %#v", tc.name, model)
		}
	}

	model := activeDeploymentMultiTenantValEModel()
	model.PassClosureManifest.ProjectionBoundaryResult = deploymentMultiTenantValEManifestProjectionBoundary
	model.PassClosureManifest.NoOverclaimGrepResult = strings.Join(deploymentMultiTenantValENoOverclaimResultTokens(), " ")
	model.PassClosureManifest.CleanRoomIPResult = strings.Join(deploymentMultiTenantValECleanRoomIPResultTokens(), " ")
	model.PassClosureManifest.CLBClosureResult = strings.Join(deploymentMultiTenantValECLBClosureResultTokens(), " ")
	model = ComputeDeploymentMultiTenantValEFoundation(model)
	if model.PassClosureManifestState != DeploymentMultiTenantValEPassClosureManifestStateActive {
		t.Fatalf("expected exact manifest result tokens to remain active, got %#v", model)
	}
}

func TestDeploymentMultiTenantValEPassClosureManifestRequiresRawExactBinding(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValEFoundation)
	}{
		{name: "leading whitespace point id blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.PointID = " " + deploymentMultiTenantValEPointID
		}},
		{name: "tab wave id blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.WaveID = "\t" + deploymentMultiTenantValEWaveID
		}},
		{name: "newline scope blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.Scope = deploymentMultiTenantValEScope + "\n"
		}},
		{name: "leading whitespace dependency gate result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.DependencyGateResult = " " + model.PassClosureManifest.DependencyGateResult
		}},
		{name: "extra dependency gate token blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.DependencyGateResult += " extra_token"
		}},
		{name: "whitespace padded command blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CommandsRun[0] = " " + model.PassClosureManifest.CommandsRun[0]
		}},
		{name: "tab padded test name blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.TestsRun[0] = "\t" + model.PassClosureManifest.TestsRun[0]
		}},
		{name: "newline padded negative fixture blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.NegativeFixturesRun[0] = model.PassClosureManifest.NegativeFixturesRun[0] + "\n"
		}},
		{name: "leading whitespace projection boundary result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.ProjectionBoundaryResult = " " + deploymentMultiTenantValEManifestProjectionBoundary
		}},
		{name: "tab padded no overclaim result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.NoOverclaimGrepResult = "\t" + strings.Join(deploymentMultiTenantValENoOverclaimResultTokens(), " ")
		}},
		{name: "newline padded clean room result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CleanRoomIPResult = strings.Join(deploymentMultiTenantValECleanRoomIPResultTokens(), " ") + "\n"
		}},
		{name: "trailing whitespace clb closure result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CLBClosureResult = strings.Join(deploymentMultiTenantValECLBClosureResultTokens(), " ") + " "
		}},
		{name: "leading whitespace evidence quality result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.EvidenceQualityResult = " " + DeploymentMultiTenantValEEvidenceQualityStateActive
		}},
		{name: "trailing whitespace cross-wave result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CrossWaveInvariantResult = DeploymentMultiTenantValEIntegratedInvariantStateActive + " "
		}},
		{name: "whitespace reviewer result blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.ReviewerResult = " " + DeploymentMultiTenantValEReviewerResultPassConfirmed
		}},
		{name: "whitespace not-yet-committed sentinel blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CommitSHAIfAvailable = deploymentMultiTenantValENotYetCommitted + " "
		}},
		{name: "whitespace padded manifest timestamp blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.Timestamp = " " + deploymentMultiTenantValEManifestTimestampActive + " "
		}},
		{name: "manifest timestamp with +00:00 offset blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.Timestamp = "2026-05-08T16:45:30+00:00"
		}},
		{name: "different canonical manifest timestamp blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.Timestamp = "2026-05-08T16:45:30Z"
		}},
		{name: "non-sentinel commit sha blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.PassClosureManifest.CommitSHAIfAvailable = "abc123"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.PassClosureManifestState != DeploymentMultiTenantValEPassClosureManifestStateBlocked {
			t.Fatalf("%s: expected blocked manifest state, got %#v", tc.name, model)
		}
		if model.CurrentState != DeploymentMultiTenantValEStateBlocked || model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("%s: expected exact blocked top-level state and no point_10_pass, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValEManifestEvidenceIdentityValidation(t *testing.T) {
	canonicalIdentity := "policy_version=" + deploymentMultiTenantValEExpectedPolicyVersion() +
		" engine_version=" + deploymentMultiTenantValEExpectedEngineVersion() +
		" schema_version=" + deploymentMultiTenantValEExpectedSchemaVersion() +
		" tenant_scope=" + deploymentMultiTenantValEExpectedTenantScope() +
		" deployment_profile=" + deploymentMultiTenantValEExpectedDeploymentProfile()
	invalidIdentities := []string{
		"policy_version= engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version= schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version= tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope= deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=",
		"policy_version=    engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=    schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=    tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=    deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=   ",
		"engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha",
		"policy_version=unknown engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=expired_engine schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=duplicate_schema tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=global_admin_scope deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=unrelated_profile",
		"policy_version=policy_v2 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v2 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v2 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant_scope_alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileTenantIsolated,
		"policy_version=policy_v1 policy_version=policy_v2 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha tenant_scope=tenant:beta deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP + " canonical_truth=foo",
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP + " point_10_pass=active",
		"policy_version engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile",
		"policy_version=<empty> engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version: engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"dashboard summary policy_version=policy_v1 engine_version=engine_v1 schema_version=schema_v1 tenant_scope=tenant:alpha deployment_profile=" + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		"policy_version: point10_vale_policy_v1 engine_version: point10_vale_engine_v1 schema_version: point10_vale_schema_v1 tenant_scope: tenant_scope_alpha deployment_profile: " + DeploymentMultiTenantProfileBoundedMarketplaceMSP,
		" " + canonicalIdentity,
		canonicalIdentity + " ",
		"\t" + canonicalIdentity,
		canonicalIdentity + "\n",
		"policy_version=" + deploymentMultiTenantValEExpectedPolicyVersion() + "\tengine_version=" + deploymentMultiTenantValEExpectedEngineVersion() + " schema_version=" + deploymentMultiTenantValEExpectedSchemaVersion() + " tenant_scope=" + deploymentMultiTenantValEExpectedTenantScope() + " deployment_profile=" + deploymentMultiTenantValEExpectedDeploymentProfile(),
		"policy_version=" + deploymentMultiTenantValEExpectedPolicyVersion() + "\nengine_version=" + deploymentMultiTenantValEExpectedEngineVersion() + " schema_version=" + deploymentMultiTenantValEExpectedSchemaVersion() + " tenant_scope=" + deploymentMultiTenantValEExpectedTenantScope() + " deployment_profile=" + deploymentMultiTenantValEExpectedDeploymentProfile(),
		"policy_version=" + deploymentMultiTenantValEExpectedPolicyVersion() + "  engine_version=" + deploymentMultiTenantValEExpectedEngineVersion() + " schema_version=" + deploymentMultiTenantValEExpectedSchemaVersion() + " tenant_scope=" + deploymentMultiTenantValEExpectedTenantScope() + " deployment_profile=" + deploymentMultiTenantValEExpectedDeploymentProfile(),
		"policy_version: " + deploymentMultiTenantValEExpectedPolicyVersion() + "\tengine_version: " + deploymentMultiTenantValEExpectedEngineVersion() + " schema_version: " + deploymentMultiTenantValEExpectedSchemaVersion() + " tenant_scope: " + deploymentMultiTenantValEExpectedTenantScope() + " deployment_profile: " + deploymentMultiTenantValEExpectedDeploymentProfile(),
		"policy_version: " + deploymentMultiTenantValEExpectedPolicyVersion() + "  engine_version: " + deploymentMultiTenantValEExpectedEngineVersion() + " schema_version: " + deploymentMultiTenantValEExpectedSchemaVersion() + " tenant_scope: " + deploymentMultiTenantValEExpectedTenantScope() + " deployment_profile: " + deploymentMultiTenantValEExpectedDeploymentProfile(),
	}

	for _, evidenceIdentity := range invalidIdentities {
		model := activeDeploymentMultiTenantValEModel()
		model.PassClosureManifest.EvidenceIdentity = evidenceIdentity
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.PassClosureManifestState != DeploymentMultiTenantValEPassClosureManifestStateBlocked {
			t.Fatalf("expected blocked manifest identity for %q, got %#v", evidenceIdentity, model)
		}
		if model.CurrentState != DeploymentMultiTenantValEStateBlocked {
			t.Fatalf("expected blocked final state for %q, got %#v", evidenceIdentity, model)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("expected exact not-complete point10 state for %q, got %#v", evidenceIdentity, model)
		}
		if strings.Contains(mustMarshalDeploymentMultiTenantValEJSON(t, model), "\"point_10_state\":\""+DeploymentMultiTenantPoint10StatePass+"\"") {
			t.Fatalf("expected blocked manifest identity for %q not to emit point_10_pass, got %#v", evidenceIdentity, model)
		}
	}

	model := activeDeploymentMultiTenantValEModel()
	model.PassClosureManifest.EvidenceIdentity = "policy_version: " + deploymentMultiTenantValEExpectedPolicyVersion() + " engine_version: " + deploymentMultiTenantValEExpectedEngineVersion() + " schema_version: " + deploymentMultiTenantValEExpectedSchemaVersion() + " tenant_scope: " + deploymentMultiTenantValEExpectedTenantScope() + " deployment_profile: " + deploymentMultiTenantValEExpectedDeploymentProfile()
	model = ComputeDeploymentMultiTenantValEFoundation(model)
	if model.PassClosureManifestState != DeploymentMultiTenantValEPassClosureManifestStateActive {
		t.Fatalf("expected valid colon syntax evidence identity to remain active, got %#v", model)
	}
	if model.CurrentState != DeploymentMultiTenantValEStatePass {
		t.Fatalf("expected valid colon syntax evidence identity to preserve final pass, got %#v", model)
	}
}

func TestDeploymentMultiTenantValENoOverclaimBlockers(t *testing.T) {
	blockedClaims := []string{
		"production approved",
		"prɔduction apprɔved",
		"production audit note approved",
		"prɔduction audit nɔte apprɔved",
		"deployment approved",
		"marketplace certified",
		"MSP certified",
		"regulator-approved",
		"compliance guaranteed",
		"compliant by default",
		"one-click secure",
		"zero-risk deployment",
		"tenant safe by default",
		"customer ready without validation",
		"deployment readiness guaranteed",
		"install success means ready",
		"marketplace install means ready",
		"marketplace production ready",
		"self-hosted production approved",
		"air-gapped certified",
		"air-gapped means fully offline verified",
		"guaranteed uptime",
		"zero downtime",
		"SLA guaranteed",
		"production SLA approved",
		"HA certified",
		"HA guaranteed",
		"failover guaranteed",
		"restore guaranteed",
		"DR guaranteed",
		"disaster recovery guaranteed",
		"backup guarantees recovery",
		"restore always works",
		"backup exists means ready",
		"healthcheck green means fully ready",
		"failover configured means ready",
		"SLA readiness means uptime guarantee",
		"supportability evidence means SLA guarantee",
		"supportability evidence review means SLA guarantee",
		"tenant isolation guaranteed",
		"data residency certified",
		"data residency guaranteed",
		"privacy guaranteed",
		"no side-channel leakage guaranteed",
		"tenant trust certified",
		"tenant trust scope certified",
		"key custody certified",
		"fleet view is canonical truth",
		"region summary is canonical truth",
		"dashboard proves tenant isolation",
		"portal view is canonical truth",
		"portal view projection is canonical truth",
		"connector is source of truth",
		"connector advisory is source of truth",
		"connector approved deployment",
		"connector certified evidence",
		"connector mutation safe by default",
		"connector can mutate without approval",
		"operator fully trusted",
		"operator approved deployment",
		"support access cannot leak",
		"break-glass safe by default",
		"break-glass permanent access",
		"MSP approved deployment",
		"MSP certified deployment",
		"partner certified deployment",
		"partner approved",
		"autonomous remediation approved",
		"agent approved deployment",
		"agent approved bounded deployment",
		"agent certified recovery",
		"AI certified fix",
		"auto-merge safe",
		"auto-deploy safe",
		"production autopatch",
		"agent guaranteed tenant isolation",
		"agent proves compliance",
		"point 10 pass",
		"point_" + "10_pass",
		"point_" + "10_pass by agent",
		"point 10 pass by agent",
		"point 10 audit pass by agent",
		"agent is source of truth",
		"external AI verified",
		"external API verified",
		"no human approval required",
		"production mutation approved by agent",
		"canonical evidence mutation by agent",
		"self-learning agent approved deployment",
		"agent self-approved model",
		"autonomous model promotion",
		"autonomous runtime activation",
		"learned output is canonical truth",
		"learned output advisory is canonical truth",
		"learned model certified",
		"AI model certified",
		"self-improving agent guarantees compliance",
		"point_" + "10_pass by learned model",
		"clean-room certified",
		"patent cleared",
		"FTO cleared",
		"legal certification",
		"copied competitor workflow",
		"same as competitor but cheaper",
		"reverse-engineered competitor platform",
		"evidence-linked readiness state production approved",
		"SLA readiness evidence SLA guaranteed",
		"supportability evidence guaranteed uptime",
		"bounded marketplace deployment profile marketplace production ready",
		"sandboxed connector execution connector is source of truth",
		"bounded operator authority operator approved deployment",
		"advisory recommendation agent approved deployment",
		"human-approved action required no human approval required",
		"offline sandbox learning pipeline learned output is canonical truth",
		"clean-room/IP guardrail patent cleared",
		"production approval",
		"deployment approval",
		"not production approval production approved",
		"\U0001d429\U0001d42b\U0001d428\U0001d41d\U0001d42e\U0001d41c\U0001d42d\U0001d422\U0001d428\U0001d427 \U0001d41a\U0001d429\U0001d429\U0001d42b\U0001d428\U0001d42f\U0001d41e\U0001d41d",
		"ρroduction approved",
		"pгoduction approved",
		"produсtion approved",
		"producτion approved",
		"productιon approved",
		"production αpproved",
		"production appro\u03bded",
		"production approvεd",
		"prod\u03c5ction approved",
		"prod\U00001d1cction approved",
		"prod\uab4ection approved",
		"prod\u028action approved",
		"production approve\u03b4",
		"productio\u03b7 approved",
		"productio\u043f approved",
		"productio\u014b approved",
		"pub\u01c0ic badge",
		"comp\u01c0iance guaranteed",
		"c\u0585mpliance guaranteed",
	}
	for _, claim := range blockedClaims {
		model := activeDeploymentMultiTenantValEModel()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValENoOverclaimStateBlocked || model.CurrentState != DeploymentMultiTenantValEStateBlocked || model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("expected blocked no-overclaim for %q, got %#v", claim, model)
		}
		for _, expected := range []string{"no_overclaim_blocked", "point_10_not_passed"} {
			if !deploymentMultiTenantValEExactReasonPresent(model.BlockingReasons, expected) {
				t.Fatalf("expected blocked no-overclaim for %q to include reason %q, got %#v", claim, expected, model.BlockingReasons)
			}
		}
	}

	splitBlockedClaims := [][]string{
		{"production", "approved"},
		{"not production approval", "approved"},
		{"not production approval", "human-approved action required", "approved"},
		{"not production approval", "apprоved"},
		{"point 10 pass", "by agent"},
		{"connector", "is source of truth"},
		{"learned output", "is canonical truth"},
		{
			"\U0001d41d\U0001d41e\U0001d429\U0001d425\U0001d428\U0001d432\U0001d426\U0001d41e\U0001d427\U0001d42d",
			"\U0001d41a\U0001d429\U0001d429\U0001d42b\U0001d428\U0001d42f\U0001d41e\U0001d41d",
		},
		{"production", "αpproved"},
		{"production", "appro\u03bded"},
		{"production", "approvεd"},
		{"prod\u03c5c", "tion approved"},
		{"prod\U00001d1cc", "tion approved"},
		{"prod\uab4e", "ction approved"},
		{"prod\u028a", "ction approved"},
		{"production", "approve\u03b4"},
		{"productio\u0578", "approved"},
		{"productio\u019e", "approved"},
		{"pub\u01c0ic", "badge"},
		{"comp\u01c0iance", "guaranteed"},
		{"c\u0585mpliance", "guaranteed"},
		{"produc", "tion approved"},
	}
	for _, claims := range splitBlockedClaims {
		model := activeDeploymentMultiTenantValEModel()
		model.NoOverclaim.ObservedClaims = claims
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValENoOverclaimStateBlocked || model.CurrentState != DeploymentMultiTenantValEStateBlocked || model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("expected blocked split no-overclaim for %q, got %#v", claims, model)
		}
		for _, expected := range []string{"no_overclaim_blocked", "point_10_not_passed"} {
			if !deploymentMultiTenantValEExactReasonPresent(model.BlockingReasons, expected) {
				t.Fatalf("expected blocked split no-overclaim for %q to include reason %q, got %#v", claims, expected, model.BlockingReasons)
			}
		}
	}

	t.Run("repetitive split scan stays bounded without backtracking", func(t *testing.T) {
		values := make([]string, 300)
		allowed := make([]bool, len(values))
		for i := range values {
			values[i] = "agent"
		}
		if deploymentMultiTenantValEForbiddenPhraseAcrossValues(values, allowed, "agent agent agent agent denied") {
			t.Fatalf("expected repetitive non-matching split phrase corpus not to block")
		}
	})

	allowedSplitClaims := [][]string{
		{"validated deployment baseline", "not production approval"},
	}
	for _, claims := range allowedSplitClaims {
		model := activeDeploymentMultiTenantValEModel()
		model.NoOverclaim.ObservedClaims = claims
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValENoOverclaimStateActive || model.CurrentState != DeploymentMultiTenantValEStatePass || model.Point10State != DeploymentMultiTenantPoint10StatePass {
			t.Fatalf("expected allowed split no-overclaim for %q to preserve active/pass closure, got %#v", claims, model)
		}
		if deploymentMultiTenantValEExactReasonPresent(model.BlockingReasons, "no_overclaim_blocked") {
			t.Fatalf("expected allowed split no-overclaim for %q not to include no_overclaim_blocked reason, got %#v", claims, model.BlockingReasons)
		}
	}

	t.Run("allowed disclaimer plus non-allowed suffix blocks with exact no-overclaim reason", func(t *testing.T) {
		model := activeDeploymentMultiTenantValEModel()
		model.NoOverclaim.ObservedClaims = []string{"not production approval", "approved"}
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValENoOverclaimStateBlocked || model.CurrentState != DeploymentMultiTenantValEStateBlocked || model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("expected blocked split exploit to fail closed, got %#v", model)
		}
		for _, expected := range []string{"no_overclaim_blocked", "point_10_not_passed"} {
			if !deploymentMultiTenantValEExactReasonPresent(model.BlockingReasons, expected) {
				t.Fatalf("expected split exploit to include reason %q, got %#v", expected, model.BlockingReasons)
			}
		}
	})

	t.Run("repetitive partial split phrase corpus remains bounded", func(t *testing.T) {
		values := make([]string, 0, 2048)
		allowed := make([]bool, 0, 2048)
		for i := 0; i < 2048; i++ {
			values = append(values, "point 10 pass by")
			allowed = append(allowed, false)
		}
		if deploymentMultiTenantValEForbiddenPhraseAcrossValues(values, allowed, "point 10 pass by agent") {
			t.Fatalf("expected repetitive partial split phrase corpus without terminal token not to block")
		}
	})

	t.Run("large compact split phrase corpus still blocks terminal token", func(t *testing.T) {
		values := make([]string, 0, 2049)
		allowed := make([]bool, 0, 2049)
		for i := 0; i < 2048; i++ {
			values = append(values, "point 10 pass by")
			allowed = append(allowed, true)
		}
		values = append(values, "agent")
		allowed = append(allowed, false)
		if !deploymentMultiTenantValEForbiddenPhraseAcrossValues(values, allowed, "point 10 pass by agent") {
			t.Fatalf("expected large split corpus with terminal unsafe bucket to block")
		}
	})

	t.Run("allowed compact phrase plus harmless non allowed suffix does not false positive", func(t *testing.T) {
		values := []string{"not production approval", "bounded evidence note"}
		allowed := []bool{true, false}
		if deploymentMultiTenantValEForbiddenPhraseAcrossValues(values, allowed, "production approval") {
			t.Fatalf("expected allowed disclaimer occurrence plus harmless suffix not to block")
		}
	})

	allowedClaims := []string{
		"validated deployment baseline",
		"evidence-linked readiness state",
		"bounded marketplace deployment profile",
		"tenant-scoped operational model",
		"advisory fleet visibility",
		"bounded operator authority",
		"sandboxed connector execution",
		"explicit connector capability",
		"connector misuse signal",
		"operator misuse signal",
		"tenant-scoped support access",
		"break-glass approval required",
		"break-glass expiry and revocation evidence",
		"HA readiness evidence",
		"failover test evidence",
		"backup freshness evidence",
		"restore test evidence",
		"tenant-scoped restore test",
		"DR drill evidence",
		"RPO/RTO target",
		"SLA readiness evidence",
		"supportability evidence",
		"known limitations",
		"tenant trust scope evidence",
		"evidence silo validation",
		"audit silo validation",
		"export silo validation",
		"privacy guard evidence",
		"side-channel negative test",
		"bounded aggregation rules",
		"advisory recommendation",
		"human-approved action required",
		"approval-gated recovery recommendation",
		"offline sandbox learning pipeline",
		"candidate model version",
		"learned output remains advisory",
		"no production autopatch",
		"no auto-merge",
		"no auto-deploy",
		"clean-room/IP guardrail",
		"public API integration",
		"standards-based evidence format",
		"not uptime guarantee",
		"not production approval",
		"not deployment approval",
		"not compliance certification",
		"not canonical truth",
		"not legal certification",
		"not patent/FTO clearance",
	}
	for _, claim := range allowedClaims {
		model := activeDeploymentMultiTenantValEModel()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValENoOverclaimStateActive || model.CurrentState != DeploymentMultiTenantValEStatePass || model.Point10State != DeploymentMultiTenantPoint10StatePass {
			t.Fatalf("expected allowed claim %q to preserve active/pass closure, got %#v", claim, model)
		}
	}
}

func TestDeploymentMultiTenantValEProjectionBoundaryBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValEFoundation)
	}{
		{name: "dashboard canonical truth blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.ProjectionBoundaryReview.Surfaces[0].CanonicalTruth = true
		}},
		{name: "fleet canonical truth blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.ProjectionBoundaryReview.Surfaces[1].CanonicalTruth = true
		}},
		{name: "marketplace production approval blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.ProjectionBoundaryReview.Surfaces[2].ApprovesProductionReadiness = true
		}},
		{name: "msp partner pass approval blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.ProjectionBoundaryReview.Surfaces[3].ApprovesPass = true
		}},
		{name: "connector summary canonical truth blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.ProjectionBoundaryReview.Surfaces[5].CanonicalTruth = true
		}},
		{name: "agent recommendation canonical truth blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.ProjectionBoundaryReview.Surfaces[7].CanonicalTruth = true
		}},
		{name: "projection hides blocked state blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.ProjectionBoundaryReview.Surfaces[0].HidesBlockedDegradedUnsupportedState = true
		}},
		{name: "missing projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.ProjectionBoundaryReview.Surfaces[0].Disclaimer = ""
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.ProjectionBoundaryState != DeploymentMultiTenantValEProjectionBoundaryStateBlocked {
			t.Fatalf("%s: expected blocked projection boundary state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValECleanRoomIPBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValEFoundation)
	}{
		{name: "copied competitor code blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.CopiedCompetitorCodePresent = true
		}},
		{name: "copied competitor text blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.CopiedCompetitorTextPresent = true
		}},
		{name: "copied competitor ui blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.CopiedCompetitorUIPresent = true
		}},
		{name: "proprietary workflow blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.ProprietaryWorkflowCopied = true
		}},
		{name: "reverse engineering language blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.ReverseEngineeringLanguagePresent = true
		}},
		{name: "confidential material blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.ConfidentialThirdPartyMaterialUsed = true
		}},
		{name: "patent cleared claim blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.PatentClearedClaim = true
		}},
		{name: "fto cleared claim blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.FTOClearedClaim = true
		}},
		{name: "legal certification claim blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.LegalCertificationClaim = true
		}},
		{name: "missing public api boundary blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.PublicAPIBoundaryPresent = false
		}},
		{name: "missing ip origin ledger blocks", mutate: func(model *DeploymentMultiTenantValEFoundation) {
			model.CleanRoomIPReview.IPOriginLedgerPresent = false
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValEModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValEFoundation(model)
		if model.CleanRoomIPState != DeploymentMultiTenantValECleanRoomIPStateBlocked {
			t.Fatalf("%s: expected blocked clean-room/ip state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValEPoint10PassSafety(t *testing.T) {
	blockedModel := activeDeploymentMultiTenantValEModel()
	blockedModel.Point10PassRule.AllTestsPassed = false
	blockedModel = ComputeDeploymentMultiTenantValEFoundation(blockedModel)
	if blockedModel.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		t.Fatalf("expected blocked model to keep point10 not complete, got %#v", blockedModel)
	}
	blockedJSON := mustMarshalDeploymentMultiTenantValEJSON(t, blockedModel)
	if strings.Contains(blockedJSON, DeploymentMultiTenantPoint10StatePass) {
		t.Fatalf("expected blocked model json not to contain point_10_pass, got %s", blockedJSON)
	}
	blockedReasons := map[string]bool{}
	for _, reason := range blockedModel.BlockingReasons {
		blockedReasons[reason] = true
	}
	for _, expected := range []string{"final_pass_rule_blocked", "point_10_not_passed"} {
		if !blockedReasons[expected] {
			t.Fatalf("expected blocked model to include reason %q, got %#v", expected, blockedModel.BlockingReasons)
		}
	}

	blockedNoOverclaim := activeDeploymentMultiTenantValEModel()
	blockedNoOverclaim.NoOverclaim.ObservedClaims = []string{"production approved"}
	blockedNoOverclaim = ComputeDeploymentMultiTenantValEFoundation(blockedNoOverclaim)
	blockedNoOverclaimReasons := map[string]bool{}
	for _, reason := range blockedNoOverclaim.BlockingReasons {
		blockedNoOverclaimReasons[reason] = true
	}
	for _, expected := range []string{"no_overclaim_blocked", "point_10_not_passed"} {
		if !blockedNoOverclaimReasons[expected] {
			t.Fatalf("expected no-overclaim blocked model to include reason %q, got %#v", expected, blockedNoOverclaim.BlockingReasons)
		}
	}

	blockedProjection := activeDeploymentMultiTenantValEModel()
	blockedProjection.ProjectionBoundaryReview.Surfaces[0].CanonicalTruth = true
	blockedProjection = ComputeDeploymentMultiTenantValEFoundation(blockedProjection)
	blockedProjectionReasons := map[string]bool{}
	for _, reason := range blockedProjection.BlockingReasons {
		blockedProjectionReasons[reason] = true
	}
	for _, expected := range []string{"projection_boundary_blocked", "point_10_not_passed"} {
		if !blockedProjectionReasons[expected] {
			t.Fatalf("expected projection blocked model to include reason %q, got %#v", expected, blockedProjection.BlockingReasons)
		}
	}

	happyModel := activeDeploymentMultiTenantValEModel()
	happyJSON := mustMarshalDeploymentMultiTenantValEJSON(t, happyModel)
	if strings.Count(happyJSON, DeploymentMultiTenantPoint10StatePass) == 0 {
		t.Fatalf("expected happy path json to contain point_10_pass, got %s", happyJSON)
	}
	if len(happyModel.BlockingReasons) != 0 {
		t.Fatalf("expected happy path to have no blocking reasons, got %#v", happyModel.BlockingReasons)
	}
	for _, forbidden := range []string{"final_pass_rule_blocked", "point_10_not_passed"} {
		for _, reason := range happyModel.BlockingReasons {
			if reason == forbidden {
				t.Fatalf("expected happy path to exclude blocking reason %q, got %#v", forbidden, happyModel.BlockingReasons)
			}
		}
	}
}

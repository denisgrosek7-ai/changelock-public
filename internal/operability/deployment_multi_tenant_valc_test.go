package operability

import (
	"encoding/json"
	"strings"
	"testing"
)

func activeDeploymentMultiTenantValCModel() DeploymentMultiTenantValCFoundation {
	model := DeploymentMultiTenantValCFoundationModel()
	return ComputeDeploymentMultiTenantValCFoundation(model)
}

func deploymentMultiTenantValCHasFinding(findings []DeploymentMultiTenantValCClosureBlockerFinding, level, surface, reason string) bool {
	for _, finding := range findings {
		if finding.BlockerLevel == level &&
			finding.Surface == surface &&
			finding.Reason == reason {
			return true
		}
	}
	return false
}

func TestDeploymentMultiTenantValCHappyPathAndPoint10NotComplete(t *testing.T) {
	model := activeDeploymentMultiTenantValCModel()
	if model.CurrentState != DeploymentMultiTenantValCStateActive {
		t.Fatalf("expected active Val C state, got %#v", model)
	}
	if model.ClosureBlockerState != DeploymentMultiTenantValCClosureBlockerStateActive {
		t.Fatalf("expected clean closure blocker state, got %#v", model)
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		t.Fatalf("expected point 10 to remain not complete, got %#v", model)
	}
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal model: %v", err)
	}
	if strings.Contains(string(payload), "point_"+"10_pass") {
		t.Fatalf("expected Val C to never emit point 10 pass, got %s", string(payload))
	}
}

func TestDeploymentMultiTenantValCAggregateProjectionDisclaimerBlocks(t *testing.T) {
	model := activeDeploymentMultiTenantValCModel()
	model.ProjectionDisclaimer = "canonical_truth"
	model = ComputeDeploymentMultiTenantValCFoundation(model)
	if model.CurrentState != DeploymentMultiTenantValCStateBlocked {
		t.Fatalf("expected malformed aggregate projection disclaimer to block ValC state, got %#v", model)
	}
	if !containsTrimmedString(model.BlockingReasons, "aggregate_projection_disclaimer_blocked") {
		t.Fatalf("expected aggregate projection disclaimer blocking reason, got %#v", model.BlockingReasons)
	}
}

func TestDeploymentMultiTenantValCProjectionDisclaimerExactBoundedBlockers(t *testing.T) {
	testCases := []struct {
		name                string
		mutate              func(*DeploymentMultiTenantValCFoundation)
		wantDisciplineState string
	}{
		{
			name: "aggregate snapshot disclaimer blocks live foundation",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.ProjectionDisclaimer = deploymentMultiTenantValCProjectionDisclaimer() + " aggregate_dependency_snapshot"
			},
		},
		{
			name: "leading whitespace aggregate disclaimer blocks live foundation",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.ProjectionDisclaimer = " " + deploymentMultiTenantValCProjectionDisclaimer()
			},
		},
		{
			name: "ha readiness uppercase disclaimer blocks live discipline",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.HAReadiness.ProjectionDisclaimer = strings.ToUpper(deploymentMultiTenantValCProjectionDisclaimer())
			},
			wantDisciplineState: DeploymentMultiTenantValCHAReadinessStateBlocked,
		},
		{
			name: "no overclaim aggregate disclaimer blocks live discipline",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.NoOverclaim.ProjectionDisclaimer = deploymentMultiTenantValCProjectionDisclaimer() + " aggregate_dependency_snapshot"
			},
			wantDisciplineState: DeploymentMultiTenantValCNoOverclaimStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.CurrentState != DeploymentMultiTenantValCStateBlocked || model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("%s: expected blocked ValC state and not-complete point10 state, got %#v", tc.name, model)
		}
		switch tc.wantDisciplineState {
		case DeploymentMultiTenantValCHAReadinessStateBlocked:
			if model.HAReadinessState != tc.wantDisciplineState {
				t.Fatalf("%s: expected blocked ha readiness state, got %#v", tc.name, model)
			}
		case DeploymentMultiTenantValCNoOverclaimStateBlocked:
			if model.NoOverclaimState != tc.wantDisciplineState {
				t.Fatalf("%s: expected blocked no-overclaim state, got %#v", tc.name, model)
			}
		}
	}
}

func TestDeploymentMultiTenantValCDependencyBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValCFoundation)
	}{
		{name: "valb current state partial blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBCurrentState = "partial"
		}},
		{name: "valb dependency state blocked blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBDependencyState = DeploymentMultiTenantValBDependencyStateBlocked
		}},
		{name: "valb tenant isolation state blocked blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBTenantIsolationState = DeploymentMultiTenantValBTenantIsolationStateBlocked
		}},
		{name: "valb data residency state blocked blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBDataResidencyState = DeploymentMultiTenantValBDataResidencyStateBlocked
		}},
		{name: "valb tenant lifecycle state blocked blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBTenantLifecycleState = DeploymentMultiTenantValBTenantLifecycleStateBlocked
		}},
		{name: "valb fair share quota state blocked blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBFairShareQuotaState = DeploymentMultiTenantValBFairShareQuotaStateBlocked
		}},
		{name: "valb no overclaim state blocked blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBNoOverclaimState = DeploymentMultiTenantValBNoOverclaimStateBlocked
		}},
		{name: "valb closure blocker state blocked blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBClosureBlockerState = DeploymentMultiTenantValBClosureBlockerStateBlocked
		}},
		{name: "valb closure blocker state cleanup blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBClosureBlockerState = DeploymentMultiTenantValBClosureBlockerStateCleanup
		}},
		{name: "valb closure blocker state advisory blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBClosureBlockerState = DeploymentMultiTenantValBClosureBlockerStateAdvisory
		}},
		{name: "point10 state complete blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.Point10State = "deployment_multi_tenant_point_10_complete"
		}},
		{name: "malformed projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ProjectionDisclaimer = "canonical_truth"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.DependencyState != DeploymentMultiTenantValCDependencyStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected blocked dependency state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValCDependencySnapshotCopiesAggregateComputedValBProjectionDisclaimer(t *testing.T) {
	valB := ComputeDeploymentMultiTenantValBFoundation(DeploymentMultiTenantValBFoundationModel())
	valB.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_valb aggregate_dependency_snapshot"
	valB.ClosureBlockerOverlay.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_valb component_closure_blocker"
	snapshot := DeploymentMultiTenantValCDependencySnapshot{
		ValBCurrentState:         valB.CurrentState,
		ValBDependencyState:      valB.DependencyState,
		ValBTenantIsolationState: valB.TenantIsolationState,
		ValBDataResidencyState:   valB.DataResidencyState,
		ValBTenantLifecycleState: valB.TenantLifecycleState,
		ValBFairShareQuotaState:  valB.FairShareQuotaState,
		ValBNoOverclaimState:     valB.NoOverclaimState,
		ValBClosureBlockerState:  valB.ClosureBlockerState,
		Point10State:             valB.Point10State,
		ProjectionDisclaimer:     valB.ProjectionDisclaimer,
	}
	if snapshot.ProjectionDisclaimer != valB.ProjectionDisclaimer {
		t.Fatalf("expected aggregate ValB disclaimer to propagate exactly, got snapshot=%q valb=%q", snapshot.ProjectionDisclaimer, valB.ProjectionDisclaimer)
	}
	if snapshot.ProjectionDisclaimer == valB.ClosureBlockerOverlay.ProjectionDisclaimer {
		t.Fatalf("expected dependency snapshot not to fallback to component disclaimer, got snapshot=%q component=%q", snapshot.ProjectionDisclaimer, valB.ClosureBlockerOverlay.ProjectionDisclaimer)
	}
	if EvaluateDeploymentMultiTenantValCDependencyState(snapshot) != DeploymentMultiTenantValCDependencyStateActive {
		t.Fatalf("expected copied aggregate disclaimer to keep dependency active, got %#v", snapshot)
	}

	valB.ProjectionDisclaimer = "canonical_truth"
	snapshot.ProjectionDisclaimer = valB.ProjectionDisclaimer
	if EvaluateDeploymentMultiTenantValCDependencyState(snapshot) != DeploymentMultiTenantValCDependencyStateBlocked {
		t.Fatalf("expected malformed aggregate disclaimer to block dependency without component fallback, got %#v", snapshot)
	}
}

func TestDeploymentMultiTenantValCWhitespaceRetaggedDependencySnapshotBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValCFoundation)
	}{
		{name: "whitespace retagged valb current state blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBCurrentState = " " + DeploymentMultiTenantValBStateActive + " "
		}},
		{name: "tab retagged valb dependency state blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBDependencyState = "\t" + DeploymentMultiTenantValBDependencyStateActive
		}},
		{name: "newline retagged point10 state blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.Point10State = DeploymentMultiTenantPoint10StateNotComplete + "\n"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.DependencyState != DeploymentMultiTenantValCDependencyStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected blocked dependency state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValCHAReadinessBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValCFoundation)
	}{
		{name: "missing topology evidence blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.TopologyEvidence = ""
		}},
		{name: "missing failover test evidence blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.FailoverTestEvidence = ""
		}},
		{name: "missing dependency degradation behavior blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.DependencyDegradationBehavior = ""
		}},
		{name: "missing healthcheck state model blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.HealthcheckStateModel = ""
		}},
		{name: "missing queue worker recovery behavior blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.QueueWorkerRecoveryBehavior = ""
		}},
		{name: "missing degraded mode semantics blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.DegradedModeSemantics = ""
		}},
		{name: "missing monitoring alert routing evidence blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.MonitoringAlertRoutingEvidence = ""
		}},
		{name: "leading whitespace evidence ref blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.EvidenceRefs[0] = " " + model.HAReadiness.EvidenceRefs[0]
		}},
		{name: "ha readiness not evidence linked blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.HAReadinessEvidenceLinked = false
		}},
		{name: "failover configured treated as ready blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.FailoverConfiguredTreatedAsReady = true
		}},
		{name: "healthcheck green treated as fully ready blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.HealthcheckGreenTreatedAsFullyReady = true
		}},
		{name: "ha readiness treated as uptime guarantee blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.HAReadinessTreatedAsUptimeGuarantee = true
		}},
		{name: "ha certified claim blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.HACertifiedClaim = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.HAReadinessState != DeploymentMultiTenantValCHAReadinessStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected blocked HA readiness state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValCRecoveryReadinessBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValCFoundation)
	}{
		{name: "backup freshness evidence missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.BackupFreshnessEvidence = ""
		}},
		{name: "stale backup evidence blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.BackupEvidenceFreshnessState = "stale"
		}},
		{name: "restore test evidence missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.RestoreTestEvidence = ""
		}},
		{name: "tenant scoped restore test missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.TenantScopedRestoreTest = ""
		}},
		{name: "restore integrity hash missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.RestoreIntegrityHash = ""
		}},
		{name: "encrypted backup custody reference missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.EncryptedBackupCustodyReference = ""
		}},
		{name: "dr drill evidence missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.DRDrillEvidence = ""
		}},
		{name: "restore target boundary missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.RestoreTargetBoundary = ""
		}},
		{name: "backup retention disposal boundary missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.BackupRetentionClass = ""
		}},
		{name: "disposal deletion boundary missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.DisposalDeletionBoundary = ""
		}},
		{name: "trailing whitespace evidence ref blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			last := len(model.RecoveryReadiness.EvidenceRefs) - 1
			model.RecoveryReadiness.EvidenceRefs[last] = model.RecoveryReadiness.EvidenceRefs[last] + " "
		}},
		{name: "rpo rto guarantee blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.RPORTOTreatedAsGuarantee = true
		}},
		{name: "backup exists treated as ready blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.BackupExistsTreatedAsReady = true
		}},
		{name: "restore guaranteed blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.RestoreGuaranteedClaim = true
		}},
		{name: "dr guaranteed blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.DRGuaranteedClaim = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.RecoveryReadinessState != DeploymentMultiTenantValCRecoveryReadinessStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected blocked recovery readiness state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValCSLAReadinessBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValCFoundation)
	}{
		{name: "supportability evidence missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.SupportabilityEvidence = ""
		}},
		{name: "alert routing evidence missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.AlertRoutingEvidence = ""
		}},
		{name: "incident escalation path missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.IncidentEscalationPath = ""
		}},
		{name: "support scope missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.SupportScope = ""
		}},
		{name: "support scope bare lower-case blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.SupportScope = "support_scope"
		}},
		{name: "support scope global blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.SupportScope = "global_support_scope"
		}},
		{name: "support scope prefixed tenant token blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.SupportScope = "ops " + deploymentMultiTenantVal0TenantScope()
		}},
		{name: "support scope suffixed tenant token blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.SupportScope = deploymentMultiTenantVal0TenantScope() + " support_scope"
		}},
		{name: "support scope unscoped blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.SupportScope = "unscoped_operator_access"
		}},
		{name: "tab retagged evidence ref blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.EvidenceRefs[0] = "\t" + model.SLAReadiness.EvidenceRefs[0]
		}},
		{name: "known limitations missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.KnownLimitations = ""
		}},
		{name: "no uptime guarantee wording missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.NoUptimeGuaranteeWordingPresent = false
		}},
		{name: "sla readiness treated as uptime guarantee blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.SLAReadinessTreatedAsUptimeGuarantee = true
		}},
		{name: "monitoring summary treated as canonical truth blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.MonitoringSummaryCanonicalTruth = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.SLAReadinessState != DeploymentMultiTenantValCSLAReadinessStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected blocked SLA readiness state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValCTenantTrustScopeBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValCFoundation)
	}{
		{name: "tenant trust scope missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustScope = ""
		}},
		{name: "issuer trust ownership missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.IssuerTrustOwnership = ""
		}},
		{name: "verification boundary missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.VerificationBoundary = ""
		}},
		{name: "key custody reference missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.KeyCustodyReference = ""
		}},
		{name: "key custody owner missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.KeyCustodyOwner = ""
		}},
		{name: "rotation behavior missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.RotationBehavior = ""
		}},
		{name: "rotation state unknown blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.RotationState = "unknown"
		}},
		{name: "rotation state stale blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.RotationState = "stale"
		}},
		{name: "offboarding transfer behavior missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.OffboardingTransferBehavior = ""
		}},
		{name: "revocation behavior missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.RevocationBehavior = ""
		}},
		{name: "trust export boundary missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustExportBoundary = ""
		}},
		{name: "trust scope inferred from dashboard view only blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.DashboardViewOnly = true
		}},
		{name: "trust scope inferred from fleet view only blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.FleetViewOnly = true
		}},
		{name: "bare lower-case tenant trust scope blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustScope = "tenant_trust_scope_evidence"
		}},
		{name: "global trust scope value blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustScope = "global"
		}},
		{name: "prefixed tenant trust scope blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustScope = "ops " + deploymentMultiTenantVal0TenantScope()
		}},
		{name: "suffixed tenant trust scope blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustScope = deploymentMultiTenantVal0TenantScope() + " trust_scope"
		}},
		{name: "unscoped trust scope value blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustScope = "unscoped"
		}},
		{name: "cross tenant trust scope value blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustScope = "cross_tenant_scope"
		}},
		{name: "all tenants trust scope value blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustScope = "all_tenants_scope"
		}},
		{name: "wildcard trust scope value blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustScope = "wildcard_tenant_scope"
		}},
		{name: "ish trust scope value blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TenantTrustScope = "tenant_scope_active_ish"
		}},
		{name: "newline retagged evidence ref blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.EvidenceRefs[0] = model.TenantTrustScope.EvidenceRefs[0] + "\n"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.TenantTrustScopeState != DeploymentMultiTenantValCTenantTrustScopeStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected blocked tenant trust scope state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValCSiloVisibilityBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValCFoundation)
	}{
		{name: "evidence silo validation missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.EvidenceSiloValidation = false
		}},
		{name: "audit silo validation missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.AuditSiloValidation = false
		}},
		{name: "export silo validation missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.ExportSiloValidation = false
		}},
		{name: "support visibility boundary missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.SupportVisibilityBoundary = ""
		}},
		{name: "support visibility exceeds tenant scope blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.SupportVisibilityExceedsTenantScope = true
		}},
		{name: "raw evidence exposed through support surface blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.RawEvidenceExposedThroughSupportSurface = true
		}},
		{name: "redaction hides missing decisive evidence blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.RedactionHidesMissingDecisiveEvidence = true
		}},
		{name: "redaction strengthens claim blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.RedactionStrengthensClaim = true
		}},
		{name: "projection surface treated as canonical truth blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.ProjectionSurfaceCanonicalTruth = true
		}},
		{name: "export silo missing exact identity blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.ExportSiloExactIdentity = ""
		}},
		{name: "prefixed whitespace evidence ref blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.EvidenceRefs[0] = " " + model.SiloVisibility.EvidenceRefs[0]
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.SiloVisibilityState != DeploymentMultiTenantValCSiloVisibilityStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected blocked silo visibility state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValCPrivacyGuardBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValCFoundation)
	}{
		{name: "cross tenant privacy guard evidence missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.CrossTenantPrivacyGuardEvidence = ""
		}},
		{name: "volume leakage check missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.VolumeLeakageCheckPresent = false
		}},
		{name: "error leakage check missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.ErrorLeakageCheckPresent = false
		}},
		{name: "timing leakage check missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.TimingLeakageCheckPresent = false
		}},
		{name: "aggregation leakage check missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.AggregationLeakageCheckPresent = false
		}},
		{name: "side channel marked safe without evidence blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.SideChannelMarkedSafeWithoutEvidence = true
		}},
		{name: "tenant private metadata leakage blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.TenantPrivateMetadataLeakage = true
		}},
		{name: "fleet aggregation treated as canonical truth blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.FleetAggregationCanonicalTruth = true
		}},
		{name: "bounded aggregation rules missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.BoundedAggregationRules = ""
		}},
		{name: "support export privacy boundary missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.SupportExportPrivacyVisibilityBoundary = ""
		}},
		{name: "suffixed whitespace evidence ref blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			last := len(model.PrivacyGuard.EvidenceRefs) - 1
			model.PrivacyGuard.EvidenceRefs[last] = model.PrivacyGuard.EvidenceRefs[last] + " "
		}},
		{name: "side channel negative tests missing blocks", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.VolumeLeakageNegativeTestPresent = false
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.PrivacyGuardState != DeploymentMultiTenantValCPrivacyGuardStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected blocked privacy guard state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValCEvidenceTokenBlockers(t *testing.T) {
	testCases := []struct {
		name          string
		mutate        func(*DeploymentMultiTenantValCFoundation)
		surface       string
		assertBlocked func(DeploymentMultiTenantValCFoundation) bool
	}{
		{
			name: "duplicate ha topology evidence blocks",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.HAReadiness.TopologyEvidence = "duplicate_ha_topology_evidence"
			},
			surface: DeploymentMultiTenantValCClosureSurfaceHAReadiness,
			assertBlocked: func(model DeploymentMultiTenantValCFoundation) bool {
				return model.HAReadinessState == DeploymentMultiTenantValCHAReadinessStateBlocked
			},
		},
		{
			name: "revoked restore test evidence blocks",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.RecoveryReadiness.RestoreTestEvidence = "revoked_restore_test_evidence"
			},
			surface: DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness,
			assertBlocked: func(model DeploymentMultiTenantValCFoundation) bool {
				return model.RecoveryReadinessState == DeploymentMultiTenantValCRecoveryReadinessStateBlocked
			},
		},
		{
			name: "expired backup freshness evidence blocks",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.RecoveryReadiness.BackupFreshnessEvidence = "expired_backup_freshness_evidence"
			},
			surface: DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness,
			assertBlocked: func(model DeploymentMultiTenantValCFoundation) bool {
				return model.RecoveryReadinessState == DeploymentMultiTenantValCRecoveryReadinessStateBlocked
			},
		},
		{
			name: "unrelated supportability evidence blocks",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.SLAReadiness.SupportabilityEvidence = "unrelated_supportability_evidence"
			},
			surface: DeploymentMultiTenantValCClosureSurfaceSLAReadiness,
			assertBlocked: func(model DeploymentMultiTenantValCFoundation) bool {
				return model.SLAReadinessState == DeploymentMultiTenantValCSLAReadinessStateBlocked
			},
		},
		{
			name: "revoked tenant trust scope evidence blocks",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.TenantTrustScope.IssuerTrustOwnership = "revoked_tenant_trust_scope_evidence"
			},
			surface: DeploymentMultiTenantValCClosureSurfaceTenantTrustScope,
			assertBlocked: func(model DeploymentMultiTenantValCFoundation) bool {
				return model.TenantTrustScopeState == DeploymentMultiTenantValCTenantTrustScopeStateBlocked
			},
		},
		{
			name: "expired key custody reference blocks",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.TenantTrustScope.KeyCustodyReference = "expired_key_custody_reference"
			},
			surface: DeploymentMultiTenantValCClosureSurfaceTenantTrustScope,
			assertBlocked: func(model DeploymentMultiTenantValCFoundation) bool {
				return model.TenantTrustScopeState == DeploymentMultiTenantValCTenantTrustScopeStateBlocked
			},
		},
		{
			name: "duplicate export silo identity blocks",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.SiloVisibility.ExportSiloExactIdentity = "duplicate_export_silo_identity"
			},
			surface: DeploymentMultiTenantValCClosureSurfaceSiloVisibility,
			assertBlocked: func(model DeploymentMultiTenantValCFoundation) bool {
				return model.SiloVisibilityState == DeploymentMultiTenantValCSiloVisibilityStateBlocked
			},
		},
		{
			name: "unrelated privacy guard evidence blocks",
			mutate: func(model *DeploymentMultiTenantValCFoundation) {
				model.PrivacyGuard.CrossTenantPrivacyGuardEvidence = "unrelated_privacy_guard_evidence"
			},
			surface: DeploymentMultiTenantValCClosureSurfacePrivacyGuard,
			assertBlocked: func(model DeploymentMultiTenantValCFoundation) bool {
				return model.PrivacyGuardState == DeploymentMultiTenantValCPrivacyGuardStateBlocked
			},
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if !tc.assertBlocked(model) {
			t.Fatalf("%s: expected affected substate to be blocked, got %#v", tc.name, model)
		}
		if model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected overall Val C state to be blocked, got %#v", tc.name, model)
		}
		if !deploymentMultiTenantValCHasFinding(model.ClosureBlockerOverlay.Findings, DeploymentMultiTenantValCBlockerLevelCLB1, tc.surface, "revoked expired duplicate or unrelated evidence token accepted") {
			t.Fatalf("%s: expected CL-B1 evidence-token finding, got %#v", tc.name, model.ClosureBlockerOverlay)
		}
	}
}

func TestDeploymentMultiTenantValCClosureBlockerOverlayCLB0AndCLB1Blockers(t *testing.T) {
	testCases := []struct {
		name    string
		mutate  func(*DeploymentMultiTenantValCFoundation)
		level   string
		surface string
		reason  string
	}{
		{name: "sla readiness treated as uptime guarantee produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.SLAReadinessTreatedAsUptimeGuarantee = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceSLAReadiness, reason: "sla readiness treated as uptime guarantee"},
		{name: "guaranteed uptime claim produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.GuaranteedUptimeClaim = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceSLAReadiness, reason: "guaranteed uptime claim"},
		{name: "zero downtime claim produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.ZeroDowntimeClaim = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceSLAReadiness, reason: "zero downtime claim"},
		{name: "production sla approved claim produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.ProductionSLAApprovedClaim = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceSLAReadiness, reason: "production sla approved claim"},
		{name: "recovery bypasses tenant isolation produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.BackupRestoreBypassesTenantIsolation = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, reason: "backup restore or dr readiness bypasses tenant isolation"},
		{name: "recovery bypasses data residency produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.BackupRestoreBypassesDataResidency = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, reason: "backup restore or dr readiness bypasses data residency"},
		{name: "support visibility leaks raw tenant evidence produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.RawEvidenceExposedThroughSupportSurface = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceSiloVisibility, reason: "support visibility leaks raw tenant evidence"},
		{name: "redaction hides decisive missing evidence produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.RedactionHidesMissingDecisiveEvidence = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceSiloVisibility, reason: "redaction hides decisive missing evidence"},
		{name: "redaction strengthens claim produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.RedactionStrengthensClaim = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceSiloVisibility, reason: "redaction strengthens claim"},
		{name: "side channel risk marked safe without evidence produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.SideChannelMarkedSafeWithoutEvidence = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfacePrivacyGuard, reason: "side-channel privacy risk marked safe without evidence"},
		{name: "tenant private metadata leakage produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.TenantPrivateMetadataLeakage = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfacePrivacyGuard, reason: "tenant-private metadata leakage"},
		{name: "projection surface treated as canonical truth produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.ProjectionSurfaceCanonicalTruth = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceSiloVisibility, reason: "projection surface treated as canonical truth"},
		{name: "copied competitor artifact produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.NoOverclaim.CleanRoomIPViolationDetected = true
		}, level: DeploymentMultiTenantValCBlockerLevelCLB0, surface: DeploymentMultiTenantValCClosureSurfaceCleanRoomIP, reason: "copied competitor deployment recovery or privacy artifact detected"},
		{name: "backup evidence missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.BackupFreshnessEvidence = ""
		}, level: DeploymentMultiTenantValCBlockerLevelCLB1, surface: DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, reason: "backup or restore evidence missing"},
		{name: "restore test evidence missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.RestoreTestEvidence = ""
		}, level: DeploymentMultiTenantValCBlockerLevelCLB1, surface: DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, reason: "restore test evidence missing"},
		{name: "stale backup evidence handling not proven produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.StaleBackupEvidenceHandlingProven = false
		}, level: DeploymentMultiTenantValCBlockerLevelCLB1, surface: DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, reason: "stale backup evidence handling not proven"},
		{name: "ha readiness not evidence linked produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.HAReadinessEvidenceLinked = false
		}, level: DeploymentMultiTenantValCBlockerLevelCLB1, surface: DeploymentMultiTenantValCClosureSurfaceHAReadiness, reason: "ha readiness not evidence-linked"},
		{name: "failover test evidence missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.FailoverTestEvidence = ""
		}, level: DeploymentMultiTenantValCBlockerLevelCLB1, surface: DeploymentMultiTenantValCClosureSurfaceHAReadiness, reason: "failover test evidence missing"},
		{name: "tenant trust scope missing issuer or trust ownership produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.IssuerTrustOwnership = ""
		}, level: DeploymentMultiTenantValCBlockerLevelCLB1, surface: DeploymentMultiTenantValCClosureSurfaceTenantTrustScope, reason: "tenant trust scope missing issuer or trust ownership"},
		{name: "key custody rotation behavior missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.RotationBehavior = ""
		}, level: DeploymentMultiTenantValCBlockerLevelCLB1, surface: DeploymentMultiTenantValCClosureSurfaceTenantTrustScope, reason: "key custody rotation or revocation behavior missing"},
		{name: "silo validation missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SiloVisibility.EvidenceSiloValidation = false
		}, level: DeploymentMultiTenantValCBlockerLevelCLB1, surface: DeploymentMultiTenantValCClosureSurfaceSiloVisibility, reason: "evidence audit or export silo validation missing"},
		{name: "privacy side channel negative tests missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.VolumeLeakageNegativeTestPresent = false
		}, level: DeploymentMultiTenantValCBlockerLevelCLB1, surface: DeploymentMultiTenantValCClosureSurfacePrivacyGuard, reason: "privacy side-channel negative tests missing"},
		{name: "dependency gate not exact active produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.Dependency.ValBCurrentState = "partial"
		}, level: DeploymentMultiTenantValCBlockerLevelCLB1, surface: DeploymentMultiTenantValCClosureSurfaceHAReadiness, reason: "dependency gate missing or not exact active"},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.ClosureBlockerState != DeploymentMultiTenantValCClosureBlockerStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected blocked closure blocker state, got %#v", tc.name, model)
		}
		if !deploymentMultiTenantValCHasFinding(model.ClosureBlockerOverlay.Findings, tc.level, tc.surface, tc.reason) {
			t.Fatalf("%s: expected finding %s/%s/%s, got %#v", tc.name, tc.level, tc.surface, tc.reason, model.ClosureBlockerOverlay)
		}
	}
}

func TestDeploymentMultiTenantValCClosureBlockerOverlayCLB2Cleanup(t *testing.T) {
	testCases := []struct {
		name    string
		mutate  func(*DeploymentMultiTenantValCFoundation)
		surface string
		reason  string
	}{
		{name: "ambiguous ha profile naming is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.HAProfileNamingExact = false
		}, surface: DeploymentMultiTenantValCClosureSurfaceHAReadiness, reason: "ambiguous ha profile naming"},
		{name: "ambiguous recovery target naming is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.RecoveryTargetNamingExact = false
		}, surface: DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, reason: "ambiguous recovery target naming"},
		{name: "ambiguous trust scope naming is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.TrustScopeNamingExact = false
		}, surface: DeploymentMultiTenantValCClosureSurfaceTenantTrustScope, reason: "ambiguous trust scope naming"},
		{name: "missing safe wording example for sla readiness is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.SafeSLAReadinessWordingExamplePresent = false
		}, surface: DeploymentMultiTenantValCClosureSurfaceSLAReadiness, reason: "missing safe wording example for sla readiness"},
		{name: "missing safe wording example for ha readiness is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.SafeHAReadinessWordingExamplePresent = false
		}, surface: DeploymentMultiTenantValCClosureSurfaceHAReadiness, reason: "missing safe wording example for ha readiness"},
		{name: "incomplete diagnostic output for ha blockers is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.HAReadiness.DiagnosticOutputComplete = false
		}, surface: DeploymentMultiTenantValCClosureSurfaceHAReadiness, reason: "incomplete diagnostic output for ha blockers"},
		{name: "incomplete diagnostic output for recovery blockers is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.DiagnosticOutputComplete = false
		}, surface: DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, reason: "incomplete diagnostic output for recovery blockers"},
		{name: "incomplete diagnostic output for sla blockers is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.SLAReadiness.DiagnosticOutputComplete = false
		}, surface: DeploymentMultiTenantValCClosureSurfaceSLAReadiness, reason: "incomplete diagnostic output for sla blockers"},
		{name: "incomplete diagnostic output for trust blockers is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.TenantTrustScope.DiagnosticOutputComplete = false
		}, surface: DeploymentMultiTenantValCClosureSurfaceTenantTrustScope, reason: "incomplete diagnostic output for trust blockers"},
		{name: "incomplete diagnostic output for privacy blockers is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.PrivacyGuard.DiagnosticOutputComplete = false
		}, surface: DeploymentMultiTenantValCClosureSurfacePrivacyGuard, reason: "incomplete diagnostic output for privacy blockers"},
		{name: "incomplete bounded runbook wording without direct pass bypass is cleanup", mutate: func(model *DeploymentMultiTenantValCFoundation) {
			model.RecoveryReadiness.RunbookWordingComplete = false
		}, surface: DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, reason: "incomplete bounded runbook wording without direct pass bypass"},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValCModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.ClosureBlockerState != DeploymentMultiTenantValCClosureBlockerStateCleanup || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected cleanup closure blocker state and blocked current state, got %#v", tc.name, model)
		}
		if !deploymentMultiTenantValCHasFinding(model.ClosureBlockerOverlay.Findings, DeploymentMultiTenantValCBlockerLevelCLB2, tc.surface, tc.reason) {
			t.Fatalf("%s: expected cleanup finding %s/%s, got %#v", tc.name, tc.surface, tc.reason, model.ClosureBlockerOverlay)
		}
	}
}

func TestDeploymentMultiTenantValCClosureBlockerOverlayCLB3Advisory(t *testing.T) {
	advisory := DeploymentMultiTenantValCClosureBlockerOverlay{
		ProjectionDisclaimer: deploymentMultiTenantValCProjectionDisclaimer(),
		Findings: []DeploymentMultiTenantValCClosureBlockerFinding{
			{
				BlockerLevel:      DeploymentMultiTenantValCBlockerLevelCLB3,
				Surface:           DeploymentMultiTenantValCClosureSurfacePrivacyGuard,
				Reason:            "advisory cleanup carried forward",
				BlocksCurrentWave: false,
				RequiredFollowup:  "record the advisory cleanup if it is carried forward",
			},
		},
	}
	if got := EvaluateDeploymentMultiTenantValCClosureBlockerState(advisory); got != DeploymentMultiTenantValCClosureBlockerStateAdvisory {
		t.Fatalf("expected advisory closure blocker state, got %s", got)
	}

	mixed := DeploymentMultiTenantValCClosureBlockerOverlay{
		ProjectionDisclaimer: deploymentMultiTenantValCProjectionDisclaimer(),
		Findings: []DeploymentMultiTenantValCClosureBlockerFinding{
			{
				BlockerLevel:      DeploymentMultiTenantValCBlockerLevelCLB3,
				Surface:           DeploymentMultiTenantValCClosureSurfacePrivacyGuard,
				Reason:            "advisory cleanup carried forward",
				BlocksCurrentWave: false,
				RequiredFollowup:  "record the advisory cleanup if it is carried forward",
			},
			{
				BlockerLevel:      DeploymentMultiTenantValCBlockerLevelCLB0,
				Surface:           DeploymentMultiTenantValCClosureSurfaceSLAReadiness,
				Reason:            "guaranteed uptime claim",
				BlocksCurrentWave: true,
				RequiredFollowup:  "remove guaranteed uptime claim",
			},
		},
	}
	if got := EvaluateDeploymentMultiTenantValCClosureBlockerState(mixed); got != DeploymentMultiTenantValCClosureBlockerStateBlocked {
		t.Fatalf("expected advisory to not mask stronger blocker, got %s", got)
	}

	model := activeDeploymentMultiTenantValCModel()
	model.ClosureBlockerState = DeploymentMultiTenantValCClosureBlockerStateAdvisory
	if got := EvaluateDeploymentMultiTenantValCState(model); got != DeploymentMultiTenantValCStateBlocked {
		t.Fatalf("expected advisory closure blocker state to block final Val C state, got %s", got)
	}
}

func TestDeploymentMultiTenantValCClosureBlockerOverlayRejectsLegacyAndUnknownLevels(t *testing.T) {
	testCases := []struct {
		name    string
		finding DeploymentMultiTenantValCClosureBlockerFinding
	}{
		{
			name: "legacy priority zero is rejected",
			finding: DeploymentMultiTenantValCClosureBlockerFinding{
				BlockerLevel:      deploymentMultiTenantLegacyPriority("0"),
				Surface:           DeploymentMultiTenantValCClosureSurfaceHAReadiness,
				Reason:            "legacy severity rejected",
				BlocksCurrentWave: true,
			},
		},
		{
			name: "legacy priority one is rejected",
			finding: DeploymentMultiTenantValCClosureBlockerFinding{
				BlockerLevel:      deploymentMultiTenantLegacyPriority("1"),
				Surface:           DeploymentMultiTenantValCClosureSurfaceHAReadiness,
				Reason:            "legacy severity rejected",
				BlocksCurrentWave: true,
				RequiredFollowup:  "use cl b blocker terminology",
			},
		},
		{
			name: "legacy priority two is rejected",
			finding: DeploymentMultiTenantValCClosureBlockerFinding{
				BlockerLevel:      deploymentMultiTenantLegacyPriority("2"),
				Surface:           DeploymentMultiTenantValCClosureSurfaceHAReadiness,
				Reason:            "legacy severity rejected",
				BlocksCurrentWave: true,
				RequiredFollowup:  "use cl b blocker terminology",
			},
		},
		{
			name: "unknown blocker level is rejected",
			finding: DeploymentMultiTenantValCClosureBlockerFinding{
				BlockerLevel:      "CL-B9",
				Surface:           DeploymentMultiTenantValCClosureSurfaceHAReadiness,
				Reason:            "unknown blocker level",
				BlocksCurrentWave: true,
			},
		},
		{
			name: "unknown surface is rejected",
			finding: DeploymentMultiTenantValCClosureBlockerFinding{
				BlockerLevel:      DeploymentMultiTenantValCBlockerLevelCLB0,
				Surface:           "unknown_surface",
				Reason:            "unknown surface",
				BlocksCurrentWave: true,
			},
		},
	}

	for _, tc := range testCases {
		overlay := DeploymentMultiTenantValCClosureBlockerOverlay{
			ProjectionDisclaimer: deploymentMultiTenantValCProjectionDisclaimer(),
			Findings:             []DeploymentMultiTenantValCClosureBlockerFinding{tc.finding},
		}
		if got := EvaluateDeploymentMultiTenantValCClosureBlockerState(overlay); got != DeploymentMultiTenantValCClosureBlockerStateBlocked {
			t.Fatalf("%s: expected blocked closure blocker state, got %s", tc.name, got)
		}
	}
}

func TestDeploymentMultiTenantValCNoOverclaimBlockers(t *testing.T) {
	blockedClaims := []string{
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
		"SLA readiness means production approval",
		"supportability evidence means SLA guarantee",
		"supportability evidence review means SLA guarantee",
		"tenant trust certified",
		"tenant trust scope certified",
		"key custody certified",
		"data residency certified",
		"tenant isolation guaranteed",
		"privacy guaranteed",
		"no side-channel leakage guaranteed",
		"fleet aggregation proves privacy",
		"support visibility cannot leak",
		"redaction proves safe",
		"portal view is canonical truth",
		"portal view projection is canonical truth",
		"dashboard proves recovery readiness",
		"dashboard routine proves recovery readiness",
		"clean-room certified",
		"patent cleared",
		"FTO cleared",
		"legal certification",
		"copied competitor workflow",
		"HA readiness evidence guaranteed uptime",
		"restore test evidence restore guaranteed",
		"SLA readiness evidence SLA guaranteed",
		"supportability evidence production SLA approved",
		"privacy guard evidence privacy guaranteed",
		"side-channel negative test no side-channel leakage guaranteed",
		"bounded support visibility support visibility cannot leak",
		"tenant trust scope evidence tenant trust certified",
	}
	allowedClaims := []string{
		"HA readiness evidence",
		"failover test evidence",
		"dependency degradation behavior",
		"degraded mode semantics",
		"backup freshness evidence",
		"restore test evidence",
		"tenant-scoped restore test",
		"restore integrity hash",
		"encrypted backup custody reference",
		"DR drill evidence",
		"RPO/RTO target",
		"SLA readiness evidence",
		"supportability evidence",
		"known limitations",
		"tenant trust scope evidence",
		"issuer/trust ownership evidence",
		"key/custody rotation evidence",
		"evidence silo validation",
		"audit silo validation",
		"export silo validation",
		"bounded support visibility",
		"privacy guard evidence",
		"side-channel negative test",
		"bounded aggregation rules",
		"advisory fleet visibility",
		"not uptime guarantee",
		"not production approval",
		"not deployment approval",
		"not compliance certification",
		"not canonical truth",
	}

	for _, claim := range blockedClaims {
		model := activeDeploymentMultiTenantValCModel()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValCNoOverclaimStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%s: expected blocked no-overclaim state, got %#v", claim, model)
		}
	}

	splitBlockedClaims := [][]string{
		{"guaranteed", "uptime"},
		{"portal view", "is canonical truth"},
		{"supportability evidence", "means SLA guarantee"},
	}
	for _, claims := range splitBlockedClaims {
		model := activeDeploymentMultiTenantValCModel()
		model.NoOverclaim.ObservedClaims = claims
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValCNoOverclaimStateBlocked || model.CurrentState != DeploymentMultiTenantValCStateBlocked {
			t.Fatalf("%q: expected blocked split no-overclaim state, got %#v", claims, model)
		}
	}

	for _, claim := range allowedClaims {
		model := activeDeploymentMultiTenantValCModel()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantValCFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValCNoOverclaimStateActive || model.CurrentState != DeploymentMultiTenantValCStateActive {
			t.Fatalf("%s: expected active no-overclaim state, got %#v", claim, model)
		}
	}
}

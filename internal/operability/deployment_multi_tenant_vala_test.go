package operability

import (
	"encoding/json"
	"strings"
	"testing"
)

func activeDeploymentMultiTenantValAModel() DeploymentMultiTenantValAFoundation {
	model := DeploymentMultiTenantValAFoundationModel()
	return ComputeDeploymentMultiTenantValAFoundation(model)
}

func deploymentMultiTenantValAHasFinding(findings []DeploymentMultiTenantValAPassBlockerFinding, severity, surface, reason string) bool {
	for _, finding := range findings {
		if finding.Severity == severity &&
			finding.Surface == surface &&
			strings.Contains(finding.Reason, reason) &&
			finding.BlocksCurrentValAPass {
			return true
		}
	}
	return false
}

func deploymentMultiTenantLegacyPriority(level string) string {
	return "P" + level
}

func TestDeploymentMultiTenantValAHappyPathAndPoint10NotComplete(t *testing.T) {
	model := activeDeploymentMultiTenantValAModel()
	if model.CurrentState != DeploymentMultiTenantValAStateActive {
		t.Fatalf("expected active Val A state, got %#v", model)
	}
	if model.DependencyState != DeploymentMultiTenantValADependencyStateActive {
		t.Fatalf("expected active Val 0 dependency, got %#v", model)
	}
	if model.PassBlockerState != DeploymentMultiTenantValAPassBlockerStateActive {
		t.Fatalf("expected clean pass blocker overlay, got %#v", model)
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		t.Fatalf("expected point 10 to remain not complete, got %#v", model)
	}
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal model: %v", err)
	}
	if strings.Contains(string(payload), "point_"+"10_pass") {
		t.Fatalf("expected Val A to never emit point 10 pass, got %s", string(payload))
	}
}

func TestDeploymentMultiTenantValAAggregateProjectionDisclaimerBlocks(t *testing.T) {
	model := activeDeploymentMultiTenantValAModel()
	model.ProjectionDisclaimer = "canonical_truth"
	model = ComputeDeploymentMultiTenantValAFoundation(model)
	if model.CurrentState != DeploymentMultiTenantValAStateBlocked {
		t.Fatalf("expected malformed aggregate projection disclaimer to block ValA state, got %#v", model)
	}
	if !containsTrimmedString(model.BlockingReasons, "aggregate_projection_disclaimer_blocked") {
		t.Fatalf("expected aggregate projection disclaimer blocking reason, got %#v", model.BlockingReasons)
	}
}

func TestDeploymentMultiTenantValAAirGappedUnsupportedDependencySemantics(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*DeploymentMultiTenantValAFoundation)
		wantState string
	}{
		{
			name: "no unsupported dependencies happy path stays ready",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"none_explicit"}
				model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateReady
			},
			wantState: DeploymentMultiTenantValAStateActive,
		},
		{
			name: "explicit unsupported dependencies accepted as degraded",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{
					"dependency_rekor_online_lookup",
					"dependency_external_advisory_sync",
				}
				model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
			},
			wantState: DeploymentMultiTenantValAStateActive,
		},
		{
			name: "explicit unsupported dependencies accepted as unsupported",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{
					"dependency_rekor_online_lookup",
				}
				model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateUnsupported
			},
			wantState: DeploymentMultiTenantValAStateActive,
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValAModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.DeploymentProfileMatrixState != DeploymentMultiTenantValADeploymentProfileMatrixStateActive {
			t.Fatalf("%s: expected active deployment profile matrix, got %#v", tc.name, model)
		}
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
	}
}

func TestDeploymentMultiTenantValADependencyBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValAFoundation)
	}{
		{name: "val0 current state partial blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.Dependency.Val0CurrentState = DeploymentMultiTenantVal0StateBlocked
		}},
		{name: "val0 dependency state blocked blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.Dependency.Val0DependencyState = DeploymentMultiTenantVal0DependencyStateBlocked
		}},
		{name: "val0 future contract state blocked blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.Dependency.Val0FutureContractState = DeploymentMultiTenantVal0FutureContractStateBlocked
		}},
		{name: "val0 no overclaim state blocked blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.Dependency.Val0NoOverclaimState = DeploymentMultiTenantVal0NoOverclaimStateBlocked
		}},
		{name: "point10 state complete blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.Dependency.Point10State = "deployment_multi_tenant_point_10_complete"
		}},
		{name: "malformed projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.Dependency.ProjectionDisclaimer = "canonical_truth"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValAModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.DependencyState != DeploymentMultiTenantValADependencyStateBlocked || model.CurrentState != DeploymentMultiTenantValAStateBlocked {
			t.Fatalf("%s: expected blocked dependency, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValADependencySnapshotCopiesAggregateComputedVal0ProjectionDisclaimer(t *testing.T) {
	val0 := ComputeDeploymentMultiTenantVal0Foundation(DeploymentMultiTenantVal0FoundationModel())
	val0.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_val0 aggregate_dependency_snapshot"
	val0.NoOverclaim.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_val0 component_no_overclaim"
	snapshot := DeploymentMultiTenantValADependencySnapshot{
		Val0CurrentState:              val0.CurrentState,
		Val0DependencyState:           val0.DependencyState,
		Val0DeploymentValidationState: val0.DeploymentValidationState,
		Val0TenantBoundaryState:       val0.TenantBoundaryState,
		Val0MSPAuthorityState:         val0.MSPAuthorityState,
		Val0PolicyEnvelopeState:       val0.PolicyEnvelopeState,
		Val0TenantTrustScopeState:     val0.TenantTrustScopeState,
		Val0ConnectorContractState:    val0.ConnectorContractState,
		Val0OperatorActionState:       val0.OperatorActionState,
		Val0PrivacyGuardState:         val0.PrivacyGuardState,
		Val0FairShareState:            val0.FairShareState,
		Val0OperationalPreflightState: val0.OperationalPreflightState,
		Val0FutureContractState:       val0.FutureContractState,
		Val0NoOverclaimState:          val0.NoOverclaimState,
		Point10State:                  val0.Point10State,
		ProjectionDisclaimer:          val0.ProjectionDisclaimer,
	}
	if snapshot.ProjectionDisclaimer != val0.ProjectionDisclaimer {
		t.Fatalf("expected aggregate Val0 disclaimer to propagate exactly, got snapshot=%q val0=%q", snapshot.ProjectionDisclaimer, val0.ProjectionDisclaimer)
	}
	if snapshot.ProjectionDisclaimer == val0.NoOverclaim.ProjectionDisclaimer {
		t.Fatalf("expected dependency snapshot not to fallback to component disclaimer, got snapshot=%q component=%q", snapshot.ProjectionDisclaimer, val0.NoOverclaim.ProjectionDisclaimer)
	}
	if EvaluateDeploymentMultiTenantValADependencyState(snapshot) != DeploymentMultiTenantValADependencyStateActive {
		t.Fatalf("expected copied aggregate disclaimer to keep dependency active, got %#v", snapshot)
	}

	val0.ProjectionDisclaimer = "canonical_truth"
	snapshot.ProjectionDisclaimer = val0.ProjectionDisclaimer
	if EvaluateDeploymentMultiTenantValADependencyState(snapshot) != DeploymentMultiTenantValADependencyStateBlocked {
		t.Fatalf("expected malformed aggregate disclaimer to block dependency without component fallback, got %#v", snapshot)
	}
}

func TestDeploymentMultiTenantValADeploymentProfileMatrixBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValAFoundation)
	}{
		{name: "missing saas tenant config blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SaaSTenantConfig = ""
		}},
		{name: "missing saas region blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SaaSRegion = ""
		}},
		{name: "missing identity bootstrap blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SaaSIdentityBootstrap = ""
		}},
		{name: "missing connector scope blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SaaSConnectorScope = ""
		}},
		{name: "missing backup policy blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SaaSBackupPolicy = ""
		}},
		{name: "missing operator support scope blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SaaSOperatorSupportScope = ""
		}},
		{name: "missing self hosted artifact provenance blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SelfHostedArtifactProvenance = ""
		}},
		{name: "missing environment manifest blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SelfHostedEnvironmentManifest = ""
		}},
		{name: "missing config validation blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SelfHostedConfigValidation = ""
		}},
		{name: "missing iam kms dependency blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SelfHostedIAMKMSDependency = ""
		}},
		{name: "missing backup target blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SelfHostedBackupTarget = ""
		}},
		{name: "missing upgrade rollback plan blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SelfHostedUpgradeRollbackPlan = ""
		}},
		{name: "missing air gapped offline artifact bundle blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedOfflineArtifactBundle = ""
		}},
		{name: "missing air gapped offline evidence bundle blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedOfflineEvidenceBundle = ""
		}},
		{name: "missing signature hash verification blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedSignatureHashVerificationState = DeploymentMultiTenantValASignatureVerificationUnknown
		}},
		{name: "hidden unsupported dependency list blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedDependenciesHidden = true
		}},
		{name: "missing offline replay export path blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedOfflineReplayExportPath = ""
		}},
		{name: "explicit unsupported dependencies cannot be ready", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"dependency_rekor_online_lookup"}
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateReady
		}},
		{name: "install success without readiness evidence blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.ReadinessEvidenceBacked = false
		}},
		{name: "marketplace install treated as ready blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.MarketplaceInstallTreatedAsReady = true
		}},
		{name: "unsupported profile ready blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.UnsupportedProfileReady = true
		}},
		{name: "unknown profile blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.SupportedProfiles = []string{"unknown"}
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValAModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.DeploymentProfileMatrixState != DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked || model.CurrentState != DeploymentMultiTenantValAStateBlocked {
			t.Fatalf("%s: expected blocked deployment profile matrix, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValAUnsupportedDependencyListBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValAFoundation)
	}{
		{name: "silently ready unsupported dependencies still block", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"dependency_rekor_online_lookup"}
			model.DeploymentProfileMatrix.AirGappedUnsupportedDependenciesSilentlyReady = true
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
		}},
		{name: "mixed sentinel and explicit dependency blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"none_explicit", "dependency_rekor_online_lookup"}
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
		}},
		{name: "nil dependency list blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = nil
		}},
		{name: "empty dependency list blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{}
		}},
		{name: "unknown dependency id blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"unknown"}
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
		}},
		{name: "partial dependency id blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"partial"}
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
		}},
		{name: "incomplete dependency id blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"incomplete"}
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
		}},
		{name: "stale dependency id blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"stale"}
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
		}},
		{name: "malformed dependency id blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"malformed"}
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
		}},
		{name: "blocked dependency id blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"blocked"}
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
		}},
		{name: "empty dependency id blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{""}
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
		}},
		{name: "whitespace dependency id blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{" "}
			model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateDegraded
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValAModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.DeploymentProfileMatrixState != DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked || model.CurrentState != DeploymentMultiTenantValAStateBlocked {
			t.Fatalf("%s: expected blocked unsupported dependency handling, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValAPreflightBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValAFoundation)
	}{
		{name: "install config validation missing blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.InstallConfigValidation = false
		}},
		{name: "upgrade config diff missing blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.UpgradeConfigDiff = false
		}},
		{name: "db schema migration dry run missing blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.DBSchemaMigrationDryRun = false
		}},
		{name: "backup before upgrade evidence missing blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.BackupBeforeUpgradeEvidence = false
		}},
		{name: "rollback plan evidence missing blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.RollbackPlanEvidence = false
		}},
		{name: "policy migration compatibility missing blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.PolicyMigrationCompatibility = false
		}},
		{name: "connector permission change review missing blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.ConnectorPermissionChangeReview = false
		}},
		{name: "key rotation readiness missing blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.KeyRotationReadiness = false
		}},
		{name: "tenant boundary validation missing blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.TenantBoundaryValidation = false
		}},
		{name: "preflight not tenant scoped blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.TenantScope = "global_admin_scope"
		}},
		{name: "production impacting change safe by default blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.PreflightGate.ProductionImpactSafeByDefault = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValAModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.PreflightGateState != DeploymentMultiTenantValAPreflightGateStateBlocked || model.CurrentState != DeploymentMultiTenantValAStateBlocked {
			t.Fatalf("%s: expected blocked preflight gate, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValAIdentityBootstrapBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValAFoundation)
	}{
		{name: "missing issuer entity id blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.IssuerEntityID = ""
		}},
		{name: "missing callback redirect url blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.CallbackRedirectURL = ""
		}},
		{name: "certificate expired blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.CertificateExpiryState = DeploymentMultiTenantValACertificateStateExpired
		}},
		{name: "certificate expiry unknown blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.CertificateExpiryState = DeploymentMultiTenantValACertificateStateUnknown
		}},
		{name: "missing group role mapping blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.GroupRoleMapping = ""
		}},
		{name: "unsafe fallback allowed blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.UnsafeFallbackEnabled = true
		}},
		{name: "break glass without expiry blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.BreakGlassExpiryPresent = false
		}},
		{name: "break glass without revocation blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.BreakGlassRevocationPresent = false
		}},
		{name: "missing tenant identity boundary blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.TenantSpecificIdentityBoundary = ""
		}},
		{name: "sso configured means secure blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.SSOConfiguredMeansSecureClaim = true
		}},
		{name: "identity readiness implies deployment readiness blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.IdentityBootstrap.IdentityReadinessImpliesDeploymentReadiness = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValAModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.IdentityBootstrapState != DeploymentMultiTenantValAIdentityBootstrapStateBlocked || model.CurrentState != DeploymentMultiTenantValAStateBlocked {
			t.Fatalf("%s: expected blocked identity bootstrap, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValAAirGappedBundleBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValAFoundation)
	}{
		{name: "missing bundle manifest blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.BundleManifest = ""
		}},
		{name: "missing artifact hashes blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.ArtifactHashes = ""
		}},
		{name: "missing proof pack hashes blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.ProofPackHashes = ""
		}},
		{name: "missing signer blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.Signer = ""
		}},
		{name: "missing policy version blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.PolicyVersion = ""
		}},
		{name: "missing engine version blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.EngineVersion = ""
		}},
		{name: "missing timestamp blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.Timestamp = ""
		}},
		{name: "unsupported online dependencies hidden blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.UnsupportedOnlineDependenciesHidden = true
		}},
		{name: "missing replay instructions blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.ReplayInstructions = ""
		}},
		{name: "missing offline replay export path blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.OfflineReplayExportPath = ""
		}},
		{name: "signature hash verification failed blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.SignatureHashVerificationState = DeploymentMultiTenantValASignatureVerificationFailed
		}},
		{name: "air gapped certified blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.AirGappedCertifiedClaim = true
		}},
		{name: "air gapped means fully offline verified blocks", mutate: func(model *DeploymentMultiTenantValAFoundation) {
			model.AirGappedEvidenceBundle.AirGappedMeansFullyOfflineVerifiedClaim = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValAModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.AirGappedEvidenceBundleState != DeploymentMultiTenantValAAirGappedEvidenceBundleStateBlocked || model.CurrentState != DeploymentMultiTenantValAStateBlocked {
			t.Fatalf("%s: expected blocked air-gapped bundle, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValANoOverclaimBlockers(t *testing.T) {
	blockedClaims := []string{
		"production approved",
		"deployment approved",
		"marketplace certified",
		"marketplace production ready",
		"one-click secure",
		"install success means ready",
		"marketplace install means ready",
		"marketplace install means production ready",
		"customer ready without validation",
		"compliant by default",
		"compliance guaranteed",
		"regulator-approved",
		"self-hosted production approved",
		"air-gapped certified",
		"air-gapped means fully offline verified",
		"sso secure by default",
		"sso configured means secure",
		"rbac complete by default",
		"deployment readiness guaranteed",
		"unsupported profile ready",
		"offline bundle certified",
		"offline replay guarantees correctness",
		"preflight passed means production approved",
		"rollback guaranteed",
		"zero-risk deployment",
		"guaranteed uptime",
		"sla guaranteed",
		"validated deployment profile production approved",
		"air-gapped offline evidence bundle air-gapped certified",
		"sso bootstrap validation sso secure by default",
		"tenant-scoped preflight preflight passed means production approved",
		"bounded marketplace deployment profile marketplace production ready",
		"offline replay/export path offline replay guarantees correctness",
	}

	for _, claim := range blockedClaims {
		model := activeDeploymentMultiTenantValAModel()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValANoOverclaimStateBlocked || model.CurrentState != DeploymentMultiTenantValAStateBlocked {
			t.Fatalf("expected blocked no-overclaim for %q, got %#v", claim, model)
		}
	}

	allowedClaims := []string{
		"validated deployment profile",
		"evidence-linked readiness state",
		"bounded marketplace deployment profile",
		"self-hosted readiness evidence",
		"air-gapped offline evidence bundle",
		"offline replay/export path",
		"sso bootstrap validation",
		"saml/oidc validation evidence",
		"tenant-scoped preflight",
		"rollback plan evidence",
		"backup-before-upgrade evidence",
		"unsupported dependency explicitly listed",
		"degraded deployment state",
		"incomplete deployment state",
		"not production approval",
		"not deployment approval",
		"not compliance certification",
		"not canonical truth",
	}

	for _, claim := range allowedClaims {
		model := activeDeploymentMultiTenantValAModel()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValANoOverclaimStateActive {
			t.Fatalf("expected allowed bounded wording for %q, got %#v", claim, model)
		}
	}
}

func TestDeploymentMultiTenantValAPassBlockerOverlayCLB0AndCLB1Blockers(t *testing.T) {
	testCases := []struct {
		name     string
		mutate   func(*DeploymentMultiTenantValAFoundation)
		severity string
		surface  string
		reason   string
	}{
		{
			name: "install success treated as readiness blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.InstallSuccessTreatedAsReady = true
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB0,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:   "install success treated as readiness",
		},
		{
			name: "marketplace install treated as readiness blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.MarketplaceInstallTreatedAsReady = true
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB0,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:   "marketplace install treated as readiness",
		},
		{
			name: "marketplace install treated as production readiness blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.MarketplaceInstallTreatedAsProductionReady = true
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB0,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:   "marketplace install treated as production readiness",
		},
		{
			name: "sso configured means secure blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.IdentityBootstrap.SSOConfiguredMeansSecureClaim = true
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB0,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceIdentityBootstrap,
			reason:   "sso configured treated as secure",
		},
		{
			name: "sso readiness treated as deployment readiness blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.IdentityBootstrap.IdentityReadinessImpliesDeploymentReadiness = true
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB0,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceIdentityBootstrap,
			reason:   "sso readiness treated as deployment readiness",
		},
		{
			name: "unsafe fallback enabled blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.IdentityBootstrap.UnsafeFallbackEnabled = true
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB0,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceIdentityBootstrap,
			reason:   "unsafe fallback enabled",
		},
		{
			name: "marketplace overclaim in deployment wording blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.ObservedClaims = []string{"bounded marketplace deployment profile marketplace production ready"}
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB0,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:   "marketplace or msp overclaim in deployment profile wording",
		},
		{
			name: "self hosted unsupported degraded semantics missing blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.SelfHostedUnsupportedSemanticsExplicit = false
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB1,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:   "self-hosted profile lacks unsupported or degraded semantics",
		},
		{
			name: "air gapped unsupported degraded semantics missing blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.AirGappedDegradedSemanticsExplicit = false
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB1,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:   "air-gapped profile lacks unsupported or degraded semantics",
		},
		{
			name: "unsupported air gapped dependency hidden blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.AirGappedUnsupportedDependenciesHidden = true
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB1,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:   "unsupported air-gapped dependency hidden or silently treated as ready",
		},
		{
			name: "explicit unsupported air gapped dependency treated as ready blocks",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies = []string{"dependency_rekor_online_lookup"}
				model.DeploymentProfileMatrix.AirGappedState = DeploymentMultiTenantDeploymentStateReady
			},
			severity: deploymentMultiTenantValAPassBlockerSeverityCLB1,
			surface:  deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:   "explicit unsupported air-gapped dependency treated as ready",
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValAModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.PassBlockerState != DeploymentMultiTenantValAPassBlockerStateBlocked || model.CurrentState != DeploymentMultiTenantValAStateBlocked {
			t.Fatalf("%s: expected blocked pass blocker overlay, got %#v", tc.name, model)
		}
		if !deploymentMultiTenantValAHasFinding(model.PassBlockerOverlay.Findings, tc.severity, tc.surface, tc.reason) {
			t.Fatalf("%s: expected finding %s/%s/%s, got %#v", tc.name, tc.severity, tc.surface, tc.reason, model.PassBlockerOverlay.Findings)
		}
	}
}

func TestDeploymentMultiTenantValAPassBlockerOverlayCLB2Cleanup(t *testing.T) {
	testCases := []struct {
		name    string
		mutate  func(*DeploymentMultiTenantValAFoundation)
		surface string
		reason  string
	}{
		{
			name: "ambiguous deployment profile naming is reported as cleanup",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.ProfileNamingExact = false
			},
			surface: deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:  "ambiguous deployment profile naming",
		},
		{
			name: "missing safe wording example is reported as cleanup",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.SafeReadinessWordingExamplePresent = false
			},
			surface: deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:  "missing safe wording example for deployment readiness",
		},
		{
			name: "incomplete readiness diagnostic output is reported as cleanup",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.DeploymentProfileMatrix.DiagnosticOutputComplete = false
			},
			surface: deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile,
			reason:  "incomplete diagnostic output for readiness blockers",
		},
		{
			name: "incomplete preflight diagnostic output is reported as cleanup",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.PreflightGate.DiagnosticOutputComplete = false
			},
			surface: deploymentMultiTenantValAPassBlockerSurfacePreflight,
			reason:  "incomplete diagnostic output for preflight blockers",
		},
		{
			name: "incomplete identity diagnostic output is reported as cleanup",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.IdentityBootstrap.DiagnosticOutputComplete = false
			},
			surface: deploymentMultiTenantValAPassBlockerSurfaceIdentityBootstrap,
			reason:  "incomplete diagnostic output for identity blockers",
		},
		{
			name: "incomplete air gapped diagnostic output is reported as cleanup",
			mutate: func(model *DeploymentMultiTenantValAFoundation) {
				model.AirGappedEvidenceBundle.DiagnosticOutputComplete = false
			},
			surface: deploymentMultiTenantValAPassBlockerSurfaceAirGappedBundle,
			reason:  "incomplete diagnostic output for air-gapped blockers",
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValAModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValAFoundation(model)
		if model.PassBlockerState != DeploymentMultiTenantValAPassBlockerStateCleanup || model.CurrentState != DeploymentMultiTenantValAStateBlocked {
			t.Fatalf("%s: expected cleanup pass blocker overlay and blocked Val A state, got %#v", tc.name, model)
		}
		if !deploymentMultiTenantValAHasFinding(model.PassBlockerOverlay.Findings, deploymentMultiTenantValAPassBlockerSeverityCLB2, tc.surface, tc.reason) {
			t.Fatalf("%s: expected cleanup finding %s/%s, got %#v", tc.name, tc.surface, tc.reason, model.PassBlockerOverlay.Findings)
		}
	}
}

func TestDeploymentMultiTenantValAPassBlockerOverlayUsesCanonicalCLBTaxonomy(t *testing.T) {
	model := activeDeploymentMultiTenantValAModel()
	model.DeploymentProfileMatrix.InstallSuccessTreatedAsReady = true
	model.DeploymentProfileMatrix.ProfileNamingExact = false
	model = ComputeDeploymentMultiTenantValAFoundation(model)
	if model.PassBlockerState != DeploymentMultiTenantValAPassBlockerStateBlocked {
		t.Fatalf("expected blocked pass blocker overlay, got %#v", model)
	}
	if !deploymentMultiTenantValAHasFinding(model.PassBlockerOverlay.Findings, deploymentMultiTenantValAPassBlockerSeverityCLB0, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "install success treated as readiness") {
		t.Fatalf("expected CL-B0 critical closure blocker finding, got %#v", model.PassBlockerOverlay.Findings)
	}
	if !deploymentMultiTenantValAHasFinding(model.PassBlockerOverlay.Findings, deploymentMultiTenantValAPassBlockerSeverityCLB2, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "ambiguous deployment profile naming") {
		t.Fatalf("expected CL-B2 cleanup finding, got %#v", model.PassBlockerOverlay.Findings)
	}
	for _, finding := range model.PassBlockerOverlay.Findings {
		if strings.HasPrefix(strings.TrimSpace(finding.Severity), "P") {
			t.Fatalf("expected canonical CL-B severity taxonomy only, got %#v", model.PassBlockerOverlay.Findings)
		}
	}
}

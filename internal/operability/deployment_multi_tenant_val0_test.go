package operability

import (
	"encoding/json"
	"strings"
	"testing"
)

func activeDeploymentMultiTenantVal0Model() DeploymentMultiTenantVal0Foundation {
	model := DeploymentMultiTenantVal0FoundationModel()
	return ComputeDeploymentMultiTenantVal0Foundation(model)
}

func TestDeploymentMultiTenantVal0HappyPathAndPoint10NotComplete(t *testing.T) {
	model := activeDeploymentMultiTenantVal0Model()
	if model.CurrentState != DeploymentMultiTenantVal0StateActive {
		t.Fatalf("expected active To\u010dka 10 Val 0 state, got %#v", model)
	}
	if model.DependencyState != DeploymentMultiTenantVal0DependencyStateActive {
		t.Fatalf("expected active To\u010dka 9 / Val E dependency, got %#v", model)
	}
	if model.FutureContractState != DeploymentMultiTenantVal0FutureContractStateActive {
		t.Fatalf("expected active Val 0 future contract foundation, got %#v", model)
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		t.Fatalf("expected point 10 to remain not complete, got %#v", model)
	}
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal model: %v", err)
	}
	if strings.Contains(string(payload), "point_"+"10_pass") {
		t.Fatalf("expected Val 0 to never emit point 10 pass, got %s", string(payload))
	}
}

func TestDeploymentMultiTenantVal0DependencyBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "vale current state partial blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ValECurrentState = OSSTrustNetworkValEStatePartial
		}},
		{name: "point 9 state not complete blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.Point9State = OSSTrustNetworkPoint9StateNotComplete
		}},
		{name: "point 9 pass allowed false blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.Point9PassAllowed = false
		}},
		{name: "wrong point 9 pass reason blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.Point9PassReason = OSSTrustNetworkValEPoint9PassReasonBlocked
		}},
		{name: "vale dependency state partial blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ValEDependencyState = OSSTrustNetworkValEDependencyStatePartial
		}},
		{name: "vale final pass rule partial blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ValEFinalPassRule = OSSTrustNetworkValEFinalPassRuleStatePartial
		}},
		{name: "vale no overclaim partial blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ValENoOverclaimState = OSSTrustNetworkValENoOverclaimStatePartial
		}},
		{name: "missing vale current state blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ValECurrentState = ""
		}},
		{name: "unknown vale proof refs block", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ProofSurfaceRefs = []string{"unknown"}
		}},
		{name: "missing vale evidence refs block", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.EvidenceRefs = nil
		}},
		{name: "malformed projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ProjectionDisclaimer = "canonical_truth"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.DependencyState != DeploymentMultiTenantVal0DependencyStateBlocked {
			t.Fatalf("%s: expected blocked dependency, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0DeploymentBlockers(t *testing.T) {
	testCases := []struct {
		name                string
		mutate              func(*DeploymentMultiTenantVal0Foundation)
		wantState           string
		wantDisciplineState string
	}{
		{
			name: "install success without readiness evidence is blocked",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.DeploymentValidation.ValidationEvidenceBacked = false
			},
			wantState:           DeploymentMultiTenantVal0StateBlocked,
			wantDisciplineState: DeploymentMultiTenantVal0DeploymentValidationStateBlocked,
		},
		{
			name: "marketplace install without validation is blocked",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.DeploymentValidation.ExplicitReadinessValidation = false
			},
			wantState:           DeploymentMultiTenantVal0StateBlocked,
			wantDisciplineState: DeploymentMultiTenantVal0DeploymentValidationStateBlocked,
		},
		{
			name: "unknown deployment state is blocked",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.DeploymentValidation.DeploymentState = DeploymentMultiTenantDeploymentStateUnknown
			},
			wantState:           DeploymentMultiTenantVal0StateBlocked,
			wantDisciplineState: DeploymentMultiTenantVal0DeploymentValidationStateBlocked,
		},
		{
			name: "incomplete deployment state is blocked",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.DeploymentValidation.DeploymentState = DeploymentMultiTenantDeploymentStateIncomplete
			},
			wantState:           DeploymentMultiTenantVal0StateBlocked,
			wantDisciplineState: DeploymentMultiTenantVal0DeploymentValidationStateBlocked,
		},
		{
			name: "unsupported deployment profile is not ready",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.DeploymentValidation.DeploymentProfile = "unsupported_profile"
			},
			wantState:           DeploymentMultiTenantVal0StateBlocked,
			wantDisciplineState: DeploymentMultiTenantVal0DeploymentValidationStateUnsupported,
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != tc.wantState {
			t.Fatalf("%s: expected %s, got %#v", tc.name, tc.wantState, model)
		}
		if model.DeploymentValidationState != tc.wantDisciplineState {
			t.Fatalf("%s: expected deployment validation state %s, got %#v", tc.name, tc.wantDisciplineState, model)
		}
	}
}

func TestDeploymentMultiTenantVal0TenantBoundaryBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "missing audit boundary is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantBoundary.AuditBoundary = ""
		}},
		{name: "missing evidence boundary is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantBoundary.EvidenceBoundary = ""
		}},
		{name: "missing export boundary is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantBoundary.ExportBoundary = ""
		}},
		{name: "fleet dashboard summary alone is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantBoundary.DashboardSummaryOnly = true
		}},
		{name: "tenant scope unknown blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantBoundary.TenantScope = "unknown"
		}},
		{name: "audit boundary partial blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantBoundary.AuditBoundary = "partial"
		}},
		{name: "evidence boundary stale blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantBoundary.EvidenceBoundary = "stale"
		}},
		{name: "export boundary unsupported blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantBoundary.ExportBoundary = "unsupported"
		}},
		{name: "credential boundary malformed blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantBoundary.CredentialBoundary = "malformed"
		}},
		{name: "operator support boundary active ish blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantBoundary.OperatorSupportBoundary = "active-ish"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.TenantBoundaryState != DeploymentMultiTenantVal0TenantBoundaryStateBlocked {
			t.Fatalf("%s: expected blocked tenant boundary, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0MSPAuthorityBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "partner source of truth is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.MSPAuthority.PartnerSourceOfTruth = true
		}},
		{name: "msp source of truth is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.MSPAuthority.MSPSourceOfTruth = true
		}},
		{name: "unscoped operator access is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.MSPAuthority.OperatorAccessScoped = false
		}},
		{name: "missing revocation path is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.MSPAuthority.RevocationPathPresent = false
		}},
		{name: "msp role scope unknown blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.MSPAuthority.RoleScope = "unknown"
		}},
		{name: "msp role scope global blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.MSPAuthority.RoleScope = "global"
		}},
		{name: "msp tenant scope unscoped blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.MSPAuthority.TenantScope = "unscoped"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.MSPAuthorityState != DeploymentMultiTenantVal0MSPAuthorityStateBlocked {
			t.Fatalf("%s: expected blocked MSP authority, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0PolicyEnvelopeBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "dangerous relaxation is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.PolicyEnvelope.DangerousRelaxation = true
		}},
		{name: "silent conflict resolution is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.PolicyEnvelope.SilentConflictResolution = true
		}},
		{name: "unknown inheritance is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.PolicyEnvelope.UnknownInheritance = true
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.PolicyEnvelopeState != DeploymentMultiTenantVal0PolicyEnvelopeStateBlocked {
			t.Fatalf("%s: expected blocked policy envelope, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0TenantTrustScopeBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "ambiguous shared trust scope is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantTrustScope.SharedAmbiguousScope = true
		}},
		{name: "missing issuer trust ownership is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantTrustScope.TrustOwner = ""
		}},
		{name: "missing offboarding revocation behavior is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantTrustScope.OffboardingBehavior = ""
		}},
		{name: "unknown trust scope blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.TenantTrustScope.TrustScope = "unknown"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.TenantTrustScopeState != DeploymentMultiTenantVal0TenantTrustScopeStateBlocked {
			t.Fatalf("%s: expected blocked tenant trust scope, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0ConnectorBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "cross tenant connector capability is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.ConnectorContract.CrossTenantAccess = true
		}},
		{name: "undeclared mutation capability is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.ConnectorContract.UndeclaredMutationCapability = true
		}},
		{name: "missing connector audit trail is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.ConnectorContract.AuditTrailPresent = false
		}},
		{name: "connector capability unknown blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.ConnectorContract.Capabilities = []string{"unknown"}
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.ConnectorContractState != DeploymentMultiTenantVal0ConnectorContractStateBlocked {
			t.Fatalf("%s: expected blocked connector contract, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0PrivacyBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "raw cross tenant evidence sharing is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.PrivacyGuard.RawCrossTenantEvidenceShare = true
		}},
		{name: "fleet view treated as canonical truth is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.PrivacyGuard.FleetViewCanonicalTruth = true
		}},
		{name: "side channel risk without evidence is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.PrivacyGuard.SideChannelEvidenceLinked = false
		}},
		{name: "cross tenant global privacy scope blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.PrivacyGuard.TenantPrivacyScope = "cross_tenant_global_scope"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.PrivacyGuardState != DeploymentMultiTenantVal0PrivacyGuardStateBlocked {
			t.Fatalf("%s: expected blocked privacy guard, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0FairShareBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "one tenant can starve another tenant is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FairShare.OneTenantCanStarveAnother = true
		}},
		{name: "alert flood spilling across tenants is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FairShare.AlertFloodSpillsAcross = true
		}},
		{name: "overload treated as ready is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FairShare.OverloadTreatedAsReady = true
		}},
		{name: "all tenants admin resource scope blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FairShare.TenantResourceScope = "all_tenants_admin"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.FairShareState != DeploymentMultiTenantVal0FairShareStateBlocked {
			t.Fatalf("%s: expected blocked fair share, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0OperationalPreflightBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "upgrade without tenant scoped preflight is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.OperationalPreflight.UpgradeTenantScoped = false
		}},
		{name: "key rotation without tenant scoped preflight is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.OperationalPreflight.KeyRotationTenantScoped = false
		}},
		{name: "connector change without tenant scoped preflight is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.OperationalPreflight.ConnectorChangeTenantScoped = false
		}},
		{name: "support access activation without validation is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.OperationalPreflight.SupportAccessActivationValidated = false
		}},
		{name: "support access revocation without validation is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.OperationalPreflight.SupportAccessRevocationValidated = false
		}},
		{name: "wildcard tenant scope blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.OperationalPreflight.TenantChangeScope = "wildcard_tenant_scope"
		}},
		{name: "tenant scope active ish blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.OperationalPreflight.TenantChangeScope = "tenant_scope_active_ish"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.OperationalPreflightState != DeploymentMultiTenantVal0OperationalPreflightStateBlocked {
			t.Fatalf("%s: expected blocked operational preflight, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0OperatorScopeBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "operator scope global blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.OperatorAction.Scope = "global"
		}},
		{name: "operator scope unscoped blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.OperatorAction.Scope = "unscoped"
		}},
		{name: "operator unscoped access blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.OperatorAction.Scope = "operator-unscoped-access"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.OperatorActionState != DeploymentMultiTenantVal0OperatorActionStateBlocked {
			t.Fatalf("%s: expected blocked operator action, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0ValidScopedValuesRemainAllowed(t *testing.T) {
	model := activeDeploymentMultiTenantVal0Model()
	model.TenantBoundary.TenantScope = "tenant_scope_evidence_linked"
	model.TenantBoundary.AuditBoundary = "tenant_audit_boundary"
	model.TenantBoundary.EvidenceBoundary = "tenant_evidence_boundary"
	model.TenantBoundary.ExportBoundary = "tenant_export_boundary"
	model.TenantBoundary.OperatorSupportBoundary = "tenant_operator_support_boundary"
	model.MSPAuthority.RoleScope = "bounded_tenant_support_scope"
	model.TenantTrustScope.VerificationBoundary = "tenant_specific_identity_boundary"
	model.ConnectorContract.Capabilities = []string{"tenant_scoped_connector_execution"}
	model.OperatorAction.Scope = "tenant_scoped_operator_action"
	model.PrivacyGuard.TenantPrivacyScope = "tenant_privacy_scope_evidence_linked"
	model.FairShare.TenantResourceScope = "tenant_resource_scope_bounded"
	model.OperationalPreflight.TenantChangeScope = "tenant_change_scope_preflight"
	model = ComputeDeploymentMultiTenantVal0Foundation(model)
	if model.CurrentState != DeploymentMultiTenantVal0StateActive {
		t.Fatalf("expected valid scoped values to remain active, got %#v", model)
	}
}

func TestDeploymentMultiTenantVal0FutureContractBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantVal0Foundation)
	}{
		{name: "missing readiness matrix contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.ReadinessEvidenceMatrixPresent = false
		}},
		{name: "missing preflight gate contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.InstallConfigValidationPresent = false
		}},
		{name: "missing sso bootstrap validator contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.IssuerEntityIDPresent = false
		}},
		{name: "missing air gapped offline evidence bundle contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.OfflineEvidenceBundleContractPresent = false
		}},
		{name: "missing backup restore dr readiness contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.BackupFreshnessEvidencePresent = false
		}},
		{name: "missing ha sla no overclaim contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.TopologyEvidencePresent = false
		}},
		{name: "missing tenant isolation data residency test pack contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.CrossTenantAuditLeakageTestPresent = false
		}},
		{name: "missing fair share quota contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.EventBudgetPerTenantPresent = false
		}},
		{name: "missing connector capability manifest contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.ConnectorCapabilityManifestPresent = false
		}},
		{name: "missing msp marketplace authority boundary contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.MSPDeploySupportTenantScopedOnly = false
		}},
		{name: "missing tenant lifecycle contract blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.FutureContract.TenantCreatePresent = false
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.FutureContractState != DeploymentMultiTenantVal0FutureContractStateBlocked {
			t.Fatalf("%s: expected blocked future contract state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0ScopeVariantRegressionBlockers(t *testing.T) {
	testCases := []struct {
		name            string
		mutate          func(*DeploymentMultiTenantVal0Foundation)
		disciplineName  string
		disciplineState func(DeploymentMultiTenantVal0Foundation) string
	}{
		{
			name: "tenant boundary global admin scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantBoundary.TenantScope = "global_admin_scope"
			},
			disciplineName: "tenant boundary",
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.TenantBoundaryState
			},
		},
		{
			name: "tenant boundary global hyphen scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantBoundary.AuditBoundary = "global-admin-scope"
			},
			disciplineName: "tenant boundary",
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.TenantBoundaryState
			},
		},
		{
			name: "msp tenant global scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.MSPAuthority.TenantScope = "tenant_global_scope"
			},
			disciplineName: "msp authority",
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.MSPAuthorityState
			},
		},
		{
			name: "msp unscoped operator access blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.MSPAuthority.RoleScope = "unscoped_operator_access"
			},
			disciplineName: "msp authority",
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.MSPAuthorityState
			},
		},
		{
			name: "tenant trust operator unscoped access blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantTrustScope.TrustScope = "operator-unscoped-access"
			},
			disciplineName: "tenant trust scope",
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.TenantTrustScopeState
			},
		},
		{
			name: "connector cross tenant global scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.ConnectorContract.TenantScope = "cross_tenant_global_scope"
			},
			disciplineName: "connector contract",
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.ConnectorContractState
			},
		},
		{
			name: "operator all tenants admin blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.OperatorAction.Scope = "all_tenants_admin"
			},
			disciplineName: "operator action",
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.OperatorActionState
			},
		},
		{
			name: "privacy any tenant support blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.PrivacyGuard.TenantPrivacyScope = "any_tenant_support"
			},
			disciplineName: "privacy guard",
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.PrivacyGuardState
			},
		},
		{
			name: "fair share wildcard tenant scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.FairShare.TenantResourceScope = "wildcard_tenant_scope"
			},
			disciplineName: "fair share",
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.FairShareState
			},
		},
		{
			name: "operational preflight tenant scope active ish blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.OperationalPreflight.TenantChangeScope = "tenant_scope_active_ish"
			},
			disciplineName: "operational preflight",
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.OperationalPreflightState
			},
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if state := tc.disciplineState(model); state == "" || !strings.HasSuffix(state, "_blocked") {
			t.Fatalf("%s: expected blocked %s discipline, got %#v", tc.name, tc.disciplineName, model)
		}
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("%s: expected overall blocked state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0NoOverclaimBlockers(t *testing.T) {
	blockedClaims := []string{
		"production approved",
		"deployment approved",
		"marketplace certified",
		"msp certified",
		"regulator-approved",
		"compliance guaranteed",
		"compliant by default",
		"one-click secure",
		"zero-risk deployment",
		"tenant safe by default",
		"globally trusted msp",
		"official marketplace trust authority",
		"partner approved",
		"customer ready without validation",
		"compliance-as-a-service",
		"certified managed trust",
		"universal trust score",
		"global trust score",
		"trust score > 90",
		"deployment passed means secure",
		"install success means ready",
		"fleet view is canonical truth",
		"partner source of truth",
		"msp source of truth",
		"cross-tenant safe without evidence",
		"automated production rollout",
		"automatic rollback approval",
		"operator fully trusted",
		"guaranteed uptime",
		"zero downtime",
		"sla guaranteed",
		"production sla approved",
		"ha certified",
		"air-gapped certified",
		"self-hosted production approved",
		"sso secure by default",
		"rbac complete by default",
		"restore guaranteed",
		"dr guaranteed",
		"tenant isolation guaranteed",
		"data residency certified",
		"marketplace production ready",
		"msp approved deployment",
		"partner certified deployment",
		"backup exists means ready",
		"failover configured means ready",
		"healthcheck green means fully ready",
		"sso configured means secure",
		"air-gapped means fully offline verified",
		"bounded marketplace deployment profile production approved",
		"tenant-scoped operational model operator fully trusted",
		"advisory fleet visibility fleet view is canonical truth",
		"HA readiness evidence guaranteed uptime",
		"SLA readiness evidence SLA guaranteed",
		"offline evidence bundle air-gapped certified",
		"tenant-scoped operational model tenant isolation guaranteed",
		"bounded marketplace deployment profile marketplace production ready",
		"supportability evidence production SLA approved",
	}
	for _, claim := range blockedClaims {
		model := activeDeploymentMultiTenantVal0Model()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked || model.NoOverclaimState != DeploymentMultiTenantVal0NoOverclaimStateBlocked {
			t.Fatalf("expected blocked no-overclaim for %q, got %#v", claim, model)
		}
	}

	allowedClaims := []string{
		"not production approval",
		"not deployment approval",
		"not compliance certification",
		"not canonical truth",
		"bounded marketplace deployment profile",
		"tenant-scoped operational model",
		"advisory fleet visibility",
		"evidence-linked readiness state",
		"bounded operator authority",
		"sandboxed connector execution",
		"HA readiness evidence",
		"SLA readiness evidence",
		"failover test evidence",
		"RPO/RTO target",
		"supportability evidence",
		"degraded mode behavior",
		"offline evidence bundle",
		"tenant-scoped restore test",
	}
	for _, claim := range allowedClaims {
		model := activeDeploymentMultiTenantVal0Model()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateActive || model.NoOverclaimState != DeploymentMultiTenantVal0NoOverclaimStateActive {
			t.Fatalf("expected allowed bounded claim %q to stay active, got %#v", claim, model)
		}
	}
}

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

func assertDeploymentMultiTenantVal0NoPoint10Pass(t *testing.T, model DeploymentMultiTenantVal0Foundation) {
	t.Helper()
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal model: %v", err)
	}
	if strings.Contains(string(payload), "point_"+"10_pass") {
		t.Fatalf("expected Val 0 to never emit point 10 pass, got %s", string(payload))
	}
}

func TestDeploymentMultiTenantVal0AggregateProjectionDisclaimerBlocks(t *testing.T) {
	model := activeDeploymentMultiTenantVal0Model()
	model.ProjectionDisclaimer = "canonical_truth"
	model = ComputeDeploymentMultiTenantVal0Foundation(model)
	if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
		t.Fatalf("expected malformed aggregate projection disclaimer to block Val0 state, got %#v", model)
	}
	if !containsTrimmedString(model.BlockingReasons, "aggregate_projection_disclaimer_blocked") {
		t.Fatalf("expected aggregate projection disclaimer blocking reason, got %#v", model.BlockingReasons)
	}
}

func TestDeploymentMultiTenantVal0CanonicalTenantTokenValueRequiresExactSingleToken(t *testing.T) {
	if !deploymentMultiTenantVal0CanonicalTenantTokenValueIsValid(deploymentMultiTenantVal0TenantScope()) {
		t.Fatalf("expected exact canonical tenant token to be valid")
	}

	invalid := []string{
		"",
		"ops " + deploymentMultiTenantVal0TenantScope(),
		deploymentMultiTenantVal0TenantScope() + " support_scope",
		deploymentMultiTenantVal0TenantScope() + "\tsupport_scope",
		deploymentMultiTenantVal0TenantScope() + "\nprofile",
	}
	for _, value := range invalid {
		if deploymentMultiTenantVal0CanonicalTenantTokenValueIsValid(value) {
			t.Fatalf("expected non-exact tenant token %q to be invalid", value)
		}
	}
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
		{name: "whitespace padded vale current state blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ValECurrentState = " " + OSSTrustNetworkValEStatePass + "\t"
		}},
		{name: "unknown vale proof refs block", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ProofSurfaceRefs = []string{"unknown"}
		}},
		{name: "newline padded point 9 state blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.Point9State = "\n" + OSSTrustNetworkPoint9StatePass
		}},
		{name: "whitespace padded point 9 pass reason blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.Point9PassReason = " " + OSSTrustNetworkValEPoint9PassReasonAllowed + "\t"
		}},
		{name: "whitespace padded vale proof refs block", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ProofSurfaceRefs = []string{" " + OSSTrustNetworkValEProofSurfaceRefs()[0] + "\t"}
		}},
		{name: "whitespace padded vale dependency state blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ValEDependencyState = " " + OSSTrustNetworkValEDependencyStateActive + "\t"
		}},
		{name: "whitespace padded vale final pass rule blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ValEFinalPassRule = " " + OSSTrustNetworkValEFinalPassRuleStateActive + "\t"
		}},
		{name: "whitespace padded vale no overclaim state blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ValENoOverclaimState = " " + OSSTrustNetworkValENoOverclaimStateActive + "\t"
		}},
		{name: "missing vale evidence refs block", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.EvidenceRefs = nil
		}},
		{name: "newline padded vale evidence refs block", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.EvidenceRefs = []string{"\n" + OSSTrustNetworkValEProofEvidenceRefs()[0]}
		}},
		{name: "malformed projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ProjectionDisclaimer = "canonical_truth"
		}},
		{name: "whitespace retagged aggregate dependency disclaimer blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ProjectionDisclaimer = " " + ossTrustNetworkValEProjectionDisclaimer() + " aggregate_dependency_snapshot"
		}},
		{name: "uppercase aggregate dependency disclaimer blocks", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.Dependency.ProjectionDisclaimer = strings.ToUpper(ossTrustNetworkValEProjectionDisclaimer() + " aggregate_dependency_snapshot")
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
			name: "whitespace padded deployment evidence ref is blocked",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.DeploymentValidation.EvidenceRefs = []string{" " + deploymentMultiTenantVal0DeploymentEvidenceRefs()[0]}
			},
			wantState:           DeploymentMultiTenantVal0StateBlocked,
			wantDisciplineState: DeploymentMultiTenantVal0DeploymentValidationStateBlocked,
		},
		{
			name: "whitespace padded deployment state is blocked",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.DeploymentValidation.DeploymentState = " " + DeploymentMultiTenantDeploymentStateReady + "\t"
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
		{
			name: "whitespace padded deployment profile is not ready",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.DeploymentValidation.DeploymentProfile = " " + DeploymentMultiTenantProfileBoundedMarketplaceMSP + "\t"
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
		{name: "whitespace padded conflict state is blocked", mutate: func(model *DeploymentMultiTenantVal0Foundation) {
			model.PolicyEnvelope.ConflictState = " " + DeploymentMultiTenantConflictStateNoConflict + "\t"
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

func TestDeploymentMultiTenantVal0CanonicalExactBindingsRemainAllowed(t *testing.T) {
	model := activeDeploymentMultiTenantVal0Model()
	model.TenantBoundary.TenantScope = deploymentMultiTenantVal0TenantScope()
	model.MSPAuthority.TenantScope = deploymentMultiTenantVal0TenantScope()
	model.ConnectorContract.TenantScope = deploymentMultiTenantVal0TenantScope()
	model.OperatorAction.TenantTarget = deploymentMultiTenantVal0OperatorTenantTarget()
	model.ConnectorContract.Capabilities = append([]string{}, model.ConnectorContract.Capabilities...)
	model.ConnectorContract.ReadBoundaries = append([]string{}, model.ConnectorContract.ReadBoundaries...)
	model.ConnectorContract.MutationBoundaries = append([]string{}, model.ConnectorContract.MutationBoundaries...)
	model.Dependency.ProofSurfaceRefs = append([]string{}, model.Dependency.ProofSurfaceRefs...)
	model.Dependency.EvidenceRefs = append([]string{}, model.Dependency.EvidenceRefs...)
	model = ComputeDeploymentMultiTenantVal0Foundation(model)
	if model.CurrentState != DeploymentMultiTenantVal0StateActive {
		t.Fatalf("expected canonical exact-bound values to remain active, got %#v", model)
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
		expectedState   string
		disciplineState func(DeploymentMultiTenantVal0Foundation) string
	}{
		{
			name: "tenant boundary global admin scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantBoundary.TenantScope = "global_admin_scope"
			},
			disciplineName: "tenant boundary",
			expectedState:  DeploymentMultiTenantVal0TenantBoundaryStateBlocked,
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
			expectedState:  DeploymentMultiTenantVal0TenantBoundaryStateBlocked,
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
			expectedState:  DeploymentMultiTenantVal0MSPAuthorityStateBlocked,
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
			expectedState:  DeploymentMultiTenantVal0MSPAuthorityStateBlocked,
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
			expectedState:  DeploymentMultiTenantVal0TenantTrustScopeStateBlocked,
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
			expectedState:  DeploymentMultiTenantVal0ConnectorContractStateBlocked,
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
			expectedState:  DeploymentMultiTenantVal0OperatorActionStateBlocked,
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
			expectedState:  DeploymentMultiTenantVal0PrivacyGuardStateBlocked,
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
			expectedState:  DeploymentMultiTenantVal0FairShareStateBlocked,
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
			expectedState:  DeploymentMultiTenantVal0OperationalPreflightStateBlocked,
			disciplineState: func(model DeploymentMultiTenantVal0Foundation) string {
				return model.OperationalPreflightState
			},
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if state := tc.disciplineState(model); state != tc.expectedState {
			t.Fatalf("%s: expected exact blocked %s discipline state %q, got %#v", tc.name, tc.disciplineName, tc.expectedState, model)
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

func TestDeploymentMultiTenantVal0ProjectionDisclaimerExactBoundedBlockers(t *testing.T) {
	disclaimer := deploymentMultiTenantVal0ProjectionDisclaimer()

	t.Run("aggregate disclaimer suffix drift blocks exact state", func(t *testing.T) {
		model := activeDeploymentMultiTenantVal0Model()
		model.ProjectionDisclaimer = disclaimer + " extra_scope"
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("expected aggregate disclaimer suffix drift to block current state, got %#v", model)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("expected point 10 to remain not complete after aggregate disclaimer drift, got %#v", model)
		}
		payload, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal model: %v", err)
		}
		if strings.Contains(string(payload), "point_"+"10_pass") {
			t.Fatalf("expected aggregate disclaimer drift not to emit point 10 pass, got %s", string(payload))
		}
	})

	t.Run("aggregate disclaimer leading whitespace blocks exact state", func(t *testing.T) {
		model := activeDeploymentMultiTenantVal0Model()
		model.ProjectionDisclaimer = " " + disclaimer
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("expected aggregate disclaimer leading whitespace to block current state, got %#v", model)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("expected point 10 to remain not complete after aggregate disclaimer leading whitespace, got %#v", model)
		}
		assertDeploymentMultiTenantVal0NoPoint10Pass(t, model)
	})

	t.Run("aggregate disclaimer uppercase retagging blocks exact state", func(t *testing.T) {
		model := activeDeploymentMultiTenantVal0Model()
		model.ProjectionDisclaimer = strings.ToUpper(disclaimer)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("expected aggregate disclaimer uppercase retagging to block current state, got %#v", model)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("expected point 10 to remain not complete after aggregate disclaimer uppercase retagging, got %#v", model)
		}
		assertDeploymentMultiTenantVal0NoPoint10Pass(t, model)
	})

	t.Run("canonical aggregate snapshot disclaimer blocks live foundation", func(t *testing.T) {
		model := activeDeploymentMultiTenantVal0Model()
		model.ProjectionDisclaimer = disclaimer + " aggregate_dependency_snapshot"
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("expected canonical aggregate snapshot disclaimer to block live current state, got %#v", model)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("expected point 10 to remain not complete after canonical aggregate snapshot disclaimer, got %#v", model)
		}
		assertDeploymentMultiTenantVal0NoPoint10Pass(t, model)
	})

	t.Run("discipline disclaimer prefix drift blocks exact no-overclaim state", func(t *testing.T) {
		model := activeDeploymentMultiTenantVal0Model()
		model.NoOverclaim.ProjectionDisclaimer = "advisory " + disclaimer
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("expected discipline disclaimer prefix drift to block current state, got %#v", model)
		}
		if model.NoOverclaimState != DeploymentMultiTenantVal0NoOverclaimStateBlocked {
			t.Fatalf("expected discipline disclaimer prefix drift to block no-overclaim state, got %#v", model)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("expected point 10 to remain not complete after discipline disclaimer drift, got %#v", model)
		}
	})

	t.Run("aggregate snapshot disclaimer dropping bounded scope blocks exact state", func(t *testing.T) {
		model := activeDeploymentMultiTenantVal0Model()
		model.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_val0 aggregate_dependency_snapshot"
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("expected short aggregate disclaimer to block current state, got %#v", model)
		}
		if !containsTrimmedString(model.BlockingReasons, "aggregate_projection_disclaimer_blocked") {
			t.Fatalf("expected short aggregate disclaimer to emit aggregate projection disclaimer blocking reason, got %#v", model.BlockingReasons)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("expected point 10 to remain not complete after short aggregate disclaimer, got %#v", model)
		}
	})
}

func TestDeploymentMultiTenantVal0DisciplineShortAggregateProjectionDisclaimerBlocksExactState(t *testing.T) {
	shortAggregateDisclaimer := "projection_only not_canonical_truth deployment_multi_tenant_val0 aggregate_dependency_snapshot"
	testCases := []struct {
		name      string
		mutate    func(*DeploymentMultiTenantVal0Foundation)
		getState  func(DeploymentMultiTenantVal0Foundation) string
		wantState string
	}{
		{
			name: "deployment validation discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.DeploymentValidation.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.DeploymentValidationState },
			wantState: DeploymentMultiTenantVal0DeploymentValidationStateBlocked,
		},
		{
			name: "tenant boundary discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantBoundary.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.TenantBoundaryState },
			wantState: DeploymentMultiTenantVal0TenantBoundaryStateBlocked,
		},
		{
			name: "msp authority discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.MSPAuthority.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.MSPAuthorityState },
			wantState: DeploymentMultiTenantVal0MSPAuthorityStateBlocked,
		},
		{
			name: "policy envelope discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.PolicyEnvelope.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.PolicyEnvelopeState },
			wantState: DeploymentMultiTenantVal0PolicyEnvelopeStateBlocked,
		},
		{
			name: "tenant trust scope discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantTrustScope.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.TenantTrustScopeState },
			wantState: DeploymentMultiTenantVal0TenantTrustScopeStateBlocked,
		},
		{
			name: "connector contract discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.ConnectorContract.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.ConnectorContractState },
			wantState: DeploymentMultiTenantVal0ConnectorContractStateBlocked,
		},
		{
			name: "operator action discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.OperatorAction.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.OperatorActionState },
			wantState: DeploymentMultiTenantVal0OperatorActionStateBlocked,
		},
		{
			name: "privacy guard discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.PrivacyGuard.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.PrivacyGuardState },
			wantState: DeploymentMultiTenantVal0PrivacyGuardStateBlocked,
		},
		{
			name: "fair share discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.FairShare.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.FairShareState },
			wantState: DeploymentMultiTenantVal0FairShareStateBlocked,
		},
		{
			name: "operational preflight discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.OperationalPreflight.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.OperationalPreflightState },
			wantState: DeploymentMultiTenantVal0OperationalPreflightStateBlocked,
		},
		{
			name: "future contract discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.FutureContract.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.FutureContractState },
			wantState: DeploymentMultiTenantVal0FutureContractStateBlocked,
		},
		{
			name: "no overclaim discipline",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.NoOverclaim.ProjectionDisclaimer = shortAggregateDisclaimer
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.NoOverclaimState },
			wantState: DeploymentMultiTenantVal0NoOverclaimStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("%s: expected blocked current state, got %#v", tc.name, model)
		}
		if tc.getState(model) != tc.wantState {
			t.Fatalf("%s: expected discipline state %s, got %#v", tc.name, tc.wantState, model)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("%s: expected point 10 to remain not complete, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0ExactIdentityBindingBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*DeploymentMultiTenantVal0Foundation)
		getState  func(DeploymentMultiTenantVal0Foundation) string
		wantState string
	}{
		{
			name: "wrong tenant boundary tenant scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantBoundary.TenantScope = "tenant:beta"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.TenantBoundaryState },
			wantState: DeploymentMultiTenantVal0TenantBoundaryStateBlocked,
		},
		{
			name: "whitespace padded tenant boundary tenant scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantBoundary.TenantScope = " " + deploymentMultiTenantVal0TenantScope() + " "
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.TenantBoundaryState },
			wantState: DeploymentMultiTenantVal0TenantBoundaryStateBlocked,
		},
		{
			name: "bare lower case tenant boundary tenant scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantBoundary.TenantScope = "alpha"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.TenantBoundaryState },
			wantState: DeploymentMultiTenantVal0TenantBoundaryStateBlocked,
		},
		{
			name: "wrong msp role scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.MSPAuthority.RoleScope = "support_readiness_operator_beta"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.MSPAuthorityState },
			wantState: DeploymentMultiTenantVal0MSPAuthorityStateBlocked,
		},
		{
			name: "wrong connector id blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.ConnectorContract.ConnectorID = "marketplace_audit_connector_beta"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.ConnectorContractState },
			wantState: DeploymentMultiTenantVal0ConnectorContractStateBlocked,
		},
		{
			name: "wrong operator actor blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.OperatorAction.Actor = "support_operator_999"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.OperatorActionState },
			wantState: DeploymentMultiTenantVal0OperatorActionStateBlocked,
		},
		{
			name: "wrong operator tenant target blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.OperatorAction.TenantTarget = "tenant:beta"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.OperatorActionState },
			wantState: DeploymentMultiTenantVal0OperatorActionStateBlocked,
		},
		{
			name: "prefixed operator tenant target blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.OperatorAction.TenantTarget = "ops " + deploymentMultiTenantVal0TenantScope()
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.OperatorActionState },
			wantState: DeploymentMultiTenantVal0OperatorActionStateBlocked,
		},
		{
			name: "wrong privacy scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.PrivacyGuard.TenantPrivacyScope = "tenant_privacy_scope_other"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.PrivacyGuardState },
			wantState: DeploymentMultiTenantVal0PrivacyGuardStateBlocked,
		},
		{
			name: "wrong fair share scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.FairShare.TenantResourceScope = "tenant_resource_scope_other"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.FairShareState },
			wantState: DeploymentMultiTenantVal0FairShareStateBlocked,
		},
		{
			name: "wrong operational preflight change scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.OperationalPreflight.TenantChangeScope = "tenant_change_scope_other"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.OperationalPreflightState },
			wantState: DeploymentMultiTenantVal0OperationalPreflightStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("%s: expected blocked current state, got %#v", tc.name, model)
		}
		if tc.getState(model) != tc.wantState {
			t.Fatalf("%s: expected discipline state %s, got %#v", tc.name, tc.wantState, model)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("%s: expected point 10 to remain not complete, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantVal0WhitespaceRetaggedFreshnessAndEnumBlockers(t *testing.T) {
	testCases := []struct {
		name      string
		mutate    func(*DeploymentMultiTenantVal0Foundation)
		getState  func(DeploymentMultiTenantVal0Foundation) string
		wantState string
	}{
		{
			name: "whitespace padded deployment freshness blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.DeploymentValidation.ValidationFreshnessState = " " + IntelligenceCalibrationFreshnessFresh + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.DeploymentValidationState },
			wantState: DeploymentMultiTenantVal0DeploymentValidationStateBlocked,
		},
		{
			name: "whitespace padded tenant boundary freshness blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantBoundary.BoundaryFreshnessState = " " + IntelligenceCalibrationFreshnessFresh + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.TenantBoundaryState },
			wantState: DeploymentMultiTenantVal0TenantBoundaryStateBlocked,
		},
		{
			name: "whitespace padded authority freshness blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.MSPAuthority.AuthorityFreshnessState = " " + IntelligenceCalibrationFreshnessFresh + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.MSPAuthorityState },
			wantState: DeploymentMultiTenantVal0MSPAuthorityStateBlocked,
		},
		{
			name: "tab padded msp tenant scope blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.MSPAuthority.TenantScope = "\t" + deploymentMultiTenantVal0TenantScope()
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.MSPAuthorityState },
			wantState: DeploymentMultiTenantVal0MSPAuthorityStateBlocked,
		},
		{
			name: "whitespace padded authority mode blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.MSPAuthority.AuthorityMode = " " + DeploymentMultiTenantAuthorityModeBounded + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.MSPAuthorityState },
			wantState: DeploymentMultiTenantVal0MSPAuthorityStateBlocked,
		},
		{
			name: "whitespace padded policy freshness blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.PolicyEnvelope.PolicyFreshnessState = " " + IntelligenceCalibrationFreshnessFresh + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.PolicyEnvelopeState },
			wantState: DeploymentMultiTenantVal0PolicyEnvelopeStateBlocked,
		},
		{
			name: "whitespace padded trust freshness blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.TenantTrustScope.TrustFreshnessState = " " + IntelligenceCalibrationFreshnessFresh + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.TenantTrustScopeState },
			wantState: DeploymentMultiTenantVal0TenantTrustScopeStateBlocked,
		},
		{
			name: "whitespace padded connector freshness blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.ConnectorContract.ConnectorFreshnessState = " " + IntelligenceCalibrationFreshnessFresh + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.ConnectorContractState },
			wantState: DeploymentMultiTenantVal0ConnectorContractStateBlocked,
		},
		{
			name: "whitespace padded connector failure behavior blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.ConnectorContract.FailureBehavior = " " + DeploymentMultiTenantConnectorFailureFailClosed + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.ConnectorContractState },
			wantState: DeploymentMultiTenantVal0ConnectorContractStateBlocked,
		},
		{
			name: "whitespace padded connector replay behavior blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.ConnectorContract.ReplayBehavior = " " + DeploymentMultiTenantConnectorReplayIdempotent + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.ConnectorContractState },
			wantState: DeploymentMultiTenantVal0ConnectorContractStateBlocked,
		},
		{
			name: "whitespace padded connector recovery behavior blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.ConnectorContract.RecoveryBehavior = " " + DeploymentMultiTenantConnectorRecoveryDeterminism + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.ConnectorContractState },
			wantState: DeploymentMultiTenantVal0ConnectorContractStateBlocked,
		},
		{
			name: "whitespace padded operator freshness blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.OperatorAction.ActionFreshnessState = " " + IntelligenceCalibrationFreshnessFresh + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.OperatorActionState },
			wantState: DeploymentMultiTenantVal0OperatorActionStateBlocked,
		},
		{
			name: "whitespace padded privacy freshness blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.PrivacyGuard.PrivacyFreshnessState = " " + IntelligenceCalibrationFreshnessFresh + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.PrivacyGuardState },
			wantState: DeploymentMultiTenantVal0PrivacyGuardStateBlocked,
		},
		{
			name: "whitespace padded fleet visibility mode blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.PrivacyGuard.FleetVisibilityMode = " " + DeploymentMultiTenantFleetVisibilityAggregated + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.PrivacyGuardState },
			wantState: DeploymentMultiTenantVal0PrivacyGuardStateBlocked,
		},
		{
			name: "whitespace padded support visibility mode blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.PrivacyGuard.SupportVisibilityMode = " " + DeploymentMultiTenantSupportVisibilityExplicit + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.PrivacyGuardState },
			wantState: DeploymentMultiTenantVal0PrivacyGuardStateBlocked,
		},
		{
			name: "whitespace padded fair share freshness blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.FairShare.FairShareFreshnessState = " " + IntelligenceCalibrationFreshnessFresh + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.FairShareState },
			wantState: DeploymentMultiTenantVal0FairShareStateBlocked,
		},
		{
			name: "whitespace padded operational preflight freshness blocks",
			mutate: func(model *DeploymentMultiTenantVal0Foundation) {
				model.OperationalPreflight.PreflightFreshnessState = " " + IntelligenceCalibrationFreshnessFresh + "\t"
			},
			getState:  func(model DeploymentMultiTenantVal0Foundation) string { return model.OperationalPreflightState },
			wantState: DeploymentMultiTenantVal0OperationalPreflightStateBlocked,
		},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("%s: expected blocked current state, got %#v", tc.name, model)
		}
		if tc.getState(model) != tc.wantState {
			t.Fatalf("%s: expected discipline state %s, got %#v", tc.name, tc.wantState, model)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("%s: expected point 10 to remain not complete, got %#v", tc.name, model)
		}
		assertDeploymentMultiTenantVal0NoPoint10Pass(t, model)
	}
}

func TestDeploymentMultiTenantVal0AggregateStateRejectsWhitespaceRetaggedDisciplineState(t *testing.T) {
	model := activeDeploymentMultiTenantVal0Model()
	model.DeploymentValidationState = " " + DeploymentMultiTenantVal0DeploymentValidationStateActive + "\t"
	if got := EvaluateDeploymentMultiTenantVal0State(model); got != DeploymentMultiTenantVal0StateBlocked {
		t.Fatalf("expected aggregate Val0 state to reject whitespace-retagged discipline state, got %q", got)
	}
}

func TestDeploymentMultiTenantVal0NoOverclaimAdversarialBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		claims []string
	}{
		{name: "direct forbidden phrase blocks exact state", claims: []string{"production approved"}},
		{name: "unicode confusable forbidden phrase blocks exact state", claims: []string{"prοductiοn apprοved"}},
		{name: "open-o confusable forbidden phrase blocks exact state", claims: []string{"prɔduction apprɔved"}},
		{name: "small capital confusable forbidden phrase blocks exact state", claims: []string{"prᴏduction approved"}},
		{name: "small capital deployment approved blocks exact state", claims: []string{"ᴅeployment approved"}},
		{name: "small capital certified managed trust blocks exact state", claims: []string{"cᴇrtified managed trust"}},
		{name: "ligature forbidden phrase blocks exact state", claims: []string{"marketplace certiﬁed"}},
		{name: "split field forbidden phrase blocks exact state", claims: []string{"production", "approved"}},
		{name: "open-o split field forbidden phrase blocks exact state", claims: []string{"prɔduction", "apprɔved"}},
		{name: "single bucket forbidden phrase with innocuous middle token blocks exact state", claims: []string{"production audit note approved"}},
		{name: "open-o innocuous middle token forbidden phrase blocks exact state", claims: []string{"prɔduction audit nɔte apprɔved"}},
		{name: "three bucket forbidden phrase with innocuous middle token blocks exact state", claims: []string{"production", "audit note", "approved"}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantVal0Model()
		model.NoOverclaim.ObservedClaims = append([]string{}, tc.claims...)
		model = ComputeDeploymentMultiTenantVal0Foundation(model)
		if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
			t.Fatalf("%s: expected blocked current state, got %#v", tc.name, model)
		}
		if model.NoOverclaimState != DeploymentMultiTenantVal0NoOverclaimStateBlocked {
			t.Fatalf("%s: expected blocked no-overclaim state, got %#v", tc.name, model)
		}
		if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
			t.Fatalf("%s: expected point 10 to remain not complete, got %#v", tc.name, model)
		}
		payload, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("%s: marshal model: %v", tc.name, err)
		}
		if strings.Contains(string(payload), "point_"+"10_pass") {
			t.Fatalf("%s: expected blocked no-overclaim path not to emit point 10 pass, got %s", tc.name, string(payload))
		}
	}
}

func TestDeploymentMultiTenantVal0FutureContractAdversarialObservedClaimsBlockers(t *testing.T) {
	model := activeDeploymentMultiTenantVal0Model()
	model.FutureContract.ObservedClaims = []string{"certifiɛd", "managed trust"}
	model = ComputeDeploymentMultiTenantVal0Foundation(model)
	if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
		t.Fatalf("expected adversarial future contract claims to block current state, got %#v", model)
	}
	if model.FutureContractState != DeploymentMultiTenantVal0FutureContractStateBlocked {
		t.Fatalf("expected adversarial future contract claims to block future contract state, got %#v", model)
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		t.Fatalf("expected point 10 to remain not complete after adversarial future contract claims, got %#v", model)
	}

	model = activeDeploymentMultiTenantVal0Model()
	model.FutureContract.ObservedClaims = []string{"production", "audit note", "approved"}
	model = ComputeDeploymentMultiTenantVal0Foundation(model)
	if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
		t.Fatalf("expected split future contract claims to block current state, got %#v", model)
	}
	if model.FutureContractState != DeploymentMultiTenantVal0FutureContractStateBlocked {
		t.Fatalf("expected split future contract claims to block future contract state, got %#v", model)
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		t.Fatalf("expected point 10 to remain not complete after split future contract claims, got %#v", model)
	}

	model = activeDeploymentMultiTenantVal0Model()
	model.FutureContract.ObservedClaims = []string{"ᴅeployment approved"}
	model = ComputeDeploymentMultiTenantVal0Foundation(model)
	if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
		t.Fatalf("expected small-capital future contract deployment approval claim to block current state, got %#v", model)
	}
	if model.FutureContractState != DeploymentMultiTenantVal0FutureContractStateBlocked {
		t.Fatalf("expected small-capital future contract deployment approval claim to block future contract state, got %#v", model)
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		t.Fatalf("expected point 10 to remain not complete after small-capital future contract deployment approval claim, got %#v", model)
	}

	model = activeDeploymentMultiTenantVal0Model()
	model.FutureContract.ObservedClaims = []string{"marketplace certiﬁed"}
	model = ComputeDeploymentMultiTenantVal0Foundation(model)
	if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
		t.Fatalf("expected ligature future contract claim to block current state, got %#v", model)
	}
	if model.FutureContractState != DeploymentMultiTenantVal0FutureContractStateBlocked {
		t.Fatalf("expected ligature future contract claim to block future contract state, got %#v", model)
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		t.Fatalf("expected point 10 to remain not complete after ligature future contract claim, got %#v", model)
	}

	model = activeDeploymentMultiTenantVal0Model()
	model.FutureContract.ObservedClaims = []string{"production audit note approved"}
	model = ComputeDeploymentMultiTenantVal0Foundation(model)
	if model.CurrentState != DeploymentMultiTenantVal0StateBlocked {
		t.Fatalf("expected single-bucket future contract claim with innocuous middle token to block current state, got %#v", model)
	}
	if model.FutureContractState != DeploymentMultiTenantVal0FutureContractStateBlocked {
		t.Fatalf("expected single-bucket future contract claim with innocuous middle token to block future contract state, got %#v", model)
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		t.Fatalf("expected point 10 to remain not complete after single-bucket future contract claim with innocuous middle token, got %#v", model)
	}
}

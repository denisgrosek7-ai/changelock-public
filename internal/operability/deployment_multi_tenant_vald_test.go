package operability

import (
	"encoding/json"
	"strings"
	"testing"
)

func activeDeploymentMultiTenantValDModel() DeploymentMultiTenantValDFoundation {
	model := DeploymentMultiTenantValDFoundationModel()
	return ComputeDeploymentMultiTenantValDFoundation(model)
}

func deploymentMultiTenantValDHasFinding(findings []DeploymentMultiTenantValDClosureBlockerFinding, level, surface, reason string) bool {
	for _, finding := range findings {
		if finding.BlockerLevel == level &&
			finding.Surface == surface &&
			strings.Contains(finding.Reason, reason) {
			return true
		}
	}
	return false
}

func TestDeploymentMultiTenantValDHappyPathAndPoint10NotComplete(t *testing.T) {
	model := activeDeploymentMultiTenantValDModel()
	if model.CurrentState != DeploymentMultiTenantValDStateActive {
		t.Fatalf("expected active Val D state, got %#v", model)
	}
	if model.AgenticOverlay.LearningLoopState != DeploymentMultiTenantValDAgentLearningLoopStateActive {
		t.Fatalf("expected active learning loop state, got %#v", model)
	}
	if model.AgenticOverlay.LearningLoop.LearningMode != "offline_sandbox_only" ||
		!model.AgenticOverlay.LearningLoop.TrainingDataPrivacyFiltered ||
		!model.AgenticOverlay.LearningLoop.TrainingDataTenantScoped ||
		!model.AgenticOverlay.LearningLoop.HumanFeedbackAuditLinked ||
		!model.AgenticOverlay.LearningLoop.LearnedOutputAdvisoryOnly ||
		model.AgenticOverlay.LearningLoop.ProductionSelfModificationAllowed ||
		model.AgenticOverlay.LearningLoop.ProductionMutationAllowed ||
		model.AgenticOverlay.LearningLoop.CanonicalMutationAllowed ||
		model.AgenticOverlay.LearningLoop.Point10PassAllowed {
		t.Fatalf("expected bounded advisory learning loop happy path, got %#v", model.AgenticOverlay.LearningLoop)
	}
	if model.ClosureBlockerState != DeploymentMultiTenantValDClosureBlockerStateActive {
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
		t.Fatalf("expected Val D to never emit point 10 pass, got %s", string(payload))
	}
}

func TestDeploymentMultiTenantValDAggregateProjectionDisclaimerBlocks(t *testing.T) {
	model := activeDeploymentMultiTenantValDModel()
	model.ProjectionDisclaimer = "canonical_truth"
	model = ComputeDeploymentMultiTenantValDFoundation(model)
	if model.CurrentState != DeploymentMultiTenantValDStateBlocked {
		t.Fatalf("expected malformed aggregate projection disclaimer to block ValD state, got %#v", model)
	}
	if !containsTrimmedString(model.BlockingReasons, "aggregate_projection_disclaimer_blocked") {
		t.Fatalf("expected aggregate projection disclaimer blocking reason, got %#v", model.BlockingReasons)
	}
}

func TestDeploymentMultiTenantValDDependencyBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValDFoundation)
	}{
		{name: "valc current state partial blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCCurrentState = "partial"
		}},
		{name: "valc dependency state blocked blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCDependencyState = DeploymentMultiTenantValCDependencyStateBlocked
		}},
		{name: "valc ha readiness state blocked blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCHAReadinessState = DeploymentMultiTenantValCHAReadinessStateBlocked
		}},
		{name: "valc recovery readiness state blocked blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCRecoveryReadinessState = DeploymentMultiTenantValCRecoveryReadinessStateBlocked
		}},
		{name: "valc sla readiness state blocked blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCSLAReadinessState = DeploymentMultiTenantValCSLAReadinessStateBlocked
		}},
		{name: "valc tenant trust scope state blocked blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCTenantTrustScopeState = DeploymentMultiTenantValCTenantTrustScopeStateBlocked
		}},
		{name: "valc silo visibility state blocked blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCSiloVisibilityState = DeploymentMultiTenantValCSiloVisibilityStateBlocked
		}},
		{name: "valc privacy guard state blocked blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCPrivacyGuardState = DeploymentMultiTenantValCPrivacyGuardStateBlocked
		}},
		{name: "valc no overclaim state blocked blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCNoOverclaimState = DeploymentMultiTenantValCNoOverclaimStateBlocked
		}},
		{name: "valc closure blocker state blocked blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCClosureBlockerState = DeploymentMultiTenantValCClosureBlockerStateBlocked
		}},
		{name: "valc closure blocker state cleanup blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCClosureBlockerState = DeploymentMultiTenantValCClosureBlockerStateCleanup
		}},
		{name: "valc closure blocker state advisory blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCClosureBlockerState = DeploymentMultiTenantValCClosureBlockerStateAdvisory
		}},
		{name: "point10 state complete blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.Point10State = "deployment_multi_tenant_point_10_complete"
		}},
		{name: "malformed projection disclaimer blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ProjectionDisclaimer = "canonical_truth"
		}},
	}

	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.DependencyState != DeploymentMultiTenantValDDependencyStateBlocked || model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%s: expected blocked dependency state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValDDependencySnapshotCopiesComputedValCProjectionDisclaimer(t *testing.T) {
	valC := ComputeDeploymentMultiTenantValCFoundation(DeploymentMultiTenantValCFoundationModel())
	valC.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_valc aggregate_dependency_snapshot"
	valC.ClosureBlockerOverlay.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_valc component_closure_blocker"
	snapshot := deploymentMultiTenantValDDependencySnapshotFromValC(valC)
	if snapshot.ProjectionDisclaimer != valC.ProjectionDisclaimer {
		t.Fatalf("expected dependency snapshot disclaimer to match aggregate computed Val C output, got snapshot=%q valc=%q", snapshot.ProjectionDisclaimer, valC.ProjectionDisclaimer)
	}
	if snapshot.ProjectionDisclaimer == valC.ClosureBlockerOverlay.ProjectionDisclaimer {
		t.Fatalf("expected dependency snapshot not to fallback to component disclaimer, got snapshot=%q component=%q", snapshot.ProjectionDisclaimer, valC.ClosureBlockerOverlay.ProjectionDisclaimer)
	}
	if EvaluateDeploymentMultiTenantValDDependencyState(snapshot) != DeploymentMultiTenantValDDependencyStateActive {
		t.Fatalf("expected copied computed disclaimer to keep dependency gate active, got %#v", snapshot)
	}
}

func TestDeploymentMultiTenantValDDependencyProjectionDisclaimerRegression(t *testing.T) {
	testCases := []string{
		"",
		"canonical_truth",
		"production_approval",
		"unknown",
		"blocked",
	}
	for _, disclaimer := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		model.Dependency.ProjectionDisclaimer = disclaimer
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.DependencyState != DeploymentMultiTenantValDDependencyStateBlocked {
			t.Fatalf("%q: expected blocked dependency state, got %#v", disclaimer, model)
		}
		if model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%q: expected blocked final Val D state, got %#v", disclaimer, model)
		}
	}
}

func TestDeploymentMultiTenantValDDependencySnapshotPropagatesUpstreamProjectionDisclaimer(t *testing.T) {
	valC := ComputeDeploymentMultiTenantValCFoundation(DeploymentMultiTenantValCFoundationModel())
	valC.ProjectionDisclaimer = "blocked"
	valC.ClosureBlockerOverlay.ProjectionDisclaimer = "projection_only not_canonical_truth deployment_multi_tenant_valc component_closure_blocker"
	snapshot := deploymentMultiTenantValDDependencySnapshotFromValC(valC)
	if snapshot.ProjectionDisclaimer != "blocked" {
		t.Fatalf("expected snapshot to propagate altered upstream disclaimer, got %#v", snapshot)
	}
	if EvaluateDeploymentMultiTenantValDDependencyState(snapshot) != DeploymentMultiTenantValDDependencyStateBlocked {
		t.Fatalf("expected altered upstream disclaimer to block dependency gate, got %#v", snapshot)
	}
}

func TestDeploymentMultiTenantValDConnectorCapabilityBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValDFoundation)
	}{
		{name: "connector id missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.ConnectorID = ""
		}},
		{name: "tenant scope missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.TenantScope = ""
		}},
		{name: "tenant scope global blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.TenantScope = "global_connector_scope"
		}},
		{name: "tenant scope unscoped blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.TenantScope = "unscoped_connector_scope"
		}},
		{name: "tenant scope cross tenant blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.TenantScope = "cross_tenant_connector_scope"
		}},
		{name: "permission manifest missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.PermissionManifest = ""
		}},
		{name: "capability manifest missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.CapabilityManifestPresent = false
		}},
		{name: "write capability undeclared blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.WriteCapabilities = nil
		}},
		{name: "mutation without explicit capability blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.MutationAllowed = true
			model.ConnectorCapability.MutationCapabilityExplicit = false
		}},
		{name: "mutation without reason blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.MutationAllowed = true
			model.ConnectorCapability.MutationCapabilityExplicit = true
			model.ConnectorCapability.Reason = ""
		}},
		{name: "mutation without audit blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.MutationAllowed = true
			model.ConnectorCapability.MutationCapabilityExplicit = true
			model.ConnectorCapability.AuditID = ""
		}},
		{name: "connector as source of truth blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.ConnectorAsSourceOfTruth = true
		}},
		{name: "connector bypasses tenant evidence deployment data residency gates blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.ConnectorBypassesTenantGate = true
		}},
		{name: "retry replay duplicate active evidence risk blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.RetryReplayDuplicatesActiveEvidenceRisk = true
		}},
		{name: "missing replay policy blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.ReplayPolicy = ""
		}},
		{name: "missing rate limit policy blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.RateLimitPolicy = ""
		}},
		{name: "invalid evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.EvidenceRefs = []string{"revoked_connector_evidence"}
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.ConnectorCapabilityState != DeploymentMultiTenantValDConnectorCapabilityStateBlocked || model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%s: expected blocked connector capability state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValDOperatorActionBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValDFoundation)
	}{
		{name: "actor missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.Actor = ""
		}},
		{name: "tenant target missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.TenantTarget = ""
		}},
		{name: "action scope global blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.ActionScope = "global_operator_scope"
		}},
		{name: "action scope unscoped blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.ActionScope = "unscoped_operator_scope"
		}},
		{name: "action scope cross tenant blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.ActionScope = "cross_tenant_operator_scope"
		}},
		{name: "reason missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.Reason = ""
		}},
		{name: "authorization basis missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.AuthorizationBasis = ""
		}},
		{name: "authority basis missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.AuthorityBasis = ""
		}},
		{name: "approval missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.Approver = ""
		}},
		{name: "expiry missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.Expiry = ""
		}},
		{name: "revocation path missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.RevocationPath = ""
		}},
		{name: "audit id missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.AuditID = ""
		}},
		{name: "rbac abac bypass blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.RBACABACEnforced = false
		}},
		{name: "sso context bypass blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.SSOContextBound = false
		}},
		{name: "support scope unbounded blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.SupportScopeBound = false
		}},
		{name: "production mutation allowed blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.ProductionMutationAllowed = true
		}},
		{name: "canonical mutation allowed blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.CanonicalMutationAllowed = true
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.OperatorActionState != DeploymentMultiTenantValDOperatorActionStateBlocked || model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%s: expected blocked operator action state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValDSupportAccessBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValDFoundation)
	}{
		{name: "support actor missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.SupportActor = ""
		}},
		{name: "tenant target missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.TenantTarget = ""
		}},
		{name: "support scope global blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.SupportScope = "global_support_scope"
		}},
		{name: "support scope unscoped blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.SupportScope = "unscoped_support_scope"
		}},
		{name: "sso session reference missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.SSOSessionReference = ""
		}},
		{name: "rbac role missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.RBACRole = ""
		}},
		{name: "abac conditions missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.ABACConditions = ""
		}},
		{name: "authority basis missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.AuthorityBasis = ""
		}},
		{name: "reason missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.Reason = ""
		}},
		{name: "approval missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.Approver = ""
		}},
		{name: "expiry missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.Expiry = ""
		}},
		{name: "revocation path missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.RevocationPath = ""
		}},
		{name: "audit id missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.AuditID = ""
		}},
		{name: "data residency boundary bypass blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.DataResidencyBoundaryRespected = false
		}},
		{name: "tenant isolation boundary bypass blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.TenantIsolationBoundaryRespected = false
		}},
		{name: "support visibility exceeds scope blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.SupportVisibilityExceedsScope = true
		}},
		{name: "raw tenant evidence exposed blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.RawTenantEvidenceExposed = true
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.SupportAccessState != DeploymentMultiTenantValDSupportAccessStateBlocked || model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%s: expected blocked support access state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValDBreakGlassBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValDFoundation)
	}{
		{name: "emergency reason missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.EmergencyReason = ""
		}},
		{name: "actor missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.Actor = ""
		}},
		{name: "tenant target missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.TenantTarget = ""
		}},
		{name: "action scope global blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.ActionScope = "global_break_glass_scope"
		}},
		{name: "action scope unscoped blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.ActionScope = "unscoped_break_glass_scope"
		}},
		{name: "authorization basis missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.AuthorizationBasis = ""
		}},
		{name: "approval missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.Approver = ""
		}},
		{name: "expiry missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.Expiry = ""
		}},
		{name: "expiry expired blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.Expiry = "expired_break_glass_expiry"
		}},
		{name: "revocation path missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.RevocationPath = ""
		}},
		{name: "audit id missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.AuditID = ""
		}},
		{name: "post action review missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.PostActionReviewRequired = false
		}},
		{name: "persistent access blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.PersistentAccessGranted = true
		}},
		{name: "data residency boundary bypass blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.DataResidencyBoundaryRespected = false
		}},
		{name: "tenant isolation boundary bypass blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.TenantIsolationBoundaryRespected = false
		}},
		{name: "break glass creates pass authority blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.CreatesPASSAuthority = true
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.BreakGlassState != DeploymentMultiTenantValDBreakGlassStateBlocked || model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%s: expected blocked break-glass state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValDMarketplaceMSPBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValDFoundation)
	}{
		{name: "marketplace profile missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.MarketplaceProfile = ""
		}},
		{name: "marketplace profile ambiguous blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.MarketplaceProfile = "marketplace_profile_ish"
		}},
		{name: "msp scope global blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.MSPOperatorScope = "global_msp_scope"
		}},
		{name: "msp scope unscoped blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.MSPOperatorScope = "unscoped_msp_scope"
		}},
		{name: "msp scope cross tenant blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.MSPOperatorScope = "cross_tenant_msp_scope"
		}},
		{name: "partner scope global blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.PartnerScope = "global_partner_scope"
		}},
		{name: "partner scope unscoped blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.PartnerScope = "unscoped_partner_scope"
		}},
		{name: "customer ready evidence missing while wording present blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.CustomerReadyWordingPresent = true
			model.MarketplaceMSPAuthority.CustomerReadyValidationEvidence = ""
		}},
		{name: "msp partner pass authority allowed blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.PassAuthorityAllowed = true
		}},
		{name: "msp partner production readiness authority allowed blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.ProductionReadinessAuthorityAllowed = true
		}},
		{name: "msp partner source of truth allowed blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.SourceOfTruthAllowed = true
		}},
		{name: "marketplace install treated as production approval blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.MarketplaceInstallTreatedAsProdApproved = true
		}},
		{name: "audit reason expiry revocation missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.AuditID = ""
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.MarketplaceMSPAuthorityState != DeploymentMultiTenantValDMarketplaceMSPAuthorityStateBlocked || model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%s: expected blocked marketplace msp authority state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValDAgenticOverlayBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValDFoundation)
	}{
		{name: "missing permission manifest blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.PermissionManifest = ""
		}},
		{name: "missing tenant scope blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.TenantScope = ""
		}},
		{name: "global scope blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.TenantScope = "global_agent_scope"
		}},
		{name: "unscoped scope blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.TenantScope = "unscoped_agent_scope"
		}},
		{name: "cross tenant scope blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.TenantScope = "cross_tenant_agent_scope"
		}},
		{name: "external api enabled by default blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.ExternalAPIAllowed = true
		}},
		{name: "missing approval blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.Approver = ""
		}},
		{name: "missing audit trail blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.AuditID = ""
		}},
		{name: "canonical mutation allowed blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.CanonicalMutationAllowed = true
		}},
		{name: "production mutation allowed blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.ProductionMutationAllowed = true
		}},
		{name: "point10 pass allowed blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.Point10PassAllowed = true
		}},
		{name: "agent treated as source of truth blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.AgentTreatedAsSourceOfTruth = true
		}},
		{name: "containment executed without approval blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.TenantBoundaryContainmentAgent.ContainmentExecutedWithoutApproval = true
		}},
		{name: "agent changes tenant state directly blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.TenantBoundaryContainmentAgent.ChangesTenantStateDirectly = true
		}},
		{name: "deployment agent treats install success as readiness blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.DeploymentHealthPreflightAgent.InstallSuccessTreatedAsReadiness = true
		}},
		{name: "preflight executed without approval blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.DeploymentHealthPreflightAgent.PreflightExecutedWithoutApproval = true
		}},
		{name: "agent marks deployment ready without canonical evaluator blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.DeploymentHealthPreflightAgent.MarksDeploymentReadyWithoutCanonicalEvaluator = true
		}},
		{name: "connector mutation executed by agent blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.ConnectorOperatorMisuseWatchAgent.ConnectorMutationExecutedByAgent = true
		}},
		{name: "connector capability missing blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.ConnectorOperatorMisuseWatchAgent.ConnectorCapabilityMissing = true
		}},
		{name: "operator support action without authority basis blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.ConnectorOperatorMisuseWatchAgent.OperatorSupportActionWithoutAuthorityBasis = true
		}},
		{name: "break glass without expiry revocation blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.ConnectorOperatorMisuseWatchAgent.BreakGlassExpiryRevocationMissing = true
		}},
		{name: "recovery execution without approval blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RecoveryRebuildRecommendationAgent.RestoreRollbackRebuildExecutedAutomatically = true
		}},
		{name: "recovery recommendation without backup restore dr evidence blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RecoveryRebuildRecommendationAgent.RecoveryEvidencePack = ""
		}},
		{name: "recovery recommendation bypasses tenant isolation blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RecoveryRebuildRecommendationAgent.RecommendationBypassesTenantIsolation = true
		}},
		{name: "recovery recommendation bypasses data residency blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RecoveryRebuildRecommendationAgent.RecommendationBypassesDataResidency = true
		}},
		{name: "recovery guaranteed wording blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RecoveryRebuildRecommendationAgent.RecoveryGuaranteedClaim = true
		}},
		{name: "stale evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.EvidenceRefs = []string{"stale_agent_evidence"}
		}},
		{name: "revoked evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.EvidenceRefs = []string{"revoked_agent_evidence"}
		}},
		{name: "expired evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.EvidenceRefs = []string{"expired_agent_evidence"}
		}},
		{name: "duplicate evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.EvidenceRefs = []string{"duplicate_agent_evidence"}
		}},
		{name: "unrelated evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.EvidenceRefs = []string{"unrelated_agent_evidence"}
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.AgenticOverlayState != DeploymentMultiTenantValDAgenticOverlayStateBlocked || model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%s: expected blocked agentic overlay state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValDAgentLearningLoopBlockers(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*DeploymentMultiTenantValDFoundation)
	}{
		{name: "online production self modification blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.LearningMode = "online_self_modification"
		}},
		{name: "self promotion blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.AgentSelfPromotes = true
		}},
		{name: "self deploy blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.AgentSelfDeploys = true
		}},
		{name: "runtime activation without human approval blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.RuntimeActivationAllowed = true
			model.AgenticOverlay.LearningLoop.RuntimeActivationApprovalStatus = ""
		}},
		{name: "model promotion without human approval blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.PromotionAllowed = true
			model.AgenticOverlay.LearningLoop.PromotionApprovalStatus = ""
		}},
		{name: "cross tenant training data blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.TrainingDataCrossTenant = true
		}},
		{name: "unapproved customer-data training blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.TrainingDataCustomerApproved = false
		}},
		{name: "missing audit-linked human feedback blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.HumanFeedbackAuditLinked = false
		}},
		{name: "missing regression tests blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.RegressionTestRefs = nil
		}},
		{name: "missing no-overclaim check blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.NoOverclaimCheckRefs = nil
		}},
		{name: "missing tenant-scope check blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.TenantScopeCheckRefs = nil
		}},
		{name: "missing approval-gate check blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.ApprovalGateCheckRefs = nil
		}},
		{name: "candidate weakening no-overclaim blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.CandidateWeakensNoOverclaim = true
		}},
		{name: "candidate weakening tenant isolation blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.CandidateWeakensTenantIsolation = true
		}},
		{name: "candidate weakening approval gates blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.CandidateWeakensApprovalGates = true
		}},
		{name: "candidate enabling external api by default blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.CandidateEnablesExternalAPIByDefault = true
		}},
		{name: "candidate enabling production mutation blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.CandidateEnablesProductionMutation = true
		}},
		{name: "candidate enabling canonical evidence mutation blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.CandidateEnablesCanonicalMutation = true
		}},
		{name: "candidate enabling point10 pass blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.CandidateEnablesPoint10Pass = true
		}},
		{name: "recommendation approval treated as execution approval blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.RecommendationApprovalMeansExecutionApproval = true
		}},
		{name: "execution approval treated as model upgrade approval blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.ExecutionApprovalMeansModelUpgradeApproval = true
		}},
		{name: "learned output treated as canonical truth blocks", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.LearnedOutputTreatedAsCanonicalTruth = true
		}},
		{name: "stale evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.EvidenceRefs = []string{"stale_learning_evidence"}
		}},
		{name: "revoked evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.EvidenceRefs = []string{"revoked_learning_evidence"}
		}},
		{name: "expired evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.EvidenceRefs = []string{"expired_learning_evidence"}
		}},
		{name: "duplicate evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.EvidenceRefs = []string{"duplicate_learning_evidence"}
		}},
		{name: "unrelated evidence refs block", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.EvidenceRefs = []string{"unrelated_learning_evidence"}
		}},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.AgenticOverlay.LearningLoopState != DeploymentMultiTenantValDAgentLearningLoopStateBlocked ||
			model.AgenticOverlayState != DeploymentMultiTenantValDAgenticOverlayStateBlocked ||
			model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%s: expected blocked learning loop and final state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValDClosureBlockerOverlayCLB0AndCLB1Blockers(t *testing.T) {
	testCases := []struct {
		name    string
		mutate  func(*DeploymentMultiTenantValDFoundation)
		level   string
		surface string
		reason  string
	}{
		{name: "connector mutation without explicit capability produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.MutationAllowed = true
			model.ConnectorCapability.MutationCapabilityExplicit = false
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, reason: "connector mutation without explicit capability"},
		{name: "connector as source of truth produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.ConnectorAsSourceOfTruth = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, reason: "connector treated as source of truth"},
		{name: "operator without authority basis produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.AuthorityBasis = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceOperatorAction, reason: "operator or support action without authority basis"},
		{name: "sso rbac bypass produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.RBACABACEnforced = false
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceSupportAccess, reason: "sso or rbac abac bypass"},
		{name: "break glass without authority basis produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.AuthorizationBasis = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceBreakGlass, reason: "break-glass without authority basis"},
		{name: "msp overclaim produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.MarketplaceMSPAuthority.PassAuthorityAllowed = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceMarketplaceMSP, reason: "marketplace or msp overclaim"},
		{name: "agent production mutation produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.ProductionMutationAllowed = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "agent executes production mutation"},
		{name: "agent canonical mutation produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.CanonicalMutationAllowed = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "agent mutates canonical evidence spine"},
		{name: "agent point10 pass allowed produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.Point10PassAllowed = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "agent emits or enables point 10 pass"},
		{name: "agent cross tenant access produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.CrossTenantAccess = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "agent performs cross-tenant access"},
		{name: "agent external api by default produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.ExternalAPIAllowed = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "agent enables external api by default"},
		{name: "learning loop self promotion produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.AgentSelfPromotes = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "agent self-promotes"},
		{name: "learning loop self deploy produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.AgentSelfDeploys = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "agent self-deploys"},
		{name: "learning loop runtime activation without approval produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.RuntimeActivationAllowed = true
			model.AgenticOverlay.LearningLoop.RuntimeActivationApprovalStatus = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "runtime activation without human approval"},
		{name: "learning loop model promotion without approval produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.PromotionAllowed = true
			model.AgenticOverlay.LearningLoop.PromotionApprovalStatus = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "model promotion without human approval"},
		{name: "learning loop cross tenant training data produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.TrainingDataCrossTenant = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "cross-tenant training data used"},
		{name: "learning loop unapproved customer data produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.TrainingDataCustomerApproved = false
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "unapproved customer data used for training"},
		{name: "learning loop recommendation approval means execution approval produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.RecommendationApprovalMeansExecutionApproval = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "recommendation approval treated as execution approval"},
		{name: "learning loop execution approval means model upgrade approval produces cl b0 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.ExecutionApprovalMeansModelUpgradeApproval = true
		}, level: DeploymentMultiTenantValDBlockerLevelCLB0, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "execution approval treated as model upgrade approval"},
		{name: "approval workflow missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.ApprovalQueue = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "approval workflow missing"},
		{name: "audit trail missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.AuditID = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "audit trail missing"},
		{name: "evidence refs missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.EvidenceRefs = nil
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "evidence refs missing"},
		{name: "permission manifest missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.PermissionManifest = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, reason: "permission manifest missing"},
		{name: "connector capability manifest missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.CapabilityManifestPresent = false
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, reason: "connector capability manifest missing"},
		{name: "break glass expiry revocation missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.BreakGlass.RevocationPath = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceBreakGlass, reason: "break-glass expiry or revocation path missing"},
		{name: "support access expiry revocation missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.SupportAccess.RevocationPath = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceSupportAccess, reason: "support access expiry or revocation missing"},
		{name: "recovery recommendation lacks required evidence pack produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RecoveryRebuildRecommendationAgent.RecoveryEvidencePack = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "recovery recommendation lacks required evidence pack"},
		{name: "learning workflow missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.TrainingApprovalStatus = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "training approval workflow missing"},
		{name: "human feedback audit link missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.HumanFeedbackAuditLinked = false
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "human feedback audit link missing"},
		{name: "evaluation result refs missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.EvaluationResultRefs = nil
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "evaluation result refs missing"},
		{name: "regression test refs missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.RegressionTestRefs = nil
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "regression test refs missing"},
		{name: "no-overclaim check refs missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.NoOverclaimCheckRefs = nil
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "no-overclaim check refs missing"},
		{name: "tenant-scope check refs missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.TenantScopeCheckRefs = nil
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "tenant-scope check refs missing"},
		{name: "approval-gate check refs missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.ApprovalGateCheckRefs = nil
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "approval-gate check refs missing"},
		{name: "model candidate versioning missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.ModelCandidateID = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "model candidate versioning missing"},
		{name: "baseline model version missing produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.BaselineModelVersion = ""
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "baseline model version missing"},
		{name: "stale revoked expired duplicate unrelated evidence handling not proven produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.ModelCandidateID = "expired_model_candidate_id"
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "stale revoked expired duplicate or unrelated evidence handling not proven"},
		{name: "dependency gate not exact active produces cl b1 blocker", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.Dependency.ValCCurrentState = "partial"
		}, level: DeploymentMultiTenantValDBlockerLevelCLB1, surface: DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, reason: "dependency gate missing or not exact active"},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.ClosureBlockerState != DeploymentMultiTenantValDClosureBlockerStateBlocked || model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%s: expected blocked closure state, got %#v", tc.name, model)
		}
		if !deploymentMultiTenantValDHasFinding(model.ClosureBlockerOverlay.Findings, tc.level, tc.surface, tc.reason) {
			t.Fatalf("%s: expected finding %s/%s containing %q, got %#v", tc.name, tc.level, tc.surface, tc.reason, model.ClosureBlockerOverlay)
		}
	}
}

func TestDeploymentMultiTenantValDClosureBlockerOverlayCLB2Cleanup(t *testing.T) {
	testCases := []struct {
		name    string
		mutate  func(*DeploymentMultiTenantValDFoundation)
		surface string
		reason  string
	}{
		{name: "ambiguous connector naming emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.ConnectorCapability.ConnectorNamingExact = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, reason: "ambiguous connector naming"},
		{name: "ambiguous operator action naming emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.OperatorAction.OperatorActionNamingExact = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceOperatorAction, reason: "ambiguous operator action naming"},
		{name: "ambiguous agent naming emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.AgentNamingExact = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "ambiguous agent naming"},
		{name: "ambiguous learning mode naming emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.LearningModeNamingExact = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "ambiguous learning mode naming"},
		{name: "ambiguous model candidate naming emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.ModelCandidateNamingExact = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "ambiguous model candidate naming"},
		{name: "missing safe wording example emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.SafeWordingExamplePresent = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "missing safe wording example"},
		{name: "missing safe wording example for advisory learning emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.SafeWordingExamplePresent = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "missing safe wording example for advisory learning"},
		{name: "incomplete diagnostic output emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.DiagnosticOutputComplete = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "incomplete diagnostic output"},
		{name: "incomplete diagnostic output for learning loop emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.DiagnosticOutputComplete = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "incomplete diagnostic output for learning-loop blockers"},
		{name: "incomplete runbook wording emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.RuntimeApprovalController.RunbookWordingComplete = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "incomplete bounded runbook wording"},
		{name: "incomplete runbook wording for model promotion emits cleanup", mutate: func(model *DeploymentMultiTenantValDFoundation) {
			model.AgenticOverlay.LearningLoop.RunbookWordingComplete = false
		}, surface: DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, reason: "incomplete runbook wording for model promotion"},
	}
	for _, tc := range testCases {
		model := activeDeploymentMultiTenantValDModel()
		tc.mutate(&model)
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.ClosureBlockerState != DeploymentMultiTenantValDClosureBlockerStateCleanup || model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("%s: expected cleanup closure blocker state and blocked final state, got %#v", tc.name, model)
		}
		if !deploymentMultiTenantValDHasFinding(model.ClosureBlockerOverlay.Findings, DeploymentMultiTenantValDBlockerLevelCLB2, tc.surface, tc.reason) {
			t.Fatalf("%s: expected CL-B2 finding, got %#v", tc.name, model.ClosureBlockerOverlay)
		}
	}
}

func TestDeploymentMultiTenantValDClosureBlockerOverlayCLB3Advisory(t *testing.T) {
	model := activeDeploymentMultiTenantValDModel()
	model.ClosureBlockerOverlay = DeploymentMultiTenantValDClosureBlockerOverlay{
		ProjectionDisclaimer: deploymentMultiTenantValDProjectionDisclaimer(),
		Findings: []DeploymentMultiTenantValDClosureBlockerFinding{
			{
				BlockerLevel:      DeploymentMultiTenantValDBlockerLevelCLB3,
				Surface:           DeploymentMultiTenantValDClosureSurfaceAgenticOverlay,
				Reason:            "advisory cleanup carried forward",
				BlocksCurrentWave: false,
				RequiredFollowup:  "record advisory cleanup if carried forward",
			},
		},
	}
	model.ClosureBlockerState = EvaluateDeploymentMultiTenantValDClosureBlockerState(model.ClosureBlockerOverlay)
	model.CurrentState = EvaluateDeploymentMultiTenantValDState(model)
	if model.ClosureBlockerState != DeploymentMultiTenantValDClosureBlockerStateAdvisory {
		t.Fatalf("expected advisory closure blocker state, got %#v", model)
	}
	if model.CurrentState != DeploymentMultiTenantValDStateBlocked {
		t.Fatalf("expected advisory closure state to keep final Val D blocked, got %#v", model)
	}
}

func TestDeploymentMultiTenantValDClosureBlockerOverlayRejectsLegacyAndUnknownLevels(t *testing.T) {
	testCases := []struct {
		name    string
		finding DeploymentMultiTenantValDClosureBlockerFinding
	}{
		{
			name: "legacy priority zero is rejected",
			finding: DeploymentMultiTenantValDClosureBlockerFinding{
				BlockerLevel:      deploymentMultiTenantLegacyPriority("0"),
				Surface:           DeploymentMultiTenantValDClosureSurfaceAgenticOverlay,
				Reason:            "legacy severity rejected",
				BlocksCurrentWave: true,
			},
		},
		{
			name: "legacy priority one is rejected",
			finding: DeploymentMultiTenantValDClosureBlockerFinding{
				BlockerLevel:      deploymentMultiTenantLegacyPriority("1"),
				Surface:           DeploymentMultiTenantValDClosureSurfaceAgenticOverlay,
				Reason:            "legacy severity rejected",
				BlocksCurrentWave: true,
				RequiredFollowup:  "use cl b blocker terminology",
			},
		},
		{
			name: "legacy priority two is rejected",
			finding: DeploymentMultiTenantValDClosureBlockerFinding{
				BlockerLevel:      deploymentMultiTenantLegacyPriority("2"),
				Surface:           DeploymentMultiTenantValDClosureSurfaceAgenticOverlay,
				Reason:            "legacy severity rejected",
				BlocksCurrentWave: true,
				RequiredFollowup:  "use cl b blocker terminology",
			},
		},
		{
			name: "unknown blocker level is rejected",
			finding: DeploymentMultiTenantValDClosureBlockerFinding{
				BlockerLevel:      "CL-B9",
				Surface:           DeploymentMultiTenantValDClosureSurfaceAgenticOverlay,
				Reason:            "unknown blocker level",
				BlocksCurrentWave: true,
			},
		},
		{
			name: "unknown surface is rejected",
			finding: DeploymentMultiTenantValDClosureBlockerFinding{
				BlockerLevel:      DeploymentMultiTenantValDBlockerLevelCLB0,
				Surface:           "unknown_surface",
				Reason:            "unknown surface",
				BlocksCurrentWave: true,
			},
		},
	}
	for _, tc := range testCases {
		model := DeploymentMultiTenantValDClosureBlockerOverlay{
			ProjectionDisclaimer: deploymentMultiTenantValDProjectionDisclaimer(),
			Findings:             []DeploymentMultiTenantValDClosureBlockerFinding{tc.finding},
		}
		if EvaluateDeploymentMultiTenantValDClosureBlockerState(model) != DeploymentMultiTenantValDClosureBlockerStateBlocked {
			t.Fatalf("%s: expected blocked state, got %#v", tc.name, model)
		}
	}
}

func TestDeploymentMultiTenantValDNoOverclaimBlockers(t *testing.T) {
	blockedClaims := []string{
		"connector approved deployment",
		"connector is source of truth",
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
		"marketplace production ready",
		"marketplace certified",
		"customer ready without validation",
		"autonomous remediation approved",
		"agent approved deployment",
		"agent certified recovery",
		"AI certified fix",
		"auto-merge safe",
		"auto-deploy safe",
		"production autopatch",
		"recovery guaranteed",
		"agent guaranteed tenant isolation",
		"agent proves compliance",
		"point_" + "10_pass by agent",
		"agent is source of truth",
		"external AI verified",
		"external API verified",
		"no human approval required",
		"production mutation approved by agent",
		"canonical evidence mutation by agent",
		"clean-room certified",
		"patent cleared",
		"FTO cleared",
		"legal certification",
		"copied competitor workflow",
		"self-learning agent approved deployment",
		"agent self-approved model",
		"autonomous model promotion",
		"autonomous runtime activation",
		"agent trained itself into production",
		"online self-modification safe",
		"model upgrade automatically approved",
		"recommendation approval means execution approval",
		"execution approval means model upgrade approval",
		"learned output is canonical truth",
		"learned model certified",
		"AI model certified",
		"agent learning guarantees security",
		"self-improving agent guarantees compliance",
		"point_" + "10_pass by learned model",
		"external AI verified model",
		"sandboxed connector execution connector is source of truth",
		"connector capability manifest connector mutation safe by default",
		"bounded operator authority operator approved deployment",
		"break-glass approval required break-glass permanent access",
		"marketplace deployment profile marketplace production ready",
		"advisory recommendation agent approved deployment",
		"human-approved action required no human approval required",
		"approval-gated recovery recommendation recovery guaranteed",
		"tenant-scoped agent runtime point_" + "10_pass by agent",
	}
	allowedClaims := []string{
		"sandboxed connector execution",
		"connector capability manifest",
		"explicit connector capability",
		"connector misuse signal",
		"operator misuse signal",
		"bounded operator authority",
		"zero-trust operator action",
		"tenant-scoped support access",
		"break-glass approval required",
		"break-glass expiry and revocation evidence",
		"MSP support surface",
		"marketplace deployment profile",
		"advisory recommendation",
		"human-approved action required",
		"evidence-linked recommendation",
		"tenant-scoped agent runtime",
		"approval-gated recovery recommendation",
		"offline sandbox learning pipeline",
		"human-approved model promotion",
		"candidate model version",
		"advisory learning improvement",
		"audit-linked human feedback",
		"regression-tested candidate",
		"no-overclaim checked candidate",
		"tenant-scope checked candidate",
		"approval-gated runtime activation",
		"learned output remains advisory",
		"no production autopatch",
		"no auto-merge",
		"no auto-deploy",
		"not canonical truth",
		"not production approval",
		"not deployment approval",
		"not compliance certification",
	}

	for _, claim := range blockedClaims {
		model := activeDeploymentMultiTenantValDModel()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValDNoOverclaimStateBlocked || model.CurrentState != DeploymentMultiTenantValDStateBlocked {
			t.Fatalf("expected blocked no-overclaim state for %q, got %#v", claim, model)
		}
	}

	for _, claim := range allowedClaims {
		model := activeDeploymentMultiTenantValDModel()
		model.NoOverclaim.ObservedClaims = []string{claim}
		model = ComputeDeploymentMultiTenantValDFoundation(model)
		if model.NoOverclaimState != DeploymentMultiTenantValDNoOverclaimStateActive {
			t.Fatalf("expected allowed claim %q to remain active, got %#v", claim, model)
		}
	}
}

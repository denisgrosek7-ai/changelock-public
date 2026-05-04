package formal

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/operability"
)

func point11Val0ActiveDependencySnapshot() Point11Val0DependencySnapshot {
	valE := operability.ComputeDeploymentMultiTenantValEFoundation(operability.DeploymentMultiTenantValEFoundationModel())
	return SnapshotPoint11Val0DependencyFromComputedPoint10ValE(valE, Point11Val0Point10RepoReview{
		LatestValEClosurePatchPresent: true,
		Point10PassOutsideValE:        false,
		CIGreenVisible:                true,
		CIGreen:                       true,
		MergeStatusVisible:            true,
		MergeAccepted:                 true,
	})
}

func activePoint11Val0Foundation() Point11Val0Foundation {
	model := Point11Val0FoundationModel()
	model.Dependency = point11Val0ActiveDependencySnapshot()
	return ComputePoint11Val0Foundation(model)
}

func point11Val0PassToken() string {
	return "point_" + "11_pass"
}

func TestPoint11Val0FoundationHappyPathActive(t *testing.T) {
	model := activePoint11Val0Foundation()
	if model.CurrentState != Point11Val0StateActive {
		t.Fatalf("expected active foundation state, got %#v", model)
	}
	if model.DependencyState != Point11Val0DependencyStateActive {
		t.Fatalf("expected active dependency state, got %#v", model)
	}
	body, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal foundation: %v", err)
	}
	if strings.Contains(string(body), point11Val0PassToken()) {
		t.Fatalf("expected no point 11 final pass token in val 0 output, got %s", body)
	}
}

func TestPoint11Val0FoundationDefaultsToDependencyReviewRequiredWhenRepoVisibilityMissing(t *testing.T) {
	model := ComputePoint11Val0Foundation(Point11Val0FoundationModel())
	if model.DependencyState != Point11Val0DependencyStateReviewRequired {
		t.Fatalf("expected dependency review required when ci or merge visibility is missing, got %#v", model)
	}
	if model.CurrentState != Point11Val0StateReviewRequired {
		t.Fatalf("expected review required aggregate state, got %#v", model)
	}
	if !point11Val0ContainsTrimmed(model.ReviewPrerequisites, "point10_ci_green_not_visible_in_repo_context") ||
		!point11Val0ContainsTrimmed(model.ReviewPrerequisites, "point10_merge_state_not_visible_in_repo_context") {
		t.Fatalf("expected review prerequisites for missing ci or merge visibility, got %#v", model)
	}
}

func TestPoint11Val0DependencyGate(t *testing.T) {
	t.Run("happy path point10 dependency active", func(t *testing.T) {
		snapshot := point11Val0ActiveDependencySnapshot()
		if snapshot.Point10State != operability.DeploymentMultiTenantPoint10StatePass {
			t.Fatalf("expected copied point10 pass from val e, got %#v", snapshot)
		}
		if got := EvaluatePoint11Val0DependencyState(snapshot); got != Point11Val0DependencyStateActive {
			t.Fatalf("expected active dependency state, got %#v", snapshot)
		}
	})

	t.Run("copied projection disclaimer propagates exactly from computed upstream output", func(t *testing.T) {
		valE := operability.ComputeDeploymentMultiTenantValEFoundation(operability.DeploymentMultiTenantValEFoundationModel())
		valE.Point10PassRule.ProjectionDisclaimer = "projection_only not_canonical_truth propagated_upstream_point10_vale"
		snapshot := SnapshotPoint11Val0DependencyFromComputedPoint10ValE(valE, Point11Val0Point10RepoReview{
			LatestValEClosurePatchPresent: true,
			CIGreenVisible:                true,
			CIGreen:                       true,
			MergeStatusVisible:            true,
			MergeAccepted:                 true,
		})
		if snapshot.ProjectionDisclaimer != valE.Point10PassRule.ProjectionDisclaimer {
			t.Fatalf("expected exact copied projection disclaimer, got snapshot=%q valE=%q", snapshot.ProjectionDisclaimer, valE.Point10PassRule.ProjectionDisclaimer)
		}
		if got := EvaluatePoint11Val0DependencyState(snapshot); got != Point11Val0DependencyStateActive {
			t.Fatalf("expected active dependency with propagated disclaimer, got %#v", snapshot)
		}
	})

	testCases := []struct {
		name      string
		mutate    func(*Point11Val0DependencySnapshot)
		wantState string
	}{
		{name: "malformed upstream projection disclaimer blocks", mutate: func(model *Point11Val0DependencySnapshot) {
			model.ProjectionDisclaimer = "canonical_truth"
		}, wantState: Point11Val0DependencyStateBlocked},
		{name: "missing point10 pass blocks", mutate: func(model *Point11Val0DependencySnapshot) {
			model.Point10State = operability.DeploymentMultiTenantPoint10StateNotComplete
		}, wantState: Point11Val0DependencyStateBlocked},
		{name: "point10 pass outside vale blocks", mutate: func(model *Point11Val0DependencySnapshot) {
			model.Point10PassOutsideValE = true
		}, wantState: Point11Val0DependencyStateBlocked},
		{name: "blocked vale pass closure manifest blocks", mutate: func(model *Point11Val0DependencySnapshot) {
			model.Point10PassClosureManifestState = operability.DeploymentMultiTenantValEPassClosureManifestStateBlocked
		}, wantState: Point11Val0DependencyStateBlocked},
		{name: "blocked vale no overclaim state blocks", mutate: func(model *Point11Val0DependencySnapshot) {
			model.Point10NoOverclaimState = operability.DeploymentMultiTenantValENoOverclaimStateBlocked
		}, wantState: Point11Val0DependencyStateBlocked},
		{name: "blocked vale clean room ip state blocks", mutate: func(model *Point11Val0DependencySnapshot) {
			model.Point10CleanRoomIPState = operability.DeploymentMultiTenantValECleanRoomIPStateBlocked
		}, wantState: Point11Val0DependencyStateBlocked},
		{name: "blocked vale clb closure state blocks", mutate: func(model *Point11Val0DependencySnapshot) {
			model.Point10CLBClosureState = operability.DeploymentMultiTenantValECLBClosureStateBlocked
		}, wantState: Point11Val0DependencyStateBlocked},
		{name: "blocked vale evidence quality state blocks", mutate: func(model *Point11Val0DependencySnapshot) {
			model.Point10EvidenceQualityState = operability.DeploymentMultiTenantValEEvidenceQualityStateBlocked
		}, wantState: Point11Val0DependencyStateBlocked},
		{name: "blocked vale projection boundary state blocks", mutate: func(model *Point11Val0DependencySnapshot) {
			model.Point10ProjectionBoundaryState = operability.DeploymentMultiTenantValEProjectionBoundaryStateBlocked
		}, wantState: Point11Val0DependencyStateBlocked},
		{name: "unknown ci merge status is review prerequisite", mutate: func(model *Point11Val0DependencySnapshot) {
			model.CIGreenVisible = false
			model.MergeStatusVisible = false
			model.ReviewPrerequisites = []string{
				"point10_ci_green_not_visible_in_repo_context",
				"point10_merge_state_not_visible_in_repo_context",
			}
		}, wantState: Point11Val0DependencyStateReviewRequired},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11Val0Foundation()
			testCase.mutate(&model.Dependency)
			model = ComputePoint11Val0Foundation(model)
			if model.DependencyState != testCase.wantState {
				t.Fatalf("expected dependency state %q, got %#v", testCase.wantState, model)
			}
			switch testCase.wantState {
			case Point11Val0DependencyStateActive:
				if model.CurrentState != Point11Val0StateActive {
					t.Fatalf("expected active aggregate state, got %#v", model)
				}
			case Point11Val0DependencyStateReviewRequired:
				if model.CurrentState != Point11Val0StateReviewRequired {
					t.Fatalf("expected review required aggregate state, got %#v", model)
				}
			default:
				if model.CurrentState != Point11Val0StateBlocked {
					t.Fatalf("expected blocked aggregate state, got %#v", model)
				}
			}
		})
	}
}

func TestPoint11Val0PolicyContractState(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point11Val0PolicyContract)
	}{
		{name: "missing policy id blocks", mutate: func(model *Point11Val0PolicyContract) { model.PolicyID = "" }},
		{name: "missing version blocks", mutate: func(model *Point11Val0PolicyContract) { model.Version = "" }},
		{name: "unsigned policy blocks", mutate: func(model *Point11Val0PolicyContract) { model.SignedState = "unsigned_policy_contract" }},
		{name: "unanchored policy blocks", mutate: func(model *Point11Val0PolicyContract) { model.AnchoredState = "unanchored_policy_contract" }},
		{name: "revoked policy blocks", mutate: func(model *Point11Val0PolicyContract) { model.RevokedBy = "revoked_by_governance" }},
		{name: "expired policy blocks", mutate: func(model *Point11Val0PolicyContract) { model.EffectiveUntil = "2000-01-01T00:00:00Z" }},
		{name: "superseded policy without compatibility context blocks", mutate: func(model *Point11Val0PolicyContract) {
			model.SupersededBy = "policy_v2"
			model.CompatibilityVersion = ""
		}},
		{name: "malformed superseded by blocks policy contract", mutate: func(model *Point11Val0PolicyContract) {
			model.SupersededBy = "revoked/invalid marker"
			model.CompatibilityVersion = "compat_v1"
		}},
		{name: "global scope blocks", mutate: func(model *Point11Val0PolicyContract) { model.Scope = "global_admin_scope" }},
		{name: "policy without approval evidence refs blocks", mutate: func(model *Point11Val0PolicyContract) { model.ApprovalEvidenceRefs = nil }},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11Val0Foundation()
			testCase.mutate(&model.PolicyContract)
			model = ComputePoint11Val0Foundation(model)
			if model.PolicyContractState != Point11Val0PolicyContractStateBlocked {
				t.Fatalf("expected blocked policy contract state, got %#v", model)
			}
			if model.CurrentState != Point11Val0StateBlocked {
				t.Fatalf("expected blocked aggregate state, got %#v", model)
			}
		})
	}

	for _, supersededBy := range []string{
		"revoked/invalid marker",
		"unknown",
		"revoked",
		"invalid",
		"placeholder",
		"global",
		"all-tenants",
		"   ",
		"policy_unknown",
		"policy_revoked",
		"policy_invalid",
		"policy_expired",
		"policy_superseded",
		"policy_malformed",
		"point11_policy_unknown",
		"point11_policy_revoked",
	} {
		t.Run("superseded by "+supersededBy+" blocks lineage", func(t *testing.T) {
			model := activePoint11Val0Foundation()
			model.PolicyContract.SupersededBy = supersededBy
			model.PolicyContract.CompatibilityVersion = "point11_val0_compat_v1"
			model = ComputePoint11Val0Foundation(model)
			if model.PolicyContractState != Point11Val0PolicyContractStateBlocked {
				t.Fatalf("expected blocked policy contract for superseded_by=%q, got %#v", supersededBy, model)
			}
		})
	}

	t.Run("canonical superseded by passes with valid compatibility version", func(t *testing.T) {
		model := activePoint11Val0Foundation()
		model.PolicyContract.SupersededBy = "policy_successor_2026_05_02"
		model.PolicyContract.CompatibilityVersion = "point11_val0_compat_v1"
		model = ComputePoint11Val0Foundation(model)
		if model.PolicyContractState != Point11Val0PolicyContractStateActive {
			t.Fatalf("expected active policy contract with canonical superseded_by, got %#v", model)
		}
		if model.CurrentState != Point11Val0StateActive {
			t.Fatalf("expected active aggregate state, got %#v", model)
		}
	})
}

func TestPoint11Val0ClaimGovernanceState(t *testing.T) {
	t.Run("valid draft claim remains draft not published", func(t *testing.T) {
		model := activePoint11Val0Foundation()
		model.ClaimGovernance.LifecycleState = Point11Val0ClaimLifecycleDraft
		model.ClaimGovernance.PublicationBoundary = point11Val0PublicationSurfaceAgentOutput
		model = ComputePoint11Val0Foundation(model)
		if model.ClaimGovernanceState != Point11Val0ClaimGovernanceStateActive {
			t.Fatalf("expected active claim governance state, got %#v", model)
		}
		if model.ClaimGovernance.LifecycleState != Point11Val0ClaimLifecycleDraft {
			t.Fatalf("expected lifecycle to remain draft, got %#v", model.ClaimGovernance)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11Val0ClaimGovernance)
	}{
		{name: "published claim requires policy version owner scope and revocation path", mutate: func(model *Point11Val0ClaimGovernance) {
			model.LifecycleState = Point11Val0ClaimLifecyclePublished
			model.PolicyVersion = ""
		}},
		{name: "blocked claim cannot appear in docs output", mutate: func(model *Point11Val0ClaimGovernance) {
			model.ClaimCategory = Point11Val0ClaimCategoryBlocked
			model.PublicationBoundary = point11Val0PublicationSurfaceDocs
		}},
		{name: "internal only claim cannot become customer visible", mutate: func(model *Point11Val0ClaimGovernance) {
			model.ClaimCategory = Point11Val0ClaimCategoryInternalOnly
			model.PublicationBoundary = point11Val0PublicationSurfacePortal
		}},
		{name: "review required claim without approval blocks publication", mutate: func(model *Point11Val0ClaimGovernance) {
			model.ClaimCategory = Point11Val0ClaimCategoryReviewRequired
			model.PublicationBoundary = point11Val0PublicationSurfaceDocs
			model.ApprovalStatus = ""
		}},
		{name: "expired claim blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.Expiry = time.Now().UTC().Add(-time.Hour).Format(time.RFC3339)
		}},
		{name: "revoked claim lifecycle blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.LifecycleState = Point11Val0ClaimLifecycleRevoked
		}},
		{name: "superseded claim lifecycle blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.LifecycleState = Point11Val0ClaimLifecycleSuperseded
		}},
		{name: "revoked status blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.RevocationOrSupersessionStatus = "revoked"
		}},
		{name: "superseded status blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.RevocationOrSupersessionStatus = "superseded"
		}},
		{name: "expired status blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.RevocationOrSupersessionStatus = "expired"
		}},
		{name: "docs without clean room review blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.PublicationBoundary = point11Val0PublicationSurfaceDocs
			model.CleanRoomIPReview = ""
		}},
		{name: "portal without clean room review blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.PublicationBoundary = point11Val0PublicationSurfacePortal
			model.CleanRoomIPReview = ""
		}},
		{name: "partner material without clean room review blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.PublicationBoundary = point11Val0PublicationSurfacePartner
			model.CleanRoomIPReview = ""
		}},
		{name: "buyer claim without clean room review blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.PublicationBoundary = point11Val0PublicationSurfaceBuyer
			model.CleanRoomIPReview = ""
		}},
		{name: "sales material without clean room review blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.PublicationBoundary = point11Val0PublicationSurfaceSales
			model.CleanRoomIPReview = ""
		}},
		{name: "demo material without clean room review blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.PublicationBoundary = point11Val0PublicationSurfaceDemo
			model.CleanRoomIPReview = ""
		}},
		{name: "export without clean room review blocks", mutate: func(model *Point11Val0ClaimGovernance) {
			model.PublicationBoundary = point11Val0PublicationSurfaceExport
			model.CleanRoomIPReview = ""
		}},
		{name: "agent output cannot become public safe claim without governance event", mutate: func(model *Point11Val0ClaimGovernance) {
			model.PublicationBoundary = point11Val0PublicationSurfaceAgentOutput
			model.ClaimCategory = Point11Val0ClaimCategoryPublicSafe
			model.GovernanceEvent = ""
		}},
		{name: "claim with missing evidence refs blocks export publication", mutate: func(model *Point11Val0ClaimGovernance) {
			model.PublicationBoundary = point11Val0PublicationSurfaceExport
			model.EvidenceRefs = nil
		}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11Val0Foundation()
			testCase.mutate(&model.ClaimGovernance)
			model = ComputePoint11Val0Foundation(model)
			if model.ClaimGovernanceState != Point11Val0ClaimGovernanceStateBlocked {
				t.Fatalf("expected blocked claim governance state, got %#v", model)
			}
			if model.CurrentState != Point11Val0StateBlocked {
				t.Fatalf("expected blocked aggregate state, got %#v", model)
			}
		})
	}

	t.Run("valid future expiry and active status passes", func(t *testing.T) {
		model := activePoint11Val0Foundation()
		model.ClaimGovernance.Expiry = time.Now().UTC().Add(24 * time.Hour).Format(time.RFC3339)
		model.ClaimGovernance.RevocationOrSupersessionStatus = "claim_active"
		model = ComputePoint11Val0Foundation(model)
		if model.ClaimGovernanceState != Point11Val0ClaimGovernanceStateActive {
			t.Fatalf("expected active claim governance state, got %#v", model)
		}
	})

	for _, surface := range []string{
		point11Val0PublicationSurfaceDocs,
		point11Val0PublicationSurfacePortal,
		point11Val0PublicationSurfaceExport,
		point11Val0PublicationSurfacePartner,
		point11Val0PublicationSurfaceDemo,
		point11Val0PublicationSurfaceSales,
		point11Val0PublicationSurfaceBuyer,
	} {
		t.Run("public facing surface "+surface+" passes with clean room review", func(t *testing.T) {
			model := activePoint11Val0Foundation()
			model.ClaimGovernance.PublicationBoundary = surface
			model.ClaimGovernance.CleanRoomIPReview = "clean_room_review_point11_val0"
			model = ComputePoint11Val0Foundation(model)
			if model.ClaimGovernanceState != Point11Val0ClaimGovernanceStateActive {
				t.Fatalf("expected active claim governance state for surface %q, got %#v", surface, model)
			}
		})
	}
}

func TestPoint11Val0AuthorityMatrixState(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point11Val0AuthorityMatrix)
	}{
		{name: "partner self approval blocks", mutate: func(model *Point11Val0AuthorityMatrix) {
			model.Proposer = "partner_operator"
			model.FinalApprover = "partner_operator"
		}},
		{name: "customer self approval blocks", mutate: func(model *Point11Val0AuthorityMatrix) {
			model.Proposer = "customer_admin"
			model.FinalApprover = "customer_admin"
		}},
		{name: "agent self approval blocks", mutate: func(model *Point11Val0AuthorityMatrix) {
			model.Proposer = "agent_runtime"
			model.FinalApprover = "agent_runtime"
		}},
		{name: "proposer as final approver for public claim blocks", mutate: func(model *Point11Val0AuthorityMatrix) {
			model.Proposer = "governance_reviewer"
			model.FinalApprover = "governance_reviewer"
			model.CustomerVisibleOrPublic = true
		}},
		{name: "policy relaxation without governance event blocks", mutate: func(model *Point11Val0AuthorityMatrix) {
			model.PolicyRelaxationRequested = true
			model.GovernanceEvent = ""
		}},
		{name: "authority expansion without governance event blocks", mutate: func(model *Point11Val0AuthorityMatrix) {
			model.AuthorityExpansionRequested = true
			model.GovernanceEvent = ""
		}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11Val0Foundation()
			testCase.mutate(&model.AuthorityMatrix)
			model = ComputePoint11Val0Foundation(model)
			if model.AuthorityMatrixState != Point11Val0AuthorityMatrixStateBlocked {
				t.Fatalf("expected blocked authority matrix state, got %#v", model)
			}
		})
	}
}

func TestPoint11Val0ExceptionGovernanceState(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point11Val0ExceptionGovernance)
	}{
		{name: "exception without expiry blocks", mutate: func(model *Point11Val0ExceptionGovernance) { model.ExpiresAt = "" }},
		{name: "exception without revocation path blocks", mutate: func(model *Point11Val0ExceptionGovernance) { model.RevocationPath = "" }},
		{name: "exception without emergency claim id blocks", mutate: func(model *Point11Val0ExceptionGovernance) { model.EmergencyClaimID = "" }},
		{name: "emergency claim without approver blocks", mutate: func(model *Point11Val0ExceptionGovernance) { model.Approver = "" }},
		{name: "exception without audit id blocks", mutate: func(model *Point11Val0ExceptionGovernance) { model.AuditID = "" }},
		{name: "exception without evidence refs blocks", mutate: func(model *Point11Val0ExceptionGovernance) { model.EvidenceRefs = nil }},
		{name: "expired exception blocks", mutate: func(model *Point11Val0ExceptionGovernance) { model.ExpiresAt = "2000-01-01T00:00:00Z" }},
		{name: "permanent silent exception blocks", mutate: func(model *Point11Val0ExceptionGovernance) { model.PermanentSilentException = true }},
		{name: "cross tenant exception scope blocks", mutate: func(model *Point11Val0ExceptionGovernance) { model.Scope = "cross-tenant_exception_scope" }},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11Val0Foundation()
			testCase.mutate(&model.ExceptionGovernance)
			model = ComputePoint11Val0Foundation(model)
			if model.ExceptionGovernanceState != Point11Val0ExceptionGovernanceStateBlocked {
				t.Fatalf("expected blocked exception governance state, got %#v", model)
			}
		})
	}

	for _, emergencyClaimID := range []string{
		"",
		"   ",
		"unknown",
		"revoked",
		"invalid",
		"placeholder",
		"revoked/invalid marker",
		"global",
		"all-tenants",
		"emergency_claim_unknown",
		"emergency_claim_revoked",
		"emergency_claim_invalid",
		"emergency_claim_expired",
		"emergency_claim_superseded",
		"point11_emergency_claim_unknown",
		"point11_emergency_claim_revoked",
	} {
		t.Run("emergency claim id "+strings.TrimSpace(emergencyClaimID)+" blocks when malformed", func(t *testing.T) {
			model := activePoint11Val0Foundation()
			model.ExceptionGovernance.EmergencyClaimID = emergencyClaimID
			model = ComputePoint11Val0Foundation(model)
			if model.ExceptionGovernanceState != Point11Val0ExceptionGovernanceStateBlocked {
				t.Fatalf("expected blocked exception governance state for emergency_claim_id=%q, got %#v", emergencyClaimID, model)
			}
		})
	}

	t.Run("canonical emergency claim id passes", func(t *testing.T) {
		model := activePoint11Val0Foundation()
		model.ExceptionGovernance.EmergencyClaimID = "point11_emergency_claim_2026_05_02"
		model = ComputePoint11Val0Foundation(model)
		if model.ExceptionGovernanceState != Point11Val0ExceptionGovernanceStateActive {
			t.Fatalf("expected active exception governance state for canonical emergency claim id, got %#v", model)
		}
	})
}

func TestPoint11Val0ABACGovernanceState(t *testing.T) {
	t.Run("unknown attribute cannot create active decision", func(t *testing.T) {
		model := activePoint11Val0Foundation()
		model.ABACGovernance.UnknownAttributes = []string{"unknown_geo_attribute"}
		model = ComputePoint11Val0Foundation(model)
		if model.ABACGovernanceState != Point11Val0ABACStateBlocked {
			t.Fatalf("expected blocked abac governance state, got %#v", model)
		}
	})

	t.Run("deny over allow precedence is visible", func(t *testing.T) {
		model := activePoint11Val0Foundation()
		model.ABACGovernance.DeniedAttributes = []string{"tenant_scope_denied"}
		model.ABACGovernance.Diagnostics = []string{"deny_over_allow_precedence_visible"}
		model = ComputePoint11Val0Foundation(model)
		if model.ABACGovernanceState != Point11Val0ABACStateBlocked {
			t.Fatalf("expected blocked abac governance state, got %#v", model)
		}
		if !point11Val0ContainsTrimmed(model.ABACGovernance.Diagnostics, "deny_over_allow_precedence_visible") {
			t.Fatalf("expected visible deny over allow diagnostic, got %#v", model.ABACGovernance)
		}
	})

	t.Run("explanation includes attributes policy refs and claim refs", func(t *testing.T) {
		model := activePoint11Val0Foundation()
		model = ComputePoint11Val0Foundation(model)
		if model.ABACGovernanceState != Point11Val0ABACStateActive {
			t.Fatalf("expected active abac governance state, got %#v", model)
		}
		if len(model.ABACGovernance.ExplanationAttributes) == 0 || len(model.ABACGovernance.ExplanationPolicyRefs) == 0 || len(model.ABACGovernance.ExplanationClaimRefs) == 0 {
			t.Fatalf("expected explanation outputs to be populated, got %#v", model.ABACGovernance)
		}
	})

	t.Run("exception interaction visible in diagnostics", func(t *testing.T) {
		model := activePoint11Val0Foundation()
		model.ABACGovernance.ExceptionState = "exception_active"
		model.ABACGovernance.ExceptionInteraction = ""
		model = ComputePoint11Val0Foundation(model)
		if model.ABACGovernanceState != Point11Val0ABACStateBlocked {
			t.Fatalf("expected blocked abac governance state when exception interaction is hidden, got %#v", model)
		}
	})
}

func TestPoint11Val0DecisionBindingState(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point11Val0DecisionBinding)
	}{
		{name: "policy to decision binding requires policy ref", mutate: func(model *Point11Val0DecisionBinding) { model.PolicyRef = "" }},
		{name: "policy to decision binding requires evidence refs", mutate: func(model *Point11Val0DecisionBinding) { model.EvidenceRefs = nil }},
		{name: "invalid policy blocks decision active state", mutate: func(model *Point11Val0DecisionBinding) { model.PolicyRefState = point11Val0DecisionRefStateUnknown }},
		{name: "invalid claim blocks decision active state", mutate: func(model *Point11Val0DecisionBinding) { model.ClaimRefState = point11Val0DecisionRefStateUnknown }},
		{name: "revoked policy or claim blocks active decision", mutate: func(model *Point11Val0DecisionBinding) { model.PolicyRefState = point11Val0DecisionRefStateRevoked }},
		{name: "decision output cannot claim legal regulatory certification authority", mutate: func(model *Point11Val0DecisionBinding) {
			model.EnforcementOutcome = "official authority"
		}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11Val0Foundation()
			testCase.mutate(&model.DecisionBinding)
			model = ComputePoint11Val0Foundation(model)
			if model.DecisionBindingState != Point11Val0DecisionBindingStateBlocked {
				t.Fatalf("expected blocked decision binding state, got %#v", model)
			}
		})
	}
}

func TestPoint11Val0NoOverclaimState(t *testing.T) {
	blockedClaims := []string{
		"certified",
		"regulator-approved",
		"compliance guaranteed",
		"production approved",
		"AI-approved",
		"AI legal proof",
		"autonomous remediation",
		"supreme authority",
		"supreme arbiter",
		"impossible to violate without detection",
	}
	for _, blockedClaim := range blockedClaims {
		t.Run(blockedClaim, func(t *testing.T) {
			model := activePoint11Val0Foundation()
			model.NoOverclaim.ObservedClaims = []string{blockedClaim}
			model = ComputePoint11Val0Foundation(model)
			if model.NoOverclaimState != Point11Val0NoOverclaimStateBlocked {
				t.Fatalf("expected blocked no overclaim state for %q, got %#v", blockedClaim, model)
			}
		})
	}

	model := activePoint11Val0Foundation()
	model.NoOverclaim.ObservedClaims = []string{
		"signed and versioned policy contract",
		"evidence-linked governance decision",
		"not regulator approval",
		"not production approval",
	}
	model = ComputePoint11Val0Foundation(model)
	if model.NoOverclaimState != Point11Val0NoOverclaimStateActive {
		t.Fatalf("expected active no overclaim state for bounded wording, got %#v", model)
	}
}

func TestPoint11Val0CrossDomainCompatibilityState(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point11Val0CrossDomainCompatibility)
		want   string
	}{
		{name: "cross domain claim without trust root ref blocks", mutate: func(model *Point11Val0CrossDomainCompatibility) { model.TrustRootRef = "" }, want: Point11Val0CrossDomainCompatibilityStateBlocked},
		{name: "cross domain claim with unknown issuer trust rule blocks", mutate: func(model *Point11Val0CrossDomainCompatibility) {
			model.IssuerTrustRule = point11Val0CrossDomainIssuerUnknown
		}, want: Point11Val0CrossDomainCompatibilityStateBlocked},
		{name: "cross domain claim with incompatible scope yields review required", mutate: func(model *Point11Val0CrossDomainCompatibility) {
			model.ScopeCompatibility = point11Val0CrossDomainScopeIncompatible
		}, want: Point11Val0CrossDomainCompatibilityStateReviewRequired},
		{name: "cross domain claim with incompatible freshness yields review required", mutate: func(model *Point11Val0CrossDomainCompatibility) {
			model.FreshnessCompatibility = point11Val0CrossDomainFreshnessIncompatible
		}, want: Point11Val0CrossDomainCompatibilityStateReviewRequired},
		{name: "revoked remote claim blocks", mutate: func(model *Point11Val0CrossDomainCompatibility) {
			model.RemoteClaimState = point11Val0DecisionRefStateRevoked
		}, want: Point11Val0CrossDomainCompatibilityStateBlocked},
		{name: "expired remote claim blocks", mutate: func(model *Point11Val0CrossDomainCompatibility) {
			model.RemoteClaimState = point11Val0DecisionRefStateExpired
		}, want: Point11Val0CrossDomainCompatibilityStateBlocked},
		{name: "remote claim cannot override local policy", mutate: func(model *Point11Val0CrossDomainCompatibility) { model.RemoteOverridesLocalPolicy = true }, want: Point11Val0CrossDomainCompatibilityStateBlocked},
		{name: "cross domain compatibility cannot create certification authority", mutate: func(model *Point11Val0CrossDomainCompatibility) { model.CreatesCertificationAuthority = true }, want: Point11Val0CrossDomainCompatibilityStateBlocked},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11Val0Foundation()
			testCase.mutate(&model.CrossDomainCompatibility)
			model = ComputePoint11Val0Foundation(model)
			if model.CrossDomainCompatibilityState != testCase.want {
				t.Fatalf("expected cross domain compatibility state %q, got %#v", testCase.want, model)
			}
			if testCase.want == Point11Val0CrossDomainCompatibilityStateReviewRequired {
				if model.CurrentState != Point11Val0StateReviewRequired {
					t.Fatalf("expected aggregate review required state, got %#v", model)
				}
				return
			}
			if model.CurrentState != Point11Val0StateBlocked {
				t.Fatalf("expected aggregate blocked state, got %#v", model)
			}
		})
	}
}

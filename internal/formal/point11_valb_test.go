package formal

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func point11ValBActiveDependencySnapshot() Point11ValBDependencySnapshot {
	valA := activePoint11ValAFoundation()
	return SnapshotPoint11ValBDependencyFromComputedValA(valA, Point11ValBValAReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
	})
}

func reviewRequiredPoint11ValBDependencySnapshot() Point11ValBDependencySnapshot {
	valA := activePoint11ValAFoundation()
	valA.CurrentState = Point11ValAStateReviewRequired
	valA.DependencyState = Point11ValADependencyStateReviewRequired
	valA.ReviewPrerequisites = []string{"vala_repo_visibility_review_prerequisite"}
	return SnapshotPoint11ValBDependencyFromComputedValA(valA, Point11ValBValAReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
	})
}

func activePoint11ValBFoundation() Point11ValBFoundation {
	model := Point11ValBFoundationModel()
	model.Dependency = point11ValBActiveDependencySnapshot()
	return ComputePoint11ValBFoundation(model)
}

func TestPoint11ValBDependencyState(t *testing.T) {
	t.Run("happy path vala dependency active", func(t *testing.T) {
		snapshot := point11ValBActiveDependencySnapshot()
		if got := EvaluatePoint11ValBDependencyState(snapshot); got != Point11ValBDependencyStateActive {
			t.Fatalf("expected active dependency state, got %#v", snapshot)
		}
	})

	t.Run("copied vala projection disclaimer propagates exactly", func(t *testing.T) {
		valA := activePoint11ValAFoundation()
		valA.ProjectionDisclaimer = "projection_only not_canonical_truth aggregate_vala_disclaimer"
		valA.Registry.ProjectionDisclaimer = "projection_only not_canonical_truth component_registry_disclaimer"
		snapshot := SnapshotPoint11ValBDependencyFromComputedValA(valA, Point11ValBValAReviewContext{
			LocalReviewAllowsDependencyReviewRequired: true,
		})
		if snapshot.ProjectionDisclaimer != valA.ProjectionDisclaimer {
			t.Fatalf("expected aggregate vala projection disclaimer, got snapshot=%q vala=%q", snapshot.ProjectionDisclaimer, valA.ProjectionDisclaimer)
		}
		if got := EvaluatePoint11ValBDependencyState(snapshot); got != Point11ValBDependencyStateActive {
			t.Fatalf("expected active dependency with propagated disclaimer, got %#v", snapshot)
		}
	})

	t.Run("malformed vala aggregate projection disclaimer blocks", func(t *testing.T) {
		valA := activePoint11ValAFoundation()
		valA.ProjectionDisclaimer = "canonical_truth"
		valA.Registry.ProjectionDisclaimer = "projection_only not_canonical_truth component_registry_disclaimer"
		snapshot := SnapshotPoint11ValBDependencyFromComputedValA(valA, Point11ValBValAReviewContext{
			LocalReviewAllowsDependencyReviewRequired: true,
		})
		if snapshot.ProjectionDisclaimer != valA.ProjectionDisclaimer {
			t.Fatalf("expected malformed aggregate disclaimer to propagate without fallback, got snapshot=%q vala=%q", snapshot.ProjectionDisclaimer, valA.ProjectionDisclaimer)
		}
		if got := EvaluatePoint11ValBDependencyState(snapshot); got != Point11ValBDependencyStateBlocked {
			t.Fatalf("expected malformed aggregate disclaimer to block dependency, got %#v", snapshot)
		}
	})

	testCases := []struct {
		name      string
		mutate    func(*Point11ValBDependencySnapshot)
		wantState string
	}{
		{name: "blocked vala registry blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValARegistryState = Point11ValARegistryStateBlocked
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "blocked vala signature blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValASignatureState = Point11ValASignatureStateBlocked
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "blocked vala anchor blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValAAnchorState = Point11ValAAnchorStateBlocked
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "blocked vala lifecycle transition blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValALifecycleTransitionState = Point11ValALifecycleTransitionStateBlocked
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "vala policy use not active blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValAPolicyUseState = Point11ValAPolicyUseStateHistoricalOnly
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "blocked vala graph blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValAGraphState = Point11ValAGraphStateBlocked
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "whitespace retagged vala current state blocks raw exact", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValACurrentState = " " + Point11ValAStateActive
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "tab newline retagged vala dependency state blocks raw exact", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValADependencyState = "\t" + Point11ValADependencyStateActive + "\n"
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "vala point11 pass emission marker blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValAPoint11PassEmitted = true
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "vala authority marker blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValACreatesAuthorityClaims = true
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "vala publication marker blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValACreatesPublicationSideEffects = true
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "vala signing marker blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValACreatesSigningSideEffects = true
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "vala anchoring marker blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValACreatesAnchoringSideEffects = true
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "vala external api marker blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValACreatesExternalAPISideEffects = true
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "vala production side effect marker blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.ValACreatesProductionSideEffects = true
		}, wantState: Point11ValBDependencyStateBlocked},
		{name: "unresolved taxonomy drift carried as review prerequisite", mutate: func(model *Point11ValBDependencySnapshot) {
			model.UnresolvedCrossPointTaxonomyDrift = true
			model.TaxonomyDriftReviewPrerequisite = true
			model.ReviewPrerequisites = append(model.ReviewPrerequisites, "cross_point_taxonomy_drift_review_prerequisite")
		}, wantState: Point11ValBDependencyStateReviewRequired},
		{name: "unresolved taxonomy drift without review prerequisite blocks", mutate: func(model *Point11ValBDependencySnapshot) {
			model.UnresolvedCrossPointTaxonomyDrift = true
			model.TaxonomyDriftReviewPrerequisite = false
		}, wantState: Point11ValBDependencyStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := point11ValBActiveDependencySnapshot()
			testCase.mutate(&model)
			if got := EvaluatePoint11ValBDependencyState(model); got != testCase.wantState {
				t.Fatalf("expected dependency state %q, got %#v", testCase.wantState, model)
			}
		})
	}

	t.Run("retagged vala inherited state records exact dependency reason", func(t *testing.T) {
		for _, mutate := range []func(*Point11ValBDependencySnapshot){
			func(model *Point11ValBDependencySnapshot) { model.ValACurrentState = " " + Point11ValAStateActive },
			func(model *Point11ValBDependencySnapshot) {
				model.ValADependencyState = "\t" + Point11ValADependencyStateActive + "\n"
			},
		} {
			model := point11ValBActiveDependencySnapshot()
			mutate(&model)
			state, reasons := point11ValBDependencyStateAndReasons(model)
			if state != Point11ValBDependencyStateBlocked {
				t.Fatalf("expected retagged vala inherited state to block, got %q for %#v", state, model)
			}
			if !point11Val0ContainsTrimmed(reasons, "vala_dependency_not_active") {
				t.Fatalf("expected exact vala dependency reason, got %#v", reasons)
			}
		}
	})

	t.Run("taxonomy drift review prerequisite cannot mask blocked vala registry", func(t *testing.T) {
		valA := activePoint11ValAFoundation()
		valA.Registry.RegistryID = ""
		valA = ComputePoint11ValAFoundation(valA)
		if valA.RegistryState != Point11ValARegistryStateBlocked {
			t.Fatalf("expected blocked computed vala registry state, got %#v", valA)
		}

		snapshot := SnapshotPoint11ValBDependencyFromComputedValA(valA, Point11ValBValAReviewContext{
			LocalReviewAllowsDependencyReviewRequired: true,
			UnresolvedCrossPointTaxonomyDrift:         true,
			TaxonomyDriftReviewPrerequisite:           true,
		})
		if got := EvaluatePoint11ValBDependencyState(snapshot); got != Point11ValBDependencyStateBlocked {
			t.Fatalf("expected blocked vala registry to override taxonomy review prerequisite, got %#v", snapshot)
		}
	})
}

func TestPoint11ValBClaimTypeState(t *testing.T) {
	t.Run("valid claim type definition active", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		if model.ClaimTypeState != Point11ValBClaimTypeStateActive {
			t.Fatalf("expected active claim type, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValBClaimTypeDefinition)
	}{
		{name: "missing claim type id blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.ClaimTypeID = ""
		}},
		{name: "claim type unknown blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.ClaimTypeID = "claim_type_unknown"
		}},
		{name: "claim type revoked blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.ClaimTypeID = "claim_type_revoked"
		}},
		{name: "claim type invalid blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.ClaimTypeID = "claim_type_invalid"
		}},
		{name: "missing allowed subject kinds blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.AllowedSubjectKinds = nil
		}},
		{name: "missing allowed issuer kinds blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.AllowedIssuerKinds = nil
		}},
		{name: "missing allowed audiences blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.AllowedAudiences = nil
		}},
		{name: "missing allowed publication surfaces blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.AllowedPublicationSurfaces = nil
		}},
		{name: "customer visible claim type without clean room ip requirement blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.CustomerVisibleAllowed = true
			model.CleanRoomIPRequired = false
		}},
		{name: "public safe claim type without governance event required blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.PublicSafeAllowed = true
			model.GovernanceEventRequired = false
		}},
		{name: "cross domain allowed without trust compatibility requirement blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.CrossDomainAllowed = true
			model.CrossDomainTrustCompatibilityRequired = false
		}},
		{name: "agent origin allowed cannot imply publish authority", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.AgentOriginAllowed = true
			model.AgentPublishAuthority = true
		}},
		{name: "agent origin allowed cannot imply approve authority", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.AgentOriginAllowed = true
			model.AgentApproveAuthority = true
		}},
		{name: "forbidden overclaim wording in claim type blocks", mutate: func(model *Point11ValBClaimTypeDefinition) {
			model.ClaimType = "certified governance claim"
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValBFoundation()
			testCase.mutate(&model.ClaimTypeDefinition)
			model = ComputePoint11ValBFoundation(model)
			if model.ClaimTypeState != Point11ValBClaimTypeStateBlocked {
				t.Fatalf("expected blocked claim type state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValBIssuanceRequestState(t *testing.T) {
	t.Run("valid claim issuance request active", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		if model.IssuanceRequestState != Point11ValBIssuanceRequestStateActive {
			t.Fatalf("expected active issuance request, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValBClaimIssuanceRequest)
	}{
		{name: "missing claim id blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.ClaimID = ""
		}},
		{name: "claim unknown blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.ClaimID = "claim_unknown"
		}},
		{name: "claim revoked blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.ClaimID = "claim_revoked"
		}},
		{name: "claim invalid blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.ClaimID = "claim_invalid"
		}},
		{name: "disallowed subject kind blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.SubjectKind = "cluster"
		}},
		{name: "disallowed issuer kind blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.IssuerKind = "customer"
		}},
		{name: "disallowed audience blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.Audience = "public_ad_hoc"
		}},
		{name: "disallowed publication surface blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.PublicationSurface = point11Val0PublicationSurfaceAgentOutput
		}},
		{name: "missing policy basis blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.PolicyBasisRef = ""
		}},
		{name: "inactive policy basis blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.PolicyBasisState = "policy_basis_inactive"
		}},
		{name: "missing evidence refs blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.EvidenceRefs = nil
		}},
		{name: "unsupported verification method blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.VerificationMethod = "unsupported_method"
		}},
		{name: "missing freshness class blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.FreshnessClass = ""
		}},
		{name: "expired expires at blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.ExpiresAt = time.Now().UTC().Add(-time.Hour).Format(time.RFC3339)
		}},
		{name: "public customer visible without governance event blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.GovernanceEventRef = ""
			model.RequestedClaimCategory = Point11Val0ClaimCategoryCustomerVisible
		}},
		{name: "proposer equals final approver for export claim blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.ProposerRef = "internal_final_approver"
		}},
		{name: "partner self approval blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.ProposerRef = "partner_operator"
			model.ApproverRef = "partner_operator"
		}},
		{name: "customer self approval blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.ProposerRef = "customer_operator"
			model.ApproverRef = "customer_operator"
		}},
		{name: "agent self approval blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.ProposerRef = "agent_operator"
			model.ApproverRef = "agent_operator"
		}},
		{name: "agent origin output to customer visible public claim without governance event blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.AgentOrigin = true
			model.GovernanceEventRef = ""
			model.RequestedClaimCategory = Point11Val0ClaimCategoryPublicSafe
		}},
		{name: "requested revoked lifecycle blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.RequestedLifecycleState = Point11Val0ClaimLifecycleRevoked
		}},
		{name: "requested expired lifecycle blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.RequestedLifecycleState = point11ValBClaimLifecycleExpired
		}},
		{name: "requested superseded lifecycle blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.RequestedLifecycleState = Point11Val0ClaimLifecycleSuperseded
		}},
		{name: "forbidden overclaim wording blocks", mutate: func(model *Point11ValBClaimIssuanceRequest) {
			model.ClaimTypeName = "production approved claim"
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValBFoundation()
			testCase.mutate(&model.IssuanceRequest)
			model = ComputePoint11ValBFoundation(model)
			if model.IssuanceRequestState != Point11ValBIssuanceRequestStateBlocked {
				t.Fatalf("expected blocked issuance request state, got %#v", model)
			}
		})
	}

	for _, surface := range []string{
		point11Val0PublicationSurfaceDocs,
		point11Val0PublicationSurfacePortal,
		point11Val0PublicationSurfaceExport,
		point11Val0PublicationSurfacePartner,
		point11Val0PublicationSurfaceDemo,
		point11Val0PublicationSurfaceSales,
		point11Val0PublicationSurfaceBuyer,
	} {
		t.Run("surface "+surface+" without clean room review blocks", func(t *testing.T) {
			model := activePoint11ValBFoundation()
			model.IssuanceRequest.PublicationSurface = surface
			model.IssuanceRequest.CleanRoomIPReviewRef = ""
			model = ComputePoint11ValBFoundation(model)
			if model.IssuanceRequestState != Point11ValBIssuanceRequestStateBlocked {
				t.Fatalf("expected missing clean-room review to block, got %#v", model)
			}
		})

		t.Run("surface "+surface+" with clean room review can pass", func(t *testing.T) {
			model := activePoint11ValBFoundation()
			model.IssuanceRequest.PublicationSurface = surface
			model.IssuanceRequest.CleanRoomIPReviewRef = "clean_room_review_point11_valb_surface_001"
			model = ComputePoint11ValBFoundation(model)
			if model.IssuanceRequestState != Point11ValBIssuanceRequestStateActive {
				t.Fatalf("expected issuance request to remain active, got %#v", model)
			}
		})
	}
}

func TestPoint11ValBIssuedClaimState(t *testing.T) {
	t.Run("valid issued claim active", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		if model.IssuedClaimState != Point11ValBIssuedClaimStateActive {
			t.Fatalf("expected active issued claim, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValBIssuedClaimRecord)
	}{
		{name: "missing claim version blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.ClaimVersion = ""
		}},
		{name: "global scope blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.Scope = "global_scope"
		}},
		{name: "missing evidence hash refs blocks where required", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.EvidenceHashRefs = nil
		}},
		{name: "verification result inactive blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.VerificationResult = point11ValBVerificationResultInactive
		}},
		{name: "expired claim blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.ExpiresAt = time.Now().UTC().Add(-time.Hour).Format(time.RFC3339)
		}},
		{name: "revoked lifecycle cannot be active use claim", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.LifecycleState = Point11Val0ClaimLifecycleRevoked
			model.RevocationRef = "governance_event_point11_valb_revocation"
		}},
		{name: "superseded lifecycle cannot be active use claim", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.LifecycleState = Point11Val0ClaimLifecycleSuperseded
			model.SupersededBy = "claim_point11_valb_successor_001"
		}},
		{name: "corrected lifecycle cannot be active use claim", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.LifecycleState = Point11Val0ClaimLifecycleCorrected
			model.CorrectionRef = "governance_event_point11_valb_correction"
		}},
		{name: "revoked lifecycle without revocation ref blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.LifecycleState = Point11Val0ClaimLifecycleRevoked
			model.RevocationRef = ""
		}},
		{name: "superseded lifecycle without valid superseded by blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.LifecycleState = Point11Val0ClaimLifecycleSuperseded
			model.SupersededBy = ""
		}},
		{name: "corrected lifecycle without correction ref blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.LifecycleState = Point11Val0ClaimLifecycleCorrected
			model.CorrectionRef = ""
		}},
		{name: "customer public export claim without governance event blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.GovernanceEventRef = ""
		}},
		{name: "customer public export claim without clean room review blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.CleanRoomIPReviewRef = ""
		}},
		{name: "missing audit id blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.AuditID = ""
		}},
		{name: "malformed projection disclaimer blocks", mutate: func(model *Point11ValBIssuedClaimRecord) {
			model.ProjectionDisclaimer = "canonical_truth"
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValBFoundation()
			testCase.mutate(&model.IssuedClaim)
			model = ComputePoint11ValBFoundation(model)
			if model.IssuedClaimState != Point11ValBIssuedClaimStateBlocked {
				t.Fatalf("expected blocked issued claim state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValBRegistryState(t *testing.T) {
	t.Run("valid claim registry active", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		if model.RegistryState != Point11ValBRegistryStateActive {
			t.Fatalf("expected active claim registry, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValBFoundation)
	}{
		{name: "missing registry id blocks", mutate: func(model *Point11ValBFoundation) {
			model.Registry.RegistryID = ""
		}},
		{name: "claim registry unknown blocks", mutate: func(model *Point11ValBFoundation) {
			model.Registry.RegistryID = "claim_registry_unknown"
		}},
		{name: "claim registry revoked blocks", mutate: func(model *Point11ValBFoundation) {
			model.Registry.RegistryID = "claim_registry_revoked"
		}},
		{name: "claim registry invalid blocks", mutate: func(model *Point11ValBFoundation) {
			model.Registry.RegistryID = "claim_registry_invalid"
		}},
		{name: "active claim refs containing revoked claim blocks", mutate: func(model *Point11ValBFoundation) {
			model.IssuedClaim.LifecycleState = Point11Val0ClaimLifecycleRevoked
			model.IssuedClaim.RevocationRef = "governance_event_point11_valb_revocation"
			model.Registry.RegisteredClaims = []Point11ValBIssuedClaimRecord{model.IssuedClaim}
		}},
		{name: "active claim refs containing superseded claim blocks", mutate: func(model *Point11ValBFoundation) {
			model.IssuedClaim.LifecycleState = Point11Val0ClaimLifecycleSuperseded
			model.IssuedClaim.SupersededBy = "claim_point11_valb_successor_001"
			model.Registry.RegisteredClaims = []Point11ValBIssuedClaimRecord{model.IssuedClaim}
		}},
		{name: "active claim refs containing blocked claim blocks", mutate: func(model *Point11ValBFoundation) {
			model.IssuedClaim.LifecycleState = Point11Val0ClaimLifecycleBlocked
			model.Registry.RegisteredClaims = []Point11ValBIssuedClaimRecord{model.IssuedClaim}
		}},
		{name: "duplicate claim identity blocks", mutate: func(model *Point11ValBFoundation) {
			model.Registry.RegisteredClaims = []Point11ValBIssuedClaimRecord{model.IssuedClaim, model.IssuedClaim}
		}},
		{name: "conflicting claim identity blocks", mutate: func(model *Point11ValBFoundation) {
			other := model.IssuedClaim
			other.SubjectRef = "subject_point11_valb_workload_beta"
			model.Registry.RegisteredClaims = []Point11ValBIssuedClaimRecord{model.IssuedClaim, other}
		}},
		{name: "cross tenant claim reuse blocks", mutate: func(model *Point11ValBFoundation) {
			other := model.IssuedClaim
			other.Scope = "tenant_scope_beta"
			model.Registry.RegisteredClaims = []Point11ValBIssuedClaimRecord{model.IssuedClaim, other}
		}},
		{name: "same claim id with different subject policy evidence hash without governance event blocks", mutate: func(model *Point11ValBFoundation) {
			other := model.IssuedClaim
			other.SubjectRef = "subject_point11_valb_workload_beta"
			other.PolicyBasisRef = "policy_point11_vala_authority_core_v2"
			other.EvidenceHashRefs = []string{"evidence_hash_point11_valb_conflict_001"}
			other.GovernanceEventRef = ""
			model.Registry.RegisteredClaims = []Point11ValBIssuedClaimRecord{model.IssuedClaim, other}
		}},
		{name: "missing governance event ref blocks", mutate: func(model *Point11ValBFoundation) {
			model.Registry.GovernanceEventRef = ""
		}},
		{name: "missing audit id blocks", mutate: func(model *Point11ValBFoundation) {
			model.Registry.AuditID = ""
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValBFoundation()
			testCase.mutate(&model)
			model = ComputePoint11ValBFoundation(model)
			if model.RegistryState != Point11ValBRegistryStateBlocked {
				t.Fatalf("expected blocked registry state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValBVerificationState(t *testing.T) {
	t.Run("valid claim verification active", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		if model.VerificationState != Point11ValBVerificationStateActive {
			t.Fatalf("expected active claim verification, got %#v", model)
		}
	})

	setHashContract := func(model *Point11ValBFoundation, required bool, verificationHashes, claimHashes []string) {
		model.IssuedClaim.EvidenceHashRequired = required
		model.IssuedClaim.EvidenceHashRefs = append([]string{}, claimHashes...)
		model.Verification.EvidenceHashRefs = append([]string{}, verificationHashes...)
		model.Verification.ClaimEvidenceHashRefs = append([]string{}, claimHashes...)
		if len(model.Registry.RegisteredClaims) > 0 {
			model.Registry.RegisteredClaims[0].EvidenceHashRequired = required
			model.Registry.RegisteredClaims[0].EvidenceHashRefs = append([]string{}, claimHashes...)
		}
	}

	t.Run("hash optional issued claim with empty evidence hash refs can verify active", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		setHashContract(&model, false, nil, nil)
		model = ComputePoint11ValBFoundation(model)
		if model.VerificationState != Point11ValBVerificationStateActive {
			t.Fatalf("expected hash-optional verification to remain active without hashes, got %#v", model)
		}
	})

	t.Run("hash required issued claim with empty evidence hash refs blocks verification", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		setHashContract(&model, true, nil, nil)
		model = ComputePoint11ValBFoundation(model)
		if model.VerificationState != Point11ValBVerificationStateBlocked {
			t.Fatalf("expected hash-required verification to block without hashes, got %#v", model)
		}
	})

	t.Run("hash required issued claim with mismatched evidence hash refs blocks verification", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		setHashContract(&model, true, []string{"evidence_hash_point11_valb_verification_only"}, []string{"evidence_hash_point11_valb_claim_only"})
		model = ComputePoint11ValBFoundation(model)
		if model.VerificationState != Point11ValBVerificationStateBlocked {
			t.Fatalf("expected hash-required mismatch to block verification, got %#v", model)
		}
	})

	t.Run("hash required issued claim with matching evidence hash refs verifies active", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		hashes := []string{"evidence_hash_point11_valb_claim_001", "evidence_hash_point11_valb_claim_002"}
		setHashContract(&model, true, hashes, hashes)
		model = ComputePoint11ValBFoundation(model)
		if model.VerificationState != Point11ValBVerificationStateActive {
			t.Fatalf("expected hash-required matching hashes to verify active, got %#v", model)
		}
	})

	t.Run("hash required claim with padded matching evidence hash refs blocks raw exact", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		hashes := []string{" evidence_hash_point11_valb_claim_001"}
		setHashContract(&model, true, hashes, hashes)
		model = ComputePoint11ValBFoundation(model)
		if model.VerificationState != Point11ValBVerificationStateBlocked {
			t.Fatalf("expected padded matching hashes to block verification, got %#v", model)
		}
		if !point11Val0ContainsTrimmed(model.Diagnostics.VerificationReasons, "claim_verification_evidence_hash_refs_mismatch") {
			t.Fatalf("expected exact evidence hash mismatch reason, got %#v", model.Diagnostics.VerificationReasons)
		}
	})

	t.Run("hash optional claim with one sided hash refs blocks", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		setHashContract(&model, false, []string{"evidence_hash_point11_valb_claim_001"}, nil)
		model = ComputePoint11ValBFoundation(model)
		if model.VerificationState != Point11ValBVerificationStateBlocked {
			t.Fatalf("expected one-sided optional hashes to block verification, got %#v", model)
		}
	})

	t.Run("hash optional claim with both hash ref sets present but mismatched blocks", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		setHashContract(&model, false, []string{"evidence_hash_point11_valb_verification_only"}, []string{"evidence_hash_point11_valb_claim_only"})
		model = ComputePoint11ValBFoundation(model)
		if model.VerificationState != Point11ValBVerificationStateBlocked {
			t.Fatalf("expected optional mismatched hashes to block verification, got %#v", model)
		}
	})

	t.Run("hash optional claim with both hash ref sets present and matching verifies active", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		hashes := []string{"evidence_hash_point11_valb_claim_001"}
		setHashContract(&model, false, hashes, hashes)
		model = ComputePoint11ValBFoundation(model)
		if model.VerificationState != Point11ValBVerificationStateActive {
			t.Fatalf("expected optional matching hashes to verify active, got %#v", model)
		}
	})

	t.Run("evidence refs mismatch still blocks even when hashes are optional", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		setHashContract(&model, false, nil, nil)
		model.Verification.EvidenceRefs = []string{"evidence:point11-valb-mismatch-001"}
		model = ComputePoint11ValBFoundation(model)
		if model.VerificationState != Point11ValBVerificationStateBlocked {
			t.Fatalf("expected evidence ref mismatch to block even when hashes are optional, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point11ValBClaimVerificationResult)
	}{
		{name: "missing verification id blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.VerificationID = ""
		}},
		{name: "verification unknown blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.VerificationID = "verification_unknown"
		}},
		{name: "verification revoked blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.VerificationID = "verification_revoked"
		}},
		{name: "verification invalid blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.VerificationID = "verification_invalid"
		}},
		{name: "claim not found in registry blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.ClaimRegistered = false
		}},
		{name: "policy basis inactive blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.PolicyBasisState = "policy_basis_inactive"
		}},
		{name: "evidence refs mismatch blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.EvidenceRefs = []string{"evidence:point11-valb-mismatch-001"}
		}},
		{name: "evidence hash refs mismatch blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.EvidenceHashRefs = []string{"evidence_hash_point11_valb_mismatch_001"}
		}},
		{name: "unsupported verification method blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.VerificationMethod = "unsupported"
		}},
		{name: "stale freshness blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.FreshnessResult = point11ValBFreshnessResultStale
		}},
		{name: "revoked claim check blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.RevocationCheckResult = point11ValBRevocationCheckBlocked
		}},
		{name: "superseded claim check blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.SupersessionCheckResult = point11ValBSupersessionCheckBlocked
		}},
		{name: "scope check blocked blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.ScopeCheckResult = point11ValBScopeCheckBlocked
		}},
		{name: "audience check blocked blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.AudienceCheckResult = point11ValBAudienceCheckBlocked
		}},
		{name: "issuer trust blocked blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.IssuerTrustResult = point11ValBIssuerTrustBlocked
		}},
		{name: "cross domain compatibility blocked blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.CrossDomainRequired = true
			model.CrossDomainCompatibilityResult = point11ValBCrossDomainCompatibilityBlocked
		}},
		{name: "overclaim wording in diagnostics blocks", mutate: func(model *Point11ValBClaimVerificationResult) {
			model.Diagnostics = []string{"official authority"}
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValBFoundation()
			testCase.mutate(&model.Verification)
			model = ComputePoint11ValBFoundation(model)
			if model.VerificationState != Point11ValBVerificationStateBlocked {
				t.Fatalf("expected blocked verification state, got %#v", model)
			}
		})
	}
}

func TestPoint11ValBCrossDomainState(t *testing.T) {
	t.Run("valid remote claim intake active under explicit trust rules", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		if model.CrossDomainIntakeState != Point11ValBCrossDomainIntakeStateActive {
			t.Fatalf("expected active cross-domain intake, got %#v", model)
		}
	})

	testCases := []struct {
		name      string
		mutate    func(*Point11ValBCrossDomainClaimIntake)
		wantState string
	}{
		{name: "missing remote trust root ref blocks", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.RemoteTrustRootRef = ""
		}, wantState: Point11ValBCrossDomainIntakeStateBlocked},
		{name: "trust root unknown blocks", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.RemoteTrustRootRef = "trust_root_unknown"
		}, wantState: Point11ValBCrossDomainIntakeStateBlocked},
		{name: "trust root placeholder blocks", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.RemoteTrustRootRef = "trust_root_placeholder"
		}, wantState: Point11ValBCrossDomainIntakeStateBlocked},
		{name: "unknown remote issuer blocks", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.RemoteIssuerRef = "issuer_unknown"
		}, wantState: Point11ValBCrossDomainIntakeStateBlocked},
		{name: "incompatible scope yields review required", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.ScopeCompatibilityResult = point11ValBGenericCompatibilityReviewRequired
		}, wantState: Point11ValBCrossDomainIntakeStateReviewRequired},
		{name: "incompatible freshness yields review required", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.FreshnessCompatibilityResult = point11ValBGenericCompatibilityReviewRequired
		}, wantState: Point11ValBCrossDomainIntakeStateReviewRequired},
		{name: "revoked remote claim blocks", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.RemoteClaimState = Point11Val0ClaimLifecycleRevoked
		}, wantState: Point11ValBCrossDomainIntakeStateBlocked},
		{name: "expired remote claim blocks", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.RemoteClaimState = point11ValBClaimLifecycleExpired
		}, wantState: Point11ValBCrossDomainIntakeStateBlocked},
		{name: "superseded remote claim blocks", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.RemoteClaimState = Point11Val0ClaimLifecycleSuperseded
		}, wantState: Point11ValBCrossDomainIntakeStateBlocked},
		{name: "remote claim cannot override local policy", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.RemoteOverridesLocalPolicy = true
		}, wantState: Point11ValBCrossDomainIntakeStateBlocked},
		{name: "remote claim cannot create certification legal regulatory authority", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.CreatesCertificationAuthority = true
		}, wantState: Point11ValBCrossDomainIntakeStateBlocked},
		{name: "missing local admissibility result blocks", mutate: func(model *Point11ValBCrossDomainClaimIntake) {
			model.LocalAdmissibilityResult = ""
		}, wantState: Point11ValBCrossDomainIntakeStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValBFoundation()
			testCase.mutate(&model.CrossDomainIntake)
			model = ComputePoint11ValBFoundation(model)
			if model.CrossDomainIntakeState != testCase.wantState {
				t.Fatalf("expected cross-domain state %q, got %#v", testCase.wantState, model)
			}
		})
	}
}

func TestPoint11ValBAggregateState(t *testing.T) {
	t.Run("aggregate active only when all components active", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		if model.CurrentState != Point11ValBStateActive {
			t.Fatalf("expected active aggregate state, got %#v", model)
		}
	})

	t.Run("dependency review required yields aggregate review required if no local blockers", func(t *testing.T) {
		model := Point11ValBFoundationModel()
		model.Dependency = reviewRequiredPoint11ValBDependencySnapshot()
		model = ComputePoint11ValBFoundation(model)
		if model.CurrentState != Point11ValBStateReviewRequired {
			t.Fatalf("expected review required aggregate state, got %#v", model)
		}
	})

	t.Run("cross domain review required yields aggregate review required if no local blockers", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		model.CrossDomainIntake.ScopeCompatibilityResult = point11ValBGenericCompatibilityReviewRequired
		model = ComputePoint11ValBFoundation(model)
		if model.CurrentState != Point11ValBStateReviewRequired {
			t.Fatalf("expected review required aggregate state, got %#v", model)
		}
	})

	t.Run("any local component blocked yields aggregate blocked", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		model.IssuanceRequest.PolicyBasisRef = ""
		model = ComputePoint11ValBFoundation(model)
		if model.CurrentState != Point11ValBStateBlocked {
			t.Fatalf("expected blocked aggregate state, got %#v", model)
		}
	})

	t.Run("dependency review required cannot mask local issued claim blocker", func(t *testing.T) {
		model := Point11ValBFoundationModel()
		model.Dependency = reviewRequiredPoint11ValBDependencySnapshot()
		model.IssuedClaim.ClaimVersion = ""
		model = ComputePoint11ValBFoundation(model)
		if model.DependencyState != Point11ValBDependencyStateReviewRequired {
			t.Fatalf("expected dependency review required state, got %#v", model)
		}
		if model.IssuedClaimState != Point11ValBIssuedClaimStateBlocked {
			t.Fatalf("expected local issued claim blocker, got %#v", model)
		}
		if model.CurrentState != Point11ValBStateBlocked {
			t.Fatalf("expected local issued claim blocker to override dependency review required, got %#v", model)
		}
	})

	t.Run("cross domain review required cannot mask local registry blocker", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		model.CrossDomainIntake.ScopeCompatibilityResult = point11ValBGenericCompatibilityReviewRequired
		model.Registry.AuditID = ""
		model = ComputePoint11ValBFoundation(model)
		if model.CrossDomainIntakeState != Point11ValBCrossDomainIntakeStateReviewRequired {
			t.Fatalf("expected cross-domain review required state, got %#v", model)
		}
		if model.RegistryState != Point11ValBRegistryStateBlocked {
			t.Fatalf("expected local registry blocker, got %#v", model)
		}
		if model.CurrentState != Point11ValBStateBlocked {
			t.Fatalf("expected local registry blocker to override cross-domain review required, got %#v", model)
		}
	})

	t.Run("diagnostics include component blocking reason", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		model.Registry.AuditID = ""
		model = ComputePoint11ValBFoundation(model)
		if !strings.Contains(strings.Join(model.Diagnostics.RegistryReasons, " "), "audit") {
			t.Fatalf("expected registry diagnostics to include blocking reason, got %#v", model.Diagnostics)
		}
	})

	t.Run("aggregate does not emit point 11 pass", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal aggregate: %v", err)
		}
		passMarker := "point_11_" + "pass"
		if strings.Contains(string(body), passMarker) {
			t.Fatalf("expected no point 11 pass emission, got %s", body)
		}
	})

	t.Run("aggregate does not create legal regulatory certification authority", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		model.CreatesAuthorityClaims = true
		model = ComputePoint11ValBFoundation(model)
		if model.CurrentState != Point11ValBStateBlocked {
			t.Fatalf("expected authority marker to block aggregate, got %#v", model)
		}
	})

	t.Run("aggregate does not create publication side effects", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		model.CreatesPublicationSideEffects = true
		model = ComputePoint11ValBFoundation(model)
		if model.CurrentState != Point11ValBStateBlocked {
			t.Fatalf("expected publication side-effect marker to block aggregate, got %#v", model)
		}
	})

	t.Run("aggregate blocks padded active local state instead of normalizing", func(t *testing.T) {
		model := activePoint11ValBFoundation()
		model.VerificationState = " " + Point11ValBVerificationStateActive + " "
		if got := EvaluatePoint11ValBFoundationState(model); got != Point11ValBStateBlocked {
			t.Fatalf("expected padded verification state to block aggregate, got %q for %#v", got, model)
		}
		if !point11Val0ContainsTrimmed(point11ValBBlockingReasons(model), "claim_verification_blocked") {
			t.Fatalf("expected exact claim verification blocking reason, got %#v", point11ValBBlockingReasons(model))
		}
	})
}

func TestPoint11ValBSemanticAntiGreenRefs(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point11ValBFoundation)
	}{
		{name: "claim unknown blocks even though it has a valid prefix", mutate: func(model *Point11ValBFoundation) {
			model.IssuanceRequest.ClaimID = "claim_unknown"
		}},
		{name: "claim revoked blocks even though it has a valid prefix", mutate: func(model *Point11ValBFoundation) {
			model.IssuanceRequest.ClaimID = "claim_revoked"
		}},
		{name: "issuer invalid blocks even though it has a valid prefix", mutate: func(model *Point11ValBFoundation) {
			model.IssuanceRequest.IssuerRef = "issuer_invalid"
		}},
		{name: "claim registry expired blocks even though it has a valid prefix", mutate: func(model *Point11ValBFoundation) {
			model.Registry.RegistryID = "claim_registry_expired"
		}},
		{name: "trust root placeholder blocks even though it has a valid prefix", mutate: func(model *Point11ValBFoundation) {
			model.CrossDomainIntake.RemoteTrustRootRef = "trust_root_placeholder"
		}},
		{name: "revoked invalid marker blocks in security critical claim ref field", mutate: func(model *Point11ValBFoundation) {
			model.IssuanceRequest.ClaimID = "revoked/invalid marker"
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint11ValBFoundation()
			testCase.mutate(&model)
			model = ComputePoint11ValBFoundation(model)
			if model.CurrentState != Point11ValBStateBlocked {
				t.Fatalf("expected blocked aggregate state, got %#v", model)
			}
		})
	}
}

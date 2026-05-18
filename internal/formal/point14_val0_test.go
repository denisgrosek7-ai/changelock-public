package formal

import (
	"testing"

	"github.com/denisgrosek/changelock/internal/operability"
)

func TestPoint14Val0DependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14Val0DependencySnapshot)
		want   string
	}{
		{
			name: "missing point13 vale closure blocks",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point13ValECurrentState = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "point13 vale blocked blocks",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point13ValECurrentState = Point13ValEStateBlocked
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "point13 vale review required prevents active",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point13ValECurrentState = Point13ValEStateReviewRequired
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "point13 vale incomplete prevents active",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point13ValECurrentState = Point13ValEStateIncomplete
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "premature point14 pass blocks",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point14PassSeen = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "local readiness cannot override missing vale closure",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point13ValEPassAllowed = true
				model.Point13ValEPassToken = point13ValEPoint13PassToken
				model.Point13ValECurrentState = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "embedded vale mismatch blocks",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point13ValE.CurrentState = Point13ValEStateBlocked
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "embedded point12 reviewer mismatch blocks",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point12.PassClosureManifest.ReviewerResult = point12ValEReviewerResultPass
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "embedded point11 publication mismatch blocks",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point11.PublicationReviewState = Point11ValDPublicationReviewStateBlocked
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "embedded point11 dependency state mismatch blocks",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point11.DependencyState = Point11ValDDependencyStateBlocked
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "embedded point11 val0 dependency cannot forge active summary",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point11.Val0Dependency.CurrentState = Point11Val0StateBlocked
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "non active point11 aggregate blocks even when final gate is active",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.InheritedPoint11CurrentState = Point11ValDStateReviewRequired
				model.Point11.CurrentState = Point11ValDStateReviewRequired
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "embedded point11 final pass gate mismatch blocks",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point11.FinalPassGateState = Point11ValDFinalPassGateStateBlocked
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "embedded point11 no overclaim submodel cannot forge active summary",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point11.NoOverclaimReview.ObservedClaims = []string{"production approved"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "inherited point11 final pass gate must be active raw exact",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.InheritedPoint11FinalPassGateState = Point11ValDFinalPassGateStateBlocked
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "whitespace retagged inherited point11 final pass gate blocks",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.InheritedPoint11FinalPassGateState = " " + Point11ValDFinalPassGateStateActive
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "embedded point10 pass rule mismatch blocks",
			mutate: func(model *Point14Val0DependencySnapshot) {
				model.Point10.Point10PassRuleState = operability.DeploymentMultiTenantValEPoint10PassRuleStateBlocked
			},
			want: Point14Val0StateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14Val0DependencySnapshotModel()
			tc.mutate(&model)
			if got := EvaluatePoint14Val0DependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14Val0ExternalSignalCandidateState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalEcosystemSignalCandidate)
		want   string
	}{
		{
			name:   "valid external signal candidate active only in candidate scope",
			mutate: func(model *ExternalEcosystemSignalCandidate) {},
			want:   Point14Val0StateActive,
		},
		{
			name: "missing signal id blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.SignalID = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "missing source type blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.SourceType = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "missing signal type blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.SignalType = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "missing evidence refs blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = nil
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "global evidence ref blocks raw evidence boundary",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{"evidence_global_point14_val0_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "singular all tenant evidence ref blocks raw evidence boundary",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{"evidence_all_tenant_point14_val0_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "padded evidence ref blocks raw evidence boundary",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{" " + model.EvidenceRefs[0]}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "zero width source ref cannot validate even with recomputed identity key",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.SourceRef += "\u200b"
				model.SignalIdentityKey = point14Val0SignalIdentityKey(*model)
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "tenant scoped without tenant scope blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.TenantScope = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "unknown source type blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.SourceType = "mystery_signal_source"
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "unknown signal type blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.SignalType = "mystery_signal_type"
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "unknown validation status blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ValidationStatus = "mystery_status"
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "canonical authority blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.CanonicalAuthority = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "pass allowed blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.PassAllowed = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "canonical mutation blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.CanonicalMutationAllowed = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "production mutation blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ProductionMutationAllowed = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "override canonical decision blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.OverrideCanonicalDecision = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "provenance pending is review required",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ValidationStatus = point14Val0ValidationProvenancePending
			},
			want: Point14Val0StateReviewRequired,
		},
		{
			name: "conflicting signal is review required",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ValidationStatus = point14Val0ValidationConflicting
			},
			want: Point14Val0StateReviewRequired,
		},
		{
			name: "superseded signal blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ValidationStatus = point14Val0ValidationSuperseded
				model.SupersededByRef = "signal_point14_val0_superseding_001"
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "duplicate signal identity cannot create duplicate active evidence",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.DuplicateSignalRefs = []string{"signal_point14_val0_duplicate_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "unrelated external signal cannot bind to artifact claim",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ArtifactBindingConsistent = false
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "global public signal cannot bypass tenant boundary",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ScopeClassification = point14Val0ScopePublicNonAuthorative
				model.TenantScope = ""
				model.ReferencedTenantScope = "tenant_point14_val0_other"
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "cross tenant external signal blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ReferencedTenantScope = "tenant_point14_val0_other"
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "cross tenant evidence ref blocks directly",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{"evidence_cross_tenant_point14_val0_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "camelcase compact cross tenant evidence ref blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{"evidence_crossTenant_point14_val0_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "camelcase compact cross tenants evidence ref blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{"evidence_crossTenants_point14_val0_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "cross scope evidence ref blocks downstream helper drift",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{"evidence_cross_scope_point14_val0_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "compact cross boundary evidence ref blocks downstream helper drift",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{"evidence_crossBoundary_point14_val0_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "camelcase compact other tenants evidence ref blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{"evidence_otherTenants_point14_val0_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "camelcase compact all tenant evidence ref blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{"evidence_allTenant_point14_val0_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "smalltenant evidence ref does not false positive as all tenant",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.EvidenceRefs = []string{"evidence_smalltenant_point14_val0_001"}
			},
			want: Point14Val0StateActive,
		},
		{
			name: "padded source ref cannot validate by trim",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.SourceRef += " "
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "padded artifact ref cannot validate by trim",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ArtifactRef = " " + model.ArtifactRef
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "tab newline hash ref cannot validate by trim",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.HashRef = model.HashRef + "\n"
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "client local received time blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ReceivedTimeSource = point14Val0TimeSourceClientLocal
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "whitespace retagged received time source blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ReceivedTimeSource = " " + point14Val0TimeSourceServerUTC
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "tab newline retagged received timestamp blocks",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.ReceivedAt = "\t" + model.ReceivedAt + "\n"
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "future dated advisory source event is review required",
			mutate: func(model *ExternalEcosystemSignalCandidate) {
				model.SourceEventAt = "2026-05-05T10:15:00Z"
			},
			want: Point14Val0StateReviewRequired,
		},
	}
	dependency := point14Val0DependencySnapshotModel()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14Val0ExternalSignalCandidateModel(dependency)
			tc.mutate(&model)
			if tc.name != "duplicate signal identity cannot create duplicate active evidence" {
				model.SignalIdentityKey = point14Val0SignalIdentityKey(model)
			}
			if got := EvaluatePoint14Val0ExternalSignalCandidateState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14Val0ExternalStakeholderAuthorityRoleState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalStakeholderAuthorityRole)
		want   string
	}{
		{
			name: "allowed stakeholder roles validate with bounded authority only",
			mutate: func(model *ExternalStakeholderAuthorityRole) {
				model.RoleType = "scanner"
				model.AllowedActionRefs = []string{"submit_evidence", "submit_annotation"}
			},
			want: Point14Val0StateActive,
		},
		{
			name: "unknown role blocks",
			mutate: func(model *ExternalStakeholderAuthorityRole) {
				model.RoleType = "global_super_reviewer"
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "scanner cannot emit pass",
			mutate: func(model *ExternalStakeholderAuthorityRole) {
				model.RoleType = "scanner"
				model.AllowedActionRefs = []string{"emit_pass"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "vex issuer cannot override canonical decision",
			mutate: func(model *ExternalStakeholderAuthorityRole) {
				model.RoleType = "vex_issuer"
				model.AllowedActionRefs = []string{"override_canonical_evidence"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "maintainer cannot publish canonical correction without governance event",
			mutate: func(model *ExternalStakeholderAuthorityRole) {
				model.RoleType = "maintainer"
				model.AllowedActionRefs = []string{"publish_authoritative_correction"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "auditor verifier note cannot approve production",
			mutate: func(model *ExternalStakeholderAuthorityRole) {
				model.RoleType = "auditor"
				model.AllowedActionRefs = []string{"approve_production"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "partner customer admin signal cannot create production approval",
			mutate: func(model *ExternalStakeholderAuthorityRole) {
				model.RoleType = "partner"
				model.AllowedActionRefs = []string{"approve_production"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "external researcher report is evidence input only",
			mutate: func(model *ExternalStakeholderAuthorityRole) {
				model.RoleType = "external_researcher"
				model.AllowedActionRefs = []string{"submit_evidence"}
			},
			want: Point14Val0StateActive,
		},
		{
			name: "agent recommendation source is evidence input only",
			mutate: func(model *ExternalStakeholderAuthorityRole) {
				model.RoleType = "agent_recommendation_source"
				model.AllowedActionRefs = []string{"submit_evidence"}
			},
			want: Point14Val0StateActive,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14Val0ExternalStakeholderAuthorityRoleModel(point14Val0DependencySnapshotModel())
			tc.mutate(&model)
			if got := EvaluatePoint14Val0ExternalStakeholderAuthorityRoleState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14Val0ExternalAuthorityConflictMatrixState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalAuthorityConflictMatrix)
		want   string
	}{
		{
			name: "conflicting external signals require explicit governance path",
			mutate: func(model *ExternalAuthorityConflictMatrix) {
				model.ConflictPresent = true
				model.ConflictID = "conflict_point14_val0_scanner_vex_001"
				model.ConflictType = "vex_vs_scanner"
				model.SignalRefs = []string{"signal_point14_val0_vendor_001", "signal_point14_val0_vex_001"}
				model.RoleRefs = []string{"role_point14_val0_scanner_001", "role_point14_val0_vex_001"}
				model.AffectedArtifactRefs = []string{"artifact_point14_val0_component_001"}
				model.AffectedClaimRefs = []string{"claim_point14_val0_001"}
				model.AffectedEvidenceRefs = []string{"evidence_point14_val0_001"}
				model.GovernancePath = "canonical_review"
			},
			want: Point14Val0StateReviewRequired,
		},
		{
			name: "conflict without governance path blocks",
			mutate: func(model *ExternalAuthorityConflictMatrix) {
				model.ConflictPresent = true
				model.ConflictID = "conflict_point14_val0_missing_path_001"
				model.ConflictType = "scanner_vs_maintainer"
				model.SignalRefs = []string{"signal_point14_val0_vendor_001", "signal_point14_val0_maintainer_001"}
				model.RoleRefs = []string{"role_point14_val0_scanner_001", "role_point14_val0_maintainer_001"}
				model.AffectedEvidenceRefs = []string{"evidence_point14_val0_001"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "crowd public consensus cannot resolve dispute",
			mutate: func(model *ExternalAuthorityConflictMatrix) {
				model.ConflictPresent = true
				model.ConflictID = "conflict_point14_val0_crowd_001"
				model.ConflictType = "public_consensus_vs_canonical"
				model.SignalRefs = []string{"signal_point14_val0_vendor_001", "signal_point14_val0_crowd_001"}
				model.RoleRefs = []string{"role_point14_val0_maintainer_001", "role_point14_val0_crowd_001"}
				model.AffectedEvidenceRefs = []string{"evidence_point14_val0_001"}
				model.GovernancePath = "contradiction_review"
				model.ConsensusResolutionRequested = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "hidden resolution flags on no conflict record block",
			mutate: func(model *ExternalAuthorityConflictMatrix) {
				model.ConsensusResolutionRequested = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "no conflict record missing audit event blocks",
			mutate: func(model *ExternalAuthorityConflictMatrix) {
				model.AuditEventRef = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "no conflict record invalid tenant scope blocks",
			mutate: func(model *ExternalAuthorityConflictMatrix) {
				model.TenantScope = ""
			},
			want: Point14Val0StateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14Val0ExternalAuthorityConflictMatrixModel(point14Val0DependencySnapshotModel())
			tc.mutate(&model)
			if got := EvaluatePoint14Val0ExternalAuthorityConflictMatrixState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14Val0ExternalSignalDisputeLifecycleState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalSignalDisputeLifecycle)
		want   string
	}{
		{
			name:   "valid opened dispute records evidence refs and state",
			mutate: func(model *ExternalSignalDisputeLifecycle) {},
			want:   Point14Val0StateActive,
		},
		{
			name: "evidence required without evidence is incomplete",
			mutate: func(model *ExternalSignalDisputeLifecycle) {
				model.DisputeState = point14Val0DisputeEvidenceNeeded
				model.EvidenceRefs = nil
			},
			want: Point14Val0StateIncomplete,
		},
		{
			name: "corrected requires governance event",
			mutate: func(model *ExternalSignalDisputeLifecycle) {
				model.DisputeState = point14Val0DisputeCorrected
				model.CorrectedAt = "2026-05-05T09:30:00Z"
				model.GovernanceEventRef = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "revoked requires governance event",
			mutate: func(model *ExternalSignalDisputeLifecycle) {
				model.DisputeState = point14Val0DisputeRevoked
				model.RevokedAt = "2026-05-05T09:30:00Z"
				model.GovernanceEventRef = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "published notice requires visibility boundary and privacy check",
			mutate: func(model *ExternalSignalDisputeLifecycle) {
				model.DisputeState = point14Val0DisputePublished
				model.PublishedAt = "2026-05-05T09:30:00Z"
				model.VisibilityBoundaryRef = ""
				model.PrivacyCheckPassed = false
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "rejected dispute cannot silently delete evidence",
			mutate: func(model *ExternalSignalDisputeLifecycle) {
				model.DisputeState = point14Val0DisputeRejected
				model.RejectedDeletesEvidence = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "superseded requires supersession ref",
			mutate: func(model *ExternalSignalDisputeLifecycle) {
				model.DisputeState = point14Val0DisputeSuperseded
				model.SupersessionRef = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "dispute cannot mutate canonical evidence directly",
			mutate: func(model *ExternalSignalDisputeLifecycle) {
				model.CanonicalMutationRequested = true
			},
			want: Point14Val0StateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14Val0ExternalSignalDisputeLifecycleModel(point14Val0DependencySnapshotModel())
			tc.mutate(&model)
			if got := EvaluatePoint14Val0ExternalSignalDisputeLifecycleState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14Val0ExternalCorrectionRevocationBoundaryState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalCorrectionRevocationBoundary)
		want   string
	}{
		{
			name: "correction notice without approver governance event blocks",
			mutate: func(model *ExternalCorrectionRevocationBoundary) {
				model.ApproverRef = ""
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "correction notice leaking tenant private data blocks",
			mutate: func(model *ExternalCorrectionRevocationBoundary) {
				model.LeaksTenantPrivateData = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "public notice strengthening claim blocks",
			mutate: func(model *ExternalCorrectionRevocationBoundary) {
				model.StrengthensClaim = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "public notice implying badge certification blocks",
			mutate: func(model *ExternalCorrectionRevocationBoundary) {
				model.ObservedTexts = []string{"vendor-certified public badge"}
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "private limitation omitted from public correction blocks",
			mutate: func(model *ExternalCorrectionRevocationBoundary) {
				model.OmitsMeaningChangingLimitation = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "redacted correction cannot hide decisive missing evidence",
			mutate: func(model *ExternalCorrectionRevocationBoundary) {
				model.HidesDecisiveMissingEvidence = true
			},
			want: Point14Val0StateBlocked,
		},
		{
			name: "client time backdating blocks",
			mutate: func(model *ExternalCorrectionRevocationBoundary) {
				model.ClientTimeBackdated = true
			},
			want: Point14Val0StateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14Val0ExternalCorrectionRevocationBoundaryModel(point14Val0DependencySnapshotModel())
			tc.mutate(&model)
			if got := EvaluatePoint14Val0ExternalCorrectionRevocationBoundaryState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14Val0AgentBoundaryAndNoOverclaim(t *testing.T) {
	t.Run("agent recommendation may be included as advisory evidence input", func(t *testing.T) {
		model := point14Val0AgentEcosystemInputBoundaryModel(point14Val0DependencySnapshotModel())
		if got := EvaluatePoint14Val0AgentEcosystemInputBoundaryState(model); got != Point14Val0StateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
	t.Run("agent recommendation cannot resolve conflict", func(t *testing.T) {
		model := point14Val0AgentEcosystemInputBoundaryModel(point14Val0DependencySnapshotModel())
		model.CanDecideDispute = true
		if got := EvaluatePoint14Val0AgentEcosystemInputBoundaryState(model); got != Point14Val0StateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("ai agent authority flags block globally", func(t *testing.T) {
		model := point14Val0AgentEcosystemInputBoundaryModel(point14Val0DependencySnapshotModel())
		model.ExternalAuthorityAllowed = true
		if got := EvaluatePoint14Val0AgentEcosystemInputBoundaryState(model); got != Point14Val0StateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("forbidden ecosystem wording blocks", func(t *testing.T) {
		model := point14Val0NoOverclaimEcosystemWordingModel()
		model.ObservedPublicationTexts = []string{"production approved"}
		if got := EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model); got != Point14Val0StateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("forbidden ecosystem wording with tab retag blocks", func(t *testing.T) {
		model := point14Val0NoOverclaimEcosystemWordingModel()
		model.ObservedPublicationTexts = []string{"production\tapproved"}
		if got := EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model); got != Point14Val0StateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("forbidden ecosystem wording with confusable rune blocks", func(t *testing.T) {
		model := point14Val0NoOverclaimEcosystemWordingModel()
		model.ObservedPublicationTexts = []string{"prоduction approved"}
		if got := EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model); got != Point14Val0StateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("split forbidden ecosystem wording across observed corpus blocks", func(t *testing.T) {
		model := point14Val0NoOverclaimEcosystemWordingModel()
		model.ObservedPublicationTexts = []string{"production", "approved"}
		if got := EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model); got != Point14Val0StateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("split forbidden ecosystem wording across observed categories blocks", func(t *testing.T) {
		model := point14Val0NoOverclaimEcosystemWordingModel()
		model.ObservedPublicationTexts = []string{"production"}
		model.ObservedAgentTexts = []string{"approved"}
		if got := EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model); got != Point14Val0StateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("long-s pass wording obfuscation blocks", func(t *testing.T) {
		model := point14Val0NoOverclaimEcosystemWordingModel()
		model.ObservedPublicationTexts = []string{"external PAſS"}
		if got := EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model); got != Point14Val0StateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("split long-s as f overclaim across observed corpus blocks", func(t *testing.T) {
		model := point14Val0NoOverclaimEcosystemWordingModel()
		model.ObservedPublicationTexts = []string{"certiſ", "ied"}
		if got := EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model); got != Point14Val0StateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("safe bounded wording passes", func(t *testing.T) {
		model := point14Val0NoOverclaimEcosystemWordingModel()
		if got := EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model); got != Point14Val0StateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
	t.Run("forbidden wording allowed only in internal blocked diagnostic context", func(t *testing.T) {
		model := point14Val0NoOverclaimEcosystemWordingModel()
		model.InternalDiagnosticTexts = []string{"external PASS"}
		model.InternalDiagnosticsClassifiedBlocked = true
		if got := EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model); got != Point14Val0StateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
}

func TestPoint14Val0NoExternalAuthorityGuardState(t *testing.T) {
	model := point14Val0NoExternalAuthorityGuardModel()
	model.ObservedAuthorityMarkers = []string{"external_pass"}
	if got := EvaluatePoint14Val0NoExternalAuthorityGuardState(model); got != Point14Val0StateBlocked {
		t.Fatalf("expected blocked, got %s", got)
	}

	model = point14Val0NoExternalAuthorityGuardModel()
	model.ObservedAuthorityMarkers = []string{"external_source_of_truth\u200b"}
	if got := EvaluatePoint14Val0NoExternalAuthorityGuardState(model); got != Point14Val0StateBlocked {
		t.Fatalf("expected zero-width authority marker blocked, got %s", got)
	}

	model = point14Val0NoExternalAuthorityGuardModel()
	model.ObservedAuthorityMarkers = []string{"external_paſs"}
	if got := EvaluatePoint14Val0NoExternalAuthorityGuardState(model); got != Point14Val0StateBlocked {
		t.Fatalf("expected long-s authority marker blocked, got %s", got)
	}

	model = point14Val0NoExternalAuthorityGuardModel()
	model.ObservedAuthorityMarkers = []string{"external", "pass"}
	if got := EvaluatePoint14Val0NoExternalAuthorityGuardState(model); got != Point14Val0StateBlocked {
		t.Fatalf("expected split authority marker blocked, got %s", got)
	}
}

func TestPoint14Val0FoundationAggregation(t *testing.T) {
	t.Run("any blocked component yields blocked", func(t *testing.T) {
		model := Point14Val0FoundationModel()
		model.ExternalSignalCandidate.CanonicalAuthority = true
		model.ExternalSignalCandidate.SignalIdentityKey = point14Val0SignalIdentityKey(model.ExternalSignalCandidate)
		got := ComputePoint14Val0Foundation(model)
		if got.CurrentState != Point14Val0StateBlocked {
			t.Fatalf("expected blocked, got %#v", got)
		}
	})
	t.Run("any review required and no blocked yields review required", func(t *testing.T) {
		model := Point14Val0FoundationModel()
		model.ExternalAuthorityConflictMatrix.ConflictPresent = true
		model.ExternalAuthorityConflictMatrix.ConflictID = "conflict_point14_val0_review_001"
		model.ExternalAuthorityConflictMatrix.ConflictType = "vex_vs_scanner"
		model.ExternalAuthorityConflictMatrix.SignalRefs = []string{"signal_point14_val0_vendor_001", "signal_point14_val0_vex_001"}
		model.ExternalAuthorityConflictMatrix.RoleRefs = []string{"role_point14_val0_scanner_001", "role_point14_val0_vex_001"}
		model.ExternalAuthorityConflictMatrix.AffectedEvidenceRefs = []string{"evidence_point14_val0_001"}
		model.ExternalAuthorityConflictMatrix.GovernancePath = "canonical_review"
		got := ComputePoint14Val0Foundation(model)
		if got.CurrentState != Point14Val0StateReviewRequired {
			t.Fatalf("expected review required, got %#v", got)
		}
	})
	t.Run("any incomplete and no blocked review yields incomplete", func(t *testing.T) {
		model := Point14Val0FoundationModel()
		model.ExternalSignalDisputeLifecycle.DisputeState = point14Val0DisputeEvidenceNeeded
		model.ExternalSignalDisputeLifecycle.EvidenceRefs = nil
		got := ComputePoint14Val0Foundation(model)
		if got.CurrentState != Point14Val0StateIncomplete {
			t.Fatalf("expected incomplete, got %#v", got)
		}
	})
	t.Run("active only when all components active", func(t *testing.T) {
		model := ComputePoint14Val0Foundation(Point14Val0FoundationModel())
		if model.CurrentState != Point14Val0StateActive {
			t.Fatalf("expected active, got %#v", model)
		}
	})
}

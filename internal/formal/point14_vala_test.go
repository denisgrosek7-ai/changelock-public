package formal

import "testing"

func TestPoint14ValADependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValADependencySnapshot)
		want   string
	}{
		{
			name:   "canonical inherited boundary snapshot stays active",
			mutate: func(model *Point14ValADependencySnapshot) {},
			want:   Point14ValAStateActive,
		},
		{
			name: "missing point14 val0 blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0CurrentState = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 state in val0 dependency blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.Dependency.InheritedPoint10CurrentState += " "
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 no overclaim state in val0 dependency blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.Dependency.InheritedPoint10NoOverclaimState += " "
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 projection state in val0 dependency blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.Dependency.InheritedPoint10ProjectionState += " "
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 pass rule state in val0 dependency blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.Dependency.InheritedPoint10PassRuleState += " "
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "tab newline retagged embedded val0 current state blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.CurrentState = "\t" + model.Point14Val0.CurrentState + "\n"
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "whitespace retagged point13 pass token from embedded point13 blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point13ValE.Point13PassToken += " "
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "whitespace retagged point11 current state blocks raw exact",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.InheritedPoint11CurrentState += " "
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "whitespace retagged point11 final pass gate blocks raw exact",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.InheritedPoint11FinalPassGateState += " "
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "nested val0 point11 current state mismatch blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.Dependency.InheritedPoint11CurrentState = Point11ValDStateReviewRequired
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "nested val0 embedded point11 current state mismatch blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.Dependency.Point11.CurrentState = Point11ValDStateReviewRequired
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "nested val0 embedded point11 dependency state mismatch blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.Dependency.Point11.DependencyState = Point11ValDDependencyStateBlocked
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "nested val0 point11 final pass gate mismatch blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.Dependency.InheritedPoint11FinalPassGateState = Point11ValDFinalPassGateStateBlocked
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing point11 current state blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.InheritedPoint11CurrentState = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing point11 final pass gate blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.InheritedPoint11FinalPassGateState = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "point14 val0 blocked blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0CurrentState = Point14Val0StateBlocked
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "point14 val0 review required prevents active",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0CurrentState = Point14Val0StateReviewRequired
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "point14 val0 incomplete prevents active",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0CurrentState = Point14Val0StateIncomplete
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "premature point14 pass blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14PassSeen = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "local vala readiness cannot override missing val0 closure",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0CurrentState = ""
				model.InheritedPoint13ValEPassAllowed = true
				model.InheritedPoint13ValEPassToken = point13ValEPoint13PassToken
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "embedded point14 val0 mismatch blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.CurrentState = Point14Val0StateBlocked
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "stale active embedded val0 recomputation blocks",
			mutate: func(model *Point14ValADependencySnapshot) {
				model.Point14Val0.NoOverclaimEcosystemWording.ObservedAgentTexts = []string{"public badge"}
			},
			want: Point14ValAStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValADependencySnapshotModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValADependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValANormalizedExternalSignalState(t *testing.T) {
	dependency := point14ValADependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*NormalizedExternalSignal)
		want   string
	}{
		{
			name:   "valid external signal normalizes to bounded candidate not pass",
			mutate: func(model *NormalizedExternalSignal) {},
			want:   Point14ValAStateActive,
		},
		{
			name: "missing normalized signal id blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.NormalizedSignalID = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing original signal id blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.OriginalSignalID = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing source identity ref blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.SourceIdentityRef = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing source type blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.SourceType = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing signal type blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.SignalType = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing normalized payload hash blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.NormalizedPayloadHash = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing tenant scope global scope blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.TenantScope = ""
				model.GlobalScopeClassification = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "malformed normalized payload blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.PayloadNormalized = false
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "normalized signal with canonical authority blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.CanonicalAuthority = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "normalized signal with pass allowed blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.PassAllowed = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "normalized signal with production approved blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.ProductionApproved = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "normalized signal with public badge allowed blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.PublicBadgeAllowed = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "normalized signal with global evidence ref blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.EvidenceRefs = []string{"evidence_global_point14_vala_001"}
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "normalized signal with singular all tenant evidence ref blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.EvidenceRefs = []string{"evidence_all_tenant_point14_vala_001"}
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "normalized signal with cross tenant evidence ref blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.EvidenceRefs = []string{"evidence_cross_tenant_point14_vala_001"}
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "normalized signal with compact cross tenant evidence ref blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.EvidenceRefs = []string{"evidence_crossTenant_point14_vala_001"}
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "normalized signal with compact all tenant evidence ref blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.EvidenceRefs = []string{"evidence_allTenant_point14_vala_001"}
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "normalized signal with padded evidence ref blocks",
			mutate: func(model *NormalizedExternalSignal) {
				model.EvidenceRefs = []string{model.EvidenceRefs[0] + " "}
			},
			want: Point14ValAStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValANormalizedExternalSignalModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValANormalizedExternalSignalState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValAExternalSignalSourceIdentityState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalSignalSourceIdentity)
		want   string
	}{
		{
			name:   "valid bounded source identity passes",
			mutate: func(model *ExternalSignalSourceIdentity) {},
			want:   Point14ValAStateActive,
		},
		{
			name: "unknown source type blocks",
			mutate: func(model *ExternalSignalSourceIdentity) {
				model.SourceType = "mystery_source_type"
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "unsupported source review required",
			mutate: func(model *ExternalSignalSourceIdentity) {
				model.SourceSupported = false
			},
			want: Point14ValAStateReviewRequired,
		},
		{
			name: "missing source ref blocks",
			mutate: func(model *ExternalSignalSourceIdentity) {
				model.SourceRef = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing provenance custody hash where required blocks",
			mutate: func(model *ExternalSignalSourceIdentity) {
				model.ProvenanceRef = ""
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "source identity cannot create canonical authority",
			mutate: func(model *ExternalSignalSourceIdentity) {
				model.CanonicalAuthorityGranted = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "scanner source cannot emit pass",
			mutate: func(model *ExternalSignalSourceIdentity) {
				model.SourceType = "scanner_finding"
				model.PassAllowed = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "vex issuer source cannot override canonical decision",
			mutate: func(model *ExternalSignalSourceIdentity) {
				model.SourceType = "vex_statement"
				model.OverrideCanonicalDecision = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "auditor verifier source cannot approve production",
			mutate: func(model *ExternalSignalSourceIdentity) {
				model.SourceType = "auditor_note"
				model.ApproveProduction = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "partner customer admin source cannot certify or approve production",
			mutate: func(model *ExternalSignalSourceIdentity) {
				model.SourceType = "partner_signal"
				model.CertifyCompliance = true
			},
			want: Point14ValAStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValASourceIdentityModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValAExternalSignalSourceIdentityState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValAExternalSignalScopeBindingState(t *testing.T) {
	dependency := point14ValADependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*ExternalSignalScopeBinding)
		want   string
	}{
		{
			name:   "valid tenant scoped artifact binding passes",
			mutate: func(model *ExternalSignalScopeBinding) {},
			want:   Point14ValAStateActive,
		},
		{
			name: "exact artifact hash mismatch blocks",
			mutate: func(model *ExternalSignalScopeBinding) {
				model.ArtifactHashesMatch = false
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "exact claim ref mismatch blocks",
			mutate: func(model *ExternalSignalScopeBinding) {
				model.ClaimBindingExact = false
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "similar package name path does not bind",
			mutate: func(model *ExternalSignalScopeBinding) {
				model.SimilarPackageOnly = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "cross tenant binding blocks clb0",
			mutate: func(model *ExternalSignalScopeBinding) {
				model.ReferencedTenantScope = "tenant_point14_vala_other"
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing artifact refs for artifact scoped signal blocks",
			mutate: func(model *ExternalSignalScopeBinding) {
				model.ArtifactRefs = nil
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "missing claim refs for claim scoped signal blocks",
			mutate: func(model *ExternalSignalScopeBinding) {
				model.ClaimRefs = nil
			},
			want: Point14ValAStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValAScopeBindingModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValAExternalSignalScopeBindingState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValAExternalSignalEvidenceBindingState(t *testing.T) {
	dependency := point14ValADependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*ExternalSignalEvidenceBinding)
		want   string
	}{
		{
			name:   "valid evidence refs pass as candidate input",
			mutate: func(model *ExternalSignalEvidenceBinding) {},
			want:   Point14ValAStateActive,
		},
		{
			name: "missing evidence refs blocks where required",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceRefs = nil
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "global evidence ref blocks in evidence binding",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceRefs = []string{"evidence_global_point14_vala_001"}
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "cross tenant evidence ref blocks in evidence binding",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceRefs = []string{"evidence_cross_tenant_point14_vala_001"}
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "singular all tenant evidence ref blocks in evidence binding",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceRefs = []string{"evidence_all_tenant_point14_vala_001"}
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "compact cross tenant evidence ref blocks in evidence binding",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceRefs = []string{"evidence_crossTenant_point14_vala_001"}
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "padded evidence ref blocks in evidence binding",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceRefs = []string{model.EvidenceRefs[0] + " "}
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "stale evidence ref review required",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceState = point14ValAEvidenceStateStale
			},
			want: Point14ValAStateReviewRequired,
		},
		{
			name: "revoked evidence ref blocks",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceState = point14ValAEvidenceStateRevoked
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "expired evidence ref blocks",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceState = point14ValAEvidenceStateExpired
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "superseded evidence ref review required",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceState = point14ValAEvidenceStateSuperseded
			},
			want: Point14ValAStateReviewRequired,
		},
		{
			name: "unrelated evidence ref blocks",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceState = point14ValAEvidenceStateUnrelated
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "evidence hash mismatch blocks",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.EvidenceHashesMatch = false
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "canonical mutation flag blocks",
			mutate: func(model *ExternalSignalEvidenceBinding) {
				model.CanonicalMutationAllowed = true
			},
			want: Point14ValAStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValAEvidenceBindingModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValAExternalSignalEvidenceBindingState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValAExternalSignalFreshnessAndTimestampBoundaryState(t *testing.T) {
	dependency := point14ValADependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*ExternalSignalFreshnessAndTimestampBoundary)
		want   string
	}{
		{
			name:   "server utc received at passes",
			mutate: func(model *ExternalSignalFreshnessAndTimestampBoundary) {},
			want:   Point14ValAStateActive,
		},
		{
			name: "approved customer controlled time source passes where modeled",
			mutate: func(model *ExternalSignalFreshnessAndTimestampBoundary) {
				model.ReceivedTimeSource = point14Val0TimeSourceApprovedCustomerTime
			},
			want: Point14ValAStateActive,
		},
		{
			name: "client local time as canonical received at blocks",
			mutate: func(model *ExternalSignalFreshnessAndTimestampBoundary) {
				model.ReceivedTimeSource = point14Val0TimeSourceClientLocal
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "source event at future dated review required",
			mutate: func(model *ExternalSignalFreshnessAndTimestampBoundary) {
				model.SourceEventAt = "2026-05-06T00:10:00Z"
			},
			want: Point14ValAStateReviewRequired,
		},
		{
			name: "source event at after received at impossible ordering review required",
			mutate: func(model *ExternalSignalFreshnessAndTimestampBoundary) {
				model.SourceEventAt = "2026-05-05T23:20:00Z"
			},
			want: Point14ValAStateReviewRequired,
		},
		{
			name: "stale signal cannot validate active without governance review path",
			mutate: func(model *ExternalSignalFreshnessAndTimestampBoundary) {
				model.StaleSignal = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "stale signal with governance review path is review required",
			mutate: func(model *ExternalSignalFreshnessAndTimestampBoundary) {
				model.StaleSignal = true
				model.GovernanceReviewPathExists = true
			},
			want: Point14ValAStateReviewRequired,
		},
		{
			name: "timestamp cannot upgrade candidate to authority",
			mutate: func(model *ExternalSignalFreshnessAndTimestampBoundary) {
				model.AuthorityUpgradeRequested = true
			},
			want: Point14ValAStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValAFreshnessAndTimestampBoundaryModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValAExternalSignalFreshnessAndTimestampBoundaryState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValAExternalSignalDuplicateAndRelationGuardState(t *testing.T) {
	dependency := point14ValADependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*ExternalSignalDuplicateAndRelationGuard)
		want   string
	}{
		{
			name: "exact duplicate does not create duplicate active candidate",
			mutate: func(model *ExternalSignalDuplicateAndRelationGuard) {
				model.DuplicateSignalRefs = []string{"signal_point14_vala_duplicate_001"}
			},
			want: Point14ValAStateReviewRequired,
		},
		{
			name: "duplicate with changed payload hash blocks",
			mutate: func(model *ExternalSignalDuplicateAndRelationGuard) {
				model.DuplicateSignalRefs = []string{"signal_point14_vala_duplicate_001"}
				model.ObservedNormalizedPayloadHash = "hash_point14_vala_mutated_payload_001"
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "unrelated signal cannot bind to artifact claim",
			mutate: func(model *ExternalSignalDuplicateAndRelationGuard) {
				model.ArtifactRelationExact = false
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "conflicting duplicate cannot silently replace prior signal",
			mutate: func(model *ExternalSignalDuplicateAndRelationGuard) {
				model.DuplicateSignalRefs = []string{"signal_point14_vala_duplicate_001"}
				model.ConflictingDuplicate = true
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "recomputed normalized hash from mutated object cannot hide drift",
			mutate: func(model *ExternalSignalDuplicateAndRelationGuard) {
				model.ObservedNormalizedPayloadHash = "hash_point14_vala_recomputed_mutated_001"
			},
			want: Point14ValAStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValADuplicateAndRelationGuardModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValAExternalSignalDuplicateAndRelationGuardState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValAExternalSignalTenantBoundaryGuardState(t *testing.T) {
	dependency := point14ValADependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*ExternalSignalTenantBoundaryGuard)
		want   string
	}{
		{
			name:   "tenant scoped boundary passes",
			mutate: func(model *ExternalSignalTenantBoundaryGuard) {},
			want:   Point14ValAStateActive,
		},
		{
			name: "cross tenant source evidence artifact claim mismatch blocks clb0",
			mutate: func(model *ExternalSignalTenantBoundaryGuard) {
				model.EvidenceTenantScope = "tenant_point14_vala_other"
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "global non tenant signal cannot attach tenant evidence unless explicit bounded rule exists",
			mutate: func(model *ExternalSignalTenantBoundaryGuard) {
				model.ScopeClassification = point14Val0ScopeGlobalAdvisory
			},
			want: Point14ValAStateBlocked,
		},
		{
			name: "tenant private data cannot be exposed in normalized public global signal",
			mutate: func(model *ExternalSignalTenantBoundaryGuard) {
				model.TenantPrivateDataExposed = true
			},
			want: Point14ValAStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValATenantBoundaryGuardModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValAExternalSignalTenantBoundaryGuardState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValAAuthorityAndWordingGuards(t *testing.T) {
	t.Run("external authority tokens block", func(t *testing.T) {
		model := point14ValANoExternalAuthorityValidationGuardModel()
		model.ObservedAuthorityMarkers = []string{"external_pass"}
		if got := EvaluatePoint14ValANoExternalAuthorityValidationGuardState(model); got != Point14ValAStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("zero-width external authority marker blocks", func(t *testing.T) {
		model := point14ValANoExternalAuthorityValidationGuardModel()
		model.ObservedAuthorityMarkers = []string{"external_source_of_truth\u200b"}
		if got := EvaluatePoint14ValANoExternalAuthorityValidationGuardState(model); got != Point14ValAStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("point14 pass marker blocks", func(t *testing.T) {
		model := point14ValANoExternalAuthorityValidationGuardModel()
		model.ObservedAuthorityMarkers = []string{point14ValABlockedPassToken}
		if got := EvaluatePoint14ValANoExternalAuthorityValidationGuardState(model); got != Point14ValAStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("ai agent external source authority flags block globally", func(t *testing.T) {
		model := point14ValANoExternalAuthorityValidationGuardModel()
		model.ExternalAuthorityAllowed = true
		if got := EvaluatePoint14ValANoExternalAuthorityValidationGuardState(model); got != Point14ValAStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("forbidden wording blocks", func(t *testing.T) {
		model := point14ValANoOverclaimValidationWordingModel()
		model.ObservedValidationTexts = []string{"production approved"}
		if got := EvaluatePoint14ValANoOverclaimValidationWordingState(model); got != Point14ValAStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("split forbidden wording blocks across observed corpus", func(t *testing.T) {
		model := point14ValANoOverclaimValidationWordingModel()
		model.ObservedValidationTexts = []string{"production", "approved"}
		if got := EvaluatePoint14ValANoOverclaimValidationWordingState(model); got != Point14ValAStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("split forbidden wording across observed categories blocks", func(t *testing.T) {
		model := point14ValANoOverclaimValidationWordingModel()
		model.ObservedNormalizationTexts = []string{"production"}
		model.ObservedValidationTexts = []string{"approved"}
		if got := EvaluatePoint14ValANoOverclaimValidationWordingState(model); got != Point14ValAStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("safe bounded wording passes", func(t *testing.T) {
		model := point14ValANoOverclaimValidationWordingModel()
		if got := EvaluatePoint14ValANoOverclaimValidationWordingState(model); got != Point14ValAStateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
	t.Run("forbidden wording allowed only in internal blocked diagnostic context", func(t *testing.T) {
		model := point14ValANoOverclaimValidationWordingModel()
		model.InternalDiagnosticTexts = []string{"external PASS"}
		model.InternalDiagnosticsClassifiedBlocked = true
		if got := EvaluatePoint14ValANoOverclaimValidationWordingState(model); got != Point14ValAStateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
}

func TestPoint14ValAExternalSignalValidationResultState(t *testing.T) {
	t.Run("validated candidate is active when all component states active", func(t *testing.T) {
		model := point14ValAValidationResultModel()
		model.NormalizedSignalState = Point14ValAStateActive
		model.SourceIdentityState = Point14ValAStateActive
		model.ScopeBindingState = Point14ValAStateActive
		model.EvidenceBindingState = Point14ValAStateActive
		model.FreshnessTimestampState = Point14ValAStateActive
		model.DuplicateRelationGuardState = Point14ValAStateActive
		model.TenantBoundaryGuardState = Point14ValAStateActive
		model.NoExternalAuthorityState = Point14ValAStateActive
		model.NoOverclaimState = Point14ValAStateActive
		if got := EvaluatePoint14ValAExternalSignalValidationResultState(model); got != Point14ValAStateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
	t.Run("blocked candidate state blocks", func(t *testing.T) {
		model := point14ValAValidationResultModel()
		model.NormalizedSignalState = Point14ValAStateActive
		model.SourceIdentityState = Point14ValAStateActive
		model.ScopeBindingState = Point14ValAStateActive
		model.EvidenceBindingState = Point14ValAStateActive
		model.FreshnessTimestampState = Point14ValAStateActive
		model.DuplicateRelationGuardState = Point14ValAStateActive
		model.TenantBoundaryGuardState = Point14ValAStateActive
		model.NoExternalAuthorityState = Point14ValAStateActive
		model.NoOverclaimState = Point14ValAStateActive
		model.ValidationState = point14ValAValidationCandidateDuplicate
		if got := EvaluatePoint14ValAExternalSignalValidationResultState(model); got != Point14ValAStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
}

func TestPoint14ValAFoundationAggregation(t *testing.T) {
	t.Run("any blocked component yields blocked", func(t *testing.T) {
		model := Point14ValAFoundationModel()
		model.NormalizedExternalSignal.PassAllowed = true
		got := ComputePoint14ValAFoundation(model)
		if got.CurrentState != Point14ValAStateBlocked {
			t.Fatalf("expected blocked, got %#v", got)
		}
	})
	t.Run("any review required and no blocked yields review required", func(t *testing.T) {
		model := Point14ValAFoundationModel()
		model.FreshnessAndTimestamp.StaleSignal = true
		model.FreshnessAndTimestamp.GovernanceReviewPathExists = true
		got := ComputePoint14ValAFoundation(model)
		if got.CurrentState != Point14ValAStateReviewRequired {
			t.Fatalf("expected review required, got %#v", got)
		}
	})
	t.Run("incomplete only when no blocked review required exists", func(t *testing.T) {
		model := Point14ValAFoundationModel()
		model.ValidationResult.ValidationState = point14ValAValidationCandidateIncomplete
		got := ComputePoint14ValAFoundation(model)
		if got.CurrentState != Point14ValAStateIncomplete {
			t.Fatalf("expected incomplete, got %#v", got)
		}
	})
	t.Run("unknown validation state blocks instead of upgrading to validated", func(t *testing.T) {
		model := Point14ValAFoundationModel()
		model.ValidationResult.ValidationState = "candidate_unknown"
		got := ComputePoint14ValAFoundation(model)
		if got.CurrentState != Point14ValAStateBlocked ||
			got.ValidationResultState != Point14ValAStateBlocked ||
			got.ValidationResult.ValidationState != point14ValAValidationCandidateBlocked {
			t.Fatalf("expected blocked unknown validation state, got %#v", got)
		}
	})
	t.Run("active only when all components active", func(t *testing.T) {
		model := ComputePoint14ValAFoundation(Point14ValAFoundationModel())
		if model.CurrentState != Point14ValAStateActive {
			t.Fatalf("expected active, got %#v", model)
		}
	})
}

package formal

import "testing"

func TestPoint14ValBDependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValBDependencySnapshot)
		want   string
	}{
		{
			name:   "canonical inherited boundary snapshot stays active",
			mutate: func(model *Point14ValBDependencySnapshot) {},
			want:   Point14ValBStateActive,
		},
		{
			name: "missing point14 vala blocks",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValACurrentState = ""
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 current state in vala dependency blocks",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValA.Dependency.InheritedPoint10CurrentState += " "
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 projection state in vala dependency blocks",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValA.Dependency.InheritedPoint10ProjectionState += " "
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 no overclaim state in vala dependency blocks",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValA.Dependency.InheritedPoint10NoOverclaimState += " "
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "whitespace retagged nested point10 pass rule state in vala dependency blocks",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValA.Dependency.InheritedPoint10PassRuleState += " "
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "tab newline retagged embedded vala current state blocks",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValA.CurrentState = "\t" + model.Point14ValA.CurrentState + "\n"
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "whitespace retagged point13 pass token from embedded point13 blocks",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point13ValE.Point13PassToken += " "
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "point14 vala blocked blocks",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValACurrentState = Point14ValAStateBlocked
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "point14 vala review required prevents active",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValACurrentState = Point14ValAStateReviewRequired
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "point14 vala incomplete prevents active",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValACurrentState = Point14ValAStateIncomplete
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "premature point14 pass blocks",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14PassSeen = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "local valb readiness cannot override missing vala closure",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValACurrentState = ""
				model.InheritedPoint14Val0CurrentState = Point14Val0StateActive
				model.InheritedPoint13ValEPassAllowed = true
				model.InheritedPoint13ValEPassToken = point13ValEPoint13PassToken
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "embedded point14 vala mismatch blocks",
			mutate: func(model *Point14ValBDependencySnapshot) {
				model.Point14ValA.CurrentState = Point14ValAStateBlocked
			},
			want: Point14ValBStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValBDependencySnapshotModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValBDependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValBExternalSignalConflictSetState(t *testing.T) {
	dependency := point14ValBDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*ExternalSignalConflictSet)
		want   string
	}{
		{
			name:   "valid no conflict set active",
			mutate: func(model *ExternalSignalConflictSet) {},
			want:   Point14ValBStateActive,
		},
		{
			name: "scanner vs vex conflict returns review required",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "scanner_vs_vex"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "scanner vs maintainer conflict returns review required",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "scanner_vs_maintainer"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "vex vs canonical evidence conflict returns review required",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "vex_vs_canonical_evidence"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "maintainer vs reviewer conflict returns review required",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "maintainer_vs_reviewer"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "auditor note vs insufficient evidence cannot upgrade to active pass",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "auditor_note_vs_insufficient_evidence"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "verifier note vs canonical decision cannot override canonical decision",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "verifier_note_vs_canonical_decision"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "duplicate external signal conflict review required",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "duplicate_external_signal_conflict"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "stale external signal conflict review required",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "stale_external_signal_conflict"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "cross tenant signal conflict blocks clb0",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "cross_tenant_signal_conflict"
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "public claim vs private evidence review required",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "public_claim_vs_private_evidence"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "agent recommendation vs governance review required",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "agent_recommendation_vs_governance"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "unknown conflict type blocks",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ConflictDetected = true
				model.ConflictType = "mystery_conflict"
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "conflict with resolve to pass blocks",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ResolveToPass = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "conflict with resolve to canonical truth blocks",
			mutate: func(model *ExternalSignalConflictSet) {
				model.ResolveToCanonicalTruth = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "conflict with publish correction blocks",
			mutate: func(model *ExternalSignalConflictSet) {
				model.PublishCorrection = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "conflict with revoke claim blocks",
			mutate: func(model *ExternalSignalConflictSet) {
				model.RevokeClaim = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "conflict with public badge blocks",
			mutate: func(model *ExternalSignalConflictSet) {
				model.CreatePublicBadge = true
			},
			want: Point14ValBStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValBConflictSetModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValBExternalSignalConflictSetState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValBStakeholderSignalComparisonState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*StakeholderSignalComparison)
		want   string
	}{
		{
			name:   "consistent bounded stakeholder comparison active",
			mutate: func(model *StakeholderSignalComparison) {},
			want:   Point14ValBStateActive,
		},
		{
			name: "conflicting stakeholder comparison review required",
			mutate: func(model *StakeholderSignalComparison) {
				model.ComparisonResult = "conflicting"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "unsupported role blocks",
			mutate: func(model *StakeholderSignalComparison) {
				model.StakeholderRoles = []string{"mystery_role"}
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "scanner cannot emit pass",
			mutate: func(model *StakeholderSignalComparison) {
				model.StakeholderRoles = []string{"scanner"}
				model.EmitsPass = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "vex issuer cannot resolve canonical conflict alone",
			mutate: func(model *StakeholderSignalComparison) {
				model.StakeholderRoles = []string{"vex_issuer"}
				model.ResolvesConflict = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "auditor verifier cannot approve production",
			mutate: func(model *StakeholderSignalComparison) {
				model.StakeholderRoles = []string{"auditor", "verifier"}
				model.ApprovesProduction = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "partner customer admin cannot certify production readiness",
			mutate: func(model *StakeholderSignalComparison) {
				model.StakeholderRoles = []string{"partner", "customer_admin"}
				model.CertifiesCompliance = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "crowd public consensus cannot resolve conflict",
			mutate: func(model *StakeholderSignalComparison) {
				model.CrowdConsensusResolutionRequested = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing authority scope blocks",
			mutate: func(model *StakeholderSignalComparison) {
				model.AuthorityScopeRefs = nil
			},
			want: Point14ValBStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValBStakeholderComparisonModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValBStakeholderSignalComparisonState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValBDisputeIntakePacketState(t *testing.T) {
	dependency := point14ValBDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*DisputeIntakePacket)
		want   string
	}{
		{
			name:   "valid opened dispute records refs and remains bounded",
			mutate: func(model *DisputeIntakePacket) { model.LifecycleState = point14Val0DisputeOpened },
			want:   Point14ValBStateActive,
		},
		{
			name: "missing dispute id blocks",
			mutate: func(model *DisputeIntakePacket) {
				model.DisputeID = ""
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing opened by role blocks",
			mutate: func(model *DisputeIntakePacket) {
				model.OpenedByRole = ""
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing opened at blocks",
			mutate: func(model *DisputeIntakePacket) {
				model.OpenedAt = ""
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "client local time as canonical opened at blocks",
			mutate: func(model *DisputeIntakePacket) {
				model.OpenedTimeSource = point14Val0TimeSourceClientLocal
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing affected signal refs blocks",
			mutate: func(model *DisputeIntakePacket) {
				model.AffectedSignalRefs = nil
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing artifact refs according to scope blocks",
			mutate: func(model *DisputeIntakePacket) {
				model.AffectedArtifactRefs = nil
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing claim refs according to scope blocks",
			mutate: func(model *DisputeIntakePacket) {
				model.AffectedClaimRefs = nil
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing evidence refs according to scope blocks",
			mutate: func(model *DisputeIntakePacket) {
				model.AffectedEvidenceRefs = nil
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "corrected lifecycle state blocks in valb",
			mutate: func(model *DisputeIntakePacket) {
				model.LifecycleState = point14Val0DisputeCorrected
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "revoked lifecycle state blocks in valb",
			mutate: func(model *DisputeIntakePacket) {
				model.LifecycleState = point14Val0DisputeRevoked
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "published notice lifecycle state blocks in valb",
			mutate: func(model *DisputeIntakePacket) {
				model.LifecycleState = point14Val0DisputePublished
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "dispute cannot mutate canonical evidence",
			mutate: func(model *DisputeIntakePacket) {
				model.CanonicalMutationRequested = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "future dated source event blocks review required",
			mutate: func(model *DisputeIntakePacket) {
				model.SourceEventAt = "2026-05-06T08:00:00Z"
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "backdated dispute closure attempt blocks",
			mutate: func(model *DisputeIntakePacket) {
				model.ClosureAttempted = true
				model.ClosedAt = "2026-05-06T07:00:00Z"
			},
			want: Point14ValBStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValBDisputeIntakeModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValBDisputeIntakePacketState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValBDisputeEvidenceRequirementGateState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*DisputeEvidenceRequirementGate)
		want   string
	}{
		{
			name: "malformed evidence refs block even when evidence types look valid",
			mutate: func(model *DisputeEvidenceRequirementGate) {
				model.RequiredEvidenceRefs = []string{"bad_ref"}
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing decisive evidence returns incomplete",
			mutate: func(model *DisputeEvidenceRequirementGate) {
				model.DecisiveEvidenceMissing = true
			},
			want: Point14ValBStateIncomplete,
		},
		{
			name: "stale evidence review required",
			mutate: func(model *DisputeEvidenceRequirementGate) {
				model.EvidenceState = point14ValAEvidenceStateStale
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "revoked evidence blocks",
			mutate: func(model *DisputeEvidenceRequirementGate) {
				model.EvidenceState = point14ValAEvidenceStateRevoked
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "expired evidence blocks",
			mutate: func(model *DisputeEvidenceRequirementGate) {
				model.EvidenceState = point14ValAEvidenceStateExpired
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "agent recommendation alone cannot satisfy evidence requirement",
			mutate: func(model *DisputeEvidenceRequirementGate) {
				model.AgentRecommendationOnly = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "crowd consensus alone cannot satisfy evidence requirement",
			mutate: func(model *DisputeEvidenceRequirementGate) {
				model.CrowdConsensusOnly = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "auditor note alone cannot satisfy evidence requirement",
			mutate: func(model *DisputeEvidenceRequirementGate) {
				model.AuditorNoteOnly = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "evidence requirement cannot be bypassed by conflict triage state",
			mutate: func(model *DisputeEvidenceRequirementGate) {
				model.ConflictTriageBypassRequested = true
			},
			want: Point14ValBStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValBEvidenceRequirementGateModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValBDisputeEvidenceRequirementGateState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValBGovernanceEscalationPathState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*GovernanceEscalationPath)
		want   string
	}{
		{
			name:   "not required governance path active",
			mutate: func(model *GovernanceEscalationPath) {},
			want:   Point14ValBStateActive,
		},
		{
			name: "governance path with owner approver reason audit active as escalation only",
			mutate: func(model *GovernanceEscalationPath) {
				model.EscalationRequired = true
				model.GovernanceEventRef = "governance_event_point14_valb_001"
				model.Owner = "owner_point14_valb_001"
				model.ApproverRole = "security_reviewer"
				model.Reason = "conflict requires governance review"
				model.AuditRef = "audit_event_point14_valb_001"
				model.EscalationState = point14ValBEscalationStateCompleted
			},
			want: Point14ValBStateActive,
		},
		{
			name: "unresolved conflict without governance path blocks",
			mutate: func(model *GovernanceEscalationPath) {
				model.EscalationRequired = true
				model.EscalationState = point14ValBEscalationStateQueued
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing owner blocks",
			mutate: func(model *GovernanceEscalationPath) {
				model.EscalationRequired = true
				model.GovernanceEventRef = "governance_event_point14_valb_001"
				model.ApproverRole = "security_reviewer"
				model.Reason = "conflict requires governance review"
				model.AuditRef = "audit_event_point14_valb_001"
				model.EscalationState = point14ValBEscalationStateCompleted
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing approver role blocks",
			mutate: func(model *GovernanceEscalationPath) {
				model.EscalationRequired = true
				model.GovernanceEventRef = "governance_event_point14_valb_001"
				model.Owner = "owner_point14_valb_001"
				model.Reason = "conflict requires governance review"
				model.AuditRef = "audit_event_point14_valb_001"
				model.EscalationState = point14ValBEscalationStateCompleted
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "missing audit ref blocks",
			mutate: func(model *GovernanceEscalationPath) {
				model.EscalationRequired = true
				model.GovernanceEventRef = "governance_event_point14_valb_001"
				model.Owner = "owner_point14_valb_001"
				model.ApproverRole = "security_reviewer"
				model.Reason = "conflict requires governance review"
				model.EscalationState = point14ValBEscalationStateCompleted
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "escalation cannot approve production",
			mutate: func(model *GovernanceEscalationPath) {
				model.ApprovesProduction = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "escalation cannot publish correction in valb",
			mutate: func(model *GovernanceEscalationPath) {
				model.PublishesCorrection = true
			},
			want: Point14ValBStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValBGovernanceEscalationPathModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValBGovernanceEscalationPathState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValBTenantPrivacyConflictBoundaryState(t *testing.T) {
	dependency := point14ValBDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*TenantPrivacyConflictBoundary)
		want   string
	}{
		{
			name:   "tenant private bounded conflict active",
			mutate: func(model *TenantPrivacyConflictBoundary) {},
			want:   Point14ValBStateActive,
		},
		{
			name: "cross tenant conflict blocks clb0",
			mutate: func(model *TenantPrivacyConflictBoundary) {
				model.TenantScope = "tenant_point14_valb_other"
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "tenant private data exposed blocks clb0",
			mutate: func(model *TenantPrivacyConflictBoundary) {
				model.TenantPrivateDataExposed = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "public visibility requested without boundary blocks",
			mutate: func(model *TenantPrivacyConflictBoundary) {
				model.PublicVisibilityRequested = true
				model.VisibilityScope = point14Val0VisibilityScopedCustomer
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "bounded public visibility request remains review required",
			mutate: func(model *TenantPrivacyConflictBoundary) {
				model.PublicVisibilityRequested = true
				model.VisibilityScope = point14Val0VisibilityPublicNoticeLimited
			},
			want: Point14ValBStateReviewRequired,
		},
		{
			name: "redacted limited conflict must keep limitation visible",
			mutate: func(model *TenantPrivacyConflictBoundary) {
				model.LimitationVisible = false
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "conflict wording cannot strengthen claim",
			mutate: func(model *TenantPrivacyConflictBoundary) {
				model.StrengthensClaim = true
			},
			want: Point14ValBStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValBTenantPrivacyConflictBoundaryModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValBTenantPrivacyConflictBoundaryState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValBAgentDisputeRecommendationBoundaryState(t *testing.T) {
	dependency := point14ValBDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*AgentDisputeRecommendationBoundary)
		want   string
	}{
		{
			name:   "agent recommendation may be advisory input",
			mutate: func(model *AgentDisputeRecommendationBoundary) {},
			want:   Point14ValBStateActive,
		},
		{
			name: "agent recommendation cannot resolve conflict",
			mutate: func(model *AgentDisputeRecommendationBoundary) {
				model.CanResolveConflict = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "agent recommendation cannot satisfy evidence requirement alone",
			mutate: func(model *AgentDisputeRecommendationBoundary) {
				model.CanSatisfyEvidenceRequirementAlone = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "agent recommendation cannot publish correction",
			mutate: func(model *AgentDisputeRecommendationBoundary) {
				model.CanPublishCorrection = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "agent recommendation cannot revoke claim",
			mutate: func(model *AgentDisputeRecommendationBoundary) {
				model.CanRevokeClaim = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "agent recommendation cannot override governance",
			mutate: func(model *AgentDisputeRecommendationBoundary) {
				model.CanOverrideGovernance = true
			},
			want: Point14ValBStateBlocked,
		},
		{
			name: "ai agent authority flags block globally",
			mutate: func(model *AgentDisputeRecommendationBoundary) {
				model.ExternalAuthorityAllowed = true
			},
			want: Point14ValBStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValBAgentDisputeRecommendationBoundaryModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValBAgentDisputeRecommendationBoundaryState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValBAuthorityAndWordingGuards(t *testing.T) {
	t.Run("external authority tokens block", func(t *testing.T) {
		model := point14ValBNoExternalAuthorityConflictGuardModel()
		model.ObservedAuthorityMarkers = []string{"external_pass"}
		if got := EvaluatePoint14ValBNoExternalAuthorityConflictGuardState(model); got != Point14ValBStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("scanner pass blocks", func(t *testing.T) {
		model := point14ValBNoExternalAuthorityConflictGuardModel()
		model.ObservedAuthorityMarkers = []string{"scanner_pass"}
		if got := EvaluatePoint14ValBNoExternalAuthorityConflictGuardState(model); got != Point14ValBStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("dispute auto resolved blocks", func(t *testing.T) {
		model := point14ValBNoExternalAuthorityConflictGuardModel()
		model.DisputeAutoResolved = true
		if got := EvaluatePoint14ValBNoExternalAuthorityConflictGuardState(model); got != Point14ValBStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("forbidden dispute wording blocks", func(t *testing.T) {
		model := point14ValBNoOverclaimDisputeWordingModel()
		model.ObservedDisputeTexts = []string{"dispute resolved by AI"}
		if got := EvaluatePoint14ValBNoOverclaimDisputeWordingState(model); got != Point14ValBStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
	t.Run("safe bounded wording passes", func(t *testing.T) {
		model := point14ValBNoOverclaimDisputeWordingModel()
		if got := EvaluatePoint14ValBNoOverclaimDisputeWordingState(model); got != Point14ValBStateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
	t.Run("forbidden wording allowed only in internal blocked diagnostic context", func(t *testing.T) {
		model := point14ValBNoOverclaimDisputeWordingModel()
		model.InternalDiagnosticTexts = []string{"scanner PASS"}
		model.InternalDiagnosticsClassifiedBlocked = true
		if got := EvaluatePoint14ValBNoOverclaimDisputeWordingState(model); got != Point14ValBStateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
	t.Run("valb specific forbidden internal diagnostic wording blocks when not classified blocked", func(t *testing.T) {
		model := point14ValBNoOverclaimDisputeWordingModel()
		model.InternalDiagnosticTexts = []string{"dispute resolved by AI"}
		if got := EvaluatePoint14ValBNoOverclaimDisputeWordingState(model); got != Point14ValBStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
}

func TestPoint14ValBExternalConflictTriageResultState(t *testing.T) {
	t.Run("no conflict detected active only when no conflict exists and all evidence valid", func(t *testing.T) {
		model := point14ValBDisputeTriageResultModel()
		model.ConflictSetState = Point14ValBStateActive
		model.StakeholderComparisonState = Point14ValBStateActive
		model.DisputeIntakeState = Point14ValBStateActive
		model.EvidenceRequirementState = Point14ValBStateActive
		model.GovernanceEscalationState = Point14ValBStateActive
		model.TenantPrivacyBoundaryState = Point14ValBStateActive
		model.AgentBoundaryState = Point14ValBStateActive
		model.NoExternalAuthorityState = Point14ValBStateActive
		model.NoOverclaimState = Point14ValBStateActive
		if got := EvaluatePoint14ValBExternalConflictTriageResultState(model); got != Point14ValBStateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
	t.Run("conflict review required prevents active", func(t *testing.T) {
		model := point14ValBDisputeTriageResultModel()
		model.ConflictSetState = Point14ValBStateReviewRequired
		model.StakeholderComparisonState = Point14ValBStateActive
		model.DisputeIntakeState = Point14ValBStateActive
		model.EvidenceRequirementState = Point14ValBStateActive
		model.GovernanceEscalationState = Point14ValBStateActive
		model.TenantPrivacyBoundaryState = Point14ValBStateActive
		model.AgentBoundaryState = Point14ValBStateActive
		model.NoExternalAuthorityState = Point14ValBStateActive
		model.NoOverclaimState = Point14ValBStateActive
		model.TriageState = point14ValBTriageReviewRequired
		if got := EvaluatePoint14ValBExternalConflictTriageResultState(model); got != Point14ValBStateReviewRequired {
			t.Fatalf("expected review required, got %s", got)
		}
	})
	t.Run("conflict evidence required prevents active", func(t *testing.T) {
		model := point14ValBDisputeTriageResultModel()
		model.ConflictSetState = Point14ValBStateActive
		model.StakeholderComparisonState = Point14ValBStateActive
		model.DisputeIntakeState = Point14ValBStateActive
		model.EvidenceRequirementState = Point14ValBStateIncomplete
		model.GovernanceEscalationState = Point14ValBStateActive
		model.TenantPrivacyBoundaryState = Point14ValBStateActive
		model.AgentBoundaryState = Point14ValBStateActive
		model.NoExternalAuthorityState = Point14ValBStateActive
		model.NoOverclaimState = Point14ValBStateActive
		model.TriageState = point14ValBTriageEvidenceRequired
		if got := EvaluatePoint14ValBExternalConflictTriageResultState(model); got != Point14ValBStateIncomplete {
			t.Fatalf("expected incomplete, got %s", got)
		}
	})
	t.Run("conflict blocked blocks", func(t *testing.T) {
		model := point14ValBDisputeTriageResultModel()
		model.ConflictSetState = Point14ValBStateBlocked
		model.StakeholderComparisonState = Point14ValBStateActive
		model.DisputeIntakeState = Point14ValBStateActive
		model.EvidenceRequirementState = Point14ValBStateActive
		model.GovernanceEscalationState = Point14ValBStateActive
		model.TenantPrivacyBoundaryState = Point14ValBStateActive
		model.AgentBoundaryState = Point14ValBStateActive
		model.NoExternalAuthorityState = Point14ValBStateActive
		model.NoOverclaimState = Point14ValBStateActive
		if got := EvaluatePoint14ValBExternalConflictTriageResultState(model); got != Point14ValBStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})
}

func TestPoint14ValBFoundationAggregation(t *testing.T) {
	t.Run("any blocked component yields blocked", func(t *testing.T) {
		model := Point14ValBFoundationModel()
		model.ConflictSet.ResolveToPass = true
		got := ComputePoint14ValBFoundation(model)
		if got.CurrentState != Point14ValBStateBlocked {
			t.Fatalf("expected blocked, got %#v", got)
		}
	})
	t.Run("any review required and no blocked yields review required", func(t *testing.T) {
		model := Point14ValBFoundationModel()
		model.ConflictSet.ConflictDetected = true
		model.ConflictSet.ConflictType = "scanner_vs_vex"
		got := ComputePoint14ValBFoundation(model)
		if got.CurrentState != Point14ValBStateReviewRequired {
			t.Fatalf("expected review required, got %#v", got)
		}
	})
	t.Run("incomplete only when no blocked review required exists", func(t *testing.T) {
		model := Point14ValBFoundationModel()
		model.EvidenceRequirementGate.DecisiveEvidenceMissing = true
		model.DisputeTriageResult.TriageState = point14ValBTriageEvidenceRequired
		got := ComputePoint14ValBFoundation(model)
		if got.CurrentState != Point14ValBStateIncomplete {
			t.Fatalf("expected incomplete, got %#v", got)
		}
	})
	t.Run("unknown triage state blocks instead of upgrading to active", func(t *testing.T) {
		model := Point14ValBFoundationModel()
		model.DisputeTriageResult.TriageState = "conflict_unknown"
		got := ComputePoint14ValBFoundation(model)
		if got.CurrentState != Point14ValBStateBlocked ||
			got.DisputeTriageResultState != Point14ValBStateBlocked ||
			got.DisputeTriageResult.TriageState != point14ValBTriageBlocked {
			t.Fatalf("expected blocked unknown triage state, got %#v", got)
		}
	})
	t.Run("active only when all components active", func(t *testing.T) {
		model := ComputePoint14ValBFoundation(Point14ValBFoundationModel())
		if model.CurrentState != Point14ValBStateActive {
			t.Fatalf("expected active, got %#v", model)
		}
	})
}

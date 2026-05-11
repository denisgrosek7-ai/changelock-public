package formal

import (
	"encoding/json"
	"testing"
)

func clonePoint14ValDDependencySnapshot(t *testing.T, model Point14ValDDependencySnapshot) Point14ValDDependencySnapshot {
	t.Helper()
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal point14 vald dependency snapshot: %v", err)
	}
	var cloned Point14ValDDependencySnapshot
	if err := json.Unmarshal(payload, &cloned); err != nil {
		t.Fatalf("unmarshal point14 vald dependency snapshot: %v", err)
	}
	return cloned
}

func TestPoint14ValDDependencyState(t *testing.T) {
	base := point14ValDDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*Point14ValDDependencySnapshot)
		want   string
	}{
		{
			name: "missing point14 valc blocks",
			mutate: func(model *Point14ValDDependencySnapshot) {
				model.Point14ValCCurrentState = ""
			},
			want: Point14ValDStateBlocked,
		},
		{
			name: "point14 valc blocked blocks",
			mutate: func(model *Point14ValDDependencySnapshot) {
				model.Point14ValCCurrentState = Point14ValCStateBlocked
			},
			want: Point14ValDStateBlocked,
		},
		{
			name: "point14 valc review required prevents active",
			mutate: func(model *Point14ValDDependencySnapshot) {
				model.Point14ValCCurrentState = Point14ValCStateReviewRequired
			},
			want: Point14ValDStateBlocked,
		},
		{
			name: "point14 valc incomplete prevents active",
			mutate: func(model *Point14ValDDependencySnapshot) {
				model.Point14ValCCurrentState = Point14ValCStateIncomplete
			},
			want: Point14ValDStateBlocked,
		},
		{
			name: "premature point14 pass blocks",
			mutate: func(model *Point14ValDDependencySnapshot) {
				model.Point14PassSeen = true
			},
			want: Point14ValDStateBlocked,
		},
		{
			name: "local vald readiness cannot override missing valc closure",
			mutate: func(model *Point14ValDDependencySnapshot) {
				model.Point14ValCCurrentState = ""
				model.InheritedPoint14ValBCurrentState = Point14ValBStateActive
				model.InheritedPoint13ValEPassAllowed = true
				model.InheritedPoint13ValEPassToken = point13ValEPoint13PassToken
			},
			want: Point14ValDStateBlocked,
		},
		{
			name: "embedded point14 valc mismatch blocks",
			mutate: func(model *Point14ValDDependencySnapshot) {
				model.Point14ValC.CurrentState = Point14ValCStateBlocked
			},
			want: Point14ValDStateBlocked,
		},
		{
			name: "point11 final pass gate blocked blocks",
			mutate: func(model *Point14ValDDependencySnapshot) {
				model.InheritedPoint11FinalPassGateState = Point11ValDFinalPassGateStateBlocked
				model.Point14ValC.Dependency.InheritedPoint11FinalPassGateState = Point11ValDFinalPassGateStateBlocked
			},
			want: Point14ValDStateBlocked,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := clonePoint14ValDDependencySnapshot(t, base)
			tc.mutate(&model)
			if got := EvaluatePoint14ValDDependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDTimelineProjectionState(t *testing.T) {
	base := point14ValDDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*ExternalEcosystemTimelineProjection)
		want   string
	}{
		{"valid read only ecosystem timeline active", func(model *ExternalEcosystemTimelineProjection) {}, Point14ValDStateActive},
		{"missing timeline projection id blocks", func(model *ExternalEcosystemTimelineProjection) { model.TimelineProjectionID = "" }, Point14ValDStateBlocked},
		{"missing tenant scope global classification blocks", func(model *ExternalEcosystemTimelineProjection) {
			model.TenantScope = ""
			model.GlobalScopeClassification = ""
		}, Point14ValDStateBlocked},
		{"missing source projection refs blocks", func(model *ExternalEcosystemTimelineProjection) { model.SourceProjectionRefs = nil }, Point14ValDStateBlocked},
		{"missing governance trace refs blocks", func(model *ExternalEcosystemTimelineProjection) { model.GovernanceTraceRefs = nil }, Point14ValDStateBlocked},
		{"missing generated at blocks", func(model *ExternalEcosystemTimelineProjection) { model.GeneratedAt = "" }, Point14ValDStateBlocked},
		{"client local generated time source blocks", func(model *ExternalEcosystemTimelineProjection) {
			model.GeneratedTimeSource = point14Val0TimeSourceClientLocal
		}, Point14ValDStateBlocked},
		{"read only false blocks", func(model *ExternalEcosystemTimelineProjection) { model.ReadOnly = false }, Point14ValDStateBlocked},
		{"projection only false blocks", func(model *ExternalEcosystemTimelineProjection) { model.ProjectionOnly = false }, Point14ValDStateBlocked},
		{"timeline mutation flag blocks", func(model *ExternalEcosystemTimelineProjection) { model.MutatesCanonicalEvidence = true }, Point14ValDStateBlocked},
		{"timeline dispute resolution flag blocks", func(model *ExternalEcosystemTimelineProjection) { model.ResolvesDispute = true }, Point14ValDStateBlocked},
		{"timeline correction publication flag blocks", func(model *ExternalEcosystemTimelineProjection) { model.PublishesCorrection = true }, Point14ValDStateBlocked},
		{"timeline revocation execution flag blocks", func(model *ExternalEcosystemTimelineProjection) { model.RevokesClaim = true }, Point14ValDStateBlocked},
		{"timeline public badge flag blocks", func(model *ExternalEcosystemTimelineProjection) { model.CreatesPublicBadge = true }, Point14ValDStateBlocked},
		{"timeline pass flag blocks", func(model *ExternalEcosystemTimelineProjection) { model.EmitsPass = true }, Point14ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValDTimelineProjectionModel(clonePoint14ValDDependencySnapshot(t, base))
			tc.mutate(&model)
			if got := EvaluatePoint14ValDTimelineProjectionState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDSignalTimelineEntryState(t *testing.T) {
	base := point14ValDDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*ExternalSignalTimelineEntryProjection)
		want   string
	}{
		{"valid signal timeline entry active", func(model *ExternalSignalTimelineEntryProjection) {}, Point14ValDStateActive},
		{"missing timeline entry id blocks", func(model *ExternalSignalTimelineEntryProjection) { model.TimelineEntryID = "" }, Point14ValDStateBlocked},
		{"missing normalized signal ref blocks", func(model *ExternalSignalTimelineEntryProjection) { model.NormalizedSignalRef = "" }, Point14ValDStateBlocked},
		{"missing validation result ref blocks", func(model *ExternalSignalTimelineEntryProjection) { model.ValidationResultRef = "" }, Point14ValDStateBlocked},
		{"missing source identity ref blocks", func(model *ExternalSignalTimelineEntryProjection) { model.SourceIdentityRef = "" }, Point14ValDStateBlocked},
		{"missing event at blocks", func(model *ExternalSignalTimelineEntryProjection) { model.EventAt = "" }, Point14ValDStateBlocked},
		{"source event at as canonical authority blocks", func(model *ExternalSignalTimelineEntryProjection) { model.SourceEventAsCanonicalAuthority = true }, Point14ValDStateBlocked},
		{"advisory only false blocks", func(model *ExternalSignalTimelineEntryProjection) { model.AdvisoryOnly = false }, Point14ValDStateBlocked},
		{"authority granted true blocks", func(model *ExternalSignalTimelineEntryProjection) { model.AuthorityGranted = true }, Point14ValDStateBlocked},
		{"signal validity upgrade flag blocks", func(model *ExternalSignalTimelineEntryProjection) { model.UpgradesSignalValidity = true }, Point14ValDStateBlocked},
		{"event after received review required", func(model *ExternalSignalTimelineEntryProjection) { model.EventAt = "2026-05-06T09:02:00Z" }, Point14ValDStateReviewRequired},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValDSignalTimelineEntryModel(clonePoint14ValDDependencySnapshot(t, base))
			tc.mutate(&model)
			if got := EvaluatePoint14ValDSignalTimelineEntryState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDDisputeTimelineProjectionState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*ExternalDisputeTimelineProjection)
		want   string
	}{
		{"valid dispute timeline projection active", func(model *ExternalDisputeTimelineProjection) {}, Point14ValDStateActive},
		{"missing dispute ref blocks", func(model *ExternalDisputeTimelineProjection) { model.DisputeRef = "" }, Point14ValDStateBlocked},
		{"missing conflict set ref blocks", func(model *ExternalDisputeTimelineProjection) { model.ConflictSetRef = "" }, Point14ValDStateBlocked},
		{"missing triage result ref blocks", func(model *ExternalDisputeTimelineProjection) { model.TriageResultRef = "" }, Point14ValDStateBlocked},
		{"missing lifecycle state blocks", func(model *ExternalDisputeTimelineProjection) { model.LifecycleState = "" }, Point14ValDStateBlocked},
		{"evidence required state remains visible", func(model *ExternalDisputeTimelineProjection) {
			model.LifecycleState = point14Val0DisputeEvidenceNeeded
			model.TimelineState = point14ValDDisputeTimelineVisible
		}, Point14ValDStateIncomplete},
		{"dispute timeline cannot resolve dispute", func(model *ExternalDisputeTimelineProjection) { model.ResolvesDispute = true }, Point14ValDStateBlocked},
		{"dispute timeline cannot move lifecycle to corrected", func(model *ExternalDisputeTimelineProjection) { model.MovesLifecycleToCorrected = true }, Point14ValDStateBlocked},
		{"dispute timeline cannot move lifecycle to revoked", func(model *ExternalDisputeTimelineProjection) { model.MovesLifecycleToRevoked = true }, Point14ValDStateBlocked},
		{"dispute timeline cannot move lifecycle to published notice", func(model *ExternalDisputeTimelineProjection) { model.MovesLifecycleToPublishedNotice = true }, Point14ValDStateBlocked},
		{"dispute timeline cannot convert review required incomplete to active", func(model *ExternalDisputeTimelineProjection) { model.ConvertsReviewIncompleteToActive = true }, Point14ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValDDisputeTimelineProjectionModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValDDisputeTimelineProjectionState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDCorrectionReadProjectionState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*CorrectionRevocationPublicationReadProjection)
		want   string
	}{
		{"valid correction read projection active", func(model *CorrectionRevocationPublicationReadProjection) {
			model.RevocationRequestRefs = nil
			model.SupersessionRecordRefs = nil
			model.PublicationApprovalRefs = nil
		}, Point14ValDStateActive},
		{"valid revocation request read projection active", func(model *CorrectionRevocationPublicationReadProjection) {
			model.CorrectionNoticeRefs = nil
			model.SupersessionRecordRefs = nil
			model.PublicationApprovalRefs = nil
		}, Point14ValDStateActive},
		{"valid supersession read projection active", func(model *CorrectionRevocationPublicationReadProjection) {
			model.CorrectionNoticeRefs = nil
			model.RevocationRequestRefs = nil
			model.PublicationApprovalRefs = nil
		}, Point14ValDStateActive},
		{"valid publication approval read projection active", func(model *CorrectionRevocationPublicationReadProjection) {
			model.CorrectionNoticeRefs = nil
			model.RevocationRequestRefs = nil
			model.SupersessionRecordRefs = nil
		}, Point14ValDStateActive},
		{"read only false blocks", func(model *CorrectionRevocationPublicationReadProjection) { model.ReadOnly = false }, Point14ValDStateBlocked},
		{"missing visibility boundary blocks", func(model *CorrectionRevocationPublicationReadProjection) { model.VisibilityBoundaryRefs = nil }, Point14ValDStateBlocked},
		{"missing limitation refs where limited blocks", func(model *CorrectionRevocationPublicationReadProjection) {
			model.LimitationRefs = nil
			model.RedactionRefs = []string{"redaction_ref_point14_valc_001"}
		}, Point14ValDStateBlocked},
		{"missing redaction refs where redacted blocks", func(model *CorrectionRevocationPublicationReadProjection) {
			model.RedactionRefs = nil
			model.LimitationRefs = []string{"limitation_ref_point14_valc_001"}
		}, Point14ValDStateActive},
		{"correction publication flag blocks", func(model *CorrectionRevocationPublicationReadProjection) { model.PublishesCorrection = true }, Point14ValDStateBlocked},
		{"revocation execution flag blocks", func(model *CorrectionRevocationPublicationReadProjection) { model.ExecutesRevocation = true }, Point14ValDStateBlocked},
		{"supersession silent replacement flag blocks", func(model *CorrectionRevocationPublicationReadProjection) {
			model.SilentReplacesSupersededSignal = true
		}, Point14ValDStateBlocked},
		{"limitation omission blocks", func(model *CorrectionRevocationPublicationReadProjection) { model.OmitsLimitations = true }, Point14ValDStateBlocked},
		{"redaction hiding blocks", func(model *CorrectionRevocationPublicationReadProjection) { model.HidesRedaction = true }, Point14ValDStateBlocked},
		{"production compliance legal approval wording blocks", func(model *CorrectionRevocationPublicationReadProjection) {
			model.ObservedReadTexts = []string{"publication proves safety"}
		}, Point14ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValDCorrectionReadProjectionModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValDCorrectionReadProjectionState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDGovernanceTraceProjectionState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*GovernanceTraceReadProjection)
		want   string
	}{
		{"valid read only governance trace projection active", func(model *GovernanceTraceReadProjection) {}, Point14ValDStateActive},
		{"missing governance trace refs blocks", func(model *GovernanceTraceReadProjection) { model.GovernanceTraceRefs = nil }, Point14ValDStateBlocked},
		{"missing owner refs blocks", func(model *GovernanceTraceReadProjection) { model.OwnerRefs = nil }, Point14ValDStateBlocked},
		{"missing approver role refs blocks", func(model *GovernanceTraceReadProjection) { model.ApproverRoleRefs = nil }, Point14ValDStateBlocked},
		{"missing audit refs blocks", func(model *GovernanceTraceReadProjection) { model.AuditRefs = nil }, Point14ValDStateBlocked},
		{"missing evidence refs blocks", func(model *GovernanceTraceReadProjection) { model.EvidenceRefs = nil }, Point14ValDStateBlocked},
		{"missing decision reason refs blocks", func(model *GovernanceTraceReadProjection) { model.DecisionReasonRefs = nil }, Point14ValDStateBlocked},
		{"read only false blocks", func(model *GovernanceTraceReadProjection) { model.ReadOnly = false }, Point14ValDStateBlocked},
		{"governance trace projection cannot approve anything", func(model *GovernanceTraceReadProjection) { model.ApprovesAnything = true }, Point14ValDStateBlocked},
		{"display alone cannot satisfy missing governance", func(model *GovernanceTraceReadProjection) { model.SatisfiesMissingGovernanceByDisplay = true }, Point14ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValDGovernanceTraceReadProjectionModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValDGovernanceTraceProjectionState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDQueryProjectionState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*EcosystemTimelineQueryProjection)
		want   string
	}{
		{"valid query projection active", func(model *EcosystemTimelineQueryProjection) {}, Point14ValDStateActive},
		{"missing query projection id blocks", func(model *EcosystemTimelineQueryProjection) { model.QueryProjectionID = "" }, Point14ValDStateBlocked},
		{"missing query scope blocks", func(model *EcosystemTimelineQueryProjection) { model.QueryScope = "" }, Point14ValDStateBlocked},
		{"missing allowed filters blocks", func(model *EcosystemTimelineQueryProjection) { model.AllowedFilters = nil }, Point14ValDStateBlocked},
		{"missing result refs blocks", func(model *EcosystemTimelineQueryProjection) { model.ResultRefs = nil }, Point14ValDStateBlocked},
		{"read only false blocks", func(model *EcosystemTimelineQueryProjection) { model.ReadOnly = false }, Point14ValDStateBlocked},
		{"query is projection only false blocks", func(model *EcosystemTimelineQueryProjection) { model.QueryIsProjectionOnly = false }, Point14ValDStateBlocked},
		{"query mutation write flag blocks", func(model *EcosystemTimelineQueryProjection) { model.MutationRequested = true }, Point14ValDStateBlocked},
		{"filter hiding decisive missing evidence blocks", func(model *EcosystemTimelineQueryProjection) { model.HidesDecisiveMissingEvidence = true }, Point14ValDStateBlocked},
		{"filter omitting limitations without disclosure blocks", func(model *EcosystemTimelineQueryProjection) { model.OmitsLimitationsWithoutDisclosure = true }, Point14ValDStateBlocked},
		{"cross tenant query result blocks", func(model *EcosystemTimelineQueryProjection) { model.CrossTenantResults = true }, Point14ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValDQueryProjectionModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValDQueryProjectionState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDAccessBoundaryState(t *testing.T) {
	base := point14ValDDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*EcosystemTimelineAccessBoundary)
		want   string
	}{
		{"valid customer tenant scoped access active", func(model *EcosystemTimelineAccessBoundary) {}, Point14ValDStateActive},
		{"valid auditor tenant scoped access active", func(model *EcosystemTimelineAccessBoundary) {
			model.ViewerRole = "auditor"
			model.AllowedViewScope = point14ValCVisibilityAuditorBounded
		}, Point14ValDStateActive},
		{"missing access boundary id blocks", func(model *EcosystemTimelineAccessBoundary) { model.AccessBoundaryID = "" }, Point14ValDStateBlocked},
		{"missing viewer role blocks", func(model *EcosystemTimelineAccessBoundary) { model.ViewerRole = "" }, Point14ValDStateBlocked},
		{"missing tenant scope global classification blocks", func(model *EcosystemTimelineAccessBoundary) {
			model.TenantScope = ""
			model.GlobalScopeClassification = ""
		}, Point14ValDStateBlocked},
		{"missing audit ref blocks", func(model *EcosystemTimelineAccessBoundary) { model.AuditRef = "" }, Point14ValDStateBlocked},
		{"missing access time blocks", func(model *EcosystemTimelineAccessBoundary) { model.AccessTime = "" }, Point14ValDStateBlocked},
		{"client local access time source blocks", func(model *EcosystemTimelineAccessBoundary) {
			model.AccessTimeSource = point14Val0TimeSourceClientLocal
		}, Point14ValDStateBlocked},
		{"expired access blocks", func(model *EcosystemTimelineAccessBoundary) { model.AccessExpired = true }, Point14ValDStateBlocked},
		{"revoked access blocks", func(model *EcosystemTimelineAccessBoundary) { model.AccessRevoked = true }, Point14ValDStateBlocked},
		{"cross tenant access blocks", func(model *EcosystemTimelineAccessBoundary) { model.CrossTenantAccess = true }, Point14ValDStateBlocked},
		{"auditor customer public viewer authority grant blocks", func(model *EcosystemTimelineAccessBoundary) { model.AuthorityGranted = true }, Point14ValDStateBlocked},
		{"public viewer tenant private data exposure blocks", func(model *EcosystemTimelineAccessBoundary) {
			model.ViewerRole = point14ValDViewerPublic
			model.AllowedViewScope = point14ValCVisibilityPrivateTenantOnly
			model.TenantPrivateDataExposed = true
		}, Point14ValDStateBlocked},
		{"customer tenant private data exposure blocks", func(model *EcosystemTimelineAccessBoundary) {
			model.ViewerRole = "customer_admin"
			model.AllowedViewScope = point14ValCVisibilityCustomerBounded
			model.TenantPrivateDataExposed = true
		}, Point14ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dependency := clonePoint14ValDDependencySnapshot(t, base)
			model := point14ValDAccessBoundaryModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValDAccessBoundaryState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDTenantPrivacyTimelineProjectionGuardState(t *testing.T) {
	base := point14ValDDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*TenantPrivacyTimelineProjectionGuard)
		want   string
	}{
		{"valid private tenant projection active", func(model *TenantPrivacyTimelineProjectionGuard) {}, Point14ValDStateActive},
		{"cross tenant projection blocks", func(model *TenantPrivacyTimelineProjectionGuard) { model.CrossTenantProjection = true }, Point14ValDStateBlocked},
		{"tenant private data exposed blocks", func(model *TenantPrivacyTimelineProjectionGuard) { model.TenantPrivateDataExposed = true }, Point14ValDStateBlocked},
		{"public private classification missing blocks", func(model *TenantPrivacyTimelineProjectionGuard) { model.PublicPrivateClassification = "" }, Point14ValDStateBlocked},
		{"redaction limitation refs missing when limited blocks", func(model *TenantPrivacyTimelineProjectionGuard) { model.LimitationRefs = nil }, Point14ValDStateBlocked},
		{"timeline summary leaking private evidence blocks", func(model *TenantPrivacyTimelineProjectionGuard) {
			model.ObservedSummaryTexts = []string{"publication proves safety"}
		}, Point14ValDStateBlocked},
		{"timeline summary strengthening claim blocks", func(model *TenantPrivacyTimelineProjectionGuard) { model.StrengthensClaim = true }, Point14ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dependency := clonePoint14ValDDependencySnapshot(t, base)
			model := point14ValDTenantPrivacyTimelineProjectionGuardModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValDTenantPrivacyTimelineProjectionGuardState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDAgentTimelineProjectionState(t *testing.T) {
	base := point14ValDDependencySnapshotModel()
	tests := []struct {
		name   string
		mutate func(*AgentEcosystemTimelineProjection)
		want   string
	}{
		{"agent recommendation may be displayed as advisory input", func(model *AgentEcosystemTimelineProjection) {}, Point14ValDStateActive},
		{"agent approval flag blocks", func(model *AgentEcosystemTimelineProjection) { model.AgentApprovalFlags = true }, Point14ValDStateBlocked},
		{"agent authority flag blocks", func(model *AgentEcosystemTimelineProjection) { model.AgentAuthorityFlags = true }, Point14ValDStateBlocked},
		{"agent timeline cannot resolve dispute", func(model *AgentEcosystemTimelineProjection) { model.CanResolveDispute = true }, Point14ValDStateBlocked},
		{"agent timeline cannot publish correction", func(model *AgentEcosystemTimelineProjection) { model.CanPublishCorrection = true }, Point14ValDStateBlocked},
		{"agent timeline cannot revoke claim", func(model *AgentEcosystemTimelineProjection) { model.CanRevokeClaim = true }, Point14ValDStateBlocked},
		{"agent timeline cannot satisfy governance trace", func(model *AgentEcosystemTimelineProjection) { model.CanSatisfyGovernanceTrace = true }, Point14ValDStateBlocked},
		{"ai agent authority flags block globally", func(model *AgentEcosystemTimelineProjection) { model.ExternalAuthorityAllowed = true }, Point14ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dependency := clonePoint14ValDDependencySnapshot(t, base)
			model := point14ValDAgentTimelineProjectionModel(dependency)
			tc.mutate(&model)
			if got := EvaluatePoint14ValDAgentTimelineProjectionState(model, dependency); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDTimestampIntegrityGuardState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*TimelineTimestampIntegrityGuard)
		want   string
	}{
		{"server utc generated at event at access time passes", func(model *TimelineTimestampIntegrityGuard) {}, Point14ValDStateActive},
		{"approved customer controlled time source passes", func(model *TimelineTimestampIntegrityGuard) {
			model.EventTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.GeneratedTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.AccessTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.ReceivedTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.PublicationApprovedTimeSource = point14Val0TimeSourceApprovedCustomerTime
			model.DisputeOpenedTimeSource = point14Val0TimeSourceApprovedCustomerTime
		}, Point14ValDStateActive},
		{"client local time as canonical generated at blocks", func(model *TimelineTimestampIntegrityGuard) {
			model.GeneratedTimeSource = point14Val0TimeSourceClientLocal
		}, Point14ValDStateBlocked},
		{"client local time as canonical event at blocks", func(model *TimelineTimestampIntegrityGuard) { model.EventTimeSource = point14Val0TimeSourceClientLocal }, Point14ValDStateBlocked},
		{"future dated timeline event blocks review required", func(model *TimelineTimestampIntegrityGuard) { model.EventAt = "2026-05-06T09:12:00Z" }, Point14ValDStateReviewRequired},
		{"impossible ordering blocks review required", func(model *TimelineTimestampIntegrityGuard) { model.GeneratedAt = "2026-05-06T08:59:00Z" }, Point14ValDStateReviewRequired},
		{"backdated publication approval display blocks review required", func(model *TimelineTimestampIntegrityGuard) { model.PublicationApprovedAt = "2026-05-06T08:59:00Z" }, Point14ValDStateReviewRequired},
		{"timeline ordering cannot upgrade validity", func(model *TimelineTimestampIntegrityGuard) { model.AttemptsValidityUpgrade = true }, Point14ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValDTimestampIntegrityGuardModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValDTimestampIntegrityGuardState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint14ValDNoMutationProjectionGuardState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point14ValDNoMutationProjectionGuard)
	}{
		{"canonical evidence mutation attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.MutatesCanonicalEvidence = true }},
		{"signal mutation attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.MutatesNormalizedSignal = true }},
		{"dispute lifecycle mutation attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.MutatesDisputeLifecycle = true }},
		{"correction notice mutation attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.MutatesCorrectionNotice = true }},
		{"revocation request mutation attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.MutatesRevocationRequest = true }},
		{"publication approval mutation attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.MutatesPublicationApproval = true }},
		{"governance trace mutation attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.MutatesGovernanceTrace = true }},
		{"publish correction attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.PublishesCorrection = true }},
		{"execute revocation attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.ExecutesRevocation = true }},
		{"resolve dispute attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.ResolvesDispute = true }},
		{"approve production attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.ApprovesProduction = true }},
		{"certify compliance attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.CertifiesCompliance = true }},
		{"create public badge attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.CreatesPublicBadge = true }},
		{"emit pass attempt blocks", func(model *Point14ValDNoMutationProjectionGuard) { model.EmitsPass = true }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point14ValDNoMutationProjectionGuardModel()
			tc.mutate(&model)
			if got := EvaluatePoint14ValDNoMutationProjectionGuardState(model); got != Point14ValDStateBlocked {
				t.Fatalf("expected blocked, got %s", got)
			}
		})
	}
}

func TestPoint14ValDNoOverclaimTimelineWordingState(t *testing.T) {
	t.Run("forbidden timeline wording blocks", func(t *testing.T) {
		model := point14ValDNoOverclaimTimelineWordingModel()
		model.ObservedTimelineTexts = []string{"timeline proves truth"}
		if got := EvaluatePoint14ValDNoOverclaimTimelineWordingState(model); got != Point14ValDStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("safe bounded wording passes", func(t *testing.T) {
		model := point14ValDNoOverclaimTimelineWordingModel()
		if got := EvaluatePoint14ValDNoOverclaimTimelineWordingState(model); got != Point14ValDStateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})

	t.Run("forbidden wording allowed only in internal blocked diagnostic context", func(t *testing.T) {
		model := point14ValDNoOverclaimTimelineWordingModel()
		model.InternalDiagnosticTexts = []string{"timeline approved"}
		model.InternalDiagnosticsClassifiedBlocked = true
		if got := EvaluatePoint14ValDNoOverclaimTimelineWordingState(model); got != Point14ValDStateActive {
			t.Fatalf("expected active, got %s", got)
		}
	})
}

func TestPoint14ValDStateAggregation(t *testing.T) {
	t.Run("any component blocked returns blocked", func(t *testing.T) {
		model := Point14ValDFoundationModel()
		model.QueryProjectionState = Point14ValDStateBlocked
		if got := point14ValDFoundationState(
			model.DependencyState,
			model.TimelineProjectionState,
			model.SignalTimelineEntryState,
			model.DisputeTimelineState,
			model.CorrectionReadProjectionState,
			model.GovernanceTraceProjectionState,
			model.QueryProjectionState,
			model.AccessBoundaryState,
			model.TenantPrivacyTimelineState,
			model.AgentTimelineProjectionState,
			model.TimestampIntegrityState,
			model.NoMutationProjectionGuardState,
			model.NoOverclaimTimelineWordingState,
		); got != Point14ValDStateBlocked {
			t.Fatalf("expected blocked, got %s", got)
		}
	})

	t.Run("any review required and no blocked returns review required", func(t *testing.T) {
		model := Point14ValDFoundationModel()
		model.TimestampIntegrityState = Point14ValDStateReviewRequired
		if got := point14ValDFoundationState(
			model.DependencyState,
			model.TimelineProjectionState,
			model.SignalTimelineEntryState,
			model.DisputeTimelineState,
			model.CorrectionReadProjectionState,
			model.GovernanceTraceProjectionState,
			model.QueryProjectionState,
			model.AccessBoundaryState,
			model.TenantPrivacyTimelineState,
			model.AgentTimelineProjectionState,
			model.TimestampIntegrityState,
			model.NoMutationProjectionGuardState,
			model.NoOverclaimTimelineWordingState,
		); got != Point14ValDStateReviewRequired {
			t.Fatalf("expected review_required, got %s", got)
		}
	})

	t.Run("any incomplete and no blocked review required returns incomplete", func(t *testing.T) {
		model := Point14ValDFoundationModel()
		model.DisputeTimelineState = Point14ValDStateIncomplete
		if got := point14ValDFoundationState(
			model.DependencyState,
			model.TimelineProjectionState,
			model.SignalTimelineEntryState,
			model.DisputeTimelineState,
			model.CorrectionReadProjectionState,
			model.GovernanceTraceProjectionState,
			model.QueryProjectionState,
			model.AccessBoundaryState,
			model.TenantPrivacyTimelineState,
			model.AgentTimelineProjectionState,
			model.TimestampIntegrityState,
			model.NoMutationProjectionGuardState,
			model.NoOverclaimTimelineWordingState,
		); got != Point14ValDStateIncomplete {
			t.Fatalf("expected incomplete, got %s", got)
		}
	})

	t.Run("active only when all components active", func(t *testing.T) {
		model := ComputePoint14ValDFoundation(Point14ValDFoundationModel())
		if model.CurrentState != Point14ValDStateActive {
			t.Fatalf("expected active, got %#v", model)
		}
	})
}

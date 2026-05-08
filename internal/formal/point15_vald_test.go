package formal

import (
	"encoding/json"
	"sync"
	"testing"
)

var (
	point15ValDFoundationOnce sync.Once
	point15ValDFoundationBase Point15ValDAssuranceProjectionFoundation
)

func point15ValDCloneStrings(values []string) []string {
	return append([]string(nil), values...)
}

func clonePoint15ValDFoundation(model Point15ValDAssuranceProjectionFoundation) Point15ValDAssuranceProjectionFoundation {
	model.BlockingReasons = point15ValDCloneStrings(model.BlockingReasons)
	model.ReviewPrerequisites = point15ValDCloneStrings(model.ReviewPrerequisites)
	model.Dependency.ReviewPrerequisites = point15ValDCloneStrings(model.Dependency.ReviewPrerequisites)
	model.Dependency.Point15ValC = clonePoint15ValCFoundation(model.Dependency.Point15ValC)
	model.Query.Filters = point15ValDCloneStrings(model.Query.Filters)
	model.Query.ResultRefs = point15ValDCloneStrings(model.Query.ResultRefs)
	model.NoOverclaimGuard.ObservedTexts = point15ValDCloneStrings(model.NoOverclaimGuard.ObservedTexts)
	model.NoOverclaimGuard.InternalDiagnosticTexts = point15ValDCloneStrings(model.NoOverclaimGuard.InternalDiagnosticTexts)
	model.NoOverclaimGuard.AllowedSafeWording = point15ValDCloneStrings(model.NoOverclaimGuard.AllowedSafeWording)
	model.NoOverclaimGuard.BlockedWording = point15ValDCloneStrings(model.NoOverclaimGuard.BlockedWording)
	return model
}

func point15ValDValidFoundationModel() Point15ValDAssuranceProjectionFoundation {
	point15ValDFoundationOnce.Do(func() {
		point15ValDFoundationBase = Point15ValDFoundationModel()
	})
	return clonePoint15ValDFoundation(point15ValDFoundationBase)
}

func point15ValDValidDependencyModel() Point15ValDDependencySnapshot {
	return point15ValDValidFoundationModel().Dependency
}

func point15ValDValidTimelineModel() Point15ValDAssuranceTimelineEntry {
	return point15ValDValidFoundationModel().Timeline
}

func point15ValDValidDashboardModel() Point15ValDDashboardSummary {
	return point15ValDValidFoundationModel().Dashboard
}

func point15ValDValidQueryModel() Point15ValDQueryProjection {
	return point15ValDValidFoundationModel().Query
}

func point15ValDValidEvidenceDetailModel() Point15ValDEvidenceDetailProjection {
	return point15ValDValidFoundationModel().EvidenceDetail
}

func point15ValDValidRevalidationDetailModel() Point15ValDRevalidationDetailProjection {
	return point15ValDValidFoundationModel().RevalidationDetail
}

func point15ValDValidEnforcementDetailModel() Point15ValDEnforcementDetailProjection {
	return point15ValDValidFoundationModel().EnforcementDetail
}

func point15ValDValidReplayHistoryModel() Point15ValDReplayProofHistoryProjection {
	return point15ValDValidFoundationModel().ReplayProofHistory
}

func point15ValDValidAccessTenantModel() Point15ValDAccessTenantPrivacyBoundary {
	return point15ValDValidFoundationModel().AccessTenantPrivacy
}

func point15ValDValidTimestampDisplayModel() Point15ValDTimestampDisplayDiscipline {
	return point15ValDValidFoundationModel().TimestampDisplayDiscipline
}

func point15ValDValidNoMutationModel() Point15ValDNoMutationProjectionGuard {
	return point15ValDValidFoundationModel().NoMutationGuard
}

func point15ValDValidAuthorityModel() Point15ValDAuthorityBoundary {
	return point15ValDValidFoundationModel().AuthorityBoundary
}

func point15ValDValidNoOverclaimModel() Point15ValDNoOverclaimGuard {
	return point15ValDValidFoundationModel().NoOverclaimGuard
}

func TestPoint15ValDDependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValDDependencySnapshot)
		want   string
	}{
		{"active when valc clean", func(*Point15ValDDependencySnapshot) {}, Point15ValDStateActive},
		{"blocks when valc missing", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValCCurrentState = ""
		}, Point15ValDStateBlocked},
		{"blocks when valc blocked", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValCCurrentState = Point15ValCStateBlocked
		}, Point15ValDStateBlocked},
		{"blocks when valc review required", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValCCurrentState = Point15ValCStateReviewRequired
		}, Point15ValDStateBlocked},
		{"blocks when valc incomplete", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValCCurrentState = Point15ValCStateIncomplete
		}, Point15ValDStateBlocked},
		{"blocks when valc not merged", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValCMerged = false
		}, Point15ValDStateBlocked},
		{"blocks when valc ci not green", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValCCIGreen = false
		}, Point15ValDStateBlocked},
		{"blocks when valc not reviewed on main", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValCReviewedOnMain = false
		}, Point15ValDStateBlocked},
		{"blocks on point15 pass token", func(model *Point15ValDDependencySnapshot) {
			model.Point15PassSeen = true
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValDValidDependencyModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValDDependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDAssuranceTimelineState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDAssuranceTimelineEntry
		want  string
	}{
		{"timeline active when display is exact", func() Point15ValDAssuranceTimelineEntry {
			return point15ValDValidTimelineModel()
		}, Point15ValDStateActive},
		{"missing blocked reason for blocked current state blocks", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.CurrentState = Point15Val0StateBlocked
			model.EnforcementReason = point15ValCReasonExpired
			model.BlockedReasonVisible = false
			return model
		}, Point15ValDStateBlocked},
		{"missing decisive evidence visibility becomes review required for review state", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.CurrentState = Point15Val0StateReviewRequired
			model.EnforcementReason = point15ValCReasonStale
			model.DecisiveEvidenceVisible = false
			return model
		}, Point15ValDStateReviewRequired},
		{"timeline cannot create validity by ordering", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.TimelineCreatesValidity = true
			return model
		}, Point15ValDStateBlocked},
		{"timeline cannot hide prior pass or later downgrade", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.PriorPassVisible = false
			return model
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDAssuranceTimelineState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDDashboardSummaryState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDDashboardSummary
		want  string
	}{
		{"summary is display only", func() Point15ValDDashboardSummary {
			return point15ValDValidDashboardModel()
		}, Point15ValDStateActive},
		{"blocked count cannot be hidden", func() Point15ValDDashboardSummary {
			model := point15ValDValidDashboardModel()
			model.HiddenBlockedCount = true
			return model
		}, Point15ValDStateBlocked},
		{"active count cannot include disallowed evidence", func() Point15ValDDashboardSummary {
			model := point15ValDValidDashboardModel()
			model.ActiveCountIncludesDisallowed = true
			return model
		}, Point15ValDStateBlocked},
		{"dashboard cannot restore active closure", func() Point15ValDDashboardSummary {
			model := point15ValDValidDashboardModel()
			model.RestoresActiveClosure = true
			return model
		}, Point15ValDStateBlocked},
		{"missing tenant scope is incomplete", func() Point15ValDDashboardSummary {
			model := point15ValDValidDashboardModel()
			model.TenantScope = ""
			return model
		}, Point15ValDStateIncomplete},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDDashboardSummaryState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDQueryProjectionState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDQueryProjection
		want  string
	}{
		{"query may filter only", func() Point15ValDQueryProjection {
			return point15ValDValidQueryModel()
		}, Point15ValDStateActive},
		{"query cannot mutate", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.QueryMutationAttempted = true
			return model
		}, Point15ValDStateBlocked},
		{"cross tenant query blocks", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.CrossTenantQuery = true
			return model
		}, Point15ValDStateBlocked},
		{"query cannot suppress decisive evidence", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.ResultState = point15ValDQueryRedacted
			model.DecisiveEvidenceVisible = false
			return model
		}, Point15ValDStateReviewRequired},
		{"query cannot strengthen claims", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.StrengthensClaims = true
			return model
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDQueryProjectionState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDEvidenceDetailProjectionState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDEvidenceDetailProjection
		want  string
	}{
		{"evidence detail exact binding active", func() Point15ValDEvidenceDetailProjection {
			return point15ValDValidEvidenceDetailModel()
		}, Point15ValDStateActive},
		{"missing limitations visibility requires review", func() Point15ValDEvidenceDetailProjection {
			model := point15ValDValidEvidenceDetailModel()
			model.LimitationsVisible = false
			return model
		}, Point15ValDStateReviewRequired},
		{"similar names do not imply identity", func() Point15ValDEvidenceDetailProjection {
			model := point15ValDValidEvidenceDetailModel()
			model.IdentityDerivedFromNameOnly = true
			return model
		}, Point15ValDStateBlocked},
		{"missing hash blocks", func() Point15ValDEvidenceDetailProjection {
			model := point15ValDValidEvidenceDetailModel()
			model.EvidenceHash = ""
			return model
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDEvidenceDetailProjectionState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDRevalidationDetailProjectionState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDRevalidationDetailProjection
		want  string
	}{
		{"revalidation detail display only active", func() Point15ValDRevalidationDetailProjection {
			return point15ValDValidRevalidationDetailModel()
		}, Point15ValDStateActive},
		{"cannot schedule from projection", func() Point15ValDRevalidationDetailProjection {
			model := point15ValDValidRevalidationDetailModel()
			model.ScheduleMutationAttempted = true
			return model
		}, Point15ValDStateBlocked},
		{"cannot retry from projection", func() Point15ValDRevalidationDetailProjection {
			model := point15ValDValidRevalidationDetailModel()
			model.RetryTriggered = true
			return model
		}, Point15ValDStateBlocked},
		{"cannot reset budget", func() Point15ValDRevalidationDetailProjection {
			model := point15ValDValidRevalidationDetailModel()
			model.RetryBudgetResetAttempted = true
			return model
		}, Point15ValDStateBlocked},
		{"cannot mark fresh or restore active", func() Point15ValDRevalidationDetailProjection {
			model := point15ValDValidRevalidationDetailModel()
			model.MarksFresh = true
			return model
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDRevalidationDetailProjectionState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDEnforcementDetailProjectionState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDEnforcementDetailProjection
		want  string
	}{
		{"displays enforcement only", func() Point15ValDEnforcementDetailProjection {
			return point15ValDValidEnforcementDetailModel()
		}, Point15ValDStateActive},
		{"cannot perform enforcement", func() Point15ValDEnforcementDetailProjection {
			model := point15ValDValidEnforcementDetailModel()
			model.PerformsEnforcement = true
			return model
		}, Point15ValDStateBlocked},
		{"cannot auto revoke publish or delete", func() Point15ValDEnforcementDetailProjection {
			model := point15ValDValidEnforcementDetailModel()
			model.AutoPublishes = true
			return model
		}, Point15ValDStateBlocked},
		{"history preserved must remain visible", func() Point15ValDEnforcementDetailProjection {
			model := point15ValDValidEnforcementDetailModel()
			model.HistoryPreserved = false
			return model
		}, Point15ValDStateBlocked},
		{"blocked reason required for blocked current state", func() Point15ValDEnforcementDetailProjection {
			model := point15ValDValidEnforcementDetailModel()
			model.EnforcementAction = point15ValCActionBlocked
			model.EnforcementReason = point15ValCReasonExpired
			model.TargetState = Point15Val0StateBlocked
			model.CurrentState = Point15Val0StateBlocked
			model.BlockedReasonVisible = false
			return model
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDEnforcementDetailProjectionState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDReplayProofHistoryProjectionState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDReplayProofHistoryProjection
		want  string
	}{
		{"replay history active", func() Point15ValDReplayProofHistoryProjection {
			return point15ValDValidReplayHistoryModel()
		}, Point15ValDStateActive},
		{"replay refs required", func() Point15ValDReplayProofHistoryProjection {
			model := point15ValDValidReplayHistoryModel()
			model.ReplayRef = ""
			return model
		}, Point15ValDStateBlocked},
		{"decisive evidence visible required", func() Point15ValDReplayProofHistoryProjection {
			model := point15ValDValidReplayHistoryModel()
			model.DecisiveEvidenceVisible = false
			return model
		}, Point15ValDStateBlocked},
		{"hash binding visible required", func() Point15ValDReplayProofHistoryProjection {
			model := point15ValDValidReplayHistoryModel()
			model.HashBindingVisible = false
			return model
		}, Point15ValDStateBlocked},
		{"projection cannot hide proof history", func() Point15ValDReplayProofHistoryProjection {
			model := point15ValDValidReplayHistoryModel()
			model.ProofHistoryHidden = true
			return model
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDReplayProofHistoryProjectionState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDAccessTenantPrivacyBoundaryState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDAccessTenantPrivacyBoundary
		want  string
	}{
		{"tenant scoped access active", func() Point15ValDAccessTenantPrivacyBoundary {
			return point15ValDValidAccessTenantModel()
		}, Point15ValDStateActive},
		{"cross tenant access blocks", func() Point15ValDAccessTenantPrivacyBoundary {
			model := point15ValDValidAccessTenantModel()
			model.CrossTenantDetected = true
			return model
		}, Point15ValDStateBlocked},
		{"tenant private exposure blocks", func() Point15ValDAccessTenantPrivacyBoundary {
			model := point15ValDValidAccessTenantModel()
			model.TenantPrivateDataExposed = true
			return model
		}, Point15ValDStateBlocked},
		{"public visibility forbidden by default", func() Point15ValDAccessTenantPrivacyBoundary {
			model := point15ValDValidAccessTenantModel()
			model.Visibility = point15ValDVisibilityPublicBlocked
			return model
		}, Point15ValDStateBlocked},
		{"redaction cannot hide decisive failure without review", func() Point15ValDAccessTenantPrivacyBoundary {
			model := point15ValDValidAccessTenantModel()
			model.DecisiveFailureHidden = true
			model.RedactionState = point15ValDRedactionLimited
			return model
		}, Point15ValDStateReviewRequired},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDAccessTenantPrivacyBoundaryState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDTimestampDisplayDisciplineState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDTimestampDisplayDiscipline
		want  string
	}{
		{"trusted display timing active", func() Point15ValDTimestampDisplayDiscipline {
			return point15ValDValidTimestampDisplayModel()
		}, Point15ValDStateActive},
		{"client local display time cannot create canonical validity", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.ClientLocalCreatesCanonical = true
			return model
		}, Point15ValDStateBlocked},
		{"source event alone cannot determine current state", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.SourceEventAt = "2026-05-07T08:00:00Z"
			model.EventAt = ""
			return model
		}, Point15ValDStateBlocked},
		{"displayed before event blocks", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.DisplayedAt = "2026-05-07T07:00:00Z"
			model.EventAt = "2026-05-07T08:00:00Z"
			return model
		}, Point15ValDStateBlocked},
		{"enforced before validated requires review", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.ValidatedAt = "2026-05-07T09:00:00Z"
			model.EnforcedAt = "2026-05-07T08:59:00Z"
			return model
		}, Point15ValDStateReviewRequired},
		{"future canonical event requires review", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.EventAt = "2026-05-07T12:00:00Z"
			model.ReferenceNow = "2026-05-07T11:00:00Z"
			model.DisplayedAt = "2026-05-07T12:01:00Z"
			return model
		}, Point15ValDStateReviewRequired},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDTimestampDisplayDisciplineState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDNoMutationProjectionGuardState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDNoMutationProjectionGuard
		want  string
	}{
		{"no mutation active", func() Point15ValDNoMutationProjectionGuard {
			return point15ValDValidNoMutationModel()
		}, Point15ValDStateActive},
		{"evidence mutation attempt blocks", func() Point15ValDNoMutationProjectionGuard {
			model := point15ValDValidNoMutationModel()
			model.EvidenceMutationAttempted = true
			return model
		}, Point15ValDStateBlocked},
		{"enforcement mutation attempt blocks", func() Point15ValDNoMutationProjectionGuard {
			model := point15ValDValidNoMutationModel()
			model.EnforcementMutationAttempted = true
			return model
		}, Point15ValDStateBlocked},
		{"pass restore attempt blocks", func() Point15ValDNoMutationProjectionGuard {
			model := point15ValDValidNoMutationModel()
			model.PassRestoreAttempted = true
			return model
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDNoMutationProjectionGuardState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDAuthorityBoundaryState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDAuthorityBoundary
		want  string
	}{
		{"authority boundary active", func() Point15ValDAuthorityBoundary {
			return point15ValDValidAuthorityModel()
		}, Point15ValDStateActive},
		{"dashboard cannot approve pass", func() Point15ValDAuthorityBoundary {
			model := point15ValDValidAuthorityModel()
			model.DashboardApprovesPass = true
			return model
		}, Point15ValDStateBlocked},
		{"timeline cannot create authority", func() Point15ValDAuthorityBoundary {
			model := point15ValDValidAuthorityModel()
			model.TimelineCreatesAuthority = true
			return model
		}, Point15ValDStateBlocked},
		{"query cannot enforce", func() Point15ValDAuthorityBoundary {
			model := point15ValDValidAuthorityModel()
			model.QueryEnforcesState = true
			return model
		}, Point15ValDStateBlocked},
		{"export preview cannot publish", func() Point15ValDAuthorityBoundary {
			model := point15ValDValidAuthorityModel()
			model.ExportPreviewPublishes = true
			return model
		}, Point15ValDStateBlocked},
		{"connector scheduler and agent cannot become authority", func() Point15ValDAuthorityBoundary {
			model := point15ValDValidAuthorityModel()
			model.AgentAuthorityGranted = true
			return model
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDAuthorityBoundaryState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDNoOverclaimGuardState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValDNoOverclaimGuard
		want  string
	}{
		{"safe bounded wording passes", func() Point15ValDNoOverclaimGuard {
			return point15ValDValidNoOverclaimModel()
		}, Point15ValDStateActive},
		{"forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"continuous assurance guaranteed"}
			return model
		}, Point15ValDStateBlocked},
		{"classified internal blocked diagnostics allowed", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.InternalDiagnosticTexts = []string{"production approved"}
			model.InternalDiagnosticsClassifiedBlocked = true
			return model
		}, Point15ValDStateActive},
		{"unclassified internal blocked diagnostic blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.InternalDiagnosticTexts = []string{"global truth"}
			return model
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValDNoOverclaimGuardState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValDAssuranceProjectionFoundationState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValDAssuranceProjectionFoundation)
		want   string
	}{
		{"happy path active", func(*Point15ValDAssuranceProjectionFoundation) {}, Point15ValDStateActive},
		{"dashboard count mismatch blocks", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.Dashboard.ActiveCount = 0
		}, Point15ValDStateBlocked},
		{"query ref mismatch blocks", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.Query.ResultRefs = []string{"wrong_ref"}
		}, Point15ValDStateBlocked},
		{"evidence detail hash mismatch blocks", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.EvidenceDetail.EvidenceHash = "wrong_hash"
		}, Point15ValDStateBlocked},
		{"revalidation detail ref mismatch blocks", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.RevalidationDetail.ScheduleRef = "wrong_schedule"
		}, Point15ValDStateBlocked},
		{"timeline source valc ref mismatch blocks", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.Timeline.SourceValCRef = "wrong_enforcement_ref"
		}, Point15ValDStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValDValidFoundationModel()
			tc.mutate(&model)
			if got := ComputePoint15ValDAssuranceProjectionFoundation(model).CurrentState; got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint10ThroughPoint15ValDCurrentSweep(t *testing.T) {
	model := ComputePoint15ValDAssuranceProjectionFoundation(Point15ValDFoundationModel())
	if model.CurrentState != Point15ValDStateActive {
		t.Fatalf("expected %s, got %s", Point15ValDStateActive, model.CurrentState)
	}
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if string(payload) == "" {
		t.Fatalf("expected non-empty payload")
	}
	if contains := point15ValDValCPayloadContainsPoint15Pass(model.Dependency.Point15ValC); contains {
		t.Fatalf("unexpected point15 pass token in upstream payload")
	}
	if string(payload) != "" && json.Valid(payload) == false {
		t.Fatalf("expected valid json payload")
	}
}

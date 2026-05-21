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

func point15ValDStringSliceContains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
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
		{"blocks whitespace retagged valc current state", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValCCurrentState = " " + Point15ValCStateActive + " "
		}, Point15ValDStateBlocked},
		{"blocks tab newline retagged inherited point14 state", func(model *Point15ValDDependencySnapshot) {
			model.InheritedPoint14ValECurrentState = "\t" + Point14ValEStatePassConfirmed + "\n"
		}, Point15ValDStateBlocked},
		{"blocks padded inherited tenant scope", func(model *Point15ValDDependencySnapshot) {
			model.InheritedTenantScope = " " + model.InheritedTenantScope + " "
		}, Point15ValDStateBlocked},
		{"blocks embedded valc state laundering behind clean flat fields", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.CurrentState = Point15ValCStateBlocked
		}, Point15ValDStateBlocked},
		{"blocks stale embedded valc no-overclaim allowed ledger mutation", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.NoOverclaimGuard.AllowedSafeWording = append(model.Point15ValC.NoOverclaimGuard.AllowedSafeWording, "production approved")
		}, Point15ValDStateBlocked},
		{"blocks stale embedded valc split no-overclaim wording", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.NoOverclaimGuard.ObservedTexts = []string{"continuous assurance", "guaranteed"}
		}, Point15ValDStateBlocked},
		{"blocks stale embedded valc confusable no-overclaim wording", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.NoOverclaimGuard.ObservedTexts = []string{"production appro\u03bded"}
		}, Point15ValDStateBlocked},
		{"blocks stale embedded valc timestamp ordering mutation", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.TimestampDiscipline.ReceivedAt = "2026-05-07T09:06:00Z"
		}, Point15ValDStateBlocked},
		{"blocks stale embedded valb no-overclaim blocked ledger mutation", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.Dependency.Point15ValB.NoOverclaimGuard.BlockedWording = append(model.Point15ValC.Dependency.Point15ValB.NoOverclaimGuard.BlockedWording, "validated revalidation schedule")
		}, Point15ValDStateBlocked},
		{"blocks stale embedded vala no-overclaim disclaimer mutation", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.NoOverclaimGuard.TriggerDisclaimer = " " + point15ValATriggerDisclaimer + " "
		}, Point15ValDStateBlocked},
		{"blocks stale embedded val0 no-overclaim allowed ledger mutation", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.NoOverclaimGuard.AllowedSafeWording = append(model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.NoOverclaimGuard.AllowedSafeWording, "freshness certified")
		}, Point15ValDStateBlocked},
		{"blocks stale embedded val0 top-level freshness disclaimer overclaim", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.FreshnessDisclaimer = "continuous assurance guaranteed"
		}, Point15ValDStateBlocked},
		{"blocks stale embedded vala top-level trigger disclaimer retag", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.TriggerDisclaimer = " " + point15ValATriggerDisclaimer + " "
		}, Point15ValDStateBlocked},
		{"blocks stale embedded valb top-level revalidation disclaimer overclaim", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.Dependency.Point15ValB.RevalidationDisclaimer = "production approved"
		}, Point15ValDStateBlocked},
		{"blocks stale embedded valc top-level enforcement disclaimer retag", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.EnforcementDisclaimer = "\t" + point15ValCEnforcementDisclaimer + "\n"
		}, Point15ValDStateBlocked},
		{"blocks stale embedded val0 nested point11 dependency mutation", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.Dependency.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point11.Val0Dependency.CurrentState = Point11Val0StateBlocked
		}, Point15ValDStateBlocked},
		{"blocks stale embedded val0 point14 closure evaluator mutation", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.Dependency.Point14ValE.ClosureEvaluator.CurrentState = Point14ValEStateBlocked
		}, Point15ValDStateBlocked},
		{"blocks stale embedded val0 point14 vala chain mutation", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.Dependency.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point14ValA.CurrentState = Point14ValAStateBlocked
		}, Point15ValDStateBlocked},
		{"blocks stale embedded val0 point14 val0 chain mutation", func(model *Point15ValDDependencySnapshot) {
			model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.Dependency.Point14ValE.Dependency.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point14ValA.Dependency.Point14Val0.CurrentState = Point14Val0StateBlocked
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
		{"padded review current state blocks instead of laundering review path", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.CurrentState = " " + Point15Val0StateReviewRequired + " "
			model.EnforcementReason = point15ValCReasonStale
			model.DecisiveEvidenceVisible = false
			return model
		}, Point15ValDStateBlocked},
		{"tab newline prior state blocks raw exact timeline state", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.PriorState = "\t" + model.PriorState + "\n"
			return model
		}, Point15ValDStateBlocked},
		{"padded tenant scope blocks raw exact timeline scope", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.TenantScope = " " + model.TenantScope + " "
			return model
		}, Point15ValDStateBlocked},
		{"padded evidence id blocks raw exact timeline identity", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.EvidenceID = " " + model.EvidenceID + " "
			return model
		}, Point15ValDStateBlocked},
		{"padded downgrade reason blocks raw exact timeline reason", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.CurrentState = Point15Val0StateReviewRequired
			model.DowngradeReason = " " + point15ValATriggerStale + " "
			model.EnforcementReason = ""
			return model
		}, Point15ValDStateBlocked},
		{"timeline cannot create validity by ordering", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.TimelineCreatesValidity = true
			return model
		}, Point15ValDStateBlocked},
		{"non UTC offset event time blocks raw exact timeline", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.EventAt = "2026-05-07T10:00:00+01:00"
			return model
		}, Point15ValDStateBlocked},
		{"non UTC offset displayed time blocks raw exact timeline", func() Point15ValDAssuranceTimelineEntry {
			model := point15ValDValidTimelineModel()
			model.DisplayedAt = "2026-05-07T10:00:00+01:00"
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
		{"padded dashboard projection mode blocks raw exact discriminator", func() Point15ValDDashboardSummary {
			model := point15ValDValidDashboardModel()
			model.ProjectionMode = " " + model.ProjectionMode + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline dashboard action blocks raw exact discriminator", func() Point15ValDDashboardSummary {
			model := point15ValDValidDashboardModel()
			model.ProjectionAction = "\t" + model.ProjectionAction + "\n"
			return model
		}, Point15ValDStateBlocked},
		{"padded dashboard visibility blocks raw exact discriminator", func() Point15ValDDashboardSummary {
			model := point15ValDValidDashboardModel()
			model.Visibility = " " + model.Visibility + " "
			return model
		}, Point15ValDStateBlocked},
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
		{"padded tenant scope blocks raw exact dashboard scope", func() Point15ValDDashboardSummary {
			model := point15ValDValidDashboardModel()
			model.TenantScope = " " + model.TenantScope + " "
			return model
		}, Point15ValDStateBlocked},
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
		{"padded result state blocks raw exact query status", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.ResultState = " " + model.ResultState + " "
			return model
		}, Point15ValDStateBlocked},
		{"padded public forbidden visibility blocks instead of bypassing public guard", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.Visibility = " " + point15ValDVisibilityPublicBlocked + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline query filter blocks raw exact query discriminator", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.Filters[0] = "\t" + model.Filters[0] + "\n"
			return model
		}, Point15ValDStateBlocked},
		{"tab newline result state blocks raw exact query status", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.ResultState = "\t" + model.ResultState + "\n"
			return model
		}, Point15ValDStateBlocked},
		{"padded redaction state blocks raw exact query status", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.RedactionState = " " + model.RedactionState + " "
			return model
		}, Point15ValDStateBlocked},
		{"padded query tenant scope blocks raw exact query boundary", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.TenantScope = " " + model.TenantScope + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline query viewer scope blocks raw exact query boundary", func() Point15ValDQueryProjection {
			model := point15ValDValidQueryModel()
			model.ViewerScope = "\t" + model.ViewerScope + "\n"
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
		{"padded enforcement status blocks raw exact evidence detail state", func() Point15ValDEvidenceDetailProjection {
			model := point15ValDValidEvidenceDetailModel()
			model.EnforcementStatus = " " + model.EnforcementStatus + " "
			return model
		}, Point15ValDStateBlocked},
		{"padded freshness status blocks raw exact evidence detail status", func() Point15ValDEvidenceDetailProjection {
			model := point15ValDValidEvidenceDetailModel()
			model.FreshnessStatus = " " + model.FreshnessStatus + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline lifecycle status blocks raw exact evidence detail status", func() Point15ValDEvidenceDetailProjection {
			model := point15ValDValidEvidenceDetailModel()
			model.LifecycleStatus = "\t" + model.LifecycleStatus + "\n"
			return model
		}, Point15ValDStateBlocked},
		{"padded tenant scope blocks raw exact evidence detail boundary", func() Point15ValDEvidenceDetailProjection {
			model := point15ValDValidEvidenceDetailModel()
			model.TenantScope = " " + model.TenantScope + " "
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
		{"padded scheduled status blocks raw exact revalidation detail", func() Point15ValDRevalidationDetailProjection {
			model := point15ValDValidRevalidationDetailModel()
			model.ScheduledStatus = " " + model.ScheduledStatus + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline run result blocks raw exact revalidation detail", func() Point15ValDRevalidationDetailProjection {
			model := point15ValDValidRevalidationDetailModel()
			model.RunResult = "\t" + model.RunResult + "\n"
			return model
		}, Point15ValDStateBlocked},
		{"padded retry and throttle statuses block raw exact revalidation detail", func() Point15ValDRevalidationDetailProjection {
			model := point15ValDValidRevalidationDetailModel()
			model.RetryStatus = " " + model.RetryStatus + " "
			model.ThrottleStatus = " " + model.ThrottleStatus + " "
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
		{"padded target state blocks raw exact enforcement detail state", func() Point15ValDEnforcementDetailProjection {
			model := point15ValDValidEnforcementDetailModel()
			model.TargetState = " " + model.TargetState + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline prior state blocks raw exact enforcement detail state", func() Point15ValDEnforcementDetailProjection {
			model := point15ValDValidEnforcementDetailModel()
			model.PriorState = "\t" + model.PriorState + "\n"
			return model
		}, Point15ValDStateBlocked},
		{"padded current state blocks raw exact enforcement detail state", func() Point15ValDEnforcementDetailProjection {
			model := point15ValDValidEnforcementDetailModel()
			model.CurrentState = " " + model.CurrentState + " "
			return model
		}, Point15ValDStateBlocked},
		{"padded enforcement action blocks raw exact enforcement detail action", func() Point15ValDEnforcementDetailProjection {
			model := point15ValDValidEnforcementDetailModel()
			model.EnforcementAction = " " + model.EnforcementAction + " "
			return model
		}, Point15ValDStateBlocked},
		{"padded enforcement reason blocks raw exact enforcement detail reason", func() Point15ValDEnforcementDetailProjection {
			model := point15ValDValidEnforcementDetailModel()
			model.EnforcementAction = point15ValCActionReview
			model.EnforcementReason = " " + point15ValCReasonStale + " "
			model.TargetState = Point15Val0StateReviewRequired
			model.CurrentState = Point15Val0StateReviewRequired
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
		{"padded replay ref blocks raw exact replay binding", func() Point15ValDReplayProofHistoryProjection {
			model := point15ValDValidReplayHistoryModel()
			model.ReplayRef = " " + model.ReplayRef + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline proof pack ref blocks raw exact replay binding", func() Point15ValDReplayProofHistoryProjection {
			model := point15ValDValidReplayHistoryModel()
			model.ProofPackRef = "\t" + model.ProofPackRef + "\n"
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
		{"padded public visibility blocks raw exact access boundary", func() Point15ValDAccessTenantPrivacyBoundary {
			model := point15ValDValidAccessTenantModel()
			model.Visibility = " " + point15ValDVisibilityPublicBlocked + " "
			return model
		}, Point15ValDStateBlocked},
		{"redaction cannot hide decisive failure without review", func() Point15ValDAccessTenantPrivacyBoundary {
			model := point15ValDValidAccessTenantModel()
			model.DecisiveFailureHidden = true
			model.RedactionState = point15ValDRedactionLimited
			return model
		}, Point15ValDStateReviewRequired},
		{"padded access redaction state blocks raw exact boundary", func() Point15ValDAccessTenantPrivacyBoundary {
			model := point15ValDValidAccessTenantModel()
			model.RedactionState = " " + model.RedactionState + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline access redaction state blocks raw exact boundary", func() Point15ValDAccessTenantPrivacyBoundary {
			model := point15ValDValidAccessTenantModel()
			model.RedactionState = "\t" + model.RedactionState + "\n"
			return model
		}, Point15ValDStateBlocked},
		{"padded access tenant scope blocks raw exact boundary", func() Point15ValDAccessTenantPrivacyBoundary {
			model := point15ValDValidAccessTenantModel()
			model.TenantScope = " " + model.TenantScope + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline access viewer scope blocks raw exact boundary", func() Point15ValDAccessTenantPrivacyBoundary {
			model := point15ValDValidAccessTenantModel()
			model.ViewerScope = "\t" + model.ViewerScope + "\n"
			return model
		}, Point15ValDStateBlocked},
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
		{"padded timestamp projection mode blocks raw exact discriminator", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.ProjectionMode = " " + model.ProjectionMode + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline timestamp projection mode blocks raw exact discriminator", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.ProjectionMode = "\t" + model.ProjectionMode + "\n"
			return model
		}, Point15ValDStateBlocked},
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
		{"non UTC offset reference time blocks raw exact timestamp display", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.ReferenceNow = "2026-05-07T12:00:00+01:00"
			return model
		}, Point15ValDStateBlocked},
		{"non UTC offset displayed time blocks raw exact timestamp display", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.DisplayedAt = "2026-05-07T12:00:00+01:00"
			return model
		}, Point15ValDStateBlocked},
		{"non UTC offset source event time blocks raw exact timestamp display", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.SourceEventAt = "2026-05-07T08:00:00+01:00"
			return model
		}, Point15ValDStateBlocked},
		{"non UTC offset validated time blocks raw exact timestamp display", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.ValidatedAt = "2026-05-07T09:00:00+01:00"
			model.EnforcedAt = "2026-05-07T10:00:00Z"
			return model
		}, Point15ValDStateBlocked},
		{"non UTC offset enforced time blocks raw exact timestamp display", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.ValidatedAt = "2026-05-07T09:00:00Z"
			model.EnforcedAt = "2026-05-07T10:00:00+01:00"
			return model
		}, Point15ValDStateBlocked},
		{"non UTC offset received time blocks raw exact timestamp display", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.ReceivedAt = "2026-05-07T08:00:00+01:00"
			model.ValidatedAt = "2026-05-07T09:00:00Z"
			return model
		}, Point15ValDStateBlocked},
		{"padded tenant scope blocks raw exact timestamp display boundary", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.TenantScope = " " + model.TenantScope + " "
			return model
		}, Point15ValDStateBlocked},
		{"whitespace only optional source event blocks instead of becoming absent", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.SourceEventAt = " \t\n"
			return model
		}, Point15ValDStateBlocked},
		{"whitespace only optional validated time blocks instead of becoming absent", func() Point15ValDTimestampDisplayDiscipline {
			model := point15ValDValidTimestampDisplayModel()
			model.ValidatedAt = " \t\n"
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
		{"padded authority tenant scope blocks raw exact boundary", func() Point15ValDAuthorityBoundary {
			model := point15ValDValidAuthorityModel()
			model.TenantScope = " " + model.TenantScope + " "
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
		{"inherited point10 readiness overclaim wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"marketplace production ready"}
			return model
		}, Point15ValDStateBlocked},
		{"unhyphenated inherited readiness overclaim wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"air gapped certified"}
			return model
		}, Point15ValDStateBlocked},
		{"unicode dash forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"deployment\u2011approved"}
			return model
		}, Point15ValDStateBlocked},
		{"slash separator forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"deployment/approved"}
			return model
		}, Point15ValDStateBlocked},
		{"dot separator forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"production.approved"}
			return model
		}, Point15ValDStateBlocked},
		{"underscore separator forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"production_approved"}
			return model
		}, Point15ValDStateBlocked},
		{"intra bucket filler forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"production is approved"}
			return model
		}, Point15ValDStateBlocked},
		{"confusable forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"de\u1d18loyment approved"}
			return model
		}, Point15ValDStateBlocked},
		{"zero width forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"deployment appro\u200dved"}
			return model
		}, Point15ValDStateBlocked},
		{"zero width separator forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"production\u200bapproved"}
			return model
		}, Point15ValDStateBlocked},
		{"math bold forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"\U0001d429\U0001d42b\U0001d428\U0001d41d\U0001d42e\U0001d41c\U0001d42d\U0001d422\U0001d428\U0001d427 \U0001d41a\U0001d429\U0001d429\U0001d42b\U0001d428\U0001d42f\U0001d41e\U0001d41d"}
			return model
		}, Point15ValDStateBlocked},
		{"greek nu forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"deployment appro\u03bded"}
			return model
		}, Point15ValDStateBlocked},
		{"greek upsilon forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"prod\u03c5ction approved"}
			return model
		}, Point15ValDStateBlocked},
		{"small cap u forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"prod\U00001d1cction approved"}
			return model
		}, Point15ValDStateBlocked},
		{"latin upsilon forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"prod\u028action approved"}
			return model
		}, Point15ValDStateBlocked},
		{"greek delta forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"production approve\u03b4"}
			return model
		}, Point15ValDStateBlocked},
		{"small cap t forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"global \U00001d1bruth"}
			return model
		}, Point15ValDStateBlocked},
		{"latin alpha forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"glob\u0251l truth"}
			return model
		}, Point15ValDStateBlocked},
		{"latin iota forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"off\u0269cial authority"}
			return model
		}, Point15ValDStateBlocked},
		{"dental click forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"comp\u01c0iance guaranteed"}
			return model
		}, Point15ValDStateBlocked},
		{"armenian oh forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"c\u0585mpliance guaranteed"}
			return model
		}, Point15ValDStateBlocked},
		{"greek eta forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"productio\u03b7 approved"}
			return model
		}, Point15ValDStateBlocked},
		{"latin eng forbidden wording blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"productio\u014b approved"}
			return model
		}, Point15ValDStateBlocked},
		{"internal underscore machine token remains non-boundary safe wording", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.InternalDiagnosticTexts = append(model.InternalDiagnosticTexts, "internal_production_approved_metric")
			return model
		}, Point15ValDStateActive},
		{"split forbidden wording across observed corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"deployment", "approved"}
			return model
		}, Point15ValDStateBlocked},
		{"split unhyphenated inherited readiness overclaim blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"air gapped", "certified"}
			return model
		}, Point15ValDStateBlocked},
		{"tab newline forbidden wording blocks after normalized token scan", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"deployment\tapproved"}
			return model
		}, Point15ValDStateBlocked},
		{"unicode dash split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"production\u2013", "approved"}
			return model
		}, Point15ValDStateBlocked},
		{"slash split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"production/", "approved"}
			return model
		}, Point15ValDStateBlocked},
		{"zero width split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"production", "appro\u2060ved"}
			return model
		}, Point15ValDStateBlocked},
		{"word fragment split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"produc", "tion approved"}
			return model
		}, Point15ValDStateBlocked},
		{"right leg u split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"prod\uab4e", "ction approved"}
			return model
		}, Point15ValDStateBlocked},
		{"latin upsilon split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"prod\u028a", "ction approved"}
			return model
		}, Point15ValDStateBlocked},
		{"greek nu split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"production", "appro\u03bded"}
			return model
		}, Point15ValDStateBlocked},
		{"greek delta split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"production", "approve\u03b4"}
			return model
		}, Point15ValDStateBlocked},
		{"small cap t split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"global", "\U00001d1bruth"}
			return model
		}, Point15ValDStateBlocked},
		{"latin alpha split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"glob\u0251l", "truth"}
			return model
		}, Point15ValDStateBlocked},
		{"latin iota split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"off\u0269cial", "authority"}
			return model
		}, Point15ValDStateBlocked},
		{"dental click split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"comp\u01c0iance", "guaranteed"}
			return model
		}, Point15ValDStateBlocked},
		{"armenian oh split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"c\u0585mpliance", "guaranteed"}
			return model
		}, Point15ValDStateBlocked},
		{"armenian vo split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"productio\u0578", "approved"}
			return model
		}, Point15ValDStateBlocked},
		{"latin n with long right leg split forbidden wording across corpus blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ObservedTexts = []string{"productio\u019e", "approved"}
			return model
		}, Point15ValDStateBlocked},
		{"padded projection disclaimer blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ProjectionDisclaimer = " " + model.ProjectionDisclaimer + " "
			return model
		}, Point15ValDStateBlocked},
		{"tab newline projection disclaimer blocks", func() Point15ValDNoOverclaimGuard {
			model := point15ValDValidNoOverclaimModel()
			model.ProjectionDisclaimer = "\t" + model.ProjectionDisclaimer + "\n"
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
		{"padded dashboard projection mode blocks foundation before vale", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.Dashboard.ProjectionMode = " " + model.Dashboard.ProjectionMode + " "
		}, Point15ValDStateBlocked},
		{"tab newline query projection action blocks foundation before vale", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.Query.ProjectionAction = "\t" + model.Query.ProjectionAction + "\n"
		}, Point15ValDStateBlocked},
		{"padded evidence detail id blocks raw exact projection", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.EvidenceDetail.EvidenceID = " " + model.EvidenceDetail.EvidenceID + " "
		}, Point15ValDStateBlocked},
		{"tab newline tenant scope blocks raw exact projection", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.Timeline.TenantScope = "\t" + model.Timeline.TenantScope + "\n"
		}, Point15ValDStateBlocked},
		{"padded timestamp projection mode blocks foundation before vale", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.TimestampDisplayDiscipline.ProjectionMode = " " + model.TimestampDisplayDiscipline.ProjectionMode + " "
		}, Point15ValDStateBlocked},
		{"split no overclaim wording blocks full ValD foundation", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.NoOverclaimGuard.ObservedTexts = []string{"deployment", "approved"}
		}, Point15ValDStateBlocked},
		{"top-level projection disclaimer overclaim blocks full ValD foundation", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.ProjectionDisclaimer = "production approved"
		}, Point15ValDStateBlocked},
		{"padded top-level projection disclaimer blocks full ValD foundation", func(model *Point15ValDAssuranceProjectionFoundation) {
			model.ProjectionDisclaimer = " " + point15ValDProjectionDisclaimer + " "
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

	t.Run("split no overclaim wording records exact blocking reason", func(t *testing.T) {
		model := point15ValDValidFoundationModel()
		model.NoOverclaimGuard.ObservedTexts = []string{"deployment", "approved"}
		computed := ComputePoint15ValDAssuranceProjectionFoundation(model)
		if computed.CurrentState != Point15ValDStateBlocked {
			t.Fatalf("expected split no-overclaim wording to block, got %#v", computed)
		}
		if !point15ValDStringSliceContains(computed.BlockingReasons, "no_overclaim") {
			t.Fatalf("expected exact no_overclaim blocking reason, got %#v", computed.BlockingReasons)
		}
	})

	t.Run("stale embedded valc timestamp mutation records exact dependency reason", func(t *testing.T) {
		model := point15ValDValidFoundationModel()
		model.Dependency.Point15ValC.TimestampDiscipline.ReceivedAt = "2026-05-07T09:06:00Z"
		computed := ComputePoint15ValDAssuranceProjectionFoundation(model)
		if computed.CurrentState != Point15ValDStateBlocked {
			t.Fatalf("expected stale embedded ValC timestamp mutation to block, got %#v", computed)
		}
		if !point15ValDStringSliceContains(computed.BlockingReasons, "dependency") {
			t.Fatalf("expected exact dependency blocking reason, got %#v", computed.BlockingReasons)
		}
	})

	t.Run("stale embedded valc no overclaim mutation records exact dependency reason", func(t *testing.T) {
		model := point15ValDValidFoundationModel()
		model.Dependency.Point15ValC.NoOverclaimGuard.AllowedSafeWording = append(model.Dependency.Point15ValC.NoOverclaimGuard.AllowedSafeWording, "production approved")
		computed := ComputePoint15ValDAssuranceProjectionFoundation(model)
		if computed.CurrentState != Point15ValDStateBlocked {
			t.Fatalf("expected stale embedded ValC no-overclaim mutation to block, got %#v", computed)
		}
		if !point15ValDStringSliceContains(computed.BlockingReasons, "dependency") {
			t.Fatalf("expected exact dependency blocking reason, got %#v", computed.BlockingReasons)
		}
	})

	t.Run("stale embedded valc split no overclaim wording records exact dependency reason", func(t *testing.T) {
		model := point15ValDValidFoundationModel()
		model.Dependency.Point15ValC.NoOverclaimGuard.ObservedTexts = []string{"continuous assurance", "guaranteed"}
		computed := ComputePoint15ValDAssuranceProjectionFoundation(model)
		if computed.CurrentState != Point15ValDStateBlocked {
			t.Fatalf("expected stale embedded ValC split no-overclaim wording to block, got %#v", computed)
		}
		if !point15ValDStringSliceContains(computed.BlockingReasons, "dependency") {
			t.Fatalf("expected exact dependency blocking reason, got %#v", computed.BlockingReasons)
		}
	})
}

func TestPoint15ValDAggregateRawExact(t *testing.T) {
	tests := []struct {
		name   string
		states []string
		want   string
	}{
		{
			name:   "happy path exact active states remain active",
			states: []string{Point15ValDStateActive, Point15ValDStateActive},
			want:   Point15ValDStateActive,
		},
		{
			name:   "direct exploit padded active aggregate state blocks",
			states: []string{" " + Point15ValDStateActive + " ", Point15ValDStateActive},
			want:   Point15ValDStateBlocked,
		},
		{
			name:   "hard invalid tab newline active aggregate state blocks",
			states: []string{Point15ValDStateActive, "\t" + Point15ValDStateActive + "\n"},
			want:   Point15ValDStateBlocked,
		},
		{
			name:   "unknown aggregate state fails closed",
			states: []string{Point15ValDStateActive, "point15_vald_unknown_active"},
			want:   Point15ValDStateBlocked,
		},
		{
			name:   "exact review state preserves review precedence",
			states: []string{Point15ValDStateActive, Point15ValDStateReviewRequired},
			want:   Point15ValDStateReviewRequired,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := point15ValDAggregate(tc.states...); got != tc.want {
				t.Fatalf("expected aggregate %s, got %s", tc.want, got)
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

package formal

import (
	"sync"
	"testing"
)

var (
	point15ValBFoundationOnce sync.Once
	point15ValBFoundationBase Point15ValBScheduledRevalidationFoundation
)

func point15ValBCloneStrings(values []string) []string {
	return append([]string(nil), values...)
}

func clonePoint15ValBFoundation(model Point15ValBScheduledRevalidationFoundation) Point15ValBScheduledRevalidationFoundation {
	model.BlockingReasons = point15ValBCloneStrings(model.BlockingReasons)
	model.ReviewPrerequisites = point15ValBCloneStrings(model.ReviewPrerequisites)
	model.Dependency.ReviewPrerequisites = point15ValBCloneStrings(model.Dependency.ReviewPrerequisites)
	model.Dependency.Point15ValA = clonePoint15ValAFoundation(model.Dependency.Point15ValA)
	model.NoOverclaimGuard.ObservedTexts = point15ValBCloneStrings(model.NoOverclaimGuard.ObservedTexts)
	model.NoOverclaimGuard.InternalDiagnosticTexts = point15ValBCloneStrings(model.NoOverclaimGuard.InternalDiagnosticTexts)
	model.NoOverclaimGuard.AllowedSafeWording = point15ValBCloneStrings(model.NoOverclaimGuard.AllowedSafeWording)
	model.NoOverclaimGuard.BlockedWording = point15ValBCloneStrings(model.NoOverclaimGuard.BlockedWording)
	return model
}

func point15ValBValidFoundationModel() Point15ValBScheduledRevalidationFoundation {
	point15ValBFoundationOnce.Do(func() {
		point15ValBFoundationBase = Point15ValBFoundationModel()
	})
	return clonePoint15ValBFoundation(point15ValBFoundationBase)
}

func point15ValBValidDependencyModel() Point15ValBDependencySnapshot {
	return point15ValBValidFoundationModel().Dependency
}

func point15ValBValidScheduleModel() Point15ValBRevalidationSchedule {
	return point15ValBValidFoundationModel().Schedule
}

func point15ValBValidRunModel() Point15ValBRevalidationRun {
	return point15ValBValidFoundationModel().Run
}

func point15ValBValidRetryBudgetModel() Point15ValBRetryBudget {
	return point15ValBValidFoundationModel().RetryBudget
}

func point15ValBValidThrottleModel() Point15ValBTenantThrottle {
	return point15ValBValidFoundationModel().TenantThrottle
}

func point15ValBValidBindingModel() Point15ValBDowngradeBinding {
	return point15ValBValidFoundationModel().DowngradeBinding
}

func point15ValBValidTimestampModel() Point15ValBTimestampDiscipline {
	return point15ValBValidFoundationModel().TimestampDiscipline
}

func point15ValBValidAuthorityBoundaryModel() Point15ValBAuthorityBoundary {
	return point15ValBValidFoundationModel().AuthorityBoundary
}

func point15ValBValidNoOverclaimGuardModel() Point15ValBNoOverclaimGuard {
	return point15ValBValidFoundationModel().NoOverclaimGuard
}

func point15ValBCompletedCleanFoundationModel() Point15ValBScheduledRevalidationFoundation {
	model := point15ValBValidFoundationModel()
	model.Schedule.ScheduledStatus = point15ValBScheduleCompleted
	model.Schedule.LastCompletedAt = "2026-05-07T09:20:00Z"
	model.Run.RunID = "run_point15_valb_001"
	model.Run.RunResult = point15ValBRunCompletedClean
	model.Run.StartedAt = "2026-05-07T09:10:00Z"
	model.Run.CompletedAt = "2026-05-07T09:20:00Z"
	model.Run.RunEvidenceHash = model.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.EvidenceHash
	model.Run.ConnectorResultRef = "connector_result_point15_valb_001"
	model.Run.PolicyVersion = model.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.PolicyVersion
	model.Run.EngineVersion = model.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.EngineVersion
	model.Run.SchemaVersion = model.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.SchemaVersion
	model.DowngradeBinding = point15ValBDowngradeBindingModel(model.Schedule, model.Run, model.RetryBudget, model.TenantThrottle)
	model.DowngradeBinding.RunRef = model.Run.RunID
	model.TimestampDiscipline.ScheduledStatus = model.Schedule.ScheduledStatus
	model.TimestampDiscipline.StartedAt = model.Run.StartedAt
	model.TimestampDiscipline.StartedAtTimeSource = point14Val0TimeSourceServerUTC
	model.TimestampDiscipline.CompletedAt = model.Run.CompletedAt
	model.TimestampDiscipline.CompletedAtTimeSource = point14Val0TimeSourceServerUTC
	return model
}

func TestPoint15ValBDependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValBDependencySnapshot)
		want   string
	}{
		{"active when vala clean", func(*Point15ValBDependencySnapshot) {}, Point15ValBStateActive},
		{"blocks when vala missing", func(model *Point15ValBDependencySnapshot) {
			model.Point15ValACurrentState = ""
		}, Point15ValBStateBlocked},
		{"blocks when vala blocked", func(model *Point15ValBDependencySnapshot) {
			model.Point15ValACurrentState = Point15ValAStateBlocked
		}, Point15ValBStateBlocked},
		{"blocks when vala review required", func(model *Point15ValBDependencySnapshot) {
			model.Point15ValACurrentState = Point15ValAStateReviewRequired
		}, Point15ValBStateBlocked},
		{"blocks when vala incomplete", func(model *Point15ValBDependencySnapshot) {
			model.Point15ValACurrentState = Point15ValAStateIncomplete
		}, Point15ValBStateBlocked},
		{"blocks when vala not merged", func(model *Point15ValBDependencySnapshot) {
			model.Point15ValAMerged = false
		}, Point15ValBStateBlocked},
		{"blocks when vala ci not green", func(model *Point15ValBDependencySnapshot) {
			model.Point15ValACIGreen = false
		}, Point15ValBStateBlocked},
		{"blocks when vala not reviewed on main", func(model *Point15ValBDependencySnapshot) {
			model.Point15ValAReviewedOnMain = false
		}, Point15ValBStateBlocked},
		{"blocks on point15 pass token", func(model *Point15ValBDependencySnapshot) {
			model.Point15PassSeen = true
		}, Point15ValBStateBlocked},
		{"blocks when computed provenance mismatches upstream", func(model *Point15ValBDependencySnapshot) {
			model.Point15ValAComputedFromUpstream = false
		}, Point15ValBStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValBValidDependencyModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValBDependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValBRevalidationScheduleState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValBRevalidationSchedule)
		want   string
	}{
		{"scheduled active before due", func(*Point15ValBRevalidationSchedule) {}, Point15ValBStateActive},
		{"not_required active only when not required", func(model *Point15ValBRevalidationSchedule) {
			model.Required = false
			model.ScheduledStatus = point15ValBScheduleNotRequired
			model.RevalidationDueAt = ""
			model.ScheduledAt = ""
			model.SchedulerTimeSource = ""
		}, Point15ValBStateActive},
		{"due is review required", func(model *Point15ValBRevalidationSchedule) {
			model.ScheduledStatus = point15ValBScheduleDue
		}, Point15ValBStateReviewRequired},
		{"overdue is review required", func(model *Point15ValBRevalidationSchedule) {
			model.ScheduledStatus = point15ValBScheduleOverdue
		}, Point15ValBStateReviewRequired},
		{"missed is review required", func(model *Point15ValBRevalidationSchedule) {
			model.ScheduledStatus = point15ValBScheduleMissed
		}, Point15ValBStateReviewRequired},
		{"failed maps to review required", func(model *Point15ValBRevalidationSchedule) {
			model.ScheduledStatus = point15ValBScheduleFailed
		}, Point15ValBStateReviewRequired},
		{"retry pending cannot hide downgrade", func(model *Point15ValBRevalidationSchedule) {
			model.ScheduledStatus = point15ValBScheduleRetryPending
		}, Point15ValBStateReviewRequired},
		{"retry exhausted blocks", func(model *Point15ValBRevalidationSchedule) {
			model.ScheduledStatus = point15ValBScheduleRetryExhausted
		}, Point15ValBStateBlocked},
		{"throttled review required", func(model *Point15ValBRevalidationSchedule) {
			model.ScheduledStatus = point15ValBScheduleThrottled
		}, Point15ValBStateReviewRequired},
		{"unknown status blocks", func(model *Point15ValBRevalidationSchedule) {
			model.ScheduledStatus = "later"
		}, Point15ValBStateBlocked},
		{"required schedule missing due time incomplete", func(model *Point15ValBRevalidationSchedule) {
			model.RevalidationDueAt = ""
		}, Point15ValBStateIncomplete},
		{"not required with scheduled status blocks", func(model *Point15ValBRevalidationSchedule) {
			model.Required = false
		}, Point15ValBStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValBValidScheduleModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValBRevalidationScheduleState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValBRevalidationRunState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValBRevalidationRun
		want  string
	}{
		{"not_run is active", func() Point15ValBRevalidationRun {
			return point15ValBValidRunModel()
		}, Point15ValBStateActive},
		{"completed_clean active when fully bound", func() Point15ValBRevalidationRun {
			return point15ValBCompletedCleanFoundationModel().Run
		}, Point15ValBStateActive},
		{"completed_with_downgrade review required", func() Point15ValBRevalidationRun {
			model := point15ValBCompletedCleanFoundationModel().Run
			model.RunResult = point15ValBRunCompletedWithDowngrade
			model.DowngradeTriggerRef = "trigger_point15_vala_001"
			return model
		}, Point15ValBStateReviewRequired},
		{"failed review required", func() Point15ValBRevalidationRun {
			model := point15ValBCompletedCleanFoundationModel().Run
			model.RunResult = point15ValBRunFailed
			return model
		}, Point15ValBStateReviewRequired},
		{"missed review required", func() Point15ValBRevalidationRun {
			model := point15ValBCompletedCleanFoundationModel().Run
			model.RunResult = point15ValBRunMissed
			return model
		}, Point15ValBStateReviewRequired},
		{"unauthorized blocks", func() Point15ValBRevalidationRun {
			model := point15ValBCompletedCleanFoundationModel().Run
			model.RunResult = point15ValBRunUnauthorized
			return model
		}, Point15ValBStateBlocked},
		{"tenant mismatch blocks", func() Point15ValBRevalidationRun {
			model := point15ValBCompletedCleanFoundationModel().Run
			model.RunResult = point15ValBRunTenantMismatch
			return model
		}, Point15ValBStateBlocked},
		{"timeout review required", func() Point15ValBRevalidationRun {
			model := point15ValBCompletedCleanFoundationModel().Run
			model.RunResult = point15ValBRunTimeout
			return model
		}, Point15ValBStateReviewRequired},
		{"tampered blocks", func() Point15ValBRevalidationRun {
			model := point15ValBCompletedCleanFoundationModel().Run
			model.RunResult = point15ValBRunTampered
			return model
		}, Point15ValBStateBlocked},
		{"missing run evidence incomplete", func() Point15ValBRevalidationRun {
			model := point15ValBCompletedCleanFoundationModel().Run
			model.RunEvidenceHash = ""
			return model
		}, Point15ValBStateIncomplete},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValBRevalidationRunState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValBRetryBudgetState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValBRetryBudget)
		want   string
	}{
		{"available active", func(*Point15ValBRetryBudget) {}, Point15ValBStateActive},
		{"exhausted manual review review required", func(model *Point15ValBRetryBudget) {
			model.AttemptsUsed = model.MaxRetries
			model.RetryBudgetStatus = point15ValBRetryExhausted
			model.RetryReason = point15ValBRetryReasonManualReview
		}, Point15ValBStateReviewRequired},
		{"exhausted terminal blocks", func(model *Point15ValBRetryBudget) {
			model.AttemptsUsed = model.MaxRetries
			model.RetryBudgetStatus = point15ValBRetryExhausted
			model.RetryReason = point15ValBRetryReasonTerminal
		}, Point15ValBStateBlocked},
		{"self reset blocks", func(model *Point15ValBRetryBudget) {
			model.SelfResetDetected = true
		}, Point15ValBStateBlocked},
		{"next retry requires trusted time source", func(model *Point15ValBRetryBudget) {
			model.NextRetryAt = "2026-05-07T10:15:00Z"
			model.NextRetryTimeSource = point14Val0TimeSourceClientLocal
		}, Point15ValBStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValBValidRetryBudgetModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValBRetryBudgetState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValBTenantThrottleState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValBTenantThrottle)
		want   string
	}{
		{"within limit active", func(*Point15ValBTenantThrottle) {}, Point15ValBStateActive},
		{"throttled review required", func(model *Point15ValBTenantThrottle) {
			model.RequestedRevalidations = 20
			model.ThrottleStatus = point15ValBThrottleReviewRequired
		}, Point15ValBStateReviewRequired},
		{"throttled blocked", func(model *Point15ValBTenantThrottle) {
			model.RequestedRevalidations = 20
			model.ThrottleStatus = point15ValBThrottleBlocked
		}, Point15ValBStateBlocked},
		{"cross tenant blocked", func(model *Point15ValBTenantThrottle) {
			model.ThrottleStatus = point15ValBThrottleCrossTenantBlocked
			model.CrossTenantDetected = true
		}, Point15ValBStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValBValidThrottleModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValBTenantThrottleState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValBDowngradeBindingState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValBDowngradeBinding
		want  string
	}{
		{"no trigger active", func() Point15ValBDowngradeBinding {
			return point15ValBValidBindingModel()
		}, Point15ValBStateActive},
		{"completed clean without run ref blocks", func() Point15ValBDowngradeBinding {
			model := point15ValBCompletedCleanFoundationModel().DowngradeBinding
			model.RunRef = ""
			return model
		}, Point15ValBStateBlocked},
		{"overdue maps to stale review when completion exists", func() Point15ValBDowngradeBinding {
			model := point15ValBValidBindingModel()
			model.ScheduleStatus = point15ValBScheduleOverdue
			model.LastCompletedAt = "2026-05-07T08:00:00Z"
			model.TriggerType = point15ValATriggerStale
			model.TargetState = Point15Val0StateReviewRequired
			model.TargetDowngradeOutcome = point15Val0DowngradeReview
			model.RetainsActiveClosure = false
			return model
		}, Point15ValBStateReviewRequired},
		{"overdue maps to missing freshness proof when no completion exists", func() Point15ValBDowngradeBinding {
			model := point15ValBValidBindingModel()
			model.ScheduleStatus = point15ValBScheduleOverdue
			model.LastCompletedAt = ""
			model.TriggerType = point15ValATriggerMissing
			model.TargetState = Point15Val0StateIncomplete
			model.TargetDowngradeOutcome = point15Val0DowngradeIncomplete
			model.RetainsActiveClosure = false
			return model
		}, Point15ValBStateIncomplete},
		{"connector timeout maps to connector timeout", func() Point15ValBDowngradeBinding {
			model := point15ValBValidBindingModel()
			model.ScheduleStatus = point15ValBScheduleCompleted
			model.RunResult = point15ValBRunTimeout
			model.TriggerType = point15ValATriggerConnTimeout
			model.TargetState = Point15Val0StateReviewRequired
			model.TargetDowngradeOutcome = point15Val0DowngradeReview
			model.RetainsActiveClosure = false
			model.RunRef = "run_point15_valb_001"
			return model
		}, Point15ValBStateReviewRequired},
		{"connector unauthorized blocks", func() Point15ValBDowngradeBinding {
			model := point15ValBValidBindingModel()
			model.ScheduleStatus = point15ValBScheduleCompleted
			model.RunResult = point15ValBRunUnauthorized
			model.TriggerType = point15ValATriggerConnAuth
			model.TargetState = Point15Val0StateBlocked
			model.TargetDowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
			model.RunRef = "run_point15_valb_001"
			return model
		}, Point15ValBStateBlocked},
		{"connector tenant mismatch blocks", func() Point15ValBDowngradeBinding {
			model := point15ValBValidBindingModel()
			model.ScheduleStatus = point15ValBScheduleCompleted
			model.RunResult = point15ValBRunTenantMismatch
			model.TriggerType = point15ValATriggerConnTenant
			model.TargetState = Point15Val0StateBlocked
			model.TargetDowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
			model.RunRef = "run_point15_valb_001"
			return model
		}, Point15ValBStateBlocked},
		{"hash mismatch blocks", func() Point15ValBDowngradeBinding {
			model := point15ValBValidBindingModel()
			model.ScheduleStatus = point15ValBScheduleCompleted
			model.RunResult = point15ValBRunCompletedClean
			model.RunEvidenceHashMatches = false
			model.TriggerType = point15ValATriggerHash
			model.TargetState = Point15Val0StateBlocked
			model.TargetDowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
			model.RunRef = "run_point15_valb_001"
			return model
		}, Point15ValBStateBlocked},
		{"tampered run blocks", func() Point15ValBDowngradeBinding {
			model := point15ValBValidBindingModel()
			model.ScheduleStatus = point15ValBScheduleCompleted
			model.RunResult = point15ValBRunTampered
			model.TriggerType = point15ValATriggerTampered
			model.TargetState = Point15Val0StateBlocked
			model.TargetDowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
			model.RunRef = "run_point15_valb_001"
			return model
		}, Point15ValBStateBlocked},
		{"retains pass with trigger blocks", func() Point15ValBDowngradeBinding {
			model := point15ValBValidBindingModel()
			model.ScheduleStatus = point15ValBScheduleOverdue
			model.LastCompletedAt = "2026-05-07T08:00:00Z"
			model.TriggerType = point15ValATriggerStale
			model.TargetState = Point15Val0StateReviewRequired
			model.TargetDowngradeOutcome = point15Val0DowngradeReview
			model.RetainsPass = true
			model.RetainsActiveClosure = false
			return model
		}, Point15ValBStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvaluatePoint15ValBDowngradeBindingState(tc.model()); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValBTimestampDisciplineState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValBTimestampDiscipline)
		want   string
	}{
		{"scheduled active before due", func(*Point15ValBTimestampDiscipline) {}, Point15ValBStateActive},
		{"client local time blocks", func(model *Point15ValBTimestampDiscipline) {
			model.ClientLocalCreatesCanonical = true
		}, Point15ValBStateBlocked},
		{"source event creates canonical blocks", func(model *Point15ValBTimestampDiscipline) {
			model.SourceEventCreatesCanonical = true
		}, Point15ValBStateBlocked},
		{"scheduled after due time blocks", func(model *Point15ValBTimestampDiscipline) {
			model.DueAt = "2026-05-07T09:00:00Z"
			model.ReferenceNow = "2026-05-07T09:30:00Z"
		}, Point15ValBStateBlocked},
		{"future completed_at blocks", func(model *Point15ValBTimestampDiscipline) {
			model.ScheduledStatus = point15ValBScheduleCompleted
			model.StartedAt = "2026-05-07T09:10:00Z"
			model.StartedAtTimeSource = point14Val0TimeSourceServerUTC
			model.CompletedAt = "2026-05-07T10:30:00Z"
			model.CompletedAtTimeSource = point14Val0TimeSourceServerUTC
		}, Point15ValBStateBlocked},
		{"completed before started blocks", func(model *Point15ValBTimestampDiscipline) {
			model.ScheduledStatus = point15ValBScheduleCompleted
			model.StartedAt = "2026-05-07T09:10:00Z"
			model.StartedAtTimeSource = point14Val0TimeSourceServerUTC
			model.CompletedAt = "2026-05-07T09:05:00Z"
			model.CompletedAtTimeSource = point14Val0TimeSourceServerUTC
		}, Point15ValBStateBlocked},
		{"due after next retry blocks", func(model *Point15ValBTimestampDiscipline) {
			model.NextRetryAt = "2026-05-07T09:45:00Z"
			model.NextRetryAtTimeSource = point14Val0TimeSourceServerUTC
			model.DueAt = "2026-05-07T10:00:00Z"
		}, Point15ValBStateBlocked},
		{"missing due_at incomplete", func(model *Point15ValBTimestampDiscipline) {
			model.DueAt = ""
		}, Point15ValBStateIncomplete},
		{"backdated completion review required", func(model *Point15ValBTimestampDiscipline) {
			model.ScheduledStatus = point15ValBScheduleCompleted
			model.StartedAt = "2026-05-07T08:50:00Z"
			model.StartedAtTimeSource = point14Val0TimeSourceServerUTC
			model.CompletedAt = "2026-05-07T08:55:00Z"
			model.CompletedAtTimeSource = point14Val0TimeSourceServerUTC
		}, Point15ValBStateReviewRequired},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValBValidTimestampModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValBTimestampDisciplineState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValBAuthorityBoundaryState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValBAuthorityBoundary)
	}{
		{"scheduler cannot mark evidence fresh", func(model *Point15ValBAuthorityBoundary) {
			model.SchedulerMarksEvidenceFresh = true
		}},
		{"connector cannot restore active closure", func(model *Point15ValBAuthorityBoundary) {
			model.ConnectorRestoresActiveClosure = true
		}},
		{"dashboard cannot suppress overdue status", func(model *Point15ValBAuthorityBoundary) {
			model.DashboardSuppressesOverdueStatus = true
		}},
		{"portal projection cannot mutate revalidation", func(model *Point15ValBAuthorityBoundary) {
			model.PortalProjectionMutatesRevalidation = true
		}},
		{"customer projection cannot mutate revalidation", func(model *Point15ValBAuthorityBoundary) {
			model.CustomerProjectionMutatesRevalidation = true
		}},
		{"auditor projection cannot mutate revalidation", func(model *Point15ValBAuthorityBoundary) {
			model.AuditorProjectionMutatesRevalidation = true
		}},
		{"agent cannot satisfy revalidation", func(model *Point15ValBAuthorityBoundary) {
			model.AgentSatisfiesRevalidation = true
		}},
		{"retry budget cannot be reset by authority surface", func(model *Point15ValBAuthorityBoundary) {
			model.RetryBudgetResetAllowed = true
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValBValidAuthorityBoundaryModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValBAuthorityBoundaryState(model); got != Point15ValBStateBlocked {
				t.Fatalf("expected blocked, got %s", got)
			}
		})
	}
}

func TestPoint15ValBNoOverclaimGuardState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValBNoOverclaimGuard)
		want   string
	}{
		{"forbidden wording blocks", func(model *Point15ValBNoOverclaimGuard) {
			model.ObservedTexts = []string{"continuous assurance guaranteed"}
		}, Point15ValBStateBlocked},
		{"safe bounded wording passes", func(*Point15ValBNoOverclaimGuard) {}, Point15ValBStateActive},
		{"classified internal blocked diagnostics allowed", func(model *Point15ValBNoOverclaimGuard) {
			model.InternalDiagnosticTexts = []string{"production approved"}
			model.InternalDiagnosticsClassifiedBlocked = true
		}, Point15ValBStateActive},
		{"unclassified internal blocked diagnostics block", func(model *Point15ValBNoOverclaimGuard) {
			model.InternalDiagnosticTexts = []string{"automatically verified forever"}
		}, Point15ValBStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValBValidNoOverclaimGuardModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValBNoOverclaimGuardState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValBScheduledRevalidationFoundationState(t *testing.T) {
	t.Run("happy path active", func(t *testing.T) {
		model := point15ValBValidFoundationModel()
		got := ComputePoint15ValBScheduledRevalidationFoundation(model)
		if got.CurrentState != Point15ValBStateActive ||
			got.DependencyState != Point15ValBStateActive ||
			got.ScheduleState != Point15ValBStateActive ||
			got.RunState != Point15ValBStateActive ||
			got.RetryBudgetState != Point15ValBStateActive ||
			got.TenantThrottleState != Point15ValBStateActive ||
			got.DowngradeBindingState != Point15ValBStateActive ||
			got.TimestampDisciplineState != Point15ValBStateActive ||
			got.AuthorityBoundaryState != Point15ValBStateActive ||
			got.NoOverclaimState != Point15ValBStateActive {
			t.Fatalf("expected full point15 valb foundation active, got %#v", got)
		}
	})

	tests := []struct {
		name   string
		mutate func(*Point15ValBScheduledRevalidationFoundation)
		assert func(*testing.T, Point15ValBScheduledRevalidationFoundation)
	}{
		{
			name: "due revalidation cannot retain active closure",
			mutate: func(model *Point15ValBScheduledRevalidationFoundation) {
				model.Schedule.ScheduledStatus = point15ValBScheduleDue
				model.DowngradeBinding.ScheduleStatus = point15ValBScheduleDue
			},
			assert: func(t *testing.T, got Point15ValBScheduledRevalidationFoundation) {
				if got.CurrentState != Point15ValBStateBlocked || got.DowngradeBindingState != Point15ValBStateBlocked {
					t.Fatalf("expected due-retention mismatch to block foundation, got %#v", got)
				}
			},
		},
		{
			name: "completed clean requires exact evidence hash binding",
			mutate: func(model *Point15ValBScheduledRevalidationFoundation) {
				*model = point15ValBCompletedCleanFoundationModel()
				model.Run.RunEvidenceHash = "hash_point15_valb_other"
			},
			assert: func(t *testing.T, got Point15ValBScheduledRevalidationFoundation) {
				if got.CurrentState != Point15ValBStateBlocked || got.RunState != Point15ValBStateBlocked {
					t.Fatalf("expected run hash mismatch to block foundation, got %#v", got)
				}
			},
		},
		{
			name: "completed run without downgrade binding run ref blocks",
			mutate: func(model *Point15ValBScheduledRevalidationFoundation) {
				*model = point15ValBCompletedCleanFoundationModel()
				model.DowngradeBinding.RunRef = ""
			},
			assert: func(t *testing.T, got Point15ValBScheduledRevalidationFoundation) {
				if got.CurrentState != Point15ValBStateBlocked || got.DowngradeBindingState != Point15ValBStateBlocked {
					t.Fatalf("expected missing completed-run ref to block foundation, got %#v", got)
				}
			},
		},
		{
			name: "cross tenant schedule reuse blocks",
			mutate: func(model *Point15ValBScheduledRevalidationFoundation) {
				model.Schedule.TenantScope = "tenant_point15_valb_other"
			},
			assert: func(t *testing.T, got Point15ValBScheduledRevalidationFoundation) {
				if got.CurrentState != Point15ValBStateBlocked || got.ScheduleState != Point15ValBStateBlocked {
					t.Fatalf("expected cross-tenant schedule mismatch to block foundation, got %#v", got)
				}
			},
		},
		{
			name: "schedule completed without run proof blocks",
			mutate: func(model *Point15ValBScheduledRevalidationFoundation) {
				model.Schedule.ScheduledStatus = point15ValBScheduleCompleted
				model.Schedule.LastCompletedAt = "2026-05-07T09:20:00Z"
				model.DowngradeBinding.ScheduleStatus = point15ValBScheduleCompleted
			},
			assert: func(t *testing.T, got Point15ValBScheduledRevalidationFoundation) {
				if got.CurrentState != Point15ValBStateBlocked || got.ScheduleState != Point15ValBStateBlocked || got.RunState != Point15ValBStateBlocked {
					t.Fatalf("expected completed schedule without run proof to block foundation, got %#v", got)
				}
			},
		},
		{
			name: "throttled blocked cannot remain active",
			mutate: func(model *Point15ValBScheduledRevalidationFoundation) {
				model.TenantThrottle.RequestedRevalidations = 20
				model.TenantThrottle.ThrottleStatus = point15ValBThrottleBlocked
				model.DowngradeBinding.ThrottleStatus = point15ValBThrottleBlocked
			},
			assert: func(t *testing.T, got Point15ValBScheduledRevalidationFoundation) {
				if got.CurrentState != Point15ValBStateBlocked || got.TenantThrottleState != Point15ValBStateBlocked {
					t.Fatalf("expected throttled blocked to block foundation, got %#v", got)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValBValidFoundationModel()
			tc.mutate(&model)
			got := ComputePoint15ValBScheduledRevalidationFoundation(model)
			tc.assert(t, got)
		})
	}
}

func TestPoint10ThroughPoint15ValBCurrentSweep(t *testing.T) {
	computed := ComputePoint15ValBScheduledRevalidationFoundation(point15ValBValidFoundationModel())
	if computed.DependencyState != Point15ValBStateActive {
		t.Fatalf("expected dependency active, got %s", computed.DependencyState)
	}
	if computed.CurrentState != Point15ValBStateActive {
		t.Fatalf("expected current state active, got %s", computed.CurrentState)
	}
	if computed.Dependency.Point15PassSeen {
		t.Fatal("expected no point_15_pass in point15 val b sweep")
	}
}

func TestPoint15ValBCachedHelperIsolation(t *testing.T) {
	model := point15ValBValidFoundationModel()
	originalAllowed := model.NoOverclaimGuard.AllowedSafeWording[0]
	model.NoOverclaimGuard.AllowedSafeWording[0] = "mutated"

	fresh := point15ValBValidFoundationModel()
	if fresh.NoOverclaimGuard.AllowedSafeWording[0] != originalAllowed {
		t.Fatalf("expected cached point15 valb helper to return isolated copy, got %#v", fresh.NoOverclaimGuard.AllowedSafeWording)
	}
}

func TestPoint15ValBCachedHelperNestedDependencyIsolation(t *testing.T) {
	model := point15ValBValidFoundationModel()
	originalAllowed := model.Dependency.Point15ValA.NoOverclaimGuard.AllowedSafeWording[0]
	model.Dependency.Point15ValA.NoOverclaimGuard.AllowedSafeWording[0] = "mutated"

	fresh := point15ValBValidFoundationModel()
	if fresh.Dependency.Point15ValA.NoOverclaimGuard.AllowedSafeWording[0] != originalAllowed {
		t.Fatalf("expected cached point15 valb nested dependency helper to return isolated copy, got %#v", fresh.Dependency.Point15ValA.NoOverclaimGuard.AllowedSafeWording)
	}
}

package formal

import (
	"encoding/json"
	"sync"
	"testing"
)

var (
	point15ValEFoundationOnce sync.Once
	point15ValEFoundationBase Point15ValEContinuousVerificationClosureFoundation
)

func point15ValECloneStrings(values []string) []string {
	return append([]string(nil), values...)
}

func clonePoint15ValEFoundation(model Point15ValEContinuousVerificationClosureFoundation) Point15ValEContinuousVerificationClosureFoundation {
	model.BlockingReasons = point15ValECloneStrings(model.BlockingReasons)
	model.ReviewPrerequisites = point15ValECloneStrings(model.ReviewPrerequisites)
	model.Dependency.ReviewPrerequisites = point15ValECloneStrings(model.Dependency.ReviewPrerequisites)
	model.Dependency.Point15ValD = clonePoint15ValDFoundation(model.Dependency.Point15ValD)
	model.ClosureEvaluator.CommandsRun = point15ValECloneStrings(model.ClosureEvaluator.CommandsRun)
	model.ClosureEvaluator.TestsRun = point15ValECloneStrings(model.ClosureEvaluator.TestsRun)
	model.ClosureEvaluator.GrepsRun = point15ValECloneStrings(model.ClosureEvaluator.GrepsRun)
	model.ClosureEvaluator.NegativeFixturesRun = point15ValECloneStrings(model.ClosureEvaluator.NegativeFixturesRun)
	model.PassClosureManifest.ExplicitNonGoals = point15ValECloneStrings(model.PassClosureManifest.ExplicitNonGoals)
	model.PassClosureManifest.CommandsRun = point15ValECloneStrings(model.PassClosureManifest.CommandsRun)
	model.PassClosureManifest.TestsRun = point15ValECloneStrings(model.PassClosureManifest.TestsRun)
	model.PassClosureManifest.GrepsRun = point15ValECloneStrings(model.PassClosureManifest.GrepsRun)
	model.PassClosureManifest.NegativeFixturesRun = point15ValECloneStrings(model.PassClosureManifest.NegativeFixturesRun)
	model.NoOverclaimFinalCheck.ObservedTexts = point15ValECloneStrings(model.NoOverclaimFinalCheck.ObservedTexts)
	model.NoOverclaimFinalCheck.InternalDiagnosticTexts = point15ValECloneStrings(model.NoOverclaimFinalCheck.InternalDiagnosticTexts)
	model.NoOverclaimFinalCheck.AllowedSafeWording = point15ValECloneStrings(model.NoOverclaimFinalCheck.AllowedSafeWording)
	model.NoOverclaimFinalCheck.BlockedWording = point15ValECloneStrings(model.NoOverclaimFinalCheck.BlockedWording)
	model.CLBFinalCheck.CLB3Advisories = point15ValECloneStrings(model.CLBFinalCheck.CLB3Advisories)
	return model
}

func point15ValEValidFoundationModel() Point15ValEContinuousVerificationClosureFoundation {
	point15ValEFoundationOnce.Do(func() {
		point15ValEFoundationBase = ComputePoint15ValEFoundation(Point15ValEFoundationModel())
	})
	return clonePoint15ValEFoundation(point15ValEFoundationBase)
}

func point15ValEValidDependencyModel() Point15ValEDependencySnapshot {
	return point15ValEValidFoundationModel().Dependency
}

func point15ValEValidClosureEvaluatorModel() Point15ValEClosureEvaluator {
	model := point15ValEClosureEvaluatorModel()
	model.DependencyState = Point15ValEStatePassConfirmed
	model.FreshnessTaxonomyState = Point15ValEStatePassConfirmed
	model.DowngradeTriggerState = Point15ValEStatePassConfirmed
	model.ScheduledRevalidationState = Point15ValEStatePassConfirmed
	model.EnforcementBoundaryState = Point15ValEStatePassConfirmed
	model.ProjectionBoundaryState = Point15ValEStatePassConfirmed
	model.ReplayProofHistoryState = Point15ValEStatePassConfirmed
	model.TenantPrivacyState = Point15ValEStatePassConfirmed
	model.TimestampIntegrityState = Point15ValEStatePassConfirmed
	model.AuthorityBoundaryState = Point15ValEStatePassConfirmed
	model.NoMutationState = Point15ValEStatePassConfirmed
	model.NoOverclaimState = Point15ValEStatePassConfirmed
	model.CLBFinalState = Point15ValEStatePassConfirmed
	model.FinalPassAllowed = true
	return model
}

func point15ValEValidPassClosureManifestModel() Point15PassClosureManifest {
	model := point15ValEPassClosureManifestModel(point15ValEDependencySnapshotModel())
	model.DependencyGateResult = Point15ValEStatePassConfirmed
	model.FreshnessTaxonomyResult = Point15ValEStatePassConfirmed
	model.DowngradeTriggerResult = Point15ValEStatePassConfirmed
	model.ScheduledRevalidationResult = Point15ValEStatePassConfirmed
	model.EnforcementBoundaryResult = Point15ValEStatePassConfirmed
	model.ProjectionBoundaryResult = Point15ValEStatePassConfirmed
	model.ReplayProofHistoryResult = Point15ValEStatePassConfirmed
	model.TenantPrivacyResult = Point15ValEStatePassConfirmed
	model.TimestampIntegrityResult = Point15ValEStatePassConfirmed
	model.AuthorityBoundaryResult = Point15ValEStatePassConfirmed
	model.NoMutationResult = Point15ValEStatePassConfirmed
	model.NoOverclaimResult = Point15ValEStatePassConfirmed
	model.CLBResult = Point15ValEStatePassConfirmed
	model.Point15PassAllowed = true
	model.Point15PassToken = point15Val0BlockedPassToken
	return model
}

func point15ValEValidFreshnessModel() Point15ValEFreshnessTaxonomyClosureCheck {
	return point15ValEValidFoundationModel().FreshnessTaxonomyClosureCheck
}

func point15ValEValidDowngradeModel() Point15ValEDowngradeTriggerClosureCheck {
	return point15ValEValidFoundationModel().DowngradeTriggerClosureCheck
}

func point15ValEValidRevalidationModel() Point15ValEScheduledRevalidationClosureCheck {
	return point15ValEValidFoundationModel().ScheduledRevalidationClosureCheck
}

func point15ValEValidEnforcementModel() Point15ValEEnforcementClosureCheck {
	return point15ValEValidFoundationModel().EnforcementClosureCheck
}

func point15ValEValidProjectionModel() Point15ValEProjectionClosureCheck {
	return point15ValEValidFoundationModel().ProjectionClosureCheck
}

func point15ValEValidReplayModel() Point15ValEReplayProofHistoryClosureCheck {
	return point15ValEValidFoundationModel().ReplayProofHistoryClosureCheck
}

func point15ValEValidTenantModel() Point15ValETenantPrivacyClosureCheck {
	return point15ValEValidFoundationModel().TenantPrivacyClosureCheck
}

func point15ValEValidTimestampModel() Point15ValETimestampIntegrityClosureCheck {
	return point15ValEValidFoundationModel().TimestampIntegrityClosureCheck
}

func point15ValEValidAuthorityModel() Point15ValEAuthorityBoundaryClosureCheck {
	return point15ValEValidFoundationModel().AuthorityBoundaryClosureCheck
}

func point15ValEValidNoMutationModel() Point15ValENoMutationClosureCheck {
	return point15ValEValidFoundationModel().NoMutationClosureCheck
}

func point15ValEValidNoOverclaimModel() Point15ValENoOverclaimFinalCheck {
	return point15ValEValidFoundationModel().NoOverclaimFinalCheck
}

func point15ValEValidCLBModel() Point15ValECLBFinalCheck {
	return point15ValEValidFoundationModel().CLBFinalCheck
}

func TestPoint15ValEDependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValEDependencySnapshot)
		want   string
	}{
		{"pass confirmed when vald clean", func(*Point15ValEDependencySnapshot) {}, Point15ValEStatePassConfirmed},
		{"blocks when vald missing", func(model *Point15ValEDependencySnapshot) { model.Point15ValDCurrentState = "" }, Point15ValEStateBlocked},
		{"blocks when vald blocked", func(model *Point15ValEDependencySnapshot) { model.Point15ValDCurrentState = Point15ValDStateBlocked }, Point15ValEStateBlocked},
		{"blocks when vald review required", func(model *Point15ValEDependencySnapshot) {
			model.Point15ValDCurrentState = Point15ValDStateReviewRequired
		}, Point15ValEStateBlocked},
		{"blocks when vald incomplete", func(model *Point15ValEDependencySnapshot) { model.Point15ValDCurrentState = Point15ValDStateIncomplete }, Point15ValEStateBlocked},
		{"blocks when vald not merged", func(model *Point15ValEDependencySnapshot) { model.Point15ValDMerged = false }, Point15ValEStateBlocked},
		{"blocks when vald ci not green", func(model *Point15ValEDependencySnapshot) { model.Point15ValDCIGreen = false }, Point15ValEStateBlocked},
		{"blocks when vald not reviewed on main", func(model *Point15ValEDependencySnapshot) { model.Point15ValDReviewedOnMain = false }, Point15ValEStateBlocked},
		{"blocks when inherited valc missing", func(model *Point15ValEDependencySnapshot) { model.InheritedPoint15ValCCurrentState = "" }, Point15ValEStateBlocked},
		{"blocks when point15 pass appears before final path", func(model *Point15ValEDependencySnapshot) { model.Point15PassSeen = true }, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidDependencyModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValEDependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEClosureEvaluatorState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValEClosureEvaluator)
		want   string
	}{
		{"all closure checks clean allows pass confirmed", func(*Point15ValEClosureEvaluator) {}, Point15ValEStatePassConfirmed},
		{"dependency blocked blocks final closure", func(model *Point15ValEClosureEvaluator) { model.DependencyState = Point15ValEStateBlocked }, Point15ValEStateBlocked},
		{"review required component stays review required", func(model *Point15ValEClosureEvaluator) {
			model.ProjectionBoundaryState = Point15ValEStateReviewRequired
		}, Point15ValEStateReviewRequired},
		{"incomplete component stays incomplete", func(model *Point15ValEClosureEvaluator) { model.TimestampIntegrityState = Point15ValEStateIncomplete }, Point15ValEStateIncomplete},
		{"no mutation paths false blocks", func(model *Point15ValEClosureEvaluator) { model.NoMutationPathsDetected = false }, Point15ValEStateBlocked},
		{"premature point15 pass blocks", func(model *Point15ValEClosureEvaluator) { model.NoPrematurePoint15Pass = false }, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidClosureEvaluatorModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValEClosureEvaluatorState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEPassClosureManifestState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15PassClosureManifest)
		want   string
	}{
		{"complete manifest allows final pass confirmed", func(*Point15PassClosureManifest) {}, Point15ValEStatePassConfirmed},
		{"missing point id blocks", func(model *Point15PassClosureManifest) { model.PointID = "" }, Point15ValEStateBlocked},
		{"missing wave id blocks", func(model *Point15PassClosureManifest) { model.WaveID = "" }, Point15ValEStateBlocked},
		{"wrong closure token blocks", func(model *Point15PassClosureManifest) { model.ClosureToken = "point_15_prepass" }, Point15ValEStateBlocked},
		{"missing dependency gate result blocks", func(model *Point15PassClosureManifest) { model.DependencyGateResult = "" }, Point15ValEStateBlocked},
		{"missing commands run blocks", func(model *Point15PassClosureManifest) { model.CommandsRun = nil }, Point15ValEStateBlocked},
		{"missing tests run blocks", func(model *Point15PassClosureManifest) { model.TestsRun = nil }, Point15ValEStateBlocked},
		{"missing greps run blocks", func(model *Point15PassClosureManifest) { model.GrepsRun = nil }, Point15ValEStateBlocked},
		{"missing negative fixtures blocks", func(model *Point15PassClosureManifest) { model.NegativeFixturesRun = nil }, Point15ValEStateBlocked},
		{"missing evidence identity blocks", func(model *Point15PassClosureManifest) { model.EvidenceIdentity = "" }, Point15ValEStateBlocked},
		{"missing evidence hash blocks", func(model *Point15PassClosureManifest) { model.EvidenceHash = "" }, Point15ValEStateBlocked},
		{"missing tenant scope blocks", func(model *Point15PassClosureManifest) { model.TenantScope = "" }, Point15ValEStateBlocked},
		{"missing policy version blocks", func(model *Point15PassClosureManifest) { model.PolicyVersion = "" }, Point15ValEStateBlocked},
		{"missing engine version blocks", func(model *Point15PassClosureManifest) { model.EngineVersion = "" }, Point15ValEStateBlocked},
		{"missing schema version blocks", func(model *Point15PassClosureManifest) { model.SchemaVersion = "" }, Point15ValEStateBlocked},
		{"missing projection boundary result blocks", func(model *Point15PassClosureManifest) { model.ProjectionBoundaryResult = "" }, Point15ValEStateBlocked},
		{"missing tenant privacy result blocks", func(model *Point15PassClosureManifest) { model.TenantPrivacyResult = "" }, Point15ValEStateBlocked},
		{"missing no overclaim result blocks", func(model *Point15PassClosureManifest) { model.NoOverclaimResult = "" }, Point15ValEStateBlocked},
		{"missing clb result blocks", func(model *Point15PassClosureManifest) { model.CLBResult = "" }, Point15ValEStateBlocked},
		{"review required dependency returns review required", func(model *Point15PassClosureManifest) {
			model.DependencyGateResult = Point15ValEStateReviewRequired
			model.Point15PassAllowed = false
			model.Point15PassToken = ""
		}, Point15ValEStateReviewRequired},
		{"review required tenant privacy returns review required", func(model *Point15PassClosureManifest) {
			model.TenantPrivacyResult = Point15ValEStateReviewRequired
			model.Point15PassAllowed = false
			model.Point15PassToken = ""
		}, Point15ValEStateReviewRequired},
		{"incomplete dependency returns incomplete", func(model *Point15PassClosureManifest) {
			model.DependencyGateResult = Point15ValEStateIncomplete
			model.Point15PassAllowed = false
			model.Point15PassToken = ""
		}, Point15ValEStateIncomplete},
		{"point15 pass allowed only in final clean path", func(model *Point15PassClosureManifest) {
			model.DependencyGateResult = Point15ValEStateReviewRequired
			model.Point15PassAllowed = true
			model.Point15PassToken = point15Val0BlockedPassToken
		}, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidPassClosureManifestModel()
			tc.mutate(&model)
			if got := EvaluatePoint15PassClosureManifestState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEFreshnessTaxonomyClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValEFreshnessTaxonomyClosureCheck)
		want   string
	}{
		{"fresh exact binding allows pass confirmed", func(*Point15ValEFreshnessTaxonomyClosureCheck) {}, Point15ValEStatePassConfirmed},
		{"stale cannot retain pass", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.FreshnessStatus = point15Val0FreshnessStale
			model.DowngradeOutcome = point15Val0DowngradeReview
			model.MappedState = Point15Val0StateReviewRequired
			model.RetainsActiveClosure = false
		}, Point15ValEStateReviewRequired},
		{"expired blocks", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.FreshnessStatus = point15Val0FreshnessExpired
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.MappedState = Point15Val0StateBlocked
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
		{"revoked blocks", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.FreshnessStatus = point15Val0FreshnessRevoked
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.MappedState = Point15Val0StateBlocked
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
		{"superseded without lineage blocks", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.FreshnessStatus = point15Val0FreshnessSuperseded
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.MappedState = Point15Val0StateBlocked
			model.SupersessionLineageRef = ""
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
		{"superseded with lineage remains bounded review path", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.FreshnessStatus = point15Val0FreshnessSuperseded
			model.DowngradeOutcome = point15Val0DowngradeReview
			model.MappedState = Point15Val0StateReviewRequired
			model.SupersessionLineageRef = "lineage_point15_vale_001"
			model.RetainsActiveClosure = false
		}, Point15ValEStateReviewRequired},
		{"drifted decisive blocks", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.FreshnessStatus = point15Val0FreshnessDrifted
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.MappedState = Point15Val0StateBlocked
			model.DriftIsDecisive = true
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
		{"missing proof can be incomplete", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.FreshnessStatus = point15Val0FreshnessMissing
			model.DowngradeOutcome = point15Val0DowngradeIncomplete
			model.MappedState = Point15Val0StateIncomplete
			model.MissingFreshnessProofDecisive = false
			model.FreshnessProofPresent = true
			model.RetainsActiveClosure = false
		}, Point15ValEStateIncomplete},
		{"unsupported blocks", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.FreshnessStatus = point15Val0FreshnessUnsupported
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.MappedState = Point15Val0StateBlocked
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
		{"tampered blocks", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.FreshnessStatus = point15Val0FreshnessTampered
			model.DowngradeOutcome = point15Val0DowngradeBlocked
			model.MappedState = Point15Val0StateBlocked
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidFreshnessModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValEFreshnessTaxonomyClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEDowngradeTriggerClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValEDowngradeTriggerClosureCheck)
		want   string
	}{
		{"no trigger allows pass confirmed", func(*Point15ValEDowngradeTriggerClosureCheck) {}, Point15ValEStatePassConfirmed},
		{"expired evidence trigger blocks", func(model *Point15ValEDowngradeTriggerClosureCheck) {
			model.TriggerDetected = true
			model.TriggerType = point15ValATriggerExpired
			model.TriggerIsDecisive = true
			model.TargetState = Point15Val0StateBlocked
			model.TargetDowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
		{"stale evidence trigger becomes review required", func(model *Point15ValEDowngradeTriggerClosureCheck) {
			model.TriggerDetected = true
			model.TriggerType = point15ValATriggerStale
			model.TriggerIsDecisive = false
			model.TargetState = Point15Val0StateReviewRequired
			model.TargetDowngradeOutcome = point15Val0DowngradeReview
			model.RetainsActiveClosure = false
		}, Point15ValEStateReviewRequired},
		{"connector unauthorized trigger blocks", func(model *Point15ValEDowngradeTriggerClosureCheck) {
			model.TriggerDetected = true
			model.TriggerType = point15ValATriggerConnAuth
			model.TriggerIsDecisive = true
			model.TargetState = Point15Val0StateBlocked
			model.TargetDowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
		{"connector tenant mismatch trigger blocks", func(model *Point15ValEDowngradeTriggerClosureCheck) {
			model.TriggerDetected = true
			model.TriggerType = point15ValATriggerConnTenant
			model.TriggerIsDecisive = true
			model.TargetState = Point15Val0StateBlocked
			model.TargetDowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
		{"trigger retaining pass blocks", func(model *Point15ValEDowngradeTriggerClosureCheck) {
			model.TriggerDetected = true
			model.TriggerType = point15ValATriggerHash
			model.TriggerIsDecisive = true
			model.TargetState = Point15Val0StateBlocked
			model.TargetDowngradeOutcome = point15Val0DowngradeBlocked
			model.RetainsPass = true
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidDowngradeModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValEDowngradeTriggerClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEScheduledRevalidationClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValEScheduledRevalidationClosureCheck)
		want   string
	}{
		{"scheduled revalidation clean allows pass confirmed", func(*Point15ValEScheduledRevalidationClosureCheck) {}, Point15ValEStatePassConfirmed},
		{"missed required revalidation cannot retain active", func(model *Point15ValEScheduledRevalidationClosureCheck) {
			model.ScheduledStatus = point15ValBScheduleMissed
			model.TargetState = Point15Val0StateBlocked
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
		{"overdue required revalidation cannot retain active", func(model *Point15ValEScheduledRevalidationClosureCheck) {
			model.ScheduledStatus = point15ValBScheduleOverdue
			model.TargetState = Point15Val0StateReviewRequired
			model.RetainsActiveClosure = false
		}, Point15ValEStateReviewRequired},
		{"retry exhausted cannot retain active", func(model *Point15ValEScheduledRevalidationClosureCheck) {
			model.RetryBudgetStatus = point15ValBRetryExhausted
			model.TargetState = Point15Val0StateBlocked
			model.RetainsActiveClosure = false
		}, Point15ValEStateBlocked},
		{"throttled cannot hide downgrade", func(model *Point15ValEScheduledRevalidationClosureCheck) {
			model.ThrottleStatus = point15ValBThrottleReviewRequired
			model.TargetState = Point15Val0StateReviewRequired
			model.RetainsActiveClosure = false
		}, Point15ValEStateReviewRequired},
		{"scheduler cannot mark fresh or restore pass", func(model *Point15ValEScheduledRevalidationClosureCheck) {
			model.SchedulerAuthorityGranted = true
		}, Point15ValEStateBlocked},
		{"completed clean run requires exact binding", func(model *Point15ValEScheduledRevalidationClosureCheck) {
			model.ScheduledStatus = point15ValBScheduleCompleted
			model.RunResult = point15ValBRunCompletedClean
			model.ExactBindingConfirmed = false
		}, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidRevalidationModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValEScheduledRevalidationClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEEnforcementClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValEEnforcementClosureCheck)
		want   string
	}{
		{"clean enforcement boundary allows pass confirmed", func(*Point15ValEEnforcementClosureCheck) {}, Point15ValEStatePassConfirmed},
		{"expired enforcement blocks", func(model *Point15ValEEnforcementClosureCheck) {
			model.EnforcementAction = point15ValCActionBlocked
			model.EnforcementReason = point15ValCReasonExpired
			model.TargetState = Point15Val0StateBlocked
			model.LifecycleStatus = point15ValCLifecycleExpired
		}, Point15ValEStateBlocked},
		{"supersession without lineage equivalent blocks via silent replacement", func(model *Point15ValEEnforcementClosureCheck) {
			model.SilentReplacementDetected = true
		}, Point15ValEStateBlocked},
		{"evidence deletion hiding blocks", func(model *Point15ValEEnforcementClosureCheck) {
			model.EvidenceDeletionDetected = true
		}, Point15ValEStateBlocked},
		{"revocation auto publish blocks", func(model *Point15ValEEnforcementClosureCheck) {
			model.AutomaticPublicationDetected = true
		}, Point15ValEStateBlocked},
		{"canonical mutation blocks", func(model *Point15ValEEnforcementClosureCheck) {
			model.CanonicalMutationAttempted = true
		}, Point15ValEStateBlocked},
		{"history must remain preserved", func(model *Point15ValEEnforcementClosureCheck) {
			model.HistoryPreserved = false
		}, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidEnforcementModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValEEnforcementClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEProjectionClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValEProjectionClosureCheck)
		want   string
	}{
		{"projection remains display only", func(*Point15ValEProjectionClosureCheck) {}, Point15ValEStatePassConfirmed},
		{"timeline display cannot mutate", func(model *Point15ValEProjectionClosureCheck) { model.MutatesState = true }, Point15ValEStateBlocked},
		{"dashboard cannot restore pass", func(model *Point15ValEProjectionClosureCheck) { model.RestoresActive = true }, Point15ValEStateBlocked},
		{"query cannot enforce", func(model *Point15ValEProjectionClosureCheck) { model.PerformsEnforcement = true }, Point15ValEStateBlocked},
		{"projection cannot hide decisive evidence", func(model *Point15ValEProjectionClosureCheck) { model.HidesDecisiveEvidence = true }, Point15ValEStateBlocked},
		{"projection cannot strengthen claims", func(model *Point15ValEProjectionClosureCheck) { model.StrengthensClaims = true }, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidProjectionModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValEProjectionClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEReplayProofHistoryClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValEReplayProofHistoryClosureCheck)
		want   string
	}{
		{"replay proof history remains visible", func(*Point15ValEReplayProofHistoryClosureCheck) {}, Point15ValEStatePassConfirmed},
		{"missing replay ref blocks", func(model *Point15ValEReplayProofHistoryClosureCheck) { model.ReplayRef = "" }, Point15ValEStateBlocked},
		{"prior state must remain visible", func(model *Point15ValEReplayProofHistoryClosureCheck) { model.PriorStateVisible = false }, Point15ValEStateBlocked},
		{"hash binding must remain visible", func(model *Point15ValEReplayProofHistoryClosureCheck) { model.HashBindingVisible = false }, Point15ValEStateBlocked},
		{"hidden proof history blocks", func(model *Point15ValEReplayProofHistoryClosureCheck) { model.ProofHistoryHidden = true }, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidReplayModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValEReplayProofHistoryClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValETenantPrivacyClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValETenantPrivacyClosureCheck)
		want   string
	}{
		{"tenant privacy clean allows pass confirmed", func(*Point15ValETenantPrivacyClosureCheck) {}, Point15ValEStatePassConfirmed},
		{"cross tenant proof blocks", func(model *Point15ValETenantPrivacyClosureCheck) { model.CrossTenantProofDetected = true }, Point15ValEStateBlocked},
		{"cross tenant schedule run blocks", func(model *Point15ValETenantPrivacyClosureCheck) { model.CrossTenantScheduleRunDetected = true }, Point15ValEStateBlocked},
		{"cross tenant enforcement blocks", func(model *Point15ValETenantPrivacyClosureCheck) { model.CrossTenantEnforcementDetected = true }, Point15ValEStateBlocked},
		{"cross tenant projection blocks", func(model *Point15ValETenantPrivacyClosureCheck) { model.CrossTenantProjectionDetected = true }, Point15ValEStateBlocked},
		{"tenant private data exposure blocks", func(model *Point15ValETenantPrivacyClosureCheck) { model.TenantPrivateDataExposed = true }, Point15ValEStateBlocked},
		{"redaction hiding decisive evidence requires review", func(model *Point15ValETenantPrivacyClosureCheck) { model.RedactionHidesDecisiveEvidence = true }, Point15ValEStateReviewRequired},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidTenantModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValETenantPrivacyClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValETimestampIntegrityClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValETimestampIntegrityClosureCheck)
		want   string
	}{
		{"timestamp integrity clean allows pass confirmed", func(*Point15ValETimestampIntegrityClosureCheck) {}, Point15ValEStatePassConfirmed},
		{"client local canonical time blocks", func(model *Point15ValETimestampIntegrityClosureCheck) { model.ClientLocalCreatesCanonical = true }, Point15ValEStateBlocked},
		{"source event only canonical validity blocks", func(model *Point15ValETimestampIntegrityClosureCheck) { model.SourceEventCreatesCanonical = true }, Point15ValEStateBlocked},
		{"future enforcement projection event requires review", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.EnforcementEnforcedAt = "2027-01-01T00:00:00Z"
			model.ProjectionDisplayedAt = "2027-01-01T00:00:01Z"
		}, Point15ValEStateReviewRequired},
		{"backdated enforcement requires review", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.EnforcementEnforcedAt = "2020-01-01T00:00:00Z"
		}, Point15ValEStateReviewRequired},
		{"missing trusted displayed time is incomplete", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.ProjectionDisplayedAt = ""
		}, Point15ValEStateIncomplete},
		{"untrusted displayed time source blocks", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.ProjectionDisplayedTimeSource = point14Val0TimeSourceClientLocal
		}, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidTimestampModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValETimestampIntegrityClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEAuthorityBoundaryClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValEAuthorityBoundaryClosureCheck)
		want   string
	}{
		{"authority boundary clean allows pass confirmed", func(*Point15ValEAuthorityBoundaryClosureCheck) {}, Point15ValEStatePassConfirmed},
		{"agent cannot pass", func(model *Point15ValEAuthorityBoundaryClosureCheck) { model.AgentPassAllowed = true }, Point15ValEStateBlocked},
		{"ai cannot pass", func(model *Point15ValEAuthorityBoundaryClosureCheck) { model.AIPassAllowed = true }, Point15ValEStateBlocked},
		{"scheduler cannot pass", func(model *Point15ValEAuthorityBoundaryClosureCheck) { model.SchedulerPassAllowed = true }, Point15ValEStateBlocked},
		{"dashboard cannot pass", func(model *Point15ValEAuthorityBoundaryClosureCheck) { model.DashboardPassAllowed = true }, Point15ValEStateBlocked},
		{"timeline cannot create authority", func(model *Point15ValEAuthorityBoundaryClosureCheck) { model.TimelineAuthorityAllowed = true }, Point15ValEStateBlocked},
		{"query cannot mutate", func(model *Point15ValEAuthorityBoundaryClosureCheck) { model.QueryMutationAllowed = true }, Point15ValEStateBlocked},
		{"external source cannot become authority", func(model *Point15ValEAuthorityBoundaryClosureCheck) { model.ExternalAuthorityAllowed = true }, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidAuthorityModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValEAuthorityBoundaryClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValENoMutationClosureCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValENoMutationClosureCheck)
		want   string
	}{
		{"no mutation path allows pass confirmed", func(*Point15ValENoMutationClosureCheck) {}, Point15ValEStatePassConfirmed},
		{"canonical mutation blocks", func(model *Point15ValENoMutationClosureCheck) { model.CanonicalMutationDetected = true }, Point15ValEStateBlocked},
		{"production mutation blocks", func(model *Point15ValENoMutationClosureCheck) { model.ProductionMutationDetected = true }, Point15ValEStateBlocked},
		{"evidence deletion blocks", func(model *Point15ValENoMutationClosureCheck) { model.EvidenceDeletionDetected = true }, Point15ValEStateBlocked},
		{"history hiding blocks", func(model *Point15ValENoMutationClosureCheck) { model.HistoryHidingDetected = true }, Point15ValEStateBlocked},
		{"revocation execution blocks", func(model *Point15ValENoMutationClosureCheck) { model.RevocationExecutionDetected = true }, Point15ValEStateBlocked},
		{"automatic publication blocks", func(model *Point15ValENoMutationClosureCheck) { model.AutomaticPublicationDetected = true }, Point15ValEStateBlocked},
		{"silent supersession replacement blocks", func(model *Point15ValENoMutationClosureCheck) { model.SilentSupersessionReplacement = true }, Point15ValEStateBlocked},
		{"retry budget reset by non core blocks", func(model *Point15ValENoMutationClosureCheck) { model.RetryBudgetResetByNonCore = true }, Point15ValEStateBlocked},
		{"pass restoration blocks", func(model *Point15ValENoMutationClosureCheck) { model.PassRestorationDetected = true }, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidNoMutationModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValENoMutationClosureCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValENoOverclaimFinalCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValENoOverclaimFinalCheck)
		want   string
	}{
		{"safe bounded wording passes", func(*Point15ValENoOverclaimFinalCheck) {}, Point15ValEStatePassConfirmed},
		{"forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "continuous assurance guaranteed")
		}, Point15ValEStateBlocked},
		{"mutated allowed wording set blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.AllowedSafeWording = model.AllowedSafeWording[:len(model.AllowedSafeWording)-1]
		}, Point15ValEStateBlocked},
		{"internal blocked diagnostics must stay classified", func(model *Point15ValENoOverclaimFinalCheck) {
			model.InternalDiagnosticTexts = []string{"public badge"}
			model.InternalDiagnosticsClassifiedBlocked = false
		}, Point15ValEStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidNoOverclaimModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValENoOverclaimFinalCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValECLBFinalCheckState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValECLBFinalCheck)
		want   string
	}{
		{"no open clb blocker allows pass confirmed", func(*Point15ValECLBFinalCheck) {}, Point15ValEStatePassConfirmed},
		{"clb0 blocks", func(model *Point15ValECLBFinalCheck) { model.CLB0Present = true }, Point15ValEStateBlocked},
		{"clb1 blocks", func(model *Point15ValECLBFinalCheck) { model.CLB1Present = true }, Point15ValEStateBlocked},
		{"clb2 blocks", func(model *Point15ValECLBFinalCheck) { model.CLB2Present = true }, Point15ValEStateBlocked},
		{"clb3 advisory remains recorded but does not block", func(model *Point15ValECLBFinalCheck) {
			model.CLB3Advisories = []string{"manual follow-through only"}
		}, Point15ValEStatePassConfirmed},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidCLBModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValECLBFinalCheckState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEFoundationState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValEContinuousVerificationClosureFoundation)
		want   string
	}{
		{"final closure emits point15 pass only on clean vale path", func(*Point15ValEContinuousVerificationClosureFoundation) {}, Point15ValEStatePassConfirmed},
		{"missing projection boundary result blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.ProjectionClosureCheck.DisplayOnly = false
		}, Point15ValEStateBlocked},
		{"clb blocker prevents final pass", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.CLBFinalCheck.CLB1Present = true
		}, Point15ValEStateBlocked},
		{"manifest evidence identity drift blocks", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.EvidenceIdentity = "wrong_identity"
		}, Point15ValEStateBlocked},
		{"missing trusted timestamp keeps final closure incomplete", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TimestampIntegrityClosureCheck.ProjectionDisplayedAt = ""
		}, Point15ValEStateIncomplete},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValEValidFoundationModel()
			tc.mutate(&model)
			result := ComputePoint15ValEFoundation(model)
			if result.CurrentState != tc.want {
				t.Fatalf("expected %s, got %s (blocking=%v review=%v dep=%s fresh=%s downgrade=%s revalidation=%s enforcement=%s projection=%s replay=%s tenant=%s timestamp=%s authority=%s mutation=%s nooverclaim=%s clb=%s evaluator=%s manifest=%s)",
					tc.want,
					result.CurrentState,
					result.BlockingReasons,
					result.ReviewPrerequisites,
					result.DependencyState,
					result.FreshnessTaxonomyClosureState,
					result.DowngradeTriggerClosureState,
					result.ScheduledRevalidationClosureState,
					result.EnforcementClosureState,
					result.ProjectionClosureState,
					result.ReplayProofHistoryClosureState,
					result.TenantPrivacyClosureState,
					result.TimestampIntegrityClosureState,
					result.AuthorityBoundaryClosureState,
					result.NoMutationClosureState,
					result.NoOverclaimFinalCheckState,
					result.CLBFinalCheckState,
					result.ClosureEvaluatorState,
					result.PassClosureManifestState,
				)
			}
			if tc.want == Point15ValEStatePassConfirmed {
				if !result.PassClosureManifest.Point15PassAllowed || result.PassClosureManifest.Point15PassToken != point15Val0BlockedPassToken {
					t.Fatalf("expected final point15 pass token in clean vale path")
				}
			} else if result.PassClosureManifest.Point15PassAllowed || result.PassClosureManifest.Point15PassToken != "" {
				t.Fatalf("expected no point15 pass token outside clean vale path")
			}
		})
	}
}

func TestPoint15ValETwitterPrivacyReviewPathPreservesSeverity(t *testing.T) {
	model := point15ValEValidFoundationModel()
	model.TenantPrivacyClosureCheck.RedactionHidesDecisiveEvidence = true

	result := ComputePoint15ValEFoundation(model)

	if result.TenantPrivacyClosureState != Point15ValEStateReviewRequired {
		t.Fatalf("expected tenant privacy review_required, got %s", result.TenantPrivacyClosureState)
	}
	if result.ClosureEvaluatorState != Point15ValEStateReviewRequired {
		t.Fatalf("expected closure evaluator review_required, got %s", result.ClosureEvaluatorState)
	}
	if result.PassClosureManifestState != Point15ValEStateReviewRequired {
		t.Fatalf("expected pass closure manifest review_required, got %s", result.PassClosureManifestState)
	}
	if result.CurrentState != Point15ValEStateReviewRequired {
		t.Fatalf("expected final state review_required, got %s", result.CurrentState)
	}
	if result.PassClosureManifest.Point15PassAllowed || result.PassClosureManifest.Point15PassToken != "" {
		t.Fatalf("expected no point15 pass token on tenant privacy review path")
	}
}

func TestPoint15ValECachedHelperIsolation(t *testing.T) {
	first := point15ValEValidFoundationModel()
	first.BlockingReasons = append(first.BlockingReasons, "mutated")
	first.ClosureEvaluator.CommandsRun = append(first.ClosureEvaluator.CommandsRun, "mutated")
	first.PassClosureManifest.TestsRun = append(first.PassClosureManifest.TestsRun, "mutated")
	first.NoOverclaimFinalCheck.ObservedTexts = append(first.NoOverclaimFinalCheck.ObservedTexts, "mutated")
	first.Dependency.Point15ValD.Query.Filters = append(first.Dependency.Point15ValD.Query.Filters, "mutated")

	second := point15ValEValidFoundationModel()
	for _, value := range second.BlockingReasons {
		if value == "mutated" {
			t.Fatalf("expected fresh blocking reasons")
		}
	}
	for _, value := range second.ClosureEvaluator.CommandsRun {
		if value == "mutated" {
			t.Fatalf("expected fresh closure evaluator commands")
		}
	}
	for _, value := range second.PassClosureManifest.TestsRun {
		if value == "mutated" {
			t.Fatalf("expected fresh pass closure manifest tests")
		}
	}
	for _, value := range second.NoOverclaimFinalCheck.ObservedTexts {
		if value == "mutated" {
			t.Fatalf("expected fresh no overclaim observed texts")
		}
	}
	for _, value := range second.Dependency.Point15ValD.Query.Filters {
		if value == "mutated" {
			t.Fatalf("expected fresh nested vald query filters")
		}
	}
}

func TestPoint10ThroughPoint15ValECurrentSweep(t *testing.T) {
	model := ComputePoint15ValEFoundation(Point15ValEFoundationModel())
	if model.CurrentState != Point15ValEStatePassConfirmed {
		t.Fatalf("expected %s, got %s (blocking=%v review=%v dep=%s fresh=%s downgrade=%s revalidation=%s enforcement=%s projection=%s replay=%s tenant=%s timestamp=%s authority=%s mutation=%s nooverclaim=%s clb=%s evaluator=%s manifest=%s)",
			Point15ValEStatePassConfirmed,
			model.CurrentState,
			model.BlockingReasons,
			model.ReviewPrerequisites,
			model.DependencyState,
			model.FreshnessTaxonomyClosureState,
			model.DowngradeTriggerClosureState,
			model.ScheduledRevalidationClosureState,
			model.EnforcementClosureState,
			model.ProjectionClosureState,
			model.ReplayProofHistoryClosureState,
			model.TenantPrivacyClosureState,
			model.TimestampIntegrityClosureState,
			model.AuthorityBoundaryClosureState,
			model.NoMutationClosureState,
			model.NoOverclaimFinalCheckState,
			model.CLBFinalCheckState,
			model.ClosureEvaluatorState,
			model.PassClosureManifestState,
		)
	}
	if !model.PassClosureManifest.Point15PassAllowed || model.PassClosureManifest.Point15PassToken != point15Val0BlockedPassToken {
		t.Fatalf("expected final point15 pass token in clean vale sweep")
	}
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if string(payload) == "" {
		t.Fatalf("expected non-empty payload")
	}
	if !json.Valid(payload) {
		t.Fatalf("expected valid json payload")
	}
}

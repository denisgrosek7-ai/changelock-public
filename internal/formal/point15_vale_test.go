package formal

import (
	"encoding/json"
	"strings"
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

func point15ValEStringSliceContains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
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
		{"blocks whitespace retagged embedded vald state", func(model *Point15ValEDependencySnapshot) {
			retagged := " " + Point15ValDStateActive + " "
			model.Point15ValDCurrentState = retagged
			model.Point15ValD.CurrentState = retagged
		}, Point15ValEStateBlocked},
		{"blocks tab newline retagged inherited point14 state", func(model *Point15ValEDependencySnapshot) {
			retagged := "\t" + Point14ValEStatePassConfirmed + "\n"
			model.InheritedPoint14ValECurrentState = retagged
			model.Point15ValD.Dependency.InheritedPoint14ValECurrentState = retagged
		}, Point15ValEStateBlocked},
		{"blocks padded inherited tenant scope even when nested snapshot matches", func(model *Point15ValEDependencySnapshot) {
			retagged := " " + model.InheritedTenantScope + " "
			model.InheritedTenantScope = retagged
			model.Point15ValD.Dependency.InheritedTenantScope = retagged
		}, Point15ValEStateBlocked},
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
		{"padded closure evaluator id blocks raw exact identity", func(model *Point15ValEClosureEvaluator) {
			model.ClosureEvaluatorID = " " + model.ClosureEvaluatorID + " "
		}, Point15ValEStateBlocked},
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
		{"missing evidence id blocks", func(model *Point15PassClosureManifest) { model.EvidenceID = "" }, Point15ValEStateBlocked},
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
		{"padded point id blocks", func(model *Point15PassClosureManifest) {
			model.PointID = " " + model.PointID + " "
		}, Point15ValEStateBlocked},
		{"tab newline closure token blocks", func(model *Point15PassClosureManifest) {
			model.ClosureToken = "\t" + point15Val0BlockedPassToken + "\n"
		}, Point15ValEStateBlocked},
		{"padded final point15 pass token blocks", func(model *Point15PassClosureManifest) {
			model.Point15PassToken = point15Val0BlockedPassToken + " "
		}, Point15ValEStateBlocked},
		{"padded tenant scope blocks", func(model *Point15PassClosureManifest) {
			model.TenantScope = " " + model.TenantScope + " "
		}, Point15ValEStateBlocked},
		{"padded evidence hash blocks", func(model *Point15PassClosureManifest) {
			model.EvidenceHash = "\t" + model.EvidenceHash + "\n"
		}, Point15ValEStateBlocked},
		{"padded evidence id blocks raw exact manifest binding", func(model *Point15PassClosureManifest) {
			model.EvidenceID = " " + model.EvidenceID + " "
		}, Point15ValEStateBlocked},
		{"offset generated at blocks raw canonical manifest time", func(model *Point15PassClosureManifest) {
			model.GeneratedAt = "2026-05-07T09:30:00+00:00"
		}, Point15ValEStateBlocked},
		{"non utc offset generated at blocks raw canonical manifest time", func(model *Point15PassClosureManifest) {
			model.GeneratedAt = "2026-05-07T10:30:00+01:00"
		}, Point15ValEStateBlocked},
		{"manifest evidence identity evidence id mismatch blocks", func(model *Point15PassClosureManifest) {
			model.EvidenceIdentity = strings.Replace(model.EvidenceIdentity, "evidence_id="+model.EvidenceID, "evidence_id=evidence:foreign-tenant", 1)
		}, Point15ValEStateBlocked},
		{"manifest evidence identity hash mismatch blocks", func(model *Point15PassClosureManifest) {
			model.EvidenceIdentity = strings.Replace(model.EvidenceIdentity, "evidence_hash="+model.EvidenceHash, "evidence_hash=sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 1)
		}, Point15ValEStateBlocked},
		{"manifest evidence id field mismatch blocks", func(model *Point15PassClosureManifest) {
			model.EvidenceID = "evidence:foreign-tenant"
		}, Point15ValEStateBlocked},
		{"manifest evidence identity extra key blocks", func(model *Point15PassClosureManifest) {
			model.EvidenceIdentity += " extra=authority"
		}, Point15ValEStateBlocked},
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
		{"padded freshness taxonomy state blocks raw exact component", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.FreshnessTaxonomyState = " " + model.FreshnessTaxonomyState + " "
		}, Point15ValEStateBlocked},
		{"padded freshness check id blocks raw exact component", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.CheckID = " " + model.CheckID + " "
		}, Point15ValEStateBlocked},
		{"padded tenant scope blocks raw exact freshness component", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.TenantScope = " " + model.TenantScope + " "
		}, Point15ValEStateBlocked},
		{"tab newline evidence id blocks raw exact component", func(model *Point15ValEFreshnessTaxonomyClosureCheck) {
			model.EvidenceID = "\t" + model.EvidenceID + "\n"
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
		{"padded schedule state blocks raw exact component", func(model *Point15ValEScheduledRevalidationClosureCheck) {
			model.ScheduleState = " " + model.ScheduleState + " "
		}, Point15ValEStateBlocked},
		{"tab newline run result blocks raw exact component", func(model *Point15ValEScheduledRevalidationClosureCheck) {
			model.RunResult = "\t" + model.RunResult + "\n"
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
		{"padded replay ref blocks raw exact closure binding", func(model *Point15ValEReplayProofHistoryClosureCheck) {
			model.ReplayRef = " " + model.ReplayRef + " "
		}, Point15ValEStateBlocked},
		{"tab newline proof pack ref blocks raw exact closure binding", func(model *Point15ValEReplayProofHistoryClosureCheck) {
			model.ProofPackRef = "\t" + model.ProofPackRef + "\n"
		}, Point15ValEStateBlocked},
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
		{"tab newline tenant scope blocks raw exact tenant privacy component", func(model *Point15ValETenantPrivacyClosureCheck) {
			model.TenantScope = "\t" + model.TenantScope + "\n"
		}, Point15ValEStateBlocked},
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
		{"whitespace only enforcement timestamp blocks instead of falling back", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.EnforcementEnforcedAt = " "
		}, Point15ValEStateBlocked},
		{"untrusted displayed time source blocks", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.ProjectionDisplayedTimeSource = point14Val0TimeSourceClientLocal
		}, Point15ValEStateBlocked},
		{"non UTC offset reference time blocks raw exact timestamp closure", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.ReferenceNow = "2026-05-07T10:30:00+01:00"
		}, Point15ValEStateBlocked},
		{"padded scheduled status token blocks raw exact timestamp closure", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.ScheduledStatus = " " + model.ScheduledStatus + " "
		}, Point15ValEStateBlocked},
		{"padded tenant scope blocks raw exact timestamp closure", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.TenantScope = " " + model.TenantScope + " "
		}, Point15ValEStateBlocked},
		{"tab newline run result token blocks raw exact timestamp closure", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.RunResult = "\t" + model.RunResult + "\n"
		}, Point15ValEStateBlocked},
		{"whitespace only revalidation due timestamp blocks raw exact timestamp closure", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.RevalidationRequired = true
			model.RevalidationDueAt = " "
		}, Point15ValEStateBlocked},
		{"whitespace only revalidation completed timestamp blocks raw exact timestamp closure", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.RevalidationCompletedAt = " "
		}, Point15ValEStateBlocked},
		{"whitespace only source event timestamp blocks raw exact timestamp closure", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.SourceEventAt = " "
		}, Point15ValEStateBlocked},
		{"padded source event time source blocks raw exact timestamp closure", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.SourceEventAt = "2026-05-07T10:29:00Z"
			model.SourceEventTimeSource = " " + model.SourceEventTimeSource + " "
		}, Point15ValEStateBlocked},
		{"tab newline source event time source blocks raw exact timestamp closure", func(model *Point15ValETimestampIntegrityClosureCheck) {
			model.SourceEventAt = "2026-05-07T10:29:00Z"
			model.SourceEventTimeSource = "\t" + model.SourceEventTimeSource + "\n"
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
		{"tab newline tenant scope blocks raw exact authority boundary component", func(model *Point15ValEAuthorityBoundaryClosureCheck) {
			model.TenantScope = "\t" + model.TenantScope + "\n"
		}, Point15ValEStateBlocked},
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
		{"inherited point10 readiness overclaim wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "marketplace production ready")
		}, Point15ValEStateBlocked},
		{"unhyphenated inherited readiness overclaim wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "self hosted production approved")
		}, Point15ValEStateBlocked},
		{"unicode dash forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "regulator\u2011approved")
		}, Point15ValEStateBlocked},
		{"slash separator forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "production/approved")
		}, Point15ValEStateBlocked},
		{"dot separator forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "deployment.approved")
		}, Point15ValEStateBlocked},
		{"underscore separator forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "production_approved")
		}, Point15ValEStateBlocked},
		{"intra bucket filler forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "production is approved")
		}, Point15ValEStateBlocked},
		{"long filler forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "production is now fully globally approved")
		}, Point15ValEStateBlocked},
		{"confusable forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "pr\u0254duction approved")
		}, Point15ValEStateBlocked},
		{"zero width forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "pr\u200doduction approved")
		}, Point15ValEStateBlocked},
		{"zero width separator forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "production\u200bapproved")
		}, Point15ValEStateBlocked},
		{"math bold forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "\U0001d429\U0001d42b\U0001d428\U0001d41d\U0001d42e\U0001d41c\U0001d42d\U0001d422\U0001d428\U0001d427 \U0001d41a\U0001d429\U0001d429\U0001d42b\U0001d428\U0001d42f\U0001d41e\U0001d41d")
		}, Point15ValEStateBlocked},
		{"greek nu forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "production appro\u03bded")
		}, Point15ValEStateBlocked},
		{"greek upsilon forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "prod\u03c5ction approved")
		}, Point15ValEStateBlocked},
		{"small cap u forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "prod\U00001d1cction approved")
		}, Point15ValEStateBlocked},
		{"latin upsilon forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "prod\u028action approved")
		}, Point15ValEStateBlocked},
		{"greek delta forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "production approve\u03b4")
		}, Point15ValEStateBlocked},
		{"small cap t forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "official au\U00001d1bhority")
		}, Point15ValEStateBlocked},
		{"latin alpha forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "glob\u0251l truth")
		}, Point15ValEStateBlocked},
		{"latin iota forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "off\u0269cial authority")
		}, Point15ValEStateBlocked},
		{"dental click forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "comp\u01c0iance guaranteed")
		}, Point15ValEStateBlocked},
		{"armenian oh forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "c\u0585mpliance guaranteed")
		}, Point15ValEStateBlocked},
		{"greek eta forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "productio\u03b7 approved")
		}, Point15ValEStateBlocked},
		{"latin eng forbidden wording blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = append(model.ObservedTexts, "productio\u014b approved")
		}, Point15ValEStateBlocked},
		{"internal underscore machine token remains non-boundary safe wording", func(model *Point15ValENoOverclaimFinalCheck) {
			model.InternalDiagnosticTexts = append(model.InternalDiagnosticTexts, "internal_production_approved_metric")
		}, Point15ValEStatePassConfirmed},
		{"split forbidden wording across observed corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"continuous assurance", "guaranteed"}
		}, Point15ValEStateBlocked},
		{"split unhyphenated inherited readiness overclaim blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"air gapped", "certified"}
		}, Point15ValEStateBlocked},
		{"tab newline split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"production\t", "\napproved"}
		}, Point15ValEStateBlocked},
		{"unicode dash split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"deployment\u2013", "approved"}
		}, Point15ValEStateBlocked},
		{"slash split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"deployment/", "approved"}
		}, Point15ValEStateBlocked},
		{"zero width split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"deployment", "appro\u200dved"}
		}, Point15ValEStateBlocked},
		{"word fragment split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"produc", "tion approved"}
		}, Point15ValEStateBlocked},
		{"right leg u split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"prod\uab4e", "ction approved"}
		}, Point15ValEStateBlocked},
		{"latin upsilon split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"prod\u028a", "ction approved"}
		}, Point15ValEStateBlocked},
		{"greek nu split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"production", "appro\u03bded"}
		}, Point15ValEStateBlocked},
		{"greek delta split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"production", "approve\u03b4"}
		}, Point15ValEStateBlocked},
		{"small cap t split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"official au", "\U00001d1bhority"}
		}, Point15ValEStateBlocked},
		{"latin alpha split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"glob\u0251l", "truth"}
		}, Point15ValEStateBlocked},
		{"latin iota split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"off\u0269cial", "authority"}
		}, Point15ValEStateBlocked},
		{"dental click split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"comp\u01c0iance", "guaranteed"}
		}, Point15ValEStateBlocked},
		{"armenian oh split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"c\u0585mpliance", "guaranteed"}
		}, Point15ValEStateBlocked},
		{"armenian vo split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"productio\u0578", "approved"}
		}, Point15ValEStateBlocked},
		{"latin n with long right leg split forbidden wording across corpus blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ObservedTexts = []string{"productio\u019e", "approved"}
		}, Point15ValEStateBlocked},
		{"padded projection disclaimer blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ProjectionDisclaimer = " " + model.ProjectionDisclaimer + " "
		}, Point15ValEStateBlocked},
		{"tab newline projection disclaimer blocks", func(model *Point15ValENoOverclaimFinalCheck) {
			model.ProjectionDisclaimer = "\t" + model.ProjectionDisclaimer + "\n"
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

func TestPoint15ValEComponentAggregateRawExact(t *testing.T) {
	tests := []struct {
		name   string
		states []string
		want   string
	}{
		{
			name:   "empty aggregate blocks instead of pass confirming",
			states: nil,
			want:   Point15ValEStateBlocked,
		},
		{
			name:   "happy path exact pass confirmed states remain pass confirmed",
			states: []string{Point15ValEStatePassConfirmed, Point15ValEStatePassConfirmed},
			want:   Point15ValEStatePassConfirmed,
		},
		{
			name:   "padded pass confirmed state blocks raw exact aggregate",
			states: []string{" " + Point15ValEStatePassConfirmed + " ", Point15ValEStatePassConfirmed},
			want:   Point15ValEStateBlocked,
		},
		{
			name:   "tab newline pass confirmed state blocks raw exact aggregate",
			states: []string{Point15ValEStatePassConfirmed, "\t" + Point15ValEStatePassConfirmed + "\n"},
			want:   Point15ValEStateBlocked,
		},
		{
			name:   "review required state preserves review precedence",
			states: []string{Point15ValEStatePassConfirmed, Point15ValEStateReviewRequired},
			want:   Point15ValEStateReviewRequired,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := point15ValEComponentAggregate(tc.states...); got != tc.want {
				t.Fatalf("expected aggregate %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValEFoundationState(t *testing.T) {
	tests := []struct {
		name        string
		mutate      func(*Point15ValEContinuousVerificationClosureFoundation)
		want        string
		wantReasons []string
	}{
		{"final closure emits point15 pass only on clean vale path", func(*Point15ValEContinuousVerificationClosureFoundation) {}, Point15ValEStatePassConfirmed, nil},
		{"missing projection boundary result blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.ProjectionClosureCheck.DisplayOnly = false
		}, Point15ValEStateBlocked, nil},
		{"padded nested vald dashboard mode blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Dashboard.ProjectionMode = " " + model.Dependency.Point15ValD.Dashboard.ProjectionMode + " "
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"tab newline nested vald query action blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Query.ProjectionAction = "\t" + model.Dependency.Point15ValD.Query.ProjectionAction + "\n"
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"padded nested vald query public visibility blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Query.Visibility = " " + point15ValDVisibilityPublicBlocked + " "
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"padded nested vald no-overclaim disclaimer blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.NoOverclaimGuard.ProjectionDisclaimer = " " + model.Dependency.Point15ValD.NoOverclaimGuard.ProjectionDisclaimer + " "
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"padded nested valb schedule tenant scope blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Schedule.TenantScope = " " + model.Dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Schedule.TenantScope + " "
		}, Point15ValEStateBlocked, []string{"tenant_privacy"}},
		{"tab newline nested vald viewer scope blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.AccessTenantPrivacy.ViewerScope = "\t" + model.Dependency.Point15ValD.AccessTenantPrivacy.ViewerScope + "\n"
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"padded nested vald timestamp projection mode blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.TimestampDisplayDiscipline.ProjectionMode = " " + model.Dependency.Point15ValD.TimestampDisplayDiscipline.ProjectionMode + " "
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"stale nested vald timeline validity creation blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Timeline.TimelineCreatesValidity = true
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"padded nested vald replay ref blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.ReplayProofHistory.ReplayRef = " " + model.Dependency.Point15ValD.ReplayProofHistory.ReplayRef + " "
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"tab newline nested vald proof pack ref blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.ReplayProofHistory.ProofPackRef = "\t" + model.Dependency.Point15ValD.ReplayProofHistory.ProofPackRef + "\n"
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"stale nested vald authority pass allowed blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.AuthorityBoundary.PassAllowed = true
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"stale nested vald dashboard pass authority blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.AuthorityBoundary.DashboardApprovesPass = true
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"stale nested vald evidence mutation blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.NoMutationGuard.EvidenceMutationAttempted = true
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"stale nested vald pass restore blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.NoMutationGuard.PassRestoreAttempted = true
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"stale nested vald no overclaim forbidden wording blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.NoOverclaimGuard.ObservedTexts = append(model.Dependency.Point15ValD.NoOverclaimGuard.ObservedTexts, "production approved")
		}, Point15ValEStateBlocked, []string{"no_overclaim"}},
		{"stale nested vald inherited point10 readiness overclaim blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.NoOverclaimGuard.ObservedTexts = append(model.Dependency.Point15ValD.NoOverclaimGuard.ObservedTexts, "marketplace production ready")
		}, Point15ValEStateBlocked, []string{"no_overclaim"}},
		{"direct vale split no overclaim blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.NoOverclaimFinalCheck.ObservedTexts = []string{"continuous assurance", "guaranteed"}
		}, Point15ValEStateBlocked, []string{"no_overclaim"}},
		{"direct vale tab newline no overclaim blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.NoOverclaimFinalCheck.ObservedTexts = []string{"production\t", "\napproved"}
		}, Point15ValEStateBlocked, []string{"no_overclaim"}},
		{"direct vale zero width split no overclaim blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.NoOverclaimFinalCheck.ObservedTexts = []string{"deployment", "appro\u200dved"}
		}, Point15ValEStateBlocked, []string{"no_overclaim"}},
		{"direct vale confusable no overclaim blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.NoOverclaimFinalCheck.ObservedTexts = []string{"production appro\u03bded"}
		}, Point15ValEStateBlocked, []string{"no_overclaim"}},
		{"direct vale long filler no overclaim blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.NoOverclaimFinalCheck.ObservedTexts = []string{"production is now fully globally approved"}
		}, Point15ValEStateBlocked, []string{"no_overclaim"}},
		{"direct vale internal blocked diagnostic must stay classified in aggregate", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.NoOverclaimFinalCheck.InternalDiagnosticTexts = []string{"public badge"}
			model.NoOverclaimFinalCheck.InternalDiagnosticsClassifiedBlocked = false
		}, Point15ValEStateBlocked, []string{"no_overclaim"}},
		{"stale nested vald query redaction state whitespace retag blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Query.RedactionState = " " + point15ValDRedactionNone + " "
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"stale nested vald query redaction state tab newline retag blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Query.RedactionState = "\t" + point15ValDRedactionNone + "\n"
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"stale nested valc no overclaim forbidden wording blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Dependency.Point15ValC.NoOverclaimGuard.ObservedTexts = append(model.Dependency.Point15ValD.Dependency.Point15ValC.NoOverclaimGuard.ObservedTexts, "deployment approved")
		}, Point15ValEStateBlocked, []string{"no_overclaim"}},
		{"stale nested valc no-overclaim allowed ledger mutation blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Dependency.Point15ValC.NoOverclaimGuard.AllowedSafeWording = append(model.Dependency.Point15ValD.Dependency.Point15ValC.NoOverclaimGuard.AllowedSafeWording, "deployment approved")
		}, Point15ValEStateBlocked, []string{"dependency", "no_overclaim"}},
		{"stale nested valc timestamp ordering mutation blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Dependency.Point15ValC.TimestampDiscipline.ReceivedAt = "2026-05-07T09:06:00Z"
		}, Point15ValEStateBlocked, []string{"dependency"}},
		{"stale nested valb no-overclaim blocked ledger mutation blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.NoOverclaimGuard.BlockedWording = append(model.Dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.NoOverclaimGuard.BlockedWording, "validated revalidation schedule")
		}, Point15ValEStateBlocked, []string{"dependency", "no_overclaim"}},
		{"stale nested vala no-overclaim disclaimer mutation blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.NoOverclaimGuard.TriggerDisclaimer = " " + point15ValATriggerDisclaimer + " "
		}, Point15ValEStateBlocked, []string{"dependency", "no_overclaim"}},
		{"stale nested val0 no-overclaim allowed ledger mutation blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.NoOverclaimGuard.AllowedSafeWording = append(model.Dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.NoOverclaimGuard.AllowedSafeWording, "freshness certified")
		}, Point15ValEStateBlocked, []string{"dependency", "no_overclaim"}},
		{"clb blocker prevents final pass", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.CLBFinalCheck.CLB1Present = true
		}, Point15ValEStateBlocked, nil},
		{"manifest evidence identity drift blocks", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.EvidenceIdentity = "wrong_identity"
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"manifest evidence id field drift blocks final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.EvidenceID = "evidence:foreign-tenant"
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"manifest evidence id missing blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.EvidenceID = ""
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"manifest evidence id padding blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.EvidenceID = " " + model.PassClosureManifest.EvidenceID + " "
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"manifest evidence hash drift blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.EvidenceHash = "sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"manifest policy version padding blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.PolicyVersion = " " + model.PassClosureManifest.PolicyVersion + " "
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"manifest engine version drift blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.EngineVersion = "engine:foreign-version"
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"manifest schema version tab newline blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.SchemaVersion = "\t" + model.PassClosureManifest.SchemaVersion + "\n"
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"manifest evidence identity extra key blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.EvidenceIdentity += " extra=authority"
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"self consistent foreign manifest identity blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.EvidenceID = "evidence:foreign-tenant"
			model.PassClosureManifest.EvidenceHash = "sha256:eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
			model.PassClosureManifest.PolicyVersion = "policy:foreign-version"
			model.PassClosureManifest.EngineVersion = "engine:foreign-version"
			model.PassClosureManifest.SchemaVersion = "schema:foreign-version"
			model.PassClosureManifest.TenantScope = "tenant:foreign-scope"
			model.PassClosureManifest.EvidenceIdentity = point15ValEManifestEvidenceIdentity(
				model.PassClosureManifest.EvidenceID,
				model.PassClosureManifest.EvidenceHash,
				model.PassClosureManifest.PolicyVersion,
				model.PassClosureManifest.EngineVersion,
				model.PassClosureManifest.SchemaVersion,
				model.PassClosureManifest.TenantScope,
			)
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"padded freshness check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.FreshnessTaxonomyClosureCheck.CheckID = " " + model.FreshnessTaxonomyClosureCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"freshness_taxonomy"}},
		{"padded freshness tenant scope blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.FreshnessTaxonomyClosureCheck.TenantScope = " " + model.FreshnessTaxonomyClosureCheck.TenantScope + " "
		}, Point15ValEStateBlocked, []string{"freshness_taxonomy"}},
		{"padded downgrade trigger check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.DowngradeTriggerClosureCheck.CheckID = " " + model.DowngradeTriggerClosureCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"downgrade_trigger"}},
		{"padded scheduled revalidation check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.ScheduledRevalidationClosureCheck.CheckID = " " + model.ScheduledRevalidationClosureCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"scheduled_revalidation"}},
		{"padded enforcement check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.EnforcementClosureCheck.CheckID = " " + model.EnforcementClosureCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"enforcement"}},
		{"padded projection check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.ProjectionClosureCheck.CheckID = " " + model.ProjectionClosureCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"projection"}},
		{"padded replay proof history check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.ReplayProofHistoryClosureCheck.CheckID = " " + model.ReplayProofHistoryClosureCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"replay_proof_history"}},
		{"padded tenant privacy check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TenantPrivacyClosureCheck.CheckID = " " + model.TenantPrivacyClosureCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"tenant_privacy"}},
		{"tab newline tenant privacy tenant scope blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TenantPrivacyClosureCheck.TenantScope = "\t" + model.TenantPrivacyClosureCheck.TenantScope + "\n"
		}, Point15ValEStateBlocked, []string{"tenant_privacy"}},
		{"padded timestamp integrity check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TimestampIntegrityClosureCheck.CheckID = " " + model.TimestampIntegrityClosureCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"timestamp_integrity"}},
		{"padded timestamp integrity tenant scope blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TimestampIntegrityClosureCheck.TenantScope = " " + model.TimestampIntegrityClosureCheck.TenantScope + " "
		}, Point15ValEStateBlocked, []string{"timestamp_integrity"}},
		{"whitespace source event timestamp blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TimestampIntegrityClosureCheck.SourceEventAt = " "
		}, Point15ValEStateBlocked, []string{"timestamp_integrity"}},
		{"padded source event time source blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TimestampIntegrityClosureCheck.SourceEventAt = "2026-05-07T10:29:00Z"
			model.TimestampIntegrityClosureCheck.SourceEventTimeSource = " " + model.TimestampIntegrityClosureCheck.SourceEventTimeSource + " "
		}, Point15ValEStateBlocked, []string{"timestamp_integrity"}},
		{"tab newline source event time source blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TimestampIntegrityClosureCheck.SourceEventAt = "2026-05-07T10:29:00Z"
			model.TimestampIntegrityClosureCheck.SourceEventTimeSource = "\t" + model.TimestampIntegrityClosureCheck.SourceEventTimeSource + "\n"
		}, Point15ValEStateBlocked, []string{"timestamp_integrity"}},
		{"whitespace revalidation completed timestamp blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TimestampIntegrityClosureCheck.RevalidationCompletedAt = " "
		}, Point15ValEStateBlocked, []string{"timestamp_integrity"}},
		{"whitespace revalidation due timestamp blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TimestampIntegrityClosureCheck.RevalidationRequired = true
			model.TimestampIntegrityClosureCheck.RevalidationDueAt = " "
		}, Point15ValEStateBlocked, []string{"timestamp_integrity"}},
		{"upstream whitespace enforcement timestamp blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.Dependency.Point15ValD.Dependency.Point15ValC.TimestampDiscipline.EnforcedAt = " "
		}, Point15ValEStateBlocked, []string{"timestamp_integrity"}},
		{"padded authority boundary check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.AuthorityBoundaryClosureCheck.CheckID = " " + model.AuthorityBoundaryClosureCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"authority_boundary"}},
		{"tab newline authority boundary tenant scope blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.AuthorityBoundaryClosureCheck.TenantScope = "\t" + model.AuthorityBoundaryClosureCheck.TenantScope + "\n"
		}, Point15ValEStateBlocked, []string{"authority_boundary"}},
		{"padded no mutation check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.NoMutationClosureCheck.CheckID = " " + model.NoMutationClosureCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"no_mutation"}},
		{"padded clb check id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.CLBFinalCheck.CheckID = " " + model.CLBFinalCheck.CheckID + " "
		}, Point15ValEStateBlocked, []string{"clb"}},
		{"padded closure evaluator id blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.ClosureEvaluator.ClosureEvaluatorID = " " + model.ClosureEvaluator.ClosureEvaluatorID + " "
		}, Point15ValEStateBlocked, []string{"closure_evaluator"}},
		{"non utc generated at offset blocks aggregate final closure", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			offsetTime := "2026-05-07T10:30:00+01:00"
			model.Dependency.Point15ValD.TimestampDisplayDiscipline.ReferenceNow = offsetTime
			model.TimestampIntegrityClosureCheck.ReferenceNow = offsetTime
			model.PassClosureManifest.GeneratedAt = offsetTime
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"manifest generated at must bind dependency reference now", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.PassClosureManifest.GeneratedAt = "2026-05-07T11:01:00Z"
		}, Point15ValEStateBlocked, []string{"pass_closure_manifest"}},
		{"missing trusted timestamp keeps final closure incomplete", func(model *Point15ValEContinuousVerificationClosureFoundation) {
			model.TimestampIntegrityClosureCheck.ProjectionDisplayedAt = ""
		}, Point15ValEStateIncomplete, nil},
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
			for _, expected := range tc.wantReasons {
				if !point15ValEStringSliceContains(result.BlockingReasons, expected) {
					t.Fatalf("expected exact blocking reason %q, got %#v", expected, result.BlockingReasons)
				}
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

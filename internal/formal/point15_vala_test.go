package formal

import (
	"strings"
	"sync"
	"testing"
)

var (
	point15ValAFoundationOnce sync.Once
	point15ValAFoundationBase Point15ValADowngradeTriggerFoundation
)

func point15ValACloneStrings(values []string) []string {
	return append([]string(nil), values...)
}

func clonePoint15ValAFoundation(model Point15ValADowngradeTriggerFoundation) Point15ValADowngradeTriggerFoundation {
	model.BlockingReasons = point15ValACloneStrings(model.BlockingReasons)
	model.ReviewPrerequisites = point15ValACloneStrings(model.ReviewPrerequisites)
	model.Dependency.ReviewPrerequisites = point15ValACloneStrings(model.Dependency.ReviewPrerequisites)
	model.Dependency.Point15Val0 = clonePoint15Val0Foundation(model.Dependency.Point15Val0)
	model.TriggerTable.AllowedTriggers = point15ValACloneStrings(model.TriggerTable.AllowedTriggers)
	model.NoOverclaimGuard.ObservedTexts = point15ValACloneStrings(model.NoOverclaimGuard.ObservedTexts)
	model.NoOverclaimGuard.InternalDiagnosticTexts = point15ValACloneStrings(model.NoOverclaimGuard.InternalDiagnosticTexts)
	model.NoOverclaimGuard.AllowedSafeWording = point15ValACloneStrings(model.NoOverclaimGuard.AllowedSafeWording)
	model.NoOverclaimGuard.BlockedWording = point15ValACloneStrings(model.NoOverclaimGuard.BlockedWording)
	return model
}

func point15ValAValidFoundationModel() Point15ValADowngradeTriggerFoundation {
	point15ValAFoundationOnce.Do(func() {
		point15ValAFoundationBase = Point15ValAFoundationModel()
	})
	return clonePoint15ValAFoundation(point15ValAFoundationBase)
}

func point15ValAFoundationWithTrigger(trigger string, decisive bool, lineage string) Point15ValADowngradeTriggerFoundation {
	model := point15ValAValidFoundationModel()
	if strings.TrimSpace(trigger) == "" {
		return model
	}
	lineageValid := point15Val0LineageRefValid(lineage)
	targetState := point15ValATriggerExpectedState(trigger, decisive, lineageValid)
	targetOutcome := point15ValATriggerExpectedOutcome(trigger, decisive, lineageValid)
	observedFreshness := point15ValATriggerObservedFreshnessStatus(trigger)

	model.TriggerTable.CurrentTriggerDetected = true
	model.TriggerTable.CurrentTriggerRef = model.Trigger.TriggerID
	model.TriggerTable.CurrentReasonRef = model.Reason.ReasonID
	model.TriggerTable.CurrentDecisionRef = model.Decision.DecisionID
	model.TriggerTable.CurrentTriggerType = trigger
	model.TriggerTable.CurrentTargetState = targetState
	model.TriggerTable.CurrentDowngradeOutcome = targetOutcome

	model.Trigger.TriggerDetected = true
	model.Trigger.TriggerType = trigger
	model.Trigger.ObservedFreshnessStatus = observedFreshness
	model.Trigger.TriggerIsDecisive = decisive
	model.Trigger.SupersessionLineageRef = lineage
	model.Trigger.TargetState = targetState
	model.Trigger.TargetDowngradeOutcome = targetOutcome
	model.Trigger.RetainsActiveClosure = false

	model.Reason.TriggerType = trigger
	model.Reason.ReasonCode = point15ValAExpectedReasonCode(trigger, lineageValid)
	model.Reason.ObservedFreshnessStatus = observedFreshness
	model.Reason.Decisive = decisive
	model.Reason.SupersessionLineageRef = lineage
	model.Reason.TargetState = targetState
	model.Reason.TargetDowngradeOutcome = targetOutcome

	model.Decision.TriggerDetected = true
	model.Decision.TriggerRef = model.Trigger.TriggerID
	model.Decision.ReasonRef = model.Reason.ReasonID
	model.Decision.TriggerType = trigger
	model.Decision.TargetState = targetState
	model.Decision.TargetDowngradeOutcome = targetOutcome
	model.Decision.RetainsActiveClosure = false

	return model
}

func TestPoint15ValADependencyState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValADependencySnapshot)
		want   string
	}{
		{"active when point15 val0 foundation is clean", func(model *Point15ValADependencySnapshot) {}, Point15ValAStateActive},
		{"blocks when point15 val0 missing", func(model *Point15ValADependencySnapshot) { model.Point15Val0CurrentState = "" }, Point15ValAStateBlocked},
		{"blocks when point15 val0 blocked", func(model *Point15ValADependencySnapshot) { model.Point15Val0CurrentState = Point15Val0StateBlocked }, Point15ValAStateBlocked},
		{"blocks when point15 val0 review required", func(model *Point15ValADependencySnapshot) {
			model.Point15Val0CurrentState = Point15Val0StateReviewRequired
		}, Point15ValAStateBlocked},
		{"blocks when point15 val0 incomplete", func(model *Point15ValADependencySnapshot) { model.Point15Val0CurrentState = Point15Val0StateIncomplete }, Point15ValAStateBlocked},
		{"blocks when point15 val0 not merged", func(model *Point15ValADependencySnapshot) { model.Point15Val0Merged = false }, Point15ValAStateBlocked},
		{"blocks when point15 val0 ci not green", func(model *Point15ValADependencySnapshot) { model.Point15Val0CIGreen = false }, Point15ValAStateBlocked},
		{"blocks when point15 val0 not reviewed on main", func(model *Point15ValADependencySnapshot) { model.Point15Val0ReviewedOnMain = false }, Point15ValAStateBlocked},
		{"blocks when embedded point15 val0 snapshot is not computed from upstream", func(model *Point15ValADependencySnapshot) {
			model.Point15Val0ComputedFromUpstream = true
			model.Point15Val0.Dependency.SnapshotFromComputedOutput = false
		}, Point15ValAStateBlocked},
		{"blocks padded point15 val0 current state raw exact", func(model *Point15ValADependencySnapshot) {
			model.Point15Val0CurrentState = " " + Point15Val0StateActive + " "
			model.Point15Val0.CurrentState = model.Point15Val0CurrentState
		}, Point15ValAStateBlocked},
		{"blocks tab newline inherited point14 pass state raw exact", func(model *Point15ValADependencySnapshot) {
			model.InheritedPoint14ValECurrentState = "\t" + Point14ValEStatePassConfirmed + "\n"
			model.Point15Val0.Dependency.Point14ValECurrentState = model.InheritedPoint14ValECurrentState
		}, Point15ValAStateBlocked},
		{"blocks padded inherited tenant scope raw exact", func(model *Point15ValADependencySnapshot) {
			model.InheritedTenantScope = " " + model.InheritedTenantScope + " "
			model.Point15Val0.Dependency.InheritedTenantScope = model.InheritedTenantScope
		}, Point15ValAStateBlocked},
		{"blocks stale embedded val0 timestamp ordering mutation", func(model *Point15ValADependencySnapshot) {
			model.Point15Val0.TimestampDiscipline.ObservedAt = "2026-05-06T18:30:00Z"
		}, Point15ValAStateBlocked},
		{"blocks stale embedded val0 no-overclaim allowed ledger mutation", func(model *Point15ValADependencySnapshot) {
			model.Point15Val0.NoOverclaimGuard.AllowedSafeWording = append(model.Point15Val0.NoOverclaimGuard.AllowedSafeWording, "freshness certified")
		}, Point15ValAStateBlocked},
		{"blocks when point15 pass already appears", func(model *Point15ValADependencySnapshot) { model.Point15PassSeen = true }, Point15ValAStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValAValidFoundationModel().Dependency
			tc.mutate(&model)
			if got := EvaluatePoint15ValADependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint15ValADowngradeTriggerFoundationState(t *testing.T) {
	tests := []struct {
		name  string
		model func() Point15ValADowngradeTriggerFoundation
		want  string
	}{
		{"happy path active with complete trigger table and no trigger", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAValidFoundationModel()
		}, Point15ValAStateActive},
		{"expired evidence trigger blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerExpired, false, "")
		}, Point15ValAStateBlocked},
		{"revoked signal trigger blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerRevoked, false, "")
		}, Point15ValAStateBlocked},
		{"stale evidence trigger requires review", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerStale, false, "")
		}, Point15ValAStateReviewRequired},
		{"superseded evidence with lineage requires review", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerSuperseded, false, "supersession_lineage_point15_vala_001")
		}, Point15ValAStateReviewRequired},
		{"superseded evidence without lineage blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerSuperseded, false, "")
		}, Point15ValAStateBlocked},
		{"policy drift non decisive requires review", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerPolicyDrift, false, "")
		}, Point15ValAStateReviewRequired},
		{"policy drift decisive blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerPolicyDrift, true, "")
		}, Point15ValAStateBlocked},
		{"artifact drift non decisive requires review", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerArtifact, false, "")
		}, Point15ValAStateReviewRequired},
		{"artifact drift decisive blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerArtifact, true, "")
		}, Point15ValAStateBlocked},
		{"verifier drift non decisive requires review", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerVerifier, false, "")
		}, Point15ValAStateReviewRequired},
		{"verifier drift decisive blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerVerifier, true, "")
		}, Point15ValAStateBlocked},
		{"connector failure requires review", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerConnFail, false, "")
		}, Point15ValAStateReviewRequired},
		{"connector timeout requires review", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerConnTimeout, false, "")
		}, Point15ValAStateReviewRequired},
		{"connector unauthorized blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerConnAuth, false, "")
		}, Point15ValAStateBlocked},
		{"connector tenant mismatch blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerConnTenant, false, "")
		}, Point15ValAStateBlocked},
		{"tampered freshness proof blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerTampered, false, "")
		}, Point15ValAStateBlocked},
		{"unsupported freshness status blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerUnsupported, false, "")
		}, Point15ValAStateBlocked},
		{"missing freshness proof non decisive is incomplete", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerMissing, false, "")
		}, Point15ValAStateIncomplete},
		{"missing freshness proof decisive blocks", func() Point15ValADowngradeTriggerFoundation {
			return point15ValAFoundationWithTrigger(point15ValATriggerMissing, true, "")
		}, Point15ValAStateBlocked},
		{"pass preservation forbidden for any downgrade trigger", func() Point15ValADowngradeTriggerFoundation {
			model := point15ValAFoundationWithTrigger(point15ValATriggerExpired, false, "")
			model.Trigger.RetainsPass = true
			return model
		}, Point15ValAStateBlocked},
		{"active closure retention forbidden when trigger exists", func() Point15ValADowngradeTriggerFoundation {
			model := point15ValAFoundationWithTrigger(point15ValATriggerStale, false, "")
			model.Trigger.RetainsActiveClosure = true
			return model
		}, Point15ValAStateBlocked},
		{"foundation blocks on trigger table reason state mismatch", func() Point15ValADowngradeTriggerFoundation {
			model := point15ValAFoundationWithTrigger(point15ValATriggerStale, false, "")
			model.Reason.TargetState = Point15Val0StateBlocked
			model.Reason.TargetDowngradeOutcome = point15Val0DowngradeBlocked
			return model
		}, Point15ValAStateBlocked},
		{"foundation blocks when expired trigger targets review required", func() Point15ValADowngradeTriggerFoundation {
			model := point15ValAFoundationWithTrigger(point15ValATriggerExpired, false, "")
			model.Trigger.TargetState = Point15Val0StateReviewRequired
			model.Trigger.TargetDowngradeOutcome = point15Val0DowngradeReview
			return model
		}, Point15ValAStateBlocked},
		{"foundation blocks when stale trigger targets active", func() Point15ValADowngradeTriggerFoundation {
			model := point15ValAFoundationWithTrigger(point15ValATriggerStale, false, "")
			model.Trigger.TargetState = Point15Val0StateActive
			model.Trigger.TargetDowngradeOutcome = point15Val0DowngradeRetainActive
			model.Trigger.RetainsActiveClosure = true
			return model
		}, Point15ValAStateBlocked},
		{"blocks whitespace-only optional trigger table ref raw exact", func() Point15ValADowngradeTriggerFoundation {
			model := point15ValAValidFoundationModel()
			model.TriggerTable.CurrentTriggerRef = " "
			return model
		}, Point15ValAStateBlocked},
		{"blocks hard invalid optional decision ref raw exact", func() Point15ValADowngradeTriggerFoundation {
			model := point15ValAValidFoundationModel()
			model.Decision.ReasonRef = "\t\n"
			return model
		}, Point15ValAStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := tc.model()
			computed := ComputePoint15ValADowngradeTriggerFoundation(model)
			if computed.CurrentState != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, computed.CurrentState)
			}
		})
	}

	t.Run("stale embedded val0 timestamp mutation records exact dependency reason", func(t *testing.T) {
		model := point15ValAValidFoundationModel()
		model.Dependency.Point15Val0.TimestampDiscipline.ObservedAt = "2026-05-06T18:30:00Z"
		computed := ComputePoint15ValADowngradeTriggerFoundation(model)
		if computed.CurrentState != Point15ValAStateBlocked {
			t.Fatalf("expected stale embedded Val0 timestamp mutation to block, got %#v", computed)
		}
		if !point15ValDStringSliceContains(computed.BlockingReasons, "dependency") {
			t.Fatalf("expected exact dependency blocking reason, got %#v", computed.BlockingReasons)
		}
	})

	t.Run("stale embedded val0 no overclaim mutation records exact dependency reason", func(t *testing.T) {
		model := point15ValAValidFoundationModel()
		model.Dependency.Point15Val0.NoOverclaimGuard.AllowedSafeWording = append(model.Dependency.Point15Val0.NoOverclaimGuard.AllowedSafeWording, "freshness certified")
		computed := ComputePoint15ValADowngradeTriggerFoundation(model)
		if computed.CurrentState != Point15ValAStateBlocked {
			t.Fatalf("expected stale embedded Val0 no-overclaim mutation to block, got %#v", computed)
		}
		if !point15ValDStringSliceContains(computed.BlockingReasons, "dependency") {
			t.Fatalf("expected exact dependency blocking reason, got %#v", computed.BlockingReasons)
		}
	})
}

func TestPoint15ValAAuthorityBoundaryState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValAAuthorityBoundary)
	}{
		{"scheduler cannot map trigger to blocked", func(model *Point15ValAAuthorityBoundary) { model.SchedulerMapsTriggerToDowngrade = true }},
		{"dashboard cannot map trigger to blocked", func(model *Point15ValAAuthorityBoundary) { model.DashboardMapsTriggerToDowngrade = true }},
		{"connector cannot map trigger to blocked", func(model *Point15ValAAuthorityBoundary) { model.ConnectorMapsTriggerToDowngrade = true }},
		{"agent cannot map trigger to blocked", func(model *Point15ValAAuthorityBoundary) { model.AgentMapsTriggerToDowngrade = true }},
		{"customer projection cannot mutate downgrade", func(model *Point15ValAAuthorityBoundary) { model.CustomerProjectionMutatesDowngrade = true }},
		{"auditor projection cannot mutate downgrade", func(model *Point15ValAAuthorityBoundary) { model.AuditorProjectionMutatesDowngrade = true }},
		{"portal projection cannot mutate downgrade", func(model *Point15ValAAuthorityBoundary) { model.PortalProjectionMutatesDowngrade = true }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValAAuthorityBoundaryModel(point15ValADependencySnapshotModel())
			tc.mutate(&model)
			if got := EvaluatePoint15ValAAuthorityBoundaryState(model); got != Point15ValAStateBlocked {
				t.Fatalf("expected %s, got %s", Point15ValAStateBlocked, got)
			}
		})
	}
}

func TestPoint15ValANoOverclaimGuardState(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Point15ValANoOverclaimGuard)
		want   string
	}{
		{"safe bounded wording passes", func(model *Point15ValANoOverclaimGuard) {}, Point15ValAStateActive},
		{"forbidden public wording blocks", func(model *Point15ValANoOverclaimGuard) {
			model.ObservedTexts = []string{"continuous assurance guaranteed"}
		}, Point15ValAStateBlocked},
		{"split forbidden public wording blocks", func(model *Point15ValANoOverclaimGuard) {
			model.ObservedTexts = []string{"continuous assurance", "guaranteed"}
		}, Point15ValAStateBlocked},
		{"confusable forbidden public wording blocks", func(model *Point15ValANoOverclaimGuard) {
			model.ObservedTexts = []string{"production appro\u03bded"}
		}, Point15ValAStateBlocked},
		{"unclassified internal diagnostic with forbidden wording blocks", func(model *Point15ValANoOverclaimGuard) {
			model.InternalDiagnosticTexts = []string{"autonomous assurance pass"}
		}, Point15ValAStateBlocked},
		{"classified blocked diagnostic remains internal only", func(model *Point15ValANoOverclaimGuard) {
			model.InternalDiagnosticTexts = []string{"autonomous assurance pass"}
			model.InternalDiagnosticsClassifiedBlocked = true
		}, Point15ValAStateActive},
		{"padded trigger disclaimer blocks raw exact", func(model *Point15ValANoOverclaimGuard) {
			model.TriggerDisclaimer = " " + point15ValATriggerDisclaimer + " "
		}, Point15ValAStateBlocked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point15ValANoOverclaimGuardModel()
			tc.mutate(&model)
			if got := EvaluatePoint15ValANoOverclaimGuardState(model); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

func TestPoint10ThroughPoint15ValACurrentSweep(t *testing.T) {
	computed := ComputePoint15ValADowngradeTriggerFoundation(point15ValAValidFoundationModel())
	if computed.DependencyState != Point15ValAStateActive {
		t.Fatalf("expected dependency active, got %s", computed.DependencyState)
	}
	if computed.CurrentState != Point15ValAStateActive {
		t.Fatalf("expected current state active, got %s", computed.CurrentState)
	}
	if computed.Dependency.Point15PassSeen {
		t.Fatal("expected no point_15_pass in point15 val a sweep")
	}
}

func TestPoint15ValACachedHelperIsolation(t *testing.T) {
	model := point15ValAValidFoundationModel()
	originalAllowed := model.NoOverclaimGuard.AllowedSafeWording[0]
	model.NoOverclaimGuard.AllowedSafeWording[0] = "mutated"

	fresh := point15ValAValidFoundationModel()
	if fresh.NoOverclaimGuard.AllowedSafeWording[0] != originalAllowed {
		t.Fatalf("expected cached point15 vala helper to return isolated copy, got %#v", fresh.NoOverclaimGuard.AllowedSafeWording)
	}
}

func TestPoint15ValACachedHelperNestedDependencyIsolation(t *testing.T) {
	model := point15ValAValidFoundationModel()
	originalAllowed := model.Dependency.Point15Val0.NoOverclaimGuard.AllowedSafeWording[0]
	model.Dependency.Point15Val0.NoOverclaimGuard.AllowedSafeWording[0] = "mutated"

	fresh := point15ValAValidFoundationModel()
	if fresh.Dependency.Point15Val0.NoOverclaimGuard.AllowedSafeWording[0] != originalAllowed {
		t.Fatalf("expected cached point15 vala nested dependency helper to return isolated copy, got %#v", fresh.Dependency.Point15Val0.NoOverclaimGuard.AllowedSafeWording)
	}
}

func TestPoint15ValAAggregateRawExact(t *testing.T) {
	tests := []struct {
		name   string
		states []string
		want   string
	}{
		{"happy path active", []string{Point15ValAStateActive, Point15ValAStateActive}, Point15ValAStateActive},
		{"direct exploit padded active blocks", []string{Point15ValAStateActive, " " + Point15ValAStateActive + " "}, Point15ValAStateBlocked},
		{"hard invalid tab newline active blocks", []string{Point15ValAStateActive, "\t" + Point15ValAStateActive + "\n"}, Point15ValAStateBlocked},
		{"sibling review path preserved", []string{Point15ValAStateActive, Point15ValAStateReviewRequired}, Point15ValAStateReviewRequired},
		{"blocked path preserved", []string{Point15ValAStateActive, Point15ValAStateBlocked}, Point15ValAStateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := point15ValAAggregate(tc.states...); got != tc.want {
				t.Fatalf("expected %s, got %s", tc.want, got)
			}
		})
	}
}

package formal

import (
	"encoding/json"
	"strings"
	"sync"
	"testing"
)

var (
	point16Val0FoundationOnce sync.Once
	point16Val0FoundationBase Point16Val0Foundation
)

func point16Val0CloneStrings(values []string) []string {
	return append([]string(nil), values...)
}

func clonePoint16Val0Foundation(model Point16Val0Foundation) Point16Val0Foundation {
	model.BlockingReasons = point16Val0CloneStrings(model.BlockingReasons)
	model.ReviewPrerequisites = point16Val0CloneStrings(model.ReviewPrerequisites)
	model.Dependency.ReviewPrerequisites = point16Val0CloneStrings(model.Dependency.ReviewPrerequisites)
	model.Dependency.Point15ValE = clonePoint15ValEFoundation(model.Dependency.Point15ValE)
	model.ReplayTaxonomy.AllowedStatuses = point16Val0CloneStrings(model.ReplayTaxonomy.AllowedStatuses)
	model.NoOverclaimBaseline.ObservedTexts = point16Val0CloneStrings(model.NoOverclaimBaseline.ObservedTexts)
	model.NoOverclaimBaseline.InternalDiagnosticTexts = point16Val0CloneStrings(model.NoOverclaimBaseline.InternalDiagnosticTexts)
	model.NoOverclaimBaseline.AllowedSafeWording = point16Val0CloneStrings(model.NoOverclaimBaseline.AllowedSafeWording)
	model.NoOverclaimBaseline.BlockedWording = point16Val0CloneStrings(model.NoOverclaimBaseline.BlockedWording)
	return model
}

func point16Val0ValidFoundationModel() Point16Val0Foundation {
	point16Val0FoundationOnce.Do(func() {
		point16Val0FoundationBase = ComputePoint16Val0Foundation(Point16Val0FoundationModel())
	})
	return clonePoint16Val0Foundation(point16Val0FoundationBase)
}

func point16Val0ValidDependencyModel() Point16Val0DependencySnapshot {
	return point16Val0ValidFoundationModel().Dependency
}

func point16Val0ValidHistoricalReplayContextModel() Point16Val0HistoricalReplayContext {
	return point16Val0ValidFoundationModel().HistoricalReplayContext
}

func point16Val0ValidOriginalDecisionBindingModel() Point16Val0OriginalDecisionBinding {
	return point16Val0ValidFoundationModel().OriginalDecisionBinding
}

func point16Val0ValidReplayTaxonomyModel() Point16Val0ReplayTaxonomy {
	return point16Val0ValidFoundationModel().ReplayTaxonomy
}

func point16Val0ValidCurrentSubstitutionGuardModel() Point16Val0CurrentSubstitutionGuard {
	return point16Val0ValidFoundationModel().CurrentSubstitutionGuard
}

func point16Val0ValidReplayReadinessEvaluationModel() Point16Val0ReplayReadinessEvaluation {
	model := point16Val0ReplayReadinessEvaluationModel()
	model.DependencyState = Point16Val0StateActive
	model.HistoricalReplayContextState = Point16Val0StateActive
	model.OriginalDecisionBindingState = Point16Val0StateActive
	model.ReplayTaxonomyState = Point16Val0StateActive
	model.CurrentSubstitutionGuardState = Point16Val0StateActive
	model.NoOverclaimState = Point16Val0StateActive
	return model
}

func point16Val0ValidNoOverclaimBaselineModel() Point16Val0NoOverclaimBaseline {
	return point16Val0ValidFoundationModel().NoOverclaimBaseline
}

func TestPoint16Val0DependencyState(t *testing.T) {
	t.Run("clean point15 closure dependency is active", func(t *testing.T) {
		model := point16Val0ValidDependencyModel()
		if got := EvaluatePoint16Val0DependencyState(model); got != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s", Point16Val0StateActive, got)
		}
	})

	tests := []struct {
		name   string
		mutate func(*Point16Val0DependencySnapshot)
	}{
		{name: "missing merged point15 vale blocks", mutate: func(model *Point16Val0DependencySnapshot) { model.Point15ValEMerged = false }},
		{name: "non-pass point15 vale current state blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15ValECurrentState = Point15ValEStateReviewRequired
		}},
		{name: "missing point15 pass allowance blocks", mutate: func(model *Point16Val0DependencySnapshot) { model.Point15PassAllowed = false }},
		{name: "wrong point15 closure token blocks", mutate: func(model *Point16Val0DependencySnapshot) { model.Point15PassToken = "" }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point16Val0ValidDependencyModel()
			tc.mutate(&model)
			if got := EvaluatePoint16Val0DependencyState(model); got != Point16Val0StateBlocked {
				t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
			}
		})
	}
}

func TestPoint16Val0HistoricalReplayContextState(t *testing.T) {
	t.Run("clean original historical context is active readiness", func(t *testing.T) {
		model := point16Val0ValidHistoricalReplayContextModel()
		if got := EvaluatePoint16Val0HistoricalReplayContextState(model); got != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s", Point16Val0StateActive, got)
		}
	})

	tests := []struct {
		name     string
		mutate   func(*Point16Val0HistoricalReplayContext)
		expected string
	}{
		{name: "missing original evidence fails closed", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.OriginalEvidenceID = ""
			model.OriginalEvidenceHash = ""
		}, expected: Point16Val0StateIncomplete},
		{name: "missing original policy identity version hash fails closed", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.OriginalPolicyID = ""
			model.OriginalPolicyVersion = ""
			model.OriginalPolicyHash = ""
		}, expected: Point16Val0StateIncomplete},
		{name: "missing original engine identity version hash fails closed", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.OriginalEngineID = ""
			model.OriginalEngineVersion = ""
			model.OriginalEngineHash = ""
		}, expected: Point16Val0StateIncomplete},
		{name: "timestamp unsafe blocks", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.ReplayTimeSource = "client_local_time"
		}, expected: Point16Val0StateBlocked},
		{name: "timestamp backdating blocks", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.ReplayAt = "2026-05-08T08:59:00Z"
		}, expected: Point16Val0StateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point16Val0ValidHistoricalReplayContextModel()
			tc.mutate(&model)
			if got := EvaluatePoint16Val0HistoricalReplayContextState(model); got != tc.expected {
				t.Fatalf("expected %s, got %s", tc.expected, got)
			}
		})
	}
}

func TestPoint16Val0ReplayTaxonomyState(t *testing.T) {
	t.Run("clean replay taxonomy is active", func(t *testing.T) {
		model := point16Val0ValidReplayTaxonomyModel()
		if got := EvaluatePoint16Val0ReplayTaxonomyState(model); got != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s", Point16Val0StateActive, got)
		}
	})

	tests := []struct {
		name     string
		mutate   func(*Point16Val0ReplayTaxonomy)
		expected string
	}{
		{name: "evidence hash mismatch blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.EvidenceHashMatches = false
			model.ReplayStatus = point16Val0EvidenceHashMismatch
		}, expected: Point16Val0StateBlocked},
		{name: "policy hash mismatch review requires", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.PolicyHashMatches = false
			model.ReplayStatus = point16Val0PolicyHashMismatch
		}, expected: Point16Val0StateReviewRequired},
		{name: "engine hash mismatch review requires", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.EngineHashMatches = false
			model.ReplayStatus = point16Val0EngineHashMismatch
		}, expected: Point16Val0StateReviewRequired},
		{name: "tenant scope mismatch blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.TenantScopeMatches = false
			model.ReplayStatus = point16Val0TenantScopeMismatch
		}, expected: Point16Val0StateBlocked},
		{name: "artifact scope mismatch review requires", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.ArtifactScopeMatches = false
			model.ReplayStatus = point16Val0ArtifactScopeMismatch
		}, expected: Point16Val0StateReviewRequired},
		{name: "claim scope mismatch review requires", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.ClaimScopeMatches = false
			model.ReplayStatus = point16Val0ClaimScopeMismatch
		}, expected: Point16Val0StateReviewRequired},
		{name: "governance scope mismatch review requires", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.GovernanceScopeMatches = false
			model.ReplayStatus = point16Val0GovernanceScopeMismatch
		}, expected: Point16Val0StateReviewRequired},
		{name: "unsupported replay blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.ReplaySupported = false
			model.ReplayStatus = point16Val0UnsupportedReplay
		}, expected: Point16Val0StateBlocked},
		{name: "tampered history blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.HistoryTampered = true
			model.ReplayStatus = point16Val0TamperedHistory
		}, expected: Point16Val0StateBlocked},
		{name: "missing lineage fails closed", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.LineagePresent = false
			model.ReplayStatus = point16Val0LineageMissing
		}, expected: Point16Val0StateIncomplete},
		{name: "current policy substitution attempt blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.CurrentPolicySubstitutionAttempted = true
			model.ReplayStatus = point16Val0CurrentPolicySubstitution
		}, expected: Point16Val0StateBlocked},
		{name: "current engine substitution attempt blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.CurrentEngineSubstitutionAttempted = true
			model.ReplayStatus = point16Val0CurrentEngineSubstitution
		}, expected: Point16Val0StateBlocked},
		{name: "current evidence substitution attempt blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.CurrentEvidenceSubstitutionAttempted = true
			model.ReplayStatus = point16Val0CurrentEvidenceSubstitution
		}, expected: Point16Val0StateBlocked},
		{name: "current time substitution attempt blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.CurrentTimeSubstitutionAttempted = true
			model.ReplayStatus = point16Val0CurrentTimeSubstitution
		}, expected: Point16Val0StateBlocked},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point16Val0ValidReplayTaxonomyModel()
			tc.mutate(&model)
			if got := EvaluatePoint16Val0ReplayTaxonomyState(model); got != tc.expected {
				t.Fatalf("expected %s, got %s", tc.expected, got)
			}
		})
	}
}

func TestPoint16Val0CurrentSubstitutionGuardState(t *testing.T) {
	t.Run("clean current substitution guard is active", func(t *testing.T) {
		model := point16Val0ValidCurrentSubstitutionGuardModel()
		if got := EvaluatePoint16Val0CurrentSubstitutionGuardState(model); got != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s", Point16Val0StateActive, got)
		}
	})

	tests := []struct {
		name   string
		mutate func(*Point16Val0CurrentSubstitutionGuard)
	}{
		{name: "current policy substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentPolicyAuthoritative = true }},
		{name: "current engine substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentEngineAuthoritative = true }},
		{name: "current evidence substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentEvidenceAuthoritative = true }},
		{name: "current timestamp substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentTimestampAuthoritative = true }},
		{name: "current tenant substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentTenantAuthoritative = true }},
		{name: "current claim substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentClaimAuthoritative = true }},
		{name: "current governance substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentGovernanceAuthoritative = true }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point16Val0ValidCurrentSubstitutionGuardModel()
			tc.mutate(&model)
			if got := EvaluatePoint16Val0CurrentSubstitutionGuardState(model); got != Point16Val0StateBlocked {
				t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
			}
		})
	}
}

func TestPoint16Val0ReplayReadinessEvaluationState(t *testing.T) {
	t.Run("clean replay readiness stays active and never emits point16 pass", func(t *testing.T) {
		model := point16Val0ValidReplayReadinessEvaluationModel()
		if got := EvaluatePoint16Val0ReplayReadinessState(model); got != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s", Point16Val0StateActive, got)
		}
	})

	tests := []struct {
		name   string
		mutate func(*Point16Val0ReplayReadinessEvaluation)
	}{
		{name: "no mutation publication revocation external api default connector path introduced blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoMutationPathsDetected = false
		}},
		{name: "publication path blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoPublicationPathDetected = false
		}},
		{name: "revocation execution path blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoRevocationExecutionDetected = false
		}},
		{name: "evidence deletion path blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoEvidenceDeletionDetected = false
		}},
		{name: "external api default path blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoExternalAPIDefaultDetected = false
		}},
		{name: "connector mutation path blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoConnectorMutationDetected = false
		}},
		{name: "scheduler path blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoSchedulerPathDetected = false
		}},
		{name: "customer authority path blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoCustomerAuthorityDetected = false
		}},
		{name: "auditor authority path blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoAuditorAuthorityDetected = false
		}},
		{name: "portal authority path blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoPortalAuthorityDetected = false
		}},
		{name: "ai agent authority path blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoAIAgentAuthorityDetected = false
		}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point16Val0ValidReplayReadinessEvaluationModel()
			tc.mutate(&model)
			if got := EvaluatePoint16Val0ReplayReadinessState(model); got != Point16Val0StateBlocked {
				t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
			}
		})
	}
}

func TestPoint16Val0NoOverclaimBaselineState(t *testing.T) {
	t.Run("safe bounded wording passes", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s", Point16Val0StateActive, got)
		}
	})

	t.Run("forbidden wording blocks", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.ObservedTexts = []string{"certified replay correctness"}
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateBlocked {
			t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
		}
	})

	t.Run("mutated safe wording set blocks", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.AllowedSafeWording = []string{"mutated"}
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateBlocked {
			t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
		}
	})
}

func TestPoint16Val0Foundation(t *testing.T) {
	t.Run("clean original historical context produces readiness not point16 pass", func(t *testing.T) {
		model := ComputePoint16Val0Foundation(point16Val0ValidFoundationModel())
		if model.CurrentState != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s (blocking=%v review=%v dep=%s ctx=%s binding=%s taxonomy=%s guard=%s readiness=%s nooverclaim=%s)",
				Point16Val0StateActive,
				model.CurrentState,
				model.BlockingReasons,
				model.ReviewPrerequisites,
				model.DependencyState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
				model.NoOverclaimState,
			)
		}
		payload, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		body := string(payload)
		if strings.Contains(body, "point_16_pass") || strings.Contains(body, "point16_pass") {
			t.Fatalf("expected no point16 pass token anywhere in production model payload")
		}
	})

	t.Run("point15 vale dependency must be present and clean", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.Dependency.Point15PassAllowed = false
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked || model.DependencyState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked dependency, got current=%s dependency=%s", model.CurrentState, model.DependencyState)
		}
	})

	t.Run("current policy substitution attempt fails closed in full foundation", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.CurrentSubstitutionGuard.CurrentPolicyAuthoritative = true
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked || model.CurrentSubstitutionGuardState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked substitution guard, got current=%s guard=%s", model.CurrentState, model.CurrentSubstitutionGuardState)
		}
	})

	t.Run("tampered substitution guard original decision timestamp fails closed", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.CurrentSubstitutionGuard.OriginalDecisionAt = "2026-05-08T09:05:00Z"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked || model.CurrentSubstitutionGuardState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked timestamp-bound substitution guard, got current=%s guard=%s readiness=%s taxonomy=%s",
				model.CurrentState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
				model.ReplayTaxonomyState,
			)
		}
	})
}

func TestPoint16Val0CachedHelperIsolation(t *testing.T) {
	first := point16Val0ValidFoundationModel()
	first.BlockingReasons = append(first.BlockingReasons, "mutated")
	first.ReviewPrerequisites = append(first.ReviewPrerequisites, "mutated")
	first.Dependency.ReviewPrerequisites = append(first.Dependency.ReviewPrerequisites, "mutated")
	first.ReplayTaxonomy.AllowedStatuses[0] = "mutated"
	first.NoOverclaimBaseline.ObservedTexts = append(first.NoOverclaimBaseline.ObservedTexts, "mutated")
	first.NoOverclaimBaseline.AllowedSafeWording[0] = "mutated"
	first.Dependency.Point15ValE.NoOverclaimFinalCheck.ObservedTexts = append(first.Dependency.Point15ValE.NoOverclaimFinalCheck.ObservedTexts, "mutated")

	second := point16Val0ValidFoundationModel()
	for _, value := range second.BlockingReasons {
		if value == "mutated" {
			t.Fatalf("expected fresh blocking reasons")
		}
	}
	for _, value := range second.ReviewPrerequisites {
		if value == "mutated" {
			t.Fatalf("expected fresh review prerequisites")
		}
	}
	for _, value := range second.Dependency.ReviewPrerequisites {
		if value == "mutated" {
			t.Fatalf("expected fresh dependency review prerequisites")
		}
	}
	for _, value := range second.ReplayTaxonomy.AllowedStatuses {
		if value == "mutated" {
			t.Fatalf("expected fresh replay taxonomy statuses")
		}
	}
	for _, value := range second.NoOverclaimBaseline.ObservedTexts {
		if value == "mutated" {
			t.Fatalf("expected fresh no overclaim observed texts")
		}
	}
	for _, value := range second.NoOverclaimBaseline.AllowedSafeWording {
		if value == "mutated" {
			t.Fatalf("expected fresh no overclaim wording set")
		}
	}
	for _, value := range second.Dependency.Point15ValE.NoOverclaimFinalCheck.ObservedTexts {
		if value == "mutated" {
			t.Fatalf("expected fresh nested point15 vale observed texts")
		}
	}
}

func TestPoint10ThroughPoint16Val0CurrentSweep(t *testing.T) {
	model := ComputePoint16Val0Foundation(Point16Val0FoundationModel())
	if model.CurrentState != Point16Val0StateActive {
		t.Fatalf("expected %s, got %s (blocking=%v review=%v dep=%s ctx=%s binding=%s taxonomy=%s guard=%s readiness=%s nooverclaim=%s)",
			Point16Val0StateActive,
			model.CurrentState,
			model.BlockingReasons,
			model.ReviewPrerequisites,
			model.DependencyState,
			model.HistoricalReplayContextState,
			model.OriginalDecisionBindingState,
			model.ReplayTaxonomyState,
			model.CurrentSubstitutionGuardState,
			model.ReplayReadinessState,
			model.NoOverclaimState,
		)
	}
	payload, err := json.Marshal(model)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	body := string(payload)
	if strings.Contains(body, "point_16_pass") || strings.Contains(body, "point16_pass") {
		t.Fatalf("expected no point16 pass token in point16 val0 sweep")
	}
}

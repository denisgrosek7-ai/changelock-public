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

	t.Run("upstream point15 whitespace retagged manifest provenance blocks at point16 boundary", func(t *testing.T) {
		valE := ComputePoint15ValEFoundation(Point15ValEFoundationModel())
		valE.Dependency.InheritedTenantScope = " " + valE.Dependency.InheritedTenantScope + " "
		valE.PassClosureManifest.EvidenceIdentity = " " + valE.PassClosureManifest.EvidenceIdentity + " "
		valE.PassClosureManifest.EvidenceHash = " " + valE.PassClosureManifest.EvidenceHash + " "
		valE.PassClosureManifest.PolicyVersion = " " + valE.PassClosureManifest.PolicyVersion + " "
		valE.PassClosureManifest.EngineVersion = " " + valE.PassClosureManifest.EngineVersion + " "
		valE.PassClosureManifest.SchemaVersion = " " + valE.PassClosureManifest.SchemaVersion + " "
		valE.PassClosureManifest.GeneratedAt = " " + valE.PassClosureManifest.GeneratedAt + " "
		valE.PassClosureManifest.TenantScope = " " + valE.PassClosureManifest.TenantScope + " "
		valE.PassClosureManifest.Scope = " " + valE.PassClosureManifest.Scope + " "

		model := point16Val0DependencySnapshotFromUpstream(valE)
		if got := EvaluatePoint16Val0DependencyState(model); got != Point16Val0StateBlocked {
			t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
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
		{name: "whitespace point15 pass token retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassToken = " " + model.Point15PassToken + " "
		}},
		{name: "whitespace point15 manifest point id retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestPointID = " " + model.Point15PassManifestPointID + " "
		}},
		{name: "whitespace point15 manifest wave id retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestWaveID = " " + model.Point15PassManifestWaveID + " "
		}},
		{name: "whitespace point15 manifest closure token retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestClosureToken = " " + model.Point15PassManifestClosureToken + " "
		}},
		{name: "whitespace point15 manifest scope retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestScope = " " + model.Point15PassManifestScope + " "
		}},
		{name: "nested pass manifest evidence identity drift blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15ValE.PassClosureManifest.EvidenceIdentity = "evidence_id=evidence_point16_val0_alt evidence_hash=sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa policy=policy.alt.v1 engine=engine.alt.v1 schema=schema.alt.v1 tenant=tenant_alt"
		}},
		{name: "nested pass manifest evidence hash drift blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15ValE.PassClosureManifest.EvidenceHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		}},
		{name: "joint upstream evidence hash whitespace retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestEvidenceHash += " "
			model.Point15ValE.PassClosureManifest.EvidenceHash += " "
		}},
		{name: "nested pass manifest policy version drift blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15ValE.PassClosureManifest.PolicyVersion = "policy.alt.v1"
		}},
		{name: "joint upstream policy version whitespace retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestPolicyVersion += " "
			model.Point15ValE.PassClosureManifest.PolicyVersion += " "
		}},
		{name: "nested pass manifest engine version drift blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15ValE.PassClosureManifest.EngineVersion = "engine.alt.v1"
		}},
		{name: "joint upstream engine version whitespace retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestEngineVersion += " "
			model.Point15ValE.PassClosureManifest.EngineVersion += " "
		}},
		{name: "nested pass manifest schema version drift blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15ValE.PassClosureManifest.SchemaVersion = "schema.alt.v1"
		}},
		{name: "joint upstream schema version whitespace retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestSchemaVersion += " "
			model.Point15ValE.PassClosureManifest.SchemaVersion += " "
		}},
		{name: "nested pass manifest generated at drift blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15ValE.PassClosureManifest.GeneratedAt = "2026-05-08T10:00:00Z"
		}},
		{name: "joint upstream generated at whitespace retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestGeneratedAt += " "
			model.Point15ValE.PassClosureManifest.GeneratedAt += " "
		}},
		{name: "joint upstream generated at retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestGeneratedAt = "2026-05-08T10:00:00Z"
			model.Point15ValE.PassClosureManifest.GeneratedAt = model.Point15PassManifestGeneratedAt
		}},
		{name: "nested pass manifest scope drift blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15ValE.PassClosureManifest.Scope = "retagged_final_continuous_verification_closure_gate"
		}},
		{name: "joint upstream pass manifest scope whitespace retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestScope += " "
			model.Point15ValE.PassClosureManifest.Scope += " "
		}},
		{name: "joint upstream pass manifest scope retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestScope = "retagged_final_continuous_verification_closure_gate"
			model.Point15ValE.PassClosureManifest.Scope = model.Point15PassManifestScope
		}},
		{name: "nested pass manifest tenant scope drift blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15ValE.PassClosureManifest.TenantScope = "tenant_alt"
		}},
		{name: "joint upstream pass manifest tenant scope whitespace retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestTenantScope += " "
			model.Point15ValE.PassClosureManifest.TenantScope += " "
		}},
		{name: "joint upstream tenant scope whitespace retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.InheritedTenantScope += " "
			model.Point15ValE.Dependency.InheritedTenantScope += " "
		}},
		{name: "joint upstream manifest provenance retag blocks", mutate: func(model *Point16Val0DependencySnapshot) {
			model.Point15PassManifestEvidenceID = "evidence_id=evidence_point16_val0_alt evidence_hash=sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa policy=policy.alt.v1 engine=engine.alt.v1 schema=schema.alt.v1 tenant=tenant_alt"
			model.Point15PassManifestEvidenceHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
			model.Point15PassManifestPolicyVersion = "policy.alt.v1"
			model.Point15PassManifestEngineVersion = "engine.alt.v1"
			model.Point15PassManifestScope = "retagged_final_continuous_verification_closure_gate"
			model.InheritedTenantScope = "tenant_alt"
			model.Point15ValE.PassClosureManifest.EvidenceIdentity = model.Point15PassManifestEvidenceID
			model.Point15ValE.PassClosureManifest.EvidenceHash = model.Point15PassManifestEvidenceHash
			model.Point15ValE.PassClosureManifest.PolicyVersion = model.Point15PassManifestPolicyVersion
			model.Point15ValE.PassClosureManifest.EngineVersion = model.Point15PassManifestEngineVersion
			model.Point15ValE.PassClosureManifest.Scope = model.Point15PassManifestScope
			model.Point15ValE.Dependency.InheritedTenantScope = model.InheritedTenantScope
		}},
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
		{name: "bare evidence ref without structured identity blocks", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.OriginalEvidenceID = "evidence_point16_val0_original"
		}, expected: Point16Val0StateBlocked},
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
		{name: "replay time source whitespace retag blocks", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.ReplayTimeSource = " " + model.ReplayTimeSource + " "
		}, expected: Point16Val0StateBlocked},
		{name: "approved customer original decision time source blocks", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.OriginalDecisionTimeSource = point14Val0TimeSourceApprovedCustomerTime
		}, expected: Point16Val0StateBlocked},
		{name: "original decision time source whitespace retag blocks", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.OriginalDecisionTimeSource = " " + model.OriginalDecisionTimeSource + " "
		}, expected: Point16Val0StateBlocked},
		{name: "approved customer original evaluated time source blocks", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.OriginalEvaluatedTimeSource = point14Val0TimeSourceApprovedCustomerTime
		}, expected: Point16Val0StateBlocked},
		{name: "original evaluated time source whitespace retag blocks", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.OriginalEvaluatedTimeSource = " " + model.OriginalEvaluatedTimeSource + " "
		}, expected: Point16Val0StateBlocked},
		{name: "timestamp backdating blocks", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.ReplayAt = "2026-05-08T08:59:00Z"
		}, expected: Point16Val0StateBlocked},
		{name: "context id exact mismatch blocks", mutate: func(model *Point16Val0HistoricalReplayContext) {
			model.ContextID = point16Val0ContextID + " "
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

func TestPoint16Val0OriginalDecisionBindingState(t *testing.T) {
	t.Run("clean original decision binding is active", func(t *testing.T) {
		model := point16Val0ValidOriginalDecisionBindingModel()
		if got := EvaluatePoint16Val0OriginalDecisionBindingState(model); got != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s", Point16Val0StateActive, got)
		}
	})

	tests := []struct {
		name   string
		mutate func(*Point16Val0OriginalDecisionBinding)
	}{
		{name: "binding id exact mismatch blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.BindingID = point16Val0BindingID + " "
		}},
		{name: "historical replay context ref exact mismatch blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.HistoricalReplayContextRef = point16Val0ContextID + " "
		}},
		{name: "original decision id exact mismatch blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.OriginalDecisionID = "decision_point16_val0_original_other"
		}},
		{name: "original decision hash exact mismatch blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.OriginalDecisionHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		}},
		{name: "bare evidence ref without structured identity blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.OriginalEvidenceID = "evidence_point16_val0_original"
		}},
		{name: "whitespace structured evidence identity blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.OriginalEvidenceID = " " + model.OriginalEvidenceID + " "
		}},
		{name: "fabricated composite evidence identity blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.OriginalEvidenceID = "junk evidence_id=x evidence_hash=y policy=z engine=w tenant=t junk"
		}},
		{name: "current evidence hash whitespace retag blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.CurrentEvidenceHash = " " + model.OriginalEvidenceHash + " "
		}},
		{name: "current policy version exact mismatch blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.CurrentPolicyVersion = model.OriginalPolicyVersion + ".other"
		}},
		{name: "current policy version whitespace retag blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.CurrentPolicyVersion = " " + model.OriginalPolicyVersion + " "
		}},
		{name: "current engine version exact mismatch blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.CurrentEngineVersion = model.OriginalEngineVersion + ".other"
		}},
		{name: "current engine version whitespace retag blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.CurrentEngineVersion = " " + model.OriginalEngineVersion + " "
		}},
		{name: "current tenant scope whitespace retag blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.CurrentTenantScope = " " + model.OriginalTenantScope + " "
		}},
		{name: "current artifact scope whitespace retag blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.CurrentArtifactScope = " " + model.OriginalArtifactScope + " "
		}},
		{name: "current claim scope whitespace retag blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.CurrentClaimScope = " " + model.OriginalClaimScope + " "
		}},
		{name: "current governance scope whitespace retag blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.CurrentGovernanceScope = " " + model.OriginalGovernanceScope + " "
		}},
		{name: "original evaluated at drift blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.OriginalEvaluatedAt = "2026-05-08T09:06:00Z"
		}},
		{name: "original evaluated at whitespace retag blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.OriginalEvaluatedAt = " " + model.OriginalEvaluatedAt + " "
		}},
		{name: "original evaluated at substituted from replay time blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.OriginalEvaluatedAt = point16Val0ReplayAt
		}},
		{name: "current context comparison disabled blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.CurrentContextComparisonOnly = false
		}},
		{name: "lineage ref exact mismatch blocks", mutate: func(model *Point16Val0OriginalDecisionBinding) {
			model.LineageRef = "lineage_point16_val0_replay_other"
		}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			model := point16Val0ValidOriginalDecisionBindingModel()
			tc.mutate(&model)
			if got := EvaluatePoint16Val0OriginalDecisionBindingState(model); got != Point16Val0StateBlocked {
				t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
			}
		})
	}

	t.Run("missing original evaluated at is incomplete", func(t *testing.T) {
		model := point16Val0ValidOriginalDecisionBindingModel()
		model.OriginalEvaluatedAt = ""
		if got := EvaluatePoint16Val0OriginalDecisionBindingState(model); got != Point16Val0StateIncomplete {
			t.Fatalf("expected %s, got %s", Point16Val0StateIncomplete, got)
		}
	})
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
		{name: "replay status whitespace retag blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.ReplayStatus = " " + model.ReplayStatus + " "
		}, expected: Point16Val0StateBlocked},
		{name: "allowed statuses whitespace retag blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.AllowedStatuses[0] = " " + model.AllowedStatuses[0] + " "
		}, expected: Point16Val0StateBlocked},
		{name: "taxonomy id exact mismatch blocks", mutate: func(model *Point16Val0ReplayTaxonomy) {
			model.TaxonomyID = point16Val0TaxonomyID + " "
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
		{name: "guard id exact mismatch fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.GuardID = point16Val0GuardID + " " }},
		{name: "current policy version substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) {
			model.CurrentPolicyVersion = model.OriginalPolicyVersion + ".other"
		}},
		{name: "current policy version whitespace retag fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) {
			model.CurrentPolicyVersion = " " + model.OriginalPolicyVersion + " "
		}},
		{name: "current policy substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentPolicyAuthoritative = true }},
		{name: "current engine version substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) {
			model.CurrentEngineVersion = model.OriginalEngineVersion + ".other"
		}},
		{name: "current engine version whitespace retag fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) {
			model.CurrentEngineVersion = " " + model.OriginalEngineVersion + " "
		}},
		{name: "current engine substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentEngineAuthoritative = true }},
		{name: "current evidence substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentEvidenceAuthoritative = true }},
		{name: "current evidence hash whitespace retag fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) {
			model.CurrentEvidenceHash = " " + model.OriginalEvidenceHash + " "
		}},
		{name: "current timestamp substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentTimestampAuthoritative = true }},
		{name: "whitespace original decision timestamp fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) {
			model.OriginalDecisionAt = " " + model.OriginalDecisionAt + " "
		}},
		{name: "current tenant substitution attempt fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) { model.CurrentTenantAuthoritative = true }},
		{name: "current tenant scope whitespace retag fails closed", mutate: func(model *Point16Val0CurrentSubstitutionGuard) {
			model.CurrentTenantScope = " " + model.OriginalTenantScope + " "
		}},
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
		{name: "evaluation id exact mismatch blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.EvaluationID = point16Val0EvaluationID + " "
		}},
		{name: "whitespace retagged component state blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.ReplayTaxonomyState = " " + Point16Val0StateActive + " "
		}},
		{name: "original context exact bound false blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.OriginalContextExactBound = false
		}},
		{name: "timestamp unsafe blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.TimestampSafe = false
		}},
		{name: "replay unsupported blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.ReplaySupported = false
		}},
		{name: "premature final token detection blocks", mutate: func(model *Point16Val0ReplayReadinessEvaluation) {
			model.NoPrematureFinalTokenDetected = false
		}},
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

	t.Run("replay disclaimer whitespace retag blocks", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.ReplayDisclaimer = " " + model.ReplayDisclaimer + " "
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateBlocked {
			t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
		}
	})

	t.Run("allowed safe wording whitespace retag blocks", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.AllowedSafeWording[0] = " " + model.AllowedSafeWording[0] + " "
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateBlocked {
			t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
		}
	})

	t.Run("blocked wording whitespace retag blocks", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.BlockedWording[0] = " " + model.BlockedWording[0] + " "
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateBlocked {
			t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
		}
	})

	obfuscatedForbiddenTexts := []string{
		"certified-secure",
		"certified secure",
		" deployment approved ",
		"deployment approved",
		"deployment-approved",
		"ᴅeployment approved",
		"production-ᴀpproved",
		"production approᴠed",
		"certified\nsecure",
		"certiﬁed secure",
		"ᴄertified secure",
		"cértified secure",
		"cɛrtified secure",
		"certifıed secure",
		"certιfied secure",
		"certiſied secure",
		"certifieԁ secure",
		"certifiеd secure",
		"certifieɗ secure",
		"certiƒied secure",
		"certiꜰied secure",
		"certifiᴇd secure",
		"public\nbadge",
		"public-badge",
		"publicbadge",
		"public ʙadge",
		"public bɑdge",
		"pubɩic badge",
		"gυaranteed secure",
		"guaʀanteed secure",
		"guaɹanteed secure",
		"gᴜaranteed secure",
		"guaranteed ꜱecure",
		"regulator-approved",
		"ƈompliance guaranteed",
		"ȼompliance guaranteed",
		"regulátor-approved",
		"regulatоr-approved",
		"compliance\nguaranteed",
		"production-approved",
		"prɔduction-approved",
		"prɔduction approved",
		"pr0duction appr0ved",
		"deᴘloyment approved",
		"legal/proof",
		"ʟegal proof",
		"le9al proof",
		"financial_guarantee",
		"fiɴancial guarantee",
		"public b8dge",
		"gl0bal truth",
		"9lobal truth",
		"global ᴛruth",
		"ɢlobal truth",
		"0fficial authority",
		"official autʜority",
		"ᴏfficial authority",
		"օfficial authority",
		"cᴇrtified secure",
		"CeRtIfIeD---SeCuRe",
	}
	for _, observed := range obfuscatedForbiddenTexts {
		observed := observed
		t.Run("obfuscated forbidden wording blocks "+strings.ReplaceAll(observed, "\n", "\\n"), func(t *testing.T) {
			model := point16Val0ValidNoOverclaimBaselineModel()
			model.ObservedTexts = []string{observed}
			if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateBlocked {
				t.Fatalf("expected %s for observed text %q, got %s", Point16Val0StateBlocked, observed, got)
			}
		})
	}

	t.Run("benign text remains active", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.ObservedTexts = []string{"historical replay comparison remains bounded to the original decision context and fails closed when unsupported"}
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s", Point16Val0StateActive, got)
		}
	})

	t.Run("split forbidden wording across observed texts blocks", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.ObservedTexts = []string{"certified", "secure"}
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateBlocked {
			t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
		}
	})

	t.Run("unclassified forbidden internal diagnostics block", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.InternalDiagnosticTexts = []string{"legal proof"}
		model.InternalDiagnosticsClassifiedBlocked = false
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateBlocked {
			t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
		}
	})

	t.Run("classified forbidden internal diagnostics remain bounded", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.InternalDiagnosticTexts = []string{"legal proof"}
		model.InternalDiagnosticsClassifiedBlocked = true
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s", Point16Val0StateActive, got)
		}
	})

	t.Run("split forbidden wording across observed and internal diagnostics blocks", func(t *testing.T) {
		tests := []struct {
			name        string
			observed    []string
			diagnostics []string
		}{
			{
				name:        "deployment approved across buckets",
				observed:    []string{"deployment"},
				diagnostics: []string{"approved"},
			},
			{
				name:        "production approved across buckets",
				observed:    []string{"production"},
				diagnostics: []string{"approved"},
			},
			{
				name:        "public badge across buckets",
				observed:    []string{"public"},
				diagnostics: []string{"badge"},
			},
			{
				name:        "legal proof across buckets",
				observed:    []string{"legal"},
				diagnostics: []string{"proof"},
			},
			{
				name:        "guaranteed secure across buckets",
				observed:    []string{"guaranteed"},
				diagnostics: []string{"secure"},
			},
		}
		for _, tc := range tests {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				model := point16Val0ValidNoOverclaimBaselineModel()
				model.ObservedTexts = tc.observed
				model.InternalDiagnosticTexts = tc.diagnostics
				model.InternalDiagnosticsClassifiedBlocked = false
				if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateBlocked {
					t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
				}
			})
		}
	})

	t.Run("cross bucket forbidden wording still blocks when diagnostics already contain classified forbidden phrase", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.ObservedTexts = []string{"guaranteed"}
		model.InternalDiagnosticTexts = []string{"secure", "legal proof"}
		model.InternalDiagnosticsClassifiedBlocked = true
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateBlocked {
			t.Fatalf("expected %s, got %s", Point16Val0StateBlocked, got)
		}
	})

	t.Run("benign split text across buckets remains active", func(t *testing.T) {
		model := point16Val0ValidNoOverclaimBaselineModel()
		model.ObservedTexts = []string{"historical replay"}
		model.InternalDiagnosticTexts = []string{"remains bounded"}
		model.InternalDiagnosticsClassifiedBlocked = false
		if got := EvaluatePoint16Val0NoOverclaimBaselineState(model); got != Point16Val0StateActive {
			t.Fatalf("expected %s, got %s", Point16Val0StateActive, got)
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

	t.Run("approved customer timestamp provenance cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.HistoricalReplayContext.OriginalDecisionTimeSource = point14Val0TimeSourceApprovedCustomerTime
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked customer-authored timestamp provenance, got current=%s context=%s binding=%s taxonomy=%s guard=%s readiness=%s",
				model.CurrentState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("whitespace original decision time source provenance cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.HistoricalReplayContext.OriginalDecisionTimeSource = " " + model.HistoricalReplayContext.OriginalDecisionTimeSource + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace decision-time provenance drift, got current=%s context=%s binding=%s taxonomy=%s guard=%s readiness=%s",
				model.CurrentState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("approved customer evaluated timestamp provenance cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.HistoricalReplayContext.OriginalEvaluatedTimeSource = point14Val0TimeSourceApprovedCustomerTime
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked evaluated-time provenance drift, got current=%s context=%s binding=%s taxonomy=%s guard=%s readiness=%s",
				model.CurrentState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("whitespace original evaluated time source provenance cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.HistoricalReplayContext.OriginalEvaluatedTimeSource = " " + model.HistoricalReplayContext.OriginalEvaluatedTimeSource + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace evaluated-time provenance drift, got current=%s context=%s binding=%s taxonomy=%s guard=%s readiness=%s",
				model.CurrentState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("nested dependency manifest evidence provenance drift cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.Dependency.Point15ValE.PassClosureManifest.EvidenceIdentity = "evidence_id=evidence_point16_val0_alt evidence_hash=sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa policy=policy.alt.v1 engine=engine.alt.v1 schema=schema.alt.v1 tenant=tenant_alt"
		model.Dependency.Point15ValE.PassClosureManifest.EvidenceHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		model.Dependency.Point15ValE.PassClosureManifest.PolicyVersion = "policy.alt.v1"
		model.Dependency.Point15ValE.PassClosureManifest.EngineVersion = "engine.alt.v1"
		model.Dependency.Point15ValE.PassClosureManifest.SchemaVersion = "schema.alt.v1"
		model.Dependency.Point15ValE.PassClosureManifest.TenantScope = "tenant_alt"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked || model.DependencyState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked dependency provenance drift, got current=%s dependency=%s", model.CurrentState, model.DependencyState)
		}
	})

	t.Run("joint upstream manifest provenance retag cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.Dependency.Point15PassManifestEvidenceID = "evidence_id=evidence_point16_val0_alt evidence_hash=sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa policy=policy.alt.v1 engine=engine.alt.v1 schema=schema.alt.v1 tenant=tenant_alt"
		model.Dependency.Point15PassManifestEvidenceHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		model.Dependency.Point15PassManifestPolicyVersion = "policy.alt.v1"
		model.Dependency.Point15PassManifestEngineVersion = "engine.alt.v1"
		model.Dependency.Point15PassManifestSchemaVersion = "schema.alt.v1"
		model.Dependency.Point15PassManifestTenantScope = "tenant_alt"
		model.Dependency.InheritedTenantScope = "tenant_alt"
		model.Dependency.Point15ValE.PassClosureManifest.EvidenceIdentity = model.Dependency.Point15PassManifestEvidenceID
		model.Dependency.Point15ValE.PassClosureManifest.EvidenceHash = model.Dependency.Point15PassManifestEvidenceHash
		model.Dependency.Point15ValE.PassClosureManifest.PolicyVersion = model.Dependency.Point15PassManifestPolicyVersion
		model.Dependency.Point15ValE.PassClosureManifest.EngineVersion = model.Dependency.Point15PassManifestEngineVersion
		model.Dependency.Point15ValE.PassClosureManifest.SchemaVersion = model.Dependency.Point15PassManifestSchemaVersion
		model.Dependency.Point15ValE.PassClosureManifest.TenantScope = model.Dependency.Point15PassManifestTenantScope
		model.Dependency.Point15ValE.Dependency.InheritedTenantScope = model.Dependency.InheritedTenantScope
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.DependencyState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked jointly retagged upstream manifest provenance, got current=%s dependency=%s ctx=%s binding=%s guard=%s readiness=%s",
				model.CurrentState,
				model.DependencyState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("joint upstream evidence hash whitespace retag cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.Dependency.Point15PassManifestEvidenceHash += " "
		model.Dependency.Point15ValE.PassClosureManifest.EvidenceHash += " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.DependencyState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace-retagged upstream evidence hash, got current=%s dependency=%s ctx=%s binding=%s guard=%s readiness=%s",
				model.CurrentState,
				model.DependencyState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("joint upstream tenant scope whitespace retag cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.Dependency.InheritedTenantScope += " "
		model.Dependency.Point15ValE.Dependency.InheritedTenantScope += " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.DependencyState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace-retagged upstream tenant scope, got current=%s dependency=%s ctx=%s binding=%s guard=%s readiness=%s",
				model.CurrentState,
				model.DependencyState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("replay status whitespace retag cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.ReplayTaxonomy.ReplayStatus = " " + model.ReplayTaxonomy.ReplayStatus + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace-retagged replay status, got current=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("replay disclaimer whitespace retag cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.NoOverclaimBaseline.ReplayDisclaimer = " " + model.NoOverclaimBaseline.ReplayDisclaimer + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.NoOverclaimState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace-retagged replay disclaimer, got current=%s nooverclaim=%s readiness=%s",
				model.CurrentState,
				model.NoOverclaimState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("current policy version substitution attempt fails closed in full foundation", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.CurrentPolicyVersion = model.OriginalDecisionBinding.OriginalPolicyVersion + ".other"
		model.CurrentSubstitutionGuard.CurrentPolicyVersion = model.CurrentSubstitutionGuard.OriginalPolicyVersion + ".other"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked version-bound current policy substitution, got current=%s binding=%s guard=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("whitespace current evidence hash substitution attempt fails closed in full foundation", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.CurrentEvidenceHash = " " + model.OriginalDecisionBinding.OriginalEvidenceHash + " "
		model.CurrentSubstitutionGuard.CurrentEvidenceHash = " " + model.CurrentSubstitutionGuard.OriginalEvidenceHash + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace evidence-hash substitution, got current=%s binding=%s guard=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("whitespace current policy version substitution attempt fails closed in full foundation", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.CurrentPolicyVersion = " " + model.OriginalDecisionBinding.OriginalPolicyVersion + " "
		model.CurrentSubstitutionGuard.CurrentPolicyVersion = " " + model.CurrentSubstitutionGuard.OriginalPolicyVersion + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace version-bound current policy substitution, got current=%s binding=%s guard=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("current engine version substitution attempt fails closed in full foundation", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.CurrentEngineVersion = model.OriginalDecisionBinding.OriginalEngineVersion + ".other"
		model.CurrentSubstitutionGuard.CurrentEngineVersion = model.CurrentSubstitutionGuard.OriginalEngineVersion + ".other"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked version-bound current engine substitution, got current=%s binding=%s guard=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("whitespace current engine version substitution attempt fails closed in full foundation", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.CurrentEngineVersion = " " + model.OriginalDecisionBinding.OriginalEngineVersion + " "
		model.CurrentSubstitutionGuard.CurrentEngineVersion = " " + model.CurrentSubstitutionGuard.OriginalEngineVersion + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace version-bound current engine substitution, got current=%s binding=%s guard=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("whitespace current tenant scope substitution attempt fails closed in full foundation", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.CurrentTenantScope = " " + model.OriginalDecisionBinding.OriginalTenantScope + " "
		model.CurrentSubstitutionGuard.CurrentTenantScope = " " + model.CurrentSubstitutionGuard.OriginalTenantScope + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace tenant-scope substitution, got current=%s binding=%s guard=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("whitespace current artifact scope substitution attempt fails closed in full foundation", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.CurrentArtifactScope = " " + model.OriginalDecisionBinding.OriginalArtifactScope + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace artifact-scope substitution, got current=%s binding=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("whitespace tampered substitution guard original decision timestamp fails closed", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.CurrentSubstitutionGuard.OriginalDecisionAt = " " + model.CurrentSubstitutionGuard.OriginalDecisionAt + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked || model.CurrentSubstitutionGuardState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace timestamp-bound substitution guard, got current=%s guard=%s readiness=%s taxonomy=%s",
				model.CurrentState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
				model.ReplayTaxonomyState,
			)
		}
	})

	t.Run("jointly retagged timestamp provenance cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.HistoricalReplayContext.OriginalDecisionAt = "2026-05-08T10:00:00Z"
		model.HistoricalReplayContext.OriginalEvaluatedAt = "2026-05-08T10:05:00Z"
		model.HistoricalReplayContext.ReplayAt = "2026-05-08T10:10:00Z"
		model.OriginalDecisionBinding.OriginalDecisionAt = "2026-05-08T10:00:00Z"
		model.OriginalDecisionBinding.OriginalEvaluatedAt = "2026-05-08T10:05:00Z"
		model.CurrentSubstitutionGuard.OriginalDecisionAt = "2026-05-08T10:00:00Z"
		model.CurrentSubstitutionGuard.ReplayAt = "2026-05-08T10:10:00Z"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked jointly retagged timestamp provenance, got current=%s context=%s binding=%s guard=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("mismatched HistoricalReplayContextRef fails closed", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.HistoricalReplayContextRef = "historical_replay_context_point16_val0_other"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked provenance retagging, got current=%s binding=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("whitespace mutated HistoricalReplayContextRef fails closed", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.HistoricalReplayContextRef = model.OriginalDecisionBinding.HistoricalReplayContextRef + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace provenance retagging, got current=%s binding=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("mismatched OriginalDecisionID fails closed", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.OriginalDecisionID = "decision_point16_val0_original_other"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked provenance retagging, got current=%s binding=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("whitespace mutated OriginalDecisionID fails closed", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.OriginalDecisionID = model.OriginalDecisionBinding.OriginalDecisionID + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked whitespace decision retagging, got current=%s binding=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("mismatched OriginalDecisionHash fails closed", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.OriginalDecisionHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked provenance retagging, got current=%s binding=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("mismatched LineageRef fails closed", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.LineageRef = "lineage_point16_val0_replay_other"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked provenance retagging, got current=%s binding=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("retagging payload to another decision lineage cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.OriginalDecisionBinding.OriginalDecisionID = "decision_point16_val0_original_other"
		model.OriginalDecisionBinding.OriginalDecisionHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		model.OriginalDecisionBinding.LineageRef = "lineage_point16_val0_replay_other"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked retagged payload, got current=%s binding=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("jointly retagged provenance cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.HistoricalReplayContext.ContextID = "historical_replay_context_point16_val0_other"
		model.HistoricalReplayContext.OriginalDecisionID = "decision_point16_val0_original_other"
		model.HistoricalReplayContext.OriginalDecisionHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		model.HistoricalReplayContext.LineageRef = "lineage_point16_val0_replay_other"
		model.OriginalDecisionBinding.HistoricalReplayContextRef = "historical_replay_context_point16_val0_other"
		model.OriginalDecisionBinding.OriginalDecisionID = "decision_point16_val0_original_other"
		model.OriginalDecisionBinding.OriginalDecisionHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		model.OriginalDecisionBinding.LineageRef = "lineage_point16_val0_replay_other"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked jointly retagged provenance, got current=%s context=%s binding=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("jointly retagged policy engine and scope provenance cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.HistoricalReplayContext.OriginalPolicyID = "policy_historical_replay_other"
		model.HistoricalReplayContext.OriginalPolicyHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		model.HistoricalReplayContext.OriginalEngineID = "engine_historical_replay_other"
		model.HistoricalReplayContext.OriginalEngineHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		model.HistoricalReplayContext.OriginalArtifactScope = "artifact_scope_historical_replay_other"
		model.HistoricalReplayContext.OriginalClaimScope = "claim_scope_historical_replay_other"
		model.HistoricalReplayContext.OriginalGovernanceScope = "governance_scope_historical_replay_other"
		model.OriginalDecisionBinding.OriginalPolicyID = "policy_historical_replay_other"
		model.OriginalDecisionBinding.OriginalPolicyHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		model.OriginalDecisionBinding.OriginalEngineID = "engine_historical_replay_other"
		model.OriginalDecisionBinding.OriginalEngineHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		model.OriginalDecisionBinding.OriginalArtifactScope = "artifact_scope_historical_replay_other"
		model.OriginalDecisionBinding.CurrentArtifactScope = "artifact_scope_historical_replay_other"
		model.OriginalDecisionBinding.OriginalClaimScope = "claim_scope_historical_replay_other"
		model.OriginalDecisionBinding.CurrentClaimScope = "claim_scope_historical_replay_other"
		model.OriginalDecisionBinding.OriginalGovernanceScope = "governance_scope_historical_replay_other"
		model.OriginalDecisionBinding.CurrentGovernanceScope = "governance_scope_historical_replay_other"
		model.CurrentSubstitutionGuard.OriginalPolicyHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		model.CurrentSubstitutionGuard.CurrentPolicyHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		model.CurrentSubstitutionGuard.OriginalEngineHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		model.CurrentSubstitutionGuard.CurrentEngineHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		model.CurrentSubstitutionGuard.OriginalClaimScope = "claim_scope_historical_replay_other"
		model.CurrentSubstitutionGuard.CurrentClaimScope = "claim_scope_historical_replay_other"
		model.CurrentSubstitutionGuard.OriginalGovernanceScope = "governance_scope_historical_replay_other"
		model.CurrentSubstitutionGuard.CurrentGovernanceScope = "governance_scope_historical_replay_other"
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked jointly retagged policy/engine/scope provenance, got current=%s context=%s binding=%s guard=%s taxonomy=%s readiness=%s",
				model.CurrentState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.CurrentSubstitutionGuardState,
				model.ReplayTaxonomyState,
				model.ReplayReadinessState,
			)
		}
	})

	t.Run("mutated internal ids cannot compute active", func(t *testing.T) {
		model := point16Val0ValidFoundationModel()
		model.HistoricalReplayContext.ContextID = point16Val0ContextID + " "
		model.OriginalDecisionBinding.BindingID = point16Val0BindingID + " "
		model.ReplayTaxonomy.TaxonomyID = point16Val0TaxonomyID + " "
		model.CurrentSubstitutionGuard.GuardID = point16Val0GuardID + " "
		model.ReplayReadinessEvaluation.EvaluationID = point16Val0EvaluationID + " "
		model = ComputePoint16Val0Foundation(model)
		if model.CurrentState != Point16Val0StateBlocked ||
			model.HistoricalReplayContextState != Point16Val0StateBlocked ||
			model.OriginalDecisionBindingState != Point16Val0StateBlocked ||
			model.ReplayTaxonomyState != Point16Val0StateBlocked ||
			model.CurrentSubstitutionGuardState != Point16Val0StateBlocked ||
			model.ReplayReadinessState != Point16Val0StateBlocked {
			t.Fatalf("expected blocked mutated internal ids, got current=%s context=%s binding=%s taxonomy=%s guard=%s readiness=%s",
				model.CurrentState,
				model.HistoricalReplayContextState,
				model.OriginalDecisionBindingState,
				model.ReplayTaxonomyState,
				model.CurrentSubstitutionGuardState,
				model.ReplayReadinessState,
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

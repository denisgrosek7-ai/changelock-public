package formal

import (
	"encoding/json"
	"sync"
	"testing"
)

var (
	point12ValEActiveFoundationBaselineJSON []byte
	point12ValEActiveFoundationBaselineOnce sync.Once
	point12ValERawFoundationModelJSON       []byte
	point12ValERawFoundationModelOnce       sync.Once
)

func mustMarshalPoint12ValEFoundation(model Point12ValEFoundation) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint12ValEFoundation(payload []byte) Point12ValEFoundation {
	var clone Point12ValEFoundation
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

func activePoint12ValDFoundationFromValC(valC Point12ValCFoundation) Point12ValDFoundation {
	model := Point12ValDFoundationModel()
	model.Dependency = SnapshotPoint12ValDDependencyFromComputedValC(valC, point12ValDDependencyReviewContextModel())
	return ComputePoint12ValDFoundation(model)
}

func uncachedActivePoint12ValEFoundation() Point12ValEFoundation {
	val0 := activePoint12Val0Foundation()
	valA := activePoint12ValAFoundationFromVal0(val0)
	valB := activePoint12ValBFoundationFromValA(valA)
	valC := activePoint12ValCFoundationFromValB(valB)
	valD := activePoint12ValDFoundationFromValC(valC)
	return ComputePoint12ValEFoundation(point12ValEFoundationModelFromUpstream(val0, valA, valB, valC, valD))
}

func rawPoint12ValEFoundationModel() Point12ValEFoundation {
	point12ValERawFoundationModelOnce.Do(func() {
		point12ValERawFoundationModelJSON = mustMarshalPoint12ValEFoundation(Point12ValEFoundationModel())
	})
	return clonePoint12ValEFoundation(point12ValERawFoundationModelJSON)
}

func activePoint12ValEFoundation() Point12ValEFoundation {
	point12ValEActiveFoundationBaselineOnce.Do(func() {
		point12ValEActiveFoundationBaselineJSON = mustMarshalPoint12ValEFoundation(uncachedActivePoint12ValEFoundation())
	})
	return clonePoint12ValEFoundation(point12ValEActiveFoundationBaselineJSON)
}

func assertPoint12ValENoPass(t *testing.T, model Point12ValEFoundation) {
	t.Helper()
	if model.Point12PassAllowed {
		t.Fatalf("expected point_12_pass to remain disallowed, got %#v", model)
	}
	if model.Point12PassToken != "" {
		t.Fatalf("expected empty point_12_pass token, got %#v", model)
	}
	if model.PassClosureManifest.Point12PassAllowed {
		t.Fatalf("expected nested pass closure manifest to disallow point_12_pass, got %#v", model.PassClosureManifest)
	}
	if model.PassClosureManifest.Point12PassToken != "" {
		t.Fatalf("expected nested pass closure manifest token to be cleared, got %#v", model.PassClosureManifest)
	}
	if model.CurrentState == Point12ValEStatePassConfirmed {
		t.Fatalf("expected state other than pass_confirmed, got %#v", model)
	}
}

func assertPoint12ValEReason(t *testing.T, reasons []string, want string) {
	t.Helper()
	for _, reason := range reasons {
		if reason == want {
			return
		}
	}
	t.Fatalf("expected exact reason %q, got %#v", want, reasons)
}

func TestPoint12ValEFoundationFixtureIsolation(t *testing.T) {
	t.Run("connected active path still computes", func(t *testing.T) {
		model := uncachedActivePoint12ValEFoundation()
		if model.CurrentState != Point12ValEStatePassConfirmed {
			t.Fatalf("expected connected active path to compute pass-confirmed baseline, got %#v", model)
		}
		if !model.Point12PassAllowed || model.Point12PassToken != point12ValEPoint12PassToken {
			t.Fatalf("expected connected active path to emit point_12_pass on final happy path, got %#v", model)
		}
	})

	t.Run("direct foundation model remains pass confirmed", func(t *testing.T) {
		model := Point12ValEFoundationModel()
		if model.CurrentState != Point12ValEStatePassConfirmed {
			t.Fatalf("expected direct foundation fixture to stay pass confirmed, got %#v", model)
		}
		if !model.Point12PassAllowed || model.Point12PassToken != point12ValEPoint12PassToken {
			t.Fatalf("expected direct foundation fixture to keep final point_12_pass token, got %#v", model)
		}
	})

	t.Run("cached active baseline remains pass confirmed", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if model.CurrentState != Point12ValEStatePassConfirmed {
			t.Fatalf("expected cached baseline to stay pass confirmed, got %#v", model)
		}
		if !model.Point12PassAllowed || model.Point12PassToken != point12ValEPoint12PassToken {
			t.Fatalf("expected cached baseline to keep final point_12_pass token, got %#v", model)
		}
	})

	t.Run("cached fixture mutation does not contaminate next clone", func(t *testing.T) {
		mutated := activePoint12ValEFoundation()
		mutated.Dependency.ValB.ReplayResult.UnsupportedVersion = true
		mutated.PassClosureManifest.ClosureManifestID = ""
		mutated.EvidenceQualityMap.MissingRefs = []string{"evidence:point12-vale-mutation-001"}

		fresh := activePoint12ValEFoundation()
		if fresh.Dependency.ValB.ReplayResult.UnsupportedVersion {
			t.Fatalf("expected cached baseline clone to keep unsupported version false, got %#v", fresh.Dependency.ValB.ReplayResult)
		}
		if fresh.PassClosureManifest.ClosureManifestID == "" {
			t.Fatalf("expected cached baseline clone to keep closure manifest id, got %#v", fresh.PassClosureManifest)
		}
		if len(fresh.EvidenceQualityMap.MissingRefs) != 0 {
			t.Fatalf("expected cached baseline clone to keep missing refs empty, got %#v", fresh.EvidenceQualityMap)
		}
	})
}

func TestPoint12ValEDependencyState(t *testing.T) {
	t.Run("valid computed vald output active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got := EvaluatePoint12ValEDependencyState(model.Dependency); got != Point12ValEStateActive {
			t.Fatalf("expected active dependency state, got %#v", model.Dependency)
		}
	})

	cases := []struct {
		name   string
		mutate func(*Point12ValEDependencySnapshot)
		want   string
	}{
		{name: "missing computed snapshot blocks", mutate: func(model *Point12ValEDependencySnapshot) { model.SnapshotFromComputedOutput = false }, want: Point12ValEStateBlocked},
		{name: "vald proof chain blocked blocks", mutate: func(model *Point12ValEDependencySnapshot) {
			model.ValDProofChainState = Point12ValDProofChainStateBlocked
		}, want: Point12ValEStateBlocked},
		{name: "external api use blocks", mutate: func(model *Point12ValEDependencySnapshot) { model.ValDExternalAPIUsed = true }, want: Point12ValEStateBlocked},
		{name: "premature point12 pass blocks", mutate: func(model *Point12ValEDependencySnapshot) { model.ValDPrematurePoint12PassSeen = true }, want: Point12ValEStateBlocked},
		{name: "wrong point id blocks", mutate: func(model *Point12ValEDependencySnapshot) { model.ValDPointID = "point_11" }, want: Point12ValEStateBlocked},
		{name: "whitespace retagged vald point id blocks", mutate: func(model *Point12ValEDependencySnapshot) {
			model.ValDPointID = " " + point12Val0PointID + " "
		}, want: Point12ValEStateBlocked},
		{name: "tab newline retagged vald active state blocks", mutate: func(model *Point12ValEDependencySnapshot) {
			model.ValDCurrentState = "\t" + Point12ValDStateActive + "\n"
		}, want: Point12ValEStateBlocked},
		{name: "padded vald proof chain state blocks", mutate: func(model *Point12ValEDependencySnapshot) {
			model.ValDProofChainState = Point12ValDProofChainStateActive + " "
		}, want: Point12ValEStateBlocked},
		{name: "point12 pass as input proof blocks", mutate: func(model *Point12ValEDependencySnapshot) {
			model.ValD.Query.RequestedExplanation = point12ValEPoint12PassToken
		}, want: Point12ValEStateBlocked},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint12ValEFoundation().Dependency
			tc.mutate(&model)
			if got := EvaluatePoint12ValEDependencyState(model); got != tc.want {
				t.Fatalf("expected %s, got %#v", tc.want, model)
			}
		})
	}

	t.Run("computed upstream point pass emission propagates into final dependency block", func(t *testing.T) {
		val0 := activePoint12Val0Foundation()
		valA := activePoint12ValAFoundationFromVal0(val0)
		valB := activePoint12ValBFoundationFromValA(valA)
		valB.ReplayResult.PointPassEmitted = true
		valC := activePoint12ValCFoundationFromValB(valB)
		valD := activePoint12ValDFoundationFromValC(valC)

		snapshot := SnapshotPoint12ValEDependencyFromComputed(val0, valA, valB, valC, valD, point12ValEDependencyReviewContextModel())
		if !snapshot.ValDPointPassEmitted {
			t.Fatalf("expected computed upstream point pass emission to propagate into final snapshot, got %#v", snapshot)
		}
		state, reasons := point12ValEDependencyStateAndReasons(snapshot)
		if state != Point12ValEStateBlocked {
			t.Fatalf("expected propagated point pass emission to block final dependency, state=%s reasons=%v snapshot=%#v", state, reasons, snapshot)
		}
		assertPoint12ValEReason(t, reasons, "dependency_valc_premature_point12_pass")
		if !point12Val0StringSliceContains(reasons, "dependency_identity_or_preflight_invalid") {
			t.Fatalf("expected exact final preflight reason for propagated point pass emission, got %v", reasons)
		}
	})

	t.Run("unsafe nested ValC ValA profile context blocks dependency with exact reason", func(t *testing.T) {
		model := activePoint12ValEFoundation().Dependency
		model.ValC.Dependency.ValAManifest.ProfileContext.CurrentProfileHash = ""
		got, reasons := point12ValEDependencyStateAndReasons(model)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected nested ValC dependency profile mutation to block, got state=%s reasons=%v", got, reasons)
		}
		assertPoint12ValEReason(t, reasons, "dependency_valc_profile_context_binding_invalid")
	})

	t.Run("stale nested ValC dependency pass emission blocks dependency with exact reason", func(t *testing.T) {
		model := activePoint12ValEFoundation().Dependency
		model.ValC.Dependency.ValBPointPassEmitted = true
		got, reasons := point12ValEDependencyStateAndReasons(model)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected stale nested ValC dependency pass emission to block, got state=%s reasons=%v", got, reasons)
		}
		assertPoint12ValEReason(t, reasons, "dependency_valc_premature_point12_pass")
	})

	t.Run("stale embedded ValC ValB replay result pass emission blocks dependency with exact reason", func(t *testing.T) {
		model := activePoint12ValEFoundation().Dependency
		model.ValC.Dependency.ValBReplayResult.PointPassEmitted = true
		got, reasons := point12ValEDependencyStateAndReasons(model)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected stale embedded ValC ValB replay result pass emission to block, got state=%s reasons=%v", got, reasons)
		}
		assertPoint12ValEReason(t, reasons, "dependency_valc_premature_point12_pass")
	})

	t.Run("nested ValC review-required dependency blocks final dependency with exact reason", func(t *testing.T) {
		model := activePoint12ValEFoundation().Dependency
		model.ValC.Dependency.ValBReplayTaxonomy = Point12Val0ReplayResultUnsupportedVersion
		model.ValC.Dependency.ValBReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultUnsupportedVersion
		model.ValC.Dependency.ValBReplayResult.UnsupportedVersion = true
		got, reasons := point12ValEDependencyStateAndReasons(model)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected nested ValC review-required dependency to block final dependency, got state=%s reasons=%v", got, reasons)
		}
		assertPoint12ValEReason(t, reasons, "dependency_valc_review_required_or_non_same_decision")
	})

	t.Run("nested ValC ValB current review-required blocks final dependency with exact reason", func(t *testing.T) {
		model := activePoint12ValEFoundation().Dependency
		model.ValC.Dependency.ValBCurrentState = Point12ValBStateReviewRequired
		got, reasons := point12ValEDependencyStateAndReasons(model)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected nested ValC ValB review-required current state to block final dependency, got state=%s reasons=%v", got, reasons)
		}
		assertPoint12ValEReason(t, reasons, "dependency_valc_review_required_or_non_same_decision")
	})

	t.Run("nested ValC ValB dependency review-required blocks final dependency with exact reason", func(t *testing.T) {
		model := activePoint12ValEFoundation().Dependency
		model.ValC.Dependency.ValBDependencyState = Point12ValBDependencyStateReviewRequired
		got, reasons := point12ValEDependencyStateAndReasons(model)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected nested ValC ValB dependency review-required state to block final dependency, got state=%s reasons=%v", got, reasons)
		}
		assertPoint12ValEReason(t, reasons, "dependency_valc_review_required_or_non_same_decision")
	})

	t.Run("nested ValC review prerequisites block final dependency with exact reason", func(t *testing.T) {
		model := activePoint12ValEFoundation().Dependency
		model.ValC.Dependency.ReviewPrerequisites = []string{"manual_review_required"}
		got, reasons := point12ValEDependencyStateAndReasons(model)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected nested ValC review prerequisite to block final dependency, got state=%s reasons=%v", got, reasons)
		}
		assertPoint12ValEReason(t, reasons, "dependency_valc_review_required_or_non_same_decision")
	})
}

func TestPoint12ValEFinalReplayInvariantState(t *testing.T) {
	t.Run("valid replay invariants active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got := ComputePoint12ValEFoundation(model).ReplayInvariantState; got != Point12ValEStateActive {
			t.Fatalf("expected active replay invariants, got %#v", model.ReplayInvariants)
		}
	})

	cases := []struct {
		name   string
		mutate func(*Point12ValEFinalReplayInvariants, *Point12ValEDependencySnapshot)
		want   string
	}{
		{name: "original context silently using current policy blocks", mutate: func(model *Point12ValEFinalReplayInvariants, _ *Point12ValEDependencySnapshot) {
			model.OriginalContextUsesCurrentPolicySilently = true
		}, want: Point12ValEStateBlocked},
		{name: "comparison mode missing drift explanation blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultDifferentDecision
			model.ReplayMode = point12Val0ReplayModeComparisonMode
			model.ComparisonModeDriftExplanationPresent = false
			model.DecisionDriftReasonsPresent = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency tamper blocks even if local model is benign", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.TamperDetected = true
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.TamperDetected = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency tamper diagnostics block even if flags stay false", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ManifestIntegrityCheckResult = point12ValBCheckResultTampered
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.TamperDetected = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency unsupported version blocks even if local model is benign", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.UnsupportedVersion = true
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.UnsupportedVersion = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency blocked replay blocks even if local model says same decision", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultBlockedReplay
			dependency.ValB.ReplayResult.ReplayState = Point12ValBReplayResultStateBlocked
			dependency.ValB.ReplayResult.BlockedReason = "tamper_detected upstream"
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.BlockedReplay = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency insufficient evidence blocks even if local model is benign", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultInsufficientEvidence
			dependency.ValB.ReplayResult.InsufficientEvidence = true
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.InsufficientEvidence = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency different decision blocks even if local model says same decision", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultDifferentDecision
			dependency.ValB.ReplayResult.DecisionDriftExplanation = "policy drift changed decision"
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.DifferentDecision = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency external api use blocks even if local model is benign", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ExternalAPIUsed = true
			model.ExternalAPIUsed = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency point pass emission blocks even if local model is benign", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.PointPassEmitted = true
			model.PointPassEmittedOutsideValE = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency evidence mismatch blocks even if ids and hashes in model are benign", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultEvidenceMismatch
			dependency.ValB.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeEvidenceMismatch, point12ValBDriftDueToEvidence)}
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.EvidenceMismatch = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency policy mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultPolicyMismatch
			dependency.ValB.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBPolicyMismatch(true)}
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.PolicyMismatch = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency engine mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultEngineMismatch
			dependency.ValB.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeEngineMismatch, point12ValBDriftDueToEngine)}
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.EngineMismatch = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency schema mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultSchemaMismatch
			dependency.ValB.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeSchemaMismatch, point12ValBDriftDueToSchema)}
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.SchemaMismatch = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency claim mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultClaimMismatch
			dependency.ValB.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeClaimMismatch, point12ValBDriftDueToClaim)}
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.ClaimMismatch = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency governance mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultGovernanceMismatch
			dependency.ValB.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeGovernanceMismatch, point12ValBDriftDueToGovernance)}
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.GovernanceMismatch = false
		}, want: Point12ValEStateBlocked},
		{name: "dependency redaction limitations block deterministic final replay", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultRedactedLimitations
			dependency.ValB.ReplayResult.RedactionLimitations = true
			dependency.ValB.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeRedactionMismatch, point12ValBDriftDueToRedaction)}
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
			model.RedactionLimitations = false
		}, want: Point12ValEStateBlocked},
		{name: "replay mode mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayMode = point12Val0ReplayModeOriginalContext
			model.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		}, want: Point12ValEStateBlocked},
		{name: "replay taxonomy mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultBlockedReplay
			dependency.ValB.ReplayResult.ReplayState = Point12ValBReplayResultStateBlocked
			model.ReplayResultTaxonomy = Point12Val0ReplayResultSameDecision
		}, want: Point12ValEStateBlocked},
		{name: "original decision state mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.OriginalDecisionState = "decision_state_original_dependency"
			model.OriginalDecisionState = "decision_state_original_local"
		}, want: Point12ValEStateBlocked},
		{name: "replayed decision state mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayedDecisionState = "decision_state_replayed_dependency"
			model.ReplayedDecisionState = "decision_state_replayed_local"
		}, want: Point12ValEStateBlocked},
		{name: "match original mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.MatchOriginal = false
			model.MatchOriginal = true
		}, want: Point12ValEStateBlocked},
		{name: "decision drift explanation mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultDifferentDecision
			dependency.ValB.ReplayResult.DecisionDriftExplanation = "dependency drift explanation"
			model.ReplayResultTaxonomy = Point12Val0ReplayResultDifferentDecision
			model.DecisionDriftExplanation = "local benign explanation"
			model.DecisionDriftReasonsPresent = true
			model.ComparisonModeDriftExplanationPresent = true
		}, want: Point12ValEStateBlocked},
		{name: "mismatch explanations mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultPolicyMismatch
			dependency.ValB.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBPolicyMismatch(true)}
			dependency.ValB.ReplayResult.MismatchExplanations = []string{"dependency mismatch explanation"}
			model.ReplayResultTaxonomy = Point12Val0ReplayResultPolicyMismatch
			model.MismatchExplanations = []string{"local benign explanation"}
			model.MismatchExpectedActualPresent = true
		}, want: Point12ValEStateBlocked},
		{name: "manifest integrity check result mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ManifestIntegrityCheckResult = point12ValBCheckResultMismatch
			model.ManifestIntegrityCheckResult = point12ValBCheckResultActive
		}, want: Point12ValEStateBlocked},
		{name: "signature metadata check result mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.SignatureMetadataCheckResult = point12ValBCheckResultMismatch
			model.SignatureMetadataCheckResult = point12ValBCheckResultActive
		}, want: Point12ValEStateBlocked},
		{name: "compatibility check result mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.CompatibilityCheckResult = point12ValBCheckResultMismatch
			model.CompatibilityCheckResult = point12ValBCheckResultActive
		}, want: Point12ValEStateBlocked},
		{name: "evidence hash check result mismatch blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.EvidenceHashCheckResult = point12ValBCheckResultMismatch
			model.EvidenceHashCheckResult = point12ValBCheckResultActive
		}, want: Point12ValEStateBlocked},
		{name: "self consistent blocked manifest check result blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.ManifestIntegrityCheckResult = point12ValBCheckResultBlocked
			model.ManifestIntegrityCheckResult = point12ValBCheckResultBlocked
		}, want: Point12ValEStateBlocked},
		{name: "self consistent blocked compatibility check result blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.CompatibilityCheckResult = point12ValBCheckResultBlocked
			model.CompatibilityCheckResult = point12ValBCheckResultBlocked
		}, want: Point12ValEStateBlocked},
		{name: "missing mismatch expected actual blocks", mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot) {
			dependency.ValB.ReplayResult.Mismatches = []Point12ValBReplayMismatch{{
				MismatchID:      "mismatch_point12_vale_001",
				MismatchType:    point12ValBMismatchTypePolicyMismatch,
				ExpectedRef:     "policy_ref_point12_expected",
				ActualRef:       "policy_ref_point12_actual",
				ExpectedHash:    "sha256:abababababababababababababababababababababababababababababababab",
				ActualHash:      "sha256:bcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbcbc",
				ExpectedVersion: "policy_version_point12_expected",
				ActualVersion:   "policy_version_point12_actual",
			}}
			model.MismatchExpectedActualPresent = false
		}, want: Point12ValEStateBlocked},
		{name: "point pass outside vale blocks", mutate: func(model *Point12ValEFinalReplayInvariants, _ *Point12ValEDependencySnapshot) {
			model.PointPassEmittedOutsideValE = true
		}, want: Point12ValEStateBlocked},
		{name: "redacted decisive evidence blocks", mutate: func(model *Point12ValEFinalReplayInvariants, _ *Point12ValEDependencySnapshot) {
			model.RedactedDecisiveEvidence = true
		}, want: Point12ValEStateBlocked},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			foundation := activePoint12ValEFoundation()
			tc.mutate(&foundation.ReplayInvariants, &foundation.Dependency)
			if got, _ := point12ValEFinalReplayInvariantStateAndReasons(foundation.ReplayInvariants, foundation.Dependency); got != tc.want {
				t.Fatalf("expected %s, got %#v", tc.want, foundation.ReplayInvariants)
			}
		})
	}

	t.Run("tab newline replay mode retag fails closed with exact reason", func(t *testing.T) {
		foundation := activePoint12ValEFoundation()
		foundation.ReplayInvariants.ReplayMode = "\t" + foundation.ReplayInvariants.ReplayMode + "\n"
		got, reasons := point12ValEFinalReplayInvariantStateAndReasons(foundation.ReplayInvariants, foundation.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked replay invariant for raw retag, got %#v", foundation.ReplayInvariants)
		}
		found := false
		for _, reason := range reasons {
			if reason == "replay_invariant_identity_or_metadata_invalid" || reason == "replay_invariant_dependency_semantics_mismatch:replay_mode" {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected exact replay-mode raw mismatch reason, got %#v", reasons)
		}
	})

	t.Run("padded replay review id fails raw exact identity", func(t *testing.T) {
		foundation := activePoint12ValEFoundation()
		foundation.ReplayInvariants.ReviewID = " " + foundation.ReplayInvariants.ReviewID + " "
		got, reasons := point12ValEFinalReplayInvariantStateAndReasons(foundation.ReplayInvariants, foundation.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked replay invariant for padded review id, got %#v", foundation.ReplayInvariants)
		}
		assertPoint12ValEReason(t, reasons, "replay_invariant_identity_or_metadata_invalid")
	})

	t.Run("padded dependency replay taxonomy fails raw exact semantics", func(t *testing.T) {
		foundation := activePoint12ValEFoundation()
		foundation.Dependency.ValB.ReplayResult.ReplayResultTaxonomy = foundation.Dependency.ValB.ReplayResult.ReplayResultTaxonomy + " "
		got, reasons := point12ValEFinalReplayInvariantStateAndReasons(foundation.ReplayInvariants, foundation.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked replay invariant for padded taxonomy, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "replay_invariant_dependency_semantics_mismatch:replay_result_taxonomy")
	})

	t.Run("tab newline dependency replay state fails raw exact semantics", func(t *testing.T) {
		foundation := activePoint12ValEFoundation()
		foundation.Dependency.ValB.ReplayResult.ReplayState = foundation.Dependency.ValB.ReplayResult.ReplayState + "\n"
		got, reasons := point12ValEFinalReplayInvariantStateAndReasons(foundation.ReplayInvariants, foundation.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked replay invariant for retagged replay state, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "replay_invariant_dependency_semantics_mismatch:blocked_replay")
		assertPoint12ValEReason(t, reasons, "replay_invariant_dependency_blocked_replay")
	})

	t.Run("self consistent blocked check result fails closed with exact reason", func(t *testing.T) {
		foundation := activePoint12ValEFoundation()
		foundation.Dependency.ValB.ReplayResult.ManifestIntegrityCheckResult = point12ValBCheckResultBlocked
		foundation.ReplayInvariants.ManifestIntegrityCheckResult = point12ValBCheckResultBlocked
		got, reasons := point12ValEFinalReplayInvariantStateAndReasons(foundation.ReplayInvariants, foundation.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked replay invariant for self-consistent blocked check result, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "replay_invariant_dependency_semantics_mismatch:blocked_replay")
	})

	t.Run("self consistent blocked compatibility check fails closed with exact reason", func(t *testing.T) {
		foundation := activePoint12ValEFoundation()
		foundation.Dependency.ValB.ReplayResult.CompatibilityCheckResult = point12ValBCheckResultBlocked
		foundation.ReplayInvariants.CompatibilityCheckResult = point12ValBCheckResultBlocked
		got, reasons := point12ValEFinalReplayInvariantStateAndReasons(foundation.ReplayInvariants, foundation.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked replay invariant for self-consistent blocked compatibility check, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "replay_invariant_dependency_semantics_mismatch:blocked_replay")
	})

	for _, tc := range []struct {
		name   string
		value  string
		mutate func(*Point12ValEFinalReplayInvariants, *Point12ValEDependencySnapshot, string)
	}{
		{
			name:  "self consistent unsupported signature metadata check fails closed",
			value: point12ValBCheckResultUnsupported,
			mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot, value string) {
				dependency.ValB.ReplayResult.SignatureMetadataCheckResult = value
				model.SignatureMetadataCheckResult = value
			},
		},
		{
			name:  "self consistent missing signature metadata check fails closed",
			value: point12ValBCheckResultMissing,
			mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot, value string) {
				dependency.ValB.ReplayResult.SignatureMetadataCheckResult = value
				model.SignatureMetadataCheckResult = value
			},
		},
		{
			name:  "self consistent missing compatibility check fails closed",
			value: point12ValBCheckResultMissing,
			mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot, value string) {
				dependency.ValB.ReplayResult.CompatibilityCheckResult = value
				model.CompatibilityCheckResult = value
			},
		},
		{
			name:  "self consistent tampered compatibility check fails closed",
			value: point12ValBCheckResultTampered,
			mutate: func(model *Point12ValEFinalReplayInvariants, dependency *Point12ValEDependencySnapshot, value string) {
				dependency.ValB.ReplayResult.CompatibilityCheckResult = value
				model.CompatibilityCheckResult = value
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			foundation := activePoint12ValEFoundation()
			tc.mutate(&foundation.ReplayInvariants, &foundation.Dependency, tc.value)
			got, reasons := point12ValEFinalReplayInvariantStateAndReasons(foundation.ReplayInvariants, foundation.Dependency)
			if got != Point12ValEStateBlocked {
				t.Fatalf("expected self-consistent non-active check result to block, got state=%s reasons=%v", got, reasons)
			}
			assertPoint12ValEReason(t, reasons, "replay_invariant_dependency_semantics_mismatch:blocked_replay")
			assertPoint12ValEReason(t, reasons, "replay_invariant_dependency_blocked_replay")
		})
	}

	t.Run("same decision dependency replay with matching semantics remains active", func(t *testing.T) {
		foundation := activePoint12ValEFoundation()
		if got, _ := point12ValEFinalReplayInvariantStateAndReasons(foundation.ReplayInvariants, foundation.Dependency); got != Point12ValEStateActive {
			t.Fatalf("expected active replay invariants, got %#v", foundation.ReplayInvariants)
		}
	})
}

func TestPoint12ValEEvidenceQualityState(t *testing.T) {
	t.Run("valid evidence quality active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got, _ := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency); got != Point12ValEStateActive {
			t.Fatalf("expected active evidence quality state, got %#v", model.EvidenceQualityMap)
		}
	})

	t.Run("missing decisive evidence incomplete", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.EvidenceQualityMap.MissingRefs = []string{model.EvidenceQualityMap.EvidenceRefs[0]}
		if got, _ := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency); got != Point12ValEStateIncomplete {
			t.Fatalf("expected incomplete evidence quality state, got %#v", model.EvidenceQualityMap)
		}
	})

	t.Run("stale evidence degrades and does not pass", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.EvidenceQualityMap.StaleRefs = []string{model.EvidenceQualityMap.EvidenceRefs[0]}
		if got, _ := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency); got == Point12ValEStateActive {
			t.Fatalf("expected stale evidence to avoid active state, got %#v", model.EvidenceQualityMap)
		}
	})

	t.Run("revoked evidence blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.EvidenceQualityMap.RevokedRefs = []string{model.EvidenceQualityMap.EvidenceRefs[0]}
		if got, _ := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked evidence quality state, got %#v", model.EvidenceQualityMap)
		}
	})

	t.Run("cross tenant evidence blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.EvidenceQualityMap.CrossTenantRefs = []string{model.EvidenceQualityMap.EvidenceRefs[0]}
		if got, _ := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked evidence quality state, got %#v", model.EvidenceQualityMap)
		}
	})

	t.Run("tampered evidence blocks as tampered", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.EvidenceQualityMap.TamperedRefs = []string{model.EvidenceQualityMap.EvidenceRefs[0]}
		if got, _ := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency); got != Point12ValEStateTampered {
			t.Fatalf("expected tampered evidence quality state, got %#v", model.EvidenceQualityMap)
		}
	})

	t.Run("unsupported evidence blocks as unsupported", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.EvidenceQualityMap.UnsupportedRefs = []string{model.EvidenceQualityMap.EvidenceRefs[0]}
		if got, _ := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency); got != Point12ValEStateUnsupported {
			t.Fatalf("expected unsupported evidence quality state, got %#v", model.EvidenceQualityMap)
		}
	})

	t.Run("padded dependency policy hash fails exact evidence binding", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValA.Manifest.PolicyHash = model.EvidenceQualityMap.PolicyHash + " "
		got, reasons := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked evidence quality state for padded policy hash, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "evidence_quality_map_binding_mismatch")
	})

	t.Run("tab newline dependency artifact ref fails exact evidence binding", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValD.ProofChain.ArtifactRef = model.EvidenceQualityMap.ArtifactRef + "\n"
		got, reasons := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked evidence quality state for retagged artifact ref, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "evidence_quality_map_binding_mismatch")
	})

	t.Run("padded evidence tenant scope fails raw exact identity", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.EvidenceQualityMap.TenantScope = " " + model.EvidenceQualityMap.TenantScope + " "
		got, reasons := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked evidence quality state for padded tenant scope, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "evidence_quality_map_identity_or_metadata_invalid")
	})

	t.Run("tab newline dependency tenant scope fails exact evidence binding", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValD.ProofChain.TenantScope = model.EvidenceQualityMap.TenantScope + "\n"
		got, reasons := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked evidence quality state for retagged tenant scope, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "evidence_quality_map_binding_mismatch")
	})

	t.Run("padded manifest tenant scope fails sibling evidence binding", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValA.Manifest.TenantScope = model.EvidenceQualityMap.TenantScope + " "
		got, reasons := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked evidence quality state for padded manifest tenant scope, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "evidence_quality_map_binding_mismatch")
	})

	t.Run("padded evidence quality state fails exact state validation", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.EvidenceQualityMap.QualityState = Point12ValEStateActive + " "
		got, reasons := point12ValEEvidenceQualityStateAndReasons(model.EvidenceQualityMap, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked evidence quality state for padded state, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "evidence_quality_map_identity_or_metadata_invalid")
	})
}

func TestPoint12ValEBindingMutationClosure(t *testing.T) {
	t.Run("valid binding mutation closure active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got, reasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency); got != Point12ValEStateActive {
			t.Fatalf("expected active binding mutation state, got %s reasons=%v model=%#v", got, reasons, model.BindingMutationClosure)
		}
	})

	t.Run("padded binding mutation review id blocks raw exact metadata", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.BindingMutationClosure.ReviewID = " " + model.BindingMutationClosure.ReviewID + " "
		got, reasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected padded binding mutation review id to block, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "binding_mutation_identity_or_metadata_invalid")
	})

	t.Run("vala schema drift plus recomputed payload hash still blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValA.Manifest.SchemaHash = "sha256:abababababababababababababababababababababababababababababababab"
		model.Dependency.ValA.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Dependency.ValA.Manifest)
		model.Dependency.ValA.Manifest.SignatureBoundManifestPayloadHash = model.Dependency.ValA.Manifest.ManifestPayloadHash
		if got, _ := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked binding mutation state, got %#v", model.Dependency.ValA.Manifest)
		}
	})

	t.Run("valc coordinated redaction manifest substitution still blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValB.ReplayRequest.RedactionManifestRef = "redaction_manifest_point12_vale_substituted"
		model.Dependency.ValC.RedactionManifest.RedactionManifestID = "redaction_manifest_point12_vale_substituted"
		model.Dependency.ValC.RedactionImpactVerdict.RedactionManifestID = "redaction_manifest_point12_vale_substituted"
		if got, _ := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked coordinated redaction substitution, got %#v", model.Dependency.ValC)
		}
	})

	t.Run("vald wrong lineage endpoint plus recomputed projection hash still blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValD.ProofChain.LineageEdges[1].ToRef = "artifact_point12_vale_other"
		recomputePoint12ValDLocalHashes(&model.Dependency.ValD)
		if got, _ := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked wrong lineage endpoint, got %#v", model.Dependency.ValD.ProofChain)
		}
	})

	t.Run("vald wrong lineage hash plus recomputed projection hash still blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValD.ProofChain.LineageEdges[1].FromHash = "sha256:cdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcd"
		recomputePoint12ValDLocalHashes(&model.Dependency.ValD)
		if got, _ := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked wrong lineage hash, got %#v", model.Dependency.ValD.ProofChain)
		}
	})

	t.Run("intentionally not bound field without reason blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		for idx := range model.Dependency.ValD.BindingMatrix.BoundFields {
			if model.Dependency.ValD.BindingMatrix.BoundFields[idx].BindingClass == point12ValDBindingClassIntentionallyNotBound {
				model.Dependency.ValD.BindingMatrix.BoundFields[idx].Reason = ""
				break
			}
		}
		if got, _ := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked binding mutation state, got %#v", model.Dependency.ValD.BindingMatrix)
		}
	})

	t.Run("padded valb request manifest binding fails raw exact", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValB.ReplayRequest.ProofPackID = model.Dependency.ValA.Manifest.ProofPackID + " "
		got, reasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked binding mutation state for padded proof pack, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "binding_mutation_valb_request_manifest_binding_invalid")
	})

	t.Run("tab newline valb result binding fails raw exact", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValB.ReplayResult.ReplayRequestID = model.Dependency.ValB.ReplayRequest.ReplayRequestID + "\n"
		got, reasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked binding mutation state for retagged replay result, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "binding_mutation_valb_result_request_binding_invalid")
	})

	t.Run("padded valc export binding fails raw exact", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValC.ExportBundle.ManifestID = model.Dependency.ValB.ReplayRequest.ManifestID + " "
		got, reasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked binding mutation state for padded export manifest, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "binding_mutation_valc_export_binding_invalid")
	})

	t.Run("tab newline valc offline binding fails raw exact", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValC.OfflineBundle.ArtifactHash = model.Dependency.ValB.ReplayRequest.ArtifactHash + "\t"
		got, reasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked binding mutation state for retagged offline artifact hash, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "binding_mutation_valc_offline_binding_invalid")
	})

	t.Run("padded binding matrix field name no longer satisfies required field", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		mutated := false
		for idx := range model.Dependency.ValD.BindingMatrix.BoundFields {
			if model.Dependency.ValD.BindingMatrix.BoundFields[idx].FieldName == "export_id" {
				model.Dependency.ValD.BindingMatrix.BoundFields[idx].FieldName = " export_id "
				mutated = true
			}
		}
		if !mutated {
			t.Fatalf("expected fixture to contain export_id binding field")
		}
		got, reasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked binding mutation state for padded binding field name, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "binding_mutation_required_vald_field_missing:export_id")
	})
}

func TestPoint12ValEProjectionBoundaryState(t *testing.T) {
	t.Run("valid projection boundary active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got, _ := point12ValEProjectionBoundaryStateAndReasons(model.ProjectionBoundary); got != Point12ValEStateActive {
			t.Fatalf("expected active projection boundary state, got %#v", model.ProjectionBoundary)
		}
	})

	t.Run("emitting point12 pass from projection blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.ProjectionBoundary.EmitsPoint12Pass = true
		if got, _ := point12ValEProjectionBoundaryStateAndReasons(model.ProjectionBoundary); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked projection boundary state, got %#v", model.ProjectionBoundary)
		}
	})

	t.Run("authority creating buyer text blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.ProjectionBoundary.BuyerProductCustomerTextCreatesAuthority = true
		if got, _ := point12ValEProjectionBoundaryStateAndReasons(model.ProjectionBoundary); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked projection boundary state, got %#v", model.ProjectionBoundary)
		}
	})
}

func TestPoint12ValENoOverclaimState(t *testing.T) {
	t.Run("valid no overclaim review active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got, _ := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview); got != Point12ValEStateActive {
			t.Fatalf("expected active no-overclaim state, got %#v", model.NoOverclaimReview)
		}
	})

	t.Run("padded no overclaim review id blocks raw exact metadata", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ReviewID = " " + model.NoOverclaimReview.ReviewID + " "
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected padded no-overclaim review id to block, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_identity_or_metadata_invalid")
	})

	t.Run("forbidden customer wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"regulator-approved deployment approved"}
		if got, _ := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked no-overclaim state, got %#v", model.NoOverclaimReview)
		}
	})

	t.Run("unhyphenated regulator approval wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"regulator approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected unhyphenated regulator approval to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("unicode dash regulator approval wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"regulator\u2011approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected unicode dash regulator approval to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("slash regulator approval wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"regulator/approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected slash-separated regulator approval to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("dot production approval wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"production.approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected dot-separated production approval to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("inherited point10 readiness overclaim wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"marketplace production ready"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected inherited readiness overclaim to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("split forbidden export wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"deployment", "approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected split forbidden export wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("slash split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"deployment/", "approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected slash split forbidden export wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("unicode dash split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"deployment\u2013", "approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected unicode dash split forbidden wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("confusable split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"pr\u0254duction", "approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected confusable split forbidden wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("zero width production approval wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"production appro\u200dved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected zero-width production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("zero width separator production approval wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"production\u200bapproved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected zero-width separator production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("greek nu production approval wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"production appro\u03bded"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected greek nu production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("greek upsilon production wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"prod\u03c5ction approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected greek upsilon production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("small cap u production wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"prod\U00001d1cction approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected small-cap u production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("latin upsilon production wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"prod\u028action approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected latin upsilon production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("greek delta approved wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"production approve\u03b4"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected greek delta production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("small cap t official authority wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"official au\U00001d1bhority"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected small-cap t official authority wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("latin alpha global truth wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"glob\u0251l truth"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected latin alpha global truth wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("latin iota official authority wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"off\u0269cial authority"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected latin iota official authority wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("dental click public badge wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"pub\u01c0ic badge"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected dental-click public badge wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("armenian oh compliance guarantee wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"c\u0585mpliance guaranteed"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected armenian-oh compliance guarantee wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("greek eta production wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"productio\u03b7 approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected greek eta production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("latin eng production wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"productio\u014b approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected latin eng production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("cyrillic pe production wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"productio\u043f approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected cyrillic pe production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("zero width split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"deployment", "appro\u2060ved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected zero-width split forbidden wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("math bold split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{
			"\U0001d41d\U0001d41e\U0001d429\U0001d425\U0001d428\U0001d432\U0001d426\U0001d41e\U0001d427\U0001d42d",
			"\U0001d41a\U0001d429\U0001d429\U0001d42b\U0001d428\U0001d42f\U0001d41e\U0001d41d",
		}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected math bold split forbidden wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("word fragment split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"produc", "tion approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected word-fragment split forbidden wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("right leg u split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"prod\uab4e", "ction approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected right-leg u split forbidden wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("latin upsilon split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"prod\u028a", "ction approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected latin upsilon split production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("greek nu split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"production", "appro\u03bded"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected greek nu split production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("greek delta split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"production", "approve\u03b4"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected greek delta split production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("small cap t split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"official au", "\U00001d1bhority"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected small-cap t split official authority wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("latin alpha split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"glob\u0251l", "truth"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected latin alpha split global truth wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("latin iota split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"off\u0269cial", "authority"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected latin iota split official authority wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("dental click split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"pub\u01c0ic", "badge"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected dental-click split public badge wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("armenian oh split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"c\u0585mpliance", "guaranteed"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected armenian-oh split compliance guarantee wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("armenian vo split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"productio\u0578", "approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected armenian vo split production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("latin n with long right leg split forbidden wording across buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"productio\u019e", "approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected latin n with long right leg split production approved wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("underscore machine token remains non-boundary safe wording", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.InternalDiagnosticTexts = []string{"internal_production_approved_metric"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateActive {
			t.Fatalf("expected underscore machine token not to become a forbidden phrase, got %#v reasons %#v", model.NoOverclaimReview, reasons)
		}
	})

	t.Run("split forbidden wording across customer and export buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"deployment"}
		model.NoOverclaimReview.ObservedExportTexts = []string{"approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected cross-surface split forbidden wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("split regulator approval across customer and export buckets blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"regulator"}
		model.NoOverclaimReview.ObservedExportTexts = []string{"approved"}
		got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected split regulator approval wording to block, got %#v", model.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, reasons, "no_overclaim_customer_or_export_overclaim_detected")
	})

	t.Run("all allowed disclaimer only split context remains active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedExportTexts = []string{"not deployment approval", "not production approval"}
		if got, reasons := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview); got != Point12ValEStateActive {
			t.Fatalf("expected allowed disclaimer-only export wording to remain active, got %s reasons=%v", got, reasons)
		}
	})

	t.Run("forbidden wording in blocked ledger allowed when classified", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.BlockedClaimLedger = []string{"proves DORA compliance"}
		model.NoOverclaimReview.BlockedClaimLedgerClassified = true
		if got, _ := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview); got != Point12ValEStateActive {
			t.Fatalf("expected active no-overclaim state for classified blocked ledger, got %#v", model.NoOverclaimReview)
		}
	})

	t.Run("internal diagnostic may mention removed claim", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.InternalDiagnosticTexts = []string{"internal diagnostic: removed proves insurance eligibility claim from export"}
		if got, _ := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview); got != Point12ValEStateActive {
			t.Fatalf("expected active no-overclaim state for internal diagnostic, got %#v", model.NoOverclaimReview)
		}
	})
}

func TestPoint12ValECleanRoomIPState(t *testing.T) {
	t.Run("valid clean room review active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got, _ := point12ValECleanRoomIPStateAndReasons(model.CleanRoomIPReview); got != Point12ValEStateActive {
			t.Fatalf("expected active clean-room state, got %#v", model.CleanRoomIPReview)
		}
	})

	t.Run("padded clean room review id blocks raw exact metadata", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.CleanRoomIPReview.ReviewID = " " + model.CleanRoomIPReview.ReviewID + " "
		got, reasons := point12ValECleanRoomIPStateAndReasons(model.CleanRoomIPReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected padded clean-room review id to block, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "clean_room_ip_identity_or_metadata_invalid")
	})

	t.Run("unreviewed customer facing dependency requires review", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.CleanRoomIPReview.UnreviewedCustomerFacingDependency = true
		if got, _ := point12ValECleanRoomIPStateAndReasons(model.CleanRoomIPReview); got != Point12ValEStateReviewRequired {
			t.Fatalf("expected review-required clean-room state, got %#v", model.CleanRoomIPReview)
		}
	})

	t.Run("competitor copy blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.CleanRoomIPReview.CompetitorCopyDetected = true
		if got, _ := point12ValECleanRoomIPStateAndReasons(model.CleanRoomIPReview); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked clean-room state, got %#v", model.CleanRoomIPReview)
		}
	})

	t.Run("padded third party ref blocks raw exact clean room review", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.CleanRoomIPReview.ThirdPartyRefs[0] = " " + model.CleanRoomIPReview.ThirdPartyRefs[0] + " "
		got, reasons := point12ValECleanRoomIPStateAndReasons(model.CleanRoomIPReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected padded third-party ref to block clean-room review, got %#v", model.CleanRoomIPReview)
		}
		assertPoint12ValEReason(t, reasons, "clean_room_ip_identity_or_metadata_invalid")
	})

	t.Run("padded license review ref blocks raw exact clean room review", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.CleanRoomIPReview.LicenseReviewRefs[0] = " " + model.CleanRoomIPReview.LicenseReviewRefs[0] + " "
		got, reasons := point12ValECleanRoomIPStateAndReasons(model.CleanRoomIPReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected padded license review ref to block clean-room review, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "clean_room_ip_identity_or_metadata_invalid")
	})

	t.Run("tab newline ip review ref blocks raw exact clean room review", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.CleanRoomIPReview.IPReviewRefs[0] = "\t" + model.CleanRoomIPReview.IPReviewRefs[0] + "\n"
		got, reasons := point12ValECleanRoomIPStateAndReasons(model.CleanRoomIPReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected retagged IP review ref to block clean-room review, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "clean_room_ip_identity_or_metadata_invalid")
	})

	t.Run("padded ai review package ref blocks raw exact clean room review", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.CleanRoomIPReview.AIReviewPackageRefs[0] = " " + model.CleanRoomIPReview.AIReviewPackageRefs[0] + " "
		got, reasons := point12ValECleanRoomIPStateAndReasons(model.CleanRoomIPReview)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected padded AI review package ref to block clean-room review, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "clean_room_ip_identity_or_metadata_invalid")
	})
}

func TestPoint12ValERetentionProvenanceState(t *testing.T) {
	t.Run("valid retention provenance active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got, _ := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency); got != Point12ValEStateActive {
			t.Fatalf("expected active retention/provenance state, got %#v", model.RetentionProvenanceReview)
		}
	})

	t.Run("padded retention review id blocks raw exact metadata", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.RetentionProvenanceReview.ReviewID = " " + model.RetentionProvenanceReview.ReviewID + " "
		got, reasons := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected padded retention review id to block, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "retention_provenance_identity_or_metadata_invalid")
	})

	t.Run("missing retention class blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.RetentionProvenanceReview.ProofPackRetentionClassRef = ""
		if got, _ := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked retention/provenance state, got %#v", model.RetentionProvenanceReview)
		}
	})

	t.Run("missing decisive toolchain provenance requires review", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.RetentionProvenanceReview.ToolchainProvenanceRefs = nil
		if got, _ := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency); got != Point12ValEStateReviewRequired {
			t.Fatalf("expected review-required retention/provenance state, got %#v", model.RetentionProvenanceReview)
		}
	})

	t.Run("agent lineage cannot become authoritative", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.RetentionProvenanceReview.AgentLineageAdvisoryOnly = false
		if got, _ := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked retention/provenance state, got %#v", model.RetentionProvenanceReview)
		}
	})

	t.Run("padded dependency retention class fails raw exact binding", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValA.Manifest.RetentionClassRef = model.RetentionProvenanceReview.ProofPackRetentionClassRef + " "
		got, reasons := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked retention/provenance state for padded retention class, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "retention_provenance_binding_mismatch")
	})

	t.Run("tab newline dependency disposal path fails raw exact binding", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValC.ExportBundle.DisposalPathRef = model.RetentionProvenanceReview.DisposalPathRef + "\n"
		got, reasons := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked retention/provenance state for retagged disposal path, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "retention_provenance_binding_mismatch")
	})

	t.Run("padded retention tenant scope fails raw exact identity", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.RetentionProvenanceReview.TenantScope = " " + model.RetentionProvenanceReview.TenantScope + " "
		got, reasons := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked retention/provenance state for padded tenant scope, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "retention_provenance_identity_or_metadata_invalid")
	})

	t.Run("tab newline proof chain tenant scope fails retention binding", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValD.ProofChain.TenantScope = model.RetentionProvenanceReview.TenantScope + "\n"
		got, reasons := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked retention/provenance state for retagged tenant scope, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "retention_provenance_binding_mismatch")
	})

	t.Run("padded export tenant scope fails sibling retention binding", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.Dependency.ValC.ExportBundle.TenantScope = model.RetentionProvenanceReview.TenantScope + " "
		got, reasons := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked retention/provenance state for padded export tenant scope, got %s", got)
		}
		assertPoint12ValEReason(t, reasons, "retention_provenance_binding_mismatch")
	})
}

func TestPoint12ValEPassClosureManifestState(t *testing.T) {
	t.Run("valid complete pass closure manifest active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got, reasons := point12ValEPassClosureManifestStateAndReasons(model.PassClosureManifest, model, true); got != Point12ValEStateActive {
			t.Fatalf("expected active pass closure manifest state, got %s reasons=%v model=%#v", got, reasons, model.PassClosureManifest)
		}
	})

	cases := []struct {
		name       string
		expected   bool
		mutate     func(*Point12ValEPassClosureManifest)
		want       string
		wantReason string
	}{
		{name: "missing closure manifest id blocks", expected: true, mutate: func(model *Point12ValEPassClosureManifest) { model.ClosureManifestID = "" }, want: Point12ValEStateBlocked},
		{name: "wrong point id blocks", expected: true, mutate: func(model *Point12ValEPassClosureManifest) { model.PointID = "point_11" }, want: Point12ValEStateBlocked},
		{name: "padded point id blocks", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			model.PointID = " " + model.PointID + " "
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_identity_or_metadata_invalid"},
		{name: "tab newline wave id blocks", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			model.WaveID = "\t" + model.WaveID + "\n"
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_identity_or_metadata_invalid"},
		{name: "missing dependency refs block", expected: true, mutate: func(model *Point12ValEPassClosureManifest) { model.ValDSnapshotRef = "" }, want: Point12ValEStateBlocked},
		{name: "commit sha present before commit blocks", expected: true, mutate: func(model *Point12ValEPassClosureManifest) { model.CommitSHAIfAvailable = "abc123" }, want: Point12ValEStateBlocked},
		{name: "padded final point12 token blocks", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			model.Point12PassToken = point12ValEPoint12PassToken + " "
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_token_invalid"},
		{name: "padded generated timestamp blocks raw exact closure manifest", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			model.GeneratedAt = " " + model.GeneratedAt + " "
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "fake prefix-shaped gate refs cannot replace canonical run evidence", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			model.CommandsRun = []string{"command_run_point12_vale_fake_001", "command_run_point12_vale_fake_002", "command_run_point12_vale_fake_003"}
			model.TestsRun = []string{"test_run_point12_vale_fake_001", "test_run_point12_vale_fake_002", "test_run_point12_vale_fake_003"}
			model.NegativeFixturesRun = []string{"negative_fixture_point12_vale_fake_001", "negative_fixture_point12_vale_fake_002", "negative_fixture_point12_vale_fake_003"}
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "duplicate command run cannot satisfy missing canonical command run", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			expected := point12ValECommandsRun()
			model.CommandsRun = []string{expected[0], expected[0], expected[2]}
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "duplicate test run cannot satisfy missing canonical test run", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			expected := point12ValETestsRun()
			model.TestsRun = []string{expected[0], expected[0], expected[2]}
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "duplicate negative fixture cannot satisfy missing canonical negative fixture", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			expected := point12ValENegativeFixturesRun()
			model.NegativeFixturesRun = []string{expected[0], expected[0], expected[2]}
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "zero width fake command run blocks raw exact gate refs", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			model.CommandsRun = []string{"command_run_point12_vale_\u200dfake_001", "command_run_point12_vale_go_test_formal_001", "command_run_point12_vale_go_test_all_001"}
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "zero width fake test run blocks raw exact gate refs", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			model.TestsRun = []string{"test_run_point12_vale_\u200dfake_001", "test_run_point12_vale_point11_regressions_001", "test_run_point12_vale_go_test_all_001"}
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "zero width fake negative fixture blocks raw exact gate refs", expected: true, mutate: func(model *Point12ValEPassClosureManifest) {
			model.NegativeFixturesRun = []string{"negative_fixture_point12_vale_\u200dfake_001", "negative_fixture_point12_vale_no_overclaim_001", "negative_fixture_point12_vale_binding_mutation_001"}
		}, want: Point12ValEStateBlocked, wantReason: "pass_closure_manifest_required_fields_invalid"},
		{name: "point12 token present before final path blocks", expected: false, mutate: func(model *Point12ValEPassClosureManifest) {}, want: Point12ValEStateBlocked},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			foundation := activePoint12ValEFoundation()
			tc.mutate(&foundation.PassClosureManifest)
			got, reasons := point12ValEPassClosureManifestStateAndReasons(foundation.PassClosureManifest, foundation, tc.expected)
			if got != tc.want {
				t.Fatalf("expected %s, got %#v", tc.want, foundation.PassClosureManifest)
			}
			if tc.wantReason != "" {
				found := false
				for _, reason := range reasons {
					if reason == tc.wantReason {
						found = true
						break
					}
				}
				if !found {
					t.Fatalf("expected exact reason %q, got %#v", tc.wantReason, reasons)
				}
			}
		})
	}

	t.Run("polluted dependency and manifest val0 snapshot ref still blocks", func(t *testing.T) {
		foundation := activePoint12ValEFoundation()
		polluted := " " + foundation.Dependency.Val0SnapshotRef + " "
		foundation.Dependency.Val0SnapshotRef = polluted
		foundation.PassClosureManifest.Val0SnapshotRef = polluted
		got, reasons := point12ValEPassClosureManifestStateAndReasons(foundation.PassClosureManifest, foundation, true)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected polluted dependency-bound val0 snapshot ref to block, got %s reasons=%v", got, reasons)
		}
		assertPoint12ValEReason(t, reasons, "pass_closure_manifest_required_fields_invalid")
	})

	t.Run("polluted dependency and manifest tenant scope still blocks", func(t *testing.T) {
		foundation := activePoint12ValEFoundation()
		polluted := foundation.Dependency.ValD.ProofChain.TenantScope + "\n"
		foundation.Dependency.ValD.ProofChain.TenantScope = polluted
		foundation.PassClosureManifest.TenantScope = polluted
		got, reasons := point12ValEPassClosureManifestStateAndReasons(foundation.PassClosureManifest, foundation, true)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected polluted dependency-bound tenant scope to block, got %s reasons=%v", got, reasons)
		}
		assertPoint12ValEReason(t, reasons, "pass_closure_manifest_required_fields_invalid")
	})

	t.Run("review required cannot mask malformed manifest", func(t *testing.T) {
		foundation := activePoint12ValEFoundation()
		foundation.PassClosureManifest.ReviewerResult = point12ValEReviewerResultReviewRequired
		foundation.PassClosureManifest.Point12PassAllowed = false
		foundation.PassClosureManifest.Point12PassToken = ""
		foundation.PassClosureManifest.ClosureManifestID = ""
		got, reasons := point12ValEPassClosureManifestStateAndReasons(foundation.PassClosureManifest, foundation, true)
		if got != Point12ValEStateBlocked {
			t.Fatalf("expected malformed review-required manifest to block, got %s reasons=%v", got, reasons)
		}
		assertPoint12ValEReason(t, reasons, "pass_closure_manifest_identity_or_metadata_invalid")
	})
}

func TestPoint12ValEAggregateState(t *testing.T) {
	t.Run("valid final happy path emits point12 pass", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if model.CurrentState != Point12ValEStatePassConfirmed {
			t.Fatalf("expected pass confirmed state, got %#v", model)
		}
		if !model.Point12PassAllowed || model.Point12PassToken != point12ValEPoint12PassToken {
			t.Fatalf("expected final point_12_pass token, got %#v", model)
		}
	})

	t.Run("active subgates are not final pass without pass confirmed manifest", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.PassClosureManifest.ReviewerResult = point12ValEReviewerResultPass
		model.PassClosureManifest.Point12PassAllowed = false
		model.PassClosureManifest.Point12PassToken = ""
		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateActive {
			t.Fatalf("expected active aggregate state without final pass token, got %#v", computed)
		}
	})

	t.Run("malformed review required manifest stays blocked in aggregate", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.PassClosureManifest.ReviewerResult = point12ValEReviewerResultReviewRequired
		model.PassClosureManifest.Point12PassAllowed = false
		model.PassClosureManifest.Point12PassToken = ""
		model.PassClosureManifest.ClosureManifestID = ""
		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked || computed.PassClosureManifestState != Point12ValEStateBlocked {
			t.Fatalf("expected malformed review-required aggregate to block, got %#v", computed)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "pass_closure_manifest:"+Point12ValEStateBlocked)
	})

	cases := []struct {
		name   string
		mutate func(*Point12ValEFoundation)
	}{
		{name: "dependency blocked removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValD.Query.RequestedExplanation = point12ValEPoint12PassToken
		}},
		{name: "replay invariant failure removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.ReplayInvariants.OriginalContextUsesCurrentPolicySilently = true
		}},
		{name: "dependency replay tamper removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.TamperDetected = true
		}},
		{name: "dependency replay unsupported removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.UnsupportedVersion = true
		}},
		{name: "dependency replay blocked removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultBlockedReplay
			model.Dependency.ValB.ReplayResult.ReplayState = Point12ValBReplayResultStateBlocked
			model.Dependency.ValB.ReplayResult.BlockedReason = "upstream blocked replay"
		}},
		{name: "dependency replay insufficient evidence removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultInsufficientEvidence
			model.Dependency.ValB.ReplayResult.InsufficientEvidence = true
		}},
		{name: "dependency replay different decision removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultDifferentDecision
			model.Dependency.ValB.ReplayResult.DecisionDriftExplanation = "upstream different decision"
		}},
		{name: "dependency replay mismatch removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultPolicyMismatch
			model.Dependency.ValB.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBPolicyMismatch(true)}
		}},
		{name: "dependency replay external api use removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.ExternalAPIUsed = true
		}},
		{name: "dependency replay point pass emission removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.PointPassEmitted = true
		}},
		{name: "self consistent blocked replay check removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.ManifestIntegrityCheckResult = point12ValBCheckResultBlocked
			model.ReplayInvariants.ManifestIntegrityCheckResult = point12ValBCheckResultBlocked
		}},
		{name: "self consistent blocked compatibility check removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.CompatibilityCheckResult = point12ValBCheckResultBlocked
			model.ReplayInvariants.CompatibilityCheckResult = point12ValBCheckResultBlocked
		}},
		{name: "self consistent unsupported signature metadata check removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.SignatureMetadataCheckResult = point12ValBCheckResultUnsupported
			model.ReplayInvariants.SignatureMetadataCheckResult = point12ValBCheckResultUnsupported
		}},
		{name: "self consistent tampered compatibility check removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValB.ReplayResult.CompatibilityCheckResult = point12ValBCheckResultTampered
			model.ReplayInvariants.CompatibilityCheckResult = point12ValBCheckResultTampered
		}},
		{name: "evidence quality failure removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.EvidenceQualityMap.MissingRefs = []string{model.EvidenceQualityMap.EvidenceRefs[0]}
		}},
		{name: "binding mutation failure removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.Dependency.ValA.Manifest.SchemaHash = "sha256:dededededededededededededededededededededededededededededededede"
			model.Dependency.ValA.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Dependency.ValA.Manifest)
			model.Dependency.ValA.Manifest.SignatureBoundManifestPayloadHash = model.Dependency.ValA.Manifest.ManifestPayloadHash
		}},
		{name: "no overclaim failure removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.NoOverclaimReview.ObservedExportTexts = []string{"deployment approved"}
		}},
		{name: "clean room failure removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.CleanRoomIPReview.CompetitorCopyDetected = true
		}},
		{name: "retention provenance failure removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.RetentionProvenanceReview.RetentionOwnerRef = ""
		}},
		{name: "incomplete pass closure manifest removes point12 pass", mutate: func(model *Point12ValEFoundation) {
			model.PassClosureManifest.ClosureManifestID = ""
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			model := rawPoint12ValEFoundationModel()
			tc.mutate(&model)
			computed := ComputePoint12ValEFoundation(model)
			assertPoint12ValENoPass(t, computed)
		})
	}

	assertBlockedAggregate := func(t *testing.T, computed Point12ValEFoundation, componentState, reason string) {
		t.Helper()
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked || componentState != Point12ValEStateBlocked {
			t.Fatalf("expected blocked aggregate with component %s, got state=%s component=%s model=%#v", reason, computed.CurrentState, componentState, computed)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, reason)
	}

	t.Run("long filler no overclaim failure blocks exact aggregate state", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.NoOverclaimReview.ObservedExportTexts = []string{"production is now fully globally approved"}
		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.NoOverclaimState, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("padded binding mutation review id blocks exact aggregate state", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.BindingMutationClosure.ReviewID = " " + model.BindingMutationClosure.ReviewID + " "
		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("unsafe inherited Val0 profile context blocks final binding mutation gate", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValA.Dependency.Val0Manifest.ProfileContext.ProfileMatchOriginal = false
		model.Dependency.ValA.Dependency.Val0Manifest.ProfileContext.ProfileBindingStatus = Point12Val0ProfileBindingStatusUnsupported
		model.Dependency.ValA.Dependency.Val0Manifest.ProfileContext.ProfileMismatchReason = "profile_unsupported"

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected unsafe inherited profile context to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_vala_dependency_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("unsafe Val0 replay assessment profile context blocks final binding mutation gate", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.Val0.ReplayAssessment.ProfileContext.CurrentProfileHash = ""
		model.Dependency.Val0.ReplayAssessment.ProfileContext.ProfileMatchOriginal = false
		model.Dependency.Val0.ReplayAssessment.ProfileContext.ProfileBindingStatus = Point12Val0ProfileBindingStatusMissingCurrent
		model.Dependency.Val0.ReplayAssessment.ProfileContext.ProfileMismatchReason = "profile_current_hash_missing"

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected unsafe Val0 replay profile context to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_val0_profile_context_binding_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("nested ValB ValA profile context blocks final binding mutation gate", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValB.Dependency.ValAManifest.ProfileContext.CurrentProfileHash = ""
		model.Dependency.ValB.Dependency.ValAManifest.ProfileContext.ProfileMatchOriginal = false
		model.Dependency.ValB.Dependency.ValAManifest.ProfileContext.ProfileBindingStatus = Point12Val0ProfileBindingStatusMissingCurrent
		model.Dependency.ValB.Dependency.ValAManifest.ProfileContext.ProfileMismatchReason = "profile_current_hash_missing"

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected unsafe nested ValB profile context to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_valb_profile_context_binding_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("nested ValC ValA profile context blocks final dependency and binding gates", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.Dependency.ValAManifest.ProfileContext.ProfileApprovalRef = "profile_approval_point_12_pass"

		dependencyState, dependencyReasons := point12ValEDependencyStateAndReasons(model.Dependency)
		if dependencyState != Point12ValEStateBlocked {
			t.Fatalf("expected nested ValC dependency profile mutation to block dependency, got state=%s reasons=%v", dependencyState, dependencyReasons)
		}
		assertPoint12ValEReason(t, dependencyReasons, "dependency_valc_profile_context_binding_invalid")

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected nested ValC dependency profile mutation to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_valc_profile_context_binding_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.DependencyState, "dependency:"+Point12ValEStateBlocked)
	})

	t.Run("stale nested ValC dependency pass emission blocks final dependency and binding gates", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.Dependency.ValBPointPassEmitted = true

		dependencyState, dependencyReasons := point12ValEDependencyStateAndReasons(model.Dependency)
		if dependencyState != Point12ValEStateBlocked {
			t.Fatalf("expected stale nested ValC dependency pass emission to block dependency, got state=%s reasons=%v", dependencyState, dependencyReasons)
		}
		assertPoint12ValEReason(t, dependencyReasons, "dependency_valc_premature_point12_pass")

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected stale nested ValC dependency pass emission to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_valc_premature_point12_pass")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.DependencyState, "dependency:"+Point12ValEStateBlocked)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("stale embedded ValC ValB replay result pass emission blocks final dependency and binding gates", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.Dependency.ValBReplayResult.PointPassEmitted = true

		dependencyState, dependencyReasons := point12ValEDependencyStateAndReasons(model.Dependency)
		if dependencyState != Point12ValEStateBlocked {
			t.Fatalf("expected stale embedded ValC ValB replay result pass emission to block dependency, got state=%s reasons=%v", dependencyState, dependencyReasons)
		}
		assertPoint12ValEReason(t, dependencyReasons, "dependency_valc_premature_point12_pass")

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected stale embedded ValC ValB replay result pass emission to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_valc_premature_point12_pass")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.DependencyState, "dependency:"+Point12ValEStateBlocked)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("nested ValC review-required dependency blocks final dependency and binding gates", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.Dependency.ValBReplayTaxonomy = Point12Val0ReplayResultUnsupportedVersion
		model.Dependency.ValC.Dependency.ValBReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultUnsupportedVersion
		model.Dependency.ValC.Dependency.ValBReplayResult.UnsupportedVersion = true

		dependencyState, dependencyReasons := point12ValEDependencyStateAndReasons(model.Dependency)
		if dependencyState != Point12ValEStateBlocked {
			t.Fatalf("expected nested ValC review-required dependency to block dependency, got state=%s reasons=%v", dependencyState, dependencyReasons)
		}
		assertPoint12ValEReason(t, dependencyReasons, "dependency_valc_review_required_or_non_same_decision")

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected nested ValC review-required dependency to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_valc_review_required_or_non_same_decision")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.DependencyState, "dependency:"+Point12ValEStateBlocked)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("valc export policy hash drift blocks final binding mutation gate", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.ExportBundle.PolicyHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"

		redactionManifestState := EvaluatePoint12ValCRedactionManifestState(model.Dependency.ValC.RedactionManifest, model.Dependency.ValC.Dependency, model.Dependency.ValC.ExportBundle)
		redactionImpactState := EvaluatePoint12ValCRedactionImpactState(model.Dependency.ValC.RedactionImpactVerdict, model.Dependency.ValC.RedactionManifest, model.Dependency.ValC.Dependency)
		offlineState := EvaluatePoint12ValCOfflineBundleState(model.Dependency.ValC.OfflineBundle, model.Dependency.ValC.Dependency, redactionImpactState)
		boundaryState := EvaluatePoint12ValCPublicPrivateBoundaryState(model.Dependency.ValC.PublicPrivateBoundary, model.Dependency.ValC.Dependency, model.Dependency.ValC.ExportBundle, model.Dependency.ValC.OfflineBundle, model.Dependency.ValC.RedactionManifest)
		exportState, exportReasons := point12ValCAuditExportStateAndReasons(model.Dependency.ValC.ExportBundle, model.Dependency.ValC.Dependency, redactionManifestState, redactionImpactState, offlineState, boundaryState)
		if exportState != Point12ValCExportStateBlocked {
			t.Fatalf("expected direct ValC export policy drift to block, got state=%s reasons=%v", exportState, exportReasons)
		}
		assertPoint12ValEReason(t, exportReasons, "audit_export_dependency_binding_mismatch")

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected ValC export policy hash drift to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_valc_export_offline_redaction_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("coordinated ValC embedded request and export drift cannot self-consistently bypass ValE", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		driftHash := "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		model.Dependency.ValC.Dependency.ValBReplayRequest.PolicyHash = driftHash
		model.Dependency.ValC.ExportBundle.PolicyHash = driftHash
		model.Dependency.ValC.OfflineBundle.PolicyHash = driftHash

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected coordinated ValC embedded request/export drift to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_valc_replay_request_binding_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("coordinated top-level ValB and ValC policy drift cannot bypass ValA binding", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		driftHash := "sha256:8888888888888888888888888888888888888888888888888888888888888888"
		model.Dependency.ValB.ReplayRequest.PolicyHash = driftHash
		model.Dependency.ValC.Dependency.ValBReplayRequest.PolicyHash = driftHash
		model.Dependency.ValC.ExportBundle.PolicyHash = driftHash
		model.Dependency.ValC.OfflineBundle.PolicyHash = driftHash

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected coordinated top-level ValB/ValC policy drift to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_valb_request_manifest_binding_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("coordinated ValC and ValD detached signature drift cannot bypass ValE", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.OfflineBundle.DetachedSignatureRef = "detached_signature_point12_vala_999"
		model.Dependency.ValD.Dependency.ValCOfflineBundle.DetachedSignatureRef = "detached_signature_point12_vala_999"

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected coordinated detached signature drift to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_valc_offline_binding_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("coordinated ValC and ValD retention class drift cannot bypass ValE", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		driftRetention := "retention_class_point12_downgraded"
		model.Dependency.ValC.ExportBundle.RetentionClassRef = driftRetention
		model.Dependency.ValC.OfflineBundle.RetentionClassRef = driftRetention
		model.Dependency.ValC.RedactionManifest.RetentionClassRef = driftRetention
		model.Dependency.ValD.Dependency.ValCAuditExportBundle.RetentionClassRef = driftRetention
		model.Dependency.ValD.Dependency.ValCOfflineBundle.RetentionClassRef = driftRetention
		model.Dependency.ValD.Dependency.ValCRedactionManifest.RetentionClassRef = driftRetention

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected coordinated retention class drift to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_valc_export_binding_invalid")

		retentionState, retentionReasons := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency)
		if retentionState != Point12ValEStateBlocked {
			t.Fatalf("expected coordinated retention class drift to block retention provenance, got state=%s reasons=%v", retentionState, retentionReasons)
		}
		assertPoint12ValEReason(t, retentionReasons, "retention_provenance_binding_mismatch")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
		assertBlockedAggregate(t, computed, computed.RetentionProvenanceState, "retention_provenance:"+Point12ValEStateBlocked)
	})

	t.Run("vald embedded valb request profile drift blocks final binding mutation gate", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.Dependency.ValBReplayRequest.ProfileContext.CurrentProfileHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		model.Dependency.ValD.Dependency.ValBReplayRequest.ProfileContext.ProfileMatchOriginal = false
		model.Dependency.ValD.Dependency.ValBReplayRequest.ProfileContext.ProfileBindingStatus = Point12Val0ProfileBindingStatusMismatch
		model.Dependency.ValD.Dependency.ValBReplayRequest.ProfileContext.ProfileMismatchReason = "profile_hash_mismatch"

		if got := EvaluatePoint12ValDDependencyState(model.Dependency.ValD.Dependency); got != Point12ValDDependencyStateBlocked {
			t.Fatalf("expected direct ValD dependency state to block embedded request profile drift, got %s", got)
		}

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected ValD embedded ValB replay request profile drift to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_vald_dependency_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("vald embedded valb result profile drift blocks final binding mutation gate", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.Dependency.ValBReplayResult.ProfileContext.CurrentProfileHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		model.Dependency.ValD.Dependency.ValBReplayResult.ProfileContext.ProfileMatchOriginal = false
		model.Dependency.ValD.Dependency.ValBReplayResult.ProfileContext.ProfileBindingStatus = Point12Val0ProfileBindingStatusMismatch
		model.Dependency.ValD.Dependency.ValBReplayResult.ProfileContext.ProfileMismatchReason = "profile_hash_mismatch"

		if got := EvaluatePoint12ValDDependencyState(model.Dependency.ValD.Dependency); got != Point12ValDDependencyStateBlocked {
			t.Fatalf("expected direct ValD dependency state to block embedded result profile drift, got %s", got)
		}

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected ValD embedded ValB replay result profile drift to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_vald_dependency_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("coordinated ValD embedded ValC export drift cannot self-consistently bypass ValE", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		driftHash := "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		model.Dependency.ValD.Dependency.ValCAuditExportBundle.PolicyHash = driftHash
		model.Dependency.ValD.ProofChain.PolicyHash = driftHash
		model.Dependency.ValD.ProofChain.ProjectionHash = point12ValDComputedProjectionHash(model.Dependency.ValD.ProofChain)

		if got := EvaluatePoint12ValDProofChainProjectionState(model.Dependency.ValD.ProofChain, model.Dependency.ValD.Dependency); got != Point12ValDProofChainStateActive {
			t.Fatalf("expected coordinated ValD local proof chain to remain internally active before ValE cross-snapshot check, got %s", got)
		}

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected coordinated ValD embedded export drift to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_vald_dependency_cross_snapshot_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("self consistent substituted ValA profile context blocks final binding mutation gate", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValA.Manifest.ProfileContext.OriginalProfileID = "profile_point12_replay_substituted_002"
		model.Dependency.ValA.Manifest.ProfileContext.CurrentProfileID = "profile_point12_replay_substituted_002"
		model.Dependency.ValA.Manifest.ProfileContext.OriginalProfileVersion = "profile_version_point12_replay_v2"
		model.Dependency.ValA.Manifest.ProfileContext.CurrentProfileVersion = "profile_version_point12_replay_v2"
		model.Dependency.ValA.Manifest.ProfileContext.OriginalProfileHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		model.Dependency.ValA.Manifest.ProfileContext.CurrentProfileHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		model.Dependency.ValA.Manifest.ProfileContext.ProfileApprovalRef = "profile_approval_point12_replay_002"
		model.Dependency.ValA.Manifest.ProfileContext.ProfileSignatureRef = "profile_signature_point12_replay_002"
		model.Dependency.ValA.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Dependency.ValA.Manifest)
		model.Dependency.ValA.Manifest.SignatureBoundManifestPayloadHash = model.Dependency.ValA.Manifest.ManifestPayloadHash

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected substituted ValA profile context to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_profile_context_binding_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("coordinated ValA profile substitution still binds back to Val0 source", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		substituted := model.Dependency.ValA.Manifest.ProfileContext
		substituted.OriginalProfileID = "profile_point12_replay_substituted_002"
		substituted.CurrentProfileID = "profile_point12_replay_substituted_002"
		substituted.OriginalProfileVersion = "profile_version_point12_replay_v2"
		substituted.CurrentProfileVersion = "profile_version_point12_replay_v2"
		substituted.OriginalProfileHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		substituted.CurrentProfileHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		substituted.ProfileApprovalRef = "profile_approval_point12_replay_002"
		substituted.ProfileSignatureRef = "profile_signature_point12_replay_002"
		model.Dependency.ValA.Dependency.Val0Manifest.ProfileContext = substituted
		model.Dependency.ValA.Manifest.ProfileContext = substituted
		model.Dependency.ValA.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Dependency.ValA.Manifest)
		model.Dependency.ValA.Manifest.SignatureBoundManifestPayloadHash = model.Dependency.ValA.Manifest.ManifestPayloadHash

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected coordinated profile substitution to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_val0_profile_context_binding_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("coordinated Val0 and ValA profile substitution still requires original source profile", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		substituted := model.Dependency.Val0.Manifest.ProfileContext
		substituted.OriginalProfileID = "profile_point12_replay_substituted_002"
		substituted.CurrentProfileID = "profile_point12_replay_substituted_002"
		substituted.OriginalProfileVersion = "profile_version_point12_replay_v2"
		substituted.CurrentProfileVersion = "profile_version_point12_replay_v2"
		substituted.OriginalProfileHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		substituted.CurrentProfileHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		substituted.ProfileApprovalRef = "profile_approval_point12_replay_002"
		substituted.ProfileSignatureRef = "profile_signature_point12_replay_002"
		model.Dependency.Val0.Manifest.ProfileContext = substituted
		model.Dependency.Val0.ReplayAssessment.ProfileContext = substituted
		model.Dependency.ValA.Dependency.Val0Manifest.ProfileContext = substituted
		model.Dependency.ValA.Manifest.ProfileContext = substituted
		model.Dependency.ValA.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Dependency.ValA.Manifest)
		model.Dependency.ValA.Manifest.SignatureBoundManifestPayloadHash = model.Dependency.ValA.Manifest.ManifestPayloadHash

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected coordinated Val0 and ValA profile substitution to block binding mutation, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_val0_profile_context_binding_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("coordinated tenant retag across Val0 ValA and ValB still requires immutable source profile", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		retagged := point12Val0DefaultProfileContext("tenant_scope_point12_beta")
		model.Dependency.Val0.Manifest.TenantScope = "tenant_scope_point12_beta"
		model.Dependency.Val0.Manifest.ProfileContext = retagged
		model.Dependency.Val0.ReplayAssessment.ProfileContext = retagged
		model.Dependency.ValA.Dependency.Val0Manifest.TenantScope = "tenant_scope_point12_beta"
		model.Dependency.ValA.Dependency.Val0Manifest.ProfileContext = retagged
		model.Dependency.ValA.Manifest.TenantScope = "tenant_scope_point12_beta"
		model.Dependency.ValA.Manifest.ProfileContext = retagged
		model.Dependency.ValA.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Dependency.ValA.Manifest)
		model.Dependency.ValA.Manifest.SignatureBoundManifestPayloadHash = model.Dependency.ValA.Manifest.ManifestPayloadHash
		model.Dependency.ValB.Dependency.ValAManifest = model.Dependency.ValA.Manifest

		bindingState, bindingReasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency)
		if bindingState != Point12ValEStateBlocked {
			t.Fatalf("expected coordinated tenant retag to block immutable source profile, got state=%s reasons=%v", bindingState, bindingReasons)
		}
		assertPoint12ValEReason(t, bindingReasons, "binding_mutation_val0_profile_context_binding_invalid")

		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("padded no overclaim review id blocks exact aggregate state", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.NoOverclaimReview.ReviewID = " " + model.NoOverclaimReview.ReviewID + " "
		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.NoOverclaimState, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("padded clean room review id blocks exact aggregate state", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.CleanRoomIPReview.ReviewID = " " + model.CleanRoomIPReview.ReviewID + " "
		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.CleanRoomIPState, "clean_room_ip:"+Point12ValEStateBlocked)
	})

	t.Run("padded clean room license review ref blocks exact aggregate state", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.CleanRoomIPReview.LicenseReviewRefs[0] = " " + model.CleanRoomIPReview.LicenseReviewRefs[0] + " "
		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.CleanRoomIPState, "clean_room_ip:"+Point12ValEStateBlocked)
	})

	t.Run("tab newline clean room ip review ref blocks exact aggregate state", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.CleanRoomIPReview.IPReviewRefs[0] = "\t" + model.CleanRoomIPReview.IPReviewRefs[0] + "\n"
		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.CleanRoomIPState, "clean_room_ip:"+Point12ValEStateBlocked)
	})

	t.Run("padded clean room ai review package ref blocks exact aggregate state", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.CleanRoomIPReview.AIReviewPackageRefs[0] = " " + model.CleanRoomIPReview.AIReviewPackageRefs[0] + " "
		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.CleanRoomIPState, "clean_room_ip:"+Point12ValEStateBlocked)
	})

	t.Run("padded retention review id blocks exact aggregate state", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.RetentionProvenanceReview.ReviewID = " " + model.RetentionProvenanceReview.ReviewID + " "
		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.RetentionProvenanceState, "retention_provenance:"+Point12ValEStateBlocked)
	})

	t.Run("retagged vald query generated at blocks exact aggregate state", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.Query.GeneratedAt = "\t" + model.Dependency.ValD.Query.GeneratedAt + "\n"
		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})

	t.Run("padded vald binding matrix generated at blocks exact aggregate state", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.BindingMatrix.GeneratedAt = " " + model.Dependency.ValD.BindingMatrix.GeneratedAt + " "
		computed := ComputePoint12ValEFoundation(model)
		assertBlockedAggregate(t, computed, computed.BindingMutationState, "binding_mutation:"+Point12ValEStateBlocked)
	})
}

func TestPoint12ValEDependencyDerivedBoundaryRecomputedBeforePass(t *testing.T) {
	t.Run("stale clean projection review cannot hide unsafe dependency projection", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.ExportBundle.AdvisoryOnly = false
		model.ProjectionBoundary = point12ValEProjectionBoundaryModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected stale clean projection review to block from dependency recomputation, got %#v", computed)
		}
		if computed.ProjectionBoundaryState != Point12ValEStateBlocked {
			t.Fatalf("expected projection boundary state blocked, got %#v", computed.ProjectionBoundary)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "projection_boundary:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide unsafe dependency wording", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.ExportBundle.CustomerVisibleSummary = "deployment approved"
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected stale clean no-overclaim review to block from dependency recomputation, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide split unsafe dependency wording", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.ExportBundle.CustomerVisibleSummary = "deployment"
		model.Dependency.ValC.OfflineBundle.CustomerVisibleExplanation = "approved"
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected split dependency no-overclaim wording to block from recomputation, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide cross-surface dependency split wording", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.Explanation.CustomerVisibleStatement = "deployment"
		model.Dependency.ValC.ExportBundle.CustomerVisibleSummary = "approved"
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected cross-surface split dependency no-overclaim wording to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide inherited point10 readiness overclaim wording", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.OfflineBundle.CustomerVisibleExplanation = "marketplace production ready"
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited readiness overclaim dependency wording to block, got %#v", computed)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide inherited overclaim in export output claims", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.ExportBundle.ExportOutputClaims = []string{"marketplace production ready"}
		model.Dependency.ValD.Dependency.ValCAuditExportBundle = model.Dependency.ValC.ExportBundle
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited export output claim overclaim to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide inherited overclaim in offline output claims", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValC.OfflineBundle.OfflineOutputClaims = []string{"ha guaranteed"}
		model.Dependency.ValD.Dependency.ValCOfflineBundle = model.Dependency.ValC.OfflineBundle
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited offline output claim overclaim to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	valCNoOverclaimSiblingCases := []struct {
		name   string
		mutate func(*Point12ValEFoundation)
	}{
		{
			name: "export limitations",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValC.ExportBundle.Limitations = []string{"deployment approved"}
			},
		},
		{
			name: "redaction surviving claims",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValC.RedactionManifest.SurvivingClaimsAfterRedaction = []string{"production approved"}
			},
		},
		{
			name: "redaction customer visible claims",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValC.RedactionManifest.CustomerVisibleClaimsAfterRedaction = []string{"public badge"}
			},
		},
		{
			name: "redaction exported claims",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValC.RedactionManifest.ExportedClaimsAfterRedaction = []string{"global truth"}
			},
		},
		{
			name: "redaction replay result claims",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValC.RedactionManifest.ReplayResultClaims = []string{"official authority"}
			},
		},
		{
			name: "redaction limitations",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValC.RedactionManifest.Limitations = []string{"compliance guaranteed"}
			},
		},
		{
			name: "redaction impact limitations",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValC.RedactionImpactVerdict.Limitations = []string{"financial guarantee"}
			},
		},
		{
			name: "offline limitations",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValC.OfflineBundle.Limitations = []string{"deployment approved"}
			},
		},
	}
	for _, tc := range valCNoOverclaimSiblingCases {
		t.Run("stale clean no-overclaim review cannot hide inherited overclaim in ValC "+tc.name, func(t *testing.T) {
			model := rawPoint12ValEFoundationModel()
			tc.mutate(&model)
			model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
			model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

			computed := ComputePoint12ValEFoundation(model)
			assertPoint12ValENoPass(t, computed)
			if computed.CurrentState != Point12ValEStateBlocked {
				t.Fatalf("expected inherited ValC %s overclaim to block, got %#v", tc.name, computed)
			}
			if computed.NoOverclaimState != Point12ValEStateBlocked {
				t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
			}
			assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
		})
	}

	t.Run("stale clean no-overclaim review cannot hide inherited overclaim in replay result claims", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValB.ReplayResult.ReplayOutputClaims = []string{"production sla approved"}
		model.Dependency.ValC.Dependency.ValBReplayResult = model.Dependency.ValB.ReplayResult
		model.Dependency.ValD.Dependency.ValBReplayResult = model.Dependency.ValB.ReplayResult
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited replay result claim overclaim to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide inherited overclaim in vald explanation lists", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.Explanation.MismatchExplanations = []string{"marketplace production ready"}
		model.Dependency.ValD.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Dependency.ValD.Explanation)
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited ValD explanation list overclaim to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide inherited overclaim in vald expected refs", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.Explanation.ExpectedRefs = []string{"marketplace production ready"}
		model.Dependency.ValD.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Dependency.ValD.Explanation)
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited ValD expected refs overclaim to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide inherited overclaim in vald actual refs", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.Explanation.ActualRefs = []string{"ha guaranteed"}
		model.Dependency.ValD.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Dependency.ValD.Explanation)
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited ValD actual refs overclaim to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide inherited overclaim in vald expected versions", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.Explanation.ExpectedVersions = []string{"marketplace_certified"}
		model.Dependency.ValD.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Dependency.ValD.Explanation)
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited ValD expected versions overclaim to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide inherited overclaim in vald actual versions", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.Explanation.ActualVersions = []string{"ha_guaranteed"}
		model.Dependency.ValD.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Dependency.ValD.Explanation)
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited ValD actual versions overclaim to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	valDExplanationSiblingCases := []struct {
		name   string
		mutate func(*Point12ValEFoundation)
	}{
		{
			name: "decision context summary",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValD.Explanation.DecisionContextSummary = "production approved"
			},
		},
		{
			name: "why decision summary",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValD.Explanation.WhyDecisionSummary = "deployment approved"
			},
		},
		{
			name: "why changed summary",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValD.Explanation.WhyChangedSummary = "public badge"
			},
		},
		{
			name: "missing evidence explanations",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValD.Explanation.MissingEvidenceExplanations = []string{"global truth"}
			},
		},
		{
			name: "redaction limitations",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValD.Explanation.RedactionLimitations = []string{"compliance guaranteed"}
			},
		},
		{
			name: "limitations",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValD.Explanation.Limitations = []string{"financial guarantee"}
			},
		},
	}
	for _, tc := range valDExplanationSiblingCases {
		t.Run("stale clean no-overclaim review cannot hide inherited overclaim in vald explanation "+tc.name, func(t *testing.T) {
			model := rawPoint12ValEFoundationModel()
			tc.mutate(&model)
			model.Dependency.ValD.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Dependency.ValD.Explanation)
			model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
			model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

			computed := ComputePoint12ValEFoundation(model)
			assertPoint12ValENoPass(t, computed)
			if computed.CurrentState != Point12ValEStateBlocked {
				t.Fatalf("expected inherited ValD explanation %s overclaim to block, got %#v", tc.name, computed)
			}
			if computed.NoOverclaimState != Point12ValEStateBlocked {
				t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
			}
			assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
		})
	}

	t.Run("stale clean no-overclaim review cannot hide inherited overclaim in vald support limitations", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.SupportProfile.Limitations = []string{"ha_guaranteed"}
		model.Dependency.ValD.SupportProfile.ProfileHash = point12ValDComputedSupportProfileHash(model.Dependency.ValD.SupportProfile)
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited ValD support limitation overclaim to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	t.Run("stale clean no-overclaim review cannot hide inherited overclaim in vald support statement", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.SupportProfile.SupportStatement = "deployment approved"
		model.Dependency.ValD.SupportProfile.ProfileHash = point12ValDComputedSupportProfileHash(model.Dependency.ValD.SupportProfile)
		model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected inherited ValD support statement overclaim to block, got %#v", computed)
		}
		if computed.NoOverclaimState != Point12ValEStateBlocked {
			t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
		}
		assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
	})

	valDSupportProfileListCases := []struct {
		name   string
		mutate func(*Point12ValEFoundation)
	}{
		{
			name: "evidence support categories",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValD.SupportProfile.EvidenceSupportCategories = []string{"production approved"}
			},
		},
		{
			name: "risk context metadata",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValD.SupportProfile.RiskContextMetadata = []string{"compliance guaranteed"}
			},
		},
		{
			name: "allowed wording refs",
			mutate: func(model *Point12ValEFoundation) {
				model.Dependency.ValD.SupportProfile.AllowedWordingRefs = []string{"production approved"}
			},
		},
	}
	for _, tc := range valDSupportProfileListCases {
		t.Run("stale clean no-overclaim review cannot hide inherited overclaim in vald support profile "+tc.name, func(t *testing.T) {
			model := rawPoint12ValEFoundationModel()
			tc.mutate(&model)
			model.Dependency.ValD.SupportProfile.ProfileHash = point12ValDComputedSupportProfileHash(model.Dependency.ValD.SupportProfile)
			model.NoOverclaimReview = point12ValENoOverclaimReviewModel(activePoint12ValEFoundation().Dependency)
			model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

			computed := ComputePoint12ValEFoundation(model)
			assertPoint12ValENoPass(t, computed)
			if computed.CurrentState != Point12ValEStateBlocked {
				t.Fatalf("expected inherited ValD support profile %s overclaim to block, got %#v", tc.name, computed)
			}
			if computed.NoOverclaimState != Point12ValEStateBlocked {
				t.Fatalf("expected no-overclaim state blocked, got %#v", computed.NoOverclaimReview)
			}
			assertPoint12ValEReason(t, computed.BlockingReasons, "no_overclaim:"+Point12ValEStateBlocked)
		})
	}

	t.Run("vald explanation list point12 pass token cannot reach final closure", func(t *testing.T) {
		model := rawPoint12ValEFoundationModel()
		model.Dependency.ValD.Explanation.MismatchExplanations = []string{"point_12_pass"}
		model.Dependency.ValD.Explanation.ExplanationHash = point12ValDComputedExplanationHash(model.Dependency.ValD.Explanation)
		model.PassClosureManifest = point12ValEPassClosureManifestModel(model.Dependency)

		state, reasons := point12ValEDependencyStateAndReasons(model.Dependency)
		if state != Point12ValEStateBlocked {
			t.Fatalf("expected dependency state blocked for ValD explanation pass token, got state=%s reasons=%#v", state, reasons)
		}
		assertPoint12ValEReason(t, reasons, "dependency_contains_point12_pass_input")

		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateBlocked {
			t.Fatalf("expected ValD explanation pass token to block final closure, got %#v", computed)
		}
	})
}

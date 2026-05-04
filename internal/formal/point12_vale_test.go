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

func uncachedActivePoint12ValEFoundation() Point12ValEFoundation {
	return ComputePoint12ValEFoundation(Point12ValEFoundationModel())
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
	if model.CurrentState == Point12ValEStatePassConfirmed {
		t.Fatalf("expected state other than pass_confirmed, got %#v", model)
	}
}

func TestPoint12ValEFoundationFixtureIsolation(t *testing.T) {
	t.Run("raw production path still computes", func(t *testing.T) {
		model := uncachedActivePoint12ValEFoundation()
		if model.CurrentState != Point12ValEStatePassConfirmed {
			t.Fatalf("expected raw production path to compute pass-confirmed baseline, got %#v", model)
		}
		if !model.Point12PassAllowed || model.Point12PassToken != point12ValEPoint12PassToken {
			t.Fatalf("expected raw production path to emit point_12_pass on final happy path, got %#v", model)
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
}

func TestPoint12ValEBindingMutationClosure(t *testing.T) {
	t.Run("valid binding mutation closure active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got, reasons := point12ValEBindingMutationStateAndReasons(model.BindingMutationClosure, model.Dependency); got != Point12ValEStateActive {
			t.Fatalf("expected active binding mutation state, got %s reasons=%v model=%#v", got, reasons, model.BindingMutationClosure)
		}
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

	t.Run("forbidden customer wording blocks", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		model.NoOverclaimReview.ObservedCustomerTexts = []string{"regulator-approved deployment approved"}
		if got, _ := point12ValENoOverclaimStateAndReasons(model.NoOverclaimReview); got != Point12ValEStateBlocked {
			t.Fatalf("expected blocked no-overclaim state, got %#v", model.NoOverclaimReview)
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
}

func TestPoint12ValERetentionProvenanceState(t *testing.T) {
	t.Run("valid retention provenance active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got, _ := point12ValERetentionProvenanceStateAndReasons(model.RetentionProvenanceReview, model.Dependency); got != Point12ValEStateActive {
			t.Fatalf("expected active retention/provenance state, got %#v", model.RetentionProvenanceReview)
		}
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
}

func TestPoint12ValEPassClosureManifestState(t *testing.T) {
	t.Run("valid complete pass closure manifest active", func(t *testing.T) {
		model := activePoint12ValEFoundation()
		if got, reasons := point12ValEPassClosureManifestStateAndReasons(model.PassClosureManifest, model, true); got != Point12ValEStateActive {
			t.Fatalf("expected active pass closure manifest state, got %s reasons=%v model=%#v", got, reasons, model.PassClosureManifest)
		}
	})

	cases := []struct {
		name     string
		expected bool
		mutate   func(*Point12ValEPassClosureManifest)
		want     string
	}{
		{name: "missing closure manifest id blocks", expected: true, mutate: func(model *Point12ValEPassClosureManifest) { model.ClosureManifestID = "" }, want: Point12ValEStateBlocked},
		{name: "wrong point id blocks", expected: true, mutate: func(model *Point12ValEPassClosureManifest) { model.PointID = "point_11" }, want: Point12ValEStateBlocked},
		{name: "missing dependency refs block", expected: true, mutate: func(model *Point12ValEPassClosureManifest) { model.ValDSnapshotRef = "" }, want: Point12ValEStateBlocked},
		{name: "commit sha present before commit blocks", expected: true, mutate: func(model *Point12ValEPassClosureManifest) { model.CommitSHAIfAvailable = "abc123" }, want: Point12ValEStateBlocked},
		{name: "point12 token present before final path blocks", expected: false, mutate: func(model *Point12ValEPassClosureManifest) {}, want: Point12ValEStateBlocked},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			foundation := activePoint12ValEFoundation()
			tc.mutate(&foundation.PassClosureManifest)
			if got, _ := point12ValEPassClosureManifestStateAndReasons(foundation.PassClosureManifest, foundation, tc.expected); got != tc.want {
				t.Fatalf("expected %s, got %#v", tc.want, foundation.PassClosureManifest)
			}
		})
	}
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
		model := rawPoint12ValEFoundationModel()
		model.PassClosureManifest.ReviewerResult = point12ValEReviewerResultPass
		model.PassClosureManifest.Point12PassAllowed = false
		model.PassClosureManifest.Point12PassToken = ""
		computed := ComputePoint12ValEFoundation(model)
		assertPoint12ValENoPass(t, computed)
		if computed.CurrentState != Point12ValEStateActive {
			t.Fatalf("expected active aggregate state without final pass token, got %#v", computed)
		}
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
}

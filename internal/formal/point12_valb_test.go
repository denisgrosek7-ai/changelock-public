package formal

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
	"testing"
)

var (
	point12ValBActiveFoundationBaselineJSON []byte
	point12ValBActiveFoundationBaselineOnce sync.Once
)

func mustMarshalPoint12ValBFoundation(model Point12ValBFoundation) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint12ValBFoundation(payload []byte) Point12ValBFoundation {
	var clone Point12ValBFoundation
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

func activePoint12ValBDependencySnapshot() Point12ValBDependencySnapshot {
	valA := activePoint12ValAFoundation()
	return SnapshotPoint12ValBDependencyFromComputedValA(valA, point12ValBDependencyReviewContextModel())
}

func syncPoint12ValBFoundationToDependency(model *Point12ValBFoundation) {
	model.ReplayCommand.ProofPackID = model.Dependency.ValAManifest.ProofPackID
	model.ReplayCommand.ManifestID = model.Dependency.ValAManifest.ManifestID
	model.ReplayCommand.TenantScope = model.Dependency.ValAManifest.TenantScope
	model.ReplayCommand.ArtifactRef = model.Dependency.ValAManifest.ArtifactRef
	model.ReplayCommand.CompatibilityProfileRef = model.Dependency.ValAManifest.CompatibilityProfileRef

	model.ReplayRequest.ProofPackID = model.Dependency.ValAManifest.ProofPackID
	model.ReplayRequest.ManifestID = model.Dependency.ValAManifest.ManifestID
	model.ReplayRequest.DecisionID = model.Dependency.ValAManifest.DecisionID
	model.ReplayRequest.TenantScope = model.Dependency.ValAManifest.TenantScope
	model.ReplayRequest.ArtifactRef = model.Dependency.ValAManifest.ArtifactRef
	model.ReplayRequest.ArtifactHash = model.Dependency.ValAManifest.ArtifactHash
	model.ReplayRequest.EvidenceRefs = append([]string{}, model.Dependency.ValAManifest.EvidenceRefs...)
	model.ReplayRequest.EvidenceHashRefs = append([]string{}, model.Dependency.ValAManifest.EvidenceHashRefs...)
	model.ReplayRequest.PolicyRef = model.Dependency.ValAManifest.PolicyRef
	model.ReplayRequest.PolicyVersion = model.Dependency.ValAManifest.PolicyVersion
	model.ReplayRequest.PolicyHash = model.Dependency.ValAManifest.PolicyHash
	model.ReplayRequest.EngineVersion = model.Dependency.ValAManifest.EngineVersion
	model.ReplayRequest.EngineHash = model.Dependency.ValAManifest.EngineHash
	model.ReplayRequest.SchemaVersion = model.Dependency.ValAManifest.SchemaVersion
	model.ReplayRequest.SchemaHash = model.Dependency.ValAManifest.SchemaHash
	model.ReplayRequest.ClaimRefs = append([]string{}, model.Dependency.ValAManifest.ClaimRefs...)
	model.ReplayRequest.GovernanceEventRefs = append([]string{}, model.Dependency.ValAManifest.GovernanceEventRefs...)
	model.ReplayRequest.ManifestPayloadHash = model.Dependency.ValAManifest.ManifestPayloadHash
	model.ReplayRequest.CompatibilityProfileRef = model.Dependency.ValAManifest.CompatibilityProfileRef
	model.ReplayRequest.RedactionManifestRef = model.Dependency.ValAManifest.RedactionManifestRef
	model.ReplayRequest.SourceManifestIntegrityState = model.Dependency.ValAManifestIntegrityState

	model.ReplayResult.ReplayRequestID = model.ReplayRequest.ReplayRequestID
	model.ReplayResult.ProofPackID = model.ReplayRequest.ProofPackID
	model.ReplayResult.ManifestID = model.ReplayRequest.ManifestID
	model.ReplayResult.ReplayMode = model.ReplayRequest.ReplayMode
	model.ReplayResult.OriginalDecisionState = model.ReplayRequest.OriginalDecisionState
	if strings.TrimSpace(model.ReplayResult.ReplayedDecisionState) == "" {
		model.ReplayResult.ReplayedDecisionState = model.ReplayRequest.OriginalDecisionState
	}
	model.ReplayResult.EvaluatedPolicyVersion = model.ReplayRequest.PolicyVersion
	model.ReplayResult.EvaluatedEngineVersion = model.ReplayRequest.EngineVersion
	model.ReplayResult.EvaluatedSchemaVersion = model.ReplayRequest.SchemaVersion
}

func uncachedActivePoint12ValBFoundation() Point12ValBFoundation {
	model := Point12ValBFoundationModel()
	model.Dependency = activePoint12ValBDependencySnapshot()
	syncPoint12ValBFoundationToDependency(&model)
	return ComputePoint12ValBFoundation(model)
}

func activePoint12ValBFoundation() Point12ValBFoundation {
	point12ValBActiveFoundationBaselineOnce.Do(func() {
		point12ValBActiveFoundationBaselineJSON = mustMarshalPoint12ValBFoundation(uncachedActivePoint12ValBFoundation())
	})
	return clonePoint12ValBFoundation(point12ValBActiveFoundationBaselineJSON)
}

func activePoint12ValBFoundationFromValA(valA Point12ValAFoundation) Point12ValBFoundation {
	model := Point12ValBFoundationModel()
	model.Dependency = SnapshotPoint12ValBDependencyFromComputedValA(valA, point12ValBDependencyReviewContextModel())
	syncPoint12ValBFoundationToDependency(&model)
	return ComputePoint12ValBFoundation(model)
}

func readPoint12ValBSource(t *testing.T) string {
	t.Helper()
	for _, path := range []string{"point12_valb.go", "internal/formal/point12_valb.go"} {
		body, err := os.ReadFile(path)
		if err == nil {
			return string(body)
		}
	}
	t.Fatal("failed to read point12_valb.go source")
	return ""
}

func point12ValBPolicyMismatch(decisive bool) Point12ValBReplayMismatch {
	return Point12ValBReplayMismatch{
		MismatchID:      "replay_mismatch_policy_001",
		MismatchType:    point12ValBMismatchTypePolicyMismatch,
		ExpectedRef:     "policy_point12_original_001",
		ActualRef:       "policy_point12_current_001",
		ExpectedHash:    "sha256:1111111111111111111111111111111111111111111111111111111111111111",
		ActualHash:      "sha256:2222222222222222222222222222222222222222222222222222222222222222",
		ExpectedVersion: "policy_version_point12_original_001",
		ActualVersion:   "policy_version_point12_current_001",
		AffectedSurface: "policy_context",
		Decisive:        decisive,
		DriftReason:     point12ValBDriftDueToPolicy,
		Explanation:     "policy context changed",
		BlocksReplay:    decisive,
	}
}

func point12ValBMismatchForType(mismatchType string, drift string) Point12ValBReplayMismatch {
	return Point12ValBReplayMismatch{
		MismatchID:      "replay_mismatch_generic_001",
		MismatchType:    mismatchType,
		ExpectedRef:     "expected_ref_001",
		ActualRef:       "actual_ref_001",
		ExpectedHash:    "sha256:1111111111111111111111111111111111111111111111111111111111111111",
		ActualHash:      "sha256:2222222222222222222222222222222222222222222222222222222222222222",
		ExpectedVersion: "expected_version_001",
		ActualVersion:   "actual_version_002",
		AffectedSurface: "replay_surface",
		Decisive:        true,
		DriftReason:     drift,
		Explanation:     "mismatch is explicitly explained",
		BlocksReplay:    strings.TrimSpace(mismatchType) == point12ValBMismatchTypeTamperDetected || strings.TrimSpace(mismatchType) == point12ValBMismatchTypeMissingEvidence,
	}
}

func TestPoint12ValBDependencyState(t *testing.T) {
	t.Run("valid computed vala foundation allows replay semantics", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		if model.DependencyState != Point12ValBDependencyStateActive {
			t.Fatalf("expected active dependency state, got %#v", model)
		}
		if model.CurrentState != Point12ValBStateActive {
			t.Fatalf("expected active valb foundation state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point12ValBDependencySnapshot)
		want   string
	}{
		{name: "missing vala dependency blocks", mutate: func(model *Point12ValBDependencySnapshot) { *model = Point12ValBDependencySnapshot{} }, want: Point12ValBDependencyStateBlocked},
		{name: "fallback regenerated vala snapshot blocks", mutate: func(model *Point12ValBDependencySnapshot) { model.SnapshotFromComputedOutput = false }, want: Point12ValBDependencyStateBlocked},
		{name: "vala manifest integrity tampered blocks", mutate: func(model *Point12ValBDependencySnapshot) {
			model.ValAManifestIntegrityState = Point12ValAManifestIntegrityStateTampered
		}, want: Point12ValBDependencyStateBlocked},
		{name: "vala manifest integrity unsupported blocks", mutate: func(model *Point12ValBDependencySnapshot) {
			model.ValAManifestIntegrityState = Point12ValAManifestIntegrityStateUnsupported
		}, want: Point12ValBDependencyStateBlocked},
		{name: "vala blocked dependency blocks", mutate: func(model *Point12ValBDependencySnapshot) {
			model.ValADependencyState = Point12ValADependencyStateBlocked
		}, want: Point12ValBDependencyStateBlocked},
		{name: "vala review required propagates review required", mutate: func(model *Point12ValBDependencySnapshot) {
			model.ValADependencyState = Point12ValADependencyStateReviewRequired
		}, want: Point12ValBDependencyStateReviewRequired},
		{name: "premature point12 pass in dependency blocks", mutate: func(model *Point12ValBDependencySnapshot) { model.ValAPrematurePoint12PassSeen = true }, want: Point12ValBDependencyStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValBDependencySnapshot()
			testCase.mutate(&model)
			if got := EvaluatePoint12ValBDependencyState(model); got != testCase.want {
				t.Fatalf("expected dependency state %q, got %#v", testCase.want, model)
			}
		})
	}
}

func TestPoint12ValBReplayCommandState(t *testing.T) {
	t.Run("valid replay proof pack original context command is active", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		if model.ReplayCommandState != Point12ValBReplayCommandStateActive {
			t.Fatalf("expected active replay command state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point12ValBReplayCommandContract)
	}{
		{name: "missing replay mode blocks", mutate: func(model *Point12ValBReplayCommandContract) { model.ReplayMode = "" }},
		{name: "unknown replay mode blocks", mutate: func(model *Point12ValBReplayCommandContract) { model.ReplayMode = "unknown_mode" }},
		{name: "allow external api true blocks", mutate: func(model *Point12ValBReplayCommandContract) { model.AllowExternalAPI = true }},
		{name: "command attempting export portal mutation blocks", mutate: func(model *Point12ValBReplayCommandContract) {
			model.MutatesEvidenceSpine = true
			model.CreatesAuditExportBundle = true
			model.OpensPortalPath = true
		}},
		{name: "command attempting point12 pass blocks", mutate: func(model *Point12ValBReplayCommandContract) { model.RequestsPoint12Pass = true }},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValBFoundation()
			testCase.mutate(&model.ReplayCommand)
			model = ComputePoint12ValBFoundation(model)
			if model.ReplayCommandState != Point12ValBReplayCommandStateBlocked {
				t.Fatalf("expected blocked replay command state, got %#v", model)
			}
		})
	}
}

func TestPoint12ValBReplayRequestState(t *testing.T) {
	t.Run("valid original context request bound to vala manifest is active", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		if model.ReplayRequestState != Point12ValBReplayRequestStateActive {
			t.Fatalf("expected active replay request state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point12ValBFoundation)
	}{
		{name: "missing proof pack id blocks", mutate: func(model *Point12ValBFoundation) { model.ReplayRequest.ProofPackID = "" }},
		{name: "missing manifest id blocks", mutate: func(model *Point12ValBFoundation) { model.ReplayRequest.ManifestID = "" }},
		{name: "tenant mismatch blocks", mutate: func(model *Point12ValBFoundation) { model.ReplayRequest.TenantScope = "tenant_scope_point12_other_001" }},
		{name: "artifact mismatch blocks", mutate: func(model *Point12ValBFoundation) {
			model.ReplayRequest.ArtifactHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		}},
		{name: "malformed refs block", mutate: func(model *Point12ValBFoundation) { model.ReplayRequest.ReplayRequestID = "bad request id" }},
		{name: "canonical looking junk refs block", mutate: func(model *Point12ValBFoundation) { model.ReplayRequest.ManifestID = "manifest_unknown" }},
		{name: "cross tenant evidence blocks", mutate: func(model *Point12ValBFoundation) {
			model.ReplayRequest.EvidenceRefs = []string{"evidence:cross-tenant-pack-001"}
		}},
		{name: "current policy context must be explicit", mutate: func(model *Point12ValBFoundation) {
			model.ReplayCommand.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
			model.ReplayCommand.AllowCurrentPolicy = true
			model.ReplayCommand.RequestedPolicyContextRef = model.ReplayRequest.PolicyRef
			model.ReplayCommand.RequestedEngineContextRef = model.ReplayRequest.EngineVersion
			model.ReplayCommand.RequestedSchemaContextRef = model.ReplayRequest.SchemaVersion
			model.ReplayRequest.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		}},
		{name: "comparison mode requires current context", mutate: func(model *Point12ValBFoundation) {
			model.ReplayCommand.ReplayMode = point12Val0ReplayModeComparisonMode
			model.ReplayCommand.AllowCurrentPolicy = true
			model.ReplayCommand.ExplainMismatch = true
			model.ReplayCommand.RequestedPolicyContextRef = model.ReplayRequest.PolicyRef
			model.ReplayCommand.RequestedEngineContextRef = model.ReplayRequest.EngineVersion
			model.ReplayCommand.RequestedSchemaContextRef = model.ReplayRequest.SchemaVersion
			model.ReplayRequest.ReplayMode = point12Val0ReplayModeComparisonMode
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValBFoundation()
			testCase.mutate(&model)
			model = ComputePoint12ValBFoundation(model)
			if model.ReplayRequestState != Point12ValBReplayRequestStateBlocked {
				t.Fatalf("expected blocked replay request state, got %#v", model)
			}
		})
	}
}

func TestPoint12ValBReplayModes(t *testing.T) {
	t.Run("original context exact replay returns same decision taxonomy only", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		if model.ReplayResultState != Point12ValBReplayResultStateActive {
			t.Fatalf("expected active replay result state, got %#v", model)
		}
		if model.ReplayResult.ReplayResultTaxonomy != Point12Val0ReplayResultSameDecision {
			t.Fatalf("expected same_decision taxonomy, got %#v", model)
		}
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal valb foundation: %v", err)
		}
		if strings.Contains(string(body), "point_12_pass") {
			t.Fatalf("expected same_decision to remain non-pass taxonomy, got %s", body)
		}
	})

	t.Run("original context cannot silently use current policy", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		model.ReplayRequest.CurrentPolicyRef = "policy_point12_current_001"
		model.ReplayRequest.CurrentPolicyVersion = "policy_version_point12_current_001"
		model.ReplayRequest.CurrentPolicyHash = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
		model = ComputePoint12ValBFoundation(model)
		if model.ReplayResultState != Point12ValBReplayResultStateBlocked {
			t.Fatalf("expected original_context with silent current policy to block, got %#v", model)
		}
	})

	t.Run("current policy context must be explicit and may stay same decision", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		model.ReplayCommand.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model.ReplayCommand.AllowCurrentPolicy = true
		model.ReplayCommand.RequestedPolicyContextRef = model.ReplayRequest.PolicyRef
		model.ReplayCommand.RequestedEngineContextRef = model.ReplayRequest.EngineVersion
		model.ReplayCommand.RequestedSchemaContextRef = model.ReplayRequest.SchemaVersion
		model.ReplayRequest.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model.ReplayRequest.CurrentPolicyRef = model.ReplayRequest.PolicyRef
		model.ReplayRequest.CurrentPolicyVersion = model.ReplayRequest.PolicyVersion
		model.ReplayRequest.CurrentPolicyHash = model.ReplayRequest.PolicyHash
		model.ReplayRequest.CurrentEngineVersion = model.ReplayRequest.EngineVersion
		model.ReplayRequest.CurrentEngineHash = model.ReplayRequest.EngineHash
		model.ReplayRequest.CurrentSchemaVersion = model.ReplayRequest.SchemaVersion
		model.ReplayRequest.CurrentSchemaHash = model.ReplayRequest.SchemaHash
		model.ReplayRequest.CurrentEvidenceRefs = append([]string{}, model.ReplayRequest.EvidenceRefs...)
		model.ReplayRequest.CurrentEvidenceHashRefs = append([]string{}, model.ReplayRequest.EvidenceHashRefs...)
		model.ReplayResult.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model = ComputePoint12ValBFoundation(model)
		if model.ReplayRequestState != Point12ValBReplayRequestStateActive || model.ReplayResultState != Point12ValBReplayResultStateActive {
			t.Fatalf("expected explicit current_policy_context replay to stay active, got %#v", model)
		}
	})

	t.Run("current policy change returns different decision with drift explanation", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		model.ReplayCommand.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model.ReplayCommand.AllowCurrentPolicy = true
		model.ReplayCommand.RequestedPolicyContextRef = "policy_point12_current_001"
		model.ReplayCommand.RequestedEngineContextRef = model.ReplayRequest.EngineVersion
		model.ReplayCommand.RequestedSchemaContextRef = model.ReplayRequest.SchemaVersion
		model.ReplayRequest.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model.ReplayRequest.CurrentPolicyRef = "policy_point12_current_001"
		model.ReplayRequest.CurrentPolicyVersion = "policy_version_point12_current_001"
		model.ReplayRequest.CurrentPolicyHash = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
		model.ReplayRequest.CurrentEngineVersion = model.ReplayRequest.EngineVersion
		model.ReplayRequest.CurrentEngineHash = model.ReplayRequest.EngineHash
		model.ReplayRequest.CurrentSchemaVersion = model.ReplayRequest.SchemaVersion
		model.ReplayRequest.CurrentSchemaHash = model.ReplayRequest.SchemaHash
		model.ReplayRequest.CurrentEvidenceRefs = append([]string{}, model.ReplayRequest.EvidenceRefs...)
		model.ReplayRequest.CurrentEvidenceHashRefs = append([]string{}, model.ReplayRequest.EvidenceHashRefs...)
		model.ReplayResult.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultDifferentDecision
		model.ReplayResult.MatchOriginal = false
		model.ReplayResult.ReplayedDecisionState = "decision_state_block"
		model.ReplayResult.DecisionDriftExplanation = "policy update changed the replayed decision"
		model.ReplayResult.DecisionDriftClassification = point12ValBDriftDueToPolicy
		model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBPolicyMismatch(false)}
		model = ComputePoint12ValBFoundation(model)
		if model.ReplayResultState != Point12ValBReplayResultStateActive {
			t.Fatalf("expected explicit current policy drift replay to stay active, got %#v", model)
		}
	})

	t.Run("current policy replay cannot rewrite original decision", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		model.ReplayCommand.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model.ReplayCommand.AllowCurrentPolicy = true
		model.ReplayCommand.RequestedPolicyContextRef = model.ReplayRequest.PolicyRef
		model.ReplayCommand.RequestedEngineContextRef = model.ReplayRequest.EngineVersion
		model.ReplayCommand.RequestedSchemaContextRef = model.ReplayRequest.SchemaVersion
		model.ReplayRequest.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model.ReplayRequest.CurrentPolicyRef = model.ReplayRequest.PolicyRef
		model.ReplayRequest.CurrentPolicyVersion = model.ReplayRequest.PolicyVersion
		model.ReplayRequest.CurrentPolicyHash = model.ReplayRequest.PolicyHash
		model.ReplayRequest.CurrentEngineVersion = model.ReplayRequest.EngineVersion
		model.ReplayRequest.CurrentEngineHash = model.ReplayRequest.EngineHash
		model.ReplayRequest.CurrentSchemaVersion = model.ReplayRequest.SchemaVersion
		model.ReplayRequest.CurrentSchemaHash = model.ReplayRequest.SchemaHash
		model.ReplayRequest.CurrentEvidenceRefs = append([]string{}, model.ReplayRequest.EvidenceRefs...)
		model.ReplayRequest.CurrentEvidenceHashRefs = append([]string{}, model.ReplayRequest.EvidenceHashRefs...)
		model.ReplayResult.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model.ReplayResult.OriginalDecisionState = "decision_state_changed"
		model = ComputePoint12ValBFoundation(model)
		if model.ReplayResultState != Point12ValBReplayResultStateBlocked {
			t.Fatalf("expected rewritten original decision to block, got %#v", model)
		}
	})

	t.Run("comparison mode requires drift explanation", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		model.ReplayCommand.ReplayMode = point12Val0ReplayModeComparisonMode
		model.ReplayCommand.AllowCurrentPolicy = true
		model.ReplayCommand.RequestedPolicyContextRef = "policy_point12_current_001"
		model.ReplayCommand.RequestedEngineContextRef = model.ReplayRequest.EngineVersion
		model.ReplayCommand.RequestedSchemaContextRef = model.ReplayRequest.SchemaVersion
		model.ReplayRequest.ReplayMode = point12Val0ReplayModeComparisonMode
		model.ReplayRequest.CurrentPolicyRef = "policy_point12_current_001"
		model.ReplayRequest.CurrentPolicyVersion = "policy_version_point12_current_001"
		model.ReplayRequest.CurrentPolicyHash = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
		model.ReplayRequest.CurrentEngineVersion = model.ReplayRequest.EngineVersion
		model.ReplayRequest.CurrentEngineHash = model.ReplayRequest.EngineHash
		model.ReplayRequest.CurrentSchemaVersion = model.ReplayRequest.SchemaVersion
		model.ReplayRequest.CurrentSchemaHash = model.ReplayRequest.SchemaHash
		model.ReplayRequest.CurrentEvidenceRefs = append([]string{}, model.ReplayRequest.EvidenceRefs...)
		model.ReplayRequest.CurrentEvidenceHashRefs = append([]string{}, model.ReplayRequest.EvidenceHashRefs...)
		model.ReplayResult.ReplayMode = point12Val0ReplayModeComparisonMode
		model.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultDifferentDecision
		model.ReplayResult.MatchOriginal = false
		model.ReplayResult.ReplayedDecisionState = "decision_state_block"
		model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBPolicyMismatch(true)}
		model.ReplayResult.DecisionDriftClassification = point12ValBDriftDueToPolicy
		model.ReplayResult.DecisionDriftExplanation = ""
		model = ComputePoint12ValBFoundation(model)
		if model.ReplayResultState != Point12ValBReplayResultStateReviewRequired {
			t.Fatalf("expected comparison mode without drift explanation to require review, got %#v", model)
		}
	})

	t.Run("comparison mode with policy drift explanation stays active", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		model.ReplayCommand.ReplayMode = point12Val0ReplayModeComparisonMode
		model.ReplayCommand.AllowCurrentPolicy = true
		model.ReplayCommand.RequestedPolicyContextRef = "policy_point12_current_001"
		model.ReplayCommand.RequestedEngineContextRef = model.ReplayRequest.EngineVersion
		model.ReplayCommand.RequestedSchemaContextRef = model.ReplayRequest.SchemaVersion
		model.ReplayRequest.ReplayMode = point12Val0ReplayModeComparisonMode
		model.ReplayRequest.CurrentPolicyRef = "policy_point12_current_001"
		model.ReplayRequest.CurrentPolicyVersion = "policy_version_point12_current_001"
		model.ReplayRequest.CurrentPolicyHash = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
		model.ReplayRequest.CurrentEngineVersion = model.ReplayRequest.EngineVersion
		model.ReplayRequest.CurrentEngineHash = model.ReplayRequest.EngineHash
		model.ReplayRequest.CurrentSchemaVersion = model.ReplayRequest.SchemaVersion
		model.ReplayRequest.CurrentSchemaHash = model.ReplayRequest.SchemaHash
		model.ReplayRequest.CurrentEvidenceRefs = append([]string{}, model.ReplayRequest.EvidenceRefs...)
		model.ReplayRequest.CurrentEvidenceHashRefs = append([]string{}, model.ReplayRequest.EvidenceHashRefs...)
		model.ReplayResult.ReplayMode = point12Val0ReplayModeComparisonMode
		model.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultDifferentDecision
		model.ReplayResult.MatchOriginal = false
		model.ReplayResult.ReplayedDecisionState = "decision_state_block"
		model.ReplayResult.DecisionDriftExplanation = "policy drift changed the replay outcome"
		model.ReplayResult.DecisionDriftClassification = point12ValBDriftDueToPolicy
		model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBPolicyMismatch(false)}
		model = ComputePoint12ValBFoundation(model)
		if model.ReplayResultState != Point12ValBReplayResultStateActive {
			t.Fatalf("expected comparison mode with drift explanation to stay active, got %#v", model)
		}
	})
}

func TestPoint12ValBReplayResultTaxonomy(t *testing.T) {
	t.Run("point12 pass emission is blocked", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		model.ReplayResult.PointPassEmitted = true
		model = ComputePoint12ValBFoundation(model)
		if model.ReplayResultState != Point12ValBReplayResultStateBlocked {
			t.Fatalf("expected point pass emission to block, got %#v", model)
		}
	})

	testCases := []struct {
		name     string
		taxonomy string
		mutate   func(*Point12ValBFoundation)
	}{
		{name: "tampered input returns tamper detected", taxonomy: Point12Val0ReplayResultTamperDetected, mutate: func(model *Point12ValBFoundation) {
			model.ReplayResult.TamperDetected = true
			model.ReplayResult.ManifestIntegrityCheckResult = point12ValBCheckResultTampered
			model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeTamperDetected, point12ValBDriftDueToPolicy)}
		}},
		{name: "unsupported version returns unsupported version", taxonomy: Point12Val0ReplayResultUnsupportedVersion, mutate: func(model *Point12ValBFoundation) {
			model.ReplayResult.UnsupportedVersion = true
			model.ReplayResult.CompatibilityCheckResult = point12ValBCheckResultUnsupported
			model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeUnsupportedVersion, point12ValBDriftDueToSchema)}
		}},
		{name: "missing decisive evidence returns insufficient evidence", taxonomy: Point12Val0ReplayResultInsufficientEvidence, mutate: func(model *Point12ValBFoundation) {
			model.ReplayResult.InsufficientEvidence = true
			model.ReplayResult.EvidenceHashCheckResult = point12ValBCheckResultMissing
			model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeMissingEvidence, point12ValBDriftDueToEvidence)}
		}},
		{name: "redacted decisive evidence returns redacted limitations", taxonomy: Point12Val0ReplayResultRedactedLimitations, mutate: func(model *Point12ValBFoundation) {
			model.ReplayResult.RedactionLimitations = true
			model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeRedactionMismatch, point12ValBDriftDueToRedaction)}
		}},
		{name: "policy mismatch taxonomy works", taxonomy: Point12Val0ReplayResultPolicyMismatch, mutate: func(model *Point12ValBFoundation) {
			model.ReplayRequest.PolicyHash = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
			model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBPolicyMismatch(false)}
		}},
		{name: "engine mismatch taxonomy works", taxonomy: Point12Val0ReplayResultEngineMismatch, mutate: func(model *Point12ValBFoundation) {
			model.ReplayRequest.EngineHash = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
			model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeEngineMismatch, point12ValBDriftDueToEngine)}
		}},
		{name: "schema mismatch taxonomy works", taxonomy: Point12Val0ReplayResultSchemaMismatch, mutate: func(model *Point12ValBFoundation) {
			model.ReplayRequest.SchemaHash = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
			model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeSchemaMismatch, point12ValBDriftDueToSchema)}
		}},
		{name: "evidence mismatch taxonomy works", taxonomy: Point12Val0ReplayResultEvidenceMismatch, mutate: func(model *Point12ValBFoundation) {
			model.ReplayRequest.EvidenceHashRefs = []string{"evidence_hash_point12_proof_pack_099"}
			model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeEvidenceMismatch, point12ValBDriftDueToEvidence)}
		}},
		{name: "claim mismatch taxonomy works", taxonomy: Point12Val0ReplayResultClaimMismatch, mutate: func(model *Point12ValBFoundation) {
			model.ReplayRequest.ClaimRefs = []string{"claim_point12_changed_001"}
			model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeClaimMismatch, point12ValBDriftDueToClaim)}
		}},
		{name: "governance mismatch taxonomy works", taxonomy: Point12Val0ReplayResultGovernanceMismatch, mutate: func(model *Point12ValBFoundation) {
			model.ReplayRequest.GovernanceEventRefs = []string{"governance_event_point12_changed_001"}
			model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{point12ValBMismatchForType(point12ValBMismatchTypeGovernanceMismatch, point12ValBDriftDueToGovernance)}
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValBFoundation()
			model.ReplayResult.ReplayResultTaxonomy = testCase.taxonomy
			testCase.mutate(&model)
			model = ComputePoint12ValBFoundation(model)
			if model.ReplayResultState != Point12ValBReplayResultStateActive {
				t.Fatalf("expected active replay result state for %s, got %#v", testCase.name, model)
			}
		})
	}

	t.Run("missing explanation for decisive mismatch requires review", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		model.ReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultPolicyMismatch
		mismatch := point12ValBPolicyMismatch(true)
		mismatch.Explanation = ""
		model.ReplayResult.Mismatches = []Point12ValBReplayMismatch{mismatch}
		model.ReplayRequest.PolicyHash = mismatch.ActualHash
		model = ComputePoint12ValBFoundation(model)
		if model.ReplayResultState != Point12ValBReplayResultStateReviewRequired {
			t.Fatalf("expected decisive mismatch without explanation to require review, got %#v", model)
		}
	})
}

func TestPoint12ValBNoOverclaimAndRegression(t *testing.T) {
	t.Run("forbidden wording in replay output claim blocks", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		model.ReplayResult.ReplayOutputClaims = []string{"production approved"}
		model = ComputePoint12ValBFoundation(model)
		if model.ReplayResultState != Point12ValBReplayResultStateBlocked {
			t.Fatalf("expected forbidden replay output claim to block, got %#v", model)
		}
	})

	t.Run("forbidden wording in customer visible explanation blocks", func(t *testing.T) {
		model := activePoint12ValBFoundation()
		model.ReplayResult.CustomerVisibleExplanation = "compliance guaranteed"
		model = ComputePoint12ValBFoundation(model)
		if model.ReplayResultState != Point12ValBReplayResultStateBlocked {
			t.Fatalf("expected forbidden customer explanation to block, got %#v", model)
		}
	})

	t.Run("val0 disallowed claims ledger bugfix remains allowed through valb dependency", func(t *testing.T) {
		val0 := activePoint12Val0Foundation()
		val0.RedactionBoundary.DisallowedClaimsAfterRedaction = []string{"production approved"}
		val0.RedactionBoundary.RedactionSummary = "internal summary: disallowed production approved claim removed during redaction"
		val0.RedactionBoundary.PartialOrAdvisoryOnly = true
		val0 = ComputePoint12Val0Foundation(val0)
		valA := activePoint12ValAFoundationFromVal0(val0)
		model := activePoint12ValBFoundationFromValA(valA)
		if model.DependencyState != Point12ValBDependencyStateActive {
			t.Fatalf("expected dependency to preserve val0 denylist ledger fix, got %#v", model)
		}
	})

	t.Run("val0 minimum safe claim forbidden wording still blocks through vala and valb", func(t *testing.T) {
		val0 := activePoint12Val0Foundation()
		val0.RedactionBoundary.MinimumSafeClaimAfterRedaction = "production approved"
		val0 = ComputePoint12Val0Foundation(val0)
		valA := activePoint12ValAFoundationFromVal0(val0)
		model := activePoint12ValBFoundationFromValA(valA)
		if model.DependencyState != Point12ValBDependencyStateBlocked {
			t.Fatalf("expected blocked valb dependency when val0 minimum safe claim is forbidden, got %#v", model)
		}
	})

	t.Run("vala manifest tamper behavior preserved", func(t *testing.T) {
		valA := activePoint12ValAFoundation()
		valA.ManifestIntegrityState = Point12ValAManifestIntegrityStateTampered
		model := activePoint12ValBFoundationFromValA(valA)
		if model.DependencyState != Point12ValBDependencyStateBlocked {
			t.Fatalf("expected valb dependency to block on tampered vala manifest, got %#v", model)
		}
	})

	t.Run("detached signature metadata remains metadata only", func(t *testing.T) {
		valA := activePoint12ValAFoundation()
		valA.Manifest.SignatureRef = ""
		valA.Manifest.DetachedSignatureRef = "detached_signature_point12_vala_metadata_001"
		valA.Manifest.SignatureBoundManifestPayloadHash = valA.Manifest.ManifestPayloadHash
		valA = ComputePoint12ValAFoundation(valA)
		model := activePoint12ValBFoundationFromValA(valA)
		if model.DependencyState != Point12ValBDependencyStateActive {
			t.Fatalf("expected detached signature metadata to remain valid for valb dependency, got %#v", model)
		}
	})
}

func TestPoint12ValBSourceBoundaries(t *testing.T) {
	body := readPoint12ValBSource(t)
	for _, forbidden := range []string{
		"http.Get",
		"http.Post",
		"fetch(",
		"kms",
		"hsm",
		"Sign(",
		"GenerateKey",
		"crypto/rsa",
		"crypto/ecdsa",
		"crypto/ed25519",
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("unexpected valb source boundary violation %q", forbidden)
		}
	}
}

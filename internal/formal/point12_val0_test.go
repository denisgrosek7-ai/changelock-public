package formal

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
)

var (
	point12Val0ActiveFoundationBaselineJSON []byte
	point12Val0ActiveFoundationBaselineOnce sync.Once
)

func mustMarshalPoint12Val0Foundation(model Point12Val0Foundation) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint12Val0Foundation(payload []byte) Point12Val0Foundation {
	var clone Point12Val0Foundation
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

func point12Val0ActiveDependencySnapshot() Point12Val0DependencySnapshot {
	valD := activePoint11ValDFoundation()
	return SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD, point12Val0DependencyReviewContextModel())
}

func uncachedActivePoint12Val0Foundation() Point12Val0Foundation {
	model := Point12Val0FoundationModel()
	model.Dependency = point12Val0ActiveDependencySnapshot()
	return ComputePoint12Val0Foundation(model)
}

func activePoint12Val0Foundation() Point12Val0Foundation {
	point12Val0ActiveFoundationBaselineOnce.Do(func() {
		point12Val0ActiveFoundationBaselineJSON = mustMarshalPoint12Val0Foundation(uncachedActivePoint12Val0Foundation())
	})
	return clonePoint12Val0Foundation(point12Val0ActiveFoundationBaselineJSON)
}

func readPoint12Val0Source(t *testing.T) string {
	t.Helper()
	for _, path := range []string{"point12_val0.go", "internal/formal/point12_val0.go"} {
		body, err := os.ReadFile(path)
		if err == nil {
			return string(body)
		}
	}
	t.Fatal("failed to read point12_val0.go source")
	return ""
}

func TestPoint12Val0DependencyState(t *testing.T) {
	t.Run("valid point11 vald final dependency snapshot foundation ready", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		if model.DependencyState != Point12Val0DependencyStateActive {
			t.Fatalf("expected active dependency state, got %#v", model)
		}
		if model.CurrentState != Point12Val0StateActive {
			t.Fatalf("expected active foundation state, got %#v", model)
		}
	})

	t.Run("valid computed snapshot preserves computed provenance", func(t *testing.T) {
		valD := activePoint11ValDFoundation()
		snapshot := SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD, Point12Val0Point11ReviewContext{
			SnapshotFromComputedOutput: true,
		})
		if !snapshot.SnapshotFromComputedOutput {
			t.Fatalf("expected helper to preserve computed provenance, got %#v", snapshot)
		}
		if got := EvaluatePoint12Val0DependencyState(snapshot); got != Point12Val0DependencyStateActive {
			t.Fatalf("expected active dependency state for computed snapshot, got %#v", snapshot)
		}
	})

	t.Run("non computed review provenance through helper blocks", func(t *testing.T) {
		valD := activePoint11ValDFoundation()
		snapshot := SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD, Point12Val0Point11ReviewContext{
			SnapshotFromComputedOutput: false,
		})
		if snapshot.SnapshotFromComputedOutput {
			t.Fatalf("expected helper to preserve non-computed provenance, got %#v", snapshot)
		}
		if got := EvaluatePoint12Val0DependencyState(snapshot); got == Point12Val0DependencyStateActive {
			t.Fatalf("expected non-computed snapshot to block, got %#v", snapshot)
		}
	})

	t.Run("zero review context does not become active accidentally", func(t *testing.T) {
		valD := activePoint11ValDFoundation()
		snapshot := SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD, Point12Val0Point11ReviewContext{})
		if snapshot.SnapshotFromComputedOutput {
			t.Fatalf("expected zero review context to remain non-computed, got %#v", snapshot)
		}
		if got := EvaluatePoint12Val0DependencyState(snapshot); got == Point12Val0DependencyStateActive {
			t.Fatalf("expected zero review context snapshot to not become active, got %#v", snapshot)
		}
	})
	t.Run("copied point11 vald projection disclaimer propagates exactly", func(t *testing.T) {
		valD := ComputePoint11ValDFoundation(Point11ValDFoundationModel())
		valD.ProjectionDisclaimer = "projection_only not_canonical_truth point11_vald_final_projection_disclaimer"
		snapshot := SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD, point12Val0DependencyReviewContextModel())
		if snapshot.ProjectionDisclaimer != valD.ProjectionDisclaimer {
			t.Fatalf("expected exact copied projection disclaimer, got snapshot=%q vald=%q", snapshot.ProjectionDisclaimer, valD.ProjectionDisclaimer)
		}
	})

	t.Run("full authority context sets are preserved from point11 vald", func(t *testing.T) {
		valD := activePoint11ValDFoundation()
		valD.QualityMap.PolicyRefs = []string{
			"policy_point11_vala_authority_core_v1",
			"policy_point11_vala_authority_core_v2",
		}
		valD.QualityMap.ClaimRefs = []string{
			"claim_point11_valb_customer_scope_001",
			"claim_point11_valb_customer_scope_002",
		}
		valD.QualityMap.GovernanceEventRefs = []string{
			"governance_event_point11_vald_quality_001",
			"governance_event_point11_vald_quality_002",
		}
		snapshot := SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD, point12Val0DependencyReviewContextModel())
		if !reflect.DeepEqual(snapshot.PolicyAuthorityContextRefs, valD.QualityMap.PolicyRefs) ||
			!reflect.DeepEqual(snapshot.ClaimAuthorityContextRefs, valD.QualityMap.ClaimRefs) ||
			!reflect.DeepEqual(snapshot.GovernanceAuthorityContextRefs, valD.QualityMap.GovernanceEventRefs) {
			t.Fatalf("expected exact authority context set copy, got %#v", snapshot)
		}
		if snapshot.PolicyAuthorityContextRef != valD.QualityMap.PolicyRefs[0] ||
			snapshot.ClaimAuthorityContextRef != valD.QualityMap.ClaimRefs[0] ||
			snapshot.GovernanceAuthorityContextRef != valD.QualityMap.GovernanceEventRefs[0] {
			t.Fatalf("expected primary authority refs to remain aligned with preserved sets, got %#v", snapshot)
		}
	})

	t.Run("foundation blocks when helper snapshot provenance is non computed", func(t *testing.T) {
		model := Point12Val0FoundationModel()
		valD := activePoint11ValDFoundation()
		model.Dependency = SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD, Point12Val0Point11ReviewContext{
			SnapshotFromComputedOutput: false,
		})
		model = ComputePoint12Val0Foundation(model)
		if model.DependencyState == Point12Val0DependencyStateActive || model.CurrentState == Point12Val0StateActive {
			t.Fatalf("expected full val0 foundation to block on non-computed dependency snapshot, got %#v", model)
		}
	})
	testCases := []struct {
		name   string
		mutate func(*Point12Val0DependencySnapshot)
		want   string
	}{
		{name: "missing point11 dependency blocks", mutate: func(model *Point12Val0DependencySnapshot) { *model = Point12Val0DependencySnapshot{} }, want: Point12Val0DependencyStateBlocked},
		{name: "point11 valc dependency blocks", mutate: func(model *Point12Val0DependencySnapshot) { model.UpstreamWaveID = "val_c" }, want: Point12Val0DependencyStateBlocked},
		{name: "regenerated fallback dependency snapshot blocks", mutate: func(model *Point12Val0DependencySnapshot) { model.SnapshotFromComputedOutput = false }, want: Point12Val0DependencyStateBlocked},
		{name: "malformed upstream closure manifest ref blocks", mutate: func(model *Point12Val0DependencySnapshot) { model.UpstreamClosureManifestRef = "manifest_unknown" }, want: Point12Val0DependencyStateBlocked},
		{name: "stale revoked unsupported upstream dependency blocks", mutate: func(model *Point12Val0DependencySnapshot) { model.UpstreamClosureManifestRef = "manifest_revoked" }, want: Point12Val0DependencyStateBlocked},
		{name: "missing authority context blocks", mutate: func(model *Point12Val0DependencySnapshot) {
			model.PolicyAuthorityContextRefs = nil
			model.PolicyAuthorityContextRef = ""
		}, want: Point12Val0DependencyStateBlocked},
		{name: "canonical looking junk policy authority ref blocks", mutate: func(model *Point12Val0DependencySnapshot) {
			model.PolicyAuthorityContextRefs = []string{"policy_unknown"}
			model.PolicyAuthorityContextRef = "policy_unknown"
		}, want: Point12Val0DependencyStateBlocked},
		{name: "canonical looking junk claim authority ref blocks", mutate: func(model *Point12Val0DependencySnapshot) {
			model.ClaimAuthorityContextRefs = []string{"claim_unknown"}
			model.ClaimAuthorityContextRef = "claim_unknown"
		}, want: Point12Val0DependencyStateBlocked},
		{name: "canonical looking junk governance authority ref blocks", mutate: func(model *Point12Val0DependencySnapshot) {
			model.GovernanceAuthorityContextRefs = []string{"governance_event_unknown"}
			model.GovernanceAuthorityContextRef = "governance_event_unknown"
		}, want: Point12Val0DependencyStateBlocked},
		{name: "primary authority ref outside preserved set blocks", mutate: func(model *Point12Val0DependencySnapshot) {
			model.PolicyAuthorityContextRef = "policy_point11_vala_authority_core_v9"
		}, want: Point12Val0DependencyStateBlocked},
		{name: "wrong point id blocks", mutate: func(model *Point12Val0DependencySnapshot) { model.UpstreamPointID = "point_10" }, want: Point12Val0DependencyStateBlocked},
		{name: "wrong wave id blocks", mutate: func(model *Point12Val0DependencySnapshot) { model.UpstreamWaveID = "val_b" }, want: Point12Val0DependencyStateBlocked},
		{name: "blocked upstream state wins over review required", mutate: func(model *Point12Val0DependencySnapshot) {
			model.UpstreamCurrentState = Point11ValDStateBlocked
			model.ReviewPrerequisites = []string{"external_review_prerequisite_point11_vald"}
		}, want: Point12Val0DependencyStateBlocked},
		{name: "review required upstream blocks final readiness", mutate: func(model *Point12Val0DependencySnapshot) {
			model.UpstreamCurrentState = Point11ValDStateReviewRequired
			model.UpstreamDependencyState = Point11ValDDependencyStateReviewRequired
			model.ReviewPrerequisites = []string{"external_review_prerequisite_point11_vald"}
		}, want: Point12Val0DependencyStateReviewRequired},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := point12Val0ActiveDependencySnapshot()
			testCase.mutate(&model)
			if got := EvaluatePoint12Val0DependencyState(model); got != testCase.want {
				t.Fatalf("expected dependency state %q, got %#v", testCase.want, model)
			}
		})
	}
}

func TestPoint12Val0ProofPackStateTaxonomy(t *testing.T) {
	for _, state := range []string{
		Point12Val0ProofPackStateDraft,
		Point12Val0ProofPackStateGenerated,
		Point12Val0ProofPackStateSignedMetadataValidated,
	} {
		t.Run(state+"_does_not_imply_pass", func(t *testing.T) {
			model := activePoint12Val0Foundation()
			model.Manifest.ProofPackState = state
			model.ReplayAssessment.ProofPackState = state
			model = ComputePoint12Val0Foundation(model)
			body, err := json.Marshal(model)
			if err != nil {
				t.Fatalf("marshal foundation: %v", err)
			}
			if strings.Contains(string(body), "point_12_pass") {
				t.Fatalf("expected no point12 pass emission, got %s", body)
			}
		})
	}

	for _, state := range []string{
		Point12Val0ProofPackStateTampered,
		Point12Val0ProofPackStateUnsupported,
		Point12Val0ProofPackStateExpired,
		Point12Val0ProofPackStateRevoked,
		Point12Val0ProofPackStateSuperseded,
		Point12Val0ProofPackStateBlocked,
	} {
		t.Run(state+"_cannot_become_active", func(t *testing.T) {
			model := activePoint12Val0Foundation()
			model.Manifest.ProofPackState = state
			model.ReplayAssessment.ProofPackState = state
			switch state {
			case Point12Val0ProofPackStateTampered:
				model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultTamperDetected
			case Point12Val0ProofPackStateUnsupported:
				model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultUnsupportedVersion
			default:
				model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultBlockedReplay
			}
			model = ComputePoint12Val0Foundation(model)
			if model.ManifestState != Point12Val0ManifestStateBlocked || model.CurrentState != Point12Val0StateBlocked {
				t.Fatalf("expected invalidating proof pack state to block, got %#v", model)
			}
		})
	}
}

func TestPoint12Val0ReplayResultTaxonomy(t *testing.T) {
	testCases := []struct {
		name   string
		mutate func(*Point12Val0Foundation)
	}{
		{name: "tampered manifest evidence yields tamper detected", mutate: func(model *Point12Val0Foundation) {
			model.Manifest.ProofPackState = Point12Val0ProofPackStateTampered
			model.ReplayAssessment.ProofPackState = Point12Val0ProofPackStateTampered
			model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultTamperDetected
		}},
		{name: "missing decisive evidence yields insufficient evidence", mutate: func(model *Point12Val0Foundation) {
			model.ReplayAssessment.DecisiveEvidencePresent = false
			model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultInsufficientEvidence
		}},
		{name: "unsupported schema engine policy yields unsupported version", mutate: func(model *Point12Val0Foundation) {
			model.Manifest.ProofPackState = Point12Val0ProofPackStateUnsupported
			model.ReplayAssessment.ProofPackState = Point12Val0ProofPackStateUnsupported
			model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultUnsupportedVersion
		}},
		{name: "policy mismatch yields policy mismatch", mutate: func(model *Point12Val0Foundation) {
			model.ReplayAssessment.ReplayPolicyRef = "policy_point12_val0_replay_changed_001"
			model.ReplayAssessment.ReplayPolicyHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
			model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultPolicyMismatch
		}},
		{name: "engine mismatch yields engine mismatch", mutate: func(model *Point12Val0Foundation) {
			model.ReplayAssessment.ReplayEngineHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
			model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultEngineMismatch
		}},
		{name: "schema mismatch yields schema mismatch", mutate: func(model *Point12Val0Foundation) {
			model.ReplayAssessment.ReplaySchemaVersion = "schema_version_point12_val0_v2"
			model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultSchemaMismatch
		}},
		{name: "evidence mismatch yields evidence mismatch", mutate: func(model *Point12Val0Foundation) {
			model.ReplayAssessment.ReplayEvidenceRefs = []string{"evidence:point12-proof-pack-evidence-002"}
			model.ReplayAssessment.ReplayEvidenceHashRefs = []string{"evidence_hash_point12_proof_pack_002"}
			model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultEvidenceMismatch
		}},
		{name: "claim mismatch yields claim mismatch", mutate: func(model *Point12Val0Foundation) {
			model.ReplayAssessment.ReplayClaimRefs = []string{"claim_point12_other_scope_001"}
			model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultClaimMismatch
		}},
		{name: "governance mismatch yields governance mismatch", mutate: func(model *Point12Val0Foundation) {
			model.ReplayAssessment.ReplayGovernanceRefs = []string{"governance_event_point12_val0_changed_001"}
			model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultGovernanceMismatch
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12Val0Foundation()
			testCase.mutate(&model)
			model = ComputePoint12Val0Foundation(model)
			if model.ReplayAssessmentState != Point12Val0ReplayAssessmentStateActive {
				t.Fatalf("expected active replay assessment classification, got %#v", model)
			}
		})
	}
}

func TestPoint12Val0DeterminismAndCompatibility(t *testing.T) {
	t.Run("original context does not silently use current policy", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.ReplayAssessment.ReplayPolicyRef = "policy_point12_val0_replay_changed_001"
		model.ReplayAssessment.ReplayPolicyHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultSameDecision
		model = ComputePoint12Val0Foundation(model)
		if model.ReplayAssessmentState != Point12Val0ReplayAssessmentStateBlocked {
			t.Fatalf("expected blocked replay assessment under silent policy drift, got %#v", model)
		}
	})

	t.Run("current policy context must be explicit", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.DeterminismContract.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model.CompatibilityProfile.ReplayMode = point12Val0ReplayModeCurrentPolicyContext
		model.CompatibilityProfile.PolicyCompatibility = point12Val0CompatibilityCompatibleAllowed
		model.CompatibilityProfile.CompatibilityEvidenceRefs = []string{"evidence:point12-compatibility-evidence-001"}
		model.ReplayAssessment.ReplayPolicyRef = "policy_point12_val0_replay_changed_001"
		model.ReplayAssessment.ReplayPolicyHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultDifferentDecision
		model.ReplayAssessment.DriftExplanation = "current_policy_context_explicitly_selected"
		model = ComputePoint12Val0Foundation(model)
		if model.CurrentState != Point12Val0StateActive {
			t.Fatalf("expected explicit current policy context to remain valid, got %#v", model)
		}
	})

	t.Run("comparison mode requires drift explanation", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.DeterminismContract.ReplayMode = point12Val0ReplayModeComparisonMode
		model.CompatibilityProfile.ReplayMode = point12Val0ReplayModeComparisonMode
		model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultDifferentDecision
		model.ReplayAssessment.DriftExplanation = ""
		model = ComputePoint12Val0Foundation(model)
		if model.ReplayAssessmentState != Point12Val0ReplayAssessmentStateBlocked {
			t.Fatalf("expected missing drift explanation to block, got %#v", model)
		}
	})

	t.Run("exact required mismatch blocks incorrect replay result", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.ReplayAssessment.ReplayEngineHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultSameDecision
		model = ComputePoint12Val0Foundation(model)
		if model.ReplayAssessmentState != Point12Val0ReplayAssessmentStateBlocked {
			t.Fatalf("expected exact-required engine mismatch to block same_decision, got %#v", model)
		}
	})

	t.Run("compatible allowed without explicit compatibility evidence blocks", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.CompatibilityProfile.PolicyCompatibility = point12Val0CompatibilityCompatibleAllowed
		model.CompatibilityProfile.CompatibilityEvidenceRefs = nil
		model = ComputePoint12Val0Foundation(model)
		if model.CompatibilityProfileState != Point12Val0CompatibilityProfileStateBlocked {
			t.Fatalf("expected missing compatibility evidence to block, got %#v", model)
		}
	})

	t.Run("unsupported profile blocks", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.CompatibilityProfile.UnsupportedBehavior = "unsupported_profile"
		model = ComputePoint12Val0Foundation(model)
		if model.CompatibilityProfileState != Point12Val0CompatibilityProfileStateBlocked {
			t.Fatalf("expected unsupported profile to block, got %#v", model)
		}
	})
}

func TestPoint12Val0ManifestValidation(t *testing.T) {
	t.Run("valid minimal val0 manifest passes foundation validation but does not emit point12 pass", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		if model.ManifestState != Point12Val0ManifestStateActive {
			t.Fatalf("expected active manifest state, got %#v", model)
		}
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal foundation: %v", err)
		}
		if strings.Contains(string(body), "point_12_pass") {
			t.Fatalf("expected no point12 pass emission, got %s", body)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point12Val0SignedProofPackManifest)
	}{
		{name: "missing proof pack id blocks", mutate: func(model *Point12Val0SignedProofPackManifest) { model.ProofPackID = "" }},
		{name: "missing artifact identity blocks", mutate: func(model *Point12Val0SignedProofPackManifest) { model.ArtifactRef = "" }},
		{name: "missing evidence identity blocks", mutate: func(model *Point12Val0SignedProofPackManifest) { model.EvidenceRefs = nil }},
		{name: "missing policy identity blocks", mutate: func(model *Point12Val0SignedProofPackManifest) { model.PolicyRef = "" }},
		{name: "missing engine identity blocks", mutate: func(model *Point12Val0SignedProofPackManifest) { model.EngineHash = "" }},
		{name: "missing schema identity blocks", mutate: func(model *Point12Val0SignedProofPackManifest) { model.SchemaVersion = "" }},
		{name: "canonical looking junk refs block", mutate: func(model *Point12Val0SignedProofPackManifest) { model.ProofPackID = "proof_pack_unknown" }},
		{name: "malformed non empty refs block", mutate: func(model *Point12Val0SignedProofPackManifest) {
			model.DependencySnapshotRef = "dependency snapshot invalid"
		}},
		{name: "cross tenant evidence ref blocks", mutate: func(model *Point12Val0SignedProofPackManifest) {
			model.EvidenceRefs = []string{"evidence:cross-tenant-proof-pack"}
		}},
		{name: "stale revoked expired superseded refs block", mutate: func(model *Point12Val0SignedProofPackManifest) { model.ClaimRefs = []string{"claim_revoked"} }},
		{name: "missing projection disclaimer blocks projection export readiness", mutate: func(model *Point12Val0SignedProofPackManifest) { model.ProjectionDisclaimer = "" }},
		{name: "missing retention class blocks export advisory readiness", mutate: func(model *Point12Val0SignedProofPackManifest) { model.RetentionClassRef = "" }},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12Val0Foundation()
			testCase.mutate(&model.Manifest)
			model = ComputePoint12Val0Foundation(model)
			if model.ManifestState != Point12Val0ManifestStateBlocked {
				t.Fatalf("expected blocked manifest state, got %#v", model)
			}
		})
	}
}

func TestPoint12Val0RedactionBoundaryState(t *testing.T) {
	t.Run("redacted non decisive field may remain partial advisory with limitation", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"non_decisive_annotation"}
		model.RedactionBoundary.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		model.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		model.RedactionBoundary.PostRedactionResult = Point12Val0ReplayResultRedactedLimitations
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateActive {
			t.Fatalf("expected active bounded redaction state, got %#v", model)
		}
	})

	t.Run("redacted decisive evidence yields insufficient evidence", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"decisive_evidence"}
		model.RedactionBoundary.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		model.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		model.RedactionBoundary.RedactionAffectsReplay = true
		model.RedactionBoundary.PostRedactionResult = Point12Val0ReplayResultInsufficientEvidence
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model.ReplayAssessment.ReplayResult = Point12Val0ReplayResultInsufficientEvidence
		model.ReplayAssessment.DecisiveEvidencePresent = false
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateActive || model.ReplayAssessmentState != Point12Val0ReplayAssessmentStateActive {
			t.Fatalf("expected insufficient evidence classification after decisive redaction, got %#v", model)
		}
	})

	t.Run("disallowed forbidden claim does not block by itself", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"marketing_claim"}
		model.RedactionBoundary.RedactionReasons = []string{"overclaim_removed"}
		model.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		model.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		model.RedactionBoundary.DisallowedClaimsAfterRedaction = []string{"production approved"}
		model.RedactionBoundary.RedactionSummary = "overclaim removed and bounded advisory wording preserved"
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateActive {
			t.Fatalf("expected disallowed forbidden ledger entry to remain active, got %#v", model)
		}
	})

	t.Run("multiple disallowed forbidden claims do not block by themselves", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"marketing_claim"}
		model.RedactionBoundary.RedactionReasons = []string{"overclaim_removed"}
		model.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		model.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		model.RedactionBoundary.DisallowedClaimsAfterRedaction = []string{
			"production approved",
			"compliance guaranteed",
			"certified",
		}
		model.RedactionBoundary.RedactionSummary = "forbidden marketing claims removed from the advisory pack"
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateActive {
			t.Fatalf("expected multiple disallowed forbidden ledger entries to remain active, got %#v", model)
		}
	})

	t.Run("forbidden claim in surviving output blocks", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"marketing_claim"}
		model.RedactionBoundary.RedactionReasons = []string{"overclaim_removed"}
		model.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		model.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		model.RedactionBoundary.DisallowedClaimsAfterRedaction = []string{"production approved"}
		model.RedactionBoundary.SurvivingClaimsAfterRedaction = []string{"production approved"}
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateBlocked {
			t.Fatalf("expected forbidden surviving claim to block, got %#v", model)
		}
	})

	t.Run("forbidden claim in customer visible export output blocks", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"marketing_claim"}
		model.RedactionBoundary.RedactionReasons = []string{"overclaim_removed"}
		model.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		model.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		model.RedactionBoundary.CustomerVisibleClaimsAfterRedaction = []string{"compliance guaranteed"}
		model.RedactionBoundary.ExportedClaimsAfterRedaction = []string{"bounded claim"}
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateBlocked {
			t.Fatalf("expected forbidden customer visible claim to block, got %#v", model)
		}
	})

	t.Run("forbidden claim in replay result claim blocks", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"marketing_claim"}
		model.RedactionBoundary.RedactionReasons = []string{"overclaim_removed"}
		model.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		model.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		model.RedactionBoundary.ReplayResultClaims = []string{"deployment approved"}
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateBlocked {
			t.Fatalf("expected forbidden replay result claim to block, got %#v", model)
		}
	})

	t.Run("disallowed field cannot hide surviving claim", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"marketing_claim"}
		model.RedactionBoundary.RedactionReasons = []string{"overclaim_removed"}
		model.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		model.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		model.RedactionBoundary.DisallowedClaimsAfterRedaction = []string{"production approved"}
		model.RedactionBoundary.CustomerVisibleClaimsAfterRedaction = []string{"production approved"}
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateBlocked {
			t.Fatalf("expected overlapping disallowed and surviving claim to block, got %#v", model)
		}
	})

	t.Run("redaction without reason approval where required blocks", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"decisive_evidence"}
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateBlocked {
			t.Fatalf("expected missing redaction governance to block, got %#v", model)
		}
	})

	t.Run("redaction cannot strengthen claim", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.MinimumSafeClaimAfterRedaction = "production approved"
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateBlocked {
			t.Fatalf("expected strengthened claim after redaction to block, got %#v", model)
		}
	})

	t.Run("minimum safe claim with compliance guaranteed blocks", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.MinimumSafeClaimAfterRedaction = "compliance guaranteed"
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateBlocked {
			t.Fatalf("expected forbidden minimum safe claim to block, got %#v", model)
		}
	})

	t.Run("redaction summary may describe disallowed claim as internal diagnostic context", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"marketing_claim"}
		model.RedactionBoundary.RedactionReasons = []string{"overclaim_removed"}
		model.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		model.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		model.RedactionBoundary.DisallowedClaimsAfterRedaction = []string{"production approved"}
		model.RedactionBoundary.RedactionSummary = "internal summary: disallowed production approved claim removed during redaction"
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateActive {
			t.Fatalf("expected internal redaction summary context to remain active, got %#v", model)
		}
	})

	t.Run("redaction cannot convert insufficient evidence into pass", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.RedactionBoundary.RedactedFields = []string{"decisive_evidence"}
		model.RedactionBoundary.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		model.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		model.RedactionBoundary.RedactionAffectsReplay = true
		model.RedactionBoundary.PostRedactionResult = Point12Val0ReplayResultSameDecision
		model.RedactionBoundary.PartialOrAdvisoryOnly = true
		model = ComputePoint12Val0Foundation(model)
		if model.RedactionBoundaryState != Point12Val0RedactionBoundaryStateBlocked {
			t.Fatalf("expected same_decision after decisive redaction to block, got %#v", model)
		}
	})
}

func TestPoint12Val0FinancialInsuranceEvidenceSupportState(t *testing.T) {
	t.Run("valid bounded evidence support profile passes metadata validation", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		if model.FinancialEvidenceSupportState != Point12Val0FinancialEvidenceSupportStateActive {
			t.Fatalf("expected active financial evidence support profile, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point12Val0FinancialInsuranceEvidenceSupportProfile)
	}{
		{name: "profile claiming premium reduction fails", mutate: func(model *Point12Val0FinancialInsuranceEvidenceSupportProfile) {
			model.SupportStatement = "This lowers insurance premium."
		}},
		{name: "profile claiming rating improvement fails", mutate: func(model *Point12Val0FinancialInsuranceEvidenceSupportProfile) {
			model.SupportStatement = "This increases credit rating."
		}},
		{name: "profile claiming compliance guarantee fails", mutate: func(model *Point12Val0FinancialInsuranceEvidenceSupportProfile) {
			model.SupportStatement = "Compliance guaranteed."
		}},
		{name: "profile claiming dora hipaa compliance proof fails", mutate: func(model *Point12Val0FinancialInsuranceEvidenceSupportProfile) {
			model.SupportStatement = "This proves DORA compliance."
		}},
		{name: "profile missing required customer review fails", mutate: func(model *Point12Val0FinancialInsuranceEvidenceSupportProfile) { model.RequiredCustomerReview = false }},
		{name: "profile missing guarantee guard flags fails", mutate: func(model *Point12Val0FinancialInsuranceEvidenceSupportProfile) {
			model.NoPremiumGuarantee = false
			model.NoRatingClaim = false
			model.NoComplianceGuarantee = false
			model.NoFinancialGuarantee = false
		}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12Val0Foundation()
			testCase.mutate(&model.FinancialEvidenceSupportProfile)
			model = ComputePoint12Val0Foundation(model)
			if model.FinancialEvidenceSupportState != Point12Val0FinancialEvidenceSupportStateBlocked {
				t.Fatalf("expected blocked financial support state, got %#v", model)
			}
		})
	}
}

func TestPoint12Val0ProvenanceState(t *testing.T) {
	t.Run("agent finding can be referenced as lineage input only", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		if model.ProvenanceState != Point12Val0ProvenanceStateActive {
			t.Fatalf("expected active provenance profile, got %#v", model)
		}
	})

	t.Run("allowed ai evidence candidate types remain advisory lineage only", func(t *testing.T) {
		for _, agentType := range point12Val0AllowedAIEvidenceCandidateTypes() {
			if !point12Val0AIEvidenceCandidateTypeValid(agentType) {
				t.Fatalf("expected allowed AI candidate type %q to pass taxonomy validation", agentType)
			}
			model := activePoint12Val0Foundation()
			model.ProvenanceProfile.AgentLineages[0].AgentType = agentType
			model = ComputePoint12Val0Foundation(model)
			if model.ProvenanceState != Point12Val0ProvenanceStateActive || model.CurrentState != Point12Val0StateActive {
				t.Fatalf("expected allowed AI candidate type %q to remain active advisory lineage, got %#v", agentType, model)
			}
		}
	})

	t.Run("blocked ai evidence candidate types are rejected", func(t *testing.T) {
		for _, agentType := range point12Val0BlockedAIEvidenceCandidateTypes() {
			model := activePoint12Val0Foundation()
			model.ProvenanceProfile.AgentLineages[0].AgentType = agentType
			model = ComputePoint12Val0Foundation(model)
			if model.ProvenanceState != Point12Val0ProvenanceStateBlocked {
				t.Fatalf("expected blocked AI candidate type %q to block, got %#v", agentType, model)
			}
		}
	})

	t.Run("agent lineage binding matrix covers exact required fields", func(t *testing.T) {
		record := activePoint12Val0Foundation().ProvenanceProfile.AgentLineages[0]
		fields := point12Val0AgentLineageBindingFields(record)
		required := map[string]bool{
			"agent_id":                     false,
			"agent_type":                   false,
			"model_or_rule_version_ref":    false,
			"permission_manifest_hash":     false,
			"input_evidence_refs":          false,
			"audit_id":                     false,
			"recommendation_id":            false,
			"lineage_input_only":           false,
			"claims_certification_false":   false,
			"claims_source_of_truth_false": false,
			"emits_premature_pass_false":   false,
		}
		for _, field := range fields {
			if field.BindingClass == point12ValDBindingClassExactRequired {
				if _, ok := required[field.FieldName]; ok {
					required[field.FieldName] = true
				}
			}
		}
		for fieldName, seen := range required {
			if !seen {
				t.Fatalf("expected exact required AI provenance field %s in %#v", fieldName, fields)
			}
		}
		if !point12Val0AgentLineageBindingMatrixValid(record) {
			t.Fatalf("expected valid AI provenance binding matrix for %#v", record)
		}
	})

	t.Run("agent finding cannot certify or emit pass", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.ProvenanceProfile.AgentLineages[0].ClaimsCertification = true
		model.ProvenanceProfile.AgentLineages[0].EmitsPrematurePass = true
		model = ComputePoint12Val0Foundation(model)
		if model.ProvenanceState != Point12Val0ProvenanceStateBlocked {
			t.Fatalf("expected certification or pass-emitting lineage to block, got %#v", model)
		}
	})

	t.Run("missing decisive toolchain provenance returns review required", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.ProvenanceProfile.CIJobIDRef = ""
		model = ComputePoint12Val0Foundation(model)
		if model.ProvenanceState != Point12Val0ProvenanceStateReviewRequired || model.CurrentState != Point12Val0StateReviewRequired {
			t.Fatalf("expected review required provenance gap, got %#v", model)
		}
	})

	t.Run("malformed agent lineage ref blocks", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.ProvenanceProfile.AgentLineages[0].AgentID = "agent placeholder"
		model = ComputePoint12Val0Foundation(model)
		if model.ProvenanceState != Point12Val0ProvenanceStateBlocked {
			t.Fatalf("expected malformed lineage ref to block, got %#v", model)
		}
	})

	mutationTests := []struct {
		name   string
		mutate func(*Point12Val0AgentLineageRecord)
	}{
		{name: "agent id mutation blocks", mutate: func(record *Point12Val0AgentLineageRecord) { record.AgentID = "agent_lineage_point12_val0_999" }},
		{name: "model or rule version mutation blocks", mutate: func(record *Point12Val0AgentLineageRecord) { record.ModelOrRuleVersionRef = "model_version_invalid" }},
		{name: "permission manifest hash mutation blocks", mutate: func(record *Point12Val0AgentLineageRecord) {
			record.PermissionManifestHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		}},
		{name: "input evidence refs mutation blocks", mutate: func(record *Point12Val0AgentLineageRecord) {
			record.InputEvidenceRefs = []string{"evidence:point12-proof-pack-evidence-999"}
		}},
		{name: "audit id mutation blocks", mutate: func(record *Point12Val0AgentLineageRecord) { record.AuditID = "audit_point12_val0_agent_999" }},
		{name: "recommendation id mutation blocks", mutate: func(record *Point12Val0AgentLineageRecord) {
			record.RecommendationID = "recommendation_point12_val0_999"
		}},
		{name: "lineage input only false blocks", mutate: func(record *Point12Val0AgentLineageRecord) { record.LineageInputOnly = false }},
		{name: "claims certification true blocks", mutate: func(record *Point12Val0AgentLineageRecord) { record.ClaimsCertification = true }},
		{name: "claims source of truth true blocks", mutate: func(record *Point12Val0AgentLineageRecord) { record.ClaimsSourceOfTruth = true }},
		{name: "emits premature pass true blocks", mutate: func(record *Point12Val0AgentLineageRecord) { record.EmitsPrematurePass = true }},
	}
	for _, tc := range mutationTests {
		t.Run(tc.name, func(t *testing.T) {
			model := activePoint12Val0Foundation()
			tc.mutate(&model.ProvenanceProfile.AgentLineages[0])
			model = ComputePoint12Val0Foundation(model)
			if model.ProvenanceState != Point12Val0ProvenanceStateBlocked || model.CurrentState == Point12Val0StateActive {
				t.Fatalf("expected exact-required AI provenance mutation to fail closed, got %#v", model)
			}
		})
	}

	t.Run("no external call path introduced", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.ProvenanceProfile.IntroducesNetworkCallPath = true
		model = ComputePoint12Val0Foundation(model)
		if model.ProvenanceState != Point12Val0ProvenanceStateBlocked {
			t.Fatalf("expected introduced call path to block, got %#v", model)
		}
	})
}

func TestPoint12Val0NoOverclaimState(t *testing.T) {
	t.Run("forbidden phrases blocked in customer export facing output", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.NoOverclaimReview.ObservedCustomerFacingTexts = []string{"guaranteed secure"}
		model = ComputePoint12Val0Foundation(model)
		if model.NoOverclaimState != Point12Val0NoOverclaimStateBlocked {
			t.Fatalf("expected forbidden claim to block, got %#v", model)
		}
	})

	t.Run("ai authority wording blocked in customer export facing output", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.NoOverclaimReview.ObservedCustomerFacingTexts = []string{"AI-approved evidence pack"}
		model = ComputePoint12Val0Foundation(model)
		if model.NoOverclaimState != Point12Val0NoOverclaimStateBlocked {
			t.Fatalf("expected AI authority wording to block, got %#v", model)
		}
	})

	t.Run("allowed safe wording remains allowed", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.NoOverclaimReview.ObservedCustomerFacingTexts = []string{"This proof pack contains evidence that may support customer, auditor, financial, or insurance review."}
		model.NoOverclaimReview.ObservedExportFacingTexts = []string{"advisory projection only"}
		model = ComputePoint12Val0Foundation(model)
		if model.NoOverclaimState != Point12Val0NoOverclaimStateActive {
			t.Fatalf("expected allowed safe wording to remain active, got %#v", model)
		}
	})
}

func TestPoint12Val0PassTokenGuard(t *testing.T) {
	t.Run("val0 cannot emit point12 pass", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal foundation: %v", err)
		}
		if strings.Contains(string(body), "point_12_pass") {
			t.Fatalf("expected no point12 pass token in val0 output, got %s", body)
		}
	})

	t.Run("val0 cannot accept point12 pass as proof", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.Manifest.SignatureRef = "point_12_pass"
		model = ComputePoint12Val0Foundation(model)
		if model.ManifestState != Point12Val0ManifestStateBlocked {
			t.Fatalf("expected premature point12 pass proof to be rejected, got %#v", model)
		}
	})

	t.Run("point12 pass fixture if present is rejected as premature", func(t *testing.T) {
		model := activePoint12Val0Foundation()
		model.Manifest.ProofPackID = "point_12_pass"
		model = ComputePoint12Val0Foundation(model)
		if model.ManifestState != Point12Val0ManifestStateBlocked {
			t.Fatalf("expected premature point12 pass fixture to block, got %#v", model)
		}
	})
}

func TestPoint12Val0SourceBoundaryGuards(t *testing.T) {
	body := readPoint12Val0Source(t)
	for _, forbidden := range []string{
		"http.Get",
		"http.Post",
		"fetch(",
		"net/http",
		"internal/connectors",
		"internal/verifier",
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("expected point12 val0 production source to stay outside live connector/verify/api boundaries, found %q", forbidden)
		}
	}
	if strings.Contains(body, point12Val0PrematurePassToken()) {
		t.Fatalf("expected point12 val0 production source to avoid literal premature pass token emission")
	}
}

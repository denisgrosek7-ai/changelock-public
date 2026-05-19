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
	point12ValAActiveFoundationBaselineJSON []byte
	point12ValAActiveFoundationBaselineOnce sync.Once
)

func mustMarshalPoint12ValAFoundation(model Point12ValAFoundation) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint12ValAFoundation(payload []byte) Point12ValAFoundation {
	var clone Point12ValAFoundation
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

func activePoint12ValADependencySnapshot() Point12ValADependencySnapshot {
	val0 := activePoint12Val0Foundation()
	return SnapshotPoint12ValADependencyFromComputedVal0(val0, point12ValADependencyReviewContextModel())
}

func syncPoint12ValAManifestToDependency(model *Point12ValAFoundation) {
	model.Manifest.ProofPackID = model.Dependency.Val0Manifest.ProofPackID
	model.Manifest.DecisionID = model.Dependency.Val0Manifest.DecisionID
	model.Manifest.TenantScope = model.Dependency.Val0Manifest.TenantScope
	model.Manifest.ArtifactRef = model.Dependency.Val0Manifest.ArtifactRef
	model.Manifest.ArtifactHash = model.Dependency.Val0Manifest.ArtifactHash
	model.Manifest.EvidenceRefs = append([]string{}, model.Dependency.Val0Manifest.EvidenceRefs...)
	model.Manifest.EvidenceHashRefs = append([]string{}, model.Dependency.Val0Manifest.EvidenceHashRefs...)
	model.Manifest.PolicyRef = model.Dependency.Val0Manifest.PolicyRef
	model.Manifest.PolicyVersion = model.Dependency.Val0Manifest.PolicyVersion
	model.Manifest.PolicyHash = model.Dependency.Val0Manifest.PolicyHash
	model.Manifest.EngineVersion = model.Dependency.Val0Manifest.EngineVersion
	model.Manifest.EngineHash = model.Dependency.Val0Manifest.EngineHash
	model.Manifest.SchemaVersion = model.Dependency.Val0Manifest.SchemaVersion
	model.Manifest.SchemaHash = model.Dependency.Val0Manifest.SchemaHash
	model.Manifest.ClaimRefs = append([]string{}, model.Dependency.Val0Manifest.ClaimRefs...)
	model.Manifest.GovernanceEventRefs = append([]string{}, model.Dependency.Val0Manifest.GovernanceEventRefs...)
	model.Manifest.CompatibilityProfileRef = model.Dependency.Val0Manifest.CompatibilityProfileRef
	model.Manifest.UpstreamVal0SnapshotRef = model.Dependency.SnapshotRef
	model.Manifest.SigningKeyRef = model.Dependency.Val0Manifest.SigningKeyRef
	model.Manifest.SignatureRef = model.Dependency.Val0Manifest.SignatureRef
	model.Manifest.RedactionManifestRef = model.Dependency.Val0Manifest.RedactionManifestRef
	model.Manifest.RetentionClassRef = model.Dependency.Val0Manifest.RetentionClassRef
	model.Manifest.ToolchainProvenanceRefs = append([]string{}, model.Dependency.Val0Manifest.ToolchainProvenanceRefs...)
	model.Manifest.AgentLineageRefs = append([]string{}, model.Dependency.Val0Manifest.AgentLineageRefs...)
	model.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Manifest)
	model.Manifest.SignatureBoundManifestID = model.Manifest.ManifestID
	model.Manifest.SignatureBoundManifestPayloadHash = model.Manifest.ManifestPayloadHash
}

func uncachedActivePoint12ValAFoundation() Point12ValAFoundation {
	model := Point12ValAFoundationModel()
	model.Dependency = activePoint12ValADependencySnapshot()
	syncPoint12ValAManifestToDependency(&model)
	return ComputePoint12ValAFoundation(model)
}

func activePoint12ValAFoundation() Point12ValAFoundation {
	point12ValAActiveFoundationBaselineOnce.Do(func() {
		point12ValAActiveFoundationBaselineJSON = mustMarshalPoint12ValAFoundation(uncachedActivePoint12ValAFoundation())
	})
	return clonePoint12ValAFoundation(point12ValAActiveFoundationBaselineJSON)
}

func activePoint12ValAFoundationFromVal0(val0 Point12Val0Foundation) Point12ValAFoundation {
	model := Point12ValAFoundationModel()
	model.Dependency = SnapshotPoint12ValADependencyFromComputedVal0(val0, point12ValADependencyReviewContextModel())
	syncPoint12ValAManifestToDependency(&model)
	return ComputePoint12ValAFoundation(model)
}

func readPoint12ValASource(t *testing.T) string {
	t.Helper()
	for _, path := range []string{"point12_vala.go", "internal/formal/point12_vala.go"} {
		body, err := os.ReadFile(path)
		if err == nil {
			return string(body)
		}
	}
	t.Fatal("failed to read point12_vala.go source")
	return ""
}

func TestPoint12ValADependencyState(t *testing.T) {
	t.Run("valid computed val0 foundation allows manifest readiness", func(t *testing.T) {
		model := activePoint12ValAFoundation()
		if model.DependencyState != Point12ValADependencyStateActive {
			t.Fatalf("expected active dependency state, got %#v", model)
		}
		if model.CurrentState != Point12ValAStateActive {
			t.Fatalf("expected active vala foundation state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point12ValADependencySnapshot)
		want   string
	}{
		{name: "missing val0 dependency blocks", mutate: func(model *Point12ValADependencySnapshot) {
			*model = Point12ValADependencySnapshot{}
		}, want: Point12ValADependencyStateBlocked},
		{name: "fallback regenerated val0 snapshot blocks", mutate: func(model *Point12ValADependencySnapshot) {
			model.SnapshotFromComputedOutput = false
		}, want: Point12ValADependencyStateBlocked},
		{name: "wrong val0 point id blocks", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0PointID = "point_11"
		}, want: Point12ValADependencyStateBlocked},
		{name: "wrong val0 wave id blocks", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0WaveID = "val_x"
		}, want: Point12ValADependencyStateBlocked},
		{name: "val0 manifest blocked blocks vala", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0ManifestState = Point12Val0ManifestStateBlocked
		}, want: Point12ValADependencyStateBlocked},
		{name: "val0 redaction blocked blocks vala", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0RedactionBoundaryState = Point12Val0RedactionBoundaryStateBlocked
		}, want: Point12ValADependencyStateBlocked},
		{name: "val0 compatibility blocked blocks vala", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0CompatibilityProfileState = Point12Val0CompatibilityProfileStateBlocked
		}, want: Point12ValADependencyStateBlocked},
		{name: "val0 provenance blocked blocks vala", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0ProvenanceState = Point12Val0ProvenanceStateBlocked
		}, want: Point12ValADependencyStateBlocked},
		{name: "premature point12 pass in dependency blocks", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0PrematurePoint12PassSeen = true
		}, want: Point12ValADependencyStateBlocked},
		{name: "val0 review required propagates review required", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0CurrentState = Point12Val0StateReviewRequired
			model.Val0DependencyState = Point12Val0DependencyStateReviewRequired
			model.Val0ProvenanceState = Point12Val0ProvenanceStateReviewRequired
		}, want: Point12ValADependencyStateReviewRequired},
		{name: "padded val0 active state blocks raw inherited dependency binding", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0CurrentState = Point12Val0StateActive + " "
		}, want: Point12ValADependencyStateBlocked},
		{name: "tab newline val0 manifest point id blocks raw inherited dependency binding", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0Manifest.PointID = "\t" + point12Val0PointID + "\n"
		}, want: Point12ValADependencyStateBlocked},
		{name: "unsupported val0 profile context blocks inherited dependency binding", mutate: func(model *Point12ValADependencySnapshot) {
			model.Val0Manifest.ProfileContext.ProfileMatchOriginal = false
			model.Val0Manifest.ProfileContext.ProfileBindingStatus = Point12Val0ProfileBindingStatusUnsupported
			model.Val0Manifest.ProfileContext.ProfileMismatchReason = "profile_unsupported"
		}, want: Point12ValADependencyStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValADependencySnapshot()
			testCase.mutate(&model)
			if got := EvaluatePoint12ValADependencyState(model); got != testCase.want {
				t.Fatalf("expected dependency state %q, got %#v", testCase.want, model)
			}
		})
	}

	t.Run("profile approval pass token blocks inherited dependency binding with exact reason", func(t *testing.T) {
		model := activePoint12ValADependencySnapshot()
		model.Val0Manifest.ProfileContext.ProfileApprovalRef = "profile_approval_point_12_pass"
		state, reasons := point12ValADependencyStateAndReasons(model)
		if state != Point12ValADependencyStateBlocked ||
			!reflect.DeepEqual(reasons, []string{
				"dependency_inherited_profile_context_premature_point12_pass",
				"dependency_identity_or_profile_context_invalid",
			}) {
			t.Fatalf("expected inherited profile pass token exact block, state=%q reasons=%v model=%#v", state, reasons, model)
		}
	})
}

func TestPoint12ValAManifestIntegrityState(t *testing.T) {
	t.Run("valid minimal vala signed proof pack manifest metadata is active readiness", func(t *testing.T) {
		model := activePoint12ValAFoundation()
		if model.ManifestIntegrityState != Point12ValAManifestIntegrityStateActive {
			t.Fatalf("expected active manifest integrity state, got %#v", model)
		}
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal vala foundation: %v", err)
		}
		if strings.Contains(string(body), "point_12_pass") {
			t.Fatalf("expected active manifest readiness to not emit point12 pass, got %s", body)
		}
	})

	requiredFieldCases := []struct {
		name   string
		mutate func(*Point12ValASignedProofPackManifestCore)
	}{
		{name: "missing proof_pack_id blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.ProofPackID = "" }},
		{name: "missing manifest_id blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.ManifestID = "" }},
		{name: "missing decision_id blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.DecisionID = "" }},
		{name: "missing tenant_scope blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.TenantScope = "" }},
		{name: "missing artifact_ref blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.ArtifactRef = "" }},
		{name: "missing artifact_hash blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.ArtifactHash = "" }},
		{name: "missing evidence_refs blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.EvidenceRefs = nil }},
		{name: "missing evidence_hash_refs blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.EvidenceHashRefs = nil }},
		{name: "missing policy hash blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.PolicyHash = "" }},
		{name: "missing engine hash blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.EngineHash = "" }},
		{name: "missing schema hash blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.SchemaHash = "" }},
		{name: "missing manifest payload hash blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.ManifestPayloadHash = "" }},
		{name: "missing signing_key_ref blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.SigningKeyRef = "" }},
		{name: "missing signature metadata blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.SignatureMetadataRef = "" }},
		{name: "missing retention_class_ref blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.RetentionClassRef = "" }},
		{name: "missing projection_disclaimer blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.ProjectionDisclaimer = "" }},
	}

	for _, testCase := range requiredFieldCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := Point12ValAFoundationModel()
			testCase.mutate(&model.Manifest)
			if got := ComputePoint12ValAFoundation(model).ManifestIntegrityState; got != Point12ValAManifestIntegrityStateBlocked {
				t.Fatalf("expected blocked manifest integrity for %s, got %#v", testCase.name, model)
			}
		})
	}

	malformedCases := []struct {
		name   string
		mutate func(*Point12ValASignedProofPackManifestCore)
	}{
		{name: "malformed proof_pack_id blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.ProofPackID = "proof pack bad" }},
		{name: "canonical looking junk manifest_id blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.ManifestID = "manifest_unknown" }},
		{name: "malformed evidence ref blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.EvidenceRefs = []string{"evidence invalid"} }},
		{name: "malformed hash blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.ArtifactHash = "sha256:not-a-real-hash" }},
		{name: "malformed signing key ref blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.SigningKeyRef = "signing key invalid" }},
		{name: "malformed signature ref blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.SignatureRef = "signature invalid" }},
		{name: "malformed algorithm ref blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) { model.HashAlgorithmRef = "hash algorithm invalid" }},
		{name: "padded signature algorithm cannot trim into supported algorithm", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.SignatureAlgorithmRef = point12ValASignatureAlgorithmEd25519 + " "
		}},
		{name: "tab newline signing key state cannot trim into active state", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.SigningKeyState = "\t" + point12ValASigningKeyStateActive + "\n"
		}},
		{name: "cross tenant evidence ref blocks", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.EvidenceRefs = []string{"evidence:cross-tenant-pack-001"}
		}},
	}

	for _, testCase := range malformedCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := Point12ValAFoundationModel()
			testCase.mutate(&model.Manifest)
			if got := ComputePoint12ValAFoundation(model).ManifestIntegrityState; got != Point12ValAManifestIntegrityStateBlocked {
				t.Fatalf("expected blocked manifest integrity for malformed input, got %#v", model)
			}
		})
	}

	tamperCases := []struct {
		name   string
		mutate func(*Point12ValASignedProofPackManifestCore)
	}{
		{name: "artifact hash mismatch yields tampered", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.ArtifactHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		}},
		{name: "evidence hash mismatch yields tampered", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.EvidenceHashRefs = []string{"evidence_hash_point12_proof_pack_002"}
		}},
		{name: "policy hash mismatch yields tampered", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.PolicyHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		}},
		{name: "engine hash mismatch yields tampered", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.EngineHash = "sha256:cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"
		}},
		{name: "schema hash mismatch yields tampered", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.SchemaHash = "sha256:dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd"
		}},
		{name: "manifest payload hash mismatch yields tampered", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.ManifestPayloadHash = "sha256:eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
		}},
		{name: "signature metadata binding mismatch yields tampered", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.SignatureBoundManifestPayloadHash = "sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
		}},
		{name: "tenant scope mismatch yields tampered", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.TenantScope = "tenant_scope_point12_beta"
		}},
		{name: "decision id mismatch yields tampered", mutate: func(model *Point12ValASignedProofPackManifestCore) {
			model.DecisionID = "decision_point12_vala_changed_001"
		}},
	}

	for _, testCase := range tamperCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := Point12ValAFoundationModel()
			testCase.mutate(&model.Manifest)
			if got := ComputePoint12ValAFoundation(model).ManifestIntegrityState; got != Point12ValAManifestIntegrityStateTampered {
				t.Fatalf("expected tampered manifest integrity, got %#v", model)
			}
		})
	}

	t.Run("supported hash and signature algorithms remain active", func(t *testing.T) {
		model := activePoint12ValAFoundation()
		if model.ManifestIntegrityState != Point12ValAManifestIntegrityStateActive {
			t.Fatalf("expected active manifest integrity for supported algorithms, got %#v", model)
		}
	})
	t.Run("schema version match alone is insufficient when schema hash drifts", func(t *testing.T) {
		model := activePoint12ValAFoundation()
		model.Manifest.SchemaVersion = model.Dependency.Val0Manifest.SchemaVersion
		model.Manifest.SchemaHash = "sha256:abababababababababababababababababababababababababababababababab"
		model.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Manifest)
		model.Manifest.SignatureBoundManifestPayloadHash = model.Manifest.ManifestPayloadHash
		model = ComputePoint12ValAFoundation(model)
		if model.ManifestIntegrityState == Point12ValAManifestIntegrityStateActive {
			t.Fatalf("expected schema hash drift to block active manifest integrity, got %#v", model)
		}
	})
	t.Run("padded signature manifest id cannot trim into valid binding", func(t *testing.T) {
		model := activePoint12ValAFoundation()
		model.Manifest.SignatureBoundManifestID = " " + model.Manifest.ManifestID
		state, reasons := point12ValAManifestIntegrityStateAndReasons(model.Manifest, model.Dependency)
		if state != Point12ValAManifestIntegrityStateTampered ||
			!point12Val0StringSliceContains(reasons, "manifest_signature_manifest_id_binding_mismatch") {
			t.Fatalf("expected raw signature manifest id mismatch, state=%q reasons=%v model=%#v", state, reasons, model)
		}
	})
	t.Run("tab newline signature payload hash cannot trim into valid binding", func(t *testing.T) {
		model := activePoint12ValAFoundation()
		model.Manifest.SignatureBoundManifestPayloadHash = "\t" + model.Manifest.ManifestPayloadHash + "\n"
		state, reasons := point12ValAManifestIntegrityStateAndReasons(model.Manifest, model.Dependency)
		if state != Point12ValAManifestIntegrityStateTampered ||
			!point12Val0StringSliceContains(reasons, "manifest_signature_payload_hash_binding_mismatch") {
			t.Fatalf("expected raw signature payload hash mismatch, state=%q reasons=%v model=%#v", state, reasons, model)
		}
	})
	t.Run("self consistent substituted profile context is signed and bound to val0 dependency", func(t *testing.T) {
		model := activePoint12ValAFoundation()
		model.Manifest.ProfileContext.OriginalProfileID = "profile_point12_replay_substituted_002"
		model.Manifest.ProfileContext.CurrentProfileID = "profile_point12_replay_substituted_002"
		model.Manifest.ProfileContext.OriginalProfileVersion = "profile_version_point12_replay_v2"
		model.Manifest.ProfileContext.CurrentProfileVersion = "profile_version_point12_replay_v2"
		model.Manifest.ProfileContext.OriginalProfileHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		model.Manifest.ProfileContext.CurrentProfileHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		model.Manifest.ProfileContext.ProfileApprovalRef = "profile_approval_point12_replay_002"
		model.Manifest.ProfileContext.ProfileSignatureRef = "profile_signature_point12_replay_002"
		model.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Manifest)
		model.Manifest.SignatureBoundManifestPayloadHash = model.Manifest.ManifestPayloadHash
		state, reasons := point12ValAManifestIntegrityStateAndReasons(model.Manifest, model.Dependency)
		if state != Point12ValAManifestIntegrityStateTampered ||
			!point12Val0StringSliceContains(reasons, "manifest_profile_context_binding_mismatch") {
			t.Fatalf("expected profile context binding mismatch, state=%q reasons=%v model=%#v", state, reasons, model)
		}
	})
	for _, testCase := range []struct {
		name   string
		mutate func(*Point12Val0ReplayProfileContext)
	}{
		{name: "isolated profile id drift", mutate: func(profile *Point12Val0ReplayProfileContext) {
			profile.OriginalProfileID = "profile_point12_replay_substituted_002"
			profile.CurrentProfileID = "profile_point12_replay_substituted_002"
		}},
		{name: "isolated profile version drift", mutate: func(profile *Point12Val0ReplayProfileContext) {
			profile.OriginalProfileVersion = "profile_version_point12_replay_v2"
			profile.CurrentProfileVersion = "profile_version_point12_replay_v2"
		}},
		{name: "isolated profile approval ref drift", mutate: func(profile *Point12Val0ReplayProfileContext) {
			profile.ProfileApprovalRef = "profile_approval_point12_replay_002"
		}},
		{name: "isolated profile signature ref drift", mutate: func(profile *Point12Val0ReplayProfileContext) {
			profile.ProfileSignatureRef = "profile_signature_point12_replay_002"
		}},
	} {
		t.Run(testCase.name+" is signed but still bound to val0 dependency", func(t *testing.T) {
			model := activePoint12ValAFoundation()
			testCase.mutate(&model.Manifest.ProfileContext)
			model.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Manifest)
			model.Manifest.SignatureBoundManifestPayloadHash = model.Manifest.ManifestPayloadHash
			state, reasons := point12ValAManifestIntegrityStateAndReasons(model.Manifest, model.Dependency)
			if state != Point12ValAManifestIntegrityStateTampered ||
				!reflect.DeepEqual(reasons, []string{"manifest_profile_context_binding_mismatch"}) {
				t.Fatalf("expected isolated profile context binding mismatch, state=%q reasons=%v model=%#v", state, reasons, model)
			}
		})
	}
	t.Run("unsupported hash algorithm yields unsupported", func(t *testing.T) {
		model := Point12ValAFoundationModel()
		model.Manifest.HashAlgorithmRef = "hash_algorithm_blake3"
		if got := ComputePoint12ValAFoundation(model).ManifestIntegrityState; got != Point12ValAManifestIntegrityStateUnsupported {
			t.Fatalf("expected unsupported manifest integrity state, got %#v", model)
		}
	})

	t.Run("unsupported signature algorithm yields unsupported", func(t *testing.T) {
		model := Point12ValAFoundationModel()
		model.Manifest.SignatureAlgorithmRef = "signature_algorithm_rsassa_pss_metadata"
		if got := ComputePoint12ValAFoundation(model).ManifestIntegrityState; got != Point12ValAManifestIntegrityStateUnsupported {
			t.Fatalf("expected unsupported manifest integrity state, got %#v", model)
		}
	})

	for _, state := range []struct {
		name string
		key  string
		want string
	}{
		{name: "signing key active allowed", key: point12ValASigningKeyStateActive, want: Point12ValAManifestIntegrityStateActive},
		{name: "signing key revoked blocks", key: point12ValASigningKeyStateRevoked, want: Point12ValAManifestIntegrityStateBlocked},
		{name: "signing key expired blocks", key: point12ValASigningKeyStateExpired, want: Point12ValAManifestIntegrityStateBlocked},
		{name: "signing key compromised blocks", key: point12ValASigningKeyStateCompromised, want: Point12ValAManifestIntegrityStateBlocked},
		{name: "signing key unknown returns review required", key: point12ValASigningKeyStateUnknown, want: Point12ValAManifestIntegrityStateReviewRequired},
	} {
		t.Run(state.name, func(t *testing.T) {
			model := Point12ValAFoundationModel()
			model.Manifest.SigningKeyState = state.key
			if got := ComputePoint12ValAFoundation(model).ManifestIntegrityState; got != state.want {
				t.Fatalf("expected %q, got %#v", state.want, model)
			}
		})
	}

	t.Run("detached signature metadata can be validated as metadata only", func(t *testing.T) {
		model := Point12ValAFoundationModel()
		model.Manifest.SignatureRef = ""
		model.Manifest.DetachedSignatureRef = "detached_signature_point12_vala_metadata_001"
		model.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Manifest)
		model.Manifest.SignatureBoundManifestPayloadHash = model.Manifest.ManifestPayloadHash
		model = ComputePoint12ValAFoundation(model)
		if model.ManifestIntegrityState != Point12ValAManifestIntegrityStateActive {
			t.Fatalf("expected detached signature metadata validation to stay active, got %#v", model)
		}
	})
}

func TestPoint12ValANoRealSigningAndNoOverclaimBoundaries(t *testing.T) {
	t.Run("validator does not call external api or real signing primitives", func(t *testing.T) {
		body := readPoint12ValASource(t)
		for _, forbidden := range []string{
			"http.Get",
			"http.Post",
			"fetch(",
			"external_api",
			"external AI",
			"kms",
			"hsm",
			"KMS",
			"HSM",
			"Sign(",
			"GenerateKey",
			"PrivateKey",
			"crypto/rsa",
			"crypto/ecdsa",
			"crypto/ed25519",
		} {
			if strings.Contains(body, forbidden) {
				t.Fatalf("expected no external signing/runtime call path, found %q", forbidden)
			}
		}
	})

	t.Run("forbidden wording in manifest output claim blocks", func(t *testing.T) {
		model := Point12ValAFoundationModel()
		model.Manifest.ManifestOutputClaims = []string{"production approved"}
		model.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Manifest)
		model.Manifest.SignatureBoundManifestPayloadHash = model.Manifest.ManifestPayloadHash
		model = ComputePoint12ValAFoundation(model)
		if model.ManifestIntegrityState != Point12ValAManifestIntegrityStateBlocked {
			t.Fatalf("expected blocked manifest output overclaim, got %#v", model)
		}
	})

	t.Run("financial insurance premium and compliance guarantee wording remains blocked", func(t *testing.T) {
		model := Point12ValAFoundationModel()
		model.Manifest.ManifestOutputClaims = []string{"lower insurance premium", "compliance guaranteed"}
		model.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Manifest)
		model.Manifest.SignatureBoundManifestPayloadHash = model.Manifest.ManifestPayloadHash
		model = ComputePoint12ValAFoundation(model)
		if model.ManifestIntegrityState != Point12ValAManifestIntegrityStateBlocked {
			t.Fatalf("expected blocked financial or compliance overclaim, got %#v", model)
		}
	})

	t.Run("safe manifest wording remains active", func(t *testing.T) {
		model := Point12ValAFoundationModel()
		model.Manifest.ManifestOutputClaims = []string{"bounded claim"}
		model.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(model.Manifest)
		model.Manifest.SignatureBoundManifestPayloadHash = model.Manifest.ManifestPayloadHash
		model = ComputePoint12ValAFoundation(model)
		if model.ManifestIntegrityState != Point12ValAManifestIntegrityStateActive {
			t.Fatalf("expected safe manifest wording to remain active, got %#v", model)
		}
	})
}

func TestPoint12ValARedactionRegressionGuards(t *testing.T) {
	t.Run("disallowed claims after redaction may contain forbidden wording as denylist content", func(t *testing.T) {
		val0 := activePoint12Val0Foundation()
		val0.RedactionBoundary.DisallowedClaimsAfterRedaction = []string{"production approved"}
		val0.RedactionBoundary.RedactionSummary = "internal summary: disallowed production approved claim removed during redaction"
		val0.RedactionBoundary.PartialOrAdvisoryOnly = true
		val0 = ComputePoint12Val0Foundation(val0)
		if val0.RedactionBoundaryState != Point12Val0RedactionBoundaryStateActive {
			t.Fatalf("expected active val0 redaction boundary, got %#v", val0)
		}
		model := activePoint12ValAFoundationFromVal0(val0)
		if model.DependencyState != Point12ValADependencyStateActive {
			t.Fatalf("expected vala dependency to preserve active val0 redaction boundary, got %#v", model)
		}
	})

	t.Run("minimum safe claim after redaction containing forbidden wording blocks", func(t *testing.T) {
		val0 := activePoint12Val0Foundation()
		val0.RedactionBoundary.MinimumSafeClaimAfterRedaction = "production approved"
		val0 = ComputePoint12Val0Foundation(val0)
		model := activePoint12ValAFoundationFromVal0(val0)
		if model.DependencyState != Point12ValADependencyStateBlocked {
			t.Fatalf("expected blocked vala dependency on invalid val0 redaction boundary, got %#v", model)
		}
	})

	t.Run("redaction summary internal diagnostic may mention removed forbidden claim", func(t *testing.T) {
		val0 := activePoint12Val0Foundation()
		val0.RedactionBoundary.RedactedFields = []string{"marketing_claim"}
		val0.RedactionBoundary.RedactionReasons = []string{"overclaim_removed"}
		val0.RedactionBoundary.RedactionApproverRef = "redaction_approver_point12_val0"
		val0.RedactionBoundary.RedactionApprovalEventRef = "governance_event_point12_val0_redaction_001"
		val0.RedactionBoundary.DisallowedClaimsAfterRedaction = []string{"production approved"}
		val0.RedactionBoundary.RedactionSummary = "internal summary: disallowed production approved claim removed during redaction"
		val0.RedactionBoundary.PartialOrAdvisoryOnly = true
		val0 = ComputePoint12Val0Foundation(val0)
		model := activePoint12ValAFoundationFromVal0(val0)
		if model.DependencyState != Point12ValADependencyStateActive {
			t.Fatalf("expected active vala dependency with internal diagnostic redaction summary, got %#v", model)
		}
	})

	t.Run("surviving export customer replay claims containing forbidden wording block", func(t *testing.T) {
		val0 := activePoint12Val0Foundation()
		val0.RedactionBoundary.CustomerVisibleClaimsAfterRedaction = []string{"production approved"}
		val0 = ComputePoint12Val0Foundation(val0)
		model := activePoint12ValAFoundationFromVal0(val0)
		if model.DependencyState != Point12ValADependencyStateBlocked {
			t.Fatalf("expected blocked vala dependency after forbidden surviving redaction claim, got %#v", model)
		}
	})
}

func TestPoint12ValAPassTokenGuard(t *testing.T) {
	t.Run("vala cannot emit point12 pass", func(t *testing.T) {
		model := activePoint12ValAFoundation()
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal vala foundation: %v", err)
		}
		if strings.Contains(string(body), "point_12_pass") {
			t.Fatalf("expected no point12 pass emission, got %s", body)
		}
	})

	t.Run("vala cannot accept point12 pass as proof", func(t *testing.T) {
		model := Point12ValAFoundationModel()
		model.Manifest.SignatureRef = "point_12_pass"
		model = ComputePoint12ValAFoundation(model)
		if model.ManifestIntegrityState != Point12ValAManifestIntegrityStateBlocked {
			t.Fatalf("expected premature point12 pass proof to block, got %#v", model)
		}
	})

	t.Run("vala manifest artifact ref cannot carry point12 pass before final closure", func(t *testing.T) {
		model := activePoint12ValAFoundation()
		model.Manifest.ArtifactRef = "artifact_point_12_pass"
		state, reasons := point12ValAManifestIntegrityStateAndReasons(model.Manifest, model.Dependency)
		if state != Point12ValAManifestIntegrityStateBlocked || !point12Val0StringSliceContains(reasons, "manifest_premature_point12_pass") {
			t.Fatalf("expected premature point12 pass artifact ref to block ValA manifest, state=%s reasons=%#v", state, reasons)
		}
	})

	t.Run("vala inherited val0 manifest artifact ref cannot carry point12 pass", func(t *testing.T) {
		model := activePoint12ValAFoundation()
		model.Dependency.Val0Manifest.ArtifactRef = "artifact_point_12_pass"
		state := EvaluatePoint12ValADependencyState(model.Dependency)
		if state != Point12ValADependencyStateBlocked {
			t.Fatalf("expected inherited premature point12 pass artifact ref to block ValA dependency, got state=%s", state)
		}
	})

	t.Run("point12 pass fixture is rejected as premature", func(t *testing.T) {
		model := Point12ValAFoundationModel()
		model.Manifest.ProofPackID = "point_12_pass"
		model = ComputePoint12ValAFoundation(model)
		if model.ManifestIntegrityState != Point12ValAManifestIntegrityStateBlocked {
			t.Fatalf("expected premature point12 pass fixture to block, got %#v", model)
		}
	})
}

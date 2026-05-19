package formal

import (
	"encoding/json"
	"os"
	"strings"
	"sync"
	"testing"
)

type point12ValCBindingMatrixEntry struct {
	DownstreamModel string
	Field           string
	BindingClass    string
	Reason          string
}

var point12ValCBindingMatrixEntries = []point12ValCBindingMatrixEntry{
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "tenant_scope", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "proof_pack_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "manifest_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "replay_result_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "decision_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "artifact_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "artifact_hash", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "evidence_refs", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "evidence_hash_refs", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "policy_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "policy_version", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "policy_hash", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "engine_version", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "engine_hash", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "schema_version", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "schema_hash", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "manifest_payload_hash", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "signature_metadata_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "redaction_manifest_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "offline_bundle_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "public_private_classification", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "retention_class_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "retention_owner_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "disposal_path_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "projection_disclaimer", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "tenant_scope", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "proof_pack_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "manifest_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "replay_request_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "replay_result_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "artifact_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "artifact_hash", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "evidence_refs", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "evidence_hash_refs", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "policy_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "policy_version", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "policy_hash", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "engine_version", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "engine_hash", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "schema_version", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "schema_hash", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "manifest_payload_hash", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "signature_metadata_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "compatibility_profile_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "redaction_manifest_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "verification_policy_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "no_external_api_required", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "external_api_used", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "public_private_classification", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "retention_class_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCRedactionManifest", Field: "redaction_manifest_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCRedactionManifest", Field: "proof_pack_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCRedactionManifest", Field: "export_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCRedactionManifest", Field: "tenant_scope", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCRedactionManifest", Field: "redaction_reasons", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCRedactionManifest", Field: "redaction_approval_event_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCRedactionManifest", Field: "post_redaction_result", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCRedactionManifest", Field: "retention_class_ref", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCRedactionImpactVerdict", Field: "redaction_manifest_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCPublicPrivateBoundary", Field: "tenant_scope", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCPublicPrivateBoundary", Field: "export_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCPublicPrivateBoundary", Field: "offline_bundle_id", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCPublicPrivateBoundary", Field: "classification", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCPublicPrivateBoundary", Field: "allowed_audience", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCPublicPrivateBoundary", Field: "customer_visible_fields", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCPublicPrivateBoundary", Field: "auditor_visible_fields", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCPublicPrivateBoundary", Field: "internal_only_fields", BindingClass: "exact_required"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "limitations", BindingClass: "advisory_only", Reason: "bounded projection text does not replace canonical upstream identity"},
	{DownstreamModel: "Point12ValCRedactionManifest", Field: "redaction_summary", BindingClass: "advisory_only", Reason: "internal diagnostic context may explain removed claims but is not surviving/export output"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "offline_output_claims", BindingClass: "advisory_only", Reason: "output claims are bounded surfaces and cannot change upstream proof identity"},
	{DownstreamModel: "Point12ValCAuditExportBundle", Field: "generated_at", BindingClass: "intentionally_not_bound", Reason: "local bundle generation timestamp is packaging metadata, not upstream identity"},
	{DownstreamModel: "Point12ValCOfflineVerificationBundle", Field: "supported_verifier_versions", BindingClass: "intentionally_not_bound", Reason: "local verifier capability list constrains use but does not redefine upstream proof context"},
}

var (
	point12ValCActiveFoundationBaselineJSON []byte
	point12ValCActiveFoundationBaselineOnce sync.Once
)

func mustMarshalPoint12ValCFoundation(model Point12ValCFoundation) []byte {
	payload, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	return payload
}

func clonePoint12ValCFoundation(payload []byte) Point12ValCFoundation {
	var clone Point12ValCFoundation
	if err := json.Unmarshal(payload, &clone); err != nil {
		panic(err)
	}
	return clone
}

func activePoint12ValCDependencySnapshot() Point12ValCDependencySnapshot {
	valB := activePoint12ValBFoundation()
	return SnapshotPoint12ValCDependencyFromComputedValB(valB, point12ValCDependencyReviewContextModel())
}

func syncPoint12ValCFoundationToDependency(model *Point12ValCFoundation) {
	model.ExportBundle.ProofPackID = model.Dependency.ValAManifest.ProofPackID
	model.ExportBundle.ManifestID = model.Dependency.ValAManifest.ManifestID
	model.ExportBundle.ReplayResultID = model.Dependency.ValBReplayResult.ReplayResultID
	model.ExportBundle.DecisionID = model.Dependency.ValBReplayRequest.DecisionID
	model.ExportBundle.TenantScope = model.Dependency.ValBReplayRequest.TenantScope
	model.ExportBundle.ArtifactRef = model.Dependency.ValBReplayRequest.ArtifactRef
	model.ExportBundle.ArtifactHash = model.Dependency.ValBReplayRequest.ArtifactHash
	model.ExportBundle.EvidenceRefs = append([]string{}, model.Dependency.ValBReplayRequest.EvidenceRefs...)
	model.ExportBundle.EvidenceHashRefs = append([]string{}, model.Dependency.ValBReplayRequest.EvidenceHashRefs...)
	model.ExportBundle.PolicyRef = model.Dependency.ValBReplayRequest.PolicyRef
	model.ExportBundle.PolicyVersion = model.Dependency.ValBReplayRequest.PolicyVersion
	model.ExportBundle.PolicyHash = model.Dependency.ValBReplayRequest.PolicyHash
	model.ExportBundle.EngineVersion = model.Dependency.ValBReplayRequest.EngineVersion
	model.ExportBundle.EngineHash = model.Dependency.ValBReplayRequest.EngineHash
	model.ExportBundle.SchemaVersion = model.Dependency.ValBReplayRequest.SchemaVersion
	model.ExportBundle.SchemaHash = model.Dependency.ValBReplayRequest.SchemaHash
	model.ExportBundle.ClaimRefs = append([]string{}, model.Dependency.ValBReplayRequest.ClaimRefs...)
	model.ExportBundle.GovernanceEventRefs = append([]string{}, model.Dependency.ValBReplayRequest.GovernanceEventRefs...)
	model.ExportBundle.CompatibilityProfileRef = model.Dependency.ValBReplayRequest.CompatibilityProfileRef
	model.ExportBundle.RedactionManifestRef = model.Dependency.ValBReplayRequest.RedactionManifestRef
	model.ExportBundle.ManifestPayloadHash = model.Dependency.ValBReplayRequest.ManifestPayloadHash
	model.ExportBundle.SignatureMetadataRef = model.Dependency.ValAManifest.SignatureMetadataRef
	model.ExportBundle.RetentionClassRef = model.Dependency.ValAManifest.RetentionClassRef

	model.RedactionManifest.ProofPackID = model.Dependency.ValAManifest.ProofPackID
	model.RedactionManifest.ExportID = model.ExportBundle.ExportID
	model.RedactionManifest.TenantScope = model.Dependency.ValBReplayRequest.TenantScope
	model.RedactionManifest.RedactionManifestID = model.Dependency.ValBReplayRequest.RedactionManifestRef
	model.RedactionManifest.RedactionPolicyRef = model.Dependency.ValBReplayRequest.PolicyRef
	model.RedactionManifest.RedactionPolicyVersion = model.Dependency.ValBReplayRequest.PolicyVersion
	model.RedactionImpactVerdict.RedactionManifestID = model.RedactionManifest.RedactionManifestID

	model.OfflineBundle.ProofPackID = model.Dependency.ValAManifest.ProofPackID
	model.OfflineBundle.ManifestID = model.Dependency.ValAManifest.ManifestID
	model.OfflineBundle.ReplayRequestID = model.Dependency.ValBReplayRequest.ReplayRequestID
	model.OfflineBundle.ReplayResultID = model.Dependency.ValBReplayResult.ReplayResultID
	model.OfflineBundle.TenantScope = model.Dependency.ValBReplayRequest.TenantScope
	model.OfflineBundle.ArtifactRef = model.Dependency.ValBReplayRequest.ArtifactRef
	model.OfflineBundle.ArtifactHash = model.Dependency.ValBReplayRequest.ArtifactHash
	model.OfflineBundle.EvidenceRefs = append([]string{}, model.Dependency.ValBReplayRequest.EvidenceRefs...)
	model.OfflineBundle.EvidenceHashRefs = append([]string{}, model.Dependency.ValBReplayRequest.EvidenceHashRefs...)
	model.OfflineBundle.PolicyRef = model.Dependency.ValBReplayRequest.PolicyRef
	model.OfflineBundle.PolicyVersion = model.Dependency.ValBReplayRequest.PolicyVersion
	model.OfflineBundle.PolicyHash = model.Dependency.ValBReplayRequest.PolicyHash
	model.OfflineBundle.EngineVersion = model.Dependency.ValBReplayRequest.EngineVersion
	model.OfflineBundle.EngineHash = model.Dependency.ValBReplayRequest.EngineHash
	model.OfflineBundle.SchemaVersion = model.Dependency.ValBReplayRequest.SchemaVersion
	model.OfflineBundle.SchemaHash = model.Dependency.ValBReplayRequest.SchemaHash
	model.OfflineBundle.ManifestPayloadHash = model.Dependency.ValBReplayRequest.ManifestPayloadHash
	model.OfflineBundle.SignatureMetadataRef = model.Dependency.ValAManifest.SignatureMetadataRef
	model.OfflineBundle.DetachedSignatureRef = model.Dependency.ValAManifest.DetachedSignatureRef
	model.OfflineBundle.CompatibilityProfileRef = model.Dependency.ValBReplayRequest.CompatibilityProfileRef
	model.OfflineBundle.RedactionManifestRef = model.Dependency.ValBReplayRequest.RedactionManifestRef
	model.OfflineBundle.RetentionClassRef = model.Dependency.ValAManifest.RetentionClassRef

	model.PublicPrivateBoundary.TenantScope = model.Dependency.ValBReplayRequest.TenantScope
	model.PublicPrivateBoundary.ExportID = model.ExportBundle.ExportID
	model.PublicPrivateBoundary.OfflineBundleID = model.OfflineBundle.OfflineBundleID
}

func uncachedActivePoint12ValCFoundation() Point12ValCFoundation {
	model := Point12ValCFoundationModel()
	model.Dependency = activePoint12ValCDependencySnapshot()
	syncPoint12ValCFoundationToDependency(&model)
	return ComputePoint12ValCFoundation(model)
}

func activePoint12ValCFoundation() Point12ValCFoundation {
	point12ValCActiveFoundationBaselineOnce.Do(func() {
		point12ValCActiveFoundationBaselineJSON = mustMarshalPoint12ValCFoundation(uncachedActivePoint12ValCFoundation())
	})
	return clonePoint12ValCFoundation(point12ValCActiveFoundationBaselineJSON)
}

func activePoint12ValCFoundationFromValB(valB Point12ValBFoundation) Point12ValCFoundation {
	model := Point12ValCFoundationModel()
	model.Dependency = SnapshotPoint12ValCDependencyFromComputedValB(valB, point12ValCDependencyReviewContextModel())
	syncPoint12ValCFoundationToDependency(&model)
	return ComputePoint12ValCFoundation(model)
}

func readPoint12ValCSource(t *testing.T) string {
	t.Helper()
	for _, path := range []string{"point12_valc.go", "internal/formal/point12_valc.go"} {
		body, err := os.ReadFile(path)
		if err == nil {
			return string(body)
		}
	}
	t.Fatal("failed to read point12_valc.go source")
	return ""
}

func TestPoint12ValCBindingMatrixCoverage(t *testing.T) {
	if len(point12ValCBindingMatrixEntries) == 0 {
		t.Fatal("expected valc binding matrix entries")
	}

	requiredModels := map[string]bool{
		"Point12ValCAuditExportBundle":         false,
		"Point12ValCOfflineVerificationBundle": false,
		"Point12ValCRedactionManifest":         false,
		"Point12ValCRedactionImpactVerdict":    false,
		"Point12ValCPublicPrivateBoundary":     false,
	}
	for _, entry := range point12ValCBindingMatrixEntries {
		if entry.DownstreamModel == "" || entry.Field == "" || entry.BindingClass == "" {
			t.Fatalf("expected binding matrix entry identity to be populated, got %#v", entry)
		}
		switch entry.BindingClass {
		case "exact_required":
			requiredModels[entry.DownstreamModel] = true
		case "advisory_only", "intentionally_not_bound":
			if strings.TrimSpace(entry.Reason) == "" {
				t.Fatalf("expected non-exact binding to explain reason, got %#v", entry)
			}
		default:
			t.Fatalf("unexpected binding class %#v", entry)
		}
	}
	for model, seen := range requiredModels {
		if !seen {
			t.Fatalf("expected exact_required binding matrix coverage for %s", model)
		}
	}
}

func TestPoint12ValCDependencyState(t *testing.T) {
	t.Run("valid computed valb output allows valc to proceed", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		if model.DependencyState != Point12ValCDependencyStateActive {
			t.Fatalf("expected active valc dependency state, got %#v", model)
		}
		if model.CurrentState != Point12ValCStateActive {
			t.Fatalf("expected active valc state, got %#v", model)
		}
	})

	testCases := []struct {
		name   string
		mutate func(*Point12ValCDependencySnapshot)
		want   string
	}{
		{name: "missing valb dependency blocks", mutate: func(model *Point12ValCDependencySnapshot) { *model = Point12ValCDependencySnapshot{} }, want: Point12ValCDependencyStateBlocked},
		{name: "fallback regenerated valb snapshot blocks", mutate: func(model *Point12ValCDependencySnapshot) { model.SnapshotFromComputedOutput = false }, want: Point12ValCDependencyStateBlocked},
		{name: "unsupported manifest integrity check requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBManifestIntegrityResult = point12ValBCheckResultUnsupported
			model.ValBReplayResult.ManifestIntegrityCheckResult = point12ValBCheckResultUnsupported
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "unsupported signature metadata check requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBSignatureMetadataResult = point12ValBCheckResultUnsupported
			model.ValBReplayResult.SignatureMetadataCheckResult = point12ValBCheckResultUnsupported
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "unsupported compatibility check requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBCompatibilityResult = point12ValBCheckResultUnsupported
			model.ValBReplayResult.CompatibilityCheckResult = point12ValBCheckResultUnsupported
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "tampered manifest integrity check blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBManifestIntegrityResult = point12ValBCheckResultTampered
		}, want: Point12ValCDependencyStateBlocked},
		{name: "padded valb active state blocks raw inherited dependency binding", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBCurrentState = Point12ValBStateActive + " "
		}, want: Point12ValCDependencyStateBlocked},
		{name: "tab newline valb replay taxonomy blocks raw inherited dependency binding", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayTaxonomy = "\t" + Point12Val0ReplayResultSameDecision + "\n"
		}, want: Point12ValCDependencyStateBlocked},
		{name: "blocked signature metadata check blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBSignatureMetadataResult = point12ValBCheckResultBlocked
		}, want: Point12ValCDependencyStateBlocked},
		{name: "unknown manifest integrity check does not become active", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBManifestIntegrityResult = "manifest_integrity_check_result_unknown_001"
		}, want: Point12ValCDependencyStateBlocked},
		{name: "valb tamper detected blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayTaxonomy = Point12Val0ReplayResultTamperDetected
		}, want: Point12ValCDependencyStateBlocked},
		{name: "valb unsupported version requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayTaxonomy = Point12Val0ReplayResultUnsupportedVersion
			model.ValBReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultUnsupportedVersion
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "valb insufficient evidence requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayTaxonomy = Point12Val0ReplayResultInsufficientEvidence
			model.ValBReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultInsufficientEvidence
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "valb redacted limitations requires review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayTaxonomy = Point12Val0ReplayResultRedactedLimitations
			model.ValBReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultRedactedLimitations
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "premature point12 pass in dependency blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBPrematurePoint12PassSeen = true
		}, want: Point12ValCDependencyStateBlocked},
		{name: "stale embedded valb replay result pass emission blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayResult.PointPassEmitted = true
		}, want: Point12ValCDependencyStateBlocked},
		{name: "stale embedded valb replay result external api use blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayResult.ExternalAPIUsed = true
		}, want: Point12ValCDependencyStateBlocked},
		{name: "stale embedded valb replay result request id blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayResult.ReplayRequestID = "replay_request_point12_valb_stale_002"
		}, want: Point12ValCDependencyStateBlocked},
		{name: "stale embedded valb replay result replay mode blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayResult.ReplayMode = point12Val0ReplayModeComparisonMode
		}, want: Point12ValCDependencyStateBlocked},
		{name: "stale embedded valb replay result profile context blocks", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayResult.ProfileContext.ProfileApprovalBoundHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
		}, want: Point12ValCDependencyStateBlocked},
		{name: "embedded valb replay result review-required state propagates", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayResult.ReplayState = Point12ValBReplayResultStateReviewRequired
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "embedded valb replay result unsupported flag propagates review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayResult.UnsupportedVersion = true
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "embedded valb replay result match-original drift propagates review", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayResult.MatchOriginal = false
		}, want: Point12ValCDependencyStateReviewRequired},
		{name: "nested vala manifest missing current profile hash blocks dependency", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValAManifest.ProfileContext.CurrentProfileHash = ""
		}, want: Point12ValCDependencyStateBlocked},
		{name: "nested vala manifest profile pass token blocks dependency", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValAManifest.ProfileContext.ProfileApprovalRef = "profile_approval_point_12_pass"
		}, want: Point12ValCDependencyStateBlocked},
		{name: "nested valb replay request policy ref point12 pass blocks dependency", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValBReplayRequest.PolicyRef = "policy_point_12_pass"
		}, want: Point12ValCDependencyStateBlocked},
		{name: "nested vala manifest coordinated tenant retag blocks dependency", mutate: func(model *Point12ValCDependencySnapshot) {
			model.ValAManifest.TenantScope = "tenant_scope_point12_beta"
			model.ValAManifest.ProfileContext.OriginalTenantScope = "tenant_scope_point12_beta"
			model.ValAManifest.ProfileContext.CurrentTenantScope = "tenant_scope_point12_beta"
		}, want: Point12ValCDependencyStateBlocked},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValCDependencySnapshot()
			testCase.mutate(&model)
			if got := EvaluatePoint12ValCDependencyState(model); got != testCase.want {
				t.Fatalf("expected dependency state %q, got %#v", testCase.want, model)
			}
		})
	}

	t.Run("unsupported dependency enters review path and cannot remain export_ready", func(t *testing.T) {
		model := Point12ValCFoundationModel()
		model.Dependency = activePoint12ValCDependencySnapshot()
		model.Dependency.ValBManifestIntegrityResult = point12ValBCheckResultUnsupported
		model.Dependency.ValBReplayResult.ManifestIntegrityCheckResult = point12ValBCheckResultUnsupported
		syncPoint12ValCFoundationToDependency(&model)
		model = ComputePoint12ValCFoundation(model)
		if model.DependencyState != Point12ValCDependencyStateReviewRequired {
			t.Fatalf("expected dependency review-required for unsupported manifest check, got %#v", model)
		}
		if model.ExportState == Point12ValCExportStateReady || model.CurrentState == Point12ValCStateActive {
			t.Fatalf("expected unsupported dependency to avoid full export_ready/active path, got %#v", model)
		}
	})

	t.Run("tampered dependency still blocks full valc foundation", func(t *testing.T) {
		model := Point12ValCFoundationModel()
		model.Dependency = activePoint12ValCDependencySnapshot()
		model.Dependency.ValBManifestIntegrityResult = point12ValBCheckResultTampered
		syncPoint12ValCFoundationToDependency(&model)
		model = ComputePoint12ValCFoundation(model)
		if model.DependencyState != Point12ValCDependencyStateBlocked || model.CurrentState != Point12ValCStateBlocked {
			t.Fatalf("expected tampered dependency to block valc foundation, got %#v", model)
		}
	})
}

func TestPoint12ValCAuditExportBundle(t *testing.T) {
	t.Run("valid audit ready export metadata remains bounded and active", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		if model.ExportState != Point12ValCExportStateReady {
			t.Fatalf("expected export_ready, got %#v", model)
		}
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal valc foundation: %v", err)
		}
		if strings.Contains(string(body), "point_12_pass") {
			t.Fatalf("expected bounded export metadata to not emit point12 pass, got %s", body)
		}
	})

	testCases := []struct {
		name       string
		mutate     func(*Point12ValCFoundation)
		wantState  string
		wantReason string
	}{
		{name: "missing export id blocks", mutate: func(model *Point12ValCFoundation) { model.ExportBundle.ExportID = "" }, wantState: Point12ValCExportStateBlocked},
		{name: "missing tenant scope blocks", mutate: func(model *Point12ValCFoundation) { model.ExportBundle.TenantScope = "" }, wantState: Point12ValCExportStateTenantMismatch},
		{name: "missing proof pack manifest replay refs block", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.ProofPackID = ""
			model.ExportBundle.ManifestID = ""
			model.ExportBundle.ReplayResultID = ""
		}, wantState: Point12ValCExportStateBlocked},
		{name: "missing projection disclaimer blocks", mutate: func(model *Point12ValCFoundation) { model.ExportBundle.ProjectionDisclaimer = "" }, wantState: Point12ValCExportStateBlocked},
		{name: "missing retention owner and disposal path returns retention missing", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.RetentionOwnerRef = ""
			model.ExportBundle.DisposalPathRef = ""
		}, wantState: Point12ValCExportStateRetentionMissing},
		{name: "public private classification missing blocks", mutate: func(model *Point12ValCFoundation) { model.ExportBundle.PublicPrivateClassification = "" }, wantState: Point12ValCExportStateBlocked},
		{name: "padded policy hash cannot trim into export dependency binding", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.PolicyHash = model.Dependency.ValBReplayRequest.PolicyHash + " "
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_identity_or_metadata_invalid"},
		{name: "tab newline export state cannot trim into export ready", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.ExportState = "\t" + Point12ValCExportStateReady + "\n"
		}, wantState: Point12ValCExportStateBlocked},
		{name: "advisory only false blocks", mutate: func(model *Point12ValCFoundation) { model.ExportBundle.AdvisoryOnly = false }, wantState: Point12ValCExportStateBlocked},
		{name: "insufficient evidence cannot become export ready", mutate: func(model *Point12ValCFoundation) {
			model.Dependency.ValBReplayTaxonomy = Point12Val0ReplayResultInsufficientEvidence
			model.Dependency.ValBReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultInsufficientEvidence
			model.Dependency.ValBReplayResult.InsufficientEvidence = true
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked},
		{name: "verifier export requires offline bundle ref", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.ExportKind = point12ValCExportKindVerifierPackageMetadata
			model.ExportBundle.OfflineBundleRef = ""
		}, wantState: Point12ValCExportStateBlocked},
		{name: "export cannot accept point12 pass fixture", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.CustomerVisibleSummary = "point_12_pass"
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_premature_point12_pass"},
		{name: "export policy ref cannot carry point12 pass before final closure", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.PolicyRef = "policy_point_12_pass"
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_premature_point12_pass"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValCFoundation()
			testCase.mutate(&model)
			if testCase.wantReason != "" {
				redactionManifestState := EvaluatePoint12ValCRedactionManifestState(model.RedactionManifest, model.Dependency, model.ExportBundle)
				redactionImpactState := EvaluatePoint12ValCRedactionImpactState(model.RedactionImpactVerdict, model.RedactionManifest, model.Dependency)
				offlineState := EvaluatePoint12ValCOfflineBundleState(model.OfflineBundle, model.Dependency, redactionImpactState)
				boundaryState := EvaluatePoint12ValCPublicPrivateBoundaryState(model.PublicPrivateBoundary, model.Dependency, model.ExportBundle, model.OfflineBundle, model.RedactionManifest)
				exportState, exportReasons := point12ValCAuditExportStateAndReasons(model.ExportBundle, model.Dependency, redactionManifestState, redactionImpactState, offlineState, boundaryState)
				if exportState != testCase.wantState {
					t.Fatalf("expected direct export state %q, got state=%s reasons=%#v", testCase.wantState, exportState, exportReasons)
				}
				if !point12Val0StringSliceContains(exportReasons, testCase.wantReason) {
					t.Fatalf("expected exact export reason %q, got %#v", testCase.wantReason, exportReasons)
				}
			}
			model = ComputePoint12ValCFoundation(model)
			if model.ExportState != testCase.wantState {
				t.Fatalf("expected export state %q, got %#v", testCase.wantState, model)
			}
		})
	}

	t.Run("unsupported version may remain bounded review required", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.Dependency.ValBReplayTaxonomy = Point12Val0ReplayResultUnsupportedVersion
		model.Dependency.ValBReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultUnsupportedVersion
		model.Dependency.ValBReplayResult.UnsupportedVersion = true
		model.ExportBundle.ExportState = Point12ValCExportStateUnsupported
		model = ComputePoint12ValCFoundation(model)
		if model.ExportState != Point12ValCExportStateUnsupported || model.CurrentState != Point12ValCStateReviewRequired {
			t.Fatalf("expected unsupported export taxonomy with overall review-required state, got %#v", model)
		}
	})
}

func TestPoint12ValCRedactionManifestAndImpact(t *testing.T) {
	t.Run("valid non decisive redaction remains active with no decision impact", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		if model.RedactionManifestState != Point12ValCRedactionManifestStateActive {
			t.Fatalf("expected active redaction manifest state, got %#v", model)
		}
		if model.RedactionImpactState != Point12ValCRedactionImpactNoDecisionImpact {
			t.Fatalf("expected no_decision_impact, got %#v", model)
		}
		if model.RedactionManifest.RedactionManifestID != model.Dependency.ValBReplayRequest.RedactionManifestRef ||
			model.RedactionImpactVerdict.RedactionManifestID != model.RedactionManifest.RedactionManifestID {
			t.Fatalf("expected redaction chain to remain bound to upstream replay request, got %#v", model)
		}
	})

	t.Run("missing redaction reason blocks when redacted fields exist", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactedFields = []string{"customer_email"}
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
		model.RedactionManifest.RedactionApprovalEventRef = "governance_event_point12_valc_redaction_001"
		model.RedactionManifest.RedactionReasons = nil
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected blocked redaction manifest, got %#v", model)
		}
	})

	t.Run("missing redaction approval event blocks where required", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactedFields = []string{"customer_email"}
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
		model.RedactionManifest.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionManifest.RedactionApprovalEventRef = ""
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected blocked redaction manifest, got %#v", model)
		}
	})

	t.Run("missing redaction manifest export id blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.ExportID = ""
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected missing redaction export id to block, got %#v", model)
		}
	})

	t.Run("malformed redaction manifest export id blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.ExportID = "redaction_export_bad"
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected malformed redaction export id to block, got %#v", model)
		}
	})

	t.Run("redaction manifest export id drift blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.ExportID = "export_point12_valc_999"
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected redaction export id drift to block, got %#v", model)
		}
	})

	t.Run("redaction manifest id drift blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactionManifestID = "redaction_manifest_point12_valc_999"
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState == Point12ValCRedactionManifestStateActive && model.CurrentState == Point12ValCStateActive {
			t.Fatalf("expected redaction manifest id drift to avoid active/full-ready state, got %#v", model)
		}
	})

	t.Run("disallowed claims may contain production approved as denylist content", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.DisallowedClaimsAfterRedaction = []string{"production approved"}
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateActive {
			t.Fatalf("expected denylist-only disallowed claim to remain active, got %#v", model)
		}
	})

	t.Run("minimum safe claim production approved blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.MinimumSafeClaimAfterRedaction = "production approved"
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected forbidden minimum safe claim to block, got %#v", model)
		}
	})

	t.Run("redaction summary cannot carry point12 pass before final closure", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactionSummary = "point_12_pass"
		state, reasons := point12ValCRedactionManifestStateAndReasons(model.RedactionManifest, model.Dependency, model.ExportBundle)
		if state != Point12ValCRedactionManifestStateBlocked || !point12Val0StringSliceContains(reasons, "redaction_manifest_premature_point12_pass") {
			t.Fatalf("expected premature point12 pass redaction summary to block, state=%s reasons=%#v", state, reasons)
		}
	})

	t.Run("customer visible exported surviving replay claims with compliance guaranteed block", func(t *testing.T) {
		for _, mutate := range []func(*Point12ValCRedactionManifest){
			func(model *Point12ValCRedactionManifest) {
				model.CustomerVisibleClaimsAfterRedaction = []string{"compliance guaranteed"}
			},
			func(model *Point12ValCRedactionManifest) {
				model.ExportedClaimsAfterRedaction = []string{"compliance guaranteed"}
			},
			func(model *Point12ValCRedactionManifest) {
				model.SurvivingClaimsAfterRedaction = []string{"compliance guaranteed"}
			},
			func(model *Point12ValCRedactionManifest) {
				model.ReplayResultClaims = []string{"compliance guaranteed"}
			},
		} {
			model := activePoint12ValCFoundation()
			mutate(&model.RedactionManifest)
			model = ComputePoint12ValCFoundation(model)
			if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
				t.Fatalf("expected forbidden surviving/export claim to block, got %#v", model)
			}
		}
	})

	t.Run("internal redaction summary may describe removed forbidden claim", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.DisallowedClaimsAfterRedaction = []string{"production approved"}
		model.RedactionManifest.RedactionSummary = "internal summary: disallowed production approved claim removed during redaction"
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateActive {
			t.Fatalf("expected internal diagnostic summary to remain active, got %#v", model)
		}
	})

	t.Run("redaction cannot hide decisive missing evidence", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactedFields = []string{"decisive_evidence_hash"}
		model.RedactionManifest.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
		model.RedactionManifest.RedactionApprovalEventRef = "governance_event_point12_valc_redaction_001"
		model.RedactionManifest.RedactionAffectsDecision = true
		model.RedactionManifest.PostRedactionResult = Point12Val0ReplayResultSameDecision
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected decisive hidden evidence to block, got %#v", model)
		}
	})

	t.Run("decisive evidence removed cannot be no decision impact", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactedFields = []string{"decisive_evidence_hash"}
		model.RedactionManifest.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
		model.RedactionManifest.RedactionApprovalEventRef = "governance_event_point12_valc_redaction_001"
		model.RedactionManifest.RedactionAffectsReplay = true
		model.RedactionManifest.PostRedactionResult = Point12Val0ReplayResultRedactedLimitations
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model.RedactionManifest.Limitations = []string{"decisive evidence removed"}
		model.RedactionImpactVerdict.DecisiveEvidenceRemoved = true
		model.RedactionImpactVerdict.AffectsReplay = true
		model.RedactionImpactVerdict.PostRedactionResult = Point12Val0ReplayResultRedactedLimitations
		model.RedactionImpactVerdict.RedactionImpactState = Point12ValCRedactionImpactNoDecisionImpact
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionImpactState != Point12ValCRedactionImpactReviewRequired {
			t.Fatalf("expected review required redaction impact verdict, got %#v", model)
		}
	})

	t.Run("redaction impact manifest id drift blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionImpactVerdict.RedactionManifestID = "redaction_manifest_point12_valc_999"
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionImpactState == Point12ValCRedactionImpactNoDecisionImpact && model.CurrentState == Point12ValCStateActive {
			t.Fatalf("expected redaction impact id drift to avoid active/full-ready state, got %#v", model)
		}
	})

	t.Run("coordinated local redaction manifest substitution still blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactionManifestID = "redaction_manifest_point12_valc_999"
		model.RedactionImpactVerdict.RedactionManifestID = "redaction_manifest_point12_valc_999"
		model.ExportBundle.RedactionManifestRef = "redaction_manifest_point12_valc_999"
		model.OfflineBundle.RedactionManifestRef = "redaction_manifest_point12_valc_999"
		model.ExportBundle.ExportState = Point12ValCExportStateReady
		model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState == Point12ValCRedactionManifestStateActive &&
			model.RedactionImpactState == Point12ValCRedactionImpactNoDecisionImpact &&
			model.ExportState == Point12ValCExportStateReady &&
			model.OfflineBundleState == Point12ValCOfflineBundleStateActive &&
			model.CurrentState == Point12ValCStateActive {
			t.Fatalf("expected coordinated local redaction substitution to fail closed, got %#v", model)
		}
	})

	t.Run("redaction retention class cannot drift from proof pack class", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RetentionClassRef = "retention_class_point12_downgraded"
		state, reasons := point12ValCRedactionManifestStateAndReasons(model.RedactionManifest, model.Dependency, model.ExportBundle)
		if state != Point12ValCRedactionManifestStateBlocked {
			t.Fatalf("expected redaction retention class drift to block, got state=%s reasons=%#v", state, reasons)
		}
		if !point12Val0StringSliceContains(reasons, "redaction_manifest_dependency_binding_mismatch") {
			t.Fatalf("expected exact redaction retention binding reason, got %#v", reasons)
		}
	})

	t.Run("replay affecting redaction can produce redacted limitations with required limits", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.RedactionManifest.RedactedFields = []string{"customer_identifier"}
		model.RedactionManifest.RedactionReasons = []string{"privacy_redaction"}
		model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
		model.RedactionManifest.RedactionApprovalEventRef = "governance_event_point12_valc_redaction_001"
		model.RedactionManifest.RedactionAffectsReplay = true
		model.RedactionManifest.PostRedactionResult = Point12Val0ReplayResultRedactedLimitations
		model.RedactionManifest.PartialOrAdvisoryOnly = true
		model.RedactionManifest.Limitations = []string{"replay bounded by redaction"}
		model.RedactionImpactVerdict.AffectsReplay = true
		model.RedactionImpactVerdict.PostRedactionResult = Point12Val0ReplayResultRedactedLimitations
		model.RedactionImpactVerdict.Limitations = []string{"replay bounded by redaction"}
		model.RedactionImpactVerdict.RequiresPartialAdvisoryExport = true
		model.RedactionImpactVerdict.RedactionImpactState = Point12ValCRedactionImpactRedactedLimits
		model = ComputePoint12ValCFoundation(model)
		if model.RedactionManifestState != Point12ValCRedactionManifestStateActive || model.RedactionImpactState != Point12ValCRedactionImpactRedactedLimits {
			t.Fatalf("expected active redaction manifest with redacted limitations impact, got %#v", model)
		}
	})
}

func TestPoint12ValCOfflineBundleAndBoundary(t *testing.T) {
	t.Run("valid offline verification metadata remains active", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		if model.OfflineBundleState != Point12ValCOfflineBundleStateActive {
			t.Fatalf("expected active offline bundle state, got %#v", model)
		}
	})

	testCases := []struct {
		name       string
		mutate     func(*Point12ValCFoundation)
		wantState  string
		wantReason string
	}{
		{name: "no external api required false blocks", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.NoExternalAPIRequired = false
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "external api used true blocks", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ExternalAPIUsed = true
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "missing manifest proof replay refs block", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ManifestID = ""
			model.OfflineBundle.ProofPackID = ""
			model.OfflineBundle.ReplayResultID = ""
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "unsupported verifier version returns unsupported", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.RequestedVerifierVersion = "verifier_version_point12_valc_002"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateUnsupported
		}, wantState: Point12ValCOfflineBundleStateUnsupported},
		{name: "tenant mismatch blocks", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.TenantScope = "tenant_scope_point12_cross_001"
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "padded proof pack id cannot trim into offline dependency binding", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ProofPackID = " " + model.Dependency.ValAManifest.ProofPackID
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_identity_or_metadata_invalid"},
		{name: "tab newline offline state cannot trim into active state", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.OfflineState = "\t" + Point12ValCOfflineBundleStateActive + "\n"
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "cross tenant evidence blocks", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.EvidenceRefs = []string{"evidence:cross-tenant-pack-001"}
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "redacted decisive evidence forces limitations or blocks", mutate: func(model *Point12ValCFoundation) {
			model.RedactionManifest.RedactedFields = []string{"decisive_evidence_hash"}
			model.RedactionManifest.RedactionReasons = []string{"privacy_redaction"}
			model.RedactionManifest.RedactionApproverRef = "redaction_approver_point12_valc"
			model.RedactionManifest.RedactionApprovalEventRef = "governance_event_point12_valc_redaction_001"
			model.RedactionManifest.RedactionAffectsReplay = true
			model.RedactionManifest.PostRedactionResult = Point12Val0ReplayResultBlockedReplay
			model.RedactionManifest.PartialOrAdvisoryOnly = true
			model.RedactionManifest.Limitations = []string{"decisive evidence removed"}
			model.RedactionImpactVerdict.DecisiveEvidenceRemoved = true
			model.RedactionImpactVerdict.AffectsReplay = true
			model.RedactionImpactVerdict.PostRedactionResult = Point12Val0ReplayResultBlockedReplay
			model.RedactionImpactVerdict.Limitations = []string{"decisive evidence removed"}
			model.RedactionImpactVerdict.RequiresPartialAdvisoryExport = true
			model.RedactionImpactVerdict.RedactionImpactState = Point12ValCRedactionImpactBlockedReplay
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "offline bundle cannot accept point12 pass fixture", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.CustomerVisibleExplanation = "point_12_pass"
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_premature_point12_pass"},
		{name: "offline verification policy ref cannot carry point12 pass before final closure", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.VerificationPolicyRef = "policy_point_12_pass"
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_premature_point12_pass"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValCFoundation()
			testCase.mutate(&model)
			if testCase.wantReason != "" {
				redactionImpactState := EvaluatePoint12ValCRedactionImpactState(model.RedactionImpactVerdict, model.RedactionManifest, model.Dependency)
				offlineState, offlineReasons := point12ValCOfflineBundleStateAndReasons(model.OfflineBundle, model.Dependency, redactionImpactState)
				if offlineState != testCase.wantState {
					t.Fatalf("expected direct offline state %q, got state=%s reasons=%#v", testCase.wantState, offlineState, offlineReasons)
				}
				if !point12Val0StringSliceContains(offlineReasons, testCase.wantReason) {
					t.Fatalf("expected exact offline reason %q, got %#v", testCase.wantReason, offlineReasons)
				}
			}
			model = ComputePoint12ValCFoundation(model)
			if model.OfflineBundleState != testCase.wantState {
				t.Fatalf("expected offline state %q, got %#v", testCase.wantState, model)
			}
		})
	}

	t.Run("all exported fields classified keeps boundary active", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateActive {
			t.Fatalf("expected active public/private boundary, got %#v", model)
		}
	})

	t.Run("missing offline bundle id blocks boundary", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.PublicPrivateBoundary.OfflineBundleID = ""
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected missing offline bundle id to block boundary, got %#v", model)
		}
	})

	t.Run("malformed offline bundle id blocks boundary", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.PublicPrivateBoundary.OfflineBundleID = "offline_boundary_bad"
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected malformed offline bundle id to block boundary, got %#v", model)
		}
	})

	t.Run("offline bundle id drift blocks boundary", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.PublicPrivateBoundary.OfflineBundleID = "offline_bundle_point12_valc_999"
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected offline bundle id drift to block boundary, got %#v", model)
		}
	})

	t.Run("missing field classification blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.PublicPrivateBoundary.PrivateFields = []string{"artifact_hash"}
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected boundary classification failure to block, got %#v", model)
		}
	})

	t.Run("private field in customer visible output blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.ExportAudience = point12ValCExportAudienceCustomer
		model.PublicPrivateBoundary.CustomerVisibleFields = []string{"artifact_hash"}
		model.ExportBundle.CustomerVisibleSummary = "customer summary"
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected private customer visible field to block, got %#v", model)
		}
	})

	t.Run("unknown audience blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.PublicPrivateBoundary.AllowedAudience = "unknown audience"
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected unknown audience to block boundary, got %#v", model)
		}
	})

	t.Run("padded allowed audience blocks raw public private boundary binding", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.PublicPrivateBoundary.AllowedAudience = point12ValCExportAudienceInternalAudit + " "
		state, reasons := point12ValCPublicPrivateBoundaryStateAndReasons(model.PublicPrivateBoundary, model.Dependency, model.ExportBundle, model.OfflineBundle, model.RedactionManifest)
		if state != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected padded allowed audience to block, got state=%q reasons=%v", state, reasons)
		}
		if !point12Val0StringSliceContains(reasons, "public_private_boundary_identity_or_metadata_invalid") {
			t.Fatalf("expected exact metadata reason, got %v", reasons)
		}
	})

	t.Run("point12 pass field token blocks public private boundary", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.PublicPrivateBoundary.ExportedFields = append(model.PublicPrivateBoundary.ExportedFields, "point_12_pass")
		model.PublicPrivateBoundary.InternalOnlyFields = append(model.PublicPrivateBoundary.InternalOnlyFields, "point_12_pass")
		state, reasons := point12ValCPublicPrivateBoundaryStateAndReasons(model.PublicPrivateBoundary, model.Dependency, model.ExportBundle, model.OfflineBundle, model.RedactionManifest)
		if state != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected point12 pass field token to block boundary, got state=%q reasons=%v", state, reasons)
		}
		if !point12Val0StringSliceContains(reasons, "public_private_boundary_premature_point12_pass") {
			t.Fatalf("expected exact pass-token boundary reason, got %v", reasons)
		}
	})

	t.Run("redaction summary cannot leak private field into public output", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.ExportAudience = point12ValCExportAudienceCustomer
		model.RedactionManifest.RedactionSummary = "internal summary: artifact_hash removed during redaction"
		model = ComputePoint12ValCFoundation(model)
		if model.PublicPrivateBoundaryState != Point12ValCPublicPrivateBoundaryStateBlocked {
			t.Fatalf("expected private field leak through redaction summary to block, got %#v", model)
		}
	})
}

func TestPoint12ValCExactBindingMutationClosure(t *testing.T) {
	exportMutations := []struct {
		name       string
		mutate     func(*Point12ValCFoundation)
		wantState  string
		wantReason string
	}{
		{name: "artifact_hash exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.ArtifactHash = "sha256:1010101010101010101010101010101010101010101010101010101010101010"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "evidence_hash_refs exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.EvidenceHashRefs = []string{"sha256:2222222222222222222222222222222222222222222222222222222222222222"}
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "policy_hash exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.PolicyHash = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "engine_hash exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.EngineHash = "sha256:4444444444444444444444444444444444444444444444444444444444444444"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "schema_hash exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.SchemaHash = "sha256:5555555555555555555555555555555555555555555555555555555555555555"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "manifest_payload_hash exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.ManifestPayloadHash = "sha256:6666666666666666666666666666666666666666666666666666666666666666"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "retention_class_ref exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.RetentionClassRef = "retention_class_point12_downgraded"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "redaction_manifest_ref exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.RedactionManifestRef = "redaction_manifest_point12_valc_999"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "replay_result_id exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.ReplayResultID = "replay_result_point12_valb_999"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "tenant_scope exact_required mutation blocks full export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.TenantScope = "tenant_scope_point12_cross_001"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateTenantMismatch, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "manifest_id exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.ManifestID = "manifest_point12_vala_999"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
		{name: "proof_pack_id exact_required mutation blocks export", mutate: func(model *Point12ValCFoundation) {
			model.ExportBundle.ProofPackID = "proof_pack_point12_val0_999"
			model.ExportBundle.ExportState = Point12ValCExportStateReady
		}, wantState: Point12ValCExportStateBlocked, wantReason: "audit_export_dependency_binding_mismatch"},
	}
	for _, testCase := range exportMutations {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValCFoundation()
			testCase.mutate(&model)
			redactionManifestState := EvaluatePoint12ValCRedactionManifestState(model.RedactionManifest, model.Dependency, model.ExportBundle)
			redactionImpactState := EvaluatePoint12ValCRedactionImpactState(model.RedactionImpactVerdict, model.RedactionManifest, model.Dependency)
			offlineState := EvaluatePoint12ValCOfflineBundleState(model.OfflineBundle, model.Dependency, redactionImpactState)
			boundaryState := EvaluatePoint12ValCPublicPrivateBoundaryState(model.PublicPrivateBoundary, model.Dependency, model.ExportBundle, model.OfflineBundle, model.RedactionManifest)
			exportState, exportReasons := point12ValCAuditExportStateAndReasons(model.ExportBundle, model.Dependency, redactionManifestState, redactionImpactState, offlineState, boundaryState)
			if exportState != testCase.wantState {
				t.Fatalf("expected direct export state %q, got state=%s reasons=%#v", testCase.wantState, exportState, exportReasons)
			}
			if !point12Val0StringSliceContains(exportReasons, testCase.wantReason) {
				t.Fatalf("expected exact export reason %q, got %#v", testCase.wantReason, exportReasons)
			}
			model = ComputePoint12ValCFoundation(model)
			if model.ExportState != testCase.wantState {
				t.Fatalf("expected export state %q, got %#v", testCase.wantState, model)
			}
			if model.ExportState == Point12ValCExportStateReady || model.ExportState == Point12ValCExportStateProjectionOnly || model.CurrentState == Point12ValCStateActive {
				t.Fatalf("expected mutated exact_required export field to avoid full-ready state, got %#v", model)
			}
		})
	}

	offlineMutations := []struct {
		name       string
		mutate     func(*Point12ValCFoundation)
		wantState  string
		wantReason string
	}{
		{name: "artifact_hash exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ArtifactHash = "sha256:7777777777777777777777777777777777777777777777777777777777777777"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "evidence_hash_refs exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.EvidenceHashRefs = []string{"sha256:8888888888888888888888888888888888888888888888888888888888888888"}
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "policy_hash exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.PolicyHash = "sha256:9999999999999999999999999999999999999999999999999999999999999999"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "engine_hash exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.EngineHash = "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "schema_hash exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.SchemaHash = "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "manifest_payload_hash exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ManifestPayloadHash = "sha256:cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "detached_signature_ref exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.DetachedSignatureRef = "detached_signature_point12_vala_999"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "retention_class_ref exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.RetentionClassRef = "retention_class_point12_downgraded"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "redaction_manifest_ref exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.RedactionManifestRef = "redaction_manifest_point12_valc_999"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "replay_result_id exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ReplayResultID = "replay_result_point12_valb_999"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "tenant_scope exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.TenantScope = "tenant_scope_point12_cross_001"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "manifest_id exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ManifestID = "manifest_point12_vala_999"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "proof_pack_id exact_required mutation blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ProofPackID = "proof_pack_point12_val0_999"
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked, wantReason: "offline_bundle_dependency_binding_mismatch"},
		{name: "no_external_api_required false blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.NoExternalAPIRequired = false
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked},
		{name: "external_api_used true blocks offline bundle", mutate: func(model *Point12ValCFoundation) {
			model.OfflineBundle.ExternalAPIUsed = true
			model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		}, wantState: Point12ValCOfflineBundleStateBlocked},
	}
	for _, testCase := range offlineMutations {
		t.Run(testCase.name, func(t *testing.T) {
			model := activePoint12ValCFoundation()
			testCase.mutate(&model)
			if testCase.wantReason != "" {
				redactionImpactState := EvaluatePoint12ValCRedactionImpactState(model.RedactionImpactVerdict, model.RedactionManifest, model.Dependency)
				offlineState, offlineReasons := point12ValCOfflineBundleStateAndReasons(model.OfflineBundle, model.Dependency, redactionImpactState)
				if offlineState != testCase.wantState {
					t.Fatalf("expected direct offline state %q, got state=%s reasons=%#v", testCase.wantState, offlineState, offlineReasons)
				}
				if !point12Val0StringSliceContains(offlineReasons, testCase.wantReason) {
					t.Fatalf("expected exact offline reason %q, got %#v", testCase.wantReason, offlineReasons)
				}
			}
			model = ComputePoint12ValCFoundation(model)
			if model.OfflineBundleState != testCase.wantState {
				t.Fatalf("expected offline state %q, got %#v", testCase.wantState, model)
			}
			if model.OfflineBundleState == Point12ValCOfflineBundleStateActive || model.CurrentState == Point12ValCStateActive {
				t.Fatalf("expected mutated exact_required offline field to avoid full-ready state, got %#v", model)
			}
		})
	}

	t.Run("export self-consistency bypass cannot hide upstream drift", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.PolicyHash = "sha256:dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd"
		model.ExportBundle.ExportState = Point12ValCExportStateReady
		model = ComputePoint12ValCFoundation(model)
		if model.ExportState == Point12ValCExportStateReady || model.ExportState == Point12ValCExportStateProjectionOnly || model.CurrentState == Point12ValCStateActive {
			t.Fatalf("expected export self-consistency bypass to fail closed, got %#v", model)
		}
	})

	t.Run("offline self-consistency bypass cannot hide upstream drift", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.OfflineBundle.EvidenceHashRefs = []string{"sha256:eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"}
		model.OfflineBundle.OfflineState = Point12ValCOfflineBundleStateActive
		model = ComputePoint12ValCFoundation(model)
		if model.OfflineBundleState == Point12ValCOfflineBundleStateActive || model.CurrentState == Point12ValCStateActive {
			t.Fatalf("expected offline self-consistency bypass to fail closed, got %#v", model)
		}
	})
}

func TestPoint12ValCNoOverclaimAndTaxonomy(t *testing.T) {
	t.Run("forbidden wording in export output blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.ExportOutputClaims = []string{"production approved"}
		model = ComputePoint12ValCFoundation(model)
		if model.ExportState != Point12ValCExportStateBlocked {
			t.Fatalf("expected export overclaim to block, got %#v", model)
		}
	})

	t.Run("forbidden wording in customer visible summary blocks", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.CustomerVisibleSummary = "compliance guaranteed"
		model = ComputePoint12ValCFoundation(model)
		if model.ExportState != Point12ValCExportStateBlocked {
			t.Fatalf("expected customer summary overclaim to block, got %#v", model)
		}
	})

	t.Run("customer facing limitations cannot overclaim", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.ExportAudience = point12ValCExportAudienceCustomer
		model.ExportBundle.Limitations = []string{"certified"}
		model = ComputePoint12ValCFoundation(model)
		if model.ExportState != Point12ValCExportStateBlocked {
			t.Fatalf("expected customer-facing limitation overclaim to block, got %#v", model)
		}
	})

	t.Run("projection only is not final pass", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.ExportBundle.ExportState = Point12ValCExportStateProjectionOnly
		model = ComputePoint12ValCFoundation(model)
		body, err := json.Marshal(model)
		if err != nil {
			t.Fatalf("marshal valc foundation: %v", err)
		}
		if strings.Contains(string(body), "point_12_pass") {
			t.Fatalf("expected projection_only to remain non-pass, got %s", body)
		}
	})

	t.Run("partial advisory export is review required not pass", func(t *testing.T) {
		model := activePoint12ValCFoundation()
		model.Dependency.ValBReplayTaxonomy = Point12Val0ReplayResultInsufficientEvidence
		model.Dependency.ValBReplayResult.ReplayResultTaxonomy = Point12Val0ReplayResultInsufficientEvidence
		model.Dependency.ValBReplayResult.InsufficientEvidence = true
		model.ExportBundle.ExportState = Point12ValCExportStatePartialAdvisory
		model.ExportBundle.Limitations = []string{"insufficient evidence for full export"}
		model = ComputePoint12ValCFoundation(model)
		if model.CurrentState != Point12ValCStateReviewRequired {
			t.Fatalf("expected partial advisory export to remain review required, got %#v", model)
		}
	})
}

func TestPoint12ValCRegressionGuards(t *testing.T) {
	t.Run("val0 computed provenance fix preserved through vala valb valc chain", func(t *testing.T) {
		valD := activePoint11ValDFoundation()
		val0 := Point12Val0FoundationModel()
		val0.Dependency = SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD, Point12Val0Point11ReviewContext{
			SnapshotFromComputedOutput: false,
		})
		val0 = ComputePoint12Val0Foundation(val0)
		valA := activePoint12ValAFoundationFromVal0(val0)
		valB := activePoint12ValBFoundationFromValA(valA)
		model := activePoint12ValCFoundationFromValB(valB)
		if model.DependencyState != Point12ValCDependencyStateBlocked {
			t.Fatalf("expected non-computed upstream provenance to stay blocked through valc, got %#v", model)
		}
	})

	t.Run("vala manifest tamper behavior preserved", func(t *testing.T) {
		valA := activePoint12ValAFoundation()
		valA.ManifestIntegrityState = Point12ValAManifestIntegrityStateTampered
		valB := activePoint12ValBFoundationFromValA(valA)
		model := activePoint12ValCFoundationFromValB(valB)
		if model.DependencyState != Point12ValCDependencyStateBlocked {
			t.Fatalf("expected tampered vala manifest to block valc dependency, got %#v", model)
		}
	})

	t.Run("vala schema hash drift with recomputed payload and signature remains blocked", func(t *testing.T) {
		valA := activePoint12ValAFoundation()
		valA.Manifest.SchemaHash = "sha256:abababababababababababababababababababababababababababababababab"
		valA.Manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(valA.Manifest)
		valA.Manifest.SignatureBoundManifestPayloadHash = valA.Manifest.ManifestPayloadHash
		valA = ComputePoint12ValAFoundation(valA)
		if valA.ManifestIntegrityState == Point12ValAManifestIntegrityStateActive {
			t.Fatalf("expected schema hash drift to stay non-active in vala, got %#v", valA)
		}
		valB := activePoint12ValBFoundationFromValA(valA)
		model := activePoint12ValCFoundationFromValB(valB)
		if model.DependencyState != Point12ValCDependencyStateBlocked {
			t.Fatalf("expected schema hash drifted vala manifest to block valc dependency, got %#v", model)
		}
	})

	t.Run("valb original context cannot silently use current policy preserved", func(t *testing.T) {
		valB := activePoint12ValBFoundation()
		valB.ReplayRequest.CurrentPolicyRef = "policy_point12_current_001"
		valB.ReplayRequest.CurrentPolicyVersion = "policy_version_point12_current_001"
		valB.ReplayRequest.CurrentPolicyHash = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
		valB = ComputePoint12ValBFoundation(valB)
		model := activePoint12ValCFoundationFromValB(valB)
		if model.DependencyState != Point12ValCDependencyStateBlocked {
			t.Fatalf("expected blocked valc dependency from invalid original_context replay, got %#v", model)
		}
	})

	t.Run("no real signing or external api side effects introduced", func(t *testing.T) {
		body := readPoint12ValCSource(t)
		for _, forbidden := range []string{
			"http.Get",
			"http.Post",
			"fetch(",
			"Sign(",
			"GenerateKey",
			"crypto/rsa",
			"crypto/ecdsa",
			"crypto/ed25519",
		} {
			if strings.Contains(body, forbidden) {
				t.Fatalf("unexpected valc source boundary violation %q", forbidden)
			}
		}
	})
}

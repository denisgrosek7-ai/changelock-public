package formal

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const (
	Point12ValAStateActive         = "point12_vala_signed_proof_pack_manifest_core_active"
	Point12ValAStateBlocked        = "point12_vala_signed_proof_pack_manifest_core_blocked"
	Point12ValAStateReviewRequired = "point12_vala_signed_proof_pack_manifest_core_review_required"

	Point12ValADependencyStateActive         = "point12_vala_dependency_active"
	Point12ValADependencyStateBlocked        = "point12_vala_dependency_blocked"
	Point12ValADependencyStateReviewRequired = "point12_vala_dependency_review_required"

	Point12ValAManifestIntegrityStateActive         = "point12_vala_manifest_integrity_active"
	Point12ValAManifestIntegrityStateBlocked        = "point12_vala_manifest_integrity_blocked"
	Point12ValAManifestIntegrityStateIncomplete     = "point12_vala_manifest_integrity_incomplete"
	Point12ValAManifestIntegrityStateUnsupported    = "point12_vala_manifest_integrity_unsupported"
	Point12ValAManifestIntegrityStateTampered       = "point12_vala_manifest_integrity_tampered"
	Point12ValAManifestIntegrityStateReviewRequired = "point12_vala_manifest_integrity_review_required"
)

const (
	point12ValAWaveID                        = "val_a"
	point12ValAPreviousWaveID                = point12Val0WaveID
	point12ValAProjectionDisclaimerBaseline  = "projection_only not_canonical_truth point12_vala_signed_proof_pack_manifest_core"
	point12ValAHashAlgorithmSHA256           = "hash_algorithm_sha256"
	point12ValASignatureAlgorithmEd25519     = "signature_algorithm_ed25519_detached_metadata"
	point12ValASignatureAlgorithmECDSAP256   = "signature_algorithm_ecdsa_p256_detached_metadata"
	point12ValASigningKeyStateActive         = "active"
	point12ValASigningKeyStateRevoked        = "revoked"
	point12ValASigningKeyStateExpired        = "expired"
	point12ValASigningKeyStateCompromised    = "compromised"
	point12ValASigningKeyStateUnknown        = "unknown"
	point12ValADependencySnapshotRefBaseline = "dependency_snapshot_point12_vala_val0_computed_001"
)

type Point12ValAVal0ReviewContext struct {
	SnapshotFromComputedOutput   bool     `json:"snapshot_from_computed_output"`
	Val0PrematurePoint12PassSeen bool     `json:"val0_premature_point12_pass_seen"`
	ReviewPrerequisites          []string `json:"review_prerequisites,omitempty"`
}

type Point12ValADependencySnapshot struct {
	Val0CurrentState              string                             `json:"val0_current_state"`
	Val0DependencyState           string                             `json:"val0_dependency_state"`
	Val0PointID                   string                             `json:"val0_point_id"`
	Val0WaveID                    string                             `json:"val0_wave_id"`
	Val0NoOverclaimState          string                             `json:"val0_no_overclaim_state"`
	Val0ManifestState             string                             `json:"val0_manifest_state"`
	Val0RedactionBoundaryState    string                             `json:"val0_redaction_boundary_state"`
	Val0CompatibilityProfileState string                             `json:"val0_compatibility_profile_state"`
	Val0ProvenanceState           string                             `json:"val0_provenance_state"`
	Val0Manifest                  Point12Val0SignedProofPackManifest `json:"val0_manifest"`
	ProjectionDisclaimer          string                             `json:"projection_disclaimer"`
	SnapshotRef                   string                             `json:"snapshot_ref"`
	SnapshotFromComputedOutput    bool                               `json:"snapshot_from_computed_output"`
	Val0PrematurePoint12PassSeen  bool                               `json:"val0_premature_point12_pass_seen"`
	ReviewPrerequisites           []string                           `json:"review_prerequisites,omitempty"`
}

type Point12ValASignedProofPackManifestCore struct {
	ProofPackID                       string   `json:"proof_pack_id"`
	ManifestID                        string   `json:"manifest_id"`
	DecisionID                        string   `json:"decision_id"`
	PointID                           string   `json:"point_id"`
	WaveID                            string   `json:"wave_id"`
	TenantScope                       string   `json:"tenant_scope"`
	ArtifactRef                       string   `json:"artifact_ref"`
	ArtifactHash                      string   `json:"artifact_hash"`
	EvidenceRefs                      []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs                  []string `json:"evidence_hash_refs,omitempty"`
	PolicyRef                         string   `json:"policy_ref"`
	PolicyVersion                     string   `json:"policy_version"`
	PolicyHash                        string   `json:"policy_hash"`
	EngineVersion                     string   `json:"engine_version"`
	EngineHash                        string   `json:"engine_hash"`
	SchemaVersion                     string   `json:"schema_version"`
	SchemaHash                        string   `json:"schema_hash"`
	ClaimRefs                         []string `json:"claim_refs,omitempty"`
	GovernanceEventRefs               []string `json:"governance_event_refs,omitempty"`
	CompatibilityProfileRef           string   `json:"compatibility_profile_ref"`
	UpstreamVal0SnapshotRef           string   `json:"upstream_val0_snapshot_ref"`
	GeneratedAt                       string   `json:"generated_at"`
	FreshnessWindow                   string   `json:"freshness_window"`
	ManifestPayloadHash               string   `json:"manifest_payload_hash"`
	ManifestSchemaVersion             string   `json:"manifest_schema_version"`
	ManifestSchemaHash                string   `json:"manifest_schema_hash"`
	HashAlgorithmRef                  string   `json:"hash_algorithm_ref"`
	SignatureAlgorithmRef             string   `json:"signature_algorithm_ref"`
	SigningKeyRef                     string   `json:"signing_key_ref"`
	SigningKeyState                   string   `json:"signing_key_state"`
	SignatureRef                      string   `json:"signature_ref"`
	DetachedSignatureRef              string   `json:"detached_signature_ref"`
	SignatureMetadataRef              string   `json:"signature_metadata_ref"`
	SignatureTimestamp                string   `json:"signature_timestamp"`
	SignatureBoundManifestID          string   `json:"signature_bound_manifest_id"`
	SignatureBoundManifestPayloadHash string   `json:"signature_bound_manifest_payload_hash"`
	RedactionManifestRef              string   `json:"redaction_manifest_ref"`
	RetentionClassRef                 string   `json:"retention_class_ref"`
	ProjectionDisclaimer              string   `json:"projection_disclaimer"`
	ToolchainProvenanceRefs           []string `json:"toolchain_provenance_refs,omitempty"`
	AgentLineageRefs                  []string `json:"agent_lineage_refs,omitempty"`
	ManifestOutputClaims              []string `json:"manifest_output_claims,omitempty"`
	ManifestState                     string   `json:"manifest_state"`
}

type Point12ValAFoundation struct {
	CurrentState           string                                 `json:"current_state"`
	BlockingReasons        []string                               `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites    []string                               `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer   string                                 `json:"projection_disclaimer"`
	DependencyState        string                                 `json:"dependency_state"`
	ManifestIntegrityState string                                 `json:"manifest_integrity_state"`
	Dependency             Point12ValADependencySnapshot          `json:"dependency"`
	Manifest               Point12ValASignedProofPackManifestCore `json:"manifest"`
}

func point12ValAManifestCoreStates() []string {
	return []string{
		Point12ValAManifestIntegrityStateActive,
		Point12ValAManifestIntegrityStateBlocked,
		Point12ValAManifestIntegrityStateIncomplete,
		Point12ValAManifestIntegrityStateUnsupported,
		Point12ValAManifestIntegrityStateTampered,
		Point12ValAManifestIntegrityStateReviewRequired,
	}
}

func point12ValADependencySnapshotRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"dependency_snapshot_", "val0_snapshot_"})
}

func point12ValAManifestRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"manifest_"})
}

func point12ValAHashAlgorithmRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"hash_algorithm_"})
}

func point12ValASignatureAlgorithmRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"signature_algorithm_"})
}

func point12ValASignatureMetadataRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"signature_metadata_"})
}

func point12ValADetachedSignatureRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"detached_signature_", "signature_"})
}

func point12ValAHashAlgorithmSupported(value string) bool {
	return strings.TrimSpace(value) == point12ValAHashAlgorithmSHA256
}

func point12ValASignatureAlgorithmSupported(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValASignatureAlgorithmEd25519,
		point12ValASignatureAlgorithmECDSAP256,
	}, value)
}

func point12ValASigningKeyStateValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValASigningKeyStateActive,
		point12ValASigningKeyStateRevoked,
		point12ValASigningKeyStateExpired,
		point12ValASigningKeyStateCompromised,
		point12ValASigningKeyStateUnknown,
	}, value)
}

func point12ValAManifestPayloadParts(model Point12ValASignedProofPackManifestCore) []string {
	parts := []string{
		strings.TrimSpace(model.ProofPackID),
		strings.TrimSpace(model.ManifestID),
		strings.TrimSpace(model.DecisionID),
		strings.TrimSpace(model.PointID),
		strings.TrimSpace(model.WaveID),
		strings.TrimSpace(model.TenantScope),
		strings.TrimSpace(model.ArtifactRef),
		strings.TrimSpace(model.ArtifactHash),
		strings.Join(model.EvidenceRefs, ","),
		strings.Join(model.EvidenceHashRefs, ","),
		strings.TrimSpace(model.PolicyRef),
		strings.TrimSpace(model.PolicyVersion),
		strings.TrimSpace(model.PolicyHash),
		strings.TrimSpace(model.EngineVersion),
		strings.TrimSpace(model.EngineHash),
		strings.TrimSpace(model.SchemaVersion),
		strings.TrimSpace(model.SchemaHash),
		strings.Join(model.ClaimRefs, ","),
		strings.Join(model.GovernanceEventRefs, ","),
		strings.TrimSpace(model.CompatibilityProfileRef),
		strings.TrimSpace(model.UpstreamVal0SnapshotRef),
		strings.TrimSpace(model.ManifestSchemaVersion),
		strings.TrimSpace(model.ManifestSchemaHash),
		strings.TrimSpace(model.RedactionManifestRef),
		strings.TrimSpace(model.RetentionClassRef),
		strings.Join(model.ToolchainProvenanceRefs, ","),
		strings.Join(model.AgentLineageRefs, ","),
		strings.Join(model.ManifestOutputClaims, ","),
	}
	return parts
}

func point12ValAComputedManifestPayloadHash(model Point12ValASignedProofPackManifestCore) string {
	if !point12ValAHashAlgorithmSupported(model.HashAlgorithmRef) {
		return ""
	}
	sum := sha256.Sum256([]byte(strings.Join(point12ValAManifestPayloadParts(model), "\n")))
	return "sha256:" + hex.EncodeToString(sum[:])
}

func SnapshotPoint12ValADependencyFromComputedVal0(val0 Point12Val0Foundation, review Point12ValAVal0ReviewContext) Point12ValADependencySnapshot {
	reviewPrerequisites := append([]string{}, val0.ReviewPrerequisites...)
	reviewPrerequisites = append(reviewPrerequisites, review.ReviewPrerequisites...)
	return Point12ValADependencySnapshot{
		Val0CurrentState:              val0.CurrentState,
		Val0DependencyState:           val0.DependencyState,
		Val0PointID:                   val0.Manifest.PointID,
		Val0WaveID:                    val0.Manifest.WaveID,
		Val0NoOverclaimState:          val0.NoOverclaimState,
		Val0ManifestState:             val0.ManifestState,
		Val0RedactionBoundaryState:    val0.RedactionBoundaryState,
		Val0CompatibilityProfileState: val0.CompatibilityProfileState,
		Val0ProvenanceState:           val0.ProvenanceState,
		Val0Manifest:                  val0.Manifest,
		ProjectionDisclaimer:          val0.ProjectionDisclaimer,
		SnapshotRef:                   point12ValADependencySnapshotRefBaseline,
		SnapshotFromComputedOutput:    review.SnapshotFromComputedOutput,
		Val0PrematurePoint12PassSeen:  review.Val0PrematurePoint12PassSeen,
		ReviewPrerequisites:           reviewPrerequisites,
	}
}

func point12ValADependencyReviewContextModel() Point12ValAVal0ReviewContext {
	return Point12ValAVal0ReviewContext{
		SnapshotFromComputedOutput: true,
	}
}

func point12ValADependencySnapshotModel() Point12ValADependencySnapshot {
	val0 := ComputePoint12Val0Foundation(Point12Val0FoundationModel())
	return SnapshotPoint12ValADependencyFromComputedVal0(val0, point12ValADependencyReviewContextModel())
}

func EvaluatePoint12ValADependencyState(model Point12ValADependencySnapshot) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!model.SnapshotFromComputedOutput ||
		!point12ValADependencySnapshotRefValid(model.SnapshotRef) ||
		strings.TrimSpace(model.Val0PointID) != point12Val0PointID ||
		strings.TrimSpace(model.Val0WaveID) != point12Val0WaveID ||
		strings.TrimSpace(model.Val0Manifest.PointID) != point12Val0PointID ||
		strings.TrimSpace(model.Val0Manifest.WaveID) != point12Val0WaveID ||
		model.Val0PrematurePoint12PassSeen ||
		point12Val0ContainsPrematurePassToken(
			model.Val0Manifest.ProofPackID,
			model.Val0Manifest.DecisionID,
			model.Val0Manifest.SigningKeyRef,
			model.Val0Manifest.SignatureRef,
		) {
		return Point12ValADependencyStateBlocked
	}
	if strings.TrimSpace(model.Val0CurrentState) == Point12Val0StateBlocked ||
		strings.TrimSpace(model.Val0DependencyState) == Point12Val0DependencyStateBlocked ||
		strings.TrimSpace(model.Val0NoOverclaimState) != Point12Val0NoOverclaimStateActive ||
		strings.TrimSpace(model.Val0ManifestState) != Point12Val0ManifestStateActive ||
		strings.TrimSpace(model.Val0RedactionBoundaryState) != Point12Val0RedactionBoundaryStateActive ||
		strings.TrimSpace(model.Val0CompatibilityProfileState) != Point12Val0CompatibilityProfileStateActive ||
		strings.TrimSpace(model.Val0ProvenanceState) == Point12Val0ProvenanceStateBlocked {
		return Point12ValADependencyStateBlocked
	}
	if strings.TrimSpace(model.Val0CurrentState) == Point12Val0StateReviewRequired ||
		strings.TrimSpace(model.Val0DependencyState) == Point12Val0DependencyStateReviewRequired ||
		strings.TrimSpace(model.Val0ProvenanceState) == Point12Val0ProvenanceStateReviewRequired ||
		len(model.ReviewPrerequisites) > 0 {
		return Point12ValADependencyStateReviewRequired
	}
	if strings.TrimSpace(model.Val0CurrentState) != Point12Val0StateActive ||
		strings.TrimSpace(model.Val0DependencyState) != Point12Val0DependencyStateActive ||
		strings.TrimSpace(model.Val0ProvenanceState) != Point12Val0ProvenanceStateActive {
		return Point12ValADependencyStateBlocked
	}
	return Point12ValADependencyStateActive
}

func point12ValAManifestIntegrityStateAndReasons(
	model Point12ValASignedProofPackManifestCore,
	dependency Point12ValADependencySnapshot,
) (string, []string) {
	blockedReasons := []string{}
	tamperedReasons := []string{}
	unsupportedReasons := []string{}
	reviewReasons := []string{}

	if !point11Val0ContainsTrimmed(point12ValAManifestCoreStates(), model.ManifestState) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point12Val0DecisionRefValid(model.DecisionID) ||
		strings.TrimSpace(model.PointID) != point12Val0PointID ||
		strings.TrimSpace(model.WaveID) != point12ValAWaveID ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12Val0ArtifactRefValid(model.ArtifactRef) ||
		!point12Val0HashValid(model.ArtifactHash) ||
		!point12Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point12Val0StringListValid(model.EvidenceHashRefs, point12Val0EvidenceHashRefValid) ||
		!point12Val0PolicyRefValid(model.PolicyRef) ||
		!point12Val0VersionIdentityValid(model.PolicyVersion) ||
		!point12Val0HashValid(model.PolicyHash) ||
		!point12Val0VersionIdentityValid(model.EngineVersion) ||
		!point12Val0HashValid(model.EngineHash) ||
		!point12Val0VersionIdentityValid(model.SchemaVersion) ||
		!point12Val0HashValid(model.SchemaHash) ||
		!point12Val0StringListValid(model.ClaimRefs, point12Val0ClaimRefValid) ||
		!point12Val0StringListValid(model.GovernanceEventRefs, point12Val0GovernanceEventRefValid) ||
		!point12Val0CompatibilityProfileRefValid(model.CompatibilityProfileRef) ||
		!point12ValADependencySnapshotRefValid(model.UpstreamVal0SnapshotRef) ||
		!point11Val0ValidTimestamp(model.GeneratedAt) ||
		!point12Val0VersionIdentityValid(model.FreshnessWindow) ||
		!point12Val0HashValid(model.ManifestPayloadHash) ||
		!point12Val0VersionIdentityValid(model.ManifestSchemaVersion) ||
		!point12Val0HashValid(model.ManifestSchemaHash) ||
		!point12ValAHashAlgorithmRefValid(model.HashAlgorithmRef) ||
		!point12ValASignatureAlgorithmRefValid(model.SignatureAlgorithmRef) ||
		!point12Val0SigningKeyRefValid(model.SigningKeyRef) ||
		!point12ValASigningKeyStateValid(model.SigningKeyState) ||
		!point12ValASignatureMetadataRefValid(model.SignatureMetadataRef) ||
		!point11Val0ValidTimestamp(model.SignatureTimestamp) ||
		!point12Val0RedactionManifestRefValid(model.RedactionManifestRef) ||
		!point12Val0RetentionClassRefValid(model.RetentionClassRef) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12Val0StringListValid(model.ToolchainProvenanceRefs, point12Val0ToolchainProvenanceRefValid) ||
		!point12Val0StringListValid(model.AgentLineageRefs, point12Val0AgentLineageRefValid) ||
		!point12Val0OptionalClaimTextListValid(model.ManifestOutputClaims) {
		blockedReasons = append(blockedReasons, "manifest_identity_or_metadata_invalid")
	}

	if len(model.EvidenceRefs) != len(model.EvidenceHashRefs) {
		blockedReasons = append(blockedReasons, "manifest_evidence_hash_alignment_invalid")
	}
	if strings.TrimSpace(model.SignatureRef) == "" && strings.TrimSpace(model.DetachedSignatureRef) == "" {
		blockedReasons = append(blockedReasons, "manifest_signature_reference_missing")
	}
	if strings.TrimSpace(model.SignatureRef) != "" && !point12Val0SignatureRefValid(model.SignatureRef) {
		blockedReasons = append(blockedReasons, "manifest_signature_ref_invalid")
	}
	if strings.TrimSpace(model.DetachedSignatureRef) != "" && !point12ValADetachedSignatureRefValid(model.DetachedSignatureRef) {
		blockedReasons = append(blockedReasons, "manifest_detached_signature_ref_invalid")
	}
	if point12Val0ContainsPrematurePassToken(
		model.ProofPackID,
		model.ManifestID,
		model.DecisionID,
		model.SigningKeyRef,
		model.SignatureRef,
		model.DetachedSignatureRef,
		model.SignatureMetadataRef,
		strings.Join(model.ManifestOutputClaims, " "),
	) {
		blockedReasons = append(blockedReasons, "manifest_premature_point12_pass")
	}
	if point12Val0ContainsForbiddenClaim(strings.Join(model.ManifestOutputClaims, " ")) {
		blockedReasons = append(blockedReasons, "manifest_output_overclaim_detected")
	}

	if point12ValAHashAlgorithmRefValid(model.HashAlgorithmRef) && !point12ValAHashAlgorithmSupported(model.HashAlgorithmRef) {
		unsupportedReasons = append(unsupportedReasons, "manifest_hash_algorithm_unsupported")
	}
	if point12ValASignatureAlgorithmRefValid(model.SignatureAlgorithmRef) && !point12ValASignatureAlgorithmSupported(model.SignatureAlgorithmRef) {
		unsupportedReasons = append(unsupportedReasons, "manifest_signature_algorithm_unsupported")
	}

	switch strings.TrimSpace(model.SigningKeyState) {
	case point12ValASigningKeyStateRevoked:
		blockedReasons = append(blockedReasons, "manifest_signing_key_revoked")
	case point12ValASigningKeyStateExpired:
		blockedReasons = append(blockedReasons, "manifest_signing_key_expired")
	case point12ValASigningKeyStateCompromised:
		blockedReasons = append(blockedReasons, "manifest_signing_key_compromised")
	case point12ValASigningKeyStateUnknown:
		reviewReasons = append(reviewReasons, "manifest_signing_key_unknown")
	}

	if strings.TrimSpace(model.UpstreamVal0SnapshotRef) != strings.TrimSpace(dependency.SnapshotRef) {
		tamperedReasons = append(tamperedReasons, "manifest_snapshot_binding_mismatch")
	}
	if strings.TrimSpace(model.DecisionID) != strings.TrimSpace(dependency.Val0Manifest.DecisionID) {
		tamperedReasons = append(tamperedReasons, "manifest_decision_binding_mismatch")
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.Val0Manifest.TenantScope) {
		tamperedReasons = append(tamperedReasons, "manifest_tenant_scope_binding_mismatch")
	}
	if strings.TrimSpace(model.ArtifactRef) != strings.TrimSpace(dependency.Val0Manifest.ArtifactRef) ||
		strings.TrimSpace(model.ArtifactHash) != strings.TrimSpace(dependency.Val0Manifest.ArtifactHash) {
		tamperedReasons = append(tamperedReasons, "manifest_artifact_binding_mismatch")
	}
	if !point12Val0ExactStringSetMatch(model.EvidenceRefs, dependency.Val0Manifest.EvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.EvidenceHashRefs, dependency.Val0Manifest.EvidenceHashRefs) {
		tamperedReasons = append(tamperedReasons, "manifest_evidence_binding_mismatch")
	}
	if strings.TrimSpace(model.PolicyRef) != strings.TrimSpace(dependency.Val0Manifest.PolicyRef) ||
		strings.TrimSpace(model.PolicyVersion) != strings.TrimSpace(dependency.Val0Manifest.PolicyVersion) ||
		strings.TrimSpace(model.PolicyHash) != strings.TrimSpace(dependency.Val0Manifest.PolicyHash) {
		tamperedReasons = append(tamperedReasons, "manifest_policy_binding_mismatch")
	}
	if strings.TrimSpace(model.EngineVersion) != strings.TrimSpace(dependency.Val0Manifest.EngineVersion) ||
		strings.TrimSpace(model.EngineHash) != strings.TrimSpace(dependency.Val0Manifest.EngineHash) {
		tamperedReasons = append(tamperedReasons, "manifest_engine_binding_mismatch")
	}
	if strings.TrimSpace(model.SchemaVersion) != strings.TrimSpace(dependency.Val0Manifest.SchemaVersion) {
		tamperedReasons = append(tamperedReasons, "manifest_schema_binding_mismatch")
	}
	if !point12Val0ExactStringSetMatch(model.ClaimRefs, dependency.Val0Manifest.ClaimRefs) {
		tamperedReasons = append(tamperedReasons, "manifest_claim_binding_mismatch")
	}
	if !point12Val0ExactStringSetMatch(model.GovernanceEventRefs, dependency.Val0Manifest.GovernanceEventRefs) {
		tamperedReasons = append(tamperedReasons, "manifest_governance_binding_mismatch")
	}
	if strings.TrimSpace(model.CompatibilityProfileRef) != strings.TrimSpace(dependency.Val0Manifest.CompatibilityProfileRef) {
		tamperedReasons = append(tamperedReasons, "manifest_compatibility_profile_binding_mismatch")
	}
	if strings.TrimSpace(model.RedactionManifestRef) != strings.TrimSpace(dependency.Val0Manifest.RedactionManifestRef) {
		tamperedReasons = append(tamperedReasons, "manifest_redaction_binding_mismatch")
	}
	if strings.TrimSpace(model.RetentionClassRef) != strings.TrimSpace(dependency.Val0Manifest.RetentionClassRef) {
		tamperedReasons = append(tamperedReasons, "manifest_retention_binding_mismatch")
	}
	if !point12Val0ExactStringSetMatch(model.ToolchainProvenanceRefs, dependency.Val0Manifest.ToolchainProvenanceRefs) {
		tamperedReasons = append(tamperedReasons, "manifest_toolchain_binding_mismatch")
	}
	if !point12Val0ExactStringSetMatch(model.AgentLineageRefs, dependency.Val0Manifest.AgentLineageRefs) {
		tamperedReasons = append(tamperedReasons, "manifest_agent_lineage_binding_mismatch")
	}
	if strings.TrimSpace(model.PolicyRef) != strings.TrimSpace(dependency.Val0Manifest.PolicyAuthorityContextRef) &&
		strings.TrimSpace(model.PolicyRef) != strings.TrimSpace(dependency.Val0Manifest.PolicyRef) {
		tamperedReasons = append(tamperedReasons, "manifest_policy_authority_binding_mismatch")
	}

	expectedPayloadHash := point12ValAComputedManifestPayloadHash(model)
	if expectedPayloadHash == "" {
		unsupportedReasons = append(unsupportedReasons, "manifest_payload_hash_algorithm_unsupported")
	} else if strings.TrimSpace(model.ManifestPayloadHash) != expectedPayloadHash {
		tamperedReasons = append(tamperedReasons, "manifest_payload_hash_mismatch")
	}
	if strings.TrimSpace(model.SignatureBoundManifestID) != strings.TrimSpace(model.ManifestID) {
		tamperedReasons = append(tamperedReasons, "manifest_signature_manifest_id_binding_mismatch")
	}
	if strings.TrimSpace(model.SignatureBoundManifestPayloadHash) != strings.TrimSpace(model.ManifestPayloadHash) {
		tamperedReasons = append(tamperedReasons, "manifest_signature_payload_hash_binding_mismatch")
	}

	if len(blockedReasons) > 0 {
		return Point12ValAManifestIntegrityStateBlocked, blockedReasons
	}
	if len(unsupportedReasons) > 0 {
		return Point12ValAManifestIntegrityStateUnsupported, unsupportedReasons
	}
	if len(tamperedReasons) > 0 {
		return Point12ValAManifestIntegrityStateTampered, tamperedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValAManifestIntegrityStateReviewRequired, reviewReasons
	}
	return Point12ValAManifestIntegrityStateActive, nil
}

func EvaluatePoint12ValAManifestIntegrityState(
	model Point12ValASignedProofPackManifestCore,
	dependency Point12ValADependencySnapshot,
) string {
	state, _ := point12ValAManifestIntegrityStateAndReasons(model, dependency)
	return state
}

func EvaluatePoint12ValAState(model Point12ValAFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		strings.TrimSpace(model.DependencyState) == Point12ValADependencyStateBlocked ||
		strings.TrimSpace(model.ManifestIntegrityState) == Point12ValAManifestIntegrityStateBlocked ||
		strings.TrimSpace(model.ManifestIntegrityState) == Point12ValAManifestIntegrityStateIncomplete ||
		strings.TrimSpace(model.ManifestIntegrityState) == Point12ValAManifestIntegrityStateUnsupported ||
		strings.TrimSpace(model.ManifestIntegrityState) == Point12ValAManifestIntegrityStateTampered {
		return Point12ValAStateBlocked
	}
	if strings.TrimSpace(model.DependencyState) == Point12ValADependencyStateReviewRequired ||
		strings.TrimSpace(model.ManifestIntegrityState) == Point12ValAManifestIntegrityStateReviewRequired {
		return Point12ValAStateReviewRequired
	}
	return Point12ValAStateActive
}

func point12ValABlockingReasons(model Point12ValAFoundation) []string {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "aggregate_projection_disclaimer_blocked")
	}
	if model.DependencyState == Point12ValADependencyStateBlocked {
		reasons = append(reasons, "point12_val0_dependency_blocked")
	}
	switch model.ManifestIntegrityState {
	case Point12ValAManifestIntegrityStateBlocked:
		reasons = append(reasons, "manifest_integrity_blocked")
	case Point12ValAManifestIntegrityStateIncomplete:
		reasons = append(reasons, "manifest_integrity_incomplete")
	case Point12ValAManifestIntegrityStateUnsupported:
		reasons = append(reasons, "manifest_integrity_unsupported")
	case Point12ValAManifestIntegrityStateTampered:
		reasons = append(reasons, "manifest_integrity_tampered")
	}
	return reasons
}

func Point12ValAFoundationModel() Point12ValAFoundation {
	dependency := point12ValADependencySnapshotModel()
	manifest := Point12ValASignedProofPackManifestCore{
		ProofPackID:              dependency.Val0Manifest.ProofPackID,
		ManifestID:               "manifest_point12_vala_001",
		DecisionID:               dependency.Val0Manifest.DecisionID,
		PointID:                  point12Val0PointID,
		WaveID:                   point12ValAWaveID,
		TenantScope:              dependency.Val0Manifest.TenantScope,
		ArtifactRef:              dependency.Val0Manifest.ArtifactRef,
		ArtifactHash:             dependency.Val0Manifest.ArtifactHash,
		EvidenceRefs:             append([]string{}, dependency.Val0Manifest.EvidenceRefs...),
		EvidenceHashRefs:         append([]string{}, dependency.Val0Manifest.EvidenceHashRefs...),
		PolicyRef:                dependency.Val0Manifest.PolicyRef,
		PolicyVersion:            dependency.Val0Manifest.PolicyVersion,
		PolicyHash:               dependency.Val0Manifest.PolicyHash,
		EngineVersion:            dependency.Val0Manifest.EngineVersion,
		EngineHash:               dependency.Val0Manifest.EngineHash,
		SchemaVersion:            dependency.Val0Manifest.SchemaVersion,
		SchemaHash:               "sha256:7777777777777777777777777777777777777777777777777777777777777777",
		ClaimRefs:                append([]string{}, dependency.Val0Manifest.ClaimRefs...),
		GovernanceEventRefs:      append([]string{}, dependency.Val0Manifest.GovernanceEventRefs...),
		CompatibilityProfileRef:  dependency.Val0Manifest.CompatibilityProfileRef,
		UpstreamVal0SnapshotRef:  dependency.SnapshotRef,
		GeneratedAt:              "2026-05-03T12:00:00Z",
		FreshnessWindow:          dependency.Val0Manifest.FreshnessWindow,
		ManifestSchemaVersion:    "manifest_schema_version_point12_vala_v1",
		ManifestSchemaHash:       "sha256:8888888888888888888888888888888888888888888888888888888888888888",
		HashAlgorithmRef:         point12ValAHashAlgorithmSHA256,
		SignatureAlgorithmRef:    point12ValASignatureAlgorithmEd25519,
		SigningKeyRef:            dependency.Val0Manifest.SigningKeyRef,
		SigningKeyState:          point12ValASigningKeyStateActive,
		SignatureRef:             dependency.Val0Manifest.SignatureRef,
		DetachedSignatureRef:     "detached_signature_point12_vala_001",
		SignatureMetadataRef:     "signature_metadata_point12_vala_001",
		SignatureTimestamp:       "2026-05-03T12:05:00Z",
		SignatureBoundManifestID: "manifest_point12_vala_001",
		RedactionManifestRef:     dependency.Val0Manifest.RedactionManifestRef,
		RetentionClassRef:        dependency.Val0Manifest.RetentionClassRef,
		ProjectionDisclaimer:     point12ValAProjectionDisclaimerBaseline,
		ToolchainProvenanceRefs:  append([]string{}, dependency.Val0Manifest.ToolchainProvenanceRefs...),
		AgentLineageRefs:         append([]string{}, dependency.Val0Manifest.AgentLineageRefs...),
		ManifestOutputClaims:     []string{"bounded claim"},
		ManifestState:            Point12ValAManifestIntegrityStateActive,
	}
	manifest.ManifestPayloadHash = point12ValAComputedManifestPayloadHash(manifest)
	manifest.SignatureBoundManifestPayloadHash = manifest.ManifestPayloadHash
	return Point12ValAFoundation{
		CurrentState:           Point12ValAStateActive,
		ProjectionDisclaimer:   point12ValAProjectionDisclaimerBaseline,
		DependencyState:        Point12ValADependencyStateActive,
		ManifestIntegrityState: Point12ValAManifestIntegrityStateActive,
		Dependency:             dependency,
		Manifest:               manifest,
	}
}

func ComputePoint12ValAFoundation(model Point12ValAFoundation) Point12ValAFoundation {
	model.DependencyState = EvaluatePoint12ValADependencyState(model.Dependency)
	manifestState, manifestReasons := point12ValAManifestIntegrityStateAndReasons(model.Manifest, model.Dependency)
	model.ManifestIntegrityState = manifestState
	model.Manifest.ManifestState = manifestState
	model.CurrentState = EvaluatePoint12ValAState(model)
	model.BlockingReasons = point12ValABlockingReasons(model)
	model.ReviewPrerequisites = append([]string{}, model.Dependency.ReviewPrerequisites...)
	if manifestState == Point12ValAManifestIntegrityStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, manifestReasons...)
	}
	return model
}

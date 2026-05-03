package formal

import "strings"

const (
	Point12ValBStateActive         = "point12_valb_replay_engine_cli_semantics_active"
	Point12ValBStateBlocked        = "point12_valb_replay_engine_cli_semantics_blocked"
	Point12ValBStateReviewRequired = "point12_valb_replay_engine_cli_semantics_review_required"

	Point12ValBDependencyStateActive         = "point12_valb_dependency_active"
	Point12ValBDependencyStateBlocked        = "point12_valb_dependency_blocked"
	Point12ValBDependencyStateReviewRequired = "point12_valb_dependency_review_required"

	Point12ValBReplayCommandStateActive  = "point12_valb_replay_command_active"
	Point12ValBReplayCommandStateBlocked = "point12_valb_replay_command_blocked"

	Point12ValBReplayRequestStateActive  = "point12_valb_replay_request_active"
	Point12ValBReplayRequestStateBlocked = "point12_valb_replay_request_blocked"

	Point12ValBReplayResultStateActive         = "point12_valb_replay_result_active"
	Point12ValBReplayResultStateBlocked        = "point12_valb_replay_result_blocked"
	Point12ValBReplayResultStateReviewRequired = "point12_valb_replay_result_review_required"
)

const (
	point12ValBWaveID                        = "val_b"
	point12ValBPreviousWaveID                = point12ValAWaveID
	point12ValBProjectionDisclaimerBaseline  = "projection_only not_canonical_truth point12_valb_replay_engine_cli_semantics"
	point12ValBDependencySnapshotRefBaseline = "dependency_snapshot_point12_valb_vala_computed_001"

	point12ValBCommandKindReplayProofPack        = "replay_proof_pack"
	point12ValBCommandKindVerifyManifestContext  = "verify_manifest_context"
	point12ValBCommandKindExplainMismatch        = "explain_mismatch"
	point12ValBCommandKindExplainMissingEvidence = "explain_missing_evidence"
	point12ValBCommandKindExplainUnsupported     = "explain_unsupported_version"

	point12ValBCheckResultActive      = "active"
	point12ValBCheckResultMismatch    = "mismatch"
	point12ValBCheckResultTampered    = "tampered"
	point12ValBCheckResultUnsupported = "unsupported"
	point12ValBCheckResultMissing     = "missing"
	point12ValBCheckResultBlocked     = "blocked"

	point12ValBMismatchTypePolicyMismatch            = Point12Val0ReplayResultPolicyMismatch
	point12ValBMismatchTypeEngineMismatch            = Point12Val0ReplayResultEngineMismatch
	point12ValBMismatchTypeSchemaMismatch            = Point12Val0ReplayResultSchemaMismatch
	point12ValBMismatchTypeEvidenceMismatch          = Point12Val0ReplayResultEvidenceMismatch
	point12ValBMismatchTypeClaimMismatch             = Point12Val0ReplayResultClaimMismatch
	point12ValBMismatchTypeGovernanceMismatch        = Point12Val0ReplayResultGovernanceMismatch
	point12ValBMismatchTypeTenantScopeMismatch       = "tenant_scope_mismatch"
	point12ValBMismatchTypeArtifactMismatch          = "artifact_mismatch"
	point12ValBMismatchTypeManifestPayloadMismatch   = "manifest_payload_hash_mismatch"
	point12ValBMismatchTypeSignatureMetadataMismatch = "signature_metadata_mismatch"
	point12ValBMismatchTypeRedactionMismatch         = "redaction_mismatch"
	point12ValBMismatchTypeUnsupportedVersion        = Point12Val0ReplayResultUnsupportedVersion
	point12ValBMismatchTypeTamperDetected            = Point12Val0ReplayResultTamperDetected
	point12ValBMismatchTypeMissingEvidence           = "missing_evidence"
	point12ValBMismatchTypeRevokedExpiredEvidence    = "revoked_or_expired_evidence"
	point12ValBMismatchTypeSupersededPolicyOrClaim   = "superseded_policy_or_claim"

	point12ValBDriftDueToPolicy              = "changed_due_to_policy"
	point12ValBDriftDueToEngine              = "changed_due_to_engine"
	point12ValBDriftDueToSchema              = "changed_due_to_schema"
	point12ValBDriftDueToEvidence            = "changed_due_to_evidence"
	point12ValBDriftDueToClaim               = "changed_due_to_claim"
	point12ValBDriftDueToGovernance          = "changed_due_to_governance"
	point12ValBDriftDueToRedaction           = "changed_due_to_redaction"
	point12ValBDriftDueToRevocation          = "changed_due_to_revocation"
	point12ValBDriftDueToToolchainProvenance = "changed_due_to_toolchain_provenance"
	point12ValBDriftDueToTenantScope         = "changed_due_to_tenant_scope"
)

type Point12ValBValAReviewContext struct {
	SnapshotFromComputedOutput   bool     `json:"snapshot_from_computed_output"`
	ValAPrematurePoint12PassSeen bool     `json:"vala_premature_point12_pass_seen"`
	ReviewPrerequisites          []string `json:"review_prerequisites,omitempty"`
}

type Point12ValBDependencySnapshot struct {
	ValACurrentState             string                                 `json:"vala_current_state"`
	ValADependencyState          string                                 `json:"vala_dependency_state"`
	ValAManifestIntegrityState   string                                 `json:"vala_manifest_integrity_state"`
	ValAPointID                  string                                 `json:"vala_point_id"`
	ValAWaveID                   string                                 `json:"vala_wave_id"`
	Val0RedactionBoundaryState   string                                 `json:"val0_redaction_boundary_state"`
	ProjectionDisclaimer         string                                 `json:"projection_disclaimer"`
	SnapshotRef                  string                                 `json:"snapshot_ref"`
	SnapshotFromComputedOutput   bool                                   `json:"snapshot_from_computed_output"`
	ValAPrematurePoint12PassSeen bool                                   `json:"vala_premature_point12_pass_seen"`
	ReviewPrerequisites          []string                               `json:"review_prerequisites,omitempty"`
	ValAManifest                 Point12ValASignedProofPackManifestCore `json:"vala_manifest"`
}

type Point12ValBReplayCommandContract struct {
	CommandID                   string `json:"command_id"`
	CommandName                 string `json:"command_name"`
	CommandKind                 string `json:"command_kind"`
	ProofPackID                 string `json:"proof_pack_id"`
	ManifestID                  string `json:"manifest_id"`
	TenantScope                 string `json:"tenant_scope"`
	ArtifactRef                 string `json:"artifact_ref"`
	ReplayMode                  string `json:"replay_mode"`
	RequestedPolicyContextRef   string `json:"requested_policy_context_ref"`
	RequestedEngineContextRef   string `json:"requested_engine_context_ref"`
	RequestedSchemaContextRef   string `json:"requested_schema_context_ref"`
	CompatibilityProfileRef     string `json:"compatibility_profile_ref"`
	AllowCurrentPolicy          bool   `json:"allow_current_policy"`
	AllowExternalAPI            bool   `json:"allow_external_api"`
	OfflineBundleRequired       bool   `json:"offline_bundle_required"`
	ExplainMismatch             bool   `json:"explain_mismatch"`
	ExplainMissingEvidence      bool   `json:"explain_missing_evidence"`
	ExplainUnsupportedVersion   bool   `json:"explain_unsupported_version"`
	ExplainRedactionLimitations bool   `json:"explain_redaction_limitations"`
	GeneratedAt                 string `json:"generated_at"`
	ProjectionDisclaimer        string `json:"projection_disclaimer"`
	MutatesEvidenceSpine        bool   `json:"mutates_evidence_spine"`
	CreatesProofPack            bool   `json:"creates_proof_pack"`
	CreatesAuditExportBundle    bool   `json:"creates_audit_export_bundle"`
	OpensPortalPath             bool   `json:"opens_portal_path"`
	RequestsPoint12Pass         bool   `json:"requests_point12_pass"`
}

type Point12ValBReplayRequest struct {
	ReplayRequestID              string   `json:"replay_request_id"`
	ProofPackID                  string   `json:"proof_pack_id"`
	ManifestID                   string   `json:"manifest_id"`
	DecisionID                   string   `json:"decision_id"`
	TenantScope                  string   `json:"tenant_scope"`
	ArtifactRef                  string   `json:"artifact_ref"`
	ArtifactHash                 string   `json:"artifact_hash"`
	EvidenceRefs                 []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs             []string `json:"evidence_hash_refs,omitempty"`
	PolicyRef                    string   `json:"policy_ref"`
	PolicyVersion                string   `json:"policy_version"`
	PolicyHash                   string   `json:"policy_hash"`
	EngineVersion                string   `json:"engine_version"`
	EngineHash                   string   `json:"engine_hash"`
	SchemaVersion                string   `json:"schema_version"`
	SchemaHash                   string   `json:"schema_hash"`
	ClaimRefs                    []string `json:"claim_refs,omitempty"`
	GovernanceEventRefs          []string `json:"governance_event_refs,omitempty"`
	ManifestPayloadHash          string   `json:"manifest_payload_hash"`
	CompatibilityProfileRef      string   `json:"compatibility_profile_ref"`
	ReplayMode                   string   `json:"replay_mode"`
	DeclaredCompatibilityContext string   `json:"declared_compatibility_context"`
	OriginalDecisionState        string   `json:"original_decision_state"`
	OriginalDecisionHash         string   `json:"original_decision_hash"`
	CurrentPolicyRef             string   `json:"current_policy_ref"`
	CurrentPolicyVersion         string   `json:"current_policy_version"`
	CurrentPolicyHash            string   `json:"current_policy_hash"`
	CurrentEngineVersion         string   `json:"current_engine_version"`
	CurrentEngineHash            string   `json:"current_engine_hash"`
	CurrentSchemaVersion         string   `json:"current_schema_version"`
	CurrentSchemaHash            string   `json:"current_schema_hash"`
	CurrentEvidenceRefs          []string `json:"current_evidence_refs,omitempty"`
	CurrentEvidenceHashRefs      []string `json:"current_evidence_hash_refs,omitempty"`
	CurrentClaimRefs             []string `json:"current_claim_refs,omitempty"`
	CurrentGovernanceEventRefs   []string `json:"current_governance_event_refs,omitempty"`
	RedactionManifestRef         string   `json:"redaction_manifest_ref"`
	GeneratedAt                  string   `json:"generated_at"`
	FreshnessContext             string   `json:"freshness_context"`
	SourceManifestIntegrityState string   `json:"source_manifest_integrity_state"`
}

type Point12ValBReplayMismatch struct {
	MismatchID      string `json:"mismatch_id"`
	MismatchType    string `json:"mismatch_type"`
	ExpectedRef     string `json:"expected_ref"`
	ActualRef       string `json:"actual_ref"`
	ExpectedHash    string `json:"expected_hash"`
	ActualHash      string `json:"actual_hash"`
	ExpectedVersion string `json:"expected_version"`
	ActualVersion   string `json:"actual_version"`
	AffectedSurface string `json:"affected_surface"`
	Decisive        bool   `json:"decisive"`
	DriftReason     string `json:"drift_reason"`
	Explanation     string `json:"explanation"`
	BlocksReplay    bool   `json:"blocks_replay"`
}

type Point12ValBReplayResult struct {
	ReplayResultID               string                      `json:"replay_result_id"`
	ReplayRequestID              string                      `json:"replay_request_id"`
	ProofPackID                  string                      `json:"proof_pack_id"`
	ManifestID                   string                      `json:"manifest_id"`
	ReplayMode                   string                      `json:"replay_mode"`
	ReplayState                  string                      `json:"replay_state"`
	ReplayResultTaxonomy         string                      `json:"replay_result_taxonomy"`
	OriginalDecisionState        string                      `json:"original_decision_state"`
	ReplayedDecisionState        string                      `json:"replayed_decision_state"`
	MatchOriginal                bool                        `json:"match_original"`
	TamperDetected               bool                        `json:"tamper_detected"`
	UnsupportedVersion           bool                        `json:"unsupported_version"`
	InsufficientEvidence         bool                        `json:"insufficient_evidence"`
	RedactionLimitations         bool                        `json:"redaction_limitations"`
	Mismatches                   []Point12ValBReplayMismatch `json:"mismatches,omitempty"`
	MismatchExplanations         []string                    `json:"mismatch_explanations,omitempty"`
	DecisionDriftExplanation     string                      `json:"decision_drift_explanation"`
	DecisionDriftClassification  string                      `json:"decision_drift_classification"`
	EvaluatedPolicyVersion       string                      `json:"evaluated_policy_version"`
	EvaluatedEngineVersion       string                      `json:"evaluated_engine_version"`
	EvaluatedSchemaVersion       string                      `json:"evaluated_schema_version"`
	EvidenceHashCheckResult      string                      `json:"evidence_hash_check_result"`
	ManifestIntegrityCheckResult string                      `json:"manifest_integrity_check_result"`
	SignatureMetadataCheckResult string                      `json:"signature_metadata_check_result"`
	CompatibilityCheckResult     string                      `json:"compatibility_check_result"`
	ExternalAPIUsed              bool                        `json:"external_api_used"`
	PointPassEmitted             bool                        `json:"point_pass_emitted"`
	ProjectionDisclaimer         string                      `json:"projection_disclaimer"`
	ReplayOutputClaims           []string                    `json:"replay_output_claims,omitempty"`
	CustomerVisibleExplanation   string                      `json:"customer_visible_explanation"`
	BlockedReason                string                      `json:"blocked_reason"`
}

type Point12ValBFoundation struct {
	CurrentState         string                           `json:"current_state"`
	BlockingReasons      []string                         `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites  []string                         `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer string                           `json:"projection_disclaimer"`
	DependencyState      string                           `json:"dependency_state"`
	ReplayCommandState   string                           `json:"replay_command_state"`
	ReplayRequestState   string                           `json:"replay_request_state"`
	ReplayResultState    string                           `json:"replay_result_state"`
	Dependency           Point12ValBDependencySnapshot    `json:"dependency"`
	ReplayCommand        Point12ValBReplayCommandContract `json:"replay_command"`
	ReplayRequest        Point12ValBReplayRequest         `json:"replay_request"`
	ReplayResult         Point12ValBReplayResult          `json:"replay_result"`
}

func point12ValBDependencySnapshotRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"dependency_snapshot_", "vala_snapshot_"})
}

func point12ValBCommandRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"command_", "replay_command_"})
}

func point12ValBCommandNameValid(value string) bool {
	trimmed := strings.TrimSpace(value)
	return point11Val0IdentityValueValid(trimmed) && !strings.Contains(trimmed, " ") && !strings.Contains(trimmed, "/")
}

func point12ValBCommandKindValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValBCommandKindReplayProofPack,
		point12ValBCommandKindVerifyManifestContext,
		point12ValBCommandKindExplainMismatch,
		point12ValBCommandKindExplainMissingEvidence,
		point12ValBCommandKindExplainUnsupported,
	}, value)
}

func point12ValBRequestedEngineContextRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"engine_context_", "engine_version_"})
}

func point12ValBRequestedSchemaContextRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"schema_context_", "schema_version_"})
}

func point12ValBDeclaredCompatibilityContextValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"compatibility_context_", "declared_compatibility_context_"})
}

func point12ValBFreshnessContextValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"freshness_context_", "freshness_"})
}

func point12ValBReplayRequestRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"replay_request_"})
}

func point12ValBReplayResultRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"replay_result_"})
}

func point12ValBMismatchRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"replay_mismatch_"})
}

func point12ValBReplayResultStates() []string {
	return []string{
		Point12ValBReplayResultStateActive,
		Point12ValBReplayResultStateBlocked,
		Point12ValBReplayResultStateReviewRequired,
	}
}

func point12ValBReplayResultTaxonomyValid(value string) bool {
	return point11Val0ContainsTrimmed(point12Val0ReplayResults(), value)
}

func point12ValBDecisionStateValueValid(value string) bool {
	trimmed := strings.TrimSpace(value)
	return point11Val0IdentityValueValid(trimmed) && !strings.Contains(trimmed, " ") && !strings.Contains(trimmed, "/")
}

func point12ValBCheckResultValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValBCheckResultActive,
		point12ValBCheckResultMismatch,
		point12ValBCheckResultTampered,
		point12ValBCheckResultUnsupported,
		point12ValBCheckResultMissing,
		point12ValBCheckResultBlocked,
	}, value)
}

func point12ValBMismatchTypeValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValBMismatchTypePolicyMismatch,
		point12ValBMismatchTypeEngineMismatch,
		point12ValBMismatchTypeSchemaMismatch,
		point12ValBMismatchTypeEvidenceMismatch,
		point12ValBMismatchTypeClaimMismatch,
		point12ValBMismatchTypeGovernanceMismatch,
		point12ValBMismatchTypeTenantScopeMismatch,
		point12ValBMismatchTypeArtifactMismatch,
		point12ValBMismatchTypeManifestPayloadMismatch,
		point12ValBMismatchTypeSignatureMetadataMismatch,
		point12ValBMismatchTypeRedactionMismatch,
		point12ValBMismatchTypeUnsupportedVersion,
		point12ValBMismatchTypeTamperDetected,
		point12ValBMismatchTypeMissingEvidence,
		point12ValBMismatchTypeRevokedExpiredEvidence,
		point12ValBMismatchTypeSupersededPolicyOrClaim,
	}, value)
}

func point12ValBDriftClassificationValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValBDriftDueToPolicy,
		point12ValBDriftDueToEngine,
		point12ValBDriftDueToSchema,
		point12ValBDriftDueToEvidence,
		point12ValBDriftDueToClaim,
		point12ValBDriftDueToGovernance,
		point12ValBDriftDueToRedaction,
		point12ValBDriftDueToRevocation,
		point12ValBDriftDueToToolchainProvenance,
		point12ValBDriftDueToTenantScope,
	}, value)
}

func point12ValBAffectedSurfaceValid(value string) bool {
	trimmed := strings.TrimSpace(value)
	return point11Val0IdentityValueValid(trimmed) && !strings.Contains(trimmed, " ") && !strings.Contains(trimmed, "/")
}

func point12ValBHasMismatchType(mismatches []Point12ValBReplayMismatch, mismatchType string) bool {
	for _, mismatch := range mismatches {
		if strings.TrimSpace(mismatch.MismatchType) == strings.TrimSpace(mismatchType) {
			return true
		}
	}
	return false
}

func point12ValBHasDecisiveMismatch(mismatches []Point12ValBReplayMismatch) bool {
	for _, mismatch := range mismatches {
		if mismatch.Decisive {
			return true
		}
	}
	return false
}

func point12ValBHasReplayBlockingMismatch(mismatches []Point12ValBReplayMismatch) bool {
	for _, mismatch := range mismatches {
		if mismatch.BlocksReplay {
			return true
		}
	}
	return false
}

func point12ValBCurrentPolicySupplied(model Point12ValBReplayRequest) bool {
	return strings.TrimSpace(model.CurrentPolicyRef) != "" ||
		strings.TrimSpace(model.CurrentPolicyVersion) != "" ||
		strings.TrimSpace(model.CurrentPolicyHash) != ""
}

func point12ValBCurrentEngineSupplied(model Point12ValBReplayRequest) bool {
	return strings.TrimSpace(model.CurrentEngineVersion) != "" ||
		strings.TrimSpace(model.CurrentEngineHash) != ""
}

func point12ValBCurrentSchemaSupplied(model Point12ValBReplayRequest) bool {
	return strings.TrimSpace(model.CurrentSchemaVersion) != "" ||
		strings.TrimSpace(model.CurrentSchemaHash) != ""
}

func point12ValBCurrentEvidenceSupplied(model Point12ValBReplayRequest) bool {
	return len(model.CurrentEvidenceRefs) > 0 || len(model.CurrentEvidenceHashRefs) > 0
}

func point12ValBCurrentClaimsSupplied(model Point12ValBReplayRequest) bool {
	return len(model.CurrentClaimRefs) > 0
}

func point12ValBCurrentGovernanceSupplied(model Point12ValBReplayRequest) bool {
	return len(model.CurrentGovernanceEventRefs) > 0
}

func point12ValBCurrentPolicyDiffers(model Point12ValBReplayRequest) bool {
	return point12ValBCurrentPolicySupplied(model) &&
		(strings.TrimSpace(model.CurrentPolicyRef) != strings.TrimSpace(model.PolicyRef) ||
			strings.TrimSpace(model.CurrentPolicyVersion) != strings.TrimSpace(model.PolicyVersion) ||
			strings.TrimSpace(model.CurrentPolicyHash) != strings.TrimSpace(model.PolicyHash))
}

func point12ValBCurrentEngineDiffers(model Point12ValBReplayRequest) bool {
	return point12ValBCurrentEngineSupplied(model) &&
		(strings.TrimSpace(model.CurrentEngineVersion) != strings.TrimSpace(model.EngineVersion) ||
			strings.TrimSpace(model.CurrentEngineHash) != strings.TrimSpace(model.EngineHash))
}

func point12ValBCurrentSchemaDiffers(model Point12ValBReplayRequest) bool {
	return point12ValBCurrentSchemaSupplied(model) &&
		(strings.TrimSpace(model.CurrentSchemaVersion) != strings.TrimSpace(model.SchemaVersion) ||
			strings.TrimSpace(model.CurrentSchemaHash) != strings.TrimSpace(model.SchemaHash))
}

func point12ValBCurrentEvidenceDiffers(model Point12ValBReplayRequest) bool {
	return point12ValBCurrentEvidenceSupplied(model) &&
		(!point12Val0ExactStringSetMatch(model.CurrentEvidenceRefs, model.EvidenceRefs) ||
			!point12Val0ExactStringSetMatch(model.CurrentEvidenceHashRefs, model.EvidenceHashRefs))
}

func point12ValBCurrentClaimsDiffer(model Point12ValBReplayRequest) bool {
	return point12ValBCurrentClaimsSupplied(model) &&
		!point12Val0ExactStringSetMatch(model.CurrentClaimRefs, model.ClaimRefs)
}

func point12ValBCurrentGovernanceDiffers(model Point12ValBReplayRequest) bool {
	return point12ValBCurrentGovernanceSupplied(model) &&
		!point12Val0ExactStringSetMatch(model.CurrentGovernanceEventRefs, model.GovernanceEventRefs)
}

func point12ValBDependencyReviewContextModel() Point12ValBValAReviewContext {
	return Point12ValBValAReviewContext{
		SnapshotFromComputedOutput: true,
	}
}

func SnapshotPoint12ValBDependencyFromComputedValA(valA Point12ValAFoundation, review Point12ValBValAReviewContext) Point12ValBDependencySnapshot {
	reviewPrerequisites := append([]string{}, valA.ReviewPrerequisites...)
	reviewPrerequisites = append(reviewPrerequisites, review.ReviewPrerequisites...)
	return Point12ValBDependencySnapshot{
		ValACurrentState:             valA.CurrentState,
		ValADependencyState:          valA.DependencyState,
		ValAManifestIntegrityState:   valA.ManifestIntegrityState,
		ValAPointID:                  valA.Manifest.PointID,
		ValAWaveID:                   valA.Manifest.WaveID,
		Val0RedactionBoundaryState:   valA.Dependency.Val0RedactionBoundaryState,
		ProjectionDisclaimer:         valA.ProjectionDisclaimer,
		SnapshotRef:                  point12ValBDependencySnapshotRefBaseline,
		SnapshotFromComputedOutput:   review.SnapshotFromComputedOutput,
		ValAPrematurePoint12PassSeen: review.ValAPrematurePoint12PassSeen,
		ReviewPrerequisites:          reviewPrerequisites,
		ValAManifest:                 valA.Manifest,
	}
}

func point12ValBDependencySnapshotModel() Point12ValBDependencySnapshot {
	valA := ComputePoint12ValAFoundation(Point12ValAFoundationModel())
	return SnapshotPoint12ValBDependencyFromComputedValA(valA, point12ValBDependencyReviewContextModel())
}

func EvaluatePoint12ValBDependencyState(model Point12ValBDependencySnapshot) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!model.SnapshotFromComputedOutput ||
		!point12ValBDependencySnapshotRefValid(model.SnapshotRef) ||
		strings.TrimSpace(model.ValAPointID) != point12Val0PointID ||
		strings.TrimSpace(model.ValAWaveID) != point12ValAWaveID ||
		strings.TrimSpace(model.ValAManifest.PointID) != point12Val0PointID ||
		strings.TrimSpace(model.ValAManifest.WaveID) != point12ValAWaveID ||
		!point12Val0ProofPackRefValid(model.ValAManifest.ProofPackID) ||
		!point12ValAManifestRefValid(model.ValAManifest.ManifestID) ||
		!point12Val0HashValid(model.ValAManifest.ManifestPayloadHash) ||
		model.ValAPrematurePoint12PassSeen ||
		point12Val0ContainsPrematurePassToken(
			model.ValAManifest.ProofPackID,
			model.ValAManifest.ManifestID,
			model.ValAManifest.DecisionID,
			model.ValAManifest.SignatureRef,
			model.ValAManifest.DetachedSignatureRef,
			model.ValAManifest.SignatureMetadataRef,
			strings.Join(model.ValAManifest.ManifestOutputClaims, " "),
		) {
		return Point12ValBDependencyStateBlocked
	}
	if strings.TrimSpace(model.ValACurrentState) == Point12ValAStateBlocked ||
		strings.TrimSpace(model.ValADependencyState) == Point12ValADependencyStateBlocked ||
		strings.TrimSpace(model.ValAManifestIntegrityState) == Point12ValAManifestIntegrityStateBlocked ||
		strings.TrimSpace(model.ValAManifestIntegrityState) == Point12ValAManifestIntegrityStateIncomplete ||
		strings.TrimSpace(model.ValAManifestIntegrityState) == Point12ValAManifestIntegrityStateUnsupported ||
		strings.TrimSpace(model.ValAManifestIntegrityState) == Point12ValAManifestIntegrityStateTampered ||
		strings.TrimSpace(model.Val0RedactionBoundaryState) == Point12Val0RedactionBoundaryStateBlocked {
		return Point12ValBDependencyStateBlocked
	}
	if strings.TrimSpace(model.ValACurrentState) == Point12ValAStateReviewRequired ||
		strings.TrimSpace(model.ValADependencyState) == Point12ValADependencyStateReviewRequired ||
		strings.TrimSpace(model.ValAManifestIntegrityState) == Point12ValAManifestIntegrityStateReviewRequired ||
		len(model.ReviewPrerequisites) > 0 {
		return Point12ValBDependencyStateReviewRequired
	}
	if strings.TrimSpace(model.ValACurrentState) != Point12ValAStateActive ||
		strings.TrimSpace(model.ValADependencyState) != Point12ValADependencyStateActive ||
		strings.TrimSpace(model.ValAManifestIntegrityState) != Point12ValAManifestIntegrityStateActive {
		return Point12ValBDependencyStateBlocked
	}
	return Point12ValBDependencyStateActive
}

func point12ValBReplayCommandStateAndReasons(
	model Point12ValBReplayCommandContract,
	dependency Point12ValBDependencySnapshot,
) (string, []string) {
	reasons := []string{}
	if !point12ValBCommandRefValid(model.CommandID) ||
		!point12ValBCommandNameValid(model.CommandName) ||
		!point12ValBCommandKindValid(model.CommandKind) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12Val0ArtifactRefValid(model.ArtifactRef) ||
		!point12Val0ReplayModeValid(model.ReplayMode) ||
		!point12Val0CompatibilityProfileRefValid(model.CompatibilityProfileRef) ||
		!point11Val0ValidTimestamp(model.GeneratedAt) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "replay_command_identity_or_contract_invalid")
	}
	if model.AllowExternalAPI {
		reasons = append(reasons, "replay_command_external_api_forbidden")
	}
	if model.OfflineBundleRequired {
		reasons = append(reasons, "replay_command_offline_bundle_not_supported")
	}
	if model.MutatesEvidenceSpine {
		reasons = append(reasons, "replay_command_mutates_evidence_spine")
	}
	if model.CreatesProofPack {
		reasons = append(reasons, "replay_command_creates_proof_pack_forbidden")
	}
	if model.CreatesAuditExportBundle {
		reasons = append(reasons, "replay_command_audit_export_bundle_forbidden")
	}
	if model.OpensPortalPath {
		reasons = append(reasons, "replay_command_portal_path_forbidden")
	}
	if model.RequestsPoint12Pass {
		reasons = append(reasons, "replay_command_point12_pass_forbidden")
	}
	if point12Val0ContainsPrematurePassToken(model.CommandID, model.CommandName) {
		reasons = append(reasons, "replay_command_premature_point12_pass")
	}
	if strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(dependency.ValAManifest.ProofPackID) ||
		strings.TrimSpace(model.ManifestID) != strings.TrimSpace(dependency.ValAManifest.ManifestID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValAManifest.TenantScope) ||
		strings.TrimSpace(model.ArtifactRef) != strings.TrimSpace(dependency.ValAManifest.ArtifactRef) ||
		strings.TrimSpace(model.CompatibilityProfileRef) != strings.TrimSpace(dependency.ValAManifest.CompatibilityProfileRef) {
		reasons = append(reasons, "replay_command_manifest_binding_mismatch")
	}
	switch strings.TrimSpace(model.ReplayMode) {
	case point12Val0ReplayModeOriginalContext:
		if model.AllowCurrentPolicy {
			reasons = append(reasons, "replay_command_original_context_cannot_allow_current_policy")
		}
	case point12Val0ReplayModeCurrentPolicyContext, point12Val0ReplayModeComparisonMode:
		if !model.AllowCurrentPolicy {
			reasons = append(reasons, "replay_command_current_policy_mode_requires_explicit_allow_current_policy")
		}
		if !point12Val0PolicyRefValid(model.RequestedPolicyContextRef) ||
			!point12ValBRequestedEngineContextRefValid(model.RequestedEngineContextRef) ||
			!point12ValBRequestedSchemaContextRefValid(model.RequestedSchemaContextRef) {
			reasons = append(reasons, "replay_command_current_context_refs_invalid")
		}
	}
	switch strings.TrimSpace(model.CommandKind) {
	case point12ValBCommandKindExplainMismatch:
		if !model.ExplainMismatch {
			reasons = append(reasons, "replay_command_explain_mismatch_flag_missing")
		}
	case point12ValBCommandKindExplainMissingEvidence:
		if !model.ExplainMissingEvidence {
			reasons = append(reasons, "replay_command_explain_missing_evidence_flag_missing")
		}
	case point12ValBCommandKindExplainUnsupported:
		if !model.ExplainUnsupportedVersion {
			reasons = append(reasons, "replay_command_explain_unsupported_flag_missing")
		}
	}
	if strings.TrimSpace(model.ReplayMode) == point12Val0ReplayModeComparisonMode && !model.ExplainMismatch {
		reasons = append(reasons, "replay_command_comparison_mode_requires_mismatch_explanation")
	}
	if len(reasons) > 0 {
		return Point12ValBReplayCommandStateBlocked, reasons
	}
	return Point12ValBReplayCommandStateActive, nil
}

func EvaluatePoint12ValBReplayCommandState(
	model Point12ValBReplayCommandContract,
	dependency Point12ValBDependencySnapshot,
) string {
	state, _ := point12ValBReplayCommandStateAndReasons(model, dependency)
	return state
}

func point12ValBReplayRequestStateAndReasons(
	model Point12ValBReplayRequest,
	dependency Point12ValBDependencySnapshot,
	command Point12ValBReplayCommandContract,
) (string, []string) {
	reasons := []string{}
	if !point12ValBReplayRequestRefValid(model.ReplayRequestID) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point12Val0DecisionRefValid(model.DecisionID) ||
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
		!point12Val0HashValid(model.ManifestPayloadHash) ||
		!point12Val0CompatibilityProfileRefValid(model.CompatibilityProfileRef) ||
		!point12Val0ReplayModeValid(model.ReplayMode) ||
		!point12ValBDeclaredCompatibilityContextValid(model.DeclaredCompatibilityContext) ||
		!point12ValBDecisionStateValueValid(model.OriginalDecisionState) ||
		!point12Val0HashValid(model.OriginalDecisionHash) ||
		!point11Val0ValidTimestamp(model.GeneratedAt) ||
		!point12ValBFreshnessContextValid(model.FreshnessContext) ||
		!point11Val0ContainsTrimmed(point12ValAManifestCoreStates(), model.SourceManifestIntegrityState) {
		reasons = append(reasons, "replay_request_identity_or_context_invalid")
	}
	if strings.TrimSpace(model.RedactionManifestRef) != "" && !point12Val0RedactionManifestRefValid(model.RedactionManifestRef) {
		reasons = append(reasons, "replay_request_redaction_manifest_invalid")
	}
	if len(model.EvidenceRefs) != len(model.EvidenceHashRefs) {
		reasons = append(reasons, "replay_request_evidence_hash_alignment_invalid")
	}
	if point12ValBCurrentPolicySupplied(model) &&
		(!point12Val0PolicyRefValid(model.CurrentPolicyRef) ||
			!point12Val0VersionIdentityValid(model.CurrentPolicyVersion) ||
			!point12Val0HashValid(model.CurrentPolicyHash)) {
		reasons = append(reasons, "replay_request_current_policy_context_invalid")
	}
	if point12ValBCurrentEngineSupplied(model) &&
		(!point12Val0VersionIdentityValid(model.CurrentEngineVersion) ||
			!point12Val0HashValid(model.CurrentEngineHash)) {
		reasons = append(reasons, "replay_request_current_engine_context_invalid")
	}
	if point12ValBCurrentSchemaSupplied(model) &&
		(!point12Val0VersionIdentityValid(model.CurrentSchemaVersion) ||
			!point12Val0HashValid(model.CurrentSchemaHash)) {
		reasons = append(reasons, "replay_request_current_schema_context_invalid")
	}
	if point12ValBCurrentEvidenceSupplied(model) &&
		(!point12Val0EvidenceRefsValid(model.CurrentEvidenceRefs) ||
			!point12Val0StringListValid(model.CurrentEvidenceHashRefs, point12Val0EvidenceHashRefValid) ||
			len(model.CurrentEvidenceRefs) != len(model.CurrentEvidenceHashRefs)) {
		reasons = append(reasons, "replay_request_current_evidence_context_invalid")
	}
	if point12ValBCurrentClaimsSupplied(model) && !point12Val0StringListValid(model.CurrentClaimRefs, point12Val0ClaimRefValid) {
		reasons = append(reasons, "replay_request_current_claim_context_invalid")
	}
	if point12ValBCurrentGovernanceSupplied(model) && !point12Val0StringListValid(model.CurrentGovernanceEventRefs, point12Val0GovernanceEventRefValid) {
		reasons = append(reasons, "replay_request_current_governance_context_invalid")
	}
	if point12Val0ContainsPrematurePassToken(
		model.ReplayRequestID,
		model.ProofPackID,
		model.ManifestID,
		model.DecisionID,
		model.DeclaredCompatibilityContext,
	) {
		reasons = append(reasons, "replay_request_premature_point12_pass")
	}
	if strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(dependency.ValAManifest.ProofPackID) ||
		strings.TrimSpace(model.ManifestID) != strings.TrimSpace(dependency.ValAManifest.ManifestID) ||
		strings.TrimSpace(model.DecisionID) != strings.TrimSpace(dependency.ValAManifest.DecisionID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValAManifest.TenantScope) ||
		strings.TrimSpace(model.ArtifactRef) != strings.TrimSpace(dependency.ValAManifest.ArtifactRef) ||
		strings.TrimSpace(model.ArtifactHash) != strings.TrimSpace(dependency.ValAManifest.ArtifactHash) ||
		strings.TrimSpace(model.CompatibilityProfileRef) != strings.TrimSpace(dependency.ValAManifest.CompatibilityProfileRef) {
		reasons = append(reasons, "replay_request_manifest_identity_binding_mismatch")
	}
	if strings.TrimSpace(model.SourceManifestIntegrityState) != strings.TrimSpace(dependency.ValAManifestIntegrityState) ||
		strings.TrimSpace(model.SourceManifestIntegrityState) != Point12ValAManifestIntegrityStateActive {
		reasons = append(reasons, "replay_request_source_manifest_integrity_invalid")
	}
	if strings.TrimSpace(model.ReplayMode) != strings.TrimSpace(command.ReplayMode) {
		reasons = append(reasons, "replay_request_command_mode_mismatch")
	}
	if point12ValBCurrentPolicySupplied(model) &&
		strings.TrimSpace(command.RequestedPolicyContextRef) != "" &&
		strings.TrimSpace(command.RequestedPolicyContextRef) != strings.TrimSpace(model.CurrentPolicyRef) {
		reasons = append(reasons, "replay_request_command_policy_context_binding_mismatch")
	}
	if point12ValBCurrentEngineSupplied(model) &&
		strings.TrimSpace(command.RequestedEngineContextRef) != "" &&
		strings.TrimSpace(command.RequestedEngineContextRef) != strings.TrimSpace(model.CurrentEngineVersion) {
		reasons = append(reasons, "replay_request_command_engine_context_binding_mismatch")
	}
	if point12ValBCurrentSchemaSupplied(model) &&
		strings.TrimSpace(command.RequestedSchemaContextRef) != "" &&
		strings.TrimSpace(command.RequestedSchemaContextRef) != strings.TrimSpace(model.CurrentSchemaVersion) {
		reasons = append(reasons, "replay_request_command_schema_context_binding_mismatch")
	}
	switch strings.TrimSpace(model.ReplayMode) {
	case point12Val0ReplayModeCurrentPolicyContext, point12Val0ReplayModeComparisonMode:
		if !point12ValBCurrentPolicySupplied(model) ||
			!point12ValBCurrentEngineSupplied(model) ||
			!point12ValBCurrentSchemaSupplied(model) ||
			!point12ValBCurrentEvidenceSupplied(model) {
			reasons = append(reasons, "replay_request_current_policy_mode_requires_explicit_current_context")
		}
	}
	if len(reasons) > 0 {
		return Point12ValBReplayRequestStateBlocked, reasons
	}
	return Point12ValBReplayRequestStateActive, nil
}

func EvaluatePoint12ValBReplayRequestState(
	model Point12ValBReplayRequest,
	dependency Point12ValBDependencySnapshot,
	command Point12ValBReplayCommandContract,
) string {
	state, _ := point12ValBReplayRequestStateAndReasons(model, dependency, command)
	return state
}

func point12ValBMismatchNeedsComparisonData(mismatchType string) bool {
	switch strings.TrimSpace(mismatchType) {
	case point12ValBMismatchTypePolicyMismatch,
		point12ValBMismatchTypeEngineMismatch,
		point12ValBMismatchTypeSchemaMismatch,
		point12ValBMismatchTypeEvidenceMismatch,
		point12ValBMismatchTypeClaimMismatch,
		point12ValBMismatchTypeGovernanceMismatch,
		point12ValBMismatchTypeTenantScopeMismatch,
		point12ValBMismatchTypeArtifactMismatch,
		point12ValBMismatchTypeManifestPayloadMismatch,
		point12ValBMismatchTypeSignatureMetadataMismatch,
		point12ValBMismatchTypeRedactionMismatch,
		point12ValBMismatchTypeUnsupportedVersion,
		point12ValBMismatchTypeTamperDetected,
		point12ValBMismatchTypeMissingEvidence,
		point12ValBMismatchTypeRevokedExpiredEvidence,
		point12ValBMismatchTypeSupersededPolicyOrClaim:
		return true
	default:
		return false
	}
}

func point12ValBMismatchHasComparisonData(model Point12ValBReplayMismatch) bool {
	return strings.TrimSpace(model.ExpectedRef) != "" && strings.TrimSpace(model.ActualRef) != "" ||
		strings.TrimSpace(model.ExpectedHash) != "" && strings.TrimSpace(model.ActualHash) != "" ||
		strings.TrimSpace(model.ExpectedVersion) != "" && strings.TrimSpace(model.ActualVersion) != ""
}

func point12ValBMismatchStateAndReasons(model Point12ValBReplayMismatch) []string {
	reasons := []string{}
	if !point12ValBMismatchRefValid(model.MismatchID) ||
		!point12ValBMismatchTypeValid(model.MismatchType) ||
		!point12ValBAffectedSurfaceValid(model.AffectedSurface) {
		reasons = append(reasons, "replay_mismatch_identity_invalid")
	}
	if point12ValBMismatchNeedsComparisonData(model.MismatchType) && !point12ValBMismatchHasComparisonData(model) {
		reasons = append(reasons, "replay_mismatch_expected_actual_missing")
	}
	if strings.TrimSpace(model.DriftReason) != "" && !point12ValBDriftClassificationValid(model.DriftReason) {
		reasons = append(reasons, "replay_mismatch_drift_reason_invalid")
	}
	return reasons
}

func point12ValBReplayResultStateAndReasons(
	model Point12ValBReplayResult,
	request Point12ValBReplayRequest,
	command Point12ValBReplayCommandContract,
	dependency Point12ValBDependencySnapshot,
) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}

	if !point12ValBReplayResultRefValid(model.ReplayResultID) ||
		!point12ValBReplayRequestRefValid(model.ReplayRequestID) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point11Val0ContainsTrimmed(point12ValBReplayResultStates(), model.ReplayState) ||
		!point12Val0ReplayModeValid(model.ReplayMode) ||
		!point12ValBReplayResultTaxonomyValid(model.ReplayResultTaxonomy) ||
		!point12ValBDecisionStateValueValid(model.OriginalDecisionState) ||
		!point12ValBDecisionStateValueValid(model.ReplayedDecisionState) ||
		!point12Val0VersionIdentityValid(model.EvaluatedPolicyVersion) ||
		!point12Val0VersionIdentityValid(model.EvaluatedEngineVersion) ||
		!point12Val0VersionIdentityValid(model.EvaluatedSchemaVersion) ||
		!point12ValBCheckResultValid(model.EvidenceHashCheckResult) ||
		!point12ValBCheckResultValid(model.ManifestIntegrityCheckResult) ||
		!point12ValBCheckResultValid(model.SignatureMetadataCheckResult) ||
		!point12ValBCheckResultValid(model.CompatibilityCheckResult) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12Val0OptionalClaimTextListValid(model.ReplayOutputClaims) {
		blockedReasons = append(blockedReasons, "replay_result_identity_or_metadata_invalid")
	}
	if strings.TrimSpace(model.ReplayRequestID) != strings.TrimSpace(request.ReplayRequestID) ||
		strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(request.ProofPackID) ||
		strings.TrimSpace(model.ManifestID) != strings.TrimSpace(request.ManifestID) ||
		strings.TrimSpace(model.ReplayMode) != strings.TrimSpace(request.ReplayMode) ||
		strings.TrimSpace(model.OriginalDecisionState) != strings.TrimSpace(request.OriginalDecisionState) {
		blockedReasons = append(blockedReasons, "replay_result_request_binding_mismatch")
	}
	if model.ExternalAPIUsed {
		blockedReasons = append(blockedReasons, "replay_result_external_api_used")
	}
	if model.PointPassEmitted {
		blockedReasons = append(blockedReasons, "replay_result_point12_pass_emitted")
	}
	if point12Val0ContainsPrematurePassToken(
		model.ReplayResultID,
		model.ReplayRequestID,
		model.BlockedReason,
		model.DecisionDriftExplanation,
		model.CustomerVisibleExplanation,
		strings.Join(model.MismatchExplanations, " "),
		strings.Join(model.ReplayOutputClaims, " "),
	) {
		blockedReasons = append(blockedReasons, "replay_result_premature_point12_pass")
	}
	if point12Val0ContainsForbiddenClaim(strings.Join(model.ReplayOutputClaims, " "), model.CustomerVisibleExplanation) {
		blockedReasons = append(blockedReasons, "replay_result_overclaim_detected")
	}
	for _, mismatch := range model.Mismatches {
		if mismatchReasons := point12ValBMismatchStateAndReasons(mismatch); len(mismatchReasons) > 0 {
			blockedReasons = append(blockedReasons, mismatchReasons...)
		}
		if mismatch.Decisive && strings.TrimSpace(mismatch.Explanation) == "" {
			reviewReasons = append(reviewReasons, "replay_result_decisive_mismatch_explanation_missing")
		}
	}

	originalPolicyMismatch := strings.TrimSpace(request.PolicyRef) != strings.TrimSpace(dependency.ValAManifest.PolicyRef) ||
		strings.TrimSpace(request.PolicyVersion) != strings.TrimSpace(dependency.ValAManifest.PolicyVersion) ||
		strings.TrimSpace(request.PolicyHash) != strings.TrimSpace(dependency.ValAManifest.PolicyHash)
	originalEngineMismatch := strings.TrimSpace(request.EngineVersion) != strings.TrimSpace(dependency.ValAManifest.EngineVersion) ||
		strings.TrimSpace(request.EngineHash) != strings.TrimSpace(dependency.ValAManifest.EngineHash)
	originalSchemaMismatch := strings.TrimSpace(request.SchemaVersion) != strings.TrimSpace(dependency.ValAManifest.SchemaVersion) ||
		strings.TrimSpace(request.SchemaHash) != strings.TrimSpace(dependency.ValAManifest.SchemaHash)
	originalEvidenceMismatch := !point12Val0ExactStringSetMatch(request.EvidenceRefs, dependency.ValAManifest.EvidenceRefs) ||
		!point12Val0ExactStringSetMatch(request.EvidenceHashRefs, dependency.ValAManifest.EvidenceHashRefs)
	originalClaimMismatch := !point12Val0ExactStringSetMatch(request.ClaimRefs, dependency.ValAManifest.ClaimRefs)
	originalGovernanceMismatch := !point12Val0ExactStringSetMatch(request.GovernanceEventRefs, dependency.ValAManifest.GovernanceEventRefs)
	manifestPayloadMismatch := strings.TrimSpace(request.ManifestPayloadHash) != strings.TrimSpace(dependency.ValAManifest.ManifestPayloadHash)
	redactionBindingMismatch := strings.TrimSpace(request.RedactionManifestRef) != strings.TrimSpace(dependency.ValAManifest.RedactionManifestRef)
	currentPolicyDiff := point12ValBCurrentPolicyDiffers(request)
	currentEngineDiff := point12ValBCurrentEngineDiffers(request)
	currentSchemaDiff := point12ValBCurrentSchemaDiffers(request)
	currentEvidenceDiff := point12ValBCurrentEvidenceDiffers(request)
	currentClaimDiff := point12ValBCurrentClaimsDiffer(request)
	currentGovernanceDiff := point12ValBCurrentGovernanceDiffers(request)

	hasTamper := model.TamperDetected ||
		model.ManifestIntegrityCheckResult == point12ValBCheckResultTampered ||
		model.SignatureMetadataCheckResult == point12ValBCheckResultTampered ||
		model.EvidenceHashCheckResult == point12ValBCheckResultTampered ||
		manifestPayloadMismatch ||
		redactionBindingMismatch ||
		point12ValBHasMismatchType(model.Mismatches, point12ValBMismatchTypeTamperDetected) ||
		point12ValBHasMismatchType(model.Mismatches, point12ValBMismatchTypeManifestPayloadMismatch) ||
		point12ValBHasMismatchType(model.Mismatches, point12ValBMismatchTypeSignatureMetadataMismatch)
	hasUnsupported := model.UnsupportedVersion ||
		model.CompatibilityCheckResult == point12ValBCheckResultUnsupported ||
		model.ManifestIntegrityCheckResult == point12ValBCheckResultUnsupported ||
		point12ValBHasMismatchType(model.Mismatches, point12ValBMismatchTypeUnsupportedVersion)
	hasInsufficient := model.InsufficientEvidence ||
		model.EvidenceHashCheckResult == point12ValBCheckResultMissing ||
		point12ValBHasMismatchType(model.Mismatches, point12ValBMismatchTypeMissingEvidence)
	hasRedactionLimitations := model.RedactionLimitations ||
		point12ValBHasMismatchType(model.Mismatches, point12ValBMismatchTypeRedactionMismatch)

	if hasTamper &&
		strings.TrimSpace(model.ReplayResultTaxonomy) != Point12Val0ReplayResultTamperDetected &&
		strings.TrimSpace(model.ReplayResultTaxonomy) != Point12Val0ReplayResultBlockedReplay {
		blockedReasons = append(blockedReasons, "replay_result_tamper_must_outrank_same_decision")
	}
	if hasUnsupported &&
		strings.TrimSpace(model.ReplayResultTaxonomy) != Point12Val0ReplayResultUnsupportedVersion &&
		strings.TrimSpace(model.ReplayResultTaxonomy) != Point12Val0ReplayResultBlockedReplay {
		blockedReasons = append(blockedReasons, "replay_result_unsupported_version_must_outrank_same_decision")
	}
	if hasInsufficient &&
		strings.TrimSpace(model.ReplayResultTaxonomy) != Point12Val0ReplayResultInsufficientEvidence &&
		strings.TrimSpace(model.ReplayResultTaxonomy) != Point12Val0ReplayResultBlockedReplay {
		blockedReasons = append(blockedReasons, "replay_result_insufficient_evidence_must_outrank_same_decision")
	}
	if hasRedactionLimitations &&
		strings.TrimSpace(model.ReplayResultTaxonomy) != Point12Val0ReplayResultRedactedLimitations &&
		strings.TrimSpace(model.ReplayResultTaxonomy) != Point12Val0ReplayResultInsufficientEvidence &&
		strings.TrimSpace(model.ReplayResultTaxonomy) != Point12Val0ReplayResultBlockedReplay {
		blockedReasons = append(blockedReasons, "replay_result_redaction_limitations_must_fail_closed")
	}

	if originalPolicyMismatch &&
		!point11Val0ContainsTrimmed([]string{
			Point12Val0ReplayResultPolicyMismatch,
			Point12Val0ReplayResultTamperDetected,
			Point12Val0ReplayResultBlockedReplay,
		}, model.ReplayResultTaxonomy) {
		blockedReasons = append(blockedReasons, "replay_result_original_policy_mismatch_unexplained")
	}
	if originalEngineMismatch &&
		!point11Val0ContainsTrimmed([]string{
			Point12Val0ReplayResultEngineMismatch,
			Point12Val0ReplayResultTamperDetected,
			Point12Val0ReplayResultBlockedReplay,
		}, model.ReplayResultTaxonomy) {
		blockedReasons = append(blockedReasons, "replay_result_original_engine_mismatch_unexplained")
	}
	if originalSchemaMismatch &&
		!point11Val0ContainsTrimmed([]string{
			Point12Val0ReplayResultSchemaMismatch,
			Point12Val0ReplayResultTamperDetected,
			Point12Val0ReplayResultBlockedReplay,
		}, model.ReplayResultTaxonomy) {
		blockedReasons = append(blockedReasons, "replay_result_original_schema_mismatch_unexplained")
	}
	if originalEvidenceMismatch &&
		!point11Val0ContainsTrimmed([]string{
			Point12Val0ReplayResultEvidenceMismatch,
			Point12Val0ReplayResultInsufficientEvidence,
			Point12Val0ReplayResultBlockedReplay,
			Point12Val0ReplayResultTamperDetected,
		}, model.ReplayResultTaxonomy) {
		blockedReasons = append(blockedReasons, "replay_result_original_evidence_mismatch_unexplained")
	}
	if originalClaimMismatch &&
		!point11Val0ContainsTrimmed([]string{
			Point12Val0ReplayResultClaimMismatch,
			Point12Val0ReplayResultBlockedReplay,
		}, model.ReplayResultTaxonomy) {
		blockedReasons = append(blockedReasons, "replay_result_original_claim_mismatch_unexplained")
	}
	if originalGovernanceMismatch &&
		!point11Val0ContainsTrimmed([]string{
			Point12Val0ReplayResultGovernanceMismatch,
			Point12Val0ReplayResultBlockedReplay,
		}, model.ReplayResultTaxonomy) {
		blockedReasons = append(blockedReasons, "replay_result_original_governance_mismatch_unexplained")
	}

	if strings.TrimSpace(request.ReplayMode) == point12Val0ReplayModeOriginalContext {
		if currentPolicyDiff &&
			!point11Val0ContainsTrimmed([]string{Point12Val0ReplayResultPolicyMismatch, Point12Val0ReplayResultBlockedReplay}, model.ReplayResultTaxonomy) {
			blockedReasons = append(blockedReasons, "replay_result_original_context_cannot_silently_use_current_policy")
		}
		if currentEngineDiff &&
			!point11Val0ContainsTrimmed([]string{Point12Val0ReplayResultEngineMismatch, Point12Val0ReplayResultBlockedReplay}, model.ReplayResultTaxonomy) {
			blockedReasons = append(blockedReasons, "replay_result_original_context_cannot_silently_use_current_engine")
		}
		if currentSchemaDiff &&
			!point11Val0ContainsTrimmed([]string{Point12Val0ReplayResultSchemaMismatch, Point12Val0ReplayResultBlockedReplay}, model.ReplayResultTaxonomy) {
			blockedReasons = append(blockedReasons, "replay_result_original_context_cannot_silently_use_current_schema")
		}
		if currentEvidenceDiff &&
			!point11Val0ContainsTrimmed([]string{Point12Val0ReplayResultEvidenceMismatch, Point12Val0ReplayResultBlockedReplay}, model.ReplayResultTaxonomy) {
			blockedReasons = append(blockedReasons, "replay_result_original_context_cannot_silently_use_current_evidence")
		}
		if currentClaimDiff &&
			!point11Val0ContainsTrimmed([]string{Point12Val0ReplayResultClaimMismatch, Point12Val0ReplayResultBlockedReplay}, model.ReplayResultTaxonomy) {
			blockedReasons = append(blockedReasons, "replay_result_original_context_cannot_silently_use_current_claims")
		}
		if currentGovernanceDiff &&
			!point11Val0ContainsTrimmed([]string{Point12Val0ReplayResultGovernanceMismatch, Point12Val0ReplayResultBlockedReplay}, model.ReplayResultTaxonomy) {
			blockedReasons = append(blockedReasons, "replay_result_original_context_cannot_silently_use_current_governance")
		}
	}

	if strings.TrimSpace(model.ReplayResultTaxonomy) == Point12Val0ReplayResultSameDecision {
		if !model.MatchOriginal ||
			strings.TrimSpace(model.OriginalDecisionState) != strings.TrimSpace(model.ReplayedDecisionState) ||
			hasTamper ||
			hasUnsupported ||
			hasInsufficient ||
			hasRedactionLimitations ||
			point12ValBHasDecisiveMismatch(model.Mismatches) {
			blockedReasons = append(blockedReasons, "replay_result_same_decision_overclaims_replay")
		}
		if model.EvidenceHashCheckResult != point12ValBCheckResultActive ||
			model.ManifestIntegrityCheckResult != point12ValBCheckResultActive ||
			model.SignatureMetadataCheckResult != point12ValBCheckResultActive ||
			model.CompatibilityCheckResult != point12ValBCheckResultActive {
			blockedReasons = append(blockedReasons, "replay_result_same_decision_requires_all_checks_active")
		}
	}

	if strings.TrimSpace(model.ReplayResultTaxonomy) == Point12Val0ReplayResultDifferentDecision {
		if model.MatchOriginal || strings.TrimSpace(model.OriginalDecisionState) == strings.TrimSpace(model.ReplayedDecisionState) {
			blockedReasons = append(blockedReasons, "replay_result_different_decision_must_not_match_original")
		}
		if strings.TrimSpace(model.DecisionDriftExplanation) == "" {
			reviewReasons = append(reviewReasons, "replay_result_different_decision_explanation_missing")
		}
	}

	if strings.TrimSpace(model.ReplayResultTaxonomy) == Point12Val0ReplayResultBlockedReplay &&
		strings.TrimSpace(model.BlockedReason) == "" &&
		!point12ValBHasReplayBlockingMismatch(model.Mismatches) &&
		!hasTamper &&
		!hasUnsupported &&
		!hasInsufficient &&
		!hasRedactionLimitations {
		blockedReasons = append(blockedReasons, "replay_result_blocked_replay_reason_missing")
	}

	for taxonomy, mismatchType := range map[string]string{
		Point12Val0ReplayResultPolicyMismatch:     point12ValBMismatchTypePolicyMismatch,
		Point12Val0ReplayResultEngineMismatch:     point12ValBMismatchTypeEngineMismatch,
		Point12Val0ReplayResultSchemaMismatch:     point12ValBMismatchTypeSchemaMismatch,
		Point12Val0ReplayResultEvidenceMismatch:   point12ValBMismatchTypeEvidenceMismatch,
		Point12Val0ReplayResultClaimMismatch:      point12ValBMismatchTypeClaimMismatch,
		Point12Val0ReplayResultGovernanceMismatch: point12ValBMismatchTypeGovernanceMismatch,
	} {
		if strings.TrimSpace(model.ReplayResultTaxonomy) == taxonomy && !point12ValBHasMismatchType(model.Mismatches, mismatchType) {
			blockedReasons = append(blockedReasons, "replay_result_specific_mismatch_missing_entry")
		}
	}

	if strings.TrimSpace(request.ReplayMode) == point12Val0ReplayModeComparisonMode {
		if strings.TrimSpace(model.DecisionDriftExplanation) == "" {
			reviewReasons = append(reviewReasons, "replay_result_comparison_mode_requires_drift_explanation")
		}
		if (currentPolicyDiff || currentEngineDiff || currentSchemaDiff || currentEvidenceDiff || currentClaimDiff || currentGovernanceDiff ||
			strings.TrimSpace(model.ReplayResultTaxonomy) == Point12Val0ReplayResultDifferentDecision ||
			point12ValBHasDecisiveMismatch(model.Mismatches)) &&
			!point12ValBDriftClassificationValid(model.DecisionDriftClassification) {
			reviewReasons = append(reviewReasons, "replay_result_comparison_mode_requires_drift_classification")
		}
	}

	if strings.TrimSpace(request.ReplayMode) == point12Val0ReplayModeCurrentPolicyContext &&
		strings.TrimSpace(model.ReplayResultTaxonomy) == Point12Val0ReplayResultDifferentDecision &&
		strings.TrimSpace(model.DecisionDriftExplanation) == "" {
		reviewReasons = append(reviewReasons, "replay_result_current_policy_drift_explanation_missing")
	}

	if strings.TrimSpace(request.ReplayMode) == point12Val0ReplayModeCurrentPolicyContext &&
		strings.TrimSpace(model.OriginalDecisionState) != strings.TrimSpace(request.OriginalDecisionState) {
		blockedReasons = append(blockedReasons, "replay_result_current_policy_cannot_rewrite_original_decision")
	}

	if len(blockedReasons) > 0 {
		return Point12ValBReplayResultStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValBReplayResultStateReviewRequired, reviewReasons
	}
	return Point12ValBReplayResultStateActive, nil
}

func EvaluatePoint12ValBReplayResultState(
	model Point12ValBReplayResult,
	request Point12ValBReplayRequest,
	command Point12ValBReplayCommandContract,
	dependency Point12ValBDependencySnapshot,
) string {
	state, _ := point12ValBReplayResultStateAndReasons(model, request, command, dependency)
	return state
}

func EvaluatePoint12ValBState(model Point12ValBFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		strings.TrimSpace(model.DependencyState) == Point12ValBDependencyStateBlocked ||
		strings.TrimSpace(model.ReplayCommandState) == Point12ValBReplayCommandStateBlocked ||
		strings.TrimSpace(model.ReplayRequestState) == Point12ValBReplayRequestStateBlocked ||
		strings.TrimSpace(model.ReplayResultState) == Point12ValBReplayResultStateBlocked {
		return Point12ValBStateBlocked
	}
	if strings.TrimSpace(model.DependencyState) == Point12ValBDependencyStateReviewRequired ||
		strings.TrimSpace(model.ReplayResultState) == Point12ValBReplayResultStateReviewRequired {
		return Point12ValBStateReviewRequired
	}
	return Point12ValBStateActive
}

func point12ValBBlockingReasons(model Point12ValBFoundation) []string {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "aggregate_projection_disclaimer_blocked")
	}
	if model.DependencyState == Point12ValBDependencyStateBlocked {
		reasons = append(reasons, "point12_vala_dependency_blocked")
	}
	if model.ReplayCommandState == Point12ValBReplayCommandStateBlocked {
		reasons = append(reasons, "replay_command_blocked")
	}
	if model.ReplayRequestState == Point12ValBReplayRequestStateBlocked {
		reasons = append(reasons, "replay_request_blocked")
	}
	if model.ReplayResultState == Point12ValBReplayResultStateBlocked {
		reasons = append(reasons, "replay_result_blocked")
	}
	return reasons
}

func Point12ValBFoundationModel() Point12ValBFoundation {
	dependency := point12ValBDependencySnapshotModel()
	command := Point12ValBReplayCommandContract{
		CommandID:                   "replay_command_point12_valb_001",
		CommandName:                 point12ValBCommandKindReplayProofPack,
		CommandKind:                 point12ValBCommandKindReplayProofPack,
		ProofPackID:                 dependency.ValAManifest.ProofPackID,
		ManifestID:                  dependency.ValAManifest.ManifestID,
		TenantScope:                 dependency.ValAManifest.TenantScope,
		ArtifactRef:                 dependency.ValAManifest.ArtifactRef,
		ReplayMode:                  point12Val0ReplayModeOriginalContext,
		CompatibilityProfileRef:     dependency.ValAManifest.CompatibilityProfileRef,
		AllowCurrentPolicy:          false,
		AllowExternalAPI:            false,
		OfflineBundleRequired:       false,
		ExplainMismatch:             true,
		ExplainMissingEvidence:      true,
		ExplainUnsupportedVersion:   true,
		ExplainRedactionLimitations: true,
		GeneratedAt:                 "2026-05-03T13:00:00Z",
		ProjectionDisclaimer:        point12ValBProjectionDisclaimerBaseline,
	}
	request := Point12ValBReplayRequest{
		ReplayRequestID:              "replay_request_point12_valb_001",
		ProofPackID:                  dependency.ValAManifest.ProofPackID,
		ManifestID:                   dependency.ValAManifest.ManifestID,
		DecisionID:                   dependency.ValAManifest.DecisionID,
		TenantScope:                  dependency.ValAManifest.TenantScope,
		ArtifactRef:                  dependency.ValAManifest.ArtifactRef,
		ArtifactHash:                 dependency.ValAManifest.ArtifactHash,
		EvidenceRefs:                 append([]string{}, dependency.ValAManifest.EvidenceRefs...),
		EvidenceHashRefs:             append([]string{}, dependency.ValAManifest.EvidenceHashRefs...),
		PolicyRef:                    dependency.ValAManifest.PolicyRef,
		PolicyVersion:                dependency.ValAManifest.PolicyVersion,
		PolicyHash:                   dependency.ValAManifest.PolicyHash,
		EngineVersion:                dependency.ValAManifest.EngineVersion,
		EngineHash:                   dependency.ValAManifest.EngineHash,
		SchemaVersion:                dependency.ValAManifest.SchemaVersion,
		SchemaHash:                   dependency.ValAManifest.SchemaHash,
		ClaimRefs:                    append([]string{}, dependency.ValAManifest.ClaimRefs...),
		GovernanceEventRefs:          append([]string{}, dependency.ValAManifest.GovernanceEventRefs...),
		ManifestPayloadHash:          dependency.ValAManifest.ManifestPayloadHash,
		CompatibilityProfileRef:      dependency.ValAManifest.CompatibilityProfileRef,
		ReplayMode:                   point12Val0ReplayModeOriginalContext,
		DeclaredCompatibilityContext: "compatibility_context_point12_valb_original",
		OriginalDecisionState:        "decision_state_allow",
		OriginalDecisionHash:         "sha256:9999999999999999999999999999999999999999999999999999999999999999",
		RedactionManifestRef:         dependency.ValAManifest.RedactionManifestRef,
		GeneratedAt:                  "2026-05-03T13:01:00Z",
		FreshnessContext:             "freshness_context_point12_valb_24h",
		SourceManifestIntegrityState: Point12ValAManifestIntegrityStateActive,
	}
	result := Point12ValBReplayResult{
		ReplayResultID:               "replay_result_point12_valb_001",
		ReplayRequestID:              request.ReplayRequestID,
		ProofPackID:                  request.ProofPackID,
		ManifestID:                   request.ManifestID,
		ReplayMode:                   request.ReplayMode,
		ReplayState:                  Point12ValBReplayResultStateActive,
		ReplayResultTaxonomy:         Point12Val0ReplayResultSameDecision,
		OriginalDecisionState:        request.OriginalDecisionState,
		ReplayedDecisionState:        request.OriginalDecisionState,
		MatchOriginal:                true,
		TamperDetected:               false,
		UnsupportedVersion:           false,
		InsufficientEvidence:         false,
		RedactionLimitations:         false,
		DecisionDriftExplanation:     "original_context_replay_matches_declared_decision",
		EvaluatedPolicyVersion:       request.PolicyVersion,
		EvaluatedEngineVersion:       request.EngineVersion,
		EvaluatedSchemaVersion:       request.SchemaVersion,
		EvidenceHashCheckResult:      point12ValBCheckResultActive,
		ManifestIntegrityCheckResult: point12ValBCheckResultActive,
		SignatureMetadataCheckResult: point12ValBCheckResultActive,
		CompatibilityCheckResult:     point12ValBCheckResultActive,
		ExternalAPIUsed:              false,
		PointPassEmitted:             false,
		ProjectionDisclaimer:         point12ValBProjectionDisclaimerBaseline,
		ReplayOutputClaims:           []string{"bounded claim"},
		CustomerVisibleExplanation:   "advisory projection only",
	}
	return Point12ValBFoundation{
		CurrentState:         Point12ValBStateActive,
		ProjectionDisclaimer: point12ValBProjectionDisclaimerBaseline,
		DependencyState:      Point12ValBDependencyStateActive,
		ReplayCommandState:   Point12ValBReplayCommandStateActive,
		ReplayRequestState:   Point12ValBReplayRequestStateActive,
		ReplayResultState:    Point12ValBReplayResultStateActive,
		Dependency:           dependency,
		ReplayCommand:        command,
		ReplayRequest:        request,
		ReplayResult:         result,
	}
}

func ComputePoint12ValBFoundation(model Point12ValBFoundation) Point12ValBFoundation {
	model.DependencyState = EvaluatePoint12ValBDependencyState(model.Dependency)
	commandState, _ := point12ValBReplayCommandStateAndReasons(model.ReplayCommand, model.Dependency)
	model.ReplayCommandState = commandState
	requestState, _ := point12ValBReplayRequestStateAndReasons(model.ReplayRequest, model.Dependency, model.ReplayCommand)
	model.ReplayRequestState = requestState
	resultState, resultReasons := point12ValBReplayResultStateAndReasons(model.ReplayResult, model.ReplayRequest, model.ReplayCommand, model.Dependency)
	model.ReplayResultState = resultState
	model.ReplayResult.ReplayState = resultState
	model.CurrentState = EvaluatePoint12ValBState(model)
	model.BlockingReasons = point12ValBBlockingReasons(model)
	model.ReviewPrerequisites = append([]string{}, model.Dependency.ReviewPrerequisites...)
	if resultState == Point12ValBReplayResultStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, resultReasons...)
	}
	return model
}

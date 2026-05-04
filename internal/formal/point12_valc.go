package formal

import "strings"

const (
	Point12ValCStateActive         = "point12_valc_audit_export_redaction_offline_verification_active"
	Point12ValCStateBlocked        = "point12_valc_audit_export_redaction_offline_verification_blocked"
	Point12ValCStateReviewRequired = "point12_valc_audit_export_redaction_offline_verification_review_required"

	Point12ValCDependencyStateActive         = "point12_valc_dependency_active"
	Point12ValCDependencyStateBlocked        = "point12_valc_dependency_blocked"
	Point12ValCDependencyStateReviewRequired = "point12_valc_dependency_review_required"

	Point12ValCRedactionManifestStateActive         = "point12_valc_redaction_manifest_active"
	Point12ValCRedactionManifestStateBlocked        = "point12_valc_redaction_manifest_blocked"
	Point12ValCRedactionManifestStateReviewRequired = "point12_valc_redaction_manifest_review_required"

	Point12ValCOfflineBundleStateActive              = "point12_valc_offline_bundle_active"
	Point12ValCOfflineBundleStateBlocked             = "point12_valc_offline_bundle_blocked"
	Point12ValCOfflineBundleStateReviewRequired      = "point12_valc_offline_bundle_review_required"
	Point12ValCOfflineBundleStateUnsupported         = "point12_valc_offline_bundle_unsupported"
	Point12ValCOfflineBundleStatePartialAdvisoryOnly = "point12_valc_offline_bundle_partial_advisory_only"
	Point12ValCOfflineBundleStateRedactedLimitations = "point12_valc_offline_bundle_redacted_limitations"

	Point12ValCPublicPrivateBoundaryStateActive  = "point12_valc_public_private_boundary_active"
	Point12ValCPublicPrivateBoundaryStateBlocked = "point12_valc_public_private_boundary_blocked"
)

const (
	point12ValCWaveID                        = "val_c"
	point12ValCPreviousWaveID                = point12ValBWaveID
	point12ValCProjectionDisclaimerBaseline  = "projection_only not_canonical_truth point12_valc_audit_export_redaction_offline_verification"
	point12ValCDependencySnapshotRefBaseline = "dependency_snapshot_point12_valc_valb_computed_001"

	point12ValCExportKindAuditReadyJSON           = "audit_ready_json"
	point12ValCExportKindAuditReadyStaticMetadata = "audit_ready_static_report_metadata"
	point12ValCExportKindVerifierPackageMetadata  = "verifier_package_metadata"
	point12ValCExportKindCustomerReviewPackage    = "customer_review_package_metadata"

	point12ValCExportScopeTenantScoped   = "tenant_scoped"
	point12ValCExportScopeAuditorReview  = "auditor_review"
	point12ValCExportScopeCustomerReview = "customer_review"
	point12ValCExportScopeVerifier       = "verifier_package"

	point12ValCExportAudienceInternalAudit = "internal_audit"
	point12ValCExportAudienceAuditor       = "auditor"
	point12ValCExportAudienceCustomer      = "customer"
	point12ValCExportAudienceVerifier      = "verifier"

	point12ValCClassificationInternalOnly      = "internal_only"
	point12ValCClassificationTenantPrivate     = "tenant_private"
	point12ValCClassificationAuditorRestricted = "auditor_restricted"
	point12ValCClassificationCustomerRedacted  = "customer_redacted"
	point12ValCClassificationPublicRedacted    = "public_redacted"

	Point12ValCExportStateReady               = "export_ready"
	Point12ValCExportStatePartialAdvisory     = "partial_advisory_export"
	Point12ValCExportStateBlocked             = "export_blocked"
	Point12ValCExportStateRedactedLimitations = Point12Val0ReplayResultRedactedLimitations
	Point12ValCExportStateInsufficient        = Point12Val0ReplayResultInsufficientEvidence
	Point12ValCExportStateUnsupported         = Point12Val0ReplayResultUnsupportedVersion
	Point12ValCExportStateTampered            = Point12Val0ReplayResultTamperDetected
	Point12ValCExportStateTenantMismatch      = "tenant_scope_mismatch"
	Point12ValCExportStateBoundaryViolation   = "public_private_boundary_violation"
	Point12ValCExportStateRetentionMissing    = "retention_missing"
	Point12ValCExportStateProjectionOnly      = "projection_only"
	Point12ValCExportStateReviewRequired      = "review_required"

	Point12ValCRedactionImpactNoDecisionImpact = "no_decision_impact"
	Point12ValCRedactionImpactRedactedLimits   = Point12Val0ReplayResultRedactedLimitations
	Point12ValCRedactionImpactInsufficient     = Point12Val0ReplayResultInsufficientEvidence
	Point12ValCRedactionImpactBlockedReplay    = Point12Val0ReplayResultBlockedReplay
	Point12ValCRedactionImpactPartialAdvisory  = "partial_advisory_only"
	Point12ValCRedactionImpactExportBlocked    = Point12ValCExportStateBlocked
	Point12ValCRedactionImpactReviewRequired   = "review_required"
)

type Point12ValCValBReviewContext struct {
	SnapshotFromComputedOutput   bool     `json:"snapshot_from_computed_output"`
	ValBPrematurePoint12PassSeen bool     `json:"valb_premature_point12_pass_seen"`
	ReviewPrerequisites          []string `json:"review_prerequisites,omitempty"`
}

type Point12ValCDependencySnapshot struct {
	ValBCurrentState             string                                 `json:"valb_current_state"`
	ValBDependencyState          string                                 `json:"valb_dependency_state"`
	ValBReplayCommandState       string                                 `json:"valb_replay_command_state"`
	ValBReplayRequestState       string                                 `json:"valb_replay_request_state"`
	ValBReplayResultState        string                                 `json:"valb_replay_result_state"`
	ValBPointID                  string                                 `json:"valb_point_id"`
	ValBWaveID                   string                                 `json:"valb_wave_id"`
	ValBManifestIntegrityResult  string                                 `json:"valb_manifest_integrity_check_result"`
	ValBSignatureMetadataResult  string                                 `json:"valb_signature_metadata_check_result"`
	ValBCompatibilityResult      string                                 `json:"valb_compatibility_check_result"`
	ValBReplayMode               string                                 `json:"valb_replay_mode"`
	ValBReplayTaxonomy           string                                 `json:"valb_replay_taxonomy"`
	ValBExternalAPIUsed          bool                                   `json:"valb_external_api_used"`
	ValBPointPassEmitted         bool                                   `json:"valb_point_pass_emitted"`
	Val0RedactionBoundaryState   string                                 `json:"val0_redaction_boundary_state"`
	ProjectionDisclaimer         string                                 `json:"projection_disclaimer"`
	SnapshotRef                  string                                 `json:"snapshot_ref"`
	SnapshotFromComputedOutput   bool                                   `json:"snapshot_from_computed_output"`
	ValBPrematurePoint12PassSeen bool                                   `json:"valb_premature_point12_pass_seen"`
	ReviewPrerequisites          []string                               `json:"review_prerequisites,omitempty"`
	ValAManifest                 Point12ValASignedProofPackManifestCore `json:"vala_manifest"`
	ValBReplayCommand            Point12ValBReplayCommandContract       `json:"valb_replay_command"`
	ValBReplayRequest            Point12ValBReplayRequest               `json:"valb_replay_request"`
	ValBReplayResult             Point12ValBReplayResult                `json:"valb_replay_result"`
}

type Point12ValCAuditExportBundle struct {
	ExportID                    string   `json:"export_id"`
	ExportKind                  string   `json:"export_kind"`
	ProofPackID                 string   `json:"proof_pack_id"`
	ManifestID                  string   `json:"manifest_id"`
	ReplayResultID              string   `json:"replay_result_id"`
	DecisionID                  string   `json:"decision_id"`
	TenantScope                 string   `json:"tenant_scope"`
	ArtifactRef                 string   `json:"artifact_ref"`
	ArtifactHash                string   `json:"artifact_hash"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs            []string `json:"evidence_hash_refs,omitempty"`
	PolicyRef                   string   `json:"policy_ref"`
	PolicyVersion               string   `json:"policy_version"`
	PolicyHash                  string   `json:"policy_hash"`
	EngineVersion               string   `json:"engine_version"`
	EngineHash                  string   `json:"engine_hash"`
	SchemaVersion               string   `json:"schema_version"`
	SchemaHash                  string   `json:"schema_hash"`
	ClaimRefs                   []string `json:"claim_refs,omitempty"`
	GovernanceEventRefs         []string `json:"governance_event_refs,omitempty"`
	CompatibilityProfileRef     string   `json:"compatibility_profile_ref"`
	RedactionManifestRef        string   `json:"redaction_manifest_ref"`
	OfflineBundleRef            string   `json:"offline_bundle_ref"`
	ManifestPayloadHash         string   `json:"manifest_payload_hash"`
	SignatureMetadataRef        string   `json:"signature_metadata_ref"`
	GeneratedAt                 string   `json:"generated_at"`
	FreshnessWindow             string   `json:"freshness_window"`
	ExportScope                 string   `json:"export_scope"`
	ExportAudience              string   `json:"export_audience"`
	PublicPrivateClassification string   `json:"public_private_classification"`
	RetentionClassRef           string   `json:"retention_class_ref"`
	RetentionOwnerRef           string   `json:"retention_owner_ref"`
	DisposalPathRef             string   `json:"disposal_path_ref"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
	Limitations                 []string `json:"limitations,omitempty"`
	AdvisoryOnly                bool     `json:"advisory_only"`
	NoOverclaimState            string   `json:"no_overclaim_state"`
	TenantBoundaryState         string   `json:"tenant_boundary_state"`
	RedactionImpactState        string   `json:"redaction_impact_state"`
	ExportState                 string   `json:"export_state"`
	ExportOutputClaims          []string `json:"export_output_claims,omitempty"`
	CustomerVisibleSummary      string   `json:"customer_visible_summary"`
}

type Point12ValCRedactionManifest struct {
	RedactionManifestID                 string   `json:"redaction_manifest_id"`
	ProofPackID                         string   `json:"proof_pack_id"`
	ExportID                            string   `json:"export_id"`
	TenantScope                         string   `json:"tenant_scope"`
	RedactedFields                      []string `json:"redacted_fields,omitempty"`
	RedactionReasons                    []string `json:"redaction_reasons,omitempty"`
	RedactionApproverRef                string   `json:"redaction_approver_ref"`
	RedactionApprovalEventRef           string   `json:"redaction_approval_event_ref"`
	RedactionPolicyRef                  string   `json:"redaction_policy_ref"`
	RedactionPolicyVersion              string   `json:"redaction_policy_version"`
	RedactionAffectsDecision            bool     `json:"redaction_affects_decision"`
	RedactionAffectsReplay              bool     `json:"redaction_affects_replay"`
	RedactionAffectsEvidenceHashes      bool     `json:"redaction_affects_evidence_hashes"`
	RedactionAffectsPolicyContext       bool     `json:"redaction_affects_policy_context"`
	RedactionAffectsClaimContext        bool     `json:"redaction_affects_claim_context"`
	RedactionAffectsGovernanceContext   bool     `json:"redaction_affects_governance_context"`
	PostRedactionResult                 string   `json:"post_redaction_result"`
	MinimumSafeClaimAfterRedaction      string   `json:"minimum_safe_claim_after_redaction"`
	DisallowedClaimsAfterRedaction      []string `json:"disallowed_claims_after_redaction,omitempty"`
	SurvivingClaimsAfterRedaction       []string `json:"surviving_claims_after_redaction,omitempty"`
	CustomerVisibleClaimsAfterRedaction []string `json:"customer_visible_claims_after_redaction,omitempty"`
	ExportedClaimsAfterRedaction        []string `json:"exported_claims_after_redaction,omitempty"`
	ReplayResultClaims                  []string `json:"replay_result_claims,omitempty"`
	RedactionSummary                    string   `json:"redaction_summary"`
	PartialOrAdvisoryOnly               bool     `json:"partial_or_advisory_only"`
	Limitations                         []string `json:"limitations,omitempty"`
	GeneratedAt                         string   `json:"generated_at"`
	RetentionClassRef                   string   `json:"retention_class_ref"`
	RetentionOwnerRef                   string   `json:"retention_owner_ref"`
	DisposalPathRef                     string   `json:"disposal_path_ref"`
}

type Point12ValCRedactionImpactVerdict struct {
	RedactionImpactID                string   `json:"redaction_impact_id"`
	RedactionManifestID              string   `json:"redaction_manifest_id"`
	AffectsDecision                  bool     `json:"affects_decision"`
	AffectsReplay                    bool     `json:"affects_replay"`
	DecisiveEvidenceRemoved          bool     `json:"decisive_evidence_removed"`
	DecisivePolicyContextRemoved     bool     `json:"decisive_policy_context_removed"`
	DecisiveClaimContextRemoved      bool     `json:"decisive_claim_context_removed"`
	DecisiveGovernanceContextRemoved bool     `json:"decisive_governance_context_removed"`
	PostRedactionResult              string   `json:"post_redaction_result"`
	MinimumSafeClaimAfterRedaction   string   `json:"minimum_safe_claim_after_redaction"`
	AllowedExportState               string   `json:"allowed_export_state"`
	DisallowedExportState            string   `json:"disallowed_export_state"`
	Limitations                      []string `json:"limitations,omitempty"`
	BlocksFullExport                 bool     `json:"blocks_full_export"`
	RequiresPartialAdvisoryExport    bool     `json:"requires_partial_advisory_export"`
	RequiresCustomerReview           bool     `json:"requires_customer_review"`
	RedactionImpactState             string   `json:"redaction_impact_state"`
}

type Point12ValCOfflineVerificationBundle struct {
	OfflineBundleID             string   `json:"offline_bundle_id"`
	ProofPackID                 string   `json:"proof_pack_id"`
	ManifestID                  string   `json:"manifest_id"`
	ReplayRequestID             string   `json:"replay_request_id"`
	ReplayResultID              string   `json:"replay_result_id"`
	TenantScope                 string   `json:"tenant_scope"`
	ArtifactRef                 string   `json:"artifact_ref"`
	ArtifactHash                string   `json:"artifact_hash"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs            []string `json:"evidence_hash_refs,omitempty"`
	PolicyRef                   string   `json:"policy_ref"`
	PolicyVersion               string   `json:"policy_version"`
	PolicyHash                  string   `json:"policy_hash"`
	EngineVersion               string   `json:"engine_version"`
	EngineHash                  string   `json:"engine_hash"`
	SchemaVersion               string   `json:"schema_version"`
	SchemaHash                  string   `json:"schema_hash"`
	ManifestPayloadHash         string   `json:"manifest_payload_hash"`
	SignatureMetadataRef        string   `json:"signature_metadata_ref"`
	DetachedSignatureRef        string   `json:"detached_signature_ref"`
	CompatibilityProfileRef     string   `json:"compatibility_profile_ref"`
	RedactionManifestRef        string   `json:"redaction_manifest_ref"`
	VerificationPolicyRef       string   `json:"verification_policy_ref"`
	GeneratedAt                 string   `json:"generated_at"`
	BundleFormatVersion         string   `json:"bundle_format_version"`
	SupportedVerifierVersions   []string `json:"supported_verifier_versions,omitempty"`
	RequestedVerifierVersion    string   `json:"requested_verifier_version"`
	NoExternalAPIRequired       bool     `json:"no_external_api_required"`
	ExternalAPIUsed             bool     `json:"external_api_used"`
	ContainsPrivateData         bool     `json:"contains_private_data"`
	PublicPrivateClassification string   `json:"public_private_classification"`
	RetentionClassRef           string   `json:"retention_class_ref"`
	RetentionOwnerRef           string   `json:"retention_owner_ref"`
	DisposalPathRef             string   `json:"disposal_path_ref"`
	Limitations                 []string `json:"limitations,omitempty"`
	OfflineOutputClaims         []string `json:"offline_output_claims,omitempty"`
	CustomerVisibleExplanation  string   `json:"customer_visible_explanation"`
	OfflineState                string   `json:"offline_state"`
}

type Point12ValCPublicPrivateBoundary struct {
	BoundaryID            string   `json:"boundary_id"`
	TenantScope           string   `json:"tenant_scope"`
	ExportID              string   `json:"export_id"`
	OfflineBundleID       string   `json:"offline_bundle_id"`
	ExportedFields        []string `json:"exported_fields,omitempty"`
	PublicFields          []string `json:"public_fields,omitempty"`
	PrivateFields         []string `json:"private_fields,omitempty"`
	RedactedFields        []string `json:"redacted_fields,omitempty"`
	SensitiveFields       []string `json:"sensitive_fields,omitempty"`
	CustomerVisibleFields []string `json:"customer_visible_fields,omitempty"`
	AuditorVisibleFields  []string `json:"auditor_visible_fields,omitempty"`
	InternalOnlyFields    []string `json:"internal_only_fields,omitempty"`
	Classification        string   `json:"classification"`
	DataResidencyRef      string   `json:"data_residency_ref"`
	AllowedAudience       string   `json:"allowed_audience"`
	BoundaryState         string   `json:"boundary_state"`
}

type Point12ValCFoundation struct {
	CurrentState               string                               `json:"current_state"`
	BlockingReasons            []string                             `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites        []string                             `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer       string                               `json:"projection_disclaimer"`
	DependencyState            string                               `json:"dependency_state"`
	RedactionManifestState     string                               `json:"redaction_manifest_state"`
	RedactionImpactState       string                               `json:"redaction_impact_state"`
	OfflineBundleState         string                               `json:"offline_bundle_state"`
	PublicPrivateBoundaryState string                               `json:"public_private_boundary_state"`
	ExportState                string                               `json:"export_state"`
	Dependency                 Point12ValCDependencySnapshot        `json:"dependency"`
	ExportBundle               Point12ValCAuditExportBundle         `json:"export_bundle"`
	RedactionManifest          Point12ValCRedactionManifest         `json:"redaction_manifest"`
	RedactionImpactVerdict     Point12ValCRedactionImpactVerdict    `json:"redaction_impact_verdict"`
	OfflineBundle              Point12ValCOfflineVerificationBundle `json:"offline_bundle"`
	PublicPrivateBoundary      Point12ValCPublicPrivateBoundary     `json:"public_private_boundary"`
}

func point12ValCDependencySnapshotRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"dependency_snapshot_", "valb_snapshot_"})
}

func point12ValCExportRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"export_"})
}

func point12ValCRedactionImpactRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"redaction_impact_"})
}

func point12ValCOfflineBundleRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"offline_bundle_"})
}

func point12ValCBoundaryRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"boundary_", "public_private_boundary_"})
}

func point12ValCRetentionOwnerRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"retention_owner_"})
}

func point12ValCDisposalPathRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"disposal_path_"})
}

func point12ValCVerificationPolicyRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"verification_policy_"})
}

func point12ValCDataResidencyRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"data_residency_"})
}

func point12ValCExportKindValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValCExportKindAuditReadyJSON,
		point12ValCExportKindAuditReadyStaticMetadata,
		point12ValCExportKindVerifierPackageMetadata,
		point12ValCExportKindCustomerReviewPackage,
	}, value)
}

func point12ValCExportScopeValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValCExportScopeTenantScoped,
		point12ValCExportScopeAuditorReview,
		point12ValCExportScopeCustomerReview,
		point12ValCExportScopeVerifier,
	}, value)
}

func point12ValCExportAudienceValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValCExportAudienceInternalAudit,
		point12ValCExportAudienceAuditor,
		point12ValCExportAudienceCustomer,
		point12ValCExportAudienceVerifier,
	}, value)
}

func point12ValCPublicPrivateClassificationValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValCClassificationInternalOnly,
		point12ValCClassificationTenantPrivate,
		point12ValCClassificationAuditorRestricted,
		point12ValCClassificationCustomerRedacted,
		point12ValCClassificationPublicRedacted,
	}, value)
}

func point12ValCExportStates() []string {
	return []string{
		Point12ValCExportStateReady,
		Point12ValCExportStatePartialAdvisory,
		Point12ValCExportStateBlocked,
		Point12ValCExportStateRedactedLimitations,
		Point12ValCExportStateInsufficient,
		Point12ValCExportStateUnsupported,
		Point12ValCExportStateTampered,
		Point12ValCExportStateTenantMismatch,
		Point12ValCExportStateBoundaryViolation,
		Point12ValCExportStateRetentionMissing,
		Point12ValCExportStateProjectionOnly,
		Point12ValCExportStateReviewRequired,
	}
}

func point12ValCDependencyCheckResultBlocked(value string) bool {
	trimmed := strings.TrimSpace(value)
	if !point12ValBCheckResultValid(trimmed) {
		return true
	}
	return point11Val0ContainsTrimmed([]string{
		point12ValBCheckResultMismatch,
		point12ValBCheckResultTampered,
		point12ValBCheckResultMissing,
		point12ValBCheckResultBlocked,
	}, trimmed)
}

func point12ValCDependencyCheckResultReviewRequired(value string) bool {
	return strings.TrimSpace(value) == point12ValBCheckResultUnsupported
}

func point12ValCRedactionImpactStates() []string {
	return []string{
		Point12ValCRedactionImpactNoDecisionImpact,
		Point12ValCRedactionImpactRedactedLimits,
		Point12ValCRedactionImpactInsufficient,
		Point12ValCRedactionImpactBlockedReplay,
		Point12ValCRedactionImpactPartialAdvisory,
		Point12ValCRedactionImpactExportBlocked,
		Point12ValCRedactionImpactReviewRequired,
	}
}

func point12ValCOfflineBundleStates() []string {
	return []string{
		Point12ValCOfflineBundleStateActive,
		Point12ValCOfflineBundleStateBlocked,
		Point12ValCOfflineBundleStateReviewRequired,
		Point12ValCOfflineBundleStateUnsupported,
		Point12ValCOfflineBundleStatePartialAdvisoryOnly,
		Point12ValCOfflineBundleStateRedactedLimitations,
	}
}

func point12ValCCustomerFacingAudience(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValCExportAudienceCustomer,
	}, value)
}

func point12ValCStringFieldValid(value string) bool {
	return point11Val0IdentityValueValid(strings.TrimSpace(value))
}

func point12ValCStringFieldListValid(values []string) bool {
	return point12Val0StringListValid(values, point12ValCStringFieldValid)
}

func point12ValCStringFieldListSubset(subset, superset []string) bool {
	if len(subset) == 0 {
		return true
	}
	seen := map[string]struct{}{}
	for _, item := range superset {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			seen[trimmed] = struct{}{}
		}
	}
	for _, item := range subset {
		if _, ok := seen[strings.TrimSpace(item)]; !ok {
			return false
		}
	}
	return true
}

func point12ValCFieldListsOverlap(left, right []string) bool {
	for _, l := range left {
		for _, r := range right {
			if strings.TrimSpace(l) != "" && strings.TrimSpace(l) == strings.TrimSpace(r) {
				return true
			}
		}
	}
	return false
}

func point12ValCAllExportedFieldsClassified(model Point12ValCPublicPrivateBoundary) bool {
	classified := append([]string{}, model.PublicFields...)
	classified = append(classified, model.PrivateFields...)
	classified = append(classified, model.RedactedFields...)
	classified = append(classified, model.SensitiveFields...)
	classified = append(classified, model.InternalOnlyFields...)
	return point12ValCStringFieldListSubset(model.ExportedFields, classified)
}

func point12ValCStringMentionedInTexts(fields []string, texts ...string) bool {
	for _, field := range fields {
		needle := strings.TrimSpace(field)
		if needle == "" {
			continue
		}
		for _, text := range texts {
			if strings.Contains(strings.ToLower(text), strings.ToLower(needle)) {
				return true
			}
		}
	}
	return false
}

func point12ValCDependencyReviewContextModel() Point12ValCValBReviewContext {
	return Point12ValCValBReviewContext{
		SnapshotFromComputedOutput: true,
	}
}

func SnapshotPoint12ValCDependencyFromComputedValB(valB Point12ValBFoundation, review Point12ValCValBReviewContext) Point12ValCDependencySnapshot {
	reviewPrerequisites := append([]string{}, valB.ReviewPrerequisites...)
	reviewPrerequisites = append(reviewPrerequisites, review.ReviewPrerequisites...)
	return Point12ValCDependencySnapshot{
		ValBCurrentState:             valB.CurrentState,
		ValBDependencyState:          valB.DependencyState,
		ValBReplayCommandState:       valB.ReplayCommandState,
		ValBReplayRequestState:       valB.ReplayRequestState,
		ValBReplayResultState:        valB.ReplayResultState,
		ValBPointID:                  valB.Dependency.ValAPointID,
		ValBWaveID:                   point12ValBWaveID,
		ValBManifestIntegrityResult:  valB.ReplayResult.ManifestIntegrityCheckResult,
		ValBSignatureMetadataResult:  valB.ReplayResult.SignatureMetadataCheckResult,
		ValBCompatibilityResult:      valB.ReplayResult.CompatibilityCheckResult,
		ValBReplayMode:               valB.ReplayRequest.ReplayMode,
		ValBReplayTaxonomy:           valB.ReplayResult.ReplayResultTaxonomy,
		ValBExternalAPIUsed:          valB.ReplayResult.ExternalAPIUsed,
		ValBPointPassEmitted:         valB.ReplayResult.PointPassEmitted,
		Val0RedactionBoundaryState:   valB.Dependency.Val0RedactionBoundaryState,
		ProjectionDisclaimer:         valB.ProjectionDisclaimer,
		SnapshotRef:                  point12ValCDependencySnapshotRefBaseline,
		SnapshotFromComputedOutput:   review.SnapshotFromComputedOutput,
		ValBPrematurePoint12PassSeen: review.ValBPrematurePoint12PassSeen,
		ReviewPrerequisites:          reviewPrerequisites,
		ValAManifest:                 valB.Dependency.ValAManifest,
		ValBReplayCommand:            valB.ReplayCommand,
		ValBReplayRequest:            valB.ReplayRequest,
		ValBReplayResult:             valB.ReplayResult,
	}
}

func point12ValCDependencySnapshotModel() Point12ValCDependencySnapshot {
	valB := ComputePoint12ValBFoundation(Point12ValBFoundationModel())
	return SnapshotPoint12ValCDependencyFromComputedValB(valB, point12ValCDependencyReviewContextModel())
}

func EvaluatePoint12ValCDependencyState(model Point12ValCDependencySnapshot) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!model.SnapshotFromComputedOutput ||
		!point12ValCDependencySnapshotRefValid(model.SnapshotRef) ||
		strings.TrimSpace(model.ValBPointID) != point12Val0PointID ||
		strings.TrimSpace(model.ValBWaveID) != point12ValBWaveID ||
		!point12Val0ReplayModeValid(model.ValBReplayMode) ||
		model.ValBExternalAPIUsed ||
		model.ValBPointPassEmitted ||
		model.ValBPrematurePoint12PassSeen ||
		point12Val0ContainsPrematurePassToken(
			model.ValBReplayCommand.CommandID,
			model.ValBReplayRequest.ReplayRequestID,
			model.ValBReplayResult.ReplayResultID,
			model.ValBReplayResult.CustomerVisibleExplanation,
			strings.Join(model.ValBReplayResult.ReplayOutputClaims, " "),
		) {
		return Point12ValCDependencyStateBlocked
	}
	if !point12Val0ProofPackRefValid(model.ValAManifest.ProofPackID) ||
		!point12ValAManifestRefValid(model.ValAManifest.ManifestID) ||
		!point12ValBReplayRequestRefValid(model.ValBReplayRequest.ReplayRequestID) ||
		!point12ValBReplayResultRefValid(model.ValBReplayResult.ReplayResultID) {
		return Point12ValCDependencyStateBlocked
	}
	if strings.TrimSpace(model.ValBCurrentState) == Point12ValBStateBlocked ||
		strings.TrimSpace(model.ValBDependencyState) == Point12ValBDependencyStateBlocked ||
		strings.TrimSpace(model.ValBReplayCommandState) == Point12ValBReplayCommandStateBlocked ||
		strings.TrimSpace(model.ValBReplayRequestState) == Point12ValBReplayRequestStateBlocked ||
		strings.TrimSpace(model.ValBReplayResultState) == Point12ValBReplayResultStateBlocked ||
		point12ValCDependencyCheckResultBlocked(model.ValBManifestIntegrityResult) ||
		point12ValCDependencyCheckResultBlocked(model.ValBSignatureMetadataResult) ||
		point12ValCDependencyCheckResultBlocked(model.ValBCompatibilityResult) ||
		strings.TrimSpace(model.ValBReplayTaxonomy) == Point12Val0ReplayResultTamperDetected ||
		strings.TrimSpace(model.ValBReplayTaxonomy) == Point12Val0ReplayResultBlockedReplay ||
		strings.TrimSpace(model.Val0RedactionBoundaryState) == Point12Val0RedactionBoundaryStateBlocked {
		return Point12ValCDependencyStateBlocked
	}
	if strings.TrimSpace(model.ValBCurrentState) == Point12ValBStateReviewRequired ||
		strings.TrimSpace(model.ValBDependencyState) == Point12ValBDependencyStateReviewRequired ||
		strings.TrimSpace(model.ValBReplayResultState) == Point12ValBReplayResultStateReviewRequired ||
		point12ValCDependencyCheckResultReviewRequired(model.ValBManifestIntegrityResult) ||
		point12ValCDependencyCheckResultReviewRequired(model.ValBSignatureMetadataResult) ||
		point12ValCDependencyCheckResultReviewRequired(model.ValBCompatibilityResult) ||
		strings.TrimSpace(model.ValBReplayTaxonomy) == Point12Val0ReplayResultUnsupportedVersion ||
		strings.TrimSpace(model.ValBReplayTaxonomy) == Point12Val0ReplayResultInsufficientEvidence ||
		strings.TrimSpace(model.ValBReplayTaxonomy) == Point12Val0ReplayResultRedactedLimitations ||
		len(model.ReviewPrerequisites) > 0 {
		return Point12ValCDependencyStateReviewRequired
	}
	if strings.TrimSpace(model.ValBCurrentState) != Point12ValBStateActive ||
		strings.TrimSpace(model.ValBDependencyState) != Point12ValBDependencyStateActive ||
		strings.TrimSpace(model.ValBReplayCommandState) != Point12ValBReplayCommandStateActive ||
		strings.TrimSpace(model.ValBReplayRequestState) != Point12ValBReplayRequestStateActive ||
		strings.TrimSpace(model.ValBReplayResultState) != Point12ValBReplayResultStateActive ||
		strings.TrimSpace(model.ValBManifestIntegrityResult) != point12ValBCheckResultActive ||
		strings.TrimSpace(model.ValBSignatureMetadataResult) != point12ValBCheckResultActive ||
		strings.TrimSpace(model.ValBCompatibilityResult) != point12ValBCheckResultActive {
		return Point12ValCDependencyStateBlocked
	}
	return Point12ValCDependencyStateActive
}

func point12ValCRedactionManifestStateAndReasons(model Point12ValCRedactionManifest, dependency Point12ValCDependencySnapshot, export Point12ValCAuditExportBundle) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12Val0RedactionManifestRefValid(model.RedactionManifestID) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValCExportRefValid(model.ExportID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12Val0RedactionFieldValuesValid(model.RedactedFields) ||
		!point12Val0OptionalStringListValid(model.RedactionReasons, point11Val0IdentityValueValid) ||
		!point12Val0PolicyRefValid(model.RedactionPolicyRef) ||
		!point12Val0VersionIdentityValid(model.RedactionPolicyVersion) ||
		!point11Val0ContainsTrimmed(point12Val0ReplayResults(), model.PostRedactionResult) ||
		!point12Val0OptionalClaimTextListValid(model.DisallowedClaimsAfterRedaction) ||
		!point12Val0OptionalClaimTextListValid(model.SurvivingClaimsAfterRedaction) ||
		!point12Val0OptionalClaimTextListValid(model.CustomerVisibleClaimsAfterRedaction) ||
		!point12Val0OptionalClaimTextListValid(model.ExportedClaimsAfterRedaction) ||
		!point12Val0OptionalClaimTextListValid(model.ReplayResultClaims) ||
		!point12Val0OptionalStringListValid(model.Limitations, point11Val0IdentityValueValid) ||
		!point11Val0ValidTimestamp(model.GeneratedAt) ||
		!point12Val0RetentionClassRefValid(model.RetentionClassRef) ||
		!point12ValCRetentionOwnerRefValid(model.RetentionOwnerRef) ||
		!point12ValCDisposalPathRefValid(model.DisposalPathRef) {
		blockedReasons = append(blockedReasons, "redaction_manifest_identity_or_metadata_invalid")
	}
	if strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(dependency.ValAManifest.ProofPackID) ||
		strings.TrimSpace(model.RedactionManifestID) != strings.TrimSpace(dependency.ValBReplayRequest.RedactionManifestRef) ||
		strings.TrimSpace(model.ExportID) != strings.TrimSpace(export.ExportID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValBReplayRequest.TenantScope) {
		blockedReasons = append(blockedReasons, "redaction_manifest_dependency_binding_mismatch")
	}
	if point12Val0ContainsPrematurePassToken(
		model.RedactionManifestID,
		model.RedactionApproverRef,
		model.RedactionApprovalEventRef,
		model.RedactionSummary,
		strings.Join(model.DisallowedClaimsAfterRedaction, " "),
		strings.Join(model.SurvivingClaimsAfterRedaction, " "),
		strings.Join(model.CustomerVisibleClaimsAfterRedaction, " "),
		strings.Join(model.ExportedClaimsAfterRedaction, " "),
		strings.Join(model.ReplayResultClaims, " "),
	) {
		blockedReasons = append(blockedReasons, "redaction_manifest_premature_point12_pass")
	}
	if len(model.RedactedFields) > 0 {
		if !point12Val0StringListValid(model.RedactionReasons, point11Val0IdentityValueValid) ||
			!point11Val0IdentityValueValid(model.RedactionApproverRef) ||
			!point12Val0GovernanceEventRefValid(model.RedactionApprovalEventRef) ||
			!model.PartialOrAdvisoryOnly {
			blockedReasons = append(blockedReasons, "redaction_manifest_approval_or_reason_missing")
		}
	}
	if strings.TrimSpace(model.MinimumSafeClaimAfterRedaction) != "" &&
		point12Val0ContainsForbiddenClaim(model.MinimumSafeClaimAfterRedaction) {
		blockedReasons = append(blockedReasons, "redaction_manifest_minimum_safe_claim_overclaim")
	}
	if point12Val0ContainsForbiddenClaim(
		strings.Join(model.SurvivingClaimsAfterRedaction, " "),
		strings.Join(model.CustomerVisibleClaimsAfterRedaction, " "),
		strings.Join(model.ExportedClaimsAfterRedaction, " "),
		strings.Join(model.ReplayResultClaims, " "),
	) {
		blockedReasons = append(blockedReasons, "redaction_manifest_surviving_claim_overclaim")
	}
	if point12Val0ClaimTextOverlap(model.DisallowedClaimsAfterRedaction, model.SurvivingClaimsAfterRedaction) ||
		point12Val0ClaimTextOverlap(model.DisallowedClaimsAfterRedaction, model.CustomerVisibleClaimsAfterRedaction) ||
		point12Val0ClaimTextOverlap(model.DisallowedClaimsAfterRedaction, model.ExportedClaimsAfterRedaction) ||
		point12Val0ClaimTextOverlap(model.DisallowedClaimsAfterRedaction, model.ReplayResultClaims) {
		blockedReasons = append(blockedReasons, "redaction_manifest_disallowed_claim_survives")
	}
	if (model.RedactionAffectsDecision || model.RedactionAffectsReplay ||
		model.RedactionAffectsEvidenceHashes || model.RedactionAffectsPolicyContext ||
		model.RedactionAffectsClaimContext || model.RedactionAffectsGovernanceContext) &&
		strings.TrimSpace(model.PostRedactionResult) != Point12Val0ReplayResultInsufficientEvidence &&
		strings.TrimSpace(model.PostRedactionResult) != Point12Val0ReplayResultBlockedReplay &&
		strings.TrimSpace(model.PostRedactionResult) != Point12Val0ReplayResultRedactedLimitations {
		blockedReasons = append(blockedReasons, "redaction_manifest_post_result_must_fail_closed")
	}
	if (model.RedactionAffectsDecision || model.RedactionAffectsReplay) && !model.PartialOrAdvisoryOnly {
		blockedReasons = append(blockedReasons, "redaction_manifest_decisive_redaction_requires_partial_advisory_only")
	}
	if (model.RedactionAffectsDecision || model.RedactionAffectsReplay) && len(model.Limitations) == 0 {
		reviewReasons = append(reviewReasons, "redaction_manifest_decisive_redaction_limitations_missing")
	}
	if len(blockedReasons) > 0 {
		return Point12ValCRedactionManifestStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValCRedactionManifestStateReviewRequired, reviewReasons
	}
	return Point12ValCRedactionManifestStateActive, nil
}

func EvaluatePoint12ValCRedactionManifestState(model Point12ValCRedactionManifest, dependency Point12ValCDependencySnapshot, export Point12ValCAuditExportBundle) string {
	state, _ := point12ValCRedactionManifestStateAndReasons(model, dependency, export)
	return state
}

func point12ValCRedactionImpactStateAndReasons(model Point12ValCRedactionImpactVerdict, manifest Point12ValCRedactionManifest, dependency Point12ValCDependencySnapshot) (string, []string) {
	reasons := []string{}
	if !point12ValCRedactionImpactRefValid(model.RedactionImpactID) ||
		!point12Val0RedactionManifestRefValid(model.RedactionManifestID) ||
		!point11Val0ContainsTrimmed(point12Val0ReplayResults(), model.PostRedactionResult) ||
		!point11Val0ContainsTrimmed(point12ValCExportStates(), model.AllowedExportState) ||
		!point11Val0ContainsTrimmed(point12ValCExportStates(), model.DisallowedExportState) ||
		!point12Val0OptionalStringListValid(model.Limitations, point11Val0IdentityValueValid) ||
		!point11Val0ContainsTrimmed(point12ValCRedactionImpactStates(), model.RedactionImpactState) {
		reasons = append(reasons, "redaction_impact_identity_or_metadata_invalid")
	}
	if point12Val0ContainsPrematurePassToken(
		model.RedactionImpactID,
		model.MinimumSafeClaimAfterRedaction,
		strings.Join(model.Limitations, " "),
	) {
		reasons = append(reasons, "redaction_impact_premature_point12_pass")
	}
	if strings.TrimSpace(model.RedactionManifestID) != strings.TrimSpace(manifest.RedactionManifestID) ||
		strings.TrimSpace(model.RedactionManifestID) != strings.TrimSpace(dependency.ValBReplayRequest.RedactionManifestRef) ||
		strings.TrimSpace(model.PostRedactionResult) != strings.TrimSpace(manifest.PostRedactionResult) ||
		strings.TrimSpace(model.MinimumSafeClaimAfterRedaction) != strings.TrimSpace(manifest.MinimumSafeClaimAfterRedaction) {
		reasons = append(reasons, "redaction_impact_manifest_binding_mismatch")
	}
	if strings.TrimSpace(model.MinimumSafeClaimAfterRedaction) != "" &&
		point12Val0ContainsForbiddenClaim(model.MinimumSafeClaimAfterRedaction) {
		reasons = append(reasons, "redaction_impact_minimum_safe_claim_overclaim")
	}
	decisiveRemoved := model.DecisiveEvidenceRemoved ||
		model.DecisivePolicyContextRemoved ||
		model.DecisiveClaimContextRemoved ||
		model.DecisiveGovernanceContextRemoved
	if decisiveRemoved || model.AffectsDecision {
		if strings.TrimSpace(model.RedactionImpactState) == Point12ValCRedactionImpactNoDecisionImpact {
			reasons = append(reasons, "redaction_impact_decisive_change_cannot_be_no_decision_impact")
		}
		if !model.BlocksFullExport && !model.RequiresPartialAdvisoryExport {
			reasons = append(reasons, "redaction_impact_decisive_change_requires_export_limitation")
		}
	}
	if model.AffectsReplay && strings.TrimSpace(model.RedactionImpactState) == Point12ValCRedactionImpactNoDecisionImpact {
		reasons = append(reasons, "redaction_impact_replay_change_requires_limitations")
	}
	if (model.AffectsDecision || model.AffectsReplay || decisiveRemoved) && len(model.Limitations) == 0 {
		reasons = append(reasons, "redaction_impact_limitations_missing")
	}
	if len(reasons) > 0 {
		return Point12ValCRedactionImpactReviewRequired, reasons
	}
	return strings.TrimSpace(model.RedactionImpactState), nil
}

func EvaluatePoint12ValCRedactionImpactState(model Point12ValCRedactionImpactVerdict, manifest Point12ValCRedactionManifest, dependency Point12ValCDependencySnapshot) string {
	state, _ := point12ValCRedactionImpactStateAndReasons(model, manifest, dependency)
	return state
}

func point12ValCOfflineBundleStateAndReasons(model Point12ValCOfflineVerificationBundle, dependency Point12ValCDependencySnapshot, redactionImpactState string) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValCOfflineBundleRefValid(model.OfflineBundleID) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point12ValBReplayRequestRefValid(model.ReplayRequestID) ||
		!point12ValBReplayResultRefValid(model.ReplayResultID) ||
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
		!point12Val0HashValid(model.ManifestPayloadHash) ||
		!point12ValASignatureMetadataRefValid(model.SignatureMetadataRef) ||
		!point12ValADetachedSignatureRefValid(model.DetachedSignatureRef) ||
		!point12Val0CompatibilityProfileRefValid(model.CompatibilityProfileRef) ||
		!point12Val0RedactionManifestRefValid(model.RedactionManifestRef) ||
		!point12ValCVerificationPolicyRefValid(model.VerificationPolicyRef) ||
		!point11Val0ValidTimestamp(model.GeneratedAt) ||
		!point12Val0VersionIdentityValid(model.BundleFormatVersion) ||
		!point12Val0StringListValid(model.SupportedVerifierVersions, point12Val0VersionIdentityValid) ||
		!point12Val0VersionIdentityValid(model.RequestedVerifierVersion) ||
		!point12ValCPublicPrivateClassificationValid(model.PublicPrivateClassification) ||
		!point12Val0RetentionClassRefValid(model.RetentionClassRef) ||
		!point12ValCRetentionOwnerRefValid(model.RetentionOwnerRef) ||
		!point12ValCDisposalPathRefValid(model.DisposalPathRef) ||
		!point12Val0OptionalStringListValid(model.Limitations, point11Val0IdentityValueValid) ||
		!point12Val0OptionalClaimTextListValid(model.OfflineOutputClaims) ||
		!point11Val0ContainsTrimmed(point12ValCOfflineBundleStates(), model.OfflineState) {
		blockedReasons = append(blockedReasons, "offline_bundle_identity_or_metadata_invalid")
	}
	if len(model.EvidenceRefs) != len(model.EvidenceHashRefs) {
		blockedReasons = append(blockedReasons, "offline_bundle_evidence_hash_alignment_invalid")
	}
	if !model.NoExternalAPIRequired {
		blockedReasons = append(blockedReasons, "offline_bundle_no_external_api_required_missing")
	}
	if model.ExternalAPIUsed {
		blockedReasons = append(blockedReasons, "offline_bundle_external_api_used")
	}
	if point12Val0ContainsPrematurePassToken(
		model.OfflineBundleID,
		model.ReplayResultID,
		model.CustomerVisibleExplanation,
		strings.Join(model.OfflineOutputClaims, " "),
	) {
		blockedReasons = append(blockedReasons, "offline_bundle_premature_point12_pass")
	}
	if point12Val0ContainsForbiddenClaim(strings.Join(model.OfflineOutputClaims, " "), model.CustomerVisibleExplanation) {
		blockedReasons = append(blockedReasons, "offline_bundle_overclaim_detected")
	}
	if strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(dependency.ValAManifest.ProofPackID) ||
		strings.TrimSpace(model.ManifestID) != strings.TrimSpace(dependency.ValAManifest.ManifestID) ||
		strings.TrimSpace(model.ReplayRequestID) != strings.TrimSpace(dependency.ValBReplayRequest.ReplayRequestID) ||
		strings.TrimSpace(model.ReplayResultID) != strings.TrimSpace(dependency.ValBReplayResult.ReplayResultID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValBReplayRequest.TenantScope) ||
		strings.TrimSpace(model.ArtifactRef) != strings.TrimSpace(dependency.ValBReplayRequest.ArtifactRef) ||
		strings.TrimSpace(model.ArtifactHash) != strings.TrimSpace(dependency.ValBReplayRequest.ArtifactHash) ||
		!point12Val0ExactStringSetMatch(model.EvidenceRefs, dependency.ValBReplayRequest.EvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.EvidenceHashRefs, dependency.ValBReplayRequest.EvidenceHashRefs) ||
		strings.TrimSpace(model.PolicyRef) != strings.TrimSpace(dependency.ValBReplayRequest.PolicyRef) ||
		strings.TrimSpace(model.PolicyVersion) != strings.TrimSpace(dependency.ValBReplayRequest.PolicyVersion) ||
		strings.TrimSpace(model.PolicyHash) != strings.TrimSpace(dependency.ValBReplayRequest.PolicyHash) ||
		strings.TrimSpace(model.EngineVersion) != strings.TrimSpace(dependency.ValBReplayRequest.EngineVersion) ||
		strings.TrimSpace(model.EngineHash) != strings.TrimSpace(dependency.ValBReplayRequest.EngineHash) ||
		strings.TrimSpace(model.SchemaVersion) != strings.TrimSpace(dependency.ValBReplayRequest.SchemaVersion) ||
		strings.TrimSpace(model.SchemaHash) != strings.TrimSpace(dependency.ValBReplayRequest.SchemaHash) ||
		strings.TrimSpace(model.ManifestPayloadHash) != strings.TrimSpace(dependency.ValBReplayRequest.ManifestPayloadHash) ||
		strings.TrimSpace(model.RedactionManifestRef) != strings.TrimSpace(dependency.ValBReplayRequest.RedactionManifestRef) ||
		strings.TrimSpace(model.SignatureMetadataRef) != strings.TrimSpace(dependency.ValAManifest.SignatureMetadataRef) ||
		strings.TrimSpace(model.CompatibilityProfileRef) != strings.TrimSpace(dependency.ValBReplayRequest.CompatibilityProfileRef) {
		blockedReasons = append(blockedReasons, "offline_bundle_dependency_binding_mismatch")
	}
	if !point11Val0ContainsTrimmed(model.SupportedVerifierVersions, model.RequestedVerifierVersion) {
		reviewReasons = append(reviewReasons, "offline_bundle_requested_verifier_version_unsupported")
	}
	if strings.TrimSpace(redactionImpactState) == Point12ValCRedactionImpactBlockedReplay ||
		strings.TrimSpace(redactionImpactState) == Point12ValCRedactionImpactInsufficient {
		if strings.TrimSpace(model.OfflineState) != Point12ValCOfflineBundleStatePartialAdvisoryOnly &&
			strings.TrimSpace(model.OfflineState) != Point12ValCOfflineBundleStateBlocked &&
			strings.TrimSpace(model.OfflineState) != Point12ValCOfflineBundleStateRedactedLimitations {
			blockedReasons = append(blockedReasons, "offline_bundle_decisive_redaction_requires_limitations")
		}
	}
	if strings.TrimSpace(dependency.ValBReplayTaxonomy) == Point12Val0ReplayResultTamperDetected ||
		strings.TrimSpace(dependency.ValBReplayTaxonomy) == Point12Val0ReplayResultBlockedReplay {
		blockedReasons = append(blockedReasons, "offline_bundle_tampered_or_blocked_replay_dependency")
	}
	if strings.TrimSpace(dependency.ValBReplayTaxonomy) == Point12Val0ReplayResultUnsupportedVersion &&
		strings.TrimSpace(model.OfflineState) != Point12ValCOfflineBundleStateUnsupported &&
		strings.TrimSpace(model.OfflineState) != Point12ValCOfflineBundleStateReviewRequired {
		reviewReasons = append(reviewReasons, "offline_bundle_unsupported_version_requires_unsupported_or_review")
	}
	if len(blockedReasons) > 0 {
		return Point12ValCOfflineBundleStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		if strings.TrimSpace(model.OfflineState) == Point12ValCOfflineBundleStateUnsupported {
			return Point12ValCOfflineBundleStateUnsupported, reviewReasons
		}
		return Point12ValCOfflineBundleStateReviewRequired, reviewReasons
	}
	return strings.TrimSpace(model.OfflineState), nil
}

func EvaluatePoint12ValCOfflineBundleState(model Point12ValCOfflineVerificationBundle, dependency Point12ValCDependencySnapshot, redactionImpactState string) string {
	state, _ := point12ValCOfflineBundleStateAndReasons(model, dependency, redactionImpactState)
	return state
}

func point12ValCPublicPrivateBoundaryStateAndReasons(model Point12ValCPublicPrivateBoundary, dependency Point12ValCDependencySnapshot, export Point12ValCAuditExportBundle, offline Point12ValCOfflineVerificationBundle, redaction Point12ValCRedactionManifest) (string, []string) {
	reasons := []string{}
	if !point12ValCBoundaryRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12ValCExportRefValid(model.ExportID) ||
		!point12ValCOfflineBundleRefValid(model.OfflineBundleID) ||
		!point12ValCStringFieldListValid(model.ExportedFields) ||
		!point12Val0OptionalStringListValid(model.PublicFields, point12ValCStringFieldValid) ||
		!point12Val0OptionalStringListValid(model.PrivateFields, point12ValCStringFieldValid) ||
		!point12Val0OptionalStringListValid(model.RedactedFields, point12ValCStringFieldValid) ||
		!point12Val0OptionalStringListValid(model.SensitiveFields, point12ValCStringFieldValid) ||
		!point12Val0OptionalStringListValid(model.CustomerVisibleFields, point12ValCStringFieldValid) ||
		!point12Val0OptionalStringListValid(model.AuditorVisibleFields, point12ValCStringFieldValid) ||
		!point12Val0OptionalStringListValid(model.InternalOnlyFields, point12ValCStringFieldValid) ||
		!point12ValCPublicPrivateClassificationValid(model.Classification) ||
		!point12ValCDataResidencyRefValid(model.DataResidencyRef) ||
		!point12ValCExportAudienceValid(model.AllowedAudience) ||
		!point11Val0ContainsTrimmed([]string{
			Point12ValCPublicPrivateBoundaryStateActive,
			Point12ValCPublicPrivateBoundaryStateBlocked,
		}, model.BoundaryState) {
		reasons = append(reasons, "public_private_boundary_identity_or_metadata_invalid")
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValBReplayRequest.TenantScope) ||
		strings.TrimSpace(model.ExportID) != strings.TrimSpace(export.ExportID) ||
		strings.TrimSpace(model.OfflineBundleID) != strings.TrimSpace(offline.OfflineBundleID) {
		reasons = append(reasons, "public_private_boundary_binding_mismatch")
	}
	if len(model.ExportedFields) == 0 || !point12ValCAllExportedFieldsClassified(model) {
		reasons = append(reasons, "public_private_boundary_unclassified_exported_field")
	}
	if !point12ValCStringFieldListSubset(model.PublicFields, model.ExportedFields) ||
		!point12ValCStringFieldListSubset(model.PrivateFields, model.ExportedFields) ||
		!point12ValCStringFieldListSubset(model.RedactedFields, model.ExportedFields) ||
		!point12ValCStringFieldListSubset(model.SensitiveFields, model.ExportedFields) ||
		!point12ValCStringFieldListSubset(model.CustomerVisibleFields, model.ExportedFields) ||
		!point12ValCStringFieldListSubset(model.AuditorVisibleFields, model.ExportedFields) ||
		!point12ValCStringFieldListSubset(model.InternalOnlyFields, model.ExportedFields) {
		reasons = append(reasons, "public_private_boundary_field_subset_invalid")
	}
	if point12ValCFieldListsOverlap(model.CustomerVisibleFields, model.PrivateFields) ||
		point12ValCFieldListsOverlap(model.CustomerVisibleFields, model.InternalOnlyFields) ||
		point12ValCFieldListsOverlap(model.PublicFields, model.PrivateFields) ||
		point12ValCFieldListsOverlap(model.PublicFields, model.InternalOnlyFields) {
		reasons = append(reasons, "public_private_boundary_private_field_leak")
	}
	if point12ValCCustomerFacingAudience(export.ExportAudience) &&
		point12ValCStringMentionedInTexts(append([]string{}, model.PrivateFields...), export.CustomerVisibleSummary, strings.Join(export.Limitations, " "), redaction.RedactionSummary) {
		reasons = append(reasons, "public_private_boundary_text_leak")
	}
	if len(reasons) > 0 {
		return Point12ValCPublicPrivateBoundaryStateBlocked, reasons
	}
	return Point12ValCPublicPrivateBoundaryStateActive, nil
}

func EvaluatePoint12ValCPublicPrivateBoundaryState(model Point12ValCPublicPrivateBoundary, dependency Point12ValCDependencySnapshot, export Point12ValCAuditExportBundle, offline Point12ValCOfflineVerificationBundle, redaction Point12ValCRedactionManifest) string {
	state, _ := point12ValCPublicPrivateBoundaryStateAndReasons(model, dependency, export, offline, redaction)
	return state
}

func point12ValCAuditExportStateAndReasons(model Point12ValCAuditExportBundle, dependency Point12ValCDependencySnapshot, redactionManifestState string, redactionImpactState string, offlineState string, boundaryState string) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValCExportRefValid(model.ExportID) ||
		!point12ValCExportKindValid(model.ExportKind) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point12ValBReplayResultRefValid(model.ReplayResultID) ||
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
		!point12Val0CompatibilityProfileRefValid(model.CompatibilityProfileRef) ||
		!point12Val0RedactionManifestRefValid(model.RedactionManifestRef) ||
		(strings.TrimSpace(model.OfflineBundleRef) != "" && !point12ValCOfflineBundleRefValid(model.OfflineBundleRef)) ||
		!point12Val0HashValid(model.ManifestPayloadHash) ||
		!point12ValASignatureMetadataRefValid(model.SignatureMetadataRef) ||
		!point11Val0ValidTimestamp(model.GeneratedAt) ||
		!point12Val0VersionIdentityValid(model.FreshnessWindow) ||
		!point12ValCExportScopeValid(model.ExportScope) ||
		!point12ValCExportAudienceValid(model.ExportAudience) ||
		!point12ValCPublicPrivateClassificationValid(model.PublicPrivateClassification) ||
		!point12Val0RetentionClassRefValid(model.RetentionClassRef) ||
		!point12ValCRetentionOwnerRefValid(model.RetentionOwnerRef) ||
		!point12ValCDisposalPathRefValid(model.DisposalPathRef) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12Val0OptionalStringListValid(model.Limitations, point11Val0IdentityValueValid) ||
		!point11Val0ContainsTrimmed(point12ValCExportStates(), model.ExportState) ||
		!point12Val0OptionalClaimTextListValid(model.ExportOutputClaims) {
		blockedReasons = append(blockedReasons, "audit_export_identity_or_metadata_invalid")
	}
	if len(model.EvidenceRefs) != len(model.EvidenceHashRefs) {
		blockedReasons = append(blockedReasons, "audit_export_evidence_hash_alignment_invalid")
	}
	if !model.AdvisoryOnly {
		blockedReasons = append(blockedReasons, "audit_export_must_remain_advisory_only")
	}
	if strings.TrimSpace(model.NoOverclaimState) != Point12Val0NoOverclaimStateActive {
		blockedReasons = append(blockedReasons, "audit_export_no_overclaim_state_invalid")
	}
	if strings.TrimSpace(model.TenantBoundaryState) != strings.TrimSpace(boundaryState) {
		blockedReasons = append(blockedReasons, "audit_export_boundary_state_binding_mismatch")
	}
	if strings.TrimSpace(model.RedactionImpactState) != strings.TrimSpace(redactionImpactState) {
		blockedReasons = append(blockedReasons, "audit_export_redaction_impact_state_binding_mismatch")
	}
	if point12Val0ContainsPrematurePassToken(
		model.ExportID,
		model.ReplayResultID,
		model.CustomerVisibleSummary,
		strings.Join(model.ExportOutputClaims, " "),
		strings.Join(model.Limitations, " "),
	) {
		blockedReasons = append(blockedReasons, "audit_export_premature_point12_pass")
	}
	if point12Val0ContainsForbiddenClaim(strings.Join(model.ExportOutputClaims, " "), model.CustomerVisibleSummary) {
		blockedReasons = append(blockedReasons, "audit_export_overclaim_detected")
	}
	if point12ValCCustomerFacingAudience(model.ExportAudience) &&
		point12Val0ContainsForbiddenClaim(strings.Join(model.Limitations, " ")) {
		blockedReasons = append(blockedReasons, "audit_export_customer_facing_limitations_overclaim")
	}
	if strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(dependency.ValAManifest.ProofPackID) ||
		strings.TrimSpace(model.ManifestID) != strings.TrimSpace(dependency.ValAManifest.ManifestID) ||
		strings.TrimSpace(model.ReplayResultID) != strings.TrimSpace(dependency.ValBReplayResult.ReplayResultID) ||
		strings.TrimSpace(model.DecisionID) != strings.TrimSpace(dependency.ValBReplayRequest.DecisionID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValBReplayRequest.TenantScope) ||
		strings.TrimSpace(model.ArtifactRef) != strings.TrimSpace(dependency.ValBReplayRequest.ArtifactRef) ||
		strings.TrimSpace(model.ArtifactHash) != strings.TrimSpace(dependency.ValBReplayRequest.ArtifactHash) ||
		!point12Val0ExactStringSetMatch(model.EvidenceRefs, dependency.ValBReplayRequest.EvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.EvidenceHashRefs, dependency.ValBReplayRequest.EvidenceHashRefs) ||
		strings.TrimSpace(model.PolicyRef) != strings.TrimSpace(dependency.ValBReplayRequest.PolicyRef) ||
		strings.TrimSpace(model.PolicyVersion) != strings.TrimSpace(dependency.ValBReplayRequest.PolicyVersion) ||
		strings.TrimSpace(model.PolicyHash) != strings.TrimSpace(dependency.ValBReplayRequest.PolicyHash) ||
		strings.TrimSpace(model.EngineVersion) != strings.TrimSpace(dependency.ValBReplayRequest.EngineVersion) ||
		strings.TrimSpace(model.EngineHash) != strings.TrimSpace(dependency.ValBReplayRequest.EngineHash) ||
		strings.TrimSpace(model.SchemaVersion) != strings.TrimSpace(dependency.ValBReplayRequest.SchemaVersion) ||
		strings.TrimSpace(model.SchemaHash) != strings.TrimSpace(dependency.ValBReplayRequest.SchemaHash) ||
		!point12Val0ExactStringSetMatch(model.ClaimRefs, dependency.ValBReplayRequest.ClaimRefs) ||
		!point12Val0ExactStringSetMatch(model.GovernanceEventRefs, dependency.ValBReplayRequest.GovernanceEventRefs) ||
		strings.TrimSpace(model.CompatibilityProfileRef) != strings.TrimSpace(dependency.ValBReplayRequest.CompatibilityProfileRef) ||
		strings.TrimSpace(model.RedactionManifestRef) != strings.TrimSpace(dependency.ValBReplayRequest.RedactionManifestRef) ||
		strings.TrimSpace(model.ManifestPayloadHash) != strings.TrimSpace(dependency.ValBReplayRequest.ManifestPayloadHash) ||
		strings.TrimSpace(model.SignatureMetadataRef) != strings.TrimSpace(dependency.ValAManifest.SignatureMetadataRef) {
		blockedReasons = append(blockedReasons, "audit_export_dependency_binding_mismatch")
	}
	if !point12ValCExportRefValid(model.ExportID) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point12ValBReplayResultRefValid(model.ReplayResultID) ||
		!point12Val0DecisionRefValid(model.DecisionID) {
		return Point12ValCExportStateBlocked, blockedReasons
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValBReplayRequest.TenantScope) {
		return Point12ValCExportStateTenantMismatch, append(blockedReasons, "audit_export_tenant_scope_mismatch")
	}
	if boundaryState != Point12ValCPublicPrivateBoundaryStateActive {
		return Point12ValCExportStateBoundaryViolation, append(blockedReasons, "audit_export_public_private_boundary_violation")
	}
	if strings.TrimSpace(model.RetentionClassRef) == "" || strings.TrimSpace(model.RetentionOwnerRef) == "" || strings.TrimSpace(model.DisposalPathRef) == "" {
		return Point12ValCExportStateRetentionMissing, append(blockedReasons, "audit_export_retention_metadata_missing")
	}
	if strings.TrimSpace(dependency.ValBReplayTaxonomy) == Point12Val0ReplayResultTamperDetected ||
		strings.TrimSpace(dependency.ValBReplayTaxonomy) == Point12Val0ReplayResultBlockedReplay {
		return Point12ValCExportStateTampered, append(blockedReasons, "audit_export_tampered_or_blocked_replay_dependency")
	}
	if strings.TrimSpace(dependency.ValBReplayTaxonomy) == Point12Val0ReplayResultUnsupportedVersion {
		if strings.TrimSpace(model.ExportState) != Point12ValCExportStateUnsupported &&
			strings.TrimSpace(model.ExportState) != Point12ValCExportStateReviewRequired {
			reviewReasons = append(reviewReasons, "audit_export_unsupported_version_requires_limited_state")
		}
	}
	if point12ValCDependencyCheckResultReviewRequired(dependency.ValBManifestIntegrityResult) ||
		point12ValCDependencyCheckResultReviewRequired(dependency.ValBSignatureMetadataResult) ||
		point12ValCDependencyCheckResultReviewRequired(dependency.ValBCompatibilityResult) {
		if strings.TrimSpace(model.ExportState) == Point12ValCExportStateReady ||
			strings.TrimSpace(model.ExportState) == Point12ValCExportStateProjectionOnly {
			reviewReasons = append(reviewReasons, "audit_export_unsupported_dependency_check_requires_limited_state")
		}
	}
	if strings.TrimSpace(dependency.ValBReplayTaxonomy) == Point12Val0ReplayResultInsufficientEvidence {
		if strings.TrimSpace(model.ExportState) == Point12ValCExportStateReady || strings.TrimSpace(model.ExportState) == Point12ValCExportStateProjectionOnly {
			blockedReasons = append(blockedReasons, "audit_export_insufficient_evidence_cannot_be_export_ready")
		}
		if len(model.Limitations) == 0 {
			reviewReasons = append(reviewReasons, "audit_export_insufficient_evidence_limitations_missing")
		}
	}
	if strings.TrimSpace(dependency.ValBReplayTaxonomy) == Point12Val0ReplayResultRedactedLimitations ||
		strings.TrimSpace(redactionImpactState) == Point12ValCRedactionImpactRedactedLimits ||
		strings.TrimSpace(redactionImpactState) == Point12ValCRedactionImpactBlockedReplay ||
		strings.TrimSpace(redactionImpactState) == Point12ValCRedactionImpactInsufficient ||
		strings.TrimSpace(redactionImpactState) == Point12ValCRedactionImpactPartialAdvisory ||
		strings.TrimSpace(redactionManifestState) != Point12ValCRedactionManifestStateActive {
		if strings.TrimSpace(model.ExportState) == Point12ValCExportStateReady {
			blockedReasons = append(blockedReasons, "audit_export_redaction_limitations_cannot_be_export_ready")
		}
		if len(model.Limitations) == 0 {
			reviewReasons = append(reviewReasons, "audit_export_redaction_limitations_missing")
		}
	}
	if strings.TrimSpace(model.ExportKind) == point12ValCExportKindVerifierPackageMetadata && strings.TrimSpace(model.OfflineBundleRef) == "" {
		blockedReasons = append(blockedReasons, "audit_export_offline_bundle_ref_missing")
	}
	if strings.TrimSpace(model.ExportKind) == point12ValCExportKindVerifierPackageMetadata &&
		strings.TrimSpace(offlineState) != Point12ValCOfflineBundleStateActive &&
		strings.TrimSpace(offlineState) != Point12ValCOfflineBundleStatePartialAdvisoryOnly &&
		strings.TrimSpace(offlineState) != Point12ValCOfflineBundleStateRedactedLimitations &&
		strings.TrimSpace(offlineState) != Point12ValCOfflineBundleStateReviewRequired &&
		strings.TrimSpace(offlineState) != Point12ValCOfflineBundleStateUnsupported {
		blockedReasons = append(blockedReasons, "audit_export_offline_bundle_state_invalid")
	}
	if len(blockedReasons) > 0 {
		return Point12ValCExportStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValCExportStateReviewRequired, reviewReasons
	}
	return strings.TrimSpace(model.ExportState), nil
}

func EvaluatePoint12ValCAuditExportState(model Point12ValCAuditExportBundle, dependency Point12ValCDependencySnapshot, redactionManifestState string, redactionImpactState string, offlineState string, boundaryState string) string {
	state, _ := point12ValCAuditExportStateAndReasons(model, dependency, redactionManifestState, redactionImpactState, offlineState, boundaryState)
	return state
}

func EvaluatePoint12ValCState(model Point12ValCFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		strings.TrimSpace(model.DependencyState) == Point12ValCDependencyStateBlocked ||
		strings.TrimSpace(model.RedactionManifestState) == Point12ValCRedactionManifestStateBlocked ||
		strings.TrimSpace(model.PublicPrivateBoundaryState) == Point12ValCPublicPrivateBoundaryStateBlocked ||
		strings.TrimSpace(model.OfflineBundleState) == Point12ValCOfflineBundleStateBlocked ||
		strings.TrimSpace(model.ExportState) == Point12ValCExportStateBlocked ||
		strings.TrimSpace(model.ExportState) == Point12ValCExportStateTampered ||
		strings.TrimSpace(model.ExportState) == Point12ValCExportStateTenantMismatch ||
		strings.TrimSpace(model.ExportState) == Point12ValCExportStateBoundaryViolation ||
		strings.TrimSpace(model.ExportState) == Point12ValCExportStateRetentionMissing {
		return Point12ValCStateBlocked
	}
	if strings.TrimSpace(model.DependencyState) == Point12ValCDependencyStateReviewRequired ||
		strings.TrimSpace(model.RedactionManifestState) == Point12ValCRedactionManifestStateReviewRequired ||
		strings.TrimSpace(model.OfflineBundleState) == Point12ValCOfflineBundleStateReviewRequired ||
		strings.TrimSpace(model.OfflineBundleState) == Point12ValCOfflineBundleStateUnsupported ||
		strings.TrimSpace(model.ExportState) == Point12ValCExportStatePartialAdvisory ||
		strings.TrimSpace(model.ExportState) == Point12ValCExportStateRedactedLimitations ||
		strings.TrimSpace(model.ExportState) == Point12ValCExportStateInsufficient ||
		strings.TrimSpace(model.ExportState) == Point12ValCExportStateUnsupported ||
		strings.TrimSpace(model.ExportState) == Point12ValCExportStateReviewRequired {
		return Point12ValCStateReviewRequired
	}
	return Point12ValCStateActive
}

func point12ValCBlockingReasons(model Point12ValCFoundation) []string {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "aggregate_projection_disclaimer_blocked")
	}
	if model.DependencyState == Point12ValCDependencyStateBlocked {
		reasons = append(reasons, "point12_valb_dependency_blocked")
	}
	if model.RedactionManifestState == Point12ValCRedactionManifestStateBlocked {
		reasons = append(reasons, "redaction_manifest_blocked")
	}
	if model.PublicPrivateBoundaryState == Point12ValCPublicPrivateBoundaryStateBlocked {
		reasons = append(reasons, "public_private_boundary_blocked")
	}
	if model.OfflineBundleState == Point12ValCOfflineBundleStateBlocked {
		reasons = append(reasons, "offline_bundle_blocked")
	}
	switch model.ExportState {
	case Point12ValCExportStateBlocked:
		reasons = append(reasons, "audit_export_blocked")
	case Point12ValCExportStateTampered:
		reasons = append(reasons, "audit_export_tamper_detected")
	case Point12ValCExportStateTenantMismatch:
		reasons = append(reasons, "audit_export_tenant_scope_mismatch")
	case Point12ValCExportStateBoundaryViolation:
		reasons = append(reasons, "audit_export_public_private_boundary_violation")
	case Point12ValCExportStateRetentionMissing:
		reasons = append(reasons, "audit_export_retention_missing")
	}
	return reasons
}

func Point12ValCFoundationModel() Point12ValCFoundation {
	dependency := point12ValCDependencySnapshotModel()
	export := Point12ValCAuditExportBundle{
		ExportID:                    "export_point12_valc_001",
		ExportKind:                  point12ValCExportKindAuditReadyJSON,
		ProofPackID:                 dependency.ValAManifest.ProofPackID,
		ManifestID:                  dependency.ValAManifest.ManifestID,
		ReplayResultID:              dependency.ValBReplayResult.ReplayResultID,
		DecisionID:                  dependency.ValBReplayRequest.DecisionID,
		TenantScope:                 dependency.ValBReplayRequest.TenantScope,
		ArtifactRef:                 dependency.ValBReplayRequest.ArtifactRef,
		ArtifactHash:                dependency.ValBReplayRequest.ArtifactHash,
		EvidenceRefs:                append([]string{}, dependency.ValBReplayRequest.EvidenceRefs...),
		EvidenceHashRefs:            append([]string{}, dependency.ValBReplayRequest.EvidenceHashRefs...),
		PolicyRef:                   dependency.ValBReplayRequest.PolicyRef,
		PolicyVersion:               dependency.ValBReplayRequest.PolicyVersion,
		PolicyHash:                  dependency.ValBReplayRequest.PolicyHash,
		EngineVersion:               dependency.ValBReplayRequest.EngineVersion,
		EngineHash:                  dependency.ValBReplayRequest.EngineHash,
		SchemaVersion:               dependency.ValBReplayRequest.SchemaVersion,
		SchemaHash:                  dependency.ValBReplayRequest.SchemaHash,
		ClaimRefs:                   append([]string{}, dependency.ValBReplayRequest.ClaimRefs...),
		GovernanceEventRefs:         append([]string{}, dependency.ValBReplayRequest.GovernanceEventRefs...),
		CompatibilityProfileRef:     dependency.ValBReplayRequest.CompatibilityProfileRef,
		RedactionManifestRef:        dependency.ValBReplayRequest.RedactionManifestRef,
		OfflineBundleRef:            "offline_bundle_point12_valc_001",
		ManifestPayloadHash:         dependency.ValBReplayRequest.ManifestPayloadHash,
		SignatureMetadataRef:        dependency.ValAManifest.SignatureMetadataRef,
		GeneratedAt:                 "2026-05-03T15:00:00Z",
		FreshnessWindow:             "freshness_window_point12_valc_24h",
		ExportScope:                 point12ValCExportScopeTenantScoped,
		ExportAudience:              point12ValCExportAudienceInternalAudit,
		PublicPrivateClassification: point12ValCClassificationAuditorRestricted,
		RetentionClassRef:           dependency.ValAManifest.RetentionClassRef,
		RetentionOwnerRef:           "retention_owner_point12_valc_001",
		DisposalPathRef:             "disposal_path_point12_valc_001",
		ProjectionDisclaimer:        point12ValCProjectionDisclaimerBaseline,
		Limitations:                 []string{"advisory projection only"},
		AdvisoryOnly:                true,
		NoOverclaimState:            Point12Val0NoOverclaimStateActive,
		TenantBoundaryState:         Point12ValCPublicPrivateBoundaryStateActive,
		RedactionImpactState:        Point12ValCRedactionImpactNoDecisionImpact,
		ExportState:                 Point12ValCExportStateReady,
		ExportOutputClaims:          []string{"bounded claim"},
		CustomerVisibleSummary:      "advisory projection only",
	}
	redaction := Point12ValCRedactionManifest{
		RedactionManifestID:            dependency.ValBReplayRequest.RedactionManifestRef,
		ProofPackID:                    dependency.ValAManifest.ProofPackID,
		ExportID:                       export.ExportID,
		TenantScope:                    dependency.ValBReplayRequest.TenantScope,
		RedactionPolicyRef:             dependency.ValBReplayRequest.PolicyRef,
		RedactionPolicyVersion:         dependency.ValBReplayRequest.PolicyVersion,
		PostRedactionResult:            Point12Val0ReplayResultSameDecision,
		MinimumSafeClaimAfterRedaction: "bounded claim",
		RedactionSummary:               "internal summary: no decisive redaction applied",
		PartialOrAdvisoryOnly:          false,
		GeneratedAt:                    "2026-05-03T15:01:00Z",
		RetentionClassRef:              dependency.ValAManifest.RetentionClassRef,
		RetentionOwnerRef:              export.RetentionOwnerRef,
		DisposalPathRef:                export.DisposalPathRef,
	}
	impact := Point12ValCRedactionImpactVerdict{
		RedactionImpactID:              "redaction_impact_point12_valc_001",
		RedactionManifestID:            redaction.RedactionManifestID,
		AffectsDecision:                false,
		AffectsReplay:                  false,
		PostRedactionResult:            redaction.PostRedactionResult,
		MinimumSafeClaimAfterRedaction: redaction.MinimumSafeClaimAfterRedaction,
		AllowedExportState:             Point12ValCExportStateReady,
		DisallowedExportState:          Point12ValCExportStateBlocked,
		BlocksFullExport:               false,
		RequiresPartialAdvisoryExport:  false,
		RequiresCustomerReview:         false,
		RedactionImpactState:           Point12ValCRedactionImpactNoDecisionImpact,
	}
	offline := Point12ValCOfflineVerificationBundle{
		OfflineBundleID:             export.OfflineBundleRef,
		ProofPackID:                 dependency.ValAManifest.ProofPackID,
		ManifestID:                  dependency.ValAManifest.ManifestID,
		ReplayRequestID:             dependency.ValBReplayRequest.ReplayRequestID,
		ReplayResultID:              dependency.ValBReplayResult.ReplayResultID,
		TenantScope:                 dependency.ValBReplayRequest.TenantScope,
		ArtifactRef:                 dependency.ValBReplayRequest.ArtifactRef,
		ArtifactHash:                dependency.ValBReplayRequest.ArtifactHash,
		EvidenceRefs:                append([]string{}, dependency.ValBReplayRequest.EvidenceRefs...),
		EvidenceHashRefs:            append([]string{}, dependency.ValBReplayRequest.EvidenceHashRefs...),
		PolicyRef:                   dependency.ValBReplayRequest.PolicyRef,
		PolicyVersion:               dependency.ValBReplayRequest.PolicyVersion,
		PolicyHash:                  dependency.ValBReplayRequest.PolicyHash,
		EngineVersion:               dependency.ValBReplayRequest.EngineVersion,
		EngineHash:                  dependency.ValBReplayRequest.EngineHash,
		SchemaVersion:               dependency.ValBReplayRequest.SchemaVersion,
		SchemaHash:                  dependency.ValBReplayRequest.SchemaHash,
		ManifestPayloadHash:         dependency.ValBReplayRequest.ManifestPayloadHash,
		SignatureMetadataRef:        dependency.ValAManifest.SignatureMetadataRef,
		DetachedSignatureRef:        dependency.ValAManifest.DetachedSignatureRef,
		CompatibilityProfileRef:     dependency.ValBReplayRequest.CompatibilityProfileRef,
		RedactionManifestRef:        dependency.ValBReplayRequest.RedactionManifestRef,
		VerificationPolicyRef:       "verification_policy_point12_valc_001",
		GeneratedAt:                 "2026-05-03T15:02:00Z",
		BundleFormatVersion:         "bundle_format_version_point12_valc_v1",
		SupportedVerifierVersions:   []string{"verifier_version_point12_valc_001"},
		RequestedVerifierVersion:    "verifier_version_point12_valc_001",
		NoExternalAPIRequired:       true,
		ExternalAPIUsed:             false,
		ContainsPrivateData:         true,
		PublicPrivateClassification: point12ValCClassificationAuditorRestricted,
		RetentionClassRef:           dependency.ValAManifest.RetentionClassRef,
		RetentionOwnerRef:           export.RetentionOwnerRef,
		DisposalPathRef:             export.DisposalPathRef,
		Limitations:                 []string{"offline advisory verification only"},
		OfflineOutputClaims:         []string{"bounded claim"},
		CustomerVisibleExplanation:  "advisory projection only",
		OfflineState:                Point12ValCOfflineBundleStateActive,
	}
	boundary := Point12ValCPublicPrivateBoundary{
		BoundaryID:           "public_private_boundary_point12_valc_001",
		TenantScope:          dependency.ValBReplayRequest.TenantScope,
		ExportID:             export.ExportID,
		OfflineBundleID:      offline.OfflineBundleID,
		ExportedFields:       []string{"artifact_hash", "evidence_hash_refs", "policy_hash", "engine_hash", "schema_hash", "limitations"},
		PrivateFields:        []string{"artifact_hash", "evidence_hash_refs", "policy_hash", "engine_hash", "schema_hash", "limitations"},
		SensitiveFields:      []string{"artifact_hash", "evidence_hash_refs"},
		AuditorVisibleFields: []string{"artifact_hash", "evidence_hash_refs", "policy_hash", "engine_hash", "schema_hash", "limitations"},
		Classification:       point12ValCClassificationAuditorRestricted,
		DataResidencyRef:     "data_residency_eu_001",
		AllowedAudience:      point12ValCExportAudienceInternalAudit,
		BoundaryState:        Point12ValCPublicPrivateBoundaryStateActive,
	}
	return Point12ValCFoundation{
		CurrentState:               Point12ValCStateActive,
		ProjectionDisclaimer:       point12ValCProjectionDisclaimerBaseline,
		DependencyState:            Point12ValCDependencyStateActive,
		RedactionManifestState:     Point12ValCRedactionManifestStateActive,
		RedactionImpactState:       Point12ValCRedactionImpactNoDecisionImpact,
		OfflineBundleState:         Point12ValCOfflineBundleStateActive,
		PublicPrivateBoundaryState: Point12ValCPublicPrivateBoundaryStateActive,
		ExportState:                Point12ValCExportStateReady,
		Dependency:                 dependency,
		ExportBundle:               export,
		RedactionManifest:          redaction,
		RedactionImpactVerdict:     impact,
		OfflineBundle:              offline,
		PublicPrivateBoundary:      boundary,
	}
}

func ComputePoint12ValCFoundation(model Point12ValCFoundation) Point12ValCFoundation {
	model.DependencyState = EvaluatePoint12ValCDependencyState(model.Dependency)
	redactionManifestState, redactionReasons := point12ValCRedactionManifestStateAndReasons(model.RedactionManifest, model.Dependency, model.ExportBundle)
	model.RedactionManifestState = redactionManifestState
	redactionImpactState, redactionImpactReasons := point12ValCRedactionImpactStateAndReasons(model.RedactionImpactVerdict, model.RedactionManifest, model.Dependency)
	model.RedactionImpactState = redactionImpactState
	model.RedactionImpactVerdict.RedactionImpactState = redactionImpactState
	offlineState, offlineReasons := point12ValCOfflineBundleStateAndReasons(model.OfflineBundle, model.Dependency, model.RedactionImpactState)
	model.OfflineBundleState = offlineState
	model.OfflineBundle.OfflineState = offlineState
	boundaryState, boundaryReasons := point12ValCPublicPrivateBoundaryStateAndReasons(model.PublicPrivateBoundary, model.Dependency, model.ExportBundle, model.OfflineBundle, model.RedactionManifest)
	model.PublicPrivateBoundaryState = boundaryState
	model.PublicPrivateBoundary.BoundaryState = boundaryState
	exportState, exportReasons := point12ValCAuditExportStateAndReasons(model.ExportBundle, model.Dependency, model.RedactionManifestState, model.RedactionImpactState, model.OfflineBundleState, model.PublicPrivateBoundaryState)
	model.ExportState = exportState
	model.ExportBundle.ExportState = exportState
	model.CurrentState = EvaluatePoint12ValCState(model)
	model.BlockingReasons = point12ValCBlockingReasons(model)
	model.ReviewPrerequisites = append([]string{}, model.Dependency.ReviewPrerequisites...)
	if model.RedactionManifestState == Point12ValCRedactionManifestStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, redactionReasons...)
	}
	if model.RedactionImpactState == Point12ValCRedactionImpactReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, redactionImpactReasons...)
	}
	if model.OfflineBundleState == Point12ValCOfflineBundleStateReviewRequired || model.OfflineBundleState == Point12ValCOfflineBundleStateUnsupported {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, offlineReasons...)
	}
	if model.ExportState == Point12ValCExportStateReviewRequired ||
		model.ExportState == Point12ValCExportStateUnsupported ||
		model.ExportState == Point12ValCExportStateInsufficient ||
		model.ExportState == Point12ValCExportStateRedactedLimitations ||
		model.ExportState == Point12ValCExportStatePartialAdvisory {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, exportReasons...)
	}
	if model.PublicPrivateBoundaryState == Point12ValCPublicPrivateBoundaryStateBlocked {
		model.BlockingReasons = append(model.BlockingReasons, boundaryReasons...)
	}
	return model
}

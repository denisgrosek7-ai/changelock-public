package formal

import (
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

const (
	Point12Val0StateActive         = "point12_val0_replay_discipline_foundation_active"
	Point12Val0StateBlocked        = "point12_val0_replay_discipline_foundation_blocked"
	Point12Val0StateReviewRequired = "point12_val0_replay_discipline_foundation_review_required"

	Point12Val0DependencyStateActive         = "point12_val0_dependency_active"
	Point12Val0DependencyStateBlocked        = "point12_val0_dependency_blocked"
	Point12Val0DependencyStateReviewRequired = "point12_val0_dependency_review_required"

	Point12Val0DeterminismContractStateActive  = "point12_val0_determinism_contract_active"
	Point12Val0DeterminismContractStateBlocked = "point12_val0_determinism_contract_blocked"

	Point12Val0CompatibilityProfileStateActive  = "point12_val0_compatibility_profile_active"
	Point12Val0CompatibilityProfileStateBlocked = "point12_val0_compatibility_profile_blocked"

	Point12Val0ManifestStateActive  = "point12_val0_manifest_active"
	Point12Val0ManifestStateBlocked = "point12_val0_manifest_blocked"

	Point12Val0ReplayAssessmentStateActive  = "point12_val0_replay_assessment_active"
	Point12Val0ReplayAssessmentStateBlocked = "point12_val0_replay_assessment_blocked"

	Point12Val0RedactionBoundaryStateActive  = "point12_val0_redaction_boundary_active"
	Point12Val0RedactionBoundaryStateBlocked = "point12_val0_redaction_boundary_blocked"

	Point12Val0FinancialEvidenceSupportStateActive  = "point12_val0_financial_evidence_support_active"
	Point12Val0FinancialEvidenceSupportStateBlocked = "point12_val0_financial_evidence_support_blocked"

	Point12Val0ProvenanceStateActive         = "point12_val0_provenance_active"
	Point12Val0ProvenanceStateBlocked        = "point12_val0_provenance_blocked"
	Point12Val0ProvenanceStateReviewRequired = "point12_val0_provenance_review_required"

	Point12Val0NoOverclaimStateActive  = "point12_val0_no_overclaim_active"
	Point12Val0NoOverclaimStateBlocked = "point12_val0_no_overclaim_blocked"
)

const (
	Point12Val0ProofPackStateDraft                   = "draft"
	Point12Val0ProofPackStateGenerated               = "generated"
	Point12Val0ProofPackStateSignedMetadataValidated = "signed_metadata_validated"
	Point12Val0ProofPackStateReplayable              = "replayable"
	Point12Val0ProofPackStatePartiallyReplayable     = "partially_replayable"
	Point12Val0ProofPackStateRedactedLimited         = "redacted_limited"
	Point12Val0ProofPackStateTampered                = "tampered"
	Point12Val0ProofPackStateUnsupported             = "unsupported"
	Point12Val0ProofPackStateExpired                 = "expired"
	Point12Val0ProofPackStateRevoked                 = "revoked"
	Point12Val0ProofPackStateSuperseded              = "superseded"
	Point12Val0ProofPackStateBlocked                 = "blocked"
)

const (
	Point12Val0ReplayResultSameDecision         = "same_decision"
	Point12Val0ReplayResultDifferentDecision    = "different_decision"
	Point12Val0ReplayResultBlockedReplay        = "blocked_replay"
	Point12Val0ReplayResultInsufficientEvidence = "insufficient_evidence"
	Point12Val0ReplayResultUnsupportedVersion   = "unsupported_version"
	Point12Val0ReplayResultTamperDetected       = "tamper_detected"
	Point12Val0ReplayResultRedactedLimitations  = "redacted_limitations"
	Point12Val0ReplayResultPolicyMismatch       = "policy_mismatch"
	Point12Val0ReplayResultEngineMismatch       = "engine_mismatch"
	Point12Val0ReplayResultSchemaMismatch       = "schema_mismatch"
	Point12Val0ReplayResultEvidenceMismatch     = "evidence_mismatch"
	Point12Val0ReplayResultClaimMismatch        = "claim_mismatch"
	Point12Val0ReplayResultGovernanceMismatch   = "governance_mismatch"
)

const (
	point12Val0ReplayModeOriginalContext      = "original_context"
	point12Val0ReplayModeCurrentPolicyContext = "current_policy_context"
	point12Val0ReplayModeComparisonMode       = "comparison_mode"

	point12Val0CompatibilityExactRequired     = "exact_required"
	point12Val0CompatibilityCompatibleAllowed = "compatible_allowed"
	point12Val0CompatibilityBlocked           = "blocked"

	point12Val0EvidenceCompatibilityExactHashRequired = "exact_hash_required"

	point12Val0ToolchainCompatibilityRequiredIfDecisive = "required_if_decisive"
	point12Val0ToolchainCompatibilityNotRequired        = "not_required"

	point12Val0UnsupportedBehaviorBlockedReplay      = "blocked_replay"
	point12Val0UnsupportedBehaviorUnsupportedVersion = "unsupported_version"

	point12Val0ProfileTypeFinancialReview = "financial_review"
	point12Val0ProfileTypeInsuranceReview = "insurance_review"
	point12Val0ProfileTypeAuditReview     = "audit_review"

	point12Val0PointID                      = "point_12"
	point12Val0WaveID                       = "val_0"
	point12Val0ProjectionDisclaimerBaseline = "projection_only not_canonical_truth point12_val0_replay_discipline_foundation"
)

type Point12Val0Point11ReviewContext struct {
	SnapshotFromComputedOutput bool     `json:"snapshot_from_computed_output"`
	ReviewPrerequisites        []string `json:"review_prerequisites,omitempty"`
}

type Point12Val0DependencySnapshot struct {
	UpstreamCurrentState                        string   `json:"upstream_current_state"`
	UpstreamDependencyState                     string   `json:"upstream_dependency_state"`
	UpstreamPointID                             string   `json:"upstream_point_id"`
	UpstreamWaveID                              string   `json:"upstream_wave_id"`
	UpstreamPassClosureManifestState            string   `json:"upstream_pass_closure_manifest_state"`
	UpstreamFinalPassGateState                  string   `json:"upstream_final_pass_gate_state"`
	UpstreamPoint11PassAllowed                  bool     `json:"upstream_point11_pass_allowed"`
	UpstreamPoint11PassToken                    string   `json:"upstream_point11_pass_token"`
	UpstreamClosureManifestRef                  string   `json:"upstream_closure_manifest_ref"`
	PolicyAuthorityContextRefs                  []string `json:"policy_authority_context_refs,omitempty"`
	ClaimAuthorityContextRefs                   []string `json:"claim_authority_context_refs,omitempty"`
	GovernanceAuthorityContextRefs              []string `json:"governance_authority_context_refs,omitempty"`
	PolicyAuthorityContextRef                   string   `json:"policy_authority_context_ref"`
	ClaimAuthorityContextRef                    string   `json:"claim_authority_context_ref"`
	GovernanceAuthorityContextRef               string   `json:"governance_authority_context_ref"`
	ProjectionDisclaimer                        string   `json:"projection_disclaimer"`
	UpstreamPoint11PassObservedOutsideFinalPath bool     `json:"upstream_point11_pass_observed_outside_final_path"`
	SnapshotFromComputedOutput                  bool     `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                         []string `json:"review_prerequisites,omitempty"`
}

type Point12Val0ReplayDeterminismContract struct {
	ReplayMode                     string `json:"replay_mode"`
	StableEvidenceIdentityRequired bool   `json:"stable_evidence_identity_required"`
	StablePolicyIdentityRequired   bool   `json:"stable_policy_identity_required"`
	StableEngineIdentityRequired   bool   `json:"stable_engine_identity_required"`
	StableSchemaIdentityRequired   bool   `json:"stable_schema_identity_required"`
	StableTenantScopeRequired      bool   `json:"stable_tenant_scope_required"`
	StableArtifactIdentityRequired bool   `json:"stable_artifact_identity_required"`
	DriftExplanationRequired       bool   `json:"drift_explanation_required"`
	UnsupportedBehavior            string `json:"unsupported_behavior"`
}

type Point12Val0ProofPackCompatibilityProfile struct {
	CompatibilityProfileRef          string   `json:"compatibility_profile_ref"`
	ReplayMode                       string   `json:"replay_mode"`
	PolicyCompatibility              string   `json:"policy_compatibility"`
	EngineCompatibility              string   `json:"engine_compatibility"`
	SchemaCompatibility              string   `json:"schema_compatibility"`
	EvidenceCompatibility            string   `json:"evidence_compatibility"`
	ToolchainCompatibility           string   `json:"toolchain_compatibility"`
	UnsupportedBehavior              string   `json:"unsupported_behavior"`
	DecisionDriftExplanationRequired bool     `json:"decision_drift_explanation_required"`
	CompatibilityEvidenceRefs        []string `json:"compatibility_evidence_refs,omitempty"`
}

type Point12Val0SignedProofPackManifest struct {
	ProofPackID                   string   `json:"proof_pack_id"`
	DecisionID                    string   `json:"decision_id"`
	PointID                       string   `json:"point_id"`
	WaveID                        string   `json:"wave_id"`
	ProofPackState                string   `json:"proof_pack_state"`
	TenantScope                   string   `json:"tenant_scope"`
	ArtifactRef                   string   `json:"artifact_ref"`
	ArtifactHash                  string   `json:"artifact_hash"`
	EvidenceRefs                  []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs              []string `json:"evidence_hash_refs,omitempty"`
	PolicyRef                     string   `json:"policy_ref"`
	PolicyVersion                 string   `json:"policy_version"`
	PolicyHash                    string   `json:"policy_hash"`
	EngineVersion                 string   `json:"engine_version"`
	EngineHash                    string   `json:"engine_hash"`
	SchemaVersion                 string   `json:"schema_version"`
	SchemaHash                    string   `json:"schema_hash"`
	ClaimRefs                     []string `json:"claim_refs,omitempty"`
	GovernanceEventRefs           []string `json:"governance_event_refs,omitempty"`
	UpstreamClosureManifestRef    string   `json:"upstream_closure_manifest_ref"`
	DependencySnapshotRef         string   `json:"dependency_snapshot_ref"`
	PolicyAuthorityContextRef     string   `json:"policy_authority_context_ref"`
	ClaimAuthorityContextRef      string   `json:"claim_authority_context_ref"`
	GovernanceAuthorityContextRef string   `json:"governance_authority_context_ref"`
	CompatibilityProfileRef       string   `json:"compatibility_profile_ref"`
	GeneratedAt                   string   `json:"generated_at"`
	FreshnessWindow               string   `json:"freshness_window"`
	SigningKeyRef                 string   `json:"signing_key_ref"`
	SignatureRef                  string   `json:"signature_ref"`
	RedactionManifestRef          string   `json:"redaction_manifest_ref"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	RetentionClassRef             string   `json:"retention_class_ref"`
	ToolchainProvenanceRefs       []string `json:"toolchain_provenance_refs,omitempty"`
	AgentLineageRefs              []string `json:"agent_lineage_refs,omitempty"`
}

type Point12Val0ReplayAssessment struct {
	ReplayAssessmentID      string   `json:"replay_assessment_id"`
	ProofPackState          string   `json:"proof_pack_state"`
	ReplayResult            string   `json:"replay_result"`
	DriftExplanation        string   `json:"drift_explanation"`
	DeterminismContractRef  string   `json:"determinism_contract_ref"`
	CompatibilityProfileRef string   `json:"compatibility_profile_ref"`
	OriginalPolicyRef       string   `json:"original_policy_ref"`
	ReplayPolicyRef         string   `json:"replay_policy_ref"`
	OriginalPolicyHash      string   `json:"original_policy_hash"`
	ReplayPolicyHash        string   `json:"replay_policy_hash"`
	OriginalEngineHash      string   `json:"original_engine_hash"`
	ReplayEngineHash        string   `json:"replay_engine_hash"`
	OriginalSchemaVersion   string   `json:"original_schema_version"`
	ReplaySchemaVersion     string   `json:"replay_schema_version"`
	EvidenceRefs            []string `json:"evidence_refs,omitempty"`
	ReplayEvidenceRefs      []string `json:"replay_evidence_refs,omitempty"`
	EvidenceHashRefs        []string `json:"evidence_hash_refs,omitempty"`
	ReplayEvidenceHashRefs  []string `json:"replay_evidence_hash_refs,omitempty"`
	ClaimRefs               []string `json:"claim_refs,omitempty"`
	ReplayClaimRefs         []string `json:"replay_claim_refs,omitempty"`
	GovernanceEventRefs     []string `json:"governance_event_refs,omitempty"`
	ReplayGovernanceRefs    []string `json:"replay_governance_refs,omitempty"`
	DecisiveEvidencePresent bool     `json:"decisive_evidence_present"`
	ProjectionDisclaimer    string   `json:"projection_disclaimer"`
}

type Point12Val0RedactionBoundary struct {
	RedactionManifestRef                string   `json:"redaction_manifest_ref"`
	RedactedFields                      []string `json:"redacted_fields,omitempty"`
	RedactionReasons                    []string `json:"redaction_reasons,omitempty"`
	RedactionApproverRef                string   `json:"redaction_approver_ref"`
	RedactionApprovalEventRef           string   `json:"redaction_approval_event_ref"`
	RedactionAffectsDecision            bool     `json:"redaction_affects_decision"`
	RedactionAffectsReplay              bool     `json:"redaction_affects_replay"`
	PostRedactionResult                 string   `json:"post_redaction_result"`
	MinimumSafeClaimAfterRedaction      string   `json:"minimum_safe_claim_after_redaction"`
	DisallowedClaimsAfterRedaction      []string `json:"disallowed_claims_after_redaction,omitempty"`
	SurvivingClaimsAfterRedaction       []string `json:"surviving_claims_after_redaction,omitempty"`
	CustomerVisibleClaimsAfterRedaction []string `json:"customer_visible_claims_after_redaction,omitempty"`
	ExportedClaimsAfterRedaction        []string `json:"exported_claims_after_redaction,omitempty"`
	ReplayResultClaims                  []string `json:"replay_result_claims,omitempty"`
	RedactionSummary                    string   `json:"redaction_summary"`
	PartialOrAdvisoryOnly               bool     `json:"partial_or_advisory_only"`
}

type Point12Val0FinancialInsuranceEvidenceSupportProfile struct {
	ProfileType                       string   `json:"profile_type"`
	EvidenceSupportCategories         []string `json:"evidence_support_categories,omitempty"`
	RiskContextMetadata               []string `json:"risk_context_metadata,omitempty"`
	Limitations                       []string `json:"limitations,omitempty"`
	RequiredCustomerReview            bool     `json:"required_customer_review"`
	LegalReviewRequiredForExternalUse bool     `json:"legal_review_required_for_external_use"`
	NoPremiumGuarantee                bool     `json:"no_premium_guarantee"`
	NoRatingClaim                     bool     `json:"no_rating_claim"`
	NoComplianceGuarantee             bool     `json:"no_compliance_guarantee"`
	NoFinancialGuarantee              bool     `json:"no_financial_guarantee"`
	AllowedWordingRefs                []string `json:"allowed_wording_refs,omitempty"`
	BlockedWordingRefs                []string `json:"blocked_wording_refs,omitempty"`
	SupportStatement                  string   `json:"support_statement"`
}

type Point12Val0AgentLineageRecord struct {
	AgentID                string   `json:"agent_id"`
	AgentType              string   `json:"agent_type"`
	ModelOrRuleVersionRef  string   `json:"model_or_rule_version_ref"`
	PermissionManifestHash string   `json:"permission_manifest_hash"`
	InputEvidenceRefs      []string `json:"input_evidence_refs,omitempty"`
	AuditID                string   `json:"audit_id"`
	RecommendationID       string   `json:"recommendation_id"`
	HumanFeedbackRefs      []string `json:"human_feedback_refs,omitempty"`
	LineageInputOnly       bool     `json:"lineage_input_only"`
	ClaimsCertification    bool     `json:"claims_certification"`
	ClaimsSourceOfTruth    bool     `json:"claims_source_of_truth"`
	EmitsPrematurePass     bool     `json:"emits_premature_pass"`
}

type point12Val0AIEvidenceCandidateBindingField struct {
	FieldName            string
	BindingClass         string
	Reason               string
	ValidationRequired   bool
	MutationTestRequired bool
}

type Point12Val0ToolchainAgentProvenanceProfile struct {
	DecisiveToolchainProvenanceRequired bool                            `json:"decisive_toolchain_provenance_required"`
	CIJobIDRef                          string                          `json:"ci_job_id_ref"`
	RunnerImageHash                     string                          `json:"runner_image_hash"`
	BuildToolVersionRef                 string                          `json:"build_tool_version_ref"`
	ScannerVersionRef                   string                          `json:"scanner_version_ref"`
	SBOMGeneratorVersionRef             string                          `json:"sbom_generator_version_ref"`
	SigningToolVersionRef               string                          `json:"signing_tool_version_ref"`
	PolicyEngineBuildHash               string                          `json:"policy_engine_build_hash"`
	ExecutionEnvironmentClassRef        string                          `json:"execution_environment_class_ref"`
	AgentLineages                       []Point12Val0AgentLineageRecord `json:"agent_lineages,omitempty"`
	IntroducesNetworkCallPath           bool                            `json:"introduces_network_call_path"`
}

type Point12Val0NoOverclaimReview struct {
	ObservedCustomerFacingTexts []string `json:"observed_customer_facing_texts,omitempty"`
	ObservedExportFacingTexts   []string `json:"observed_export_facing_texts,omitempty"`
	ObservedDiagnostics         []string `json:"observed_diagnostics,omitempty"`
	AllowedSafeWording          []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording              []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
}

type Point12Val0Foundation struct {
	CurrentState                    string                                              `json:"current_state"`
	BlockingReasons                 []string                                            `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites             []string                                            `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer            string                                              `json:"projection_disclaimer"`
	DependencyState                 string                                              `json:"dependency_state"`
	DeterminismContractState        string                                              `json:"determinism_contract_state"`
	CompatibilityProfileState       string                                              `json:"compatibility_profile_state"`
	ManifestState                   string                                              `json:"manifest_state"`
	ReplayAssessmentState           string                                              `json:"replay_assessment_state"`
	RedactionBoundaryState          string                                              `json:"redaction_boundary_state"`
	FinancialEvidenceSupportState   string                                              `json:"financial_evidence_support_state"`
	ProvenanceState                 string                                              `json:"provenance_state"`
	NoOverclaimState                string                                              `json:"no_overclaim_state"`
	Dependency                      Point12Val0DependencySnapshot                       `json:"dependency"`
	DeterminismContract             Point12Val0ReplayDeterminismContract                `json:"determinism_contract"`
	CompatibilityProfile            Point12Val0ProofPackCompatibilityProfile            `json:"compatibility_profile"`
	Manifest                        Point12Val0SignedProofPackManifest                  `json:"manifest"`
	ReplayAssessment                Point12Val0ReplayAssessment                         `json:"replay_assessment"`
	RedactionBoundary               Point12Val0RedactionBoundary                        `json:"redaction_boundary"`
	FinancialEvidenceSupportProfile Point12Val0FinancialInsuranceEvidenceSupportProfile `json:"financial_evidence_support_profile"`
	ProvenanceProfile               Point12Val0ToolchainAgentProvenanceProfile          `json:"provenance_profile"`
	NoOverclaimReview               Point12Val0NoOverclaimReview                        `json:"no_overclaim_review"`
}

func point12Val0PrematurePassToken() string {
	return strings.Join([]string{"point", "12", "pass"}, "_")
}

func point12Val0ContainsPrematurePassToken(values ...string) bool {
	normalizedToken := formalNoOverclaimNormalizePublicText(point12Val0PrematurePassToken())
	for _, value := range values {
		normalized := formalNoOverclaimNormalizePublicText(value)
		if normalized == "" {
			continue
		}
		if formalNoOverclaimContainsForbidden(normalized, normalizedToken) {
			return true
		}
	}
	return false
}

func point12Val0ProofPackStates() []string {
	return []string{
		Point12Val0ProofPackStateDraft,
		Point12Val0ProofPackStateGenerated,
		Point12Val0ProofPackStateSignedMetadataValidated,
		Point12Val0ProofPackStateReplayable,
		Point12Val0ProofPackStatePartiallyReplayable,
		Point12Val0ProofPackStateRedactedLimited,
		Point12Val0ProofPackStateTampered,
		Point12Val0ProofPackStateUnsupported,
		Point12Val0ProofPackStateExpired,
		Point12Val0ProofPackStateRevoked,
		Point12Val0ProofPackStateSuperseded,
		Point12Val0ProofPackStateBlocked,
	}
}

func point12Val0ReplayResults() []string {
	return []string{
		Point12Val0ReplayResultSameDecision,
		Point12Val0ReplayResultDifferentDecision,
		Point12Val0ReplayResultBlockedReplay,
		Point12Val0ReplayResultInsufficientEvidence,
		Point12Val0ReplayResultUnsupportedVersion,
		Point12Val0ReplayResultTamperDetected,
		Point12Val0ReplayResultRedactedLimitations,
		Point12Val0ReplayResultPolicyMismatch,
		Point12Val0ReplayResultEngineMismatch,
		Point12Val0ReplayResultSchemaMismatch,
		Point12Val0ReplayResultEvidenceMismatch,
		Point12Val0ReplayResultClaimMismatch,
		Point12Val0ReplayResultGovernanceMismatch,
	}
}

func point12Val0HashValid(value string) bool {
	if !formalRawExactNonEmpty(value) || strings.Contains(value, " ") || !strings.HasPrefix(value, "sha256:") {
		return false
	}
	digest := strings.TrimPrefix(value, "sha256:")
	if len(digest) != 64 {
		return false
	}
	for _, char := range digest {
		switch {
		case char >= '0' && char <= '9':
		case char >= 'a' && char <= 'f':
		default:
			return false
		}
	}
	return true
}

func point12Val0VersionIdentityValid(value string) bool {
	return formalRawExactNonEmpty(value) && point11Val0IdentityValueValid(value) && !strings.Contains(value, "/")
}

func point12Val0StringListValid(values []string, validator func(string) bool) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !formalRawExactNonEmpty(value) || !validator(value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point12Val0OptionalStringListValid(values []string, validator func(string) bool) bool {
	if len(values) == 0 {
		return true
	}
	return point12Val0StringListValid(values, validator)
}

func point12Val0AuthorityContextRefListValid(values []string) bool {
	return point12Val0StringListValid(values, point12Val0AuthorityContextRefValid)
}

func point12Val0PrimaryAuthorityContextRefValid(primary string, values []string) bool {
	if !point12Val0AuthorityContextRefValid(primary) || !point12Val0AuthorityContextRefListValid(values) {
		return false
	}
	for _, value := range values {
		if value == primary {
			return true
		}
	}
	return false
}

func point12Val0ExactStringSetMatch(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	seen := map[string]int{}
	for _, value := range left {
		if !formalRawExactNonEmpty(value) {
			return false
		}
		seen[value]++
	}
	for _, value := range right {
		if !formalRawExactNonEmpty(value) {
			return false
		}
		if seen[value] == 0 {
			return false
		}
		seen[value]--
	}
	for _, count := range seen {
		if count != 0 {
			return false
		}
	}
	return true
}

func point12Val0ExactOneOf(value string, allowed []string) bool {
	if !formalRawExactNonEmpty(value) {
		return false
	}
	for _, candidate := range allowed {
		if value == candidate {
			return true
		}
	}
	return false
}

func point12Val0FirstValue(values []string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func point12Val0RawCanonicalRefWithPrefixes(value string, prefixes []string) bool {
	return formalRawExactNonEmpty(value) && point11Val0CanonicalRefWithPrefixes(value, prefixes)
}

func point12Val0RawTimestampValid(value string) bool {
	if !formalRawExactNonEmpty(value) {
		return false
	}
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return false
	}
	return parsed.UTC().Format(time.RFC3339) == value
}

func point12Val0ProofPackRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"proof_pack_"})
}

func point12Val0DecisionRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"decision_", "enforcement_"})
}

func point12Val0ArtifactRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"artifact_", "sbom_", "package_"})
}

func point12Val0PolicyRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"policy_", "point11_policy_"})
}

func point12Val0ClaimRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"claim_", "claim:"})
}

func point12Val0GovernanceEventRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"governance_event_"})
}

func point12Val0ClosureManifestRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"manifest_", "closure_manifest_"})
}

func point12Val0DependencySnapshotRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"dependency_snapshot_", "dependency_review_"})
}

func point12Val0AuthorityContextRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{
		"policy_",
		"claim_",
		"governance_event_",
		"authority_context_",
	})
}

func point12Val0CompatibilityProfileRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"compatibility_profile_"})
}

func point12Val0RedactionManifestRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"redaction_manifest_"})
}

func point12Val0RetentionClassRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"retention_class_"})
}

func point12Val0SigningKeyRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"signing_key_", "metadata_signing_key_"})
}

func point12Val0SignatureRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"signature_", "signature_metadata_"})
}

func point12Val0ToolchainProvenanceRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{
		"toolchain_provenance_",
		"ci_job_",
		"build_tool_",
		"scanner_",
		"sbom_generator_",
		"signing_tool_",
		"execution_environment_",
	})
}

func point12Val0AgentLineageRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"agent_lineage_", "agent_"})
}

func point12Val0AllowedAIEvidenceCandidateTypes() []string {
	return []string{
		"AI_FINDING",
		"AI_RECOMMENDATION",
		"AI_PATCH_PROPOSAL",
		"AI_SANDBOX_RESULT",
		"AI_APPROVAL_REQUEST",
		"AI_POST_ACTION_REPORT",
		"AI_MODEL_CANDIDATE",
		"AI_EVIDENCE_SUMMARY",
	}
}

func point12Val0BlockedAIEvidenceCandidateTypes() []string {
	return []string{
		"AI_PASS",
		"AI_CERTIFIED",
		"AI_PRODUCTION_APPROVED",
		"AI_COMPLIANCE_APPROVED",
		"AI_DEPLOYED",
		"AI_AUTOPATCHED",
		"AI_LEGAL_PROOF",
		"AI_PUBLIC_BADGE",
	}
}

func point12Val0AIEvidenceCandidateTypeValid(value string) bool {
	for _, allowed := range point12Val0AllowedAIEvidenceCandidateTypes() {
		if value == allowed {
			return true
		}
	}
	return false
}

func point12Val0DefaultAgentLineageRecord() Point12Val0AgentLineageRecord {
	return Point12Val0AgentLineageRecord{
		AgentID:                "agent_lineage_point12_val0_001",
		AgentType:              "AI_RECOMMENDATION",
		ModelOrRuleVersionRef:  "model_version_point12_val0_001",
		PermissionManifestHash: "sha256:6666666666666666666666666666666666666666666666666666666666666666",
		InputEvidenceRefs:      []string{"evidence:point12-proof-pack-evidence-001"},
		AuditID:                "audit_point12_val0_agent_001",
		RecommendationID:       "recommendation_point12_val0_001",
		HumanFeedbackRefs:      []string{"human_feedback_point12_val0_001"},
		LineageInputOnly:       true,
	}
}

func point12Val0AgentLineageBindingFields(record Point12Val0AgentLineageRecord) []point12Val0AIEvidenceCandidateBindingField {
	return []point12Val0AIEvidenceCandidateBindingField{
		{FieldName: "agent_id", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "agent_type", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "model_or_rule_version_ref", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "permission_manifest_hash", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "input_evidence_refs", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "input_evidence_hash_refs", BindingClass: point12ValDBindingClassCompatibilityAllowed, Reason: "agent lineage record carries canonical input evidence refs while exact evidence hash refs remain bound by replay, export, and proof-chain context"},
		{FieldName: "tenant_scope", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "tenant scope is bound by canonical proof pack, replay, redaction, export, and proof-chain context rather than duplicated per lineage record"},
		{FieldName: "policy_version", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "policy version remains bound by canonical replay and export context rather than duplicated per lineage record"},
		{FieldName: "engine_version", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "engine version remains bound by canonical replay and export context rather than duplicated per lineage record"},
		{FieldName: "schema_version", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "schema version remains bound by canonical replay and export context rather than duplicated per lineage record"},
		{FieldName: "audit_id", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "recommendation_id", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "human_feedback_refs", BindingClass: point12ValDBindingClassCompatibilityAllowed, Reason: "human feedback refs are optional corroboration and may expand with reviewed evidence without changing canonical identity"},
		{FieldName: "lineage_input_only", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "claims_certification_false", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "claims_source_of_truth_false", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "emits_premature_pass_false", BindingClass: point12ValDBindingClassExactRequired, ValidationRequired: true, MutationTestRequired: true},
		{FieldName: "advisory_only", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "lineage_input_only plus blocked certification/source-of-truth/pass flags are the Val 0 no-authority equivalent for advisory AI evidence candidates"},
		{FieldName: "external_api_allowed", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "network call authority is governed at the toolchain provenance profile level through introduces_network_call_path rather than per lineage record"},
		{FieldName: "production_mutation_allowed", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "production mutation authority is not modeled on lineage records and remains blocked by higher-layer governance gates"},
		{FieldName: "canonical_mutation_allowed", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "canonical mutation authority is not modeled on lineage records and remains blocked by higher-layer governance gates"},
		{FieldName: "pass_allowed", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "lineage records cannot create point pass authority and premature pass emission is blocked explicitly"},
		{FieldName: "approval_required", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "human approval for production-impacting actions is governed by bounded deployment and runtime models, not by Val 0 advisory lineage records"},
		{FieldName: "generated_diff_hash", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "generated patch diff hashing belongs to a future bounded patch-proposal model and is not carried by the current lineage record"},
		{FieldName: "sandbox_result_refs", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "sandbox result refs are not represented on the current lineage record and would require a broader bounded execution model"},
		{FieldName: "approval_request_ref", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "approval request refs are not represented on the current lineage record and belong to approval workflow models"},
		{FieldName: "approval_event_ref", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "approval event refs are not represented on the current lineage record and belong to approval workflow models"},
		{FieldName: "reviewer_ref", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "reviewer refs are not represented on the current lineage record and belong to explicit review workflow models"},
	}
}

func point12Val0AgentLineageBindingMatrixValid(record Point12Val0AgentLineageRecord) bool {
	requiredExact := map[string]bool{
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
	return point12Val0AgentLineageBindingFieldsValid(point12Val0AgentLineageBindingFields(record), requiredExact)
}

func point12Val0AgentLineageBindingFieldsValid(fields []point12Val0AIEvidenceCandidateBindingField, requiredExact map[string]bool) bool {
	for _, field := range fields {
		switch field.BindingClass {
		case point12ValDBindingClassExactRequired:
			if !field.ValidationRequired || !field.MutationTestRequired {
				return false
			}
			if _, ok := requiredExact[field.FieldName]; ok {
				requiredExact[field.FieldName] = true
			}
		case point12ValDBindingClassCompatibilityAllowed, point12ValDBindingClassAdvisoryOnly, point12ValDBindingClassIntentionallyNotBound:
			if strings.TrimSpace(field.Reason) == "" {
				return false
			}
		default:
			return false
		}
	}
	for _, seen := range requiredExact {
		if !seen {
			return false
		}
	}
	return true
}

func point12Val0PrimaryExpectedAgentLineageRecord() Point12Val0AgentLineageRecord {
	return point12Val0DefaultAgentLineageRecord()
}

func point12Val0AgentLineageExactBindingValid(record Point12Val0AgentLineageRecord) bool {
	expected := point12Val0PrimaryExpectedAgentLineageRecord()
	if !formalRawExactNonEmpty(expected.AgentID) {
		return false
	}
	return record.AgentID == expected.AgentID &&
		record.AgentType == expected.AgentType &&
		record.ModelOrRuleVersionRef == expected.ModelOrRuleVersionRef &&
		record.PermissionManifestHash == expected.PermissionManifestHash &&
		point12Val0ExactStringSetMatch(record.InputEvidenceRefs, expected.InputEvidenceRefs) &&
		record.AuditID == expected.AuditID &&
		record.RecommendationID == expected.RecommendationID &&
		record.LineageInputOnly == expected.LineageInputOnly &&
		record.ClaimsCertification == expected.ClaimsCertification &&
		record.ClaimsSourceOfTruth == expected.ClaimsSourceOfTruth &&
		record.EmitsPrematurePass == expected.EmitsPrematurePass
}

func point12Val0EvidenceRefValid(value string) bool {
	if !formalRawExactNonEmpty(value) || !point11Val0IdentityValueValid(value) {
		return false
	}
	if strings.Contains(value, " ") || strings.Contains(value, "/") {
		return false
	}
	normalized := point11Val0NormalizeText(value)
	for _, blocked := range []string{"cross-tenant", "other-tenant", "global", "unscoped", "wildcard", "all-tenants"} {
		if strings.Contains(normalized, blocked) {
			return false
		}
	}
	return strings.HasPrefix(value, "evidence:") || strings.HasPrefix(value, "evidence_")
}

func point12Val0EvidenceRefsValid(values []string) bool {
	return point12Val0StringListValid(values, point12Val0EvidenceRefValid)
}

func point12Val0EvidenceHashRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"evidence_hash_"})
}

func point12Val0AuditRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"audit_"})
}

func point12Val0RecommendationRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"recommendation_"})
}

func point12Val0PermissionManifestHashValid(value string) bool {
	return point12Val0HashValid(value)
}

func point12Val0HumanFeedbackRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"feedback_", "human_feedback_"})
}

func point12Val0CIJobRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"ci_job_"})
}

func point12Val0VersionRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{
		"build_tool_",
		"scanner_",
		"sbom_generator_",
		"signing_tool_",
		"agent_model_",
		"model_version_",
		"rule_version_",
		"engine_version_",
		"schema_version_",
	})
}

func point12Val0ExecutionEnvironmentClassRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"execution_environment_"})
}

func point12Val0ReplayAssessmentRefValid(value string) bool {
	return point12Val0RawCanonicalRefWithPrefixes(value, []string{"replay_assessment_"})
}

func point12Val0RedactionFieldValuesValid(values []string) bool {
	if len(values) == 0 {
		return true
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !formalRawExactTokenValid(value, point11Val0IdentityValueValid) || strings.Contains(value, "/") {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point12Val0OptionalClaimTextListValid(values []string) bool {
	if len(values) == 0 {
		return true
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			return false
		}
		normalized := point11Val0NormalizeText(trimmed)
		if _, exists := seen[normalized]; exists {
			return false
		}
		seen[normalized] = struct{}{}
	}
	return true
}

func point12Val0ClaimTextOverlap(left []string, right []string) bool {
	if len(left) == 0 || len(right) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range left {
		normalized := point11Val0NormalizeText(value)
		if normalized == "" {
			continue
		}
		seen[normalized] = struct{}{}
	}
	for _, value := range right {
		normalized := point11Val0NormalizeText(value)
		if normalized == "" {
			continue
		}
		if _, exists := seen[normalized]; exists {
			return true
		}
	}
	return false
}

func point12Val0CompatibilityValueValid(value string) bool {
	return point12Val0ExactOneOf(value, []string{
		point12Val0CompatibilityExactRequired,
		point12Val0CompatibilityCompatibleAllowed,
		point12Val0CompatibilityBlocked,
	})
}

func point12Val0ReplayModeValid(value string) bool {
	return point12Val0ExactOneOf(value, []string{
		point12Val0ReplayModeOriginalContext,
		point12Val0ReplayModeCurrentPolicyContext,
		point12Val0ReplayModeComparisonMode,
	})
}

func point12Val0ProfileTypeValid(value string) bool {
	return point12Val0ExactOneOf(value, []string{
		point12Val0ProfileTypeFinancialReview,
		point12Val0ProfileTypeInsuranceReview,
		point12Val0ProfileTypeAuditReview,
	})
}

func point12Val0ForbiddenClaims() []string {
	return []string{
		"certified",
		"guaranteed secure",
		"regulator-approved",
		"production approved",
		"deployment approved",
		"compliance guaranteed",
		"public badge",
		"global truth",
		"official authority",
		"absolute proof",
		"absolute incontestability",
		"mathematical accountability",
		"premium reduction",
		"lowers insurance premium",
		"credit rating",
		"financial guarantee",
		"compliance-as-an-asset",
		"proves dora compliance",
		"proves hipaa compliance",
		"proves insurance eligibility",
		"cyber insurance premium",
		"lower insurance premium",
		"ai approved",
		"ai-approved",
		"ai decision",
		"ai legal proof",
		"autonomous remediation",
		"continuous compliance attestation",
		"source of truth",
		"canonical truth",
	}
}

func point12Val0AllowedClaims() []string {
	return []string{
		"bounded claim",
		"evidence-linked governance decision",
		"policy-bound decision support",
		"audit-ready governance trail",
		"not a certification",
		"not regulator approval",
		"not production approval",
		"not deployment approval",
		"not compliance guarantee",
		"advisory projection only",
		"canonical evidence spine remains source of truth",
		"this proof pack contains evidence that may support customer, auditor, financial, or insurance review.",
	}
}

func formalNoOverclaimNormalizeText(value string) string {
	return formalNoOverclaimNormalizeTextWithOptions(value, false)
}

func formalNoOverclaimNormalizePublicText(value string) string {
	return formalNoOverclaimNormalizeTextWithOptions(value, true)
}

func formalNoOverclaimNormalizeTextWithOptions(value string, underscoreBoundary bool) string {
	compatible := norm.NFKD.String(value)
	folded := strings.Map(func(r rune) rune {
		if formalNoOverclaimIgnorableFormat(r) {
			return -1
		}
		if formalNoOverclaimBoundarySeparator(r, underscoreBoundary) {
			return ' '
		}
		return formalNoOverclaimConfusableLetter(r)
	}, strings.ToLower(strings.TrimSpace(compatible)))
	return strings.Join(strings.Fields(folded), " ")
}

func formalNoOverclaimIgnorableFormat(r rune) bool {
	return unicode.Is(unicode.Cf, r) || unicode.IsMark(r)
}

func formalNoOverclaimBoundarySeparator(r rune, underscoreBoundary bool) bool {
	if r == '_' && !underscoreBoundary {
		return false
	}
	if formalNoOverclaimDashSeparator(r) {
		return true
	}
	return unicode.IsPunct(r) || unicode.IsSymbol(r)
}

func formalNoOverclaimDashSeparator(r rune) bool {
	switch r {
	case '-', '\u00ad', '\u2010', '\u2011', '\u2012', '\u2013', '\u2014', '\u2015', '\u2043', '\u2212', '\ufe58', '\ufe63', '\uff0d':
		return true
	default:
		return false
	}
}

func formalNoOverclaimConfusableLetter(r rune) rune {
	if r >= '\uff41' && r <= '\uff5a' {
		return 'a' + (r - '\uff41')
	}
	switch r {
	case '\u0430', '\u03b1', '\u0251', '\u1d00': // Cyrillic/Greek/Latin/small-cap a
		return 'a'
	case '\u0432', '\u0299': // Cyrillic/small-cap b
		return 'b'
	case '\u0441', '\u03f2', '\u1d04': // Cyrillic/Greek/small-cap c
		return 'c'
	case '\u0501', '\u03b4', '\u1d05', '\u217e': // Cyrillic/Greek/small-cap d
		return 'd'
	case '\u0435', '\u0454', '\u03b5', '\u025b', '\u04bd', '\u212e', '\u212f', '\u1d07': // Cyrillic/Greek/Latin/small-cap e
		return 'e'
	case '\ua730': // small-cap f
		return 'f'
	case '\u0261', '\u0262': // Latin/small-cap g
		return 'g'
	case '\u04bb', '\u029c': // Cyrillic/small-cap h
		return 'h'
	case '\u0456', '\u03b9', '\u026a', '\u0269', '\u0131': // Cyrillic/Greek/Latin/small-cap i
		return 'i'
	case '\u0458', '\u1d0a': // Cyrillic/small-cap j
		return 'j'
	case '\u043a', '\u03ba', '\u1d0b': // Cyrillic/Greek/small-cap k
		return 'k'
	case '\u217c', '\u04cf', '\u01c0', '\u026d', '\u026b', '\u0142', '\u019a', '\u1e37', '\u029f': // Roman/Cyrillic/Latin/small-cap l
		return 'l'
	case '\u043c', '\u1d0d': // Cyrillic/small-cap m
		return 'm'
	case '\u043d', '\u043f', '\u03b7', '\u0578', '\u0272', '\u0273', '\u0274', '\u014b', '\u019e', '\ua791': // Cyrillic/Greek/Armenian/Latin/small-cap n
		return 'n'
	case '\u043e', '\u03bf', '\u0585', '\u0254', '\u1d0f': // Cyrillic/Greek/Armenian/Latin o
		return 'o'
	case '\u0440', '\u03c1', '\u1d18': // Cyrillic/Greek/small-cap p
		return 'p'
	case '\u051b': // Cyrillic qa
		return 'q'
	case '\u0433', '\u0280': // Cyrillic/small-cap r
		return 'r'
	case '\u0455', '\ua731': // Cyrillic/small-cap s
		return 's'
	case '\u0442', '\u03c4', '\u1d1b': // Cyrillic/Greek/small-cap t
		return 't'
	case '\u03c5', '\u1d1c', '\uab4e', '\u028a': // Greek/Latin/small-cap U confusables
		return 'u'
	case '\u0445', '\u03c7': // Cyrillic/Greek x
		return 'x'
	case '\u0443', '\u028f': // Cyrillic/small-cap y
		return 'y'
	case '\u03bd', '\u0475', '\u1d20', '\u2174': // Greek/Cyrillic/small-cap v
		return 'v'
	case '\u1d21': // small-cap w
		return 'w'
	case '\u1d22': // small-cap z
		return 'z'
	default:
		return r
	}
}

func formalNoOverclaimContainsForbidden(normalized, forbidden string) bool {
	if strings.Contains(normalized, forbidden) {
		return true
	}
	if formalNoOverclaimContainsForbiddenTokenSequence(normalized, forbidden) {
		return true
	}
	compactNormalized := formalNoOverclaimCompact(normalized)
	compactForbidden := formalNoOverclaimCompact(forbidden)
	return compactForbidden != "" && strings.Contains(compactNormalized, compactForbidden)
}

func formalNoOverclaimContainsForbiddenTokenSequence(normalized, forbidden string) bool {
	valueTokens := strings.Fields(normalized)
	phraseTokens := strings.Fields(forbidden)
	if len(valueTokens) == 0 || len(phraseTokens) < 2 || len(valueTokens) < len(phraseTokens) {
		return false
	}
	for start, token := range valueTokens {
		if token != phraseTokens[0] {
			continue
		}
		matched := 1
		for idx := start + 1; idx < len(valueTokens) && matched < len(phraseTokens); idx++ {
			if valueTokens[idx] != phraseTokens[matched] {
				continue
			}
			matched++
		}
		if matched == len(phraseTokens) {
			return true
		}
	}
	return false
}

func formalNoOverclaimCompact(value string) string {
	return strings.Join(strings.Fields(value), "")
}

func formalNoOverclaimForbiddenCompactAcrossValues(values []string, allowed []bool, phrase string) bool {
	if len(values) != len(allowed) {
		return false
	}
	compactPhrase := formalNoOverclaimCompact(phrase)
	if compactPhrase == "" {
		return false
	}
	for start := range values {
		var compact strings.Builder
		allAllowed := true
		parts := 0
		for end := start; end < len(values); end++ {
			part := formalNoOverclaimCompact(values[end])
			if part == "" {
				continue
			}
			compact.WriteString(part)
			allAllowed = allAllowed && allowed[end]
			parts++
			if parts > 1 && !allAllowed && strings.Contains(compact.String(), compactPhrase) {
				return true
			}
		}
	}
	return false
}

func point12Val0ContainsForbiddenClaim(values ...string) bool {
	return point12Val0ContainsForbiddenClaimWithNormalizer(formalNoOverclaimNormalizePublicText, values...)
}

func point12Val0ContainsForbiddenInternalClaim(values ...string) bool {
	return point12Val0ContainsForbiddenClaimWithNormalizer(formalNoOverclaimNormalizeText, values...)
}

func point12Val0ContainsForbiddenClaimWithNormalizer(normalize func(string) string, values ...string) bool {
	allowed := map[string]struct{}{}
	for _, value := range point12Val0AllowedClaims() {
		allowed[normalize(value)] = struct{}{}
	}
	crossNormalizedParts := make([]string, 0, len(values))
	crossPartAllowed := make([]bool, 0, len(values))
	for _, value := range values {
		normalized := normalize(value)
		if normalized == "" {
			continue
		}
		_, isAllowed := allowed[normalized]
		crossNormalizedParts = append(crossNormalizedParts, normalized)
		crossPartAllowed = append(crossPartAllowed, isAllowed)
		if _, ok := allowed[normalized]; ok {
			continue
		}
		for _, forbidden := range point12Val0ForbiddenClaims() {
			if formalNoOverclaimContainsForbidden(normalized, normalize(forbidden)) {
				return true
			}
		}
	}
	for _, forbidden := range point12Val0ForbiddenClaims() {
		if point12Val0ForbiddenPhraseAcrossValues(crossNormalizedParts, crossPartAllowed, normalize(forbidden)) {
			return true
		}
	}
	return false
}

func point12Val0ForbiddenPhraseAcrossValues(values []string, allowed []bool, phrase string) bool {
	if len(values) != len(allowed) {
		return false
	}
	if formalNoOverclaimForbiddenCompactAcrossValues(values, allowed, phrase) {
		return true
	}
	phraseTokens := strings.Fields(phrase)
	if len(phraseTokens) < 2 {
		return false
	}
	matched := 0
	distinctBuckets := 0
	lastBucket := -1
	matchedIncludesNonAllowed := false
	for bucketIndex, value := range values {
		bucketTokens := strings.Fields(value)
		if len(bucketTokens) == 0 {
			continue
		}
		for _, token := range bucketTokens {
			if token != phraseTokens[matched] {
				continue
			}
			if bucketIndex != lastBucket {
				distinctBuckets++
				lastBucket = bucketIndex
			}
			if !allowed[bucketIndex] {
				matchedIncludesNonAllowed = true
			}
			matched++
			if matched == len(phraseTokens) {
				return distinctBuckets > 1 && matchedIncludesNonAllowed
			}
		}
	}
	return false
}

func SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD Point11ValDFoundation, review Point12Val0Point11ReviewContext) Point12Val0DependencySnapshot {
	reviewPrerequisites := append([]string{}, valD.ReviewPrerequisites...)
	reviewPrerequisites = append(reviewPrerequisites, review.ReviewPrerequisites...)
	return Point12Val0DependencySnapshot{
		UpstreamCurrentState:             valD.CurrentState,
		UpstreamDependencyState:          valD.DependencyState,
		UpstreamPointID:                  valD.PassClosureManifest.PointID,
		UpstreamWaveID:                   valD.PassClosureManifest.WaveID,
		UpstreamPassClosureManifestState: valD.PassClosureManifestState,
		UpstreamFinalPassGateState:       valD.FinalPassGateState,
		UpstreamPoint11PassAllowed:       valD.FinalPassGate.Point11PassAllowed,
		UpstreamPoint11PassToken:         valD.Point11PassToken,
		UpstreamClosureManifestRef:       valD.PassClosureManifest.ManifestID,
		PolicyAuthorityContextRefs:       append([]string{}, valD.QualityMap.PolicyRefs...),
		ClaimAuthorityContextRefs:        append([]string{}, valD.QualityMap.ClaimRefs...),
		GovernanceAuthorityContextRefs:   append([]string{}, valD.QualityMap.GovernanceEventRefs...),
		// Preserve the full upstream authority context set. The singular refs remain
		// only as primary refs for the current Val 0 manifest skeleton contract.
		PolicyAuthorityContextRef:                   point12Val0FirstValue(valD.QualityMap.PolicyRefs),
		ClaimAuthorityContextRef:                    point12Val0FirstValue(valD.QualityMap.ClaimRefs),
		GovernanceAuthorityContextRef:               point12Val0FirstValue(valD.QualityMap.GovernanceEventRefs),
		ProjectionDisclaimer:                        valD.ProjectionDisclaimer,
		UpstreamPoint11PassObservedOutsideFinalPath: valD.FinalPassGate.Point11PassObservedOutsideFinalClosure,
		SnapshotFromComputedOutput:                  review.SnapshotFromComputedOutput,
		ReviewPrerequisites:                         reviewPrerequisites,
	}
}

func point12Val0DependencyReviewContextModel() Point12Val0Point11ReviewContext {
	return Point12Val0Point11ReviewContext{
		SnapshotFromComputedOutput: true,
	}
}

func point12Val0DependencySnapshotModel() Point12Val0DependencySnapshot {
	valD := ComputePoint11ValDFoundation(Point11ValDFoundationModel())
	return SnapshotPoint12Val0DependencyFromComputedPoint11ValD(valD, point12Val0DependencyReviewContextModel())
}

func EvaluatePoint12Val0DependencyState(model Point12Val0DependencySnapshot) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!model.SnapshotFromComputedOutput ||
		!point12Val0ClosureManifestRefValid(model.UpstreamClosureManifestRef) ||
		!point12Val0PrimaryAuthorityContextRefValid(model.PolicyAuthorityContextRef, model.PolicyAuthorityContextRefs) ||
		!point12Val0PrimaryAuthorityContextRefValid(model.ClaimAuthorityContextRef, model.ClaimAuthorityContextRefs) ||
		!point12Val0PrimaryAuthorityContextRefValid(model.GovernanceAuthorityContextRef, model.GovernanceAuthorityContextRefs) ||
		model.UpstreamPoint11PassObservedOutsideFinalPath ||
		model.UpstreamPointID != point11ValDPointID ||
		model.UpstreamWaveID != point11ValDWaveID {
		return Point12Val0DependencyStateBlocked
	}
	if model.UpstreamCurrentState == Point11ValDStateBlocked ||
		model.UpstreamDependencyState == Point11ValDDependencyStateBlocked {
		return Point12Val0DependencyStateBlocked
	}
	if model.UpstreamPassClosureManifestState == Point11ValDPassClosureManifestStateBlocked ||
		model.UpstreamFinalPassGateState == Point11ValDFinalPassGateStateBlocked {
		if model.UpstreamCurrentState == Point11ValDStateReviewRequired ||
			model.UpstreamDependencyState == Point11ValDDependencyStateReviewRequired ||
			len(model.ReviewPrerequisites) > 0 {
			return Point12Val0DependencyStateReviewRequired
		}
		return Point12Val0DependencyStateBlocked
	}
	if model.UpstreamCurrentState == Point11ValDStateReviewRequired ||
		model.UpstreamDependencyState == Point11ValDDependencyStateReviewRequired ||
		len(model.ReviewPrerequisites) > 0 {
		return Point12Val0DependencyStateReviewRequired
	}
	if model.UpstreamCurrentState != Point11ValDStateActive ||
		model.UpstreamDependencyState != Point11ValDDependencyStateActive ||
		model.UpstreamPassClosureManifestState != Point11ValDPassClosureManifestStateActive ||
		model.UpstreamFinalPassGateState != Point11ValDFinalPassGateStateActive ||
		!model.UpstreamPoint11PassAllowed ||
		model.UpstreamPoint11PassToken != point11ValDPoint11PassToken {
		return Point12Val0DependencyStateBlocked
	}
	return Point12Val0DependencyStateActive
}

func EvaluatePoint12Val0DeterminismContractState(model Point12Val0ReplayDeterminismContract) string {
	if !formalRawExactValid(model.ReplayMode, point12Val0ReplayModeValid) ||
		!point12Val0ExactOneOf(model.UnsupportedBehavior, []string{
			point12Val0UnsupportedBehaviorBlockedReplay,
			point12Val0UnsupportedBehaviorUnsupportedVersion,
		}) {
		return Point12Val0DeterminismContractStateBlocked
	}
	if model.ReplayMode == point12Val0ReplayModeOriginalContext {
		if !model.StableEvidenceIdentityRequired ||
			!model.StablePolicyIdentityRequired ||
			!model.StableEngineIdentityRequired ||
			!model.StableSchemaIdentityRequired ||
			!model.StableTenantScopeRequired ||
			!model.StableArtifactIdentityRequired {
			return Point12Val0DeterminismContractStateBlocked
		}
	}
	if model.ReplayMode == point12Val0ReplayModeComparisonMode && !model.DriftExplanationRequired {
		return Point12Val0DeterminismContractStateBlocked
	}
	return Point12Val0DeterminismContractStateActive
}

func EvaluatePoint12Val0CompatibilityProfileState(model Point12Val0ProofPackCompatibilityProfile) string {
	if !point12Val0CompatibilityProfileRefValid(model.CompatibilityProfileRef) ||
		!formalRawExactValid(model.ReplayMode, point12Val0ReplayModeValid) ||
		!formalRawExactValid(model.PolicyCompatibility, point12Val0CompatibilityValueValid) ||
		!formalRawExactValid(model.EngineCompatibility, point12Val0CompatibilityValueValid) ||
		!formalRawExactValid(model.SchemaCompatibility, point12Val0CompatibilityValueValid) ||
		model.EvidenceCompatibility != point12Val0EvidenceCompatibilityExactHashRequired ||
		!point12Val0ExactOneOf(model.ToolchainCompatibility, []string{
			point12Val0ToolchainCompatibilityRequiredIfDecisive,
			point12Val0ToolchainCompatibilityNotRequired,
		}) ||
		!point12Val0ExactOneOf(model.UnsupportedBehavior, []string{
			point12Val0UnsupportedBehaviorBlockedReplay,
			point12Val0UnsupportedBehaviorUnsupportedVersion,
		}) {
		return Point12Val0CompatibilityProfileStateBlocked
	}
	if (model.PolicyCompatibility == point12Val0CompatibilityCompatibleAllowed ||
		model.EngineCompatibility == point12Val0CompatibilityCompatibleAllowed ||
		model.SchemaCompatibility == point12Val0CompatibilityCompatibleAllowed) &&
		!point12Val0EvidenceRefsValid(model.CompatibilityEvidenceRefs) {
		return Point12Val0CompatibilityProfileStateBlocked
	}
	if model.ReplayMode == point12Val0ReplayModeComparisonMode && !model.DecisionDriftExplanationRequired {
		return Point12Val0CompatibilityProfileStateBlocked
	}
	return Point12Val0CompatibilityProfileStateActive
}

func EvaluatePoint12Val0ManifestState(model Point12Val0SignedProofPackManifest) string {
	if !point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12Val0DecisionRefValid(model.DecisionID) ||
		model.PointID != point12Val0PointID ||
		model.WaveID != point12Val0WaveID ||
		!point12Val0ExactOneOf(model.ProofPackState, point12Val0ProofPackStates()) ||
		!formalRawExactTokenValid(model.TenantScope, point11Val0ScopeValid) ||
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
		!point12Val0ClosureManifestRefValid(model.UpstreamClosureManifestRef) ||
		!point12Val0DependencySnapshotRefValid(model.DependencySnapshotRef) ||
		!point12Val0AuthorityContextRefValid(model.PolicyAuthorityContextRef) ||
		!point12Val0AuthorityContextRefValid(model.ClaimAuthorityContextRef) ||
		!point12Val0AuthorityContextRefValid(model.GovernanceAuthorityContextRef) ||
		!point12Val0CompatibilityProfileRefValid(model.CompatibilityProfileRef) ||
		!point12Val0RawTimestampValid(model.GeneratedAt) ||
		!point12Val0VersionIdentityValid(model.FreshnessWindow) ||
		!point12Val0SigningKeyRefValid(model.SigningKeyRef) ||
		!point12Val0SignatureRefValid(model.SignatureRef) ||
		!point12Val0RedactionManifestRefValid(model.RedactionManifestRef) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12Val0RetentionClassRefValid(model.RetentionClassRef) ||
		!point12Val0StringListValid(model.ToolchainProvenanceRefs, point12Val0ToolchainProvenanceRefValid) ||
		!point12Val0StringListValid(model.AgentLineageRefs, point12Val0AgentLineageRefValid) ||
		point12Val0ContainsPrematurePassToken(
			model.ProofPackID,
			model.DecisionID,
			model.SigningKeyRef,
			model.SignatureRef,
		) {
		return Point12Val0ManifestStateBlocked
	}
	switch model.ProofPackState {
	case Point12Val0ProofPackStateTampered,
		Point12Val0ProofPackStateUnsupported,
		Point12Val0ProofPackStateExpired,
		Point12Val0ProofPackStateRevoked,
		Point12Val0ProofPackStateSuperseded,
		Point12Val0ProofPackStateBlocked:
		return Point12Val0ManifestStateBlocked
	}
	return Point12Val0ManifestStateActive
}

func point12Val0MismatchNeedsSpecificResult(actual bool, result, expected string) bool {
	if !actual {
		return false
	}
	return result != expected
}

func point12Val0ReplayAssessmentStateAndReasons(
	model Point12Val0ReplayAssessment,
	contract Point12Val0ReplayDeterminismContract,
	compat Point12Val0ProofPackCompatibilityProfile,
	manifest Point12Val0SignedProofPackManifest,
	redaction Point12Val0RedactionBoundary,
) (string, []string) {
	reasons := []string{}
	if !point12Val0ReplayAssessmentRefValid(model.ReplayAssessmentID) ||
		!point12Val0ExactOneOf(model.ProofPackState, point12Val0ProofPackStates()) ||
		!point12Val0ExactOneOf(model.ReplayResult, point12Val0ReplayResults()) ||
		!point12Val0DependencySnapshotRefValid(model.DeterminismContractRef) ||
		!point12Val0CompatibilityProfileRefValid(model.CompatibilityProfileRef) ||
		!point12Val0PolicyRefValid(model.OriginalPolicyRef) ||
		!point12Val0PolicyRefValid(model.ReplayPolicyRef) ||
		!point12Val0HashValid(model.OriginalPolicyHash) ||
		!point12Val0HashValid(model.ReplayPolicyHash) ||
		!point12Val0HashValid(model.OriginalEngineHash) ||
		!point12Val0HashValid(model.ReplayEngineHash) ||
		!point12Val0VersionIdentityValid(model.OriginalSchemaVersion) ||
		!point12Val0VersionIdentityValid(model.ReplaySchemaVersion) ||
		!point12Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point12Val0EvidenceRefsValid(model.ReplayEvidenceRefs) ||
		!point12Val0StringListValid(model.EvidenceHashRefs, point12Val0EvidenceHashRefValid) ||
		!point12Val0StringListValid(model.ReplayEvidenceHashRefs, point12Val0EvidenceHashRefValid) ||
		!point12Val0StringListValid(model.ClaimRefs, point12Val0ClaimRefValid) ||
		!point12Val0StringListValid(model.ReplayClaimRefs, point12Val0ClaimRefValid) ||
		!point12Val0StringListValid(model.GovernanceEventRefs, point12Val0GovernanceEventRefValid) ||
		!point12Val0StringListValid(model.ReplayGovernanceRefs, point12Val0GovernanceEventRefValid) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "replay_assessment_identity_invalid")
	}

	if contract.ReplayMode == point12Val0ReplayModeComparisonMode &&
		contract.DriftExplanationRequired &&
		!point11Val0IdentityValueValid(model.DriftExplanation) {
		reasons = append(reasons, "replay_assessment_missing_drift_explanation")
	}

	if point12Val0ContainsPrematurePassToken(model.ReplayAssessmentID, model.DriftExplanation) {
		reasons = append(reasons, "replay_assessment_premature_pass_token")
	}

	policyMismatch := model.OriginalPolicyRef != model.ReplayPolicyRef || model.OriginalPolicyHash != model.ReplayPolicyHash
	engineMismatch := model.OriginalEngineHash != model.ReplayEngineHash
	schemaMismatch := model.OriginalSchemaVersion != model.ReplaySchemaVersion
	evidenceMismatch := !point12Val0ExactStringSetMatch(model.EvidenceRefs, model.ReplayEvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.EvidenceHashRefs, model.ReplayEvidenceHashRefs)
	claimMismatch := !point12Val0ExactStringSetMatch(model.ClaimRefs, model.ReplayClaimRefs)
	governanceMismatch := !point12Val0ExactStringSetMatch(model.GovernanceEventRefs, model.ReplayGovernanceRefs)

	if !model.DecisiveEvidencePresent &&
		model.ReplayResult != Point12Val0ReplayResultInsufficientEvidence &&
		model.ReplayResult != Point12Val0ReplayResultBlockedReplay {
		reasons = append(reasons, "replay_assessment_missing_decisive_evidence")
	}
	if redaction.RedactionAffectsReplay &&
		model.ReplayResult != Point12Val0ReplayResultRedactedLimitations &&
		model.ReplayResult != Point12Val0ReplayResultInsufficientEvidence &&
		model.ReplayResult != Point12Val0ReplayResultBlockedReplay {
		reasons = append(reasons, "replay_assessment_redaction_limitation_missing")
	}

	switch model.ProofPackState {
	case Point12Val0ProofPackStateTampered:
		if model.ReplayResult != Point12Val0ReplayResultTamperDetected &&
			model.ReplayResult != Point12Val0ReplayResultBlockedReplay {
			reasons = append(reasons, "replay_assessment_tamper_result_invalid")
		}
	case Point12Val0ProofPackStateUnsupported:
		if model.ReplayResult != Point12Val0ReplayResultUnsupportedVersion &&
			model.ReplayResult != Point12Val0ReplayResultBlockedReplay {
			reasons = append(reasons, "replay_assessment_unsupported_result_invalid")
		}
	case Point12Val0ProofPackStateExpired, Point12Val0ProofPackStateRevoked, Point12Val0ProofPackStateSuperseded:
		if model.ReplayResult == Point12Val0ReplayResultSameDecision {
			reasons = append(reasons, "replay_assessment_invalidated_pack_cannot_replay_same_decision")
		}
	}

	if compat.PolicyCompatibility == point12Val0CompatibilityBlocked &&
		model.ReplayResult != Point12Val0ReplayResultBlockedReplay {
		reasons = append(reasons, "replay_assessment_policy_compatibility_blocked")
	}
	if compat.EngineCompatibility == point12Val0CompatibilityBlocked &&
		model.ReplayResult != Point12Val0ReplayResultBlockedReplay {
		reasons = append(reasons, "replay_assessment_engine_compatibility_blocked")
	}
	if compat.SchemaCompatibility == point12Val0CompatibilityBlocked &&
		model.ReplayResult != Point12Val0ReplayResultBlockedReplay {
		reasons = append(reasons, "replay_assessment_schema_compatibility_blocked")
	}
	if point12Val0MismatchNeedsSpecificResult(compat.PolicyCompatibility == point12Val0CompatibilityExactRequired && policyMismatch, model.ReplayResult, Point12Val0ReplayResultPolicyMismatch) {
		reasons = append(reasons, "replay_assessment_policy_mismatch_result_invalid")
	}
	if point12Val0MismatchNeedsSpecificResult(compat.EngineCompatibility == point12Val0CompatibilityExactRequired && engineMismatch, model.ReplayResult, Point12Val0ReplayResultEngineMismatch) {
		reasons = append(reasons, "replay_assessment_engine_mismatch_result_invalid")
	}
	if point12Val0MismatchNeedsSpecificResult(compat.SchemaCompatibility == point12Val0CompatibilityExactRequired && schemaMismatch, model.ReplayResult, Point12Val0ReplayResultSchemaMismatch) {
		reasons = append(reasons, "replay_assessment_schema_mismatch_result_invalid")
	}
	if point12Val0MismatchNeedsSpecificResult(policyMismatch && contract.ReplayMode == point12Val0ReplayModeOriginalContext, model.ReplayResult, Point12Val0ReplayResultPolicyMismatch) {
		reasons = append(reasons, "replay_assessment_original_context_requires_policy_mismatch_result")
	}
	if point12Val0MismatchNeedsSpecificResult(evidenceMismatch, model.ReplayResult, Point12Val0ReplayResultEvidenceMismatch) {
		reasons = append(reasons, "replay_assessment_evidence_mismatch_result_invalid")
	}
	if point12Val0MismatchNeedsSpecificResult(claimMismatch, model.ReplayResult, Point12Val0ReplayResultClaimMismatch) {
		reasons = append(reasons, "replay_assessment_claim_mismatch_result_invalid")
	}
	if point12Val0MismatchNeedsSpecificResult(governanceMismatch, model.ReplayResult, Point12Val0ReplayResultGovernanceMismatch) {
		reasons = append(reasons, "replay_assessment_governance_mismatch_result_invalid")
	}
	if model.ReplayResult == Point12Val0ReplayResultSameDecision &&
		(policyMismatch || engineMismatch || schemaMismatch || evidenceMismatch || claimMismatch || governanceMismatch || redaction.RedactionAffectsReplay || !model.DecisiveEvidencePresent) {
		reasons = append(reasons, "replay_assessment_same_decision_overclaims_replay")
	}
	if model.ReplayResult == Point12Val0ReplayResultDifferentDecision &&
		contract.DriftExplanationRequired &&
		!point11Val0IdentityValueValid(model.DriftExplanation) {
		reasons = append(reasons, "replay_assessment_different_decision_requires_explanation")
	}
	if point12Val0ContainsForbiddenClaim(model.DriftExplanation, strings.Join(manifest.ClaimRefs, " ")) {
		reasons = append(reasons, "replay_assessment_overclaim_detected")
	}

	if len(reasons) > 0 {
		return Point12Val0ReplayAssessmentStateBlocked, reasons
	}
	return Point12Val0ReplayAssessmentStateActive, nil
}

func EvaluatePoint12Val0RedactionBoundaryState(model Point12Val0RedactionBoundary) string {
	if !point12Val0RedactionManifestRefValid(model.RedactionManifestRef) ||
		!point12Val0RedactionFieldValuesValid(model.RedactedFields) ||
		!point12Val0OptionalStringListValid(model.RedactionReasons, point11Val0IdentityValueValid) ||
		!point12Val0OptionalClaimTextListValid(model.DisallowedClaimsAfterRedaction) ||
		!point12Val0OptionalClaimTextListValid(model.SurvivingClaimsAfterRedaction) ||
		!point12Val0OptionalClaimTextListValid(model.CustomerVisibleClaimsAfterRedaction) ||
		!point12Val0OptionalClaimTextListValid(model.ExportedClaimsAfterRedaction) ||
		!point12Val0OptionalClaimTextListValid(model.ReplayResultClaims) ||
		!point12Val0ExactOneOf(model.PostRedactionResult, point12Val0ReplayResults()) {
		return Point12Val0RedactionBoundaryStateBlocked
	}
	if len(model.RedactedFields) > 0 {
		if !point12Val0StringListValid(model.RedactionReasons, point11Val0IdentityValueValid) ||
			!formalRawExactTokenValid(model.RedactionApproverRef, point11Val0IdentityValueValid) ||
			!point12Val0GovernanceEventRefValid(model.RedactionApprovalEventRef) ||
			!model.PartialOrAdvisoryOnly {
			return Point12Val0RedactionBoundaryStateBlocked
		}
	}
	if (model.RedactionAffectsDecision || model.RedactionAffectsReplay) &&
		model.PostRedactionResult != Point12Val0ReplayResultInsufficientEvidence &&
		model.PostRedactionResult != Point12Val0ReplayResultBlockedReplay &&
		model.PostRedactionResult != Point12Val0ReplayResultRedactedLimitations {
		return Point12Val0RedactionBoundaryStateBlocked
	}
	if (model.RedactionAffectsDecision || model.RedactionAffectsReplay) && !model.PartialOrAdvisoryOnly {
		return Point12Val0RedactionBoundaryStateBlocked
	}
	if strings.TrimSpace(model.MinimumSafeClaimAfterRedaction) != "" &&
		point12Val0ContainsForbiddenClaim(model.MinimumSafeClaimAfterRedaction) {
		return Point12Val0RedactionBoundaryStateBlocked
	}
	if point12Val0ContainsForbiddenClaim(
		strings.Join(model.SurvivingClaimsAfterRedaction, " "),
		strings.Join(model.CustomerVisibleClaimsAfterRedaction, " "),
		strings.Join(model.ExportedClaimsAfterRedaction, " "),
		strings.Join(model.ReplayResultClaims, " "),
	) {
		return Point12Val0RedactionBoundaryStateBlocked
	}
	if point12Val0ClaimTextOverlap(model.DisallowedClaimsAfterRedaction, model.SurvivingClaimsAfterRedaction) ||
		point12Val0ClaimTextOverlap(model.DisallowedClaimsAfterRedaction, model.CustomerVisibleClaimsAfterRedaction) ||
		point12Val0ClaimTextOverlap(model.DisallowedClaimsAfterRedaction, model.ExportedClaimsAfterRedaction) ||
		point12Val0ClaimTextOverlap(model.DisallowedClaimsAfterRedaction, model.ReplayResultClaims) {
		return Point12Val0RedactionBoundaryStateBlocked
	}
	return Point12Val0RedactionBoundaryStateActive
}

func EvaluatePoint12Val0FinancialEvidenceSupportState(model Point12Val0FinancialInsuranceEvidenceSupportProfile) string {
	if !point12Val0ProfileTypeValid(model.ProfileType) ||
		!point12Val0StringListValid(model.EvidenceSupportCategories, point11Val0IdentityValueValid) ||
		!point12Val0StringListValid(model.RiskContextMetadata, point11Val0IdentityValueValid) ||
		!point12Val0StringListValid(model.Limitations, point11Val0IdentityValueValid) ||
		!model.RequiredCustomerReview ||
		!model.LegalReviewRequiredForExternalUse ||
		!model.NoPremiumGuarantee ||
		!model.NoRatingClaim ||
		!model.NoComplianceGuarantee ||
		!model.NoFinancialGuarantee ||
		!point12Val0StringListValid(model.AllowedWordingRefs, point11Val0IdentityValueValid) ||
		!point12Val0StringListValid(model.BlockedWordingRefs, point11Val0IdentityValueValid) ||
		!point11Val0IdentityValueValid(model.SupportStatement) {
		return Point12Val0FinancialEvidenceSupportStateBlocked
	}
	values := append([]string{model.SupportStatement}, model.Limitations...)
	if point12Val0ContainsForbiddenClaim(values...) {
		return Point12Val0FinancialEvidenceSupportStateBlocked
	}
	return Point12Val0FinancialEvidenceSupportStateActive
}

func point12Val0AgentLineageState(record Point12Val0AgentLineageRecord) string {
	if !point12Val0AgentLineageBindingMatrixValid(record) ||
		!point12Val0AgentLineageExactBindingValid(record) ||
		!point12Val0AgentLineageRefValid(record.AgentID) ||
		!point12Val0AIEvidenceCandidateTypeValid(record.AgentType) ||
		!point12Val0VersionRefValid(record.ModelOrRuleVersionRef) ||
		!point12Val0PermissionManifestHashValid(record.PermissionManifestHash) ||
		!point12Val0EvidenceRefsValid(record.InputEvidenceRefs) ||
		!point12Val0AuditRefValid(record.AuditID) ||
		!point12Val0RecommendationRefValid(record.RecommendationID) ||
		!point12Val0OptionalStringListValid(record.HumanFeedbackRefs, point12Val0HumanFeedbackRefValid) ||
		!record.LineageInputOnly ||
		record.ClaimsCertification ||
		record.ClaimsSourceOfTruth ||
		record.EmitsPrematurePass {
		return Point12Val0ProvenanceStateBlocked
	}
	return Point12Val0ProvenanceStateActive
}

func point12Val0ToolchainProvenanceMissing(model Point12Val0ToolchainAgentProvenanceProfile) bool {
	return model.CIJobIDRef == "" ||
		model.RunnerImageHash == "" ||
		model.BuildToolVersionRef == "" ||
		model.ScannerVersionRef == "" ||
		model.SBOMGeneratorVersionRef == "" ||
		model.SigningToolVersionRef == "" ||
		model.PolicyEngineBuildHash == "" ||
		model.ExecutionEnvironmentClassRef == ""
}

func point12Val0ToolchainProvenanceInvalid(model Point12Val0ToolchainAgentProvenanceProfile) bool {
	return (model.CIJobIDRef != "" && !point12Val0CIJobRefValid(model.CIJobIDRef)) ||
		(model.RunnerImageHash != "" && !point12Val0HashValid(model.RunnerImageHash)) ||
		(model.BuildToolVersionRef != "" && !point12Val0VersionRefValid(model.BuildToolVersionRef)) ||
		(model.ScannerVersionRef != "" && !point12Val0VersionRefValid(model.ScannerVersionRef)) ||
		(model.SBOMGeneratorVersionRef != "" && !point12Val0VersionRefValid(model.SBOMGeneratorVersionRef)) ||
		(model.SigningToolVersionRef != "" && !point12Val0VersionRefValid(model.SigningToolVersionRef)) ||
		(model.PolicyEngineBuildHash != "" && !point12Val0HashValid(model.PolicyEngineBuildHash)) ||
		(model.ExecutionEnvironmentClassRef != "" && !point12Val0ExecutionEnvironmentClassRefValid(model.ExecutionEnvironmentClassRef))
}

func EvaluatePoint12Val0ProvenanceState(model Point12Val0ToolchainAgentProvenanceProfile) string {
	if model.IntroducesNetworkCallPath {
		return Point12Val0ProvenanceStateBlocked
	}
	for _, lineage := range model.AgentLineages {
		if point12Val0AgentLineageState(lineage) != Point12Val0ProvenanceStateActive {
			return Point12Val0ProvenanceStateBlocked
		}
	}
	if !model.DecisiveToolchainProvenanceRequired {
		if len(model.AgentLineages) == 0 {
			return Point12Val0ProvenanceStateReviewRequired
		}
		return Point12Val0ProvenanceStateActive
	}
	if point12Val0ToolchainProvenanceInvalid(model) {
		return Point12Val0ProvenanceStateBlocked
	}
	if point12Val0ToolchainProvenanceMissing(model) {
		return Point12Val0ProvenanceStateReviewRequired
	}
	if len(model.AgentLineages) == 0 {
		return Point12Val0ProvenanceStateReviewRequired
	}
	return Point12Val0ProvenanceStateActive
}

func EvaluatePoint12Val0NoOverclaimState(model Point12Val0NoOverclaimReview) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point12Val0StringListValid(model.AllowedSafeWording, point11Val0IdentityValueValid) ||
		!point12Val0StringListValid(model.BlockedWording, point11Val0IdentityValueValid) {
		return Point12Val0NoOverclaimStateBlocked
	}
	observedValues := append([]string{}, model.ObservedCustomerFacingTexts...)
	observedValues = append(observedValues, model.ObservedExportFacingTexts...)
	if point12Val0ContainsForbiddenClaim(observedValues...) ||
		point12Val0ContainsForbiddenInternalClaim(model.ObservedDiagnostics...) {
		return Point12Val0NoOverclaimStateBlocked
	}
	return Point12Val0NoOverclaimStateActive
}

func EvaluatePoint12Val0State(model Point12Val0Foundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		len(point12Val0InvalidComponentStateReasons(model)) > 0 ||
		model.DependencyState == Point12Val0DependencyStateBlocked ||
		model.DeterminismContractState == Point12Val0DeterminismContractStateBlocked ||
		model.CompatibilityProfileState == Point12Val0CompatibilityProfileStateBlocked ||
		model.ManifestState == Point12Val0ManifestStateBlocked ||
		model.ReplayAssessmentState == Point12Val0ReplayAssessmentStateBlocked ||
		model.RedactionBoundaryState == Point12Val0RedactionBoundaryStateBlocked ||
		model.FinancialEvidenceSupportState == Point12Val0FinancialEvidenceSupportStateBlocked ||
		model.ProvenanceState == Point12Val0ProvenanceStateBlocked ||
		model.NoOverclaimState == Point12Val0NoOverclaimStateBlocked {
		return Point12Val0StateBlocked
	}
	if model.DependencyState == Point12Val0DependencyStateReviewRequired ||
		model.ProvenanceState == Point12Val0ProvenanceStateReviewRequired {
		return Point12Val0StateReviewRequired
	}
	return Point12Val0StateActive
}

func point12Val0InvalidComponentStateReasons(model Point12Val0Foundation) []string {
	reasons := []string{}
	if !point12Val0ExactOneOf(model.DependencyState, []string{Point12Val0DependencyStateActive, Point12Val0DependencyStateBlocked, Point12Val0DependencyStateReviewRequired}) {
		reasons = append(reasons, "dependency_state_invalid")
	}
	if !point12Val0ExactOneOf(model.DeterminismContractState, []string{Point12Val0DeterminismContractStateActive, Point12Val0DeterminismContractStateBlocked}) {
		reasons = append(reasons, "determinism_contract_state_invalid")
	}
	if !point12Val0ExactOneOf(model.CompatibilityProfileState, []string{Point12Val0CompatibilityProfileStateActive, Point12Val0CompatibilityProfileStateBlocked}) {
		reasons = append(reasons, "compatibility_profile_state_invalid")
	}
	if !point12Val0ExactOneOf(model.ManifestState, []string{Point12Val0ManifestStateActive, Point12Val0ManifestStateBlocked}) {
		reasons = append(reasons, "manifest_state_invalid")
	}
	if !point12Val0ExactOneOf(model.ReplayAssessmentState, []string{Point12Val0ReplayAssessmentStateActive, Point12Val0ReplayAssessmentStateBlocked}) {
		reasons = append(reasons, "replay_assessment_state_invalid")
	}
	if !point12Val0ExactOneOf(model.RedactionBoundaryState, []string{Point12Val0RedactionBoundaryStateActive, Point12Val0RedactionBoundaryStateBlocked}) {
		reasons = append(reasons, "redaction_boundary_state_invalid")
	}
	if !point12Val0ExactOneOf(model.FinancialEvidenceSupportState, []string{Point12Val0FinancialEvidenceSupportStateActive, Point12Val0FinancialEvidenceSupportStateBlocked}) {
		reasons = append(reasons, "financial_evidence_support_state_invalid")
	}
	if !point12Val0ExactOneOf(model.ProvenanceState, []string{Point12Val0ProvenanceStateActive, Point12Val0ProvenanceStateBlocked, Point12Val0ProvenanceStateReviewRequired}) {
		reasons = append(reasons, "provenance_state_invalid")
	}
	if !point12Val0ExactOneOf(model.NoOverclaimState, []string{Point12Val0NoOverclaimStateActive, Point12Val0NoOverclaimStateBlocked}) {
		reasons = append(reasons, "no_overclaim_state_invalid")
	}
	return reasons
}

func point12Val0BlockingReasons(model Point12Val0Foundation) []string {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "aggregate_projection_disclaimer_blocked")
	}
	reasons = append(reasons, point12Val0InvalidComponentStateReasons(model)...)
	if model.DependencyState == Point12Val0DependencyStateBlocked {
		reasons = append(reasons, "point11_vald_dependency_blocked")
	}
	if model.DeterminismContractState == Point12Val0DeterminismContractStateBlocked {
		reasons = append(reasons, "determinism_contract_blocked")
	}
	if model.CompatibilityProfileState == Point12Val0CompatibilityProfileStateBlocked {
		reasons = append(reasons, "compatibility_profile_blocked")
	}
	if model.ManifestState == Point12Val0ManifestStateBlocked {
		reasons = append(reasons, "manifest_blocked")
	}
	if model.ReplayAssessmentState == Point12Val0ReplayAssessmentStateBlocked {
		reasons = append(reasons, "replay_assessment_blocked")
	}
	if model.RedactionBoundaryState == Point12Val0RedactionBoundaryStateBlocked {
		reasons = append(reasons, "redaction_boundary_blocked")
	}
	if model.FinancialEvidenceSupportState == Point12Val0FinancialEvidenceSupportStateBlocked {
		reasons = append(reasons, "financial_evidence_support_blocked")
	}
	if model.ProvenanceState == Point12Val0ProvenanceStateBlocked {
		reasons = append(reasons, "provenance_blocked")
	}
	if model.NoOverclaimState == Point12Val0NoOverclaimStateBlocked {
		reasons = append(reasons, "no_overclaim_blocked")
	}
	return reasons
}

func Point12Val0FoundationModel() Point12Val0Foundation {
	disclaimer := point12Val0ProjectionDisclaimerBaseline
	dependency := point12Val0DependencySnapshotModel()
	return Point12Val0Foundation{
		CurrentState:                  Point12Val0StateActive,
		ProjectionDisclaimer:          disclaimer,
		DependencyState:               Point12Val0DependencyStateActive,
		DeterminismContractState:      Point12Val0DeterminismContractStateActive,
		CompatibilityProfileState:     Point12Val0CompatibilityProfileStateActive,
		ManifestState:                 Point12Val0ManifestStateActive,
		ReplayAssessmentState:         Point12Val0ReplayAssessmentStateActive,
		RedactionBoundaryState:        Point12Val0RedactionBoundaryStateActive,
		FinancialEvidenceSupportState: Point12Val0FinancialEvidenceSupportStateActive,
		ProvenanceState:               Point12Val0ProvenanceStateActive,
		NoOverclaimState:              Point12Val0NoOverclaimStateActive,
		Dependency:                    dependency,
		DeterminismContract: Point12Val0ReplayDeterminismContract{
			ReplayMode:                     point12Val0ReplayModeOriginalContext,
			StableEvidenceIdentityRequired: true,
			StablePolicyIdentityRequired:   true,
			StableEngineIdentityRequired:   true,
			StableSchemaIdentityRequired:   true,
			StableTenantScopeRequired:      true,
			StableArtifactIdentityRequired: true,
			DriftExplanationRequired:       true,
			UnsupportedBehavior:            point12Val0UnsupportedBehaviorBlockedReplay,
		},
		CompatibilityProfile: Point12Val0ProofPackCompatibilityProfile{
			CompatibilityProfileRef:          "compatibility_profile_point12_val0_default_001",
			ReplayMode:                       point12Val0ReplayModeOriginalContext,
			PolicyCompatibility:              point12Val0CompatibilityExactRequired,
			EngineCompatibility:              point12Val0CompatibilityExactRequired,
			SchemaCompatibility:              point12Val0CompatibilityExactRequired,
			EvidenceCompatibility:            point12Val0EvidenceCompatibilityExactHashRequired,
			ToolchainCompatibility:           point12Val0ToolchainCompatibilityRequiredIfDecisive,
			UnsupportedBehavior:              point12Val0UnsupportedBehaviorBlockedReplay,
			DecisionDriftExplanationRequired: true,
		},
		Manifest: Point12Val0SignedProofPackManifest{
			ProofPackID:                   "proof_pack_point12_val0_001",
			DecisionID:                    "decision_point12_val0_replay_001",
			PointID:                       point12Val0PointID,
			WaveID:                        point12Val0WaveID,
			ProofPackState:                Point12Val0ProofPackStateGenerated,
			TenantScope:                   "tenant_scope_point12_alpha",
			ArtifactRef:                   "artifact_point12_target_001",
			ArtifactHash:                  "sha256:1111111111111111111111111111111111111111111111111111111111111111",
			EvidenceRefs:                  []string{"evidence:point12-proof-pack-evidence-001"},
			EvidenceHashRefs:              []string{"evidence_hash_point12_proof_pack_001"},
			PolicyRef:                     dependency.PolicyAuthorityContextRef,
			PolicyVersion:                 "policy_version_point12_val0_v1",
			PolicyHash:                    "sha256:2222222222222222222222222222222222222222222222222222222222222222",
			EngineVersion:                 "engine_version_point12_val0_v1",
			EngineHash:                    "sha256:3333333333333333333333333333333333333333333333333333333333333333",
			SchemaVersion:                 "schema_version_point12_val0_v1",
			SchemaHash:                    "sha256:7777777777777777777777777777777777777777777777777777777777777777",
			ClaimRefs:                     []string{dependency.ClaimAuthorityContextRef},
			GovernanceEventRefs:           []string{dependency.GovernanceAuthorityContextRef},
			UpstreamClosureManifestRef:    dependency.UpstreamClosureManifestRef,
			DependencySnapshotRef:         "dependency_snapshot_point12_val0_001",
			PolicyAuthorityContextRef:     dependency.PolicyAuthorityContextRef,
			ClaimAuthorityContextRef:      dependency.ClaimAuthorityContextRef,
			GovernanceAuthorityContextRef: dependency.GovernanceAuthorityContextRef,
			CompatibilityProfileRef:       "compatibility_profile_point12_val0_default_001",
			GeneratedAt:                   "2026-05-03T10:00:00Z",
			FreshnessWindow:               "freshness_window_48h",
			SigningKeyRef:                 "metadata_signing_key_point12_val0_001",
			SignatureRef:                  "signature_metadata_point12_val0_001",
			RedactionManifestRef:          "redaction_manifest_point12_val0_001",
			ProjectionDisclaimer:          disclaimer,
			RetentionClassRef:             "retention_class_point12_advisory_export",
			ToolchainProvenanceRefs:       []string{"toolchain_provenance_point12_val0_001"},
			AgentLineageRefs:              []string{"agent_lineage_point12_val0_001"},
		},
		ReplayAssessment: Point12Val0ReplayAssessment{
			ReplayAssessmentID:      "replay_assessment_point12_val0_001",
			ProofPackState:          Point12Val0ProofPackStateGenerated,
			ReplayResult:            Point12Val0ReplayResultSameDecision,
			DriftExplanation:        "original_context_replay_matches_decision",
			DeterminismContractRef:  "dependency_snapshot_point12_val0_determinism_001",
			CompatibilityProfileRef: "compatibility_profile_point12_val0_default_001",
			OriginalPolicyRef:       dependency.PolicyAuthorityContextRef,
			ReplayPolicyRef:         dependency.PolicyAuthorityContextRef,
			OriginalPolicyHash:      "sha256:2222222222222222222222222222222222222222222222222222222222222222",
			ReplayPolicyHash:        "sha256:2222222222222222222222222222222222222222222222222222222222222222",
			OriginalEngineHash:      "sha256:3333333333333333333333333333333333333333333333333333333333333333",
			ReplayEngineHash:        "sha256:3333333333333333333333333333333333333333333333333333333333333333",
			OriginalSchemaVersion:   "schema_version_point12_val0_v1",
			ReplaySchemaVersion:     "schema_version_point12_val0_v1",
			EvidenceRefs:            []string{"evidence:point12-proof-pack-evidence-001"},
			ReplayEvidenceRefs:      []string{"evidence:point12-proof-pack-evidence-001"},
			EvidenceHashRefs:        []string{"evidence_hash_point12_proof_pack_001"},
			ReplayEvidenceHashRefs:  []string{"evidence_hash_point12_proof_pack_001"},
			ClaimRefs:               []string{dependency.ClaimAuthorityContextRef},
			ReplayClaimRefs:         []string{dependency.ClaimAuthorityContextRef},
			GovernanceEventRefs:     []string{dependency.GovernanceAuthorityContextRef},
			ReplayGovernanceRefs:    []string{dependency.GovernanceAuthorityContextRef},
			DecisiveEvidencePresent: true,
			ProjectionDisclaimer:    disclaimer,
		},
		RedactionBoundary: Point12Val0RedactionBoundary{
			RedactionManifestRef:           "redaction_manifest_point12_val0_001",
			PostRedactionResult:            Point12Val0ReplayResultSameDecision,
			MinimumSafeClaimAfterRedaction: "bounded claim",
		},
		FinancialEvidenceSupportProfile: Point12Val0FinancialInsuranceEvidenceSupportProfile{
			ProfileType:                       point12Val0ProfileTypeFinancialReview,
			EvidenceSupportCategories:         []string{"evidence_support_customer_review", "evidence_support_audit_review"},
			RiskContextMetadata:               []string{"risk_context_metadata_point12_val0"},
			Limitations:                       []string{"not compliance guarantee", "advisory projection only"},
			RequiredCustomerReview:            true,
			LegalReviewRequiredForExternalUse: true,
			NoPremiumGuarantee:                true,
			NoRatingClaim:                     true,
			NoComplianceGuarantee:             true,
			NoFinancialGuarantee:              true,
			AllowedWordingRefs:                []string{"allowed_wording_point12_val0_evidence_support"},
			BlockedWordingRefs:                []string{"denylist_wording_point12_val0_forbidden"},
			SupportStatement:                  "This proof pack contains evidence that may support customer, auditor, financial, or insurance review.",
		},
		ProvenanceProfile: Point12Val0ToolchainAgentProvenanceProfile{
			DecisiveToolchainProvenanceRequired: true,
			CIJobIDRef:                          "ci_job_point12_val0_001",
			RunnerImageHash:                     "sha256:4444444444444444444444444444444444444444444444444444444444444444",
			BuildToolVersionRef:                 "build_tool_point12_val0_001",
			ScannerVersionRef:                   "scanner_point12_val0_001",
			SBOMGeneratorVersionRef:             "sbom_generator_point12_val0_001",
			SigningToolVersionRef:               "signing_tool_point12_val0_001",
			PolicyEngineBuildHash:               "sha256:5555555555555555555555555555555555555555555555555555555555555555",
			ExecutionEnvironmentClassRef:        "execution_environment_point12_val0_ci",
			AgentLineages: []Point12Val0AgentLineageRecord{
				point12Val0DefaultAgentLineageRecord(),
			},
		},
		NoOverclaimReview: Point12Val0NoOverclaimReview{
			ObservedCustomerFacingTexts: []string{"bounded claim"},
			ObservedExportFacingTexts:   []string{"advisory projection only"},
			ObservedDiagnostics:         []string{"canonical evidence spine remains source of truth"},
			AllowedSafeWording:          point12Val0AllowedClaims(),
			BlockedWording:              point12Val0ForbiddenClaims(),
			ProjectionDisclaimer:        disclaimer,
		},
	}
}

func ComputePoint12Val0Foundation(model Point12Val0Foundation) Point12Val0Foundation {
	model.DependencyState = EvaluatePoint12Val0DependencyState(model.Dependency)
	model.DeterminismContractState = EvaluatePoint12Val0DeterminismContractState(model.DeterminismContract)
	model.CompatibilityProfileState = EvaluatePoint12Val0CompatibilityProfileState(model.CompatibilityProfile)
	model.ManifestState = EvaluatePoint12Val0ManifestState(model.Manifest)
	model.RedactionBoundaryState = EvaluatePoint12Val0RedactionBoundaryState(model.RedactionBoundary)
	model.FinancialEvidenceSupportState = EvaluatePoint12Val0FinancialEvidenceSupportState(model.FinancialEvidenceSupportProfile)
	model.ProvenanceState = EvaluatePoint12Val0ProvenanceState(model.ProvenanceProfile)
	model.NoOverclaimState = EvaluatePoint12Val0NoOverclaimState(model.NoOverclaimReview)
	model.ReplayAssessmentState, _ = point12Val0ReplayAssessmentStateAndReasons(
		model.ReplayAssessment,
		model.DeterminismContract,
		model.CompatibilityProfile,
		model.Manifest,
		model.RedactionBoundary,
	)
	model.CurrentState = EvaluatePoint12Val0State(model)
	model.BlockingReasons = point12Val0BlockingReasons(model)
	model.ReviewPrerequisites = append([]string{}, model.Dependency.ReviewPrerequisites...)
	if model.ProvenanceState == Point12Val0ProvenanceStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, "decisive_toolchain_provenance_review_required")
	}
	return model
}

package formal

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const (
	Point12ValDStateActive         = "point12_vald_interrogatable_proof_chain_financial_evidence_support_active"
	Point12ValDStateBlocked        = "point12_vald_interrogatable_proof_chain_financial_evidence_support_blocked"
	Point12ValDStateReviewRequired = "point12_vald_interrogatable_proof_chain_financial_evidence_support_review_required"

	Point12ValDDependencyStateActive         = "point12_vald_dependency_active"
	Point12ValDDependencyStateBlocked        = "point12_vald_dependency_blocked"
	Point12ValDDependencyStateReviewRequired = "point12_vald_dependency_review_required"

	Point12ValDBindingMatrixStateActive         = "point12_vald_binding_matrix_active"
	Point12ValDBindingMatrixStateBlocked        = "point12_vald_binding_matrix_blocked"
	Point12ValDBindingMatrixStateReviewRequired = "point12_vald_binding_matrix_review_required"

	Point12ValDProofChainStateActive         = "point12_vald_proof_chain_projection_active"
	Point12ValDProofChainStateBlocked        = "point12_vald_proof_chain_projection_blocked"
	Point12ValDProofChainStateReviewRequired = "point12_vald_proof_chain_projection_review_required"

	Point12ValDLineageEdgeStateActive         = "point12_vald_lineage_edge_active"
	Point12ValDLineageEdgeStateBlocked        = "point12_vald_lineage_edge_blocked"
	Point12ValDLineageEdgeStateReviewRequired = "point12_vald_lineage_edge_review_required"

	Point12ValDQueryStateActive         = "point12_vald_query_active"
	Point12ValDQueryStateBlocked        = "point12_vald_query_blocked"
	Point12ValDQueryStateReviewRequired = "point12_vald_query_review_required"

	Point12ValDExplanationStateActive         = "point12_vald_explanation_active"
	Point12ValDExplanationStateBlocked        = "point12_vald_explanation_blocked"
	Point12ValDExplanationStateReviewRequired = "point12_vald_explanation_review_required"

	Point12ValDSupportProfileStateActive         = "point12_vald_support_profile_active"
	Point12ValDSupportProfileStateBlocked        = "point12_vald_support_profile_blocked"
	Point12ValDSupportProfileStateReviewRequired = "point12_vald_support_profile_review_required"

	Point12ValDPortalCompatibilityStateActive  = "point12_vald_portal_compatibility_active"
	Point12ValDPortalCompatibilityStateBlocked = "point12_vald_portal_compatibility_blocked"
)

const (
	point12ValDWaveID                        = "val_d"
	point12ValDPreviousWaveID                = point12ValCWaveID
	point12ValDProjectionDisclaimerBaseline  = "projection_only not_canonical_truth point12_vald_interrogatable_proof_chain_financial_evidence_support"
	point12ValDDependencySnapshotRefBaseline = "dependency_snapshot_point12_vald_valc_computed_001"

	point12ValDBindingClassExactRequired            = "exact_required"
	point12ValDBindingClassCompatibilityAllowed     = "compatibility_allowed_with_evidence"
	point12ValDBindingClassAdvisoryOnly             = "advisory_only"
	point12ValDBindingClassIntentionallyNotBound    = "intentionally_not_bound"
	point12ValDQueryKindWhyDecision                 = "why_decision"
	point12ValDQueryKindWhyChanged                  = "why_changed"
	point12ValDQueryKindExplainMismatch             = "explain_mismatch"
	point12ValDQueryKindExplainMissingEvidence      = "explain_missing_evidence"
	point12ValDQueryKindExplainRedactionLimitations = "explain_redaction_limitations"
	point12ValDQueryKindEvidenceLineage             = "evidence_lineage"
	point12ValDQueryKindFinancialEvidenceSupport    = "financial_evidence_support"
	point12ValDQueryKindInsuranceEvidenceSupport    = "insurance_evidence_support"
	point12ValDQueryKindAuditEvidenceSupport        = "audit_evidence_support"
	point12ValDQueryKindPortalCompatibility         = "portal_compatibility"
	point12ValDLineageEdgeTypeSourceToEvidence      = "source_to_evidence"
	point12ValDLineageEdgeTypeEvidenceToArtifact    = "evidence_to_artifact"
	point12ValDLineageEdgeTypeArtifactToDecision    = "artifact_to_decision"
	point12ValDLineageEdgeTypeDecisionToProofPack   = "decision_to_proof_pack"
	point12ValDLineageEdgeTypeProofPackToManifest   = "proof_pack_to_manifest"
	point12ValDLineageEdgeTypeManifestToReplay      = "manifest_to_replay"
	point12ValDLineageEdgeTypeReplayToExport        = "replay_to_export"
	point12ValDLineageEdgeTypeExportToOfflineBundle = "export_to_offline_bundle"
	point12ValDLineageEdgeTypeRedactionToExport     = "redaction_to_export"
	point12ValDLineageEdgeTypeClaimToDecision       = "claim_to_decision"
	point12ValDLineageEdgeTypeGovernanceToDecision  = "governance_to_decision"
	point12ValDLineageEdgeTypeAgentFindingAdvisory  = "agent_finding_to_lineage_advisory"
	point12ValDProfileTypeAuditReview               = "audit_review"
)

type Point12ValDValCReviewContext struct {
	SnapshotFromComputedOutput   bool     `json:"snapshot_from_computed_output"`
	ValCPrematurePoint12PassSeen bool     `json:"valc_premature_point12_pass_seen"`
	ReviewPrerequisites          []string `json:"review_prerequisites,omitempty"`
}

type Point12ValDDependencySnapshot struct {
	ValCCurrentState               string                               `json:"valc_current_state"`
	ValCDependencyState            string                               `json:"valc_dependency_state"`
	ValCExportState                string                               `json:"valc_export_state"`
	ValCRedactionManifestState     string                               `json:"valc_redaction_manifest_state"`
	ValCRedactionImpactState       string                               `json:"valc_redaction_impact_state"`
	ValCOfflineBundleState         string                               `json:"valc_offline_bundle_state"`
	ValCPublicPrivateBoundaryState string                               `json:"valc_public_private_boundary_state"`
	ValCPointID                    string                               `json:"valc_point_id"`
	ValCWaveID                     string                               `json:"valc_wave_id"`
	ProjectionDisclaimer           string                               `json:"projection_disclaimer"`
	SnapshotRef                    string                               `json:"snapshot_ref"`
	SnapshotFromComputedOutput     bool                                 `json:"snapshot_from_computed_output"`
	ValCExternalAPIUsed            bool                                 `json:"valc_external_api_used"`
	ValCPointPassEmitted           bool                                 `json:"valc_point_pass_emitted"`
	ValCPrematurePoint12PassSeen   bool                                 `json:"valc_premature_point12_pass_seen"`
	ReviewPrerequisites            []string                             `json:"review_prerequisites,omitempty"`
	ValCAuditExportBundle          Point12ValCAuditExportBundle         `json:"valc_export_bundle"`
	ValCRedactionManifest          Point12ValCRedactionManifest         `json:"valc_redaction_manifest"`
	ValCRedactionImpactVerdict     Point12ValCRedactionImpactVerdict    `json:"valc_redaction_impact_verdict"`
	ValCOfflineBundle              Point12ValCOfflineVerificationBundle `json:"valc_offline_bundle"`
	ValCPublicPrivateBoundary      Point12ValCPublicPrivateBoundary     `json:"valc_public_private_boundary"`
	ValBReplayRequest              Point12ValBReplayRequest             `json:"valb_replay_request"`
	ValBReplayResult               Point12ValBReplayResult              `json:"valb_replay_result"`
}

type Point12ValDBindingMatrixField struct {
	FieldName            string `json:"field_name"`
	DownstreamModel      string `json:"downstream_model"`
	UpstreamSource       string `json:"upstream_source"`
	BindingClass         string `json:"binding_class"`
	DownstreamValueRef   string `json:"downstream_value_ref"`
	UpstreamValueRef     string `json:"upstream_value_ref"`
	DownstreamHash       string `json:"downstream_hash"`
	UpstreamHash         string `json:"upstream_hash"`
	DownstreamVersion    string `json:"downstream_version"`
	UpstreamVersion      string `json:"upstream_version"`
	ValidationRequired   bool   `json:"validation_required"`
	MutationTestRequired bool   `json:"mutation_test_required"`
	Reason               string `json:"reason"`
}

type Point12ValDBindingMatrix struct {
	BindingMatrixID       string                          `json:"binding_matrix_id"`
	PointID               string                          `json:"point_id"`
	WaveID                string                          `json:"wave_id"`
	UpstreamDependencyRef string                          `json:"upstream_dependency_ref"`
	BoundFields           []Point12ValDBindingMatrixField `json:"bound_fields,omitempty"`
	BindingLimitations    []string                        `json:"binding_limitations,omitempty"`
	GeneratedAt           string                          `json:"generated_at"`
	MatrixState           string                          `json:"matrix_state"`
}

type Point12ValDLineageEdge struct {
	EdgeID                 string   `json:"edge_id"`
	EdgeType               string   `json:"edge_type"`
	FromRef                string   `json:"from_ref"`
	ToRef                  string   `json:"to_ref"`
	FromHash               string   `json:"from_hash"`
	ToHash                 string   `json:"to_hash"`
	TenantScope            string   `json:"tenant_scope"`
	EvidenceSpineRef       string   `json:"evidence_spine_ref"`
	SourceTimestamp        string   `json:"source_timestamp"`
	TargetTimestamp        string   `json:"target_timestamp"`
	Decisive               bool     `json:"decisive"`
	Inferred               bool     `json:"inferred"`
	AdvisoryOnly           bool     `json:"advisory_only"`
	EdgeState              string   `json:"edge_state"`
	Explanation            string   `json:"explanation"`
	AgentID                string   `json:"agent_id"`
	AgentType              string   `json:"agent_type"`
	PermissionManifestHash string   `json:"permission_manifest_hash"`
	InputEvidenceRefs      []string `json:"input_evidence_refs,omitempty"`
	AuditID                string   `json:"audit_id"`
	RecommendationID       string   `json:"recommendation_id"`
	LineageInputOnly       bool     `json:"lineage_input_only"`
	ClaimsCertification    bool     `json:"claims_certification"`
	ClaimsSourceOfTruth    bool     `json:"claims_source_of_truth"`
	EmitsPrematurePass     bool     `json:"emits_premature_pass"`
}

type Point12ValDProofChainProjection struct {
	ProofChainID             string                   `json:"proof_chain_id"`
	ProofPackID              string                   `json:"proof_pack_id"`
	ManifestID               string                   `json:"manifest_id"`
	ReplayResultID           string                   `json:"replay_result_id"`
	ExportID                 string                   `json:"export_id"`
	OfflineBundleID          string                   `json:"offline_bundle_id"`
	RedactionManifestID      string                   `json:"redaction_manifest_id"`
	TenantScope              string                   `json:"tenant_scope"`
	ArtifactRef              string                   `json:"artifact_ref"`
	ArtifactHash             string                   `json:"artifact_hash"`
	EvidenceRefs             []string                 `json:"evidence_refs,omitempty"`
	EvidenceHashRefs         []string                 `json:"evidence_hash_refs,omitempty"`
	PolicyRef                string                   `json:"policy_ref"`
	PolicyVersion            string                   `json:"policy_version"`
	PolicyHash               string                   `json:"policy_hash"`
	EngineVersion            string                   `json:"engine_version"`
	EngineHash               string                   `json:"engine_hash"`
	SchemaVersion            string                   `json:"schema_version"`
	SchemaHash               string                   `json:"schema_hash"`
	ClaimRefs                []string                 `json:"claim_refs,omitempty"`
	GovernanceEventRefs      []string                 `json:"governance_event_refs,omitempty"`
	CompatibilityProfileRef  string                   `json:"compatibility_profile_ref"`
	ManifestPayloadHash      string                   `json:"manifest_payload_hash"`
	SignatureMetadataRef     string                   `json:"signature_metadata_ref"`
	PublicPrivateBoundaryRef string                   `json:"public_private_boundary_ref"`
	RetentionClassRef        string                   `json:"retention_class_ref"`
	LineageEdges             []Point12ValDLineageEdge `json:"lineage_edges,omitempty"`
	SourceEvidenceSpineRefs  []string                 `json:"source_evidence_spine_refs,omitempty"`
	ProjectionHash           string                   `json:"projection_hash"`
	ProjectionDisclaimer     string                   `json:"projection_disclaimer"`
	AdvisoryOnly             bool                     `json:"advisory_only"`
	ProjectionState          string                   `json:"projection_state"`
}

type Point12ValDProofChainQuery struct {
	QueryID                         string `json:"query_id"`
	QueryKind                       string `json:"query_kind"`
	ProofChainID                    string `json:"proof_chain_id"`
	ProofPackID                     string `json:"proof_pack_id"`
	ManifestID                      string `json:"manifest_id"`
	ReplayResultID                  string `json:"replay_result_id"`
	TenantScope                     string `json:"tenant_scope"`
	ArtifactRef                     string `json:"artifact_ref"`
	RequestedExplanation            string `json:"requested_explanation"`
	RequestedAudience               string `json:"requested_audience"`
	IncludeRedactionLimitations     bool   `json:"include_redaction_limitations"`
	IncludeMismatchDetails          bool   `json:"include_mismatch_details"`
	IncludeFinancialEvidenceSupport bool   `json:"include_financial_evidence_support"`
	IncludePortalCompatibility      bool   `json:"include_portal_compatibility"`
	AllowExternalAPI                bool   `json:"allow_external_api"`
	AllowMutation                   bool   `json:"allow_mutation"`
	GeneratedAt                     string `json:"generated_at"`
	QueryState                      string `json:"query_state"`
}

type Point12ValDExplanationResult struct {
	ExplanationID               string   `json:"explanation_id"`
	QueryID                     string   `json:"query_id"`
	ExplanationKind             string   `json:"explanation_kind"`
	ProofChainID                string   `json:"proof_chain_id"`
	TenantScope                 string   `json:"tenant_scope"`
	BasedOnRefs                 []string `json:"based_on_refs,omitempty"`
	BasedOnHashes               []string `json:"based_on_hashes,omitempty"`
	ExpectedRefs                []string `json:"expected_refs,omitempty"`
	ActualRefs                  []string `json:"actual_refs,omitempty"`
	ExpectedHashes              []string `json:"expected_hashes,omitempty"`
	ActualHashes                []string `json:"actual_hashes,omitempty"`
	ExpectedVersions            []string `json:"expected_versions,omitempty"`
	ActualVersions              []string `json:"actual_versions,omitempty"`
	DecisionContextSummary      string   `json:"decision_context_summary"`
	MismatchExplanations        []string `json:"mismatch_explanations,omitempty"`
	MissingEvidenceExplanations []string `json:"missing_evidence_explanations,omitempty"`
	RedactionLimitations        []string `json:"redaction_limitations,omitempty"`
	WhyDecisionSummary          string   `json:"why_decision_summary"`
	WhyChangedSummary           string   `json:"why_changed_summary"`
	DriftReasons                []string `json:"drift_reasons,omitempty"`
	Limitations                 []string `json:"limitations,omitempty"`
	CustomerVisibleStatement    string   `json:"customer_visible_statement"`
	InternalDiagnosticSummary   string   `json:"internal_diagnostic_summary"`
	AdvisoryOnly                bool     `json:"advisory_only"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
	NoOverclaimState            string   `json:"no_overclaim_state"`
	ExplanationHash             string   `json:"explanation_hash"`
	ExplanationState            string   `json:"explanation_state"`
}

type Point12ValDFinancialInsuranceEvidenceSupportProfile struct {
	ProfileID                         string   `json:"profile_id"`
	ProfileType                       string   `json:"profile_type"`
	ProofChainID                      string   `json:"proof_chain_id"`
	ProofPackID                       string   `json:"proof_pack_id"`
	ExportID                          string   `json:"export_id"`
	TenantScope                       string   `json:"tenant_scope"`
	EvidenceSupportCategories         []string `json:"evidence_support_categories,omitempty"`
	RiskContextMetadata               []string `json:"risk_context_metadata,omitempty"`
	SupportingEvidenceRefs            []string `json:"supporting_evidence_refs,omitempty"`
	SupportingEvidenceHashRefs        []string `json:"supporting_evidence_hash_refs,omitempty"`
	Limitations                       []string `json:"limitations,omitempty"`
	RequiredCustomerReview            bool     `json:"required_customer_review"`
	LegalReviewRequiredForExternalUse bool     `json:"legal_review_required_for_external_use"`
	NoPremiumGuarantee                bool     `json:"no_premium_guarantee"`
	NoRatingClaim                     bool     `json:"no_rating_claim"`
	NoComplianceGuarantee             bool     `json:"no_compliance_guarantee"`
	NoFinancialGuarantee              bool     `json:"no_financial_guarantee"`
	NoLegalProtectionGuarantee        bool     `json:"no_legal_protection_guarantee"`
	AllowedWordingRefs                []string `json:"allowed_wording_refs,omitempty"`
	BlockedWordingRefs                []string `json:"blocked_wording_refs,omitempty"`
	SupportStatement                  string   `json:"support_statement"`
	InternalDiagnosticSummary         string   `json:"internal_diagnostic_summary"`
	AdvisoryOnly                      bool     `json:"advisory_only"`
	ProfileHash                       string   `json:"profile_hash"`
	ProfileState                      string   `json:"profile_state"`
}

type Point12ValDPortalCompatibilityContract struct {
	PortalContractID             string   `json:"portal_contract_id"`
	ProofChainID                 string   `json:"proof_chain_id"`
	ProofPackID                  string   `json:"proof_pack_id"`
	ManifestID                   string   `json:"manifest_id"`
	ReplayResultID               string   `json:"replay_result_id"`
	ExportID                     string   `json:"export_id"`
	TenantScope                  string   `json:"tenant_scope"`
	ReadOnly                     bool     `json:"read_only"`
	NotesAnnotationOnly          bool     `json:"notes_annotation_only"`
	EvidenceMutationAllowed      bool     `json:"evidence_mutation_allowed"`
	DecisionMutationAllowed      bool     `json:"decision_mutation_allowed"`
	CertificationAllowed         bool     `json:"certification_allowed"`
	PointPassAllowed             bool     `json:"point_pass_allowed"`
	RequiredProjectionDisclaimer string   `json:"required_projection_disclaimer"`
	AllowedSurfaces              []string `json:"allowed_surfaces,omitempty"`
	BlockedSurfaces              []string `json:"blocked_surfaces,omitempty"`
	CompatibilityState           string   `json:"compatibility_state"`
}

type Point12ValDFoundation struct {
	CurrentState             string                                              `json:"current_state"`
	BlockingReasons          []string                                            `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites      []string                                            `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer     string                                              `json:"projection_disclaimer"`
	DependencyState          string                                              `json:"dependency_state"`
	BindingMatrixState       string                                              `json:"binding_matrix_state"`
	ProofChainState          string                                              `json:"proof_chain_state"`
	QueryState               string                                              `json:"query_state"`
	ExplanationState         string                                              `json:"explanation_state"`
	SupportProfileState      string                                              `json:"support_profile_state"`
	PortalCompatibilityState string                                              `json:"portal_compatibility_state"`
	Dependency               Point12ValDDependencySnapshot                       `json:"dependency"`
	BindingMatrix            Point12ValDBindingMatrix                            `json:"binding_matrix"`
	ProofChain               Point12ValDProofChainProjection                     `json:"proof_chain"`
	Query                    Point12ValDProofChainQuery                          `json:"query"`
	Explanation              Point12ValDExplanationResult                        `json:"explanation"`
	SupportProfile           Point12ValDFinancialInsuranceEvidenceSupportProfile `json:"support_profile"`
	PortalCompatibility      Point12ValDPortalCompatibilityContract              `json:"portal_compatibility"`
}

type point12ValDExpectedLineageBinding struct {
	EdgeType                     string
	FromRef                      string
	ToRef                        string
	FromHash                     string
	ToHash                       string
	TenantScope                  string
	EvidenceSpineRef             string
	AllowedEvidenceSpineRefs     []string
	MatchFromRef                 bool
	MatchToRef                   bool
	MatchFromHash                bool
	MatchToHash                  bool
	MatchEvidenceSpineRef        bool
	MatchAllowedEvidenceSpineRef bool
	RequireAdvisoryOnly          bool
	RequireDecisive              bool
	ForbidDecisive               bool
}

type point12ValDExpectedLineageBindingGroup struct {
	EdgeType   string
	Expected   []point12ValDExpectedLineageBinding
	BlockExtra bool
}

func point12ValDHash(parts ...string) string {
	sum := sha256.Sum256([]byte(strings.Join(parts, "\n")))
	return "sha256:" + hex.EncodeToString(sum[:])
}

func point12ValDComputedProjectionHash(model Point12ValDProofChainProjection) string {
	parts := []string{
		model.ProofChainID,
		model.ProofPackID,
		model.ManifestID,
		model.ReplayResultID,
		model.ExportID,
		model.OfflineBundleID,
		model.RedactionManifestID,
		model.TenantScope,
		model.ArtifactRef,
		model.ArtifactHash,
		strings.Join(model.EvidenceRefs, ","),
		strings.Join(model.EvidenceHashRefs, ","),
		model.PolicyRef,
		model.PolicyVersion,
		model.PolicyHash,
		model.EngineVersion,
		model.EngineHash,
		model.SchemaVersion,
		model.SchemaHash,
		strings.Join(model.ClaimRefs, ","),
		strings.Join(model.GovernanceEventRefs, ","),
		model.CompatibilityProfileRef,
		model.ManifestPayloadHash,
		model.SignatureMetadataRef,
		model.PublicPrivateBoundaryRef,
		model.RetentionClassRef,
		strings.Join(model.SourceEvidenceSpineRefs, ","),
		model.ProjectionDisclaimer,
	}
	for _, edge := range model.LineageEdges {
		parts = append(parts,
			edge.EdgeID,
			edge.EdgeType,
			edge.FromRef,
			edge.ToRef,
			edge.FromHash,
			edge.ToHash,
			edge.TenantScope,
			edge.EvidenceSpineRef,
			edge.AgentID,
			edge.PermissionManifestHash,
		)
	}
	return point12ValDHash(parts...)
}

func point12ValDComputedExplanationHash(model Point12ValDExplanationResult) string {
	return point12ValDHash(
		model.ExplanationID,
		model.QueryID,
		model.ExplanationKind,
		model.ProofChainID,
		model.TenantScope,
		strings.Join(model.BasedOnRefs, ","),
		strings.Join(model.BasedOnHashes, ","),
		strings.Join(model.ExpectedRefs, ","),
		strings.Join(model.ActualRefs, ","),
		strings.Join(model.ExpectedHashes, ","),
		strings.Join(model.ActualHashes, ","),
		strings.Join(model.ExpectedVersions, ","),
		strings.Join(model.ActualVersions, ","),
		model.DecisionContextSummary,
		strings.Join(model.MismatchExplanations, ","),
		strings.Join(model.MissingEvidenceExplanations, ","),
		strings.Join(model.RedactionLimitations, ","),
		model.WhyDecisionSummary,
		model.WhyChangedSummary,
		strings.Join(model.DriftReasons, ","),
		strings.Join(model.Limitations, ","),
		model.CustomerVisibleStatement,
		model.ProjectionDisclaimer,
	)
}

func point12ValDComputedSupportProfileHash(model Point12ValDFinancialInsuranceEvidenceSupportProfile) string {
	return point12ValDHash(
		model.ProfileID,
		model.ProfileType,
		model.ProofChainID,
		model.ProofPackID,
		model.ExportID,
		model.TenantScope,
		strings.Join(model.EvidenceSupportCategories, ","),
		strings.Join(model.RiskContextMetadata, ","),
		strings.Join(model.SupportingEvidenceRefs, ","),
		strings.Join(model.SupportingEvidenceHashRefs, ","),
		strings.Join(model.Limitations, ","),
		strings.Join(model.AllowedWordingRefs, ","),
		strings.Join(model.BlockedWordingRefs, ","),
		model.SupportStatement,
		model.AdvisoryOnlyString(),
	)
}

func (m Point12ValDFinancialInsuranceEvidenceSupportProfile) AdvisoryOnlyString() string {
	if m.AdvisoryOnly {
		return "true"
	}
	return "false"
}

func point12ValDDependencySnapshotRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"dependency_snapshot_", "valc_snapshot_"})
}

func point12ValDBindingMatrixRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"binding_matrix_"})
}

func point12ValDProofChainRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"proof_chain_"})
}

func point12ValDLineageEdgeRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"lineage_edge_", "proof_lineage_edge_"})
}

func point12ValDQueryRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"query_", "proof_chain_query_"})
}

func point12ValDExplanationRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"explanation_", "proof_chain_explanation_"})
}

func point12ValDSupportProfileRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"support_profile_", "financial_support_profile_", "insurance_support_profile_", "audit_support_profile_"})
}

func point12ValDPortalContractRefValid(value string) bool {
	return point11Val0CanonicalRefWithPrefixes(value, []string{"portal_contract_", "portal_compatibility_"})
}

func point12ValDBindingClassValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValDBindingClassExactRequired,
		point12ValDBindingClassCompatibilityAllowed,
		point12ValDBindingClassAdvisoryOnly,
		point12ValDBindingClassIntentionallyNotBound,
	}, value)
}

func point12ValDPrimaryAgentLineageRecord() Point12Val0AgentLineageRecord {
	return point12Val0DefaultAgentLineageRecord()
}

func point12ValDStringSetSubset(values []string, allowed []string) bool {
	allowedSet := map[string]struct{}{}
	for _, value := range allowed {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			allowedSet[trimmed] = struct{}{}
		}
	}
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			return false
		}
		if _, ok := allowedSet[trimmed]; !ok {
			return false
		}
	}
	return true
}

func point12ValDBoolString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}
func point12ValDLineageEdgeTypeValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValDLineageEdgeTypeSourceToEvidence,
		point12ValDLineageEdgeTypeEvidenceToArtifact,
		point12ValDLineageEdgeTypeArtifactToDecision,
		point12ValDLineageEdgeTypeDecisionToProofPack,
		point12ValDLineageEdgeTypeProofPackToManifest,
		point12ValDLineageEdgeTypeManifestToReplay,
		point12ValDLineageEdgeTypeReplayToExport,
		point12ValDLineageEdgeTypeExportToOfflineBundle,
		point12ValDLineageEdgeTypeRedactionToExport,
		point12ValDLineageEdgeTypeClaimToDecision,
		point12ValDLineageEdgeTypeGovernanceToDecision,
		point12ValDLineageEdgeTypeAgentFindingAdvisory,
	}, value)
}

func point12ValDQueryKindValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12ValDQueryKindWhyDecision,
		point12ValDQueryKindWhyChanged,
		point12ValDQueryKindExplainMismatch,
		point12ValDQueryKindExplainMissingEvidence,
		point12ValDQueryKindExplainRedactionLimitations,
		point12ValDQueryKindEvidenceLineage,
		point12ValDQueryKindFinancialEvidenceSupport,
		point12ValDQueryKindInsuranceEvidenceSupport,
		point12ValDQueryKindAuditEvidenceSupport,
		point12ValDQueryKindPortalCompatibility,
	}, value)
}

func point12ValDProfileTypeValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		point12Val0ProfileTypeFinancialReview,
		point12Val0ProfileTypeInsuranceReview,
		point12ValDProfileTypeAuditReview,
	}, value)
}

func point12ValDRefValueValid(value string) bool {
	return strings.TrimSpace(value) != ""
}

func point12ValDTextListValid(values []string) bool {
	return point12Val0OptionalStringListValid(values, point12ValDRefValueValid)
}

func point12ValDHashBindingValueValid(value string) bool {
	return point12Val0HashValid(value) || point12Val0EvidenceHashRefValid(value)
}

func point12ValDHashBindingListValid(values []string) bool {
	return point12Val0StringListValid(values, point12ValDHashBindingValueValid)
}

func point12ValDOptionalHashBindingListValid(values []string) bool {
	if len(values) == 0 {
		return true
	}
	return point12ValDHashBindingListValid(values)
}

func point12ValDStateValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		Point12ValDStateActive,
		Point12ValDStateBlocked,
		Point12ValDStateReviewRequired,
		Point12ValDBindingMatrixStateActive,
		Point12ValDBindingMatrixStateBlocked,
		Point12ValDBindingMatrixStateReviewRequired,
		Point12ValDProofChainStateActive,
		Point12ValDProofChainStateBlocked,
		Point12ValDProofChainStateReviewRequired,
		Point12ValDLineageEdgeStateActive,
		Point12ValDLineageEdgeStateBlocked,
		Point12ValDLineageEdgeStateReviewRequired,
		Point12ValDQueryStateActive,
		Point12ValDQueryStateBlocked,
		Point12ValDQueryStateReviewRequired,
		Point12ValDExplanationStateActive,
		Point12ValDExplanationStateBlocked,
		Point12ValDExplanationStateReviewRequired,
		Point12ValDSupportProfileStateActive,
		Point12ValDSupportProfileStateBlocked,
		Point12ValDSupportProfileStateReviewRequired,
		Point12ValDPortalCompatibilityStateActive,
		Point12ValDPortalCompatibilityStateBlocked,
	}, value)
}

func point12ValDLineageEdgeStateValid(value string) bool {
	return point11Val0ContainsTrimmed([]string{
		Point12ValDLineageEdgeStateActive,
		Point12ValDLineageEdgeStateBlocked,
		Point12ValDLineageEdgeStateReviewRequired,
	}, value)
}

func point12ValDLineageBindingGroup(edgeType string, expected []point12ValDExpectedLineageBinding) point12ValDExpectedLineageBindingGroup {
	return point12ValDExpectedLineageBindingGroup{
		EdgeType:   edgeType,
		Expected:   expected,
		BlockExtra: true,
	}
}

func point12ValDExpectedEvidenceBindings(dependency Point12ValDDependencySnapshot) []point12ValDExpectedLineageBinding {
	expected := make([]point12ValDExpectedLineageBinding, 0, len(dependency.ValCAuditExportBundle.EvidenceRefs))
	for idx, evidenceRef := range dependency.ValCAuditExportBundle.EvidenceRefs {
		if idx >= len(dependency.ValCAuditExportBundle.EvidenceHashRefs) {
			break
		}
		evidenceHash := dependency.ValCAuditExportBundle.EvidenceHashRefs[idx]
		expected = append(expected, point12ValDExpectedLineageBinding{
			EdgeType:              point12ValDLineageEdgeTypeSourceToEvidence,
			ToRef:                 evidenceRef,
			FromHash:              evidenceHash,
			ToHash:                evidenceHash,
			TenantScope:           dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef:      evidenceRef,
			MatchToRef:            true,
			MatchFromHash:         true,
			MatchToHash:           true,
			MatchEvidenceSpineRef: true,
			RequireAdvisoryOnly:   true,
			ForbidDecisive:        true,
		})
	}
	return expected
}

func point12ValDExpectedEvidenceToArtifactBindings(dependency Point12ValDDependencySnapshot) []point12ValDExpectedLineageBinding {
	expected := make([]point12ValDExpectedLineageBinding, 0, len(dependency.ValCAuditExportBundle.EvidenceRefs))
	for idx, evidenceRef := range dependency.ValCAuditExportBundle.EvidenceRefs {
		if idx >= len(dependency.ValCAuditExportBundle.EvidenceHashRefs) {
			break
		}
		evidenceHash := dependency.ValCAuditExportBundle.EvidenceHashRefs[idx]
		expected = append(expected, point12ValDExpectedLineageBinding{
			EdgeType:              point12ValDLineageEdgeTypeEvidenceToArtifact,
			FromRef:               evidenceRef,
			ToRef:                 dependency.ValCAuditExportBundle.ArtifactRef,
			FromHash:              evidenceHash,
			ToHash:                dependency.ValCAuditExportBundle.ArtifactHash,
			TenantScope:           dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef:      evidenceRef,
			MatchFromRef:          true,
			MatchToRef:            true,
			MatchFromHash:         true,
			MatchToHash:           true,
			MatchEvidenceSpineRef: true,
		})
	}
	return expected
}

func point12ValDExpectedLineageBindingGroups(projection Point12ValDProofChainProjection, dependency Point12ValDDependencySnapshot) []point12ValDExpectedLineageBindingGroup {
	groups := []point12ValDExpectedLineageBindingGroup{
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeSourceToEvidence, point12ValDExpectedEvidenceBindings(dependency)),
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeEvidenceToArtifact, point12ValDExpectedEvidenceToArtifactBindings(dependency)),
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeArtifactToDecision, []point12ValDExpectedLineageBinding{{
			EdgeType:                     point12ValDLineageEdgeTypeArtifactToDecision,
			FromRef:                      dependency.ValCAuditExportBundle.ArtifactRef,
			ToRef:                        dependency.ValBReplayRequest.DecisionID,
			FromHash:                     dependency.ValCAuditExportBundle.ArtifactHash,
			ToHash:                       dependency.ValBReplayRequest.OriginalDecisionHash,
			TenantScope:                  dependency.ValCAuditExportBundle.TenantScope,
			AllowedEvidenceSpineRefs:     append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
			MatchFromRef:                 true,
			MatchToRef:                   true,
			MatchFromHash:                true,
			MatchToHash:                  true,
			MatchAllowedEvidenceSpineRef: true,
			RequireDecisive:              true,
		}}),
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeDecisionToProofPack, []point12ValDExpectedLineageBinding{{
			EdgeType:                     point12ValDLineageEdgeTypeDecisionToProofPack,
			FromRef:                      dependency.ValBReplayRequest.DecisionID,
			ToRef:                        dependency.ValCAuditExportBundle.ProofPackID,
			FromHash:                     dependency.ValBReplayRequest.OriginalDecisionHash,
			ToHash:                       dependency.ValCAuditExportBundle.ManifestPayloadHash,
			TenantScope:                  dependency.ValCAuditExportBundle.TenantScope,
			AllowedEvidenceSpineRefs:     append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
			MatchFromRef:                 true,
			MatchToRef:                   true,
			MatchFromHash:                true,
			MatchToHash:                  true,
			MatchAllowedEvidenceSpineRef: true,
			RequireDecisive:              true,
		}}),
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeProofPackToManifest, []point12ValDExpectedLineageBinding{{
			EdgeType:                     point12ValDLineageEdgeTypeProofPackToManifest,
			FromRef:                      dependency.ValCAuditExportBundle.ProofPackID,
			ToRef:                        dependency.ValCAuditExportBundle.ManifestID,
			FromHash:                     dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:                       dependency.ValCAuditExportBundle.ManifestPayloadHash,
			TenantScope:                  dependency.ValCAuditExportBundle.TenantScope,
			AllowedEvidenceSpineRefs:     append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
			MatchFromRef:                 true,
			MatchToRef:                   true,
			MatchFromHash:                true,
			MatchToHash:                  true,
			MatchAllowedEvidenceSpineRef: true,
			RequireDecisive:              true,
		}}),
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeManifestToReplay, []point12ValDExpectedLineageBinding{{
			EdgeType:                     point12ValDLineageEdgeTypeManifestToReplay,
			FromRef:                      dependency.ValCAuditExportBundle.ManifestID,
			ToRef:                        dependency.ValCAuditExportBundle.ReplayResultID,
			FromHash:                     dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:                       dependency.ValCAuditExportBundle.ManifestPayloadHash,
			TenantScope:                  dependency.ValCAuditExportBundle.TenantScope,
			AllowedEvidenceSpineRefs:     append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
			MatchFromRef:                 true,
			MatchToRef:                   true,
			MatchFromHash:                true,
			MatchToHash:                  true,
			MatchAllowedEvidenceSpineRef: true,
			RequireDecisive:              true,
		}}),
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeReplayToExport, []point12ValDExpectedLineageBinding{{
			EdgeType:                     point12ValDLineageEdgeTypeReplayToExport,
			FromRef:                      dependency.ValCAuditExportBundle.ReplayResultID,
			ToRef:                        dependency.ValCAuditExportBundle.ExportID,
			FromHash:                     dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:                       dependency.ValCAuditExportBundle.ManifestPayloadHash,
			TenantScope:                  dependency.ValCAuditExportBundle.TenantScope,
			AllowedEvidenceSpineRefs:     append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
			MatchFromRef:                 true,
			MatchToRef:                   true,
			MatchFromHash:                true,
			MatchToHash:                  true,
			MatchAllowedEvidenceSpineRef: true,
		}}),
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeExportToOfflineBundle, []point12ValDExpectedLineageBinding{{
			EdgeType:                     point12ValDLineageEdgeTypeExportToOfflineBundle,
			FromRef:                      dependency.ValCAuditExportBundle.ExportID,
			ToRef:                        dependency.ValCOfflineBundle.OfflineBundleID,
			FromHash:                     dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:                       dependency.ValCOfflineBundle.ManifestPayloadHash,
			TenantScope:                  dependency.ValCAuditExportBundle.TenantScope,
			AllowedEvidenceSpineRefs:     append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
			MatchFromRef:                 true,
			MatchToRef:                   true,
			MatchFromHash:                true,
			MatchToHash:                  true,
			MatchAllowedEvidenceSpineRef: true,
		}}),
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeRedactionToExport, []point12ValDExpectedLineageBinding{{
			EdgeType:                     point12ValDLineageEdgeTypeRedactionToExport,
			FromRef:                      dependency.ValCRedactionManifest.RedactionManifestID,
			ToRef:                        dependency.ValCAuditExportBundle.ExportID,
			FromHash:                     dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:                       dependency.ValCAuditExportBundle.ManifestPayloadHash,
			TenantScope:                  dependency.ValCAuditExportBundle.TenantScope,
			AllowedEvidenceSpineRefs:     append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
			MatchFromRef:                 true,
			MatchToRef:                   true,
			MatchFromHash:                true,
			MatchToHash:                  true,
			MatchAllowedEvidenceSpineRef: true,
		}}),
	}
	claimBindings := make([]point12ValDExpectedLineageBinding, 0, len(dependency.ValCAuditExportBundle.ClaimRefs))
	for _, claimRef := range dependency.ValCAuditExportBundle.ClaimRefs {
		claimBindings = append(claimBindings, point12ValDExpectedLineageBinding{
			EdgeType:                     point12ValDLineageEdgeTypeClaimToDecision,
			FromRef:                      claimRef,
			ToRef:                        dependency.ValBReplayRequest.DecisionID,
			FromHash:                     dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:                       dependency.ValBReplayRequest.OriginalDecisionHash,
			TenantScope:                  dependency.ValCAuditExportBundle.TenantScope,
			AllowedEvidenceSpineRefs:     append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
			MatchFromRef:                 true,
			MatchToRef:                   true,
			MatchFromHash:                true,
			MatchToHash:                  true,
			MatchAllowedEvidenceSpineRef: true,
			RequireDecisive:              true,
		})
	}
	governanceBindings := make([]point12ValDExpectedLineageBinding, 0, len(dependency.ValCAuditExportBundle.GovernanceEventRefs))
	for _, governanceRef := range dependency.ValCAuditExportBundle.GovernanceEventRefs {
		governanceBindings = append(governanceBindings, point12ValDExpectedLineageBinding{
			EdgeType:                     point12ValDLineageEdgeTypeGovernanceToDecision,
			FromRef:                      governanceRef,
			ToRef:                        dependency.ValBReplayRequest.DecisionID,
			FromHash:                     dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:                       dependency.ValBReplayRequest.OriginalDecisionHash,
			TenantScope:                  dependency.ValCAuditExportBundle.TenantScope,
			AllowedEvidenceSpineRefs:     append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
			MatchFromRef:                 true,
			MatchToRef:                   true,
			MatchFromHash:                true,
			MatchToHash:                  true,
			MatchAllowedEvidenceSpineRef: true,
			RequireDecisive:              true,
		})
	}
	groups = append(groups,
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeClaimToDecision, claimBindings),
		point12ValDLineageBindingGroup(point12ValDLineageEdgeTypeGovernanceToDecision, governanceBindings),
	)
	_ = projection
	return groups
}

func point12ValDLineageEdgeMatchesExpectedBinding(edge Point12ValDLineageEdge, expected point12ValDExpectedLineageBinding) bool {
	if strings.TrimSpace(edge.EdgeType) != strings.TrimSpace(expected.EdgeType) ||
		strings.TrimSpace(edge.TenantScope) != strings.TrimSpace(expected.TenantScope) {
		return false
	}
	if expected.MatchFromRef && strings.TrimSpace(edge.FromRef) != strings.TrimSpace(expected.FromRef) {
		return false
	}
	if expected.MatchToRef && strings.TrimSpace(edge.ToRef) != strings.TrimSpace(expected.ToRef) {
		return false
	}
	if expected.MatchFromHash && strings.TrimSpace(edge.FromHash) != strings.TrimSpace(expected.FromHash) {
		return false
	}
	if expected.MatchToHash && strings.TrimSpace(edge.ToHash) != strings.TrimSpace(expected.ToHash) {
		return false
	}
	if expected.MatchEvidenceSpineRef && strings.TrimSpace(edge.EvidenceSpineRef) != strings.TrimSpace(expected.EvidenceSpineRef) {
		return false
	}
	if expected.MatchAllowedEvidenceSpineRef && !point11Val0ContainsTrimmed(expected.AllowedEvidenceSpineRefs, edge.EvidenceSpineRef) {
		return false
	}
	if expected.RequireAdvisoryOnly && !edge.AdvisoryOnly {
		return false
	}
	if expected.RequireDecisive && !edge.Decisive {
		return false
	}
	if expected.ForbidDecisive && edge.Decisive {
		return false
	}
	return true
}

func point12ValDRequiredLineageBindingsReasons(projection Point12ValDProofChainProjection, dependency Point12ValDDependencySnapshot) []string {
	reasons := []string{}
	if len(dependency.ValCAuditExportBundle.EvidenceRefs) != len(dependency.ValCAuditExportBundle.EvidenceHashRefs) {
		reasons = append(reasons, "proof_chain_expected_evidence_binding_invalid")
		return reasons
	}
	for _, group := range point12ValDExpectedLineageBindingGroups(projection, dependency) {
		candidateIndexes := []int{}
		for idx, edge := range projection.LineageEdges {
			if strings.TrimSpace(edge.EdgeType) == strings.TrimSpace(group.EdgeType) {
				candidateIndexes = append(candidateIndexes, idx)
			}
		}
		if len(group.Expected) == 0 {
			if len(candidateIndexes) > 0 {
				reasons = append(reasons, "proof_chain_required_lineage_edge_unexpected:"+group.EdgeType)
			}
			continue
		}
		if len(candidateIndexes) == 0 {
			reasons = append(reasons, "proof_chain_required_lineage_edge_missing:"+group.EdgeType)
			continue
		}
		used := map[int]bool{}
		for _, expected := range group.Expected {
			matched := false
			for _, idx := range candidateIndexes {
				if used[idx] {
					continue
				}
				if point12ValDLineageEdgeMatchesExpectedBinding(projection.LineageEdges[idx], expected) {
					used[idx] = true
					matched = true
					break
				}
			}
			if !matched {
				reasons = append(reasons, "proof_chain_required_lineage_edge_binding_mismatch:"+group.EdgeType)
				break
			}
		}
		if group.BlockExtra {
			for _, idx := range candidateIndexes {
				if !used[idx] {
					reasons = append(reasons, "proof_chain_required_lineage_edge_unexpected:"+group.EdgeType)
					break
				}
			}
		}
	}
	return reasons
}

func point12ValDProofChainRequiredEdgeTypes(model Point12ValDProofChainProjection) []string {
	required := []string{
		point12ValDLineageEdgeTypeSourceToEvidence,
		point12ValDLineageEdgeTypeEvidenceToArtifact,
		point12ValDLineageEdgeTypeArtifactToDecision,
		point12ValDLineageEdgeTypeDecisionToProofPack,
		point12ValDLineageEdgeTypeProofPackToManifest,
		point12ValDLineageEdgeTypeManifestToReplay,
		point12ValDLineageEdgeTypeReplayToExport,
		point12ValDLineageEdgeTypeExportToOfflineBundle,
		point12ValDLineageEdgeTypeRedactionToExport,
	}
	if len(model.ClaimRefs) > 0 {
		required = append(required, point12ValDLineageEdgeTypeClaimToDecision)
	}
	if len(model.GovernanceEventRefs) > 0 {
		required = append(required, point12ValDLineageEdgeTypeGovernanceToDecision)
	}
	return required
}

func point12ValDHasEdgeType(edges []Point12ValDLineageEdge, edgeType string) bool {
	for _, edge := range edges {
		if strings.TrimSpace(edge.EdgeType) == strings.TrimSpace(edgeType) {
			return true
		}
	}
	return false
}

func point12ValDExpectedExplanationRefs(projection Point12ValDProofChainProjection) []string {
	refs := []string{projection.ArtifactRef, projection.ProofPackID, projection.ManifestID, projection.ReplayResultID}
	return append(refs, projection.SourceEvidenceSpineRefs...)
}

func point12ValDExpectedExplanationHashes(projection Point12ValDProofChainProjection) []string {
	hashes := []string{projection.ArtifactHash, projection.PolicyHash, projection.EngineHash, projection.SchemaHash, projection.ManifestPayloadHash}
	return append(hashes, projection.EvidenceHashRefs...)
}

func point12ValDDependencyReviewContextModel() Point12ValDValCReviewContext {
	return Point12ValDValCReviewContext{
		SnapshotFromComputedOutput: true,
	}
}

func SnapshotPoint12ValDDependencyFromComputedValC(valC Point12ValCFoundation, review Point12ValDValCReviewContext) Point12ValDDependencySnapshot {
	reviewPrerequisites := append([]string{}, valC.ReviewPrerequisites...)
	reviewPrerequisites = append(reviewPrerequisites, review.ReviewPrerequisites...)
	return Point12ValDDependencySnapshot{
		ValCCurrentState:               valC.CurrentState,
		ValCDependencyState:            valC.DependencyState,
		ValCExportState:                valC.ExportState,
		ValCRedactionManifestState:     valC.RedactionManifestState,
		ValCRedactionImpactState:       valC.RedactionImpactState,
		ValCOfflineBundleState:         valC.OfflineBundleState,
		ValCPublicPrivateBoundaryState: valC.PublicPrivateBoundaryState,
		ValCPointID:                    point12Val0PointID,
		ValCWaveID:                     point12ValCWaveID,
		ProjectionDisclaimer:           valC.ProjectionDisclaimer,
		SnapshotRef:                    point12ValDDependencySnapshotRefBaseline,
		SnapshotFromComputedOutput:     review.SnapshotFromComputedOutput,
		ValCExternalAPIUsed:            valC.OfflineBundle.ExternalAPIUsed,
		ValCPointPassEmitted:           review.ValCPrematurePoint12PassSeen,
		ValCPrematurePoint12PassSeen:   review.ValCPrematurePoint12PassSeen,
		ReviewPrerequisites:            reviewPrerequisites,
		ValCAuditExportBundle:          valC.ExportBundle,
		ValCRedactionManifest:          valC.RedactionManifest,
		ValCRedactionImpactVerdict:     valC.RedactionImpactVerdict,
		ValCOfflineBundle:              valC.OfflineBundle,
		ValCPublicPrivateBoundary:      valC.PublicPrivateBoundary,
		ValBReplayRequest:              valC.Dependency.ValBReplayRequest,
		ValBReplayResult:               valC.Dependency.ValBReplayResult,
	}
}

func point12ValDDependencySnapshotModel() Point12ValDDependencySnapshot {
	valC := ComputePoint12ValCFoundation(Point12ValCFoundationModel())
	return SnapshotPoint12ValDDependencyFromComputedValC(valC, point12ValDDependencyReviewContextModel())
}

func EvaluatePoint12ValDDependencyState(model Point12ValDDependencySnapshot) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!model.SnapshotFromComputedOutput ||
		!point12ValDDependencySnapshotRefValid(model.SnapshotRef) ||
		strings.TrimSpace(model.ValCPointID) != point12Val0PointID ||
		strings.TrimSpace(model.ValCWaveID) != point12ValCWaveID ||
		model.ValCExternalAPIUsed ||
		model.ValCPointPassEmitted ||
		model.ValCPrematurePoint12PassSeen ||
		!point12ValCExportRefValid(model.ValCAuditExportBundle.ExportID) ||
		!point12ValCOfflineBundleRefValid(model.ValCOfflineBundle.OfflineBundleID) ||
		!point12Val0RedactionManifestRefValid(model.ValCRedactionManifest.RedactionManifestID) ||
		!point12ValCRedactionImpactRefValid(model.ValCRedactionImpactVerdict.RedactionImpactID) ||
		!point12ValCBoundaryRefValid(model.ValCPublicPrivateBoundary.BoundaryID) ||
		point12Val0ContainsPrematurePassToken(
			model.ValCAuditExportBundle.ExportID,
			model.ValCRedactionManifest.RedactionManifestID,
			model.ValCRedactionImpactVerdict.RedactionImpactID,
			model.ValCOfflineBundle.OfflineBundleID,
			model.ValCAuditExportBundle.CustomerVisibleSummary,
			model.ValCOfflineBundle.CustomerVisibleExplanation,
		) {
		return Point12ValDDependencyStateBlocked
	}
	if strings.TrimSpace(model.ValCCurrentState) == Point12ValCStateBlocked ||
		strings.TrimSpace(model.ValCDependencyState) == Point12ValCDependencyStateBlocked ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStateBlocked ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStateTampered ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStateTenantMismatch ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStateBoundaryViolation ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStateRetentionMissing ||
		strings.TrimSpace(model.ValCRedactionImpactState) == Point12ValCRedactionImpactExportBlocked ||
		strings.TrimSpace(model.ValCPublicPrivateBoundaryState) == Point12ValCPublicPrivateBoundaryStateBlocked ||
		strings.TrimSpace(model.ValCOfflineBundleState) == Point12ValCOfflineBundleStateBlocked {
		return Point12ValDDependencyStateBlocked
	}
	if strings.TrimSpace(model.ValCCurrentState) == Point12ValCStateReviewRequired ||
		strings.TrimSpace(model.ValCDependencyState) == Point12ValCDependencyStateReviewRequired ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStateProjectionOnly ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStatePartialAdvisory ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStateRedactedLimitations ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStateInsufficient ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStateUnsupported ||
		strings.TrimSpace(model.ValCExportState) == Point12ValCExportStateReviewRequired ||
		strings.TrimSpace(model.ValCRedactionManifestState) == Point12ValCRedactionManifestStateReviewRequired ||
		strings.TrimSpace(model.ValCRedactionImpactState) == Point12ValCRedactionImpactReviewRequired ||
		strings.TrimSpace(model.ValCRedactionImpactState) == Point12ValCRedactionImpactRedactedLimits ||
		strings.TrimSpace(model.ValCRedactionImpactState) == Point12ValCRedactionImpactInsufficient ||
		strings.TrimSpace(model.ValCRedactionImpactState) == Point12ValCRedactionImpactBlockedReplay ||
		strings.TrimSpace(model.ValCRedactionImpactState) == Point12ValCRedactionImpactPartialAdvisory ||
		strings.TrimSpace(model.ValCOfflineBundleState) == Point12ValCOfflineBundleStateReviewRequired ||
		strings.TrimSpace(model.ValCOfflineBundleState) == Point12ValCOfflineBundleStateUnsupported ||
		strings.TrimSpace(model.ValCOfflineBundleState) == Point12ValCOfflineBundleStatePartialAdvisoryOnly ||
		strings.TrimSpace(model.ValCOfflineBundleState) == Point12ValCOfflineBundleStateRedactedLimitations ||
		len(model.ReviewPrerequisites) > 0 {
		return Point12ValDDependencyStateReviewRequired
	}
	if strings.TrimSpace(model.ValCCurrentState) != Point12ValCStateActive ||
		strings.TrimSpace(model.ValCDependencyState) != Point12ValCDependencyStateActive ||
		strings.TrimSpace(model.ValCExportState) != Point12ValCExportStateReady ||
		strings.TrimSpace(model.ValCRedactionManifestState) != Point12ValCRedactionManifestStateActive ||
		strings.TrimSpace(model.ValCRedactionImpactState) != Point12ValCRedactionImpactNoDecisionImpact ||
		strings.TrimSpace(model.ValCOfflineBundleState) != Point12ValCOfflineBundleStateActive ||
		strings.TrimSpace(model.ValCPublicPrivateBoundaryState) != Point12ValCPublicPrivateBoundaryStateActive {
		return Point12ValDDependencyStateBlocked
	}
	return Point12ValDDependencyStateActive
}

func point12ValDLineageEdgeStateAndReasons(model Point12ValDLineageEdge, projection Point12ValDProofChainProjection) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValDLineageEdgeRefValid(model.EdgeID) ||
		!point12ValDLineageEdgeTypeValid(model.EdgeType) ||
		!point11Val0IdentityValueValid(model.FromRef) ||
		!point11Val0IdentityValueValid(model.ToRef) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point11Val0IdentityValueValid(model.EvidenceSpineRef) ||
		!point11Val0ValidTimestamp(model.SourceTimestamp) ||
		!point11Val0ValidTimestamp(model.TargetTimestamp) ||
		!point12ValDLineageEdgeStateValid(model.EdgeState) {
		blockedReasons = append(blockedReasons, "lineage_edge_identity_or_metadata_invalid")
	}
	if strings.TrimSpace(model.TenantScope) != strings.TrimSpace(projection.TenantScope) {
		blockedReasons = append(blockedReasons, "lineage_edge_cross_tenant_mismatch")
	}
	if strings.TrimSpace(model.FromHash) == "" || strings.TrimSpace(model.ToHash) == "" {
		blockedReasons = append(blockedReasons, "lineage_edge_hash_missing")
	}
	if model.Decisive && model.Inferred {
		blockedReasons = append(blockedReasons, "lineage_edge_decisive_inferred_blocked")
	}
	if model.Inferred && !model.Decisive && !model.AdvisoryOnly {
		reviewReasons = append(reviewReasons, "lineage_edge_inferred_non_decisive_must_be_advisory_only")
	}
	if strings.TrimSpace(model.EdgeType) == point12ValDLineageEdgeTypeAgentFindingAdvisory {
		expectedLineage := point12ValDPrimaryAgentLineageRecord()
		if !point12Val0AgentLineageRefValid(model.AgentID) ||
			!point12Val0AIEvidenceCandidateTypeValid(model.AgentType) ||
			!point12Val0HashValid(model.PermissionManifestHash) ||
			!point12Val0EvidenceRefsValid(model.InputEvidenceRefs) ||
			!point11Val0IdentityValueValid(model.AuditID) ||
			!point11Val0IdentityValueValid(model.RecommendationID) {
			blockedReasons = append(blockedReasons, "lineage_edge_agent_advisory_identity_invalid")
		}
		if !model.AdvisoryOnly || !model.LineageInputOnly || model.Decisive || model.ClaimsCertification || model.ClaimsSourceOfTruth || model.EmitsPrematurePass {
			blockedReasons = append(blockedReasons, "lineage_edge_agent_advisory_authority_violation")
		}
		if strings.TrimSpace(expectedLineage.AgentID) != "" {
			if strings.TrimSpace(model.AgentID) != strings.TrimSpace(expectedLineage.AgentID) ||
				strings.TrimSpace(model.PermissionManifestHash) != strings.TrimSpace(expectedLineage.PermissionManifestHash) ||
				!point12Val0ExactStringSetMatch(model.InputEvidenceRefs, expectedLineage.InputEvidenceRefs) ||
				strings.TrimSpace(model.AuditID) != strings.TrimSpace(expectedLineage.AuditID) ||
				strings.TrimSpace(model.RecommendationID) != strings.TrimSpace(expectedLineage.RecommendationID) ||
				model.LineageInputOnly != expectedLineage.LineageInputOnly {
				blockedReasons = append(blockedReasons, "lineage_edge_agent_advisory_binding_mismatch")
			}
		}
		if !point12ValDStringSetSubset(model.InputEvidenceRefs, projection.EvidenceRefs) ||
			!point11Val0ContainsTrimmed(model.InputEvidenceRefs, model.EvidenceSpineRef) ||
			strings.TrimSpace(model.ToRef) != strings.TrimSpace(projection.ArtifactRef) ||
			strings.TrimSpace(model.ToHash) != strings.TrimSpace(projection.ArtifactHash) {
			blockedReasons = append(blockedReasons, "lineage_edge_agent_advisory_projection_binding_mismatch")
		}
		if point12Val0ContainsPrematurePassToken(model.AgentID, model.AuditID, model.RecommendationID, model.Explanation) {
			blockedReasons = append(blockedReasons, "lineage_edge_agent_advisory_premature_point12_pass")
		}
	}
	if len(blockedReasons) > 0 {
		return Point12ValDLineageEdgeStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValDLineageEdgeStateReviewRequired, reviewReasons
	}
	return Point12ValDLineageEdgeStateActive, nil
}

func EvaluatePoint12ValDLineageEdgeState(model Point12ValDLineageEdge, projection Point12ValDProofChainProjection) string {
	state, _ := point12ValDLineageEdgeStateAndReasons(model, projection)
	return state
}

func point12ValDProofChainProjectionStateAndReasons(model Point12ValDProofChainProjection, dependency Point12ValDDependencySnapshot) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValDProofChainRefValid(model.ProofChainID) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point12ValBReplayResultRefValid(model.ReplayResultID) ||
		!point12ValCExportRefValid(model.ExportID) ||
		!point12ValCOfflineBundleRefValid(model.OfflineBundleID) ||
		!point12Val0RedactionManifestRefValid(model.RedactionManifestID) ||
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
		!point12Val0HashValid(model.ManifestPayloadHash) ||
		!point12ValASignatureMetadataRefValid(model.SignatureMetadataRef) ||
		!point12ValCBoundaryRefValid(model.PublicPrivateBoundaryRef) ||
		!point12Val0RetentionClassRefValid(model.RetentionClassRef) ||
		!point12Val0EvidenceRefsValid(model.SourceEvidenceSpineRefs) ||
		!point12Val0HashValid(model.ProjectionHash) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!point11Val0ContainsTrimmed([]string{
			Point12ValDProofChainStateActive,
			Point12ValDProofChainStateBlocked,
			Point12ValDProofChainStateReviewRequired,
		}, model.ProjectionState) {
		blockedReasons = append(blockedReasons, "proof_chain_identity_or_metadata_invalid")
	}
	if !model.AdvisoryOnly {
		blockedReasons = append(blockedReasons, "proof_chain_projection_must_remain_advisory_only")
	}
	if strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(dependency.ValCAuditExportBundle.ProofPackID) ||
		strings.TrimSpace(model.ManifestID) != strings.TrimSpace(dependency.ValCAuditExportBundle.ManifestID) ||
		strings.TrimSpace(model.ReplayResultID) != strings.TrimSpace(dependency.ValCAuditExportBundle.ReplayResultID) ||
		strings.TrimSpace(model.ExportID) != strings.TrimSpace(dependency.ValCAuditExportBundle.ExportID) ||
		strings.TrimSpace(model.OfflineBundleID) != strings.TrimSpace(dependency.ValCOfflineBundle.OfflineBundleID) ||
		strings.TrimSpace(model.RedactionManifestID) != strings.TrimSpace(dependency.ValBReplayRequest.RedactionManifestRef) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.ValCAuditExportBundle.TenantScope) ||
		strings.TrimSpace(model.ArtifactRef) != strings.TrimSpace(dependency.ValCAuditExportBundle.ArtifactRef) ||
		strings.TrimSpace(model.ArtifactHash) != strings.TrimSpace(dependency.ValCAuditExportBundle.ArtifactHash) ||
		!point12Val0ExactStringSetMatch(model.EvidenceRefs, dependency.ValCAuditExportBundle.EvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.EvidenceHashRefs, dependency.ValCAuditExportBundle.EvidenceHashRefs) ||
		strings.TrimSpace(model.PolicyRef) != strings.TrimSpace(dependency.ValCAuditExportBundle.PolicyRef) ||
		strings.TrimSpace(model.PolicyVersion) != strings.TrimSpace(dependency.ValCAuditExportBundle.PolicyVersion) ||
		strings.TrimSpace(model.PolicyHash) != strings.TrimSpace(dependency.ValCAuditExportBundle.PolicyHash) ||
		strings.TrimSpace(model.EngineVersion) != strings.TrimSpace(dependency.ValCAuditExportBundle.EngineVersion) ||
		strings.TrimSpace(model.EngineHash) != strings.TrimSpace(dependency.ValCAuditExportBundle.EngineHash) ||
		strings.TrimSpace(model.SchemaVersion) != strings.TrimSpace(dependency.ValCAuditExportBundle.SchemaVersion) ||
		strings.TrimSpace(model.SchemaHash) != strings.TrimSpace(dependency.ValCAuditExportBundle.SchemaHash) ||
		!point12Val0ExactStringSetMatch(model.ClaimRefs, dependency.ValCAuditExportBundle.ClaimRefs) ||
		!point12Val0ExactStringSetMatch(model.GovernanceEventRefs, dependency.ValCAuditExportBundle.GovernanceEventRefs) ||
		strings.TrimSpace(model.CompatibilityProfileRef) != strings.TrimSpace(dependency.ValCAuditExportBundle.CompatibilityProfileRef) ||
		strings.TrimSpace(model.ManifestPayloadHash) != strings.TrimSpace(dependency.ValCAuditExportBundle.ManifestPayloadHash) ||
		strings.TrimSpace(model.SignatureMetadataRef) != strings.TrimSpace(dependency.ValCAuditExportBundle.SignatureMetadataRef) ||
		strings.TrimSpace(model.PublicPrivateBoundaryRef) != strings.TrimSpace(dependency.ValCPublicPrivateBoundary.BoundaryID) ||
		strings.TrimSpace(model.RetentionClassRef) != strings.TrimSpace(dependency.ValCAuditExportBundle.RetentionClassRef) {
		blockedReasons = append(blockedReasons, "proof_chain_dependency_binding_mismatch")
	}
	if !point12Val0ExactStringSetMatch(model.SourceEvidenceSpineRefs, dependency.ValCAuditExportBundle.EvidenceRefs) {
		blockedReasons = append(blockedReasons, "proof_chain_source_evidence_spine_binding_mismatch")
	}
	if model.ProjectionHash != point12ValDComputedProjectionHash(model) {
		blockedReasons = append(blockedReasons, "proof_chain_projection_hash_mismatch")
	}
	for _, edge := range model.LineageEdges {
		edgeState, edgeReasons := point12ValDLineageEdgeStateAndReasons(edge, model)
		if edgeState == Point12ValDLineageEdgeStateBlocked {
			blockedReasons = append(blockedReasons, edgeReasons...)
		}
		if edgeState == Point12ValDLineageEdgeStateReviewRequired {
			reviewReasons = append(reviewReasons, edgeReasons...)
		}
	}
	blockedReasons = append(blockedReasons, point12ValDRequiredLineageBindingsReasons(model, dependency)...)
	if len(model.SourceEvidenceSpineRefs) == 0 {
		blockedReasons = append(blockedReasons, "proof_chain_source_evidence_spine_missing")
	}
	if point12Val0ContainsPrematurePassToken(model.ProofChainID, model.ExportID, model.ProjectionDisclaimer) {
		blockedReasons = append(blockedReasons, "proof_chain_premature_point12_pass")
	}
	if len(blockedReasons) > 0 {
		return Point12ValDProofChainStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValDProofChainStateReviewRequired, reviewReasons
	}
	return Point12ValDProofChainStateActive, nil
}

func EvaluatePoint12ValDProofChainProjectionState(model Point12ValDProofChainProjection, dependency Point12ValDDependencySnapshot) string {
	state, _ := point12ValDProofChainProjectionStateAndReasons(model, dependency)
	return state
}

func point12ValDProofChainQueryStateAndReasons(model Point12ValDProofChainQuery, proofChain Point12ValDProofChainProjection, dependency Point12ValDDependencySnapshot) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValDQueryRefValid(model.QueryID) ||
		!point12ValDQueryKindValid(model.QueryKind) ||
		!point12ValDProofChainRefValid(model.ProofChainID) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point12ValBReplayResultRefValid(model.ReplayResultID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12Val0ArtifactRefValid(model.ArtifactRef) ||
		!point12ValCExportAudienceValid(model.RequestedAudience) ||
		!point11Val0ValidTimestamp(model.GeneratedAt) ||
		!point11Val0ContainsTrimmed([]string{
			Point12ValDQueryStateActive,
			Point12ValDQueryStateBlocked,
			Point12ValDQueryStateReviewRequired,
		}, model.QueryState) {
		blockedReasons = append(blockedReasons, "proof_chain_query_identity_or_metadata_invalid")
	}
	if model.AllowExternalAPI {
		blockedReasons = append(blockedReasons, "proof_chain_query_external_api_forbidden")
	}
	if model.AllowMutation {
		blockedReasons = append(blockedReasons, "proof_chain_query_mutation_forbidden")
	}
	if point12Val0ContainsPrematurePassToken(model.QueryID, model.RequestedExplanation) {
		blockedReasons = append(blockedReasons, "proof_chain_query_premature_point12_pass")
	}
	if strings.TrimSpace(model.ProofChainID) != strings.TrimSpace(proofChain.ProofChainID) ||
		strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(proofChain.ProofPackID) ||
		strings.TrimSpace(model.ManifestID) != strings.TrimSpace(proofChain.ManifestID) ||
		strings.TrimSpace(model.ReplayResultID) != strings.TrimSpace(proofChain.ReplayResultID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(proofChain.TenantScope) ||
		strings.TrimSpace(model.ArtifactRef) != strings.TrimSpace(proofChain.ArtifactRef) {
		blockedReasons = append(blockedReasons, "proof_chain_query_projection_binding_mismatch")
	}
	switch strings.TrimSpace(model.QueryKind) {
	case point12ValDQueryKindWhyChanged:
		if strings.TrimSpace(dependency.ValBReplayRequest.ReplayMode) != point12Val0ReplayModeComparisonMode {
			reviewReasons = append(reviewReasons, "proof_chain_query_why_changed_requires_comparison_context")
		}
	case point12ValDQueryKindExplainMismatch:
		if !model.IncludeMismatchDetails {
			reviewReasons = append(reviewReasons, "proof_chain_query_mismatch_details_required")
		}
	case point12ValDQueryKindExplainMissingEvidence:
		if strings.TrimSpace(dependency.ValBReplayResult.ReplayResultTaxonomy) != Point12Val0ReplayResultInsufficientEvidence &&
			strings.TrimSpace(dependency.ValBReplayResult.ReplayResultTaxonomy) != Point12Val0ReplayResultBlockedReplay {
			reviewReasons = append(reviewReasons, "proof_chain_query_missing_evidence_requires_missing_evidence_context")
		}
	case point12ValDQueryKindExplainRedactionLimitations:
		if !model.IncludeRedactionLimitations {
			reviewReasons = append(reviewReasons, "proof_chain_query_redaction_limitations_required")
		}
	case point12ValDQueryKindFinancialEvidenceSupport, point12ValDQueryKindInsuranceEvidenceSupport, point12ValDQueryKindAuditEvidenceSupport:
		if !model.IncludeFinancialEvidenceSupport {
			reviewReasons = append(reviewReasons, "proof_chain_query_support_profile_flag_missing")
		}
	case point12ValDQueryKindPortalCompatibility:
		if !model.IncludePortalCompatibility {
			reviewReasons = append(reviewReasons, "proof_chain_query_portal_compatibility_flag_missing")
		}
	}
	if len(blockedReasons) > 0 {
		return Point12ValDQueryStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValDQueryStateReviewRequired, reviewReasons
	}
	return Point12ValDQueryStateActive, nil
}

func EvaluatePoint12ValDProofChainQueryState(model Point12ValDProofChainQuery, proofChain Point12ValDProofChainProjection, dependency Point12ValDDependencySnapshot) string {
	state, _ := point12ValDProofChainQueryStateAndReasons(model, proofChain, dependency)
	return state
}

func point12ValDExplanationStateAndReasons(model Point12ValDExplanationResult, query Point12ValDProofChainQuery, proofChain Point12ValDProofChainProjection, dependency Point12ValDDependencySnapshot) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValDExplanationRefValid(model.ExplanationID) ||
		!point12ValDQueryRefValid(model.QueryID) ||
		!point12ValDQueryKindValid(model.ExplanationKind) ||
		!point12ValDProofChainRefValid(model.ProofChainID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12Val0StringListValid(model.BasedOnRefs, point12ValDRefValueValid) ||
		!point12ValDHashBindingListValid(model.BasedOnHashes) ||
		!point12ValDTextListValid(model.ExpectedRefs) ||
		!point12ValDTextListValid(model.ActualRefs) ||
		!point12ValDOptionalHashBindingListValid(model.ExpectedHashes) ||
		!point12ValDOptionalHashBindingListValid(model.ActualHashes) ||
		!point12Val0OptionalStringListValid(model.ExpectedVersions, point12Val0VersionIdentityValid) ||
		!point12Val0OptionalStringListValid(model.ActualVersions, point12Val0VersionIdentityValid) ||
		!point12ValDTextListValid(model.MismatchExplanations) ||
		!point12ValDTextListValid(model.MissingEvidenceExplanations) ||
		!point12ValDTextListValid(model.RedactionLimitations) ||
		!point12Val0OptionalStringListValid(model.DriftReasons, point12ValBDriftClassificationValid) ||
		!point12ValDTextListValid(model.Limitations) ||
		!point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		strings.TrimSpace(model.NoOverclaimState) != Point12Val0NoOverclaimStateActive ||
		!point12Val0HashValid(model.ExplanationHash) ||
		!point11Val0ContainsTrimmed([]string{
			Point12ValDExplanationStateActive,
			Point12ValDExplanationStateBlocked,
			Point12ValDExplanationStateReviewRequired,
		}, model.ExplanationState) {
		blockedReasons = append(blockedReasons, "explanation_identity_or_metadata_invalid")
	}
	if !model.AdvisoryOnly {
		blockedReasons = append(blockedReasons, "explanation_must_remain_advisory_only")
	}
	if strings.TrimSpace(model.QueryID) != strings.TrimSpace(query.QueryID) ||
		strings.TrimSpace(model.ExplanationKind) != strings.TrimSpace(query.QueryKind) ||
		strings.TrimSpace(model.ProofChainID) != strings.TrimSpace(proofChain.ProofChainID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(proofChain.TenantScope) {
		blockedReasons = append(blockedReasons, "explanation_query_or_projection_binding_mismatch")
	}
	if !point12Val0ExactStringSetMatch(model.BasedOnRefs, point12ValDExpectedExplanationRefs(proofChain)) ||
		!point12Val0ExactStringSetMatch(model.BasedOnHashes, point12ValDExpectedExplanationHashes(proofChain)) {
		blockedReasons = append(blockedReasons, "explanation_evidence_binding_mismatch")
	}
	if model.ExplanationHash != point12ValDComputedExplanationHash(model) {
		blockedReasons = append(blockedReasons, "explanation_hash_mismatch")
	}
	if point12Val0ContainsPrematurePassToken(
		model.ExplanationID,
		model.CustomerVisibleStatement,
		model.WhyDecisionSummary,
		model.WhyChangedSummary,
		model.DecisionContextSummary,
	) {
		blockedReasons = append(blockedReasons, "explanation_premature_point12_pass")
	}
	if point12Val0ContainsForbiddenClaim(
		model.CustomerVisibleStatement,
		model.DecisionContextSummary,
		model.WhyDecisionSummary,
		model.WhyChangedSummary,
	) {
		blockedReasons = append(blockedReasons, "explanation_overclaim_detected")
	}
	switch strings.TrimSpace(query.QueryKind) {
	case point12ValDQueryKindWhyDecision:
		if strings.TrimSpace(model.DecisionContextSummary) == "" || strings.TrimSpace(model.WhyDecisionSummary) == "" {
			reviewReasons = append(reviewReasons, "explanation_why_decision_context_missing")
		}
		if strings.TrimSpace(dependency.ValBReplayResult.ReplayResultTaxonomy) == Point12Val0ReplayResultInsufficientEvidence ||
			strings.TrimSpace(dependency.ValBReplayResult.ReplayResultTaxonomy) == Point12Val0ReplayResultBlockedReplay {
			if len(model.MissingEvidenceExplanations) == 0 {
				reviewReasons = append(reviewReasons, "explanation_missing_decisive_evidence_requires_missing_evidence_explanation")
			}
		}
	case point12ValDQueryKindWhyChanged:
		if len(model.DriftReasons) == 0 || strings.TrimSpace(model.WhyChangedSummary) == "" {
			reviewReasons = append(reviewReasons, "explanation_why_changed_requires_drift_reasons")
		}
	case point12ValDQueryKindExplainMismatch:
		if len(model.MismatchExplanations) == 0 ||
			(len(model.ExpectedRefs) == 0 && len(model.ExpectedHashes) == 0 && len(model.ExpectedVersions) == 0) ||
			(len(model.ActualRefs) == 0 && len(model.ActualHashes) == 0 && len(model.ActualVersions) == 0) {
			reviewReasons = append(reviewReasons, "explanation_mismatch_expected_actual_missing")
		}
	case point12ValDQueryKindExplainMissingEvidence:
		if len(model.MissingEvidenceExplanations) == 0 {
			reviewReasons = append(reviewReasons, "explanation_missing_evidence_details_missing")
		}
	case point12ValDQueryKindExplainRedactionLimitations:
		if len(model.RedactionLimitations) == 0 {
			reviewReasons = append(reviewReasons, "explanation_redaction_limitations_missing")
		}
	}
	if len(blockedReasons) > 0 {
		return Point12ValDExplanationStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValDExplanationStateReviewRequired, reviewReasons
	}
	return Point12ValDExplanationStateActive, nil
}

func EvaluatePoint12ValDExplanationState(model Point12ValDExplanationResult, query Point12ValDProofChainQuery, proofChain Point12ValDProofChainProjection, dependency Point12ValDDependencySnapshot) string {
	state, _ := point12ValDExplanationStateAndReasons(model, query, proofChain, dependency)
	return state
}

func point12ValDSupportProfileStateAndReasons(model Point12ValDFinancialInsuranceEvidenceSupportProfile, proofChain Point12ValDProofChainProjection) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValDSupportProfileRefValid(model.ProfileID) ||
		!point12ValDProfileTypeValid(model.ProfileType) ||
		!point12ValDProofChainRefValid(model.ProofChainID) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValCExportRefValid(model.ExportID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point12Val0StringListValid(model.EvidenceSupportCategories, point11Val0IdentityValueValid) ||
		!point12Val0OptionalStringListValid(model.RiskContextMetadata, point11Val0IdentityValueValid) ||
		!point12Val0EvidenceRefsValid(model.SupportingEvidenceRefs) ||
		!point12Val0StringListValid(model.SupportingEvidenceHashRefs, point12Val0EvidenceHashRefValid) ||
		!point12Val0OptionalStringListValid(model.Limitations, point11Val0IdentityValueValid) ||
		!point12Val0OptionalStringListValid(model.AllowedWordingRefs, point11Val0IdentityValueValid) ||
		!point12Val0OptionalStringListValid(model.BlockedWordingRefs, point11Val0IdentityValueValid) ||
		!point12Val0HashValid(model.ProfileHash) ||
		!point11Val0ContainsTrimmed([]string{
			Point12ValDSupportProfileStateActive,
			Point12ValDSupportProfileStateBlocked,
			Point12ValDSupportProfileStateReviewRequired,
		}, model.ProfileState) {
		blockedReasons = append(blockedReasons, "support_profile_identity_or_metadata_invalid")
	}
	if !model.RequiredCustomerReview ||
		!model.LegalReviewRequiredForExternalUse ||
		!model.NoPremiumGuarantee ||
		!model.NoRatingClaim ||
		!model.NoComplianceGuarantee ||
		!model.NoFinancialGuarantee ||
		!model.NoLegalProtectionGuarantee {
		blockedReasons = append(blockedReasons, "support_profile_required_guard_flags_missing")
	}
	if !model.AdvisoryOnly {
		blockedReasons = append(blockedReasons, "support_profile_must_remain_advisory_only")
	}
	if point12Val0ContainsForbiddenClaim(model.SupportStatement) {
		blockedReasons = append(blockedReasons, "support_profile_overclaim_detected")
	}
	if point12Val0ContainsPrematurePassToken(model.ProfileID, model.SupportStatement) {
		blockedReasons = append(blockedReasons, "support_profile_premature_point12_pass")
	}
	if strings.TrimSpace(model.ProofChainID) != strings.TrimSpace(proofChain.ProofChainID) ||
		strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(proofChain.ProofPackID) ||
		strings.TrimSpace(model.ExportID) != strings.TrimSpace(proofChain.ExportID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(proofChain.TenantScope) ||
		!point12Val0ExactStringSetMatch(model.SupportingEvidenceRefs, proofChain.EvidenceRefs) ||
		!point12Val0ExactStringSetMatch(model.SupportingEvidenceHashRefs, proofChain.EvidenceHashRefs) {
		blockedReasons = append(blockedReasons, "support_profile_projection_binding_mismatch")
	}
	if model.ProfileHash != point12ValDComputedSupportProfileHash(model) {
		blockedReasons = append(blockedReasons, "support_profile_hash_mismatch")
	}
	if len(model.Limitations) == 0 {
		reviewReasons = append(reviewReasons, "support_profile_limitations_missing")
	}
	if len(blockedReasons) > 0 {
		return Point12ValDSupportProfileStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValDSupportProfileStateReviewRequired, reviewReasons
	}
	return Point12ValDSupportProfileStateActive, nil
}

func EvaluatePoint12ValDSupportProfileState(model Point12ValDFinancialInsuranceEvidenceSupportProfile, proofChain Point12ValDProofChainProjection) string {
	state, _ := point12ValDSupportProfileStateAndReasons(model, proofChain)
	return state
}

func point12ValDPortalCompatibilityStateAndReasons(model Point12ValDPortalCompatibilityContract, proofChain Point12ValDProofChainProjection) (string, []string) {
	reasons := []string{}
	if !point12ValDPortalContractRefValid(model.PortalContractID) ||
		!point12ValDProofChainRefValid(model.ProofChainID) ||
		!point12Val0ProofPackRefValid(model.ProofPackID) ||
		!point12ValAManifestRefValid(model.ManifestID) ||
		!point12ValBReplayResultRefValid(model.ReplayResultID) ||
		!point12ValCExportRefValid(model.ExportID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point11Val0ValidProjectionDisclaimer(model.RequiredProjectionDisclaimer) ||
		!point12Val0OptionalStringListValid(model.AllowedSurfaces, point11Val0IdentityValueValid) ||
		!point12Val0OptionalStringListValid(model.BlockedSurfaces, point11Val0IdentityValueValid) ||
		!point11Val0ContainsTrimmed([]string{
			Point12ValDPortalCompatibilityStateActive,
			Point12ValDPortalCompatibilityStateBlocked,
		}, model.CompatibilityState) {
		reasons = append(reasons, "portal_compatibility_identity_or_metadata_invalid")
	}
	if !model.ReadOnly ||
		!model.NotesAnnotationOnly ||
		model.EvidenceMutationAllowed ||
		model.DecisionMutationAllowed ||
		model.CertificationAllowed ||
		model.PointPassAllowed {
		reasons = append(reasons, "portal_compatibility_read_only_contract_violated")
	}
	if point12Val0ContainsForbiddenClaim(model.RequiredProjectionDisclaimer) ||
		point12Val0ContainsPrematurePassToken(model.PortalContractID, model.RequiredProjectionDisclaimer) {
		reasons = append(reasons, "portal_compatibility_projection_disclaimer_invalid")
	}
	if strings.TrimSpace(model.ProofChainID) != strings.TrimSpace(proofChain.ProofChainID) ||
		strings.TrimSpace(model.ProofPackID) != strings.TrimSpace(proofChain.ProofPackID) ||
		strings.TrimSpace(model.ManifestID) != strings.TrimSpace(proofChain.ManifestID) ||
		strings.TrimSpace(model.ReplayResultID) != strings.TrimSpace(proofChain.ReplayResultID) ||
		strings.TrimSpace(model.ExportID) != strings.TrimSpace(proofChain.ExportID) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(proofChain.TenantScope) ||
		strings.TrimSpace(model.RequiredProjectionDisclaimer) != strings.TrimSpace(proofChain.ProjectionDisclaimer) {
		reasons = append(reasons, "portal_compatibility_projection_binding_mismatch")
	}
	if len(reasons) > 0 {
		return Point12ValDPortalCompatibilityStateBlocked, reasons
	}
	return Point12ValDPortalCompatibilityStateActive, nil
}

func EvaluatePoint12ValDPortalCompatibilityState(model Point12ValDPortalCompatibilityContract, proofChain Point12ValDProofChainProjection) string {
	state, _ := point12ValDPortalCompatibilityStateAndReasons(model, proofChain)
	return state
}

func point12ValDBindingMatrixStateAndReasons(model Point12ValDBindingMatrix) (string, []string) {
	blockedReasons := []string{}
	reviewReasons := []string{}
	if !point12ValDBindingMatrixRefValid(model.BindingMatrixID) ||
		strings.TrimSpace(model.PointID) != point12Val0PointID ||
		strings.TrimSpace(model.WaveID) != point12ValDWaveID ||
		!point12ValDDependencySnapshotRefValid(model.UpstreamDependencyRef) ||
		!point12Val0OptionalStringListValid(model.BindingLimitations, point11Val0IdentityValueValid) ||
		!point11Val0ValidTimestamp(model.GeneratedAt) ||
		!point11Val0ContainsTrimmed([]string{
			Point12ValDBindingMatrixStateActive,
			Point12ValDBindingMatrixStateBlocked,
			Point12ValDBindingMatrixStateReviewRequired,
		}, model.MatrixState) {
		blockedReasons = append(blockedReasons, "binding_matrix_identity_or_metadata_invalid")
	}
	if len(model.BoundFields) == 0 {
		blockedReasons = append(blockedReasons, "binding_matrix_fields_missing")
	}
	requiredModels := map[string]bool{
		"Point12ValDDependencySnapshot":                       false,
		"Point12ValDProofChainProjection":                     false,
		"Point12ValDLineageEdge":                              false,
		"Point12ValDProofChainQuery":                          false,
		"Point12ValDExplanationResult":                        false,
		"Point12ValDFinancialInsuranceEvidenceSupportProfile": false,
		"Point12ValDPortalCompatibilityContract":              false,
	}
	for _, entry := range model.BoundFields {
		requiredModels[entry.DownstreamModel] = true
		if strings.TrimSpace(entry.FieldName) == "" ||
			strings.TrimSpace(entry.DownstreamModel) == "" ||
			!point12ValDBindingClassValid(entry.BindingClass) {
			blockedReasons = append(blockedReasons, "binding_matrix_entry_identity_invalid")
			continue
		}
		switch strings.TrimSpace(entry.BindingClass) {
		case point12ValDBindingClassExactRequired:
			if !entry.ValidationRequired || !entry.MutationTestRequired {
				blockedReasons = append(blockedReasons, "binding_matrix_exact_required_validator_missing")
			}
			if strings.TrimSpace(entry.UpstreamSource) == "" {
				blockedReasons = append(blockedReasons, "binding_matrix_exact_required_upstream_source_missing")
			}
			if strings.TrimSpace(entry.UpstreamValueRef) == "" &&
				strings.TrimSpace(entry.UpstreamHash) == "" &&
				strings.TrimSpace(entry.UpstreamVersion) == "" {
				blockedReasons = append(blockedReasons, "binding_matrix_exact_required_upstream_identity_missing")
			}
		case point12ValDBindingClassCompatibilityAllowed, point12ValDBindingClassAdvisoryOnly, point12ValDBindingClassIntentionallyNotBound:
			if strings.TrimSpace(entry.Reason) == "" {
				reviewReasons = append(reviewReasons, "binding_matrix_non_exact_reason_missing")
			}
		}
	}
	for modelName, seen := range requiredModels {
		if !seen {
			reviewReasons = append(reviewReasons, "binding_matrix_downstream_model_missing:"+modelName)
		}
	}
	if len(blockedReasons) > 0 {
		return Point12ValDBindingMatrixStateBlocked, blockedReasons
	}
	if len(reviewReasons) > 0 {
		return Point12ValDBindingMatrixStateReviewRequired, reviewReasons
	}
	return Point12ValDBindingMatrixStateActive, nil
}

func EvaluatePoint12ValDBindingMatrixState(model Point12ValDBindingMatrix) string {
	state, _ := point12ValDBindingMatrixStateAndReasons(model)
	return state
}

func EvaluatePoint12ValDState(model Point12ValDFoundation) string {
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		strings.TrimSpace(model.DependencyState) == Point12ValDDependencyStateBlocked ||
		strings.TrimSpace(model.BindingMatrixState) == Point12ValDBindingMatrixStateBlocked ||
		strings.TrimSpace(model.ProofChainState) == Point12ValDProofChainStateBlocked ||
		strings.TrimSpace(model.QueryState) == Point12ValDQueryStateBlocked ||
		strings.TrimSpace(model.ExplanationState) == Point12ValDExplanationStateBlocked ||
		strings.TrimSpace(model.SupportProfileState) == Point12ValDSupportProfileStateBlocked ||
		strings.TrimSpace(model.PortalCompatibilityState) == Point12ValDPortalCompatibilityStateBlocked {
		return Point12ValDStateBlocked
	}
	if strings.TrimSpace(model.DependencyState) == Point12ValDDependencyStateReviewRequired ||
		strings.TrimSpace(model.BindingMatrixState) == Point12ValDBindingMatrixStateReviewRequired ||
		strings.TrimSpace(model.ProofChainState) == Point12ValDProofChainStateReviewRequired ||
		strings.TrimSpace(model.QueryState) == Point12ValDQueryStateReviewRequired ||
		strings.TrimSpace(model.ExplanationState) == Point12ValDExplanationStateReviewRequired ||
		strings.TrimSpace(model.SupportProfileState) == Point12ValDSupportProfileStateReviewRequired {
		return Point12ValDStateReviewRequired
	}
	return Point12ValDStateActive
}

func point12ValDBlockingReasons(model Point12ValDFoundation) []string {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "aggregate_projection_disclaimer_blocked")
	}
	if model.DependencyState == Point12ValDDependencyStateBlocked {
		reasons = append(reasons, "point12_valc_dependency_blocked")
	}
	if model.BindingMatrixState == Point12ValDBindingMatrixStateBlocked {
		reasons = append(reasons, "binding_matrix_blocked")
	}
	if model.ProofChainState == Point12ValDProofChainStateBlocked {
		reasons = append(reasons, "proof_chain_projection_blocked")
	}
	if model.QueryState == Point12ValDQueryStateBlocked {
		reasons = append(reasons, "proof_chain_query_blocked")
	}
	if model.ExplanationState == Point12ValDExplanationStateBlocked {
		reasons = append(reasons, "explanation_blocked")
	}
	if model.SupportProfileState == Point12ValDSupportProfileStateBlocked {
		reasons = append(reasons, "support_profile_blocked")
	}
	if model.PortalCompatibilityState == Point12ValDPortalCompatibilityStateBlocked {
		reasons = append(reasons, "portal_compatibility_blocked")
	}
	return reasons
}

func Point12ValDFoundationModel() Point12ValDFoundation {
	dependency := point12ValDDependencySnapshotModel()
	dependency.ValCCurrentState = Point12ValCStateActive
	dependency.ValCDependencyState = Point12ValCDependencyStateActive
	dependency.ReviewPrerequisites = nil
	agentLineage := point12ValDPrimaryAgentLineageRecord()
	proofChain := Point12ValDProofChainProjection{
		ProofChainID:             "proof_chain_point12_vald_001",
		ProofPackID:              dependency.ValCAuditExportBundle.ProofPackID,
		ManifestID:               dependency.ValCAuditExportBundle.ManifestID,
		ReplayResultID:           dependency.ValCAuditExportBundle.ReplayResultID,
		ExportID:                 dependency.ValCAuditExportBundle.ExportID,
		OfflineBundleID:          dependency.ValCOfflineBundle.OfflineBundleID,
		RedactionManifestID:      dependency.ValBReplayRequest.RedactionManifestRef,
		TenantScope:              dependency.ValCAuditExportBundle.TenantScope,
		ArtifactRef:              dependency.ValCAuditExportBundle.ArtifactRef,
		ArtifactHash:             dependency.ValCAuditExportBundle.ArtifactHash,
		EvidenceRefs:             append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
		EvidenceHashRefs:         append([]string{}, dependency.ValCAuditExportBundle.EvidenceHashRefs...),
		PolicyRef:                dependency.ValCAuditExportBundle.PolicyRef,
		PolicyVersion:            dependency.ValCAuditExportBundle.PolicyVersion,
		PolicyHash:               dependency.ValCAuditExportBundle.PolicyHash,
		EngineVersion:            dependency.ValCAuditExportBundle.EngineVersion,
		EngineHash:               dependency.ValCAuditExportBundle.EngineHash,
		SchemaVersion:            dependency.ValCAuditExportBundle.SchemaVersion,
		SchemaHash:               dependency.ValCAuditExportBundle.SchemaHash,
		ClaimRefs:                append([]string{}, dependency.ValCAuditExportBundle.ClaimRefs...),
		GovernanceEventRefs:      append([]string{}, dependency.ValCAuditExportBundle.GovernanceEventRefs...),
		CompatibilityProfileRef:  dependency.ValCAuditExportBundle.CompatibilityProfileRef,
		ManifestPayloadHash:      dependency.ValCAuditExportBundle.ManifestPayloadHash,
		SignatureMetadataRef:     dependency.ValCAuditExportBundle.SignatureMetadataRef,
		PublicPrivateBoundaryRef: dependency.ValCPublicPrivateBoundary.BoundaryID,
		RetentionClassRef:        dependency.ValCAuditExportBundle.RetentionClassRef,
		SourceEvidenceSpineRefs:  append([]string{}, dependency.ValCAuditExportBundle.EvidenceRefs...),
		ProjectionDisclaimer:     point12ValDProjectionDisclaimerBaseline,
		AdvisoryOnly:             true,
		ProjectionState:          Point12ValDProofChainStateActive,
	}
	proofChain.LineageEdges = []Point12ValDLineageEdge{
		{
			EdgeID:           "lineage_edge_point12_vald_source_001",
			EdgeType:         point12ValDLineageEdgeTypeSourceToEvidence,
			FromRef:          "source_point12_vald_001",
			ToRef:            dependency.ValCAuditExportBundle.EvidenceRefs[0],
			FromHash:         dependency.ValCAuditExportBundle.EvidenceHashRefs[0],
			ToHash:           dependency.ValCAuditExportBundle.EvidenceHashRefs[0],
			TenantScope:      dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:00Z",
			TargetTimestamp:  "2026-05-04T08:00:01Z",
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "source evidence captured in canonical spine",
		},
		{
			EdgeID:           "lineage_edge_point12_vald_artifact_001",
			EdgeType:         point12ValDLineageEdgeTypeEvidenceToArtifact,
			FromRef:          dependency.ValCAuditExportBundle.EvidenceRefs[0],
			ToRef:            dependency.ValCAuditExportBundle.ArtifactRef,
			FromHash:         dependency.ValCAuditExportBundle.EvidenceHashRefs[0],
			ToHash:           dependency.ValCAuditExportBundle.ArtifactHash,
			TenantScope:      dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:02Z",
			TargetTimestamp:  "2026-05-04T08:00:03Z",
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "evidence bound to artifact hash",
		},
		{
			EdgeID:           "lineage_edge_point12_vald_decision_001",
			EdgeType:         point12ValDLineageEdgeTypeArtifactToDecision,
			FromRef:          dependency.ValCAuditExportBundle.ArtifactRef,
			ToRef:            dependency.ValBReplayRequest.DecisionID,
			FromHash:         dependency.ValCAuditExportBundle.ArtifactHash,
			ToHash:           dependency.ValBReplayRequest.OriginalDecisionHash,
			TenantScope:      dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:04Z",
			TargetTimestamp:  "2026-05-04T08:00:05Z",
			Decisive:         true,
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "artifact participates in decision context",
		},
		{
			EdgeID:           "lineage_edge_point12_vald_proofpack_001",
			EdgeType:         point12ValDLineageEdgeTypeDecisionToProofPack,
			FromRef:          dependency.ValBReplayRequest.DecisionID,
			ToRef:            dependency.ValCAuditExportBundle.ProofPackID,
			FromHash:         dependency.ValBReplayRequest.OriginalDecisionHash,
			ToHash:           dependency.ValCAuditExportBundle.ManifestPayloadHash,
			TenantScope:      dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:06Z",
			TargetTimestamp:  "2026-05-04T08:00:07Z",
			Decisive:         true,
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "decision contributes to proof pack assembly",
		},
		{
			EdgeID:           "lineage_edge_point12_vald_manifest_001",
			EdgeType:         point12ValDLineageEdgeTypeProofPackToManifest,
			FromRef:          dependency.ValCAuditExportBundle.ProofPackID,
			ToRef:            dependency.ValCAuditExportBundle.ManifestID,
			FromHash:         dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:           dependency.ValCAuditExportBundle.ManifestPayloadHash,
			TenantScope:      dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:08Z",
			TargetTimestamp:  "2026-05-04T08:00:09Z",
			Decisive:         true,
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "proof pack points at manifest",
		},
		{
			EdgeID:           "lineage_edge_point12_vald_replay_001",
			EdgeType:         point12ValDLineageEdgeTypeManifestToReplay,
			FromRef:          dependency.ValCAuditExportBundle.ManifestID,
			ToRef:            dependency.ValCAuditExportBundle.ReplayResultID,
			FromHash:         dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:           dependency.ValCAuditExportBundle.ManifestPayloadHash,
			TenantScope:      dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:10Z",
			TargetTimestamp:  "2026-05-04T08:00:11Z",
			Decisive:         true,
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "manifest context participates in replay",
		},
		{
			EdgeID:           "lineage_edge_point12_vald_export_001",
			EdgeType:         point12ValDLineageEdgeTypeReplayToExport,
			FromRef:          dependency.ValCAuditExportBundle.ReplayResultID,
			ToRef:            dependency.ValCAuditExportBundle.ExportID,
			FromHash:         dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:           dependency.ValCAuditExportBundle.ManifestPayloadHash,
			TenantScope:      dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:12Z",
			TargetTimestamp:  "2026-05-04T08:00:13Z",
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "replay output packaged into export metadata",
		},
		{
			EdgeID:           "lineage_edge_point12_vald_offline_001",
			EdgeType:         point12ValDLineageEdgeTypeExportToOfflineBundle,
			FromRef:          dependency.ValCAuditExportBundle.ExportID,
			ToRef:            dependency.ValCOfflineBundle.OfflineBundleID,
			FromHash:         dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:           dependency.ValCOfflineBundle.ManifestPayloadHash,
			TenantScope:      dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:14Z",
			TargetTimestamp:  "2026-05-04T08:00:15Z",
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "export metadata references offline bundle",
		},
		{
			EdgeID:           "lineage_edge_point12_vald_redaction_001",
			EdgeType:         point12ValDLineageEdgeTypeRedactionToExport,
			FromRef:          dependency.ValCRedactionManifest.RedactionManifestID,
			ToRef:            dependency.ValCAuditExportBundle.ExportID,
			FromHash:         dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:           dependency.ValCAuditExportBundle.ManifestPayloadHash,
			TenantScope:      dependency.ValCAuditExportBundle.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:16Z",
			TargetTimestamp:  "2026-05-04T08:00:17Z",
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "redaction context remains attached to export",
		},
	}
	if len(proofChain.ClaimRefs) > 0 {
		proofChain.LineageEdges = append(proofChain.LineageEdges, Point12ValDLineageEdge{
			EdgeID:           "lineage_edge_point12_vald_claim_001",
			EdgeType:         point12ValDLineageEdgeTypeClaimToDecision,
			FromRef:          proofChain.ClaimRefs[0],
			ToRef:            dependency.ValBReplayRequest.DecisionID,
			FromHash:         dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:           dependency.ValBReplayRequest.OriginalDecisionHash,
			TenantScope:      proofChain.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:18Z",
			TargetTimestamp:  "2026-05-04T08:00:19Z",
			Decisive:         true,
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "claim reference remains decisive in decision context",
		})
	}
	if len(proofChain.GovernanceEventRefs) > 0 {
		proofChain.LineageEdges = append(proofChain.LineageEdges, Point12ValDLineageEdge{
			EdgeID:           "lineage_edge_point12_vald_governance_001",
			EdgeType:         point12ValDLineageEdgeTypeGovernanceToDecision,
			FromRef:          proofChain.GovernanceEventRefs[0],
			ToRef:            dependency.ValBReplayRequest.DecisionID,
			FromHash:         dependency.ValCAuditExportBundle.ManifestPayloadHash,
			ToHash:           dependency.ValBReplayRequest.OriginalDecisionHash,
			TenantScope:      proofChain.TenantScope,
			EvidenceSpineRef: dependency.ValCAuditExportBundle.EvidenceRefs[0],
			SourceTimestamp:  "2026-05-04T08:00:20Z",
			TargetTimestamp:  "2026-05-04T08:00:21Z",
			Decisive:         true,
			AdvisoryOnly:     true,
			EdgeState:        Point12ValDLineageEdgeStateActive,
			Explanation:      "governance event remains decisive in decision context",
		})
	}
	if strings.TrimSpace(agentLineage.AgentID) != "" && len(agentLineage.InputEvidenceRefs) > 0 {
		proofChain.LineageEdges = append(proofChain.LineageEdges, Point12ValDLineageEdge{
			EdgeID:                 "lineage_edge_point12_vald_agent_001",
			EdgeType:               point12ValDLineageEdgeTypeAgentFindingAdvisory,
			FromRef:                agentLineage.AgentID,
			ToRef:                  proofChain.ArtifactRef,
			FromHash:               agentLineage.PermissionManifestHash,
			ToHash:                 proofChain.ArtifactHash,
			TenantScope:            proofChain.TenantScope,
			EvidenceSpineRef:       agentLineage.InputEvidenceRefs[0],
			SourceTimestamp:        "2026-05-04T08:00:22Z",
			TargetTimestamp:        "2026-05-04T08:00:23Z",
			AdvisoryOnly:           true,
			EdgeState:              Point12ValDLineageEdgeStateActive,
			Explanation:            "AI evidence candidate lineage remains advisory only and cannot satisfy canonical evidence requirements by itself",
			AgentID:                agentLineage.AgentID,
			AgentType:              agentLineage.AgentType,
			PermissionManifestHash: agentLineage.PermissionManifestHash,
			InputEvidenceRefs:      append([]string{}, agentLineage.InputEvidenceRefs...),
			AuditID:                agentLineage.AuditID,
			RecommendationID:       agentLineage.RecommendationID,
			LineageInputOnly:       agentLineage.LineageInputOnly,
			ClaimsCertification:    agentLineage.ClaimsCertification,
			ClaimsSourceOfTruth:    agentLineage.ClaimsSourceOfTruth,
			EmitsPrematurePass:     agentLineage.EmitsPrematurePass,
		})
	}
	proofChain.ProjectionHash = point12ValDComputedProjectionHash(proofChain)
	agentLineageEdge := Point12ValDLineageEdge{}
	for _, edge := range proofChain.LineageEdges {
		if strings.TrimSpace(edge.EdgeType) == point12ValDLineageEdgeTypeAgentFindingAdvisory {
			agentLineageEdge = edge
			break
		}
	}

	query := Point12ValDProofChainQuery{
		QueryID:                         "proof_chain_query_point12_vald_001",
		QueryKind:                       point12ValDQueryKindWhyDecision,
		ProofChainID:                    proofChain.ProofChainID,
		ProofPackID:                     proofChain.ProofPackID,
		ManifestID:                      proofChain.ManifestID,
		ReplayResultID:                  proofChain.ReplayResultID,
		TenantScope:                     proofChain.TenantScope,
		ArtifactRef:                     proofChain.ArtifactRef,
		RequestedExplanation:            "why decision",
		RequestedAudience:               point12ValCExportAudienceAuditor,
		IncludeRedactionLimitations:     true,
		IncludeMismatchDetails:          true,
		IncludeFinancialEvidenceSupport: true,
		IncludePortalCompatibility:      true,
		AllowExternalAPI:                false,
		AllowMutation:                   false,
		GeneratedAt:                     "2026-05-04T08:10:00Z",
		QueryState:                      Point12ValDQueryStateActive,
	}
	explanation := Point12ValDExplanationResult{
		ExplanationID:             "explanation_point12_vald_001",
		QueryID:                   query.QueryID,
		ExplanationKind:           query.QueryKind,
		ProofChainID:              proofChain.ProofChainID,
		TenantScope:               proofChain.TenantScope,
		BasedOnRefs:               point12ValDExpectedExplanationRefs(proofChain),
		BasedOnHashes:             point12ValDExpectedExplanationHashes(proofChain),
		DecisionContextSummary:    "This decision was derived from artifact, evidence, policy, engine, schema, claim, and governance refs.",
		MismatchExplanations:      []string{"replay matched original context within declared compatibility profile"},
		WhyDecisionSummary:        "This decision was derived from these evidence, policy, engine, schema, tenant, artifact, claim, and governance refs.",
		Limitations:               []string{"advisory projection only"},
		CustomerVisibleStatement:  "This proof chain contains evidence that may support customer, auditor, financial, or insurance review.",
		InternalDiagnosticSummary: "internal diagnostic: removed disallowed claims remain blocked",
		AdvisoryOnly:              true,
		ProjectionDisclaimer:      point12ValDProjectionDisclaimerBaseline,
		NoOverclaimState:          Point12Val0NoOverclaimStateActive,
		ExplanationState:          Point12ValDExplanationStateActive,
	}
	explanation.ExplanationHash = point12ValDComputedExplanationHash(explanation)
	supportProfile := Point12ValDFinancialInsuranceEvidenceSupportProfile{
		ProfileID:                         "support_profile_point12_vald_001",
		ProfileType:                       point12Val0ProfileTypeFinancialReview,
		ProofChainID:                      proofChain.ProofChainID,
		ProofPackID:                       proofChain.ProofPackID,
		ExportID:                          proofChain.ExportID,
		TenantScope:                       proofChain.TenantScope,
		EvidenceSupportCategories:         []string{"bounded_evidence_support"},
		RiskContextMetadata:               []string{"advisory_only"},
		SupportingEvidenceRefs:            append([]string{}, proofChain.EvidenceRefs...),
		SupportingEvidenceHashRefs:        append([]string{}, proofChain.EvidenceHashRefs...),
		Limitations:                       []string{"not compliance guarantee", "advisory projection only"},
		RequiredCustomerReview:            true,
		LegalReviewRequiredForExternalUse: true,
		NoPremiumGuarantee:                true,
		NoRatingClaim:                     true,
		NoComplianceGuarantee:             true,
		NoFinancialGuarantee:              true,
		NoLegalProtectionGuarantee:        true,
		AllowedWordingRefs:                []string{"allowed_wording_ref_point12_vald_001"},
		BlockedWordingRefs:                []string{"production approved", "financial guarantee"},
		SupportStatement:                  "This proof chain contains evidence that may support customer, auditor, financial, or insurance review.",
		InternalDiagnosticSummary:         "internal diagnostic: disallowed wording remains blocked",
		AdvisoryOnly:                      true,
		ProfileState:                      Point12ValDSupportProfileStateActive,
	}
	supportProfile.ProfileHash = point12ValDComputedSupportProfileHash(supportProfile)
	portalCompatibility := Point12ValDPortalCompatibilityContract{
		PortalContractID:             "portal_contract_point12_vald_001",
		ProofChainID:                 proofChain.ProofChainID,
		ProofPackID:                  proofChain.ProofPackID,
		ManifestID:                   proofChain.ManifestID,
		ReplayResultID:               proofChain.ReplayResultID,
		ExportID:                     proofChain.ExportID,
		TenantScope:                  proofChain.TenantScope,
		ReadOnly:                     true,
		NotesAnnotationOnly:          true,
		EvidenceMutationAllowed:      false,
		DecisionMutationAllowed:      false,
		CertificationAllowed:         false,
		PointPassAllowed:             false,
		RequiredProjectionDisclaimer: point12ValDProjectionDisclaimerBaseline,
		AllowedSurfaces:              []string{"read_only_projection", "notes_annotation_only"},
		BlockedSurfaces:              []string{"portal_ui", "auditor_account_workflow", "evidence_mutation", "decision_mutation"},
		CompatibilityState:           Point12ValDPortalCompatibilityStateActive,
	}
	bindingMatrix := Point12ValDBindingMatrix{
		BindingMatrixID:       "binding_matrix_point12_vald_001",
		PointID:               point12Val0PointID,
		WaveID:                point12ValDWaveID,
		UpstreamDependencyRef: dependency.SnapshotRef,
		BindingLimitations:    []string{"projection only"},
		GeneratedAt:           "2026-05-04T08:15:00Z",
		MatrixState:           Point12ValDBindingMatrixStateActive,
		BoundFields: []Point12ValDBindingMatrixField{
			{FieldName: "export_id", DownstreamModel: "Point12ValDDependencySnapshot", UpstreamSource: "valc.export_bundle.export_id", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: dependency.ValCAuditExportBundle.ExportID, UpstreamValueRef: dependency.ValCAuditExportBundle.ExportID, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "redaction_manifest_id", DownstreamModel: "Point12ValDDependencySnapshot", UpstreamSource: "valb.replay_request.redaction_manifest_ref", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: dependency.ValCRedactionManifest.RedactionManifestID, UpstreamValueRef: dependency.ValBReplayRequest.RedactionManifestRef, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "tenant_scope", DownstreamModel: "Point12ValDProofChainProjection", UpstreamSource: "valc.export_bundle.tenant_scope", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: proofChain.TenantScope, UpstreamValueRef: dependency.ValCAuditExportBundle.TenantScope, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "artifact_hash", DownstreamModel: "Point12ValDProofChainProjection", UpstreamSource: "valc.export_bundle.artifact_hash", BindingClass: point12ValDBindingClassExactRequired, DownstreamHash: proofChain.ArtifactHash, UpstreamHash: dependency.ValCAuditExportBundle.ArtifactHash, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "evidence_hash_refs", DownstreamModel: "Point12ValDProofChainProjection", UpstreamSource: "valc.export_bundle.evidence_hash_refs", BindingClass: point12ValDBindingClassExactRequired, DownstreamHash: strings.Join(proofChain.EvidenceHashRefs, ","), UpstreamHash: strings.Join(dependency.ValCAuditExportBundle.EvidenceHashRefs, ","), ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "policy_hash", DownstreamModel: "Point12ValDProofChainProjection", UpstreamSource: "valc.export_bundle.policy_hash", BindingClass: point12ValDBindingClassExactRequired, DownstreamHash: proofChain.PolicyHash, UpstreamHash: dependency.ValCAuditExportBundle.PolicyHash, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "engine_hash", DownstreamModel: "Point12ValDProofChainProjection", UpstreamSource: "valc.export_bundle.engine_hash", BindingClass: point12ValDBindingClassExactRequired, DownstreamHash: proofChain.EngineHash, UpstreamHash: dependency.ValCAuditExportBundle.EngineHash, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "schema_hash", DownstreamModel: "Point12ValDProofChainProjection", UpstreamSource: "valc.export_bundle.schema_hash", BindingClass: point12ValDBindingClassExactRequired, DownstreamHash: proofChain.SchemaHash, UpstreamHash: dependency.ValCAuditExportBundle.SchemaHash, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "manifest_payload_hash", DownstreamModel: "Point12ValDProofChainProjection", UpstreamSource: "valc.export_bundle.manifest_payload_hash", BindingClass: point12ValDBindingClassExactRequired, DownstreamHash: proofChain.ManifestPayloadHash, UpstreamHash: dependency.ValCAuditExportBundle.ManifestPayloadHash, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "proof_chain_id", DownstreamModel: "Point12ValDProofChainQuery", UpstreamSource: "vald.proof_chain.proof_chain_id", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: query.ProofChainID, UpstreamValueRef: proofChain.ProofChainID, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "query_kind", DownstreamModel: "Point12ValDProofChainQuery", UpstreamSource: "user_requested_query", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "query kind is caller-selected and cannot rewrite upstream evidence"},
			{FieldName: "based_on_hashes", DownstreamModel: "Point12ValDExplanationResult", UpstreamSource: "vald.proof_chain.exact_hashes", BindingClass: point12ValDBindingClassExactRequired, DownstreamHash: strings.Join(explanation.BasedOnHashes, ","), UpstreamHash: strings.Join(point12ValDExpectedExplanationHashes(proofChain), ","), ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "supporting_evidence_hash_refs", DownstreamModel: "Point12ValDFinancialInsuranceEvidenceSupportProfile", UpstreamSource: "vald.proof_chain.evidence_hash_refs", BindingClass: point12ValDBindingClassExactRequired, DownstreamHash: strings.Join(supportProfile.SupportingEvidenceHashRefs, ","), UpstreamHash: strings.Join(proofChain.EvidenceHashRefs, ","), ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "profile_type", DownstreamModel: "Point12ValDFinancialInsuranceEvidenceSupportProfile", UpstreamSource: "review_intent", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "support profile type is an advisory review lens, not canonical evidence identity"},
			{FieldName: "proof_pack_id", DownstreamModel: "Point12ValDPortalCompatibilityContract", UpstreamSource: "vald.proof_chain.proof_pack_id", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: portalCompatibility.ProofPackID, UpstreamValueRef: proofChain.ProofPackID, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "manifest_id", DownstreamModel: "Point12ValDPortalCompatibilityContract", UpstreamSource: "vald.proof_chain.manifest_id", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: portalCompatibility.ManifestID, UpstreamValueRef: proofChain.ManifestID, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "replay_result_id", DownstreamModel: "Point12ValDPortalCompatibilityContract", UpstreamSource: "vald.proof_chain.replay_result_id", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: portalCompatibility.ReplayResultID, UpstreamValueRef: proofChain.ReplayResultID, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "export_id", DownstreamModel: "Point12ValDPortalCompatibilityContract", UpstreamSource: "vald.proof_chain.export_id", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: portalCompatibility.ExportID, UpstreamValueRef: proofChain.ExportID, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "lineage_edge_hashes", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "canonical_evidence_spine", BindingClass: point12ValDBindingClassExactRequired, DownstreamHash: proofChain.LineageEdges[0].FromHash, UpstreamHash: dependency.ValCAuditExportBundle.EvidenceHashRefs[0], ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "source_to_evidence_from_ref", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "source_system_origin_unmodeled", BindingClass: point12ValDBindingClassIntentionallyNotBound, DownstreamValueRef: proofChain.LineageEdges[0].FromRef, Reason: "source_to_evidence FromRef origin is not modeled upstream; the edge stays advisory-only and exact binding applies to evidence refs, hashes, spine refs, and tenant scope"},
			{FieldName: "ai_agent_id", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "val0.provenance_profile.agent_lineages[0].agent_id", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: agentLineageEdge.AgentID, UpstreamValueRef: agentLineage.AgentID, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "ai_agent_type", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "val0.provenance_profile.agent_lineages[0].agent_type", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: agentLineageEdge.AgentType, UpstreamValueRef: agentLineage.AgentType, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "ai_model_or_rule_version_ref", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "val0.provenance_profile.agent_lineages[0].model_or_rule_version_ref", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "Val D advisory lineage edge does not carry model or rule version and exact validation remains anchored in Val 0 provenance"},
			{FieldName: "ai_permission_manifest_hash", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "val0.provenance_profile.agent_lineages[0].permission_manifest_hash", BindingClass: point12ValDBindingClassExactRequired, DownstreamHash: agentLineageEdge.PermissionManifestHash, UpstreamHash: agentLineage.PermissionManifestHash, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "ai_input_evidence_refs", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "val0.provenance_profile.agent_lineages[0].input_evidence_refs", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: strings.Join(agentLineageEdge.InputEvidenceRefs, ","), UpstreamValueRef: strings.Join(agentLineage.InputEvidenceRefs, ","), ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "ai_input_evidence_hash_refs", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "proof_chain.evidence_hash_refs", BindingClass: point12ValDBindingClassCompatibilityAllowed, Reason: "agent lineage edge carries input evidence refs only while exact hash binding remains enforced on the canonical proof-chain projection"},
			{FieldName: "ai_tenant_scope", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "valc.export_bundle.tenant_scope", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: agentLineageEdge.TenantScope, UpstreamValueRef: proofChain.TenantScope, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "ai_policy_version", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "proof_chain.policy_version", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "policy version remains exact-bound on the proof-chain projection and is not duplicated on the advisory AI lineage edge"},
			{FieldName: "ai_engine_version", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "proof_chain.engine_version", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "engine version remains exact-bound on the proof-chain projection and is not duplicated on the advisory AI lineage edge"},
			{FieldName: "ai_schema_version", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "proof_chain.schema_version", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "schema version remains exact-bound on the proof-chain projection and is not duplicated on the advisory AI lineage edge"},
			{FieldName: "ai_audit_id", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "val0.provenance_profile.agent_lineages[0].audit_id", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: agentLineageEdge.AuditID, UpstreamValueRef: agentLineage.AuditID, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "ai_recommendation_id", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "val0.provenance_profile.agent_lineages[0].recommendation_id", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: agentLineageEdge.RecommendationID, UpstreamValueRef: agentLineage.RecommendationID, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "ai_lineage_input_only", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "val0.provenance_profile.agent_lineages[0].lineage_input_only", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: point12ValDBoolString(agentLineageEdge.LineageInputOnly), UpstreamValueRef: point12ValDBoolString(agentLineage.LineageInputOnly), ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "ai_advisory_only", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "bounded_ai_lineage_contract", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: point12ValDBoolString(agentLineageEdge.AdvisoryOnly), UpstreamValueRef: "true", ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "ai_human_feedback_refs", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "val0.provenance_profile.agent_lineages[0].human_feedback_refs", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "human feedback refs are not projected onto the Val D advisory lineage edge and remain upstream provenance detail"},
			{FieldName: "ai_external_api_allowed", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "bounded_query_and_portal_contracts", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "external API authority is not represented on the advisory lineage edge and remains fail-closed in query, portal, and higher-layer gates"},
			{FieldName: "ai_production_mutation_allowed", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "bounded_query_and_portal_contracts", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "production mutation authority is not represented on the advisory lineage edge and remains blocked by higher-layer governance gates"},
			{FieldName: "ai_canonical_mutation_allowed", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "bounded_query_and_portal_contracts", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "canonical mutation authority is not represented on the advisory lineage edge and remains blocked by higher-layer governance gates"},
			{FieldName: "ai_pass_allowed", DownstreamModel: "Point12ValDLineageEdge", UpstreamSource: "bounded_final_pass_gate", BindingClass: point12ValDBindingClassIntentionallyNotBound, Reason: "advisory lineage edges cannot create pass authority and premature pass emission remains explicitly blocked"},
			{FieldName: "projection_disclaimer", DownstreamModel: "Point12ValDPortalCompatibilityContract", UpstreamSource: "vald.proof_chain.projection_disclaimer", BindingClass: point12ValDBindingClassExactRequired, DownstreamValueRef: portalCompatibility.RequiredProjectionDisclaimer, UpstreamValueRef: proofChain.ProjectionDisclaimer, ValidationRequired: true, MutationTestRequired: true},
			{FieldName: "limitations", DownstreamModel: "Point12ValDExplanationResult", UpstreamSource: "advisory_context", BindingClass: point12ValDBindingClassAdvisoryOnly, Reason: "limitations inform bounded interpretation but do not become source-of-truth identity"},
		},
	}
	return Point12ValDFoundation{
		CurrentState:             Point12ValDStateActive,
		ProjectionDisclaimer:     point12ValDProjectionDisclaimerBaseline,
		DependencyState:          Point12ValDDependencyStateActive,
		BindingMatrixState:       Point12ValDBindingMatrixStateActive,
		ProofChainState:          Point12ValDProofChainStateActive,
		QueryState:               Point12ValDQueryStateActive,
		ExplanationState:         Point12ValDExplanationStateActive,
		SupportProfileState:      Point12ValDSupportProfileStateActive,
		PortalCompatibilityState: Point12ValDPortalCompatibilityStateActive,
		Dependency:               dependency,
		BindingMatrix:            bindingMatrix,
		ProofChain:               proofChain,
		Query:                    query,
		Explanation:              explanation,
		SupportProfile:           supportProfile,
		PortalCompatibility:      portalCompatibility,
	}
}

func ComputePoint12ValDFoundation(model Point12ValDFoundation) Point12ValDFoundation {
	model.DependencyState = EvaluatePoint12ValDDependencyState(model.Dependency)
	matrixState, matrixReasons := point12ValDBindingMatrixStateAndReasons(model.BindingMatrix)
	model.BindingMatrixState = matrixState
	model.BindingMatrix.MatrixState = matrixState
	proofChainState, proofChainReasons := point12ValDProofChainProjectionStateAndReasons(model.ProofChain, model.Dependency)
	model.ProofChainState = proofChainState
	model.ProofChain.ProjectionState = proofChainState
	queryState, queryReasons := point12ValDProofChainQueryStateAndReasons(model.Query, model.ProofChain, model.Dependency)
	model.QueryState = queryState
	model.Query.QueryState = queryState
	explanationState, explanationReasons := point12ValDExplanationStateAndReasons(model.Explanation, model.Query, model.ProofChain, model.Dependency)
	model.ExplanationState = explanationState
	model.Explanation.ExplanationState = explanationState
	supportState, supportReasons := point12ValDSupportProfileStateAndReasons(model.SupportProfile, model.ProofChain)
	model.SupportProfileState = supportState
	model.SupportProfile.ProfileState = supportState
	portalState, portalReasons := point12ValDPortalCompatibilityStateAndReasons(model.PortalCompatibility, model.ProofChain)
	model.PortalCompatibilityState = portalState
	model.PortalCompatibility.CompatibilityState = portalState
	model.CurrentState = EvaluatePoint12ValDState(model)
	model.BlockingReasons = point12ValDBlockingReasons(model)
	model.ReviewPrerequisites = append([]string{}, model.Dependency.ReviewPrerequisites...)
	if model.BindingMatrixState == Point12ValDBindingMatrixStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, matrixReasons...)
	}
	if model.ProofChainState == Point12ValDProofChainStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, proofChainReasons...)
	}
	if model.QueryState == Point12ValDQueryStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, queryReasons...)
	}
	if model.ExplanationState == Point12ValDExplanationStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, explanationReasons...)
	}
	if model.SupportProfileState == Point12ValDSupportProfileStateReviewRequired {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, supportReasons...)
	}
	if model.PortalCompatibilityState == Point12ValDPortalCompatibilityStateBlocked {
		model.BlockingReasons = append(model.BlockingReasons, portalReasons...)
	}
	return model
}

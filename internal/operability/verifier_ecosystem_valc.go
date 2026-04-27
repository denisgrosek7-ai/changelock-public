package operability

import "strings"

const (
	VerifierEcosystemValCAudienceSurfaceStateActive     = "verifier_ecosystem_valc_audience_surfaces_active"
	VerifierEcosystemValCAudienceSurfaceStatePartial    = "verifier_ecosystem_valc_audience_surfaces_partial"
	VerifierEcosystemValCAudienceSurfaceStateIncomplete = "verifier_ecosystem_valc_audience_surfaces_incomplete"
	VerifierEcosystemValCAudienceSurfaceStateBlocked    = "verifier_ecosystem_valc_audience_surfaces_blocked"
	VerifierEcosystemValCAudienceSurfaceStateUnknown    = "verifier_ecosystem_valc_audience_surfaces_unknown"

	VerifierEcosystemValCPublicOutputStateActive     = "verifier_ecosystem_valc_public_output_active"
	VerifierEcosystemValCPublicOutputStatePartial    = "verifier_ecosystem_valc_public_output_partial"
	VerifierEcosystemValCPublicOutputStateIncomplete = "verifier_ecosystem_valc_public_output_incomplete"
	VerifierEcosystemValCPublicOutputStateBlocked    = "verifier_ecosystem_valc_public_output_blocked"
	VerifierEcosystemValCPublicOutputStateUnknown    = "verifier_ecosystem_valc_public_output_unknown"

	VerifierEcosystemValCPartnerOutputStateActive     = "verifier_ecosystem_valc_partner_output_active"
	VerifierEcosystemValCPartnerOutputStatePartial    = "verifier_ecosystem_valc_partner_output_partial"
	VerifierEcosystemValCPartnerOutputStateIncomplete = "verifier_ecosystem_valc_partner_output_incomplete"
	VerifierEcosystemValCPartnerOutputStateBlocked    = "verifier_ecosystem_valc_partner_output_blocked"
	VerifierEcosystemValCPartnerOutputStateUnknown    = "verifier_ecosystem_valc_partner_output_unknown"

	VerifierEcosystemValCAuditorFlowStateActive     = "verifier_ecosystem_valc_auditor_flow_active"
	VerifierEcosystemValCAuditorFlowStatePartial    = "verifier_ecosystem_valc_auditor_flow_partial"
	VerifierEcosystemValCAuditorFlowStateIncomplete = "verifier_ecosystem_valc_auditor_flow_incomplete"
	VerifierEcosystemValCAuditorFlowStateBlocked    = "verifier_ecosystem_valc_auditor_flow_blocked"
	VerifierEcosystemValCAuditorFlowStateUnknown    = "verifier_ecosystem_valc_auditor_flow_unknown"

	VerifierEcosystemValCRequestContractStateActive     = "verifier_ecosystem_valc_request_contract_active"
	VerifierEcosystemValCRequestContractStatePartial    = "verifier_ecosystem_valc_request_contract_partial"
	VerifierEcosystemValCRequestContractStateIncomplete = "verifier_ecosystem_valc_request_contract_incomplete"
	VerifierEcosystemValCRequestContractStateBlocked    = "verifier_ecosystem_valc_request_contract_blocked"
	VerifierEcosystemValCRequestContractStateUnknown    = "verifier_ecosystem_valc_request_contract_unknown"

	VerifierEcosystemValCPublisherProfileStateActive     = "verifier_ecosystem_valc_publisher_profile_active"
	VerifierEcosystemValCPublisherProfileStatePartial    = "verifier_ecosystem_valc_publisher_profile_partial"
	VerifierEcosystemValCPublisherProfileStateIncomplete = "verifier_ecosystem_valc_publisher_profile_incomplete"
	VerifierEcosystemValCPublisherProfileStateBlocked    = "verifier_ecosystem_valc_publisher_profile_blocked"
	VerifierEcosystemValCPublisherProfileStateUnknown    = "verifier_ecosystem_valc_publisher_profile_unknown"

	VerifierEcosystemValCArtifactRuleStateActive     = "verifier_ecosystem_valc_artifact_rules_active"
	VerifierEcosystemValCArtifactRuleStatePartial    = "verifier_ecosystem_valc_artifact_rules_partial"
	VerifierEcosystemValCArtifactRuleStateIncomplete = "verifier_ecosystem_valc_artifact_rules_incomplete"
	VerifierEcosystemValCArtifactRuleStateBlocked    = "verifier_ecosystem_valc_artifact_rules_blocked"
	VerifierEcosystemValCArtifactRuleStateUnknown    = "verifier_ecosystem_valc_artifact_rules_unknown"

	VerifierEcosystemValCTrustDistributionStateActive     = "verifier_ecosystem_valc_trust_distribution_active"
	VerifierEcosystemValCTrustDistributionStatePartial    = "verifier_ecosystem_valc_trust_distribution_partial"
	VerifierEcosystemValCTrustDistributionStateIncomplete = "verifier_ecosystem_valc_trust_distribution_incomplete"
	VerifierEcosystemValCTrustDistributionStateBlocked    = "verifier_ecosystem_valc_trust_distribution_blocked"
	VerifierEcosystemValCTrustDistributionStateUnknown    = "verifier_ecosystem_valc_trust_distribution_unknown"

	VerifierEcosystemValCStateActive     = "verifier_ecosystem_valc_active"
	VerifierEcosystemValCStatePartial    = "verifier_ecosystem_valc_partial"
	VerifierEcosystemValCStateIncomplete = "verifier_ecosystem_valc_incomplete"
	VerifierEcosystemValCStateBlocked    = "verifier_ecosystem_valc_blocked"
	VerifierEcosystemValCStateUnknown    = "verifier_ecosystem_valc_unknown"

	VerifierEcosystemValCAudiencePublic             = "public"
	VerifierEcosystemValCAudiencePartner            = "partner"
	VerifierEcosystemValCAudienceAuditor            = "auditor"
	VerifierEcosystemValCAudienceInternal           = "internal_diagnostic"
	VerifierEcosystemValCAudiencePublisherSelfCheck = "publisher_self_check"
	VerifierEcosystemValCAudienceUnknown            = "unknown"

	VerifierEcosystemValCRequestModeUploadDescriptor        = "upload_descriptor"
	VerifierEcosystemValCRequestModeReferenceDescriptor     = "reference_descriptor"
	VerifierEcosystemValCRequestModeOfflineBundleDescriptor = "offline_bundle_descriptor"
	VerifierEcosystemValCRequestModeAPIReferenceDescriptor  = "api_reference_descriptor"
	VerifierEcosystemValCRequestModeUnknown                 = "unknown"

	VerifierEcosystemValCPublisherTypeVendor   = "vendor"
	VerifierEcosystemValCPublisherTypePartner  = "partner"
	VerifierEcosystemValCPublisherTypeCustomer = "customer"
	VerifierEcosystemValCPublisherTypeInternal = "internal"
	VerifierEcosystemValCPublisherTypeUnknown  = "unknown"

	VerifierEcosystemValCDistributionModeBundledTrustMaterial = "bundled_trust_material"
	VerifierEcosystemValCDistributionModeReferenceLookup      = "reference_lookup"
	VerifierEcosystemValCDistributionModeOfflineBundle        = "offline_bundle"
	VerifierEcosystemValCDistributionModePartnerScopedDir     = "partner_scoped_directory"
	VerifierEcosystemValCDistributionModeAuditorScopedDir     = "auditor_scoped_directory"
	VerifierEcosystemValCDistributionModeUnknown              = "unknown"
)

type VerifierEcosystemValCDependencySnapshot struct {
	Point5State                    string `json:"point_5_state"`
	Point5DependencyState          string `json:"point_5_dependency_state"`
	Point6State                    string `json:"point_6_state"`
	Point6ClosureState             string `json:"point_6_closure_state"`
	Point6ClosurePrerequisiteState string `json:"point_6_closure_prerequisite_state"`
	Point6ClosureInvariantState    string `json:"point_6_closure_invariant_state"`
	Point6ProofSurfaceState        string `json:"point_6_proof_surface_state"`
	Point6PassRuleState            string `json:"point_6_pass_rule_state"`
	Point6PassAllowed              bool   `json:"point_6_pass_allowed"`
	Val0CurrentState               string `json:"val_0_current_state"`
	Val0State                      string `json:"val_0_state"`
	ValACurrentState               string `json:"val_a_current_state"`
	ValAState                      string `json:"val_a_state"`
	ValBCurrentState               string `json:"val_b_current_state"`
	ValBState                      string `json:"val_b_state"`
	Point7State                    string `json:"point_7_state"`
}

type VerifierEcosystemValCAudienceSurface struct {
	CurrentState             string   `json:"current_state"`
	AudienceSurfaceID        string   `json:"audience_surface_id"`
	Version                  string   `json:"version"`
	AudienceType             string   `json:"audience_type"`
	AllowedScopeClasses      []string `json:"allowed_scope_classes,omitempty"`
	AllowedOutputClasses     []string `json:"allowed_output_classes,omitempty"`
	AllowedDiagnosticClasses []string `json:"allowed_diagnostic_classes,omitempty"`
	RedactionPolicyRef       string   `json:"redaction_policy_ref"`
	TrustMaterialVisibility  string   `json:"trust_material_visibility"`
	EvidenceVisibilityPolicy string   `json:"evidence_visibility_policy"`
	ExportAllowed            bool     `json:"export_allowed"`
	RepeatabilityRequired    bool     `json:"repeatability_required"`
	Caveats                  []string `json:"caveats,omitempty"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
	LifecycleState           string   `json:"lifecycle_state"`
	CompatibilityState       string   `json:"compatibility_state"`
	CreatedAt                string   `json:"created_at"`
	UpdatedAt                string   `json:"updated_at"`
	PublicReuseAllowed       bool     `json:"public_reuse_allowed"`
	CertificationClaim       bool     `json:"certification_claim"`
}

type VerifierEcosystemValCAudienceSurfaceCatalog struct {
	CurrentState           string                                 `json:"current_state"`
	AudienceCatalogID      string                                 `json:"audience_catalog_id"`
	SupportedAudienceTypes []string                               `json:"supported_audience_types,omitempty"`
	Surfaces               []VerifierEcosystemValCAudienceSurface `json:"surfaces,omitempty"`
	ProjectionDisclaimer   string                                 `json:"projection_disclaimer"`
}

type VerifierEcosystemValCPublicOutputContract struct {
	CurrentState                  string   `json:"current_state"`
	OutputID                      string   `json:"output_id"`
	AudienceSurfaceRef            string   `json:"audience_surface_ref"`
	ProofFingerprint              string   `json:"proof_fingerprint"`
	ProofType                     string   `json:"proof_type"`
	SchemaVersion                 string   `json:"schema_version"`
	VerificationSummary           string   `json:"verification_summary"`
	OverallResult                 string   `json:"overall_result"`
	DiagnosticClass               string   `json:"diagnostic_class"`
	OutputClass                   string   `json:"output_class"`
	FreshnessState                string   `json:"freshness_state"`
	CompatibilityState            string   `json:"compatibility_state"`
	IssuerVisibility              string   `json:"issuer_visibility"`
	TrustRootVisibility           string   `json:"trust_root_visibility"`
	RedactedFields                []string `json:"redacted_fields,omitempty"`
	RequiredCaveats               []string `json:"required_caveats,omitempty"`
	Limitations                   []string `json:"limitations,omitempty"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	SensitiveTrustMaterialExposed bool     `json:"sensitive_trust_material_exposed"`
	CertificationClaim            bool     `json:"certification_claim"`
	UniversalTruthClaim           bool     `json:"universal_truth_claim"`
	RegulatorApprovalClaim        bool     `json:"regulator_approval_claim"`
	DeploymentApprovalClaim       bool     `json:"deployment_approval_claim"`
}

type VerifierEcosystemValCPartnerOutputContract struct {
	CurrentState                   string   `json:"current_state"`
	OutputID                       string   `json:"output_id"`
	AudienceSurfaceRef             string   `json:"audience_surface_ref"`
	ProofFingerprint               string   `json:"proof_fingerprint"`
	ProofType                      string   `json:"proof_type"`
	SchemaVersion                  string   `json:"schema_version"`
	OverallResult                  string   `json:"overall_result"`
	DiagnosticClass                string   `json:"diagnostic_class"`
	OutputClass                    string   `json:"output_class"`
	CompatibilityState             string   `json:"compatibility_state"`
	FreshnessState                 string   `json:"freshness_state"`
	IssuerRefVisibility            string   `json:"issuer_ref_visibility"`
	TrustRootRefVisibility         string   `json:"trust_root_ref_visibility"`
	IntegrationContext             string   `json:"integration_context"`
	AllowedDetailLevel             string   `json:"allowed_detail_level"`
	EvidenceRefPolicy              string   `json:"evidence_ref_policy"`
	RedactedFields                 []string `json:"redacted_fields,omitempty"`
	Caveats                        []string `json:"caveats,omitempty"`
	ProjectionDisclaimer           string   `json:"projection_disclaimer"`
	InternalOnlyDiagnosticsExposed bool     `json:"internal_only_diagnostics_exposed"`
	MutatesEvidence                bool     `json:"mutates_evidence"`
	ApprovesDeployment             bool     `json:"approves_deployment"`
	SuppressesFailures             bool     `json:"suppresses_failures"`
	PublishesClaims                bool     `json:"publishes_claims"`
	CanonicalTruthClaim            bool     `json:"canonical_truth_claim"`
}

type VerifierEcosystemValCAuditorFlowContract struct {
	CurrentState                 string   `json:"current_state"`
	AuditorFlowID                string   `json:"auditor_flow_id"`
	Version                      string   `json:"version"`
	VerifierContractRef          string   `json:"verifier_contract_ref"`
	ProofEnvelopeRef             string   `json:"proof_envelope_ref"`
	FixtureOrCaseRefs            []string `json:"fixture_or_case_refs,omitempty"`
	RequiredEvidenceRefs         []string `json:"required_evidence_refs,omitempty"`
	RepeatabilityInputs          []string `json:"repeatability_inputs,omitempty"`
	DeterministicReportRef       string   `json:"deterministic_report_ref"`
	OutputBoundaryRef            string   `json:"output_boundary_ref"`
	TrustRootMaterialRef         string   `json:"trust_root_material_ref"`
	RevocationMaterialRef        string   `json:"revocation_material_ref"`
	SupersessionMaterialRef      string   `json:"supersession_material_ref"`
	CompatibilityPolicyRef       string   `json:"compatibility_policy_ref"`
	FreshnessPolicyRef           string   `json:"freshness_policy_ref"`
	PreservedDiagnostics         []string `json:"preserved_diagnostics,omitempty"`
	Caveats                      []string `json:"caveats,omitempty"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
	Repeatable                   bool     `json:"repeatable"`
	EvidenceLinked               bool     `json:"evidence_linked"`
	HiddenMainInstanceDependency bool     `json:"hidden_main_instance_dependency"`
	CertifiesOrganization        bool     `json:"certifies_organization"`
}

type VerifierEcosystemValCRequestContract struct {
	CurrentState              string   `json:"current_state"`
	RequestContractID         string   `json:"request_contract_id"`
	RequestMode               string   `json:"request_mode"`
	AcceptedInputDescriptors  []string `json:"accepted_input_descriptors,omitempty"`
	AllowedArtifactTypes      []string `json:"allowed_artifact_types,omitempty"`
	RequiredMetadata          []string `json:"required_metadata,omitempty"`
	MaxScopeClass             string   `json:"max_scope_class"`
	AllowedOutputBoundaries   []string `json:"allowed_output_boundaries,omitempty"`
	TrustMaterialRequirements []string `json:"trust_material_requirements,omitempty"`
	RedactionPolicyRef        string   `json:"redaction_policy_ref"`
	AbuseOrMisuseLimitations  []string `json:"abuse_or_misuse_limitations,omitempty"`
	Caveats                   []string `json:"caveats,omitempty"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
	IngestsCanonicalEvidence  bool     `json:"ingests_canonical_evidence"`
	ApprovesPublication       bool     `json:"approves_publication"`
	InternalArtifactForPublic bool     `json:"internal_artifact_for_public"`
}

type VerifierEcosystemValCPublisherCompatibilityProfile struct {
	CurrentState                string   `json:"current_state"`
	PublisherProfileID          string   `json:"publisher_profile_id"`
	Version                     string   `json:"version"`
	PublisherType               string   `json:"publisher_type"`
	SupportedArtifactTypes      []string `json:"supported_artifact_types,omitempty"`
	RequiredSchemaVersions      []string `json:"required_schema_versions,omitempty"`
	RequiredProofEnvelopeFields []string `json:"required_proof_envelope_fields,omitempty"`
	RequiredSignaturePolicy     string   `json:"required_signature_policy"`
	RequiredDigestPolicy        string   `json:"required_digest_policy"`
	IssuerBindingPolicy         string   `json:"issuer_binding_policy"`
	TrustRootPolicy             string   `json:"trust_root_policy"`
	RevocationSuppressionPolicy string   `json:"revocation_suppression_policy"`
	CompatibilityPolicyRef      string   `json:"compatibility_policy_ref"`
	ConformanceCaseRefs         []string `json:"conformance_case_refs,omitempty"`
	ForbiddenClaims             []string `json:"forbidden_claims,omitempty"`
	ObservedClaims              []string `json:"observed_claims,omitempty"`
	Caveats                     []string `json:"caveats,omitempty"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
	ApprovedVendorClaim         bool     `json:"approved_vendor_claim"`
	AutomaticallyTrustedClaim   bool     `json:"automatically_trusted_claim"`
}

type VerifierEcosystemValCArtifactPublishingRule struct {
	CurrentState              string   `json:"current_state"`
	RuleID                    string   `json:"rule_id"`
	PublisherProfileRef       string   `json:"publisher_profile_ref"`
	ArtifactType              string   `json:"artifact_type"`
	SchemaVersion             string   `json:"schema_version"`
	ProofType                 string   `json:"proof_type"`
	RequiredFields            []string `json:"required_fields,omitempty"`
	ForbiddenFields           []string `json:"forbidden_fields,omitempty"`
	RequiredDiagnostics       []string `json:"required_diagnostics,omitempty"`
	OutputBoundaryConstraints []string `json:"output_boundary_constraints,omitempty"`
	CompatibilityState        string   `json:"compatibility_state"`
	Caveats                   []string `json:"caveats,omitempty"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
	ObservedFields            []string `json:"observed_fields,omitempty"`
	ObservedClaims            []string `json:"observed_claims,omitempty"`
	SelectedOutputBoundary    string   `json:"selected_output_boundary"`
}

type VerifierEcosystemValCArtifactPublishingRuleCatalog struct {
	CurrentState              string                                        `json:"current_state"`
	RuleCatalogID             string                                        `json:"rule_catalog_id"`
	SupportedArtifactTypes    []string                                      `json:"supported_artifact_types,omitempty"`
	SupportedOutputBoundaries []string                                      `json:"supported_output_boundaries,omitempty"`
	Rules                     []VerifierEcosystemValCArtifactPublishingRule `json:"rules,omitempty"`
	ProjectionDisclaimer      string                                        `json:"projection_disclaimer"`
}

type VerifierEcosystemValCTrustDistributionVisibility struct {
	CurrentState                 string   `json:"current_state"`
	TrustDistributionID          string   `json:"trust_distribution_id"`
	IssuerDiscoveryPolicy        string   `json:"issuer_discovery_policy"`
	TrustRootDistributionMode    string   `json:"trust_root_distribution_mode"`
	TrustRootVersion             string   `json:"trust_root_version"`
	KeyRotationState             string   `json:"key_rotation_state"`
	RolloverMetadataRef          string   `json:"rollover_metadata_ref"`
	RevocationMaterialRef        string   `json:"revocation_material_ref"`
	SupersessionMaterialRef      string   `json:"supersession_material_ref"`
	OfflineDistributionSupported bool     `json:"offline_distribution_supported"`
	AudienceVisibilityScope      string   `json:"audience_visibility_scope"`
	Caveats                      []string `json:"caveats,omitempty"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
	TrustRootState               string   `json:"trust_root_state"`
	RevocationState              string   `json:"revocation_state"`
	GlobalKeyDirectoryClaim      bool     `json:"global_key_directory_claim"`
	SensitiveKeyMaterialExposed  bool     `json:"sensitive_key_material_exposed"`
}

func verifierEcosystemValCProjectionDisclaimer() string {
	return "projection_only not_canonical_truth public_partner_auditor_publisher advisory_projection"
}

func verifierEcosystemValCAudienceTypes() []string {
	return []string{
		VerifierEcosystemValCAudiencePublic,
		VerifierEcosystemValCAudiencePartner,
		VerifierEcosystemValCAudienceAuditor,
		VerifierEcosystemValCAudienceInternal,
		VerifierEcosystemValCAudiencePublisherSelfCheck,
		VerifierEcosystemValCAudienceUnknown,
	}
}

func verifierEcosystemValCRequestModes() []string {
	return []string{
		VerifierEcosystemValCRequestModeUploadDescriptor,
		VerifierEcosystemValCRequestModeReferenceDescriptor,
		VerifierEcosystemValCRequestModeOfflineBundleDescriptor,
		VerifierEcosystemValCRequestModeAPIReferenceDescriptor,
		VerifierEcosystemValCRequestModeUnknown,
	}
}

func verifierEcosystemValCPublisherTypes() []string {
	return []string{
		VerifierEcosystemValCPublisherTypeVendor,
		VerifierEcosystemValCPublisherTypePartner,
		VerifierEcosystemValCPublisherTypeCustomer,
		VerifierEcosystemValCPublisherTypeInternal,
		VerifierEcosystemValCPublisherTypeUnknown,
	}
}

func verifierEcosystemValCDistributionModes() []string {
	return []string{
		VerifierEcosystemValCDistributionModeBundledTrustMaterial,
		VerifierEcosystemValCDistributionModeReferenceLookup,
		VerifierEcosystemValCDistributionModeOfflineBundle,
		VerifierEcosystemValCDistributionModePartnerScopedDir,
		VerifierEcosystemValCDistributionModeAuditorScopedDir,
		VerifierEcosystemValCDistributionModeUnknown,
	}
}

func verifierEcosystemValCRequiredAudienceSurfaceIDs() []string {
	return []string{
		"audience-surface:public",
		"audience-surface:partner",
		"audience-surface:auditor",
		"audience-surface:internal-diagnostic",
		"audience-surface:publisher-self-check",
	}
}

func verifierEcosystemValCRequiredAudienceTypes() []string {
	return []string{
		VerifierEcosystemValCAudiencePublic,
		VerifierEcosystemValCAudiencePartner,
		VerifierEcosystemValCAudienceAuditor,
		VerifierEcosystemValCAudienceInternal,
		VerifierEcosystemValCAudiencePublisherSelfCheck,
	}
}

func verifierEcosystemValCRequiredArtifactRuleIDs() []string {
	return []string{
		"artifact-rule:signed-attestation-public",
		"artifact-rule:sealed-artifact-partner",
		"artifact-rule:lineage-bundle-auditor",
	}
}

func verifierEcosystemValCRequiredRequestMetadata() []string {
	return []string{
		"proof_fingerprint",
		"proof_type",
		"schema_version",
		"artifact_descriptor",
		"requested_scope",
		"output_boundary",
	}
}

func verifierEcosystemValCRequiredPublisherEnvelopeFields() []string {
	return []string{
		"proof_type",
		"schema_version",
		"artifact_digest_ref",
		"signature_ref",
		"issuer_ref",
		"trust_root_ref",
	}
}

func verifierEcosystemValCFreshnessStates() []string {
	return []string{
		IntelligenceCalibrationFreshnessFresh,
		IntelligenceCalibrationFreshnessStale,
		IntelligenceCalibrationFreshnessExpired,
		IntelligenceCalibrationFreshnessUnknown,
		IntelligenceCalibrationFreshnessUnsupported,
	}
}

func verifierEcosystemValCRequiredPreservedDiagnostics() []string {
	return []string{
		VerifierEcosystemDiagnosticStaleArtifact,
		VerifierEcosystemDiagnosticIncompleteArtifact,
		VerifierEcosystemDiagnosticUnsupportedSchema,
		VerifierEcosystemDiagnosticRevokedIssuer,
		VerifierEcosystemDiagnosticSupersededProof,
	}
}

func VerifierEcosystemValCAudienceSurfaceCatalogModel() VerifierEcosystemValCAudienceSurfaceCatalog {
	return verifierEcosystemValCAudienceSurfaceCatalogModel()
}

func verifierEcosystemValCAudienceSurfaceCatalogModel() VerifierEcosystemValCAudienceSurfaceCatalog {
	disclaimer := verifierEcosystemValCProjectionDisclaimer()
	return VerifierEcosystemValCAudienceSurfaceCatalog{
		CurrentState:           "verifier_ecosystem_valc_audience_surface_catalog_ready",
		AudienceCatalogID:      "verifier-audience-surfaces-valc",
		SupportedAudienceTypes: verifierEcosystemValCAudienceTypes(),
		Surfaces: []VerifierEcosystemValCAudienceSurface{
			{
				CurrentState:             "audience_surface_ready",
				AudienceSurfaceID:        "audience-surface:public",
				Version:                  "2026.04",
				AudienceType:             VerifierEcosystemValCAudiencePublic,
				AllowedScopeClasses:      []string{VerifierEcosystemScopePublicSafe},
				AllowedOutputClasses:     []string{VerifierEcosystemValBOutputClassVerified, VerifierEcosystemValBOutputClassVerifiedWithWarnings, VerifierEcosystemValBOutputClassNonVerifiedInvalid, VerifierEcosystemValBOutputClassNonVerifiedUnsupported, VerifierEcosystemValBOutputClassNonVerifiedStale, VerifierEcosystemValBOutputClassNonVerifiedRevoked, VerifierEcosystemValBOutputClassNonVerifiedSuperseded, VerifierEcosystemValBOutputClassRedactionBlocked},
				AllowedDiagnosticClasses: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticInvalidSignature, VerifierEcosystemDiagnosticDigestMismatch, VerifierEcosystemDiagnosticUnsupportedSchema, VerifierEcosystemDiagnosticUnsupportedProofType, VerifierEcosystemDiagnosticStaleArtifact, VerifierEcosystemDiagnosticExpiredArtifact, VerifierEcosystemDiagnosticRevokedIssuer, VerifierEcosystemDiagnosticSupersededProof, VerifierEcosystemDiagnosticRedactionViolation},
				RedactionPolicyRef:       "redaction-policy:public-safe",
				TrustMaterialVisibility:  "summary_only",
				EvidenceVisibilityPolicy: "public_caveated",
				ExportAllowed:            true,
				Caveats:                  []string{"public-safe verifier output remains redacted and bounded to declared proof scope"},
				ProjectionDisclaimer:     disclaimer,
				LifecycleState:           "active",
				CompatibilityState:       ReferenceArchitectureCompatibilityCompatible,
				CreatedAt:                "2026-04-27T11:40:00Z",
				UpdatedAt:                "2026-04-27T11:40:00Z",
				PublicReuseAllowed:       true,
			},
			{
				CurrentState:             "audience_surface_ready",
				AudienceSurfaceID:        "audience-surface:partner",
				Version:                  "2026.04",
				AudienceType:             VerifierEcosystemValCAudiencePartner,
				AllowedScopeClasses:      []string{VerifierEcosystemScopePartnerSafe},
				AllowedOutputClasses:     []string{VerifierEcosystemValBOutputClassVerified, VerifierEcosystemValBOutputClassVerifiedWithWarnings, VerifierEcosystemValBOutputClassNonVerifiedInvalid, VerifierEcosystemValBOutputClassNonVerifiedIncomplete, VerifierEcosystemValBOutputClassNonVerifiedUnsupported, VerifierEcosystemValBOutputClassNonVerifiedStale, VerifierEcosystemValBOutputClassNonVerifiedRevoked, VerifierEcosystemValBOutputClassNonVerifiedSuperseded, VerifierEcosystemValBOutputClassNonVerifiedScopeMismatch, VerifierEcosystemValBOutputClassRedactionBlocked},
				AllowedDiagnosticClasses: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticInvalidSignature, VerifierEcosystemDiagnosticDigestMismatch, VerifierEcosystemDiagnosticUnsupportedSchema, VerifierEcosystemDiagnosticUnsupportedProofType, VerifierEcosystemDiagnosticStaleArtifact, VerifierEcosystemDiagnosticExpiredArtifact, VerifierEcosystemDiagnosticRevokedIssuer, VerifierEcosystemDiagnosticSupersededProof, VerifierEcosystemDiagnosticScopeMismatch, VerifierEcosystemDiagnosticRedactionViolation},
				RedactionPolicyRef:       "redaction-policy:partner-safe",
				TrustMaterialVisibility:  "issuer_scoped",
				EvidenceVisibilityPolicy: "partner_scoped",
				ExportAllowed:            true,
				Caveats:                  []string{"partner-safe verifier output remains scoped to bounded integration context and cannot become canonical truth"},
				ProjectionDisclaimer:     disclaimer,
				LifecycleState:           "active",
				CompatibilityState:       ReferenceArchitectureCompatibilityCompatible,
				CreatedAt:                "2026-04-27T11:41:00Z",
				UpdatedAt:                "2026-04-27T11:41:00Z",
			},
			{
				CurrentState:             "audience_surface_ready",
				AudienceSurfaceID:        "audience-surface:auditor",
				Version:                  "2026.04",
				AudienceType:             VerifierEcosystemValCAudienceAuditor,
				AllowedScopeClasses:      []string{VerifierEcosystemScopeAuditorSafe},
				AllowedOutputClasses:     verifierEcosystemValBOutputClasses(),
				AllowedDiagnosticClasses: verifierEcosystemVal0DiagnosticClasses(),
				RedactionPolicyRef:       "redaction-policy:auditor-safe",
				TrustMaterialVisibility:  "auditor_scoped",
				EvidenceVisibilityPolicy: "auditor_evidence_linked",
				ExportAllowed:            true,
				RepeatabilityRequired:    true,
				Caveats:                  []string{"auditor-safe verification remains repeatable and evidence-linked rather than certifying organizations"},
				ProjectionDisclaimer:     disclaimer,
				LifecycleState:           "active",
				CompatibilityState:       ReferenceArchitectureCompatibilityCompatible,
				CreatedAt:                "2026-04-27T11:42:00Z",
				UpdatedAt:                "2026-04-27T11:42:00Z",
			},
			{
				CurrentState:             "audience_surface_ready",
				AudienceSurfaceID:        "audience-surface:internal-diagnostic",
				Version:                  "2026.04",
				AudienceType:             VerifierEcosystemValCAudienceInternal,
				AllowedScopeClasses:      []string{VerifierEcosystemScopeInternalDiagnostic},
				AllowedOutputClasses:     verifierEcosystemValBOutputClasses(),
				AllowedDiagnosticClasses: verifierEcosystemVal0DiagnosticClasses(),
				RedactionPolicyRef:       "redaction-policy:internal-diagnostic",
				TrustMaterialVisibility:  "internal_scoped",
				EvidenceVisibilityPolicy: "internal_only",
				Caveats:                  []string{"internal diagnostic verifier output must not be reused as public-safe output"},
				ProjectionDisclaimer:     disclaimer,
				LifecycleState:           "active",
				CompatibilityState:       ReferenceArchitectureCompatibilityCompatible,
				CreatedAt:                "2026-04-27T11:43:00Z",
				UpdatedAt:                "2026-04-27T11:43:00Z",
			},
			{
				CurrentState:             "audience_surface_ready",
				AudienceSurfaceID:        "audience-surface:publisher-self-check",
				Version:                  "2026.04",
				AudienceType:             VerifierEcosystemValCAudiencePublisherSelfCheck,
				AllowedScopeClasses:      []string{VerifierEcosystemScopePartnerSafe, VerifierEcosystemScopeRestrictedOffline},
				AllowedOutputClasses:     []string{VerifierEcosystemValBOutputClassVerified, VerifierEcosystemValBOutputClassVerifiedWithWarnings, VerifierEcosystemValBOutputClassNonVerifiedInvalid, VerifierEcosystemValBOutputClassNonVerifiedIncomplete, VerifierEcosystemValBOutputClassNonVerifiedUnsupported, VerifierEcosystemValBOutputClassNonVerifiedStale, VerifierEcosystemValBOutputClassNonVerifiedRevoked, VerifierEcosystemValBOutputClassNonVerifiedSuperseded},
				AllowedDiagnosticClasses: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticInvalidSignature, VerifierEcosystemDiagnosticDigestMismatch, VerifierEcosystemDiagnosticUnsupportedSchema, VerifierEcosystemDiagnosticUnsupportedProofType, VerifierEcosystemDiagnosticStaleArtifact, VerifierEcosystemDiagnosticExpiredArtifact, VerifierEcosystemDiagnosticRevokedIssuer, VerifierEcosystemDiagnosticSupersededProof, VerifierEcosystemDiagnosticIncompleteArtifact},
				RedactionPolicyRef:       "redaction-policy:publisher-self-check",
				TrustMaterialVisibility:  "publisher_scoped",
				EvidenceVisibilityPolicy: "publisher_scoped",
				ExportAllowed:            false,
				Caveats:                  []string{"publisher self-check output remains compatibility guidance only and does not imply certification"},
				ProjectionDisclaimer:     disclaimer,
				LifecycleState:           "active",
				CompatibilityState:       ReferenceArchitectureCompatibilityCompatible,
				CreatedAt:                "2026-04-27T11:44:00Z",
				UpdatedAt:                "2026-04-27T11:44:00Z",
			},
		},
		ProjectionDisclaimer: disclaimer,
	}
}

func VerifierEcosystemValCPublicOutputContractModel() VerifierEcosystemValCPublicOutputContract {
	return VerifierEcosystemValCPublicOutputContract{
		CurrentState:         "verifier_ecosystem_valc_public_output_ready",
		OutputID:             "public-output:verification-summary",
		AudienceSurfaceRef:   "audience-surface:public",
		ProofFingerprint:     "sha256:public-proof-fingerprint",
		ProofType:            VerifierEcosystemProofTypeSignedAttestation,
		SchemaVersion:        "changelock.verifier.proof_envelope.v1",
		VerificationSummary:  "public-safe verifier output remains bounded to advisory proof summary",
		OverallResult:        VerifierEcosystemValAOverallResultVerified,
		DiagnosticClass:      VerifierEcosystemDiagnosticVerified,
		OutputClass:          VerifierEcosystemValBOutputClassVerified,
		FreshnessState:       IntelligenceCalibrationFreshnessFresh,
		CompatibilityState:   ReferenceArchitectureCompatibilityCompatible,
		IssuerVisibility:     "summary_only",
		TrustRootVisibility:  "summary_only",
		RedactedFields:       []string{"issuer_ref", "trust_root_ref", "revocation_material_ref"},
		RequiredCaveats:      []string{"public-safe output remains redacted and cannot become canonical truth"},
		Limitations:          []string{"public-safe output remains bounded to proof summary, compatibility state, freshness state, and verifier diagnostic class"},
		ProjectionDisclaimer: verifierEcosystemValCProjectionDisclaimer(),
	}
}

func VerifierEcosystemValCPartnerOutputContractModel() VerifierEcosystemValCPartnerOutputContract {
	return VerifierEcosystemValCPartnerOutputContract{
		CurrentState:           "verifier_ecosystem_valc_partner_output_ready",
		OutputID:               "partner-output:verification-summary",
		AudienceSurfaceRef:     "audience-surface:partner",
		ProofFingerprint:       "sha256:partner-proof-fingerprint",
		ProofType:              VerifierEcosystemProofTypeSignedAttestation,
		SchemaVersion:          "changelock.verifier.proof_envelope.v1",
		OverallResult:          VerifierEcosystemValAOverallResultVerified,
		DiagnosticClass:        VerifierEcosystemDiagnosticVerified,
		OutputClass:            VerifierEcosystemValBOutputClassVerified,
		CompatibilityState:     ReferenceArchitectureCompatibilityCompatible,
		FreshnessState:         IntelligenceCalibrationFreshnessFresh,
		IssuerRefVisibility:    "issuer_scoped",
		TrustRootRefVisibility: "reference_scoped",
		IntegrationContext:     "partner_api_reference",
		AllowedDetailLevel:     "integration_scoped",
		EvidenceRefPolicy:      "partner_scoped",
		RedactedFields:         []string{"key_or_material_ref"},
		Caveats:                []string{"partner-safe output remains scoped and cannot become canonical truth or deployment approval"},
		ProjectionDisclaimer:   verifierEcosystemValCProjectionDisclaimer(),
	}
}

func VerifierEcosystemValCAuditorFlowContractModel() VerifierEcosystemValCAuditorFlowContract {
	return VerifierEcosystemValCAuditorFlowContract{
		CurrentState:            "verifier_ecosystem_valc_auditor_flow_ready",
		AuditorFlowID:           "auditor-flow:repeatable-verification",
		Version:                 "2026.04",
		VerifierContractRef:     "verifier-contract-val0",
		ProofEnvelopeRef:        "proof-envelope-reference-val0",
		FixtureOrCaseRefs:       []string{"fixture:valid-proof-envelope", "conformance:valid-proof-envelope"},
		RequiredEvidenceRefs:    []string{"evidence:verifier-contract-001", "evidence:proof-envelope-001", "evidence:compatibility-matrix-001", "evidence:conformance-suite-001"},
		RepeatabilityInputs:     []string{"proof_fingerprint", "schema_version", "proof_type", "trust_root_version"},
		DeterministicReportRef:  "verification-report:vala-current",
		OutputBoundaryRef:       "boundary/auditor-safe",
		TrustRootMaterialRef:    "trust-root:reference-program",
		RevocationMaterialRef:   "evidence:revocation-001",
		SupersessionMaterialRef: "supersession:lineage-bundle-current",
		CompatibilityPolicyRef:  "verifier-compatibility-matrix-valb",
		FreshnessPolicyRef:      "freshness-policy:bounded-window",
		PreservedDiagnostics:    verifierEcosystemValCRequiredPreservedDiagnostics(),
		Caveats:                 []string{"auditor-safe flow remains repeatable and evidence-linked without certifying an organization"},
		ProjectionDisclaimer:    verifierEcosystemValCProjectionDisclaimer(),
		Repeatable:              true,
		EvidenceLinked:          true,
	}
}

func VerifierEcosystemValCRequestContractModel() VerifierEcosystemValCRequestContract {
	return VerifierEcosystemValCRequestContract{
		CurrentState:              "verifier_ecosystem_valc_request_contract_ready",
		RequestContractID:         "verification-request-contract-valc",
		RequestMode:               VerifierEcosystemValCRequestModeUploadDescriptor,
		AcceptedInputDescriptors:  []string{VerifierEcosystemValCRequestModeUploadDescriptor, VerifierEcosystemValCRequestModeReferenceDescriptor, VerifierEcosystemValCRequestModeOfflineBundleDescriptor, VerifierEcosystemValCRequestModeAPIReferenceDescriptor},
		AllowedArtifactTypes:      verifierEcosystemVal0SupportedProofTypes(),
		RequiredMetadata:          verifierEcosystemValCRequiredRequestMetadata(),
		MaxScopeClass:             VerifierEcosystemScopeAuditorSafe,
		AllowedOutputBoundaries:   []string{VerifierEcosystemScopePublicSafe, VerifierEcosystemScopePartnerSafe, VerifierEcosystemScopeAuditorSafe, VerifierEcosystemScopeRestrictedOffline},
		TrustMaterialRequirements: []string{"issuer_ref_required", "trust_root_ref_required", "revocation_metadata_required"},
		RedactionPolicyRef:        "redaction-policy:request-contract",
		AbuseOrMisuseLimitations:  []string{"request contract remains descriptor-only and does not ingest canonical evidence or approve publication"},
		Caveats:                   []string{"upload and reference descriptors remain bounded submission contracts only"},
		ProjectionDisclaimer:      verifierEcosystemValCProjectionDisclaimer(),
	}
}

func VerifierEcosystemValCPublisherCompatibilityProfileModel() VerifierEcosystemValCPublisherCompatibilityProfile {
	return VerifierEcosystemValCPublisherCompatibilityProfile{
		CurrentState:                "verifier_ecosystem_valc_publisher_profile_ready",
		PublisherProfileID:          "publisher-profile:vendor-compatible",
		Version:                     "2026.04",
		PublisherType:               VerifierEcosystemValCPublisherTypeVendor,
		SupportedArtifactTypes:      verifierEcosystemVal0SupportedProofTypes(),
		RequiredSchemaVersions:      VerifierEcosystemValBCompatibilityMatrixModel().SupportedSchemaVersions,
		RequiredProofEnvelopeFields: verifierEcosystemValCRequiredPublisherEnvelopeFields(),
		RequiredSignaturePolicy:     "signature_ref_required",
		RequiredDigestPolicy:        "sha256_or_sha512_required",
		IssuerBindingPolicy:         "issuer_ref_bound",
		TrustRootPolicy:             "versioned_trust_root_required",
		RevocationSuppressionPolicy: "suppression_forbidden",
		CompatibilityPolicyRef:      "verifier-compatibility-matrix-valb",
		ConformanceCaseRefs:         verifierEcosystemValBRequiredConformanceCaseIDs(),
		ForbiddenClaims:             []string{"approved vendor", "certified publisher", "integrity rating", "regulator-approved verifier"},
		Caveats:                     []string{"publisher profile is compatibility guidance only and verifier-compatible artifacts are not automatically trusted"},
		ProjectionDisclaimer:        verifierEcosystemValCProjectionDisclaimer(),
	}
}

func VerifierEcosystemValCArtifactPublishingRuleCatalogModel() VerifierEcosystemValCArtifactPublishingRuleCatalog {
	disclaimer := verifierEcosystemValCProjectionDisclaimer()
	requiredFields := verifierEcosystemValCRequiredPublisherEnvelopeFields()
	return VerifierEcosystemValCArtifactPublishingRuleCatalog{
		CurrentState:              "verifier_ecosystem_valc_artifact_rule_catalog_ready",
		RuleCatalogID:             "artifact-publishing-rules-valc",
		SupportedArtifactTypes:    verifierEcosystemVal0SupportedProofTypes(),
		SupportedOutputBoundaries: []string{VerifierEcosystemScopePublicSafe, VerifierEcosystemScopePartnerSafe, VerifierEcosystemScopeAuditorSafe},
		Rules: []VerifierEcosystemValCArtifactPublishingRule{
			{
				CurrentState:              "artifact_rule_ready",
				RuleID:                    "artifact-rule:signed-attestation-public",
				PublisherProfileRef:       "publisher-profile:vendor-compatible",
				ArtifactType:              VerifierEcosystemProofTypeSignedAttestation,
				SchemaVersion:             "changelock.verifier.proof_envelope.v1",
				ProofType:                 VerifierEcosystemProofTypeSignedAttestation,
				RequiredFields:            requiredFields,
				ForbiddenFields:           []string{"certification_claim", "approved_vendor_claim"},
				RequiredDiagnostics:       []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticCompatibilityWarning},
				OutputBoundaryConstraints: []string{VerifierEcosystemScopePublicSafe, VerifierEcosystemScopePartnerSafe},
				CompatibilityState:        ReferenceArchitectureCompatibilityCompatible,
				Caveats:                   []string{"compatible artifact rule remains guidance only and does not certify the artifact"},
				ProjectionDisclaimer:      disclaimer,
				ObservedFields:            requiredFields,
				SelectedOutputBoundary:    VerifierEcosystemScopePublicSafe,
			},
			{
				CurrentState:              "artifact_rule_ready",
				RuleID:                    "artifact-rule:sealed-artifact-partner",
				PublisherProfileRef:       "publisher-profile:vendor-compatible",
				ArtifactType:              VerifierEcosystemProofTypeSealedArtifact,
				SchemaVersion:             "changelock.verifier.proof_envelope.v1.1",
				ProofType:                 VerifierEcosystemProofTypeSealedArtifact,
				RequiredFields:            requiredFields,
				ForbiddenFields:           []string{"certification_claim", "approved_vendor_claim"},
				RequiredDiagnostics:       []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticCompatibilityWarning},
				OutputBoundaryConstraints: []string{VerifierEcosystemScopePartnerSafe, VerifierEcosystemScopeAuditorSafe},
				CompatibilityState:        ReferenceArchitectureCompatibilityCompatible,
				Caveats:                   []string{"partner-safe sealed artifact publishing remains verifier-compatible but not automatically trusted"},
				ProjectionDisclaimer:      disclaimer,
				ObservedFields:            requiredFields,
				SelectedOutputBoundary:    VerifierEcosystemScopePartnerSafe,
			},
			{
				CurrentState:              "artifact_rule_ready",
				RuleID:                    "artifact-rule:lineage-bundle-auditor",
				PublisherProfileRef:       "publisher-profile:vendor-compatible",
				ArtifactType:              VerifierEcosystemProofTypeLineageBundle,
				SchemaVersion:             "changelock.verifier.proof_envelope.v1",
				ProofType:                 VerifierEcosystemProofTypeLineageBundle,
				RequiredFields:            requiredFields,
				ForbiddenFields:           []string{"certification_claim", "approved_vendor_claim"},
				RequiredDiagnostics:       []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticSupersededProof},
				OutputBoundaryConstraints: []string{VerifierEcosystemScopeAuditorSafe},
				CompatibilityState:        ReferenceArchitectureCompatibilityCompatible,
				Caveats:                   []string{"auditor-safe lineage publishing remains compatible only within declared trust and compatibility policy"},
				ProjectionDisclaimer:      disclaimer,
				ObservedFields:            requiredFields,
				SelectedOutputBoundary:    VerifierEcosystemScopeAuditorSafe,
			},
		},
		ProjectionDisclaimer: disclaimer,
	}
}

func VerifierEcosystemValCTrustDistributionVisibilityModel() VerifierEcosystemValCTrustDistributionVisibility {
	return VerifierEcosystemValCTrustDistributionVisibility{
		CurrentState:              "verifier_ecosystem_valc_trust_distribution_ready",
		TrustDistributionID:       "trust-distribution:partner-scoped",
		IssuerDiscoveryPolicy:     "issuer_ref_bounded_lookup",
		TrustRootDistributionMode: VerifierEcosystemValCDistributionModePartnerScopedDir,
		TrustRootVersion:          "2026.04",
		KeyRotationState:          VerifierEcosystemKeyRotationCurrent,
		RevocationMaterialRef:     "evidence:revocation-001",
		SupersessionMaterialRef:   "rollover:trust-root-2026.04",
		AudienceVisibilityScope:   VerifierEcosystemScopePartnerSafe,
		Caveats:                   []string{"trust-root distribution remains scoped, bounded, and not a global key registry"},
		ProjectionDisclaimer:      verifierEcosystemValCProjectionDisclaimer(),
		TrustRootState:            VerifierEcosystemTrustRootTrusted,
		RevocationState:           VerifierEcosystemRevocationNotRevoked,
	}
}

func verifierEcosystemValCAudienceLookup(model VerifierEcosystemValCAudienceSurfaceCatalog) map[string]VerifierEcosystemValCAudienceSurface {
	lookup := make(map[string]VerifierEcosystemValCAudienceSurface, len(model.Surfaces))
	for _, item := range model.Surfaces {
		lookup[strings.TrimSpace(item.AudienceSurfaceID)] = item
	}
	return lookup
}

func verifierEcosystemValCAudienceSurfaceKey(item VerifierEcosystemValCAudienceSurface) string {
	return strings.TrimSpace(item.AudienceSurfaceID)
}

func verifierEcosystemValCExpectedOutputClass(overallResult, diagnostic string, caveats []string) string {
	return deriveVerifierEcosystemValBOutputClass(overallResult, diagnostic, caveats)
}

func verifierEcosystemValCAllowedPublicVisibilities() []string {
	return []string{"summary_only", "redacted_summary"}
}

func verifierEcosystemValCRequiredEvidenceIDs() []string {
	return []string{
		"evidence:audience-surfaces-001",
		"evidence:public-output-001",
		"evidence:partner-output-001",
		"evidence:auditor-flow-001",
		"evidence:request-contract-001",
		"evidence:publisher-profile-001",
		"evidence:artifact-rules-001",
		"evidence:trust-distribution-001",
		"evidence:point7-governance-003",
	}
}

func verifierEcosystemValCRequiredEvidenceScopes() []string {
	return []string{
		"audience_surfaces",
		"public_output_contract",
		"partner_output_contract",
		"auditor_flow_contract",
		"request_contract",
		"publisher_profile",
		"artifact_publishing_rules",
		"trust_distribution_visibility",
		"point7_governance",
	}
}

func VerifierEcosystemValCVerifierEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:audience-surfaces-001", EvidenceType: "audience_surface_catalog", Source: "verifier/audience-surfaces", Timestamp: "2026-04-27T11:50:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "audience_surfaces", Caveats: []string{"bounded to public, partner, auditor, internal, and publisher self-check audience surfaces only"}},
		{EvidenceID: "evidence:public-output-001", EvidenceType: "public_output_contract", Source: "verifier/public-output", Timestamp: "2026-04-27T11:51:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "public_output_contract", Caveats: []string{"public-safe output remains redacted and advisory only"}},
		{EvidenceID: "evidence:partner-output-001", EvidenceType: "partner_output_contract", Source: "verifier/partner-output", Timestamp: "2026-04-27T11:52:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "partner_output_contract", Caveats: []string{"partner-safe output remains scoped and non-canonical"}},
		{EvidenceID: "evidence:auditor-flow-001", EvidenceType: "auditor_flow_contract", Source: "verifier/auditor-flow", Timestamp: "2026-04-27T11:53:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "auditor_flow_contract", Caveats: []string{"auditor-safe flow remains repeatable and evidence-linked"}},
		{EvidenceID: "evidence:request-contract-001", EvidenceType: "request_contract", Source: "verifier/request-contract", Timestamp: "2026-04-27T11:54:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "request_contract", Caveats: []string{"upload and reference descriptors remain bounded request contracts only"}},
		{EvidenceID: "evidence:publisher-profile-001", EvidenceType: "publisher_profile", Source: "verifier/publisher-profile", Timestamp: "2026-04-27T11:55:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "publisher_profile", Caveats: []string{"publisher profile remains compatibility guidance and not certification"}},
		{EvidenceID: "evidence:artifact-rules-001", EvidenceType: "artifact_publishing_rules", Source: "verifier/artifact-rules", Timestamp: "2026-04-27T11:56:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "artifact_publishing_rules", Caveats: []string{"artifact publishing rules remain bounded verifier-compatible guidance only"}},
		{EvidenceID: "evidence:trust-distribution-001", EvidenceType: "trust_distribution_visibility", Source: "verifier/trust-distribution", Timestamp: "2026-04-27T11:57:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "trust_distribution_visibility", Caveats: []string{"trust-root distribution remains scoped and bounded rather than global protocol authority"}},
		{EvidenceID: "evidence:point7-governance-003", EvidenceType: "state_governance", Source: "verifier/point7-governance", Timestamp: "2026-04-27T11:58:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point7_governance", Caveats: []string{"Val C keeps point_7_state not complete and cannot return point_7_pass"}},
	}
}

func VerifierEcosystemValCProofEvidenceRefs() []string {
	return []string{
		"point6_integrated_closure",
		"point7_verifier_discipline_foundation",
		"point7_reference_verifier_tooling",
		"point7_compatibility_diagnostics_conformance",
		"point7_public_partner_auditor_publisher_ecosystem",
		"evidence:audience-surfaces-001",
		"evidence:public-output-001",
		"evidence:partner-output-001",
		"evidence:auditor-flow-001",
		"evidence:request-contract-001",
		"evidence:publisher-profile-001",
		"evidence:artifact-rules-001",
		"evidence:trust-distribution-001",
		"evidence:point7-governance-003",
	}
}

func verifierEcosystemValCProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, evidenceRefs []string) bool {
	if !containsExactTrimmedStringSet(evidenceRefs, VerifierEcosystemValCProofEvidenceRefs()...) {
		return false
	}
	allFresh, stale, ok := verifierEcosystemVal0EvidenceValid(evidence)
	if !ok || !allFresh || stale {
		return false
	}
	evidenceIDs := make([]string, 0, len(evidence))
	evidenceScopes := make([]string, 0, len(evidence))
	for _, item := range evidence {
		evidenceIDs = append(evidenceIDs, item.EvidenceID)
		evidenceScopes = append(evidenceScopes, item.Scope)
	}
	return containsExactTrimmedStringSet(evidenceIDs, verifierEcosystemValCRequiredEvidenceIDs()...) &&
		containsExactTrimmedStringSet(evidenceScopes, verifierEcosystemValCRequiredEvidenceScopes()...)
}

func EvaluateVerifierEcosystemValCAudienceSurfaceState(model VerifierEcosystemValCAudienceSurfaceCatalog) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.AudienceCatalogID, model.ProjectionDisclaimer) || len(model.Surfaces) == 0 {
		return VerifierEcosystemValCAudienceSurfaceStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedAudienceTypes, verifierEcosystemValCAudienceTypes()...) || !verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValCAudienceSurfaceStatePartial
	}
	if verifierEcosystemVal0HasOverclaim(model.ProjectionDisclaimer) {
		return VerifierEcosystemValCAudienceSurfaceStateBlocked
	}
	surfaceIDs := make([]string, 0, len(model.Surfaces))
	surfaceTypes := make([]string, 0, len(model.Surfaces))
	outputClassCountsByAudienceType := make(map[string]int, len(model.Surfaces))
	for _, item := range model.Surfaces {
		if !referenceArchitectureValBRequiredRefsPresent(
			item.AudienceSurfaceID,
			item.Version,
			item.AudienceType,
			item.RedactionPolicyRef,
			item.TrustMaterialVisibility,
			item.EvidenceVisibilityPolicy,
			item.LifecycleState,
			item.CompatibilityState,
			item.ProjectionDisclaimer,
			item.CreatedAt,
			item.UpdatedAt,
		) || len(item.AllowedScopeClasses) == 0 || len(item.AllowedOutputClasses) == 0 || len(item.AllowedDiagnosticClasses) == 0 || len(item.Caveats) == 0 {
			return VerifierEcosystemValCAudienceSurfaceStateIncomplete
		}
		if !containsTrimmedString(model.SupportedAudienceTypes, item.AudienceType) ||
			!containsAllTrimmedStrings(verifierEcosystemVal0SupportedScopeClasses(), item.AllowedScopeClasses...) ||
			!containsAllTrimmedStrings(verifierEcosystemValBOutputClasses(), item.AllowedOutputClasses...) ||
			!containsAllTrimmedStrings(verifierEcosystemVal0DiagnosticClasses(), item.AllowedDiagnosticClasses...) ||
			!containsTrimmedString([]string{"active"}, item.LifecycleState) ||
			!containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), item.CompatibilityState) ||
			!verifierEcosystemVal0HasProjectionDisclaimer(item.ProjectionDisclaimer) {
			return VerifierEcosystemValCAudienceSurfaceStateUnknown
		}
		if _, ok := referenceArchitectureVal0ParseTimestamp(item.CreatedAt); !ok {
			return VerifierEcosystemValCAudienceSurfaceStatePartial
		}
		if _, ok := referenceArchitectureVal0ParseTimestamp(item.UpdatedAt); !ok {
			return VerifierEcosystemValCAudienceSurfaceStatePartial
		}
		if item.CertificationClaim || verifierEcosystemVal0HasOverclaim(strings.Join(item.Caveats, " "), item.ProjectionDisclaimer) {
			return VerifierEcosystemValCAudienceSurfaceStateBlocked
		}
		audienceType := strings.TrimSpace(item.AudienceType)
		if _, duplicate := outputClassCountsByAudienceType[audienceType]; duplicate {
			return VerifierEcosystemValCAudienceSurfaceStatePartial
		}
		switch audienceType {
		case VerifierEcosystemValCAudienceUnknown:
			return VerifierEcosystemValCAudienceSurfaceStateUnknown
		case VerifierEcosystemValCAudiencePublic:
			if !containsExactTrimmedStringSet(item.AllowedScopeClasses, VerifierEcosystemScopePublicSafe) ||
				!containsTrimmedString(verifierEcosystemValCAllowedPublicVisibilities(), item.TrustMaterialVisibility) {
				return VerifierEcosystemValCAudienceSurfaceStateBlocked
			}
		case VerifierEcosystemValCAudiencePartner:
			if !containsExactTrimmedStringSet(item.AllowedScopeClasses, VerifierEcosystemScopePartnerSafe) {
				return VerifierEcosystemValCAudienceSurfaceStatePartial
			}
		case VerifierEcosystemValCAudienceAuditor:
			if !item.RepeatabilityRequired || strings.TrimSpace(item.EvidenceVisibilityPolicy) != "auditor_evidence_linked" || !containsExactTrimmedStringSet(item.AllowedScopeClasses, VerifierEcosystemScopeAuditorSafe) {
				return VerifierEcosystemValCAudienceSurfaceStateBlocked
			}
		case VerifierEcosystemValCAudienceInternal:
			if item.PublicReuseAllowed || !containsExactTrimmedStringSet(item.AllowedScopeClasses, VerifierEcosystemScopeInternalDiagnostic) {
				return VerifierEcosystemValCAudienceSurfaceStateBlocked
			}
		case VerifierEcosystemValCAudiencePublisherSelfCheck:
			if containsTrimmedString(item.AllowedOutputClasses, VerifierEcosystemValBOutputClassRedactionBlocked) {
				return VerifierEcosystemValCAudienceSurfaceStatePartial
			}
		}
		surfaceIDs = append(surfaceIDs, verifierEcosystemValCAudienceSurfaceKey(item))
		surfaceTypes = append(surfaceTypes, audienceType)
		outputClassCountsByAudienceType[audienceType] = len(item.AllowedOutputClasses)
	}
	if !containsExactTrimmedStringSet(surfaceIDs, verifierEcosystemValCRequiredAudienceSurfaceIDs()...) {
		return VerifierEcosystemValCAudienceSurfaceStatePartial
	}
	if !containsExactTrimmedStringSet(surfaceTypes, verifierEcosystemValCRequiredAudienceTypes()...) {
		return VerifierEcosystemValCAudienceSurfaceStatePartial
	}
	publicOutputClassCount, publicFound := outputClassCountsByAudienceType[VerifierEcosystemValCAudiencePublic]
	partnerOutputClassCount, partnerFound := outputClassCountsByAudienceType[VerifierEcosystemValCAudiencePartner]
	if !publicFound || !partnerFound {
		return VerifierEcosystemValCAudienceSurfaceStatePartial
	}
	if partnerOutputClassCount <= publicOutputClassCount {
		return VerifierEcosystemValCAudienceSurfaceStatePartial
	}
	return VerifierEcosystemValCAudienceSurfaceStateActive
}

func EvaluateVerifierEcosystemValCPublicOutputState(model VerifierEcosystemValCPublicOutputContract, audiences VerifierEcosystemValCAudienceSurfaceCatalog) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.OutputID,
		model.AudienceSurfaceRef,
		model.ProofFingerprint,
		model.ProofType,
		model.SchemaVersion,
		model.VerificationSummary,
		model.OverallResult,
		model.DiagnosticClass,
		model.OutputClass,
		model.FreshnessState,
		model.CompatibilityState,
		model.IssuerVisibility,
		model.TrustRootVisibility,
		model.ProjectionDisclaimer,
	) || len(model.RedactedFields) == 0 || len(model.RequiredCaveats) == 0 || len(model.Limitations) == 0 {
		return VerifierEcosystemValCPublicOutputStateIncomplete
	}
	if EvaluateVerifierEcosystemValCAudienceSurfaceState(audiences) != VerifierEcosystemValCAudienceSurfaceStateActive {
		return VerifierEcosystemValCPublicOutputStateBlocked
	}
	if !containsTrimmedString(verifierEcosystemVal0SupportedProofTypes(), model.ProofType) ||
		!containsTrimmedString(VerifierEcosystemValBCompatibilityMatrixModel().SupportedSchemaVersions, model.SchemaVersion) ||
		!containsTrimmedString(verifierEcosystemValAOverallResults(), model.OverallResult) ||
		!containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), model.DiagnosticClass) ||
		!containsTrimmedString(verifierEcosystemValBOutputClasses(), model.OutputClass) ||
		!containsTrimmedString(verifierEcosystemValCFreshnessStates(), model.FreshnessState) ||
		!containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), model.CompatibilityState) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValCPublicOutputStateUnknown
	}
	if model.SensitiveTrustMaterialExposed || model.CertificationClaim || model.UniversalTruthClaim || model.RegulatorApprovalClaim || model.DeploymentApprovalClaim ||
		!containsTrimmedString(verifierEcosystemValCAllowedPublicVisibilities(), model.IssuerVisibility) ||
		!containsTrimmedString(verifierEcosystemValCAllowedPublicVisibilities(), model.TrustRootVisibility) ||
		verifierEcosystemVal0HasOverclaim(model.VerificationSummary, strings.Join(model.RequiredCaveats, " "), strings.Join(model.Limitations, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValCPublicOutputStateBlocked
	}
	audience, ok := verifierEcosystemValCAudienceLookup(audiences)[strings.TrimSpace(model.AudienceSurfaceRef)]
	if !ok || strings.TrimSpace(audience.AudienceType) != VerifierEcosystemValCAudiencePublic {
		return VerifierEcosystemValCPublicOutputStateBlocked
	}
	if !containsTrimmedString(audience.AllowedDiagnosticClasses, model.DiagnosticClass) || !containsTrimmedString(audience.AllowedOutputClasses, model.OutputClass) {
		return VerifierEcosystemValCPublicOutputStateBlocked
	}
	expectedOutputClass := verifierEcosystemValCExpectedOutputClass(model.OverallResult, model.DiagnosticClass, model.RequiredCaveats)
	if expectedOutputClass == VerifierEcosystemValBOutputClassUnknown {
		return VerifierEcosystemValCPublicOutputStateUnknown
	}
	if strings.TrimSpace(model.OutputClass) != expectedOutputClass {
		if expectedOutputClass != VerifierEcosystemValBOutputClassVerified && strings.TrimSpace(model.OutputClass) == VerifierEcosystemValBOutputClassVerified {
			return VerifierEcosystemValCPublicOutputStateBlocked
		}
		return VerifierEcosystemValCPublicOutputStatePartial
	}
	return VerifierEcosystemValCPublicOutputStateActive
}

func EvaluateVerifierEcosystemValCPartnerOutputState(model VerifierEcosystemValCPartnerOutputContract, audiences VerifierEcosystemValCAudienceSurfaceCatalog) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.OutputID,
		model.AudienceSurfaceRef,
		model.ProofFingerprint,
		model.ProofType,
		model.SchemaVersion,
		model.OverallResult,
		model.DiagnosticClass,
		model.OutputClass,
		model.CompatibilityState,
		model.FreshnessState,
		model.IssuerRefVisibility,
		model.TrustRootRefVisibility,
		model.IntegrationContext,
		model.AllowedDetailLevel,
		model.EvidenceRefPolicy,
		model.ProjectionDisclaimer,
	) || len(model.RedactedFields) == 0 || len(model.Caveats) == 0 {
		return VerifierEcosystemValCPartnerOutputStateIncomplete
	}
	if EvaluateVerifierEcosystemValCAudienceSurfaceState(audiences) != VerifierEcosystemValCAudienceSurfaceStateActive {
		return VerifierEcosystemValCPartnerOutputStateBlocked
	}
	if !containsTrimmedString(verifierEcosystemVal0SupportedProofTypes(), model.ProofType) ||
		!containsTrimmedString(VerifierEcosystemValBCompatibilityMatrixModel().SupportedSchemaVersions, model.SchemaVersion) ||
		!containsTrimmedString(verifierEcosystemValAOverallResults(), model.OverallResult) ||
		!containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), model.DiagnosticClass) ||
		!containsTrimmedString(verifierEcosystemValBOutputClasses(), model.OutputClass) ||
		!containsTrimmedString(verifierEcosystemValCFreshnessStates(), model.FreshnessState) ||
		!containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), model.CompatibilityState) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValCPartnerOutputStateUnknown
	}
	if model.InternalOnlyDiagnosticsExposed || model.MutatesEvidence || model.ApprovesDeployment || model.SuppressesFailures || model.PublishesClaims || model.CanonicalTruthClaim ||
		verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValCPartnerOutputStateBlocked
	}
	audience, ok := verifierEcosystemValCAudienceLookup(audiences)[strings.TrimSpace(model.AudienceSurfaceRef)]
	if !ok || strings.TrimSpace(audience.AudienceType) != VerifierEcosystemValCAudiencePartner {
		return VerifierEcosystemValCPartnerOutputStateBlocked
	}
	if !containsTrimmedString(audience.AllowedDiagnosticClasses, model.DiagnosticClass) || !containsTrimmedString(audience.AllowedOutputClasses, model.OutputClass) {
		return VerifierEcosystemValCPartnerOutputStateBlocked
	}
	expectedOutputClass := verifierEcosystemValCExpectedOutputClass(model.OverallResult, model.DiagnosticClass, model.Caveats)
	if expectedOutputClass == VerifierEcosystemValBOutputClassUnknown {
		return VerifierEcosystemValCPartnerOutputStateUnknown
	}
	if strings.TrimSpace(model.OutputClass) != expectedOutputClass {
		return VerifierEcosystemValCPartnerOutputStatePartial
	}
	return VerifierEcosystemValCPartnerOutputStateActive
}

func EvaluateVerifierEcosystemValCAuditorFlowState(model VerifierEcosystemValCAuditorFlowContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.AuditorFlowID,
		model.Version,
		model.VerifierContractRef,
		model.ProofEnvelopeRef,
		model.DeterministicReportRef,
		model.OutputBoundaryRef,
		model.TrustRootMaterialRef,
		model.RevocationMaterialRef,
		model.SupersessionMaterialRef,
		model.CompatibilityPolicyRef,
		model.FreshnessPolicyRef,
		model.ProjectionDisclaimer,
	) || len(model.FixtureOrCaseRefs) == 0 || len(model.RequiredEvidenceRefs) == 0 || len(model.RepeatabilityInputs) == 0 || len(model.PreservedDiagnostics) == 0 || len(model.Caveats) == 0 {
		return VerifierEcosystemValCAuditorFlowStateIncomplete
	}
	if !model.Repeatable || !model.EvidenceLinked || strings.TrimSpace(model.OutputBoundaryRef) != "boundary/auditor-safe" {
		return VerifierEcosystemValCAuditorFlowStateBlocked
	}
	if model.HiddenMainInstanceDependency || model.CertifiesOrganization || verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValCAuditorFlowStateBlocked
	}
	for _, value := range append(append([]string{}, model.RequiredEvidenceRefs...), model.RepeatabilityInputs...) {
		if strings.TrimSpace(value) == "" {
			return VerifierEcosystemValCAuditorFlowStatePartial
		}
	}
	if !containsAllTrimmedStrings(model.PreservedDiagnostics, verifierEcosystemValCRequiredPreservedDiagnostics()...) || !verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValCAuditorFlowStatePartial
	}
	return VerifierEcosystemValCAuditorFlowStateActive
}

func EvaluateVerifierEcosystemValCRequestContractState(model VerifierEcosystemValCRequestContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.RequestContractID,
		model.RequestMode,
		model.MaxScopeClass,
		model.RedactionPolicyRef,
		model.ProjectionDisclaimer,
	) || len(model.AcceptedInputDescriptors) == 0 || len(model.AllowedArtifactTypes) == 0 || len(model.RequiredMetadata) == 0 || len(model.AllowedOutputBoundaries) == 0 || len(model.TrustMaterialRequirements) == 0 || len(model.AbuseOrMisuseLimitations) == 0 || len(model.Caveats) == 0 {
		return VerifierEcosystemValCRequestContractStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemValCRequestModes(), model.RequestMode) ||
		!containsExactTrimmedStringSet(model.AcceptedInputDescriptors, VerifierEcosystemValCRequestModeUploadDescriptor, VerifierEcosystemValCRequestModeReferenceDescriptor, VerifierEcosystemValCRequestModeOfflineBundleDescriptor, VerifierEcosystemValCRequestModeAPIReferenceDescriptor) ||
		!containsExactTrimmedStringSet(model.RequiredMetadata, verifierEcosystemValCRequiredRequestMetadata()...) ||
		!containsAllTrimmedStrings(verifierEcosystemVal0SupportedProofTypes(), model.AllowedArtifactTypes...) ||
		!containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), model.MaxScopeClass) ||
		!containsAllTrimmedStrings(verifierEcosystemVal0SupportedScopeClasses(), model.AllowedOutputBoundaries...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValCRequestContractStateUnknown
	}
	if strings.TrimSpace(model.RequestMode) == VerifierEcosystemValCRequestModeUnknown {
		return VerifierEcosystemValCRequestContractStateUnknown
	}
	if model.IngestsCanonicalEvidence || model.ApprovesPublication || model.InternalArtifactForPublic || verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), strings.Join(model.AbuseOrMisuseLimitations, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValCRequestContractStateBlocked
	}
	return VerifierEcosystemValCRequestContractStateActive
}

func EvaluateVerifierEcosystemValCPublisherProfileState(model VerifierEcosystemValCPublisherCompatibilityProfile) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.PublisherProfileID,
		model.Version,
		model.PublisherType,
		model.RequiredSignaturePolicy,
		model.RequiredDigestPolicy,
		model.IssuerBindingPolicy,
		model.TrustRootPolicy,
		model.RevocationSuppressionPolicy,
		model.CompatibilityPolicyRef,
		model.ProjectionDisclaimer,
	) || len(model.SupportedArtifactTypes) == 0 || len(model.RequiredSchemaVersions) == 0 || len(model.RequiredProofEnvelopeFields) == 0 || len(model.ConformanceCaseRefs) == 0 || len(model.ForbiddenClaims) == 0 || len(model.Caveats) == 0 {
		return VerifierEcosystemValCPublisherProfileStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemValCPublisherTypes(), model.PublisherType) ||
		!containsAllTrimmedStrings(verifierEcosystemVal0SupportedProofTypes(), model.SupportedArtifactTypes...) ||
		!containsAllTrimmedStrings(VerifierEcosystemValBCompatibilityMatrixModel().SupportedSchemaVersions, model.RequiredSchemaVersions...) ||
		!containsExactTrimmedStringSet(model.RequiredProofEnvelopeFields, verifierEcosystemValCRequiredPublisherEnvelopeFields()...) ||
		!containsAllTrimmedStrings(verifierEcosystemValBRequiredConformanceCaseIDs(), model.ConformanceCaseRefs...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValCPublisherProfileStateUnknown
	}
	if strings.TrimSpace(model.PublisherType) == VerifierEcosystemValCPublisherTypeUnknown {
		return VerifierEcosystemValCPublisherProfileStateUnknown
	}
	if model.ApprovedVendorClaim || model.AutomaticallyTrustedClaim || verifierEcosystemVal0HasOverclaim(strings.Join(model.ObservedClaims, " "), strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValCPublisherProfileStateBlocked
	}
	for _, claim := range model.ObservedClaims {
		if containsTrimmedString(model.ForbiddenClaims, claim) {
			return VerifierEcosystemValCPublisherProfileStateBlocked
		}
	}
	return VerifierEcosystemValCPublisherProfileStateActive
}

func EvaluateVerifierEcosystemValCArtifactRuleState(model VerifierEcosystemValCArtifactPublishingRuleCatalog) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.RuleCatalogID, model.ProjectionDisclaimer) || len(model.Rules) == 0 {
		return VerifierEcosystemValCArtifactRuleStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedArtifactTypes, verifierEcosystemVal0SupportedProofTypes()...) ||
		!containsExactTrimmedStringSet(model.SupportedOutputBoundaries, VerifierEcosystemScopePublicSafe, VerifierEcosystemScopePartnerSafe, VerifierEcosystemScopeAuditorSafe) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValCArtifactRuleStatePartial
	}
	ruleIDs := make([]string, 0, len(model.Rules))
	for _, item := range model.Rules {
		if !referenceArchitectureValBRequiredRefsPresent(
			item.RuleID,
			item.PublisherProfileRef,
			item.ArtifactType,
			item.SchemaVersion,
			item.ProofType,
			item.CompatibilityState,
			item.ProjectionDisclaimer,
			item.SelectedOutputBoundary,
		) || len(item.RequiredFields) == 0 || len(item.ForbiddenFields) == 0 || len(item.RequiredDiagnostics) == 0 || len(item.OutputBoundaryConstraints) == 0 {
			return VerifierEcosystemValCArtifactRuleStateIncomplete
		}
		if !containsTrimmedString(model.SupportedArtifactTypes, item.ArtifactType) || !containsTrimmedString(verifierEcosystemVal0SupportedProofTypes(), item.ProofType) {
			return VerifierEcosystemValCArtifactRuleStateUnknown
		}
		if !containsTrimmedString(VerifierEcosystemValBCompatibilityMatrixModel().SupportedSchemaVersions, item.SchemaVersion) {
			return VerifierEcosystemValCArtifactRuleStateBlocked
		}
		if !containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), item.CompatibilityState) ||
			!containsAllTrimmedStrings(verifierEcosystemVal0DiagnosticClasses(), item.RequiredDiagnostics...) ||
			!containsAllTrimmedStrings(model.SupportedOutputBoundaries, item.OutputBoundaryConstraints...) ||
			!containsAllTrimmedStrings(item.ObservedFields, item.RequiredFields...) ||
			!verifierEcosystemVal0HasProjectionDisclaimer(item.ProjectionDisclaimer) {
			return VerifierEcosystemValCArtifactRuleStatePartial
		}
		if !containsTrimmedString(item.OutputBoundaryConstraints, item.SelectedOutputBoundary) {
			return VerifierEcosystemValCArtifactRuleStateBlocked
		}
		if strings.TrimSpace(item.CompatibilityState) == ReferenceArchitectureCompatibilityUnsupported || strings.TrimSpace(item.CompatibilityState) == ReferenceArchitectureCompatibilityUnknown {
			return VerifierEcosystemValCArtifactRuleStateBlocked
		}
		for _, claim := range item.ObservedClaims {
			if containsTrimmedString(item.ForbiddenFields, claim) {
				return VerifierEcosystemValCArtifactRuleStateBlocked
			}
		}
		if verifierEcosystemVal0HasOverclaim(strings.Join(item.ObservedClaims, " "), strings.Join(item.Caveats, " "), item.ProjectionDisclaimer) {
			return VerifierEcosystemValCArtifactRuleStateBlocked
		}
		ruleIDs = append(ruleIDs, item.RuleID)
	}
	if !containsExactTrimmedStringSet(ruleIDs, verifierEcosystemValCRequiredArtifactRuleIDs()...) {
		return VerifierEcosystemValCArtifactRuleStatePartial
	}
	return VerifierEcosystemValCArtifactRuleStateActive
}

func EvaluateVerifierEcosystemValCTrustDistributionState(model VerifierEcosystemValCTrustDistributionVisibility) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.TrustDistributionID,
		model.IssuerDiscoveryPolicy,
		model.TrustRootDistributionMode,
		model.TrustRootVersion,
		model.KeyRotationState,
		model.RevocationMaterialRef,
		model.SupersessionMaterialRef,
		model.AudienceVisibilityScope,
		model.ProjectionDisclaimer,
	) || len(model.Caveats) == 0 {
		return VerifierEcosystemValCTrustDistributionStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemValCDistributionModes(), model.TrustRootDistributionMode) ||
		!containsTrimmedString(verifierEcosystemVal0TrustRootStates(), model.TrustRootState) ||
		!containsTrimmedString(verifierEcosystemVal0RevocationStates(), model.RevocationState) ||
		!containsTrimmedString([]string{VerifierEcosystemKeyRotationCurrent, VerifierEcosystemKeyRotationRollover}, model.KeyRotationState) ||
		!containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), model.AudienceVisibilityScope) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValCTrustDistributionStateUnknown
	}
	if strings.TrimSpace(model.TrustRootDistributionMode) == VerifierEcosystemValCDistributionModeUnknown {
		return VerifierEcosystemValCTrustDistributionStateUnknown
	}
	if model.GlobalKeyDirectoryClaim || (strings.TrimSpace(model.AudienceVisibilityScope) == VerifierEcosystemScopePublicSafe && model.SensitiveKeyMaterialExposed) ||
		verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValCTrustDistributionStateBlocked
	}
	switch strings.TrimSpace(model.RevocationState) {
	case VerifierEcosystemRevocationNotRevoked:
	case VerifierEcosystemRevocationRevoked, VerifierEcosystemRevocationExpired, VerifierEcosystemRevocationUnsupported:
		return VerifierEcosystemValCTrustDistributionStateBlocked
	case VerifierEcosystemRevocationUnknown:
		return VerifierEcosystemValCTrustDistributionStateUnknown
	default:
		return VerifierEcosystemValCTrustDistributionStateUnknown
	}
	trustDistributionState := VerifierEcosystemValCTrustDistributionStateUnknown
	switch strings.TrimSpace(model.TrustRootState) {
	case VerifierEcosystemTrustRootTrusted:
		trustDistributionState = VerifierEcosystemValCTrustDistributionStateActive
	case VerifierEcosystemTrustRootTrustedWithWarnings:
		trustDistributionState = VerifierEcosystemValCTrustDistributionStatePartial
	case VerifierEcosystemTrustRootRotated:
		if strings.TrimSpace(model.RolloverMetadataRef) == "" {
			return VerifierEcosystemValCTrustDistributionStateBlocked
		}
		trustDistributionState = VerifierEcosystemValCTrustDistributionStatePartial
	case VerifierEcosystemTrustRootRevoked, VerifierEcosystemTrustRootExpired, VerifierEcosystemTrustRootUnsupported:
		return VerifierEcosystemValCTrustDistributionStateBlocked
	case VerifierEcosystemTrustRootUnknown:
		return VerifierEcosystemValCTrustDistributionStateUnknown
	default:
		return VerifierEcosystemValCTrustDistributionStateUnknown
	}
	if strings.TrimSpace(model.KeyRotationState) == VerifierEcosystemKeyRotationRollover {
		if strings.TrimSpace(model.RolloverMetadataRef) == "" {
			return VerifierEcosystemValCTrustDistributionStateBlocked
		}
		if trustDistributionState == VerifierEcosystemValCTrustDistributionStateActive || trustDistributionState == VerifierEcosystemValCTrustDistributionStatePartial {
			trustDistributionState = VerifierEcosystemValCTrustDistributionStatePartial
		} else {
			return trustDistributionState
		}
	}
	if strings.TrimSpace(model.TrustRootDistributionMode) == VerifierEcosystemValCDistributionModeOfflineBundle && strings.TrimSpace(model.AudienceVisibilityScope) == VerifierEcosystemScopePublicSafe {
		return VerifierEcosystemValCTrustDistributionStateBlocked
	}
	return trustDistributionState
}

func verifierEcosystemValCPoint6DependencyHealthy(snapshot VerifierEcosystemValCDependencySnapshot) bool {
	return strings.TrimSpace(snapshot.Point5State) == IntelligenceCalibrationPoint5StatePass &&
		strings.TrimSpace(snapshot.Point5DependencyState) == IntelligenceCalibrationValEStateActive &&
		strings.TrimSpace(snapshot.Point6State) == ReferenceArchitecturePoint6StatePass &&
		strings.TrimSpace(snapshot.Point6ClosureState) == ReferenceArchitectureValEStateActive &&
		strings.TrimSpace(snapshot.Point6ClosurePrerequisiteState) == ReferenceArchitectureValEPrerequisiteStateActive &&
		strings.TrimSpace(snapshot.Point6ClosureInvariantState) == ReferenceArchitectureValEInvariantStateActive &&
		strings.TrimSpace(snapshot.Point6ProofSurfaceState) == ReferenceArchitectureValEProofSurfaceStateActive &&
		strings.TrimSpace(snapshot.Point6PassRuleState) == ReferenceArchitectureValEPassRuleStateActive &&
		snapshot.Point6PassAllowed
}

func verifierEcosystemValCStateSeverity(state, active, partial, incomplete, blocked, unknown string) int {
	switch strings.TrimSpace(state) {
	case active:
		return 0
	case partial:
		return 1
	case incomplete:
		return 2
	case unknown:
		return 3
	case blocked:
		return 4
	default:
		return 3
	}
}

func EvaluateVerifierEcosystemValCState(
	dependency VerifierEcosystemValCDependencySnapshot,
	audienceSurfaceState, publicOutputState, partnerOutputState, auditorFlowState, requestContractState, publisherProfileState, artifactRuleState, trustDistributionState string,
) string {
	if !verifierEcosystemValCPoint6DependencyHealthy(dependency) ||
		strings.TrimSpace(dependency.Val0CurrentState) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(dependency.Val0State) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(dependency.ValACurrentState) != VerifierEcosystemValAStateActive ||
		strings.TrimSpace(dependency.ValAState) != VerifierEcosystemValAStateActive ||
		strings.TrimSpace(dependency.ValBCurrentState) != VerifierEcosystemValBStateActive ||
		strings.TrimSpace(dependency.ValBState) != VerifierEcosystemValBStateActive ||
		strings.TrimSpace(dependency.Point7State) != VerifierEcosystemPoint7StateNotComplete {
		return VerifierEcosystemValCStateBlocked
	}
	highestSeverity := 0
	componentSeverities := []int{
		verifierEcosystemValCStateSeverity(audienceSurfaceState, VerifierEcosystemValCAudienceSurfaceStateActive, VerifierEcosystemValCAudienceSurfaceStatePartial, VerifierEcosystemValCAudienceSurfaceStateIncomplete, VerifierEcosystemValCAudienceSurfaceStateBlocked, VerifierEcosystemValCAudienceSurfaceStateUnknown),
		verifierEcosystemValCStateSeverity(publicOutputState, VerifierEcosystemValCPublicOutputStateActive, VerifierEcosystemValCPublicOutputStatePartial, VerifierEcosystemValCPublicOutputStateIncomplete, VerifierEcosystemValCPublicOutputStateBlocked, VerifierEcosystemValCPublicOutputStateUnknown),
		verifierEcosystemValCStateSeverity(partnerOutputState, VerifierEcosystemValCPartnerOutputStateActive, VerifierEcosystemValCPartnerOutputStatePartial, VerifierEcosystemValCPartnerOutputStateIncomplete, VerifierEcosystemValCPartnerOutputStateBlocked, VerifierEcosystemValCPartnerOutputStateUnknown),
		verifierEcosystemValCStateSeverity(auditorFlowState, VerifierEcosystemValCAuditorFlowStateActive, VerifierEcosystemValCAuditorFlowStatePartial, VerifierEcosystemValCAuditorFlowStateIncomplete, VerifierEcosystemValCAuditorFlowStateBlocked, VerifierEcosystemValCAuditorFlowStateUnknown),
		verifierEcosystemValCStateSeverity(requestContractState, VerifierEcosystemValCRequestContractStateActive, VerifierEcosystemValCRequestContractStatePartial, VerifierEcosystemValCRequestContractStateIncomplete, VerifierEcosystemValCRequestContractStateBlocked, VerifierEcosystemValCRequestContractStateUnknown),
		verifierEcosystemValCStateSeverity(publisherProfileState, VerifierEcosystemValCPublisherProfileStateActive, VerifierEcosystemValCPublisherProfileStatePartial, VerifierEcosystemValCPublisherProfileStateIncomplete, VerifierEcosystemValCPublisherProfileStateBlocked, VerifierEcosystemValCPublisherProfileStateUnknown),
		verifierEcosystemValCStateSeverity(artifactRuleState, VerifierEcosystemValCArtifactRuleStateActive, VerifierEcosystemValCArtifactRuleStatePartial, VerifierEcosystemValCArtifactRuleStateIncomplete, VerifierEcosystemValCArtifactRuleStateBlocked, VerifierEcosystemValCArtifactRuleStateUnknown),
		verifierEcosystemValCStateSeverity(trustDistributionState, VerifierEcosystemValCTrustDistributionStateActive, VerifierEcosystemValCTrustDistributionStatePartial, VerifierEcosystemValCTrustDistributionStateIncomplete, VerifierEcosystemValCTrustDistributionStateBlocked, VerifierEcosystemValCTrustDistributionStateUnknown),
	}
	for _, severity := range componentSeverities {
		if severity > highestSeverity {
			highestSeverity = severity
		}
	}
	switch highestSeverity {
	case 4:
		return VerifierEcosystemValCStateBlocked
	case 3:
		return VerifierEcosystemValCStateUnknown
	case 2:
		return VerifierEcosystemValCStateIncomplete
	case 1:
		return VerifierEcosystemValCStatePartial
	default:
		return VerifierEcosystemValCStateActive
	}
}

func VerifierEcosystemValCProofSurfaceRefs() []string {
	return []string{
		"/v1/verifier-ecosystem/val0/proofs",
		"/v1/verifier-ecosystem/vala/proofs",
		"/v1/verifier-ecosystem/valb/proofs",
		"/v1/verifier-ecosystem/valc/audience-surfaces",
		"/v1/verifier-ecosystem/valc/public-output",
		"/v1/verifier-ecosystem/valc/partner-output",
		"/v1/verifier-ecosystem/valc/auditor-flow",
		"/v1/verifier-ecosystem/valc/request-contract",
		"/v1/verifier-ecosystem/valc/publisher-profile",
		"/v1/verifier-ecosystem/valc/artifact-rules",
		"/v1/verifier-ecosystem/valc/trust-distribution",
		"/v1/verifier-ecosystem/valc/proofs",
	}
}

func EvaluateVerifierEcosystemValCProofsState(
	currentState string,
	point7State string,
	val0CurrentState string,
	valACurrentState string,
	valBCurrentState string,
	surfaceRefs, evidenceRefs, limitations, whyPoint7NotPass []string,
	projectionDisclaimer string,
) string {
	baseState := strings.TrimSpace(currentState)
	if strings.TrimSpace(val0CurrentState) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(valACurrentState) != VerifierEcosystemValAStateActive ||
		strings.TrimSpace(valBCurrentState) != VerifierEcosystemValBStateActive ||
		!containsExactTrimmedStringSet(surfaceRefs, VerifierEcosystemValCProofSurfaceRefs()...) ||
		!verifierEcosystemValCProofEvidenceQualityValid(VerifierEcosystemValCVerifierEvidence(), evidenceRefs) ||
		len(limitations) == 0 ||
		len(whyPoint7NotPass) == 0 ||
		!verifierEcosystemVal0HasProjectionDisclaimer(projectionDisclaimer) {
		if baseState == VerifierEcosystemValCStateActive {
			return VerifierEcosystemValCStatePartial
		}
		return baseState
	}
	if baseState == VerifierEcosystemValCStateActive && strings.TrimSpace(point7State) != VerifierEcosystemPoint7StateNotComplete {
		return VerifierEcosystemValCStatePartial
	}
	return baseState
}

func EvaluateVerifierEcosystemValCPoint7State(valCState string) string {
	_ = valCState
	return VerifierEcosystemPoint7StateNotComplete
}

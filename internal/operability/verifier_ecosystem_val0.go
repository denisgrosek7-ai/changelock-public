package operability

import (
	"strings"
)

const (
	VerifierEcosystemPoint7StatePass        = "verifier_ecosystem_point_7_pass"
	VerifierEcosystemPoint7StateNotComplete = "verifier_ecosystem_point_7_not_complete"

	VerifierEcosystemVal0ContractStateActive     = "verifier_ecosystem_val0_contract_active"
	VerifierEcosystemVal0ContractStatePartial    = "verifier_ecosystem_val0_contract_partial"
	VerifierEcosystemVal0ContractStateIncomplete = "verifier_ecosystem_val0_contract_incomplete"
	VerifierEcosystemVal0ContractStateBlocked    = "verifier_ecosystem_val0_contract_blocked"
	VerifierEcosystemVal0ContractStateUnknown    = "verifier_ecosystem_val0_contract_unknown"

	VerifierEcosystemVal0EnvelopeStateActive     = "verifier_ecosystem_val0_proof_envelope_active"
	VerifierEcosystemVal0EnvelopeStatePartial    = "verifier_ecosystem_val0_proof_envelope_partial"
	VerifierEcosystemVal0EnvelopeStateIncomplete = "verifier_ecosystem_val0_proof_envelope_incomplete"
	VerifierEcosystemVal0EnvelopeStateBlocked    = "verifier_ecosystem_val0_proof_envelope_blocked"
	VerifierEcosystemVal0EnvelopeStateUnknown    = "verifier_ecosystem_val0_proof_envelope_unknown"

	VerifierEcosystemVal0ScopeStateActive     = "verifier_ecosystem_val0_scope_active"
	VerifierEcosystemVal0ScopeStatePartial    = "verifier_ecosystem_val0_scope_partial"
	VerifierEcosystemVal0ScopeStateIncomplete = "verifier_ecosystem_val0_scope_incomplete"
	VerifierEcosystemVal0ScopeStateBlocked    = "verifier_ecosystem_val0_scope_blocked"
	VerifierEcosystemVal0ScopeStateUnknown    = "verifier_ecosystem_val0_scope_unknown"

	VerifierEcosystemVal0CompatibilityStateActive     = "verifier_ecosystem_val0_schema_compatibility_active"
	VerifierEcosystemVal0CompatibilityStatePartial    = "verifier_ecosystem_val0_schema_compatibility_partial"
	VerifierEcosystemVal0CompatibilityStateIncomplete = "verifier_ecosystem_val0_schema_compatibility_incomplete"
	VerifierEcosystemVal0CompatibilityStateBlocked    = "verifier_ecosystem_val0_schema_compatibility_blocked"
	VerifierEcosystemVal0CompatibilityStateUnknown    = "verifier_ecosystem_val0_schema_compatibility_unknown"

	VerifierEcosystemVal0TrustStateActive     = "verifier_ecosystem_val0_trust_root_issuer_active"
	VerifierEcosystemVal0TrustStatePartial    = "verifier_ecosystem_val0_trust_root_issuer_partial"
	VerifierEcosystemVal0TrustStateIncomplete = "verifier_ecosystem_val0_trust_root_issuer_incomplete"
	VerifierEcosystemVal0TrustStateBlocked    = "verifier_ecosystem_val0_trust_root_issuer_blocked"
	VerifierEcosystemVal0TrustStateUnknown    = "verifier_ecosystem_val0_trust_root_issuer_unknown"

	VerifierEcosystemVal0DiagnosticsStateActive     = "verifier_ecosystem_val0_diagnostics_active"
	VerifierEcosystemVal0DiagnosticsStatePartial    = "verifier_ecosystem_val0_diagnostics_partial"
	VerifierEcosystemVal0DiagnosticsStateIncomplete = "verifier_ecosystem_val0_diagnostics_incomplete"
	VerifierEcosystemVal0DiagnosticsStateBlocked    = "verifier_ecosystem_val0_diagnostics_blocked"
	VerifierEcosystemVal0DiagnosticsStateUnknown    = "verifier_ecosystem_val0_diagnostics_unknown"

	VerifierEcosystemVal0OutputBoundaryStateActive     = "verifier_ecosystem_val0_output_boundaries_active"
	VerifierEcosystemVal0OutputBoundaryStatePartial    = "verifier_ecosystem_val0_output_boundaries_partial"
	VerifierEcosystemVal0OutputBoundaryStateIncomplete = "verifier_ecosystem_val0_output_boundaries_incomplete"
	VerifierEcosystemVal0OutputBoundaryStateBlocked    = "verifier_ecosystem_val0_output_boundaries_blocked"
	VerifierEcosystemVal0OutputBoundaryStateUnknown    = "verifier_ecosystem_val0_output_boundaries_unknown"

	VerifierEcosystemVal0StateActive     = "verifier_ecosystem_val0_active"
	VerifierEcosystemVal0StatePartial    = "verifier_ecosystem_val0_partial"
	VerifierEcosystemVal0StateIncomplete = "verifier_ecosystem_val0_incomplete"
	VerifierEcosystemVal0StateBlocked    = "verifier_ecosystem_val0_blocked"
	VerifierEcosystemVal0StateUnknown    = "verifier_ecosystem_val0_unknown"

	VerifierEcosystemProfilePublic           = "public"
	VerifierEcosystemProfilePartner          = "partner"
	VerifierEcosystemProfileAuditor          = "auditor"
	VerifierEcosystemProfileInternal         = "internal"
	VerifierEcosystemProfileOfflineAirGap    = "offline_or_air_gapped"
	VerifierEcosystemModeStandalone          = "standalone"
	VerifierEcosystemModeEmbeddedSDK         = "embedded_sdk"
	VerifierEcosystemModeServiceSurface      = "service_surface"
	VerifierEcosystemModeOfflineBundle       = "offline_bundle"
	VerifierEcosystemModeUnknown             = "unknown"
	VerifierEcosystemScopePublicSafe         = "public_safe"
	VerifierEcosystemScopePartnerSafe        = "partner_safe"
	VerifierEcosystemScopeAuditorSafe        = "auditor_safe"
	VerifierEcosystemScopeInternalDiagnostic = "internal_diagnostic"
	VerifierEcosystemScopeRestrictedOffline  = "restricted_offline"

	VerifierEcosystemProofTypeSignedAttestation = "signed_attestation_envelope"
	VerifierEcosystemProofTypeSealedArtifact    = "sealed_artifact_envelope"
	VerifierEcosystemProofTypeLineageBundle     = "lineage_bundle_envelope"

	VerifierEcosystemTrustModeManagedOnline = "managed_online"
	VerifierEcosystemTrustModePartnerScoped = "partner_scoped"
	VerifierEcosystemTrustModeOfflineBundle = "offline_bundle"
	VerifierEcosystemTrustModeAirGapped     = "air_gapped_bundle"

	VerifierEcosystemTrustRootTrusted             = "trusted"
	VerifierEcosystemTrustRootTrustedWithWarnings = "trusted_with_warnings"
	VerifierEcosystemTrustRootRotated             = "rotated"
	VerifierEcosystemTrustRootRevoked             = "revoked"
	VerifierEcosystemTrustRootExpired             = "expired"
	VerifierEcosystemTrustRootUnsupported         = "unsupported"
	VerifierEcosystemTrustRootUnknown             = "unknown"
	VerifierEcosystemRevocationNotRevoked         = "not_revoked"
	VerifierEcosystemRevocationRevoked            = "revoked"
	VerifierEcosystemRevocationExpired            = "expired"
	VerifierEcosystemRevocationUnsupported        = "unsupported"
	VerifierEcosystemRevocationUnknown            = "unknown"

	VerifierEcosystemKeyRotationCurrent  = "current"
	VerifierEcosystemKeyRotationRollover = "rollover_in_progress"

	VerifierEcosystemDiagnosticVerified                  = "verified"
	VerifierEcosystemDiagnosticInvalidSignature          = "invalid_signature"
	VerifierEcosystemDiagnosticDigestMismatch            = "digest_mismatch"
	VerifierEcosystemDiagnosticSchemaMismatch            = "schema_mismatch"
	VerifierEcosystemDiagnosticUnsupportedSchema         = "unsupported_schema"
	VerifierEcosystemDiagnosticUnsupportedProofType      = "unsupported_proof_type"
	VerifierEcosystemDiagnosticStaleArtifact             = "stale_artifact"
	VerifierEcosystemDiagnosticExpiredArtifact           = "expired_artifact"
	VerifierEcosystemDiagnosticRevokedIssuer             = "revoked_issuer"
	VerifierEcosystemDiagnosticSupersededProof           = "superseded_proof"
	VerifierEcosystemDiagnosticInsufficientTrustMaterial = "insufficient_trust_material"
	VerifierEcosystemDiagnosticIncompleteArtifact        = "incomplete_artifact"
	VerifierEcosystemDiagnosticScopeMismatch             = "scope_mismatch"
	VerifierEcosystemDiagnosticRedactionViolation        = "redaction_boundary_violation"
	VerifierEcosystemDiagnosticCompatibilityWarning      = "compatibility_warning"
	VerifierEcosystemDiagnosticUnknown                   = "unknown"
)

type VerifierEcosystemVal0VerifierContract struct {
	CurrentState              string   `json:"current_state"`
	VerifierContractID        string   `json:"verifier_contract_id"`
	Version                   string   `json:"version"`
	VerifierProfile           string   `json:"verifier_profile"`
	SupportedVerifierProfiles []string `json:"supported_verifier_profiles,omitempty"`
	VerifierMode              string   `json:"verifier_mode"`
	SupportedVerifierModes    []string `json:"supported_verifier_modes,omitempty"`
	SupportedProofTypes       []string `json:"supported_proof_types,omitempty"`
	SupportedSchemaVersions   []string `json:"supported_schema_versions,omitempty"`
	SupportedTrustRootModes   []string `json:"supported_trust_root_modes,omitempty"`
	SupportedOutputBoundaries []string `json:"supported_output_boundaries,omitempty"`
	RequiredArtifactFields    []string `json:"required_artifact_fields,omitempty"`
	RequiredEvidenceTypes     []string `json:"required_evidence_types,omitempty"`
	RequiredDiagnosticClasses []string `json:"required_diagnostic_classes,omitempty"`
	LifecycleState            string   `json:"lifecycle_state"`
	CompatibilityState        string   `json:"compatibility_state"`
	Owner                     string   `json:"owner"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
	CreatedAt                 string   `json:"created_at"`
	UpdatedAt                 string   `json:"updated_at"`
	CertifiedLanguagePresent  bool     `json:"certified_language_present"`
	UniversalTruthClaim       bool     `json:"universal_truth_claim"`
	RegulatorApprovedClaim    bool     `json:"regulator_approved_claim"`
	CanonicalTruthClaim       bool     `json:"canonical_truth_claim"`
}

type VerifierEcosystemVal0ProofEnvelope struct {
	CurrentState            string   `json:"current_state"`
	EnvelopeID              string   `json:"envelope_id"`
	ProofType               string   `json:"proof_type"`
	SchemaVersion           string   `json:"schema_version"`
	ArtifactDigestRef       string   `json:"artifact_digest_ref"`
	SignatureRef            string   `json:"signature_ref"`
	IssuerRef               string   `json:"issuer_ref"`
	TrustRootRef            string   `json:"trust_root_ref"`
	Scope                   string   `json:"scope"`
	IssuedAt                string   `json:"issued_at"`
	ExpiresAt               string   `json:"expires_at"`
	LineageRef              string   `json:"lineage_ref"`
	CompatibilityRef        string   `json:"compatibility_ref"`
	RevocationRef           string   `json:"revocation_ref"`
	SupersessionRef         string   `json:"supersession_ref"`
	Caveats                 []string `json:"caveats,omitempty"`
	ProjectionDisclaimer    string   `json:"projection_disclaimer"`
	ClaimsTruthOutsideScope bool     `json:"claims_truth_outside_scope"`
}

type VerifierEcosystemVal0VerificationScope struct {
	CurrentState              string   `json:"current_state"`
	ScopeID                   string   `json:"scope_id"`
	ScopeClass                string   `json:"scope_class"`
	AllowedDiagnosticDetails  []string `json:"allowed_diagnostic_details,omitempty"`
	RedactionAware            bool     `json:"redaction_aware"`
	EvidenceTraceable         bool     `json:"evidence_traceable"`
	InternalOnly              bool     `json:"internal_only"`
	OfflineTrustConstraintRef string   `json:"offline_trust_constraint_ref"`
	RequiredCaveats           []string `json:"required_caveats,omitempty"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
}

type VerifierEcosystemVal0VerificationScopeCatalog struct {
	CurrentState          string                                   `json:"current_state"`
	CatalogID             string                                   `json:"catalog_id"`
	SupportedScopeClasses []string                                 `json:"supported_scope_classes,omitempty"`
	Scopes                []VerifierEcosystemVal0VerificationScope `json:"scopes,omitempty"`
	ProjectionDisclaimer  string                                   `json:"projection_disclaimer"`
}

type VerifierEcosystemVal0SchemaCompatibilityBaseline struct {
	CurrentState                  string   `json:"current_state"`
	BaselineID                    string   `json:"baseline_id"`
	PrimarySchemaVersion          string   `json:"primary_schema_version"`
	SupportedSchemaVersions       []string `json:"supported_schema_versions,omitempty"`
	SupportedCompatibilityResults []string `json:"supported_compatibility_results,omitempty"`
	CompatibilityState            string   `json:"compatibility_state"`
	SupersessionHandlingRef       string   `json:"supersession_handling_ref"`
	MixedVersionDiagnostics       []string `json:"mixed_version_diagnostics,omitempty"`
	Caveats                       []string `json:"caveats,omitempty"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	UniversalSupportClaim         bool     `json:"universal_support_claim"`
}

type VerifierEcosystemVal0TrustIssuerDiscipline struct {
	CurrentState                 string   `json:"current_state"`
	IssuerID                     string   `json:"issuer_id"`
	IssuerScope                  string   `json:"issuer_scope"`
	TrustRootRef                 string   `json:"trust_root_ref"`
	TrustRootVersion             string   `json:"trust_root_version"`
	KeyOrMaterialRef             string   `json:"key_or_material_ref"`
	KeyRotationState             string   `json:"key_rotation_state"`
	TrustRootState               string   `json:"trust_root_state"`
	RevocationState              string   `json:"revocation_state"`
	TrustScope                   string   `json:"trust_scope"`
	OfflineDistributionSupported bool     `json:"offline_distribution_supported"`
	OfflineDistributionScope     string   `json:"offline_distribution_scope"`
	RolloverMetadataRef          string   `json:"rollover_metadata_ref"`
	Caveats                      []string `json:"caveats,omitempty"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
	GlobalKeyDirectoryClaim      bool     `json:"global_key_directory_claim"`
}

type VerifierEcosystemVal0DiagnosticsModel struct {
	CurrentState                  string   `json:"current_state"`
	DiagnosticsID                 string   `json:"diagnostics_id"`
	SupportedDiagnosticClasses    []string `json:"supported_diagnostic_classes,omitempty"`
	RequiredDiagnosticClasses     []string `json:"required_diagnostic_classes,omitempty"`
	ObservedDiagnosticClass       string   `json:"observed_diagnostic_class"`
	ReasonDetails                 []string `json:"reason_details,omitempty"`
	RequiredChecksPassed          bool     `json:"required_checks_passed"`
	EvidenceTraceable             bool     `json:"evidence_traceable"`
	RedactionKeepsFailuresVisible bool     `json:"redaction_keeps_failures_visible"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	CertifiedLanguagePresent      bool     `json:"certified_language_present"`
	UniversalTruthClaim           bool     `json:"universal_truth_claim"`
	CanonicalTruthClaim           bool     `json:"canonical_truth_claim"`
}

type VerifierEcosystemVal0OutputBoundary struct {
	CurrentState                string   `json:"current_state"`
	BoundaryID                  string   `json:"boundary_id"`
	ScopeClass                  string   `json:"scope_class"`
	AllowedFields               []string `json:"allowed_fields,omitempty"`
	RedactedFields              []string `json:"redacted_fields,omitempty"`
	RequiredCaveats             []string `json:"required_caveats,omitempty"`
	RequiredDiagnostics         []string `json:"required_diagnostics,omitempty"`
	EvidenceRefPolicy           string   `json:"evidence_ref_policy"`
	TrustMaterialVisibility     string   `json:"trust_material_visibility"`
	ExportAllowed               bool     `json:"export_allowed"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
	PublicReuseAllowed          bool     `json:"public_reuse_allowed"`
	PreservesInvalidDiagnostics bool     `json:"preserves_invalid_diagnostics"`
}

type VerifierEcosystemVal0OutputBoundaryCollection struct {
	CurrentState          string                                `json:"current_state"`
	CollectionID          string                                `json:"collection_id"`
	SupportedScopeClasses []string                              `json:"supported_scope_classes,omitempty"`
	Boundaries            []VerifierEcosystemVal0OutputBoundary `json:"boundaries,omitempty"`
	ProjectionDisclaimer  string                                `json:"projection_disclaimer"`
}

type VerifierEcosystemVal0DependencySnapshot struct {
	Point5State                    string `json:"point_5_state"`
	Point5DependencyState          string `json:"point_5_dependency_state"`
	Point6State                    string `json:"point_6_state"`
	Point6ClosureState             string `json:"point_6_closure_state"`
	Point6ClosurePrerequisiteState string `json:"point_6_closure_prerequisite_state"`
	Point6ClosureInvariantState    string `json:"point_6_closure_invariant_state"`
	Point6ProofSurfaceState        string `json:"point_6_proof_surface_state"`
	Point6PassRuleState            string `json:"point_6_pass_rule_state"`
	Point6PassAllowed              bool   `json:"point_6_pass_allowed"`
}

func verifierEcosystemVal0SupportedProfiles() []string {
	return []string{
		VerifierEcosystemProfilePublic,
		VerifierEcosystemProfilePartner,
		VerifierEcosystemProfileAuditor,
		VerifierEcosystemProfileInternal,
		VerifierEcosystemProfileOfflineAirGap,
	}
}

func verifierEcosystemVal0SupportedModes() []string {
	return []string{
		VerifierEcosystemModeStandalone,
		VerifierEcosystemModeEmbeddedSDK,
		VerifierEcosystemModeServiceSurface,
		VerifierEcosystemModeOfflineBundle,
		VerifierEcosystemModeUnknown,
	}
}

func verifierEcosystemVal0SupportedProofTypes() []string {
	return []string{
		VerifierEcosystemProofTypeSignedAttestation,
		VerifierEcosystemProofTypeSealedArtifact,
		VerifierEcosystemProofTypeLineageBundle,
	}
}

func verifierEcosystemVal0SupportedTrustRootModes() []string {
	return []string{
		VerifierEcosystemTrustModeManagedOnline,
		VerifierEcosystemTrustModePartnerScoped,
		VerifierEcosystemTrustModeOfflineBundle,
		VerifierEcosystemTrustModeAirGapped,
	}
}

func verifierEcosystemVal0SupportedScopeClasses() []string {
	return []string{
		VerifierEcosystemScopePublicSafe,
		VerifierEcosystemScopePartnerSafe,
		VerifierEcosystemScopeAuditorSafe,
		VerifierEcosystemScopeInternalDiagnostic,
		VerifierEcosystemScopeRestrictedOffline,
	}
}

func verifierEcosystemVal0SupportedSchemaVersions() []string {
	return []string{
		"changelock.verifier.proof_envelope.v1",
		"changelock.verifier.proof_envelope.v1.1",
	}
}

func verifierEcosystemVal0CompatibilityResults() []string {
	return []string{
		ReferenceArchitectureCompatibilityCompatible,
		ReferenceArchitectureCompatibilityCompatibleWithWarning,
		ReferenceArchitectureCompatibilityDeprecated,
		ReferenceArchitectureCompatibilitySuperseded,
		ReferenceArchitectureCompatibilityUnsupported,
		ReferenceArchitectureCompatibilityUnknown,
	}
}

func verifierEcosystemVal0TrustRootStates() []string {
	return []string{
		VerifierEcosystemTrustRootTrusted,
		VerifierEcosystemTrustRootTrustedWithWarnings,
		VerifierEcosystemTrustRootRotated,
		VerifierEcosystemTrustRootRevoked,
		VerifierEcosystemTrustRootExpired,
		VerifierEcosystemTrustRootUnsupported,
		VerifierEcosystemTrustRootUnknown,
	}
}

func verifierEcosystemVal0RevocationStates() []string {
	return []string{
		VerifierEcosystemRevocationNotRevoked,
		VerifierEcosystemRevocationRevoked,
		VerifierEcosystemRevocationExpired,
		VerifierEcosystemRevocationUnsupported,
		VerifierEcosystemRevocationUnknown,
	}
}

func verifierEcosystemVal0DiagnosticClasses() []string {
	return []string{
		VerifierEcosystemDiagnosticVerified,
		VerifierEcosystemDiagnosticInvalidSignature,
		VerifierEcosystemDiagnosticDigestMismatch,
		VerifierEcosystemDiagnosticSchemaMismatch,
		VerifierEcosystemDiagnosticUnsupportedSchema,
		VerifierEcosystemDiagnosticUnsupportedProofType,
		VerifierEcosystemDiagnosticStaleArtifact,
		VerifierEcosystemDiagnosticExpiredArtifact,
		VerifierEcosystemDiagnosticRevokedIssuer,
		VerifierEcosystemDiagnosticSupersededProof,
		VerifierEcosystemDiagnosticInsufficientTrustMaterial,
		VerifierEcosystemDiagnosticIncompleteArtifact,
		VerifierEcosystemDiagnosticScopeMismatch,
		VerifierEcosystemDiagnosticRedactionViolation,
		VerifierEcosystemDiagnosticCompatibilityWarning,
		VerifierEcosystemDiagnosticUnknown,
	}
}

func verifierEcosystemVal0RequiredArtifactFields() []string {
	return []string{
		"envelope_id",
		"proof_type",
		"schema_version",
		"artifact_digest_ref",
		"signature_ref",
		"issuer_ref",
		"trust_root_ref",
		"scope",
		"issued_at",
		"expires_at",
		"lineage_ref",
		"compatibility_ref",
		"revocation_ref",
		"supersession_ref",
	}
}

func verifierEcosystemVal0RequiredEvidenceTypes() []string {
	return []string{
		"signature_material",
		"schema_definition",
		"issuer_metadata",
		"trust_root_metadata",
		"compatibility_metadata",
		"revocation_metadata",
	}
}

func verifierEcosystemVal0OutputBoundaries() []string {
	return verifierEcosystemVal0SupportedScopeClasses()
}

func verifierEcosystemVal0ProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_verifier_discipline_foundation advisory_projection"
}

func verifierEcosystemVal0HasProjectionDisclaimer(value string) bool {
	return strings.Contains(strings.TrimSpace(value), "projection_only") &&
		strings.Contains(strings.TrimSpace(value), "not_canonical_truth")
}

func verifierEcosystemVal0HasOverclaim(values ...string) bool {
	disallowed := []string{
		"anyone can verify everything",
		"mathematically proves total truth",
		"universal trust protocol",
		"changelock certification",
		"integrity rating",
		"regulator-approved verifier",
		"global key registry for all instances",
		"zero-knowledge verifier as core requirement",
		"formal certification",
		"absolute proof",
		"universal authority",
		"certified architecture",
		"guaranteed secure architecture",
		"regulator approved",
	}
	for _, value := range values {
		lower := strings.ToLower(strings.TrimSpace(value))
		for _, item := range disallowed {
			if strings.Contains(lower, item) {
				return true
			}
		}
	}
	return false
}

func verifierEcosystemVal0EvidenceRefs() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:verifier-contract-001", EvidenceType: "schema_definition", Source: "verifier/contracts", Timestamp: "2026-04-27T07:00:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "verifier_contract", Caveats: []string{"bounded to verifier discipline foundation"}},
		{EvidenceID: "evidence:proof-envelope-001", EvidenceType: "signature_material", Source: "sealed-artifact/reference-pack", Timestamp: "2026-04-27T07:01:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "proof_envelope", Caveats: []string{"bounded to declared proof envelope schema line"}},
		{EvidenceID: "evidence:verification-scope-001", EvidenceType: "scope_catalog", Source: "verifier/scopes", Timestamp: "2026-04-27T07:01:30Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "verification_scope", Caveats: []string{"bounded to declared verification scope classes and redaction discipline"}},
		{EvidenceID: "evidence:trust-root-001", EvidenceType: "trust_root_metadata", Source: "trust-root/catalog", Timestamp: "2026-04-27T07:02:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "trust_material", Caveats: []string{"bounded to versioned trust-root discovery discipline"}},
		{EvidenceID: "evidence:revocation-001", EvidenceType: "revocation_metadata", Source: "revocation/metadata", Timestamp: "2026-04-27T07:03:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "revocation_supersession", Caveats: []string{"bounded to current bounded revocation metadata window"}},
		{EvidenceID: "evidence:compatibility-001", EvidenceType: "compatibility_metadata", Source: "compatibility/baseline", Timestamp: "2026-04-27T07:04:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "schema_compatibility", Caveats: []string{"bounded to supported verifier schema line set"}},
		{EvidenceID: "evidence:diagnostics-001", EvidenceType: "diagnostics_contract", Source: "verifier/diagnostics", Timestamp: "2026-04-27T07:04:30Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "diagnostics_model", Caveats: []string{"bounded to deterministic diagnostic classes and precedence rules"}},
		{EvidenceID: "evidence:output-boundary-001", EvidenceType: "output_boundary_policy", Source: "verifier/output-boundaries", Timestamp: "2026-04-27T07:05:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "output_boundary_discipline", Caveats: []string{"bounded to public, partner, auditor, internal, and restricted offline output rules"}},
		{EvidenceID: "evidence:point7-governance-001", EvidenceType: "state_governance", Source: "verifier/point7-governance", Timestamp: "2026-04-27T07:05:30Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point7_governance", Caveats: []string{"bounded to point_7_state not complete and no-pass rule for Val 0"}},
	}
}

func VerifierEcosystemVal0VerifierEvidence() []ReferenceArchitectureEvidenceReference {
	return verifierEcosystemVal0EvidenceRefs()
}

func verifierEcosystemVal0RequiredVerifierEvidenceIDs() []string {
	return []string{
		"evidence:verifier-contract-001",
		"evidence:proof-envelope-001",
		"evidence:verification-scope-001",
		"evidence:trust-root-001",
		"evidence:revocation-001",
		"evidence:compatibility-001",
		"evidence:diagnostics-001",
		"evidence:output-boundary-001",
		"evidence:point7-governance-001",
	}
}

func verifierEcosystemVal0RequiredVerifierEvidenceScopes() []string {
	return []string{
		"verifier_contract",
		"proof_envelope",
		"verification_scope",
		"trust_material",
		"revocation_supersession",
		"schema_compatibility",
		"diagnostics_model",
		"output_boundary_discipline",
		"point7_governance",
	}
}

func VerifierEcosystemVal0ProofEvidenceRefs() []string {
	return []string{
		"point6_integrated_closure",
		"verifier_discipline_foundation",
		"evidence:verifier-contract-001",
		"evidence:proof-envelope-001",
		"evidence:verification-scope-001",
		"evidence:trust-root-001",
		"evidence:revocation-001",
		"evidence:compatibility-001",
		"evidence:diagnostics-001",
		"evidence:output-boundary-001",
		"evidence:point7-governance-001",
	}
}

func verifierEcosystemVal0EvidenceValid(evidenceRefs []ReferenceArchitectureEvidenceReference) (allFresh bool, stale bool, ok bool) {
	if len(evidenceRefs) == 0 {
		return false, false, false
	}
	allFresh = true
	for _, evidence := range evidenceRefs {
		if strings.TrimSpace(evidence.EvidenceID) == "" ||
			strings.TrimSpace(evidence.EvidenceType) == "" ||
			strings.TrimSpace(evidence.Source) == "" ||
			strings.TrimSpace(evidence.Scope) == "" {
			return false, false, false
		}
		if _, valid := referenceArchitectureVal0ParseTimestamp(evidence.Timestamp); !valid {
			return false, false, false
		}
		switch strings.TrimSpace(evidence.FreshnessState) {
		case IntelligenceCalibrationFreshnessFresh:
		case IntelligenceCalibrationFreshnessStale, IntelligenceCalibrationFreshnessExpired:
			allFresh = false
			stale = true
		case IntelligenceCalibrationFreshnessUnsupported, IntelligenceCalibrationFreshnessUnknown:
			return false, false, false
		default:
			return false, false, false
		}
	}
	return allFresh, stale, true
}

func verifierEcosystemVal0ProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, evidenceRefs []string) bool {
	if !containsExactTrimmedStringSet(evidenceRefs, VerifierEcosystemVal0ProofEvidenceRefs()...) {
		return false
	}
	allFresh, stale, ok := verifierEcosystemVal0EvidenceValid(evidence)
	if !ok || !allFresh || stale {
		return false
	}
	evidenceIDs := make([]string, 0, len(evidence))
	evidenceScopes := make([]string, 0, len(evidence))
	for _, evidenceRef := range evidence {
		evidenceIDs = append(evidenceIDs, evidenceRef.EvidenceID)
		evidenceScopes = append(evidenceScopes, evidenceRef.Scope)
	}
	return containsExactTrimmedStringSet(evidenceIDs, verifierEcosystemVal0RequiredVerifierEvidenceIDs()...) &&
		containsExactTrimmedStringSet(evidenceScopes, verifierEcosystemVal0RequiredVerifierEvidenceScopes()...)
}

func VerifierEcosystemVal0VerifierContractModel() VerifierEcosystemVal0VerifierContract {
	return VerifierEcosystemVal0VerifierContract{
		CurrentState:              "verifier_ecosystem_val0_contract_ready",
		VerifierContractID:        "verifier-contract-val0-foundation",
		Version:                   "2026.04",
		VerifierProfile:           VerifierEcosystemProfileAuditor,
		SupportedVerifierProfiles: verifierEcosystemVal0SupportedProfiles(),
		VerifierMode:              VerifierEcosystemModeServiceSurface,
		SupportedVerifierModes:    verifierEcosystemVal0SupportedModes(),
		SupportedProofTypes:       verifierEcosystemVal0SupportedProofTypes(),
		SupportedSchemaVersions:   verifierEcosystemVal0SupportedSchemaVersions(),
		SupportedTrustRootModes:   verifierEcosystemVal0SupportedTrustRootModes(),
		SupportedOutputBoundaries: verifierEcosystemVal0OutputBoundaries(),
		RequiredArtifactFields:    verifierEcosystemVal0RequiredArtifactFields(),
		RequiredEvidenceTypes:     verifierEcosystemVal0RequiredEvidenceTypes(),
		RequiredDiagnosticClasses: verifierEcosystemVal0DiagnosticClasses(),
		LifecycleState:            ReferenceArchitectureLifecycleActive,
		CompatibilityState:        ReferenceArchitectureCompatibilityCompatible,
		Owner:                     "verifier_ecosystem_program",
		ProjectionDisclaimer:      verifierEcosystemVal0ProjectionDisclaimer(),
		CreatedAt:                 "2026-04-27T07:00:00Z",
		UpdatedAt:                 "2026-04-27T07:05:00Z",
	}
}

func VerifierEcosystemVal0ProofEnvelopeModel() VerifierEcosystemVal0ProofEnvelope {
	return VerifierEcosystemVal0ProofEnvelope{
		CurrentState:         "verifier_ecosystem_val0_proof_envelope_ready",
		EnvelopeID:           "proof-envelope-val0-reference",
		ProofType:            VerifierEcosystemProofTypeSignedAttestation,
		SchemaVersion:        "changelock.verifier.proof_envelope.v1",
		ArtifactDigestRef:    "digest:sha256:reference-artifact",
		SignatureRef:         "signature:reference-artifact",
		IssuerRef:            "issuer:reference-signer",
		TrustRootRef:         "trust-root:reference-program",
		Scope:                VerifierEcosystemScopeAuditorSafe,
		IssuedAt:             "2026-04-27T07:05:00Z",
		ExpiresAt:            "2026-05-27T07:05:00Z",
		LineageRef:           "lineage:reference-blueprint",
		CompatibilityRef:     "compatibility:reference-proof-envelope",
		RevocationRef:        "revocation:reference-proof-envelope",
		SupersessionRef:      "supersession:reference-proof-envelope",
		Caveats:              []string{"cryptographic and semantic verification remains bounded to the proof envelope scope"},
		ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer(),
	}
}

func VerifierEcosystemVal0VerificationScopeCatalogModel() VerifierEcosystemVal0VerificationScopeCatalog {
	return VerifierEcosystemVal0VerificationScopeCatalog{
		CurrentState:          "verifier_ecosystem_val0_scope_catalog_ready",
		CatalogID:             "verifier-scope-catalog-val0",
		SupportedScopeClasses: verifierEcosystemVal0SupportedScopeClasses(),
		Scopes: []VerifierEcosystemVal0VerificationScope{
			{CurrentState: "verifier_scope_ready", ScopeID: "scope/public-safe", ScopeClass: VerifierEcosystemScopePublicSafe, AllowedDiagnosticDetails: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticStaleArtifact, VerifierEcosystemDiagnosticUnsupportedSchema}, RedactionAware: true, EvidenceTraceable: true, RequiredCaveats: []string{"public output is redaction-aware and bounded to artifact scope"}, ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer()},
			{CurrentState: "verifier_scope_ready", ScopeID: "scope/partner-safe", ScopeClass: VerifierEcosystemScopePartnerSafe, AllowedDiagnosticDetails: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticScopeMismatch}, RedactionAware: true, EvidenceTraceable: true, RequiredCaveats: []string{"partner output remains bounded and scoped to agreed verifier contract"}, ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer()},
			{CurrentState: "verifier_scope_ready", ScopeID: "scope/auditor-safe", ScopeClass: VerifierEcosystemScopeAuditorSafe, AllowedDiagnosticDetails: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticSupersededProof}, RedactionAware: true, EvidenceTraceable: true, RequiredCaveats: []string{"auditor output remains repeatable and evidence-linked"}, ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer()},
			{CurrentState: "verifier_scope_ready", ScopeID: "scope/internal-diagnostic", ScopeClass: VerifierEcosystemScopeInternalDiagnostic, AllowedDiagnosticDetails: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticIncompleteArtifact, VerifierEcosystemDiagnosticInvalidSignature}, RedactionAware: true, EvidenceTraceable: true, InternalOnly: true, RequiredCaveats: []string{"internal diagnostic output must not be reused as public-safe output"}, ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer()},
			{CurrentState: "verifier_scope_ready", ScopeID: "scope/restricted-offline", ScopeClass: VerifierEcosystemScopeRestrictedOffline, AllowedDiagnosticDetails: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticInsufficientTrustMaterial, VerifierEcosystemDiagnosticStaleArtifact}, RedactionAware: true, EvidenceTraceable: true, OfflineTrustConstraintRef: "offline-trust/distribution-reference", RequiredCaveats: []string{"offline verification remains bounded to scoped trust-root material and air-gapped distribution rules"}, ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer()},
		},
		ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer(),
	}
}

func VerifierEcosystemVal0SchemaCompatibilityBaselineModel() VerifierEcosystemVal0SchemaCompatibilityBaseline {
	return VerifierEcosystemVal0SchemaCompatibilityBaseline{
		CurrentState:                  "verifier_ecosystem_val0_schema_compatibility_ready",
		BaselineID:                    "verifier-schema-compatibility-val0",
		PrimarySchemaVersion:          "changelock.verifier.proof_envelope.v1",
		SupportedSchemaVersions:       verifierEcosystemVal0SupportedSchemaVersions(),
		SupportedCompatibilityResults: verifierEcosystemVal0CompatibilityResults(),
		CompatibilityState:            ReferenceArchitectureCompatibilityCompatible,
		SupersessionHandlingRef:       "supersession-handling/verifier-envelope",
		MixedVersionDiagnostics:       []string{"mixed-version replay remains representable through compatibility_warning without implying universal compatibility"},
		Caveats:                       []string{"mixed-version diagnostics remain bounded and require explicit caveat handling"},
		ProjectionDisclaimer:          verifierEcosystemVal0ProjectionDisclaimer(),
	}
}

func VerifierEcosystemVal0TrustIssuerDisciplineModel() VerifierEcosystemVal0TrustIssuerDiscipline {
	return VerifierEcosystemVal0TrustIssuerDiscipline{
		CurrentState:                 "verifier_ecosystem_val0_trust_issuer_ready",
		IssuerID:                     "issuer:reference-signer",
		IssuerScope:                  VerifierEcosystemScopeAuditorSafe,
		TrustRootRef:                 "trust-root:reference-program",
		TrustRootVersion:             "2026.04",
		KeyOrMaterialRef:             "key-material:reference-program",
		KeyRotationState:             VerifierEcosystemKeyRotationCurrent,
		TrustRootState:               VerifierEcosystemTrustRootTrusted,
		RevocationState:              VerifierEcosystemRevocationNotRevoked,
		TrustScope:                   VerifierEcosystemScopeAuditorSafe,
		OfflineDistributionSupported: true,
		OfflineDistributionScope:     "restricted_offline trust-root bundle for bounded air-gapped verification",
		RolloverMetadataRef:          "rollover:reference-program",
		Caveats:                      []string{"issuer discovery remains bounded to versioned trust-root material and scoped rollover metadata"},
		ProjectionDisclaimer:         verifierEcosystemVal0ProjectionDisclaimer(),
	}
}

func VerifierEcosystemVal0DiagnosticsCatalogModel() VerifierEcosystemVal0DiagnosticsModel {
	return VerifierEcosystemVal0DiagnosticsModel{
		CurrentState:                  "verifier_ecosystem_val0_diagnostics_ready",
		DiagnosticsID:                 "verifier-diagnostics-val0",
		SupportedDiagnosticClasses:    verifierEcosystemVal0DiagnosticClasses(),
		RequiredDiagnosticClasses:     verifierEcosystemVal0DiagnosticClasses(),
		ObservedDiagnosticClass:       VerifierEcosystemDiagnosticVerified,
		ReasonDetails:                 []string{"digest integrity, signature reference, issuer linkage, trust-root compatibility, freshness, and compatibility checks remain within declared verifier scope"},
		RequiredChecksPassed:          true,
		EvidenceTraceable:             true,
		RedactionKeepsFailuresVisible: true,
		ProjectionDisclaimer:          verifierEcosystemVal0ProjectionDisclaimer(),
	}
}

func VerifierEcosystemVal0OutputBoundaryCollectionModel() VerifierEcosystemVal0OutputBoundaryCollection {
	return VerifierEcosystemVal0OutputBoundaryCollection{
		CurrentState:          "verifier_ecosystem_val0_output_boundaries_ready",
		CollectionID:          "verifier-output-boundaries-val0",
		SupportedScopeClasses: verifierEcosystemVal0SupportedScopeClasses(),
		Boundaries: []VerifierEcosystemVal0OutputBoundary{
			{CurrentState: "verifier_output_boundary_ready", BoundaryID: "boundary/public-safe", ScopeClass: VerifierEcosystemScopePublicSafe, AllowedFields: []string{"artifact_digest_ref", "schema_version", "compatibility_state", "observed_diagnostic_class"}, RedactedFields: []string{"issuer_ref", "trust_root_ref", "lineage_ref"}, RequiredCaveats: []string{"public-safe output remains redacted and bounded to verifier scope"}, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticStaleArtifact, VerifierEcosystemDiagnosticUnsupportedSchema}, EvidenceRefPolicy: "public_caveated", TrustMaterialVisibility: "summary_only", ExportAllowed: true, ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer(), PublicReuseAllowed: true, PreservesInvalidDiagnostics: true},
			{CurrentState: "verifier_output_boundary_ready", BoundaryID: "boundary/partner-safe", ScopeClass: VerifierEcosystemScopePartnerSafe, AllowedFields: []string{"artifact_digest_ref", "issuer_ref", "schema_version", "compatibility_state", "scope"}, RedactedFields: []string{"key_or_material_ref"}, RequiredCaveats: []string{"partner output remains scoped and bounded to agreed verifier audience"}, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticScopeMismatch}, EvidenceRefPolicy: "partner_scoped", TrustMaterialVisibility: "issuer_scoped", ExportAllowed: true, ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer(), PublicReuseAllowed: false, PreservesInvalidDiagnostics: true},
			{CurrentState: "verifier_output_boundary_ready", BoundaryID: "boundary/auditor-safe", ScopeClass: VerifierEcosystemScopeAuditorSafe, AllowedFields: []string{"artifact_digest_ref", "issuer_ref", "trust_root_ref", "schema_version", "scope", "lineage_ref"}, RedactedFields: []string{}, RequiredCaveats: []string{"auditor output remains repeatable and evidence-linked"}, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticCompatibilityWarning, VerifierEcosystemDiagnosticSupersededProof}, EvidenceRefPolicy: "auditor_evidence_linked", TrustMaterialVisibility: "auditor_scoped", ExportAllowed: true, ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer(), PublicReuseAllowed: false, PreservesInvalidDiagnostics: true},
			{CurrentState: "verifier_output_boundary_ready", BoundaryID: "boundary/internal-diagnostic", ScopeClass: VerifierEcosystemScopeInternalDiagnostic, AllowedFields: []string{"artifact_digest_ref", "issuer_ref", "trust_root_ref", "schema_version", "scope", "lineage_ref", "revocation_ref", "supersession_ref"}, RedactedFields: []string{}, RequiredCaveats: []string{"internal diagnostic output must not be reused as public-safe output"}, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticIncompleteArtifact, VerifierEcosystemDiagnosticInvalidSignature}, EvidenceRefPolicy: "internal_only", TrustMaterialVisibility: "internal_scoped", ExportAllowed: false, ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer(), PublicReuseAllowed: false, PreservesInvalidDiagnostics: true},
			{CurrentState: "verifier_output_boundary_ready", BoundaryID: "boundary/restricted-offline", ScopeClass: VerifierEcosystemScopeRestrictedOffline, AllowedFields: []string{"artifact_digest_ref", "schema_version", "scope", "trust_root_ref"}, RedactedFields: []string{"issuer_ref", "lineage_ref"}, RequiredCaveats: []string{"offline output remains bounded to scoped air-gapped trust material"}, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticVerified, VerifierEcosystemDiagnosticInsufficientTrustMaterial, VerifierEcosystemDiagnosticStaleArtifact}, EvidenceRefPolicy: "offline_scoped", TrustMaterialVisibility: "offline_scoped", ExportAllowed: false, ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer(), PublicReuseAllowed: false, PreservesInvalidDiagnostics: true},
		},
		ProjectionDisclaimer: verifierEcosystemVal0ProjectionDisclaimer(),
	}
}

func EvaluateVerifierEcosystemVal0VerifierContractState(contract VerifierEcosystemVal0VerifierContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		contract.VerifierContractID,
		contract.Version,
		contract.VerifierProfile,
		contract.VerifierMode,
		contract.LifecycleState,
		contract.CompatibilityState,
		contract.Owner,
		contract.ProjectionDisclaimer,
		contract.CreatedAt,
		contract.UpdatedAt,
	) {
		return VerifierEcosystemVal0ContractStateIncomplete
	}
	if !containsExactTrimmedStringSet(contract.SupportedVerifierProfiles, verifierEcosystemVal0SupportedProfiles()...) ||
		!containsExactTrimmedStringSet(contract.SupportedVerifierModes, verifierEcosystemVal0SupportedModes()...) ||
		!containsExactTrimmedStringSet(contract.SupportedProofTypes, verifierEcosystemVal0SupportedProofTypes()...) ||
		!containsExactTrimmedStringSet(contract.SupportedSchemaVersions, verifierEcosystemVal0SupportedSchemaVersions()...) ||
		!containsExactTrimmedStringSet(contract.SupportedTrustRootModes, verifierEcosystemVal0SupportedTrustRootModes()...) ||
		!containsExactTrimmedStringSet(contract.SupportedOutputBoundaries, verifierEcosystemVal0OutputBoundaries()...) ||
		!containsExactTrimmedStringSet(contract.RequiredArtifactFields, verifierEcosystemVal0RequiredArtifactFields()...) ||
		!containsExactTrimmedStringSet(contract.RequiredEvidenceTypes, verifierEcosystemVal0RequiredEvidenceTypes()...) ||
		!containsExactTrimmedStringSet(contract.RequiredDiagnosticClasses, verifierEcosystemVal0DiagnosticClasses()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(contract.ProjectionDisclaimer) {
		return VerifierEcosystemVal0ContractStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(contract.CreatedAt); !ok {
		return VerifierEcosystemVal0ContractStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(contract.UpdatedAt); !ok {
		return VerifierEcosystemVal0ContractStatePartial
	}
	if !containsTrimmedString(verifierEcosystemVal0SupportedProfiles(), contract.VerifierProfile) ||
		!containsTrimmedString(verifierEcosystemVal0SupportedModes(), contract.VerifierMode) ||
		!containsTrimmedString(referenceArchitectureVal0LifecycleStates(), contract.LifecycleState) ||
		!containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), contract.CompatibilityState) {
		return VerifierEcosystemVal0ContractStateUnknown
	}
	if contract.CertifiedLanguagePresent || contract.UniversalTruthClaim || contract.RegulatorApprovedClaim || contract.CanonicalTruthClaim ||
		verifierEcosystemVal0HasOverclaim(contract.ProjectionDisclaimer) {
		return VerifierEcosystemVal0ContractStateBlocked
	}
	if strings.TrimSpace(contract.CompatibilityState) == ReferenceArchitectureCompatibilityUnsupported {
		return VerifierEcosystemVal0ContractStateBlocked
	}
	if strings.TrimSpace(contract.CompatibilityState) == ReferenceArchitectureCompatibilityUnknown ||
		strings.TrimSpace(contract.VerifierMode) == VerifierEcosystemModeUnknown {
		return VerifierEcosystemVal0ContractStateUnknown
	}
	if strings.TrimSpace(contract.CompatibilityState) == ReferenceArchitectureCompatibilityCompatibleWithWarning ||
		strings.TrimSpace(contract.CompatibilityState) == ReferenceArchitectureCompatibilityDeprecated ||
		strings.TrimSpace(contract.CompatibilityState) == ReferenceArchitectureCompatibilitySuperseded {
		return VerifierEcosystemVal0ContractStatePartial
	}
	return VerifierEcosystemVal0ContractStateActive
}

func EvaluateVerifierEcosystemVal0ProofEnvelopeState(envelope VerifierEcosystemVal0ProofEnvelope) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		envelope.EnvelopeID,
		envelope.ProofType,
		envelope.SchemaVersion,
		envelope.ArtifactDigestRef,
		envelope.SignatureRef,
		envelope.IssuerRef,
		envelope.TrustRootRef,
		envelope.Scope,
		envelope.IssuedAt,
		envelope.ExpiresAt,
		envelope.LineageRef,
		envelope.CompatibilityRef,
		envelope.RevocationRef,
		envelope.SupersessionRef,
		envelope.ProjectionDisclaimer,
	) {
		return VerifierEcosystemVal0EnvelopeStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemVal0SupportedProofTypes(), envelope.ProofType) ||
		!containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), envelope.Scope) {
		return VerifierEcosystemVal0EnvelopeStateUnknown
	}
	if !containsTrimmedString(verifierEcosystemVal0SupportedSchemaVersions(), envelope.SchemaVersion) {
		return VerifierEcosystemVal0EnvelopeStateBlocked
	}
	issuedAt, issuedOK := referenceArchitectureVal0ParseTimestamp(envelope.IssuedAt)
	expiresAt, expiresOK := referenceArchitectureVal0ParseTimestamp(envelope.ExpiresAt)
	if !issuedOK || !expiresOK {
		return VerifierEcosystemVal0EnvelopeStatePartial
	}
	if !expiresAt.After(issuedAt) || strings.TrimSpace(envelope.ExpiresAt) == "" {
		return VerifierEcosystemVal0EnvelopeStateBlocked
	}
	if envelope.ClaimsTruthOutsideScope || verifierEcosystemVal0HasOverclaim(strings.Join(envelope.Caveats, " ")) {
		return VerifierEcosystemVal0EnvelopeStateBlocked
	}
	if !verifierEcosystemVal0HasProjectionDisclaimer(envelope.ProjectionDisclaimer) {
		return VerifierEcosystemVal0EnvelopeStatePartial
	}
	return VerifierEcosystemVal0EnvelopeStateActive
}

func EvaluateVerifierEcosystemVal0VerificationScopeState(catalog VerifierEcosystemVal0VerificationScopeCatalog) string {
	if !referenceArchitectureValBRequiredRefsPresent(catalog.CatalogID, catalog.ProjectionDisclaimer) || len(catalog.Scopes) == 0 {
		return VerifierEcosystemVal0ScopeStateIncomplete
	}
	if !containsExactTrimmedStringSet(catalog.SupportedScopeClasses, verifierEcosystemVal0SupportedScopeClasses()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(catalog.ProjectionDisclaimer) ||
		len(catalog.Scopes) != len(verifierEcosystemVal0SupportedScopeClasses()) {
		return VerifierEcosystemVal0ScopeStatePartial
	}
	seen := map[string]struct{}{}
	for _, scope := range catalog.Scopes {
		normalized := strings.TrimSpace(scope.ScopeClass)
		if normalized == "" || !containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), normalized) {
			return VerifierEcosystemVal0ScopeStateUnknown
		}
		if _, duplicate := seen[normalized]; duplicate {
			return VerifierEcosystemVal0ScopeStatePartial
		}
		seen[normalized] = struct{}{}
		if !referenceArchitectureValBRequiredRefsPresent(scope.ScopeID, scope.ScopeClass, scope.ProjectionDisclaimer) || len(scope.AllowedDiagnosticDetails) == 0 || len(scope.RequiredCaveats) == 0 {
			return VerifierEcosystemVal0ScopeStateIncomplete
		}
		if !containsAllTrimmedStrings(verifierEcosystemVal0DiagnosticClasses(), scope.AllowedDiagnosticDetails...) || !verifierEcosystemVal0HasProjectionDisclaimer(scope.ProjectionDisclaimer) {
			return VerifierEcosystemVal0ScopeStatePartial
		}
		switch normalized {
		case VerifierEcosystemScopePublicSafe:
			if !scope.RedactionAware {
				return VerifierEcosystemVal0ScopeStateBlocked
			}
		case VerifierEcosystemScopeAuditorSafe:
			if !scope.EvidenceTraceable {
				return VerifierEcosystemVal0ScopeStateBlocked
			}
		case VerifierEcosystemScopeInternalDiagnostic:
			if !scope.InternalOnly {
				return VerifierEcosystemVal0ScopeStateBlocked
			}
		case VerifierEcosystemScopeRestrictedOffline:
			if strings.TrimSpace(scope.OfflineTrustConstraintRef) == "" {
				return VerifierEcosystemVal0ScopeStateBlocked
			}
		}
	}
	if !containsExactTrimmedStringSet(keysFromMap(seen), verifierEcosystemVal0SupportedScopeClasses()...) {
		return VerifierEcosystemVal0ScopeStatePartial
	}
	return VerifierEcosystemVal0ScopeStateActive
}

func EvaluateVerifierEcosystemVal0SchemaCompatibilityBaselineState(baseline VerifierEcosystemVal0SchemaCompatibilityBaseline) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		baseline.BaselineID,
		baseline.PrimarySchemaVersion,
		baseline.CompatibilityState,
		baseline.ProjectionDisclaimer,
	) {
		return VerifierEcosystemVal0CompatibilityStateIncomplete
	}
	if !containsExactTrimmedStringSet(baseline.SupportedSchemaVersions, verifierEcosystemVal0SupportedSchemaVersions()...) ||
		!containsExactTrimmedStringSet(baseline.SupportedCompatibilityResults, verifierEcosystemVal0CompatibilityResults()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(baseline.ProjectionDisclaimer) {
		return VerifierEcosystemVal0CompatibilityStatePartial
	}
	if !containsTrimmedString(verifierEcosystemVal0SupportedSchemaVersions(), baseline.PrimarySchemaVersion) ||
		!containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), baseline.CompatibilityState) {
		return VerifierEcosystemVal0CompatibilityStateUnknown
	}
	if baseline.UniversalSupportClaim || verifierEcosystemVal0HasOverclaim(strings.Join(baseline.Caveats, " ")) {
		return VerifierEcosystemVal0CompatibilityStateBlocked
	}
	switch strings.TrimSpace(baseline.CompatibilityState) {
	case ReferenceArchitectureCompatibilityCompatible:
		return VerifierEcosystemVal0CompatibilityStateActive
	case ReferenceArchitectureCompatibilityCompatibleWithWarning:
		if len(baseline.Caveats) == 0 {
			return VerifierEcosystemVal0CompatibilityStateBlocked
		}
		return VerifierEcosystemVal0CompatibilityStatePartial
	case ReferenceArchitectureCompatibilityDeprecated, ReferenceArchitectureCompatibilitySuperseded:
		if strings.TrimSpace(baseline.SupersessionHandlingRef) == "" {
			return VerifierEcosystemVal0CompatibilityStateBlocked
		}
		return VerifierEcosystemVal0CompatibilityStatePartial
	case ReferenceArchitectureCompatibilityUnsupported:
		return VerifierEcosystemVal0CompatibilityStateBlocked
	case ReferenceArchitectureCompatibilityUnknown:
		return VerifierEcosystemVal0CompatibilityStateUnknown
	default:
		return VerifierEcosystemVal0CompatibilityStateUnknown
	}
}

func EvaluateVerifierEcosystemVal0TrustIssuerDisciplineState(model VerifierEcosystemVal0TrustIssuerDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.IssuerID,
		model.IssuerScope,
		model.TrustRootRef,
		model.TrustRootVersion,
		model.KeyOrMaterialRef,
		model.KeyRotationState,
		model.TrustRootState,
		model.RevocationState,
		model.TrustScope,
		model.ProjectionDisclaimer,
	) {
		return VerifierEcosystemVal0TrustStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), model.IssuerScope) ||
		!containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), model.TrustScope) ||
		!containsTrimmedString(verifierEcosystemVal0TrustRootStates(), model.TrustRootState) ||
		!containsTrimmedString(verifierEcosystemVal0RevocationStates(), model.RevocationState) ||
		!containsTrimmedString([]string{VerifierEcosystemKeyRotationCurrent, VerifierEcosystemKeyRotationRollover}, model.KeyRotationState) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemVal0TrustStateUnknown
	}
	if model.GlobalKeyDirectoryClaim || verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " ")) {
		return VerifierEcosystemVal0TrustStateBlocked
	}
	if model.OfflineDistributionSupported && strings.TrimSpace(model.OfflineDistributionScope) == "" {
		return VerifierEcosystemVal0TrustStateBlocked
	}
	switch strings.TrimSpace(model.RevocationState) {
	case VerifierEcosystemRevocationNotRevoked:
	case VerifierEcosystemRevocationRevoked, VerifierEcosystemRevocationExpired, VerifierEcosystemRevocationUnsupported:
		return VerifierEcosystemVal0TrustStateBlocked
	case VerifierEcosystemRevocationUnknown:
		return VerifierEcosystemVal0TrustStateUnknown
	default:
		return VerifierEcosystemVal0TrustStateUnknown
	}
	if strings.TrimSpace(model.KeyRotationState) == VerifierEcosystemKeyRotationRollover {
		if strings.TrimSpace(model.RolloverMetadataRef) == "" {
			return VerifierEcosystemVal0TrustStateBlocked
		}
		return VerifierEcosystemVal0TrustStatePartial
	}
	switch strings.TrimSpace(model.TrustRootState) {
	case VerifierEcosystemTrustRootTrusted:
		return VerifierEcosystemVal0TrustStateActive
	case VerifierEcosystemTrustRootTrustedWithWarnings:
		return VerifierEcosystemVal0TrustStatePartial
	case VerifierEcosystemTrustRootRotated:
		if strings.TrimSpace(model.RolloverMetadataRef) == "" {
			return VerifierEcosystemVal0TrustStateBlocked
		}
		return VerifierEcosystemVal0TrustStatePartial
	case VerifierEcosystemTrustRootRevoked, VerifierEcosystemTrustRootExpired, VerifierEcosystemTrustRootUnsupported:
		return VerifierEcosystemVal0TrustStateBlocked
	case VerifierEcosystemTrustRootUnknown:
		return VerifierEcosystemVal0TrustStateUnknown
	default:
		return VerifierEcosystemVal0TrustStateUnknown
	}
}

func EvaluateVerifierEcosystemVal0DiagnosticsState(model VerifierEcosystemVal0DiagnosticsModel) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DiagnosticsID, model.ObservedDiagnosticClass, model.ProjectionDisclaimer) || len(model.ReasonDetails) == 0 {
		return VerifierEcosystemVal0DiagnosticsStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedDiagnosticClasses, verifierEcosystemVal0DiagnosticClasses()...) ||
		!containsExactTrimmedStringSet(model.RequiredDiagnosticClasses, verifierEcosystemVal0DiagnosticClasses()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemVal0DiagnosticsStatePartial
	}
	if !containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), model.ObservedDiagnosticClass) {
		return VerifierEcosystemVal0DiagnosticsStateUnknown
	}
	if model.CertifiedLanguagePresent || model.UniversalTruthClaim || model.CanonicalTruthClaim || verifierEcosystemVal0HasOverclaim(strings.Join(model.ReasonDetails, " ")) {
		return VerifierEcosystemVal0DiagnosticsStateBlocked
	}
	if !model.EvidenceTraceable || !model.RedactionKeepsFailuresVisible {
		return VerifierEcosystemVal0DiagnosticsStateBlocked
	}
	if !model.RequiredChecksPassed {
		return VerifierEcosystemVal0DiagnosticsStateBlocked
	}
	if strings.TrimSpace(model.ObservedDiagnosticClass) != VerifierEcosystemDiagnosticVerified {
		switch strings.TrimSpace(model.ObservedDiagnosticClass) {
		case VerifierEcosystemDiagnosticUnknown:
			return VerifierEcosystemVal0DiagnosticsStateUnknown
		case VerifierEcosystemDiagnosticCompatibilityWarning:
			return VerifierEcosystemVal0DiagnosticsStatePartial
		default:
			return VerifierEcosystemVal0DiagnosticsStateBlocked
		}
	}
	return VerifierEcosystemVal0DiagnosticsStateActive
}

func EvaluateVerifierEcosystemVal0OutputBoundaryState(collection VerifierEcosystemVal0OutputBoundaryCollection) string {
	if !referenceArchitectureValBRequiredRefsPresent(collection.CollectionID, collection.ProjectionDisclaimer) || len(collection.Boundaries) == 0 {
		return VerifierEcosystemVal0OutputBoundaryStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedScopeClasses, verifierEcosystemVal0SupportedScopeClasses()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Boundaries) != len(verifierEcosystemVal0SupportedScopeClasses()) {
		return VerifierEcosystemVal0OutputBoundaryStatePartial
	}
	seen := map[string]struct{}{}
	for _, boundary := range collection.Boundaries {
		scopeClass := strings.TrimSpace(boundary.ScopeClass)
		if scopeClass == "" || !containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), scopeClass) {
			return VerifierEcosystemVal0OutputBoundaryStateUnknown
		}
		if _, duplicate := seen[scopeClass]; duplicate {
			return VerifierEcosystemVal0OutputBoundaryStatePartial
		}
		seen[scopeClass] = struct{}{}
		if !referenceArchitectureValBRequiredRefsPresent(boundary.BoundaryID, boundary.ScopeClass, boundary.EvidenceRefPolicy, boundary.TrustMaterialVisibility, boundary.ProjectionDisclaimer) ||
			len(boundary.AllowedFields) == 0 || len(boundary.RequiredCaveats) == 0 || len(boundary.RequiredDiagnostics) == 0 {
			return VerifierEcosystemVal0OutputBoundaryStateIncomplete
		}
		if !containsAllTrimmedStrings(verifierEcosystemVal0DiagnosticClasses(), boundary.RequiredDiagnostics...) || !verifierEcosystemVal0HasProjectionDisclaimer(boundary.ProjectionDisclaimer) {
			return VerifierEcosystemVal0OutputBoundaryStatePartial
		}
		if !boundary.PreservesInvalidDiagnostics {
			return VerifierEcosystemVal0OutputBoundaryStateBlocked
		}
		switch scopeClass {
		case VerifierEcosystemScopePublicSafe:
			if len(boundary.RedactedFields) == 0 {
				return VerifierEcosystemVal0OutputBoundaryStateBlocked
			}
		case VerifierEcosystemScopeAuditorSafe:
			if strings.TrimSpace(boundary.EvidenceRefPolicy) != "auditor_evidence_linked" {
				return VerifierEcosystemVal0OutputBoundaryStateBlocked
			}
		case VerifierEcosystemScopeInternalDiagnostic:
			if boundary.PublicReuseAllowed {
				return VerifierEcosystemVal0OutputBoundaryStateBlocked
			}
		}
	}
	if !containsExactTrimmedStringSet(keysFromMap(seen), verifierEcosystemVal0SupportedScopeClasses()...) {
		return VerifierEcosystemVal0OutputBoundaryStatePartial
	}
	return VerifierEcosystemVal0OutputBoundaryStateActive
}

func verifierEcosystemVal0Point6DependencyHealthy(snapshot VerifierEcosystemVal0DependencySnapshot) bool {
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

func verifierEcosystemVal0AggregateStateSeverity(state, active, partial, incomplete, blocked, unknown string) int {
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

func EvaluateVerifierEcosystemVal0State(
	dependency VerifierEcosystemVal0DependencySnapshot,
	contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState string,
) string {
	if !verifierEcosystemVal0Point6DependencyHealthy(dependency) {
		return VerifierEcosystemVal0StateBlocked
	}
	highestSeverity := 0
	componentSeverities := []int{
		verifierEcosystemVal0AggregateStateSeverity(
			contractState,
			VerifierEcosystemVal0ContractStateActive,
			VerifierEcosystemVal0ContractStatePartial,
			VerifierEcosystemVal0ContractStateIncomplete,
			VerifierEcosystemVal0ContractStateBlocked,
			VerifierEcosystemVal0ContractStateUnknown,
		),
		verifierEcosystemVal0AggregateStateSeverity(
			envelopeState,
			VerifierEcosystemVal0EnvelopeStateActive,
			VerifierEcosystemVal0EnvelopeStatePartial,
			VerifierEcosystemVal0EnvelopeStateIncomplete,
			VerifierEcosystemVal0EnvelopeStateBlocked,
			VerifierEcosystemVal0EnvelopeStateUnknown,
		),
		verifierEcosystemVal0AggregateStateSeverity(
			scopeState,
			VerifierEcosystemVal0ScopeStateActive,
			VerifierEcosystemVal0ScopeStatePartial,
			VerifierEcosystemVal0ScopeStateIncomplete,
			VerifierEcosystemVal0ScopeStateBlocked,
			VerifierEcosystemVal0ScopeStateUnknown,
		),
		verifierEcosystemVal0AggregateStateSeverity(
			compatibilityState,
			VerifierEcosystemVal0CompatibilityStateActive,
			VerifierEcosystemVal0CompatibilityStatePartial,
			VerifierEcosystemVal0CompatibilityStateIncomplete,
			VerifierEcosystemVal0CompatibilityStateBlocked,
			VerifierEcosystemVal0CompatibilityStateUnknown,
		),
		verifierEcosystemVal0AggregateStateSeverity(
			trustState,
			VerifierEcosystemVal0TrustStateActive,
			VerifierEcosystemVal0TrustStatePartial,
			VerifierEcosystemVal0TrustStateIncomplete,
			VerifierEcosystemVal0TrustStateBlocked,
			VerifierEcosystemVal0TrustStateUnknown,
		),
		verifierEcosystemVal0AggregateStateSeverity(
			diagnosticsState,
			VerifierEcosystemVal0DiagnosticsStateActive,
			VerifierEcosystemVal0DiagnosticsStatePartial,
			VerifierEcosystemVal0DiagnosticsStateIncomplete,
			VerifierEcosystemVal0DiagnosticsStateBlocked,
			VerifierEcosystemVal0DiagnosticsStateUnknown,
		),
		verifierEcosystemVal0AggregateStateSeverity(
			outputBoundaryState,
			VerifierEcosystemVal0OutputBoundaryStateActive,
			VerifierEcosystemVal0OutputBoundaryStatePartial,
			VerifierEcosystemVal0OutputBoundaryStateIncomplete,
			VerifierEcosystemVal0OutputBoundaryStateBlocked,
			VerifierEcosystemVal0OutputBoundaryStateUnknown,
		),
	}
	for _, severity := range componentSeverities {
		if severity > highestSeverity {
			highestSeverity = severity
		}
	}
	switch highestSeverity {
	case 4:
		return VerifierEcosystemVal0StateBlocked
	case 3:
		return VerifierEcosystemVal0StateUnknown
	case 2:
		return VerifierEcosystemVal0StateIncomplete
	case 1:
		return VerifierEcosystemVal0StatePartial
	default:
		return VerifierEcosystemVal0StateActive
	}
}

func VerifierEcosystemVal0ProofSurfaceRefs() []string {
	return []string{
		"/v1/verifier-ecosystem/val0/contract",
		"/v1/verifier-ecosystem/val0/proof-envelope",
		"/v1/verifier-ecosystem/val0/verification-scope",
		"/v1/verifier-ecosystem/val0/schema-compatibility",
		"/v1/verifier-ecosystem/val0/trust-root-issuer",
		"/v1/verifier-ecosystem/val0/diagnostics",
		"/v1/verifier-ecosystem/val0/output-boundaries",
		"/v1/verifier-ecosystem/val0/proofs",
	}
}

func EvaluateVerifierEcosystemVal0ProofsState(
	currentState string,
	point7State string,
	supportedProfiles, supportedModes, surfaceRefs, evidenceRefs, limitations []string,
	projectionDisclaimer string,
) string {
	baseState := strings.TrimSpace(currentState)
	if !containsExactTrimmedStringSet(supportedProfiles, verifierEcosystemVal0SupportedProfiles()...) ||
		!containsExactTrimmedStringSet(supportedModes, verifierEcosystemVal0SupportedModes()...) ||
		!containsExactTrimmedStringSet(surfaceRefs, VerifierEcosystemVal0ProofSurfaceRefs()...) ||
		!verifierEcosystemVal0ProofEvidenceQualityValid(VerifierEcosystemVal0VerifierEvidence(), evidenceRefs) ||
		len(limitations) == 0 ||
		!verifierEcosystemVal0HasProjectionDisclaimer(projectionDisclaimer) {
		if baseState == VerifierEcosystemVal0StateActive {
			return VerifierEcosystemVal0StatePartial
		}
		return baseState
	}
	if baseState == VerifierEcosystemVal0StateActive && strings.TrimSpace(point7State) != VerifierEcosystemPoint7StateNotComplete {
		return VerifierEcosystemVal0StatePartial
	}
	return baseState
}

func EvaluateVerifierEcosystemPoint7State(val0State string) string {
	if strings.TrimSpace(val0State) == VerifierEcosystemVal0StateActive {
		return VerifierEcosystemPoint7StateNotComplete
	}
	return VerifierEcosystemPoint7StateNotComplete
}

func keysFromMap(values map[string]struct{}) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	return keys
}

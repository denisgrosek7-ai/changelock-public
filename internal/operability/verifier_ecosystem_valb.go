package operability

import "strings"

const (
	VerifierEcosystemValBCompatibilityMatrixStateActive     = "verifier_ecosystem_valb_compatibility_matrix_active"
	VerifierEcosystemValBCompatibilityMatrixStatePartial    = "verifier_ecosystem_valb_compatibility_matrix_partial"
	VerifierEcosystemValBCompatibilityMatrixStateIncomplete = "verifier_ecosystem_valb_compatibility_matrix_incomplete"
	VerifierEcosystemValBCompatibilityMatrixStateBlocked    = "verifier_ecosystem_valb_compatibility_matrix_blocked"
	VerifierEcosystemValBCompatibilityMatrixStateUnknown    = "verifier_ecosystem_valb_compatibility_matrix_unknown"

	VerifierEcosystemValBSchemaProofCompatibilityStateActive     = "verifier_ecosystem_valb_schema_proof_compatibility_active"
	VerifierEcosystemValBSchemaProofCompatibilityStatePartial    = "verifier_ecosystem_valb_schema_proof_compatibility_partial"
	VerifierEcosystemValBSchemaProofCompatibilityStateIncomplete = "verifier_ecosystem_valb_schema_proof_compatibility_incomplete"
	VerifierEcosystemValBSchemaProofCompatibilityStateBlocked    = "verifier_ecosystem_valb_schema_proof_compatibility_blocked"
	VerifierEcosystemValBSchemaProofCompatibilityStateUnknown    = "verifier_ecosystem_valb_schema_proof_compatibility_unknown"

	VerifierEcosystemValBMixedVersionStateActive     = "verifier_ecosystem_valb_mixed_version_diagnostics_active"
	VerifierEcosystemValBMixedVersionStatePartial    = "verifier_ecosystem_valb_mixed_version_diagnostics_partial"
	VerifierEcosystemValBMixedVersionStateIncomplete = "verifier_ecosystem_valb_mixed_version_diagnostics_incomplete"
	VerifierEcosystemValBMixedVersionStateBlocked    = "verifier_ecosystem_valb_mixed_version_diagnostics_blocked"
	VerifierEcosystemValBMixedVersionStateUnknown    = "verifier_ecosystem_valb_mixed_version_diagnostics_unknown"

	VerifierEcosystemValBDiagnosticPrecedenceStateActive     = "verifier_ecosystem_valb_diagnostic_precedence_active"
	VerifierEcosystemValBDiagnosticPrecedenceStatePartial    = "verifier_ecosystem_valb_diagnostic_precedence_partial"
	VerifierEcosystemValBDiagnosticPrecedenceStateIncomplete = "verifier_ecosystem_valb_diagnostic_precedence_incomplete"
	VerifierEcosystemValBDiagnosticPrecedenceStateBlocked    = "verifier_ecosystem_valb_diagnostic_precedence_blocked"
	VerifierEcosystemValBDiagnosticPrecedenceStateUnknown    = "verifier_ecosystem_valb_diagnostic_precedence_unknown"

	VerifierEcosystemValBFixtureDescriptorStateActive     = "verifier_ecosystem_valb_fixture_descriptors_active"
	VerifierEcosystemValBFixtureDescriptorStatePartial    = "verifier_ecosystem_valb_fixture_descriptors_partial"
	VerifierEcosystemValBFixtureDescriptorStateIncomplete = "verifier_ecosystem_valb_fixture_descriptors_incomplete"
	VerifierEcosystemValBFixtureDescriptorStateBlocked    = "verifier_ecosystem_valb_fixture_descriptors_blocked"
	VerifierEcosystemValBFixtureDescriptorStateUnknown    = "verifier_ecosystem_valb_fixture_descriptors_unknown"

	VerifierEcosystemValBConformanceCaseStateActive     = "verifier_ecosystem_valb_conformance_cases_active"
	VerifierEcosystemValBConformanceCaseStatePartial    = "verifier_ecosystem_valb_conformance_cases_partial"
	VerifierEcosystemValBConformanceCaseStateIncomplete = "verifier_ecosystem_valb_conformance_cases_incomplete"
	VerifierEcosystemValBConformanceCaseStateBlocked    = "verifier_ecosystem_valb_conformance_cases_blocked"
	VerifierEcosystemValBConformanceCaseStateUnknown    = "verifier_ecosystem_valb_conformance_cases_unknown"

	VerifierEcosystemValBConformanceSuiteStateActive     = "verifier_ecosystem_valb_conformance_suite_active"
	VerifierEcosystemValBConformanceSuiteStatePartial    = "verifier_ecosystem_valb_conformance_suite_partial"
	VerifierEcosystemValBConformanceSuiteStateIncomplete = "verifier_ecosystem_valb_conformance_suite_incomplete"
	VerifierEcosystemValBConformanceSuiteStateBlocked    = "verifier_ecosystem_valb_conformance_suite_blocked"
	VerifierEcosystemValBConformanceSuiteStateUnknown    = "verifier_ecosystem_valb_conformance_suite_unknown"

	VerifierEcosystemValBOutputClassStateActive     = "verifier_ecosystem_valb_output_classes_active"
	VerifierEcosystemValBOutputClassStatePartial    = "verifier_ecosystem_valb_output_classes_partial"
	VerifierEcosystemValBOutputClassStateIncomplete = "verifier_ecosystem_valb_output_classes_incomplete"
	VerifierEcosystemValBOutputClassStateBlocked    = "verifier_ecosystem_valb_output_classes_blocked"
	VerifierEcosystemValBOutputClassStateUnknown    = "verifier_ecosystem_valb_output_classes_unknown"

	VerifierEcosystemValBStateActive     = "verifier_ecosystem_valb_active"
	VerifierEcosystemValBStatePartial    = "verifier_ecosystem_valb_partial"
	VerifierEcosystemValBStateIncomplete = "verifier_ecosystem_valb_incomplete"
	VerifierEcosystemValBStateBlocked    = "verifier_ecosystem_valb_blocked"
	VerifierEcosystemValBStateUnknown    = "verifier_ecosystem_valb_unknown"

	VerifierEcosystemValBOutputClassVerified                     = "verified"
	VerifierEcosystemValBOutputClassVerifiedWithWarnings         = "verified_with_warnings"
	VerifierEcosystemValBOutputClassNonVerifiedInvalid           = "non_verified_invalid"
	VerifierEcosystemValBOutputClassNonVerifiedIncomplete        = "non_verified_incomplete"
	VerifierEcosystemValBOutputClassNonVerifiedUnsupported       = "non_verified_unsupported"
	VerifierEcosystemValBOutputClassNonVerifiedStale             = "non_verified_stale"
	VerifierEcosystemValBOutputClassNonVerifiedRevoked           = "non_verified_revoked"
	VerifierEcosystemValBOutputClassNonVerifiedSuperseded        = "non_verified_superseded"
	VerifierEcosystemValBOutputClassNonVerifiedScopeMismatch     = "non_verified_scope_mismatch"
	VerifierEcosystemValBOutputClassNonVerifiedTrustInsufficient = "non_verified_trust_material_insufficient"
	VerifierEcosystemValBOutputClassRedactionBlocked             = "redaction_blocked"
	VerifierEcosystemValBOutputClassUnknown                      = "unknown"
	VerifierEcosystemValBFixtureTypeValidProofEnvelope           = "valid_proof_envelope"
	VerifierEcosystemValBFixtureTypeStaleProofEnvelope           = "stale_proof_envelope"
	VerifierEcosystemValBFixtureTypeExpiredProofEnvelope         = "expired_proof_envelope"
	VerifierEcosystemValBFixtureTypeRevokedIssuer                = "revoked_issuer"
	VerifierEcosystemValBFixtureTypeSupersededProof              = "superseded_proof"
	VerifierEcosystemValBFixtureTypeUnsupportedSchema            = "unsupported_schema"
	VerifierEcosystemValBFixtureTypeUnsupportedProofType         = "unsupported_proof_type"
	VerifierEcosystemValBFixtureTypeMalformedTimestamp           = "malformed_timestamp"
	VerifierEcosystemValBFixtureTypeMissingSignature             = "missing_signature"
	VerifierEcosystemValBFixtureTypeMissingDigest                = "missing_digest"
	VerifierEcosystemValBFixtureTypeDigestMismatch               = "digest_mismatch"
	VerifierEcosystemValBFixtureTypeInvalidSignature             = "invalid_signature"
	VerifierEcosystemValBFixtureTypeInsufficientTrustMaterial    = "insufficient_trust_material"
	VerifierEcosystemValBFixtureTypeScopeMismatch                = "scope_mismatch"
	VerifierEcosystemValBFixtureTypeRedactionBoundaryViolation   = "redaction_boundary_violation"
)

type VerifierEcosystemValBDependencySnapshot struct {
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
	Point7State                    string `json:"point_7_state"`
}

type VerifierEcosystemValBCompatibilityEntry struct {
	CurrentState               string   `json:"current_state"`
	EntryID                    string   `json:"entry_id"`
	SchemaVersion              string   `json:"schema_version"`
	ProofType                  string   `json:"proof_type"`
	VerifierVersion            string   `json:"verifier_version"`
	TrustRootVersion           string   `json:"trust_root_version"`
	CompatibilityState         string   `json:"compatibility_state"`
	RequiredDiagnostics        []string `json:"required_diagnostics,omitempty"`
	Caveats                    []string `json:"caveats,omitempty"`
	MigrationOrSupersessionRef string   `json:"migration_or_supersession_ref"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValBCompatibilityMatrix struct {
	CurrentState               string                                    `json:"current_state"`
	CompatibilityMatrixID      string                                    `json:"compatibility_matrix_id"`
	Version                    string                                    `json:"version"`
	SupportedSchemaVersions    []string                                  `json:"supported_schema_versions,omitempty"`
	SupportedProofTypes        []string                                  `json:"supported_proof_types,omitempty"`
	SupportedVerifierVersions  []string                                  `json:"supported_verifier_versions,omitempty"`
	SupportedTrustRootVersions []string                                  `json:"supported_trust_root_versions,omitempty"`
	CompatibilityEntries       []VerifierEcosystemValBCompatibilityEntry `json:"compatibility_entries,omitempty"`
	DeprecatedVersions         []string                                  `json:"deprecated_versions,omitempty"`
	SupersededVersions         []string                                  `json:"superseded_versions,omitempty"`
	UnsupportedVersions        []string                                  `json:"unsupported_versions,omitempty"`
	MixedVersionRules          []string                                  `json:"mixed_version_rules,omitempty"`
	Caveats                    []string                                  `json:"caveats,omitempty"`
	ProjectionDisclaimer       string                                    `json:"projection_disclaimer"`
	CreatedAt                  string                                    `json:"created_at"`
	UpdatedAt                  string                                    `json:"updated_at"`
}

type VerifierEcosystemValBSchemaProofCompatibility struct {
	CurrentState               string   `json:"current_state"`
	CompatibilityID            string   `json:"compatibility_id"`
	CompatibilityMatrixRef     string   `json:"compatibility_matrix_ref"`
	SchemaVersion              string   `json:"schema_version"`
	ProofType                  string   `json:"proof_type"`
	VerifierVersion            string   `json:"verifier_version"`
	TrustRootVersion           string   `json:"trust_root_version"`
	CompatibilityState         string   `json:"compatibility_state"`
	DerivedDiagnosticClass     string   `json:"derived_diagnostic_class"`
	RequiredDiagnostics        []string `json:"required_diagnostics,omitempty"`
	MigrationOrSupersessionRef string   `json:"migration_or_supersession_ref"`
	Caveats                    []string `json:"caveats,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValBMixedVersionCase struct {
	CurrentState                string   `json:"current_state"`
	MixedVersionCaseID          string   `json:"mixed_version_case_id"`
	SchemaVersion               string   `json:"schema_version"`
	ProofType                   string   `json:"proof_type"`
	VerifierVersion             string   `json:"verifier_version"`
	TrustRootVersion            string   `json:"trust_root_version"`
	ExpectedCompatibilityState  string   `json:"expected_compatibility_state"`
	ExpectedDiagnosticClass     string   `json:"expected_diagnostic_class"`
	MigrationOrSupersessionRef  string   `json:"migration_or_supersession_ref"`
	Caveats                     []string `json:"caveats,omitempty"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
	UniversalCompatibilityClaim bool     `json:"universal_compatibility_claim"`
}

type VerifierEcosystemValBMixedVersionDiagnosticsCatalog struct {
	CurrentState                 string                                  `json:"current_state"`
	CatalogID                    string                                  `json:"catalog_id"`
	SupportedCompatibilityStates []string                                `json:"supported_compatibility_states,omitempty"`
	SupportedDiagnosticClasses   []string                                `json:"supported_diagnostic_classes,omitempty"`
	Cases                        []VerifierEcosystemValBMixedVersionCase `json:"cases,omitempty"`
	ProjectionDisclaimer         string                                  `json:"projection_disclaimer"`
}

type VerifierEcosystemValBDiagnosticPrecedence struct {
	CurrentState               string   `json:"current_state"`
	PrecedenceID               string   `json:"precedence_id"`
	SupportedDiagnosticClasses []string `json:"supported_diagnostic_classes,omitempty"`
	PrecedenceOrder            []string `json:"precedence_order,omitempty"`
	ObservedDiagnostics        []string `json:"observed_diagnostics,omitempty"`
	DerivedDiagnosticClass     string   `json:"derived_diagnostic_class"`
	Caveats                    []string `json:"caveats,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValBFixtureDescriptor struct {
	CurrentState            string   `json:"current_state"`
	FixtureID               string   `json:"fixture_id"`
	FixtureType             string   `json:"fixture_type"`
	ProofType               string   `json:"proof_type"`
	SchemaVersion           string   `json:"schema_version"`
	ExpectedResult          string   `json:"expected_result"`
	ExpectedDiagnostic      string   `json:"expected_diagnostic"`
	ExpectedOutputBoundary  string   `json:"expected_output_boundary"`
	RequiredEvidenceRefs    []string `json:"required_evidence_refs,omitempty"`
	Caveats                 []string `json:"caveats,omitempty"`
	ProjectionDisclaimer    string   `json:"projection_disclaimer"`
	ProductionEvidenceClaim bool     `json:"production_evidence_claim"`
}

type VerifierEcosystemValBFixtureCatalog struct {
	CurrentState          string                                   `json:"current_state"`
	FixtureCatalogID      string                                   `json:"fixture_catalog_id"`
	SupportedFixtureTypes []string                                 `json:"supported_fixture_types,omitempty"`
	Fixtures              []VerifierEcosystemValBFixtureDescriptor `json:"fixtures,omitempty"`
	ProjectionDisclaimer  string                                   `json:"projection_disclaimer"`
}

type VerifierEcosystemValBConformanceCase struct {
	CurrentState            string   `json:"current_state"`
	ConformanceCaseID       string   `json:"conformance_case_id"`
	FixtureRef              string   `json:"fixture_ref"`
	VerifierContractRef     string   `json:"verifier_contract_ref"`
	InputRef                string   `json:"input_ref"`
	ExpectedOverallResult   string   `json:"expected_overall_result"`
	ExpectedDiagnosticClass string   `json:"expected_diagnostic_class"`
	ExpectedOutputClass     string   `json:"expected_output_class"`
	RequiredFields          []string `json:"required_fields,omitempty"`
	ForbiddenClaims         []string `json:"forbidden_claims,omitempty"`
	ObservedClaims          []string `json:"observed_claims,omitempty"`
	Caveats                 []string `json:"caveats,omitempty"`
	ProjectionDisclaimer    string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValBConformanceCaseCatalog struct {
	CurrentState         string                                 `json:"current_state"`
	ConformanceCatalogID string                                 `json:"conformance_catalog_id"`
	Cases                []VerifierEcosystemValBConformanceCase `json:"cases,omitempty"`
	ProjectionDisclaimer string                                 `json:"projection_disclaimer"`
}

type VerifierEcosystemValBConformanceSuite struct {
	CurrentState               string   `json:"current_state"`
	ConformanceSuiteID         string   `json:"conformance_suite_id"`
	RequiredCaseRefs           []string `json:"required_case_refs,omitempty"`
	RequiredFixtureRefs        []string `json:"required_fixture_refs,omitempty"`
	SupportedOutputClasses     []string `json:"supported_output_classes,omitempty"`
	SupportedDiagnosticClasses []string `json:"supported_diagnostic_classes,omitempty"`
	Caveats                    []string `json:"caveats,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
	CertificationClaim         bool     `json:"certification_claim"`
}

type VerifierEcosystemValBOutputClassMapping struct {
	CurrentState         string   `json:"current_state"`
	OutputClassID        string   `json:"output_class_id"`
	OverallResult        string   `json:"overall_result"`
	DiagnosticClass      string   `json:"diagnostic_class"`
	OutputClass          string   `json:"output_class"`
	Caveats              []string `json:"caveats,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValBOutputClassCatalog struct {
	CurrentState           string                                    `json:"current_state"`
	OutputClassCatalogID   string                                    `json:"output_class_catalog_id"`
	SupportedOutputClasses []string                                  `json:"supported_output_classes,omitempty"`
	Mappings               []VerifierEcosystemValBOutputClassMapping `json:"mappings,omitempty"`
	ProjectionDisclaimer   string                                    `json:"projection_disclaimer"`
}

func verifierEcosystemValBProjectionDisclaimer() string {
	return "projection_only not_canonical_truth compatibility_diagnostics_conformance advisory_projection"
}

func verifierEcosystemValBSupportedVerifierVersions() []string {
	return []string{
		"reference-verifier/vala-2026.04",
		"reference-verifier/vala-2025.12",
	}
}

func verifierEcosystemValBSupportedTrustRootVersions() []string {
	return []string{
		"2026.04",
		"2025.12",
	}
}

func verifierEcosystemValBOutputClasses() []string {
	return []string{
		VerifierEcosystemValBOutputClassVerified,
		VerifierEcosystemValBOutputClassVerifiedWithWarnings,
		VerifierEcosystemValBOutputClassNonVerifiedInvalid,
		VerifierEcosystemValBOutputClassNonVerifiedIncomplete,
		VerifierEcosystemValBOutputClassNonVerifiedUnsupported,
		VerifierEcosystemValBOutputClassNonVerifiedStale,
		VerifierEcosystemValBOutputClassNonVerifiedRevoked,
		VerifierEcosystemValBOutputClassNonVerifiedSuperseded,
		VerifierEcosystemValBOutputClassNonVerifiedScopeMismatch,
		VerifierEcosystemValBOutputClassNonVerifiedTrustInsufficient,
		VerifierEcosystemValBOutputClassRedactionBlocked,
		VerifierEcosystemValBOutputClassUnknown,
	}
}

func verifierEcosystemValBFixtureTypes() []string {
	return []string{
		VerifierEcosystemValBFixtureTypeValidProofEnvelope,
		VerifierEcosystemValBFixtureTypeStaleProofEnvelope,
		VerifierEcosystemValBFixtureTypeExpiredProofEnvelope,
		VerifierEcosystemValBFixtureTypeRevokedIssuer,
		VerifierEcosystemValBFixtureTypeSupersededProof,
		VerifierEcosystemValBFixtureTypeUnsupportedSchema,
		VerifierEcosystemValBFixtureTypeUnsupportedProofType,
		VerifierEcosystemValBFixtureTypeMalformedTimestamp,
		VerifierEcosystemValBFixtureTypeMissingSignature,
		VerifierEcosystemValBFixtureTypeMissingDigest,
		VerifierEcosystemValBFixtureTypeDigestMismatch,
		VerifierEcosystemValBFixtureTypeInvalidSignature,
		VerifierEcosystemValBFixtureTypeInsufficientTrustMaterial,
		VerifierEcosystemValBFixtureTypeScopeMismatch,
		VerifierEcosystemValBFixtureTypeRedactionBoundaryViolation,
	}
}

func verifierEcosystemValBDiagnosticPrecedenceOrder() []string {
	return []string{
		VerifierEcosystemDiagnosticUnknown,
		VerifierEcosystemDiagnosticInvalidSignature,
		VerifierEcosystemDiagnosticDigestMismatch,
		VerifierEcosystemDiagnosticRevokedIssuer,
		VerifierEcosystemDiagnosticExpiredArtifact,
		VerifierEcosystemDiagnosticStaleArtifact,
		VerifierEcosystemDiagnosticUnsupportedSchema,
		VerifierEcosystemDiagnosticUnsupportedProofType,
		VerifierEcosystemDiagnosticSchemaMismatch,
		VerifierEcosystemDiagnosticInsufficientTrustMaterial,
		VerifierEcosystemDiagnosticSupersededProof,
		VerifierEcosystemDiagnosticIncompleteArtifact,
		VerifierEcosystemDiagnosticScopeMismatch,
		VerifierEcosystemDiagnosticRedactionViolation,
		VerifierEcosystemDiagnosticCompatibilityWarning,
		VerifierEcosystemDiagnosticVerified,
	}
}

func verifierEcosystemValBRequiredConformanceFields() []string {
	return []string{
		"proof_type",
		"schema_version",
		"expected_overall_result",
		"expected_diagnostic_class",
		"expected_output_class",
	}
}

func verifierEcosystemValBRequiredForbiddenClaims() []string {
	return []string{
		"verifier certification",
		"deployment approved",
		"universal authority",
	}
}

func verifierEcosystemValBCompatibilityEntryKey(entry VerifierEcosystemValBCompatibilityEntry) string {
	return strings.Join([]string{
		strings.TrimSpace(entry.SchemaVersion),
		strings.TrimSpace(entry.ProofType),
		strings.TrimSpace(entry.VerifierVersion),
		strings.TrimSpace(entry.TrustRootVersion),
	}, "|")
}

func verifierEcosystemValBRequiredCompatibilityEntryKeys() []string {
	return []string{
		"changelock.verifier.proof_envelope.v1|signed_attestation_envelope|reference-verifier/vala-2026.04|2026.04",
		"changelock.verifier.proof_envelope.v1.1|sealed_artifact_envelope|reference-verifier/vala-2026.04|2026.04",
		"changelock.verifier.proof_envelope.v1|lineage_bundle_envelope|reference-verifier/vala-2026.04|2026.04",
		"changelock.verifier.proof_envelope.v0|signed_attestation_envelope|reference-verifier/vala-2026.04|2026.04",
		"changelock.verifier.proof_envelope.v1.1|signed_attestation_envelope|reference-verifier/vala-2026.04|2025.12",
		"changelock.verifier.proof_envelope.v1|lineage_bundle_envelope|reference-verifier/vala-2025.12|2025.12",
	}
}

func verifierEcosystemValBRequiredMixedVersionCaseIDs() []string {
	return []string{
		"mixed-version:deprecated-schema-v0",
		"mixed-version:trust-root-warning",
		"mixed-version:superseded-verifier-lineage",
	}
}

func verifierEcosystemValBRequiredFixtureIDs() []string {
	return []string{
		"fixture:valid-proof-envelope",
		"fixture:stale-proof-envelope",
		"fixture:expired-proof-envelope",
		"fixture:revoked-issuer",
		"fixture:superseded-proof",
		"fixture:unsupported-schema",
		"fixture:unsupported-proof-type",
		"fixture:malformed-timestamp",
		"fixture:missing-signature",
		"fixture:missing-digest",
		"fixture:digest-mismatch",
		"fixture:invalid-signature",
		"fixture:insufficient-trust-material",
		"fixture:scope-mismatch",
		"fixture:redaction-boundary-violation",
	}
}

func verifierEcosystemValBRequiredConformanceCaseIDs() []string {
	return []string{
		"conformance:valid-proof-envelope",
		"conformance:stale-proof-envelope",
		"conformance:expired-proof-envelope",
		"conformance:revoked-issuer",
		"conformance:superseded-proof",
		"conformance:unsupported-schema",
		"conformance:unsupported-proof-type",
		"conformance:malformed-timestamp",
		"conformance:missing-signature",
		"conformance:missing-digest",
		"conformance:digest-mismatch",
		"conformance:invalid-signature",
		"conformance:insufficient-trust-material",
		"conformance:scope-mismatch",
		"conformance:redaction-boundary-violation",
	}
}

func verifierEcosystemValBRequiredOutputClassMappingKeys() []string {
	return []string{
		"verified|verified",
		"verified_with_warnings|compatibility_warning",
		"invalid|invalid_signature",
		"invalid|digest_mismatch",
		"invalid|scope_mismatch",
		"invalid|redaction_boundary_violation",
		"incomplete|incomplete_artifact",
		"incomplete|insufficient_trust_material",
		"unsupported|unsupported_schema",
		"unsupported|unsupported_proof_type",
		"stale|stale_artifact",
		"stale|expired_artifact",
		"revoked|revoked_issuer",
		"superseded|superseded_proof",
	}
}

func verifierEcosystemValBCompatibilityMatrixEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:compatibility-matrix-001", EvidenceType: "compatibility_matrix", Source: "verifier/compatibility-matrix", Timestamp: "2026-04-27T10:20:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "compatibility_matrix", Caveats: []string{"bounded to declared verifier compatibility matrix entries and mixed-version rules"}},
		{EvidenceID: "evidence:schema-proof-compatibility-001", EvidenceType: "compatibility_evaluator", Source: "verifier/schema-proof-compatibility", Timestamp: "2026-04-27T10:21:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "schema_proof_compatibility", Caveats: []string{"bounded to schema and proof-type compatibility evaluation only"}},
		{EvidenceID: "evidence:mixed-version-diagnostics-001", EvidenceType: "mixed_version_diagnostics", Source: "verifier/mixed-version", Timestamp: "2026-04-27T10:22:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "mixed_version_diagnostics", Caveats: []string{"bounded to explicit mixed-version diagnostic cases"}},
		{EvidenceID: "evidence:diagnostic-precedence-001", EvidenceType: "diagnostic_precedence", Source: "verifier/diagnostic-precedence", Timestamp: "2026-04-27T10:23:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "diagnostic_precedence", Caveats: []string{"bounded to deterministic diagnostic precedence rules"}},
		{EvidenceID: "evidence:fixture-descriptors-001", EvidenceType: "fixture_descriptor_catalog", Source: "verifier/fixtures", Timestamp: "2026-04-27T10:24:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "fixture_descriptors", Caveats: []string{"fixtures are test descriptors only and not fake production evidence"}},
		{EvidenceID: "evidence:conformance-cases-001", EvidenceType: "conformance_cases", Source: "verifier/conformance-cases", Timestamp: "2026-04-27T10:25:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "conformance_cases", Caveats: []string{"bounded to declared conformance cases and forbidden-claim discipline"}},
		{EvidenceID: "evidence:conformance-suite-001", EvidenceType: "conformance_suite", Source: "verifier/conformance-suite", Timestamp: "2026-04-27T10:26:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "conformance_suite", Caveats: []string{"conformance suite remains advisory and does not certify a verifier"}},
		{EvidenceID: "evidence:output-classes-001", EvidenceType: "output_class_catalog", Source: "verifier/output-classes", Timestamp: "2026-04-27T10:27:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "output_classes", Caveats: []string{"output classes remain bounded non-verified or warning-bearing report classes only"}},
		{EvidenceID: "evidence:point7-governance-002", EvidenceType: "state_governance", Source: "verifier/point7-governance", Timestamp: "2026-04-27T10:28:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point7_governance", Caveats: []string{"Val B keeps point_7_state not complete and cannot return point_7_pass"}},
	}
}

func VerifierEcosystemValBVerifierEvidence() []ReferenceArchitectureEvidenceReference {
	return verifierEcosystemValBCompatibilityMatrixEvidence()
}

func verifierEcosystemValBRequiredEvidenceIDs() []string {
	return []string{
		"evidence:compatibility-matrix-001",
		"evidence:schema-proof-compatibility-001",
		"evidence:mixed-version-diagnostics-001",
		"evidence:diagnostic-precedence-001",
		"evidence:fixture-descriptors-001",
		"evidence:conformance-cases-001",
		"evidence:conformance-suite-001",
		"evidence:output-classes-001",
		"evidence:point7-governance-002",
	}
}

func verifierEcosystemValBRequiredEvidenceScopes() []string {
	return []string{
		"compatibility_matrix",
		"schema_proof_compatibility",
		"mixed_version_diagnostics",
		"diagnostic_precedence",
		"fixture_descriptors",
		"conformance_cases",
		"conformance_suite",
		"output_classes",
		"point7_governance",
	}
}

func VerifierEcosystemValBProofEvidenceRefs() []string {
	return []string{
		"point6_integrated_closure",
		"point7_verifier_discipline_foundation",
		"point7_reference_verifier_tooling",
		"point7_compatibility_diagnostics_conformance",
		"evidence:compatibility-matrix-001",
		"evidence:schema-proof-compatibility-001",
		"evidence:mixed-version-diagnostics-001",
		"evidence:diagnostic-precedence-001",
		"evidence:fixture-descriptors-001",
		"evidence:conformance-cases-001",
		"evidence:conformance-suite-001",
		"evidence:output-classes-001",
		"evidence:point7-governance-002",
	}
}

func verifierEcosystemValBProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, evidenceRefs []string) bool {
	if !containsExactTrimmedStringSet(evidenceRefs, VerifierEcosystemValBProofEvidenceRefs()...) {
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
	return containsExactTrimmedStringSet(evidenceIDs, verifierEcosystemValBRequiredEvidenceIDs()...) &&
		containsExactTrimmedStringSet(evidenceScopes, verifierEcosystemValBRequiredEvidenceScopes()...)
}

func VerifierEcosystemValBCompatibilityMatrixModel() VerifierEcosystemValBCompatibilityMatrix {
	disclaimer := verifierEcosystemValBProjectionDisclaimer()
	return VerifierEcosystemValBCompatibilityMatrix{
		CurrentState:               "verifier_ecosystem_valb_compatibility_matrix_ready",
		CompatibilityMatrixID:      "verifier-compatibility-matrix-valb",
		Version:                    "2026.04",
		SupportedSchemaVersions:    []string{"changelock.verifier.proof_envelope.v0", "changelock.verifier.proof_envelope.v1", "changelock.verifier.proof_envelope.v1.1"},
		SupportedProofTypes:        verifierEcosystemVal0SupportedProofTypes(),
		SupportedVerifierVersions:  verifierEcosystemValBSupportedVerifierVersions(),
		SupportedTrustRootVersions: verifierEcosystemValBSupportedTrustRootVersions(),
		CompatibilityEntries: []VerifierEcosystemValBCompatibilityEntry{
			{CurrentState: "compatibility_entry_ready", EntryID: "matrix-entry:signed-attestation-v1", SchemaVersion: "changelock.verifier.proof_envelope.v1", ProofType: VerifierEcosystemProofTypeSignedAttestation, VerifierVersion: "reference-verifier/vala-2026.04", TrustRootVersion: "2026.04", CompatibilityState: ReferenceArchitectureCompatibilityCompatible, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticVerified}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "compatibility_entry_ready", EntryID: "matrix-entry:sealed-artifact-v1.1", SchemaVersion: "changelock.verifier.proof_envelope.v1.1", ProofType: VerifierEcosystemProofTypeSealedArtifact, VerifierVersion: "reference-verifier/vala-2026.04", TrustRootVersion: "2026.04", CompatibilityState: ReferenceArchitectureCompatibilityCompatible, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticVerified}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "compatibility_entry_ready", EntryID: "matrix-entry:lineage-bundle-v1", SchemaVersion: "changelock.verifier.proof_envelope.v1", ProofType: VerifierEcosystemProofTypeLineageBundle, VerifierVersion: "reference-verifier/vala-2026.04", TrustRootVersion: "2026.04", CompatibilityState: ReferenceArchitectureCompatibilityCompatible, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticVerified}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "compatibility_entry_ready", EntryID: "matrix-entry:deprecated-schema-v0", SchemaVersion: "changelock.verifier.proof_envelope.v0", ProofType: VerifierEcosystemProofTypeSignedAttestation, VerifierVersion: "reference-verifier/vala-2026.04", TrustRootVersion: "2026.04", CompatibilityState: ReferenceArchitectureCompatibilityDeprecated, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticCompatibilityWarning}, Caveats: []string{"deprecated schema remains recognized only with explicit migration caveat"}, MigrationOrSupersessionRef: "migration:proof-envelope-v1", ProjectionDisclaimer: disclaimer},
			{CurrentState: "compatibility_entry_ready", EntryID: "matrix-entry:trust-root-warning-v1.1", SchemaVersion: "changelock.verifier.proof_envelope.v1.1", ProofType: VerifierEcosystemProofTypeSignedAttestation, VerifierVersion: "reference-verifier/vala-2026.04", TrustRootVersion: "2025.12", CompatibilityState: ReferenceArchitectureCompatibilityCompatibleWithWarning, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticCompatibilityWarning}, Caveats: []string{"older trust-root version remains bounded and warning-bearing until rollover completes"}, MigrationOrSupersessionRef: "rollover:trust-root-2026.04", ProjectionDisclaimer: disclaimer},
			{CurrentState: "compatibility_entry_ready", EntryID: "matrix-entry:superseded-lineage-v1", SchemaVersion: "changelock.verifier.proof_envelope.v1", ProofType: VerifierEcosystemProofTypeLineageBundle, VerifierVersion: "reference-verifier/vala-2025.12", TrustRootVersion: "2025.12", CompatibilityState: ReferenceArchitectureCompatibilitySuperseded, RequiredDiagnostics: []string{VerifierEcosystemDiagnosticSupersededProof}, Caveats: []string{"superseded lineage verifier path remains representable but not cleanly verified"}, MigrationOrSupersessionRef: "supersession:lineage-bundle-current", ProjectionDisclaimer: disclaimer},
		},
		DeprecatedVersions:  []string{"changelock.verifier.proof_envelope.v0"},
		SupersededVersions:  []string{"reference-verifier/vala-2025.12"},
		UnsupportedVersions: []string{"changelock.verifier.proof_envelope.v9", "proof_type:unknown"},
		MixedVersionRules: []string{
			"deprecated schema versions remain warning-bearing and require explicit migration guidance",
			"superseded verifier versions remain non-clean and require supersession guidance",
			"older trust-root versions remain warning-bearing until explicit rollover metadata is present",
		},
		Caveats:              []string{"compatibility matrix is advisory, version-bound, and does not certify a verifier"},
		ProjectionDisclaimer: disclaimer,
		CreatedAt:            "2026-04-27T10:20:00Z",
		UpdatedAt:            "2026-04-27T10:20:00Z",
	}
}

func VerifierEcosystemValBSchemaProofCompatibilityModel() VerifierEcosystemValBSchemaProofCompatibility {
	return VerifierEcosystemValBSchemaProofCompatibility{
		CurrentState:           "verifier_ecosystem_valb_schema_proof_compatibility_ready",
		CompatibilityID:        "schema-proof-compatibility-current",
		CompatibilityMatrixRef: "verifier-compatibility-matrix-valb",
		SchemaVersion:          "changelock.verifier.proof_envelope.v1",
		ProofType:              VerifierEcosystemProofTypeSignedAttestation,
		VerifierVersion:        "reference-verifier/vala-2026.04",
		TrustRootVersion:       "2026.04",
		CompatibilityState:     ReferenceArchitectureCompatibilityCompatible,
		DerivedDiagnosticClass: VerifierEcosystemDiagnosticVerified,
		RequiredDiagnostics:    []string{VerifierEcosystemDiagnosticVerified},
		ProjectionDisclaimer:   verifierEcosystemValBProjectionDisclaimer(),
	}
}

func VerifierEcosystemValBMixedVersionDiagnosticsCatalogModel() VerifierEcosystemValBMixedVersionDiagnosticsCatalog {
	disclaimer := verifierEcosystemValBProjectionDisclaimer()
	return VerifierEcosystemValBMixedVersionDiagnosticsCatalog{
		CurrentState:                 "verifier_ecosystem_valb_mixed_version_catalog_ready",
		CatalogID:                    "verifier-mixed-version-catalog-valb",
		SupportedCompatibilityStates: verifierEcosystemVal0CompatibilityResults(),
		SupportedDiagnosticClasses:   verifierEcosystemVal0DiagnosticClasses(),
		Cases: []VerifierEcosystemValBMixedVersionCase{
			{CurrentState: "mixed_version_case_ready", MixedVersionCaseID: "mixed-version:deprecated-schema-v0", SchemaVersion: "changelock.verifier.proof_envelope.v0", ProofType: VerifierEcosystemProofTypeSignedAttestation, VerifierVersion: "reference-verifier/vala-2026.04", TrustRootVersion: "2026.04", ExpectedCompatibilityState: ReferenceArchitectureCompatibilityDeprecated, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticCompatibilityWarning, MigrationOrSupersessionRef: "migration:proof-envelope-v1", Caveats: []string{"deprecated schema remains warning-bearing and does not cleanly verify"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "mixed_version_case_ready", MixedVersionCaseID: "mixed-version:trust-root-warning", SchemaVersion: "changelock.verifier.proof_envelope.v1.1", ProofType: VerifierEcosystemProofTypeSignedAttestation, VerifierVersion: "reference-verifier/vala-2026.04", TrustRootVersion: "2025.12", ExpectedCompatibilityState: ReferenceArchitectureCompatibilityCompatibleWithWarning, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticCompatibilityWarning, MigrationOrSupersessionRef: "rollover:trust-root-2026.04", Caveats: []string{"mixed-version trust-root rollover remains warning-bearing and explicitly bounded"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "mixed_version_case_ready", MixedVersionCaseID: "mixed-version:superseded-verifier-lineage", SchemaVersion: "changelock.verifier.proof_envelope.v1", ProofType: VerifierEcosystemProofTypeLineageBundle, VerifierVersion: "reference-verifier/vala-2025.12", TrustRootVersion: "2025.12", ExpectedCompatibilityState: ReferenceArchitectureCompatibilitySuperseded, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticSupersededProof, MigrationOrSupersessionRef: "supersession:lineage-bundle-current", Caveats: []string{"superseded verifier remains representable but cannot become clean verified"}, ProjectionDisclaimer: disclaimer},
		},
		ProjectionDisclaimer: disclaimer,
	}
}

func VerifierEcosystemValBDiagnosticPrecedenceModel() VerifierEcosystemValBDiagnosticPrecedence {
	return VerifierEcosystemValBDiagnosticPrecedence{
		CurrentState:               "verifier_ecosystem_valb_diagnostic_precedence_ready",
		PrecedenceID:               "verifier-diagnostic-precedence-valb",
		SupportedDiagnosticClasses: verifierEcosystemVal0DiagnosticClasses(),
		PrecedenceOrder:            verifierEcosystemValBDiagnosticPrecedenceOrder(),
		ObservedDiagnostics:        []string{VerifierEcosystemDiagnosticVerified},
		DerivedDiagnosticClass:     VerifierEcosystemDiagnosticVerified,
		ProjectionDisclaimer:       verifierEcosystemValBProjectionDisclaimer(),
	}
}

func VerifierEcosystemValBFixtureCatalogModel() VerifierEcosystemValBFixtureCatalog {
	disclaimer := verifierEcosystemValBProjectionDisclaimer()
	return VerifierEcosystemValBFixtureCatalog{
		CurrentState:          "verifier_ecosystem_valb_fixture_catalog_ready",
		FixtureCatalogID:      "verifier-fixture-catalog-valb",
		SupportedFixtureTypes: verifierEcosystemValBFixtureTypes(),
		Fixtures: []VerifierEcosystemValBFixtureDescriptor{
			{CurrentState: "fixture_ready", FixtureID: "fixture:valid-proof-envelope", FixtureType: VerifierEcosystemValBFixtureTypeValidProofEnvelope, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultVerified, ExpectedDiagnostic: VerifierEcosystemDiagnosticVerified, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:verifier-contract-001", "evidence:proof-envelope-001"}, Caveats: []string{"fixture remains bounded to declared verifier scope only"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:stale-proof-envelope", FixtureType: VerifierEcosystemValBFixtureTypeStaleProofEnvelope, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultStale, ExpectedDiagnostic: VerifierEcosystemDiagnosticStaleArtifact, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:proof-envelope-001"}, Caveats: []string{"stale artifact remains explicitly non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:expired-proof-envelope", FixtureType: VerifierEcosystemValBFixtureTypeExpiredProofEnvelope, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultStale, ExpectedDiagnostic: VerifierEcosystemDiagnosticExpiredArtifact, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:proof-envelope-001"}, Caveats: []string{"expired artifact remains explicitly non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:revoked-issuer", FixtureType: VerifierEcosystemValBFixtureTypeRevokedIssuer, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultRevoked, ExpectedDiagnostic: VerifierEcosystemDiagnosticRevokedIssuer, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:trust-root-001", "evidence:revocation-001"}, Caveats: []string{"revoked issuer remains explicitly non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:superseded-proof", FixtureType: VerifierEcosystemValBFixtureTypeSupersededProof, ProofType: VerifierEcosystemProofTypeLineageBundle, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultSuperseded, ExpectedDiagnostic: VerifierEcosystemDiagnosticSupersededProof, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:compatibility-001"}, Caveats: []string{"superseded proof remains non-clean and warning-bearing"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:unsupported-schema", FixtureType: VerifierEcosystemValBFixtureTypeUnsupportedSchema, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v9", ExpectedResult: VerifierEcosystemValAOverallResultUnsupported, ExpectedDiagnostic: VerifierEcosystemDiagnosticUnsupportedSchema, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:compatibility-001"}, Caveats: []string{"unsupported schema remains explicitly non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:unsupported-proof-type", FixtureType: VerifierEcosystemValBFixtureTypeUnsupportedProofType, ProofType: "unsupported_proof_type_fixture", SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultUnsupported, ExpectedDiagnostic: VerifierEcosystemDiagnosticUnsupportedProofType, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:compatibility-001"}, Caveats: []string{"unsupported proof type remains explicitly non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:malformed-timestamp", FixtureType: VerifierEcosystemValBFixtureTypeMalformedTimestamp, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultIncomplete, ExpectedDiagnostic: VerifierEcosystemDiagnosticIncompleteArtifact, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:proof-envelope-001"}, Caveats: []string{"malformed timestamps remain non-verified conformance fixtures"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:missing-signature", FixtureType: VerifierEcosystemValBFixtureTypeMissingSignature, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultIncomplete, ExpectedDiagnostic: VerifierEcosystemDiagnosticIncompleteArtifact, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:proof-envelope-001"}, Caveats: []string{"missing signature fixtures remain incomplete and non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:missing-digest", FixtureType: VerifierEcosystemValBFixtureTypeMissingDigest, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultIncomplete, ExpectedDiagnostic: VerifierEcosystemDiagnosticIncompleteArtifact, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:proof-envelope-001"}, Caveats: []string{"missing digest fixtures remain incomplete and non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:digest-mismatch", FixtureType: VerifierEcosystemValBFixtureTypeDigestMismatch, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultInvalid, ExpectedDiagnostic: VerifierEcosystemDiagnosticDigestMismatch, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:proof-envelope-001"}, Caveats: []string{"digest mismatch fixtures remain non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:invalid-signature", FixtureType: VerifierEcosystemValBFixtureTypeInvalidSignature, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultInvalid, ExpectedDiagnostic: VerifierEcosystemDiagnosticInvalidSignature, ExpectedOutputBoundary: VerifierEcosystemScopeAuditorSafe, RequiredEvidenceRefs: []string{"evidence:proof-envelope-001"}, Caveats: []string{"invalid signature fixtures remain non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:insufficient-trust-material", FixtureType: VerifierEcosystemValBFixtureTypeInsufficientTrustMaterial, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultIncomplete, ExpectedDiagnostic: VerifierEcosystemDiagnosticInsufficientTrustMaterial, ExpectedOutputBoundary: VerifierEcosystemScopeRestrictedOffline, RequiredEvidenceRefs: []string{"evidence:trust-root-001"}, Caveats: []string{"insufficient trust material remains incomplete and non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:scope-mismatch", FixtureType: VerifierEcosystemValBFixtureTypeScopeMismatch, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultInvalid, ExpectedDiagnostic: VerifierEcosystemDiagnosticScopeMismatch, ExpectedOutputBoundary: VerifierEcosystemScopePartnerSafe, RequiredEvidenceRefs: []string{"evidence:verification-scope-001"}, Caveats: []string{"scope mismatch fixtures remain non-verified"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "fixture_ready", FixtureID: "fixture:redaction-boundary-violation", FixtureType: VerifierEcosystemValBFixtureTypeRedactionBoundaryViolation, ProofType: VerifierEcosystemProofTypeSignedAttestation, SchemaVersion: "changelock.verifier.proof_envelope.v1", ExpectedResult: VerifierEcosystemValAOverallResultInvalid, ExpectedDiagnostic: VerifierEcosystemDiagnosticRedactionViolation, ExpectedOutputBoundary: VerifierEcosystemScopePublicSafe, RequiredEvidenceRefs: []string{"evidence:output-boundary-001"}, Caveats: []string{"redaction boundary violations remain blocked from public-safe output"}, ProjectionDisclaimer: disclaimer},
		},
		ProjectionDisclaimer: disclaimer,
	}
}

func VerifierEcosystemValBConformanceCaseCatalogModel() VerifierEcosystemValBConformanceCaseCatalog {
	disclaimer := verifierEcosystemValBProjectionDisclaimer()
	return VerifierEcosystemValBConformanceCaseCatalog{
		CurrentState:         "verifier_ecosystem_valb_conformance_catalog_ready",
		ConformanceCatalogID: "verifier-conformance-catalog-valb",
		Cases: []VerifierEcosystemValBConformanceCase{
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:valid-proof-envelope", FixtureRef: "fixture:valid-proof-envelope", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/valid-proof-envelope", ExpectedOverallResult: VerifierEcosystemValAOverallResultVerified, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticVerified, ExpectedOutputClass: VerifierEcosystemValBOutputClassVerified, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:stale-proof-envelope", FixtureRef: "fixture:stale-proof-envelope", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/stale-proof-envelope", ExpectedOverallResult: VerifierEcosystemValAOverallResultStale, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticStaleArtifact, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedStale, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:expired-proof-envelope", FixtureRef: "fixture:expired-proof-envelope", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/expired-proof-envelope", ExpectedOverallResult: VerifierEcosystemValAOverallResultStale, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticExpiredArtifact, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedStale, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:revoked-issuer", FixtureRef: "fixture:revoked-issuer", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/revoked-issuer", ExpectedOverallResult: VerifierEcosystemValAOverallResultRevoked, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticRevokedIssuer, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedRevoked, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:superseded-proof", FixtureRef: "fixture:superseded-proof", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/superseded-proof", ExpectedOverallResult: VerifierEcosystemValAOverallResultSuperseded, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticSupersededProof, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedSuperseded, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:unsupported-schema", FixtureRef: "fixture:unsupported-schema", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/unsupported-schema", ExpectedOverallResult: VerifierEcosystemValAOverallResultUnsupported, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticUnsupportedSchema, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedUnsupported, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:unsupported-proof-type", FixtureRef: "fixture:unsupported-proof-type", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/unsupported-proof-type", ExpectedOverallResult: VerifierEcosystemValAOverallResultUnsupported, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticUnsupportedProofType, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedUnsupported, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:malformed-timestamp", FixtureRef: "fixture:malformed-timestamp", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/malformed-timestamp", ExpectedOverallResult: VerifierEcosystemValAOverallResultIncomplete, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticIncompleteArtifact, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedIncomplete, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:missing-signature", FixtureRef: "fixture:missing-signature", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/missing-signature", ExpectedOverallResult: VerifierEcosystemValAOverallResultIncomplete, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticIncompleteArtifact, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedIncomplete, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:missing-digest", FixtureRef: "fixture:missing-digest", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/missing-digest", ExpectedOverallResult: VerifierEcosystemValAOverallResultIncomplete, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticIncompleteArtifact, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedIncomplete, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:digest-mismatch", FixtureRef: "fixture:digest-mismatch", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/digest-mismatch", ExpectedOverallResult: VerifierEcosystemValAOverallResultInvalid, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticDigestMismatch, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedInvalid, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:invalid-signature", FixtureRef: "fixture:invalid-signature", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/invalid-signature", ExpectedOverallResult: VerifierEcosystemValAOverallResultInvalid, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticInvalidSignature, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedInvalid, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:insufficient-trust-material", FixtureRef: "fixture:insufficient-trust-material", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/insufficient-trust-material", ExpectedOverallResult: VerifierEcosystemValAOverallResultIncomplete, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticInsufficientTrustMaterial, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedTrustInsufficient, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:scope-mismatch", FixtureRef: "fixture:scope-mismatch", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/scope-mismatch", ExpectedOverallResult: VerifierEcosystemValAOverallResultInvalid, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticScopeMismatch, ExpectedOutputClass: VerifierEcosystemValBOutputClassNonVerifiedScopeMismatch, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
			{CurrentState: "conformance_case_ready", ConformanceCaseID: "conformance:redaction-boundary-violation", FixtureRef: "fixture:redaction-boundary-violation", VerifierContractRef: "verifier-contract-ref/val0", InputRef: "input-ref/redaction-boundary-violation", ExpectedOverallResult: VerifierEcosystemValAOverallResultInvalid, ExpectedDiagnosticClass: VerifierEcosystemDiagnosticRedactionViolation, ExpectedOutputClass: VerifierEcosystemValBOutputClassRedactionBlocked, RequiredFields: verifierEcosystemValBRequiredConformanceFields(), ForbiddenClaims: verifierEcosystemValBRequiredForbiddenClaims(), ProjectionDisclaimer: disclaimer},
		},
		ProjectionDisclaimer: disclaimer,
	}
}

func VerifierEcosystemValBConformanceSuiteModel() VerifierEcosystemValBConformanceSuite {
	return VerifierEcosystemValBConformanceSuite{
		CurrentState:               "verifier_ecosystem_valb_conformance_suite_ready",
		ConformanceSuiteID:         "verifier-conformance-suite-valb",
		RequiredCaseRefs:           verifierEcosystemValBRequiredConformanceCaseIDs(),
		RequiredFixtureRefs:        verifierEcosystemValBRequiredFixtureIDs(),
		SupportedOutputClasses:     verifierEcosystemValBOutputClasses(),
		SupportedDiagnosticClasses: verifierEcosystemVal0DiagnosticClasses(),
		Caveats:                    []string{"conformance suite proves only declared case behavior and does not certify a verifier"},
		ProjectionDisclaimer:       verifierEcosystemValBProjectionDisclaimer(),
	}
}

func VerifierEcosystemValBOutputClassCatalogModel() VerifierEcosystemValBOutputClassCatalog {
	disclaimer := verifierEcosystemValBProjectionDisclaimer()
	return VerifierEcosystemValBOutputClassCatalog{
		CurrentState:           "verifier_ecosystem_valb_output_class_catalog_ready",
		OutputClassCatalogID:   "verifier-output-classes-valb",
		SupportedOutputClasses: verifierEcosystemValBOutputClasses(),
		Mappings: []VerifierEcosystemValBOutputClassMapping{
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:verified", OverallResult: VerifierEcosystemValAOverallResultVerified, DiagnosticClass: VerifierEcosystemDiagnosticVerified, OutputClass: VerifierEcosystemValBOutputClassVerified, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:verified-with-warnings", OverallResult: VerifierEcosystemValAOverallResultWarnings, DiagnosticClass: VerifierEcosystemDiagnosticCompatibilityWarning, OutputClass: VerifierEcosystemValBOutputClassVerifiedWithWarnings, Caveats: []string{"warning-bearing verification remains advisory and explicit"}, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:invalid-signature", OverallResult: VerifierEcosystemValAOverallResultInvalid, DiagnosticClass: VerifierEcosystemDiagnosticInvalidSignature, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedInvalid, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:digest-mismatch", OverallResult: VerifierEcosystemValAOverallResultInvalid, DiagnosticClass: VerifierEcosystemDiagnosticDigestMismatch, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedInvalid, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:scope-mismatch", OverallResult: VerifierEcosystemValAOverallResultInvalid, DiagnosticClass: VerifierEcosystemDiagnosticScopeMismatch, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedScopeMismatch, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:redaction-blocked", OverallResult: VerifierEcosystemValAOverallResultInvalid, DiagnosticClass: VerifierEcosystemDiagnosticRedactionViolation, OutputClass: VerifierEcosystemValBOutputClassRedactionBlocked, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:incomplete-artifact", OverallResult: VerifierEcosystemValAOverallResultIncomplete, DiagnosticClass: VerifierEcosystemDiagnosticIncompleteArtifact, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedIncomplete, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:insufficient-trust", OverallResult: VerifierEcosystemValAOverallResultIncomplete, DiagnosticClass: VerifierEcosystemDiagnosticInsufficientTrustMaterial, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedTrustInsufficient, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:unsupported-schema", OverallResult: VerifierEcosystemValAOverallResultUnsupported, DiagnosticClass: VerifierEcosystemDiagnosticUnsupportedSchema, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedUnsupported, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:unsupported-proof-type", OverallResult: VerifierEcosystemValAOverallResultUnsupported, DiagnosticClass: VerifierEcosystemDiagnosticUnsupportedProofType, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedUnsupported, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:stale-artifact", OverallResult: VerifierEcosystemValAOverallResultStale, DiagnosticClass: VerifierEcosystemDiagnosticStaleArtifact, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedStale, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:expired-artifact", OverallResult: VerifierEcosystemValAOverallResultStale, DiagnosticClass: VerifierEcosystemDiagnosticExpiredArtifact, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedStale, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:revoked-issuer", OverallResult: VerifierEcosystemValAOverallResultRevoked, DiagnosticClass: VerifierEcosystemDiagnosticRevokedIssuer, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedRevoked, ProjectionDisclaimer: disclaimer},
			{CurrentState: "output_class_mapping_ready", OutputClassID: "output-class:superseded-proof", OverallResult: VerifierEcosystemValAOverallResultSuperseded, DiagnosticClass: VerifierEcosystemDiagnosticSupersededProof, OutputClass: VerifierEcosystemValBOutputClassNonVerifiedSuperseded, ProjectionDisclaimer: disclaimer},
		},
		ProjectionDisclaimer: disclaimer,
	}
}

func verifierEcosystemValBExpectedDiagnosticForCompatibilityState(state string) string {
	switch strings.TrimSpace(state) {
	case ReferenceArchitectureCompatibilityCompatible:
		return VerifierEcosystemDiagnosticVerified
	case ReferenceArchitectureCompatibilityCompatibleWithWarning, ReferenceArchitectureCompatibilityDeprecated:
		return VerifierEcosystemDiagnosticCompatibilityWarning
	case ReferenceArchitectureCompatibilitySuperseded:
		return VerifierEcosystemDiagnosticSupersededProof
	case ReferenceArchitectureCompatibilityUnsupported:
		return VerifierEcosystemDiagnosticUnsupportedSchema
	default:
		return VerifierEcosystemDiagnosticUnknown
	}
}

func deriveVerifierEcosystemValBOutputClass(overallResult, diagnostic string, caveats []string) string {
	switch strings.TrimSpace(diagnostic) {
	case VerifierEcosystemDiagnosticVerified:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultVerified {
			return VerifierEcosystemValBOutputClassVerified
		}
	case VerifierEcosystemDiagnosticCompatibilityWarning:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultWarnings && len(caveats) > 0 {
			return VerifierEcosystemValBOutputClassVerifiedWithWarnings
		}
	case VerifierEcosystemDiagnosticInvalidSignature, VerifierEcosystemDiagnosticDigestMismatch:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultInvalid {
			return VerifierEcosystemValBOutputClassNonVerifiedInvalid
		}
	case VerifierEcosystemDiagnosticIncompleteArtifact:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultIncomplete {
			return VerifierEcosystemValBOutputClassNonVerifiedIncomplete
		}
	case VerifierEcosystemDiagnosticUnsupportedSchema, VerifierEcosystemDiagnosticUnsupportedProofType:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultUnsupported {
			return VerifierEcosystemValBOutputClassNonVerifiedUnsupported
		}
	case VerifierEcosystemDiagnosticStaleArtifact, VerifierEcosystemDiagnosticExpiredArtifact:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultStale {
			return VerifierEcosystemValBOutputClassNonVerifiedStale
		}
	case VerifierEcosystemDiagnosticRevokedIssuer:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultRevoked {
			return VerifierEcosystemValBOutputClassNonVerifiedRevoked
		}
	case VerifierEcosystemDiagnosticSupersededProof:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultSuperseded {
			return VerifierEcosystemValBOutputClassNonVerifiedSuperseded
		}
	case VerifierEcosystemDiagnosticScopeMismatch:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultInvalid {
			return VerifierEcosystemValBOutputClassNonVerifiedScopeMismatch
		}
	case VerifierEcosystemDiagnosticInsufficientTrustMaterial:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultIncomplete {
			return VerifierEcosystemValBOutputClassNonVerifiedTrustInsufficient
		}
	case VerifierEcosystemDiagnosticRedactionViolation:
		if strings.TrimSpace(overallResult) == VerifierEcosystemValAOverallResultInvalid {
			return VerifierEcosystemValBOutputClassRedactionBlocked
		}
	}
	return VerifierEcosystemValBOutputClassUnknown
}

func DeriveVerifierEcosystemValBDiagnostic(observedDiagnostics, caveats []string) string {
	if len(observedDiagnostics) == 0 {
		return VerifierEcosystemDiagnosticUnknown
	}
	for _, diagnostic := range observedDiagnostics {
		if !containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), diagnostic) {
			return VerifierEcosystemDiagnosticUnknown
		}
	}
	for _, item := range verifierEcosystemValBDiagnosticPrecedenceOrder() {
		if containsTrimmedString(observedDiagnostics, item) {
			if item == VerifierEcosystemDiagnosticCompatibilityWarning && len(caveats) == 0 {
				return VerifierEcosystemDiagnosticUnknown
			}
			return item
		}
	}
	return VerifierEcosystemDiagnosticUnknown
}

func verifierEcosystemValBFixtureExpectedOutcome(fixtureType string) (string, string) {
	switch strings.TrimSpace(fixtureType) {
	case VerifierEcosystemValBFixtureTypeValidProofEnvelope:
		return VerifierEcosystemValAOverallResultVerified, VerifierEcosystemDiagnosticVerified
	case VerifierEcosystemValBFixtureTypeStaleProofEnvelope:
		return VerifierEcosystemValAOverallResultStale, VerifierEcosystemDiagnosticStaleArtifact
	case VerifierEcosystemValBFixtureTypeExpiredProofEnvelope:
		return VerifierEcosystemValAOverallResultStale, VerifierEcosystemDiagnosticExpiredArtifact
	case VerifierEcosystemValBFixtureTypeRevokedIssuer:
		return VerifierEcosystemValAOverallResultRevoked, VerifierEcosystemDiagnosticRevokedIssuer
	case VerifierEcosystemValBFixtureTypeSupersededProof:
		return VerifierEcosystemValAOverallResultSuperseded, VerifierEcosystemDiagnosticSupersededProof
	case VerifierEcosystemValBFixtureTypeUnsupportedSchema:
		return VerifierEcosystemValAOverallResultUnsupported, VerifierEcosystemDiagnosticUnsupportedSchema
	case VerifierEcosystemValBFixtureTypeUnsupportedProofType:
		return VerifierEcosystemValAOverallResultUnsupported, VerifierEcosystemDiagnosticUnsupportedProofType
	case VerifierEcosystemValBFixtureTypeMalformedTimestamp, VerifierEcosystemValBFixtureTypeMissingSignature, VerifierEcosystemValBFixtureTypeMissingDigest:
		return VerifierEcosystemValAOverallResultIncomplete, VerifierEcosystemDiagnosticIncompleteArtifact
	case VerifierEcosystemValBFixtureTypeDigestMismatch:
		return VerifierEcosystemValAOverallResultInvalid, VerifierEcosystemDiagnosticDigestMismatch
	case VerifierEcosystemValBFixtureTypeInvalidSignature:
		return VerifierEcosystemValAOverallResultInvalid, VerifierEcosystemDiagnosticInvalidSignature
	case VerifierEcosystemValBFixtureTypeInsufficientTrustMaterial:
		return VerifierEcosystemValAOverallResultIncomplete, VerifierEcosystemDiagnosticInsufficientTrustMaterial
	case VerifierEcosystemValBFixtureTypeScopeMismatch:
		return VerifierEcosystemValAOverallResultInvalid, VerifierEcosystemDiagnosticScopeMismatch
	case VerifierEcosystemValBFixtureTypeRedactionBoundaryViolation:
		return VerifierEcosystemValAOverallResultInvalid, VerifierEcosystemDiagnosticRedactionViolation
	default:
		return VerifierEcosystemValAOverallResultUnknown, VerifierEcosystemDiagnosticUnknown
	}
}

func verifierEcosystemValBFixtureProofTypeAllowed(model VerifierEcosystemValBFixtureDescriptor) bool {
	if containsTrimmedString(verifierEcosystemVal0SupportedProofTypes(), model.ProofType) {
		return true
	}
	return strings.TrimSpace(model.FixtureType) == VerifierEcosystemValBFixtureTypeUnsupportedProofType
}

func verifierEcosystemValBFixtureSchemaAllowed(model VerifierEcosystemValBFixtureDescriptor) bool {
	if containsTrimmedString(VerifierEcosystemValBCompatibilityMatrixModel().SupportedSchemaVersions, model.SchemaVersion) {
		return true
	}
	return strings.TrimSpace(model.FixtureType) == VerifierEcosystemValBFixtureTypeUnsupportedSchema
}

func EvaluateVerifierEcosystemValBCompatibilityMatrixState(matrix VerifierEcosystemValBCompatibilityMatrix) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		matrix.CompatibilityMatrixID,
		matrix.Version,
		matrix.ProjectionDisclaimer,
		matrix.CreatedAt,
		matrix.UpdatedAt,
	) || len(matrix.CompatibilityEntries) == 0 {
		return VerifierEcosystemValBCompatibilityMatrixStateIncomplete
	}
	if !containsExactTrimmedStringSet(matrix.SupportedProofTypes, verifierEcosystemVal0SupportedProofTypes()...) ||
		!containsExactTrimmedStringSet(matrix.SupportedVerifierVersions, verifierEcosystemValBSupportedVerifierVersions()...) ||
		!containsExactTrimmedStringSet(matrix.SupportedTrustRootVersions, verifierEcosystemValBSupportedTrustRootVersions()...) ||
		!containsExactTrimmedStringSet(matrix.SupportedSchemaVersions, "changelock.verifier.proof_envelope.v0", "changelock.verifier.proof_envelope.v1", "changelock.verifier.proof_envelope.v1.1") ||
		!containsExactTrimmedStringSet(matrix.DeprecatedVersions, "changelock.verifier.proof_envelope.v0") ||
		!containsExactTrimmedStringSet(matrix.SupersededVersions, "reference-verifier/vala-2025.12") ||
		!containsExactTrimmedStringSet(matrix.UnsupportedVersions, "changelock.verifier.proof_envelope.v9", "proof_type:unknown") ||
		!containsExactTrimmedStringSet(matrix.MixedVersionRules,
			"deprecated schema versions remain warning-bearing and require explicit migration guidance",
			"superseded verifier versions remain non-clean and require supersession guidance",
			"older trust-root versions remain warning-bearing until explicit rollover metadata is present",
		) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(matrix.ProjectionDisclaimer) {
		return VerifierEcosystemValBCompatibilityMatrixStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(matrix.CreatedAt); !ok {
		return VerifierEcosystemValBCompatibilityMatrixStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(matrix.UpdatedAt); !ok {
		return VerifierEcosystemValBCompatibilityMatrixStatePartial
	}
	if verifierEcosystemVal0HasOverclaim(strings.Join(matrix.Caveats, " "), matrix.ProjectionDisclaimer) {
		return VerifierEcosystemValBCompatibilityMatrixStateBlocked
	}
	entryKeys := make([]string, 0, len(matrix.CompatibilityEntries))
	for _, entry := range matrix.CompatibilityEntries {
		if !referenceArchitectureValBRequiredRefsPresent(
			entry.EntryID,
			entry.SchemaVersion,
			entry.ProofType,
			entry.VerifierVersion,
			entry.TrustRootVersion,
			entry.CompatibilityState,
			entry.ProjectionDisclaimer,
		) || len(entry.RequiredDiagnostics) == 0 {
			return VerifierEcosystemValBCompatibilityMatrixStateIncomplete
		}
		if !containsTrimmedString(matrix.SupportedSchemaVersions, entry.SchemaVersion) ||
			!containsTrimmedString(matrix.SupportedProofTypes, entry.ProofType) ||
			!containsTrimmedString(matrix.SupportedVerifierVersions, entry.VerifierVersion) ||
			!containsTrimmedString(matrix.SupportedTrustRootVersions, entry.TrustRootVersion) ||
			!containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), entry.CompatibilityState) ||
			!containsAllTrimmedStrings(verifierEcosystemVal0DiagnosticClasses(), entry.RequiredDiagnostics...) ||
			!verifierEcosystemVal0HasProjectionDisclaimer(entry.ProjectionDisclaimer) {
			return VerifierEcosystemValBCompatibilityMatrixStateUnknown
		}
		if verifierEcosystemVal0HasOverclaim(strings.Join(entry.Caveats, " "), entry.ProjectionDisclaimer) {
			return VerifierEcosystemValBCompatibilityMatrixStateBlocked
		}
		switch strings.TrimSpace(entry.CompatibilityState) {
		case ReferenceArchitectureCompatibilityCompatible:
			if !containsExactTrimmedStringSet(entry.RequiredDiagnostics, VerifierEcosystemDiagnosticVerified) {
				return VerifierEcosystemValBCompatibilityMatrixStatePartial
			}
		case ReferenceArchitectureCompatibilityCompatibleWithWarning:
			if len(entry.Caveats) == 0 || !containsExactTrimmedStringSet(entry.RequiredDiagnostics, VerifierEcosystemDiagnosticCompatibilityWarning) {
				return VerifierEcosystemValBCompatibilityMatrixStateBlocked
			}
		case ReferenceArchitectureCompatibilityDeprecated:
			if len(entry.Caveats) == 0 || strings.TrimSpace(entry.MigrationOrSupersessionRef) == "" || !containsExactTrimmedStringSet(entry.RequiredDiagnostics, VerifierEcosystemDiagnosticCompatibilityWarning) {
				return VerifierEcosystemValBCompatibilityMatrixStateBlocked
			}
		case ReferenceArchitectureCompatibilitySuperseded:
			if len(entry.Caveats) == 0 || strings.TrimSpace(entry.MigrationOrSupersessionRef) == "" || !containsExactTrimmedStringSet(entry.RequiredDiagnostics, VerifierEcosystemDiagnosticSupersededProof) {
				return VerifierEcosystemValBCompatibilityMatrixStateBlocked
			}
		case ReferenceArchitectureCompatibilityUnsupported:
			return VerifierEcosystemValBCompatibilityMatrixStateBlocked
		case ReferenceArchitectureCompatibilityUnknown:
			return VerifierEcosystemValBCompatibilityMatrixStateUnknown
		default:
			return VerifierEcosystemValBCompatibilityMatrixStateUnknown
		}
		entryKeys = append(entryKeys, verifierEcosystemValBCompatibilityEntryKey(entry))
	}
	if !containsExactTrimmedStringSet(entryKeys, verifierEcosystemValBRequiredCompatibilityEntryKeys()...) {
		return VerifierEcosystemValBCompatibilityMatrixStatePartial
	}
	return VerifierEcosystemValBCompatibilityMatrixStateActive
}

func EvaluateVerifierEcosystemValBSchemaProofCompatibilityState(model VerifierEcosystemValBSchemaProofCompatibility, matrix VerifierEcosystemValBCompatibilityMatrix) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CompatibilityID,
		model.CompatibilityMatrixRef,
		model.SchemaVersion,
		model.ProofType,
		model.VerifierVersion,
		model.TrustRootVersion,
		model.CompatibilityState,
		model.DerivedDiagnosticClass,
		model.ProjectionDisclaimer,
	) || len(model.RequiredDiagnostics) == 0 {
		return VerifierEcosystemValBSchemaProofCompatibilityStateIncomplete
	}
	if !verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValBSchemaProofCompatibilityStatePartial
	}
	if EvaluateVerifierEcosystemValBCompatibilityMatrixState(matrix) != VerifierEcosystemValBCompatibilityMatrixStateActive {
		return VerifierEcosystemValBSchemaProofCompatibilityStateBlocked
	}
	if !containsTrimmedString(matrix.SupportedSchemaVersions, model.SchemaVersion) ||
		!containsTrimmedString(matrix.SupportedProofTypes, model.ProofType) ||
		!containsTrimmedString(matrix.SupportedVerifierVersions, model.VerifierVersion) ||
		!containsTrimmedString(matrix.SupportedTrustRootVersions, model.TrustRootVersion) ||
		!containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), model.CompatibilityState) ||
		!containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), model.DerivedDiagnosticClass) {
		return VerifierEcosystemValBSchemaProofCompatibilityStateUnknown
	}
	switch strings.TrimSpace(model.CompatibilityState) {
	case ReferenceArchitectureCompatibilityUnknown:
		return VerifierEcosystemValBSchemaProofCompatibilityStateUnknown
	case ReferenceArchitectureCompatibilityUnsupported:
		return VerifierEcosystemValBSchemaProofCompatibilityStateBlocked
	}
	found := false
	var matched VerifierEcosystemValBCompatibilityEntry
	modelKey := strings.Join([]string{
		strings.TrimSpace(model.SchemaVersion),
		strings.TrimSpace(model.ProofType),
		strings.TrimSpace(model.VerifierVersion),
		strings.TrimSpace(model.TrustRootVersion),
	}, "|")
	for _, entry := range matrix.CompatibilityEntries {
		if verifierEcosystemValBCompatibilityEntryKey(entry) == modelKey {
			matched = entry
			found = true
			break
		}
	}
	if !found {
		return VerifierEcosystemValBSchemaProofCompatibilityStateBlocked
	}
	if strings.TrimSpace(model.CompatibilityState) != strings.TrimSpace(matched.CompatibilityState) ||
		!containsExactTrimmedStringSet(model.RequiredDiagnostics, matched.RequiredDiagnostics...) ||
		strings.TrimSpace(model.DerivedDiagnosticClass) != verifierEcosystemValBExpectedDiagnosticForCompatibilityState(matched.CompatibilityState) {
		return VerifierEcosystemValBSchemaProofCompatibilityStatePartial
	}
	if verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValBSchemaProofCompatibilityStateBlocked
	}
	switch strings.TrimSpace(model.CompatibilityState) {
	case ReferenceArchitectureCompatibilityCompatible:
		return VerifierEcosystemValBSchemaProofCompatibilityStateActive
	case ReferenceArchitectureCompatibilityCompatibleWithWarning:
		if len(model.Caveats) == 0 {
			return VerifierEcosystemValBSchemaProofCompatibilityStateBlocked
		}
		return VerifierEcosystemValBSchemaProofCompatibilityStatePartial
	case ReferenceArchitectureCompatibilityDeprecated:
		if len(model.Caveats) == 0 || strings.TrimSpace(model.MigrationOrSupersessionRef) == "" {
			return VerifierEcosystemValBSchemaProofCompatibilityStateBlocked
		}
		return VerifierEcosystemValBSchemaProofCompatibilityStatePartial
	case ReferenceArchitectureCompatibilitySuperseded:
		if len(model.Caveats) == 0 || strings.TrimSpace(model.MigrationOrSupersessionRef) == "" {
			return VerifierEcosystemValBSchemaProofCompatibilityStateBlocked
		}
		return VerifierEcosystemValBSchemaProofCompatibilityStatePartial
	case ReferenceArchitectureCompatibilityUnsupported:
		return VerifierEcosystemValBSchemaProofCompatibilityStateBlocked
	case ReferenceArchitectureCompatibilityUnknown:
		return VerifierEcosystemValBSchemaProofCompatibilityStateUnknown
	default:
		return VerifierEcosystemValBSchemaProofCompatibilityStateUnknown
	}
}

func EvaluateVerifierEcosystemValBMixedVersionDiagnosticsState(model VerifierEcosystemValBMixedVersionDiagnosticsCatalog) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.CatalogID, model.ProjectionDisclaimer) || len(model.Cases) == 0 {
		return VerifierEcosystemValBMixedVersionStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedCompatibilityStates, verifierEcosystemVal0CompatibilityResults()...) ||
		!containsExactTrimmedStringSet(model.SupportedDiagnosticClasses, verifierEcosystemVal0DiagnosticClasses()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValBMixedVersionStatePartial
	}
	caseIDs := make([]string, 0, len(model.Cases))
	for _, item := range model.Cases {
		if !referenceArchitectureValBRequiredRefsPresent(
			item.MixedVersionCaseID,
			item.SchemaVersion,
			item.ProofType,
			item.VerifierVersion,
			item.TrustRootVersion,
			item.ExpectedCompatibilityState,
			item.ExpectedDiagnosticClass,
			item.ProjectionDisclaimer,
		) {
			return VerifierEcosystemValBMixedVersionStateIncomplete
		}
		if !containsTrimmedString(model.SupportedCompatibilityStates, item.ExpectedCompatibilityState) ||
			!containsTrimmedString(model.SupportedDiagnosticClasses, item.ExpectedDiagnosticClass) ||
			!containsTrimmedString(VerifierEcosystemValBCompatibilityMatrixModel().SupportedSchemaVersions, item.SchemaVersion) ||
			!containsTrimmedString(verifierEcosystemVal0SupportedProofTypes(), item.ProofType) ||
			!containsTrimmedString(verifierEcosystemValBSupportedVerifierVersions(), item.VerifierVersion) ||
			!containsTrimmedString(verifierEcosystemValBSupportedTrustRootVersions(), item.TrustRootVersion) ||
			!verifierEcosystemVal0HasProjectionDisclaimer(item.ProjectionDisclaimer) {
			return VerifierEcosystemValBMixedVersionStateUnknown
		}
		if item.UniversalCompatibilityClaim || verifierEcosystemVal0HasOverclaim(strings.Join(item.Caveats, " "), item.ProjectionDisclaimer) {
			return VerifierEcosystemValBMixedVersionStateBlocked
		}
		switch strings.TrimSpace(item.ExpectedCompatibilityState) {
		case ReferenceArchitectureCompatibilityDeprecated:
			if len(item.Caveats) == 0 || strings.TrimSpace(item.MigrationOrSupersessionRef) == "" || strings.TrimSpace(item.ExpectedDiagnosticClass) != VerifierEcosystemDiagnosticCompatibilityWarning {
				return VerifierEcosystemValBMixedVersionStateBlocked
			}
		case ReferenceArchitectureCompatibilitySuperseded:
			if len(item.Caveats) == 0 || strings.TrimSpace(item.MigrationOrSupersessionRef) == "" || strings.TrimSpace(item.ExpectedDiagnosticClass) != VerifierEcosystemDiagnosticSupersededProof {
				return VerifierEcosystemValBMixedVersionStateBlocked
			}
		case ReferenceArchitectureCompatibilityCompatibleWithWarning:
			if len(item.Caveats) == 0 || strings.TrimSpace(item.ExpectedDiagnosticClass) != VerifierEcosystemDiagnosticCompatibilityWarning {
				return VerifierEcosystemValBMixedVersionStateBlocked
			}
		case ReferenceArchitectureCompatibilityCompatible:
			if strings.TrimSpace(item.ExpectedDiagnosticClass) != VerifierEcosystemDiagnosticVerified {
				return VerifierEcosystemValBMixedVersionStatePartial
			}
		case ReferenceArchitectureCompatibilityUnsupported:
			return VerifierEcosystemValBMixedVersionStateBlocked
		case ReferenceArchitectureCompatibilityUnknown:
			return VerifierEcosystemValBMixedVersionStateUnknown
		default:
			return VerifierEcosystemValBMixedVersionStateUnknown
		}
		caseIDs = append(caseIDs, item.MixedVersionCaseID)
	}
	if !containsExactTrimmedStringSet(caseIDs, verifierEcosystemValBRequiredMixedVersionCaseIDs()...) {
		return VerifierEcosystemValBMixedVersionStatePartial
	}
	return VerifierEcosystemValBMixedVersionStateActive
}

func EvaluateVerifierEcosystemValBDiagnosticPrecedenceState(model VerifierEcosystemValBDiagnosticPrecedence) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.PrecedenceID, model.DerivedDiagnosticClass, model.ProjectionDisclaimer) || len(model.ObservedDiagnostics) == 0 {
		return VerifierEcosystemValBDiagnosticPrecedenceStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedDiagnosticClasses, verifierEcosystemVal0DiagnosticClasses()...) ||
		!containsExactTrimmedStringSet(model.PrecedenceOrder, verifierEcosystemValBDiagnosticPrecedenceOrder()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValBDiagnosticPrecedenceStatePartial
	}
	if verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValBDiagnosticPrecedenceStateBlocked
	}
	if !containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), model.DerivedDiagnosticClass) {
		return VerifierEcosystemValBDiagnosticPrecedenceStateUnknown
	}
	expected := DeriveVerifierEcosystemValBDiagnostic(model.ObservedDiagnostics, model.Caveats)
	if expected == VerifierEcosystemDiagnosticUnknown {
		return VerifierEcosystemValBDiagnosticPrecedenceStateUnknown
	}
	if strings.TrimSpace(model.DerivedDiagnosticClass) != expected {
		return VerifierEcosystemValBDiagnosticPrecedenceStatePartial
	}
	if strings.TrimSpace(model.DerivedDiagnosticClass) == VerifierEcosystemDiagnosticCompatibilityWarning && len(model.Caveats) == 0 {
		return VerifierEcosystemValBDiagnosticPrecedenceStateBlocked
	}
	return VerifierEcosystemValBDiagnosticPrecedenceStateActive
}

func EvaluateVerifierEcosystemValBFixtureDescriptorState(model VerifierEcosystemValBFixtureCatalog) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.FixtureCatalogID, model.ProjectionDisclaimer) || len(model.Fixtures) == 0 {
		return VerifierEcosystemValBFixtureDescriptorStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedFixtureTypes, verifierEcosystemValBFixtureTypes()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValBFixtureDescriptorStatePartial
	}
	fixtureIDs := make([]string, 0, len(model.Fixtures))
	for _, item := range model.Fixtures {
		if !referenceArchitectureValBRequiredRefsPresent(
			item.FixtureID,
			item.FixtureType,
			item.SchemaVersion,
			item.ExpectedResult,
			item.ExpectedDiagnostic,
			item.ExpectedOutputBoundary,
			item.ProjectionDisclaimer,
		) || len(item.RequiredEvidenceRefs) == 0 {
			return VerifierEcosystemValBFixtureDescriptorStateIncomplete
		}
		if !containsTrimmedString(model.SupportedFixtureTypes, item.FixtureType) ||
			!verifierEcosystemValBFixtureProofTypeAllowed(item) ||
			!verifierEcosystemValBFixtureSchemaAllowed(item) ||
			!containsTrimmedString(verifierEcosystemValAOverallResults(), item.ExpectedResult) ||
			!containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), item.ExpectedDiagnostic) ||
			!containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), item.ExpectedOutputBoundary) ||
			!verifierEcosystemVal0HasProjectionDisclaimer(item.ProjectionDisclaimer) {
			return VerifierEcosystemValBFixtureDescriptorStateUnknown
		}
		if item.ProductionEvidenceClaim || verifierEcosystemVal0HasOverclaim(strings.Join(item.Caveats, " "), item.ProjectionDisclaimer) {
			return VerifierEcosystemValBFixtureDescriptorStateBlocked
		}
		for _, evidenceRef := range item.RequiredEvidenceRefs {
			if strings.TrimSpace(evidenceRef) == "" {
				return VerifierEcosystemValBFixtureDescriptorStatePartial
			}
		}
		expectedResult, expectedDiagnostic := verifierEcosystemValBFixtureExpectedOutcome(item.FixtureType)
		if strings.TrimSpace(item.ExpectedResult) != expectedResult || strings.TrimSpace(item.ExpectedDiagnostic) != expectedDiagnostic {
			return VerifierEcosystemValBFixtureDescriptorStatePartial
		}
		fixtureIDs = append(fixtureIDs, item.FixtureID)
	}
	if !containsExactTrimmedStringSet(fixtureIDs, verifierEcosystemValBRequiredFixtureIDs()...) {
		return VerifierEcosystemValBFixtureDescriptorStatePartial
	}
	return VerifierEcosystemValBFixtureDescriptorStateActive
}

func EvaluateVerifierEcosystemValBConformanceCaseState(model VerifierEcosystemValBConformanceCaseCatalog, fixtures VerifierEcosystemValBFixtureCatalog, outputs VerifierEcosystemValBOutputClassCatalog) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.ConformanceCatalogID, model.ProjectionDisclaimer) || len(model.Cases) == 0 {
		return VerifierEcosystemValBConformanceCaseStateIncomplete
	}
	if !verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValBConformanceCaseStatePartial
	}
	if EvaluateVerifierEcosystemValBFixtureDescriptorState(fixtures) != VerifierEcosystemValBFixtureDescriptorStateActive ||
		EvaluateVerifierEcosystemValBOutputClassState(outputs) != VerifierEcosystemValBOutputClassStateActive {
		return VerifierEcosystemValBConformanceCaseStateBlocked
	}
	fixtureLookup := make(map[string]VerifierEcosystemValBFixtureDescriptor, len(fixtures.Fixtures))
	for _, fixture := range fixtures.Fixtures {
		fixtureLookup[strings.TrimSpace(fixture.FixtureID)] = fixture
	}
	caseIDs := make([]string, 0, len(model.Cases))
	for _, item := range model.Cases {
		if !referenceArchitectureValBRequiredRefsPresent(
			item.ConformanceCaseID,
			item.FixtureRef,
			item.VerifierContractRef,
			item.InputRef,
			item.ExpectedOverallResult,
			item.ExpectedDiagnosticClass,
			item.ExpectedOutputClass,
			item.ProjectionDisclaimer,
		) || len(item.RequiredFields) == 0 || len(item.ForbiddenClaims) == 0 {
			return VerifierEcosystemValBConformanceCaseStateIncomplete
		}
		if !containsExactTrimmedStringSet(item.RequiredFields, verifierEcosystemValBRequiredConformanceFields()...) ||
			!containsExactTrimmedStringSet(item.ForbiddenClaims, verifierEcosystemValBRequiredForbiddenClaims()...) ||
			!containsTrimmedString(verifierEcosystemValAOverallResults(), item.ExpectedOverallResult) ||
			!containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), item.ExpectedDiagnosticClass) ||
			!containsTrimmedString(verifierEcosystemValBOutputClasses(), item.ExpectedOutputClass) ||
			!verifierEcosystemVal0HasProjectionDisclaimer(item.ProjectionDisclaimer) {
			return VerifierEcosystemValBConformanceCaseStateUnknown
		}
		if verifierEcosystemVal0HasOverclaim(strings.Join(item.ObservedClaims, " "), strings.Join(item.Caveats, " "), item.ProjectionDisclaimer) {
			return VerifierEcosystemValBConformanceCaseStateBlocked
		}
		for _, claim := range item.ObservedClaims {
			if containsTrimmedString(item.ForbiddenClaims, claim) {
				return VerifierEcosystemValBConformanceCaseStateBlocked
			}
		}
		fixture, ok := fixtureLookup[strings.TrimSpace(item.FixtureRef)]
		if !ok {
			return VerifierEcosystemValBConformanceCaseStateIncomplete
		}
		if strings.TrimSpace(item.ExpectedOverallResult) != strings.TrimSpace(fixture.ExpectedResult) ||
			strings.TrimSpace(item.ExpectedDiagnosticClass) != strings.TrimSpace(fixture.ExpectedDiagnostic) {
			return VerifierEcosystemValBConformanceCaseStatePartial
		}
		if strings.TrimSpace(item.ExpectedOutputClass) != deriveVerifierEcosystemValBOutputClass(item.ExpectedOverallResult, item.ExpectedDiagnosticClass, item.Caveats) {
			return VerifierEcosystemValBConformanceCaseStatePartial
		}
		caseIDs = append(caseIDs, item.ConformanceCaseID)
	}
	if !containsExactTrimmedStringSet(caseIDs, verifierEcosystemValBRequiredConformanceCaseIDs()...) {
		return VerifierEcosystemValBConformanceCaseStatePartial
	}
	return VerifierEcosystemValBConformanceCaseStateActive
}

func EvaluateVerifierEcosystemValBConformanceSuiteState(model VerifierEcosystemValBConformanceSuite, cases VerifierEcosystemValBConformanceCaseCatalog, fixtures VerifierEcosystemValBFixtureCatalog, outputs VerifierEcosystemValBOutputClassCatalog) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.ConformanceSuiteID, model.ProjectionDisclaimer) || len(model.RequiredCaseRefs) == 0 || len(model.RequiredFixtureRefs) == 0 {
		return VerifierEcosystemValBConformanceSuiteStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredCaseRefs, verifierEcosystemValBRequiredConformanceCaseIDs()...) ||
		!containsExactTrimmedStringSet(model.RequiredFixtureRefs, verifierEcosystemValBRequiredFixtureIDs()...) {
		return VerifierEcosystemValBConformanceSuiteStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedOutputClasses, verifierEcosystemValBOutputClasses()...) ||
		!containsExactTrimmedStringSet(model.SupportedDiagnosticClasses, verifierEcosystemVal0DiagnosticClasses()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValBConformanceSuiteStatePartial
	}
	if model.CertificationClaim || verifierEcosystemVal0HasOverclaim(strings.Join(model.Caveats, " "), model.ProjectionDisclaimer) {
		return VerifierEcosystemValBConformanceSuiteStateBlocked
	}
	actualCaseRefs := make([]string, 0, len(cases.Cases))
	for _, item := range cases.Cases {
		actualCaseRefs = append(actualCaseRefs, item.ConformanceCaseID)
	}
	if !containsAllTrimmedStrings(actualCaseRefs, model.RequiredCaseRefs...) {
		return VerifierEcosystemValBConformanceSuiteStateIncomplete
	}
	actualFixtureRefs := make([]string, 0, len(fixtures.Fixtures))
	for _, item := range fixtures.Fixtures {
		actualFixtureRefs = append(actualFixtureRefs, item.FixtureID)
	}
	if !containsAllTrimmedStrings(actualFixtureRefs, model.RequiredFixtureRefs...) {
		return VerifierEcosystemValBConformanceSuiteStateIncomplete
	}
	caseState := EvaluateVerifierEcosystemValBConformanceCaseState(cases, fixtures, outputs)
	if caseState != VerifierEcosystemValBConformanceCaseStateActive {
		switch caseState {
		case VerifierEcosystemValBConformanceCaseStateBlocked:
			return VerifierEcosystemValBConformanceSuiteStateBlocked
		case VerifierEcosystemValBConformanceCaseStateUnknown:
			return VerifierEcosystemValBConformanceSuiteStateUnknown
		case VerifierEcosystemValBConformanceCaseStateIncomplete:
			return VerifierEcosystemValBConformanceSuiteStateIncomplete
		default:
			return VerifierEcosystemValBConformanceSuiteStatePartial
		}
	}
	if EvaluateVerifierEcosystemValBFixtureDescriptorState(fixtures) != VerifierEcosystemValBFixtureDescriptorStateActive ||
		EvaluateVerifierEcosystemValBOutputClassState(outputs) != VerifierEcosystemValBOutputClassStateActive {
		return VerifierEcosystemValBConformanceSuiteStateBlocked
	}
	return VerifierEcosystemValBConformanceSuiteStateActive
}

func EvaluateVerifierEcosystemValBOutputClassState(model VerifierEcosystemValBOutputClassCatalog) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.OutputClassCatalogID, model.ProjectionDisclaimer) || len(model.Mappings) == 0 {
		return VerifierEcosystemValBOutputClassStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedOutputClasses, verifierEcosystemValBOutputClasses()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValBOutputClassStatePartial
	}
	mappingKeys := make([]string, 0, len(model.Mappings))
	for _, item := range model.Mappings {
		if !referenceArchitectureValBRequiredRefsPresent(
			item.OutputClassID,
			item.OverallResult,
			item.DiagnosticClass,
			item.OutputClass,
			item.ProjectionDisclaimer,
		) {
			return VerifierEcosystemValBOutputClassStateIncomplete
		}
		if !containsTrimmedString(verifierEcosystemValAOverallResults(), item.OverallResult) ||
			!containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), item.DiagnosticClass) ||
			!containsTrimmedString(verifierEcosystemValBOutputClasses(), item.OutputClass) ||
			!verifierEcosystemVal0HasProjectionDisclaimer(item.ProjectionDisclaimer) {
			return VerifierEcosystemValBOutputClassStateUnknown
		}
		if verifierEcosystemVal0HasOverclaim(strings.Join(item.Caveats, " "), item.ProjectionDisclaimer) {
			return VerifierEcosystemValBOutputClassStateBlocked
		}
		expected := deriveVerifierEcosystemValBOutputClass(item.OverallResult, item.DiagnosticClass, item.Caveats)
		if strings.TrimSpace(item.OutputClass) != expected {
			return VerifierEcosystemValBOutputClassStatePartial
		}
		mappingKeys = append(mappingKeys, strings.TrimSpace(item.OverallResult)+"|"+strings.TrimSpace(item.DiagnosticClass))
	}
	if !containsExactTrimmedStringSet(mappingKeys, verifierEcosystemValBRequiredOutputClassMappingKeys()...) {
		return VerifierEcosystemValBOutputClassStatePartial
	}
	return VerifierEcosystemValBOutputClassStateActive
}

func verifierEcosystemValBPoint6DependencyHealthy(snapshot VerifierEcosystemValBDependencySnapshot) bool {
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

func verifierEcosystemValBStateSeverity(state, active, partial, incomplete, blocked, unknown string) int {
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

func EvaluateVerifierEcosystemValBState(
	dependency VerifierEcosystemValBDependencySnapshot,
	compatibilityMatrixState, schemaProofCompatibilityState, mixedVersionState, diagnosticPrecedenceState, fixtureDescriptorState, conformanceCaseState, conformanceSuiteState, outputClassState string,
) string {
	if !verifierEcosystemValBPoint6DependencyHealthy(dependency) ||
		strings.TrimSpace(dependency.Val0CurrentState) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(dependency.Val0State) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(dependency.ValACurrentState) != VerifierEcosystemValAStateActive ||
		strings.TrimSpace(dependency.ValAState) != VerifierEcosystemValAStateActive ||
		strings.TrimSpace(dependency.Point7State) != VerifierEcosystemPoint7StateNotComplete {
		return VerifierEcosystemValBStateBlocked
	}
	highestSeverity := 0
	componentSeverities := []int{
		verifierEcosystemValBStateSeverity(compatibilityMatrixState, VerifierEcosystemValBCompatibilityMatrixStateActive, VerifierEcosystemValBCompatibilityMatrixStatePartial, VerifierEcosystemValBCompatibilityMatrixStateIncomplete, VerifierEcosystemValBCompatibilityMatrixStateBlocked, VerifierEcosystemValBCompatibilityMatrixStateUnknown),
		verifierEcosystemValBStateSeverity(schemaProofCompatibilityState, VerifierEcosystemValBSchemaProofCompatibilityStateActive, VerifierEcosystemValBSchemaProofCompatibilityStatePartial, VerifierEcosystemValBSchemaProofCompatibilityStateIncomplete, VerifierEcosystemValBSchemaProofCompatibilityStateBlocked, VerifierEcosystemValBSchemaProofCompatibilityStateUnknown),
		verifierEcosystemValBStateSeverity(mixedVersionState, VerifierEcosystemValBMixedVersionStateActive, VerifierEcosystemValBMixedVersionStatePartial, VerifierEcosystemValBMixedVersionStateIncomplete, VerifierEcosystemValBMixedVersionStateBlocked, VerifierEcosystemValBMixedVersionStateUnknown),
		verifierEcosystemValBStateSeverity(diagnosticPrecedenceState, VerifierEcosystemValBDiagnosticPrecedenceStateActive, VerifierEcosystemValBDiagnosticPrecedenceStatePartial, VerifierEcosystemValBDiagnosticPrecedenceStateIncomplete, VerifierEcosystemValBDiagnosticPrecedenceStateBlocked, VerifierEcosystemValBDiagnosticPrecedenceStateUnknown),
		verifierEcosystemValBStateSeverity(fixtureDescriptorState, VerifierEcosystemValBFixtureDescriptorStateActive, VerifierEcosystemValBFixtureDescriptorStatePartial, VerifierEcosystemValBFixtureDescriptorStateIncomplete, VerifierEcosystemValBFixtureDescriptorStateBlocked, VerifierEcosystemValBFixtureDescriptorStateUnknown),
		verifierEcosystemValBStateSeverity(conformanceCaseState, VerifierEcosystemValBConformanceCaseStateActive, VerifierEcosystemValBConformanceCaseStatePartial, VerifierEcosystemValBConformanceCaseStateIncomplete, VerifierEcosystemValBConformanceCaseStateBlocked, VerifierEcosystemValBConformanceCaseStateUnknown),
		verifierEcosystemValBStateSeverity(conformanceSuiteState, VerifierEcosystemValBConformanceSuiteStateActive, VerifierEcosystemValBConformanceSuiteStatePartial, VerifierEcosystemValBConformanceSuiteStateIncomplete, VerifierEcosystemValBConformanceSuiteStateBlocked, VerifierEcosystemValBConformanceSuiteStateUnknown),
		verifierEcosystemValBStateSeverity(outputClassState, VerifierEcosystemValBOutputClassStateActive, VerifierEcosystemValBOutputClassStatePartial, VerifierEcosystemValBOutputClassStateIncomplete, VerifierEcosystemValBOutputClassStateBlocked, VerifierEcosystemValBOutputClassStateUnknown),
	}
	for _, severity := range componentSeverities {
		if severity > highestSeverity {
			highestSeverity = severity
		}
	}
	switch highestSeverity {
	case 4:
		return VerifierEcosystemValBStateBlocked
	case 3:
		return VerifierEcosystemValBStateUnknown
	case 2:
		return VerifierEcosystemValBStateIncomplete
	case 1:
		return VerifierEcosystemValBStatePartial
	default:
		return VerifierEcosystemValBStateActive
	}
}

func VerifierEcosystemValBProofSurfaceRefs() []string {
	return []string{
		"/v1/verifier-ecosystem/val0/proofs",
		"/v1/verifier-ecosystem/vala/proofs",
		"/v1/verifier-ecosystem/valb/compatibility-matrix",
		"/v1/verifier-ecosystem/valb/schema-proof-compatibility",
		"/v1/verifier-ecosystem/valb/mixed-version-diagnostics",
		"/v1/verifier-ecosystem/valb/diagnostic-precedence",
		"/v1/verifier-ecosystem/valb/fixture-descriptors",
		"/v1/verifier-ecosystem/valb/conformance-cases",
		"/v1/verifier-ecosystem/valb/conformance-suite",
		"/v1/verifier-ecosystem/valb/output-classes",
		"/v1/verifier-ecosystem/valb/proofs",
	}
}

func EvaluateVerifierEcosystemValBProofsState(
	currentState string,
	point7State string,
	val0CurrentState string,
	valACurrentState string,
	surfaceRefs, evidenceRefs, limitations, whyPoint7NotPass []string,
	projectionDisclaimer string,
) string {
	baseState := strings.TrimSpace(currentState)
	if strings.TrimSpace(val0CurrentState) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(valACurrentState) != VerifierEcosystemValAStateActive ||
		!containsExactTrimmedStringSet(surfaceRefs, VerifierEcosystemValBProofSurfaceRefs()...) ||
		!verifierEcosystemValBProofEvidenceQualityValid(VerifierEcosystemValBVerifierEvidence(), evidenceRefs) ||
		len(limitations) == 0 ||
		len(whyPoint7NotPass) == 0 ||
		!verifierEcosystemVal0HasProjectionDisclaimer(projectionDisclaimer) {
		if baseState == VerifierEcosystemValBStateActive {
			return VerifierEcosystemValBStatePartial
		}
		return baseState
	}
	if baseState == VerifierEcosystemValBStateActive && strings.TrimSpace(point7State) != VerifierEcosystemPoint7StateNotComplete {
		return VerifierEcosystemValBStatePartial
	}
	return baseState
}

func EvaluateVerifierEcosystemValBPoint7State(valBState string) string {
	_ = valBState
	return VerifierEcosystemPoint7StateNotComplete
}

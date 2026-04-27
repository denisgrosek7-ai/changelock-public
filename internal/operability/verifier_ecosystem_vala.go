package operability

import "strings"

const (
	VerifierEcosystemValAInputStateActive     = "verifier_ecosystem_vala_input_model_active"
	VerifierEcosystemValAInputStatePartial    = "verifier_ecosystem_vala_input_model_partial"
	VerifierEcosystemValAInputStateIncomplete = "verifier_ecosystem_vala_input_model_incomplete"
	VerifierEcosystemValAInputStateBlocked    = "verifier_ecosystem_vala_input_model_blocked"
	VerifierEcosystemValAInputStateUnknown    = "verifier_ecosystem_vala_input_model_unknown"

	VerifierEcosystemValAEngineStateActive     = "verifier_ecosystem_vala_engine_active"
	VerifierEcosystemValAEngineStatePartial    = "verifier_ecosystem_vala_engine_partial"
	VerifierEcosystemValAEngineStateIncomplete = "verifier_ecosystem_vala_engine_incomplete"
	VerifierEcosystemValAEngineStateBlocked    = "verifier_ecosystem_vala_engine_blocked"
	VerifierEcosystemValAEngineStateUnknown    = "verifier_ecosystem_vala_engine_unknown"

	VerifierEcosystemValAResultStateActive     = "verifier_ecosystem_vala_result_active"
	VerifierEcosystemValAResultStatePartial    = "verifier_ecosystem_vala_result_partial"
	VerifierEcosystemValAResultStateIncomplete = "verifier_ecosystem_vala_result_incomplete"
	VerifierEcosystemValAResultStateBlocked    = "verifier_ecosystem_vala_result_blocked"
	VerifierEcosystemValAResultStateUnknown    = "verifier_ecosystem_vala_result_unknown"

	VerifierEcosystemValADiagnosticsMappingStateActive     = "verifier_ecosystem_vala_diagnostics_mapping_active"
	VerifierEcosystemValADiagnosticsMappingStatePartial    = "verifier_ecosystem_vala_diagnostics_mapping_partial"
	VerifierEcosystemValADiagnosticsMappingStateIncomplete = "verifier_ecosystem_vala_diagnostics_mapping_incomplete"
	VerifierEcosystemValADiagnosticsMappingStateBlocked    = "verifier_ecosystem_vala_diagnostics_mapping_blocked"
	VerifierEcosystemValADiagnosticsMappingStateUnknown    = "verifier_ecosystem_vala_diagnostics_mapping_unknown"

	VerifierEcosystemValACommandContractStateActive     = "verifier_ecosystem_vala_command_contract_active"
	VerifierEcosystemValACommandContractStatePartial    = "verifier_ecosystem_vala_command_contract_partial"
	VerifierEcosystemValACommandContractStateIncomplete = "verifier_ecosystem_vala_command_contract_incomplete"
	VerifierEcosystemValACommandContractStateBlocked    = "verifier_ecosystem_vala_command_contract_blocked"
	VerifierEcosystemValACommandContractStateUnknown    = "verifier_ecosystem_vala_command_contract_unknown"

	VerifierEcosystemValASDKEntrypointStateActive     = "verifier_ecosystem_vala_sdk_entrypoint_active"
	VerifierEcosystemValASDKEntrypointStatePartial    = "verifier_ecosystem_vala_sdk_entrypoint_partial"
	VerifierEcosystemValASDKEntrypointStateIncomplete = "verifier_ecosystem_vala_sdk_entrypoint_incomplete"
	VerifierEcosystemValASDKEntrypointStateBlocked    = "verifier_ecosystem_vala_sdk_entrypoint_blocked"
	VerifierEcosystemValASDKEntrypointStateUnknown    = "verifier_ecosystem_vala_sdk_entrypoint_unknown"

	VerifierEcosystemValAStateActive     = "verifier_ecosystem_vala_active"
	VerifierEcosystemValAStatePartial    = "verifier_ecosystem_vala_partial"
	VerifierEcosystemValAStateIncomplete = "verifier_ecosystem_vala_incomplete"
	VerifierEcosystemValAStateBlocked    = "verifier_ecosystem_vala_blocked"
	VerifierEcosystemValAStateUnknown    = "verifier_ecosystem_vala_unknown"

	VerifierEcosystemValACommandReportFormatJSON        = "json"
	VerifierEcosystemValACommandReportFormatTextSummary = "text_summary"
	VerifierEcosystemValACommandReportFormatUnknown     = "unknown"

	VerifierEcosystemValADigestMatch       = "digest_match"
	VerifierEcosystemValADigestMismatch    = "digest_mismatch"
	VerifierEcosystemValADigestMissing     = "digest_missing"
	VerifierEcosystemValADigestUnsupported = "digest_unsupported"

	VerifierEcosystemValASignatureValid       = "signature_valid"
	VerifierEcosystemValASignatureInvalid     = "signature_invalid"
	VerifierEcosystemValASignatureMissing     = "signature_missing"
	VerifierEcosystemValASignatureUnsupported = "signature_unsupported"

	VerifierEcosystemValASchemaValid       = "schema_valid"
	VerifierEcosystemValASchemaMismatch    = "schema_mismatch"
	VerifierEcosystemValASchemaUnsupported = "schema_unsupported"

	VerifierEcosystemValAScopeValid       = "scope_valid"
	VerifierEcosystemValAScopeMismatch    = "scope_mismatch"
	VerifierEcosystemValAScopeUnsupported = "scope_unsupported"

	VerifierEcosystemValAIssuerTrusted = "issuer_trusted"
	VerifierEcosystemValAIssuerRevoked = "issuer_revoked"
	VerifierEcosystemValAIssuerUnknown = "issuer_unknown"

	VerifierEcosystemValARevocationNotRevoked      = "not_revoked"
	VerifierEcosystemValARevocationRevoked         = "revoked"
	VerifierEcosystemValARevocationMaterialMissing = "revocation_material_missing"
	VerifierEcosystemValARevocationUnknown         = "unknown"
	VerifierEcosystemValASupersessionCurrent       = "current"
	VerifierEcosystemValASupersessionSuperseded    = "superseded"
	VerifierEcosystemValASupersessionUnknown       = "unknown"
	VerifierEcosystemValALineagePresent            = "lineage_present"
	VerifierEcosystemValALineageMissing            = "lineage_missing"
	VerifierEcosystemValALineageUnknown            = "lineage_unknown"
	VerifierEcosystemValAOutputBoundaryValid       = "output_boundary_valid"
	VerifierEcosystemValAOutputBoundaryViolation   = "output_boundary_violation"
	VerifierEcosystemValAOutputBoundaryUnknown     = "output_boundary_unknown"
	VerifierEcosystemValAOverallResultVerified     = "verified"
	VerifierEcosystemValAOverallResultWarnings     = "verified_with_warnings"
	VerifierEcosystemValAOverallResultInvalid      = "invalid"
	VerifierEcosystemValAOverallResultIncomplete   = "incomplete"
	VerifierEcosystemValAOverallResultUnsupported  = "unsupported"
	VerifierEcosystemValAOverallResultStale        = "stale"
	VerifierEcosystemValAOverallResultRevoked      = "revoked"
	VerifierEcosystemValAOverallResultSuperseded   = "superseded"
	VerifierEcosystemValAOverallResultUnknown      = "unknown"
)

type VerifierEcosystemValADependencySnapshot struct {
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
	Point7State                    string `json:"point_7_state"`
}

type VerifierEcosystemValAReferenceVerifierInput struct {
	CurrentState                    string   `json:"current_state"`
	VerificationRequestID           string   `json:"verification_request_id"`
	VerifierContractRef             string   `json:"verifier_contract_ref"`
	ProofEnvelopeRef                string   `json:"proof_envelope_ref"`
	ArtifactRef                     string   `json:"artifact_ref"`
	ArtifactDigest                  string   `json:"artifact_digest"`
	ArtifactDigestAlgorithm         string   `json:"artifact_digest_algorithm"`
	SignatureRef                    string   `json:"signature_ref"`
	IssuerRef                       string   `json:"issuer_ref"`
	TrustRootRef                    string   `json:"trust_root_ref"`
	SchemaVersion                   string   `json:"schema_version"`
	ProofType                       string   `json:"proof_type"`
	RequestedScope                  string   `json:"requested_scope"`
	VerificationTime                string   `json:"verification_time"`
	ExpectedOutputBoundary          string   `json:"expected_output_boundary"`
	CompatibilityPolicyRef          string   `json:"compatibility_policy_ref"`
	RevocationMaterialRef           string   `json:"revocation_material_ref"`
	SupersessionMaterialRef         string   `json:"supersession_material_ref"`
	EvidenceRefs                    []string `json:"evidence_refs,omitempty"`
	DigestVerificationState         string   `json:"digest_verification_state"`
	SignatureVerificationState      string   `json:"signature_verification_state"`
	SchemaVerificationState         string   `json:"schema_verification_state"`
	ScopeVerificationState          string   `json:"scope_verification_state"`
	FreshnessVerificationState      string   `json:"freshness_verification_state"`
	TrustRootVerificationState      string   `json:"trust_root_verification_state"`
	IssuerVerificationState         string   `json:"issuer_verification_state"`
	CompatibilityEvaluationState    string   `json:"compatibility_evaluation_state"`
	RevocationEvaluationState       string   `json:"revocation_evaluation_state"`
	SupersessionEvaluationState     string   `json:"supersession_evaluation_state"`
	LineageVerificationState        string   `json:"lineage_verification_state"`
	OutputBoundaryVerificationState string   `json:"output_boundary_verification_state"`
	StrictFailClosed                bool     `json:"strict_fail_closed"`
	CanonicalEvidenceSpineAccess    bool     `json:"canonical_evidence_spine_access"`
	TruthOutsideScopeClaim          bool     `json:"truth_outside_scope_claim"`
	ClaimsActualCryptoVerification  bool     `json:"claims_actual_crypto_verification"`
	Caveats                         []string `json:"caveats,omitempty"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValAReferenceVerifierEngine struct {
	CurrentState               string   `json:"current_state"`
	EngineID                   string   `json:"engine_id"`
	Version                    string   `json:"version"`
	InputModelRef              string   `json:"input_model_ref"`
	ResultModelRef             string   `json:"result_model_ref"`
	VerifierContractState      string   `json:"verifier_contract_state"`
	ProofEnvelopeState         string   `json:"proof_envelope_state"`
	RequestedScopeState        string   `json:"requested_scope_state"`
	SchemaCompatibilityState   string   `json:"schema_compatibility_state"`
	TrustRootIssuerState       string   `json:"trust_root_issuer_state"`
	DiagnosticsState           string   `json:"diagnostics_state"`
	OutputBoundaryState        string   `json:"output_boundary_state"`
	SupportedDigestAlgorithms  []string `json:"supported_digest_algorithms,omitempty"`
	SupportedOverallResults    []string `json:"supported_overall_results,omitempty"`
	DeterministicOutput        bool     `json:"deterministic_output"`
	ExplicitFixtureSemantics   bool     `json:"explicit_fixture_semantics"`
	UsesRealCryptoPrimitives   bool     `json:"uses_real_crypto_primitives"`
	NetworkDependency          bool     `json:"network_dependency"`
	MutatesEvidence            bool     `json:"mutates_evidence"`
	ClaimsDeploymentApproval   bool     `json:"claims_deployment_approval"`
	ClaimsCanonicalTruth       bool     `json:"claims_canonical_truth"`
	ClaimsCertification        bool     `json:"claims_certification"`
	ClaimsActualCryptoValidity bool     `json:"claims_actual_crypto_validity"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValAVerificationResult struct {
	CurrentState           string   `json:"current_state"`
	VerificationResultID   string   `json:"verification_result_id"`
	RequestID              string   `json:"request_id"`
	VerifierVersion        string   `json:"verifier_version"`
	ProofType              string   `json:"proof_type"`
	SchemaVersion          string   `json:"schema_version"`
	Scope                  string   `json:"scope"`
	OutputBoundary         string   `json:"output_boundary"`
	OverallResult          string   `json:"overall_result"`
	DiagnosticClass        string   `json:"diagnostic_class"`
	DigestResult           string   `json:"digest_result"`
	SignatureResult        string   `json:"signature_result"`
	SchemaResult           string   `json:"schema_result"`
	ScopeResult            string   `json:"scope_result"`
	FreshnessResult        string   `json:"freshness_result"`
	TrustRootResult        string   `json:"trust_root_result"`
	IssuerResult           string   `json:"issuer_result"`
	CompatibilityResult    string   `json:"compatibility_result"`
	RevocationResult       string   `json:"revocation_result"`
	SupersessionResult     string   `json:"supersession_result"`
	LineageResult          string   `json:"lineage_result"`
	OutputBoundaryResult   string   `json:"output_boundary_result"`
	EvidenceRefs           []string `json:"evidence_refs,omitempty"`
	Caveats                []string `json:"caveats,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
	ProjectionDisclaimer   string   `json:"projection_disclaimer"`
	VerifiedAt             string   `json:"verified_at"`
	TruthOutsideScopeClaim bool     `json:"truth_outside_scope_claim"`
}

type VerifierEcosystemValADiagnosticsMapping struct {
	CurrentState               string   `json:"current_state"`
	MappingID                  string   `json:"mapping_id"`
	ProofType                  string   `json:"proof_type"`
	SupportedDiagnosticClasses []string `json:"supported_diagnostic_classes,omitempty"`
	DeterministicPrecedence    []string `json:"deterministic_precedence,omitempty"`
	DigestResult               string   `json:"digest_result"`
	SignatureResult            string   `json:"signature_result"`
	SchemaResult               string   `json:"schema_result"`
	ScopeResult                string   `json:"scope_result"`
	FreshnessResult            string   `json:"freshness_result"`
	TrustRootResult            string   `json:"trust_root_result"`
	IssuerResult               string   `json:"issuer_result"`
	CompatibilityResult        string   `json:"compatibility_result"`
	RevocationResult           string   `json:"revocation_result"`
	SupersessionResult         string   `json:"supersession_result"`
	LineageResult              string   `json:"lineage_result"`
	OutputBoundaryResult       string   `json:"output_boundary_result"`
	DerivedDiagnosticClass     string   `json:"derived_diagnostic_class"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type VerifierEcosystemValACommandContract struct {
	CurrentState                string `json:"current_state"`
	CommandContractID           string `json:"command_contract_id"`
	ProofEnvelopeInputRef       string `json:"proof_envelope_input_ref"`
	ArtifactDigestDescriptor    string `json:"artifact_digest_descriptor"`
	SignatureDescriptor         string `json:"signature_descriptor"`
	TrustRootMaterialDescriptor string `json:"trust_root_material_descriptor"`
	RequestedScope              string `json:"requested_scope"`
	OutputBoundary              string `json:"output_boundary"`
	ReportFormat                string `json:"report_format"`
	StrictFailClosedMode        bool   `json:"strict_fail_closed_mode"`
	ExampleCommand              string `json:"example_command"`
	MutatesEvidence             bool   `json:"mutates_evidence"`
	ApprovesDeployment          bool   `json:"approves_deployment"`
	SuppressesFailures          bool   `json:"suppresses_failures"`
	PublishesClaims             bool   `json:"publishes_claims"`
	CertificationClaim          bool   `json:"certification_claim"`
	ProjectionDisclaimer        string `json:"projection_disclaimer"`
}

type VerifierEcosystemValASDKEntrypoint struct {
	CurrentState                 string `json:"current_state"`
	EntrypointID                 string `json:"entrypoint_id"`
	PackagePath                  string `json:"package_path"`
	FunctionName                 string `json:"function_name"`
	InputModelRef                string `json:"input_model_ref"`
	ResultModelRef               string `json:"result_model_ref"`
	DeterministicOutput          bool   `json:"deterministic_output"`
	HiddenMainInstanceDependency bool   `json:"hidden_main_instance_dependency"`
	NetworkDependency            bool   `json:"network_dependency"`
	MutatesCanonicalState        bool   `json:"mutates_canonical_state"`
	GeneratesFakeEvidence        bool   `json:"generates_fake_evidence"`
	ProjectionDisclaimer         string `json:"projection_disclaimer"`
}

func verifierEcosystemValAProjectionDisclaimer() string {
	return "projection_only not_canonical_truth reference_verifier_tooling advisory_projection"
}

func verifierEcosystemValASupportedDigestAlgorithms() []string {
	return []string{"sha256", "sha512"}
}

func verifierEcosystemValASupportedReportFormats() []string {
	return []string{
		VerifierEcosystemValACommandReportFormatJSON,
		VerifierEcosystemValACommandReportFormatTextSummary,
		VerifierEcosystemValACommandReportFormatUnknown,
	}
}

func verifierEcosystemValAOverallResults() []string {
	return []string{
		VerifierEcosystemValAOverallResultVerified,
		VerifierEcosystemValAOverallResultWarnings,
		VerifierEcosystemValAOverallResultInvalid,
		VerifierEcosystemValAOverallResultIncomplete,
		VerifierEcosystemValAOverallResultUnsupported,
		VerifierEcosystemValAOverallResultStale,
		VerifierEcosystemValAOverallResultRevoked,
		VerifierEcosystemValAOverallResultSuperseded,
		VerifierEcosystemValAOverallResultUnknown,
	}
}

func verifierEcosystemValADigestResults() []string {
	return []string{
		VerifierEcosystemValADigestMatch,
		VerifierEcosystemValADigestMismatch,
		VerifierEcosystemValADigestMissing,
		VerifierEcosystemValADigestUnsupported,
	}
}

func verifierEcosystemValASignatureResults() []string {
	return []string{
		VerifierEcosystemValASignatureValid,
		VerifierEcosystemValASignatureInvalid,
		VerifierEcosystemValASignatureMissing,
		VerifierEcosystemValASignatureUnsupported,
	}
}

func verifierEcosystemValASchemaResults() []string {
	return []string{
		VerifierEcosystemValASchemaValid,
		VerifierEcosystemValASchemaMismatch,
		VerifierEcosystemValASchemaUnsupported,
	}
}

func verifierEcosystemValAScopeResults() []string {
	return []string{
		VerifierEcosystemValAScopeValid,
		VerifierEcosystemValAScopeMismatch,
		VerifierEcosystemValAScopeUnsupported,
	}
}

func verifierEcosystemValAIssuerResults() []string {
	return []string{
		VerifierEcosystemValAIssuerTrusted,
		VerifierEcosystemValAIssuerRevoked,
		VerifierEcosystemValAIssuerUnknown,
	}
}

func verifierEcosystemValARevocationResults() []string {
	return []string{
		VerifierEcosystemValARevocationNotRevoked,
		VerifierEcosystemValARevocationRevoked,
		VerifierEcosystemValARevocationMaterialMissing,
		VerifierEcosystemValARevocationUnknown,
	}
}

func verifierEcosystemValASupersessionResults() []string {
	return []string{
		VerifierEcosystemValASupersessionCurrent,
		VerifierEcosystemValASupersessionSuperseded,
		VerifierEcosystemValASupersessionUnknown,
	}
}

func verifierEcosystemValALineageResults() []string {
	return []string{
		VerifierEcosystemValALineagePresent,
		VerifierEcosystemValALineageMissing,
		VerifierEcosystemValALineageUnknown,
	}
}

func verifierEcosystemValAOutputBoundaryResults() []string {
	return []string{
		VerifierEcosystemValAOutputBoundaryValid,
		VerifierEcosystemValAOutputBoundaryViolation,
		VerifierEcosystemValAOutputBoundaryUnknown,
	}
}

func verifierEcosystemValADiagnosticPrecedence() []string {
	return []string{
		VerifierEcosystemDiagnosticInvalidSignature,
		VerifierEcosystemDiagnosticDigestMismatch,
		VerifierEcosystemDiagnosticUnsupportedProofType,
		VerifierEcosystemDiagnosticUnsupportedSchema,
		VerifierEcosystemDiagnosticSchemaMismatch,
		VerifierEcosystemDiagnosticRevokedIssuer,
		VerifierEcosystemDiagnosticExpiredArtifact,
		VerifierEcosystemDiagnosticStaleArtifact,
		VerifierEcosystemDiagnosticSupersededProof,
		VerifierEcosystemDiagnosticInsufficientTrustMaterial,
		VerifierEcosystemDiagnosticIncompleteArtifact,
		VerifierEcosystemDiagnosticScopeMismatch,
		VerifierEcosystemDiagnosticRedactionViolation,
		VerifierEcosystemDiagnosticCompatibilityWarning,
		VerifierEcosystemDiagnosticVerified,
		VerifierEcosystemDiagnosticUnknown,
	}
}

func verifierEcosystemValAEvidenceRefs() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:verifier-input-vala-001", EvidenceType: "schema_definition", Source: "verifier/input-model", Timestamp: "2026-04-27T08:00:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "verifier_input", Caveats: []string{"bounded to reference verifier request descriptors"}},
		{EvidenceID: "evidence:verifier-engine-vala-001", EvidenceType: "signature_material", Source: "verifier/engine", Timestamp: "2026-04-27T08:01:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "verifier_engine", Caveats: []string{"verification semantics remain bounded and explicit"}},
		{EvidenceID: "evidence:verifier-report-vala-001", EvidenceType: "compatibility_metadata", Source: "verifier/report", Timestamp: "2026-04-27T08:02:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "verification_report", Caveats: []string{"report remains advisory and scope-bounded"}},
		{EvidenceID: "evidence:verifier-sdk-vala-001", EvidenceType: "trust_root_metadata", Source: "verifier/sdk", Timestamp: "2026-04-27T08:03:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "sdk_entrypoint", Caveats: []string{"sdk entrypoint is deterministic and non-mutating"}},
	}
}

func VerifierEcosystemValAVerifierEvidence() []ReferenceArchitectureEvidenceReference {
	return verifierEcosystemValAEvidenceRefs()
}

func VerifierEcosystemValAReferenceVerifierInputModel() VerifierEcosystemValAReferenceVerifierInput {
	return VerifierEcosystemValAReferenceVerifierInput{
		CurrentState:                    "verifier_ecosystem_vala_input_ready",
		VerificationRequestID:           "reference-verifier-request-vala-001",
		VerifierContractRef:             "/v1/verifier-ecosystem/val0/contract",
		ProofEnvelopeRef:                "/v1/verifier-ecosystem/val0/proof-envelope",
		ArtifactRef:                     "artifact:reference-verifier-bundle",
		ArtifactDigest:                  "2d8a84e4e1ec70e3cf6d0e6d7f1b9bc8a1e33f6f7bb4ab4fc1c2b6d9e9ac0021",
		ArtifactDigestAlgorithm:         "sha256",
		SignatureRef:                    "signature:reference-verifier-bundle",
		IssuerRef:                       "issuer:reference-signer",
		TrustRootRef:                    "trust-root:reference-program",
		SchemaVersion:                   "changelock.verifier.proof_envelope.v1",
		ProofType:                       VerifierEcosystemProofTypeSignedAttestation,
		RequestedScope:                  VerifierEcosystemScopeAuditorSafe,
		VerificationTime:                "2026-04-27T08:04:00Z",
		ExpectedOutputBoundary:          VerifierEcosystemScopeAuditorSafe,
		CompatibilityPolicyRef:          "compatibility:reference-proof-envelope",
		RevocationMaterialRef:           "revocation:reference-proof-envelope",
		SupersessionMaterialRef:         "supersession:reference-proof-envelope",
		EvidenceRefs:                    []string{"evidence:verifier-input-vala-001", "evidence:verifier-engine-vala-001", "evidence:verifier-report-vala-001"},
		DigestVerificationState:         VerifierEcosystemValADigestMatch,
		SignatureVerificationState:      VerifierEcosystemValASignatureValid,
		SchemaVerificationState:         VerifierEcosystemValASchemaValid,
		ScopeVerificationState:          VerifierEcosystemValAScopeValid,
		FreshnessVerificationState:      IntelligenceCalibrationFreshnessFresh,
		TrustRootVerificationState:      VerifierEcosystemTrustRootTrusted,
		IssuerVerificationState:         VerifierEcosystemValAIssuerTrusted,
		CompatibilityEvaluationState:    ReferenceArchitectureCompatibilityCompatible,
		RevocationEvaluationState:       VerifierEcosystemValARevocationNotRevoked,
		SupersessionEvaluationState:     VerifierEcosystemValASupersessionCurrent,
		LineageVerificationState:        VerifierEcosystemValALineagePresent,
		OutputBoundaryVerificationState: VerifierEcosystemValAOutputBoundaryValid,
		StrictFailClosed:                true,
		Caveats:                         []string{"verification remains bounded to the declared proof envelope, schema, scope, trust-root material, freshness, and compatibility window"},
		ProjectionDisclaimer:            verifierEcosystemValAProjectionDisclaimer(),
	}
}

func VerifierEcosystemValAReferenceVerifierEngineModel(
	contractState, envelopeState, scopeState, compatibilityState, trustState, diagnosticsState, outputBoundaryState string,
) VerifierEcosystemValAReferenceVerifierEngine {
	return VerifierEcosystemValAReferenceVerifierEngine{
		CurrentState:              "verifier_ecosystem_vala_engine_ready",
		EngineID:                  "reference-verifier-engine-vala",
		Version:                   "2026.04",
		InputModelRef:             "/v1/verifier-ecosystem/vala/input-model",
		ResultModelRef:            "/v1/verifier-ecosystem/vala/verification-report",
		VerifierContractState:     contractState,
		ProofEnvelopeState:        envelopeState,
		RequestedScopeState:       scopeState,
		SchemaCompatibilityState:  compatibilityState,
		TrustRootIssuerState:      trustState,
		DiagnosticsState:          diagnosticsState,
		OutputBoundaryState:       outputBoundaryState,
		SupportedDigestAlgorithms: verifierEcosystemValASupportedDigestAlgorithms(),
		SupportedOverallResults:   verifierEcosystemValAOverallResults(),
		DeterministicOutput:       true,
		ExplicitFixtureSemantics:  true,
		UsesRealCryptoPrimitives:  false,
		ProjectionDisclaimer:      verifierEcosystemValAProjectionDisclaimer(),
	}
}

func VerifierEcosystemValADiagnosticsMappingModel(result VerifierEcosystemValAVerificationResult) VerifierEcosystemValADiagnosticsMapping {
	return VerifierEcosystemValADiagnosticsMapping{
		CurrentState:               "verifier_ecosystem_vala_diagnostics_mapping_ready",
		MappingID:                  "reference-verifier-diagnostics-mapping-vala",
		ProofType:                  result.ProofType,
		SupportedDiagnosticClasses: verifierEcosystemVal0DiagnosticClasses(),
		DeterministicPrecedence:    verifierEcosystemValADiagnosticPrecedence(),
		DigestResult:               result.DigestResult,
		SignatureResult:            result.SignatureResult,
		SchemaResult:               result.SchemaResult,
		ScopeResult:                result.ScopeResult,
		FreshnessResult:            result.FreshnessResult,
		TrustRootResult:            result.TrustRootResult,
		IssuerResult:               result.IssuerResult,
		CompatibilityResult:        result.CompatibilityResult,
		RevocationResult:           result.RevocationResult,
		SupersessionResult:         result.SupersessionResult,
		LineageResult:              result.LineageResult,
		OutputBoundaryResult:       result.OutputBoundaryResult,
		DerivedDiagnosticClass:     result.DiagnosticClass,
		ProjectionDisclaimer:       verifierEcosystemValAProjectionDisclaimer(),
	}
}

func VerifierEcosystemValACommandContractModel() VerifierEcosystemValACommandContract {
	return VerifierEcosystemValACommandContract{
		CurrentState:                "verifier_ecosystem_vala_command_contract_ready",
		CommandContractID:           "reference-verifier-command-contract-vala",
		ProofEnvelopeInputRef:       "<proof-envelope-ref-or-path>",
		ArtifactDigestDescriptor:    "<sha256:artifact-digest>",
		SignatureDescriptor:         "<signature-ref-or-path>",
		TrustRootMaterialDescriptor: "<trust-root-bundle-ref-or-path>",
		RequestedScope:              VerifierEcosystemScopeAuditorSafe,
		OutputBoundary:              VerifierEcosystemScopeAuditorSafe,
		ReportFormat:                VerifierEcosystemValACommandReportFormatJSON,
		StrictFailClosedMode:        true,
		ExampleCommand:              "changelock-cli verify-proof --input <proof-envelope-ref> --scope auditor_safe --output-boundary auditor_safe --format json --strict",
		ProjectionDisclaimer:        verifierEcosystemValAProjectionDisclaimer(),
	}
}

func VerifierEcosystemValASDKEntrypointModel() VerifierEcosystemValASDKEntrypoint {
	return VerifierEcosystemValASDKEntrypoint{
		CurrentState:                 "verifier_ecosystem_vala_sdk_entrypoint_ready",
		EntrypointID:                 "reference-verifier-sdk-entrypoint-vala",
		PackagePath:                  "github.com/denisgrosek/changelock/internal/verifier",
		FunctionName:                 "VerifyReferenceVerifierRequest",
		InputModelRef:                "/v1/verifier-ecosystem/vala/input-model",
		ResultModelRef:               "/v1/verifier-ecosystem/vala/verification-report",
		DeterministicOutput:          true,
		HiddenMainInstanceDependency: false,
		NetworkDependency:            false,
		MutatesCanonicalState:        false,
		GeneratesFakeEvidence:        false,
		ProjectionDisclaimer:         verifierEcosystemValAProjectionDisclaimer(),
	}
}

func verifierEcosystemValAAllRequiredInputRefs(input VerifierEcosystemValAReferenceVerifierInput) bool {
	return referenceArchitectureValBRequiredRefsPresent(
		input.VerificationRequestID,
		input.VerifierContractRef,
		input.ProofEnvelopeRef,
		input.ArtifactRef,
		input.ArtifactDigest,
		input.ArtifactDigestAlgorithm,
		input.SignatureRef,
		input.IssuerRef,
		input.TrustRootRef,
		input.SchemaVersion,
		input.ProofType,
		input.RequestedScope,
		input.VerificationTime,
		input.ExpectedOutputBoundary,
		input.CompatibilityPolicyRef,
		input.RevocationMaterialRef,
		input.SupersessionMaterialRef,
		input.ProjectionDisclaimer,
	)
}

func EvaluateVerifierEcosystemValAReferenceVerifierInputState(input VerifierEcosystemValAReferenceVerifierInput) string {
	if !verifierEcosystemValAAllRequiredInputRefs(input) {
		return VerifierEcosystemValAInputStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemValASupportedDigestAlgorithms(), input.ArtifactDigestAlgorithm) ||
		!containsTrimmedString(verifierEcosystemVal0SupportedProofTypes(), input.ProofType) ||
		!containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), input.RequestedScope) ||
		!containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), input.ExpectedOutputBoundary) {
		return VerifierEcosystemValAInputStateUnknown
	}
	if !containsTrimmedString(verifierEcosystemValADigestResults(), input.DigestVerificationState) ||
		!containsTrimmedString(verifierEcosystemValASignatureResults(), input.SignatureVerificationState) ||
		!containsTrimmedString(verifierEcosystemValASchemaResults(), input.SchemaVerificationState) ||
		!containsTrimmedString(verifierEcosystemValAScopeResults(), input.ScopeVerificationState) ||
		!containsTrimmedString([]string{IntelligenceCalibrationFreshnessFresh, IntelligenceCalibrationFreshnessStale, IntelligenceCalibrationFreshnessExpired, IntelligenceCalibrationFreshnessUnsupported, IntelligenceCalibrationFreshnessUnknown}, input.FreshnessVerificationState) ||
		!containsTrimmedString(verifierEcosystemVal0TrustRootStates(), input.TrustRootVerificationState) ||
		!containsTrimmedString(verifierEcosystemValAIssuerResults(), input.IssuerVerificationState) ||
		!containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), input.CompatibilityEvaluationState) ||
		!containsTrimmedString(verifierEcosystemValARevocationResults(), input.RevocationEvaluationState) ||
		!containsTrimmedString(verifierEcosystemValASupersessionResults(), input.SupersessionEvaluationState) ||
		!containsTrimmedString(verifierEcosystemValALineageResults(), input.LineageVerificationState) ||
		!containsTrimmedString(verifierEcosystemValAOutputBoundaryResults(), input.OutputBoundaryVerificationState) {
		return VerifierEcosystemValAInputStateUnknown
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(input.VerificationTime); !ok {
		return VerifierEcosystemValAInputStatePartial
	}
	if !verifierEcosystemVal0HasProjectionDisclaimer(input.ProjectionDisclaimer) {
		return VerifierEcosystemValAInputStatePartial
	}
	if input.CanonicalEvidenceSpineAccess || input.TruthOutsideScopeClaim ||
		verifierEcosystemVal0HasOverclaim(strings.Join(input.Caveats, " "), input.ProjectionDisclaimer) {
		return VerifierEcosystemValAInputStateBlocked
	}
	if strings.TrimSpace(input.FreshnessVerificationState) == IntelligenceCalibrationFreshnessUnsupported ||
		strings.TrimSpace(input.FreshnessVerificationState) == IntelligenceCalibrationFreshnessUnknown ||
		strings.TrimSpace(input.CompatibilityEvaluationState) == ReferenceArchitectureCompatibilityUnsupported ||
		strings.TrimSpace(input.CompatibilityEvaluationState) == ReferenceArchitectureCompatibilityUnknown {
		return VerifierEcosystemValAInputStateBlocked
	}
	return VerifierEcosystemValAInputStateActive
}

func verifierEcosystemValAStateSeverity(state, active, partial, incomplete, blocked, unknown string) int {
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

func EvaluateVerifierEcosystemValAReferenceVerifierEngineState(engine VerifierEcosystemValAReferenceVerifierEngine) string {
	if !engine.DeterministicOutput || !engine.ExplicitFixtureSemantics ||
		engine.NetworkDependency || engine.MutatesEvidence || engine.ClaimsDeploymentApproval || engine.ClaimsCanonicalTruth || engine.ClaimsCertification ||
		(engine.ClaimsActualCryptoValidity && !engine.UsesRealCryptoPrimitives) ||
		verifierEcosystemVal0HasOverclaim(engine.ProjectionDisclaimer) {
		return VerifierEcosystemValAEngineStateBlocked
	}
	if !referenceArchitectureValBRequiredRefsPresent(
		engine.EngineID,
		engine.Version,
		engine.InputModelRef,
		engine.ResultModelRef,
		engine.VerifierContractState,
		engine.ProofEnvelopeState,
		engine.RequestedScopeState,
		engine.SchemaCompatibilityState,
		engine.TrustRootIssuerState,
		engine.DiagnosticsState,
		engine.OutputBoundaryState,
		engine.ProjectionDisclaimer,
	) {
		return VerifierEcosystemValAEngineStateIncomplete
	}
	if !containsExactTrimmedStringSet(engine.SupportedDigestAlgorithms, verifierEcosystemValASupportedDigestAlgorithms()...) ||
		!containsExactTrimmedStringSet(engine.SupportedOverallResults, verifierEcosystemValAOverallResults()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(engine.ProjectionDisclaimer) {
		return VerifierEcosystemValAEngineStatePartial
	}
	highestSeverity := 0
	componentSeverities := []int{
		verifierEcosystemValAStateSeverity(
			engine.VerifierContractState,
			VerifierEcosystemVal0ContractStateActive,
			VerifierEcosystemVal0ContractStatePartial,
			VerifierEcosystemVal0ContractStateIncomplete,
			VerifierEcosystemVal0ContractStateBlocked,
			VerifierEcosystemVal0ContractStateUnknown,
		),
		verifierEcosystemValAStateSeverity(
			engine.ProofEnvelopeState,
			VerifierEcosystemVal0EnvelopeStateActive,
			VerifierEcosystemVal0EnvelopeStatePartial,
			VerifierEcosystemVal0EnvelopeStateIncomplete,
			VerifierEcosystemVal0EnvelopeStateBlocked,
			VerifierEcosystemVal0EnvelopeStateUnknown,
		),
		verifierEcosystemValAStateSeverity(
			engine.RequestedScopeState,
			VerifierEcosystemVal0ScopeStateActive,
			VerifierEcosystemVal0ScopeStatePartial,
			VerifierEcosystemVal0ScopeStateIncomplete,
			VerifierEcosystemVal0ScopeStateBlocked,
			VerifierEcosystemVal0ScopeStateUnknown,
		),
		verifierEcosystemValAStateSeverity(
			engine.SchemaCompatibilityState,
			VerifierEcosystemVal0CompatibilityStateActive,
			VerifierEcosystemVal0CompatibilityStatePartial,
			VerifierEcosystemVal0CompatibilityStateIncomplete,
			VerifierEcosystemVal0CompatibilityStateBlocked,
			VerifierEcosystemVal0CompatibilityStateUnknown,
		),
		verifierEcosystemValAStateSeverity(
			engine.TrustRootIssuerState,
			VerifierEcosystemVal0TrustStateActive,
			VerifierEcosystemVal0TrustStatePartial,
			VerifierEcosystemVal0TrustStateIncomplete,
			VerifierEcosystemVal0TrustStateBlocked,
			VerifierEcosystemVal0TrustStateUnknown,
		),
		verifierEcosystemValAStateSeverity(
			engine.DiagnosticsState,
			VerifierEcosystemVal0DiagnosticsStateActive,
			VerifierEcosystemVal0DiagnosticsStatePartial,
			VerifierEcosystemVal0DiagnosticsStateIncomplete,
			VerifierEcosystemVal0DiagnosticsStateBlocked,
			VerifierEcosystemVal0DiagnosticsStateUnknown,
		),
		verifierEcosystemValAStateSeverity(
			engine.OutputBoundaryState,
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
		return VerifierEcosystemValAEngineStateBlocked
	case 3:
		return VerifierEcosystemValAEngineStateUnknown
	case 2:
		return VerifierEcosystemValAEngineStateIncomplete
	case 1:
		return VerifierEcosystemValAEngineStatePartial
	}
	return VerifierEcosystemValAEngineStateActive
}

func deriveVerifierEcosystemValADiagnosticClassFromComponents(
	proofType, digestResult, signatureResult, schemaResult, scopeResult, freshnessResult, trustRootResult, issuerResult, compatibilityResult, revocationResult, supersessionResult, lineageResult, outputBoundaryResult string,
) string {
	if !containsTrimmedString(verifierEcosystemVal0SupportedProofTypes(), proofType) {
		return VerifierEcosystemDiagnosticUnsupportedProofType
	}
	if signatureResult == VerifierEcosystemValASignatureInvalid {
		return VerifierEcosystemDiagnosticInvalidSignature
	}
	if digestResult == VerifierEcosystemValADigestMismatch {
		return VerifierEcosystemDiagnosticDigestMismatch
	}
	if schemaResult == VerifierEcosystemValASchemaUnsupported {
		return VerifierEcosystemDiagnosticUnsupportedSchema
	}
	if schemaResult == VerifierEcosystemValASchemaMismatch {
		return VerifierEcosystemDiagnosticSchemaMismatch
	}
	if strings.TrimSpace(trustRootResult) == VerifierEcosystemTrustRootRevoked ||
		strings.TrimSpace(trustRootResult) == VerifierEcosystemTrustRootExpired ||
		strings.TrimSpace(issuerResult) == VerifierEcosystemValAIssuerRevoked ||
		strings.TrimSpace(revocationResult) == VerifierEcosystemValARevocationRevoked {
		return VerifierEcosystemDiagnosticRevokedIssuer
	}
	if strings.TrimSpace(freshnessResult) == IntelligenceCalibrationFreshnessExpired {
		return VerifierEcosystemDiagnosticExpiredArtifact
	}
	if strings.TrimSpace(freshnessResult) == IntelligenceCalibrationFreshnessStale {
		return VerifierEcosystemDiagnosticStaleArtifact
	}
	if strings.TrimSpace(supersessionResult) == VerifierEcosystemValASupersessionSuperseded {
		return VerifierEcosystemDiagnosticSupersededProof
	}
	if strings.TrimSpace(outputBoundaryResult) == VerifierEcosystemValAOutputBoundaryViolation {
		return VerifierEcosystemDiagnosticRedactionViolation
	}
	if strings.TrimSpace(signatureResult) == VerifierEcosystemValASignatureUnsupported ||
		strings.TrimSpace(digestResult) == VerifierEcosystemValADigestUnsupported ||
		strings.TrimSpace(trustRootResult) == VerifierEcosystemTrustRootUnsupported ||
		strings.TrimSpace(trustRootResult) == VerifierEcosystemTrustRootUnknown ||
		strings.TrimSpace(revocationResult) == VerifierEcosystemValARevocationMaterialMissing {
		return VerifierEcosystemDiagnosticInsufficientTrustMaterial
	}
	if strings.TrimSpace(signatureResult) == VerifierEcosystemValASignatureMissing ||
		strings.TrimSpace(digestResult) == VerifierEcosystemValADigestMissing ||
		strings.TrimSpace(lineageResult) == VerifierEcosystemValALineageMissing {
		return VerifierEcosystemDiagnosticIncompleteArtifact
	}
	if strings.TrimSpace(scopeResult) == VerifierEcosystemValAScopeMismatch ||
		strings.TrimSpace(scopeResult) == VerifierEcosystemValAScopeUnsupported {
		return VerifierEcosystemDiagnosticScopeMismatch
	}
	if strings.TrimSpace(compatibilityResult) == ReferenceArchitectureCompatibilityCompatibleWithWarning ||
		strings.TrimSpace(compatibilityResult) == ReferenceArchitectureCompatibilityDeprecated ||
		strings.TrimSpace(compatibilityResult) == ReferenceArchitectureCompatibilitySuperseded ||
		strings.TrimSpace(trustRootResult) == VerifierEcosystemTrustRootTrustedWithWarnings {
		return VerifierEcosystemDiagnosticCompatibilityWarning
	}
	validStates := containsTrimmedString(verifierEcosystemValADigestResults(), digestResult) &&
		containsTrimmedString(verifierEcosystemValASignatureResults(), signatureResult) &&
		containsTrimmedString(verifierEcosystemValASchemaResults(), schemaResult) &&
		containsTrimmedString(verifierEcosystemValAScopeResults(), scopeResult) &&
		containsTrimmedString([]string{IntelligenceCalibrationFreshnessFresh, IntelligenceCalibrationFreshnessStale, IntelligenceCalibrationFreshnessExpired}, freshnessResult) &&
		containsTrimmedString(verifierEcosystemVal0TrustRootStates(), trustRootResult) &&
		containsTrimmedString(verifierEcosystemValAIssuerResults(), issuerResult) &&
		containsTrimmedString(verifierEcosystemVal0CompatibilityResults(), compatibilityResult) &&
		containsTrimmedString(verifierEcosystemValARevocationResults(), revocationResult) &&
		containsTrimmedString(verifierEcosystemValASupersessionResults(), supersessionResult) &&
		containsTrimmedString(verifierEcosystemValALineageResults(), lineageResult) &&
		containsTrimmedString(verifierEcosystemValAOutputBoundaryResults(), outputBoundaryResult)
	if !validStates ||
		strings.TrimSpace(freshnessResult) == IntelligenceCalibrationFreshnessUnknown ||
		strings.TrimSpace(freshnessResult) == IntelligenceCalibrationFreshnessUnsupported ||
		strings.TrimSpace(compatibilityResult) == ReferenceArchitectureCompatibilityUnknown ||
		strings.TrimSpace(outputBoundaryResult) == VerifierEcosystemValAOutputBoundaryUnknown ||
		strings.TrimSpace(issuerResult) == VerifierEcosystemValAIssuerUnknown ||
		strings.TrimSpace(revocationResult) == VerifierEcosystemValARevocationUnknown ||
		strings.TrimSpace(supersessionResult) == VerifierEcosystemValASupersessionUnknown ||
		strings.TrimSpace(lineageResult) == VerifierEcosystemValALineageUnknown {
		return VerifierEcosystemDiagnosticUnknown
	}
	return VerifierEcosystemDiagnosticVerified
}

func DeriveVerifierEcosystemValADiagnosticClass(result VerifierEcosystemValAVerificationResult) string {
	return deriveVerifierEcosystemValADiagnosticClassFromComponents(
		result.ProofType,
		result.DigestResult,
		result.SignatureResult,
		result.SchemaResult,
		result.ScopeResult,
		result.FreshnessResult,
		result.TrustRootResult,
		result.IssuerResult,
		result.CompatibilityResult,
		result.RevocationResult,
		result.SupersessionResult,
		result.LineageResult,
		result.OutputBoundaryResult,
	)
}

func deriveVerifierEcosystemValAOverallResultFromDiagnostic(result VerifierEcosystemValAVerificationResult, diagnostic string) string {
	switch diagnostic {
	case VerifierEcosystemDiagnosticVerified:
		return VerifierEcosystemValAOverallResultVerified
	case VerifierEcosystemDiagnosticCompatibilityWarning:
		if len(result.Caveats) == 0 {
			return VerifierEcosystemValAOverallResultIncomplete
		}
		return VerifierEcosystemValAOverallResultWarnings
	case VerifierEcosystemDiagnosticInvalidSignature,
		VerifierEcosystemDiagnosticDigestMismatch,
		VerifierEcosystemDiagnosticSchemaMismatch,
		VerifierEcosystemDiagnosticScopeMismatch,
		VerifierEcosystemDiagnosticRedactionViolation:
		return VerifierEcosystemValAOverallResultInvalid
	case VerifierEcosystemDiagnosticUnsupportedSchema,
		VerifierEcosystemDiagnosticUnsupportedProofType:
		return VerifierEcosystemValAOverallResultUnsupported
	case VerifierEcosystemDiagnosticStaleArtifact,
		VerifierEcosystemDiagnosticExpiredArtifact:
		return VerifierEcosystemValAOverallResultStale
	case VerifierEcosystemDiagnosticRevokedIssuer:
		return VerifierEcosystemValAOverallResultRevoked
	case VerifierEcosystemDiagnosticSupersededProof:
		return VerifierEcosystemValAOverallResultSuperseded
	case VerifierEcosystemDiagnosticInsufficientTrustMaterial,
		VerifierEcosystemDiagnosticIncompleteArtifact:
		return VerifierEcosystemValAOverallResultIncomplete
	default:
		return VerifierEcosystemValAOverallResultUnknown
	}
}

func EvaluateVerifierEcosystemValAVerificationResultState(result VerifierEcosystemValAVerificationResult) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		result.VerificationResultID,
		result.RequestID,
		result.VerifierVersion,
		result.ProofType,
		result.SchemaVersion,
		result.Scope,
		result.OutputBoundary,
		result.OverallResult,
		result.DiagnosticClass,
		result.DigestResult,
		result.SignatureResult,
		result.SchemaResult,
		result.ScopeResult,
		result.FreshnessResult,
		result.TrustRootResult,
		result.IssuerResult,
		result.CompatibilityResult,
		result.RevocationResult,
		result.SupersessionResult,
		result.LineageResult,
		result.OutputBoundaryResult,
		result.ProjectionDisclaimer,
		result.VerifiedAt,
	) {
		return VerifierEcosystemValAResultStateIncomplete
	}
	if len(result.EvidenceRefs) == 0 || len(result.Limitations) == 0 {
		return VerifierEcosystemValAResultStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemValAOverallResults(), result.OverallResult) ||
		!containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), result.DiagnosticClass) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(result.ProjectionDisclaimer) {
		return VerifierEcosystemValAResultStateUnknown
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(result.VerifiedAt); !ok {
		return VerifierEcosystemValAResultStatePartial
	}
	if result.TruthOutsideScopeClaim || verifierEcosystemVal0HasOverclaim(strings.Join(result.Caveats, " "), strings.Join(result.Limitations, " "), result.ProjectionDisclaimer) {
		return VerifierEcosystemValAResultStateBlocked
	}
	expectedDiagnostic := DeriveVerifierEcosystemValADiagnosticClass(result)
	if strings.TrimSpace(result.DiagnosticClass) != expectedDiagnostic {
		return VerifierEcosystemValAResultStatePartial
	}
	expectedOverall := deriveVerifierEcosystemValAOverallResultFromDiagnostic(result, expectedDiagnostic)
	if strings.TrimSpace(result.OverallResult) != expectedOverall {
		return VerifierEcosystemValAResultStatePartial
	}
	switch strings.TrimSpace(result.OverallResult) {
	case VerifierEcosystemValAOverallResultVerified:
		return VerifierEcosystemValAResultStateActive
	case VerifierEcosystemValAOverallResultWarnings:
		if len(result.Caveats) == 0 {
			return VerifierEcosystemValAResultStateBlocked
		}
		return VerifierEcosystemValAResultStatePartial
	case VerifierEcosystemValAOverallResultIncomplete:
		return VerifierEcosystemValAResultStateIncomplete
	case VerifierEcosystemValAOverallResultUnknown:
		return VerifierEcosystemValAResultStateUnknown
	default:
		return VerifierEcosystemValAResultStateBlocked
	}
}

func EvaluateVerifierEcosystemValADiagnosticsMappingState(model VerifierEcosystemValADiagnosticsMapping) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.MappingID,
		model.ProofType,
		model.DerivedDiagnosticClass,
		model.ProjectionDisclaimer,
	) {
		return VerifierEcosystemValADiagnosticsMappingStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedDiagnosticClasses, verifierEcosystemVal0DiagnosticClasses()...) ||
		!containsExactTrimmedStringSet(model.DeterministicPrecedence, verifierEcosystemValADiagnosticPrecedence()...) ||
		!verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValADiagnosticsMappingStatePartial
	}
	expected := deriveVerifierEcosystemValADiagnosticClassFromComponents(
		model.ProofType,
		model.DigestResult,
		model.SignatureResult,
		model.SchemaResult,
		model.ScopeResult,
		model.FreshnessResult,
		model.TrustRootResult,
		model.IssuerResult,
		model.CompatibilityResult,
		model.RevocationResult,
		model.SupersessionResult,
		model.LineageResult,
		model.OutputBoundaryResult,
	)
	if !containsTrimmedString(verifierEcosystemVal0DiagnosticClasses(), model.DerivedDiagnosticClass) {
		return VerifierEcosystemValADiagnosticsMappingStateUnknown
	}
	if strings.TrimSpace(model.DerivedDiagnosticClass) != expected {
		return VerifierEcosystemValADiagnosticsMappingStatePartial
	}
	if strings.TrimSpace(model.DerivedDiagnosticClass) == VerifierEcosystemDiagnosticUnknown {
		return VerifierEcosystemValADiagnosticsMappingStateUnknown
	}
	return VerifierEcosystemValADiagnosticsMappingStateActive
}

func EvaluateVerifierEcosystemValACommandContractState(model VerifierEcosystemValACommandContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CommandContractID,
		model.ProofEnvelopeInputRef,
		model.ArtifactDigestDescriptor,
		model.SignatureDescriptor,
		model.TrustRootMaterialDescriptor,
		model.RequestedScope,
		model.OutputBoundary,
		model.ReportFormat,
		model.ExampleCommand,
		model.ProjectionDisclaimer,
	) {
		return VerifierEcosystemValACommandContractStateIncomplete
	}
	if !containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), model.RequestedScope) ||
		!containsTrimmedString(verifierEcosystemVal0SupportedScopeClasses(), model.OutputBoundary) ||
		!containsTrimmedString(verifierEcosystemValASupportedReportFormats(), model.ReportFormat) {
		return VerifierEcosystemValACommandContractStateUnknown
	}
	if !verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValACommandContractStatePartial
	}
	if model.MutatesEvidence || model.ApprovesDeployment || model.SuppressesFailures || model.PublishesClaims || model.CertificationClaim ||
		verifierEcosystemVal0HasOverclaim(model.ExampleCommand, model.ProjectionDisclaimer) {
		return VerifierEcosystemValACommandContractStateBlocked
	}
	if strings.TrimSpace(model.ReportFormat) == VerifierEcosystemValACommandReportFormatUnknown {
		return VerifierEcosystemValACommandContractStateUnknown
	}
	if !model.StrictFailClosedMode {
		return VerifierEcosystemValACommandContractStatePartial
	}
	return VerifierEcosystemValACommandContractStateActive
}

func EvaluateVerifierEcosystemValASDKEntrypointState(model VerifierEcosystemValASDKEntrypoint) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.EntrypointID,
		model.PackagePath,
		model.FunctionName,
		model.InputModelRef,
		model.ResultModelRef,
		model.ProjectionDisclaimer,
	) {
		return VerifierEcosystemValASDKEntrypointStateIncomplete
	}
	if !verifierEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return VerifierEcosystemValASDKEntrypointStatePartial
	}
	if !model.DeterministicOutput {
		return VerifierEcosystemValASDKEntrypointStatePartial
	}
	if model.HiddenMainInstanceDependency || model.NetworkDependency || model.MutatesCanonicalState || model.GeneratesFakeEvidence ||
		verifierEcosystemVal0HasOverclaim(model.PackagePath, model.FunctionName, model.ProjectionDisclaimer) {
		return VerifierEcosystemValASDKEntrypointStateBlocked
	}
	return VerifierEcosystemValASDKEntrypointStateActive
}

func verifierEcosystemValAPoint6DependencyHealthy(snapshot VerifierEcosystemValADependencySnapshot) bool {
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

func EvaluateVerifierEcosystemValAState(
	dependency VerifierEcosystemValADependencySnapshot,
	inputState, engineState, resultState, diagnosticsMappingState, commandState, sdkState string,
) string {
	if !verifierEcosystemValAPoint6DependencyHealthy(dependency) ||
		strings.TrimSpace(dependency.Val0CurrentState) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(dependency.Val0State) != VerifierEcosystemVal0StateActive ||
		strings.TrimSpace(dependency.Point7State) != VerifierEcosystemPoint7StateNotComplete {
		return VerifierEcosystemValAStateBlocked
	}
	highestSeverity := 0
	componentSeverities := []int{
		verifierEcosystemValAStateSeverity(inputState, VerifierEcosystemValAInputStateActive, VerifierEcosystemValAInputStatePartial, VerifierEcosystemValAInputStateIncomplete, VerifierEcosystemValAInputStateBlocked, VerifierEcosystemValAInputStateUnknown),
		verifierEcosystemValAStateSeverity(engineState, VerifierEcosystemValAEngineStateActive, VerifierEcosystemValAEngineStatePartial, VerifierEcosystemValAEngineStateIncomplete, VerifierEcosystemValAEngineStateBlocked, VerifierEcosystemValAEngineStateUnknown),
		verifierEcosystemValAStateSeverity(resultState, VerifierEcosystemValAResultStateActive, VerifierEcosystemValAResultStatePartial, VerifierEcosystemValAResultStateIncomplete, VerifierEcosystemValAResultStateBlocked, VerifierEcosystemValAResultStateUnknown),
		verifierEcosystemValAStateSeverity(diagnosticsMappingState, VerifierEcosystemValADiagnosticsMappingStateActive, VerifierEcosystemValADiagnosticsMappingStatePartial, VerifierEcosystemValADiagnosticsMappingStateIncomplete, VerifierEcosystemValADiagnosticsMappingStateBlocked, VerifierEcosystemValADiagnosticsMappingStateUnknown),
		verifierEcosystemValAStateSeverity(commandState, VerifierEcosystemValACommandContractStateActive, VerifierEcosystemValACommandContractStatePartial, VerifierEcosystemValACommandContractStateIncomplete, VerifierEcosystemValACommandContractStateBlocked, VerifierEcosystemValACommandContractStateUnknown),
		verifierEcosystemValAStateSeverity(sdkState, VerifierEcosystemValASDKEntrypointStateActive, VerifierEcosystemValASDKEntrypointStatePartial, VerifierEcosystemValASDKEntrypointStateIncomplete, VerifierEcosystemValASDKEntrypointStateBlocked, VerifierEcosystemValASDKEntrypointStateUnknown),
	}
	for _, severity := range componentSeverities {
		if severity > highestSeverity {
			highestSeverity = severity
		}
	}
	switch highestSeverity {
	case 4:
		return VerifierEcosystemValAStateBlocked
	case 3:
		return VerifierEcosystemValAStateUnknown
	case 2:
		return VerifierEcosystemValAStateIncomplete
	case 1:
		return VerifierEcosystemValAStatePartial
	}
	return VerifierEcosystemValAStateActive
}

func VerifierEcosystemValAProofSurfaceRefs() []string {
	return []string{
		"/v1/verifier-ecosystem/val0/proofs",
		"/v1/verifier-ecosystem/vala/input-model",
		"/v1/verifier-ecosystem/vala/verifier-engine",
		"/v1/verifier-ecosystem/vala/verification-report",
		"/v1/verifier-ecosystem/vala/diagnostics-mapping",
		"/v1/verifier-ecosystem/vala/command-contract",
		"/v1/verifier-ecosystem/vala/sdk-entrypoint",
		"/v1/verifier-ecosystem/vala/proofs",
	}
}

func EvaluateVerifierEcosystemValAProofsState(
	currentState string,
	point7State string,
	val0CurrentState string,
	surfaceRefs, evidenceRefs, limitations, whyPoint7NotPass []string,
	projectionDisclaimer string,
) string {
	baseState := strings.TrimSpace(currentState)
	if strings.TrimSpace(val0CurrentState) != VerifierEcosystemVal0StateActive ||
		!containsExactTrimmedStringSet(surfaceRefs, VerifierEcosystemValAProofSurfaceRefs()...) ||
		len(evidenceRefs) < 4 ||
		len(limitations) == 0 ||
		len(whyPoint7NotPass) == 0 ||
		!verifierEcosystemVal0HasProjectionDisclaimer(projectionDisclaimer) {
		if baseState == VerifierEcosystemValAStateActive {
			return VerifierEcosystemValAStatePartial
		}
		return baseState
	}
	if baseState == VerifierEcosystemValAStateActive && strings.TrimSpace(point7State) != VerifierEcosystemPoint7StateNotComplete {
		return VerifierEcosystemValAStatePartial
	}
	return baseState
}

func EvaluateVerifierEcosystemValAPoint7State(valAState string) string {
	_ = valAState
	return VerifierEcosystemPoint7StateNotComplete
}

package verifier

import (
	"strings"
	"time"
)

const (
	referenceVerifierReportFormatJSON        = "json"
	referenceVerifierReportFormatTextSummary = "text_summary"
	referenceVerifierReportFormatUnknown     = "unknown"

	referenceVerifierDigestMatch       = "digest_match"
	referenceVerifierDigestMismatch    = "digest_mismatch"
	referenceVerifierDigestMissing     = "digest_missing"
	referenceVerifierDigestUnsupported = "digest_unsupported"

	referenceVerifierSignatureValid       = "signature_valid"
	referenceVerifierSignatureInvalid     = "signature_invalid"
	referenceVerifierSignatureMissing     = "signature_missing"
	referenceVerifierSignatureUnsupported = "signature_unsupported"

	referenceVerifierSchemaValid       = "schema_valid"
	referenceVerifierSchemaMismatch    = "schema_mismatch"
	referenceVerifierSchemaUnsupported = "schema_unsupported"

	referenceVerifierScopeValid       = "scope_valid"
	referenceVerifierScopeMismatch    = "scope_mismatch"
	referenceVerifierScopeUnsupported = "scope_unsupported"

	referenceVerifierIssuerTrusted = "issuer_trusted"
	referenceVerifierIssuerRevoked = "issuer_revoked"
	referenceVerifierIssuerUnknown = "issuer_unknown"

	referenceVerifierRevocationNotRevoked      = "not_revoked"
	referenceVerifierRevocationRevoked         = "revoked"
	referenceVerifierRevocationMaterialMissing = "revocation_material_missing"
	referenceVerifierRevocationUnknown         = "unknown"

	referenceVerifierSupersessionCurrent    = "current"
	referenceVerifierSupersessionSuperseded = "superseded"
	referenceVerifierSupersessionUnknown    = "unknown"

	referenceVerifierLineagePresent = "lineage_present"
	referenceVerifierLineageMissing = "lineage_missing"
	referenceVerifierLineageUnknown = "lineage_unknown"

	referenceVerifierOutputBoundaryValid     = "output_boundary_valid"
	referenceVerifierOutputBoundaryViolation = "output_boundary_violation"
	referenceVerifierOutputBoundaryUnknown   = "output_boundary_unknown"

	referenceVerifierOverallVerified    = "verified"
	referenceVerifierOverallWarnings    = "verified_with_warnings"
	referenceVerifierOverallInvalid     = "invalid"
	referenceVerifierOverallIncomplete  = "incomplete"
	referenceVerifierOverallUnsupported = "unsupported"
	referenceVerifierOverallStale       = "stale"
	referenceVerifierOverallRevoked     = "revoked"
	referenceVerifierOverallSuperseded  = "superseded"
	referenceVerifierOverallUnknown     = "unknown"

	referenceVerifierDiagnosticVerified             = "verified"
	referenceVerifierDiagnosticInvalidSignature     = "invalid_signature"
	referenceVerifierDiagnosticDigestMismatch       = "digest_mismatch"
	referenceVerifierDiagnosticSchemaMismatch       = "schema_mismatch"
	referenceVerifierDiagnosticUnsupportedSchema    = "unsupported_schema"
	referenceVerifierDiagnosticUnsupportedProofType = "unsupported_proof_type"
	referenceVerifierDiagnosticStaleArtifact        = "stale_artifact"
	referenceVerifierDiagnosticExpiredArtifact      = "expired_artifact"
	referenceVerifierDiagnosticRevokedIssuer        = "revoked_issuer"
	referenceVerifierDiagnosticSupersededProof      = "superseded_proof"
	referenceVerifierDiagnosticInsufficientTrust    = "insufficient_trust_material"
	referenceVerifierDiagnosticIncompleteArtifact   = "incomplete_artifact"
	referenceVerifierDiagnosticScopeMismatch        = "scope_mismatch"
	referenceVerifierDiagnosticRedactionViolation   = "redaction_boundary_violation"
	referenceVerifierDiagnosticCompatibilityWarning = "compatibility_warning"
	referenceVerifierDiagnosticUnknown              = "unknown"
)

type ReferenceVerifierRequest struct {
	RequestID                       string   `json:"request_id"`
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
	ReportFormat                    string   `json:"report_format"`
	StrictFailClosed                bool     `json:"strict_fail_closed"`
	TruthOutsideScopeClaim          bool     `json:"truth_outside_scope_claim"`
	ClaimsActualCryptoVerification  bool     `json:"claims_actual_crypto_verification"`
	Caveats                         []string `json:"caveats,omitempty"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
}

type ReferenceVerificationResult struct {
	ResultID               string   `json:"result_id"`
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

func VerifyReferenceVerifierRequest(request ReferenceVerifierRequest) ReferenceVerificationResult {
	request.RequestID = strings.TrimSpace(request.RequestID)
	request.ProofType = strings.TrimSpace(request.ProofType)
	request.SchemaVersion = strings.TrimSpace(request.SchemaVersion)
	request.RequestedScope = strings.TrimSpace(request.RequestedScope)
	request.ExpectedOutputBoundary = strings.TrimSpace(request.ExpectedOutputBoundary)
	request.VerificationTime = strings.TrimSpace(request.VerificationTime)
	request.ReportFormat = strings.TrimSpace(request.ReportFormat)

	digestResult := strings.TrimSpace(request.DigestVerificationState)
	if strings.TrimSpace(request.ArtifactDigest) == "" {
		digestResult = referenceVerifierDigestMissing
	} else if !containsString([]string{"sha256", "sha512"}, request.ArtifactDigestAlgorithm) {
		digestResult = referenceVerifierDigestUnsupported
	}

	signatureResult := strings.TrimSpace(request.SignatureVerificationState)
	if strings.TrimSpace(request.SignatureRef) == "" {
		signatureResult = referenceVerifierSignatureMissing
	}

	schemaResult := strings.TrimSpace(request.SchemaVerificationState)
	if strings.TrimSpace(request.SchemaVersion) == "" {
		schemaResult = referenceVerifierSchemaUnsupported
	}

	scopeResult := strings.TrimSpace(request.ScopeVerificationState)
	if request.RequestedScope == "" || request.ExpectedOutputBoundary == "" || request.RequestedScope != request.ExpectedOutputBoundary {
		scopeResult = referenceVerifierScopeMismatch
	}

	outputBoundaryResult := strings.TrimSpace(request.OutputBoundaryVerificationState)
	if request.ExpectedOutputBoundary == "" {
		outputBoundaryResult = referenceVerifierOutputBoundaryUnknown
	}

	verifiedAt := request.VerificationTime
	if _, err := time.Parse(time.RFC3339, request.VerificationTime); err != nil {
		verifiedAt = request.VerificationTime
	}

	diagnostic := deriveReferenceVerifierDiagnostic(
		request.ProofType,
		digestResult,
		signatureResult,
		schemaResult,
		scopeResult,
		strings.TrimSpace(request.FreshnessVerificationState),
		strings.TrimSpace(request.TrustRootVerificationState),
		strings.TrimSpace(request.IssuerVerificationState),
		strings.TrimSpace(request.CompatibilityEvaluationState),
		strings.TrimSpace(request.RevocationEvaluationState),
		strings.TrimSpace(request.SupersessionEvaluationState),
		strings.TrimSpace(request.LineageVerificationState),
		outputBoundaryResult,
	)
	overall := deriveReferenceVerifierOverallResult(diagnostic, request.Caveats)

	limitations := []string{
		"Reference verifier output remains bounded to the supplied proof envelope, schema line, scope, trust-root material, freshness inputs, and compatibility descriptors.",
		"Verification output is advisory and does not create canonical truth, deployment approval, publication authority, or certification authority.",
	}
	if !request.ClaimsActualCryptoVerification {
		limitations = append(limitations, "If repository-level cryptographic primitives are not invoked, digest and signature handling remains modeled verification semantics and must not be read as claimed real cryptographic validation.")
	}
	if request.StrictFailClosed && overall == referenceVerifierOverallWarnings {
		limitations = append(limitations, "Strict fail-closed mode keeps warning-bearing outputs non-clean and prevents them from being treated as clean verified results.")
	}

	return ReferenceVerificationResult{
		ResultID:               "reference-verifier-result-" + firstNonEmpty(request.RequestID, "unknown"),
		RequestID:              request.RequestID,
		VerifierVersion:        "reference-verifier/vala-2026.04",
		ProofType:              request.ProofType,
		SchemaVersion:          request.SchemaVersion,
		Scope:                  request.RequestedScope,
		OutputBoundary:         request.ExpectedOutputBoundary,
		OverallResult:          overall,
		DiagnosticClass:        diagnostic,
		DigestResult:           digestResult,
		SignatureResult:        signatureResult,
		SchemaResult:           schemaResult,
		ScopeResult:            scopeResult,
		FreshnessResult:        strings.TrimSpace(request.FreshnessVerificationState),
		TrustRootResult:        strings.TrimSpace(request.TrustRootVerificationState),
		IssuerResult:           strings.TrimSpace(request.IssuerVerificationState),
		CompatibilityResult:    strings.TrimSpace(request.CompatibilityEvaluationState),
		RevocationResult:       strings.TrimSpace(request.RevocationEvaluationState),
		SupersessionResult:     strings.TrimSpace(request.SupersessionEvaluationState),
		LineageResult:          strings.TrimSpace(request.LineageVerificationState),
		OutputBoundaryResult:   outputBoundaryResult,
		EvidenceRefs:           uniqueStrings(request.EvidenceRefs),
		Caveats:                uniqueStrings(request.Caveats),
		Limitations:            uniqueStrings(limitations),
		ProjectionDisclaimer:   strings.TrimSpace(request.ProjectionDisclaimer),
		VerifiedAt:             verifiedAt,
		TruthOutsideScopeClaim: request.TruthOutsideScopeClaim,
	}
}

func deriveReferenceVerifierDiagnostic(
	proofType, digestResult, signatureResult, schemaResult, scopeResult, freshnessResult, trustRootResult, issuerResult, compatibilityResult, revocationResult, supersessionResult, lineageResult, outputBoundaryResult string,
) string {
	if !containsString([]string{"signed_attestation_envelope", "sealed_artifact_envelope", "lineage_bundle_envelope"}, proofType) {
		return referenceVerifierDiagnosticUnsupportedProofType
	}
	if signatureResult == referenceVerifierSignatureInvalid {
		return referenceVerifierDiagnosticInvalidSignature
	}
	if digestResult == referenceVerifierDigestMismatch {
		return referenceVerifierDiagnosticDigestMismatch
	}
	if schemaResult == referenceVerifierSchemaUnsupported {
		return referenceVerifierDiagnosticUnsupportedSchema
	}
	if schemaResult == referenceVerifierSchemaMismatch {
		return referenceVerifierDiagnosticSchemaMismatch
	}
	if trustRootResult == "revoked" || trustRootResult == "expired" || issuerResult == referenceVerifierIssuerRevoked || revocationResult == referenceVerifierRevocationRevoked {
		return referenceVerifierDiagnosticRevokedIssuer
	}
	if freshnessResult == "expired" {
		return referenceVerifierDiagnosticExpiredArtifact
	}
	if freshnessResult == "stale" {
		return referenceVerifierDiagnosticStaleArtifact
	}
	if supersessionResult == referenceVerifierSupersessionSuperseded {
		return referenceVerifierDiagnosticSupersededProof
	}
	if outputBoundaryResult == referenceVerifierOutputBoundaryViolation {
		return referenceVerifierDiagnosticRedactionViolation
	}
	if signatureResult == referenceVerifierSignatureUnsupported ||
		digestResult == referenceVerifierDigestUnsupported ||
		trustRootResult == "unsupported" ||
		trustRootResult == "unknown" ||
		revocationResult == referenceVerifierRevocationMaterialMissing {
		return referenceVerifierDiagnosticInsufficientTrust
	}
	if signatureResult == referenceVerifierSignatureMissing ||
		digestResult == referenceVerifierDigestMissing ||
		lineageResult == referenceVerifierLineageMissing {
		return referenceVerifierDiagnosticIncompleteArtifact
	}
	if scopeResult == referenceVerifierScopeMismatch || scopeResult == referenceVerifierScopeUnsupported {
		return referenceVerifierDiagnosticScopeMismatch
	}
	if compatibilityResult == "compatible_with_warnings" || compatibilityResult == "deprecated" || compatibilityResult == "superseded" || trustRootResult == "trusted_with_warnings" {
		return referenceVerifierDiagnosticCompatibilityWarning
	}
	validStates := containsString([]string{referenceVerifierDigestMatch, referenceVerifierDigestMismatch, referenceVerifierDigestMissing, referenceVerifierDigestUnsupported}, digestResult) &&
		containsString([]string{referenceVerifierSignatureValid, referenceVerifierSignatureInvalid, referenceVerifierSignatureMissing, referenceVerifierSignatureUnsupported}, signatureResult) &&
		containsString([]string{referenceVerifierSchemaValid, referenceVerifierSchemaMismatch, referenceVerifierSchemaUnsupported}, schemaResult) &&
		containsString([]string{referenceVerifierScopeValid, referenceVerifierScopeMismatch, referenceVerifierScopeUnsupported}, scopeResult) &&
		containsString([]string{"fresh", "stale", "expired"}, freshnessResult) &&
		containsString([]string{"trusted", "trusted_with_warnings", "rotated", "revoked", "expired", "unsupported", "unknown"}, trustRootResult) &&
		containsString([]string{referenceVerifierIssuerTrusted, referenceVerifierIssuerRevoked, referenceVerifierIssuerUnknown}, issuerResult) &&
		containsString([]string{"compatible", "compatible_with_warnings", "deprecated", "superseded", "unsupported", "unknown"}, compatibilityResult) &&
		containsString([]string{referenceVerifierRevocationNotRevoked, referenceVerifierRevocationRevoked, referenceVerifierRevocationMaterialMissing, referenceVerifierRevocationUnknown}, revocationResult) &&
		containsString([]string{referenceVerifierSupersessionCurrent, referenceVerifierSupersessionSuperseded, referenceVerifierSupersessionUnknown}, supersessionResult) &&
		containsString([]string{referenceVerifierLineagePresent, referenceVerifierLineageMissing, referenceVerifierLineageUnknown}, lineageResult) &&
		containsString([]string{referenceVerifierOutputBoundaryValid, referenceVerifierOutputBoundaryViolation, referenceVerifierOutputBoundaryUnknown}, outputBoundaryResult)
	if !validStates ||
		freshnessResult == "unknown" ||
		freshnessResult == "unsupported" ||
		issuerResult == referenceVerifierIssuerUnknown ||
		compatibilityResult == "unknown" ||
		revocationResult == referenceVerifierRevocationUnknown ||
		supersessionResult == referenceVerifierSupersessionUnknown ||
		lineageResult == referenceVerifierLineageUnknown ||
		outputBoundaryResult == referenceVerifierOutputBoundaryUnknown {
		return referenceVerifierDiagnosticUnknown
	}
	return referenceVerifierDiagnosticVerified
}

func deriveReferenceVerifierOverallResult(diagnostic string, caveats []string) string {
	switch diagnostic {
	case referenceVerifierDiagnosticVerified:
		return referenceVerifierOverallVerified
	case referenceVerifierDiagnosticCompatibilityWarning:
		if len(uniqueStrings(caveats)) == 0 {
			return referenceVerifierOverallIncomplete
		}
		return referenceVerifierOverallWarnings
	case referenceVerifierDiagnosticInvalidSignature,
		referenceVerifierDiagnosticDigestMismatch,
		referenceVerifierDiagnosticSchemaMismatch,
		referenceVerifierDiagnosticScopeMismatch,
		referenceVerifierDiagnosticRedactionViolation:
		return referenceVerifierOverallInvalid
	case referenceVerifierDiagnosticUnsupportedSchema,
		referenceVerifierDiagnosticUnsupportedProofType:
		return referenceVerifierOverallUnsupported
	case referenceVerifierDiagnosticStaleArtifact,
		referenceVerifierDiagnosticExpiredArtifact:
		return referenceVerifierOverallStale
	case referenceVerifierDiagnosticRevokedIssuer:
		return referenceVerifierOverallRevoked
	case referenceVerifierDiagnosticSupersededProof:
		return referenceVerifierOverallSuperseded
	case referenceVerifierDiagnosticInsufficientTrust,
		referenceVerifierDiagnosticIncompleteArtifact:
		return referenceVerifierOverallIncomplete
	default:
		return referenceVerifierOverallUnknown
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			return value
		}
	}
	return ""
}

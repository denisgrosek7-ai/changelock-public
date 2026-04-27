package verifier

import (
	"reflect"
	"testing"
)

func activeReferenceVerifierRequest() ReferenceVerifierRequest {
	return ReferenceVerifierRequest{
		RequestID:                       "reference-verifier-request-vala-001",
		VerifierContractRef:             "/v1/verifier-ecosystem/val0/contract",
		ProofEnvelopeRef:                "/v1/verifier-ecosystem/val0/proof-envelope",
		ArtifactRef:                     "artifact:reference-verifier-bundle",
		ArtifactDigest:                  "2d8a84e4e1ec70e3cf6d0e6d7f1b9bc8a1e33f6f7bb4ab4fc1c2b6d9e9ac0021",
		ArtifactDigestAlgorithm:         "sha256",
		SignatureRef:                    "signature:reference-verifier-bundle",
		IssuerRef:                       "issuer:reference-signer",
		TrustRootRef:                    "trust-root:reference-program",
		SchemaVersion:                   "changelock.verifier.proof_envelope.v1",
		ProofType:                       "signed_attestation_envelope",
		RequestedScope:                  "auditor_safe",
		VerificationTime:                "2026-04-27T08:04:00Z",
		ExpectedOutputBoundary:          "auditor_safe",
		CompatibilityPolicyRef:          "compatibility:reference-proof-envelope",
		RevocationMaterialRef:           "revocation:reference-proof-envelope",
		SupersessionMaterialRef:         "supersession:reference-proof-envelope",
		EvidenceRefs:                    []string{"evidence:verifier-input-vala-001", "evidence:verifier-engine-vala-001", "evidence:verifier-report-vala-001"},
		DigestVerificationState:         "digest_match",
		SignatureVerificationState:      "signature_valid",
		SchemaVerificationState:         "schema_valid",
		ScopeVerificationState:          "scope_valid",
		FreshnessVerificationState:      "fresh",
		TrustRootVerificationState:      "trusted",
		IssuerVerificationState:         "issuer_trusted",
		CompatibilityEvaluationState:    "compatible",
		RevocationEvaluationState:       "not_revoked",
		SupersessionEvaluationState:     "current",
		LineageVerificationState:        "lineage_present",
		OutputBoundaryVerificationState: "output_boundary_valid",
		ReportFormat:                    "json",
		StrictFailClosed:                true,
		Caveats:                         []string{"verification remains bounded to declared scope"},
		ProjectionDisclaimer:            "projection_only not_canonical_truth reference_verifier_tooling advisory_projection",
	}
}

func TestVerifyReferenceVerifierRequestDeterministic(t *testing.T) {
	request := activeReferenceVerifierRequest()
	first := VerifyReferenceVerifierRequest(request)
	second := VerifyReferenceVerifierRequest(request)

	if !reflect.DeepEqual(first, second) {
		t.Fatalf("expected deterministic verifier result, got %#v and %#v", first, second)
	}
	if first.OverallResult != referenceVerifierOverallVerified || first.DiagnosticClass != referenceVerifierDiagnosticVerified {
		t.Fatalf("expected fully valid request to verify cleanly, got %#v", first)
	}
}

func TestVerifyReferenceVerifierRequestInvalidInputsStayNonVerified(t *testing.T) {
	request := activeReferenceVerifierRequest()
	request.SignatureVerificationState = referenceVerifierSignatureInvalid
	invalidSignature := VerifyReferenceVerifierRequest(request)
	if invalidSignature.DiagnosticClass != referenceVerifierDiagnosticInvalidSignature || invalidSignature.OverallResult == referenceVerifierOverallVerified {
		t.Fatalf("expected invalid signature not to verify, got %#v", invalidSignature)
	}

	request = activeReferenceVerifierRequest()
	request.ArtifactDigestAlgorithm = "sha3"
	unsupportedDigest := VerifyReferenceVerifierRequest(request)
	if unsupportedDigest.OverallResult == referenceVerifierOverallVerified {
		t.Fatalf("expected unsupported digest algorithm not to verify, got %#v", unsupportedDigest)
	}

	request = activeReferenceVerifierRequest()
	request.EvidenceRefs = nil
	noEvidence := VerifyReferenceVerifierRequest(request)
	if len(noEvidence.EvidenceRefs) != 0 {
		t.Fatalf("expected missing evidence to remain missing, got %#v", noEvidence)
	}
}

func TestVerifyReferenceVerifierRequestDoesNotGenerateFakeEvidence(t *testing.T) {
	request := activeReferenceVerifierRequest()
	request.EvidenceRefs = []string{"evidence:verifier-input-vala-001"}
	result := VerifyReferenceVerifierRequest(request)
	if len(result.EvidenceRefs) != 1 || result.EvidenceRefs[0] != "evidence:verifier-input-vala-001" {
		t.Fatalf("expected verifier not to generate fake evidence, got %#v", result)
	}
}

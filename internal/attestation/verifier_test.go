package attestation

import "testing"

func TestVerifierAllowsCredentialReleaseWhenVerified(t *testing.T) {
	result := NewVerifier().Verify(VerificationRequest{
		SubjectRef:               "cluster-a/acme-prod/Deployment/api",
		Provider:                 "sgx",
		QuoteType:                "sgx_quote",
		Measurement:              "m-1",
		LifecycleState:           "active",
		SubstrateClass:           "confidential",
		TrustedMeasurements:      []string{"m-1"},
		RequireCredentialRelease: true,
	})

	if result.CurrentState != VerdictVerified || !result.CredentialReleaseAllowed {
		t.Fatalf("expected verified credential release, got %#v", result)
	}
}

func TestVerifierRejectsRevokedLifecycle(t *testing.T) {
	result := NewVerifier().Verify(VerificationRequest{
		Provider:       "tdx",
		QuoteType:      "tdx_quote",
		Measurement:    "m-1",
		LifecycleState: "revoked",
		SubstrateClass: "confidential",
	})

	if result.CurrentState != VerdictMismatch || result.CredentialReleaseAllowed {
		t.Fatalf("expected mismatch without credential release, got %#v", result)
	}
}

func TestVerifierRejectsUntrustedMeasurement(t *testing.T) {
	result := NewVerifier().Verify(VerificationRequest{
		Provider:            "sev",
		QuoteType:           "snp_report",
		Measurement:         "bad-measurement",
		LifecycleState:      "active",
		SubstrateClass:      "confidential",
		TrustedMeasurements: []string{"good-measurement"},
	})

	if result.CurrentState != VerdictMismatch {
		t.Fatalf("expected mismatch, got %#v", result)
	}
}

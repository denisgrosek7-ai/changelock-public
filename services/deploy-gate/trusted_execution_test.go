package main

import (
	"strings"
	"testing"

	"github.com/denisgrosek/changelock/internal/signingidentity"
	"github.com/denisgrosek/changelock/internal/verify"
)

func TestAdmissionReviewAllowsTrustedExecutionProfileMatch(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	previousVerifier := artifactVerifier
	previousSigner := signerIdentityEnforcer
	artifactVerifier = fakeArtifactVerifier{
		result: verify.ArtifactVerification{
			SignatureValid:   true,
			AttestationValid: true,
			VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
			VerifiedRepo:     "my-org/acme-app",
			VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
			VerifiedSubject:  "repo:my-org/acme-app",
			VerifiedDigest:   "sha256:abc123",
		},
	}
	signerIdentityEnforcer = &fakeSignerIdentityEvaluator{enabled: false, mode: signingidentity.EnforcementDisabled}
	defer func() {
		artifactVerifier = previousVerifier
		signerIdentityEnforcer = previousSigner
	}()

	review := trustedAdmissionReview()
	review.Request.UID = "allow-profile-match"
	review.Request.Object.Metadata.Name = "payments-api"
	review.Request.Object.Metadata.Annotations["changelock.io/trusted-execution-profile"] = "confidential-strict"
	review.Request.Object.Metadata.Annotations["changelock.io/cluster-id"] = "cluster-a"
	review.Request.Object.Metadata.Annotations["changelock.io/node-id"] = "node-a"
	review.Request.Object.Metadata.Annotations["changelock.io/node-substrate-class"] = "confidential"
	review.Request.Object.Metadata.Annotations["changelock.io/attestation-provider"] = "sgx"
	review.Request.Object.Metadata.Annotations["changelock.io/attestation-measurement"] = "m-1"
	review.Request.Object.Metadata.Annotations["changelock.io/trusted-measurement"] = "m-1"

	response := executeAdmissionRequest(t, review)
	if !response.Response.Allowed {
		t.Fatalf("expected allow for matching trusted execution profile, got %#v", response.Response)
	}
}

func TestAdmissionReviewDeniesTrustedExecutionProfileMismatch(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	previousVerifier := artifactVerifier
	artifactVerifier = fakeArtifactVerifier{
		result: verify.ArtifactVerification{
			SignatureValid:   true,
			AttestationValid: true,
			VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
			VerifiedRepo:     "my-org/acme-app",
			VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
			VerifiedSubject:  "repo:my-org/acme-app",
			VerifiedDigest:   "sha256:abc123",
		},
	}
	defer func() {
		artifactVerifier = previousVerifier
	}()

	review := trustedAdmissionReview()
	review.Request.UID = "deny-profile-mismatch"
	review.Request.Object.Metadata.Name = "payments-api"
	review.Request.Object.Metadata.Annotations["changelock.io/trusted-execution-profile"] = "confidential-strict"
	review.Request.Object.Metadata.Annotations["changelock.io/cluster-id"] = "cluster-a"
	review.Request.Object.Metadata.Annotations["changelock.io/node-id"] = "node-a"
	review.Request.Object.Metadata.Annotations["changelock.io/node-substrate-class"] = "confidential"
	review.Request.Object.Metadata.Annotations["changelock.io/attestation-provider"] = "sgx"
	review.Request.Object.Metadata.Annotations["changelock.io/attestation-measurement"] = "m-wrong"
	review.Request.Object.Metadata.Annotations["changelock.io/trusted-measurement"] = "m-expected"

	response := executeAdmissionRequest(t, review)
	if response.Response.Allowed {
		t.Fatalf("expected denial for mismatched trusted execution profile, got %#v", response.Response)
	}
	if response.Response.Status == nil || response.Response.Status.Message == "" || !containsReason(response.Response.Status.Message, "trusted execution profile confidential-strict rejected workload") {
		t.Fatalf("expected trusted execution mismatch in denial message, got %#v", response.Response)
	}
}

func TestAdmissionReviewDeniesUnknownTrustedExecutionProfile(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	previousVerifier := artifactVerifier
	previousSigner := signerIdentityEnforcer
	artifactVerifier = fakeArtifactVerifier{
		result: verify.ArtifactVerification{
			SignatureValid:   true,
			AttestationValid: true,
			VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
			VerifiedRepo:     "my-org/acme-app",
			VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
			VerifiedSubject:  "repo:my-org/acme-app",
			VerifiedDigest:   "sha256:abc123",
		},
	}
	signerIdentityEnforcer = &fakeSignerIdentityEvaluator{enabled: false, mode: signingidentity.EnforcementDisabled}
	defer func() {
		artifactVerifier = previousVerifier
		signerIdentityEnforcer = previousSigner
	}()

	review := trustedAdmissionReview()
	review.Request.UID = "deny-profile-unknown"
	review.Request.Object.Metadata.Name = "payments-api"
	review.Request.Object.Metadata.Annotations["changelock.io/trusted-execution-profile"] = "does-not-exist"

	response := executeAdmissionRequest(t, review)
	if response.Response.Allowed {
		t.Fatalf("expected denial for unknown trusted execution profile, got %#v", response.Response)
	}
	if response.Response.Status == nil || response.Response.Status.Message == "" || !containsReason(response.Response.Status.Message, "trusted execution profile is unknown") {
		t.Fatalf("expected unknown trusted execution profile in denial message, got %#v", response.Response)
	}
}

func TestAdmissionReviewDeniesInvalidTrustedExecutionAttestationExpiry(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")
	previousVerifier := artifactVerifier
	previousSigner := signerIdentityEnforcer
	artifactVerifier = fakeArtifactVerifier{
		result: verify.ArtifactVerification{
			SignatureValid:   true,
			AttestationValid: true,
			VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
			VerifiedRepo:     "my-org/acme-app",
			VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
			VerifiedSubject:  "repo:my-org/acme-app",
			VerifiedDigest:   "sha256:abc123",
		},
	}
	signerIdentityEnforcer = &fakeSignerIdentityEvaluator{enabled: false, mode: signingidentity.EnforcementDisabled}
	defer func() {
		artifactVerifier = previousVerifier
		signerIdentityEnforcer = previousSigner
	}()

	review := trustedAdmissionReview()
	review.Request.UID = "deny-invalid-attestation-expiry"
	review.Request.Object.Metadata.Name = "payments-api"
	review.Request.Object.Metadata.Annotations["changelock.io/trusted-execution-profile"] = "confidential-strict"
	review.Request.Object.Metadata.Annotations["changelock.io/attestation-valid-until"] = "not-a-time"

	response := executeAdmissionRequest(t, review)
	if response.Response.Allowed {
		t.Fatalf("expected denial for invalid attestation expiry annotation, got %#v", response.Response)
	}
	if response.Response.Status == nil || response.Response.Status.Message == "" || !containsReason(response.Response.Status.Message, "attestation expiry annotation is invalid") {
		t.Fatalf("expected invalid attestation expiry in denial message, got %#v", response.Response)
	}
}

func containsReason(value, expected string) bool {
	return value != "" && expected != "" && strings.Contains(value, expected)
}

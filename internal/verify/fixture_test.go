package verify

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestFixtureVerifierReturnsConfiguredResult(t *testing.T) {
	path := filepath.Join(t.TempDir(), "fixture.yaml")
	data := []byte(`
artifacts:
  - image: ghcr.io/my-org/acme-app@sha256:abc123
    expectedRepository: my-org/acme-app
    result:
      signatureValid: true
      attestationValid: true
      verifiedIdentity: https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main
      verifiedIssuer: https://token.actions.githubusercontent.com
      verifiedRepo: my-org/acme-app
      verifiedWorkflow: .github/workflows/build-sign-attest.yml
`)
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	verifier, err := NewFixtureVerifier(path)
	if err != nil {
		t.Fatalf("NewFixtureVerifier() error = %v", err)
	}

	result, err := verifier.VerifyArtifact(context.Background(), ArtifactVerificationRequest{
		Image:              "ghcr.io/my-org/acme-app@sha256:abc123",
		ExpectedRepository: "my-org/acme-app",
		ExpectedRef:        "refs/heads/main",
		ExpectedCommitSHA:  "abc123",
	})
	if err != nil {
		t.Fatalf("VerifyArtifact() error = %v", err)
	}
	if !result.SignatureValid || !result.AttestationValid {
		t.Fatalf("expected valid fixture result, got %#v", result)
	}
	if result.VerifiedDigest != "sha256:abc123" || result.VerifiedCommitSHA != "abc123" {
		t.Fatalf("expected defaults to be populated, got %#v", result)
	}
}

func TestFixtureVerifierReturnsConfiguredError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "fixture.yaml")
	data := []byte(`
artifacts:
  - image: ghcr.io/my-org/acme-app@sha256:def456
    error: simulated verification failure
`)
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	verifier, err := NewFixtureVerifier(path)
	if err != nil {
		t.Fatalf("NewFixtureVerifier() error = %v", err)
	}

	if _, err := verifier.VerifyArtifact(context.Background(), ArtifactVerificationRequest{Image: "ghcr.io/my-org/acme-app@sha256:def456"}); err == nil {
		t.Fatalf("expected configured error")
	}
}

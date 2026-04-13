package verify

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"
)

type commandRunnerFunc func(ctx context.Context, name string, args ...string) (commandOutput, error)

func (f commandRunnerFunc) Run(ctx context.Context, name string, args ...string) (commandOutput, error) {
	return f(ctx, name, args...)
}

func TestVerifyArtifactValidSignatureAndAttestation(t *testing.T) {
	verifier := &CosignVerifier{
		binary: "cosign",
		runner: commandRunnerFunc(func(_ context.Context, _ string, args ...string) (commandOutput, error) {
			switch args[0] {
			case "verify":
				return commandOutput{Stdout: []byte(`[{"critical":{"image":{"docker-manifest-digest":"sha256:abc123"}}}]`)}, nil
			case "verify-attestation":
				return commandOutput{Stdout: []byte(validAttestationJSON(t, "ghcr.io/my-org/acme-app:deadbeef", "sha256:abc123"))}, nil
			default:
				return commandOutput{}, fmt.Errorf("unexpected command %q", strings.Join(args, " "))
			}
		}),
	}

	result, err := verifier.VerifyArtifact(context.Background(), ArtifactVerificationRequest{
		Image:                   "ghcr.io/my-org/acme-app@sha256:abc123",
		ExpectedRepository:      "my-org/acme-app",
		AllowedSignerIdentities: []string{"https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main"},
		AllowedOIDCIssuers:      []string{"https://token.actions.githubusercontent.com"},
	})
	if err != nil {
		t.Fatalf("VerifyArtifact() error = %v", err)
	}
	if !result.SignatureValid || !result.AttestationValid {
		t.Fatalf("expected verification success, got %#v", result)
	}
	if result.VerifiedRepo != "my-org/acme-app" {
		t.Fatalf("expected verified repo, got %#v", result)
	}
	if result.VerifiedWorkflow != ".github/workflows/build-sign-attest.yml" {
		t.Fatalf("expected verified workflow, got %#v", result)
	}
	if len(result.Reasons) != 0 {
		t.Fatalf("expected no reasons, got %#v", result.Reasons)
	}
}

func TestVerifyArtifactMissingSignatureDenies(t *testing.T) {
	verifier := &CosignVerifier{
		binary: "cosign",
		runner: commandRunnerFunc(func(_ context.Context, _ string, args ...string) (commandOutput, error) {
			if args[0] == "verify" {
				return commandOutput{Stderr: []byte("no signatures found")}, errors.New("exit status 1")
			}
			return commandOutput{Stderr: []byte("no attestations found")}, errors.New("exit status 1")
		}),
	}

	result, err := verifier.VerifyArtifact(context.Background(), ArtifactVerificationRequest{
		Image:                   "ghcr.io/my-org/acme-app@sha256:abc123",
		ExpectedRepository:      "my-org/acme-app",
		AllowedSignerIdentities: []string{"https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main"},
		AllowedOIDCIssuers:      []string{"https://token.actions.githubusercontent.com"},
	})
	if err != nil {
		t.Fatalf("VerifyArtifact() error = %v", err)
	}
	if result.SignatureValid {
		t.Fatalf("expected signature failure, got %#v", result)
	}
	if len(result.Reasons) == 0 {
		t.Fatalf("expected deny reasons")
	}
}

func TestVerifyArtifactMissingAttestationDenies(t *testing.T) {
	verifier := &CosignVerifier{
		binary: "cosign",
		runner: commandRunnerFunc(func(_ context.Context, _ string, args ...string) (commandOutput, error) {
			if args[0] == "verify" {
				return commandOutput{Stdout: []byte(`[{}]`)}, nil
			}
			return commandOutput{Stderr: []byte("no matching attestations")}, errors.New("exit status 1")
		}),
	}

	result, err := verifier.VerifyArtifact(context.Background(), ArtifactVerificationRequest{
		Image:                   "ghcr.io/my-org/acme-app@sha256:abc123",
		ExpectedRepository:      "my-org/acme-app",
		AllowedSignerIdentities: []string{"https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main"},
		AllowedOIDCIssuers:      []string{"https://token.actions.githubusercontent.com"},
	})
	if err != nil {
		t.Fatalf("VerifyArtifact() error = %v", err)
	}
	if result.AttestationValid {
		t.Fatalf("expected attestation failure, got %#v", result)
	}
}

func TestVerifyArtifactIssuerMismatchDenies(t *testing.T) {
	verifier := &CosignVerifier{
		binary: "cosign",
		runner: commandRunnerFunc(func(_ context.Context, _ string, args ...string) (commandOutput, error) {
			if flagValue(args, "--certificate-oidc-issuer") != "https://token.actions.githubusercontent.com" {
				return commandOutput{Stderr: []byte("issuer mismatch")}, errors.New("exit status 1")
			}
			if args[0] == "verify" {
				return commandOutput{Stdout: []byte(`[{}]`)}, nil
			}
			return commandOutput{Stdout: []byte(validAttestationJSON(t, "ghcr.io/my-org/acme-app:deadbeef", "sha256:abc123"))}, nil
		}),
	}

	result, err := verifier.VerifyArtifact(context.Background(), ArtifactVerificationRequest{
		Image:                   "ghcr.io/my-org/acme-app@sha256:abc123",
		ExpectedRepository:      "my-org/acme-app",
		AllowedSignerIdentities: []string{"https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main"},
		AllowedOIDCIssuers:      []string{"https://wrong-issuer.example.com"},
	})
	if err != nil {
		t.Fatalf("VerifyArtifact() error = %v", err)
	}
	if result.SignatureValid || result.AttestationValid {
		t.Fatalf("expected verification failure, got %#v", result)
	}
}

func TestVerifyArtifactDigestMismatchDenies(t *testing.T) {
	verifier := &CosignVerifier{
		binary: "cosign",
		runner: commandRunnerFunc(func(_ context.Context, _ string, args ...string) (commandOutput, error) {
			if args[0] == "verify" {
				return commandOutput{Stdout: []byte(`[{}]`)}, nil
			}
			return commandOutput{Stdout: []byte(validAttestationJSON(t, "ghcr.io/my-org/acme-app:deadbeef", "sha256:ffff"))}, nil
		}),
	}

	result, err := verifier.VerifyArtifact(context.Background(), ArtifactVerificationRequest{
		Image:                   "ghcr.io/my-org/acme-app@sha256:abc123",
		ExpectedRepository:      "my-org/acme-app",
		AllowedSignerIdentities: []string{"https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main"},
		AllowedOIDCIssuers:      []string{"https://token.actions.githubusercontent.com"},
	})
	if err != nil {
		t.Fatalf("VerifyArtifact() error = %v", err)
	}
	if result.AttestationValid {
		t.Fatalf("expected attestation mismatch failure, got %#v", result)
	}
}

func TestVerifyArtifactSubjectMismatchDenies(t *testing.T) {
	verifier := &CosignVerifier{
		binary: "cosign",
		runner: commandRunnerFunc(func(_ context.Context, _ string, args ...string) (commandOutput, error) {
			if args[0] == "verify" {
				return commandOutput{Stdout: []byte(`[{}]`)}, nil
			}
			return commandOutput{Stdout: []byte(validAttestationJSON(t, "ghcr.io/other-org/other-app:deadbeef", "sha256:abc123"))}, nil
		}),
	}

	result, err := verifier.VerifyArtifact(context.Background(), ArtifactVerificationRequest{
		Image:                   "ghcr.io/my-org/acme-app@sha256:abc123",
		ExpectedRepository:      "my-org/acme-app",
		AllowedSignerIdentities: []string{"https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main"},
		AllowedOIDCIssuers:      []string{"https://token.actions.githubusercontent.com"},
	})
	if err != nil {
		t.Fatalf("VerifyArtifact() error = %v", err)
	}
	if result.AttestationValid {
		t.Fatalf("expected subject mismatch failure, got %#v", result)
	}
}

func validAttestationJSON(t *testing.T, subjectName, digest string) string {
	t.Helper()

	statement, err := json.Marshal(inTotoStatement{
		PredicateType: DefaultPredicateType,
		Subject: []inTotoSubject{
			{
				Name: subjectName,
				Digest: map[string]string{
					"sha256": strings.TrimPrefix(digest, "sha256:"),
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	envelope, err := json.Marshal(map[string]string{
		"payload": base64.StdEncoding.EncodeToString(statement),
	})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	return "[" + string(envelope) + "]"
}

func flagValue(args []string, flag string) string {
	for i := 0; i < len(args)-1; i++ {
		if args[i] == flag {
			return args[i+1]
		}
	}
	return ""
}

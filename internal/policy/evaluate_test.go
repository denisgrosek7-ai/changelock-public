package policy

import (
	"testing"

	"github.com/denisgrosek/changelock/internal/verify"
)

func TestEvaluateChangeAllowsApprovedCriticalPathChange(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")

	bundle, err := LoadBundle(DefaultPoliciesDir(), "acme")
	if err != nil {
		t.Fatalf("LoadBundle() error = %v", err)
	}

	decision := EvaluateChange(bundle, ChangeEvaluationRequest{
		Tenant:             "acme",
		Repository:         "my-org/acme-app",
		Branch:             "main",
		PullRequest:        true,
		SignedCommits:      true,
		Approvals:          2,
		SecurityApprovals:  1,
		CodeOwnersApproved: true,
		ChangedFiles: []string{
			"deploy/k8s/api-deployment.yaml",
		},
	})

	if decision.Decision != "ALLOW" {
		t.Fatalf("expected ALLOW, got %#v", decision)
	}
}

func TestEvaluateChangeDeniesUnsignedCommit(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")

	bundle, err := LoadBundle(DefaultPoliciesDir(), "acme")
	if err != nil {
		t.Fatalf("LoadBundle() error = %v", err)
	}

	decision := EvaluateChange(bundle, ChangeEvaluationRequest{
		Tenant:            "acme",
		Repository:        "my-org/acme-app",
		Branch:            "main",
		PullRequest:       true,
		SignedCommits:     false,
		Approvals:         2,
		SecurityApprovals: 1,
		ChangedFiles: []string{
			"services/policy-engine/main.go",
		},
	})

	if decision.Decision != "DENY" {
		t.Fatalf("expected DENY, got %#v", decision)
	}
	if len(decision.Reasons) == 0 {
		t.Fatalf("expected deny reasons")
	}
}

func TestEvaluateArtifactAllowsVerifiedArtifact(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")

	bundle, err := LoadBundle(DefaultPoliciesDir(), "acme")
	if err != nil {
		t.Fatalf("LoadBundle() error = %v", err)
	}

	decision := EvaluateArtifact(bundle, ArtifactEvaluationRequest{
		Tenant:       "acme",
		Repository:   "my-org/acme-app",
		Image:        "ghcr.io/my-org/acme-app@sha256:abc123",
		DigestPinned: true,
		Verification: &verify.ArtifactVerification{
			SignatureValid:   true,
			AttestationValid: true,
			VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
			VerifiedRepo:     "my-org/acme-app",
			VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
			VerifiedSubject:  "repo:my-org/acme-app",
			VerifiedDigest:   "sha256:abc123",
		},
	})

	if decision.Decision != "ALLOW" {
		t.Fatalf("expected ALLOW, got %#v", decision)
	}
}

func TestEvaluateArtifactDeniesWorkflowMismatchFromVerifiedFacts(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")

	bundle, err := LoadBundle(DefaultPoliciesDir(), "acme")
	if err != nil {
		t.Fatalf("LoadBundle() error = %v", err)
	}

	decision := EvaluateArtifact(bundle, ArtifactEvaluationRequest{
		Tenant:     "acme",
		Repository: "my-org/acme-app",
		Image:      "ghcr.io/my-org/acme-app@sha256:abc123",
		Verification: &verify.ArtifactVerification{
			SignatureValid:   true,
			AttestationValid: true,
			VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/other.yml@refs/heads/main",
			VerifiedRepo:     "my-org/acme-app",
			VerifiedWorkflow: ".github/workflows/other.yml",
			VerifiedSubject:  "repo:my-org/acme-app",
			VerifiedDigest:   "sha256:abc123",
		},
	})

	if decision.Decision != "DENY" {
		t.Fatalf("expected DENY, got %#v", decision)
	}
}

func TestEvaluateArtifactDeniesRepoMismatchFromVerifiedFacts(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")

	bundle, err := LoadBundle(DefaultPoliciesDir(), "acme")
	if err != nil {
		t.Fatalf("LoadBundle() error = %v", err)
	}

	decision := EvaluateArtifact(bundle, ArtifactEvaluationRequest{
		Tenant:     "acme",
		Repository: "my-org/acme-app",
		Image:      "ghcr.io/my-org/acme-app@sha256:abc123",
		Verification: &verify.ArtifactVerification{
			SignatureValid:   true,
			AttestationValid: true,
			VerifiedIdentity: "https://github.com/my-org/other-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
			VerifiedRepo:     "my-org/other-app",
			VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
			VerifiedSubject:  "repo:my-org/other-app",
			VerifiedDigest:   "sha256:abc123",
		},
	})

	if decision.Decision != "DENY" {
		t.Fatalf("expected DENY, got %#v", decision)
	}
}

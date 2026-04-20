package policy

import (
	"fmt"
	"testing"

	"github.com/denisgrosek/changelock/internal/verify"
)

type policyBenchmarkProfile struct {
	name  string
	count int
}

var policyBenchmarkProfiles = []policyBenchmarkProfile{
	{name: "small", count: 10},
	{name: "medium", count: 100},
	{name: "large", count: 1000},
}

func BenchmarkEvaluateChange(b *testing.B) {
	b.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")

	bundle, err := LoadBundle(DefaultPoliciesDir(), "acme")
	if err != nil {
		b.Fatalf("LoadBundle() error = %v", err)
	}

	for _, profile := range policyBenchmarkProfiles {
		profile := profile
		request := ChangeEvaluationRequest{
			Tenant:             "acme",
			Repository:         "my-org/acme-app",
			Branch:             "main",
			PullRequest:        true,
			SignedCommits:      true,
			Approvals:          2,
			SecurityApprovals:  1,
			CodeOwnersApproved: true,
			ChangedFiles:       benchmarkChangedFiles(profile.count),
		}
		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				decision := EvaluateChange(bundle, request)
				if decision.Decision == "" {
					b.Fatal("expected decision")
				}
			}
		})
	}
}

func BenchmarkEvaluateArtifact(b *testing.B) {
	b.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")

	baseBundle, err := LoadBundle(DefaultPoliciesDir(), "acme")
	if err != nil {
		b.Fatalf("LoadBundle() error = %v", err)
	}

	for _, profile := range policyBenchmarkProfiles {
		profile := profile
		bundle := *baseBundle
		bundle.Artifact.Spec.AllowedSignerIdentities = benchmarkArtifactValues("https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main", "signer", profile.count)
		bundle.Artifact.Spec.AllowedWorkflowFiles = benchmarkArtifactValues(".github/workflows/build-sign-attest.yml", "workflow", profile.count)
		bundle.Artifact.Spec.AllowedSubjects = benchmarkArtifactValues("repo:my-org/acme-app", "subject", profile.count)

		request := ArtifactEvaluationRequest{
			Tenant:     "acme",
			Repository: "my-org/acme-app",
			Image:      "ghcr.io/my-org/acme-app@sha256:abc123",
			Verification: &verify.ArtifactVerification{
				SignatureValid:   true,
				AttestationValid: true,
				VerifiedIdentity: "https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main",
				VerifiedRepo:     "my-org/acme-app",
				VerifiedWorkflow: ".github/workflows/build-sign-attest.yml",
				VerifiedSubject:  "repo:my-org/acme-app",
				VerifiedDigest:   "sha256:abc123",
			},
		}

		b.Run(profile.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				decision := EvaluateArtifact(&bundle, request)
				if decision.Decision == "" {
					b.Fatal("expected decision")
				}
			}
		})
	}
}

func benchmarkChangedFiles(count int) []string {
	files := make([]string, 0, count)
	for i := 0; i < count-1; i++ {
		files = append(files, fmt.Sprintf("services/api/file-%04d.go", i))
	}
	files = append(files, "deploy/k8s/api-deployment.yaml")
	return files
}

func benchmarkArtifactValues(match, prefix string, count int) []string {
	values := make([]string, 0, count)
	for i := 0; i < count-1; i++ {
		values = append(values, fmt.Sprintf("%s-%04d", prefix, i))
	}
	values = append(values, match)
	return values
}

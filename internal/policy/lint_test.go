package policy

import (
	"strings"
	"testing"
)

func TestLoadBundleCarriesDeterministicShadowFindings(t *testing.T) {
	t.Setenv("CHANGELOCK_POLICIES_DIR", "../../policies")

	bundle, err := LoadBundle(DefaultPoliciesDir(), "acme")
	if err != nil {
		t.Fatalf("LoadBundle() error = %v", err)
	}

	found := false
	for _, finding := range bundle.LintFindings {
		if finding.Kind == "critical_path_shadow" && finding.Scope == ".github/workflows/**" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected critical_path_shadow finding, got %#v", bundle.LintFindings)
	}
}

func TestLoadBundleRejectsDuplicateRepositoryPolicies(t *testing.T) {
	root := t.TempDir()
	writePolicyFixtures(t, root, map[string]string{
		"global/change-policy.yaml": `metadata:
  name: global-change-policy
spec:
  allowedBranches:
    - main
  requireSignedCommits: true
  requirePullRequest: true
  minimumApprovals: 1
  minimumSecurityApprovals: 1
  criticalPaths: []
  criticalPathRules:
    minimumSecurityApprovals: 1
    requireCodeOwnersApproval: false
  blockForcePushOnProtectedBranches: true
`,
		"global/artifact-policy.yaml": `metadata:
  name: global-artifact-policy
spec:
  allowedRegistries:
    - ghcr.io/my-org/acme-app
  requireDigestPinning: true
  requireProvenance: true
  requireSignature: true
  allowedSignerIdentities:
    - https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main
  allowedWorkflowFiles:
    - .github/workflows/build-sign-attest.yml
  allowedSubjects:
    - repo:my-org/acme-app
`,
		"global/runtime-policy.yaml": `metadata:
  name: global-runtime-policy
spec:
  blockLatestTag: true
  requireReadOnlyRootFilesystem: true
  allowPrivilegeEscalation: false
  allowHostNetwork: false
  allowHostPID: false
  allowHostIPC: false
  requireNonRoot: true
  maxContainerCapabilities: []
`,
		"tenants/acme/tenant.yaml": `metadata:
  name: acme
spec:
  repositories:
    - my-org/acme-app
  environments:
    - prod
  namespaces:
    - acme-prod
`,
		"tenants/acme/repositories.yaml": `repositories:
  - name: my-org/acme-app
    defaultBranch: main
    workflowAllowlist:
      - .github/workflows/build-sign-attest.yml
  - name: my-org/acme-app
    defaultBranch: release
    workflowAllowlist:
      - .github/workflows/release.yml
`,
	})

	_, err := LoadBundle(root, "acme")
	if err == nil {
		t.Fatalf("expected duplicate repository policy error")
	}
	if !strings.Contains(err.Error(), "policy consistency check failed") || !strings.Contains(err.Error(), "shadow") {
		t.Fatalf("expected shadow consistency error, got %v", err)
	}
}

func TestLintBundleWarnsWhenSignerAndProvenanceRulesAreShadowed(t *testing.T) {
	bundle := &Bundle{
		Artifact: ArtifactPolicy{
			Spec: ArtifactPolicySpec{
				RequireSignature:        false,
				RequireProvenance:       false,
				AllowedSignerIdentities: []string{"signer-a"},
				AllowedWorkflowFiles:    []string{".github/workflows/build.yml"},
				AllowedSubjects:         []string{"repo:my-org/acme-app"},
			},
		},
	}

	findings := lintBundle(bundle)
	kinds := map[string]struct{}{}
	for _, finding := range findings {
		kinds[finding.Kind] = struct{}{}
	}
	if _, ok := kinds["artifact_signer_shadow"]; !ok {
		t.Fatalf("expected artifact_signer_shadow finding, got %#v", findings)
	}
	if _, ok := kinds["artifact_provenance_shadow"]; !ok {
		t.Fatalf("expected artifact_provenance_shadow finding, got %#v", findings)
	}
}

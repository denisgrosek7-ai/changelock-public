package policy

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadBundleComputesDeterministicIdentity(t *testing.T) {
	root := t.TempDir()
	writePolicyFixtures(t, root, map[string]string{
		"global/change-policy.yaml":        "metadata:\n  name: global-change-policy\nspec:\n  allowedBranches:\n    - main\n  requireSignedCommits: true\n  requirePullRequest: true\n  minimumApprovals: 1\n  minimumSecurityApprovals: 1\n  criticalPaths: []\n  criticalPathRules:\n    minimumSecurityApprovals: 1\n    requireCodeOwnersApproval: false\n  blockForcePushOnProtectedBranches: true\n",
		"global/artifact-policy.yaml":      "metadata:\n  name: global-artifact-policy\nspec:\n  allowedRegistries:\n    - ghcr.io/\n  requireDigestPinning: true\n  requireProvenance: true\n  requireSignature: true\n  allowedSignerIdentities:\n    - https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main\n  allowedWorkflowFiles:\n    - .github/workflows/build-sign-attest.yml\n  allowedSubjects:\n    - repo:my-org/acme-app\n",
		"global/runtime-policy.yaml":       "metadata:\n  name: global-runtime-policy\nspec:\n  blockLatestTag: true\n  requireReadOnlyRootFilesystem: true\n  allowPrivilegeEscalation: false\n  allowHostNetwork: false\n  allowHostPID: false\n  allowHostIPC: false\n  requireNonRoot: true\n  maxContainerCapabilities: []\n",
		"tenants/acme/tenant.yaml":         "metadata:\n  name: acme\nspec:\n  repositories:\n    - my-org/acme-app\n  environments:\n    - prod\n  namespaces:\n    - acme-prod\n",
		"tenants/acme/repositories.yaml":   "repositories:\n  - name: my-org/acme-app\n    defaultBranch: main\n    workflowAllowlist:\n      - .github/workflows/build-sign-attest.yml\n    releaseBranches:\n      - main\n",
		"tenants/acme/critical-paths.yaml": "criticalPaths:\n  - path: deploy/**\n    securityOwnerGroup: secops\n",
	})

	first, err := LoadBundle(root, "acme")
	if err != nil {
		t.Fatalf("LoadBundle() error = %v", err)
	}
	second, err := LoadBundle(root, "acme")
	if err != nil {
		t.Fatalf("LoadBundle() second error = %v", err)
	}

	if first.BundleID != "tenant:acme" {
		t.Fatalf("expected bundle id tenant:acme, got %q", first.BundleID)
	}
	if first.BundleHash == "" {
		t.Fatal("expected non-empty bundle hash")
	}
	if first.BundleHash != second.BundleHash {
		t.Fatalf("expected deterministic bundle hash, got %q and %q", first.BundleHash, second.BundleHash)
	}
}

func TestLoadBundleHashChangesWhenPolicyContentChanges(t *testing.T) {
	root := t.TempDir()
	writePolicyFixtures(t, root, map[string]string{
		"global/change-policy.yaml":      "metadata:\n  name: global-change-policy\nspec:\n  allowedBranches:\n    - main\n  requireSignedCommits: true\n  requirePullRequest: true\n  minimumApprovals: 1\n  minimumSecurityApprovals: 1\n  criticalPaths: []\n  criticalPathRules:\n    minimumSecurityApprovals: 1\n    requireCodeOwnersApproval: false\n  blockForcePushOnProtectedBranches: true\n",
		"global/artifact-policy.yaml":    "metadata:\n  name: global-artifact-policy\nspec:\n  allowedRegistries:\n    - ghcr.io/\n  requireDigestPinning: true\n  requireProvenance: true\n  requireSignature: true\n  allowedSignerIdentities:\n    - https://github.com/my-org/acme-app/.github/workflows/build-sign-attest.yml@refs/heads/main\n  allowedWorkflowFiles:\n    - .github/workflows/build-sign-attest.yml\n  allowedSubjects:\n    - repo:my-org/acme-app\n",
		"global/runtime-policy.yaml":     "metadata:\n  name: global-runtime-policy\nspec:\n  blockLatestTag: true\n  requireReadOnlyRootFilesystem: true\n  allowPrivilegeEscalation: false\n  allowHostNetwork: false\n  allowHostPID: false\n  allowHostIPC: false\n  requireNonRoot: true\n  maxContainerCapabilities: []\n",
		"tenants/acme/tenant.yaml":       "metadata:\n  name: acme\nspec:\n  repositories:\n    - my-org/acme-app\n  environments:\n    - prod\n  namespaces:\n    - acme-prod\n",
		"tenants/acme/repositories.yaml": "repositories:\n  - name: my-org/acme-app\n    defaultBranch: main\n    workflowAllowlist:\n      - .github/workflows/build-sign-attest.yml\n    releaseBranches:\n      - main\n",
	})

	before, err := LoadBundle(root, "acme")
	if err != nil {
		t.Fatalf("LoadBundle() error = %v", err)
	}

	writePolicyFixtures(t, root, map[string]string{
		"global/runtime-policy.yaml": "metadata:\n  name: global-runtime-policy\nspec:\n  blockLatestTag: false\n  requireReadOnlyRootFilesystem: true\n  allowPrivilegeEscalation: false\n  allowHostNetwork: false\n  allowHostPID: false\n  allowHostIPC: false\n  requireNonRoot: true\n  maxContainerCapabilities: []\n",
	})

	after, err := LoadBundle(root, "acme")
	if err != nil {
		t.Fatalf("LoadBundle() second error = %v", err)
	}

	if before.BundleHash == after.BundleHash {
		t.Fatalf("expected bundle hash to change after policy content change, got %q", before.BundleHash)
	}
}

func writePolicyFixtures(t *testing.T, root string, files map[string]string) {
	t.Helper()
	for relativePath, content := range files {
		path := filepath.Join(root, filepath.FromSlash(relativePath))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("MkdirAll(%q) error = %v", path, err)
		}
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatalf("WriteFile(%q) error = %v", path, err)
		}
	}
}

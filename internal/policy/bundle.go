package policy

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/denisgrosek/changelock/internal/identity"
	"gopkg.in/yaml.v3"
)

type ChangePolicy struct {
	Metadata PolicyMetadata   `yaml:"metadata"`
	Spec     ChangePolicySpec `yaml:"spec"`
}

type ChangePolicySpec struct {
	AllowedBranches                   []string          `yaml:"allowedBranches"`
	RequireSignedCommits              bool              `yaml:"requireSignedCommits"`
	RequirePullRequest                bool              `yaml:"requirePullRequest"`
	MinimumApprovals                  int               `yaml:"minimumApprovals"`
	MinimumSecurityApprovals          int               `yaml:"minimumSecurityApprovals"`
	CriticalPaths                     []string          `yaml:"criticalPaths"`
	CriticalPathRules                 CriticalPathRules `yaml:"criticalPathRules"`
	BlockForcePushOnProtectedBranches bool              `yaml:"blockForcePushOnProtectedBranches"`
}

type CriticalPathRules struct {
	MinimumSecurityApprovals  int  `yaml:"minimumSecurityApprovals"`
	RequireCodeOwnersApproval bool `yaml:"requireCodeOwnersApproval"`
}

type ArtifactPolicy struct {
	Metadata PolicyMetadata     `yaml:"metadata"`
	Spec     ArtifactPolicySpec `yaml:"spec"`
}

type ArtifactPolicySpec struct {
	AllowedRegistries       []string `yaml:"allowedRegistries"`
	RequireDigestPinning    bool     `yaml:"requireDigestPinning"`
	RequireProvenance       bool     `yaml:"requireProvenance"`
	RequireSignature        bool     `yaml:"requireSignature"`
	AllowedSignerIdentities []string `yaml:"allowedSignerIdentities"`
	AllowedWorkflowFiles    []string `yaml:"allowedWorkflowFiles"`
	AllowedSubjects         []string `yaml:"allowedSubjects"`
}

type RuntimePolicy struct {
	Metadata PolicyMetadata    `yaml:"metadata"`
	Spec     RuntimePolicySpec `yaml:"spec"`
}

type PolicyMetadata struct {
	Name string `yaml:"name"`
}

type RuntimePolicySpec struct {
	BlockLatestTag                bool     `yaml:"blockLatestTag"`
	RequireReadOnlyRootFilesystem bool     `yaml:"requireReadOnlyRootFilesystem"`
	AllowPrivilegeEscalation      bool     `yaml:"allowPrivilegeEscalation"`
	AllowHostNetwork              bool     `yaml:"allowHostNetwork"`
	AllowHostPID                  bool     `yaml:"allowHostPID"`
	AllowHostIPC                  bool     `yaml:"allowHostIPC"`
	RequireNonRoot                bool     `yaml:"requireNonRoot"`
	MaxContainerCapabilities      []string `yaml:"maxContainerCapabilities"`
}

type Tenant struct {
	Metadata TenantMetadata `yaml:"metadata"`
	Spec     TenantSpec     `yaml:"spec"`
}

type TenantMetadata struct {
	Name string `yaml:"name"`
}

type TenantSpec struct {
	Repositories []string `yaml:"repositories"`
	Environments []string `yaml:"environments"`
	Namespaces   []string `yaml:"namespaces"`
}

type RepositoryPolicies struct {
	Repositories []RepositoryPolicy `yaml:"repositories"`
}

type RepositoryPolicy struct {
	Name              string   `yaml:"name"`
	DefaultBranch     string   `yaml:"defaultBranch"`
	WorkflowAllowlist []string `yaml:"workflowAllowlist"`
	ReleaseBranches   []string `yaml:"releaseBranches"`
}

type CriticalPaths struct {
	CriticalPaths []CriticalPathEntry `yaml:"criticalPaths"`
}

type CriticalPathEntry struct {
	Path               string `yaml:"path"`
	SecurityOwnerGroup string `yaml:"securityOwnerGroup"`
}

type Bundle struct {
	Change            ChangePolicy
	Artifact          ArtifactPolicy
	Runtime           RuntimePolicy
	Tenant            Tenant
	RepositoryConfigs map[string]RepositoryPolicy
	RepositoryEntries []RepositoryPolicy
	CriticalPaths     []CriticalPathEntry
	BundleID          string
	BundleHash        string
	LintFindings      []LintFinding
}

func DefaultPoliciesDir() string {
	if dir := os.Getenv("CHANGELOCK_POLICIES_DIR"); dir != "" {
		return dir
	}
	return "policies"
}

func LoadBundle(policiesDir, tenant string) (*Bundle, error) {
	if tenant == "" {
		tenant = "acme"
	}

	bundle := &Bundle{
		RepositoryConfigs: map[string]RepositoryPolicy{},
	}
	rawFiles := map[string][]byte{}

	if err := loadBundleYAML(policiesDir, "global/change-policy.yaml", &bundle.Change, rawFiles); err != nil {
		return nil, fmt.Errorf("load global change policy: %w", err)
	}
	if err := loadBundleYAML(policiesDir, "global/artifact-policy.yaml", &bundle.Artifact, rawFiles); err != nil {
		return nil, fmt.Errorf("load global artifact policy: %w", err)
	}
	if err := loadBundleYAML(policiesDir, "global/runtime-policy.yaml", &bundle.Runtime, rawFiles); err != nil {
		return nil, fmt.Errorf("load global runtime policy: %w", err)
	}

	tenantDir := filepath.ToSlash(filepath.Join("tenants", tenant))
	if err := loadBundleYAML(policiesDir, filepath.ToSlash(filepath.Join(tenantDir, "tenant.yaml")), &bundle.Tenant, rawFiles); err != nil {
		return nil, fmt.Errorf("load tenant policy: %w", err)
	}

	var repositories RepositoryPolicies
	if err := loadBundleYAML(policiesDir, filepath.ToSlash(filepath.Join(tenantDir, "repositories.yaml")), &repositories, rawFiles); err != nil {
		return nil, fmt.Errorf("load repository policy: %w", err)
	}
	for _, repository := range repositories.Repositories {
		bundle.RepositoryEntries = append(bundle.RepositoryEntries, repository)
		bundle.RepositoryConfigs[repository.Name] = repository
	}

	var criticalPaths CriticalPaths
	if err := loadBundleYAML(policiesDir, filepath.ToSlash(filepath.Join(tenantDir, "critical-paths.yaml")), &criticalPaths, rawFiles); err == nil {
		bundle.CriticalPaths = criticalPaths.CriticalPaths
	}
	bundle.BundleID = bundleID(bundle.Tenant.Metadata.Name, tenant)
	bundle.BundleHash = identity.CanonicalFileSetHash(rawFiles)
	bundle.LintFindings = lintBundle(bundle)
	if err := lintError(bundle.LintFindings); err != nil {
		return nil, err
	}

	return bundle, nil
}

func loadYAML(path string, dst any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, dst)
}

func loadBundleYAML(policiesDir, relativePath string, dst any, rawFiles map[string][]byte) error {
	path := filepath.Join(policiesDir, filepath.FromSlash(relativePath))
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	rawFiles[filepath.ToSlash(relativePath)] = canonicalPolicyBytes(data)
	return yaml.Unmarshal(data, dst)
}

func canonicalPolicyBytes(data []byte) []byte {
	return []byte(strings.ReplaceAll(string(data), "\r\n", "\n"))
}

func bundleID(metadataName, tenant string) string {
	name := strings.TrimSpace(metadataName)
	if name == "" {
		name = strings.TrimSpace(tenant)
	}
	if name == "" {
		name = "acme"
	}
	return "tenant:" + name
}

func (b *Bundle) RepositoryAllowed(repository string) bool {
	if repository == "" {
		return false
	}
	for _, allowed := range b.Tenant.Spec.Repositories {
		if repository == allowed {
			return true
		}
	}
	return false
}

func (b *Bundle) AllowedWorkflowFiles(repository string) []string {
	if repository != "" {
		if repositoryPolicy, ok := b.RepositoryConfigs[repository]; ok && len(repositoryPolicy.WorkflowAllowlist) > 0 {
			return repositoryPolicy.WorkflowAllowlist
		}
	}
	return b.Artifact.Spec.AllowedWorkflowFiles
}

func (b *Bundle) AllCriticalPathPatterns() []string {
	patterns := make([]string, 0, len(b.Change.Spec.CriticalPaths)+len(b.CriticalPaths))
	patterns = append(patterns, b.Change.Spec.CriticalPaths...)
	for _, path := range b.CriticalPaths {
		patterns = append(patterns, path.Path)
	}
	return patterns
}

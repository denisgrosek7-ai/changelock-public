package policy

import (
	"fmt"
	"sort"
	"strings"
)

type LintSeverity string

const (
	LintSeverityWarning LintSeverity = "warning"
	LintSeverityError   LintSeverity = "error"
)

type LintFinding struct {
	Severity LintSeverity `json:"severity"`
	Kind     string       `json:"kind"`
	Scope    string       `json:"scope"`
	Message  string       `json:"message"`
}

func lintBundle(bundle *Bundle) []LintFinding {
	if bundle == nil {
		return nil
	}

	findings := []LintFinding{}
	findings = append(findings, duplicateValueFindings("change.allowed_branches", bundle.Change.Spec.AllowedBranches)...)
	findings = append(findings, duplicateValueFindings("change.critical_paths", bundle.Change.Spec.CriticalPaths)...)
	findings = append(findings, duplicateValueFindings("artifact.allowed_registries", bundle.Artifact.Spec.AllowedRegistries)...)
	findings = append(findings, duplicateValueFindings("artifact.allowed_signer_identities", bundle.Artifact.Spec.AllowedSignerIdentities)...)
	findings = append(findings, duplicateValueFindings("artifact.allowed_workflow_files", bundle.Artifact.Spec.AllowedWorkflowFiles)...)
	findings = append(findings, duplicateValueFindings("artifact.allowed_subjects", bundle.Artifact.Spec.AllowedSubjects)...)
	findings = append(findings, duplicateCriticalPathEntryFindings(bundle.CriticalPaths)...)
	findings = append(findings, duplicateRepositoryEntryFindings(bundle.RepositoryEntries)...)
	findings = append(findings, criticalPathShadowFindings(bundle.Change.Spec.CriticalPaths, bundle.CriticalPaths)...)

	for _, repository := range bundle.RepositoryEntries {
		scope := fmt.Sprintf("repository.%s.workflow_allowlist", strings.TrimSpace(repository.Name))
		findings = append(findings, duplicateValueFindings(scope, repository.WorkflowAllowlist)...)
	}

	signerConfigured := len(nonEmptyValues(bundle.Artifact.Spec.AllowedSignerIdentities)) > 0
	if !bundle.Artifact.Spec.RequireSignature && signerConfigured {
		findings = append(findings, LintFinding{
			Severity: LintSeverityWarning,
			Kind:     "artifact_signer_shadow",
			Scope:    "artifact.require_signature",
			Message:  "allowed signer identities are configured but signature enforcement is disabled; signer rules will not affect ALLOW or DENY decisions",
		})
	}

	provenanceScoped := len(nonEmptyValues(bundle.Artifact.Spec.AllowedWorkflowFiles)) > 0 || len(nonEmptyValues(bundle.Artifact.Spec.AllowedSubjects)) > 0
	if !bundle.Artifact.Spec.RequireProvenance && provenanceScoped {
		findings = append(findings, LintFinding{
			Severity: LintSeverityWarning,
			Kind:     "artifact_provenance_shadow",
			Scope:    "artifact.require_provenance",
			Message:  "workflow or subject allowlists are configured but provenance enforcement is disabled; verified workflow and subject rules will not affect ALLOW or DENY decisions",
		})
	}

	sort.Slice(findings, func(i, j int) bool {
		if findings[i].Severity != findings[j].Severity {
			return findings[i].Severity == LintSeverityError
		}
		if findings[i].Kind != findings[j].Kind {
			return findings[i].Kind < findings[j].Kind
		}
		if findings[i].Scope != findings[j].Scope {
			return findings[i].Scope < findings[j].Scope
		}
		return findings[i].Message < findings[j].Message
	})

	return findings
}

func lintError(findings []LintFinding) error {
	messages := []string{}
	for _, finding := range findings {
		if finding.Severity != LintSeverityError {
			continue
		}
		messages = append(messages, fmt.Sprintf("%s (%s)", finding.Message, finding.Scope))
	}
	if len(messages) == 0 {
		return nil
	}
	sort.Strings(messages)
	return fmt.Errorf("policy consistency check failed: %s", strings.Join(messages, "; "))
}

func duplicateRepositoryEntryFindings(entries []RepositoryPolicy) []LintFinding {
	findings := []LintFinding{}
	seen := map[string]struct{}{}
	for _, entry := range entries {
		name := strings.TrimSpace(entry.Name)
		if name == "" {
			continue
		}
		if _, ok := seen[name]; ok {
			findings = append(findings, LintFinding{
				Severity: LintSeverityError,
				Kind:     "repository_shadow",
				Scope:    "repository." + name,
				Message:  fmt.Sprintf("repository policy %q is defined more than once; later entries would shadow earlier entries", name),
			})
			continue
		}
		seen[name] = struct{}{}
	}
	return findings
}

func duplicateCriticalPathEntryFindings(entries []CriticalPathEntry) []LintFinding {
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		values = append(values, entry.Path)
	}
	return duplicateValueFindings("tenant.critical_paths", values)
}

func criticalPathShadowFindings(global []string, tenant []CriticalPathEntry) []LintFinding {
	findings := []LintFinding{}
	globalSet := map[string]struct{}{}
	for _, value := range nonEmptyValues(global) {
		globalSet[value] = struct{}{}
	}
	shadowed := map[string]struct{}{}
	for _, entry := range tenant {
		path := strings.TrimSpace(entry.Path)
		if path == "" {
			continue
		}
		if _, ok := globalSet[path]; !ok {
			continue
		}
		if _, seen := shadowed[path]; seen {
			continue
		}
		shadowed[path] = struct{}{}
		findings = append(findings, LintFinding{
			Severity: LintSeverityWarning,
			Kind:     "critical_path_shadow",
			Scope:    path,
			Message:  fmt.Sprintf("critical path pattern %q is defined in both global change policy and tenant critical-paths; evaluation remains deterministic but the later list does not add new coverage", path),
		})
	}
	return findings
}

func duplicateValueFindings(scope string, values []string) []LintFinding {
	findings := []LintFinding{}
	seen := map[string]struct{}{}
	for _, value := range nonEmptyValues(values) {
		if _, ok := seen[value]; ok {
			findings = append(findings, LintFinding{
				Severity: LintSeverityWarning,
				Kind:     "duplicate_entry",
				Scope:    scope,
				Message:  fmt.Sprintf("duplicate policy value %q is configured more than once in %s", value, scope),
			})
			continue
		}
		seen[value] = struct{}{}
	}
	return findings
}

func nonEmptyValues(values []string) []string {
	filtered := make([]string, 0, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue
		}
		filtered = append(filtered, trimmed)
	}
	return filtered
}

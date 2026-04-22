package config

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	ProductionConfigAPIVersion = "changelock.io/v1alpha1"
	ProductionConfigKind       = "ProductionConfig"
	ReportSchemaVersion        = "5.production_config_report.v1"

	IssueSeverityError   = "error"
	IssueSeverityWarning = "warning"

	SyncStateInSync       = "in_sync"
	SyncStateConflict     = "sync_conflict"
	SyncStateStale        = "sync_stale"
	SyncStateLocalOnly    = "local_only"
	SyncStateRemoteOnly   = "remote_only"
	SyncStateUnconfigured = "not_configured"
)

type Metadata struct {
	Name string `yaml:"name" json:"name"`
}

type CLISpec struct {
	Output          string `yaml:"output,omitempty" json:"output,omitempty"`
	Offline         bool   `yaml:"offline,omitempty" json:"offline,omitempty"`
	Scanner         string `yaml:"scanner,omitempty" json:"scanner,omitempty"`
	FailureSeverity string `yaml:"failure_severity,omitempty" json:"failure_severity,omitempty"`
}

type WorkflowSpec struct {
	ValidationRequired bool `yaml:"validation_required,omitempty" json:"validation_required,omitempty"`
	ApprovalRequired   bool `yaml:"approval_required,omitempty" json:"approval_required,omitempty"`
}

type SyncSpec struct {
	LocalRevision  string    `yaml:"local_revision,omitempty" json:"local_revision,omitempty"`
	RemoteRevision string    `yaml:"remote_revision,omitempty" json:"remote_revision,omitempty"`
	Precedence     string    `yaml:"precedence,omitempty" json:"precedence,omitempty"`
	LastSyncedAt   time.Time `yaml:"last_synced_at,omitempty" json:"last_synced_at,omitempty"`
}

type ProductionConfigSpec struct {
	TenantID         string       `yaml:"tenant_id" json:"tenant_id"`
	Environment      string       `yaml:"environment" json:"environment"`
	Repository       string       `yaml:"repository,omitempty" json:"repository,omitempty"`
	APIURL           string       `yaml:"api_url,omitempty" json:"api_url,omitempty"`
	PolicyBundleDir  string       `yaml:"policy_bundle_dir" json:"policy_bundle_dir"`
	KyvernoPolicyDir string       `yaml:"kyverno_policy_dir" json:"kyverno_policy_dir"`
	CLI              CLISpec      `yaml:"cli,omitempty" json:"cli,omitempty"`
	Sync             SyncSpec     `yaml:"sync,omitempty" json:"sync,omitempty"`
	Workflow         WorkflowSpec `yaml:"workflow,omitempty" json:"workflow,omitempty"`
}

type ProductionConfig struct {
	APIVersion string               `yaml:"apiVersion" json:"apiVersion"`
	Kind       string               `yaml:"kind" json:"kind"`
	Metadata   Metadata             `yaml:"metadata" json:"metadata"`
	Spec       ProductionConfigSpec `yaml:"spec" json:"spec"`
}

type ValidationIssue struct {
	Severity string `json:"severity"`
	Code     string `json:"code"`
	Field    string `json:"field,omitempty"`
	Message  string `json:"message"`
}

type ValidationReport struct {
	SchemaVersion string            `json:"schema_version"`
	CurrentState  string            `json:"current_state"`
	SyncState     string            `json:"sync_state"`
	ResolvedPaths map[string]string `json:"resolved_paths,omitempty"`
	Issues        []ValidationIssue `json:"issues,omitempty"`
	Limitations   []string          `json:"limitations,omitempty"`
}

func LoadProductionConfig(path string) (ProductionConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ProductionConfig{}, fmt.Errorf("read production config: %w", err)
	}
	var cfg ProductionConfig
	decoder := yaml.NewDecoder(strings.NewReader(string(data)))
	decoder.KnownFields(true)
	if err := decoder.Decode(&cfg); err != nil {
		return ProductionConfig{}, fmt.Errorf("decode production config: %w", err)
	}
	return cfg, nil
}

func ValidateProductionConfig(path string, cfg ProductionConfig, now func() time.Time) ValidationReport {
	if now == nil {
		now = time.Now
	}
	baseDir := filepath.Dir(path)
	issues := []ValidationIssue{}
	resolvedPaths := map[string]string{
		"policy_bundle_dir":  resolvePath(baseDir, cfg.Spec.PolicyBundleDir),
		"kyverno_policy_dir": resolvePath(baseDir, cfg.Spec.KyvernoPolicyDir),
	}

	if strings.TrimSpace(cfg.APIVersion) != ProductionConfigAPIVersion {
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "unsupported_api_version",
			Field:    "apiVersion",
			Message:  fmt.Sprintf("apiVersion must be %q", ProductionConfigAPIVersion),
		})
	}
	if strings.TrimSpace(cfg.Kind) != ProductionConfigKind {
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "unsupported_kind",
			Field:    "kind",
			Message:  fmt.Sprintf("kind must be %q", ProductionConfigKind),
		})
	}
	if strings.TrimSpace(cfg.Metadata.Name) == "" {
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "metadata_name_required",
			Field:    "metadata.name",
			Message:  "metadata.name is required",
		})
	}
	if strings.TrimSpace(cfg.Spec.TenantID) == "" {
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "tenant_required",
			Field:    "spec.tenant_id",
			Message:  "spec.tenant_id is required",
		})
	}
	if strings.TrimSpace(cfg.Spec.Environment) == "" {
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "environment_required",
			Field:    "spec.environment",
			Message:  "spec.environment is required",
		})
	}
	if strings.TrimSpace(cfg.Spec.PolicyBundleDir) == "" {
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "policy_bundle_dir_required",
			Field:    "spec.policy_bundle_dir",
			Message:  "spec.policy_bundle_dir is required",
		})
	}
	if strings.TrimSpace(cfg.Spec.KyvernoPolicyDir) == "" {
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "kyverno_policy_dir_required",
			Field:    "spec.kyverno_policy_dir",
			Message:  "spec.kyverno_policy_dir is required",
		})
	}
	if apiURL := strings.TrimSpace(cfg.Spec.APIURL); apiURL != "" {
		parsed, err := url.Parse(apiURL)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			issues = append(issues, ValidationIssue{
				Severity: IssueSeverityError,
				Code:     "api_url_invalid",
				Field:    "spec.api_url",
				Message:  "spec.api_url must be an absolute http or https URL",
			})
		} else if parsed.Scheme != "http" && parsed.Scheme != "https" {
			issues = append(issues, ValidationIssue{
				Severity: IssueSeverityError,
				Code:     "api_url_scheme_invalid",
				Field:    "spec.api_url",
				Message:  "spec.api_url must use http or https",
			})
		}
	}
	switch strings.ToLower(strings.TrimSpace(cfg.Spec.Sync.Precedence)) {
	case "", "local", "remote":
	default:
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "sync_precedence_invalid",
			Field:    "spec.sync.precedence",
			Message:  "spec.sync.precedence must be local or remote when set",
		})
	}
	if cfg.Spec.CLI.Offline && strings.TrimSpace(cfg.Spec.APIURL) != "" {
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityWarning,
			Code:     "api_url_ignored_offline",
			Field:    "spec.cli.offline",
			Message:  "spec.cli.offline=true keeps remote API URL configured but disables API-assisted preview and check paths",
		})
	}
	validatePath := func(field, key string) {
		resolved := resolvedPaths[key]
		info, err := os.Stat(resolved)
		if err != nil {
			issues = append(issues, ValidationIssue{
				Severity: IssueSeverityError,
				Code:     "path_missing",
				Field:    field,
				Message:  fmt.Sprintf("%s does not exist at %s", field, resolved),
			})
			return
		}
		if !info.IsDir() {
			issues = append(issues, ValidationIssue{
				Severity: IssueSeverityError,
				Code:     "path_not_directory",
				Field:    field,
				Message:  fmt.Sprintf("%s must point to a directory", field),
			})
		}
	}
	if strings.TrimSpace(cfg.Spec.PolicyBundleDir) != "" {
		validatePath("spec.policy_bundle_dir", "policy_bundle_dir")
	}
	if strings.TrimSpace(cfg.Spec.KyvernoPolicyDir) != "" {
		validatePath("spec.kyverno_policy_dir", "kyverno_policy_dir")
	}

	switch strings.ToLower(strings.TrimSpace(cfg.Spec.CLI.Output)) {
	case "", "human", "json":
	default:
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "cli_output_invalid",
			Field:    "spec.cli.output",
			Message:  "spec.cli.output must be human or json",
		})
	}
	switch strings.ToLower(strings.TrimSpace(cfg.Spec.CLI.Scanner)) {
	case "", "auto", "trivy", "grype":
	default:
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "scanner_invalid",
			Field:    "spec.cli.scanner",
			Message:  "spec.cli.scanner must be auto, trivy, or grype",
		})
	}
	switch normalizeSeverity(cfg.Spec.CLI.FailureSeverity) {
	case "", "CRITICAL", "HIGH", "MEDIUM", "LOW":
	default:
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "failure_severity_invalid",
			Field:    "spec.cli.failure_severity",
			Message:  "spec.cli.failure_severity must be CRITICAL, HIGH, MEDIUM, or LOW",
		})
	}

	syncState := evaluateSyncState(cfg.Spec.Sync, now())
	switch syncState {
	case SyncStateConflict:
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityError,
			Code:     "sync_conflict_visible",
			Field:    "spec.sync",
			Message:  "local and remote revisions diverge; resolve precedence before production use",
		})
	case SyncStateStale:
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityWarning,
			Code:     "sync_state_stale",
			Field:    "spec.sync.last_synced_at",
			Message:  "config sync state is stale and should be refreshed before rollout",
		})
	case SyncStateLocalOnly, SyncStateRemoteOnly:
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityWarning,
			Code:     "sync_state_partial",
			Field:    "spec.sync",
			Message:  "only one side of sync revision is present; conflict detection is bounded",
		})
	}
	if cfg.Spec.Workflow.ValidationRequired && !cfg.Spec.Workflow.ApprovalRequired {
		issues = append(issues, ValidationIssue{
			Severity: IssueSeverityWarning,
			Code:     "workflow_governance_thin",
			Field:    "spec.workflow",
			Message:  "validation is required but approval_required is disabled; confirm this is intentional for production closure paths",
		})
	}

	currentState := "valid"
	for _, issue := range issues {
		if issue.Severity == IssueSeverityError {
			currentState = "invalid"
			break
		}
	}
	if currentState == "valid" && len(issues) > 0 {
		currentState = "valid_with_warnings"
	}

	return ValidationReport{
		SchemaVersion: ReportSchemaVersion,
		CurrentState:  currentState,
		SyncState:     syncState,
		ResolvedPaths: resolvedPaths,
		Issues:        issues,
		Limitations: []string{
			"Production config validation is schema-strict and fail-fast for declared fields, but it only evaluates the file and local path state in scope.",
			"Sync-state checks surface visible divergence and staleness; they do not silently auto-resolve precedence.",
		},
	}
}

func resolvePath(baseDir, value string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return ""
	}
	if filepath.IsAbs(trimmed) {
		return filepath.Clean(trimmed)
	}
	return filepath.Clean(filepath.Join(baseDir, trimmed))
}

func normalizeSeverity(value string) string {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "CRITICAL":
		return "CRITICAL"
	case "HIGH":
		return "HIGH"
	case "MEDIUM":
		return "MEDIUM"
	case "LOW":
		return "LOW"
	default:
		return strings.ToUpper(strings.TrimSpace(value))
	}
}

func evaluateSyncState(sync SyncSpec, now time.Time) string {
	local := strings.TrimSpace(sync.LocalRevision)
	remote := strings.TrimSpace(sync.RemoteRevision)
	precedence := strings.ToLower(strings.TrimSpace(sync.Precedence))

	switch {
	case local == "" && remote == "":
		return SyncStateUnconfigured
	case local != "" && remote == "":
		return SyncStateLocalOnly
	case local == "" && remote != "":
		return SyncStateRemoteOnly
	case local != remote:
		if precedence != "local" && precedence != "remote" {
			return SyncStateConflict
		}
		return SyncStateConflict
	}

	if !sync.LastSyncedAt.IsZero() && now.Sub(sync.LastSyncedAt.UTC()) > 24*time.Hour {
		return SyncStateStale
	}
	return SyncStateInSync
}

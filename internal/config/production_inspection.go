package config

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

const InspectionSchemaVersion = "5.production_config_inspection.v1"

type DefaultApplication struct {
	Field  string `json:"field"`
	Value  string `json:"value"`
	Source string `json:"source"`
}

type EffectiveCLISpec struct {
	Output          string `json:"output"`
	Offline         bool   `json:"offline"`
	Scanner         string `json:"scanner"`
	FailureSeverity string `json:"failure_severity"`
}

type EffectiveWorkflowSpec struct {
	ValidationRequired bool `json:"validation_required"`
	ApprovalRequired   bool `json:"approval_required"`
}

type SyncInspection struct {
	CurrentState         string    `json:"current_state"`
	CurrentRevision      string    `json:"current_revision,omitempty"`
	DesiredRevision      string    `json:"desired_revision,omitempty"`
	SourceRevision       string    `json:"source_revision,omitempty"`
	Precedence           string    `json:"precedence"`
	LastSyncedAt         time.Time `json:"last_synced_at,omitempty"`
	LocalOverrideVisible bool      `json:"local_override_visible"`
	ConflictVisible      bool      `json:"conflict_visible"`
	StaleVisible         bool      `json:"stale_visible"`
	Explanation          []string  `json:"explanation,omitempty"`
}

type EffectiveProductionConfig struct {
	Name             string                `json:"name"`
	TenantID         string                `json:"tenant_id"`
	Environment      string                `json:"environment"`
	Repository       string                `json:"repository,omitempty"`
	APIURL           string                `json:"api_url,omitempty"`
	PolicyBundleDir  string                `json:"policy_bundle_dir"`
	KyvernoPolicyDir string                `json:"kyverno_policy_dir"`
	CLI              EffectiveCLISpec      `json:"cli"`
	Workflow         EffectiveWorkflowSpec `json:"workflow"`
	Sync             SyncInspection        `json:"sync"`
}

type InspectionReport struct {
	SchemaVersion   string                    `json:"schema_version"`
	CurrentState    string                    `json:"current_state"`
	ConfigPath      string                    `json:"config_path"`
	DeclaredConfig  ProductionConfigSpec      `json:"declared_config"`
	EffectiveConfig EffectiveProductionConfig `json:"effective_config"`
	Validation      ValidationReport          `json:"validation"`
	DefaultsApplied []DefaultApplication      `json:"defaults_applied,omitempty"`
	ReasonCodes     []string                  `json:"reason_codes,omitempty"`
	Limitations     []string                  `json:"limitations,omitempty"`
}

func InspectProductionConfig(path string, cfg ProductionConfig, now func() time.Time) InspectionReport {
	if now == nil {
		now = time.Now
	}
	validation := ValidateProductionConfig(path, cfg, now)
	defaults := []DefaultApplication{}

	output := strings.ToLower(strings.TrimSpace(cfg.Spec.CLI.Output))
	if output == "" {
		output = "human"
		defaults = append(defaults, DefaultApplication{Field: "spec.cli.output", Value: output, Source: "default"})
	}
	scanner := strings.ToLower(strings.TrimSpace(cfg.Spec.CLI.Scanner))
	if scanner == "" {
		scanner = "auto"
		defaults = append(defaults, DefaultApplication{Field: "spec.cli.scanner", Value: scanner, Source: "default"})
	}
	failureSeverity := normalizeSeverity(cfg.Spec.CLI.FailureSeverity)
	if failureSeverity == "" {
		failureSeverity = "CRITICAL"
		defaults = append(defaults, DefaultApplication{Field: "spec.cli.failure_severity", Value: failureSeverity, Source: "default"})
	}

	absolutePath := filepath.Clean(path)
	if abs, err := filepath.Abs(path); err == nil {
		absolutePath = abs
	}

	syncInspection := buildSyncInspection(cfg.Spec.Sync, now())
	reasonCodes := make([]string, 0, len(validation.Issues)+2)
	for _, issue := range validation.Issues {
		reasonCodes = append(reasonCodes, issue.Code)
	}
	if syncInspection.CurrentState != "" && syncInspection.CurrentState != SyncStateInSync {
		reasonCodes = append(reasonCodes, syncInspection.CurrentState)
	}
	if len(defaults) > 0 {
		reasonCodes = append(reasonCodes, "defaults_applied_visible")
	}

	currentState := validation.CurrentState
	if currentState == "valid" && len(defaults) > 0 {
		currentState = "valid_with_defaults"
	}

	return InspectionReport{
		SchemaVersion:  InspectionSchemaVersion,
		CurrentState:   currentState,
		ConfigPath:     absolutePath,
		DeclaredConfig: cfg.Spec,
		EffectiveConfig: EffectiveProductionConfig{
			Name:             strings.TrimSpace(cfg.Metadata.Name),
			TenantID:         strings.TrimSpace(cfg.Spec.TenantID),
			Environment:      strings.TrimSpace(cfg.Spec.Environment),
			Repository:       strings.TrimSpace(cfg.Spec.Repository),
			APIURL:           strings.TrimSpace(cfg.Spec.APIURL),
			PolicyBundleDir:  resolvePath(filepath.Dir(path), cfg.Spec.PolicyBundleDir),
			KyvernoPolicyDir: resolvePath(filepath.Dir(path), cfg.Spec.KyvernoPolicyDir),
			CLI: EffectiveCLISpec{
				Output:          output,
				Offline:         cfg.Spec.CLI.Offline,
				Scanner:         scanner,
				FailureSeverity: failureSeverity,
			},
			Workflow: EffectiveWorkflowSpec{
				ValidationRequired: cfg.Spec.Workflow.ValidationRequired,
				ApprovalRequired:   cfg.Spec.Workflow.ApprovalRequired,
			},
			Sync: syncInspection,
		},
		Validation:      validation,
		DefaultsApplied: defaults,
		ReasonCodes:     uniqueStrings(reasonCodes),
		Limitations: []string{
			"Effective config inspection shows declared and normalized local production config only; it does not claim live remote runtime state beyond the sync revision data in scope.",
			"Sync precedence explanation keeps divergence visible instead of silently auto-resolving conflicting revisions.",
		},
	}
}

func buildSyncInspection(sync SyncSpec, now time.Time) SyncInspection {
	local := strings.TrimSpace(sync.LocalRevision)
	remote := strings.TrimSpace(sync.RemoteRevision)
	precedence := strings.ToLower(strings.TrimSpace(sync.Precedence))
	state := evaluateSyncState(sync, now)
	lastSyncedAt := sync.LastSyncedAt.UTC()
	if sync.LastSyncedAt.IsZero() {
		lastSyncedAt = time.Time{}
	}

	sourceRevision := ""
	switch precedence {
	case "local":
		sourceRevision = local
	case "remote":
		sourceRevision = remote
	default:
		sourceRevision = firstNonEmpty(remote, local)
	}

	explanation := []string{}
	switch {
	case local == "" && remote == "":
		explanation = append(explanation, "No local or remote revision is declared, so sync state is not configured.")
	case local != "" && remote == "":
		explanation = append(explanation, fmt.Sprintf("Only local revision %q is declared, so remote convergence cannot be confirmed.", local))
	case local == "" && remote != "":
		explanation = append(explanation, fmt.Sprintf("Only remote revision %q is declared, so local effective state is bounded.", remote))
	case local != remote:
		explanation = append(explanation, fmt.Sprintf("Local revision %q diverges from remote revision %q.", local, remote))
		if precedence == "" {
			explanation = append(explanation, "No explicit precedence is declared, so ChangeLock keeps the divergence visible and does not auto-resolve it.")
		} else {
			explanation = append(explanation, fmt.Sprintf("Precedence is explicitly set to %q, but divergence remains operator-visible until revisions converge.", precedence))
		}
	default:
		explanation = append(explanation, fmt.Sprintf("Local and remote revision both resolve to %q.", local))
	}
	if !lastSyncedAt.IsZero() && now.Sub(lastSyncedAt) > 24*time.Hour {
		explanation = append(explanation, "The last sync timestamp is older than 24 hours, so the revision projection is stale.")
	}

	return SyncInspection{
		CurrentState:         state,
		CurrentRevision:      local,
		DesiredRevision:      remote,
		SourceRevision:       sourceRevision,
		Precedence:           firstNonEmpty(precedence, "unspecified"),
		LastSyncedAt:         lastSyncedAt,
		LocalOverrideVisible: precedence == "local" && local != "" && remote != "" && local != remote,
		ConflictVisible:      state == SyncStateConflict,
		StaleVisible:         state == SyncStateStale,
		Explanation:          uniqueStrings(explanation),
	}
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]struct{}{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if trimmed := strings.TrimSpace(value); trimmed != "" {
			return trimmed
		}
	}
	return ""
}

package preflightcli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	prodcfg "github.com/denisgrosek/changelock/internal/config"
	runtimecfg "github.com/denisgrosek/changelock/internal/runtime"
)

const (
	inspectSchemaVersion = "5b.cli_inspect.v1"
	explainSchemaVersion = "5b.cli_explain.v1"
)

type valBOptions struct {
	ConfigPath              string
	Output                  string
	TrustedExecutionProfile string
	SubjectRef              string
	Topic                   string
}

type inspectResponse struct {
	SchemaVersion     string                           `json:"schema_version"`
	CurrentState      string                           `json:"current_state"`
	ConfigPath        string                           `json:"config_path"`
	ConfigInspection  prodcfg.InspectionReport         `json:"config_inspection,omitempty"`
	RuntimeInspection runtimecfg.SelfHealingInspection `json:"runtime_inspection,omitempty"`
	ReasonCodes       []string                         `json:"reason_codes,omitempty"`
	Limitations       []string                         `json:"limitations,omitempty"`
}

type explainResponse struct {
	SchemaVersion  string         `json:"schema_version"`
	CurrentState   string         `json:"current_state"`
	Topic          string         `json:"topic"`
	ReasonCodes    []string       `json:"reason_codes,omitempty"`
	Explanations   []string       `json:"explanations,omitempty"`
	EffectiveState map[string]any `json:"effective_state,omitempty"`
	Limitations    []string       `json:"limitations,omitempty"`
}

type productionContext struct {
	Path              string
	Config            prodcfg.ProductionConfig
	Inspection        prodcfg.InspectionReport
	RuntimeInspection runtimecfg.SelfHealingInspection
}

func (a *App) runCheck(ctx context.Context, args []string) (Result, error) {
	_ = ctx
	options, err := a.parseValBOptions("check", args)
	if err != nil {
		return Result{}, err
	}
	a.config.Output = options.Output
	contextReport, loadErr := a.loadProductionContext(options.ConfigPath)
	if loadErr != nil {
		return configLoadFailureResult("check", options.ConfigPath, loadErr), nil
	}

	checks := []CheckResult{
		{
			Name:    "config-schema",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  statusForValidationState(contextReport.Inspection.Validation.CurrentState),
			Summary: configValidationSummary(contextReport.Inspection.Validation.CurrentState),
			Details: validationMessages(contextReport.Inspection.Validation.Issues),
			Metadata: map[string]any{
				"sync_state": contextReport.Inspection.Validation.SyncState,
			},
		},
		{
			Name:    "config-compatibility",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  statusForCompatibility(contextReport.Inspection.Validation.Issues),
			Summary: compatibilitySummary(contextReport.Inspection.Validation.Issues),
			Details: validationMessages(contextReport.Inspection.Validation.Issues),
		},
		{
			Name:    "sync-state",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  statusForSyncState(contextReport.Inspection.EffectiveConfig.Sync.CurrentState),
			Summary: syncSummary(contextReport.Inspection.EffectiveConfig.Sync.CurrentState),
			Details: contextReport.Inspection.EffectiveConfig.Sync.Explanation,
		},
		{
			Name:    "workflow-closure",
			Mode:    ModeLocal,
			Target:  firstNonEmpty(options.SubjectRef, contextReport.Config.Metadata.Name),
			Status:  statusForWorkflow(contextReport.Config.Spec.Workflow),
			Summary: workflowSummary(contextReport.Config.Spec.Workflow),
			Details: workflowDetails(contextReport.Config.Spec.Workflow),
		},
		{
			Name:    "runtime-self-healing",
			Mode:    ModeLocal,
			Target:  "runtime-self-healing",
			Status:  statusForInspectionState(contextReport.RuntimeInspection.CurrentState),
			Summary: runtimeInspectionSummary(contextReport.RuntimeInspection),
			Details: runtimeInspectionDetails(contextReport.RuntimeInspection),
		},
	}
	checks = append(checks, trustedExecutionCheck(options.TrustedExecutionProfile))

	result := finalizeResult(Result{
		Command: "check",
		Mode:    ModeLocalOnly,
		Inputs: map[string]string{
			"config":                    contextReport.Path,
			"trusted_execution_profile": options.TrustedExecutionProfile,
			"subject_ref":               options.SubjectRef,
		},
		Checks:      checks,
		ReasonCodes: uniqueStrings(append([]string{}, contextReport.Inspection.ReasonCodes...)),
		Limitations: []string{"Check evaluates local production config, runtime env config, and optional trusted-execution profile ID validity; it does not claim live runtime subject verification."},
	})
	return result, nil
}

func (a *App) runPreview(ctx context.Context, args []string) (Result, error) {
	_ = ctx
	options, err := a.parseValBOptions("preview", args)
	if err != nil {
		return Result{}, err
	}
	a.config.Output = options.Output
	contextReport, loadErr := a.loadProductionContext(options.ConfigPath)
	if loadErr != nil {
		return configLoadFailureResult("preview", options.ConfigPath, loadErr), nil
	}

	checks := []CheckResult{
		{
			Name:    "startup-preview",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  statusForValidationState(contextReport.Inspection.Validation.CurrentState),
			Summary: previewStartupSummary(contextReport.Inspection.Validation.CurrentState),
			Details: validationMessages(contextReport.Inspection.Validation.Issues),
		},
		{
			Name:    "sync-preview",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  statusForSyncState(contextReport.Inspection.EffectiveConfig.Sync.CurrentState),
			Summary: previewSyncSummary(contextReport.Inspection.EffectiveConfig.Sync.CurrentState),
			Details: contextReport.Inspection.EffectiveConfig.Sync.Explanation,
		},
		{
			Name:    "workflow-closure-preview",
			Mode:    ModeLocal,
			Target:  firstNonEmpty(options.SubjectRef, contextReport.Config.Metadata.Name),
			Status:  statusForWorkflowPreview(contextReport.Config.Spec.Workflow),
			Summary: previewWorkflowSummary(contextReport.Config.Spec.Workflow),
			Details: workflowDetails(contextReport.Config.Spec.Workflow),
		},
		trustedExecutionPreview(options.TrustedExecutionProfile),
		{
			Name:    "preview-honesty",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  StatusInfo,
			Summary: "preview is bounded to deterministic local config and declared sync state",
			Details: []string{
				"Preview does not include live runtime attestation freshness, enterprise workflow evidence, or Phase 3 intelligence beyond the config and profile data in scope.",
				"Trusted execution preview is limited to deterministic profile existence unless a live substrate truth record is provided elsewhere.",
			},
		},
	}

	result := finalizeResult(Result{
		Command: "preview",
		Mode:    ModeLocalOnly,
		Inputs: map[string]string{
			"config":                    contextReport.Path,
			"trusted_execution_profile": options.TrustedExecutionProfile,
			"subject_ref":               options.SubjectRef,
		},
		Checks: checks,
		ReasonCodes: uniqueStrings(append([]string{
			"preview_bounded_local_only",
		}, contextReport.Inspection.ReasonCodes...)),
		Limitations: []string{
			"Preview is honest about bounded scope and does not promise live runtime success from local config alone.",
			"Conflict, stale, and override visibility come from declared sync revision data and are not silently auto-resolved.",
		},
	})
	return result, nil
}

func (a *App) runInspect(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	_ = ctx
	options, err := a.parseValBOptions("inspect", args)
	if err != nil {
		return a.writeError(stderr, err)
	}
	a.config.Output = options.Output
	contextReport, loadErr := a.loadProductionContext(options.ConfigPath)
	response := inspectResponse{
		SchemaVersion: inspectSchemaVersion,
		ConfigPath:    options.ConfigPath,
		Limitations: []string{
			"Inspect shows declared, normalized, and effective local config state; it does not replace authoritative server-side truth.",
			"Runtime self-healing inspection reflects local env config and defaults, not live remediation execution history.",
		},
	}
	if loadErr != nil {
		response.CurrentState = "invalid"
		response.ReasonCodes = []string{"config_load_failed"}
		if err := renderInspectResponse(stdout, a.config.Output, response, []string{loadErr.Error()}); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitFailed
	}

	response.CurrentState = combinedInspectionState(contextReport.Inspection.CurrentState, contextReport.RuntimeInspection.CurrentState)
	response.ConfigPath = contextReport.Path
	response.ConfigInspection = contextReport.Inspection
	response.RuntimeInspection = contextReport.RuntimeInspection
	response.ReasonCodes = uniqueStrings(append(contextReport.Inspection.ReasonCodes, runtimeIssueCodes(contextReport.RuntimeInspection)...))

	if err := renderInspectResponse(stdout, a.config.Output, response, nil); err != nil {
		return a.writeError(stderr, err)
	}
	return exitCodeForInspection(response.CurrentState)
}

func (a *App) runExplain(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	_ = ctx
	options, err := a.parseValBOptions("explain", args)
	if err != nil {
		return a.writeError(stderr, err)
	}
	a.config.Output = options.Output
	contextReport, loadErr := a.loadProductionContext(options.ConfigPath)
	response := explainResponse{
		SchemaVersion: explainSchemaVersion,
		Topic:         firstNonEmpty(options.Topic, "config"),
		Limitations: []string{
			"Explain is bounded to local config, sync revision, runtime env config, and optional trusted-execution profile identity in scope.",
			"It does not invent runtime-only or server-only evidence that is not currently available.",
		},
	}
	if loadErr != nil {
		response.CurrentState = "invalid"
		response.ReasonCodes = []string{"config_load_failed"}
		response.Explanations = []string{loadErr.Error()}
		if err := renderExplainResponse(stdout, a.config.Output, response); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitFailed
	}

	buildExplainResponse(&response, contextReport, options)
	if err := renderExplainResponse(stdout, a.config.Output, response); err != nil {
		return a.writeError(stderr, err)
	}
	return exitCodeForInspection(response.CurrentState)
}

func (a *App) parseValBOptions(command string, args []string) (valBOptions, error) {
	fs := newFlagSet(command)
	options := valBOptions{
		ConfigPath: a.config.ProductionConfigPath,
		Output:     a.config.Output,
		Topic:      "config",
	}
	fs.StringVar(&options.ConfigPath, "config", options.ConfigPath, "production config file")
	fs.StringVar(&options.Output, "output", options.Output, "output mode: human|json")
	fs.StringVar(&options.TrustedExecutionProfile, "trusted-execution-profile", "", "trusted execution profile to check or preview")
	fs.StringVar(&options.SubjectRef, "subject-ref", "", "optional subject reference for bounded workflow wording")
	fs.StringVar(&options.Topic, "topic", options.Topic, "explain topic: config|sync|workflow|trusted-execution")
	if err := fs.Parse(args); err != nil {
		return valBOptions{}, usageError{message: err.Error()}
	}
	if strings.TrimSpace(options.ConfigPath) == "" {
		return valBOptions{}, usageError{message: fmt.Sprintf("%s requires --config or CHANGELOCK_CLI_CONFIG", command)}
	}
	switch strings.ToLower(strings.TrimSpace(options.Output)) {
	case "human", "json", "":
	default:
		return valBOptions{}, usageError{message: fmt.Sprintf("unsupported output mode %q", options.Output)}
	}
	switch strings.ToLower(strings.TrimSpace(options.Topic)) {
	case "config", "sync", "workflow", "trusted-execution":
	default:
		return valBOptions{}, usageError{message: fmt.Sprintf("unsupported explain topic %q", options.Topic)}
	}
	return options, nil
}

func (a *App) loadProductionContext(path string) (productionContext, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return productionContext{}, err
	}
	cfg, err := prodcfg.LoadProductionConfig(absPath)
	if err != nil {
		return productionContext{}, err
	}
	return productionContext{
		Path:              absPath,
		Config:            cfg,
		Inspection:        prodcfg.InspectProductionConfig(absPath, cfg, time.Now),
		RuntimeInspection: runtimecfg.InspectSelfHealingConfig(os.Getenv),
	}, nil
}

func configLoadFailureResult(command, path string, err error) Result {
	return finalizeResult(Result{
		Command: command,
		Mode:    ModeLocalOnly,
		Inputs: map[string]string{
			"config": strings.TrimSpace(path),
		},
		Checks: []CheckResult{{
			Name:    "config-load",
			Mode:    ModeLocal,
			Target:  strings.TrimSpace(path),
			Status:  StatusFail,
			Summary: "production config failed strict loading",
			Details: []string{err.Error()},
		}},
		ReasonCodes: []string{"config_load_failed"},
		Limitations: []string{"Strict config loading is fail-fast and intentionally does not fall back to a permissive parse path."},
	})
}

func trustedExecutionCheck(profileID string) CheckResult {
	profileID = strings.TrimSpace(profileID)
	if profileID == "" {
		return CheckResult{
			Name:    "trusted-execution-profile",
			Mode:    ModeLocal,
			Status:  StatusInfo,
			Summary: "no trusted execution profile was requested for deterministic validation",
		}
	}
	profile, ok := runtimecfg.ExecutionProfileByID(profileID)
	if !ok {
		return CheckResult{
			Name:    "trusted-execution-profile",
			Mode:    ModeLocal,
			Target:  profileID,
			Status:  StatusFail,
			Summary: "trusted execution profile is unknown",
		}
	}
	return CheckResult{
		Name:    "trusted-execution-profile",
		Mode:    ModeLocal,
		Target:  profileID,
		Status:  StatusPass,
		Summary: "trusted execution profile exists and is locally explainable",
		Details: []string{
			fmt.Sprintf("Trust boundary: %s", profile.TrustBoundary),
			fmt.Sprintf("Minimum confidence: %d", profile.MinimumConfidence),
		},
	}
}

func trustedExecutionPreview(profileID string) CheckResult {
	profileID = strings.TrimSpace(profileID)
	if profileID == "" {
		return CheckResult{
			Name:    "trusted-execution-preview",
			Mode:    ModeLocal,
			Status:  StatusInfo,
			Summary: "no trusted execution profile was requested for preview",
		}
	}
	if _, ok := runtimecfg.ExecutionProfileByID(profileID); !ok {
		return CheckResult{
			Name:    "trusted-execution-preview",
			Mode:    ModeLocal,
			Target:  profileID,
			Status:  StatusFail,
			Summary: "preview cannot continue because the trusted execution profile is unknown",
		}
	}
	return CheckResult{
		Name:    "trusted-execution-preview",
		Mode:    ModeLocal,
		Target:  profileID,
		Status:  StatusInfo,
		Summary: "trusted execution preview confirms profile identity only; live substrate matching remains bounded",
		Details: []string{
			"Runtime-only substrate truth, attestation freshness, and credential-release state are not part of this local preview.",
		},
	}
}

func buildExplainResponse(response *explainResponse, contextReport productionContext, options valBOptions) {
	response.CurrentState = contextReport.Inspection.CurrentState
	topic := strings.ToLower(strings.TrimSpace(options.Topic))
	switch topic {
	case "sync":
		response.ReasonCodes = uniqueStrings(append(issueCodesForField(contextReport.Inspection.Validation.Issues, "spec.sync"), contextReport.Inspection.EffectiveConfig.Sync.CurrentState))
		response.Explanations = append([]string{}, contextReport.Inspection.EffectiveConfig.Sync.Explanation...)
		response.EffectiveState = map[string]any{
			"current_state":          contextReport.Inspection.EffectiveConfig.Sync.CurrentState,
			"current_revision":       contextReport.Inspection.EffectiveConfig.Sync.CurrentRevision,
			"desired_revision":       contextReport.Inspection.EffectiveConfig.Sync.DesiredRevision,
			"source_revision":        contextReport.Inspection.EffectiveConfig.Sync.SourceRevision,
			"precedence":             contextReport.Inspection.EffectiveConfig.Sync.Precedence,
			"local_override_visible": contextReport.Inspection.EffectiveConfig.Sync.LocalOverrideVisible,
			"conflict_visible":       contextReport.Inspection.EffectiveConfig.Sync.ConflictVisible,
		}
	case "workflow":
		response.ReasonCodes = uniqueStrings(append(issueCodesForField(contextReport.Inspection.Validation.Issues, "spec.workflow"), workflowReasonCodes(contextReport.Config.Spec.Workflow)...))
		response.Explanations = workflowDetails(contextReport.Config.Spec.Workflow)
		response.EffectiveState = map[string]any{
			"validation_required": contextReport.Config.Spec.Workflow.ValidationRequired,
			"approval_required":   contextReport.Config.Spec.Workflow.ApprovalRequired,
		}
	case "trusted-execution":
		profileID := strings.TrimSpace(options.TrustedExecutionProfile)
		if profileID == "" {
			response.ReasonCodes = []string{"trusted_execution_profile_not_requested"}
			response.Explanations = []string{"No trusted execution profile was requested, so explain can only state that no deterministic profile check was performed."}
			response.EffectiveState = map[string]any{"profile_requested": false}
			return
		}
		profile, ok := runtimecfg.ExecutionProfileByID(profileID)
		if !ok {
			response.CurrentState = "invalid"
			response.ReasonCodes = []string{"trusted_execution_profile_unknown"}
			response.Explanations = []string{"The requested trusted execution profile is not registered locally, so any preview or check would fail before runtime admission."}
			response.EffectiveState = map[string]any{"profile_id": profileID}
			return
		}
		response.ReasonCodes = []string{"trusted_execution_profile_known", "runtime_subject_truth_required"}
		response.Explanations = []string{
			fmt.Sprintf("Trusted execution profile %q exists locally and requires trust boundary %q.", profile.ProfileID, profile.TrustBoundary),
			"Profile explanation remains bounded until a live substrate truth record is supplied; this CLI path does not synthesize runtime-only evidence.",
		}
		response.EffectiveState = map[string]any{
			"profile_id":                 profile.ProfileID,
			"trust_boundary":             profile.TrustBoundary,
			"minimum_confidence":         profile.MinimumConfidence,
			"require_attestation":        profile.RequireAttestation,
			"require_credential_release": profile.RequireCredentialRelease,
		}
	default:
		response.ReasonCodes = uniqueStrings(contextReport.Inspection.ReasonCodes)
		response.Explanations = validationMessages(contextReport.Inspection.Validation.Issues)
		if len(response.Explanations) == 0 {
			response.Explanations = []string{"The production config is valid with no blocking schema or compatibility findings in the current local scope."}
		}
		response.EffectiveState = map[string]any{
			"current_state":     contextReport.Inspection.CurrentState,
			"sync_state":        contextReport.Inspection.Validation.SyncState,
			"effective_output":  contextReport.Inspection.EffectiveConfig.CLI.Output,
			"effective_scanner": contextReport.Inspection.EffectiveConfig.CLI.Scanner,
			"defaults_applied":  contextReport.Inspection.DefaultsApplied,
		}
	}
}

func renderInspectResponse(w io.Writer, output string, response inspectResponse, loadErrors []string) error {
	switch strings.ToLower(strings.TrimSpace(output)) {
	case "", "human":
		if _, err := fmt.Fprintf(w, "Config: %s\nState: %s\n", response.ConfigPath, response.CurrentState); err != nil {
			return err
		}
		if len(loadErrors) > 0 {
			for _, item := range loadErrors {
				if _, err := fmt.Fprintf(w, "- %s\n", item); err != nil {
					return err
				}
			}
			return nil
		}
		cfg := response.ConfigInspection.EffectiveConfig
		if _, err := fmt.Fprintf(w, "\nEffective config:\n- tenant_id: %s\n- environment: %s\n- repository: %s\n- policy_bundle_dir: %s\n- kyverno_policy_dir: %s\n- cli.output: %s\n- cli.offline: %t\n- cli.scanner: %s\n- cli.failure_severity: %s\n- workflow.validation_required: %t\n- workflow.approval_required: %t\n", cfg.TenantID, cfg.Environment, firstNonEmpty(cfg.Repository, "-"), cfg.PolicyBundleDir, cfg.KyvernoPolicyDir, cfg.CLI.Output, cfg.CLI.Offline, cfg.CLI.Scanner, cfg.CLI.FailureSeverity, cfg.Workflow.ValidationRequired, cfg.Workflow.ApprovalRequired); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(w, "\nSync state:\n- current_state: %s\n- current_revision: %s\n- desired_revision: %s\n- source_revision: %s\n- precedence: %s\n", cfg.Sync.CurrentState, firstNonEmpty(cfg.Sync.CurrentRevision, "-"), firstNonEmpty(cfg.Sync.DesiredRevision, "-"), firstNonEmpty(cfg.Sync.SourceRevision, "-"), cfg.Sync.Precedence); err != nil {
			return err
		}
		for _, item := range cfg.Sync.Explanation {
			if _, err := fmt.Fprintf(w, "  - %s\n", item); err != nil {
				return err
			}
		}
		if len(response.ConfigInspection.DefaultsApplied) > 0 {
			if _, err := fmt.Fprintln(w, "\nDefaults applied:"); err != nil {
				return err
			}
			for _, item := range response.ConfigInspection.DefaultsApplied {
				if _, err := fmt.Fprintf(w, "- %s => %s\n", item.Field, item.Value); err != nil {
					return err
				}
			}
		}
		if _, err := fmt.Fprintf(w, "\nRuntime self-healing:\n- state: %s\n- mode: %s\n- fail_mode: %s\n- require_signed_desired_state: %t\n", response.RuntimeInspection.CurrentState, response.RuntimeInspection.EffectiveConfig.Mode, response.RuntimeInspection.EffectiveConfig.FailMode, response.RuntimeInspection.EffectiveConfig.RequireSignedDesiredState); err != nil {
			return err
		}
		for _, issue := range response.RuntimeInspection.Issues {
			if _, err := fmt.Fprintf(w, "- %s: %s\n", issue.Code, issue.Message); err != nil {
				return err
			}
		}
		if len(response.ReasonCodes) > 0 {
			if _, err := fmt.Fprintln(w, "\nReason codes:"); err != nil {
				return err
			}
			for _, reason := range response.ReasonCodes {
				if _, err := fmt.Fprintf(w, "- %s\n", reason); err != nil {
					return err
				}
			}
		}
		if len(response.Limitations) > 0 {
			if _, err := fmt.Fprintln(w, "\nLimitations:"); err != nil {
				return err
			}
			for _, limitation := range response.Limitations {
				if _, err := fmt.Fprintf(w, "- %s\n", limitation); err != nil {
					return err
				}
			}
		}
		return nil
	case "json":
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return encoder.Encode(response)
	default:
		return fmt.Errorf("unsupported output mode %q", output)
	}
}

func renderExplainResponse(w io.Writer, output string, response explainResponse) error {
	switch strings.ToLower(strings.TrimSpace(output)) {
	case "", "human":
		if _, err := fmt.Fprintf(w, "Topic: %s\nState: %s\n", response.Topic, response.CurrentState); err != nil {
			return err
		}
		if len(response.ReasonCodes) > 0 {
			if _, err := fmt.Fprintln(w, "\nReason codes:"); err != nil {
				return err
			}
			for _, reason := range response.ReasonCodes {
				if _, err := fmt.Fprintf(w, "- %s\n", reason); err != nil {
					return err
				}
			}
		}
		if len(response.Explanations) > 0 {
			if _, err := fmt.Fprintln(w, "\nExplanations:"); err != nil {
				return err
			}
			for _, item := range response.Explanations {
				if _, err := fmt.Fprintf(w, "- %s\n", item); err != nil {
					return err
				}
			}
		}
		if len(response.Limitations) > 0 {
			if _, err := fmt.Fprintln(w, "\nLimitations:"); err != nil {
				return err
			}
			for _, limitation := range response.Limitations {
				if _, err := fmt.Fprintf(w, "- %s\n", limitation); err != nil {
					return err
				}
			}
		}
		return nil
	case "json":
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return encoder.Encode(response)
	default:
		return fmt.Errorf("unsupported output mode %q", output)
	}
}

func configValidationSummary(state string) string {
	switch state {
	case "invalid":
		return "production config failed strict schema or compatibility validation"
	case "valid_with_warnings":
		return "production config passed strict parsing with visible compatibility warnings"
	case "valid_with_defaults":
		return "production config passed strict parsing and uses explicit safe defaults"
	default:
		return "production config passed strict validation"
	}
}

func previewStartupSummary(state string) string {
	switch state {
	case "invalid":
		return "startup would fail fast with the current production config"
	case "valid_with_warnings":
		return "startup would proceed, but visible compatibility warnings remain"
	default:
		return "startup would accept the current production config in the local scope"
	}
}

func previewSyncSummary(state string) string {
	switch state {
	case prodcfg.SyncStateConflict:
		return "preview is blocked by visible sync revision conflict"
	case prodcfg.SyncStateStale:
		return "preview is degraded because sync revision data is stale"
	case prodcfg.SyncStateLocalOnly, prodcfg.SyncStateRemoteOnly:
		return "preview is bounded because only one side of sync revision is present"
	default:
		return syncSummary(state)
	}
}

func syncSummary(state string) string {
	switch state {
	case prodcfg.SyncStateConflict:
		return "sync revisions conflict and precedence remains operator-visible"
	case prodcfg.SyncStateStale:
		return "sync revision projection is stale"
	case prodcfg.SyncStateLocalOnly:
		return "only a local revision is declared"
	case prodcfg.SyncStateRemoteOnly:
		return "only a remote revision is declared"
	case prodcfg.SyncStateUnconfigured:
		return "sync revision state is not configured"
	default:
		return "sync revision state is in sync"
	}
}

func compatibilitySummary(issues []prodcfg.ValidationIssue) string {
	if len(issues) == 0 {
		return "no config compatibility issues are visible in the current local scope"
	}
	hasError := false
	for _, issue := range issues {
		if issue.Severity == prodcfg.IssueSeverityError {
			hasError = true
			break
		}
	}
	if hasError {
		return "compatibility issues remain and must be resolved before production use"
	}
	return "compatibility warnings remain visible and should be reviewed before rollout"
}

func statusForValidationState(state string) Status {
	switch state {
	case "invalid":
		return StatusFail
	case "valid_with_warnings":
		return StatusWarning
	case "valid_with_defaults":
		return StatusInfo
	default:
		return StatusPass
	}
}

func statusForCompatibility(issues []prodcfg.ValidationIssue) Status {
	hasWarning := false
	for _, issue := range issues {
		if issue.Severity == prodcfg.IssueSeverityError {
			return StatusFail
		}
		if issue.Severity == prodcfg.IssueSeverityWarning {
			hasWarning = true
		}
	}
	if hasWarning {
		return StatusWarning
	}
	return StatusPass
}

func statusForSyncState(state string) Status {
	switch state {
	case prodcfg.SyncStateConflict:
		return StatusFail
	case prodcfg.SyncStateStale:
		return StatusDegraded
	case prodcfg.SyncStateLocalOnly, prodcfg.SyncStateRemoteOnly:
		return StatusWarning
	case prodcfg.SyncStateUnconfigured:
		return StatusInfo
	default:
		return StatusPass
	}
}

func statusForWorkflow(spec prodcfg.WorkflowSpec) Status {
	switch {
	case spec.ValidationRequired && spec.ApprovalRequired:
		return StatusPass
	case spec.ValidationRequired || spec.ApprovalRequired:
		return StatusWarning
	default:
		return StatusInfo
	}
}

func statusForWorkflowPreview(spec prodcfg.WorkflowSpec) Status {
	switch {
	case spec.ValidationRequired || spec.ApprovalRequired:
		return StatusWarning
	default:
		return StatusInfo
	}
}

func statusForInspectionState(state string) Status {
	switch state {
	case "invalid":
		return StatusFail
	case "valid_with_warnings":
		return StatusWarning
	case "valid_with_defaults":
		return StatusInfo
	default:
		return StatusPass
	}
}

func workflowSummary(spec prodcfg.WorkflowSpec) string {
	switch {
	case spec.ValidationRequired && spec.ApprovalRequired:
		return "workflow closure is explicitly gated by validation and approval"
	case spec.ValidationRequired:
		return "workflow closure requires validation, but approval is not configured"
	case spec.ApprovalRequired:
		return "workflow closure requires approval, but validation is not configured"
	default:
		return "workflow closure has no explicit validation or approval gate in the local config"
	}
}

func previewWorkflowSummary(spec prodcfg.WorkflowSpec) string {
	if spec.ValidationRequired || spec.ApprovalRequired {
		return "workflow preview shows explicit closure blockers that remain until evidence is supplied"
	}
	return "workflow preview cannot claim gated closure because the local config does not declare approval or validation requirements"
}

func workflowDetails(spec prodcfg.WorkflowSpec) []string {
	details := []string{}
	if spec.ValidationRequired {
		details = append(details, "Validation evidence is required before closure.")
	}
	if spec.ApprovalRequired {
		details = append(details, "Approval evidence is required before closure.")
	}
	if len(details) == 0 {
		details = append(details, "No explicit closure blocker is declared in the local workflow config.")
	}
	return details
}

func runtimeInspectionSummary(report runtimecfg.SelfHealingInspection) string {
	switch report.CurrentState {
	case "invalid":
		return "runtime self-healing env config is invalid"
	case "valid_with_warnings":
		return "runtime self-healing env config is valid with visible compatibility warnings"
	case "valid_with_defaults":
		return "runtime self-healing env config is valid and uses explicit defaults"
	default:
		return "runtime self-healing env config is valid"
	}
}

func runtimeInspectionDetails(report runtimecfg.SelfHealingInspection) []string {
	details := []string{}
	for _, issue := range report.Issues {
		details = append(details, fmt.Sprintf("%s: %s", issue.Code, issue.Message))
	}
	for _, item := range report.DefaultsApplied {
		details = append(details, fmt.Sprintf("%s defaulted to %s", item.Field, item.Value))
	}
	return details
}

func runtimeIssueCodes(report runtimecfg.SelfHealingInspection) []string {
	values := []string{}
	for _, issue := range report.Issues {
		values = append(values, issue.Code)
	}
	return values
}

func workflowReasonCodes(spec prodcfg.WorkflowSpec) []string {
	reasons := []string{}
	if spec.ValidationRequired {
		reasons = append(reasons, "workflow_validation_required")
	}
	if spec.ApprovalRequired {
		reasons = append(reasons, "workflow_approval_required")
	}
	if len(reasons) == 0 {
		reasons = append(reasons, "workflow_gates_not_declared")
	}
	return reasons
}

func issueCodesForField(issues []prodcfg.ValidationIssue, prefix string) []string {
	values := []string{}
	for _, issue := range issues {
		if strings.HasPrefix(issue.Field, prefix) {
			values = append(values, issue.Code)
		}
	}
	return values
}

func validationMessages(issues []prodcfg.ValidationIssue) []string {
	lines := []string{}
	for _, issue := range issues {
		lines = append(lines, fmt.Sprintf("%s: %s", issue.Code, issue.Message))
	}
	return lines
}

func combinedInspectionState(configState, runtimeState string) string {
	if configState == "invalid" || runtimeState == "invalid" {
		return "invalid"
	}
	if configState == "valid_with_warnings" || runtimeState == "valid_with_warnings" {
		return "valid_with_warnings"
	}
	if configState == "valid_with_defaults" || runtimeState == "valid_with_defaults" {
		return "valid_with_defaults"
	}
	return "valid"
}

func exitCodeForInspection(state string) int {
	if state == "invalid" {
		return ExitFailed
	}
	return ExitSuccess
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

package preflightcli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	prodcfg "github.com/denisgrosek/changelock/internal/config"
	runtimecfg "github.com/denisgrosek/changelock/internal/runtime"
)

const (
	supportBundleSchemaVersion    = "5c.cli_support_bundle.v1"
	supportHealthSchemaVersion    = "5c.cli_health_snapshot.v1"
	supportMatrixSchemaVersion    = "5c.cli_support_matrix.v1"
	upgradeReadinessSchemaVersion = "5c.cli_upgrade_readiness.v1"

	profileLocal      = "local"
	profileStaging    = "staging"
	profileProduction = "production"
	profileFederated  = "federated"
	profileOffline    = "offline"

	healthStateHealthy  = "healthy"
	healthStateDegraded = "degraded"
	healthStateFailing  = "failing"
	healthStateUnknown  = "unknown"
)

type readinessOptions struct {
	ConfigPath string
	Output     string
	Profile    string
}

type supportOptions struct {
	ConfigPath string
	Output     string
	Profile    string
}

type upgradeReadinessOptions struct {
	ConfigPath     string
	Output         string
	Profile        string
	TargetVersion  string
	CurrentVersion string
}

type supportBundleScope struct {
	Profile     string `json:"profile"`
	TenantID    string `json:"tenant_id"`
	Environment string `json:"environment"`
	Repository  string `json:"repository,omitempty"`
	ConfigPath  string `json:"config_path"`
}

type healthComponent struct {
	Component    string   `json:"component"`
	CurrentState string   `json:"current_state"`
	Summary      string   `json:"summary"`
	ReasonCodes  []string `json:"reason_codes,omitempty"`
	NextAction   string   `json:"next_action,omitempty"`
	OwnerHint    string   `json:"owner_hint,omitempty"`
}

type healthSnapshot struct {
	SchemaVersion string            `json:"schema_version"`
	CurrentState  string            `json:"current_state"`
	Items         []healthComponent `json:"items"`
}

type operatorIssue struct {
	Severity   string `json:"severity"`
	ReasonCode string `json:"reason_code"`
	Subsystem  string `json:"subsystem"`
	Summary    string `json:"summary"`
	NextAction string `json:"next_action,omitempty"`
	OwnerHint  string `json:"owner_hint,omitempty"`
	Blocking   bool   `json:"blocking"`
}

type readinessSummary struct {
	OverallResult     Status            `json:"overall_result"`
	ExitCode          int               `json:"exit_code"`
	Checks            []CheckResult     `json:"checks"`
	DiagnosticSummary DiagnosticSummary `json:"diagnostic_summary"`
	ReasonCodes       []string          `json:"reason_codes,omitempty"`
	Limitations       []string          `json:"limitations,omitempty"`
}

type supportBundleResponse struct {
	SchemaVersion     string                           `json:"schema_version"`
	GeneratedAt       time.Time                        `json:"generated_at"`
	CurrentState      string                           `json:"current_state"`
	Scope             supportBundleScope               `json:"scope"`
	RedactionState    string                           `json:"redaction_state"`
	Version           VersionInfo                      `json:"version"`
	Health            healthSnapshot                   `json:"health"`
	ConfigInspection  prodcfg.InspectionReport         `json:"config_inspection"`
	RuntimeInspection runtimecfg.SelfHealingInspection `json:"runtime_inspection"`
	Readiness         readinessSummary                 `json:"readiness"`
	OperatorIssues    []operatorIssue                  `json:"operator_issues,omitempty"`
	DiagnosticRefs    []string                         `json:"diagnostic_refs,omitempty"`
	Limitations       []string                         `json:"limitations,omitempty"`
}

type supportMatrix struct {
	SchemaVersion         string   `json:"schema_version"`
	CurrentVersion        string   `json:"current_version"`
	TargetVersion         string   `json:"target_version"`
	CurrentLine           string   `json:"current_line,omitempty"`
	TargetLine            string   `json:"target_line,omitempty"`
	CurrentState          string   `json:"current_state"`
	RollbackSupported     bool     `json:"rollback_supported"`
	CompatibilityGuidance []string `json:"compatibility_guidance,omitempty"`
	RollbackCautions      []string `json:"rollback_cautions,omitempty"`
}

type upgradeReadinessResponse struct {
	SchemaVersion string        `json:"schema_version"`
	CurrentState  string        `json:"current_state"`
	Profile       string        `json:"profile"`
	ConfigPath    string        `json:"config_path"`
	SupportMatrix supportMatrix `json:"support_matrix"`
	Result        Result        `json:"result"`
	Limitations   []string      `json:"limitations,omitempty"`
}

type parsedReleaseVersion struct {
	Raw   string
	Line  string
	Major int
	Minor int
	OK    bool
}

type apiProbeOutcome struct {
	client *APIClient
	check  CheckResult
}

func (a *App) runReadiness(ctx context.Context, args []string) (Result, error) {
	options, err := a.parseReadinessOptions(args)
	if err != nil {
		return Result{}, err
	}
	a.config.Output = options.Output
	contextReport, loadErr := a.loadProductionContext(options.ConfigPath)
	if loadErr != nil {
		return configLoadFailureResult("readiness", options.ConfigPath, loadErr), nil
	}
	return a.buildReadinessResult(ctx, contextReport, options), nil
}

func (a *App) runSupport(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	options, err := a.parseSupportOptions(args)
	if err != nil {
		return a.writeError(stderr, err)
	}
	a.config.Output = options.Output
	contextReport, loadErr := a.loadProductionContext(options.ConfigPath)
	response := supportBundleResponse{
		SchemaVersion:  supportBundleSchemaVersion,
		GeneratedAt:    time.Now().UTC(),
		CurrentState:   "invalid",
		RedactionState: "redacted_by_default",
		Version:        a.runtime.VersionInfo,
		Scope: supportBundleScope{
			Profile:    options.Profile,
			ConfigPath: options.ConfigPath,
		},
		Limitations: []string{
			"Support bundle is bounded, redacted by default, and built from the same config, runtime inspection, and audit-aware supportability surfaces already present in ChangeLock.",
			"Surface health probes reflect lightweight API reachability and proof/timeline checks; they do not create a second diagnostic truth store.",
		},
	}
	if loadErr != nil {
		response.OperatorIssues = []operatorIssue{{
			Severity:   "error",
			ReasonCode: "config_load_failed",
			Subsystem:  "config",
			Summary:    loadErr.Error(),
			NextAction: "Fix the production config path or contents before collecting a support bundle.",
			OwnerHint:  "platform_operator",
			Blocking:   true,
		}}
		if err := renderSupportBundle(stdout, options.Output, response); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitFailed
	}

	effectiveProfile := normalizeProfile(options.Profile, contextReport.Config.Spec.Environment, contextReport.Config.Spec.CLI.Offline)
	readiness := finalizeResult(a.buildReadinessResult(ctx, contextReport, readinessOptions{
		ConfigPath: options.ConfigPath,
		Output:     options.Output,
		Profile:    effectiveProfile,
	}))
	health := a.buildHealthSnapshot(ctx, contextReport, effectiveProfile)
	response.CurrentState = supportBundleState(readiness, health)
	response.Scope = supportBundleScope{
		Profile:     effectiveProfile,
		TenantID:    strings.TrimSpace(contextReport.Config.Spec.TenantID),
		Environment: strings.TrimSpace(contextReport.Config.Spec.Environment),
		Repository:  strings.TrimSpace(contextReport.Config.Spec.Repository),
		ConfigPath:  contextReport.Path,
	}
	response.Health = health
	response.ConfigInspection = contextReport.Inspection
	response.RuntimeInspection = redactRuntimeInspection(contextReport.RuntimeInspection)
	response.Readiness = readinessSummary{
		OverallResult:     readiness.OverallResult,
		ExitCode:          readiness.ExitCode,
		Checks:            readiness.Checks,
		DiagnosticSummary: readiness.DiagnosticSummary,
		ReasonCodes:       readiness.ReasonCodes,
		Limitations:       readiness.Limitations,
	}
	response.OperatorIssues = diagnosticsToOperatorIssues(readiness.Diagnostics)
	response.DiagnosticRefs = []string{
		"changelock-cli readiness --config <path>",
		"changelock-cli diagnostics --input <preflight-json> --format markdown",
		"docs/production-phase5-valb.md",
		"docs/production-phase5-valc.md",
	}
	if err := renderSupportBundle(stdout, options.Output, response); err != nil {
		return a.writeError(stderr, err)
	}
	if response.CurrentState == healthStateFailing {
		return ExitFailed
	}
	return ExitSuccess
}

func (a *App) runUpgradeReadiness(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	options, err := a.parseUpgradeReadinessOptions(args)
	if err != nil {
		return a.writeError(stderr, err)
	}
	a.config.Output = options.Output
	contextReport, loadErr := a.loadProductionContext(options.ConfigPath)
	response := upgradeReadinessResponse{
		SchemaVersion: upgradeReadinessSchemaVersion,
		Profile:       options.Profile,
		ConfigPath:    options.ConfigPath,
		CurrentState:  "invalid",
		Limitations: []string{
			"Upgrade readiness is bounded to local config compatibility, runtime env config, sync revision visibility, and the built-in support matrix baseline.",
			"Rollback guidance is advisory and does not claim that runtime, enterprise workflow, or evidence history is already validated post-transition without re-running readiness.",
		},
	}
	if loadErr != nil {
		response.Result = finalizeResult(configLoadFailureResult("upgrade-readiness", options.ConfigPath, loadErr))
		response.SupportMatrix = evaluateSupportMatrix(options.CurrentVersion, options.TargetVersion)
		if err := renderUpgradeReadiness(stdout, options.Output, response); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitFailed
	}

	effectiveProfile := normalizeProfile(options.Profile, contextReport.Config.Spec.Environment, contextReport.Config.Spec.CLI.Offline)
	options.Profile = effectiveProfile
	result, matrix := a.buildUpgradeReadinessResult(ctx, contextReport, options)
	response.Result = finalizeResult(result)
	response.SupportMatrix = matrix
	response.CurrentState = matrix.CurrentState
	response.Profile = effectiveProfile
	response.ConfigPath = contextReport.Path
	if err := renderUpgradeReadiness(stdout, options.Output, response); err != nil {
		return a.writeError(stderr, err)
	}
	return response.Result.ExitCode
}

func (a *App) parseReadinessOptions(args []string) (readinessOptions, error) {
	fs := newFlagSet("readiness")
	options := readinessOptions{
		ConfigPath: a.config.ProductionConfigPath,
		Output:     a.config.Output,
		Profile:    "",
	}
	fs.StringVar(&options.ConfigPath, "config", options.ConfigPath, "production config file")
	fs.StringVar(&options.Output, "output", options.Output, "output mode: human|json")
	fs.StringVar(&options.Profile, "profile", options.Profile, "environment profile: local|staging|production|federated|offline")
	if err := fs.Parse(args); err != nil {
		return readinessOptions{}, usageError{message: err.Error()}
	}
	if strings.TrimSpace(options.ConfigPath) == "" {
		return readinessOptions{}, usageError{message: "readiness requires --config or CHANGELOCK_CLI_CONFIG"}
	}
	if err := validateOutputMode(options.Output); err != nil {
		return readinessOptions{}, err
	}
	if err := validateProfileValue(options.Profile); err != nil {
		return readinessOptions{}, err
	}
	return options, nil
}

func (a *App) parseSupportOptions(args []string) (supportOptions, error) {
	fs := newFlagSet("support")
	options := supportOptions{
		ConfigPath: a.config.ProductionConfigPath,
		Output:     a.config.Output,
		Profile:    "",
	}
	fs.StringVar(&options.ConfigPath, "config", options.ConfigPath, "production config file")
	fs.StringVar(&options.Output, "output", options.Output, "output mode: human|json")
	fs.StringVar(&options.Profile, "profile", options.Profile, "environment profile: local|staging|production|federated|offline")
	if err := fs.Parse(args); err != nil {
		return supportOptions{}, usageError{message: err.Error()}
	}
	if strings.TrimSpace(options.ConfigPath) == "" {
		return supportOptions{}, usageError{message: "support requires --config or CHANGELOCK_CLI_CONFIG"}
	}
	if err := validateOutputMode(options.Output); err != nil {
		return supportOptions{}, err
	}
	if err := validateProfileValue(options.Profile); err != nil {
		return supportOptions{}, err
	}
	return options, nil
}

func (a *App) parseUpgradeReadinessOptions(args []string) (upgradeReadinessOptions, error) {
	fs := newFlagSet("upgrade-readiness")
	options := upgradeReadinessOptions{
		ConfigPath:     a.config.ProductionConfigPath,
		Output:         a.config.Output,
		Profile:        "",
		CurrentVersion: strings.TrimSpace(a.runtime.VersionInfo.Version),
	}
	fs.StringVar(&options.ConfigPath, "config", options.ConfigPath, "production config file")
	fs.StringVar(&options.Output, "output", options.Output, "output mode: human|json")
	fs.StringVar(&options.Profile, "profile", options.Profile, "environment profile: local|staging|production|federated|offline")
	fs.StringVar(&options.TargetVersion, "target-version", "", "target ChangeLock version")
	fs.StringVar(&options.CurrentVersion, "current-version", options.CurrentVersion, "current ChangeLock version override")
	if err := fs.Parse(args); err != nil {
		return upgradeReadinessOptions{}, usageError{message: err.Error()}
	}
	if strings.TrimSpace(options.ConfigPath) == "" {
		return upgradeReadinessOptions{}, usageError{message: "upgrade-readiness requires --config or CHANGELOCK_CLI_CONFIG"}
	}
	if strings.TrimSpace(options.TargetVersion) == "" {
		return upgradeReadinessOptions{}, usageError{message: "upgrade-readiness requires --target-version"}
	}
	if err := validateOutputMode(options.Output); err != nil {
		return upgradeReadinessOptions{}, err
	}
	if err := validateProfileValue(options.Profile); err != nil {
		return upgradeReadinessOptions{}, err
	}
	return options, nil
}

func validateOutputMode(output string) error {
	switch strings.ToLower(strings.TrimSpace(output)) {
	case "", "human", "json":
		return nil
	default:
		return usageError{message: fmt.Sprintf("unsupported output mode %q", output)}
	}
}

func validateProfileValue(profile string) error {
	switch strings.ToLower(strings.TrimSpace(profile)) {
	case "", profileLocal, profileStaging, profileProduction, profileFederated, profileOffline:
		return nil
	default:
		return usageError{message: fmt.Sprintf("unsupported profile %q", profile)}
	}
}

func (a *App) buildReadinessResult(ctx context.Context, contextReport productionContext, options readinessOptions) Result {
	profile := normalizeProfile(options.Profile, contextReport.Config.Spec.Environment, contextReport.Config.Spec.CLI.Offline)
	apiProbe := a.apiConnectivityCheck(ctx, contextReport, profile)
	checks := []CheckResult{
		{
			Name:    "config-readiness",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  statusForValidationState(contextReport.Inspection.Validation.CurrentState),
			Summary: configValidationSummary(contextReport.Inspection.Validation.CurrentState),
			Details: validationMessages(contextReport.Inspection.Validation.Issues),
		},
		{
			Name:    "dependency-readiness",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  a.dependencyReadinessStatus(ctx, contextReport, profile),
			Summary: a.dependencyReadinessSummary(ctx, contextReport, profile),
			Details: a.dependencyReadinessDetails(ctx, contextReport, profile),
		},
		apiProbe.check,
		{
			Name:    "sync-readiness",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  readinessStatusForSyncState(contextReport.Inspection.EffectiveConfig.Sync.CurrentState, profile),
			Summary: readinessSyncSummary(contextReport.Inspection.EffectiveConfig.Sync.CurrentState, profile),
			Details: contextReport.Inspection.EffectiveConfig.Sync.Explanation,
		},
		{
			Name:    "workflow-readiness",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  readinessStatusForWorkflow(contextReport.Config.Spec.Workflow, profile),
			Summary: workflowReadinessSummary(contextReport.Config.Spec.Workflow, profile),
			Details: workflowDetails(contextReport.Config.Spec.Workflow),
		},
		{
			Name:    "runtime-self-healing-readiness",
			Mode:    ModeLocal,
			Target:  "runtime-self-healing",
			Status:  readinessStatusForRuntimeInspection(contextReport.RuntimeInspection, profile),
			Summary: runtimeInspectionSummary(contextReport.RuntimeInspection),
			Details: runtimeInspectionDetails(contextReport.RuntimeInspection),
		},
	}

	reasonCodes := uniqueStrings(append(append([]string{
		"go_live_profile_" + profile,
	}, contextReport.Inspection.ReasonCodes...), runtimeIssueCodes(contextReport.RuntimeInspection)...))
	return finalizeResult(Result{
		Command: "readiness",
		Mode:    executionMode(contextReport.Config.Spec.CLI.Offline, contextReport.Config.Spec.APIURL),
		Inputs: map[string]string{
			"config":      contextReport.Path,
			"profile":     profile,
			"tenant_id":   contextReport.Config.Spec.TenantID,
			"environment": contextReport.Config.Spec.Environment,
			"repository":  contextReport.Config.Spec.Repository,
		},
		Checks:      checks,
		ReasonCodes: reasonCodes,
		Limitations: []string{
			"Readiness is a bounded go-live validator over local config, env config, binary prerequisites, and optional API reachability; it does not claim runtime workload evidence was exercised.",
			"Warnings and degraded states remain visible and are not silently upgraded into a full production pass.",
		},
	})
}

func (a *App) buildHealthSnapshot(ctx context.Context, contextReport productionContext, profile string) healthSnapshot {
	profile = normalizeProfile(profile, contextReport.Config.Spec.Environment, contextReport.Config.Spec.CLI.Offline)
	items := []healthComponent{
		{
			Component:    "config",
			CurrentState: healthStateForStatus(statusForValidationState(contextReport.Inspection.Validation.CurrentState)),
			Summary:      configValidationSummary(contextReport.Inspection.Validation.CurrentState),
			ReasonCodes:  uniqueStrings(contextReport.Inspection.ReasonCodes),
			NextAction:   "Resolve any invalid or warning-level config findings before rollout.",
			OwnerHint:    "platform_operator",
		},
		{
			Component:    "sync",
			CurrentState: healthStateForStatus(readinessStatusForSyncState(contextReport.Inspection.EffectiveConfig.Sync.CurrentState, profile)),
			Summary:      readinessSyncSummary(contextReport.Inspection.EffectiveConfig.Sync.CurrentState, profile),
			ReasonCodes:  uniqueStrings([]string{contextReport.Inspection.EffectiveConfig.Sync.CurrentState}),
			NextAction:   "Converge local and remote revisions or document why bounded divergence is acceptable.",
			OwnerHint:    "platform_operator",
		},
		{
			Component:    "runtime-self-healing",
			CurrentState: healthStateForStatus(readinessStatusForRuntimeInspection(contextReport.RuntimeInspection, profile)),
			Summary:      runtimeInspectionSummary(contextReport.RuntimeInspection),
			ReasonCodes:  runtimeIssueCodes(contextReport.RuntimeInspection),
			NextAction:   "Confirm remediation mode, signed desired-state requirements, and quarantine posture before go-live.",
			OwnerHint:    "platform_operator",
		},
	}

	apiProbe := a.apiConnectivityCheck(ctx, contextReport, profile)
	items = append(items, healthComponent{
		Component:    "api",
		CurrentState: healthStateForStatus(apiProbe.check.Status),
		Summary:      apiProbe.check.Summary,
		ReasonCodes:  diagnosticReasonCodesForDetails(apiProbe.check),
		NextAction:   "Restore API reachability or declare this node offline if API-assisted support surfaces are intentionally unavailable.",
		OwnerHint:    "platform_operator",
	})

	if apiProbe.client != nil {
		items = append(items,
			a.surfaceHealthCheck(ctx, apiProbe.client, "command-center-surface", commandCenterProbePath(contextReport), "Bounded operator timeline and grouped notification surfaces are reachable."),
			a.surfaceHealthCheck(ctx, apiProbe.client, "runtime-surface", phase2ProofProbePath(contextReport), "Runtime proofs surface is reachable for support verification."),
			a.surfaceHealthCheck(ctx, apiProbe.client, "intelligence-surface", phase3ProofProbePath(contextReport), "Intelligence proofs surface is reachable for support verification."),
			a.surfaceHealthCheck(ctx, apiProbe.client, "enterprise-surface", phase4ProofProbePath(contextReport), "Enterprise proofs surface is reachable for support verification."),
		)
	} else {
		for _, component := range []string{"command-center-surface", "runtime-surface", "intelligence-surface", "enterprise-surface"} {
			items = append(items, healthComponent{
				Component:    component,
				CurrentState: healthStateUnknown,
				Summary:      "API-assisted health probe is not configured for this support snapshot.",
				ReasonCodes:  []string{"api_probe_unavailable"},
				NextAction:   "Set spec.api_url and a read-only token if you need remote surface verification in support snapshots.",
				OwnerHint:    "platform_operator",
			})
		}
	}

	return healthSnapshot{
		SchemaVersion: supportHealthSchemaVersion,
		CurrentState:  summarizeHealthStates(items),
		Items:         items,
	}
}

func (a *App) buildUpgradeReadinessResult(ctx context.Context, contextReport productionContext, options upgradeReadinessOptions) (Result, supportMatrix) {
	profile := normalizeProfile(options.Profile, contextReport.Config.Spec.Environment, contextReport.Config.Spec.CLI.Offline)
	matrix := evaluateSupportMatrix(options.CurrentVersion, options.TargetVersion)
	apiProbe := a.apiConnectivityCheck(ctx, contextReport, profile)
	apiUpgradeCheck := apiProbe.check
	if strings.TrimSpace(contextReport.Config.Spec.APIURL) == "" && apiUpgradeCheck.Status == StatusFail {
		apiUpgradeCheck.Status = StatusWarning
		apiUpgradeCheck.Summary = "API-assisted upgrade context is not configured, so upgrade assessment remains local-only"
		apiUpgradeCheck.Details = []string{"Set spec.api_url and a read-only token if you want remote proof and health verification during upgrade planning."}
	}
	checks := []CheckResult{
		{
			Name:    "config-compatibility-readiness",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  statusForCompatibility(contextReport.Inspection.Validation.Issues),
			Summary: compatibilitySummary(contextReport.Inspection.Validation.Issues),
			Details: validationMessages(contextReport.Inspection.Validation.Issues),
		},
		{
			Name:    "sync-migration-readiness",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  readinessStatusForSyncState(contextReport.Inspection.EffectiveConfig.Sync.CurrentState, profile),
			Summary: readinessSyncSummary(contextReport.Inspection.EffectiveConfig.Sync.CurrentState, profile),
			Details: contextReport.Inspection.EffectiveConfig.Sync.Explanation,
		},
		{
			Name:    "version-support-matrix",
			Mode:    ModeLocal,
			Target:  firstNonEmpty(options.TargetVersion, options.CurrentVersion),
			Status:  statusForSupportMatrix(matrix),
			Summary: supportMatrixSummary(matrix),
			Details: uniqueStrings(append(append([]string{}, matrix.CompatibilityGuidance...), matrix.RollbackCautions...)),
			Metadata: map[string]any{
				"current_version":    matrix.CurrentVersion,
				"target_version":     matrix.TargetVersion,
				"current_line":       matrix.CurrentLine,
				"target_line":        matrix.TargetLine,
				"rollback_supported": matrix.RollbackSupported,
			},
		},
		{
			Name:    "rollback-readiness",
			Mode:    ModeLocal,
			Target:  firstNonEmpty(options.TargetVersion, options.CurrentVersion),
			Status:  statusForRollbackMatrix(matrix),
			Summary: rollbackSummary(matrix),
			Details: matrix.RollbackCautions,
		},
		{
			Name:    "api-upgrade-context",
			Mode:    ModeRemote,
			Target:  firstNonEmpty(strings.TrimSpace(contextReport.Config.Spec.APIURL), "(offline)"),
			Status:  apiUpgradeCheck.Status,
			Summary: apiUpgradeCheck.Summary,
			Details: apiUpgradeCheck.Details,
		},
	}

	result := finalizeResult(Result{
		Command: "upgrade-readiness",
		Mode:    executionMode(contextReport.Config.Spec.CLI.Offline, contextReport.Config.Spec.APIURL),
		Inputs: map[string]string{
			"config":          contextReport.Path,
			"profile":         profile,
			"current_version": options.CurrentVersion,
			"target_version":  options.TargetVersion,
		},
		Checks: checks,
		ReasonCodes: uniqueStrings(append(append([]string{
			matrix.CurrentState,
			"upgrade_profile_" + profile,
		}, contextReport.Inspection.ReasonCodes...), runtimeIssueCodes(contextReport.RuntimeInspection)...)),
		Limitations: []string{
			"Upgrade readiness does not execute migrations; it only classifies local compatibility, API reachability, revision visibility, and bounded support matrix risk.",
			"Rollback notes stay advisory and must be paired with post-change readiness and evidence review.",
		},
	})
	return result, matrix
}

func (a *App) apiConnectivityCheck(ctx context.Context, contextReport productionContext, profile string) apiProbeOutcome {
	config := Config{
		APIURL:  firstNonEmpty(strings.TrimSpace(contextReport.Config.Spec.APIURL), a.config.APIURL),
		Token:   strings.TrimSpace(a.config.Token),
		Timeout: a.config.Timeout,
		Offline: contextReport.Config.Spec.CLI.Offline || profile == profileOffline,
	}
	client := NewAPIClient(config, a.runtime.HTTPClient)
	if client == nil {
		status := StatusInfo
		summary := "API-assisted readiness checks are not configured in this scope"
		if profile == profileProduction || profile == profileFederated {
			status = StatusFail
			summary = "API-assisted readiness is required for production or federated go-live"
		} else if profile == profileStaging {
			status = StatusWarning
			summary = "API-assisted readiness is not configured, so supportability remains partially local-only"
		}
		return apiProbeOutcome{
			check: CheckResult{
				Name:    "api-readiness",
				Mode:    ModeRemote,
				Target:  firstNonEmpty(strings.TrimSpace(contextReport.Config.Spec.APIURL), "(offline)"),
				Status:  status,
				Summary: summary,
				Details: []string{"Set spec.api_url plus a read-only token if this node should verify server-side health and proof surfaces."},
			},
		}
	}
	if err := client.Healthz(ctx); err != nil {
		status := StatusDegraded
		if profile == profileProduction || profile == profileFederated {
			status = StatusFail
		}
		return apiProbeOutcome{
			check: CheckResult{
				Name:    "api-readiness",
				Mode:    ModeRemote,
				Target:  config.APIURL,
				Status:  status,
				Summary: "ChangeLock API health probe failed",
				Details: []string{err.Error()},
			},
		}
	}
	authInfo, err := client.AuthMe(ctx)
	if err != nil {
		status := StatusWarning
		if profile == profileProduction || profile == profileFederated {
			status = StatusFail
		}
		return apiProbeOutcome{
			client: client,
			check: CheckResult{
				Name:    "api-readiness",
				Mode:    ModeRemote,
				Target:  config.APIURL,
				Status:  status,
				Summary: "ChangeLock API is reachable, but auth scope lookup failed",
				Details: []string{err.Error()},
			},
		}
	}
	return apiProbeOutcome{
		client: client,
		check: CheckResult{
			Name:    "api-readiness",
			Mode:    ModeRemote,
			Target:  config.APIURL,
			Status:  StatusPass,
			Summary: fmt.Sprintf("ChangeLock API is reachable and authenticated as %s", firstNonEmpty(authInfo.Role, authInfo.AuthMode)),
			Details: []string{
				fmt.Sprintf("tenant scope: %s", firstNonEmpty(authInfo.TenantID, "global")),
				fmt.Sprintf("auth mode: %s", firstNonEmpty(authInfo.AuthMode, "unknown")),
			},
		},
	}
}

func (a *App) dependencyReadinessStatus(ctx context.Context, contextReport productionContext, profile string) Status {
	issues := a.toolReadinessIssues(ctx, contextReport)
	if len(issues) == 0 {
		return StatusPass
	}
	if profile == profileProduction || profile == profileFederated {
		return StatusFail
	}
	if profile == profileStaging {
		return StatusWarning
	}
	return StatusInfo
}

func (a *App) dependencyReadinessSummary(ctx context.Context, contextReport productionContext, profile string) string {
	issues := a.toolReadinessIssues(ctx, contextReport)
	if len(issues) == 0 {
		return "local binaries and policy bundle prerequisites are present"
	}
	if profile == profileProduction || profile == profileFederated {
		return "required local prerequisites are missing for bounded production supportability"
	}
	return "some local prerequisites are missing, but this profile can continue only with reduced supportability"
}

func (a *App) dependencyReadinessDetails(ctx context.Context, contextReport productionContext, profile string) []string {
	issues := a.toolReadinessIssues(ctx, contextReport)
	if len(issues) == 0 {
		return []string{"Kyverno, cosign, and vulnerability scanner prerequisites are available for local supportability checks."}
	}
	sort.Strings(issues)
	return issues
}

func (a *App) toolReadinessIssues(ctx context.Context, contextReport productionContext) []string {
	issues := []string{}
	if _, err := a.runtime.RunCommand(ctx, a.config.KyvernoBin); err != nil {
		issues = append(issues, fmt.Sprintf("kyverno binary %q is unavailable: %v", a.config.KyvernoBin, err))
	}
	if _, err := a.runtime.RunCommand(ctx, a.config.CosignBin); err != nil {
		issues = append(issues, fmt.Sprintf("cosign binary %q is unavailable: %v", a.config.CosignBin, err))
	}
	scanner := strings.ToLower(strings.TrimSpace(contextReport.Inspection.EffectiveConfig.CLI.Scanner))
	switch scanner {
	case "", "auto":
		if _, err := a.runtime.RunCommand(ctx, a.config.TrivyBin); err != nil {
			if _, fallbackErr := a.runtime.RunCommand(ctx, a.config.GrypeBin); fallbackErr != nil {
				issues = append(issues, fmt.Sprintf("no scanner is available via %q or %q", a.config.TrivyBin, a.config.GrypeBin))
			}
		}
	case "trivy":
		if _, err := a.runtime.RunCommand(ctx, a.config.TrivyBin); err != nil {
			issues = append(issues, fmt.Sprintf("trivy binary %q is unavailable: %v", a.config.TrivyBin, err))
		}
	case "grype":
		if _, err := a.runtime.RunCommand(ctx, a.config.GrypeBin); err != nil {
			issues = append(issues, fmt.Sprintf("grype binary %q is unavailable: %v", a.config.GrypeBin, err))
		}
	}
	return uniqueStrings(issues)
}

func readinessStatusForSyncState(state, profile string) Status {
	switch state {
	case prodcfg.SyncStateConflict:
		return StatusFail
	case prodcfg.SyncStateStale:
		if profile == profileProduction || profile == profileFederated {
			return StatusDegraded
		}
		return StatusWarning
	case prodcfg.SyncStateLocalOnly, prodcfg.SyncStateRemoteOnly:
		if profile == profileProduction || profile == profileFederated {
			return StatusWarning
		}
		return StatusInfo
	case prodcfg.SyncStateUnconfigured:
		if profile == profileProduction || profile == profileFederated {
			return StatusFail
		}
		if profile == profileStaging {
			return StatusWarning
		}
		return StatusInfo
	default:
		return StatusPass
	}
}

func readinessSyncSummary(state, profile string) string {
	switch state {
	case prodcfg.SyncStateConflict:
		return "sync revisions conflict and block a bounded go-live decision"
	case prodcfg.SyncStateStale:
		return "sync revision projection is stale and keeps rollout posture degraded"
	case prodcfg.SyncStateLocalOnly, prodcfg.SyncStateRemoteOnly:
		return "only one side of sync revision is declared, so convergence remains partial"
	case prodcfg.SyncStateUnconfigured:
		if profile == profileProduction || profile == profileFederated {
			return "production or federated go-live requires explicit sync revision state"
		}
		return "sync revision state is not configured for this profile"
	default:
		return "sync revision state is in sync"
	}
}

func readinessStatusForWorkflow(spec prodcfg.WorkflowSpec, profile string) Status {
	switch {
	case spec.ValidationRequired && spec.ApprovalRequired:
		return StatusPass
	case profile == profileProduction || profile == profileFederated:
		if spec.ValidationRequired || spec.ApprovalRequired {
			return StatusWarning
		}
		return StatusDegraded
	case spec.ValidationRequired || spec.ApprovalRequired:
		return StatusInfo
	default:
		return StatusInfo
	}
}

func workflowReadinessSummary(spec prodcfg.WorkflowSpec, profile string) string {
	switch {
	case spec.ValidationRequired && spec.ApprovalRequired:
		return "workflow closure is gated by validation and approval"
	case profile == profileProduction || profile == profileFederated:
		return "workflow closure posture is partially declared and should be tightened before production rollout"
	case spec.ValidationRequired || spec.ApprovalRequired:
		return "workflow closure has one explicit gate in the current config"
	default:
		return "workflow closure remains lightly governed in this profile"
	}
}

func readinessStatusForRuntimeInspection(report runtimecfg.SelfHealingInspection, profile string) Status {
	switch report.CurrentState {
	case "invalid":
		return StatusFail
	case "valid_with_warnings":
		if profile == profileProduction || profile == profileFederated {
			return StatusWarning
		}
		return StatusInfo
	case "valid_with_defaults":
		return StatusInfo
	default:
		return StatusPass
	}
}

func healthStateForStatus(status Status) string {
	switch status {
	case StatusPass:
		return healthStateHealthy
	case StatusWarning, StatusDegraded:
		return healthStateDegraded
	case StatusFail, StatusError:
		return healthStateFailing
	default:
		return healthStateUnknown
	}
}

func summarizeHealthStates(items []healthComponent) string {
	hasHealthy := false
	hasUnknown := false
	for _, item := range items {
		switch item.CurrentState {
		case healthStateFailing:
			return healthStateFailing
		case healthStateDegraded:
			return healthStateDegraded
		case healthStateHealthy:
			hasHealthy = true
		case healthStateUnknown:
			hasUnknown = true
		}
	}
	if hasHealthy && !hasUnknown {
		return healthStateHealthy
	}
	if hasHealthy || hasUnknown {
		return healthStateDegraded
	}
	return healthStateUnknown
}

func supportBundleState(readiness Result, health healthSnapshot) string {
	switch readiness.OverallResult {
	case StatusFail, StatusError:
		return healthStateFailing
	case StatusWarning, StatusDegraded:
		if health.CurrentState == healthStateFailing {
			return healthStateFailing
		}
		return healthStateDegraded
	}
	switch health.CurrentState {
	case healthStateFailing:
		return healthStateFailing
	case healthStateDegraded:
		return healthStateDegraded
	}
	return health.CurrentState
}

func redactRuntimeInspection(report runtimecfg.SelfHealingInspection) runtimecfg.SelfHealingInspection {
	clone := report
	if len(clone.DeclaredValues) > 0 {
		redacted := make(map[string]string, len(clone.DeclaredValues))
		for key, value := range clone.DeclaredValues {
			if strings.TrimSpace(value) == "" {
				continue
			}
			redacted[key] = "<redacted>"
		}
		clone.DeclaredValues = redacted
	}
	return clone
}

func diagnosticsToOperatorIssues(diagnostics []Diagnostic) []operatorIssue {
	issues := make([]operatorIssue, 0, len(diagnostics))
	for _, diagnostic := range filterDiagnostics(diagnostics, false) {
		issues = append(issues, operatorIssue{
			Severity:   string(diagnostic.Severity),
			ReasonCode: diagnostic.ReasonCode,
			Subsystem:  diagnostic.Source,
			Summary:    diagnostic.Summary,
			NextAction: diagnostic.FixHint,
			OwnerHint:  ownerHintForDiagnostic(diagnostic.Source),
			Blocking:   diagnostic.Blocking,
		})
	}
	if len(issues) > 8 {
		issues = issues[:8]
	}
	return issues
}

func ownerHintForDiagnostic(source string) string {
	switch strings.TrimSpace(source) {
	case "policy", "evidence", "runtime":
		return "platform_operator"
	case "vulnerability", "vex", "signer_identity":
		return "security_engineer"
	default:
		return "platform_operator"
	}
}

func diagnosticReasonCodesForDetails(check CheckResult) []string {
	reasonCodes := []string{}
	if strings.TrimSpace(check.Name) != "" {
		reasonCodes = append(reasonCodes, strings.ReplaceAll(check.Name, "-", "_"))
	}
	for _, line := range check.Details {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			reasonCodes = append(reasonCodes, trimmed)
		}
	}
	return uniqueStrings(reasonCodes)
}

func renderSupportBundle(w io.Writer, output string, response supportBundleResponse) error {
	switch strings.ToLower(strings.TrimSpace(output)) {
	case "", "human":
		if _, err := fmt.Fprintf(w, "Support bundle\nConfig: %s\nProfile: %s\nState: %s\nRedaction: %s\nVersion: %s\nGenerated at: %s\n", response.Scope.ConfigPath, response.Scope.Profile, response.CurrentState, response.RedactionState, firstNonEmpty(response.Version.Version, "dev"), response.GeneratedAt.Format(time.RFC3339)); err != nil {
			return err
		}
		if _, err := fmt.Fprintln(w, "\nHealth snapshot:"); err != nil {
			return err
		}
		for _, item := range response.Health.Items {
			if _, err := fmt.Fprintf(w, "- %s [%s]: %s\n", item.Component, item.CurrentState, item.Summary); err != nil {
				return err
			}
		}
		if len(response.OperatorIssues) > 0 {
			if _, err := fmt.Fprintln(w, "\nOperator issues:"); err != nil {
				return err
			}
			for _, issue := range response.OperatorIssues {
				if _, err := fmt.Fprintf(w, "- %s %s (%s): %s\n", issue.Severity, issue.ReasonCode, issue.Subsystem, issue.Summary); err != nil {
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

func renderUpgradeReadiness(w io.Writer, output string, response upgradeReadinessResponse) error {
	switch strings.ToLower(strings.TrimSpace(output)) {
	case "", "human":
		if _, err := fmt.Fprintf(w, "Upgrade readiness\nConfig: %s\nProfile: %s\nState: %s\nCurrent version: %s\nTarget version: %s\nSupport matrix: %s\nRollback supported: %t\n", response.ConfigPath, response.Profile, response.CurrentState, firstNonEmpty(response.SupportMatrix.CurrentVersion, "unknown"), response.SupportMatrix.TargetVersion, response.SupportMatrix.CurrentState, response.SupportMatrix.RollbackSupported); err != nil {
			return err
		}
		if len(response.SupportMatrix.CompatibilityGuidance) > 0 {
			if _, err := fmt.Fprintln(w, "\nCompatibility guidance:"); err != nil {
				return err
			}
			for _, item := range response.SupportMatrix.CompatibilityGuidance {
				if _, err := fmt.Fprintf(w, "- %s\n", item); err != nil {
					return err
				}
			}
		}
		if len(response.SupportMatrix.RollbackCautions) > 0 {
			if _, err := fmt.Fprintln(w, "\nRollback cautions:"); err != nil {
				return err
			}
			for _, item := range response.SupportMatrix.RollbackCautions {
				if _, err := fmt.Fprintf(w, "- %s\n", item); err != nil {
					return err
				}
			}
		}
		if _, err := fmt.Fprintln(w); err != nil {
			return err
		}
		return renderHuman(w, response.Result)
	case "json":
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return encoder.Encode(response)
	default:
		return fmt.Errorf("unsupported output mode %q", output)
	}
}

func commandCenterProbePath(contextReport productionContext) string {
	query := url.Values{}
	addSupportScopeQuery(query, contextReport)
	query.Set("limit", "1")
	return "/v1/command-center/timeline?" + query.Encode()
}

func phase2ProofProbePath(contextReport productionContext) string {
	query := url.Values{}
	addSupportScopeQuery(query, contextReport)
	return "/v1/runtime/phase2/proofs?" + query.Encode()
}

func phase3ProofProbePath(contextReport productionContext) string {
	query := url.Values{}
	addSupportScopeQuery(query, contextReport)
	return "/v1/intelligence/phase3/proofs?" + query.Encode()
}

func phase4ProofProbePath(contextReport productionContext) string {
	query := url.Values{}
	addSupportScopeQuery(query, contextReport)
	return "/v1/enterprise/phase4/proofs?" + query.Encode()
}

func addSupportScopeQuery(query url.Values, contextReport productionContext) {
	if tenant := strings.TrimSpace(contextReport.Config.Spec.TenantID); tenant != "" {
		query.Set("tenant_id", tenant)
	}
	if environment := strings.TrimSpace(contextReport.Config.Spec.Environment); environment != "" {
		query.Set("environment", environment)
	}
	if repo := strings.TrimSpace(contextReport.Config.Spec.Repository); repo != "" {
		query.Set("repo", repo)
	}
}

func (a *App) surfaceHealthCheck(ctx context.Context, client *APIClient, component, path, healthySummary string) healthComponent {
	if err := client.ProbeJSON(ctx, path); err != nil {
		return healthComponent{
			Component:    component,
			CurrentState: healthStateFailing,
			Summary:      fmt.Sprintf("%s probe failed", component),
			ReasonCodes:  []string{"surface_probe_failed"},
			NextAction:   "Verify API reachability, auth scope, and the underlying audit-writer surface before relying on this support projection.",
			OwnerHint:    "platform_operator",
		}
	}
	return healthComponent{
		Component:    component,
		CurrentState: healthStateHealthy,
		Summary:      healthySummary,
		ReasonCodes:  []string{"surface_probe_ok"},
		NextAction:   "No action required.",
		OwnerHint:    "platform_operator",
	}
}

func normalizeProfile(profile, environment string, offline bool) string {
	switch strings.ToLower(strings.TrimSpace(profile)) {
	case profileLocal, profileStaging, profileProduction, profileFederated, profileOffline:
		return strings.ToLower(strings.TrimSpace(profile))
	}
	if offline {
		return profileOffline
	}
	switch strings.ToLower(strings.TrimSpace(environment)) {
	case "prod", "production":
		return profileProduction
	case "stage", "staging":
		return profileStaging
	case "federated", "spoke", "hub":
		return profileFederated
	default:
		return profileLocal
	}
}

func evaluateSupportMatrix(currentVersion, targetVersion string) supportMatrix {
	current := parseReleaseVersion(currentVersion)
	target := parseReleaseVersion(targetVersion)
	matrix := supportMatrix{
		SchemaVersion:  supportMatrixSchemaVersion,
		CurrentVersion: strings.TrimSpace(currentVersion),
		TargetVersion:  strings.TrimSpace(targetVersion),
		CurrentLine:    current.Line,
		TargetLine:     target.Line,
		CurrentState:   "bounded_support_matrix",
	}

	switch {
	case !target.OK:
		matrix.CurrentState = "unsupported_target_version"
		matrix.RollbackSupported = false
		matrix.CompatibilityGuidance = []string{"Target version is not parseable as a bounded ChangeLock release line."}
		matrix.RollbackCautions = []string{"Do not attempt upgrade or rollback planning until the target version string is normalized."}
	case !current.OK:
		matrix.CurrentState = "bounded_current_version"
		matrix.RollbackSupported = true
		matrix.CompatibilityGuidance = []string{"Current CLI version is not a stable release line, so the support matrix remains bounded and advisory."}
		matrix.RollbackCautions = []string{"Treat rollback posture cautiously because the current version does not map cleanly onto a supported release line."}
	case current.Major != target.Major:
		matrix.CurrentState = "unsupported_major_transition"
		matrix.RollbackSupported = false
		matrix.CompatibilityGuidance = []string{"Major version transitions are outside the bounded support matrix in this slice."}
		matrix.RollbackCautions = []string{"Cross-major rollback is not represented as a safe path in the current support matrix baseline."}
	case target.Minor > current.Minor+1:
		matrix.CurrentState = "unsupported_multi_minor_jump"
		matrix.RollbackSupported = false
		matrix.CompatibilityGuidance = []string{"Skipping more than one minor line at a time is outside the bounded support matrix."}
		matrix.RollbackCautions = []string{"Upgrade through the next supported minor line first and rerun readiness after each step."}
	case target.Minor < current.Minor:
		matrix.CurrentState = "rollback_transition"
		matrix.RollbackSupported = true
		matrix.CompatibilityGuidance = []string{"The target line is older than the current line, so this operation is treated as a rollback transition."}
		matrix.RollbackCautions = []string{"Re-validate config schema compatibility and audit/evidence projections before rolling back to an older line."}
	default:
		matrix.CurrentState = "supported_transition"
		matrix.RollbackSupported = true
		matrix.CompatibilityGuidance = []string{"Target version is within the current bounded support matrix and can proceed after local readiness passes."}
		if target.Minor > current.Minor {
			matrix.RollbackCautions = []string{"Review config schema changes and rerun readiness immediately after the upgrade before treating rollback as safe."}
		}
	}
	return matrix
}

func parseReleaseVersion(raw string) parsedReleaseVersion {
	raw = strings.TrimSpace(strings.TrimPrefix(raw, "v"))
	parts := strings.Split(raw, ".")
	if len(parts) < 2 {
		return parsedReleaseVersion{Raw: raw}
	}
	major, majorErr := strconv.Atoi(parts[0])
	minor, minorErr := strconv.Atoi(parts[1])
	if majorErr != nil || minorErr != nil {
		return parsedReleaseVersion{Raw: raw}
	}
	return parsedReleaseVersion{
		Raw:   raw,
		Line:  fmt.Sprintf("v%d.%d", major, minor),
		Major: major,
		Minor: minor,
		OK:    true,
	}
}

func statusForSupportMatrix(matrix supportMatrix) Status {
	switch matrix.CurrentState {
	case "supported_transition":
		return StatusPass
	case "rollback_transition", "bounded_current_version":
		return StatusWarning
	default:
		return StatusFail
	}
}

func supportMatrixSummary(matrix supportMatrix) string {
	switch matrix.CurrentState {
	case "supported_transition":
		return "target version is inside the bounded support matrix"
	case "rollback_transition":
		return "target version is older than current and is treated as a rollback transition"
	case "bounded_current_version":
		return "support matrix is advisory because the current version is not a stable release line"
	case "unsupported_target_version":
		return "target version is not parseable as a bounded ChangeLock release line"
	case "unsupported_major_transition":
		return "target version crosses a major support boundary"
	default:
		return "target version falls outside the bounded support matrix"
	}
}

func statusForRollbackMatrix(matrix supportMatrix) Status {
	if !matrix.RollbackSupported {
		return StatusFail
	}
	if len(matrix.RollbackCautions) > 0 {
		return StatusWarning
	}
	return StatusPass
}

func rollbackSummary(matrix supportMatrix) string {
	if !matrix.RollbackSupported {
		return "rollback posture is not safe in the current bounded support matrix"
	}
	if len(matrix.RollbackCautions) > 0 {
		return "rollback posture remains possible but requires explicit caution"
	}
	return "rollback posture is acceptable inside the current bounded support matrix"
}

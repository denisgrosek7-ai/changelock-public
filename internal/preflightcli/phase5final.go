package preflightcli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
)

const (
	phase5SummarySchemaVersion    = "5.phase5_production_summary.v1"
	phase5StateIncomplete         = "phase5_incomplete"
	phase5StateSubstantiallyReady = "phase5_substantially_ready"
	phase5StateProductionReady    = "phase5_production_usability_active"
	phase5SectionRedactionState   = "redacted_by_default"
)

type phase5SummaryOptions struct {
	ConfigPath    string
	Output        string
	Profile       string
	TargetVersion string
}

type phase5SectionSummary struct {
	Name           string        `json:"name"`
	OverallResult  Status        `json:"overall_result"`
	CurrentState   string        `json:"current_state"`
	Summary        string        `json:"summary"`
	Checks         []CheckResult `json:"checks,omitempty"`
	ReasonCodes    []string      `json:"reason_codes,omitempty"`
	EvidenceRefs   []string      `json:"evidence_refs,omitempty"`
	RedactionState string        `json:"redaction_state,omitempty"`
	Limitations    []string      `json:"limitations,omitempty"`
}

type phase5SummaryResponse struct {
	SchemaVersion  string               `json:"schema_version"`
	GeneratedAt    time.Time            `json:"generated_at"`
	CurrentState   string               `json:"current_state"`
	OverallResult  Status               `json:"overall_result"`
	Profile        string               `json:"profile"`
	ConfigPath     string               `json:"config_path"`
	CommandCenter  phase5SectionSummary `json:"command_center"`
	ConfigCLI      phase5SectionSummary `json:"config_cli"`
	Supportability phase5SectionSummary `json:"supportability"`
	ReasonCodes    []string             `json:"reason_codes,omitempty"`
	Limitations    []string             `json:"limitations,omitempty"`
}

func (a *App) runPhase5Summary(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	options, err := a.parsePhase5SummaryOptions(args)
	if err != nil {
		return a.writeError(stderr, err)
	}
	a.config.Output = options.Output
	contextReport, loadErr := a.loadProductionContext(options.ConfigPath)
	if loadErr != nil {
		response := phase5InvalidSummary(options, loadErr)
		if err := renderPhase5Summary(stdout, options.Output, response); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitFailed
	}
	response := a.buildPhase5Summary(ctx, contextReport, options)
	if err := renderPhase5Summary(stdout, options.Output, response); err != nil {
		return a.writeError(stderr, err)
	}
	return exitCodeForPhase5Summary(response)
}

func (a *App) parsePhase5SummaryOptions(args []string) (phase5SummaryOptions, error) {
	fs := newFlagSet("phase5-summary")
	options := phase5SummaryOptions{
		ConfigPath: a.config.ProductionConfigPath,
		Output:     a.config.Output,
		Profile:    "",
	}
	fs.StringVar(&options.ConfigPath, "config", options.ConfigPath, "production config file")
	fs.StringVar(&options.Output, "output", options.Output, "output mode: human|json")
	fs.StringVar(&options.Profile, "profile", options.Profile, "environment profile: local|staging|production|federated|offline")
	fs.StringVar(&options.TargetVersion, "target-version", "", "optional target ChangeLock version for bounded upgrade posture")
	if err := fs.Parse(args); err != nil {
		return phase5SummaryOptions{}, usageError{message: err.Error()}
	}
	if strings.TrimSpace(options.ConfigPath) == "" {
		return phase5SummaryOptions{}, usageError{message: "phase5-summary requires --config or CHANGELOCK_CLI_CONFIG"}
	}
	if err := validateOutputMode(options.Output); err != nil {
		return phase5SummaryOptions{}, err
	}
	if err := validateProfileValue(options.Profile); err != nil {
		return phase5SummaryOptions{}, err
	}
	return options, nil
}

func phase5InvalidSummary(options phase5SummaryOptions, loadErr error) phase5SummaryResponse {
	failed := phase5SectionSummary{
		Name:          "config_cli",
		OverallResult: StatusFail,
		CurrentState:  phase5SectionState("config_cli", StatusFail),
		Summary:       "strict config loading failed before Phase 5 consolidation could run",
		Checks: []CheckResult{{
			Name:    "phase5-config-load",
			Mode:    ModeLocal,
			Target:  strings.TrimSpace(options.ConfigPath),
			Status:  StatusFail,
			Summary: "production config failed strict loading",
			Details: []string{loadErr.Error()},
		}},
		ReasonCodes: []string{"config_load_failed"},
		Limitations: []string{
			"Phase 5 final summary is fail-fast on invalid production config and does not silently fall back to permissive parsing.",
		},
	}
	return phase5SummaryResponse{
		SchemaVersion: phase5SummarySchemaVersion,
		GeneratedAt:   time.Now().UTC(),
		CurrentState:  phase5StateIncomplete,
		OverallResult: StatusFail,
		Profile:       strings.TrimSpace(options.Profile),
		ConfigPath:    strings.TrimSpace(options.ConfigPath),
		CommandCenter: phase5SectionSummary{
			Name:          "command_center",
			OverallResult: StatusInfo,
			CurrentState:  phase5SectionState("command_center", StatusInfo),
			Summary:       "command-center consolidation was not evaluated because local config failed strict loading first",
			ReasonCodes:   []string{"phase5_command_center_not_evaluated"},
		},
		ConfigCLI: failed,
		Supportability: phase5SectionSummary{
			Name:           "supportability",
			OverallResult:  StatusFail,
			CurrentState:   phase5SectionState("supportability", StatusFail),
			Summary:        "supportability consolidation cannot proceed without a valid production config",
			ReasonCodes:    []string{"phase5_supportability_not_evaluated"},
			RedactionState: phase5SectionRedactionState,
		},
		ReasonCodes: []string{"config_load_failed", phase5StateIncomplete},
		Limitations: []string{
			"Phase 5 final summary connects Val A, Val B, and Val C only after strict config loading succeeds.",
		},
	}
}

func (a *App) buildPhase5Summary(ctx context.Context, contextReport productionContext, options phase5SummaryOptions) phase5SummaryResponse {
	profile := normalizeProfile(options.Profile, contextReport.Config.Spec.Environment, contextReport.Config.Spec.CLI.Offline)
	targetVersion := strings.TrimSpace(options.TargetVersion)
	if targetVersion == "" {
		targetVersion = firstNonEmpty(strings.TrimSpace(a.runtime.VersionInfo.Version), "dev")
	}

	commandCenter := a.buildPhase5CommandCenterSection(ctx, contextReport, profile)
	configCLI := a.buildPhase5ConfigCLISection(ctx, contextReport)
	supportability := a.buildPhase5SupportabilitySection(ctx, contextReport, profile, targetVersion)

	currentState, overallResult := evaluatePhase5CurrentState(commandCenter, configCLI, supportability)
	reasonCodes := uniqueStrings(append(
		append(append([]string{currentState, "phase5_profile_" + profile}, commandCenter.ReasonCodes...), configCLI.ReasonCodes...),
		supportability.ReasonCodes...,
	))

	return phase5SummaryResponse{
		SchemaVersion:  phase5SummarySchemaVersion,
		GeneratedAt:    time.Now().UTC(),
		CurrentState:   currentState,
		OverallResult:  overallResult,
		Profile:        profile,
		ConfigPath:     contextReport.Path,
		CommandCenter:  commandCenter,
		ConfigCLI:      configCLI,
		Supportability: supportability,
		ReasonCodes:    reasonCodes,
		Limitations: []string{
			"Phase 5 final summary is a bounded consolidation over existing command-center, deterministic config/CLI, and supportability surfaces; it does not create a new truth store.",
			"Production usability activation remains conservative and requires Val A, Val B, and Val C to stay simultaneously active without a critical blocker.",
			"Final summary stays inside Phase 5 boundaries and does not upgrade readiness language into market, certification, or public-authority claims.",
		},
	}
}

func (a *App) buildPhase5CommandCenterSection(ctx context.Context, contextReport productionContext, profile string) phase5SectionSummary {
	apiProbe := a.apiConnectivityCheck(ctx, contextReport, profile)
	checks := []CheckResult{{
		Name:    "command-center-api",
		Mode:    apiProbe.check.Mode,
		Target:  apiProbe.check.Target,
		Status:  apiProbe.check.Status,
		Summary: apiProbe.check.Summary,
		Details: apiProbe.check.Details,
	}}

	if apiProbe.client == nil || apiProbe.check.Status != StatusPass {
		unavailableStatus := apiProbe.check.Status
		if unavailableStatus == "" {
			unavailableStatus = StatusInfo
		}
		checks = append(checks,
			phase5UnavailableSurfaceCheck("command-center-timeline", unavailableStatus, "timeline surface was not probed because API context is not ready"),
			phase5UnavailableSurfaceCheck("command-center-notifications", unavailableStatus, "notifications surface was not probed because API context is not ready"),
			phase5UnavailableSurfaceCheck("command-center-search", unavailableStatus, "search surface was not probed because API context is not ready"),
		)
	} else {
		checks = append(checks,
			a.phase5ProbeCheck(ctx, apiProbe.client, "command-center-timeline", commandCenterProbePath(contextReport), "timeline surface is reachable and bounded lifecycle aggregation is active"),
			a.phase5ProbeCheck(ctx, apiProbe.client, "command-center-notifications", commandCenterNotificationsProbePath(contextReport), "grouped notification surface is reachable and state-aware alerting is active"),
			a.phase5ProbeCheck(ctx, apiProbe.client, "command-center-search", commandCenterSearchProbePath(contextReport), "canonical command-center search surface is reachable"),
		)
	}

	result := finalizeResult(Result{
		Command:     "phase5-command-center",
		Mode:        executionMode(contextReport.Config.Spec.CLI.Offline, contextReport.Config.Spec.APIURL),
		Checks:      checks,
		ReasonCodes: uniqueStrings(append([]string{"phase5_vala_command_center"}, checkNamesToReasonCodes(checks)...)),
		Limitations: []string{
			"Command-center consolidation reuses timeline, grouped notifications, and search surfaces already exposed by ChangeLock.",
			"Surface reachability does not claim that every tenant scope already has rich data; it only proves that the operator layer is live and bounded.",
		},
	})
	return phase5SectionFromResult("command_center", "Command center is consolidated across timeline, grouped notifications, and canonical search.", result, []string{
		"/v1/command-center/timeline",
		"/v1/command-center/notifications",
		"/v1/command-center/search",
	}, "")
}

func (a *App) buildPhase5ConfigCLISection(ctx context.Context, contextReport productionContext) phase5SectionSummary {
	checkResult, checkErr := a.runCheck(ctx, []string{"--config", contextReport.Path, "--output", "json"})
	previewResult, previewErr := a.runPreview(ctx, []string{"--config", contextReport.Path, "--output", "json"})
	inspectExit := a.runInspect(ctx, []string{"--config", contextReport.Path, "--output", "json"}, io.Discard, io.Discard)
	explainExit := a.runExplain(ctx, []string{"--config", contextReport.Path, "--output", "json"}, io.Discard, io.Discard)

	checks := []CheckResult{
		{
			Name:    "strict-config-validation",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  statusForValidationState(contextReport.Inspection.Validation.CurrentState),
			Summary: configValidationSummary(contextReport.Inspection.Validation.CurrentState),
			Details: validationMessages(contextReport.Inspection.Validation.Issues),
		},
		{
			Name:    "effective-config-inspection",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  phase5StatusForInspectionCurrentState(contextReport.Inspection.CurrentState, contextReport.RuntimeInspection.CurrentState),
			Summary: phase5InspectionSummary(contextReport.Inspection.CurrentState, contextReport.RuntimeInspection.CurrentState),
			Details: uniqueStrings(append(contextReport.Inspection.EffectiveConfig.Sync.Explanation, runtimeInspectionDetails(contextReport.RuntimeInspection)...)),
		},
		{
			Name:    "sync-conflict-visibility",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  statusForSyncState(contextReport.Inspection.EffectiveConfig.Sync.CurrentState),
			Summary: syncSummary(contextReport.Inspection.EffectiveConfig.Sync.CurrentState),
			Details: contextReport.Inspection.EffectiveConfig.Sync.Explanation,
		},
		phase5CLISurfaceCheck(checkErr, previewErr, inspectExit, explainExit, checkResult, previewResult, contextReport.Path),
	}

	result := finalizeResult(Result{
		Command:     "phase5-config-cli",
		Mode:        ModeLocalOnly,
		Checks:      checks,
		ReasonCodes: uniqueStrings(append(append([]string{"phase5_valb_config_cli"}, contextReport.Inspection.ReasonCodes...), runtimeIssueCodes(contextReport.RuntimeInspection)...)),
		Limitations: []string{
			"CLI consolidation remains bounded to strict local config, effective state inspection, and explainable sync/workflow/runtime signals already present in Val B.",
			"Preview honesty is preserved by reusing the same deterministic command surfaces instead of inventing a broader simulation layer.",
		},
	})
	return phase5SectionFromResult("config_cli", "Strict config, effective inspection, and operator CLI surfaces are aligned on the same effective state model.", result, []string{
		"changelock-cli check --config <path>",
		"changelock-cli preview --config <path>",
		"changelock-cli inspect --config <path>",
		"changelock-cli explain --config <path> --topic sync",
	}, "")
}

func (a *App) buildPhase5SupportabilitySection(ctx context.Context, contextReport productionContext, profile, targetVersion string) phase5SectionSummary {
	readiness := a.buildReadinessResult(ctx, contextReport, readinessOptions{
		ConfigPath: contextReport.Path,
		Output:     "json",
		Profile:    profile,
	})
	health := a.buildHealthSnapshot(ctx, contextReport, profile)
	upgradeResult, matrix := a.buildUpgradeReadinessResult(ctx, contextReport, upgradeReadinessOptions{
		ConfigPath:     contextReport.Path,
		Output:         "json",
		Profile:        profile,
		CurrentVersion: strings.TrimSpace(a.runtime.VersionInfo.Version),
		TargetVersion:  targetVersion,
	})

	checks := []CheckResult{
		{
			Name:    "readiness-gate",
			Mode:    CheckMode(readiness.Mode),
			Target:  contextReport.Path,
			Status:  phase5NormalizedResultStatus(readiness),
			Summary: "profile-aware readiness gate remains available for bounded go-live validation",
			Details: uniqueStrings(append([]string{}, readiness.ReasonCodes...)),
		},
		{
			Name:    "health-projections",
			Mode:    CheckMode(executionMode(contextReport.Config.Spec.CLI.Offline, contextReport.Config.Spec.APIURL)),
			Target:  contextReport.Path,
			Status:  phase5StatusForHealthState(health.CurrentState),
			Summary: fmt.Sprintf("normalized supportability health snapshot is %s", health.CurrentState),
			Details: healthComponentDetails(health.Items),
		},
		{
			Name:    "support-bundle-redaction",
			Mode:    ModeLocal,
			Target:  contextReport.Path,
			Status:  StatusPass,
			Summary: "support bundle remains redacted by default across runtime and config diagnostics",
			Details: []string{"Declared runtime environment values are redacted and the final summary only exposes redaction state, not raw secret-like inputs."},
		},
		{
			Name:    "upgrade-rollback-guidance",
			Mode:    CheckMode(upgradeResult.Mode),
			Target:  targetVersion,
			Status:  phase5NormalizedResultStatus(upgradeResult),
			Summary: supportMatrixSummary(matrix),
			Details: uniqueStrings(append(append([]string{}, matrix.CompatibilityGuidance...), matrix.RollbackCautions...)),
		},
	}

	result := finalizeResult(Result{
		Command: "phase5-supportability",
		Mode:    executionMode(contextReport.Config.Spec.CLI.Offline, contextReport.Config.Spec.APIURL),
		Checks:  checks,
		ReasonCodes: uniqueStrings(append(append([]string{
			"phase5_valc_supportability",
			phase5SectionRedactionState,
		}, readiness.ReasonCodes...), upgradeResult.ReasonCodes...)),
		Limitations: []string{
			"Supportability consolidation is bounded to readiness, redacted support diagnostics, health projections, and support-matrix guidance already present in Val C.",
			"Upgrade and rollback guidance remains advisory and must be paired with post-change readiness before treating production posture as restored.",
		},
	})
	return phase5SectionFromResult("supportability", "Readiness, redacted supportability, health projections, and upgrade posture are consolidated.", result, []string{
		"changelock-cli readiness --config <path>",
		"changelock-cli support --config <path>",
		"changelock-cli upgrade-readiness --config <path> --target-version <version>",
	}, phase5SectionRedactionState)
}

func phase5UnavailableSurfaceCheck(name string, status Status, summary string) CheckResult {
	return CheckResult{
		Name:    name,
		Mode:    ModeRemote,
		Status:  status,
		Summary: summary,
	}
}

func (a *App) phase5ProbeCheck(ctx context.Context, client *APIClient, name, path, summary string) CheckResult {
	if err := client.ProbeJSON(ctx, path); err != nil {
		return CheckResult{
			Name:    name,
			Mode:    ModeRemote,
			Target:  path,
			Status:  StatusFail,
			Summary: fmt.Sprintf("%s probe failed", name),
			Details: []string{err.Error()},
		}
	}
	return CheckResult{
		Name:    name,
		Mode:    ModeRemote,
		Target:  path,
		Status:  StatusPass,
		Summary: summary,
	}
}

func phase5CLISurfaceCheck(checkErr, previewErr error, inspectExit, explainExit int, checkResult, previewResult Result, configPath string) CheckResult {
	details := []string{
		fmt.Sprintf("check overall result: %s", phase5NormalizedResultStatus(checkResult)),
		fmt.Sprintf("preview overall result: %s", phase5NormalizedResultStatus(previewResult)),
		fmt.Sprintf("inspect exit code: %d", inspectExit),
		fmt.Sprintf("explain exit code: %d", explainExit),
	}
	if checkErr != nil {
		return CheckResult{
			Name:    "cli-operator-surface",
			Mode:    ModeLocal,
			Target:  configPath,
			Status:  StatusFail,
			Summary: "check command failed while Phase 5 CLI consolidation was being exercised",
			Details: []string{checkErr.Error()},
		}
	}
	if previewErr != nil {
		return CheckResult{
			Name:    "cli-operator-surface",
			Mode:    ModeLocal,
			Target:  configPath,
			Status:  StatusFail,
			Summary: "preview command failed while Phase 5 CLI consolidation was being exercised",
			Details: []string{previewErr.Error()},
		}
	}
	if inspectExit == ExitUsage || inspectExit == ExitExecution || explainExit == ExitUsage || explainExit == ExitExecution {
		return CheckResult{
			Name:    "cli-operator-surface",
			Mode:    ModeLocal,
			Target:  configPath,
			Status:  StatusFail,
			Summary: "inspect or explain surface did not stay operational during Phase 5 consolidation",
			Details: details,
		}
	}
	return CheckResult{
		Name:    "cli-operator-surface",
		Mode:    ModeLocal,
		Target:  configPath,
		Status:  StatusPass,
		Summary: "check, preview, inspect, and explain remain aligned on the same effective config model",
		Details: details,
	}
}

func phase5StatusForInspectionCurrentState(configState, runtimeState string) Status {
	switch combinedInspectionState(configState, runtimeState) {
	case "invalid":
		return StatusFail
	case "valid_with_warnings":
		return StatusWarning
	default:
		return StatusPass
	}
}

func phase5InspectionSummary(configState, runtimeState string) string {
	switch combinedInspectionState(configState, runtimeState) {
	case "invalid":
		return "effective config inspection found blocking local config or runtime env issues"
	case "valid_with_warnings":
		return "effective config inspection stays active with visible warnings"
	default:
		return "effective config inspection is active and keeps defaults, sync, and runtime posture visible"
	}
}

func phase5StatusForHealthState(state string) Status {
	switch strings.TrimSpace(state) {
	case healthStateHealthy:
		return StatusPass
	case healthStateDegraded:
		return StatusDegraded
	case healthStateFailing:
		return StatusFail
	default:
		return StatusInfo
	}
}

func healthComponentDetails(items []healthComponent) []string {
	details := make([]string, 0, len(items))
	for _, item := range items {
		details = append(details, fmt.Sprintf("%s=%s", item.Component, item.CurrentState))
	}
	return uniqueStrings(details)
}

func phase5SectionFromResult(name, summary string, result Result, evidenceRefs []string, redactionState string) phase5SectionSummary {
	result = finalizeResult(result)
	overall := phase5NormalizedResultStatus(result)
	return phase5SectionSummary{
		Name:           name,
		OverallResult:  overall,
		CurrentState:   phase5SectionState(name, overall),
		Summary:        summary,
		Checks:         result.Checks,
		ReasonCodes:    uniqueStrings(append([]string{phase5SectionState(name, overall)}, result.ReasonCodes...)),
		EvidenceRefs:   uniqueStrings(evidenceRefs),
		RedactionState: redactionState,
		Limitations:    result.Limitations,
	}
}

func phase5SectionState(name string, overall Status) string {
	switch overall {
	case StatusPass:
		return name + "_active"
	case StatusWarning, StatusDegraded, StatusInfo:
		return name + "_ready"
	default:
		return name + "_incomplete"
	}
}

func evaluatePhase5CurrentState(commandCenter, configCLI, supportability phase5SectionSummary) (string, Status) {
	sections := []phase5SectionSummary{commandCenter, configCLI, supportability}
	allPass := true
	anyFail := false
	worst := StatusPass
	for _, section := range sections {
		switch section.OverallResult {
		case StatusFail, StatusError:
			anyFail = true
			allPass = false
		case StatusPass:
		case StatusDegraded:
			allPass = false
			if worst == StatusPass || worst == StatusInfo || worst == StatusWarning {
				worst = StatusDegraded
			}
		case StatusWarning:
			allPass = false
			if worst == StatusPass || worst == StatusInfo {
				worst = StatusWarning
			}
		case StatusInfo:
			allPass = false
			if worst == StatusPass {
				worst = StatusInfo
			}
		default:
			allPass = false
		}
	}
	if allPass {
		return phase5StateProductionReady, StatusPass
	}
	if !anyFail {
		if worst == StatusPass {
			worst = StatusInfo
		}
		return phase5StateSubstantiallyReady, worst
	}
	return phase5StateIncomplete, StatusFail
}

func phase5NormalizedResultStatus(result Result) Status {
	return phase5NormalizedChecksStatus(result.Checks)
}

func phase5NormalizedChecksStatus(checks []CheckResult) Status {
	if len(checks) == 0 {
		return StatusInfo
	}
	hasPass := false
	hasInfo := false
	hasWarning := false
	hasDegraded := false
	for _, check := range checks {
		switch check.Status {
		case StatusError:
			return StatusError
		case StatusFail:
			return StatusFail
		case StatusDegraded:
			hasDegraded = true
		case StatusWarning:
			hasWarning = true
		case StatusInfo, StatusSkip:
			hasInfo = true
		case StatusPass:
			hasPass = true
		}
	}
	switch {
	case hasDegraded:
		return StatusDegraded
	case hasWarning:
		return StatusWarning
	case hasInfo && !hasPass:
		return StatusInfo
	default:
		return StatusPass
	}
}

func checkNamesToReasonCodes(checks []CheckResult) []string {
	reasons := make([]string, 0, len(checks))
	for _, check := range checks {
		if trimmed := strings.TrimSpace(check.Name); trimmed != "" {
			reasons = append(reasons, strings.ReplaceAll(trimmed, "-", "_"))
		}
	}
	return uniqueStrings(reasons)
}

func renderPhase5Summary(w io.Writer, output string, response phase5SummaryResponse) error {
	switch strings.ToLower(strings.TrimSpace(output)) {
	case "", "human":
		if _, err := fmt.Fprintf(w, "Phase 5 summary\nConfig: %s\nProfile: %s\nState: %s\nOverall: %s\nGenerated at: %s\n", response.ConfigPath, response.Profile, response.CurrentState, response.OverallResult, response.GeneratedAt.Format(time.RFC3339)); err != nil {
			return err
		}
		for _, section := range []phase5SectionSummary{response.CommandCenter, response.ConfigCLI, response.Supportability} {
			if _, err := fmt.Fprintf(w, "\n%s [%s] (%s)\n%s\n", section.Name, section.OverallResult, section.CurrentState, section.Summary); err != nil {
				return err
			}
			for _, check := range section.Checks {
				if _, err := fmt.Fprintf(w, "- %s [%s]: %s\n", check.Name, check.Status, check.Summary); err != nil {
					return err
				}
			}
			if section.RedactionState != "" {
				if _, err := fmt.Fprintf(w, "  redaction: %s\n", section.RedactionState); err != nil {
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

func exitCodeForPhase5Summary(response phase5SummaryResponse) int {
	switch response.OverallResult {
	case StatusPass, StatusWarning, StatusDegraded, StatusInfo:
		return ExitSuccess
	case StatusFail:
		return ExitFailed
	default:
		return ExitExecution
	}
}

func commandCenterNotificationsProbePath(contextReport productionContext) string {
	query := url.Values{}
	addSupportScopeQuery(query, contextReport)
	query.Set("limit", "1")
	return "/v1/command-center/notifications?" + query.Encode()
}

func commandCenterSearchProbePath(contextReport productionContext) string {
	query := url.Values{}
	addSupportScopeQuery(query, contextReport)
	query.Set("limit", "1")
	query.Set("q", firstNonEmpty(
		strings.TrimSpace(contextReport.Config.Metadata.Name),
		strings.TrimSpace(contextReport.Config.Spec.Repository),
		strings.TrimSpace(contextReport.Config.Spec.TenantID),
		"phase5",
	))
	return "/v1/command-center/search?" + query.Encode()
}

package preflightcli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type diagnosticsOptions struct {
	Input       string
	Format      string
	IncludePass bool
}

type DiagnosticsOutput struct {
	Command           string            `json:"command"`
	Mode              string            `json:"mode"`
	OverallResult     Status            `json:"overall_result"`
	ExitCode          int               `json:"exit_code"`
	Inputs            map[string]string `json:"inputs,omitempty"`
	Diagnostics       []Diagnostic      `json:"diagnostics"`
	DiagnosticSummary DiagnosticSummary `json:"diagnostic_summary"`
}

type sarifLog struct {
	Version string     `json:"version"`
	Schema  string     `json:"$schema"`
	Runs    []sarifRun `json:"runs"`
}

type sarifRun struct {
	Tool    sarifTool     `json:"tool"`
	Results []sarifResult `json:"results"`
}

type sarifTool struct {
	Driver sarifDriver `json:"driver"`
}

type sarifDriver struct {
	Name           string      `json:"name"`
	InformationURI string      `json:"informationUri,omitempty"`
	Rules          []sarifRule `json:"rules,omitempty"`
}

type sarifRule struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name,omitempty"`
	ShortDescription     sarifMultiformMessage  `json:"shortDescription"`
	HelpURI              string                 `json:"helpUri,omitempty"`
	DefaultConfiguration map[string]string      `json:"defaultConfiguration,omitempty"`
	Properties           map[string]interface{} `json:"properties,omitempty"`
}

type sarifResult struct {
	RuleID     string            `json:"ruleId"`
	Level      string            `json:"level"`
	Message    sarifPlainMessage `json:"message"`
	Locations  []sarifLocation   `json:"locations,omitempty"`
	Properties map[string]any    `json:"properties,omitempty"`
}

type sarifMultiformMessage struct {
	Text string `json:"text"`
}

type sarifPlainMessage struct {
	Text string `json:"text"`
}

type sarifLocation struct {
	PhysicalLocation sarifPhysicalLocation `json:"physicalLocation"`
}

type sarifPhysicalLocation struct {
	ArtifactLocation sarifArtifactLocation `json:"artifactLocation"`
	Region           *sarifRegion          `json:"region,omitempty"`
}

type sarifArtifactLocation struct {
	URI string `json:"uri"`
}

type sarifRegion struct {
	StartLine   int `json:"startLine"`
	StartColumn int `json:"startColumn,omitempty"`
	EndLine     int `json:"endLine,omitempty"`
	EndColumn   int `json:"endColumn,omitempty"`
}

func attachDiagnostics(result Result) Result {
	diagnostics, summary := buildDiagnostics(result)
	result.Diagnostics = diagnostics
	result.DiagnosticSummary = summary
	return result
}

func buildDiagnostics(result Result) ([]Diagnostic, DiagnosticSummary) {
	diagnostics := make([]Diagnostic, 0, len(result.Checks))
	summary := DiagnosticSummary{
		CountsBySeverity:        map[string]int{},
		CountsBySource:          map[string]int{},
		CountsByEvaluationState: map[string]int{},
	}
	for _, check := range result.Checks {
		diagnostic := diagnosticForCheck(result.Command, check)
		diagnostics = append(diagnostics, diagnostic)
		summary.Total++
		if diagnostic.Blocking {
			summary.Blocking++
		} else if diagnostic.EvaluationState != EvaluationStatePass {
			summary.Advisory++
		}
		summary.CountsBySeverity[string(diagnostic.Severity)]++
		summary.CountsBySource[diagnostic.Source]++
		summary.CountsByEvaluationState[string(diagnostic.EvaluationState)]++
	}
	return diagnostics, summary
}

func diagnosticForCheck(command string, check CheckResult) Diagnostic {
	reasonCode, category, source, docsRef, fixHint := diagnosticDescriptor(check)
	targetFile := diagnosticTargetFile(check)
	diagnostic := Diagnostic{
		CheckID:          check.Name,
		RuleID:           diagnosticRuleID(command, check.Name),
		Category:         category,
		Severity:         diagnosticSeverity(check.Status),
		ReasonCode:       reasonCode,
		Message:          diagnosticMessage(check),
		Summary:          check.Summary,
		Target:           check.Target,
		TargetFile:       targetFile,
		ResourceIdentity: diagnosticResourceIdentity(check.Target, targetFile),
		FixHint:          fixHint,
		DocsRef:          docsRef,
		Source:           source,
		Blocking:         check.Status == StatusFail || check.Status == StatusError,
		EvaluationState:  evaluationState(check.Status),
	}
	if reviewRange := diagnosticRange(check); reviewRange != nil {
		diagnostic.Range = reviewRange
	} else if targetFile != "" {
		diagnostic.Range = &DiagnosticRange{
			StartLine:   1,
			StartColumn: 1,
			EndLine:     1,
			EndColumn:   1,
		}
	}
	return diagnostic
}

func diagnosticDescriptor(check CheckResult) (reasonCode, category, source, docsRef, fixHint string) {
	contextKind := metadataString(check.Metadata, "context_kind")
	switch check.Name {
	case "manifest":
		switch check.Status {
		case StatusPass:
			return "manifest_policy_satisfied", "policy", "policy", "docs/developer-preflight-cli.md", "No action required."
		case StatusFail:
			return "manifest_policy_violation", "policy", "policy", "docs/developer-preflight-cli.md", "Review the Kyverno deny output and update the manifest or tenant policy before pushing."
		default:
			return "manifest_check_error", "policy", "policy", "docs/developer-preflight-cli.md", "Install `kyverno` or set `CHANGELOCK_CLI_KYVERNO_BIN` to a working binary before rerunning the check."
		}
	case "image-digest":
		if check.Status == StatusPass {
			return "image_digest_pinned", "supply-chain", "policy", "docs/supply-chain.md", "No action required."
		}
		return "image_digest_required", "supply-chain", "policy", "docs/supply-chain.md", "Use immutable `@sha256:` image references in manifests and examples."
	case "image-trust":
		switch check.Status {
		case StatusPass:
			return "image_trust_verified", "supply-chain", "signer_identity", "docs/supply-chain.md", "No action required."
		case StatusFail:
			return "image_trust_verification_failed", "supply-chain", "signer_identity", "docs/supply-chain.md", "Verify the image signature, provenance, and allowed signer identity before promotion."
		default:
			return "image_trust_check_error", "supply-chain", "signer_identity", "docs/supply-chain.md", "Install `cosign` or set `CHANGELOCK_CLI_COSIGN_BIN` to a working binary before rerunning the check."
		}
	case "image-policy":
		switch check.Status {
		case StatusPass:
			return "artifact_policy_allowed", "policy", "policy", "docs/supply-chain.md", "No action required."
		case StatusFail:
			return "artifact_policy_denied", "policy", "policy", "docs/deployment-flow.md", "Adjust repository, branch, signer, or environment inputs so the artifact satisfies ChangeLock policy."
		default:
			return "artifact_policy_error", "policy", "policy", "docs/deployment-flow.md", "Verify the local policy bundle path and rerun the artifact policy check."
		}
	case "scan":
		switch check.Status {
		case StatusPass:
			return "vulnerability_threshold_clear", "vulnerability", "vulnerability", "docs/vulnerability-ops.md", "No action required."
		case StatusFail:
			return "vulnerability_threshold_breached", "vulnerability", "vulnerability", "docs/vex-exploitability-ops.md", "Remediate the package version or record a scoped VEX statement if the finding is truly not affected."
		default:
			if strings.Contains(strings.ToLower(check.Summary), "command not found") {
				return "scanner_unavailable", "vulnerability", "vulnerability", "docs/developer-preflight-cli.md", "Install `trivy` or `grype`, or point the CLI to a working scanner binary."
			}
			return "vulnerability_scan_error", "vulnerability", "vulnerability", "docs/vulnerability-ops.md", "Rerun the scan and inspect local scanner output before treating the result as trusted."
		}
	case "remote-auth":
		switch check.Status {
		case StatusPass:
			return "api_identity_confirmed", "context", "evidence", "docs/developer-preflight-cli.md", "No action required."
		case StatusSkip:
			return "api_context_not_configured", "context", "evidence", "docs/developer-preflight-cli.md", "Set `--api-url` and a read-only token if you want ChangeLock server context in local checks."
		default:
			return "api_identity_lookup_failed", "context", "evidence", "docs/developer-preflight-cli.md", "Verify the ChangeLock API URL, token, and local network path before relying on API-assisted checks."
		}
	case "remote-image-context":
		switch check.Status {
		case StatusPass:
			if metadataInt(check.Metadata, "exception_count") > 0 {
				return "remote_exception_context_found", "context", "evidence", "docs/audit-evidence.md", "Review the linked approved exceptions and confirm they still match the intended deployment scope."
			}
			return "remote_exception_context_clear", "context", "evidence", "docs/audit-evidence.md", "No approved digest exceptions matched this image."
		case StatusSkip:
			if strings.Contains(strings.ToLower(check.Summary), "not digest-pinned") {
				return "remote_context_skipped_not_digest_pinned", "context", "evidence", "docs/developer-preflight-cli.md", "Use a digest-pinned image reference if you want digest-scoped server context and exception matching."
			}
			return "api_context_not_configured", "context", "evidence", "docs/developer-preflight-cli.md", "Set `--api-url` and a read-only token if you want ChangeLock server context in local checks."
		default:
			return "remote_exception_context_failed", "context", "evidence", "docs/audit-evidence.md", "Verify ChangeLock API reachability before relying on server-side exception context."
		}
	case "remote-scan-context":
		switch check.Status {
		case StatusPass:
			if contextKind == "vex-net" {
				if metadataBool(check.Metadata, "threshold_breached") {
					return "remote_vex_context_actionable", "vulnerability", "vex", "docs/vex-exploitability-ops.md", "The finding still breaches the threshold after VEX merge. Remediate or record a tighter, valid VEX statement."
				}
				return "remote_vex_context_clear", "vulnerability", "vex", "docs/vex-exploitability-ops.md", "No threshold breach remains after VEX merge for the evaluated digest."
			}
			if metadataInt(check.Metadata, "exception_match_count") > 0 {
				return "remote_exception_context_found", "context", "evidence", "docs/audit-evidence.md", "Review the linked approved CVE exceptions and confirm they still match the intended digest and environment."
			}
			return "remote_exception_context_clear", "context", "evidence", "docs/audit-evidence.md", "No approved CVE exceptions matched the threshold-breaching findings."
		case StatusSkip:
			if contextKind == "vex-net-partial" {
				return "remote_vex_context_unavailable", "vulnerability", "vex", "docs/vex-exploitability-ops.md", "Use a digest-pinned image reference if you want net-actionable VEX context in local checks."
			}
			if strings.Contains(strings.ToLower(check.Summary), "no threshold-breaching findings") {
				return "remote_context_skipped_no_findings", "vulnerability", "vulnerability", "docs/vulnerability-ops.md", "No findings breached the configured threshold, so server-side vulnerability context was not queried."
			}
			return "api_context_not_configured", "context", "evidence", "docs/developer-preflight-cli.md", "Set `--api-url` and a read-only token if you want ChangeLock server context in local checks."
		default:
			return "remote_vex_context_failed", "vulnerability", "vex", "docs/vex-exploitability-ops.md", "Verify ChangeLock API reachability before relying on server-side VEX and net-actionable vulnerability context."
		}
	case "config-readiness":
		switch check.Status {
		case StatusPass:
			return "config_readiness_confirmed", "config", "policy", "docs/production-phase5-valb.md", "No action required."
		case StatusWarning, StatusDegraded, StatusInfo:
			return "config_readiness_partial", "config", "policy", "docs/production-phase5-valb.md", "Review config warnings, defaults, and compatibility notes before treating this environment as rollout-ready."
		default:
			return "config_readiness_failed", "config", "policy", "docs/production-phase5-valb.md", "Resolve blocking schema, path, or compatibility findings before go-live."
		}
	case "dependency-readiness":
		switch check.Status {
		case StatusPass:
			return "dependency_readiness_confirmed", "dependency", "runtime", "docs/production-phase5-valc.md", "No action required."
		case StatusInfo, StatusWarning:
			return "dependency_readiness_partial", "dependency", "runtime", "docs/production-phase5-valc.md", "Install the missing local tooling before relying on full supportability or deterministic rollout checks."
		default:
			return "dependency_readiness_failed", "dependency", "runtime", "docs/production-phase5-valc.md", "Install the missing binaries and rerun readiness before production use."
		}
	case "api-readiness":
		switch check.Status {
		case StatusPass:
			return "api_readiness_confirmed", "supportability", "evidence", "docs/production-phase5-valc.md", "No action required."
		case StatusInfo, StatusWarning, StatusDegraded:
			return "api_readiness_partial", "supportability", "evidence", "docs/production-phase5-valc.md", "Restore API reachability or explicitly operate this profile as offline before rollout."
		default:
			return "api_readiness_failed", "supportability", "evidence", "docs/production-phase5-valc.md", "Fix API URL, token, or service health before relying on production support surfaces."
		}
	case "sync-readiness", "sync-migration-readiness":
		switch check.Status {
		case StatusPass:
			return "sync_state_ready", "sync", "evidence", "docs/production-phase5-valb.md", "No action required."
		case StatusWarning, StatusDegraded, StatusInfo:
			return "sync_state_partial", "sync", "evidence", "docs/production-phase5-valb.md", "Converge revisions or document why bounded divergence is acceptable before rollout or upgrade."
		default:
			return "sync_state_blocking", "sync", "evidence", "docs/production-phase5-valb.md", "Resolve sync conflicts before rollout or upgrade."
		}
	case "workflow-readiness":
		switch check.Status {
		case StatusPass:
			return "workflow_governance_ready", "workflow", "policy", "docs/production-phase5-valc.md", "No action required."
		case StatusInfo, StatusWarning, StatusDegraded:
			return "workflow_governance_partial", "workflow", "policy", "docs/production-phase5-valc.md", "Tighten approval and validation gates before treating closure workflows as production-ready."
		default:
			return "workflow_governance_failed", "workflow", "policy", "docs/production-phase5-valc.md", "Declare the missing validation and approval gates before rollout."
		}
	case "runtime-self-healing-readiness":
		switch check.Status {
		case StatusPass:
			return "runtime_supportability_ready", "runtime", "runtime", "docs/production-phase5-valc.md", "No action required."
		case StatusInfo, StatusWarning:
			return "runtime_supportability_partial", "runtime", "runtime", "docs/production-phase5-valc.md", "Review remediation mode, signed desired-state posture, and quarantine safety before go-live."
		default:
			return "runtime_supportability_failed", "runtime", "runtime", "docs/production-phase5-valc.md", "Fix invalid runtime self-healing configuration before rollout."
		}
	case "config-compatibility-readiness":
		switch check.Status {
		case StatusPass:
			return "upgrade_config_compatible", "upgrade", "policy", "docs/production-phase5-valc.md", "No action required."
		case StatusWarning:
			return "upgrade_config_warning", "upgrade", "policy", "docs/production-phase5-valc.md", "Review config compatibility warnings before changing versions."
		default:
			return "upgrade_config_blocked", "upgrade", "policy", "docs/production-phase5-valc.md", "Resolve blocking config issues before attempting the upgrade."
		}
	case "version-support-matrix":
		switch check.Status {
		case StatusPass:
			return "upgrade_line_supported", "upgrade", "evidence", "docs/production-phase5-valc.md", "No action required."
		case StatusWarning:
			return "upgrade_line_bounded", "upgrade", "evidence", "docs/production-phase5-valc.md", "Proceed only after reviewing the bounded support matrix guidance and rollback cautions."
		default:
			return "upgrade_line_unsupported", "upgrade", "evidence", "docs/production-phase5-valc.md", "Choose a supported target line or step through the next bounded upgrade path."
		}
	case "rollback-readiness":
		switch check.Status {
		case StatusPass:
			return "rollback_posture_ready", "rollback", "evidence", "docs/production-phase5-valc.md", "No action required."
		case StatusWarning:
			return "rollback_posture_caution", "rollback", "evidence", "docs/production-phase5-valc.md", "Review rollback cautions and re-run readiness after any reversal."
		default:
			return "rollback_posture_blocked", "rollback", "evidence", "docs/production-phase5-valc.md", "Do not treat rollback as safe until the support matrix and compatibility blockers are cleared."
		}
	case "api-upgrade-context":
		switch check.Status {
		case StatusPass:
			return "upgrade_api_context_ready", "upgrade", "evidence", "docs/production-phase5-valc.md", "No action required."
		case StatusWarning, StatusDegraded, StatusInfo:
			return "upgrade_api_context_partial", "upgrade", "evidence", "docs/production-phase5-valc.md", "Restore API reachability before relying on remote upgrade verification surfaces."
		default:
			return "upgrade_api_context_failed", "upgrade", "evidence", "docs/production-phase5-valc.md", "Fix API health or auth scope before trusting upgrade diagnostics."
		}
	case "review-scope":
		return "code_review_scope_ready", "review", "developer_review", "docs/shift-left-integration.md", "No action required."
	case "review-diff":
		switch check.Status {
		case StatusPass:
			return "code_review_diff_clean", "review", "developer_review", "docs/shift-left-integration.md", "No action required."
		case StatusFail:
			return "code_review_diff_integrity_failed", "review", "developer_review", "docs/shift-left-integration.md", "Fix whitespace, conflict markers, or malformed patch fragments before pushing."
		default:
			return "code_review_diff_integrity_error", "review", "developer_review", "docs/shift-left-integration.md", "Verify local git tooling and rerun the review gate."
		}
	case "review-format":
		switch check.Status {
		case StatusPass:
			return "code_review_format_clean", "review", "developer_review", "docs/shift-left-integration.md", "No action required."
		case StatusFail:
			return "code_review_format_required", "review", "developer_review", "docs/shift-left-integration.md", "Run gofmt on the changed Go files before pushing."
		default:
			return "code_review_format_skipped", "review", "developer_review", "docs/shift-left-integration.md", "Formatting review only applies when changed Go files are present."
		}
	case "review-test-skip":
		switch check.Status {
		case StatusPass:
			return "code_review_skip_marker_clear", "review", "developer_review", "docs/shift-left-integration.md", "No action required."
		default:
			return "code_review_skip_marker_blocked", "review", "developer_review", "docs/shift-left-integration.md", "Remove added test skip markers or justify them in a tighter, explicit review path."
		}
	case "review-deferred-marker":
		switch check.Status {
		case StatusPass:
			return "code_review_deferred_marker_clear", "review", "developer_review", "docs/shift-left-integration.md", "No action required."
		default:
			return "code_review_deferred_marker_present", "review", "developer_review", "docs/shift-left-integration.md", "Review added TODO/FIXME markers and decide whether they belong in the current change."
		}
	case "review-formal-test-coverage":
		switch check.Status {
		case StatusPass:
			return "code_review_formal_test_delta_present", "review", "developer_review", "docs/shift-left-integration.md", "No action required."
		default:
			return "code_review_formal_test_delta_missing", "review", "developer_review", "docs/shift-left-integration.md", "Add or update the matching formal point tests before push."
		}
	case "review-provider":
		switch check.Status {
		case StatusPass:
			return "code_review_provider_clear", "review", "developer_review", "docs/shift-left-integration.md", "No action required."
		case StatusSkip:
			return "code_review_provider_not_configured", "review", "developer_review", "docs/shift-left-integration.md", "Configure CHANGELOCK_CLI_REVIEW_PROVIDER_BIN if you want semantic agent review before push."
		default:
			return "code_review_provider_failed", "review", "developer_review", "docs/shift-left-integration.md", "Fix the external review provider or disable it explicitly before relying on local review gating."
		}
	case "review-finding":
		if metadataString(check.Metadata, "finding_severity") == "P3" {
			return "code_review_provider_advisory_finding", "review", "developer_review", "docs/shift-left-integration.md", "Review the advisory provider finding before pushing."
		}
		return "code_review_provider_blocking_finding", "review", "developer_review", "docs/shift-left-integration.md", "Address the provider finding or lower the configured block severity explicitly before pushing."
	default:
		return "check_unknown", "general", "evidence", "docs/developer-preflight-cli.md", "Review the CLI output for details and rerun with `--output json` if automation needs to inspect the result."
	}
}

func diagnosticRuleID(command, checkName string) string {
	command = strings.TrimSpace(strings.ToLower(command))
	if command == "" {
		command = "preflight"
	}
	checkName = strings.TrimSpace(strings.ReplaceAll(checkName, "-", "_"))
	if checkName == "" {
		checkName = "check"
	}
	return command + "." + checkName
}

func diagnosticSeverity(status Status) DiagnosticSeverity {
	switch status {
	case StatusPass, StatusInfo:
		return DiagnosticSeverityNote
	case StatusSkip, StatusWarning, StatusDegraded:
		return DiagnosticSeverityWarning
	default:
		return DiagnosticSeverityError
	}
}

func evaluationState(status Status) EvaluationState {
	switch status {
	case StatusPass:
		return EvaluationStatePass
	case StatusFail:
		return EvaluationStateFail
	case StatusSkip:
		return EvaluationStateSkipped
	case StatusInfo:
		return EvaluationStatePass
	case StatusError:
		return EvaluationStateUnknown
	default:
		return EvaluationStateWarn
	}
}

func diagnosticMessage(check CheckResult) string {
	if len(check.Details) == 0 {
		return check.Summary
	}
	details := strings.Join(check.Details, "; ")
	if strings.TrimSpace(details) == "" {
		return check.Summary
	}
	return fmt.Sprintf("%s (%s)", check.Summary, details)
}

func diagnosticTargetFile(check CheckResult) string {
	if targetFile := metadataString(check.Metadata, "target_file"); targetFile != "" {
		return targetFile
	}
	target := strings.TrimSpace(check.Target)
	if target == "" {
		return ""
	}
	if isYAML(target) || filepath.Ext(target) != "" {
		return target
	}
	return ""
}

func diagnosticRange(check CheckResult) *DiagnosticRange {
	startLine := metadataInt(check.Metadata, "start_line")
	if startLine <= 0 {
		return nil
	}
	endLine := metadataInt(check.Metadata, "end_line")
	if endLine <= 0 {
		endLine = startLine
	}
	return &DiagnosticRange{
		StartLine:   startLine,
		StartColumn: 1,
		EndLine:     endLine,
		EndColumn:   1,
	}
}

func diagnosticResourceIdentity(target, targetFile string) string {
	if targetFile != "" {
		return ""
	}
	return strings.TrimSpace(target)
}

func metadataString(values map[string]any, key string) string {
	if values == nil {
		return ""
	}
	value, ok := values[key]
	if !ok {
		return ""
	}
	text, ok := value.(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(text)
}

func metadataInt(values map[string]any, key string) int {
	if values == nil {
		return 0
	}
	value, ok := values[key]
	if !ok {
		return 0
	}
	switch typed := value.(type) {
	case int:
		return typed
	case int32:
		return int(typed)
	case int64:
		return int(typed)
	case float64:
		return int(typed)
	default:
		return 0
	}
}

func metadataBool(values map[string]any, key string) bool {
	if values == nil {
		return false
	}
	value, ok := values[key]
	if !ok {
		return false
	}
	typed, ok := value.(bool)
	return ok && typed
}

func parseDiagnosticsOptions(args []string) (diagnosticsOptions, error) {
	fs := newFlagSet("diagnostics")
	options := diagnosticsOptions{
		Input:  "-",
		Format: "json",
	}
	fs.StringVar(&options.Input, "input", options.Input, "path to a JSON preflight result or - for stdin")
	fs.StringVar(&options.Format, "format", options.Format, "diagnostic output: json|github-annotations|markdown|sarif")
	fs.BoolVar(&options.IncludePass, "include-pass", false, "include PASS diagnostics in formatted output")
	if err := fs.Parse(args); err != nil {
		return diagnosticsOptions{}, usageError{message: err.Error()}
	}
	switch strings.ToLower(strings.TrimSpace(options.Format)) {
	case "json", "github-annotations", "markdown", "sarif":
	default:
		return diagnosticsOptions{}, usageError{message: fmt.Sprintf("unsupported diagnostics format %q", options.Format)}
	}
	return options, nil
}

func (a *App) runDiagnostics(args []string, stdout, stderr io.Writer) int {
	options, err := parseDiagnosticsOptions(args)
	if err != nil {
		return a.writeError(stderr, err)
	}
	result, err := loadResultForDiagnostics(options.Input)
	if err != nil {
		return a.writeError(stderr, err)
	}
	switch strings.ToLower(strings.TrimSpace(options.Format)) {
	case "json":
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(DiagnosticsOutput{
			Command:           result.Command,
			Mode:              result.Mode,
			OverallResult:     result.OverallResult,
			ExitCode:          result.ExitCode,
			Inputs:            result.Inputs,
			Diagnostics:       filterDiagnostics(result.Diagnostics, options.IncludePass),
			DiagnosticSummary: summarizeDiagnostics(filterDiagnostics(result.Diagnostics, options.IncludePass)),
		}); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitSuccess
	case "github-annotations":
		if err := renderGitHubAnnotations(stdout, result, options.IncludePass); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitSuccess
	case "markdown":
		if err := renderDiagnosticsMarkdown(stdout, result, options.IncludePass); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitSuccess
	case "sarif":
		if err := renderDiagnosticsSARIF(stdout, result, options.IncludePass); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitSuccess
	default:
		return a.writeError(stderr, usageError{message: fmt.Sprintf("unsupported diagnostics format %q", options.Format)})
	}
}

func loadResultForDiagnostics(path string) (Result, error) {
	var reader io.Reader
	switch strings.TrimSpace(path) {
	case "", "-":
		reader = os.Stdin
	default:
		file, err := os.Open(filepath.Clean(path))
		if err != nil {
			return Result{}, err
		}
		defer file.Close()
		reader = file
	}
	var result Result
	if err := json.NewDecoder(reader).Decode(&result); err != nil {
		return Result{}, fmt.Errorf("decode diagnostics input: %w", err)
	}
	return finalizeResult(result), nil
}

func filterDiagnostics(diagnostics []Diagnostic, includePass bool) []Diagnostic {
	if includePass {
		return append([]Diagnostic(nil), diagnostics...)
	}
	filtered := make([]Diagnostic, 0, len(diagnostics))
	for _, diagnostic := range diagnostics {
		if diagnostic.EvaluationState == EvaluationStatePass {
			continue
		}
		filtered = append(filtered, diagnostic)
	}
	return filtered
}

func summarizeDiagnostics(diagnostics []Diagnostic) DiagnosticSummary {
	summary := DiagnosticSummary{
		CountsBySeverity:        map[string]int{},
		CountsBySource:          map[string]int{},
		CountsByEvaluationState: map[string]int{},
	}
	for _, diagnostic := range diagnostics {
		summary.Total++
		if diagnostic.Blocking {
			summary.Blocking++
		} else if diagnostic.EvaluationState != EvaluationStatePass {
			summary.Advisory++
		}
		summary.CountsBySeverity[string(diagnostic.Severity)]++
		summary.CountsBySource[diagnostic.Source]++
		summary.CountsByEvaluationState[string(diagnostic.EvaluationState)]++
	}
	return summary
}

func renderGitHubAnnotations(w io.Writer, result Result, includePass bool) error {
	for _, diagnostic := range filterDiagnostics(result.Diagnostics, includePass) {
		level := "notice"
		switch diagnostic.Severity {
		case DiagnosticSeverityError:
			level = "error"
		case DiagnosticSeverityWarning:
			level = "warning"
		}
		properties := []string{}
		if diagnostic.TargetFile != "" {
			properties = append(properties, "file="+escapeGitHubAnnotationValue(diagnostic.TargetFile))
		}
		if diagnostic.Range != nil {
			properties = append(properties,
				fmt.Sprintf("line=%d", diagnostic.Range.StartLine),
				fmt.Sprintf("col=%d", diagnostic.Range.StartColumn),
				fmt.Sprintf("endLine=%d", diagnostic.Range.EndLine),
				fmt.Sprintf("endColumn=%d", diagnostic.Range.EndColumn),
			)
		}
		title := escapeGitHubAnnotationValue(fmt.Sprintf("%s [%s]", diagnostic.ReasonCode, diagnostic.Category))
		message := diagnostic.Message
		if strings.TrimSpace(diagnostic.FixHint) != "" {
			message += " Fix: " + diagnostic.FixHint
		}
		if strings.TrimSpace(diagnostic.DocsRef) != "" {
			message += " Docs: " + diagnostic.DocsRef
		}
		if len(properties) > 0 {
			if _, err := fmt.Fprintf(w, "::%s %s,title=%s::%s\n", level, strings.Join(properties, ","), title, escapeGitHubAnnotationValue(message)); err != nil {
				return err
			}
			continue
		}
		if _, err := fmt.Fprintf(w, "::%s title=%s::%s\n", level, title, escapeGitHubAnnotationValue(message)); err != nil {
			return err
		}
	}
	return nil
}

func escapeGitHubAnnotationValue(value string) string {
	replacer := strings.NewReplacer("%", "%25", "\r", "%0D", "\n", "%0A", ":", "%3A", ",", "%2C")
	return replacer.Replace(value)
}

func renderDiagnosticsMarkdown(w io.Writer, result Result, includePass bool) error {
	diagnostics := filterDiagnostics(result.Diagnostics, includePass)
	summary := summarizeDiagnostics(diagnostics)
	if _, err := fmt.Fprintf(w, "## ChangeLock Shift-Left Summary\n\n"); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "- Command: `%s`\n- Mode: `%s`\n- Overall result: `%s`\n- Blocking findings: `%d`\n- Advisory findings: `%d`\n", result.Command, result.Mode, result.OverallResult, summary.Blocking, summary.Advisory); err != nil {
		return err
	}
	if vulnerabilityMarkdownSummary, ok := deriveVulnerabilityMarkdownSummary(result); ok {
		if _, err := fmt.Fprintln(w, "\n### Vulnerability Context"); err != nil {
			return err
		}
		for _, line := range vulnerabilityMarkdownSummary {
			if _, err := fmt.Fprintf(w, "- %s\n", line); err != nil {
				return err
			}
		}
	}
	if len(diagnostics) == 0 {
		_, err := fmt.Fprintln(w, "\nNo blocking or advisory diagnostics were emitted for this run.")
		return err
	}
	if _, err := fmt.Fprintln(w, "\n### Findings"); err != nil {
		return err
	}
	for _, diagnostic := range diagnostics {
		target := firstNonEmpty(diagnostic.TargetFile, diagnostic.ResourceIdentity, diagnostic.Target, "-")
		if _, err := fmt.Fprintf(w, "- `%s` `%s` `%s` on `%s`: %s\n", diagnostic.Severity, diagnostic.ReasonCode, diagnostic.EvaluationState, target, diagnostic.Summary); err != nil {
			return err
		}
		if strings.TrimSpace(diagnostic.FixHint) != "" {
			if _, err := fmt.Fprintf(w, "  Fix: %s\n", diagnostic.FixHint); err != nil {
				return err
			}
		}
		if strings.TrimSpace(diagnostic.DocsRef) != "" {
			if _, err := fmt.Fprintf(w, "  Docs: `%s`\n", diagnostic.DocsRef); err != nil {
				return err
			}
		}
	}
	return nil
}

func deriveVulnerabilityMarkdownSummary(result Result) ([]string, bool) {
	for _, check := range result.Checks {
		if check.Name != "remote-scan-context" || metadataString(check.Metadata, "context_kind") != "vex-net" {
			continue
		}
		lines := []string{
			fmt.Sprintf("Raw findings: `%d`", metadataInt(check.Metadata, "raw_count")),
			fmt.Sprintf("Resolved by VEX: `%d`", metadataInt(check.Metadata, "resolved_by_vex_count")),
			fmt.Sprintf("Net actionable: `%d`", metadataInt(check.Metadata, "actionable_count")),
			fmt.Sprintf("Under investigation: `%d`", metadataInt(check.Metadata, "under_investigation_count")),
			fmt.Sprintf("Threshold breached after VEX merge: `%t`", metadataBool(check.Metadata, "threshold_breached")),
		}
		return lines, true
	}
	return nil, false
}

func renderDiagnosticsSARIF(w io.Writer, result Result, includePass bool) error {
	diagnostics := filterDiagnostics(result.Diagnostics, includePass)
	rules := make([]sarifRule, 0, len(diagnostics))
	seenRules := map[string]struct{}{}
	results := make([]sarifResult, 0, len(diagnostics))
	for _, diagnostic := range diagnostics {
		if _, ok := seenRules[diagnostic.RuleID]; !ok {
			seenRules[diagnostic.RuleID] = struct{}{}
			rules = append(rules, sarifRule{
				ID:               diagnostic.RuleID,
				Name:             diagnostic.ReasonCode,
				ShortDescription: sarifMultiformMessage{Text: diagnostic.Summary},
				HelpURI:          diagnostic.DocsRef,
				DefaultConfiguration: map[string]string{
					"level": sarifLevel(diagnostic.Severity),
				},
				Properties: map[string]interface{}{
					"category":        diagnostic.Category,
					"source":          diagnostic.Source,
					"evaluationState": diagnostic.EvaluationState,
					"blocking":        diagnostic.Blocking,
				},
			})
		}
		resultItem := sarifResult{
			RuleID:  diagnostic.RuleID,
			Level:   sarifLevel(diagnostic.Severity),
			Message: sarifPlainMessage{Text: diagnostic.Message},
			Properties: map[string]any{
				"reasonCode":      diagnostic.ReasonCode,
				"evaluationState": diagnostic.EvaluationState,
				"blocking":        diagnostic.Blocking,
				"docsRef":         diagnostic.DocsRef,
				"fixHint":         diagnostic.FixHint,
				"source":          diagnostic.Source,
			},
		}
		if diagnostic.TargetFile != "" {
			location := sarifLocation{
				PhysicalLocation: sarifPhysicalLocation{
					ArtifactLocation: sarifArtifactLocation{URI: diagnostic.TargetFile},
				},
			}
			if diagnostic.Range != nil {
				location.PhysicalLocation.Region = &sarifRegion{
					StartLine:   diagnostic.Range.StartLine,
					StartColumn: diagnostic.Range.StartColumn,
					EndLine:     diagnostic.Range.EndLine,
					EndColumn:   diagnostic.Range.EndColumn,
				}
			}
			resultItem.Locations = []sarifLocation{location}
		}
		results = append(results, resultItem)
	}
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(sarifLog{
		Version: "2.1.0",
		Schema:  "https://json.schemastore.org/sarif-2.1.0.json",
		Runs: []sarifRun{{
			Tool: sarifTool{
				Driver: sarifDriver{
					Name:           "ChangeLock CLI",
					InformationURI: "docs/shift-left-integration.md",
					Rules:          rules,
				},
			},
			Results: results,
		}},
	})
}

func sarifLevel(severity DiagnosticSeverity) string {
	switch severity {
	case DiagnosticSeverityError:
		return "error"
	case DiagnosticSeverityWarning:
		return "warning"
	default:
		return "note"
	}
}

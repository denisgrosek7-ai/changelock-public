package preflightcli

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

type Status string

const (
	StatusPass     Status = "PASS"
	StatusFail     Status = "FAIL"
	StatusSkip     Status = "SKIP"
	StatusError    Status = "ERROR"
	StatusWarning  Status = "WARNING"
	StatusDegraded Status = "DEGRADED"
	StatusInfo     Status = "INFO"
)

type CheckMode string

const (
	ModeLocal       CheckMode = "local"
	ModeRemote      CheckMode = "remote"
	ModeLocalOnly   string    = "local-only"
	ModeOffline     string    = "offline"
	ModeAPIAssisted string    = "api-assisted"
)

const (
	ExitSuccess   = 0
	ExitFailed    = 1
	ExitUsage     = 2
	ExitExecution = 3
)

type CheckResult struct {
	Name     string         `json:"name"`
	Mode     CheckMode      `json:"mode"`
	Status   Status         `json:"status"`
	Summary  string         `json:"summary"`
	Target   string         `json:"target,omitempty"`
	Details  []string       `json:"details,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

type DiagnosticSeverity string

const (
	DiagnosticSeverityError   DiagnosticSeverity = "error"
	DiagnosticSeverityWarning DiagnosticSeverity = "warning"
	DiagnosticSeverityNote    DiagnosticSeverity = "note"
)

type EvaluationState string

const (
	EvaluationStatePass    EvaluationState = "pass"
	EvaluationStateFail    EvaluationState = "fail"
	EvaluationStateWarn    EvaluationState = "warn"
	EvaluationStateUnknown EvaluationState = "unknown"
	EvaluationStateSkipped EvaluationState = "skipped"
)

type DiagnosticRange struct {
	StartLine   int `json:"start_line"`
	StartColumn int `json:"start_column"`
	EndLine     int `json:"end_line"`
	EndColumn   int `json:"end_column"`
}

type Diagnostic struct {
	CheckID          string             `json:"check_id"`
	RuleID           string             `json:"rule_id"`
	Category         string             `json:"category"`
	Severity         DiagnosticSeverity `json:"severity"`
	ReasonCode       string             `json:"reason_code"`
	Message          string             `json:"message"`
	Summary          string             `json:"summary"`
	Target           string             `json:"target,omitempty"`
	TargetFile       string             `json:"target_file,omitempty"`
	Range            *DiagnosticRange   `json:"range,omitempty"`
	ResourceIdentity string             `json:"resource_identity,omitempty"`
	FixHint          string             `json:"fix_hint,omitempty"`
	DocsRef          string             `json:"docs_ref,omitempty"`
	Source           string             `json:"source"`
	Blocking         bool               `json:"blocking"`
	EvaluationState  EvaluationState    `json:"evaluation_state"`
}

type DiagnosticSummary struct {
	Total                   int            `json:"total"`
	Blocking                int            `json:"blocking"`
	Advisory                int            `json:"advisory"`
	CountsBySeverity        map[string]int `json:"counts_by_severity,omitempty"`
	CountsBySource          map[string]int `json:"counts_by_source,omitempty"`
	CountsByEvaluationState map[string]int `json:"counts_by_evaluation_state,omitempty"`
}

type Result struct {
	Command           string            `json:"command"`
	Mode              string            `json:"mode"`
	Inputs            map[string]string `json:"inputs,omitempty"`
	Checks            []CheckResult     `json:"checks"`
	Diagnostics       []Diagnostic      `json:"diagnostics,omitempty"`
	DiagnosticSummary DiagnosticSummary `json:"diagnostic_summary,omitempty"`
	ReasonCodes       []string          `json:"reason_codes,omitempty"`
	Limitations       []string          `json:"limitations,omitempty"`
	OverallResult     Status            `json:"overall_result"`
	ExitCode          int               `json:"exit_code"`
}

func (r *Result) add(checks ...CheckResult) {
	r.Checks = append(r.Checks, checks...)
}

func finalizeResult(result Result) Result {
	if len(result.Checks) == 0 {
		result.OverallResult = StatusError
		result.ExitCode = ExitExecution
		return attachDiagnostics(result)
	}
	hasPass := false
	hasWarning := false
	hasDegraded := false
	hasInfo := false
	for _, check := range result.Checks {
		switch check.Status {
		case StatusError:
			result.OverallResult = StatusError
			result.ExitCode = ExitExecution
			return attachDiagnostics(result)
		case StatusFail:
			result.OverallResult = StatusFail
			result.ExitCode = ExitFailed
			return attachDiagnostics(result)
		case StatusPass:
			hasPass = true
		case StatusWarning:
			hasWarning = true
		case StatusDegraded:
			hasDegraded = true
		case StatusInfo:
			hasInfo = true
		}
	}
	if hasDegraded {
		result.OverallResult = StatusDegraded
		result.ExitCode = ExitSuccess
		return attachDiagnostics(result)
	}
	if hasWarning {
		result.OverallResult = StatusWarning
		result.ExitCode = ExitSuccess
		return attachDiagnostics(result)
	}
	if hasPass {
		result.OverallResult = StatusPass
		result.ExitCode = ExitSuccess
		return attachDiagnostics(result)
	}
	if hasInfo {
		result.OverallResult = StatusInfo
		result.ExitCode = ExitSuccess
		return attachDiagnostics(result)
	}
	result.OverallResult = StatusSkip
	result.ExitCode = ExitExecution
	return attachDiagnostics(result)
}

func exitCodeForResult(result Result) int {
	switch finalizeResult(result).OverallResult {
	case StatusPass:
		return ExitSuccess
	case StatusFail:
		return ExitFailed
	case StatusError:
		return ExitExecution
	case StatusWarning, StatusDegraded, StatusInfo:
		return ExitSuccess
	default:
		return ExitExecution
	}
}

func renderResult(w io.Writer, output string, result Result) error {
	result = finalizeResult(result)
	switch strings.ToLower(strings.TrimSpace(output)) {
	case "", "human":
		return renderHuman(w, result)
	case "json":
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return encoder.Encode(result)
	default:
		return fmt.Errorf("unsupported output mode %q", output)
	}
}

func renderHuman(w io.Writer, result Result) error {
	if _, err := fmt.Fprintf(w, "Command: %s\n", result.Command); err != nil {
		return err
	}
	if len(result.Inputs) > 0 {
		for key, value := range result.Inputs {
			if strings.TrimSpace(value) == "" {
				continue
			}
			if _, err := fmt.Fprintf(w, "%s: %s\n", key, value); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}

	tw := tabwriter.NewWriter(w, 0, 2, 2, ' ', 0)
	for _, check := range result.Checks {
		target := check.Target
		if target == "" {
			target = "-"
		}
		if _, err := fmt.Fprintf(tw, "[%s]\t%s\t%s\t%s\t%s\n", check.Status, check.Mode, check.Name, target, check.Summary); err != nil {
			return err
		}
	}
	if err := tw.Flush(); err != nil {
		return err
	}
	for _, check := range result.Checks {
		if len(check.Details) == 0 {
			continue
		}
		if _, err := fmt.Fprintf(w, "\n%s details:\n", check.Name); err != nil {
			return err
		}
		for _, detail := range check.Details {
			if _, err := fmt.Fprintf(w, "- %s\n", detail); err != nil {
				return err
			}
		}
	}
	if len(result.ReasonCodes) > 0 {
		if _, err := fmt.Fprintln(w, "\nReason codes:"); err != nil {
			return err
		}
		for _, reason := range result.ReasonCodes {
			if _, err := fmt.Fprintf(w, "- %s\n", reason); err != nil {
				return err
			}
		}
	}
	if len(result.Limitations) > 0 {
		if _, err := fmt.Fprintln(w, "\nLimitations:"); err != nil {
			return err
		}
		for _, limitation := range result.Limitations {
			if _, err := fmt.Fprintf(w, "- %s\n", limitation); err != nil {
				return err
			}
		}
	}
	_, err := fmt.Fprintf(w, "\nMode: %s\nExit code: %d\n[RESULT] %s\n", result.Mode, result.ExitCode, result.OverallResult)
	return err
}

func truncateLines(text string, limit int) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}
	lines := strings.Split(text, "\n")
	if limit > 0 && len(lines) > limit {
		lines = append(lines[:limit], "...output truncated...")
	}
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return lines
}

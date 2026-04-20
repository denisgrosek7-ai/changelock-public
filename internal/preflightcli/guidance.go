package preflightcli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	aiguidance "github.com/denisgrosek/changelock/internal/guidance"
)

type guidanceOptions struct {
	Input       string
	Format      string
	IncludePass bool
}

type GuidanceOutput struct {
	Command       string              `json:"command"`
	Mode          string              `json:"mode"`
	OverallResult Status              `json:"overall_result"`
	ExitCode      int                 `json:"exit_code"`
	Inputs        map[string]string   `json:"inputs,omitempty"`
	Guidance      aiguidance.Response `json:"guidance"`
}

func parseGuidanceOptions(args []string) (guidanceOptions, error) {
	fs := newFlagSet("guidance")
	options := guidanceOptions{
		Input:  "-",
		Format: "json",
	}
	fs.StringVar(&options.Input, "input", options.Input, "path to a JSON preflight result or - for stdin")
	fs.StringVar(&options.Format, "format", options.Format, "guidance output: json|markdown")
	fs.BoolVar(&options.IncludePass, "include-pass", false, "include PASS diagnostics when building guidance facts")
	if err := fs.Parse(args); err != nil {
		return guidanceOptions{}, usageError{message: err.Error()}
	}
	switch strings.ToLower(strings.TrimSpace(options.Format)) {
	case "json", "markdown":
	default:
		return guidanceOptions{}, usageError{message: fmt.Sprintf("unsupported guidance format %q", options.Format)}
	}
	return options, nil
}

func (a *App) runGuidance(args []string, stdout, stderr io.Writer) int {
	options, err := parseGuidanceOptions(args)
	if err != nil {
		return a.writeError(stderr, err)
	}
	result, err := loadResultForDiagnostics(options.Input)
	if err != nil {
		return a.writeError(stderr, err)
	}
	config, err := aiguidance.ParseConfig(os.Getenv)
	if err != nil {
		return a.writeError(stderr, err)
	}
	response := buildPreflightGuidance(result, config, options.IncludePass, time.Now().UTC())
	switch strings.ToLower(strings.TrimSpace(options.Format)) {
	case "json":
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(GuidanceOutput{
			Command:       result.Command,
			Mode:          result.Mode,
			OverallResult: result.OverallResult,
			ExitCode:      result.ExitCode,
			Inputs:        result.Inputs,
			Guidance:      response,
		}); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitSuccess
	case "markdown":
		if err := renderGuidanceMarkdown(stdout, result, response); err != nil {
			return a.writeError(stderr, err)
		}
		return ExitSuccess
	default:
		return a.writeError(stderr, usageError{message: fmt.Sprintf("unsupported guidance format %q", options.Format)})
	}
}

func buildPreflightGuidance(result Result, config aiguidance.Config, includePass bool, now time.Time) aiguidance.Response {
	result = finalizeResult(result)
	scope := guidanceScopeFromResult(result)
	facts := guidanceFactsFromResult(result, includePass)
	return aiguidance.Build(scope, facts, config, now)
}

func guidanceScopeFromResult(result Result) aiguidance.Scope {
	repository := strings.TrimSpace(result.Inputs["repository"])
	scopeType := "workspace"
	scopeRef := firstNonEmpty(repository, result.Inputs["image"], result.Inputs["policy_dir"], result.Command)
	if repository != "" {
		scopeType = "repository"
	}
	return aiguidance.Scope{
		ScopeType:   scopeType,
		ScopeRef:    firstNonEmpty(scopeRef, scopeType+":default"),
		TenantID:    strings.TrimSpace(result.Inputs["tenant"]),
		Environment: strings.TrimSpace(result.Inputs["environment"]),
		Repository:  repository,
	}
}

func guidanceFactsFromResult(result Result, includePass bool) []aiguidance.InputFact {
	diagnostics := filterDiagnostics(result.Diagnostics, includePass)
	facts := make([]aiguidance.InputFact, 0, len(diagnostics))
	for _, diagnostic := range diagnostics {
		check := matchingCheck(result.Checks, diagnostic)
		metadata := map[string]string{}
		if check != nil {
			metadata = stringifyMetadata(check.Metadata)
		}
		reasonCodes := []string{diagnostic.ReasonCode}
		if diagnostic.EvaluationState == EvaluationStateUnknown || diagnostic.EvaluationState == EvaluationStateSkipped {
			reasonCodes = append(reasonCodes, aiguidance.ReasonGuidanceMissingContext)
		}
		scopeRef := firstNonEmpty(diagnostic.ResourceIdentity, diagnostic.TargetFile, diagnostic.Target, result.Inputs["repository"], result.Command)
		facts = append(facts, aiguidance.InputFact{
			ID:                 firstNonEmpty(diagnostic.RuleID+":"+scopeRef, diagnostic.CheckID+":"+scopeRef),
			Category:           guidanceCategoryForDiagnostic(diagnostic),
			SourceComponent:    "changelock-cli",
			RelatedReasonCodes: reasonCodes,
			FindingRefs:        []string{firstNonEmpty(diagnostic.RuleID, diagnostic.CheckID)},
			EvidenceRefs:       guidanceEvidenceRefs(result, diagnostic),
			DocsRefs:           nonEmptyStringSlice(diagnostic.DocsRef),
			ScopeType:          guidanceScopeTypeForDiagnostic(result, diagnostic),
			ScopeRef:           scopeRef,
			TenantID:           strings.TrimSpace(result.Inputs["tenant"]),
			Environment:        strings.TrimSpace(result.Inputs["environment"]),
			Repository:         strings.TrimSpace(result.Inputs["repository"]),
			Severity:           string(diagnostic.Severity),
			Summary:            diagnostic.Summary,
			Detail:             diagnostic.Message,
			Metadata:           metadata,
			Blocking:           diagnostic.Blocking,
			Deterministic:      true,
		})
	}
	return facts
}

func matchingCheck(checks []CheckResult, diagnostic Diagnostic) *CheckResult {
	for i := range checks {
		check := &checks[i]
		if check.Name != diagnostic.CheckID {
			continue
		}
		if diagnostic.Target == "" || check.Target == diagnostic.Target {
			return check
		}
	}
	return nil
}

func stringifyMetadata(metadata map[string]any) map[string]string {
	if len(metadata) == 0 {
		return map[string]string{}
	}
	result := make(map[string]string, len(metadata))
	for key, value := range metadata {
		switch typed := value.(type) {
		case string:
			result[key] = strings.TrimSpace(typed)
		case int:
			result[key] = fmt.Sprintf("%d", typed)
		case int32:
			result[key] = fmt.Sprintf("%d", typed)
		case int64:
			result[key] = fmt.Sprintf("%d", typed)
		case float64:
			result[key] = fmt.Sprintf("%.0f", typed)
		case bool:
			if typed {
				result[key] = "true"
			} else {
				result[key] = "false"
			}
		}
	}
	return result
}

func guidanceCategoryForDiagnostic(diagnostic Diagnostic) string {
	switch diagnostic.Source {
	case "vulnerability", "vex":
		return aiguidance.CategoryVulnerability
	case "signer_identity":
		return aiguidance.CategorySigning
	case "policy":
		return aiguidance.CategoryPolicy
	case "evidence":
		if diagnostic.Category == "supply-chain" {
			return aiguidance.CategoryArtifact
		}
		return aiguidance.CategoryContext
	default:
		switch diagnostic.Category {
		case "vulnerability":
			return aiguidance.CategoryVulnerability
		case "policy":
			return aiguidance.CategoryPolicy
		case "supply-chain":
			return aiguidance.CategoryArtifact
		default:
			return aiguidance.CategoryShiftLeft
		}
	}
}

func guidanceEvidenceRefs(result Result, diagnostic Diagnostic) []string {
	refs := []string{}
	if diagnostic.TargetFile != "" {
		refs = append(refs, "file:"+filepath.Clean(diagnostic.TargetFile))
	}
	if target := strings.TrimSpace(diagnostic.Target); target != "" && target != diagnostic.TargetFile {
		refs = append(refs, "target:"+target)
	}
	if image := strings.TrimSpace(result.Inputs["image"]); image != "" {
		refs = append(refs, "image:"+image)
	}
	return uniqueNonEmpty(refs...)
}

func guidanceScopeTypeForDiagnostic(result Result, diagnostic Diagnostic) string {
	if strings.TrimSpace(result.Inputs["repository"]) != "" {
		return "repository"
	}
	if diagnostic.TargetFile != "" {
		return "file"
	}
	return "workspace"
}

func renderGuidanceMarkdown(w io.Writer, result Result, response aiguidance.Response) error {
	if _, err := fmt.Fprintln(w, "## ChangeLock Contextual Guidance"); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "\n- Command: `%s`\n- Mode: `%s`\n- Guidance mode: `%s`\n- Deterministic only: `%t`\n- Total guidance items: `%d`\n", result.Command, result.Mode, response.Summary.GuidanceMode, response.Summary.DeterministicOnly, response.Summary.TotalItems); err != nil {
		return err
	}
	for _, limitation := range response.Summary.Limitations {
		if strings.TrimSpace(limitation) == "" {
			continue
		}
		if _, err := fmt.Fprintf(w, "- Limitation: %s\n", limitation); err != nil {
			return err
		}
	}
	if len(response.Items) == 0 {
		_, err := fmt.Fprintln(w, "\nNo contextual guidance items were derived from the current deterministic findings.")
		return err
	}
	if _, err := fmt.Fprintln(w, "\n### Prioritized Items"); err != nil {
		return err
	}
	for _, item := range response.Items {
		if _, err := fmt.Fprintf(w, "- `%s` `%s` `%s`: %s\n", item.Priority, item.Category, item.Confidence, firstNonEmpty(item.RecommendationSummary, item.Explanation)); err != nil {
			return err
		}
		if strings.TrimSpace(item.Explanation) != "" {
			if _, err := fmt.Fprintf(w, "  Why: %s\n", item.Explanation); err != nil {
				return err
			}
		}
		for _, step := range item.RecommendationSteps {
			if strings.TrimSpace(step) == "" {
				continue
			}
			if _, err := fmt.Fprintf(w, "  Step: %s\n", step); err != nil {
				return err
			}
		}
		if strings.TrimSpace(item.SaferAlternative) != "" {
			if _, err := fmt.Fprintf(w, "  Safer alternative: %s\n", item.SaferAlternative); err != nil {
				return err
			}
		}
		if strings.TrimSpace(item.ImpactSummary) != "" {
			if _, err := fmt.Fprintf(w, "  Impact: %s\n", item.ImpactSummary); err != nil {
				return err
			}
		}
		for _, limitation := range item.DataLimitations {
			if strings.TrimSpace(limitation) == "" {
				continue
			}
			if _, err := fmt.Fprintf(w, "  Limitation: %s\n", limitation); err != nil {
				return err
			}
		}
		if item.VEXDraft != nil {
			if _, err := fmt.Fprintf(w, "  VEX draft: `%s` · %s\n", item.VEXDraft.CandidateStatus, item.VEXDraft.Justification); err != nil {
				return err
			}
		}
		if item.BreakGlassGuidance != nil {
			if _, err := fmt.Fprintf(w, "  Break-glass scope: %s\n", item.BreakGlassGuidance.ScopeExplanation); err != nil {
				return err
			}
		}
		if len(item.DocsRefs) > 0 {
			if _, err := fmt.Fprintf(w, "  Docs: `%s`\n", strings.Join(item.DocsRefs, "`, `")); err != nil {
				return err
			}
		}
	}
	return nil
}

func nonEmptyStringSlice(values ...string) []string {
	items := make([]string, 0, len(values))
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			continue
		}
		items = append(items, strings.TrimSpace(value))
	}
	return items
}

func uniqueNonEmpty(values ...string) []string {
	seen := map[string]struct{}{}
	items := make([]string, 0, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		items = append(items, trimmed)
	}
	return items
}

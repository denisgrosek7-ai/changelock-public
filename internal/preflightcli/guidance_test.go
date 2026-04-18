package preflightcli

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	aiguidance "github.com/denisgrosek/changelock/internal/guidance"
)

func TestGuidanceCommandMarkdownIncludesVEXDraftSuggestion(t *testing.T) {
	t.Setenv("CHANGELOCK_AI_GUIDANCE_MODE", "local-template")
	inputPath := writeDiagnosticsInput(t, Result{
		Command: "preflight",
		Mode:    ModeAPIAssisted,
		Inputs: map[string]string{
			"tenant":      "acme",
			"repository":  "my-org/acme-app",
			"image":       "ghcr.io/example/api@sha256:abcd",
			"environment": "prod",
		},
		Checks: []CheckResult{
			{
				Name:    "scan",
				Mode:    ModeLocal,
				Status:  StatusFail,
				Summary: "trivy scan found 1 finding at or above HIGH",
				Target:  "ghcr.io/example/api@sha256:abcd",
			},
			{
				Name:    "remote-scan-context",
				Mode:    ModeRemote,
				Status:  StatusFail,
				Summary: "net actionable vulnerability context still breaches HIGH after VEX merge",
				Target:  "ghcr.io/example/api@sha256:abcd",
				Metadata: map[string]any{
					"context_kind":              "vex-net",
					"raw_count":                 2,
					"resolved_by_vex_count":     0,
					"actionable_count":          1,
					"under_investigation_count": 0,
					"threshold_breached":        true,
				},
			},
		},
	})

	app := newTestApp(t, Runtime{})
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"guidance", "--input", inputPath, "--format", "markdown"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success exit code, got %d stderr=%q", code, stderr.String())
	}
	body := stdout.String()
	for _, expected := range []string{
		"## ChangeLock Contextual Guidance",
		"Guidance mode: `local-template`",
		"VEX draft: `under_investigation`",
		"Prefer digest-scoped remediation or a tightly scoped VEX statement",
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %q in guidance markdown, got %q", expected, body)
		}
	}
}

func TestGuidanceCommandJSONKeepsDeterministicFallbackWhenDisabled(t *testing.T) {
	t.Setenv("CHANGELOCK_AI_GUIDANCE_MODE", "disabled")
	manifestPath := writeTempYAML(t, "deployment.yaml")
	inputPath := writeDiagnosticsInput(t, Result{
		Command: "manifest",
		Mode:    ModeLocalOnly,
		Inputs: map[string]string{
			"tenant":     "acme",
			"repository": "my-org/acme-app",
			"policy_dir": "deploy/kyverno",
		},
		Checks: []CheckResult{
			{
				Name:    "manifest",
				Mode:    ModeLocal,
				Status:  StatusFail,
				Summary: "Kyverno reported policy violations",
				Target:  manifestPath,
			},
		},
	})

	app := newTestApp(t, Runtime{})
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"guidance", "--input", inputPath, "--format", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success exit code, got %d stderr=%q", code, stderr.String())
	}
	var output GuidanceOutput
	if err := json.NewDecoder(stdout).Decode(&output); err != nil {
		t.Fatalf("decode guidance json: %v", err)
	}
	if !output.Guidance.Summary.DeterministicOnly {
		t.Fatalf("expected deterministic fallback, got %#v", output.Guidance.Summary)
	}
	if len(output.Guidance.Items) == 0 {
		t.Fatalf("expected at least one guidance item, got %#v", output.Guidance)
	}
	if output.Guidance.Items[0].Confidence == "" || output.Guidance.Items[0].RecommendationSummary == "" {
		t.Fatalf("expected populated guidance item, got %#v", output.Guidance.Items[0])
	}
}

func TestBuildPreflightGuidanceMarksUnknownContextLimited(t *testing.T) {
	response := buildPreflightGuidance(finalizeResult(Result{
		Command: "scan",
		Mode:    ModeAPIAssisted,
		Inputs: map[string]string{
			"tenant": "acme",
			"image":  "ghcr.io/example/api@sha256:abcd",
		},
		Checks: []CheckResult{
			{
				Name:    "remote-scan-context",
				Mode:    ModeRemote,
				Status:  StatusError,
				Summary: "failed to query net actionable vulnerability context",
				Target:  "ghcr.io/example/api@sha256:abcd",
			},
		},
	}), mustGuidanceConfig("local-template"), false, testTime())
	if len(response.Items) != 1 {
		t.Fatalf("expected one guidance item, got %#v", response.Items)
	}
	if response.Items[0].Confidence != "limited" {
		t.Fatalf("expected limited confidence, got %#v", response.Items[0])
	}
}

func mustGuidanceConfig(mode string) aiguidance.Config {
	return aiguidance.Config{
		Mode:            mode,
		MaxItems:        12,
		IncludeDocs:     true,
		RedactSensitive: true,
	}
}

func testTime() time.Time {
	return time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC)
}

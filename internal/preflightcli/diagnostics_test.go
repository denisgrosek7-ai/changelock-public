package preflightcli

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
)

func TestFinalizeResultAddsDiagnostics(t *testing.T) {
	manifestPath := writeTempYAML(t, "deployment.yaml")
	result := finalizeResult(Result{
		Command: "preflight",
		Mode:    ModeOffline,
		Checks: []CheckResult{
			{
				Name:    "manifest",
				Mode:    ModeLocal,
				Status:  StatusFail,
				Summary: "Kyverno reported policy violations",
				Target:  manifestPath,
			},
			{
				Name:    "remote-auth",
				Mode:    ModeRemote,
				Status:  StatusSkip,
				Summary: "API-assisted checks disabled or no API URL configured",
			},
		},
	})

	if len(result.Diagnostics) != 2 {
		t.Fatalf("expected 2 diagnostics, got %d", len(result.Diagnostics))
	}
	if result.Diagnostics[0].ReasonCode != "manifest_policy_violation" {
		t.Fatalf("unexpected first reason code: %+v", result.Diagnostics[0])
	}
	if result.Diagnostics[0].TargetFile != manifestPath {
		t.Fatalf("expected manifest target file, got %+v", result.Diagnostics[0])
	}
	if result.DiagnosticSummary.Blocking != 1 || result.DiagnosticSummary.Advisory != 1 {
		t.Fatalf("unexpected diagnostic summary: %+v", result.DiagnosticSummary)
	}
}

func TestDiagnosticsCommandRendersGitHubAnnotations(t *testing.T) {
	manifestPath := writeTempYAML(t, "deployment.yaml")
	inputPath := writeDiagnosticsInput(t, Result{
		Command: "manifest",
		Mode:    ModeLocalOnly,
		Checks: []CheckResult{{
			Name:    "manifest",
			Mode:    ModeLocal,
			Status:  StatusFail,
			Summary: "Kyverno reported policy violations",
			Target:  manifestPath,
			Details: []string{"disallow latest tag"},
		}},
	})

	app := newTestApp(t, Runtime{})
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"diagnostics", "--input", inputPath, "--format", "github-annotations"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success exit code, got %d stderr=%q", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "::error") || !strings.Contains(stdout.String(), "manifest_policy_violation") {
		t.Fatalf("expected GitHub annotation output, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), manifestPath) {
		t.Fatalf("expected manifest path in annotation output, got %q", stdout.String())
	}
}

func TestDiagnosticsCommandMarkdownIncludesVEXSummary(t *testing.T) {
	inputPath := writeDiagnosticsInput(t, Result{
		Command: "scan",
		Mode:    ModeAPIAssisted,
		Checks: []CheckResult{
			{
				Name:    "scan",
				Mode:    ModeLocal,
				Status:  StatusFail,
				Summary: "trivy scan found 2 findings at or above HIGH",
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
					"raw_count":                 4,
					"resolved_by_vex_count":     2,
					"actionable_count":          2,
					"under_investigation_count": 1,
					"threshold_breached":        true,
				},
			},
		},
	})

	app := newTestApp(t, Runtime{})
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"diagnostics", "--input", inputPath, "--format", "markdown"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success exit code, got %d stderr=%q", code, stderr.String())
	}
	body := stdout.String()
	for _, expected := range []string{"## ChangeLock Shift-Left Summary", "### Vulnerability Context", "Resolved by VEX: `2`", "Net actionable: `2`"} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %q in markdown output, got %q", expected, body)
		}
	}
}

func TestRemoteScanContextUsesNetVulnerabilitiesForDigestPinnedImages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/vulnerabilities/net" {
			http.NotFound(w, r)
			return
		}
		response := audit.VulnerabilityNetResponse{
			RawCount:                3,
			ResolvedByVEXCount:      1,
			ActionableCount:         2,
			UnderInvestigationCount: 1,
			SeverityThreshold:       "HIGH",
			ThresholdBreached:       true,
			Findings: []audit.VulnerabilityFinding{
				{CVEID: "CVE-2026-1000", Severity: "CRITICAL"},
				{CVEID: "CVE-2026-1001", Severity: "HIGH"},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	app := newTestApp(t, Runtime{
		ScanImage: func(_ context.Context, _ Config, image string) (ScanSummary, error) {
			return ScanSummary{
				Scanner: "trivy",
				Image:   image,
				Counts:  map[string]int{"CRITICAL": 1, "HIGH": 1},
				Findings: []ScanFinding{
					{CVEID: "CVE-2026-1000", Severity: "CRITICAL", PackageName: "openssl"},
					{CVEID: "CVE-2026-1001", Severity: "HIGH", PackageName: "glibc"},
				},
			}, nil
		},
		HTTPClient: server.Client(),
	})

	result, err := app.runScan(context.Background(), []string{
		"--image", "ghcr.io/example/api@sha256:abcd",
		"--api-url", server.URL,
		"--fail-severity", "HIGH",
	})
	if err != nil {
		t.Fatalf("runScan error: %v", err)
	}

	var remote CheckResult
	for _, check := range result.Checks {
		if check.Name == "remote-scan-context" {
			remote = check
			break
		}
	}
	if remote.Name == "" {
		t.Fatalf("expected remote-scan-context check, got %+v", result.Checks)
	}
	if remote.Status != StatusFail {
		t.Fatalf("expected remote-scan-context FAIL, got %+v", remote)
	}
	if metadataString(remote.Metadata, "context_kind") != "vex-net" || metadataInt(remote.Metadata, "resolved_by_vex_count") != 1 {
		t.Fatalf("unexpected remote metadata: %+v", remote.Metadata)
	}
}

func writeDiagnosticsInput(t *testing.T, result Result) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "result.json")
	buffer := &bytes.Buffer{}
	if err := renderResult(buffer, "json", result); err != nil {
		t.Fatalf("renderResult: %v", err)
	}
	if err := os.WriteFile(path, buffer.Bytes(), 0o644); err != nil {
		t.Fatalf("write result input: %v", err)
	}
	return path
}

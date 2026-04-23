package preflightcli

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	prodcfg "github.com/denisgrosek/changelock/internal/config"
)

func TestCheckCommandValidatesConfigAndTrustedExecutionProfile(t *testing.T) {
	app := newTestApp(t, Runtime{})
	configPath := writeProductionConfigForCLI(t, withFreshLastSyncedAt(`
apiVersion: changelock.io/v1alpha1
kind: ProductionConfig
metadata:
  name: acme-prod
spec:
  tenant_id: acme
  environment: prod
  repository: my-org/acme-app
  policy_bundle_dir: policies
  kyverno_policy_dir: kyverno
  workflow:
    validation_required: true
    approval_required: true
  sync:
    local_revision: rev-1
    remote_revision: rev-1
    precedence: remote
    last_synced_at: {{fresh_last_synced_at}}
`))

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"check", "--config", configPath, "--trusted-execution-profile", "confidential-strict", "--output", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success, got %d stderr=%q", code, stderr.String())
	}
	result := decodeJSONResult(t, stdout)
	if result.Command != "check" || result.OverallResult != StatusPass {
		t.Fatalf("unexpected result %#v", result)
	}
	if !containsCheck(result.Checks, "trusted-execution-profile", StatusPass) {
		t.Fatalf("expected trusted execution PASS, got %#v", result.Checks)
	}
	if !containsStringValue(result.ReasonCodes, "defaults_applied_visible") {
		t.Fatalf("expected defaults reason code, got %#v", result.ReasonCodes)
	}
}

func TestPreviewCommandKeepsConflictVisible(t *testing.T) {
	app := newTestApp(t, Runtime{})
	configPath := writeProductionConfigForCLI(t, `
apiVersion: changelock.io/v1alpha1
kind: ProductionConfig
metadata:
  name: acme-prod
spec:
  tenant_id: acme
  environment: prod
  policy_bundle_dir: policies
  kyverno_policy_dir: kyverno
  sync:
    local_revision: rev-a
    remote_revision: rev-b
    precedence: local
`)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"preview", "--config", configPath, "--output", "json"}, stdout, stderr)
	if code != ExitFailed {
		t.Fatalf("expected failed preview for sync conflict, got %d stderr=%q", code, stderr.String())
	}
	result := decodeJSONResult(t, stdout)
	if result.Command != "preview" || result.OverallResult != StatusFail {
		t.Fatalf("unexpected result %#v", result)
	}
	if !containsCheck(result.Checks, "sync-preview", StatusFail) {
		t.Fatalf("expected failing sync-preview check, got %#v", result.Checks)
	}
	if !containsStringValue(result.ReasonCodes, "preview_bounded_local_only") {
		t.Fatalf("expected bounded preview reason code, got %#v", result.ReasonCodes)
	}
}

func TestInspectCommandJSONShowsEffectiveConfig(t *testing.T) {
	app := newTestApp(t, Runtime{})
	configPath := writeProductionConfigForCLI(t, `
apiVersion: changelock.io/v1alpha1
kind: ProductionConfig
metadata:
  name: acme-prod
spec:
  tenant_id: acme
  environment: prod
  policy_bundle_dir: policies
  kyverno_policy_dir: kyverno
`)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"inspect", "--config", configPath, "--output", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success, got %d stderr=%q", code, stderr.String())
	}
	var response inspectResponse
	if err := json.Unmarshal(stdout.Bytes(), &response); err != nil {
		t.Fatalf("decode inspect response: %v body=%q", err, stdout.String())
	}
	if response.ConfigInspection.EffectiveConfig.CLI.Output != "human" || response.ConfigInspection.EffectiveConfig.Sync.CurrentState != prodcfg.SyncStateUnconfigured {
		t.Fatalf("unexpected effective config %#v", response.ConfigInspection.EffectiveConfig)
	}
	if response.RuntimeInspection.CurrentState == "" {
		t.Fatalf("expected runtime inspection state, got %#v", response)
	}
}

func TestExplainSyncCommandShowsPrecedence(t *testing.T) {
	app := newTestApp(t, Runtime{})
	configPath := writeProductionConfigForCLI(t, `
apiVersion: changelock.io/v1alpha1
kind: ProductionConfig
metadata:
  name: acme-prod
spec:
  tenant_id: acme
  environment: prod
  policy_bundle_dir: policies
  kyverno_policy_dir: kyverno
  sync:
    local_revision: rev-a
    remote_revision: rev-b
    precedence: local
`)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"explain", "--config", configPath, "--topic", "sync", "--output", "json"}, stdout, stderr)
	if code != ExitFailed {
		t.Fatalf("expected failed explain for invalid sync state, got %d stderr=%q", code, stderr.String())
	}
	var response explainResponse
	if err := json.Unmarshal(stdout.Bytes(), &response); err != nil {
		t.Fatalf("decode explain response: %v body=%q", err, stdout.String())
	}
	if !containsStringValue(response.ReasonCodes, "sync_conflict_visible") {
		t.Fatalf("expected sync conflict reason code, got %#v", response)
	}
	if response.EffectiveState["precedence"] != "local" {
		t.Fatalf("expected local precedence, got %#v", response.EffectiveState)
	}
}

func writeProductionConfigForCLI(t *testing.T, body string) string {
	t.Helper()
	dir := t.TempDir()
	for _, child := range []string{"policies", "kyverno"} {
		if err := os.MkdirAll(filepath.Join(dir, child), 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", child, err)
		}
	}
	path := filepath.Join(dir, "production.yaml")
	if err := os.WriteFile(path, []byte(strings.TrimSpace(body)+"\n"), 0o644); err != nil {
		t.Fatalf("write config: %v", err)
	}
	return path
}

func withFreshLastSyncedAt(body string) string {
	return strings.ReplaceAll(body, "{{fresh_last_synced_at}}", time.Now().Add(-1*time.Hour).UTC().Format(time.RFC3339))
}

func containsCheck(checks []CheckResult, name string, status Status) bool {
	for _, check := range checks {
		if check.Name == name && check.Status == status {
			return true
		}
	}
	return false
}

func containsStringValue(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

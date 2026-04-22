package preflightcli

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadinessCommandProductionFailsWithoutAPI(t *testing.T) {
	app := newTestApp(t, Runtime{})
	configPath := writeProductionConfigForCLI(t, `
apiVersion: changelock.io/v1alpha1
kind: ProductionConfig
metadata:
  name: acme-prod
spec:
  tenant_id: acme
  environment: production
  repository: my-org/acme-app
  policy_bundle_dir: policies
  kyverno_policy_dir: kyverno
  sync:
    local_revision: rev-1
    remote_revision: rev-1
    precedence: remote
    last_synced_at: 2026-04-22T08:00:00Z
  workflow:
    validation_required: true
    approval_required: true
`)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"readiness", "--config", configPath, "--output", "json"}, stdout, stderr)
	if code != ExitFailed {
		t.Fatalf("expected failed readiness without API in production, got %d stderr=%q", code, stderr.String())
	}
	result := decodeJSONResult(t, stdout)
	if result.Command != "readiness" || result.OverallResult != StatusFail {
		t.Fatalf("unexpected readiness result %#v", result)
	}
	if !containsCheck(result.Checks, "api-readiness", StatusFail) {
		t.Fatalf("expected blocking api-readiness check, got %#v", result.Checks)
	}
	if result.DiagnosticSummary.Blocking == 0 {
		t.Fatalf("expected blocking diagnostics, got %#v", result.DiagnosticSummary)
	}
}

func TestSupportCommandReturnsRedactedBundleAndHealthSnapshot(t *testing.T) {
	t.Setenv("CHANGELOCK_SELF_HEALING_MODE", "alert-only")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/healthz":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ok"))
		case r.URL.Path == "/v1/auth/me":
			_ = json.NewEncoder(w).Encode(AuthInfo{
				Authenticated: true,
				AuthMode:      "token",
				Role:          "viewer",
				TenantID:      "acme",
			})
		case strings.HasPrefix(r.URL.Path, "/v1/command-center/timeline"):
			_ = json.NewEncoder(w).Encode(map[string]any{"schema_version": "5a.security_timeline.v1", "entries": []any{}})
		case strings.HasPrefix(r.URL.Path, "/v1/runtime/phase2/proofs"):
			_ = json.NewEncoder(w).Encode(map[string]any{"schema_version": "2.runtime_phase2_proofs.v1", "current_state": "phase2_core_slice_active"})
		case strings.HasPrefix(r.URL.Path, "/v1/intelligence/phase3/proofs"):
			_ = json.NewEncoder(w).Encode(map[string]any{"schema_version": "3.intelligence_phase3_proofs.v1", "current_state": "phase3_core_slice_active"})
		case strings.HasPrefix(r.URL.Path, "/v1/enterprise/phase4/proofs"):
			_ = json.NewEncoder(w).Encode(map[string]any{"schema_version": "4.enterprise_phase4_proofs.v1", "current_state": "phase4_core_slice_active"})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	app := newTestApp(t, Runtime{
		HTTPClient: server.Client(),
		VersionInfo: VersionInfo{
			Version: "5.5.0",
			Commit:  "abc123",
			Date:    "2026-04-22",
		},
	})
	app.config.Token = "demo-token"

	configPath := writeProductionConfigForCLI(t, `
apiVersion: changelock.io/v1alpha1
kind: ProductionConfig
metadata:
  name: acme-prod
spec:
  tenant_id: acme
  environment: production
  repository: my-org/acme-app
  api_url: `+server.URL+`
  policy_bundle_dir: policies
  kyverno_policy_dir: kyverno
  cli:
    output: json
    scanner: auto
    failure_severity: CRITICAL
  sync:
    local_revision: rev-1
    remote_revision: rev-1
    precedence: remote
    last_synced_at: 2026-04-22T08:00:00Z
  workflow:
    validation_required: true
    approval_required: true
`)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"support", "--config", configPath, "--output", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected support bundle success, got %d stderr=%q stdout=%q", code, stderr.String(), stdout.String())
	}

	var response supportBundleResponse
	if err := json.Unmarshal(stdout.Bytes(), &response); err != nil {
		t.Fatalf("decode support bundle: %v body=%q", err, stdout.String())
	}
	if response.RedactionState != "redacted_by_default" {
		t.Fatalf("expected redacted support bundle, got %#v", response)
	}
	if got := response.RuntimeInspection.DeclaredValues["CHANGELOCK_SELF_HEALING_MODE"]; got != "<redacted>" {
		t.Fatalf("expected declared runtime values to be redacted, got %#v", response.RuntimeInspection.DeclaredValues)
	}
	if response.Scope.Profile != profileProduction {
		t.Fatalf("expected effective production profile, got %#v", response.Scope)
	}
	if !supportHealthContains(response.Health.Items, "api", healthStateHealthy) || !supportHealthContains(response.Health.Items, "enterprise-surface", healthStateHealthy) {
		t.Fatalf("expected healthy API and enterprise surface probes, got %#v", response.Health.Items)
	}
	if len(response.Readiness.Checks) == 0 || response.Readiness.OverallResult == "" {
		t.Fatalf("expected embedded readiness summary, got %#v", response.Readiness)
	}
}

func TestUpgradeReadinessCommandShowsSupportMatrixAndRollbackCautions(t *testing.T) {
	app := newTestApp(t, Runtime{
		VersionInfo: VersionInfo{
			Version: "5.5.0",
		},
	})

	configPath := writeProductionConfigForCLI(t, `
apiVersion: changelock.io/v1alpha1
kind: ProductionConfig
metadata:
  name: acme-prod
spec:
  tenant_id: acme
  environment: production
  repository: my-org/acme-app
  policy_bundle_dir: policies
  kyverno_policy_dir: kyverno
  sync:
    local_revision: rev-1
    remote_revision: rev-1
    precedence: remote
    last_synced_at: 2026-04-22T08:00:00Z
  workflow:
    validation_required: true
    approval_required: true
`)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	code := app.Run(context.Background(), []string{"upgrade-readiness", "--config", configPath, "--target-version", "5.6.0", "--output", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected upgrade readiness success, got %d stderr=%q stdout=%q", code, stderr.String(), stdout.String())
	}

	var response upgradeReadinessResponse
	if err := json.Unmarshal(stdout.Bytes(), &response); err != nil {
		t.Fatalf("decode upgrade readiness: %v body=%q", err, stdout.String())
	}
	if response.SupportMatrix.CurrentState != "supported_transition" {
		t.Fatalf("expected supported transition, got %#v", response.SupportMatrix)
	}
	if !response.SupportMatrix.RollbackSupported || len(response.SupportMatrix.RollbackCautions) == 0 {
		t.Fatalf("expected rollback cautions on next-minor upgrade, got %#v", response.SupportMatrix)
	}
	if !containsCheck(response.Result.Checks, "version-support-matrix", StatusPass) {
		t.Fatalf("expected version-support-matrix PASS, got %#v", response.Result.Checks)
	}
	if !containsCheck(response.Result.Checks, "rollback-readiness", StatusWarning) {
		t.Fatalf("expected rollback-readiness warning, got %#v", response.Result.Checks)
	}
}

func supportHealthContains(items []healthComponent, component, state string) bool {
	for _, item := range items {
		if item.Component == component && item.CurrentState == state {
			return true
		}
	}
	return false
}

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

func TestPhase5SummaryRequiresAllSectionsForActiveState(t *testing.T) {
	commandOnly, commandStatus := evaluatePhase5CurrentState(
		phase5SectionSummary{Name: "command_center", OverallResult: StatusPass},
		phase5SectionSummary{Name: "config_cli", OverallResult: StatusFail},
		phase5SectionSummary{Name: "supportability", OverallResult: StatusFail},
	)
	if commandOnly != phase5StateIncomplete || commandStatus != StatusFail {
		t.Fatalf("expected incomplete state with command-center only, got state=%s status=%s", commandOnly, commandStatus)
	}

	commandAndConfig, commandAndConfigStatus := evaluatePhase5CurrentState(
		phase5SectionSummary{Name: "command_center", OverallResult: StatusPass},
		phase5SectionSummary{Name: "config_cli", OverallResult: StatusPass},
		phase5SectionSummary{Name: "supportability", OverallResult: StatusFail},
	)
	if commandAndConfig != phase5StateIncomplete || commandAndConfigStatus != StatusFail {
		t.Fatalf("expected incomplete state without supportability, got state=%s status=%s", commandAndConfig, commandAndConfigStatus)
	}

	fullyReady, fullyReadyStatus := evaluatePhase5CurrentState(
		phase5SectionSummary{Name: "command_center", OverallResult: StatusPass},
		phase5SectionSummary{Name: "config_cli", OverallResult: StatusPass},
		phase5SectionSummary{Name: "supportability", OverallResult: StatusPass},
	)
	if fullyReady != phase5StateProductionReady || fullyReadyStatus != StatusPass {
		t.Fatalf("expected active phase5 state, got state=%s status=%s", fullyReady, fullyReadyStatus)
	}
}

func TestPhase5SummaryCommandConsolidatesValAValBValC(t *testing.T) {
	t.Setenv("CHANGELOCK_SELF_HEALING_MODE", "alert-only")
	t.Setenv("CHANGELOCK_SELF_HEALING_MAX_ATTEMPTS", "3")
	t.Setenv("CHANGELOCK_SELF_HEALING_WINDOW", "15m")
	t.Setenv("CHANGELOCK_SELF_HEALING_ALLOWED_KINDS", "Deployment,DaemonSet,StatefulSet")
	t.Setenv("CHANGELOCK_CLOSED_LOOP_FAIL_MODE", "quarantine")
	t.Setenv("CHANGELOCK_CLOSED_LOOP_RECONCILE_INTERVAL", "2m")
	t.Setenv("CHANGELOCK_CLOSED_LOOP_REQUIRE_SIGNED_DESIRED_STATE", "true")
	t.Setenv("CHANGELOCK_CLOSED_LOOP_PROTECTED_NAMESPACES", "changelock,changelock-system")
	t.Setenv("CHANGELOCK_RUNTIME_VEX_QUARANTINE_SEVERITY", "critical")
	t.Setenv("CHANGELOCK_RUNTIME_VEX_QUARANTINE_REQUIRE_NET_ACTIONABLE", "true")

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
			_ = json.NewEncoder(w).Encode(map[string]any{"schema_version": "2a.security_timeline.v1", "entries": []any{}})
		case strings.HasPrefix(r.URL.Path, "/v1/command-center/notifications"):
			_ = json.NewEncoder(w).Encode(map[string]any{"schema_version": "5a.command_notifications.v1", "items": []any{}})
		case strings.HasPrefix(r.URL.Path, "/v1/command-center/search"):
			_ = json.NewEncoder(w).Encode(map[string]any{"schema_version": "2a.command_search.v1", "results": []any{}})
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
	code := app.Run(context.Background(), []string{"phase5-summary", "--config", configPath, "--output", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected phase5 summary success, got %d stderr=%q stdout=%q", code, stderr.String(), stdout.String())
	}

	var response phase5SummaryResponse
	if err := json.Unmarshal(stdout.Bytes(), &response); err != nil {
		t.Fatalf("decode phase5 summary: %v body=%q", err, stdout.String())
	}
	if response.CurrentState != phase5StateProductionReady || response.OverallResult != StatusPass {
		t.Fatalf("expected active phase5 summary, got %#v", response)
	}
	for _, status := range []Status{
		response.CommandCenter.OverallResult,
		response.ConfigCLI.OverallResult,
		response.Supportability.OverallResult,
	} {
		switch status {
		case StatusPass, StatusFail, StatusWarning, StatusDegraded, StatusInfo:
		default:
			t.Fatalf("unexpected section status taxonomy %q in %#v", status, response)
		}
	}
	if response.Supportability.RedactionState != phase5SectionRedactionState {
		t.Fatalf("expected redacted supportability section, got %#v", response.Supportability)
	}
}

func TestPhase5SummaryRemainsRedacted(t *testing.T) {
	t.Setenv("CHANGELOCK_SELF_HEALING_MODE", "alert-only")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/healthz":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ok"))
		case r.URL.Path == "/v1/auth/me":
			_ = json.NewEncoder(w).Encode(AuthInfo{Authenticated: true, AuthMode: "token", Role: "viewer", TenantID: "acme"})
		case strings.HasPrefix(r.URL.Path, "/v1/command-center/"), strings.HasPrefix(r.URL.Path, "/v1/runtime/phase2/proofs"), strings.HasPrefix(r.URL.Path, "/v1/intelligence/phase3/proofs"), strings.HasPrefix(r.URL.Path, "/v1/enterprise/phase4/proofs"):
			_ = json.NewEncoder(w).Encode(map[string]any{"ok": true})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	app := newTestApp(t, Runtime{HTTPClient: server.Client(), VersionInfo: VersionInfo{Version: "5.5.0"}})
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
	code := app.Run(context.Background(), []string{"phase5-summary", "--config", configPath, "--output", "json"}, stdout, stderr)
	if code != ExitSuccess {
		t.Fatalf("expected phase5 summary success, got %d stderr=%q stdout=%q", code, stderr.String(), stdout.String())
	}
	if strings.Contains(stdout.String(), "alert-only") {
		t.Fatalf("expected redacted phase5 summary, got raw runtime value in %q", stdout.String())
	}
}

func TestPhase5SummaryDoesNotActivateOnCriticalSupportabilityBaseline(t *testing.T) {
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
	code := app.Run(context.Background(), []string{"phase5-summary", "--config", configPath, "--output", "json"}, stdout, stderr)
	if code != ExitFailed {
		t.Fatalf("expected failed phase5 summary without critical supportability baseline, got %d stderr=%q stdout=%q", code, stderr.String(), stdout.String())
	}

	var response phase5SummaryResponse
	if err := json.Unmarshal(stdout.Bytes(), &response); err != nil {
		t.Fatalf("decode phase5 summary: %v body=%q", err, stdout.String())
	}
	if response.CurrentState == phase5StateProductionReady {
		t.Fatalf("expected non-active phase5 summary without API-backed supportability, got %#v", response)
	}
	if response.OverallResult != StatusFail {
		t.Fatalf("expected failed phase5 summary, got %#v", response)
	}
}

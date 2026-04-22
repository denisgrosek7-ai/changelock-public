package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestLoadProductionConfigRejectsUnknownFields(t *testing.T) {
	path := writeProductionConfigFixture(t, `
apiVersion: changelock.io/v1alpha1
kind: ProductionConfig
metadata:
  name: acme-prod
spec:
  tenant_id: acme
  environment: prod
  policy_bundle_dir: policies
  kyverno_policy_dir: kyverno
  unknown_field: nope
`)

	_, err := LoadProductionConfig(path)
	if err == nil || !strings.Contains(err.Error(), "unknown_field") {
		t.Fatalf("expected strict unknown field error, got %v", err)
	}
}

func TestInspectProductionConfigShowsDefaultsAndConflictVisibility(t *testing.T) {
	path := writeProductionConfigFixture(t, `
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
    local_revision: rev-a
    remote_revision: rev-b
    precedence: local
    last_synced_at: 2026-04-20T08:00:00Z
`)

	cfg, err := LoadProductionConfig(path)
	if err != nil {
		t.Fatalf("LoadProductionConfig: %v", err)
	}
	report := InspectProductionConfig(path, cfg, func() time.Time {
		return time.Date(2026, 4, 22, 8, 0, 0, 0, time.UTC)
	})
	if report.CurrentState != "invalid" {
		t.Fatalf("expected invalid conflict state, got %#v", report)
	}
	if !report.EffectiveConfig.Sync.ConflictVisible || !report.EffectiveConfig.Sync.LocalOverrideVisible {
		t.Fatalf("expected visible sync conflict and local override, got %#v", report.EffectiveConfig.Sync)
	}
	if report.EffectiveConfig.CLI.Output != "human" || report.EffectiveConfig.CLI.Scanner != "auto" || report.EffectiveConfig.CLI.FailureSeverity != "CRITICAL" {
		t.Fatalf("expected explicit effective defaults, got %#v", report.EffectiveConfig.CLI)
	}
	if len(report.DefaultsApplied) == 0 {
		t.Fatalf("expected defaults to be visible, got %#v", report)
	}
	if !slicesContain(report.ReasonCodes, "sync_conflict_visible") || !slicesContain(report.ReasonCodes, "defaults_applied_visible") {
		t.Fatalf("expected sync and default reason codes, got %#v", report.ReasonCodes)
	}
}

func writeProductionConfigFixture(t *testing.T, body string) string {
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

func slicesContain(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

package runtime

import (
	"os"
	"testing"
)

func TestLoadSelfHealingConfigDefaults(t *testing.T) {
	t.Setenv("CHANGELOCK_SELF_HEALING_MODE", "")
	t.Setenv("CHANGELOCK_SELF_HEALING_ALLOWED_KINDS", "")

	config, err := LoadSelfHealingConfig()
	if err != nil {
		t.Fatalf("LoadSelfHealingConfig() error = %v", err)
	}
	if config.Mode != RemediationModeDisabled {
		t.Fatalf("expected disabled mode, got %q", config.Mode)
	}
	if _, ok := config.AllowedKinds["Deployment"]; !ok {
		t.Fatalf("expected Deployment to be allowed by default")
	}
	if config.FailMode != RemediationModeQuarantine {
		t.Fatalf("expected quarantine fail mode, got %q", config.FailMode)
	}
	if config.ReconcileInterval <= 0 {
		t.Fatalf("expected positive reconcile interval, got %v", config.ReconcileInterval)
	}
}

func TestLoadSelfHealingConfigRejectsInvalidMode(t *testing.T) {
	t.Setenv("CHANGELOCK_SELF_HEALING_MODE", "mutate-everything")
	if _, err := LoadSelfHealingConfig(); err == nil {
		t.Fatalf("expected invalid mode error")
	}
}

func TestLoadSelfHealingConfigRejectsInvalidKind(t *testing.T) {
	t.Setenv("CHANGELOCK_SELF_HEALING_MODE", string(RemediationModeAlertOnly))
	t.Setenv("CHANGELOCK_SELF_HEALING_ALLOWED_KINDS", "Deployment,Job")
	if _, err := LoadSelfHealingConfig(); err == nil {
		t.Fatalf("expected invalid kind error")
	}
}

func TestParseBoolEnv(t *testing.T) {
	t.Setenv("CHANGELOCK_SELF_HEALING_CRITICAL_ONLY", "true")
	if !parseBoolEnv("CHANGELOCK_SELF_HEALING_CRITICAL_ONLY") {
		t.Fatalf("expected true")
	}
	if parseBoolEnv("CHANGELOCK_SELF_HEALING_NOT_SET") {
		t.Fatalf("expected false for missing key")
	}
}

func TestLoadSelfHealingConfigRejectsInvalidWindow(t *testing.T) {
	t.Setenv("CHANGELOCK_SELF_HEALING_MODE", string(RemediationModeAlertOnly))
	t.Setenv("CHANGELOCK_SELF_HEALING_WINDOW", "nope")
	if _, err := LoadSelfHealingConfig(); err == nil {
		t.Fatalf("expected invalid window error")
	}
}

func TestLoadSelfHealingConfigRejectsInvalidClosedLoopFailMode(t *testing.T) {
	t.Setenv("CHANGELOCK_CLOSED_LOOP_FAIL_MODE", "mutate")
	if _, err := LoadSelfHealingConfig(); err == nil {
		t.Fatalf("expected invalid closed loop fail mode error")
	}
}

func TestLoadSelfHealingConfigRejectsInvalidVEXSeverity(t *testing.T) {
	t.Setenv("CHANGELOCK_RUNTIME_VEX_QUARANTINE_SEVERITY", "panic")
	if _, err := LoadSelfHealingConfig(); err == nil {
		t.Fatalf("expected invalid vex severity error")
	}
}

func TestEnvOrDefault(t *testing.T) {
	t.Setenv("CHANGELOCK_SELF_HEALING_TEST", "value")
	if got := envOrDefault("CHANGELOCK_SELF_HEALING_TEST", "fallback"); got != "value" {
		t.Fatalf("expected value, got %q", got)
	}
	if got := envOrDefault("CHANGELOCK_SELF_HEALING_UNKNOWN_"+os.Getenv("USER"), "fallback"); got != "fallback" {
		t.Fatalf("expected fallback, got %q", got)
	}
}

package runtime

import "testing"

func TestInspectSelfHealingConfigShowsDefaults(t *testing.T) {
	report := InspectSelfHealingConfig(func(string) string { return "" })
	if report.CurrentState != "valid_with_defaults" {
		t.Fatalf("expected valid_with_defaults, got %#v", report)
	}
	if report.EffectiveConfig.Mode != RemediationModeDisabled {
		t.Fatalf("expected disabled mode by default, got %#v", report.EffectiveConfig)
	}
	if len(report.DefaultsApplied) == 0 {
		t.Fatalf("expected defaults to be visible, got %#v", report)
	}
}

func TestInspectSelfHealingConfigReportsInvalidMode(t *testing.T) {
	report := InspectSelfHealingConfig(func(key string) string {
		if key == "CHANGELOCK_SELF_HEALING_MODE" {
			return "explode"
		}
		return ""
	})
	if report.CurrentState != "invalid" {
		t.Fatalf("expected invalid state, got %#v", report)
	}
	if len(report.Issues) != 1 || report.Issues[0].Code != "runtime_self_healing_invalid" {
		t.Fatalf("expected invalid issue code, got %#v", report.Issues)
	}
}

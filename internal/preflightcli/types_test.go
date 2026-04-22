package preflightcli

import "testing"

func TestFinalizeResultPrefersDegradedOverPass(t *testing.T) {
	result := finalizeResult(Result{
		Command: "check",
		Mode:    ModeLocalOnly,
		Checks: []CheckResult{
			{Name: "config", Status: StatusPass, Summary: "config valid"},
			{Name: "sync", Status: StatusDegraded, Summary: "sync is stale"},
		},
	})

	if result.OverallResult != StatusDegraded {
		t.Fatalf("expected degraded overall result, got %#v", result)
	}
}

func TestFinalizeResultPrefersWarningOverPass(t *testing.T) {
	result := finalizeResult(Result{
		Command: "check",
		Mode:    ModeLocalOnly,
		Checks: []CheckResult{
			{Name: "config", Status: StatusPass, Summary: "config valid"},
			{Name: "preview", Status: StatusWarning, Summary: "preview is bounded"},
		},
	})

	if result.OverallResult != StatusWarning {
		t.Fatalf("expected warning overall result, got %#v", result)
	}
}

package runtime

import (
	"testing"
	"time"
)

func TestSelectRemediationModeHonorsSignedRequirement(t *testing.T) {
	config := SelfHealingConfig{
		Mode:                      RemediationModePatchApprovedState,
		AllowedKinds:              map[string]struct{}{"Deployment": {}},
		RequireSignedDesiredState: true,
	}
	desired := ApprovedWorkloadState{
		Namespace:                     "acme-prod",
		WorkloadKind:                  "Deployment",
		Workload:                      "booking-api",
		DesiredStateVerificationState: VerificationStateUnverified,
	}
	result := ComparisonResult{Severity: DriftSeverityCritical, Remediable: true}

	mode, reason := SelectRemediationMode(config, desired, result)
	if mode != RemediationModeQuarantine {
		t.Fatalf("expected quarantine, got %q", mode)
	}
	if reason == "" {
		t.Fatalf("expected quarantine reason")
	}
}

func TestSelectRemediationModeHonorsAlertOnlyFailMode(t *testing.T) {
	config := SelfHealingConfig{
		Mode:                      RemediationModePatchApprovedState,
		AllowedKinds:              map[string]struct{}{"Deployment": {}},
		RequireSignedDesiredState: true,
		FailMode:                  RemediationModeAlertOnly,
	}
	desired := ApprovedWorkloadState{
		Namespace:                     "acme-prod",
		WorkloadKind:                  "Deployment",
		Workload:                      "booking-api",
		DesiredStateVerificationState: VerificationStateUnverified,
	}
	result := ComparisonResult{Severity: DriftSeverityCritical, Remediable: true}

	mode, reason := SelectRemediationMode(config, desired, result)
	if mode != RemediationModeAlertOnly {
		t.Fatalf("expected alert-only, got %q", mode)
	}
	if reason == "" {
		t.Fatalf("expected reason")
	}
}

func TestTrackerQuarantineAndAttempts(t *testing.T) {
	tracker := NewTracker()
	now := time.Date(2026, 4, 16, 12, 0, 0, 0, time.UTC)
	tracker.now = func() time.Time { return now }
	key := "local/acme-prod/Deployment/booking-api"

	if count := tracker.RecordAttempt(key, 5*time.Minute); count != 1 {
		t.Fatalf("expected first attempt count 1, got %d", count)
	}
	now = now.Add(2 * time.Minute)
	if count := tracker.RecordAttempt(key, 5*time.Minute); count != 2 {
		t.Fatalf("expected second attempt count 2, got %d", count)
	}
	tracker.Quarantine(key, "flapping")

	count, quarantined, reason := tracker.Current(key, 5*time.Minute)
	if count != 2 || !quarantined || reason != "flapping" {
		t.Fatalf("unexpected tracker state: count=%d quarantined=%v reason=%q", count, quarantined, reason)
	}
}

func TestProtectedTargetMatchesNamespaceAndWorkload(t *testing.T) {
	config := SelfHealingConfig{
		ProtectedNamespaces: map[string]struct{}{"changelock": {}},
		ProtectedWorkloads:  map[string]struct{}{"acme-prod/Deployment/booking-api": {}},
	}
	if ok, _ := ProtectedTarget(config, ApprovedWorkloadState{Namespace: "changelock", WorkloadKind: "Deployment", Workload: "audit-writer"}); !ok {
		t.Fatalf("expected protected namespace match")
	}
	if ok, _ := ProtectedTarget(config, ApprovedWorkloadState{Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "booking-api"}); !ok {
		t.Fatalf("expected protected workload match")
	}
	if ok, _ := ProtectedTarget(config, ApprovedWorkloadState{Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "frontend"}); ok {
		t.Fatalf("did not expect unrelated workload to be protected")
	}
}

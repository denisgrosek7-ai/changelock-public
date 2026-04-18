package audit

import (
	"testing"
	"time"
)

func TestDeriveRuntimeActiveStatesUsesLatestSnapshot(t *testing.T) {
	now := time.Date(2026, 4, 16, 18, 0, 0, 0, time.UTC)
	events := []StoredEvent{
		{
			ReceivedAt: now,
			Event: Event{
				EventType:             EventTypeRuntimeActiveStateObserved,
				Component:             "runtime-agent",
				TenantID:              "acme",
				ClusterID:             "prod-eu",
				Namespace:             "acme-prod",
				WorkloadKind:          "Deployment",
				Workload:              "booking-api",
				Digest:                "sha256:new",
				DriftResult:           "image_digest_drift",
				DriftSeverity:         "high",
				ReconciliationStatus:  "quarantined",
				RemediationMode:       "quarantine",
				RemediationAttempt:    3,
				QuarantineReason:      "repeat drift",
				QuarantineType:        "repeat-drift",
				DesiredStateVerification: "verified",
				Timestamp:             now,
				Evidence: &Evidence{Runtime: &RuntimeEvidence{
					ApprovedDigest:     "sha256:approved",
					ActualConfigHash:   "cfg-live",
					ExpectedConfigHash: "cfg-approved",
				}},
			},
		},
		{
			ReceivedAt: now.Add(-time.Minute),
			Event: Event{
				EventType:            EventTypeRuntimeActiveStateObserved,
				Component:            "runtime-agent",
				TenantID:             "acme",
				ClusterID:            "prod-eu",
				Namespace:            "acme-prod",
				WorkloadKind:         "Deployment",
				Workload:             "booking-api",
				Digest:               "sha256:old",
				ReconciliationStatus: "in_sync",
				Timestamp:            now.Add(-time.Minute),
			},
		},
	}

	items := DeriveRuntimeActiveStates(events, RuntimeActiveStateFilter{TenantID: "acme", Limit: 10})
	if len(items) != 1 {
		t.Fatalf("expected one active state, got %#v", items)
	}
	if items[0].ObservedDigest != "sha256:new" || items[0].ReconciliationStatus != "quarantined" {
		t.Fatalf("expected latest snapshot, got %#v", items[0])
	}
}

func TestDeriveRuntimeClosedLoopStatusCountsStates(t *testing.T) {
	now := time.Date(2026, 4, 16, 18, 0, 0, 0, time.UTC)
	status := DeriveRuntimeClosedLoopStatus([]RuntimeActiveStateView{
		{ID: "one", ReconciliationStatus: "in_sync", LastReconciledAt: now},
		{ID: "two", ReconciliationStatus: "remediated", LastReconciledAt: now.Add(time.Minute)},
		{ID: "three", ReconciliationStatus: "quarantined", QuarantineType: "vex", ProtectedTarget: true, LastReconciledAt: now.Add(2 * time.Minute)},
	})
	if status.TotalTargets != 3 || status.InSync != 1 || status.Remediated != 1 || status.Quarantined != 1 {
		t.Fatalf("unexpected status summary %#v", status)
	}
	if status.ProtectedBlocked != 1 {
		t.Fatalf("expected protected block count, got %#v", status)
	}
	if status.CountsByQuarantine["vex"] != 1 {
		t.Fatalf("expected vex quarantine count, got %#v", status)
	}
}

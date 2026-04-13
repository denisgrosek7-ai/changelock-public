package audit

import (
	"context"
	"testing"
	"time"
)

func TestMemoryStorePersistsAndFiltersEvents(t *testing.T) {
	store := NewMemoryStore()
	store.now = func() time.Time { return time.Date(2026, 4, 13, 20, 0, 0, 0, time.UTC) }

	_, err := store.Ingest(context.Background(), Event{
		Component:   "deploy-gate",
		EventType:   EventTypeDeployGateDecision,
		TenantID:    "acme",
		Repo:        "my-org/acme-app",
		Environment: "prod",
		Decision:    DecisionDeny,
		Reasons:     []string{"workflow mismatch"},
	})
	if err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	_, err = store.Ingest(context.Background(), Event{
		Component:   "runtime-agent",
		EventType:   EventTypeRuntimeDriftResult,
		TenantID:    "acme",
		Environment: "prod",
		Decision:    DecisionDeny,
		Reasons:     []string{"config drift"},
	})
	if err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	events, err := store.ListEvents(context.Background(), EventFilter{
		Decision: DecisionDeny,
		TenantID: "acme",
		Limit:    10,
	})
	if err != nil {
		t.Fatalf("ListEvents() error = %v", err)
	}
	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d", len(events))
	}
	if events[0].TenantID != "acme" {
		t.Fatalf("unexpected tenant %#v", events[0])
	}
}

func TestMemoryStoreSummaryAggregates(t *testing.T) {
	store := NewMemoryStore()
	base := time.Date(2026, 4, 13, 20, 0, 0, 0, time.UTC)
	store.now = func() time.Time { return base }

	mustIngest := func(event Event) {
		t.Helper()
		if _, err := store.Ingest(context.Background(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}

	mustIngest(Event{Component: "deploy-gate", EventType: EventTypeDeployGateDecision, Decision: DecisionAllow, TenantID: "acme"})
	mustIngest(Event{Component: "deploy-gate", EventType: EventTypeDeployGateDecision, Decision: DecisionDeny, TenantID: "acme", Reasons: []string{"workflow mismatch"}})
	mustIngest(Event{Component: "runtime-agent", EventType: EventTypeRuntimeDriftResult, Decision: DecisionDeny, TenantID: "acme", Reasons: []string{"config drift"}})
	mustIngest(Event{Component: "attestation-verifier", EventType: EventTypeArtifactVerificationResult, Decision: DecisionError, TenantID: "acme", Reasons: []string{"timeout"}})

	summary, err := store.Summary(context.Background(), EventFilter{TenantID: "acme"})
	if err != nil {
		t.Fatalf("Summary() error = %v", err)
	}
	if summary.TotalEvents != 4 || summary.TotalAllow != 1 || summary.TotalDeny != 2 || summary.TotalError != 1 {
		t.Fatalf("unexpected summary %#v", summary)
	}
	if summary.CountsByEventType[EventTypeRuntimeDriftResult] != 1 {
		t.Fatalf("expected runtime drift count, got %#v", summary.CountsByEventType)
	}
	if summary.RecentRuntimeDriftDeny != 1 {
		t.Fatalf("expected recent runtime drift deny count, got %#v", summary)
	}
	if len(summary.TopDenyReasons) == 0 {
		t.Fatalf("expected deny reasons, got %#v", summary)
	}
}

func TestValidateEventRejectsMissingFields(t *testing.T) {
	if err := ValidateEvent(Event{}); err == nil {
		t.Fatalf("expected validation error")
	}
}

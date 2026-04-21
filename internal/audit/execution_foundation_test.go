package audit

import (
	"testing"
	"time"
)

func TestNormalizeEventAddsExecutionEnvelope(t *testing.T) {
	event := NormalizeEvent(Event{
		Component:        "deploy-gate",
		EventType:        EventTypeDeployGateDecision,
		Decision:         DecisionAllow,
		Repo:             "acme/platform-edge",
		Environment:      "prod",
		Namespace:        "acme-prod",
		Workload:         "edge-gateway",
		Digest:           "sha256:edge-v1",
		PolicyBundleHash: "bundle-v1",
	}, func() time.Time {
		return mustParseTime(t, "2026-04-21T10:15:00Z")
	})

	if event.SchemaVersion != ExecutionEventSchemaVersion {
		t.Fatalf("expected schema version %q, got %#v", ExecutionEventSchemaVersion, event)
	}
	if event.EventID == "" || event.TraceID == "" || event.CorrelationID == "" || event.DecisionID == "" || event.IdempotencyKey == "" || event.PayloadHash == "" {
		t.Fatalf("expected populated execution envelope, got %#v", event)
	}
	if event.TraceID != event.RequestID || event.CorrelationID != event.RequestID {
		t.Fatalf("expected request-linked trace/correlation IDs, got %#v", event)
	}
	if event.DecisionID != event.DecisionHash {
		t.Fatalf("expected decision ID to follow decision hash, got %#v", event)
	}
}

func TestNormalizeEventPreservesExecutionEnvelopeOverrides(t *testing.T) {
	event := NormalizeEvent(Event{
		SchemaVersion:  "custom.event.v1",
		EventID:        "event-123",
		TraceID:        "trace-123",
		CorrelationID:  "corr-123",
		DecisionID:     "decision-123",
		CausalParent:   "parent-123",
		IdempotencyKey: "idem-123",
		PayloadHash:    "sha256:payload-123",
		RequestID:      "req-123",
		Component:      "runtime-agent",
		EventType:      EventTypeRuntimeObservationRecorded,
		Decision:       DecisionAllow,
	}, func() time.Time {
		return mustParseTime(t, "2026-04-21T10:15:00Z")
	})

	if event.SchemaVersion != "custom.event.v1" || event.EventID != "event-123" || event.TraceID != "trace-123" || event.CorrelationID != "corr-123" || event.DecisionID != "decision-123" || event.IdempotencyKey != "idem-123" || event.PayloadHash != "sha256:payload-123" {
		t.Fatalf("expected custom execution envelope to be preserved, got %#v", event)
	}
}

func TestExecutionEnvelopeFieldSet(t *testing.T) {
	fields := ExecutionEnvelopeFieldSet()
	if len(fields) < 8 {
		t.Fatalf("expected canonical execution envelope fields, got %#v", fields)
	}
	if fields[0] != "schema_version" || fields[len(fields)-1] != "payload_hash" {
		t.Fatalf("unexpected canonical execution envelope field order %#v", fields)
	}
}

func mustParseTime(t *testing.T, raw string) time.Time {
	t.Helper()
	ts, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		t.Fatalf("parse time %q: %v", raw, err)
	}
	return ts
}

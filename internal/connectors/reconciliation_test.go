package connectors

import (
	"testing"
	"time"
)

func TestReconcileExternalClosureRequiresValidation(t *testing.T) {
	record := Reconcile(ReconciliationInput{
		WorkflowID:      "wf-1",
		ConnectorSystem: "jira",
		ConnectorRef:    "JIRA-101",
		InternalState:   "resolved",
		ExternalState:   "closed",
		ValidationState: "pending",
	}, func() time.Time { return time.Unix(1710000000, 0).UTC() })

	if record.CurrentState != StateExternalClosurePendingValidation {
		t.Fatalf("expected external closure pending validation, got %#v", record)
	}
	if record.SafeToAutoClose {
		t.Fatalf("expected auto-close to remain blocked, got %#v", record)
	}
}

func TestReconcileConnectorFailurePreservesCore(t *testing.T) {
	record := Reconcile(ReconciliationInput{
		ConnectorSystem: "servicenow",
		ConnectorRef:    "SNOW-202",
		InternalState:   "assigned",
		ExternalState:   "in_progress",
		Health: ConnectorHealth{
			CurrentState: HealthFailing,
			LastError:    "rate limited",
		},
	}, func() time.Time { return time.Unix(1710000100, 0).UTC() })

	if record.CurrentState != StateConnectorDegraded {
		t.Fatalf("expected connector degraded state, got %#v", record)
	}
}

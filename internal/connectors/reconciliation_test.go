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

func TestReconcileReopenedOverridesExternalClosure(t *testing.T) {
	testCases := []struct {
		name            string
		validationState string
	}{
		{name: "pending validation", validationState: "pending"},
		{name: "verified validation", validationState: "verified"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			record := Reconcile(ReconciliationInput{
				WorkflowID:      "wf-2",
				ConnectorSystem: "jira",
				ConnectorRef:    "JIRA-202",
				InternalState:   "reopened",
				ExternalState:   "closed",
				ValidationState: tc.validationState,
			}, func() time.Time { return time.Unix(1710000200, 0).UTC() })

			if record.CurrentState != StateReopenedForValidation {
				t.Fatalf("expected reopened precedence over external closure, got %#v", record)
			}
			if record.SafeToAutoClose {
				t.Fatalf("expected auto-close to stay blocked on reopened workflow, got %#v", record)
			}
		})
	}
}

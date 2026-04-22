package workflow

import (
	"testing"
	"time"
)

func TestEvaluateLifecycleRequiresValidationBeforeResolution(t *testing.T) {
	record := EvaluateLifecycle(LifecycleInput{
		WorkflowID:         "wf-1",
		ArtifactType:       "finding",
		SubjectRef:         "cluster-a/prod/Deployment/payments-api",
		Severity:           "critical",
		RequestedState:     StateResolved,
		ValidationRequired: true,
		ValidationState:    ValidationStatePending,
		Owners: Ownership{
			FindingOwner: "team-payments",
		},
	}, func() time.Time { return time.Unix(1710000000, 0).UTC() })

	if record.CurrentState != StateUnderValidation {
		t.Fatalf("expected under_validation, got %#v", record)
	}
	if record.ClosureReady {
		t.Fatalf("expected closure to be blocked before validation, got %#v", record)
	}
}

func TestEvaluateLifecycleFeedbackRejectsWithIdentityTrail(t *testing.T) {
	record := EvaluateLifecycle(LifecycleInput{
		WorkflowID:     "wf-2",
		ArtifactType:   "recommendation",
		RequestedState: StateTriaged,
		Owners: Ownership{
			FindingOwner: "team-platform",
		},
		Feedback: &FeedbackSignal{
			SourceSystem: "slack",
			Actor:        "security-admin",
			Decision:     "reject",
			Reason:       "insufficient evidence",
		},
	}, func() time.Time { return time.Unix(1710000100, 0).UTC() })

	if record.CurrentState != StateRejected {
		t.Fatalf("expected rejected state, got %#v", record)
	}
	if record.Feedback == nil || record.Feedback.Actor != "security-admin" {
		t.Fatalf("expected feedback identity trail, got %#v", record)
	}
}

func TestEvaluateLifecycleFeedbackWithoutActorCannotRejectCanonicalState(t *testing.T) {
	record := EvaluateLifecycle(LifecycleInput{
		WorkflowID:     "wf-3",
		ArtifactType:   "recommendation",
		RequestedState: StateTriaged,
		Owners: Ownership{
			FindingOwner: "team-platform",
		},
		Feedback: &FeedbackSignal{
			SourceSystem: "slack",
			Decision:     "reject",
			Reason:       "missing actor should not reject workflow",
		},
	}, func() time.Time { return time.Unix(1710000200, 0).UTC() })

	if record.CanonicalState == StateRejected || record.CurrentState == StateRejected {
		t.Fatalf("expected feedback without actor to preserve workflow state, got %#v", record)
	}
	if !containsWorkflowReason(record.ReasonCodes, "feedback_missing_identity_trail") {
		t.Fatalf("expected missing identity reason, got %#v", record.ReasonCodes)
	}
	if !containsWorkflowReason(record.ReasonCodes, "feedback_reject_recorded_without_authority") {
		t.Fatalf("expected non-authoritative reject reason, got %#v", record.ReasonCodes)
	}
}

func containsWorkflowReason(values []string, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
}

package handoff

import (
	"testing"
	"time"
)

func TestEvaluateIntakeAcceptedProducesPartnerSafeDashboard(t *testing.T) {
	record := EvaluateIntake(IntakeInput{
		PartnerID:                "vendor-a",
		Organization:             "Vendor A",
		HandoffRef:               "handoff-1",
		VerificationStatus:       "verified",
		FreshnessState:           "fresh",
		PolicyCompatibility:      "compatible",
		IncidentDisclosureStatus: "shared",
		PartnerVisibleEvidence:   []string{"sealed://proof/1"},
	}, func() time.Time { return time.Unix(1710000000, 0).UTC() })

	if record.CurrentState != IntakeStateAccepted {
		t.Fatalf("expected accepted intake, got %#v", record)
	}
	if !record.Dashboard.SensitiveSignalsRedacted {
		t.Fatalf("expected partner dashboard to redact sensitive signals, got %#v", record.Dashboard)
	}
}

func TestEvaluateIntakeRejectsFailedVerification(t *testing.T) {
	record := EvaluateIntake(IntakeInput{
		PartnerID:           "vendor-b",
		VerificationStatus:  "failed",
		FreshnessState:      "fresh",
		PolicyCompatibility: "compatible",
	}, func() time.Time { return time.Unix(1710000100, 0).UTC() })

	if record.CurrentState != IntakeStateRejected {
		t.Fatalf("expected rejected intake, got %#v", record)
	}
}

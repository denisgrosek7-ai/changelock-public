package compliance

import (
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/connectors"
	"github.com/denisgrosek/changelock/internal/handoff"
	"github.com/denisgrosek/changelock/internal/workflow"
)

func TestEvaluateComplianceMappingDistinguishesInferredCoverage(t *testing.T) {
	record := EvaluateComplianceMapping(MappingInput{
		SubjectRef:         "cluster-a/prod/Deployment/payments-api",
		ControlFamily:      "soc2.cc7",
		ControlID:          "CC7.2",
		CoverageState:      CoverageInferred,
		EvidenceRefs:       []string{"event://1"},
		TechnicalEventRefs: []string{"deploy_gate_decision"},
	}, func() time.Time { return time.Unix(1710000000, 0).UTC() })

	if record.CurrentState != "control_partially_supported" {
		t.Fatalf("expected partial support state, got %#v", record)
	}
}

func TestEvaluatePolicyDriftDetectsSoftening(t *testing.T) {
	record := EvaluatePolicyDrift(DriftInput{
		SubjectRef:   "cluster-a/prod/Deployment/payments-api",
		Actor:        "security-admin",
		PreviousMode: "deny",
		CurrentMode:  "exception",
		ExceptionID:  "exc-1",
	}, func() time.Time { return time.Unix(1710000100, 0).UTC() })

	if record.CurrentState != DriftStateSoftened {
		t.Fatalf("expected softened drift state, got %#v", record)
	}
}

func TestBuildExecutiveReportAggregatesEvidence(t *testing.T) {
	report := BuildExecutiveReport(ExecutiveReportInput{
		ScopeRef: "tenant:acme",
		WorkflowArtifacts: []workflow.LifecycleRecord{{
			CurrentState:   workflow.StateUnderValidation,
			RequestedState: workflow.StateResolved,
			EvidenceRefs:   []string{"wf://1"},
		}},
		ReconciliationArtifacts: []connectors.ReconciliationRecord{{
			CurrentState: connectors.StateConnectorDegraded,
			EvidenceRefs: []string{"cn://1"},
		}},
		PartnerArtifacts: []handoff.IntakeRecord{{
			CurrentState: handoff.IntakeStateExpired,
			EvidenceRefs: []string{"pt://1"},
		}},
		ComplianceArtifacts: []ComplianceMappingRecord{{
			CoverageState: CoverageMissing,
			EvidenceRefs:  []string{"cm://1"},
		}},
		DriftArtifacts: []PolicyDriftRecord{{
			CurrentState: DriftStateSoftened,
			EvidenceRefs: []string{"dr://1"},
		}},
	}, func() time.Time { return time.Unix(1710000200, 0).UTC() })

	if report.CurrentState != "executive_governance_attention_required" {
		t.Fatalf("expected attention required report, got %#v", report)
	}
	if len(report.EvidenceTraceRefs) < 5 {
		t.Fatalf("expected aggregated evidence refs, got %#v", report)
	}
}

package workflow

import "testing"

func TestEnterpriseWorkflowAuthorityValAStateRequiresActiveVal0(t *testing.T) {
	got := EvaluateEnterpriseWorkflowAuthorityValAState(
		EnterpriseWorkflowAuthorityVal0StateSubstantial,
		EnterpriseWorkflowAuthorityValAEventOrchestrationStateActive,
		EnterpriseWorkflowAuthorityValALifecycleConnectorsStateActive,
		EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStateActive,
		EnterpriseWorkflowAuthorityValATicketChangeProjectionStateActive,
		EnterpriseWorkflowAuthorityValAReconciliationBaselineStateActive,
		EnterpriseWorkflowAuthorityValAIdempotentMutationStateActive,
	)
	if got != EnterpriseWorkflowAuthorityValAStateIncomplete {
		t.Fatalf("expected incomplete Val A state without active Val 0, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValAEventOrchestrationIsPartialWithoutReplayRecovery(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValAEventOrchestration()
	model.ReplayRecovery = nil
	if got := EvaluateEnterpriseWorkflowAuthorityValAEventOrchestrationState(model); got != EnterpriseWorkflowAuthorityValAEventOrchestrationStatePartial {
		t.Fatalf("expected partial event orchestration without replay recovery, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValALifecycleConnectorsStateIsPartialWithSingleConnector(t *testing.T) {
	items := EnterpriseWorkflowAuthorityValALifecycleConnectors()[:1]
	if got := EvaluateEnterpriseWorkflowAuthorityValALifecycleConnectorsState(items); got != EnterpriseWorkflowAuthorityValALifecycleConnectorsStatePartial {
		t.Fatalf("expected partial lifecycle connectors with single connector coverage, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValALifecycleConnectorsStateIsPartialWithTwoConnectors(t *testing.T) {
	items := EnterpriseWorkflowAuthorityValALifecycleConnectors()[:2]
	if got := EvaluateEnterpriseWorkflowAuthorityValALifecycleConnectorsState(items); got != EnterpriseWorkflowAuthorityValALifecycleConnectorsStatePartial {
		t.Fatalf("expected partial lifecycle connectors with incomplete connector coverage, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValALifecycleConnectorsStateIsPartialWithDuplicateConnector(t *testing.T) {
	items := EnterpriseWorkflowAuthorityValALifecycleConnectors()
	items[1].ConnectorSystem = WorkflowAuthorityConnectorJira
	if got := EvaluateEnterpriseWorkflowAuthorityValALifecycleConnectorsState(items); got != EnterpriseWorkflowAuthorityValALifecycleConnectorsStatePartial {
		t.Fatalf("expected partial lifecycle connectors with duplicate connector coverage, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValALifecycleConnectorsStateIsActiveWithExpectedConnectorSet(t *testing.T) {
	items := EnterpriseWorkflowAuthorityValALifecycleConnectors()
	if got := EvaluateEnterpriseWorkflowAuthorityValALifecycleConnectorsState(items); got != EnterpriseWorkflowAuthorityValALifecycleConnectorsStateActive {
		t.Fatalf("expected active lifecycle connectors with jira/servicenow/github coverage, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValAEvidenceInjectionIsPartialWithoutExternalTicketSafeTier(t *testing.T) {
	items := EnterpriseWorkflowAuthorityValAEvidenceBundleInjection()
	items[0].SupportedRedactionTiers = []string{WorkflowAuthorityEvidenceTierInternalFull, WorkflowAuthorityEvidenceTierPartnerScoped}
	if got := EvaluateEnterpriseWorkflowAuthorityValAEvidenceBundleInjectionState(items); got != EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStatePartial {
		t.Fatalf("expected partial evidence injection without external-ticket-safe tier, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValAIdempotentMutationIsPartialWithoutDuplicateSuppression(t *testing.T) {
	items := EnterpriseWorkflowAuthorityValAIdempotentMutationDiscipline()
	items[0].DuplicateSuppressionRules = nil
	if got := EvaluateEnterpriseWorkflowAuthorityValAIdempotentMutationState(items); got != EnterpriseWorkflowAuthorityValAIdempotentMutationStatePartial {
		t.Fatalf("expected partial idempotent mutation without duplicate suppression, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValAFoundationIsActive(t *testing.T) {
	eventOrchestration := EnterpriseWorkflowAuthorityValAEventOrchestration()
	if got := EvaluateEnterpriseWorkflowAuthorityValAEventOrchestrationState(eventOrchestration); got != EnterpriseWorkflowAuthorityValAEventOrchestrationStateActive {
		t.Fatalf("expected active event orchestration, got %q", got)
	}

	lifecycleConnectors := EnterpriseWorkflowAuthorityValALifecycleConnectors()
	if got := EvaluateEnterpriseWorkflowAuthorityValALifecycleConnectorsState(lifecycleConnectors); got != EnterpriseWorkflowAuthorityValALifecycleConnectorsStateActive {
		t.Fatalf("expected active lifecycle connectors, got %q", got)
	}

	evidenceInjection := EnterpriseWorkflowAuthorityValAEvidenceBundleInjection()
	if got := EvaluateEnterpriseWorkflowAuthorityValAEvidenceBundleInjectionState(evidenceInjection); got != EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStateActive {
		t.Fatalf("expected active evidence injection, got %q", got)
	}

	projection := EnterpriseWorkflowAuthorityValATicketChangeProjection()
	if got := EvaluateEnterpriseWorkflowAuthorityValATicketChangeProjectionState(projection); got != EnterpriseWorkflowAuthorityValATicketChangeProjectionStateActive {
		t.Fatalf("expected active ticket/change projection, got %q", got)
	}

	reconciliation := EnterpriseWorkflowAuthorityValAReconciliationBaseline()
	if got := EvaluateEnterpriseWorkflowAuthorityValAReconciliationBaselineState(reconciliation); got != EnterpriseWorkflowAuthorityValAReconciliationBaselineStateActive {
		t.Fatalf("expected active reconciliation baseline, got %q", got)
	}

	idempotent := EnterpriseWorkflowAuthorityValAIdempotentMutationDiscipline()
	if got := EvaluateEnterpriseWorkflowAuthorityValAIdempotentMutationState(idempotent); got != EnterpriseWorkflowAuthorityValAIdempotentMutationStateActive {
		t.Fatalf("expected active idempotent mutation baseline, got %q", got)
	}

	if got := EvaluateEnterpriseWorkflowAuthorityValAState(
		EnterpriseWorkflowAuthorityVal0StateActive,
		EvaluateEnterpriseWorkflowAuthorityValAEventOrchestrationState(eventOrchestration),
		EvaluateEnterpriseWorkflowAuthorityValALifecycleConnectorsState(lifecycleConnectors),
		EvaluateEnterpriseWorkflowAuthorityValAEvidenceBundleInjectionState(evidenceInjection),
		EvaluateEnterpriseWorkflowAuthorityValATicketChangeProjectionState(projection),
		EvaluateEnterpriseWorkflowAuthorityValAReconciliationBaselineState(reconciliation),
		EvaluateEnterpriseWorkflowAuthorityValAIdempotentMutationState(idempotent),
	); got != EnterpriseWorkflowAuthorityValAStateActive {
		t.Fatalf("expected active overall Val A state, got %q", got)
	}
}

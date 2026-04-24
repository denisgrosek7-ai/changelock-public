package workflow

import "testing"

func TestEnterpriseWorkflowAuthorityValDStateRequiresActiveValC(t *testing.T) {
	got := EvaluateEnterpriseWorkflowAuthorityValDState(
		EnterpriseWorkflowAuthorityValCStateSubstantial,
		EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStateActive,
		EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStateActive,
		EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStateActive,
		EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStateActive,
		EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStateActive,
		EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStateActive,
		EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStateActive,
		EnterpriseWorkflowAuthorityValDReopenRollbackReviewStateActive,
	)
	if got != EnterpriseWorkflowAuthorityValDStateIncomplete {
		t.Fatalf("expected incomplete Val D state without active Val C, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewIsPartialWithoutGitHubConnector(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValDConnectorCorrectnessReview()
	model.RequiredConnectors = []string{WorkflowAuthorityConnectorJira, WorkflowAuthorityConnectorServiceNow}
	if got := EvaluateEnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewState(model); got != EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStatePartial {
		t.Fatalf("expected partial connector correctness review without full connector set, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValDApprovalBoundaryReviewIsPartialWithoutSessionBoundConsumption(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValDApprovalBoundaryReview()
	model.RequiredConsumptionModes = []string{
		WorkflowAuthorityConsumptionSingleUse,
		WorkflowAuthorityConsumptionMultiUseBounded,
	}
	if got := EvaluateEnterpriseWorkflowAuthorityValDApprovalBoundaryReviewState(model); got != EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStatePartial {
		t.Fatalf("expected partial approval boundary review without full consumption mode coverage, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValDExceptionExpiryReviewIsPartialWithoutSupersessionEffects(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValDExceptionExpiryReview()
	model.SupersessionEffectRules = nil
	if got := EvaluateEnterpriseWorkflowAuthorityValDExceptionExpiryReviewState(model); got != EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStatePartial {
		t.Fatalf("expected partial exception expiry review without supersession effects, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValDClosureCorrectnessReviewIsPartialWithoutAdministrativeCloseRules(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValDClosureCorrectnessReview()
	model.AdministrativeCloseRules = nil
	if got := EvaluateEnterpriseWorkflowAuthorityValDClosureCorrectnessReviewState(model); got != EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStatePartial {
		t.Fatalf("expected partial closure correctness review without administrative close rules, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValDReconciliationConflictReviewIsPartialWithoutCanonicalPrecedence(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValDReconciliationConflictReview()
	model.CanonicalPrecedenceRules = []string{"external_close_never_overwrites_canonical_validation_state"}
	if got := EvaluateEnterpriseWorkflowAuthorityValDReconciliationConflictReviewState(model); got != EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStatePartial {
		t.Fatalf("expected partial reconciliation conflict review without full canonical precedence rules, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValDWorkflowLedgerReviewIsPartialWithoutRevocationChecks(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValDWorkflowLedgerReview()
	model.RevocationChecks = nil
	if got := EvaluateEnterpriseWorkflowAuthorityValDWorkflowLedgerReviewState(model); got != EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStatePartial {
		t.Fatalf("expected partial workflow ledger review without revocation checks, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewIsPartialWithoutRollbackDecisionClass(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReview()
	model.RequiredDecisionClasses = []string{"approval", "break_glass", "exception", "validation_close", "reopen"}
	if got := EvaluateEnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewState(model); got != EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStatePartial {
		t.Fatalf("expected partial governance traceability review without rollback decision coverage, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValDReopenRollbackReviewIsPartialWithoutConnectorVisibilityRules(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValDReopenRollbackReview()
	model.ConnectorVisibilityRules = nil
	if got := EvaluateEnterpriseWorkflowAuthorityValDReopenRollbackReviewState(model); got != EnterpriseWorkflowAuthorityValDReopenRollbackReviewStatePartial {
		t.Fatalf("expected partial reopen/rollback review without connector visibility rules, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValDFoundationIsActive(t *testing.T) {
	connectorCorrectness := EnterpriseWorkflowAuthorityValDConnectorCorrectnessReview()
	if got := EvaluateEnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewState(connectorCorrectness); got != EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStateActive {
		t.Fatalf("expected active connector correctness review, got %q", got)
	}

	approvalBoundary := EnterpriseWorkflowAuthorityValDApprovalBoundaryReview()
	if got := EvaluateEnterpriseWorkflowAuthorityValDApprovalBoundaryReviewState(approvalBoundary); got != EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStateActive {
		t.Fatalf("expected active approval boundary review, got %q", got)
	}

	exceptionExpiry := EnterpriseWorkflowAuthorityValDExceptionExpiryReview()
	if got := EvaluateEnterpriseWorkflowAuthorityValDExceptionExpiryReviewState(exceptionExpiry); got != EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStateActive {
		t.Fatalf("expected active exception expiry review, got %q", got)
	}

	closureCorrectness := EnterpriseWorkflowAuthorityValDClosureCorrectnessReview()
	if got := EvaluateEnterpriseWorkflowAuthorityValDClosureCorrectnessReviewState(closureCorrectness); got != EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStateActive {
		t.Fatalf("expected active closure correctness review, got %q", got)
	}

	reconciliation := EnterpriseWorkflowAuthorityValDReconciliationConflictReview()
	if got := EvaluateEnterpriseWorkflowAuthorityValDReconciliationConflictReviewState(reconciliation); got != EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStateActive {
		t.Fatalf("expected active reconciliation conflict review, got %q", got)
	}

	ledger := EnterpriseWorkflowAuthorityValDWorkflowLedgerReview()
	if got := EvaluateEnterpriseWorkflowAuthorityValDWorkflowLedgerReviewState(ledger); got != EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStateActive {
		t.Fatalf("expected active workflow ledger review, got %q", got)
	}

	governance := EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReview()
	if got := EvaluateEnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewState(governance); got != EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStateActive {
		t.Fatalf("expected active governance traceability review, got %q", got)
	}

	reopenRollback := EnterpriseWorkflowAuthorityValDReopenRollbackReview()
	if got := EvaluateEnterpriseWorkflowAuthorityValDReopenRollbackReviewState(reopenRollback); got != EnterpriseWorkflowAuthorityValDReopenRollbackReviewStateActive {
		t.Fatalf("expected active reopen/rollback review, got %q", got)
	}

	if got := EvaluateEnterpriseWorkflowAuthorityValDState(
		EnterpriseWorkflowAuthorityValCStateActive,
		EvaluateEnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewState(connectorCorrectness),
		EvaluateEnterpriseWorkflowAuthorityValDApprovalBoundaryReviewState(approvalBoundary),
		EvaluateEnterpriseWorkflowAuthorityValDExceptionExpiryReviewState(exceptionExpiry),
		EvaluateEnterpriseWorkflowAuthorityValDClosureCorrectnessReviewState(closureCorrectness),
		EvaluateEnterpriseWorkflowAuthorityValDReconciliationConflictReviewState(reconciliation),
		EvaluateEnterpriseWorkflowAuthorityValDWorkflowLedgerReviewState(ledger),
		EvaluateEnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewState(governance),
		EvaluateEnterpriseWorkflowAuthorityValDReopenRollbackReviewState(reopenRollback),
	); got != EnterpriseWorkflowAuthorityValDStateActive {
		t.Fatalf("expected active overall Val D state, got %q", got)
	}
}

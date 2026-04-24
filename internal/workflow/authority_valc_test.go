package workflow

import "testing"

func TestEnterpriseWorkflowAuthorityValCStateRequiresActiveValB(t *testing.T) {
	got := EvaluateEnterpriseWorkflowAuthorityValCState(
		EnterpriseWorkflowAuthorityValBStateSubstantial,
		EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStateActive,
		EnterpriseWorkflowAuthorityValCWorkflowLedgerStateActive,
		EnterpriseWorkflowAuthorityValCStaleReopenHandlingStateActive,
		EnterpriseWorkflowAuthorityValCRollbackLinkageStateActive,
		EnterpriseWorkflowAuthorityValCGovernanceMappingStateActive,
		EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStateActive,
	)
	if got != EnterpriseWorkflowAuthorityValCStateIncomplete {
		t.Fatalf("expected incomplete Val C state without active Val B, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValCClosureValidationEnforcementIsPartialWithoutRevokedExpiryConsequences(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValCClosureValidationEnforcement()
	model.FailureStateConsequences = []string{"superseded_authorization_requires_new_effective_artifact_before_close"}
	if got := EvaluateEnterpriseWorkflowAuthorityValCClosureValidationEnforcementState(model); got != EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStatePartial {
		t.Fatalf("expected partial closure validation enforcement without revoked/expired consequences, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValCWorkflowLedgerIsPartialWithoutAppendOnlySignedEntries(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValCWorkflowLedger()
	model.AppendOnly = false
	if got := EvaluateEnterpriseWorkflowAuthorityValCWorkflowLedgerState(model); got != EnterpriseWorkflowAuthorityValCWorkflowLedgerStatePartial {
		t.Fatalf("expected partial workflow ledger without append-only enforcement, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValCStaleReopenHandlingIsPartialWithoutEvidenceBoundReopen(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValCStaleReopenHandling()
	model.ReopenEvidenceRules = []string{"reopen_must_link_to_prior_closure_or_validation_attempt"}
	if got := EvaluateEnterpriseWorkflowAuthorityValCStaleReopenHandlingState(model); got != EnterpriseWorkflowAuthorityValCStaleReopenHandlingStatePartial {
		t.Fatalf("expected partial stale/reopen handling without evidence-bound reopen rules, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValCRollbackLinkageIsPartialWithoutClosureConsequences(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValCRollbackLinkage()
	model.ClosureConsequences = nil
	if got := EvaluateEnterpriseWorkflowAuthorityValCRollbackLinkageState(model); got != EnterpriseWorkflowAuthorityValCRollbackLinkageStatePartial {
		t.Fatalf("expected partial rollback linkage without closure consequences, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValCReplayRecoveryHardeningIsPartialWithoutCanonicalPrecedence(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValCReplayRecoveryHardening()
	model.CanonicalPrecedenceRules = []string{"external_close_never_overwrites_canonical_validation_state"}
	if got := EvaluateEnterpriseWorkflowAuthorityValCReplayRecoveryHardeningState(model); got != EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStatePartial {
		t.Fatalf("expected partial replay/recovery hardening without full canonical precedence rules, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValCFoundationIsActive(t *testing.T) {
	closureValidation := EnterpriseWorkflowAuthorityValCClosureValidationEnforcement()
	if got := EvaluateEnterpriseWorkflowAuthorityValCClosureValidationEnforcementState(closureValidation); got != EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStateActive {
		t.Fatalf("expected active closure validation enforcement, got %q", got)
	}

	ledger := EnterpriseWorkflowAuthorityValCWorkflowLedger()
	if got := EvaluateEnterpriseWorkflowAuthorityValCWorkflowLedgerState(ledger); got != EnterpriseWorkflowAuthorityValCWorkflowLedgerStateActive {
		t.Fatalf("expected active workflow ledger, got %q", got)
	}

	staleReopen := EnterpriseWorkflowAuthorityValCStaleReopenHandling()
	if got := EvaluateEnterpriseWorkflowAuthorityValCStaleReopenHandlingState(staleReopen); got != EnterpriseWorkflowAuthorityValCStaleReopenHandlingStateActive {
		t.Fatalf("expected active stale/reopen handling, got %q", got)
	}

	rollbackLinkage := EnterpriseWorkflowAuthorityValCRollbackLinkage()
	if got := EvaluateEnterpriseWorkflowAuthorityValCRollbackLinkageState(rollbackLinkage); got != EnterpriseWorkflowAuthorityValCRollbackLinkageStateActive {
		t.Fatalf("expected active rollback linkage, got %q", got)
	}

	governanceMapping := EnterpriseWorkflowAuthorityValCGovernanceMapping()
	if got := EvaluateEnterpriseWorkflowAuthorityValCGovernanceMappingState(governanceMapping); got != EnterpriseWorkflowAuthorityValCGovernanceMappingStateActive {
		t.Fatalf("expected active governance mapping, got %q", got)
	}

	replayRecovery := EnterpriseWorkflowAuthorityValCReplayRecoveryHardening()
	if got := EvaluateEnterpriseWorkflowAuthorityValCReplayRecoveryHardeningState(replayRecovery); got != EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStateActive {
		t.Fatalf("expected active replay/recovery hardening, got %q", got)
	}

	if got := EvaluateEnterpriseWorkflowAuthorityValCState(
		EnterpriseWorkflowAuthorityValBStateActive,
		EvaluateEnterpriseWorkflowAuthorityValCClosureValidationEnforcementState(closureValidation),
		EvaluateEnterpriseWorkflowAuthorityValCWorkflowLedgerState(ledger),
		EvaluateEnterpriseWorkflowAuthorityValCStaleReopenHandlingState(staleReopen),
		EvaluateEnterpriseWorkflowAuthorityValCRollbackLinkageState(rollbackLinkage),
		EvaluateEnterpriseWorkflowAuthorityValCGovernanceMappingState(governanceMapping),
		EvaluateEnterpriseWorkflowAuthorityValCReplayRecoveryHardeningState(replayRecovery),
	); got != EnterpriseWorkflowAuthorityValCStateActive {
		t.Fatalf("expected active overall Val C state, got %q", got)
	}
}

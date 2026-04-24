package workflow

import "testing"

func TestEnterpriseWorkflowAuthorityValBStateRequiresActiveValA(t *testing.T) {
	got := EvaluateEnterpriseWorkflowAuthorityValBState(
		EnterpriseWorkflowAuthorityValAStateSubstantial,
		EnterpriseWorkflowAuthorityValBSignedAuthorizationsStateActive,
		EnterpriseWorkflowAuthorityValBBreakGlassStateActive,
		EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStateActive,
		EnterpriseWorkflowAuthorityValBExpiryRevocationStateActive,
		EnterpriseWorkflowAuthorityValBAntiReplayStateActive,
		EnterpriseWorkflowAuthorityValBApprovalTraceabilityStateActive,
	)
	if got != EnterpriseWorkflowAuthorityValBStateIncomplete {
		t.Fatalf("expected incomplete Val B state without active Val A, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValBSignedAuthorizationsIsPartialWithoutJTIConsumption(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValBSignedAuthorizations()
	model.AntiReplayFields = []string{"subject_nonce_binding", "consumed_at"}
	if got := EvaluateEnterpriseWorkflowAuthorityValBSignedAuthorizationsState(model); got != EnterpriseWorkflowAuthorityValBSignedAuthorizationsStatePartial {
		t.Fatalf("expected partial signed authorizations without anti-replay jti coverage, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValBBreakGlassIsPartialWithoutDistinctApproverExecutor(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValBBreakGlassFlow()
	model.DistinctApproverExecutor = false
	if got := EvaluateEnterpriseWorkflowAuthorityValBBreakGlassState(model); got != EnterpriseWorkflowAuthorityValBBreakGlassStatePartial {
		t.Fatalf("expected partial break-glass flow without distinct approver/executor, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValBManagedExceptionRegistryIsPartialWithoutRevokedStage(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValBManagedExceptionRegistry()
	model.LifecycleStages = []string{"requested", "approved", "activated", "expiring", "expired", "superseded", "revalidated"}
	if got := EvaluateEnterpriseWorkflowAuthorityValBManagedExceptionRegistryState(model); got != EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStatePartial {
		t.Fatalf("expected partial managed exception registry without revoked lifecycle stage, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValBAntiReplayIsPartialWithoutReplayCache(t *testing.T) {
	model := EnterpriseWorkflowAuthorityValBAntiReplayProtection()
	model.ReplayCacheRules = nil
	if got := EvaluateEnterpriseWorkflowAuthorityValBAntiReplayState(model); got != EnterpriseWorkflowAuthorityValBAntiReplayStatePartial {
		t.Fatalf("expected partial anti-replay protection without replay cache rules, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValBFoundationIsActive(t *testing.T) {
	signedAuthorizations := EnterpriseWorkflowAuthorityValBSignedAuthorizations()
	if got := EvaluateEnterpriseWorkflowAuthorityValBSignedAuthorizationsState(signedAuthorizations); got != EnterpriseWorkflowAuthorityValBSignedAuthorizationsStateActive {
		t.Fatalf("expected active signed authorizations, got %q", got)
	}

	breakGlass := EnterpriseWorkflowAuthorityValBBreakGlassFlow()
	if got := EvaluateEnterpriseWorkflowAuthorityValBBreakGlassState(breakGlass); got != EnterpriseWorkflowAuthorityValBBreakGlassStateActive {
		t.Fatalf("expected active break-glass flow, got %q", got)
	}

	exceptionRegistry := EnterpriseWorkflowAuthorityValBManagedExceptionRegistry()
	if got := EvaluateEnterpriseWorkflowAuthorityValBManagedExceptionRegistryState(exceptionRegistry); got != EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStateActive {
		t.Fatalf("expected active managed exception registry, got %q", got)
	}

	expiryRevocation := EnterpriseWorkflowAuthorityValBExpiryRevocationEnforcement()
	if got := EvaluateEnterpriseWorkflowAuthorityValBExpiryRevocationState(expiryRevocation); got != EnterpriseWorkflowAuthorityValBExpiryRevocationStateActive {
		t.Fatalf("expected active expiry/revocation enforcement, got %q", got)
	}

	antiReplay := EnterpriseWorkflowAuthorityValBAntiReplayProtection()
	if got := EvaluateEnterpriseWorkflowAuthorityValBAntiReplayState(antiReplay); got != EnterpriseWorkflowAuthorityValBAntiReplayStateActive {
		t.Fatalf("expected active anti-replay protection, got %q", got)
	}

	approvalTraceability := EnterpriseWorkflowAuthorityValBApprovalTraceability()
	if got := EvaluateEnterpriseWorkflowAuthorityValBApprovalTraceabilityState(approvalTraceability); got != EnterpriseWorkflowAuthorityValBApprovalTraceabilityStateActive {
		t.Fatalf("expected active approval traceability, got %q", got)
	}

	if got := EvaluateEnterpriseWorkflowAuthorityValBState(
		EnterpriseWorkflowAuthorityValAStateActive,
		EvaluateEnterpriseWorkflowAuthorityValBSignedAuthorizationsState(signedAuthorizations),
		EvaluateEnterpriseWorkflowAuthorityValBBreakGlassState(breakGlass),
		EvaluateEnterpriseWorkflowAuthorityValBManagedExceptionRegistryState(exceptionRegistry),
		EvaluateEnterpriseWorkflowAuthorityValBExpiryRevocationState(expiryRevocation),
		EvaluateEnterpriseWorkflowAuthorityValBAntiReplayState(antiReplay),
		EvaluateEnterpriseWorkflowAuthorityValBApprovalTraceabilityState(approvalTraceability),
	); got != EnterpriseWorkflowAuthorityValBStateActive {
		t.Fatalf("expected active overall Val B state, got %q", got)
	}
}

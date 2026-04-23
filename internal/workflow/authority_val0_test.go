package workflow

import "testing"

func TestEnterpriseWorkflowAuthorityVal0StateRequiresAllFoundationSlices(t *testing.T) {
	got := EvaluateEnterpriseWorkflowAuthorityVal0State(
		EnterpriseWorkflowAuthorityVal0BoundaryStateActive,
		EnterpriseWorkflowAuthorityVal0StateMachineStateActive,
		EnterpriseWorkflowAuthorityVal0ProjectionStateActive,
		EnterpriseWorkflowAuthorityVal0ApprovalContractStateActive,
		EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStateActive,
		EnterpriseWorkflowAuthorityVal0ClosureValidationStateActive,
		EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStateActive,
		EnterpriseWorkflowAuthorityVal0TimeAuthorityStatePartial,
	)
	if got != EnterpriseWorkflowAuthorityVal0StateSubstantial {
		t.Fatalf("expected substantial val0 state, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityVal0StateMachineRequiresClosureInvariant(t *testing.T) {
	model := EnterpriseWorkflowAuthorityVal0StateMachine()
	model.GlobalInvariants = nil
	if got := EvaluateEnterpriseWorkflowAuthorityVal0StateMachineState(model); got != EnterpriseWorkflowAuthorityVal0StateMachineStateIncomplete {
		t.Fatalf("expected incomplete state machine without invariants, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityVal0ApprovalContractRequiresAntiReplayDiscipline(t *testing.T) {
	model := EnterpriseWorkflowAuthorityVal0ApprovalContract()
	model.AntiReplayRules = nil
	if got := EvaluateEnterpriseWorkflowAuthorityVal0ApprovalContractState(model); got != EnterpriseWorkflowAuthorityVal0ApprovalContractStatePartial {
		t.Fatalf("expected partial approval contract without anti-replay rules, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityVal0TimeAuthorityRequiresClockSkewRules(t *testing.T) {
	model := EnterpriseWorkflowAuthorityVal0TimeAuthority()
	model.ClockSkewTolerance = ""
	if got := EvaluateEnterpriseWorkflowAuthorityVal0TimeAuthorityState(model); got != EnterpriseWorkflowAuthorityVal0TimeAuthorityStateIncomplete {
		t.Fatalf("expected incomplete time authority without clock skew tolerance, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityVal0FoundationIsActive(t *testing.T) {
	boundary := EnterpriseWorkflowAuthorityVal0BoundaryRules()
	if got := EvaluateEnterpriseWorkflowAuthorityVal0BoundaryState(boundary); got != EnterpriseWorkflowAuthorityVal0BoundaryStateActive {
		t.Fatalf("expected active authority boundary, got %q", got)
	}

	stateMachine := EnterpriseWorkflowAuthorityVal0StateMachine()
	if got := EvaluateEnterpriseWorkflowAuthorityVal0StateMachineState(stateMachine); got != EnterpriseWorkflowAuthorityVal0StateMachineStateActive {
		t.Fatalf("expected active state machine, got %q", got)
	}

	projection := EnterpriseWorkflowAuthorityVal0ExternalProjectionRules()
	if got := EvaluateEnterpriseWorkflowAuthorityVal0ProjectionState(projection); got != EnterpriseWorkflowAuthorityVal0ProjectionStateActive {
		t.Fatalf("expected active projection rules, got %q", got)
	}

	approval := EnterpriseWorkflowAuthorityVal0ApprovalContract()
	if got := EvaluateEnterpriseWorkflowAuthorityVal0ApprovalContractState(approval); got != EnterpriseWorkflowAuthorityVal0ApprovalContractStateActive {
		t.Fatalf("expected active approval contract, got %q", got)
	}

	exceptionLifecycle := EnterpriseWorkflowAuthorityVal0ExceptionLifecycle()
	if got := EvaluateEnterpriseWorkflowAuthorityVal0ExceptionLifecycleState(exceptionLifecycle); got != EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStateActive {
		t.Fatalf("expected active exception lifecycle, got %q", got)
	}

	closure := EnterpriseWorkflowAuthorityVal0ClosureValidation()
	if got := EvaluateEnterpriseWorkflowAuthorityVal0ClosureValidationState(closure); got != EnterpriseWorkflowAuthorityVal0ClosureValidationStateActive {
		t.Fatalf("expected active closure validation, got %q", got)
	}

	separation := EnterpriseWorkflowAuthorityVal0SeparationOfDuties()
	if got := EvaluateEnterpriseWorkflowAuthorityVal0SeparationOfDutiesState(separation); got != EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStateActive {
		t.Fatalf("expected active separation-of-duties, got %q", got)
	}

	timeAuthority := EnterpriseWorkflowAuthorityVal0TimeAuthority()
	if got := EvaluateEnterpriseWorkflowAuthorityVal0TimeAuthorityState(timeAuthority); got != EnterpriseWorkflowAuthorityVal0TimeAuthorityStateActive {
		t.Fatalf("expected active time authority, got %q", got)
	}

	if got := EvaluateEnterpriseWorkflowAuthorityVal0State(
		EvaluateEnterpriseWorkflowAuthorityVal0BoundaryState(boundary),
		stateMachine.CurrentState,
		EvaluateEnterpriseWorkflowAuthorityVal0ProjectionState(projection),
		approval.CurrentState,
		EvaluateEnterpriseWorkflowAuthorityVal0ExceptionLifecycleState(exceptionLifecycle),
		closure.CurrentState,
		EvaluateEnterpriseWorkflowAuthorityVal0SeparationOfDutiesState(separation),
		timeAuthority.CurrentState,
	); got != EnterpriseWorkflowAuthorityVal0StateActive {
		t.Fatalf("expected active overall val0 state, got %q", got)
	}
}

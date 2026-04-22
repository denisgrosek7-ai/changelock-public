package formal

import "testing"

func TestEvaluateFoundationStateRequiresAllCoreSurfaces(t *testing.T) {
	presence := buildCoreSurfacePresence()
	delete(presence, "formal.evidence_custody")

	if got := evaluateFoundationStateForPresence(presence); got != FoundationStateIncomplete {
		t.Fatalf("expected incomplete foundation when evidence custody surface is missing, got %q", got)
	}
	if got := EvaluateFoundationState(ContractsCoverage()); got != FoundationStateActive {
		t.Fatalf("expected active foundation from baseline coverage, got %q", got)
	}
}

func TestEvaluatePhase8StateRequiresGroupSpecificStates(t *testing.T) {
	if got := EvaluatePhase8State(EntryGateStateReady, FoundationStateActive, FormalDisciplineStateActive, ComplianceCodificationStateActive, GovernedAutonomyStateActive); got != Phase8StateActive {
		t.Fatalf("expected active phase8 state, got %q", got)
	}
	if got := EvaluatePhase8State(EntryGateStateReady, FoundationStateActive, ComplianceCodificationStateActive, ComplianceCodificationStateActive, GovernedAutonomyStateActive); got != Phase8StateIncomplete {
		t.Fatalf("expected incomplete phase8 state for wrong slot-specific formal state, got %q", got)
	}
	if got := EvaluatePhase8State(EntryGateStateReady, FoundationStateActive, FormalDisciplineStateActive, ComplianceCodificationStatePartial, GovernedAutonomyStateActive); got != Phase8StateSubstantial {
		t.Fatalf("expected substantial phase8 state when one group is partial, got %q", got)
	}
}

func TestDeferredInstitutionalExpansionStaysOutsideCoreState(t *testing.T) {
	deferred := DeferredInstitutionalExpansion()
	if len(deferred) == 0 {
		t.Fatal("expected deferred institutional expansion to remain visible")
	}
	if got := EvaluateFormalDisciplineState(); got != FormalDisciplineStateActive {
		t.Fatalf("expected active formal discipline state, got %q", got)
	}
	if got := EvaluateComplianceCodificationState(); got != ComplianceCodificationStateActive {
		t.Fatalf("expected active compliance codification state, got %q", got)
	}
	if got := EvaluateGovernedAutonomyState(); got != GovernedAutonomyStateActive {
		t.Fatalf("expected active governed autonomy state, got %q", got)
	}
}

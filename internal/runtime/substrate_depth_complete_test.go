package runtime

import "testing"

func TestRuntimeSubstratePoint1StateRequiresAllValsActive(t *testing.T) {
	got := EvaluateRuntimeSubstratePoint1State(
		RuntimeSubstrateValAStateActive,
		RuntimeSubstrateValBStateActive,
		RuntimeSubstrateValCStateActive,
		RuntimeSubstrateValDStateActive,
		RuntimeSubstrateValEStateSubstantial,
		RuntimeSubstratePoint1DocumentationRefs(),
		RuntimeSubstratePoint1SurfaceRefs(),
		RuntimeSubstratePoint1EvidenceRefs(),
		[]string{"integrated closure remains bounded to A-E surfaces"},
		RuntimeSubstratePoint1DeferredScope(),
		"Point 1 closure binds A-E into one fail-closed summary.",
	)
	if got != RuntimeSubstratePoint1StateSubstantial {
		t.Fatalf("expected substantial point 1 state when val e is not active, got %q", got)
	}
}

func TestRuntimeSubstratePoint1StateRequiresClosureMetadata(t *testing.T) {
	got := EvaluateRuntimeSubstratePoint1State(
		RuntimeSubstrateValAStateActive,
		RuntimeSubstrateValBStateActive,
		RuntimeSubstrateValCStateActive,
		RuntimeSubstrateValDStateActive,
		RuntimeSubstrateValEStateActive,
		nil,
		RuntimeSubstratePoint1SurfaceRefs(),
		RuntimeSubstratePoint1EvidenceRefs(),
		[]string{"integrated closure remains bounded to A-E surfaces"},
		RuntimeSubstratePoint1DeferredScope(),
		"Point 1 closure binds A-E into one fail-closed summary.",
	)
	if got != RuntimeSubstratePoint1StateSubstantial {
		t.Fatalf("expected substantial point 1 state without documentation refs, got %q", got)
	}
}

func TestRuntimeSubstratePoint1StateIsActiveWhenAllValsAndMetadataArePresent(t *testing.T) {
	got := EvaluateRuntimeSubstratePoint1State(
		RuntimeSubstrateValAStateActive,
		RuntimeSubstrateValBStateActive,
		RuntimeSubstrateValCStateActive,
		RuntimeSubstrateValDStateActive,
		RuntimeSubstrateValEStateActive,
		RuntimeSubstratePoint1DocumentationRefs(),
		RuntimeSubstratePoint1SurfaceRefs(),
		RuntimeSubstratePoint1EvidenceRefs(),
		[]string{"Integrated closure remains a summary layer and does not add new runtime mechanics."},
		RuntimeSubstratePoint1DeferredScope(),
		"Point 1 closure binds Val A through Val E into one canonical runtime/substrate completion summary.",
	)
	if got != RuntimeSubstratePoint1StateActive {
		t.Fatalf("expected active point 1 state, got %q", got)
	}
}

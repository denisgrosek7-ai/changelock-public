package runtime

import (
	"testing"
	"time"
)

func TestRuntimeSubstrateValABaselineIsContractDefined(t *testing.T) {
	entryGate := RuntimeSubstrateEntryGateBaseline()
	schema := RuntimeSubstrateValAEventSchema()
	matrix := RuntimeSubstrateValASupportMatrix()
	events := RuntimeSubstrateValAObservedEvents()

	if entryGate.CurrentState != RuntimeSubstrateEntryGateStateSubstantial {
		t.Fatalf("expected substantially ready entry gate, got %#v", entryGate)
	}
	if got := EvaluateRuntimeSubstrateValAEventSchemaState(schema); got != RuntimeSubstrateValAEventSchemaStateActive {
		t.Fatalf("expected active event schema state, got %q", got)
	}
	if got := EvaluateRuntimeSubstrateValASupportMatrixState(matrix); got != RuntimeSubstrateValASupportMatrixStateActive {
		t.Fatalf("expected active support matrix state, got %q", got)
	}
	if got := EvaluateRuntimeSubstrateValAObservabilityState(events); got != RuntimeSubstrateValAObservabilityStateActive {
		t.Fatalf("expected active observability state, got %q", got)
	}
	if got := EvaluateRuntimeSubstrateValAState(entryGate.CurrentState, schema.CurrentState, EvaluateRuntimeSubstrateValASupportMatrixState(matrix), EvaluateRuntimeSubstrateValAObservabilityState(events)); got != RuntimeSubstrateValAStateContractDefined {
		t.Fatalf("expected contract-defined Val A state, got %q", got)
	}
}

func TestRuntimeSubstrateValAStateRequiresMandatoryFamilies(t *testing.T) {
	events := RuntimeSubstrateValAObservedEvents()
	filtered := make([]RuntimeSubstrateObservedEvent, 0, len(events))
	for _, event := range events {
		if event.EventFamily == RuntimeSubstrateEventFamilyNetworkActivity {
			continue
		}
		filtered = append(filtered, event)
	}

	if got := EvaluateRuntimeSubstrateValAObservabilityState(filtered); got != RuntimeSubstrateValAObservabilityStateIncomplete {
		t.Fatalf("expected incomplete observability state without network family, got %q", got)
	}
}

func TestRuntimeSubstrateEntryGateRuntimeBaselineIsReady(t *testing.T) {
	gate := RuntimeSubstrateEntryGateRuntimeBaseline()
	if gate.CurrentState != RuntimeSubstrateEntryGateStateReady {
		t.Fatalf("expected ready runtime baseline gate, got %#v", gate)
	}
}

func TestNormalizeAndValidateRuntimeSubstrateObservedEvent(t *testing.T) {
	event := NormalizeRuntimeSubstrateObservedEvent(RuntimeSubstrateObservedEvent{
		EventFamily:           RuntimeSubstrateEventFamilyExecLifecycle,
		CurrentState:          RuntimeSubstrateEventStateObserved,
		Process:               ProcessIdentity{ProcessName: "api", PID: 42},
		Workload:              WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "api"},
		Node:                  NodeIdentity{NodeID: "node-a", SubstrateClass: "standard", TrustBoundary: TrustBoundaryKernelRuntimeLayer},
		AttributionConfidence: RuntimeSubstrateConfidenceHighFidelity,
		FreshnessState:        RuntimeSubstrateFreshnessFresh,
	}, runtimeSubstrateNow)
	if event.EventID == "" {
		t.Fatal("expected generated event id")
	}
	if event.SchemaVersion != RuntimeSubstrateValAEventRecordSchema {
		t.Fatalf("expected schema version %q, got %#v", RuntimeSubstrateValAEventRecordSchema, event)
	}
	if err := ValidateRuntimeSubstrateObservedEvent(event); err != nil {
		t.Fatalf("expected valid runtime observation, got %v", err)
	}
}

func TestValidateRuntimeSubstrateObservedEventRejectsMissingIdentity(t *testing.T) {
	event := NormalizeRuntimeSubstrateObservedEvent(RuntimeSubstrateObservedEvent{
		EventFamily:           RuntimeSubstrateEventFamilyFileActivity,
		CurrentState:          RuntimeSubstrateEventStatePartiallyCorrelated,
		Workload:              WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "api"},
		Node:                  NodeIdentity{NodeID: "node-a", SubstrateClass: "standard", TrustBoundary: TrustBoundaryKernelRuntimeLayer},
		AttributionConfidence: RuntimeSubstrateConfidenceBoundedCorrelation,
		FreshnessState:        RuntimeSubstrateFreshnessFresh,
	}, runtimeSubstrateNow)
	if err := ValidateRuntimeSubstrateObservedEvent(event); err == nil {
		t.Fatal("expected missing process identity validation error")
	}
}

func runtimeSubstrateNow() time.Time {
	return time.Date(2026, time.April, 22, 12, 0, 0, 0, time.UTC)
}

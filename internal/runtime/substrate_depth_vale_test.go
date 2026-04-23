package runtime

import (
	"testing"
	"time"
)

func TestRuntimeSubstrateValEStateRequiresActiveValD(t *testing.T) {
	if got := EvaluateRuntimeSubstrateValEState(
		RuntimeSubstrateValDStateSubstantial,
		RuntimeSubstrateValELatencyPackStateActive,
		RuntimeSubstrateValEFalsePositiveBudgetStateActive,
		RuntimeSubstrateValEReplayableBenchmarkPackStateActive,
		RuntimeSubstrateValEPerformanceGateStateActive,
	); got != RuntimeSubstrateValEStateIncomplete {
		t.Fatalf("expected incomplete val e without active val d, got %q", got)
	}
}

func TestRuntimeSubstrateValELatencyPackIsPartialWithoutP99(t *testing.T) {
	items := fullyMeasuredValELatencyItems()
	items[0].CaptureP99Micros = 0
	if got := EvaluateRuntimeSubstrateValELatencyPackState(items); got != RuntimeSubstrateValELatencyPackStatePartial {
		t.Fatalf("expected partial latency pack without p99, got %q", got)
	}
}

func TestRuntimeSubstrateValEFalsePositiveBudgetIsPartialWithoutMeasurementSource(t *testing.T) {
	items := fullyMeasuredValEFalsePositiveBudgetItems()
	items[1].MeasurementSource = ""
	if got := EvaluateRuntimeSubstrateValEFalsePositiveBudgetState(items); got != RuntimeSubstrateValEFalsePositiveBudgetStatePartial {
		t.Fatalf("expected partial false-positive budget without measurement source, got %q", got)
	}
}

func TestRuntimeSubstrateValEReplayableBenchmarkPackIsPartialWithoutMethodology(t *testing.T) {
	items := fullyMeasuredValEReplayableBenchmarkPacks()
	items[0].MethodologyRef = ""
	if got := EvaluateRuntimeSubstrateValEReplayableBenchmarkPackState(items); got != RuntimeSubstrateValEReplayableBenchmarkPackStatePartial {
		t.Fatalf("expected partial replayable benchmark pack without methodology ref, got %q", got)
	}
}

func TestRuntimeSubstrateValESurfacesAreActiveWithMeasuredPack(t *testing.T) {
	latency := fullyMeasuredValELatencyItems()
	if got := EvaluateRuntimeSubstrateValELatencyPackState(latency); got != RuntimeSubstrateValELatencyPackStateActive {
		t.Fatalf("expected active latency pack state, got %q", got)
	}

	falsePositive := fullyMeasuredValEFalsePositiveBudgetItems()
	if got := EvaluateRuntimeSubstrateValEFalsePositiveBudgetState(falsePositive); got != RuntimeSubstrateValEFalsePositiveBudgetStateActive {
		t.Fatalf("expected active false-positive budget state, got %q", got)
	}

	replayable := fullyMeasuredValEReplayableBenchmarkPacks()
	if got := EvaluateRuntimeSubstrateValEReplayableBenchmarkPackState(replayable); got != RuntimeSubstrateValEReplayableBenchmarkPackStateActive {
		t.Fatalf("expected active replayable benchmark pack state, got %q", got)
	}

	performanceGate := fullyMeasuredValEPerformanceGates()
	if got := EvaluateRuntimeSubstrateValEPerformanceGateState(performanceGate); got != RuntimeSubstrateValEPerformanceGateStateActive {
		t.Fatalf("expected active performance gate state, got %q", got)
	}

	if got := EvaluateRuntimeSubstrateValEState(
		RuntimeSubstrateValDStateActive,
		RuntimeSubstrateValELatencyPackStateActive,
		RuntimeSubstrateValEFalsePositiveBudgetStateActive,
		RuntimeSubstrateValEReplayableBenchmarkPackStateActive,
		RuntimeSubstrateValEPerformanceGateStateActive,
	); got != RuntimeSubstrateValEStateActive {
		t.Fatalf("expected active val e state, got %q", got)
	}
}

func fullyMeasuredValELatencyItems() []RuntimeSubstrateExecutionClassLatencyPackItem {
	measuredAt := time.Date(2026, time.April, 23, 13, 0, 0, 0, time.UTC)
	return []RuntimeSubstrateExecutionClassLatencyPackItem{
		{
			ExecutionClass:                     RuntimeExecutionClassStandardNode,
			CurrentState:                       RuntimeSubstrateExecutionClassStateSupported,
			MeasurementBasis:                   "runtime_capture_and_response_pack",
			MeasuredAt:                         measuredAt,
			MeasurementSource:                  "runtime_substrate_vale_latency.standard_node.v1",
			MethodologyRef:                     "/v1/public/benchmarks/methodology",
			EvidenceRefs:                       []string{"/v1/runtime/substrate-depth/vald/overhead-visibility", "/v1/foundation/execution/benchmarks/evaluate"},
			CaptureP50Micros:                   120,
			CaptureP95Micros:                   210,
			CaptureP99Micros:                   340,
			CorrelationP50Micros:               240,
			CorrelationP95Micros:               410,
			CorrelationP99Micros:               580,
			EnforcementDecisionP50Micros:       300,
			EnforcementDecisionP95Micros:       520,
			EnforcementDecisionP99Micros:       760,
			EndToEndP50Micros:                  660,
			EndToEndP95Micros:                  1140,
			EndToEndP99Micros:                  1620,
			CaptureBudgetP99Micros:             400,
			CorrelationBudgetP99Micros:         700,
			EnforcementDecisionBudgetP99Micros: 900,
			EndToEndBudgetP99Micros:            1800,
		},
		{
			ExecutionClass:                     RuntimeExecutionClassHardenedNode,
			CurrentState:                       RuntimeSubstrateExecutionClassStateSupported,
			MeasurementBasis:                   "runtime_capture_and_response_pack",
			MeasuredAt:                         measuredAt.Add(10 * time.Minute),
			MeasurementSource:                  "runtime_substrate_vale_latency.hardened_node.v1",
			MethodologyRef:                     "/v1/public/benchmarks/methodology",
			EvidenceRefs:                       []string{"/v1/runtime/substrate-depth/vald/overhead-visibility", "/v1/foundation/execution/benchmarks/evaluate"},
			CaptureP50Micros:                   140,
			CaptureP95Micros:                   250,
			CaptureP99Micros:                   390,
			CorrelationP50Micros:               280,
			CorrelationP95Micros:               450,
			CorrelationP99Micros:               620,
			EnforcementDecisionP50Micros:       340,
			EnforcementDecisionP95Micros:       560,
			EnforcementDecisionP99Micros:       800,
			EndToEndP50Micros:                  740,
			EndToEndP95Micros:                  1230,
			EndToEndP99Micros:                  1710,
			CaptureBudgetP99Micros:             450,
			CorrelationBudgetP99Micros:         750,
			EnforcementDecisionBudgetP99Micros: 950,
			EndToEndBudgetP99Micros:            1900,
		},
		{
			ExecutionClass:                     RuntimeExecutionClassConfidentialCapableNode,
			CurrentState:                       RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementBasis:                   "runtime_capture_and_response_pack",
			MeasuredAt:                         measuredAt.Add(20 * time.Minute),
			MeasurementSource:                  "runtime_substrate_vale_latency.confidential_capable_node.v1",
			MethodologyRef:                     "/v1/public/benchmarks/methodology",
			EvidenceRefs:                       []string{"/v1/runtime/substrate-depth/vald/overhead-visibility", "/v1/foundation/execution/benchmarks/evaluate"},
			CaptureP50Micros:                   180,
			CaptureP95Micros:                   320,
			CaptureP99Micros:                   470,
			CorrelationP50Micros:               320,
			CorrelationP95Micros:               520,
			CorrelationP99Micros:               720,
			EnforcementDecisionP50Micros:       390,
			EnforcementDecisionP95Micros:       640,
			EnforcementDecisionP99Micros:       910,
			EndToEndP50Micros:                  860,
			EndToEndP95Micros:                  1440,
			EndToEndP99Micros:                  1980,
			CaptureBudgetP99Micros:             520,
			CorrelationBudgetP99Micros:         780,
			EnforcementDecisionBudgetP99Micros: 980,
			EndToEndBudgetP99Micros:            2100,
		},
		{
			ExecutionClass:                     RuntimeExecutionClassVMBackedNode,
			CurrentState:                       RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementBasis:                   "runtime_capture_and_response_pack",
			MeasuredAt:                         measuredAt.Add(30 * time.Minute),
			MeasurementSource:                  "runtime_substrate_vale_latency.vm_backed_node.v1",
			MethodologyRef:                     "/v1/public/benchmarks/methodology",
			EvidenceRefs:                       []string{"/v1/runtime/substrate-depth/vald/overhead-visibility", "/v1/foundation/execution/benchmarks/evaluate"},
			CaptureP50Micros:                   170,
			CaptureP95Micros:                   310,
			CaptureP99Micros:                   450,
			CorrelationP50Micros:               340,
			CorrelationP95Micros:               550,
			CorrelationP99Micros:               760,
			EnforcementDecisionP50Micros:       410,
			EnforcementDecisionP95Micros:       670,
			EnforcementDecisionP99Micros:       940,
			EndToEndP50Micros:                  920,
			EndToEndP95Micros:                  1500,
			EndToEndP99Micros:                  2060,
			CaptureBudgetP99Micros:             500,
			CorrelationBudgetP99Micros:         820,
			EnforcementDecisionBudgetP99Micros: 1020,
			EndToEndBudgetP99Micros:            2200,
		},
		{
			ExecutionClass:                     RuntimeExecutionClassOfflineAirgappedNode,
			CurrentState:                       RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementBasis:                   "runtime_capture_and_response_pack",
			MeasuredAt:                         measuredAt.Add(40 * time.Minute),
			MeasurementSource:                  "runtime_substrate_vale_latency.offline_airgapped_node.v1",
			MethodologyRef:                     "/v1/public/benchmarks/methodology",
			EvidenceRefs:                       []string{"/v1/runtime/substrate-depth/vald/overhead-visibility", "/v1/foundation/execution/benchmarks/evaluate"},
			CaptureP50Micros:                   150,
			CaptureP95Micros:                   260,
			CaptureP99Micros:                   390,
			CorrelationP50Micros:               310,
			CorrelationP95Micros:               500,
			CorrelationP99Micros:               700,
			EnforcementDecisionP50Micros:       360,
			EnforcementDecisionP95Micros:       600,
			EnforcementDecisionP99Micros:       860,
			EndToEndP50Micros:                  800,
			EndToEndP95Micros:                  1350,
			EndToEndP99Micros:                  1880,
			CaptureBudgetP99Micros:             440,
			CorrelationBudgetP99Micros:         760,
			EnforcementDecisionBudgetP99Micros: 940,
			EndToEndBudgetP99Micros:            2050,
		},
	}
}

func fullyMeasuredValEFalsePositiveBudgetItems() []RuntimeSubstrateExecutionClassFalsePositiveBudgetItem {
	measuredAt := time.Date(2026, time.April, 23, 14, 0, 0, 0, time.UTC)
	return []RuntimeSubstrateExecutionClassFalsePositiveBudgetItem{
		{ExecutionClass: RuntimeExecutionClassStandardNode, CurrentState: RuntimeSubstrateExecutionClassStateSupported, MeasurementBasis: "runtime_findings_review_window", MeasuredAt: measuredAt, MeasurementSource: "runtime_substrate_vale_false_positive.standard_node.v1", EvidenceRefs: []string{"/v1/runtime/findings", "/v1/runtime/substrate-depth/vala/observability"}, ObservationWindow: "14d", DetectionCount: 320, FalsePositiveCount: 3, FalsePositiveRatePct: 0.94, AllowedFalsePositiveRatePct: 2.00, BudgetState: "within_budget"},
		{ExecutionClass: RuntimeExecutionClassHardenedNode, CurrentState: RuntimeSubstrateExecutionClassStateSupported, MeasurementBasis: "runtime_findings_review_window", MeasuredAt: measuredAt.Add(10 * time.Minute), MeasurementSource: "runtime_substrate_vale_false_positive.hardened_node.v1", EvidenceRefs: []string{"/v1/runtime/findings", "/v1/runtime/substrate-depth/vala/observability"}, ObservationWindow: "14d", DetectionCount: 290, FalsePositiveCount: 2, FalsePositiveRatePct: 0.69, AllowedFalsePositiveRatePct: 1.80, BudgetState: "within_budget"},
		{ExecutionClass: RuntimeExecutionClassConfidentialCapableNode, CurrentState: RuntimeSubstrateExecutionClassStateDegraded, MeasurementBasis: "runtime_findings_review_window", MeasuredAt: measuredAt.Add(20 * time.Minute), MeasurementSource: "runtime_substrate_vale_false_positive.confidential_capable_node.v1", EvidenceRefs: []string{"/v1/runtime/findings", "/v1/runtime/substrate-depth/vala/observability"}, ObservationWindow: "14d", DetectionCount: 180, FalsePositiveCount: 2, FalsePositiveRatePct: 1.11, AllowedFalsePositiveRatePct: 2.20, BudgetState: "within_budget"},
		{ExecutionClass: RuntimeExecutionClassVMBackedNode, CurrentState: RuntimeSubstrateExecutionClassStateDegraded, MeasurementBasis: "runtime_findings_review_window", MeasuredAt: measuredAt.Add(30 * time.Minute), MeasurementSource: "runtime_substrate_vale_false_positive.vm_backed_node.v1", EvidenceRefs: []string{"/v1/runtime/findings", "/v1/runtime/substrate-depth/vala/observability"}, ObservationWindow: "14d", DetectionCount: 210, FalsePositiveCount: 3, FalsePositiveRatePct: 1.43, AllowedFalsePositiveRatePct: 2.40, BudgetState: "within_budget"},
		{ExecutionClass: RuntimeExecutionClassOfflineAirgappedNode, CurrentState: RuntimeSubstrateExecutionClassStateDegraded, MeasurementBasis: "runtime_findings_review_window", MeasuredAt: measuredAt.Add(40 * time.Minute), MeasurementSource: "runtime_substrate_vale_false_positive.offline_airgapped_node.v1", EvidenceRefs: []string{"/v1/runtime/findings", "/v1/runtime/substrate-depth/vala/observability"}, ObservationWindow: "14d", DetectionCount: 140, FalsePositiveCount: 2, FalsePositiveRatePct: 1.43, AllowedFalsePositiveRatePct: 2.50, BudgetState: "within_budget"},
	}
}

func fullyMeasuredValEReplayableBenchmarkPacks() []RuntimeSubstrateReplayableBenchmarkPackItem {
	replayedAt := time.Date(2026, time.April, 23, 15, 0, 0, 0, time.UTC)
	executionClasses := []string{RuntimeExecutionClassStandardNode, RuntimeExecutionClassHardenedNode, RuntimeExecutionClassConfidentialCapableNode, RuntimeExecutionClassVMBackedNode, RuntimeExecutionClassOfflineAirgappedNode}
	commandHints := []string{
		"go test ./internal/runtime -run '^$' -bench BenchmarkCompare -benchmem",
		"go test ./services/audit-writer -run '^$' -bench 'BenchmarkAuditWriter(TopologyBlastRadius|ForensicsState|RuntimeFindings)' -benchmem",
		"go test ./services/audit-writer -run '^$' -bench 'BenchmarkAuditWriter(HandoffSeal|HandoffVerify|FederationProofVerify|ValidationExecute)' -benchmem",
	}
	return []RuntimeSubstrateReplayableBenchmarkPackItem{
		{PackID: "runtime_substrate_vale_local_baseline", CurrentState: "replayable_pack_ready", ProfileID: "local_baseline", MethodologyRef: "/v1/public/benchmarks/methodology", HarnessRef: "/v1/foundation/execution/benchmarks/harness", EvaluationRef: "/v1/foundation/execution/benchmarks/evaluate", ReplayedAt: replayedAt, Replayable: true, ExecutionClasses: executionClasses, CommandHints: commandHints, MeasuredOutputs: []string{"capture_p50_p95_p99=reported", "false_positive_budget=reported", "performance_gate=passed"}, EvidenceRefs: []string{"/v1/public/benchmarks/packs", "/v1/public/benchmarks/methodology", "/v1/foundation/execution/benchmarks/harness"}},
		{PackID: "runtime_substrate_vale_production_like", CurrentState: "replayable_pack_ready", ProfileID: "production_like", MethodologyRef: "/v1/public/benchmarks/methodology", HarnessRef: "/v1/foundation/execution/benchmarks/harness", EvaluationRef: "/v1/foundation/execution/benchmarks/evaluate", ReplayedAt: replayedAt.Add(10 * time.Minute), Replayable: true, ExecutionClasses: executionClasses, CommandHints: commandHints, MeasuredOutputs: []string{"capture_p50_p95_p99=reported", "false_positive_budget=reported", "performance_gate=passed"}, EvidenceRefs: []string{"/v1/public/benchmarks/packs", "/v1/public/benchmarks/methodology", "/v1/foundation/execution/benchmarks/harness"}},
		{PackID: "runtime_substrate_vale_stress", CurrentState: "replayable_pack_ready", ProfileID: "stress", MethodologyRef: "/v1/public/benchmarks/methodology", HarnessRef: "/v1/foundation/execution/benchmarks/harness", EvaluationRef: "/v1/foundation/execution/benchmarks/evaluate", ReplayedAt: replayedAt.Add(20 * time.Minute), Replayable: true, ExecutionClasses: executionClasses, CommandHints: commandHints, MeasuredOutputs: []string{"capture_p50_p95_p99=reported", "false_positive_budget=reported", "performance_gate=passed"}, EvidenceRefs: []string{"/v1/public/benchmarks/packs", "/v1/public/benchmarks/methodology", "/v1/foundation/execution/benchmarks/harness"}},
	}
}

func fullyMeasuredValEPerformanceGates() []RuntimeSubstratePerformanceGateItem {
	measuredAt := time.Date(2026, time.April, 23, 16, 0, 0, 0, time.UTC)
	dimensions := []string{"capture_latency", "correlation_latency", "enforcement_decision_latency", "end_to_end_latency", "false_positive_budget"}
	return []RuntimeSubstratePerformanceGateItem{
		{GateID: "runtime_substrate_vale_local_baseline", CurrentState: "performance_gate_passed", ProfileID: "local_baseline", EvaluationState: "passed", EvaluationRef: "/v1/foundation/execution/benchmarks/evaluate", MeasuredAt: measuredAt, ObservationCount: 4, GatedDimensions: dimensions, EvidenceRefs: []string{"/v1/foundation/execution/benchmarks/evaluate", "/v1/runtime/substrate-depth/vale/latency-pack", "/v1/runtime/substrate-depth/vale/false-positive-budget"}},
		{GateID: "runtime_substrate_vale_production_like", CurrentState: "performance_gate_passed", ProfileID: "production_like", EvaluationState: "passed", EvaluationRef: "/v1/foundation/execution/benchmarks/evaluate", MeasuredAt: measuredAt.Add(10 * time.Minute), ObservationCount: 4, GatedDimensions: dimensions, EvidenceRefs: []string{"/v1/foundation/execution/benchmarks/evaluate", "/v1/runtime/substrate-depth/vale/latency-pack", "/v1/runtime/substrate-depth/vale/false-positive-budget"}},
		{GateID: "runtime_substrate_vale_stress", CurrentState: "performance_gate_passed", ProfileID: "stress", EvaluationState: "passed", EvaluationRef: "/v1/foundation/execution/benchmarks/evaluate", MeasuredAt: measuredAt.Add(20 * time.Minute), ObservationCount: 4, GatedDimensions: dimensions, EvidenceRefs: []string{"/v1/foundation/execution/benchmarks/evaluate", "/v1/runtime/substrate-depth/vale/latency-pack", "/v1/runtime/substrate-depth/vale/false-positive-budget"}},
	}
}

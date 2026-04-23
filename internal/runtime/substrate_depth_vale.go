package runtime

import (
	"strings"
	"time"
)

const (
	RuntimeSubstrateValELatencyPackStateActive     = "runtime_substrate_vale_latency_pack_active"
	RuntimeSubstrateValELatencyPackStatePartial    = "runtime_substrate_vale_latency_pack_partial"
	RuntimeSubstrateValELatencyPackStateIncomplete = "runtime_substrate_vale_latency_pack_incomplete"

	RuntimeSubstrateValEFalsePositiveBudgetStateActive     = "runtime_substrate_vale_false_positive_budget_active"
	RuntimeSubstrateValEFalsePositiveBudgetStatePartial    = "runtime_substrate_vale_false_positive_budget_partial"
	RuntimeSubstrateValEFalsePositiveBudgetStateIncomplete = "runtime_substrate_vale_false_positive_budget_incomplete"

	RuntimeSubstrateValEReplayableBenchmarkPackStateActive     = "runtime_substrate_vale_replayable_benchmark_pack_active"
	RuntimeSubstrateValEReplayableBenchmarkPackStatePartial    = "runtime_substrate_vale_replayable_benchmark_pack_partial"
	RuntimeSubstrateValEReplayableBenchmarkPackStateIncomplete = "runtime_substrate_vale_replayable_benchmark_pack_incomplete"

	RuntimeSubstrateValEPerformanceGateStateActive     = "runtime_substrate_vale_performance_gate_active"
	RuntimeSubstrateValEPerformanceGateStatePartial    = "runtime_substrate_vale_performance_gate_partial"
	RuntimeSubstrateValEPerformanceGateStateIncomplete = "runtime_substrate_vale_performance_gate_incomplete"

	RuntimeSubstrateValEStateIncomplete  = "runtime_substrate_vale_incomplete"
	RuntimeSubstrateValEStateSubstantial = "runtime_substrate_vale_substantially_ready"
	RuntimeSubstrateValEStateActive      = "runtime_substrate_vale_active"
)

type RuntimeSubstrateExecutionClassLatencyPackItem struct {
	ExecutionClass                     string    `json:"execution_class"`
	CurrentState                       string    `json:"current_state"`
	MeasurementBasis                   string    `json:"measurement_basis,omitempty"`
	MeasuredAt                         time.Time `json:"measured_at,omitempty"`
	MeasurementSource                  string    `json:"measurement_source,omitempty"`
	MethodologyRef                     string    `json:"methodology_ref,omitempty"`
	EvidenceRefs                       []string  `json:"evidence_refs,omitempty"`
	CaptureP50Micros                   int       `json:"capture_p50_micros,omitempty"`
	CaptureP95Micros                   int       `json:"capture_p95_micros,omitempty"`
	CaptureP99Micros                   int       `json:"capture_p99_micros,omitempty"`
	CorrelationP50Micros               int       `json:"correlation_p50_micros,omitempty"`
	CorrelationP95Micros               int       `json:"correlation_p95_micros,omitempty"`
	CorrelationP99Micros               int       `json:"correlation_p99_micros,omitempty"`
	EnforcementDecisionP50Micros       int       `json:"enforcement_decision_p50_micros,omitempty"`
	EnforcementDecisionP95Micros       int       `json:"enforcement_decision_p95_micros,omitempty"`
	EnforcementDecisionP99Micros       int       `json:"enforcement_decision_p99_micros,omitempty"`
	EndToEndP50Micros                  int       `json:"end_to_end_p50_micros,omitempty"`
	EndToEndP95Micros                  int       `json:"end_to_end_p95_micros,omitempty"`
	EndToEndP99Micros                  int       `json:"end_to_end_p99_micros,omitempty"`
	CaptureBudgetP99Micros             int       `json:"capture_budget_p99_micros,omitempty"`
	CorrelationBudgetP99Micros         int       `json:"correlation_budget_p99_micros,omitempty"`
	EnforcementDecisionBudgetP99Micros int       `json:"enforcement_decision_budget_p99_micros,omitempty"`
	EndToEndBudgetP99Micros            int       `json:"end_to_end_budget_p99_micros,omitempty"`
	Limitations                        []string  `json:"limitations,omitempty"`
}

type RuntimeSubstrateExecutionClassFalsePositiveBudgetItem struct {
	ExecutionClass              string    `json:"execution_class"`
	CurrentState                string    `json:"current_state"`
	MeasurementBasis            string    `json:"measurement_basis,omitempty"`
	MeasuredAt                  time.Time `json:"measured_at,omitempty"`
	MeasurementSource           string    `json:"measurement_source,omitempty"`
	EvidenceRefs                []string  `json:"evidence_refs,omitempty"`
	ObservationWindow           string    `json:"observation_window,omitempty"`
	DetectionCount              int       `json:"detection_count,omitempty"`
	FalsePositiveCount          int       `json:"false_positive_count,omitempty"`
	FalsePositiveRatePct        float64   `json:"false_positive_rate_pct,omitempty"`
	AllowedFalsePositiveRatePct float64   `json:"allowed_false_positive_rate_pct,omitempty"`
	BudgetState                 string    `json:"budget_state,omitempty"`
	Limitations                 []string  `json:"limitations,omitempty"`
}

type RuntimeSubstrateReplayableBenchmarkPackItem struct {
	PackID           string    `json:"pack_id"`
	CurrentState     string    `json:"current_state"`
	ProfileID        string    `json:"profile_id"`
	MethodologyRef   string    `json:"methodology_ref,omitempty"`
	HarnessRef       string    `json:"harness_ref,omitempty"`
	EvaluationRef    string    `json:"evaluation_ref,omitempty"`
	ReplayedAt       time.Time `json:"replayed_at,omitempty"`
	Replayable       bool      `json:"replayable"`
	ExecutionClasses []string  `json:"execution_classes,omitempty"`
	CommandHints     []string  `json:"command_hints,omitempty"`
	MeasuredOutputs  []string  `json:"measured_outputs,omitempty"`
	EvidenceRefs     []string  `json:"evidence_refs,omitempty"`
	Limitations      []string  `json:"limitations,omitempty"`
}

type RuntimeSubstratePerformanceGateItem struct {
	GateID            string    `json:"gate_id"`
	CurrentState      string    `json:"current_state"`
	ProfileID         string    `json:"profile_id"`
	EvaluationState   string    `json:"evaluation_state,omitempty"`
	EvaluationRef     string    `json:"evaluation_ref,omitempty"`
	MeasuredAt        time.Time `json:"measured_at,omitempty"`
	ObservationCount  int       `json:"observation_count,omitempty"`
	GatedDimensions   []string  `json:"gated_dimensions,omitempty"`
	EvidenceRefs      []string  `json:"evidence_refs,omitempty"`
	OverridePermitted bool      `json:"override_permitted"`
	Limitations       []string  `json:"limitations,omitempty"`
}

func RuntimeSubstrateValERemainingDeferredScope() []string {
	return nil
}

func EvaluateRuntimeSubstrateValELatencyPackState(items []RuntimeSubstrateExecutionClassLatencyPackItem) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.ExecutionClass) == "" || strings.TrimSpace(item.CurrentState) == "" {
					return false
				}
				if strings.TrimSpace(item.MethodologyRef) == "" {
					return false
				}
			}
			return true
		},
		func() bool {
			expected := map[string]struct{}{
				RuntimeExecutionClassStandardNode:            {},
				RuntimeExecutionClassHardenedNode:            {},
				RuntimeExecutionClassConfidentialCapableNode: {},
				RuntimeExecutionClassVMBackedNode:            {},
				RuntimeExecutionClassOfflineAirgappedNode:    {},
			}
			for _, item := range items {
				if strings.TrimSpace(item.MeasurementBasis) == "" ||
					item.MeasuredAt.IsZero() ||
					strings.TrimSpace(item.MeasurementSource) == "" ||
					len(item.EvidenceRefs) == 0 ||
					!runtimeSubstrateValELatencyPercentilesReady(item) ||
					!runtimeSubstrateValELatencyWithinBudget(item) {
					return false
				}
				delete(expected, strings.TrimSpace(item.ExecutionClass))
			}
			return len(expected) == 0
		},
		RuntimeSubstrateValELatencyPackStateIncomplete,
		RuntimeSubstrateValELatencyPackStatePartial,
		RuntimeSubstrateValELatencyPackStateActive,
	)
}

func EvaluateRuntimeSubstrateValEFalsePositiveBudgetState(items []RuntimeSubstrateExecutionClassFalsePositiveBudgetItem) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.ExecutionClass) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ObservationWindow) == "" {
					return false
				}
			}
			return true
		},
		func() bool {
			expected := map[string]struct{}{
				RuntimeExecutionClassStandardNode:            {},
				RuntimeExecutionClassHardenedNode:            {},
				RuntimeExecutionClassConfidentialCapableNode: {},
				RuntimeExecutionClassVMBackedNode:            {},
				RuntimeExecutionClassOfflineAirgappedNode:    {},
			}
			for _, item := range items {
				if strings.TrimSpace(item.MeasurementBasis) == "" ||
					item.MeasuredAt.IsZero() ||
					strings.TrimSpace(item.MeasurementSource) == "" ||
					len(item.EvidenceRefs) == 0 ||
					item.DetectionCount <= 0 ||
					item.FalsePositiveCount < 0 ||
					item.FalsePositiveRatePct < 0 ||
					item.AllowedFalsePositiveRatePct <= 0 ||
					strings.TrimSpace(item.BudgetState) != "within_budget" ||
					item.FalsePositiveCount > item.DetectionCount ||
					item.FalsePositiveRatePct > item.AllowedFalsePositiveRatePct {
					return false
				}
				delete(expected, strings.TrimSpace(item.ExecutionClass))
			}
			return len(expected) == 0
		},
		RuntimeSubstrateValEFalsePositiveBudgetStateIncomplete,
		RuntimeSubstrateValEFalsePositiveBudgetStatePartial,
		RuntimeSubstrateValEFalsePositiveBudgetStateActive,
	)
}

func EvaluateRuntimeSubstrateValEReplayableBenchmarkPackState(items []RuntimeSubstrateReplayableBenchmarkPackItem) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.PackID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ProfileID) == "" {
					return false
				}
			}
			return true
		},
		func() bool {
			expected := map[string]struct{}{
				"local_baseline":  {},
				"production_like": {},
				"stress":          {},
			}
			for _, item := range items {
				if strings.TrimSpace(item.MethodologyRef) == "" ||
					strings.TrimSpace(item.HarnessRef) == "" ||
					strings.TrimSpace(item.EvaluationRef) == "" ||
					item.ReplayedAt.IsZero() ||
					!item.Replayable ||
					len(item.ExecutionClasses) != 5 ||
					len(item.CommandHints) == 0 ||
					len(item.MeasuredOutputs) == 0 ||
					len(item.EvidenceRefs) == 0 {
					return false
				}
				delete(expected, strings.TrimSpace(item.ProfileID))
			}
			return len(expected) == 0
		},
		RuntimeSubstrateValEReplayableBenchmarkPackStateIncomplete,
		RuntimeSubstrateValEReplayableBenchmarkPackStatePartial,
		RuntimeSubstrateValEReplayableBenchmarkPackStateActive,
	)
}

func EvaluateRuntimeSubstrateValEPerformanceGateState(items []RuntimeSubstratePerformanceGateItem) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.GateID) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.ProfileID) == "" {
					return false
				}
			}
			return true
		},
		func() bool {
			expected := map[string]struct{}{
				"local_baseline":  {},
				"production_like": {},
				"stress":          {},
			}
			for _, item := range items {
				if strings.TrimSpace(item.EvaluationState) != "passed" ||
					strings.TrimSpace(item.EvaluationRef) == "" ||
					item.MeasuredAt.IsZero() ||
					item.ObservationCount <= 0 ||
					len(item.EvidenceRefs) == 0 ||
					!runtimeSubstrateValEHasGatedDimensions(item.GatedDimensions) ||
					item.OverridePermitted {
					return false
				}
				delete(expected, strings.TrimSpace(item.ProfileID))
			}
			return len(expected) == 0
		},
		RuntimeSubstrateValEPerformanceGateStateIncomplete,
		RuntimeSubstrateValEPerformanceGateStatePartial,
		RuntimeSubstrateValEPerformanceGateStateActive,
	)
}

func EvaluateRuntimeSubstrateValEState(valDState, latencyPackState, falsePositiveBudgetState, replayableBenchmarkPackState, performanceGateState string) string {
	if strings.TrimSpace(valDState) != RuntimeSubstrateValDStateActive {
		return RuntimeSubstrateValEStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(latencyPackState),
		strings.TrimSpace(falsePositiveBudgetState),
		strings.TrimSpace(replayableBenchmarkPackState),
		strings.TrimSpace(performanceGateState),
	} {
		switch state {
		case RuntimeSubstrateValELatencyPackStateActive,
			RuntimeSubstrateValEFalsePositiveBudgetStateActive,
			RuntimeSubstrateValEReplayableBenchmarkPackStateActive,
			RuntimeSubstrateValEPerformanceGateStateActive:
		case RuntimeSubstrateValELatencyPackStatePartial,
			RuntimeSubstrateValEFalsePositiveBudgetStatePartial,
			RuntimeSubstrateValEReplayableBenchmarkPackStatePartial,
			RuntimeSubstrateValEPerformanceGateStatePartial:
			hasPartial = true
		default:
			return RuntimeSubstrateValEStateIncomplete
		}
	}
	if hasPartial {
		return RuntimeSubstrateValEStateSubstantial
	}
	return RuntimeSubstrateValEStateActive
}

func runtimeSubstrateValELatencyPercentilesReady(item RuntimeSubstrateExecutionClassLatencyPackItem) bool {
	return runtimeSubstrateValEPercentilesReady(item.CaptureP50Micros, item.CaptureP95Micros, item.CaptureP99Micros) &&
		runtimeSubstrateValEPercentilesReady(item.CorrelationP50Micros, item.CorrelationP95Micros, item.CorrelationP99Micros) &&
		runtimeSubstrateValEPercentilesReady(item.EnforcementDecisionP50Micros, item.EnforcementDecisionP95Micros, item.EnforcementDecisionP99Micros) &&
		runtimeSubstrateValEPercentilesReady(item.EndToEndP50Micros, item.EndToEndP95Micros, item.EndToEndP99Micros) &&
		item.CaptureBudgetP99Micros > 0 &&
		item.CorrelationBudgetP99Micros > 0 &&
		item.EnforcementDecisionBudgetP99Micros > 0 &&
		item.EndToEndBudgetP99Micros > 0
}

func runtimeSubstrateValELatencyWithinBudget(item RuntimeSubstrateExecutionClassLatencyPackItem) bool {
	return item.CaptureP99Micros <= item.CaptureBudgetP99Micros &&
		item.CorrelationP99Micros <= item.CorrelationBudgetP99Micros &&
		item.EnforcementDecisionP99Micros <= item.EnforcementDecisionBudgetP99Micros &&
		item.EndToEndP99Micros <= item.EndToEndBudgetP99Micros
}

func runtimeSubstrateValEPercentilesReady(p50, p95, p99 int) bool {
	return p50 > 0 && p95 >= p50 && p99 >= p95
}

func runtimeSubstrateValEHasGatedDimensions(values []string) bool {
	required := map[string]struct{}{
		"capture_latency":              {},
		"correlation_latency":          {},
		"enforcement_decision_latency": {},
		"end_to_end_latency":           {},
		"false_positive_budget":        {},
	}
	for _, value := range values {
		delete(required, strings.TrimSpace(value))
	}
	return len(required) == 0
}

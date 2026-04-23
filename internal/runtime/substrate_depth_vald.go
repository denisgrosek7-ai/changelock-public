package runtime

import (
	"strings"
	"time"
)

const (
	RuntimeSubstrateExecutionClassStateSupported   = "runtime_substrate_execution_class_supported"
	RuntimeSubstrateExecutionClassStateDegraded    = "runtime_substrate_execution_class_degraded"
	RuntimeSubstrateExecutionClassStateUnsupported = "runtime_substrate_execution_class_unsupported"

	RuntimeSubstrateValDExecutionClassMatrixStateActive     = "runtime_substrate_vald_execution_class_matrix_active"
	RuntimeSubstrateValDExecutionClassMatrixStatePartial    = "runtime_substrate_vald_execution_class_matrix_partial"
	RuntimeSubstrateValDExecutionClassMatrixStateIncomplete = "runtime_substrate_vald_execution_class_matrix_incomplete"

	RuntimeSubstrateValDSignalCoverageStateActive     = "runtime_substrate_vald_signal_coverage_active"
	RuntimeSubstrateValDSignalCoverageStatePartial    = "runtime_substrate_vald_signal_coverage_partial"
	RuntimeSubstrateValDSignalCoverageStateIncomplete = "runtime_substrate_vald_signal_coverage_incomplete"

	RuntimeSubstrateValDEnforcementAvailabilityStateActive     = "runtime_substrate_vald_enforcement_availability_active"
	RuntimeSubstrateValDEnforcementAvailabilityStatePartial    = "runtime_substrate_vald_enforcement_availability_partial"
	RuntimeSubstrateValDEnforcementAvailabilityStateIncomplete = "runtime_substrate_vald_enforcement_availability_incomplete"

	RuntimeSubstrateValDOverheadVisibilityStateActive     = "runtime_substrate_vald_overhead_visibility_active"
	RuntimeSubstrateValDOverheadVisibilityStatePartial    = "runtime_substrate_vald_overhead_visibility_partial"
	RuntimeSubstrateValDOverheadVisibilityStateIncomplete = "runtime_substrate_vald_overhead_visibility_incomplete"

	RuntimeSubstrateValDStateIncomplete  = "runtime_substrate_vald_incomplete"
	RuntimeSubstrateValDStateSubstantial = "runtime_substrate_vald_substantially_ready"
	RuntimeSubstrateValDStateActive      = "runtime_substrate_vald_active"
)

type RuntimeSubstrateExecutionClassMatrixItem struct {
	ExecutionClass          string   `json:"execution_class"`
	CurrentState            string   `json:"current_state"`
	ObservabilityState      string   `json:"observability_state"`
	CorrelationState        string   `json:"correlation_state"`
	EnforcementState        string   `json:"enforcement_state"`
	RequiredSignalFamilies  []string `json:"required_signal_families,omitempty"`
	DegradedCapabilities    []string `json:"degraded_capabilities,omitempty"`
	UnsupportedCapabilities []string `json:"unsupported_capabilities,omitempty"`
	CapabilityAssumptions   []string `json:"capability_assumptions,omitempty"`
	Limitations             []string `json:"limitations,omitempty"`
}

type RuntimeSubstrateExecutionClassSignalCoverageItem struct {
	ExecutionClass      string   `json:"execution_class"`
	CurrentState        string   `json:"current_state"`
	ObservedFamilies    []string `json:"observed_families,omitempty"`
	PartialFamilies     []string `json:"partial_families,omitempty"`
	UnsupportedFamilies []string `json:"unsupported_families,omitempty"`
	HookCoverageRefs    []string `json:"hook_coverage_refs,omitempty"`
	DegradedReasons     []string `json:"degraded_reasons,omitempty"`
}

type RuntimeSubstrateExecutionClassEnforcementAvailabilityItem struct {
	ExecutionClass         string   `json:"execution_class"`
	CurrentState           string   `json:"current_state"`
	SupportedActions       []string `json:"supported_actions,omitempty"`
	UnsupportedActions     []string `json:"unsupported_actions,omitempty"`
	SupportedDecisionModes []string `json:"supported_decision_modes,omitempty"`
	GuaranteeBoundaries    []string `json:"guarantee_boundaries,omitempty"`
	NonGuarantees          []string `json:"non_guarantees,omitempty"`
}

type RuntimeSubstrateExecutionClassOverheadVisibilityItem struct {
	ExecutionClass                   string    `json:"execution_class"`
	CurrentState                     string    `json:"current_state"`
	MeasurementStatus                string    `json:"measurement_status"`
	MeasurementBasis                 string    `json:"measurement_basis,omitempty"`
	MeasuredAt                       time.Time `json:"measured_at,omitempty"`
	MeasurementSource                string    `json:"measurement_source,omitempty"`
	EvidenceRefs                     []string  `json:"evidence_refs,omitempty"`
	BudgetClass                      string    `json:"budget_class"`
	BenchmarkRefs                    []string  `json:"benchmark_refs,omitempty"`
	MethodologyRefs                  []string  `json:"methodology_refs,omitempty"`
	ObservedCPUOverheadMillicores    int       `json:"observed_cpu_overhead_millicores,omitempty"`
	ObservedMemoryOverheadMiB        int       `json:"observed_memory_overhead_mib,omitempty"`
	ObservedCaptureLatencyMicros     int       `json:"observed_capture_latency_micros,omitempty"`
	ObservedCorrelationLatencyMicros int       `json:"observed_correlation_latency_micros,omitempty"`
	VisibilityRules                  []string  `json:"visibility_rules,omitempty"`
	Limitations                      []string  `json:"limitations,omitempty"`
}

func RuntimeSubstrateValDRemainingDeferredScope() []string {
	return []string{
		"performance_and_proof_pack",
	}
}

func EvaluateRuntimeSubstrateValDExecutionClassMatrixState(items []RuntimeSubstrateExecutionClassMatrixItem) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.ExecutionClass) == "" || strings.TrimSpace(item.CurrentState) == "" {
					return false
				}
				if len(item.RequiredSignalFamilies) == 0 || len(item.CapabilityAssumptions) == 0 {
					return false
				}
				if strings.TrimSpace(item.ObservabilityState) == "" || strings.TrimSpace(item.CorrelationState) == "" || strings.TrimSpace(item.EnforcementState) == "" {
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
				delete(expected, strings.TrimSpace(item.ExecutionClass))
			}
			return len(expected) == 0
		},
		RuntimeSubstrateValDExecutionClassMatrixStateIncomplete,
		RuntimeSubstrateValDExecutionClassMatrixStatePartial,
		RuntimeSubstrateValDExecutionClassMatrixStateActive,
	)
}

func EvaluateRuntimeSubstrateValDSignalCoverageState(items []RuntimeSubstrateExecutionClassSignalCoverageItem) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.ExecutionClass) == "" || strings.TrimSpace(item.CurrentState) == "" {
					return false
				}
				if len(item.ObservedFamilies) == 0 && len(item.PartialFamilies) == 0 && len(item.UnsupportedFamilies) == 0 {
					return false
				}
				if len(item.HookCoverageRefs) == 0 {
					return false
				}
			}
			return true
		},
		func() bool {
			for _, item := range items {
				total := len(item.ObservedFamilies) + len(item.PartialFamilies) + len(item.UnsupportedFamilies)
				if total < 4 {
					return false
				}
			}
			return true
		},
		RuntimeSubstrateValDSignalCoverageStateIncomplete,
		RuntimeSubstrateValDSignalCoverageStatePartial,
		RuntimeSubstrateValDSignalCoverageStateActive,
	)
}

func EvaluateRuntimeSubstrateValDEnforcementAvailabilityState(items []RuntimeSubstrateExecutionClassEnforcementAvailabilityItem) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.ExecutionClass) == "" || strings.TrimSpace(item.CurrentState) == "" {
					return false
				}
				if len(item.SupportedDecisionModes) == 0 || len(item.GuaranteeBoundaries) == 0 || len(item.NonGuarantees) == 0 {
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
				delete(expected, strings.TrimSpace(item.ExecutionClass))
			}
			return len(expected) == 0
		},
		RuntimeSubstrateValDEnforcementAvailabilityStateIncomplete,
		RuntimeSubstrateValDEnforcementAvailabilityStatePartial,
		RuntimeSubstrateValDEnforcementAvailabilityStateActive,
	)
}

func EvaluateRuntimeSubstrateValDOverheadVisibilityState(items []RuntimeSubstrateExecutionClassOverheadVisibilityItem) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.ExecutionClass) == "" || strings.TrimSpace(item.CurrentState) == "" {
					return false
				}
				if strings.TrimSpace(item.MeasurementStatus) == "" || strings.TrimSpace(item.BudgetClass) == "" {
					return false
				}
				if len(item.BenchmarkRefs) == 0 || len(item.MethodologyRefs) == 0 || len(item.VisibilityRules) == 0 {
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
				if !runtimeSubstrateValDOverheadMeasurementReady(item.MeasurementStatus) ||
					strings.TrimSpace(item.MeasurementBasis) == "" ||
					item.MeasuredAt.IsZero() ||
					strings.TrimSpace(item.MeasurementSource) == "" ||
					len(item.EvidenceRefs) == 0 ||
					!runtimeSubstrateValDHasConcreteOverheadFields(item) {
					return false
				}
				delete(expected, strings.TrimSpace(item.ExecutionClass))
			}
			return len(expected) == 0
		},
		RuntimeSubstrateValDOverheadVisibilityStateIncomplete,
		RuntimeSubstrateValDOverheadVisibilityStatePartial,
		RuntimeSubstrateValDOverheadVisibilityStateActive,
	)
}

func EvaluateRuntimeSubstrateValDState(valCState, matrixState, signalCoverageState, enforcementAvailabilityState, overheadVisibilityState string) string {
	if strings.TrimSpace(valCState) != RuntimeSubstrateValCStateActive {
		return RuntimeSubstrateValDStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(matrixState),
		strings.TrimSpace(signalCoverageState),
		strings.TrimSpace(enforcementAvailabilityState),
		strings.TrimSpace(overheadVisibilityState),
	} {
		switch state {
		case RuntimeSubstrateValDExecutionClassMatrixStateActive,
			RuntimeSubstrateValDSignalCoverageStateActive,
			RuntimeSubstrateValDEnforcementAvailabilityStateActive,
			RuntimeSubstrateValDOverheadVisibilityStateActive:
		case RuntimeSubstrateValDExecutionClassMatrixStatePartial,
			RuntimeSubstrateValDSignalCoverageStatePartial,
			RuntimeSubstrateValDEnforcementAvailabilityStatePartial,
			RuntimeSubstrateValDOverheadVisibilityStatePartial:
			hasPartial = true
		default:
			return RuntimeSubstrateValDStateIncomplete
		}
	}
	if hasPartial {
		return RuntimeSubstrateValDStateSubstantial
	}
	return RuntimeSubstrateValDStateActive
}

func runtimeSubstrateValDOverheadMeasurementReady(status string) bool {
	switch strings.TrimSpace(status) {
	case "class_specific_measurement_recorded", "class_specific_measurement_verified":
		return true
	default:
		return false
	}
}

func runtimeSubstrateValDHasConcreteOverheadFields(item RuntimeSubstrateExecutionClassOverheadVisibilityItem) bool {
	return item.ObservedCPUOverheadMillicores > 0 ||
		item.ObservedMemoryOverheadMiB > 0 ||
		item.ObservedCaptureLatencyMicros > 0 ||
		item.ObservedCorrelationLatencyMicros > 0
}

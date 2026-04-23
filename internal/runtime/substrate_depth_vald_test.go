package runtime

import (
	"testing"
	"time"
)

func TestRuntimeSubstrateValDStateRequiresActiveValC(t *testing.T) {
	if got := EvaluateRuntimeSubstrateValDState(
		RuntimeSubstrateValCStateSubstantial,
		RuntimeSubstrateValDExecutionClassMatrixStateActive,
		RuntimeSubstrateValDSignalCoverageStateActive,
		RuntimeSubstrateValDEnforcementAvailabilityStateActive,
		RuntimeSubstrateValDOverheadVisibilityStateActive,
	); got != RuntimeSubstrateValDStateIncomplete {
		t.Fatalf("expected incomplete val d without active val c, got %q", got)
	}
}

func TestRuntimeSubstrateValDOverheadVisibilityIsPartialWithoutMeasuredAt(t *testing.T) {
	items := fullyMeasuredValDOverheadItems()
	items[0].MeasuredAt = time.Time{}
	if got := EvaluateRuntimeSubstrateValDOverheadVisibilityState(items); got != RuntimeSubstrateValDOverheadVisibilityStatePartial {
		t.Fatalf("expected partial overhead visibility without measured_at, got %q", got)
	}
}

func TestRuntimeSubstrateValDOverheadVisibilityIsPartialWithoutMeasurementSource(t *testing.T) {
	items := fullyMeasuredValDOverheadItems()
	items[1].MeasurementSource = ""
	if got := EvaluateRuntimeSubstrateValDOverheadVisibilityState(items); got != RuntimeSubstrateValDOverheadVisibilityStatePartial {
		t.Fatalf("expected partial overhead visibility without measurement source, got %q", got)
	}
}

func TestRuntimeSubstrateValDOverheadVisibilityIsPartialWithoutEvidenceRefs(t *testing.T) {
	items := fullyMeasuredValDOverheadItems()
	items[1].EvidenceRefs = nil
	if got := EvaluateRuntimeSubstrateValDOverheadVisibilityState(items); got != RuntimeSubstrateValDOverheadVisibilityStatePartial {
		t.Fatalf("expected partial overhead visibility without evidence refs, got %q", got)
	}
}

func TestRuntimeSubstrateValDOverheadVisibilityIsPartialWithPlaceholderStatus(t *testing.T) {
	items := fullyMeasuredValDOverheadItems()
	items[2].MeasurementStatus = "class_specific_measurement_pending"
	if got := EvaluateRuntimeSubstrateValDOverheadVisibilityState(items); got != RuntimeSubstrateValDOverheadVisibilityStatePartial {
		t.Fatalf("expected partial overhead visibility with placeholder status, got %q", got)
	}
}

func TestRuntimeSubstrateValDSurfacesAreActiveWhenOverheadIsMeasured(t *testing.T) {
	matrix := []RuntimeSubstrateExecutionClassMatrixItem{
		{
			ExecutionClass:         RuntimeExecutionClassStandardNode,
			CurrentState:           RuntimeSubstrateExecutionClassStateSupported,
			ObservabilityState:     RuntimeSubstrateValASupportMatrixStateActive,
			CorrelationState:       RuntimeSubstrateValBStateActive,
			EnforcementState:       RuntimeSubstrateValCStateActive,
			RequiredSignalFamilies: []string{RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage, RuntimeSubstrateEventFamilyFileActivity, RuntimeSubstrateEventFamilyNetworkActivity},
			CapabilityAssumptions:  []string{"linux_tracepoint_exec_available"},
		},
		{
			ExecutionClass:         RuntimeExecutionClassHardenedNode,
			CurrentState:           RuntimeSubstrateExecutionClassStateSupported,
			ObservabilityState:     RuntimeSubstrateValASupportMatrixStateActive,
			CorrelationState:       RuntimeSubstrateValBStateActive,
			EnforcementState:       RuntimeSubstrateValCStateActive,
			RequiredSignalFamilies: []string{RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage, RuntimeSubstrateEventFamilyFileActivity, RuntimeSubstrateEventFamilyNetworkActivity},
			CapabilityAssumptions:  []string{"lsm_exec_hook_exposed"},
		},
		{
			ExecutionClass:         RuntimeExecutionClassConfidentialCapableNode,
			CurrentState:           RuntimeSubstrateExecutionClassStateDegraded,
			ObservabilityState:     RuntimeSubstrateValASupportMatrixStateActive,
			CorrelationState:       RuntimeSubstrateValBStateSubstantial,
			EnforcementState:       RuntimeSubstrateValCStateActive,
			RequiredSignalFamilies: []string{RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage, RuntimeSubstrateEventFamilyFileActivity, RuntimeSubstrateEventFamilyNetworkActivity},
			DegradedCapabilities:   []string{"capture_only_correlation_boundary"},
			CapabilityAssumptions:  []string{"guest_exec_visibility_present"},
		},
		{
			ExecutionClass:         RuntimeExecutionClassVMBackedNode,
			CurrentState:           RuntimeSubstrateExecutionClassStateDegraded,
			ObservabilityState:     RuntimeSubstrateValASupportMatrixStatePartial,
			CorrelationState:       RuntimeSubstrateValBStateSubstantial,
			EnforcementState:       RuntimeSubstrateValCStateActive,
			RequiredSignalFamilies: []string{RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage, RuntimeSubstrateEventFamilyFileActivity, RuntimeSubstrateEventFamilyNetworkActivity},
			DegradedCapabilities:   []string{"guest_only_file_and_network_context"},
			CapabilityAssumptions:  []string{"guest_kernel_signal_path_available"},
		},
		{
			ExecutionClass:          RuntimeExecutionClassOfflineAirgappedNode,
			CurrentState:            RuntimeSubstrateExecutionClassStateDegraded,
			ObservabilityState:      RuntimeSubstrateValASupportMatrixStatePartial,
			CorrelationState:        RuntimeSubstrateValBStateSubstantial,
			EnforcementState:        RuntimeSubstrateValCStateSubstantial,
			RequiredSignalFamilies:  []string{RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage, RuntimeSubstrateEventFamilyFileActivity, RuntimeSubstrateEventFamilyNetworkActivity},
			UnsupportedCapabilities: []string{RuntimeSubstrateEventFamilyNetworkActivity},
			CapabilityAssumptions:   []string{"offline_capture_pipeline_present"},
		},
	}
	if got := EvaluateRuntimeSubstrateValDExecutionClassMatrixState(matrix); got != RuntimeSubstrateValDExecutionClassMatrixStateActive {
		t.Fatalf("expected active execution class matrix, got %q", got)
	}

	coverage := []RuntimeSubstrateExecutionClassSignalCoverageItem{
		{ExecutionClass: RuntimeExecutionClassStandardNode, CurrentState: RuntimeSubstrateExecutionClassStateSupported, ObservedFamilies: []string{RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage, RuntimeSubstrateEventFamilyFileActivity, RuntimeSubstrateEventFamilyNetworkActivity}, HookCoverageRefs: []string{"tracepoint_exec", "tracepoint_tcp_connect"}},
		{ExecutionClass: RuntimeExecutionClassHardenedNode, CurrentState: RuntimeSubstrateExecutionClassStateSupported, ObservedFamilies: []string{RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage, RuntimeSubstrateEventFamilyFileActivity, RuntimeSubstrateEventFamilyNetworkActivity}, HookCoverageRefs: []string{"tracepoint_exec_and_lsm_exec", "tracepoint_tcp_connect"}},
		{ExecutionClass: RuntimeExecutionClassConfidentialCapableNode, CurrentState: RuntimeSubstrateExecutionClassStateDegraded, ObservedFamilies: []string{RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage, RuntimeSubstrateEventFamilyFileActivity, RuntimeSubstrateEventFamilyNetworkActivity}, HookCoverageRefs: []string{"tracepoint_exec", "tracepoint_tcp_connect"}, DegradedReasons: []string{"capture_only_confidential_boundary"}},
		{ExecutionClass: RuntimeExecutionClassVMBackedNode, CurrentState: RuntimeSubstrateExecutionClassStateDegraded, ObservedFamilies: []string{RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage}, PartialFamilies: []string{RuntimeSubstrateEventFamilyFileActivity, RuntimeSubstrateEventFamilyNetworkActivity}, HookCoverageRefs: []string{"guest_tracepoint_exec", "guest_tracepoint_tcp_connect"}, DegradedReasons: []string{"guest_only_context"}},
		{ExecutionClass: RuntimeExecutionClassOfflineAirgappedNode, CurrentState: RuntimeSubstrateExecutionClassStateDegraded, ObservedFamilies: []string{RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage, RuntimeSubstrateEventFamilyFileActivity}, UnsupportedFamilies: []string{RuntimeSubstrateEventFamilyNetworkActivity}, HookCoverageRefs: []string{"tracepoint_exec", "not_available_in_airgapped_capture_path"}, DegradedReasons: []string{"offline_network_signal_unavailable"}},
	}
	if got := EvaluateRuntimeSubstrateValDSignalCoverageState(coverage); got != RuntimeSubstrateValDSignalCoverageStateActive {
		t.Fatalf("expected active signal coverage state, got %q", got)
	}

	enforcement := []RuntimeSubstrateExecutionClassEnforcementAvailabilityItem{
		{ExecutionClass: RuntimeExecutionClassStandardNode, CurrentState: RuntimeSubstrateExecutionClassStateSupported, SupportedActions: []string{"runtime.apply_network_isolation", "runtime.restart_from_trusted_image", "hardening.block_exec_class_next_restart"}, SupportedDecisionModes: []string{RuntimeSubstrateDecisionModeImmediateContainment, RuntimeSubstrateDecisionModeNextRestartPreventive, RuntimeSubstrateDecisionModeTerminateAndRecover}, GuaranteeBoundaries: []string{"bounded_runtime_and_hardening_hooks_only"}, NonGuarantees: []string{"no_universal_prevention_claim"}},
		{ExecutionClass: RuntimeExecutionClassHardenedNode, CurrentState: RuntimeSubstrateExecutionClassStateSupported, SupportedActions: []string{"runtime.apply_network_isolation", "runtime.restart_from_trusted_image", "hardening.block_exec_class_next_restart"}, SupportedDecisionModes: []string{RuntimeSubstrateDecisionModeImmediateContainment, RuntimeSubstrateDecisionModeNextRestartPreventive, RuntimeSubstrateDecisionModeTerminateAndRecover}, GuaranteeBoundaries: []string{"bounded_runtime_and_hardening_hooks_only"}, NonGuarantees: []string{"no_universal_prevention_claim"}},
		{ExecutionClass: RuntimeExecutionClassConfidentialCapableNode, CurrentState: RuntimeSubstrateExecutionClassStateSupported, SupportedActions: []string{"runtime.apply_network_isolation", "runtime.restart_from_trusted_image", "hardening.block_exec_class_next_restart"}, SupportedDecisionModes: []string{RuntimeSubstrateDecisionModeImmediateContainment, RuntimeSubstrateDecisionModeNextRestartPreventive, RuntimeSubstrateDecisionModeTerminateAndRecover}, GuaranteeBoundaries: []string{"confidential_capture_does_not_change_enforcement_claims"}, NonGuarantees: []string{"no_absolute_truth_claim"}},
		{ExecutionClass: RuntimeExecutionClassVMBackedNode, CurrentState: RuntimeSubstrateExecutionClassStateDegraded, SupportedActions: []string{"runtime.apply_network_isolation", "runtime.restart_from_trusted_image", "hardening.block_exec_class_next_restart"}, SupportedDecisionModes: []string{RuntimeSubstrateDecisionModeImmediateContainment, RuntimeSubstrateDecisionModeNextRestartPreventive, RuntimeSubstrateDecisionModeTerminateAndRecover}, GuaranteeBoundaries: []string{"guest_scoped_containment_only"}, NonGuarantees: []string{"no_hypervisor_wide_claim"}},
		{ExecutionClass: RuntimeExecutionClassOfflineAirgappedNode, CurrentState: RuntimeSubstrateExecutionClassStateDegraded, SupportedActions: []string{"runtime.restart_from_trusted_image", "hardening.block_exec_class_next_restart"}, UnsupportedActions: []string{"runtime.apply_network_isolation"}, SupportedDecisionModes: []string{RuntimeSubstrateDecisionModeObserveOnly, RuntimeSubstrateDecisionModeNextRestartPreventive, RuntimeSubstrateDecisionModeTerminateAndRecover}, GuaranteeBoundaries: []string{"no_network_containment_when_signal_path_is_absent"}, NonGuarantees: []string{"no_fake_network_block_claim"}},
	}
	if got := EvaluateRuntimeSubstrateValDEnforcementAvailabilityState(enforcement); got != RuntimeSubstrateValDEnforcementAvailabilityStateActive {
		t.Fatalf("expected active enforcement availability state, got %q", got)
	}

	overhead := fullyMeasuredValDOverheadItems()
	if got := EvaluateRuntimeSubstrateValDOverheadVisibilityState(overhead); got != RuntimeSubstrateValDOverheadVisibilityStateActive {
		t.Fatalf("expected active overhead visibility state, got %q", got)
	}

	if got := EvaluateRuntimeSubstrateValDState(
		RuntimeSubstrateValCStateActive,
		RuntimeSubstrateValDExecutionClassMatrixStateActive,
		RuntimeSubstrateValDSignalCoverageStateActive,
		RuntimeSubstrateValDEnforcementAvailabilityStateActive,
		RuntimeSubstrateValDOverheadVisibilityStateActive,
	); got != RuntimeSubstrateValDStateActive {
		t.Fatalf("expected active val d state, got %q", got)
	}
}

func fullyMeasuredValDOverheadItems() []RuntimeSubstrateExecutionClassOverheadVisibilityItem {
	measuredAt := time.Date(2026, time.April, 23, 12, 0, 0, 0, time.UTC)
	return []RuntimeSubstrateExecutionClassOverheadVisibilityItem{
		{
			ExecutionClass:                   RuntimeExecutionClassStandardNode,
			CurrentState:                     RuntimeSubstrateExecutionClassStateSupported,
			MeasurementStatus:                "class_specific_measurement_verified",
			MeasurementBasis:                 "bounded_runtime_overhead_capture_window",
			MeasuredAt:                       measuredAt,
			MeasurementSource:                "runtime_substrate_vald_overhead_pack.standard_node.v1",
			EvidenceRefs:                     []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set/runtime_overhead/standard_node"},
			BudgetClass:                      "runtime_overhead_standard_node",
			BenchmarkRefs:                    []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set"},
			MethodologyRefs:                  []string{"/v1/runtime/boundaries", "/v1/public/benchmarks/methodology"},
			ObservedCPUOverheadMillicores:    14,
			ObservedMemoryOverheadMiB:        22,
			ObservedCaptureLatencyMicros:     340,
			ObservedCorrelationLatencyMicros: 580,
			VisibilityRules:                  []string{"class_specific_measured_overhead_must_remain_scoped_to_standard_node"},
		},
		{
			ExecutionClass:                   RuntimeExecutionClassHardenedNode,
			CurrentState:                     RuntimeSubstrateExecutionClassStateSupported,
			MeasurementStatus:                "class_specific_measurement_verified",
			MeasurementBasis:                 "bounded_runtime_overhead_capture_window",
			MeasuredAt:                       measuredAt.Add(10 * time.Minute),
			MeasurementSource:                "runtime_substrate_vald_overhead_pack.hardened_node.v1",
			EvidenceRefs:                     []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set/runtime_overhead/hardened_node"},
			BudgetClass:                      "runtime_overhead_hardened_node",
			BenchmarkRefs:                    []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set"},
			MethodologyRefs:                  []string{"/v1/runtime/boundaries", "/v1/public/benchmarks/methodology"},
			ObservedCPUOverheadMillicores:    18,
			ObservedMemoryOverheadMiB:        26,
			ObservedCaptureLatencyMicros:     390,
			ObservedCorrelationLatencyMicros: 610,
			VisibilityRules:                  []string{"class_specific_measured_overhead_must_remain_scoped_to_hardened_node"},
		},
		{
			ExecutionClass:                   RuntimeExecutionClassConfidentialCapableNode,
			CurrentState:                     RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementStatus:                "class_specific_measurement_recorded",
			MeasurementBasis:                 "guest_visible_runtime_overhead_capture_window",
			MeasuredAt:                       measuredAt.Add(20 * time.Minute),
			MeasurementSource:                "runtime_substrate_vald_overhead_pack.confidential_capable_node.v1",
			EvidenceRefs:                     []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set/runtime_overhead/confidential_capable_node"},
			BudgetClass:                      "runtime_overhead_confidential_boundary",
			BenchmarkRefs:                    []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set"},
			MethodologyRefs:                  []string{"/v1/runtime/boundaries", "/v1/public/benchmarks/methodology"},
			ObservedCPUOverheadMillicores:    24,
			ObservedMemoryOverheadMiB:        34,
			ObservedCaptureLatencyMicros:     470,
			ObservedCorrelationLatencyMicros: 720,
			VisibilityRules:                  []string{"class_specific_measured_overhead_must_remain_scoped_to_confidential_capable_node"},
		},
		{
			ExecutionClass:                   RuntimeExecutionClassVMBackedNode,
			CurrentState:                     RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementStatus:                "class_specific_measurement_recorded",
			MeasurementBasis:                 "guest_scoped_runtime_overhead_capture_window",
			MeasuredAt:                       measuredAt.Add(30 * time.Minute),
			MeasurementSource:                "runtime_substrate_vald_overhead_pack.vm_backed_node.v1",
			EvidenceRefs:                     []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set/runtime_overhead/vm_backed_node"},
			BudgetClass:                      "runtime_overhead_vm_guest_boundary",
			BenchmarkRefs:                    []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set"},
			MethodologyRefs:                  []string{"/v1/runtime/boundaries", "/v1/public/benchmarks/methodology"},
			ObservedCPUOverheadMillicores:    21,
			ObservedMemoryOverheadMiB:        30,
			ObservedCaptureLatencyMicros:     510,
			ObservedCorrelationLatencyMicros: 760,
			VisibilityRules:                  []string{"class_specific_measured_overhead_must_remain_scoped_to_vm_backed_node"},
		},
		{
			ExecutionClass:                   RuntimeExecutionClassOfflineAirgappedNode,
			CurrentState:                     RuntimeSubstrateExecutionClassStateDegraded,
			MeasurementStatus:                "class_specific_measurement_recorded",
			MeasurementBasis:                 "offline_runtime_overhead_capture_window",
			MeasuredAt:                       measuredAt.Add(40 * time.Minute),
			MeasurementSource:                "runtime_substrate_vald_overhead_pack.offline_airgapped_node.v1",
			EvidenceRefs:                     []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set/runtime_overhead/offline_airgapped_node"},
			BudgetClass:                      "runtime_overhead_offline_boundary",
			BenchmarkRefs:                    []string{"/v1/foundation/execution/benchmarks/harness", "/v1/public/benchmarks/set"},
			MethodologyRefs:                  []string{"/v1/runtime/boundaries", "/v1/public/benchmarks/methodology"},
			ObservedCPUOverheadMillicores:    16,
			ObservedMemoryOverheadMiB:        24,
			ObservedCaptureLatencyMicros:     360,
			ObservedCorrelationLatencyMicros: 640,
			VisibilityRules:                  []string{"class_specific_measured_overhead_must_remain_scoped_to_offline_airgapped_node"},
		},
	}
}

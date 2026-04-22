package runtime

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	RuntimeSubstrateEntryGateStateIncomplete         = "runtime_substrate_entry_gate_incomplete"
	RuntimeSubstrateEntryGateStateSubstantial        = "runtime_substrate_entry_gate_substantially_ready"
	RuntimeSubstrateEntryGateStateReady              = "runtime_substrate_entry_gate_ready"
	RuntimeSubstrateImplementationStateContractOnly  = "contract_only"
	RuntimeSubstrateImplementationStateRuntimePath   = "runtime_baseline_present"
	RuntimeSubstrateValAEventRecordSchema            = "runtime.substrate.vala.event_record.v1"
	RuntimeSubstrateValAEventSchemaStateActive       = "runtime_substrate_vala_event_schema_active"
	RuntimeSubstrateValAEventSchemaStatePartial      = "runtime_substrate_vala_event_schema_partial"
	RuntimeSubstrateValAEventSchemaStateIncomplete   = "runtime_substrate_vala_event_schema_incomplete"
	RuntimeSubstrateValASupportMatrixStateActive     = "runtime_substrate_vala_support_matrix_active"
	RuntimeSubstrateValASupportMatrixStatePartial    = "runtime_substrate_vala_support_matrix_partial"
	RuntimeSubstrateValASupportMatrixStateIncomplete = "runtime_substrate_vala_support_matrix_incomplete"
	RuntimeSubstrateValAObservabilityStateActive     = "runtime_substrate_vala_observability_active"
	RuntimeSubstrateValAObservabilityStatePartial    = "runtime_substrate_vala_observability_partial"
	RuntimeSubstrateValAObservabilityStateIncomplete = "runtime_substrate_vala_observability_incomplete"
	RuntimeSubstrateValAObservabilityContractDefined = "runtime_substrate_vala_observability_contract_defined"
	RuntimeSubstrateValAStateIncomplete              = "runtime_substrate_vala_incomplete"
	RuntimeSubstrateValAStateSubstantial             = "runtime_substrate_vala_substantially_ready"
	RuntimeSubstrateValAStateContractDefined         = "runtime_substrate_vala_contract_defined"
	RuntimeSubstrateValAStateActive                  = "runtime_substrate_vala_active"
	RuntimeSubstrateRecordKindExample                = "example_records"
	RuntimeSubstrateRecordKindObserved               = "observed_records"
	RuntimeSubstrateEventStateObserved               = "observed"
	RuntimeSubstrateEventStatePartiallyCorrelated    = "partially_correlated"
	RuntimeSubstrateEventStateStale                  = "stale"
	RuntimeSubstrateEventStateUnsupported            = "unsupported"
	RuntimeSubstrateFreshnessFresh                   = "fresh"
	RuntimeSubstrateFreshnessStale                   = "stale"
	RuntimeSubstrateFreshnessUnavailable             = "unavailable"
	RuntimeSubstrateConfidenceHighFidelity           = "high_fidelity_capture"
	RuntimeSubstrateConfidenceBoundedCorrelation     = "bounded_correlation"
	RuntimeSubstrateConfidenceLimitedContext         = "limited_context"
	RuntimeSubstrateConfidenceUnsupportedSignal      = "unsupported_signal"
	RuntimeSubstrateEventFamilyExecLifecycle         = "exec_lifecycle"
	RuntimeSubstrateEventFamilyProcessLineage        = "process_lineage"
	RuntimeSubstrateEventFamilyFileActivity          = "file_activity"
	RuntimeSubstrateEventFamilyNetworkActivity       = "network_activity"
	RuntimeExecutionClassStandardNode                = "standard_node"
	RuntimeExecutionClassHardenedNode                = "hardened_node"
	RuntimeExecutionClassConfidentialCapableNode     = "confidential_capable_node"
	RuntimeExecutionClassVMBackedNode                = "vm_backed_node"
	RuntimeExecutionClassOfflineAirgappedNode        = "offline_airgapped_node"
)

var ErrInvalidRuntimeSubstrateObservation = errors.New("invalid runtime substrate observation")

type RuntimeSubstrateEntryGate struct {
	CurrentState             string   `json:"current_state"`
	ImplementationPermission string   `json:"implementation_permission,omitempty"`
	RequiredBeforeValA       []string `json:"required_before_val_a,omitempty"`
	BoundaryRules            []string `json:"boundary_rules,omitempty"`
	ExplicitExclusions       []string `json:"explicit_exclusions,omitempty"`
	RequiredMeasurementGates []string `json:"required_measurement_gates,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type RuntimeSubstrateEventField struct {
	FieldName     string `json:"field_name"`
	Meaning       string `json:"meaning"`
	Required      bool   `json:"required"`
	MissingState  string `json:"missing_state"`
	FieldBoundary string `json:"field_boundary,omitempty"`
}

type RuntimeSubstrateEventSchema struct {
	CurrentState                   string                       `json:"current_state"`
	EventFamilies                  []string                     `json:"event_families,omitempty"`
	RequiredFields                 []RuntimeSubstrateEventField `json:"required_fields,omitempty"`
	CorrelationStates              []string                     `json:"correlation_states,omitempty"`
	FreshnessStates                []string                     `json:"freshness_states,omitempty"`
	AttributionConfidenceSemantics []string                     `json:"attribution_confidence_semantics,omitempty"`
	Limitations                    []string                     `json:"limitations,omitempty"`
}

type RuntimeSubstrateSignalCapability struct {
	SignalFamily          string   `json:"signal_family"`
	HookModel             string   `json:"hook_model"`
	CoverageState         string   `json:"coverage_state"`
	RequiredForValA       bool     `json:"required_for_val_a"`
	SupportBasis          string   `json:"support_basis,omitempty"`
	ExpectedHookCoverage  []string `json:"expected_hook_coverage,omitempty"`
	CapabilityAssumptions []string `json:"capability_assumptions,omitempty"`
	DegradedBehavior      string   `json:"degraded_behavior,omitempty"`
	UnsupportedReason     string   `json:"unsupported_reason,omitempty"`
}

type RuntimeSubstrateExecutionClassSupport struct {
	ExecutionClass string                             `json:"execution_class"`
	CurrentState   string                             `json:"current_state"`
	Capabilities   []RuntimeSubstrateSignalCapability `json:"capabilities,omitempty"`
	Limitations    []string                           `json:"limitations,omitempty"`
}

type RuntimeSubstrateObservedEvent struct {
	SchemaVersion         string           `json:"schema_version"`
	EventID               string           `json:"event_id"`
	EventFamily           string           `json:"event_family"`
	CurrentState          string           `json:"current_state"`
	Process               ProcessIdentity  `json:"process"`
	ParentPID             int              `json:"parent_pid,omitempty"`
	ThreadID              int              `json:"thread_id,omitempty"`
	ContainerRuntime      string           `json:"container_runtime,omitempty"`
	NamespaceMode         string           `json:"namespace_mode,omitempty"`
	Workload              WorkloadIdentity `json:"workload"`
	Node                  NodeIdentity     `json:"node"`
	AttributionConfidence string           `json:"attribution_confidence"`
	FreshnessState        string           `json:"freshness_state"`
	UnsupportedFields     []string         `json:"unsupported_fields,omitempty"`
	EvidenceRefs          []string         `json:"evidence_refs,omitempty"`
	Reasons               []string         `json:"reasons,omitempty"`
	ObservedAt            time.Time        `json:"observed_at"`
}

func RuntimeSubstrateEntryGateBaseline() RuntimeSubstrateEntryGate {
	return runtimeSubstrateEntryGate(RuntimeSubstrateImplementationStateContractOnly)
}

func RuntimeSubstrateEntryGateRuntimeBaseline() RuntimeSubstrateEntryGate {
	return runtimeSubstrateEntryGate(RuntimeSubstrateImplementationStateRuntimePath)
}

func runtimeSubstrateEntryGate(implementationPermission string) RuntimeSubstrateEntryGate {
	gate := RuntimeSubstrateEntryGate{
		ImplementationPermission: strings.TrimSpace(implementationPermission),
		RequiredBeforeValA: []string{
			"phase8_complete_as_bounded_formal_authority_phase",
			"canonical_evidence_spine_remains_single_truth_base",
			"runtime_boundaries_already_reject_absolute_truth_and_generic_memory_safety_claims",
			"degraded_and_unsupported_states_are_required_contract_outputs",
		},
		BoundaryRules: []string{
			"substrate observation, derived correlation, enforcement-capable signal, and unsupported-or-unknown remain explicitly separated",
			"val_a remains observability-only and does not claim binary provenance truth or inline prevention authority",
			"attribution confidence is capture and correlation quality only and not an enforcement risk score",
		},
		ExplicitExclusions: []string{
			"absolute_truth_rhetoric",
			"generic_memory_safety_claims",
			"universal_inline_prevention_claims",
			"unmeasured_latency_promises",
		},
		RequiredMeasurementGates: []string{
			"performance_budget_required_before_later_enforcement_claims",
			"false_positive_budget_required_before_later_policy_expansion",
		},
	}
	if gate.ImplementationPermission == RuntimeSubstrateImplementationStateRuntimePath {
		gate.Limitations = []string{
			"Val A runtime baseline is now present through bounded canonical observation ingest and correlation, but later provenance, enforcement, and benchmark waves remain deferred.",
			"Runtime baseline readiness does not widen claims into binary provenance truth, inline prevention, or measured enforcement latency.",
		}
	} else {
		gate.Limitations = []string{
			"Current entry-gate status is only substantially ready for the Val A contract slice and is not permission to claim that the full runtime implementation baseline already exists.",
			"Later provenance, enforcement, and benchmark waves remain deferred until a real runtime event path and measurement gate exist.",
		}
	}
	gate.CurrentState = EvaluateRuntimeSubstrateEntryGateState(gate)
	return gate
}

func EvaluateRuntimeSubstrateEntryGateState(gate RuntimeSubstrateEntryGate) string {
	if len(gate.RequiredBeforeValA) == 0 || len(gate.BoundaryRules) == 0 || len(gate.ExplicitExclusions) == 0 || len(gate.RequiredMeasurementGates) == 0 {
		return RuntimeSubstrateEntryGateStateIncomplete
	}
	if strings.TrimSpace(gate.ImplementationPermission) == RuntimeSubstrateImplementationStateRuntimePath {
		return RuntimeSubstrateEntryGateStateReady
	}
	return RuntimeSubstrateEntryGateStateSubstantial
}

func RuntimeSubstrateValAEventSchema() RuntimeSubstrateEventSchema {
	schema := RuntimeSubstrateEventSchema{
		EventFamilies: []string{
			RuntimeSubstrateEventFamilyExecLifecycle,
			RuntimeSubstrateEventFamilyProcessLineage,
			RuntimeSubstrateEventFamilyFileActivity,
			RuntimeSubstrateEventFamilyNetworkActivity,
		},
		RequiredFields: []RuntimeSubstrateEventField{
			{FieldName: "event_family", Meaning: "Stable event family classification for exec, process, file, or network signal.", Required: true, MissingState: RuntimeSubstrateEventStateUnsupported},
			{FieldName: "process.process_name", Meaning: "Observed process name for runtime attribution.", Required: true, MissingState: RuntimeSubstrateEventStatePartiallyCorrelated},
			{FieldName: "process.pid", Meaning: "Stable process identity anchor for bounded lineage and workload correlation.", Required: true, MissingState: RuntimeSubstrateEventStatePartiallyCorrelated},
			{FieldName: "process.lineage_ref", Meaning: "Parent or lineage context where supportable.", Required: true, MissingState: RuntimeSubstrateEventStatePartiallyCorrelated},
			{FieldName: "workload.namespace", Meaning: "Namespace or workload scope for the observed process or event.", Required: true, MissingState: RuntimeSubstrateEventStatePartiallyCorrelated},
			{FieldName: "workload.workload", Meaning: "Bounded workload identity tied to the captured event.", Required: true, MissingState: RuntimeSubstrateEventStatePartiallyCorrelated},
			{FieldName: "node.node_id", Meaning: "Node context for the observed event path.", Required: true, MissingState: RuntimeSubstrateEventStatePartiallyCorrelated},
			{FieldName: "observed_at", Meaning: "Capture timestamp for the substrate-backed event record.", Required: true, MissingState: RuntimeSubstrateEventStateUnsupported},
			{FieldName: "attribution_confidence", Meaning: "Capture and correlation quality class only; not an enforcement risk score.", Required: true, MissingState: RuntimeSubstrateEventStateUnsupported, FieldBoundary: "not_an_enforcement_score"},
			{FieldName: "freshness_state", Meaning: "Fresh, stale, or unavailable timing status for the event record.", Required: true, MissingState: RuntimeSubstrateEventStateUnsupported},
			{FieldName: "unsupported_fields", Meaning: "Explicit list of fields that are not available on the current path or execution class.", Required: true, MissingState: RuntimeSubstrateEventStateObserved, FieldBoundary: "missing_values_must_not_be_silent"},
		},
		CorrelationStates: []string{
			RuntimeSubstrateEventStateObserved,
			RuntimeSubstrateEventStatePartiallyCorrelated,
			RuntimeSubstrateEventStateStale,
			RuntimeSubstrateEventStateUnsupported,
		},
		FreshnessStates: []string{
			RuntimeSubstrateFreshnessFresh,
			RuntimeSubstrateFreshnessStale,
			RuntimeSubstrateFreshnessUnavailable,
		},
		AttributionConfidenceSemantics: []string{
			RuntimeSubstrateConfidenceHighFidelity + " = kernel-near capture with bounded process, workload, and node linkage present",
			RuntimeSubstrateConfidenceBoundedCorrelation + " = kernel-near or userspace-linked event with explicit correlation limits still visible",
			RuntimeSubstrateConfidenceLimitedContext + " = attribution remains bounded but some freshness or workload context is stale",
			RuntimeSubstrateConfidenceUnsupportedSignal + " = signal family or required field is unsupported on the current execution path and must not be inferred",
		},
		Limitations: []string{
			"Val A event schema is for substrate observability only and does not claim signed binary provenance, attestation linkage, or inline prevention semantics.",
			"Unsupported fields remain first-class output instead of being flattened into inferred truth.",
		},
	}
	schema.CurrentState = EvaluateRuntimeSubstrateValAEventSchemaState(schema)
	return schema
}

func RuntimeSubstrateValASupportMatrix() []RuntimeSubstrateExecutionClassSupport {
	return []RuntimeSubstrateExecutionClassSupport{
		{
			ExecutionClass: RuntimeExecutionClassStandardNode,
			CurrentState:   "runtime_substrate_observability_supported",
			Capabilities: []RuntimeSubstrateSignalCapability{
				{SignalFamily: RuntimeSubstrateEventFamilyExecLifecycle, HookModel: "tracepoint_exec", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"exec_enter", "exec_identity_fields"}, CapabilityAssumptions: []string{"linux_tracepoint_exec_available", "pid_capture_available"}},
				{SignalFamily: RuntimeSubstrateEventFamilyProcessLineage, HookModel: "tracepoint_sched_process", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"process_fork", "parent_child_lineage"}, CapabilityAssumptions: []string{"sched_process_events_available", "lineage_ref_builder_available"}},
				{SignalFamily: RuntimeSubstrateEventFamilyFileActivity, HookModel: "kprobe_file_open", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"open_path", "process_to_file_linkage"}, CapabilityAssumptions: []string{"kprobe_file_open_available", "file_path_capture_permitted"}},
				{SignalFamily: RuntimeSubstrateEventFamilyNetworkActivity, HookModel: "tracepoint_tcp_connect", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"socket_connect", "process_to_socket_linkage"}, CapabilityAssumptions: []string{"tcp_connect_signal_available", "socket_identity_capture_available"}},
			},
		},
		{
			ExecutionClass: RuntimeExecutionClassHardenedNode,
			CurrentState:   "runtime_substrate_observability_supported",
			Capabilities: []RuntimeSubstrateSignalCapability{
				{SignalFamily: RuntimeSubstrateEventFamilyExecLifecycle, HookModel: "tracepoint_exec_and_lsm_exec", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"exec_enter", "lsm_exec_boundary"}, CapabilityAssumptions: []string{"lsm_exec_hook_exposed", "kernel_tracepoint_exec_available"}},
				{SignalFamily: RuntimeSubstrateEventFamilyProcessLineage, HookModel: "tracepoint_sched_process", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"process_fork", "parent_child_lineage"}, CapabilityAssumptions: []string{"sched_process_events_available", "lineage_ref_builder_available"}},
				{SignalFamily: RuntimeSubstrateEventFamilyFileActivity, HookModel: "kprobe_file_open", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"open_path", "process_to_file_linkage"}, CapabilityAssumptions: []string{"kprobe_file_open_available", "filesystem_open_events_accessible"}},
				{SignalFamily: RuntimeSubstrateEventFamilyNetworkActivity, HookModel: "tracepoint_tcp_connect", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"socket_connect", "process_to_socket_linkage"}, CapabilityAssumptions: []string{"tcp_connect_signal_available", "socket_identity_capture_available"}},
			},
		},
		{
			ExecutionClass: RuntimeExecutionClassConfidentialCapableNode,
			CurrentState:   "runtime_substrate_observability_supported",
			Capabilities: []RuntimeSubstrateSignalCapability{
				{SignalFamily: RuntimeSubstrateEventFamilyExecLifecycle, HookModel: "tracepoint_exec", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"exec_enter", "process_identity_fields"}, CapabilityAssumptions: []string{"guest_exec_visibility_present", "node_confidential_capability_does_not_imply_provenance_truth"}},
				{SignalFamily: RuntimeSubstrateEventFamilyProcessLineage, HookModel: "tracepoint_sched_process", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"process_fork", "parent_child_lineage"}, CapabilityAssumptions: []string{"sched_process_events_available", "lineage_ref_builder_available"}},
				{SignalFamily: RuntimeSubstrateEventFamilyFileActivity, HookModel: "kprobe_file_open", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"open_path", "process_to_file_linkage"}, CapabilityAssumptions: []string{"file_open_signal_available", "guest_filesystem_context_only"}},
				{SignalFamily: RuntimeSubstrateEventFamilyNetworkActivity, HookModel: "tracepoint_tcp_connect", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"socket_connect", "process_to_socket_linkage"}, CapabilityAssumptions: []string{"connect_signal_available", "confidential_capable_node_claims_remain_capture_only"}},
			},
			Limitations: []string{
				"Confidential-capable observability remains capture-only in Val A and does not yet claim binary provenance or attestation-backed truth.",
			},
		},
		{
			ExecutionClass: RuntimeExecutionClassVMBackedNode,
			CurrentState:   "runtime_substrate_observability_degraded",
			Capabilities: []RuntimeSubstrateSignalCapability{
				{SignalFamily: RuntimeSubstrateEventFamilyExecLifecycle, HookModel: "guest_tracepoint_exec", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"guest_exec_enter", "guest_process_identity_fields"}, CapabilityAssumptions: []string{"guest_kernel_signal_path_available", "host_hypervisor_context_not_in_scope"}},
				{SignalFamily: RuntimeSubstrateEventFamilyProcessLineage, HookModel: "guest_tracepoint_sched_process", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"guest_process_fork", "guest_parent_child_lineage"}, CapabilityAssumptions: []string{"guest_sched_process_events_available", "host_lineage_not_in_scope"}},
				{SignalFamily: RuntimeSubstrateEventFamilyFileActivity, HookModel: "guest_kprobe_file_open", CoverageState: RuntimeSubstrateEventStatePartiallyCorrelated, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"guest_open_path", "guest_process_to_file_linkage"}, CapabilityAssumptions: []string{"guest_file_open_signal_available", "host_or_hypervisor_file_context_unavailable"}, DegradedBehavior: "file activity remains guest-scoped and must not imply host or hypervisor provenance"},
				{SignalFamily: RuntimeSubstrateEventFamilyNetworkActivity, HookModel: "guest_tracepoint_tcp_connect", CoverageState: RuntimeSubstrateEventStatePartiallyCorrelated, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"guest_socket_connect", "guest_process_to_socket_linkage"}, CapabilityAssumptions: []string{"guest_connect_signal_available", "cross_boundary_socket_context_unavailable"}, DegradedBehavior: "socket attribution remains guest-scoped when cross-boundary metadata is unavailable"},
			},
		},
		{
			ExecutionClass: RuntimeExecutionClassOfflineAirgappedNode,
			CurrentState:   "runtime_substrate_observability_degraded",
			Capabilities: []RuntimeSubstrateSignalCapability{
				{SignalFamily: RuntimeSubstrateEventFamilyExecLifecycle, HookModel: "tracepoint_exec", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"exec_enter", "process_identity_fields"}, CapabilityAssumptions: []string{"local_exec_signal_available", "offline_capture_pipeline_present"}},
				{SignalFamily: RuntimeSubstrateEventFamilyProcessLineage, HookModel: "tracepoint_sched_process", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"process_fork", "parent_child_lineage"}, CapabilityAssumptions: []string{"local_sched_process_events_available", "offline_lineage_builder_available"}},
				{SignalFamily: RuntimeSubstrateEventFamilyFileActivity, HookModel: "kprobe_file_open", CoverageState: RuntimeSubstrateEventStateObserved, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"open_path", "process_to_file_linkage"}, CapabilityAssumptions: []string{"local_file_open_signal_available", "filesystem_context_local_only"}},
				{SignalFamily: RuntimeSubstrateEventFamilyNetworkActivity, HookModel: "not_available_in_airgapped_capture_path", CoverageState: RuntimeSubstrateEventStateUnsupported, RequiredForValA: true, SupportBasis: "contract_assumption_only", ExpectedHookCoverage: []string{"network_signal_explicitly_absent"}, CapabilityAssumptions: []string{"offline_or_airgapped_capture_path"}, DegradedBehavior: "network activity must return unsupported instead of inferring absent egress", UnsupportedReason: "offline_or_airgapped_runtime_path"},
			},
			Limitations: []string{
				"Offline or air-gapped support stays explicit about missing network signal paths instead of translating isolation into fake negative evidence.",
			},
		},
	}
}

func RuntimeSubstrateValAObservedEvents() []RuntimeSubstrateObservedEvent {
	return []RuntimeSubstrateObservedEvent{
		{
			SchemaVersion:         RuntimeSubstrateValAEventRecordSchema,
			EventID:               "exec-api-20260422t090000z",
			EventFamily:           RuntimeSubstrateEventFamilyExecLifecycle,
			CurrentState:          RuntimeSubstrateEventStateObserved,
			Process:               ProcessIdentity{ProcessName: "api", ProcessPath: "/usr/local/bin/api", PID: 2048, CgroupID: "cg-api", NamespaceID: "ns-acme-prod", LineageRef: "init->containerd-shim->api"},
			ParentPID:             1024,
			ThreadID:              2048,
			ContainerRuntime:      "containerd",
			NamespaceMode:         "kubernetes_pod_namespace",
			Workload:              WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "api", PodUID: "pod-api-0", ImageDigest: "sha256:111"},
			Node:                  NodeIdentity{NodeID: "node-a", SubstrateClass: SubstrateClassStandard, TrustBoundary: TrustBoundaryKernelRuntimeLayer},
			AttributionConfidence: RuntimeSubstrateConfidenceHighFidelity,
			FreshnessState:        RuntimeSubstrateFreshnessFresh,
			EvidenceRefs:          []string{"audit://runtime/substrate/exec/api", "trace://exec/cluster-a/acme-prod/api"},
			Reasons:               []string{"exec_event_captured", "process_and_workload_identity_bound"},
			ObservedAt:            runtimeSubstrateSampleTime(0),
		},
		{
			SchemaVersion:         RuntimeSubstrateValAEventRecordSchema,
			EventID:               "process-lineage-worker-20260422t090100z",
			EventFamily:           RuntimeSubstrateEventFamilyProcessLineage,
			CurrentState:          RuntimeSubstrateEventStateObserved,
			Process:               ProcessIdentity{ProcessName: "worker", ProcessPath: "/usr/local/bin/worker", PID: 2210, CgroupID: "cg-worker", NamespaceID: "ns-acme-prod", LineageRef: "init->containerd-shim->worker"},
			ParentPID:             1024,
			ThreadID:              2211,
			ContainerRuntime:      "containerd",
			NamespaceMode:         "kubernetes_pod_namespace",
			Workload:              WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "worker", PodUID: "pod-worker-0", ImageDigest: "sha256:222"},
			Node:                  NodeIdentity{NodeID: "node-h1", SubstrateClass: SubstrateClassHardened, TrustBoundary: TrustBoundaryKernelRuntimeLayer},
			AttributionConfidence: RuntimeSubstrateConfidenceHighFidelity,
			FreshnessState:        RuntimeSubstrateFreshnessFresh,
			EvidenceRefs:          []string{"audit://runtime/substrate/process/worker", "trace://sched/cluster-a/acme-prod/worker"},
			Reasons:               []string{"lineage_capture_present", "parent_child_relationship_bound"},
			ObservedAt:            runtimeSubstrateSampleTime(1),
		},
		{
			SchemaVersion:         RuntimeSubstrateValAEventRecordSchema,
			EventID:               "file-open-batch-20260422t090200z",
			EventFamily:           RuntimeSubstrateEventFamilyFileActivity,
			CurrentState:          RuntimeSubstrateEventStatePartiallyCorrelated,
			Process:               ProcessIdentity{ProcessName: "batch", ProcessPath: "/usr/local/bin/batch", PID: 3010, CgroupID: "vm-cg-batch", NamespaceID: "vm-namespace", LineageRef: "systemd->batch"},
			ParentPID:             1,
			ContainerRuntime:      "vm_guest_agent",
			NamespaceMode:         "vm_guest_namespace",
			Workload:              WorkloadIdentity{ClusterID: "cluster-vm", Namespace: "finance", WorkloadKind: "VirtualMachine", Workload: "billing-batch", ImageDigest: "sha256:333"},
			Node:                  NodeIdentity{NodeID: "vm-node-a", SubstrateClass: SubstrateClassStandard, TrustBoundary: TrustBoundaryKernelRuntimeLayer},
			AttributionConfidence: RuntimeSubstrateConfidenceBoundedCorrelation,
			FreshnessState:        RuntimeSubstrateFreshnessFresh,
			UnsupportedFields:     []string{"thread_id", "host_hypervisor_ref"},
			EvidenceRefs:          []string{"audit://runtime/substrate/file/billing-batch", "trace://file-open/cluster-vm/finance/billing-batch"},
			Reasons:               []string{"guest_file_signal_captured", "host_level_hypervisor_context_not_in_scope"},
			ObservedAt:            runtimeSubstrateSampleTime(2),
		},
		{
			SchemaVersion:         RuntimeSubstrateValAEventRecordSchema,
			EventID:               "net-connect-gateway-20260422t090300z",
			EventFamily:           RuntimeSubstrateEventFamilyNetworkActivity,
			CurrentState:          RuntimeSubstrateEventStateObserved,
			Process:               ProcessIdentity{ProcessName: "gateway", ProcessPath: "/usr/local/bin/gateway", PID: 4120, CgroupID: "cg-gateway", NamespaceID: "ns-edge", LineageRef: "init->containerd-shim->gateway"},
			ParentPID:             1001,
			ThreadID:              4120,
			ContainerRuntime:      "containerd",
			NamespaceMode:         "kubernetes_pod_namespace",
			Workload:              WorkloadIdentity{ClusterID: "cluster-edge", Namespace: "edge", WorkloadKind: "Deployment", Workload: "gateway", PodUID: "pod-gateway-0", ImageDigest: "sha256:444"},
			Node:                  NodeIdentity{NodeID: "node-c1", SubstrateClass: SubstrateClassConfidential, TrustBoundary: TrustBoundaryKernelRuntimeLayer},
			AttributionConfidence: RuntimeSubstrateConfidenceBoundedCorrelation,
			FreshnessState:        RuntimeSubstrateFreshnessFresh,
			EvidenceRefs:          []string{"audit://runtime/substrate/network/gateway", "trace://tcp-connect/cluster-edge/edge/gateway"},
			Reasons:               []string{"socket_to_process_attribution_bound", "confidential_capable_node_claims_remain_capture_only"},
			ObservedAt:            runtimeSubstrateSampleTime(3),
		},
		{
			SchemaVersion:         RuntimeSubstrateValAEventRecordSchema,
			EventID:               "exec-reporter-20260422t053000z",
			EventFamily:           RuntimeSubstrateEventFamilyExecLifecycle,
			CurrentState:          RuntimeSubstrateEventStateStale,
			Process:               ProcessIdentity{ProcessName: "reporter", ProcessPath: "/usr/local/bin/reporter", PID: 1800, CgroupID: "cg-reporter", NamespaceID: "ns-archive", LineageRef: "init->containerd-shim->reporter"},
			ParentPID:             900,
			ThreadID:              1800,
			ContainerRuntime:      "containerd",
			NamespaceMode:         "kubernetes_pod_namespace",
			Workload:              WorkloadIdentity{ClusterID: "cluster-a", Namespace: "archive", WorkloadKind: "CronJob", Workload: "daily-report", PodUID: "pod-report-0", ImageDigest: "sha256:555"},
			Node:                  NodeIdentity{NodeID: "node-a", SubstrateClass: SubstrateClassStandard, TrustBoundary: TrustBoundaryKernelRuntimeLayer},
			AttributionConfidence: RuntimeSubstrateConfidenceLimitedContext,
			FreshnessState:        RuntimeSubstrateFreshnessStale,
			EvidenceRefs:          []string{"audit://runtime/substrate/exec/daily-report"},
			Reasons:               []string{"captured_exec_record_remains_visible", "freshness_window_expired"},
			ObservedAt:            runtimeSubstrateSampleTime(-210),
		},
		{
			SchemaVersion:         RuntimeSubstrateValAEventRecordSchema,
			EventID:               "net-airgap-sync-unsupported",
			EventFamily:           RuntimeSubstrateEventFamilyNetworkActivity,
			CurrentState:          RuntimeSubstrateEventStateUnsupported,
			Process:               ProcessIdentity{ProcessName: "sync-agent", ProcessPath: "/usr/local/bin/sync-agent", PID: 5100, CgroupID: "cg-sync", NamespaceID: "ns-offline", LineageRef: "systemd->sync-agent"},
			ParentPID:             1,
			ContainerRuntime:      "systemd_service",
			NamespaceMode:         "host_namespace",
			Workload:              WorkloadIdentity{ClusterID: "cluster-offline", Namespace: "offline", WorkloadKind: "DaemonSet", Workload: "sync-agent", PodUID: "pod-sync-0", ImageDigest: "sha256:666"},
			Node:                  NodeIdentity{NodeID: "node-offline-a", SubstrateClass: SubstrateClassHardened, TrustBoundary: TrustBoundaryKernelRuntimeLayer},
			AttributionConfidence: RuntimeSubstrateConfidenceUnsupportedSignal,
			FreshnessState:        RuntimeSubstrateFreshnessUnavailable,
			UnsupportedFields:     []string{"destination_ip", "socket_inode", "remote_cluster_context"},
			EvidenceRefs:          []string{"audit://runtime/substrate/network/offline-sync"},
			Reasons:               []string{"airgapped_network_capture_path_unavailable", "unsupported_is_explicit_not_negative_evidence"},
			ObservedAt:            runtimeSubstrateSampleTime(4),
		},
	}
}

func RuntimeSubstrateRemainingDeferredScope() []string {
	return []string{
		"binary_and_provenance_correlation",
		"enforcement_taxonomy_baseline",
		"execution_class_matrix_depth",
		"performance_and_proof_pack",
	}
}

func NormalizeRuntimeSubstrateObservedEvent(event RuntimeSubstrateObservedEvent, now func() time.Time) RuntimeSubstrateObservedEvent {
	if now == nil {
		now = time.Now
	}
	event.SchemaVersion = firstNonEmptyTrimmed(event.SchemaVersion, RuntimeSubstrateValAEventRecordSchema)
	event.EventFamily = normalizeRuntimeSubstrateEventFamily(event.EventFamily)
	event.CurrentState = normalizeRuntimeSubstrateEventState(event.CurrentState)
	event.AttributionConfidence = normalizeRuntimeSubstrateConfidence(event.AttributionConfidence)
	event.FreshnessState = normalizeRuntimeSubstrateFreshnessState(event.FreshnessState)
	event.ContainerRuntime = strings.TrimSpace(event.ContainerRuntime)
	event.NamespaceMode = strings.TrimSpace(event.NamespaceMode)
	event.Process.ProcessName = strings.TrimSpace(event.Process.ProcessName)
	event.Process.ProcessPath = strings.TrimSpace(event.Process.ProcessPath)
	event.Process.CgroupID = strings.TrimSpace(event.Process.CgroupID)
	event.Process.NamespaceID = strings.TrimSpace(event.Process.NamespaceID)
	event.Process.LineageRef = strings.TrimSpace(event.Process.LineageRef)
	event.Workload.ClusterID = strings.TrimSpace(event.Workload.ClusterID)
	event.Workload.Namespace = strings.TrimSpace(event.Workload.Namespace)
	event.Workload.WorkloadKind = strings.TrimSpace(event.Workload.WorkloadKind)
	event.Workload.Workload = strings.TrimSpace(event.Workload.Workload)
	event.Workload.PodUID = strings.TrimSpace(event.Workload.PodUID)
	event.Workload.ImageDigest = strings.TrimSpace(event.Workload.ImageDigest)
	event.Workload.PolicySubject = strings.TrimSpace(event.Workload.PolicySubject)
	event.Node.NodeID = strings.TrimSpace(event.Node.NodeID)
	event.Node.SubstrateClass = normalizeSubstrateClass(event.Node.SubstrateClass)
	event.Node.TrustBoundary = normalizeTrustBoundary(event.Node.TrustBoundary)
	event.Node.AttestationRef = strings.TrimSpace(event.Node.AttestationRef)
	event.UnsupportedFields = uniqueStrings(event.UnsupportedFields)
	event.EvidenceRefs = uniqueStrings(event.EvidenceRefs)
	event.Reasons = uniqueStrings(event.Reasons)
	if event.ObservedAt.IsZero() {
		event.ObservedAt = now().UTC()
	}
	if strings.TrimSpace(event.EventID) == "" {
		event.EventID = runtimeSubstrateObservedEventID(event)
	}
	return event
}

func ValidateRuntimeSubstrateObservedEvent(event RuntimeSubstrateObservedEvent) error {
	if strings.TrimSpace(event.SchemaVersion) != RuntimeSubstrateValAEventRecordSchema {
		return fmt.Errorf("%w: unsupported schema version", ErrInvalidRuntimeSubstrateObservation)
	}
	if strings.TrimSpace(event.EventFamily) == "" {
		return fmt.Errorf("%w: event_family is required", ErrInvalidRuntimeSubstrateObservation)
	}
	if strings.TrimSpace(event.CurrentState) == "" {
		return fmt.Errorf("%w: current_state is required", ErrInvalidRuntimeSubstrateObservation)
	}
	if strings.TrimSpace(event.AttributionConfidence) == "" {
		return fmt.Errorf("%w: attribution_confidence is required", ErrInvalidRuntimeSubstrateObservation)
	}
	if strings.TrimSpace(event.FreshnessState) == "" {
		return fmt.Errorf("%w: freshness_state is required", ErrInvalidRuntimeSubstrateObservation)
	}
	if event.ObservedAt.IsZero() {
		return fmt.Errorf("%w: observed_at is required", ErrInvalidRuntimeSubstrateObservation)
	}
	if event.CurrentState == RuntimeSubstrateEventStateUnsupported {
		if len(event.UnsupportedFields) == 0 {
			return fmt.Errorf("%w: unsupported observations must declare unsupported_fields", ErrInvalidRuntimeSubstrateObservation)
		}
		return nil
	}
	if event.Process.PID == 0 || strings.TrimSpace(event.Process.ProcessName) == "" {
		return fmt.Errorf("%w: stable process identity requires pid and process_name", ErrInvalidRuntimeSubstrateObservation)
	}
	if strings.TrimSpace(event.Workload.Namespace) == "" || strings.TrimSpace(event.Workload.WorkloadKind) == "" || strings.TrimSpace(event.Workload.Workload) == "" {
		return fmt.Errorf("%w: workload namespace, workload_kind, and workload are required", ErrInvalidRuntimeSubstrateObservation)
	}
	if strings.TrimSpace(event.Node.NodeID) == "" {
		return fmt.Errorf("%w: node.node_id is required", ErrInvalidRuntimeSubstrateObservation)
	}
	return nil
}

func EvaluateRuntimeSubstrateValAEventSchemaState(schema RuntimeSubstrateEventSchema) string {
	if len(schema.EventFamilies) == 0 || len(schema.RequiredFields) == 0 || len(schema.CorrelationStates) == 0 {
		return RuntimeSubstrateValAEventSchemaStateIncomplete
	}
	requiredFamilies := map[string]struct{}{
		RuntimeSubstrateEventFamilyExecLifecycle:   {},
		RuntimeSubstrateEventFamilyProcessLineage:  {},
		RuntimeSubstrateEventFamilyFileActivity:    {},
		RuntimeSubstrateEventFamilyNetworkActivity: {},
	}
	requiredStates := map[string]struct{}{
		RuntimeSubstrateEventStateObserved:            {},
		RuntimeSubstrateEventStatePartiallyCorrelated: {},
		RuntimeSubstrateEventStateStale:               {},
		RuntimeSubstrateEventStateUnsupported:         {},
	}
	partial := false
	for family := range requiredFamilies {
		if !containsString(schema.EventFamilies, family) {
			return RuntimeSubstrateValAEventSchemaStateIncomplete
		}
	}
	for state := range requiredStates {
		if !containsString(schema.CorrelationStates, state) {
			return RuntimeSubstrateValAEventSchemaStateIncomplete
		}
	}
	if !containsString(schema.FreshnessStates, RuntimeSubstrateFreshnessFresh) || !containsString(schema.FreshnessStates, RuntimeSubstrateFreshnessStale) {
		return RuntimeSubstrateValAEventSchemaStateIncomplete
	}
	if !containsString(schema.FreshnessStates, RuntimeSubstrateFreshnessUnavailable) {
		partial = true
	}
	if len(schema.RequiredFields) < 10 {
		partial = true
	}
	for _, field := range schema.RequiredFields {
		if strings.TrimSpace(field.FieldName) == "" || strings.TrimSpace(field.MissingState) == "" {
			partial = true
		}
	}
	if partial {
		return RuntimeSubstrateValAEventSchemaStatePartial
	}
	return RuntimeSubstrateValAEventSchemaStateActive
}

func EvaluateRuntimeSubstrateValASupportMatrixState(matrix []RuntimeSubstrateExecutionClassSupport) string {
	if len(matrix) == 0 {
		return RuntimeSubstrateValASupportMatrixStateIncomplete
	}
	expectedClasses := []string{
		RuntimeExecutionClassStandardNode,
		RuntimeExecutionClassHardenedNode,
		RuntimeExecutionClassConfidentialCapableNode,
		RuntimeExecutionClassVMBackedNode,
		RuntimeExecutionClassOfflineAirgappedNode,
	}
	expectedFamilies := []string{
		RuntimeSubstrateEventFamilyExecLifecycle,
		RuntimeSubstrateEventFamilyProcessLineage,
		RuntimeSubstrateEventFamilyFileActivity,
		RuntimeSubstrateEventFamilyNetworkActivity,
	}
	classes := map[string]RuntimeSubstrateExecutionClassSupport{}
	for _, class := range matrix {
		classes[strings.TrimSpace(class.ExecutionClass)] = class
	}
	partial := false
	for _, classID := range expectedClasses {
		class, ok := classes[classID]
		if !ok {
			return RuntimeSubstrateValASupportMatrixStateIncomplete
		}
		capabilities := map[string]RuntimeSubstrateSignalCapability{}
		for _, capability := range class.Capabilities {
			capabilities[strings.TrimSpace(capability.SignalFamily)] = capability
		}
		for _, family := range expectedFamilies {
			capability, ok := capabilities[family]
			if !ok {
				return RuntimeSubstrateValASupportMatrixStateIncomplete
			}
			if strings.TrimSpace(capability.CoverageState) == "" {
				partial = true
			}
			if capability.CoverageState != RuntimeSubstrateEventStateObserved && strings.TrimSpace(capability.DegradedBehavior) == "" {
				partial = true
			}
		}
	}
	if partial {
		return RuntimeSubstrateValASupportMatrixStatePartial
	}
	return RuntimeSubstrateValASupportMatrixStateActive
}

func EvaluateRuntimeSubstrateValAObservabilityState(events []RuntimeSubstrateObservedEvent) string {
	if len(events) == 0 {
		return RuntimeSubstrateValAObservabilityStateIncomplete
	}
	requiredFamilies := []string{
		RuntimeSubstrateEventFamilyExecLifecycle,
		RuntimeSubstrateEventFamilyProcessLineage,
		RuntimeSubstrateEventFamilyFileActivity,
		RuntimeSubstrateEventFamilyNetworkActivity,
	}
	requiredStates := []string{
		RuntimeSubstrateEventStateObserved,
		RuntimeSubstrateEventStatePartiallyCorrelated,
		RuntimeSubstrateEventStateStale,
		RuntimeSubstrateEventStateUnsupported,
	}
	partial := false
	for _, family := range requiredFamilies {
		if !hasRuntimeSubstrateEventFamily(events, family) {
			return RuntimeSubstrateValAObservabilityStateIncomplete
		}
	}
	for _, state := range requiredStates {
		if !hasRuntimeSubstrateEventState(events, state) {
			return RuntimeSubstrateValAObservabilityStateIncomplete
		}
	}
	for _, event := range events {
		if strings.TrimSpace(event.EventID) == "" || strings.TrimSpace(event.EventFamily) == "" || event.ObservedAt.IsZero() {
			return RuntimeSubstrateValAObservabilityStateIncomplete
		}
		if strings.TrimSpace(event.AttributionConfidence) == "" || strings.TrimSpace(event.FreshnessState) == "" {
			return RuntimeSubstrateValAObservabilityStateIncomplete
		}
		if event.CurrentState != RuntimeSubstrateEventStateUnsupported {
			if event.Process.PID == 0 && strings.TrimSpace(event.Process.ProcessName) == "" {
				partial = true
			}
			if strings.TrimSpace(event.Workload.ClusterID) == "" || strings.TrimSpace(event.Workload.Namespace) == "" {
				partial = true
			}
			if strings.TrimSpace(event.Node.NodeID) == "" {
				partial = true
			}
		}
	}
	if partial {
		return RuntimeSubstrateValAObservabilityStatePartial
	}
	return RuntimeSubstrateValAObservabilityStateActive
}

func EvaluateRuntimeSubstrateValAState(entryGateState, schemaState, supportMatrixState, observabilityState string) string {
	entryGateState = strings.TrimSpace(entryGateState)
	schemaState = strings.TrimSpace(schemaState)
	supportMatrixState = strings.TrimSpace(supportMatrixState)
	observabilityState = strings.TrimSpace(observabilityState)

	if entryGateState != RuntimeSubstrateEntryGateStateReady && entryGateState != RuntimeSubstrateEntryGateStateSubstantial {
		return RuntimeSubstrateValAStateIncomplete
	}
	states := []string{schemaState, supportMatrixState, observabilityState}
	hasPartial := false
	for _, state := range states {
		switch state {
		case RuntimeSubstrateValAEventSchemaStateActive, RuntimeSubstrateValASupportMatrixStateActive, RuntimeSubstrateValAObservabilityStateActive:
		case RuntimeSubstrateValAEventSchemaStatePartial, RuntimeSubstrateValASupportMatrixStatePartial, RuntimeSubstrateValAObservabilityStatePartial:
			hasPartial = true
		default:
			return RuntimeSubstrateValAStateIncomplete
		}
	}
	if entryGateState == RuntimeSubstrateEntryGateStateSubstantial {
		return RuntimeSubstrateValAStateContractDefined
	}
	if hasPartial {
		return RuntimeSubstrateValAStateSubstantial
	}
	return RuntimeSubstrateValAStateActive
}

func hasRuntimeSubstrateEventFamily(events []RuntimeSubstrateObservedEvent, family string) bool {
	for _, event := range events {
		if strings.TrimSpace(event.EventFamily) == family {
			return true
		}
	}
	return false
}

func hasRuntimeSubstrateEventState(events []RuntimeSubstrateObservedEvent, state string) bool {
	for _, event := range events {
		if strings.TrimSpace(event.CurrentState) == state {
			return true
		}
	}
	return false
}

func firstNonEmptyTrimmed(values ...string) string {
	for _, value := range values {
		if value = strings.TrimSpace(value); value != "" {
			return value
		}
	}
	return ""
}

func normalizeRuntimeSubstrateEventFamily(value string) string {
	switch strings.TrimSpace(value) {
	case RuntimeSubstrateEventFamilyExecLifecycle, RuntimeSubstrateEventFamilyProcessLineage, RuntimeSubstrateEventFamilyFileActivity, RuntimeSubstrateEventFamilyNetworkActivity:
		return strings.TrimSpace(value)
	default:
		return ""
	}
}

func normalizeRuntimeSubstrateEventState(value string) string {
	switch strings.TrimSpace(value) {
	case RuntimeSubstrateEventStateObserved, RuntimeSubstrateEventStatePartiallyCorrelated, RuntimeSubstrateEventStateStale, RuntimeSubstrateEventStateUnsupported:
		return strings.TrimSpace(value)
	default:
		return ""
	}
}

func normalizeRuntimeSubstrateFreshnessState(value string) string {
	switch strings.TrimSpace(value) {
	case RuntimeSubstrateFreshnessFresh, RuntimeSubstrateFreshnessStale, RuntimeSubstrateFreshnessUnavailable:
		return strings.TrimSpace(value)
	default:
		return ""
	}
}

func normalizeRuntimeSubstrateConfidence(value string) string {
	switch strings.TrimSpace(value) {
	case RuntimeSubstrateConfidenceHighFidelity, RuntimeSubstrateConfidenceBoundedCorrelation, RuntimeSubstrateConfidenceLimitedContext, RuntimeSubstrateConfidenceUnsupportedSignal:
		return strings.TrimSpace(value)
	default:
		return ""
	}
}

func runtimeSubstrateObservedEventID(event RuntimeSubstrateObservedEvent) string {
	return fmt.Sprintf(
		"%s-%s-%s-%d",
		firstNonEmptyTrimmed(event.EventFamily, "unknown-family"),
		firstNonEmptyTrimmed(event.Workload.Workload, "unknown-workload"),
		firstNonEmptyTrimmed(event.Process.ProcessName, "unknown-process"),
		event.ObservedAt.UTC().Unix(),
	)
}

func runtimeSubstrateSampleTime(offsetMinutes int) time.Time {
	base := time.Date(2026, time.April, 22, 9, 0, 0, 0, time.UTC)
	return base.Add(time.Duration(offsetMinutes) * time.Minute)
}

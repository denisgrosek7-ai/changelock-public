package runtime

import "testing"

func TestRuntimeSubstrateValCEnforcementTaxonomyIsActive(t *testing.T) {
	taxonomy := RuntimeSubstrateValCEnforcementTaxonomy()
	if taxonomy.CurrentState != RuntimeSubstrateEnforcementTaxonomyStateActive {
		t.Fatalf("expected active taxonomy, got %#v", taxonomy)
	}
	if !containsString(taxonomy.EnforcementClasses, RuntimeSubstrateEnforcementClassPrevent) {
		t.Fatalf("expected prevent class in taxonomy, got %#v", taxonomy.EnforcementClasses)
	}
	if !containsString(taxonomy.DecisionModes, RuntimeSubstrateDecisionModeTerminateAndRecover) {
		t.Fatalf("expected terminate-and-recover mode, got %#v", taxonomy.DecisionModes)
	}
}

func TestRuntimeSubstrateValCStateRequiresActiveValB(t *testing.T) {
	if got := EvaluateRuntimeSubstrateValCState(
		RuntimeSubstrateValBStateSubstantial,
		RuntimeSubstrateEnforcementTaxonomyStateActive,
		RuntimeSubstrateActionCatalogStateActive,
		RuntimeSubstratePolicyHookMappingStateActive,
		RuntimeSubstrateDecisionAuditStateActive,
	); got != RuntimeSubstrateValCStateIncomplete {
		t.Fatalf("expected incomplete val c without active val b, got %q", got)
	}
}

func TestRuntimeSubstrateValCActionHookAndAuditStates(t *testing.T) {
	actions := []RuntimeSubstrateEnforcementActionCatalogItem{
		{
			ActionID:                  "runtime.capture_forensics",
			SourceKind:                "runtime_enforcement",
			GuaranteeClass:            RuntimeSubstrateEnforcementClassObserve,
			DecisionMode:              RuntimeSubstrateDecisionModeObserveOnly,
			PolicyRef:                 "runtime_assurance_policy.v1",
			HookMappingRefs:           []string{"hook.observe"},
			Guarantees:                []string{"records bounded evidence capture"},
			NonGuarantees:             []string{"does not mutate workload state"},
			SupportedExecutionClasses: []string{"standard_node"},
			AuditTrailExpectations:    []string{"runtime_enforcement_evaluated"},
		},
		{
			ActionID:                  "runtime.network_isolation",
			SourceKind:                "runtime_enforcement",
			GuaranteeClass:            RuntimeSubstrateEnforcementClassContain,
			DecisionMode:              RuntimeSubstrateDecisionModeImmediateContainment,
			PolicyRef:                 "runtime_assurance_policy.v1",
			HookMappingRefs:           []string{"hook.contain"},
			ApprovalRequired:          true,
			RollbackRequired:          true,
			Guarantees:                []string{"reduces blast radius after detection"},
			NonGuarantees:             []string{"does not claim universal prevention"},
			SupportedExecutionClasses: []string{"standard_node"},
			AuditTrailExpectations:    []string{"runtime_network_isolation_applied"},
		},
		{
			ActionID:                  "hardening.next_restart_exec_block",
			SourceKind:                "hardening_execution",
			GuaranteeClass:            RuntimeSubstrateEnforcementClassPrevent,
			DecisionMode:              RuntimeSubstrateDecisionModeNextRestartPreventive,
			PolicyRef:                 "runtime_closed_loop_hardening.v1",
			HookMappingRefs:           []string{"hook.prevent"},
			RollbackRequired:          true,
			Guarantees:                []string{"stages deny semantics for later trusted restart"},
			NonGuarantees:             []string{"does not block the current process immediately"},
			SupportedExecutionClasses: []string{"standard_node"},
			AuditTrailExpectations:    []string{"hardening_action_applied"},
		},
		{
			ActionID:                  "runtime.trusted_restart",
			SourceKind:                "runtime_enforcement",
			GuaranteeClass:            RuntimeSubstrateEnforcementClassTerminate,
			DecisionMode:              RuntimeSubstrateDecisionModeTerminateAndRecover,
			PolicyRef:                 "runtime_assurance_policy.v1",
			HookMappingRefs:           []string{"hook.terminate"},
			ApprovalRequired:          true,
			RollbackRequired:          true,
			Guarantees:                []string{"requests trusted restart after detection and review"},
			NonGuarantees:             []string{"does not imply pre-execution omniscience"},
			SupportedExecutionClasses: []string{"standard_node"},
			AuditTrailExpectations:    []string{"runtime_trusted_restart_requested"},
		},
	}
	if got := EvaluateRuntimeSubstrateValCActionCatalogState(actions); got != RuntimeSubstrateActionCatalogStateActive {
		t.Fatalf("expected active action catalog, got %q", got)
	}

	hooks := []RuntimeSubstratePolicyHookMapping{
		{
			MappingID:                 "hook.observe",
			PolicyRef:                 "runtime_assurance_policy.v1",
			ActionID:                  "runtime.capture_forensics",
			HookModel:                 "audit_backed_forensic_request",
			GuaranteeClass:            RuntimeSubstrateEnforcementClassObserve,
			DecisionMode:              RuntimeSubstrateDecisionModeObserveOnly,
			GuaranteeSemantics:        []string{"request is audit-trailed"},
			NonGuarantees:             []string{"not a containment or kill action"},
			AuditTrailSources:         []string{"/v1/runtime/enforcement"},
			SupportedExecutionClasses: []string{"standard_node"},
		},
		{
			MappingID:                 "hook.contain",
			PolicyRef:                 "runtime_assurance_policy.v1",
			ActionID:                  "runtime.network_isolation",
			HookModel:                 "network_isolation_apply",
			GuaranteeClass:            RuntimeSubstrateEnforcementClassContain,
			DecisionMode:              RuntimeSubstrateDecisionModeImmediateContainment,
			GuaranteeSemantics:        []string{"isolation applies after approval and execution"},
			NonGuarantees:             []string{"no universal prevention claim"},
			AuditTrailSources:         []string{"/v1/runtime/enforcement"},
			SupportedExecutionClasses: []string{"standard_node"},
		},
		{
			MappingID:                 "hook.prevent",
			PolicyRef:                 "runtime_closed_loop_hardening.v1",
			ActionID:                  "hardening.next_restart_exec_block",
			HookModel:                 "next_restart_exec_class_block",
			GuaranteeClass:            RuntimeSubstrateEnforcementClassPrevent,
			DecisionMode:              RuntimeSubstrateDecisionModeNextRestartPreventive,
			GuaranteeSemantics:        []string{"deny semantics become active on later trusted restart"},
			NonGuarantees:             []string{"no immediate in-place block claim"},
			AuditTrailSources:         []string{"/v1/hardening/actions"},
			SupportedExecutionClasses: []string{"standard_node"},
		},
		{
			MappingID:                 "hook.terminate",
			PolicyRef:                 "runtime_assurance_policy.v1",
			ActionID:                  "runtime.trusted_restart",
			HookModel:                 "trusted_restart_request",
			GuaranteeClass:            RuntimeSubstrateEnforcementClassTerminate,
			DecisionMode:              RuntimeSubstrateDecisionModeTerminateAndRecover,
			GuaranteeSemantics:        []string{"restart request is explicit and audit-trailed"},
			NonGuarantees:             []string{"does not promise universal workload recovery"},
			AuditTrailSources:         []string{"/v1/runtime/enforcement"},
			SupportedExecutionClasses: []string{"standard_node"},
		},
	}
	if got := EvaluateRuntimeSubstrateValCPolicyHookMappingState(hooks); got != RuntimeSubstratePolicyHookMappingStateActive {
		t.Fatalf("expected active hook mapping state, got %q", got)
	}

	audit := []RuntimeSubstrateDecisionAuditRecord{
		{
			SourceKind:       "runtime_enforcement",
			SubjectRef:       "cluster-a|ns|Deployment|api",
			DecisionRef:      "decision-observe",
			ActionID:         "runtime.capture_forensics",
			GuaranteeClass:   RuntimeSubstrateEnforcementClassObserve,
			DecisionMode:     RuntimeSubstrateDecisionModeObserveOnly,
			ExecutionResult:  "forensic_snapshot_requested",
			AuditEventType:   "runtime_forensic_snapshot_requested",
			AuditTrailRefs:   []string{"/v1/runtime/enforcement"},
			EvidenceRefs:     []string{"audit://evidence/1"},
			Guarantees:       []string{"bounded evidence capture"},
			NonGuarantees:    []string{"not a kill action"},
			Executed:         true,
			ApprovalRequired: false,
			RollbackRequired: false,
		},
		{
			SourceKind:       "runtime_enforcement",
			SubjectRef:       "cluster-a|ns|Deployment|api",
			DecisionRef:      "decision-contain",
			ActionID:         "runtime.network_isolation",
			GuaranteeClass:   RuntimeSubstrateEnforcementClassContain,
			DecisionMode:     RuntimeSubstrateDecisionModeImmediateContainment,
			ExecutionResult:  "network_isolation_applied",
			AuditEventType:   "runtime_network_isolation_applied",
			AuditTrailRefs:   []string{"/v1/runtime/enforcement"},
			EvidenceRefs:     []string{"audit://evidence/2"},
			Guarantees:       []string{"blast radius reduced"},
			NonGuarantees:    []string{"no pre-execution block claim"},
			Executed:         true,
			ApprovalRequired: true,
			ApprovalState:    "approved_and_executed",
			RollbackRequired: true,
			RollbackState:    "rollback_ready",
		},
	}
	if got := EvaluateRuntimeSubstrateValCDecisionAuditState(audit); got != RuntimeSubstrateDecisionAuditStateActive {
		t.Fatalf("expected active decision audit state, got %q", got)
	}
}

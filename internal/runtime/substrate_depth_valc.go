package runtime

import "strings"

const (
	RuntimeSubstrateEnforcementClassObserve     = "observe"
	RuntimeSubstrateEnforcementClassPrevent     = "prevent"
	RuntimeSubstrateEnforcementClassContain     = "contain"
	RuntimeSubstrateEnforcementClassTerminate   = "terminate"
	RuntimeSubstrateEnforcementClassUnsupported = "unsupported"

	RuntimeSubstrateDecisionModeObserveOnly           = "observe_only"
	RuntimeSubstrateDecisionModeSampleOrEscalate      = "sample_or_escalate"
	RuntimeSubstrateDecisionModeImmediateContainment  = "immediate_containment"
	RuntimeSubstrateDecisionModeNextRestartPreventive = "next_restart_preventive"
	RuntimeSubstrateDecisionModeTerminateAndRecover   = "terminate_and_recover"
	RuntimeSubstrateDecisionModeUnsupported           = "unsupported"

	RuntimeSubstrateActionCatalogStateActive     = "runtime_substrate_valc_action_catalog_active"
	RuntimeSubstrateActionCatalogStatePartial    = "runtime_substrate_valc_action_catalog_partial"
	RuntimeSubstrateActionCatalogStateIncomplete = "runtime_substrate_valc_action_catalog_incomplete"

	RuntimeSubstratePolicyHookMappingStateActive     = "runtime_substrate_valc_policy_hook_mapping_active"
	RuntimeSubstratePolicyHookMappingStatePartial    = "runtime_substrate_valc_policy_hook_mapping_partial"
	RuntimeSubstratePolicyHookMappingStateIncomplete = "runtime_substrate_valc_policy_hook_mapping_incomplete"

	RuntimeSubstrateDecisionAuditStateActive     = "runtime_substrate_valc_decision_audit_active"
	RuntimeSubstrateDecisionAuditStatePartial    = "runtime_substrate_valc_decision_audit_partial"
	RuntimeSubstrateDecisionAuditStateIncomplete = "runtime_substrate_valc_decision_audit_incomplete"

	RuntimeSubstrateEnforcementTaxonomyStateActive     = "runtime_substrate_valc_taxonomy_active"
	RuntimeSubstrateEnforcementTaxonomyStatePartial    = "runtime_substrate_valc_taxonomy_partial"
	RuntimeSubstrateEnforcementTaxonomyStateIncomplete = "runtime_substrate_valc_taxonomy_incomplete"

	RuntimeSubstrateValCStateIncomplete  = "runtime_substrate_valc_incomplete"
	RuntimeSubstrateValCStateSubstantial = "runtime_substrate_valc_substantially_ready"
	RuntimeSubstrateValCStateActive      = "runtime_substrate_valc_active"
)

type RuntimeSubstrateEnforcementTaxonomy struct {
	CurrentState           string   `json:"current_state"`
	EnforcementClasses     []string `json:"enforcement_classes,omitempty"`
	DecisionModes          []string `json:"decision_modes,omitempty"`
	RequiredAuditInputs    []string `json:"required_audit_inputs,omitempty"`
	GuaranteePrinciples    []string `json:"guarantee_principles,omitempty"`
	NonGuaranteePrinciples []string `json:"non_guarantee_principles,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type RuntimeSubstrateEnforcementActionCatalogItem struct {
	ActionID                  string   `json:"action_id"`
	SourceKind                string   `json:"source_kind"`
	GuaranteeClass            string   `json:"guarantee_class"`
	DecisionMode              string   `json:"decision_mode"`
	PolicyRef                 string   `json:"policy_ref"`
	HookMappingRefs           []string `json:"hook_mapping_refs,omitempty"`
	ApprovalRequired          bool     `json:"approval_required"`
	RollbackRequired          bool     `json:"rollback_required"`
	Guarantees                []string `json:"guarantees,omitempty"`
	NonGuarantees             []string `json:"non_guarantees,omitempty"`
	SupportedExecutionClasses []string `json:"supported_execution_classes,omitempty"`
	UnsupportedClasses        []string `json:"unsupported_classes,omitempty"`
	AuditTrailExpectations    []string `json:"audit_trail_expectations,omitempty"`
}

type RuntimeSubstratePolicyHookMapping struct {
	MappingID                 string   `json:"mapping_id"`
	PolicyRef                 string   `json:"policy_ref"`
	ActionID                  string   `json:"action_id"`
	HookModel                 string   `json:"hook_model"`
	GuaranteeClass            string   `json:"guarantee_class"`
	DecisionMode              string   `json:"decision_mode"`
	GuaranteeSemantics        []string `json:"guarantee_semantics,omitempty"`
	NonGuarantees             []string `json:"non_guarantees,omitempty"`
	AuditTrailSources         []string `json:"audit_trail_sources,omitempty"`
	SupportedExecutionClasses []string `json:"supported_execution_classes,omitempty"`
	UnsupportedClasses        []string `json:"unsupported_classes,omitempty"`
}

type RuntimeSubstrateDecisionAuditRecord struct {
	SourceKind       string   `json:"source_kind"`
	SubjectRef       string   `json:"subject_ref"`
	DecisionRef      string   `json:"decision_ref"`
	ActionID         string   `json:"action_id"`
	GuaranteeClass   string   `json:"guarantee_class"`
	DecisionMode     string   `json:"decision_mode"`
	ApprovalRequired bool     `json:"approval_required"`
	ApprovalState    string   `json:"approval_state,omitempty"`
	RollbackRequired bool     `json:"rollback_required"`
	RollbackState    string   `json:"rollback_state,omitempty"`
	Executed         bool     `json:"executed"`
	ExecutionResult  string   `json:"execution_result"`
	AuditEventType   string   `json:"audit_event_type"`
	AuditTrailRefs   []string `json:"audit_trail_refs,omitempty"`
	EvidenceRefs     []string `json:"evidence_refs,omitempty"`
	Guarantees       []string `json:"guarantees,omitempty"`
	NonGuarantees    []string `json:"non_guarantees,omitempty"`
}

func RuntimeSubstrateValCEnforcementTaxonomy() RuntimeSubstrateEnforcementTaxonomy {
	taxonomy := RuntimeSubstrateEnforcementTaxonomy{
		EnforcementClasses: []string{
			RuntimeSubstrateEnforcementClassObserve,
			RuntimeSubstrateEnforcementClassPrevent,
			RuntimeSubstrateEnforcementClassContain,
			RuntimeSubstrateEnforcementClassTerminate,
			RuntimeSubstrateEnforcementClassUnsupported,
		},
		DecisionModes: []string{
			RuntimeSubstrateDecisionModeObserveOnly,
			RuntimeSubstrateDecisionModeSampleOrEscalate,
			RuntimeSubstrateDecisionModeImmediateContainment,
			RuntimeSubstrateDecisionModeNextRestartPreventive,
			RuntimeSubstrateDecisionModeTerminateAndRecover,
			RuntimeSubstrateDecisionModeUnsupported,
		},
		RequiredAuditInputs: []string{
			"decision.policy_ref",
			"decision.approval_required",
			"decision.rollback_required",
			"decision.execution_result",
			"canonical_runtime_or_hardening_audit_event",
			"canonical_evidence_refs",
		},
		GuaranteePrinciples: []string{
			"observe means bounded evaluation or evidence-capture semantics only and does not mutate workload state by itself",
			"contain means blast-radius reduction after detection and must not be described as universal prevention",
			"prevent is only claimed for bounded next-restart or hook-scoped deny paths whose timing is explicitly declared",
			"terminate means detection-triggered recovery or shutdown semantics after review or approval, not pre-execution omniscience",
		},
		NonGuaranteePrinciples: []string{
			"unsupported or degraded enforcement paths remain explicit instead of being flattened into prevent or contain claims",
			"approval-gated decisions do not imply execution until a corresponding audit-trailed execution record exists",
			"rollback discipline describes reversibility and cleanup semantics, not guaranteed instant undo on every substrate class",
		},
		Limitations: []string{
			"Val C documents bounded enforcement taxonomy over canonical runtime and hardening trails; it does not claim universal inline blocking or kernel-level monopoly.",
			"Selected prevent paths remain scoped to declared next-restart or hook-limited semantics and do not expand Val C into generic memory-safety or broad syscall-override claims.",
		},
	}
	taxonomy.CurrentState = EvaluateRuntimeSubstrateValCEnforcementTaxonomyState(taxonomy)
	return taxonomy
}

func RuntimeSubstrateValCRemainingDeferredScope() []string {
	return []string{
		"execution_class_matrix_depth",
		"performance_and_proof_pack",
	}
}

func EvaluateRuntimeSubstrateValCEnforcementTaxonomyState(taxonomy RuntimeSubstrateEnforcementTaxonomy) string {
	if len(taxonomy.EnforcementClasses) == 0 || len(taxonomy.DecisionModes) == 0 || len(taxonomy.RequiredAuditInputs) == 0 {
		return RuntimeSubstrateEnforcementTaxonomyStateIncomplete
	}
	for _, class := range []string{
		RuntimeSubstrateEnforcementClassObserve,
		RuntimeSubstrateEnforcementClassPrevent,
		RuntimeSubstrateEnforcementClassContain,
		RuntimeSubstrateEnforcementClassTerminate,
		RuntimeSubstrateEnforcementClassUnsupported,
	} {
		if !containsString(taxonomy.EnforcementClasses, class) {
			return RuntimeSubstrateEnforcementTaxonomyStateIncomplete
		}
	}
	for _, mode := range []string{
		RuntimeSubstrateDecisionModeObserveOnly,
		RuntimeSubstrateDecisionModeSampleOrEscalate,
		RuntimeSubstrateDecisionModeImmediateContainment,
		RuntimeSubstrateDecisionModeNextRestartPreventive,
		RuntimeSubstrateDecisionModeTerminateAndRecover,
		RuntimeSubstrateDecisionModeUnsupported,
	} {
		if !containsString(taxonomy.DecisionModes, mode) {
			return RuntimeSubstrateEnforcementTaxonomyStateIncomplete
		}
	}
	if len(taxonomy.GuaranteePrinciples) < 3 || len(taxonomy.NonGuaranteePrinciples) < 2 {
		return RuntimeSubstrateEnforcementTaxonomyStatePartial
	}
	return RuntimeSubstrateEnforcementTaxonomyStateActive
}

func EvaluateRuntimeSubstrateValCActionCatalogState(items []RuntimeSubstrateEnforcementActionCatalogItem) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.ActionID) == "" || strings.TrimSpace(item.SourceKind) == "" || strings.TrimSpace(item.PolicyRef) == "" {
					return false
				}
				if !isRuntimeSubstrateEnforcementClass(item.GuaranteeClass) || !isRuntimeSubstrateDecisionMode(item.DecisionMode) {
					return false
				}
				if len(item.HookMappingRefs) == 0 || len(item.Guarantees) == 0 || len(item.NonGuarantees) == 0 || len(item.AuditTrailExpectations) == 0 {
					return false
				}
			}
			return true
		},
		func() bool {
			return runtimeSubstrateValCActionCatalogCovers(items, RuntimeSubstrateEnforcementClassObserve) &&
				runtimeSubstrateValCActionCatalogCovers(items, RuntimeSubstrateEnforcementClassContain) &&
				runtimeSubstrateValCActionCatalogCovers(items, RuntimeSubstrateEnforcementClassPrevent) &&
				runtimeSubstrateValCActionCatalogCovers(items, RuntimeSubstrateEnforcementClassTerminate) &&
				runtimeSubstrateValCActionCatalogHasApprovalAndRollback(items)
		},
		RuntimeSubstrateActionCatalogStateIncomplete,
		RuntimeSubstrateActionCatalogStatePartial,
		RuntimeSubstrateActionCatalogStateActive,
	)
}

func EvaluateRuntimeSubstrateValCPolicyHookMappingState(items []RuntimeSubstratePolicyHookMapping) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.MappingID) == "" || strings.TrimSpace(item.PolicyRef) == "" || strings.TrimSpace(item.ActionID) == "" || strings.TrimSpace(item.HookModel) == "" {
					return false
				}
				if !isRuntimeSubstrateEnforcementClass(item.GuaranteeClass) || !isRuntimeSubstrateDecisionMode(item.DecisionMode) {
					return false
				}
				if len(item.GuaranteeSemantics) == 0 || len(item.NonGuarantees) == 0 || len(item.AuditTrailSources) == 0 {
					return false
				}
			}
			return true
		},
		func() bool {
			return runtimeSubstrateValCHookCoverage(items, RuntimeSubstrateEnforcementClassContain) &&
				runtimeSubstrateValCHookCoverage(items, RuntimeSubstrateEnforcementClassPrevent) &&
				runtimeSubstrateValCHookCoverage(items, RuntimeSubstrateEnforcementClassTerminate)
		},
		RuntimeSubstratePolicyHookMappingStateIncomplete,
		RuntimeSubstratePolicyHookMappingStatePartial,
		RuntimeSubstratePolicyHookMappingStateActive,
	)
}

func EvaluateRuntimeSubstrateValCDecisionAuditState(items []RuntimeSubstrateDecisionAuditRecord) string {
	return evaluateRuntimeSubstrateCorrelationItems(
		len(items),
		func() bool {
			for _, item := range items {
				if strings.TrimSpace(item.SourceKind) == "" || strings.TrimSpace(item.SubjectRef) == "" || strings.TrimSpace(item.DecisionRef) == "" || strings.TrimSpace(item.ActionID) == "" {
					return false
				}
				if !isRuntimeSubstrateEnforcementClass(item.GuaranteeClass) || !isRuntimeSubstrateDecisionMode(item.DecisionMode) {
					return false
				}
				if strings.TrimSpace(item.ExecutionResult) == "" || strings.TrimSpace(item.AuditEventType) == "" {
					return false
				}
				if len(item.AuditTrailRefs) == 0 || len(item.Guarantees) == 0 || len(item.NonGuarantees) == 0 {
					return false
				}
				if item.RollbackRequired && strings.TrimSpace(item.RollbackState) == "" {
					return false
				}
				if item.ApprovalRequired && strings.TrimSpace(item.ApprovalState) == "" {
					return false
				}
			}
			return true
		},
		func() bool {
			hasExecuted := false
			hasApprovalOrRollback := false
			for _, item := range items {
				if item.Executed {
					hasExecuted = true
				}
				if item.ApprovalRequired || item.RollbackRequired {
					hasApprovalOrRollback = true
				}
			}
			return hasExecuted && hasApprovalOrRollback
		},
		RuntimeSubstrateDecisionAuditStateIncomplete,
		RuntimeSubstrateDecisionAuditStatePartial,
		RuntimeSubstrateDecisionAuditStateActive,
	)
}

func EvaluateRuntimeSubstrateValCState(valBState, taxonomyState, actionCatalogState, hookMappingState, decisionAuditState string) string {
	if strings.TrimSpace(valBState) != RuntimeSubstrateValBStateActive {
		return RuntimeSubstrateValCStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(taxonomyState),
		strings.TrimSpace(actionCatalogState),
		strings.TrimSpace(hookMappingState),
		strings.TrimSpace(decisionAuditState),
	} {
		switch state {
		case RuntimeSubstrateEnforcementTaxonomyStateActive,
			RuntimeSubstrateActionCatalogStateActive,
			RuntimeSubstratePolicyHookMappingStateActive,
			RuntimeSubstrateDecisionAuditStateActive:
		case RuntimeSubstrateEnforcementTaxonomyStatePartial,
			RuntimeSubstrateActionCatalogStatePartial,
			RuntimeSubstratePolicyHookMappingStatePartial,
			RuntimeSubstrateDecisionAuditStatePartial:
			hasPartial = true
		default:
			return RuntimeSubstrateValCStateIncomplete
		}
	}
	if hasPartial {
		return RuntimeSubstrateValCStateSubstantial
	}
	return RuntimeSubstrateValCStateActive
}

func isRuntimeSubstrateEnforcementClass(value string) bool {
	switch strings.TrimSpace(value) {
	case RuntimeSubstrateEnforcementClassObserve,
		RuntimeSubstrateEnforcementClassPrevent,
		RuntimeSubstrateEnforcementClassContain,
		RuntimeSubstrateEnforcementClassTerminate,
		RuntimeSubstrateEnforcementClassUnsupported:
		return true
	default:
		return false
	}
}

func isRuntimeSubstrateDecisionMode(value string) bool {
	switch strings.TrimSpace(value) {
	case RuntimeSubstrateDecisionModeObserveOnly,
		RuntimeSubstrateDecisionModeSampleOrEscalate,
		RuntimeSubstrateDecisionModeImmediateContainment,
		RuntimeSubstrateDecisionModeNextRestartPreventive,
		RuntimeSubstrateDecisionModeTerminateAndRecover,
		RuntimeSubstrateDecisionModeUnsupported:
		return true
	default:
		return false
	}
}

func runtimeSubstrateValCActionCatalogCovers(items []RuntimeSubstrateEnforcementActionCatalogItem, class string) bool {
	for _, item := range items {
		if item.GuaranteeClass == class {
			return true
		}
	}
	return false
}

func runtimeSubstrateValCActionCatalogHasApprovalAndRollback(items []RuntimeSubstrateEnforcementActionCatalogItem) bool {
	hasApproval := false
	hasRollback := false
	for _, item := range items {
		if item.ApprovalRequired {
			hasApproval = true
		}
		if item.RollbackRequired {
			hasRollback = true
		}
	}
	return hasApproval && hasRollback
}

func runtimeSubstrateValCHookCoverage(items []RuntimeSubstratePolicyHookMapping, class string) bool {
	for _, item := range items {
		if item.GuaranteeClass == class {
			return true
		}
	}
	return false
}

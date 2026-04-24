package workflow

import "strings"

const (
	EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStateActive     = "enterprise_workflow_authority_valc_closure_validation_enforcement_active"
	EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStatePartial    = "enterprise_workflow_authority_valc_closure_validation_enforcement_partial"
	EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStateIncomplete = "enterprise_workflow_authority_valc_closure_validation_enforcement_incomplete"

	EnterpriseWorkflowAuthorityValCWorkflowLedgerStateActive     = "enterprise_workflow_authority_valc_workflow_ledger_active"
	EnterpriseWorkflowAuthorityValCWorkflowLedgerStatePartial    = "enterprise_workflow_authority_valc_workflow_ledger_partial"
	EnterpriseWorkflowAuthorityValCWorkflowLedgerStateIncomplete = "enterprise_workflow_authority_valc_workflow_ledger_incomplete"

	EnterpriseWorkflowAuthorityValCStaleReopenHandlingStateActive     = "enterprise_workflow_authority_valc_stale_reopen_handling_active"
	EnterpriseWorkflowAuthorityValCStaleReopenHandlingStatePartial    = "enterprise_workflow_authority_valc_stale_reopen_handling_partial"
	EnterpriseWorkflowAuthorityValCStaleReopenHandlingStateIncomplete = "enterprise_workflow_authority_valc_stale_reopen_handling_incomplete"

	EnterpriseWorkflowAuthorityValCRollbackLinkageStateActive     = "enterprise_workflow_authority_valc_rollback_linkage_active"
	EnterpriseWorkflowAuthorityValCRollbackLinkageStatePartial    = "enterprise_workflow_authority_valc_rollback_linkage_partial"
	EnterpriseWorkflowAuthorityValCRollbackLinkageStateIncomplete = "enterprise_workflow_authority_valc_rollback_linkage_incomplete"

	EnterpriseWorkflowAuthorityValCGovernanceMappingStateActive     = "enterprise_workflow_authority_valc_governance_mapping_active"
	EnterpriseWorkflowAuthorityValCGovernanceMappingStatePartial    = "enterprise_workflow_authority_valc_governance_mapping_partial"
	EnterpriseWorkflowAuthorityValCGovernanceMappingStateIncomplete = "enterprise_workflow_authority_valc_governance_mapping_incomplete"

	EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStateActive     = "enterprise_workflow_authority_valc_replay_recovery_hardening_active"
	EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStatePartial    = "enterprise_workflow_authority_valc_replay_recovery_hardening_partial"
	EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStateIncomplete = "enterprise_workflow_authority_valc_replay_recovery_hardening_incomplete"

	EnterpriseWorkflowAuthorityValCStateIncomplete  = "enterprise_workflow_authority_valc_incomplete"
	EnterpriseWorkflowAuthorityValCStateSubstantial = "enterprise_workflow_authority_valc_substantially_ready"
	EnterpriseWorkflowAuthorityValCStateActive      = "enterprise_workflow_authority_valc_active"
)

type WorkflowClosureValidationEnforcementBaseline struct {
	CurrentState             string   `json:"current_state"`
	RequiredChecks           []string `json:"required_checks,omitempty"`
	ClosureAuthorityRules    []string `json:"closure_authority_rules,omitempty"`
	AlternativePolicyRules   []string `json:"alternative_policy_rules,omitempty"`
	FailureStateConsequences []string `json:"failure_state_consequences,omitempty"`
	ReopenRules              []string `json:"reopen_rules,omitempty"`
	RollbackConsequences     []string `json:"rollback_consequences,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type WorkflowAppendOnlyLedgerBaseline struct {
	CurrentState      string   `json:"current_state"`
	RequiredFields    []string `json:"required_fields,omitempty"`
	RecordTypes       []string `json:"record_types,omitempty"`
	AppendOnly        bool     `json:"append_only"`
	SignedEntries     bool     `json:"signed_entries"`
	SupersessionRules []string `json:"supersession_rules,omitempty"`
	RevocationRules   []string `json:"revocation_rules,omitempty"`
	Limitations       []string `json:"limitations,omitempty"`
}

type WorkflowStaleReopenHandlingBaseline struct {
	CurrentState           string   `json:"current_state"`
	StaleSignals           []string `json:"stale_signals,omitempty"`
	ReopenTriggers         []string `json:"reopen_triggers,omitempty"`
	ReopenEvidenceRules    []string `json:"reopen_evidence_rules,omitempty"`
	ClosedConflictEffects  []string `json:"closed_conflict_effects,omitempty"`
	ConnectorVisibility    []string `json:"connector_visibility,omitempty"`
	OperationalEffectRules []string `json:"operational_effect_rules,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type WorkflowRollbackLinkageBaseline struct {
	CurrentState           string   `json:"current_state"`
	RequiredRefs           []string `json:"required_refs,omitempty"`
	RollbackStates         []string `json:"rollback_states,omitempty"`
	ClosureConsequences    []string `json:"closure_consequences,omitempty"`
	ReopenRules            []string `json:"reopen_rules,omitempty"`
	AuthorityInteraction   []string `json:"authority_interaction_rules,omitempty"`
	ValidationLinkageRules []string `json:"validation_linkage_rules,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type WorkflowGovernanceMappingBaseline struct {
	CurrentState        string   `json:"current_state"`
	RequiredMappings    []string `json:"required_mappings,omitempty"`
	DecisionClasses     []string `json:"decision_classes,omitempty"`
	EvidenceRules       []string `json:"evidence_rules,omitempty"`
	ComplianceScopes    []string `json:"compliance_scopes,omitempty"`
	VisibilityRules     []string `json:"visibility_rules,omitempty"`
	SupersessionEffects []string `json:"supersession_effects,omitempty"`
	Limitations         []string `json:"limitations,omitempty"`
}

type WorkflowReplayRecoveryHardeningBaseline struct {
	CurrentState             string   `json:"current_state"`
	RequiredSources          []string `json:"required_sources,omitempty"`
	RecoveryRules            []string `json:"recovery_rules,omitempty"`
	DuplicateSuppression     []string `json:"duplicate_suppression_rules,omitempty"`
	OutageRecoveryRules      []string `json:"outage_recovery_rules,omitempty"`
	CanonicalPrecedenceRules []string `json:"canonical_precedence_rules,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

func EnterpriseWorkflowAuthorityValCClosureValidationEnforcement() WorkflowClosureValidationEnforcementBaseline {
	model := WorkflowClosureValidationEnforcementBaseline{
		RequiredChecks: []string{
			"remediation_declared",
			"remediation_evidence_present",
			"validation_result_verified",
			"freshness_check_passed",
			"no_still_active_expired_exception",
			"no_still_active_revoked_override_or_authorization",
			"closure_policy_satisfied",
		},
		ClosureAuthorityRules: []string{
			"administrative_external_close_never_equals_canonical_close",
			"validated_fixed_or_bounded_alternative_policy_required_before_closed",
			"superseded_authority_or_exception_requires_replacement_link_before_close",
		},
		AlternativePolicyRules: []string{
			"bounded_closure_override_must_reference_signed_authorization_and_policy_exception",
			"production_closure_override_requires_distinct_approver_and_executor",
		},
		FailureStateConsequences: []string{
			"expired_exception_blocks_or_reopens_close",
			"revoked_override_blocks_or_reopens_close",
			"superseded_authorization_requires_new_effective_artifact_before_close",
		},
		ReopenRules: []string{
			"reopen_requires_reason_and_evidence_ref",
			"later_conflict_or_failed_validation_reopens_canonical_workflow",
			"reopened_state_must_propagate_to_connector_projection_without_losing_prior_close_trace",
		},
		RollbackConsequences: []string{
			"rollback_applied_does_not_imply_close_without_validation",
			"rollback_after_close_reopens_canonical_workflow",
			"rollback_failure_blocks_close_and_marks_validation_pending",
		},
		Limitations: []string{
			"Val C hardens closure-by-validation and governance semantics before the later final workflow authority gate review.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValCClosureValidationEnforcementState(model)
	return model
}

func EnterpriseWorkflowAuthorityValCWorkflowLedger() WorkflowAppendOnlyLedgerBaseline {
	model := WorkflowAppendOnlyLedgerBaseline{
		RequiredFields: []string{
			"event_id",
			"actor_identity",
			"subject",
			"action_class",
			"scope",
			"evidence_refs",
			"external_system_refs",
			"resulting_canonical_transition",
			"resulting_external_mutation_outcome",
			"recorded_at",
		},
		RecordTypes: []string{
			"authorization_issued",
			"authorization_consumed",
			"exception_activated",
			"exception_revoked",
			"closure_validated",
			"workflow_reopened",
			"rollback_linked",
		},
		AppendOnly:    true,
		SignedEntries: true,
		SupersessionRules: []string{
			"superseded_ledger_entry_must_link_to_replacement_entry",
			"supersession_cannot_delete_or_hide_prior_decision_trace",
		},
		RevocationRules: []string{
			"revoked_authority_or_exception_effect_must_be_recorded_as_new_append_only_entry",
			"revocation_entry_must_remain_visible_to_replay_and_audit_paths",
		},
		Limitations: []string{
			"Val C defines append-only workflow ledger posture without yet declaring the final workflow authority gate complete.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValCWorkflowLedgerState(model)
	return model
}

func EnterpriseWorkflowAuthorityValCStaleReopenHandling() WorkflowStaleReopenHandlingBaseline {
	model := WorkflowStaleReopenHandlingBaseline{
		StaleSignals: []string{
			"stale_external_close_signal",
			"expired_exception_still_projected_active",
			"revoked_override_after_external_resolve",
			"rollback_signal_after_close",
			"validation_freshness_expired",
		},
		ReopenTriggers: []string{
			"later_conflict_evidence",
			"rollback_after_close",
			"expired_exception_while_subject_unresolved",
			"revoked_authority_during_close_pending",
		},
		ReopenEvidenceRules: []string{
			"reopen_requires_reason_and_evidence_ref",
			"reopen_must_link_to_prior_closure_or_validation_attempt",
			"external_closed_state_cannot_block_canonical_reopen",
		},
		ClosedConflictEffects: []string{
			"stale_or_conflicting_close_marks_canonical_workflow_reopened_or_pending_review",
			"external_projection_must_show_stale_or_reopened_signal",
		},
		ConnectorVisibility: []string{
			"jira_and_servicenow_must_show_reopen_or_stale_notice",
			"github_projection_must_preserve_reopen_and_rollback_visibility",
		},
		OperationalEffectRules: []string{
			"revoked_or_expired_authority_has_operational_effect_even_if_external_projection_is_stale",
			"superseded_close_notice_requires_replacement_link_visibility",
		},
		Limitations: []string{
			"Val C stale and reopen handling remains evidence-bound and does not let external systems suppress canonical reopen decisions.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValCStaleReopenHandlingState(model)
	return model
}

func EnterpriseWorkflowAuthorityValCRollbackLinkage() WorkflowRollbackLinkageBaseline {
	model := WorkflowRollbackLinkageBaseline{
		RequiredRefs: []string{
			"rollback_request_ref",
			"rollback_execution_ref",
			"rollback_validation_ref",
			"canonical_workflow_ref",
			"evidence_bundle_ref",
		},
		RollbackStates: []string{
			"rollback_requested",
			"rollback_applied",
			"rollback_validated",
			"rollback_rejected",
		},
		ClosureConsequences: []string{
			"rollback_applied_does_not_imply_close",
			"rollback_after_close_reopens_workflow",
			"rollback_rejected_blocks_close_and_keeps_validation_pending",
		},
		ReopenRules: []string{
			"rollback_triggered_reopen_must_link_to_prior_close_or_validation_ref",
			"reopen_after_rollback_must_remain_auditable",
		},
		AuthorityInteraction: []string{
			"rollback_using_break_glass_or_override_must_link_to_authorization_artifact",
			"superseded_or_revoked_authority_cannot_reuse_prior_rollback_linkage",
		},
		ValidationLinkageRules: []string{
			"rollback_validation_ref_must_show_result_before_reclose",
			"post_rollback_close_requires_fresh_validation_evidence",
		},
		Limitations: []string{
			"Val C rollback linkage defines operational and closure consequences before later Point 3 waves add the final workflow authority gate.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValCRollbackLinkageState(model)
	return model
}

func EnterpriseWorkflowAuthorityValCGovernanceMapping() WorkflowGovernanceMappingBaseline {
	model := WorkflowGovernanceMappingBaseline{
		RequiredMappings: []string{
			"compliance_control_ref",
			"policy_rule_ref",
			"exception_policy_ref",
			"approval_decision_ref",
			"closure_reason_taxonomy",
		},
		DecisionClasses: []string{
			"approval",
			"break_glass",
			"exception",
			"validation_close",
			"reopen",
			"rollback",
		},
		EvidenceRules: []string{
			"governance_mapping_must_link_to_evidence_bundle",
			"closure_validation_must_link_to_policy_and_control_mapping",
		},
		ComplianceScopes: []string{
			"change_governance",
			"exception_governance",
			"closure_governance",
		},
		VisibilityRules: []string{
			"revoked_and_expired_states_remain_governance_visible",
			"reopened_and_rollback_states_must_remain_traceable_in_governance_projection",
		},
		SupersessionEffects: []string{
			"superseded_decision_must_link_to_replacement_mapping",
			"superseded_mapping_cannot_erase_prior_audit_trace",
		},
		Limitations: []string{
			"Val C governance mapping remains a hardened traceability layer and does not replace the later final workflow authority gate review.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValCGovernanceMappingState(model)
	return model
}

func EnterpriseWorkflowAuthorityValCReplayRecoveryHardening() WorkflowReplayRecoveryHardeningBaseline {
	model := WorkflowReplayRecoveryHardeningBaseline{
		RequiredSources: []string{
			"canonical_event_log",
			"connector_delivery_log",
			"authorization_consumption_log",
		},
		RecoveryRules: []string{
			"replay_must_reconstruct_last_authoritative_transition_before_projection",
			"duplicate_external_close_must_not_skip_validation_or_reopen_requirements",
			"recovery_must_preserve_revoked_expired_and_superseded_effects",
		},
		DuplicateSuppression: []string{
			"suppress_duplicate_connector_delivery_replay",
			"reject_duplicate_authorization_consumption_replay",
			"replayed_rollback_signal_must_not_duplicate_reopen_effect",
		},
		OutageRecoveryRules: []string{
			"degraded_connector_mode_marks_recovery_required_not_authoritative_success",
			"outage_recovery_must_reconcile_external_projection_after_canonical_state_is_restored",
		},
		CanonicalPrecedenceRules: []string{
			"external_close_never_overwrites_canonical_validation_state",
			"revoked_or_expired_authority_effect_overrides_stale_external_success_signal",
		},
		Limitations: []string{
			"Val C replay and recovery hardening protects canonical authority during connector drift before the final Point 3 gate review.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValCReplayRecoveryHardeningState(model)
	return model
}

func EvaluateEnterpriseWorkflowAuthorityValCClosureValidationEnforcementState(model WorkflowClosureValidationEnforcementBaseline) string {
	if len(model.RequiredChecks) == 0 {
		return EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredChecks,
		"remediation_declared",
		"remediation_evidence_present",
		"validation_result_verified",
		"freshness_check_passed",
		"no_still_active_expired_exception",
		"no_still_active_revoked_override_or_authorization",
		"closure_policy_satisfied",
	) || len(model.ClosureAuthorityRules) == 0 || len(model.AlternativePolicyRules) == 0 || !containsAllTrimmedStrings(model.FailureStateConsequences,
		"expired_exception_blocks_or_reopens_close",
		"revoked_override_blocks_or_reopens_close",
		"superseded_authorization_requires_new_effective_artifact_before_close",
	) || len(model.ReopenRules) == 0 || len(model.RollbackConsequences) == 0 {
		return EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStatePartial
	}
	return EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValCWorkflowLedgerState(model WorkflowAppendOnlyLedgerBaseline) string {
	if len(model.RequiredFields) == 0 || len(model.RecordTypes) == 0 {
		return EnterpriseWorkflowAuthorityValCWorkflowLedgerStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredFields,
		"event_id",
		"actor_identity",
		"subject",
		"action_class",
		"scope",
		"evidence_refs",
		"external_system_refs",
		"resulting_canonical_transition",
		"resulting_external_mutation_outcome",
		"recorded_at",
	) || !containsAllTrimmedStrings(model.RecordTypes,
		"authorization_issued",
		"authorization_consumed",
		"exception_activated",
		"exception_revoked",
		"closure_validated",
		"workflow_reopened",
		"rollback_linked",
	) || !model.AppendOnly || !model.SignedEntries || len(model.SupersessionRules) == 0 || len(model.RevocationRules) == 0 {
		return EnterpriseWorkflowAuthorityValCWorkflowLedgerStatePartial
	}
	return EnterpriseWorkflowAuthorityValCWorkflowLedgerStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValCStaleReopenHandlingState(model WorkflowStaleReopenHandlingBaseline) string {
	if len(model.StaleSignals) == 0 || len(model.ReopenTriggers) == 0 {
		return EnterpriseWorkflowAuthorityValCStaleReopenHandlingStateIncomplete
	}
	if !containsAllTrimmedStrings(model.StaleSignals,
		"stale_external_close_signal",
		"expired_exception_still_projected_active",
		"revoked_override_after_external_resolve",
		"rollback_signal_after_close",
		"validation_freshness_expired",
	) || !containsAllTrimmedStrings(model.ReopenEvidenceRules,
		"reopen_requires_reason_and_evidence_ref",
		"reopen_must_link_to_prior_closure_or_validation_attempt",
		"external_closed_state_cannot_block_canonical_reopen",
	) || len(model.ClosedConflictEffects) == 0 || len(model.ConnectorVisibility) == 0 || len(model.OperationalEffectRules) == 0 {
		return EnterpriseWorkflowAuthorityValCStaleReopenHandlingStatePartial
	}
	return EnterpriseWorkflowAuthorityValCStaleReopenHandlingStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValCRollbackLinkageState(model WorkflowRollbackLinkageBaseline) string {
	if len(model.RequiredRefs) == 0 || len(model.RollbackStates) == 0 {
		return EnterpriseWorkflowAuthorityValCRollbackLinkageStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredRefs,
		"rollback_request_ref",
		"rollback_execution_ref",
		"rollback_validation_ref",
		"canonical_workflow_ref",
		"evidence_bundle_ref",
	) || !containsAllTrimmedStrings(model.RollbackStates,
		"rollback_requested",
		"rollback_applied",
		"rollback_validated",
		"rollback_rejected",
	) || !containsAllTrimmedStrings(model.ClosureConsequences,
		"rollback_applied_does_not_imply_close",
		"rollback_after_close_reopens_workflow",
		"rollback_rejected_blocks_close_and_keeps_validation_pending",
	) || len(model.ReopenRules) == 0 || len(model.AuthorityInteraction) == 0 || len(model.ValidationLinkageRules) == 0 {
		return EnterpriseWorkflowAuthorityValCRollbackLinkageStatePartial
	}
	return EnterpriseWorkflowAuthorityValCRollbackLinkageStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValCGovernanceMappingState(model WorkflowGovernanceMappingBaseline) string {
	if len(model.RequiredMappings) == 0 || len(model.DecisionClasses) == 0 {
		return EnterpriseWorkflowAuthorityValCGovernanceMappingStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredMappings,
		"compliance_control_ref",
		"policy_rule_ref",
		"exception_policy_ref",
		"approval_decision_ref",
		"closure_reason_taxonomy",
	) || !containsAllTrimmedStrings(model.DecisionClasses,
		"approval",
		"break_glass",
		"exception",
		"validation_close",
		"reopen",
		"rollback",
	) || len(model.EvidenceRules) == 0 || len(model.ComplianceScopes) == 0 || len(model.VisibilityRules) == 0 || len(model.SupersessionEffects) == 0 {
		return EnterpriseWorkflowAuthorityValCGovernanceMappingStatePartial
	}
	return EnterpriseWorkflowAuthorityValCGovernanceMappingStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValCReplayRecoveryHardeningState(model WorkflowReplayRecoveryHardeningBaseline) string {
	if len(model.RequiredSources) == 0 || len(model.RecoveryRules) == 0 {
		return EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredSources,
		"canonical_event_log",
		"connector_delivery_log",
		"authorization_consumption_log",
	) || len(model.DuplicateSuppression) == 0 || len(model.OutageRecoveryRules) == 0 || !containsAllTrimmedStrings(model.CanonicalPrecedenceRules,
		"external_close_never_overwrites_canonical_validation_state",
		"revoked_or_expired_authority_effect_overrides_stale_external_success_signal",
	) {
		return EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStatePartial
	}
	return EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValCState(
	valBState,
	closureValidationEnforcementState,
	workflowLedgerState,
	staleReopenHandlingState,
	rollbackLinkageState,
	governanceMappingState,
	replayRecoveryHardeningState string,
) string {
	if strings.TrimSpace(valBState) != EnterpriseWorkflowAuthorityValBStateActive {
		return EnterpriseWorkflowAuthorityValCStateIncomplete
	}

	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(closureValidationEnforcementState),
		strings.TrimSpace(workflowLedgerState),
		strings.TrimSpace(staleReopenHandlingState),
		strings.TrimSpace(rollbackLinkageState),
		strings.TrimSpace(governanceMappingState),
		strings.TrimSpace(replayRecoveryHardeningState),
	} {
		switch state {
		case EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStateActive,
			EnterpriseWorkflowAuthorityValCWorkflowLedgerStateActive,
			EnterpriseWorkflowAuthorityValCStaleReopenHandlingStateActive,
			EnterpriseWorkflowAuthorityValCRollbackLinkageStateActive,
			EnterpriseWorkflowAuthorityValCGovernanceMappingStateActive,
			EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStateActive:
		case EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStatePartial,
			EnterpriseWorkflowAuthorityValCWorkflowLedgerStatePartial,
			EnterpriseWorkflowAuthorityValCStaleReopenHandlingStatePartial,
			EnterpriseWorkflowAuthorityValCRollbackLinkageStatePartial,
			EnterpriseWorkflowAuthorityValCGovernanceMappingStatePartial,
			EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStatePartial:
			hasPartial = true
		default:
			return EnterpriseWorkflowAuthorityValCStateIncomplete
		}
	}
	if hasPartial {
		return EnterpriseWorkflowAuthorityValCStateSubstantial
	}
	return EnterpriseWorkflowAuthorityValCStateActive
}

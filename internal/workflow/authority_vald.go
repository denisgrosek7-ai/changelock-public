package workflow

import "strings"

const (
	EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStateActive     = "enterprise_workflow_authority_vald_connector_correctness_review_active"
	EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStatePartial    = "enterprise_workflow_authority_vald_connector_correctness_review_partial"
	EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStateIncomplete = "enterprise_workflow_authority_vald_connector_correctness_review_incomplete"

	EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStateActive     = "enterprise_workflow_authority_vald_approval_boundary_review_active"
	EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStatePartial    = "enterprise_workflow_authority_vald_approval_boundary_review_partial"
	EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStateIncomplete = "enterprise_workflow_authority_vald_approval_boundary_review_incomplete"

	EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStateActive     = "enterprise_workflow_authority_vald_exception_expiry_review_active"
	EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStatePartial    = "enterprise_workflow_authority_vald_exception_expiry_review_partial"
	EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStateIncomplete = "enterprise_workflow_authority_vald_exception_expiry_review_incomplete"

	EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStateActive     = "enterprise_workflow_authority_vald_closure_correctness_review_active"
	EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStatePartial    = "enterprise_workflow_authority_vald_closure_correctness_review_partial"
	EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStateIncomplete = "enterprise_workflow_authority_vald_closure_correctness_review_incomplete"

	EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStateActive     = "enterprise_workflow_authority_vald_reconciliation_conflict_review_active"
	EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStatePartial    = "enterprise_workflow_authority_vald_reconciliation_conflict_review_partial"
	EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStateIncomplete = "enterprise_workflow_authority_vald_reconciliation_conflict_review_incomplete"

	EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStateActive     = "enterprise_workflow_authority_vald_workflow_ledger_review_active"
	EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStatePartial    = "enterprise_workflow_authority_vald_workflow_ledger_review_partial"
	EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStateIncomplete = "enterprise_workflow_authority_vald_workflow_ledger_review_incomplete"

	EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStateActive     = "enterprise_workflow_authority_vald_governance_traceability_review_active"
	EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStatePartial    = "enterprise_workflow_authority_vald_governance_traceability_review_partial"
	EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStateIncomplete = "enterprise_workflow_authority_vald_governance_traceability_review_incomplete"

	EnterpriseWorkflowAuthorityValDReopenRollbackReviewStateActive     = "enterprise_workflow_authority_vald_reopen_rollback_review_active"
	EnterpriseWorkflowAuthorityValDReopenRollbackReviewStatePartial    = "enterprise_workflow_authority_vald_reopen_rollback_review_partial"
	EnterpriseWorkflowAuthorityValDReopenRollbackReviewStateIncomplete = "enterprise_workflow_authority_vald_reopen_rollback_review_incomplete"

	EnterpriseWorkflowAuthorityValDStateIncomplete  = "enterprise_workflow_authority_vald_incomplete"
	EnterpriseWorkflowAuthorityValDStateSubstantial = "enterprise_workflow_authority_vald_substantially_ready"
	EnterpriseWorkflowAuthorityValDStateActive      = "enterprise_workflow_authority_vald_active"
)

type WorkflowConnectorCorrectnessReviewBaseline struct {
	CurrentState            string   `json:"current_state"`
	RequiredConnectors      []string `json:"required_connectors,omitempty"`
	SyncBackAllowlistRules  []string `json:"sync_back_allowlist_rules,omitempty"`
	IdempotentMutationRules []string `json:"idempotent_mutation_rules,omitempty"`
	ConflictPrecedenceRules []string `json:"conflict_precedence_rules,omitempty"`
	DegradedModeRules       []string `json:"degraded_mode_rules,omitempty"`
	Limitations             []string `json:"limitations,omitempty"`
}

type WorkflowApprovalBoundaryReviewBaseline struct {
	CurrentState              string   `json:"current_state"`
	RequiredActionClasses     []string `json:"required_action_classes,omitempty"`
	RequiredConsumptionModes  []string `json:"required_consumption_modes,omitempty"`
	RequiredBoundaryRules     []string `json:"required_boundary_rules,omitempty"`
	SeparationOfDutiesChecks  []string `json:"separation_of_duties_checks,omitempty"`
	RevocationAndExpiryChecks []string `json:"revocation_and_expiry_checks,omitempty"`
	Limitations               []string `json:"limitations,omitempty"`
}

type WorkflowExceptionExpiryReviewBaseline struct {
	CurrentState            string   `json:"current_state"`
	RequiredLifecycleStages []string `json:"required_lifecycle_stages,omitempty"`
	ExpiryEffectRules       []string `json:"expiry_effect_rules,omitempty"`
	RevocationEffectRules   []string `json:"revocation_effect_rules,omitempty"`
	SupersessionEffectRules []string `json:"supersession_effect_rules,omitempty"`
	RevalidationRules       []string `json:"revalidation_rules,omitempty"`
	Limitations             []string `json:"limitations,omitempty"`
}

type WorkflowClosureCorrectnessReviewBaseline struct {
	CurrentState             string   `json:"current_state"`
	RequiredChecks           []string `json:"required_checks,omitempty"`
	FailureStateConsequences []string `json:"failure_state_consequences,omitempty"`
	AdministrativeCloseRules []string `json:"administrative_close_rules,omitempty"`
	AlternativePolicyRules   []string `json:"alternative_policy_rules,omitempty"`
	ValidationLinkageRules   []string `json:"validation_linkage_rules,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type WorkflowReconciliationConflictReviewBaseline struct {
	CurrentState             string   `json:"current_state"`
	RequiredSignals          []string `json:"required_signals,omitempty"`
	ConflictResolutionRules  []string `json:"conflict_resolution_rules,omitempty"`
	ReplayRecoveryRules      []string `json:"replay_recovery_rules,omitempty"`
	OutageRecoveryRules      []string `json:"outage_recovery_rules,omitempty"`
	CanonicalPrecedenceRules []string `json:"canonical_precedence_rules,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type WorkflowLedgerReviewBaseline struct {
	CurrentState         string   `json:"current_state"`
	RequiredRecordTypes  []string `json:"required_record_types,omitempty"`
	RequiredRecordFields []string `json:"required_record_fields,omitempty"`
	AppendOnlyChecks     []string `json:"append_only_checks,omitempty"`
	SignedEntryChecks    []string `json:"signed_entry_checks,omitempty"`
	SupersessionChecks   []string `json:"supersession_checks,omitempty"`
	RevocationChecks     []string `json:"revocation_checks,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
}

type WorkflowGovernanceTraceabilityReviewBaseline struct {
	CurrentState            string   `json:"current_state"`
	RequiredMappings        []string `json:"required_mappings,omitempty"`
	RequiredDecisionClasses []string `json:"required_decision_classes,omitempty"`
	EvidenceLinkageRules    []string `json:"evidence_linkage_rules,omitempty"`
	ComplianceScopeRules    []string `json:"compliance_scope_rules,omitempty"`
	VisibilityRules         []string `json:"visibility_rules,omitempty"`
	Limitations             []string `json:"limitations,omitempty"`
}

type WorkflowReopenRollbackReviewBaseline struct {
	CurrentState             string   `json:"current_state"`
	RequiredReopenRules      []string `json:"required_reopen_rules,omitempty"`
	RequiredRollbackStates   []string `json:"required_rollback_states,omitempty"`
	OperationalEffectRules   []string `json:"operational_effect_rules,omitempty"`
	ValidationConsequences   []string `json:"validation_consequences,omitempty"`
	ConnectorVisibilityRules []string `json:"connector_visibility_rules,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

func EnterpriseWorkflowAuthorityValDConnectorCorrectnessReview() WorkflowConnectorCorrectnessReviewBaseline {
	model := WorkflowConnectorCorrectnessReviewBaseline{
		RequiredConnectors: []string{
			WorkflowAuthorityConnectorJira,
			WorkflowAuthorityConnectorServiceNow,
			WorkflowAuthorityConnectorGitHub,
		},
		SyncBackAllowlistRules: []string{
			"sync_back_is_limited_to_source_specific_allowlisted_fields",
			"external_projection_cannot_expand_scope_or_authority_class",
		},
		IdempotentMutationRules: []string{
			"connector_mutations_require_idempotency_keys",
			"duplicate_delivery_or_write_attempts_must_be_suppressed",
		},
		ConflictPrecedenceRules: []string{
			"canonical_state_has_precedence_over_external_close_or_resolve_labels",
			"stale_external_state_must_be_marked_not_promoted_to_truth",
		},
		DegradedModeRules: []string{
			"connector_outage_preserves_canonical_authority",
			"degraded_mode_requires_replay_or_recovery_path_visibility",
		},
		Limitations: []string{
			"Val D connector correctness review validates bounded connector behavior and does not create a new orchestration authority layer outside the canonical workflow spine.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewState(model)
	return model
}

func EnterpriseWorkflowAuthorityValDApprovalBoundaryReview() WorkflowApprovalBoundaryReviewBaseline {
	model := WorkflowApprovalBoundaryReviewBaseline{
		RequiredActionClasses: []string{
			WorkflowAuthorityActionApprovalRequired,
			WorkflowAuthoritySensitiveActionBreakGlass,
			WorkflowAuthoritySensitiveActionBroadScopeOverride,
			WorkflowAuthoritySensitiveActionProductionClosure,
			WorkflowAuthoritySensitiveActionLongLivedException,
		},
		RequiredConsumptionModes: []string{
			WorkflowAuthorityConsumptionSingleUse,
			WorkflowAuthorityConsumptionMultiUseBounded,
			WorkflowAuthorityConsumptionSessionBound,
		},
		RequiredBoundaryRules: []string{
			"approval_is_identity_subject_scope_and_time_bound",
			"approval_cannot_be_replayed_after_consumption_or_revocation",
			"external_projection_is_not_equivalent_to_canonical_approval",
		},
		SeparationOfDutiesChecks: []string{
			"break_glass_requires_distinct_approver_executor_or_dual_control",
			"production_closure_override_requires_dual_control",
		},
		RevocationAndExpiryChecks: []string{
			"expired_or_revoked_authority_must_fail_closed",
			"superseded_authority_requires_replacement_artifact_before_effect",
		},
		Limitations: []string{
			"Val D approval boundary review confirms bounded authority posture and does not widen action classes beyond the delegated-authority model already defined in Val B.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValDApprovalBoundaryReviewState(model)
	return model
}

func EnterpriseWorkflowAuthorityValDExceptionExpiryReview() WorkflowExceptionExpiryReviewBaseline {
	model := WorkflowExceptionExpiryReviewBaseline{
		RequiredLifecycleStages: []string{
			"requested",
			"approved",
			"activated",
			"expiring",
			"expired",
			"revoked",
			"superseded",
			"revalidated",
		},
		ExpiryEffectRules: []string{
			"expired_exception_blocks_or_reopens_subject",
			"grace_denied_without_fresh_revalidation",
		},
		RevocationEffectRules: []string{
			"revoked_exception_immediately_loses_operational_effect",
			"revoked_exception_remains_visible_in_governance_and_connector_projection",
		},
		SupersessionEffectRules: []string{
			"superseded_exception_must_link_to_replacement_object",
			"replacement_exception_requires_new_evidence_and_expiry_window",
		},
		RevalidationRules: []string{
			"revalidation_requires_fresh_evidence_and_policy_link",
			"revalidation_issues_new_effective_window_not_silent_extension",
		},
		Limitations: []string{
			"Val D exception expiry review validates operational effects and does not replace the managed exception registry itself.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValDExceptionExpiryReviewState(model)
	return model
}

func EnterpriseWorkflowAuthorityValDClosureCorrectnessReview() WorkflowClosureCorrectnessReviewBaseline {
	model := WorkflowClosureCorrectnessReviewBaseline{
		RequiredChecks: []string{
			"remediation_declared",
			"remediation_evidence_present",
			"validation_result_verified",
			"freshness_check_passed",
			"closure_policy_satisfied",
		},
		FailureStateConsequences: []string{
			"expired_exception_blocks_or_reopens_close",
			"revoked_override_blocks_or_reopens_close",
			"superseded_authorization_requires_replacement_before_close",
			"rollback_after_close_reopens_workflow",
		},
		AdministrativeCloseRules: []string{
			"external_resolved_or_closed_state_is_not_canonical_close",
			"administrative_close_cannot_bypass_validation_or_governance_failures",
		},
		AlternativePolicyRules: []string{
			"bounded_closure_override_must_reference_signed_authorization",
			"production_closure_override_requires_distinct_human_control",
		},
		ValidationLinkageRules: []string{
			"closure_links_to_validation_evidence_and_policy_mapping",
			"post_rollback_close_requires_fresh_validation_result",
		},
		Limitations: []string{
			"Val D closure correctness review validates canonical close semantics without creating a new mutable closure database.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValDClosureCorrectnessReviewState(model)
	return model
}

func EnterpriseWorkflowAuthorityValDReconciliationConflictReview() WorkflowReconciliationConflictReviewBaseline {
	model := WorkflowReconciliationConflictReviewBaseline{
		RequiredSignals: []string{
			"external_closed_without_validation",
			"stale_external_state",
			"revoked_or_expired_authority_after_projection_success",
			"rollback_signal_after_close",
			"duplicate_delivery_or_replay_attempt",
		},
		ConflictResolutionRules: []string{
			"canonical_state_wins_over_external_success_or_resolution_labels",
			"conflicts_trigger_review_or_reopen_not_silent_acceptance",
		},
		ReplayRecoveryRules: []string{
			"recovery_reconstructs_last_authoritative_transition_before_projection",
			"replay_preserves_revoked_expired_superseded_effects",
		},
		OutageRecoveryRules: []string{
			"connector_outage_cannot_mark_workflow_authoritatively_successful",
			"recovery_path_must_reconcile_projection_after_canonical_restore",
		},
		CanonicalPrecedenceRules: []string{
			"external_close_never_overwrites_canonical_validation_state",
			"revoked_or_expired_authority_effect_overrides_stale_external_success_signal",
		},
		Limitations: []string{
			"Val D reconciliation conflict review confirms fail-closed conflict posture and does not replace the underlying connector reconciliation engine.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValDReconciliationConflictReviewState(model)
	return model
}

func EnterpriseWorkflowAuthorityValDWorkflowLedgerReview() WorkflowLedgerReviewBaseline {
	model := WorkflowLedgerReviewBaseline{
		RequiredRecordTypes: []string{
			"authorization_issued",
			"authorization_consumed",
			"exception_activated",
			"exception_revoked",
			"closure_validated",
			"workflow_reopened",
			"rollback_linked",
		},
		RequiredRecordFields: []string{
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
		AppendOnlyChecks: []string{
			"original_entry_remains_visible_after_supersession",
			"corrective_or_withdrawn_effect_is_recorded_as_new_entry_not_mutation",
		},
		SignedEntryChecks: []string{
			"signed_entries_bind_actor_and_effect",
			"signed_entries_remain_verifiable_during_replay_and_audit",
		},
		SupersessionChecks: []string{
			"superseding_entry_links_to_original_decision",
			"supersession_does_not_delete_prior_event_trace",
		},
		RevocationChecks: []string{
			"revocation_is_recorded_as_new_append_only_event",
			"revoked_effect_and_original_effect_remain_distinctly_visible",
		},
		Limitations: []string{
			"Val D workflow ledger review validates semantic separation between original, corrective, superseding, and withdrawn effects without introducing a different ledger model.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValDWorkflowLedgerReviewState(model)
	return model
}

func EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReview() WorkflowGovernanceTraceabilityReviewBaseline {
	model := WorkflowGovernanceTraceabilityReviewBaseline{
		RequiredMappings: []string{
			"compliance_control_ref",
			"policy_rule_ref",
			"exception_policy_ref",
			"approval_decision_ref",
			"closure_reason_taxonomy",
		},
		RequiredDecisionClasses: []string{
			"approval",
			"break_glass",
			"exception",
			"validation_close",
			"reopen",
			"rollback",
		},
		EvidenceLinkageRules: []string{
			"governance_mapping_links_to_evidence_bundle",
			"closure_validation_links_to_policy_and_control_mapping",
			"reopen_and_rollback_links_preserve_governance_reasoning",
		},
		ComplianceScopeRules: []string{
			"change_governance_scope_is_visible",
			"exception_governance_scope_is_visible",
			"closure_governance_scope_is_visible",
		},
		VisibilityRules: []string{
			"revoked_expired_reopened_and_rollback_states_remain_traceable",
			"superseded_governance_mappings_link_to_replacements",
		},
		Limitations: []string{
			"Val D governance traceability review validates evidence and policy lineage without claiming external systems as a new source of governance truth.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewState(model)
	return model
}

func EnterpriseWorkflowAuthorityValDReopenRollbackReview() WorkflowReopenRollbackReviewBaseline {
	model := WorkflowReopenRollbackReviewBaseline{
		RequiredReopenRules: []string{
			"reopen_requires_reason_and_evidence_ref",
			"external_closed_state_cannot_block_canonical_reopen",
			"reopened_state_must_propagate_to_connector_projection_without_losing_prior_close_trace",
		},
		RequiredRollbackStates: []string{
			"rollback_requested",
			"rollback_applied",
			"rollback_validated",
			"rollback_rejected",
		},
		OperationalEffectRules: []string{
			"expired_or_revoked_authority_changes_operational_effect_even_if_projection_is_stale",
			"rollback_after_close_has_reopen_effect",
			"superseded_authority_or_exception_requires_visible_replacement_link",
		},
		ValidationConsequences: []string{
			"post_rollback_close_requires_fresh_validation",
			"rollback_rejected_keeps_workflow_validation_pending",
		},
		ConnectorVisibilityRules: []string{
			"jira_and_servicenow_show_reopen_or_stale_notice",
			"github_projection_preserves_reopen_and_rollback_visibility",
		},
		Limitations: []string{
			"Val D reopen and rollback review validates final pre-gate consistency and does not yet perform the final workflow authority gate aggregation itself.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValDReopenRollbackReviewState(model)
	return model
}

func EvaluateEnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewState(model WorkflowConnectorCorrectnessReviewBaseline) string {
	if len(model.RequiredConnectors) == 0 {
		return EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredConnectors,
		WorkflowAuthorityConnectorJira,
		WorkflowAuthorityConnectorServiceNow,
		WorkflowAuthorityConnectorGitHub,
	) || len(model.SyncBackAllowlistRules) == 0 || len(model.IdempotentMutationRules) == 0 || len(model.ConflictPrecedenceRules) == 0 || len(model.DegradedModeRules) == 0 {
		return EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStatePartial
	}
	return EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValDApprovalBoundaryReviewState(model WorkflowApprovalBoundaryReviewBaseline) string {
	if len(model.RequiredActionClasses) == 0 || len(model.RequiredConsumptionModes) == 0 {
		return EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredActionClasses,
		WorkflowAuthorityActionApprovalRequired,
		WorkflowAuthoritySensitiveActionBreakGlass,
		WorkflowAuthoritySensitiveActionBroadScopeOverride,
		WorkflowAuthoritySensitiveActionProductionClosure,
		WorkflowAuthoritySensitiveActionLongLivedException,
	) || !containsExactTrimmedStringSet(model.RequiredConsumptionModes,
		WorkflowAuthorityConsumptionSingleUse,
		WorkflowAuthorityConsumptionMultiUseBounded,
		WorkflowAuthorityConsumptionSessionBound,
	) || len(model.RequiredBoundaryRules) == 0 || len(model.SeparationOfDutiesChecks) == 0 || len(model.RevocationAndExpiryChecks) == 0 {
		return EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStatePartial
	}
	return EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValDExceptionExpiryReviewState(model WorkflowExceptionExpiryReviewBaseline) string {
	if len(model.RequiredLifecycleStages) == 0 {
		return EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredLifecycleStages,
		"requested",
		"approved",
		"activated",
		"expiring",
		"expired",
		"revoked",
		"superseded",
		"revalidated",
	) || len(model.ExpiryEffectRules) == 0 || len(model.RevocationEffectRules) == 0 || len(model.SupersessionEffectRules) == 0 || len(model.RevalidationRules) == 0 {
		return EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStatePartial
	}
	return EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValDClosureCorrectnessReviewState(model WorkflowClosureCorrectnessReviewBaseline) string {
	if len(model.RequiredChecks) == 0 {
		return EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredChecks,
		"remediation_declared",
		"remediation_evidence_present",
		"validation_result_verified",
		"freshness_check_passed",
		"closure_policy_satisfied",
	) || !containsAllTrimmedStrings(model.FailureStateConsequences,
		"expired_exception_blocks_or_reopens_close",
		"revoked_override_blocks_or_reopens_close",
		"superseded_authorization_requires_replacement_before_close",
		"rollback_after_close_reopens_workflow",
	) || len(model.AdministrativeCloseRules) == 0 || len(model.AlternativePolicyRules) == 0 || len(model.ValidationLinkageRules) == 0 {
		return EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStatePartial
	}
	return EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValDReconciliationConflictReviewState(model WorkflowReconciliationConflictReviewBaseline) string {
	if len(model.RequiredSignals) == 0 {
		return EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredSignals,
		"external_closed_without_validation",
		"stale_external_state",
		"revoked_or_expired_authority_after_projection_success",
		"rollback_signal_after_close",
		"duplicate_delivery_or_replay_attempt",
	) || len(model.ConflictResolutionRules) == 0 || len(model.ReplayRecoveryRules) == 0 || len(model.OutageRecoveryRules) == 0 || !containsAllTrimmedStrings(model.CanonicalPrecedenceRules,
		"external_close_never_overwrites_canonical_validation_state",
		"revoked_or_expired_authority_effect_overrides_stale_external_success_signal",
	) {
		return EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStatePartial
	}
	return EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValDWorkflowLedgerReviewState(model WorkflowLedgerReviewBaseline) string {
	if len(model.RequiredRecordTypes) == 0 || len(model.RequiredRecordFields) == 0 {
		return EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredRecordTypes,
		"authorization_issued",
		"authorization_consumed",
		"exception_activated",
		"exception_revoked",
		"closure_validated",
		"workflow_reopened",
		"rollback_linked",
	) || !containsAllTrimmedStrings(model.RequiredRecordFields,
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
	) || len(model.AppendOnlyChecks) == 0 || len(model.SignedEntryChecks) == 0 || len(model.SupersessionChecks) == 0 || len(model.RevocationChecks) == 0 {
		return EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStatePartial
	}
	return EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewState(model WorkflowGovernanceTraceabilityReviewBaseline) string {
	if len(model.RequiredMappings) == 0 || len(model.RequiredDecisionClasses) == 0 {
		return EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredMappings,
		"compliance_control_ref",
		"policy_rule_ref",
		"exception_policy_ref",
		"approval_decision_ref",
		"closure_reason_taxonomy",
	) || !containsAllTrimmedStrings(model.RequiredDecisionClasses,
		"approval",
		"break_glass",
		"exception",
		"validation_close",
		"reopen",
		"rollback",
	) || len(model.EvidenceLinkageRules) == 0 || len(model.ComplianceScopeRules) == 0 || len(model.VisibilityRules) == 0 {
		return EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStatePartial
	}
	return EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValDReopenRollbackReviewState(model WorkflowReopenRollbackReviewBaseline) string {
	if len(model.RequiredReopenRules) == 0 || len(model.RequiredRollbackStates) == 0 {
		return EnterpriseWorkflowAuthorityValDReopenRollbackReviewStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredReopenRules,
		"reopen_requires_reason_and_evidence_ref",
		"external_closed_state_cannot_block_canonical_reopen",
		"reopened_state_must_propagate_to_connector_projection_without_losing_prior_close_trace",
	) || !containsExactTrimmedStringSet(model.RequiredRollbackStates,
		"rollback_requested",
		"rollback_applied",
		"rollback_validated",
		"rollback_rejected",
	) || len(model.OperationalEffectRules) == 0 || len(model.ValidationConsequences) == 0 || len(model.ConnectorVisibilityRules) == 0 {
		return EnterpriseWorkflowAuthorityValDReopenRollbackReviewStatePartial
	}
	return EnterpriseWorkflowAuthorityValDReopenRollbackReviewStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValDState(
	valCState,
	connectorCorrectnessReviewState,
	approvalBoundaryReviewState,
	exceptionExpiryReviewState,
	closureCorrectnessReviewState,
	reconciliationConflictReviewState,
	workflowLedgerReviewState,
	governanceTraceabilityReviewState,
	reopenRollbackReviewState string,
) string {
	if strings.TrimSpace(valCState) != EnterpriseWorkflowAuthorityValCStateActive {
		return EnterpriseWorkflowAuthorityValDStateIncomplete
	}

	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(connectorCorrectnessReviewState),
		strings.TrimSpace(approvalBoundaryReviewState),
		strings.TrimSpace(exceptionExpiryReviewState),
		strings.TrimSpace(closureCorrectnessReviewState),
		strings.TrimSpace(reconciliationConflictReviewState),
		strings.TrimSpace(workflowLedgerReviewState),
		strings.TrimSpace(governanceTraceabilityReviewState),
		strings.TrimSpace(reopenRollbackReviewState),
	} {
		switch state {
		case EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStateActive,
			EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStateActive,
			EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStateActive,
			EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStateActive,
			EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStateActive,
			EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStateActive,
			EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStateActive,
			EnterpriseWorkflowAuthorityValDReopenRollbackReviewStateActive:
		case EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStatePartial,
			EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStatePartial,
			EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStatePartial,
			EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStatePartial,
			EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStatePartial,
			EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStatePartial,
			EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStatePartial,
			EnterpriseWorkflowAuthorityValDReopenRollbackReviewStatePartial:
			hasPartial = true
		default:
			return EnterpriseWorkflowAuthorityValDStateIncomplete
		}
	}
	if hasPartial {
		return EnterpriseWorkflowAuthorityValDStateSubstantial
	}
	return EnterpriseWorkflowAuthorityValDStateActive
}

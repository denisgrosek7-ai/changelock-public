package workflow

import "strings"

const (
	EnterpriseWorkflowAuthorityValAEventOrchestrationStateActive     = "enterprise_workflow_authority_vala_event_orchestration_active"
	EnterpriseWorkflowAuthorityValAEventOrchestrationStatePartial    = "enterprise_workflow_authority_vala_event_orchestration_partial"
	EnterpriseWorkflowAuthorityValAEventOrchestrationStateIncomplete = "enterprise_workflow_authority_vala_event_orchestration_incomplete"

	EnterpriseWorkflowAuthorityValALifecycleConnectorsStateActive     = "enterprise_workflow_authority_vala_lifecycle_connectors_active"
	EnterpriseWorkflowAuthorityValALifecycleConnectorsStatePartial    = "enterprise_workflow_authority_vala_lifecycle_connectors_partial"
	EnterpriseWorkflowAuthorityValALifecycleConnectorsStateIncomplete = "enterprise_workflow_authority_vala_lifecycle_connectors_incomplete"

	EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStateActive     = "enterprise_workflow_authority_vala_evidence_bundle_injection_active"
	EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStatePartial    = "enterprise_workflow_authority_vala_evidence_bundle_injection_partial"
	EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStateIncomplete = "enterprise_workflow_authority_vala_evidence_bundle_injection_incomplete"

	EnterpriseWorkflowAuthorityValATicketChangeProjectionStateActive     = "enterprise_workflow_authority_vala_ticket_change_projection_active"
	EnterpriseWorkflowAuthorityValATicketChangeProjectionStatePartial    = "enterprise_workflow_authority_vala_ticket_change_projection_partial"
	EnterpriseWorkflowAuthorityValATicketChangeProjectionStateIncomplete = "enterprise_workflow_authority_vala_ticket_change_projection_incomplete"

	EnterpriseWorkflowAuthorityValAReconciliationBaselineStateActive     = "enterprise_workflow_authority_vala_reconciliation_baseline_active"
	EnterpriseWorkflowAuthorityValAReconciliationBaselineStatePartial    = "enterprise_workflow_authority_vala_reconciliation_baseline_partial"
	EnterpriseWorkflowAuthorityValAReconciliationBaselineStateIncomplete = "enterprise_workflow_authority_vala_reconciliation_baseline_incomplete"

	EnterpriseWorkflowAuthorityValAIdempotentMutationStateActive     = "enterprise_workflow_authority_vala_idempotent_mutation_active"
	EnterpriseWorkflowAuthorityValAIdempotentMutationStatePartial    = "enterprise_workflow_authority_vala_idempotent_mutation_partial"
	EnterpriseWorkflowAuthorityValAIdempotentMutationStateIncomplete = "enterprise_workflow_authority_vala_idempotent_mutation_incomplete"

	EnterpriseWorkflowAuthorityValAStateIncomplete  = "enterprise_workflow_authority_vala_incomplete"
	EnterpriseWorkflowAuthorityValAStateSubstantial = "enterprise_workflow_authority_vala_substantially_ready"
	EnterpriseWorkflowAuthorityValAStateActive      = "enterprise_workflow_authority_vala_active"
)

type WorkflowEventOrchestrationBaseline struct {
	CurrentState      string   `json:"current_state"`
	CanonicalSource   string   `json:"canonical_source"`
	EventClasses      []string `json:"event_classes,omitempty"`
	ProjectionTargets []string `json:"projection_targets,omitempty"`
	SyncBackSources   []string `json:"sync_back_sources,omitempty"`
	ReplayRecovery    []string `json:"replay_recovery,omitempty"`
	DegradedMode      string   `json:"degraded_mode"`
	Limitations       []string `json:"limitations,omitempty"`
}

type WorkflowLifecycleConnectorBaseline struct {
	ConnectorSystem         string   `json:"connector_system"`
	CurrentState            string   `json:"current_state"`
	ObjectClasses           []string `json:"object_classes,omitempty"`
	CreateSupported         bool     `json:"create_supported"`
	UpdateSupported         bool     `json:"update_supported"`
	SyncBackSupported       bool     `json:"sync_back_supported"`
	EvidenceBundleField     string   `json:"evidence_bundle_field"`
	ClosureFeedbackField    string   `json:"closure_feedback_field"`
	ApprovalReflectionField string   `json:"approval_reflection_field"`
	ReplayRecoveryPath      []string `json:"replay_recovery_path,omitempty"`
	Limitations             []string `json:"limitations,omitempty"`
}

type WorkflowEvidenceBundleInjectionBaseline struct {
	ConnectorSystem         string   `json:"connector_system"`
	CurrentState            string   `json:"current_state"`
	SupportedRedactionTiers []string `json:"supported_redaction_tiers,omitempty"`
	DefaultOutboundTier     string   `json:"default_outbound_tier"`
	RequiredInjectedRefs    []string `json:"required_injected_refs,omitempty"`
	RemediationExpectations []string `json:"remediation_expectations,omitempty"`
	ClosureConditions       []string `json:"closure_conditions,omitempty"`
	ExceptionContextFields  []string `json:"exception_context_fields,omitempty"`
	CanonicalPermalinkField string   `json:"canonical_permalink_field"`
	Limitations             []string `json:"limitations,omitempty"`
}

type WorkflowTicketChangeProjectionBaseline struct {
	ConnectorSystem               string   `json:"connector_system"`
	CurrentState                  string   `json:"current_state"`
	ProjectionTargets             []string `json:"projection_targets,omitempty"`
	AdvisoryOnlyFields            []string `json:"advisory_only_fields,omitempty"`
	SyncBackEligibleFields        []string `json:"sync_back_eligible_fields,omitempty"`
	NeverOverwriteCanonicalFields []string `json:"never_overwrite_canonical_fields,omitempty"`
	ConflictSignals               []string `json:"conflict_signals,omitempty"`
	Limitations                   []string `json:"limitations,omitempty"`
}

type WorkflowReconciliationBaseline struct {
	ConnectorSystem     string   `json:"connector_system"`
	CurrentState        string   `json:"current_state"`
	ConflictPrecedence  string   `json:"conflict_precedence"`
	LastSyncedField     string   `json:"last_synced_field"`
	DriftSignals        []string `json:"drift_signals,omitempty"`
	ReplayRecoveryPath  []string `json:"replay_recovery_path,omitempty"`
	StaleExternalMarker string   `json:"stale_external_marker"`
	DegradedMode        string   `json:"degraded_mode"`
	Limitations         []string `json:"limitations,omitempty"`
}

type WorkflowIdempotentMutationBaseline struct {
	ConnectorSystem           string   `json:"connector_system"`
	CurrentState              string   `json:"current_state"`
	MutationKeys              []string `json:"mutation_keys,omitempty"`
	DuplicateSuppressionRules []string `json:"duplicate_suppression_rules,omitempty"`
	ReplayProtectionRules     []string `json:"replay_protection_rules,omitempty"`
	OutageBehavior            string   `json:"outage_behavior"`
	Limitations               []string `json:"limitations,omitempty"`
}

func EnterpriseWorkflowAuthorityValAEventOrchestration() WorkflowEventOrchestrationBaseline {
	model := WorkflowEventOrchestrationBaseline{
		CanonicalSource: "canonical_workflow_orchestrator",
		EventClasses: []string{
			"runtime_drift_detected",
			"remediation_declared",
			"approval_decision_recorded",
			"override_or_exception_requested",
			"external_connector_signal_received",
			"validation_result_recorded",
			"closure_or_reopen_recorded",
			"rollback_signal_recorded",
		},
		ProjectionTargets: []string{
			WorkflowAuthorityConnectorJira,
			WorkflowAuthorityConnectorServiceNow,
			WorkflowAuthorityConnectorGitHub,
		},
		SyncBackSources: []string{
			WorkflowAuthorityConnectorJira,
			WorkflowAuthorityConnectorServiceNow,
			WorkflowAuthorityConnectorGitHub,
		},
		ReplayRecovery: []string{
			"read_last_canonical_transition",
			"replay_unapplied_connector_signals",
			"reissue_idempotent_projection_mutations",
		},
		DegradedMode: "connector_outage_or_sync_failure_keeps_canonical_orchestrator_authoritative_and_marks_projection_recovery_required",
		Limitations: []string{
			"Val A defines the unified event orchestration baseline before later Point 3 waves add signed delegated authority, live ledger enforcement, and final workflow authority review.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValAEventOrchestrationState(model)
	return model
}

func EnterpriseWorkflowAuthorityValALifecycleConnectors() []WorkflowLifecycleConnectorBaseline {
	return []WorkflowLifecycleConnectorBaseline{
		{
			ConnectorSystem:         WorkflowAuthorityConnectorJira,
			CurrentState:            "lifecycle_connector_ready",
			ObjectClasses:           []string{"ticket", "comment", "assignee_projection"},
			CreateSupported:         true,
			UpdateSupported:         true,
			SyncBackSupported:       true,
			EvidenceBundleField:     "issue_description_or_attachment",
			ClosureFeedbackField:    "resolution_or_comment_feedback",
			ApprovalReflectionField: "approval_comment_or_transition_note",
			ReplayRecoveryPath:      []string{"load_last_synced_issue_state", "replay_idempotent_issue_update", "reinject_evidence_bundle_if_missing"},
			Limitations:             []string{"Jira remains a workflow projection target and cannot directly close canonical workflow state."},
		},
		{
			ConnectorSystem:         WorkflowAuthorityConnectorServiceNow,
			CurrentState:            "lifecycle_connector_ready",
			ObjectClasses:           []string{"change_request", "work_note", "approval_projection"},
			CreateSupported:         true,
			UpdateSupported:         true,
			SyncBackSupported:       true,
			EvidenceBundleField:     "change_description_or_attachment",
			ClosureFeedbackField:    "change_state_feedback",
			ApprovalReflectionField: "cab_or_approval_note",
			ReplayRecoveryPath:      []string{"load_last_synced_change_state", "replay_idempotent_change_update", "reinject_evidence_bundle_if_missing"},
			Limitations:             []string{"ServiceNow remains administrative projection unless matched by canonical validation evidence."},
		},
		{
			ConnectorSystem:         WorkflowAuthorityConnectorGitHub,
			CurrentState:            "lifecycle_connector_ready",
			ObjectClasses:           []string{"issue_or_pr_comment", "status_check_projection", "workflow_run_feedback"},
			CreateSupported:         true,
			UpdateSupported:         true,
			SyncBackSupported:       true,
			EvidenceBundleField:     "issue_or_pr_comment_body",
			ClosureFeedbackField:    "workflow_run_or_comment_feedback",
			ApprovalReflectionField: "deployment_or_comment_projection",
			ReplayRecoveryPath:      []string{"read_latest_issue_pr_state", "replay_webhook_delivery_log", "reapply_status_check_or_comment_projection"},
			Limitations:             []string{"GitHub status, merge, or close remains advisory unless canonical workflow validation agrees."},
		},
	}
}

func EnterpriseWorkflowAuthorityValAEvidenceBundleInjection() []WorkflowEvidenceBundleInjectionBaseline {
	return []WorkflowEvidenceBundleInjectionBaseline{
		{
			ConnectorSystem:         WorkflowAuthorityConnectorJira,
			CurrentState:            "evidence_bundle_injection_ready",
			SupportedRedactionTiers: []string{WorkflowAuthorityEvidenceTierInternalFull, WorkflowAuthorityEvidenceTierPartnerScoped, WorkflowAuthorityEvidenceTierExternalTicketSafe},
			DefaultOutboundTier:     WorkflowAuthorityEvidenceTierExternalTicketSafe,
			RequiredInjectedRefs:    []string{"evidence_bundle_ref", "runtime_evidence_ref", "policy_decision_summary", "canonical_permalink"},
			RemediationExpectations: []string{"owner_routed_remediation_expectations", "validation_required_before_canonical_close"},
			ClosureConditions:       []string{"validated_fixed_required_for_close", "external_close_is_projection_only"},
			ExceptionContextFields:  []string{"exception_context", "approval_context", "expiry_or_revocation_notice"},
			CanonicalPermalinkField: "canonical_workflow_permalink",
			Limitations:             []string{"Outbound ticket payload remains external_ticket_safe and does not expose internal_full evidence by default."},
		},
		{
			ConnectorSystem:         WorkflowAuthorityConnectorServiceNow,
			CurrentState:            "evidence_bundle_injection_ready",
			SupportedRedactionTiers: []string{WorkflowAuthorityEvidenceTierInternalFull, WorkflowAuthorityEvidenceTierPartnerScoped, WorkflowAuthorityEvidenceTierExternalTicketSafe},
			DefaultOutboundTier:     WorkflowAuthorityEvidenceTierExternalTicketSafe,
			RequiredInjectedRefs:    []string{"evidence_bundle_ref", "runtime_evidence_ref", "closure_conditions", "canonical_permalink"},
			RemediationExpectations: []string{"change_executor_expectations", "rollback_visibility_required"},
			ClosureConditions:       []string{"validation_result_required_before_canonical_close", "expired_exception_blocks_close"},
			ExceptionContextFields:  []string{"exception_context", "approval_context", "rollback_expectations"},
			CanonicalPermalinkField: "canonical_workflow_permalink",
			Limitations:             []string{"Change payload remains bounded and external-ticket-safe even when richer internal evidence exists."},
		},
		{
			ConnectorSystem:         WorkflowAuthorityConnectorGitHub,
			CurrentState:            "evidence_bundle_injection_ready",
			SupportedRedactionTiers: []string{WorkflowAuthorityEvidenceTierInternalFull, WorkflowAuthorityEvidenceTierPartnerScoped, WorkflowAuthorityEvidenceTierExternalTicketSafe},
			DefaultOutboundTier:     WorkflowAuthorityEvidenceTierExternalTicketSafe,
			RequiredInjectedRefs:    []string{"evidence_bundle_ref", "remediation_expectations", "closure_conditions", "canonical_permalink"},
			RemediationExpectations: []string{"workflow_run_or_pr_expectations", "validation_required_before_close"},
			ClosureConditions:       []string{"merge_or_close_is_not_canonical_close", "rollback_signal_may_reopen"},
			ExceptionContextFields:  []string{"exception_context", "approval_context", "blast_radius_note"},
			CanonicalPermalinkField: "canonical_workflow_permalink",
			Limitations:             []string{"GitHub projection remains bounded to external-ticket-safe disclosure and does not publish internal_full evidence by default."},
		},
	}
}

func EnterpriseWorkflowAuthorityValATicketChangeProjection() []WorkflowTicketChangeProjectionBaseline {
	return []WorkflowTicketChangeProjectionBaseline{
		{
			ConnectorSystem:               WorkflowAuthorityConnectorJira,
			CurrentState:                  "ticket_change_projection_ready",
			ProjectionTargets:             []string{"issue_summary", "issue_description", "comment", "assignee"},
			AdvisoryOnlyFields:            []string{"resolution_label", "admin_close_note"},
			SyncBackEligibleFields:        []string{"status", "comment", "assignee", "change_request_link"},
			NeverOverwriteCanonicalFields: []string{"canonical_state", "validated_fixed", "closed", "reopened"},
			ConflictSignals:               []string{"external_closed_without_validation", "status_drift", "stale_external_state"},
			Limitations:                   []string{"Jira resolution remains advisory and cannot overwrite canonical closure."},
		},
		{
			ConnectorSystem:               WorkflowAuthorityConnectorServiceNow,
			CurrentState:                  "ticket_change_projection_ready",
			ProjectionTargets:             []string{"change_summary", "work_note", "approval_projection", "rollback_note"},
			AdvisoryOnlyFields:            []string{"resolved_flag", "cab_note"},
			SyncBackEligibleFields:        []string{"change_state", "work_note", "approval_comment", "executor_identity"},
			NeverOverwriteCanonicalFields: []string{"canonical_state", "closure_pending", "closed", "reopened"},
			ConflictSignals:               []string{"external_resolved_without_validation", "change_window_expired", "executor_drift"},
			Limitations:                   []string{"ServiceNow close semantics remain administrative and bounded by canonical workflow validation."},
		},
		{
			ConnectorSystem:               WorkflowAuthorityConnectorGitHub,
			CurrentState:                  "ticket_change_projection_ready",
			ProjectionTargets:             []string{"issue_or_pr_comment", "status_check_context", "workflow_run_feedback"},
			AdvisoryOnlyFields:            []string{"issue_closed_state", "pr_merged"},
			SyncBackEligibleFields:        []string{"workflow_run_state", "comment", "deployment_result", "rollback_signal"},
			NeverOverwriteCanonicalFields: []string{"canonical_state", "validated_fixed", "closed", "exception_active"},
			ConflictSignals:               []string{"pr_merged_without_validation", "workflow_run_divergence", "rollback_signal_after_close"},
			Limitations:                   []string{"GitHub close, merge, and workflow completion remain projection signals instead of canonical truth."},
		},
	}
}

func EnterpriseWorkflowAuthorityValAReconciliationBaseline() []WorkflowReconciliationBaseline {
	return []WorkflowReconciliationBaseline{
		{
			ConnectorSystem:     WorkflowAuthorityConnectorJira,
			CurrentState:        "reconciliation_baseline_ready",
			ConflictPrecedence:  "canonical_state_precedes_external_close_labels",
			LastSyncedField:     "last_synced_at",
			DriftSignals:        []string{"external_closed_without_validation", "status_drift", "assignee_drift"},
			ReplayRecoveryPath:  []string{"read_last_synced_issue_state", "compare_connector_and_canonical_state", "replay_unapplied_projection_mutations"},
			StaleExternalMarker: "stale_external_issue_state",
			DegradedMode:        "connector_outage_keeps_canonical_state_local_and_marks_reconciliation_replay_required",
			Limitations:         []string{"Reconciliation keeps canonical state local and treats stale or conflicting Jira state as workflow drift."},
		},
		{
			ConnectorSystem:     WorkflowAuthorityConnectorServiceNow,
			CurrentState:        "reconciliation_baseline_ready",
			ConflictPrecedence:  "canonical_state_precedes_external_change_resolution",
			LastSyncedField:     "last_synced_at",
			DriftSignals:        []string{"external_resolved_without_validation", "change_window_expired", "executor_identity_drift"},
			ReplayRecoveryPath:  []string{"read_last_synced_change_state", "compare_external_and_canonical_state", "replay_idempotent_change_updates"},
			StaleExternalMarker: "stale_external_change_state",
			DegradedMode:        "rate_limit_or_outage_marks_projection_degraded_but_preserves_canonical_workflow_authority",
			Limitations:         []string{"Administrative closure remains advisory until canonical workflow validation agrees."},
		},
		{
			ConnectorSystem:     WorkflowAuthorityConnectorGitHub,
			CurrentState:        "reconciliation_baseline_ready",
			ConflictPrecedence:  "canonical_state_precedes_issue_close_merge_or_workflow_completion",
			LastSyncedField:     "last_delivery_at",
			DriftSignals:        []string{"pr_merged_without_validation", "duplicate_webhook_delivery", "rollback_signal_after_close"},
			ReplayRecoveryPath:  []string{"read_last_webhook_delivery_state", "deduplicate_delivery_id", "replay_projection_or_reopen_signal"},
			StaleExternalMarker: "stale_external_github_state",
			DegradedMode:        "webhook_or_actions_outage_preserves_canonical_state_and_marks_sync_back_review_required",
			Limitations:         []string{"GitHub workflow and merge feedback stay bounded to reconciliation and cannot directly close canonical state."},
		},
	}
}

func EnterpriseWorkflowAuthorityValAIdempotentMutationDiscipline() []WorkflowIdempotentMutationBaseline {
	return []WorkflowIdempotentMutationBaseline{
		{
			ConnectorSystem:           WorkflowAuthorityConnectorJira,
			CurrentState:              "idempotent_mutation_ready",
			MutationKeys:              []string{"workflow_id", "connector_ref", "request_idempotency_key"},
			DuplicateSuppressionRules: []string{"ignore_replayed_transition_with_same_idempotency_key", "ignore_duplicate_comment_injection_hash"},
			ReplayProtectionRules:     []string{"replayed_connector_write_must_not_expand_scope", "duplicate_transition_attempts_must_be_suppressed"},
			OutageBehavior:            "buffer_retry_or_manual_replay_without_mutating_canonical_state",
			Limitations:               []string{"Idempotent Jira mutation discipline defines keys and duplicate suppression before later Point 3 waves add live action authorization."},
		},
		{
			ConnectorSystem:           WorkflowAuthorityConnectorServiceNow,
			CurrentState:              "idempotent_mutation_ready",
			MutationKeys:              []string{"workflow_id", "connector_ref", "change_request_idempotency_key"},
			DuplicateSuppressionRules: []string{"ignore_replayed_state_write_with_same_key", "suppress_duplicate_evidence_bundle_upload"},
			ReplayProtectionRules:     []string{"replayed_change_update_must_not_silently_reopen_scope", "duplicate_approval_reflection_must_be_suppressed"},
			OutageBehavior:            "queue_or_manual_replay_without_letting_external_state_overwrite_canonical_state",
			Limitations:               []string{"Idempotent ServiceNow mutation discipline remains bounded and replayable without granting autonomous closure authority."},
		},
		{
			ConnectorSystem:           WorkflowAuthorityConnectorGitHub,
			CurrentState:              "idempotent_mutation_ready",
			MutationKeys:              []string{"workflow_id", "connector_ref", "github_delivery_id"},
			DuplicateSuppressionRules: []string{"ignore_duplicate_webhook_delivery", "suppress_duplicate_comment_body_hash"},
			ReplayProtectionRules:     []string{"replayed_status_check_or_comment_must_not_change_canonical_state", "duplicate_delivery_must_be_marked_replayed"},
			OutageBehavior:            "replay_from_delivery_log_or_manual_projection_without_losing_canonical_authority",
			Limitations:               []string{"Idempotent GitHub mutation discipline stays projection-scoped and cannot bypass canonical workflow validation."},
		},
	}
}

func EvaluateEnterpriseWorkflowAuthorityValAEventOrchestrationState(model WorkflowEventOrchestrationBaseline) string {
	if strings.TrimSpace(model.CanonicalSource) == "" || len(model.EventClasses) == 0 || len(model.ProjectionTargets) == 0 {
		return EnterpriseWorkflowAuthorityValAEventOrchestrationStateIncomplete
	}
	if len(model.SyncBackSources) == 0 || len(model.ReplayRecovery) == 0 || strings.TrimSpace(model.DegradedMode) == "" {
		return EnterpriseWorkflowAuthorityValAEventOrchestrationStatePartial
	}
	return EnterpriseWorkflowAuthorityValAEventOrchestrationStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValALifecycleConnectorsState(items []WorkflowLifecycleConnectorBaseline) string {
	if len(items) == 0 {
		return EnterpriseWorkflowAuthorityValALifecycleConnectorsStateIncomplete
	}
	expected := map[string]struct{}{
		WorkflowAuthorityConnectorJira:       {},
		WorkflowAuthorityConnectorServiceNow: {},
		WorkflowAuthorityConnectorGitHub:     {},
	}
	seen := make(map[string]struct{}, len(items))
	for _, item := range items {
		connector := strings.TrimSpace(item.ConnectorSystem)
		if connector == "" || strings.TrimSpace(item.CurrentState) == "" {
			return EnterpriseWorkflowAuthorityValALifecycleConnectorsStateIncomplete
		}
		if _, ok := expected[connector]; !ok {
			return EnterpriseWorkflowAuthorityValALifecycleConnectorsStatePartial
		}
		if _, duplicate := seen[connector]; duplicate {
			return EnterpriseWorkflowAuthorityValALifecycleConnectorsStatePartial
		}
		seen[connector] = struct{}{}
		if len(item.ObjectClasses) == 0 || !item.CreateSupported || !item.UpdateSupported || !item.SyncBackSupported || strings.TrimSpace(item.EvidenceBundleField) == "" || strings.TrimSpace(item.ClosureFeedbackField) == "" || strings.TrimSpace(item.ApprovalReflectionField) == "" || len(item.ReplayRecoveryPath) == 0 {
			return EnterpriseWorkflowAuthorityValALifecycleConnectorsStatePartial
		}
	}
	if len(seen) != len(expected) {
		return EnterpriseWorkflowAuthorityValALifecycleConnectorsStatePartial
	}
	return EnterpriseWorkflowAuthorityValALifecycleConnectorsStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValAEvidenceBundleInjectionState(items []WorkflowEvidenceBundleInjectionBaseline) string {
	if len(items) == 0 {
		return EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStateIncomplete
	}
	for _, item := range items {
		if strings.TrimSpace(item.ConnectorSystem) == "" || strings.TrimSpace(item.CurrentState) == "" {
			return EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStateIncomplete
		}
		if !containsTrimmedString(item.SupportedRedactionTiers, WorkflowAuthorityEvidenceTierInternalFull) || !containsTrimmedString(item.SupportedRedactionTiers, WorkflowAuthorityEvidenceTierPartnerScoped) || !containsTrimmedString(item.SupportedRedactionTiers, WorkflowAuthorityEvidenceTierExternalTicketSafe) || strings.TrimSpace(item.DefaultOutboundTier) == "" || len(item.RequiredInjectedRefs) == 0 || len(item.RemediationExpectations) == 0 || len(item.ClosureConditions) == 0 || len(item.ExceptionContextFields) == 0 || strings.TrimSpace(item.CanonicalPermalinkField) == "" {
			return EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStatePartial
		}
	}
	return EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValATicketChangeProjectionState(items []WorkflowTicketChangeProjectionBaseline) string {
	if len(items) == 0 {
		return EnterpriseWorkflowAuthorityValATicketChangeProjectionStateIncomplete
	}
	for _, item := range items {
		if strings.TrimSpace(item.ConnectorSystem) == "" || strings.TrimSpace(item.CurrentState) == "" {
			return EnterpriseWorkflowAuthorityValATicketChangeProjectionStateIncomplete
		}
		if len(item.ProjectionTargets) == 0 || len(item.AdvisoryOnlyFields) == 0 || len(item.SyncBackEligibleFields) == 0 || len(item.NeverOverwriteCanonicalFields) == 0 || len(item.ConflictSignals) == 0 {
			return EnterpriseWorkflowAuthorityValATicketChangeProjectionStatePartial
		}
	}
	return EnterpriseWorkflowAuthorityValATicketChangeProjectionStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValAReconciliationBaselineState(items []WorkflowReconciliationBaseline) string {
	if len(items) == 0 {
		return EnterpriseWorkflowAuthorityValAReconciliationBaselineStateIncomplete
	}
	for _, item := range items {
		if strings.TrimSpace(item.ConnectorSystem) == "" || strings.TrimSpace(item.CurrentState) == "" {
			return EnterpriseWorkflowAuthorityValAReconciliationBaselineStateIncomplete
		}
		if strings.TrimSpace(item.ConflictPrecedence) == "" || strings.TrimSpace(item.LastSyncedField) == "" || len(item.DriftSignals) == 0 || len(item.ReplayRecoveryPath) == 0 || strings.TrimSpace(item.StaleExternalMarker) == "" || strings.TrimSpace(item.DegradedMode) == "" {
			return EnterpriseWorkflowAuthorityValAReconciliationBaselineStatePartial
		}
	}
	return EnterpriseWorkflowAuthorityValAReconciliationBaselineStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValAIdempotentMutationState(items []WorkflowIdempotentMutationBaseline) string {
	if len(items) == 0 {
		return EnterpriseWorkflowAuthorityValAIdempotentMutationStateIncomplete
	}
	for _, item := range items {
		if strings.TrimSpace(item.ConnectorSystem) == "" || strings.TrimSpace(item.CurrentState) == "" {
			return EnterpriseWorkflowAuthorityValAIdempotentMutationStateIncomplete
		}
		if len(item.MutationKeys) == 0 || len(item.DuplicateSuppressionRules) == 0 || len(item.ReplayProtectionRules) == 0 || strings.TrimSpace(item.OutageBehavior) == "" {
			return EnterpriseWorkflowAuthorityValAIdempotentMutationStatePartial
		}
	}
	return EnterpriseWorkflowAuthorityValAIdempotentMutationStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValAState(val0State, eventOrchestrationState, lifecycleConnectorsState, evidenceBundleInjectionState, ticketChangeProjectionState, reconciliationBaselineState, idempotentMutationState string) string {
	if strings.TrimSpace(val0State) != EnterpriseWorkflowAuthorityVal0StateActive {
		return EnterpriseWorkflowAuthorityValAStateIncomplete
	}

	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(eventOrchestrationState),
		strings.TrimSpace(lifecycleConnectorsState),
		strings.TrimSpace(evidenceBundleInjectionState),
		strings.TrimSpace(ticketChangeProjectionState),
		strings.TrimSpace(reconciliationBaselineState),
		strings.TrimSpace(idempotentMutationState),
	} {
		switch state {
		case EnterpriseWorkflowAuthorityValAEventOrchestrationStateActive,
			EnterpriseWorkflowAuthorityValALifecycleConnectorsStateActive,
			EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStateActive,
			EnterpriseWorkflowAuthorityValATicketChangeProjectionStateActive,
			EnterpriseWorkflowAuthorityValAReconciliationBaselineStateActive,
			EnterpriseWorkflowAuthorityValAIdempotentMutationStateActive:
		case EnterpriseWorkflowAuthorityValAEventOrchestrationStatePartial,
			EnterpriseWorkflowAuthorityValALifecycleConnectorsStatePartial,
			EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStatePartial,
			EnterpriseWorkflowAuthorityValATicketChangeProjectionStatePartial,
			EnterpriseWorkflowAuthorityValAReconciliationBaselineStatePartial,
			EnterpriseWorkflowAuthorityValAIdempotentMutationStatePartial:
			hasPartial = true
		default:
			return EnterpriseWorkflowAuthorityValAStateIncomplete
		}
	}
	if hasPartial {
		return EnterpriseWorkflowAuthorityValAStateSubstantial
	}
	return EnterpriseWorkflowAuthorityValAStateActive
}

func containsTrimmedString(values []string, needle string) bool {
	needle = strings.TrimSpace(needle)
	if needle == "" {
		return false
	}
	for _, value := range values {
		if strings.TrimSpace(value) == needle {
			return true
		}
	}
	return false
}

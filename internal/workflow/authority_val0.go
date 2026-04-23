package workflow

import "strings"

const (
	EnterpriseWorkflowAuthorityVal0BoundaryStateActive     = "enterprise_workflow_authority_val0_boundary_active"
	EnterpriseWorkflowAuthorityVal0BoundaryStatePartial    = "enterprise_workflow_authority_val0_boundary_partial"
	EnterpriseWorkflowAuthorityVal0BoundaryStateIncomplete = "enterprise_workflow_authority_val0_boundary_incomplete"

	EnterpriseWorkflowAuthorityVal0StateMachineStateActive     = "enterprise_workflow_authority_val0_state_machine_active"
	EnterpriseWorkflowAuthorityVal0StateMachineStatePartial    = "enterprise_workflow_authority_val0_state_machine_partial"
	EnterpriseWorkflowAuthorityVal0StateMachineStateIncomplete = "enterprise_workflow_authority_val0_state_machine_incomplete"

	EnterpriseWorkflowAuthorityVal0ProjectionStateActive     = "enterprise_workflow_authority_val0_projection_rules_active"
	EnterpriseWorkflowAuthorityVal0ProjectionStatePartial    = "enterprise_workflow_authority_val0_projection_rules_partial"
	EnterpriseWorkflowAuthorityVal0ProjectionStateIncomplete = "enterprise_workflow_authority_val0_projection_rules_incomplete"

	EnterpriseWorkflowAuthorityVal0ApprovalContractStateActive     = "enterprise_workflow_authority_val0_approval_contract_active"
	EnterpriseWorkflowAuthorityVal0ApprovalContractStatePartial    = "enterprise_workflow_authority_val0_approval_contract_partial"
	EnterpriseWorkflowAuthorityVal0ApprovalContractStateIncomplete = "enterprise_workflow_authority_val0_approval_contract_incomplete"

	EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStateActive     = "enterprise_workflow_authority_val0_exception_lifecycle_active"
	EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStatePartial    = "enterprise_workflow_authority_val0_exception_lifecycle_partial"
	EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStateIncomplete = "enterprise_workflow_authority_val0_exception_lifecycle_incomplete"

	EnterpriseWorkflowAuthorityVal0ClosureValidationStateActive     = "enterprise_workflow_authority_val0_closure_validation_active"
	EnterpriseWorkflowAuthorityVal0ClosureValidationStatePartial    = "enterprise_workflow_authority_val0_closure_validation_partial"
	EnterpriseWorkflowAuthorityVal0ClosureValidationStateIncomplete = "enterprise_workflow_authority_val0_closure_validation_incomplete"

	EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStateActive     = "enterprise_workflow_authority_val0_separation_of_duties_active"
	EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStatePartial    = "enterprise_workflow_authority_val0_separation_of_duties_partial"
	EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStateIncomplete = "enterprise_workflow_authority_val0_separation_of_duties_incomplete"

	EnterpriseWorkflowAuthorityVal0TimeAuthorityStateActive     = "enterprise_workflow_authority_val0_time_authority_active"
	EnterpriseWorkflowAuthorityVal0TimeAuthorityStatePartial    = "enterprise_workflow_authority_val0_time_authority_partial"
	EnterpriseWorkflowAuthorityVal0TimeAuthorityStateIncomplete = "enterprise_workflow_authority_val0_time_authority_incomplete"

	EnterpriseWorkflowAuthorityVal0StateIncomplete  = "enterprise_workflow_authority_val0_incomplete"
	EnterpriseWorkflowAuthorityVal0StateSubstantial = "enterprise_workflow_authority_val0_substantially_ready"
	EnterpriseWorkflowAuthorityVal0StateActive      = "enterprise_workflow_authority_val0_active"

	WorkflowAuthorityActionAdvisoryOnly                = "advisory_only"
	WorkflowAuthorityActionApprovalRequired            = "approval_required"
	WorkflowAuthorityActionAutoExecutablePreapproved   = "auto_executable_under_preapproved_policy"
	WorkflowAuthorityActionExternalProjectionOnly      = "external_projection_only"
	WorkflowAuthorityActionCanonicalStateMutation      = "canonical_state_mutation"
	WorkflowAuthorityActionForbiddenWithoutHumanReview = "forbidden_without_human_review"
	WorkflowAuthorityConsumptionSingleUse              = "single_use"
	WorkflowAuthorityConsumptionMultiUseBounded        = "multi_use_bounded"
	WorkflowAuthorityConsumptionSessionBound           = "session_bound"
	WorkflowAuthorityConnectorJira                     = "jira"
	WorkflowAuthorityConnectorServiceNow               = "servicenow"
	WorkflowAuthorityConnectorGitHub                   = "github"
	WorkflowAuthorityStateDetected                     = "detected"
	WorkflowAuthorityStateTriaged                      = "triaged"
	WorkflowAuthorityStatePendingExternalAction        = "pending_external_action"
	WorkflowAuthorityStateApprovalRequired             = "approval_required"
	WorkflowAuthorityStateRemediationInProgress        = "remediation_in_progress"
	WorkflowAuthorityStateOverrideRequested            = "override_requested"
	WorkflowAuthorityStateOverrideActive               = "override_active"
	WorkflowAuthorityStateExceptionActive              = "exception_active"
	WorkflowAuthorityStateValidationPending            = "validation_pending"
	WorkflowAuthorityStateValidatedFixed               = "validated_fixed"
	WorkflowAuthorityStateClosurePending               = "closure_pending"
	WorkflowAuthorityStateClosed                       = "closed"
	WorkflowAuthorityStateReopened                     = "reopened"
	WorkflowAuthorityStateSuperseded                   = "superseded"
	WorkflowAuthorityEvidenceTierInternalFull          = "internal_full"
	WorkflowAuthorityEvidenceTierPartnerScoped         = "partner_scoped"
	WorkflowAuthorityEvidenceTierExternalTicketSafe    = "external_ticket_safe"
	WorkflowAuthoritySensitiveActionBreakGlass         = "break_glass"
	WorkflowAuthoritySensitiveActionLongLivedException = "long_lived_exception"
	WorkflowAuthoritySensitiveActionBroadScopeOverride = "broad_scope_override"
	WorkflowAuthoritySensitiveActionProductionClosure  = "production_closure_override"
	WorkflowAuthorityTimeSourceCanonicalService        = "canonical_service_time"
	WorkflowAuthorityTimeSourceExternalAdvisory        = "external_clock_advisory_only"
)

type AuthorityBoundaryRule struct {
	ActionClass               string   `json:"action_class"`
	CurrentState              string   `json:"current_state"`
	AuthorityMode             string   `json:"authority_mode"`
	HumanReviewPolicy         string   `json:"human_review_policy"`
	EvidenceRequired          bool     `json:"evidence_required"`
	IdentityBound             bool     `json:"identity_bound"`
	TimeBound                 bool     `json:"time_bound"`
	ExternalProjectionAllowed bool     `json:"external_projection_allowed"`
	Limitations               []string `json:"limitations,omitempty"`
}

type CanonicalWorkflowStateDefinition struct {
	State                  string   `json:"state"`
	CurrentState           string   `json:"current_state"`
	AllowedTransitions     []string `json:"allowed_transitions,omitempty"`
	RequiredEvidence       []string `json:"required_evidence,omitempty"`
	EntryConditions        []string `json:"entry_conditions,omitempty"`
	ExternalLabelSemantics []string `json:"external_label_semantics,omitempty"`
}

type CanonicalWorkflowStateMachine struct {
	CurrentState            string                             `json:"current_state"`
	States                  []CanonicalWorkflowStateDefinition `json:"states,omitempty"`
	GlobalInvariants        []string                           `json:"global_invariants,omitempty"`
	ConflictPrecedenceRules []string                           `json:"conflict_precedence_rules,omitempty"`
	ClosureAuthorityRules   []string                           `json:"closure_authority_rules,omitempty"`
	Limitations             []string                           `json:"limitations,omitempty"`
}

type ExternalProjectionRule struct {
	ConnectorSystem               string   `json:"connector_system"`
	CurrentState                  string   `json:"current_state"`
	ProjectionFields              []string `json:"projection_fields,omitempty"`
	SyncBackEligibleFields        []string `json:"sync_back_eligible_fields,omitempty"`
	AdvisoryOnlyFields            []string `json:"advisory_only_fields,omitempty"`
	ConflictSignals               []string `json:"conflict_signals,omitempty"`
	NeverOverwriteCanonicalFields []string `json:"never_overwrite_canonical_fields,omitempty"`
	IdempotentMutationKeys        []string `json:"idempotent_mutation_keys,omitempty"`
	DuplicateSuppressionRules     []string `json:"duplicate_suppression_rules,omitempty"`
	DegradedMode                  string   `json:"degraded_mode"`
	ReplayRecoveryPath            []string `json:"replay_recovery_path,omitempty"`
	EvidenceTier                  string   `json:"evidence_tier"`
	Limitations                   []string `json:"limitations,omitempty"`
}

type WorkflowApprovalActionContract struct {
	CurrentState            string   `json:"current_state"`
	RequiredArtifactFields  []string `json:"required_artifact_fields,omitempty"`
	ConsumptionSemantics    []string `json:"consumption_semantics,omitempty"`
	RevocationRules         []string `json:"revocation_rules,omitempty"`
	AntiReplayRules         []string `json:"anti_replay_rules,omitempty"`
	ScopeRules              []string `json:"scope_rules,omitempty"`
	ExternalSystemBindings  []string `json:"external_system_bindings,omitempty"`
	SeparationOfDutiesRules []string `json:"separation_of_duties_rules,omitempty"`
	Limitations             []string `json:"limitations,omitempty"`
}

type ExceptionLifecycleStage struct {
	Status           string   `json:"status"`
	CurrentState     string   `json:"current_state"`
	RequiresApproval bool     `json:"requires_approval"`
	RequiresEvidence bool     `json:"requires_evidence"`
	EnforcedAtExpiry bool     `json:"enforced_at_expiry"`
	ExpiryEffects    []string `json:"expiry_effects,omitempty"`
	VerifierVisible  bool     `json:"verifier_visible"`
	Limitations      []string `json:"limitations,omitempty"`
}

type ClosureValidationDiscipline struct {
	CurrentState             string   `json:"current_state"`
	RequiredChecks           []string `json:"required_checks,omitempty"`
	BoundedAlternativePolicy []string `json:"bounded_alternative_policy,omitempty"`
	ReopenRules              []string `json:"reopen_rules,omitempty"`
	RollbackRules            []string `json:"rollback_rules,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type SeparationOfDutiesRule struct {
	ActionClass              string   `json:"action_class"`
	CurrentState             string   `json:"current_state"`
	ControlMode              string   `json:"control_mode"`
	DistinctApproverExecutor bool     `json:"distinct_approver_executor"`
	DualControlRequired      bool     `json:"dual_control_required"`
	BlastRadiusConstraint    string   `json:"blast_radius_constraint"`
	RevocationPath           []string `json:"revocation_path,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type TimeAuthorityDiscipline struct {
	CurrentState             string   `json:"current_state"`
	CanonicalTimeSource      string   `json:"canonical_time_source"`
	ClockSkewTolerance       string   `json:"clock_skew_tolerance"`
	ExpiryEvaluationRules    []string `json:"expiry_evaluation_rules,omitempty"`
	FreshnessEvaluationRules []string `json:"freshness_evaluation_rules,omitempty"`
	TokenEvaluationRules     []string `json:"token_evaluation_rules,omitempty"`
	ConnectorTimeRules       []string `json:"connector_time_rules,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

func EnterpriseWorkflowAuthorityVal0BoundaryRules() []AuthorityBoundaryRule {
	items := []AuthorityBoundaryRule{
		{
			ActionClass:               WorkflowAuthorityActionAdvisoryOnly,
			CurrentState:              "authority_boundary_ready",
			AuthorityMode:             "projection_only_no_canonical_mutation",
			HumanReviewPolicy:         "not_required_for_advisory_projection",
			EvidenceRequired:          true,
			IdentityBound:             true,
			TimeBound:                 false,
			ExternalProjectionAllowed: true,
			Limitations:               []string{"Advisory projection may explain or annotate workflow state but cannot close or approve it."},
		},
		{
			ActionClass:               WorkflowAuthorityActionApprovalRequired,
			CurrentState:              "authority_boundary_ready",
			AuthorityMode:             "signed_human_approval_required",
			HumanReviewPolicy:         "required_before_effective_authority",
			EvidenceRequired:          true,
			IdentityBound:             true,
			TimeBound:                 true,
			ExternalProjectionAllowed: true,
			Limitations:               []string{"Approval-required actions stay blocked until a valid authorization artifact is issued and not revoked."},
		},
		{
			ActionClass:               WorkflowAuthorityActionAutoExecutablePreapproved,
			CurrentState:              "authority_boundary_ready",
			AuthorityMode:             "preapproved_policy_gate",
			HumanReviewPolicy:         "bounded_to_preapproved_policy",
			EvidenceRequired:          true,
			IdentityBound:             true,
			TimeBound:                 true,
			ExternalProjectionAllowed: true,
			Limitations:               []string{"Auto-executable mutations remain bounded to explicit policy, scope, and expiry conditions."},
		},
		{
			ActionClass:               WorkflowAuthorityActionExternalProjectionOnly,
			CurrentState:              "authority_boundary_ready",
			AuthorityMode:             "external_system_projection",
			HumanReviewPolicy:         "not_authoritative_for_canonical_state",
			EvidenceRequired:          true,
			IdentityBound:             true,
			TimeBound:                 false,
			ExternalProjectionAllowed: true,
			Limitations:               []string{"External projection state remains advisory and never directly overwrites canonical closure or validation truth."},
		},
		{
			ActionClass:               WorkflowAuthorityActionCanonicalStateMutation,
			CurrentState:              "authority_boundary_ready",
			AuthorityMode:             "canonical_mutation_with_evidence",
			HumanReviewPolicy:         "approval_or_validation_bound",
			EvidenceRequired:          true,
			IdentityBound:             true,
			TimeBound:                 true,
			ExternalProjectionAllowed: true,
			Limitations:               []string{"Canonical mutation is only valid when evidence, identity, and policy conditions are satisfied."},
		},
		{
			ActionClass:               WorkflowAuthorityActionForbiddenWithoutHumanReview,
			CurrentState:              "authority_boundary_ready",
			AuthorityMode:             "forbidden_without_human_review",
			HumanReviewPolicy:         "mandatory_dual_or_distinct_control",
			EvidenceRequired:          true,
			IdentityBound:             true,
			TimeBound:                 true,
			ExternalProjectionAllowed: false,
			Limitations:               []string{"Broad-scope overrides and closure overrides cannot self-authorize from external ticket state or automated projection alone."},
		},
	}
	return items
}

func EnterpriseWorkflowAuthorityVal0StateMachine() CanonicalWorkflowStateMachine {
	model := CanonicalWorkflowStateMachine{
		States: []CanonicalWorkflowStateDefinition{
			{State: WorkflowAuthorityStateDetected, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateTriaged}, RequiredEvidence: []string{"detection_evidence_ref"}, EntryConditions: []string{"subject_identified"}, ExternalLabelSemantics: []string{"external_open_labels_are_projection_only"}},
			{State: WorkflowAuthorityStateTriaged, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStatePendingExternalAction, WorkflowAuthorityStateApprovalRequired, WorkflowAuthorityStateRemediationInProgress}, RequiredEvidence: []string{"triage_summary_ref"}, EntryConditions: []string{"owner_or_route_assigned"}, ExternalLabelSemantics: []string{"external_status_may_reflect_assignment_only"}},
			{State: WorkflowAuthorityStatePendingExternalAction, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateRemediationInProgress, WorkflowAuthorityStateApprovalRequired}, RequiredEvidence: []string{"connector_projection_ref"}, EntryConditions: []string{"external_projection_created"}, ExternalLabelSemantics: []string{"ticket_state_is_not_closure_authority"}},
			{State: WorkflowAuthorityStateApprovalRequired, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateRemediationInProgress, WorkflowAuthorityStateOverrideRequested}, RequiredEvidence: []string{"approval_request_ref"}, EntryConditions: []string{"action_exceeds_auto_policy"}, ExternalLabelSemantics: []string{"external_approval_comments_are_advisory_without_canonical_artifact"}},
			{State: WorkflowAuthorityStateRemediationInProgress, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateValidationPending, WorkflowAuthorityStateOverrideRequested}, RequiredEvidence: []string{"remediation_plan_ref"}, EntryConditions: []string{"approved_or_preapproved_action"}, ExternalLabelSemantics: []string{"external_work_in_progress_labels_do_not_change_validation_requirements"}},
			{State: WorkflowAuthorityStateOverrideRequested, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateOverrideActive, WorkflowAuthorityStateRemediationInProgress}, RequiredEvidence: []string{"override_request_ref"}, EntryConditions: []string{"bounded_exception_or_break_glass_requested"}, ExternalLabelSemantics: []string{"external_request_labels_do_not_activate_override"}},
			{State: WorkflowAuthorityStateOverrideActive, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateValidationPending, WorkflowAuthorityStateExceptionActive}, RequiredEvidence: []string{"authorization_artifact_ref"}, EntryConditions: []string{"signed_authorization_active"}, ExternalLabelSemantics: []string{"external_ticket_close_cannot_cancel_override_without_canonical_revoke"}},
			{State: WorkflowAuthorityStateExceptionActive, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateValidationPending, WorkflowAuthorityStateReopened}, RequiredEvidence: []string{"exception_registry_ref"}, EntryConditions: []string{"exception_approved_and_not_expired"}, ExternalLabelSemantics: []string{"external_resolution_does_not_end_exception_by_itself"}},
			{State: WorkflowAuthorityStateValidationPending, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateValidatedFixed, WorkflowAuthorityStateReopened}, RequiredEvidence: []string{"validation_harness_ref"}, EntryConditions: []string{"remediation_declared"}, ExternalLabelSemantics: []string{"external_closed_requires_validation_match"}},
			{State: WorkflowAuthorityStateValidatedFixed, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateClosurePending, WorkflowAuthorityStateReopened}, RequiredEvidence: []string{"validation_result_ref"}, EntryConditions: []string{"validation_result_verified"}, ExternalLabelSemantics: []string{"administrative_resolution_is_still_projection_only"}},
			{State: WorkflowAuthorityStateClosurePending, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateClosed, WorkflowAuthorityStateReopened}, RequiredEvidence: []string{"closure_policy_ref"}, EntryConditions: []string{"freshness_and_exception_checks_satisfied"}, ExternalLabelSemantics: []string{"external_close_label_is_not_sufficient"}},
			{State: WorkflowAuthorityStateClosed, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateReopened, WorkflowAuthorityStateSuperseded}, RequiredEvidence: []string{"closure_evidence_bundle_ref"}, EntryConditions: []string{"closure_validated"}, ExternalLabelSemantics: []string{"later_conflict_or_failure_may_reopen"}},
			{State: WorkflowAuthorityStateReopened, CurrentState: "workflow_state_defined", AllowedTransitions: []string{WorkflowAuthorityStateTriaged, WorkflowAuthorityStateValidationPending}, RequiredEvidence: []string{"reopen_reason_ref", "reopen_evidence_ref"}, EntryConditions: []string{"reopen_reason_present"}, ExternalLabelSemantics: []string{"canonical_reopen_overrides_external_closed_labels"}},
			{State: WorkflowAuthorityStateSuperseded, CurrentState: "workflow_state_defined", AllowedTransitions: nil, RequiredEvidence: []string{"replacement_object_ref"}, EntryConditions: []string{"replacement_link_present"}, ExternalLabelSemantics: []string{"historical_visibility_preserved"}},
		},
		GlobalInvariants: []string{
			"closed_requires_validated_fixed_or_bounded_alternative_policy",
			"override_active_requires_valid_authorization_artifact",
			"exception_active_forbidden_after_expires_at",
			"reopened_requires_reason_and_evidence_ref",
			"superseded_requires_replacement_link",
		},
		ConflictPrecedenceRules: []string{
			"canonical_state_precedes_external_close_labels",
			"external_assignee_or_status_drift_is_conflict_signal_not_canonical_truth",
			"stale_external_objects_cannot_overwrite_validation_or_closure_state",
		},
		ClosureAuthorityRules: []string{
			"resolved_in_jira_or_servicenow_is_not_equivalent_to_validated_fixed",
			"closure_requires_validation_result_or_explicit_bounded_alternative_policy",
			"reopen_remains_evidence_bound_and_audited",
		},
		Limitations: []string{
			"Val 0 defines the canonical workflow authority state machine and invariants before later Point 3 waves attach live orchestration and signed approval artifacts.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityVal0StateMachineState(model)
	return model
}

func EnterpriseWorkflowAuthorityVal0ExternalProjectionRules() []ExternalProjectionRule {
	items := []ExternalProjectionRule{
		{
			ConnectorSystem:               WorkflowAuthorityConnectorJira,
			CurrentState:                  "external_projection_rules_ready",
			ProjectionFields:              []string{"issue_summary", "issue_description", "evidence_bundle_ref", "closure_conditions", "exception_context"},
			SyncBackEligibleFields:        []string{"status", "comment", "change_request_link", "assignee"},
			AdvisoryOnlyFields:            []string{"resolution_label", "admin_close_note"},
			ConflictSignals:               []string{"external_closed_without_validation", "external_status_drift", "stale_external_state"},
			NeverOverwriteCanonicalFields: []string{"canonical_state", "validated_fixed", "closed", "reopened"},
			IdempotentMutationKeys:        []string{"workflow_id", "connector_ref", "request_idempotency_key"},
			DuplicateSuppressionRules:     []string{"ignore_replayed_transition_with_same_idempotency_key", "ignore_duplicate_comment_injection_hash"},
			DegradedMode:                  "connector_outage_keeps_canonical_state_active_and_marks_external_projection_degraded",
			ReplayRecoveryPath:            []string{"read_last_synced_state", "replay_canonical_events", "reapply_projection_if_idempotent"},
			EvidenceTier:                  WorkflowAuthorityEvidenceTierExternalTicketSafe,
			Limitations:                   []string{"Jira remains a projection target and cannot by itself close the canonical workflow."},
		},
		{
			ConnectorSystem:               WorkflowAuthorityConnectorServiceNow,
			CurrentState:                  "external_projection_rules_ready",
			ProjectionFields:              []string{"change_ticket_summary", "approval_context", "evidence_bundle_ref", "rollback_expectations"},
			SyncBackEligibleFields:        []string{"change_state", "approval_comment", "work_note", "executor_identity"},
			AdvisoryOnlyFields:            []string{"resolved_flag", "cab_note"},
			ConflictSignals:               []string{"external_resolved_without_validation", "change_window_expired", "external_executor_drift"},
			NeverOverwriteCanonicalFields: []string{"canonical_state", "closure_pending", "closed", "reopened"},
			IdempotentMutationKeys:        []string{"workflow_id", "connector_ref", "change_request_idempotency_key"},
			DuplicateSuppressionRules:     []string{"ignore_replayed_state_write_with_same_key", "suppress_duplicate_evidence_bundle_upload"},
			DegradedMode:                  "connector_rate_limit_or_outage_preserves_canonical_authority_and_marks_projection_replay_required",
			ReplayRecoveryPath:            []string{"reload_last_synced_change_state", "compare_external_and_canonical_timestamps", "replay_outstanding_idempotent_mutations"},
			EvidenceTier:                  WorkflowAuthorityEvidenceTierExternalTicketSafe,
			Limitations:                   []string{"ServiceNow closure remains administrative unless matched by canonical validation evidence."},
		},
		{
			ConnectorSystem:               WorkflowAuthorityConnectorGitHub,
			CurrentState:                  "external_projection_rules_ready",
			ProjectionFields:              []string{"issue_or_pr_comment", "status_check_context", "evidence_bundle_ref", "remediation_expectations"},
			SyncBackEligibleFields:        []string{"workflow_run_state", "comment", "deployment_result", "rollback_signal"},
			AdvisoryOnlyFields:            []string{"issue_closed_state", "pr_merged"},
			ConflictSignals:               []string{"pr_merged_without_validation", "workflow_run_divergence", "rollback_signal_after_close"},
			NeverOverwriteCanonicalFields: []string{"canonical_state", "validated_fixed", "closed", "exception_active"},
			IdempotentMutationKeys:        []string{"workflow_id", "connector_ref", "github_delivery_id"},
			DuplicateSuppressionRules:     []string{"ignore_duplicate_webhook_delivery", "suppress_duplicate_comment_body_hash"},
			DegradedMode:                  "webhook_or_actions_outage_keeps_canonical_state_authoritative_and flags_sync_back_review_required",
			ReplayRecoveryPath:            []string{"replay_webhook_delivery_log", "re-read_latest_issue_pr_state", "reconcile_comment_and_status_check_projection"},
			EvidenceTier:                  WorkflowAuthorityEvidenceTierExternalTicketSafe,
			Limitations:                   []string{"GitHub close, merge, or workflow completion cannot directly set canonical closure without validation evidence."},
		},
	}
	return items
}

func EnterpriseWorkflowAuthorityVal0ApprovalContract() WorkflowApprovalActionContract {
	model := WorkflowApprovalActionContract{
		RequiredArtifactFields: []string{
			"authorization_id",
			"subject",
			"actor_identity",
			"action_class",
			"scope",
			"issued_at",
			"expires_at",
			"justification",
			"evidence_refs",
			"signature_seal",
			"revocation_status",
			"consumption_mode",
			"anti_replay_marker_jti",
		},
		ConsumptionSemantics: []string{
			WorkflowAuthorityConsumptionSingleUse,
			WorkflowAuthorityConsumptionMultiUseBounded,
			WorkflowAuthorityConsumptionSessionBound,
		},
		RevocationRules: []string{
			"authorization_artifact_must_be_revocable_before_or_after_consumption_when_policy_requires",
			"revoked_artifact_cannot_activate_override_or_closure",
			"consumed_single_use_artifact_cannot_be_replayed",
		},
		AntiReplayRules: []string{
			"authorization_jti_or_nonce_must_be_unique_per_artifact",
			"duplicate_consumption_attempts_must_be_rejected_or_marked_replayed",
			"expired_artifacts_are_invalid_after_canonical_time_plus_allowed_skew",
		},
		ScopeRules: []string{
			"approval_must_be_subject_bound_scope_bound_and_time_bound",
			"approval_to_action_contract_must_identify_target_system_and_blast_radius",
			"closure_or_override_scope_cannot_expand_silently_after_issue",
		},
		ExternalSystemBindings: []string{
			"jira_transition_authorization",
			"servicenow_change_authorization",
			"github_deploy_or_comment_projection_authorization",
		},
		SeparationOfDutiesRules: []string{
			"break_glass_requires_dual_control_or_distinct_approver_executor",
			"production_closure_override_requires_separate_approver_and_executor",
			"long_lived_exception_requires_elevated_scope_review",
		},
		Limitations: []string{
			"Val 0 approval contract defines the required shape and anti-replay discipline before later Point 3 waves issue signed authorization artifacts.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityVal0ApprovalContractState(model)
	return model
}

func EnterpriseWorkflowAuthorityVal0ExceptionLifecycle() []ExceptionLifecycleStage {
	items := []ExceptionLifecycleStage{
		{Status: "requested", CurrentState: "exception_lifecycle_ready", RequiresApproval: true, RequiresEvidence: true, EnforcedAtExpiry: false, ExpiryEffects: []string{"await_approval_before_activation"}, VerifierVisible: true, Limitations: []string{"Request does not imply activation."}},
		{Status: "approved", CurrentState: "exception_lifecycle_ready", RequiresApproval: true, RequiresEvidence: true, EnforcedAtExpiry: true, ExpiryEffects: []string{"activate_before_expiry_or_reissue"}, VerifierVisible: true},
		{Status: "activated", CurrentState: "exception_lifecycle_ready", RequiresApproval: true, RequiresEvidence: true, EnforcedAtExpiry: true, ExpiryEffects: []string{"block_again", "review_required"}, VerifierVisible: true},
		{Status: "expiring", CurrentState: "exception_lifecycle_ready", RequiresApproval: false, RequiresEvidence: true, EnforcedAtExpiry: true, ExpiryEffects: []string{"grace_denied_without_revalidation"}, VerifierVisible: true},
		{Status: "expired", CurrentState: "exception_lifecycle_ready", RequiresApproval: false, RequiresEvidence: true, EnforcedAtExpiry: true, ExpiryEffects: []string{"block_again", "reopen_if_still_unresolved"}, VerifierVisible: true},
		{Status: "revoked", CurrentState: "exception_lifecycle_ready", RequiresApproval: false, RequiresEvidence: true, EnforcedAtExpiry: true, ExpiryEffects: []string{"authorization_invalidated", "projection_update_required"}, VerifierVisible: true},
		{Status: "superseded", CurrentState: "exception_lifecycle_ready", RequiresApproval: false, RequiresEvidence: true, EnforcedAtExpiry: false, ExpiryEffects: []string{"replacement_exception_link_required"}, VerifierVisible: true},
		{Status: "revalidated", CurrentState: "exception_lifecycle_ready", RequiresApproval: true, RequiresEvidence: true, EnforcedAtExpiry: true, ExpiryEffects: []string{"fresh_expiry_window_required"}, VerifierVisible: true},
	}
	return items
}

func EnterpriseWorkflowAuthorityVal0ClosureValidation() ClosureValidationDiscipline {
	model := ClosureValidationDiscipline{
		RequiredChecks: []string{
			"remediation_declared",
			"remediation_evidence_present",
			"validation_result_verified",
			"freshness_check_passed",
			"no_still_active_expired_exception",
			"closure_policy_satisfied",
		},
		BoundedAlternativePolicy: []string{
			"bounded_closure_override_must_reference_signed_authorization_and_policy_exception",
		},
		ReopenRules: []string{
			"reopen_requires_reason_and_evidence_ref",
			"later_conflict_or_failed_validation_reopens_canonical_workflow",
			"reopen_must_remain_audited_and_external_close_does_not_block_it",
		},
		RollbackRules: []string{
			"rollback_or_undo_must_link_to_canonical_workflow_and_evidence_bundle",
			"rollback_completion_does_not_imply_closure_without_validation",
			"rollback_path_must_be_visible_for_override_or_deployment_authority",
		},
		Limitations: []string{
			"Val 0 defines closure and reopen discipline before later Point 3 waves attach live validation harness and workflow ledger enforcement.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityVal0ClosureValidationState(model)
	return model
}

func EnterpriseWorkflowAuthorityVal0SeparationOfDuties() []SeparationOfDutiesRule {
	items := []SeparationOfDutiesRule{
		{
			ActionClass:              WorkflowAuthoritySensitiveActionBreakGlass,
			CurrentState:             "separation_of_duties_ready",
			ControlMode:              "dual_control",
			DistinctApproverExecutor: true,
			DualControlRequired:      true,
			BlastRadiusConstraint:    "subject_or_namespace_scoped",
			RevocationPath:           []string{"revoke_break_glass_artifact", "reopen_canonical_workflow_if_needed"},
			Limitations:              []string{"Break-glass remains bounded, expirable, and auditable."},
		},
		{
			ActionClass:              WorkflowAuthoritySensitiveActionLongLivedException,
			CurrentState:             "separation_of_duties_ready",
			ControlMode:              "distinct_approver_executor",
			DistinctApproverExecutor: true,
			DualControlRequired:      false,
			BlastRadiusConstraint:    "time_bound_and_policy_linked",
			RevocationPath:           []string{"revoke_exception", "enforce_auto_expiry"},
		},
		{
			ActionClass:              WorkflowAuthoritySensitiveActionBroadScopeOverride,
			CurrentState:             "separation_of_duties_ready",
			ControlMode:              "dual_control",
			DistinctApproverExecutor: true,
			DualControlRequired:      true,
			BlastRadiusConstraint:    "blast_radius_visible_and_bounded",
			RevocationPath:           []string{"revoke_override_artifact", "force_revalidation"},
		},
		{
			ActionClass:              WorkflowAuthoritySensitiveActionProductionClosure,
			CurrentState:             "separation_of_duties_ready",
			ControlMode:              "distinct_approver_executor",
			DistinctApproverExecutor: true,
			DualControlRequired:      true,
			BlastRadiusConstraint:    "production_scope_requires_dual_control",
			RevocationPath:           []string{"reopen_workflow", "withdraw_closure_authority"},
		},
	}
	return items
}

func EnterpriseWorkflowAuthorityVal0TimeAuthority() TimeAuthorityDiscipline {
	model := TimeAuthorityDiscipline{
		CanonicalTimeSource: WorkflowAuthorityTimeSourceCanonicalService,
		ClockSkewTolerance:  "90s",
		ExpiryEvaluationRules: []string{
			"evaluate_authorization_expiry_against_canonical_service_time",
			"apply_clock_skew_tolerance_before_rejecting_recently_issued_tokens",
			"exception_active_forbidden_after_expiry_plus_skew",
		},
		FreshnessEvaluationRules: []string{
			"validation_and_closure_freshness_use_canonical_service_time",
			"stale_external_connector_timestamps_remain_advisory_only",
			"reopen_on_later_evidence_uses_canonical_receive_time",
		},
		TokenEvaluationRules: []string{
			"authorization_artifacts_require_issued_at_and_expires_at",
			"expired_or_revoked_artifacts_cannot_be_consumed",
			"single_use_artifacts_cannot_be_reused_after_first_consumption",
		},
		ConnectorTimeRules: []string{
			WorkflowAuthorityTimeSourceExternalAdvisory,
			"external_timestamps_may_confirm_or_explain_but_not_override_canonical_expiry_evaluation",
			"connector_outage_or_clock_drift_marks_external_state_degraded_not_authoritative",
		},
		Limitations: []string{
			"Val 0 time authority defines canonical expiry and skew discipline before later Point 3 waves attach live authorization consumption and connector replay enforcement.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityVal0TimeAuthorityState(model)
	return model
}

func EvaluateEnterpriseWorkflowAuthorityVal0BoundaryState(items []AuthorityBoundaryRule) string {
	if len(items) == 0 {
		return EnterpriseWorkflowAuthorityVal0BoundaryStateIncomplete
	}
	expected := map[string]struct{}{
		WorkflowAuthorityActionAdvisoryOnly:                {},
		WorkflowAuthorityActionApprovalRequired:            {},
		WorkflowAuthorityActionAutoExecutablePreapproved:   {},
		WorkflowAuthorityActionExternalProjectionOnly:      {},
		WorkflowAuthorityActionCanonicalStateMutation:      {},
		WorkflowAuthorityActionForbiddenWithoutHumanReview: {},
	}
	for _, item := range items {
		action := strings.TrimSpace(item.ActionClass)
		if action == "" || strings.TrimSpace(item.CurrentState) != "authority_boundary_ready" || strings.TrimSpace(item.AuthorityMode) == "" || strings.TrimSpace(item.HumanReviewPolicy) == "" {
			return EnterpriseWorkflowAuthorityVal0BoundaryStatePartial
		}
		if !item.EvidenceRequired || !item.IdentityBound {
			return EnterpriseWorkflowAuthorityVal0BoundaryStatePartial
		}
		delete(expected, action)
	}
	if len(expected) != 0 {
		return EnterpriseWorkflowAuthorityVal0BoundaryStatePartial
	}
	return EnterpriseWorkflowAuthorityVal0BoundaryStateActive
}

func EvaluateEnterpriseWorkflowAuthorityVal0StateMachineState(model CanonicalWorkflowStateMachine) string {
	if len(model.States) == 0 || len(model.GlobalInvariants) == 0 || len(model.ConflictPrecedenceRules) == 0 || len(model.ClosureAuthorityRules) == 0 {
		return EnterpriseWorkflowAuthorityVal0StateMachineStateIncomplete
	}
	expectedStates := map[string]struct{}{
		WorkflowAuthorityStateDetected:              {},
		WorkflowAuthorityStateTriaged:               {},
		WorkflowAuthorityStatePendingExternalAction: {},
		WorkflowAuthorityStateApprovalRequired:      {},
		WorkflowAuthorityStateRemediationInProgress: {},
		WorkflowAuthorityStateOverrideRequested:     {},
		WorkflowAuthorityStateOverrideActive:        {},
		WorkflowAuthorityStateExceptionActive:       {},
		WorkflowAuthorityStateValidationPending:     {},
		WorkflowAuthorityStateValidatedFixed:        {},
		WorkflowAuthorityStateClosurePending:        {},
		WorkflowAuthorityStateClosed:                {},
		WorkflowAuthorityStateReopened:              {},
		WorkflowAuthorityStateSuperseded:            {},
	}
	for _, item := range model.States {
		state := strings.TrimSpace(item.State)
		if state == "" || strings.TrimSpace(item.CurrentState) != "workflow_state_defined" || len(item.AllowedTransitions) == 0 && state != WorkflowAuthorityStateSuperseded || len(item.RequiredEvidence) == 0 || len(item.EntryConditions) == 0 || len(item.ExternalLabelSemantics) == 0 {
			return EnterpriseWorkflowAuthorityVal0StateMachineStatePartial
		}
		delete(expectedStates, state)
	}
	if len(expectedStates) != 0 {
		return EnterpriseWorkflowAuthorityVal0StateMachineStatePartial
	}
	requiredInvariants := []string{
		"closed_requires_validated_fixed_or_bounded_alternative_policy",
		"override_active_requires_valid_authorization_artifact",
		"exception_active_forbidden_after_expires_at",
		"reopened_requires_reason_and_evidence_ref",
		"superseded_requires_replacement_link",
	}
	for _, invariant := range requiredInvariants {
		if !containsWorkflowValue(model.GlobalInvariants, invariant) {
			return EnterpriseWorkflowAuthorityVal0StateMachineStatePartial
		}
	}
	if !containsWorkflowValue(model.ConflictPrecedenceRules, "canonical_state_precedes_external_close_labels") {
		return EnterpriseWorkflowAuthorityVal0StateMachineStatePartial
	}
	return EnterpriseWorkflowAuthorityVal0StateMachineStateActive
}

func EvaluateEnterpriseWorkflowAuthorityVal0ProjectionState(items []ExternalProjectionRule) string {
	if len(items) == 0 {
		return EnterpriseWorkflowAuthorityVal0ProjectionStateIncomplete
	}
	expected := map[string]struct{}{
		WorkflowAuthorityConnectorJira:       {},
		WorkflowAuthorityConnectorServiceNow: {},
		WorkflowAuthorityConnectorGitHub:     {},
	}
	for _, item := range items {
		system := strings.TrimSpace(item.ConnectorSystem)
		if system == "" || strings.TrimSpace(item.CurrentState) != "external_projection_rules_ready" || len(item.ProjectionFields) == 0 || len(item.SyncBackEligibleFields) == 0 || len(item.AdvisoryOnlyFields) == 0 || len(item.ConflictSignals) == 0 || len(item.NeverOverwriteCanonicalFields) == 0 || len(item.IdempotentMutationKeys) == 0 || len(item.DuplicateSuppressionRules) == 0 || strings.TrimSpace(item.DegradedMode) == "" || len(item.ReplayRecoveryPath) == 0 || strings.TrimSpace(item.EvidenceTier) == "" {
			return EnterpriseWorkflowAuthorityVal0ProjectionStatePartial
		}
		delete(expected, system)
	}
	if len(expected) != 0 {
		return EnterpriseWorkflowAuthorityVal0ProjectionStatePartial
	}
	return EnterpriseWorkflowAuthorityVal0ProjectionStateActive
}

func EvaluateEnterpriseWorkflowAuthorityVal0ApprovalContractState(model WorkflowApprovalActionContract) string {
	if len(model.RequiredArtifactFields) == 0 || len(model.ConsumptionSemantics) == 0 {
		return EnterpriseWorkflowAuthorityVal0ApprovalContractStateIncomplete
	}
	requiredFields := []string{
		"authorization_id",
		"subject",
		"actor_identity",
		"action_class",
		"scope",
		"issued_at",
		"expires_at",
		"justification",
		"evidence_refs",
		"signature_seal",
		"revocation_status",
		"consumption_mode",
		"anti_replay_marker_jti",
	}
	for _, field := range requiredFields {
		if !containsWorkflowValue(model.RequiredArtifactFields, field) {
			return EnterpriseWorkflowAuthorityVal0ApprovalContractStatePartial
		}
	}
	for _, mode := range []string{
		WorkflowAuthorityConsumptionSingleUse,
		WorkflowAuthorityConsumptionMultiUseBounded,
		WorkflowAuthorityConsumptionSessionBound,
	} {
		if !containsWorkflowValue(model.ConsumptionSemantics, mode) {
			return EnterpriseWorkflowAuthorityVal0ApprovalContractStatePartial
		}
	}
	if len(model.RevocationRules) == 0 || len(model.AntiReplayRules) == 0 || len(model.ScopeRules) == 0 || len(model.ExternalSystemBindings) == 0 || len(model.SeparationOfDutiesRules) == 0 {
		return EnterpriseWorkflowAuthorityVal0ApprovalContractStatePartial
	}
	model.CurrentState = EnterpriseWorkflowAuthorityVal0ApprovalContractStateActive
	return EnterpriseWorkflowAuthorityVal0ApprovalContractStateActive
}

func EvaluateEnterpriseWorkflowAuthorityVal0ExceptionLifecycleState(items []ExceptionLifecycleStage) string {
	if len(items) == 0 {
		return EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStateIncomplete
	}
	expected := map[string]struct{}{
		"requested":   {},
		"approved":    {},
		"activated":   {},
		"expiring":    {},
		"expired":     {},
		"revoked":     {},
		"superseded":  {},
		"revalidated": {},
	}
	for _, item := range items {
		status := strings.TrimSpace(item.Status)
		if status == "" || strings.TrimSpace(item.CurrentState) != "exception_lifecycle_ready" || !item.RequiresEvidence || len(item.ExpiryEffects) == 0 || !item.VerifierVisible {
			return EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStatePartial
		}
		delete(expected, status)
	}
	if len(expected) != 0 {
		return EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStatePartial
	}
	return EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStateActive
}

func EvaluateEnterpriseWorkflowAuthorityVal0ClosureValidationState(model ClosureValidationDiscipline) string {
	if len(model.RequiredChecks) == 0 || len(model.ReopenRules) == 0 || len(model.RollbackRules) == 0 {
		return EnterpriseWorkflowAuthorityVal0ClosureValidationStateIncomplete
	}
	requiredChecks := []string{
		"remediation_declared",
		"remediation_evidence_present",
		"validation_result_verified",
		"freshness_check_passed",
		"no_still_active_expired_exception",
		"closure_policy_satisfied",
	}
	for _, check := range requiredChecks {
		if !containsWorkflowValue(model.RequiredChecks, check) {
			return EnterpriseWorkflowAuthorityVal0ClosureValidationStatePartial
		}
	}
	if len(model.BoundedAlternativePolicy) == 0 {
		return EnterpriseWorkflowAuthorityVal0ClosureValidationStatePartial
	}
	return EnterpriseWorkflowAuthorityVal0ClosureValidationStateActive
}

func EvaluateEnterpriseWorkflowAuthorityVal0SeparationOfDutiesState(items []SeparationOfDutiesRule) string {
	if len(items) == 0 {
		return EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStateIncomplete
	}
	expected := map[string]struct{}{
		WorkflowAuthoritySensitiveActionBreakGlass:         {},
		WorkflowAuthoritySensitiveActionLongLivedException: {},
		WorkflowAuthoritySensitiveActionBroadScopeOverride: {},
		WorkflowAuthoritySensitiveActionProductionClosure:  {},
	}
	for _, item := range items {
		action := strings.TrimSpace(item.ActionClass)
		if action == "" || strings.TrimSpace(item.CurrentState) != "separation_of_duties_ready" || strings.TrimSpace(item.ControlMode) == "" || !item.DistinctApproverExecutor || strings.TrimSpace(item.BlastRadiusConstraint) == "" || len(item.RevocationPath) == 0 {
			return EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStatePartial
		}
		if action != WorkflowAuthoritySensitiveActionLongLivedException && !item.DualControlRequired {
			return EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStatePartial
		}
		delete(expected, action)
	}
	if len(expected) != 0 {
		return EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStatePartial
	}
	return EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStateActive
}

func EvaluateEnterpriseWorkflowAuthorityVal0TimeAuthorityState(model TimeAuthorityDiscipline) string {
	if strings.TrimSpace(model.CanonicalTimeSource) == "" || strings.TrimSpace(model.ClockSkewTolerance) == "" {
		return EnterpriseWorkflowAuthorityVal0TimeAuthorityStateIncomplete
	}
	if strings.TrimSpace(model.CanonicalTimeSource) != WorkflowAuthorityTimeSourceCanonicalService || len(model.ExpiryEvaluationRules) == 0 || len(model.FreshnessEvaluationRules) == 0 || len(model.TokenEvaluationRules) == 0 || len(model.ConnectorTimeRules) == 0 {
		return EnterpriseWorkflowAuthorityVal0TimeAuthorityStatePartial
	}
	if !containsWorkflowValue(model.ConnectorTimeRules, WorkflowAuthorityTimeSourceExternalAdvisory) {
		return EnterpriseWorkflowAuthorityVal0TimeAuthorityStatePartial
	}
	return EnterpriseWorkflowAuthorityVal0TimeAuthorityStateActive
}

func EvaluateEnterpriseWorkflowAuthorityVal0State(boundaryState, stateMachineState, projectionState, approvalContractState, exceptionLifecycleState, closureValidationState, separationOfDutiesState, timeAuthorityState string) string {
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(boundaryState),
		strings.TrimSpace(stateMachineState),
		strings.TrimSpace(projectionState),
		strings.TrimSpace(approvalContractState),
		strings.TrimSpace(exceptionLifecycleState),
		strings.TrimSpace(closureValidationState),
		strings.TrimSpace(separationOfDutiesState),
		strings.TrimSpace(timeAuthorityState),
	} {
		switch state {
		case EnterpriseWorkflowAuthorityVal0BoundaryStateActive,
			EnterpriseWorkflowAuthorityVal0StateMachineStateActive,
			EnterpriseWorkflowAuthorityVal0ProjectionStateActive,
			EnterpriseWorkflowAuthorityVal0ApprovalContractStateActive,
			EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStateActive,
			EnterpriseWorkflowAuthorityVal0ClosureValidationStateActive,
			EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStateActive,
			EnterpriseWorkflowAuthorityVal0TimeAuthorityStateActive:
		case EnterpriseWorkflowAuthorityVal0BoundaryStatePartial,
			EnterpriseWorkflowAuthorityVal0StateMachineStatePartial,
			EnterpriseWorkflowAuthorityVal0ProjectionStatePartial,
			EnterpriseWorkflowAuthorityVal0ApprovalContractStatePartial,
			EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStatePartial,
			EnterpriseWorkflowAuthorityVal0ClosureValidationStatePartial,
			EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStatePartial,
			EnterpriseWorkflowAuthorityVal0TimeAuthorityStatePartial:
			hasPartial = true
		default:
			return EnterpriseWorkflowAuthorityVal0StateIncomplete
		}
	}
	if hasPartial {
		return EnterpriseWorkflowAuthorityVal0StateSubstantial
	}
	return EnterpriseWorkflowAuthorityVal0StateActive
}

func containsWorkflowValue(values []string, expected string) bool {
	expected = strings.TrimSpace(expected)
	for _, value := range values {
		if strings.TrimSpace(value) == expected {
			return true
		}
	}
	return false
}

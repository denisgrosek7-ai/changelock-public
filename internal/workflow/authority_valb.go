package workflow

import "strings"

const (
	EnterpriseWorkflowAuthorityValBSignedAuthorizationsStateActive     = "enterprise_workflow_authority_valb_signed_authorizations_active"
	EnterpriseWorkflowAuthorityValBSignedAuthorizationsStatePartial    = "enterprise_workflow_authority_valb_signed_authorizations_partial"
	EnterpriseWorkflowAuthorityValBSignedAuthorizationsStateIncomplete = "enterprise_workflow_authority_valb_signed_authorizations_incomplete"

	EnterpriseWorkflowAuthorityValBBreakGlassStateActive     = "enterprise_workflow_authority_valb_break_glass_active"
	EnterpriseWorkflowAuthorityValBBreakGlassStatePartial    = "enterprise_workflow_authority_valb_break_glass_partial"
	EnterpriseWorkflowAuthorityValBBreakGlassStateIncomplete = "enterprise_workflow_authority_valb_break_glass_incomplete"

	EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStateActive     = "enterprise_workflow_authority_valb_managed_exception_registry_active"
	EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStatePartial    = "enterprise_workflow_authority_valb_managed_exception_registry_partial"
	EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStateIncomplete = "enterprise_workflow_authority_valb_managed_exception_registry_incomplete"

	EnterpriseWorkflowAuthorityValBExpiryRevocationStateActive     = "enterprise_workflow_authority_valb_expiry_revocation_active"
	EnterpriseWorkflowAuthorityValBExpiryRevocationStatePartial    = "enterprise_workflow_authority_valb_expiry_revocation_partial"
	EnterpriseWorkflowAuthorityValBExpiryRevocationStateIncomplete = "enterprise_workflow_authority_valb_expiry_revocation_incomplete"

	EnterpriseWorkflowAuthorityValBAntiReplayStateActive     = "enterprise_workflow_authority_valb_anti_replay_active"
	EnterpriseWorkflowAuthorityValBAntiReplayStatePartial    = "enterprise_workflow_authority_valb_anti_replay_partial"
	EnterpriseWorkflowAuthorityValBAntiReplayStateIncomplete = "enterprise_workflow_authority_valb_anti_replay_incomplete"

	EnterpriseWorkflowAuthorityValBApprovalTraceabilityStateActive     = "enterprise_workflow_authority_valb_approval_traceability_active"
	EnterpriseWorkflowAuthorityValBApprovalTraceabilityStatePartial    = "enterprise_workflow_authority_valb_approval_traceability_partial"
	EnterpriseWorkflowAuthorityValBApprovalTraceabilityStateIncomplete = "enterprise_workflow_authority_valb_approval_traceability_incomplete"

	EnterpriseWorkflowAuthorityValBStateIncomplete  = "enterprise_workflow_authority_valb_incomplete"
	EnterpriseWorkflowAuthorityValBStateSubstantial = "enterprise_workflow_authority_valb_substantially_ready"
	EnterpriseWorkflowAuthorityValBStateActive      = "enterprise_workflow_authority_valb_active"
)

type WorkflowSignedAuthorizationArtifactBaseline struct {
	CurrentState           string   `json:"current_state"`
	RequiredArtifactFields []string `json:"required_artifact_fields,omitempty"`
	SupportedActionClasses []string `json:"supported_action_classes,omitempty"`
	ConsumptionModes       []string `json:"consumption_modes,omitempty"`
	SignatureRules         []string `json:"signature_rules,omitempty"`
	SubjectBindingRules    []string `json:"subject_binding_rules,omitempty"`
	ScopeBindingRules      []string `json:"scope_binding_rules,omitempty"`
	TimeAuthorityRules     []string `json:"time_authority_rules,omitempty"`
	RevocationRules        []string `json:"revocation_rules,omitempty"`
	AntiReplayFields       []string `json:"anti_replay_fields,omitempty"`
	Limitations            []string `json:"limitations,omitempty"`
}

type WorkflowBreakGlassControlBaseline struct {
	CurrentState             string   `json:"current_state"`
	RequiredStages           []string `json:"required_stages,omitempty"`
	SupportedActionClasses   []string `json:"supported_action_classes,omitempty"`
	DistinctApproverExecutor bool     `json:"distinct_approver_executor"`
	DualControlRequired      bool     `json:"dual_control_required"`
	MaximumDuration          string   `json:"maximum_duration"`
	BlastRadiusFields        []string `json:"blast_radius_fields,omitempty"`
	ActivationRules          []string `json:"activation_rules,omitempty"`
	RevocationPath           []string `json:"revocation_path,omitempty"`
	ConsumptionModes         []string `json:"consumption_modes,omitempty"`
	Limitations              []string `json:"limitations,omitempty"`
}

type WorkflowManagedExceptionRegistryBaseline struct {
	CurrentState      string   `json:"current_state"`
	RequiredFields    []string `json:"required_fields,omitempty"`
	LifecycleStages   []string `json:"lifecycle_stages,omitempty"`
	AutoExpiryEffects []string `json:"auto_expiry_effects,omitempty"`
	RevocationRules   []string `json:"revocation_rules,omitempty"`
	SupersessionRules []string `json:"supersession_rules,omitempty"`
	RevalidationRules []string `json:"revalidation_rules,omitempty"`
	VisibilityRules   []string `json:"visibility_rules,omitempty"`
	Limitations       []string `json:"limitations,omitempty"`
}

type WorkflowExpiryRevocationEnforcementBaseline struct {
	CurrentState              string   `json:"current_state"`
	CanonicalTimeSource       string   `json:"canonical_time_source"`
	ClockSkewTolerance        string   `json:"clock_skew_tolerance"`
	ExpiryEvaluationRules     []string `json:"expiry_evaluation_rules,omitempty"`
	RevocationEvaluationRules []string `json:"revocation_evaluation_rules,omitempty"`
	ConnectorPropagationRules []string `json:"connector_propagation_rules,omitempty"`
	ReopenRules               []string `json:"reopen_rules,omitempty"`
	Limitations               []string `json:"limitations,omitempty"`
}

type WorkflowAntiReplayProtectionBaseline struct {
	CurrentState               string   `json:"current_state"`
	TokenTypes                 []string `json:"token_types,omitempty"`
	NonceOrJTIFields           []string `json:"nonce_or_jti_fields,omitempty"`
	ConsumptionModes           []string `json:"consumption_modes,omitempty"`
	ReplayCacheRules           []string `json:"replay_cache_rules,omitempty"`
	DuplicateSuppressionRules  []string `json:"duplicate_suppression_rules,omitempty"`
	RevocationInteractionRules []string `json:"revocation_interaction_rules,omitempty"`
	Limitations                []string `json:"limitations,omitempty"`
}

type WorkflowApprovalTraceabilityBaseline struct {
	CurrentState         string   `json:"current_state"`
	RequiredTraceFields  []string `json:"required_trace_fields,omitempty"`
	IdentityTrailRules   []string `json:"identity_trail_rules,omitempty"`
	EvidenceLinkageRules []string `json:"evidence_linkage_rules,omitempty"`
	ExternalRefRules     []string `json:"external_ref_rules,omitempty"`
	ResultingActionRules []string `json:"resulting_action_rules,omitempty"`
	SupersessionRules    []string `json:"supersession_rules,omitempty"`
	Limitations          []string `json:"limitations,omitempty"`
}

func EnterpriseWorkflowAuthorityValBSignedAuthorizations() WorkflowSignedAuthorizationArtifactBaseline {
	model := WorkflowSignedAuthorizationArtifactBaseline{
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
		SupportedActionClasses: []string{
			WorkflowAuthorityActionApprovalRequired,
			WorkflowAuthoritySensitiveActionBreakGlass,
			WorkflowAuthoritySensitiveActionBroadScopeOverride,
			WorkflowAuthoritySensitiveActionProductionClosure,
			WorkflowAuthoritySensitiveActionLongLivedException,
		},
		ConsumptionModes: []string{
			WorkflowAuthorityConsumptionSingleUse,
			WorkflowAuthorityConsumptionMultiUseBounded,
			WorkflowAuthorityConsumptionSessionBound,
		},
		SignatureRules: []string{
			"authorization_artifact_must_be_signed_or_sealed",
			"approver_identity_must_match_signature_subject",
			"signature_verification_must_use_canonical_trust_policy",
		},
		SubjectBindingRules: []string{
			"subject_ref_must_match_target_workflow_object",
			"artifact_scope_must_bind_to_canonical_subject_and_environment",
		},
		ScopeBindingRules: []string{
			"scope_must_encode_system_environment_and_blast_radius",
			"broad_scope_override_requires_distinct_or_dual_control",
		},
		TimeAuthorityRules: []string{
			"issued_at_and_expires_at_must_use_canonical_service_time",
			"expiry_evaluation_must_honor_clock_skew_tolerance",
		},
		RevocationRules: []string{
			"revoked_artifact_becomes_immediately_ineffective",
			"revocation_handle_must_link_to_canonical_workflow_record",
		},
		AntiReplayFields: []string{
			"anti_replay_marker_jti",
			"subject_nonce_binding",
			"consumed_at",
		},
		Limitations: []string{
			"Val B defines signed delegated authority posture and does not yet expose the later workflow ledger or final authority gate.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValBSignedAuthorizationsState(model)
	return model
}

func EnterpriseWorkflowAuthorityValBBreakGlassFlow() WorkflowBreakGlassControlBaseline {
	model := WorkflowBreakGlassControlBaseline{
		RequiredStages: []string{
			"request",
			"approval",
			"activation",
			"revocation",
			"expiry",
		},
		SupportedActionClasses: []string{
			WorkflowAuthoritySensitiveActionBreakGlass,
			WorkflowAuthoritySensitiveActionBroadScopeOverride,
			WorkflowAuthoritySensitiveActionProductionClosure,
		},
		DistinctApproverExecutor: true,
		DualControlRequired:      true,
		MaximumDuration:          "4h",
		BlastRadiusFields: []string{
			"subject_scope",
			"system_scope",
			"environment_scope",
			"blast_radius",
		},
		ActivationRules: []string{
			"break_glass_cannot_activate_without_signed_authorization_artifact",
			"activation_requires_identity_bound_scope_and_justification",
		},
		RevocationPath: []string{
			"manual_revoke",
			"auto_expire",
			"revoke_on_superseding_authorization",
		},
		ConsumptionModes: []string{
			WorkflowAuthorityConsumptionSingleUse,
			WorkflowAuthorityConsumptionSessionBound,
		},
		Limitations: []string{
			"Break-glass remains bounded to emergency action classes and does not convert external workflow systems into canonical closure authority.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValBBreakGlassState(model)
	return model
}

func EnterpriseWorkflowAuthorityValBManagedExceptionRegistry() WorkflowManagedExceptionRegistryBaseline {
	model := WorkflowManagedExceptionRegistryBaseline{
		RequiredFields: []string{
			"exception_id",
			"subject",
			"policy_link",
			"approver",
			"reason",
			"issued_at",
			"expires_at",
			"current_state",
			"linked_evidence",
			"revocation_status",
			"superseded_by",
		},
		LifecycleStages: []string{
			"requested",
			"approved",
			"activated",
			"expiring",
			"expired",
			"revoked",
			"superseded",
			"revalidated",
		},
		AutoExpiryEffects: []string{
			"expired_exception_blocks_or_reopens_workflow",
			"grace_denied_without_new_authorization",
			"connector_projection_must_show_expired_state",
		},
		RevocationRules: []string{
			"revocation_must_invalidate_active_exception_effect",
			"revocation_must_be_visible_to_connector_projection_and_verifier",
		},
		SupersessionRules: []string{
			"superseded_exception_must_link_to_replacement_object",
			"replacement_must_not_inherit_scope_without_new_evidence_and_expiry",
		},
		RevalidationRules: []string{
			"revalidation_requires_fresh_evidence_and_policy_link",
			"revalidation_must_issue_new_expiry_window",
		},
		VisibilityRules: []string{
			"active_exception_must_be_canonical_and_connector_visible",
			"expired_or_revoked_exception_must_remain_auditable",
		},
		Limitations: []string{
			"Val B defines the managed exception registry baseline before later Point 3 waves attach closure hardening and workflow ledger persistence.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValBManagedExceptionRegistryState(model)
	return model
}

func EnterpriseWorkflowAuthorityValBExpiryRevocationEnforcement() WorkflowExpiryRevocationEnforcementBaseline {
	model := WorkflowExpiryRevocationEnforcementBaseline{
		CanonicalTimeSource: WorkflowAuthorityTimeSourceCanonicalService,
		ClockSkewTolerance:  "5m",
		ExpiryEvaluationRules: []string{
			"authorization_and_exception_expiry_must_use_canonical_service_time",
			"expired_artifact_or_exception_must_not_remain_effective_after_clock_skew_window",
		},
		RevocationEvaluationRules: []string{
			"revoked_artifact_must_fail_closed_even_if_external_projection_lags",
			"revocation_must_be_checked_before_consumption_and_replay",
		},
		ConnectorPropagationRules: []string{
			"connector_projection_must_reflect_expired_or_revoked_state_on_next_sync",
			"degraded_connector_mode_cannot_mask_expiry_or_revocation_in_canonical_state",
		},
		ReopenRules: []string{
			"expired_exception_or_revoked_override_may_reopen_pending_or_closed_workflow",
			"reopen_must_remain_evidence_bound_and_auditable",
		},
		Limitations: []string{
			"Expiry and revocation remain canonical-service-time bound even when external systems provide conflicting clocks or stale close signals.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValBExpiryRevocationState(model)
	return model
}

func EnterpriseWorkflowAuthorityValBAntiReplayProtection() WorkflowAntiReplayProtectionBaseline {
	model := WorkflowAntiReplayProtectionBaseline{
		TokenTypes: []string{
			"signed_authorization_artifact",
			"break_glass_authorization",
			"exception_activation_grant",
		},
		NonceOrJTIFields: []string{
			"anti_replay_marker_jti",
			"subject_nonce_binding",
			"consumed_at",
		},
		ConsumptionModes: []string{
			WorkflowAuthorityConsumptionSingleUse,
			WorkflowAuthorityConsumptionMultiUseBounded,
			WorkflowAuthorityConsumptionSessionBound,
		},
		ReplayCacheRules: []string{
			"jti_must_be_rejected_after_consumption",
			"session_bound_grant_must_be_tied_to_canonical_session_identity",
		},
		DuplicateSuppressionRules: []string{
			"duplicate_authorization_consumption_must_be_rejected",
			"replayed_connector_reflection_must_not_reactivate_consumed_grant",
		},
		RevocationInteractionRules: []string{
			"revoked_grant_must_invalidate_cached_replay_entries",
			"revocation_check_precedes_replay_cache_acceptance",
		},
		Limitations: []string{
			"Val B anti-replay protection applies to authority artifacts and remains distinct from connector mutation replay discipline introduced in Val A.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValBAntiReplayState(model)
	return model
}

func EnterpriseWorkflowAuthorityValBApprovalTraceability() WorkflowApprovalTraceabilityBaseline {
	model := WorkflowApprovalTraceabilityBaseline{
		RequiredTraceFields: []string{
			"authorization_id",
			"actor_identity",
			"subject",
			"action_class",
			"scope",
			"evidence_refs",
			"external_system_refs",
			"resulting_canonical_transition",
			"resulting_external_mutation_outcome",
		},
		IdentityTrailRules: []string{
			"approver_identity_must_be_bound_to_signed_artifact",
			"distinct_executor_identity_must_be_recorded_for_sensitive_actions",
		},
		EvidenceLinkageRules: []string{
			"approval_trace_must_link_to_evidence_bundle_and_justification",
			"exception_and_override_traces_must_link_to_subject_scope_and_policy_ref",
		},
		ExternalRefRules: []string{
			"external_ticket_or_change_refs_remain_projection_refs_not_truth_override",
			"connector_replay_or_sync_back_must_preserve_original_authorization_ref",
		},
		ResultingActionRules: []string{
			"trace_must_record_resulting_canonical_transition",
			"trace_must_record_resulting_external_mutation_outcome",
		},
		SupersessionRules: []string{
			"superseded_authorization_trace_must_link_to_replacement_artifact",
			"revoked_or_consumed_trace_entries_must_remain_visible",
		},
		Limitations: []string{
			"Val B provides approval traceability baseline before later Point 3 waves add append-only workflow ledger review and final gate enforcement.",
		},
	}
	model.CurrentState = EvaluateEnterpriseWorkflowAuthorityValBApprovalTraceabilityState(model)
	return model
}

func EvaluateEnterpriseWorkflowAuthorityValBSignedAuthorizationsState(model WorkflowSignedAuthorizationArtifactBaseline) string {
	if len(model.RequiredArtifactFields) == 0 || len(model.SupportedActionClasses) == 0 {
		return EnterpriseWorkflowAuthorityValBSignedAuthorizationsStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredArtifactFields,
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
	) || !containsAllTrimmedStrings(model.ConsumptionModes,
		WorkflowAuthorityConsumptionSingleUse,
		WorkflowAuthorityConsumptionMultiUseBounded,
		WorkflowAuthorityConsumptionSessionBound,
	) || !containsAllTrimmedStrings(model.AntiReplayFields,
		"anti_replay_marker_jti",
		"subject_nonce_binding",
		"consumed_at",
	) || len(model.SignatureRules) == 0 || len(model.SubjectBindingRules) == 0 || len(model.ScopeBindingRules) == 0 || len(model.TimeAuthorityRules) == 0 || len(model.RevocationRules) == 0 {
		return EnterpriseWorkflowAuthorityValBSignedAuthorizationsStatePartial
	}
	return EnterpriseWorkflowAuthorityValBSignedAuthorizationsStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValBBreakGlassState(model WorkflowBreakGlassControlBaseline) string {
	if len(model.RequiredStages) == 0 || len(model.SupportedActionClasses) == 0 {
		return EnterpriseWorkflowAuthorityValBBreakGlassStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredStages, "request", "approval", "activation", "revocation", "expiry") ||
		!containsAllTrimmedStrings(model.SupportedActionClasses,
			WorkflowAuthoritySensitiveActionBreakGlass,
			WorkflowAuthoritySensitiveActionBroadScopeOverride,
			WorkflowAuthoritySensitiveActionProductionClosure,
		) ||
		!model.DistinctApproverExecutor ||
		!model.DualControlRequired ||
		strings.TrimSpace(model.MaximumDuration) == "" ||
		len(model.BlastRadiusFields) == 0 ||
		len(model.ActivationRules) == 0 ||
		len(model.RevocationPath) == 0 ||
		!containsAllTrimmedStrings(model.ConsumptionModes,
			WorkflowAuthorityConsumptionSingleUse,
			WorkflowAuthorityConsumptionSessionBound,
		) {
		return EnterpriseWorkflowAuthorityValBBreakGlassStatePartial
	}
	return EnterpriseWorkflowAuthorityValBBreakGlassStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValBManagedExceptionRegistryState(model WorkflowManagedExceptionRegistryBaseline) string {
	if len(model.RequiredFields) == 0 || len(model.LifecycleStages) == 0 {
		return EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredFields,
		"exception_id",
		"subject",
		"policy_link",
		"approver",
		"reason",
		"issued_at",
		"expires_at",
		"current_state",
		"linked_evidence",
		"revocation_status",
		"superseded_by",
	) || !containsAllTrimmedStrings(model.LifecycleStages,
		"requested",
		"approved",
		"activated",
		"expiring",
		"expired",
		"revoked",
		"superseded",
		"revalidated",
	) || len(model.AutoExpiryEffects) == 0 || len(model.RevocationRules) == 0 || len(model.SupersessionRules) == 0 || len(model.RevalidationRules) == 0 || len(model.VisibilityRules) == 0 {
		return EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStatePartial
	}
	return EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValBExpiryRevocationState(model WorkflowExpiryRevocationEnforcementBaseline) string {
	if strings.TrimSpace(model.CanonicalTimeSource) == "" || strings.TrimSpace(model.ClockSkewTolerance) == "" {
		return EnterpriseWorkflowAuthorityValBExpiryRevocationStateIncomplete
	}
	if strings.TrimSpace(model.CanonicalTimeSource) != WorkflowAuthorityTimeSourceCanonicalService ||
		len(model.ExpiryEvaluationRules) == 0 ||
		len(model.RevocationEvaluationRules) == 0 ||
		len(model.ConnectorPropagationRules) == 0 ||
		len(model.ReopenRules) == 0 {
		return EnterpriseWorkflowAuthorityValBExpiryRevocationStatePartial
	}
	return EnterpriseWorkflowAuthorityValBExpiryRevocationStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValBAntiReplayState(model WorkflowAntiReplayProtectionBaseline) string {
	if len(model.TokenTypes) == 0 || len(model.NonceOrJTIFields) == 0 {
		return EnterpriseWorkflowAuthorityValBAntiReplayStateIncomplete
	}
	if !containsAllTrimmedStrings(model.NonceOrJTIFields,
		"anti_replay_marker_jti",
		"subject_nonce_binding",
		"consumed_at",
	) || !containsAllTrimmedStrings(model.ConsumptionModes,
		WorkflowAuthorityConsumptionSingleUse,
		WorkflowAuthorityConsumptionMultiUseBounded,
		WorkflowAuthorityConsumptionSessionBound,
	) || len(model.ReplayCacheRules) == 0 || len(model.DuplicateSuppressionRules) == 0 || len(model.RevocationInteractionRules) == 0 {
		return EnterpriseWorkflowAuthorityValBAntiReplayStatePartial
	}
	return EnterpriseWorkflowAuthorityValBAntiReplayStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValBApprovalTraceabilityState(model WorkflowApprovalTraceabilityBaseline) string {
	if len(model.RequiredTraceFields) == 0 {
		return EnterpriseWorkflowAuthorityValBApprovalTraceabilityStateIncomplete
	}
	if !containsAllTrimmedStrings(model.RequiredTraceFields,
		"authorization_id",
		"actor_identity",
		"subject",
		"action_class",
		"scope",
		"evidence_refs",
		"external_system_refs",
		"resulting_canonical_transition",
		"resulting_external_mutation_outcome",
	) || len(model.IdentityTrailRules) == 0 || len(model.EvidenceLinkageRules) == 0 || len(model.ExternalRefRules) == 0 || len(model.ResultingActionRules) == 0 || len(model.SupersessionRules) == 0 {
		return EnterpriseWorkflowAuthorityValBApprovalTraceabilityStatePartial
	}
	return EnterpriseWorkflowAuthorityValBApprovalTraceabilityStateActive
}

func EvaluateEnterpriseWorkflowAuthorityValBState(
	valAState,
	signedAuthorizationsState,
	breakGlassState,
	managedExceptionRegistryState,
	expiryRevocationState,
	antiReplayState,
	approvalTraceabilityState string,
) string {
	if strings.TrimSpace(valAState) != EnterpriseWorkflowAuthorityValAStateActive {
		return EnterpriseWorkflowAuthorityValBStateIncomplete
	}

	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(signedAuthorizationsState),
		strings.TrimSpace(breakGlassState),
		strings.TrimSpace(managedExceptionRegistryState),
		strings.TrimSpace(expiryRevocationState),
		strings.TrimSpace(antiReplayState),
		strings.TrimSpace(approvalTraceabilityState),
	} {
		switch state {
		case EnterpriseWorkflowAuthorityValBSignedAuthorizationsStateActive,
			EnterpriseWorkflowAuthorityValBBreakGlassStateActive,
			EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStateActive,
			EnterpriseWorkflowAuthorityValBExpiryRevocationStateActive,
			EnterpriseWorkflowAuthorityValBAntiReplayStateActive,
			EnterpriseWorkflowAuthorityValBApprovalTraceabilityStateActive:
		case EnterpriseWorkflowAuthorityValBSignedAuthorizationsStatePartial,
			EnterpriseWorkflowAuthorityValBBreakGlassStatePartial,
			EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStatePartial,
			EnterpriseWorkflowAuthorityValBExpiryRevocationStatePartial,
			EnterpriseWorkflowAuthorityValBAntiReplayStatePartial,
			EnterpriseWorkflowAuthorityValBApprovalTraceabilityStatePartial:
			hasPartial = true
		default:
			return EnterpriseWorkflowAuthorityValBStateIncomplete
		}
	}
	if hasPartial {
		return EnterpriseWorkflowAuthorityValBStateSubstantial
	}
	return EnterpriseWorkflowAuthorityValBStateActive
}

func containsAllTrimmedStrings(values []string, needles ...string) bool {
	for _, needle := range needles {
		if !containsTrimmedString(values, needle) {
			return false
		}
	}
	return true
}

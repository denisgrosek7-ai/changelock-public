package operability

import (
	"strings"

	workflowauthority "github.com/denisgrosek/changelock/internal/workflow"
)

const (
	ProductionUsabilityVal0ConfigIntegrityStateActive     = "production_usability_val0_config_integrity_active"
	ProductionUsabilityVal0ConfigIntegrityStatePartial    = "production_usability_val0_config_integrity_partial"
	ProductionUsabilityVal0ConfigIntegrityStateIncomplete = "production_usability_val0_config_integrity_incomplete"

	ProductionUsabilityVal0ExplainabilityStateActive     = "production_usability_val0_explainability_contract_active"
	ProductionUsabilityVal0ExplainabilityStatePartial    = "production_usability_val0_explainability_contract_partial"
	ProductionUsabilityVal0ExplainabilityStateIncomplete = "production_usability_val0_explainability_contract_incomplete"

	ProductionUsabilityVal0StatusModelStateActive     = "production_usability_val0_status_model_active"
	ProductionUsabilityVal0StatusModelStatePartial    = "production_usability_val0_status_model_partial"
	ProductionUsabilityVal0StatusModelStateIncomplete = "production_usability_val0_status_model_incomplete"

	ProductionUsabilityVal0OperationContractStateActive     = "production_usability_val0_operation_contract_active"
	ProductionUsabilityVal0OperationContractStatePartial    = "production_usability_val0_operation_contract_partial"
	ProductionUsabilityVal0OperationContractStateIncomplete = "production_usability_val0_operation_contract_incomplete"

	ProductionUsabilityVal0DecisionQualityStateActive     = "production_usability_val0_decision_quality_active"
	ProductionUsabilityVal0DecisionQualityStatePartial    = "production_usability_val0_decision_quality_partial"
	ProductionUsabilityVal0DecisionQualityStateIncomplete = "production_usability_val0_decision_quality_incomplete"

	ProductionUsabilityVal0NotificationStateActive     = "production_usability_val0_notification_taxonomy_active"
	ProductionUsabilityVal0NotificationStatePartial    = "production_usability_val0_notification_taxonomy_partial"
	ProductionUsabilityVal0NotificationStateIncomplete = "production_usability_val0_notification_taxonomy_incomplete"

	ProductionUsabilityVal0PermissionRedactionStateActive     = "production_usability_val0_permission_redaction_active"
	ProductionUsabilityVal0PermissionRedactionStatePartial    = "production_usability_val0_permission_redaction_partial"
	ProductionUsabilityVal0PermissionRedactionStateIncomplete = "production_usability_val0_permission_redaction_incomplete"

	ProductionUsabilityVal0RecoveryStateActive     = "production_usability_val0_recovery_contract_active"
	ProductionUsabilityVal0RecoveryStatePartial    = "production_usability_val0_recovery_contract_partial"
	ProductionUsabilityVal0RecoveryStateIncomplete = "production_usability_val0_recovery_contract_incomplete"

	ProductionUsabilityVal0ActionModeStateActive     = "production_usability_val0_action_modes_active"
	ProductionUsabilityVal0ActionModeStatePartial    = "production_usability_val0_action_modes_partial"
	ProductionUsabilityVal0ActionModeStateIncomplete = "production_usability_val0_action_modes_incomplete"

	ProductionUsabilityVal0StateIncomplete  = "production_usability_val0_incomplete"
	ProductionUsabilityVal0StateSubstantial = "production_usability_val0_substantially_ready"
	ProductionUsabilityVal0StateActive      = "production_usability_val0_active"

	ProductionUsabilityPoint4StateNotComplete = "production_usability_point_4_not_complete"

	ProductionUsabilityCompatibilityCurrent           = "current"
	ProductionUsabilityCompatibilityDeprecated        = "deprecated"
	ProductionUsabilityCompatibilityMigrationRequired = "migration_required"
	ProductionUsabilityCompatibilityUnsupported       = "unsupported"
	ProductionUsabilityCompatibilityUnknown           = "unknown"

	ProductionUsabilityUnknownFieldReject                  = "reject"
	ProductionUsabilityUnknownFieldWarn                    = "warn"
	ProductionUsabilityUnknownFieldIgnoreExplicitlyAllowed = "ignore_only_if_explicitly_allowed"

	ProductionUsabilityValidationValid       = "valid"
	ProductionUsabilityValidationInvalid     = "invalid"
	ProductionUsabilityValidationDegraded    = "degraded"
	ProductionUsabilityValidationUnsupported = "unsupported"

	ProductionUsabilitySeverityCritical = "critical"
	ProductionUsabilitySeverityError    = "error"
	ProductionUsabilitySeverityWarning  = "warning"
	ProductionUsabilitySeverityInfo     = "info"

	ProductionUsabilityVisibilityInternalAdmin = "internal_admin"
	ProductionUsabilityVisibilityOperator      = "operator"
	ProductionUsabilityVisibilityDeveloper     = "developer"
	ProductionUsabilityVisibilityPartner       = "partner"
	ProductionUsabilityVisibilityPublicSafe    = "public_safe"

	ProductionUsabilityRedactionNone       = "none"
	ProductionUsabilityRedactionLow        = "low"
	ProductionUsabilityRedactionMedium     = "medium"
	ProductionUsabilityRedactionHigh       = "high"
	ProductionUsabilityRedactionPublicSafe = "public_safe"

	ProductionUsabilityStatusFresh       = "fresh"
	ProductionUsabilityStatusStale       = "stale"
	ProductionUsabilityStatusPartial     = "partial"
	ProductionUsabilityStatusDegraded    = "degraded"
	ProductionUsabilityStatusUnavailable = "unavailable"
	ProductionUsabilityStatusUnsupported = "unsupported"

	ProductionUsabilityOperationReadOnly            = "read_only"
	ProductionUsabilityOperationPreviewOnly         = "preview_only"
	ProductionUsabilityOperationIdempotentMutation  = "idempotent_mutation"
	ProductionUsabilityOperationNonIdempotentMutate = "non_idempotent_mutation"
	ProductionUsabilityOperationSideEffecting       = "side_effecting"

	ProductionUsabilityRetrySafe              = "retry_safe"
	ProductionUsabilityRetryUnsafe            = "retry_unsafe"
	ProductionUsabilityRetryConditionallySafe = "retry_conditionally_safe"

	ProductionUsabilityDecisionBlocker       = "blocker"
	ProductionUsabilityDecisionUrgent        = "urgent"
	ProductionUsabilityDecisionNormal        = "normal"
	ProductionUsabilityDecisionLow           = "low"
	ProductionUsabilityDecisionInformational = "informational"

	ProductionUsabilityNoAction         = "no_action"
	ProductionUsabilityOperatorAction   = "operator_action"
	ProductionUsabilityAdminAction      = "admin_action"
	ProductionUsabilityGovernanceAction = "governance_action"
	ProductionUsabilitySupportAction    = "support_action"

	ProductionUsabilityAckUnacknowledged = "unacknowledged"
	ProductionUsabilityAckAcknowledged   = "acknowledged"
	ProductionUsabilityAckSuppressedDup  = "suppressed_duplicate"
	ProductionUsabilityAckResolved       = "resolved"
	ProductionUsabilityAckReopened       = "reopened"

	ProductionUsabilityEvidenceFull         = "full"
	ProductionUsabilityEvidenceMetadataOnly = "metadata_only"
	ProductionUsabilityEvidenceRedacted     = "redacted"
	ProductionUsabilityEvidenceHidden       = "hidden"

	ProductionUsabilityRemediationConfigFix       = "config_fix"
	ProductionUsabilityRemediationPermissionFix   = "permission_fix"
	ProductionUsabilityRemediationEvidenceRefresh = "evidence_refresh"
	ProductionUsabilityRemediationPolicyUpdate    = "policy_update"
	ProductionUsabilityRemediationWaitAndRetry    = "wait_and_retry"
	ProductionUsabilityRemediationManualReview    = "manual_review"
	ProductionUsabilityRemediationSupportEsc      = "support_escalation"
	ProductionUsabilityRemediationUnsupported     = "unsupported"

	ProductionUsabilityActionModeViewOnly  = "view_only"
	ProductionUsabilityActionModeExplain   = "explain_only"
	ProductionUsabilityActionModePreview   = "preview"
	ProductionUsabilityActionModeDryRun    = "dry_run"
	ProductionUsabilityActionModeAuditOnly = "audit_only"
	ProductionUsabilityActionModeEnforce   = "enforce"
	ProductionUsabilityActionModeMutate    = "mutate"
)

type ConfigIntegrityContract struct {
	CurrentState               string   `json:"current_state"`
	SchemaVersion              string   `json:"schema_version"`
	SupportedCompatibility     []string `json:"supported_compatibility_statuses,omitempty"`
	CurrentCompatibility       string   `json:"current_compatibility_status"`
	UnknownFieldPolicies       []string `json:"unknown_field_policies,omitempty"`
	CurrentUnknownFieldPolicy  string   `json:"current_unknown_field_policy"`
	ValidationStates           []string `json:"validation_states,omitempty"`
	CurrentValidationResult    string   `json:"current_validation_result"`
	MigrationWarningSemantics  []string `json:"migration_warning_semantics,omitempty"`
	ConflictDetectionSemantics []string `json:"conflict_detection_semantics,omitempty"`
	EffectiveConfigView        string   `json:"effective_config_view_contract"`
	FailFastBootstrap          string   `json:"fail_fast_bootstrap_expectation"`
	Limitations                []string `json:"limitations,omitempty"`
}

type ExplainabilityPayloadContract struct {
	CurrentState                string   `json:"current_state"`
	RequiredFields              []string `json:"required_fields,omitempty"`
	SupportedSeverities         []string `json:"supported_severities,omitempty"`
	SupportedVisibilityScopes   []string `json:"supported_visibility_scopes,omitempty"`
	SupportedRedactionTiers     []string `json:"supported_redaction_tiers,omitempty"`
	SupportedDecisionPriorities []string `json:"supported_decision_priorities,omitempty"`
	SupportedActionModes        []string `json:"supported_action_modes,omitempty"`
	SafeRedactionScopes         []string `json:"safe_redaction_scopes,omitempty"`
	TechnicalDetailBounded      bool     `json:"technical_detail_bounded"`
	PreservesFailureSemantics   bool     `json:"preserves_failure_semantics"`
	Rules                       []string `json:"rules,omitempty"`
	Limitations                 []string `json:"limitations,omitempty"`
}

type OperationalStatusDefinition struct {
	Status                   string   `json:"status"`
	CurrentState             string   `json:"current_state"`
	FreshnessIndicator       string   `json:"freshness_indicator"`
	SourceProjectionID       string   `json:"source_projection_id"`
	LimitationMessage        string   `json:"limitation_message"`
	EvidenceRefsPolicy       string   `json:"evidence_refs_policy"`
	CanonicalTruthDisclaimer string   `json:"canonical_truth_disclaimer"`
	ClaimsCanonicalTruth     bool     `json:"claims_canonical_truth"`
	Limitations              []string `json:"limitations,omitempty"`
}

type OperationalStatusModel struct {
	CurrentState         string                        `json:"current_state"`
	Items                []OperationalStatusDefinition `json:"items,omitempty"`
	RequiredDistinctions []string                      `json:"required_distinctions,omitempty"`
	Limitations          []string                      `json:"limitations,omitempty"`
}

type UsabilityOperation struct {
	OperationName          string   `json:"operation_name"`
	OperationType          string   `json:"operation_type"`
	RetrySafety            string   `json:"retry_safety"`
	SideEffects            []string `json:"side_effects,omitempty"`
	IdempotencyKeyRequired bool     `json:"idempotency_key_required"`
	SafeRetryGuidance      string   `json:"safe_retry_guidance"`
	DoNotRetryReason       string   `json:"do_not_retry_reason"`
	PermissionScope        string   `json:"permission_scope"`
}

type OperationContractModel struct {
	CurrentState           string               `json:"current_state"`
	Items                  []UsabilityOperation `json:"items,omitempty"`
	RequiredOperationTypes []string             `json:"required_operation_types,omitempty"`
	SupportedRetrySafety   []string             `json:"supported_retry_safety,omitempty"`
	Rules                  []string             `json:"rules,omitempty"`
	Limitations            []string             `json:"limitations,omitempty"`
}

type DecisionQualityContract struct {
	CurrentState          string   `json:"current_state"`
	DecisionPriorities    []string `json:"decision_priorities,omitempty"`
	ActionRequiredClasses []string `json:"action_required_classes,omitempty"`
	RequiredFields        []string `json:"required_fields,omitempty"`
	AdvisoryOnly          bool     `json:"advisory_only"`
	Limitations           []string `json:"limitations,omitempty"`
}

type NotificationTaxonomyContract struct {
	CurrentState                 string   `json:"current_state"`
	Severities                   []string `json:"severities,omitempty"`
	AcknowledgementStates        []string `json:"acknowledgement_states,omitempty"`
	RequiredFields               []string `json:"required_fields,omitempty"`
	AcknowledgementIsRemediation bool     `json:"acknowledgement_is_remediation"`
	ResolvedEqualsClosure        bool     `json:"resolved_equals_canonical_closure"`
	CriticalSuppressionAllowed   bool     `json:"critical_suppression_allowed"`
	ReopenOnChangeExplicit       bool     `json:"reopen_on_change_explicit"`
	Limitations                  []string `json:"limitations,omitempty"`
}

type PermissionAwareExplanationRule struct {
	VisibilityScope     string `json:"visibility_scope"`
	CurrentState        string `json:"current_state"`
	AllowedDetailLevel  string `json:"allowed_detail_level"`
	RedactionTier       string `json:"redaction_tier"`
	EvidenceVisibility  string `json:"evidence_visibility"`
	SafeFallbackMessage string `json:"safe_fallback_message"`
}

type PermissionRedactionContract struct {
	CurrentState                string                           `json:"current_state"`
	Items                       []PermissionAwareExplanationRule `json:"items,omitempty"`
	SupportedEvidenceVisibility []string                         `json:"supported_evidence_visibility,omitempty"`
	TechnicalDetailBounded      bool                             `json:"technical_detail_bounded"`
	PreservesFailureSemantics   bool                             `json:"preserves_failure_semantics"`
	Limitations                 []string                         `json:"limitations,omitempty"`
}

type RecoveryUXContract struct {
	CurrentState        string   `json:"current_state"`
	RemediationClasses  []string `json:"remediation_classes,omitempty"`
	RequiredFields      []string `json:"required_fields,omitempty"`
	UnsafeRetryDenied   bool     `json:"unsafe_retry_denied"`
	UnsupportedExplicit bool     `json:"unsupported_explicit"`
	Limitations         []string `json:"limitations,omitempty"`
}

type AutomationActionMode struct {
	ActionMode            string `json:"action_mode"`
	CurrentState          string `json:"current_state"`
	MutatesCanonicalState bool   `json:"mutates_canonical_state"`
	AllowedInVal0         bool   `json:"allowed_in_val0"`
	GovernedElsewhere     bool   `json:"governed_elsewhere"`
	Limitation            string `json:"limitation"`
}

type ActionModeTaxonomy struct {
	CurrentState string                 `json:"current_state"`
	Items        []AutomationActionMode `json:"items,omitempty"`
	Limitations  []string               `json:"limitations,omitempty"`
}

func ProductionUsabilityVal0ConfigIntegrity() ConfigIntegrityContract {
	return ConfigIntegrityContract{
		CurrentState:              "config_integrity_contract_ready",
		SchemaVersion:             "point4.production_usability.val0.config_integrity.v1",
		SupportedCompatibility:    []string{ProductionUsabilityCompatibilityCurrent, ProductionUsabilityCompatibilityDeprecated, ProductionUsabilityCompatibilityMigrationRequired, ProductionUsabilityCompatibilityUnsupported, ProductionUsabilityCompatibilityUnknown},
		CurrentCompatibility:      ProductionUsabilityCompatibilityCurrent,
		UnknownFieldPolicies:      []string{ProductionUsabilityUnknownFieldReject, ProductionUsabilityUnknownFieldWarn, ProductionUsabilityUnknownFieldIgnoreExplicitlyAllowed},
		CurrentUnknownFieldPolicy: ProductionUsabilityUnknownFieldReject,
		ValidationStates:          []string{ProductionUsabilityValidationValid, ProductionUsabilityValidationInvalid, ProductionUsabilityValidationDegraded, ProductionUsabilityValidationUnsupported},
		CurrentValidationResult:   ProductionUsabilityValidationValid,
		MigrationWarningSemantics: []string{"warning_only_not_successful_migration", "migration_warning_requires_operator_visibility"},
		ConflictDetectionSemantics: []string{
			"conflicting_sources_mark_config_invalid_or_degraded",
			"effective_config_view_is_projection_only_not_canonical_truth",
		},
		EffectiveConfigView: "inspection_projection_only_not_canonical_truth",
		FailFastBootstrap:   "invalid_or_unsupported_config_blocks_active_bootstrap",
		Limitations: []string{
			"Val 0 defines config integrity semantics only; later Point 4 waves may add stricter factories and migration execution paths.",
		},
	}
}

func ProductionUsabilityVal0ExplainabilityContract() ExplainabilityPayloadContract {
	return ExplainabilityPayloadContract{
		CurrentState:                "explainability_contract_ready",
		RequiredFields:              []string{"reason_code", "severity", "human_message", "technical_detail", "policy_ref", "subject_ref", "evidence_refs", "next_step", "visibility_scope", "redaction_tier", "recovery_hint", "action_mode", "decision_priority"},
		SupportedSeverities:         []string{ProductionUsabilitySeverityCritical, ProductionUsabilitySeverityError, ProductionUsabilitySeverityWarning, ProductionUsabilitySeverityInfo},
		SupportedVisibilityScopes:   []string{ProductionUsabilityVisibilityInternalAdmin, ProductionUsabilityVisibilityOperator, ProductionUsabilityVisibilityDeveloper, ProductionUsabilityVisibilityPartner, ProductionUsabilityVisibilityPublicSafe},
		SupportedRedactionTiers:     []string{ProductionUsabilityRedactionNone, ProductionUsabilityRedactionLow, ProductionUsabilityRedactionMedium, ProductionUsabilityRedactionHigh, ProductionUsabilityRedactionPublicSafe},
		SupportedDecisionPriorities: []string{ProductionUsabilityDecisionBlocker, ProductionUsabilityDecisionUrgent, ProductionUsabilityDecisionNormal, ProductionUsabilityDecisionLow, ProductionUsabilityDecisionInformational},
		SupportedActionModes:        []string{ProductionUsabilityActionModeViewOnly, ProductionUsabilityActionModeExplain, ProductionUsabilityActionModePreview, ProductionUsabilityActionModeDryRun, ProductionUsabilityActionModeAuditOnly},
		SafeRedactionScopes:         []string{ProductionUsabilityVisibilityPartner, ProductionUsabilityVisibilityPublicSafe},
		TechnicalDetailBounded:      true,
		PreservesFailureSemantics:   true,
		Rules: []string{
			"public_safe_and_partner_scopes_must_explain_without_raw_sensitive_evidence",
			"human_message_cannot_hide_technical_reason",
			"technical_detail_never_overrides_visibility_or_redaction",
		},
		Limitations: []string{
			"Val 0 explainability is a bounded payload contract only; later Point 4 waves add richer UI, CLI, API, and support projections.",
		},
	}
}

func ProductionUsabilityVal0OperationalStatusModel() OperationalStatusModel {
	return OperationalStatusModel{
		CurrentState: "operational_status_model_ready",
		Items: []OperationalStatusDefinition{
			{Status: ProductionUsabilityStatusFresh, CurrentState: "status_defined", FreshnessIndicator: "freshness_timestamp_required", SourceProjectionID: "canonical_or_projection_source_required", LimitationMessage: "fresh does not imply universal health or completeness", EvidenceRefsPolicy: "required_when_decision_bound", CanonicalTruthDisclaimer: "projection_may_reflect_canonical_state_but_is_not_canonical_truth_itself", ClaimsCanonicalTruth: false},
			{Status: ProductionUsabilityStatusStale, CurrentState: "status_defined", FreshnessIndicator: "staleness_indicator_required", SourceProjectionID: "projection_source_required", LimitationMessage: "stale is not fresh and cannot claim canonical truth", EvidenceRefsPolicy: "required_when_available", CanonicalTruthDisclaimer: "stale_projection_is_non_canonical", ClaimsCanonicalTruth: false},
			{Status: ProductionUsabilityStatusPartial, CurrentState: "status_defined", FreshnessIndicator: "freshness_indicator_required", SourceProjectionID: "projection_source_required", LimitationMessage: "partial is not complete", EvidenceRefsPolicy: "required_when_available", CanonicalTruthDisclaimer: "partial_projection_is_non_canonical", ClaimsCanonicalTruth: false},
			{Status: ProductionUsabilityStatusDegraded, CurrentState: "status_defined", FreshnessIndicator: "degraded_indicator_required", SourceProjectionID: "projection_source_required", LimitationMessage: "degraded is not healthy", EvidenceRefsPolicy: "required_when_available", CanonicalTruthDisclaimer: "degraded_projection_is_non_canonical", ClaimsCanonicalTruth: false},
			{Status: ProductionUsabilityStatusUnavailable, CurrentState: "status_defined", FreshnessIndicator: "availability_indicator_required", SourceProjectionID: "projection_source_required", LimitationMessage: "unavailable is not unsupported", EvidenceRefsPolicy: "optional_when_source_absent", CanonicalTruthDisclaimer: "unavailable_projection_is_non_canonical", ClaimsCanonicalTruth: false},
			{Status: ProductionUsabilityStatusUnsupported, CurrentState: "status_defined", FreshnessIndicator: "unsupported_indicator_required", SourceProjectionID: "projection_source_required", LimitationMessage: "unsupported is explicit and not silent absence", EvidenceRefsPolicy: "optional_when_capability_not_supported", CanonicalTruthDisclaimer: "unsupported_projection_is_non_canonical", ClaimsCanonicalTruth: false},
		},
		RequiredDistinctions: []string{
			"stale_is_not_fresh",
			"partial_is_not_complete",
			"degraded_is_not_healthy",
			"unavailable_is_not_unsupported",
			"unsupported_is_explicit_not_silent",
			"no_projection_claims_canonical_truth",
		},
		Limitations: []string{
			"Val 0 status semantics are shared contracts for later UI, CLI, API, diagnostics, cache, and support projections; they do not create a new canonical truth source.",
		},
	}
}

func ProductionUsabilityVal0OperationContractModel() OperationContractModel {
	return OperationContractModel{
		CurrentState:           "operation_contract_ready",
		RequiredOperationTypes: []string{ProductionUsabilityOperationReadOnly, ProductionUsabilityOperationPreviewOnly, ProductionUsabilityOperationIdempotentMutation, ProductionUsabilityOperationNonIdempotentMutate, ProductionUsabilityOperationSideEffecting},
		SupportedRetrySafety:   []string{ProductionUsabilityRetrySafe, ProductionUsabilityRetryUnsafe, ProductionUsabilityRetryConditionallySafe},
		Items: []UsabilityOperation{
			{OperationName: "read_status_surface", OperationType: ProductionUsabilityOperationReadOnly, RetrySafety: ProductionUsabilityRetrySafe, SideEffects: nil, IdempotencyKeyRequired: false, SafeRetryGuidance: "safe_to_repeat_reads_when_request_context_is_unchanged", DoNotRetryReason: "not_applicable_retry_safe_read", PermissionScope: "viewer"},
			{OperationName: "preview_config_validation", OperationType: ProductionUsabilityOperationPreviewOnly, RetrySafety: ProductionUsabilityRetrySafe, SideEffects: nil, IdempotencyKeyRequired: false, SafeRetryGuidance: "safe_to_repeat_preview_without_mutating_state", DoNotRetryReason: "not_applicable_retry_safe_preview", PermissionScope: "operator"},
			{OperationName: "acknowledge_notification_projection", OperationType: ProductionUsabilityOperationIdempotentMutation, RetrySafety: ProductionUsabilityRetryConditionallySafe, SideEffects: []string{"projection_ack_state_update"}, IdempotencyKeyRequired: true, SafeRetryGuidance: "retry_only_with_same_idempotency_key_and_same_target", DoNotRetryReason: "unsafe_without_same_idempotency_key", PermissionScope: "operator"},
			{OperationName: "activate_break_glass_authority", OperationType: ProductionUsabilityOperationNonIdempotentMutate, RetrySafety: ProductionUsabilityRetryUnsafe, SideEffects: []string{"authority_grant_activation", "external_projection_update"}, IdempotencyKeyRequired: true, SafeRetryGuidance: "do_not_retry_without_manual_confirmation_of_effect", DoNotRetryReason: "may_duplicate_or_extend_a_high_risk_authority_grant", PermissionScope: "security_admin"},
			{OperationName: "dispatch_support_probe", OperationType: ProductionUsabilityOperationSideEffecting, RetrySafety: ProductionUsabilityRetryConditionallySafe, SideEffects: []string{"support_channel_notification", "diagnostic_collection_start"}, IdempotencyKeyRequired: true, SafeRetryGuidance: "retry_only_if_previous_dispatch_confirmed_not_delivered", DoNotRetryReason: "duplicate_dispatch_can_amplify_noise_or_repeat_side_effects", PermissionScope: "support"},
		},
		Rules: []string{
			"mutating_operations_are_not_assumed_retry_safe",
			"preview_and_audit_style_operations_never_mutate_canonical_state",
			"idempotency_key_is_required_where_replay_or_duplicate_risk_exists",
		},
		Limitations: []string{
			"Val 0 operation semantics classify retry and mutation posture only; later Point 4 waves add richer transport, backpressure, and support tooling behavior.",
		},
	}
}

func ProductionUsabilityVal0DecisionQualityContract() DecisionQualityContract {
	return DecisionQualityContract{
		CurrentState:          "decision_quality_contract_ready",
		DecisionPriorities:    []string{ProductionUsabilityDecisionBlocker, ProductionUsabilityDecisionUrgent, ProductionUsabilityDecisionNormal, ProductionUsabilityDecisionLow, ProductionUsabilityDecisionInformational},
		ActionRequiredClasses: []string{ProductionUsabilityNoAction, ProductionUsabilityOperatorAction, ProductionUsabilityAdminAction, ProductionUsabilityGovernanceAction, ProductionUsabilitySupportAction},
		RequiredFields:        []string{"decision_priority", "action_required", "blast_radius_hint", "affected_subjects_summary", "why_this_matters", "recommended_next_step"},
		AdvisoryOnly:          true,
		Limitations: []string{
			"Decision quality remains advisory and cannot approve, mutate, or override canonical workflow truth.",
		},
	}
}

func ProductionUsabilityVal0NotificationTaxonomyContract() NotificationTaxonomyContract {
	return NotificationTaxonomyContract{
		CurrentState:                 "notification_taxonomy_ready",
		Severities:                   []string{ProductionUsabilitySeverityCritical, ProductionUsabilitySeverityError, ProductionUsabilitySeverityWarning, ProductionUsabilitySeverityInfo},
		AcknowledgementStates:        []string{ProductionUsabilityAckUnacknowledged, ProductionUsabilityAckAcknowledged, ProductionUsabilityAckSuppressedDup, ProductionUsabilityAckResolved, ProductionUsabilityAckReopened},
		RequiredFields:               []string{"severity", "grouping_key", "duplicate_suppression_key", "acknowledgement_state", "escalation_policy_ref", "reopen_on_change"},
		AcknowledgementIsRemediation: false,
		ResolvedEqualsClosure:        false,
		CriticalSuppressionAllowed:   false,
		ReopenOnChangeExplicit:       true,
		Limitations: []string{
			"Notification acknowledgement and suppression remain operator-experience projections and never become canonical remediation or closure truth.",
		},
	}
}

func ProductionUsabilityVal0PermissionRedactionContract() PermissionRedactionContract {
	return PermissionRedactionContract{
		CurrentState:                "permission_redaction_contract_ready",
		SupportedEvidenceVisibility: []string{ProductionUsabilityEvidenceFull, ProductionUsabilityEvidenceMetadataOnly, ProductionUsabilityEvidenceRedacted, ProductionUsabilityEvidenceHidden},
		Items: []PermissionAwareExplanationRule{
			{VisibilityScope: ProductionUsabilityVisibilityInternalAdmin, CurrentState: "permission_rule_ready", AllowedDetailLevel: "full_diagnostic", RedactionTier: ProductionUsabilityRedactionNone, EvidenceVisibility: ProductionUsabilityEvidenceFull, SafeFallbackMessage: "Full diagnostic detail is available to internal administrators."},
			{VisibilityScope: ProductionUsabilityVisibilityOperator, CurrentState: "permission_rule_ready", AllowedDetailLevel: "operator_bounded", RedactionTier: ProductionUsabilityRedactionLow, EvidenceVisibility: ProductionUsabilityEvidenceMetadataOnly, SafeFallbackMessage: "Operator view is bounded to decision-relevant metadata and safe explanations."},
			{VisibilityScope: ProductionUsabilityVisibilityDeveloper, CurrentState: "permission_rule_ready", AllowedDetailLevel: "developer_bounded", RedactionTier: ProductionUsabilityRedactionMedium, EvidenceVisibility: ProductionUsabilityEvidenceMetadataOnly, SafeFallbackMessage: "Developer view stays bounded to technical reason without raw protected evidence."},
			{VisibilityScope: ProductionUsabilityVisibilityPartner, CurrentState: "permission_rule_ready", AllowedDetailLevel: "partner_safe", RedactionTier: ProductionUsabilityRedactionHigh, EvidenceVisibility: ProductionUsabilityEvidenceRedacted, SafeFallbackMessage: "Partner view is redacted and limited to safe bounded explanation."},
			{VisibilityScope: ProductionUsabilityVisibilityPublicSafe, CurrentState: "permission_rule_ready", AllowedDetailLevel: "public_safe", RedactionTier: ProductionUsabilityRedactionPublicSafe, EvidenceVisibility: ProductionUsabilityEvidenceHidden, SafeFallbackMessage: "Public-safe view discloses only bounded explanation and explicit limitations."},
		},
		TechnicalDetailBounded:    true,
		PreservesFailureSemantics: true,
		Limitations: []string{
			"Permission-aware explanation and redaction remain honest about limitations and never turn a fail or degraded result into a pass.",
		},
	}
}

func ProductionUsabilityVal0RecoveryUXContract() RecoveryUXContract {
	return RecoveryUXContract{
		CurrentState:        "recovery_contract_ready",
		RemediationClasses:  []string{ProductionUsabilityRemediationConfigFix, ProductionUsabilityRemediationPermissionFix, ProductionUsabilityRemediationEvidenceRefresh, ProductionUsabilityRemediationPolicyUpdate, ProductionUsabilityRemediationWaitAndRetry, ProductionUsabilityRemediationManualReview, ProductionUsabilityRemediationSupportEsc, ProductionUsabilityRemediationUnsupported},
		RequiredFields:      []string{"recovery_hint", "remediation_class", "safe_retry_guidance", "escalation_path", "rollback_hint", "do_not_retry_reason", "inspect_or_explain_command_ref"},
		UnsafeRetryDenied:   true,
		UnsupportedExplicit: true,
		Limitations: []string{
			"Recovery hints remain bounded to safe remediation guidance and cannot weaken policy, evidence, or governance discipline.",
		},
	}
}

func ProductionUsabilityVal0ActionModeTaxonomy() ActionModeTaxonomy {
	return ActionModeTaxonomy{
		CurrentState: "action_mode_taxonomy_ready",
		Items: []AutomationActionMode{
			{ActionMode: ProductionUsabilityActionModeViewOnly, CurrentState: "action_mode_ready", MutatesCanonicalState: false, AllowedInVal0: true, GovernedElsewhere: false, Limitation: "view_only never mutates canonical state"},
			{ActionMode: ProductionUsabilityActionModeExplain, CurrentState: "action_mode_ready", MutatesCanonicalState: false, AllowedInVal0: true, GovernedElsewhere: false, Limitation: "explain_only never mutates canonical state"},
			{ActionMode: ProductionUsabilityActionModePreview, CurrentState: "action_mode_ready", MutatesCanonicalState: false, AllowedInVal0: true, GovernedElsewhere: false, Limitation: "preview never mutates canonical state"},
			{ActionMode: ProductionUsabilityActionModeDryRun, CurrentState: "action_mode_ready", MutatesCanonicalState: false, AllowedInVal0: true, GovernedElsewhere: false, Limitation: "dry_run never mutates canonical state"},
			{ActionMode: ProductionUsabilityActionModeAuditOnly, CurrentState: "action_mode_ready", MutatesCanonicalState: false, AllowedInVal0: true, GovernedElsewhere: false, Limitation: "audit_only never mutates canonical state"},
			{ActionMode: ProductionUsabilityActionModeEnforce, CurrentState: "action_mode_declared_but_governed_elsewhere", MutatesCanonicalState: true, AllowedInVal0: false, GovernedElsewhere: true, Limitation: "enforce remains unavailable in Val 0 unless governed elsewhere"},
			{ActionMode: ProductionUsabilityActionModeMutate, CurrentState: "action_mode_declared_but_governed_elsewhere", MutatesCanonicalState: true, AllowedInVal0: false, GovernedElsewhere: true, Limitation: "mutate remains unavailable in Val 0 unless governed elsewhere"},
		},
		Limitations: []string{
			"Val 0 declares safe automation action semantics only; it does not introduce new automation authority.",
		},
	}
}

func EvaluateProductionUsabilityVal0ConfigIntegrityState(model ConfigIntegrityContract) string {
	if strings.TrimSpace(model.SchemaVersion) == "" || strings.TrimSpace(model.CurrentCompatibility) == "" || strings.TrimSpace(model.CurrentUnknownFieldPolicy) == "" || strings.TrimSpace(model.CurrentValidationResult) == "" || strings.TrimSpace(model.EffectiveConfigView) == "" || strings.TrimSpace(model.FailFastBootstrap) == "" {
		return ProductionUsabilityVal0ConfigIntegrityStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedCompatibility, ProductionUsabilityCompatibilityCurrent, ProductionUsabilityCompatibilityDeprecated, ProductionUsabilityCompatibilityMigrationRequired, ProductionUsabilityCompatibilityUnsupported, ProductionUsabilityCompatibilityUnknown) ||
		!containsExactTrimmedStringSet(model.UnknownFieldPolicies, ProductionUsabilityUnknownFieldReject, ProductionUsabilityUnknownFieldWarn, ProductionUsabilityUnknownFieldIgnoreExplicitlyAllowed) ||
		!containsExactTrimmedStringSet(model.ValidationStates, ProductionUsabilityValidationValid, ProductionUsabilityValidationInvalid, ProductionUsabilityValidationDegraded, ProductionUsabilityValidationUnsupported) ||
		len(model.MigrationWarningSemantics) == 0 ||
		len(model.ConflictDetectionSemantics) == 0 {
		return ProductionUsabilityVal0ConfigIntegrityStatePartial
	}
	currentCompatibility := strings.TrimSpace(model.CurrentCompatibility)
	currentValidation := strings.TrimSpace(model.CurrentValidationResult)
	if !containsTrimmedString(model.SupportedCompatibility, currentCompatibility) ||
		!containsTrimmedString(model.ValidationStates, currentValidation) {
		return ProductionUsabilityVal0ConfigIntegrityStatePartial
	}
	if currentCompatibility == ProductionUsabilityCompatibilityUnsupported ||
		currentCompatibility == ProductionUsabilityCompatibilityUnknown ||
		currentValidation == ProductionUsabilityValidationInvalid ||
		currentValidation == ProductionUsabilityValidationUnsupported {
		return ProductionUsabilityVal0ConfigIntegrityStatePartial
	}
	switch strings.TrimSpace(model.CurrentUnknownFieldPolicy) {
	case ProductionUsabilityUnknownFieldReject, ProductionUsabilityUnknownFieldWarn, ProductionUsabilityUnknownFieldIgnoreExplicitlyAllowed:
	default:
		return ProductionUsabilityVal0ConfigIntegrityStatePartial
	}
	if !strings.Contains(strings.TrimSpace(model.EffectiveConfigView), "projection_only") || !strings.Contains(strings.TrimSpace(model.FailFastBootstrap), "blocks_active_bootstrap") {
		return ProductionUsabilityVal0ConfigIntegrityStatePartial
	}
	return ProductionUsabilityVal0ConfigIntegrityStateActive
}

func EvaluateProductionUsabilityVal0ExplainabilityState(model ExplainabilityPayloadContract) string {
	if len(model.RequiredFields) == 0 {
		return ProductionUsabilityVal0ExplainabilityStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredFields, "reason_code", "severity", "human_message", "technical_detail", "policy_ref", "subject_ref", "evidence_refs", "next_step", "visibility_scope", "redaction_tier", "recovery_hint", "action_mode", "decision_priority") ||
		!containsExactTrimmedStringSet(model.SupportedSeverities, ProductionUsabilitySeverityCritical, ProductionUsabilitySeverityError, ProductionUsabilitySeverityWarning, ProductionUsabilitySeverityInfo) ||
		!containsExactTrimmedStringSet(model.SupportedVisibilityScopes, ProductionUsabilityVisibilityInternalAdmin, ProductionUsabilityVisibilityOperator, ProductionUsabilityVisibilityDeveloper, ProductionUsabilityVisibilityPartner, ProductionUsabilityVisibilityPublicSafe) ||
		!containsExactTrimmedStringSet(model.SupportedRedactionTiers, ProductionUsabilityRedactionNone, ProductionUsabilityRedactionLow, ProductionUsabilityRedactionMedium, ProductionUsabilityRedactionHigh, ProductionUsabilityRedactionPublicSafe) ||
		!containsAllTrimmedStrings(model.SafeRedactionScopes, ProductionUsabilityVisibilityPartner, ProductionUsabilityVisibilityPublicSafe) ||
		!model.TechnicalDetailBounded ||
		!model.PreservesFailureSemantics ||
		len(model.Rules) == 0 {
		return ProductionUsabilityVal0ExplainabilityStatePartial
	}
	return ProductionUsabilityVal0ExplainabilityStateActive
}

func EvaluateProductionUsabilityVal0StatusModelState(model OperationalStatusModel) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityVal0StatusModelStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredDistinctions, "stale_is_not_fresh", "partial_is_not_complete", "degraded_is_not_healthy", "unavailable_is_not_unsupported", "unsupported_is_explicit_not_silent", "no_projection_claims_canonical_truth") {
		return ProductionUsabilityVal0StatusModelStatePartial
	}
	expectedStatuses := map[string]struct{}{
		ProductionUsabilityStatusFresh:       {},
		ProductionUsabilityStatusStale:       {},
		ProductionUsabilityStatusPartial:     {},
		ProductionUsabilityStatusDegraded:    {},
		ProductionUsabilityStatusUnavailable: {},
		ProductionUsabilityStatusUnsupported: {},
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		status := strings.TrimSpace(item.Status)
		if status == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.FreshnessIndicator) == "" || strings.TrimSpace(item.SourceProjectionID) == "" || strings.TrimSpace(item.LimitationMessage) == "" || strings.TrimSpace(item.EvidenceRefsPolicy) == "" || strings.TrimSpace(item.CanonicalTruthDisclaimer) == "" {
			return ProductionUsabilityVal0StatusModelStateIncomplete
		}
		if _, ok := expectedStatuses[status]; !ok {
			return ProductionUsabilityVal0StatusModelStatePartial
		}
		if _, duplicate := seen[status]; duplicate {
			return ProductionUsabilityVal0StatusModelStatePartial
		}
		seen[status] = struct{}{}
		if item.ClaimsCanonicalTruth {
			return ProductionUsabilityVal0StatusModelStatePartial
		}
	}
	if len(seen) != len(expectedStatuses) {
		return ProductionUsabilityVal0StatusModelStatePartial
	}
	return ProductionUsabilityVal0StatusModelStateActive
}

func EvaluateProductionUsabilityVal0OperationContractState(model OperationContractModel) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityVal0OperationContractStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredOperationTypes, ProductionUsabilityOperationReadOnly, ProductionUsabilityOperationPreviewOnly, ProductionUsabilityOperationIdempotentMutation, ProductionUsabilityOperationNonIdempotentMutate, ProductionUsabilityOperationSideEffecting) ||
		!containsExactTrimmedStringSet(model.SupportedRetrySafety, ProductionUsabilityRetrySafe, ProductionUsabilityRetryUnsafe, ProductionUsabilityRetryConditionallySafe) ||
		len(model.Rules) == 0 {
		return ProductionUsabilityVal0OperationContractStatePartial
	}
	expectedTypes := map[string]struct{}{
		ProductionUsabilityOperationReadOnly:            {},
		ProductionUsabilityOperationPreviewOnly:         {},
		ProductionUsabilityOperationIdempotentMutation:  {},
		ProductionUsabilityOperationNonIdempotentMutate: {},
		ProductionUsabilityOperationSideEffecting:       {},
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		opType := strings.TrimSpace(item.OperationType)
		if strings.TrimSpace(item.OperationName) == "" || opType == "" || strings.TrimSpace(item.RetrySafety) == "" || strings.TrimSpace(item.SafeRetryGuidance) == "" || strings.TrimSpace(item.DoNotRetryReason) == "" || strings.TrimSpace(item.PermissionScope) == "" {
			return ProductionUsabilityVal0OperationContractStateIncomplete
		}
		if _, ok := expectedTypes[opType]; !ok {
			return ProductionUsabilityVal0OperationContractStatePartial
		}
		if _, duplicate := seen[opType]; duplicate {
			return ProductionUsabilityVal0OperationContractStatePartial
		}
		seen[opType] = struct{}{}
		if (opType == ProductionUsabilityOperationNonIdempotentMutate || opType == ProductionUsabilityOperationSideEffecting) && strings.TrimSpace(item.RetrySafety) == ProductionUsabilityRetrySafe {
			return ProductionUsabilityVal0OperationContractStatePartial
		}
		if opType == ProductionUsabilityOperationPreviewOnly && len(item.SideEffects) > 0 {
			return ProductionUsabilityVal0OperationContractStatePartial
		}
	}
	if len(seen) != len(expectedTypes) {
		return ProductionUsabilityVal0OperationContractStatePartial
	}
	return ProductionUsabilityVal0OperationContractStateActive
}

func EvaluateProductionUsabilityVal0DecisionQualityState(model DecisionQualityContract) string {
	if len(model.RequiredFields) == 0 {
		return ProductionUsabilityVal0DecisionQualityStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.DecisionPriorities, ProductionUsabilityDecisionBlocker, ProductionUsabilityDecisionUrgent, ProductionUsabilityDecisionNormal, ProductionUsabilityDecisionLow, ProductionUsabilityDecisionInformational) ||
		!containsExactTrimmedStringSet(model.ActionRequiredClasses, ProductionUsabilityNoAction, ProductionUsabilityOperatorAction, ProductionUsabilityAdminAction, ProductionUsabilityGovernanceAction, ProductionUsabilitySupportAction) ||
		!containsExactTrimmedStringSet(model.RequiredFields, "decision_priority", "action_required", "blast_radius_hint", "affected_subjects_summary", "why_this_matters", "recommended_next_step") ||
		!model.AdvisoryOnly {
		return ProductionUsabilityVal0DecisionQualityStatePartial
	}
	return ProductionUsabilityVal0DecisionQualityStateActive
}

func EvaluateProductionUsabilityVal0NotificationState(model NotificationTaxonomyContract) string {
	if len(model.RequiredFields) == 0 {
		return ProductionUsabilityVal0NotificationStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.Severities, ProductionUsabilitySeverityCritical, ProductionUsabilitySeverityError, ProductionUsabilitySeverityWarning, ProductionUsabilitySeverityInfo) ||
		!containsExactTrimmedStringSet(model.AcknowledgementStates, ProductionUsabilityAckUnacknowledged, ProductionUsabilityAckAcknowledged, ProductionUsabilityAckSuppressedDup, ProductionUsabilityAckResolved, ProductionUsabilityAckReopened) ||
		!containsExactTrimmedStringSet(model.RequiredFields, "severity", "grouping_key", "duplicate_suppression_key", "acknowledgement_state", "escalation_policy_ref", "reopen_on_change") ||
		model.AcknowledgementIsRemediation ||
		model.ResolvedEqualsClosure ||
		model.CriticalSuppressionAllowed ||
		!model.ReopenOnChangeExplicit {
		return ProductionUsabilityVal0NotificationStatePartial
	}
	return ProductionUsabilityVal0NotificationStateActive
}

func EvaluateProductionUsabilityVal0PermissionRedactionState(model PermissionRedactionContract) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityVal0PermissionRedactionStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedEvidenceVisibility, ProductionUsabilityEvidenceFull, ProductionUsabilityEvidenceMetadataOnly, ProductionUsabilityEvidenceRedacted, ProductionUsabilityEvidenceHidden) ||
		!model.TechnicalDetailBounded ||
		!model.PreservesFailureSemantics {
		return ProductionUsabilityVal0PermissionRedactionStatePartial
	}
	expectedScopes := map[string]struct{}{
		ProductionUsabilityVisibilityInternalAdmin: {},
		ProductionUsabilityVisibilityOperator:      {},
		ProductionUsabilityVisibilityDeveloper:     {},
		ProductionUsabilityVisibilityPartner:       {},
		ProductionUsabilityVisibilityPublicSafe:    {},
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		scope := strings.TrimSpace(item.VisibilityScope)
		evidenceVisibility := strings.TrimSpace(item.EvidenceVisibility)
		if scope == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.AllowedDetailLevel) == "" || strings.TrimSpace(item.RedactionTier) == "" || evidenceVisibility == "" || strings.TrimSpace(item.SafeFallbackMessage) == "" {
			return ProductionUsabilityVal0PermissionRedactionStateIncomplete
		}
		if _, ok := expectedScopes[scope]; !ok {
			return ProductionUsabilityVal0PermissionRedactionStatePartial
		}
		if _, duplicate := seen[scope]; duplicate {
			return ProductionUsabilityVal0PermissionRedactionStatePartial
		}
		seen[scope] = struct{}{}
		if !containsTrimmedString(model.SupportedEvidenceVisibility, evidenceVisibility) {
			return ProductionUsabilityVal0PermissionRedactionStatePartial
		}
		if (scope == ProductionUsabilityVisibilityPartner || scope == ProductionUsabilityVisibilityPublicSafe) && evidenceVisibility == ProductionUsabilityEvidenceFull {
			return ProductionUsabilityVal0PermissionRedactionStatePartial
		}
	}
	if len(seen) != len(expectedScopes) {
		return ProductionUsabilityVal0PermissionRedactionStatePartial
	}
	return ProductionUsabilityVal0PermissionRedactionStateActive
}

func EvaluateProductionUsabilityVal0RecoveryState(model RecoveryUXContract) string {
	if len(model.RequiredFields) == 0 {
		return ProductionUsabilityVal0RecoveryStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RemediationClasses, ProductionUsabilityRemediationConfigFix, ProductionUsabilityRemediationPermissionFix, ProductionUsabilityRemediationEvidenceRefresh, ProductionUsabilityRemediationPolicyUpdate, ProductionUsabilityRemediationWaitAndRetry, ProductionUsabilityRemediationManualReview, ProductionUsabilityRemediationSupportEsc, ProductionUsabilityRemediationUnsupported) ||
		!containsExactTrimmedStringSet(model.RequiredFields, "recovery_hint", "remediation_class", "safe_retry_guidance", "escalation_path", "rollback_hint", "do_not_retry_reason", "inspect_or_explain_command_ref") ||
		!model.UnsafeRetryDenied ||
		!model.UnsupportedExplicit {
		return ProductionUsabilityVal0RecoveryStatePartial
	}
	return ProductionUsabilityVal0RecoveryStateActive
}

func EvaluateProductionUsabilityVal0ActionModeState(model ActionModeTaxonomy) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityVal0ActionModeStateIncomplete
	}
	expectedModes := map[string]struct{}{
		ProductionUsabilityActionModeViewOnly:  {},
		ProductionUsabilityActionModeExplain:   {},
		ProductionUsabilityActionModePreview:   {},
		ProductionUsabilityActionModeDryRun:    {},
		ProductionUsabilityActionModeAuditOnly: {},
		ProductionUsabilityActionModeEnforce:   {},
		ProductionUsabilityActionModeMutate:    {},
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		mode := strings.TrimSpace(item.ActionMode)
		if mode == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.Limitation) == "" {
			return ProductionUsabilityVal0ActionModeStateIncomplete
		}
		if _, ok := expectedModes[mode]; !ok {
			return ProductionUsabilityVal0ActionModeStatePartial
		}
		if _, duplicate := seen[mode]; duplicate {
			return ProductionUsabilityVal0ActionModeStatePartial
		}
		seen[mode] = struct{}{}
		switch mode {
		case ProductionUsabilityActionModeViewOnly, ProductionUsabilityActionModeExplain, ProductionUsabilityActionModePreview, ProductionUsabilityActionModeDryRun, ProductionUsabilityActionModeAuditOnly:
			if item.MutatesCanonicalState || !item.AllowedInVal0 {
				return ProductionUsabilityVal0ActionModeStatePartial
			}
		case ProductionUsabilityActionModeEnforce, ProductionUsabilityActionModeMutate:
			if item.AllowedInVal0 || !item.GovernedElsewhere || !item.MutatesCanonicalState {
				return ProductionUsabilityVal0ActionModeStatePartial
			}
		}
	}
	if len(seen) != len(expectedModes) {
		return ProductionUsabilityVal0ActionModeStatePartial
	}
	return ProductionUsabilityVal0ActionModeStateActive
}

func EvaluateProductionUsabilityVal0State(point3DependencyState, configIntegrityState, explainabilityState, statusModelState, operationContractState, decisionQualityState, notificationState, permissionRedactionState, recoveryState, actionModeState string) string {
	if strings.TrimSpace(point3DependencyState) != workflowauthority.EnterpriseWorkflowAuthorityValDStateActive {
		return ProductionUsabilityVal0StateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(configIntegrityState),
		strings.TrimSpace(explainabilityState),
		strings.TrimSpace(statusModelState),
		strings.TrimSpace(operationContractState),
		strings.TrimSpace(decisionQualityState),
		strings.TrimSpace(notificationState),
		strings.TrimSpace(permissionRedactionState),
		strings.TrimSpace(recoveryState),
		strings.TrimSpace(actionModeState),
	} {
		switch state {
		case ProductionUsabilityVal0ConfigIntegrityStateActive,
			ProductionUsabilityVal0ExplainabilityStateActive,
			ProductionUsabilityVal0StatusModelStateActive,
			ProductionUsabilityVal0OperationContractStateActive,
			ProductionUsabilityVal0DecisionQualityStateActive,
			ProductionUsabilityVal0NotificationStateActive,
			ProductionUsabilityVal0PermissionRedactionStateActive,
			ProductionUsabilityVal0RecoveryStateActive,
			ProductionUsabilityVal0ActionModeStateActive:
		case ProductionUsabilityVal0ConfigIntegrityStatePartial,
			ProductionUsabilityVal0ExplainabilityStatePartial,
			ProductionUsabilityVal0StatusModelStatePartial,
			ProductionUsabilityVal0OperationContractStatePartial,
			ProductionUsabilityVal0DecisionQualityStatePartial,
			ProductionUsabilityVal0NotificationStatePartial,
			ProductionUsabilityVal0PermissionRedactionStatePartial,
			ProductionUsabilityVal0RecoveryStatePartial,
			ProductionUsabilityVal0ActionModeStatePartial:
			hasPartial = true
		default:
			return ProductionUsabilityVal0StateIncomplete
		}
	}
	if hasPartial {
		return ProductionUsabilityVal0StateSubstantial
	}
	return ProductionUsabilityVal0StateActive
}

func EvaluateProductionUsabilityVal0ProofsState(point3DependencyState, configIntegrityState, explainabilityState, statusModelState, operationContractState, decisionQualityState, notificationState, permissionRedactionState, recoveryState, actionModeState string, surfaceRefs, evidenceRefs, limitations, whyPoint4NotPass []string) string {
	baseState := EvaluateProductionUsabilityVal0State(point3DependencyState, configIntegrityState, explainabilityState, statusModelState, operationContractState, decisionQualityState, notificationState, permissionRedactionState, recoveryState, actionModeState)
	if len(surfaceRefs) < 10 || len(evidenceRefs) < 6 || len(limitations) == 0 || len(whyPoint4NotPass) == 0 {
		if baseState == ProductionUsabilityVal0StateActive {
			return ProductionUsabilityVal0StateSubstantial
		}
		return baseState
	}
	return baseState
}

func ProductionUsabilityPoint4DocumentationRefs() []string {
	return []string{
		"docs/production-usability-operability-recovery-val0-core.md",
	}
}

func containsExactTrimmedStringSet(values []string, expected ...string) bool {
	if len(values) != len(expected) {
		return false
	}
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			return false
		}
		if _, duplicate := seen[trimmed]; duplicate {
			return false
		}
		seen[trimmed] = struct{}{}
	}
	for _, item := range expected {
		if _, ok := seen[strings.TrimSpace(item)]; !ok {
			return false
		}
	}
	return true
}

func containsAllTrimmedStrings(values []string, expected ...string) bool {
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue
		}
		seen[trimmed] = struct{}{}
	}
	for _, item := range expected {
		if _, ok := seen[strings.TrimSpace(item)]; !ok {
			return false
		}
	}
	return true
}

func containsTrimmedString(values []string, expected string) bool {
	expected = strings.TrimSpace(expected)
	if expected == "" {
		return false
	}
	for _, value := range values {
		if strings.TrimSpace(value) == expected {
			return true
		}
	}
	return false
}

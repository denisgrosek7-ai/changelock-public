package operability

import "strings"

const (
	ProductionUsabilityValAConfigFactoryStateActive     = "production_usability_vala_config_factory_active"
	ProductionUsabilityValAConfigFactoryStatePartial    = "production_usability_vala_config_factory_partial"
	ProductionUsabilityValAConfigFactoryStateIncomplete = "production_usability_vala_config_factory_incomplete"

	ProductionUsabilityValABootstrapValidationStateActive     = "production_usability_vala_bootstrap_validation_active"
	ProductionUsabilityValABootstrapValidationStatePartial    = "production_usability_vala_bootstrap_validation_partial"
	ProductionUsabilityValABootstrapValidationStateIncomplete = "production_usability_vala_bootstrap_validation_incomplete"

	ProductionUsabilityValAPolicySchemaStateActive     = "production_usability_vala_policy_schema_active"
	ProductionUsabilityValAPolicySchemaStatePartial    = "production_usability_vala_policy_schema_partial"
	ProductionUsabilityValAPolicySchemaStateIncomplete = "production_usability_vala_policy_schema_incomplete"

	ProductionUsabilityValAEffectiveConfigStateActive     = "production_usability_vala_effective_config_active"
	ProductionUsabilityValAEffectiveConfigStatePartial    = "production_usability_vala_effective_config_partial"
	ProductionUsabilityValAEffectiveConfigStateIncomplete = "production_usability_vala_effective_config_incomplete"

	ProductionUsabilityValARejectionLayerStateActive     = "production_usability_vala_rejection_layer_active"
	ProductionUsabilityValARejectionLayerStatePartial    = "production_usability_vala_rejection_layer_partial"
	ProductionUsabilityValARejectionLayerStateIncomplete = "production_usability_vala_rejection_layer_incomplete"

	ProductionUsabilityValADryRunStateActive     = "production_usability_vala_policy_dry_run_active"
	ProductionUsabilityValADryRunStatePartial    = "production_usability_vala_policy_dry_run_partial"
	ProductionUsabilityValADryRunStateIncomplete = "production_usability_vala_policy_dry_run_incomplete"

	ProductionUsabilityValAExplainStateActive     = "production_usability_vala_permission_explain_active"
	ProductionUsabilityValAExplainStatePartial    = "production_usability_vala_permission_explain_partial"
	ProductionUsabilityValAExplainStateIncomplete = "production_usability_vala_permission_explain_incomplete"

	ProductionUsabilityValARecoveryGuidanceStateActive     = "production_usability_vala_recovery_guidance_active"
	ProductionUsabilityValARecoveryGuidanceStatePartial    = "production_usability_vala_recovery_guidance_partial"
	ProductionUsabilityValARecoveryGuidanceStateIncomplete = "production_usability_vala_recovery_guidance_incomplete"

	ProductionUsabilityValAFirstRunStateActive     = "production_usability_vala_first_run_bootstrap_active"
	ProductionUsabilityValAFirstRunStatePartial    = "production_usability_vala_first_run_bootstrap_partial"
	ProductionUsabilityValAFirstRunStateIncomplete = "production_usability_vala_first_run_bootstrap_incomplete"

	ProductionUsabilityValAUpgradePreviewStateActive     = "production_usability_vala_upgrade_impact_preview_active"
	ProductionUsabilityValAUpgradePreviewStatePartial    = "production_usability_vala_upgrade_impact_preview_partial"
	ProductionUsabilityValAUpgradePreviewStateIncomplete = "production_usability_vala_upgrade_impact_preview_incomplete"

	ProductionUsabilityValAStateIncomplete  = "production_usability_vala_incomplete"
	ProductionUsabilityValAStateSubstantial = "production_usability_vala_substantially_ready"
	ProductionUsabilityValAStateActive      = "production_usability_vala_active"

	ProductionUsabilityBootstrapAllowed         = "allowed"
	ProductionUsabilityBootstrapBlocked         = "blocked"
	ProductionUsabilityBootstrapDegradedAllowed = "degraded_allowed"
	ProductionUsabilityBootstrapUnsupported     = "unsupported"
)

type SchemaStrictConfigFactory struct {
	CurrentState                  string   `json:"current_state"`
	SchemaVersion                 string   `json:"schema_version"`
	SupportedSchemaVersions       []string `json:"supported_schema_versions,omitempty"`
	CurrentCompatibility          string   `json:"current_compatibility_status"`
	CurrentUnknownFieldPolicy     string   `json:"current_unknown_field_policy"`
	UnknownFieldsDetected         []string `json:"unknown_fields_detected,omitempty"`
	ExplicitUnknownFieldAllowance bool     `json:"explicit_unknown_field_allowance"`
	RequiredFieldValidation       bool     `json:"required_field_validation"`
	TypeShapeValidation           bool     `json:"type_shape_validation"`
	ConflictDetection             bool     `json:"conflict_detection"`
	CompatibilityWarnings         []string `json:"compatibility_warnings,omitempty"`
	MigrationWarnings             []string `json:"migration_warnings,omitempty"`
	MigrationCompleted            bool     `json:"migration_completed"`
	CurrentValidationResult       string   `json:"current_validation_result"`
	FailFastBootstrap             bool     `json:"fail_fast_bootstrap"`
	InspectionMode                string   `json:"inspection_mode"`
	DefaultsApplied               []string `json:"defaults_applied,omitempty"`
	UserProvidedFields            []string `json:"user_provided_fields,omitempty"`
	DefaultsDeclaredExplicitly    bool     `json:"defaults_declared_explicitly"`
	DefaultsSurfacedInInspection  bool     `json:"defaults_surfaced_in_inspection"`
	DefaultsDistinguishable       bool     `json:"defaults_distinguishable_from_user_values"`
	SecretsRedacted               bool     `json:"secrets_redacted"`
	Limitations                   []string `json:"limitations,omitempty"`
}

type BootstrapValidationCore struct {
	CurrentState            string   `json:"current_state"`
	ConfigValidationResult  string   `json:"config_validation_result"`
	PolicyValidationResult  string   `json:"policy_validation_result"`
	BootstrapDisposition    string   `json:"bootstrap_disposition"`
	BlockingReasonCodes     []string `json:"blocking_reason_codes,omitempty"`
	WarningCodes            []string `json:"warning_codes,omitempty"`
	AggregateReportingSafe  bool     `json:"aggregate_reporting_safe"`
	RejectionFields         []string `json:"rejection_fields,omitempty"`
	RecoveryHints           []string `json:"recovery_hints,omitempty"`
	InspectCommands         []string `json:"inspect_commands,omitempty"`
	DegradedBoundaries      []string `json:"degraded_boundaries,omitempty"`
	DegradedAllowedExplicit bool     `json:"degraded_allowed_explicit"`
	ProjectionDisclaimer    string   `json:"projection_disclaimer"`
	Limitations             []string `json:"limitations,omitempty"`
}

type PolicySchemaDiscipline struct {
	CurrentState                  string   `json:"current_state"`
	PolicySchemaVersion           string   `json:"policy_schema_version"`
	SupportedSchemaVersions       []string `json:"supported_schema_versions,omitempty"`
	CurrentCompatibility          string   `json:"current_compatibility_status"`
	CurrentValidationResult       string   `json:"current_validation_result"`
	CurrentUnknownFieldPolicy     string   `json:"current_unknown_field_policy"`
	UnknownFieldsDetected         []string `json:"unknown_fields_detected,omitempty"`
	ExplicitUnknownFieldAllowance bool     `json:"explicit_unknown_field_allowance"`
	DeprecatedWarnings            []string `json:"deprecated_warnings,omitempty"`
	MigrationWarnings             []string `json:"migration_warnings,omitempty"`
	MigrationCompleted            bool     `json:"migration_completed"`
	EffectiveInspectionView       string   `json:"effective_inspection_view"`
	Limitations                   []string `json:"limitations,omitempty"`
}

type EffectiveConfigInspection struct {
	CurrentState               string   `json:"current_state"`
	SchemaVersion              string   `json:"schema_version"`
	CompatibilityStatus        string   `json:"compatibility_status"`
	ValidationResult           string   `json:"validation_result"`
	DefaultsApplied            []string `json:"defaults_applied,omitempty"`
	UserProvidedFields         []string `json:"user_provided_fields,omitempty"`
	RejectedUnknownFields      []string `json:"rejected_unknown_fields,omitempty"`
	Warnings                   []string `json:"warnings,omitempty"`
	Conflicts                  []string `json:"conflicts,omitempty"`
	MigrationWarnings          []string `json:"migration_warnings,omitempty"`
	PolicySchemaSummary        string   `json:"policy_schema_summary"`
	SourceProjectionDisclaimer string   `json:"source_projection_disclaimer"`
	GeneratedAtIndicator       string   `json:"generated_at_indicator"`
	LimitationNotes            []string `json:"limitation_notes,omitempty"`
	RedactedFields             []string `json:"redacted_fields,omitempty"`
	PermissionAware            bool     `json:"permission_aware"`
	SecretsExposed             bool     `json:"secrets_exposed"`
}

type HumanReadableRejectionLayer struct {
	CurrentState                string   `json:"current_state"`
	RequiredFields              []string `json:"required_fields,omitempty"`
	SupportedVisibilityScopes   []string `json:"supported_visibility_scopes,omitempty"`
	SupportedRedactionTiers     []string `json:"supported_redaction_tiers,omitempty"`
	SupportedActionModes        []string `json:"supported_action_modes,omitempty"`
	SupportedDecisionPriorities []string `json:"supported_decision_priorities,omitempty"`
	TechnicalDetailScopes       []string `json:"technical_detail_scopes,omitempty"`
	RestrictedScopes            []string `json:"restricted_scopes,omitempty"`
	SecretsRedacted             bool     `json:"secrets_redacted"`
	TechnicalDetailBounded      bool     `json:"technical_detail_bounded"`
	FailureDowngradedToWarning  bool     `json:"failure_downgraded_to_warning"`
	Limitations                 []string `json:"limitations,omitempty"`
}

type PolicyDryRunAuditFlow struct {
	CurrentState                 string   `json:"current_state"`
	SupportedActionModes         []string `json:"supported_action_modes,omitempty"`
	PreviewAccepted              []string `json:"preview_accepted,omitempty"`
	PreviewRejected              []string `json:"preview_rejected,omitempty"`
	BlockingRules                []string `json:"blocking_rules,omitempty"`
	NonBlockingWarnings          []string `json:"non_blocking_warnings,omitempty"`
	PermissionAwareOutput        bool     `json:"permission_aware_output"`
	RecoveryHints                []string `json:"recovery_hints,omitempty"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
	MutatesCanonicalState        bool     `json:"mutates_canonical_state"`
	AuditOnlyImpliesApproval     bool     `json:"audit_only_implies_approval"`
	DryRunSuccessImpliesActivate bool     `json:"dry_run_success_implies_activation"`
	Limitations                  []string `json:"limitations,omitempty"`
}

type ExplainOutputVariant struct {
	VisibilityScope      string `json:"visibility_scope"`
	CurrentState         string `json:"current_state"`
	DetailLevel          string `json:"detail_level"`
	EvidenceVisibility   string `json:"evidence_visibility"`
	RedactionTier        string `json:"redaction_tier"`
	HonestLimitation     string `json:"honest_limitation"`
	SecretsRedacted      bool   `json:"secrets_redacted"`
	HiddenEvidenceMarker string `json:"hidden_evidence_marker"`
}

type PermissionAwareExplainOutputs struct {
	CurrentState                string                 `json:"current_state"`
	Items                       []ExplainOutputVariant `json:"items,omitempty"`
	SupportedEvidenceVisibility []string               `json:"supported_evidence_visibility,omitempty"`
	PreservesFailureSemantics   bool                   `json:"preserves_failure_semantics"`
	Limitations                 []string               `json:"limitations,omitempty"`
}

type RecoveryGuidanceItem struct {
	FailureClass                   string `json:"failure_class"`
	CurrentState                   string `json:"current_state"`
	RemediationClass               string `json:"remediation_class"`
	RecoveryHint                   string `json:"recovery_hint"`
	SafeRetryGuidance              string `json:"safe_retry_guidance"`
	DoNotRetryReason               string `json:"do_not_retry_reason"`
	InspectOrExplainCommand        string `json:"inspect_or_explain_command"`
	EscalationPath                 string `json:"escalation_path"`
	UnsafeRetrySuggested           bool   `json:"unsafe_retry_suggested"`
	PolicyBypassSuggested          bool   `json:"policy_bypass_suggested"`
	CanonicalEvidenceEditSuggested bool   `json:"canonical_evidence_edit_suggested"`
}

type RecoveryGuidanceCore struct {
	CurrentState        string                 `json:"current_state"`
	Items               []RecoveryGuidanceItem `json:"items,omitempty"`
	UnsupportedExplicit bool                   `json:"unsupported_explicit"`
	Limitations         []string               `json:"limitations,omitempty"`
}

type FirstRunSafeBootstrap struct {
	CurrentState               string   `json:"current_state"`
	MinimalSafeShapeFields     []string `json:"minimal_safe_shape_fields,omitempty"`
	MissingReadinessFields     []string `json:"missing_readiness_fields,omitempty"`
	SampleConfigDetected       bool     `json:"sample_config_detected"`
	SampleMarkedNonProduction  bool     `json:"sample_marked_non_production"`
	AutoEnablesProduction      bool     `json:"auto_enables_production"`
	FakeDemoEvidencePresent    bool     `json:"fake_demo_evidence_present"`
	NextSteps                  []string `json:"next_steps,omitempty"`
	NonMutatingPreview         bool     `json:"non_mutating_preview"`
	BootstrapPathValid         bool     `json:"bootstrap_path_valid"`
	ClaimsProductionCompletion bool     `json:"claims_production_completion"`
}

type UpgradeImpactPreview struct {
	CurrentState             string   `json:"current_state"`
	CurrentSchemaVersion     string   `json:"current_schema_version"`
	TargetSchemaVersion      string   `json:"target_schema_version"`
	KnownTargetSchemas       []string `json:"known_target_schemas,omitempty"`
	CompatibilityStatus      string   `json:"compatibility_status"`
	DeprecatedFields         []string `json:"deprecated_fields,omitempty"`
	RemovedUnsupportedFields []string `json:"removed_unsupported_fields,omitempty"`
	MigrationRequiredItems   []string `json:"migration_required_items,omitempty"`
	Warnings                 []string `json:"warnings,omitempty"`
	BlockingIssues           []string `json:"blocking_issues,omitempty"`
	RollbackPerspective      string   `json:"rollback_perspective"`
	RollbackAppearsPossible  bool     `json:"rollback_appears_possible"`
	LimitationDisclaimer     string   `json:"limitation_disclaimer"`
	MutatesConfig            bool     `json:"mutates_config"`
}

func productionUsabilityValAConfigSchemaVersions() []string {
	return []string{
		"point4.production_usability.vala.config.v1",
		"point4.production_usability.vala.config.v0",
	}
}

func productionUsabilityValAPolicySchemaVersions() []string {
	return []string{
		"point4.production_usability.vala.policy.v1",
		"point4.production_usability.vala.policy.v0",
	}
}

func productionUsabilityValARequiredExplainFields() []string {
	return []string{"reason_code", "severity", "human_message", "technical_detail", "policy_ref", "subject_ref", "evidence_refs", "next_step", "recovery_hint", "visibility_scope", "redaction_tier", "action_mode", "decision_priority"}
}

func productionUsabilityValAExplainScopes() []string {
	return []string{ProductionUsabilityVisibilityInternalAdmin, ProductionUsabilityVisibilityOperator, ProductionUsabilityVisibilityDeveloper, ProductionUsabilityVisibilityPartner, ProductionUsabilityVisibilityPublicSafe}
}

func productionUsabilityValARecoveryFailureClasses() []string {
	return []string{
		"missing_required_field",
		"unknown_field",
		"unsupported_schema_version",
		"deprecated_schema_version",
		"migration_required",
		"conflicting_values",
		"invalid_field_shape",
		"permission_insufficient_for_detailed_explain",
		"dry_run_success_activation_not_performed",
	}
}

func ProductionUsabilityValAConfigFactory() SchemaStrictConfigFactory {
	return SchemaStrictConfigFactory{
		CurrentState:                  "schema_strict_config_factory_ready",
		SchemaVersion:                 "point4.production_usability.vala.config.v1",
		SupportedSchemaVersions:       productionUsabilityValAConfigSchemaVersions(),
		CurrentCompatibility:          ProductionUsabilityCompatibilityCurrent,
		CurrentUnknownFieldPolicy:     ProductionUsabilityUnknownFieldReject,
		UnknownFieldsDetected:         nil,
		ExplicitUnknownFieldAllowance: false,
		RequiredFieldValidation:       true,
		TypeShapeValidation:           true,
		ConflictDetection:             true,
		CompatibilityWarnings:         nil,
		MigrationWarnings:             nil,
		MigrationCompleted:            false,
		CurrentValidationResult:       ProductionUsabilityValidationValid,
		FailFastBootstrap:             true,
		InspectionMode:                "projection_only_non_mutating_effective_config_view",
		DefaultsApplied:               []string{"audit_store", "request_timeout"},
		UserProvidedFields:            []string{"schema_version", "tenant_id", "environment"},
		DefaultsDeclaredExplicitly:    true,
		DefaultsSurfacedInInspection:  true,
		DefaultsDistinguishable:       true,
		SecretsRedacted:               true,
		Limitations: []string{
			"Val A implements schema-strict config validation and projection surfaces only; later waves add wider operator UX and lifecycle flows.",
		},
	}
}

func ProductionUsabilityValABootstrapValidation() BootstrapValidationCore {
	return BootstrapValidationCore{
		CurrentState:            "fail_fast_bootstrap_validation_ready",
		ConfigValidationResult:  ProductionUsabilityValidationValid,
		PolicyValidationResult:  ProductionUsabilityValidationValid,
		BootstrapDisposition:    ProductionUsabilityBootstrapAllowed,
		BlockingReasonCodes:     nil,
		WarningCodes:            []string{"deprecated_schema_surfaces_warning_only_when_present"},
		AggregateReportingSafe:  true,
		RejectionFields:         productionUsabilityValARequiredExplainFields(),
		RecoveryHints:           []string{"run_changelock_config_inspect", "run_changelock_policy_dry_run"},
		InspectCommands:         []string{"changelock config inspect --format json", "changelock policy dry-run --audit-only"},
		DegradedBoundaries:      nil,
		DegradedAllowedExplicit: false,
		ProjectionDisclaimer:    "bootstrap_result_is_projection_only_and_never_implies_canonical_success",
		Limitations: []string{
			"Val A fail-fast bootstrap stops unsafe activation but does not add broader install or go-live orchestration.",
		},
	}
}

func ProductionUsabilityValAPolicySchemaDiscipline() PolicySchemaDiscipline {
	return PolicySchemaDiscipline{
		CurrentState:                  "versioned_policy_schema_ready",
		PolicySchemaVersion:           "point4.production_usability.vala.policy.v1",
		SupportedSchemaVersions:       productionUsabilityValAPolicySchemaVersions(),
		CurrentCompatibility:          ProductionUsabilityCompatibilityCurrent,
		CurrentValidationResult:       ProductionUsabilityValidationValid,
		CurrentUnknownFieldPolicy:     ProductionUsabilityUnknownFieldReject,
		UnknownFieldsDetected:         nil,
		ExplicitUnknownFieldAllowance: false,
		DeprecatedWarnings:            nil,
		MigrationWarnings:             nil,
		MigrationCompleted:            false,
		EffectiveInspectionView:       "projection_only_policy_effective_view_not_canonical_truth",
		Limitations: []string{
			"Val A adds policy schema strictness and inspection only; it does not rewrite the policy engine.",
		},
	}
}

func ProductionUsabilityValAEffectiveConfigInspection() EffectiveConfigInspection {
	return EffectiveConfigInspection{
		CurrentState:               "effective_config_inspection_ready",
		SchemaVersion:              "point4.production_usability.vala.config.v1",
		CompatibilityStatus:        ProductionUsabilityCompatibilityCurrent,
		ValidationResult:           ProductionUsabilityValidationValid,
		DefaultsApplied:            []string{"audit_store", "request_timeout"},
		UserProvidedFields:         []string{"schema_version", "tenant_id", "environment"},
		RejectedUnknownFields:      nil,
		Warnings:                   nil,
		Conflicts:                  nil,
		MigrationWarnings:          nil,
		PolicySchemaSummary:        "point4.production_usability.vala.policy.v1/current/valid",
		SourceProjectionDisclaimer: "effective_config_is_projection_only_not_canonical_truth",
		GeneratedAtIndicator:       "generated_at_present",
		LimitationNotes: []string{
			"Effective config inspection redacts secrets and surfaces defaults separately from user values.",
		},
		RedactedFields:  []string{"signing_key", "database_password"},
		PermissionAware: true,
		SecretsExposed:  false,
	}
}

func ProductionUsabilityValAHumanReadableRejectionLayer() HumanReadableRejectionLayer {
	return HumanReadableRejectionLayer{
		CurrentState:                "human_readable_rejection_layer_ready",
		RequiredFields:              productionUsabilityValARequiredExplainFields(),
		SupportedVisibilityScopes:   productionUsabilityValAExplainScopes(),
		SupportedRedactionTiers:     []string{ProductionUsabilityRedactionNone, ProductionUsabilityRedactionLow, ProductionUsabilityRedactionMedium, ProductionUsabilityRedactionHigh, ProductionUsabilityRedactionPublicSafe},
		SupportedActionModes:        []string{ProductionUsabilityActionModeExplain, ProductionUsabilityActionModePreview, ProductionUsabilityActionModeDryRun, ProductionUsabilityActionModeAuditOnly},
		SupportedDecisionPriorities: []string{ProductionUsabilityDecisionBlocker, ProductionUsabilityDecisionUrgent, ProductionUsabilityDecisionNormal, ProductionUsabilityDecisionLow, ProductionUsabilityDecisionInformational},
		TechnicalDetailScopes:       []string{ProductionUsabilityVisibilityInternalAdmin, ProductionUsabilityVisibilityOperator},
		RestrictedScopes:            []string{ProductionUsabilityVisibilityDeveloper, ProductionUsabilityVisibilityPartner, ProductionUsabilityVisibilityPublicSafe},
		SecretsRedacted:             true,
		TechnicalDetailBounded:      true,
		FailureDowngradedToWarning:  false,
		Limitations: []string{
			"Val A rejections are bounded explainability outputs and do not weaken validation outcomes.",
		},
	}
}

func ProductionUsabilityValAPolicyDryRunAuditFlow() PolicyDryRunAuditFlow {
	return PolicyDryRunAuditFlow{
		CurrentState:                 "policy_dry_run_audit_flow_ready",
		SupportedActionModes:         []string{ProductionUsabilityActionModeDryRun, ProductionUsabilityActionModeAuditOnly},
		PreviewAccepted:              []string{"schema_version", "tenant_id", "environment"},
		PreviewRejected:              []string{"unknown_field_example"},
		BlockingRules:                []string{"unsupported_schema_version_blocks_activation", "invalid_field_shape_blocks_activation"},
		NonBlockingWarnings:          []string{"deprecated_schema_warning", "migration_required_preview"},
		PermissionAwareOutput:        true,
		RecoveryHints:                []string{"run_changelock config inspect", "run_changelock policy dry-run --audit-only"},
		ProjectionDisclaimer:         "dry_run_and_audit_only_are_simulated_projection_only_and_non_authoritative",
		MutatesCanonicalState:        false,
		AuditOnlyImpliesApproval:     false,
		DryRunSuccessImpliesActivate: false,
		Limitations: []string{
			"Val A dry-run and audit-only remain non-mutating and cannot approve or activate production state.",
		},
	}
}

func ProductionUsabilityValAPermissionAwareExplainOutputs() PermissionAwareExplainOutputs {
	return PermissionAwareExplainOutputs{
		CurrentState:                "permission_aware_explain_outputs_ready",
		SupportedEvidenceVisibility: []string{ProductionUsabilityEvidenceFull, ProductionUsabilityEvidenceMetadataOnly, ProductionUsabilityEvidenceRedacted, ProductionUsabilityEvidenceHidden},
		Items: []ExplainOutputVariant{
			{VisibilityScope: ProductionUsabilityVisibilityInternalAdmin, CurrentState: "explain_variant_ready", DetailLevel: "full_technical_detail", EvidenceVisibility: ProductionUsabilityEvidenceFull, RedactionTier: ProductionUsabilityRedactionNone, HonestLimitation: "full_detail_still_subject_to_safe_handling", SecretsRedacted: true, HiddenEvidenceMarker: "sensitive_values_redacted_even_for_admin"},
			{VisibilityScope: ProductionUsabilityVisibilityOperator, CurrentState: "explain_variant_ready", DetailLevel: "operational_detail", EvidenceVisibility: ProductionUsabilityEvidenceMetadataOnly, RedactionTier: ProductionUsabilityRedactionLow, HonestLimitation: "operator_output_excludes_restricted_raw_evidence", SecretsRedacted: true, HiddenEvidenceMarker: "raw_evidence_hidden"},
			{VisibilityScope: ProductionUsabilityVisibilityDeveloper, CurrentState: "explain_variant_ready", DetailLevel: "actionable_remediation_detail", EvidenceVisibility: ProductionUsabilityEvidenceRedacted, RedactionTier: ProductionUsabilityRedactionMedium, HonestLimitation: "developer_output_hides_restricted_internal_evidence", SecretsRedacted: true, HiddenEvidenceMarker: "restricted_evidence_redacted"},
			{VisibilityScope: ProductionUsabilityVisibilityPartner, CurrentState: "explain_variant_ready", DetailLevel: "bounded_partner_detail", EvidenceVisibility: ProductionUsabilityEvidenceHidden, RedactionTier: ProductionUsabilityRedactionHigh, HonestLimitation: "partner_output_is_bounded_and_hidden_evidence_is_marked_as_hidden", SecretsRedacted: true, HiddenEvidenceMarker: "hidden_evidence_present"},
			{VisibilityScope: ProductionUsabilityVisibilityPublicSafe, CurrentState: "explain_variant_ready", DetailLevel: "public_safe_summary", EvidenceVisibility: ProductionUsabilityEvidenceHidden, RedactionTier: ProductionUsabilityRedactionPublicSafe, HonestLimitation: "public_safe_output_is_bounded_and_hidden_evidence_is_marked_as_hidden", SecretsRedacted: true, HiddenEvidenceMarker: "hidden_evidence_present"},
		},
		PreservesFailureSemantics: true,
		Limitations: []string{
			"Val A explain outputs remain projections and do not turn FAIL into PASS under redaction.",
		},
	}
}

func ProductionUsabilityValARecoveryGuidance() RecoveryGuidanceCore {
	return RecoveryGuidanceCore{
		CurrentState: "recovery_guidance_ready",
		Items: []RecoveryGuidanceItem{
			{FailureClass: "missing_required_field", CurrentState: "recovery_guidance_ready", RemediationClass: ProductionUsabilityRemediationConfigFix, RecoveryHint: "add_the_missing_required_field_and_reinspect", SafeRetryGuidance: "retry_after_fixing_the_required_field", DoNotRetryReason: "retry_without_fix_repeats_same_blocker", InspectOrExplainCommand: "changelock config inspect --explain missing_required_field", EscalationPath: "platform_config_owner"},
			{FailureClass: "unknown_field", CurrentState: "recovery_guidance_ready", RemediationClass: ProductionUsabilityRemediationConfigFix, RecoveryHint: "remove_the_unknown_field_or_enable_explicit_allowance", SafeRetryGuidance: "retry_after_schema_known_fields_match", DoNotRetryReason: "retry_without_schema_alignment_repeats_unknown_field_block", InspectOrExplainCommand: "changelock config inspect --explain unknown_field", EscalationPath: "platform_config_owner"},
			{FailureClass: "unsupported_schema_version", CurrentState: "recovery_guidance_ready", RemediationClass: ProductionUsabilityRemediationUnsupported, RecoveryHint: "upgrade_or_downgrade_to_a_supported_schema_version", SafeRetryGuidance: "retry_only_after_supported_schema_version_is_selected", DoNotRetryReason: "unsupported_schema_remains_blocked_until_support_exists", InspectOrExplainCommand: "changelock config inspect --explain unsupported_schema_version", EscalationPath: "platform_schema_owner"},
			{FailureClass: "deprecated_schema_version", CurrentState: "recovery_guidance_ready", RemediationClass: ProductionUsabilityRemediationPolicyUpdate, RecoveryHint: "plan_a_schema_upgrade_before_future_removal", SafeRetryGuidance: "retry_after_reviewing_deprecation_warning_output", DoNotRetryReason: "deprecated_schema_may_continue_temporarily_but_requires_review", InspectOrExplainCommand: "changelock config inspect --explain deprecated_schema_version", EscalationPath: "platform_schema_owner"},
			{FailureClass: "migration_required", CurrentState: "recovery_guidance_ready", RemediationClass: ProductionUsabilityRemediationManualReview, RecoveryHint: "review_required_migration_steps_before_activation", SafeRetryGuidance: "retry_after_required_migration_steps_are_reviewed", DoNotRetryReason: "migration_warning_does_not_mean_migration_complete", InspectOrExplainCommand: "changelock config inspect --explain migration_required", EscalationPath: "platform_schema_owner"},
			{FailureClass: "conflicting_values", CurrentState: "recovery_guidance_ready", RemediationClass: ProductionUsabilityRemediationConfigFix, RecoveryHint: "resolve_conflicting_values_before_retry", SafeRetryGuidance: "retry_after_single_consistent_value_set_is_selected", DoNotRetryReason: "conflict_persists_until_values_are_reconciled", InspectOrExplainCommand: "changelock config inspect --explain conflicting_values", EscalationPath: "platform_config_owner"},
			{FailureClass: "invalid_field_shape", CurrentState: "recovery_guidance_ready", RemediationClass: ProductionUsabilityRemediationConfigFix, RecoveryHint: "fix_field_type_or_shape_to_match_schema", SafeRetryGuidance: "retry_after_shape_matches_schema", DoNotRetryReason: "retry_without_type_fix_repeats_same_validation_failure", InspectOrExplainCommand: "changelock config inspect --explain invalid_field_shape", EscalationPath: "platform_config_owner"},
			{FailureClass: "permission_insufficient_for_detailed_explain", CurrentState: "recovery_guidance_ready", RemediationClass: ProductionUsabilityRemediationPermissionFix, RecoveryHint: "request_internal_admin_or_operator_scope_for_more_detail", SafeRetryGuidance: "retry_after_permission_scope_is_corrected", DoNotRetryReason: "bounded_output_cannot_expand_without_authorized_scope", InspectOrExplainCommand: "changelock config explain --scope operator", EscalationPath: "security_admin"},
			{FailureClass: "dry_run_success_activation_not_performed", CurrentState: "recovery_guidance_ready", RemediationClass: ProductionUsabilityRemediationManualReview, RecoveryHint: "review_dry_run_output_and_execute_governed_activation_separately", SafeRetryGuidance: "retry_dry_run_only_if_inputs_change", DoNotRetryReason: "dry_run_success_never_activates_production_state", InspectOrExplainCommand: "changelock policy dry-run --audit-only", EscalationPath: "platform_operator"},
		},
		UnsupportedExplicit: true,
		Limitations: []string{
			"Val A recovery guidance remains bounded advice and never suggests policy bypass or manual evidence tampering.",
		},
	}
}

func ProductionUsabilityValAFirstRunSafeBootstrap() FirstRunSafeBootstrap {
	return FirstRunSafeBootstrap{
		CurrentState:               "first_run_safe_bootstrap_ready",
		MinimalSafeShapeFields:     []string{"schema_version", "tenant_id", "environment", "audit_store"},
		MissingReadinessFields:     []string{"signing_identity", "policy_bundle_ref"},
		SampleConfigDetected:       true,
		SampleMarkedNonProduction:  true,
		AutoEnablesProduction:      false,
		FakeDemoEvidencePresent:    false,
		NextSteps:                  []string{"replace_sample_values_with_real_production_values", "run_changelock_config_inspect_before_activation"},
		NonMutatingPreview:         true,
		BootstrapPathValid:         true,
		ClaimsProductionCompletion: false,
	}
}

func ProductionUsabilityValAUpgradeImpactPreview() UpgradeImpactPreview {
	return UpgradeImpactPreview{
		CurrentState:             "upgrade_impact_preview_ready",
		CurrentSchemaVersion:     "point4.production_usability.vala.config.v1",
		TargetSchemaVersion:      "point4.production_usability.vala.config.v2-preview",
		KnownTargetSchemas:       []string{"point4.production_usability.vala.config.v1", "point4.production_usability.vala.config.v2-preview"},
		CompatibilityStatus:      ProductionUsabilityCompatibilityMigrationRequired,
		DeprecatedFields:         []string{"legacy_timeout_seconds"},
		RemovedUnsupportedFields: []string{"deprecated_legacy_token"},
		MigrationRequiredItems:   []string{"rename_legacy_timeout_seconds", "set_request_timeout_duration"},
		Warnings:                 []string{"v2_preview_requires_schema_review_before_activation"},
		BlockingIssues:           []string{"missing_required_timeout_duration"},
		RollbackPerspective:      "config_policy_schema_only_bounded_projection",
		RollbackAppearsPossible:  true,
		LimitationDisclaimer:     "upgrade_impact_preview_is_projection_only_and_bounded_to_config_policy_schema_perspective",
		MutatesConfig:            false,
	}
}

func EvaluateProductionUsabilityValAConfigFactoryState(model SchemaStrictConfigFactory) string {
	if strings.TrimSpace(model.SchemaVersion) == "" || strings.TrimSpace(model.CurrentCompatibility) == "" || strings.TrimSpace(model.CurrentUnknownFieldPolicy) == "" || strings.TrimSpace(model.CurrentValidationResult) == "" || strings.TrimSpace(model.InspectionMode) == "" {
		return ProductionUsabilityValAConfigFactoryStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedSchemaVersions, productionUsabilityValAConfigSchemaVersions()...) ||
		!containsTrimmedString(model.SupportedSchemaVersions, model.SchemaVersion) ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().SupportedCompatibility, model.CurrentCompatibility) ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().ValidationStates, model.CurrentValidationResult) {
		return ProductionUsabilityValAConfigFactoryStatePartial
	}
	switch strings.TrimSpace(model.CurrentUnknownFieldPolicy) {
	case ProductionUsabilityUnknownFieldReject, ProductionUsabilityUnknownFieldWarn, ProductionUsabilityUnknownFieldIgnoreExplicitlyAllowed:
	default:
		return ProductionUsabilityValAConfigFactoryStatePartial
	}
	if !model.RequiredFieldValidation || !model.TypeShapeValidation || !model.ConflictDetection || !model.FailFastBootstrap ||
		!model.DefaultsDeclaredExplicitly || !model.DefaultsSurfacedInInspection || !model.DefaultsDistinguishable || !model.SecretsRedacted {
		return ProductionUsabilityValAConfigFactoryStatePartial
	}
	if !strings.Contains(strings.TrimSpace(model.InspectionMode), "projection_only") || hasTrimmedStringOverlap(model.DefaultsApplied, model.UserProvidedFields) {
		return ProductionUsabilityValAConfigFactoryStatePartial
	}
	if len(model.UnknownFieldsDetected) > 0 {
		switch strings.TrimSpace(model.CurrentUnknownFieldPolicy) {
		case ProductionUsabilityUnknownFieldReject:
			return ProductionUsabilityValAConfigFactoryStatePartial
		case ProductionUsabilityUnknownFieldWarn, ProductionUsabilityUnknownFieldIgnoreExplicitlyAllowed:
			if !model.ExplicitUnknownFieldAllowance {
				return ProductionUsabilityValAConfigFactoryStatePartial
			}
		}
	}
	if strings.TrimSpace(model.CurrentCompatibility) == ProductionUsabilityCompatibilityUnsupported ||
		strings.TrimSpace(model.CurrentCompatibility) == ProductionUsabilityCompatibilityUnknown ||
		strings.TrimSpace(model.CurrentValidationResult) == ProductionUsabilityValidationInvalid ||
		strings.TrimSpace(model.CurrentValidationResult) == ProductionUsabilityValidationUnsupported {
		return ProductionUsabilityValAConfigFactoryStatePartial
	}
	if strings.TrimSpace(model.CurrentCompatibility) == ProductionUsabilityCompatibilityDeprecated && len(model.CompatibilityWarnings) == 0 {
		return ProductionUsabilityValAConfigFactoryStatePartial
	}
	if len(model.MigrationWarnings) > 0 && model.MigrationCompleted {
		return ProductionUsabilityValAConfigFactoryStatePartial
	}
	return ProductionUsabilityValAConfigFactoryStateActive
}

func EvaluateProductionUsabilityValABootstrapValidationState(model BootstrapValidationCore) string {
	if strings.TrimSpace(model.ConfigValidationResult) == "" || strings.TrimSpace(model.PolicyValidationResult) == "" || strings.TrimSpace(model.BootstrapDisposition) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ProductionUsabilityValABootstrapValidationStateIncomplete
	}
	if !containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().ValidationStates, model.ConfigValidationResult) ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().ValidationStates, model.PolicyValidationResult) ||
		!containsTrimmedString([]string{ProductionUsabilityBootstrapAllowed, ProductionUsabilityBootstrapBlocked, ProductionUsabilityBootstrapDegradedAllowed, ProductionUsabilityBootstrapUnsupported}, model.BootstrapDisposition) ||
		!containsExactTrimmedStringSet(model.RejectionFields, productionUsabilityValARequiredExplainFields()...) ||
		!model.AggregateReportingSafe ||
		len(model.RecoveryHints) == 0 ||
		len(model.InspectCommands) == 0 ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") {
		return ProductionUsabilityValABootstrapValidationStatePartial
	}
	configValidation := strings.TrimSpace(model.ConfigValidationResult)
	policyValidation := strings.TrimSpace(model.PolicyValidationResult)
	disposition := strings.TrimSpace(model.BootstrapDisposition)
	if configValidation == ProductionUsabilityValidationInvalid || policyValidation == ProductionUsabilityValidationInvalid {
		if disposition != ProductionUsabilityBootstrapBlocked || len(model.BlockingReasonCodes) == 0 {
			return ProductionUsabilityValABootstrapValidationStatePartial
		}
	}
	if configValidation == ProductionUsabilityValidationUnsupported || policyValidation == ProductionUsabilityValidationUnsupported {
		if disposition != ProductionUsabilityBootstrapUnsupported || len(model.BlockingReasonCodes) == 0 {
			return ProductionUsabilityValABootstrapValidationStatePartial
		}
	}
	if configValidation == ProductionUsabilityValidationDegraded || policyValidation == ProductionUsabilityValidationDegraded || disposition == ProductionUsabilityBootstrapDegradedAllowed {
		if disposition != ProductionUsabilityBootstrapDegradedAllowed || !model.DegradedAllowedExplicit || len(model.DegradedBoundaries) == 0 {
			return ProductionUsabilityValABootstrapValidationStatePartial
		}
	}
	if disposition == ProductionUsabilityBootstrapAllowed && len(model.BlockingReasonCodes) > 0 {
		return ProductionUsabilityValABootstrapValidationStatePartial
	}
	return ProductionUsabilityValABootstrapValidationStateActive
}

func EvaluateProductionUsabilityValAPolicySchemaState(model PolicySchemaDiscipline) string {
	if strings.TrimSpace(model.PolicySchemaVersion) == "" || strings.TrimSpace(model.CurrentCompatibility) == "" || strings.TrimSpace(model.CurrentValidationResult) == "" || strings.TrimSpace(model.CurrentUnknownFieldPolicy) == "" || strings.TrimSpace(model.EffectiveInspectionView) == "" {
		return ProductionUsabilityValAPolicySchemaStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedSchemaVersions, productionUsabilityValAPolicySchemaVersions()...) ||
		!containsTrimmedString(model.SupportedSchemaVersions, model.PolicySchemaVersion) ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().SupportedCompatibility, model.CurrentCompatibility) ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().ValidationStates, model.CurrentValidationResult) {
		return ProductionUsabilityValAPolicySchemaStatePartial
	}
	switch strings.TrimSpace(model.CurrentUnknownFieldPolicy) {
	case ProductionUsabilityUnknownFieldReject, ProductionUsabilityUnknownFieldWarn, ProductionUsabilityUnknownFieldIgnoreExplicitlyAllowed:
	default:
		return ProductionUsabilityValAPolicySchemaStatePartial
	}
	if !strings.Contains(strings.TrimSpace(model.EffectiveInspectionView), "projection_only") {
		return ProductionUsabilityValAPolicySchemaStatePartial
	}
	if len(model.UnknownFieldsDetected) > 0 {
		switch strings.TrimSpace(model.CurrentUnknownFieldPolicy) {
		case ProductionUsabilityUnknownFieldReject:
			return ProductionUsabilityValAPolicySchemaStatePartial
		case ProductionUsabilityUnknownFieldWarn, ProductionUsabilityUnknownFieldIgnoreExplicitlyAllowed:
			if !model.ExplicitUnknownFieldAllowance {
				return ProductionUsabilityValAPolicySchemaStatePartial
			}
		}
	}
	if strings.TrimSpace(model.CurrentCompatibility) == ProductionUsabilityCompatibilityUnsupported ||
		strings.TrimSpace(model.CurrentCompatibility) == ProductionUsabilityCompatibilityUnknown ||
		strings.TrimSpace(model.CurrentValidationResult) == ProductionUsabilityValidationInvalid ||
		strings.TrimSpace(model.CurrentValidationResult) == ProductionUsabilityValidationUnsupported {
		return ProductionUsabilityValAPolicySchemaStatePartial
	}
	if strings.TrimSpace(model.CurrentCompatibility) == ProductionUsabilityCompatibilityDeprecated && len(model.DeprecatedWarnings) == 0 {
		return ProductionUsabilityValAPolicySchemaStatePartial
	}
	if len(model.MigrationWarnings) > 0 && model.MigrationCompleted {
		return ProductionUsabilityValAPolicySchemaStatePartial
	}
	return ProductionUsabilityValAPolicySchemaStateActive
}

func EvaluateProductionUsabilityValAEffectiveConfigState(model EffectiveConfigInspection) string {
	if strings.TrimSpace(model.SchemaVersion) == "" || strings.TrimSpace(model.CompatibilityStatus) == "" || strings.TrimSpace(model.ValidationResult) == "" || strings.TrimSpace(model.PolicySchemaSummary) == "" || strings.TrimSpace(model.SourceProjectionDisclaimer) == "" || strings.TrimSpace(model.GeneratedAtIndicator) == "" {
		return ProductionUsabilityValAEffectiveConfigStateIncomplete
	}
	if !containsTrimmedString(productionUsabilityValAConfigSchemaVersions(), model.SchemaVersion) ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().SupportedCompatibility, model.CompatibilityStatus) ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().ValidationStates, model.ValidationResult) ||
		len(model.RedactedFields) == 0 ||
		!model.PermissionAware ||
		model.SecretsExposed ||
		!strings.Contains(strings.TrimSpace(model.SourceProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.SourceProjectionDisclaimer), "not_canonical_truth") ||
		hasTrimmedStringOverlap(model.DefaultsApplied, model.UserProvidedFields) {
		return ProductionUsabilityValAEffectiveConfigStatePartial
	}
	return ProductionUsabilityValAEffectiveConfigStateActive
}

func EvaluateProductionUsabilityValARejectionLayerState(model HumanReadableRejectionLayer) string {
	if len(model.RequiredFields) == 0 {
		return ProductionUsabilityValARejectionLayerStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.RequiredFields, productionUsabilityValARequiredExplainFields()...) ||
		!containsExactTrimmedStringSet(model.SupportedVisibilityScopes, productionUsabilityValAExplainScopes()...) ||
		!containsExactTrimmedStringSet(model.SupportedRedactionTiers, ProductionUsabilityRedactionNone, ProductionUsabilityRedactionLow, ProductionUsabilityRedactionMedium, ProductionUsabilityRedactionHigh, ProductionUsabilityRedactionPublicSafe) ||
		!containsAllTrimmedStrings(model.SupportedActionModes, ProductionUsabilityActionModeExplain, ProductionUsabilityActionModePreview, ProductionUsabilityActionModeDryRun, ProductionUsabilityActionModeAuditOnly) ||
		!containsExactTrimmedStringSet(model.SupportedDecisionPriorities, ProductionUsabilityDecisionBlocker, ProductionUsabilityDecisionUrgent, ProductionUsabilityDecisionNormal, ProductionUsabilityDecisionLow, ProductionUsabilityDecisionInformational) ||
		!containsAllTrimmedStrings(model.TechnicalDetailScopes, ProductionUsabilityVisibilityInternalAdmin, ProductionUsabilityVisibilityOperator) ||
		!containsAllTrimmedStrings(model.RestrictedScopes, ProductionUsabilityVisibilityDeveloper, ProductionUsabilityVisibilityPartner, ProductionUsabilityVisibilityPublicSafe) ||
		!model.SecretsRedacted ||
		!model.TechnicalDetailBounded ||
		model.FailureDowngradedToWarning {
		return ProductionUsabilityValARejectionLayerStatePartial
	}
	return ProductionUsabilityValARejectionLayerStateActive
}

func EvaluateProductionUsabilityValADryRunState(model PolicyDryRunAuditFlow) string {
	if len(model.SupportedActionModes) == 0 {
		return ProductionUsabilityValADryRunStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedActionModes, ProductionUsabilityActionModeDryRun, ProductionUsabilityActionModeAuditOnly) ||
		len(model.PreviewAccepted) == 0 ||
		len(model.PreviewRejected) == 0 ||
		len(model.BlockingRules) == 0 ||
		len(model.NonBlockingWarnings) == 0 ||
		!model.PermissionAwareOutput ||
		len(model.RecoveryHints) == 0 ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "simulated") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "non_authoritative") ||
		model.MutatesCanonicalState ||
		model.AuditOnlyImpliesApproval ||
		model.DryRunSuccessImpliesActivate {
		return ProductionUsabilityValADryRunStatePartial
	}
	return ProductionUsabilityValADryRunStateActive
}

func EvaluateProductionUsabilityValAExplainState(model PermissionAwareExplainOutputs) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValAExplainStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedEvidenceVisibility, ProductionUsabilityEvidenceFull, ProductionUsabilityEvidenceMetadataOnly, ProductionUsabilityEvidenceRedacted, ProductionUsabilityEvidenceHidden) ||
		!model.PreservesFailureSemantics {
		return ProductionUsabilityValAExplainStatePartial
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
		if scope == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.DetailLevel) == "" || strings.TrimSpace(item.EvidenceVisibility) == "" || strings.TrimSpace(item.RedactionTier) == "" || strings.TrimSpace(item.HonestLimitation) == "" {
			return ProductionUsabilityValAExplainStateIncomplete
		}
		if _, ok := expectedScopes[scope]; !ok {
			return ProductionUsabilityValAExplainStatePartial
		}
		if _, duplicate := seen[scope]; duplicate {
			return ProductionUsabilityValAExplainStatePartial
		}
		seen[scope] = struct{}{}
		if !containsTrimmedString(model.SupportedEvidenceVisibility, item.EvidenceVisibility) || !item.SecretsRedacted {
			return ProductionUsabilityValAExplainStatePartial
		}
		if item.EvidenceVisibility == ProductionUsabilityEvidenceRedacted || item.EvidenceVisibility == ProductionUsabilityEvidenceHidden {
			if strings.TrimSpace(item.HiddenEvidenceMarker) == "" {
				return ProductionUsabilityValAExplainStatePartial
			}
		}
		if (scope == ProductionUsabilityVisibilityPartner || scope == ProductionUsabilityVisibilityPublicSafe) && item.EvidenceVisibility == ProductionUsabilityEvidenceFull {
			return ProductionUsabilityValAExplainStatePartial
		}
	}
	if len(seen) != len(expectedScopes) {
		return ProductionUsabilityValAExplainStatePartial
	}
	return ProductionUsabilityValAExplainStateActive
}

func EvaluateProductionUsabilityValARecoveryGuidanceState(model RecoveryGuidanceCore) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValARecoveryGuidanceStateIncomplete
	}
	if !model.UnsupportedExplicit {
		return ProductionUsabilityValARecoveryGuidanceStatePartial
	}
	expectedFailureClasses := map[string]struct{}{}
	for _, item := range productionUsabilityValARecoveryFailureClasses() {
		expectedFailureClasses[item] = struct{}{}
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		failureClass := strings.TrimSpace(item.FailureClass)
		if failureClass == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.RemediationClass) == "" || strings.TrimSpace(item.RecoveryHint) == "" || strings.TrimSpace(item.SafeRetryGuidance) == "" || strings.TrimSpace(item.DoNotRetryReason) == "" || strings.TrimSpace(item.InspectOrExplainCommand) == "" || strings.TrimSpace(item.EscalationPath) == "" {
			return ProductionUsabilityValARecoveryGuidanceStateIncomplete
		}
		if _, ok := expectedFailureClasses[failureClass]; !ok {
			return ProductionUsabilityValARecoveryGuidanceStatePartial
		}
		if _, duplicate := seen[failureClass]; duplicate {
			return ProductionUsabilityValARecoveryGuidanceStatePartial
		}
		seen[failureClass] = struct{}{}
		if !containsTrimmedString(ProductionUsabilityVal0RecoveryUXContract().RemediationClasses, item.RemediationClass) ||
			item.UnsafeRetrySuggested ||
			item.PolicyBypassSuggested ||
			item.CanonicalEvidenceEditSuggested {
			return ProductionUsabilityValARecoveryGuidanceStatePartial
		}
	}
	if len(seen) != len(expectedFailureClasses) {
		return ProductionUsabilityValARecoveryGuidanceStatePartial
	}
	return ProductionUsabilityValARecoveryGuidanceStateActive
}

func EvaluateProductionUsabilityValAFirstRunState(model FirstRunSafeBootstrap) string {
	if len(model.MinimalSafeShapeFields) == 0 || len(model.NextSteps) == 0 {
		return ProductionUsabilityValAFirstRunStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.MinimalSafeShapeFields, "schema_version", "tenant_id", "environment", "audit_store") ||
		!model.SampleConfigDetected ||
		!model.SampleMarkedNonProduction ||
		model.AutoEnablesProduction ||
		model.FakeDemoEvidencePresent ||
		!model.NonMutatingPreview ||
		!model.BootstrapPathValid ||
		model.ClaimsProductionCompletion {
		return ProductionUsabilityValAFirstRunStatePartial
	}
	return ProductionUsabilityValAFirstRunStateActive
}

func EvaluateProductionUsabilityValAUpgradePreviewState(model UpgradeImpactPreview) string {
	if strings.TrimSpace(model.CurrentSchemaVersion) == "" || strings.TrimSpace(model.TargetSchemaVersion) == "" || strings.TrimSpace(model.CompatibilityStatus) == "" || strings.TrimSpace(model.RollbackPerspective) == "" || strings.TrimSpace(model.LimitationDisclaimer) == "" {
		return ProductionUsabilityValAUpgradePreviewStateIncomplete
	}
	if !containsTrimmedString(model.KnownTargetSchemas, model.TargetSchemaVersion) ||
		!containsTrimmedString(append(ProductionUsabilityVal0ConfigIntegrity().SupportedCompatibility, ProductionUsabilityCompatibilityMigrationRequired), model.CompatibilityStatus) ||
		!strings.Contains(strings.TrimSpace(model.RollbackPerspective), "config_policy_schema_only") ||
		!strings.Contains(strings.TrimSpace(model.LimitationDisclaimer), "projection_only") ||
		model.MutatesConfig {
		return ProductionUsabilityValAUpgradePreviewStatePartial
	}
	if strings.TrimSpace(model.CompatibilityStatus) == ProductionUsabilityCompatibilityUnknown || strings.TrimSpace(model.CompatibilityStatus) == ProductionUsabilityCompatibilityUnsupported {
		return ProductionUsabilityValAUpgradePreviewStatePartial
	}
	return ProductionUsabilityValAUpgradePreviewStateActive
}

func EvaluateProductionUsabilityValAState(val0State, configFactoryState, bootstrapValidationState, policySchemaState, effectiveConfigState, rejectionLayerState, dryRunState, explainState, recoveryGuidanceState, firstRunState, upgradePreviewState string) string {
	if strings.TrimSpace(val0State) != ProductionUsabilityVal0StateActive {
		return ProductionUsabilityValAStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(configFactoryState),
		strings.TrimSpace(bootstrapValidationState),
		strings.TrimSpace(policySchemaState),
		strings.TrimSpace(effectiveConfigState),
		strings.TrimSpace(rejectionLayerState),
		strings.TrimSpace(dryRunState),
		strings.TrimSpace(explainState),
		strings.TrimSpace(recoveryGuidanceState),
		strings.TrimSpace(firstRunState),
		strings.TrimSpace(upgradePreviewState),
	} {
		switch state {
		case ProductionUsabilityValAConfigFactoryStateActive,
			ProductionUsabilityValABootstrapValidationStateActive,
			ProductionUsabilityValAPolicySchemaStateActive,
			ProductionUsabilityValAEffectiveConfigStateActive,
			ProductionUsabilityValARejectionLayerStateActive,
			ProductionUsabilityValADryRunStateActive,
			ProductionUsabilityValAExplainStateActive,
			ProductionUsabilityValARecoveryGuidanceStateActive,
			ProductionUsabilityValAFirstRunStateActive,
			ProductionUsabilityValAUpgradePreviewStateActive:
		case ProductionUsabilityValAConfigFactoryStatePartial,
			ProductionUsabilityValABootstrapValidationStatePartial,
			ProductionUsabilityValAPolicySchemaStatePartial,
			ProductionUsabilityValAEffectiveConfigStatePartial,
			ProductionUsabilityValARejectionLayerStatePartial,
			ProductionUsabilityValADryRunStatePartial,
			ProductionUsabilityValAExplainStatePartial,
			ProductionUsabilityValARecoveryGuidanceStatePartial,
			ProductionUsabilityValAFirstRunStatePartial,
			ProductionUsabilityValAUpgradePreviewStatePartial:
			hasPartial = true
		default:
			return ProductionUsabilityValAStateIncomplete
		}
	}
	if hasPartial {
		return ProductionUsabilityValAStateSubstantial
	}
	return ProductionUsabilityValAStateActive
}

func EvaluateProductionUsabilityValAProofsState(val0State, configFactoryState, bootstrapValidationState, policySchemaState, effectiveConfigState, rejectionLayerState, dryRunState, explainState, recoveryGuidanceState, firstRunState, upgradePreviewState string, surfaceRefs, evidenceRefs, limitations, whyPoint4NotPass []string) string {
	baseState := EvaluateProductionUsabilityValAState(val0State, configFactoryState, bootstrapValidationState, policySchemaState, effectiveConfigState, rejectionLayerState, dryRunState, explainState, recoveryGuidanceState, firstRunState, upgradePreviewState)
	if len(surfaceRefs) < 11 || len(evidenceRefs) < 7 || len(limitations) == 0 || len(whyPoint4NotPass) == 0 {
		if baseState == ProductionUsabilityValAStateActive {
			return ProductionUsabilityValAStateSubstantial
		}
		return baseState
	}
	return baseState
}

func hasTrimmedStringOverlap(left, right []string) bool {
	seen := make(map[string]struct{}, len(left))
	for _, item := range left {
		trimmed := strings.TrimSpace(item)
		if trimmed == "" {
			continue
		}
		seen[trimmed] = struct{}{}
	}
	for _, item := range right {
		trimmed := strings.TrimSpace(item)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			return true
		}
	}
	return false
}

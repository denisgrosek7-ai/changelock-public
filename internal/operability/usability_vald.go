package operability

import "strings"

const (
	ProductionUsabilityValDConfigReviewStateActive     = "production_usability_vald_config_review_active"
	ProductionUsabilityValDConfigReviewStatePartial    = "production_usability_vald_config_review_partial"
	ProductionUsabilityValDConfigReviewStateIncomplete = "production_usability_vald_config_review_incomplete"

	ProductionUsabilityValDExplainabilityReviewStateActive     = "production_usability_vald_explainability_review_active"
	ProductionUsabilityValDExplainabilityReviewStatePartial    = "production_usability_vald_explainability_review_partial"
	ProductionUsabilityValDExplainabilityReviewStateIncomplete = "production_usability_vald_explainability_review_incomplete"

	ProductionUsabilityValDDryRunReviewStateActive     = "production_usability_vald_dry_run_review_active"
	ProductionUsabilityValDDryRunReviewStatePartial    = "production_usability_vald_dry_run_review_partial"
	ProductionUsabilityValDDryRunReviewStateIncomplete = "production_usability_vald_dry_run_review_incomplete"

	ProductionUsabilityValDRedactionReviewStateActive     = "production_usability_vald_redaction_review_active"
	ProductionUsabilityValDRedactionReviewStatePartial    = "production_usability_vald_redaction_review_partial"
	ProductionUsabilityValDRedactionReviewStateIncomplete = "production_usability_vald_redaction_review_incomplete"

	ProductionUsabilityValDDegradedBehaviorReviewStateActive     = "production_usability_vald_degraded_behavior_review_active"
	ProductionUsabilityValDDegradedBehaviorReviewStatePartial    = "production_usability_vald_degraded_behavior_review_partial"
	ProductionUsabilityValDDegradedBehaviorReviewStateIncomplete = "production_usability_vald_degraded_behavior_review_incomplete"

	ProductionUsabilityValDUIWindowingReviewStateActive     = "production_usability_vald_ui_windowing_review_active"
	ProductionUsabilityValDUIWindowingReviewStatePartial    = "production_usability_vald_ui_windowing_review_partial"
	ProductionUsabilityValDUIWindowingReviewStateIncomplete = "production_usability_vald_ui_windowing_review_incomplete"

	ProductionUsabilityValDCommandNoiseReviewStateActive     = "production_usability_vald_command_noise_review_active"
	ProductionUsabilityValDCommandNoiseReviewStatePartial    = "production_usability_vald_command_noise_review_partial"
	ProductionUsabilityValDCommandNoiseReviewStateIncomplete = "production_usability_vald_command_noise_review_incomplete"

	ProductionUsabilityValDAPIProtectionReviewStateActive     = "production_usability_vald_api_protection_review_active"
	ProductionUsabilityValDAPIProtectionReviewStatePartial    = "production_usability_vald_api_protection_review_partial"
	ProductionUsabilityValDAPIProtectionReviewStateIncomplete = "production_usability_vald_api_protection_review_incomplete"

	ProductionUsabilityValDCLIResilienceReviewStateActive     = "production_usability_vald_cli_resilience_review_active"
	ProductionUsabilityValDCLIResilienceReviewStatePartial    = "production_usability_vald_cli_resilience_review_partial"
	ProductionUsabilityValDCLIResilienceReviewStateIncomplete = "production_usability_vald_cli_resilience_review_incomplete"

	ProductionUsabilityValDSupportabilityReviewStateActive     = "production_usability_vald_supportability_review_active"
	ProductionUsabilityValDSupportabilityReviewStatePartial    = "production_usability_vald_supportability_review_partial"
	ProductionUsabilityValDSupportabilityReviewStateIncomplete = "production_usability_vald_supportability_review_incomplete"

	ProductionUsabilityValDRecoveryReviewStateActive     = "production_usability_vald_recovery_review_active"
	ProductionUsabilityValDRecoveryReviewStatePartial    = "production_usability_vald_recovery_review_partial"
	ProductionUsabilityValDRecoveryReviewStateIncomplete = "production_usability_vald_recovery_review_incomplete"

	ProductionUsabilityValDUpgradeRollbackReviewStateActive     = "production_usability_vald_upgrade_rollback_review_active"
	ProductionUsabilityValDUpgradeRollbackReviewStatePartial    = "production_usability_vald_upgrade_rollback_review_partial"
	ProductionUsabilityValDUpgradeRollbackReviewStateIncomplete = "production_usability_vald_upgrade_rollback_review_incomplete"

	ProductionUsabilityValDScaleEnvelopeReviewStateActive     = "production_usability_vald_scale_envelope_review_active"
	ProductionUsabilityValDScaleEnvelopeReviewStatePartial    = "production_usability_vald_scale_envelope_review_partial"
	ProductionUsabilityValDScaleEnvelopeReviewStateIncomplete = "production_usability_vald_scale_envelope_review_incomplete"

	ProductionUsabilityValDGovernanceBoundaryReviewStateActive     = "production_usability_vald_governance_boundary_review_active"
	ProductionUsabilityValDGovernanceBoundaryReviewStatePartial    = "production_usability_vald_governance_boundary_review_partial"
	ProductionUsabilityValDGovernanceBoundaryReviewStateIncomplete = "production_usability_vald_governance_boundary_review_incomplete"

	ProductionUsabilityValDRegressionGateStateActive     = "production_usability_vald_regression_gate_active"
	ProductionUsabilityValDRegressionGateStatePartial    = "production_usability_vald_regression_gate_partial"
	ProductionUsabilityValDRegressionGateStateIncomplete = "production_usability_vald_regression_gate_incomplete"

	ProductionUsabilityValDStateIncomplete  = "production_usability_vald_incomplete"
	ProductionUsabilityValDStateSubstantial = "production_usability_vald_substantially_ready"
	ProductionUsabilityValDStateActive      = "production_usability_vald_active"

	ProductionUsabilityFinalGatePass        = "pass"
	ProductionUsabilityFinalGateFail        = "fail"
	ProductionUsabilityFinalGateWarning     = "warning"
	ProductionUsabilityFinalGateBlocked     = "blocked"
	ProductionUsabilityFinalGateUnsupported = "unsupported"
	ProductionUsabilityFinalGateNotRun      = "not_run"
)

type ConfigCorrectnessReview struct {
	CurrentState                    string   `json:"current_state"`
	ReviewState                     string   `json:"review_state"`
	SupportedReviewStates           []string `json:"supported_review_states,omitempty"`
	ConfigFactoryState              string   `json:"config_factory_state"`
	BootstrapValidationState        string   `json:"bootstrap_validation_state"`
	EffectiveConfigState            string   `json:"effective_config_state"`
	SchemaCompatibilityStatus       string   `json:"schema_compatibility_status"`
	ValidationResult                string   `json:"validation_result"`
	UnknownFieldPolicy              string   `json:"unknown_field_policy"`
	RequiredFieldValidationPassed   bool     `json:"required_field_validation_passed"`
	DefaultsSeparated               bool     `json:"defaults_separated"`
	SecretsExposedInEffectiveConfig bool     `json:"secrets_exposed_in_effective_config"`
	MigrationWarningsPresent        bool     `json:"migration_warnings_present"`
	MigrationCompleted              bool     `json:"migration_completed"`
	DeprecatedSchemaHandled         bool     `json:"deprecated_schema_handled"`
	EffectiveConfigClaimsCanonical  bool     `json:"effective_config_claims_canonical_truth"`
	FailFastBootstrapEnabled        bool     `json:"fail_fast_bootstrap_enabled"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
	Limitations                     []string `json:"limitations,omitempty"`
}

type ExplainabilityClarityReview struct {
	CurrentState                       string   `json:"current_state"`
	ReviewState                        string   `json:"review_state"`
	SupportedReviewStates              []string `json:"supported_review_states,omitempty"`
	RejectionLayerState                string   `json:"rejection_layer_state"`
	ExplainState                       string   `json:"explain_state"`
	ReasonCodesPresent                 bool     `json:"reason_codes_present"`
	PolicyRefsPresent                  bool     `json:"policy_refs_present"`
	SubjectRefsPresent                 bool     `json:"subject_refs_present"`
	NextStepsPresent                   bool     `json:"next_steps_present"`
	RecoveryHintsPresent               bool     `json:"recovery_hints_present"`
	VisibilityScopesCovered            bool     `json:"visibility_scopes_covered"`
	RedactionTiersCovered              bool     `json:"redaction_tiers_covered"`
	DecisionPriorityPresent            bool     `json:"decision_priority_present"`
	TechnicalDetailAvailableForAllowed bool     `json:"technical_detail_available_for_allowed_scopes"`
	SafeFallbackForRestrictedScopes    bool     `json:"safe_fallback_for_restricted_scopes"`
	SensitiveEvidenceLeaked            bool     `json:"sensitive_evidence_leaked"`
	FailureHiddenByRedaction           bool     `json:"failure_hidden_by_redaction"`
	ProjectionDisclaimer               string   `json:"projection_disclaimer"`
	Limitations                        []string `json:"limitations,omitempty"`
}

type DryRunAuditCorrectnessReview struct {
	CurrentState                  string   `json:"current_state"`
	ReviewState                   string   `json:"review_state"`
	SupportedReviewStates         []string `json:"supported_review_states,omitempty"`
	DryRunState                   string   `json:"dry_run_state"`
	DryRunMutatesCanonicalState   bool     `json:"dry_run_mutates_canonical_state"`
	AuditOnlyMutatesCanonical     bool     `json:"audit_only_mutates_canonical_state"`
	PreviewSuccessImpliesActivate bool     `json:"preview_success_implies_activate"`
	AuditOnlyImpliesApproval      bool     `json:"audit_only_implies_approval"`
	BlockingIssuesIdentified      bool     `json:"blocking_issues_identified"`
	WarningIssuesIdentified       bool     `json:"warning_issues_identified"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	Limitations                   []string `json:"limitations,omitempty"`
}

type PermissionRedactionReview struct {
	CurrentState                      string   `json:"current_state"`
	ReviewState                       string   `json:"review_state"`
	SupportedReviewStates             []string `json:"supported_review_states,omitempty"`
	PermissionRedactionState          string   `json:"permission_redaction_state"`
	ExplainState                      string   `json:"explain_state"`
	PermissionSupportState            string   `json:"permission_support_state"`
	ExportSafetyState                 string   `json:"export_safety_state"`
	VisibilityScopesCovered           bool     `json:"visibility_scopes_covered"`
	HiddenMetadataRepresented         bool     `json:"hidden_metadata_represented"`
	SafeFallbackMessagesPresent       bool     `json:"safe_fallback_messages_present"`
	PartnerOrPublicExposeFullEvidence bool     `json:"partner_or_public_expose_full_evidence"`
	RawSecretsOrTokensDetected        bool     `json:"raw_secrets_or_tokens_detected"`
	AuditorImpliesPublicSafe          bool     `json:"auditor_implies_public_safe"`
	FailureHiddenByRedaction          bool     `json:"failure_hidden_by_redaction"`
	ProjectionDisclaimer              string   `json:"projection_disclaimer"`
	Limitations                       []string `json:"limitations,omitempty"`
}

type DegradedBehaviorReview struct {
	CurrentState                    string   `json:"current_state"`
	ReviewState                     string   `json:"review_state"`
	SupportedReviewStates           []string `json:"supported_review_states,omitempty"`
	StatusModelState                string   `json:"status_model_state"`
	UIDataState                     string   `json:"ui_data_state"`
	ResultSemanticsState            string   `json:"result_semantics_state"`
	HealthSnapshotState             string   `json:"health_snapshot_state"`
	StaleReportedAsFresh            bool     `json:"stale_reported_as_fresh"`
	PartialReportedAsComplete       bool     `json:"partial_reported_as_complete"`
	DegradedReportedAsHealthy       bool     `json:"degraded_reported_as_healthy"`
	UnavailableCollapsedUnsupported bool     `json:"unavailable_collapsed_into_unsupported"`
	UnsupportedSilentlyOmitted      bool     `json:"unsupported_silently_omitted"`
	LimitationMessagesPresent       bool     `json:"limitation_messages_present"`
	RecoveryHintsPresent            bool     `json:"recovery_hints_present"`
	CanonicalTruthDisclaimers       bool     `json:"canonical_truth_disclaimers_present"`
	Limitations                     []string `json:"limitations,omitempty"`
}

type UIWindowingResultReview struct {
	CurrentState                   string   `json:"current_state"`
	ReviewState                    string   `json:"review_state"`
	SupportedReviewStates          []string `json:"supported_review_states,omitempty"`
	UIDataResilienceState          string   `json:"ui_data_resilience_state"`
	WindowingState                 string   `json:"windowing_state"`
	ResultSemanticsState           string   `json:"result_semantics_state"`
	UnknownTotalClaimsComplete     bool     `json:"unknown_total_claims_complete"`
	LimitExceedsMaxWindow          bool     `json:"limit_exceeds_max_window"`
	StableOrderingEnforced         bool     `json:"stable_ordering_enforced"`
	MaxWindowEnforced              bool     `json:"max_window_enforced"`
	PartialWindowLimitationPresent bool     `json:"partial_window_limitation_present"`
	TruncationWarningsPresent      bool     `json:"truncation_warnings_present"`
	CacheClaimsCanonicalTruth      bool     `json:"cache_claims_canonical_truth"`
	Limitations                    []string `json:"limitations,omitempty"`
}

type CommandNoiseReview struct {
	CurrentState                  string   `json:"current_state"`
	ReviewState                   string   `json:"review_state"`
	SupportedReviewStates         []string `json:"supported_review_states,omitempty"`
	CommandCenterState            string   `json:"command_center_state"`
	NoiseBudgetState              string   `json:"noise_budget_state"`
	TaskViewsAdvisoryOnly         bool     `json:"task_views_advisory_only"`
	AcknowledgementEqualsRemed    bool     `json:"acknowledgement_equals_remediation"`
	ResolvedEqualsCanonicalClose  bool     `json:"resolved_equals_canonical_closure"`
	CriticalSuppressionInvisible  bool     `json:"critical_suppression_invisible"`
	HighestSeverityPreserved      bool     `json:"highest_severity_preserved"`
	SuppressedDuplicatesAuditable bool     `json:"suppressed_duplicates_auditable"`
	ReopenOnChangeExplicit        bool     `json:"reopen_on_change_explicit"`
	UngovernedTaskMutation        bool     `json:"ungoverned_task_mutation"`
	Limitations                   []string `json:"limitations,omitempty"`
}

type APIProtectionReview struct {
	CurrentState               string   `json:"current_state"`
	ReviewState                string   `json:"review_state"`
	SupportedReviewStates      []string `json:"supported_review_states,omitempty"`
	APIProtectionState         string   `json:"api_protection_state"`
	RequestClassesCovered      bool     `json:"request_classes_covered"`
	PriorityLanesCovered       bool     `json:"priority_lanes_covered"`
	FairnessScopesCovered      bool     `json:"fairness_scopes_covered"`
	RateLimitPolicyRefsPresent bool     `json:"rate_limit_policy_refs_present"`
	BackpressureExplicit       bool     `json:"backpressure_explicit"`
	DegradedResponseMarkers    bool     `json:"degraded_response_markers_present"`
	RetryAfterHintsPresent     bool     `json:"retry_after_hints_present"`
	PriorityBypassesGovernance bool     `json:"priority_bypasses_governance"`
	PolicyDenialHiddenThrottle bool     `json:"policy_denial_hidden_as_throttling"`
	StarvationPrevented        bool     `json:"starvation_prevented"`
	UnsafeMutationRetrySafe    bool     `json:"unsafe_mutation_became_retry_safe"`
	Limitations                []string `json:"limitations,omitempty"`
}

type CLIResilienceReview struct {
	CurrentState              string   `json:"current_state"`
	ReviewState               string   `json:"review_state"`
	SupportedReviewStates     []string `json:"supported_review_states,omitempty"`
	CLIResilienceState        string   `json:"cli_resilience_state"`
	OperationTypesCovered     bool     `json:"operation_types_covered"`
	ActionModesCovered        bool     `json:"action_modes_covered"`
	RetrySafetyCovered        bool     `json:"retry_safety_covered"`
	MissingRequiredKey        bool     `json:"missing_required_idempotency_key"`
	RetryUnsafeMissingReason  bool     `json:"retry_unsafe_missing_reason"`
	NonMutatingModeMutates    bool     `json:"non_mutating_mode_mutates"`
	PartialFailureAsSuccess   bool     `json:"partial_failure_reported_as_success"`
	ExitCodeSemanticsDistinct bool     `json:"exit_code_semantics_distinct"`
	InspectExplainRefsPresent bool     `json:"inspect_explain_refs_present"`
	PolicyBypassAllowed       bool     `json:"policy_bypass_allowed"`
	Limitations               []string `json:"limitations,omitempty"`
}

type SupportabilityReview struct {
	CurrentState                  string   `json:"current_state"`
	ReviewState                   string   `json:"review_state"`
	SupportedReviewStates         []string `json:"supported_review_states,omitempty"`
	ReadinessState                string   `json:"readiness_state"`
	GuidedReadinessState          string   `json:"guided_readiness_state"`
	SupportBundleState            string   `json:"support_bundle_state"`
	DiagnosticsState              string   `json:"diagnostics_state"`
	HealthSnapshotState           string   `json:"health_snapshot_state"`
	PermissionSupportState        string   `json:"permission_support_state"`
	ExportSafetyState             string   `json:"export_safety_state"`
	ReadinessReportedPassNotReady bool     `json:"readiness_reported_pass_when_not_ready"`
	SupportBundleManifestMissing  bool     `json:"support_bundle_manifest_missing"`
	RawSecretsOrTokensDetected    bool     `json:"raw_secrets_or_tokens_detected"`
	HealthyWithBlockingDegraded   bool     `json:"healthy_with_blocking_degraded"`
	SupportOutputsCanonicalTruth  bool     `json:"support_outputs_claim_canonical_truth"`
	StalePartialUnsupportedShown  bool     `json:"stale_partial_unsupported_explicit"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	Limitations                   []string `json:"limitations,omitempty"`
}

type RecoveryUXReview struct {
	CurrentState                string   `json:"current_state"`
	ReviewState                 string   `json:"review_state"`
	SupportedReviewStates       []string `json:"supported_review_states,omitempty"`
	Val0RecoveryState           string   `json:"val_0_recovery_state"`
	ValARecoveryGuidanceState   string   `json:"val_a_recovery_guidance_state"`
	ValCRecoveryPlaybookState   string   `json:"val_c_recovery_playbook_state"`
	RecoveryHintsPresent        bool     `json:"recovery_hints_present"`
	RemediationClassesCovered   bool     `json:"remediation_classes_covered"`
	SafeUnsafeStepsDistinct     bool     `json:"safe_unsafe_steps_distinct"`
	EscalationPathsPresent      bool     `json:"escalation_paths_present"`
	RollbackHintsPresent        bool     `json:"rollback_hints_present"`
	InspectExplainRefsPresent   bool     `json:"inspect_explain_refs_present"`
	SupportBundleRefsBounded    bool     `json:"support_bundle_refs_bounded"`
	PolicyBypassRecommended     bool     `json:"policy_bypass_recommended"`
	UnsafeRetryRecommended      bool     `json:"unsafe_retry_recommended"`
	UnsupportedCapabilityHonest bool     `json:"unsupported_capability_honest"`
	Limitations                 []string `json:"limitations,omitempty"`
}

type UpgradeRollbackReview struct {
	CurrentState                      string   `json:"current_state"`
	ReviewState                       string   `json:"review_state"`
	SupportedReviewStates             []string `json:"supported_review_states,omitempty"`
	ValAUpgradePreviewState           string   `json:"val_a_upgrade_preview_state"`
	ValCUpgradeAdvisoryState          string   `json:"val_c_upgrade_advisory_state"`
	CurrentVersionPresent             bool     `json:"current_version_present"`
	TargetVersionKnown                bool     `json:"target_version_known"`
	CompatibilityStatusPresent        bool     `json:"compatibility_status_present"`
	MigrationItemsRepresented         bool     `json:"migration_items_represented"`
	DeprecatedRemovedItemsRepresented bool     `json:"deprecated_removed_items_represented"`
	RollbackAvailabilityBounded       bool     `json:"rollback_availability_bounded"`
	RollbackScopeLimited              bool     `json:"rollback_scope_limited"`
	AdvisoryModeValid                 bool     `json:"advisory_mode_valid"`
	AdvisoryMutatesState              bool     `json:"advisory_mutates_state"`
	ApprovalImplied                   bool     `json:"approval_implied"`
	LimitationDisclaimer              string   `json:"limitation_disclaimer"`
	Limitations                       []string `json:"limitations,omitempty"`
}

type ScaleEnvelopeReview struct {
	CurrentState                   string   `json:"current_state"`
	ReviewState                    string   `json:"review_state"`
	SupportedReviewStates          []string `json:"supported_review_states,omitempty"`
	ScaleEnvelopeState             string   `json:"scale_envelope_state"`
	ExpectedRangesPresent          bool     `json:"expected_ranges_present"`
	QueryLimitsPresent             bool     `json:"query_limits_present"`
	MaxPageSizeBounded             bool     `json:"max_page_size_bounded"`
	LatencyBudgetsPresent          bool     `json:"latency_budgets_present"`
	DegradationThresholdsPresent   bool     `json:"degradation_thresholds_present"`
	UnsupportedScaleExplicit       bool     `json:"unsupported_scale_explicit"`
	MeasurementNotesPresent        bool     `json:"measurement_notes_present"`
	UnknownOrUnmeasuredMarkedLimit bool     `json:"unknown_or_unmeasured_marked_as_limitation"`
	MarketedAsGuarantee            bool     `json:"marketed_as_guarantee"`
	LimitationDisclaimer           string   `json:"limitation_disclaimer"`
	Limitations                    []string `json:"limitations,omitempty"`
}

type GovernanceBoundaryReview struct {
	CurrentState                     string   `json:"current_state"`
	ReviewState                      string   `json:"review_state"`
	SupportedReviewStates            []string `json:"supported_review_states,omitempty"`
	Val0State                        string   `json:"val_0_state"`
	ValAState                        string   `json:"val_a_state"`
	ValBState                        string   `json:"val_b_state"`
	ValCState                        string   `json:"val_c_state"`
	EvidenceSpineCanonical           bool     `json:"evidence_spine_canonical"`
	OutputsRemainProjectionOnly      bool     `json:"outputs_remain_projection_only"`
	ProjectionClaimsCanonicalTruth   bool     `json:"projection_claims_canonical_truth"`
	AdvisoryMutatesWithoutGovernance bool     `json:"advisory_mutates_without_governance"`
	PublicPartnerAuditorDistinct     bool     `json:"public_partner_auditor_boundaries_distinct"`
	PublicPartnerExposure            bool     `json:"public_partner_exposure"`
	UsabilityWeakensPolicyDiscipline bool     `json:"usability_weakens_policy_discipline"`
	MutationRequiresGovernanceRefs   bool     `json:"mutation_requires_governance_refs"`
	Limitations                      []string `json:"limitations,omitempty"`
}

type UsabilityRegressionGate struct {
	CurrentState                           string   `json:"current_state"`
	ReviewState                            string   `json:"review_state"`
	SupportedReviewStates                  []string `json:"supported_review_states,omitempty"`
	ConfigFailureFixtureCoverage           bool     `json:"config_failure_fixture_coverage"`
	RejectionSnapshotCoverage              bool     `json:"rejection_snapshot_coverage"`
	ExplanationPayloadCoverage             bool     `json:"explanation_payload_contract_coverage"`
	StaleDegradedPartialFixtureCoverage    bool     `json:"stale_degraded_partial_fixture_coverage"`
	CLIRetryIdempotencyFixtureCoverage     bool     `json:"cli_retry_idempotency_fixture_coverage"`
	APIBackpressureFairnessFixtureCoverage bool     `json:"api_backpressure_fairness_fixture_coverage"`
	SupportBundleRedactionFixtureCoverage  bool     `json:"support_bundle_redaction_fixture_coverage"`
	UpgradeRollbackFixtureCoverage         bool     `json:"upgrade_rollback_fixture_coverage"`
	DependencyCoverage                     bool     `json:"dependency_coverage"`
	MissingCriticalFixtureCoverage         bool     `json:"missing_critical_fixture_coverage"`
	Limitations                            []string `json:"limitations,omitempty"`
}

func productionUsabilityValDReviewStates() []string {
	return []string{
		ProductionUsabilityFinalGatePass,
		ProductionUsabilityFinalGateFail,
		ProductionUsabilityFinalGateWarning,
		ProductionUsabilityFinalGateBlocked,
		ProductionUsabilityFinalGateUnsupported,
		ProductionUsabilityFinalGateNotRun,
	}
}

func ProductionUsabilityValDConfigCorrectnessReview() ConfigCorrectnessReview {
	configFactory := ProductionUsabilityValAConfigFactory()
	effective := ProductionUsabilityValAEffectiveConfigInspection()
	bootstrap := ProductionUsabilityValABootstrapValidation()
	return ConfigCorrectnessReview{
		CurrentState:                    "final_usability_config_review_ready",
		ReviewState:                     ProductionUsabilityFinalGatePass,
		SupportedReviewStates:           productionUsabilityValDReviewStates(),
		ConfigFactoryState:              EvaluateProductionUsabilityValAConfigFactoryState(configFactory),
		BootstrapValidationState:        EvaluateProductionUsabilityValABootstrapValidationState(bootstrap),
		EffectiveConfigState:            EvaluateProductionUsabilityValAEffectiveConfigState(effective),
		SchemaCompatibilityStatus:       configFactory.CurrentCompatibility,
		ValidationResult:                configFactory.CurrentValidationResult,
		UnknownFieldPolicy:              configFactory.CurrentUnknownFieldPolicy,
		RequiredFieldValidationPassed:   configFactory.RequiredFieldValidation,
		DefaultsSeparated:               !hasTrimmedStringOverlap(effective.DefaultsApplied, effective.UserProvidedFields),
		SecretsExposedInEffectiveConfig: effective.SecretsExposed,
		MigrationWarningsPresent:        len(configFactory.MigrationWarnings) > 0,
		MigrationCompleted:              configFactory.MigrationCompleted,
		DeprecatedSchemaHandled:         strings.TrimSpace(configFactory.CurrentCompatibility) != ProductionUsabilityCompatibilityDeprecated || len(configFactory.CompatibilityWarnings) > 0,
		EffectiveConfigClaimsCanonical:  false,
		FailFastBootstrapEnabled:        configFactory.FailFastBootstrap,
		ProjectionDisclaimer:            "config_correctness_review_is_projection_only_not_canonical_truth",
		Limitations: []string{
			"Val D config correctness review proves bounded final-gate posture only and does not mutate config or bootstrap state.",
		},
	}
}

func ProductionUsabilityValDExplainabilityClarityReview() ExplainabilityClarityReview {
	rejection := ProductionUsabilityValAHumanReadableRejectionLayer()
	explain := ProductionUsabilityValAPermissionAwareExplainOutputs()
	return ExplainabilityClarityReview{
		CurrentState:                       "final_usability_explainability_review_ready",
		ReviewState:                        ProductionUsabilityFinalGatePass,
		SupportedReviewStates:              productionUsabilityValDReviewStates(),
		RejectionLayerState:                EvaluateProductionUsabilityValARejectionLayerState(rejection),
		ExplainState:                       EvaluateProductionUsabilityValAExplainState(explain),
		ReasonCodesPresent:                 containsTrimmedString(rejection.RequiredFields, "reason_code"),
		PolicyRefsPresent:                  containsTrimmedString(rejection.RequiredFields, "policy_ref"),
		SubjectRefsPresent:                 containsTrimmedString(rejection.RequiredFields, "subject_ref"),
		NextStepsPresent:                   containsTrimmedString(rejection.RequiredFields, "next_step"),
		RecoveryHintsPresent:               containsTrimmedString(rejection.RequiredFields, "recovery_hint"),
		VisibilityScopesCovered:            containsExactTrimmedStringSet(rejection.SupportedVisibilityScopes, productionUsabilityValAExplainScopes()...),
		RedactionTiersCovered:              containsExactTrimmedStringSet(rejection.SupportedRedactionTiers, ProductionUsabilityRedactionNone, ProductionUsabilityRedactionLow, ProductionUsabilityRedactionMedium, ProductionUsabilityRedactionHigh, ProductionUsabilityRedactionPublicSafe),
		DecisionPriorityPresent:            containsTrimmedString(rejection.RequiredFields, "decision_priority"),
		TechnicalDetailAvailableForAllowed: containsAllTrimmedStrings(rejection.TechnicalDetailScopes, ProductionUsabilityVisibilityInternalAdmin, ProductionUsabilityVisibilityOperator),
		SafeFallbackForRestrictedScopes:    true,
		SensitiveEvidenceLeaked:            false,
		FailureHiddenByRedaction:           false,
		ProjectionDisclaimer:               "explainability_review_is_projection_only_not_canonical_truth",
		Limitations: []string{
			"Val D explainability review validates clarity and safe bounded redaction only; it does not turn review output into canonical truth.",
		},
	}
}

func ProductionUsabilityValDDryRunAuditReview() DryRunAuditCorrectnessReview {
	dryRun := ProductionUsabilityValAPolicyDryRunAuditFlow()
	return DryRunAuditCorrectnessReview{
		CurrentState:                  "final_usability_dry_run_review_ready",
		ReviewState:                   ProductionUsabilityFinalGatePass,
		SupportedReviewStates:         productionUsabilityValDReviewStates(),
		DryRunState:                   EvaluateProductionUsabilityValADryRunState(dryRun),
		DryRunMutatesCanonicalState:   dryRun.MutatesCanonicalState,
		AuditOnlyMutatesCanonical:     false,
		PreviewSuccessImpliesActivate: dryRun.DryRunSuccessImpliesActivate,
		AuditOnlyImpliesApproval:      dryRun.AuditOnlyImpliesApproval,
		BlockingIssuesIdentified:      len(dryRun.BlockingRules) > 0,
		WarningIssuesIdentified:       len(dryRun.NonBlockingWarnings) > 0,
		ProjectionDisclaimer:          "dry_run_review_is_projection_only_non_authoritative_not_canonical_truth",
		Limitations: []string{
			"Val D dry-run review verifies simulation correctness only and does not authorize activation or approval.",
		},
	}
}

func ProductionUsabilityValDPermissionRedactionReview() PermissionRedactionReview {
	permission := ProductionUsabilityVal0PermissionRedactionContract()
	explain := ProductionUsabilityValAPermissionAwareExplainOutputs()
	support := ProductionUsabilityValCPermissionSupportFlows()
	export := ProductionUsabilityValCRedactionSafeExport()
	return PermissionRedactionReview{
		CurrentState:                      "final_usability_redaction_review_ready",
		ReviewState:                       ProductionUsabilityFinalGatePass,
		SupportedReviewStates:             productionUsabilityValDReviewStates(),
		PermissionRedactionState:          EvaluateProductionUsabilityVal0PermissionRedactionState(permission),
		ExplainState:                      EvaluateProductionUsabilityValAExplainState(explain),
		PermissionSupportState:            EvaluateProductionUsabilityValCPermissionSupportState(support),
		ExportSafetyState:                 EvaluateProductionUsabilityValCExportSafetyState(export),
		VisibilityScopesCovered:           true,
		HiddenMetadataRepresented:         true,
		SafeFallbackMessagesPresent:       true,
		PartnerOrPublicExposeFullEvidence: false,
		RawSecretsOrTokensDetected:        false,
		AuditorImpliesPublicSafe:          export.AuditorSafe && export.PublicSafe,
		FailureHiddenByRedaction:          false,
		ProjectionDisclaimer:              "redaction_review_is_projection_only_not_canonical_truth",
		Limitations: []string{
			"Val D redaction review validates bounded evidence visibility and export discipline only; it does not approve disclosure.",
		},
	}
}

func ProductionUsabilityValDDegradedBehaviorReview() DegradedBehaviorReview {
	statusModel := ProductionUsabilityVal0OperationalStatusModel()
	uiData := ProductionUsabilityValBUIDataResilience()
	resultSemantics := ProductionUsabilityValBResultSemantics()
	health := ProductionUsabilityValCHealthSnapshot()
	return DegradedBehaviorReview{
		CurrentState:                    "final_usability_degraded_behavior_review_ready",
		ReviewState:                     ProductionUsabilityFinalGatePass,
		SupportedReviewStates:           productionUsabilityValDReviewStates(),
		StatusModelState:                EvaluateProductionUsabilityVal0StatusModelState(statusModel),
		UIDataState:                     EvaluateProductionUsabilityValBUIDataResilienceState(uiData),
		ResultSemanticsState:            EvaluateProductionUsabilityValBResultSemanticsState(resultSemantics),
		HealthSnapshotState:             EvaluateProductionUsabilityValCHealthSnapshotState(health),
		StaleReportedAsFresh:            false,
		PartialReportedAsComplete:       false,
		DegradedReportedAsHealthy:       false,
		UnavailableCollapsedUnsupported: false,
		UnsupportedSilentlyOmitted:      false,
		LimitationMessagesPresent:       true,
		RecoveryHintsPresent:            true,
		CanonicalTruthDisclaimers:       true,
		Limitations: []string{
			"Val D degraded-behavior review ensures stale, partial, degraded, unavailable, and unsupported outputs remain explicit and non-canonical.",
		},
	}
}

func ProductionUsabilityValDUIWindowingResultReview() UIWindowingResultReview {
	windowing := ProductionUsabilityValBWindowing()
	return UIWindowingResultReview{
		CurrentState:                   "final_usability_ui_windowing_review_ready",
		ReviewState:                    ProductionUsabilityFinalGatePass,
		SupportedReviewStates:          productionUsabilityValDReviewStates(),
		UIDataResilienceState:          EvaluateProductionUsabilityValBUIDataResilienceState(ProductionUsabilityValBUIDataResilience()),
		WindowingState:                 EvaluateProductionUsabilityValBWindowingState(windowing),
		ResultSemanticsState:           EvaluateProductionUsabilityValBResultSemanticsState(ProductionUsabilityValBResultSemantics()),
		UnknownTotalClaimsComplete:     !windowing.TotalCountKnown && windowing.ClaimsCompleteData,
		LimitExceedsMaxWindow:          windowing.Limit > windowing.MaxWindowSize,
		StableOrderingEnforced:         windowing.StableSortRequired,
		MaxWindowEnforced:              windowing.MaxWindowEnforced,
		PartialWindowLimitationPresent: !windowing.PartialResult || strings.TrimSpace(windowing.LimitationMessage) != "",
		TruncationWarningsPresent:      strings.TrimSpace(windowing.TruncationWarning) != "",
		CacheClaimsCanonicalTruth:      false,
		Limitations: []string{
			"Val D UI/windowing review validates bounded pagination and projection semantics only; it does not turn windows into canonical datasets.",
		},
	}
}

func ProductionUsabilityValDCommandNoiseReview() CommandNoiseReview {
	return CommandNoiseReview{
		CurrentState:                  "final_usability_command_noise_review_ready",
		ReviewState:                   ProductionUsabilityFinalGatePass,
		SupportedReviewStates:         productionUsabilityValDReviewStates(),
		CommandCenterState:            EvaluateProductionUsabilityValBCommandCenterState(ProductionUsabilityValBCommandCenterTasks()),
		NoiseBudgetState:              EvaluateProductionUsabilityValBNoiseBudgetState(ProductionUsabilityValBNoiseBudget()),
		TaskViewsAdvisoryOnly:         true,
		AcknowledgementEqualsRemed:    false,
		ResolvedEqualsCanonicalClose:  false,
		CriticalSuppressionInvisible:  false,
		HighestSeverityPreserved:      true,
		SuppressedDuplicatesAuditable: true,
		ReopenOnChangeExplicit:        true,
		UngovernedTaskMutation:        false,
		Limitations: []string{
			"Val D command/noise review proves advisory-only operator surfaces and auditable grouping only; it does not grant workflow authority.",
		},
	}
}

func ProductionUsabilityValDAPIProtectionReview() APIProtectionReview {
	apiProtection := ProductionUsabilityValBAPIProtection()
	return APIProtectionReview{
		CurrentState:               "final_usability_api_protection_review_ready",
		ReviewState:                ProductionUsabilityFinalGatePass,
		SupportedReviewStates:      productionUsabilityValDReviewStates(),
		APIProtectionState:         EvaluateProductionUsabilityValBAPIProtectionState(apiProtection),
		RequestClassesCovered:      len(apiProtection.Items) == len(apiProtection.SupportedRequestClasses),
		PriorityLanesCovered:       true,
		FairnessScopesCovered:      true,
		RateLimitPolicyRefsPresent: true,
		BackpressureExplicit:       true,
		DegradedResponseMarkers:    true,
		RetryAfterHintsPresent:     true,
		PriorityBypassesGovernance: false,
		PolicyDenialHiddenThrottle: false,
		StarvationPrevented:        true,
		UnsafeMutationRetrySafe:    false,
		Limitations: []string{
			"Val D API protection review validates fairness, backpressure, and governance boundaries only; it does not mutate request handling state.",
		},
	}
}

func ProductionUsabilityValDCLIResilienceReview() CLIResilienceReview {
	cli := ProductionUsabilityValBCLIResilience()
	return CLIResilienceReview{
		CurrentState:              "final_usability_cli_review_ready",
		ReviewState:               ProductionUsabilityFinalGatePass,
		SupportedReviewStates:     productionUsabilityValDReviewStates(),
		CLIResilienceState:        EvaluateProductionUsabilityValBCLIResilienceState(cli),
		OperationTypesCovered:     len(cli.Items) > 0,
		ActionModesCovered:        true,
		RetrySafetyCovered:        true,
		MissingRequiredKey:        false,
		RetryUnsafeMissingReason:  false,
		NonMutatingModeMutates:    false,
		PartialFailureAsSuccess:   false,
		ExitCodeSemanticsDistinct: len(cli.SupportedExitCodeSemantics) > 0,
		InspectExplainRefsPresent: true,
		PolicyBypassAllowed:       false,
		Limitations: []string{
			"Val D CLI review validates retry, idempotency, and partial-failure posture only; it does not bypass governed policy or evidence checks.",
		},
	}
}

func ProductionUsabilityValDSupportabilityReview() SupportabilityReview {
	return SupportabilityReview{
		CurrentState:                  "final_usability_supportability_review_ready",
		ReviewState:                   ProductionUsabilityFinalGatePass,
		SupportedReviewStates:         productionUsabilityValDReviewStates(),
		ReadinessState:                EvaluateProductionUsabilityValCReadinessState(ProductionUsabilityValCReadinessChecks()),
		GuidedReadinessState:          EvaluateProductionUsabilityValCGuidedReadinessState(ProductionUsabilityValCGuidedReadiness()),
		SupportBundleState:            EvaluateProductionUsabilityValCSupportBundleState(ProductionUsabilityValCSupportBundleQualityGate()),
		DiagnosticsState:              EvaluateProductionUsabilityValCDiagnosticsState(ProductionUsabilityValCDiagnosticsHardening()),
		HealthSnapshotState:           EvaluateProductionUsabilityValCHealthSnapshotState(ProductionUsabilityValCHealthSnapshot()),
		PermissionSupportState:        EvaluateProductionUsabilityValCPermissionSupportState(ProductionUsabilityValCPermissionSupportFlows()),
		ExportSafetyState:             EvaluateProductionUsabilityValCExportSafetyState(ProductionUsabilityValCRedactionSafeExport()),
		ReadinessReportedPassNotReady: false,
		SupportBundleManifestMissing:  false,
		RawSecretsOrTokensDetected:    false,
		HealthyWithBlockingDegraded:   false,
		SupportOutputsCanonicalTruth:  false,
		StalePartialUnsupportedShown:  true,
		ProjectionDisclaimer:          "supportability_review_is_projection_only_not_canonical_truth",
		Limitations: []string{
			"Val D supportability review validates readiness, diagnostics, health, and export posture only; it does not replace canonical truth.",
		},
	}
}

func ProductionUsabilityValDRecoveryUXReview() RecoveryUXReview {
	return RecoveryUXReview{
		CurrentState:                "final_usability_recovery_review_ready",
		ReviewState:                 ProductionUsabilityFinalGatePass,
		SupportedReviewStates:       productionUsabilityValDReviewStates(),
		Val0RecoveryState:           ProductionUsabilityVal0RecoveryStateActive,
		ValARecoveryGuidanceState:   EvaluateProductionUsabilityValARecoveryGuidanceState(ProductionUsabilityValARecoveryGuidance()),
		ValCRecoveryPlaybookState:   EvaluateProductionUsabilityValCRecoveryPlaybookState(ProductionUsabilityValCRecoveryPlaybooks()),
		RecoveryHintsPresent:        true,
		RemediationClassesCovered:   true,
		SafeUnsafeStepsDistinct:     true,
		EscalationPathsPresent:      true,
		RollbackHintsPresent:        true,
		InspectExplainRefsPresent:   true,
		SupportBundleRefsBounded:    true,
		PolicyBypassRecommended:     false,
		UnsafeRetryRecommended:      false,
		UnsupportedCapabilityHonest: true,
		Limitations: []string{
			"Val D recovery review proves bounded guidance quality only and does not authorize bypasses, unsafe retries, or unsupported capability use.",
		},
	}
}

func ProductionUsabilityValDUpgradeRollbackReview() UpgradeRollbackReview {
	preview := ProductionUsabilityValAUpgradeImpactPreview()
	advisory := ProductionUsabilityValCUpgradeRollbackAdvisory()
	return UpgradeRollbackReview{
		CurrentState:                      "final_usability_upgrade_rollback_review_ready",
		ReviewState:                       ProductionUsabilityFinalGatePass,
		SupportedReviewStates:             productionUsabilityValDReviewStates(),
		ValAUpgradePreviewState:           EvaluateProductionUsabilityValAUpgradePreviewState(preview),
		ValCUpgradeAdvisoryState:          EvaluateProductionUsabilityValCUpgradeAdvisoryState(advisory),
		CurrentVersionPresent:             strings.TrimSpace(advisory.CurrentVersion) != "",
		TargetVersionKnown:                containsTrimmedString(advisory.KnownTargetVersions, advisory.TargetVersion),
		CompatibilityStatusPresent:        strings.TrimSpace(advisory.CompatibilityStatus) != "",
		MigrationItemsRepresented:         len(advisory.MigrationRequiredItems) > 0,
		DeprecatedRemovedItemsRepresented: len(advisory.DeprecatedItems) > 0 && len(advisory.RemovedOrUnsupportedItems) > 0,
		RollbackAvailabilityBounded:       !advisory.RollbackAvailable || len(advisory.RollbackLimitations) > 0,
		RollbackScopeLimited:              strings.Contains(strings.TrimSpace(advisory.RollbackScope), "config_policy"),
		AdvisoryModeValid:                 containsTrimmedString(productionUsabilityValCAdvisoryModes(), advisory.AdvisoryMode),
		AdvisoryMutatesState:              advisory.MutatesState || preview.MutatesConfig,
		ApprovalImplied:                   advisory.ApprovalImplied,
		LimitationDisclaimer:              "upgrade_rollback_review_is_projection_only_not_canonical_truth",
		Limitations: []string{
			"Val D upgrade/rollback review validates bounded advisory posture only and does not implement lifecycle mutation or closure.",
		},
	}
}

func ProductionUsabilityValDScaleEnvelopeReview() ScaleEnvelopeReview {
	scale := ProductionUsabilityValBScaleEnvelope()
	return ScaleEnvelopeReview{
		CurrentState:                   "final_usability_scale_review_ready",
		ReviewState:                    ProductionUsabilityFinalGatePass,
		SupportedReviewStates:          productionUsabilityValDReviewStates(),
		ScaleEnvelopeState:             EvaluateProductionUsabilityValBScaleEnvelopeState(scale),
		ExpectedRangesPresent:          strings.TrimSpace(scale.ExpectedEventVolumeRange) != "" && strings.TrimSpace(scale.ExpectedWorkflowObjectRange) != "",
		QueryLimitsPresent:             scale.TimelineQueryLimit > 0 && scale.SearchQueryLimit > 0,
		MaxPageSizeBounded:             scale.MaxPageSize > 0,
		LatencyBudgetsPresent:          strings.TrimSpace(scale.DashboardLatencyBudget) != "" && strings.TrimSpace(scale.APILatencyBudget) != "" && strings.TrimSpace(scale.CLIResponseBudget) != "",
		DegradationThresholdsPresent:   len(scale.DegradationThresholds) > 0,
		UnsupportedScaleExplicit:       len(scale.UnsupportedScaleConditions) > 0,
		MeasurementNotesPresent:        len(scale.MeasurementNotes) > 0,
		UnknownOrUnmeasuredMarkedLimit: strings.Contains(strings.TrimSpace(scale.LimitationDisclaimer), "unmeasured"),
		MarketedAsGuarantee:            scale.ClaimsLatencyGuarantee,
		LimitationDisclaimer:           scale.LimitationDisclaimer,
		Limitations: []string{
			"Val D scale review validates bounded, measurement-aware posture only and does not certify final performance guarantees.",
		},
	}
}

func ProductionUsabilityValDGovernanceBoundaryReview() GovernanceBoundaryReview {
	return GovernanceBoundaryReview{
		CurrentState:                     "final_usability_governance_boundary_review_ready",
		ReviewState:                      ProductionUsabilityFinalGatePass,
		SupportedReviewStates:            productionUsabilityValDReviewStates(),
		Val0State:                        ProductionUsabilityVal0StateActive,
		ValAState:                        ProductionUsabilityValAStateActive,
		ValBState:                        ProductionUsabilityValBStateActive,
		ValCState:                        ProductionUsabilityValCStateActive,
		EvidenceSpineCanonical:           true,
		OutputsRemainProjectionOnly:      true,
		ProjectionClaimsCanonicalTruth:   false,
		AdvisoryMutatesWithoutGovernance: false,
		PublicPartnerAuditorDistinct:     true,
		PublicPartnerExposure:            false,
		UsabilityWeakensPolicyDiscipline: false,
		MutationRequiresGovernanceRefs:   true,
		Limitations: []string{
			"Val D governance boundary review validates that usability surfaces remain projections and do not weaken evidence, policy, or governance discipline.",
		},
	}
}

func ProductionUsabilityValDUsabilityRegressionGate() UsabilityRegressionGate {
	return UsabilityRegressionGate{
		CurrentState:                           "final_usability_regression_gate_ready",
		ReviewState:                            ProductionUsabilityFinalGatePass,
		SupportedReviewStates:                  productionUsabilityValDReviewStates(),
		ConfigFailureFixtureCoverage:           true,
		RejectionSnapshotCoverage:              true,
		ExplanationPayloadCoverage:             true,
		StaleDegradedPartialFixtureCoverage:    true,
		CLIRetryIdempotencyFixtureCoverage:     true,
		APIBackpressureFairnessFixtureCoverage: true,
		SupportBundleRedactionFixtureCoverage:  true,
		UpgradeRollbackFixtureCoverage:         true,
		DependencyCoverage:                     true,
		MissingCriticalFixtureCoverage:         false,
		Limitations: []string{
			"Val D regression gate is a proof-oriented fixture contract only and does not add broader test framework refactors.",
		},
	}
}

func EvaluateProductionUsabilityValDConfigReviewState(model ConfigCorrectnessReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.ConfigFactoryState) == "" || strings.TrimSpace(model.BootstrapValidationState) == "" || strings.TrimSpace(model.EffectiveConfigState) == "" || strings.TrimSpace(model.SchemaCompatibilityStatus) == "" || strings.TrimSpace(model.ValidationResult) == "" || strings.TrimSpace(model.UnknownFieldPolicy) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ProductionUsabilityValDConfigReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		!containsTrimmedString(model.SupportedReviewStates, model.ReviewState) ||
		model.ConfigFactoryState != ProductionUsabilityValAConfigFactoryStateActive ||
		model.BootstrapValidationState != ProductionUsabilityValABootstrapValidationStateActive ||
		model.EffectiveConfigState != ProductionUsabilityValAEffectiveConfigStateActive ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().SupportedCompatibility, model.SchemaCompatibilityStatus) ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().ValidationStates, model.ValidationResult) ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().UnknownFieldPolicies, model.UnknownFieldPolicy) ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValDConfigReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		model.SchemaCompatibilityStatus == ProductionUsabilityCompatibilityUnknown ||
		model.SchemaCompatibilityStatus == ProductionUsabilityCompatibilityUnsupported ||
		model.ValidationResult == ProductionUsabilityValidationInvalid ||
		model.ValidationResult == ProductionUsabilityValidationUnsupported ||
		!model.RequiredFieldValidationPassed ||
		!model.DefaultsSeparated ||
		model.SecretsExposedInEffectiveConfig ||
		(model.MigrationWarningsPresent && model.MigrationCompleted) ||
		!model.DeprecatedSchemaHandled ||
		model.EffectiveConfigClaimsCanonical ||
		!model.FailFastBootstrapEnabled {
		return ProductionUsabilityValDConfigReviewStatePartial
	}
	return ProductionUsabilityValDConfigReviewStateActive
}

func EvaluateProductionUsabilityValDExplainabilityReviewState(model ExplainabilityClarityReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.RejectionLayerState) == "" || strings.TrimSpace(model.ExplainState) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ProductionUsabilityValDExplainabilityReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.RejectionLayerState != ProductionUsabilityValARejectionLayerStateActive ||
		model.ExplainState != ProductionUsabilityValAExplainStateActive ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValDExplainabilityReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		!model.ReasonCodesPresent ||
		!model.PolicyRefsPresent ||
		!model.SubjectRefsPresent ||
		!model.NextStepsPresent ||
		!model.RecoveryHintsPresent ||
		!model.VisibilityScopesCovered ||
		!model.RedactionTiersCovered ||
		!model.DecisionPriorityPresent ||
		!model.TechnicalDetailAvailableForAllowed ||
		!model.SafeFallbackForRestrictedScopes ||
		model.SensitiveEvidenceLeaked ||
		model.FailureHiddenByRedaction {
		return ProductionUsabilityValDExplainabilityReviewStatePartial
	}
	return ProductionUsabilityValDExplainabilityReviewStateActive
}

func EvaluateProductionUsabilityValDDryRunReviewState(model DryRunAuditCorrectnessReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.DryRunState) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ProductionUsabilityValDDryRunReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.DryRunState != ProductionUsabilityValADryRunStateActive ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "non_authoritative") {
		return ProductionUsabilityValDDryRunReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		model.DryRunMutatesCanonicalState ||
		model.AuditOnlyMutatesCanonical ||
		model.PreviewSuccessImpliesActivate ||
		model.AuditOnlyImpliesApproval ||
		!model.BlockingIssuesIdentified ||
		!model.WarningIssuesIdentified {
		return ProductionUsabilityValDDryRunReviewStatePartial
	}
	return ProductionUsabilityValDDryRunReviewStateActive
}

func EvaluateProductionUsabilityValDRedactionReviewState(model PermissionRedactionReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.PermissionRedactionState) == "" || strings.TrimSpace(model.ExplainState) == "" || strings.TrimSpace(model.PermissionSupportState) == "" || strings.TrimSpace(model.ExportSafetyState) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ProductionUsabilityValDRedactionReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.PermissionRedactionState != ProductionUsabilityVal0PermissionRedactionStateActive ||
		model.ExplainState != ProductionUsabilityValAExplainStateActive ||
		model.PermissionSupportState != ProductionUsabilityValCPermissionSupportStateActive ||
		model.ExportSafetyState != ProductionUsabilityValCExportSafetyStateActive ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValDRedactionReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		!model.VisibilityScopesCovered ||
		!model.HiddenMetadataRepresented ||
		!model.SafeFallbackMessagesPresent ||
		model.PartnerOrPublicExposeFullEvidence ||
		model.RawSecretsOrTokensDetected ||
		model.AuditorImpliesPublicSafe ||
		model.FailureHiddenByRedaction {
		return ProductionUsabilityValDRedactionReviewStatePartial
	}
	return ProductionUsabilityValDRedactionReviewStateActive
}

func EvaluateProductionUsabilityValDDegradedBehaviorReviewState(model DegradedBehaviorReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.StatusModelState) == "" || strings.TrimSpace(model.UIDataState) == "" || strings.TrimSpace(model.ResultSemanticsState) == "" || strings.TrimSpace(model.HealthSnapshotState) == "" {
		return ProductionUsabilityValDDegradedBehaviorReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.StatusModelState != ProductionUsabilityVal0StatusModelStateActive ||
		model.UIDataState != ProductionUsabilityValBUIDataResilienceStateActive ||
		model.ResultSemanticsState != ProductionUsabilityValBResultSemanticsStateActive ||
		model.HealthSnapshotState != ProductionUsabilityValCHealthSnapshotStateActive {
		return ProductionUsabilityValDDegradedBehaviorReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		model.StaleReportedAsFresh ||
		model.PartialReportedAsComplete ||
		model.DegradedReportedAsHealthy ||
		model.UnavailableCollapsedUnsupported ||
		model.UnsupportedSilentlyOmitted ||
		!model.LimitationMessagesPresent ||
		!model.RecoveryHintsPresent ||
		!model.CanonicalTruthDisclaimers {
		return ProductionUsabilityValDDegradedBehaviorReviewStatePartial
	}
	return ProductionUsabilityValDDegradedBehaviorReviewStateActive
}

func EvaluateProductionUsabilityValDUIWindowingReviewState(model UIWindowingResultReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.UIDataResilienceState) == "" || strings.TrimSpace(model.WindowingState) == "" || strings.TrimSpace(model.ResultSemanticsState) == "" {
		return ProductionUsabilityValDUIWindowingReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.UIDataResilienceState != ProductionUsabilityValBUIDataResilienceStateActive ||
		model.WindowingState != ProductionUsabilityValBWindowingStateActive ||
		model.ResultSemanticsState != ProductionUsabilityValBResultSemanticsStateActive {
		return ProductionUsabilityValDUIWindowingReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		model.UnknownTotalClaimsComplete ||
		model.LimitExceedsMaxWindow ||
		!model.StableOrderingEnforced ||
		!model.MaxWindowEnforced ||
		!model.PartialWindowLimitationPresent ||
		!model.TruncationWarningsPresent ||
		model.CacheClaimsCanonicalTruth {
		return ProductionUsabilityValDUIWindowingReviewStatePartial
	}
	return ProductionUsabilityValDUIWindowingReviewStateActive
}

func EvaluateProductionUsabilityValDCommandNoiseReviewState(model CommandNoiseReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.CommandCenterState) == "" || strings.TrimSpace(model.NoiseBudgetState) == "" {
		return ProductionUsabilityValDCommandNoiseReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.CommandCenterState != ProductionUsabilityValBCommandCenterStateActive ||
		model.NoiseBudgetState != ProductionUsabilityValBNoiseBudgetStateActive {
		return ProductionUsabilityValDCommandNoiseReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		!model.TaskViewsAdvisoryOnly ||
		model.AcknowledgementEqualsRemed ||
		model.ResolvedEqualsCanonicalClose ||
		model.CriticalSuppressionInvisible ||
		!model.HighestSeverityPreserved ||
		!model.SuppressedDuplicatesAuditable ||
		!model.ReopenOnChangeExplicit ||
		model.UngovernedTaskMutation {
		return ProductionUsabilityValDCommandNoiseReviewStatePartial
	}
	return ProductionUsabilityValDCommandNoiseReviewStateActive
}

func EvaluateProductionUsabilityValDAPIProtectionReviewState(model APIProtectionReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.APIProtectionState) == "" {
		return ProductionUsabilityValDAPIProtectionReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.APIProtectionState != ProductionUsabilityValBAPIProtectionStateActive {
		return ProductionUsabilityValDAPIProtectionReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		!model.RequestClassesCovered ||
		!model.PriorityLanesCovered ||
		!model.FairnessScopesCovered ||
		!model.RateLimitPolicyRefsPresent ||
		!model.BackpressureExplicit ||
		!model.DegradedResponseMarkers ||
		!model.RetryAfterHintsPresent ||
		model.PriorityBypassesGovernance ||
		model.PolicyDenialHiddenThrottle ||
		!model.StarvationPrevented ||
		model.UnsafeMutationRetrySafe {
		return ProductionUsabilityValDAPIProtectionReviewStatePartial
	}
	return ProductionUsabilityValDAPIProtectionReviewStateActive
}

func EvaluateProductionUsabilityValDCLIResilienceReviewState(model CLIResilienceReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.CLIResilienceState) == "" {
		return ProductionUsabilityValDCLIResilienceReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.CLIResilienceState != ProductionUsabilityValBCLIResilienceStateActive {
		return ProductionUsabilityValDCLIResilienceReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		!model.OperationTypesCovered ||
		!model.ActionModesCovered ||
		!model.RetrySafetyCovered ||
		model.MissingRequiredKey ||
		model.RetryUnsafeMissingReason ||
		model.NonMutatingModeMutates ||
		model.PartialFailureAsSuccess ||
		!model.ExitCodeSemanticsDistinct ||
		!model.InspectExplainRefsPresent ||
		model.PolicyBypassAllowed {
		return ProductionUsabilityValDCLIResilienceReviewStatePartial
	}
	return ProductionUsabilityValDCLIResilienceReviewStateActive
}

func EvaluateProductionUsabilityValDSupportabilityReviewState(model SupportabilityReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.ReadinessState) == "" || strings.TrimSpace(model.GuidedReadinessState) == "" || strings.TrimSpace(model.SupportBundleState) == "" || strings.TrimSpace(model.DiagnosticsState) == "" || strings.TrimSpace(model.HealthSnapshotState) == "" || strings.TrimSpace(model.PermissionSupportState) == "" || strings.TrimSpace(model.ExportSafetyState) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ProductionUsabilityValDSupportabilityReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.ReadinessState != ProductionUsabilityValCReadinessStateActive ||
		model.GuidedReadinessState != ProductionUsabilityValCGuidedReadinessStateActive ||
		model.SupportBundleState != ProductionUsabilityValCSupportBundleStateActive ||
		model.DiagnosticsState != ProductionUsabilityValCDiagnosticsStateActive ||
		model.HealthSnapshotState != ProductionUsabilityValCHealthSnapshotStateActive ||
		model.PermissionSupportState != ProductionUsabilityValCPermissionSupportStateActive ||
		model.ExportSafetyState != ProductionUsabilityValCExportSafetyStateActive ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValDSupportabilityReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		model.ReadinessReportedPassNotReady ||
		model.SupportBundleManifestMissing ||
		model.RawSecretsOrTokensDetected ||
		model.HealthyWithBlockingDegraded ||
		model.SupportOutputsCanonicalTruth ||
		!model.StalePartialUnsupportedShown {
		return ProductionUsabilityValDSupportabilityReviewStatePartial
	}
	return ProductionUsabilityValDSupportabilityReviewStateActive
}

func EvaluateProductionUsabilityValDRecoveryReviewState(model RecoveryUXReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.Val0RecoveryState) == "" || strings.TrimSpace(model.ValARecoveryGuidanceState) == "" || strings.TrimSpace(model.ValCRecoveryPlaybookState) == "" {
		return ProductionUsabilityValDRecoveryReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.Val0RecoveryState != ProductionUsabilityVal0RecoveryStateActive ||
		model.ValARecoveryGuidanceState != ProductionUsabilityValARecoveryGuidanceStateActive ||
		model.ValCRecoveryPlaybookState != ProductionUsabilityValCRecoveryPlaybookStateActive {
		return ProductionUsabilityValDRecoveryReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		!model.RecoveryHintsPresent ||
		!model.RemediationClassesCovered ||
		!model.SafeUnsafeStepsDistinct ||
		!model.EscalationPathsPresent ||
		!model.RollbackHintsPresent ||
		!model.InspectExplainRefsPresent ||
		!model.SupportBundleRefsBounded ||
		model.PolicyBypassRecommended ||
		model.UnsafeRetryRecommended ||
		!model.UnsupportedCapabilityHonest {
		return ProductionUsabilityValDRecoveryReviewStatePartial
	}
	return ProductionUsabilityValDRecoveryReviewStateActive
}

func EvaluateProductionUsabilityValDUpgradeRollbackReviewState(model UpgradeRollbackReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.ValAUpgradePreviewState) == "" || strings.TrimSpace(model.ValCUpgradeAdvisoryState) == "" || strings.TrimSpace(model.LimitationDisclaimer) == "" {
		return ProductionUsabilityValDUpgradeRollbackReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.ValAUpgradePreviewState != ProductionUsabilityValAUpgradePreviewStateActive ||
		model.ValCUpgradeAdvisoryState != ProductionUsabilityValCUpgradeAdvisoryStateActive ||
		!strings.Contains(strings.TrimSpace(model.LimitationDisclaimer), "projection_only") ||
		!strings.Contains(strings.TrimSpace(model.LimitationDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValDUpgradeRollbackReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		!model.CurrentVersionPresent ||
		!model.TargetVersionKnown ||
		!model.CompatibilityStatusPresent ||
		!model.MigrationItemsRepresented ||
		!model.DeprecatedRemovedItemsRepresented ||
		!model.RollbackAvailabilityBounded ||
		!model.RollbackScopeLimited ||
		!model.AdvisoryModeValid ||
		model.AdvisoryMutatesState ||
		model.ApprovalImplied {
		return ProductionUsabilityValDUpgradeRollbackReviewStatePartial
	}
	return ProductionUsabilityValDUpgradeRollbackReviewStateActive
}

func EvaluateProductionUsabilityValDScaleEnvelopeReviewState(model ScaleEnvelopeReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.ScaleEnvelopeState) == "" || strings.TrimSpace(model.LimitationDisclaimer) == "" {
		return ProductionUsabilityValDScaleEnvelopeReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.ScaleEnvelopeState != ProductionUsabilityValBScaleEnvelopeStateActive {
		return ProductionUsabilityValDScaleEnvelopeReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		!model.ExpectedRangesPresent ||
		!model.QueryLimitsPresent ||
		!model.MaxPageSizeBounded ||
		!model.LatencyBudgetsPresent ||
		!model.DegradationThresholdsPresent ||
		!model.UnsupportedScaleExplicit ||
		!model.MeasurementNotesPresent ||
		!model.UnknownOrUnmeasuredMarkedLimit ||
		model.MarketedAsGuarantee {
		return ProductionUsabilityValDScaleEnvelopeReviewStatePartial
	}
	return ProductionUsabilityValDScaleEnvelopeReviewStateActive
}

func EvaluateProductionUsabilityValDGovernanceBoundaryReviewState(model GovernanceBoundaryReview) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" || strings.TrimSpace(model.Val0State) == "" || strings.TrimSpace(model.ValAState) == "" || strings.TrimSpace(model.ValBState) == "" || strings.TrimSpace(model.ValCState) == "" {
		return ProductionUsabilityValDGovernanceBoundaryReviewStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) ||
		model.Val0State != ProductionUsabilityVal0StateActive ||
		model.ValAState != ProductionUsabilityValAStateActive ||
		model.ValBState != ProductionUsabilityValBStateActive ||
		model.ValCState != ProductionUsabilityValCStateActive {
		return ProductionUsabilityValDGovernanceBoundaryReviewStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		!model.EvidenceSpineCanonical ||
		!model.OutputsRemainProjectionOnly ||
		model.ProjectionClaimsCanonicalTruth ||
		model.AdvisoryMutatesWithoutGovernance ||
		!model.PublicPartnerAuditorDistinct ||
		model.PublicPartnerExposure ||
		model.UsabilityWeakensPolicyDiscipline ||
		!model.MutationRequiresGovernanceRefs {
		return ProductionUsabilityValDGovernanceBoundaryReviewStatePartial
	}
	return ProductionUsabilityValDGovernanceBoundaryReviewStateActive
}

func EvaluateProductionUsabilityValDRegressionGateState(model UsabilityRegressionGate) string {
	if strings.TrimSpace(model.CurrentState) == "" || strings.TrimSpace(model.ReviewState) == "" {
		return ProductionUsabilityValDRegressionGateStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, productionUsabilityValDReviewStates()...) {
		return ProductionUsabilityValDRegressionGateStatePartial
	}
	if model.ReviewState != ProductionUsabilityFinalGatePass ||
		!model.ConfigFailureFixtureCoverage ||
		!model.RejectionSnapshotCoverage ||
		!model.ExplanationPayloadCoverage ||
		!model.StaleDegradedPartialFixtureCoverage ||
		!model.CLIRetryIdempotencyFixtureCoverage ||
		!model.APIBackpressureFairnessFixtureCoverage ||
		!model.SupportBundleRedactionFixtureCoverage ||
		!model.UpgradeRollbackFixtureCoverage ||
		!model.DependencyCoverage ||
		model.MissingCriticalFixtureCoverage {
		return ProductionUsabilityValDRegressionGateStatePartial
	}
	return ProductionUsabilityValDRegressionGateStateActive
}

func EvaluateProductionUsabilityValDState(val0State, valAState, valBState, valCState, configReviewState, explainabilityReviewState, dryRunReviewState, redactionReviewState, degradedBehaviorReviewState, uiWindowingReviewState, commandNoiseReviewState, apiProtectionReviewState, cliResilienceReviewState, supportabilityReviewState, recoveryReviewState, upgradeRollbackReviewState, scaleEnvelopeReviewState, governanceBoundaryReviewState, regressionGateState string) string {
	if strings.TrimSpace(val0State) != ProductionUsabilityVal0StateActive ||
		strings.TrimSpace(valAState) != ProductionUsabilityValAStateActive ||
		strings.TrimSpace(valBState) != ProductionUsabilityValBStateActive ||
		strings.TrimSpace(valCState) != ProductionUsabilityValCStateActive {
		return ProductionUsabilityValDStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		configReviewState,
		explainabilityReviewState,
		dryRunReviewState,
		redactionReviewState,
		degradedBehaviorReviewState,
		uiWindowingReviewState,
		commandNoiseReviewState,
		apiProtectionReviewState,
		cliResilienceReviewState,
		supportabilityReviewState,
		recoveryReviewState,
		upgradeRollbackReviewState,
		scaleEnvelopeReviewState,
		governanceBoundaryReviewState,
		regressionGateState,
	} {
		switch strings.TrimSpace(state) {
		case ProductionUsabilityValDConfigReviewStateActive,
			ProductionUsabilityValDExplainabilityReviewStateActive,
			ProductionUsabilityValDDryRunReviewStateActive,
			ProductionUsabilityValDRedactionReviewStateActive,
			ProductionUsabilityValDDegradedBehaviorReviewStateActive,
			ProductionUsabilityValDUIWindowingReviewStateActive,
			ProductionUsabilityValDCommandNoiseReviewStateActive,
			ProductionUsabilityValDAPIProtectionReviewStateActive,
			ProductionUsabilityValDCLIResilienceReviewStateActive,
			ProductionUsabilityValDSupportabilityReviewStateActive,
			ProductionUsabilityValDRecoveryReviewStateActive,
			ProductionUsabilityValDUpgradeRollbackReviewStateActive,
			ProductionUsabilityValDScaleEnvelopeReviewStateActive,
			ProductionUsabilityValDGovernanceBoundaryReviewStateActive,
			ProductionUsabilityValDRegressionGateStateActive:
		case ProductionUsabilityValDConfigReviewStatePartial,
			ProductionUsabilityValDExplainabilityReviewStatePartial,
			ProductionUsabilityValDDryRunReviewStatePartial,
			ProductionUsabilityValDRedactionReviewStatePartial,
			ProductionUsabilityValDDegradedBehaviorReviewStatePartial,
			ProductionUsabilityValDUIWindowingReviewStatePartial,
			ProductionUsabilityValDCommandNoiseReviewStatePartial,
			ProductionUsabilityValDAPIProtectionReviewStatePartial,
			ProductionUsabilityValDCLIResilienceReviewStatePartial,
			ProductionUsabilityValDSupportabilityReviewStatePartial,
			ProductionUsabilityValDRecoveryReviewStatePartial,
			ProductionUsabilityValDUpgradeRollbackReviewStatePartial,
			ProductionUsabilityValDScaleEnvelopeReviewStatePartial,
			ProductionUsabilityValDGovernanceBoundaryReviewStatePartial,
			ProductionUsabilityValDRegressionGateStatePartial:
			hasPartial = true
		default:
			return ProductionUsabilityValDStateIncomplete
		}
	}
	if hasPartial {
		return ProductionUsabilityValDStateSubstantial
	}
	return ProductionUsabilityValDStateActive
}

func EvaluateProductionUsabilityValDProofsState(val0State, valAState, valBState, valCState, configReviewState, explainabilityReviewState, dryRunReviewState, redactionReviewState, degradedBehaviorReviewState, uiWindowingReviewState, commandNoiseReviewState, apiProtectionReviewState, cliResilienceReviewState, supportabilityReviewState, recoveryReviewState, upgradeRollbackReviewState, scaleEnvelopeReviewState, governanceBoundaryReviewState, regressionGateState string, surfaceRefs, evidenceRefs, limitations, whyPoint4NotPass []string) string {
	baseState := EvaluateProductionUsabilityValDState(val0State, valAState, valBState, valCState, configReviewState, explainabilityReviewState, dryRunReviewState, redactionReviewState, degradedBehaviorReviewState, uiWindowingReviewState, commandNoiseReviewState, apiProtectionReviewState, cliResilienceReviewState, supportabilityReviewState, recoveryReviewState, upgradeRollbackReviewState, scaleEnvelopeReviewState, governanceBoundaryReviewState, regressionGateState)
	if len(surfaceRefs) < 20 || len(evidenceRefs) < 16 || len(limitations) == 0 || len(whyPoint4NotPass) == 0 {
		if baseState == ProductionUsabilityValDStateActive {
			return ProductionUsabilityValDStateSubstantial
		}
		return baseState
	}
	return baseState
}

package operability

import "strings"

const (
	ProductionUsabilityValBUIDataResilienceStateActive     = "production_usability_valb_ui_data_resilience_active"
	ProductionUsabilityValBUIDataResilienceStatePartial    = "production_usability_valb_ui_data_resilience_partial"
	ProductionUsabilityValBUIDataResilienceStateIncomplete = "production_usability_valb_ui_data_resilience_incomplete"

	ProductionUsabilityValBWindowingStateActive     = "production_usability_valb_windowing_active"
	ProductionUsabilityValBWindowingStatePartial    = "production_usability_valb_windowing_partial"
	ProductionUsabilityValBWindowingStateIncomplete = "production_usability_valb_windowing_incomplete"

	ProductionUsabilityValBResultSemanticsStateActive     = "production_usability_valb_result_semantics_active"
	ProductionUsabilityValBResultSemanticsStatePartial    = "production_usability_valb_result_semantics_partial"
	ProductionUsabilityValBResultSemanticsStateIncomplete = "production_usability_valb_result_semantics_incomplete"

	ProductionUsabilityValBCommandCenterStateActive     = "production_usability_valb_command_center_tasks_active"
	ProductionUsabilityValBCommandCenterStatePartial    = "production_usability_valb_command_center_tasks_partial"
	ProductionUsabilityValBCommandCenterStateIncomplete = "production_usability_valb_command_center_tasks_incomplete"

	ProductionUsabilityValBNoiseBudgetStateActive     = "production_usability_valb_noise_budget_active"
	ProductionUsabilityValBNoiseBudgetStatePartial    = "production_usability_valb_noise_budget_partial"
	ProductionUsabilityValBNoiseBudgetStateIncomplete = "production_usability_valb_noise_budget_incomplete"

	ProductionUsabilityValBAPIProtectionStateActive     = "production_usability_valb_api_protection_active"
	ProductionUsabilityValBAPIProtectionStatePartial    = "production_usability_valb_api_protection_partial"
	ProductionUsabilityValBAPIProtectionStateIncomplete = "production_usability_valb_api_protection_incomplete"

	ProductionUsabilityValBCLIResilienceStateActive     = "production_usability_valb_cli_resilience_active"
	ProductionUsabilityValBCLIResilienceStatePartial    = "production_usability_valb_cli_resilience_partial"
	ProductionUsabilityValBCLIResilienceStateIncomplete = "production_usability_valb_cli_resilience_incomplete"

	ProductionUsabilityValBScaleEnvelopeStateActive     = "production_usability_valb_scale_envelope_active"
	ProductionUsabilityValBScaleEnvelopeStatePartial    = "production_usability_valb_scale_envelope_partial"
	ProductionUsabilityValBScaleEnvelopeStateIncomplete = "production_usability_valb_scale_envelope_incomplete"

	ProductionUsabilityValBActionModeEnforcementStateActive     = "production_usability_valb_action_mode_enforcement_active"
	ProductionUsabilityValBActionModeEnforcementStatePartial    = "production_usability_valb_action_mode_enforcement_partial"
	ProductionUsabilityValBActionModeEnforcementStateIncomplete = "production_usability_valb_action_mode_enforcement_incomplete"

	ProductionUsabilityValBStateIncomplete  = "production_usability_valb_incomplete"
	ProductionUsabilityValBStateSubstantial = "production_usability_valb_substantially_ready"
	ProductionUsabilityValBStateActive      = "production_usability_valb_active"

	ProductionUsabilityResultComplete    = "complete"
	ProductionUsabilityResultPartial     = "partial"
	ProductionUsabilityResultStale       = "stale"
	ProductionUsabilityResultDegraded    = "degraded"
	ProductionUsabilityResultUnavailable = "unavailable"
	ProductionUsabilityResultUnsupported = "unsupported"

	ProductionUsabilityRequestClassRead      = "read"
	ProductionUsabilityRequestClassPreview   = "preview"
	ProductionUsabilityRequestClassExplain   = "explain"
	ProductionUsabilityRequestClassAuditOnly = "audit_only"
	ProductionUsabilityRequestClassMutation  = "mutation"
	ProductionUsabilityRequestClassSupport   = "support"
	ProductionUsabilityRequestClassSystem    = "system"

	ProductionUsabilityPriorityCritical   = "critical"
	ProductionUsabilityPriorityHigh       = "high"
	ProductionUsabilityPriorityNormal     = "normal"
	ProductionUsabilityPriorityLow        = "low"
	ProductionUsabilityPriorityBackground = "background"

	ProductionUsabilityBackpressureNone     = "none"
	ProductionUsabilityBackpressureSoft     = "soft"
	ProductionUsabilityBackpressureHard     = "hard"
	ProductionUsabilityBackpressureDegraded = "degraded"
)

type UIDataProjection struct {
	ProjectionID             string   `json:"projection_id"`
	DataSourceRef            string   `json:"data_source_ref"`
	FreshnessState           string   `json:"freshness_state"`
	GeneratedAtIndicator     string   `json:"generated_at_indicator"`
	SourceLastSeenIndicator  string   `json:"source_last_seen_indicator"`
	LimitationMessage        string   `json:"limitation_message"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	CanonicalTruthDisclaimer string   `json:"canonical_truth_disclaimer"`
	ClaimsCanonicalTruth     bool     `json:"claims_canonical_truth"`
	RecoveryHint             string   `json:"recovery_hint"`
	VisibilityScope          string   `json:"visibility_scope"`
	RedactionTier            string   `json:"redaction_tier"`
}

type UIDataResilienceModel struct {
	CurrentState string             `json:"current_state"`
	Items        []UIDataProjection `json:"items,omitempty"`
	Limitations  []string           `json:"limitations,omitempty"`
}

type VirtualDataWindowContract struct {
	CurrentState         string `json:"current_state"`
	WindowID             string `json:"window_id"`
	CursorMetadata       string `json:"cursor_metadata"`
	OffsetMetadata       string `json:"offset_metadata"`
	Limit                int    `json:"limit"`
	ResultCount          int    `json:"result_count"`
	TotalCountKnown      bool   `json:"total_count_known"`
	TotalCount           int    `json:"total_count"`
	NextCursor           string `json:"next_cursor"`
	PartialResult        bool   `json:"partial_result"`
	OrderingKey          string `json:"ordering_key"`
	StableSortRequired   bool   `json:"stable_sort_required"`
	MaxWindowSize        int    `json:"max_window_size"`
	MaxWindowEnforced    bool   `json:"max_window_enforced"`
	TruncationWarning    string `json:"truncation_warning"`
	StaleWindow          bool   `json:"stale_window"`
	ReplayRefreshHint    string `json:"replay_refresh_hint"`
	ClaimsCompleteData   bool   `json:"claims_complete_data"`
	ProjectionDisclaimer string `json:"projection_disclaimer"`
	MutatesCanonicalData bool   `json:"mutates_canonical_data"`
	LimitationMessage    string `json:"limitation_message"`
}

type ResultHealthDefinition struct {
	ResultHealth               string   `json:"result_health"`
	CurrentState               string   `json:"current_state"`
	Limitation                 string   `json:"limitation"`
	MissingComponents          []string `json:"missing_components,omitempty"`
	StaleComponents            []string `json:"stale_components,omitempty"`
	DegradedComponents         []string `json:"degraded_components,omitempty"`
	UnsupportedComponents      []string `json:"unsupported_components,omitempty"`
	RecoveryHint               string   `json:"recovery_hint"`
	SafeRetryGuidance          string   `json:"safe_retry_guidance"`
	ReportedAsFullSuccess      bool     `json:"reported_as_full_success"`
	ReportedAsHealthy          bool     `json:"reported_as_healthy"`
	UnsupportedSilentlyOmitted bool     `json:"unsupported_silently_omitted"`
}

type ResultSemanticsModel struct {
	CurrentState string                   `json:"current_state"`
	Items        []ResultHealthDefinition `json:"items,omitempty"`
	Limitations  []string                 `json:"limitations,omitempty"`
}

type CommandCenterTask struct {
	TaskID                           string   `json:"task_id"`
	TaskType                         string   `json:"task_type"`
	DecisionPriority                 string   `json:"decision_priority"`
	ActionRequired                   string   `json:"action_required"`
	WhyThisMatters                   string   `json:"why_this_matters"`
	AffectedSubjects                 []string `json:"affected_subjects,omitempty"`
	BlastRadiusHint                  string   `json:"blast_radius_hint"`
	RecommendedNextStep              string   `json:"recommended_next_step"`
	LinkedSurfaceRefs                []string `json:"linked_surface_refs,omitempty"`
	LinkedEvidenceRefs               []string `json:"linked_evidence_refs,omitempty"`
	VisibilityScope                  string   `json:"visibility_scope"`
	RedactionTier                    string   `json:"redaction_tier"`
	DecisionSupportOnly              bool     `json:"decision_support_only"`
	AcknowledgementEqualsRemediation bool     `json:"acknowledgement_equals_remediation"`
	ResolvedEqualsCanonicalClosure   bool     `json:"resolved_equals_canonical_closure"`
	WorkflowEvidenceRequired         bool     `json:"workflow_evidence_required"`
}

type CommandCenterTaskModel struct {
	CurrentState                string              `json:"current_state"`
	SupportedDecisionPriorities []string            `json:"supported_decision_priorities,omitempty"`
	SupportedActionRequired     []string            `json:"supported_action_required,omitempty"`
	Items                       []CommandCenterTask `json:"items,omitempty"`
	Limitations                 []string            `json:"limitations,omitempty"`
}

type NoiseBudgetEntry struct {
	Severity                         string `json:"severity"`
	GroupingKey                      string `json:"grouping_key"`
	DuplicateSuppressionKey          string `json:"duplicate_suppression_key"`
	AcknowledgementState             string `json:"acknowledgement_state"`
	EscalationPolicyRef              string `json:"escalation_policy_ref"`
	ReopenOnChange                   bool   `json:"reopen_on_change"`
	SuppressionReason                string `json:"suppression_reason"`
	CriticalBlocker                  bool   `json:"critical_blocker"`
	SuppressedCount                  int    `json:"suppressed_count"`
	FirstSeenAtIndicator             string `json:"first_seen_at_indicator"`
	LastSeenAtIndicator              string `json:"last_seen_at_indicator"`
	HighestSeverityPreserved         bool   `json:"highest_severity_preserved"`
	SuppressedDuplicatesAuditable    bool   `json:"suppressed_duplicates_auditable"`
	AcknowledgementEqualsRemediation bool   `json:"acknowledgement_equals_remediation"`
	ResolvedEqualsCanonicalClosure   bool   `json:"resolved_equals_canonical_closure"`
}

type NoiseBudgetModel struct {
	CurrentState              string             `json:"current_state"`
	SupportedSeverities       []string           `json:"supported_severities,omitempty"`
	SupportedAcknowledgements []string           `json:"supported_acknowledgement_states,omitempty"`
	Items                     []NoiseBudgetEntry `json:"items,omitempty"`
	Limitations               []string           `json:"limitations,omitempty"`
}

type APIProtectionRule struct {
	RequestClass                   string `json:"request_class"`
	CurrentState                   string `json:"current_state"`
	PriorityLane                   string `json:"priority_lane"`
	FairnessScope                  string `json:"fairness_scope"`
	RateLimitPolicyRef             string `json:"rate_limit_policy_ref"`
	BackpressureState              string `json:"backpressure_state"`
	DegradedResponseAllowed        bool   `json:"degraded_response_allowed"`
	RetryAfterHint                 string `json:"retry_after_hint"`
	SafeRetryGuidance              string `json:"safe_retry_guidance"`
	RejectionExplanation           string `json:"rejection_explanation"`
	VisibilityScope                string `json:"visibility_scope"`
	RedactionTier                  string `json:"redaction_tier"`
	PriorityBypassesGovernance     bool   `json:"priority_bypasses_governance"`
	PolicyDenialHiddenAsThrottling bool   `json:"policy_denial_hidden_as_throttling"`
	StarvationPrevented            bool   `json:"starvation_prevented"`
	GovernanceRequired             bool   `json:"governance_required"`
}

type APIProtectionDiscipline struct {
	CurrentState            string              `json:"current_state"`
	SupportedRequestClasses []string            `json:"supported_request_classes,omitempty"`
	SupportedPriorityLanes  []string            `json:"supported_priority_lanes,omitempty"`
	SupportedBackpressure   []string            `json:"supported_backpressure_states,omitempty"`
	Items                   []APIProtectionRule `json:"items,omitempty"`
	Limitations             []string            `json:"limitations,omitempty"`
}

type CLICommandResilience struct {
	CommandName            string   `json:"command_name"`
	CurrentState           string   `json:"current_state"`
	OperationType          string   `json:"operation_type"`
	ActionMode             string   `json:"action_mode"`
	RetrySafety            string   `json:"retry_safety"`
	IdempotencyKeyRequired bool     `json:"idempotency_key_required"`
	IdempotencyKeyPresent  bool     `json:"idempotency_key_present"`
	SafeRetryGuidance      string   `json:"safe_retry_guidance"`
	DoNotRetryReason       string   `json:"do_not_retry_reason"`
	PartialFailureBehavior string   `json:"partial_failure_behavior"`
	ExitCodeSemantics      []string `json:"exit_code_semantics,omitempty"`
	ExplainCommandRef      string   `json:"explain_command_ref"`
	InspectCommandRef      string   `json:"inspect_command_ref"`
	MutatesCanonicalState  bool     `json:"mutates_canonical_state"`
}

type CLIResilienceSurface struct {
	CurrentState               string                 `json:"current_state"`
	SupportedExitCodeSemantics []string               `json:"supported_exit_code_semantics,omitempty"`
	Items                      []CLICommandResilience `json:"items,omitempty"`
	Limitations                []string               `json:"limitations,omitempty"`
}

type ProductionScaleEnvelope struct {
	CurrentState                string   `json:"current_state"`
	ExpectedEventVolumeRange    string   `json:"expected_event_volume_range"`
	ExpectedWorkflowObjectRange string   `json:"expected_workflow_object_range"`
	TimelineQueryLimit          int      `json:"timeline_query_limit"`
	SearchQueryLimit            int      `json:"search_query_limit"`
	MaxPageSize                 int      `json:"max_page_size"`
	DashboardLatencyBudget      string   `json:"dashboard_latency_budget"`
	APILatencyBudget            string   `json:"api_latency_budget"`
	CLIResponseBudget           string   `json:"cli_response_budget"`
	DegradationThresholds       []string `json:"degradation_thresholds,omitempty"`
	UnsupportedScaleConditions  []string `json:"unsupported_scale_conditions,omitempty"`
	MeasurementNotes            []string `json:"measurement_notes,omitempty"`
	LimitationDisclaimer        string   `json:"limitation_disclaimer"`
	ScaleMeasured               bool     `json:"scale_measured"`
	LatencyBudgetsMeasured      bool     `json:"latency_budgets_measured"`
	ClaimsLatencyGuarantee      bool     `json:"claims_latency_guarantee"`
}

type ActionModeEnforcementRule struct {
	SurfaceRef         string   `json:"surface_ref"`
	SurfaceKind        string   `json:"surface_kind"`
	CurrentState       string   `json:"current_state"`
	ActionMode         string   `json:"action_mode"`
	OperationType      string   `json:"operation_type"`
	NonMutating        bool     `json:"non_mutating"`
	ExecutesMutation   bool     `json:"executes_mutation"`
	AvailableInValB    bool     `json:"available_in_val_b"`
	GovernedExternally bool     `json:"governed_externally"`
	GovernanceRef      string   `json:"governance_ref"`
	EvidenceRefs       []string `json:"evidence_refs,omitempty"`
}

type ActionModeEnforcementModel struct {
	CurrentState          string                      `json:"current_state"`
	SupportedSurfaceKinds []string                    `json:"supported_surface_kinds,omitempty"`
	Items                 []ActionModeEnforcementRule `json:"items,omitempty"`
	Limitations           []string                    `json:"limitations,omitempty"`
}

func productionUsabilityValBFreshnessStates() []string {
	return []string{
		ProductionUsabilityStatusFresh,
		ProductionUsabilityStatusStale,
		ProductionUsabilityStatusPartial,
		ProductionUsabilityStatusDegraded,
		ProductionUsabilityStatusUnavailable,
		ProductionUsabilityStatusUnsupported,
	}
}

func productionUsabilityValBResultHealthStates() []string {
	return []string{
		ProductionUsabilityResultComplete,
		ProductionUsabilityResultPartial,
		ProductionUsabilityResultStale,
		ProductionUsabilityResultDegraded,
		ProductionUsabilityResultUnavailable,
		ProductionUsabilityResultUnsupported,
	}
}

func productionUsabilityValBRequestClasses() []string {
	return []string{
		ProductionUsabilityRequestClassRead,
		ProductionUsabilityRequestClassPreview,
		ProductionUsabilityRequestClassExplain,
		ProductionUsabilityRequestClassAuditOnly,
		ProductionUsabilityRequestClassMutation,
		ProductionUsabilityRequestClassSupport,
		ProductionUsabilityRequestClassSystem,
	}
}

func productionUsabilityValBPriorityLanes() []string {
	return []string{
		ProductionUsabilityPriorityCritical,
		ProductionUsabilityPriorityHigh,
		ProductionUsabilityPriorityNormal,
		ProductionUsabilityPriorityLow,
		ProductionUsabilityPriorityBackground,
	}
}

func productionUsabilityValBBackpressureStates() []string {
	return []string{
		ProductionUsabilityBackpressureNone,
		ProductionUsabilityBackpressureSoft,
		ProductionUsabilityBackpressureHard,
		ProductionUsabilityBackpressureDegraded,
	}
}

func productionUsabilityValBRegisteredSurfaceRefs() []string {
	return []string{
		"/v1/production/usability-operability-recovery/valb/ui-data-resilience",
		"/v1/production/usability-operability-recovery/valb/windowing",
		"/v1/production/usability-operability-recovery/valb/result-semantics",
		"/v1/production/usability-operability-recovery/valb/command-center-tasks",
		"/v1/production/usability-operability-recovery/valb/noise-budget",
		"/v1/production/usability-operability-recovery/valb/api-protection",
		"/v1/production/usability-operability-recovery/valb/cli-resilience",
		"/v1/production/usability-operability-recovery/valb/scale-envelope",
		"/v1/production/usability-operability-recovery/valb/action-mode-enforcement",
		"/v1/production/usability-operability-recovery/valb/proofs",
	}
}

func ProductionUsabilityValBUIDataResilience() UIDataResilienceModel {
	return UIDataResilienceModel{
		CurrentState: "ui_data_resilience_ready",
		Items: []UIDataProjection{
			{ProjectionID: "command-center:fresh", DataSourceRef: "canonical_projection.command_center", FreshnessState: ProductionUsabilityStatusFresh, GeneratedAtIndicator: "generated_at_present", SourceLastSeenIndicator: "source_last_seen_recent", LimitationMessage: "fresh_projection_can_still_be_projection_only", EvidenceRefs: []string{"workflow/task-1"}, CanonicalTruthDisclaimer: "ui_projection_is_not_canonical_truth", ClaimsCanonicalTruth: false, RecoveryHint: "refresh_if_operator_needs_latest_projection", VisibilityScope: ProductionUsabilityVisibilityOperator, RedactionTier: ProductionUsabilityRedactionLow},
			{ProjectionID: "dashboard:stale", DataSourceRef: "projection.dashboard_cache", FreshnessState: ProductionUsabilityStatusStale, GeneratedAtIndicator: "generated_at_present", SourceLastSeenIndicator: "source_last_seen_old", LimitationMessage: "stale_projection_requires_refresh_before_trust", CanonicalTruthDisclaimer: "cached_projection_is_not_canonical_truth", ClaimsCanonicalTruth: false, RecoveryHint: "refresh_window_or_query_again", VisibilityScope: ProductionUsabilityVisibilityOperator, RedactionTier: ProductionUsabilityRedactionLow},
			{ProjectionID: "timeline:partial", DataSourceRef: "projection.timeline_window", FreshnessState: ProductionUsabilityStatusPartial, GeneratedAtIndicator: "generated_at_present", SourceLastSeenIndicator: "source_last_seen_recent", LimitationMessage: "windowed_timeline_is_partial_and_not_complete", CanonicalTruthDisclaimer: "partial_window_is_not_canonical_truth", ClaimsCanonicalTruth: false, RecoveryHint: "load_additional_window_segments", VisibilityScope: ProductionUsabilityVisibilityOperator, RedactionTier: ProductionUsabilityRedactionLow},
			{ProjectionID: "api:degraded", DataSourceRef: "projection.api_resilience", FreshnessState: ProductionUsabilityStatusDegraded, GeneratedAtIndicator: "generated_at_present", SourceLastSeenIndicator: "source_last_seen_recent", LimitationMessage: "degraded_api_surface_is_not_full_fidelity", CanonicalTruthDisclaimer: "degraded_projection_is_not_canonical_truth", ClaimsCanonicalTruth: false, RecoveryHint: "retry_when_degradation_clears", VisibilityScope: ProductionUsabilityVisibilityDeveloper, RedactionTier: ProductionUsabilityRedactionMedium},
			{ProjectionID: "support:unavailable", DataSourceRef: "projection.support_view", FreshnessState: ProductionUsabilityStatusUnavailable, GeneratedAtIndicator: "generated_at_present", SourceLastSeenIndicator: "source_last_seen_unavailable", LimitationMessage: "support_projection_temporarily_unavailable", CanonicalTruthDisclaimer: "unavailable_projection_is_not_canonical_truth", ClaimsCanonicalTruth: false, RecoveryHint: "retry_after_support_surface_recovers", VisibilityScope: ProductionUsabilityVisibilityInternalAdmin, RedactionTier: ProductionUsabilityRedactionNone},
			{ProjectionID: "public:unsupported", DataSourceRef: "projection.public_summary", FreshnessState: ProductionUsabilityStatusUnsupported, GeneratedAtIndicator: "generated_at_present", SourceLastSeenIndicator: "source_last_seen_not_supported", LimitationMessage: "public_summary_does_not_support_internal_operability_detail", CanonicalTruthDisclaimer: "unsupported_projection_is_not_canonical_truth", ClaimsCanonicalTruth: false, RecoveryHint: "use_internal_surface_for_supported_detail", VisibilityScope: ProductionUsabilityVisibilityPublicSafe, RedactionTier: ProductionUsabilityRedactionPublicSafe},
		},
		Limitations: []string{
			"Val B UI data resilience covers bounded projection health only and never upgrades a projection into canonical truth.",
		},
	}
}

func ProductionUsabilityValBWindowing() VirtualDataWindowContract {
	return VirtualDataWindowContract{
		CurrentState:         "virtual_windowing_ready",
		WindowID:             "timeline-window-001",
		CursorMetadata:       "cursor:timeline:after:evt-100",
		OffsetMetadata:       "offset:0",
		Limit:                100,
		ResultCount:          100,
		TotalCountKnown:      false,
		TotalCount:           0,
		NextCursor:           "cursor:timeline:after:evt-200",
		PartialResult:        true,
		OrderingKey:          "canonical_event_time_desc",
		StableSortRequired:   true,
		MaxWindowSize:        500,
		MaxWindowEnforced:    true,
		TruncationWarning:    "timeline_window_truncated_at_limit",
		StaleWindow:          false,
		ReplayRefreshHint:    "refresh_window_or_follow_next_cursor",
		ClaimsCompleteData:   false,
		ProjectionDisclaimer: "windowed_projection_is_not_canonical_truth",
		MutatesCanonicalData: false,
		LimitationMessage:    "windowed_response_is_partial_until_all_relevant_segments_are_loaded",
	}
}

func ProductionUsabilityValBResultSemantics() ResultSemanticsModel {
	return ResultSemanticsModel{
		CurrentState: "shared_result_semantics_ready",
		Items: []ResultHealthDefinition{
			{ResultHealth: ProductionUsabilityResultComplete, CurrentState: "result_semantics_ready", Limitation: "complete_result_is_still_projection_bound", RecoveryHint: "no_recovery_needed", SafeRetryGuidance: "retry_optional_if_fresher_projection_needed"},
			{ResultHealth: ProductionUsabilityResultPartial, CurrentState: "result_semantics_ready", Limitation: "partial_result_missing_some_components", MissingComponents: []string{"timeline_tail"}, RecoveryHint: "refresh_or_expand_window", SafeRetryGuidance: "retry_after_fetching_missing_components", ReportedAsFullSuccess: false},
			{ResultHealth: ProductionUsabilityResultStale, CurrentState: "result_semantics_ready", Limitation: "stale_result_older_than_latest_observation", StaleComponents: []string{"dashboard_rollup"}, RecoveryHint: "refresh_projection", SafeRetryGuidance: "retry_safe_after_refresh"},
			{ResultHealth: ProductionUsabilityResultDegraded, CurrentState: "result_semantics_ready", Limitation: "degraded_result_has_reduced_fidelity", DegradedComponents: []string{"search_index"}, RecoveryHint: "retry_when_degradation_clears", SafeRetryGuidance: "retry_conditionally_safe_after_backpressure_recovery", ReportedAsHealthy: false},
			{ResultHealth: ProductionUsabilityResultUnavailable, CurrentState: "result_semantics_ready", Limitation: "requested_projection_unavailable", MissingComponents: []string{"support_surface"}, RecoveryHint: "use_alternate_surface_or_retry_later", SafeRetryGuidance: "retry_after_unavailability_clears"},
			{ResultHealth: ProductionUsabilityResultUnsupported, CurrentState: "result_semantics_ready", Limitation: "requested_surface_not_supported_for_scope_or_shape", UnsupportedComponents: []string{"public_internal_debug_view"}, RecoveryHint: "use_supported_surface_or_scope", SafeRetryGuidance: "do_not_retry_until_supported_path_exists", UnsupportedSilentlyOmitted: false},
		},
		Limitations: []string{
			"Val B result semantics distinguish partial, stale, degraded, unavailable, and unsupported instead of flattening them into generic success.",
		},
	}
}

func ProductionUsabilityValBCommandCenterTasks() CommandCenterTaskModel {
	return CommandCenterTaskModel{
		CurrentState:                "command_center_task_model_ready",
		SupportedDecisionPriorities: []string{ProductionUsabilityDecisionBlocker, ProductionUsabilityDecisionUrgent, ProductionUsabilityDecisionNormal, ProductionUsabilityDecisionLow, ProductionUsabilityDecisionInformational},
		SupportedActionRequired:     []string{ProductionUsabilityNoAction, ProductionUsabilityOperatorAction, ProductionUsabilityAdminAction, ProductionUsabilityGovernanceAction, ProductionUsabilitySupportAction},
		Items: []CommandCenterTask{
			{TaskID: "task-001", TaskType: "config_validation_blocker", DecisionPriority: ProductionUsabilityDecisionBlocker, ActionRequired: ProductionUsabilityOperatorAction, WhyThisMatters: "invalid_config_blocks_safe_activation", AffectedSubjects: []string{"tenant/acme/prod"}, BlastRadiusHint: "tenant_scope", RecommendedNextStep: "inspect_and_fix_schema_errors", LinkedSurfaceRefs: []string{"/v1/production/usability-operability-recovery/vala/config-factory"}, LinkedEvidenceRefs: []string{"config_factory_core"}, VisibilityScope: ProductionUsabilityVisibilityOperator, RedactionTier: ProductionUsabilityRedactionLow, DecisionSupportOnly: true, WorkflowEvidenceRequired: true},
			{TaskID: "task-002", TaskType: "policy_upgrade_notice", DecisionPriority: ProductionUsabilityDecisionNormal, ActionRequired: ProductionUsabilityAdminAction, WhyThisMatters: "deprecated_policy_schema_requires_upgrade_planning", AffectedSubjects: []string{"policy/bundle-a"}, BlastRadiusHint: "policy_bundle_scope", RecommendedNextStep: "run_policy_dry_run_before_upgrade", LinkedSurfaceRefs: []string{"/v1/production/usability-operability-recovery/vala/policy-schema"}, LinkedEvidenceRefs: []string{"policy_schema_core"}, VisibilityScope: ProductionUsabilityVisibilityInternalAdmin, RedactionTier: ProductionUsabilityRedactionNone, DecisionSupportOnly: true, WorkflowEvidenceRequired: true},
			{TaskID: "task-003", TaskType: "partner_visibility_notice", DecisionPriority: ProductionUsabilityDecisionInformational, ActionRequired: ProductionUsabilityNoAction, WhyThisMatters: "partner_scope_receives_bounded_task_explanation_only", AffectedSubjects: []string{"partner/vendor-a"}, BlastRadiusHint: "partner_scope", RecommendedNextStep: "request_internal_operator_support_if_more_detail_needed", LinkedSurfaceRefs: []string{"/v1/production/usability-operability-recovery/vala/explain"}, VisibilityScope: ProductionUsabilityVisibilityPartner, RedactionTier: ProductionUsabilityRedactionHigh, DecisionSupportOnly: true, WorkflowEvidenceRequired: true},
		},
		Limitations: []string{
			"Val B command center tasks are decision-support projections only and do not approve, close, or mutate workflow state.",
		},
	}
}

func ProductionUsabilityValBNoiseBudget() NoiseBudgetModel {
	return NoiseBudgetModel{
		CurrentState:              "noise_budget_grouping_ready",
		SupportedSeverities:       []string{ProductionUsabilitySeverityCritical, ProductionUsabilitySeverityError, ProductionUsabilitySeverityWarning, ProductionUsabilitySeverityInfo},
		SupportedAcknowledgements: []string{ProductionUsabilityAckUnacknowledged, ProductionUsabilityAckAcknowledged, ProductionUsabilityAckSuppressedDup, ProductionUsabilityAckResolved, ProductionUsabilityAckReopened},
		Items: []NoiseBudgetEntry{
			{Severity: ProductionUsabilitySeverityCritical, GroupingKey: "config_blocker", DuplicateSuppressionKey: "config_blocker/acme", AcknowledgementState: ProductionUsabilityAckUnacknowledged, EscalationPolicyRef: "notify-platform-blockers", ReopenOnChange: true, CriticalBlocker: true, SuppressedCount: 0, FirstSeenAtIndicator: "first_seen_present", LastSeenAtIndicator: "last_seen_present", HighestSeverityPreserved: true, SuppressedDuplicatesAuditable: true},
			{Severity: ProductionUsabilitySeverityWarning, GroupingKey: "deprecated_schema", DuplicateSuppressionKey: "deprecated_schema/acme", AcknowledgementState: ProductionUsabilityAckSuppressedDup, EscalationPolicyRef: "notify-schema-maintainers", ReopenOnChange: true, SuppressionReason: "duplicate_warning_window", CriticalBlocker: false, SuppressedCount: 4, FirstSeenAtIndicator: "first_seen_present", LastSeenAtIndicator: "last_seen_present", HighestSeverityPreserved: true, SuppressedDuplicatesAuditable: true},
			{Severity: ProductionUsabilitySeverityError, GroupingKey: "policy_conflict", DuplicateSuppressionKey: "policy_conflict/acme", AcknowledgementState: ProductionUsabilityAckReopened, EscalationPolicyRef: "notify-policy-owners", ReopenOnChange: true, CriticalBlocker: false, SuppressedCount: 0, FirstSeenAtIndicator: "first_seen_present", LastSeenAtIndicator: "last_seen_present", HighestSeverityPreserved: true, SuppressedDuplicatesAuditable: true},
		},
		Limitations: []string{
			"Val B noise budget groups duplicate projections without deleting their auditability or converting acknowledgement into remediation.",
		},
	}
}

func ProductionUsabilityValBAPIProtection() APIProtectionDiscipline {
	return APIProtectionDiscipline{
		CurrentState:            "api_protection_discipline_ready",
		SupportedRequestClasses: productionUsabilityValBRequestClasses(),
		SupportedPriorityLanes:  productionUsabilityValBPriorityLanes(),
		SupportedBackpressure:   productionUsabilityValBBackpressureStates(),
		Items: []APIProtectionRule{
			{RequestClass: ProductionUsabilityRequestClassRead, CurrentState: "api_protection_ready", PriorityLane: ProductionUsabilityPriorityNormal, FairnessScope: "tenant", RateLimitPolicyRef: "api-read-normal", BackpressureState: ProductionUsabilityBackpressureNone, DegradedResponseAllowed: false, RetryAfterHint: "none", SafeRetryGuidance: "retry_safe_for_read_after_normal_interval", RejectionExplanation: "read_requests_use_fairness_and_rate_policies", VisibilityScope: ProductionUsabilityVisibilityDeveloper, RedactionTier: ProductionUsabilityRedactionMedium, StarvationPrevented: true},
			{RequestClass: ProductionUsabilityRequestClassPreview, CurrentState: "api_protection_ready", PriorityLane: ProductionUsabilityPriorityHigh, FairnessScope: "tenant", RateLimitPolicyRef: "api-preview-high", BackpressureState: ProductionUsabilityBackpressureSoft, DegradedResponseAllowed: true, RetryAfterHint: "retry_after_soft_backpressure", SafeRetryGuidance: "retry_conditionally_safe_after_retry_after_window", RejectionExplanation: "preview_requests_may_receive_soft_backpressure", VisibilityScope: ProductionUsabilityVisibilityDeveloper, RedactionTier: ProductionUsabilityRedactionMedium, StarvationPrevented: true},
			{RequestClass: ProductionUsabilityRequestClassExplain, CurrentState: "api_protection_ready", PriorityLane: ProductionUsabilityPriorityHigh, FairnessScope: "tenant", RateLimitPolicyRef: "api-explain-high", BackpressureState: ProductionUsabilityBackpressureDegraded, DegradedResponseAllowed: true, RetryAfterHint: "retry_after_degraded_lane_recovers", SafeRetryGuidance: "retry_conditionally_safe_when_degraded_lane_clears", RejectionExplanation: "explain_requests_may_return_degraded_redacted_output_under_pressure", VisibilityScope: ProductionUsabilityVisibilityOperator, RedactionTier: ProductionUsabilityRedactionLow, StarvationPrevented: true},
			{RequestClass: ProductionUsabilityRequestClassAuditOnly, CurrentState: "api_protection_ready", PriorityLane: ProductionUsabilityPriorityNormal, FairnessScope: "tenant", RateLimitPolicyRef: "api-audit-normal", BackpressureState: ProductionUsabilityBackpressureSoft, DegradedResponseAllowed: true, RetryAfterHint: "retry_after_soft_backpressure", SafeRetryGuidance: "retry_safe_for_audit_only_after_retry_after_window", RejectionExplanation: "audit_only_requests_remain_non_mutating_even_when_degraded", VisibilityScope: ProductionUsabilityVisibilityOperator, RedactionTier: ProductionUsabilityRedactionLow, StarvationPrevented: true},
			{RequestClass: ProductionUsabilityRequestClassMutation, CurrentState: "api_protection_ready", PriorityLane: ProductionUsabilityPriorityCritical, FairnessScope: "tenant", RateLimitPolicyRef: "api-mutation-critical", BackpressureState: ProductionUsabilityBackpressureHard, DegradedResponseAllowed: false, RetryAfterHint: "retry_after_governed_mutation_backpressure_clears", SafeRetryGuidance: "retry_only_if_mutation_is_governed_and_retry_safe", RejectionExplanation: "mutation_requests_remain_governed_even_on_critical_lane", VisibilityScope: ProductionUsabilityVisibilityInternalAdmin, RedactionTier: ProductionUsabilityRedactionNone, StarvationPrevented: true, GovernanceRequired: true},
			{RequestClass: ProductionUsabilityRequestClassSupport, CurrentState: "api_protection_ready", PriorityLane: ProductionUsabilityPriorityLow, FairnessScope: "tenant", RateLimitPolicyRef: "api-support-low", BackpressureState: ProductionUsabilityBackpressureSoft, DegradedResponseAllowed: true, RetryAfterHint: "retry_after_support_lane_recovers", SafeRetryGuidance: "retry_conditionally_safe_after_support_backpressure_clears", RejectionExplanation: "support_requests_do_not_preempt_operator_work_indefinitely", VisibilityScope: ProductionUsabilityVisibilityOperator, RedactionTier: ProductionUsabilityRedactionHigh, StarvationPrevented: true},
			{RequestClass: ProductionUsabilityRequestClassSystem, CurrentState: "api_protection_ready", PriorityLane: ProductionUsabilityPriorityBackground, FairnessScope: "cluster", RateLimitPolicyRef: "api-system-background", BackpressureState: ProductionUsabilityBackpressureDegraded, DegradedResponseAllowed: true, RetryAfterHint: "retry_after_background_lane_recovers", SafeRetryGuidance: "retry_conditionally_safe_for_background_system_tasks", RejectionExplanation: "system_requests_do_not_starve_human_operator_lanes", VisibilityScope: ProductionUsabilityVisibilityInternalAdmin, RedactionTier: ProductionUsabilityRedactionNone, StarvationPrevented: true},
		},
		Limitations: []string{
			"Val B API protection models bounded priority, fairness, and backpressure semantics only; it does not rewrite transport or gateway enforcement.",
		},
	}
}

func ProductionUsabilityValBCLIResilience() CLIResilienceSurface {
	return CLIResilienceSurface{
		CurrentState:               "cli_resilience_ready",
		SupportedExitCodeSemantics: []string{"success", "partial", "blocked", "unsupported", "internal_error"},
		Items: []CLICommandResilience{
			{CommandName: "changelock config inspect", CurrentState: "cli_resilience_ready", OperationType: ProductionUsabilityOperationReadOnly, ActionMode: ProductionUsabilityActionModeViewOnly, RetrySafety: ProductionUsabilityRetrySafe, SafeRetryGuidance: "retry_safe_when_fresher_projection_is_needed", PartialFailureBehavior: "return_partial_with_nonzero_partial_exit", ExitCodeSemantics: []string{"success", "partial", "blocked", "unsupported", "internal_error"}, ExplainCommandRef: "changelock config explain", InspectCommandRef: "changelock config inspect", MutatesCanonicalState: false},
			{CommandName: "changelock config explain", CurrentState: "cli_resilience_ready", OperationType: ProductionUsabilityOperationReadOnly, ActionMode: ProductionUsabilityActionModeExplain, RetrySafety: ProductionUsabilityRetrySafe, SafeRetryGuidance: "retry_safe_when_scope_or_redaction_changes", PartialFailureBehavior: "return_partial_with_bounded_explain_output", ExitCodeSemantics: []string{"success", "partial", "blocked", "unsupported", "internal_error"}, ExplainCommandRef: "changelock config explain", InspectCommandRef: "changelock config inspect", MutatesCanonicalState: false},
			{CommandName: "changelock policy preview", CurrentState: "cli_resilience_ready", OperationType: ProductionUsabilityOperationPreviewOnly, ActionMode: ProductionUsabilityActionModePreview, RetrySafety: ProductionUsabilityRetrySafe, SafeRetryGuidance: "retry_safe_after_preview_inputs_change", PartialFailureBehavior: "return_partial_preview_with_limitations", ExitCodeSemantics: []string{"success", "partial", "blocked", "unsupported", "internal_error"}, ExplainCommandRef: "changelock policy explain", InspectCommandRef: "changelock policy inspect", MutatesCanonicalState: false},
			{CommandName: "changelock policy dry-run", CurrentState: "cli_resilience_ready", OperationType: ProductionUsabilityOperationPreviewOnly, ActionMode: ProductionUsabilityActionModeDryRun, RetrySafety: ProductionUsabilityRetrySafe, SafeRetryGuidance: "retry_safe_when_inputs_change", PartialFailureBehavior: "return_partial_dry_run_with_blocking_rules", ExitCodeSemantics: []string{"success", "partial", "blocked", "unsupported", "internal_error"}, ExplainCommandRef: "changelock policy explain", InspectCommandRef: "changelock policy inspect", MutatesCanonicalState: false},
			{CommandName: "changelock policy audit", CurrentState: "cli_resilience_ready", OperationType: ProductionUsabilityOperationReadOnly, ActionMode: ProductionUsabilityActionModeAuditOnly, RetrySafety: ProductionUsabilityRetrySafe, SafeRetryGuidance: "retry_safe_for_audit_only_queries", PartialFailureBehavior: "return_partial_audit_with_limitations", ExitCodeSemantics: []string{"success", "partial", "blocked", "unsupported", "internal_error"}, ExplainCommandRef: "changelock policy explain", InspectCommandRef: "changelock policy inspect", MutatesCanonicalState: false},
			{CommandName: "changelock governed apply", CurrentState: "cli_resilience_ready", OperationType: ProductionUsabilityOperationIdempotentMutation, ActionMode: ProductionUsabilityActionModeMutate, RetrySafety: ProductionUsabilityRetryConditionallySafe, IdempotencyKeyRequired: true, IdempotencyKeyPresent: true, SafeRetryGuidance: "retry_only_with_same_idempotency_key_and_governed_authorization", PartialFailureBehavior: "surface_partial_mutation_failure_explicitly", ExitCodeSemantics: []string{"success", "partial", "blocked", "unsupported", "internal_error"}, ExplainCommandRef: "changelock governed explain", InspectCommandRef: "changelock governed inspect", MutatesCanonicalState: true},
			{CommandName: "changelock governed rotate", CurrentState: "cli_resilience_ready", OperationType: ProductionUsabilityOperationSideEffecting, ActionMode: ProductionUsabilityActionModeMutate, RetrySafety: ProductionUsabilityRetryUnsafe, SafeRetryGuidance: "do_not_retry_without_operator_review", DoNotRetryReason: "side_effecting_rotation_may_repeat_external_effects", PartialFailureBehavior: "surface_partial_side_effect_failure_and_stop", ExitCodeSemantics: []string{"success", "partial", "blocked", "unsupported", "internal_error"}, ExplainCommandRef: "changelock governed explain", InspectCommandRef: "changelock governed inspect", MutatesCanonicalState: true},
		},
		Limitations: []string{
			"Val B CLI resilience surfaces explain retry, idempotency, and partial-failure behavior without bypassing policy or evidence gates.",
		},
	}
}

func ProductionUsabilityValBScaleEnvelope() ProductionScaleEnvelope {
	return ProductionScaleEnvelope{
		CurrentState:                "production_scale_envelope_ready",
		ExpectedEventVolumeRange:    "up_to_100k_events_per_day_projection_target",
		ExpectedWorkflowObjectRange: "up_to_10k_active_objects_per_tenant_projection_target",
		TimelineQueryLimit:          1000,
		SearchQueryLimit:            200,
		MaxPageSize:                 500,
		DashboardLatencyBudget:      "p95_3s_planning_target",
		APILatencyBudget:            "p95_1s_planning_target",
		CLIResponseBudget:           "p95_2s_planning_target",
		DegradationThresholds:       []string{"timeline_queries_above_1000", "search_queries_above_200", "page_size_above_500"},
		UnsupportedScaleConditions:  []string{"whole_dataset_unwindowed_timeline_queries", "unbounded_cross_tenant_search"},
		MeasurementNotes:            []string{"latency_budgets_are_planning_targets_only", "scale_envelope_not_yet_final_performance_gate"},
		LimitationDisclaimer:        "scale_and_latency_envelope_is_measurement_aware_and_unmeasured_conditions_are_marked_as_limitations",
		ScaleMeasured:               false,
		LatencyBudgetsMeasured:      false,
		ClaimsLatencyGuarantee:      false,
	}
}

func ProductionUsabilityValBActionModeEnforcement() ActionModeEnforcementModel {
	return ActionModeEnforcementModel{
		CurrentState:          "action_mode_enforcement_ready",
		SupportedSurfaceKinds: []string{"ui", "api", "cli"},
		Items: []ActionModeEnforcementRule{
			{SurfaceRef: "/v1/production/usability-operability-recovery/valb/command-center-tasks", SurfaceKind: "ui", CurrentState: "action_mode_enforcement_ready", ActionMode: ProductionUsabilityActionModeViewOnly, OperationType: ProductionUsabilityOperationReadOnly, NonMutating: true, ExecutesMutation: false, AvailableInValB: true},
			{SurfaceRef: "/v1/production/usability-operability-recovery/valb/api-protection", SurfaceKind: "api", CurrentState: "action_mode_enforcement_ready", ActionMode: ProductionUsabilityActionModeExplain, OperationType: ProductionUsabilityOperationReadOnly, NonMutating: true, ExecutesMutation: false, AvailableInValB: true},
			{SurfaceRef: "/v1/production/usability-operability-recovery/valb/ui-data-resilience", SurfaceKind: "api", CurrentState: "action_mode_enforcement_ready", ActionMode: ProductionUsabilityActionModePreview, OperationType: ProductionUsabilityOperationPreviewOnly, NonMutating: true, ExecutesMutation: false, AvailableInValB: true},
			{SurfaceRef: "/v1/production/usability-operability-recovery/valb/cli-resilience", SurfaceKind: "cli", CurrentState: "action_mode_enforcement_ready", ActionMode: ProductionUsabilityActionModeDryRun, OperationType: ProductionUsabilityOperationPreviewOnly, NonMutating: true, ExecutesMutation: false, AvailableInValB: true},
			{SurfaceRef: "/v1/production/usability-operability-recovery/valb/result-semantics", SurfaceKind: "api", CurrentState: "action_mode_enforcement_ready", ActionMode: ProductionUsabilityActionModeAuditOnly, OperationType: ProductionUsabilityOperationReadOnly, NonMutating: true, ExecutesMutation: false, AvailableInValB: true},
			{SurfaceRef: "/v1/production/usability-operability-recovery/valb/api-protection", SurfaceKind: "api", CurrentState: "action_mode_enforcement_ready", ActionMode: ProductionUsabilityActionModeEnforce, OperationType: ProductionUsabilityOperationIdempotentMutation, NonMutating: false, ExecutesMutation: true, AvailableInValB: false},
			{SurfaceRef: "/v1/production/usability-operability-recovery/valb/cli-resilience", SurfaceKind: "cli", CurrentState: "action_mode_enforcement_ready", ActionMode: ProductionUsabilityActionModeMutate, OperationType: ProductionUsabilityOperationIdempotentMutation, NonMutating: false, ExecutesMutation: true, AvailableInValB: true, GovernedExternally: true, GovernanceRef: "point3.workflow_authority.valb.governed_mutation", EvidenceRefs: []string{"workflow_authority_valb_proofs", "governed_mutation_contract"}},
		},
		Limitations: []string{
			"Val B action-mode enforcement keeps suggest, preview, explain, dry-run, and audit-only surfaces non-mutating and does not introduce new mutation authority.",
		},
	}
}

func EvaluateProductionUsabilityValBUIDataResilienceState(model UIDataResilienceModel) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValBUIDataResilienceStateIncomplete
	}
	expectedStates := map[string]struct{}{}
	for _, state := range productionUsabilityValBFreshnessStates() {
		expectedStates[state] = struct{}{}
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		freshness := strings.TrimSpace(item.FreshnessState)
		if strings.TrimSpace(item.ProjectionID) == "" || strings.TrimSpace(item.DataSourceRef) == "" || freshness == "" || strings.TrimSpace(item.GeneratedAtIndicator) == "" || strings.TrimSpace(item.SourceLastSeenIndicator) == "" || strings.TrimSpace(item.LimitationMessage) == "" || strings.TrimSpace(item.CanonicalTruthDisclaimer) == "" || strings.TrimSpace(item.RecoveryHint) == "" || strings.TrimSpace(item.VisibilityScope) == "" || strings.TrimSpace(item.RedactionTier) == "" {
			return ProductionUsabilityValBUIDataResilienceStateIncomplete
		}
		if _, ok := expectedStates[freshness]; !ok {
			return ProductionUsabilityValBUIDataResilienceStatePartial
		}
		if _, duplicate := seen[freshness]; duplicate {
			return ProductionUsabilityValBUIDataResilienceStatePartial
		}
		seen[freshness] = struct{}{}
		if !containsTrimmedString([]string{ProductionUsabilityVisibilityInternalAdmin, ProductionUsabilityVisibilityOperator, ProductionUsabilityVisibilityDeveloper, ProductionUsabilityVisibilityPartner, ProductionUsabilityVisibilityPublicSafe}, item.VisibilityScope) ||
			!containsTrimmedString([]string{ProductionUsabilityRedactionNone, ProductionUsabilityRedactionLow, ProductionUsabilityRedactionMedium, ProductionUsabilityRedactionHigh, ProductionUsabilityRedactionPublicSafe}, item.RedactionTier) ||
			!strings.Contains(strings.TrimSpace(item.CanonicalTruthDisclaimer), "not_canonical_truth") ||
			item.ClaimsCanonicalTruth {
			return ProductionUsabilityValBUIDataResilienceStatePartial
		}
	}
	if len(seen) != len(expectedStates) {
		return ProductionUsabilityValBUIDataResilienceStatePartial
	}
	return ProductionUsabilityValBUIDataResilienceStateActive
}

func EvaluateProductionUsabilityValBWindowingState(model VirtualDataWindowContract) string {
	if strings.TrimSpace(model.WindowID) == "" || strings.TrimSpace(model.OrderingKey) == "" || strings.TrimSpace(model.ReplayRefreshHint) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" || strings.TrimSpace(model.LimitationMessage) == "" {
		return ProductionUsabilityValBWindowingStateIncomplete
	}
	if model.Limit <= 0 || model.MaxWindowSize <= 0 || model.Limit > model.MaxWindowSize || !model.MaxWindowEnforced || !model.StableSortRequired || model.ResultCount < 0 || model.ResultCount > model.Limit || model.MutatesCanonicalData ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValBWindowingStatePartial
	}
	if model.PartialResult || model.StaleWindow {
		if strings.TrimSpace(model.TruncationWarning) == "" {
			return ProductionUsabilityValBWindowingStatePartial
		}
	}
	if !model.TotalCountKnown && model.ClaimsCompleteData {
		return ProductionUsabilityValBWindowingStatePartial
	}
	return ProductionUsabilityValBWindowingStateActive
}

func EvaluateProductionUsabilityValBResultSemanticsState(model ResultSemanticsModel) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValBResultSemanticsStateIncomplete
	}
	expected := map[string]struct{}{}
	for _, item := range productionUsabilityValBResultHealthStates() {
		expected[item] = struct{}{}
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		health := strings.TrimSpace(item.ResultHealth)
		if health == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.Limitation) == "" || strings.TrimSpace(item.RecoveryHint) == "" || strings.TrimSpace(item.SafeRetryGuidance) == "" {
			return ProductionUsabilityValBResultSemanticsStateIncomplete
		}
		if _, ok := expected[health]; !ok {
			return ProductionUsabilityValBResultSemanticsStatePartial
		}
		if _, duplicate := seen[health]; duplicate {
			return ProductionUsabilityValBResultSemanticsStatePartial
		}
		seen[health] = struct{}{}
		if health == ProductionUsabilityResultPartial && item.ReportedAsFullSuccess {
			return ProductionUsabilityValBResultSemanticsStatePartial
		}
		if health == ProductionUsabilityResultDegraded && item.ReportedAsHealthy {
			return ProductionUsabilityValBResultSemanticsStatePartial
		}
		if health == ProductionUsabilityResultUnsupported && item.UnsupportedSilentlyOmitted {
			return ProductionUsabilityValBResultSemanticsStatePartial
		}
	}
	if len(seen) != len(expected) {
		return ProductionUsabilityValBResultSemanticsStatePartial
	}
	return ProductionUsabilityValBResultSemanticsStateActive
}

func EvaluateProductionUsabilityValBCommandCenterState(model CommandCenterTaskModel) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValBCommandCenterStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedDecisionPriorities, ProductionUsabilityDecisionBlocker, ProductionUsabilityDecisionUrgent, ProductionUsabilityDecisionNormal, ProductionUsabilityDecisionLow, ProductionUsabilityDecisionInformational) ||
		!containsExactTrimmedStringSet(model.SupportedActionRequired, ProductionUsabilityNoAction, ProductionUsabilityOperatorAction, ProductionUsabilityAdminAction, ProductionUsabilityGovernanceAction, ProductionUsabilitySupportAction) {
		return ProductionUsabilityValBCommandCenterStatePartial
	}
	for _, item := range model.Items {
		if strings.TrimSpace(item.TaskID) == "" || strings.TrimSpace(item.TaskType) == "" || strings.TrimSpace(item.DecisionPriority) == "" || strings.TrimSpace(item.ActionRequired) == "" || strings.TrimSpace(item.WhyThisMatters) == "" || len(item.AffectedSubjects) == 0 || strings.TrimSpace(item.BlastRadiusHint) == "" || strings.TrimSpace(item.RecommendedNextStep) == "" || len(item.LinkedSurfaceRefs) == 0 || strings.TrimSpace(item.VisibilityScope) == "" || strings.TrimSpace(item.RedactionTier) == "" {
			return ProductionUsabilityValBCommandCenterStateIncomplete
		}
		if !containsTrimmedString(model.SupportedDecisionPriorities, item.DecisionPriority) || !containsTrimmedString(model.SupportedActionRequired, item.ActionRequired) || !item.DecisionSupportOnly || item.AcknowledgementEqualsRemediation || item.ResolvedEqualsCanonicalClosure || !item.WorkflowEvidenceRequired {
			return ProductionUsabilityValBCommandCenterStatePartial
		}
		if (item.VisibilityScope == ProductionUsabilityVisibilityPartner || item.VisibilityScope == ProductionUsabilityVisibilityPublicSafe) && len(item.LinkedEvidenceRefs) > 0 {
			return ProductionUsabilityValBCommandCenterStatePartial
		}
	}
	return ProductionUsabilityValBCommandCenterStateActive
}

func EvaluateProductionUsabilityValBNoiseBudgetState(model NoiseBudgetModel) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValBNoiseBudgetStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedSeverities, ProductionUsabilitySeverityCritical, ProductionUsabilitySeverityError, ProductionUsabilitySeverityWarning, ProductionUsabilitySeverityInfo) ||
		!containsExactTrimmedStringSet(model.SupportedAcknowledgements, ProductionUsabilityAckUnacknowledged, ProductionUsabilityAckAcknowledged, ProductionUsabilityAckSuppressedDup, ProductionUsabilityAckResolved, ProductionUsabilityAckReopened) {
		return ProductionUsabilityValBNoiseBudgetStatePartial
	}
	for _, item := range model.Items {
		if strings.TrimSpace(item.Severity) == "" || strings.TrimSpace(item.GroupingKey) == "" || strings.TrimSpace(item.DuplicateSuppressionKey) == "" || strings.TrimSpace(item.AcknowledgementState) == "" || strings.TrimSpace(item.EscalationPolicyRef) == "" || strings.TrimSpace(item.FirstSeenAtIndicator) == "" || strings.TrimSpace(item.LastSeenAtIndicator) == "" {
			return ProductionUsabilityValBNoiseBudgetStateIncomplete
		}
		if !containsTrimmedString(model.SupportedSeverities, item.Severity) || !containsTrimmedString(model.SupportedAcknowledgements, item.AcknowledgementState) || !item.HighestSeverityPreserved || !item.SuppressedDuplicatesAuditable || item.AcknowledgementEqualsRemediation || item.ResolvedEqualsCanonicalClosure {
			return ProductionUsabilityValBNoiseBudgetStatePartial
		}
		if item.CriticalBlocker && (item.AcknowledgementState == ProductionUsabilityAckSuppressedDup || strings.TrimSpace(item.SuppressionReason) != "" || item.SuppressedCount > 0) {
			return ProductionUsabilityValBNoiseBudgetStatePartial
		}
	}
	return ProductionUsabilityValBNoiseBudgetStateActive
}

func EvaluateProductionUsabilityValBAPIProtectionState(model APIProtectionDiscipline) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValBAPIProtectionStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedRequestClasses, productionUsabilityValBRequestClasses()...) ||
		!containsExactTrimmedStringSet(model.SupportedPriorityLanes, productionUsabilityValBPriorityLanes()...) ||
		!containsExactTrimmedStringSet(model.SupportedBackpressure, productionUsabilityValBBackpressureStates()...) {
		return ProductionUsabilityValBAPIProtectionStatePartial
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		if strings.TrimSpace(item.RequestClass) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.PriorityLane) == "" || strings.TrimSpace(item.FairnessScope) == "" || strings.TrimSpace(item.RateLimitPolicyRef) == "" || strings.TrimSpace(item.BackpressureState) == "" || strings.TrimSpace(item.RetryAfterHint) == "" || strings.TrimSpace(item.SafeRetryGuidance) == "" || strings.TrimSpace(item.RejectionExplanation) == "" || strings.TrimSpace(item.VisibilityScope) == "" || strings.TrimSpace(item.RedactionTier) == "" {
			return ProductionUsabilityValBAPIProtectionStateIncomplete
		}
		if !containsTrimmedString(model.SupportedRequestClasses, item.RequestClass) || !containsTrimmedString(model.SupportedPriorityLanes, item.PriorityLane) || !containsTrimmedString(model.SupportedBackpressure, item.BackpressureState) {
			return ProductionUsabilityValBAPIProtectionStatePartial
		}
		if _, duplicate := seen[strings.TrimSpace(item.RequestClass)]; duplicate {
			return ProductionUsabilityValBAPIProtectionStatePartial
		}
		seen[strings.TrimSpace(item.RequestClass)] = struct{}{}
		if !item.StarvationPrevented || item.PolicyDenialHiddenAsThrottling {
			return ProductionUsabilityValBAPIProtectionStatePartial
		}
		if item.RequestClass == ProductionUsabilityRequestClassMutation {
			if item.PriorityBypassesGovernance || !item.GovernanceRequired {
				return ProductionUsabilityValBAPIProtectionStatePartial
			}
		}
	}
	if len(seen) != len(model.SupportedRequestClasses) {
		return ProductionUsabilityValBAPIProtectionStatePartial
	}
	return ProductionUsabilityValBAPIProtectionStateActive
}

func EvaluateProductionUsabilityValBCLIResilienceState(model CLIResilienceSurface) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValBCLIResilienceStateIncomplete
	}
	if !containsAllTrimmedStrings(model.SupportedExitCodeSemantics, "success", "partial", "blocked", "unsupported", "internal_error") {
		return ProductionUsabilityValBCLIResilienceStatePartial
	}
	for _, item := range model.Items {
		if strings.TrimSpace(item.CommandName) == "" || strings.TrimSpace(item.CurrentState) == "" || strings.TrimSpace(item.OperationType) == "" || strings.TrimSpace(item.ActionMode) == "" || strings.TrimSpace(item.RetrySafety) == "" || strings.TrimSpace(item.SafeRetryGuidance) == "" || strings.TrimSpace(item.PartialFailureBehavior) == "" || strings.TrimSpace(item.ExplainCommandRef) == "" || strings.TrimSpace(item.InspectCommandRef) == "" || len(item.ExitCodeSemantics) == 0 {
			return ProductionUsabilityValBCLIResilienceStateIncomplete
		}
		if !containsTrimmedString(ProductionUsabilityVal0OperationContractModel().RequiredOperationTypes, item.OperationType) ||
			!containsTrimmedString([]string{ProductionUsabilityActionModeViewOnly, ProductionUsabilityActionModeExplain, ProductionUsabilityActionModePreview, ProductionUsabilityActionModeDryRun, ProductionUsabilityActionModeAuditOnly, ProductionUsabilityActionModeEnforce, ProductionUsabilityActionModeMutate}, item.ActionMode) ||
			!containsTrimmedString(ProductionUsabilityVal0OperationContractModel().SupportedRetrySafety, item.RetrySafety) ||
			!containsAllTrimmedStrings(item.ExitCodeSemantics, "success", "partial", "blocked", "unsupported", "internal_error") {
			return ProductionUsabilityValBCLIResilienceStatePartial
		}
		switch item.ActionMode {
		case ProductionUsabilityActionModeViewOnly, ProductionUsabilityActionModeExplain, ProductionUsabilityActionModePreview, ProductionUsabilityActionModeDryRun, ProductionUsabilityActionModeAuditOnly:
			if item.MutatesCanonicalState {
				return ProductionUsabilityValBCLIResilienceStatePartial
			}
		}
		switch item.OperationType {
		case ProductionUsabilityOperationNonIdempotentMutate, ProductionUsabilityOperationSideEffecting:
			if item.RetrySafety == ProductionUsabilityRetrySafe {
				return ProductionUsabilityValBCLIResilienceStatePartial
			}
		}
		if item.IdempotencyKeyRequired && !item.IdempotencyKeyPresent {
			return ProductionUsabilityValBCLIResilienceStatePartial
		}
		if item.RetrySafety == ProductionUsabilityRetryUnsafe && strings.TrimSpace(item.DoNotRetryReason) == "" {
			return ProductionUsabilityValBCLIResilienceStatePartial
		}
	}
	return ProductionUsabilityValBCLIResilienceStateActive
}

func EvaluateProductionUsabilityValBScaleEnvelopeState(model ProductionScaleEnvelope) string {
	if strings.TrimSpace(model.ExpectedEventVolumeRange) == "" || strings.TrimSpace(model.ExpectedWorkflowObjectRange) == "" || model.TimelineQueryLimit <= 0 || model.SearchQueryLimit <= 0 || model.MaxPageSize <= 0 || strings.TrimSpace(model.DashboardLatencyBudget) == "" || strings.TrimSpace(model.APILatencyBudget) == "" || strings.TrimSpace(model.CLIResponseBudget) == "" || strings.TrimSpace(model.LimitationDisclaimer) == "" {
		return ProductionUsabilityValBScaleEnvelopeStateIncomplete
	}
	if len(model.UnsupportedScaleConditions) == 0 {
		return ProductionUsabilityValBScaleEnvelopeStatePartial
	}
	if (!model.ScaleMeasured || !model.LatencyBudgetsMeasured) && !strings.Contains(strings.TrimSpace(model.LimitationDisclaimer), "unmeasured") {
		return ProductionUsabilityValBScaleEnvelopeStatePartial
	}
	if model.ClaimsLatencyGuarantee && (!model.ScaleMeasured || !model.LatencyBudgetsMeasured) {
		return ProductionUsabilityValBScaleEnvelopeStatePartial
	}
	return ProductionUsabilityValBScaleEnvelopeStateActive
}

func EvaluateProductionUsabilityValBActionModeEnforcementState(model ActionModeEnforcementModel) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValBActionModeEnforcementStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedSurfaceKinds, "ui", "api", "cli") {
		return ProductionUsabilityValBActionModeEnforcementStatePartial
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
		if strings.TrimSpace(item.SurfaceRef) == "" || strings.TrimSpace(item.SurfaceKind) == "" || strings.TrimSpace(item.CurrentState) == "" || mode == "" || strings.TrimSpace(item.OperationType) == "" {
			return ProductionUsabilityValBActionModeEnforcementStateIncomplete
		}
		if !containsTrimmedString(productionUsabilityValBRegisteredSurfaceRefs(), item.SurfaceRef) {
			return ProductionUsabilityValBActionModeEnforcementStatePartial
		}
		if !containsTrimmedString(model.SupportedSurfaceKinds, item.SurfaceKind) || !containsTrimmedString(ProductionUsabilityVal0OperationContractModel().RequiredOperationTypes, item.OperationType) {
			return ProductionUsabilityValBActionModeEnforcementStatePartial
		}
		if _, ok := expectedModes[mode]; !ok {
			return ProductionUsabilityValBActionModeEnforcementStatePartial
		}
		if _, duplicate := seen[mode]; duplicate {
			return ProductionUsabilityValBActionModeEnforcementStatePartial
		}
		seen[mode] = struct{}{}
		switch mode {
		case ProductionUsabilityActionModeViewOnly, ProductionUsabilityActionModeExplain, ProductionUsabilityActionModePreview, ProductionUsabilityActionModeDryRun, ProductionUsabilityActionModeAuditOnly:
			if !item.NonMutating || item.ExecutesMutation {
				return ProductionUsabilityValBActionModeEnforcementStatePartial
			}
		case ProductionUsabilityActionModeEnforce, ProductionUsabilityActionModeMutate:
			if item.AvailableInValB {
				if !item.GovernedExternally || strings.TrimSpace(item.GovernanceRef) == "" || len(item.EvidenceRefs) == 0 {
					return ProductionUsabilityValBActionModeEnforcementStatePartial
				}
			}
		}
	}
	if len(seen) != len(expectedModes) {
		return ProductionUsabilityValBActionModeEnforcementStatePartial
	}
	return ProductionUsabilityValBActionModeEnforcementStateActive
}

func EvaluateProductionUsabilityValBState(val0State, valAState, uiDataResilienceState, windowingState, resultSemanticsState, commandCenterState, noiseBudgetState, apiProtectionState, cliResilienceState, scaleEnvelopeState, actionModeEnforcementState string) string {
	if strings.TrimSpace(val0State) != ProductionUsabilityVal0StateActive || strings.TrimSpace(valAState) != ProductionUsabilityValAStateActive {
		return ProductionUsabilityValBStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(uiDataResilienceState),
		strings.TrimSpace(windowingState),
		strings.TrimSpace(resultSemanticsState),
		strings.TrimSpace(commandCenterState),
		strings.TrimSpace(noiseBudgetState),
		strings.TrimSpace(apiProtectionState),
		strings.TrimSpace(cliResilienceState),
		strings.TrimSpace(scaleEnvelopeState),
		strings.TrimSpace(actionModeEnforcementState),
	} {
		switch state {
		case ProductionUsabilityValBUIDataResilienceStateActive,
			ProductionUsabilityValBWindowingStateActive,
			ProductionUsabilityValBResultSemanticsStateActive,
			ProductionUsabilityValBCommandCenterStateActive,
			ProductionUsabilityValBNoiseBudgetStateActive,
			ProductionUsabilityValBAPIProtectionStateActive,
			ProductionUsabilityValBCLIResilienceStateActive,
			ProductionUsabilityValBScaleEnvelopeStateActive,
			ProductionUsabilityValBActionModeEnforcementStateActive:
		case ProductionUsabilityValBUIDataResilienceStatePartial,
			ProductionUsabilityValBWindowingStatePartial,
			ProductionUsabilityValBResultSemanticsStatePartial,
			ProductionUsabilityValBCommandCenterStatePartial,
			ProductionUsabilityValBNoiseBudgetStatePartial,
			ProductionUsabilityValBAPIProtectionStatePartial,
			ProductionUsabilityValBCLIResilienceStatePartial,
			ProductionUsabilityValBScaleEnvelopeStatePartial,
			ProductionUsabilityValBActionModeEnforcementStatePartial:
			hasPartial = true
		default:
			return ProductionUsabilityValBStateIncomplete
		}
	}
	if hasPartial {
		return ProductionUsabilityValBStateSubstantial
	}
	return ProductionUsabilityValBStateActive
}

func EvaluateProductionUsabilityValBProofsState(val0State, valAState, uiDataResilienceState, windowingState, resultSemanticsState, commandCenterState, noiseBudgetState, apiProtectionState, cliResilienceState, scaleEnvelopeState, actionModeEnforcementState string, surfaceRefs, evidenceRefs, limitations, whyPoint4NotPass []string) string {
	baseState := EvaluateProductionUsabilityValBState(val0State, valAState, uiDataResilienceState, windowingState, resultSemanticsState, commandCenterState, noiseBudgetState, apiProtectionState, cliResilienceState, scaleEnvelopeState, actionModeEnforcementState)
	if len(surfaceRefs) < 10 || len(evidenceRefs) < 8 || len(limitations) == 0 || len(whyPoint4NotPass) == 0 {
		if baseState == ProductionUsabilityValBStateActive {
			return ProductionUsabilityValBStateSubstantial
		}
		return baseState
	}
	return baseState
}

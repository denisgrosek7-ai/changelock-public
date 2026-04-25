package operability

import "strings"

const (
	ProductionUsabilityValCReadinessStateActive     = "production_usability_valc_readiness_active"
	ProductionUsabilityValCReadinessStatePartial    = "production_usability_valc_readiness_partial"
	ProductionUsabilityValCReadinessStateIncomplete = "production_usability_valc_readiness_incomplete"

	ProductionUsabilityValCGuidedReadinessStateActive     = "production_usability_valc_guided_readiness_active"
	ProductionUsabilityValCGuidedReadinessStatePartial    = "production_usability_valc_guided_readiness_partial"
	ProductionUsabilityValCGuidedReadinessStateIncomplete = "production_usability_valc_guided_readiness_incomplete"

	ProductionUsabilityValCSupportBundleStateActive     = "production_usability_valc_support_bundle_active"
	ProductionUsabilityValCSupportBundleStatePartial    = "production_usability_valc_support_bundle_partial"
	ProductionUsabilityValCSupportBundleStateIncomplete = "production_usability_valc_support_bundle_incomplete"

	ProductionUsabilityValCDiagnosticsStateActive     = "production_usability_valc_diagnostics_active"
	ProductionUsabilityValCDiagnosticsStatePartial    = "production_usability_valc_diagnostics_partial"
	ProductionUsabilityValCDiagnosticsStateIncomplete = "production_usability_valc_diagnostics_incomplete"

	ProductionUsabilityValCHealthSnapshotStateActive     = "production_usability_valc_health_snapshot_active"
	ProductionUsabilityValCHealthSnapshotStatePartial    = "production_usability_valc_health_snapshot_partial"
	ProductionUsabilityValCHealthSnapshotStateIncomplete = "production_usability_valc_health_snapshot_incomplete"

	ProductionUsabilityValCRecoveryPlaybookStateActive     = "production_usability_valc_recovery_playbooks_active"
	ProductionUsabilityValCRecoveryPlaybookStatePartial    = "production_usability_valc_recovery_playbooks_partial"
	ProductionUsabilityValCRecoveryPlaybookStateIncomplete = "production_usability_valc_recovery_playbooks_incomplete"

	ProductionUsabilityValCUpgradeAdvisoryStateActive     = "production_usability_valc_upgrade_advisory_active"
	ProductionUsabilityValCUpgradeAdvisoryStatePartial    = "production_usability_valc_upgrade_advisory_partial"
	ProductionUsabilityValCUpgradeAdvisoryStateIncomplete = "production_usability_valc_upgrade_advisory_incomplete"

	ProductionUsabilityValCPermissionSupportStateActive     = "production_usability_valc_permission_support_flows_active"
	ProductionUsabilityValCPermissionSupportStatePartial    = "production_usability_valc_permission_support_flows_partial"
	ProductionUsabilityValCPermissionSupportStateIncomplete = "production_usability_valc_permission_support_flows_incomplete"

	ProductionUsabilityValCExportSafetyStateActive     = "production_usability_valc_redaction_export_safety_active"
	ProductionUsabilityValCExportSafetyStatePartial    = "production_usability_valc_redaction_export_safety_partial"
	ProductionUsabilityValCExportSafetyStateIncomplete = "production_usability_valc_redaction_export_safety_incomplete"

	ProductionUsabilityValCStateIncomplete  = "production_usability_valc_incomplete"
	ProductionUsabilityValCStateSubstantial = "production_usability_valc_substantially_ready"
	ProductionUsabilityValCStateActive      = "production_usability_valc_active"

	ProductionUsabilityReadinessPass        = "pass"
	ProductionUsabilityReadinessFail        = "fail"
	ProductionUsabilityReadinessWarning     = "warning"
	ProductionUsabilityReadinessDegraded    = "degraded"
	ProductionUsabilityReadinessUnsupported = "unsupported"
	ProductionUsabilityReadinessNotRun      = "not_run"

	ProductionUsabilityGuidedModeInstall         = "install"
	ProductionUsabilityGuidedModeFirstRun        = "first_run"
	ProductionUsabilityGuidedModeGoLive          = "go_live"
	ProductionUsabilityGuidedModeUpgradePrecheck = "upgrade_precheck"

	ProductionUsabilitySecretScanPassed = "passed"
	ProductionUsabilitySecretScanWarn   = "warning"
	ProductionUsabilitySecretScanFailed = "failed"

	ProductionUsabilityHealthHealthy     = "healthy"
	ProductionUsabilityHealthWarning     = "warning"
	ProductionUsabilityHealthDegraded    = "degraded"
	ProductionUsabilityHealthUnhealthy   = "unhealthy"
	ProductionUsabilityHealthUnavailable = "unavailable"
	ProductionUsabilityHealthUnsupported = "unsupported"

	ProductionUsabilityScenarioConfigFailure         = "config_failure"
	ProductionUsabilityScenarioPolicyDenial          = "policy_denial"
	ProductionUsabilityScenarioDegradedHealth        = "degraded_health"
	ProductionUsabilityScenarioStaleData             = "stale_data"
	ProductionUsabilityScenarioPartialResult         = "partial_result"
	ProductionUsabilityScenarioSupportEscalation     = "support_escalation"
	ProductionUsabilityScenarioUpgradeBlocked        = "upgrade_blocked"
	ProductionUsabilityScenarioRollbackNotAvailable  = "rollback_not_available"
	ProductionUsabilityScenarioUnsupportedCapability = "unsupported_capability"

	ProductionUsabilityAdvisoryPreview     = "preview"
	ProductionUsabilityAdvisoryAuditOnly   = "audit_only"
	ProductionUsabilityAdvisoryBlocked     = "blocked"
	ProductionUsabilityAdvisoryUnsupported = "unsupported"

	ProductionUsabilitySupportActionCollectBundle = "collect_bundle"
	ProductionUsabilitySupportActionShareRedacted = "share_redacted_bundle"
	ProductionUsabilitySupportActionEscalate      = "escalate"
)

type ReadinessCheck struct {
	ReadinessCheckID     string   `json:"readiness_check_id"`
	ReadinessScope       string   `json:"readiness_scope"`
	Status               string   `json:"status"`
	Blocking             bool     `json:"blocking"`
	ReasonCode           string   `json:"reason_code"`
	HumanMessage         string   `json:"human_message"`
	TechnicalDetail      string   `json:"technical_detail"`
	EvidenceRefs         []string `json:"evidence_refs,omitempty"`
	DependencyRefs       []string `json:"dependency_refs,omitempty"`
	GeneratedAtIndicator string   `json:"generated_at_indicator"`
	FreshnessState       string   `json:"freshness_state"`
	RecoveryHint         string   `json:"recovery_hint"`
	VisibilityScope      string   `json:"visibility_scope"`
	RedactionTier        string   `json:"redaction_tier"`
}

type ReadinessCheckModel struct {
	CurrentState         string           `json:"current_state"`
	SupportedStatuses    []string         `json:"supported_statuses,omitempty"`
	Items                []ReadinessCheck `json:"items,omitempty"`
	WarningsVisible      bool             `json:"warnings_visible"`
	ProjectionDisclaimer string           `json:"projection_disclaimer"`
	Limitations          []string         `json:"limitations,omitempty"`
}

type GuidedReadinessBaseline struct {
	CurrentState               string   `json:"current_state"`
	Mode                       string   `json:"mode"`
	RequiredSteps              []string `json:"required_steps,omitempty"`
	CompletedSteps             []string `json:"completed_steps,omitempty"`
	MissingSteps               []string `json:"missing_steps,omitempty"`
	BlockingSteps              []string `json:"blocking_steps,omitempty"`
	SampleConfigDetected       bool     `json:"sample_config_detected"`
	ProductionConfigDetected   bool     `json:"production_config_detected"`
	AutoProductionEnablement   bool     `json:"auto_production_enablement"`
	FakeDemoEvidenceDetected   bool     `json:"fake_demo_evidence_detected"`
	GoLiveAllowed              bool     `json:"go_live_allowed"`
	LimitationMessage          string   `json:"limitation_message"`
	NextStepGuidance           []string `json:"next_step_guidance,omitempty"`
	MutatesCanonicalState      bool     `json:"mutates_canonical_state"`
	ClaimsProductionCompletion bool     `json:"claims_production_completion"`
}

type SupportBundleQualityGate struct {
	CurrentState                 string   `json:"current_state"`
	BundleID                     string   `json:"bundle_id"`
	GeneratedAtIndicator         string   `json:"generated_at_indicator"`
	RequestedByActorRef          string   `json:"requested_by_actor_ref"`
	VisibilityScope              string   `json:"visibility_scope"`
	RedactionTier                string   `json:"redaction_tier"`
	IncludedSections             []string `json:"included_sections,omitempty"`
	ExcludedSections             []string `json:"excluded_sections,omitempty"`
	ExclusionReasons             []string `json:"exclusion_reasons,omitempty"`
	EvidenceRefsIncluded         []string `json:"evidence_refs_included,omitempty"`
	EvidenceRefsRedacted         []string `json:"evidence_refs_redacted,omitempty"`
	RedactedEvidenceRepresented  bool     `json:"redacted_evidence_represented"`
	ConfigSummaryIncluded        bool     `json:"config_summary_included"`
	HealthSnapshotIncluded       bool     `json:"health_snapshot_included"`
	DegradedStaleMarkersIncluded bool     `json:"degraded_stale_markers_included"`
	ReproducibilityMetadata      bool     `json:"reproducibility_metadata"`
	SecretScanResult             string   `json:"secret_scan_result"`
	RawSecretDetected            bool     `json:"raw_secret_detected"`
	RawTokenDetected             bool     `json:"raw_token_detected"`
	UnfilteredEnvDetected        bool     `json:"unfiltered_env_detected"`
	CacheClaimsCanonicalTruth    bool     `json:"cache_claims_canonical_truth"`
	ManifestPresent              bool     `json:"manifest_present"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type DiagnosticsHardeningModel struct {
	CurrentState                string   `json:"current_state"`
	DiagnosticsID               string   `json:"diagnostics_id"`
	GeneratedAtIndicator        string   `json:"generated_at_indicator"`
	Sections                    []string `json:"sections,omitempty"`
	PermissionScope             string   `json:"permission_scope"`
	RedactionTier               string   `json:"redaction_tier"`
	SafeToShare                 bool     `json:"safe_to_share"`
	SensitiveFieldsRedacted     bool     `json:"sensitive_fields_redacted"`
	SecretScanStatus            string   `json:"secret_scan_status"`
	UnsupportedSections         []string `json:"unsupported_sections,omitempty"`
	StaleSections               []string `json:"stale_sections,omitempty"`
	PartialSections             []string `json:"partial_sections,omitempty"`
	UnsupportedSectionsExplicit bool     `json:"unsupported_sections_explicit"`
	StaleSectionsExplicit       bool     `json:"stale_sections_explicit"`
	PartialSectionsExplicit     bool     `json:"partial_sections_explicit"`
	LimitationMessage           string   `json:"limitation_message"`
	RecoveryHint                string   `json:"recovery_hint"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
}

type HealthSnapshotModel struct {
	CurrentState          string            `json:"current_state"`
	SnapshotID            string            `json:"snapshot_id"`
	GeneratedAtIndicator  string            `json:"generated_at_indicator"`
	HealthState           string            `json:"health_state"`
	SupportedHealthStates []string          `json:"supported_health_states,omitempty"`
	ComponentStates       map[string]string `json:"component_states,omitempty"`
	StaleComponents       []string          `json:"stale_components,omitempty"`
	DegradedComponents    []string          `json:"degraded_components,omitempty"`
	UnavailableComponents []string          `json:"unavailable_components,omitempty"`
	UnsupportedComponents []string          `json:"unsupported_components,omitempty"`
	EvidenceRefs          []string          `json:"evidence_refs,omitempty"`
	FreshnessState        string            `json:"freshness_state"`
	LimitationMessage     string            `json:"limitation_message"`
	RecoveryHint          string            `json:"recovery_hint"`
	ProjectionDisclaimer  string            `json:"projection_disclaimer"`
}

type RecoveryPlaybook struct {
	PlaybookID            string   `json:"playbook_id"`
	Scenario              string   `json:"scenario"`
	RemediationClass      string   `json:"remediation_class"`
	SafeSteps             []string `json:"safe_steps,omitempty"`
	UnsafeSteps           []string `json:"unsafe_steps,omitempty"`
	DoNotRetryReason      string   `json:"do_not_retry_reason"`
	EscalationPath        string   `json:"escalation_path"`
	RollbackHint          string   `json:"rollback_hint"`
	InspectCommandRef     string   `json:"inspect_command_ref"`
	ExplainCommandRef     string   `json:"explain_command_ref"`
	SupportBundleRef      string   `json:"support_bundle_ref"`
	RequiredPermissions   []string `json:"required_permissions,omitempty"`
	EvidenceRefs          []string `json:"evidence_refs,omitempty"`
	LimitationMessage     string   `json:"limitation_message"`
	PolicyBypassSuggested bool     `json:"policy_bypass_suggested"`
	UnsafeRetrySuggested  bool     `json:"unsafe_retry_suggested"`
}

type RecoveryPlaybookModel struct {
	CurrentState          string             `json:"current_state"`
	SupportedScenarios    []string           `json:"supported_scenarios,omitempty"`
	SupportedRemediations []string           `json:"supported_remediations,omitempty"`
	Items                 []RecoveryPlaybook `json:"items,omitempty"`
	Limitations           []string           `json:"limitations,omitempty"`
}

type UpgradeRollbackAdvisory struct {
	CurrentState              string   `json:"current_state"`
	AdvisoryID                string   `json:"advisory_id"`
	CurrentVersion            string   `json:"current_version"`
	TargetVersion             string   `json:"target_version"`
	KnownTargetVersions       []string `json:"known_target_versions,omitempty"`
	ConfigPolicySchemaImpact  []string `json:"config_policy_schema_impact,omitempty"`
	CompatibilityStatus       string   `json:"compatibility_status"`
	DeprecatedItems           []string `json:"deprecated_items,omitempty"`
	RemovedOrUnsupportedItems []string `json:"removed_or_unsupported_items,omitempty"`
	MigrationRequiredItems    []string `json:"migration_required_items,omitempty"`
	BlockingIssues            []string `json:"blocking_issues,omitempty"`
	WarningIssues             []string `json:"warning_issues,omitempty"`
	RollbackAvailable         bool     `json:"rollback_available"`
	RollbackLimitations       []string `json:"rollback_limitations,omitempty"`
	RollbackBlockers          []string `json:"rollback_blockers,omitempty"`
	RollbackScope             string   `json:"rollback_scope"`
	AdvisoryMode              string   `json:"advisory_mode"`
	MutatesState              bool     `json:"mutates_state"`
	GeneratedAtIndicator      string   `json:"generated_at_indicator"`
	LimitationDisclaimer      string   `json:"limitation_disclaimer"`
	ApprovalImplied           bool     `json:"approval_implied"`
}

type PermissionAwareSupportFlow struct {
	RequesterRole         string   `json:"requester_role"`
	VisibilityScope       string   `json:"visibility_scope"`
	CurrentState          string   `json:"current_state"`
	AllowedSections       []string `json:"allowed_sections,omitempty"`
	RedactedSections      []string `json:"redacted_sections,omitempty"`
	HiddenSections        []string `json:"hidden_sections,omitempty"`
	EvidenceVisibility    string   `json:"evidence_visibility"`
	SafeFallbackMessage   string   `json:"safe_fallback_message"`
	EscalationRequired    bool     `json:"escalation_required"`
	SupportActionMode     string   `json:"support_action_mode"`
	MutatesCanonicalState bool     `json:"mutates_canonical_state"`
}

type PermissionAwareSupportFlowModel struct {
	CurrentState                string                       `json:"current_state"`
	SupportedEvidenceVisibility []string                     `json:"supported_evidence_visibility,omitempty"`
	SupportedActionModes        []string                     `json:"supported_action_modes,omitempty"`
	Items                       []PermissionAwareSupportFlow `json:"items,omitempty"`
	Limitations                 []string                     `json:"limitations,omitempty"`
}

type RedactionSafeExportModel struct {
	CurrentState                string   `json:"current_state"`
	ExportID                    string   `json:"export_id"`
	ExportScope                 string   `json:"export_scope"`
	RedactionTier               string   `json:"redaction_tier"`
	SecretScanStatus            string   `json:"secret_scan_status"`
	AllowedContentClasses       []string `json:"allowed_content_classes,omitempty"`
	BlockedContentClasses       []string `json:"blocked_content_classes,omitempty"`
	EvidenceHandling            string   `json:"evidence_handling"`
	PublicSafe                  bool     `json:"public_safe"`
	PartnerSafe                 bool     `json:"partner_safe"`
	AuditorSafe                 bool     `json:"auditor_safe"`
	RawSecretDetected           bool     `json:"raw_secret_detected"`
	RawInternalEvidenceIncluded bool     `json:"raw_internal_evidence_included"`
	PolicyAllowsExport          bool     `json:"policy_allows_export"`
	LimitationMessage           string   `json:"limitation_message"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
}

func productionUsabilityValCReadinessStatuses() []string {
	return []string{
		ProductionUsabilityReadinessPass,
		ProductionUsabilityReadinessFail,
		ProductionUsabilityReadinessWarning,
		ProductionUsabilityReadinessDegraded,
		ProductionUsabilityReadinessUnsupported,
		ProductionUsabilityReadinessNotRun,
	}
}

func productionUsabilityValCGuidedModes() []string {
	return []string{
		ProductionUsabilityGuidedModeInstall,
		ProductionUsabilityGuidedModeFirstRun,
		ProductionUsabilityGuidedModeGoLive,
		ProductionUsabilityGuidedModeUpgradePrecheck,
	}
}

func productionUsabilityValCSecretScanStatuses() []string {
	return []string{
		ProductionUsabilitySecretScanPassed,
		ProductionUsabilitySecretScanWarn,
		ProductionUsabilitySecretScanFailed,
	}
}

func productionUsabilityValCHealthStates() []string {
	return []string{
		ProductionUsabilityHealthHealthy,
		ProductionUsabilityHealthWarning,
		ProductionUsabilityHealthDegraded,
		ProductionUsabilityHealthUnhealthy,
		ProductionUsabilityHealthUnavailable,
		ProductionUsabilityHealthUnsupported,
	}
}

func productionUsabilityValCRecoveryScenarios() []string {
	return []string{
		ProductionUsabilityScenarioConfigFailure,
		ProductionUsabilityScenarioPolicyDenial,
		ProductionUsabilityScenarioDegradedHealth,
		ProductionUsabilityScenarioStaleData,
		ProductionUsabilityScenarioPartialResult,
		ProductionUsabilityScenarioSupportEscalation,
		ProductionUsabilityScenarioUpgradeBlocked,
		ProductionUsabilityScenarioRollbackNotAvailable,
		ProductionUsabilityScenarioUnsupportedCapability,
	}
}

func productionUsabilityValCSupportActionModes() []string {
	return []string{
		ProductionUsabilityActionModeViewOnly,
		ProductionUsabilityActionModeExplain,
		ProductionUsabilitySupportActionCollectBundle,
		ProductionUsabilitySupportActionShareRedacted,
		ProductionUsabilitySupportActionEscalate,
	}
}

func productionUsabilityValCAdvisoryModes() []string {
	return []string{
		ProductionUsabilityAdvisoryPreview,
		ProductionUsabilityAdvisoryAuditOnly,
		ProductionUsabilityAdvisoryBlocked,
		ProductionUsabilityAdvisoryUnsupported,
	}
}

func productionUsabilityValCRegisteredSurfaceRefs() []string {
	return []string{
		"/v1/production/usability-operability-recovery/valc/readiness",
		"/v1/production/usability-operability-recovery/valc/guided-readiness",
		"/v1/production/usability-operability-recovery/valc/support-bundle",
		"/v1/production/usability-operability-recovery/valc/diagnostics",
		"/v1/production/usability-operability-recovery/valc/health-snapshot",
		"/v1/production/usability-operability-recovery/valc/recovery-playbooks",
		"/v1/production/usability-operability-recovery/valc/upgrade-rollback-advisory",
		"/v1/production/usability-operability-recovery/valc/permission-support-flows",
		"/v1/production/usability-operability-recovery/valc/redaction-export-safety",
		"/v1/production/usability-operability-recovery/valc/proofs",
	}
}

func ProductionUsabilityValCReadinessChecks() ReadinessCheckModel {
	return ReadinessCheckModel{
		CurrentState:      "readiness_checks_ready",
		SupportedStatuses: productionUsabilityValCReadinessStatuses(),
		Items: []ReadinessCheck{
			{ReadinessCheckID: "readiness-config", ReadinessScope: "config_integrity", Status: ProductionUsabilityReadinessPass, Blocking: true, ReasonCode: "config_valid", HumanMessage: "Production config integrity passed the bounded readiness check.", TechnicalDetail: "Schema validation, fail-fast bootstrap readiness, and unknown-field policy are aligned for the current configuration set.", EvidenceRefs: []string{"vala_config_factory", "val0_config_integrity"}, DependencyRefs: []string{"val0", "vala"}, GeneratedAtIndicator: "generated_at_present", FreshnessState: ProductionUsabilityStatusFresh, RecoveryHint: "re-run readiness after material config changes", VisibilityScope: ProductionUsabilityVisibilityOperator, RedactionTier: ProductionUsabilityRedactionLow},
			{ReadinessCheckID: "readiness-policy", ReadinessScope: "policy_schema", Status: ProductionUsabilityReadinessPass, Blocking: true, ReasonCode: "policy_supported", HumanMessage: "Policy schema is supported for the current bounded preview path.", TechnicalDetail: "Policy schema compatibility and migration status remain compatible with the current production usability baseline.", EvidenceRefs: []string{"vala_policy_schema"}, DependencyRefs: []string{"vala"}, GeneratedAtIndicator: "generated_at_present", FreshnessState: ProductionUsabilityStatusFresh, RecoveryHint: "re-run readiness after policy bundle upgrades", VisibilityScope: ProductionUsabilityVisibilityInternalAdmin, RedactionTier: ProductionUsabilityRedactionNone},
			{ReadinessCheckID: "readiness-deprecated-warning", ReadinessScope: "upgrade_advisory", Status: ProductionUsabilityReadinessWarning, Blocking: false, ReasonCode: "deprecated_schema_notice", HumanMessage: "A bounded upgrade advisory warning remains visible for deprecated compatibility paths.", TechnicalDetail: "Deprecated configuration fields do not block current readiness but require upgrade planning before future schema enforcement.", DependencyRefs: []string{"vala", "valb"}, GeneratedAtIndicator: "generated_at_present", FreshnessState: ProductionUsabilityStatusFresh, RecoveryHint: "review upgrade advisory output before the next planned version move", VisibilityScope: ProductionUsabilityVisibilityOperator, RedactionTier: ProductionUsabilityRedactionLow},
		},
		WarningsVisible:      true,
		ProjectionDisclaimer: "readiness_projection_is_not_canonical_truth",
		Limitations: []string{
			"Val C readiness checks are bounded production projections only and do not replace canonical workflow or proof state.",
		},
	}
}

func ProductionUsabilityValCGuidedReadiness() GuidedReadinessBaseline {
	return GuidedReadinessBaseline{
		CurrentState:               "guided_readiness_ready",
		Mode:                       ProductionUsabilityGuidedModeGoLive,
		RequiredSteps:              []string{"config_validated", "policy_validated", "health_snapshot_reviewed", "support_bundle_path_verified"},
		CompletedSteps:             []string{"config_validated", "policy_validated", "health_snapshot_reviewed", "support_bundle_path_verified"},
		MissingSteps:               nil,
		BlockingSteps:              nil,
		SampleConfigDetected:       false,
		ProductionConfigDetected:   true,
		AutoProductionEnablement:   false,
		FakeDemoEvidenceDetected:   false,
		GoLiveAllowed:              true,
		LimitationMessage:          "guided_go_live_baseline_is_projection_only_and_not_full_production_completion",
		NextStepGuidance:           []string{"capture bounded health snapshot before go-live window", "collect a redaction-safe support bundle for recovery readiness"},
		MutatesCanonicalState:      false,
		ClaimsProductionCompletion: false,
	}
}

func ProductionUsabilityValCSupportBundleQualityGate() SupportBundleQualityGate {
	return SupportBundleQualityGate{
		CurrentState:                 "support_bundle_quality_gate_ready",
		BundleID:                     "bundle-acme-prod-001",
		GeneratedAtIndicator:         "generated_at_present",
		RequestedByActorRef:          "actor/support-admin",
		VisibilityScope:              ProductionUsabilityVisibilityInternalAdmin,
		RedactionTier:                ProductionUsabilityRedactionHigh,
		IncludedSections:             []string{"manifest", "config_summary", "health_snapshot", "diagnostics_metadata", "readiness_summary"},
		ExcludedSections:             []string{"raw_env_dump", "raw_token_material"},
		ExclusionReasons:             []string{"secret_bearing_content_is_blocked", "token_material_is_blocked"},
		EvidenceRefsIncluded:         []string{"readiness_projection", "health_snapshot_projection"},
		EvidenceRefsRedacted:         []string{"canonical_evidence.redacted_metadata"},
		RedactedEvidenceRepresented:  true,
		ConfigSummaryIncluded:        true,
		HealthSnapshotIncluded:       true,
		DegradedStaleMarkersIncluded: true,
		ReproducibilityMetadata:      true,
		SecretScanResult:             ProductionUsabilitySecretScanPassed,
		RawSecretDetected:            false,
		RawTokenDetected:             false,
		UnfilteredEnvDetected:        false,
		CacheClaimsCanonicalTruth:    false,
		ManifestPresent:              true,
		ProjectionDisclaimer:         "support_bundle_projection_is_not_canonical_truth",
	}
}

func ProductionUsabilityValCDiagnosticsHardening() DiagnosticsHardeningModel {
	return DiagnosticsHardeningModel{
		CurrentState:                "diagnostics_hardening_ready",
		DiagnosticsID:               "diag-acme-prod-001",
		GeneratedAtIndicator:        "generated_at_present",
		Sections:                    []string{"config_summary", "health_snapshot", "api_resilience", "search_index", "public_export_capability"},
		PermissionScope:             ProductionUsabilityVisibilityInternalAdmin,
		RedactionTier:               ProductionUsabilityRedactionHigh,
		SafeToShare:                 true,
		SensitiveFieldsRedacted:     true,
		SecretScanStatus:            ProductionUsabilitySecretScanPassed,
		UnsupportedSections:         []string{"public_export_capability"},
		StaleSections:               []string{"search_index"},
		PartialSections:             []string{"api_resilience"},
		UnsupportedSectionsExplicit: true,
		StaleSectionsExplicit:       true,
		PartialSectionsExplicit:     true,
		LimitationMessage:           "diagnostics_output_is_projection_only_and_contains_explicit_stale_partial_and_unsupported_markers",
		RecoveryHint:                "collect_fresh_diagnostics_after_recovering_stale_or_partial_components",
		EvidenceRefs:                []string{"diagnostics_projection", "health_snapshot_projection"},
		ProjectionDisclaimer:        "diagnostics_projection_is_not_canonical_truth",
	}
}

func ProductionUsabilityValCHealthSnapshot() HealthSnapshotModel {
	return HealthSnapshotModel{
		CurrentState:          "health_snapshot_ready",
		SnapshotID:            "health-acme-prod-001",
		GeneratedAtIndicator:  "generated_at_present",
		HealthState:           ProductionUsabilityHealthHealthy,
		SupportedHealthStates: productionUsabilityValCHealthStates(),
		ComponentStates: map[string]string{
			"config_factory": ProductionUsabilityHealthHealthy,
			"policy_schema":  ProductionUsabilityHealthHealthy,
			"ui_resilience":  ProductionUsabilityHealthHealthy,
			"support_bundle": ProductionUsabilityHealthHealthy,
		},
		StaleComponents:       nil,
		DegradedComponents:    nil,
		UnavailableComponents: nil,
		UnsupportedComponents: nil,
		EvidenceRefs:          []string{"health_snapshot_projection"},
		FreshnessState:        ProductionUsabilityStatusFresh,
		LimitationMessage:     "health_snapshot_is_point_in_time_projection_only",
		RecoveryHint:          "capture_new_snapshot_when_component_state_changes",
		ProjectionDisclaimer:  "health_snapshot_is_not_canonical_truth",
	}
}

func ProductionUsabilityValCRecoveryPlaybooks() RecoveryPlaybookModel {
	return RecoveryPlaybookModel{
		CurrentState:          "recovery_playbooks_ready",
		SupportedScenarios:    productionUsabilityValCRecoveryScenarios(),
		SupportedRemediations: ProductionUsabilityVal0RecoveryUXContract().RemediationClasses,
		Items: []RecoveryPlaybook{
			{PlaybookID: "playbook-config-failure", Scenario: ProductionUsabilityScenarioConfigFailure, RemediationClass: ProductionUsabilityRemediationConfigFix, SafeSteps: []string{"run_changelock_config_inspect", "apply_schema_declared_fix"}, UnsafeSteps: []string{"edit_canonical_evidence_directly"}, DoNotRetryReason: "retry_without_fix_repeats_same_config_failure", EscalationPath: "platform_config_owners", RollbackHint: "rollback_limited_to_config_policy_supportability_only", InspectCommandRef: "changelock config inspect", ExplainCommandRef: "changelock config explain", RequiredPermissions: []string{"operator"}, LimitationMessage: "recovery_guidance_is_advisory_only"},
			{PlaybookID: "playbook-policy-denial", Scenario: ProductionUsabilityScenarioPolicyDenial, RemediationClass: ProductionUsabilityRemediationPolicyUpdate, SafeSteps: []string{"inspect_policy_denial_reason", "request_governed_policy_update"}, UnsafeSteps: []string{"bypass_policy_gate"}, DoNotRetryReason: "retry_without_policy_change_repeats_denial", EscalationPath: "policy_owners", RollbackHint: "policy_rollback_requires_supported_schema_path", InspectCommandRef: "changelock policy inspect", ExplainCommandRef: "changelock policy explain", RequiredPermissions: []string{"operator", "governance"}, LimitationMessage: "recovery_guidance_does_not_authorize_policy_bypass"},
			{PlaybookID: "playbook-degraded-health", Scenario: ProductionUsabilityScenarioDegradedHealth, RemediationClass: ProductionUsabilityRemediationWaitAndRetry, SafeSteps: []string{"inspect_health_snapshot", "recheck_component_after_degradation_clears"}, UnsafeSteps: []string{"assume_degraded_component_is_healthy"}, DoNotRetryReason: "retry_unsafe_if_component_is_side_effecting_and_state_unknown", EscalationPath: "support_engineering", RollbackHint: "rollback_hint_applies_only_when_supported_and_governed", InspectCommandRef: "changelock health inspect", ExplainCommandRef: "changelock health explain", SupportBundleRef: "bundle-acme-prod-001", RequiredPermissions: []string{"operator"}, LimitationMessage: "degraded_health_guidance_is_projection_only"},
			{PlaybookID: "playbook-stale-data", Scenario: ProductionUsabilityScenarioStaleData, RemediationClass: ProductionUsabilityRemediationEvidenceRefresh, SafeSteps: []string{"refresh_projection", "requery_support_surface"}, UnsafeSteps: []string{"treat_stale_projection_as_fresh"}, DoNotRetryReason: "blind_retry_without_refresh_can_repeat_stale_view", EscalationPath: "support_engineering", RollbackHint: "not_applicable_for_stale_projection_only", InspectCommandRef: "changelock support inspect", ExplainCommandRef: "changelock support explain", RequiredPermissions: []string{"operator"}, LimitationMessage: "stale_data_playbook_requires_bounded_refresh_only"},
			{PlaybookID: "playbook-partial-result", Scenario: ProductionUsabilityScenarioPartialResult, RemediationClass: ProductionUsabilityRemediationWaitAndRetry, SafeSteps: []string{"load_additional_segments", "retry_after_missing_component_recovers"}, UnsafeSteps: []string{"claim_partial_result_is_complete"}, DoNotRetryReason: "retry_without_missing_component_context_can_repeat_partial_result", EscalationPath: "support_engineering", RollbackHint: "not_applicable_for_partial_projection_only", InspectCommandRef: "changelock support inspect", ExplainCommandRef: "changelock support explain", RequiredPermissions: []string{"operator"}, LimitationMessage: "partial_result_playbook_is_guidance_only"},
			{PlaybookID: "playbook-support-escalation", Scenario: ProductionUsabilityScenarioSupportEscalation, RemediationClass: ProductionUsabilityRemediationSupportEsc, SafeSteps: []string{"collect_redacted_support_bundle", "escalate_with_manifest_and_health_snapshot"}, UnsafeSteps: []string{"share_unredacted_bundle_outside_scope"}, DoNotRetryReason: "do_not_retry_until_support_path_confirms_required_scope", EscalationPath: "support_engineering", RollbackHint: "support_escalation_does_not_change_runtime_state", InspectCommandRef: "changelock support inspect", ExplainCommandRef: "changelock support explain", SupportBundleRef: "bundle-acme-prod-001", RequiredPermissions: []string{"support"}, LimitationMessage: "support_escalation_is_not_approval"},
			{PlaybookID: "playbook-upgrade-blocked", Scenario: ProductionUsabilityScenarioUpgradeBlocked, RemediationClass: ProductionUsabilityRemediationManualReview, SafeSteps: []string{"review_blocking_upgrade_issues", "resolve_schema_or_policy_conflicts_before_retry"}, UnsafeSteps: []string{"force_unsupported_upgrade_path"}, DoNotRetryReason: "retry_without_resolving_blockers_repeats_upgrade_failure", EscalationPath: "release_engineering", RollbackHint: "rollback_availability_depends_on_bounded_config_policy_scope", InspectCommandRef: "changelock upgrade preview", ExplainCommandRef: "changelock upgrade explain", RequiredPermissions: []string{"admin"}, LimitationMessage: "upgrade_blocked_playbook_is_advisory_only"},
			{PlaybookID: "playbook-rollback-not-available", Scenario: ProductionUsabilityScenarioRollbackNotAvailable, RemediationClass: ProductionUsabilityRemediationUnsupported, SafeSteps: []string{"review_rollback_scope_limitations", "escalate_for_supported_recovery_path"}, UnsafeSteps: []string{"assume_rollback_exists_without_supported_scope"}, DoNotRetryReason: "unsupported_rollback_path_cannot_be_made_safe_by_retry", EscalationPath: "release_engineering", RollbackHint: "rollback_unavailable_until_supported_schema_and_supportability_scope_exists", InspectCommandRef: "changelock upgrade preview", ExplainCommandRef: "changelock upgrade explain", RequiredPermissions: []string{"admin"}, LimitationMessage: "rollback_unavailable_is_explicitly_unsupported"},
			{PlaybookID: "playbook-unsupported-capability", Scenario: ProductionUsabilityScenarioUnsupportedCapability, RemediationClass: ProductionUsabilityRemediationUnsupported, SafeSteps: []string{"use_supported_surface_or_scope"}, UnsafeSteps: []string{"pretend_unsupported_capability_is_available"}, DoNotRetryReason: "unsupported_capability_stays_unsupported_until_support_exists", EscalationPath: "product_supportability_owners", RollbackHint: "not_applicable_until_supported_capability_exists", InspectCommandRef: "changelock support inspect", ExplainCommandRef: "changelock support explain", RequiredPermissions: []string{"operator"}, LimitationMessage: "unsupported_capability_guidance_is_bounded_and_non-authorizing"},
		},
		Limitations: []string{
			"Val C recovery playbooks are advisory only and do not authorize policy bypass, evidence edits, or unsupported lifecycle operations.",
		},
	}
}

func ProductionUsabilityValCUpgradeRollbackAdvisory() UpgradeRollbackAdvisory {
	return UpgradeRollbackAdvisory{
		CurrentState:              "upgrade_rollback_advisory_ready",
		AdvisoryID:                "upgrade-advisory-acme-prod-001",
		CurrentVersion:            "2026.04.1",
		TargetVersion:             "2026.05.0",
		KnownTargetVersions:       []string{"2026.05.0", "2026.05.1"},
		ConfigPolicySchemaImpact:  []string{"config_schema_current", "policy_schema_current"},
		CompatibilityStatus:       ProductionUsabilityCompatibilityCurrent,
		DeprecatedItems:           []string{"legacy_policy_toggle"},
		RemovedOrUnsupportedItems: []string{"none"},
		MigrationRequiredItems:    []string{"document_future_policy_schema_deprecation"},
		BlockingIssues:            nil,
		WarningIssues:             []string{"rollback_scope_limited_to_config_policy_supportability_only"},
		RollbackAvailable:         true,
		RollbackLimitations:       []string{"rollback_does_not_cover_runtime_or_external_ticket_side_effects"},
		RollbackBlockers:          nil,
		RollbackScope:             "config_policy_supportability_only",
		AdvisoryMode:              ProductionUsabilityAdvisoryPreview,
		MutatesState:              false,
		GeneratedAtIndicator:      "generated_at_present",
		LimitationDisclaimer:      "upgrade_rollback_advisory_is_projection_only_baseline",
		ApprovalImplied:           false,
	}
}

func ProductionUsabilityValCPermissionSupportFlows() PermissionAwareSupportFlowModel {
	return PermissionAwareSupportFlowModel{
		CurrentState:                "permission_support_flows_ready",
		SupportedEvidenceVisibility: []string{ProductionUsabilityEvidenceFull, ProductionUsabilityEvidenceMetadataOnly, ProductionUsabilityEvidenceRedacted, ProductionUsabilityEvidenceHidden},
		SupportedActionModes:        productionUsabilityValCSupportActionModes(),
		Items: []PermissionAwareSupportFlow{
			{RequesterRole: "internal_admin", VisibilityScope: ProductionUsabilityVisibilityInternalAdmin, CurrentState: "support_flow_ready", AllowedSections: []string{"manifest", "config_summary", "health_snapshot", "diagnostics"}, RedactedSections: nil, HiddenSections: nil, EvidenceVisibility: ProductionUsabilityEvidenceFull, SafeFallbackMessage: "Internal admin support flow can review the full bounded support bundle.", EscalationRequired: false, SupportActionMode: ProductionUsabilitySupportActionCollectBundle, MutatesCanonicalState: false},
			{RequesterRole: "operator", VisibilityScope: ProductionUsabilityVisibilityOperator, CurrentState: "support_flow_ready", AllowedSections: []string{"manifest", "config_summary", "health_snapshot"}, RedactedSections: []string{"diagnostics_sensitive_fields"}, HiddenSections: []string{"raw_secret_material"}, EvidenceVisibility: ProductionUsabilityEvidenceMetadataOnly, SafeFallbackMessage: "Operator support flow remains bounded to metadata and redacted sections.", EscalationRequired: false, SupportActionMode: ProductionUsabilityActionModeExplain, MutatesCanonicalState: false},
			{RequesterRole: "developer", VisibilityScope: ProductionUsabilityVisibilityDeveloper, CurrentState: "support_flow_ready", AllowedSections: []string{"manifest", "health_snapshot", "diagnostics"}, RedactedSections: []string{"config_sensitive_values"}, HiddenSections: []string{"raw_token_material"}, EvidenceVisibility: ProductionUsabilityEvidenceMetadataOnly, SafeFallbackMessage: "Developer support flow exposes actionable diagnostics without raw protected evidence.", EscalationRequired: false, SupportActionMode: ProductionUsabilityActionModeViewOnly, MutatesCanonicalState: false},
			{RequesterRole: "partner_support", VisibilityScope: ProductionUsabilityVisibilityPartner, CurrentState: "support_flow_ready", AllowedSections: []string{"manifest", "health_summary"}, RedactedSections: []string{"diagnostics", "config_summary"}, HiddenSections: []string{"evidence_refs", "raw_env_dump"}, EvidenceVisibility: ProductionUsabilityEvidenceRedacted, SafeFallbackMessage: "Partner flow receives a redacted support summary and must escalate for internal-only detail.", EscalationRequired: true, SupportActionMode: ProductionUsabilitySupportActionShareRedacted, MutatesCanonicalState: false},
			{RequesterRole: "public_observer", VisibilityScope: ProductionUsabilityVisibilityPublicSafe, CurrentState: "support_flow_ready", AllowedSections: []string{"manifest_summary"}, RedactedSections: []string{"health_snapshot"}, HiddenSections: []string{"config_summary", "evidence_refs", "diagnostics"}, EvidenceVisibility: ProductionUsabilityEvidenceHidden, SafeFallbackMessage: "Public-safe flow discloses only bounded status and explicit hidden-section markers.", EscalationRequired: true, SupportActionMode: ProductionUsabilitySupportActionEscalate, MutatesCanonicalState: false},
		},
		Limitations: []string{
			"Val C support flows remain bounded, non-mutating, and permission-aware; escalation is not approval and hidden sections remain visible as hidden.",
		},
	}
}

func ProductionUsabilityValCRedactionSafeExport() RedactionSafeExportModel {
	return RedactionSafeExportModel{
		CurrentState:                "redaction_export_safety_ready",
		ExportID:                    "export-acme-prod-001",
		ExportScope:                 "auditor_support_bundle",
		RedactionTier:               ProductionUsabilityRedactionHigh,
		SecretScanStatus:            ProductionUsabilitySecretScanPassed,
		AllowedContentClasses:       []string{"manifest", "config_summary", "health_snapshot_metadata", "diagnostics_metadata"},
		BlockedContentClasses:       []string{"raw_secrets", "raw_tokens", "private_keys", "unfiltered_env"},
		EvidenceHandling:            ProductionUsabilityEvidenceMetadataOnly,
		PublicSafe:                  false,
		PartnerSafe:                 false,
		AuditorSafe:                 true,
		RawSecretDetected:           false,
		RawInternalEvidenceIncluded: false,
		PolicyAllowsExport:          true,
		LimitationMessage:           "export_baseline_is_projection_only_and_scope_limited",
		ProjectionDisclaimer:        "export_projection_is_not_canonical_truth",
	}
}

func EvaluateProductionUsabilityValCReadinessState(model ReadinessCheckModel) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValCReadinessStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedStatuses, productionUsabilityValCReadinessStatuses()...) ||
		!model.WarningsVisible ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValCReadinessStatePartial
	}
	hasWarning := false
	for _, item := range model.Items {
		status := strings.TrimSpace(item.Status)
		if strings.TrimSpace(item.ReadinessCheckID) == "" || strings.TrimSpace(item.ReadinessScope) == "" || status == "" || strings.TrimSpace(item.ReasonCode) == "" || strings.TrimSpace(item.HumanMessage) == "" || strings.TrimSpace(item.TechnicalDetail) == "" || strings.TrimSpace(item.GeneratedAtIndicator) == "" || strings.TrimSpace(item.FreshnessState) == "" || strings.TrimSpace(item.RecoveryHint) == "" || strings.TrimSpace(item.VisibilityScope) == "" || strings.TrimSpace(item.RedactionTier) == "" {
			return ProductionUsabilityValCReadinessStateIncomplete
		}
		if !containsTrimmedString(model.SupportedStatuses, status) ||
			!containsTrimmedString([]string{ProductionUsabilityStatusFresh, ProductionUsabilityStatusStale, ProductionUsabilityStatusPartial, ProductionUsabilityStatusDegraded, ProductionUsabilityStatusUnavailable, ProductionUsabilityStatusUnsupported}, item.FreshnessState) ||
			!containsTrimmedString([]string{ProductionUsabilityVisibilityInternalAdmin, ProductionUsabilityVisibilityOperator, ProductionUsabilityVisibilityDeveloper, ProductionUsabilityVisibilityPartner, ProductionUsabilityVisibilityPublicSafe}, item.VisibilityScope) ||
			!containsTrimmedString([]string{ProductionUsabilityRedactionNone, ProductionUsabilityRedactionLow, ProductionUsabilityRedactionMedium, ProductionUsabilityRedactionHigh, ProductionUsabilityRedactionPublicSafe}, item.RedactionTier) {
			return ProductionUsabilityValCReadinessStatePartial
		}
		switch status {
		case ProductionUsabilityReadinessWarning:
			hasWarning = true
		case ProductionUsabilityReadinessFail, ProductionUsabilityReadinessDegraded, ProductionUsabilityReadinessUnsupported, ProductionUsabilityReadinessNotRun:
			return ProductionUsabilityValCReadinessStatePartial
		}
		if item.Blocking && status != ProductionUsabilityReadinessPass {
			return ProductionUsabilityValCReadinessStatePartial
		}
	}
	if !hasWarning {
		return ProductionUsabilityValCReadinessStatePartial
	}
	return ProductionUsabilityValCReadinessStateActive
}

func EvaluateProductionUsabilityValCGuidedReadinessState(model GuidedReadinessBaseline) string {
	if strings.TrimSpace(model.Mode) == "" || len(model.RequiredSteps) == 0 || strings.TrimSpace(model.LimitationMessage) == "" || len(model.NextStepGuidance) == 0 {
		return ProductionUsabilityValCGuidedReadinessStateIncomplete
	}
	if !containsTrimmedString(productionUsabilityValCGuidedModes(), model.Mode) ||
		model.MutatesCanonicalState ||
		model.ClaimsProductionCompletion {
		return ProductionUsabilityValCGuidedReadinessStatePartial
	}
	if model.SampleConfigDetected && model.AutoProductionEnablement {
		return ProductionUsabilityValCGuidedReadinessStatePartial
	}
	if model.FakeDemoEvidenceDetected {
		return ProductionUsabilityValCGuidedReadinessStatePartial
	}
	if len(model.BlockingSteps) > 0 || len(model.MissingSteps) > 0 {
		if model.GoLiveAllowed {
			return ProductionUsabilityValCGuidedReadinessStatePartial
		}
	}
	if model.Mode == ProductionUsabilityGuidedModeGoLive && !model.ProductionConfigDetected {
		return ProductionUsabilityValCGuidedReadinessStatePartial
	}
	return ProductionUsabilityValCGuidedReadinessStateActive
}

func EvaluateProductionUsabilityValCSupportBundleState(model SupportBundleQualityGate) string {
	if strings.TrimSpace(model.BundleID) == "" || strings.TrimSpace(model.GeneratedAtIndicator) == "" || strings.TrimSpace(model.RequestedByActorRef) == "" || strings.TrimSpace(model.VisibilityScope) == "" || strings.TrimSpace(model.RedactionTier) == "" || len(model.IncludedSections) == 0 || strings.TrimSpace(model.SecretScanResult) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ProductionUsabilityValCSupportBundleStateIncomplete
	}
	if !containsTrimmedString(productionUsabilityValCSecretScanStatuses(), model.SecretScanResult) ||
		!containsTrimmedString(productionUsabilityValAExplainScopes(), model.VisibilityScope) ||
		!containsTrimmedString(ProductionUsabilityVal0ExplainabilityContract().SupportedRedactionTiers, model.RedactionTier) ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") ||
		!model.ManifestPresent ||
		model.RawSecretDetected ||
		model.RawTokenDetected ||
		model.UnfilteredEnvDetected ||
		model.CacheClaimsCanonicalTruth {
		return ProductionUsabilityValCSupportBundleStatePartial
	}
	if len(model.ExcludedSections) != len(model.ExclusionReasons) {
		return ProductionUsabilityValCSupportBundleStatePartial
	}
	if len(model.EvidenceRefsRedacted) > 0 && !model.RedactedEvidenceRepresented {
		return ProductionUsabilityValCSupportBundleStatePartial
	}
	if model.VisibilityScope == ProductionUsabilityVisibilityPublicSafe && model.RedactionTier != ProductionUsabilityRedactionPublicSafe {
		return ProductionUsabilityValCSupportBundleStatePartial
	}
	if model.VisibilityScope == ProductionUsabilityVisibilityPartner && model.RedactionTier == ProductionUsabilityRedactionNone {
		return ProductionUsabilityValCSupportBundleStatePartial
	}
	return ProductionUsabilityValCSupportBundleStateActive
}

func EvaluateProductionUsabilityValCDiagnosticsState(model DiagnosticsHardeningModel) string {
	if strings.TrimSpace(model.DiagnosticsID) == "" || strings.TrimSpace(model.GeneratedAtIndicator) == "" || len(model.Sections) == 0 || strings.TrimSpace(model.PermissionScope) == "" || strings.TrimSpace(model.RedactionTier) == "" || strings.TrimSpace(model.SecretScanStatus) == "" || strings.TrimSpace(model.LimitationMessage) == "" || strings.TrimSpace(model.RecoveryHint) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ProductionUsabilityValCDiagnosticsStateIncomplete
	}
	if !containsTrimmedString(productionUsabilityValCSecretScanStatuses(), model.SecretScanStatus) ||
		!containsTrimmedString(productionUsabilityValAExplainScopes(), model.PermissionScope) ||
		!containsTrimmedString(ProductionUsabilityVal0ExplainabilityContract().SupportedRedactionTiers, model.RedactionTier) ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValCDiagnosticsStatePartial
	}
	if model.SafeToShare && (!model.SensitiveFieldsRedacted || model.SecretScanStatus != ProductionUsabilitySecretScanPassed) {
		return ProductionUsabilityValCDiagnosticsStatePartial
	}
	if len(model.UnsupportedSections) > 0 && !model.UnsupportedSectionsExplicit {
		return ProductionUsabilityValCDiagnosticsStatePartial
	}
	if len(model.StaleSections) > 0 && !model.StaleSectionsExplicit {
		return ProductionUsabilityValCDiagnosticsStatePartial
	}
	if len(model.PartialSections) > 0 && !model.PartialSectionsExplicit {
		return ProductionUsabilityValCDiagnosticsStatePartial
	}
	return ProductionUsabilityValCDiagnosticsStateActive
}

func EvaluateProductionUsabilityValCHealthSnapshotState(model HealthSnapshotModel) string {
	if strings.TrimSpace(model.SnapshotID) == "" || strings.TrimSpace(model.GeneratedAtIndicator) == "" || strings.TrimSpace(model.HealthState) == "" || len(model.ComponentStates) == 0 || strings.TrimSpace(model.FreshnessState) == "" || strings.TrimSpace(model.LimitationMessage) == "" || strings.TrimSpace(model.RecoveryHint) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ProductionUsabilityValCHealthSnapshotStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedHealthStates, productionUsabilityValCHealthStates()...) ||
		!containsTrimmedString(model.SupportedHealthStates, model.HealthState) ||
		!containsTrimmedString([]string{ProductionUsabilityStatusFresh, ProductionUsabilityStatusStale, ProductionUsabilityStatusPartial, ProductionUsabilityStatusDegraded, ProductionUsabilityStatusUnavailable, ProductionUsabilityStatusUnsupported}, model.FreshnessState) ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValCHealthSnapshotStatePartial
	}
	for component, state := range model.ComponentStates {
		if strings.TrimSpace(component) == "" || !containsTrimmedString(model.SupportedHealthStates, state) {
			return ProductionUsabilityValCHealthSnapshotStatePartial
		}
		switch state {
		case ProductionUsabilityHealthDegraded, ProductionUsabilityHealthUnhealthy:
			if !containsTrimmedString(model.DegradedComponents, component) {
				return ProductionUsabilityValCHealthSnapshotStatePartial
			}
		case ProductionUsabilityHealthUnavailable:
			if !containsTrimmedString(model.UnavailableComponents, component) {
				return ProductionUsabilityValCHealthSnapshotStatePartial
			}
		case ProductionUsabilityHealthUnsupported:
			if !containsTrimmedString(model.UnsupportedComponents, component) {
				return ProductionUsabilityValCHealthSnapshotStatePartial
			}
		}
	}
	if model.HealthState == ProductionUsabilityHealthHealthy && (len(model.DegradedComponents) > 0 || len(model.UnavailableComponents) > 0 || len(model.UnsupportedComponents) > 0) {
		return ProductionUsabilityValCHealthSnapshotStatePartial
	}
	if model.FreshnessState == ProductionUsabilityStatusFresh && len(model.StaleComponents) > 0 {
		return ProductionUsabilityValCHealthSnapshotStatePartial
	}
	return ProductionUsabilityValCHealthSnapshotStateActive
}

func EvaluateProductionUsabilityValCRecoveryPlaybookState(model RecoveryPlaybookModel) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValCRecoveryPlaybookStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedScenarios, productionUsabilityValCRecoveryScenarios()...) {
		return ProductionUsabilityValCRecoveryPlaybookStatePartial
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		if strings.TrimSpace(item.PlaybookID) == "" || strings.TrimSpace(item.Scenario) == "" || strings.TrimSpace(item.RemediationClass) == "" || len(item.SafeSteps) == 0 || len(item.UnsafeSteps) == 0 || strings.TrimSpace(item.DoNotRetryReason) == "" || strings.TrimSpace(item.EscalationPath) == "" || strings.TrimSpace(item.RollbackHint) == "" || strings.TrimSpace(item.InspectCommandRef) == "" || strings.TrimSpace(item.ExplainCommandRef) == "" || len(item.RequiredPermissions) == 0 || strings.TrimSpace(item.LimitationMessage) == "" {
			return ProductionUsabilityValCRecoveryPlaybookStateIncomplete
		}
		if !containsTrimmedString(model.SupportedScenarios, item.Scenario) || !containsTrimmedString(model.SupportedRemediations, item.RemediationClass) {
			return ProductionUsabilityValCRecoveryPlaybookStatePartial
		}
		if _, duplicate := seen[item.Scenario]; duplicate {
			return ProductionUsabilityValCRecoveryPlaybookStatePartial
		}
		seen[item.Scenario] = struct{}{}
		if item.PolicyBypassSuggested || item.UnsafeRetrySuggested || hasTrimmedStringOverlap(item.SafeSteps, item.UnsafeSteps) {
			return ProductionUsabilityValCRecoveryPlaybookStatePartial
		}
	}
	if len(seen) != len(model.SupportedScenarios) {
		return ProductionUsabilityValCRecoveryPlaybookStatePartial
	}
	return ProductionUsabilityValCRecoveryPlaybookStateActive
}

func EvaluateProductionUsabilityValCUpgradeAdvisoryState(model UpgradeRollbackAdvisory) string {
	if strings.TrimSpace(model.AdvisoryID) == "" || strings.TrimSpace(model.CurrentVersion) == "" || strings.TrimSpace(model.TargetVersion) == "" || strings.TrimSpace(model.CompatibilityStatus) == "" || strings.TrimSpace(model.RollbackScope) == "" || strings.TrimSpace(model.AdvisoryMode) == "" || strings.TrimSpace(model.GeneratedAtIndicator) == "" || strings.TrimSpace(model.LimitationDisclaimer) == "" {
		return ProductionUsabilityValCUpgradeAdvisoryStateIncomplete
	}
	if !containsTrimmedString(model.KnownTargetVersions, model.TargetVersion) ||
		!containsTrimmedString(productionUsabilityValCAdvisoryModes(), model.AdvisoryMode) ||
		!containsTrimmedString(ProductionUsabilityVal0ConfigIntegrity().SupportedCompatibility, model.CompatibilityStatus) ||
		model.MutatesState ||
		!strings.Contains(strings.TrimSpace(model.LimitationDisclaimer), "projection_only") {
		return ProductionUsabilityValCUpgradeAdvisoryStatePartial
	}
	if model.RollbackAvailable && (!strings.Contains(strings.TrimSpace(model.RollbackScope), "config_policy") || len(model.RollbackLimitations) == 0) {
		return ProductionUsabilityValCUpgradeAdvisoryStatePartial
	}
	if (model.AdvisoryMode == ProductionUsabilityAdvisoryPreview || model.AdvisoryMode == ProductionUsabilityAdvisoryAuditOnly) && model.ApprovalImplied {
		return ProductionUsabilityValCUpgradeAdvisoryStatePartial
	}
	if model.CompatibilityStatus == ProductionUsabilityCompatibilityUnknown && model.AdvisoryMode != ProductionUsabilityAdvisoryUnsupported {
		return ProductionUsabilityValCUpgradeAdvisoryStatePartial
	}
	return ProductionUsabilityValCUpgradeAdvisoryStateActive
}

func EvaluateProductionUsabilityValCPermissionSupportState(model PermissionAwareSupportFlowModel) string {
	if len(model.Items) == 0 {
		return ProductionUsabilityValCPermissionSupportStateIncomplete
	}
	if !containsExactTrimmedStringSet(model.SupportedEvidenceVisibility, ProductionUsabilityEvidenceFull, ProductionUsabilityEvidenceMetadataOnly, ProductionUsabilityEvidenceRedacted, ProductionUsabilityEvidenceHidden) ||
		!containsExactTrimmedStringSet(model.SupportedActionModes, productionUsabilityValCSupportActionModes()...) {
		return ProductionUsabilityValCPermissionSupportStatePartial
	}
	expectedScopes := map[string]struct{}{}
	for _, scope := range productionUsabilityValAExplainScopes() {
		expectedScopes[scope] = struct{}{}
	}
	seen := map[string]struct{}{}
	for _, item := range model.Items {
		scope := strings.TrimSpace(item.VisibilityScope)
		if strings.TrimSpace(item.RequesterRole) == "" || scope == "" || strings.TrimSpace(item.CurrentState) == "" || len(item.AllowedSections) == 0 || strings.TrimSpace(item.EvidenceVisibility) == "" || strings.TrimSpace(item.SafeFallbackMessage) == "" || strings.TrimSpace(item.SupportActionMode) == "" {
			return ProductionUsabilityValCPermissionSupportStateIncomplete
		}
		if _, ok := expectedScopes[scope]; !ok {
			return ProductionUsabilityValCPermissionSupportStatePartial
		}
		if !containsTrimmedString(model.SupportedEvidenceVisibility, item.EvidenceVisibility) ||
			!containsTrimmedString(model.SupportedActionModes, item.SupportActionMode) ||
			item.MutatesCanonicalState {
			return ProductionUsabilityValCPermissionSupportStatePartial
		}
		if _, duplicate := seen[scope]; duplicate {
			return ProductionUsabilityValCPermissionSupportStatePartial
		}
		seen[scope] = struct{}{}
		if (scope == ProductionUsabilityVisibilityPartner || scope == ProductionUsabilityVisibilityPublicSafe) && (item.EvidenceVisibility == ProductionUsabilityEvidenceFull || len(item.HiddenSections) == 0) {
			return ProductionUsabilityValCPermissionSupportStatePartial
		}
	}
	if len(seen) != len(expectedScopes) {
		return ProductionUsabilityValCPermissionSupportStatePartial
	}
	return ProductionUsabilityValCPermissionSupportStateActive
}

func EvaluateProductionUsabilityValCExportSafetyState(model RedactionSafeExportModel) string {
	if strings.TrimSpace(model.ExportID) == "" || strings.TrimSpace(model.ExportScope) == "" || strings.TrimSpace(model.RedactionTier) == "" || strings.TrimSpace(model.SecretScanStatus) == "" || len(model.AllowedContentClasses) == 0 || len(model.BlockedContentClasses) == 0 || strings.TrimSpace(model.EvidenceHandling) == "" || strings.TrimSpace(model.LimitationMessage) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return ProductionUsabilityValCExportSafetyStateIncomplete
	}
	if !containsTrimmedString(productionUsabilityValCSecretScanStatuses(), model.SecretScanStatus) ||
		!containsTrimmedString(ProductionUsabilityVal0ExplainabilityContract().SupportedRedactionTiers, model.RedactionTier) ||
		!containsTrimmedString([]string{ProductionUsabilityEvidenceFull, ProductionUsabilityEvidenceMetadataOnly, ProductionUsabilityEvidenceRedacted, ProductionUsabilityEvidenceHidden}, model.EvidenceHandling) ||
		!strings.Contains(strings.TrimSpace(model.ProjectionDisclaimer), "not_canonical_truth") {
		return ProductionUsabilityValCExportSafetyStatePartial
	}
	if model.RawSecretDetected || !model.PolicyAllowsExport || model.RawInternalEvidenceIncluded {
		return ProductionUsabilityValCExportSafetyStatePartial
	}
	if (model.PublicSafe || model.PartnerSafe) && model.EvidenceHandling == ProductionUsabilityEvidenceFull {
		return ProductionUsabilityValCExportSafetyStatePartial
	}
	if model.AuditorSafe && model.PublicSafe {
		return ProductionUsabilityValCExportSafetyStatePartial
	}
	return ProductionUsabilityValCExportSafetyStateActive
}

func EvaluateProductionUsabilityValCState(val0State, valAState, valBState, readinessState, guidedReadinessState, supportBundleState, diagnosticsState, healthSnapshotState, recoveryPlaybookState, upgradeAdvisoryState, permissionSupportState, exportSafetyState string) string {
	if strings.TrimSpace(val0State) != ProductionUsabilityVal0StateActive ||
		strings.TrimSpace(valAState) != ProductionUsabilityValAStateActive ||
		strings.TrimSpace(valBState) != ProductionUsabilityValBStateActive {
		return ProductionUsabilityValCStateIncomplete
	}
	hasPartial := false
	for _, state := range []string{
		strings.TrimSpace(readinessState),
		strings.TrimSpace(guidedReadinessState),
		strings.TrimSpace(supportBundleState),
		strings.TrimSpace(diagnosticsState),
		strings.TrimSpace(healthSnapshotState),
		strings.TrimSpace(recoveryPlaybookState),
		strings.TrimSpace(upgradeAdvisoryState),
		strings.TrimSpace(permissionSupportState),
		strings.TrimSpace(exportSafetyState),
	} {
		switch state {
		case ProductionUsabilityValCReadinessStateActive,
			ProductionUsabilityValCGuidedReadinessStateActive,
			ProductionUsabilityValCSupportBundleStateActive,
			ProductionUsabilityValCDiagnosticsStateActive,
			ProductionUsabilityValCHealthSnapshotStateActive,
			ProductionUsabilityValCRecoveryPlaybookStateActive,
			ProductionUsabilityValCUpgradeAdvisoryStateActive,
			ProductionUsabilityValCPermissionSupportStateActive,
			ProductionUsabilityValCExportSafetyStateActive:
		case ProductionUsabilityValCReadinessStatePartial,
			ProductionUsabilityValCGuidedReadinessStatePartial,
			ProductionUsabilityValCSupportBundleStatePartial,
			ProductionUsabilityValCDiagnosticsStatePartial,
			ProductionUsabilityValCHealthSnapshotStatePartial,
			ProductionUsabilityValCRecoveryPlaybookStatePartial,
			ProductionUsabilityValCUpgradeAdvisoryStatePartial,
			ProductionUsabilityValCPermissionSupportStatePartial,
			ProductionUsabilityValCExportSafetyStatePartial:
			hasPartial = true
		default:
			return ProductionUsabilityValCStateIncomplete
		}
	}
	if hasPartial {
		return ProductionUsabilityValCStateSubstantial
	}
	return ProductionUsabilityValCStateActive
}

func EvaluateProductionUsabilityValCProofsState(val0State, valAState, valBState, readinessState, guidedReadinessState, supportBundleState, diagnosticsState, healthSnapshotState, recoveryPlaybookState, upgradeAdvisoryState, permissionSupportState, exportSafetyState string, surfaceRefs, evidenceRefs, limitations, whyPoint4NotPass []string) string {
	baseState := EvaluateProductionUsabilityValCState(val0State, valAState, valBState, readinessState, guidedReadinessState, supportBundleState, diagnosticsState, healthSnapshotState, recoveryPlaybookState, upgradeAdvisoryState, permissionSupportState, exportSafetyState)
	if len(surfaceRefs) < 11 || len(evidenceRefs) < 9 || len(limitations) == 0 || len(whyPoint4NotPass) == 0 {
		if baseState == ProductionUsabilityValCStateActive {
			return ProductionUsabilityValCStateSubstantial
		}
		return baseState
	}
	return baseState
}

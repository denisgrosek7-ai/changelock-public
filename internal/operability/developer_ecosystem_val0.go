package operability

import "strings"

const (
	DeveloperEcosystemPoint8StateNotComplete            = "developer_ecosystem_point_8_not_complete"
	DeveloperEcosystemVal0PerformanceBudgetDisciplineID = "developer-ecosystem-performance-budget"

	DeveloperEcosystemVal0DependencyStateActive     = "developer_ecosystem_val0_dependency_active"
	DeveloperEcosystemVal0DependencyStatePartial    = "developer_ecosystem_val0_dependency_partial"
	DeveloperEcosystemVal0DependencyStateIncomplete = "developer_ecosystem_val0_dependency_incomplete"
	DeveloperEcosystemVal0DependencyStateBlocked    = "developer_ecosystem_val0_dependency_blocked"
	DeveloperEcosystemVal0DependencyStateUnknown    = "developer_ecosystem_val0_dependency_unknown"

	DeveloperEcosystemVal0OutputClassificationStateActive     = "developer_ecosystem_val0_output_classification_active"
	DeveloperEcosystemVal0OutputClassificationStatePartial    = "developer_ecosystem_val0_output_classification_partial"
	DeveloperEcosystemVal0OutputClassificationStateIncomplete = "developer_ecosystem_val0_output_classification_incomplete"
	DeveloperEcosystemVal0OutputClassificationStateBlocked    = "developer_ecosystem_val0_output_classification_blocked"
	DeveloperEcosystemVal0OutputClassificationStateUnknown    = "developer_ecosystem_val0_output_classification_unknown"

	DeveloperEcosystemVal0IDEAdvisoryStateActive     = "developer_ecosystem_val0_ide_advisory_active"
	DeveloperEcosystemVal0IDEAdvisoryStatePartial    = "developer_ecosystem_val0_ide_advisory_partial"
	DeveloperEcosystemVal0IDEAdvisoryStateIncomplete = "developer_ecosystem_val0_ide_advisory_incomplete"
	DeveloperEcosystemVal0IDEAdvisoryStateBlocked    = "developer_ecosystem_val0_ide_advisory_blocked"
	DeveloperEcosystemVal0IDEAdvisoryStateUnknown    = "developer_ecosystem_val0_ide_advisory_unknown"

	DeveloperEcosystemVal0LocalProductionStateActive     = "developer_ecosystem_val0_local_production_active"
	DeveloperEcosystemVal0LocalProductionStatePartial    = "developer_ecosystem_val0_local_production_partial"
	DeveloperEcosystemVal0LocalProductionStateIncomplete = "developer_ecosystem_val0_local_production_incomplete"
	DeveloperEcosystemVal0LocalProductionStateBlocked    = "developer_ecosystem_val0_local_production_blocked"
	DeveloperEcosystemVal0LocalProductionStateUnknown    = "developer_ecosystem_val0_local_production_unknown"

	DeveloperEcosystemVal0RepoPolicyStateActive     = "developer_ecosystem_val0_repo_policy_active"
	DeveloperEcosystemVal0RepoPolicyStatePartial    = "developer_ecosystem_val0_repo_policy_partial"
	DeveloperEcosystemVal0RepoPolicyStateIncomplete = "developer_ecosystem_val0_repo_policy_incomplete"
	DeveloperEcosystemVal0RepoPolicyStateBlocked    = "developer_ecosystem_val0_repo_policy_blocked"
	DeveloperEcosystemVal0RepoPolicyStateUnknown    = "developer_ecosystem_val0_repo_policy_unknown"

	DeveloperEcosystemVal0PluginSafetyStateActive     = "developer_ecosystem_val0_plugin_safety_active"
	DeveloperEcosystemVal0PluginSafetyStatePartial    = "developer_ecosystem_val0_plugin_safety_partial"
	DeveloperEcosystemVal0PluginSafetyStateIncomplete = "developer_ecosystem_val0_plugin_safety_incomplete"
	DeveloperEcosystemVal0PluginSafetyStateBlocked    = "developer_ecosystem_val0_plugin_safety_blocked"
	DeveloperEcosystemVal0PluginSafetyStateUnknown    = "developer_ecosystem_val0_plugin_safety_unknown"

	DeveloperEcosystemVal0PerformanceBudgetStateActive     = "developer_ecosystem_val0_performance_budget_active"
	DeveloperEcosystemVal0PerformanceBudgetStatePartial    = "developer_ecosystem_val0_performance_budget_partial"
	DeveloperEcosystemVal0PerformanceBudgetStateIncomplete = "developer_ecosystem_val0_performance_budget_incomplete"
	DeveloperEcosystemVal0PerformanceBudgetStateBlocked    = "developer_ecosystem_val0_performance_budget_blocked"
	DeveloperEcosystemVal0PerformanceBudgetStateUnknown    = "developer_ecosystem_val0_performance_budget_unknown"

	DeveloperEcosystemVal0DXMetricsStateActive     = "developer_ecosystem_val0_dx_metrics_active"
	DeveloperEcosystemVal0DXMetricsStatePartial    = "developer_ecosystem_val0_dx_metrics_partial"
	DeveloperEcosystemVal0DXMetricsStateIncomplete = "developer_ecosystem_val0_dx_metrics_incomplete"
	DeveloperEcosystemVal0DXMetricsStateBlocked    = "developer_ecosystem_val0_dx_metrics_blocked"
	DeveloperEcosystemVal0DXMetricsStateUnknown    = "developer_ecosystem_val0_dx_metrics_unknown"

	DeveloperEcosystemVal0NoOverclaimStateActive     = "developer_ecosystem_val0_no_overclaim_active"
	DeveloperEcosystemVal0NoOverclaimStatePartial    = "developer_ecosystem_val0_no_overclaim_partial"
	DeveloperEcosystemVal0NoOverclaimStateIncomplete = "developer_ecosystem_val0_no_overclaim_incomplete"
	DeveloperEcosystemVal0NoOverclaimStateBlocked    = "developer_ecosystem_val0_no_overclaim_blocked"
	DeveloperEcosystemVal0NoOverclaimStateUnknown    = "developer_ecosystem_val0_no_overclaim_unknown"

	DeveloperEcosystemVal0StateActive     = "developer_ecosystem_val0_active"
	DeveloperEcosystemVal0StatePartial    = "developer_ecosystem_val0_partial"
	DeveloperEcosystemVal0StateIncomplete = "developer_ecosystem_val0_incomplete"
	DeveloperEcosystemVal0StateBlocked    = "developer_ecosystem_val0_blocked"
	DeveloperEcosystemVal0StateUnknown    = "developer_ecosystem_val0_unknown"

	DeveloperEcosystemOutputObservedFact          = "observed_fact"
	DeveloperEcosystemOutputDerivedAdvisory       = "derived_advisory_signal"
	DeveloperEcosystemOutputRecommendation        = "recommendation"
	DeveloperEcosystemOutputRemediationHint       = "remediation_hint"
	DeveloperEcosystemOutputUncertainty           = "uncertainty"
	DeveloperEcosystemOutputStaleOrPartial        = "stale_or_partial_state"
	DeveloperEcosystemOutputProductionOnlyUnknown = "production_only_unknown"

	DeveloperEcosystemSurfaceIDE    = "ide"
	DeveloperEcosystemSurfaceLocal  = "local"
	DeveloperEcosystemSurfaceSDK    = "sdk"
	DeveloperEcosystemSurfacePlugin = "plugin"

	DeveloperEcosystemSimulationLocalValidation  = "local_validation"
	DeveloperEcosystemSimulationMockVerification = "mock_verification"
	DeveloperEcosystemSimulationOfflineBundle    = "offline_bundle"

	DeveloperEcosystemRepoConfigSchemaV0         = "developer_ecosystem_repo_config.v0"
	DeveloperEcosystemRepoScopeLocalAdvisoryOnly = "repo_local_advisory_only"
	DeveloperEcosystemEnterpriseOverrideGoverned = "enterprise_governance_required"
	DeveloperEcosystemFailClosedHandling         = "fail_closed"
	DeveloperEcosystemDeprecationVisible         = "visible_with_warning"

	DeveloperEcosystemPluginCapabilityReadWorkspace      = "read_workspace"
	DeveloperEcosystemPluginCapabilityEmitAdvisorySignal = "emit_advisory_signal"
	DeveloperEcosystemPluginCapabilityCollectDebugTrace  = "collect_debug_trace"
	DeveloperEcosystemPluginCapabilityRunLocalValidation = "run_local_validation"
	DeveloperEcosystemPluginScopeAdvisoryOnly            = "advisory_only"

	DeveloperEcosystemBudgetClassIDELatency      = "ide_latency"
	DeveloperEcosystemBudgetClassLocalValidation = "local_validation_time"
	DeveloperEcosystemBudgetClassPreCommit       = "pre_commit_pre_push_blocking_time"
	DeveloperEcosystemBudgetClassPluginExecution = "plugin_execution"
	DeveloperEcosystemBudgetClassDegradedMode    = "degraded_mode_fallback"

	DeveloperEcosystemBudgetStateKnown       = "known"
	DeveloperEcosystemBudgetStateDegraded    = "degraded"
	DeveloperEcosystemBudgetStateUnknown     = "unknown"
	DeveloperEcosystemBudgetStateUnsupported = "unsupported"
)

type DeveloperEcosystemVal0DependencySnapshot struct {
	Point6State                string `json:"point_6_state"`
	Point7State                string `json:"point_7_state"`
	Point7ClosureState         string `json:"point_7_closure_state"`
	Point7PrerequisiteState    string `json:"point_7_prerequisite_state"`
	Point7InvariantState       string `json:"point_7_invariant_state"`
	Point7ProofSurfaceState    string `json:"point_7_proof_surface_state"`
	Point7EvidenceQualityState string `json:"point_7_evidence_quality_state"`
	Point7NoOverclaimState     string `json:"point_7_no_overclaim_state"`
	Point7PassRuleState        string `json:"point_7_pass_rule_state"`
	Point7PassAllowed          bool   `json:"point_7_pass_allowed"`
}

type DeveloperEcosystemVal0OutputClassification struct {
	CurrentState                    string   `json:"current_state"`
	ClassificationID                string   `json:"classification_id"`
	Version                         string   `json:"version"`
	ClassifiedSurfaceKinds          []string `json:"classified_surface_kinds,omitempty"`
	AllowedOutputClasses            []string `json:"allowed_output_classes,omitempty"`
	UncertaintyVisible              bool     `json:"uncertainty_visible"`
	StalePartialVisible             bool     `json:"stale_partial_visible"`
	ProductionOnlyUnknownVisible    bool     `json:"production_only_unknown_visible"`
	AdvisoryTreatedAsApproval       bool     `json:"advisory_treated_as_approval"`
	RecommendationsSuppressFailures bool     `json:"recommendations_suppress_failures"`
	RedactionConvertsFailureToPass  bool     `json:"redaction_converts_failure_to_pass"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
	CreatedAt                       string   `json:"created_at"`
	UpdatedAt                       string   `json:"updated_at"`
}

type DeveloperEcosystemVal0IDEAdvisoryDiscipline struct {
	CurrentState             string `json:"current_state"`
	DisciplineID             string `json:"discipline_id"`
	Version                  string `json:"version"`
	AdvisoryOnly             bool   `json:"advisory_only"`
	EvidenceLinked           bool   `json:"evidence_linked"`
	FreshnessAware           bool   `json:"freshness_aware"`
	ReasonCoded              bool   `json:"reason_coded"`
	UncertaintyAware         bool   `json:"uncertainty_aware"`
	CandidateVsReviewedAware bool   `json:"candidate_vs_reviewed_aware"`
	NonMutating              bool   `json:"non_mutating"`
	NonApproving             bool   `json:"non_approving"`
	CanonicalTruthClaim      bool   `json:"canonical_truth_claim"`
	ProductionApprovalClaim  bool   `json:"production_approval_claim"`
	PolicyOverrideClaim      bool   `json:"policy_override_claim"`
	DeploymentApprovalClaim  bool   `json:"deployment_approval_claim"`
	CertificationClaim       bool   `json:"certification_claim"`
	ProjectionDisclaimer     string `json:"projection_disclaimer"`
}

type DeveloperEcosystemVal0LocalProductionDiscipline struct {
	CurrentState                    string `json:"current_state"`
	DisciplineID                    string `json:"discipline_id"`
	Version                         string `json:"version"`
	SimulationScope                 string `json:"simulation_scope"`
	UnsupportedCaseDisclosure       bool   `json:"unsupported_case_disclosure"`
	ProductionOnlyUnknownDisclosure bool   `json:"production_only_unknown_disclosure"`
	FreshnessContextAssumptions     bool   `json:"freshness_context_assumptions"`
	ProductionEquivalenceClaim      bool   `json:"production_equivalence_claim"`
	MutatesCanonicalEvidence        bool   `json:"mutates_canonical_evidence"`
	ApprovalAuthority               bool   `json:"approval_authority"`
	SuppressesFailures              bool   `json:"suppresses_failures"`
	ProjectionDisclaimer            string `json:"projection_disclaimer"`
}

type DeveloperEcosystemVal0RepoPolicyBoundaryDiscipline struct {
	CurrentState                    string   `json:"current_state"`
	DisciplineID                    string   `json:"discipline_id"`
	Version                         string   `json:"version"`
	RepoConfigSchemaVersion         string   `json:"repo_config_schema_version"`
	SupportedSchemaVersions         []string `json:"supported_schema_versions,omitempty"`
	AllowedFields                   []string `json:"allowed_fields,omitempty"`
	ScopeBoundary                   string   `json:"scope_boundary"`
	EnterpriseOverrideBoundary      string   `json:"enterprise_override_boundary"`
	ReviewRequired                  bool     `json:"review_required"`
	UnknownFieldHandling            string   `json:"unknown_field_handling"`
	UnsupportedVersionHandling      string   `json:"unsupported_version_handling"`
	DeprecationBehavior             string   `json:"deprecation_behavior"`
	OverrideEnterprisePolicy        bool     `json:"override_enterprise_policy"`
	DisableCanonicalEvidenceRules   bool     `json:"disable_canonical_evidence_rules"`
	GrantApprovalAuthority          bool     `json:"grant_approval_authority"`
	ChangeProductionTrustWithoutGov bool     `json:"change_production_trust_without_governance"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemVal0PluginSafetyDiscipline struct {
	CurrentState             string   `json:"current_state"`
	DisciplineID             string   `json:"discipline_id"`
	Version                  string   `json:"version"`
	DeclaredCapabilities     []string `json:"declared_capabilities,omitempty"`
	BoundedScope             string   `json:"bounded_scope"`
	SandboxIsolationExpected bool     `json:"sandbox_isolation_expected"`
	PerformanceBudgetRef     string   `json:"performance_budget_ref"`
	AuditVisibility          bool     `json:"audit_visibility"`
	DebugVisibility          bool     `json:"debug_visibility"`
	FailureStateVisible      bool     `json:"failure_state_visible"`
	HiddenCanonicalMutation  bool     `json:"hidden_canonical_mutation"`
	HiddenPolicyOverride     bool     `json:"hidden_policy_override"`
	HiddenApprovalPath       bool     `json:"hidden_approval_path"`
	GovernanceBypass         bool     `json:"governance_bypass"`
	CanonicalTruthClaim      bool     `json:"canonical_truth_claim"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemVal0PerformanceBudgetClass struct {
	BudgetClass             string `json:"budget_class"`
	BudgetState             string `json:"budget_state"`
	TargetMillis            int    `json:"target_millis"`
	DegradedFallbackVisible bool   `json:"degraded_fallback_visible"`
}

type DeveloperEcosystemVal0PerformanceBudgetDiscipline struct {
	CurrentState          string                                         `json:"current_state"`
	DisciplineID          string                                         `json:"discipline_id"`
	Version               string                                         `json:"version"`
	Budgets               []DeveloperEcosystemVal0PerformanceBudgetClass `json:"budgets,omitempty"`
	DegradedModeVisible   bool                                           `json:"degraded_mode_visible"`
	SilentBypassAllowed   bool                                           `json:"silent_bypass_allowed"`
	HidesFailuresDegraded bool                                           `json:"hides_failures_degraded"`
	ProjectionDisclaimer  string                                         `json:"projection_disclaimer"`
}

type DeveloperEcosystemVal0DXMetricsDiscipline struct {
	CurrentState         string   `json:"current_state"`
	DisciplineID         string   `json:"discipline_id"`
	Version              string   `json:"version"`
	MetricNames          []string `json:"metric_names,omitempty"`
	AdvisoryOnly         bool     `json:"advisory_only"`
	DeveloperTrustScore  bool     `json:"developer_trust_score"`
	CertificationUse     bool     `json:"certification_use"`
	FastTrackApproval    bool     `json:"fast_track_approval"`
	DeploymentAuthority  bool     `json:"deployment_authority"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemVal0NoOverclaimDiscipline struct {
	CurrentState             string `json:"current_state"`
	DisciplineID             string `json:"discipline_id"`
	Version                  string `json:"version"`
	CanonicalTruthClaim      bool   `json:"canonical_truth_claim"`
	ProductionApprovalClaim  bool   `json:"production_approval_claim"`
	DeveloperTrustScoreClaim bool   `json:"developer_trust_score_claim"`
	CertificationClaim       bool   `json:"certification_claim"`
	FastTrackClaim           bool   `json:"fast_track_claim"`
	EvidenceMutationClaim    bool   `json:"evidence_mutation_claim"`
	ProjectionDisclaimer     string `json:"projection_disclaimer"`
}

type DeveloperEcosystemVal0Foundation struct {
	CurrentState              string                                             `json:"current_state"`
	Point8State               string                                             `json:"point_8_state"`
	DependencyState           string                                             `json:"dependency_state"`
	OutputClassificationState string                                             `json:"output_classification_state"`
	IDEAdvisoryState          string                                             `json:"ide_advisory_state"`
	LocalProductionState      string                                             `json:"local_production_state"`
	RepoPolicyBoundaryState   string                                             `json:"repo_policy_boundary_state"`
	PluginSafetyState         string                                             `json:"plugin_safety_state"`
	PerformanceBudgetState    string                                             `json:"performance_budget_state"`
	DXMetricsState            string                                             `json:"dx_metrics_state"`
	NoOverclaimState          string                                             `json:"no_overclaim_state"`
	FoundationID              string                                             `json:"foundation_id"`
	Version                   string                                             `json:"version"`
	Dependency                DeveloperEcosystemVal0DependencySnapshot           `json:"dependency"`
	OutputClassification      DeveloperEcosystemVal0OutputClassification         `json:"output_classification"`
	IDEAdvisory               DeveloperEcosystemVal0IDEAdvisoryDiscipline        `json:"ide_advisory"`
	LocalProduction           DeveloperEcosystemVal0LocalProductionDiscipline    `json:"local_production"`
	RepoPolicyBoundary        DeveloperEcosystemVal0RepoPolicyBoundaryDiscipline `json:"repo_policy_boundary"`
	PluginSafety              DeveloperEcosystemVal0PluginSafetyDiscipline       `json:"plugin_safety"`
	PerformanceBudget         DeveloperEcosystemVal0PerformanceBudgetDiscipline  `json:"performance_budget"`
	DXMetrics                 DeveloperEcosystemVal0DXMetricsDiscipline          `json:"dx_metrics"`
	NoOverclaim               DeveloperEcosystemVal0NoOverclaimDiscipline        `json:"no_overclaim"`
	EvidenceRefs              []string                                           `json:"evidence_refs,omitempty"`
	ProofSurfaceRefs          []string                                           `json:"proof_surface_refs,omitempty"`
	BlockingReasons           []string                                           `json:"blocking_reasons,omitempty"`
	ProjectionDisclaimer      string                                             `json:"projection_disclaimer"`
	CreatedAt                 string                                             `json:"created_at"`
	UpdatedAt                 string                                             `json:"updated_at"`
}

func developerEcosystemVal0ProjectionDisclaimer() string {
	return "projection_only not_canonical_truth developer_ecosystem_discipline_foundation advisory_projection"
}

func developerEcosystemVal0HasProjectionDisclaimer(value string) bool {
	return strings.Contains(strings.TrimSpace(value), "projection_only") &&
		strings.Contains(strings.TrimSpace(value), "not_canonical_truth")
}

func developerEcosystemVal0OutputClasses() []string {
	return []string{
		DeveloperEcosystemOutputObservedFact,
		DeveloperEcosystemOutputDerivedAdvisory,
		DeveloperEcosystemOutputRecommendation,
		DeveloperEcosystemOutputRemediationHint,
		DeveloperEcosystemOutputUncertainty,
		DeveloperEcosystemOutputStaleOrPartial,
		DeveloperEcosystemOutputProductionOnlyUnknown,
	}
}

func developerEcosystemVal0ClassifiedSurfaces() []string {
	return []string{
		DeveloperEcosystemSurfaceIDE,
		DeveloperEcosystemSurfaceLocal,
		DeveloperEcosystemSurfaceSDK,
		DeveloperEcosystemSurfacePlugin,
	}
}

func developerEcosystemVal0AllowedRepoFields() []string {
	return []string{
		"schema_version",
		"scope_boundary",
		"advisory_mode",
		"compatibility_mode",
		"review_required",
		"deprecation_behavior",
	}
}

func developerEcosystemVal0PluginCapabilities() []string {
	return []string{
		DeveloperEcosystemPluginCapabilityReadWorkspace,
		DeveloperEcosystemPluginCapabilityEmitAdvisorySignal,
		DeveloperEcosystemPluginCapabilityCollectDebugTrace,
		DeveloperEcosystemPluginCapabilityRunLocalValidation,
	}
}

func developerEcosystemVal0BudgetClasses() []string {
	return []string{
		DeveloperEcosystemBudgetClassIDELatency,
		DeveloperEcosystemBudgetClassLocalValidation,
		DeveloperEcosystemBudgetClassPreCommit,
		DeveloperEcosystemBudgetClassPluginExecution,
		DeveloperEcosystemBudgetClassDegradedMode,
	}
}

func developerEcosystemVal0DXMetricNames() []string {
	return []string{
		"ide_adoption",
		"advisory_usefulness",
		"local_validation_completion",
		"pre_commit_bypass_rate",
		"schema_invalid_repo_config_rate",
		"plugin_failure_rate",
		"developer_friction_complaint_rate",
		"mock_vs_production_mismatch_rate",
		"remediation_suggestion_acceptance",
	}
}

func developerEcosystemVal0HasOverclaim(values ...string) bool {
	disallowed := []string{
		"developer trust score",
		"developer certification",
		"certified developer tooling",
		"approved developer tooling",
		"fast-track deployment",
		"fast track deployment",
		"deployment approved",
		"production approved",
		"canonical truth",
		"point_8_pass",
	}
	for _, value := range values {
		lower := strings.ToLower(strings.TrimSpace(value))
		for _, item := range disallowed {
			if strings.Contains(lower, item) {
				return true
			}
		}
	}
	return false
}

func developerEcosystemVal0Evidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:developer-output-classification-001", EvidenceType: "output_classification", Source: "developer-ecosystem/output-classification", Timestamp: "2026-04-28T08:20:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "developer_output_classification", Caveats: []string{"developer output classes remain advisory and fail-closed"}},
		{EvidenceID: "evidence:developer-ide-advisory-001", EvidenceType: "ide_advisory", Source: "developer-ecosystem/ide-advisory", Timestamp: "2026-04-28T08:21:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "ide_advisory_discipline", Caveats: []string{"ide signals remain advisory, reason-coded, freshness-aware, and non-approving"}},
		{EvidenceID: "evidence:developer-local-production-001", EvidenceType: "local_production_boundary", Source: "developer-ecosystem/local-production", Timestamp: "2026-04-28T08:22:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "local_vs_production_discipline", Caveats: []string{"local and mock results remain non-production-equivalent and non-authoritative"}},
		{EvidenceID: "evidence:developer-repo-policy-boundary-001", EvidenceType: "repo_policy_boundary", Source: "developer-ecosystem/repo-policy", Timestamp: "2026-04-28T08:23:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "repo_policy_boundary", Caveats: []string{"repo-local config remains schema-bound and cannot override enterprise governance"}},
		{EvidenceID: "evidence:developer-plugin-safety-001", EvidenceType: "plugin_safety", Source: "developer-ecosystem/plugin-safety", Timestamp: "2026-04-28T08:24:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "plugin_safety_discipline", Caveats: []string{"plugin capabilities remain bounded, auditable, and non-mutating"}},
		{EvidenceID: "evidence:developer-performance-budget-001", EvidenceType: "performance_budget", Source: "developer-ecosystem/performance-budget", Timestamp: "2026-04-28T08:25:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "performance_budget_discipline", Caveats: []string{"developer budget regressions remain visible and cannot silently bypass failures"}},
		{EvidenceID: "evidence:developer-dx-metrics-001", EvidenceType: "dx_metrics", Source: "developer-ecosystem/dx-metrics", Timestamp: "2026-04-28T08:26:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "dx_metrics_discipline", Caveats: []string{"dx metrics remain advisory only and do not create developer trust or approval authority"}},
		{EvidenceID: "evidence:developer-no-overclaim-001", EvidenceType: "no_overclaim", Source: "developer-ecosystem/no-overclaim", Timestamp: "2026-04-28T08:27:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim_discipline", Caveats: []string{"developer tooling cannot claim certification, production approval, or point_8_pass in Val 0"}},
		{EvidenceID: "evidence:developer-canonical-boundary-001", EvidenceType: "canonical_boundary", Source: "developer-ecosystem/canonical-boundary", Timestamp: "2026-04-28T08:28:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "canonical_evidence_boundary", Caveats: []string{"developer tooling remains outside canonical evidence mutation and approval paths"}},
		{EvidenceID: "evidence:point8-governance-001", EvidenceType: "state_governance", Source: "developer-ecosystem/point8-governance", Timestamp: "2026-04-28T08:29:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point8_governance", Caveats: []string{"point_8_state remains not complete and point_8_pass is not allowed in Val 0"}},
	}
}

func developerEcosystemVal0RequiredEvidenceIDs() []string {
	return []string{
		"evidence:developer-output-classification-001",
		"evidence:developer-ide-advisory-001",
		"evidence:developer-local-production-001",
		"evidence:developer-repo-policy-boundary-001",
		"evidence:developer-plugin-safety-001",
		"evidence:developer-performance-budget-001",
		"evidence:developer-dx-metrics-001",
		"evidence:developer-no-overclaim-001",
		"evidence:developer-canonical-boundary-001",
		"evidence:point8-governance-001",
	}
}

func developerEcosystemVal0RequiredEvidenceScopes() []string {
	return []string{
		"developer_output_classification",
		"ide_advisory_discipline",
		"local_vs_production_discipline",
		"repo_policy_boundary",
		"plugin_safety_discipline",
		"performance_budget_discipline",
		"dx_metrics_discipline",
		"no_overclaim_discipline",
		"canonical_evidence_boundary",
		"point8_governance",
	}
}

func DeveloperEcosystemVal0ProofEvidenceRefs() []string {
	return []string{
		"point7_integrated_verifier_ecosystem_closure",
		"developer_ecosystem_discipline_foundation",
		"evidence:developer-output-classification-001",
		"evidence:developer-ide-advisory-001",
		"evidence:developer-local-production-001",
		"evidence:developer-repo-policy-boundary-001",
		"evidence:developer-plugin-safety-001",
		"evidence:developer-performance-budget-001",
		"evidence:developer-dx-metrics-001",
		"evidence:developer-no-overclaim-001",
		"evidence:developer-canonical-boundary-001",
		"evidence:point8-governance-001",
	}
}

func DeveloperEcosystemVal0ProofSurfaceRefs() []string {
	return []string{
		"/v1/verifier-ecosystem/vale/closure",
		"/v1/verifier-ecosystem/vale/proofs",
		"/v1/developer-ecosystem/val0/status",
		"/v1/developer-ecosystem/val0/proofs",
	}
}

func DeveloperEcosystemVal0OutputClassificationModel() DeveloperEcosystemVal0OutputClassification {
	return DeveloperEcosystemVal0OutputClassification{
		ClassificationID:             "developer-ecosystem-output-classification",
		Version:                      "2026.04",
		ClassifiedSurfaceKinds:       developerEcosystemVal0ClassifiedSurfaces(),
		AllowedOutputClasses:         developerEcosystemVal0OutputClasses(),
		UncertaintyVisible:           true,
		StalePartialVisible:          true,
		ProductionOnlyUnknownVisible: true,
		ProjectionDisclaimer:         developerEcosystemVal0ProjectionDisclaimer(),
		CreatedAt:                    "2026-04-28T08:20:00Z",
		UpdatedAt:                    "2026-04-28T08:20:00Z",
	}
}

func DeveloperEcosystemVal0IDEAdvisoryDisciplineModel() DeveloperEcosystemVal0IDEAdvisoryDiscipline {
	return DeveloperEcosystemVal0IDEAdvisoryDiscipline{
		DisciplineID:             "developer-ecosystem-ide-advisory",
		Version:                  "2026.04",
		AdvisoryOnly:             true,
		EvidenceLinked:           true,
		FreshnessAware:           true,
		ReasonCoded:              true,
		UncertaintyAware:         true,
		CandidateVsReviewedAware: true,
		NonMutating:              true,
		NonApproving:             true,
		ProjectionDisclaimer:     developerEcosystemVal0ProjectionDisclaimer(),
	}
}

func DeveloperEcosystemVal0LocalProductionDisciplineModel() DeveloperEcosystemVal0LocalProductionDiscipline {
	return DeveloperEcosystemVal0LocalProductionDiscipline{
		DisciplineID:                    "developer-ecosystem-local-production",
		Version:                         "2026.04",
		SimulationScope:                 DeveloperEcosystemSimulationLocalValidation,
		UnsupportedCaseDisclosure:       true,
		ProductionOnlyUnknownDisclosure: true,
		FreshnessContextAssumptions:     true,
		ProjectionDisclaimer:            developerEcosystemVal0ProjectionDisclaimer(),
	}
}

func DeveloperEcosystemVal0RepoPolicyBoundaryDisciplineModel() DeveloperEcosystemVal0RepoPolicyBoundaryDiscipline {
	return DeveloperEcosystemVal0RepoPolicyBoundaryDiscipline{
		DisciplineID:               "developer-ecosystem-repo-policy-boundary",
		Version:                    "2026.04",
		RepoConfigSchemaVersion:    DeveloperEcosystemRepoConfigSchemaV0,
		SupportedSchemaVersions:    []string{DeveloperEcosystemRepoConfigSchemaV0},
		AllowedFields:              developerEcosystemVal0AllowedRepoFields(),
		ScopeBoundary:              DeveloperEcosystemRepoScopeLocalAdvisoryOnly,
		EnterpriseOverrideBoundary: DeveloperEcosystemEnterpriseOverrideGoverned,
		ReviewRequired:             true,
		UnknownFieldHandling:       DeveloperEcosystemFailClosedHandling,
		UnsupportedVersionHandling: DeveloperEcosystemFailClosedHandling,
		DeprecationBehavior:        DeveloperEcosystemDeprecationVisible,
		ProjectionDisclaimer:       developerEcosystemVal0ProjectionDisclaimer(),
	}
}

func DeveloperEcosystemVal0PluginSafetyDisciplineModel() DeveloperEcosystemVal0PluginSafetyDiscipline {
	return DeveloperEcosystemVal0PluginSafetyDiscipline{
		DisciplineID:             "developer-ecosystem-plugin-safety",
		Version:                  "2026.04",
		DeclaredCapabilities:     developerEcosystemVal0PluginCapabilities(),
		BoundedScope:             DeveloperEcosystemPluginScopeAdvisoryOnly,
		SandboxIsolationExpected: true,
		PerformanceBudgetRef:     DeveloperEcosystemVal0PerformanceBudgetDisciplineID,
		AuditVisibility:          true,
		DebugVisibility:          true,
		FailureStateVisible:      true,
		ProjectionDisclaimer:     developerEcosystemVal0ProjectionDisclaimer(),
	}
}

func DeveloperEcosystemVal0PerformanceBudgetDisciplineModel() DeveloperEcosystemVal0PerformanceBudgetDiscipline {
	return DeveloperEcosystemVal0PerformanceBudgetDiscipline{
		DisciplineID: DeveloperEcosystemVal0PerformanceBudgetDisciplineID,
		Version:      "2026.04",
		Budgets: []DeveloperEcosystemVal0PerformanceBudgetClass{
			{BudgetClass: DeveloperEcosystemBudgetClassIDELatency, BudgetState: DeveloperEcosystemBudgetStateKnown, TargetMillis: 150, DegradedFallbackVisible: true},
			{BudgetClass: DeveloperEcosystemBudgetClassLocalValidation, BudgetState: DeveloperEcosystemBudgetStateKnown, TargetMillis: 2500, DegradedFallbackVisible: true},
			{BudgetClass: DeveloperEcosystemBudgetClassPreCommit, BudgetState: DeveloperEcosystemBudgetStateKnown, TargetMillis: 5000, DegradedFallbackVisible: true},
			{BudgetClass: DeveloperEcosystemBudgetClassPluginExecution, BudgetState: DeveloperEcosystemBudgetStateKnown, TargetMillis: 1000, DegradedFallbackVisible: true},
			{BudgetClass: DeveloperEcosystemBudgetClassDegradedMode, BudgetState: DeveloperEcosystemBudgetStateKnown, TargetMillis: 300, DegradedFallbackVisible: true},
		},
		DegradedModeVisible:  true,
		ProjectionDisclaimer: developerEcosystemVal0ProjectionDisclaimer(),
	}
}

func DeveloperEcosystemVal0DXMetricsDisciplineModel() DeveloperEcosystemVal0DXMetricsDiscipline {
	return DeveloperEcosystemVal0DXMetricsDiscipline{
		DisciplineID:         "developer-ecosystem-dx-metrics",
		Version:              "2026.04",
		MetricNames:          developerEcosystemVal0DXMetricNames(),
		AdvisoryOnly:         true,
		ProjectionDisclaimer: developerEcosystemVal0ProjectionDisclaimer(),
	}
}

func DeveloperEcosystemVal0NoOverclaimDisciplineModel() DeveloperEcosystemVal0NoOverclaimDiscipline {
	return DeveloperEcosystemVal0NoOverclaimDiscipline{
		DisciplineID:         "developer-ecosystem-no-overclaim",
		Version:              "2026.04",
		ProjectionDisclaimer: developerEcosystemVal0ProjectionDisclaimer(),
	}
}

func DeveloperEcosystemVal0FoundationModel() DeveloperEcosystemVal0Foundation {
	return DeveloperEcosystemVal0Foundation{
		FoundationID:         "developer-ecosystem-discipline-foundation",
		Version:              "2026.04",
		OutputClassification: DeveloperEcosystemVal0OutputClassificationModel(),
		IDEAdvisory:          DeveloperEcosystemVal0IDEAdvisoryDisciplineModel(),
		LocalProduction:      DeveloperEcosystemVal0LocalProductionDisciplineModel(),
		RepoPolicyBoundary:   DeveloperEcosystemVal0RepoPolicyBoundaryDisciplineModel(),
		PluginSafety:         DeveloperEcosystemVal0PluginSafetyDisciplineModel(),
		PerformanceBudget:    DeveloperEcosystemVal0PerformanceBudgetDisciplineModel(),
		DXMetrics:            DeveloperEcosystemVal0DXMetricsDisciplineModel(),
		NoOverclaim:          DeveloperEcosystemVal0NoOverclaimDisciplineModel(),
		EvidenceRefs:         DeveloperEcosystemVal0ProofEvidenceRefs(),
		ProofSurfaceRefs:     DeveloperEcosystemVal0ProofSurfaceRefs(),
		ProjectionDisclaimer: developerEcosystemVal0ProjectionDisclaimer(),
		CreatedAt:            "2026-04-28T08:20:00Z",
		UpdatedAt:            "2026-04-28T08:20:00Z",
	}
}

func developerEcosystemVal0StateSeverity(state, active, partial, incomplete, blocked, unknown string) int {
	switch strings.TrimSpace(state) {
	case active:
		return 0
	case partial:
		return 1
	case incomplete:
		return 2
	case unknown:
		return 3
	case blocked:
		return 4
	default:
		return 3
	}
}

func EvaluateDeveloperEcosystemVal0DependencyState(snapshot DeveloperEcosystemVal0DependencySnapshot) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		snapshot.Point6State,
		snapshot.Point7State,
		snapshot.Point7ClosureState,
		snapshot.Point7PrerequisiteState,
		snapshot.Point7InvariantState,
		snapshot.Point7ProofSurfaceState,
		snapshot.Point7EvidenceQualityState,
		snapshot.Point7NoOverclaimState,
		snapshot.Point7PassRuleState,
	) {
		return DeveloperEcosystemVal0DependencyStateIncomplete
	}
	if strings.TrimSpace(snapshot.Point6State) != ReferenceArchitecturePoint6StatePass ||
		strings.TrimSpace(snapshot.Point7State) != VerifierEcosystemPoint7StatePass ||
		strings.TrimSpace(snapshot.Point7ClosureState) != VerifierEcosystemValEStatePass ||
		strings.TrimSpace(snapshot.Point7PrerequisiteState) != VerifierEcosystemValEPrerequisiteStateActive ||
		strings.TrimSpace(snapshot.Point7InvariantState) != VerifierEcosystemValEInvariantStateActive ||
		strings.TrimSpace(snapshot.Point7ProofSurfaceState) != VerifierEcosystemValEProofSurfaceStateActive ||
		strings.TrimSpace(snapshot.Point7EvidenceQualityState) != VerifierEcosystemValEEvidenceQualityStateActive ||
		strings.TrimSpace(snapshot.Point7NoOverclaimState) != VerifierEcosystemValENoOverclaimStateActive ||
		strings.TrimSpace(snapshot.Point7PassRuleState) != VerifierEcosystemValEPassRuleStateActive ||
		!snapshot.Point7PassAllowed {
		return DeveloperEcosystemVal0DependencyStateBlocked
	}
	return DeveloperEcosystemVal0DependencyStateActive
}

func EvaluateDeveloperEcosystemVal0OutputClassificationState(model DeveloperEcosystemVal0OutputClassification) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.ClassificationID, model.Version, model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0OutputClassificationStateIncomplete
	}
	if !developerEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0OutputClassificationStateUnknown
	}
	if !containsExactTrimmedStringSet(model.ClassifiedSurfaceKinds, developerEcosystemVal0ClassifiedSurfaces()...) ||
		!containsExactTrimmedStringSet(model.AllowedOutputClasses, developerEcosystemVal0OutputClasses()...) {
		return DeveloperEcosystemVal0OutputClassificationStateUnknown
	}
	if model.AdvisoryTreatedAsApproval || model.RecommendationsSuppressFailures || model.RedactionConvertsFailureToPass {
		return DeveloperEcosystemVal0OutputClassificationStateBlocked
	}
	if !model.UncertaintyVisible || !model.StalePartialVisible || !model.ProductionOnlyUnknownVisible {
		return DeveloperEcosystemVal0OutputClassificationStatePartial
	}
	return DeveloperEcosystemVal0OutputClassificationStateActive
}

func EvaluateDeveloperEcosystemVal0IDEAdvisoryState(model DeveloperEcosystemVal0IDEAdvisoryDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0IDEAdvisoryStateIncomplete
	}
	if !developerEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0IDEAdvisoryStateUnknown
	}
	if model.CanonicalTruthClaim || model.ProductionApprovalClaim || model.PolicyOverrideClaim || model.DeploymentApprovalClaim || model.CertificationClaim {
		return DeveloperEcosystemVal0IDEAdvisoryStateBlocked
	}
	if !model.AdvisoryOnly || !model.EvidenceLinked || !model.FreshnessAware || !model.ReasonCoded || !model.UncertaintyAware || !model.CandidateVsReviewedAware || !model.NonMutating || !model.NonApproving {
		return DeveloperEcosystemVal0IDEAdvisoryStatePartial
	}
	return DeveloperEcosystemVal0IDEAdvisoryStateActive
}

func EvaluateDeveloperEcosystemVal0LocalProductionState(model DeveloperEcosystemVal0LocalProductionDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.SimulationScope, model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0LocalProductionStateIncomplete
	}
	if !developerEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0LocalProductionStateUnknown
	}
	if !containsTrimmedString([]string{
		DeveloperEcosystemSimulationLocalValidation,
		DeveloperEcosystemSimulationMockVerification,
		DeveloperEcosystemSimulationOfflineBundle,
	}, model.SimulationScope) {
		return DeveloperEcosystemVal0LocalProductionStateUnknown
	}
	if model.ProductionEquivalenceClaim || model.MutatesCanonicalEvidence || model.ApprovalAuthority || model.SuppressesFailures {
		return DeveloperEcosystemVal0LocalProductionStateBlocked
	}
	if !model.UnsupportedCaseDisclosure || !model.ProductionOnlyUnknownDisclosure || !model.FreshnessContextAssumptions {
		return DeveloperEcosystemVal0LocalProductionStatePartial
	}
	return DeveloperEcosystemVal0LocalProductionStateActive
}

func EvaluateDeveloperEcosystemVal0RepoPolicyBoundaryState(model DeveloperEcosystemVal0RepoPolicyBoundaryDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.RepoConfigSchemaVersion,
		model.ScopeBoundary,
		model.EnterpriseOverrideBoundary,
		model.UnknownFieldHandling,
		model.UnsupportedVersionHandling,
		model.DeprecationBehavior,
		model.ProjectionDisclaimer,
	) {
		return DeveloperEcosystemVal0RepoPolicyStateIncomplete
	}
	if !developerEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0RepoPolicyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.SupportedSchemaVersions, DeveloperEcosystemRepoConfigSchemaV0) ||
		strings.TrimSpace(model.RepoConfigSchemaVersion) != DeveloperEcosystemRepoConfigSchemaV0 ||
		!containsExactTrimmedStringSet(model.AllowedFields, developerEcosystemVal0AllowedRepoFields()...) {
		return DeveloperEcosystemVal0RepoPolicyStateUnknown
	}
	if strings.TrimSpace(model.ScopeBoundary) != DeveloperEcosystemRepoScopeLocalAdvisoryOnly ||
		strings.TrimSpace(model.EnterpriseOverrideBoundary) != DeveloperEcosystemEnterpriseOverrideGoverned ||
		strings.TrimSpace(model.UnknownFieldHandling) != DeveloperEcosystemFailClosedHandling ||
		strings.TrimSpace(model.UnsupportedVersionHandling) != DeveloperEcosystemFailClosedHandling ||
		strings.TrimSpace(model.DeprecationBehavior) != DeveloperEcosystemDeprecationVisible {
		return DeveloperEcosystemVal0RepoPolicyStatePartial
	}
	if model.OverrideEnterprisePolicy || model.DisableCanonicalEvidenceRules || model.GrantApprovalAuthority || model.ChangeProductionTrustWithoutGov {
		return DeveloperEcosystemVal0RepoPolicyStateBlocked
	}
	if !model.ReviewRequired {
		return DeveloperEcosystemVal0RepoPolicyStatePartial
	}
	return DeveloperEcosystemVal0RepoPolicyStateActive
}

func EvaluateDeveloperEcosystemVal0PluginSafetyState(model DeveloperEcosystemVal0PluginSafetyDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.BoundedScope, model.PerformanceBudgetRef, model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0PluginSafetyStateIncomplete
	}
	if !developerEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0PluginSafetyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.DeclaredCapabilities, developerEcosystemVal0PluginCapabilities()...) {
		return DeveloperEcosystemVal0PluginSafetyStateUnknown
	}
	if strings.TrimSpace(model.PerformanceBudgetRef) != DeveloperEcosystemVal0PerformanceBudgetDisciplineID {
		return DeveloperEcosystemVal0PluginSafetyStateBlocked
	}
	if strings.TrimSpace(model.BoundedScope) != DeveloperEcosystemPluginScopeAdvisoryOnly {
		return DeveloperEcosystemVal0PluginSafetyStatePartial
	}
	if model.HiddenCanonicalMutation || model.HiddenPolicyOverride || model.HiddenApprovalPath || model.GovernanceBypass || model.CanonicalTruthClaim {
		return DeveloperEcosystemVal0PluginSafetyStateBlocked
	}
	if !model.SandboxIsolationExpected || !model.AuditVisibility || !model.DebugVisibility || !model.FailureStateVisible {
		return DeveloperEcosystemVal0PluginSafetyStatePartial
	}
	return DeveloperEcosystemVal0PluginSafetyStateActive
}

func EvaluateDeveloperEcosystemVal0PerformanceBudgetState(model DeveloperEcosystemVal0PerformanceBudgetDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0PerformanceBudgetStateIncomplete
	}
	if !developerEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0PerformanceBudgetStateUnknown
	}
	if len(model.Budgets) == 0 {
		return DeveloperEcosystemVal0PerformanceBudgetStateIncomplete
	}
	classes := make([]string, 0, len(model.Budgets))
	degradedVisible := true
	hasDegraded := false
	for _, budget := range model.Budgets {
		if strings.TrimSpace(budget.BudgetClass) == "" || budget.TargetMillis <= 0 {
			return DeveloperEcosystemVal0PerformanceBudgetStateUnknown
		}
		classes = append(classes, budget.BudgetClass)
		switch strings.TrimSpace(budget.BudgetState) {
		case DeveloperEcosystemBudgetStateKnown:
		case DeveloperEcosystemBudgetStateDegraded:
			hasDegraded = true
			if !budget.DegradedFallbackVisible {
				degradedVisible = false
			}
		case DeveloperEcosystemBudgetStateUnknown:
			return DeveloperEcosystemVal0PerformanceBudgetStateUnknown
		case DeveloperEcosystemBudgetStateUnsupported:
			return DeveloperEcosystemVal0PerformanceBudgetStateBlocked
		default:
			return DeveloperEcosystemVal0PerformanceBudgetStateUnknown
		}
	}
	if !containsExactTrimmedStringSet(classes, developerEcosystemVal0BudgetClasses()...) {
		return DeveloperEcosystemVal0PerformanceBudgetStateUnknown
	}
	if model.SilentBypassAllowed || model.HidesFailuresDegraded {
		return DeveloperEcosystemVal0PerformanceBudgetStateBlocked
	}
	if !model.DegradedModeVisible || !degradedVisible {
		return DeveloperEcosystemVal0PerformanceBudgetStatePartial
	}
	if hasDegraded {
		return DeveloperEcosystemVal0PerformanceBudgetStatePartial
	}
	return DeveloperEcosystemVal0PerformanceBudgetStateActive
}

func EvaluateDeveloperEcosystemVal0DXMetricsState(model DeveloperEcosystemVal0DXMetricsDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0DXMetricsStateIncomplete
	}
	if !developerEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0DXMetricsStateUnknown
	}
	if !containsExactTrimmedStringSet(model.MetricNames, developerEcosystemVal0DXMetricNames()...) {
		return DeveloperEcosystemVal0DXMetricsStateUnknown
	}
	if model.DeveloperTrustScore || model.CertificationUse || model.FastTrackApproval || model.DeploymentAuthority {
		return DeveloperEcosystemVal0DXMetricsStateBlocked
	}
	if !model.AdvisoryOnly {
		return DeveloperEcosystemVal0DXMetricsStatePartial
	}
	return DeveloperEcosystemVal0DXMetricsStateActive
}

func EvaluateDeveloperEcosystemVal0NoOverclaimState(model DeveloperEcosystemVal0NoOverclaimDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0NoOverclaimStateIncomplete
	}
	if !developerEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemVal0NoOverclaimStateUnknown
	}
	if model.CanonicalTruthClaim || model.ProductionApprovalClaim || model.DeveloperTrustScoreClaim || model.CertificationClaim || model.FastTrackClaim || model.EvidenceMutationClaim {
		return DeveloperEcosystemVal0NoOverclaimStateBlocked
	}
	return DeveloperEcosystemVal0NoOverclaimStateActive
}

func DeveloperEcosystemVal0ProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, evidenceRefs []string) bool {
	allFresh, stale, ok := verifierEcosystemVal0EvidenceValid(evidence)
	if !ok || !allFresh || stale || !containsExactTrimmedStringSet(evidenceRefs, DeveloperEcosystemVal0ProofEvidenceRefs()...) {
		return false
	}
	ids := make([]string, 0, len(evidence))
	scopes := make([]string, 0, len(evidence))
	for _, item := range evidence {
		ids = append(ids, item.EvidenceID)
		scopes = append(scopes, item.Scope)
	}
	return containsExactTrimmedStringSet(ids, developerEcosystemVal0RequiredEvidenceIDs()...) &&
		containsExactTrimmedStringSet(scopes, developerEcosystemVal0RequiredEvidenceScopes()...)
}

func EvaluateDeveloperEcosystemVal0State(model DeveloperEcosystemVal0Foundation) string {
	if EvaluateDeveloperEcosystemVal0DependencyState(model.Dependency) != DeveloperEcosystemVal0DependencyStateActive {
		return DeveloperEcosystemVal0StateBlocked
	}
	highestSeverity := 0
	for _, severity := range []int{
		developerEcosystemVal0StateSeverity(model.DependencyState, DeveloperEcosystemVal0DependencyStateActive, DeveloperEcosystemVal0DependencyStatePartial, DeveloperEcosystemVal0DependencyStateIncomplete, DeveloperEcosystemVal0DependencyStateBlocked, DeveloperEcosystemVal0DependencyStateUnknown),
		developerEcosystemVal0StateSeverity(model.OutputClassificationState, DeveloperEcosystemVal0OutputClassificationStateActive, DeveloperEcosystemVal0OutputClassificationStatePartial, DeveloperEcosystemVal0OutputClassificationStateIncomplete, DeveloperEcosystemVal0OutputClassificationStateBlocked, DeveloperEcosystemVal0OutputClassificationStateUnknown),
		developerEcosystemVal0StateSeverity(model.IDEAdvisoryState, DeveloperEcosystemVal0IDEAdvisoryStateActive, DeveloperEcosystemVal0IDEAdvisoryStatePartial, DeveloperEcosystemVal0IDEAdvisoryStateIncomplete, DeveloperEcosystemVal0IDEAdvisoryStateBlocked, DeveloperEcosystemVal0IDEAdvisoryStateUnknown),
		developerEcosystemVal0StateSeverity(model.LocalProductionState, DeveloperEcosystemVal0LocalProductionStateActive, DeveloperEcosystemVal0LocalProductionStatePartial, DeveloperEcosystemVal0LocalProductionStateIncomplete, DeveloperEcosystemVal0LocalProductionStateBlocked, DeveloperEcosystemVal0LocalProductionStateUnknown),
		developerEcosystemVal0StateSeverity(model.RepoPolicyBoundaryState, DeveloperEcosystemVal0RepoPolicyStateActive, DeveloperEcosystemVal0RepoPolicyStatePartial, DeveloperEcosystemVal0RepoPolicyStateIncomplete, DeveloperEcosystemVal0RepoPolicyStateBlocked, DeveloperEcosystemVal0RepoPolicyStateUnknown),
		developerEcosystemVal0StateSeverity(model.PluginSafetyState, DeveloperEcosystemVal0PluginSafetyStateActive, DeveloperEcosystemVal0PluginSafetyStatePartial, DeveloperEcosystemVal0PluginSafetyStateIncomplete, DeveloperEcosystemVal0PluginSafetyStateBlocked, DeveloperEcosystemVal0PluginSafetyStateUnknown),
		developerEcosystemVal0StateSeverity(model.PerformanceBudgetState, DeveloperEcosystemVal0PerformanceBudgetStateActive, DeveloperEcosystemVal0PerformanceBudgetStatePartial, DeveloperEcosystemVal0PerformanceBudgetStateIncomplete, DeveloperEcosystemVal0PerformanceBudgetStateBlocked, DeveloperEcosystemVal0PerformanceBudgetStateUnknown),
		developerEcosystemVal0StateSeverity(model.DXMetricsState, DeveloperEcosystemVal0DXMetricsStateActive, DeveloperEcosystemVal0DXMetricsStatePartial, DeveloperEcosystemVal0DXMetricsStateIncomplete, DeveloperEcosystemVal0DXMetricsStateBlocked, DeveloperEcosystemVal0DXMetricsStateUnknown),
		developerEcosystemVal0StateSeverity(model.NoOverclaimState, DeveloperEcosystemVal0NoOverclaimStateActive, DeveloperEcosystemVal0NoOverclaimStatePartial, DeveloperEcosystemVal0NoOverclaimStateIncomplete, DeveloperEcosystemVal0NoOverclaimStateBlocked, DeveloperEcosystemVal0NoOverclaimStateUnknown),
	} {
		if severity > highestSeverity {
			highestSeverity = severity
		}
	}
	switch highestSeverity {
	case 4:
		return DeveloperEcosystemVal0StateBlocked
	case 3:
		return DeveloperEcosystemVal0StateUnknown
	case 2:
		return DeveloperEcosystemVal0StateIncomplete
	case 1:
		return DeveloperEcosystemVal0StatePartial
	default:
		return DeveloperEcosystemVal0StateActive
	}
}

func EvaluateDeveloperEcosystemPoint8State(currentState string) string {
	_ = currentState
	return DeveloperEcosystemPoint8StateNotComplete
}

func EvaluateDeveloperEcosystemVal0ProofsState(model DeveloperEcosystemVal0Foundation, limitations []string) string {
	baseState := strings.TrimSpace(model.CurrentState)
	if baseState == "" {
		baseState = DeveloperEcosystemVal0StateUnknown
	}
	if !developerEcosystemVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, DeveloperEcosystemVal0ProofSurfaceRefs()...) ||
		!DeveloperEcosystemVal0ProofEvidenceQualityValid(developerEcosystemVal0Evidence(), model.EvidenceRefs) ||
		len(limitations) == 0 ||
		strings.TrimSpace(model.Point8State) != DeveloperEcosystemPoint8StateNotComplete {
		if baseState == DeveloperEcosystemVal0StateActive {
			return DeveloperEcosystemVal0StatePartial
		}
		return baseState
	}
	return baseState
}

func computeDeveloperEcosystemVal0BlockingReasons(model DeveloperEcosystemVal0Foundation) []string {
	reasons := []string{}
	if model.DependencyState != DeveloperEcosystemVal0DependencyStateActive {
		reasons = append(reasons, "Točka 8 Val 0 requires accepted Točka 7 Val E closure to remain pass and evidence-backed.")
	}
	if model.OutputClassificationState != DeveloperEcosystemVal0OutputClassificationStateActive {
		reasons = append(reasons, "Developer output classification must preserve advisory, uncertainty, stale, and production-only unknown visibility.")
	}
	if model.IDEAdvisoryState != DeveloperEcosystemVal0IDEAdvisoryStateActive {
		reasons = append(reasons, "IDE advisory discipline must remain advisory, evidence-linked, reason-coded, freshness-aware, and non-approving.")
	}
	if model.LocalProductionState != DeveloperEcosystemVal0LocalProductionStateActive {
		reasons = append(reasons, "Local and mock tooling must disclose unsupported and production-only unknown cases without claiming production equivalence.")
	}
	if model.RepoPolicyBoundaryState != DeveloperEcosystemVal0RepoPolicyStateActive {
		reasons = append(reasons, "Repo-local config must remain schema-bound and cannot override enterprise governance or canonical evidence rules.")
	}
	if model.PluginSafetyState != DeveloperEcosystemVal0PluginSafetyStateActive {
		reasons = append(reasons, "Plugin outputs must remain bounded, auditable, non-mutating, and non-approving.")
	}
	if model.PerformanceBudgetState != DeveloperEcosystemVal0PerformanceBudgetStateActive {
		reasons = append(reasons, "Developer performance budget regressions must remain visible and cannot silently bypass failures.")
	}
	if model.DXMetricsState != DeveloperEcosystemVal0DXMetricsStateActive {
		reasons = append(reasons, "DX metrics must remain advisory only and cannot create developer trust or fast-track approval signals.")
	}
	if model.NoOverclaimState != DeveloperEcosystemVal0NoOverclaimStateActive {
		reasons = append(reasons, "Developer tooling cannot claim certification, production approval, canonical truth, or point_8_pass in Val 0.")
	}
	return verifierEcosystemValECollectText(reasons)
}

func ComputeDeveloperEcosystemVal0Foundation(model DeveloperEcosystemVal0Foundation) DeveloperEcosystemVal0Foundation {
	model.DependencyState = EvaluateDeveloperEcosystemVal0DependencyState(model.Dependency)
	model.OutputClassificationState = EvaluateDeveloperEcosystemVal0OutputClassificationState(model.OutputClassification)
	model.IDEAdvisoryState = EvaluateDeveloperEcosystemVal0IDEAdvisoryState(model.IDEAdvisory)
	model.LocalProductionState = EvaluateDeveloperEcosystemVal0LocalProductionState(model.LocalProduction)
	model.RepoPolicyBoundaryState = EvaluateDeveloperEcosystemVal0RepoPolicyBoundaryState(model.RepoPolicyBoundary)
	model.PluginSafetyState = EvaluateDeveloperEcosystemVal0PluginSafetyState(model.PluginSafety)
	model.PerformanceBudgetState = EvaluateDeveloperEcosystemVal0PerformanceBudgetState(model.PerformanceBudget)
	model.DXMetricsState = EvaluateDeveloperEcosystemVal0DXMetricsState(model.DXMetrics)
	model.NoOverclaimState = EvaluateDeveloperEcosystemVal0NoOverclaimState(model.NoOverclaim)
	model.CurrentState = EvaluateDeveloperEcosystemVal0State(model)
	model.Point8State = EvaluateDeveloperEcosystemPoint8State(model.CurrentState)
	model.BlockingReasons = computeDeveloperEcosystemVal0BlockingReasons(model)
	return model
}

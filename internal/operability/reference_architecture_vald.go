package operability

import "strings"

const (
	ReferenceArchitectureValDVisibilityStateActive     = "reference_architecture_vald_operational_visibility_active"
	ReferenceArchitectureValDVisibilityStatePartial    = "reference_architecture_vald_operational_visibility_partial"
	ReferenceArchitectureValDVisibilityStateIncomplete = "reference_architecture_vald_operational_visibility_incomplete"
	ReferenceArchitectureValDVisibilityStateBlocked    = "reference_architecture_vald_operational_visibility_blocked"
	ReferenceArchitectureValDVisibilityStateUnknown    = "reference_architecture_vald_operational_visibility_unknown"

	ReferenceArchitectureValDAlignmentStateActive     = "reference_architecture_vald_alignment_summary_active"
	ReferenceArchitectureValDAlignmentStatePartial    = "reference_architecture_vald_alignment_summary_partial"
	ReferenceArchitectureValDAlignmentStateIncomplete = "reference_architecture_vald_alignment_summary_incomplete"
	ReferenceArchitectureValDAlignmentStateBlocked    = "reference_architecture_vald_alignment_summary_blocked"
	ReferenceArchitectureValDAlignmentStateUnknown    = "reference_architecture_vald_alignment_summary_unknown"

	ReferenceArchitectureValDAlertStateActive     = "reference_architecture_vald_deviation_alert_active"
	ReferenceArchitectureValDAlertStatePartial    = "reference_architecture_vald_deviation_alert_partial"
	ReferenceArchitectureValDAlertStateIncomplete = "reference_architecture_vald_deviation_alert_incomplete"
	ReferenceArchitectureValDAlertStateBlocked    = "reference_architecture_vald_deviation_alert_blocked"
	ReferenceArchitectureValDAlertStateUnknown    = "reference_architecture_vald_deviation_alert_unknown"

	ReferenceArchitectureValDSupportBoundaryStateActive     = "reference_architecture_vald_support_boundary_active"
	ReferenceArchitectureValDSupportBoundaryStatePartial    = "reference_architecture_vald_support_boundary_partial"
	ReferenceArchitectureValDSupportBoundaryStateIncomplete = "reference_architecture_vald_support_boundary_incomplete"
	ReferenceArchitectureValDSupportBoundaryStateBlocked    = "reference_architecture_vald_support_boundary_blocked"
	ReferenceArchitectureValDSupportBoundaryStateUnknown    = "reference_architecture_vald_support_boundary_unknown"

	ReferenceArchitectureValDMigrationStateActive     = "reference_architecture_vald_migration_upgrade_active"
	ReferenceArchitectureValDMigrationStatePartial    = "reference_architecture_vald_migration_upgrade_partial"
	ReferenceArchitectureValDMigrationStateIncomplete = "reference_architecture_vald_migration_upgrade_incomplete"
	ReferenceArchitectureValDMigrationStateBlocked    = "reference_architecture_vald_migration_upgrade_blocked"
	ReferenceArchitectureValDMigrationStateUnknown    = "reference_architecture_vald_migration_upgrade_unknown"

	ReferenceArchitectureValDTopologyGateStateActive     = "reference_architecture_vald_topology_gate_active"
	ReferenceArchitectureValDTopologyGateStatePartial    = "reference_architecture_vald_topology_gate_partial"
	ReferenceArchitectureValDTopologyGateStateIncomplete = "reference_architecture_vald_topology_gate_incomplete"
	ReferenceArchitectureValDTopologyGateStateBlocked    = "reference_architecture_vald_topology_gate_blocked"
	ReferenceArchitectureValDTopologyGateStateUnknown    = "reference_architecture_vald_topology_gate_unknown"

	ReferenceArchitectureValDSecurityGateStateActive     = "reference_architecture_vald_security_boundary_gate_active"
	ReferenceArchitectureValDSecurityGateStatePartial    = "reference_architecture_vald_security_boundary_gate_partial"
	ReferenceArchitectureValDSecurityGateStateIncomplete = "reference_architecture_vald_security_boundary_gate_incomplete"
	ReferenceArchitectureValDSecurityGateStateBlocked    = "reference_architecture_vald_security_boundary_gate_blocked"
	ReferenceArchitectureValDSecurityGateStateUnknown    = "reference_architecture_vald_security_boundary_gate_unknown"

	ReferenceArchitectureValDOperabilityGateStateActive     = "reference_architecture_vald_operability_gate_active"
	ReferenceArchitectureValDOperabilityGateStatePartial    = "reference_architecture_vald_operability_gate_partial"
	ReferenceArchitectureValDOperabilityGateStateIncomplete = "reference_architecture_vald_operability_gate_incomplete"
	ReferenceArchitectureValDOperabilityGateStateBlocked    = "reference_architecture_vald_operability_gate_blocked"
	ReferenceArchitectureValDOperabilityGateStateUnknown    = "reference_architecture_vald_operability_gate_unknown"

	ReferenceArchitectureValDCompatibilityGateStateActive     = "reference_architecture_vald_compatibility_gate_active"
	ReferenceArchitectureValDCompatibilityGateStatePartial    = "reference_architecture_vald_compatibility_gate_partial"
	ReferenceArchitectureValDCompatibilityGateStateIncomplete = "reference_architecture_vald_compatibility_gate_incomplete"
	ReferenceArchitectureValDCompatibilityGateStateBlocked    = "reference_architecture_vald_compatibility_gate_blocked"
	ReferenceArchitectureValDCompatibilityGateStateUnknown    = "reference_architecture_vald_compatibility_gate_unknown"

	ReferenceArchitectureValDFinalGateStateActive     = "reference_architecture_vald_final_gate_active"
	ReferenceArchitectureValDFinalGateStatePartial    = "reference_architecture_vald_final_gate_partial"
	ReferenceArchitectureValDFinalGateStateIncomplete = "reference_architecture_vald_final_gate_incomplete"
	ReferenceArchitectureValDFinalGateStateBlocked    = "reference_architecture_vald_final_gate_blocked"
	ReferenceArchitectureValDFinalGateStateUnknown    = "reference_architecture_vald_final_gate_unknown"

	ReferenceArchitectureValDStateActive     = "reference_architecture_vald_active"
	ReferenceArchitectureValDStatePartial    = "reference_architecture_vald_partial"
	ReferenceArchitectureValDStateIncomplete = "reference_architecture_vald_incomplete"
	ReferenceArchitectureValDStateBlocked    = "reference_architecture_vald_blocked"
	ReferenceArchitectureValDStateUnknown    = "reference_architecture_vald_unknown"

	ReferenceArchitectureValDDeviationBlueprintContractGap = "blueprint_contract_gap"
	ReferenceArchitectureValDDeviationFamilyProfileGap     = "family_profile_gap"
	ReferenceArchitectureValDDeviationMissingArtifactGap   = "missing_required_artifact"
	ReferenceArchitectureValDDeviationReadinessGap         = "readiness_gap"
	ReferenceArchitectureValDDeviationValidationHookGap    = "validation_hook_gap"
	ReferenceArchitectureValDDeviationConformanceGap       = "conformance_gap"
	ReferenceArchitectureValDDeviationResilienceGap        = "resilience_gap"
	ReferenceArchitectureValDDeviationRecoveryGap          = "recovery_gap"
	ReferenceArchitectureValDDeviationScalingGap           = "scaling_gap"
	ReferenceArchitectureValDDeviationTrustPathGap         = "trust_path_gap"
	ReferenceArchitectureValDDeviationAuditPathGap         = "audit_path_gap"
	ReferenceArchitectureValDDeviationControlPlaneGap      = "control_plane_safety_gap"
	ReferenceArchitectureValDDeviationSupportBoundaryGap   = "support_boundary_gap"
	ReferenceArchitectureValDDeviationMigrationGap         = "migration_upgrade_gap"
	ReferenceArchitectureValDDeviationStaleEvidence        = "stale_evidence"
	ReferenceArchitectureValDDeviationUnsupportedEnv       = "unsupported_environment"
	ReferenceArchitectureValDDeviationOverclaimDetected    = "overclaim_language_detected"
	ReferenceArchitectureValDDeviationUnknown              = "unknown"
)

type ReferenceArchitectureSourceValStates struct {
	Point5DependencyState string `json:"point_5_dependency_state"`
	Point5State           string `json:"point_5_state"`
	Val0State             string `json:"val_0_state"`
	ValAState             string `json:"val_a_state"`
	ValBState             string `json:"val_b_state"`
	ValCState             string `json:"val_c_state"`
}

type ReferenceArchitectureOperationalVisibilityReport struct {
	CurrentState             string                                   `json:"current_state"`
	VisibilityReportID       string                                   `json:"visibility_report_id"`
	Version                  string                                   `json:"version"`
	BlueprintFamily          string                                   `json:"blueprint_family"`
	BlueprintID              string                                   `json:"blueprint_id"`
	SourceValStates          ReferenceArchitectureSourceValStates     `json:"source_val_states"`
	Point6State              string                                   `json:"point_6_state"`
	AlignmentStatus          string                                   `json:"alignment_status"`
	ConformanceStatus        string                                   `json:"conformance_status"`
	ReadinessStatus          string                                   `json:"readiness_status"`
	ResilienceStatus         string                                   `json:"resilience_status"`
	SupportBoundaryStatus    string                                   `json:"support_boundary_status"`
	MigrationUpgradeStatus   string                                   `json:"migration_upgrade_status"`
	EvidenceRefs             []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	Caveats                  []string                                 `json:"caveats,omitempty"`
	Limitations              []string                                 `json:"limitations,omitempty"`
	ProjectionDisclaimer     string                                   `json:"projection_disclaimer"`
	CreatedAt                string                                   `json:"created_at"`
	UpdatedAt                string                                   `json:"updated_at"`
	CertifiedLanguagePresent bool                                     `json:"certified_language_present"`
	GuaranteedSecurityClaim  bool                                     `json:"guaranteed_security_claim_present"`
	ProductionApprovedClaim  bool                                     `json:"production_approved_claim_present"`
}

type ReferenceArchitectureOperationalVisibilityCollection struct {
	CurrentState         string                                             `json:"current_state"`
	CollectionID         string                                             `json:"collection_id"`
	SupportedFamilies    []string                                           `json:"supported_families,omitempty"`
	Reports              []ReferenceArchitectureOperationalVisibilityReport `json:"reports,omitempty"`
	ProjectionDisclaimer string                                             `json:"projection_disclaimer"`
}

type ReferenceArchitectureBlueprintAlignmentSummary struct {
	CurrentState                     string   `json:"current_state"`
	SummaryID                        string   `json:"summary_id"`
	BlueprintFamily                  string   `json:"blueprint_family"`
	Val0State                        string   `json:"val_0_state"`
	ValAState                        string   `json:"val_a_state"`
	ValBState                        string   `json:"val_b_state"`
	ValCState                        string   `json:"val_c_state"`
	AlignmentStatus                  string   `json:"alignment_status"`
	BlockingDeviations               []string `json:"blocking_deviations,omitempty"`
	NonBlockingCaveats               []string `json:"non_blocking_caveats,omitempty"`
	StaleEvidenceRefs                []string `json:"stale_evidence_refs,omitempty"`
	UnsupportedEnvironmentConditions []string `json:"unsupported_environment_conditions,omitempty"`
	SupportBoundaryGaps              []string `json:"support_boundary_gaps,omitempty"`
	RedactionKeepsBlockingVisible    bool     `json:"redaction_keeps_blocking_visible"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureBlueprintAlignmentCollection struct {
	CurrentState         string                                           `json:"current_state"`
	CollectionID         string                                           `json:"collection_id"`
	SupportedFamilies    []string                                         `json:"supported_families,omitempty"`
	Summaries            []ReferenceArchitectureBlueprintAlignmentSummary `json:"summaries,omitempty"`
	ProjectionDisclaimer string                                           `json:"projection_disclaimer"`
}

type ReferenceArchitectureDeviationAlert struct {
	AlertID                string   `json:"alert_id"`
	BlueprintFamily        string   `json:"blueprint_family"`
	SourceLayer            string   `json:"source_layer"`
	DeviationCategory      string   `json:"deviation_category"`
	Severity               string   `json:"severity"`
	AffectedScope          string   `json:"affected_scope"`
	EvidenceRef            string   `json:"evidence_ref"`
	BlocksAlignment        bool     `json:"blocks_alignment"`
	OperatorActionRequired string   `json:"operator_action_required"`
	SupportBoundaryRef     string   `json:"support_boundary_ref"`
	Timestamp              string   `json:"timestamp"`
	FreshnessState         string   `json:"freshness_state"`
	Caveats                []string `json:"caveats,omitempty"`
	AdvisoryOnly           bool     `json:"advisory_only"`
}

type ReferenceArchitectureDeviationAlertReport struct {
	CurrentState         string                                `json:"current_state"`
	ReportID             string                                `json:"report_id"`
	BlueprintFamily      string                                `json:"blueprint_family"`
	Alerts               []ReferenceArchitectureDeviationAlert `json:"alerts,omitempty"`
	ProjectionDisclaimer string                                `json:"projection_disclaimer"`
}

type ReferenceArchitectureDeviationAlertCollection struct {
	CurrentState          string                                      `json:"current_state"`
	CollectionID          string                                      `json:"collection_id"`
	SupportedFamilies     []string                                    `json:"supported_families,omitempty"`
	SupportedCategories   []string                                    `json:"supported_categories,omitempty"`
	SupportedSeverities   []string                                    `json:"supported_severities,omitempty"`
	SupportedSourceLayers []string                                    `json:"supported_source_layers,omitempty"`
	Reports               []ReferenceArchitectureDeviationAlertReport `json:"reports,omitempty"`
	ProjectionDisclaimer  string                                      `json:"projection_disclaimer"`
}

type ReferenceArchitectureSupportBoundaryView struct {
	CurrentState                     string   `json:"current_state"`
	ViewID                           string   `json:"view_id"`
	BlueprintFamily                  string   `json:"blueprint_family"`
	SupportedEnvironmentScope        string   `json:"supported_environment_scope"`
	UnsupportedConditions            []string `json:"unsupported_conditions,omitempty"`
	DegradedSupportConditions        []string `json:"degraded_support_conditions,omitempty"`
	OperatorResponsibility           string   `json:"operator_responsibility"`
	PartnerMSPBoundary               string   `json:"partner_msp_boundary"`
	VerifierAuditorBoundary          string   `json:"verifier_auditor_boundary"`
	AirGappedSovereignBoundary       string   `json:"air_gapped_sovereign_boundary"`
	EvidenceExportBoundary           string   `json:"evidence_export_boundary"`
	EscalationGuidanceRef            string   `json:"escalation_guidance_ref"`
	SupportBoundaryRef               string   `json:"support_boundary_ref"`
	Caveats                          []string `json:"caveats,omitempty"`
	PartnerCanonicalAuthority        bool     `json:"partner_canonical_authority"`
	RedactionKeepsUnsupportedVisible bool     `json:"redaction_keeps_unsupported_visible"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureSupportBoundaryCollection struct {
	CurrentState         string                                     `json:"current_state"`
	CollectionID         string                                     `json:"collection_id"`
	SupportedFamilies    []string                                   `json:"supported_families,omitempty"`
	Views                []ReferenceArchitectureSupportBoundaryView `json:"views,omitempty"`
	ProjectionDisclaimer string                                     `json:"projection_disclaimer"`
}

type ReferenceArchitectureMigrationUpgradeVisibility struct {
	CurrentState            string                                   `json:"current_state"`
	VisibilityID            string                                   `json:"visibility_id"`
	BlueprintFamily         string                                   `json:"blueprint_family"`
	CurrentBlueprintVersion string                                   `json:"current_blueprint_version"`
	TargetBlueprintVersion  string                                   `json:"target_blueprint_version"`
	CompatibilityState      string                                   `json:"compatibility_state"`
	DeprecationState        string                                   `json:"deprecation_state"`
	MigrationPathRef        string                                   `json:"migration_path_ref"`
	RollbackBoundaryRef     string                                   `json:"rollback_boundary_ref"`
	RequiredValidationRefs  []string                                 `json:"required_validation_refs,omitempty"`
	EvidenceRefs            []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	StaleIndicator          bool                                     `json:"stale_indicator"`
	SupersededIndicator     bool                                     `json:"superseded_indicator"`
	Caveats                 []string                                 `json:"caveats,omitempty"`
	ProjectionDisclaimer    string                                   `json:"projection_disclaimer"`
	ExecutesMigration       bool                                     `json:"executes_migration"`
}

type ReferenceArchitectureMigrationUpgradeCollection struct {
	CurrentState         string                                            `json:"current_state"`
	CollectionID         string                                            `json:"collection_id"`
	SupportedFamilies    []string                                          `json:"supported_families,omitempty"`
	Views                []ReferenceArchitectureMigrationUpgradeVisibility `json:"views,omitempty"`
	ProjectionDisclaimer string                                            `json:"projection_disclaimer"`
}

type ReferenceArchitectureTopologyGateCheck struct {
	CurrentState                        string                                   `json:"current_state"`
	CheckID                             string                                   `json:"check_id"`
	BlueprintFamily                     string                                   `json:"blueprint_family"`
	DeploymentTopology                  string                                   `json:"deployment_topology"`
	TrustAnchorMode                     string                                   `json:"trust_anchor_mode"`
	AuditPathMode                       string                                   `json:"audit_path_mode"`
	ConnectivityMode                    string                                   `json:"connectivity_mode"`
	SupportedTopology                   bool                                     `json:"supported_topology"`
	ControlDataPlaneSeparationRequired  bool                                     `json:"control_data_plane_separation_required"`
	ControlDataPlaneSeparationSatisfied bool                                     `json:"control_data_plane_separation_satisfied"`
	TrustAnchorTopologyCompatible       bool                                     `json:"trust_anchor_topology_compatible"`
	AuditPathTopologyCompatible         bool                                     `json:"audit_path_topology_compatible"`
	OfflineTopologyCompatible           bool                                     `json:"offline_topology_compatible"`
	UnsupportedConditions               []string                                 `json:"unsupported_conditions,omitempty"`
	EvidenceRefs                        []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	RedactionKeepsMismatchVisible       bool                                     `json:"redaction_keeps_mismatch_visible"`
	ProjectionDisclaimer                string                                   `json:"projection_disclaimer"`
}

type ReferenceArchitectureTopologyGateCollection struct {
	CurrentState         string                                   `json:"current_state"`
	CollectionID         string                                   `json:"collection_id"`
	SupportedFamilies    []string                                 `json:"supported_families,omitempty"`
	Checks               []ReferenceArchitectureTopologyGateCheck `json:"checks,omitempty"`
	ProjectionDisclaimer string                                   `json:"projection_disclaimer"`
}

type ReferenceArchitectureSecurityBoundaryGateCheck struct {
	CurrentState                   string                                   `json:"current_state"`
	CheckID                        string                                   `json:"check_id"`
	BlueprintFamily                string                                   `json:"blueprint_family"`
	TrustAnchorBoundary            string                                   `json:"trust_anchor_boundary"`
	SigningCustodyBoundary         string                                   `json:"signing_custody_boundary"`
	EvidenceStorageBoundary        string                                   `json:"evidence_storage_boundary"`
	PolicyDistributionBoundary     string                                   `json:"policy_distribution_boundary"`
	OperatorAccessBoundary         string                                   `json:"operator_access_boundary"`
	PartnerVerifierAuditorBoundary string                                   `json:"partner_verifier_auditor_boundary"`
	RedactionExportBoundary        string                                   `json:"redaction_export_boundary"`
	NoShadowTruthBoundary          string                                   `json:"no_shadow_truth_boundary"`
	MutationAuthorityBlocked       bool                                     `json:"mutation_authority_blocked"`
	ApprovalAuthorityBlocked       bool                                     `json:"approval_authority_blocked"`
	EvidenceRefs                   []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer           string                                   `json:"projection_disclaimer"`
}

type ReferenceArchitectureSecurityBoundaryCollection struct {
	CurrentState         string                                           `json:"current_state"`
	CollectionID         string                                           `json:"collection_id"`
	SupportedFamilies    []string                                         `json:"supported_families,omitempty"`
	Checks               []ReferenceArchitectureSecurityBoundaryGateCheck `json:"checks,omitempty"`
	ProjectionDisclaimer string                                           `json:"projection_disclaimer"`
}

type ReferenceArchitectureOperabilityGateCheck struct {
	CurrentState             string                                   `json:"current_state"`
	CheckID                  string                                   `json:"check_id"`
	BlueprintFamily          string                                   `json:"blueprint_family"`
	ReadinessState           string                                   `json:"readiness_state"`
	ResilienceState          string                                   `json:"resilience_state"`
	RecoveryExpectationState string                                   `json:"recovery_expectation_state"`
	AuditPathState           string                                   `json:"audit_path_state"`
	ControlPlaneState        string                                   `json:"control_plane_state"`
	SupportBoundaryState     string                                   `json:"support_boundary_state"`
	OperatorActionGuidance   string                                   `json:"operator_action_guidance"`
	EvidenceRefs             []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	DegradedStateVisible     bool                                     `json:"degraded_state_visible"`
	ProjectionDisclaimer     string                                   `json:"projection_disclaimer"`
}

type ReferenceArchitectureOperabilityGateCollection struct {
	CurrentState         string                                      `json:"current_state"`
	CollectionID         string                                      `json:"collection_id"`
	SupportedFamilies    []string                                    `json:"supported_families,omitempty"`
	Checks               []ReferenceArchitectureOperabilityGateCheck `json:"checks,omitempty"`
	ProjectionDisclaimer string                                      `json:"projection_disclaimer"`
}

type ReferenceArchitectureCompatibilityGateCheck struct {
	CurrentState              string                                   `json:"current_state"`
	CheckID                   string                                   `json:"check_id"`
	BlueprintFamily           string                                   `json:"blueprint_family"`
	LifecycleState            string                                   `json:"lifecycle_state"`
	CompatibilityState        string                                   `json:"compatibility_state"`
	MigrationVisibilityRef    string                                   `json:"migration_visibility_ref"`
	ValidationRequirementRefs []string                                 `json:"validation_requirement_refs,omitempty"`
	EvidenceRefs              []ReferenceArchitectureEvidenceReference `json:"evidence_refs,omitempty"`
	UniversalSupportClaim     bool                                     `json:"universal_support_claim"`
	ProjectionDisclaimer      string                                   `json:"projection_disclaimer"`
}

type ReferenceArchitectureCompatibilityGateCollection struct {
	CurrentState         string                                        `json:"current_state"`
	CollectionID         string                                        `json:"collection_id"`
	SupportedFamilies    []string                                      `json:"supported_families,omitempty"`
	Checks               []ReferenceArchitectureCompatibilityGateCheck `json:"checks,omitempty"`
	ProjectionDisclaimer string                                        `json:"projection_disclaimer"`
}

type ReferenceArchitectureFinalGateReport struct {
	CurrentState               string   `json:"current_state"`
	GateID                     string   `json:"gate_id"`
	BlueprintFamily            string   `json:"blueprint_family"`
	OperationalVisibilityState string   `json:"operational_visibility_state"`
	AlignmentSummaryState      string   `json:"alignment_summary_state"`
	DeviationAlertState        string   `json:"deviation_alert_state"`
	SupportBoundaryState       string   `json:"support_boundary_state"`
	MigrationUpgradeState      string   `json:"migration_upgrade_state"`
	TopologyGateState          string   `json:"topology_gate_state"`
	SecurityBoundaryGateState  string   `json:"security_boundary_gate_state"`
	OperabilityGateState       string   `json:"operability_gate_state"`
	CompatibilityGateState     string   `json:"compatibility_gate_state"`
	BlockingReasons            []string `json:"blocking_reasons,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type ReferenceArchitectureFinalGateCollection struct {
	CurrentState         string                                 `json:"current_state"`
	CollectionID         string                                 `json:"collection_id"`
	SupportedFamilies    []string                               `json:"supported_families,omitempty"`
	Reports              []ReferenceArchitectureFinalGateReport `json:"reports,omitempty"`
	ProjectionDisclaimer string                                 `json:"projection_disclaimer"`
}

func referenceArchitectureValDProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_operational_visibility_final_reference_gate"
}

func referenceArchitectureValDHasProjectionDisclaimer(value string) bool {
	return strings.Contains(strings.TrimSpace(value), "projection_only") &&
		strings.Contains(strings.TrimSpace(value), "not_canonical_truth")
}

func referenceArchitectureValDSourceLayers() []string {
	return []string{"val0", "vala", "valb", "valc", "vald"}
}

func referenceArchitectureValDAlertCategories() []string {
	return []string{
		ReferenceArchitectureValDDeviationBlueprintContractGap,
		ReferenceArchitectureValDDeviationFamilyProfileGap,
		ReferenceArchitectureValDDeviationMissingArtifactGap,
		ReferenceArchitectureValDDeviationReadinessGap,
		ReferenceArchitectureValDDeviationValidationHookGap,
		ReferenceArchitectureValDDeviationConformanceGap,
		ReferenceArchitectureValDDeviationResilienceGap,
		ReferenceArchitectureValDDeviationRecoveryGap,
		ReferenceArchitectureValDDeviationScalingGap,
		ReferenceArchitectureValDDeviationTrustPathGap,
		ReferenceArchitectureValDDeviationAuditPathGap,
		ReferenceArchitectureValDDeviationControlPlaneGap,
		ReferenceArchitectureValDDeviationSupportBoundaryGap,
		ReferenceArchitectureValDDeviationMigrationGap,
		ReferenceArchitectureValDDeviationStaleEvidence,
		ReferenceArchitectureValDDeviationUnsupportedEnv,
		ReferenceArchitectureValDDeviationOverclaimDetected,
		ReferenceArchitectureValDDeviationUnknown,
	}
}

func referenceArchitectureValDRequiredAlertCategories() []string {
	return []string{
		ReferenceArchitectureValDDeviationBlueprintContractGap,
		ReferenceArchitectureValDDeviationFamilyProfileGap,
		ReferenceArchitectureValDDeviationMissingArtifactGap,
		ReferenceArchitectureValDDeviationReadinessGap,
		ReferenceArchitectureValDDeviationValidationHookGap,
		ReferenceArchitectureValDDeviationConformanceGap,
		ReferenceArchitectureValDDeviationResilienceGap,
		ReferenceArchitectureValDDeviationRecoveryGap,
		ReferenceArchitectureValDDeviationScalingGap,
		ReferenceArchitectureValDDeviationTrustPathGap,
		ReferenceArchitectureValDDeviationAuditPathGap,
		ReferenceArchitectureValDDeviationControlPlaneGap,
		ReferenceArchitectureValDDeviationSupportBoundaryGap,
		ReferenceArchitectureValDDeviationMigrationGap,
		ReferenceArchitectureValDDeviationStaleEvidence,
		ReferenceArchitectureValDDeviationUnsupportedEnv,
		ReferenceArchitectureValDDeviationOverclaimDetected,
	}
}

func referenceArchitectureValDPoint5DependencyHealthy(state string) bool {
	return strings.TrimSpace(state) == IntelligenceCalibrationValEStateActive
}

func referenceArchitectureValDNormalizedFamilyKey(value string) (string, bool) {
	normalized := strings.TrimSpace(value)
	if normalized == "" {
		return "", false
	}
	return normalized, true
}

func referenceArchitectureValDRegisterFamily(seen map[string]struct{}, family string) bool {
	normalized, ok := referenceArchitectureValDNormalizedFamilyKey(family)
	if !ok {
		return false
	}
	if _, duplicate := seen[normalized]; duplicate {
		return false
	}
	seen[normalized] = struct{}{}
	return true
}

func referenceArchitectureValDSeenFamiliesMatchSupported(seen map[string]struct{}) bool {
	families := make([]string, 0, len(seen))
	for family := range seen {
		families = append(families, family)
	}
	return containsExactTrimmedStringSet(families, referenceArchitectureVal0Families()...)
}

func referenceArchitectureValDFinalGateHasSupportedComponentStates(report ReferenceArchitectureFinalGateReport) bool {
	return containsTrimmedString([]string{
		ReferenceArchitectureValDVisibilityStateActive,
		ReferenceArchitectureValDVisibilityStatePartial,
		ReferenceArchitectureValDVisibilityStateIncomplete,
		ReferenceArchitectureValDVisibilityStateBlocked,
		ReferenceArchitectureValDVisibilityStateUnknown,
	}, report.OperationalVisibilityState) &&
		containsTrimmedString([]string{
			ReferenceArchitectureValDAlignmentStateActive,
			ReferenceArchitectureValDAlignmentStatePartial,
			ReferenceArchitectureValDAlignmentStateIncomplete,
			ReferenceArchitectureValDAlignmentStateBlocked,
			ReferenceArchitectureValDAlignmentStateUnknown,
		}, report.AlignmentSummaryState) &&
		containsTrimmedString([]string{
			ReferenceArchitectureValDAlertStateActive,
			ReferenceArchitectureValDAlertStatePartial,
			ReferenceArchitectureValDAlertStateIncomplete,
			ReferenceArchitectureValDAlertStateBlocked,
			ReferenceArchitectureValDAlertStateUnknown,
		}, report.DeviationAlertState) &&
		containsTrimmedString([]string{
			ReferenceArchitectureValDSupportBoundaryStateActive,
			ReferenceArchitectureValDSupportBoundaryStatePartial,
			ReferenceArchitectureValDSupportBoundaryStateIncomplete,
			ReferenceArchitectureValDSupportBoundaryStateBlocked,
			ReferenceArchitectureValDSupportBoundaryStateUnknown,
		}, report.SupportBoundaryState) &&
		containsTrimmedString([]string{
			ReferenceArchitectureValDMigrationStateActive,
			ReferenceArchitectureValDMigrationStatePartial,
			ReferenceArchitectureValDMigrationStateIncomplete,
			ReferenceArchitectureValDMigrationStateBlocked,
			ReferenceArchitectureValDMigrationStateUnknown,
		}, report.MigrationUpgradeState) &&
		containsTrimmedString([]string{
			ReferenceArchitectureValDTopologyGateStateActive,
			ReferenceArchitectureValDTopologyGateStatePartial,
			ReferenceArchitectureValDTopologyGateStateIncomplete,
			ReferenceArchitectureValDTopologyGateStateBlocked,
			ReferenceArchitectureValDTopologyGateStateUnknown,
		}, report.TopologyGateState) &&
		containsTrimmedString([]string{
			ReferenceArchitectureValDSecurityGateStateActive,
			ReferenceArchitectureValDSecurityGateStatePartial,
			ReferenceArchitectureValDSecurityGateStateIncomplete,
			ReferenceArchitectureValDSecurityGateStateBlocked,
			ReferenceArchitectureValDSecurityGateStateUnknown,
		}, report.SecurityBoundaryGateState) &&
		containsTrimmedString([]string{
			ReferenceArchitectureValDOperabilityGateStateActive,
			ReferenceArchitectureValDOperabilityGateStatePartial,
			ReferenceArchitectureValDOperabilityGateStateIncomplete,
			ReferenceArchitectureValDOperabilityGateStateBlocked,
			ReferenceArchitectureValDOperabilityGateStateUnknown,
		}, report.OperabilityGateState) &&
		containsTrimmedString([]string{
			ReferenceArchitectureValDCompatibilityGateStateActive,
			ReferenceArchitectureValDCompatibilityGateStatePartial,
			ReferenceArchitectureValDCompatibilityGateStateIncomplete,
			ReferenceArchitectureValDCompatibilityGateStateBlocked,
			ReferenceArchitectureValDCompatibilityGateStateUnknown,
		}, report.CompatibilityGateState)
}

func referenceArchitectureValDDerivedSourceStates() ReferenceArchitectureSourceValStates {
	return ReferenceArchitectureSourceValStates{
		Point5DependencyState: IntelligenceCalibrationValEStateActive,
		Point5State:           IntelligenceCalibrationPoint5StatePass,
		Val0State:             ReferenceArchitectureVal0StateActive,
		ValAState:             ReferenceArchitectureValAStateActive,
		ValBState:             ReferenceArchitectureValBStateActive,
		ValCState:             ReferenceArchitectureValCStateActive,
	}
}

func referenceArchitectureValDVisibilityEvidenceRefs(pack ReferenceArchitectureBlueprintPack) []ReferenceArchitectureEvidenceReference {
	refs := append([]ReferenceArchitectureEvidenceReference{}, pack.EvidenceRefs...)
	refs = append(refs, ReferenceArchitectureEvidenceReference{
		EvidenceID:     "vald-visibility/" + pack.BlueprintFamily,
		EvidenceType:   ReferenceArchitectureEvidenceSupportBoundary,
		Source:         "operational-visibility",
		Timestamp:      pack.UpdatedAt,
		FreshnessState: IntelligenceCalibrationFreshnessFresh,
		Scope:          pack.BlueprintFamily,
		Caveats:        []string{"dashboard-ready summary only"},
	})
	return refs
}

func referenceArchitectureValDOperationalVisibilityReportForProfile(profile ReferenceArchitectureBlueprintFamilyProfile, pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureOperationalVisibilityReport {
	return ReferenceArchitectureOperationalVisibilityReport{
		CurrentState:           "reference_architecture_vald_operational_visibility_ready",
		VisibilityReportID:     "visibility/" + profile.Family,
		Version:                pack.Version,
		BlueprintFamily:        profile.Family,
		BlueprintID:            profile.BlueprintID,
		SourceValStates:        referenceArchitectureValDDerivedSourceStates(),
		Point6State:            ReferenceArchitecturePoint6StateNotComplete,
		AlignmentStatus:        ReferenceArchitectureConformanceMatched,
		ConformanceStatus:      ReferenceArchitectureConformanceMatched,
		ReadinessStatus:        ReferenceArchitectureValBReadinessStateActive,
		ResilienceStatus:       ReferenceArchitectureValCStateActive,
		SupportBoundaryStatus:  ReferenceArchitectureValDSupportBoundaryStateActive,
		MigrationUpgradeStatus: ReferenceArchitectureValDMigrationStateActive,
		EvidenceRefs:           referenceArchitectureValDVisibilityEvidenceRefs(pack),
		Caveats:                append([]string{}, profile.Caveats...),
		Limitations:            []string{"dashboard-ready operational summary remains advisory"},
		ProjectionDisclaimer:   referenceArchitectureValDProjectionDisclaimer(),
		CreatedAt:              pack.CreatedAt,
		UpdatedAt:              pack.UpdatedAt,
	}
}

func ReferenceArchitectureValDOperationalVisibilityCollection() ReferenceArchitectureOperationalVisibilityCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	reports := make([]ReferenceArchitectureOperationalVisibilityReport, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		profile, _ := LookupReferenceArchitectureValAFamilyProfile(pack.BlueprintFamily)
		reports = append(reports, referenceArchitectureValDOperationalVisibilityReportForProfile(profile, pack))
	}
	return ReferenceArchitectureOperationalVisibilityCollection{
		CurrentState:         "reference_architecture_vald_operational_visibility_collection_ready",
		CollectionID:         "reference-architecture-vald-operational-visibility",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Reports:              reports,
		ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
	}
}

func referenceArchitectureValDAlignmentSummaryForProfile(profile ReferenceArchitectureBlueprintFamilyProfile) ReferenceArchitectureBlueprintAlignmentSummary {
	return ReferenceArchitectureBlueprintAlignmentSummary{
		CurrentState:                     "reference_architecture_vald_alignment_summary_ready",
		SummaryID:                        "alignment/" + profile.Family,
		BlueprintFamily:                  profile.Family,
		Val0State:                        ReferenceArchitectureVal0StateActive,
		ValAState:                        ReferenceArchitectureValAStateActive,
		ValBState:                        ReferenceArchitectureValBStateActive,
		ValCState:                        ReferenceArchitectureValCStateActive,
		AlignmentStatus:                  ReferenceArchitectureConformanceMatched,
		BlockingDeviations:               nil,
		NonBlockingCaveats:               append([]string{}, profile.Caveats...),
		StaleEvidenceRefs:                nil,
		UnsupportedEnvironmentConditions: nil,
		SupportBoundaryGaps:              nil,
		RedactionKeepsBlockingVisible:    true,
		ProjectionDisclaimer:             referenceArchitectureValDProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValDAlignmentSummaryCollection() ReferenceArchitectureBlueprintAlignmentCollection {
	registry := ReferenceArchitectureValAFamilyRegistry()
	summaries := make([]ReferenceArchitectureBlueprintAlignmentSummary, 0, len(registry.Profiles))
	for _, profile := range registry.Profiles {
		summaries = append(summaries, referenceArchitectureValDAlignmentSummaryForProfile(profile))
	}
	return ReferenceArchitectureBlueprintAlignmentCollection{
		CurrentState:         "reference_architecture_vald_alignment_collection_ready",
		CollectionID:         "reference-architecture-vald-alignment-summaries",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Summaries:            summaries,
		ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValDDeviationAlertCollection() ReferenceArchitectureDeviationAlertCollection {
	reports := make([]ReferenceArchitectureDeviationAlertReport, 0, len(referenceArchitectureVal0Families()))
	for _, family := range referenceArchitectureVal0Families() {
		reports = append(reports, ReferenceArchitectureDeviationAlertReport{
			CurrentState:         "reference_architecture_vald_deviation_alert_report_ready",
			ReportID:             "alerts/" + family,
			BlueprintFamily:      family,
			Alerts:               nil,
			ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
		})
	}
	return ReferenceArchitectureDeviationAlertCollection{
		CurrentState:          "reference_architecture_vald_deviation_alert_collection_ready",
		CollectionID:          "reference-architecture-vald-deviation-alerts",
		SupportedFamilies:     referenceArchitectureVal0Families(),
		SupportedCategories:   referenceArchitectureValDAlertCategories(),
		SupportedSeverities:   referenceArchitectureValCScenarioSeverities(),
		SupportedSourceLayers: referenceArchitectureValDSourceLayers(),
		Reports:               reports,
		ProjectionDisclaimer:  referenceArchitectureValDProjectionDisclaimer(),
	}
}

func referenceArchitectureValDSupportBoundaryViewForProfile(profile ReferenceArchitectureBlueprintFamilyProfile) ReferenceArchitectureSupportBoundaryView {
	partnerBoundary := "no partner boundary beyond support escalation"
	if profile.Family == ReferenceArchitectureFamilyPartnerMSPSuitable {
		partnerBoundary = "partner visibility is bounded; customer remains canonical authority"
	}
	airGapBoundary := "standard online support boundary"
	if profile.LocalTrustAnchorAssumptionRequired {
		airGapBoundary = "local sovereign or air-gapped evidence handoff boundary"
	}
	return ReferenceArchitectureSupportBoundaryView{
		CurrentState:                     "reference_architecture_vald_support_boundary_view_ready",
		ViewID:                           "support-view/" + profile.Family,
		BlueprintFamily:                  profile.Family,
		SupportedEnvironmentScope:        profile.TargetEnvironment.DeploymentTopology,
		UnsupportedConditions:            append([]string{}, profile.UnsupportedConditions...),
		DegradedSupportConditions:        append([]string{}, profile.DegradedConditions...),
		OperatorResponsibility:           "operators maintain bounded support, evidence review, and escalation discipline",
		PartnerMSPBoundary:               partnerBoundary,
		VerifierAuditorBoundary:          profile.TargetEnvironment.VerifierOrPartnerAccessMode,
		AirGappedSovereignBoundary:       airGapBoundary,
		EvidenceExportBoundary:           "bounded evidence export and redaction review boundary",
		EscalationGuidanceRef:            "support-escalation/" + profile.Family,
		SupportBoundaryRef:               profile.SupportBoundaryRef,
		Caveats:                          append([]string{}, profile.Caveats...),
		PartnerCanonicalAuthority:        profile.PartnerCanonicalTruthOverrideAllowed,
		RedactionKeepsUnsupportedVisible: true,
		ProjectionDisclaimer:             referenceArchitectureValDProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValDSupportBoundaryCollection() ReferenceArchitectureSupportBoundaryCollection {
	registry := ReferenceArchitectureValAFamilyRegistry()
	views := make([]ReferenceArchitectureSupportBoundaryView, 0, len(registry.Profiles))
	for _, profile := range registry.Profiles {
		views = append(views, referenceArchitectureValDSupportBoundaryViewForProfile(profile))
	}
	return ReferenceArchitectureSupportBoundaryCollection{
		CurrentState:         "reference_architecture_vald_support_boundary_collection_ready",
		CollectionID:         "reference-architecture-vald-support-boundaries",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Views:                views,
		ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
	}
}

func referenceArchitectureValDMigrationVisibilityForPack(pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureMigrationUpgradeVisibility {
	return ReferenceArchitectureMigrationUpgradeVisibility{
		CurrentState:            "reference_architecture_vald_migration_upgrade_ready",
		VisibilityID:            "migration/" + pack.BlueprintFamily,
		BlueprintFamily:         pack.BlueprintFamily,
		CurrentBlueprintVersion: pack.Version,
		TargetBlueprintVersion:  pack.Version,
		CompatibilityState:      pack.CompatibilityState,
		DeprecationState:        pack.LifecycleState,
		MigrationPathRef:        "migration-path/" + pack.BlueprintFamily,
		RollbackBoundaryRef:     pack.SupportBoundaryRef,
		RequiredValidationRefs:  []string{pack.ValidationPackRef, pack.ConformanceKitRef},
		EvidenceRefs:            append([]ReferenceArchitectureEvidenceReference{}, pack.EvidenceRefs...),
		StaleIndicator:          false,
		SupersededIndicator:     false,
		Caveats:                 append([]string{}, pack.Caveats...),
		ProjectionDisclaimer:    referenceArchitectureValDProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValDMigrationUpgradeCollection() ReferenceArchitectureMigrationUpgradeCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	views := make([]ReferenceArchitectureMigrationUpgradeVisibility, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		views = append(views, referenceArchitectureValDMigrationVisibilityForPack(pack))
	}
	return ReferenceArchitectureMigrationUpgradeCollection{
		CurrentState:         "reference_architecture_vald_migration_upgrade_collection_ready",
		CollectionID:         "reference-architecture-vald-migration-upgrade",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Views:                views,
		ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
	}
}

func referenceArchitectureValDTopologyGateCheckForProfile(profile ReferenceArchitectureBlueprintFamilyProfile, pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureTopologyGateCheck {
	return ReferenceArchitectureTopologyGateCheck{
		CurrentState:                        "reference_architecture_vald_topology_gate_ready",
		CheckID:                             "topology/" + profile.Family,
		BlueprintFamily:                     profile.Family,
		DeploymentTopology:                  profile.TargetEnvironment.DeploymentTopology,
		TrustAnchorMode:                     profile.TargetEnvironment.TrustAnchorMode,
		AuditPathMode:                       profile.TargetEnvironment.AuditPathMode,
		ConnectivityMode:                    profile.TargetEnvironment.ConnectivityMode,
		SupportedTopology:                   true,
		ControlDataPlaneSeparationRequired:  profile.StrongerTrustAnchorMode || profile.PerformanceEnvelopeRequired,
		ControlDataPlaneSeparationSatisfied: true,
		TrustAnchorTopologyCompatible:       true,
		AuditPathTopologyCompatible:         true,
		OfflineTopologyCompatible:           !profile.LocalTrustAnchorAssumptionRequired || !profile.LiveExternalDependencyAllowedOffline,
		UnsupportedConditions:               nil,
		EvidenceRefs:                        append([]ReferenceArchitectureEvidenceReference{}, pack.EvidenceRefs...),
		RedactionKeepsMismatchVisible:       true,
		ProjectionDisclaimer:                referenceArchitectureValDProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValDTopologyGateCollection() ReferenceArchitectureTopologyGateCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	checks := make([]ReferenceArchitectureTopologyGateCheck, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		profile, _ := LookupReferenceArchitectureValAFamilyProfile(pack.BlueprintFamily)
		checks = append(checks, referenceArchitectureValDTopologyGateCheckForProfile(profile, pack))
	}
	return ReferenceArchitectureTopologyGateCollection{
		CurrentState:         "reference_architecture_vald_topology_gate_collection_ready",
		CollectionID:         "reference-architecture-vald-topology-gate",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Checks:               checks,
		ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
	}
}

func referenceArchitectureValDSecurityBoundaryCheckForProfile(profile ReferenceArchitectureBlueprintFamilyProfile, pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureSecurityBoundaryGateCheck {
	noShadowTruth := "bounded no-shadow-truth rule remains explicit"
	if profile.Family == ReferenceArchitectureFamilyPartnerMSPSuitable {
		noShadowTruth = "partner must not create canonical shadow truth or approval authority"
	}
	return ReferenceArchitectureSecurityBoundaryGateCheck{
		CurrentState:                   "reference_architecture_vald_security_boundary_gate_ready",
		CheckID:                        "security/" + profile.Family,
		BlueprintFamily:                profile.Family,
		TrustAnchorBoundary:            profile.TargetEnvironment.TrustAnchorMode,
		SigningCustodyBoundary:         "governed signer custody boundary",
		EvidenceStorageBoundary:        "bounded evidence storage custody boundary",
		PolicyDistributionBoundary:     "bounded policy distribution authority boundary",
		OperatorAccessBoundary:         profile.TargetEnvironment.OperatorControlModel,
		PartnerVerifierAuditorBoundary: profile.TargetEnvironment.VerifierOrPartnerAccessMode,
		RedactionExportBoundary:        "bounded redaction and export control boundary",
		NoShadowTruthBoundary:          noShadowTruth,
		MutationAuthorityBlocked:       true,
		ApprovalAuthorityBlocked:       true,
		EvidenceRefs:                   append([]ReferenceArchitectureEvidenceReference{}, pack.EvidenceRefs...),
		ProjectionDisclaimer:           referenceArchitectureValDProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValDSecurityBoundaryCollection() ReferenceArchitectureSecurityBoundaryCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	checks := make([]ReferenceArchitectureSecurityBoundaryGateCheck, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		profile, _ := LookupReferenceArchitectureValAFamilyProfile(pack.BlueprintFamily)
		checks = append(checks, referenceArchitectureValDSecurityBoundaryCheckForProfile(profile, pack))
	}
	return ReferenceArchitectureSecurityBoundaryCollection{
		CurrentState:         "reference_architecture_vald_security_boundary_collection_ready",
		CollectionID:         "reference-architecture-vald-security-boundary-gate",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Checks:               checks,
		ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
	}
}

func referenceArchitectureValDOperabilityGateCheckForFamily(family string, pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureOperabilityGateCheck {
	return ReferenceArchitectureOperabilityGateCheck{
		CurrentState:             "reference_architecture_vald_operability_gate_ready",
		CheckID:                  "operability/" + family,
		BlueprintFamily:          family,
		ReadinessState:           ReferenceArchitectureValBReadinessStateActive,
		ResilienceState:          ReferenceArchitectureValCStateActive,
		RecoveryExpectationState: ReferenceArchitectureValCRecoveryExpectationStateActive,
		AuditPathState:           ReferenceArchitectureValCAuditPathStateActive,
		ControlPlaneState:        ReferenceArchitectureValCControlPlaneStateActive,
		SupportBoundaryState:     ReferenceArchitectureValDSupportBoundaryStateActive,
		OperatorActionGuidance:   "operators review readiness, resilience, recovery, and support guidance before treating alignment as clean",
		EvidenceRefs:             append([]ReferenceArchitectureEvidenceReference{}, pack.EvidenceRefs...),
		DegradedStateVisible:     true,
		ProjectionDisclaimer:     referenceArchitectureValDProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValDOperabilityGateCollection() ReferenceArchitectureOperabilityGateCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	checks := make([]ReferenceArchitectureOperabilityGateCheck, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		checks = append(checks, referenceArchitectureValDOperabilityGateCheckForFamily(pack.BlueprintFamily, pack))
	}
	return ReferenceArchitectureOperabilityGateCollection{
		CurrentState:         "reference_architecture_vald_operability_gate_collection_ready",
		CollectionID:         "reference-architecture-vald-operability-gate",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Checks:               checks,
		ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
	}
}

func referenceArchitectureValDCompatibilityGateCheckForPack(pack ReferenceArchitectureBlueprintPack) ReferenceArchitectureCompatibilityGateCheck {
	return ReferenceArchitectureCompatibilityGateCheck{
		CurrentState:              "reference_architecture_vald_compatibility_gate_ready",
		CheckID:                   "compatibility/" + pack.BlueprintFamily,
		BlueprintFamily:           pack.BlueprintFamily,
		LifecycleState:            pack.LifecycleState,
		CompatibilityState:        pack.CompatibilityState,
		MigrationVisibilityRef:    "migration/" + pack.BlueprintFamily,
		ValidationRequirementRefs: []string{pack.ValidationPackRef, pack.ConformanceKitRef},
		EvidenceRefs:              append([]ReferenceArchitectureEvidenceReference{}, pack.EvidenceRefs...),
		ProjectionDisclaimer:      referenceArchitectureValDProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValDCompatibilityGateCollection() ReferenceArchitectureCompatibilityGateCollection {
	registry := ReferenceArchitectureValBPackRegistry()
	checks := make([]ReferenceArchitectureCompatibilityGateCheck, 0, len(registry.Packs))
	for _, pack := range registry.Packs {
		checks = append(checks, referenceArchitectureValDCompatibilityGateCheckForPack(pack))
	}
	return ReferenceArchitectureCompatibilityGateCollection{
		CurrentState:         "reference_architecture_vald_compatibility_gate_collection_ready",
		CollectionID:         "reference-architecture-vald-compatibility-gate",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Checks:               checks,
		ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
	}
}

func referenceArchitectureValDFinalGateReportForFamily(family string) ReferenceArchitectureFinalGateReport {
	return ReferenceArchitectureFinalGateReport{
		CurrentState:               "reference_architecture_vald_final_gate_report_ready",
		GateID:                     "final-gate/" + family,
		BlueprintFamily:            family,
		OperationalVisibilityState: ReferenceArchitectureValDVisibilityStateActive,
		AlignmentSummaryState:      ReferenceArchitectureValDAlignmentStateActive,
		DeviationAlertState:        ReferenceArchitectureValDAlertStateActive,
		SupportBoundaryState:       ReferenceArchitectureValDSupportBoundaryStateActive,
		MigrationUpgradeState:      ReferenceArchitectureValDMigrationStateActive,
		TopologyGateState:          ReferenceArchitectureValDTopologyGateStateActive,
		SecurityBoundaryGateState:  ReferenceArchitectureValDSecurityGateStateActive,
		OperabilityGateState:       ReferenceArchitectureValDOperabilityGateStateActive,
		CompatibilityGateState:     ReferenceArchitectureValDCompatibilityGateStateActive,
		ProjectionDisclaimer:       referenceArchitectureValDProjectionDisclaimer(),
	}
}

func ReferenceArchitectureValDFinalGateCollection() ReferenceArchitectureFinalGateCollection {
	reports := make([]ReferenceArchitectureFinalGateReport, 0, len(referenceArchitectureVal0Families()))
	for _, family := range referenceArchitectureVal0Families() {
		reports = append(reports, referenceArchitectureValDFinalGateReportForFamily(family))
	}
	return ReferenceArchitectureFinalGateCollection{
		CurrentState:         "reference_architecture_vald_final_gate_collection_ready",
		CollectionID:         "reference-architecture-vald-final-gate",
		SupportedFamilies:    referenceArchitectureVal0Families(),
		Reports:              reports,
		ProjectionDisclaimer: referenceArchitectureValDProjectionDisclaimer(),
	}
}

func EvaluateReferenceArchitectureValDOperationalVisibilityReportState(report ReferenceArchitectureOperationalVisibilityReport) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		report.VisibilityReportID,
		report.Version,
		report.BlueprintFamily,
		report.BlueprintID,
		report.AlignmentStatus,
		report.ConformanceStatus,
		report.ReadinessStatus,
		report.ResilienceStatus,
		report.SupportBoundaryStatus,
		report.MigrationUpgradeStatus,
		report.ProjectionDisclaimer,
		report.CreatedAt,
		report.UpdatedAt,
	) || len(report.EvidenceRefs) == 0 {
		return ReferenceArchitectureValDVisibilityStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), report.BlueprintFamily) ||
		!containsTrimmedString(referenceArchitectureVal0ConformanceStates(), report.AlignmentStatus) ||
		!containsTrimmedString(referenceArchitectureVal0ConformanceStates(), report.ConformanceStatus) ||
		!containsTrimmedString([]string{ReferenceArchitectureValBReadinessStateActive, ReferenceArchitectureValBReadinessStatePartial, ReferenceArchitectureValBReadinessStateIncomplete, ReferenceArchitectureValBReadinessStateBlocked, ReferenceArchitectureValBReadinessStateUnknown}, report.ReadinessStatus) ||
		!containsTrimmedString([]string{ReferenceArchitectureValCStateActive, ReferenceArchitectureValCStatePartial, ReferenceArchitectureValCStateIncomplete, ReferenceArchitectureValCStateBlocked, ReferenceArchitectureValCStateUnknown}, report.ResilienceStatus) ||
		!containsTrimmedString([]string{ReferenceArchitectureValDSupportBoundaryStateActive, ReferenceArchitectureValDSupportBoundaryStatePartial, ReferenceArchitectureValDSupportBoundaryStateIncomplete, ReferenceArchitectureValDSupportBoundaryStateBlocked, ReferenceArchitectureValDSupportBoundaryStateUnknown}, report.SupportBoundaryStatus) ||
		!containsTrimmedString([]string{ReferenceArchitectureValDMigrationStateActive, ReferenceArchitectureValDMigrationStatePartial, ReferenceArchitectureValDMigrationStateIncomplete, ReferenceArchitectureValDMigrationStateBlocked, ReferenceArchitectureValDMigrationStateUnknown}, report.MigrationUpgradeStatus) ||
		!referenceArchitectureValDHasProjectionDisclaimer(report.ProjectionDisclaimer) {
		return ReferenceArchitectureValDVisibilityStatePartial
	}
	if report.CertifiedLanguagePresent || report.GuaranteedSecurityClaim || report.ProductionApprovedClaim {
		return ReferenceArchitectureValDVisibilityStateBlocked
	}
	if report.SourceValStates.Val0State != ReferenceArchitectureVal0StateActive ||
		report.SourceValStates.ValAState != ReferenceArchitectureValAStateActive ||
		report.SourceValStates.ValBState != ReferenceArchitectureValBStateActive ||
		report.SourceValStates.ValCState != ReferenceArchitectureValCStateActive ||
		report.SourceValStates.Point5State != IntelligenceCalibrationPoint5StatePass ||
		!referenceArchitectureValDPoint5DependencyHealthy(report.SourceValStates.Point5DependencyState) ||
		report.Point6State != ReferenceArchitecturePoint6StateNotComplete {
		return ReferenceArchitectureValDVisibilityStatePartial
	}
	allFresh, stale, ok := referenceArchitectureValBEvidenceValid(report.EvidenceRefs)
	if !ok || stale || !allFresh {
		return ReferenceArchitectureValDVisibilityStatePartial
	}
	return ReferenceArchitectureValDVisibilityStateActive
}

func EvaluateReferenceArchitectureValDOperationalVisibilityCollectionState(collection ReferenceArchitectureOperationalVisibilityCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Reports) == 0 {
		return ReferenceArchitectureValDVisibilityStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValDHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Reports) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValDVisibilityStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, report := range collection.Reports {
		if !referenceArchitectureValDRegisterFamily(seenFamilies, report.BlueprintFamily) {
			return ReferenceArchitectureValDVisibilityStatePartial
		}
		if EvaluateReferenceArchitectureValDOperationalVisibilityReportState(report) != ReferenceArchitectureValDVisibilityStateActive {
			return ReferenceArchitectureValDVisibilityStatePartial
		}
	}
	if !referenceArchitectureValDSeenFamiliesMatchSupported(seenFamilies) {
		return ReferenceArchitectureValDVisibilityStatePartial
	}
	return ReferenceArchitectureValDVisibilityStateActive
}

func EvaluateReferenceArchitectureValDAlignmentSummaryState(summary ReferenceArchitectureBlueprintAlignmentSummary) string {
	if !referenceArchitectureValBRequiredRefsPresent(summary.SummaryID, summary.BlueprintFamily, summary.AlignmentStatus, summary.ProjectionDisclaimer) {
		return ReferenceArchitectureValDAlignmentStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), summary.BlueprintFamily) ||
		!containsTrimmedString(referenceArchitectureVal0ConformanceStates(), summary.AlignmentStatus) ||
		!referenceArchitectureValDHasProjectionDisclaimer(summary.ProjectionDisclaimer) {
		return ReferenceArchitectureValDAlignmentStatePartial
	}
	if !summary.RedactionKeepsBlockingVisible || len(summary.BlockingDeviations) > 0 || len(summary.StaleEvidenceRefs) > 0 ||
		len(summary.UnsupportedEnvironmentConditions) > 0 || len(summary.SupportBoundaryGaps) > 0 {
		return ReferenceArchitectureValDAlignmentStatePartial
	}
	if summary.Val0State != ReferenceArchitectureVal0StateActive ||
		summary.ValAState != ReferenceArchitectureValAStateActive ||
		summary.ValBState != ReferenceArchitectureValBStateActive ||
		summary.ValCState != ReferenceArchitectureValCStateActive ||
		summary.AlignmentStatus != ReferenceArchitectureConformanceMatched {
		return ReferenceArchitectureValDAlignmentStatePartial
	}
	return ReferenceArchitectureValDAlignmentStateActive
}

func EvaluateReferenceArchitectureValDAlignmentSummaryCollectionState(collection ReferenceArchitectureBlueprintAlignmentCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Summaries) == 0 {
		return ReferenceArchitectureValDAlignmentStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValDHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Summaries) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValDAlignmentStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, summary := range collection.Summaries {
		if !referenceArchitectureValDRegisterFamily(seenFamilies, summary.BlueprintFamily) {
			return ReferenceArchitectureValDAlignmentStatePartial
		}
		if EvaluateReferenceArchitectureValDAlignmentSummaryState(summary) != ReferenceArchitectureValDAlignmentStateActive {
			return ReferenceArchitectureValDAlignmentStatePartial
		}
	}
	if !referenceArchitectureValDSeenFamiliesMatchSupported(seenFamilies) {
		return ReferenceArchitectureValDAlignmentStatePartial
	}
	return ReferenceArchitectureValDAlignmentStateActive
}

func EvaluateReferenceArchitectureValDDeviationAlertState(alert ReferenceArchitectureDeviationAlert) string {
	if !referenceArchitectureValBRequiredRefsPresent(alert.AlertID, alert.BlueprintFamily, alert.SourceLayer, alert.DeviationCategory, alert.Severity, alert.AffectedScope, alert.OperatorActionRequired, alert.SupportBoundaryRef, alert.Timestamp) {
		return ReferenceArchitectureValDAlertStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), alert.BlueprintFamily) ||
		!containsTrimmedString(referenceArchitectureValDSourceLayers(), alert.SourceLayer) ||
		!containsTrimmedString(referenceArchitectureValDAlertCategories(), alert.DeviationCategory) ||
		!containsTrimmedString(referenceArchitectureValCScenarioSeverities(), alert.Severity) {
		return ReferenceArchitectureValDAlertStatePartial
	}
	if _, ok := referenceArchitectureVal0ParseTimestamp(alert.Timestamp); !ok {
		return ReferenceArchitectureValDAlertStatePartial
	}
	if alert.BlocksAlignment && strings.TrimSpace(alert.EvidenceRef) == "" {
		return ReferenceArchitectureValDAlertStateIncomplete
	}
	if strings.TrimSpace(alert.FreshnessState) != IntelligenceCalibrationFreshnessFresh {
		return ReferenceArchitectureValDAlertStatePartial
	}
	if alert.AdvisoryOnly && alert.BlocksAlignment {
		return ReferenceArchitectureValDAlertStatePartial
	}
	if alert.DeviationCategory == ReferenceArchitectureValDDeviationOverclaimDetected && alert.BlocksAlignment {
		return ReferenceArchitectureValDAlertStateBlocked
	}
	return ReferenceArchitectureValDAlertStateActive
}

func EvaluateReferenceArchitectureValDDeviationAlertReportState(report ReferenceArchitectureDeviationAlertReport) string {
	if !referenceArchitectureValBRequiredRefsPresent(report.ReportID, report.BlueprintFamily, report.ProjectionDisclaimer) {
		return ReferenceArchitectureValDAlertStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), report.BlueprintFamily) || !referenceArchitectureValDHasProjectionDisclaimer(report.ProjectionDisclaimer) {
		return ReferenceArchitectureValDAlertStatePartial
	}
	for _, alert := range report.Alerts {
		state := EvaluateReferenceArchitectureValDDeviationAlertState(alert)
		if state != ReferenceArchitectureValDAlertStateActive {
			return state
		}
	}
	return ReferenceArchitectureValDAlertStateActive
}

func EvaluateReferenceArchitectureValDDeviationAlertCollectionState(collection ReferenceArchitectureDeviationAlertCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Reports) == 0 {
		return ReferenceArchitectureValDAlertStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(collection.SupportedCategories, referenceArchitectureValDAlertCategories()...) ||
		!containsExactTrimmedStringSet(collection.SupportedSeverities, referenceArchitectureValCScenarioSeverities()...) ||
		!containsExactTrimmedStringSet(collection.SupportedSourceLayers, referenceArchitectureValDSourceLayers()...) ||
		!referenceArchitectureValDHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Reports) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValDAlertStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, report := range collection.Reports {
		if !referenceArchitectureValDRegisterFamily(seenFamilies, report.BlueprintFamily) {
			return ReferenceArchitectureValDAlertStatePartial
		}
		state := EvaluateReferenceArchitectureValDDeviationAlertReportState(report)
		if state != ReferenceArchitectureValDAlertStateActive {
			return state
		}
	}
	if !referenceArchitectureValDSeenFamiliesMatchSupported(seenFamilies) {
		return ReferenceArchitectureValDAlertStatePartial
	}
	return ReferenceArchitectureValDAlertStateActive
}

func EvaluateReferenceArchitectureValDSupportBoundaryViewState(view ReferenceArchitectureSupportBoundaryView) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		view.ViewID,
		view.BlueprintFamily,
		view.SupportedEnvironmentScope,
		view.OperatorResponsibility,
		view.VerifierAuditorBoundary,
		view.EvidenceExportBoundary,
		view.EscalationGuidanceRef,
		view.SupportBoundaryRef,
		view.ProjectionDisclaimer,
	) || len(view.UnsupportedConditions) == 0 || len(view.DegradedSupportConditions) == 0 {
		return ReferenceArchitectureValDSupportBoundaryStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), view.BlueprintFamily) || !referenceArchitectureValDHasProjectionDisclaimer(view.ProjectionDisclaimer) {
		return ReferenceArchitectureValDSupportBoundaryStatePartial
	}
	if !view.RedactionKeepsUnsupportedVisible || view.PartnerCanonicalAuthority {
		return ReferenceArchitectureValDSupportBoundaryStatePartial
	}
	return ReferenceArchitectureValDSupportBoundaryStateActive
}

func EvaluateReferenceArchitectureValDSupportBoundaryCollectionState(collection ReferenceArchitectureSupportBoundaryCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Views) == 0 {
		return ReferenceArchitectureValDSupportBoundaryStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValDHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Views) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValDSupportBoundaryStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, view := range collection.Views {
		if !referenceArchitectureValDRegisterFamily(seenFamilies, view.BlueprintFamily) {
			return ReferenceArchitectureValDSupportBoundaryStatePartial
		}
		if EvaluateReferenceArchitectureValDSupportBoundaryViewState(view) != ReferenceArchitectureValDSupportBoundaryStateActive {
			return ReferenceArchitectureValDSupportBoundaryStatePartial
		}
	}
	if !referenceArchitectureValDSeenFamiliesMatchSupported(seenFamilies) {
		return ReferenceArchitectureValDSupportBoundaryStatePartial
	}
	return ReferenceArchitectureValDSupportBoundaryStateActive
}

func EvaluateReferenceArchitectureValDMigrationVisibilityState(view ReferenceArchitectureMigrationUpgradeVisibility) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		view.VisibilityID,
		view.BlueprintFamily,
		view.CurrentBlueprintVersion,
		view.TargetBlueprintVersion,
		view.CompatibilityState,
		view.DeprecationState,
		view.MigrationPathRef,
		view.RollbackBoundaryRef,
		view.ProjectionDisclaimer,
	) || len(view.RequiredValidationRefs) == 0 || len(view.EvidenceRefs) == 0 {
		return ReferenceArchitectureValDMigrationStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), view.BlueprintFamily) ||
		!containsTrimmedString(referenceArchitectureVal0CompatibilityStates(), view.CompatibilityState) ||
		!containsTrimmedString(referenceArchitectureVal0LifecycleStates(), view.DeprecationState) ||
		!referenceArchitectureValDHasProjectionDisclaimer(view.ProjectionDisclaimer) {
		return ReferenceArchitectureValDMigrationStatePartial
	}
	if view.ExecutesMigration {
		return ReferenceArchitectureValDMigrationStateBlocked
	}
	allFresh, stale, ok := referenceArchitectureValBEvidenceValid(view.EvidenceRefs)
	if !ok || stale || !allFresh || view.StaleIndicator {
		return ReferenceArchitectureValDMigrationStatePartial
	}
	if view.CompatibilityState == ReferenceArchitectureCompatibilityUnsupported || view.DeprecationState == ReferenceArchitectureLifecycleUnsupported {
		return ReferenceArchitectureValDMigrationStateBlocked
	}
	if view.CompatibilityState == ReferenceArchitectureCompatibilityDeprecated ||
		view.CompatibilityState == ReferenceArchitectureCompatibilitySuperseded ||
		view.DeprecationState == ReferenceArchitectureLifecycleDeprecated ||
		view.DeprecationState == ReferenceArchitectureLifecycleSuperseded ||
		view.SupersededIndicator {
		return ReferenceArchitectureValDMigrationStatePartial
	}
	return ReferenceArchitectureValDMigrationStateActive
}

func EvaluateReferenceArchitectureValDMigrationUpgradeCollectionState(collection ReferenceArchitectureMigrationUpgradeCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Views) == 0 {
		return ReferenceArchitectureValDMigrationStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValDHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Views) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValDMigrationStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, view := range collection.Views {
		if !referenceArchitectureValDRegisterFamily(seenFamilies, view.BlueprintFamily) {
			return ReferenceArchitectureValDMigrationStatePartial
		}
		state := EvaluateReferenceArchitectureValDMigrationVisibilityState(view)
		if state != ReferenceArchitectureValDMigrationStateActive {
			return state
		}
	}
	if !referenceArchitectureValDSeenFamiliesMatchSupported(seenFamilies) {
		return ReferenceArchitectureValDMigrationStatePartial
	}
	return ReferenceArchitectureValDMigrationStateActive
}

func EvaluateReferenceArchitectureValDTopologyGateState(check ReferenceArchitectureTopologyGateCheck) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		check.CheckID,
		check.BlueprintFamily,
		check.DeploymentTopology,
		check.TrustAnchorMode,
		check.AuditPathMode,
		check.ConnectivityMode,
		check.ProjectionDisclaimer,
	) || len(check.EvidenceRefs) == 0 {
		return ReferenceArchitectureValDTopologyGateStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), check.BlueprintFamily) ||
		!containsTrimmedString(referenceArchitectureVal0SupportedTopologies(), check.DeploymentTopology) ||
		!containsTrimmedString(referenceArchitectureVal0SupportedTrustAnchors(), check.TrustAnchorMode) ||
		!containsTrimmedString(referenceArchitectureVal0SupportedAuditPaths(), check.AuditPathMode) ||
		!containsTrimmedString(referenceArchitectureVal0SupportedConnectivityModes(), check.ConnectivityMode) ||
		!referenceArchitectureValDHasProjectionDisclaimer(check.ProjectionDisclaimer) {
		return ReferenceArchitectureValDTopologyGateStatePartial
	}
	if !check.RedactionKeepsMismatchVisible || !check.SupportedTopology || !check.TrustAnchorTopologyCompatible ||
		!check.AuditPathTopologyCompatible || !check.OfflineTopologyCompatible ||
		len(check.UnsupportedConditions) > 0 {
		return ReferenceArchitectureValDTopologyGateStatePartial
	}
	if check.ControlDataPlaneSeparationRequired && !check.ControlDataPlaneSeparationSatisfied {
		return ReferenceArchitectureValDTopologyGateStatePartial
	}
	allFresh, stale, ok := referenceArchitectureValBEvidenceValid(check.EvidenceRefs)
	if !ok || stale || !allFresh {
		return ReferenceArchitectureValDTopologyGateStatePartial
	}
	return ReferenceArchitectureValDTopologyGateStateActive
}

func EvaluateReferenceArchitectureValDTopologyGateCollectionState(collection ReferenceArchitectureTopologyGateCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Checks) == 0 {
		return ReferenceArchitectureValDTopologyGateStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValDHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Checks) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValDTopologyGateStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, check := range collection.Checks {
		if !referenceArchitectureValDRegisterFamily(seenFamilies, check.BlueprintFamily) {
			return ReferenceArchitectureValDTopologyGateStatePartial
		}
		if EvaluateReferenceArchitectureValDTopologyGateState(check) != ReferenceArchitectureValDTopologyGateStateActive {
			return ReferenceArchitectureValDTopologyGateStatePartial
		}
	}
	if !referenceArchitectureValDSeenFamiliesMatchSupported(seenFamilies) {
		return ReferenceArchitectureValDTopologyGateStatePartial
	}
	return ReferenceArchitectureValDTopologyGateStateActive
}

func EvaluateReferenceArchitectureValDSecurityBoundaryGateState(check ReferenceArchitectureSecurityBoundaryGateCheck) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		check.CheckID,
		check.BlueprintFamily,
		check.TrustAnchorBoundary,
		check.SigningCustodyBoundary,
		check.EvidenceStorageBoundary,
		check.PolicyDistributionBoundary,
		check.OperatorAccessBoundary,
		check.PartnerVerifierAuditorBoundary,
		check.RedactionExportBoundary,
		check.NoShadowTruthBoundary,
		check.ProjectionDisclaimer,
	) || len(check.EvidenceRefs) == 0 {
		return ReferenceArchitectureValDSecurityGateStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), check.BlueprintFamily) || !referenceArchitectureValDHasProjectionDisclaimer(check.ProjectionDisclaimer) {
		return ReferenceArchitectureValDSecurityGateStatePartial
	}
	if !check.MutationAuthorityBlocked || !check.ApprovalAuthorityBlocked {
		return ReferenceArchitectureValDSecurityGateStateBlocked
	}
	if strings.TrimSpace(check.NoShadowTruthBoundary) == "" {
		return ReferenceArchitectureValDSecurityGateStatePartial
	}
	allFresh, stale, ok := referenceArchitectureValBEvidenceValid(check.EvidenceRefs)
	if !ok || stale || !allFresh {
		return ReferenceArchitectureValDSecurityGateStatePartial
	}
	return ReferenceArchitectureValDSecurityGateStateActive
}

func EvaluateReferenceArchitectureValDSecurityBoundaryCollectionState(collection ReferenceArchitectureSecurityBoundaryCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Checks) == 0 {
		return ReferenceArchitectureValDSecurityGateStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValDHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Checks) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValDSecurityGateStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, check := range collection.Checks {
		if !referenceArchitectureValDRegisterFamily(seenFamilies, check.BlueprintFamily) {
			return ReferenceArchitectureValDSecurityGateStatePartial
		}
		state := EvaluateReferenceArchitectureValDSecurityBoundaryGateState(check)
		if state != ReferenceArchitectureValDSecurityGateStateActive {
			return state
		}
	}
	if !referenceArchitectureValDSeenFamiliesMatchSupported(seenFamilies) {
		return ReferenceArchitectureValDSecurityGateStatePartial
	}
	return ReferenceArchitectureValDSecurityGateStateActive
}

func EvaluateReferenceArchitectureValDOperabilityGateState(check ReferenceArchitectureOperabilityGateCheck) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		check.CheckID,
		check.BlueprintFamily,
		check.ReadinessState,
		check.ResilienceState,
		check.RecoveryExpectationState,
		check.AuditPathState,
		check.ControlPlaneState,
		check.SupportBoundaryState,
		check.OperatorActionGuidance,
		check.ProjectionDisclaimer,
	) || len(check.EvidenceRefs) == 0 {
		return ReferenceArchitectureValDOperabilityGateStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), check.BlueprintFamily) || !referenceArchitectureValDHasProjectionDisclaimer(check.ProjectionDisclaimer) {
		return ReferenceArchitectureValDOperabilityGateStatePartial
	}
	if check.ReadinessState != ReferenceArchitectureValBReadinessStateActive ||
		check.ResilienceState != ReferenceArchitectureValCStateActive ||
		check.RecoveryExpectationState != ReferenceArchitectureValCRecoveryExpectationStateActive ||
		check.AuditPathState != ReferenceArchitectureValCAuditPathStateActive ||
		check.ControlPlaneState != ReferenceArchitectureValCControlPlaneStateActive ||
		check.SupportBoundaryState != ReferenceArchitectureValDSupportBoundaryStateActive ||
		!check.DegradedStateVisible {
		return ReferenceArchitectureValDOperabilityGateStatePartial
	}
	allFresh, stale, ok := referenceArchitectureValBEvidenceValid(check.EvidenceRefs)
	if !ok || stale || !allFresh {
		return ReferenceArchitectureValDOperabilityGateStatePartial
	}
	return ReferenceArchitectureValDOperabilityGateStateActive
}

func EvaluateReferenceArchitectureValDOperabilityGateCollectionState(collection ReferenceArchitectureOperabilityGateCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Checks) == 0 {
		return ReferenceArchitectureValDOperabilityGateStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValDHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Checks) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValDOperabilityGateStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, check := range collection.Checks {
		if !referenceArchitectureValDRegisterFamily(seenFamilies, check.BlueprintFamily) {
			return ReferenceArchitectureValDOperabilityGateStatePartial
		}
		if EvaluateReferenceArchitectureValDOperabilityGateState(check) != ReferenceArchitectureValDOperabilityGateStateActive {
			return ReferenceArchitectureValDOperabilityGateStatePartial
		}
	}
	if !referenceArchitectureValDSeenFamiliesMatchSupported(seenFamilies) {
		return ReferenceArchitectureValDOperabilityGateStatePartial
	}
	return ReferenceArchitectureValDOperabilityGateStateActive
}

func EvaluateReferenceArchitectureValDCompatibilityGateState(check ReferenceArchitectureCompatibilityGateCheck) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		check.CheckID,
		check.BlueprintFamily,
		check.LifecycleState,
		check.CompatibilityState,
		check.MigrationVisibilityRef,
		check.ProjectionDisclaimer,
	) || len(check.ValidationRequirementRefs) == 0 || len(check.EvidenceRefs) == 0 {
		return ReferenceArchitectureValDCompatibilityGateStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), check.BlueprintFamily) ||
		!containsTrimmedString(referenceArchitectureVal0LifecycleStates(), check.LifecycleState) ||
		!containsTrimmedString(referenceArchitectureVal0CompatibilityStates(), check.CompatibilityState) ||
		!referenceArchitectureValDHasProjectionDisclaimer(check.ProjectionDisclaimer) {
		return ReferenceArchitectureValDCompatibilityGateStatePartial
	}
	if check.UniversalSupportClaim {
		return ReferenceArchitectureValDCompatibilityGateStateBlocked
	}
	allFresh, stale, ok := referenceArchitectureValBEvidenceValid(check.EvidenceRefs)
	if !ok || stale || !allFresh {
		return ReferenceArchitectureValDCompatibilityGateStatePartial
	}
	if check.CompatibilityState == ReferenceArchitectureCompatibilityUnsupported || check.LifecycleState == ReferenceArchitectureLifecycleUnsupported {
		return ReferenceArchitectureValDCompatibilityGateStateBlocked
	}
	if check.CompatibilityState == ReferenceArchitectureCompatibilityDeprecated ||
		check.CompatibilityState == ReferenceArchitectureCompatibilitySuperseded ||
		check.LifecycleState == ReferenceArchitectureLifecycleDeprecated ||
		check.LifecycleState == ReferenceArchitectureLifecycleSuperseded {
		return ReferenceArchitectureValDCompatibilityGateStatePartial
	}
	return ReferenceArchitectureValDCompatibilityGateStateActive
}

func EvaluateReferenceArchitectureValDCompatibilityGateCollectionState(collection ReferenceArchitectureCompatibilityGateCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Checks) == 0 {
		return ReferenceArchitectureValDCompatibilityGateStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValDHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Checks) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValDCompatibilityGateStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, check := range collection.Checks {
		if !referenceArchitectureValDRegisterFamily(seenFamilies, check.BlueprintFamily) {
			return ReferenceArchitectureValDCompatibilityGateStatePartial
		}
		state := EvaluateReferenceArchitectureValDCompatibilityGateState(check)
		if state != ReferenceArchitectureValDCompatibilityGateStateActive {
			return state
		}
	}
	if !referenceArchitectureValDSeenFamiliesMatchSupported(seenFamilies) {
		return ReferenceArchitectureValDCompatibilityGateStatePartial
	}
	return ReferenceArchitectureValDCompatibilityGateStateActive
}

func EvaluateReferenceArchitectureValDFinalGateReportState(report ReferenceArchitectureFinalGateReport) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		report.GateID,
		report.BlueprintFamily,
		report.OperationalVisibilityState,
		report.AlignmentSummaryState,
		report.DeviationAlertState,
		report.SupportBoundaryState,
		report.MigrationUpgradeState,
		report.TopologyGateState,
		report.SecurityBoundaryGateState,
		report.OperabilityGateState,
		report.CompatibilityGateState,
		report.ProjectionDisclaimer,
	) {
		return ReferenceArchitectureValDFinalGateStateIncomplete
	}
	if !containsTrimmedString(referenceArchitectureVal0Families(), report.BlueprintFamily) || !referenceArchitectureValDHasProjectionDisclaimer(report.ProjectionDisclaimer) {
		return ReferenceArchitectureValDFinalGateStatePartial
	}
	if len(report.BlockingReasons) > 0 {
		return ReferenceArchitectureValDFinalGateStatePartial
	}
	if !referenceArchitectureValDFinalGateHasSupportedComponentStates(report) {
		return ReferenceArchitectureValDFinalGateStatePartial
	}
	if report.OperationalVisibilityState != ReferenceArchitectureValDVisibilityStateActive ||
		report.AlignmentSummaryState != ReferenceArchitectureValDAlignmentStateActive ||
		report.DeviationAlertState != ReferenceArchitectureValDAlertStateActive ||
		report.SupportBoundaryState != ReferenceArchitectureValDSupportBoundaryStateActive ||
		report.MigrationUpgradeState != ReferenceArchitectureValDMigrationStateActive ||
		report.TopologyGateState != ReferenceArchitectureValDTopologyGateStateActive ||
		report.SecurityBoundaryGateState != ReferenceArchitectureValDSecurityGateStateActive ||
		report.OperabilityGateState != ReferenceArchitectureValDOperabilityGateStateActive ||
		report.CompatibilityGateState != ReferenceArchitectureValDCompatibilityGateStateActive {
		return ReferenceArchitectureValDFinalGateStatePartial
	}
	return ReferenceArchitectureValDFinalGateStateActive
}

func EvaluateReferenceArchitectureValDFinalGateCollectionState(collection ReferenceArchitectureFinalGateCollection) string {
	if strings.TrimSpace(collection.CollectionID) == "" || strings.TrimSpace(collection.ProjectionDisclaimer) == "" || len(collection.Reports) == 0 {
		return ReferenceArchitectureValDFinalGateStateIncomplete
	}
	if !containsExactTrimmedStringSet(collection.SupportedFamilies, referenceArchitectureVal0Families()...) ||
		!referenceArchitectureValDHasProjectionDisclaimer(collection.ProjectionDisclaimer) ||
		len(collection.Reports) != len(referenceArchitectureVal0Families()) {
		return ReferenceArchitectureValDFinalGateStatePartial
	}
	seenFamilies := map[string]struct{}{}
	for _, report := range collection.Reports {
		if !referenceArchitectureValDRegisterFamily(seenFamilies, report.BlueprintFamily) {
			return ReferenceArchitectureValDFinalGateStatePartial
		}
		if EvaluateReferenceArchitectureValDFinalGateReportState(report) != ReferenceArchitectureValDFinalGateStateActive {
			return ReferenceArchitectureValDFinalGateStatePartial
		}
	}
	if !referenceArchitectureValDSeenFamiliesMatchSupported(seenFamilies) {
		return ReferenceArchitectureValDFinalGateStatePartial
	}
	return ReferenceArchitectureValDFinalGateStateActive
}

func referenceArchitectureValDRequiresPriorStates(point5State, point5DependencyState, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, valCCurrentState, valCState, point6State string) bool {
	return strings.TrimSpace(point5State) == IntelligenceCalibrationPoint5StatePass &&
		referenceArchitectureValDPoint5DependencyHealthy(point5DependencyState) &&
		strings.TrimSpace(val0CurrentState) == ReferenceArchitectureVal0StateActive &&
		strings.TrimSpace(val0State) == ReferenceArchitectureVal0StateActive &&
		strings.TrimSpace(valACurrentState) == ReferenceArchitectureValAStateActive &&
		strings.TrimSpace(valAState) == ReferenceArchitectureValAStateActive &&
		strings.TrimSpace(valBCurrentState) == ReferenceArchitectureValBStateActive &&
		strings.TrimSpace(valBState) == ReferenceArchitectureValBStateActive &&
		strings.TrimSpace(valCCurrentState) == ReferenceArchitectureValCStateActive &&
		strings.TrimSpace(valCState) == ReferenceArchitectureValCStateActive &&
		strings.TrimSpace(point6State) == ReferenceArchitecturePoint6StateNotComplete
}

func EvaluateReferenceArchitectureValDState(
	point5State, point5DependencyState, val0CurrentState, val0State, valACurrentState, valAState,
	valBCurrentState, valBState, valCCurrentState, valCState, point6State,
	visibilityState, alignmentState, alertState, supportBoundaryState, migrationState, topologyGateState,
	securityGateState, operabilityGateState, compatibilityGateState, finalGateState string,
) string {
	if !referenceArchitectureValDRequiresPriorStates(point5State, point5DependencyState, val0CurrentState, val0State, valACurrentState, valAState, valBCurrentState, valBState, valCCurrentState, valCState, point6State) {
		return ReferenceArchitectureValDStateBlocked
	}
	componentStates := []string{
		visibilityState,
		alignmentState,
		alertState,
		supportBoundaryState,
		migrationState,
		topologyGateState,
		securityGateState,
		operabilityGateState,
		compatibilityGateState,
		finalGateState,
	}
	for _, state := range componentStates {
		if strings.TrimSpace(state) == "" {
			return ReferenceArchitectureValDStateIncomplete
		}
	}
	allActive := visibilityState == ReferenceArchitectureValDVisibilityStateActive &&
		alignmentState == ReferenceArchitectureValDAlignmentStateActive &&
		alertState == ReferenceArchitectureValDAlertStateActive &&
		supportBoundaryState == ReferenceArchitectureValDSupportBoundaryStateActive &&
		migrationState == ReferenceArchitectureValDMigrationStateActive &&
		topologyGateState == ReferenceArchitectureValDTopologyGateStateActive &&
		securityGateState == ReferenceArchitectureValDSecurityGateStateActive &&
		operabilityGateState == ReferenceArchitectureValDOperabilityGateStateActive &&
		compatibilityGateState == ReferenceArchitectureValDCompatibilityGateStateActive &&
		finalGateState == ReferenceArchitectureValDFinalGateStateActive
	if allActive {
		return ReferenceArchitectureValDStateActive
	}
	for _, state := range componentStates {
		switch strings.TrimSpace(state) {
		case ReferenceArchitectureValDVisibilityStateBlocked,
			ReferenceArchitectureValDAlignmentStateBlocked,
			ReferenceArchitectureValDAlertStateBlocked,
			ReferenceArchitectureValDSupportBoundaryStateBlocked,
			ReferenceArchitectureValDMigrationStateBlocked,
			ReferenceArchitectureValDTopologyGateStateBlocked,
			ReferenceArchitectureValDSecurityGateStateBlocked,
			ReferenceArchitectureValDOperabilityGateStateBlocked,
			ReferenceArchitectureValDCompatibilityGateStateBlocked,
			ReferenceArchitectureValDFinalGateStateBlocked:
			return ReferenceArchitectureValDStateBlocked
		case ReferenceArchitectureValDVisibilityStateIncomplete,
			ReferenceArchitectureValDAlignmentStateIncomplete,
			ReferenceArchitectureValDAlertStateIncomplete,
			ReferenceArchitectureValDSupportBoundaryStateIncomplete,
			ReferenceArchitectureValDMigrationStateIncomplete,
			ReferenceArchitectureValDTopologyGateStateIncomplete,
			ReferenceArchitectureValDSecurityGateStateIncomplete,
			ReferenceArchitectureValDOperabilityGateStateIncomplete,
			ReferenceArchitectureValDCompatibilityGateStateIncomplete,
			ReferenceArchitectureValDFinalGateStateIncomplete:
			return ReferenceArchitectureValDStateIncomplete
		case ReferenceArchitectureValDVisibilityStateUnknown,
			ReferenceArchitectureValDAlignmentStateUnknown,
			ReferenceArchitectureValDAlertStateUnknown,
			ReferenceArchitectureValDSupportBoundaryStateUnknown,
			ReferenceArchitectureValDMigrationStateUnknown,
			ReferenceArchitectureValDTopologyGateStateUnknown,
			ReferenceArchitectureValDSecurityGateStateUnknown,
			ReferenceArchitectureValDOperabilityGateStateUnknown,
			ReferenceArchitectureValDCompatibilityGateStateUnknown,
			ReferenceArchitectureValDFinalGateStateUnknown:
			return ReferenceArchitectureValDStateUnknown
		}
	}
	return ReferenceArchitectureValDStatePartial
}

func referenceArchitectureValDProofSurfaceRefs() []string {
	return []string{
		"/v1/reference-architecture/val0/proofs",
		"/v1/reference-architecture/vala/proofs",
		"/v1/reference-architecture/valb/proofs",
		"/v1/reference-architecture/valc/proofs",
		"/v1/reference-architecture/vald/operational-visibility",
		"/v1/reference-architecture/vald/alignment-summary",
		"/v1/reference-architecture/vald/deviation-alerts",
		"/v1/reference-architecture/vald/support-boundaries",
		"/v1/reference-architecture/vald/migration-upgrade",
		"/v1/reference-architecture/vald/topology-gate",
		"/v1/reference-architecture/vald/security-boundary-gate",
		"/v1/reference-architecture/vald/operability-gate",
		"/v1/reference-architecture/vald/compatibility-gate",
		"/v1/reference-architecture/vald/final-gate",
		"/v1/reference-architecture/vald/proofs",
	}
}

func EvaluateReferenceArchitectureValDProofsState(valDState, point5DependencyState, point6State string, supportedFamilies, surfaceRefs, evidenceRefs, limitations []string, projectionDisclaimer string) string {
	baseState := strings.TrimSpace(valDState)
	if !containsExactTrimmedStringSet(supportedFamilies, referenceArchitectureVal0Families()...) ||
		!containsExactTrimmedStringSet(surfaceRefs, referenceArchitectureValDProofSurfaceRefs()...) ||
		len(evidenceRefs) < 14 ||
		len(limitations) == 0 ||
		!referenceArchitectureValDHasProjectionDisclaimer(projectionDisclaimer) {
		if baseState == ReferenceArchitectureValDStateActive {
			return ReferenceArchitectureValDStatePartial
		}
		return baseState
	}
	if baseState == ReferenceArchitectureValDStateActive &&
		(!referenceArchitectureValDPoint5DependencyHealthy(point5DependencyState) || strings.TrimSpace(point6State) != ReferenceArchitecturePoint6StateNotComplete) {
		return ReferenceArchitectureValDStatePartial
	}
	return baseState
}

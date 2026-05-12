package operability

import "strings"

const (
	DeploymentMultiTenantValBStateActive  = "deployment_multi_tenant_valb_active"
	DeploymentMultiTenantValBStateBlocked = "deployment_multi_tenant_valb_blocked"

	DeploymentMultiTenantValBDependencyStateActive  = "deployment_multi_tenant_valb_dependency_active"
	DeploymentMultiTenantValBDependencyStateBlocked = "deployment_multi_tenant_valb_dependency_blocked"

	DeploymentMultiTenantValBTenantIsolationStateActive  = "deployment_multi_tenant_valb_tenant_isolation_active"
	DeploymentMultiTenantValBTenantIsolationStateBlocked = "deployment_multi_tenant_valb_tenant_isolation_blocked"

	DeploymentMultiTenantValBDataResidencyStateActive  = "deployment_multi_tenant_valb_data_residency_active"
	DeploymentMultiTenantValBDataResidencyStateBlocked = "deployment_multi_tenant_valb_data_residency_blocked"

	DeploymentMultiTenantValBTenantLifecycleStateActive  = "deployment_multi_tenant_valb_tenant_lifecycle_active"
	DeploymentMultiTenantValBTenantLifecycleStateBlocked = "deployment_multi_tenant_valb_tenant_lifecycle_blocked"

	DeploymentMultiTenantValBFairShareQuotaStateActive  = "deployment_multi_tenant_valb_fair_share_quota_active"
	DeploymentMultiTenantValBFairShareQuotaStateBlocked = "deployment_multi_tenant_valb_fair_share_quota_blocked"

	DeploymentMultiTenantValBNoOverclaimStateActive  = "deployment_multi_tenant_valb_no_overclaim_active"
	DeploymentMultiTenantValBNoOverclaimStateBlocked = "deployment_multi_tenant_valb_no_overclaim_blocked"

	DeploymentMultiTenantValBClosureBlockerStateActive   = "deployment_multi_tenant_valb_closure_blocker_active"
	DeploymentMultiTenantValBClosureBlockerStateCleanup  = "deployment_multi_tenant_valb_closure_blocker_cleanup"
	DeploymentMultiTenantValBClosureBlockerStateAdvisory = "deployment_multi_tenant_valb_closure_blocker_advisory"
	DeploymentMultiTenantValBClosureBlockerStateBlocked  = "deployment_multi_tenant_valb_closure_blocker_blocked"

	DeploymentMultiTenantValBBlockerLevelCLB0 = "CL-B0"
	DeploymentMultiTenantValBBlockerLevelCLB1 = "CL-B1"
	DeploymentMultiTenantValBBlockerLevelCLB2 = "CL-B2"
	DeploymentMultiTenantValBBlockerLevelCLB3 = "CL-B3"

	DeploymentMultiTenantValBClosureSurfaceTenantIsolation = "tenant_isolation"
	DeploymentMultiTenantValBClosureSurfaceDataResidency   = "data_residency"
	DeploymentMultiTenantValBClosureSurfaceTenantLifecycle = "tenant_lifecycle"
	DeploymentMultiTenantValBClosureSurfaceFairShare       = "fair_share"
	DeploymentMultiTenantValBClosureSurfaceNoOverclaim     = "no_overclaim"
	DeploymentMultiTenantValBClosureSurfaceCleanRoomIP     = "clean_room_ip"
)

type DeploymentMultiTenantValBDependencySnapshot struct {
	ValACurrentState                 string `json:"vala_current_state"`
	ValADependencyState              string `json:"vala_dependency_state"`
	ValADeploymentProfileMatrixState string `json:"vala_deployment_profile_matrix_state"`
	ValAPreflightGateState           string `json:"vala_preflight_gate_state"`
	ValAIdentityBootstrapState       string `json:"vala_identity_bootstrap_state"`
	ValAAirGappedEvidenceBundleState string `json:"vala_air_gapped_evidence_bundle_state"`
	ValANoOverclaimState             string `json:"vala_no_overclaim_state"`
	ValAPassBlockerState             string `json:"vala_pass_blocker_state"`
	Point10State                     string `json:"point_10_state"`
	ProjectionDisclaimer             string `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValBTenantIsolationTestPack struct {
	CurrentState                                string   `json:"current_state"`
	EvidenceRefs                                []string `json:"evidence_refs,omitempty"`
	FreshnessState                              string   `json:"freshness_state"`
	TenantScope                                 string   `json:"tenant_scope"`
	TenantIsolationEvidenceBacked               bool     `json:"tenant_isolation_evidence_backed"`
	CrossTenantAuditLeakageTestPresent          bool     `json:"cross_tenant_audit_leakage_test_present"`
	CrossTenantEvidenceLeakageTestPresent       bool     `json:"cross_tenant_evidence_leakage_test_present"`
	CrossTenantExportLeakageTestPresent         bool     `json:"cross_tenant_export_leakage_test_present"`
	CrossTenantCredentialLeakageTestPresent     bool     `json:"cross_tenant_credential_leakage_test_present"`
	SupportOperatorAccessLeakageTestPresent     bool     `json:"support_operator_access_leakage_test_present"`
	TenantBoundaryValidationTestPresent         bool     `json:"tenant_boundary_validation_test_present"`
	TenantScopeNegativeTestPresent              bool     `json:"tenant_scope_negative_test_present"`
	RawCrossTenantEvidenceSharingNegativeTest   bool     `json:"raw_cross_tenant_evidence_sharing_negative_test_present"`
	TenantNamespaceIsolationEvidence            string   `json:"tenant_namespace_isolation_evidence"`
	TenantScopedAuditNamespaceEvidence          string   `json:"tenant_scoped_audit_namespace_evidence"`
	TenantScopedEvidenceNamespaceEvidence       string   `json:"tenant_scoped_evidence_namespace_evidence"`
	TenantScopedExportBoundaryEvidence          string   `json:"tenant_scoped_export_boundary_evidence"`
	TenantScopedCredentialBoundaryEvidence      string   `json:"tenant_scoped_credential_boundary_evidence"`
	TenantScopedSupportOperatorBoundaryEvidence string   `json:"tenant_scoped_support_operator_boundary_evidence"`
	CrossTenantAuditLeakagePresent              bool     `json:"cross_tenant_audit_leakage_present"`
	CrossTenantEvidenceLeakagePresent           bool     `json:"cross_tenant_evidence_leakage_present"`
	CrossTenantExportLeakagePresent             bool     `json:"cross_tenant_export_leakage_present"`
	CrossTenantCredentialLeakagePresent         bool     `json:"cross_tenant_credential_leakage_present"`
	SupportOperatorAccessLeakagePresent         bool     `json:"support_operator_access_leakage_present"`
	TenantIsolationConfigOnly                   bool     `json:"tenant_isolation_config_only"`
	DashboardSummaryOnly                        bool     `json:"dashboard_summary_only"`
	FleetSummaryOnly                            bool     `json:"fleet_summary_only"`
	RegionSummaryOnly                           bool     `json:"region_summary_only"`
	DeploymentSummaryOnly                       bool     `json:"deployment_summary_only"`
	SupportSummaryOnly                          bool     `json:"support_summary_only"`
	RawCrossTenantEvidenceSharingPresent        bool     `json:"raw_cross_tenant_evidence_sharing_present"`
	TenantPrivateMetadataSideChannelLeakage     bool     `json:"tenant_private_metadata_side_channel_leakage"`
	TenantPrivateMetadataSideChannelMarkedSafe  bool     `json:"tenant_private_metadata_side_channel_marked_safe"`
	CanonicalEvidenceSpineBypass                bool     `json:"canonical_evidence_spine_bypass"`
	TenantProfileNamingExact                    bool     `json:"tenant_profile_naming_exact"`
	SafeIsolationWordingExamplePresent          bool     `json:"safe_isolation_wording_example_present"`
	DiagnosticOutputComplete                    bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer                        string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValBDataResidencyValidator struct {
	CurrentState                     string   `json:"current_state"`
	EvidenceRefs                     []string `json:"evidence_refs,omitempty"`
	FreshnessState                   string   `json:"freshness_state"`
	TenantScope                      string   `json:"tenant_scope"`
	DataResidencyEvidenceBacked      bool     `json:"data_residency_evidence_backed"`
	TenantRegion                     string   `json:"tenant_region"`
	EvidenceRegion                   string   `json:"evidence_region"`
	ExportRegion                     string   `json:"export_region"`
	BackupRegionReference            string   `json:"backup_region_reference"`
	SupportAccessRegion              string   `json:"support_access_region"`
	AllowedRegionPolicy              string   `json:"allowed_region_policy"`
	CrossRegionFlowExists            bool     `json:"cross_region_flow_exists"`
	CrossRegionExceptionPath         string   `json:"cross_region_exception_path"`
	CrossRegionExceptionScoped       bool     `json:"cross_region_exception_scoped"`
	CrossRegionExceptionAudited      bool     `json:"cross_region_exception_audited"`
	CrossRegionExceptionSilentlyOpen bool     `json:"cross_region_exception_silently_open"`
	RegionExportBoundaryValidation   bool     `json:"region_export_boundary_validation"`
	RegionConflictState              string   `json:"region_conflict_state"`
	DataResidencyBypassPresent       bool     `json:"data_residency_bypass_present"`
	BackupPathBypassesResidency      bool     `json:"backup_path_bypasses_residency"`
	ExportPathBypassesResidency      bool     `json:"export_path_bypasses_residency"`
	SupportPathBypassesResidency     bool     `json:"support_path_bypasses_residency"`
	RegionSummaryCanonicalTruth      bool     `json:"region_summary_canonical_truth"`
	CertifiedOrCompliantByDefault    bool     `json:"certified_or_compliant_by_default"`
	RegionNamingExact                bool     `json:"region_naming_exact"`
	SafeResidencyWordingExample      bool     `json:"safe_residency_wording_example_present"`
	DiagnosticOutputComplete         bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValBTenantLifecycle struct {
	CurrentState                             string   `json:"current_state"`
	EvidenceRefs                             []string `json:"evidence_refs,omitempty"`
	FreshnessState                           string   `json:"freshness_state"`
	TenantScope                              string   `json:"tenant_scope"`
	LifecycleState                           string   `json:"lifecycle_state"`
	TenantCreatePresent                      bool     `json:"tenant_create_present"`
	TenantConfigurePresent                   bool     `json:"tenant_configure_present"`
	TenantSuspendPresent                     bool     `json:"tenant_suspend_present"`
	TenantTransferPresent                    bool     `json:"tenant_transfer_present"`
	TenantOffboardPresent                    bool     `json:"tenant_offboard_present"`
	TenantDataExportPresent                  bool     `json:"tenant_data_export_present"`
	TenantEvidenceRetentionPresent           bool     `json:"tenant_evidence_retention_present"`
	TenantDeletionPresent                    bool     `json:"tenant_deletion_present"`
	SupportAccessRevokePresent               bool     `json:"support_access_revoke_present"`
	KeyCustodyRotationPresent                bool     `json:"key_custody_rotation_present"`
	OffboardingRevokesSupportAccess          bool     `json:"offboarding_revokes_support_access"`
	TransferPreservesBoundaries              bool     `json:"transfer_preserves_boundaries"`
	DeletionExportRetentionSemanticsExplicit bool     `json:"deletion_export_retention_semantics_explicit"`
	LifecycleActionTenantScoped              bool     `json:"lifecycle_action_tenant_scoped"`
	StaleRevokedHandlingProven               bool     `json:"stale_revoked_handling_proven"`
	DashboardSummaryOnly                     bool     `json:"dashboard_summary_only"`
	FleetSummaryOnly                         bool     `json:"fleet_summary_only"`
	DiagnosticOutputComplete                 bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer                     string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValBFairShareQuotaPolicy struct {
	CurrentState                                   string   `json:"current_state"`
	EvidenceRefs                                   []string `json:"evidence_refs,omitempty"`
	FreshnessState                                 string   `json:"freshness_state"`
	TenantScope                                    string   `json:"tenant_scope"`
	EventBudgetPerTenantPresent                    bool     `json:"event_budget_per_tenant_present"`
	QueueIsolationPresent                          bool     `json:"queue_isolation_present"`
	NoisyTenantContainmentPresent                  bool     `json:"noisy_tenant_containment_present"`
	AlertFloodThrottlingPresent                    bool     `json:"alert_flood_throttling_present"`
	NoStarvationRulePresent                        bool     `json:"no_starvation_rule_present"`
	OverloadDowngradeSemanticsPresent              bool     `json:"overload_downgrade_semantics_present"`
	PerTenantRateLimitEvidencePresent              bool     `json:"per_tenant_rate_limit_evidence_present"`
	PerTenantBackpressureSemanticsPresent          bool     `json:"per_tenant_backpressure_semantics_present"`
	BoundedDegradationStatePresent                 bool     `json:"bounded_degradation_state_present"`
	TenantAwareNegativeTestPresent                 bool     `json:"tenant_aware_negative_test_present"`
	VerificationIngestAuditIsolationUnderLoad      bool     `json:"verification_ingest_audit_isolation_under_load"`
	OneTenantStarvesAnother                        bool     `json:"one_tenant_starves_another"`
	NoisyTenantDegradesAnotherTenant               bool     `json:"noisy_tenant_degrades_another_tenant"`
	AlertFloodSpillsAcrossTenants                  bool     `json:"alert_flood_spills_across_tenants"`
	OverloadSilentlyTreatedAsReady                 bool     `json:"overload_silently_treated_as_ready"`
	GlobalQueueStarvationWithoutBoundedDegradation bool     `json:"global_queue_starvation_without_bounded_degradation"`
	DiagnosticOutputComplete                       bool     `json:"diagnostic_output_complete"`
	RunbookWordingComplete                         bool     `json:"runbook_wording_complete"`
	ProjectionDisclaimer                           string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValBNoOverclaimDiscipline struct {
	CurrentState                 string   `json:"current_state"`
	ObservedClaims               []string `json:"observed_claims,omitempty"`
	CleanRoomIPViolationDetected bool     `json:"clean_room_ip_violation_detected"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValBClosureBlockerFinding struct {
	BlockerLevel      string `json:"blocker_level"`
	Surface           string `json:"surface"`
	Reason            string `json:"reason"`
	BlocksCurrentWave bool   `json:"blocks_current_wave"`
	RequiredFollowup  string `json:"required_followup,omitempty"`
}

type DeploymentMultiTenantValBClosureBlockerOverlay struct {
	CurrentState         string                                           `json:"current_state"`
	Findings             []DeploymentMultiTenantValBClosureBlockerFinding `json:"findings,omitempty"`
	ProjectionDisclaimer string                                           `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValBFoundation struct {
	CurrentState          string                                           `json:"current_state"`
	Point10State          string                                           `json:"point_10_state"`
	ProjectionDisclaimer  string                                           `json:"projection_disclaimer"`
	BlockingReasons       []string                                         `json:"blocking_reasons,omitempty"`
	DependencyState       string                                           `json:"dependency_state"`
	TenantIsolationState  string                                           `json:"tenant_isolation_state"`
	DataResidencyState    string                                           `json:"data_residency_state"`
	TenantLifecycleState  string                                           `json:"tenant_lifecycle_state"`
	FairShareQuotaState   string                                           `json:"fair_share_quota_state"`
	NoOverclaimState      string                                           `json:"no_overclaim_state"`
	ClosureBlockerState   string                                           `json:"closure_blocker_state"`
	Dependency            DeploymentMultiTenantValBDependencySnapshot      `json:"dependency"`
	TenantIsolation       DeploymentMultiTenantValBTenantIsolationTestPack `json:"tenant_isolation"`
	DataResidency         DeploymentMultiTenantValBDataResidencyValidator  `json:"data_residency"`
	TenantLifecycle       DeploymentMultiTenantValBTenantLifecycle         `json:"tenant_lifecycle"`
	FairShareQuota        DeploymentMultiTenantValBFairShareQuotaPolicy    `json:"fair_share_quota"`
	NoOverclaim           DeploymentMultiTenantValBNoOverclaimDiscipline   `json:"no_overclaim"`
	ClosureBlockerOverlay DeploymentMultiTenantValBClosureBlockerOverlay   `json:"closure_blocker_overlay"`
}

func deploymentMultiTenantValBProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_marketplace_deployment_profile deployment_multi_tenant_valb"
}

func deploymentMultiTenantValBHasProjectionDisclaimer(value string) bool {
	return value == deploymentMultiTenantValBProjectionDisclaimer() ||
		value == deploymentMultiTenantValBProjectionDisclaimer()+" aggregate_dependency_snapshot" ||
		value == "projection_only not_canonical_truth deployment_multi_tenant_valb aggregate_dependency_snapshot"
}

func deploymentMultiTenantValBHasFoundationProjectionDisclaimer(value string) bool {
	return value == deploymentMultiTenantValBProjectionDisclaimer()
}

func deploymentMultiTenantValBTenantIsolationEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-valb-tenant-isolation-001"}
}

func deploymentMultiTenantValBHasExactTenantScope(value string) bool {
	return value == deploymentMultiTenantVal0TenantScope()
}

func deploymentMultiTenantValBDataResidencyEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-valb-data-residency-001"}
}

func deploymentMultiTenantValBTenantLifecycleEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-valb-tenant-lifecycle-001"}
}

func deploymentMultiTenantValBFairShareEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-valb-fair-share-001"}
}

func deploymentMultiTenantValBClosureBlockerLevels() []string {
	return []string{
		DeploymentMultiTenantValBBlockerLevelCLB0,
		DeploymentMultiTenantValBBlockerLevelCLB1,
		DeploymentMultiTenantValBBlockerLevelCLB2,
		DeploymentMultiTenantValBBlockerLevelCLB3,
	}
}

func deploymentMultiTenantValBClosureBlockerSurfaces() []string {
	return []string{
		DeploymentMultiTenantValBClosureSurfaceTenantIsolation,
		DeploymentMultiTenantValBClosureSurfaceDataResidency,
		DeploymentMultiTenantValBClosureSurfaceTenantLifecycle,
		DeploymentMultiTenantValBClosureSurfaceFairShare,
		DeploymentMultiTenantValBClosureSurfaceNoOverclaim,
		DeploymentMultiTenantValBClosureSurfaceCleanRoomIP,
	}
}

func deploymentMultiTenantValBDependencySnapshotModel() DeploymentMultiTenantValBDependencySnapshot {
	valA := ComputeDeploymentMultiTenantValAFoundation(DeploymentMultiTenantValAFoundationModel())
	return DeploymentMultiTenantValBDependencySnapshot{
		ValACurrentState:                 valA.CurrentState,
		ValADependencyState:              valA.DependencyState,
		ValADeploymentProfileMatrixState: valA.DeploymentProfileMatrixState,
		ValAPreflightGateState:           valA.PreflightGateState,
		ValAIdentityBootstrapState:       valA.IdentityBootstrapState,
		ValAAirGappedEvidenceBundleState: valA.AirGappedEvidenceBundleState,
		ValANoOverclaimState:             valA.NoOverclaimState,
		ValAPassBlockerState:             valA.PassBlockerState,
		Point10State:                     valA.Point10State,
		ProjectionDisclaimer:             valA.ProjectionDisclaimer,
	}
}

func EvaluateDeploymentMultiTenantValBDependencyState(model DeploymentMultiTenantValBDependencySnapshot) string {
	if !deploymentMultiTenantValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeploymentMultiTenantValBDependencyStateBlocked
	}
	if model.ValACurrentState != DeploymentMultiTenantValAStateActive ||
		model.ValADependencyState != DeploymentMultiTenantValADependencyStateActive ||
		model.ValADeploymentProfileMatrixState != DeploymentMultiTenantValADeploymentProfileMatrixStateActive ||
		model.ValAPreflightGateState != DeploymentMultiTenantValAPreflightGateStateActive ||
		model.ValAIdentityBootstrapState != DeploymentMultiTenantValAIdentityBootstrapStateActive ||
		model.ValAAirGappedEvidenceBundleState != DeploymentMultiTenantValAAirGappedEvidenceBundleStateActive ||
		model.ValANoOverclaimState != DeploymentMultiTenantValANoOverclaimStateActive ||
		model.ValAPassBlockerState != DeploymentMultiTenantValAPassBlockerStateActive ||
		model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		return DeploymentMultiTenantValBDependencyStateBlocked
	}
	return DeploymentMultiTenantValBDependencyStateActive
}

func EvaluateDeploymentMultiTenantValBTenantIsolationState(model DeploymentMultiTenantValBTenantIsolationTestPack) string {
	if !deploymentMultiTenantValBHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!deploymentMultiTenantVal0ContainsExactStringSet(model.EvidenceRefs, deploymentMultiTenantValBTenantIsolationEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!deploymentMultiTenantValBHasExactTenantScope(model.TenantScope) ||
		!model.TenantIsolationEvidenceBacked ||
		!model.CrossTenantAuditLeakageTestPresent ||
		!model.CrossTenantEvidenceLeakageTestPresent ||
		!model.CrossTenantExportLeakageTestPresent ||
		!model.CrossTenantCredentialLeakageTestPresent ||
		!model.SupportOperatorAccessLeakageTestPresent ||
		!model.TenantBoundaryValidationTestPresent ||
		!model.TenantScopeNegativeTestPresent ||
		!model.RawCrossTenantEvidenceSharingNegativeTest ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.TenantNamespaceIsolationEvidence) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantScopedAuditNamespaceEvidence) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantScopedEvidenceNamespaceEvidence) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantScopedExportBoundaryEvidence) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantScopedCredentialBoundaryEvidence) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantScopedSupportOperatorBoundaryEvidence) ||
		model.CrossTenantAuditLeakagePresent ||
		model.CrossTenantEvidenceLeakagePresent ||
		model.CrossTenantExportLeakagePresent ||
		model.CrossTenantCredentialLeakagePresent ||
		model.SupportOperatorAccessLeakagePresent ||
		model.TenantIsolationConfigOnly ||
		model.DashboardSummaryOnly ||
		model.FleetSummaryOnly ||
		model.RegionSummaryOnly ||
		model.DeploymentSummaryOnly ||
		model.SupportSummaryOnly ||
		model.RawCrossTenantEvidenceSharingPresent ||
		model.TenantPrivateMetadataSideChannelLeakage ||
		model.TenantPrivateMetadataSideChannelMarkedSafe ||
		model.CanonicalEvidenceSpineBypass {
		return DeploymentMultiTenantValBTenantIsolationStateBlocked
	}
	return DeploymentMultiTenantValBTenantIsolationStateActive
}

func deploymentMultiTenantValBDataResidencyHasInferredCrossRegionFlow(model DeploymentMultiTenantValBDataResidencyValidator) bool {
	tenantRegion := strings.TrimSpace(model.TenantRegion)
	if tenantRegion == "" {
		return false
	}
	regions := []string{
		strings.TrimSpace(model.EvidenceRegion),
		strings.TrimSpace(model.ExportRegion),
		strings.TrimSpace(model.BackupRegionReference),
		strings.TrimSpace(model.SupportAccessRegion),
	}
	for _, region := range regions {
		if region != "" && region != tenantRegion {
			return true
		}
	}
	return false
}

func deploymentMultiTenantValBDataResidencyHasValidCrossRegionException(model DeploymentMultiTenantValBDataResidencyValidator) bool {
	return deploymentMultiTenantVal0ExactValueIsValid(model.CrossRegionExceptionPath) &&
		model.CrossRegionExceptionScoped &&
		model.CrossRegionExceptionAudited &&
		!model.CrossRegionExceptionSilentlyOpen
}

func EvaluateDeploymentMultiTenantValBDataResidencyState(model DeploymentMultiTenantValBDataResidencyValidator) string {
	effectiveCrossRegionFlow := model.CrossRegionFlowExists || deploymentMultiTenantValBDataResidencyHasInferredCrossRegionFlow(model)
	if !deploymentMultiTenantValBHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!deploymentMultiTenantVal0ContainsExactStringSet(model.EvidenceRefs, deploymentMultiTenantValBDataResidencyEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!deploymentMultiTenantValBHasExactTenantScope(model.TenantScope) ||
		!model.DataResidencyEvidenceBacked ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.TenantRegion) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.EvidenceRegion) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.ExportRegion) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.BackupRegionReference) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SupportAccessRegion) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.AllowedRegionPolicy) ||
		!model.RegionExportBoundaryValidation ||
		model.RegionConflictState != DeploymentMultiTenantConflictStateNoConflict ||
		model.DataResidencyBypassPresent ||
		model.BackupPathBypassesResidency ||
		model.ExportPathBypassesResidency ||
		model.SupportPathBypassesResidency ||
		model.RegionSummaryCanonicalTruth ||
		model.CertifiedOrCompliantByDefault {
		return DeploymentMultiTenantValBDataResidencyStateBlocked
	}
	if effectiveCrossRegionFlow && !deploymentMultiTenantValBDataResidencyHasValidCrossRegionException(model) {
		return DeploymentMultiTenantValBDataResidencyStateBlocked
	}
	return DeploymentMultiTenantValBDataResidencyStateActive
}

func EvaluateDeploymentMultiTenantValBTenantLifecycleState(model DeploymentMultiTenantValBTenantLifecycle) string {
	if !deploymentMultiTenantValBHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!deploymentMultiTenantVal0ContainsExactStringSet(model.EvidenceRefs, deploymentMultiTenantValBTenantLifecycleEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!deploymentMultiTenantValBHasExactTenantScope(model.TenantScope) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.LifecycleState) ||
		!model.TenantCreatePresent ||
		!model.TenantConfigurePresent ||
		!model.TenantSuspendPresent ||
		!model.TenantTransferPresent ||
		!model.TenantOffboardPresent ||
		!model.TenantDataExportPresent ||
		!model.TenantEvidenceRetentionPresent ||
		!model.TenantDeletionPresent ||
		!model.SupportAccessRevokePresent ||
		!model.KeyCustodyRotationPresent ||
		!model.OffboardingRevokesSupportAccess ||
		!model.TransferPreservesBoundaries ||
		!model.DeletionExportRetentionSemanticsExplicit ||
		!model.LifecycleActionTenantScoped ||
		!model.StaleRevokedHandlingProven ||
		model.DashboardSummaryOnly ||
		model.FleetSummaryOnly {
		return DeploymentMultiTenantValBTenantLifecycleStateBlocked
	}
	return DeploymentMultiTenantValBTenantLifecycleStateActive
}

func EvaluateDeploymentMultiTenantValBFairShareQuotaState(model DeploymentMultiTenantValBFairShareQuotaPolicy) string {
	if !deploymentMultiTenantValBHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!deploymentMultiTenantVal0ContainsExactStringSet(model.EvidenceRefs, deploymentMultiTenantValBFairShareEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!deploymentMultiTenantValBHasExactTenantScope(model.TenantScope) ||
		!model.EventBudgetPerTenantPresent ||
		!model.QueueIsolationPresent ||
		!model.NoisyTenantContainmentPresent ||
		!model.AlertFloodThrottlingPresent ||
		!model.NoStarvationRulePresent ||
		!model.OverloadDowngradeSemanticsPresent ||
		!model.PerTenantRateLimitEvidencePresent ||
		!model.PerTenantBackpressureSemanticsPresent ||
		!model.BoundedDegradationStatePresent ||
		!model.TenantAwareNegativeTestPresent ||
		!model.VerificationIngestAuditIsolationUnderLoad ||
		model.OneTenantStarvesAnother ||
		model.NoisyTenantDegradesAnotherTenant ||
		model.AlertFloodSpillsAcrossTenants ||
		model.OverloadSilentlyTreatedAsReady ||
		model.GlobalQueueStarvationWithoutBoundedDegradation {
		return DeploymentMultiTenantValBFairShareQuotaStateBlocked
	}
	return DeploymentMultiTenantValBFairShareQuotaStateActive
}

func deploymentMultiTenantValBContainsForbiddenClaim(values ...string) bool {
	allowed := []string{
		"evidence-linked tenant isolation test",
		"tenant-scoped audit boundary",
		"tenant-scoped evidence boundary",
		"tenant-scoped export boundary",
		"tenant-scoped credential boundary",
		"data residency evidence",
		"region/export boundary validation",
		"bounded cross-region exception path",
		"tenant lifecycle evidence",
		"support access revoke evidence",
		"key/custody rotation evidence",
		"fair-share quota evidence",
		"tenant-aware negative test",
		"noisy tenant containment evidence",
		"bounded degradation semantics",
		"not production approval",
		"not compliance certification",
		"not canonical truth",
		"advisory fleet visibility",
	}
	disallowed := []string{
		"tenant isolation guaranteed",
		"zero cross-tenant leakage",
		"no leakage guaranteed",
		"data residency certified",
		"data residency guaranteed",
		"region compliant by default",
		"sovereign compliant by default",
		"all tenants isolated by default",
		"tenant safe by default",
		"cross-tenant safe by default",
		"support access cannot leak",
		"backup residency guaranteed",
		"export residency guaranteed",
		"fair-share guarantees no outages",
		"quotas guarantee tenant performance",
		"noisy tenant cannot impact anyone",
		"lifecycle complete means compliant",
		"offboarding guarantees deletion",
		"deletion guaranteed",
		"transfer safe by default",
		"dashboard proves tenant isolation",
		"fleet view proves data residency",
		"support summary is canonical truth",
		"region summary is canonical truth",
	}
	allowedExact := make(map[string]struct{}, len(allowed)*2)
	for _, phrase := range allowed {
		allowedExact[deploymentMultiTenantVal0NormalizeClaimText(phrase)] = struct{}{}
		allowedExact[deploymentMultiTenantVal0CompactClaimText(phrase)] = struct{}{}
	}
	blockedNormalized := make([]string, 0, len(disallowed))
	blockedCompact := make([]string, 0, len(disallowed))
	for _, phrase := range disallowed {
		blockedNormalized = append(blockedNormalized, deploymentMultiTenantVal0NormalizeClaimText(phrase))
		blockedCompact = append(blockedCompact, deploymentMultiTenantVal0CompactClaimText(phrase))
	}
	crossNormalizedParts := make([]string, 0, len(values))
	corpusNormalizedParts := make([]string, 0, len(values))
	var corpusCompact strings.Builder
	for _, value := range values {
		normalized := deploymentMultiTenantVal0NormalizeClaimText(value)
		compact := deploymentMultiTenantVal0CompactClaimText(value)
		if normalized == "" && compact == "" {
			continue
		}
		if normalized != "" {
			crossNormalizedParts = append(crossNormalizedParts, normalized)
		}
		if _, ok := allowedExact[normalized]; ok {
			continue
		}
		if _, ok := allowedExact[compact]; ok {
			continue
		}
		if normalized != "" {
			corpusNormalizedParts = append(corpusNormalizedParts, normalized)
		}
		corpusCompact.WriteString(compact)
		for i := range blockedNormalized {
			if strings.Contains(normalized, blockedNormalized[i]) ||
				strings.Contains(compact, blockedCompact[i]) ||
				deploymentMultiTenantVal0ValueContainsForbiddenPhraseTokenSequence(normalized, blockedNormalized[i]) {
				return true
			}
		}
	}
	corpusNormalized := strings.Join(corpusNormalizedParts, " ")
	corpusCompactValue := corpusCompact.String()
	for i := range blockedNormalized {
		if strings.Contains(corpusNormalized, blockedNormalized[i]) || strings.Contains(corpusCompactValue, blockedCompact[i]) {
			return true
		}
		if deploymentMultiTenantVal0BucketsContainForbiddenPhraseAcrossValues(crossNormalizedParts, blockedNormalized[i]) {
			return true
		}
	}
	return false
}

func EvaluateDeploymentMultiTenantValBNoOverclaimState(model DeploymentMultiTenantValBNoOverclaimDiscipline) string {
	if !deploymentMultiTenantValBHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		model.CleanRoomIPViolationDetected ||
		deploymentMultiTenantValBContainsForbiddenClaim(model.ObservedClaims...) {
		return DeploymentMultiTenantValBNoOverclaimStateBlocked
	}
	return DeploymentMultiTenantValBNoOverclaimStateActive
}

func deploymentMultiTenantValBClosureBlockerFinding(level, surface, reason string, blocksCurrentWave bool, requiredFollowup string) DeploymentMultiTenantValBClosureBlockerFinding {
	return DeploymentMultiTenantValBClosureBlockerFinding{
		BlockerLevel:      level,
		Surface:           surface,
		Reason:            reason,
		BlocksCurrentWave: blocksCurrentWave,
		RequiredFollowup:  requiredFollowup,
	}
}

func deploymentMultiTenantValBClosureBlockerFindings(model DeploymentMultiTenantValBFoundation) []DeploymentMultiTenantValBClosureBlockerFinding {
	findings := []DeploymentMultiTenantValBClosureBlockerFinding{}
	inferredCrossRegionFlow := deploymentMultiTenantValBDataResidencyHasInferredCrossRegionFlow(model.DataResidency)
	effectiveCrossRegionFlow := model.DataResidency.CrossRegionFlowExists || inferredCrossRegionFlow
	if model.TenantIsolation.CrossTenantAuditLeakagePresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "cross-tenant audit leakage", true, "remove cross-tenant audit leakage and rerun tenant isolation evidence"))
	}
	if model.TenantIsolation.CrossTenantEvidenceLeakagePresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "cross-tenant evidence leakage", true, "remove cross-tenant evidence leakage and rerun tenant isolation evidence"))
	}
	if model.TenantIsolation.CrossTenantExportLeakagePresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "cross-tenant export leakage", true, "restore export boundary isolation and rerun tenant isolation evidence"))
	}
	if model.TenantIsolation.CrossTenantCredentialLeakagePresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "cross-tenant credential leakage", true, "restore credential boundary isolation and rerun tenant isolation evidence"))
	}
	if model.TenantIsolation.SupportOperatorAccessLeakagePresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "support or operator access leakage", true, "constrain support and operator access scope and rerun tenant isolation evidence"))
	}
	if model.DataResidency.DataResidencyBypassPresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceDataResidency, "data residency bypass", true, "remove residency bypass path and rerun residency validation"))
	}
	if model.DataResidency.BackupPathBypassesResidency {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceDataResidency, "backup path bypasses data residency", true, "constrain backup path to tenant residency policy and rerun residency validation"))
	}
	if model.DataResidency.ExportPathBypassesResidency {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceDataResidency, "export path bypasses data residency", true, "constrain export path to tenant residency policy and rerun residency validation"))
	}
	if model.DataResidency.SupportPathBypassesResidency {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceDataResidency, "support path bypasses data residency", true, "constrain support path to tenant residency policy and rerun residency validation"))
	}
	if model.TenantIsolation.RawCrossTenantEvidenceSharingPresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "raw cross-tenant evidence sharing", true, "remove cross-tenant evidence sharing and rerun tenant isolation evidence"))
	}
	if model.TenantIsolation.TenantPrivateMetadataSideChannelMarkedSafe {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "tenant-private metadata side-channel leakage marked safe", true, "remove unsupported side-channel safety claim and rerun tenant isolation evidence"))
	}
	if model.TenantIsolation.TenantIsolationConfigOnly {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "tenant isolation treated as config-only", true, "replace config-only claim with evidence-linked tenant isolation test coverage"))
	}
	if model.TenantIsolation.DashboardSummaryOnly {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "dashboard summary treated as canonical isolation evidence", true, "replace dashboard summary inference with tenant isolation evidence"))
	}
	if model.TenantIsolation.FleetSummaryOnly {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "fleet summary treated as canonical isolation evidence", true, "replace fleet summary inference with tenant isolation evidence"))
	}
	if model.TenantIsolation.CanonicalEvidenceSpineBypass {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "canonical evidence spine bypass", true, "restore canonical evidence spine discipline before final review"))
	}
	if model.DataResidency.RegionSummaryCanonicalTruth {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceDataResidency, "region summary treated as canonical truth", true, "replace region summary inference with tenant-scoped data residency evidence"))
	}
	if inferredCrossRegionFlow && !deploymentMultiTenantValBDataResidencyHasValidCrossRegionException(model.DataResidency) {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceDataResidency, "inferred cross-region residency flow without scoped audited exception", true, "add scoped audited cross-region exception or align region boundaries and rerun residency validation"))
	}
	if effectiveCrossRegionFlow && model.DataResidency.CrossRegionExceptionSilentlyOpen {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceDataResidency, "cross-region exception silently allowed", true, "make the cross-region exception explicit, scoped, and audited before handoff"))
	}
	if model.NoOverclaim.CleanRoomIPViolationDetected {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBClosureSurfaceCleanRoomIP, "copied competitor deployment or tenant isolation artifact detected", true, "remove copied artifact and replace with clean-room implementation evidence"))
	}
	if !model.FairShareQuota.TenantAwareNegativeTestPresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB1, DeploymentMultiTenantValBClosureSurfaceFairShare, "fair-share or quota policy lacks tenant-aware negative test", true, "add tenant-aware negative test coverage for fair-share and quota containment"))
	}
	if !model.TenantIsolation.TenantScopeNegativeTestPresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB1, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "malformed unknown or stale tenant scope not tested", true, "add malformed unknown and stale tenant scope negative coverage"))
	}
	if !model.TenantLifecycle.TenantOffboardPresent || !model.TenantLifecycle.SupportAccessRevokePresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB1, DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, "tenant lifecycle lacks offboarding or revoke semantics", true, "add offboarding and support revoke lifecycle evidence"))
	}
	if !model.TenantLifecycle.OffboardingRevokesSupportAccess {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB1, DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, "offboarding does not revoke support or operator access", true, "enforce support and operator access revocation during tenant offboarding"))
	}
	if !model.TenantLifecycle.TransferPreservesBoundaries {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB1, DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, "tenant transfer weakens audit evidence or export boundaries", true, "preserve audit evidence and export boundaries across tenant transfer"))
	}
	if !model.TenantLifecycle.DeletionExportRetentionSemanticsExplicit {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB1, DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, "deletion export or retention semantics ambiguous", true, "make deletion export and retention semantics explicit"))
	}
	if model.DataResidency.CrossRegionFlowExists && !deploymentMultiTenantVal0ExactValueIsValid(model.DataResidency.CrossRegionExceptionPath) {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB1, DeploymentMultiTenantValBClosureSurfaceDataResidency, "data residency exception path missing while cross-region flow exists", true, "define audited scoped exception path for cross-region flow"))
	}
	if model.DependencyState != DeploymentMultiTenantValBDependencyStateActive {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB1, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "dependency gate missing or not exact active", true, "restore exact active Val A dependency before Val B final review"))
	}
	if !model.TenantLifecycle.StaleRevokedHandlingProven {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB1, DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, "stale or revoked handling not proven where Val B depends on it", true, "prove stale and revoked handling across tenant lifecycle transitions"))
	}
	if !model.DataResidency.RegionNamingExact {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB2, DeploymentMultiTenantValBClosureSurfaceDataResidency, "ambiguous region naming", true, "normalize region naming before handoff"))
	}
	if !model.TenantIsolation.TenantProfileNamingExact {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB2, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "ambiguous deployment or tenant profile naming", true, "normalize tenant and deployment profile naming before handoff"))
	}
	if !model.TenantIsolation.SafeIsolationWordingExamplePresent {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB2, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "missing safe wording example for tenant isolation", true, "add bounded safe wording example for tenant isolation"))
	}
	if !model.DataResidency.SafeResidencyWordingExample {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB2, DeploymentMultiTenantValBClosureSurfaceDataResidency, "missing safe wording example for data residency", true, "add bounded safe wording example for data residency"))
	}
	if !model.TenantIsolation.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB2, DeploymentMultiTenantValBClosureSurfaceTenantIsolation, "incomplete diagnostic output for tenant isolation blockers", true, "complete tenant isolation diagnostic output before handoff"))
	}
	if !model.DataResidency.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB2, DeploymentMultiTenantValBClosureSurfaceDataResidency, "incomplete diagnostic output for data residency blockers", true, "complete data residency diagnostic output before handoff"))
	}
	if !model.TenantLifecycle.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB2, DeploymentMultiTenantValBClosureSurfaceTenantLifecycle, "incomplete diagnostic output for tenant lifecycle blockers", true, "complete tenant lifecycle diagnostic output before handoff"))
	}
	if !model.FairShareQuota.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB2, DeploymentMultiTenantValBClosureSurfaceFairShare, "incomplete diagnostic output for fair-share blockers", true, "complete fair-share diagnostic output before handoff"))
	}
	if !model.FairShareQuota.RunbookWordingComplete {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB2, DeploymentMultiTenantValBClosureSurfaceFairShare, "incomplete runbook wording without direct closure bypass", true, "complete bounded runbook wording before handoff"))
	}
	if model.NoOverclaimState != DeploymentMultiTenantValBNoOverclaimStateActive {
		findings = append(findings, deploymentMultiTenantValBClosureBlockerFinding(DeploymentMultiTenantValBBlockerLevelCLB1, DeploymentMultiTenantValBClosureSurfaceNoOverclaim, "forbidden tenant isolation or residency claim present", true, "remove forbidden tenant isolation or residency overclaim"))
	}
	return findings
}

func EvaluateDeploymentMultiTenantValBClosureBlockerState(model DeploymentMultiTenantValBClosureBlockerOverlay) string {
	if !deploymentMultiTenantValBHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeploymentMultiTenantValBClosureBlockerStateBlocked
	}
	hasCleanup := false
	hasAdvisory := false
	for _, finding := range model.Findings {
		level := strings.TrimSpace(finding.BlockerLevel)
		surface := strings.TrimSpace(finding.Surface)
		if len(level) == 2 && level[0] == 'P' && level[1] >= '0' && level[1] <= '9' {
			return DeploymentMultiTenantValBClosureBlockerStateBlocked
		}
		if !containsTrimmedString(deploymentMultiTenantValBClosureBlockerLevels(), level) ||
			!containsTrimmedString(deploymentMultiTenantValBClosureBlockerSurfaces(), surface) {
			return DeploymentMultiTenantValBClosureBlockerStateBlocked
		}
		if (level == DeploymentMultiTenantValBBlockerLevelCLB1 ||
			level == DeploymentMultiTenantValBBlockerLevelCLB2 ||
			level == DeploymentMultiTenantValBBlockerLevelCLB3) &&
			strings.TrimSpace(finding.RequiredFollowup) == "" {
			return DeploymentMultiTenantValBClosureBlockerStateBlocked
		}
		switch level {
		case DeploymentMultiTenantValBBlockerLevelCLB0, DeploymentMultiTenantValBBlockerLevelCLB1:
			return DeploymentMultiTenantValBClosureBlockerStateBlocked
		case DeploymentMultiTenantValBBlockerLevelCLB2:
			hasCleanup = true
		case DeploymentMultiTenantValBBlockerLevelCLB3:
			hasAdvisory = true
		default:
			return DeploymentMultiTenantValBClosureBlockerStateBlocked
		}
	}
	if hasCleanup {
		return DeploymentMultiTenantValBClosureBlockerStateCleanup
	}
	if hasAdvisory {
		return DeploymentMultiTenantValBClosureBlockerStateAdvisory
	}
	return DeploymentMultiTenantValBClosureBlockerStateActive
}

func EvaluateDeploymentMultiTenantValBState(model DeploymentMultiTenantValBFoundation) string {
	if !deploymentMultiTenantValBHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		model.DependencyState != DeploymentMultiTenantValBDependencyStateActive ||
		model.TenantIsolationState != DeploymentMultiTenantValBTenantIsolationStateActive ||
		model.DataResidencyState != DeploymentMultiTenantValBDataResidencyStateActive ||
		model.TenantLifecycleState != DeploymentMultiTenantValBTenantLifecycleStateActive ||
		model.FairShareQuotaState != DeploymentMultiTenantValBFairShareQuotaStateActive ||
		model.NoOverclaimState != DeploymentMultiTenantValBNoOverclaimStateActive ||
		model.ClosureBlockerState != DeploymentMultiTenantValBClosureBlockerStateActive ||
		model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		return DeploymentMultiTenantValBStateBlocked
	}
	return DeploymentMultiTenantValBStateActive
}

func deploymentMultiTenantValBBlockingReasons(model DeploymentMultiTenantValBFoundation) []string {
	reasons := []string{}
	if !deploymentMultiTenantValBHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "aggregate_projection_disclaimer_blocked")
	}
	if model.DependencyState != DeploymentMultiTenantValBDependencyStateActive {
		reasons = append(reasons, "dependency_state_blocked")
	}
	if model.TenantIsolationState != DeploymentMultiTenantValBTenantIsolationStateActive {
		reasons = append(reasons, "tenant_isolation_state_blocked")
	}
	if model.DataResidencyState != DeploymentMultiTenantValBDataResidencyStateActive {
		reasons = append(reasons, "data_residency_state_blocked")
	}
	if model.TenantLifecycleState != DeploymentMultiTenantValBTenantLifecycleStateActive {
		reasons = append(reasons, "tenant_lifecycle_state_blocked")
	}
	if model.FairShareQuotaState != DeploymentMultiTenantValBFairShareQuotaStateActive {
		reasons = append(reasons, "fair_share_quota_state_blocked")
	}
	if model.NoOverclaimState != DeploymentMultiTenantValBNoOverclaimStateActive {
		reasons = append(reasons, "no_overclaim_state_blocked")
	}
	if model.ClosureBlockerState != DeploymentMultiTenantValBClosureBlockerStateActive {
		reasons = append(reasons, "closure_blocker_state_not_clean")
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		reasons = append(reasons, "point10_state_not_complete_guard_violated")
	}
	return reasons
}

func DeploymentMultiTenantValBFoundationModel() DeploymentMultiTenantValBFoundation {
	disclaimer := deploymentMultiTenantValBProjectionDisclaimer()
	return DeploymentMultiTenantValBFoundation{
		CurrentState:         DeploymentMultiTenantValBStateActive,
		Point10State:         DeploymentMultiTenantPoint10StateNotComplete,
		ProjectionDisclaimer: disclaimer,
		DependencyState:      DeploymentMultiTenantValBDependencyStateActive,
		TenantIsolationState: DeploymentMultiTenantValBTenantIsolationStateActive,
		DataResidencyState:   DeploymentMultiTenantValBDataResidencyStateActive,
		TenantLifecycleState: DeploymentMultiTenantValBTenantLifecycleStateActive,
		FairShareQuotaState:  DeploymentMultiTenantValBFairShareQuotaStateActive,
		NoOverclaimState:     DeploymentMultiTenantValBNoOverclaimStateActive,
		ClosureBlockerState:  DeploymentMultiTenantValBClosureBlockerStateActive,
		Dependency:           deploymentMultiTenantValBDependencySnapshotModel(),
		TenantIsolation: DeploymentMultiTenantValBTenantIsolationTestPack{
			CurrentState:                                DeploymentMultiTenantValBTenantIsolationStateActive,
			EvidenceRefs:                                append([]string{}, deploymentMultiTenantValBTenantIsolationEvidenceRefs()...),
			FreshnessState:                              IntelligenceCalibrationFreshnessFresh,
			TenantScope:                                 "tenant:alpha",
			TenantIsolationEvidenceBacked:               true,
			CrossTenantAuditLeakageTestPresent:          true,
			CrossTenantEvidenceLeakageTestPresent:       true,
			CrossTenantExportLeakageTestPresent:         true,
			CrossTenantCredentialLeakageTestPresent:     true,
			SupportOperatorAccessLeakageTestPresent:     true,
			TenantBoundaryValidationTestPresent:         true,
			TenantScopeNegativeTestPresent:              true,
			RawCrossTenantEvidenceSharingNegativeTest:   true,
			TenantNamespaceIsolationEvidence:            "tenant_namespace_isolation_evidence",
			TenantScopedAuditNamespaceEvidence:          "tenant_scoped_audit_namespace_evidence",
			TenantScopedEvidenceNamespaceEvidence:       "tenant_scoped_evidence_namespace_evidence",
			TenantScopedExportBoundaryEvidence:          "tenant_scoped_export_boundary_evidence",
			TenantScopedCredentialBoundaryEvidence:      "tenant_scoped_credential_boundary_evidence",
			TenantScopedSupportOperatorBoundaryEvidence: "tenant_scoped_support_operator_boundary_evidence",
			TenantProfileNamingExact:                    true,
			SafeIsolationWordingExamplePresent:          true,
			DiagnosticOutputComplete:                    true,
			ProjectionDisclaimer:                        disclaimer,
		},
		DataResidency: DeploymentMultiTenantValBDataResidencyValidator{
			CurrentState:                   DeploymentMultiTenantValBDataResidencyStateActive,
			EvidenceRefs:                   append([]string{}, deploymentMultiTenantValBDataResidencyEvidenceRefs()...),
			FreshnessState:                 IntelligenceCalibrationFreshnessFresh,
			TenantScope:                    "tenant:alpha",
			DataResidencyEvidenceBacked:    true,
			TenantRegion:                   "eu_central_1",
			EvidenceRegion:                 "eu_central_1",
			ExportRegion:                   "eu_central_1",
			BackupRegionReference:          "eu_central_1",
			SupportAccessRegion:            "eu_central_1",
			AllowedRegionPolicy:            "tenant_region_policy_exact",
			RegionExportBoundaryValidation: true,
			RegionConflictState:            DeploymentMultiTenantConflictStateNoConflict,
			RegionNamingExact:              true,
			SafeResidencyWordingExample:    true,
			DiagnosticOutputComplete:       true,
			ProjectionDisclaimer:           disclaimer,
		},
		TenantLifecycle: DeploymentMultiTenantValBTenantLifecycle{
			CurrentState:                             DeploymentMultiTenantValBTenantLifecycleStateActive,
			EvidenceRefs:                             append([]string{}, deploymentMultiTenantValBTenantLifecycleEvidenceRefs()...),
			FreshnessState:                           IntelligenceCalibrationFreshnessFresh,
			TenantScope:                              "tenant:alpha",
			LifecycleState:                           "tenant_lifecycle_governed",
			TenantCreatePresent:                      true,
			TenantConfigurePresent:                   true,
			TenantSuspendPresent:                     true,
			TenantTransferPresent:                    true,
			TenantOffboardPresent:                    true,
			TenantDataExportPresent:                  true,
			TenantEvidenceRetentionPresent:           true,
			TenantDeletionPresent:                    true,
			SupportAccessRevokePresent:               true,
			KeyCustodyRotationPresent:                true,
			OffboardingRevokesSupportAccess:          true,
			TransferPreservesBoundaries:              true,
			DeletionExportRetentionSemanticsExplicit: true,
			LifecycleActionTenantScoped:              true,
			StaleRevokedHandlingProven:               true,
			DiagnosticOutputComplete:                 true,
			ProjectionDisclaimer:                     disclaimer,
		},
		FairShareQuota: DeploymentMultiTenantValBFairShareQuotaPolicy{
			CurrentState:                              DeploymentMultiTenantValBFairShareQuotaStateActive,
			EvidenceRefs:                              append([]string{}, deploymentMultiTenantValBFairShareEvidenceRefs()...),
			FreshnessState:                            IntelligenceCalibrationFreshnessFresh,
			TenantScope:                               "tenant:alpha",
			EventBudgetPerTenantPresent:               true,
			QueueIsolationPresent:                     true,
			NoisyTenantContainmentPresent:             true,
			AlertFloodThrottlingPresent:               true,
			NoStarvationRulePresent:                   true,
			OverloadDowngradeSemanticsPresent:         true,
			PerTenantRateLimitEvidencePresent:         true,
			PerTenantBackpressureSemanticsPresent:     true,
			BoundedDegradationStatePresent:            true,
			TenantAwareNegativeTestPresent:            true,
			VerificationIngestAuditIsolationUnderLoad: true,
			DiagnosticOutputComplete:                  true,
			RunbookWordingComplete:                    true,
			ProjectionDisclaimer:                      disclaimer,
		},
		NoOverclaim: DeploymentMultiTenantValBNoOverclaimDiscipline{
			CurrentState:         DeploymentMultiTenantValBNoOverclaimStateActive,
			ProjectionDisclaimer: disclaimer,
		},
		ClosureBlockerOverlay: DeploymentMultiTenantValBClosureBlockerOverlay{
			CurrentState:         DeploymentMultiTenantValBClosureBlockerStateActive,
			ProjectionDisclaimer: disclaimer,
		},
	}
}

func ComputeDeploymentMultiTenantValBFoundation(model DeploymentMultiTenantValBFoundation) DeploymentMultiTenantValBFoundation {
	model.DependencyState = EvaluateDeploymentMultiTenantValBDependencyState(model.Dependency)
	model.TenantIsolationState = EvaluateDeploymentMultiTenantValBTenantIsolationState(model.TenantIsolation)
	model.DataResidencyState = EvaluateDeploymentMultiTenantValBDataResidencyState(model.DataResidency)
	model.TenantLifecycleState = EvaluateDeploymentMultiTenantValBTenantLifecycleState(model.TenantLifecycle)
	model.FairShareQuotaState = EvaluateDeploymentMultiTenantValBFairShareQuotaState(model.FairShareQuota)
	model.NoOverclaimState = EvaluateDeploymentMultiTenantValBNoOverclaimState(model.NoOverclaim)
	model.ClosureBlockerOverlay = DeploymentMultiTenantValBClosureBlockerOverlay{
		ProjectionDisclaimer: deploymentMultiTenantValBProjectionDisclaimer(),
		Findings:             deploymentMultiTenantValBClosureBlockerFindings(model),
	}
	model.ClosureBlockerState = EvaluateDeploymentMultiTenantValBClosureBlockerState(model.ClosureBlockerOverlay)
	model.ClosureBlockerOverlay.CurrentState = model.ClosureBlockerState
	model.Point10State = EvaluateDeploymentMultiTenantPoint10State(model.CurrentState)
	model.CurrentState = EvaluateDeploymentMultiTenantValBState(model)
	model.Point10State = EvaluateDeploymentMultiTenantPoint10State(model.CurrentState)
	model.BlockingReasons = deploymentMultiTenantValBBlockingReasons(model)
	model.TenantIsolation.CurrentState = model.TenantIsolationState
	model.DataResidency.CurrentState = model.DataResidencyState
	model.TenantLifecycle.CurrentState = model.TenantLifecycleState
	model.FairShareQuota.CurrentState = model.FairShareQuotaState
	model.NoOverclaim.CurrentState = model.NoOverclaimState
	return model
}

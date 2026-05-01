package operability

import "strings"

const (
	DeploymentMultiTenantValCStateActive  = "deployment_multi_tenant_valc_active"
	DeploymentMultiTenantValCStateBlocked = "deployment_multi_tenant_valc_blocked"

	DeploymentMultiTenantValCDependencyStateActive  = "deployment_multi_tenant_valc_dependency_active"
	DeploymentMultiTenantValCDependencyStateBlocked = "deployment_multi_tenant_valc_dependency_blocked"

	DeploymentMultiTenantValCHAReadinessStateActive  = "deployment_multi_tenant_valc_ha_readiness_active"
	DeploymentMultiTenantValCHAReadinessStateBlocked = "deployment_multi_tenant_valc_ha_readiness_blocked"

	DeploymentMultiTenantValCRecoveryReadinessStateActive  = "deployment_multi_tenant_valc_recovery_readiness_active"
	DeploymentMultiTenantValCRecoveryReadinessStateBlocked = "deployment_multi_tenant_valc_recovery_readiness_blocked"

	DeploymentMultiTenantValCSLAReadinessStateActive  = "deployment_multi_tenant_valc_sla_readiness_active"
	DeploymentMultiTenantValCSLAReadinessStateBlocked = "deployment_multi_tenant_valc_sla_readiness_blocked"

	DeploymentMultiTenantValCTenantTrustScopeStateActive  = "deployment_multi_tenant_valc_tenant_trust_scope_active"
	DeploymentMultiTenantValCTenantTrustScopeStateBlocked = "deployment_multi_tenant_valc_tenant_trust_scope_blocked"

	DeploymentMultiTenantValCSiloVisibilityStateActive  = "deployment_multi_tenant_valc_silo_visibility_active"
	DeploymentMultiTenantValCSiloVisibilityStateBlocked = "deployment_multi_tenant_valc_silo_visibility_blocked"

	DeploymentMultiTenantValCPrivacyGuardStateActive  = "deployment_multi_tenant_valc_privacy_guard_active"
	DeploymentMultiTenantValCPrivacyGuardStateBlocked = "deployment_multi_tenant_valc_privacy_guard_blocked"

	DeploymentMultiTenantValCNoOverclaimStateActive  = "deployment_multi_tenant_valc_no_overclaim_active"
	DeploymentMultiTenantValCNoOverclaimStateBlocked = "deployment_multi_tenant_valc_no_overclaim_blocked"

	DeploymentMultiTenantValCClosureBlockerStateActive   = "deployment_multi_tenant_valc_closure_blocker_active"
	DeploymentMultiTenantValCClosureBlockerStateCleanup  = "deployment_multi_tenant_valc_closure_blocker_cleanup"
	DeploymentMultiTenantValCClosureBlockerStateAdvisory = "deployment_multi_tenant_valc_closure_blocker_advisory"
	DeploymentMultiTenantValCClosureBlockerStateBlocked  = "deployment_multi_tenant_valc_closure_blocker_blocked"

	DeploymentMultiTenantValCBlockerLevelCLB0 = "CL-B0"
	DeploymentMultiTenantValCBlockerLevelCLB1 = "CL-B1"
	DeploymentMultiTenantValCBlockerLevelCLB2 = "CL-B2"
	DeploymentMultiTenantValCBlockerLevelCLB3 = "CL-B3"

	DeploymentMultiTenantValCClosureSurfaceHAReadiness       = "ha_readiness"
	DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness = "recovery_readiness"
	DeploymentMultiTenantValCClosureSurfaceSLAReadiness      = "sla_readiness"
	DeploymentMultiTenantValCClosureSurfaceTenantTrustScope  = "tenant_trust_scope"
	DeploymentMultiTenantValCClosureSurfaceSiloVisibility    = "silo_visibility"
	DeploymentMultiTenantValCClosureSurfacePrivacyGuard      = "privacy_guard"
	DeploymentMultiTenantValCClosureSurfaceNoOverclaim       = "no_overclaim"
	DeploymentMultiTenantValCClosureSurfaceCleanRoomIP       = "clean_room_ip"
)

type DeploymentMultiTenantValCDependencySnapshot struct {
	ValBCurrentState         string `json:"valb_current_state"`
	ValBDependencyState      string `json:"valb_dependency_state"`
	ValBTenantIsolationState string `json:"valb_tenant_isolation_state"`
	ValBDataResidencyState   string `json:"valb_data_residency_state"`
	ValBTenantLifecycleState string `json:"valb_tenant_lifecycle_state"`
	ValBFairShareQuotaState  string `json:"valb_fair_share_quota_state"`
	ValBNoOverclaimState     string `json:"valb_no_overclaim_state"`
	ValBClosureBlockerState  string `json:"valb_closure_blocker_state"`
	Point10State             string `json:"point_10_state"`
	ProjectionDisclaimer     string `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValCHAReadiness struct {
	CurrentState                         string   `json:"current_state"`
	EvidenceRefs                         []string `json:"evidence_refs,omitempty"`
	FreshnessState                       string   `json:"freshness_state"`
	HAReadinessEvidenceLinked            bool     `json:"ha_readiness_evidence_linked"`
	TopologyEvidence                     string   `json:"topology_evidence"`
	FailoverTestEvidence                 string   `json:"failover_test_evidence"`
	DependencyDegradationBehavior        string   `json:"dependency_degradation_behavior"`
	HealthcheckStateModel                string   `json:"healthcheck_state_model"`
	QueueWorkerRecoveryBehavior          string   `json:"queue_worker_recovery_behavior"`
	DegradedModeSemantics                string   `json:"degraded_mode_semantics"`
	MonitoringAlertRoutingEvidence       string   `json:"monitoring_alert_routing_evidence"`
	TenantAwareHAImpactBoundary          string   `json:"tenant_aware_ha_impact_boundary"`
	FailoverConfiguredTreatedAsReady     bool     `json:"failover_configured_treated_as_ready"`
	HealthcheckGreenTreatedAsFullyReady  bool     `json:"healthcheck_green_treated_as_fully_ready"`
	HAReadinessTreatedAsUptimeGuarantee  bool     `json:"ha_readiness_treated_as_uptime_guarantee"`
	HACertifiedClaim                     bool     `json:"ha_certified_claim"`
	HAProfileNamingExact                 bool     `json:"ha_profile_naming_exact"`
	SafeHAReadinessWordingExamplePresent bool     `json:"safe_ha_readiness_wording_example_present"`
	DiagnosticOutputComplete             bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValCRecoveryReadiness struct {
	CurrentState                         string   `json:"current_state"`
	EvidenceRefs                         []string `json:"evidence_refs,omitempty"`
	FreshnessState                       string   `json:"freshness_state"`
	BackupFreshnessEvidence              string   `json:"backup_freshness_evidence"`
	BackupEvidenceFreshnessState         string   `json:"backup_evidence_freshness_state"`
	StaleBackupEvidenceHandlingProven    bool     `json:"stale_backup_evidence_handling_proven"`
	RestoreTestEvidence                  string   `json:"restore_test_evidence"`
	TenantScopedRestoreTest              string   `json:"tenant_scoped_restore_test"`
	RestoreIntegrityHash                 string   `json:"restore_integrity_hash"`
	EncryptedBackupCustodyReference      string   `json:"encrypted_backup_custody_reference"`
	DRDrillEvidence                      string   `json:"dr_drill_evidence"`
	RecoveryDependencyInventory          string   `json:"recovery_dependency_inventory"`
	RestoreTargetBoundary                string   `json:"restore_target_boundary"`
	BackupRetentionClass                 string   `json:"backup_retention_class"`
	DisposalDeletionBoundary             string   `json:"disposal_deletion_boundary"`
	RPORTOTarget                         string   `json:"rpo_rto_target"`
	RPORTOTreatedAsGuarantee             bool     `json:"rpo_rto_treated_as_guarantee"`
	BackupExistsTreatedAsReady           bool     `json:"backup_exists_treated_as_ready"`
	RestoreGuaranteedClaim               bool     `json:"restore_guaranteed_claim"`
	DRGuaranteedClaim                    bool     `json:"dr_guaranteed_claim"`
	BackupRestoreBypassesTenantIsolation bool     `json:"backup_restore_bypasses_tenant_isolation"`
	BackupRestoreBypassesDataResidency   bool     `json:"backup_restore_bypasses_data_residency"`
	RecoveryTargetNamingExact            bool     `json:"recovery_target_naming_exact"`
	RunbookWordingComplete               bool     `json:"runbook_wording_complete"`
	DiagnosticOutputComplete             bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValCSLAReadiness struct {
	CurrentState                             string   `json:"current_state"`
	EvidenceRefs                             []string `json:"evidence_refs,omitempty"`
	FreshnessState                           string   `json:"freshness_state"`
	SupportabilityEvidence                   string   `json:"supportability_evidence"`
	MonitoringCoverageEvidence               string   `json:"monitoring_coverage_evidence"`
	AlertRoutingEvidence                     string   `json:"alert_routing_evidence"`
	IncidentEscalationPath                   string   `json:"incident_escalation_path"`
	SupportScope                             string   `json:"support_scope"`
	TenantImpactBoundary                     string   `json:"tenant_impact_boundary"`
	DegradedModeBehavior                     string   `json:"degraded_mode_behavior"`
	RPORTOTargetReference                    string   `json:"rpo_rto_target_reference"`
	KnownLimitations                         string   `json:"known_limitations"`
	NoUptimeGuaranteeWordingPresent          bool     `json:"no_uptime_guarantee_wording_present"`
	SLAReadinessTreatedAsUptimeGuarantee     bool     `json:"sla_readiness_treated_as_uptime_guarantee"`
	SLAReadinessTreatedAsProductionApproval  bool     `json:"sla_readiness_treated_as_production_approval"`
	SLAReadinessTreatedAsComplianceReadiness bool     `json:"sla_readiness_treated_as_compliance_readiness"`
	GuaranteedUptimeClaim                    bool     `json:"guaranteed_uptime_claim"`
	ZeroDowntimeClaim                        bool     `json:"zero_downtime_claim"`
	ProductionSLAApprovedClaim               bool     `json:"production_sla_approved_claim"`
	SLAGuaranteedClaim                       bool     `json:"sla_guaranteed_claim"`
	MonitoringSummaryCanonicalTruth          bool     `json:"monitoring_summary_canonical_truth"`
	SafeSLAReadinessWordingExamplePresent    bool     `json:"safe_sla_readiness_wording_example_present"`
	DiagnosticOutputComplete                 bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer                     string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValCTenantTrustScope struct {
	CurrentState                 string   `json:"current_state"`
	EvidenceRefs                 []string `json:"evidence_refs,omitempty"`
	FreshnessState               string   `json:"freshness_state"`
	TenantTrustScope             string   `json:"tenant_trust_scope"`
	IssuerTrustOwnership         string   `json:"issuer_trust_ownership"`
	VerificationBoundary         string   `json:"verification_boundary"`
	KeyCustodyReference          string   `json:"key_custody_reference"`
	KeyCustodyOwner              string   `json:"key_custody_owner"`
	RotationBehavior             string   `json:"rotation_behavior"`
	RotationState                string   `json:"rotation_state"`
	OffboardingTransferBehavior  string   `json:"offboarding_transfer_behavior"`
	RevocationBehavior           string   `json:"revocation_behavior"`
	TenantTrustExportBoundary    string   `json:"tenant_trust_export_boundary"`
	DashboardViewOnly            bool     `json:"dashboard_view_only"`
	FleetViewOnly                bool     `json:"fleet_view_only"`
	TenantTrustScopeOfficialAuth bool     `json:"tenant_trust_scope_official_authority"`
	TrustScopeNamingExact        bool     `json:"trust_scope_naming_exact"`
	DiagnosticOutputComplete     bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValCSiloVisibility struct {
	CurrentState                            string   `json:"current_state"`
	EvidenceRefs                            []string `json:"evidence_refs,omitempty"`
	FreshnessState                          string   `json:"freshness_state"`
	EvidenceSiloValidation                  bool     `json:"evidence_silo_validation"`
	AuditSiloValidation                     bool     `json:"audit_silo_validation"`
	ExportSiloValidation                    bool     `json:"export_silo_validation"`
	SupportVisibilityBoundary               string   `json:"support_visibility_boundary"`
	TenantScopedEvidenceNamespace           string   `json:"tenant_scoped_evidence_namespace"`
	TenantScopedAuditNamespace              string   `json:"tenant_scoped_audit_namespace"`
	TenantScopedExportBoundary              string   `json:"tenant_scoped_export_boundary"`
	SupportAccessVisibilityRules            string   `json:"support_access_visibility_rules"`
	RedactionBoundary                       string   `json:"redaction_boundary"`
	ExportSiloExactIdentity                 string   `json:"export_silo_exact_identity"`
	SupportVisibilityExceedsTenantScope     bool     `json:"support_visibility_exceeds_tenant_scope"`
	RawEvidenceExposedThroughSupportSurface bool     `json:"raw_evidence_exposed_through_support_surface"`
	RedactionHidesMissingDecisiveEvidence   bool     `json:"redaction_hides_missing_decisive_evidence"`
	RedactionStrengthensClaim               bool     `json:"redaction_strengthens_claim"`
	ProjectionSurfaceCanonicalTruth         bool     `json:"projection_surface_canonical_truth"`
	DiagnosticOutputComplete                bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer                    string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValCPrivacyGuard struct {
	CurrentState                           string   `json:"current_state"`
	EvidenceRefs                           []string `json:"evidence_refs,omitempty"`
	FreshnessState                         string   `json:"freshness_state"`
	CrossTenantPrivacyGuardEvidence        string   `json:"cross_tenant_privacy_guard_evidence"`
	SideChannelTelemetryPolicy             string   `json:"side_channel_telemetry_policy"`
	VolumeLeakageCheckPresent              bool     `json:"volume_leakage_check_present"`
	ErrorLeakageCheckPresent               bool     `json:"error_leakage_check_present"`
	TimingLeakageCheckPresent              bool     `json:"timing_leakage_check_present"`
	AggregationLeakageCheckPresent         bool     `json:"aggregation_leakage_check_present"`
	BoundedAggregationRules                string   `json:"bounded_aggregation_rules"`
	TenantPrivateMetadataClassification    string   `json:"tenant_private_metadata_classification"`
	SupportExportPrivacyVisibilityBoundary string   `json:"support_export_privacy_visibility_boundary"`
	VolumeLeakageNegativeTestPresent       bool     `json:"volume_leakage_negative_test_present"`
	ErrorLeakageNegativeTestPresent        bool     `json:"error_leakage_negative_test_present"`
	TimingLeakageNegativeTestPresent       bool     `json:"timing_leakage_negative_test_present"`
	AggregationLeakageNegativeTestPresent  bool     `json:"aggregation_leakage_negative_test_present"`
	SideChannelMarkedSafeWithoutEvidence   bool     `json:"side_channel_marked_safe_without_evidence"`
	TenantPrivateMetadataLeakage           bool     `json:"tenant_private_metadata_leakage"`
	FleetAggregationCanonicalTruth         bool     `json:"fleet_aggregation_canonical_truth"`
	DiagnosticOutputComplete               bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer                   string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValCNoOverclaimDiscipline struct {
	CurrentState                 string   `json:"current_state"`
	ObservedClaims               []string `json:"observed_claims,omitempty"`
	CleanRoomIPViolationDetected bool     `json:"clean_room_ip_violation_detected"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValCClosureBlockerFinding struct {
	BlockerLevel      string `json:"blocker_level"`
	Surface           string `json:"surface"`
	Reason            string `json:"reason"`
	BlocksCurrentWave bool   `json:"blocks_current_wave"`
	RequiredFollowup  string `json:"required_followup,omitempty"`
}

type DeploymentMultiTenantValCClosureBlockerOverlay struct {
	CurrentState         string                                           `json:"current_state"`
	Findings             []DeploymentMultiTenantValCClosureBlockerFinding `json:"findings,omitempty"`
	ProjectionDisclaimer string                                           `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValCFoundation struct {
	CurrentState           string                                         `json:"current_state"`
	Point10State           string                                         `json:"point_10_state"`
	BlockingReasons        []string                                       `json:"blocking_reasons,omitempty"`
	DependencyState        string                                         `json:"dependency_state"`
	HAReadinessState       string                                         `json:"ha_readiness_state"`
	RecoveryReadinessState string                                         `json:"recovery_readiness_state"`
	SLAReadinessState      string                                         `json:"sla_readiness_state"`
	TenantTrustScopeState  string                                         `json:"tenant_trust_scope_state"`
	SiloVisibilityState    string                                         `json:"silo_visibility_state"`
	PrivacyGuardState      string                                         `json:"privacy_guard_state"`
	NoOverclaimState       string                                         `json:"no_overclaim_state"`
	ClosureBlockerState    string                                         `json:"closure_blocker_state"`
	Dependency             DeploymentMultiTenantValCDependencySnapshot    `json:"dependency"`
	HAReadiness            DeploymentMultiTenantValCHAReadiness           `json:"ha_readiness"`
	RecoveryReadiness      DeploymentMultiTenantValCRecoveryReadiness     `json:"recovery_readiness"`
	SLAReadiness           DeploymentMultiTenantValCSLAReadiness          `json:"sla_readiness"`
	TenantTrustScope       DeploymentMultiTenantValCTenantTrustScope      `json:"tenant_trust_scope"`
	SiloVisibility         DeploymentMultiTenantValCSiloVisibility        `json:"silo_visibility"`
	PrivacyGuard           DeploymentMultiTenantValCPrivacyGuard          `json:"privacy_guard"`
	NoOverclaim            DeploymentMultiTenantValCNoOverclaimDiscipline `json:"no_overclaim"`
	ClosureBlockerOverlay  DeploymentMultiTenantValCClosureBlockerOverlay `json:"closure_blocker_overlay"`
}

func deploymentMultiTenantValCProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_marketplace_deployment_profile deployment_multi_tenant_valc"
}

func deploymentMultiTenantValCHasProjectionDisclaimer(value string) bool {
	normalized := strings.ToLower(strings.TrimSpace(value))
	return strings.Contains(normalized, "projection_only") &&
		strings.Contains(normalized, "not_canonical_truth") &&
		strings.Contains(normalized, "deployment_multi_tenant_valc")
}

func deploymentMultiTenantValCHAEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-valc-ha-readiness-001"}
}

func deploymentMultiTenantValCRecoveryEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-valc-recovery-readiness-001"}
}

func deploymentMultiTenantValCSLAEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-valc-sla-readiness-001"}
}

func deploymentMultiTenantValCTrustScopeEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-valc-trust-scope-001"}
}

func deploymentMultiTenantValCSiloVisibilityEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-valc-silo-visibility-001"}
}

func deploymentMultiTenantValCPrivacyGuardEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-valc-privacy-guard-001"}
}

func deploymentMultiTenantValCEvidenceValueIsValid(value string) bool {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" || trimmed != strings.ToLower(trimmed) {
		return false
	}
	if strings.Contains(trimmed, "*") {
		return false
	}
	normalized := deploymentMultiTenantVal0NormalizeClaimText(trimmed)
	compact := deploymentMultiTenantVal0CompactClaimText(trimmed)
	if normalized == "" || compact == "" {
		return false
	}
	for _, token := range strings.Fields(normalized) {
		switch token {
		case "unknown", "partial", "incomplete", "stale", "malformed", "unsupported", "blocked", "revoked", "expired", "duplicate", "unrelated":
			return false
		}
		if token == "ish" || strings.HasSuffix(token, "ish") {
			return false
		}
	}
	return true
}

func deploymentMultiTenantValCHasRevokedExpiredDuplicateOrUnrelatedEvidenceToken(values ...string) bool {
	for _, value := range values {
		normalized := deploymentMultiTenantVal0NormalizeClaimText(value)
		if normalized == "" {
			continue
		}
		for _, token := range strings.Fields(normalized) {
			switch token {
			case "revoked", "expired", "duplicate", "unrelated":
				return true
			}
		}
	}
	return false
}

func deploymentMultiTenantValCClosureBlockerLevels() []string {
	return []string{
		DeploymentMultiTenantValCBlockerLevelCLB0,
		DeploymentMultiTenantValCBlockerLevelCLB1,
		DeploymentMultiTenantValCBlockerLevelCLB2,
		DeploymentMultiTenantValCBlockerLevelCLB3,
	}
}

func deploymentMultiTenantValCClosureBlockerSurfaces() []string {
	return []string{
		DeploymentMultiTenantValCClosureSurfaceHAReadiness,
		DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness,
		DeploymentMultiTenantValCClosureSurfaceSLAReadiness,
		DeploymentMultiTenantValCClosureSurfaceTenantTrustScope,
		DeploymentMultiTenantValCClosureSurfaceSiloVisibility,
		DeploymentMultiTenantValCClosureSurfacePrivacyGuard,
		DeploymentMultiTenantValCClosureSurfaceNoOverclaim,
		DeploymentMultiTenantValCClosureSurfaceCleanRoomIP,
	}
}

func deploymentMultiTenantValCDependencySnapshotModel() DeploymentMultiTenantValCDependencySnapshot {
	valB := ComputeDeploymentMultiTenantValBFoundation(DeploymentMultiTenantValBFoundationModel())
	return DeploymentMultiTenantValCDependencySnapshot{
		ValBCurrentState:         valB.CurrentState,
		ValBDependencyState:      valB.DependencyState,
		ValBTenantIsolationState: valB.TenantIsolationState,
		ValBDataResidencyState:   valB.DataResidencyState,
		ValBTenantLifecycleState: valB.TenantLifecycleState,
		ValBFairShareQuotaState:  valB.FairShareQuotaState,
		ValBNoOverclaimState:     valB.NoOverclaimState,
		ValBClosureBlockerState:  valB.ClosureBlockerState,
		Point10State:             valB.Point10State,
		ProjectionDisclaimer:     deploymentMultiTenantValBProjectionDisclaimer(),
	}
}

func EvaluateDeploymentMultiTenantValCDependencyState(model DeploymentMultiTenantValCDependencySnapshot) string {
	if !deploymentMultiTenantValBHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeploymentMultiTenantValCDependencyStateBlocked
	}
	if strings.TrimSpace(model.ValBCurrentState) != DeploymentMultiTenantValBStateActive ||
		strings.TrimSpace(model.ValBDependencyState) != DeploymentMultiTenantValBDependencyStateActive ||
		strings.TrimSpace(model.ValBTenantIsolationState) != DeploymentMultiTenantValBTenantIsolationStateActive ||
		strings.TrimSpace(model.ValBDataResidencyState) != DeploymentMultiTenantValBDataResidencyStateActive ||
		strings.TrimSpace(model.ValBTenantLifecycleState) != DeploymentMultiTenantValBTenantLifecycleStateActive ||
		strings.TrimSpace(model.ValBFairShareQuotaState) != DeploymentMultiTenantValBFairShareQuotaStateActive ||
		strings.TrimSpace(model.ValBNoOverclaimState) != DeploymentMultiTenantValBNoOverclaimStateActive ||
		strings.TrimSpace(model.ValBClosureBlockerState) != DeploymentMultiTenantValBClosureBlockerStateActive ||
		strings.TrimSpace(model.Point10State) != DeploymentMultiTenantPoint10StateNotComplete {
		return DeploymentMultiTenantValCDependencyStateBlocked
	}
	return DeploymentMultiTenantValCDependencyStateActive
}

func EvaluateDeploymentMultiTenantValCHAReadinessState(model DeploymentMultiTenantValCHAReadiness) string {
	if !deploymentMultiTenantValCHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, deploymentMultiTenantValCHAEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!model.HAReadinessEvidenceLinked ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.TopologyEvidence) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.FailoverTestEvidence) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.DependencyDegradationBehavior) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.HealthcheckStateModel) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.QueueWorkerRecoveryBehavior) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.DegradedModeSemantics) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.MonitoringAlertRoutingEvidence) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantAwareHAImpactBoundary) ||
		model.FailoverConfiguredTreatedAsReady ||
		model.HealthcheckGreenTreatedAsFullyReady ||
		model.HAReadinessTreatedAsUptimeGuarantee ||
		model.HACertifiedClaim {
		return DeploymentMultiTenantValCHAReadinessStateBlocked
	}
	return DeploymentMultiTenantValCHAReadinessStateActive
}

func EvaluateDeploymentMultiTenantValCRecoveryReadinessState(model DeploymentMultiTenantValCRecoveryReadiness) string {
	if !deploymentMultiTenantValCHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, deploymentMultiTenantValCRecoveryEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.BackupFreshnessEvidence) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.BackupEvidenceFreshnessState) ||
		!model.StaleBackupEvidenceHandlingProven ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.RestoreTestEvidence) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantScopedRestoreTest) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.RestoreIntegrityHash) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.EncryptedBackupCustodyReference) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.DRDrillEvidence) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.RecoveryDependencyInventory) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.RestoreTargetBoundary) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.BackupRetentionClass) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.DisposalDeletionBoundary) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.RPORTOTarget) ||
		model.RPORTOTreatedAsGuarantee ||
		model.BackupExistsTreatedAsReady ||
		model.RestoreGuaranteedClaim ||
		model.DRGuaranteedClaim ||
		model.BackupRestoreBypassesTenantIsolation ||
		model.BackupRestoreBypassesDataResidency {
		return DeploymentMultiTenantValCRecoveryReadinessStateBlocked
	}
	return DeploymentMultiTenantValCRecoveryReadinessStateActive
}

func EvaluateDeploymentMultiTenantValCSLAReadinessState(model DeploymentMultiTenantValCSLAReadiness) string {
	if !deploymentMultiTenantValCHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, deploymentMultiTenantValCSLAEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.SupportabilityEvidence) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.MonitoringCoverageEvidence) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.AlertRoutingEvidence) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.IncidentEscalationPath) ||
		!deploymentMultiTenantVal0TenantScopedValueIsValid(model.SupportScope) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantImpactBoundary) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.DegradedModeBehavior) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.RPORTOTargetReference) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.KnownLimitations) ||
		!model.NoUptimeGuaranteeWordingPresent ||
		model.SLAReadinessTreatedAsUptimeGuarantee ||
		model.SLAReadinessTreatedAsProductionApproval ||
		model.SLAReadinessTreatedAsComplianceReadiness ||
		model.GuaranteedUptimeClaim ||
		model.ZeroDowntimeClaim ||
		model.ProductionSLAApprovedClaim ||
		model.SLAGuaranteedClaim ||
		model.MonitoringSummaryCanonicalTruth {
		return DeploymentMultiTenantValCSLAReadinessStateBlocked
	}
	return DeploymentMultiTenantValCSLAReadinessStateActive
}

func EvaluateDeploymentMultiTenantValCTenantTrustScopeState(model DeploymentMultiTenantValCTenantTrustScope) string {
	if !deploymentMultiTenantValCHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, deploymentMultiTenantValCTrustScopeEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!deploymentMultiTenantVal0TenantScopedValueIsValid(model.TenantTrustScope) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.IssuerTrustOwnership) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.VerificationBoundary) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.KeyCustodyReference) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.KeyCustodyOwner) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.RotationBehavior) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.RotationState) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.OffboardingTransferBehavior) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.RevocationBehavior) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantTrustExportBoundary) ||
		model.DashboardViewOnly ||
		model.FleetViewOnly ||
		model.TenantTrustScopeOfficialAuth {
		return DeploymentMultiTenantValCTenantTrustScopeStateBlocked
	}
	return DeploymentMultiTenantValCTenantTrustScopeStateActive
}

func EvaluateDeploymentMultiTenantValCSiloVisibilityState(model DeploymentMultiTenantValCSiloVisibility) string {
	if !deploymentMultiTenantValCHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, deploymentMultiTenantValCSiloVisibilityEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!model.EvidenceSiloValidation ||
		!model.AuditSiloValidation ||
		!model.ExportSiloValidation ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.SupportVisibilityBoundary) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantScopedEvidenceNamespace) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantScopedAuditNamespace) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantScopedExportBoundary) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.SupportAccessVisibilityRules) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.RedactionBoundary) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.ExportSiloExactIdentity) ||
		model.SupportVisibilityExceedsTenantScope ||
		model.RawEvidenceExposedThroughSupportSurface ||
		model.RedactionHidesMissingDecisiveEvidence ||
		model.RedactionStrengthensClaim ||
		model.ProjectionSurfaceCanonicalTruth {
		return DeploymentMultiTenantValCSiloVisibilityStateBlocked
	}
	return DeploymentMultiTenantValCSiloVisibilityStateActive
}

func EvaluateDeploymentMultiTenantValCPrivacyGuardState(model DeploymentMultiTenantValCPrivacyGuard) string {
	if !deploymentMultiTenantValCHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, deploymentMultiTenantValCPrivacyGuardEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.CrossTenantPrivacyGuardEvidence) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.SideChannelTelemetryPolicy) ||
		!model.VolumeLeakageCheckPresent ||
		!model.ErrorLeakageCheckPresent ||
		!model.TimingLeakageCheckPresent ||
		!model.AggregationLeakageCheckPresent ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.BoundedAggregationRules) ||
		!deploymentMultiTenantValCEvidenceValueIsValid(model.TenantPrivateMetadataClassification) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.SupportExportPrivacyVisibilityBoundary) ||
		!model.VolumeLeakageNegativeTestPresent ||
		!model.ErrorLeakageNegativeTestPresent ||
		!model.TimingLeakageNegativeTestPresent ||
		!model.AggregationLeakageNegativeTestPresent ||
		model.SideChannelMarkedSafeWithoutEvidence ||
		model.TenantPrivateMetadataLeakage ||
		model.FleetAggregationCanonicalTruth {
		return DeploymentMultiTenantValCPrivacyGuardStateBlocked
	}
	return DeploymentMultiTenantValCPrivacyGuardStateActive
}

func deploymentMultiTenantValCContainsForbiddenClaim(values ...string) bool {
	allowed := []string{
		"ha readiness evidence",
		"failover test evidence",
		"dependency degradation behavior",
		"degraded mode semantics",
		"backup freshness evidence",
		"restore test evidence",
		"tenant-scoped restore test",
		"restore integrity hash",
		"encrypted backup custody reference",
		"dr drill evidence",
		"rpo/rto target",
		"sla readiness evidence",
		"supportability evidence",
		"known limitations",
		"tenant trust scope evidence",
		"issuer/trust ownership evidence",
		"key/custody rotation evidence",
		"evidence silo validation",
		"audit silo validation",
		"export silo validation",
		"bounded support visibility",
		"privacy guard evidence",
		"side-channel negative test",
		"bounded aggregation rules",
		"advisory fleet visibility",
		"not uptime guarantee",
		"not production approval",
		"not deployment approval",
		"not compliance certification",
		"not canonical truth",
	}
	disallowed := []string{
		"guaranteed uptime",
		"zero downtime",
		"sla guaranteed",
		"production sla approved",
		"ha certified",
		"ha guaranteed",
		"failover guaranteed",
		"restore guaranteed",
		"dr guaranteed",
		"disaster recovery guaranteed",
		"backup guarantees recovery",
		"restore always works",
		"backup exists means ready",
		"healthcheck green means fully ready",
		"failover configured means ready",
		"sla readiness means uptime guarantee",
		"sla readiness means production approval",
		"supportability evidence means sla guarantee",
		"tenant trust certified",
		"tenant trust scope certified",
		"key custody certified",
		"data residency certified",
		"tenant isolation guaranteed",
		"privacy guaranteed",
		"no side-channel leakage guaranteed",
		"fleet aggregation proves privacy",
		"support visibility cannot leak",
		"redaction proves safe",
		"portal view is canonical truth",
		"dashboard proves recovery readiness",
		"clean-room certified",
		"patent cleared",
		"fto cleared",
		"legal certification",
		"copied competitor workflow",
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
	for _, value := range values {
		normalized := deploymentMultiTenantVal0NormalizeClaimText(value)
		compact := deploymentMultiTenantVal0CompactClaimText(value)
		if normalized == "" && compact == "" {
			continue
		}
		if _, ok := allowedExact[normalized]; ok {
			continue
		}
		for i := range blockedNormalized {
			if strings.Contains(normalized, blockedNormalized[i]) || strings.Contains(compact, blockedCompact[i]) {
				return true
			}
		}
	}
	return false
}

func EvaluateDeploymentMultiTenantValCNoOverclaimState(model DeploymentMultiTenantValCNoOverclaimDiscipline) string {
	if !deploymentMultiTenantValCHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		model.CleanRoomIPViolationDetected ||
		deploymentMultiTenantValCContainsForbiddenClaim(model.ObservedClaims...) {
		return DeploymentMultiTenantValCNoOverclaimStateBlocked
	}
	return DeploymentMultiTenantValCNoOverclaimStateActive
}

func deploymentMultiTenantValCClosureBlockerFinding(level, surface, reason string, blocksCurrentWave bool, requiredFollowup string) DeploymentMultiTenantValCClosureBlockerFinding {
	return DeploymentMultiTenantValCClosureBlockerFinding{
		BlockerLevel:      level,
		Surface:           surface,
		Reason:            reason,
		BlocksCurrentWave: blocksCurrentWave,
		RequiredFollowup:  requiredFollowup,
	}
}

func deploymentMultiTenantValCClosureBlockerFindings(model DeploymentMultiTenantValCFoundation) []DeploymentMultiTenantValCClosureBlockerFinding {
	findings := []DeploymentMultiTenantValCClosureBlockerFinding{}
	if model.SLAReadiness.SLAReadinessTreatedAsUptimeGuarantee {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceSLAReadiness, "sla readiness treated as uptime guarantee", true, "remove uptime guarantee interpretation from sla readiness and rerun supportability validation"))
	}
	if model.SLAReadiness.GuaranteedUptimeClaim {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceSLAReadiness, "guaranteed uptime claim", true, "remove guaranteed uptime claim and keep sla readiness bounded"))
	}
	if model.SLAReadiness.ZeroDowntimeClaim {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceSLAReadiness, "zero downtime claim", true, "remove zero downtime claim and keep degraded mode semantics explicit"))
	}
	if model.SLAReadiness.ProductionSLAApprovedClaim {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceSLAReadiness, "production sla approved claim", true, "remove production sla approval claim and keep supportability evidence bounded"))
	}
	if model.RecoveryReadiness.BackupRestoreBypassesTenantIsolation {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, "backup restore or dr readiness bypasses tenant isolation", true, "remove recovery path bypass and preserve tenant isolation boundaries"))
	}
	if model.RecoveryReadiness.BackupRestoreBypassesDataResidency {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, "backup restore or dr readiness bypasses data residency", true, "remove recovery path bypass and preserve data residency boundaries"))
	}
	if model.SiloVisibility.RawEvidenceExposedThroughSupportSurface {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceSiloVisibility, "support visibility leaks raw tenant evidence", true, "remove raw tenant evidence exposure from support visibility surfaces"))
	}
	if model.SiloVisibility.RedactionHidesMissingDecisiveEvidence {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceSiloVisibility, "redaction hides decisive missing evidence", true, "restore decisive missing evidence disclosure without using redaction to hide gaps"))
	}
	if model.SiloVisibility.RedactionStrengthensClaim {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceSiloVisibility, "redaction strengthens claim", true, "remove strengthened claim language from redacted surfaces"))
	}
	if model.PrivacyGuard.SideChannelMarkedSafeWithoutEvidence {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfacePrivacyGuard, "side-channel privacy risk marked safe without evidence", true, "add side-channel evidence or remove unsupported privacy safety claim"))
	}
	if model.PrivacyGuard.TenantPrivateMetadataLeakage {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfacePrivacyGuard, "tenant-private metadata leakage", true, "remove tenant-private metadata leakage and rerun privacy guard validation"))
	}
	if model.SiloVisibility.ProjectionSurfaceCanonicalTruth || model.SLAReadiness.MonitoringSummaryCanonicalTruth || model.PrivacyGuard.FleetAggregationCanonicalTruth {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceSiloVisibility, "projection surface treated as canonical truth", true, "restore projection-only semantics and evidence-backed canonical boundaries"))
	}
	if model.NoOverclaim.CleanRoomIPViolationDetected {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCClosureSurfaceCleanRoomIP, "copied competitor deployment recovery or privacy artifact detected", true, "remove copied artifact and replace it with clean-room implementation evidence"))
	}
	if deploymentMultiTenantValCHasRevokedExpiredDuplicateOrUnrelatedEvidenceToken(
		model.HAReadiness.TopologyEvidence,
		model.HAReadiness.FailoverTestEvidence,
		model.HAReadiness.DependencyDegradationBehavior,
		model.HAReadiness.HealthcheckStateModel,
		model.HAReadiness.QueueWorkerRecoveryBehavior,
		model.HAReadiness.DegradedModeSemantics,
		model.HAReadiness.MonitoringAlertRoutingEvidence,
	) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceHAReadiness, "revoked expired duplicate or unrelated evidence token accepted", true, "reject revoked expired duplicate unrelated evidence tokens and rerun Val C fail-closed evidence tests"))
	}
	if deploymentMultiTenantValCHasRevokedExpiredDuplicateOrUnrelatedEvidenceToken(
		model.RecoveryReadiness.BackupFreshnessEvidence,
		model.RecoveryReadiness.RestoreTestEvidence,
		model.RecoveryReadiness.RestoreIntegrityHash,
		model.RecoveryReadiness.EncryptedBackupCustodyReference,
		model.RecoveryReadiness.DRDrillEvidence,
		model.RecoveryReadiness.RecoveryDependencyInventory,
		model.RecoveryReadiness.BackupRetentionClass,
		model.RecoveryReadiness.RPORTOTarget,
	) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, "revoked expired duplicate or unrelated evidence token accepted", true, "reject revoked expired duplicate unrelated evidence tokens and rerun Val C fail-closed evidence tests"))
	}
	if deploymentMultiTenantValCHasRevokedExpiredDuplicateOrUnrelatedEvidenceToken(
		model.SLAReadiness.SupportabilityEvidence,
		model.SLAReadiness.MonitoringCoverageEvidence,
		model.SLAReadiness.AlertRoutingEvidence,
		model.SLAReadiness.IncidentEscalationPath,
		model.SLAReadiness.DegradedModeBehavior,
		model.SLAReadiness.RPORTOTargetReference,
		model.SLAReadiness.KnownLimitations,
	) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceSLAReadiness, "revoked expired duplicate or unrelated evidence token accepted", true, "reject revoked expired duplicate unrelated evidence tokens and rerun Val C fail-closed evidence tests"))
	}
	if deploymentMultiTenantValCHasRevokedExpiredDuplicateOrUnrelatedEvidenceToken(
		model.TenantTrustScope.IssuerTrustOwnership,
		model.TenantTrustScope.KeyCustodyReference,
		model.TenantTrustScope.KeyCustodyOwner,
		model.TenantTrustScope.RotationBehavior,
		model.TenantTrustScope.RotationState,
		model.TenantTrustScope.OffboardingTransferBehavior,
		model.TenantTrustScope.RevocationBehavior,
	) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceTenantTrustScope, "revoked expired duplicate or unrelated evidence token accepted", true, "reject revoked expired duplicate unrelated evidence tokens and rerun Val C fail-closed evidence tests"))
	}
	if deploymentMultiTenantValCHasRevokedExpiredDuplicateOrUnrelatedEvidenceToken(
		model.SiloVisibility.SupportAccessVisibilityRules,
		model.SiloVisibility.ExportSiloExactIdentity,
	) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceSiloVisibility, "revoked expired duplicate or unrelated evidence token accepted", true, "reject revoked expired duplicate unrelated evidence tokens and rerun Val C fail-closed evidence tests"))
	}
	if deploymentMultiTenantValCHasRevokedExpiredDuplicateOrUnrelatedEvidenceToken(
		model.PrivacyGuard.CrossTenantPrivacyGuardEvidence,
		model.PrivacyGuard.SideChannelTelemetryPolicy,
		model.PrivacyGuard.BoundedAggregationRules,
		model.PrivacyGuard.TenantPrivateMetadataClassification,
	) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfacePrivacyGuard, "revoked expired duplicate or unrelated evidence token accepted", true, "reject revoked expired duplicate unrelated evidence tokens and rerun Val C fail-closed evidence tests"))
	}

	if !deploymentMultiTenantValCEvidenceValueIsValid(model.RecoveryReadiness.BackupFreshnessEvidence) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, "backup or restore evidence missing", true, "add backup freshness evidence and bounded recovery evidence links"))
	}
	if !deploymentMultiTenantValCEvidenceValueIsValid(model.RecoveryReadiness.RestoreTestEvidence) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, "restore test evidence missing", true, "add restore test evidence before final review"))
	}
	if !model.RecoveryReadiness.StaleBackupEvidenceHandlingProven {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, "stale backup evidence handling not proven", true, "prove stale backup evidence handling and bounded downgrade behavior"))
	}
	if !model.HAReadiness.HAReadinessEvidenceLinked {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceHAReadiness, "ha readiness not evidence-linked", true, "link ha readiness to exact evidence before final review"))
	}
	if !deploymentMultiTenantValCEvidenceValueIsValid(model.HAReadiness.FailoverTestEvidence) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceHAReadiness, "failover test evidence missing", true, "add failover test evidence and rerun ha validation"))
	}
	if !deploymentMultiTenantValCEvidenceValueIsValid(model.TenantTrustScope.IssuerTrustOwnership) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceTenantTrustScope, "tenant trust scope missing issuer or trust ownership", true, "add explicit issuer and trust ownership evidence"))
	}
	if !deploymentMultiTenantValCEvidenceValueIsValid(model.TenantTrustScope.RotationBehavior) || !deploymentMultiTenantValCEvidenceValueIsValid(model.TenantTrustScope.RevocationBehavior) {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceTenantTrustScope, "key custody rotation or revocation behavior missing", true, "add key custody rotation and revocation behavior evidence"))
	}
	if !model.SiloVisibility.EvidenceSiloValidation || !model.SiloVisibility.AuditSiloValidation || !model.SiloVisibility.ExportSiloValidation {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceSiloVisibility, "evidence audit or export silo validation missing", true, "complete evidence audit and export silo validation before final review"))
	}
	if !model.PrivacyGuard.VolumeLeakageNegativeTestPresent || !model.PrivacyGuard.ErrorLeakageNegativeTestPresent || !model.PrivacyGuard.TimingLeakageNegativeTestPresent || !model.PrivacyGuard.AggregationLeakageNegativeTestPresent {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfacePrivacyGuard, "privacy side-channel negative tests missing", true, "add volume error timing and aggregation side-channel negative coverage"))
	}
	if model.DependencyState != DeploymentMultiTenantValCDependencyStateActive {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB1, DeploymentMultiTenantValCClosureSurfaceHAReadiness, "dependency gate missing or not exact active", true, "restore exact active Val B dependency before Val C final review"))
	}

	if !model.HAReadiness.HAProfileNamingExact {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfaceHAReadiness, "ambiguous ha profile naming", true, "normalize ha profile naming before handoff"))
	}
	if !model.RecoveryReadiness.RecoveryTargetNamingExact {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, "ambiguous recovery target naming", true, "normalize recovery target naming before handoff"))
	}
	if !model.TenantTrustScope.TrustScopeNamingExact {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfaceTenantTrustScope, "ambiguous trust scope naming", true, "normalize trust scope naming before handoff"))
	}
	if !model.SLAReadiness.SafeSLAReadinessWordingExamplePresent {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfaceSLAReadiness, "missing safe wording example for sla readiness", true, "add bounded safe wording example for sla readiness"))
	}
	if !model.HAReadiness.SafeHAReadinessWordingExamplePresent {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfaceHAReadiness, "missing safe wording example for ha readiness", true, "add bounded safe wording example for ha readiness"))
	}
	if !model.HAReadiness.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfaceHAReadiness, "incomplete diagnostic output for ha blockers", true, "complete ha diagnostic output before handoff"))
	}
	if !model.RecoveryReadiness.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, "incomplete diagnostic output for recovery blockers", true, "complete recovery diagnostic output before handoff"))
	}
	if !model.SLAReadiness.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfaceSLAReadiness, "incomplete diagnostic output for sla blockers", true, "complete sla diagnostic output before handoff"))
	}
	if !model.TenantTrustScope.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfaceTenantTrustScope, "incomplete diagnostic output for trust blockers", true, "complete trust scope diagnostic output before handoff"))
	}
	if !model.PrivacyGuard.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfacePrivacyGuard, "incomplete diagnostic output for privacy blockers", true, "complete privacy guard diagnostic output before handoff"))
	}
	if !model.RecoveryReadiness.RunbookWordingComplete {
		findings = append(findings, deploymentMultiTenantValCClosureBlockerFinding(DeploymentMultiTenantValCBlockerLevelCLB2, DeploymentMultiTenantValCClosureSurfaceRecoveryReadiness, "incomplete bounded runbook wording without direct pass bypass", true, "complete bounded runbook wording before handoff"))
	}
	return findings
}

func EvaluateDeploymentMultiTenantValCClosureBlockerState(model DeploymentMultiTenantValCClosureBlockerOverlay) string {
	if !deploymentMultiTenantValCHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeploymentMultiTenantValCClosureBlockerStateBlocked
	}
	hasCleanup := false
	hasAdvisory := false
	for _, finding := range model.Findings {
		level := strings.TrimSpace(finding.BlockerLevel)
		surface := strings.TrimSpace(finding.Surface)
		if len(level) == 2 && level[0] == 'P' && level[1] >= '0' && level[1] <= '9' {
			return DeploymentMultiTenantValCClosureBlockerStateBlocked
		}
		if !containsTrimmedString(deploymentMultiTenantValCClosureBlockerLevels(), level) ||
			!containsTrimmedString(deploymentMultiTenantValCClosureBlockerSurfaces(), surface) {
			return DeploymentMultiTenantValCClosureBlockerStateBlocked
		}
		if (level == DeploymentMultiTenantValCBlockerLevelCLB1 ||
			level == DeploymentMultiTenantValCBlockerLevelCLB2 ||
			level == DeploymentMultiTenantValCBlockerLevelCLB3) &&
			strings.TrimSpace(finding.RequiredFollowup) == "" {
			return DeploymentMultiTenantValCClosureBlockerStateBlocked
		}
		switch level {
		case DeploymentMultiTenantValCBlockerLevelCLB0, DeploymentMultiTenantValCBlockerLevelCLB1:
			return DeploymentMultiTenantValCClosureBlockerStateBlocked
		case DeploymentMultiTenantValCBlockerLevelCLB2:
			hasCleanup = true
		case DeploymentMultiTenantValCBlockerLevelCLB3:
			hasAdvisory = true
		default:
			return DeploymentMultiTenantValCClosureBlockerStateBlocked
		}
	}
	if hasCleanup {
		return DeploymentMultiTenantValCClosureBlockerStateCleanup
	}
	if hasAdvisory {
		return DeploymentMultiTenantValCClosureBlockerStateAdvisory
	}
	return DeploymentMultiTenantValCClosureBlockerStateActive
}

func EvaluateDeploymentMultiTenantValCState(model DeploymentMultiTenantValCFoundation) string {
	if strings.TrimSpace(model.DependencyState) != DeploymentMultiTenantValCDependencyStateActive ||
		strings.TrimSpace(model.HAReadinessState) != DeploymentMultiTenantValCHAReadinessStateActive ||
		strings.TrimSpace(model.RecoveryReadinessState) != DeploymentMultiTenantValCRecoveryReadinessStateActive ||
		strings.TrimSpace(model.SLAReadinessState) != DeploymentMultiTenantValCSLAReadinessStateActive ||
		strings.TrimSpace(model.TenantTrustScopeState) != DeploymentMultiTenantValCTenantTrustScopeStateActive ||
		strings.TrimSpace(model.SiloVisibilityState) != DeploymentMultiTenantValCSiloVisibilityStateActive ||
		strings.TrimSpace(model.PrivacyGuardState) != DeploymentMultiTenantValCPrivacyGuardStateActive ||
		strings.TrimSpace(model.NoOverclaimState) != DeploymentMultiTenantValCNoOverclaimStateActive ||
		strings.TrimSpace(model.ClosureBlockerState) != DeploymentMultiTenantValCClosureBlockerStateActive ||
		strings.TrimSpace(model.Point10State) != DeploymentMultiTenantPoint10StateNotComplete {
		return DeploymentMultiTenantValCStateBlocked
	}
	return DeploymentMultiTenantValCStateActive
}

func deploymentMultiTenantValCBlockingReasons(model DeploymentMultiTenantValCFoundation) []string {
	reasons := []string{}
	if model.DependencyState != DeploymentMultiTenantValCDependencyStateActive {
		reasons = append(reasons, "dependency_state_blocked")
	}
	if model.HAReadinessState != DeploymentMultiTenantValCHAReadinessStateActive {
		reasons = append(reasons, "ha_readiness_state_blocked")
	}
	if model.RecoveryReadinessState != DeploymentMultiTenantValCRecoveryReadinessStateActive {
		reasons = append(reasons, "recovery_readiness_state_blocked")
	}
	if model.SLAReadinessState != DeploymentMultiTenantValCSLAReadinessStateActive {
		reasons = append(reasons, "sla_readiness_state_blocked")
	}
	if model.TenantTrustScopeState != DeploymentMultiTenantValCTenantTrustScopeStateActive {
		reasons = append(reasons, "tenant_trust_scope_state_blocked")
	}
	if model.SiloVisibilityState != DeploymentMultiTenantValCSiloVisibilityStateActive {
		reasons = append(reasons, "silo_visibility_state_blocked")
	}
	if model.PrivacyGuardState != DeploymentMultiTenantValCPrivacyGuardStateActive {
		reasons = append(reasons, "privacy_guard_state_blocked")
	}
	if model.NoOverclaimState != DeploymentMultiTenantValCNoOverclaimStateActive {
		reasons = append(reasons, "no_overclaim_state_blocked")
	}
	if model.ClosureBlockerState != DeploymentMultiTenantValCClosureBlockerStateActive {
		reasons = append(reasons, "closure_blocker_state_not_clean")
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		reasons = append(reasons, "point10_state_not_complete_guard_violated")
	}
	return reasons
}

func DeploymentMultiTenantValCFoundationModel() DeploymentMultiTenantValCFoundation {
	disclaimer := deploymentMultiTenantValCProjectionDisclaimer()
	return DeploymentMultiTenantValCFoundation{
		CurrentState:           DeploymentMultiTenantValCStateActive,
		Point10State:           DeploymentMultiTenantPoint10StateNotComplete,
		DependencyState:        DeploymentMultiTenantValCDependencyStateActive,
		HAReadinessState:       DeploymentMultiTenantValCHAReadinessStateActive,
		RecoveryReadinessState: DeploymentMultiTenantValCRecoveryReadinessStateActive,
		SLAReadinessState:      DeploymentMultiTenantValCSLAReadinessStateActive,
		TenantTrustScopeState:  DeploymentMultiTenantValCTenantTrustScopeStateActive,
		SiloVisibilityState:    DeploymentMultiTenantValCSiloVisibilityStateActive,
		PrivacyGuardState:      DeploymentMultiTenantValCPrivacyGuardStateActive,
		NoOverclaimState:       DeploymentMultiTenantValCNoOverclaimStateActive,
		ClosureBlockerState:    DeploymentMultiTenantValCClosureBlockerStateActive,
		Dependency:             deploymentMultiTenantValCDependencySnapshotModel(),
		HAReadiness: DeploymentMultiTenantValCHAReadiness{
			CurrentState:                         DeploymentMultiTenantValCHAReadinessStateActive,
			EvidenceRefs:                         append([]string{}, deploymentMultiTenantValCHAEvidenceRefs()...),
			FreshnessState:                       IntelligenceCalibrationFreshnessFresh,
			HAReadinessEvidenceLinked:            true,
			TopologyEvidence:                     "ha_topology_evidence",
			FailoverTestEvidence:                 "failover_test_evidence",
			DependencyDegradationBehavior:        "dependency_degradation_behavior",
			HealthcheckStateModel:                "healthcheck_state_model",
			QueueWorkerRecoveryBehavior:          "queue_worker_recovery_behavior",
			DegradedModeSemantics:                "degraded_mode_semantics",
			MonitoringAlertRoutingEvidence:       "monitoring_alert_routing_evidence",
			TenantAwareHAImpactBoundary:          "tenant_aware_ha_impact_boundary",
			HAProfileNamingExact:                 true,
			SafeHAReadinessWordingExamplePresent: true,
			DiagnosticOutputComplete:             true,
			ProjectionDisclaimer:                 disclaimer,
		},
		RecoveryReadiness: DeploymentMultiTenantValCRecoveryReadiness{
			CurrentState:                      DeploymentMultiTenantValCRecoveryReadinessStateActive,
			EvidenceRefs:                      append([]string{}, deploymentMultiTenantValCRecoveryEvidenceRefs()...),
			FreshnessState:                    IntelligenceCalibrationFreshnessFresh,
			BackupFreshnessEvidence:           "backup_freshness_evidence",
			BackupEvidenceFreshnessState:      IntelligenceCalibrationFreshnessFresh,
			StaleBackupEvidenceHandlingProven: true,
			RestoreTestEvidence:               "restore_test_evidence",
			TenantScopedRestoreTest:           "tenant_scoped_restore_test",
			RestoreIntegrityHash:              "restore_integrity_hash",
			EncryptedBackupCustodyReference:   "encrypted_backup_custody_reference",
			DRDrillEvidence:                   "dr_drill_evidence",
			RecoveryDependencyInventory:       "recovery_dependency_inventory",
			RestoreTargetBoundary:             "restore_target_boundary",
			BackupRetentionClass:              "backup_retention_class",
			DisposalDeletionBoundary:          "disposal_deletion_boundary",
			RPORTOTarget:                      "rpo_rto_target",
			RecoveryTargetNamingExact:         true,
			RunbookWordingComplete:            true,
			DiagnosticOutputComplete:          true,
			ProjectionDisclaimer:              disclaimer,
		},
		SLAReadiness: DeploymentMultiTenantValCSLAReadiness{
			CurrentState:                          DeploymentMultiTenantValCSLAReadinessStateActive,
			EvidenceRefs:                          append([]string{}, deploymentMultiTenantValCSLAEvidenceRefs()...),
			FreshnessState:                        IntelligenceCalibrationFreshnessFresh,
			SupportabilityEvidence:                "supportability_evidence",
			MonitoringCoverageEvidence:            "monitoring_coverage_evidence",
			AlertRoutingEvidence:                  "alert_routing_evidence",
			IncidentEscalationPath:                "incident_escalation_path",
			SupportScope:                          "tenant:alpha",
			TenantImpactBoundary:                  "tenant_impact_boundary",
			DegradedModeBehavior:                  "degraded_mode_behavior",
			RPORTOTargetReference:                 "rpo_rto_target_reference",
			KnownLimitations:                      "known_limitations",
			NoUptimeGuaranteeWordingPresent:       true,
			SafeSLAReadinessWordingExamplePresent: true,
			DiagnosticOutputComplete:              true,
			ProjectionDisclaimer:                  disclaimer,
		},
		TenantTrustScope: DeploymentMultiTenantValCTenantTrustScope{
			CurrentState:                DeploymentMultiTenantValCTenantTrustScopeStateActive,
			EvidenceRefs:                append([]string{}, deploymentMultiTenantValCTrustScopeEvidenceRefs()...),
			FreshnessState:              IntelligenceCalibrationFreshnessFresh,
			TenantTrustScope:            "tenant_trust_scope_evidence",
			IssuerTrustOwnership:        "issuer_trust_ownership_evidence",
			VerificationBoundary:        "tenant_trust_verification_boundary",
			KeyCustodyReference:         "key_custody_reference",
			KeyCustodyOwner:             "key_custody_owner",
			RotationBehavior:            "key_rotation_behavior",
			RotationState:               DeploymentMultiTenantTrustRotationActive,
			OffboardingTransferBehavior: "offboarding_transfer_behavior",
			RevocationBehavior:          "revocation_behavior",
			TenantTrustExportBoundary:   "tenant_trust_export_boundary",
			TrustScopeNamingExact:       true,
			DiagnosticOutputComplete:    true,
			ProjectionDisclaimer:        disclaimer,
		},
		SiloVisibility: DeploymentMultiTenantValCSiloVisibility{
			CurrentState:                  DeploymentMultiTenantValCSiloVisibilityStateActive,
			EvidenceRefs:                  append([]string{}, deploymentMultiTenantValCSiloVisibilityEvidenceRefs()...),
			FreshnessState:                IntelligenceCalibrationFreshnessFresh,
			EvidenceSiloValidation:        true,
			AuditSiloValidation:           true,
			ExportSiloValidation:          true,
			SupportVisibilityBoundary:     "support_visibility_boundary",
			TenantScopedEvidenceNamespace: "tenant_scoped_evidence_namespace",
			TenantScopedAuditNamespace:    "tenant_scoped_audit_namespace",
			TenantScopedExportBoundary:    "tenant_scoped_export_boundary",
			SupportAccessVisibilityRules:  "support_access_visibility_rules",
			RedactionBoundary:             "redaction_boundary",
			ExportSiloExactIdentity:       "export_silo_exact_identity",
			DiagnosticOutputComplete:      true,
			ProjectionDisclaimer:          disclaimer,
		},
		PrivacyGuard: DeploymentMultiTenantValCPrivacyGuard{
			CurrentState:                           DeploymentMultiTenantValCPrivacyGuardStateActive,
			EvidenceRefs:                           append([]string{}, deploymentMultiTenantValCPrivacyGuardEvidenceRefs()...),
			FreshnessState:                         IntelligenceCalibrationFreshnessFresh,
			CrossTenantPrivacyGuardEvidence:        "cross_tenant_privacy_guard_evidence",
			SideChannelTelemetryPolicy:             "side_channel_telemetry_policy",
			VolumeLeakageCheckPresent:              true,
			ErrorLeakageCheckPresent:               true,
			TimingLeakageCheckPresent:              true,
			AggregationLeakageCheckPresent:         true,
			BoundedAggregationRules:                "bounded_aggregation_rules",
			TenantPrivateMetadataClassification:    "tenant_private_metadata_classification",
			SupportExportPrivacyVisibilityBoundary: "support_export_privacy_visibility_boundary",
			VolumeLeakageNegativeTestPresent:       true,
			ErrorLeakageNegativeTestPresent:        true,
			TimingLeakageNegativeTestPresent:       true,
			AggregationLeakageNegativeTestPresent:  true,
			DiagnosticOutputComplete:               true,
			ProjectionDisclaimer:                   disclaimer,
		},
		NoOverclaim: DeploymentMultiTenantValCNoOverclaimDiscipline{
			CurrentState:         DeploymentMultiTenantValCNoOverclaimStateActive,
			ProjectionDisclaimer: disclaimer,
		},
		ClosureBlockerOverlay: DeploymentMultiTenantValCClosureBlockerOverlay{
			CurrentState:         DeploymentMultiTenantValCClosureBlockerStateActive,
			ProjectionDisclaimer: disclaimer,
		},
	}
}

func ComputeDeploymentMultiTenantValCFoundation(model DeploymentMultiTenantValCFoundation) DeploymentMultiTenantValCFoundation {
	model.DependencyState = EvaluateDeploymentMultiTenantValCDependencyState(model.Dependency)
	model.HAReadinessState = EvaluateDeploymentMultiTenantValCHAReadinessState(model.HAReadiness)
	model.RecoveryReadinessState = EvaluateDeploymentMultiTenantValCRecoveryReadinessState(model.RecoveryReadiness)
	model.SLAReadinessState = EvaluateDeploymentMultiTenantValCSLAReadinessState(model.SLAReadiness)
	model.TenantTrustScopeState = EvaluateDeploymentMultiTenantValCTenantTrustScopeState(model.TenantTrustScope)
	model.SiloVisibilityState = EvaluateDeploymentMultiTenantValCSiloVisibilityState(model.SiloVisibility)
	model.PrivacyGuardState = EvaluateDeploymentMultiTenantValCPrivacyGuardState(model.PrivacyGuard)
	model.NoOverclaimState = EvaluateDeploymentMultiTenantValCNoOverclaimState(model.NoOverclaim)
	model.ClosureBlockerOverlay = DeploymentMultiTenantValCClosureBlockerOverlay{
		ProjectionDisclaimer: deploymentMultiTenantValCProjectionDisclaimer(),
		Findings:             deploymentMultiTenantValCClosureBlockerFindings(model),
	}
	model.ClosureBlockerState = EvaluateDeploymentMultiTenantValCClosureBlockerState(model.ClosureBlockerOverlay)
	model.ClosureBlockerOverlay.CurrentState = model.ClosureBlockerState
	model.Point10State = EvaluateDeploymentMultiTenantPoint10State(model.CurrentState)
	model.CurrentState = EvaluateDeploymentMultiTenantValCState(model)
	model.Point10State = EvaluateDeploymentMultiTenantPoint10State(model.CurrentState)
	model.BlockingReasons = deploymentMultiTenantValCBlockingReasons(model)
	return model
}

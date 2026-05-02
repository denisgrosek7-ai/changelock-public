package operability

import (
	"strings"
	"unicode"
)

const (
	DeploymentMultiTenantValAStateActive  = "deployment_multi_tenant_vala_active"
	DeploymentMultiTenantValAStateBlocked = "deployment_multi_tenant_vala_blocked"

	DeploymentMultiTenantValADependencyStateActive  = "deployment_multi_tenant_vala_dependency_active"
	DeploymentMultiTenantValADependencyStateBlocked = "deployment_multi_tenant_vala_dependency_blocked"

	DeploymentMultiTenantValADeploymentProfileMatrixStateActive  = "deployment_multi_tenant_vala_deployment_profile_matrix_active"
	DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked = "deployment_multi_tenant_vala_deployment_profile_matrix_blocked"

	DeploymentMultiTenantValAPreflightGateStateActive  = "deployment_multi_tenant_vala_preflight_gate_active"
	DeploymentMultiTenantValAPreflightGateStateBlocked = "deployment_multi_tenant_vala_preflight_gate_blocked"

	DeploymentMultiTenantValAIdentityBootstrapStateActive  = "deployment_multi_tenant_vala_identity_bootstrap_active"
	DeploymentMultiTenantValAIdentityBootstrapStateBlocked = "deployment_multi_tenant_vala_identity_bootstrap_blocked"

	DeploymentMultiTenantValAAirGappedEvidenceBundleStateActive  = "deployment_multi_tenant_vala_air_gapped_evidence_bundle_active"
	DeploymentMultiTenantValAAirGappedEvidenceBundleStateBlocked = "deployment_multi_tenant_vala_air_gapped_evidence_bundle_blocked"

	DeploymentMultiTenantValANoOverclaimStateActive  = "deployment_multi_tenant_vala_no_overclaim_active"
	DeploymentMultiTenantValANoOverclaimStateBlocked = "deployment_multi_tenant_vala_no_overclaim_blocked"

	DeploymentMultiTenantValAPassBlockerStateActive  = "deployment_multi_tenant_vala_pass_blocker_active"
	DeploymentMultiTenantValAPassBlockerStateCleanup = "deployment_multi_tenant_vala_pass_blocker_cleanup"
	DeploymentMultiTenantValAPassBlockerStateBlocked = "deployment_multi_tenant_vala_pass_blocker_blocked"

	DeploymentMultiTenantValAProfileSaaS       = "saas"
	DeploymentMultiTenantValAProfileSelfHosted = "self_hosted"
	DeploymentMultiTenantValAProfileAirGapped  = "air_gapped"

	DeploymentMultiTenantValAIdentityModeSSO  = "sso"
	DeploymentMultiTenantValAIdentityModeSAML = "saml"
	DeploymentMultiTenantValAIdentityModeOIDC = "oidc"

	DeploymentMultiTenantValACertificateStateValid   = "valid"
	DeploymentMultiTenantValACertificateStateExpired = "expired"
	DeploymentMultiTenantValACertificateStateUnknown = "unknown"

	DeploymentMultiTenantValASignatureVerificationVerified = "verified"
	DeploymentMultiTenantValASignatureVerificationFailed   = "failed"
	DeploymentMultiTenantValASignatureVerificationUnknown  = "unknown"

	deploymentMultiTenantValAPassBlockerSeverityCLB0 = "CL-B0"
	deploymentMultiTenantValAPassBlockerSeverityCLB1 = "CL-B1"
	deploymentMultiTenantValAPassBlockerSeverityCLB2 = "CL-B2"

	deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile = "deployment_profile"
	deploymentMultiTenantValAPassBlockerSurfacePreflight         = "preflight"
	deploymentMultiTenantValAPassBlockerSurfaceIdentityBootstrap = "identity_bootstrap"
	deploymentMultiTenantValAPassBlockerSurfaceAirGappedBundle   = "air_gapped_bundle"
	deploymentMultiTenantValAPassBlockerSurfaceNoOverclaim       = "no_overclaim"
)

type DeploymentMultiTenantValADependencySnapshot struct {
	Val0CurrentState              string `json:"val0_current_state"`
	Val0DependencyState           string `json:"val0_dependency_state"`
	Val0DeploymentValidationState string `json:"val0_deployment_validation_state"`
	Val0TenantBoundaryState       string `json:"val0_tenant_boundary_state"`
	Val0MSPAuthorityState         string `json:"val0_msp_authority_state"`
	Val0PolicyEnvelopeState       string `json:"val0_policy_envelope_state"`
	Val0TenantTrustScopeState     string `json:"val0_tenant_trust_scope_state"`
	Val0ConnectorContractState    string `json:"val0_connector_contract_state"`
	Val0OperatorActionState       string `json:"val0_operator_action_state"`
	Val0PrivacyGuardState         string `json:"val0_privacy_guard_state"`
	Val0FairShareState            string `json:"val0_fair_share_state"`
	Val0OperationalPreflightState string `json:"val0_operational_preflight_state"`
	Val0FutureContractState       string `json:"val0_future_contract_state"`
	Val0NoOverclaimState          string `json:"val0_no_overclaim_state"`
	Point10State                  string `json:"point_10_state"`
	ProjectionDisclaimer          string `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValADeploymentProfileMatrix struct {
	CurrentState                                  string   `json:"current_state"`
	SupportedProfiles                             []string `json:"supported_profiles,omitempty"`
	SupportedStates                               []string `json:"supported_states,omitempty"`
	EvidenceRefs                                  []string `json:"evidence_refs,omitempty"`
	FreshnessState                                string   `json:"freshness_state"`
	ReadinessEvidenceBacked                       bool     `json:"readiness_evidence_backed"`
	InstallSucceeded                              bool     `json:"install_succeeded"`
	MarketplaceInstallSucceeded                   bool     `json:"marketplace_install_succeeded"`
	InstallSuccessTreatedAsReady                  bool     `json:"install_success_treated_as_ready"`
	MarketplaceInstallTreatedAsReady              bool     `json:"marketplace_install_treated_as_ready"`
	MarketplaceInstallTreatedAsProductionReady    bool     `json:"marketplace_install_treated_as_production_ready"`
	UnsupportedProfileReady                       bool     `json:"unsupported_profile_ready"`
	SaaSState                                     string   `json:"saas_state"`
	SaaSEvidenceKeys                              []string `json:"saas_evidence_keys,omitempty"`
	SaaSTenantConfig                              string   `json:"saas_tenant_config"`
	SaaSRegion                                    string   `json:"saas_region"`
	SaaSIdentityBootstrap                         string   `json:"saas_identity_bootstrap"`
	SaaSConnectorScope                            string   `json:"saas_connector_scope"`
	SaaSBackupPolicy                              string   `json:"saas_backup_policy"`
	SaaSOperatorSupportScope                      string   `json:"saas_operator_support_scope"`
	SelfHostedState                               string   `json:"self_hosted_state"`
	SelfHostedEvidenceKeys                        []string `json:"self_hosted_evidence_keys,omitempty"`
	SelfHostedArtifactProvenance                  string   `json:"self_hosted_artifact_provenance"`
	SelfHostedEnvironmentManifest                 string   `json:"self_hosted_environment_manifest"`
	SelfHostedConfigValidation                    string   `json:"self_hosted_config_validation"`
	SelfHostedIAMKMSDependency                    string   `json:"self_hosted_iam_kms_dependency"`
	SelfHostedBackupTarget                        string   `json:"self_hosted_backup_target"`
	SelfHostedUpgradeRollbackPlan                 string   `json:"self_hosted_upgrade_rollback_plan"`
	SelfHostedUnsupportedSemanticsExplicit        bool     `json:"self_hosted_unsupported_semantics_explicit"`
	SelfHostedDegradedSemanticsExplicit           bool     `json:"self_hosted_degraded_semantics_explicit"`
	AirGappedState                                string   `json:"air_gapped_state"`
	AirGappedEvidenceKeys                         []string `json:"air_gapped_evidence_keys,omitempty"`
	AirGappedOfflineArtifactBundle                string   `json:"air_gapped_offline_artifact_bundle"`
	AirGappedOfflineEvidenceBundle                string   `json:"air_gapped_offline_evidence_bundle"`
	AirGappedSignatureHashVerificationState       string   `json:"air_gapped_signature_hash_verification_state"`
	AirGappedUnsupportedOnlineDependencies        []string `json:"air_gapped_unsupported_online_dependencies,omitempty"`
	AirGappedUnsupportedDependenciesHidden        bool     `json:"air_gapped_unsupported_dependencies_hidden"`
	AirGappedUnsupportedDependenciesSilentlyReady bool     `json:"air_gapped_unsupported_dependencies_silently_ready"`
	AirGappedOfflineReplayExportPath              string   `json:"air_gapped_offline_replay_export_path"`
	AirGappedUnsupportedSemanticsExplicit         bool     `json:"air_gapped_unsupported_semantics_explicit"`
	AirGappedDegradedSemanticsExplicit            bool     `json:"air_gapped_degraded_semantics_explicit"`
	ProfileNamingExact                            bool     `json:"profile_naming_exact"`
	SafeReadinessWordingExamplePresent            bool     `json:"safe_readiness_wording_example_present"`
	DiagnosticOutputComplete                      bool     `json:"diagnostic_output_complete"`
	ObservedClaims                                []string `json:"observed_claims,omitempty"`
	ProjectionDisclaimer                          string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValAPreflightGate struct {
	CurrentState                          string   `json:"current_state"`
	EvidenceRefs                          []string `json:"evidence_refs,omitempty"`
	CheckKeys                             []string `json:"check_keys,omitempty"`
	FreshnessState                        string   `json:"freshness_state"`
	TenantScope                           string   `json:"tenant_scope"`
	InstallConfigValidation               bool     `json:"install_config_validation"`
	UpgradeConfigDiff                     bool     `json:"upgrade_config_diff"`
	DBSchemaMigrationDryRun               bool     `json:"db_schema_migration_dry_run"`
	BackupBeforeUpgradeEvidence           bool     `json:"backup_before_upgrade_evidence"`
	RollbackPlanEvidence                  bool     `json:"rollback_plan_evidence"`
	PolicyMigrationCompatibility          bool     `json:"policy_migration_compatibility"`
	ConnectorPermissionChangeReview       bool     `json:"connector_permission_change_review"`
	KeyRotationReadiness                  bool     `json:"key_rotation_readiness"`
	TenantBoundaryValidation              bool     `json:"tenant_boundary_validation"`
	ProductionImpactSafeByDefault         bool     `json:"production_impact_safe_by_default"`
	InstallSuccessUsedAsPreflightEvidence bool     `json:"install_success_used_as_preflight_evidence"`
	DiagnosticOutputComplete              bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer                  string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValAIdentityBootstrap struct {
	CurrentState                                string   `json:"current_state"`
	SupportedModes                              []string `json:"supported_modes,omitempty"`
	EvidenceRefs                                []string `json:"evidence_refs,omitempty"`
	ValidationKeys                              []string `json:"validation_keys,omitempty"`
	FreshnessState                              string   `json:"freshness_state"`
	IdentityMode                                string   `json:"identity_mode"`
	TenantScope                                 string   `json:"tenant_scope"`
	IssuerEntityID                              string   `json:"issuer_entity_id"`
	CallbackRedirectURL                         string   `json:"callback_redirect_url"`
	CertificateExpiryState                      string   `json:"certificate_expiry_state"`
	GroupRoleMapping                            string   `json:"group_role_mapping"`
	AdminBootstrapFallback                      string   `json:"admin_bootstrap_fallback"`
	BreakGlassCompatibility                     bool     `json:"break_glass_compatibility"`
	BreakGlassExpiryPresent                     bool     `json:"break_glass_expiry_present"`
	BreakGlassRevocationPresent                 bool     `json:"break_glass_revocation_present"`
	DisabledUnsafeFallbackHandling              bool     `json:"disabled_unsafe_fallback_handling"`
	TenantSpecificIdentityBoundary              string   `json:"tenant_specific_identity_boundary"`
	UnsafeFallbackEnabled                       bool     `json:"unsafe_fallback_enabled"`
	SSOConfiguredMeansSecureClaim               bool     `json:"sso_configured_means_secure_claim"`
	IdentityReadinessImpliesDeploymentReadiness bool     `json:"identity_readiness_implies_deployment_readiness"`
	DiagnosticOutputComplete                    bool     `json:"diagnostic_output_complete"`
	ObservedClaims                              []string `json:"observed_claims,omitempty"`
	ProjectionDisclaimer                        string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValAAirGappedEvidenceBundle struct {
	CurrentState                            string   `json:"current_state"`
	EvidenceRefs                            []string `json:"evidence_refs,omitempty"`
	BundleKeys                              []string `json:"bundle_keys,omitempty"`
	FreshnessState                          string   `json:"freshness_state"`
	AirGappedModeExplicit                   bool     `json:"air_gapped_mode_explicit"`
	OfflineEvidenceBundleEvidenceBacked     bool     `json:"offline_evidence_bundle_evidence_backed"`
	BundleManifest                          string   `json:"bundle_manifest"`
	ArtifactHashes                          string   `json:"artifact_hashes"`
	ProofPackHashes                         string   `json:"proof_pack_hashes"`
	Signer                                  string   `json:"signer"`
	PolicyVersion                           string   `json:"policy_version"`
	EngineVersion                           string   `json:"engine_version"`
	Timestamp                               string   `json:"timestamp"`
	UnsupportedOnlineDependencies           []string `json:"unsupported_online_dependencies,omitempty"`
	UnsupportedOnlineDependenciesHidden     bool     `json:"unsupported_online_dependencies_hidden"`
	ReplayInstructions                      string   `json:"replay_instructions"`
	OfflineReplayExportPath                 string   `json:"offline_replay_export_path"`
	SignatureHashVerificationState          string   `json:"signature_hash_verification_state"`
	AirGappedCertifiedClaim                 bool     `json:"air_gapped_certified_claim"`
	AirGappedMeansFullyOfflineVerifiedClaim bool     `json:"air_gapped_means_fully_offline_verified_claim"`
	DiagnosticOutputComplete                bool     `json:"diagnostic_output_complete"`
	ObservedClaims                          []string `json:"observed_claims,omitempty"`
	ProjectionDisclaimer                    string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValANoOverclaimDiscipline struct {
	CurrentState         string   `json:"current_state"`
	ObservedClaims       []string `json:"observed_claims,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValAPassBlockerFinding struct {
	Severity              string `json:"severity"`
	Surface               string `json:"surface"`
	Reason                string `json:"reason"`
	BlocksCurrentValAPass bool   `json:"blocks_current_vala_pass"`
}

type DeploymentMultiTenantValAPassBlockerOverlay struct {
	CurrentState         string                                        `json:"current_state"`
	Findings             []DeploymentMultiTenantValAPassBlockerFinding `json:"findings,omitempty"`
	ProjectionDisclaimer string                                        `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValAFoundation struct {
	CurrentState                 string                                           `json:"current_state"`
	Point10State                 string                                           `json:"point_10_state"`
	ProjectionDisclaimer         string                                           `json:"projection_disclaimer"`
	BlockingReasons              []string                                         `json:"blocking_reasons,omitempty"`
	DependencyState              string                                           `json:"dependency_state"`
	DeploymentProfileMatrixState string                                           `json:"deployment_profile_matrix_state"`
	PreflightGateState           string                                           `json:"preflight_gate_state"`
	IdentityBootstrapState       string                                           `json:"identity_bootstrap_state"`
	AirGappedEvidenceBundleState string                                           `json:"air_gapped_evidence_bundle_state"`
	NoOverclaimState             string                                           `json:"no_overclaim_state"`
	PassBlockerState             string                                           `json:"pass_blocker_state"`
	Dependency                   DeploymentMultiTenantValADependencySnapshot      `json:"dependency"`
	DeploymentProfileMatrix      DeploymentMultiTenantValADeploymentProfileMatrix `json:"deployment_profile_matrix"`
	PreflightGate                DeploymentMultiTenantValAPreflightGate           `json:"preflight_gate"`
	IdentityBootstrap            DeploymentMultiTenantValAIdentityBootstrap       `json:"identity_bootstrap"`
	AirGappedEvidenceBundle      DeploymentMultiTenantValAAirGappedEvidenceBundle `json:"air_gapped_evidence_bundle"`
	NoOverclaim                  DeploymentMultiTenantValANoOverclaimDiscipline   `json:"no_overclaim"`
	PassBlockerOverlay           DeploymentMultiTenantValAPassBlockerOverlay      `json:"pass_blocker_overlay"`
}

func deploymentMultiTenantValAProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_marketplace_deployment_profile deployment_multi_tenant_vala"
}

func deploymentMultiTenantValAHasProjectionDisclaimer(value string) bool {
	normalized := strings.ToLower(strings.TrimSpace(value))
	return strings.Contains(normalized, "projection_only") &&
		strings.Contains(normalized, "not_canonical_truth") &&
		strings.Contains(normalized, "deployment_multi_tenant_vala")
}

func deploymentMultiTenantValADeploymentProfileMatrixEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vala-profile-matrix-001"}
}

func deploymentMultiTenantValAPreflightEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vala-preflight-001"}
}

func deploymentMultiTenantValAIdentityBootstrapEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vala-identity-bootstrap-001"}
}

func deploymentMultiTenantValAAirGappedEvidenceBundleEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vala-air-gapped-evidence-bundle-001"}
}

func deploymentMultiTenantValARequiredProfiles() []string {
	return []string{
		DeploymentMultiTenantValAProfileSaaS,
		DeploymentMultiTenantValAProfileSelfHosted,
		DeploymentMultiTenantValAProfileAirGapped,
	}
}

func deploymentMultiTenantValARequiredStates() []string {
	return []string{
		DeploymentMultiTenantDeploymentStateReady,
		DeploymentMultiTenantDeploymentStateDegraded,
		DeploymentMultiTenantDeploymentStateIncomplete,
		DeploymentMultiTenantDeploymentStateUnsupported,
		DeploymentMultiTenantDeploymentStateBlocked,
		DeploymentMultiTenantDeploymentStateUnknown,
	}
}

func deploymentMultiTenantValAIdentityModes() []string {
	return []string{
		DeploymentMultiTenantValAIdentityModeSSO,
		DeploymentMultiTenantValAIdentityModeSAML,
		DeploymentMultiTenantValAIdentityModeOIDC,
	}
}

func deploymentMultiTenantValAProfileMatrixSaaSKeys() []string {
	return []string{
		"tenant_config",
		"region",
		"identity_bootstrap",
		"connector_scope",
		"backup_policy",
		"operator_support_scope",
	}
}

func deploymentMultiTenantValAProfileMatrixSelfHostedKeys() []string {
	return []string{
		"artifact_provenance",
		"environment_manifest",
		"config_validation",
		"iam_kms_dependencies",
		"backup_target",
		"upgrade_rollback_plan",
	}
}

func deploymentMultiTenantValAProfileMatrixAirGappedKeys() []string {
	return []string{
		"offline_artifact_bundle",
		"offline_evidence_bundle",
		"signature_hash_verification",
		"unsupported_dependency_list",
		"offline_replay_export_path",
	}
}

func deploymentMultiTenantValAPreflightCheckKeys() []string {
	return []string{
		"install_config_validation",
		"upgrade_config_diff",
		"db_schema_migration_dry_run",
		"backup_before_upgrade_evidence",
		"rollback_plan_evidence",
		"policy_migration_compatibility",
		"connector_permission_changes_review",
		"key_rotation_readiness",
		"tenant_boundary_validation",
	}
}

func deploymentMultiTenantValAIdentityValidationKeys() []string {
	return []string{
		"issuer_entity_id",
		"callback_redirect_url",
		"certificate_expiry_state",
		"group_role_mapping",
		"admin_bootstrap_fallback",
		"break_glass_compatibility",
		"disabled_unsafe_fallback_handling",
		"tenant_specific_identity_boundary",
	}
}

func deploymentMultiTenantValAAirGappedBundleKeys() []string {
	return []string{
		"bundle_manifest",
		"artifact_hashes",
		"proof_pack_hashes",
		"signer",
		"policy_version",
		"engine_version",
		"timestamp",
		"unsupported_online_dependencies",
		"replay_instructions",
		"offline_replay_export_path",
		"signature_hash_verification",
	}
}

func deploymentMultiTenantValAAllValuesValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	for _, value := range values {
		if !deploymentMultiTenantVal0ExactValueIsValid(value) {
			return false
		}
	}
	return true
}

func deploymentMultiTenantValAHasNoUnsupportedOnlineDependencies(values []string) bool {
	return len(values) == 1 && strings.TrimSpace(values[0]) == "none_explicit"
}

func deploymentMultiTenantValAAirGappedUnsupportedDependencyIDIsValid(value string) bool {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" || trimmed != strings.ToLower(trimmed) || trimmed == "none_explicit" {
		return false
	}
	normalized := deploymentMultiTenantVal0NormalizeClaimText(trimmed)
	if normalized == "" {
		return false
	}
	switch normalized {
	case "unknown", "partial", "incomplete", "stale", "malformed", "blocked":
		return false
	}
	for _, char := range trimmed {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			continue
		}
		switch char {
		case '_', '-', ':', '/', '.':
		default:
			return false
		}
	}
	return true
}

func deploymentMultiTenantValAAirGappedUnsupportedDependencyListIsValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	if deploymentMultiTenantValAHasNoUnsupportedOnlineDependencies(values) {
		return true
	}
	for _, value := range values {
		if strings.TrimSpace(value) == "none_explicit" {
			return false
		}
		if !deploymentMultiTenantValAAirGappedUnsupportedDependencyIDIsValid(value) {
			return false
		}
	}
	return true
}

func deploymentMultiTenantValAHasExplicitUnsupportedDependencies(values []string) bool {
	return deploymentMultiTenantValAAirGappedUnsupportedDependencyListIsValid(values) &&
		!deploymentMultiTenantValAHasNoUnsupportedOnlineDependencies(values)
}

func deploymentMultiTenantValADependencySnapshotModel() DeploymentMultiTenantValADependencySnapshot {
	val0 := ComputeDeploymentMultiTenantVal0Foundation(DeploymentMultiTenantVal0FoundationModel())
	return DeploymentMultiTenantValADependencySnapshot{
		Val0CurrentState:              val0.CurrentState,
		Val0DependencyState:           val0.DependencyState,
		Val0DeploymentValidationState: val0.DeploymentValidationState,
		Val0TenantBoundaryState:       val0.TenantBoundaryState,
		Val0MSPAuthorityState:         val0.MSPAuthorityState,
		Val0PolicyEnvelopeState:       val0.PolicyEnvelopeState,
		Val0TenantTrustScopeState:     val0.TenantTrustScopeState,
		Val0ConnectorContractState:    val0.ConnectorContractState,
		Val0OperatorActionState:       val0.OperatorActionState,
		Val0PrivacyGuardState:         val0.PrivacyGuardState,
		Val0FairShareState:            val0.FairShareState,
		Val0OperationalPreflightState: val0.OperationalPreflightState,
		Val0FutureContractState:       val0.FutureContractState,
		Val0NoOverclaimState:          val0.NoOverclaimState,
		Point10State:                  val0.Point10State,
		ProjectionDisclaimer:          val0.ProjectionDisclaimer,
	}
}

func EvaluateDeploymentMultiTenantValADependencyState(model DeploymentMultiTenantValADependencySnapshot) string {
	if !deploymentMultiTenantVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeploymentMultiTenantValADependencyStateBlocked
	}
	if strings.TrimSpace(model.Val0CurrentState) != DeploymentMultiTenantVal0StateActive ||
		strings.TrimSpace(model.Val0DependencyState) != DeploymentMultiTenantVal0DependencyStateActive ||
		strings.TrimSpace(model.Val0DeploymentValidationState) != DeploymentMultiTenantVal0DeploymentValidationStateActive ||
		strings.TrimSpace(model.Val0TenantBoundaryState) != DeploymentMultiTenantVal0TenantBoundaryStateActive ||
		strings.TrimSpace(model.Val0MSPAuthorityState) != DeploymentMultiTenantVal0MSPAuthorityStateActive ||
		strings.TrimSpace(model.Val0PolicyEnvelopeState) != DeploymentMultiTenantVal0PolicyEnvelopeStateActive ||
		strings.TrimSpace(model.Val0TenantTrustScopeState) != DeploymentMultiTenantVal0TenantTrustScopeStateActive ||
		strings.TrimSpace(model.Val0ConnectorContractState) != DeploymentMultiTenantVal0ConnectorContractStateActive ||
		strings.TrimSpace(model.Val0OperatorActionState) != DeploymentMultiTenantVal0OperatorActionStateActive ||
		strings.TrimSpace(model.Val0PrivacyGuardState) != DeploymentMultiTenantVal0PrivacyGuardStateActive ||
		strings.TrimSpace(model.Val0FairShareState) != DeploymentMultiTenantVal0FairShareStateActive ||
		strings.TrimSpace(model.Val0OperationalPreflightState) != DeploymentMultiTenantVal0OperationalPreflightStateActive ||
		strings.TrimSpace(model.Val0FutureContractState) != DeploymentMultiTenantVal0FutureContractStateActive ||
		strings.TrimSpace(model.Val0NoOverclaimState) != DeploymentMultiTenantVal0NoOverclaimStateActive ||
		strings.TrimSpace(model.Point10State) != DeploymentMultiTenantPoint10StateNotComplete {
		return DeploymentMultiTenantValADependencyStateBlocked
	}
	return DeploymentMultiTenantValADependencyStateActive
}

func EvaluateDeploymentMultiTenantValADeploymentProfileMatrixState(model DeploymentMultiTenantValADeploymentProfileMatrix) string {
	hasNoUnsupportedDependencies := deploymentMultiTenantValAHasNoUnsupportedOnlineDependencies(model.AirGappedUnsupportedOnlineDependencies)
	hasExplicitUnsupportedDependencies := deploymentMultiTenantValAHasExplicitUnsupportedDependencies(model.AirGappedUnsupportedOnlineDependencies)
	if !deploymentMultiTenantValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, deploymentMultiTenantValADeploymentProfileMatrixEvidenceRefs()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!containsExactTrimmedStringSet(model.SupportedProfiles, deploymentMultiTenantValARequiredProfiles()...) ||
		!containsExactTrimmedStringSet(model.SupportedStates, deploymentMultiTenantValARequiredStates()...) ||
		!model.ReadinessEvidenceBacked {
		return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
	}
	if model.InstallSuccessTreatedAsReady || model.MarketplaceInstallTreatedAsReady ||
		model.MarketplaceInstallTreatedAsProductionReady || model.UnsupportedProfileReady {
		return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
	}
	if strings.TrimSpace(model.SaaSState) != DeploymentMultiTenantDeploymentStateReady ||
		!containsExactTrimmedStringSet(model.SaaSEvidenceKeys, deploymentMultiTenantValAProfileMatrixSaaSKeys()...) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SaaSTenantConfig) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SaaSRegion) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SaaSIdentityBootstrap) ||
		!deploymentMultiTenantVal0TenantScopedValueIsValid(model.SaaSConnectorScope) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SaaSBackupPolicy) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.SaaSOperatorSupportScope) {
		return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
	}
	if strings.TrimSpace(model.SelfHostedState) != DeploymentMultiTenantDeploymentStateReady ||
		!containsExactTrimmedStringSet(model.SelfHostedEvidenceKeys, deploymentMultiTenantValAProfileMatrixSelfHostedKeys()...) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SelfHostedArtifactProvenance) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SelfHostedEnvironmentManifest) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SelfHostedConfigValidation) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SelfHostedIAMKMSDependency) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SelfHostedBackupTarget) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.SelfHostedUpgradeRollbackPlan) {
		return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
	}
	if !containsTrimmedString(deploymentMultiTenantValARequiredStates(), model.AirGappedState) ||
		!containsExactTrimmedStringSet(model.AirGappedEvidenceKeys, deploymentMultiTenantValAProfileMatrixAirGappedKeys()...) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.AirGappedOfflineArtifactBundle) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.AirGappedOfflineEvidenceBundle) ||
		strings.TrimSpace(model.AirGappedSignatureHashVerificationState) != DeploymentMultiTenantValASignatureVerificationVerified ||
		model.AirGappedUnsupportedDependenciesHidden ||
		model.AirGappedUnsupportedDependenciesSilentlyReady ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.AirGappedOfflineReplayExportPath) {
		return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
	}
	if !deploymentMultiTenantValAAirGappedUnsupportedDependencyListIsValid(model.AirGappedUnsupportedOnlineDependencies) {
		return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
	}
	if hasExplicitUnsupportedDependencies {
		if strings.TrimSpace(model.AirGappedState) != DeploymentMultiTenantDeploymentStateDegraded &&
			strings.TrimSpace(model.AirGappedState) != DeploymentMultiTenantDeploymentStateUnsupported {
			return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
		}
		if !model.AirGappedUnsupportedSemanticsExplicit || !model.AirGappedDegradedSemanticsExplicit {
			return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
		}
	} else if !hasNoUnsupportedDependencies || strings.TrimSpace(model.AirGappedState) != DeploymentMultiTenantDeploymentStateReady {
		return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
	}
	if model.InstallSucceeded && !model.ReadinessEvidenceBacked {
		return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
	}
	if model.MarketplaceInstallSucceeded && (!model.ReadinessEvidenceBacked || model.MarketplaceInstallTreatedAsReady || model.MarketplaceInstallTreatedAsProductionReady) {
		return DeploymentMultiTenantValADeploymentProfileMatrixStateBlocked
	}
	return DeploymentMultiTenantValADeploymentProfileMatrixStateActive
}

func EvaluateDeploymentMultiTenantValAPreflightGateState(model DeploymentMultiTenantValAPreflightGate) string {
	if !deploymentMultiTenantValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, deploymentMultiTenantValAPreflightEvidenceRefs()...) ||
		!containsExactTrimmedStringSet(model.CheckKeys, deploymentMultiTenantValAPreflightCheckKeys()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!deploymentMultiTenantVal0TenantScopedValueIsValid(model.TenantScope) ||
		!model.InstallConfigValidation || !model.UpgradeConfigDiff || !model.DBSchemaMigrationDryRun ||
		!model.BackupBeforeUpgradeEvidence || !model.RollbackPlanEvidence ||
		!model.PolicyMigrationCompatibility || !model.ConnectorPermissionChangeReview ||
		!model.KeyRotationReadiness || !model.TenantBoundaryValidation ||
		model.ProductionImpactSafeByDefault || model.InstallSuccessUsedAsPreflightEvidence {
		return DeploymentMultiTenantValAPreflightGateStateBlocked
	}
	return DeploymentMultiTenantValAPreflightGateStateActive
}

func EvaluateDeploymentMultiTenantValAIdentityBootstrapState(model DeploymentMultiTenantValAIdentityBootstrap) string {
	if !deploymentMultiTenantValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.SupportedModes, deploymentMultiTenantValAIdentityModes()...) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, deploymentMultiTenantValAIdentityBootstrapEvidenceRefs()...) ||
		!containsExactTrimmedStringSet(model.ValidationKeys, deploymentMultiTenantValAIdentityValidationKeys()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!containsTrimmedString(deploymentMultiTenantValAIdentityModes(), model.IdentityMode) ||
		!deploymentMultiTenantVal0TenantScopedValueIsValid(model.TenantScope) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.IssuerEntityID) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.CallbackRedirectURL) ||
		strings.TrimSpace(model.CertificateExpiryState) != DeploymentMultiTenantValACertificateStateValid ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.GroupRoleMapping) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.AdminBootstrapFallback) ||
		!model.BreakGlassCompatibility || !model.BreakGlassExpiryPresent ||
		!model.BreakGlassRevocationPresent || !model.DisabledUnsafeFallbackHandling ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.TenantSpecificIdentityBoundary) ||
		model.UnsafeFallbackEnabled || model.SSOConfiguredMeansSecureClaim ||
		model.IdentityReadinessImpliesDeploymentReadiness {
		return DeploymentMultiTenantValAIdentityBootstrapStateBlocked
	}
	return DeploymentMultiTenantValAIdentityBootstrapStateActive
}

func EvaluateDeploymentMultiTenantValAAirGappedEvidenceBundleState(model DeploymentMultiTenantValAAirGappedEvidenceBundle) string {
	if !deploymentMultiTenantValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, deploymentMultiTenantValAAirGappedEvidenceBundleEvidenceRefs()...) ||
		!containsExactTrimmedStringSet(model.BundleKeys, deploymentMultiTenantValAAirGappedBundleKeys()...) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!model.AirGappedModeExplicit || !model.OfflineEvidenceBundleEvidenceBacked ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.BundleManifest) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.ArtifactHashes) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.ProofPackHashes) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.Signer) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.PolicyVersion) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.EngineVersion) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.Timestamp) ||
		!deploymentMultiTenantValAAllValuesValid(model.UnsupportedOnlineDependencies) ||
		model.UnsupportedOnlineDependenciesHidden ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.ReplayInstructions) ||
		!deploymentMultiTenantVal0ExactValueIsValid(model.OfflineReplayExportPath) ||
		strings.TrimSpace(model.SignatureHashVerificationState) != DeploymentMultiTenantValASignatureVerificationVerified ||
		model.AirGappedCertifiedClaim || model.AirGappedMeansFullyOfflineVerifiedClaim {
		return DeploymentMultiTenantValAAirGappedEvidenceBundleStateBlocked
	}
	return DeploymentMultiTenantValAAirGappedEvidenceBundleStateActive
}

func deploymentMultiTenantValAContainsForbiddenClaim(values ...string) bool {
	allowed := []string{
		"validated deployment profile",
		"evidence-linked readiness state",
		"bounded marketplace deployment profile",
		"self-hosted readiness evidence",
		"air-gapped offline evidence bundle",
		"offline replay/export path",
		"sso bootstrap validation",
		"saml/oidc validation evidence",
		"tenant-scoped preflight",
		"rollback plan evidence",
		"backup-before-upgrade evidence",
		"unsupported dependency explicitly listed",
		"degraded deployment state",
		"incomplete deployment state",
		"not production approval",
		"not deployment approval",
		"not compliance certification",
		"not canonical truth",
	}
	disallowed := []string{
		"production approved",
		"deployment approved",
		"marketplace certified",
		"marketplace production ready",
		"one-click secure",
		"install success means ready",
		"marketplace install means ready",
		"marketplace install means production ready",
		"customer ready without validation",
		"compliant by default",
		"compliance guaranteed",
		"regulator-approved",
		"self-hosted production approved",
		"air-gapped certified",
		"air-gapped means fully offline verified",
		"sso secure by default",
		"sso configured means secure",
		"rbac complete by default",
		"deployment readiness guaranteed",
		"unsupported profile ready",
		"offline bundle certified",
		"offline replay guarantees correctness",
		"preflight passed means production approved",
		"rollback guaranteed",
		"zero-risk deployment",
		"guaranteed uptime",
		"sla guaranteed",
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
		if _, ok := allowedExact[compact]; ok {
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

func EvaluateDeploymentMultiTenantValANoOverclaimState(model DeploymentMultiTenantValANoOverclaimDiscipline) string {
	if !deploymentMultiTenantValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		deploymentMultiTenantValAContainsForbiddenClaim(model.ObservedClaims...) {
		return DeploymentMultiTenantValANoOverclaimStateBlocked
	}
	return DeploymentMultiTenantValANoOverclaimStateActive
}

func deploymentMultiTenantValAPassBlockerFinding(severity, surface, reason string) DeploymentMultiTenantValAPassBlockerFinding {
	return DeploymentMultiTenantValAPassBlockerFinding{
		Severity:              severity,
		Surface:               surface,
		Reason:                reason,
		BlocksCurrentValAPass: true,
	}
}

func deploymentMultiTenantValAPassBlockerFindings(model DeploymentMultiTenantValAFoundation) []DeploymentMultiTenantValAPassBlockerFinding {
	findings := []DeploymentMultiTenantValAPassBlockerFinding{}
	if model.DeploymentProfileMatrix.InstallSuccessTreatedAsReady {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB0, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "install success treated as readiness"))
	}
	if model.DeploymentProfileMatrix.MarketplaceInstallTreatedAsReady {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB0, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "marketplace install treated as readiness"))
	}
	if model.DeploymentProfileMatrix.MarketplaceInstallTreatedAsProductionReady {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB0, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "marketplace install treated as production readiness"))
	}
	if model.IdentityBootstrap.SSOConfiguredMeansSecureClaim {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB0, deploymentMultiTenantValAPassBlockerSurfaceIdentityBootstrap, "sso configured treated as secure"))
	}
	if model.IdentityBootstrap.IdentityReadinessImpliesDeploymentReadiness {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB0, deploymentMultiTenantValAPassBlockerSurfaceIdentityBootstrap, "sso readiness treated as deployment readiness"))
	}
	if model.IdentityBootstrap.UnsafeFallbackEnabled {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB0, deploymentMultiTenantValAPassBlockerSurfaceIdentityBootstrap, "unsafe fallback enabled"))
	}
	if deploymentMultiTenantValAContainsForbiddenClaim(model.DeploymentProfileMatrix.ObservedClaims...) {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB0, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "marketplace or msp overclaim in deployment profile wording"))
	}
	if !model.DeploymentProfileMatrix.SelfHostedUnsupportedSemanticsExplicit || !model.DeploymentProfileMatrix.SelfHostedDegradedSemanticsExplicit {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB1, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "self-hosted profile lacks unsupported or degraded semantics"))
	}
	if !model.DeploymentProfileMatrix.AirGappedUnsupportedSemanticsExplicit || !model.DeploymentProfileMatrix.AirGappedDegradedSemanticsExplicit {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB1, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "air-gapped profile lacks unsupported or degraded semantics"))
	}
	if model.DeploymentProfileMatrix.AirGappedUnsupportedDependenciesHidden || model.DeploymentProfileMatrix.AirGappedUnsupportedDependenciesSilentlyReady {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB1, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "unsupported air-gapped dependency hidden or silently treated as ready"))
	}
	if deploymentMultiTenantValAHasExplicitUnsupportedDependencies(model.DeploymentProfileMatrix.AirGappedUnsupportedOnlineDependencies) &&
		strings.TrimSpace(model.DeploymentProfileMatrix.AirGappedState) == DeploymentMultiTenantDeploymentStateReady {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB1, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "explicit unsupported air-gapped dependency treated as ready"))
	}
	if !model.DeploymentProfileMatrix.ProfileNamingExact {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB2, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "ambiguous deployment profile naming"))
	}
	if !model.DeploymentProfileMatrix.SafeReadinessWordingExamplePresent {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB2, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "missing safe wording example for deployment readiness"))
	}
	if !model.DeploymentProfileMatrix.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB2, deploymentMultiTenantValAPassBlockerSurfaceDeploymentProfile, "incomplete diagnostic output for readiness blockers"))
	}
	if !model.PreflightGate.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB2, deploymentMultiTenantValAPassBlockerSurfacePreflight, "incomplete diagnostic output for preflight blockers"))
	}
	if !model.IdentityBootstrap.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB2, deploymentMultiTenantValAPassBlockerSurfaceIdentityBootstrap, "incomplete diagnostic output for identity blockers"))
	}
	if !model.AirGappedEvidenceBundle.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB2, deploymentMultiTenantValAPassBlockerSurfaceAirGappedBundle, "incomplete diagnostic output for air-gapped blockers"))
	}
	if model.NoOverclaimState != DeploymentMultiTenantValANoOverclaimStateActive {
		findings = append(findings, deploymentMultiTenantValAPassBlockerFinding(deploymentMultiTenantValAPassBlockerSeverityCLB0, deploymentMultiTenantValAPassBlockerSurfaceNoOverclaim, "forbidden deployment or marketplace claim present"))
	}
	return findings
}

func EvaluateDeploymentMultiTenantValAPassBlockerState(model DeploymentMultiTenantValAPassBlockerOverlay) string {
	if !deploymentMultiTenantValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeploymentMultiTenantValAPassBlockerStateBlocked
	}
	hasCleanup := false
	for _, finding := range model.Findings {
		switch strings.TrimSpace(finding.Severity) {
		case deploymentMultiTenantValAPassBlockerSeverityCLB0, deploymentMultiTenantValAPassBlockerSeverityCLB1:
			return DeploymentMultiTenantValAPassBlockerStateBlocked
		case deploymentMultiTenantValAPassBlockerSeverityCLB2:
			hasCleanup = true
		default:
			return DeploymentMultiTenantValAPassBlockerStateBlocked
		}
	}
	if hasCleanup {
		return DeploymentMultiTenantValAPassBlockerStateCleanup
	}
	return DeploymentMultiTenantValAPassBlockerStateActive
}

func EvaluateDeploymentMultiTenantValAState(model DeploymentMultiTenantValAFoundation) string {
	if !deploymentMultiTenantValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		strings.TrimSpace(model.DependencyState) != DeploymentMultiTenantValADependencyStateActive ||
		strings.TrimSpace(model.DeploymentProfileMatrixState) != DeploymentMultiTenantValADeploymentProfileMatrixStateActive ||
		strings.TrimSpace(model.PreflightGateState) != DeploymentMultiTenantValAPreflightGateStateActive ||
		strings.TrimSpace(model.IdentityBootstrapState) != DeploymentMultiTenantValAIdentityBootstrapStateActive ||
		strings.TrimSpace(model.AirGappedEvidenceBundleState) != DeploymentMultiTenantValAAirGappedEvidenceBundleStateActive ||
		strings.TrimSpace(model.NoOverclaimState) != DeploymentMultiTenantValANoOverclaimStateActive ||
		strings.TrimSpace(model.PassBlockerState) != DeploymentMultiTenantValAPassBlockerStateActive ||
		strings.TrimSpace(model.Point10State) != DeploymentMultiTenantPoint10StateNotComplete {
		return DeploymentMultiTenantValAStateBlocked
	}
	return DeploymentMultiTenantValAStateActive
}

func deploymentMultiTenantValABlockingReasons(model DeploymentMultiTenantValAFoundation) []string {
	reasons := []string{}
	if !deploymentMultiTenantValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "aggregate_projection_disclaimer_blocked")
	}
	if model.DependencyState != DeploymentMultiTenantValADependencyStateActive {
		reasons = append(reasons, "dependency")
	}
	if model.DeploymentProfileMatrixState != DeploymentMultiTenantValADeploymentProfileMatrixStateActive {
		reasons = append(reasons, "deployment_profile_matrix")
	}
	if model.PreflightGateState != DeploymentMultiTenantValAPreflightGateStateActive {
		reasons = append(reasons, "preflight_gate")
	}
	if model.IdentityBootstrapState != DeploymentMultiTenantValAIdentityBootstrapStateActive {
		reasons = append(reasons, "identity_bootstrap")
	}
	if model.AirGappedEvidenceBundleState != DeploymentMultiTenantValAAirGappedEvidenceBundleStateActive {
		reasons = append(reasons, "air_gapped_evidence_bundle")
	}
	if model.NoOverclaimState != DeploymentMultiTenantValANoOverclaimStateActive {
		reasons = append(reasons, "no_overclaim")
	}
	if model.PassBlockerState != DeploymentMultiTenantValAPassBlockerStateActive {
		reasons = append(reasons, "pass_blocker")
	}
	return reasons
}

func DeploymentMultiTenantValAFoundationModel() DeploymentMultiTenantValAFoundation {
	disclaimer := deploymentMultiTenantValAProjectionDisclaimer()
	return DeploymentMultiTenantValAFoundation{
		CurrentState:                 DeploymentMultiTenantValAStateActive,
		Point10State:                 DeploymentMultiTenantPoint10StateNotComplete,
		ProjectionDisclaimer:         disclaimer,
		DependencyState:              DeploymentMultiTenantValADependencyStateActive,
		DeploymentProfileMatrixState: DeploymentMultiTenantValADeploymentProfileMatrixStateActive,
		PreflightGateState:           DeploymentMultiTenantValAPreflightGateStateActive,
		IdentityBootstrapState:       DeploymentMultiTenantValAIdentityBootstrapStateActive,
		AirGappedEvidenceBundleState: DeploymentMultiTenantValAAirGappedEvidenceBundleStateActive,
		NoOverclaimState:             DeploymentMultiTenantValANoOverclaimStateActive,
		PassBlockerState:             DeploymentMultiTenantValAPassBlockerStateActive,
		Dependency:                   deploymentMultiTenantValADependencySnapshotModel(),
		DeploymentProfileMatrix: DeploymentMultiTenantValADeploymentProfileMatrix{
			CurrentState:                            DeploymentMultiTenantValADeploymentProfileMatrixStateActive,
			SupportedProfiles:                       append([]string{}, deploymentMultiTenantValARequiredProfiles()...),
			SupportedStates:                         append([]string{}, deploymentMultiTenantValARequiredStates()...),
			EvidenceRefs:                            append([]string{}, deploymentMultiTenantValADeploymentProfileMatrixEvidenceRefs()...),
			FreshnessState:                          IntelligenceCalibrationFreshnessFresh,
			ReadinessEvidenceBacked:                 true,
			InstallSucceeded:                        true,
			MarketplaceInstallSucceeded:             true,
			SaaSState:                               DeploymentMultiTenantDeploymentStateReady,
			SaaSEvidenceKeys:                        append([]string{}, deploymentMultiTenantValAProfileMatrixSaaSKeys()...),
			SaaSTenantConfig:                        "tenant_config_bundle",
			SaaSRegion:                              "eu_central_1",
			SaaSIdentityBootstrap:                   "tenant_identity_bootstrap",
			SaaSConnectorScope:                      "tenant_scoped_connector_execution",
			SaaSBackupPolicy:                        "tenant_backup_policy",
			SaaSOperatorSupportScope:                "tenant_operator_support_boundary",
			SelfHostedState:                         DeploymentMultiTenantDeploymentStateReady,
			SelfHostedEvidenceKeys:                  append([]string{}, deploymentMultiTenantValAProfileMatrixSelfHostedKeys()...),
			SelfHostedArtifactProvenance:            "artifact_provenance_bundle",
			SelfHostedEnvironmentManifest:           "environment_manifest_bundle",
			SelfHostedConfigValidation:              "config_validation_evidence",
			SelfHostedIAMKMSDependency:              "iam_kms_dependency_contract",
			SelfHostedBackupTarget:                  "tenant_backup_target",
			SelfHostedUpgradeRollbackPlan:           "upgrade_rollback_plan",
			SelfHostedUnsupportedSemanticsExplicit:  true,
			SelfHostedDegradedSemanticsExplicit:     true,
			AirGappedState:                          DeploymentMultiTenantDeploymentStateReady,
			AirGappedEvidenceKeys:                   append([]string{}, deploymentMultiTenantValAProfileMatrixAirGappedKeys()...),
			AirGappedOfflineArtifactBundle:          "offline_artifact_bundle",
			AirGappedOfflineEvidenceBundle:          "offline_evidence_bundle",
			AirGappedSignatureHashVerificationState: DeploymentMultiTenantValASignatureVerificationVerified,
			AirGappedUnsupportedOnlineDependencies:  []string{"none_explicit"},
			AirGappedOfflineReplayExportPath:        "offline_replay_export_path",
			AirGappedUnsupportedSemanticsExplicit:   true,
			AirGappedDegradedSemanticsExplicit:      true,
			ProfileNamingExact:                      true,
			SafeReadinessWordingExamplePresent:      true,
			DiagnosticOutputComplete:                true,
			ObservedClaims: []string{
				"validated deployment profile",
				"self-hosted readiness evidence",
				"unsupported dependency explicitly listed",
			},
			ProjectionDisclaimer: disclaimer,
		},
		PreflightGate: DeploymentMultiTenantValAPreflightGate{
			CurrentState:                    DeploymentMultiTenantValAPreflightGateStateActive,
			EvidenceRefs:                    append([]string{}, deploymentMultiTenantValAPreflightEvidenceRefs()...),
			CheckKeys:                       append([]string{}, deploymentMultiTenantValAPreflightCheckKeys()...),
			FreshnessState:                  IntelligenceCalibrationFreshnessFresh,
			TenantScope:                     "tenant:alpha",
			InstallConfigValidation:         true,
			UpgradeConfigDiff:               true,
			DBSchemaMigrationDryRun:         true,
			BackupBeforeUpgradeEvidence:     true,
			RollbackPlanEvidence:            true,
			PolicyMigrationCompatibility:    true,
			ConnectorPermissionChangeReview: true,
			KeyRotationReadiness:            true,
			TenantBoundaryValidation:        true,
			DiagnosticOutputComplete:        true,
			ProjectionDisclaimer:            disclaimer,
		},
		IdentityBootstrap: DeploymentMultiTenantValAIdentityBootstrap{
			CurrentState:                   DeploymentMultiTenantValAIdentityBootstrapStateActive,
			SupportedModes:                 append([]string{}, deploymentMultiTenantValAIdentityModes()...),
			EvidenceRefs:                   append([]string{}, deploymentMultiTenantValAIdentityBootstrapEvidenceRefs()...),
			ValidationKeys:                 append([]string{}, deploymentMultiTenantValAIdentityValidationKeys()...),
			FreshnessState:                 IntelligenceCalibrationFreshnessFresh,
			IdentityMode:                   DeploymentMultiTenantValAIdentityModeOIDC,
			TenantScope:                    "tenant:alpha",
			IssuerEntityID:                 "tenant_identity_issuer",
			CallbackRedirectURL:            "tenant_identity_callback_url",
			CertificateExpiryState:         DeploymentMultiTenantValACertificateStateValid,
			GroupRoleMapping:               "tenant_group_role_mapping",
			AdminBootstrapFallback:         "tenant_admin_bootstrap_fallback",
			BreakGlassCompatibility:        true,
			BreakGlassExpiryPresent:        true,
			BreakGlassRevocationPresent:    true,
			DisabledUnsafeFallbackHandling: true,
			TenantSpecificIdentityBoundary: "tenant_specific_identity_boundary",
			DiagnosticOutputComplete:       true,
			ObservedClaims:                 []string{"sso bootstrap validation", "saml/oidc validation evidence"},
			ProjectionDisclaimer:           disclaimer,
		},
		AirGappedEvidenceBundle: DeploymentMultiTenantValAAirGappedEvidenceBundle{
			CurrentState:                        DeploymentMultiTenantValAAirGappedEvidenceBundleStateActive,
			EvidenceRefs:                        append([]string{}, deploymentMultiTenantValAAirGappedEvidenceBundleEvidenceRefs()...),
			BundleKeys:                          append([]string{}, deploymentMultiTenantValAAirGappedBundleKeys()...),
			FreshnessState:                      IntelligenceCalibrationFreshnessFresh,
			AirGappedModeExplicit:               true,
			OfflineEvidenceBundleEvidenceBacked: true,
			BundleManifest:                      "bundle_manifest",
			ArtifactHashes:                      "artifact_hashes",
			ProofPackHashes:                     "proof_pack_hashes",
			Signer:                              "offline_bundle_signer",
			PolicyVersion:                       "policy_version",
			EngineVersion:                       "engine_version",
			Timestamp:                           "timestamp_value",
			UnsupportedOnlineDependencies:       []string{"none_explicit"},
			ReplayInstructions:                  "replay_instructions",
			OfflineReplayExportPath:             "offline_replay_export_path",
			SignatureHashVerificationState:      DeploymentMultiTenantValASignatureVerificationVerified,
			DiagnosticOutputComplete:            true,
			ObservedClaims:                      []string{"air-gapped offline evidence bundle", "offline replay/export path"},
			ProjectionDisclaimer:                disclaimer,
		},
		NoOverclaim: DeploymentMultiTenantValANoOverclaimDiscipline{
			CurrentState:         DeploymentMultiTenantValANoOverclaimStateActive,
			ProjectionDisclaimer: disclaimer,
		},
		PassBlockerOverlay: DeploymentMultiTenantValAPassBlockerOverlay{
			CurrentState:         DeploymentMultiTenantValAPassBlockerStateActive,
			ProjectionDisclaimer: disclaimer,
		},
	}
}

func ComputeDeploymentMultiTenantValAFoundation(model DeploymentMultiTenantValAFoundation) DeploymentMultiTenantValAFoundation {
	model.DependencyState = EvaluateDeploymentMultiTenantValADependencyState(model.Dependency)
	model.DeploymentProfileMatrixState = EvaluateDeploymentMultiTenantValADeploymentProfileMatrixState(model.DeploymentProfileMatrix)
	model.PreflightGateState = EvaluateDeploymentMultiTenantValAPreflightGateState(model.PreflightGate)
	model.IdentityBootstrapState = EvaluateDeploymentMultiTenantValAIdentityBootstrapState(model.IdentityBootstrap)
	model.AirGappedEvidenceBundleState = EvaluateDeploymentMultiTenantValAAirGappedEvidenceBundleState(model.AirGappedEvidenceBundle)
	model.NoOverclaimState = EvaluateDeploymentMultiTenantValANoOverclaimState(model.NoOverclaim)
	model.PassBlockerOverlay = DeploymentMultiTenantValAPassBlockerOverlay{
		ProjectionDisclaimer: deploymentMultiTenantValAProjectionDisclaimer(),
		Findings:             deploymentMultiTenantValAPassBlockerFindings(model),
	}
	model.PassBlockerState = EvaluateDeploymentMultiTenantValAPassBlockerState(model.PassBlockerOverlay)
	model.PassBlockerOverlay.CurrentState = model.PassBlockerState
	model.Point10State = EvaluateDeploymentMultiTenantPoint10State(model.CurrentState)
	model.CurrentState = EvaluateDeploymentMultiTenantValAState(model)
	model.Point10State = EvaluateDeploymentMultiTenantPoint10State(model.CurrentState)
	model.BlockingReasons = deploymentMultiTenantValABlockingReasons(model)
	model.DeploymentProfileMatrix.CurrentState = model.DeploymentProfileMatrixState
	model.PreflightGate.CurrentState = model.PreflightGateState
	model.IdentityBootstrap.CurrentState = model.IdentityBootstrapState
	model.AirGappedEvidenceBundle.CurrentState = model.AirGappedEvidenceBundleState
	model.NoOverclaim.CurrentState = model.NoOverclaimState
	return model
}

package operability

import "strings"

const (
	DeploymentMultiTenantValDStateActive  = "deployment_multi_tenant_vald_active"
	DeploymentMultiTenantValDStateBlocked = "deployment_multi_tenant_vald_blocked"

	DeploymentMultiTenantValDDependencyStateActive  = "deployment_multi_tenant_vald_dependency_active"
	DeploymentMultiTenantValDDependencyStateBlocked = "deployment_multi_tenant_vald_dependency_blocked"

	DeploymentMultiTenantValDConnectorCapabilityStateActive  = "deployment_multi_tenant_vald_connector_capability_active"
	DeploymentMultiTenantValDConnectorCapabilityStateBlocked = "deployment_multi_tenant_vald_connector_capability_blocked"

	DeploymentMultiTenantValDOperatorActionStateActive  = "deployment_multi_tenant_vald_operator_action_active"
	DeploymentMultiTenantValDOperatorActionStateBlocked = "deployment_multi_tenant_vald_operator_action_blocked"

	DeploymentMultiTenantValDSupportAccessStateActive  = "deployment_multi_tenant_vald_support_access_active"
	DeploymentMultiTenantValDSupportAccessStateBlocked = "deployment_multi_tenant_vald_support_access_blocked"

	DeploymentMultiTenantValDBreakGlassStateActive  = "deployment_multi_tenant_vald_break_glass_active"
	DeploymentMultiTenantValDBreakGlassStateBlocked = "deployment_multi_tenant_vald_break_glass_blocked"

	DeploymentMultiTenantValDMarketplaceMSPAuthorityStateActive  = "deployment_multi_tenant_vald_marketplace_msp_authority_active"
	DeploymentMultiTenantValDMarketplaceMSPAuthorityStateBlocked = "deployment_multi_tenant_vald_marketplace_msp_authority_blocked"

	DeploymentMultiTenantValDAgenticOverlayStateActive  = "deployment_multi_tenant_vald_agentic_overlay_active"
	DeploymentMultiTenantValDAgenticOverlayStateBlocked = "deployment_multi_tenant_vald_agentic_overlay_blocked"

	DeploymentMultiTenantValDAgentLearningLoopStateActive  = "deployment_multi_tenant_vald_agent_learning_loop_active"
	DeploymentMultiTenantValDAgentLearningLoopStateBlocked = "deployment_multi_tenant_vald_agent_learning_loop_blocked"

	DeploymentMultiTenantValDNoOverclaimStateActive  = "deployment_multi_tenant_vald_no_overclaim_active"
	DeploymentMultiTenantValDNoOverclaimStateBlocked = "deployment_multi_tenant_vald_no_overclaim_blocked"

	DeploymentMultiTenantValDClosureBlockerStateActive   = "deployment_multi_tenant_vald_closure_blocker_active"
	DeploymentMultiTenantValDClosureBlockerStateCleanup  = "deployment_multi_tenant_vald_closure_blocker_cleanup"
	DeploymentMultiTenantValDClosureBlockerStateAdvisory = "deployment_multi_tenant_vald_closure_blocker_advisory"
	DeploymentMultiTenantValDClosureBlockerStateBlocked  = "deployment_multi_tenant_vald_closure_blocker_blocked"

	DeploymentMultiTenantValDBlockerLevelCLB0 = "CL-B0"
	DeploymentMultiTenantValDBlockerLevelCLB1 = "CL-B1"
	DeploymentMultiTenantValDBlockerLevelCLB2 = "CL-B2"
	DeploymentMultiTenantValDBlockerLevelCLB3 = "CL-B3"

	DeploymentMultiTenantValDClosureSurfaceConnectorSandbox = "connector_sandbox"
	DeploymentMultiTenantValDClosureSurfaceOperatorAction   = "operator_action"
	DeploymentMultiTenantValDClosureSurfaceSupportAccess    = "support_access"
	DeploymentMultiTenantValDClosureSurfaceBreakGlass       = "break_glass"
	DeploymentMultiTenantValDClosureSurfaceMarketplaceMSP   = "marketplace_msp"
	DeploymentMultiTenantValDClosureSurfaceAgenticOverlay   = "agentic_overlay"
	DeploymentMultiTenantValDClosureSurfaceNoOverclaim      = "no_overclaim"
	DeploymentMultiTenantValDClosureSurfaceCleanRoomIP      = "clean_room_ip"

	deploymentMultiTenantValDApprovalStatusApproved = "approved_active"
	deploymentMultiTenantValDBreakGlassReviewActive = "post_review_active"
)

type DeploymentMultiTenantValDDependencySnapshot struct {
	ValCCurrentState           string `json:"valc_current_state"`
	ValCDependencyState        string `json:"valc_dependency_state"`
	ValCHAReadinessState       string `json:"valc_ha_readiness_state"`
	ValCRecoveryReadinessState string `json:"valc_recovery_readiness_state"`
	ValCSLAReadinessState      string `json:"valc_sla_readiness_state"`
	ValCTenantTrustScopeState  string `json:"valc_tenant_trust_scope_state"`
	ValCSiloVisibilityState    string `json:"valc_silo_visibility_state"`
	ValCPrivacyGuardState      string `json:"valc_privacy_guard_state"`
	ValCNoOverclaimState       string `json:"valc_no_overclaim_state"`
	ValCClosureBlockerState    string `json:"valc_closure_blocker_state"`
	Point10State               string `json:"point_10_state"`
	ProjectionDisclaimer       string `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValDConnectorCapabilityManifest struct {
	CurrentState                            string   `json:"current_state"`
	ConnectorID                             string   `json:"connector_id"`
	ConnectorType                           string   `json:"connector_type"`
	TenantScope                             string   `json:"tenant_scope"`
	PermissionManifest                      string   `json:"permission_manifest"`
	CapabilityManifestPresent               bool     `json:"capability_manifest_present"`
	ReadCapabilities                        []string `json:"read_capabilities,omitempty"`
	WriteCapabilities                       []string `json:"write_capabilities,omitempty"`
	MutationAllowed                         bool     `json:"mutation_allowed"`
	MutationCapabilityExplicit              bool     `json:"mutation_capability_explicit"`
	EvidenceTypes                           []string `json:"evidence_types,omitempty"`
	RetryPolicy                             string   `json:"retry_policy"`
	ReplayPolicy                            string   `json:"replay_policy"`
	RateLimitPolicy                         string   `json:"rate_limit_policy"`
	AuditRequired                           bool     `json:"audit_required"`
	AuditID                                 string   `json:"audit_id"`
	Reason                                  string   `json:"reason"`
	SourceOfTruth                           string   `json:"source_of_truth"`
	TenantScopedExecution                   bool     `json:"tenant_scoped_execution"`
	RecoveryBehavior                        string   `json:"recovery_behavior"`
	FailureBehavior                         string   `json:"failure_behavior"`
	EvidenceRefs                            []string `json:"evidence_refs,omitempty"`
	FreshnessState                          string   `json:"freshness_state"`
	ProjectionDisclaimer                    string   `json:"projection_disclaimer"`
	DiagnosticOutputComplete                bool     `json:"diagnostic_output_complete"`
	ConnectorAsSourceOfTruth                bool     `json:"connector_as_source_of_truth"`
	ConnectorBypassesDeploymentGate         bool     `json:"connector_bypasses_deployment_gate"`
	ConnectorBypassesTenantGate             bool     `json:"connector_bypasses_tenant_gate"`
	ConnectorBypassesEvidenceGate           bool     `json:"connector_bypasses_evidence_gate"`
	ConnectorBypassesDataResidencyGate      bool     `json:"connector_bypasses_data_residency_gate"`
	RetryReplayDuplicatesActiveEvidenceRisk bool     `json:"retry_replay_duplicates_active_evidence_risk"`
	ConnectorNamingExact                    bool     `json:"connector_naming_exact"`
	SafeConnectorWordingExamplePresent      bool     `json:"safe_connector_wording_example_present"`
}

type DeploymentMultiTenantValDOperatorActionPolicy struct {
	CurrentState                      string   `json:"current_state"`
	Actor                             string   `json:"actor"`
	ActorType                         string   `json:"actor_type"`
	TenantTarget                      string   `json:"tenant_target"`
	ActionScope                       string   `json:"action_scope"`
	ActionType                        string   `json:"action_type"`
	Reason                            string   `json:"reason"`
	AuthorizationBasis                string   `json:"authorization_basis"`
	Approver                          string   `json:"approver"`
	ApprovalStatus                    string   `json:"approval_status"`
	AuthorityBasis                    string   `json:"authority_basis"`
	Expiry                            string   `json:"expiry"`
	RevocationPath                    string   `json:"revocation_path"`
	AuditID                           string   `json:"audit_id"`
	EvidenceRefs                      []string `json:"evidence_refs,omitempty"`
	RBACABACEnforced                  bool     `json:"rbac_abac_enforced"`
	SSOContextBound                   bool     `json:"sso_context_bound"`
	TenantScopeBound                  bool     `json:"tenant_scope_bound"`
	SupportScopeBound                 bool     `json:"support_scope_bound"`
	ProductionMutationAllowed         bool     `json:"production_mutation_allowed"`
	CanonicalMutationAllowed          bool     `json:"canonical_mutation_allowed"`
	DiagnosticOutputComplete          bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer              string   `json:"projection_disclaimer"`
	OperatorActionCanonicalApproval   bool     `json:"operator_action_canonical_approval"`
	OperatorActionNamingExact         bool     `json:"operator_action_naming_exact"`
	SafeOperatorWordingExamplePresent bool     `json:"safe_operator_wording_example_present"`
}

type DeploymentMultiTenantValDSupportAccessEnforcement struct {
	CurrentState                     string   `json:"current_state"`
	SupportAccessID                  string   `json:"support_access_id"`
	SupportActor                     string   `json:"support_actor"`
	TenantTarget                     string   `json:"tenant_target"`
	SupportScope                     string   `json:"support_scope"`
	SSOSessionReference              string   `json:"sso_session_reference"`
	RBACRole                         string   `json:"rbac_role"`
	ABACConditions                   string   `json:"abac_conditions"`
	AuthorityBasis                   string   `json:"authority_basis"`
	Reason                           string   `json:"reason"`
	Approver                         string   `json:"approver"`
	ApprovalStatus                   string   `json:"approval_status"`
	Expiry                           string   `json:"expiry"`
	RevocationPath                   string   `json:"revocation_path"`
	AuditID                          string   `json:"audit_id"`
	EvidenceRefs                     []string `json:"evidence_refs,omitempty"`
	SupportVisibilityBoundary        string   `json:"support_visibility_boundary"`
	DataResidencyBoundaryRespected   bool     `json:"data_residency_boundary_respected"`
	TenantIsolationBoundaryRespected bool     `json:"tenant_isolation_boundary_respected"`
	DiagnosticOutputComplete         bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
	SupportVisibilityExceedsScope    bool     `json:"support_visibility_exceeds_scope"`
	RawTenantEvidenceExposed         bool     `json:"raw_tenant_evidence_exposed"`
	SupportSummaryCanonicalTruth     bool     `json:"support_summary_canonical_truth"`
}

type DeploymentMultiTenantValDBreakGlassAccess struct {
	CurrentState                     string   `json:"current_state"`
	BreakGlassID                     string   `json:"break_glass_id"`
	EmergencyReason                  string   `json:"emergency_reason"`
	Actor                            string   `json:"actor"`
	TenantTarget                     string   `json:"tenant_target"`
	ActionScope                      string   `json:"action_scope"`
	AuthorizationBasis               string   `json:"authorization_basis"`
	Approver                         string   `json:"approver"`
	ApprovalStatus                   string   `json:"approval_status"`
	Expiry                           string   `json:"expiry"`
	RevocationPath                   string   `json:"revocation_path"`
	AuditID                          string   `json:"audit_id"`
	EvidenceRefs                     []string `json:"evidence_refs,omitempty"`
	PostActionReviewRequired         bool     `json:"post_action_review_required"`
	PostActionReviewState            string   `json:"post_action_review_state"`
	TenantScopeBound                 bool     `json:"tenant_scope_bound"`
	DataResidencyBoundaryRespected   bool     `json:"data_residency_boundary_respected"`
	TenantIsolationBoundaryRespected bool     `json:"tenant_isolation_boundary_respected"`
	SupportVisibilityBoundary        string   `json:"support_visibility_boundary"`
	DiagnosticOutputComplete         bool     `json:"diagnostic_output_complete"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
	PersistentAccessGranted          bool     `json:"persistent_access_granted"`
	CreatesPASSAuthority             bool     `json:"creates_pass_authority"`
}

type DeploymentMultiTenantValDMarketplaceMSPAuthorityBoundary struct {
	CurrentState                            string   `json:"current_state"`
	MarketplaceProfile                      string   `json:"marketplace_profile"`
	MSPOperatorScope                        string   `json:"msp_operator_scope"`
	PartnerScope                            string   `json:"partner_scope"`
	TenantScope                             string   `json:"tenant_scope"`
	DeploymentScope                         string   `json:"deployment_scope"`
	SupportScope                            string   `json:"support_scope"`
	CustomerReadyValidationEvidence         string   `json:"customer_ready_validation_evidence"`
	CustomerReadyWordingPresent             bool     `json:"customer_ready_wording_present"`
	AuthorityBoundary                       string   `json:"authority_boundary"`
	ApprovalBoundary                        string   `json:"approval_boundary"`
	PassAuthorityAllowed                    bool     `json:"pass_authority_allowed"`
	ProductionReadinessAuthorityAllowed     bool     `json:"production_readiness_authority_allowed"`
	SourceOfTruthAllowed                    bool     `json:"source_of_truth_allowed"`
	AuditID                                 string   `json:"audit_id"`
	Reason                                  string   `json:"reason"`
	Expiry                                  string   `json:"expiry"`
	RevocationPath                          string   `json:"revocation_path"`
	EvidenceRefs                            []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer                    string   `json:"projection_disclaimer"`
	DiagnosticOutputComplete                bool     `json:"diagnostic_output_complete"`
	MarketplaceInstallTreatedAsProdApproved bool     `json:"marketplace_install_treated_as_production_approved"`
	MSPApprovedDeploymentClaim              bool     `json:"msp_approved_deployment_claim"`
	PartnerCertifiedDeploymentClaim         bool     `json:"partner_certified_deployment_claim"`
}

type DeploymentMultiTenantValDAgentRecommendation struct {
	AgentID                                       string   `json:"agent_id"`
	AgentType                                     string   `json:"agent_type"`
	TenantScope                                   string   `json:"tenant_scope"`
	PermissionManifest                            string   `json:"permission_manifest"`
	AllowedReadSurfaces                           []string `json:"allowed_read_surfaces,omitempty"`
	AllowedRecommendationSurfaces                 []string `json:"allowed_recommendation_surfaces,omitempty"`
	ForbiddenMutationSurfaces                     []string `json:"forbidden_mutation_surfaces,omitempty"`
	ExternalAPIAllowed                            bool     `json:"external_api_allowed"`
	ApprovalRequired                              bool     `json:"approval_required"`
	ApprovalStatus                                string   `json:"approval_status"`
	Approver                                      string   `json:"approver"`
	ApprovalReason                                string   `json:"approval_reason"`
	ApprovalQueue                                 string   `json:"approval_queue"`
	AuditID                                       string   `json:"audit_id"`
	EvidenceRefs                                  []string `json:"evidence_refs,omitempty"`
	RecommendationID                              string   `json:"recommendation_id"`
	RecommendationType                            string   `json:"recommendation_type"`
	HumanReviewRequired                           bool     `json:"human_review_required"`
	ExecutionAllowed                              bool     `json:"execution_allowed"`
	CanonicalMutationAllowed                      bool     `json:"canonical_mutation_allowed"`
	ProductionMutationAllowed                     bool     `json:"production_mutation_allowed"`
	Point10PassAllowed                            bool     `json:"point10_pass_allowed"`
	ProjectionDisclaimer                          string   `json:"projection_disclaimer"`
	DiagnosticOutputComplete                      bool     `json:"diagnostic_output_complete"`
	AgentTreatedAsSourceOfTruth                   bool     `json:"agent_treated_as_source_of_truth"`
	CrossTenantAccess                             bool     `json:"cross_tenant_access"`
	ContainmentExecutedWithoutApproval            bool     `json:"containment_executed_without_approval"`
	ChangesTenantStateDirectly                    bool     `json:"changes_tenant_state_directly"`
	ExpandsOrLimitsAccessDirectlyWithoutApproval  bool     `json:"expands_or_limits_access_directly_without_approval"`
	InstallSuccessTreatedAsReadiness              bool     `json:"install_success_treated_as_readiness"`
	PreflightExecutedWithoutApproval              bool     `json:"preflight_executed_without_approval"`
	MarksDeploymentReadyWithoutCanonicalEvaluator bool     `json:"marks_deployment_ready_without_canonical_evaluator"`
	SLAReadinessTreatedAsUptimeGuarantee          bool     `json:"sla_readiness_treated_as_uptime_guarantee"`
	ConnectorMutationExecutedByAgent              bool     `json:"connector_mutation_executed_by_agent"`
	ConnectorCapabilityMissing                    bool     `json:"connector_capability_missing"`
	OperatorSupportActionWithoutAuthorityBasis    bool     `json:"operator_support_action_without_authority_basis"`
	BreakGlassExpiryRevocationMissing             bool     `json:"break_glass_expiry_revocation_missing"`
	RestoreRollbackRebuildExecutedAutomatically   bool     `json:"restore_rollback_rebuild_executed_automatically"`
	RecoveryEvidencePack                          string   `json:"recovery_evidence_pack"`
	RecommendationBypassesTenantIsolation         bool     `json:"recommendation_bypasses_tenant_isolation"`
	RecommendationBypassesDataResidency           bool     `json:"recommendation_bypasses_data_residency"`
	RecoveryGuaranteedClaim                       bool     `json:"recovery_guaranteed_claim"`
	TreatsRecommendationAsApproval                bool     `json:"treats_recommendation_as_approval"`
	AgentNamingExact                              bool     `json:"agent_naming_exact"`
	SafeWordingExamplePresent                     bool     `json:"safe_wording_example_present"`
	RunbookWordingComplete                        bool     `json:"runbook_wording_complete"`
}

type DeploymentMultiTenantValDAgentRuntimeApprovalController = DeploymentMultiTenantValDAgentRecommendation

type DeploymentMultiTenantValDAgentLearningLoop struct {
	LearningAllowed                              bool     `json:"learning_allowed"`
	LearningMode                                 string   `json:"learning_mode"`
	TrainingDataScope                            string   `json:"training_data_scope"`
	TrainingDataRefs                             []string `json:"training_data_refs,omitempty"`
	TrainingDataPrivacyFiltered                  bool     `json:"training_data_privacy_filtered"`
	TrainingDataTenantScoped                     bool     `json:"training_data_tenant_scoped"`
	TrainingDataCrossTenant                      bool     `json:"training_data_cross_tenant"`
	TrainingDataCustomerApproved                 bool     `json:"training_data_customer_approved"`
	HumanFeedbackRefs                            []string `json:"human_feedback_refs,omitempty"`
	HumanFeedbackAuditLinked                     bool     `json:"human_feedback_audit_linked"`
	TrainingApprovalRequired                     bool     `json:"training_approval_required"`
	TrainingApprovalStatus                       string   `json:"training_approval_status"`
	TrainingApprover                             string   `json:"training_approver"`
	ModelCandidateID                             string   `json:"model_candidate_id"`
	ModelVersion                                 string   `json:"model_version"`
	BaselineModelVersion                         string   `json:"baseline_model_version"`
	EvaluationResultRefs                         []string `json:"evaluation_result_refs,omitempty"`
	RegressionTestRefs                           []string `json:"regression_test_refs,omitempty"`
	NoOverclaimCheckRefs                         []string `json:"no_overclaim_check_refs,omitempty"`
	TenantScopeCheckRefs                         []string `json:"tenant_scope_check_refs,omitempty"`
	ApprovalGateCheckRefs                        []string `json:"approval_gate_check_refs,omitempty"`
	PromotionAllowed                             bool     `json:"promotion_allowed"`
	PromotionApprovalStatus                      string   `json:"promotion_approval_status"`
	RuntimeActivationAllowed                     bool     `json:"runtime_activation_allowed"`
	RuntimeActivationApprovalStatus              string   `json:"runtime_activation_approval_status"`
	RecommendationApprovalStatus                 string   `json:"recommendation_approval_status"`
	ExecutionApprovalStatus                      string   `json:"execution_approval_status"`
	ModelUpgradeApprovalStatus                   string   `json:"model_upgrade_approval_status"`
	LearnedOutputAdvisoryOnly                    bool     `json:"learned_output_advisory_only"`
	ExternalAPIAllowed                           bool     `json:"external_api_allowed"`
	ProductionSelfModificationAllowed            bool     `json:"production_self_modification_allowed"`
	ProductionMutationAllowed                    bool     `json:"production_mutation_allowed"`
	CanonicalMutationAllowed                     bool     `json:"canonical_mutation_allowed"`
	Point10PassAllowed                           bool     `json:"point10_pass_allowed"`
	ApprovalRequired                             bool     `json:"approval_required"`
	AuditID                                      string   `json:"audit_id"`
	EvidenceRefs                                 []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer                         string   `json:"projection_disclaimer"`
	DiagnosticOutputComplete                     bool     `json:"diagnostic_output_complete"`
	CandidateWeakensNoOverclaim                  bool     `json:"candidate_weakens_no_overclaim"`
	CandidateWeakensTenantIsolation              bool     `json:"candidate_weakens_tenant_isolation"`
	CandidateWeakensApprovalGates                bool     `json:"candidate_weakens_approval_gates"`
	CandidateEnablesExternalAPIByDefault         bool     `json:"candidate_enables_external_api_by_default"`
	CandidateEnablesProductionMutation           bool     `json:"candidate_enables_production_mutation"`
	CandidateEnablesCanonicalMutation            bool     `json:"candidate_enables_canonical_mutation"`
	CandidateEnablesPoint10Pass                  bool     `json:"candidate_enables_point10_pass"`
	AgentSelfPromotes                            bool     `json:"agent_self_promotes"`
	AgentSelfDeploys                             bool     `json:"agent_self_deploys"`
	AgentModifiesProductionPolicy                bool     `json:"agent_modifies_production_policy"`
	RecommendationApprovalMeansExecutionApproval bool     `json:"recommendation_approval_means_execution_approval"`
	ExecutionApprovalMeansModelUpgradeApproval   bool     `json:"execution_approval_means_model_upgrade_approval"`
	LearnedOutputTreatedAsCanonicalTruth         bool     `json:"learned_output_treated_as_canonical_truth"`
	LearningModeNamingExact                      bool     `json:"learning_mode_naming_exact"`
	ModelCandidateNamingExact                    bool     `json:"model_candidate_naming_exact"`
	SafeWordingExamplePresent                    bool     `json:"safe_wording_example_present"`
	RunbookWordingComplete                       bool     `json:"runbook_wording_complete"`
}

type DeploymentMultiTenantValDAgenticOverlay struct {
	CurrentState                       string                                                  `json:"current_state"`
	LearningLoopState                  string                                                  `json:"learning_loop_state"`
	RuntimeApprovalController          DeploymentMultiTenantValDAgentRuntimeApprovalController `json:"runtime_approval_controller"`
	TenantBoundaryContainmentAgent     DeploymentMultiTenantValDAgentRecommendation            `json:"tenant_boundary_containment_agent"`
	DeploymentHealthPreflightAgent     DeploymentMultiTenantValDAgentRecommendation            `json:"deployment_health_preflight_agent"`
	ConnectorOperatorMisuseWatchAgent  DeploymentMultiTenantValDAgentRecommendation            `json:"connector_operator_misuse_watch_agent"`
	RecoveryRebuildRecommendationAgent DeploymentMultiTenantValDAgentRecommendation            `json:"recovery_rebuild_recommendation_agent"`
	LearningLoop                       DeploymentMultiTenantValDAgentLearningLoop              `json:"learning_loop"`
	ProjectionDisclaimer               string                                                  `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValDNoOverclaimDiscipline struct {
	CurrentState                 string   `json:"current_state"`
	ObservedClaims               []string `json:"observed_claims,omitempty"`
	CleanRoomIPViolationDetected bool     `json:"clean_room_ip_violation_detected"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValDClosureBlockerFinding struct {
	BlockerLevel      string `json:"blocker_level"`
	Surface           string `json:"surface"`
	Reason            string `json:"reason"`
	BlocksCurrentWave bool   `json:"blocks_current_wave"`
	RequiredFollowup  string `json:"required_followup,omitempty"`
}

type DeploymentMultiTenantValDClosureBlockerOverlay struct {
	CurrentState         string                                           `json:"current_state"`
	Findings             []DeploymentMultiTenantValDClosureBlockerFinding `json:"findings,omitempty"`
	ProjectionDisclaimer string                                           `json:"projection_disclaimer"`
}

type DeploymentMultiTenantValDFoundation struct {
	CurrentState                 string                                                   `json:"current_state"`
	Point10State                 string                                                   `json:"point_10_state"`
	ProjectionDisclaimer         string                                                   `json:"projection_disclaimer"`
	BlockingReasons              []string                                                 `json:"blocking_reasons,omitempty"`
	DependencyState              string                                                   `json:"dependency_state"`
	ConnectorCapabilityState     string                                                   `json:"connector_capability_state"`
	OperatorActionState          string                                                   `json:"operator_action_state"`
	SupportAccessState           string                                                   `json:"support_access_state"`
	BreakGlassState              string                                                   `json:"break_glass_state"`
	MarketplaceMSPAuthorityState string                                                   `json:"marketplace_msp_authority_state"`
	AgenticOverlayState          string                                                   `json:"agentic_overlay_state"`
	NoOverclaimState             string                                                   `json:"no_overclaim_state"`
	ClosureBlockerState          string                                                   `json:"closure_blocker_state"`
	Dependency                   DeploymentMultiTenantValDDependencySnapshot              `json:"dependency"`
	ConnectorCapability          DeploymentMultiTenantValDConnectorCapabilityManifest     `json:"connector_capability"`
	OperatorAction               DeploymentMultiTenantValDOperatorActionPolicy            `json:"operator_action"`
	SupportAccess                DeploymentMultiTenantValDSupportAccessEnforcement        `json:"support_access"`
	BreakGlass                   DeploymentMultiTenantValDBreakGlassAccess                `json:"break_glass"`
	MarketplaceMSPAuthority      DeploymentMultiTenantValDMarketplaceMSPAuthorityBoundary `json:"marketplace_msp_authority"`
	AgenticOverlay               DeploymentMultiTenantValDAgenticOverlay                  `json:"agentic_overlay"`
	NoOverclaim                  DeploymentMultiTenantValDNoOverclaimDiscipline           `json:"no_overclaim"`
	ClosureBlockerOverlay        DeploymentMultiTenantValDClosureBlockerOverlay           `json:"closure_blocker_overlay"`
}

func deploymentMultiTenantValDProjectionDisclaimer() string {
	return "projection_only not_canonical_truth bounded_marketplace_deployment_profile deployment_multi_tenant_vald"
}

func deploymentMultiTenantValDHasProjectionDisclaimer(value string) bool {
	return value == deploymentMultiTenantValDProjectionDisclaimer() ||
		value == deploymentMultiTenantValDProjectionDisclaimer()+" aggregate_dependency_snapshot" ||
		value == "projection_only not_canonical_truth deployment_multi_tenant_vald aggregate_dependency_snapshot"
}

func deploymentMultiTenantValDHasFoundationProjectionDisclaimer(value string) bool {
	return value == deploymentMultiTenantValDProjectionDisclaimer()
}

func deploymentMultiTenantValDConnectorEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vald-connector-capability-001"}
}

func deploymentMultiTenantValDOperatorEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vald-operator-action-001"}
}

func deploymentMultiTenantValDSupportEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vald-support-access-001"}
}

func deploymentMultiTenantValDBreakGlassEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vald-break-glass-001"}
}

func deploymentMultiTenantValDMarketplaceMSPEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vald-marketplace-msp-001"}
}

func deploymentMultiTenantValDAgenticOverlayEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vald-agentic-overlay-001"}
}

func deploymentMultiTenantValDAgentLearningLoopEvidenceRefs() []string {
	return []string{"evidence:deployment-multi-tenant-vald-agent-learning-loop-001"}
}

func deploymentMultiTenantValDAgentLearningLoopHumanFeedbackRefs() []string {
	return []string{"human_feedback_ref"}
}

func deploymentMultiTenantValDAgentLearningLoopEvaluationResultRefs() []string {
	return []string{"evaluation_result_ref"}
}

func deploymentMultiTenantValDAgentLearningLoopRegressionTestRefs() []string {
	return []string{"regression_test_ref"}
}

func deploymentMultiTenantValDAgentLearningLoopNoOverclaimCheckRefs() []string {
	return []string{"no_overclaim_check_ref"}
}

func deploymentMultiTenantValDAgentLearningLoopTenantScopeCheckRefs() []string {
	return []string{"tenant_scope_check_ref"}
}

func deploymentMultiTenantValDAgentLearningLoopApprovalGateCheckRefs() []string {
	return []string{"approval_gate_check_ref"}
}

func deploymentMultiTenantValDEvidenceValueIsValid(value string) bool {
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

func deploymentMultiTenantValDHasExactEvidenceRefs(values []string, expected ...string) bool {
	return deploymentMultiTenantVal0ContainsExactStringSet(values, expected...)
}

func deploymentMultiTenantValDAllExactValuesValid(values []string) bool {
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

func deploymentMultiTenantValDHasRevokedExpiredDuplicateOrUnrelatedEvidenceToken(values ...string) bool {
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

func deploymentMultiTenantValDClosureBlockerLevels() []string {
	return []string{
		DeploymentMultiTenantValDBlockerLevelCLB0,
		DeploymentMultiTenantValDBlockerLevelCLB1,
		DeploymentMultiTenantValDBlockerLevelCLB2,
		DeploymentMultiTenantValDBlockerLevelCLB3,
	}
}

func deploymentMultiTenantValDClosureBlockerSurfaces() []string {
	return []string{
		DeploymentMultiTenantValDClosureSurfaceConnectorSandbox,
		DeploymentMultiTenantValDClosureSurfaceOperatorAction,
		DeploymentMultiTenantValDClosureSurfaceSupportAccess,
		DeploymentMultiTenantValDClosureSurfaceBreakGlass,
		DeploymentMultiTenantValDClosureSurfaceMarketplaceMSP,
		DeploymentMultiTenantValDClosureSurfaceAgenticOverlay,
		DeploymentMultiTenantValDClosureSurfaceNoOverclaim,
		DeploymentMultiTenantValDClosureSurfaceCleanRoomIP,
	}
}

func deploymentMultiTenantValDApprovalStatusIsApproved(value string) bool {
	return value == deploymentMultiTenantValDApprovalStatusApproved
}

func deploymentMultiTenantValDExactEvidenceValue(value, expected string) bool {
	return deploymentMultiTenantValDEvidenceValueIsValid(value) && value == expected
}

func deploymentMultiTenantValDSourceOfTruthIsBounded(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "advisory_evidence_input")
}

func deploymentMultiTenantValDPostActionReviewStateValid(value string) bool {
	return value == deploymentMultiTenantValDBreakGlassReviewActive
}

func deploymentMultiTenantValDLearningModeIsSandboxOnly(value string) bool {
	return value == "offline_sandbox_only"
}

func deploymentMultiTenantValDTrainingApprovalStatusValid(value string) bool {
	return value == "training_approved_active"
}

func deploymentMultiTenantValDPromotionApprovalStatusValid(value string, promotionAllowed bool) bool {
	if promotionAllowed {
		return value == "promotion_approved_active"
	}
	return value == "promotion_not_approved"
}

func deploymentMultiTenantValDRuntimeActivationApprovalStatusValid(value string, runtimeActivationAllowed bool) bool {
	if runtimeActivationAllowed {
		return value == "runtime_activation_approved_active"
	}
	return value == "runtime_activation_not_approved"
}

func deploymentMultiTenantValDRecommendationApprovalStatusValid(value string) bool {
	return value == "recommendation_reviewed"
}

func deploymentMultiTenantValDExecutionApprovalStatusValid(value string) bool {
	return value == "execution_not_approved"
}

func deploymentMultiTenantValDModelUpgradeApprovalStatusValid(value string) bool {
	return value == "model_upgrade_not_approved"
}

func deploymentMultiTenantValDConnectorPermissionManifestValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "connector_permission_manifest")
}

func deploymentMultiTenantValDConnectorAuditIDValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "connector_audit_id")
}

func deploymentMultiTenantValDConnectorReasonValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "bounded_connector_reason")
}

func deploymentMultiTenantValDOperatorReasonValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "bounded_operator_reason")
}

func deploymentMultiTenantValDHumanApprovalAuthorityValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "human_approval_authority")
}

func deploymentMultiTenantValDOperatorAuditIDValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "operator_action_audit_id")
}

func deploymentMultiTenantValDSupportReasonValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "support_reason")
}

func deploymentMultiTenantValDSupportAuditIDValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "support_access_audit_id")
}

func deploymentMultiTenantValDBreakGlassEmergencyReasonValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "emergency_reason")
}

func deploymentMultiTenantValDBreakGlassAuditIDValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "break_glass_audit_id")
}

func deploymentMultiTenantValDMarketplaceReasonValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "bounded_marketplace_reason")
}

func deploymentMultiTenantValDMarketplaceAuditIDValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "marketplace_msp_audit_id")
}

func deploymentMultiTenantValDMarketplaceCustomerReadyValidationEvidenceValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "customer_ready_validation_evidence")
}

func deploymentMultiTenantValDExpectedAgentPermissionManifest(agentType string) string {
	switch agentType {
	case "runtime_approval_controller":
		return "agent_permission_manifest"
	case "tenant_boundary_containment_agent":
		return "containment_permission_manifest"
	case "deployment_health_preflight_agent":
		return "deployment_agent_permission_manifest"
	case "connector_operator_misuse_watch_agent":
		return "misuse_watch_permission_manifest"
	case "recovery_rebuild_recommendation_agent":
		return "recovery_agent_permission_manifest"
	default:
		return ""
	}
}

func deploymentMultiTenantValDAgentPermissionManifestValid(agentType, value string) bool {
	expected := deploymentMultiTenantValDExpectedAgentPermissionManifest(agentType)
	return expected != "" && deploymentMultiTenantValDExactEvidenceValue(value, expected)
}

func deploymentMultiTenantValDAgentApprovalReasonValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "human_approved_action_required")
}

func deploymentMultiTenantValDExpectedAgentAuditID(agentType string) string {
	switch agentType {
	case "runtime_approval_controller":
		return "agent_runtime_audit_id"
	case "tenant_boundary_containment_agent":
		return "containment_agent_audit_id"
	case "deployment_health_preflight_agent":
		return "deployment_agent_audit_id"
	case "connector_operator_misuse_watch_agent":
		return "misuse_watch_audit_id"
	case "recovery_rebuild_recommendation_agent":
		return "recovery_agent_audit_id"
	default:
		return ""
	}
}

func deploymentMultiTenantValDAgentAuditIDValid(agentType, value string) bool {
	expected := deploymentMultiTenantValDExpectedAgentAuditID(agentType)
	return expected != "" && deploymentMultiTenantValDExactEvidenceValue(value, expected)
}

func deploymentMultiTenantValDLearningLoopTrainingApproverValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "human_training_approver")
}

func deploymentMultiTenantValDLearningLoopModelCandidateIDValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "candidate_model_id")
}

func deploymentMultiTenantValDLearningLoopModelVersionValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "candidate_model_version")
}

func deploymentMultiTenantValDLearningLoopBaselineModelVersionValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "baseline_model_version")
}

func deploymentMultiTenantValDLearningLoopAuditIDValid(value string) bool {
	return deploymentMultiTenantValDExactEvidenceValue(value, "agent_learning_loop_audit_id")
}

func deploymentMultiTenantValDExactTenantScopedValue(value, suffix string) bool {
	return value == deploymentMultiTenantVal0TenantScope()+" "+suffix
}

func deploymentMultiTenantValDLearningLoopTrainingDataScopeValid(value string) bool {
	return deploymentMultiTenantValDExactTenantScopedValue(value, "training_data_scope")
}

func deploymentMultiTenantValDOperatorActionScopeValid(value string) bool {
	return deploymentMultiTenantValDExactTenantScopedValue(value, "operator_action_scope")
}

func deploymentMultiTenantValDSupportAccessScopeValid(value string) bool {
	return deploymentMultiTenantValDExactTenantScopedValue(value, "support_scope")
}

func deploymentMultiTenantValDBreakGlassActionScopeValid(value string) bool {
	return deploymentMultiTenantValDExactTenantScopedValue(value, "break_glass_scope")
}

func deploymentMultiTenantValDMarketplaceMSPOperatorScopeValid(value string) bool {
	return deploymentMultiTenantValDExactTenantScopedValue(value, "msp_operator_scope")
}

func deploymentMultiTenantValDMarketplacePartnerScopeValid(value string) bool {
	return deploymentMultiTenantValDExactTenantScopedValue(value, "partner_scope")
}

func deploymentMultiTenantValDAgentCoreFieldsValid(agent DeploymentMultiTenantValDAgentRecommendation) bool {
	return deploymentMultiTenantValDEvidenceValueIsValid(agent.AgentID) &&
		deploymentMultiTenantValDEvidenceValueIsValid(agent.AgentType) &&
		deploymentMultiTenantVal0ExactTenantScopeValueIsValid(agent.TenantScope) &&
		deploymentMultiTenantValDAgentPermissionManifestValid(agent.AgentType, agent.PermissionManifest) &&
		deploymentMultiTenantValDAllExactValuesValid(agent.AllowedReadSurfaces) &&
		deploymentMultiTenantValDAllExactValuesValid(agent.AllowedRecommendationSurfaces) &&
		deploymentMultiTenantValDAllExactValuesValid(agent.ForbiddenMutationSurfaces) &&
		agent.ApprovalRequired &&
		deploymentMultiTenantValDApprovalStatusIsApproved(agent.ApprovalStatus) &&
		deploymentMultiTenantValDHumanApprovalAuthorityValid(agent.Approver) &&
		deploymentMultiTenantValDAgentApprovalReasonValid(agent.ApprovalReason) &&
		deploymentMultiTenantValDAgentAuditIDValid(agent.AgentType, agent.AuditID) &&
		deploymentMultiTenantValDHasExactEvidenceRefs(agent.EvidenceRefs, deploymentMultiTenantValDAgenticOverlayEvidenceRefs()...) &&
		deploymentMultiTenantValDEvidenceValueIsValid(agent.RecommendationID) &&
		deploymentMultiTenantValDEvidenceValueIsValid(agent.RecommendationType) &&
		agent.HumanReviewRequired &&
		!agent.CanonicalMutationAllowed &&
		!agent.ProductionMutationAllowed &&
		!agent.Point10PassAllowed &&
		!agent.ExternalAPIAllowed &&
		agent.DiagnosticOutputComplete &&
		deploymentMultiTenantValDHasFoundationProjectionDisclaimer(agent.ProjectionDisclaimer)
}

func EvaluateDeploymentMultiTenantValDAgentLearningLoopState(model DeploymentMultiTenantValDAgentLearningLoop) string {
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!model.LearningAllowed ||
		!deploymentMultiTenantValDLearningModeIsSandboxOnly(model.LearningMode) ||
		!deploymentMultiTenantValDLearningLoopTrainingDataScopeValid(model.TrainingDataScope) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.TrainingDataRefs, deploymentMultiTenantValDAgentLearningLoopEvidenceRefs()...) ||
		!model.TrainingDataPrivacyFiltered ||
		!model.TrainingDataTenantScoped ||
		model.TrainingDataCrossTenant ||
		!model.TrainingDataCustomerApproved ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.HumanFeedbackRefs, deploymentMultiTenantValDAgentLearningLoopHumanFeedbackRefs()...) ||
		!model.HumanFeedbackAuditLinked ||
		!model.TrainingApprovalRequired ||
		!deploymentMultiTenantValDTrainingApprovalStatusValid(model.TrainingApprovalStatus) ||
		!deploymentMultiTenantValDLearningLoopTrainingApproverValid(model.TrainingApprover) ||
		!deploymentMultiTenantValDLearningLoopModelCandidateIDValid(model.ModelCandidateID) ||
		!deploymentMultiTenantValDLearningLoopModelVersionValid(model.ModelVersion) ||
		!deploymentMultiTenantValDLearningLoopBaselineModelVersionValid(model.BaselineModelVersion) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.EvaluationResultRefs, deploymentMultiTenantValDAgentLearningLoopEvaluationResultRefs()...) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.RegressionTestRefs, deploymentMultiTenantValDAgentLearningLoopRegressionTestRefs()...) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.NoOverclaimCheckRefs, deploymentMultiTenantValDAgentLearningLoopNoOverclaimCheckRefs()...) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.TenantScopeCheckRefs, deploymentMultiTenantValDAgentLearningLoopTenantScopeCheckRefs()...) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.ApprovalGateCheckRefs, deploymentMultiTenantValDAgentLearningLoopApprovalGateCheckRefs()...) ||
		!model.LearnedOutputAdvisoryOnly ||
		model.ExternalAPIAllowed ||
		model.ProductionSelfModificationAllowed ||
		model.ProductionMutationAllowed ||
		model.CanonicalMutationAllowed ||
		model.Point10PassAllowed ||
		!model.ApprovalRequired ||
		!deploymentMultiTenantValDLearningLoopAuditIDValid(model.AuditID) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.EvidenceRefs, deploymentMultiTenantValDAgentLearningLoopEvidenceRefs()...) ||
		!model.DiagnosticOutputComplete ||
		model.CandidateWeakensNoOverclaim ||
		model.CandidateWeakensTenantIsolation ||
		model.CandidateWeakensApprovalGates ||
		model.CandidateEnablesExternalAPIByDefault ||
		model.CandidateEnablesProductionMutation ||
		model.CandidateEnablesCanonicalMutation ||
		model.CandidateEnablesPoint10Pass ||
		model.AgentSelfPromotes ||
		model.AgentSelfDeploys ||
		model.AgentModifiesProductionPolicy ||
		model.RecommendationApprovalMeansExecutionApproval ||
		model.ExecutionApprovalMeansModelUpgradeApproval ||
		model.LearnedOutputTreatedAsCanonicalTruth {
		return DeploymentMultiTenantValDAgentLearningLoopStateBlocked
	}
	if !deploymentMultiTenantValDPromotionApprovalStatusValid(model.PromotionApprovalStatus, model.PromotionAllowed) {
		return DeploymentMultiTenantValDAgentLearningLoopStateBlocked
	}
	if !deploymentMultiTenantValDRuntimeActivationApprovalStatusValid(model.RuntimeActivationApprovalStatus, model.RuntimeActivationAllowed) {
		return DeploymentMultiTenantValDAgentLearningLoopStateBlocked
	}
	if !deploymentMultiTenantValDRecommendationApprovalStatusValid(model.RecommendationApprovalStatus) ||
		!deploymentMultiTenantValDExecutionApprovalStatusValid(model.ExecutionApprovalStatus) ||
		!deploymentMultiTenantValDModelUpgradeApprovalStatusValid(model.ModelUpgradeApprovalStatus) {
		return DeploymentMultiTenantValDAgentLearningLoopStateBlocked
	}
	return DeploymentMultiTenantValDAgentLearningLoopStateActive
}

func deploymentMultiTenantValDDependencySnapshotFromValC(valC DeploymentMultiTenantValCFoundation) DeploymentMultiTenantValDDependencySnapshot {
	return DeploymentMultiTenantValDDependencySnapshot{
		ValCCurrentState:           valC.CurrentState,
		ValCDependencyState:        valC.DependencyState,
		ValCHAReadinessState:       valC.HAReadinessState,
		ValCRecoveryReadinessState: valC.RecoveryReadinessState,
		ValCSLAReadinessState:      valC.SLAReadinessState,
		ValCTenantTrustScopeState:  valC.TenantTrustScopeState,
		ValCSiloVisibilityState:    valC.SiloVisibilityState,
		ValCPrivacyGuardState:      valC.PrivacyGuardState,
		ValCNoOverclaimState:       valC.NoOverclaimState,
		ValCClosureBlockerState:    valC.ClosureBlockerState,
		Point10State:               valC.Point10State,
		ProjectionDisclaimer:       valC.ProjectionDisclaimer,
	}
}

func deploymentMultiTenantValDDependencySnapshotModel() DeploymentMultiTenantValDDependencySnapshot {
	valC := ComputeDeploymentMultiTenantValCFoundation(DeploymentMultiTenantValCFoundationModel())
	return deploymentMultiTenantValDDependencySnapshotFromValC(valC)
}

func EvaluateDeploymentMultiTenantValDDependencyState(model DeploymentMultiTenantValDDependencySnapshot) string {
	if !deploymentMultiTenantValCHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeploymentMultiTenantValDDependencyStateBlocked
	}
	if model.ValCCurrentState != DeploymentMultiTenantValCStateActive ||
		model.ValCDependencyState != DeploymentMultiTenantValCDependencyStateActive ||
		model.ValCHAReadinessState != DeploymentMultiTenantValCHAReadinessStateActive ||
		model.ValCRecoveryReadinessState != DeploymentMultiTenantValCRecoveryReadinessStateActive ||
		model.ValCSLAReadinessState != DeploymentMultiTenantValCSLAReadinessStateActive ||
		model.ValCTenantTrustScopeState != DeploymentMultiTenantValCTenantTrustScopeStateActive ||
		model.ValCSiloVisibilityState != DeploymentMultiTenantValCSiloVisibilityStateActive ||
		model.ValCPrivacyGuardState != DeploymentMultiTenantValCPrivacyGuardStateActive ||
		model.ValCNoOverclaimState != DeploymentMultiTenantValCNoOverclaimStateActive ||
		model.ValCClosureBlockerState != DeploymentMultiTenantValCClosureBlockerStateActive ||
		model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		return DeploymentMultiTenantValDDependencyStateBlocked
	}
	return DeploymentMultiTenantValDDependencyStateActive
}

func EvaluateDeploymentMultiTenantValDConnectorCapabilityState(model DeploymentMultiTenantValDConnectorCapabilityManifest) string {
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!deploymentMultiTenantVal0FreshnessIsFresh(model.FreshnessState) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.ConnectorID) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.ConnectorType) ||
		!deploymentMultiTenantVal0ExactTenantScopeValueIsValid(model.TenantScope) ||
		!deploymentMultiTenantValDConnectorPermissionManifestValid(model.PermissionManifest) ||
		!model.CapabilityManifestPresent ||
		!deploymentMultiTenantValDAllExactValuesValid(model.ReadCapabilities) ||
		!deploymentMultiTenantValDAllExactValuesValid(model.WriteCapabilities) ||
		!deploymentMultiTenantValDAllExactValuesValid(model.EvidenceTypes) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.RetryPolicy) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.ReplayPolicy) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.RateLimitPolicy) ||
		!model.AuditRequired ||
		!deploymentMultiTenantValDConnectorAuditIDValid(model.AuditID) ||
		!deploymentMultiTenantValDSourceOfTruthIsBounded(model.SourceOfTruth) ||
		!model.TenantScopedExecution ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.RecoveryBehavior) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.FailureBehavior) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.EvidenceRefs, deploymentMultiTenantValDConnectorEvidenceRefs()...) ||
		!model.DiagnosticOutputComplete ||
		model.ConnectorAsSourceOfTruth ||
		model.ConnectorBypassesDeploymentGate ||
		model.ConnectorBypassesTenantGate ||
		model.ConnectorBypassesEvidenceGate ||
		model.ConnectorBypassesDataResidencyGate ||
		model.RetryReplayDuplicatesActiveEvidenceRisk {
		return DeploymentMultiTenantValDConnectorCapabilityStateBlocked
	}
	if model.MutationAllowed && (!model.MutationCapabilityExplicit || !deploymentMultiTenantValDConnectorReasonValid(model.Reason) || !deploymentMultiTenantValDConnectorAuditIDValid(model.AuditID)) {
		return DeploymentMultiTenantValDConnectorCapabilityStateBlocked
	}
	return DeploymentMultiTenantValDConnectorCapabilityStateActive
}

func EvaluateDeploymentMultiTenantValDOperatorActionState(model DeploymentMultiTenantValDOperatorActionPolicy) string {
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.Actor) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.ActorType) ||
		!deploymentMultiTenantVal0ExactTenantScopeValueIsValid(model.TenantTarget) ||
		!deploymentMultiTenantValDOperatorActionScopeValid(model.ActionScope) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.ActionType) ||
		!deploymentMultiTenantValDOperatorReasonValid(model.Reason) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.AuthorizationBasis) ||
		!deploymentMultiTenantValDHumanApprovalAuthorityValid(model.Approver) ||
		!deploymentMultiTenantValDApprovalStatusIsApproved(model.ApprovalStatus) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.AuthorityBasis) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.Expiry) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.RevocationPath) ||
		!deploymentMultiTenantValDOperatorAuditIDValid(model.AuditID) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.EvidenceRefs, deploymentMultiTenantValDOperatorEvidenceRefs()...) ||
		!model.RBACABACEnforced ||
		!model.SSOContextBound ||
		!model.TenantScopeBound ||
		!model.SupportScopeBound ||
		model.ProductionMutationAllowed ||
		model.CanonicalMutationAllowed ||
		model.OperatorActionCanonicalApproval ||
		!model.DiagnosticOutputComplete ||
		!deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeploymentMultiTenantValDOperatorActionStateBlocked
	}
	return DeploymentMultiTenantValDOperatorActionStateActive
}

func EvaluateDeploymentMultiTenantValDSupportAccessState(model DeploymentMultiTenantValDSupportAccessEnforcement) string {
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.SupportAccessID) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.SupportActor) ||
		!deploymentMultiTenantVal0ExactTenantScopeValueIsValid(model.TenantTarget) ||
		!deploymentMultiTenantValDSupportAccessScopeValid(model.SupportScope) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.SSOSessionReference) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.RBACRole) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.ABACConditions) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.AuthorityBasis) ||
		!deploymentMultiTenantValDSupportReasonValid(model.Reason) ||
		!deploymentMultiTenantValDHumanApprovalAuthorityValid(model.Approver) ||
		!deploymentMultiTenantValDApprovalStatusIsApproved(model.ApprovalStatus) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.Expiry) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.RevocationPath) ||
		!deploymentMultiTenantValDSupportAuditIDValid(model.AuditID) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.EvidenceRefs, deploymentMultiTenantValDSupportEvidenceRefs()...) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.SupportVisibilityBoundary) ||
		!model.DataResidencyBoundaryRespected ||
		!model.TenantIsolationBoundaryRespected ||
		!model.DiagnosticOutputComplete ||
		model.SupportVisibilityExceedsScope ||
		model.RawTenantEvidenceExposed ||
		model.SupportSummaryCanonicalTruth {
		return DeploymentMultiTenantValDSupportAccessStateBlocked
	}
	return DeploymentMultiTenantValDSupportAccessStateActive
}

func EvaluateDeploymentMultiTenantValDBreakGlassState(model DeploymentMultiTenantValDBreakGlassAccess) string {
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.BreakGlassID) ||
		!deploymentMultiTenantValDBreakGlassEmergencyReasonValid(model.EmergencyReason) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.Actor) ||
		!deploymentMultiTenantVal0ExactTenantScopeValueIsValid(model.TenantTarget) ||
		!deploymentMultiTenantValDBreakGlassActionScopeValid(model.ActionScope) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.AuthorizationBasis) ||
		!deploymentMultiTenantValDHumanApprovalAuthorityValid(model.Approver) ||
		!deploymentMultiTenantValDApprovalStatusIsApproved(model.ApprovalStatus) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.Expiry) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.RevocationPath) ||
		!deploymentMultiTenantValDBreakGlassAuditIDValid(model.AuditID) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.EvidenceRefs, deploymentMultiTenantValDBreakGlassEvidenceRefs()...) ||
		!model.PostActionReviewRequired ||
		!deploymentMultiTenantValDPostActionReviewStateValid(model.PostActionReviewState) ||
		!model.TenantScopeBound ||
		!model.DataResidencyBoundaryRespected ||
		!model.TenantIsolationBoundaryRespected ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.SupportVisibilityBoundary) ||
		!model.DiagnosticOutputComplete ||
		model.PersistentAccessGranted ||
		model.CreatesPASSAuthority {
		return DeploymentMultiTenantValDBreakGlassStateBlocked
	}
	return DeploymentMultiTenantValDBreakGlassStateActive
}

func EvaluateDeploymentMultiTenantValDMarketplaceMSPAuthorityState(model DeploymentMultiTenantValDMarketplaceMSPAuthorityBoundary) string {
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.MarketplaceProfile) ||
		!deploymentMultiTenantValDMarketplaceMSPOperatorScopeValid(model.MSPOperatorScope) ||
		!deploymentMultiTenantValDMarketplacePartnerScopeValid(model.PartnerScope) ||
		!deploymentMultiTenantVal0ExactTenantScopeValueIsValid(model.TenantScope) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.DeploymentScope) ||
		!deploymentMultiTenantVal0BoundaryValueIsValid(model.SupportScope) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.AuthorityBoundary) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.ApprovalBoundary) ||
		!deploymentMultiTenantValDMarketplaceAuditIDValid(model.AuditID) ||
		!deploymentMultiTenantValDMarketplaceReasonValid(model.Reason) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.Expiry) ||
		!deploymentMultiTenantValDEvidenceValueIsValid(model.RevocationPath) ||
		!deploymentMultiTenantValDHasExactEvidenceRefs(model.EvidenceRefs, deploymentMultiTenantValDMarketplaceMSPEvidenceRefs()...) ||
		!model.DiagnosticOutputComplete ||
		model.PassAuthorityAllowed ||
		model.ProductionReadinessAuthorityAllowed ||
		model.SourceOfTruthAllowed ||
		model.MarketplaceInstallTreatedAsProdApproved ||
		model.MSPApprovedDeploymentClaim ||
		model.PartnerCertifiedDeploymentClaim {
		return DeploymentMultiTenantValDMarketplaceMSPAuthorityStateBlocked
	}
	if model.CustomerReadyWordingPresent && !deploymentMultiTenantValDMarketplaceCustomerReadyValidationEvidenceValid(model.CustomerReadyValidationEvidence) {
		return DeploymentMultiTenantValDMarketplaceMSPAuthorityStateBlocked
	}
	return DeploymentMultiTenantValDMarketplaceMSPAuthorityStateActive
}

func deploymentMultiTenantValDAgentRecommendationRequiresApprovalQueue(agentType string) bool {
	return strings.TrimSpace(agentType) == "runtime_approval_controller"
}

func deploymentMultiTenantValDAgentRecommendationRequiresRecoveryEvidencePack(agentType string) bool {
	return strings.Contains(strings.TrimSpace(agentType), "recovery")
}

func deploymentMultiTenantValDAgentRecommendationState(agent DeploymentMultiTenantValDAgentRecommendation) string {
	if !deploymentMultiTenantValDAgentCoreFieldsValid(agent) {
		return DeploymentMultiTenantValDAgenticOverlayStateBlocked
	}
	if deploymentMultiTenantValDAgentRecommendationRequiresApprovalQueue(agent.AgentType) && !deploymentMultiTenantValDEvidenceValueIsValid(agent.ApprovalQueue) {
		return DeploymentMultiTenantValDAgenticOverlayStateBlocked
	}
	if deploymentMultiTenantValDAgentRecommendationRequiresRecoveryEvidencePack(agent.AgentType) && !deploymentMultiTenantValDEvidenceValueIsValid(agent.RecoveryEvidencePack) {
		return DeploymentMultiTenantValDAgenticOverlayStateBlocked
	}
	if agent.AgentTreatedAsSourceOfTruth ||
		agent.CrossTenantAccess ||
		agent.CanonicalMutationAllowed ||
		agent.ProductionMutationAllowed ||
		agent.Point10PassAllowed ||
		agent.ExternalAPIAllowed ||
		agent.TreatsRecommendationAsApproval {
		return DeploymentMultiTenantValDAgenticOverlayStateBlocked
	}
	if agent.ContainmentExecutedWithoutApproval ||
		agent.ChangesTenantStateDirectly ||
		agent.ExpandsOrLimitsAccessDirectlyWithoutApproval ||
		agent.InstallSuccessTreatedAsReadiness ||
		agent.PreflightExecutedWithoutApproval ||
		agent.MarksDeploymentReadyWithoutCanonicalEvaluator ||
		agent.SLAReadinessTreatedAsUptimeGuarantee ||
		agent.ConnectorMutationExecutedByAgent ||
		agent.ConnectorCapabilityMissing ||
		agent.OperatorSupportActionWithoutAuthorityBasis ||
		agent.BreakGlassExpiryRevocationMissing ||
		agent.RestoreRollbackRebuildExecutedAutomatically ||
		agent.RecommendationBypassesTenantIsolation ||
		agent.RecommendationBypassesDataResidency ||
		agent.RecoveryGuaranteedClaim {
		return DeploymentMultiTenantValDAgenticOverlayStateBlocked
	}
	return DeploymentMultiTenantValDAgenticOverlayStateActive
}

func EvaluateDeploymentMultiTenantValDAgenticOverlayState(model DeploymentMultiTenantValDAgenticOverlay) string {
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeploymentMultiTenantValDAgenticOverlayStateBlocked
	}
	if EvaluateDeploymentMultiTenantValDAgentLearningLoopState(model.LearningLoop) != DeploymentMultiTenantValDAgentLearningLoopStateActive {
		return DeploymentMultiTenantValDAgenticOverlayStateBlocked
	}
	agents := []DeploymentMultiTenantValDAgentRecommendation{
		DeploymentMultiTenantValDAgentRecommendation(model.RuntimeApprovalController),
		model.TenantBoundaryContainmentAgent,
		model.DeploymentHealthPreflightAgent,
		model.ConnectorOperatorMisuseWatchAgent,
		model.RecoveryRebuildRecommendationAgent,
	}
	for _, agent := range agents {
		if deploymentMultiTenantValDAgentRecommendationState(agent) != DeploymentMultiTenantValDAgenticOverlayStateActive {
			return DeploymentMultiTenantValDAgenticOverlayStateBlocked
		}
	}
	return DeploymentMultiTenantValDAgenticOverlayStateActive
}

func deploymentMultiTenantValDContainsForbiddenClaim(values ...string) bool {
	allowedExact := []string{
		"sandboxed connector execution",
		"connector capability manifest",
		"explicit connector capability",
		"connector misuse signal",
		"operator misuse signal",
		"bounded operator authority",
		"zero-trust operator action",
		"tenant-scoped support access",
		"break-glass approval required",
		"break-glass expiry and revocation evidence",
		"msp support surface",
		"marketplace deployment profile",
		"advisory recommendation",
		"human-approved action required",
		"evidence-linked recommendation",
		"tenant-scoped agent runtime",
		"approval-gated recovery recommendation",
		"offline sandbox learning pipeline",
		"human-approved model promotion",
		"candidate model version",
		"advisory learning improvement",
		"audit-linked human feedback",
		"regression-tested candidate",
		"no-overclaim checked candidate",
		"tenant-scope checked candidate",
		"approval-gated runtime activation",
		"learned output remains advisory",
		"no production autopatch",
		"no auto-merge",
		"no auto-deploy",
		"not canonical truth",
		"not production approval",
		"not deployment approval",
		"not compliance certification",
	}
	disallowed := []string{
		"connector approved deployment",
		"connector is source of truth",
		"connector certified evidence",
		"connector mutation safe by default",
		"connector can mutate without approval",
		"operator fully trusted",
		"operator approved deployment",
		"support access cannot leak",
		"break-glass safe by default",
		"break-glass permanent access",
		"msp approved deployment",
		"msp certified deployment",
		"partner certified deployment",
		"marketplace production ready",
		"marketplace certified",
		"customer ready without validation",
		"autonomous remediation approved",
		"agent approved deployment",
		"agent certified recovery",
		"ai certified fix",
		"auto-merge safe",
		"auto-deploy safe",
		"production autopatch",
		"recovery guaranteed",
		"agent guaranteed tenant isolation",
		"agent proves compliance",
		"point 10 pass by agent",
		"agent is source of truth",
		"external ai verified",
		"external api verified",
		"no human approval required",
		"production mutation approved by agent",
		"canonical evidence mutation by agent",
		"self-learning agent approved deployment",
		"agent self-approved model",
		"autonomous model promotion",
		"autonomous runtime activation",
		"agent trained itself into production",
		"online self-modification safe",
		"model upgrade automatically approved",
		"recommendation approval means execution approval",
		"execution approval means model upgrade approval",
		"learned output is canonical truth",
		"learned model certified",
		"ai model certified",
		"agent learning guarantees security",
		"self-improving agent guarantees compliance",
		"point 10 pass by learned model",
		"external ai verified model",
		"clean-room certified",
		"patent cleared",
		"fto cleared",
		"legal certification",
		"copied competitor workflow",
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
		isAllowedExact := false
		for _, allowed := range allowedExact {
			if normalized == deploymentMultiTenantVal0NormalizeClaimText(allowed) {
				isAllowedExact = true
				break
			}
		}
		if isAllowedExact {
			continue
		}
		if normalized != "" {
			corpusNormalizedParts = append(corpusNormalizedParts, normalized)
		}
		corpusCompact.WriteString(compact)
		for _, forbidden := range disallowed {
			blockedNormalized := deploymentMultiTenantVal0NormalizeClaimText(forbidden)
			blockedCompact := deploymentMultiTenantVal0CompactClaimText(forbidden)
			if strings.Contains(normalized, blockedNormalized) ||
				strings.Contains(compact, blockedCompact) ||
				deploymentMultiTenantVal0ValueContainsForbiddenPhraseTokenSequence(normalized, blockedNormalized) {
				return true
			}
		}
	}
	corpusNormalized := strings.Join(corpusNormalizedParts, " ")
	corpusCompactValue := corpusCompact.String()
	for _, forbidden := range disallowed {
		blockedNormalized := deploymentMultiTenantVal0NormalizeClaimText(forbidden)
		blockedCompact := deploymentMultiTenantVal0CompactClaimText(forbidden)
		if strings.Contains(corpusNormalized, blockedNormalized) || strings.Contains(corpusCompactValue, blockedCompact) {
			return true
		}
		if deploymentMultiTenantVal0BucketsContainForbiddenPhraseAcrossValues(crossNormalizedParts, blockedNormalized) {
			return true
		}
	}
	return false
}

func EvaluateDeploymentMultiTenantValDNoOverclaimState(model DeploymentMultiTenantValDNoOverclaimDiscipline) string {
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		model.CleanRoomIPViolationDetected ||
		deploymentMultiTenantValDContainsForbiddenClaim(model.ObservedClaims...) {
		return DeploymentMultiTenantValDNoOverclaimStateBlocked
	}
	return DeploymentMultiTenantValDNoOverclaimStateActive
}

func deploymentMultiTenantValDClosureBlockerFinding(level, surface, reason string, blocks bool, followup string) DeploymentMultiTenantValDClosureBlockerFinding {
	return DeploymentMultiTenantValDClosureBlockerFinding{
		BlockerLevel:      level,
		Surface:           surface,
		Reason:            reason,
		BlocksCurrentWave: blocks,
		RequiredFollowup:  followup,
	}
}

func deploymentMultiTenantValDClosureBlockerFindings(model DeploymentMultiTenantValDFoundation) []DeploymentMultiTenantValDClosureBlockerFinding {
	findings := []DeploymentMultiTenantValDClosureBlockerFinding{}
	if model.ConnectorCapability.MutationAllowed && !model.ConnectorCapability.MutationCapabilityExplicit {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, "connector mutation without explicit capability", true, "remove connector mutation path or require explicit tenant-scoped capability with audit"))
	}
	if model.ConnectorCapability.ConnectorAsSourceOfTruth {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, "connector treated as source of truth", true, "restore advisory-only connector semantics and preserve canonical evaluators"))
	}
	if model.ConnectorCapability.ConnectorBypassesTenantGate || model.ConnectorCapability.ConnectorBypassesEvidenceGate || model.ConnectorCapability.ConnectorBypassesDeploymentGate || model.ConnectorCapability.ConnectorBypassesDataResidencyGate {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, "connector bypasses tenant evidence deployment or data residency gate", true, "remove connector bypass path and rerun bounded connector validation"))
	}
	if !deploymentMultiTenantValDEvidenceValueIsValid(model.OperatorAction.AuthorityBasis) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceOperatorAction, "operator or support action without authority basis", true, "add explicit authority basis before any operator or support workflow proceeds"))
	}
	if !model.OperatorAction.RBACABACEnforced || !model.OperatorAction.SSOContextBound || !model.SupportAccess.DataResidencyBoundaryRespected || !model.SupportAccess.TenantIsolationBoundaryRespected {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceSupportAccess, "sso or rbac abac bypass", true, "restore sso and rbac abac enforcement before final review"))
	}
	if !deploymentMultiTenantValDEvidenceValueIsValid(model.BreakGlass.AuthorizationBasis) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceBreakGlass, "break-glass without authority basis", true, "add explicit break-glass authority basis and rerun break-glass validation"))
	}
	if model.BreakGlass.PersistentAccessGranted || !model.BreakGlass.TenantScopeBound {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceBreakGlass, "break-glass creates persistent or global access", true, "remove persistent or unscoped break-glass access and restore expiry-bounded tenant scope"))
	}
	if model.MarketplaceMSPAuthority.PassAuthorityAllowed || model.MarketplaceMSPAuthority.MSPApprovedDeploymentClaim || model.MarketplaceMSPAuthority.PartnerCertifiedDeploymentClaim || model.MarketplaceMSPAuthority.MarketplaceInstallTreatedAsProdApproved {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceMarketplaceMSP, "marketplace or msp overclaim", true, "remove marketplace or msp approval overclaim and keep authority boundaries bounded"))
	}
	if model.MarketplaceMSPAuthority.PassAuthorityAllowed {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceMarketplaceMSP, "msp or partner allowed to approve pass", true, "remove pass authority from msp or partner surfaces"))
	}
	if model.MarketplaceMSPAuthority.ProductionReadinessAuthorityAllowed {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceMarketplaceMSP, "msp or partner allowed to approve production readiness", true, "remove production readiness authority from msp or partner surfaces"))
	}
	if model.MarketplaceMSPAuthority.SourceOfTruthAllowed {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceMarketplaceMSP, "msp or partner becomes source of truth", true, "restore non-canonical marketplace and msp boundaries"))
	}
	agents := []DeploymentMultiTenantValDAgentRecommendation{
		DeploymentMultiTenantValDAgentRecommendation(model.AgenticOverlay.RuntimeApprovalController),
		model.AgenticOverlay.TenantBoundaryContainmentAgent,
		model.AgenticOverlay.DeploymentHealthPreflightAgent,
		model.AgenticOverlay.ConnectorOperatorMisuseWatchAgent,
		model.AgenticOverlay.RecoveryRebuildRecommendationAgent,
	}
	for _, agent := range agents {
		if agent.ProductionMutationAllowed {
			findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent executes production mutation", true, "disable production mutation and keep agent recommendations advisory only"))
		}
		if agent.CanonicalMutationAllowed {
			findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent mutates canonical evidence spine", true, "remove canonical mutation ability from agent runtime"))
		}
		if agent.Point10PassAllowed {
			findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent emits or enables point 10 pass", true, "remove any point 10 pass path from agent runtime and approval models"))
		}
		if agent.CrossTenantAccess {
			findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent performs cross-tenant access", true, "restore tenant-scoped agent runtime and remove cross-tenant access"))
		}
		if agent.ExternalAPIAllowed {
			findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent enables external api by default", true, "disable external api access by default and require bounded manual review"))
		}
		if agent.ConnectorMutationExecutedByAgent && agent.ConnectorCapabilityMissing {
			findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent executes connector mutation without capability", true, "remove connector mutation execution and require explicit capability plus human approval"))
		}
		if agent.OperatorSupportActionWithoutAuthorityBasis {
			findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent executes operator or support action without authority basis", true, "restore authority basis enforcement before agent recommendations can proceed"))
		}
		if agent.RestoreRollbackRebuildExecutedAutomatically {
			findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent executes restore rollback or rebuild without approval", true, "remove automatic recovery execution and keep recovery recommendation approval-gated"))
		}
		if agent.TreatsRecommendationAsApproval {
			findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent treats recommendation as approval", true, "restore human approval as the only approval authority"))
		}
	}
	if model.AgenticOverlay.LearningLoop.AgentSelfPromotes {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent self-promotes", true, "disable self-promotion and require human-approved model promotion"))
	}
	if model.AgenticOverlay.LearningLoop.AgentSelfDeploys {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent self-deploys", true, "disable self-deploy and keep runtime activation human-approved"))
	}
	if model.AgenticOverlay.LearningLoop.AgentModifiesProductionPolicy {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "agent modifies production policy", true, "remove production policy self-modification and keep learning offline-only"))
	}
	if !deploymentMultiTenantValDRuntimeActivationApprovalStatusValid(model.AgenticOverlay.LearningLoop.RuntimeActivationApprovalStatus, model.AgenticOverlay.LearningLoop.RuntimeActivationAllowed) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "runtime activation without human approval", true, "require human approval before any runtime activation step"))
	}
	if !deploymentMultiTenantValDPromotionApprovalStatusValid(model.AgenticOverlay.LearningLoop.PromotionApprovalStatus, model.AgenticOverlay.LearningLoop.PromotionAllowed) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "model promotion without human approval", true, "require human-approved model promotion before any candidate advances"))
	}
	if model.AgenticOverlay.LearningLoop.ProductionSelfModificationAllowed {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "production self-modification allowed", true, "remove production self-modification and keep learning sandboxed"))
	}
	if model.AgenticOverlay.LearningLoop.ProductionMutationAllowed || model.AgenticOverlay.LearningLoop.CandidateEnablesProductionMutation {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "candidate enables production mutation", true, "remove production mutation enablement from the learning loop"))
	}
	if model.AgenticOverlay.LearningLoop.CanonicalMutationAllowed || model.AgenticOverlay.LearningLoop.CandidateEnablesCanonicalMutation {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "candidate enables canonical evidence mutation", true, "remove canonical evidence mutation enablement from the learning loop"))
	}
	if model.AgenticOverlay.LearningLoop.Point10PassAllowed || model.AgenticOverlay.LearningLoop.CandidateEnablesPoint10Pass {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "candidate enables point 10 pass", true, "remove any point 10 pass path from the learning loop"))
	}
	if model.AgenticOverlay.LearningLoop.TrainingDataCrossTenant {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "cross-tenant training data used", true, "restore tenant-scoped training data boundaries before learning proceeds"))
	}
	if !model.AgenticOverlay.LearningLoop.TrainingDataCustomerApproved {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "unapproved customer data used for training", true, "require customer approval before any customer data enters the learning loop"))
	}
	if model.AgenticOverlay.LearningLoop.ExternalAPIAllowed || model.AgenticOverlay.LearningLoop.CandidateEnablesExternalAPIByDefault {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "external api enabled by default", true, "disable external api access by default in the learning loop"))
	}
	if model.AgenticOverlay.LearningLoop.LearnedOutputTreatedAsCanonicalTruth {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "learned output treated as canonical truth", true, "restore advisory-only learned output semantics"))
	}
	if model.AgenticOverlay.LearningLoop.RecommendationApprovalMeansExecutionApproval {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "recommendation approval treated as execution approval", true, "separate recommendation approval from execution approval"))
	}
	if model.AgenticOverlay.LearningLoop.ExecutionApprovalMeansModelUpgradeApproval {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "execution approval treated as model upgrade approval", true, "separate execution approval from model upgrade approval"))
	}
	if model.NoOverclaimState != DeploymentMultiTenantValDNoOverclaimStateActive {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceNoOverclaim, "forbidden connector operator marketplace or agent overclaim present", true, "remove forbidden overclaim before final review"))
	}
	if model.NoOverclaim.CleanRoomIPViolationDetected {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDClosureSurfaceCleanRoomIP, "copied competitor connector operator deployment or agent workflow detected", true, "remove copied artifact and replace it with clean-room implementation evidence"))
	}
	if !deploymentMultiTenantValDEvidenceValueIsValid(model.ConnectorCapability.PermissionManifest) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, "permission manifest missing", true, "add explicit connector permission manifest before final review"))
	}
	if !model.ConnectorCapability.CapabilityManifestPresent {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, "connector capability manifest missing", true, "add explicit connector capability manifest and rerun connector validation"))
	}
	if !deploymentMultiTenantValDEvidenceValueIsValid(model.ConnectorCapability.ReplayPolicy) || !deploymentMultiTenantValDEvidenceValueIsValid(model.ConnectorCapability.RetryPolicy) || !deploymentMultiTenantValDEvidenceValueIsValid(model.ConnectorCapability.RateLimitPolicy) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, "connector replay retry or rate-limit semantics missing", true, "add bounded replay retry and rate-limit semantics before final review"))
	}
	if !deploymentMultiTenantValDEvidenceValueIsValid(model.BreakGlass.Expiry) || !deploymentMultiTenantValDEvidenceValueIsValid(model.BreakGlass.RevocationPath) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceBreakGlass, "break-glass expiry or revocation path missing", true, "add bounded break-glass expiry and revocation path before final review"))
	}
	if !deploymentMultiTenantValDEvidenceValueIsValid(model.SupportAccess.Expiry) || !deploymentMultiTenantValDEvidenceValueIsValid(model.SupportAccess.RevocationPath) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceSupportAccess, "support access expiry or revocation missing", true, "add support access expiry and revocation path before final review"))
	}
	if !model.OperatorAction.RBACABACEnforced || !model.SupportAccess.DataResidencyBoundaryRespected || !model.SupportAccess.TenantIsolationBoundaryRespected {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceSupportAccess, "rbac abac enforcement not proven", true, "prove support and operator boundary enforcement before final review"))
	}
	if !deploymentMultiTenantValDEvidenceValueIsValid(model.AgenticOverlay.RuntimeApprovalController.ApprovalQueue) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "approval workflow missing", true, "add approval workflow and bounded approval queue before final review"))
	}
	if !deploymentMultiTenantValDAgentAuditIDValid(model.AgenticOverlay.RuntimeApprovalController.AgentType, model.AgenticOverlay.RuntimeApprovalController.AuditID) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "audit trail missing", true, "add agent audit trail before final review"))
	}
	if !deploymentMultiTenantValDHasExactEvidenceRefs(model.AgenticOverlay.RuntimeApprovalController.EvidenceRefs, deploymentMultiTenantValDAgenticOverlayEvidenceRefs()...) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "evidence refs missing", true, "add exact evidence refs before final review"))
	}
	if !deploymentMultiTenantValDEvidenceValueIsValid(model.AgenticOverlay.RuntimeApprovalController.PermissionManifest) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "permission manifest missing", true, "add agent permission manifest before final review"))
	}
	if model.AgenticOverlay.RuntimeApprovalController.ExternalAPIAllowed {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "no-external-api default not proven", true, "restore no-external-api default and rerun agent overlay validation"))
	}
	if !model.AgenticOverlay.RuntimeApprovalController.HumanReviewRequired {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "human review requirement not enforced", true, "enforce human review requirement before final review"))
	}
	if !deploymentMultiTenantValDEvidenceValueIsValid(model.AgenticOverlay.RecoveryRebuildRecommendationAgent.RecoveryEvidencePack) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "recovery recommendation lacks required evidence pack", true, "attach backup restore and dr evidence pack before recovery recommendation review"))
	}
	if deploymentMultiTenantValDHasRevokedExpiredDuplicateOrUnrelatedEvidenceToken(
		model.ConnectorCapability.PermissionManifest,
		model.OperatorAction.AuthorizationBasis,
		model.SupportAccess.AuthorityBasis,
		model.BreakGlass.AuthorizationBasis,
		model.AgenticOverlay.RuntimeApprovalController.PermissionManifest,
		model.AgenticOverlay.RecoveryRebuildRecommendationAgent.RecoveryEvidencePack,
		model.AgenticOverlay.LearningLoop.ModelCandidateID,
		model.AgenticOverlay.LearningLoop.ModelVersion,
		model.AgenticOverlay.LearningLoop.BaselineModelVersion,
	) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "stale revoked expired duplicate or unrelated evidence handling not proven", true, "reject stale revoked expired duplicate unrelated evidence tokens and rerun Val D fail-closed tests"))
	}
	if !model.AgenticOverlay.LearningLoop.TrainingApprovalRequired || !deploymentMultiTenantValDTrainingApprovalStatusValid(model.AgenticOverlay.LearningLoop.TrainingApprovalStatus) || !deploymentMultiTenantValDLearningLoopTrainingApproverValid(model.AgenticOverlay.LearningLoop.TrainingApprover) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "training approval workflow missing", true, "add human-approved learning workflow before final review"))
	}
	if !model.AgenticOverlay.LearningLoop.HumanFeedbackAuditLinked {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "human feedback audit link missing", true, "link human feedback to audit evidence before final review"))
	}
	if !deploymentMultiTenantValDHasExactEvidenceRefs(model.AgenticOverlay.LearningLoop.EvaluationResultRefs, deploymentMultiTenantValDAgentLearningLoopEvaluationResultRefs()...) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "evaluation result refs missing", true, "add evaluation result refs before final review"))
	}
	if !deploymentMultiTenantValDHasExactEvidenceRefs(model.AgenticOverlay.LearningLoop.RegressionTestRefs, deploymentMultiTenantValDAgentLearningLoopRegressionTestRefs()...) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "regression test refs missing", true, "add regression test refs before final review"))
	}
	if !deploymentMultiTenantValDHasExactEvidenceRefs(model.AgenticOverlay.LearningLoop.NoOverclaimCheckRefs, deploymentMultiTenantValDAgentLearningLoopNoOverclaimCheckRefs()...) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "no-overclaim check refs missing", true, "add no-overclaim check refs before final review"))
	}
	if !deploymentMultiTenantValDHasExactEvidenceRefs(model.AgenticOverlay.LearningLoop.TenantScopeCheckRefs, deploymentMultiTenantValDAgentLearningLoopTenantScopeCheckRefs()...) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "tenant-scope check refs missing", true, "add tenant-scope check refs before final review"))
	}
	if !deploymentMultiTenantValDHasExactEvidenceRefs(model.AgenticOverlay.LearningLoop.ApprovalGateCheckRefs, deploymentMultiTenantValDAgentLearningLoopApprovalGateCheckRefs()...) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "approval-gate check refs missing", true, "add approval-gate check refs before final review"))
	}
	if !deploymentMultiTenantValDLearningLoopModelCandidateIDValid(model.AgenticOverlay.LearningLoop.ModelCandidateID) || !deploymentMultiTenantValDLearningLoopModelVersionValid(model.AgenticOverlay.LearningLoop.ModelVersion) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "model candidate versioning missing", true, "add bounded candidate model versioning before final review"))
	}
	if !deploymentMultiTenantValDLearningLoopBaselineModelVersionValid(model.AgenticOverlay.LearningLoop.BaselineModelVersion) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "baseline model version missing", true, "add bounded baseline model version before final review"))
	}
	if !deploymentMultiTenantValDLearningLoopAuditIDValid(model.AgenticOverlay.LearningLoop.AuditID) {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "learning-loop audit trail missing", true, "add bounded learning-loop audit trail before final review"))
	}
	if model.DependencyState != DeploymentMultiTenantValDDependencyStateActive {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB1, DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, "dependency gate missing or not exact active", true, "restore exact active Val C dependency before Val D final review"))
	}
	if !model.ConnectorCapability.ConnectorNamingExact {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceConnectorSandbox, "ambiguous connector naming", true, "normalize connector naming before handoff"))
	}
	if !model.OperatorAction.OperatorActionNamingExact {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceOperatorAction, "ambiguous operator action naming", true, "normalize operator action naming before handoff"))
	}
	if !model.AgenticOverlay.RuntimeApprovalController.AgentNamingExact {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "ambiguous agent naming", true, "normalize agent naming before handoff"))
	}
	if !model.AgenticOverlay.LearningLoop.LearningModeNamingExact {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "ambiguous learning mode naming", true, "normalize learning mode naming before handoff"))
	}
	if !model.AgenticOverlay.LearningLoop.ModelCandidateNamingExact {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "ambiguous model candidate naming", true, "normalize model candidate naming before handoff"))
	}
	if !model.ConnectorCapability.SafeConnectorWordingExamplePresent || !model.OperatorAction.SafeOperatorWordingExamplePresent || !model.AgenticOverlay.RuntimeApprovalController.SafeWordingExamplePresent {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "missing safe wording example for connector operator or agent recommendations", true, "add bounded safe wording examples before handoff"))
	}
	if !model.AgenticOverlay.LearningLoop.SafeWordingExamplePresent {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "missing safe wording example for advisory learning", true, "add bounded safe wording example for advisory learning before handoff"))
	}
	if !model.ConnectorCapability.DiagnosticOutputComplete || !model.OperatorAction.DiagnosticOutputComplete || !model.SupportAccess.DiagnosticOutputComplete || !model.BreakGlass.DiagnosticOutputComplete || !model.AgenticOverlay.RuntimeApprovalController.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "incomplete diagnostic output for connector operator support break-glass or agent blockers", true, "complete bounded diagnostic output before handoff"))
	}
	if !model.AgenticOverlay.LearningLoop.DiagnosticOutputComplete {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "incomplete diagnostic output for learning-loop blockers", true, "complete learning-loop diagnostic output before handoff"))
	}
	if !model.AgenticOverlay.RuntimeApprovalController.RunbookWordingComplete {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "incomplete bounded runbook wording without direct pass bypass", true, "complete bounded runbook wording before handoff"))
	}
	if !model.AgenticOverlay.LearningLoop.RunbookWordingComplete {
		findings = append(findings, deploymentMultiTenantValDClosureBlockerFinding(DeploymentMultiTenantValDBlockerLevelCLB2, DeploymentMultiTenantValDClosureSurfaceAgenticOverlay, "incomplete runbook wording for model promotion without direct pass bypass", true, "complete bounded model promotion runbook wording before handoff"))
	}
	return findings
}

func EvaluateDeploymentMultiTenantValDClosureBlockerState(model DeploymentMultiTenantValDClosureBlockerOverlay) string {
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeploymentMultiTenantValDClosureBlockerStateBlocked
	}
	hasCleanup := false
	hasAdvisory := false
	for _, finding := range model.Findings {
		level := strings.TrimSpace(finding.BlockerLevel)
		surface := strings.TrimSpace(finding.Surface)
		if len(level) == 2 && level[0] == 'P' && level[1] >= '0' && level[1] <= '9' {
			return DeploymentMultiTenantValDClosureBlockerStateBlocked
		}
		if !containsTrimmedString(deploymentMultiTenantValDClosureBlockerLevels(), level) ||
			!containsTrimmedString(deploymentMultiTenantValDClosureBlockerSurfaces(), surface) {
			return DeploymentMultiTenantValDClosureBlockerStateBlocked
		}
		if (level == DeploymentMultiTenantValDBlockerLevelCLB1 ||
			level == DeploymentMultiTenantValDBlockerLevelCLB2 ||
			level == DeploymentMultiTenantValDBlockerLevelCLB3) &&
			strings.TrimSpace(finding.RequiredFollowup) == "" {
			return DeploymentMultiTenantValDClosureBlockerStateBlocked
		}
		switch level {
		case DeploymentMultiTenantValDBlockerLevelCLB0, DeploymentMultiTenantValDBlockerLevelCLB1:
			return DeploymentMultiTenantValDClosureBlockerStateBlocked
		case DeploymentMultiTenantValDBlockerLevelCLB2:
			hasCleanup = true
		case DeploymentMultiTenantValDBlockerLevelCLB3:
			hasAdvisory = true
		default:
			return DeploymentMultiTenantValDClosureBlockerStateBlocked
		}
	}
	if hasCleanup {
		return DeploymentMultiTenantValDClosureBlockerStateCleanup
	}
	if hasAdvisory {
		return DeploymentMultiTenantValDClosureBlockerStateAdvisory
	}
	return DeploymentMultiTenantValDClosureBlockerStateActive
}

func EvaluateDeploymentMultiTenantValDState(model DeploymentMultiTenantValDFoundation) string {
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) ||
		strings.TrimSpace(model.DependencyState) != DeploymentMultiTenantValDDependencyStateActive ||
		strings.TrimSpace(model.ConnectorCapabilityState) != DeploymentMultiTenantValDConnectorCapabilityStateActive ||
		strings.TrimSpace(model.OperatorActionState) != DeploymentMultiTenantValDOperatorActionStateActive ||
		strings.TrimSpace(model.SupportAccessState) != DeploymentMultiTenantValDSupportAccessStateActive ||
		strings.TrimSpace(model.BreakGlassState) != DeploymentMultiTenantValDBreakGlassStateActive ||
		strings.TrimSpace(model.MarketplaceMSPAuthorityState) != DeploymentMultiTenantValDMarketplaceMSPAuthorityStateActive ||
		strings.TrimSpace(model.AgenticOverlayState) != DeploymentMultiTenantValDAgenticOverlayStateActive ||
		strings.TrimSpace(model.NoOverclaimState) != DeploymentMultiTenantValDNoOverclaimStateActive ||
		strings.TrimSpace(model.ClosureBlockerState) != DeploymentMultiTenantValDClosureBlockerStateActive ||
		strings.TrimSpace(model.Point10State) != DeploymentMultiTenantPoint10StateNotComplete {
		return DeploymentMultiTenantValDStateBlocked
	}
	return DeploymentMultiTenantValDStateActive
}

func deploymentMultiTenantValDBlockingReasons(model DeploymentMultiTenantValDFoundation) []string {
	reasons := []string{}
	if !deploymentMultiTenantValDHasFoundationProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "aggregate_projection_disclaimer_blocked")
	}
	if model.DependencyState != DeploymentMultiTenantValDDependencyStateActive {
		reasons = append(reasons, "dependency_state_blocked")
	}
	if model.ConnectorCapabilityState != DeploymentMultiTenantValDConnectorCapabilityStateActive {
		reasons = append(reasons, "connector_capability_state_blocked")
	}
	if model.OperatorActionState != DeploymentMultiTenantValDOperatorActionStateActive {
		reasons = append(reasons, "operator_action_state_blocked")
	}
	if model.SupportAccessState != DeploymentMultiTenantValDSupportAccessStateActive {
		reasons = append(reasons, "support_access_state_blocked")
	}
	if model.BreakGlassState != DeploymentMultiTenantValDBreakGlassStateActive {
		reasons = append(reasons, "break_glass_state_blocked")
	}
	if model.MarketplaceMSPAuthorityState != DeploymentMultiTenantValDMarketplaceMSPAuthorityStateActive {
		reasons = append(reasons, "marketplace_msp_authority_state_blocked")
	}
	if model.AgenticOverlayState != DeploymentMultiTenantValDAgenticOverlayStateActive {
		reasons = append(reasons, "agentic_overlay_state_blocked")
	}
	if model.NoOverclaimState != DeploymentMultiTenantValDNoOverclaimStateActive {
		reasons = append(reasons, "no_overclaim_state_blocked")
	}
	if model.ClosureBlockerState != DeploymentMultiTenantValDClosureBlockerStateActive {
		reasons = append(reasons, "closure_blocker_state_not_clean")
	}
	if model.Point10State != DeploymentMultiTenantPoint10StateNotComplete {
		reasons = append(reasons, "point10_state_not_complete_guard_violated")
	}
	return reasons
}

func DeploymentMultiTenantValDFoundationModel() DeploymentMultiTenantValDFoundation {
	disclaimer := deploymentMultiTenantValDProjectionDisclaimer()
	runtimeAgent := DeploymentMultiTenantValDAgentRuntimeApprovalController{
		AgentID:                       "agent_runtime_approval_controller",
		AgentType:                     "runtime_approval_controller",
		TenantScope:                   "tenant:alpha",
		PermissionManifest:            "agent_permission_manifest",
		AllowedReadSurfaces:           []string{"deployment_health_surface", "tenant_boundary_surface"},
		AllowedRecommendationSurfaces: []string{"approval_queue_surface", "advisory_recommendation_surface"},
		ForbiddenMutationSurfaces:     []string{"canonical_evidence_spine", "production_mutation_surface"},
		ApprovalRequired:              true,
		ApprovalStatus:                deploymentMultiTenantValDApprovalStatusApproved,
		Approver:                      "human_approval_authority",
		ApprovalReason:                "human_approved_action_required",
		ApprovalQueue:                 "bounded_approval_queue",
		AuditID:                       "agent_runtime_audit_id",
		EvidenceRefs:                  append([]string{}, deploymentMultiTenantValDAgenticOverlayEvidenceRefs()...),
		RecommendationID:              "agent_runtime_recommendation",
		RecommendationType:            "approval_queue_review",
		HumanReviewRequired:           true,
		ExecutionAllowed:              false,
		ProjectionDisclaimer:          disclaimer,
		DiagnosticOutputComplete:      true,
		AgentNamingExact:              true,
		SafeWordingExamplePresent:     true,
		RunbookWordingComplete:        true,
	}
	learningLoop := DeploymentMultiTenantValDAgentLearningLoop{
		LearningAllowed:                 true,
		LearningMode:                    "offline_sandbox_only",
		TrainingDataScope:               "tenant:alpha training_data_scope",
		TrainingDataRefs:                append([]string{}, deploymentMultiTenantValDAgentLearningLoopEvidenceRefs()...),
		TrainingDataPrivacyFiltered:     true,
		TrainingDataTenantScoped:        true,
		TrainingDataCustomerApproved:    true,
		HumanFeedbackRefs:               []string{"human_feedback_ref"},
		HumanFeedbackAuditLinked:        true,
		TrainingApprovalRequired:        true,
		TrainingApprovalStatus:          "training_approved_active",
		TrainingApprover:                "human_training_approver",
		ModelCandidateID:                "candidate_model_id",
		ModelVersion:                    "candidate_model_version",
		BaselineModelVersion:            "baseline_model_version",
		EvaluationResultRefs:            []string{"evaluation_result_ref"},
		RegressionTestRefs:              []string{"regression_test_ref"},
		NoOverclaimCheckRefs:            []string{"no_overclaim_check_ref"},
		TenantScopeCheckRefs:            []string{"tenant_scope_check_ref"},
		ApprovalGateCheckRefs:           []string{"approval_gate_check_ref"},
		PromotionApprovalStatus:         "promotion_not_approved",
		RuntimeActivationApprovalStatus: "runtime_activation_not_approved",
		RecommendationApprovalStatus:    "recommendation_reviewed",
		ExecutionApprovalStatus:         "execution_not_approved",
		ModelUpgradeApprovalStatus:      "model_upgrade_not_approved",
		LearnedOutputAdvisoryOnly:       true,
		ApprovalRequired:                true,
		AuditID:                         "agent_learning_loop_audit_id",
		EvidenceRefs:                    append([]string{}, deploymentMultiTenantValDAgentLearningLoopEvidenceRefs()...),
		ProjectionDisclaimer:            disclaimer,
		DiagnosticOutputComplete:        true,
		LearningModeNamingExact:         true,
		ModelCandidateNamingExact:       true,
		SafeWordingExamplePresent:       true,
		RunbookWordingComplete:          true,
	}
	return DeploymentMultiTenantValDFoundation{
		CurrentState:                 DeploymentMultiTenantValDStateActive,
		Point10State:                 DeploymentMultiTenantPoint10StateNotComplete,
		ProjectionDisclaimer:         disclaimer,
		DependencyState:              DeploymentMultiTenantValDDependencyStateActive,
		ConnectorCapabilityState:     DeploymentMultiTenantValDConnectorCapabilityStateActive,
		OperatorActionState:          DeploymentMultiTenantValDOperatorActionStateActive,
		SupportAccessState:           DeploymentMultiTenantValDSupportAccessStateActive,
		BreakGlassState:              DeploymentMultiTenantValDBreakGlassStateActive,
		MarketplaceMSPAuthorityState: DeploymentMultiTenantValDMarketplaceMSPAuthorityStateActive,
		AgenticOverlayState:          DeploymentMultiTenantValDAgenticOverlayStateActive,
		NoOverclaimState:             DeploymentMultiTenantValDNoOverclaimStateActive,
		ClosureBlockerState:          DeploymentMultiTenantValDClosureBlockerStateActive,
		Dependency:                   deploymentMultiTenantValDDependencySnapshotModel(),
		ConnectorCapability: DeploymentMultiTenantValDConnectorCapabilityManifest{
			CurrentState:                       DeploymentMultiTenantValDConnectorCapabilityStateActive,
			ConnectorID:                        "sandboxed_connector_alpha",
			ConnectorType:                      "bounded_connector_type",
			TenantScope:                        "tenant:alpha",
			PermissionManifest:                 "connector_permission_manifest",
			CapabilityManifestPresent:          true,
			ReadCapabilities:                   []string{"read_evidence_surface", "read_diagnostic_surface"},
			WriteCapabilities:                  []string{"write_advisory_surface"},
			EvidenceTypes:                      []string{"connector_signal_evidence"},
			RetryPolicy:                        "bounded_retry_policy",
			ReplayPolicy:                       "bounded_replay_policy",
			RateLimitPolicy:                    "bounded_rate_limit_policy",
			AuditRequired:                      true,
			AuditID:                            "connector_audit_id",
			Reason:                             "bounded_connector_reason",
			SourceOfTruth:                      "advisory_evidence_input",
			TenantScopedExecution:              true,
			RecoveryBehavior:                   "bounded_recovery_behavior",
			FailureBehavior:                    "bounded_failure_behavior",
			EvidenceRefs:                       append([]string{}, deploymentMultiTenantValDConnectorEvidenceRefs()...),
			FreshnessState:                     IntelligenceCalibrationFreshnessFresh,
			ProjectionDisclaimer:               disclaimer,
			DiagnosticOutputComplete:           true,
			ConnectorNamingExact:               true,
			SafeConnectorWordingExamplePresent: true,
		},
		OperatorAction: DeploymentMultiTenantValDOperatorActionPolicy{
			CurrentState:                      DeploymentMultiTenantValDOperatorActionStateActive,
			Actor:                             "operator_actor_alpha",
			ActorType:                         "operator",
			TenantTarget:                      "tenant:alpha",
			ActionScope:                       "tenant:alpha operator_action_scope",
			ActionType:                        "bounded_operator_review",
			Reason:                            "bounded_operator_reason",
			AuthorizationBasis:                "authorization_basis",
			Approver:                          "human_approval_authority",
			ApprovalStatus:                    deploymentMultiTenantValDApprovalStatusApproved,
			AuthorityBasis:                    "authority_basis",
			Expiry:                            "expiry_2026_12_31t00_00_00z",
			RevocationPath:                    "revocation_path",
			AuditID:                           "operator_action_audit_id",
			EvidenceRefs:                      append([]string{}, deploymentMultiTenantValDOperatorEvidenceRefs()...),
			RBACABACEnforced:                  true,
			SSOContextBound:                   true,
			TenantScopeBound:                  true,
			SupportScopeBound:                 true,
			DiagnosticOutputComplete:          true,
			ProjectionDisclaimer:              disclaimer,
			OperatorActionNamingExact:         true,
			SafeOperatorWordingExamplePresent: true,
		},
		SupportAccess: DeploymentMultiTenantValDSupportAccessEnforcement{
			CurrentState:                     DeploymentMultiTenantValDSupportAccessStateActive,
			SupportAccessID:                  "support_access_id",
			SupportActor:                     "support_actor_alpha",
			TenantTarget:                     "tenant:alpha",
			SupportScope:                     "tenant:alpha support_scope",
			SSOSessionReference:              "sso_session_reference",
			RBACRole:                         "rbac_role",
			ABACConditions:                   "abac_conditions",
			AuthorityBasis:                   "authority_basis",
			Reason:                           "support_reason",
			Approver:                         "human_approval_authority",
			ApprovalStatus:                   deploymentMultiTenantValDApprovalStatusApproved,
			Expiry:                           "expiry_2026_12_31t00_00_00z",
			RevocationPath:                   "revocation_path",
			AuditID:                          "support_access_audit_id",
			EvidenceRefs:                     append([]string{}, deploymentMultiTenantValDSupportEvidenceRefs()...),
			SupportVisibilityBoundary:        "tenant_support_visibility_boundary",
			DataResidencyBoundaryRespected:   true,
			TenantIsolationBoundaryRespected: true,
			DiagnosticOutputComplete:         true,
			ProjectionDisclaimer:             disclaimer,
		},
		BreakGlass: DeploymentMultiTenantValDBreakGlassAccess{
			CurrentState:                     DeploymentMultiTenantValDBreakGlassStateActive,
			BreakGlassID:                     "break_glass_id",
			EmergencyReason:                  "emergency_reason",
			Actor:                            "break_glass_actor",
			TenantTarget:                     "tenant:alpha",
			ActionScope:                      "tenant:alpha break_glass_scope",
			AuthorizationBasis:               "break_glass_authority_basis",
			Approver:                         "human_approval_authority",
			ApprovalStatus:                   deploymentMultiTenantValDApprovalStatusApproved,
			Expiry:                           "expiry_2026_12_31t00_00_00z",
			RevocationPath:                   "revocation_path",
			AuditID:                          "break_glass_audit_id",
			EvidenceRefs:                     append([]string{}, deploymentMultiTenantValDBreakGlassEvidenceRefs()...),
			PostActionReviewRequired:         true,
			PostActionReviewState:            deploymentMultiTenantValDBreakGlassReviewActive,
			TenantScopeBound:                 true,
			DataResidencyBoundaryRespected:   true,
			TenantIsolationBoundaryRespected: true,
			SupportVisibilityBoundary:        "tenant_support_visibility_boundary",
			DiagnosticOutputComplete:         true,
			ProjectionDisclaimer:             disclaimer,
		},
		MarketplaceMSPAuthority: DeploymentMultiTenantValDMarketplaceMSPAuthorityBoundary{
			CurrentState:                    DeploymentMultiTenantValDMarketplaceMSPAuthorityStateActive,
			MarketplaceProfile:              "marketplace_deployment_profile",
			MSPOperatorScope:                "tenant:alpha msp_operator_scope",
			PartnerScope:                    "tenant:alpha partner_scope",
			TenantScope:                     "tenant:alpha",
			DeploymentScope:                 "tenant_deployment_scope",
			SupportScope:                    "tenant_support_scope",
			CustomerReadyValidationEvidence: "customer_ready_validation_evidence",
			AuthorityBoundary:               "bounded_authority_boundary",
			ApprovalBoundary:                "bounded_approval_boundary",
			AuditID:                         "marketplace_msp_audit_id",
			Reason:                          "bounded_marketplace_reason",
			Expiry:                          "expiry_2026_12_31t00_00_00z",
			RevocationPath:                  "revocation_path",
			EvidenceRefs:                    append([]string{}, deploymentMultiTenantValDMarketplaceMSPEvidenceRefs()...),
			ProjectionDisclaimer:            disclaimer,
			DiagnosticOutputComplete:        true,
		},
		AgenticOverlay: DeploymentMultiTenantValDAgenticOverlay{
			CurrentState:              DeploymentMultiTenantValDAgenticOverlayStateActive,
			LearningLoopState:         DeploymentMultiTenantValDAgentLearningLoopStateActive,
			RuntimeApprovalController: runtimeAgent,
			TenantBoundaryContainmentAgent: DeploymentMultiTenantValDAgentRecommendation{
				AgentID:                       "tenant_boundary_containment_agent",
				AgentType:                     "tenant_boundary_containment_agent",
				TenantScope:                   "tenant:alpha",
				PermissionManifest:            "containment_permission_manifest",
				AllowedReadSurfaces:           []string{"tenant_boundary_surface", "data_residency_surface"},
				AllowedRecommendationSurfaces: []string{"containment_recommendation_surface"},
				ForbiddenMutationSurfaces:     []string{"canonical_evidence_spine", "production_mutation_surface"},
				ApprovalRequired:              true,
				ApprovalStatus:                deploymentMultiTenantValDApprovalStatusApproved,
				Approver:                      "human_approval_authority",
				ApprovalReason:                "human_approved_action_required",
				AuditID:                       "containment_agent_audit_id",
				EvidenceRefs:                  append([]string{}, deploymentMultiTenantValDAgenticOverlayEvidenceRefs()...),
				RecommendationID:              "containment_recommendation",
				RecommendationType:            "containment_recommendation",
				HumanReviewRequired:           true,
				ExecutionAllowed:              false,
				ProjectionDisclaimer:          disclaimer,
				DiagnosticOutputComplete:      true,
				AgentNamingExact:              true,
				SafeWordingExamplePresent:     true,
			},
			DeploymentHealthPreflightAgent: DeploymentMultiTenantValDAgentRecommendation{
				AgentID:                       "deployment_health_preflight_agent",
				AgentType:                     "deployment_health_preflight_agent",
				TenantScope:                   "tenant:alpha",
				PermissionManifest:            "deployment_agent_permission_manifest",
				AllowedReadSurfaces:           []string{"deployment_preflight_surface", "ha_recovery_surface"},
				AllowedRecommendationSurfaces: []string{"deployment_diagnostic_surface"},
				ForbiddenMutationSurfaces:     []string{"canonical_evidence_spine", "production_mutation_surface"},
				ApprovalRequired:              true,
				ApprovalStatus:                deploymentMultiTenantValDApprovalStatusApproved,
				Approver:                      "human_approval_authority",
				ApprovalReason:                "human_approved_action_required",
				AuditID:                       "deployment_agent_audit_id",
				EvidenceRefs:                  append([]string{}, deploymentMultiTenantValDAgenticOverlayEvidenceRefs()...),
				RecommendationID:              "deployment_preflight_recommendation",
				RecommendationType:            "deployment_preflight_recommendation",
				HumanReviewRequired:           true,
				ExecutionAllowed:              false,
				ProjectionDisclaimer:          disclaimer,
				DiagnosticOutputComplete:      true,
				AgentNamingExact:              true,
				SafeWordingExamplePresent:     true,
			},
			ConnectorOperatorMisuseWatchAgent: DeploymentMultiTenantValDAgentRecommendation{
				AgentID:                       "connector_operator_misuse_watch_agent",
				AgentType:                     "connector_operator_misuse_watch_agent",
				TenantScope:                   "tenant:alpha",
				PermissionManifest:            "misuse_watch_permission_manifest",
				AllowedReadSurfaces:           []string{"connector_surface", "operator_action_surface", "support_access_surface"},
				AllowedRecommendationSurfaces: []string{"misuse_signal_surface"},
				ForbiddenMutationSurfaces:     []string{"canonical_evidence_spine", "production_mutation_surface"},
				ApprovalRequired:              true,
				ApprovalStatus:                deploymentMultiTenantValDApprovalStatusApproved,
				Approver:                      "human_approval_authority",
				ApprovalReason:                "human_approved_action_required",
				AuditID:                       "misuse_watch_audit_id",
				EvidenceRefs:                  append([]string{}, deploymentMultiTenantValDAgenticOverlayEvidenceRefs()...),
				RecommendationID:              "misuse_signal_recommendation",
				RecommendationType:            "connector_operator_misuse_signal",
				HumanReviewRequired:           true,
				ExecutionAllowed:              false,
				ProjectionDisclaimer:          disclaimer,
				DiagnosticOutputComplete:      true,
				AgentNamingExact:              true,
				SafeWordingExamplePresent:     true,
			},
			RecoveryRebuildRecommendationAgent: DeploymentMultiTenantValDAgentRecommendation{
				AgentID:                       "recovery_rebuild_recommendation_agent",
				AgentType:                     "recovery_rebuild_recommendation_agent",
				TenantScope:                   "tenant:alpha",
				PermissionManifest:            "recovery_agent_permission_manifest",
				AllowedReadSurfaces:           []string{"recovery_readiness_surface", "tenant_isolation_surface", "data_residency_surface"},
				AllowedRecommendationSurfaces: []string{"recovery_recommendation_surface"},
				ForbiddenMutationSurfaces:     []string{"canonical_evidence_spine", "production_mutation_surface"},
				ApprovalRequired:              true,
				ApprovalStatus:                deploymentMultiTenantValDApprovalStatusApproved,
				Approver:                      "human_approval_authority",
				ApprovalReason:                "human_approved_action_required",
				AuditID:                       "recovery_agent_audit_id",
				EvidenceRefs:                  append([]string{}, deploymentMultiTenantValDAgenticOverlayEvidenceRefs()...),
				RecommendationID:              "recovery_recommendation",
				RecommendationType:            "recovery_rebuild_recommendation",
				HumanReviewRequired:           true,
				ExecutionAllowed:              false,
				ProjectionDisclaimer:          disclaimer,
				DiagnosticOutputComplete:      true,
				RecoveryEvidencePack:          "backup_restore_dr_evidence_pack",
				AgentNamingExact:              true,
				SafeWordingExamplePresent:     true,
			},
			LearningLoop:         learningLoop,
			ProjectionDisclaimer: disclaimer,
		},
		NoOverclaim: DeploymentMultiTenantValDNoOverclaimDiscipline{
			CurrentState:         DeploymentMultiTenantValDNoOverclaimStateActive,
			ProjectionDisclaimer: disclaimer,
		},
		ClosureBlockerOverlay: DeploymentMultiTenantValDClosureBlockerOverlay{
			CurrentState:         DeploymentMultiTenantValDClosureBlockerStateActive,
			ProjectionDisclaimer: disclaimer,
		},
	}
}

func ComputeDeploymentMultiTenantValDFoundation(model DeploymentMultiTenantValDFoundation) DeploymentMultiTenantValDFoundation {
	model.DependencyState = EvaluateDeploymentMultiTenantValDDependencyState(model.Dependency)
	model.ConnectorCapabilityState = EvaluateDeploymentMultiTenantValDConnectorCapabilityState(model.ConnectorCapability)
	model.OperatorActionState = EvaluateDeploymentMultiTenantValDOperatorActionState(model.OperatorAction)
	model.SupportAccessState = EvaluateDeploymentMultiTenantValDSupportAccessState(model.SupportAccess)
	model.BreakGlassState = EvaluateDeploymentMultiTenantValDBreakGlassState(model.BreakGlass)
	model.MarketplaceMSPAuthorityState = EvaluateDeploymentMultiTenantValDMarketplaceMSPAuthorityState(model.MarketplaceMSPAuthority)
	model.AgenticOverlay.LearningLoopState = EvaluateDeploymentMultiTenantValDAgentLearningLoopState(model.AgenticOverlay.LearningLoop)
	model.AgenticOverlayState = EvaluateDeploymentMultiTenantValDAgenticOverlayState(model.AgenticOverlay)
	model.AgenticOverlay.CurrentState = model.AgenticOverlayState
	model.NoOverclaimState = EvaluateDeploymentMultiTenantValDNoOverclaimState(model.NoOverclaim)
	model.ClosureBlockerOverlay = DeploymentMultiTenantValDClosureBlockerOverlay{
		ProjectionDisclaimer: deploymentMultiTenantValDProjectionDisclaimer(),
		Findings:             deploymentMultiTenantValDClosureBlockerFindings(model),
	}
	model.ClosureBlockerState = EvaluateDeploymentMultiTenantValDClosureBlockerState(model.ClosureBlockerOverlay)
	model.ClosureBlockerOverlay.CurrentState = model.ClosureBlockerState
	model.Point10State = EvaluateDeploymentMultiTenantPoint10State(model.CurrentState)
	model.CurrentState = EvaluateDeploymentMultiTenantValDState(model)
	model.Point10State = EvaluateDeploymentMultiTenantPoint10State(model.CurrentState)
	model.BlockingReasons = deploymentMultiTenantValDBlockingReasons(model)
	return model
}

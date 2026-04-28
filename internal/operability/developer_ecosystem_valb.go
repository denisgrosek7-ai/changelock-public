package operability

import "strings"

const (
	DeveloperEcosystemValBValECompatibilityStateActive     = "developer_ecosystem_valb_vale_compatibility_active"
	DeveloperEcosystemValBValECompatibilityStatePartial    = "developer_ecosystem_valb_vale_compatibility_partial"
	DeveloperEcosystemValBValECompatibilityStateIncomplete = "developer_ecosystem_valb_vale_compatibility_incomplete"
	DeveloperEcosystemValBValECompatibilityStateBlocked    = "developer_ecosystem_valb_vale_compatibility_blocked"
	DeveloperEcosystemValBValECompatibilityStateUnknown    = "developer_ecosystem_valb_vale_compatibility_unknown"

	DeveloperEcosystemValBDependencyStateActive     = "developer_ecosystem_valb_dependency_active"
	DeveloperEcosystemValBDependencyStatePartial    = "developer_ecosystem_valb_dependency_partial"
	DeveloperEcosystemValBDependencyStateIncomplete = "developer_ecosystem_valb_dependency_incomplete"
	DeveloperEcosystemValBDependencyStateBlocked    = "developer_ecosystem_valb_dependency_blocked"
	DeveloperEcosystemValBDependencyStateUnknown    = "developer_ecosystem_valb_dependency_unknown"

	DeveloperEcosystemValBRepoConfigSchemaStateActive     = "developer_ecosystem_valb_repo_config_schema_active"
	DeveloperEcosystemValBRepoConfigSchemaStatePartial    = "developer_ecosystem_valb_repo_config_schema_partial"
	DeveloperEcosystemValBRepoConfigSchemaStateIncomplete = "developer_ecosystem_valb_repo_config_schema_incomplete"
	DeveloperEcosystemValBRepoConfigSchemaStateBlocked    = "developer_ecosystem_valb_repo_config_schema_blocked"
	DeveloperEcosystemValBRepoConfigSchemaStateUnknown    = "developer_ecosystem_valb_repo_config_schema_unknown"

	DeveloperEcosystemValBRepoConfigValidationStateActive     = "developer_ecosystem_valb_repo_config_validation_active"
	DeveloperEcosystemValBRepoConfigValidationStatePartial    = "developer_ecosystem_valb_repo_config_validation_partial"
	DeveloperEcosystemValBRepoConfigValidationStateIncomplete = "developer_ecosystem_valb_repo_config_validation_incomplete"
	DeveloperEcosystemValBRepoConfigValidationStateBlocked    = "developer_ecosystem_valb_repo_config_validation_blocked"
	DeveloperEcosystemValBRepoConfigValidationStateUnknown    = "developer_ecosystem_valb_repo_config_validation_unknown"

	DeveloperEcosystemValBPolicyPreviewStateActive     = "developer_ecosystem_valb_policy_preview_active"
	DeveloperEcosystemValBPolicyPreviewStatePartial    = "developer_ecosystem_valb_policy_preview_partial"
	DeveloperEcosystemValBPolicyPreviewStateIncomplete = "developer_ecosystem_valb_policy_preview_incomplete"
	DeveloperEcosystemValBPolicyPreviewStateBlocked    = "developer_ecosystem_valb_policy_preview_blocked"
	DeveloperEcosystemValBPolicyPreviewStateUnknown    = "developer_ecosystem_valb_policy_preview_unknown"

	DeveloperEcosystemValBLocalCIContinuityStateActive     = "developer_ecosystem_valb_local_ci_continuity_active"
	DeveloperEcosystemValBLocalCIContinuityStatePartial    = "developer_ecosystem_valb_local_ci_continuity_partial"
	DeveloperEcosystemValBLocalCIContinuityStateIncomplete = "developer_ecosystem_valb_local_ci_continuity_incomplete"
	DeveloperEcosystemValBLocalCIContinuityStateBlocked    = "developer_ecosystem_valb_local_ci_continuity_blocked"
	DeveloperEcosystemValBLocalCIContinuityStateUnknown    = "developer_ecosystem_valb_local_ci_continuity_unknown"

	DeveloperEcosystemValBAPISDKSurfaceStateActive     = "developer_ecosystem_valb_api_sdk_surface_active"
	DeveloperEcosystemValBAPISDKSurfaceStatePartial    = "developer_ecosystem_valb_api_sdk_surface_partial"
	DeveloperEcosystemValBAPISDKSurfaceStateIncomplete = "developer_ecosystem_valb_api_sdk_surface_incomplete"
	DeveloperEcosystemValBAPISDKSurfaceStateBlocked    = "developer_ecosystem_valb_api_sdk_surface_blocked"
	DeveloperEcosystemValBAPISDKSurfaceStateUnknown    = "developer_ecosystem_valb_api_sdk_surface_unknown"

	DeveloperEcosystemValBExamplesTemplatesStateActive     = "developer_ecosystem_valb_examples_templates_active"
	DeveloperEcosystemValBExamplesTemplatesStatePartial    = "developer_ecosystem_valb_examples_templates_partial"
	DeveloperEcosystemValBExamplesTemplatesStateIncomplete = "developer_ecosystem_valb_examples_templates_incomplete"
	DeveloperEcosystemValBExamplesTemplatesStateBlocked    = "developer_ecosystem_valb_examples_templates_blocked"
	DeveloperEcosystemValBExamplesTemplatesStateUnknown    = "developer_ecosystem_valb_examples_templates_unknown"

	DeveloperEcosystemValBAPIVersioningStateActive     = "developer_ecosystem_valb_api_versioning_active"
	DeveloperEcosystemValBAPIVersioningStatePartial    = "developer_ecosystem_valb_api_versioning_partial"
	DeveloperEcosystemValBAPIVersioningStateIncomplete = "developer_ecosystem_valb_api_versioning_incomplete"
	DeveloperEcosystemValBAPIVersioningStateBlocked    = "developer_ecosystem_valb_api_versioning_blocked"
	DeveloperEcosystemValBAPIVersioningStateUnknown    = "developer_ecosystem_valb_api_versioning_unknown"

	DeveloperEcosystemValBNoOverclaimStateActive     = "developer_ecosystem_valb_no_overclaim_active"
	DeveloperEcosystemValBNoOverclaimStatePartial    = "developer_ecosystem_valb_no_overclaim_partial"
	DeveloperEcosystemValBNoOverclaimStateIncomplete = "developer_ecosystem_valb_no_overclaim_incomplete"
	DeveloperEcosystemValBNoOverclaimStateBlocked    = "developer_ecosystem_valb_no_overclaim_blocked"
	DeveloperEcosystemValBNoOverclaimStateUnknown    = "developer_ecosystem_valb_no_overclaim_unknown"

	DeveloperEcosystemValBStateActive     = "developer_ecosystem_valb_active"
	DeveloperEcosystemValBStatePartial    = "developer_ecosystem_valb_partial"
	DeveloperEcosystemValBStateIncomplete = "developer_ecosystem_valb_incomplete"
	DeveloperEcosystemValBStateBlocked    = "developer_ecosystem_valb_blocked"
	DeveloperEcosystemValBStateUnknown    = "developer_ecosystem_valb_unknown"

	DeveloperEcosystemRepoConfigSchemaV1Advisory = "changelock_repo_config.v1alpha_advisory"

	DeveloperEcosystemPolicyPreviewResultAdvisory    = "policy_preview_advisory"
	DeveloperEcosystemPolicyPreviewResultDegraded    = "policy_preview_degraded"
	DeveloperEcosystemPolicyPreviewResultUnavailable = "policy_preview_unavailable"

	DeveloperEcosystemLocalOnlyFixtureVisible = "visible_local_only_fixture"
	DeveloperEcosystemContinuityFailClosed    = "fail_closed_on_missing_continuity"

	DeveloperEcosystemSDKRateExpectationBounded = "bounded_developer_assist_budget"

	DeveloperEcosystemValBRepoConfigCompatibilityBehaviorBounded = "visible_compatibility_window"
	DeveloperEcosystemSDKCompatibilityVisible                    = "visible_compatibility_window"
	DeveloperEcosystemSDKSchemaExactModels                       = "exact_versioned_models"
	DeveloperEcosystemValBAPIVersionIdentity                     = "developer_api_sdk_surface.v1"
	DeveloperEcosystemValBAPICompatibilityWindow                 = "one_major_visible_window"
)

type DeveloperEcosystemValBValECompatibilityGate struct {
	CurrentState         string   `json:"current_state"`
	GateID               string   `json:"gate_id"`
	Version              string   `json:"version"`
	ValECurrentState     string   `json:"vale_current_state"`
	Point7State          string   `json:"point_7_state"`
	PassRuleState        string   `json:"pass_rule_state"`
	NoOverclaimState     string   `json:"no_overclaim_state"`
	ProofSurfaceState    string   `json:"proof_surface_state"`
	EvidenceQualityState string   `json:"evidence_quality_state"`
	Point7PassAllowed    bool     `json:"point_7_pass_allowed"`
	Point7PassReason     string   `json:"point_7_pass_reason"`
	SurfaceRefs          []string `json:"surface_refs,omitempty"`
	EvidenceRefs         []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValBDependencySnapshot struct {
	ValACurrentState         string   `json:"vala_current_state"`
	ValAPoint8State          string   `json:"vala_point_8_state"`
	ValADependencyState      string   `json:"vala_dependency_state"`
	IDEBaselineState         string   `json:"ide_baseline_state"`
	TrustFeedbackState       string   `json:"trust_feedback_state"`
	CAVIVEXContextState      string   `json:"cavi_vex_context_state"`
	LocalAdvisoryState       string   `json:"local_advisory_state"`
	ValidationHarnessState   string   `json:"validation_harness_state"`
	MockVerificationState    string   `json:"mock_verification_state"`
	InspectExplainState      string   `json:"inspect_explain_state"`
	DegradedModeState        string   `json:"degraded_mode_state"`
	NoOverclaimState         string   `json:"no_overclaim_state"`
	ValAProofSurfaceRefs     []string `json:"vala_proof_surface_refs,omitempty"`
	ValAEvidenceRefs         []string `json:"vala_evidence_refs,omitempty"`
	ValAProjectionDisclaimer string   `json:"vala_projection_disclaimer"`
}

type DeveloperEcosystemValBRepoConfigSchemaContract struct {
	CurrentState                     string   `json:"current_state"`
	ContractID                       string   `json:"contract_id"`
	Version                          string   `json:"version"`
	SchemaVersion                    string   `json:"schema_version"`
	SupportedSchemaVersions          []string `json:"supported_schema_versions,omitempty"`
	AllowedTopLevelFields            []string `json:"allowed_top_level_fields,omitempty"`
	AllowedAdvisoryConfigFields      []string `json:"allowed_advisory_config_fields,omitempty"`
	AllowedLocalValidationFields     []string `json:"allowed_local_validation_fields,omitempty"`
	AllowedCIContinuityFields        []string `json:"allowed_ci_continuity_fields,omitempty"`
	AllowedTemplateProfileRefs       []string `json:"allowed_template_profile_refs,omitempty"`
	DisallowedEnterpriseOverrideRefs []string `json:"disallowed_enterprise_override_fields,omitempty"`
	UnknownFieldHandling             string   `json:"unknown_field_handling"`
	UnsupportedVersionHandling       string   `json:"unsupported_version_handling"`
	CompatibilityBehavior            string   `json:"compatibility_behavior"`
	LocalAdvisoryScope               bool     `json:"local_advisory_scope"`
	ScopeBounded                     bool     `json:"scope_bounded"`
	ReviewBound                      bool     `json:"review_bound"`
	EnterprisePolicyOverrideAttempt  bool     `json:"enterprise_policy_override_attempt"`
	DisableCanonicalEvidenceRules    bool     `json:"disable_canonical_evidence_rules"`
	GrantApprovalAuthority           bool     `json:"grant_approval_authority"`
	ChangeProductionTrustBoundary    bool     `json:"change_production_trust_boundary"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValBRepoConfigValidationDiscipline struct {
	CurrentState                        string `json:"current_state"`
	DisciplineID                        string `json:"discipline_id"`
	Version                             string `json:"version"`
	SchemaVersionValid                  bool   `json:"schema_version_valid"`
	FieldAllowlistValid                 bool   `json:"field_allowlist_valid"`
	NormalizedExactSetValid             bool   `json:"normalized_exact_set_valid"`
	DuplicateEntriesDetected            bool   `json:"duplicate_entries_detected"`
	UnsupportedValueDetected            bool   `json:"unsupported_value_detected"`
	DeprecatedFieldDetected             bool   `json:"deprecated_field_detected"`
	DeprecatedFieldCompatibilityVisible bool   `json:"deprecated_field_compatibility_visible"`
	MalformedConfigDetected             bool   `json:"malformed_config_detected"`
	StaleTemplateReferenceDetected      bool   `json:"stale_template_reference_detected"`
	EvidenceContextLinked               bool   `json:"evidence_context_linked"`
	MissingGovernanceField              bool   `json:"missing_governance_field"`
	PermissiveFallback                  bool   `json:"permissive_fallback"`
	AdvisoryOnly                        bool   `json:"advisory_only"`
	ProjectionDisclaimer                string `json:"projection_disclaimer"`
}

type DeveloperEcosystemValBPolicyPreviewPath struct {
	CurrentState                     string   `json:"current_state"`
	ModelID                          string   `json:"model_id"`
	Version                          string   `json:"version"`
	RepoLocalInputSnapshot           string   `json:"repo_local_input_snapshot"`
	EnterprisePolicyBoundarySnapshot string   `json:"enterprise_policy_boundary_snapshot"`
	LocalAdvisoryPreviewResult       string   `json:"local_advisory_preview_result"`
	AffectedRuleHints                []string `json:"affected_rule_hints,omitempty"`
	ProductionOnlyUnknownsVisible    bool     `json:"production_only_unknowns_visible"`
	MismatchDegradedReason           string   `json:"mismatch_degraded_reason"`
	EvidenceContextRefs              []string `json:"evidence_context_refs,omitempty"`
	RemediationHints                 []string `json:"remediation_hints,omitempty"`
	ApprovesDeployment               bool     `json:"approves_deployment"`
	OverridesPolicy                  bool     `json:"overrides_policy"`
	MutatesCanonicalEvidence         bool     `json:"mutates_canonical_evidence"`
	SuppressesFailures               bool     `json:"suppresses_failures"`
	CanonicalDecisionClaim           bool     `json:"canonical_decision_claim"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValBLocalCIContinuityContract struct {
	CurrentState                  string   `json:"current_state"`
	ContractID                    string   `json:"contract_id"`
	Version                       string   `json:"version"`
	LocalValidationDescriptor     string   `json:"local_validation_descriptor"`
	CIValidationDescriptor        string   `json:"ci_validation_descriptor"`
	SharedInputIdentityFields     []string `json:"shared_input_identity_fields,omitempty"`
	ExpectedOutputClasses         []string `json:"expected_output_classes,omitempty"`
	MismatchVisible               bool     `json:"mismatch_visible"`
	LocalOnlyFixtureBehavior      string   `json:"local_only_fixture_behavior"`
	ProductionOnlyCIChecksVisible bool     `json:"production_only_ci_checks_visible"`
	EvidenceHandoffRefs           []string `json:"evidence_handoff_refs,omitempty"`
	FailureExplanationVisible     bool     `json:"failure_explanation_visible"`
	LocalPassBecomesCIPass        bool     `json:"local_pass_becomes_ci_pass"`
	ProductionCIEquivalenceClaim  bool     `json:"production_ci_equivalence_claim"`
	MissingDescriptors            bool     `json:"missing_descriptors"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValBDeveloperAPISDKSurface struct {
	CurrentState                  string   `json:"current_state"`
	ContractID                    string   `json:"contract_id"`
	Version                       string   `json:"version"`
	StableSurfaceDescriptors      []string `json:"stable_surface_descriptors,omitempty"`
	RequestResponseModelVersion   string   `json:"request_response_model_version"`
	SupportedVersions             []string `json:"supported_versions,omitempty"`
	DeprecatedVersions            []string `json:"deprecated_versions,omitempty"`
	PermissionAwareOperations     bool     `json:"permission_aware_operations"`
	ReadOnlyAdvisoryDefault       bool     `json:"read_only_advisory_default"`
	InspectExplainHelpers         []string `json:"inspect_explain_helpers,omitempty"`
	ValidationReplayHelpers       []string `json:"validation_replay_helpers,omitempty"`
	ErrorDegradedStateModel       bool     `json:"error_degraded_state_model"`
	RatePerformanceExpectation    string   `json:"rate_performance_expectation"`
	CompatibilityBehavior         string   `json:"compatibility_behavior"`
	FailClosedUnsupportedBehavior bool     `json:"fail_closed_unsupported_behavior"`
	HiddenMutationPath            bool     `json:"hidden_mutation_path"`
	ApprovesDeployment            bool     `json:"approves_deployment"`
	CertifiesTrust                bool     `json:"certifies_trust"`
	BypassesCanonicalEvidence     bool     `json:"bypasses_canonical_evidence"`
	SuppressesFailures            bool     `json:"suppresses_failures"`
	UnsupportedVersionActive      bool     `json:"unsupported_version_active"`
	DeprecatedVersionVisible      bool     `json:"deprecated_version_visible"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValBExamplesTemplatesContract struct {
	CurrentState                        string `json:"current_state"`
	ContractID                          string `json:"contract_id"`
	Version                             string `json:"version"`
	NewServiceTemplateDescriptor        string `json:"new_service_template_descriptor"`
	StarterPackDescriptor               string `json:"starter_pack_descriptor"`
	DefaultConfigDescriptor             string `json:"default_config_descriptor"`
	LocalValidationExampleDescriptor    string `json:"local_validation_example_descriptor"`
	CAVIVEXExampleDescriptor            string `json:"cavi_vex_example_descriptor"`
	SDKUsageExampleDescriptor           string `json:"sdk_usage_example_descriptor"`
	CIContinuityExampleDescriptor       string `json:"ci_continuity_example_descriptor"`
	TemplateVersion                     string `json:"template_version"`
	FreshnessState                      string `json:"freshness_state"`
	CompatibilityMetadataVisible        bool   `json:"compatibility_metadata_visible"`
	ComplianceCertificationClaim        bool   `json:"compliance_certification_claim"`
	ProductionApprovalClaim             bool   `json:"production_approval_claim"`
	CanonicalEvidenceClaim              bool   `json:"canonical_evidence_claim"`
	StaleTemplateReferenceDetected      bool   `json:"stale_template_reference_detected"`
	DeprecatedTemplateReferenceDetected bool   `json:"deprecated_template_reference_detected"`
	DeprecatedCompatibilityVisible      bool   `json:"deprecated_compatibility_visible"`
	ProjectionDisclaimer                string `json:"projection_disclaimer"`
}

type DeveloperEcosystemValBAPIVersioningCompatibility struct {
	CurrentState                   string   `json:"current_state"`
	DisciplineID                   string   `json:"discipline_id"`
	Version                        string   `json:"version"`
	VersionIdentity                string   `json:"version_identity"`
	SupportedVersions              []string `json:"supported_versions,omitempty"`
	DeprecatedVersions             []string `json:"deprecated_versions,omitempty"`
	UnsupportedVersions            []string `json:"unsupported_versions,omitempty"`
	CompatibilityWindow            string   `json:"compatibility_window"`
	MigrationHint                  string   `json:"migration_hint"`
	FailClosedUnsupportedBehavior  bool     `json:"fail_closed_unsupported_behavior"`
	StaleClientBehaviorVisible     bool     `json:"stale_client_behavior_visible"`
	SchemaCompatibilityBehavior    string   `json:"schema_compatibility_behavior"`
	UnknownVersionDetected         bool     `json:"unknown_version_detected"`
	UnsupportedVersionDetected     bool     `json:"unsupported_version_detected"`
	DeprecatedVersionDetected      bool     `json:"deprecated_version_detected"`
	DeprecatedCompatibilityVisible bool     `json:"deprecated_compatibility_visible"`
	ProjectionDisclaimer           string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValBNoOverclaimDiscipline struct {
	CurrentState                        string `json:"current_state"`
	DisciplineID                        string `json:"discipline_id"`
	Version                             string `json:"version"`
	ProductionApprovalClaim             bool   `json:"production_approval_claim"`
	CertificationClaim                  bool   `json:"certification_claim"`
	GovernanceReplacementClaim          bool   `json:"governance_replacement_claim"`
	EnterprisePolicyOverrideClaim       bool   `json:"enterprise_policy_override_claim"`
	CanonicalTruthClaim                 bool   `json:"canonical_truth_claim"`
	ComplianceGuaranteeClaim            bool   `json:"compliance_guarantee_claim"`
	DeveloperFastTrackApprovalClaim     bool   `json:"developer_fast_track_approval_claim"`
	TemplateFormalEvidenceClaim         bool   `json:"template_formal_evidence_claim"`
	SDKProductionAuthorizationClaim     bool   `json:"sdk_production_authorization_claim"`
	RepoLocalPolicyAuthorityClaim       bool   `json:"repo_local_policy_authority_claim"`
	PolicyPreviewCanonicalDecisionClaim bool   `json:"policy_preview_canonical_decision_claim"`
	Point8PassClaim                     bool   `json:"point8_pass_claim"`
	ProjectionDisclaimer                string `json:"projection_disclaimer"`
}

type DeveloperEcosystemValBIntegration struct {
	CurrentState              string                                               `json:"current_state"`
	Point8State               string                                               `json:"point_8_state"`
	ValECompatibilityState    string                                               `json:"vale_compatibility_state"`
	DependencyState           string                                               `json:"dependency_state"`
	RepoConfigSchemaState     string                                               `json:"repo_config_schema_state"`
	RepoConfigValidationState string                                               `json:"repo_config_validation_state"`
	PolicyPreviewState        string                                               `json:"policy_preview_state"`
	LocalCIContinuityState    string                                               `json:"local_ci_continuity_state"`
	APISDKSurfaceState        string                                               `json:"api_sdk_surface_state"`
	ExamplesTemplatesState    string                                               `json:"examples_templates_state"`
	APIVersioningState        string                                               `json:"api_versioning_state"`
	NoOverclaimState          string                                               `json:"no_overclaim_state"`
	IntegrationID             string                                               `json:"integration_id"`
	Version                   string                                               `json:"version"`
	ValECompatibility         DeveloperEcosystemValBValECompatibilityGate          `json:"vale_compatibility"`
	Dependency                DeveloperEcosystemValBDependencySnapshot             `json:"dependency"`
	RepoConfigSchema          DeveloperEcosystemValBRepoConfigSchemaContract       `json:"repo_config_schema"`
	RepoConfigValidation      DeveloperEcosystemValBRepoConfigValidationDiscipline `json:"repo_config_validation"`
	PolicyPreview             DeveloperEcosystemValBPolicyPreviewPath              `json:"policy_preview"`
	LocalCIContinuity         DeveloperEcosystemValBLocalCIContinuityContract      `json:"local_ci_continuity"`
	DeveloperAPISDK           DeveloperEcosystemValBDeveloperAPISDKSurface         `json:"developer_api_sdk"`
	ExamplesTemplates         DeveloperEcosystemValBExamplesTemplatesContract      `json:"examples_templates"`
	APIVersioning             DeveloperEcosystemValBAPIVersioningCompatibility     `json:"api_versioning"`
	NoOverclaim               DeveloperEcosystemValBNoOverclaimDiscipline          `json:"no_overclaim"`
	EvidenceRefs              []string                                             `json:"evidence_refs,omitempty"`
	ProofSurfaceRefs          []string                                             `json:"proof_surface_refs,omitempty"`
	BlockingReasons           []string                                             `json:"blocking_reasons,omitempty"`
	ProjectionDisclaimer      string                                               `json:"projection_disclaimer"`
	CreatedAt                 string                                               `json:"created_at"`
	UpdatedAt                 string                                               `json:"updated_at"`
}

func developerEcosystemValBProjectionDisclaimer() string {
	return "projection_only not_canonical_truth developer_ecosystem_valb advisory_projection repo_sdk_integration"
}

func developerEcosystemValBHasProjectionDisclaimer(value string) bool {
	normalized := strings.TrimSpace(value)
	return strings.Contains(normalized, "projection_only") &&
		strings.Contains(normalized, "not_canonical_truth") &&
		strings.Contains(normalized, "advisory_projection") &&
		strings.Contains(normalized, "developer_ecosystem_valb")
}

func developerEcosystemValBRepoTopLevelFields() []string {
	return []string{"schema_version", "advisory", "local_validation", "ci_continuity", "templates"}
}

func developerEcosystemValBRepoAdvisoryFields() []string {
	return []string{"policy_preview", "sdk_surfaces", "inspect_explain"}
}

func developerEcosystemValBRepoLocalValidationFields() []string {
	return []string{"validation_classes", "fixture_profile", "degraded_mode"}
}

func developerEcosystemValBRepoCIContinuityFields() []string {
	return []string{"ci_descriptor", "shared_input_identity", "production_only_checks_visible"}
}

func developerEcosystemValBTemplateProfileRefs() []string {
	return []string{"starter_pack_ref", "default_config_ref", "sdk_usage_ref", "ci_continuity_ref"}
}

func developerEcosystemValBRepoDisallowedEnterpriseOverrideFields() []string {
	return []string{
		"enterprise_policy_override",
		"disable_canonical_evidence_rules",
		"grant_approval_authority",
		"change_production_trust_boundary",
	}
}

func developerEcosystemValBPolicyPreviewResults() []string {
	return []string{
		DeveloperEcosystemPolicyPreviewResultAdvisory,
		DeveloperEcosystemPolicyPreviewResultDegraded,
		DeveloperEcosystemPolicyPreviewResultUnavailable,
	}
}

func developerEcosystemValBLocalCIBehaviors() []string {
	return []string{
		DeveloperEcosystemLocalOnlyFixtureVisible,
		DeveloperEcosystemContinuityFailClosed,
	}
}

func developerEcosystemValBStableSurfaceDescriptors() []string {
	return []string{"status_reader", "proofs_reader", "inspect_explain_helper", "validation_replay_helper"}
}

func developerEcosystemValBInspectHelpers() []string {
	return []string{"local_inspect_helper", "local_explain_helper"}
}

func developerEcosystemValBValidationReplayHelpers() []string {
	return []string{"validation_replay_descriptor", "policy_preview_replay_descriptor"}
}

func developerEcosystemValBSupportedSDKVersions() []string {
	return []string{"v1_advisory", "v1_readonly"}
}

func developerEcosystemValBDeprecatedSDKVersions() []string {
	return []string{"v0_preview"}
}

func developerEcosystemValBUnsupportedSDKVersions() []string {
	return []string{"v_next_unknown"}
}

func developerEcosystemValBRequiredEvidenceScopes() []string {
	return []string{
		"point7_vale_compatibility_gate",
		"point8_developer_discipline_foundation",
		"point8_ide_local_tooling_core",
		"repo_config_schema_contract",
		"repo_config_validation_discipline",
		"policy_preview_path",
		"local_ci_continuity_contract",
		"developer_api_sdk_surface_contract",
		"examples_templates_contract",
		"api_versioning_compatibility_discipline",
		"no_overclaim_discipline",
		"canonical_evidence_boundary",
		"point8_governance",
	}
}

func DeveloperEcosystemValBProofEvidenceRefs() []string {
	return []string{
		"point7_vale_compatibility_gate",
		"point8_developer_discipline_foundation",
		"point8_ide_local_tooling_core",
		"evidence:developer-repo-config-schema-001",
		"evidence:developer-repo-config-validation-001",
		"evidence:developer-policy-preview-001",
		"evidence:developer-local-ci-continuity-001",
		"evidence:developer-api-sdk-surface-001",
		"evidence:developer-examples-templates-001",
		"evidence:developer-api-versioning-001",
		"evidence:developer-valb-no-overclaim-001",
		"evidence:developer-valb-canonical-boundary-001",
		"evidence:point8-valb-governance-001",
	}
}

func DeveloperEcosystemValBProofSurfaceRefs() []string {
	return []string{
		"/v1/verifier-ecosystem/vale/closure",
		"/v1/verifier-ecosystem/vale/proofs",
		"/v1/developer-ecosystem/val0/status",
		"/v1/developer-ecosystem/val0/proofs",
		"/v1/developer-ecosystem/vala/status",
		"/v1/developer-ecosystem/vala/proofs",
		"/v1/developer-ecosystem/valb/status",
		"/v1/developer-ecosystem/valb/proofs",
	}
}

func developerEcosystemValBEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "point7_vale_compatibility_gate", EvidenceType: "vale_compatibility", Source: "developer-ecosystem/valb/vale-compatibility", Timestamp: "2026-04-28T10:40:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point7_vale_compatibility_gate", Caveats: []string{"Val B requires patched Val E proof outputs and exact pass/no-overclaim semantics to remain intact"}},
		{EvidenceID: "point8_developer_discipline_foundation", EvidenceType: "developer_dependency", Source: "developer-ecosystem/val0", Timestamp: "2026-04-28T10:41:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point8_developer_discipline_foundation", Caveats: []string{"Val B depends on accepted developer discipline foundation and exact proof surfaces"}},
		{EvidenceID: "point8_ide_local_tooling_core", EvidenceType: "developer_dependency", Source: "developer-ecosystem/vala", Timestamp: "2026-04-28T10:42:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point8_ide_local_tooling_core", Caveats: []string{"Val B depends on accepted Val A IDE and local tooling core outputs"}},
		{EvidenceID: "evidence:developer-repo-config-schema-001", EvidenceType: "repo_config_schema", Source: "developer-ecosystem/valb/repo-config-schema", Timestamp: "2026-04-28T10:43:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "repo_config_schema_contract", Caveats: []string{"Repo-local ChangeLock-as-code schema remains scope-bound, review-bound, and advisory"}},
		{EvidenceID: "evidence:developer-repo-config-validation-001", EvidenceType: "repo_config_validation", Source: "developer-ecosystem/valb/repo-config-validation", Timestamp: "2026-04-28T10:44:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "repo_config_validation_discipline", Caveats: []string{"Repo config validation fails closed on unsupported versions, malformed config, and governance bypass fields"}},
		{EvidenceID: "evidence:developer-policy-preview-001", EvidenceType: "policy_preview", Source: "developer-ecosystem/valb/policy-preview", Timestamp: "2026-04-28T10:45:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "policy_preview_path", Caveats: []string{"Policy preview remains advisory and cannot approve deployment or become canonical decision authority"}},
		{EvidenceID: "evidence:developer-local-ci-continuity-001", EvidenceType: "local_ci_continuity", Source: "developer-ecosystem/valb/local-ci-continuity", Timestamp: "2026-04-28T10:46:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "local_ci_continuity_contract", Caveats: []string{"Local PASS-like output does not become CI PASS and CI-only checks remain visible as production-only unknowns locally"}},
		{EvidenceID: "evidence:developer-api-sdk-surface-001", EvidenceType: "developer_api_sdk", Source: "developer-ecosystem/valb/api-sdk", Timestamp: "2026-04-28T10:47:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "developer_api_sdk_surface_contract", Caveats: []string{"API and SDK surfaces remain read-only, permission-aware, advisory helpers and cannot bypass canonical evidence"}},
		{EvidenceID: "evidence:developer-examples-templates-001", EvidenceType: "examples_templates", Source: "developer-ecosystem/valb/examples", Timestamp: "2026-04-28T10:48:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "examples_templates_contract", Caveats: []string{"Examples and templates remain adoption aids and do not imply compliance certification or production approval"}},
		{EvidenceID: "evidence:developer-api-versioning-001", EvidenceType: "api_versioning", Source: "developer-ecosystem/valb/api-versioning", Timestamp: "2026-04-28T10:49:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "api_versioning_compatibility_discipline", Caveats: []string{"Deprecated versions remain visibly compatibility-bound and unsupported versions fail closed"}},
		{EvidenceID: "evidence:developer-valb-no-overclaim-001", EvidenceType: "no_overclaim", Source: "developer-ecosystem/valb/no-overclaim", Timestamp: "2026-04-28T10:50:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim_discipline", Caveats: []string{"Repo and SDK integration cannot approve deployment, certify trust, or replace enterprise governance"}},
		{EvidenceID: "evidence:developer-valb-canonical-boundary-001", EvidenceType: "canonical_boundary", Source: "developer-ecosystem/valb/canonical-boundary", Timestamp: "2026-04-28T10:51:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "canonical_evidence_boundary", Caveats: []string{"Repo config, policy preview, SDK, templates, and CI continuity remain projections over the canonical execution/audit/evidence spine"}},
		{EvidenceID: "evidence:point8-valb-governance-001", EvidenceType: "state_governance", Source: "developer-ecosystem/point8-governance", Timestamp: "2026-04-28T10:52:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point8_governance", Caveats: []string{"Val B keeps point_8_state not_complete and leaves integrated closure to later waves"}},
	}
}

func developerEcosystemValBRequiredEvidenceIDs() []string {
	ids := make([]string, 0, len(developerEcosystemValBEvidence()))
	for _, item := range developerEcosystemValBEvidence() {
		ids = append(ids, item.EvidenceID)
	}
	return ids
}

func DeveloperEcosystemValBValECompatibilityGateModel() DeveloperEcosystemValBValECompatibilityGate {
	return DeveloperEcosystemValBValECompatibilityGate{
		GateID:               "developer-ecosystem-vale-compatibility",
		Version:              "2026.04",
		ValECurrentState:     VerifierEcosystemValEStatePass,
		Point7State:          VerifierEcosystemPoint7StatePass,
		PassRuleState:        VerifierEcosystemValEPassRuleStateActive,
		NoOverclaimState:     VerifierEcosystemValENoOverclaimStateActive,
		ProofSurfaceState:    VerifierEcosystemValEProofSurfaceStateActive,
		EvidenceQualityState: VerifierEcosystemValEEvidenceQualityStateActive,
		Point7PassAllowed:    true,
		Point7PassReason:     VerifierEcosystemValEPoint7PassReasonAllowed,
		SurfaceRefs:          VerifierEcosystemValEProofSurfaceRefs(),
		EvidenceRefs:         VerifierEcosystemValEProofEvidenceRefs(),
		ProjectionDisclaimer: verifierEcosystemValEProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValBRepoConfigSchemaContractModel() DeveloperEcosystemValBRepoConfigSchemaContract {
	return DeveloperEcosystemValBRepoConfigSchemaContract{
		ContractID:                       "developer-ecosystem-repo-config-schema",
		Version:                          "2026.04",
		SchemaVersion:                    DeveloperEcosystemRepoConfigSchemaV1Advisory,
		SupportedSchemaVersions:          []string{DeveloperEcosystemRepoConfigSchemaV1Advisory},
		AllowedTopLevelFields:            developerEcosystemValBRepoTopLevelFields(),
		AllowedAdvisoryConfigFields:      developerEcosystemValBRepoAdvisoryFields(),
		AllowedLocalValidationFields:     developerEcosystemValBRepoLocalValidationFields(),
		AllowedCIContinuityFields:        developerEcosystemValBRepoCIContinuityFields(),
		AllowedTemplateProfileRefs:       developerEcosystemValBTemplateProfileRefs(),
		DisallowedEnterpriseOverrideRefs: developerEcosystemValBRepoDisallowedEnterpriseOverrideFields(),
		UnknownFieldHandling:             DeveloperEcosystemFailClosedHandling,
		UnsupportedVersionHandling:       DeveloperEcosystemFailClosedHandling,
		CompatibilityBehavior:            DeveloperEcosystemValBRepoConfigCompatibilityBehaviorBounded,
		LocalAdvisoryScope:               true,
		ScopeBounded:                     true,
		ReviewBound:                      true,
		ProjectionDisclaimer:             developerEcosystemValBProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValBRepoConfigValidationDisciplineModel() DeveloperEcosystemValBRepoConfigValidationDiscipline {
	return DeveloperEcosystemValBRepoConfigValidationDiscipline{
		DisciplineID:            "developer-ecosystem-repo-config-validation",
		Version:                 "2026.04",
		SchemaVersionValid:      true,
		FieldAllowlistValid:     true,
		NormalizedExactSetValid: true,
		EvidenceContextLinked:   true,
		AdvisoryOnly:            true,
		ProjectionDisclaimer:    developerEcosystemValBProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValBPolicyPreviewPathModel() DeveloperEcosystemValBPolicyPreviewPath {
	return DeveloperEcosystemValBPolicyPreviewPath{
		ModelID:                          "developer-ecosystem-policy-preview",
		Version:                          "2026.04",
		RepoLocalInputSnapshot:           "repo-local config snapshot for advisory preview only",
		EnterprisePolicyBoundarySnapshot: "enterprise policy boundary snapshot for local advisory comparison",
		LocalAdvisoryPreviewResult:       DeveloperEcosystemPolicyPreviewResultAdvisory,
		AffectedRuleHints:                []string{"policy.impact.hint", "preview.boundary.note"},
		ProductionOnlyUnknownsVisible:    true,
		EvidenceContextRefs:              []string{"preview:evidence:repo-local-config", "preview:evidence:enterprise-boundary"},
		RemediationHints:                 []string{"compare local advisory preview with CI policy evaluation before treating preview output as stable"},
		ProjectionDisclaimer:             developerEcosystemValBProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValBLocalCIContinuityContractModel() DeveloperEcosystemValBLocalCIContinuityContract {
	return DeveloperEcosystemValBLocalCIContinuityContract{
		ContractID:                    "developer-ecosystem-local-ci-continuity",
		Version:                       "2026.04",
		LocalValidationDescriptor:     "local advisory validation descriptor",
		CIValidationDescriptor:        "ci advisory validation descriptor",
		SharedInputIdentityFields:     []string{"repo_ref", "config_schema_version", "validation_profile"},
		ExpectedOutputClasses:         developerEcosystemVal0OutputClasses(),
		MismatchVisible:               true,
		LocalOnlyFixtureBehavior:      DeveloperEcosystemLocalOnlyFixtureVisible,
		ProductionOnlyCIChecksVisible: true,
		EvidenceHandoffRefs:           []string{"handoff:local-advisory-summary", "handoff:ci-input-identity"},
		FailureExplanationVisible:     true,
		ProjectionDisclaimer:          developerEcosystemValBProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValBDeveloperAPISDKSurfaceModel() DeveloperEcosystemValBDeveloperAPISDKSurface {
	return DeveloperEcosystemValBDeveloperAPISDKSurface{
		ContractID:                    "developer-ecosystem-api-sdk-surface",
		Version:                       "2026.04",
		StableSurfaceDescriptors:      developerEcosystemValBStableSurfaceDescriptors(),
		RequestResponseModelVersion:   DeveloperEcosystemValBAPIVersionIdentity,
		SupportedVersions:             developerEcosystemValBSupportedSDKVersions(),
		DeprecatedVersions:            developerEcosystemValBDeprecatedSDKVersions(),
		PermissionAwareOperations:     true,
		ReadOnlyAdvisoryDefault:       true,
		InspectExplainHelpers:         developerEcosystemValBInspectHelpers(),
		ValidationReplayHelpers:       developerEcosystemValBValidationReplayHelpers(),
		ErrorDegradedStateModel:       true,
		RatePerformanceExpectation:    DeveloperEcosystemSDKRateExpectationBounded,
		CompatibilityBehavior:         DeveloperEcosystemSDKCompatibilityVisible,
		FailClosedUnsupportedBehavior: true,
		DeprecatedVersionVisible:      true,
		ProjectionDisclaimer:          developerEcosystemValBProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValBExamplesTemplatesContractModel() DeveloperEcosystemValBExamplesTemplatesContract {
	return DeveloperEcosystemValBExamplesTemplatesContract{
		ContractID:                       "developer-ecosystem-examples-templates",
		Version:                          "2026.04",
		NewServiceTemplateDescriptor:     "template:new-service-advisory",
		StarterPackDescriptor:            "template:reference-aligned-starter-pack",
		DefaultConfigDescriptor:          "template:policy-compliant-default-config",
		LocalValidationExampleDescriptor: "template:local-validation-example",
		CAVIVEXExampleDescriptor:         "template:cavi-vex-context-example",
		SDKUsageExampleDescriptor:        "template:sdk-usage-example",
		CIContinuityExampleDescriptor:    "template:ci-continuity-example",
		TemplateVersion:                  "2026.04",
		FreshnessState:                   DeveloperEcosystemLocalFreshnessFresh,
		CompatibilityMetadataVisible:     true,
		DeprecatedCompatibilityVisible:   true,
		ProjectionDisclaimer:             developerEcosystemValBProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValBAPIVersioningCompatibilityModel() DeveloperEcosystemValBAPIVersioningCompatibility {
	return DeveloperEcosystemValBAPIVersioningCompatibility{
		DisciplineID:                   "developer-ecosystem-api-versioning",
		Version:                        "2026.04",
		VersionIdentity:                DeveloperEcosystemValBAPIVersionIdentity,
		SupportedVersions:              developerEcosystemValBSupportedSDKVersions(),
		DeprecatedVersions:             developerEcosystemValBDeprecatedSDKVersions(),
		UnsupportedVersions:            developerEcosystemValBUnsupportedSDKVersions(),
		CompatibilityWindow:            DeveloperEcosystemValBAPICompatibilityWindow,
		MigrationHint:                  "upgrade deprecated clients before relying on new repo or SDK advisory descriptors",
		FailClosedUnsupportedBehavior:  true,
		StaleClientBehaviorVisible:     true,
		SchemaCompatibilityBehavior:    DeveloperEcosystemSDKSchemaExactModels,
		DeprecatedCompatibilityVisible: true,
		ProjectionDisclaimer:           developerEcosystemValBProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValBNoOverclaimDisciplineModel() DeveloperEcosystemValBNoOverclaimDiscipline {
	return DeveloperEcosystemValBNoOverclaimDiscipline{
		DisciplineID:         "developer-ecosystem-valb-no-overclaim",
		Version:              "2026.04",
		ProjectionDisclaimer: developerEcosystemValBProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValBIntegrationModel() DeveloperEcosystemValBIntegration {
	return DeveloperEcosystemValBIntegration{
		IntegrationID:        "developer-ecosystem-repo-sdk-integration",
		Version:              "2026.04",
		ValECompatibility:    DeveloperEcosystemValBValECompatibilityGateModel(),
		RepoConfigSchema:     DeveloperEcosystemValBRepoConfigSchemaContractModel(),
		RepoConfigValidation: DeveloperEcosystemValBRepoConfigValidationDisciplineModel(),
		PolicyPreview:        DeveloperEcosystemValBPolicyPreviewPathModel(),
		LocalCIContinuity:    DeveloperEcosystemValBLocalCIContinuityContractModel(),
		DeveloperAPISDK:      DeveloperEcosystemValBDeveloperAPISDKSurfaceModel(),
		ExamplesTemplates:    DeveloperEcosystemValBExamplesTemplatesContractModel(),
		APIVersioning:        DeveloperEcosystemValBAPIVersioningCompatibilityModel(),
		NoOverclaim:          DeveloperEcosystemValBNoOverclaimDisciplineModel(),
		EvidenceRefs:         DeveloperEcosystemValBProofEvidenceRefs(),
		ProofSurfaceRefs:     DeveloperEcosystemValBProofSurfaceRefs(),
		ProjectionDisclaimer: developerEcosystemValBProjectionDisclaimer(),
		CreatedAt:            "2026-04-28T10:40:00Z",
		UpdatedAt:            "2026-04-28T10:40:00Z",
	}
}

func developerEcosystemValBStateSeverity(state, active, partial, incomplete, blocked, unknown string) int {
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

func EvaluateDeveloperEcosystemValBValECompatibilityState(model DeveloperEcosystemValBValECompatibilityGate) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.GateID,
		model.Version,
		model.ValECurrentState,
		model.Point7State,
		model.PassRuleState,
		model.NoOverclaimState,
		model.ProofSurfaceState,
		model.EvidenceQualityState,
		model.Point7PassReason,
		model.ProjectionDisclaimer,
	) {
		return DeveloperEcosystemValBValECompatibilityStateIncomplete
	}
	if !verifierEcosystemValEHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemValBValECompatibilityStateUnknown
	}
	if !containsExactTrimmedStringSet(model.SurfaceRefs, VerifierEcosystemValEProofSurfaceRefs()...) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, VerifierEcosystemValEProofEvidenceRefs()...) {
		return DeveloperEcosystemValBValECompatibilityStateBlocked
	}
	if verifierEcosystemValEContainsDisallowedClaim(model.Point7PassReason) {
		return DeveloperEcosystemValBValECompatibilityStateBlocked
	}
	if strings.TrimSpace(model.ValECurrentState) != VerifierEcosystemValEStatePass ||
		strings.TrimSpace(model.Point7State) != VerifierEcosystemPoint7StatePass ||
		strings.TrimSpace(model.PassRuleState) != VerifierEcosystemValEPassRuleStateActive ||
		strings.TrimSpace(model.NoOverclaimState) != VerifierEcosystemValENoOverclaimStateActive ||
		strings.TrimSpace(model.ProofSurfaceState) != VerifierEcosystemValEProofSurfaceStateActive ||
		strings.TrimSpace(model.EvidenceQualityState) != VerifierEcosystemValEEvidenceQualityStateActive ||
		!model.Point7PassAllowed ||
		strings.TrimSpace(model.Point7PassReason) != VerifierEcosystemValEPoint7PassReasonAllowed {
		return DeveloperEcosystemValBValECompatibilityStateBlocked
	}
	return DeveloperEcosystemValBValECompatibilityStateActive
}

func EvaluateDeveloperEcosystemValBDependencyState(snapshot DeveloperEcosystemValBDependencySnapshot) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		snapshot.ValACurrentState,
		snapshot.ValAPoint8State,
		snapshot.ValADependencyState,
		snapshot.IDEBaselineState,
		snapshot.TrustFeedbackState,
		snapshot.CAVIVEXContextState,
		snapshot.LocalAdvisoryState,
		snapshot.ValidationHarnessState,
		snapshot.MockVerificationState,
		snapshot.InspectExplainState,
		snapshot.DegradedModeState,
		snapshot.NoOverclaimState,
		snapshot.ValAProjectionDisclaimer,
	) {
		return DeveloperEcosystemValBDependencyStateIncomplete
	}
	if !developerEcosystemValAHasProjectionDisclaimer(snapshot.ValAProjectionDisclaimer) {
		return DeveloperEcosystemValBDependencyStateUnknown
	}
	if !containsExactTrimmedStringSet(snapshot.ValAProofSurfaceRefs, DeveloperEcosystemValAProofSurfaceRefs()...) ||
		!DeveloperEcosystemValAProofEvidenceQualityValid(developerEcosystemValAEvidence(), snapshot.ValAEvidenceRefs) {
		return DeveloperEcosystemValBDependencyStateBlocked
	}
	if strings.TrimSpace(snapshot.ValACurrentState) != DeveloperEcosystemValAStateActive ||
		strings.TrimSpace(snapshot.ValAPoint8State) != DeveloperEcosystemPoint8StateNotComplete ||
		strings.TrimSpace(snapshot.ValADependencyState) != DeveloperEcosystemValADependencyStateActive ||
		strings.TrimSpace(snapshot.IDEBaselineState) != DeveloperEcosystemValAIDEBaselineStateActive ||
		strings.TrimSpace(snapshot.TrustFeedbackState) != DeveloperEcosystemValATrustFeedbackStateActive ||
		strings.TrimSpace(snapshot.CAVIVEXContextState) != DeveloperEcosystemValACAVIVEXStateActive ||
		strings.TrimSpace(snapshot.LocalAdvisoryState) != DeveloperEcosystemValALocalAdvisoryStateActive ||
		strings.TrimSpace(snapshot.ValidationHarnessState) != DeveloperEcosystemValAValidationHarnessStateActive ||
		strings.TrimSpace(snapshot.MockVerificationState) != DeveloperEcosystemValAMockVerificationStateActive ||
		strings.TrimSpace(snapshot.InspectExplainState) != DeveloperEcosystemValAInspectExplainStateActive ||
		strings.TrimSpace(snapshot.DegradedModeState) != DeveloperEcosystemValADegradedModeStateActive ||
		strings.TrimSpace(snapshot.NoOverclaimState) != DeveloperEcosystemValANoOverclaimStateActive {
		return DeveloperEcosystemValBDependencyStateBlocked
	}
	return DeveloperEcosystemValBDependencyStateActive
}

func EvaluateDeveloperEcosystemValBRepoConfigSchemaState(model DeveloperEcosystemValBRepoConfigSchemaContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ContractID,
		model.Version,
		model.SchemaVersion,
		model.UnknownFieldHandling,
		model.UnsupportedVersionHandling,
		model.CompatibilityBehavior,
		model.ProjectionDisclaimer,
	) {
		return DeveloperEcosystemValBRepoConfigSchemaStateIncomplete
	}
	if !developerEcosystemValBHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemValBRepoConfigSchemaStateUnknown
	}
	if !containsExactTrimmedStringSet(model.SupportedSchemaVersions, DeveloperEcosystemRepoConfigSchemaV1Advisory) ||
		!containsExactTrimmedStringSet(model.AllowedTopLevelFields, developerEcosystemValBRepoTopLevelFields()...) ||
		!containsExactTrimmedStringSet(model.AllowedAdvisoryConfigFields, developerEcosystemValBRepoAdvisoryFields()...) ||
		!containsExactTrimmedStringSet(model.AllowedLocalValidationFields, developerEcosystemValBRepoLocalValidationFields()...) ||
		!containsExactTrimmedStringSet(model.AllowedCIContinuityFields, developerEcosystemValBRepoCIContinuityFields()...) ||
		!containsExactTrimmedStringSet(model.AllowedTemplateProfileRefs, developerEcosystemValBTemplateProfileRefs()...) ||
		!containsExactTrimmedStringSet(model.DisallowedEnterpriseOverrideRefs, developerEcosystemValBRepoDisallowedEnterpriseOverrideFields()...) {
		return DeveloperEcosystemValBRepoConfigSchemaStateUnknown
	}
	if strings.TrimSpace(model.SchemaVersion) != DeveloperEcosystemRepoConfigSchemaV1Advisory ||
		strings.TrimSpace(model.UnknownFieldHandling) != DeveloperEcosystemFailClosedHandling ||
		strings.TrimSpace(model.UnsupportedVersionHandling) != DeveloperEcosystemFailClosedHandling {
		return DeveloperEcosystemValBRepoConfigSchemaStateBlocked
	}
	if strings.TrimSpace(model.CompatibilityBehavior) == "permissive" ||
		strings.TrimSpace(model.CompatibilityBehavior) == "unsupported" {
		return DeveloperEcosystemValBRepoConfigSchemaStateBlocked
	}
	if strings.TrimSpace(model.CompatibilityBehavior) != DeveloperEcosystemValBRepoConfigCompatibilityBehaviorBounded {
		return DeveloperEcosystemValBRepoConfigSchemaStateUnknown
	}
	if model.EnterprisePolicyOverrideAttempt || model.DisableCanonicalEvidenceRules || model.GrantApprovalAuthority || model.ChangeProductionTrustBoundary {
		return DeveloperEcosystemValBRepoConfigSchemaStateBlocked
	}
	if !model.LocalAdvisoryScope || !model.ScopeBounded || !model.ReviewBound {
		return DeveloperEcosystemValBRepoConfigSchemaStatePartial
	}
	return DeveloperEcosystemValBRepoConfigSchemaStateActive
}

func EvaluateDeveloperEcosystemValBRepoConfigValidationState(model DeveloperEcosystemValBRepoConfigValidationDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return DeveloperEcosystemValBRepoConfigValidationStateIncomplete
	}
	if !developerEcosystemValBHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemValBRepoConfigValidationStateUnknown
	}
	if !model.SchemaVersionValid || !model.FieldAllowlistValid || model.UnsupportedValueDetected || model.MalformedConfigDetected ||
		model.MissingGovernanceField || model.PermissiveFallback || !model.AdvisoryOnly {
		return DeveloperEcosystemValBRepoConfigValidationStateBlocked
	}
	if model.DeprecatedFieldDetected && !model.DeprecatedFieldCompatibilityVisible {
		return DeveloperEcosystemValBRepoConfigValidationStateBlocked
	}
	if !model.NormalizedExactSetValid || model.DuplicateEntriesDetected || model.DeprecatedFieldDetected ||
		model.StaleTemplateReferenceDetected || !model.EvidenceContextLinked {
		return DeveloperEcosystemValBRepoConfigValidationStatePartial
	}
	return DeveloperEcosystemValBRepoConfigValidationStateActive
}

func EvaluateDeveloperEcosystemValBPolicyPreviewState(model DeveloperEcosystemValBPolicyPreviewPath) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ModelID,
		model.Version,
		model.RepoLocalInputSnapshot,
		model.EnterprisePolicyBoundarySnapshot,
		model.LocalAdvisoryPreviewResult,
		model.ProjectionDisclaimer,
	) || len(model.AffectedRuleHints) == 0 || len(model.EvidenceContextRefs) == 0 || len(model.RemediationHints) == 0 {
		return DeveloperEcosystemValBPolicyPreviewStateIncomplete
	}
	if !developerEcosystemValBHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsTrimmedString(developerEcosystemValBPolicyPreviewResults(), model.LocalAdvisoryPreviewResult) {
		return DeveloperEcosystemValBPolicyPreviewStateUnknown
	}
	if model.ApprovesDeployment || model.OverridesPolicy || model.MutatesCanonicalEvidence || model.SuppressesFailures || model.CanonicalDecisionClaim {
		return DeveloperEcosystemValBPolicyPreviewStateBlocked
	}
	if !model.ProductionOnlyUnknownsVisible {
		return DeveloperEcosystemValBPolicyPreviewStatePartial
	}
	if (strings.TrimSpace(model.LocalAdvisoryPreviewResult) == DeveloperEcosystemPolicyPreviewResultDegraded ||
		strings.TrimSpace(model.LocalAdvisoryPreviewResult) == DeveloperEcosystemPolicyPreviewResultUnavailable) &&
		strings.TrimSpace(model.MismatchDegradedReason) == "" {
		return DeveloperEcosystemValBPolicyPreviewStatePartial
	}
	return DeveloperEcosystemValBPolicyPreviewStateActive
}

func EvaluateDeveloperEcosystemValBLocalCIContinuityState(model DeveloperEcosystemValBLocalCIContinuityContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ContractID,
		model.Version,
		model.LocalValidationDescriptor,
		model.CIValidationDescriptor,
		model.LocalOnlyFixtureBehavior,
		model.ProjectionDisclaimer,
	) || len(model.SharedInputIdentityFields) == 0 || len(model.ExpectedOutputClasses) == 0 || len(model.EvidenceHandoffRefs) == 0 {
		return DeveloperEcosystemValBLocalCIContinuityStateIncomplete
	}
	if !developerEcosystemValBHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.ExpectedOutputClasses, developerEcosystemVal0OutputClasses()...) ||
		!containsTrimmedString(developerEcosystemValBLocalCIBehaviors(), model.LocalOnlyFixtureBehavior) {
		return DeveloperEcosystemValBLocalCIContinuityStateUnknown
	}
	if model.MissingDescriptors || model.LocalPassBecomesCIPass || model.ProductionCIEquivalenceClaim {
		return DeveloperEcosystemValBLocalCIContinuityStateBlocked
	}
	if !model.MismatchVisible || !model.ProductionOnlyCIChecksVisible || !model.FailureExplanationVisible {
		return DeveloperEcosystemValBLocalCIContinuityStatePartial
	}
	return DeveloperEcosystemValBLocalCIContinuityStateActive
}

func EvaluateDeveloperEcosystemValBAPISDKSurfaceState(model DeveloperEcosystemValBDeveloperAPISDKSurface) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ContractID,
		model.Version,
		model.RequestResponseModelVersion,
		model.RatePerformanceExpectation,
		model.CompatibilityBehavior,
		model.ProjectionDisclaimer,
	) || len(model.StableSurfaceDescriptors) == 0 || len(model.SupportedVersions) == 0 || len(model.InspectExplainHelpers) == 0 || len(model.ValidationReplayHelpers) == 0 {
		return DeveloperEcosystemValBAPISDKSurfaceStateIncomplete
	}
	if !developerEcosystemValBHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.StableSurfaceDescriptors, developerEcosystemValBStableSurfaceDescriptors()...) ||
		!containsExactTrimmedStringSet(model.SupportedVersions, developerEcosystemValBSupportedSDKVersions()...) ||
		!containsExactTrimmedStringSet(model.DeprecatedVersions, developerEcosystemValBDeprecatedSDKVersions()...) ||
		!containsExactTrimmedStringSet(model.InspectExplainHelpers, developerEcosystemValBInspectHelpers()...) ||
		!containsExactTrimmedStringSet(model.ValidationReplayHelpers, developerEcosystemValBValidationReplayHelpers()...) ||
		strings.TrimSpace(model.CompatibilityBehavior) != DeveloperEcosystemSDKCompatibilityVisible ||
		strings.TrimSpace(model.RatePerformanceExpectation) != DeveloperEcosystemSDKRateExpectationBounded {
		return DeveloperEcosystemValBAPISDKSurfaceStateUnknown
	}
	if model.HiddenMutationPath || model.ApprovesDeployment || model.CertifiesTrust || model.BypassesCanonicalEvidence ||
		model.SuppressesFailures || model.UnsupportedVersionActive || !model.FailClosedUnsupportedBehavior {
		return DeveloperEcosystemValBAPISDKSurfaceStateBlocked
	}
	if !model.PermissionAwareOperations || !model.ReadOnlyAdvisoryDefault || !model.ErrorDegradedStateModel || !model.DeprecatedVersionVisible {
		return DeveloperEcosystemValBAPISDKSurfaceStatePartial
	}
	return DeveloperEcosystemValBAPISDKSurfaceStateActive
}

func EvaluateDeveloperEcosystemValBExamplesTemplatesState(model DeveloperEcosystemValBExamplesTemplatesContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ContractID,
		model.Version,
		model.NewServiceTemplateDescriptor,
		model.StarterPackDescriptor,
		model.DefaultConfigDescriptor,
		model.LocalValidationExampleDescriptor,
		model.CAVIVEXExampleDescriptor,
		model.SDKUsageExampleDescriptor,
		model.CIContinuityExampleDescriptor,
		model.TemplateVersion,
		model.FreshnessState,
		model.ProjectionDisclaimer,
	) {
		return DeveloperEcosystemValBExamplesTemplatesStateIncomplete
	}
	if !developerEcosystemValBHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsTrimmedString(developerEcosystemValALocalFreshnessStates(), model.FreshnessState) {
		return DeveloperEcosystemValBExamplesTemplatesStateUnknown
	}
	if model.ComplianceCertificationClaim || model.ProductionApprovalClaim || model.CanonicalEvidenceClaim {
		return DeveloperEcosystemValBExamplesTemplatesStateBlocked
	}
	if model.DeprecatedTemplateReferenceDetected && !model.DeprecatedCompatibilityVisible {
		return DeveloperEcosystemValBExamplesTemplatesStateBlocked
	}
	if strings.TrimSpace(model.FreshnessState) == DeveloperEcosystemLocalFreshnessStale || model.StaleTemplateReferenceDetected ||
		model.DeprecatedTemplateReferenceDetected || !model.CompatibilityMetadataVisible {
		return DeveloperEcosystemValBExamplesTemplatesStatePartial
	}
	return DeveloperEcosystemValBExamplesTemplatesStateActive
}

func EvaluateDeveloperEcosystemValBAPIVersioningState(model DeveloperEcosystemValBAPIVersioningCompatibility) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.VersionIdentity,
		model.CompatibilityWindow,
		model.MigrationHint,
		model.SchemaCompatibilityBehavior,
		model.ProjectionDisclaimer,
	) || len(model.SupportedVersions) == 0 || len(model.DeprecatedVersions) == 0 || len(model.UnsupportedVersions) == 0 {
		return DeveloperEcosystemValBAPIVersioningStateIncomplete
	}
	if !developerEcosystemValBHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.SupportedVersions, developerEcosystemValBSupportedSDKVersions()...) ||
		!containsExactTrimmedStringSet(model.DeprecatedVersions, developerEcosystemValBDeprecatedSDKVersions()...) ||
		!containsExactTrimmedStringSet(model.UnsupportedVersions, developerEcosystemValBUnsupportedSDKVersions()...) ||
		strings.TrimSpace(model.VersionIdentity) != DeveloperEcosystemValBAPIVersionIdentity ||
		strings.TrimSpace(model.CompatibilityWindow) != DeveloperEcosystemValBAPICompatibilityWindow ||
		strings.TrimSpace(model.SchemaCompatibilityBehavior) != DeveloperEcosystemSDKSchemaExactModels {
		return DeveloperEcosystemValBAPIVersioningStateUnknown
	}
	if !model.FailClosedUnsupportedBehavior || model.UnsupportedVersionDetected {
		return DeveloperEcosystemValBAPIVersioningStateBlocked
	}
	if model.UnknownVersionDetected {
		return DeveloperEcosystemValBAPIVersioningStateUnknown
	}
	if model.DeprecatedVersionDetected && !model.DeprecatedCompatibilityVisible {
		return DeveloperEcosystemValBAPIVersioningStateBlocked
	}
	if model.DeprecatedVersionDetected || !model.StaleClientBehaviorVisible {
		return DeveloperEcosystemValBAPIVersioningStatePartial
	}
	return DeveloperEcosystemValBAPIVersioningStateActive
}

func EvaluateDeveloperEcosystemValBNoOverclaimState(model DeveloperEcosystemValBNoOverclaimDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return DeveloperEcosystemValBNoOverclaimStateIncomplete
	}
	if !developerEcosystemValBHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemValBNoOverclaimStateUnknown
	}
	if model.ProductionApprovalClaim || model.CertificationClaim || model.GovernanceReplacementClaim ||
		model.EnterprisePolicyOverrideClaim || model.CanonicalTruthClaim || model.ComplianceGuaranteeClaim ||
		model.DeveloperFastTrackApprovalClaim || model.TemplateFormalEvidenceClaim || model.SDKProductionAuthorizationClaim ||
		model.RepoLocalPolicyAuthorityClaim || model.PolicyPreviewCanonicalDecisionClaim || model.Point8PassClaim {
		return DeveloperEcosystemValBNoOverclaimStateBlocked
	}
	return DeveloperEcosystemValBNoOverclaimStateActive
}

func DeveloperEcosystemValBProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, evidenceRefs []string) bool {
	allFresh, stale, ok := verifierEcosystemVal0EvidenceValid(evidence)
	if !ok || !allFresh || stale || !containsExactTrimmedStringSet(evidenceRefs, DeveloperEcosystemValBProofEvidenceRefs()...) {
		return false
	}
	ids := make([]string, 0, len(evidence))
	scopes := make([]string, 0, len(evidence))
	for _, item := range evidence {
		ids = append(ids, item.EvidenceID)
		scopes = append(scopes, item.Scope)
	}
	return containsExactTrimmedStringSet(ids, developerEcosystemValBRequiredEvidenceIDs()...) &&
		containsExactTrimmedStringSet(scopes, developerEcosystemValBRequiredEvidenceScopes()...)
}

func EvaluateDeveloperEcosystemValBState(model DeveloperEcosystemValBIntegration) string {
	if EvaluateDeveloperEcosystemValBValECompatibilityState(model.ValECompatibility) != DeveloperEcosystemValBValECompatibilityStateActive {
		return DeveloperEcosystemValBStateBlocked
	}
	if EvaluateDeveloperEcosystemValBDependencyState(model.Dependency) != DeveloperEcosystemValBDependencyStateActive {
		return DeveloperEcosystemValBStateBlocked
	}
	highestSeverity := 0
	for _, severity := range []int{
		developerEcosystemValBStateSeverity(model.ValECompatibilityState, DeveloperEcosystemValBValECompatibilityStateActive, DeveloperEcosystemValBValECompatibilityStatePartial, DeveloperEcosystemValBValECompatibilityStateIncomplete, DeveloperEcosystemValBValECompatibilityStateBlocked, DeveloperEcosystemValBValECompatibilityStateUnknown),
		developerEcosystemValBStateSeverity(model.DependencyState, DeveloperEcosystemValBDependencyStateActive, DeveloperEcosystemValBDependencyStatePartial, DeveloperEcosystemValBDependencyStateIncomplete, DeveloperEcosystemValBDependencyStateBlocked, DeveloperEcosystemValBDependencyStateUnknown),
		developerEcosystemValBStateSeverity(model.RepoConfigSchemaState, DeveloperEcosystemValBRepoConfigSchemaStateActive, DeveloperEcosystemValBRepoConfigSchemaStatePartial, DeveloperEcosystemValBRepoConfigSchemaStateIncomplete, DeveloperEcosystemValBRepoConfigSchemaStateBlocked, DeveloperEcosystemValBRepoConfigSchemaStateUnknown),
		developerEcosystemValBStateSeverity(model.RepoConfigValidationState, DeveloperEcosystemValBRepoConfigValidationStateActive, DeveloperEcosystemValBRepoConfigValidationStatePartial, DeveloperEcosystemValBRepoConfigValidationStateIncomplete, DeveloperEcosystemValBRepoConfigValidationStateBlocked, DeveloperEcosystemValBRepoConfigValidationStateUnknown),
		developerEcosystemValBStateSeverity(model.PolicyPreviewState, DeveloperEcosystemValBPolicyPreviewStateActive, DeveloperEcosystemValBPolicyPreviewStatePartial, DeveloperEcosystemValBPolicyPreviewStateIncomplete, DeveloperEcosystemValBPolicyPreviewStateBlocked, DeveloperEcosystemValBPolicyPreviewStateUnknown),
		developerEcosystemValBStateSeverity(model.LocalCIContinuityState, DeveloperEcosystemValBLocalCIContinuityStateActive, DeveloperEcosystemValBLocalCIContinuityStatePartial, DeveloperEcosystemValBLocalCIContinuityStateIncomplete, DeveloperEcosystemValBLocalCIContinuityStateBlocked, DeveloperEcosystemValBLocalCIContinuityStateUnknown),
		developerEcosystemValBStateSeverity(model.APISDKSurfaceState, DeveloperEcosystemValBAPISDKSurfaceStateActive, DeveloperEcosystemValBAPISDKSurfaceStatePartial, DeveloperEcosystemValBAPISDKSurfaceStateIncomplete, DeveloperEcosystemValBAPISDKSurfaceStateBlocked, DeveloperEcosystemValBAPISDKSurfaceStateUnknown),
		developerEcosystemValBStateSeverity(model.ExamplesTemplatesState, DeveloperEcosystemValBExamplesTemplatesStateActive, DeveloperEcosystemValBExamplesTemplatesStatePartial, DeveloperEcosystemValBExamplesTemplatesStateIncomplete, DeveloperEcosystemValBExamplesTemplatesStateBlocked, DeveloperEcosystemValBExamplesTemplatesStateUnknown),
		developerEcosystemValBStateSeverity(model.APIVersioningState, DeveloperEcosystemValBAPIVersioningStateActive, DeveloperEcosystemValBAPIVersioningStatePartial, DeveloperEcosystemValBAPIVersioningStateIncomplete, DeveloperEcosystemValBAPIVersioningStateBlocked, DeveloperEcosystemValBAPIVersioningStateUnknown),
		developerEcosystemValBStateSeverity(model.NoOverclaimState, DeveloperEcosystemValBNoOverclaimStateActive, DeveloperEcosystemValBNoOverclaimStatePartial, DeveloperEcosystemValBNoOverclaimStateIncomplete, DeveloperEcosystemValBNoOverclaimStateBlocked, DeveloperEcosystemValBNoOverclaimStateUnknown),
	} {
		if severity > highestSeverity {
			highestSeverity = severity
		}
	}
	switch highestSeverity {
	case 4:
		return DeveloperEcosystemValBStateBlocked
	case 3:
		return DeveloperEcosystemValBStateUnknown
	case 2:
		return DeveloperEcosystemValBStateIncomplete
	case 1:
		return DeveloperEcosystemValBStatePartial
	default:
		return DeveloperEcosystemValBStateActive
	}
}

func EvaluateDeveloperEcosystemValBProofsState(model DeveloperEcosystemValBIntegration, limitations []string) string {
	baseState := strings.TrimSpace(model.CurrentState)
	if baseState == "" {
		baseState = DeveloperEcosystemValBStateUnknown
	}
	if !developerEcosystemValBHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, DeveloperEcosystemValBProofSurfaceRefs()...) ||
		!DeveloperEcosystemValBProofEvidenceQualityValid(developerEcosystemValBEvidence(), model.EvidenceRefs) ||
		len(limitations) == 0 ||
		strings.TrimSpace(model.Point8State) != DeveloperEcosystemPoint8StateNotComplete {
		if baseState == DeveloperEcosystemValBStateActive {
			return DeveloperEcosystemValBStatePartial
		}
		return baseState
	}
	return baseState
}

func developerEcosystemValBBlockingReasons(model DeveloperEcosystemValBIntegration) []string {
	reasons := []string{}
	if model.ValECompatibilityState != DeveloperEcosystemValBValECompatibilityStateActive {
		reasons = append(reasons, "Val B requires the patched Točka 7 Val E compatibility gate with exact Point7PassReason allowlist, active no-overclaim, and active pass rule.")
	}
	if model.DependencyState != DeveloperEcosystemValBDependencyStateActive {
		reasons = append(reasons, "Val B requires actual Val A proof/status outputs with exact proof surfaces and exact evidence quality.")
	}
	if model.RepoConfigSchemaState != DeveloperEcosystemValBRepoConfigSchemaStateActive {
		reasons = append(reasons, "Repo-local ChangeLock-as-code schema must remain schema-bound, scope-bound, review-bound, and fail-closed on unsupported fields or versions.")
	}
	if model.RepoConfigValidationState != DeveloperEcosystemValBRepoConfigValidationStateActive {
		reasons = append(reasons, "Repo config validation must fail closed on malformed, deprecated-without-compatibility, or governance-bypass configuration.")
	}
	if model.PolicyPreviewState != DeveloperEcosystemValBPolicyPreviewStateActive {
		reasons = append(reasons, "Policy preview must remain advisory and cannot override policy, approve deployment, or hide production-only unknowns.")
	}
	if model.LocalCIContinuityState != DeveloperEcosystemValBLocalCIContinuityStateActive {
		reasons = append(reasons, "Local-to-CI continuity must keep CI-only checks visible locally and must not convert local PASS-like output into CI PASS.")
	}
	if model.APISDKSurfaceState != DeveloperEcosystemValBAPISDKSurfaceStateActive {
		reasons = append(reasons, "Developer API and SDK contracts must remain read-only, permission-aware, non-mutating, non-approving, and fail-closed on unsupported versions.")
	}
	if model.ExamplesTemplatesState != DeveloperEcosystemValBExamplesTemplatesStateActive {
		reasons = append(reasons, "Examples and templates must remain bounded adoption helpers with visible freshness and compatibility metadata.")
	}
	if model.APIVersioningState != DeveloperEcosystemValBAPIVersioningStateActive {
		reasons = append(reasons, "API versioning and compatibility discipline must visibly degrade deprecated versions and fail closed on unknown or unsupported versions.")
	}
	if model.NoOverclaimState != DeveloperEcosystemValBNoOverclaimStateActive {
		reasons = append(reasons, "Val B cannot approve deployment, certify trust, replace governance, create canonical truth, or return point_8_pass.")
	}
	return verifierEcosystemValECollectText(reasons)
}

func ComputeDeveloperEcosystemValBIntegration(model DeveloperEcosystemValBIntegration) DeveloperEcosystemValBIntegration {
	model.ValECompatibilityState = EvaluateDeveloperEcosystemValBValECompatibilityState(model.ValECompatibility)
	model.DependencyState = EvaluateDeveloperEcosystemValBDependencyState(model.Dependency)
	model.RepoConfigSchemaState = EvaluateDeveloperEcosystemValBRepoConfigSchemaState(model.RepoConfigSchema)
	model.RepoConfigValidationState = EvaluateDeveloperEcosystemValBRepoConfigValidationState(model.RepoConfigValidation)
	model.PolicyPreviewState = EvaluateDeveloperEcosystemValBPolicyPreviewState(model.PolicyPreview)
	model.LocalCIContinuityState = EvaluateDeveloperEcosystemValBLocalCIContinuityState(model.LocalCIContinuity)
	model.APISDKSurfaceState = EvaluateDeveloperEcosystemValBAPISDKSurfaceState(model.DeveloperAPISDK)
	model.ExamplesTemplatesState = EvaluateDeveloperEcosystemValBExamplesTemplatesState(model.ExamplesTemplates)
	model.APIVersioningState = EvaluateDeveloperEcosystemValBAPIVersioningState(model.APIVersioning)
	model.NoOverclaimState = EvaluateDeveloperEcosystemValBNoOverclaimState(model.NoOverclaim)
	model.CurrentState = EvaluateDeveloperEcosystemValBState(model)
	model.Point8State = EvaluateDeveloperEcosystemPoint8State(model.CurrentState)
	model.BlockingReasons = developerEcosystemValBBlockingReasons(model)
	return model
}

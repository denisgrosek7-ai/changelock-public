package operability

import "strings"

const (
	DeveloperEcosystemValADependencyStateActive     = "developer_ecosystem_vala_dependency_active"
	DeveloperEcosystemValADependencyStatePartial    = "developer_ecosystem_vala_dependency_partial"
	DeveloperEcosystemValADependencyStateIncomplete = "developer_ecosystem_vala_dependency_incomplete"
	DeveloperEcosystemValADependencyStateBlocked    = "developer_ecosystem_vala_dependency_blocked"
	DeveloperEcosystemValADependencyStateUnknown    = "developer_ecosystem_vala_dependency_unknown"

	DeveloperEcosystemValAIDEBaselineStateActive     = "developer_ecosystem_vala_ide_baseline_active"
	DeveloperEcosystemValAIDEBaselineStatePartial    = "developer_ecosystem_vala_ide_baseline_partial"
	DeveloperEcosystemValAIDEBaselineStateIncomplete = "developer_ecosystem_vala_ide_baseline_incomplete"
	DeveloperEcosystemValAIDEBaselineStateBlocked    = "developer_ecosystem_vala_ide_baseline_blocked"
	DeveloperEcosystemValAIDEBaselineStateUnknown    = "developer_ecosystem_vala_ide_baseline_unknown"

	DeveloperEcosystemValATrustFeedbackStateActive     = "developer_ecosystem_vala_trust_feedback_active"
	DeveloperEcosystemValATrustFeedbackStatePartial    = "developer_ecosystem_vala_trust_feedback_partial"
	DeveloperEcosystemValATrustFeedbackStateIncomplete = "developer_ecosystem_vala_trust_feedback_incomplete"
	DeveloperEcosystemValATrustFeedbackStateBlocked    = "developer_ecosystem_vala_trust_feedback_blocked"
	DeveloperEcosystemValATrustFeedbackStateUnknown    = "developer_ecosystem_vala_trust_feedback_unknown"

	DeveloperEcosystemValACAVIVEXStateActive     = "developer_ecosystem_vala_cavi_vex_active"
	DeveloperEcosystemValACAVIVEXStatePartial    = "developer_ecosystem_vala_cavi_vex_partial"
	DeveloperEcosystemValACAVIVEXStateIncomplete = "developer_ecosystem_vala_cavi_vex_incomplete"
	DeveloperEcosystemValACAVIVEXStateBlocked    = "developer_ecosystem_vala_cavi_vex_blocked"
	DeveloperEcosystemValACAVIVEXStateUnknown    = "developer_ecosystem_vala_cavi_vex_unknown"

	DeveloperEcosystemValALocalAdvisoryStateActive     = "developer_ecosystem_vala_local_advisory_active"
	DeveloperEcosystemValALocalAdvisoryStatePartial    = "developer_ecosystem_vala_local_advisory_partial"
	DeveloperEcosystemValALocalAdvisoryStateIncomplete = "developer_ecosystem_vala_local_advisory_incomplete"
	DeveloperEcosystemValALocalAdvisoryStateBlocked    = "developer_ecosystem_vala_local_advisory_blocked"
	DeveloperEcosystemValALocalAdvisoryStateUnknown    = "developer_ecosystem_vala_local_advisory_unknown"

	DeveloperEcosystemValAValidationHarnessStateActive     = "developer_ecosystem_vala_validation_harness_active"
	DeveloperEcosystemValAValidationHarnessStatePartial    = "developer_ecosystem_vala_validation_harness_partial"
	DeveloperEcosystemValAValidationHarnessStateIncomplete = "developer_ecosystem_vala_validation_harness_incomplete"
	DeveloperEcosystemValAValidationHarnessStateBlocked    = "developer_ecosystem_vala_validation_harness_blocked"
	DeveloperEcosystemValAValidationHarnessStateUnknown    = "developer_ecosystem_vala_validation_harness_unknown"

	DeveloperEcosystemValAMockVerificationStateActive     = "developer_ecosystem_vala_mock_verification_active"
	DeveloperEcosystemValAMockVerificationStatePartial    = "developer_ecosystem_vala_mock_verification_partial"
	DeveloperEcosystemValAMockVerificationStateIncomplete = "developer_ecosystem_vala_mock_verification_incomplete"
	DeveloperEcosystemValAMockVerificationStateBlocked    = "developer_ecosystem_vala_mock_verification_blocked"
	DeveloperEcosystemValAMockVerificationStateUnknown    = "developer_ecosystem_vala_mock_verification_unknown"

	DeveloperEcosystemValAInspectExplainStateActive     = "developer_ecosystem_vala_inspect_explain_active"
	DeveloperEcosystemValAInspectExplainStatePartial    = "developer_ecosystem_vala_inspect_explain_partial"
	DeveloperEcosystemValAInspectExplainStateIncomplete = "developer_ecosystem_vala_inspect_explain_incomplete"
	DeveloperEcosystemValAInspectExplainStateBlocked    = "developer_ecosystem_vala_inspect_explain_blocked"
	DeveloperEcosystemValAInspectExplainStateUnknown    = "developer_ecosystem_vala_inspect_explain_unknown"

	DeveloperEcosystemValADegradedModeStateActive     = "developer_ecosystem_vala_degraded_mode_active"
	DeveloperEcosystemValADegradedModeStatePartial    = "developer_ecosystem_vala_degraded_mode_partial"
	DeveloperEcosystemValADegradedModeStateIncomplete = "developer_ecosystem_vala_degraded_mode_incomplete"
	DeveloperEcosystemValADegradedModeStateBlocked    = "developer_ecosystem_vala_degraded_mode_blocked"
	DeveloperEcosystemValADegradedModeStateUnknown    = "developer_ecosystem_vala_degraded_mode_unknown"

	DeveloperEcosystemValANoOverclaimStateActive     = "developer_ecosystem_vala_no_overclaim_active"
	DeveloperEcosystemValANoOverclaimStatePartial    = "developer_ecosystem_vala_no_overclaim_partial"
	DeveloperEcosystemValANoOverclaimStateIncomplete = "developer_ecosystem_vala_no_overclaim_incomplete"
	DeveloperEcosystemValANoOverclaimStateBlocked    = "developer_ecosystem_vala_no_overclaim_blocked"
	DeveloperEcosystemValANoOverclaimStateUnknown    = "developer_ecosystem_vala_no_overclaim_unknown"

	DeveloperEcosystemValAStateActive     = "developer_ecosystem_vala_active"
	DeveloperEcosystemValAStatePartial    = "developer_ecosystem_vala_partial"
	DeveloperEcosystemValAStateIncomplete = "developer_ecosystem_vala_incomplete"
	DeveloperEcosystemValAStateBlocked    = "developer_ecosystem_vala_blocked"
	DeveloperEcosystemValAStateUnknown    = "developer_ecosystem_vala_unknown"

	DeveloperEcosystemEditorVSCode   = "vscode"
	DeveloperEcosystemEditorIntelliJ = "intellij"

	DeveloperEcosystemTrustSignalDependencyTrust      = "dependency_trust_signal"
	DeveloperEcosystemTrustSignalProvenanceSeal       = "provenance_seal_presence_signal"
	DeveloperEcosystemTrustSignalRelevance            = "relevance_signal"
	DeveloperEcosystemTrustSignalPolicyImpact         = "policy_impact_hint"
	DeveloperEcosystemTrustSignalRemediationGuidance  = "remediation_guidance"
	DeveloperEcosystemTrustSignalStaleUnavailable     = "stale_unavailable_signal"
	DeveloperEcosystemTrustSignalProductionOnlyUnkown = "production_only_unknown_signal"

	DeveloperEcosystemCAVIVerdictRelevant    = "bounded_relevant"
	DeveloperEcosystemCAVIVerdictNotRelevant = "bounded_not_relevant"

	DeveloperEcosystemVEXReviewStateCandidate = "candidate"
	DeveloperEcosystemVEXReviewStateReviewed  = "reviewed"

	DeveloperEcosystemLocalFreshnessFresh   = "fresh"
	DeveloperEcosystemLocalFreshnessStale   = "stale"
	DeveloperEcosystemLocalFreshnessUnknown = "unknown"

	DeveloperEcosystemLocalAdvisoryResultAdvisory    = "local_advisory_result"
	DeveloperEcosystemLocalAdvisoryResultDegraded    = "local_advisory_degraded"
	DeveloperEcosystemLocalAdvisoryResultUnavailable = "local_advisory_unavailable"

	DeveloperEcosystemValidationClassDependencyGraph = "dependency_graph_validation"
	DeveloperEcosystemValidationClassPolicyImpact    = "policy_impact_validation"
	DeveloperEcosystemValidationClassSealPresence    = "seal_presence_validation"
	DeveloperEcosystemValidationClassCAVIVEX         = "cavi_vex_context_validation"

	DeveloperEcosystemMockFixtureTrustSynthetic = "synthetic_fixture_trust"
	DeveloperEcosystemMockFixtureTrustBounded   = "bounded_example_fixture_trust"

	DeveloperEcosystemMockStaleFixtureFailClosed     = "fail_closed"
	DeveloperEcosystemMockStaleFixtureVisibleDegrade = "visible_degraded"
)

type DeveloperEcosystemValADependencySnapshot struct {
	Val0CurrentState           string   `json:"val0_current_state"`
	Val0Point8State            string   `json:"val0_point_8_state"`
	Val0OutputClassification   string   `json:"val0_output_classification_state"`
	Val0IDEAdvisoryState       string   `json:"val0_ide_advisory_state"`
	Val0LocalProductionState   string   `json:"val0_local_production_state"`
	Val0RepoPolicyBoundary     string   `json:"val0_repo_policy_boundary_state"`
	Val0PluginSafetyState      string   `json:"val0_plugin_safety_state"`
	Val0PerformanceBudgetState string   `json:"val0_performance_budget_state"`
	Val0DXMetricsState         string   `json:"val0_dx_metrics_state"`
	Val0NoOverclaimState       string   `json:"val0_no_overclaim_state"`
	Val0ProofSurfaceRefs       []string `json:"val0_proof_surface_refs,omitempty"`
	Val0EvidenceRefs           []string `json:"val0_evidence_refs,omitempty"`
	Val0ProjectionDisclaimer   string   `json:"val0_projection_disclaimer"`
}

type DeveloperEcosystemValAIDEBaselineContract struct {
	CurrentState               string   `json:"current_state"`
	ContractID                 string   `json:"contract_id"`
	Version                    string   `json:"version"`
	SupportedEditors           []string `json:"supported_editors,omitempty"`
	AdvisoryOnlyRendering      bool     `json:"advisory_only_rendering"`
	EvidenceContextLinkSupport bool     `json:"evidence_context_link_support"`
	ReasonCodeDisplay          bool     `json:"reason_code_display"`
	FreshnessDisplay           bool     `json:"freshness_display"`
	UncertaintyDisplay         bool     `json:"uncertainty_display"`
	CandidateReviewedDisplay   bool     `json:"candidate_reviewed_display"`
	DegradedUnavailableDisplay bool     `json:"degraded_unavailable_display"`
	NonMutatingBehavior        bool     `json:"non_mutating_behavior"`
	NoApprovalCertificationUI  bool     `json:"no_approval_certification_ui"`
	CanonicalTruthClaim        bool     `json:"canonical_truth_claim"`
	ProductionApprovalClaim    bool     `json:"production_approval_claim"`
	PolicyOverrideClaim        bool     `json:"policy_override_claim"`
	DeploymentApprovalClaim    bool     `json:"deployment_approval_claim"`
	CertificationClaim         bool     `json:"certification_claim"`
	HiddenDegradedState        bool     `json:"hidden_degraded_state"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValATrustFeedbackModel struct {
	CurrentState               string   `json:"current_state"`
	ModelID                    string   `json:"model_id"`
	Version                    string   `json:"version"`
	OutputClassificationRef    string   `json:"output_classification_ref"`
	SignalClasses              []string `json:"signal_classes,omitempty"`
	OutputClasses              []string `json:"output_classes,omitempty"`
	ReasonCoded                bool     `json:"reason_coded"`
	UncertaintyVisible         bool     `json:"uncertainty_visible"`
	StalePartialVisible        bool     `json:"stale_partial_visible"`
	ProductionOnlyUnknownShown bool     `json:"production_only_unknown_shown"`
	RecommendationsSuppress    bool     `json:"recommendations_suppress_failures"`
	ImpliesPass                bool     `json:"implies_pass"`
	ApprovalClaim              bool     `json:"approval_claim"`
	CertificationClaim         bool     `json:"certification_claim"`
	CanonicalTruthClaim        bool     `json:"canonical_truth_claim"`
	DeploymentApprovalClaim    bool     `json:"deployment_approval_claim"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValACAVIVEXContextModel struct {
	CurrentState                   string   `json:"current_state"`
	ModelID                        string   `json:"model_id"`
	Version                        string   `json:"version"`
	AdvisoryIdentifier             string   `json:"advisory_identifier"`
	AffectedComponent              string   `json:"affected_component"`
	BoundedRelevanceVerdict        string   `json:"bounded_relevance_verdict"`
	ContextualExploitabilityNote   string   `json:"contextual_exploitability_note"`
	VEXReviewState                 string   `json:"vex_review_state"`
	EvidenceContextLinks           []string `json:"evidence_context_links,omitempty"`
	UncertaintyNote                string   `json:"uncertainty_note"`
	FreshnessState                 string   `json:"freshness_state"`
	ProductionOnlyUnknownsVisible  bool     `json:"production_only_unknowns_visible"`
	RemediationHint                string   `json:"remediation_hint"`
	CandidatePromotedToReviewed    bool     `json:"candidate_promoted_to_reviewed"`
	ReviewedContextEvidenceLinked  bool     `json:"reviewed_context_evidence_linked"`
	RedactionConvertsUnknownToPass bool     `json:"redaction_converts_unknown_to_pass"`
	DeploymentApprovalClaim        bool     `json:"deployment_approval_claim"`
	ProjectionDisclaimer           string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValALocalAdvisoryModel struct {
	CurrentState                 string   `json:"current_state"`
	ModelID                      string   `json:"model_id"`
	Version                      string   `json:"version"`
	LocalInputSnapshot           string   `json:"local_input_snapshot"`
	LocalEnvironmentAssumptions  []string `json:"local_environment_assumptions,omitempty"`
	LocalEvidenceRefs            []string `json:"local_evidence_refs,omitempty"`
	UnavailableProductionRefs    []string `json:"unavailable_production_refs,omitempty"`
	LocalAdvisoryResult          string   `json:"local_advisory_result"`
	DegradedUnavailableReason    string   `json:"degraded_unavailable_reason"`
	InspectExplainOutput         bool     `json:"inspect_explain_output"`
	ProductionEquivalenceClaim   bool     `json:"production_equivalence_claim"`
	MutatesCanonicalEvidence     bool     `json:"mutates_canonical_evidence"`
	ApprovesDeployment           bool     `json:"approves_deployment"`
	SuppressesProductionUnknowns bool     `json:"suppresses_production_only_unknowns"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValAValidationHarnessContract struct {
	CurrentState               string   `json:"current_state"`
	ContractID                 string   `json:"contract_id"`
	Version                    string   `json:"version"`
	SupportedValidationClasses []string `json:"supported_validation_classes,omitempty"`
	RequiredInputDescriptors   []string `json:"required_input_descriptors,omitempty"`
	LocalFixtureDescriptors    []string `json:"local_fixture_descriptors,omitempty"`
	ExpectedOutputClasses      []string `json:"expected_output_classes,omitempty"`
	MismatchVisible            bool     `json:"mismatch_visible"`
	LocalCIContinuityExpected  bool     `json:"local_ci_continuity_expected"`
	UnsupportedCaseHandling    string   `json:"unsupported_case_handling"`
	UnknownValidationClass     bool     `json:"unknown_validation_class"`
	ProductionEquivalenceClaim bool     `json:"production_equivalence_claim"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValAMockVerificationServerContract struct {
	CurrentState                 string   `json:"current_state"`
	ContractID                   string   `json:"contract_id"`
	Version                      string   `json:"version"`
	SimulationScope              string   `json:"simulation_scope"`
	SupportedProofExamples       []string `json:"supported_proof_examples,omitempty"`
	UnsupportedProofClasses      []string `json:"unsupported_proof_classes,omitempty"`
	ProductionOnlyUnknownVisible bool     `json:"production_only_unknown_visible"`
	FixtureTrustLevel            string   `json:"fixture_trust_level"`
	StaleFixtureBehavior         string   `json:"stale_fixture_behavior"`
	ProductionEquivalenceClaim   bool     `json:"production_equivalence_claim"`
	CanonicalProofClaim          bool     `json:"canonical_proof_claim"`
	ApprovalAuthority            bool     `json:"approval_authority"`
	CertificationClaim           bool     `json:"certification_claim"`
	StaleFixtureDetected         bool     `json:"stale_fixture_detected"`
	UnsupportedProofClassActive  bool     `json:"unsupported_proof_class_active"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValAInspectExplainFlow struct {
	CurrentState                   string   `json:"current_state"`
	FlowID                         string   `json:"flow_id"`
	Version                        string   `json:"version"`
	OutputClasses                  []string `json:"output_classes,omitempty"`
	FailureReasonsVisible          bool     `json:"failure_reasons_visible"`
	ProductionOnlyUnknownVisible   bool     `json:"production_only_unknown_visible"`
	RecommendationAsApproval       bool     `json:"recommendation_as_approval"`
	AdvisorySignalAsPass           bool     `json:"advisory_signal_as_pass"`
	RedactionConvertsUnknownToPass bool     `json:"redaction_converts_unknown_to_pass"`
	ProjectionDisclaimer           string   `json:"projection_disclaimer"`
}

type DeveloperEcosystemValADegradedModeDiscipline struct {
	CurrentState             string `json:"current_state"`
	DisciplineID             string `json:"discipline_id"`
	Version                  string `json:"version"`
	DegradedReasonVisible    bool   `json:"degraded_reason_visible"`
	StaleUnavailableVisible  bool   `json:"stale_unavailable_visible"`
	FallbackModeExplicit     bool   `json:"fallback_mode_explicit"`
	SilentBypassAllowed      bool   `json:"silent_bypass_allowed"`
	HiddenFailureSuppression bool   `json:"hidden_failure_suppression"`
	FalseActiveClaim         bool   `json:"false_active_claim"`
	PerformanceBudgetRef     string `json:"performance_budget_ref"`
	ProjectionDisclaimer     string `json:"projection_disclaimer"`
}

type DeveloperEcosystemValANoOverclaimDiscipline struct {
	CurrentState               string `json:"current_state"`
	DisciplineID               string `json:"discipline_id"`
	Version                    string `json:"version"`
	CanonicalTruthClaim        bool   `json:"canonical_truth_claim"`
	ProductionApprovalClaim    bool   `json:"production_approval_claim"`
	PolicyOverrideClaim        bool   `json:"policy_override_claim"`
	DeploymentApprovalClaim    bool   `json:"deployment_approval_claim"`
	CertificationClaim         bool   `json:"certification_claim"`
	Point8PassClaim            bool   `json:"point8_pass_claim"`
	EvidenceMutationClaim      bool   `json:"evidence_mutation_claim"`
	ProductionEquivalenceClaim bool   `json:"production_equivalence_claim"`
	ProjectionDisclaimer       string `json:"projection_disclaimer"`
}

type DeveloperEcosystemValACore struct {
	CurrentState           string                                               `json:"current_state"`
	Point8State            string                                               `json:"point_8_state"`
	DependencyState        string                                               `json:"dependency_state"`
	IDEBaselineState       string                                               `json:"ide_baseline_state"`
	TrustFeedbackState     string                                               `json:"trust_feedback_state"`
	CAVIVEXContextState    string                                               `json:"cavi_vex_context_state"`
	LocalAdvisoryState     string                                               `json:"local_advisory_state"`
	ValidationHarnessState string                                               `json:"local_validation_harness_state"`
	MockVerificationState  string                                               `json:"mock_verification_server_state"`
	InspectExplainState    string                                               `json:"inspect_explain_state"`
	DegradedModeState      string                                               `json:"degraded_mode_state"`
	NoOverclaimState       string                                               `json:"no_overclaim_state"`
	CoreID                 string                                               `json:"core_id"`
	Version                string                                               `json:"version"`
	Dependency             DeveloperEcosystemValADependencySnapshot             `json:"dependency"`
	IDEBaseline            DeveloperEcosystemValAIDEBaselineContract            `json:"ide_baseline"`
	TrustFeedback          DeveloperEcosystemValATrustFeedbackModel             `json:"trust_feedback"`
	CAVIVEXContext         DeveloperEcosystemValACAVIVEXContextModel            `json:"cavi_vex_context"`
	LocalAdvisory          DeveloperEcosystemValALocalAdvisoryModel             `json:"local_advisory"`
	ValidationHarness      DeveloperEcosystemValAValidationHarnessContract      `json:"validation_harness"`
	MockVerificationServer DeveloperEcosystemValAMockVerificationServerContract `json:"mock_verification_server"`
	InspectExplain         DeveloperEcosystemValAInspectExplainFlow             `json:"inspect_explain"`
	DegradedMode           DeveloperEcosystemValADegradedModeDiscipline         `json:"degraded_mode"`
	NoOverclaim            DeveloperEcosystemValANoOverclaimDiscipline          `json:"no_overclaim"`
	EvidenceRefs           []string                                             `json:"evidence_refs,omitempty"`
	ProofSurfaceRefs       []string                                             `json:"proof_surface_refs,omitempty"`
	BlockingReasons        []string                                             `json:"blocking_reasons,omitempty"`
	ProjectionDisclaimer   string                                               `json:"projection_disclaimer"`
	CreatedAt              string                                               `json:"created_at"`
	UpdatedAt              string                                               `json:"updated_at"`
}

func developerEcosystemValAProjectionDisclaimer() string {
	return "projection_only not_canonical_truth developer_ecosystem_vala advisory_projection ide_local_tooling_core"
}

func developerEcosystemValAHasProjectionDisclaimer(value string) bool {
	normalized := strings.TrimSpace(value)
	return strings.Contains(normalized, "projection_only") &&
		strings.Contains(normalized, "not_canonical_truth") &&
		strings.Contains(normalized, "advisory_projection") &&
		strings.Contains(normalized, "developer_ecosystem_vala")
}

func developerEcosystemValAEditors() []string {
	return []string{
		DeveloperEcosystemEditorVSCode,
		DeveloperEcosystemEditorIntelliJ,
	}
}

func developerEcosystemValATrustSignalClasses() []string {
	return []string{
		DeveloperEcosystemTrustSignalDependencyTrust,
		DeveloperEcosystemTrustSignalProvenanceSeal,
		DeveloperEcosystemTrustSignalRelevance,
		DeveloperEcosystemTrustSignalPolicyImpact,
		DeveloperEcosystemTrustSignalRemediationGuidance,
		DeveloperEcosystemTrustSignalStaleUnavailable,
		DeveloperEcosystemTrustSignalProductionOnlyUnkown,
	}
}

func developerEcosystemValACAVIVerdicts() []string {
	return []string{
		DeveloperEcosystemCAVIVerdictRelevant,
		DeveloperEcosystemCAVIVerdictNotRelevant,
	}
}

func developerEcosystemValAVEXReviewStates() []string {
	return []string{
		DeveloperEcosystemVEXReviewStateCandidate,
		DeveloperEcosystemVEXReviewStateReviewed,
	}
}

func developerEcosystemValALocalFreshnessStates() []string {
	return []string{
		DeveloperEcosystemLocalFreshnessFresh,
		DeveloperEcosystemLocalFreshnessStale,
		DeveloperEcosystemLocalFreshnessUnknown,
	}
}

func developerEcosystemValALocalAdvisoryResults() []string {
	return []string{
		DeveloperEcosystemLocalAdvisoryResultAdvisory,
		DeveloperEcosystemLocalAdvisoryResultDegraded,
		DeveloperEcosystemLocalAdvisoryResultUnavailable,
	}
}

func developerEcosystemValAValidationClasses() []string {
	return []string{
		DeveloperEcosystemValidationClassDependencyGraph,
		DeveloperEcosystemValidationClassPolicyImpact,
		DeveloperEcosystemValidationClassSealPresence,
		DeveloperEcosystemValidationClassCAVIVEX,
	}
}

func developerEcosystemValAMockFixtureTrustLevels() []string {
	return []string{
		DeveloperEcosystemMockFixtureTrustSynthetic,
		DeveloperEcosystemMockFixtureTrustBounded,
	}
}

func developerEcosystemValAMockStaleBehaviors() []string {
	return []string{
		DeveloperEcosystemMockStaleFixtureFailClosed,
		DeveloperEcosystemMockStaleFixtureVisibleDegrade,
	}
}

func developerEcosystemValARequiredEvidenceScopes() []string {
	return []string{
		"point8_developer_discipline_foundation",
		"point8_ide_local_tooling_core",
		"ide_baseline_contract",
		"ide_trust_feedback_model",
		"cavi_vex_context_model",
		"local_trust_advisory_model",
		"local_validation_harness_contract",
		"mock_verification_server_contract",
		"inspect_explain_flow",
		"degraded_mode_behavior",
		"no_overclaim_discipline",
		"canonical_evidence_boundary",
		"point8_governance",
	}
}

func DeveloperEcosystemValAProofEvidenceRefs() []string {
	return []string{
		"point8_developer_discipline_foundation",
		"developer_ecosystem_ide_local_tooling_core",
		"evidence:developer-ide-baseline-001",
		"evidence:developer-trust-feedback-001",
		"evidence:developer-cavi-vex-context-001",
		"evidence:developer-local-advisory-001",
		"evidence:developer-validation-harness-001",
		"evidence:developer-mock-verification-server-001",
		"evidence:developer-inspect-explain-001",
		"evidence:developer-degraded-mode-001",
		"evidence:developer-vala-no-overclaim-001",
		"evidence:developer-vala-canonical-boundary-001",
		"evidence:point8-vala-governance-001",
	}
}

func DeveloperEcosystemValAProofSurfaceRefs() []string {
	return []string{
		"/v1/developer-ecosystem/val0/status",
		"/v1/developer-ecosystem/val0/proofs",
		"/v1/developer-ecosystem/vala/status",
		"/v1/developer-ecosystem/vala/proofs",
	}
}

func developerEcosystemValAEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "point8_developer_discipline_foundation", EvidenceType: "developer_dependency", Source: "developer-ecosystem/val0", Timestamp: "2026-04-28T09:25:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point8_developer_discipline_foundation", Caveats: []string{"Val A requires active Val 0 discipline states and exact Val 0 proof surfaces"}},
		{EvidenceID: "developer_ecosystem_ide_local_tooling_core", EvidenceType: "developer_core", Source: "developer-ecosystem/vala", Timestamp: "2026-04-28T09:26:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point8_ide_local_tooling_core", Caveats: []string{"Val A remains advisory only and does not implement production IDE/runtime tooling"}},
		{EvidenceID: "evidence:developer-ide-baseline-001", EvidenceType: "ide_contract", Source: "developer-ecosystem/vala/ide-baseline", Timestamp: "2026-04-28T09:27:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "ide_baseline_contract", Caveats: []string{"VS Code and IntelliJ contracts remain advisory and non-mutating"}},
		{EvidenceID: "evidence:developer-trust-feedback-001", EvidenceType: "trust_feedback", Source: "developer-ecosystem/vala/trust-feedback", Timestamp: "2026-04-28T09:28:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "ide_trust_feedback_model", Caveats: []string{"In-editor trust feedback remains reason-coded and does not imply approval"}},
		{EvidenceID: "evidence:developer-cavi-vex-context-001", EvidenceType: "cavi_vex_context", Source: "developer-ecosystem/vala/cavi-vex", Timestamp: "2026-04-28T09:29:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "cavi_vex_context_model", Caveats: []string{"Candidate VEX remains review-aware and evidence-linked"}},
		{EvidenceID: "evidence:developer-local-advisory-001", EvidenceType: "local_advisory", Source: "developer-ecosystem/vala/local-advisory", Timestamp: "2026-04-28T09:30:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "local_trust_advisory_model", Caveats: []string{"Local advisory remains non-canonical and non-equivalent to production verification"}},
		{EvidenceID: "evidence:developer-validation-harness-001", EvidenceType: "validation_harness", Source: "developer-ecosystem/vala/validation-harness", Timestamp: "2026-04-28T09:31:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "local_validation_harness_contract", Caveats: []string{"Local harness mismatches must remain visible and fail closed on unsupported classes"}},
		{EvidenceID: "evidence:developer-mock-verification-server-001", EvidenceType: "mock_verification_server", Source: "developer-ecosystem/vala/mock-server", Timestamp: "2026-04-28T09:32:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "mock_verification_server_contract", Caveats: []string{"Mock verification is developer-assist only and cannot create canonical proof"}},
		{EvidenceID: "evidence:developer-inspect-explain-001", EvidenceType: "inspect_explain", Source: "developer-ecosystem/vala/inspect-explain", Timestamp: "2026-04-28T09:33:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "inspect_explain_flow", Caveats: []string{"Inspect/explain must preserve uncertainty, stale state, and production-only unknowns"}},
		{EvidenceID: "evidence:developer-degraded-mode-001", EvidenceType: "degraded_mode", Source: "developer-ecosystem/vala/degraded-mode", Timestamp: "2026-04-28T09:34:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "degraded_mode_behavior", Caveats: []string{"Degraded mode must remain explicit and linked to Val 0 performance budget discipline"}},
		{EvidenceID: "evidence:developer-vala-no-overclaim-001", EvidenceType: "no_overclaim", Source: "developer-ecosystem/vala/no-overclaim", Timestamp: "2026-04-28T09:35:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim_discipline", Caveats: []string{"Val A cannot return point_8_pass or production approval"}},
		{EvidenceID: "evidence:developer-vala-canonical-boundary-001", EvidenceType: "canonical_boundary", Source: "developer-ecosystem/vala/canonical-boundary", Timestamp: "2026-04-28T09:36:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "canonical_evidence_boundary", Caveats: []string{"IDE, local, harness, and mock outputs remain projections over the canonical execution/audit/evidence spine"}},
		{EvidenceID: "evidence:point8-vala-governance-001", EvidenceType: "state_governance", Source: "developer-ecosystem/point8-governance", Timestamp: "2026-04-28T09:37:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point8_governance", Caveats: []string{"Val A keeps point_8_state not_complete and leaves integrated closure to later waves"}},
	}
}

func developerEcosystemValARequiredEvidenceIDs() []string {
	ids := make([]string, 0, len(developerEcosystemValAEvidence()))
	for _, item := range developerEcosystemValAEvidence() {
		ids = append(ids, item.EvidenceID)
	}
	return ids
}

func DeveloperEcosystemValAIDEBaselineContractModel() DeveloperEcosystemValAIDEBaselineContract {
	return DeveloperEcosystemValAIDEBaselineContract{
		ContractID:                 "developer-ecosystem-ide-baseline",
		Version:                    "2026.04",
		SupportedEditors:           developerEcosystemValAEditors(),
		AdvisoryOnlyRendering:      true,
		EvidenceContextLinkSupport: true,
		ReasonCodeDisplay:          true,
		FreshnessDisplay:           true,
		UncertaintyDisplay:         true,
		CandidateReviewedDisplay:   true,
		DegradedUnavailableDisplay: true,
		NonMutatingBehavior:        true,
		NoApprovalCertificationUI:  true,
		ProjectionDisclaimer:       developerEcosystemValAProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValATrustFeedbackModelDefinition() DeveloperEcosystemValATrustFeedbackModel {
	return DeveloperEcosystemValATrustFeedbackModel{
		ModelID:                    "developer-ecosystem-ide-trust-feedback",
		Version:                    "2026.04",
		OutputClassificationRef:    DeveloperEcosystemVal0OutputClassificationModel().ClassificationID,
		SignalClasses:              developerEcosystemValATrustSignalClasses(),
		OutputClasses:              developerEcosystemVal0OutputClasses(),
		ReasonCoded:                true,
		UncertaintyVisible:         true,
		StalePartialVisible:        true,
		ProductionOnlyUnknownShown: true,
		ProjectionDisclaimer:       developerEcosystemValAProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValACAVIVEXContextModelDefinition() DeveloperEcosystemValACAVIVEXContextModel {
	return DeveloperEcosystemValACAVIVEXContextModel{
		ModelID:                       "developer-ecosystem-cavi-vex-context",
		Version:                       "2026.04",
		AdvisoryIdentifier:            "GHSA-example-2026-0001",
		AffectedComponent:             "pkg:npm/example-lib@1.2.3",
		BoundedRelevanceVerdict:       DeveloperEcosystemCAVIVerdictRelevant,
		ContextualExploitabilityNote:  "bounded local context only; production-only unknowns remain visible",
		VEXReviewState:                DeveloperEcosystemVEXReviewStateCandidate,
		EvidenceContextLinks:          []string{"context:cavi-vex:ghsa-example-2026-0001"},
		UncertaintyNote:               "candidate exploitability context remains advisory until reviewed evidence is present",
		FreshnessState:                DeveloperEcosystemLocalFreshnessFresh,
		ProductionOnlyUnknownsVisible: true,
		RemediationHint:               "inspect dependency provenance and review applicable VEX before treating local relevance as stable",
		ProjectionDisclaimer:          developerEcosystemValAProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValALocalAdvisoryModelDefinition() DeveloperEcosystemValALocalAdvisoryModel {
	return DeveloperEcosystemValALocalAdvisoryModel{
		ModelID:            "developer-ecosystem-local-advisory",
		Version:            "2026.04",
		LocalInputSnapshot: "workspace dependency graph snapshot with local fixture-backed trust context",
		LocalEnvironmentAssumptions: []string{
			"local filesystem snapshot",
			"developer workstation advisory context",
		},
		LocalEvidenceRefs: []string{
			"local:evidence:dependency-graph-snapshot",
			"local:evidence:trust-signal-cache",
		},
		UnavailableProductionRefs: []string{
			"prod:evidence:tenant-policy-snapshot",
		},
		LocalAdvisoryResult:  DeveloperEcosystemLocalAdvisoryResultAdvisory,
		InspectExplainOutput: true,
		ProjectionDisclaimer: developerEcosystemValAProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValAValidationHarnessContractModel() DeveloperEcosystemValAValidationHarnessContract {
	return DeveloperEcosystemValAValidationHarnessContract{
		ContractID:                 "developer-ecosystem-validation-harness",
		Version:                    "2026.04",
		SupportedValidationClasses: developerEcosystemValAValidationClasses(),
		RequiredInputDescriptors: []string{
			"workspace_input_snapshot",
			"local_advisory_context",
			"fixture_descriptor",
		},
		LocalFixtureDescriptors: []string{
			"fixture:dependency-graph-sample",
			"fixture:cavi-vex-sample",
		},
		ExpectedOutputClasses:     developerEcosystemVal0OutputClasses(),
		MismatchVisible:           true,
		LocalCIContinuityExpected: true,
		UnsupportedCaseHandling:   DeveloperEcosystemFailClosedHandling,
		ProjectionDisclaimer:      developerEcosystemValAProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValAMockVerificationServerContractModel() DeveloperEcosystemValAMockVerificationServerContract {
	return DeveloperEcosystemValAMockVerificationServerContract{
		ContractID:                   "developer-ecosystem-mock-verification-server",
		Version:                      "2026.04",
		SimulationScope:              DeveloperEcosystemSimulationMockVerification,
		SupportedProofExamples:       []string{"signed-provenance-sample", "sealed-vex-sample"},
		UnsupportedProofClasses:      []string{"live-production-trust-root-lookup"},
		ProductionOnlyUnknownVisible: true,
		FixtureTrustLevel:            DeveloperEcosystemMockFixtureTrustSynthetic,
		StaleFixtureBehavior:         DeveloperEcosystemMockStaleFixtureFailClosed,
		ProjectionDisclaimer:         developerEcosystemValAProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValAInspectExplainFlowModel() DeveloperEcosystemValAInspectExplainFlow {
	return DeveloperEcosystemValAInspectExplainFlow{
		FlowID:                       "developer-ecosystem-inspect-explain",
		Version:                      "2026.04",
		OutputClasses:                developerEcosystemVal0OutputClasses(),
		FailureReasonsVisible:        true,
		ProductionOnlyUnknownVisible: true,
		ProjectionDisclaimer:         developerEcosystemValAProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValADegradedModeDisciplineModel() DeveloperEcosystemValADegradedModeDiscipline {
	return DeveloperEcosystemValADegradedModeDiscipline{
		DisciplineID:            "developer-ecosystem-vala-degraded-mode",
		Version:                 "2026.04",
		DegradedReasonVisible:   true,
		StaleUnavailableVisible: true,
		FallbackModeExplicit:    true,
		PerformanceBudgetRef:    DeveloperEcosystemVal0PerformanceBudgetDisciplineModel().DisciplineID,
		ProjectionDisclaimer:    developerEcosystemValAProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValANoOverclaimDisciplineModel() DeveloperEcosystemValANoOverclaimDiscipline {
	return DeveloperEcosystemValANoOverclaimDiscipline{
		DisciplineID:         "developer-ecosystem-vala-no-overclaim",
		Version:              "2026.04",
		ProjectionDisclaimer: developerEcosystemValAProjectionDisclaimer(),
	}
}

func DeveloperEcosystemValACoreModel() DeveloperEcosystemValACore {
	return DeveloperEcosystemValACore{
		CoreID:                 "developer-ecosystem-ide-local-tooling-core",
		Version:                "2026.04",
		IDEBaseline:            DeveloperEcosystemValAIDEBaselineContractModel(),
		TrustFeedback:          DeveloperEcosystemValATrustFeedbackModelDefinition(),
		CAVIVEXContext:         DeveloperEcosystemValACAVIVEXContextModelDefinition(),
		LocalAdvisory:          DeveloperEcosystemValALocalAdvisoryModelDefinition(),
		ValidationHarness:      DeveloperEcosystemValAValidationHarnessContractModel(),
		MockVerificationServer: DeveloperEcosystemValAMockVerificationServerContractModel(),
		InspectExplain:         DeveloperEcosystemValAInspectExplainFlowModel(),
		DegradedMode:           DeveloperEcosystemValADegradedModeDisciplineModel(),
		NoOverclaim:            DeveloperEcosystemValANoOverclaimDisciplineModel(),
		EvidenceRefs:           DeveloperEcosystemValAProofEvidenceRefs(),
		ProofSurfaceRefs:       DeveloperEcosystemValAProofSurfaceRefs(),
		ProjectionDisclaimer:   developerEcosystemValAProjectionDisclaimer(),
		CreatedAt:              "2026-04-28T09:25:00Z",
		UpdatedAt:              "2026-04-28T09:25:00Z",
	}
}

func developerEcosystemValAStateSeverity(state, active, partial, incomplete, blocked, unknown string) int {
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

func EvaluateDeveloperEcosystemValADependencyState(snapshot DeveloperEcosystemValADependencySnapshot) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		snapshot.Val0CurrentState,
		snapshot.Val0Point8State,
		snapshot.Val0OutputClassification,
		snapshot.Val0IDEAdvisoryState,
		snapshot.Val0LocalProductionState,
		snapshot.Val0RepoPolicyBoundary,
		snapshot.Val0PluginSafetyState,
		snapshot.Val0PerformanceBudgetState,
		snapshot.Val0DXMetricsState,
		snapshot.Val0NoOverclaimState,
		snapshot.Val0ProjectionDisclaimer,
	) {
		return DeveloperEcosystemValADependencyStateIncomplete
	}
	if !developerEcosystemVal0HasProjectionDisclaimer(snapshot.Val0ProjectionDisclaimer) {
		return DeveloperEcosystemValADependencyStateUnknown
	}
	if !containsExactTrimmedStringSet(snapshot.Val0ProofSurfaceRefs, DeveloperEcosystemVal0ProofSurfaceRefs()...) ||
		!DeveloperEcosystemVal0ProofEvidenceQualityValid(developerEcosystemVal0Evidence(), snapshot.Val0EvidenceRefs) {
		return DeveloperEcosystemValADependencyStateBlocked
	}
	if strings.TrimSpace(snapshot.Val0CurrentState) != DeveloperEcosystemVal0StateActive ||
		strings.TrimSpace(snapshot.Val0Point8State) != DeveloperEcosystemPoint8StateNotComplete ||
		strings.TrimSpace(snapshot.Val0OutputClassification) != DeveloperEcosystemVal0OutputClassificationStateActive ||
		strings.TrimSpace(snapshot.Val0IDEAdvisoryState) != DeveloperEcosystemVal0IDEAdvisoryStateActive ||
		strings.TrimSpace(snapshot.Val0LocalProductionState) != DeveloperEcosystemVal0LocalProductionStateActive ||
		strings.TrimSpace(snapshot.Val0RepoPolicyBoundary) != DeveloperEcosystemVal0RepoPolicyStateActive ||
		strings.TrimSpace(snapshot.Val0PluginSafetyState) != DeveloperEcosystemVal0PluginSafetyStateActive ||
		strings.TrimSpace(snapshot.Val0PerformanceBudgetState) != DeveloperEcosystemVal0PerformanceBudgetStateActive ||
		strings.TrimSpace(snapshot.Val0DXMetricsState) != DeveloperEcosystemVal0DXMetricsStateActive ||
		strings.TrimSpace(snapshot.Val0NoOverclaimState) != DeveloperEcosystemVal0NoOverclaimStateActive {
		return DeveloperEcosystemValADependencyStateBlocked
	}
	return DeveloperEcosystemValADependencyStateActive
}

func EvaluateDeveloperEcosystemValAIDEBaselineState(model DeveloperEcosystemValAIDEBaselineContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.ContractID, model.Version, model.ProjectionDisclaimer) {
		return DeveloperEcosystemValAIDEBaselineStateIncomplete
	}
	if !developerEcosystemValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.SupportedEditors, developerEcosystemValAEditors()...) {
		return DeveloperEcosystemValAIDEBaselineStateUnknown
	}
	if model.CanonicalTruthClaim || model.ProductionApprovalClaim || model.PolicyOverrideClaim ||
		model.DeploymentApprovalClaim || model.CertificationClaim || model.HiddenDegradedState {
		return DeveloperEcosystemValAIDEBaselineStateBlocked
	}
	if !model.AdvisoryOnlyRendering || !model.EvidenceContextLinkSupport || !model.ReasonCodeDisplay ||
		!model.FreshnessDisplay || !model.UncertaintyDisplay || !model.CandidateReviewedDisplay ||
		!model.DegradedUnavailableDisplay || !model.NonMutatingBehavior || !model.NoApprovalCertificationUI {
		return DeveloperEcosystemValAIDEBaselineStatePartial
	}
	return DeveloperEcosystemValAIDEBaselineStateActive
}

func EvaluateDeveloperEcosystemValATrustFeedbackState(model DeveloperEcosystemValATrustFeedbackModel) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.ModelID, model.Version, model.OutputClassificationRef, model.ProjectionDisclaimer) {
		return DeveloperEcosystemValATrustFeedbackStateIncomplete
	}
	if !developerEcosystemValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.SignalClasses, developerEcosystemValATrustSignalClasses()...) ||
		!containsExactTrimmedStringSet(model.OutputClasses, developerEcosystemVal0OutputClasses()...) {
		return DeveloperEcosystemValATrustFeedbackStateUnknown
	}
	if model.OutputClassificationRef != DeveloperEcosystemVal0OutputClassificationModel().ClassificationID {
		return DeveloperEcosystemValATrustFeedbackStateUnknown
	}
	if model.RecommendationsSuppress || model.ImpliesPass || model.ApprovalClaim ||
		model.CertificationClaim || model.CanonicalTruthClaim || model.DeploymentApprovalClaim {
		return DeveloperEcosystemValATrustFeedbackStateBlocked
	}
	if !model.ReasonCoded || !model.UncertaintyVisible || !model.StalePartialVisible || !model.ProductionOnlyUnknownShown {
		return DeveloperEcosystemValATrustFeedbackStatePartial
	}
	return DeveloperEcosystemValATrustFeedbackStateActive
}

func EvaluateDeveloperEcosystemValACAVIVEXContextState(model DeveloperEcosystemValACAVIVEXContextModel) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ModelID,
		model.Version,
		model.AdvisoryIdentifier,
		model.AffectedComponent,
		model.BoundedRelevanceVerdict,
		model.VEXReviewState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceContextLinks) == 0 {
		return DeveloperEcosystemValACAVIVEXStateIncomplete
	}
	if !developerEcosystemValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsTrimmedString(developerEcosystemValACAVIVerdicts(), model.BoundedRelevanceVerdict) ||
		!containsTrimmedString(developerEcosystemValAVEXReviewStates(), model.VEXReviewState) ||
		!containsTrimmedString(developerEcosystemValALocalFreshnessStates(), model.FreshnessState) {
		return DeveloperEcosystemValACAVIVEXStateUnknown
	}
	if model.CandidatePromotedToReviewed || model.RedactionConvertsUnknownToPass || model.DeploymentApprovalClaim {
		return DeveloperEcosystemValACAVIVEXStateBlocked
	}
	if strings.TrimSpace(model.VEXReviewState) == DeveloperEcosystemVEXReviewStateReviewed && !model.ReviewedContextEvidenceLinked {
		return DeveloperEcosystemValACAVIVEXStateBlocked
	}
	if strings.TrimSpace(model.UncertaintyNote) == "" || strings.TrimSpace(model.RemediationHint) == "" || !model.ProductionOnlyUnknownsVisible {
		return DeveloperEcosystemValACAVIVEXStatePartial
	}
	return DeveloperEcosystemValACAVIVEXStateActive
}

func EvaluateDeveloperEcosystemValALocalAdvisoryState(model DeveloperEcosystemValALocalAdvisoryModel) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.ModelID, model.Version, model.LocalInputSnapshot, model.LocalAdvisoryResult, model.ProjectionDisclaimer) ||
		len(model.LocalEnvironmentAssumptions) == 0 || len(model.LocalEvidenceRefs) == 0 {
		return DeveloperEcosystemValALocalAdvisoryStateIncomplete
	}
	if !developerEcosystemValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsTrimmedString(developerEcosystemValALocalAdvisoryResults(), model.LocalAdvisoryResult) {
		return DeveloperEcosystemValALocalAdvisoryStateUnknown
	}
	if model.ProductionEquivalenceClaim || model.MutatesCanonicalEvidence || model.ApprovesDeployment || model.SuppressesProductionUnknowns {
		return DeveloperEcosystemValALocalAdvisoryStateBlocked
	}
	if len(model.UnavailableProductionRefs) == 0 || !model.InspectExplainOutput {
		return DeveloperEcosystemValALocalAdvisoryStatePartial
	}
	if (strings.TrimSpace(model.LocalAdvisoryResult) == DeveloperEcosystemLocalAdvisoryResultDegraded ||
		strings.TrimSpace(model.LocalAdvisoryResult) == DeveloperEcosystemLocalAdvisoryResultUnavailable) &&
		strings.TrimSpace(model.DegradedUnavailableReason) == "" {
		return DeveloperEcosystemValALocalAdvisoryStatePartial
	}
	return DeveloperEcosystemValALocalAdvisoryStateActive
}

func EvaluateDeveloperEcosystemValAValidationHarnessState(model DeveloperEcosystemValAValidationHarnessContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.ContractID, model.Version, model.UnsupportedCaseHandling, model.ProjectionDisclaimer) ||
		len(model.RequiredInputDescriptors) == 0 || len(model.LocalFixtureDescriptors) == 0 {
		return DeveloperEcosystemValAValidationHarnessStateIncomplete
	}
	if !developerEcosystemValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.SupportedValidationClasses, developerEcosystemValAValidationClasses()...) ||
		!containsExactTrimmedStringSet(model.ExpectedOutputClasses, developerEcosystemVal0OutputClasses()...) {
		return DeveloperEcosystemValAValidationHarnessStateUnknown
	}
	if strings.TrimSpace(model.UnsupportedCaseHandling) != DeveloperEcosystemFailClosedHandling ||
		model.UnknownValidationClass || model.ProductionEquivalenceClaim {
		return DeveloperEcosystemValAValidationHarnessStateBlocked
	}
	if !model.MismatchVisible || !model.LocalCIContinuityExpected {
		return DeveloperEcosystemValAValidationHarnessStatePartial
	}
	return DeveloperEcosystemValAValidationHarnessStateActive
}

func EvaluateDeveloperEcosystemValAMockVerificationState(model DeveloperEcosystemValAMockVerificationServerContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.ContractID, model.Version, model.SimulationScope, model.FixtureTrustLevel, model.StaleFixtureBehavior, model.ProjectionDisclaimer) ||
		len(model.SupportedProofExamples) == 0 || len(model.UnsupportedProofClasses) == 0 {
		return DeveloperEcosystemValAMockVerificationStateIncomplete
	}
	if !developerEcosystemValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		strings.TrimSpace(model.SimulationScope) != DeveloperEcosystemSimulationMockVerification ||
		!containsTrimmedString(developerEcosystemValAMockFixtureTrustLevels(), model.FixtureTrustLevel) ||
		!containsTrimmedString(developerEcosystemValAMockStaleBehaviors(), model.StaleFixtureBehavior) {
		return DeveloperEcosystemValAMockVerificationStateUnknown
	}
	if model.ProductionEquivalenceClaim || model.CanonicalProofClaim || model.ApprovalAuthority ||
		model.CertificationClaim || model.StaleFixtureDetected || model.UnsupportedProofClassActive {
		return DeveloperEcosystemValAMockVerificationStateBlocked
	}
	if !model.ProductionOnlyUnknownVisible {
		return DeveloperEcosystemValAMockVerificationStatePartial
	}
	return DeveloperEcosystemValAMockVerificationStateActive
}

func EvaluateDeveloperEcosystemValAInspectExplainState(model DeveloperEcosystemValAInspectExplainFlow) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.FlowID, model.Version, model.ProjectionDisclaimer) {
		return DeveloperEcosystemValAInspectExplainStateIncomplete
	}
	if !developerEcosystemValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.OutputClasses, developerEcosystemVal0OutputClasses()...) {
		return DeveloperEcosystemValAInspectExplainStateUnknown
	}
	if !model.FailureReasonsVisible || !model.ProductionOnlyUnknownVisible ||
		model.RecommendationAsApproval || model.AdvisorySignalAsPass || model.RedactionConvertsUnknownToPass {
		return DeveloperEcosystemValAInspectExplainStateBlocked
	}
	return DeveloperEcosystemValAInspectExplainStateActive
}

func EvaluateDeveloperEcosystemValADegradedModeState(model DeveloperEcosystemValADegradedModeDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.PerformanceBudgetRef, model.ProjectionDisclaimer) {
		return DeveloperEcosystemValADegradedModeStateIncomplete
	}
	if !developerEcosystemValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		strings.TrimSpace(model.PerformanceBudgetRef) != DeveloperEcosystemVal0PerformanceBudgetDisciplineModel().DisciplineID {
		return DeveloperEcosystemValADegradedModeStateUnknown
	}
	if model.SilentBypassAllowed || model.HiddenFailureSuppression || model.FalseActiveClaim {
		return DeveloperEcosystemValADegradedModeStateBlocked
	}
	if !model.DegradedReasonVisible || !model.StaleUnavailableVisible || !model.FallbackModeExplicit {
		return DeveloperEcosystemValADegradedModeStatePartial
	}
	return DeveloperEcosystemValADegradedModeStateActive
}

func EvaluateDeveloperEcosystemValANoOverclaimState(model DeveloperEcosystemValANoOverclaimDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return DeveloperEcosystemValANoOverclaimStateIncomplete
	}
	if !developerEcosystemValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return DeveloperEcosystemValANoOverclaimStateUnknown
	}
	if model.CanonicalTruthClaim || model.ProductionApprovalClaim || model.PolicyOverrideClaim ||
		model.DeploymentApprovalClaim || model.CertificationClaim || model.Point8PassClaim ||
		model.EvidenceMutationClaim || model.ProductionEquivalenceClaim {
		return DeveloperEcosystemValANoOverclaimStateBlocked
	}
	return DeveloperEcosystemValANoOverclaimStateActive
}

func DeveloperEcosystemValAProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, evidenceRefs []string) bool {
	allFresh, stale, ok := verifierEcosystemVal0EvidenceValid(evidence)
	if !ok || !allFresh || stale || !containsExactTrimmedStringSet(evidenceRefs, DeveloperEcosystemValAProofEvidenceRefs()...) {
		return false
	}
	ids := make([]string, 0, len(evidence))
	scopes := make([]string, 0, len(evidence))
	for _, item := range evidence {
		ids = append(ids, item.EvidenceID)
		scopes = append(scopes, item.Scope)
	}
	return containsExactTrimmedStringSet(ids, developerEcosystemValARequiredEvidenceIDs()...) &&
		containsExactTrimmedStringSet(scopes, developerEcosystemValARequiredEvidenceScopes()...)
}

func EvaluateDeveloperEcosystemValAState(model DeveloperEcosystemValACore) string {
	if EvaluateDeveloperEcosystemValADependencyState(model.Dependency) != DeveloperEcosystemValADependencyStateActive {
		return DeveloperEcosystemValAStateBlocked
	}
	highestSeverity := 0
	for _, severity := range []int{
		developerEcosystemValAStateSeverity(model.DependencyState, DeveloperEcosystemValADependencyStateActive, DeveloperEcosystemValADependencyStatePartial, DeveloperEcosystemValADependencyStateIncomplete, DeveloperEcosystemValADependencyStateBlocked, DeveloperEcosystemValADependencyStateUnknown),
		developerEcosystemValAStateSeverity(model.IDEBaselineState, DeveloperEcosystemValAIDEBaselineStateActive, DeveloperEcosystemValAIDEBaselineStatePartial, DeveloperEcosystemValAIDEBaselineStateIncomplete, DeveloperEcosystemValAIDEBaselineStateBlocked, DeveloperEcosystemValAIDEBaselineStateUnknown),
		developerEcosystemValAStateSeverity(model.TrustFeedbackState, DeveloperEcosystemValATrustFeedbackStateActive, DeveloperEcosystemValATrustFeedbackStatePartial, DeveloperEcosystemValATrustFeedbackStateIncomplete, DeveloperEcosystemValATrustFeedbackStateBlocked, DeveloperEcosystemValATrustFeedbackStateUnknown),
		developerEcosystemValAStateSeverity(model.CAVIVEXContextState, DeveloperEcosystemValACAVIVEXStateActive, DeveloperEcosystemValACAVIVEXStatePartial, DeveloperEcosystemValACAVIVEXStateIncomplete, DeveloperEcosystemValACAVIVEXStateBlocked, DeveloperEcosystemValACAVIVEXStateUnknown),
		developerEcosystemValAStateSeverity(model.LocalAdvisoryState, DeveloperEcosystemValALocalAdvisoryStateActive, DeveloperEcosystemValALocalAdvisoryStatePartial, DeveloperEcosystemValALocalAdvisoryStateIncomplete, DeveloperEcosystemValALocalAdvisoryStateBlocked, DeveloperEcosystemValALocalAdvisoryStateUnknown),
		developerEcosystemValAStateSeverity(model.ValidationHarnessState, DeveloperEcosystemValAValidationHarnessStateActive, DeveloperEcosystemValAValidationHarnessStatePartial, DeveloperEcosystemValAValidationHarnessStateIncomplete, DeveloperEcosystemValAValidationHarnessStateBlocked, DeveloperEcosystemValAValidationHarnessStateUnknown),
		developerEcosystemValAStateSeverity(model.MockVerificationState, DeveloperEcosystemValAMockVerificationStateActive, DeveloperEcosystemValAMockVerificationStatePartial, DeveloperEcosystemValAMockVerificationStateIncomplete, DeveloperEcosystemValAMockVerificationStateBlocked, DeveloperEcosystemValAMockVerificationStateUnknown),
		developerEcosystemValAStateSeverity(model.InspectExplainState, DeveloperEcosystemValAInspectExplainStateActive, DeveloperEcosystemValAInspectExplainStatePartial, DeveloperEcosystemValAInspectExplainStateIncomplete, DeveloperEcosystemValAInspectExplainStateBlocked, DeveloperEcosystemValAInspectExplainStateUnknown),
		developerEcosystemValAStateSeverity(model.DegradedModeState, DeveloperEcosystemValADegradedModeStateActive, DeveloperEcosystemValADegradedModeStatePartial, DeveloperEcosystemValADegradedModeStateIncomplete, DeveloperEcosystemValADegradedModeStateBlocked, DeveloperEcosystemValADegradedModeStateUnknown),
		developerEcosystemValAStateSeverity(model.NoOverclaimState, DeveloperEcosystemValANoOverclaimStateActive, DeveloperEcosystemValANoOverclaimStatePartial, DeveloperEcosystemValANoOverclaimStateIncomplete, DeveloperEcosystemValANoOverclaimStateBlocked, DeveloperEcosystemValANoOverclaimStateUnknown),
	} {
		if severity > highestSeverity {
			highestSeverity = severity
		}
	}
	switch highestSeverity {
	case 4:
		return DeveloperEcosystemValAStateBlocked
	case 3:
		return DeveloperEcosystemValAStateUnknown
	case 2:
		return DeveloperEcosystemValAStateIncomplete
	case 1:
		return DeveloperEcosystemValAStatePartial
	default:
		return DeveloperEcosystemValAStateActive
	}
}

func EvaluateDeveloperEcosystemValAProofsState(model DeveloperEcosystemValACore, limitations []string) string {
	baseState := strings.TrimSpace(model.CurrentState)
	if baseState == "" {
		baseState = DeveloperEcosystemValAStateUnknown
	}
	if !developerEcosystemValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, DeveloperEcosystemValAProofSurfaceRefs()...) ||
		!DeveloperEcosystemValAProofEvidenceQualityValid(developerEcosystemValAEvidence(), model.EvidenceRefs) ||
		len(limitations) == 0 ||
		strings.TrimSpace(model.Point8State) != DeveloperEcosystemPoint8StateNotComplete {
		if baseState == DeveloperEcosystemValAStateActive {
			return DeveloperEcosystemValAStatePartial
		}
		return baseState
	}
	return baseState
}

func computeDeveloperEcosystemValABlockingReasons(model DeveloperEcosystemValACore) []string {
	reasons := []string{}
	if model.DependencyState != DeveloperEcosystemValADependencyStateActive {
		reasons = append(reasons, "Točka 8 Val A requires actual Val 0 proof/status outputs with exact proof surfaces and evidence quality.")
	}
	if model.IDEBaselineState != DeveloperEcosystemValAIDEBaselineStateActive {
		reasons = append(reasons, "IDE baseline contracts must remain advisory-only, evidence-linked, freshness-aware, uncertainty-aware, and non-approving.")
	}
	if model.TrustFeedbackState != DeveloperEcosystemValATrustFeedbackStateActive {
		reasons = append(reasons, "In-editor trust feedback must remain classified, reason-coded, and must preserve stale, partial, and production-only unknown signals.")
	}
	if model.CAVIVEXContextState != DeveloperEcosystemValACAVIVEXStateActive {
		reasons = append(reasons, "CAVI/VEX context must remain bounded, review-aware, evidence-linked, and must not convert candidate or unknown context into reviewed or verified state.")
	}
	if model.LocalAdvisoryState != DeveloperEcosystemValALocalAdvisoryStateActive {
		reasons = append(reasons, "Local advisory output must remain non-canonical, non-mutating, non-approving, and explicit about unavailable production refs.")
	}
	if model.ValidationHarnessState != DeveloperEcosystemValAValidationHarnessStateActive {
		reasons = append(reasons, "Local validation harness integration must fail closed on unsupported or unknown validation classes and cannot claim production verifier equivalence.")
	}
	if model.MockVerificationState != DeveloperEcosystemValAMockVerificationStateActive {
		reasons = append(reasons, "Mock verification server output must remain developer-assist only and must not create canonical proof or hide stale fixtures.")
	}
	if model.InspectExplainState != DeveloperEcosystemValAInspectExplainStateActive {
		reasons = append(reasons, "Inspect/explain flow must keep failure reasons, uncertainty, stale state, and production-only unknowns visible.")
	}
	if model.DegradedModeState != DeveloperEcosystemValADegradedModeStateActive {
		reasons = append(reasons, "Developer degraded mode must remain explicit, linked to performance budget discipline, and must not silently bypass failures.")
	}
	if model.NoOverclaimState != DeveloperEcosystemValANoOverclaimStateActive {
		reasons = append(reasons, "Val A cannot claim point_8_pass, deployment approval, certification, policy override, or canonical truth.")
	}
	return verifierEcosystemValECollectText(reasons)
}

func ComputeDeveloperEcosystemValACore(model DeveloperEcosystemValACore) DeveloperEcosystemValACore {
	model.DependencyState = EvaluateDeveloperEcosystemValADependencyState(model.Dependency)
	model.IDEBaselineState = EvaluateDeveloperEcosystemValAIDEBaselineState(model.IDEBaseline)
	model.TrustFeedbackState = EvaluateDeveloperEcosystemValATrustFeedbackState(model.TrustFeedback)
	model.CAVIVEXContextState = EvaluateDeveloperEcosystemValACAVIVEXContextState(model.CAVIVEXContext)
	model.LocalAdvisoryState = EvaluateDeveloperEcosystemValALocalAdvisoryState(model.LocalAdvisory)
	model.ValidationHarnessState = EvaluateDeveloperEcosystemValAValidationHarnessState(model.ValidationHarness)
	model.MockVerificationState = EvaluateDeveloperEcosystemValAMockVerificationState(model.MockVerificationServer)
	model.InspectExplainState = EvaluateDeveloperEcosystemValAInspectExplainState(model.InspectExplain)
	model.DegradedModeState = EvaluateDeveloperEcosystemValADegradedModeState(model.DegradedMode)
	model.NoOverclaimState = EvaluateDeveloperEcosystemValANoOverclaimState(model.NoOverclaim)
	model.CurrentState = EvaluateDeveloperEcosystemValAState(model)
	model.Point8State = EvaluateDeveloperEcosystemPoint8State(model.CurrentState)
	model.BlockingReasons = computeDeveloperEcosystemValABlockingReasons(model)
	return model
}

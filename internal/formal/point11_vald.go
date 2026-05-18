package formal

import (
	"strings"
	"time"
)

const (
	Point11ValDStateActive         = "point11_vald_final_governance_authority_gate_active"
	Point11ValDStateBlocked        = "point11_vald_final_governance_authority_gate_blocked"
	Point11ValDStateReviewRequired = "point11_vald_final_governance_authority_gate_review_required"

	Point11ValDDependencyStateActive         = "point11_vald_dependency_active"
	Point11ValDDependencyStateBlocked        = "point11_vald_dependency_blocked"
	Point11ValDDependencyStateReviewRequired = "point11_vald_dependency_review_required"

	Point11ValDIntegratedInvariantStateActive  = "point11_vald_integrated_invariant_active"
	Point11ValDIntegratedInvariantStateBlocked = "point11_vald_integrated_invariant_blocked"

	Point11ValDQualityMapStateActive  = "point11_vald_quality_map_active"
	Point11ValDQualityMapStateBlocked = "point11_vald_quality_map_blocked"

	Point11ValDPublicationReviewStateActive  = "point11_vald_publication_review_active"
	Point11ValDPublicationReviewStateBlocked = "point11_vald_publication_review_blocked"

	Point11ValDNoOverclaimReviewStateActive  = "point11_vald_no_overclaim_review_active"
	Point11ValDNoOverclaimReviewStateBlocked = "point11_vald_no_overclaim_review_blocked"

	Point11ValDCleanRoomIPReviewStateActive  = "point11_vald_clean_room_ip_review_active"
	Point11ValDCleanRoomIPReviewStateBlocked = "point11_vald_clean_room_ip_review_blocked"

	Point11ValDCLBClosureStateActive  = "point11_vald_clb_closure_active"
	Point11ValDCLBClosureStateBlocked = "point11_vald_clb_closure_blocked"

	Point11ValDPassClosureManifestStateActive  = "point11_vald_pass_closure_manifest_active"
	Point11ValDPassClosureManifestStateBlocked = "point11_vald_pass_closure_manifest_blocked"

	Point11ValDFinalPassGateStateActive  = "point11_vald_final_pass_gate_active"
	Point11ValDFinalPassGateStateBlocked = "point11_vald_final_pass_gate_blocked"
)

const (
	point11ValDProjectionDisclaimerBaseline   = "projection_only not_canonical_truth point11_vald_final_governance_authority_gate"
	point11ValDPointID                        = "point_11"
	point11ValDWaveID                         = "val_d"
	point11ValDScope                          = "final_governance_authority_closure_gate"
	point11ValDReviewerResultPassConfirmed    = "PASS_CONFIRMED"
	point11ValDReviewerResultPass             = "PASS"
	point11ValDPoint11PassToken               = "point_11_pass"
	point11ValDCheckStateActive               = "check_active"
	point11ValDCheckStateBlocked              = "check_blocked"
	point11ValDCheckStateReviewRequired       = "check_review_required"
	point11ValDFreshnessStateActive           = "freshness_active"
	point11ValDFreshnessStateStale            = "freshness_stale"
	point11ValDRevocationStateActive          = "revocation_active"
	point11ValDRevocationStateBlocked         = "revocation_blocked"
	point11ValDSupersessionStateActive        = "supersession_active"
	point11ValDSupersessionStateBlocked       = "supersession_blocked"
	point11ValDDuplicateStateActive           = "duplicate_none"
	point11ValDDuplicateStateBlocked          = "duplicate_detected"
	point11ValDConflictStateActive            = "conflict_none"
	point11ValDConflictStateBlocked           = "conflict_detected"
	point11ValDUnrelatedStateActive           = "unrelated_none"
	point11ValDUnrelatedStateBlocked          = "unrelated_detected"
	point11ValDTenantScopeStateActive         = "tenant_scope_active"
	point11ValDTenantScopeStateBlocked        = "tenant_scope_blocked"
	point11ValDProjectionBoundaryResultActive = "projection_boundary_active"
	point11ValDCommandRunRefPrefix            = "command_run_"
	point11ValDTestRunRefPrefix               = "test_run_"
	point11ValDGrepRunRefPrefix               = "grep_run_"
	point11ValDNegativeFixtureRunRefPrefix    = "negative_fixture_"
	point11ValDPublicationSurfaceAgentOutput  = "agent_output"
	point11ValDDependencyRefPrefix            = "dependency_review_"
	point11ValDThirdPartyRefPrefix            = "third_party_"
	point11ValDLicenseReviewRefPrefix         = "license_review_"
	point11ValDIPReviewRefPrefix              = "ip_review_"
	point11ValDExternalLegalReviewRefPrefix   = "external_legal_review_"
	point11ValDExternalFTOReviewRefPrefix     = "external_fto_review_"
	point11ValDFindingRefPrefix               = "finding_"
	point11ValDAcceptedRiskRefPrefix          = "accepted_risk_"
	point11ValDDeferredItemRefPrefix          = "deferred_item_"
	point11ValDResolvedFindingRefPrefix       = "resolved_finding_"
)

type Point11ValDVal0ReviewContext struct {
	Val0Point11PassEmitted bool     `json:"val0_point11_pass_emitted"`
	ReviewPrerequisites    []string `json:"review_prerequisites,omitempty"`
}

type Point11ValDValAReviewContext struct {
	ValAPoint11PassEmitted            bool     `json:"vala_point11_pass_emitted"`
	ValACreatesSigningSideEffects     bool     `json:"vala_creates_signing_side_effects"`
	ValACreatesAnchoringSideEffects   bool     `json:"vala_creates_anchoring_side_effects"`
	ValACreatesExternalAPISideEffects bool     `json:"vala_creates_external_api_side_effects"`
	ValACreatesProductionSideEffects  bool     `json:"vala_creates_production_side_effects"`
	OpenCLB0Findings                  int      `json:"open_clb0_findings"`
	OpenCLB1Findings                  int      `json:"open_clb1_findings"`
	OpenCLB2Findings                  int      `json:"open_clb2_findings"`
	ReviewPrerequisites               []string `json:"review_prerequisites,omitempty"`
}

type Point11ValDValBReviewContext struct {
	ValBPoint11PassEmitted            bool     `json:"valb_point11_pass_emitted"`
	ValBCreatesSigningSideEffects     bool     `json:"valb_creates_signing_side_effects"`
	ValBCreatesAnchoringSideEffects   bool     `json:"valb_creates_anchoring_side_effects"`
	ValBCreatesExternalAPISideEffects bool     `json:"valb_creates_external_api_side_effects"`
	ValBCreatesProductionSideEffects  bool     `json:"valb_creates_production_side_effects"`
	OpenCLB0Findings                  int      `json:"open_clb0_findings"`
	OpenCLB1Findings                  int      `json:"open_clb1_findings"`
	OpenCLB2Findings                  int      `json:"open_clb2_findings"`
	ReviewPrerequisites               []string `json:"review_prerequisites,omitempty"`
}

type Point11ValDValCReviewContext struct {
	ValCPoint11PassEmitted            bool     `json:"valc_point11_pass_emitted"`
	ValCCreatesExternalAPISideEffects bool     `json:"valc_creates_external_api_side_effects"`
	ValCCreatesProductionSideEffects  bool     `json:"valc_creates_production_side_effects"`
	ValCDashboardSourceOfTruthBypass  bool     `json:"valc_dashboard_source_of_truth_bypass"`
	OpenCLB0Findings                  int      `json:"open_clb0_findings"`
	OpenCLB1Findings                  int      `json:"open_clb1_findings"`
	OpenCLB2Findings                  int      `json:"open_clb2_findings"`
	ReviewPrerequisites               []string `json:"review_prerequisites,omitempty"`
}

type Point11ValDVal0DependencySnapshot struct {
	CurrentState                  string   `json:"current_state"`
	DependencyState               string   `json:"dependency_state"`
	PolicyContractState           string   `json:"policy_contract_state"`
	ClaimGovernanceState          string   `json:"claim_governance_state"`
	AuthorityMatrixState          string   `json:"authority_matrix_state"`
	ExceptionGovernanceState      string   `json:"exception_governance_state"`
	ABACState                     string   `json:"abac_state"`
	DecisionBindingState          string   `json:"decision_binding_state"`
	NoOverclaimState              string   `json:"no_overclaim_state"`
	CrossDomainCompatibilityState string   `json:"cross_domain_compatibility_state"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	Point11PassEmitted            bool     `json:"point11_pass_emitted"`
	ReviewPrerequisites           []string `json:"review_prerequisites,omitempty"`
}

type Point11ValDValADependencySnapshot struct {
	CurrentState                             string   `json:"current_state"`
	DependencyState                          string   `json:"dependency_state"`
	RegistryState                            string   `json:"registry_state"`
	SignatureState                           string   `json:"signature_state"`
	AnchorState                              string   `json:"anchor_state"`
	LifecycleTransitionState                 string   `json:"lifecycle_transition_state"`
	PolicyUseState                           string   `json:"policy_use_state"`
	GraphState                               string   `json:"graph_state"`
	ProjectionDisclaimer                     string   `json:"projection_disclaimer"`
	CreatesLegalRegulatoryCertificationClaim bool     `json:"creates_legal_regulatory_certification_claim"`
	CreatesPublicationSideEffects            bool     `json:"creates_publication_side_effects"`
	Point11PassEmitted                       bool     `json:"point11_pass_emitted"`
	CreatesSigningSideEffects                bool     `json:"creates_signing_side_effects"`
	CreatesAnchoringSideEffects              bool     `json:"creates_anchoring_side_effects"`
	CreatesExternalAPISideEffects            bool     `json:"creates_external_api_side_effects"`
	CreatesProductionSideEffects             bool     `json:"creates_production_side_effects"`
	OpenCLB0Findings                         int      `json:"open_clb0_findings"`
	OpenCLB1Findings                         int      `json:"open_clb1_findings"`
	OpenCLB2Findings                         int      `json:"open_clb2_findings"`
	ReviewPrerequisites                      []string `json:"review_prerequisites,omitempty"`
}

type Point11ValDValBDependencySnapshot struct {
	CurrentState                  string   `json:"current_state"`
	DependencyState               string   `json:"dependency_state"`
	ClaimTypeState                string   `json:"claim_type_state"`
	IssuanceRequestState          string   `json:"issuance_request_state"`
	IssuedClaimState              string   `json:"issued_claim_state"`
	RegistryState                 string   `json:"registry_state"`
	VerificationState             string   `json:"verification_state"`
	CrossDomainIntakeState        string   `json:"cross_domain_intake_state"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	CreatesAuthorityClaims        bool     `json:"creates_authority_claims"`
	CreatesPublicationSideEffects bool     `json:"creates_publication_side_effects"`
	Point11PassEmitted            bool     `json:"point11_pass_emitted"`
	CreatesSigningSideEffects     bool     `json:"creates_signing_side_effects"`
	CreatesAnchoringSideEffects   bool     `json:"creates_anchoring_side_effects"`
	CreatesExternalAPISideEffects bool     `json:"creates_external_api_side_effects"`
	CreatesProductionSideEffects  bool     `json:"creates_production_side_effects"`
	OpenCLB0Findings              int      `json:"open_clb0_findings"`
	OpenCLB1Findings              int      `json:"open_clb1_findings"`
	OpenCLB2Findings              int      `json:"open_clb2_findings"`
	ReviewPrerequisites           []string `json:"review_prerequisites,omitempty"`
}

type Point11ValDValCDependencySnapshot struct {
	CurrentState                      string   `json:"current_state"`
	DependencyState                   string   `json:"dependency_state"`
	EnforcementInputState             string   `json:"enforcement_input_state"`
	EnforcementResultState            string   `json:"enforcement_result_state"`
	ABACDecisionState                 string   `json:"abac_decision_state"`
	ExceptionDecisionState            string   `json:"exception_decision_state"`
	PrecedenceState                   string   `json:"precedence_state"`
	MonitoringState                   string   `json:"monitoring_state"`
	DashboardState                    string   `json:"dashboard_state"`
	ProjectionDisclaimer              string   `json:"projection_disclaimer"`
	CreatesAuthorityClaims            bool     `json:"creates_authority_claims"`
	CreatesPublicationSideEffects     bool     `json:"creates_publication_side_effects"`
	CreatesRealEnforcementSideEffects bool     `json:"creates_real_enforcement_side_effects"`
	Point11PassEmitted                bool     `json:"point11_pass_emitted"`
	CreatesExternalAPISideEffects     bool     `json:"creates_external_api_side_effects"`
	CreatesProductionSideEffects      bool     `json:"creates_production_side_effects"`
	DashboardSourceOfTruthBypass      bool     `json:"dashboard_source_of_truth_bypass"`
	OpenCLB0Findings                  int      `json:"open_clb0_findings"`
	OpenCLB1Findings                  int      `json:"open_clb1_findings"`
	OpenCLB2Findings                  int      `json:"open_clb2_findings"`
	ReviewPrerequisites               []string `json:"review_prerequisites,omitempty"`
}

type Point11ValDDependencyBundle struct {
	Val0 Point11ValDVal0DependencySnapshot `json:"val0"`
	ValA Point11ValDValADependencySnapshot `json:"vala"`
	ValB Point11ValDValBDependencySnapshot `json:"valb"`
	ValC Point11ValDValCDependencySnapshot `json:"valc"`
}

type Point11ValDIntegratedGovernanceInvariantReview struct {
	CurrentState                       string   `json:"current_state"`
	InvariantReviewID                  string   `json:"invariant_review_id"`
	Val0Ref                            string   `json:"val0_ref"`
	ValARef                            string   `json:"vala_ref"`
	ValBRef                            string   `json:"valb_ref"`
	ValCRef                            string   `json:"valc_ref"`
	PolicyAuthorityConsistencyState    string   `json:"policy_authority_consistency_state"`
	ClaimAuthorityConsistencyState     string   `json:"claim_authority_consistency_state"`
	PublicationBoundaryState           string   `json:"publication_boundary_state"`
	NoOverclaimState                   string   `json:"no_overclaim_state"`
	CleanRoomIPState                   string   `json:"clean_room_ip_state"`
	ProjectionBoundaryState            string   `json:"projection_boundary_state"`
	ExceptionEmergencyConsistencyState string   `json:"exception_emergency_consistency_state"`
	ABACEnforcementConsistencyState    string   `json:"abac_enforcement_consistency_state"`
	DashboardProjectionState           string   `json:"dashboard_projection_state"`
	Point11PassRuleState               string   `json:"point11_pass_rule_state"`
	Diagnostics                        []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer               string   `json:"projection_disclaimer"`
}

type Point11ValDEvidenceGovernanceQualityMap struct {
	CurrentState          string   `json:"current_state"`
	QualityMapID          string   `json:"quality_map_id"`
	PolicyRefs            []string `json:"policy_refs,omitempty"`
	ClaimRefs             []string `json:"claim_refs,omitempty"`
	VerificationRefs      []string `json:"verification_refs,omitempty"`
	RegistryRefs          []string `json:"registry_refs,omitempty"`
	EnforcementRefs       []string `json:"enforcement_refs,omitempty"`
	ExceptionRefs         []string `json:"exception_refs,omitempty"`
	EmergencyRefs         []string `json:"emergency_refs,omitempty"`
	MonitoringRefs        []string `json:"monitoring_refs,omitempty"`
	DashboardRefs         []string `json:"dashboard_refs,omitempty"`
	EvidenceRefs          []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs      []string `json:"evidence_hash_refs,omitempty"`
	AuditRefs             []string `json:"audit_refs,omitempty"`
	GovernanceEventRefs   []string `json:"governance_event_refs,omitempty"`
	CleanRoomIPReviewRefs []string `json:"clean_room_ip_review_refs,omitempty"`
	FreshnessState        string   `json:"freshness_state"`
	RevocationState       string   `json:"revocation_state"`
	SupersessionState     string   `json:"supersession_state"`
	DuplicateState        string   `json:"duplicate_state"`
	ConflictState         string   `json:"conflict_state"`
	UnrelatedState        string   `json:"unrelated_state"`
	TenantScopeState      string   `json:"tenant_scope_state"`
	Diagnostics           []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer  string   `json:"projection_disclaimer"`
}

type Point11ValDPublicationBoundaryFinalReview struct {
	CurrentState                  string   `json:"current_state"`
	PublicationReviewID           string   `json:"publication_review_id"`
	ModeledSurfaces               []string `json:"modeled_surfaces,omitempty"`
	CustomerVisibleSurfaces       []string `json:"customer_visible_surfaces,omitempty"`
	PublicSurfaces                []string `json:"public_surfaces,omitempty"`
	ExportSurfaces                []string `json:"export_surfaces,omitempty"`
	PartnerSurfaces               []string `json:"partner_surfaces,omitempty"`
	BuyerSurfaces                 []string `json:"buyer_surfaces,omitempty"`
	AgentOutputSurfaces           []string `json:"agent_output_surfaces,omitempty"`
	CleanRoomIPReviewRefs         []string `json:"clean_room_ip_review_refs,omitempty"`
	GovernanceEventRefs           []string `json:"governance_event_refs,omitempty"`
	CreatesPublicationSideEffects bool     `json:"creates_publication_side_effects"`
	CreatesCustomerFacingMaterial bool     `json:"creates_customer_facing_material"`
	CreatesAuthorityClaim         bool     `json:"creates_authority_claim"`
	CreatesCertificationClaim     bool     `json:"creates_certification_claim"`
	CreatesRegulatoryClaim        bool     `json:"creates_regulatory_claim"`
	CreatesComplianceGuarantee    bool     `json:"creates_compliance_guarantee"`
	Diagnostics                   []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
}

type Point11ValDFinalNoOverclaimReview struct {
	CurrentState            string   `json:"current_state"`
	NoOverclaimReviewID     string   `json:"no_overclaim_review_id"`
	ObservedClaims          []string `json:"observed_claims,omitempty"`
	ObservedDiagnostics     []string `json:"observed_diagnostics,omitempty"`
	ObservedDashboardText   []string `json:"observed_dashboard_text,omitempty"`
	ObservedPublicationText []string `json:"observed_publication_text,omitempty"`
	Denylist                []string `json:"denylist,omitempty"`
	SafeWordingExamples     []string `json:"safe_wording_examples,omitempty"`
	ReviewState             string   `json:"review_state"`
	Diagnostics             []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer    string   `json:"projection_disclaimer"`
}

type Point11ValDCleanRoomIPFinalReview struct {
	CurrentState                     string   `json:"current_state"`
	CleanRoomReviewID                string   `json:"clean_room_review_id"`
	PublicClaimRefs                  []string `json:"public_claim_refs,omitempty"`
	BuyerClaimRefs                   []string `json:"buyer_claim_refs,omitempty"`
	PartnerClaimRefs                 []string `json:"partner_claim_refs,omitempty"`
	CustomerVisibleClaimRefs         []string `json:"customer_visible_claim_refs,omitempty"`
	ThirdPartyRefs                   []string `json:"third_party_refs,omitempty"`
	LicenseReviewRefs                []string `json:"license_review_refs,omitempty"`
	IPReviewRefs                     []string `json:"ip_review_refs,omitempty"`
	CopiedCompetitorMaterialDetected bool     `json:"copied_competitor_material_detected"`
	CopiedUIDetected                 bool     `json:"copied_ui_detected"`
	CopiedWorkflowDetected           bool     `json:"copied_workflow_detected"`
	CopiedPrivateDocsDetected        bool     `json:"copied_private_docs_detected"`
	LegalClearanceClaimed            bool     `json:"legal_clearance_claimed"`
	FTOClaimed                       bool     `json:"fto_claimed"`
	ExternalLegalReviewRef           string   `json:"external_legal_review_ref"`
	ExternalFTOReviewRef             string   `json:"external_fto_review_ref"`
	Diagnostics                      []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
}

type Point11ValDAcceptedRisk struct {
	RiskID      string `json:"risk_id"`
	EvidenceRef string `json:"evidence_ref"`
	Scope       string `json:"scope"`
	OwnerRef    string `json:"owner_ref"`
	Expiry      string `json:"expiry"`
}

type Point11ValDCLBClosureLedger struct {
	CurrentState         string                    `json:"current_state"`
	ClosureLedgerID      string                    `json:"closure_ledger_id"`
	CLB0Findings         []string                  `json:"clb0_findings,omitempty"`
	CLB1Findings         []string                  `json:"clb1_findings,omitempty"`
	CLB2Findings         []string                  `json:"clb2_findings,omitempty"`
	CLB3Findings         []string                  `json:"clb3_findings,omitempty"`
	ResolvedFindings     []string                  `json:"resolved_findings,omitempty"`
	AcceptedRisks        []Point11ValDAcceptedRisk `json:"accepted_risks,omitempty"`
	DeferredItems        []string                  `json:"deferred_items,omitempty"`
	ReviewerResult       string                    `json:"reviewer_result"`
	Diagnostics          []string                  `json:"diagnostics,omitempty"`
	ProjectionDisclaimer string                    `json:"projection_disclaimer"`
}

type Point11ValDPassClosureManifest struct {
	CurrentState                    string   `json:"current_state"`
	ManifestID                      string   `json:"manifest_id"`
	PointID                         string   `json:"point_id"`
	WaveID                          string   `json:"wave_id"`
	Scope                           string   `json:"scope"`
	DependencyGateResult            string   `json:"dependency_gate_result"`
	IntegratedInvariantResult       string   `json:"integrated_invariant_result"`
	EvidenceGovernanceQualityResult string   `json:"evidence_governance_quality_result"`
	PublicationBoundaryResult       string   `json:"publication_boundary_result"`
	NoOverclaimResult               string   `json:"no_overclaim_result"`
	CleanRoomIPResult               string   `json:"clean_room_ip_result"`
	CLBClosureResult                string   `json:"clb_closure_result"`
	CommandsRun                     []string `json:"commands_run,omitempty"`
	TestsRun                        []string `json:"tests_run,omitempty"`
	GrepsRun                        []string `json:"greps_run,omitempty"`
	NegativeFixturesRun             []string `json:"negative_fixtures_run,omitempty"`
	EvidenceIdentitySummary         string   `json:"evidence_identity_summary"`
	PolicyClaimGovernanceSummary    string   `json:"policy_claim_governance_summary"`
	ExceptionEmergencySummary       string   `json:"exception_emergency_summary"`
	ABACEnforcementSummary          string   `json:"abac_enforcement_summary"`
	DashboardProjectionSummary      string   `json:"dashboard_projection_summary"`
	ProjectionBoundaryResultToken   string   `json:"projection_boundary_result"`
	ReviewerResult                  string   `json:"reviewer_result"`
	Point11PassAllowed              bool     `json:"point11_pass_allowed"`
	Point11PassToken                string   `json:"point11_pass_token"`
	Diagnostics                     []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer            string   `json:"projection_disclaimer"`
}

type Point11ValDFinalPoint11PassGate struct {
	CurrentState                           string   `json:"current_state"`
	FinalGateID                            string   `json:"final_gate_id"`
	DependencyState                        string   `json:"dependency_state"`
	InvariantState                         string   `json:"invariant_state"`
	QualityState                           string   `json:"quality_state"`
	PublicationState                       string   `json:"publication_state"`
	NoOverclaimState                       string   `json:"no_overclaim_state"`
	CleanRoomIPState                       string   `json:"clean_room_ip_state"`
	CLBClosureState                        string   `json:"clb_closure_state"`
	ManifestState                          string   `json:"manifest_state"`
	Point11PassAllowed                     bool     `json:"point11_pass_allowed"`
	Point11PassEmitted                     bool     `json:"point11_pass_emitted"`
	Point11PassToken                       string   `json:"point11_pass_token"`
	Point11PassObservedOutsideFinalClosure bool     `json:"point11_pass_observed_outside_final_closure"`
	Diagnostics                            []string `json:"diagnostics,omitempty"`
	ProjectionDisclaimer                   string   `json:"projection_disclaimer"`
}

type Point11ValDDiagnostics struct {
	CurrentState               string   `json:"current_state"`
	BlockingReasons            []string `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites        []string `json:"review_prerequisites,omitempty"`
	ComponentStates            []string `json:"component_states,omitempty"`
	DependencyReasons          []string `json:"dependency_reasons,omitempty"`
	IntegratedInvariantReasons []string `json:"integrated_invariant_reasons,omitempty"`
	QualityMapReasons          []string `json:"quality_map_reasons,omitempty"`
	PublicationReviewReasons   []string `json:"publication_review_reasons,omitempty"`
	NoOverclaimReasons         []string `json:"no_overclaim_reasons,omitempty"`
	CleanRoomIPReasons         []string `json:"clean_room_ip_reasons,omitempty"`
	CLBLedgerReasons           []string `json:"clb_ledger_reasons,omitempty"`
	ManifestReasons            []string `json:"manifest_reasons,omitempty"`
	FinalPassGateReasons       []string `json:"final_pass_gate_reasons,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type Point11ValDFoundation struct {
	CurrentState                  string                                         `json:"current_state"`
	BlockingReasons               []string                                       `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites           []string                                       `json:"review_prerequisites,omitempty"`
	DependencyState               string                                         `json:"dependency_state"`
	IntegratedInvariantState      string                                         `json:"integrated_invariant_state"`
	QualityMapState               string                                         `json:"quality_map_state"`
	PublicationReviewState        string                                         `json:"publication_review_state"`
	NoOverclaimReviewState        string                                         `json:"no_overclaim_review_state"`
	CleanRoomIPReviewState        string                                         `json:"clean_room_ip_review_state"`
	CLBClosureState               string                                         `json:"clb_closure_state"`
	PassClosureManifestState      string                                         `json:"pass_closure_manifest_state"`
	FinalPassGateState            string                                         `json:"final_pass_gate_state"`
	Point11PassToken              string                                         `json:"point11_pass_token,omitempty"`
	Diagnostics                   Point11ValDDiagnostics                         `json:"diagnostics"`
	ProjectionDisclaimer          string                                         `json:"projection_disclaimer"`
	CreatesAuthorityClaims        bool                                           `json:"creates_authority_claims"`
	CreatesPublicationSideEffects bool                                           `json:"creates_publication_side_effects"`
	CreatesSigningSideEffects     bool                                           `json:"creates_signing_side_effects"`
	CreatesAnchoringSideEffects   bool                                           `json:"creates_anchoring_side_effects"`
	CreatesExternalAPISideEffects bool                                           `json:"creates_external_api_side_effects"`
	CreatesProductionSideEffects  bool                                           `json:"creates_production_side_effects"`
	Val0Dependency                Point11ValDVal0DependencySnapshot              `json:"val0_dependency"`
	ValADependency                Point11ValDValADependencySnapshot              `json:"vala_dependency"`
	ValBDependency                Point11ValDValBDependencySnapshot              `json:"valb_dependency"`
	ValCDependency                Point11ValDValCDependencySnapshot              `json:"valc_dependency"`
	IntegratedInvariantReview     Point11ValDIntegratedGovernanceInvariantReview `json:"integrated_invariant_review"`
	QualityMap                    Point11ValDEvidenceGovernanceQualityMap        `json:"quality_map"`
	PublicationReview             Point11ValDPublicationBoundaryFinalReview      `json:"publication_review"`
	NoOverclaimReview             Point11ValDFinalNoOverclaimReview              `json:"no_overclaim_review"`
	CleanRoomIPReview             Point11ValDCleanRoomIPFinalReview              `json:"clean_room_ip_review"`
	CLBLedger                     Point11ValDCLBClosureLedger                    `json:"clb_ledger"`
	PassClosureManifest           Point11ValDPassClosureManifest                 `json:"pass_closure_manifest"`
	FinalPassGate                 Point11ValDFinalPoint11PassGate                `json:"final_pass_gate"`
}

func point11ValDContainsForbiddenText(values ...string) bool {
	return point11ValCContainsForbiddenText(values...)
}

func point11ValDRawExactNonEmpty(value string) bool {
	return value != "" && value == strings.TrimSpace(value) && !strings.ContainsAny(value, "\t\r\n")
}

func point11ValDRawCanonicalRefWithPrefixes(value string, prefixes []string) bool {
	return point11ValDRawExactNonEmpty(value) && point11Val0CanonicalRefWithPrefixes(value, prefixes)
}

func point11ValDRawTimestampValid(value string) bool {
	if !point11ValDRawExactNonEmpty(value) {
		return false
	}
	parsed, err := time.Parse(time.RFC3339, value)
	return err == nil && parsed.UTC().Format(time.RFC3339) == value
}

func point11ValDClosureRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{"closure_", "point11_closure_"})
}

func point11ValDManifestRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{"manifest_", "point11_manifest_"})
}

func point11ValDQualityMapRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{"quality_map_", "point11_quality_map_"})
}

func point11ValDPublicationReviewRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{"publication_review_", "point11_publication_review_"})
}

func point11ValDNoOverclaimReviewRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{"no_overclaim_review_", "point11_no_overclaim_review_"})
}

func point11ValDCleanRoomReviewRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{"clean_room_review_", "point11_clean_room_review_"})
}

func point11ValDCLBClosureRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{"clb_closure_", "point11_clb_closure_"})
}

func point11ValDFinalGateRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{"final_gate_", "point11_final_gate_"})
}

func point11ValDDependencyRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDDependencyRefPrefix, "point11_dependency_review_"})
}

func point11ValDThirdPartyRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDThirdPartyRefPrefix, "point11_third_party_"})
}

func point11ValDLicenseReviewRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDLicenseReviewRefPrefix, "point11_license_review_"})
}

func point11ValDIPReviewRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDIPReviewRefPrefix, "point11_ip_review_"})
}

func point11ValDExternalLegalReviewRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDExternalLegalReviewRefPrefix, "point11_external_legal_review_"})
}

func point11ValDExternalFTOReviewRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDExternalFTOReviewRefPrefix, "point11_external_fto_review_"})
}

func point11ValDCommandRunRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDCommandRunRefPrefix, "point11_command_run_"})
}

func point11ValDTestRunRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDTestRunRefPrefix, "point11_test_run_"})
}

func point11ValDGrepRunRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDGrepRunRefPrefix, "point11_grep_run_"})
}

func point11ValDNegativeFixtureRunRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDNegativeFixtureRunRefPrefix, "point11_negative_fixture_"})
}

func point11ValDFindingRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDFindingRefPrefix, "point11_finding_"})
}

func point11ValDAcceptedRiskRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDAcceptedRiskRefPrefix, "point11_accepted_risk_"})
}

func point11ValDDeferredItemRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDDeferredItemRefPrefix, "point11_deferred_item_"})
}

func point11ValDResolvedFindingRefValid(value string) bool {
	return point11ValDRawCanonicalRefWithPrefixes(value, []string{point11ValDResolvedFindingRefPrefix, "point11_resolved_finding_"})
}

func point11ValDStringListValid(values []string, validator func(string) bool) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point11ValDRawExactNonEmpty(value) || !validator(value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point11ValDOptionalStringListValid(values []string, validator func(string) bool) bool {
	if len(values) == 0 {
		return true
	}
	return point11ValDStringListValid(values, validator)
}

func point11ValDPublicationSurfacesValid(values []string) bool {
	if len(values) == 0 {
		return true
	}
	seen := map[string]struct{}{}
	allowed := append([]string{point11ValDPublicationSurfaceAgentOutput}, point11Val0PublicationSurfaces()...)
	for _, value := range values {
		if !point11ValDRawExactNonEmpty(value) || !point11Val0ContainsTrimmed(allowed, value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point11ValDEvidenceRefsValid(values []string) bool {
	if !point11Val0EvidenceRefsValid(values) {
		return false
	}
	for _, value := range values {
		if !point11ValDRawExactNonEmpty(value) {
			return false
		}
		lower := strings.ToLower(value)
		for _, blocked := range []string{
			"unknown",
			"unsupported",
			"invalid",
			"revoked",
			"expired",
			"superseded",
			"corrected",
			"blocked",
			"malformed",
			"placeholder",
			"junk",
			"marker",
			"global",
			"unscoped",
			"all-tenants",
			"wildcard",
		} {
			if strings.Contains(lower, blocked) {
				return false
			}
		}
	}
	return true
}

func point11ValDAnyPublicFacingSurfaces(surfaceLists ...[]string) bool {
	for _, surfaces := range surfaceLists {
		for _, surface := range surfaces {
			if point11Val0PublicFacingSurface(surface) {
				return true
			}
		}
	}
	return false
}

func point11ValDAnySurface(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func point11ValDGenericCheckStateActive(value string) bool {
	return value == point11ValDCheckStateActive
}

func point11ValDNoOverclaimDenylist() []string {
	return []string{
		"certified",
		"regulator-approved",
		"compliance guaranteed",
		"production approved",
		"deployment approved",
		"legally certified",
		"public badge",
		"official authority",
		"universal authority",
		"global truth",
		"mathematically proves total truth",
		"impossible to violate without detection",
		"supreme authority",
		"supreme arbiter",
		"self-healed and secure",
		"agent approved",
		"AI certified",
		"AI approved",
		"AI-approved",
		"AI decision",
		"AI legal proof",
		"autonomous remediation",
		"continuous compliance attestation",
		"guaranteed secure",
		"zero-risk",
		"compliant by default",
		"marketplace certified",
		"install success means ready",
		"SLA readiness means uptime guarantee",
	}
}

func point11ValDSafeWordingExamples() []string {
	return []string{
		"bounded claim",
		"evidence-linked governance decision",
		"policy-bound decision support",
		"audit-ready governance trail",
		"not a certification",
		"not regulator approval",
		"not production approval",
		"not deployment approval",
		"not compliance guarantee",
		"advisory projection only",
		"canonical evidence spine remains " + "source of truth",
	}
}

func SnapshotPoint11ValDVal0DependencyFromComputed(val0 Point11Val0Foundation, review Point11ValDVal0ReviewContext) Point11ValDVal0DependencySnapshot {
	return Point11ValDVal0DependencySnapshot{
		CurrentState:                  val0.CurrentState,
		DependencyState:               val0.DependencyState,
		PolicyContractState:           val0.PolicyContractState,
		ClaimGovernanceState:          val0.ClaimGovernanceState,
		AuthorityMatrixState:          val0.AuthorityMatrixState,
		ExceptionGovernanceState:      val0.ExceptionGovernanceState,
		ABACState:                     val0.ABACGovernanceState,
		DecisionBindingState:          val0.DecisionBindingState,
		NoOverclaimState:              val0.NoOverclaimState,
		CrossDomainCompatibilityState: val0.CrossDomainCompatibilityState,
		ProjectionDisclaimer:          val0.ProjectionDisclaimer,
		Point11PassEmitted:            review.Val0Point11PassEmitted,
		ReviewPrerequisites:           append([]string{}, review.ReviewPrerequisites...),
	}
}

func SnapshotPoint11ValDValADependencyFromComputed(valA Point11ValAFoundation, review Point11ValDValAReviewContext) Point11ValDValADependencySnapshot {
	reviewPrereqs := append([]string{}, valA.ReviewPrerequisites...)
	reviewPrereqs = append(reviewPrereqs, review.ReviewPrerequisites...)
	return Point11ValDValADependencySnapshot{
		CurrentState:                             valA.CurrentState,
		DependencyState:                          valA.DependencyState,
		RegistryState:                            valA.RegistryState,
		SignatureState:                           valA.SignatureState,
		AnchorState:                              valA.AnchorState,
		LifecycleTransitionState:                 valA.LifecycleTransitionState,
		PolicyUseState:                           valA.PolicyUseState,
		GraphState:                               valA.GraphState,
		ProjectionDisclaimer:                     valA.ProjectionDisclaimer,
		CreatesLegalRegulatoryCertificationClaim: valA.CreatesLegalRegulatoryCertificationClaim,
		CreatesPublicationSideEffects:            valA.CreatesPublicationSideEffects,
		Point11PassEmitted:                       review.ValAPoint11PassEmitted,
		CreatesSigningSideEffects:                review.ValACreatesSigningSideEffects,
		CreatesAnchoringSideEffects:              review.ValACreatesAnchoringSideEffects,
		CreatesExternalAPISideEffects:            review.ValACreatesExternalAPISideEffects,
		CreatesProductionSideEffects:             review.ValACreatesProductionSideEffects,
		OpenCLB0Findings:                         review.OpenCLB0Findings,
		OpenCLB1Findings:                         review.OpenCLB1Findings,
		OpenCLB2Findings:                         review.OpenCLB2Findings,
		ReviewPrerequisites:                      reviewPrereqs,
	}
}

func SnapshotPoint11ValDValBDependencyFromComputed(valB Point11ValBFoundation, review Point11ValDValBReviewContext) Point11ValDValBDependencySnapshot {
	reviewPrereqs := append([]string{}, valB.ReviewPrerequisites...)
	reviewPrereqs = append(reviewPrereqs, review.ReviewPrerequisites...)
	return Point11ValDValBDependencySnapshot{
		CurrentState:                  valB.CurrentState,
		DependencyState:               valB.DependencyState,
		ClaimTypeState:                valB.ClaimTypeState,
		IssuanceRequestState:          valB.IssuanceRequestState,
		IssuedClaimState:              valB.IssuedClaimState,
		RegistryState:                 valB.RegistryState,
		VerificationState:             valB.VerificationState,
		CrossDomainIntakeState:        valB.CrossDomainIntakeState,
		ProjectionDisclaimer:          valB.ProjectionDisclaimer,
		CreatesAuthorityClaims:        valB.CreatesAuthorityClaims,
		CreatesPublicationSideEffects: valB.CreatesPublicationSideEffects,
		Point11PassEmitted:            review.ValBPoint11PassEmitted,
		CreatesSigningSideEffects:     review.ValBCreatesSigningSideEffects,
		CreatesAnchoringSideEffects:   review.ValBCreatesAnchoringSideEffects,
		CreatesExternalAPISideEffects: review.ValBCreatesExternalAPISideEffects,
		CreatesProductionSideEffects:  review.ValBCreatesProductionSideEffects,
		OpenCLB0Findings:              review.OpenCLB0Findings,
		OpenCLB1Findings:              review.OpenCLB1Findings,
		OpenCLB2Findings:              review.OpenCLB2Findings,
		ReviewPrerequisites:           reviewPrereqs,
	}
}

func SnapshotPoint11ValDValCDependencyFromComputed(valC Point11ValCFoundation, review Point11ValDValCReviewContext) Point11ValDValCDependencySnapshot {
	reviewPrereqs := append([]string{}, valC.ReviewPrerequisites...)
	reviewPrereqs = append(reviewPrereqs, review.ReviewPrerequisites...)
	return Point11ValDValCDependencySnapshot{
		CurrentState:                      valC.CurrentState,
		DependencyState:                   valC.DependencyState,
		EnforcementInputState:             valC.EnforcementInputState,
		EnforcementResultState:            valC.EnforcementResultState,
		ABACDecisionState:                 valC.ABACDecisionState,
		ExceptionDecisionState:            valC.ExceptionDecisionState,
		PrecedenceState:                   valC.PrecedenceState,
		MonitoringState:                   valC.MonitoringState,
		DashboardState:                    valC.DashboardState,
		ProjectionDisclaimer:              valC.ProjectionDisclaimer,
		CreatesAuthorityClaims:            valC.CreatesAuthorityClaims,
		CreatesPublicationSideEffects:     valC.CreatesPublicationSideEffects,
		CreatesRealEnforcementSideEffects: valC.CreatesRealEnforcementSideEffects,
		Point11PassEmitted:                review.ValCPoint11PassEmitted,
		CreatesExternalAPISideEffects:     review.ValCCreatesExternalAPISideEffects,
		CreatesProductionSideEffects:      review.ValCCreatesProductionSideEffects,
		DashboardSourceOfTruthBypass:      review.ValCDashboardSourceOfTruthBypass,
		OpenCLB0Findings:                  review.OpenCLB0Findings,
		OpenCLB1Findings:                  review.OpenCLB1Findings,
		OpenCLB2Findings:                  review.OpenCLB2Findings,
		ReviewPrerequisites:               reviewPrereqs,
	}
}

func point11ValDDefaultVal0ReviewContext() Point11ValDVal0ReviewContext {
	return Point11ValDVal0ReviewContext{}
}

func point11ValDDefaultValAReviewContext() Point11ValDValAReviewContext {
	return Point11ValDValAReviewContext{}
}

func point11ValDDefaultValBReviewContext() Point11ValDValBReviewContext {
	return Point11ValDValBReviewContext{}
}

func point11ValDDefaultValCReviewContext() Point11ValDValCReviewContext {
	return Point11ValDValCReviewContext{}
}

func point11ValDVal0SnapshotStateAndReasons(model Point11ValDVal0DependencySnapshot) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "val0_projection_disclaimer_blocked")
	}
	if model.Point11PassEmitted {
		reasons = append(reasons, "val0_point11_pass_emitted")
	}
	if model.PolicyContractState != Point11Val0PolicyContractStateActive {
		reasons = append(reasons, "val0_policy_contract_not_active")
	}
	if model.ClaimGovernanceState != Point11Val0ClaimGovernanceStateActive {
		reasons = append(reasons, "val0_claim_governance_not_active")
	}
	if model.AuthorityMatrixState != Point11Val0AuthorityMatrixStateActive {
		reasons = append(reasons, "val0_authority_matrix_not_active")
	}
	if model.ExceptionGovernanceState != Point11Val0ExceptionGovernanceStateActive {
		reasons = append(reasons, "val0_exception_governance_not_active")
	}
	if model.ABACState != Point11Val0ABACStateActive {
		reasons = append(reasons, "val0_abac_not_active")
	}
	if model.DecisionBindingState != Point11Val0DecisionBindingStateActive {
		reasons = append(reasons, "val0_decision_binding_not_active")
	}
	if model.NoOverclaimState != Point11Val0NoOverclaimStateActive {
		reasons = append(reasons, "val0_no_overclaim_not_active")
	}
	if model.CrossDomainCompatibilityState == Point11Val0CrossDomainCompatibilityStateBlocked {
		reasons = append(reasons, "val0_cross_domain_compatibility_blocked")
	}
	if len(reasons) > 0 {
		return Point11ValDDependencyStateBlocked, reasons
	}
	if model.CurrentState == Point11Val0StateActive &&
		model.DependencyState == Point11Val0DependencyStateActive &&
		model.CrossDomainCompatibilityState == Point11Val0CrossDomainCompatibilityStateActive &&
		len(model.ReviewPrerequisites) == 0 {
		return Point11ValDDependencyStateActive, nil
	}
	if len(model.ReviewPrerequisites) > 0 ||
		model.CurrentState == Point11Val0StateReviewRequired ||
		model.DependencyState == Point11Val0DependencyStateReviewRequired ||
		model.CrossDomainCompatibilityState == Point11Val0CrossDomainCompatibilityStateReviewRequired {
		return Point11ValDDependencyStateReviewRequired, []string{"val0_dependency_review_required"}
	}
	return Point11ValDDependencyStateBlocked, []string{"val0_dependency_not_active"}
}

func point11ValDValASnapshotStateAndReasons(model Point11ValDValADependencySnapshot) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "vala_projection_disclaimer_blocked")
	}
	if model.Point11PassEmitted {
		reasons = append(reasons, "vala_point11_pass_emitted")
	}
	if model.CreatesLegalRegulatoryCertificationClaim {
		reasons = append(reasons, "vala_authority_claim_blocked")
	}
	if model.CreatesPublicationSideEffects {
		reasons = append(reasons, "vala_publication_side_effects_blocked")
	}
	if model.CreatesSigningSideEffects {
		reasons = append(reasons, "vala_signing_side_effects_blocked")
	}
	if model.CreatesAnchoringSideEffects {
		reasons = append(reasons, "vala_anchoring_side_effects_blocked")
	}
	if model.CreatesExternalAPISideEffects {
		reasons = append(reasons, "vala_external_api_side_effects_blocked")
	}
	if model.CreatesProductionSideEffects {
		reasons = append(reasons, "vala_production_side_effects_blocked")
	}
	if model.OpenCLB0Findings > 0 {
		reasons = append(reasons, "vala_open_clb0_findings")
	}
	if model.OpenCLB1Findings > 0 {
		reasons = append(reasons, "vala_open_clb1_findings")
	}
	if model.OpenCLB2Findings > 0 {
		reasons = append(reasons, "vala_open_clb2_findings")
	}
	if model.RegistryState != Point11ValARegistryStateActive {
		reasons = append(reasons, "vala_registry_not_active")
	}
	if model.SignatureState != Point11ValASignatureStateActive {
		reasons = append(reasons, "vala_signature_not_active")
	}
	if model.AnchorState != Point11ValAAnchorStateActive {
		reasons = append(reasons, "vala_anchor_not_active")
	}
	if model.LifecycleTransitionState != Point11ValALifecycleTransitionStateActive {
		reasons = append(reasons, "vala_lifecycle_transition_not_active")
	}
	if model.PolicyUseState != Point11ValAPolicyUseStateActive {
		reasons = append(reasons, "vala_policy_use_not_active")
	}
	if model.GraphState != Point11ValAGraphStateActive {
		reasons = append(reasons, "vala_graph_not_active")
	}
	if len(reasons) > 0 {
		return Point11ValDDependencyStateBlocked, reasons
	}
	if model.CurrentState == Point11ValAStateActive &&
		model.DependencyState == Point11ValADependencyStateActive &&
		len(model.ReviewPrerequisites) == 0 {
		return Point11ValDDependencyStateActive, nil
	}
	if len(model.ReviewPrerequisites) > 0 ||
		model.CurrentState == Point11ValAStateReviewRequired ||
		model.DependencyState == Point11ValADependencyStateReviewRequired {
		return Point11ValDDependencyStateReviewRequired, []string{"vala_dependency_review_required"}
	}
	return Point11ValDDependencyStateBlocked, []string{"vala_dependency_not_active"}
}

func point11ValDValBSnapshotStateAndReasons(model Point11ValDValBDependencySnapshot) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "valb_projection_disclaimer_blocked")
	}
	if model.Point11PassEmitted {
		reasons = append(reasons, "valb_point11_pass_emitted")
	}
	if model.CreatesAuthorityClaims {
		reasons = append(reasons, "valb_authority_claim_blocked")
	}
	if model.CreatesPublicationSideEffects {
		reasons = append(reasons, "valb_publication_side_effects_blocked")
	}
	if model.CreatesSigningSideEffects {
		reasons = append(reasons, "valb_signing_side_effects_blocked")
	}
	if model.CreatesAnchoringSideEffects {
		reasons = append(reasons, "valb_anchoring_side_effects_blocked")
	}
	if model.CreatesExternalAPISideEffects {
		reasons = append(reasons, "valb_external_api_side_effects_blocked")
	}
	if model.CreatesProductionSideEffects {
		reasons = append(reasons, "valb_production_side_effects_blocked")
	}
	if model.OpenCLB0Findings > 0 {
		reasons = append(reasons, "valb_open_clb0_findings")
	}
	if model.OpenCLB1Findings > 0 {
		reasons = append(reasons, "valb_open_clb1_findings")
	}
	if model.OpenCLB2Findings > 0 {
		reasons = append(reasons, "valb_open_clb2_findings")
	}
	if model.ClaimTypeState != Point11ValBClaimTypeStateActive {
		reasons = append(reasons, "valb_claim_type_not_active")
	}
	if model.IssuanceRequestState != Point11ValBIssuanceRequestStateActive {
		reasons = append(reasons, "valb_issuance_request_not_active")
	}
	if model.IssuedClaimState != Point11ValBIssuedClaimStateActive {
		reasons = append(reasons, "valb_issued_claim_not_active")
	}
	if model.RegistryState != Point11ValBRegistryStateActive {
		reasons = append(reasons, "valb_registry_not_active")
	}
	if model.VerificationState != Point11ValBVerificationStateActive {
		reasons = append(reasons, "valb_verification_not_active")
	}
	if model.CrossDomainIntakeState == Point11ValBCrossDomainIntakeStateBlocked {
		reasons = append(reasons, "valb_cross_domain_intake_blocked")
	}
	if len(reasons) > 0 {
		return Point11ValDDependencyStateBlocked, reasons
	}
	if model.CurrentState == Point11ValBStateActive &&
		model.DependencyState == Point11ValBDependencyStateActive &&
		model.CrossDomainIntakeState == Point11ValBCrossDomainIntakeStateActive &&
		len(model.ReviewPrerequisites) == 0 {
		return Point11ValDDependencyStateActive, nil
	}
	if len(model.ReviewPrerequisites) > 0 ||
		model.CurrentState == Point11ValBStateReviewRequired ||
		model.DependencyState == Point11ValBDependencyStateReviewRequired ||
		model.CrossDomainIntakeState == Point11ValBCrossDomainIntakeStateReviewRequired {
		return Point11ValDDependencyStateReviewRequired, []string{"valb_dependency_review_required"}
	}
	return Point11ValDDependencyStateBlocked, []string{"valb_dependency_not_active"}
}

func point11ValDValCSnapshotStateAndReasons(model Point11ValDValCDependencySnapshot) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "valc_projection_disclaimer_blocked")
	}
	if model.Point11PassEmitted {
		reasons = append(reasons, "valc_point11_pass_emitted")
	}
	if model.CreatesAuthorityClaims {
		reasons = append(reasons, "valc_authority_claim_blocked")
	}
	if model.CreatesPublicationSideEffects {
		reasons = append(reasons, "valc_publication_side_effects_blocked")
	}
	if model.CreatesRealEnforcementSideEffects {
		reasons = append(reasons, "valc_real_enforcement_side_effects_blocked")
	}
	if model.CreatesExternalAPISideEffects {
		reasons = append(reasons, "valc_external_api_side_effects_blocked")
	}
	if model.CreatesProductionSideEffects {
		reasons = append(reasons, "valc_production_side_effects_blocked")
	}
	if model.DashboardSourceOfTruthBypass {
		reasons = append(reasons, "valc_dashboard_source_of_truth_bypass")
	}
	if model.OpenCLB0Findings > 0 {
		reasons = append(reasons, "valc_open_clb0_findings")
	}
	if model.OpenCLB1Findings > 0 {
		reasons = append(reasons, "valc_open_clb1_findings")
	}
	if model.OpenCLB2Findings > 0 {
		reasons = append(reasons, "valc_open_clb2_findings")
	}
	if model.EnforcementInputState != Point11ValCEnforcementInputStateActive {
		reasons = append(reasons, "valc_enforcement_input_not_active")
	}
	if model.EnforcementResultState == Point11ValCEnforcementResultStateBlocked {
		reasons = append(reasons, "valc_enforcement_result_blocked")
	}
	if model.ABACDecisionState != Point11ValCABACDecisionStateActive {
		reasons = append(reasons, "valc_abac_not_active")
	}
	if model.ExceptionDecisionState == Point11ValCExceptionDecisionStateBlocked {
		reasons = append(reasons, "valc_exception_decision_blocked")
	}
	if model.PrecedenceState == Point11ValCPrecedenceStateBlocked {
		reasons = append(reasons, "valc_precedence_blocked")
	}
	if model.MonitoringState == Point11ValCMonitoringStateBlocked {
		reasons = append(reasons, "valc_monitoring_blocked")
	}
	if model.DashboardState != Point11ValCDashboardStateActive {
		reasons = append(reasons, "valc_dashboard_not_active")
	}
	if len(reasons) > 0 {
		return Point11ValDDependencyStateBlocked, reasons
	}
	if model.CurrentState == Point11ValCStateActive &&
		model.DependencyState == Point11ValCDependencyStateActive &&
		model.EnforcementResultState == Point11ValCEnforcementResultStateActive &&
		model.ExceptionDecisionState == Point11ValCExceptionDecisionStateActive &&
		model.PrecedenceState == Point11ValCPrecedenceStateActive &&
		model.MonitoringState == Point11ValCMonitoringStateActive &&
		len(model.ReviewPrerequisites) == 0 {
		return Point11ValDDependencyStateActive, nil
	}
	if len(model.ReviewPrerequisites) > 0 ||
		model.CurrentState == Point11ValCStateReviewRequired ||
		model.DependencyState == Point11ValCDependencyStateReviewRequired ||
		model.EnforcementResultState == Point11ValCEnforcementResultStateReviewRequired ||
		model.ExceptionDecisionState == Point11ValCExceptionDecisionStateReviewRequired ||
		model.PrecedenceState == Point11ValCPrecedenceStateReviewRequired ||
		model.MonitoringState == Point11ValCMonitoringStateReviewRequired {
		return Point11ValDDependencyStateReviewRequired, []string{"valc_dependency_review_required"}
	}
	return Point11ValDDependencyStateBlocked, []string{"valc_dependency_not_active"}
}

func point11ValDDependencyStateAndReasons(model Point11ValDDependencyBundle) (string, []string) {
	reasons := []string{}
	val0State, val0Reasons := point11ValDVal0SnapshotStateAndReasons(model.Val0)
	valAState, valAReasons := point11ValDValASnapshotStateAndReasons(model.ValA)
	valBState, valBReasons := point11ValDValBSnapshotStateAndReasons(model.ValB)
	valCState, valCReasons := point11ValDValCSnapshotStateAndReasons(model.ValC)
	if val0State == Point11ValDDependencyStateBlocked {
		reasons = append(reasons, val0Reasons...)
	}
	if valAState == Point11ValDDependencyStateBlocked {
		reasons = append(reasons, valAReasons...)
	}
	if valBState == Point11ValDDependencyStateBlocked {
		reasons = append(reasons, valBReasons...)
	}
	if valCState == Point11ValDDependencyStateBlocked {
		reasons = append(reasons, valCReasons...)
	}
	if len(reasons) > 0 {
		return Point11ValDDependencyStateBlocked, reasons
	}
	if val0State == Point11ValDDependencyStateActive &&
		valAState == Point11ValDDependencyStateActive &&
		valBState == Point11ValDDependencyStateActive &&
		valCState == Point11ValDDependencyStateActive {
		return Point11ValDDependencyStateActive, nil
	}
	return Point11ValDDependencyStateReviewRequired, []string{"upstream_dependency_review_required"}
}

func EvaluatePoint11ValDDependencyState(model Point11ValDDependencyBundle) string {
	state, _ := point11ValDDependencyStateAndReasons(model)
	return state
}

func point11ValDIntegratedInvariantStateAndReasons(model Point11ValDIntegratedGovernanceInvariantReview) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "integrated_invariant_projection_disclaimer_blocked")
	}
	if !point11ValDClosureRefValid(model.InvariantReviewID) ||
		!point11ValDDependencyRefValid(model.Val0Ref) ||
		!point11ValDDependencyRefValid(model.ValARef) ||
		!point11ValDDependencyRefValid(model.ValBRef) ||
		!point11ValDDependencyRefValid(model.ValCRef) {
		reasons = append(reasons, "integrated_invariant_identity_invalid")
	}
	for _, state := range []string{
		model.PolicyAuthorityConsistencyState,
		model.ClaimAuthorityConsistencyState,
		model.PublicationBoundaryState,
		model.NoOverclaimState,
		model.CleanRoomIPState,
		model.ProjectionBoundaryState,
		model.ExceptionEmergencyConsistencyState,
		model.ABACEnforcementConsistencyState,
		model.DashboardProjectionState,
		model.Point11PassRuleState,
	} {
		if !point11ValDGenericCheckStateActive(state) {
			reasons = append(reasons, "integrated_invariant_component_blocked")
			break
		}
	}
	if point11ValDContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "integrated_invariant_overclaim_detected")
	}
	if len(reasons) > 0 {
		return Point11ValDIntegratedInvariantStateBlocked, reasons
	}
	return Point11ValDIntegratedInvariantStateActive, nil
}

func EvaluatePoint11ValDIntegratedInvariantState(model Point11ValDIntegratedGovernanceInvariantReview) string {
	state, _ := point11ValDIntegratedInvariantStateAndReasons(model)
	return state
}

func point11ValDQualityMapStateAndReasons(model Point11ValDEvidenceGovernanceQualityMap) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "quality_map_projection_disclaimer_blocked")
	}
	if !point11ValDQualityMapRefValid(model.QualityMapID) {
		reasons = append(reasons, "quality_map_id_invalid")
	}
	if !point11ValDStringListValid(model.PolicyRefs, point11ValAPolicyRefValid) {
		reasons = append(reasons, "quality_map_policy_refs_invalid")
	}
	if !point11ValDStringListValid(model.ClaimRefs, point11ValBClaimRefValid) {
		reasons = append(reasons, "quality_map_claim_refs_invalid")
	}
	if !point11ValDStringListValid(model.VerificationRefs, point11ValBVerificationRefValid) {
		reasons = append(reasons, "quality_map_verification_refs_invalid")
	}
	if !point11ValDStringListValid(model.RegistryRefs, point11ValBClaimRegistryRefValid) {
		reasons = append(reasons, "quality_map_registry_refs_invalid")
	}
	if !point11ValDStringListValid(model.EnforcementRefs, point11ValCEnforcementRefValid) {
		reasons = append(reasons, "quality_map_enforcement_refs_invalid")
	}
	if !point11ValDStringListValid(model.ExceptionRefs, point11ValCExceptionRefValid) {
		reasons = append(reasons, "quality_map_exception_refs_invalid")
	}
	if !point11ValDStringListValid(model.EmergencyRefs, point11ValCEmergencyRefValid) {
		reasons = append(reasons, "quality_map_emergency_refs_invalid")
	}
	if !point11ValDStringListValid(model.MonitoringRefs, point11ValCMonitoringRefValid) {
		reasons = append(reasons, "quality_map_monitoring_refs_invalid")
	}
	if !point11ValDStringListValid(model.DashboardRefs, point11ValCDashboardRefValid) {
		reasons = append(reasons, "quality_map_dashboard_refs_invalid")
	}
	if !point11ValDEvidenceRefsValid(model.EvidenceRefs) {
		reasons = append(reasons, "quality_map_evidence_refs_invalid")
	}
	if !point11ValBEvidenceHashRefsValid(model.EvidenceHashRefs) {
		reasons = append(reasons, "quality_map_evidence_hash_refs_invalid")
	}
	if !point11ValDStringListValid(model.AuditRefs, point11ValBAuditRefValid) {
		reasons = append(reasons, "quality_map_audit_refs_invalid")
	}
	if !point11ValDStringListValid(model.GovernanceEventRefs, point11ValBGovernanceEventRefValid) {
		reasons = append(reasons, "quality_map_governance_event_refs_invalid")
	}
	if !point11ValDStringListValid(model.CleanRoomIPReviewRefs, point11ValBCleanRoomReviewRefValid) {
		reasons = append(reasons, "quality_map_clean_room_ip_review_refs_invalid")
	}
	if model.FreshnessState != point11ValDFreshnessStateActive {
		reasons = append(reasons, "quality_map_freshness_blocked")
	}
	if model.RevocationState != point11ValDRevocationStateActive {
		reasons = append(reasons, "quality_map_revocation_blocked")
	}
	if model.SupersessionState != point11ValDSupersessionStateActive {
		reasons = append(reasons, "quality_map_supersession_blocked")
	}
	if model.DuplicateState != point11ValDDuplicateStateActive {
		reasons = append(reasons, "quality_map_duplicate_blocked")
	}
	if model.ConflictState != point11ValDConflictStateActive {
		reasons = append(reasons, "quality_map_conflict_blocked")
	}
	if model.UnrelatedState != point11ValDUnrelatedStateActive {
		reasons = append(reasons, "quality_map_unrelated_blocked")
	}
	if model.TenantScopeState != point11ValDTenantScopeStateActive {
		reasons = append(reasons, "quality_map_tenant_scope_blocked")
	}
	if point11ValDContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "quality_map_overclaim_detected")
	}
	if len(reasons) > 0 {
		return Point11ValDQualityMapStateBlocked, reasons
	}
	return Point11ValDQualityMapStateActive, nil
}

func EvaluatePoint11ValDQualityMapState(model Point11ValDEvidenceGovernanceQualityMap) string {
	state, _ := point11ValDQualityMapStateAndReasons(model)
	return state
}

func point11ValDPublicationReviewStateAndReasons(model Point11ValDPublicationBoundaryFinalReview) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "publication_review_projection_disclaimer_blocked")
	}
	if !point11ValDPublicationReviewRefValid(model.PublicationReviewID) {
		reasons = append(reasons, "publication_review_id_invalid")
	}
	for _, surfaces := range [][]string{
		model.ModeledSurfaces,
		model.CustomerVisibleSurfaces,
		model.PublicSurfaces,
		model.ExportSurfaces,
		model.PartnerSurfaces,
		model.BuyerSurfaces,
		model.AgentOutputSurfaces,
	} {
		if !point11ValDPublicationSurfacesValid(surfaces) {
			reasons = append(reasons, "publication_review_surface_list_invalid")
			break
		}
	}
	if model.CreatesPublicationSideEffects {
		reasons = append(reasons, "publication_review_side_effects_blocked")
	}
	if model.CreatesCustomerFacingMaterial {
		reasons = append(reasons, "publication_review_customer_material_blocked")
	}
	if model.CreatesAuthorityClaim {
		reasons = append(reasons, "publication_review_authority_claim_blocked")
	}
	if model.CreatesCertificationClaim {
		reasons = append(reasons, "publication_review_certification_claim_blocked")
	}
	if model.CreatesRegulatoryClaim {
		reasons = append(reasons, "publication_review_regulatory_claim_blocked")
	}
	if model.CreatesComplianceGuarantee {
		reasons = append(reasons, "publication_review_compliance_guarantee_blocked")
	}
	if point11ValDAnyPublicFacingSurfaces(model.CustomerVisibleSurfaces, model.PublicSurfaces, model.ExportSurfaces, model.PartnerSurfaces, model.BuyerSurfaces) {
		if !point11ValDStringListValid(model.CleanRoomIPReviewRefs, point11ValBCleanRoomReviewRefValid) {
			reasons = append(reasons, "publication_review_clean_room_ip_missing")
		}
		if !point11ValDStringListValid(model.GovernanceEventRefs, point11ValBGovernanceEventRefValid) {
			reasons = append(reasons, "publication_review_governance_event_missing")
		}
	}
	if len(model.AgentOutputSurfaces) > 0 &&
		(point11ValDAnyPublicFacingSurfaces(model.CustomerVisibleSurfaces, model.PublicSurfaces) ||
			point11ValDAnySurface(model.AgentOutputSurfaces, point11ValDPublicationSurfaceAgentOutput)) &&
		!point11ValDStringListValid(model.GovernanceEventRefs, point11ValBGovernanceEventRefValid) {
		reasons = append(reasons, "publication_review_agent_output_without_governance")
	}
	if point11ValDContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "publication_review_overclaim_detected")
	}
	if len(reasons) > 0 {
		return Point11ValDPublicationReviewStateBlocked, reasons
	}
	return Point11ValDPublicationReviewStateActive, nil
}

func EvaluatePoint11ValDPublicationReviewState(model Point11ValDPublicationBoundaryFinalReview) string {
	state, _ := point11ValDPublicationReviewStateAndReasons(model)
	return state
}

func point11ValDNoOverclaimReviewStateAndReasons(model Point11ValDFinalNoOverclaimReview) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "no_overclaim_projection_disclaimer_blocked")
	}
	if !point11ValDNoOverclaimReviewRefValid(model.NoOverclaimReviewID) {
		reasons = append(reasons, "no_overclaim_review_id_invalid")
	}
	if len(model.Denylist) == 0 || len(model.SafeWordingExamples) == 0 {
		reasons = append(reasons, "no_overclaim_reference_sets_missing")
	}
	if model.ReviewState != Point11ValDNoOverclaimReviewStateActive {
		reasons = append(reasons, "no_overclaim_review_state_invalid")
	}
	if point11ValDContainsForbiddenText(model.ObservedClaims...) ||
		point11ValDContainsForbiddenText(model.ObservedDiagnostics...) ||
		point11ValDContainsForbiddenText(model.ObservedDashboardText...) ||
		point11ValDContainsForbiddenText(model.ObservedPublicationText...) {
		reasons = append(reasons, "no_overclaim_forbidden_text_detected")
	}
	if len(reasons) > 0 {
		return Point11ValDNoOverclaimReviewStateBlocked, reasons
	}
	return Point11ValDNoOverclaimReviewStateActive, nil
}

func EvaluatePoint11ValDNoOverclaimReviewState(model Point11ValDFinalNoOverclaimReview) string {
	state, _ := point11ValDNoOverclaimReviewStateAndReasons(model)
	return state
}

func point11ValDCleanRoomIPReviewStateAndReasons(model Point11ValDCleanRoomIPFinalReview) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "clean_room_projection_disclaimer_blocked")
	}
	if !point11ValDCleanRoomReviewRefValid(model.CleanRoomReviewID) {
		reasons = append(reasons, "clean_room_review_id_invalid")
	}
	if len(model.PublicClaimRefs) > 0 && !point11ValBClaimRefListValid(model.PublicClaimRefs) {
		reasons = append(reasons, "clean_room_public_claim_refs_invalid")
	}
	if len(model.BuyerClaimRefs) > 0 && !point11ValBClaimRefListValid(model.BuyerClaimRefs) {
		reasons = append(reasons, "clean_room_buyer_claim_refs_invalid")
	}
	if len(model.PartnerClaimRefs) > 0 && !point11ValBClaimRefListValid(model.PartnerClaimRefs) {
		reasons = append(reasons, "clean_room_partner_claim_refs_invalid")
	}
	if len(model.CustomerVisibleClaimRefs) > 0 && !point11ValBClaimRefListValid(model.CustomerVisibleClaimRefs) {
		reasons = append(reasons, "clean_room_customer_visible_claim_refs_invalid")
	}
	if len(model.ThirdPartyRefs) > 0 && !point11ValDStringListValid(model.ThirdPartyRefs, point11ValDThirdPartyRefValid) {
		reasons = append(reasons, "clean_room_third_party_refs_invalid")
	}
	if len(model.ThirdPartyRefs) > 0 {
		if !point11ValDStringListValid(model.LicenseReviewRefs, point11ValDLicenseReviewRefValid) {
			reasons = append(reasons, "clean_room_license_review_missing")
		}
		if !point11ValDStringListValid(model.IPReviewRefs, point11ValDIPReviewRefValid) {
			reasons = append(reasons, "clean_room_ip_review_missing")
		}
	}
	if len(model.PublicClaimRefs) > 0 || len(model.BuyerClaimRefs) > 0 || len(model.PartnerClaimRefs) > 0 || len(model.CustomerVisibleClaimRefs) > 0 {
		if !point11ValDStringListValid(model.IPReviewRefs, point11ValDIPReviewRefValid) {
			reasons = append(reasons, "clean_room_customer_public_ip_review_missing")
		}
	}
	if model.CopiedCompetitorMaterialDetected {
		reasons = append(reasons, "clean_room_copied_competitor_material_detected")
	}
	if model.CopiedUIDetected {
		reasons = append(reasons, "clean_room_copied_ui_detected")
	}
	if model.CopiedWorkflowDetected {
		reasons = append(reasons, "clean_room_copied_workflow_detected")
	}
	if model.CopiedPrivateDocsDetected {
		reasons = append(reasons, "clean_room_copied_private_docs_detected")
	}
	if model.LegalClearanceClaimed && !point11ValDExternalLegalReviewRefValid(model.ExternalLegalReviewRef) {
		reasons = append(reasons, "clean_room_legal_clearance_without_external_review")
	}
	if model.FTOClaimed && !point11ValDExternalFTOReviewRefValid(model.ExternalFTOReviewRef) {
		reasons = append(reasons, "clean_room_fto_without_external_review")
	}
	if point11ValDContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "clean_room_overclaim_detected")
	}
	if len(reasons) > 0 {
		return Point11ValDCleanRoomIPReviewStateBlocked, reasons
	}
	return Point11ValDCleanRoomIPReviewStateActive, nil
}

func EvaluatePoint11ValDCleanRoomIPReviewState(model Point11ValDCleanRoomIPFinalReview) string {
	state, _ := point11ValDCleanRoomIPReviewStateAndReasons(model)
	return state
}

func point11ValDCLBClosureStateAndReasons(model Point11ValDCLBClosureLedger) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "clb_ledger_projection_disclaimer_blocked")
	}
	if !point11ValDCLBClosureRefValid(model.ClosureLedgerID) {
		reasons = append(reasons, "clb_ledger_id_invalid")
	}
	if len(model.CLB0Findings) > 0 {
		reasons = append(reasons, "clb0_open_findings")
	}
	if len(model.CLB1Findings) > 0 {
		reasons = append(reasons, "clb1_open_findings")
	}
	if len(model.CLB2Findings) > 0 {
		reasons = append(reasons, "clb2_open_findings")
	}
	for _, finding := range model.CLB0Findings {
		if !point11ValDFindingRefValid(finding) {
			reasons = append(reasons, "clb0_finding_invalid")
			break
		}
	}
	for _, finding := range model.CLB1Findings {
		if !point11ValDFindingRefValid(finding) {
			reasons = append(reasons, "clb1_finding_invalid")
			break
		}
	}
	for _, finding := range model.CLB2Findings {
		if !point11ValDFindingRefValid(finding) {
			reasons = append(reasons, "clb2_finding_invalid")
			break
		}
	}
	if len(model.CLB3Findings) > 0 && !point11ValDStringListValid(model.CLB3Findings, point11ValDFindingRefValid) {
		reasons = append(reasons, "clb3_finding_invalid")
	}
	if len(model.ResolvedFindings) > 0 && !point11ValDStringListValid(model.ResolvedFindings, point11ValDResolvedFindingRefValid) {
		reasons = append(reasons, "resolved_finding_invalid")
	}
	if len(model.DeferredItems) > 0 && !point11ValDStringListValid(model.DeferredItems, point11ValDDeferredItemRefValid) {
		reasons = append(reasons, "deferred_item_invalid")
	}
	for _, risk := range model.AcceptedRisks {
		if !point11ValDAcceptedRiskRefValid(risk.RiskID) ||
			!point11ValDEvidenceRefsValid([]string{risk.EvidenceRef}) ||
			!point11Val0ScopeValid(risk.Scope) ||
			!point11ValCActorRefValid(risk.OwnerRef) ||
			!point11ValDRawTimestampValid(risk.Expiry) {
			reasons = append(reasons, "accepted_risk_missing_required_fields")
			break
		}
		expiry, _ := time.Parse(time.RFC3339, risk.Expiry)
		if expiry.Before(time.Now().UTC()) {
			reasons = append(reasons, "accepted_risk_expired")
			break
		}
	}
	if model.ReviewerResult != point11ValDReviewerResultPassConfirmed {
		reasons = append(reasons, "clb_ledger_reviewer_not_pass_confirmed")
	}
	if point11ValDContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "clb_ledger_overclaim_detected")
	}
	if len(reasons) > 0 {
		return Point11ValDCLBClosureStateBlocked, reasons
	}
	return Point11ValDCLBClosureStateActive, nil
}

func EvaluatePoint11ValDCLBClosureState(model Point11ValDCLBClosureLedger) string {
	state, _ := point11ValDCLBClosureStateAndReasons(model)
	return state
}

func point11ValDPassClosureManifestStateAndReasons(model Point11ValDPassClosureManifest, foundation Point11ValDFoundation) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "manifest_projection_disclaimer_blocked")
	}
	if !point11ValDManifestRefValid(model.ManifestID) {
		reasons = append(reasons, "manifest_id_invalid")
	}
	if model.PointID != point11ValDPointID {
		reasons = append(reasons, "manifest_point_id_invalid")
	}
	if model.WaveID != point11ValDWaveID {
		reasons = append(reasons, "manifest_wave_id_invalid")
	}
	if model.Scope != point11ValDScope {
		reasons = append(reasons, "manifest_scope_missing")
	}
	if model.DependencyGateResult != Point11ValDDependencyStateActive {
		reasons = append(reasons, "manifest_dependency_gate_result_invalid")
	}
	if model.IntegratedInvariantResult != Point11ValDIntegratedInvariantStateActive {
		reasons = append(reasons, "manifest_integrated_invariant_result_invalid")
	}
	if model.EvidenceGovernanceQualityResult != Point11ValDQualityMapStateActive {
		reasons = append(reasons, "manifest_quality_result_invalid")
	}
	if model.PublicationBoundaryResult != Point11ValDPublicationReviewStateActive {
		reasons = append(reasons, "manifest_publication_result_invalid")
	}
	if model.NoOverclaimResult != Point11ValDNoOverclaimReviewStateActive {
		reasons = append(reasons, "manifest_no_overclaim_result_invalid")
	}
	if model.CleanRoomIPResult != Point11ValDCleanRoomIPReviewStateActive {
		reasons = append(reasons, "manifest_clean_room_ip_result_invalid")
	}
	if model.CLBClosureResult != Point11ValDCLBClosureStateActive {
		reasons = append(reasons, "manifest_clb_closure_result_invalid")
	}
	if !point11ValDStringListValid(model.CommandsRun, point11ValDCommandRunRefValid) {
		reasons = append(reasons, "manifest_commands_run_missing")
	}
	if !point11ValDStringListValid(model.TestsRun, point11ValDTestRunRefValid) {
		reasons = append(reasons, "manifest_tests_run_missing")
	}
	if !point11ValDStringListValid(model.GrepsRun, point11ValDGrepRunRefValid) {
		reasons = append(reasons, "manifest_greps_run_missing")
	}
	if !point11ValDStringListValid(model.NegativeFixturesRun, point11ValDNegativeFixtureRunRefValid) {
		reasons = append(reasons, "manifest_negative_fixtures_missing")
	}
	if !point11Val0IdentityValueValid(model.EvidenceIdentitySummary) ||
		!point11Val0IdentityValueValid(model.PolicyClaimGovernanceSummary) ||
		!point11Val0IdentityValueValid(model.ExceptionEmergencySummary) ||
		!point11Val0IdentityValueValid(model.ABACEnforcementSummary) ||
		!point11Val0IdentityValueValid(model.DashboardProjectionSummary) {
		reasons = append(reasons, "manifest_summaries_missing")
	}
	if model.ProjectionBoundaryResultToken != point11ValDProjectionBoundaryResultActive {
		reasons = append(reasons, "manifest_projection_boundary_result_invalid")
	}
	if model.ReviewerResult != point11ValDReviewerResultPassConfirmed {
		reasons = append(reasons, "manifest_reviewer_not_pass_confirmed")
	}
	if !model.Point11PassAllowed {
		reasons = append(reasons, "manifest_point11_pass_not_allowed")
	}
	if model.Point11PassToken != point11ValDPoint11PassToken {
		reasons = append(reasons, "manifest_point11_pass_token_invalid")
	}
	if point11ValDContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "manifest_overclaim_detected")
	}
	if foundation.DependencyState != Point11ValDDependencyStateActive ||
		foundation.IntegratedInvariantState != Point11ValDIntegratedInvariantStateActive ||
		foundation.QualityMapState != Point11ValDQualityMapStateActive ||
		foundation.PublicationReviewState != Point11ValDPublicationReviewStateActive ||
		foundation.NoOverclaimReviewState != Point11ValDNoOverclaimReviewStateActive ||
		foundation.CleanRoomIPReviewState != Point11ValDCleanRoomIPReviewStateActive ||
		foundation.CLBClosureState != Point11ValDCLBClosureStateActive {
		reasons = append(reasons, "manifest_foundation_gates_not_active")
	}
	if len(reasons) > 0 {
		return Point11ValDPassClosureManifestStateBlocked, reasons
	}
	return Point11ValDPassClosureManifestStateActive, nil
}

func EvaluatePoint11ValDPassClosureManifestState(model Point11ValDPassClosureManifest, foundation Point11ValDFoundation) string {
	state, _ := point11ValDPassClosureManifestStateAndReasons(model, foundation)
	return state
}

func point11ValDFinalPassGateStateAndReasons(model Point11ValDFinalPoint11PassGate, foundation Point11ValDFoundation) (string, []string) {
	reasons := []string{}
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) {
		reasons = append(reasons, "final_pass_gate_projection_disclaimer_blocked")
	}
	if !point11ValDFinalGateRefValid(model.FinalGateID) {
		reasons = append(reasons, "final_pass_gate_id_invalid")
	}
	if model.DependencyState != Point11ValDDependencyStateActive {
		reasons = append(reasons, "final_pass_gate_dependency_state_invalid")
	}
	if model.InvariantState != Point11ValDIntegratedInvariantStateActive {
		reasons = append(reasons, "final_pass_gate_invariant_state_invalid")
	}
	if model.QualityState != Point11ValDQualityMapStateActive {
		reasons = append(reasons, "final_pass_gate_quality_state_invalid")
	}
	if model.PublicationState != Point11ValDPublicationReviewStateActive {
		reasons = append(reasons, "final_pass_gate_publication_state_invalid")
	}
	if model.NoOverclaimState != Point11ValDNoOverclaimReviewStateActive {
		reasons = append(reasons, "final_pass_gate_no_overclaim_state_invalid")
	}
	if model.CleanRoomIPState != Point11ValDCleanRoomIPReviewStateActive {
		reasons = append(reasons, "final_pass_gate_clean_room_ip_state_invalid")
	}
	if model.CLBClosureState != Point11ValDCLBClosureStateActive {
		reasons = append(reasons, "final_pass_gate_clb_closure_state_invalid")
	}
	if model.ManifestState != Point11ValDPassClosureManifestStateActive {
		reasons = append(reasons, "final_pass_gate_manifest_state_invalid")
	}
	if !model.Point11PassAllowed {
		reasons = append(reasons, "final_pass_gate_point11_pass_not_allowed")
	}
	if !model.Point11PassEmitted {
		reasons = append(reasons, "final_pass_gate_point11_pass_not_emitted")
	}
	if model.Point11PassToken != point11ValDPoint11PassToken {
		reasons = append(reasons, "final_pass_gate_point11_pass_token_invalid")
	}
	if model.Point11PassObservedOutsideFinalClosure {
		reasons = append(reasons, "final_pass_gate_point11_pass_outside_final_closure")
	}
	if point11ValDContainsForbiddenText(strings.Join(model.Diagnostics, " ")) {
		reasons = append(reasons, "final_pass_gate_overclaim_detected")
	}
	if foundation.DependencyState != Point11ValDDependencyStateActive ||
		foundation.IntegratedInvariantState != Point11ValDIntegratedInvariantStateActive ||
		foundation.QualityMapState != Point11ValDQualityMapStateActive ||
		foundation.PublicationReviewState != Point11ValDPublicationReviewStateActive ||
		foundation.NoOverclaimReviewState != Point11ValDNoOverclaimReviewStateActive ||
		foundation.CleanRoomIPReviewState != Point11ValDCleanRoomIPReviewStateActive ||
		foundation.CLBClosureState != Point11ValDCLBClosureStateActive ||
		foundation.PassClosureManifestState != Point11ValDPassClosureManifestStateActive {
		reasons = append(reasons, "final_pass_gate_foundation_gates_not_active")
	}
	if len(reasons) > 0 {
		return Point11ValDFinalPassGateStateBlocked, reasons
	}
	return Point11ValDFinalPassGateStateActive, nil
}

func EvaluatePoint11ValDFinalPassGateState(model Point11ValDFinalPoint11PassGate, foundation Point11ValDFoundation) string {
	state, _ := point11ValDFinalPassGateStateAndReasons(model, foundation)
	return state
}

func point11ValDComponentStates(model Point11ValDFoundation) []string {
	return []string{
		"dependency:" + model.DependencyState,
		"invariant:" + model.IntegratedInvariantState,
		"quality:" + model.QualityMapState,
		"publication:" + model.PublicationReviewState,
		"no_overclaim:" + model.NoOverclaimReviewState,
		"clean_room_ip:" + model.CleanRoomIPReviewState,
		"clb:" + model.CLBClosureState,
		"manifest:" + model.PassClosureManifestState,
		"final_pass_gate:" + model.FinalPassGateState,
	}
}

func point11ValDReasonsMatchExactly(reasons []string, expected string) bool {
	return len(reasons) == 1 && strings.TrimSpace(reasons[0]) == expected
}

func point11ValDBlockingReasons(model Point11ValDFoundation) []string {
	reasons := []string{}
	if model.DependencyState == Point11ValDDependencyStateBlocked {
		reasons = append(reasons, "dependency_blocked")
	}
	if model.IntegratedInvariantState == Point11ValDIntegratedInvariantStateBlocked {
		reasons = append(reasons, "integrated_invariant_blocked")
	}
	if model.QualityMapState == Point11ValDQualityMapStateBlocked {
		reasons = append(reasons, "quality_map_blocked")
	}
	if model.PublicationReviewState == Point11ValDPublicationReviewStateBlocked {
		reasons = append(reasons, "publication_review_blocked")
	}
	if model.NoOverclaimReviewState == Point11ValDNoOverclaimReviewStateBlocked {
		reasons = append(reasons, "no_overclaim_review_blocked")
	}
	if model.CleanRoomIPReviewState == Point11ValDCleanRoomIPReviewStateBlocked {
		reasons = append(reasons, "clean_room_ip_review_blocked")
	}
	if model.CLBClosureState == Point11ValDCLBClosureStateBlocked {
		reasons = append(reasons, "clb_closure_blocked")
	}
	if model.PassClosureManifestState == Point11ValDPassClosureManifestStateBlocked {
		reasons = append(reasons, "pass_closure_manifest_blocked")
	}
	if model.FinalPassGateState == Point11ValDFinalPassGateStateBlocked {
		reasons = append(reasons, "final_pass_gate_blocked")
	}
	if model.CreatesAuthorityClaims {
		reasons = append(reasons, "authority_claim_surface_blocked")
	}
	if model.CreatesPublicationSideEffects {
		reasons = append(reasons, "publication_side_effects_blocked")
	}
	if model.CreatesSigningSideEffects {
		reasons = append(reasons, "signing_side_effects_blocked")
	}
	if model.CreatesAnchoringSideEffects {
		reasons = append(reasons, "anchoring_side_effects_blocked")
	}
	if model.CreatesExternalAPISideEffects {
		reasons = append(reasons, "external_api_side_effects_blocked")
	}
	if model.CreatesProductionSideEffects {
		reasons = append(reasons, "production_side_effects_blocked")
	}
	return reasons
}

func point11ValDDiagnosticsModel(
	model Point11ValDFoundation,
	dependencyReasons []string,
	integratedReasons []string,
	qualityReasons []string,
	publicationReasons []string,
	noOverclaimReasons []string,
	cleanRoomReasons []string,
	clbReasons []string,
	manifestReasons []string,
	finalGateReasons []string,
) Point11ValDDiagnostics {
	return Point11ValDDiagnostics{
		CurrentState:               model.CurrentState,
		BlockingReasons:            append([]string{}, model.BlockingReasons...),
		ReviewPrerequisites:        append([]string{}, model.ReviewPrerequisites...),
		ComponentStates:            point11ValDComponentStates(model),
		DependencyReasons:          append([]string{}, dependencyReasons...),
		IntegratedInvariantReasons: append([]string{}, integratedReasons...),
		QualityMapReasons:          append([]string{}, qualityReasons...),
		PublicationReviewReasons:   append([]string{}, publicationReasons...),
		NoOverclaimReasons:         append([]string{}, noOverclaimReasons...),
		CleanRoomIPReasons:         append([]string{}, cleanRoomReasons...),
		CLBLedgerReasons:           append([]string{}, clbReasons...),
		ManifestReasons:            append([]string{}, manifestReasons...),
		FinalPassGateReasons:       append([]string{}, finalGateReasons...),
		ProjectionDisclaimer:       model.ProjectionDisclaimer,
	}
}

func EvaluatePoint11ValDFoundationState(model Point11ValDFoundation) string {
	dependencyActive := model.DependencyState == Point11ValDDependencyStateActive
	dependencyReviewRequired := model.DependencyState == Point11ValDDependencyStateReviewRequired
	if !point11Val0ValidProjectionDisclaimer(model.ProjectionDisclaimer) ||
		model.CreatesAuthorityClaims ||
		model.CreatesPublicationSideEffects ||
		model.CreatesSigningSideEffects ||
		model.CreatesAnchoringSideEffects ||
		model.CreatesExternalAPISideEffects ||
		model.CreatesProductionSideEffects ||
		(!dependencyActive && !dependencyReviewRequired) ||
		model.IntegratedInvariantState != Point11ValDIntegratedInvariantStateActive ||
		model.QualityMapState != Point11ValDQualityMapStateActive ||
		model.PublicationReviewState != Point11ValDPublicationReviewStateActive ||
		model.NoOverclaimReviewState != Point11ValDNoOverclaimReviewStateActive ||
		model.CleanRoomIPReviewState != Point11ValDCleanRoomIPReviewStateActive ||
		model.CLBClosureState != Point11ValDCLBClosureStateActive ||
		model.PassClosureManifestState != Point11ValDPassClosureManifestStateActive ||
		model.FinalPassGateState != Point11ValDFinalPassGateStateActive {
		return Point11ValDStateBlocked
	}
	if dependencyReviewRequired {
		return Point11ValDStateReviewRequired
	}
	return Point11ValDStateActive
}

func Point11ValDFoundationModel() Point11ValDFoundation {
	disclaimer := point11ValDProjectionDisclaimerBaseline
	return Point11ValDFoundation{
		CurrentState:             Point11ValDStateBlocked,
		DependencyState:          Point11ValDDependencyStateReviewRequired,
		IntegratedInvariantState: Point11ValDIntegratedInvariantStateActive,
		QualityMapState:          Point11ValDQualityMapStateActive,
		PublicationReviewState:   Point11ValDPublicationReviewStateActive,
		NoOverclaimReviewState:   Point11ValDNoOverclaimReviewStateActive,
		CleanRoomIPReviewState:   Point11ValDCleanRoomIPReviewStateActive,
		CLBClosureState:          Point11ValDCLBClosureStateActive,
		PassClosureManifestState: Point11ValDPassClosureManifestStateActive,
		FinalPassGateState:       Point11ValDFinalPassGateStateActive,
		ProjectionDisclaimer:     disclaimer,
		Val0Dependency:           SnapshotPoint11ValDVal0DependencyFromComputed(ComputePoint11Val0Foundation(Point11Val0FoundationModel()), point11ValDDefaultVal0ReviewContext()),
		ValADependency:           SnapshotPoint11ValDValADependencyFromComputed(ComputePoint11ValAFoundation(Point11ValAFoundationModel()), point11ValDDefaultValAReviewContext()),
		ValBDependency:           SnapshotPoint11ValDValBDependencyFromComputed(ComputePoint11ValBFoundation(Point11ValBFoundationModel()), point11ValDDefaultValBReviewContext()),
		ValCDependency:           SnapshotPoint11ValDValCDependencyFromComputed(ComputePoint11ValCFoundation(Point11ValCFoundationModel()), point11ValDDefaultValCReviewContext()),
		IntegratedInvariantReview: Point11ValDIntegratedGovernanceInvariantReview{
			InvariantReviewID:                  "closure_point11_vald_invariant_001",
			Val0Ref:                            "dependency_review_point11_vald_val0_001",
			ValARef:                            "dependency_review_point11_vald_vala_001",
			ValBRef:                            "dependency_review_point11_vald_valb_001",
			ValCRef:                            "dependency_review_point11_vald_valc_001",
			PolicyAuthorityConsistencyState:    point11ValDCheckStateActive,
			ClaimAuthorityConsistencyState:     point11ValDCheckStateActive,
			PublicationBoundaryState:           point11ValDCheckStateActive,
			NoOverclaimState:                   point11ValDCheckStateActive,
			CleanRoomIPState:                   point11ValDCheckStateActive,
			ProjectionBoundaryState:            point11ValDCheckStateActive,
			ExceptionEmergencyConsistencyState: point11ValDCheckStateActive,
			ABACEnforcementConsistencyState:    point11ValDCheckStateActive,
			DashboardProjectionState:           point11ValDCheckStateActive,
			Point11PassRuleState:               point11ValDCheckStateActive,
			Diagnostics:                        []string{"integrated_policy_claim_governance_consistent"},
			ProjectionDisclaimer:               disclaimer,
		},
		QualityMap: Point11ValDEvidenceGovernanceQualityMap{
			QualityMapID:          "quality_map_point11_vald_001",
			PolicyRefs:            []string{"policy_point11_vala_authority_core_v1"},
			ClaimRefs:             []string{"claim_point11_valb_customer_scope_001"},
			VerificationRefs:      []string{"verification_point11_valb_claim_001"},
			RegistryRefs:          []string{"claim_registry_point11_valb_core"},
			EnforcementRefs:       []string{"enforcement_point11_valc_decision_001"},
			ExceptionRefs:         []string{"exception_point11_valc_scope_override_001"},
			EmergencyRefs:         []string{"emergency_point11_valc_scope_override_001"},
			MonitoringRefs:        []string{"monitoring_point11_valc_link_001"},
			DashboardRefs:         []string{"dashboard_point11_valc_governance_001"},
			EvidenceRefs:          []string{"evidence:point11-vald-quality-001"},
			EvidenceHashRefs:      []string{"evidence_hash_point11_vald_quality_001"},
			AuditRefs:             []string{"audit_point11_vald_quality_001"},
			GovernanceEventRefs:   []string{"governance_event_point11_vald_quality_001"},
			CleanRoomIPReviewRefs: []string{"clean_room_review_point11_vald_quality_001"},
			FreshnessState:        point11ValDFreshnessStateActive,
			RevocationState:       point11ValDRevocationStateActive,
			SupersessionState:     point11ValDSupersessionStateActive,
			DuplicateState:        point11ValDDuplicateStateActive,
			ConflictState:         point11ValDConflictStateActive,
			UnrelatedState:        point11ValDUnrelatedStateActive,
			TenantScopeState:      point11ValDTenantScopeStateActive,
			Diagnostics:           []string{"evidence_linked_and_tenant_scoped"},
			ProjectionDisclaimer:  disclaimer,
		},
		PublicationReview: Point11ValDPublicationBoundaryFinalReview{
			PublicationReviewID:           "publication_review_point11_vald_001",
			ModeledSurfaces:               []string{point11Val0PublicationSurfaceDocs, point11Val0PublicationSurfacePortal, point11Val0PublicationSurfaceExport},
			CustomerVisibleSurfaces:       []string{point11Val0PublicationSurfacePortal},
			PublicSurfaces:                []string{point11Val0PublicationSurfaceDocs},
			ExportSurfaces:                []string{point11Val0PublicationSurfaceExport},
			PartnerSurfaces:               []string{point11Val0PublicationSurfacePartner},
			BuyerSurfaces:                 []string{point11Val0PublicationSurfaceBuyer},
			AgentOutputSurfaces:           []string{point11ValDPublicationSurfaceAgentOutput},
			CleanRoomIPReviewRefs:         []string{"clean_room_review_point11_vald_publication_001"},
			GovernanceEventRefs:           []string{"governance_event_point11_vald_publication_001"},
			CreatesPublicationSideEffects: false,
			CreatesCustomerFacingMaterial: false,
			CreatesAuthorityClaim:         false,
			CreatesCertificationClaim:     false,
			CreatesRegulatoryClaim:        false,
			CreatesComplianceGuarantee:    false,
			Diagnostics:                   []string{"modeled_surfaces_only"},
			ProjectionDisclaimer:          disclaimer,
		},
		NoOverclaimReview: Point11ValDFinalNoOverclaimReview{
			NoOverclaimReviewID:     "no_overclaim_review_point11_vald_001",
			ObservedClaims:          []string{"bounded claim"},
			ObservedDiagnostics:     []string{"policy_bound_decision_support"},
			ObservedDashboardText:   []string{"advisory projection only"},
			ObservedPublicationText: []string{"not a certification"},
			Denylist:                point11ValDNoOverclaimDenylist(),
			SafeWordingExamples:     point11ValDSafeWordingExamples(),
			ReviewState:             Point11ValDNoOverclaimReviewStateActive,
			Diagnostics:             []string{"no_forbidden_claims_outside_denylist"},
			ProjectionDisclaimer:    disclaimer,
		},
		CleanRoomIPReview: Point11ValDCleanRoomIPFinalReview{
			CleanRoomReviewID:        "clean_room_review_point11_vald_001",
			PublicClaimRefs:          []string{"claim_point11_valb_customer_scope_001"},
			BuyerClaimRefs:           []string{"claim_point11_valb_buyer_scope_001"},
			PartnerClaimRefs:         []string{"claim_point11_valb_partner_scope_001"},
			CustomerVisibleClaimRefs: []string{"claim_point11_valb_customer_scope_001"},
			ThirdPartyRefs:           []string{"third_party_point11_vald_dependency_001"},
			LicenseReviewRefs:        []string{"license_review_point11_vald_001"},
			IPReviewRefs:             []string{"ip_review_point11_vald_001"},
			ExternalLegalReviewRef:   "external_legal_review_point11_vald_001",
			ExternalFTOReviewRef:     "external_fto_review_point11_vald_001",
			Diagnostics:              []string{"clean_room_and_ip_review_active"},
			ProjectionDisclaimer:     disclaimer,
		},
		CLBLedger: Point11ValDCLBClosureLedger{
			ClosureLedgerID:      "clb_closure_point11_vald_001",
			CLB0Findings:         nil,
			CLB1Findings:         nil,
			CLB2Findings:         nil,
			CLB3Findings:         []string{"finding_point11_vald_advisory_001"},
			ResolvedFindings:     []string{"resolved_finding_point11_vald_001"},
			AcceptedRisks:        nil,
			DeferredItems:        []string{"deferred_item_point11_twelve_future_wave"},
			ReviewerResult:       point11ValDReviewerResultPassConfirmed,
			Diagnostics:          []string{"clb3_advisory_only"},
			ProjectionDisclaimer: disclaimer,
		},
		PassClosureManifest: Point11ValDPassClosureManifest{
			ManifestID:                      "manifest_point11_vald_001",
			PointID:                         point11ValDPointID,
			WaveID:                          point11ValDWaveID,
			Scope:                           point11ValDScope,
			DependencyGateResult:            Point11ValDDependencyStateActive,
			IntegratedInvariantResult:       Point11ValDIntegratedInvariantStateActive,
			EvidenceGovernanceQualityResult: Point11ValDQualityMapStateActive,
			PublicationBoundaryResult:       Point11ValDPublicationReviewStateActive,
			NoOverclaimResult:               Point11ValDNoOverclaimReviewStateActive,
			CleanRoomIPResult:               Point11ValDCleanRoomIPReviewStateActive,
			CLBClosureResult:                Point11ValDCLBClosureStateActive,
			CommandsRun:                     []string{"command_run_point11_vald_001"},
			TestsRun:                        []string{"test_run_point11_vald_001"},
			GrepsRun:                        []string{"grep_run_point11_vald_001"},
			NegativeFixturesRun:             []string{"negative_fixture_point11_vald_001"},
			EvidenceIdentitySummary:         "evidence_identity_exact_and_tenant_scoped",
			PolicyClaimGovernanceSummary:    "policy_claim_governance_consistent",
			ExceptionEmergencySummary:       "exception_emergency_scoped_time_bound_monitored_revocable",
			ABACEnforcementSummary:          "abac_deny_over_allow_enforced",
			DashboardProjectionSummary:      "dashboard_projection_bounded",
			ProjectionBoundaryResultToken:   point11ValDProjectionBoundaryResultActive,
			ReviewerResult:                  point11ValDReviewerResultPassConfirmed,
			Point11PassAllowed:              true,
			Point11PassToken:                point11ValDPoint11PassToken,
			Diagnostics:                     []string{"pass_closure_manifest_complete"},
			ProjectionDisclaimer:            disclaimer,
		},
		FinalPassGate: Point11ValDFinalPoint11PassGate{
			FinalGateID:          "final_gate_point11_vald_001",
			DependencyState:      Point11ValDDependencyStateActive,
			InvariantState:       Point11ValDIntegratedInvariantStateActive,
			QualityState:         Point11ValDQualityMapStateActive,
			PublicationState:     Point11ValDPublicationReviewStateActive,
			NoOverclaimState:     Point11ValDNoOverclaimReviewStateActive,
			CleanRoomIPState:     Point11ValDCleanRoomIPReviewStateActive,
			CLBClosureState:      Point11ValDCLBClosureStateActive,
			ManifestState:        Point11ValDPassClosureManifestStateActive,
			Point11PassAllowed:   true,
			Point11PassEmitted:   true,
			Point11PassToken:     point11ValDPoint11PassToken,
			Diagnostics:          []string{"final_pass_gate_active"},
			ProjectionDisclaimer: disclaimer,
		},
	}
}

func ComputePoint11ValDFoundation(model Point11ValDFoundation) Point11ValDFoundation {
	dependencyBundle := Point11ValDDependencyBundle{
		Val0: model.Val0Dependency,
		ValA: model.ValADependency,
		ValB: model.ValBDependency,
		ValC: model.ValCDependency,
	}
	dependencyState, dependencyReasons := point11ValDDependencyStateAndReasons(dependencyBundle)
	invariantState, invariantReasons := point11ValDIntegratedInvariantStateAndReasons(model.IntegratedInvariantReview)
	qualityState, qualityReasons := point11ValDQualityMapStateAndReasons(model.QualityMap)
	publicationState, publicationReasons := point11ValDPublicationReviewStateAndReasons(model.PublicationReview)
	noOverclaimState, noOverclaimReasons := point11ValDNoOverclaimReviewStateAndReasons(model.NoOverclaimReview)
	cleanRoomState, cleanRoomReasons := point11ValDCleanRoomIPReviewStateAndReasons(model.CleanRoomIPReview)
	clbState, clbReasons := point11ValDCLBClosureStateAndReasons(model.CLBLedger)

	model.DependencyState = dependencyState
	model.IntegratedInvariantState = invariantState
	model.QualityMapState = qualityState
	model.PublicationReviewState = publicationState
	model.NoOverclaimReviewState = noOverclaimState
	model.CleanRoomIPReviewState = cleanRoomState
	model.CLBClosureState = clbState

	manifestState, manifestReasons := point11ValDPassClosureManifestStateAndReasons(model.PassClosureManifest, model)
	model.PassClosureManifestState = manifestState
	model.PassClosureManifest.CurrentState = manifestState
	finalGateState, finalGateReasons := point11ValDFinalPassGateStateAndReasons(model.FinalPassGate, model)
	model.FinalPassGateState = finalGateState
	model.FinalPassGate.CurrentState = finalGateState

	reviewOnlyDueToUpstreamPrereq := model.DependencyState == Point11ValDDependencyStateReviewRequired &&
		model.IntegratedInvariantState == Point11ValDIntegratedInvariantStateActive &&
		model.QualityMapState == Point11ValDQualityMapStateActive &&
		model.PublicationReviewState == Point11ValDPublicationReviewStateActive &&
		model.NoOverclaimReviewState == Point11ValDNoOverclaimReviewStateActive &&
		model.CleanRoomIPReviewState == Point11ValDCleanRoomIPReviewStateActive &&
		model.CLBClosureState == Point11ValDCLBClosureStateActive &&
		!model.CreatesAuthorityClaims &&
		!model.CreatesPublicationSideEffects &&
		!model.CreatesSigningSideEffects &&
		!model.CreatesAnchoringSideEffects &&
		!model.CreatesExternalAPISideEffects &&
		!model.CreatesProductionSideEffects &&
		point11ValDReasonsMatchExactly(manifestReasons, "manifest_foundation_gates_not_active") &&
		point11ValDReasonsMatchExactly(finalGateReasons, "final_pass_gate_foundation_gates_not_active")

	if reviewOnlyDueToUpstreamPrereq {
		model.CurrentState = Point11ValDStateReviewRequired
		model.BlockingReasons = nil
	} else {
		model.CurrentState = EvaluatePoint11ValDFoundationState(model)
		model.BlockingReasons = point11ValDBlockingReasons(model)
	}
	model.ReviewPrerequisites = nil
	for _, prereq := range model.Val0Dependency.ReviewPrerequisites {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, prereq)
	}
	for _, prereq := range model.ValADependency.ReviewPrerequisites {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, prereq)
	}
	for _, prereq := range model.ValBDependency.ReviewPrerequisites {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, prereq)
	}
	for _, prereq := range model.ValCDependency.ReviewPrerequisites {
		model.ReviewPrerequisites = append(model.ReviewPrerequisites, prereq)
	}
	if model.CurrentState == Point11ValDStateActive && model.FinalPassGateState == Point11ValDFinalPassGateStateActive {
		model.Point11PassToken = point11ValDPoint11PassToken
	} else {
		model.Point11PassToken = ""
	}
	model.Diagnostics = point11ValDDiagnosticsModel(
		model,
		dependencyReasons,
		invariantReasons,
		qualityReasons,
		publicationReasons,
		noOverclaimReasons,
		cleanRoomReasons,
		clbReasons,
		manifestReasons,
		finalGateReasons,
	)
	return model
}

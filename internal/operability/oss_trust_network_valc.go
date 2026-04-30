package operability

import "strings"

const (
	OSSTrustNetworkValCStateActive     = "oss_trust_network_valc_active"
	OSSTrustNetworkValCStatePartial    = "oss_trust_network_valc_partial"
	OSSTrustNetworkValCStateIncomplete = "oss_trust_network_valc_incomplete"
	OSSTrustNetworkValCStateBlocked    = "oss_trust_network_valc_blocked"
	OSSTrustNetworkValCStateUnknown    = "oss_trust_network_valc_unknown"

	OSSTrustNetworkValCDependencyStateActive     = "oss_trust_network_valc_dependency_active"
	OSSTrustNetworkValCDependencyStatePartial    = "oss_trust_network_valc_dependency_partial"
	OSSTrustNetworkValCDependencyStateIncomplete = "oss_trust_network_valc_dependency_incomplete"
	OSSTrustNetworkValCDependencyStateBlocked    = "oss_trust_network_valc_dependency_blocked"
	OSSTrustNetworkValCDependencyStateUnknown    = "oss_trust_network_valc_dependency_unknown"

	OSSTrustNetworkValCTrustVisibilityStateActive     = "oss_trust_network_valc_trust_visibility_active"
	OSSTrustNetworkValCTrustVisibilityStatePartial    = "oss_trust_network_valc_trust_visibility_partial"
	OSSTrustNetworkValCTrustVisibilityStateIncomplete = "oss_trust_network_valc_trust_visibility_incomplete"
	OSSTrustNetworkValCTrustVisibilityStateBlocked    = "oss_trust_network_valc_trust_visibility_blocked"
	OSSTrustNetworkValCTrustVisibilityStateUnknown    = "oss_trust_network_valc_trust_visibility_unknown"

	OSSTrustNetworkValCPackageTrustStatusStateActive     = "oss_trust_network_valc_package_trust_status_active"
	OSSTrustNetworkValCPackageTrustStatusStatePartial    = "oss_trust_network_valc_package_trust_status_partial"
	OSSTrustNetworkValCPackageTrustStatusStateIncomplete = "oss_trust_network_valc_package_trust_status_incomplete"
	OSSTrustNetworkValCPackageTrustStatusStateBlocked    = "oss_trust_network_valc_package_trust_status_blocked"
	OSSTrustNetworkValCPackageTrustStatusStateUnknown    = "oss_trust_network_valc_package_trust_status_unknown"

	OSSTrustNetworkValCExportBoundaryStateActive     = "oss_trust_network_valc_export_boundary_active"
	OSSTrustNetworkValCExportBoundaryStatePartial    = "oss_trust_network_valc_export_boundary_partial"
	OSSTrustNetworkValCExportBoundaryStateIncomplete = "oss_trust_network_valc_export_boundary_incomplete"
	OSSTrustNetworkValCExportBoundaryStateBlocked    = "oss_trust_network_valc_export_boundary_blocked"
	OSSTrustNetworkValCExportBoundaryStateUnknown    = "oss_trust_network_valc_export_boundary_unknown"

	OSSTrustNetworkValCRemediationSuggestionStateActive     = "oss_trust_network_valc_remediation_suggestion_active"
	OSSTrustNetworkValCRemediationSuggestionStatePartial    = "oss_trust_network_valc_remediation_suggestion_partial"
	OSSTrustNetworkValCRemediationSuggestionStateIncomplete = "oss_trust_network_valc_remediation_suggestion_incomplete"
	OSSTrustNetworkValCRemediationSuggestionStateBlocked    = "oss_trust_network_valc_remediation_suggestion_blocked"
	OSSTrustNetworkValCRemediationSuggestionStateUnknown    = "oss_trust_network_valc_remediation_suggestion_unknown"

	OSSTrustNetworkValCPRProposalStateActive     = "oss_trust_network_valc_pr_proposal_active"
	OSSTrustNetworkValCPRProposalStatePartial    = "oss_trust_network_valc_pr_proposal_partial"
	OSSTrustNetworkValCPRProposalStateIncomplete = "oss_trust_network_valc_pr_proposal_incomplete"
	OSSTrustNetworkValCPRProposalStateBlocked    = "oss_trust_network_valc_pr_proposal_blocked"
	OSSTrustNetworkValCPRProposalStateUnknown    = "oss_trust_network_valc_pr_proposal_unknown"

	OSSTrustNetworkValCLocalOverrideStateActive     = "oss_trust_network_valc_local_override_active"
	OSSTrustNetworkValCLocalOverrideStatePartial    = "oss_trust_network_valc_local_override_partial"
	OSSTrustNetworkValCLocalOverrideStateIncomplete = "oss_trust_network_valc_local_override_incomplete"
	OSSTrustNetworkValCLocalOverrideStateBlocked    = "oss_trust_network_valc_local_override_blocked"
	OSSTrustNetworkValCLocalOverrideStateUnknown    = "oss_trust_network_valc_local_override_unknown"

	OSSTrustNetworkValCRemediationSafetyStateActive     = "oss_trust_network_valc_remediation_safety_active"
	OSSTrustNetworkValCRemediationSafetyStatePartial    = "oss_trust_network_valc_remediation_safety_partial"
	OSSTrustNetworkValCRemediationSafetyStateIncomplete = "oss_trust_network_valc_remediation_safety_incomplete"
	OSSTrustNetworkValCRemediationSafetyStateBlocked    = "oss_trust_network_valc_remediation_safety_blocked"
	OSSTrustNetworkValCRemediationSafetyStateUnknown    = "oss_trust_network_valc_remediation_safety_unknown"

	OSSTrustNetworkValCEcosystemConsistencyStateActive     = "oss_trust_network_valc_ecosystem_consistency_active"
	OSSTrustNetworkValCEcosystemConsistencyStatePartial    = "oss_trust_network_valc_ecosystem_consistency_partial"
	OSSTrustNetworkValCEcosystemConsistencyStateIncomplete = "oss_trust_network_valc_ecosystem_consistency_incomplete"
	OSSTrustNetworkValCEcosystemConsistencyStateBlocked    = "oss_trust_network_valc_ecosystem_consistency_blocked"
	OSSTrustNetworkValCEcosystemConsistencyStateUnknown    = "oss_trust_network_valc_ecosystem_consistency_unknown"

	OSSTrustNetworkValCNoOverclaimStateActive     = "oss_trust_network_valc_no_overclaim_active"
	OSSTrustNetworkValCNoOverclaimStatePartial    = "oss_trust_network_valc_no_overclaim_partial"
	OSSTrustNetworkValCNoOverclaimStateIncomplete = "oss_trust_network_valc_no_overclaim_incomplete"
	OSSTrustNetworkValCNoOverclaimStateBlocked    = "oss_trust_network_valc_no_overclaim_blocked"
	OSSTrustNetworkValCNoOverclaimStateUnknown    = "oss_trust_network_valc_no_overclaim_unknown"

	OSSTrustNetworkValCVisibilityVisible     = "visible"
	OSSTrustNetworkValCVisibilityLimited     = "limited"
	OSSTrustNetworkValCVisibilityHidden      = "hidden"
	OSSTrustNetworkValCVisibilityUnsupported = "unsupported"
	OSSTrustNetworkValCVisibilityStale       = "stale"
	OSSTrustNetworkValCVisibilityUnknown     = "unknown"

	OSSTrustNetworkValCPackageStatusReviewedSignalAvailable  = "reviewed_signal_available"
	OSSTrustNetworkValCPackageStatusCandidateSignalAvailable = "candidate_signal_available"
	OSSTrustNetworkValCPackageStatusLocalReviewNeeded        = "local_review_needed"
	OSSTrustNetworkValCPackageStatusSupersededSignal         = "superseded_signal"
	OSSTrustNetworkValCPackageStatusRevokedSignal            = "revoked_signal"
	OSSTrustNetworkValCPackageStatusUnsupportedSignal        = "unsupported_signal"
	OSSTrustNetworkValCPackageStatusUnknownSignal            = "unknown_signal"

	OSSTrustNetworkValCExportClassInternalOperatorView   = "internal_operator_view"
	OSSTrustNetworkValCExportClassEnterpriseCustomerView = "enterprise_customer_view"
	OSSTrustNetworkValCExportClassMaintainerView         = "maintainer_view"
	OSSTrustNetworkValCExportClassPartnerView            = "partner_view"
	OSSTrustNetworkValCExportClassPublicSummaryView      = "public_summary_view"
	OSSTrustNetworkValCExportClassUnsupportedView        = "unsupported_view"

	OSSTrustNetworkValCSuggestionClassVersionUpgrade    = "version_upgrade_suggestion"
	OSSTrustNetworkValCSuggestionClassPinOrHold         = "pin_or_hold_suggestion"
	OSSTrustNetworkValCSuggestionClassReplaceDependency = "replace_dependency_suggestion"
	OSSTrustNetworkValCSuggestionClassMaintainerContact = "maintainer_contact_suggestion"
	OSSTrustNetworkValCSuggestionClassReviewRequired    = "review_required_suggestion"
	OSSTrustNetworkValCSuggestionClassNoAction          = "no_action_suggestion"
	OSSTrustNetworkValCSuggestionClassUnsupported       = "unsupported_suggestion"

	OSSTrustNetworkValCProposalStateProposalReady = "proposal_ready"
	OSSTrustNetworkValCProposalStateNeedsReview   = "needs_review"
	OSSTrustNetworkValCProposalStateUnsupported   = "unsupported"
	OSSTrustNetworkValCProposalStateBlocked       = "blocked"
	OSSTrustNetworkValCProposalStateUnknown       = "unknown"

	OSSTrustNetworkValCOverrideStateNoOverride             = "no_override"
	OSSTrustNetworkValCOverrideStateOverridePresent        = "override_present"
	OSSTrustNetworkValCOverrideStateOverrideRequiresReview = "override_requires_review"
	OSSTrustNetworkValCOverrideStateOverrideRejected       = "override_rejected"
	OSSTrustNetworkValCOverrideStateUnsupported            = "unsupported"
	OSSTrustNetworkValCOverrideStateUnknown                = "unknown"

	OSSTrustNetworkValCRiskClassLow    = "low"
	OSSTrustNetworkValCRiskClassMedium = "medium"
	OSSTrustNetworkValCRiskClassHigh   = "high"
)

type OSSTrustNetworkValCDependencySnapshot struct {
	ValBCurrentState                string   `json:"valb_current_state"`
	ValBPoint9State                 string   `json:"valb_point_9_state"`
	ValBDependencyState             string   `json:"valb_dependency_state"`
	ValBCandidateSignalIntakeState  string   `json:"valb_candidate_signal_intake_state"`
	ValBReviewWorkflowState         string   `json:"valb_review_workflow_state"`
	ValBSharedVEXTriageState        string   `json:"valb_shared_vex_triage_state"`
	ValBSourceWeightingState        string   `json:"valb_source_weighting_state"`
	ValBLocalApplicabilityState     string   `json:"valb_local_applicability_state"`
	ValBPropagationExchangeState    string   `json:"valb_propagation_exchange_state"`
	ValBSupersessionRevocationState string   `json:"valb_supersession_revocation_state"`
	ValBReviewerAuditabilityState   string   `json:"valb_reviewer_auditability_state"`
	ValBNoOverclaimState            string   `json:"valb_no_overclaim_state"`
	ValBProofSurfaceRefs            []string `json:"valb_proof_surface_refs,omitempty"`
	ValBEvidenceRefs                []string `json:"valb_evidence_refs,omitempty"`
	ValBProjectionDisclaimer        string   `json:"valb_projection_disclaimer"`
}

type OSSTrustNetworkValCTrustVisibility struct {
	CurrentState             string   `json:"current_state"`
	VisibilityProfileID      string   `json:"visibility_profile_id"`
	PackageOrProjectIdentity string   `json:"package_or_project_identity"`
	RegistryOrEcosystem      string   `json:"registry_or_ecosystem"`
	ReleaseOrVersionRef      string   `json:"release_or_version_ref"`
	VisibilityState          string   `json:"visibility_state"`
	ReviewedSignalState      string   `json:"reviewed_signal_state"`
	LocalApplicabilityState  string   `json:"local_applicability_state"`
	SourceWeightingState     string   `json:"source_weighting_state"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	FreshnessState           string   `json:"freshness_state"`
	Caveats                  []string `json:"caveats,omitempty"`
	PackageSafetyClaim       bool     `json:"package_safety_claim"`
	CertifiedPackageClaim    bool     `json:"certified_package_claim"`
	ProductionApprovalClaim  bool     `json:"production_approval_claim"`
	DeploymentApprovalClaim  bool     `json:"deployment_approval_claim"`
	GlobalTruthClaim         bool     `json:"global_truth_claim"`
	EnterpriseOverrideClaim  bool     `json:"enterprise_override_claim"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValCPackageTrustStatus struct {
	CurrentState             string   `json:"current_state"`
	StatusSummaryID          string   `json:"status_summary_id"`
	PackageOrProjectIdentity string   `json:"package_or_project_identity"`
	ReleaseOrVersionRef      string   `json:"release_or_version_ref"`
	StatusClass              string   `json:"status_class"`
	ReviewedSignalState      string   `json:"reviewed_signal_state"`
	LocalApplicabilityState  string   `json:"local_applicability_state"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	FreshnessState           string   `json:"freshness_state"`
	Caveats                  []string `json:"caveats,omitempty"`
	ReplacementRef           string   `json:"replacement_ref"`
	DisplayedAsReviewed      bool     `json:"displayed_as_reviewed"`
	UniversalTrustScoreClaim bool     `json:"universal_trust_score_claim"`
	IntegrityScoreClaim      bool     `json:"integrity_score_claim"`
	BadgeScoreClaim          bool     `json:"badge_score_claim"`
	GenericSafetyClaim       bool     `json:"generic_safety_claim"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValCExportBoundary struct {
	CurrentState                string   `json:"current_state"`
	BoundaryID                  string   `json:"boundary_id"`
	ExportClass                 string   `json:"export_class"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	Caveats                     []string `json:"caveats,omitempty"`
	CanonicalInternalExposure   bool     `json:"canonical_internal_exposure"`
	CertificationClaim          bool     `json:"certification_claim"`
	ApprovalClaim               bool     `json:"approval_claim"`
	GlobalBlocklistClaim        bool     `json:"global_blocklist_claim"`
	RedactionStripsCaveats      bool     `json:"redaction_strips_caveats"`
	RedactionStripsEvidenceRefs bool     `json:"redaction_strips_evidence_refs"`
	CandidatePromotedToReviewed bool     `json:"candidate_promoted_to_reviewed"`
	RejectedPromotedToActive    bool     `json:"rejected_promoted_to_active"`
	RevokedPromotedToActive     bool     `json:"revoked_promoted_to_active"`
	UnknownPromotedToActive     bool     `json:"unknown_promoted_to_active"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValCRemediationSuggestion struct {
	CurrentState              string   `json:"current_state"`
	SuggestionID              string   `json:"suggestion_id"`
	SuggestionClass           string   `json:"suggestion_class"`
	PackageOrProjectIdentity  string   `json:"package_or_project_identity"`
	AffectedReleaseOrVersion  string   `json:"affected_release_or_version"`
	TargetReleaseOrVersion    string   `json:"target_release_or_version"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	ConfidenceClass           string   `json:"confidence_class"`
	CompatibilityNote         string   `json:"compatibility_note"`
	RiskNote                  string   `json:"risk_note"`
	LocalApplicabilityNote    string   `json:"local_applicability_note"`
	Rationale                 string   `json:"rationale"`
	Caveats                   []string `json:"caveats,omitempty"`
	DependencyMutationAttempt bool     `json:"dependency_mutation_attempt"`
	PolicyOverrideAttempt     bool     `json:"policy_override_attempt"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValCPRProposal struct {
	CurrentState           string   `json:"current_state"`
	ProposalID             string   `json:"proposal_id"`
	ProposalState          string   `json:"proposal_state"`
	EvidenceRefs           []string `json:"evidence_refs,omitempty"`
	CompatibilityNote      string   `json:"compatibility_note"`
	ReviewerRequired       bool     `json:"reviewer_required"`
	NoAutomerge            bool     `json:"no_automerge"`
	NoHiddenMutation       bool     `json:"no_hidden_mutation"`
	LocalApplicabilityNote string   `json:"local_applicability_note"`
	Caveats                []string `json:"caveats,omitempty"`
	AdvisoryOnly           bool     `json:"advisory_only"`
	BranchWrite            bool     `json:"branch_write"`
	NetworkAction          bool     `json:"network_action"`
	DependencyMutation     bool     `json:"dependency_mutation"`
	PRCreation             bool     `json:"pr_creation"`
	AutoMerge              bool     `json:"auto_merge"`
	ProjectionDisclaimer   string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValCLocalOverride struct {
	CurrentState                        string   `json:"current_state"`
	OverrideID                          string   `json:"override_id"`
	OverrideState                       string   `json:"override_state"`
	EvidenceRefs                        []string `json:"evidence_refs,omitempty"`
	Rationale                           string   `json:"rationale"`
	Scope                               string   `json:"scope"`
	OwnerOrReviewerClass                string   `json:"owner_or_reviewer_class"`
	Caveats                             []string `json:"caveats,omitempty"`
	LocalOnlyBoundary                   bool     `json:"local_only_boundary"`
	RewriteCanonicalEvidence            bool     `json:"rewrite_canonical_evidence"`
	SilentlySuppressNetworkIntelligence bool     `json:"silently_suppress_network_intelligence"`
	SharedSignalOverridesLocalDecision  bool     `json:"shared_signal_overrides_local_decision"`
	ProjectionDisclaimer                string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValCRemediationSafety struct {
	CurrentState            string   `json:"current_state"`
	SafetyID                string   `json:"safety_id"`
	SuggestionID            string   `json:"suggestion_id"`
	CompatibilityNote       string   `json:"compatibility_note"`
	TestValidationNote      string   `json:"test_validation_note"`
	RollbackNote            string   `json:"rollback_note"`
	RiskClass               string   `json:"risk_class"`
	ReviewerRequired        bool     `json:"reviewer_required"`
	EvidenceRefs            []string `json:"evidence_refs,omitempty"`
	Caveats                 []string `json:"caveats,omitempty"`
	ProductionApprovalClaim bool     `json:"production_approval_claim"`
	DeploymentApprovalClaim bool     `json:"deployment_approval_claim"`
	HiddenMutationPath      bool     `json:"hidden_mutation_path"`
	ProjectionDisclaimer    string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValCEcosystemConsistency struct {
	CurrentState                       string   `json:"current_state"`
	ConsistencyID                      string   `json:"consistency_id"`
	PackageOrProjectIdentity           string   `json:"package_or_project_identity"`
	ReleaseOrVersionRef                string   `json:"release_or_version_ref"`
	SuggestionPackageOrProjectIdentity string   `json:"suggestion_package_or_project_identity"`
	SuggestionReleaseOrVersion         string   `json:"suggestion_release_or_version"`
	PackageStatusClass                 string   `json:"package_status_class"`
	ReviewState                        string   `json:"review_state"`
	ReviewerDecisionState              string   `json:"reviewer_decision_state"`
	PropagationState                   string   `json:"propagation_state"`
	LocalApplicabilityState            string   `json:"local_applicability_state"`
	DisplayedAsReviewed                bool     `json:"displayed_as_reviewed"`
	DisplayedAsApplicable              bool     `json:"displayed_as_applicable"`
	ReviewedExchangePresentedActive    bool     `json:"reviewed_exchange_presented_active"`
	FreshnessState                     string   `json:"freshness_state"`
	EvidenceRefs                       []string `json:"evidence_refs,omitempty"`
	Caveats                            []string `json:"caveats,omitempty"`
	ProjectionDisclaimer               string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValCNoOverclaim struct {
	CurrentState               string   `json:"current_state"`
	DisciplineID               string   `json:"discipline_id"`
	Version                    string   `json:"version"`
	ObservedClaims             []string `json:"observed_claims,omitempty"`
	GlobalTruthClaim           bool     `json:"global_truth_claim"`
	ReviewedMeansSafeClaim     bool     `json:"reviewed_means_safe_claim"`
	CommunityTruthClaim        bool     `json:"community_truth_claim"`
	NetworkTruthClaim          bool     `json:"network_truth_claim"`
	CertifiedPackageClaim      bool     `json:"certified_package_claim"`
	OfficiallySafePackageClaim bool     `json:"officially_safe_package_claim"`
	RegulatorApproved          bool     `json:"regulator_approved"`
	AuditPassed                bool     `json:"audit_passed"`
	ComplianceGuaranteed       bool     `json:"compliance_guaranteed"`
	ProductionApproved         bool     `json:"production_approved"`
	DeploymentApproved         bool     `json:"deployment_approved"`
	LegalCertification         bool     `json:"legal_certification"`
	PatentCleared              bool     `json:"patent_cleared"`
	FTOPCleared                bool     `json:"fto_cleared"`
	AutoRemediated             bool     `json:"auto_remediated"`
	AutoMerged                 bool     `json:"auto_merged"`
	ProductionAutopatch        bool     `json:"production_autopatch"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValCCore struct {
	CurrentState               string                                   `json:"current_state"`
	Point9State                string                                   `json:"point_9_state"`
	DependencyState            string                                   `json:"dependency_state"`
	TrustVisibilityState       string                                   `json:"trust_visibility_state"`
	PackageTrustStatusState    string                                   `json:"package_trust_status_state"`
	ExportBoundaryState        string                                   `json:"export_boundary_state"`
	RemediationSuggestionState string                                   `json:"remediation_suggestion_state"`
	PRProposalState            string                                   `json:"pr_proposal_state"`
	LocalOverrideState         string                                   `json:"local_override_state"`
	RemediationSafetyState     string                                   `json:"remediation_safety_state"`
	EcosystemConsistencyState  string                                   `json:"ecosystem_consistency_state"`
	NoOverclaimState           string                                   `json:"no_overclaim_state"`
	Dependency                 OSSTrustNetworkValCDependencySnapshot    `json:"dependency"`
	TrustVisibility            OSSTrustNetworkValCTrustVisibility       `json:"trust_visibility"`
	PackageTrustStatus         OSSTrustNetworkValCPackageTrustStatus    `json:"package_trust_status"`
	ExportBoundary             OSSTrustNetworkValCExportBoundary        `json:"export_boundary"`
	RemediationSuggestion      OSSTrustNetworkValCRemediationSuggestion `json:"remediation_suggestion"`
	PRProposal                 OSSTrustNetworkValCPRProposal            `json:"pr_proposal"`
	LocalOverride              OSSTrustNetworkValCLocalOverride         `json:"local_override"`
	RemediationSafety          OSSTrustNetworkValCRemediationSafety     `json:"remediation_safety"`
	EcosystemConsistency       OSSTrustNetworkValCEcosystemConsistency  `json:"ecosystem_consistency"`
	NoOverclaim                OSSTrustNetworkValCNoOverclaim           `json:"no_overclaim"`
	ProofSurfaceRefs           []string                                 `json:"proof_surface_refs,omitempty"`
	EvidenceRefs               []string                                 `json:"evidence_refs,omitempty"`
	BlockingReasons            []string                                 `json:"blocking_reasons,omitempty"`
	WhyPoint9NotComplete       []string                                 `json:"why_point_9_not_complete,omitempty"`
	ProjectionDisclaimer       string                                   `json:"projection_disclaimer"`
}

func ossTrustNetworkValCProjectionDisclaimer() string {
	return "projection_only not_canonical_truth oss_trust_network_valc advisory_projection no_hidden_mutation_path"
}

func ossTrustNetworkValCHasProjectionDisclaimer(value string) bool {
	normalized := strings.ToLower(strings.TrimSpace(value))
	return strings.Contains(normalized, "projection_only") &&
		strings.Contains(normalized, "not_canonical_truth") &&
		strings.Contains(normalized, "oss_trust_network_valc")
}

func ossTrustNetworkValCVisibilityStates() []string {
	return []string{
		OSSTrustNetworkValCVisibilityVisible,
		OSSTrustNetworkValCVisibilityLimited,
		OSSTrustNetworkValCVisibilityHidden,
		OSSTrustNetworkValCVisibilityUnsupported,
		OSSTrustNetworkValCVisibilityStale,
		OSSTrustNetworkValCVisibilityUnknown,
	}
}

func ossTrustNetworkValCPackageStatusClasses() []string {
	return []string{
		OSSTrustNetworkValCPackageStatusReviewedSignalAvailable,
		OSSTrustNetworkValCPackageStatusCandidateSignalAvailable,
		OSSTrustNetworkValCPackageStatusLocalReviewNeeded,
		OSSTrustNetworkValCPackageStatusSupersededSignal,
		OSSTrustNetworkValCPackageStatusRevokedSignal,
		OSSTrustNetworkValCPackageStatusUnsupportedSignal,
		OSSTrustNetworkValCPackageStatusUnknownSignal,
	}
}

func ossTrustNetworkValCExportClasses() []string {
	return []string{
		OSSTrustNetworkValCExportClassInternalOperatorView,
		OSSTrustNetworkValCExportClassEnterpriseCustomerView,
		OSSTrustNetworkValCExportClassMaintainerView,
		OSSTrustNetworkValCExportClassPartnerView,
		OSSTrustNetworkValCExportClassPublicSummaryView,
		OSSTrustNetworkValCExportClassUnsupportedView,
	}
}

func ossTrustNetworkValCSuggestionClasses() []string {
	return []string{
		OSSTrustNetworkValCSuggestionClassVersionUpgrade,
		OSSTrustNetworkValCSuggestionClassPinOrHold,
		OSSTrustNetworkValCSuggestionClassReplaceDependency,
		OSSTrustNetworkValCSuggestionClassMaintainerContact,
		OSSTrustNetworkValCSuggestionClassReviewRequired,
		OSSTrustNetworkValCSuggestionClassNoAction,
		OSSTrustNetworkValCSuggestionClassUnsupported,
	}
}

func ossTrustNetworkValCProposalStates() []string {
	return []string{
		OSSTrustNetworkValCProposalStateProposalReady,
		OSSTrustNetworkValCProposalStateNeedsReview,
		OSSTrustNetworkValCProposalStateUnsupported,
		OSSTrustNetworkValCProposalStateBlocked,
		OSSTrustNetworkValCProposalStateUnknown,
	}
}

func ossTrustNetworkValCLocalOverrideStates() []string {
	return []string{
		OSSTrustNetworkValCOverrideStateNoOverride,
		OSSTrustNetworkValCOverrideStateOverridePresent,
		OSSTrustNetworkValCOverrideStateOverrideRequiresReview,
		OSSTrustNetworkValCOverrideStateOverrideRejected,
		OSSTrustNetworkValCOverrideStateUnsupported,
		OSSTrustNetworkValCOverrideStateUnknown,
	}
}

func ossTrustNetworkValCRiskClasses() []string {
	return []string{
		OSSTrustNetworkValCRiskClassLow,
		OSSTrustNetworkValCRiskClassMedium,
		OSSTrustNetworkValCRiskClassHigh,
	}
}

func OSSTrustNetworkValCProofSurfaceRefs() []string {
	return []string{
		"/v1/oss-trust-network/valc/status",
		"/v1/oss-trust-network/valc/proofs",
	}
}

func OSSTrustNetworkValCProofEvidenceRefs() []string {
	return []string{
		"evidence:ostn-valc-dependency-001",
		"evidence:ostn-valc-visibility-001",
		"evidence:ostn-valc-package-status-001",
		"evidence:ostn-valc-export-boundary-001",
		"evidence:ostn-valc-remediation-suggestion-001",
		"evidence:ostn-valc-pr-proposal-001",
		"evidence:ostn-valc-local-override-001",
		"evidence:ostn-valc-remediation-safety-001",
		"evidence:ostn-valc-ecosystem-consistency-001",
		"evidence:ostn-valc-no-overclaim-001",
		"evidence:ostn-valc-point9-governance-001",
	}
}

type ossTrustNetworkValCExpectedEvidenceMetadata struct {
	EvidenceType string
	Source       string
	Scope        string
}

func ossTrustNetworkValCExpectedEvidenceMetadataByID() map[string]ossTrustNetworkValCExpectedEvidenceMetadata {
	return map[string]ossTrustNetworkValCExpectedEvidenceMetadata{
		"evidence:ostn-valc-dependency-001":             {EvidenceType: "dependency_state", Source: "oss-trust-network/valc/dependency", Scope: "valb_dependency"},
		"evidence:ostn-valc-visibility-001":             {EvidenceType: "trust_visibility", Source: "oss-trust-network/valc/visibility", Scope: "oss_trust_visibility"},
		"evidence:ostn-valc-package-status-001":         {EvidenceType: "package_trust_status", Source: "oss-trust-network/valc/package-status", Scope: "package_trust_status"},
		"evidence:ostn-valc-export-boundary-001":        {EvidenceType: "export_boundary", Source: "oss-trust-network/valc/export-boundary", Scope: "ecosystem_export_boundary"},
		"evidence:ostn-valc-remediation-suggestion-001": {EvidenceType: "remediation_suggestion", Source: "oss-trust-network/valc/remediation-suggestion", Scope: "remediation_suggestion"},
		"evidence:ostn-valc-pr-proposal-001":            {EvidenceType: "pr_proposal", Source: "oss-trust-network/valc/pr-proposal", Scope: "pr_proposal_descriptor"},
		"evidence:ostn-valc-local-override-001":         {EvidenceType: "local_override", Source: "oss-trust-network/valc/local-override", Scope: "local_override_visibility"},
		"evidence:ostn-valc-remediation-safety-001":     {EvidenceType: "remediation_safety", Source: "oss-trust-network/valc/remediation-safety", Scope: "remediation_safety"},
		"evidence:ostn-valc-ecosystem-consistency-001":  {EvidenceType: "ecosystem_consistency", Source: "oss-trust-network/valc/ecosystem-consistency", Scope: "ecosystem_consistency"},
		"evidence:ostn-valc-no-overclaim-001":           {EvidenceType: "no_overclaim", Source: "oss-trust-network/valc/no-overclaim", Scope: "no_overclaim_discipline"},
		"evidence:ostn-valc-point9-governance-001":      {EvidenceType: "state_governance", Source: "oss-trust-network/valc/point9-governance", Scope: "point9_governance"},
	}
}

func ossTrustNetworkValCEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:ostn-valc-dependency-001", EvidenceType: "dependency_state", Source: "oss-trust-network/valc/dependency", Timestamp: "2026-04-30T08:00:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "valb_dependency", Caveats: []string{"Val C depends on exact and active OSTN Val B only."}},
		{EvidenceID: "evidence:ostn-valc-visibility-001", EvidenceType: "trust_visibility", Source: "oss-trust-network/valc/visibility", Timestamp: "2026-04-30T08:01:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "oss_trust_visibility", Caveats: []string{"Visibility remains bounded ecosystem projection and not canonical truth."}},
		{EvidenceID: "evidence:ostn-valc-package-status-001", EvidenceType: "package_trust_status", Source: "oss-trust-network/valc/package-status", Timestamp: "2026-04-30T08:02:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "package_trust_status", Caveats: []string{"Package trust status remains bounded summary only and not a score or safety badge."}},
		{EvidenceID: "evidence:ostn-valc-export-boundary-001", EvidenceType: "export_boundary", Source: "oss-trust-network/valc/export-boundary", Timestamp: "2026-04-30T08:03:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "ecosystem_export_boundary", Caveats: []string{"Export boundaries remain caveated and cannot create partner or public certification claims."}},
		{EvidenceID: "evidence:ostn-valc-remediation-suggestion-001", EvidenceType: "remediation_suggestion", Source: "oss-trust-network/valc/remediation-suggestion", Timestamp: "2026-04-30T08:04:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "remediation_suggestion", Caveats: []string{"Remediation suggestions remain evidence-linked advisory descriptors and not mutation authority."}},
		{EvidenceID: "evidence:ostn-valc-pr-proposal-001", EvidenceType: "pr_proposal", Source: "oss-trust-network/valc/pr-proposal", Timestamp: "2026-04-30T08:05:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "pr_proposal_descriptor", Caveats: []string{"PR proposal descriptors remain reviewer-required metadata with no branch write, no PR creation, and no auto-merge."}},
		{EvidenceID: "evidence:ostn-valc-local-override-001", EvidenceType: "local_override", Source: "oss-trust-network/valc/local-override", Timestamp: "2026-04-30T08:06:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "local_override_visibility", Caveats: []string{"Local override visibility remains evidence-linked local context and cannot suppress reviewed network intelligence silently."}},
		{EvidenceID: "evidence:ostn-valc-remediation-safety-001", EvidenceType: "remediation_safety", Source: "oss-trust-network/valc/remediation-safety", Timestamp: "2026-04-30T08:07:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "remediation_safety", Caveats: []string{"Remediation safety remains advisory-only and blocks hidden mutation or approval claims."}},
		{EvidenceID: "evidence:ostn-valc-ecosystem-consistency-001", EvidenceType: "ecosystem_consistency", Source: "oss-trust-network/valc/ecosystem-consistency", Timestamp: "2026-04-30T08:08:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "ecosystem_consistency", Caveats: []string{"Visibility, remediation, and local context must remain internally consistent and freshness-bounded."}},
		{EvidenceID: "evidence:ostn-valc-no-overclaim-001", EvidenceType: "no_overclaim", Source: "oss-trust-network/valc/no-overclaim", Timestamp: "2026-04-30T08:09:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim_discipline", Caveats: []string{"Val C cannot create marketing badges, hidden mutation paths, or forbidden point pass claims."}},
		{EvidenceID: "evidence:ostn-valc-point9-governance-001", EvidenceType: "state_governance", Source: "oss-trust-network/valc/point9-governance", Timestamp: "2026-04-30T08:10:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point9_governance", Caveats: []string{"point_9_state remains not complete and final pass semantics remain deferred beyond Val C."}},
	}
}

func OSSTrustNetworkValCProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, refs []string) bool {
	if !containsExactTrimmedStringSet(refs, OSSTrustNetworkValCProofEvidenceRefs()...) {
		return false
	}
	expected := ossTrustNetworkValCExpectedEvidenceMetadataByID()
	if len(evidence) != len(expected) {
		return false
	}
	allFresh, stale, ok := verifierEcosystemVal0EvidenceValid(evidence)
	if !ok || !allFresh || stale {
		return false
	}
	seen := map[string]struct{}{}
	for _, evidenceRef := range evidence {
		id := strings.TrimSpace(evidenceRef.EvidenceID)
		if id == "" {
			return false
		}
		expectedMetadata, exists := expected[id]
		if !exists {
			return false
		}
		if _, duplicate := seen[id]; duplicate {
			return false
		}
		seen[id] = struct{}{}
		if strings.TrimSpace(evidenceRef.EvidenceType) != expectedMetadata.EvidenceType ||
			strings.TrimSpace(evidenceRef.Source) != expectedMetadata.Source ||
			strings.TrimSpace(evidenceRef.Scope) != expectedMetadata.Scope {
			return false
		}
	}
	return len(seen) == len(expected)
}

func ossTrustNetworkValCContainsForbiddenClaim(values ...string) bool {
	allowed := map[string]struct{}{
		"bounded ecosystem visibility":    {},
		"package trust visibility":        {},
		"advisory remediation suggestion": {},
		"pr proposal descriptor":          {},
		"reviewer-required proposal":      {},
		"local override visibility":       {},
		"reviewed oss trust signal":       {},
		"candidate oss trust signal":      {},
		"source-weighted reviewed signal": {},
		"bounded reviewed exchange":       {},
		"local applicability context":     {},
		"evidence-linked suggestion":      {},
		"not canonical truth":             {},
		"not formal certification":        {},
		"not production approval":         {},
		"no hidden mutation path":         {},
	}
	// Keep this denylist aligned with the required Val C blocked-phrase set.
	disallowed := []string{
		"changelock verified",
		"certified package",
		"officially safe package",
		"regulator-approved",
		"audit passed",
		"compliance guaranteed",
		"production approved",
		"deployment approved",
		"legal certification",
		"patent cleared",
		"fto cleared",
		"de-facto standard",
		"immune system for open source",
		"universal trust score",
		"integrity score",
		"score > 90",
		"globally safe",
		"automatically blocked everywhere",
		"global truth",
		"crowd-sourced truth layer",
		"reviewed means safe",
		"community truth",
		"network truth",
		"auto-remediated",
		"auto-merged",
		"production autopatch",
		"point_9_pass",
	}
	for _, value := range values {
		normalized := strings.ToLower(strings.TrimSpace(value))
		if normalized == "" {
			continue
		}
		if _, ok := allowed[normalized]; ok {
			continue
		}
		for _, blocked := range disallowed {
			if strings.Contains(normalized, blocked) {
				return true
			}
		}
	}
	return false
}

func OSSTrustNetworkValCDependencySnapshotModel() OSSTrustNetworkValCDependencySnapshot {
	model := ComputeOSSTrustNetworkValBCore(OSSTrustNetworkValBCoreModel())
	return OSSTrustNetworkValCDependencySnapshot{
		ValBCurrentState:                model.CurrentState,
		ValBPoint9State:                 model.Point9State,
		ValBDependencyState:             model.DependencyState,
		ValBCandidateSignalIntakeState:  model.CandidateSignalIntakeState,
		ValBReviewWorkflowState:         model.ReviewWorkflowState,
		ValBSharedVEXTriageState:        model.SharedVEXTriageState,
		ValBSourceWeightingState:        model.SourceWeightingState,
		ValBLocalApplicabilityState:     model.LocalApplicabilityState,
		ValBPropagationExchangeState:    model.PropagationExchangeState,
		ValBSupersessionRevocationState: model.SupersessionRevocationState,
		ValBReviewerAuditabilityState:   model.ReviewerAuditabilityState,
		ValBNoOverclaimState:            model.NoOverclaimState,
		ValBProofSurfaceRefs:            append([]string{}, model.ProofSurfaceRefs...),
		ValBEvidenceRefs:                append([]string{}, model.EvidenceRefs...),
		ValBProjectionDisclaimer:        model.ProjectionDisclaimer,
	}
}

func OSSTrustNetworkValCTrustVisibilityModel() OSSTrustNetworkValCTrustVisibility {
	return OSSTrustNetworkValCTrustVisibility{
		CurrentState:             OSSTrustNetworkValCTrustVisibilityStateActive,
		VisibilityProfileID:      "ostn-valc-visibility-001",
		PackageOrProjectIdentity: "github.com/example/project",
		RegistryOrEcosystem:      "github",
		ReleaseOrVersionRef:      "refs/tags/v1.2.3",
		VisibilityState:          OSSTrustNetworkValCVisibilityVisible,
		ReviewedSignalState:      OSSTrustNetworkValBReviewWorkflowStateActive,
		LocalApplicabilityState:  OSSTrustNetworkValBLocalApplicabilityStateActive,
		SourceWeightingState:     OSSTrustNetworkValBSourceWeightingStateActive,
		EvidenceRefs:             []string{"evidence:ostn-valc-visibility-001"},
		FreshnessState:           IntelligenceCalibrationFreshnessFresh,
		Caveats:                  []string{"visibility remains bounded ecosystem projection"},
		ProjectionDisclaimer:     ossTrustNetworkValCProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValCPackageTrustStatusModel() OSSTrustNetworkValCPackageTrustStatus {
	return OSSTrustNetworkValCPackageTrustStatus{
		CurrentState:             OSSTrustNetworkValCPackageTrustStatusStateActive,
		StatusSummaryID:          "ostn-valc-package-status-001",
		PackageOrProjectIdentity: "github.com/example/project",
		ReleaseOrVersionRef:      "refs/tags/v1.2.3",
		StatusClass:              OSSTrustNetworkValCPackageStatusReviewedSignalAvailable,
		ReviewedSignalState:      OSSTrustNetworkValBReviewWorkflowStateActive,
		LocalApplicabilityState:  OSSTrustNetworkValBLocalApplicabilityStateActive,
		EvidenceRefs:             []string{"evidence:ostn-valc-package-status-001"},
		FreshnessState:           IntelligenceCalibrationFreshnessFresh,
		Caveats:                  []string{"package trust status remains bounded summary only"},
		DisplayedAsReviewed:      true,
		ProjectionDisclaimer:     ossTrustNetworkValCProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValCExportBoundaryModel() OSSTrustNetworkValCExportBoundary {
	return OSSTrustNetworkValCExportBoundary{
		CurrentState:         OSSTrustNetworkValCExportBoundaryStateActive,
		BoundaryID:           "ostn-valc-export-001",
		ExportClass:          OSSTrustNetworkValCExportClassEnterpriseCustomerView,
		EvidenceRefs:         []string{"evidence:ostn-valc-export-boundary-001"},
		Caveats:              []string{"export remains caveated and evidence-linked"},
		ProjectionDisclaimer: ossTrustNetworkValCProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValCRemediationSuggestionModel() OSSTrustNetworkValCRemediationSuggestion {
	return OSSTrustNetworkValCRemediationSuggestion{
		CurrentState:             OSSTrustNetworkValCRemediationSuggestionStateActive,
		SuggestionID:             "ostn-valc-remediation-001",
		SuggestionClass:          OSSTrustNetworkValCSuggestionClassVersionUpgrade,
		PackageOrProjectIdentity: "github.com/example/project",
		AffectedReleaseOrVersion: "refs/tags/v1.2.3",
		TargetReleaseOrVersion:   "refs/tags/v1.2.4",
		EvidenceRefs:             []string{"evidence:ostn-valc-remediation-suggestion-001"},
		ConfidenceClass:          OSSTrustNetworkConfidenceBounded,
		CompatibilityNote:        "bounded compatibility note with reviewer validation required",
		RiskNote:                 "moderate compatibility and rollout risk requires bounded review",
		LocalApplicabilityNote:   "apply only after enterprise-local applicability review",
		Rationale:                "reviewed signal suggests a bounded version upgrade path",
		Caveats:                  []string{"suggestions remain advisory only"},
		ProjectionDisclaimer:     ossTrustNetworkValCProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValCPRProposalModel() OSSTrustNetworkValCPRProposal {
	return OSSTrustNetworkValCPRProposal{
		CurrentState:           OSSTrustNetworkValCPRProposalStateActive,
		ProposalID:             "ostn-valc-pr-proposal-001",
		ProposalState:          OSSTrustNetworkValCProposalStateProposalReady,
		EvidenceRefs:           []string{"evidence:ostn-valc-pr-proposal-001"},
		CompatibilityNote:      "proposal remains reviewer-required and compatibility-bounded",
		ReviewerRequired:       true,
		NoAutomerge:            true,
		NoHiddenMutation:       true,
		LocalApplicabilityNote: "proposal remains enterprise-local and reviewer-gated",
		Caveats:                []string{"proposal descriptor only; no PR is created in Val C"},
		AdvisoryOnly:           true,
		ProjectionDisclaimer:   ossTrustNetworkValCProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValCLocalOverrideModel() OSSTrustNetworkValCLocalOverride {
	return OSSTrustNetworkValCLocalOverride{
		CurrentState:         OSSTrustNetworkValCLocalOverrideStateActive,
		OverrideID:           "ostn-valc-local-override-001",
		OverrideState:        OSSTrustNetworkValCOverrideStateNoOverride,
		EvidenceRefs:         []string{"evidence:ostn-valc-local-override-001"},
		Rationale:            "no local override recorded for this release context",
		Scope:                OSSTrustNetworkApplicabilityEnterpriseLocal,
		OwnerOrReviewerClass: "enterprise_security_reviewer",
		Caveats:              []string{"local override visibility remains local-only and evidence-linked"},
		LocalOnlyBoundary:    true,
		ProjectionDisclaimer: ossTrustNetworkValCProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValCRemediationSafetyModel() OSSTrustNetworkValCRemediationSafety {
	return OSSTrustNetworkValCRemediationSafety{
		CurrentState:         OSSTrustNetworkValCRemediationSafetyStateActive,
		SafetyID:             "ostn-valc-remediation-safety-001",
		SuggestionID:         "ostn-valc-remediation-001",
		CompatibilityNote:    "compatibility validated through bounded reviewer workflow",
		TestValidationNote:   "run unit, integration, and release validation before adoption",
		RollbackNote:         "rollback to refs/tags/v1.2.3 if bounded validation fails",
		RiskClass:            OSSTrustNetworkValCRiskClassMedium,
		ReviewerRequired:     true,
		EvidenceRefs:         []string{"evidence:ostn-valc-remediation-safety-001"},
		Caveats:              []string{"remediation safety remains advisory and reviewer-gated"},
		ProjectionDisclaimer: ossTrustNetworkValCProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValCEcosystemConsistencyModel() OSSTrustNetworkValCEcosystemConsistency {
	return OSSTrustNetworkValCEcosystemConsistency{
		CurrentState:                       OSSTrustNetworkValCEcosystemConsistencyStateActive,
		ConsistencyID:                      "ostn-valc-consistency-001",
		PackageOrProjectIdentity:           "github.com/example/project",
		ReleaseOrVersionRef:                "refs/tags/v1.2.3",
		SuggestionPackageOrProjectIdentity: "github.com/example/project",
		SuggestionReleaseOrVersion:         "refs/tags/v1.2.3",
		PackageStatusClass:                 OSSTrustNetworkValCPackageStatusReviewedSignalAvailable,
		ReviewState:                        OSSTrustNetworkValBReviewStateReviewed,
		ReviewerDecisionState:              OSSTrustNetworkValBReviewerDecisionStateAccepted,
		PropagationState:                   OSSTrustNetworkValBPropagationStateReviewedExchange,
		LocalApplicabilityState:            OSSTrustNetworkValBLocalApplicabilityStatusApplicable,
		DisplayedAsReviewed:                true,
		DisplayedAsApplicable:              true,
		ReviewedExchangePresentedActive:    true,
		FreshnessState:                     IntelligenceCalibrationFreshnessFresh,
		EvidenceRefs:                       []string{"evidence:ostn-valc-ecosystem-consistency-001"},
		Caveats:                            []string{"ecosystem visibility remains consistent with reviewed intelligence and local context"},
		ProjectionDisclaimer:               ossTrustNetworkValCProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValCNoOverclaimModel() OSSTrustNetworkValCNoOverclaim {
	return OSSTrustNetworkValCNoOverclaim{
		CurrentState:         OSSTrustNetworkValCNoOverclaimStateActive,
		DisciplineID:         "oss_no_overclaim_valc",
		Version:              "v0",
		ObservedClaims:       []string{"bounded ecosystem visibility", "advisory remediation suggestion", "no hidden mutation path"},
		ProjectionDisclaimer: ossTrustNetworkValCProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValCCoreModel() OSSTrustNetworkValCCore {
	return OSSTrustNetworkValCCore{
		CurrentState:          OSSTrustNetworkValCStateActive,
		Point9State:           OSSTrustNetworkPoint9StateNotComplete,
		Dependency:            OSSTrustNetworkValCDependencySnapshotModel(),
		TrustVisibility:       OSSTrustNetworkValCTrustVisibilityModel(),
		PackageTrustStatus:    OSSTrustNetworkValCPackageTrustStatusModel(),
		ExportBoundary:        OSSTrustNetworkValCExportBoundaryModel(),
		RemediationSuggestion: OSSTrustNetworkValCRemediationSuggestionModel(),
		PRProposal:            OSSTrustNetworkValCPRProposalModel(),
		LocalOverride:         OSSTrustNetworkValCLocalOverrideModel(),
		RemediationSafety:     OSSTrustNetworkValCRemediationSafetyModel(),
		EcosystemConsistency:  OSSTrustNetworkValCEcosystemConsistencyModel(),
		NoOverclaim:           OSSTrustNetworkValCNoOverclaimModel(),
		ProofSurfaceRefs:      OSSTrustNetworkValCProofSurfaceRefs(),
		EvidenceRefs:          OSSTrustNetworkValCProofEvidenceRefs(),
		WhyPoint9NotComplete: []string{
			"Val C provides bounded remediation and ecosystem visibility only and cannot complete Točka 9.",
			"Final gates, integrated closure, and any final pass semantics remain reserved for later OSTN waves.",
			"Visibility, remediation suggestions, proposal descriptors, and local overrides remain advisory and cannot create hidden mutation, approval authority, or canonical truth.",
		},
		ProjectionDisclaimer: ossTrustNetworkValCProjectionDisclaimer(),
	}
}

func EvaluateOSSTrustNetworkValCDependencyState(model OSSTrustNetworkValCDependencySnapshot) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ValBCurrentState,
		model.ValBPoint9State,
		model.ValBDependencyState,
		model.ValBCandidateSignalIntakeState,
		model.ValBReviewWorkflowState,
		model.ValBSharedVEXTriageState,
		model.ValBSourceWeightingState,
		model.ValBLocalApplicabilityState,
		model.ValBPropagationExchangeState,
		model.ValBSupersessionRevocationState,
		model.ValBReviewerAuditabilityState,
		model.ValBNoOverclaimState,
		model.ValBProjectionDisclaimer,
	) || len(model.ValBProofSurfaceRefs) == 0 || len(model.ValBEvidenceRefs) == 0 {
		return OSSTrustNetworkValCDependencyStateIncomplete
	}
	if !ossTrustNetworkValBHasProjectionDisclaimer(model.ValBProjectionDisclaimer) {
		return OSSTrustNetworkValCDependencyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.ValBProofSurfaceRefs, OSSTrustNetworkValBProofSurfaceRefs()...) ||
		!containsExactTrimmedStringSet(model.ValBEvidenceRefs, OSSTrustNetworkValBProofEvidenceRefs()...) ||
		!OSSTrustNetworkValBProofEvidenceQualityValid(ossTrustNetworkValBEvidence(), model.ValBEvidenceRefs) {
		return OSSTrustNetworkValCDependencyStateBlocked
	}
	if strings.TrimSpace(model.ValBCurrentState) != OSSTrustNetworkValBStateActive ||
		strings.TrimSpace(model.ValBPoint9State) != OSSTrustNetworkPoint9StateNotComplete ||
		strings.TrimSpace(model.ValBDependencyState) != OSSTrustNetworkValBDependencyStateActive ||
		strings.TrimSpace(model.ValBCandidateSignalIntakeState) != OSSTrustNetworkValBCandidateSignalIntakeStateActive ||
		strings.TrimSpace(model.ValBReviewWorkflowState) != OSSTrustNetworkValBReviewWorkflowStateActive ||
		strings.TrimSpace(model.ValBSharedVEXTriageState) != OSSTrustNetworkValBSharedVEXTriageStateActive ||
		strings.TrimSpace(model.ValBSourceWeightingState) != OSSTrustNetworkValBSourceWeightingStateActive ||
		strings.TrimSpace(model.ValBLocalApplicabilityState) != OSSTrustNetworkValBLocalApplicabilityStateActive ||
		strings.TrimSpace(model.ValBPropagationExchangeState) != OSSTrustNetworkValBPropagationExchangeStateActive ||
		strings.TrimSpace(model.ValBSupersessionRevocationState) != OSSTrustNetworkValBSupersessionRevocationStateActive ||
		strings.TrimSpace(model.ValBReviewerAuditabilityState) != OSSTrustNetworkValBReviewerAuditabilityStateActive ||
		strings.TrimSpace(model.ValBNoOverclaimState) != OSSTrustNetworkValBNoOverclaimStateActive {
		return OSSTrustNetworkValCDependencyStateBlocked
	}
	return OSSTrustNetworkValCDependencyStateActive
}

func EvaluateOSSTrustNetworkValCTrustVisibilityState(model OSSTrustNetworkValCTrustVisibility) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.VisibilityProfileID,
		model.PackageOrProjectIdentity,
		model.RegistryOrEcosystem,
		model.ReleaseOrVersionRef,
		model.VisibilityState,
		model.ReviewedSignalState,
		model.LocalApplicabilityState,
		model.SourceWeightingState,
		model.FreshnessState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValCTrustVisibilityStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValCTrustVisibilityStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valc-visibility-001") ||
		!containsTrimmedString(ossTrustNetworkValCVisibilityStates(), model.VisibilityState) ||
		model.PackageSafetyClaim ||
		model.CertifiedPackageClaim ||
		model.ProductionApprovalClaim ||
		model.DeploymentApprovalClaim ||
		model.GlobalTruthClaim ||
		model.EnterpriseOverrideClaim {
		return OSSTrustNetworkValCTrustVisibilityStateBlocked
	}
	switch strings.TrimSpace(model.VisibilityState) {
	case OSSTrustNetworkValCVisibilityVisible:
		if ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) &&
			strings.TrimSpace(model.ReviewedSignalState) == OSSTrustNetworkValBReviewWorkflowStateActive &&
			strings.TrimSpace(model.LocalApplicabilityState) == OSSTrustNetworkValBLocalApplicabilityStateActive &&
			strings.TrimSpace(model.SourceWeightingState) == OSSTrustNetworkValBSourceWeightingStateActive {
			return OSSTrustNetworkValCTrustVisibilityStateActive
		}
		return OSSTrustNetworkValCTrustVisibilityStateBlocked
	case OSSTrustNetworkValCVisibilityLimited:
		return OSSTrustNetworkValCTrustVisibilityStatePartial
	case OSSTrustNetworkValCVisibilityHidden,
		OSSTrustNetworkValCVisibilityUnsupported,
		OSSTrustNetworkValCVisibilityStale,
		OSSTrustNetworkValCVisibilityUnknown:
		return OSSTrustNetworkValCTrustVisibilityStateBlocked
	default:
		return OSSTrustNetworkValCTrustVisibilityStateBlocked
	}
}

func EvaluateOSSTrustNetworkValCPackageTrustStatusState(model OSSTrustNetworkValCPackageTrustStatus) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.StatusSummaryID,
		model.PackageOrProjectIdentity,
		model.ReleaseOrVersionRef,
		model.StatusClass,
		model.ReviewedSignalState,
		model.LocalApplicabilityState,
		model.FreshnessState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValCPackageTrustStatusStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValCPackageTrustStatusStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valc-package-status-001") ||
		!containsTrimmedString(ossTrustNetworkValCPackageStatusClasses(), model.StatusClass) ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) ||
		model.UniversalTrustScoreClaim ||
		model.IntegrityScoreClaim ||
		model.BadgeScoreClaim ||
		model.GenericSafetyClaim {
		return OSSTrustNetworkValCPackageTrustStatusStateBlocked
	}
	switch strings.TrimSpace(model.StatusClass) {
	case OSSTrustNetworkValCPackageStatusReviewedSignalAvailable:
		if strings.TrimSpace(model.ReviewedSignalState) == OSSTrustNetworkValBReviewWorkflowStateActive &&
			strings.TrimSpace(model.LocalApplicabilityState) == OSSTrustNetworkValBLocalApplicabilityStateActive {
			return OSSTrustNetworkValCPackageTrustStatusStateActive
		}
		return OSSTrustNetworkValCPackageTrustStatusStateBlocked
	case OSSTrustNetworkValCPackageStatusCandidateSignalAvailable:
		if model.DisplayedAsReviewed {
			return OSSTrustNetworkValCPackageTrustStatusStateBlocked
		}
		return OSSTrustNetworkValCPackageTrustStatusStatePartial
	case OSSTrustNetworkValCPackageStatusLocalReviewNeeded:
		if model.DisplayedAsReviewed {
			return OSSTrustNetworkValCPackageTrustStatusStateBlocked
		}
		return OSSTrustNetworkValCPackageTrustStatusStatePartial
	case OSSTrustNetworkValCPackageStatusSupersededSignal:
		if strings.TrimSpace(model.ReplacementRef) == "" {
			return OSSTrustNetworkValCPackageTrustStatusStateBlocked
		}
		return OSSTrustNetworkValCPackageTrustStatusStatePartial
	case OSSTrustNetworkValCPackageStatusRevokedSignal,
		OSSTrustNetworkValCPackageStatusUnsupportedSignal,
		OSSTrustNetworkValCPackageStatusUnknownSignal:
		return OSSTrustNetworkValCPackageTrustStatusStateBlocked
	default:
		return OSSTrustNetworkValCPackageTrustStatusStateBlocked
	}
}

func EvaluateOSSTrustNetworkValCExportBoundaryState(model OSSTrustNetworkValCExportBoundary) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.BoundaryID,
		model.ExportClass,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValCExportBoundaryStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValCExportBoundaryStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valc-export-boundary-001") ||
		!containsTrimmedString(ossTrustNetworkValCExportClasses(), model.ExportClass) ||
		model.CandidatePromotedToReviewed ||
		model.RejectedPromotedToActive ||
		model.RevokedPromotedToActive ||
		model.UnknownPromotedToActive ||
		model.GlobalBlocklistClaim {
		return OSSTrustNetworkValCExportBoundaryStateBlocked
	}
	switch strings.TrimSpace(model.ExportClass) {
	case OSSTrustNetworkValCExportClassInternalOperatorView, OSSTrustNetworkValCExportClassEnterpriseCustomerView:
		if model.CertificationClaim || model.ApprovalClaim {
			return OSSTrustNetworkValCExportBoundaryStateBlocked
		}
		return OSSTrustNetworkValCExportBoundaryStateActive
	case OSSTrustNetworkValCExportClassMaintainerView, OSSTrustNetworkValCExportClassPartnerView:
		if model.RedactionStripsCaveats || model.RedactionStripsEvidenceRefs || model.CertificationClaim || model.ApprovalClaim {
			return OSSTrustNetworkValCExportBoundaryStateBlocked
		}
		return OSSTrustNetworkValCExportBoundaryStateActive
	case OSSTrustNetworkValCExportClassPublicSummaryView:
		if model.CanonicalInternalExposure || model.CertificationClaim || model.ApprovalClaim {
			return OSSTrustNetworkValCExportBoundaryStateBlocked
		}
		return OSSTrustNetworkValCExportBoundaryStateActive
	case OSSTrustNetworkValCExportClassUnsupportedView:
		return OSSTrustNetworkValCExportBoundaryStatePartial
	default:
		return OSSTrustNetworkValCExportBoundaryStateBlocked
	}
}

func EvaluateOSSTrustNetworkValCRemediationSuggestionState(model OSSTrustNetworkValCRemediationSuggestion) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.SuggestionID,
		model.SuggestionClass,
		model.PackageOrProjectIdentity,
		model.AffectedReleaseOrVersion,
		model.TargetReleaseOrVersion,
		model.ConfidenceClass,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValCRemediationSuggestionStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValCRemediationSuggestionStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valc-remediation-suggestion-001") ||
		!containsTrimmedString(ossTrustNetworkValCSuggestionClasses(), model.SuggestionClass) ||
		!containsTrimmedString(ossTrustNetworkVal0ConfidenceClasses(), model.ConfidenceClass) ||
		strings.TrimSpace(model.CompatibilityNote) == "" ||
		strings.TrimSpace(model.RiskNote) == "" ||
		strings.TrimSpace(model.LocalApplicabilityNote) == "" ||
		model.DependencyMutationAttempt ||
		model.PolicyOverrideAttempt {
		return OSSTrustNetworkValCRemediationSuggestionStateBlocked
	}
	switch strings.TrimSpace(model.SuggestionClass) {
	case OSSTrustNetworkValCSuggestionClassNoAction:
		if strings.TrimSpace(model.Rationale) == "" {
			return OSSTrustNetworkValCRemediationSuggestionStateBlocked
		}
		return OSSTrustNetworkValCRemediationSuggestionStateActive
	case OSSTrustNetworkValCSuggestionClassUnsupported:
		return OSSTrustNetworkValCRemediationSuggestionStatePartial
	case OSSTrustNetworkValCSuggestionClassVersionUpgrade,
		OSSTrustNetworkValCSuggestionClassPinOrHold,
		OSSTrustNetworkValCSuggestionClassReplaceDependency,
		OSSTrustNetworkValCSuggestionClassMaintainerContact,
		OSSTrustNetworkValCSuggestionClassReviewRequired:
		return OSSTrustNetworkValCRemediationSuggestionStateActive
	default:
		return OSSTrustNetworkValCRemediationSuggestionStateBlocked
	}
}

func EvaluateOSSTrustNetworkValCPRProposalState(model OSSTrustNetworkValCPRProposal) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ProposalID,
		model.ProposalState,
		model.CompatibilityNote,
		model.LocalApplicabilityNote,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValCPRProposalStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValCPRProposalStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valc-pr-proposal-001") ||
		!containsTrimmedString(ossTrustNetworkValCProposalStates(), model.ProposalState) ||
		!model.AdvisoryOnly ||
		model.BranchWrite ||
		model.NetworkAction ||
		model.DependencyMutation ||
		model.PRCreation ||
		model.AutoMerge {
		return OSSTrustNetworkValCPRProposalStateBlocked
	}
	switch strings.TrimSpace(model.ProposalState) {
	case OSSTrustNetworkValCProposalStateProposalReady:
		if model.ReviewerRequired && model.NoAutomerge && model.NoHiddenMutation {
			return OSSTrustNetworkValCPRProposalStateActive
		}
		return OSSTrustNetworkValCPRProposalStateBlocked
	case OSSTrustNetworkValCProposalStateNeedsReview:
		if model.NoAutomerge && model.NoHiddenMutation {
			return OSSTrustNetworkValCPRProposalStatePartial
		}
		return OSSTrustNetworkValCPRProposalStateBlocked
	case OSSTrustNetworkValCProposalStateUnsupported,
		OSSTrustNetworkValCProposalStateBlocked,
		OSSTrustNetworkValCProposalStateUnknown:
		return OSSTrustNetworkValCPRProposalStateBlocked
	default:
		return OSSTrustNetworkValCPRProposalStateBlocked
	}
}

func EvaluateOSSTrustNetworkValCLocalOverrideState(model OSSTrustNetworkValCLocalOverride) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.OverrideID,
		model.OverrideState,
		model.Scope,
		model.OwnerOrReviewerClass,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValCLocalOverrideStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValCLocalOverrideStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valc-local-override-001") ||
		!containsTrimmedString(ossTrustNetworkValCLocalOverrideStates(), model.OverrideState) ||
		!containsTrimmedString(ossTrustNetworkVal0ApplicabilityScopes(), model.Scope) ||
		model.RewriteCanonicalEvidence ||
		model.SilentlySuppressNetworkIntelligence ||
		model.SharedSignalOverridesLocalDecision {
		return OSSTrustNetworkValCLocalOverrideStateBlocked
	}
	switch strings.TrimSpace(model.OverrideState) {
	case OSSTrustNetworkValCOverrideStateNoOverride:
		return OSSTrustNetworkValCLocalOverrideStateActive
	case OSSTrustNetworkValCOverrideStateOverridePresent:
		if strings.TrimSpace(model.Rationale) != "" && model.LocalOnlyBoundary {
			return OSSTrustNetworkValCLocalOverrideStateActive
		}
		return OSSTrustNetworkValCLocalOverrideStateBlocked
	case OSSTrustNetworkValCOverrideStateOverrideRequiresReview:
		return OSSTrustNetworkValCLocalOverrideStatePartial
	case OSSTrustNetworkValCOverrideStateOverrideRejected,
		OSSTrustNetworkValCOverrideStateUnsupported,
		OSSTrustNetworkValCOverrideStateUnknown:
		return OSSTrustNetworkValCLocalOverrideStateBlocked
	default:
		return OSSTrustNetworkValCLocalOverrideStateBlocked
	}
}

func EvaluateOSSTrustNetworkValCRemediationSafetyState(model OSSTrustNetworkValCRemediationSafety) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.SafetyID,
		model.SuggestionID,
		model.RiskClass,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValCRemediationSafetyStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValCRemediationSafetyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valc-remediation-safety-001") ||
		!containsTrimmedString(ossTrustNetworkValCRiskClasses(), model.RiskClass) ||
		strings.TrimSpace(model.CompatibilityNote) == "" ||
		strings.TrimSpace(model.TestValidationNote) == "" ||
		strings.TrimSpace(model.RollbackNote) == "" ||
		model.ProductionApprovalClaim ||
		model.DeploymentApprovalClaim ||
		model.HiddenMutationPath ||
		!model.ReviewerRequired {
		return OSSTrustNetworkValCRemediationSafetyStateBlocked
	}
	if strings.TrimSpace(model.TestValidationNote) == "" || strings.TrimSpace(model.RollbackNote) == "" {
		return OSSTrustNetworkValCRemediationSafetyStateBlocked
	}
	return OSSTrustNetworkValCRemediationSafetyStateActive
}

func EvaluateOSSTrustNetworkValCEcosystemConsistencyState(model OSSTrustNetworkValCEcosystemConsistency) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ConsistencyID,
		model.PackageOrProjectIdentity,
		model.ReleaseOrVersionRef,
		model.SuggestionPackageOrProjectIdentity,
		model.SuggestionReleaseOrVersion,
		model.PackageStatusClass,
		model.ReviewState,
		model.ReviewerDecisionState,
		model.PropagationState,
		model.LocalApplicabilityState,
		model.FreshnessState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValCEcosystemConsistencyStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValCEcosystemConsistencyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valc-ecosystem-consistency-001") ||
		!containsTrimmedString(ossTrustNetworkValCPackageStatusClasses(), model.PackageStatusClass) ||
		!containsTrimmedString(ossTrustNetworkValBReviewStates(), model.ReviewState) ||
		!containsTrimmedString(ossTrustNetworkValBReviewerDecisionStates(), model.ReviewerDecisionState) ||
		!containsTrimmedString(ossTrustNetworkValBPropagationStates(), model.PropagationState) ||
		!containsTrimmedString(ossTrustNetworkValBLocalApplicabilityStates(), model.LocalApplicabilityState) ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) {
		return OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	if strings.TrimSpace(model.SuggestionPackageOrProjectIdentity) != strings.TrimSpace(model.PackageOrProjectIdentity) ||
		strings.TrimSpace(model.SuggestionReleaseOrVersion) != strings.TrimSpace(model.ReleaseOrVersionRef) {
		return OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	if strings.TrimSpace(model.PackageStatusClass) == OSSTrustNetworkValCPackageStatusReviewedSignalAvailable &&
		(strings.TrimSpace(model.ReviewState) != OSSTrustNetworkValBReviewStateReviewed ||
			strings.TrimSpace(model.ReviewerDecisionState) != OSSTrustNetworkValBReviewerDecisionStateAccepted) {
		return OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	if strings.TrimSpace(model.PackageStatusClass) == OSSTrustNetworkValCPackageStatusCandidateSignalAvailable &&
		model.DisplayedAsReviewed {
		return OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	if strings.TrimSpace(model.PackageStatusClass) == OSSTrustNetworkValCPackageStatusRevokedSignal &&
		model.ReviewedExchangePresentedActive {
		return OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	if strings.TrimSpace(model.PackageStatusClass) == OSSTrustNetworkValCPackageStatusSupersededSignal &&
		model.ReviewedExchangePresentedActive {
		return OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	if (strings.TrimSpace(model.PackageStatusClass) == OSSTrustNetworkValCPackageStatusUnsupportedSignal ||
		strings.TrimSpace(model.PackageStatusClass) == OSSTrustNetworkValCPackageStatusUnknownSignal) &&
		model.ReviewedExchangePresentedActive {
		return OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	if model.ReviewedExchangePresentedActive &&
		strings.TrimSpace(model.PropagationState) != OSSTrustNetworkValBPropagationStateReviewedExchange {
		return OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	if (strings.TrimSpace(model.PropagationState) == OSSTrustNetworkValBPropagationStateRevoked ||
		strings.TrimSpace(model.PropagationState) == OSSTrustNetworkValBPropagationStateSuperseded) &&
		model.ReviewedExchangePresentedActive {
		return OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	if strings.TrimSpace(model.LocalApplicabilityState) == OSSTrustNetworkValBLocalApplicabilityStatusUnknown &&
		model.DisplayedAsApplicable {
		return OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	return OSSTrustNetworkValCEcosystemConsistencyStateActive
}

func ossTrustNetworkValCCrossSurfaceConsistencyMismatch(model OSSTrustNetworkValCCore) bool {
	packageIdentity := strings.TrimSpace(model.EcosystemConsistency.PackageOrProjectIdentity)
	releaseRef := strings.TrimSpace(model.EcosystemConsistency.ReleaseOrVersionRef)

	if strings.TrimSpace(model.TrustVisibility.PackageOrProjectIdentity) != packageIdentity ||
		strings.TrimSpace(model.TrustVisibility.ReleaseOrVersionRef) != releaseRef {
		return true
	}
	if strings.TrimSpace(model.PackageTrustStatus.PackageOrProjectIdentity) != packageIdentity ||
		strings.TrimSpace(model.PackageTrustStatus.ReleaseOrVersionRef) != releaseRef {
		return true
	}
	if strings.TrimSpace(model.RemediationSuggestion.PackageOrProjectIdentity) != packageIdentity ||
		strings.TrimSpace(model.RemediationSuggestion.AffectedReleaseOrVersion) != releaseRef {
		return true
	}
	if strings.TrimSpace(model.PackageTrustStatus.StatusClass) != strings.TrimSpace(model.EcosystemConsistency.PackageStatusClass) {
		return true
	}
	if model.PackageTrustStatus.DisplayedAsReviewed != model.EcosystemConsistency.DisplayedAsReviewed {
		return true
	}
	if model.EcosystemConsistency.ReviewedExchangePresentedActive {
		if strings.TrimSpace(model.PackageTrustStatus.StatusClass) != OSSTrustNetworkValCPackageStatusReviewedSignalAvailable ||
			strings.TrimSpace(model.PackageTrustStatusState) != OSSTrustNetworkValCPackageTrustStatusStateActive ||
			strings.TrimSpace(model.Dependency.ValBReviewWorkflowState) != OSSTrustNetworkValBReviewWorkflowStateActive ||
			strings.TrimSpace(model.Dependency.ValBPropagationExchangeState) != OSSTrustNetworkValBPropagationExchangeStateActive ||
			strings.TrimSpace(model.Dependency.ValBLocalApplicabilityState) != OSSTrustNetworkValBLocalApplicabilityStateActive ||
			strings.TrimSpace(model.PackageTrustStatus.ReviewedSignalState) != OSSTrustNetworkValBReviewWorkflowStateActive ||
			strings.TrimSpace(model.PackageTrustStatus.LocalApplicabilityState) != OSSTrustNetworkValBLocalApplicabilityStateActive ||
			strings.TrimSpace(model.TrustVisibility.ReviewedSignalState) != OSSTrustNetworkValBReviewWorkflowStateActive ||
			strings.TrimSpace(model.TrustVisibility.LocalApplicabilityState) != OSSTrustNetworkValBLocalApplicabilityStateActive {
			return true
		}
	}
	if model.EcosystemConsistency.DisplayedAsApplicable {
		if strings.TrimSpace(model.Dependency.ValBLocalApplicabilityState) != OSSTrustNetworkValBLocalApplicabilityStateActive ||
			strings.TrimSpace(model.PackageTrustStatus.LocalApplicabilityState) != OSSTrustNetworkValBLocalApplicabilityStateActive ||
			strings.TrimSpace(model.TrustVisibility.LocalApplicabilityState) != OSSTrustNetworkValBLocalApplicabilityStateActive {
			return true
		}
		switch strings.TrimSpace(model.LocalOverride.OverrideState) {
		case OSSTrustNetworkValCOverrideStateOverrideRequiresReview,
			OSSTrustNetworkValCOverrideStateUnsupported,
			OSSTrustNetworkValCOverrideStateUnknown:
			return true
		}
	}
	return false
}

func EvaluateOSSTrustNetworkValCNoOverclaimState(model OSSTrustNetworkValCNoOverclaim) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return OSSTrustNetworkValCNoOverclaimStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValCNoOverclaimStateUnknown
	}
	if model.GlobalTruthClaim ||
		model.ReviewedMeansSafeClaim ||
		model.CommunityTruthClaim ||
		model.NetworkTruthClaim ||
		model.CertifiedPackageClaim ||
		model.OfficiallySafePackageClaim ||
		model.RegulatorApproved ||
		model.AuditPassed ||
		model.ComplianceGuaranteed ||
		model.ProductionApproved ||
		model.DeploymentApproved ||
		model.LegalCertification ||
		model.PatentCleared ||
		model.FTOPCleared ||
		model.AutoRemediated ||
		model.AutoMerged ||
		model.ProductionAutopatch ||
		ossTrustNetworkValCContainsForbiddenClaim(model.ObservedClaims...) {
		return OSSTrustNetworkValCNoOverclaimStateBlocked
	}
	return OSSTrustNetworkValCNoOverclaimStateActive
}

func EvaluateOSSTrustNetworkValCState(model OSSTrustNetworkValCCore) string {
	if strings.TrimSpace(model.Point9State) == "" || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return OSSTrustNetworkValCStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValCStateUnknown
	}
	if strings.TrimSpace(model.Point9State) != OSSTrustNetworkPoint9StateNotComplete ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValCProofSurfaceRefs()...) ||
		!OSSTrustNetworkValCProofEvidenceQualityValid(ossTrustNetworkValCEvidence(), model.EvidenceRefs) {
		return OSSTrustNetworkValCStateBlocked
	}
	states := []string{
		model.DependencyState,
		model.TrustVisibilityState,
		model.PackageTrustStatusState,
		model.ExportBoundaryState,
		model.RemediationSuggestionState,
		model.PRProposalState,
		model.LocalOverrideState,
		model.RemediationSafetyState,
		model.EcosystemConsistencyState,
		model.NoOverclaimState,
	}
	allActive := true
	for _, state := range states {
		if strings.TrimSpace(state) == "" {
			return OSSTrustNetworkValCStateIncomplete
		}
		if !strings.HasSuffix(strings.TrimSpace(state), "_active") {
			allActive = false
		}
	}
	if allActive {
		return OSSTrustNetworkValCStateActive
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_blocked") {
			return OSSTrustNetworkValCStateBlocked
		}
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_incomplete") {
			return OSSTrustNetworkValCStateIncomplete
		}
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_unknown") {
			return OSSTrustNetworkValCStateUnknown
		}
	}
	return OSSTrustNetworkValCStatePartial
}

func EvaluateOSSTrustNetworkValCPointsState(currentState string) string {
	_ = currentState
	return OSSTrustNetworkPoint9StateNotComplete
}

func EvaluateOSSTrustNetworkValCProofsState(model OSSTrustNetworkValCCore, limitations []string) string {
	baseState := strings.TrimSpace(model.CurrentState)
	if baseState == "" {
		baseState = OSSTrustNetworkValCStateUnknown
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValCProofSurfaceRefs()...) ||
		!OSSTrustNetworkValCProofEvidenceQualityValid(ossTrustNetworkValCEvidence(), model.EvidenceRefs) ||
		len(limitations) == 0 ||
		strings.TrimSpace(model.Point9State) != OSSTrustNetworkPoint9StateNotComplete {
		if baseState == OSSTrustNetworkValCStateActive {
			return OSSTrustNetworkValCStatePartial
		}
		return baseState
	}
	return baseState
}

func ossTrustNetworkValCBlockingReasons(model OSSTrustNetworkValCCore) []string {
	reasons := []string{}
	if model.DependencyState != OSSTrustNetworkValCDependencyStateActive {
		reasons = append(reasons, "OSTN Val B dependency is not exact, active, and evidence-safe.")
	}
	if model.TrustVisibilityState != OSSTrustNetworkValCTrustVisibilityStateActive {
		reasons = append(reasons, "OSS trust visibility is not exact, fresh, reviewed, and locally bounded.")
	}
	if model.PackageTrustStatusState != OSSTrustNetworkValCPackageTrustStatusStateActive {
		reasons = append(reasons, "Package trust status summary is not reviewed, fresh, and local-context-bounded.")
	}
	if model.ExportBoundaryState != OSSTrustNetworkValCExportBoundaryStateActive {
		reasons = append(reasons, "Ecosystem export boundary is not exact, caveated, and redaction-safe.")
	}
	if model.RemediationSuggestionState != OSSTrustNetworkValCRemediationSuggestionStateActive {
		reasons = append(reasons, "Remediation suggestion descriptor is not evidence-linked, scoped, and advisory-only.")
	}
	if model.PRProposalState != OSSTrustNetworkValCPRProposalStateActive {
		reasons = append(reasons, "PR proposal descriptor is not reviewer-required, no-automerge, and free of mutation paths.")
	}
	if model.LocalOverrideState != OSSTrustNetworkValCLocalOverrideStateActive {
		reasons = append(reasons, "Local override visibility is not evidence-linked, local-only, and suppression-safe.")
	}
	if model.RemediationSafetyState != OSSTrustNetworkValCRemediationSafetyStateActive {
		reasons = append(reasons, "Remediation safety is not reviewer-gated, rollback-ready, and validation-backed.")
	}
	if model.EcosystemConsistencyState != OSSTrustNetworkValCEcosystemConsistencyStateActive {
		reasons = append(reasons, "Ecosystem visibility, package status, remediation, and local applicability are not internally consistent.")
	}
	if model.NoOverclaimState != OSSTrustNetworkValCNoOverclaimStateActive {
		reasons = append(reasons, "Val C no-overclaim and no-hidden-mutation guard is not active.")
	}
	return developerEcosystemValECollectText(reasons)
}

func ComputeOSSTrustNetworkValCCore(model OSSTrustNetworkValCCore) OSSTrustNetworkValCCore {
	model.DependencyState = EvaluateOSSTrustNetworkValCDependencyState(model.Dependency)
	model.TrustVisibilityState = EvaluateOSSTrustNetworkValCTrustVisibilityState(model.TrustVisibility)
	model.PackageTrustStatusState = EvaluateOSSTrustNetworkValCPackageTrustStatusState(model.PackageTrustStatus)
	model.ExportBoundaryState = EvaluateOSSTrustNetworkValCExportBoundaryState(model.ExportBoundary)
	model.RemediationSuggestionState = EvaluateOSSTrustNetworkValCRemediationSuggestionState(model.RemediationSuggestion)
	model.PRProposalState = EvaluateOSSTrustNetworkValCPRProposalState(model.PRProposal)
	model.LocalOverrideState = EvaluateOSSTrustNetworkValCLocalOverrideState(model.LocalOverride)
	model.RemediationSafetyState = EvaluateOSSTrustNetworkValCRemediationSafetyState(model.RemediationSafety)
	model.EcosystemConsistencyState = EvaluateOSSTrustNetworkValCEcosystemConsistencyState(model.EcosystemConsistency)
	if model.EcosystemConsistencyState != OSSTrustNetworkValCEcosystemConsistencyStateIncomplete &&
		model.EcosystemConsistencyState != OSSTrustNetworkValCEcosystemConsistencyStateUnknown &&
		!strings.HasSuffix(strings.TrimSpace(model.TrustVisibilityState), "_unknown") &&
		!strings.HasSuffix(strings.TrimSpace(model.TrustVisibilityState), "_incomplete") &&
		!strings.HasSuffix(strings.TrimSpace(model.PackageTrustStatusState), "_unknown") &&
		!strings.HasSuffix(strings.TrimSpace(model.PackageTrustStatusState), "_incomplete") &&
		!strings.HasSuffix(strings.TrimSpace(model.RemediationSuggestionState), "_unknown") &&
		!strings.HasSuffix(strings.TrimSpace(model.RemediationSuggestionState), "_incomplete") &&
		!strings.HasSuffix(strings.TrimSpace(model.LocalOverrideState), "_unknown") &&
		!strings.HasSuffix(strings.TrimSpace(model.LocalOverrideState), "_incomplete") &&
		ossTrustNetworkValCCrossSurfaceConsistencyMismatch(model) {
		model.EcosystemConsistencyState = OSSTrustNetworkValCEcosystemConsistencyStateBlocked
	}
	model.NoOverclaimState = EvaluateOSSTrustNetworkValCNoOverclaimState(model.NoOverclaim)
	model.Point9State = EvaluateOSSTrustNetworkValCPointsState(model.CurrentState)
	model.CurrentState = EvaluateOSSTrustNetworkValCState(model)
	model.Point9State = EvaluateOSSTrustNetworkValCPointsState(model.CurrentState)
	model.BlockingReasons = ossTrustNetworkValCBlockingReasons(model)

	model.TrustVisibility.CurrentState = model.TrustVisibilityState
	model.PackageTrustStatus.CurrentState = model.PackageTrustStatusState
	model.ExportBoundary.CurrentState = model.ExportBoundaryState
	model.RemediationSuggestion.CurrentState = model.RemediationSuggestionState
	model.PRProposal.CurrentState = model.PRProposalState
	model.LocalOverride.CurrentState = model.LocalOverrideState
	model.RemediationSafety.CurrentState = model.RemediationSafetyState
	model.EcosystemConsistency.CurrentState = model.EcosystemConsistencyState
	model.NoOverclaim.CurrentState = model.NoOverclaimState

	return model
}

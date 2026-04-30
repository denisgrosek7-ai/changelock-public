package operability

import "strings"

const (
	OSSTrustNetworkValBStateActive     = "oss_trust_network_valb_active"
	OSSTrustNetworkValBStatePartial    = "oss_trust_network_valb_partial"
	OSSTrustNetworkValBStateIncomplete = "oss_trust_network_valb_incomplete"
	OSSTrustNetworkValBStateBlocked    = "oss_trust_network_valb_blocked"
	OSSTrustNetworkValBStateUnknown    = "oss_trust_network_valb_unknown"

	OSSTrustNetworkValBDependencyStateActive     = "oss_trust_network_valb_dependency_active"
	OSSTrustNetworkValBDependencyStatePartial    = "oss_trust_network_valb_dependency_partial"
	OSSTrustNetworkValBDependencyStateIncomplete = "oss_trust_network_valb_dependency_incomplete"
	OSSTrustNetworkValBDependencyStateBlocked    = "oss_trust_network_valb_dependency_blocked"
	OSSTrustNetworkValBDependencyStateUnknown    = "oss_trust_network_valb_dependency_unknown"

	OSSTrustNetworkValBCandidateSignalIntakeStateActive     = "oss_trust_network_valb_candidate_signal_intake_active"
	OSSTrustNetworkValBCandidateSignalIntakeStatePartial    = "oss_trust_network_valb_candidate_signal_intake_partial"
	OSSTrustNetworkValBCandidateSignalIntakeStateIncomplete = "oss_trust_network_valb_candidate_signal_intake_incomplete"
	OSSTrustNetworkValBCandidateSignalIntakeStateBlocked    = "oss_trust_network_valb_candidate_signal_intake_blocked"
	OSSTrustNetworkValBCandidateSignalIntakeStateUnknown    = "oss_trust_network_valb_candidate_signal_intake_unknown"

	OSSTrustNetworkValBReviewWorkflowStateActive     = "oss_trust_network_valb_review_workflow_active"
	OSSTrustNetworkValBReviewWorkflowStatePartial    = "oss_trust_network_valb_review_workflow_partial"
	OSSTrustNetworkValBReviewWorkflowStateIncomplete = "oss_trust_network_valb_review_workflow_incomplete"
	OSSTrustNetworkValBReviewWorkflowStateBlocked    = "oss_trust_network_valb_review_workflow_blocked"
	OSSTrustNetworkValBReviewWorkflowStateUnknown    = "oss_trust_network_valb_review_workflow_unknown"

	OSSTrustNetworkValBSharedVEXTriageStateActive     = "oss_trust_network_valb_shared_vex_triage_active"
	OSSTrustNetworkValBSharedVEXTriageStatePartial    = "oss_trust_network_valb_shared_vex_triage_partial"
	OSSTrustNetworkValBSharedVEXTriageStateIncomplete = "oss_trust_network_valb_shared_vex_triage_incomplete"
	OSSTrustNetworkValBSharedVEXTriageStateBlocked    = "oss_trust_network_valb_shared_vex_triage_blocked"
	OSSTrustNetworkValBSharedVEXTriageStateUnknown    = "oss_trust_network_valb_shared_vex_triage_unknown"

	OSSTrustNetworkValBSourceWeightingStateActive     = "oss_trust_network_valb_source_weighting_active"
	OSSTrustNetworkValBSourceWeightingStatePartial    = "oss_trust_network_valb_source_weighting_partial"
	OSSTrustNetworkValBSourceWeightingStateIncomplete = "oss_trust_network_valb_source_weighting_incomplete"
	OSSTrustNetworkValBSourceWeightingStateBlocked    = "oss_trust_network_valb_source_weighting_blocked"
	OSSTrustNetworkValBSourceWeightingStateUnknown    = "oss_trust_network_valb_source_weighting_unknown"

	OSSTrustNetworkValBLocalApplicabilityStateActive     = "oss_trust_network_valb_local_applicability_active"
	OSSTrustNetworkValBLocalApplicabilityStatePartial    = "oss_trust_network_valb_local_applicability_partial"
	OSSTrustNetworkValBLocalApplicabilityStateIncomplete = "oss_trust_network_valb_local_applicability_incomplete"
	OSSTrustNetworkValBLocalApplicabilityStateBlocked    = "oss_trust_network_valb_local_applicability_blocked"
	OSSTrustNetworkValBLocalApplicabilityStateUnknown    = "oss_trust_network_valb_local_applicability_unknown"

	OSSTrustNetworkValBPropagationExchangeStateActive     = "oss_trust_network_valb_propagation_exchange_active"
	OSSTrustNetworkValBPropagationExchangeStatePartial    = "oss_trust_network_valb_propagation_exchange_partial"
	OSSTrustNetworkValBPropagationExchangeStateIncomplete = "oss_trust_network_valb_propagation_exchange_incomplete"
	OSSTrustNetworkValBPropagationExchangeStateBlocked    = "oss_trust_network_valb_propagation_exchange_blocked"
	OSSTrustNetworkValBPropagationExchangeStateUnknown    = "oss_trust_network_valb_propagation_exchange_unknown"

	OSSTrustNetworkValBSupersessionRevocationStateActive     = "oss_trust_network_valb_supersession_revocation_active"
	OSSTrustNetworkValBSupersessionRevocationStatePartial    = "oss_trust_network_valb_supersession_revocation_partial"
	OSSTrustNetworkValBSupersessionRevocationStateIncomplete = "oss_trust_network_valb_supersession_revocation_incomplete"
	OSSTrustNetworkValBSupersessionRevocationStateBlocked    = "oss_trust_network_valb_supersession_revocation_blocked"
	OSSTrustNetworkValBSupersessionRevocationStateUnknown    = "oss_trust_network_valb_supersession_revocation_unknown"

	OSSTrustNetworkValBReviewerAuditabilityStateActive     = "oss_trust_network_valb_reviewer_auditability_active"
	OSSTrustNetworkValBReviewerAuditabilityStatePartial    = "oss_trust_network_valb_reviewer_auditability_partial"
	OSSTrustNetworkValBReviewerAuditabilityStateIncomplete = "oss_trust_network_valb_reviewer_auditability_incomplete"
	OSSTrustNetworkValBReviewerAuditabilityStateBlocked    = "oss_trust_network_valb_reviewer_auditability_blocked"
	OSSTrustNetworkValBReviewerAuditabilityStateUnknown    = "oss_trust_network_valb_reviewer_auditability_unknown"

	OSSTrustNetworkValBNoOverclaimStateActive     = "oss_trust_network_valb_no_overclaim_active"
	OSSTrustNetworkValBNoOverclaimStatePartial    = "oss_trust_network_valb_no_overclaim_partial"
	OSSTrustNetworkValBNoOverclaimStateIncomplete = "oss_trust_network_valb_no_overclaim_incomplete"
	OSSTrustNetworkValBNoOverclaimStateBlocked    = "oss_trust_network_valb_no_overclaim_blocked"
	OSSTrustNetworkValBNoOverclaimStateUnknown    = "oss_trust_network_valb_no_overclaim_unknown"

	OSSTrustNetworkValBCandidateSourceClassMaintainer            = "maintainer"
	OSSTrustNetworkValBCandidateSourceClassRegistry              = "registry"
	OSSTrustNetworkValBCandidateSourceClassCommunity             = "community"
	OSSTrustNetworkValBCandidateSourceClassEnterpriseObservation = "enterprise_observation"
	OSSTrustNetworkValBCandidateSourceClassVendor                = "vendor"
	OSSTrustNetworkValBCandidateSourceClassVerifier              = "verifier"
	OSSTrustNetworkValBCandidateSourceClassAutomatedHeuristic    = "automated_heuristic"

	OSSTrustNetworkValBCandidateIntakeStateReceived         = "received"
	OSSTrustNetworkValBCandidateIntakeStateNormalized       = "normalized"
	OSSTrustNetworkValBCandidateIntakeStateRejectedAtIntake = "rejected_at_intake"
	OSSTrustNetworkValBCandidateIntakeStateUnsupported      = "unsupported"
	OSSTrustNetworkValBCandidateIntakeStateStale            = "stale"
	OSSTrustNetworkValBCandidateIntakeStateMalformed        = "malformed"
	OSSTrustNetworkValBCandidateIntakeStateUnknown          = "unknown"

	OSSTrustNetworkValBReviewStateCandidate  = "candidate"
	OSSTrustNetworkValBReviewStateInReview   = "in_review"
	OSSTrustNetworkValBReviewStateReviewed   = "reviewed"
	OSSTrustNetworkValBReviewStateRejected   = "rejected"
	OSSTrustNetworkValBReviewStateSuperseded = "superseded"
	OSSTrustNetworkValBReviewStateRevoked    = "revoked"

	OSSTrustNetworkValBReviewerDecisionStateNone              = "none"
	OSSTrustNetworkValBReviewerDecisionStateAccepted          = "accepted"
	OSSTrustNetworkValBReviewerDecisionStateRejected          = "rejected"
	OSSTrustNetworkValBReviewerDecisionStateSuperseded        = "superseded"
	OSSTrustNetworkValBReviewerDecisionStateRevoked           = "revoked"
	OSSTrustNetworkValBReviewerDecisionStateNeedsMoreEvidence = "needs_more_evidence"

	OSSTrustNetworkValBSharedVEXStateCandidate   = "candidate"
	OSSTrustNetworkValBSharedVEXStateReviewed    = "reviewed"
	OSSTrustNetworkValBSharedVEXStateRejected    = "rejected"
	OSSTrustNetworkValBSharedVEXStateSuperseded  = "superseded"
	OSSTrustNetworkValBSharedVEXStateRevoked     = "revoked"
	OSSTrustNetworkValBSharedVEXStateUnsupported = "unsupported"
	OSSTrustNetworkValBSharedVEXStateUnknown     = "unknown"

	OSSTrustNetworkValBSourceWeightClassLow     = "low"
	OSSTrustNetworkValBSourceWeightClassMedium  = "medium"
	OSSTrustNetworkValBSourceWeightClassHigh    = "high"
	OSSTrustNetworkValBSourceWeightClassBounded = "bounded"

	OSSTrustNetworkValBLocalApplicabilityStatusApplicable       = "applicable"
	OSSTrustNetworkValBLocalApplicabilityStatusNotApplicable    = "not_applicable"
	OSSTrustNetworkValBLocalApplicabilityStatusUnknown          = "unknown"
	OSSTrustNetworkValBLocalApplicabilityStatusNeedsLocalReview = "needs_local_review"
	OSSTrustNetworkValBLocalApplicabilityStatusUnsupported      = "unsupported"

	OSSTrustNetworkValBPropagationStateNotShared         = "not_shared"
	OSSTrustNetworkValBPropagationStateCandidateExchange = "candidate_exchange"
	OSSTrustNetworkValBPropagationStateReviewedExchange  = "reviewed_exchange"
	OSSTrustNetworkValBPropagationStateRejected          = "rejected"
	OSSTrustNetworkValBPropagationStateRevoked           = "revoked"
	OSSTrustNetworkValBPropagationStateSuperseded        = "superseded"
	OSSTrustNetworkValBPropagationStateUnsupported       = "unsupported"
	OSSTrustNetworkValBPropagationStateUnknown           = "unknown"

	OSSTrustNetworkValBLifecycleStateActive     = "active"
	OSSTrustNetworkValBLifecycleStateSuperseded = "superseded"
	OSSTrustNetworkValBLifecycleStateRevoked    = "revoked"
	OSSTrustNetworkValBLifecycleStateUnknown    = "unknown"
)

type OSSTrustNetworkValBDependencySnapshot struct {
	ValACurrentState            string   `json:"vala_current_state"`
	ValAPoint9State             string   `json:"vala_point_9_state"`
	ValADependencyState         string   `json:"vala_dependency_state"`
	ValAReleaseTrustIntakeState string   `json:"vala_release_trust_intake_state"`
	ValASigningSignalState      string   `json:"vala_signing_signal_state"`
	ValAMaintainerState         string   `json:"vala_maintainer_attestation_state"`
	ValAProvenanceState         string   `json:"vala_provenance_material_state"`
	ValARegistryDescriptorState string   `json:"vala_registry_descriptor_state"`
	ValARegistryMetadataState   string   `json:"vala_registry_metadata_state"`
	ValATypoWarningState        string   `json:"vala_typo_squatting_warning_state"`
	ValADriftSignalState        string   `json:"vala_drift_signal_state"`
	ValANoOverclaimState        string   `json:"vala_no_overclaim_state"`
	ValAProofSurfaceRefs        []string `json:"vala_proof_surface_refs,omitempty"`
	ValAEvidenceRefs            []string `json:"vala_evidence_refs,omitempty"`
	ValAProjectionDisclaimer    string   `json:"vala_projection_disclaimer"`
}

type OSSTrustNetworkValBCandidateSignalIntake struct {
	CurrentState             string   `json:"current_state"`
	CandidateSignalID        string   `json:"candidate_signal_id"`
	PackageOrProjectIdentity string   `json:"package_or_project_identity"`
	RegistryOrEcosystem      string   `json:"registry_or_ecosystem"`
	ReleaseOrVersionRef      string   `json:"release_or_version_ref"`
	CandidateSourceClass     string   `json:"candidate_source_class"`
	SourceWeightClass        string   `json:"source_weight_class"`
	IntakeState              string   `json:"intake_state"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	FreshnessState           string   `json:"freshness_state"`
	ApplicabilityScope       string   `json:"applicability_scope"`
	Caveats                  []string `json:"caveats,omitempty"`
	CreatesReviewedTrust     bool     `json:"creates_reviewed_trust"`
	CanonicalTruthClaim      bool     `json:"canonical_truth_claim"`
	GlobalBlocklistClaim     bool     `json:"global_blocklist_claim"`
	EnterpriseOverrideClaim  bool     `json:"enterprise_override_claim"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValBReviewWorkflow struct {
	CurrentState                  string   `json:"current_state"`
	WorkflowID                    string   `json:"workflow_id"`
	ReviewState                   string   `json:"review_state"`
	ReviewerDecisionState         string   `json:"reviewer_decision_state"`
	SourceWeightClass             string   `json:"source_weight_class"`
	EvidenceRefs                  []string `json:"evidence_refs,omitempty"`
	FreshnessState                string   `json:"freshness_state"`
	Caveats                       []string `json:"caveats,omitempty"`
	ReviewerRationale             string   `json:"reviewer_rationale"`
	ReplacementRef                string   `json:"replacement_ref"`
	RejectedPropagatedUsable      bool     `json:"rejected_propagated_usable"`
	RewritesCanonicalEvidence     bool     `json:"rewrites_canonical_evidence"`
	RewritesLocalEnterpriseResult bool     `json:"rewrites_local_enterprise_result"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValBSharedVEXTriage struct {
	CurrentState                          string   `json:"current_state"`
	VEXOrTriageID                         string   `json:"vex_or_triage_id"`
	VulnerabilityOrAdvisoryRef            string   `json:"vulnerability_or_advisory_ref"`
	PackageOrProjectIdentity              string   `json:"package_or_project_identity"`
	AffectedReleaseScope                  string   `json:"affected_release_scope"`
	ExploitabilityContext                 string   `json:"exploitability_context"`
	LocalApplicabilityNote                string   `json:"local_applicability_note"`
	ReviewState                           string   `json:"review_state"`
	EvidenceRefs                          []string `json:"evidence_refs,omitempty"`
	Caveats                               []string `json:"caveats,omitempty"`
	SupersedesRef                         string   `json:"supersedes_ref"`
	LocalApplicabilityBounded             bool     `json:"local_applicability_bounded"`
	OverridesLocalEnterpriseApplicability bool     `json:"overrides_local_enterprise_applicability"`
	CanonicalTruthClaim                   bool     `json:"canonical_truth_claim"`
	AutomaticGlobalTruthClaim             bool     `json:"automatic_global_truth_claim"`
	ProjectionDisclaimer                  string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValBSourceWeighting struct {
	CurrentState                 string   `json:"current_state"`
	DisciplineID                 string   `json:"discipline_id"`
	Version                      string   `json:"version"`
	SupportedSourceClasses       []string `json:"supported_source_classes,omitempty"`
	SupportedSourceWeightClasses []string `json:"supported_source_weight_classes,omitempty"`
	SourceClass                  string   `json:"source_class"`
	SourceWeightClass            string   `json:"source_weight_class"`
	EvidenceRefs                 []string `json:"evidence_refs,omitempty"`
	Caveats                      []string `json:"caveats,omitempty"`
	FreshnessState               string   `json:"freshness_state"`
	ReviewState                  string   `json:"review_state"`
	AutomatedHeuristicStandalone bool     `json:"automated_heuristic_standalone"`
	CommunityInputWithoutReview  bool     `json:"community_input_without_review"`
	UniversalTrustScoreClaim     bool     `json:"universal_trust_score_claim"`
	IntegrityScoreClaim          bool     `json:"integrity_score_claim"`
	BadgeScoreClaim              bool     `json:"badge_score_claim"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValBLocalApplicability struct {
	CurrentState         string   `json:"current_state"`
	DisciplineID         string   `json:"discipline_id"`
	Version              string   `json:"version"`
	ApplicabilityState   string   `json:"applicability_state"`
	LocalScope           string   `json:"local_scope"`
	EvidenceRefs         []string `json:"evidence_refs,omitempty"`
	Caveats              []string `json:"caveats,omitempty"`
	Rationale            string   `json:"rationale"`
	LocalEvidenceLinked  bool     `json:"local_evidence_linked"`
	OverrideClaim        bool     `json:"override_claim"`
	SharedSignalHidden   bool     `json:"shared_signal_hidden"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValBPropagationExchange struct {
	CurrentState            string   `json:"current_state"`
	DisciplineID            string   `json:"discipline_id"`
	Version                 string   `json:"version"`
	PropagationState        string   `json:"propagation_state"`
	ReviewState             string   `json:"review_state"`
	SourceWeightClass       string   `json:"source_weight_class"`
	FreshnessState          string   `json:"freshness_state"`
	LocalApplicabilityState string   `json:"local_applicability_state"`
	ApplicabilityScope      string   `json:"applicability_scope"`
	EvidenceRefs            []string `json:"evidence_refs,omitempty"`
	Caveats                 []string `json:"caveats,omitempty"`
	ReplacementRef          string   `json:"replacement_ref"`
	PresentedAsReviewed     bool     `json:"presented_as_reviewed"`
	SimilarityContextGating bool     `json:"similarity_context_gating"`
	AutomaticGlobalSpread   bool     `json:"automatic_global_spread"`
	GlobalBlocklistClaim    bool     `json:"global_blocklist_claim"`
	EnterpriseOverride      bool     `json:"enterprise_override"`
	ProjectionDisclaimer    string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValBSupersessionRevocation struct {
	CurrentState                string   `json:"current_state"`
	DisciplineID                string   `json:"discipline_id"`
	Version                     string   `json:"version"`
	LifecycleState              string   `json:"lifecycle_state"`
	PreviousSignalID            string   `json:"previous_signal_id"`
	ReplacementRef              string   `json:"replacement_ref"`
	SupersessionReason          string   `json:"supersession_reason"`
	RevocationReason            string   `json:"revocation_reason"`
	RevocationTimestamp         string   `json:"revocation_timestamp"`
	EvidenceRefs                []string `json:"evidence_refs,omitempty"`
	Caveats                     []string `json:"caveats,omitempty"`
	PreviousIdentityPreserved   bool     `json:"previous_identity_preserved"`
	ReviewedExchangeStillActive bool     `json:"reviewed_exchange_still_active"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValBReviewerAuditability struct {
	CurrentState              string   `json:"current_state"`
	DisciplineID              string   `json:"discipline_id"`
	Version                   string   `json:"version"`
	ReviewerRoleClass         string   `json:"reviewer_role_class"`
	DecisionState             string   `json:"decision_state"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	Rationale                 string   `json:"rationale"`
	DecisionTimestamp         string   `json:"decision_timestamp"`
	Caveats                   []string `json:"caveats,omitempty"`
	LegalCertificationClaim   bool     `json:"legal_certification_claim"`
	RegulatorApprovalClaim    bool     `json:"regulator_approval_claim"`
	OfficialOSSAuthorityClaim bool     `json:"official_oss_authority_claim"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValBNoOverclaim struct {
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
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValBCore struct {
	CurrentState                string                                    `json:"current_state"`
	Point9State                 string                                    `json:"point_9_state"`
	DependencyState             string                                    `json:"dependency_state"`
	CandidateSignalIntakeState  string                                    `json:"candidate_signal_intake_state"`
	ReviewWorkflowState         string                                    `json:"review_workflow_state"`
	SharedVEXTriageState        string                                    `json:"shared_vex_triage_state"`
	SourceWeightingState        string                                    `json:"source_weighting_state"`
	LocalApplicabilityState     string                                    `json:"local_applicability_state"`
	PropagationExchangeState    string                                    `json:"propagation_exchange_state"`
	SupersessionRevocationState string                                    `json:"supersession_revocation_state"`
	ReviewerAuditabilityState   string                                    `json:"reviewer_auditability_state"`
	NoOverclaimState            string                                    `json:"no_overclaim_state"`
	Dependency                  OSSTrustNetworkValBDependencySnapshot     `json:"dependency"`
	CandidateSignalIntake       OSSTrustNetworkValBCandidateSignalIntake  `json:"candidate_signal_intake"`
	ReviewWorkflow              OSSTrustNetworkValBReviewWorkflow         `json:"review_workflow"`
	SharedVEXTriage             OSSTrustNetworkValBSharedVEXTriage        `json:"shared_vex_triage"`
	SourceWeighting             OSSTrustNetworkValBSourceWeighting        `json:"source_weighting"`
	LocalApplicability          OSSTrustNetworkValBLocalApplicability     `json:"local_applicability"`
	PropagationExchange         OSSTrustNetworkValBPropagationExchange    `json:"propagation_exchange"`
	SupersessionRevocation      OSSTrustNetworkValBSupersessionRevocation `json:"supersession_revocation"`
	ReviewerAuditability        OSSTrustNetworkValBReviewerAuditability   `json:"reviewer_auditability"`
	NoOverclaim                 OSSTrustNetworkValBNoOverclaim            `json:"no_overclaim"`
	ProofSurfaceRefs            []string                                  `json:"proof_surface_refs,omitempty"`
	EvidenceRefs                []string                                  `json:"evidence_refs,omitempty"`
	BlockingReasons             []string                                  `json:"blocking_reasons,omitempty"`
	WhyPoint9NotComplete        []string                                  `json:"why_point_9_not_complete,omitempty"`
	ProjectionDisclaimer        string                                    `json:"projection_disclaimer"`
}

func ossTrustNetworkValBProjectionDisclaimer() string {
	return "projection_only not_canonical_truth oss_trust_network_valb advisory_projection"
}

func ossTrustNetworkValBHasProjectionDisclaimer(value string) bool {
	normalized := strings.ToLower(strings.TrimSpace(value))
	return strings.Contains(normalized, "projection_only") &&
		strings.Contains(normalized, "not_canonical_truth") &&
		strings.Contains(normalized, "oss_trust_network_valb")
}

func ossTrustNetworkValBCandidateSourceClasses() []string {
	return []string{
		OSSTrustNetworkValBCandidateSourceClassMaintainer,
		OSSTrustNetworkValBCandidateSourceClassRegistry,
		OSSTrustNetworkValBCandidateSourceClassCommunity,
		OSSTrustNetworkValBCandidateSourceClassEnterpriseObservation,
		OSSTrustNetworkValBCandidateSourceClassVendor,
		OSSTrustNetworkValBCandidateSourceClassVerifier,
		OSSTrustNetworkValBCandidateSourceClassAutomatedHeuristic,
	}
}

func ossTrustNetworkValBCandidateIntakeStates() []string {
	return []string{
		OSSTrustNetworkValBCandidateIntakeStateReceived,
		OSSTrustNetworkValBCandidateIntakeStateNormalized,
		OSSTrustNetworkValBCandidateIntakeStateRejectedAtIntake,
		OSSTrustNetworkValBCandidateIntakeStateUnsupported,
		OSSTrustNetworkValBCandidateIntakeStateStale,
		OSSTrustNetworkValBCandidateIntakeStateMalformed,
		OSSTrustNetworkValBCandidateIntakeStateUnknown,
	}
}

func ossTrustNetworkValBReviewStates() []string {
	return []string{
		OSSTrustNetworkValBReviewStateCandidate,
		OSSTrustNetworkValBReviewStateInReview,
		OSSTrustNetworkValBReviewStateReviewed,
		OSSTrustNetworkValBReviewStateRejected,
		OSSTrustNetworkValBReviewStateSuperseded,
		OSSTrustNetworkValBReviewStateRevoked,
	}
}

func ossTrustNetworkValBReviewerDecisionStates() []string {
	return []string{
		OSSTrustNetworkValBReviewerDecisionStateNone,
		OSSTrustNetworkValBReviewerDecisionStateAccepted,
		OSSTrustNetworkValBReviewerDecisionStateRejected,
		OSSTrustNetworkValBReviewerDecisionStateSuperseded,
		OSSTrustNetworkValBReviewerDecisionStateRevoked,
		OSSTrustNetworkValBReviewerDecisionStateNeedsMoreEvidence,
	}
}

func ossTrustNetworkValBSharedVEXStates() []string {
	return []string{
		OSSTrustNetworkValBSharedVEXStateCandidate,
		OSSTrustNetworkValBSharedVEXStateReviewed,
		OSSTrustNetworkValBSharedVEXStateRejected,
		OSSTrustNetworkValBSharedVEXStateSuperseded,
		OSSTrustNetworkValBSharedVEXStateRevoked,
		OSSTrustNetworkValBSharedVEXStateUnsupported,
		OSSTrustNetworkValBSharedVEXStateUnknown,
	}
}

func ossTrustNetworkValBSourceWeightClasses() []string {
	return []string{
		OSSTrustNetworkValBSourceWeightClassLow,
		OSSTrustNetworkValBSourceWeightClassMedium,
		OSSTrustNetworkValBSourceWeightClassHigh,
		OSSTrustNetworkValBSourceWeightClassBounded,
	}
}

func ossTrustNetworkValBLocalApplicabilityStates() []string {
	return []string{
		OSSTrustNetworkValBLocalApplicabilityStatusApplicable,
		OSSTrustNetworkValBLocalApplicabilityStatusNotApplicable,
		OSSTrustNetworkValBLocalApplicabilityStatusUnknown,
		OSSTrustNetworkValBLocalApplicabilityStatusNeedsLocalReview,
		OSSTrustNetworkValBLocalApplicabilityStatusUnsupported,
	}
}

func ossTrustNetworkValBPropagationStates() []string {
	return []string{
		OSSTrustNetworkValBPropagationStateNotShared,
		OSSTrustNetworkValBPropagationStateCandidateExchange,
		OSSTrustNetworkValBPropagationStateReviewedExchange,
		OSSTrustNetworkValBPropagationStateRejected,
		OSSTrustNetworkValBPropagationStateRevoked,
		OSSTrustNetworkValBPropagationStateSuperseded,
		OSSTrustNetworkValBPropagationStateUnsupported,
		OSSTrustNetworkValBPropagationStateUnknown,
	}
}

func ossTrustNetworkValBLifecycleStates() []string {
	return []string{
		OSSTrustNetworkValBLifecycleStateActive,
		OSSTrustNetworkValBLifecycleStateSuperseded,
		OSSTrustNetworkValBLifecycleStateRevoked,
		OSSTrustNetworkValBLifecycleStateUnknown,
	}
}

func OSSTrustNetworkValBProofSurfaceRefs() []string {
	return []string{
		"/v1/oss-trust-network/valb/status",
		"/v1/oss-trust-network/valb/proofs",
	}
}

func OSSTrustNetworkValBProofEvidenceRefs() []string {
	return []string{
		"evidence:ostn-valb-dependency-001",
		"evidence:ostn-valb-candidate-intake-001",
		"evidence:ostn-valb-review-workflow-001",
		"evidence:ostn-valb-shared-vex-001",
		"evidence:ostn-valb-source-weighting-001",
		"evidence:ostn-valb-local-applicability-001",
		"evidence:ostn-valb-propagation-001",
		"evidence:ostn-valb-supersession-revocation-001",
		"evidence:ostn-valb-reviewer-auditability-001",
		"evidence:ostn-valb-no-overclaim-001",
		"evidence:ostn-valb-point9-governance-001",
	}
}

type ossTrustNetworkValBExpectedEvidenceMetadata struct {
	EvidenceType string
	Source       string
	Scope        string
}

func ossTrustNetworkValBExpectedEvidenceMetadataByID() map[string]ossTrustNetworkValBExpectedEvidenceMetadata {
	return map[string]ossTrustNetworkValBExpectedEvidenceMetadata{
		"evidence:ostn-valb-dependency-001":              {EvidenceType: "dependency_state", Source: "oss-trust-network/valb/dependency", Scope: "vala_dependency"},
		"evidence:ostn-valb-candidate-intake-001":        {EvidenceType: "candidate_signal_intake", Source: "oss-trust-network/valb/candidate-intake", Scope: "candidate_signal_intake"},
		"evidence:ostn-valb-review-workflow-001":         {EvidenceType: "review_workflow", Source: "oss-trust-network/valb/review-workflow", Scope: "review_workflow"},
		"evidence:ostn-valb-shared-vex-001":              {EvidenceType: "shared_vex_triage", Source: "oss-trust-network/valb/shared-vex", Scope: "shared_vex_triage"},
		"evidence:ostn-valb-source-weighting-001":        {EvidenceType: "source_weighting", Source: "oss-trust-network/valb/source-weighting", Scope: "source_weighting"},
		"evidence:ostn-valb-local-applicability-001":     {EvidenceType: "local_applicability", Source: "oss-trust-network/valb/local-applicability", Scope: "local_applicability"},
		"evidence:ostn-valb-propagation-001":             {EvidenceType: "propagation_exchange", Source: "oss-trust-network/valb/propagation", Scope: "reviewed_exchange"},
		"evidence:ostn-valb-supersession-revocation-001": {EvidenceType: "supersession_revocation", Source: "oss-trust-network/valb/supersession-revocation", Scope: "supersession_revocation"},
		"evidence:ostn-valb-reviewer-auditability-001":   {EvidenceType: "reviewer_auditability", Source: "oss-trust-network/valb/reviewer-auditability", Scope: "reviewer_auditability"},
		"evidence:ostn-valb-no-overclaim-001":            {EvidenceType: "no_overclaim", Source: "oss-trust-network/valb/no-overclaim", Scope: "no_overclaim_discipline"},
		"evidence:ostn-valb-point9-governance-001":       {EvidenceType: "state_governance", Source: "oss-trust-network/valb/point9-governance", Scope: "point9_governance"},
	}
}

func ossTrustNetworkValBEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:ostn-valb-dependency-001", EvidenceType: "dependency_state", Source: "oss-trust-network/valb/dependency", Timestamp: "2026-04-29T11:10:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "vala_dependency", Caveats: []string{"Val B depends on exact and active OSTN Val A only."}},
		{EvidenceID: "evidence:ostn-valb-candidate-intake-001", EvidenceType: "candidate_signal_intake", Source: "oss-trust-network/valb/candidate-intake", Timestamp: "2026-04-29T11:11:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "candidate_signal_intake", Caveats: []string{"Candidate OSS trust signals remain reviewable and bounded intake only."}},
		{EvidenceID: "evidence:ostn-valb-review-workflow-001", EvidenceType: "review_workflow", Source: "oss-trust-network/valb/review-workflow", Timestamp: "2026-04-29T11:12:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "review_workflow", Caveats: []string{"Reviewed shared intelligence requires explicit decision, rationale, freshness, and caveats."}},
		{EvidenceID: "evidence:ostn-valb-shared-vex-001", EvidenceType: "shared_vex_triage", Source: "oss-trust-network/valb/shared-vex", Timestamp: "2026-04-29T11:13:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "shared_vex_triage", Caveats: []string{"Shared VEX remains reviewed intelligence, not global truth or enterprise override."}},
		{EvidenceID: "evidence:ostn-valb-source-weighting-001", EvidenceType: "source_weighting", Source: "oss-trust-network/valb/source-weighting", Timestamp: "2026-04-29T11:14:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "source_weighting", Caveats: []string{"Source weighting remains advisory and cannot become a universal trust or integrity score."}},
		{EvidenceID: "evidence:ostn-valb-local-applicability-001", EvidenceType: "local_applicability", Source: "oss-trust-network/valb/local-applicability", Timestamp: "2026-04-29T11:15:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "local_applicability", Caveats: []string{"Shared intelligence remains bounded by local applicability and local evidence context."}},
		{EvidenceID: "evidence:ostn-valb-propagation-001", EvidenceType: "propagation_exchange", Source: "oss-trust-network/valb/propagation", Timestamp: "2026-04-29T11:16:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "reviewed_exchange", Caveats: []string{"Propagation remains bounded reviewed exchange only and cannot auto-spread as network truth."}},
		{EvidenceID: "evidence:ostn-valb-supersession-revocation-001", EvidenceType: "supersession_revocation", Source: "oss-trust-network/valb/supersession-revocation", Timestamp: "2026-04-29T11:17:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "supersession_revocation", Caveats: []string{"Superseded and revoked signals remain bounded lifecycle states with preserved identity and reasons."}},
		{EvidenceID: "evidence:ostn-valb-reviewer-auditability-001", EvidenceType: "reviewer_auditability", Source: "oss-trust-network/valb/reviewer-auditability", Timestamp: "2026-04-29T11:18:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "reviewer_auditability", Caveats: []string{"Reviewer auditability remains evidence-linked decision traceability, not legal or regulatory authority."}},
		{EvidenceID: "evidence:ostn-valb-no-overclaim-001", EvidenceType: "no_overclaim", Source: "oss-trust-network/valb/no-overclaim", Timestamp: "2026-04-29T11:19:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim_discipline", Caveats: []string{"Val B cannot create global truth, reviewed-means-safe, or forbidden point pass claims."}},
		{EvidenceID: "evidence:ostn-valb-point9-governance-001", EvidenceType: "state_governance", Source: "oss-trust-network/valb/point9-governance", Timestamp: "2026-04-29T11:20:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point9_governance", Caveats: []string{"point_9_state remains not complete and later waves are required for any final pass."}},
	}
}

func OSSTrustNetworkValBProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, refs []string) bool {
	if !containsExactTrimmedStringSet(refs, OSSTrustNetworkValBProofEvidenceRefs()...) {
		return false
	}
	expected := ossTrustNetworkValBExpectedEvidenceMetadataByID()
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

func ossTrustNetworkValBContainsForbiddenClaim(values ...string) bool {
	allowed := map[string]struct{}{
		"bounded shared intelligence":     {},
		"reviewed oss trust signal":       {},
		"candidate oss trust signal":      {},
		"source-weighted reviewed signal": {},
		"shared vex review workflow":      {},
		"local applicability context":     {},
		"bounded reviewed exchange":       {},
		"advisory oss network signal":     {},
		"evidence-linked review decision": {},
		"superseded signal":               {},
		"revoked signal":                  {},
		"not canonical truth":             {},
		"not formal certification":        {},
	}
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

func OSSTrustNetworkValBDependencySnapshotModel() OSSTrustNetworkValBDependencySnapshot {
	model := ComputeOSSTrustNetworkValACore(OSSTrustNetworkValACoreModel())
	return OSSTrustNetworkValBDependencySnapshot{
		ValACurrentState:            model.CurrentState,
		ValAPoint9State:             model.Point9State,
		ValADependencyState:         model.DependencyState,
		ValAReleaseTrustIntakeState: model.ReleaseTrustIntakeState,
		ValASigningSignalState:      model.SigningSignalState,
		ValAMaintainerState:         model.MaintainerAttestationState,
		ValAProvenanceState:         model.ProvenanceMaterialState,
		ValARegistryDescriptorState: model.RegistryDescriptorState,
		ValARegistryMetadataState:   model.RegistryMetadataState,
		ValATypoWarningState:        model.TypoSquattingWarningState,
		ValADriftSignalState:        model.DriftSignalState,
		ValANoOverclaimState:        model.NoOverclaimState,
		ValAProofSurfaceRefs:        append([]string{}, model.ProofSurfaceRefs...),
		ValAEvidenceRefs:            append([]string{}, model.EvidenceRefs...),
		ValAProjectionDisclaimer:    model.ProjectionDisclaimer,
	}
}

func OSSTrustNetworkValBCandidateSignalIntakeModel() OSSTrustNetworkValBCandidateSignalIntake {
	return OSSTrustNetworkValBCandidateSignalIntake{
		CurrentState:             OSSTrustNetworkValBCandidateSignalIntakeStateActive,
		CandidateSignalID:        "ostn-valb-candidate-001",
		PackageOrProjectIdentity: "github.com/example/project",
		RegistryOrEcosystem:      "github",
		ReleaseOrVersionRef:      "refs/tags/v1.2.3",
		CandidateSourceClass:     OSSTrustNetworkValBCandidateSourceClassVerifier,
		SourceWeightClass:        OSSTrustNetworkValBSourceWeightClassMedium,
		IntakeState:              OSSTrustNetworkValBCandidateIntakeStateNormalized,
		EvidenceRefs:             []string{"evidence:ostn-valb-candidate-intake-001"},
		FreshnessState:           IntelligenceCalibrationFreshnessFresh,
		ApplicabilityScope:       OSSTrustNetworkApplicabilityEnterpriseLocal,
		Caveats:                  []string{"candidate intake remains bounded and reviewable"},
		ProjectionDisclaimer:     ossTrustNetworkValBProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValBReviewWorkflowModel() OSSTrustNetworkValBReviewWorkflow {
	return OSSTrustNetworkValBReviewWorkflow{
		CurrentState:          OSSTrustNetworkValBReviewWorkflowStateActive,
		WorkflowID:            "ostn-valb-review-001",
		ReviewState:           OSSTrustNetworkValBReviewStateReviewed,
		ReviewerDecisionState: OSSTrustNetworkValBReviewerDecisionStateAccepted,
		SourceWeightClass:     OSSTrustNetworkValBSourceWeightClassMedium,
		EvidenceRefs:          []string{"evidence:ostn-valb-review-workflow-001"},
		FreshnessState:        IntelligenceCalibrationFreshnessFresh,
		Caveats:               []string{"reviewed signal remains bounded shared intelligence"},
		ReviewerRationale:     "bounded reviewed signal accepted with freshness, weighting, and evidence linkage",
		ProjectionDisclaimer:  ossTrustNetworkValBProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValBSharedVEXTriageModel() OSSTrustNetworkValBSharedVEXTriage {
	return OSSTrustNetworkValBSharedVEXTriage{
		CurrentState:               OSSTrustNetworkValBSharedVEXTriageStateActive,
		VEXOrTriageID:              "ostn-valb-vex-001",
		VulnerabilityOrAdvisoryRef: "CVE-2026-0001",
		PackageOrProjectIdentity:   "github.com/example/project",
		AffectedReleaseScope:       "refs/tags/v1.2.3",
		ExploitabilityContext:      "bounded reviewed exploitability context",
		LocalApplicabilityNote:     "local applicability remains enterprise-bounded",
		ReviewState:                OSSTrustNetworkValBSharedVEXStateReviewed,
		EvidenceRefs:               []string{"evidence:ostn-valb-shared-vex-001"},
		Caveats:                    []string{"shared VEX remains advisory reviewed intelligence"},
		LocalApplicabilityBounded:  true,
		ProjectionDisclaimer:       ossTrustNetworkValBProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValBSourceWeightingModel() OSSTrustNetworkValBSourceWeighting {
	return OSSTrustNetworkValBSourceWeighting{
		CurrentState:                 OSSTrustNetworkValBSourceWeightingStateActive,
		DisciplineID:                 "oss_source_weighting",
		Version:                      "v0",
		SupportedSourceClasses:       ossTrustNetworkValBCandidateSourceClasses(),
		SupportedSourceWeightClasses: ossTrustNetworkValBSourceWeightClasses(),
		SourceClass:                  OSSTrustNetworkValBCandidateSourceClassVerifier,
		SourceWeightClass:            OSSTrustNetworkValBSourceWeightClassMedium,
		EvidenceRefs:                 []string{"evidence:ostn-valb-source-weighting-001"},
		Caveats:                      []string{"source weighting remains bounded advisory weighting"},
		FreshnessState:               IntelligenceCalibrationFreshnessFresh,
		ReviewState:                  OSSTrustNetworkValBReviewStateReviewed,
		ProjectionDisclaimer:         ossTrustNetworkValBProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValBLocalApplicabilityModel() OSSTrustNetworkValBLocalApplicability {
	return OSSTrustNetworkValBLocalApplicability{
		CurrentState:         OSSTrustNetworkValBLocalApplicabilityStateActive,
		DisciplineID:         "oss_local_applicability_valb",
		Version:              "v0",
		ApplicabilityState:   OSSTrustNetworkValBLocalApplicabilityStatusApplicable,
		LocalScope:           OSSTrustNetworkApplicabilityEnterpriseLocal,
		EvidenceRefs:         []string{"evidence:ostn-valb-local-applicability-001"},
		Caveats:              []string{"local applicability remains evidence-linked and local"},
		Rationale:            "local applicability confirmed within bounded enterprise scope",
		LocalEvidenceLinked:  true,
		ProjectionDisclaimer: ossTrustNetworkValBProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValBPropagationExchangeModel() OSSTrustNetworkValBPropagationExchange {
	return OSSTrustNetworkValBPropagationExchange{
		CurrentState:            OSSTrustNetworkValBPropagationExchangeStateActive,
		DisciplineID:            "oss_propagation_exchange",
		Version:                 "v0",
		PropagationState:        OSSTrustNetworkValBPropagationStateReviewedExchange,
		ReviewState:             OSSTrustNetworkValBReviewStateReviewed,
		SourceWeightClass:       OSSTrustNetworkValBSourceWeightClassMedium,
		FreshnessState:          IntelligenceCalibrationFreshnessFresh,
		LocalApplicabilityState: OSSTrustNetworkValBLocalApplicabilityStatusApplicable,
		ApplicabilityScope:      OSSTrustNetworkApplicabilityEnterpriseLocal,
		EvidenceRefs:            []string{"evidence:ostn-valb-propagation-001"},
		Caveats:                 []string{"reviewed exchange remains bounded and context-gated"},
		SimilarityContextGating: true,
		ProjectionDisclaimer:    ossTrustNetworkValBProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValBSupersessionRevocationModel() OSSTrustNetworkValBSupersessionRevocation {
	return OSSTrustNetworkValBSupersessionRevocation{
		CurrentState:              OSSTrustNetworkValBSupersessionRevocationStateActive,
		DisciplineID:              "oss_supersession_revocation",
		Version:                   "v0",
		LifecycleState:            OSSTrustNetworkValBLifecycleStateActive,
		PreviousSignalID:          "ostn-valb-review-001",
		EvidenceRefs:              []string{"evidence:ostn-valb-supersession-revocation-001"},
		Caveats:                   []string{"lifecycle discipline preserves signal identity and reasons"},
		PreviousIdentityPreserved: true,
		ProjectionDisclaimer:      ossTrustNetworkValBProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValBReviewerAuditabilityModel() OSSTrustNetworkValBReviewerAuditability {
	return OSSTrustNetworkValBReviewerAuditability{
		CurrentState:         OSSTrustNetworkValBReviewerAuditabilityStateActive,
		DisciplineID:         "oss_reviewer_auditability",
		Version:              "v0",
		ReviewerRoleClass:    "enterprise_reviewer",
		DecisionState:        OSSTrustNetworkValBReviewerDecisionStateAccepted,
		EvidenceRefs:         []string{"evidence:ostn-valb-reviewer-auditability-001"},
		Rationale:            "decision accepted with evidence-linked reviewer rationale",
		DecisionTimestamp:    "2026-04-29T11:18:30Z",
		Caveats:              []string{"reviewer decision remains bounded shared intelligence metadata"},
		ProjectionDisclaimer: ossTrustNetworkValBProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValBNoOverclaimModel() OSSTrustNetworkValBNoOverclaim {
	return OSSTrustNetworkValBNoOverclaim{
		CurrentState:         OSSTrustNetworkValBNoOverclaimStateActive,
		DisciplineID:         "oss_no_overclaim_valb",
		Version:              "v0",
		ObservedClaims:       []string{"bounded shared intelligence", "reviewed OSS trust signal", "not canonical truth"},
		ProjectionDisclaimer: ossTrustNetworkValBProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValBCoreModel() OSSTrustNetworkValBCore {
	return OSSTrustNetworkValBCore{
		CurrentState:                OSSTrustNetworkValBStateActive,
		Point9State:                 OSSTrustNetworkPoint9StateNotComplete,
		DependencyState:             OSSTrustNetworkValBDependencyStateActive,
		CandidateSignalIntakeState:  OSSTrustNetworkValBCandidateSignalIntakeStateActive,
		ReviewWorkflowState:         OSSTrustNetworkValBReviewWorkflowStateActive,
		SharedVEXTriageState:        OSSTrustNetworkValBSharedVEXTriageStateActive,
		SourceWeightingState:        OSSTrustNetworkValBSourceWeightingStateActive,
		LocalApplicabilityState:     OSSTrustNetworkValBLocalApplicabilityStateActive,
		PropagationExchangeState:    OSSTrustNetworkValBPropagationExchangeStateActive,
		SupersessionRevocationState: OSSTrustNetworkValBSupersessionRevocationStateActive,
		ReviewerAuditabilityState:   OSSTrustNetworkValBReviewerAuditabilityStateActive,
		NoOverclaimState:            OSSTrustNetworkValBNoOverclaimStateActive,
		Dependency:                  OSSTrustNetworkValBDependencySnapshotModel(),
		CandidateSignalIntake:       OSSTrustNetworkValBCandidateSignalIntakeModel(),
		ReviewWorkflow:              OSSTrustNetworkValBReviewWorkflowModel(),
		SharedVEXTriage:             OSSTrustNetworkValBSharedVEXTriageModel(),
		SourceWeighting:             OSSTrustNetworkValBSourceWeightingModel(),
		LocalApplicability:          OSSTrustNetworkValBLocalApplicabilityModel(),
		PropagationExchange:         OSSTrustNetworkValBPropagationExchangeModel(),
		SupersessionRevocation:      OSSTrustNetworkValBSupersessionRevocationModel(),
		ReviewerAuditability:        OSSTrustNetworkValBReviewerAuditabilityModel(),
		NoOverclaim:                 OSSTrustNetworkValBNoOverclaimModel(),
		ProofSurfaceRefs:            OSSTrustNetworkValBProofSurfaceRefs(),
		EvidenceRefs:                OSSTrustNetworkValBProofEvidenceRefs(),
		WhyPoint9NotComplete: []string{
			"Val B establishes bounded shared reviewed intelligence only and cannot complete Točka 9.",
			"Dashboards, remediation workflows, final gates, integrated closure, and later OSTN waves remain for later work.",
			"Shared reviewed intelligence remains advisory, evidence-linked, and not canonical truth or final pass authority.",
		},
		ProjectionDisclaimer: ossTrustNetworkValBProjectionDisclaimer(),
	}
}

func EvaluateOSSTrustNetworkValBDependencyState(model OSSTrustNetworkValBDependencySnapshot) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ValACurrentState,
		model.ValAPoint9State,
		model.ValADependencyState,
		model.ValAReleaseTrustIntakeState,
		model.ValASigningSignalState,
		model.ValAMaintainerState,
		model.ValAProvenanceState,
		model.ValARegistryDescriptorState,
		model.ValARegistryMetadataState,
		model.ValATypoWarningState,
		model.ValADriftSignalState,
		model.ValANoOverclaimState,
		model.ValAProjectionDisclaimer,
	) || len(model.ValAProofSurfaceRefs) == 0 || len(model.ValAEvidenceRefs) == 0 {
		return OSSTrustNetworkValBDependencyStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ValAProjectionDisclaimer) {
		return OSSTrustNetworkValBDependencyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.ValAProofSurfaceRefs, OSSTrustNetworkValAProofSurfaceRefs()...) ||
		!containsExactTrimmedStringSet(model.ValAEvidenceRefs, OSSTrustNetworkValAProofEvidenceRefs()...) ||
		!OSSTrustNetworkValAProofEvidenceQualityValid(ossTrustNetworkValAEvidence(), model.ValAEvidenceRefs) {
		return OSSTrustNetworkValBDependencyStateBlocked
	}
	if strings.TrimSpace(model.ValACurrentState) != OSSTrustNetworkValAStateActive ||
		strings.TrimSpace(model.ValAPoint9State) != OSSTrustNetworkPoint9StateNotComplete ||
		strings.TrimSpace(model.ValADependencyState) != OSSTrustNetworkValADependencyStateActive ||
		strings.TrimSpace(model.ValAReleaseTrustIntakeState) != OSSTrustNetworkValAReleaseTrustIntakeStateActive ||
		strings.TrimSpace(model.ValASigningSignalState) != OSSTrustNetworkValASigningSignalStateActive ||
		strings.TrimSpace(model.ValAMaintainerState) != OSSTrustNetworkValAMaintainerAttestationStateActive ||
		strings.TrimSpace(model.ValAProvenanceState) != OSSTrustNetworkValAProvenanceMaterialStateActive ||
		strings.TrimSpace(model.ValARegistryDescriptorState) != OSSTrustNetworkValARegistryDescriptorStateActive ||
		strings.TrimSpace(model.ValARegistryMetadataState) != OSSTrustNetworkValARegistryMetadataStateActive ||
		strings.TrimSpace(model.ValATypoWarningState) != OSSTrustNetworkValATypoSquattingWarningStateActive ||
		strings.TrimSpace(model.ValADriftSignalState) != OSSTrustNetworkValADriftSignalStateActive ||
		strings.TrimSpace(model.ValANoOverclaimState) != OSSTrustNetworkValANoOverclaimStateActive {
		return OSSTrustNetworkValBDependencyStateBlocked
	}
	return OSSTrustNetworkValBDependencyStateActive
}

func EvaluateOSSTrustNetworkValBCandidateSignalIntakeState(model OSSTrustNetworkValBCandidateSignalIntake) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CandidateSignalID,
		model.PackageOrProjectIdentity,
		model.RegistryOrEcosystem,
		model.ReleaseOrVersionRef,
		model.CandidateSourceClass,
		model.SourceWeightClass,
		model.IntakeState,
		model.FreshnessState,
		model.ApplicabilityScope,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValBCandidateSignalIntakeStateIncomplete
	}
	if !ossTrustNetworkValBHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValBCandidateSignalIntakeStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valb-candidate-intake-001") ||
		!containsTrimmedString(ossTrustNetworkValBCandidateSourceClasses(), model.CandidateSourceClass) ||
		!containsTrimmedString(ossTrustNetworkValBSourceWeightClasses(), model.SourceWeightClass) ||
		!containsTrimmedString(ossTrustNetworkValBCandidateIntakeStates(), model.IntakeState) ||
		!containsTrimmedString(ossTrustNetworkVal0ApplicabilityScopes(), model.ApplicabilityScope) ||
		model.CreatesReviewedTrust ||
		model.CanonicalTruthClaim ||
		model.GlobalBlocklistClaim ||
		model.EnterpriseOverrideClaim {
		return OSSTrustNetworkValBCandidateSignalIntakeStateBlocked
	}
	switch strings.TrimSpace(model.IntakeState) {
	case OSSTrustNetworkValBCandidateIntakeStateNormalized:
		if ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) {
			return OSSTrustNetworkValBCandidateSignalIntakeStateActive
		}
		return OSSTrustNetworkValBCandidateSignalIntakeStateBlocked
	case OSSTrustNetworkValBCandidateIntakeStateReceived:
		return OSSTrustNetworkValBCandidateSignalIntakeStatePartial
	case OSSTrustNetworkValBCandidateIntakeStateRejectedAtIntake,
		OSSTrustNetworkValBCandidateIntakeStateUnsupported,
		OSSTrustNetworkValBCandidateIntakeStateStale,
		OSSTrustNetworkValBCandidateIntakeStateMalformed,
		OSSTrustNetworkValBCandidateIntakeStateUnknown:
		return OSSTrustNetworkValBCandidateSignalIntakeStateBlocked
	default:
		return OSSTrustNetworkValBCandidateSignalIntakeStateBlocked
	}
}

func EvaluateOSSTrustNetworkValBReviewWorkflowState(model OSSTrustNetworkValBReviewWorkflow) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.WorkflowID,
		model.ReviewState,
		model.ReviewerDecisionState,
		model.SourceWeightClass,
		model.FreshnessState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValBReviewWorkflowStateIncomplete
	}
	if !ossTrustNetworkValBHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValBReviewWorkflowStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valb-review-workflow-001") ||
		!containsTrimmedString(ossTrustNetworkValBReviewStates(), model.ReviewState) ||
		!containsTrimmedString(ossTrustNetworkValBReviewerDecisionStates(), model.ReviewerDecisionState) ||
		!containsTrimmedString(ossTrustNetworkValBSourceWeightClasses(), model.SourceWeightClass) ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) ||
		model.RewritesCanonicalEvidence ||
		model.RewritesLocalEnterpriseResult {
		return OSSTrustNetworkValBReviewWorkflowStateBlocked
	}
	switch strings.TrimSpace(model.ReviewState) {
	case OSSTrustNetworkValBReviewStateReviewed:
		if strings.TrimSpace(model.ReviewerDecisionState) == OSSTrustNetworkValBReviewerDecisionStateAccepted &&
			strings.TrimSpace(model.ReviewerRationale) != "" {
			return OSSTrustNetworkValBReviewWorkflowStateActive
		}
		return OSSTrustNetworkValBReviewWorkflowStateBlocked
	case OSSTrustNetworkValBReviewStateCandidate:
		if strings.TrimSpace(model.ReviewerDecisionState) == OSSTrustNetworkValBReviewerDecisionStateNone ||
			strings.TrimSpace(model.ReviewerDecisionState) == OSSTrustNetworkValBReviewerDecisionStateNeedsMoreEvidence {
			return OSSTrustNetworkValBReviewWorkflowStatePartial
		}
		return OSSTrustNetworkValBReviewWorkflowStateBlocked
	case OSSTrustNetworkValBReviewStateInReview:
		if strings.TrimSpace(model.ReviewerDecisionState) == OSSTrustNetworkValBReviewerDecisionStateNone ||
			strings.TrimSpace(model.ReviewerDecisionState) == OSSTrustNetworkValBReviewerDecisionStateNeedsMoreEvidence {
			return OSSTrustNetworkValBReviewWorkflowStatePartial
		}
		return OSSTrustNetworkValBReviewWorkflowStateBlocked
	case OSSTrustNetworkValBReviewStateRejected:
		if model.RejectedPropagatedUsable || strings.TrimSpace(model.ReviewerDecisionState) != OSSTrustNetworkValBReviewerDecisionStateRejected || strings.TrimSpace(model.ReviewerRationale) == "" {
			return OSSTrustNetworkValBReviewWorkflowStateBlocked
		}
		return OSSTrustNetworkValBReviewWorkflowStatePartial
	case OSSTrustNetworkValBReviewStateSuperseded:
		if strings.TrimSpace(model.ReviewerDecisionState) != OSSTrustNetworkValBReviewerDecisionStateSuperseded ||
			strings.TrimSpace(model.ReviewerRationale) == "" ||
			strings.TrimSpace(model.ReplacementRef) == "" {
			return OSSTrustNetworkValBReviewWorkflowStateBlocked
		}
		return OSSTrustNetworkValBReviewWorkflowStatePartial
	case OSSTrustNetworkValBReviewStateRevoked:
		return OSSTrustNetworkValBReviewWorkflowStateBlocked
	default:
		return OSSTrustNetworkValBReviewWorkflowStateBlocked
	}
}

func EvaluateOSSTrustNetworkValBSharedVEXTriageState(model OSSTrustNetworkValBSharedVEXTriage) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.VEXOrTriageID,
		model.VulnerabilityOrAdvisoryRef,
		model.PackageOrProjectIdentity,
		model.AffectedReleaseScope,
		model.ExploitabilityContext,
		model.LocalApplicabilityNote,
		model.ReviewState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValBSharedVEXTriageStateIncomplete
	}
	if !ossTrustNetworkValBHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValBSharedVEXTriageStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valb-shared-vex-001") ||
		!containsTrimmedString(ossTrustNetworkValBSharedVEXStates(), model.ReviewState) ||
		!model.LocalApplicabilityBounded ||
		model.OverridesLocalEnterpriseApplicability ||
		model.CanonicalTruthClaim ||
		model.AutomaticGlobalTruthClaim {
		return OSSTrustNetworkValBSharedVEXTriageStateBlocked
	}
	switch strings.TrimSpace(model.ReviewState) {
	case OSSTrustNetworkValBSharedVEXStateReviewed:
		return OSSTrustNetworkValBSharedVEXTriageStateActive
	case OSSTrustNetworkValBSharedVEXStateCandidate:
		return OSSTrustNetworkValBSharedVEXTriageStatePartial
	case OSSTrustNetworkValBSharedVEXStateSuperseded:
		if strings.TrimSpace(model.SupersedesRef) == "" {
			return OSSTrustNetworkValBSharedVEXTriageStateBlocked
		}
		return OSSTrustNetworkValBSharedVEXTriageStatePartial
	case OSSTrustNetworkValBSharedVEXStateRejected,
		OSSTrustNetworkValBSharedVEXStateRevoked,
		OSSTrustNetworkValBSharedVEXStateUnsupported,
		OSSTrustNetworkValBSharedVEXStateUnknown:
		return OSSTrustNetworkValBSharedVEXTriageStateBlocked
	default:
		return OSSTrustNetworkValBSharedVEXTriageStateBlocked
	}
}

func EvaluateOSSTrustNetworkValBSourceWeightingState(model OSSTrustNetworkValBSourceWeighting) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.SourceClass,
		model.SourceWeightClass,
		model.FreshnessState,
		model.ReviewState,
		model.ProjectionDisclaimer,
	) || len(model.SupportedSourceClasses) == 0 || len(model.SupportedSourceWeightClasses) == 0 || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValBSourceWeightingStateIncomplete
	}
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	sourceClass := strings.TrimSpace(model.SourceClass)
	sourceWeightClass := strings.TrimSpace(model.SourceWeightClass)
	reviewState := strings.TrimSpace(model.ReviewState)
	if !ossTrustNetworkValBHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValBSourceWeightingStateUnknown
	}
	if !containsExactTrimmedStringSet(model.SupportedSourceClasses, ossTrustNetworkValBCandidateSourceClasses()...) ||
		!containsExactTrimmedStringSet(model.SupportedSourceWeightClasses, ossTrustNetworkValBSourceWeightClasses()...) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valb-source-weighting-001") ||
		!containsTrimmedString(ossTrustNetworkValBCandidateSourceClasses(), sourceClass) ||
		!containsTrimmedString(ossTrustNetworkValBSourceWeightClasses(), sourceWeightClass) ||
		!containsTrimmedString(ossTrustNetworkValBReviewStates(), reviewState) ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) ||
		model.UniversalTrustScoreClaim ||
		model.IntegrityScoreClaim ||
		model.BadgeScoreClaim {
		return OSSTrustNetworkValBSourceWeightingStateBlocked
	}
	if sourceClass == OSSTrustNetworkValBCandidateSourceClassAutomatedHeuristic &&
		sourceWeightClass == OSSTrustNetworkValBSourceWeightClassHigh &&
		reviewState == OSSTrustNetworkValBReviewStateReviewed &&
		model.AutomatedHeuristicStandalone {
		return OSSTrustNetworkValBSourceWeightingStateBlocked
	}
	if sourceClass == OSSTrustNetworkValBCandidateSourceClassCommunity &&
		reviewState == OSSTrustNetworkValBReviewStateReviewed &&
		model.CommunityInputWithoutReview {
		return OSSTrustNetworkValBSourceWeightingStateBlocked
	}
	switch reviewState {
	case OSSTrustNetworkValBReviewStateReviewed:
		return OSSTrustNetworkValBSourceWeightingStateActive
	case OSSTrustNetworkValBReviewStateCandidate, OSSTrustNetworkValBReviewStateInReview:
		return OSSTrustNetworkValBSourceWeightingStatePartial
	default:
		return OSSTrustNetworkValBSourceWeightingStateBlocked
	}
}

func EvaluateOSSTrustNetworkValBLocalApplicabilityState(model OSSTrustNetworkValBLocalApplicability) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.ApplicabilityState,
		model.LocalScope,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValBLocalApplicabilityStateIncomplete
	}
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	applicabilityState := strings.TrimSpace(model.ApplicabilityState)
	if !ossTrustNetworkValBHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValBLocalApplicabilityStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valb-local-applicability-001") ||
		!containsTrimmedString(ossTrustNetworkValBLocalApplicabilityStates(), applicabilityState) ||
		!containsTrimmedString(ossTrustNetworkVal0ApplicabilityScopes(), model.LocalScope) ||
		model.OverrideClaim ||
		model.SharedSignalHidden {
		return OSSTrustNetworkValBLocalApplicabilityStateBlocked
	}
	switch applicabilityState {
	case OSSTrustNetworkValBLocalApplicabilityStatusApplicable:
		if model.LocalEvidenceLinked {
			return OSSTrustNetworkValBLocalApplicabilityStateActive
		}
		return OSSTrustNetworkValBLocalApplicabilityStateBlocked
	case OSSTrustNetworkValBLocalApplicabilityStatusNotApplicable:
		if model.LocalEvidenceLinked && strings.TrimSpace(model.Rationale) != "" {
			return OSSTrustNetworkValBLocalApplicabilityStateActive
		}
		return OSSTrustNetworkValBLocalApplicabilityStateBlocked
	case OSSTrustNetworkValBLocalApplicabilityStatusNeedsLocalReview:
		return OSSTrustNetworkValBLocalApplicabilityStatePartial
	case OSSTrustNetworkValBLocalApplicabilityStatusUnknown, OSSTrustNetworkValBLocalApplicabilityStatusUnsupported:
		return OSSTrustNetworkValBLocalApplicabilityStateBlocked
	default:
		return OSSTrustNetworkValBLocalApplicabilityStateBlocked
	}
}

func EvaluateOSSTrustNetworkValBPropagationExchangeState(model OSSTrustNetworkValBPropagationExchange) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.PropagationState,
		model.ReviewState,
		model.SourceWeightClass,
		model.FreshnessState,
		model.LocalApplicabilityState,
		model.ApplicabilityScope,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValBPropagationExchangeStateIncomplete
	}
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	propagationState := strings.TrimSpace(model.PropagationState)
	reviewState := strings.TrimSpace(model.ReviewState)
	sourceWeightClass := strings.TrimSpace(model.SourceWeightClass)
	localApplicabilityState := strings.TrimSpace(model.LocalApplicabilityState)
	applicabilityScope := strings.TrimSpace(model.ApplicabilityScope)
	if !ossTrustNetworkValBHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValBPropagationExchangeStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valb-propagation-001") ||
		!containsTrimmedString(ossTrustNetworkValBPropagationStates(), propagationState) ||
		!containsTrimmedString(ossTrustNetworkValBReviewStates(), reviewState) ||
		!containsTrimmedString(ossTrustNetworkValBSourceWeightClasses(), sourceWeightClass) ||
		!containsTrimmedString(ossTrustNetworkValBLocalApplicabilityStates(), localApplicabilityState) ||
		!containsTrimmedString(ossTrustNetworkVal0ApplicabilityScopes(), applicabilityScope) ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) ||
		!model.SimilarityContextGating ||
		model.AutomaticGlobalSpread ||
		model.GlobalBlocklistClaim ||
		model.EnterpriseOverride {
		return OSSTrustNetworkValBPropagationExchangeStateBlocked
	}
	switch propagationState {
	case OSSTrustNetworkValBPropagationStateReviewedExchange:
		if reviewState == OSSTrustNetworkValBReviewStateReviewed &&
			localApplicabilityState == OSSTrustNetworkValBLocalApplicabilityStatusApplicable {
			return OSSTrustNetworkValBPropagationExchangeStateActive
		}
		return OSSTrustNetworkValBPropagationExchangeStateBlocked
	case OSSTrustNetworkValBPropagationStateCandidateExchange:
		if model.PresentedAsReviewed {
			return OSSTrustNetworkValBPropagationExchangeStateBlocked
		}
		return OSSTrustNetworkValBPropagationExchangeStatePartial
	case OSSTrustNetworkValBPropagationStateNotShared:
		return OSSTrustNetworkValBPropagationExchangeStatePartial
	case OSSTrustNetworkValBPropagationStateSuperseded:
		if strings.TrimSpace(model.ReplacementRef) == "" {
			return OSSTrustNetworkValBPropagationExchangeStateBlocked
		}
		return OSSTrustNetworkValBPropagationExchangeStatePartial
	case OSSTrustNetworkValBPropagationStateRejected,
		OSSTrustNetworkValBPropagationStateRevoked,
		OSSTrustNetworkValBPropagationStateUnsupported,
		OSSTrustNetworkValBPropagationStateUnknown:
		return OSSTrustNetworkValBPropagationExchangeStateBlocked
	default:
		return OSSTrustNetworkValBPropagationExchangeStateBlocked
	}
}

func EvaluateOSSTrustNetworkValBSupersessionRevocationState(model OSSTrustNetworkValBSupersessionRevocation) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.LifecycleState,
		model.PreviousSignalID,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValBSupersessionRevocationStateIncomplete
	}
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	lifecycleState := strings.TrimSpace(model.LifecycleState)
	if !ossTrustNetworkValBHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValBSupersessionRevocationStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valb-supersession-revocation-001") ||
		!containsTrimmedString(ossTrustNetworkValBLifecycleStates(), lifecycleState) ||
		!model.PreviousIdentityPreserved ||
		model.ReviewedExchangeStillActive {
		return OSSTrustNetworkValBSupersessionRevocationStateBlocked
	}
	switch lifecycleState {
	case OSSTrustNetworkValBLifecycleStateActive:
		return OSSTrustNetworkValBSupersessionRevocationStateActive
	case OSSTrustNetworkValBLifecycleStateSuperseded:
		if strings.TrimSpace(model.ReplacementRef) == "" || strings.TrimSpace(model.SupersessionReason) == "" {
			return OSSTrustNetworkValBSupersessionRevocationStateBlocked
		}
		return OSSTrustNetworkValBSupersessionRevocationStatePartial
	case OSSTrustNetworkValBLifecycleStateRevoked:
		if strings.TrimSpace(model.RevocationReason) == "" || strings.TrimSpace(model.RevocationTimestamp) == "" {
			return OSSTrustNetworkValBSupersessionRevocationStateBlocked
		}
		if _, valid := referenceArchitectureVal0ParseTimestamp(model.RevocationTimestamp); !valid {
			return OSSTrustNetworkValBSupersessionRevocationStateBlocked
		}
		return OSSTrustNetworkValBSupersessionRevocationStateBlocked
	default:
		return OSSTrustNetworkValBSupersessionRevocationStateBlocked
	}
}

func EvaluateOSSTrustNetworkValBReviewerAuditabilityState(model OSSTrustNetworkValBReviewerAuditability) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.ReviewerRoleClass,
		model.DecisionState,
		model.DecisionTimestamp,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValBReviewerAuditabilityStateIncomplete
	}
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	decisionState := strings.TrimSpace(model.DecisionState)
	if !ossTrustNetworkValBHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValBReviewerAuditabilityStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-valb-reviewer-auditability-001") ||
		!containsTrimmedString(ossTrustNetworkValBReviewerDecisionStates(), decisionState) ||
		model.LegalCertificationClaim ||
		model.RegulatorApprovalClaim ||
		model.OfficialOSSAuthorityClaim {
		return OSSTrustNetworkValBReviewerAuditabilityStateBlocked
	}
	if _, valid := referenceArchitectureVal0ParseTimestamp(model.DecisionTimestamp); !valid {
		return OSSTrustNetworkValBReviewerAuditabilityStateBlocked
	}
	switch decisionState {
	case OSSTrustNetworkValBReviewerDecisionStateNone:
		return OSSTrustNetworkValBReviewerAuditabilityStatePartial
	default:
		if strings.TrimSpace(model.Rationale) == "" {
			return OSSTrustNetworkValBReviewerAuditabilityStateBlocked
		}
		return OSSTrustNetworkValBReviewerAuditabilityStateActive
	}
}

func EvaluateOSSTrustNetworkValBNoOverclaimState(model OSSTrustNetworkValBNoOverclaim) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return OSSTrustNetworkValBNoOverclaimStateIncomplete
	}
	if !ossTrustNetworkValBHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValBNoOverclaimStateUnknown
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
		ossTrustNetworkValBContainsForbiddenClaim(model.ObservedClaims...) {
		return OSSTrustNetworkValBNoOverclaimStateBlocked
	}
	return OSSTrustNetworkValBNoOverclaimStateActive
}

func EvaluateOSSTrustNetworkValBState(model OSSTrustNetworkValBCore) string {
	if strings.TrimSpace(model.Point9State) == "" || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return OSSTrustNetworkValBStateIncomplete
	}
	if !ossTrustNetworkValBHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValBStateUnknown
	}
	if strings.TrimSpace(model.Point9State) != OSSTrustNetworkPoint9StateNotComplete ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValBProofSurfaceRefs()...) ||
		!OSSTrustNetworkValBProofEvidenceQualityValid(ossTrustNetworkValBEvidence(), model.EvidenceRefs) {
		return OSSTrustNetworkValBStateBlocked
	}
	states := []string{
		model.DependencyState,
		model.CandidateSignalIntakeState,
		model.ReviewWorkflowState,
		model.SharedVEXTriageState,
		model.SourceWeightingState,
		model.LocalApplicabilityState,
		model.PropagationExchangeState,
		model.SupersessionRevocationState,
		model.ReviewerAuditabilityState,
		model.NoOverclaimState,
	}
	allActive := true
	for _, state := range states {
		if strings.TrimSpace(state) == "" {
			return OSSTrustNetworkValBStateIncomplete
		}
		if !strings.HasSuffix(strings.TrimSpace(state), "_active") {
			allActive = false
		}
	}
	if allActive {
		return OSSTrustNetworkValBStateActive
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_blocked") {
			return OSSTrustNetworkValBStateBlocked
		}
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_incomplete") {
			return OSSTrustNetworkValBStateIncomplete
		}
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_unknown") {
			return OSSTrustNetworkValBStateUnknown
		}
	}
	return OSSTrustNetworkValBStatePartial
}

func EvaluateOSSTrustNetworkValBPointsState(currentState string) string {
	_ = currentState
	return OSSTrustNetworkPoint9StateNotComplete
}

func EvaluateOSSTrustNetworkValBProofsState(model OSSTrustNetworkValBCore, limitations []string) string {
	baseState := strings.TrimSpace(model.CurrentState)
	if baseState == "" {
		baseState = OSSTrustNetworkValBStateUnknown
	}
	if !ossTrustNetworkValBHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValBProofSurfaceRefs()...) ||
		!OSSTrustNetworkValBProofEvidenceQualityValid(ossTrustNetworkValBEvidence(), model.EvidenceRefs) ||
		len(limitations) == 0 ||
		strings.TrimSpace(model.Point9State) != OSSTrustNetworkPoint9StateNotComplete {
		if baseState == OSSTrustNetworkValBStateActive {
			return OSSTrustNetworkValBStatePartial
		}
		return baseState
	}
	return baseState
}

func ossTrustNetworkValBBlockingReasons(model OSSTrustNetworkValBCore) []string {
	reasons := []string{}
	if model.DependencyState != OSSTrustNetworkValBDependencyStateActive {
		reasons = append(reasons, "OSTN Val A dependency is not exact, active, and evidence-safe.")
	}
	if model.CandidateSignalIntakeState != OSSTrustNetworkValBCandidateSignalIntakeStateActive {
		reasons = append(reasons, "Candidate signal intake is not normalized, fresh, scoped, and bounded.")
	}
	if model.ReviewWorkflowState != OSSTrustNetworkValBReviewWorkflowStateActive {
		reasons = append(reasons, "Review workflow is not accepted, fresh, rationale-backed, and bounded.")
	}
	if model.SharedVEXTriageState != OSSTrustNetworkValBSharedVEXTriageStateActive {
		reasons = append(reasons, "Shared VEX / triage workflow is not reviewed, locally bounded, and evidence-linked.")
	}
	if model.SourceWeightingState != OSSTrustNetworkValBSourceWeightingStateActive {
		reasons = append(reasons, "Source weighting is not exact, bounded, and free of score overclaim.")
	}
	if model.LocalApplicabilityState != OSSTrustNetworkValBLocalApplicabilityStateActive {
		reasons = append(reasons, "Local applicability is not evidence-linked, scoped, and explicitly bounded.")
	}
	if model.PropagationExchangeState != OSSTrustNetworkValBPropagationExchangeStateActive {
		reasons = append(reasons, "Propagation exchange is not reviewed, fresh, context-gated, and locally bounded.")
	}
	if model.SupersessionRevocationState != OSSTrustNetworkValBSupersessionRevocationStateActive {
		reasons = append(reasons, "Supersession and revocation discipline is not preserving identity and fail-closed behavior.")
	}
	if model.ReviewerAuditabilityState != OSSTrustNetworkValBReviewerAuditabilityStateActive {
		reasons = append(reasons, "Reviewer decision auditability is not evidence-linked, rationalized, and bounded.")
	}
	if model.NoOverclaimState != OSSTrustNetworkValBNoOverclaimStateActive {
		reasons = append(reasons, "Val B no-overclaim and no-network-truth guard is not active.")
	}
	return developerEcosystemValECollectText(reasons)
}

func ComputeOSSTrustNetworkValBCore(model OSSTrustNetworkValBCore) OSSTrustNetworkValBCore {
	model.DependencyState = EvaluateOSSTrustNetworkValBDependencyState(model.Dependency)
	model.CandidateSignalIntakeState = EvaluateOSSTrustNetworkValBCandidateSignalIntakeState(model.CandidateSignalIntake)
	model.ReviewWorkflowState = EvaluateOSSTrustNetworkValBReviewWorkflowState(model.ReviewWorkflow)
	model.SharedVEXTriageState = EvaluateOSSTrustNetworkValBSharedVEXTriageState(model.SharedVEXTriage)
	model.SourceWeightingState = EvaluateOSSTrustNetworkValBSourceWeightingState(model.SourceWeighting)
	model.LocalApplicabilityState = EvaluateOSSTrustNetworkValBLocalApplicabilityState(model.LocalApplicability)
	model.PropagationExchangeState = EvaluateOSSTrustNetworkValBPropagationExchangeState(model.PropagationExchange)
	model.SupersessionRevocationState = EvaluateOSSTrustNetworkValBSupersessionRevocationState(model.SupersessionRevocation)
	model.ReviewerAuditabilityState = EvaluateOSSTrustNetworkValBReviewerAuditabilityState(model.ReviewerAuditability)
	model.NoOverclaimState = EvaluateOSSTrustNetworkValBNoOverclaimState(model.NoOverclaim)
	model.Point9State = EvaluateOSSTrustNetworkValBPointsState(model.CurrentState)
	model.CurrentState = EvaluateOSSTrustNetworkValBState(model)
	model.Point9State = EvaluateOSSTrustNetworkValBPointsState(model.CurrentState)
	model.BlockingReasons = ossTrustNetworkValBBlockingReasons(model)

	model.CandidateSignalIntake.CurrentState = model.CandidateSignalIntakeState
	model.ReviewWorkflow.CurrentState = model.ReviewWorkflowState
	model.SharedVEXTriage.CurrentState = model.SharedVEXTriageState
	model.SourceWeighting.CurrentState = model.SourceWeightingState
	model.LocalApplicability.CurrentState = model.LocalApplicabilityState
	model.PropagationExchange.CurrentState = model.PropagationExchangeState
	model.SupersessionRevocation.CurrentState = model.SupersessionRevocationState
	model.ReviewerAuditability.CurrentState = model.ReviewerAuditabilityState
	model.NoOverclaim.CurrentState = model.NoOverclaimState

	return model
}

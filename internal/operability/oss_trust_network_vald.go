package operability

import "strings"

const (
	OSSTrustNetworkValDStateActive     = "oss_trust_network_vald_active"
	OSSTrustNetworkValDStatePartial    = "oss_trust_network_vald_partial"
	OSSTrustNetworkValDStateIncomplete = "oss_trust_network_vald_incomplete"
	OSSTrustNetworkValDStateBlocked    = "oss_trust_network_vald_blocked"
	OSSTrustNetworkValDStateUnknown    = "oss_trust_network_vald_unknown"

	OSSTrustNetworkValDDependencyStateActive     = "oss_trust_network_vald_dependency_active"
	OSSTrustNetworkValDDependencyStatePartial    = "oss_trust_network_vald_dependency_partial"
	OSSTrustNetworkValDDependencyStateIncomplete = "oss_trust_network_vald_dependency_incomplete"
	OSSTrustNetworkValDDependencyStateBlocked    = "oss_trust_network_vald_dependency_blocked"
	OSSTrustNetworkValDDependencyStateUnknown    = "oss_trust_network_vald_dependency_unknown"

	OSSTrustNetworkValDSignalCorrectnessStateActive     = "oss_trust_network_vald_signal_correctness_active"
	OSSTrustNetworkValDSignalCorrectnessStatePartial    = "oss_trust_network_vald_signal_correctness_partial"
	OSSTrustNetworkValDSignalCorrectnessStateIncomplete = "oss_trust_network_vald_signal_correctness_incomplete"
	OSSTrustNetworkValDSignalCorrectnessStateBlocked    = "oss_trust_network_vald_signal_correctness_blocked"
	OSSTrustNetworkValDSignalCorrectnessStateUnknown    = "oss_trust_network_vald_signal_correctness_unknown"

	OSSTrustNetworkValDReleaseFoundationStateActive     = "oss_trust_network_vald_release_foundation_active"
	OSSTrustNetworkValDReleaseFoundationStatePartial    = "oss_trust_network_vald_release_foundation_partial"
	OSSTrustNetworkValDReleaseFoundationStateIncomplete = "oss_trust_network_vald_release_foundation_incomplete"
	OSSTrustNetworkValDReleaseFoundationStateBlocked    = "oss_trust_network_vald_release_foundation_blocked"
	OSSTrustNetworkValDReleaseFoundationStateUnknown    = "oss_trust_network_vald_release_foundation_unknown"

	OSSTrustNetworkValDReviewedIntelligenceStateActive     = "oss_trust_network_vald_reviewed_intelligence_active"
	OSSTrustNetworkValDReviewedIntelligenceStatePartial    = "oss_trust_network_vald_reviewed_intelligence_partial"
	OSSTrustNetworkValDReviewedIntelligenceStateIncomplete = "oss_trust_network_vald_reviewed_intelligence_incomplete"
	OSSTrustNetworkValDReviewedIntelligenceStateBlocked    = "oss_trust_network_vald_reviewed_intelligence_blocked"
	OSSTrustNetworkValDReviewedIntelligenceStateUnknown    = "oss_trust_network_vald_reviewed_intelligence_unknown"

	OSSTrustNetworkValDPropagationSafetyStateActive     = "oss_trust_network_vald_propagation_safety_active"
	OSSTrustNetworkValDPropagationSafetyStatePartial    = "oss_trust_network_vald_propagation_safety_partial"
	OSSTrustNetworkValDPropagationSafetyStateIncomplete = "oss_trust_network_vald_propagation_safety_incomplete"
	OSSTrustNetworkValDPropagationSafetyStateBlocked    = "oss_trust_network_vald_propagation_safety_blocked"
	OSSTrustNetworkValDPropagationSafetyStateUnknown    = "oss_trust_network_vald_propagation_safety_unknown"

	OSSTrustNetworkValDRemediationPRSafetyStateActive     = "oss_trust_network_vald_remediation_pr_safety_active"
	OSSTrustNetworkValDRemediationPRSafetyStatePartial    = "oss_trust_network_vald_remediation_pr_safety_partial"
	OSSTrustNetworkValDRemediationPRSafetyStateIncomplete = "oss_trust_network_vald_remediation_pr_safety_incomplete"
	OSSTrustNetworkValDRemediationPRSafetyStateBlocked    = "oss_trust_network_vald_remediation_pr_safety_blocked"
	OSSTrustNetworkValDRemediationPRSafetyStateUnknown    = "oss_trust_network_vald_remediation_pr_safety_unknown"

	OSSTrustNetworkValDEcosystemVisibilityConsistencyStateActive     = "oss_trust_network_vald_ecosystem_visibility_consistency_active"
	OSSTrustNetworkValDEcosystemVisibilityConsistencyStatePartial    = "oss_trust_network_vald_ecosystem_visibility_consistency_partial"
	OSSTrustNetworkValDEcosystemVisibilityConsistencyStateIncomplete = "oss_trust_network_vald_ecosystem_visibility_consistency_incomplete"
	OSSTrustNetworkValDEcosystemVisibilityConsistencyStateBlocked    = "oss_trust_network_vald_ecosystem_visibility_consistency_blocked"
	OSSTrustNetworkValDEcosystemVisibilityConsistencyStateUnknown    = "oss_trust_network_vald_ecosystem_visibility_consistency_unknown"

	OSSTrustNetworkValDEvidenceQualityStateActive     = "oss_trust_network_vald_evidence_quality_active"
	OSSTrustNetworkValDEvidenceQualityStatePartial    = "oss_trust_network_vald_evidence_quality_partial"
	OSSTrustNetworkValDEvidenceQualityStateIncomplete = "oss_trust_network_vald_evidence_quality_incomplete"
	OSSTrustNetworkValDEvidenceQualityStateBlocked    = "oss_trust_network_vald_evidence_quality_blocked"
	OSSTrustNetworkValDEvidenceQualityStateUnknown    = "oss_trust_network_vald_evidence_quality_unknown"

	OSSTrustNetworkValDNoOverclaimStateActive     = "oss_trust_network_vald_no_overclaim_active"
	OSSTrustNetworkValDNoOverclaimStatePartial    = "oss_trust_network_vald_no_overclaim_partial"
	OSSTrustNetworkValDNoOverclaimStateIncomplete = "oss_trust_network_vald_no_overclaim_incomplete"
	OSSTrustNetworkValDNoOverclaimStateBlocked    = "oss_trust_network_vald_no_overclaim_blocked"
	OSSTrustNetworkValDNoOverclaimStateUnknown    = "oss_trust_network_vald_no_overclaim_unknown"

	OSSTrustNetworkValDSignalLifecycleCandidate   = "candidate"
	OSSTrustNetworkValDSignalLifecycleReviewed    = "reviewed"
	OSSTrustNetworkValDSignalLifecycleRejected    = "rejected"
	OSSTrustNetworkValDSignalLifecycleSuperseded  = "superseded"
	OSSTrustNetworkValDSignalLifecycleRevoked     = "revoked"
	OSSTrustNetworkValDSignalLifecycleUnsupported = "unsupported"
	OSSTrustNetworkValDSignalLifecycleUnknown     = "unknown"

	OSSTrustNetworkValDSigningVerificationVerified    = "verified"
	OSSTrustNetworkValDProvenanceVerificationVerified = "verified"
	OSSTrustNetworkValDRegistryDescriptorModeOnly     = "descriptor_only"
)

type OSSTrustNetworkValDDependencySnapshot struct {
	ValCCurrentState               string   `json:"valc_current_state"`
	ValCPoint9State                string   `json:"valc_point_9_state"`
	ValCDependencyState            string   `json:"valc_dependency_state"`
	ValCTrustVisibilityState       string   `json:"valc_trust_visibility_state"`
	ValCPackageTrustStatusState    string   `json:"valc_package_trust_status_state"`
	ValCExportBoundaryState        string   `json:"valc_export_boundary_state"`
	ValCRemediationSuggestionState string   `json:"valc_remediation_suggestion_state"`
	ValCPRProposalState            string   `json:"valc_pr_proposal_state"`
	ValCLocalOverrideState         string   `json:"valc_local_override_state"`
	ValCRemediationSafetyState     string   `json:"valc_remediation_safety_state"`
	ValCEcosystemConsistencyState  string   `json:"valc_ecosystem_consistency_state"`
	ValCNoOverclaimState           string   `json:"valc_no_overclaim_state"`
	ValCProofSurfaceRefs           []string `json:"valc_proof_surface_refs,omitempty"`
	ValCEvidenceRefs               []string `json:"valc_evidence_refs,omitempty"`
	ValCProjectionDisclaimer       string   `json:"valc_projection_disclaimer"`
}

type OSSTrustNetworkValDSignalCorrectness struct {
	SignalID                         string   `json:"signal_id"`
	SignalLifecycleState             string   `json:"signal_lifecycle_state"`
	ReviewState                      string   `json:"review_state"`
	ReviewerDecisionState            string   `json:"reviewer_decision_state"`
	SourceClass                      string   `json:"source_class"`
	SourceWeightClass                string   `json:"source_weight_class"`
	FreshnessState                   string   `json:"freshness_state"`
	LocalApplicabilityStatus         string   `json:"local_applicability_status"`
	PropagationState                 string   `json:"propagation_state"`
	EvidenceRefs                     []string `json:"evidence_refs,omitempty"`
	Caveats                          []string `json:"caveats,omitempty"`
	ReplacementRef                   string   `json:"replacement_ref"`
	ProjectionDisclaimer             string   `json:"projection_disclaimer"`
	CandidateDisplayedAsReviewed     bool     `json:"candidate_displayed_as_reviewed"`
	RejectedUsableTrust              bool     `json:"rejected_usable_trust"`
	RevokedUsableTrust               bool     `json:"revoked_usable_trust"`
	SupersededUsableReviewedExchange bool     `json:"superseded_usable_reviewed_exchange"`
	CanonicalTruthClaim              bool     `json:"canonical_truth_claim"`
	EnterpriseApprovalAuthorityClaim bool     `json:"enterprise_approval_authority_claim"`
}

type OSSTrustNetworkValDReleaseFoundation struct {
	ReleaseTrustIntakeState                string   `json:"release_trust_intake_state"`
	SigningSignalState                     string   `json:"signing_signal_state"`
	SigningVerificationState               string   `json:"signing_verification_state"`
	SigningScoped                          bool     `json:"signing_scoped"`
	SigningEvidenceLinked                  bool     `json:"signing_evidence_linked"`
	SigningNonAuthoritative                bool     `json:"signing_non_authoritative"`
	MaintainerAttestationState             string   `json:"maintainer_attestation_state"`
	MaintainerIdentityBinding              bool     `json:"maintainer_identity_binding"`
	MaintainerKeyLinked                    bool     `json:"maintainer_key_linked"`
	MaintainerDelegationHandled            bool     `json:"maintainer_delegation_handled"`
	MaintainerRotationHandled              bool     `json:"maintainer_rotation_handled"`
	MaintainerRevocationHandled            bool     `json:"maintainer_revocation_handled"`
	MaintainerCompromiseHandled            bool     `json:"maintainer_compromise_handled"`
	ProvenanceMaterialState                string   `json:"provenance_material_state"`
	ProvenanceVerificationState            string   `json:"provenance_verification_state"`
	ProvenanceReleaseArtifactScoped        bool     `json:"provenance_release_artifact_scoped"`
	ProvenanceEvidenceLinked               bool     `json:"provenance_evidence_linked"`
	RegistryDescriptorState                string   `json:"registry_descriptor_state"`
	RegistryDescriptorMode                 string   `json:"registry_descriptor_mode"`
	LiveRegistryFetch                      bool     `json:"live_registry_fetch"`
	RegistryMetadataState                  string   `json:"registry_metadata_state"`
	RegistryMetadataNormalized             bool     `json:"registry_metadata_normalized"`
	RegistryMetadataFresh                  bool     `json:"registry_metadata_fresh"`
	RegistryMetadataScoped                 bool     `json:"registry_metadata_scoped"`
	RegistryMetadataEvidenceLinked         bool     `json:"registry_metadata_evidence_linked"`
	RegistryMetadataCreatesReviewedTrust   bool     `json:"registry_metadata_creates_reviewed_trust"`
	RegistryMetadataCreatesGlobalBlocklist bool     `json:"registry_metadata_creates_global_blocklist"`
	TypoSquattingWarningState              string   `json:"typo_squatting_warning_state"`
	TypoWarningCandidateBounded            bool     `json:"typo_warning_candidate_bounded"`
	TypoWarningAutoGlobalBlock             bool     `json:"typo_warning_auto_global_block"`
	TypoWarningCanonicalTruthClaim         bool     `json:"typo_warning_canonical_truth_claim"`
	DriftSignalState                       string   `json:"drift_signal_state"`
	DriftSignalEvidenceLinked              bool     `json:"drift_signal_evidence_linked"`
	DriftSignalSourceWeighted              bool     `json:"drift_signal_source_weighted"`
	DriftSignalScoped                      bool     `json:"drift_signal_scoped"`
	DriftSignalNonOverriding               bool     `json:"drift_signal_non_overriding"`
	DriftOverridesLocalEnterprise          bool     `json:"drift_overrides_local_enterprise"`
	EvidenceRefs                           []string `json:"evidence_refs,omitempty"`
	Caveats                                []string `json:"caveats,omitempty"`
	ProjectionDisclaimer                   string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValDReviewedIntelligence struct {
	CandidateIntakeState          string   `json:"candidate_intake_state"`
	CandidateIntakeNormalized     bool     `json:"candidate_intake_normalized"`
	ReviewWorkflowState           string   `json:"review_workflow_state"`
	ReviewState                   string   `json:"review_state"`
	ReviewerDecisionState         string   `json:"reviewer_decision_state"`
	ReviewerRationale             string   `json:"reviewer_rationale"`
	SharedVEXTriageState          string   `json:"shared_vex_triage_state"`
	SharedVEXState                string   `json:"shared_vex_state"`
	SharedVEXDisplayedAsReviewed  bool     `json:"shared_vex_displayed_as_reviewed"`
	SourceWeightingState          string   `json:"source_weighting_state"`
	SourceClass                   string   `json:"source_class"`
	SourceWeightClass             string   `json:"source_weight_class"`
	LocalApplicabilityGateState   string   `json:"local_applicability_gate_state"`
	LocalApplicabilityStatus      string   `json:"local_applicability_status"`
	DisplayedAsApplicable         bool     `json:"displayed_as_applicable"`
	ReviewerAuditabilityState     string   `json:"reviewer_auditability_state"`
	ReviewerRoleClass             string   `json:"reviewer_role_class"`
	ReviewerTimestamp             string   `json:"reviewer_timestamp"`
	ReviewerEvidenceLinked        bool     `json:"reviewer_evidence_linked"`
	FreshnessState                string   `json:"freshness_state"`
	EvidenceRefs                  []string `json:"evidence_refs,omitempty"`
	Caveats                       []string `json:"caveats,omitempty"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
	UniversalTrustScoreClaim      bool     `json:"universal_trust_score_claim"`
	IntegrityScoreClaim           bool     `json:"integrity_score_claim"`
	BadgeScoreClaim               bool     `json:"badge_score_claim"`
	RewritesCanonicalEvidence     bool     `json:"rewrites_canonical_evidence"`
	RewritesLocalEnterpriseResult bool     `json:"rewrites_local_enterprise_result"`
}

type OSSTrustNetworkValDPropagationSafety struct {
	PropagationState             string   `json:"propagation_state"`
	ReviewState                  string   `json:"review_state"`
	LocalApplicabilityStatus     string   `json:"local_applicability_status"`
	SourceWeightClass            string   `json:"source_weight_class"`
	FreshnessState               string   `json:"freshness_state"`
	ReplacementRef               string   `json:"replacement_ref"`
	EvidenceRefs                 []string `json:"evidence_refs,omitempty"`
	Caveats                      []string `json:"caveats,omitempty"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
	SimilarityContextGating      bool     `json:"similarity_context_gating"`
	CandidateDisplayedAsReviewed bool     `json:"candidate_displayed_as_reviewed"`
	AutomaticGlobalSpread        bool     `json:"automatic_global_spread"`
	GlobalBlocklistClaim         bool     `json:"global_blocklist_claim"`
	EnterpriseOverride           bool     `json:"enterprise_override"`
}

type OSSTrustNetworkValDRemediationPRSafety struct {
	SuggestionClass            string   `json:"suggestion_class"`
	ProposalState              string   `json:"proposal_state"`
	RiskClass                  string   `json:"risk_class"`
	CompatibilityNote          string   `json:"compatibility_note"`
	RiskNote                   string   `json:"risk_note"`
	RollbackNote               string   `json:"rollback_note"`
	TestValidationNote         string   `json:"test_validation_note"`
	LocalApplicabilityNote     string   `json:"local_applicability_note"`
	Rationale                  string   `json:"rationale"`
	EvidenceRefs               []string `json:"evidence_refs,omitempty"`
	Caveats                    []string `json:"caveats,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
	ReviewerRequired           bool     `json:"reviewer_required"`
	ProposalReviewerRequired   bool     `json:"proposal_reviewer_required"`
	ProposalNoAutomerge        bool     `json:"proposal_no_automerge"`
	ProposalNoHiddenMutation   bool     `json:"proposal_no_hidden_mutation"`
	ProposalAdvisoryOnly       bool     `json:"proposal_advisory_only"`
	ProposalBranchWrite        bool     `json:"proposal_branch_write"`
	ProposalNetworkAction      bool     `json:"proposal_network_action"`
	ProposalDependencyMutation bool     `json:"proposal_dependency_mutation"`
	ProposalPRCreation         bool     `json:"proposal_pr_creation"`
	ProposalAutoMerge          bool     `json:"proposal_auto_merge"`
	DependencyMutationAttempt  bool     `json:"dependency_mutation_attempt"`
	PolicyOverrideAttempt      bool     `json:"policy_override_attempt"`
	NoActionHidesRisk          bool     `json:"no_action_hides_risk"`
	HiddenMutationPath         bool     `json:"hidden_mutation_path"`
	ProductionApprovalClaim    bool     `json:"production_approval_claim"`
	DeploymentApprovalClaim    bool     `json:"deployment_approval_claim"`
}

type OSSTrustNetworkValDEcosystemVisibilityConsistency struct {
	VisibilityState                             string   `json:"visibility_state"`
	ReviewedSignalState                         string   `json:"reviewed_signal_state"`
	LocalApplicabilityState                     string   `json:"local_applicability_state"`
	SourceWeightingState                        string   `json:"source_weighting_state"`
	VisibilityFreshnessState                    string   `json:"visibility_freshness_state"`
	PackageStatusClass                          string   `json:"package_status_class"`
	PackageDisplayedAsReviewed                  bool     `json:"package_displayed_as_reviewed"`
	ExportClass                                 string   `json:"export_class"`
	LocalOverrideState                          string   `json:"local_override_state"`
	LocalOverrideScope                          string   `json:"local_override_scope"`
	LocalOverrideRationale                      string   `json:"local_override_rationale"`
	EvidenceRefs                                []string `json:"evidence_refs,omitempty"`
	Caveats                                     []string `json:"caveats,omitempty"`
	ProjectionDisclaimer                        string   `json:"projection_disclaimer"`
	CanonicalInternalExposure                   bool     `json:"canonical_internal_exposure"`
	CertificationClaim                          bool     `json:"certification_claim"`
	ApprovalClaim                               bool     `json:"approval_claim"`
	CandidatePromotedToReviewed                 bool     `json:"candidate_promoted_to_reviewed"`
	RejectedPromotedToActive                    bool     `json:"rejected_promoted_to_active"`
	RevokedPromotedToActive                     bool     `json:"revoked_promoted_to_active"`
	UnknownPromotedToActive                     bool     `json:"unknown_promoted_to_active"`
	LocalOnlyBoundary                           bool     `json:"local_only_boundary"`
	RewriteCanonicalEvidence                    bool     `json:"rewrite_canonical_evidence"`
	SilentlySuppressReviewedNetworkIntelligence bool     `json:"silently_suppress_reviewed_network_intelligence"`
	SharedSignalOverridesLocalDecision          bool     `json:"shared_signal_overrides_local_decision"`
	ReviewedExchangePresentedActive             bool     `json:"reviewed_exchange_presented_active"`
}

type OSSTrustNetworkValDEvidenceQuality struct {
	Evidence                       []ReferenceArchitectureEvidenceReference `json:"evidence,omitempty"`
	ProofSurfaceRefs               []string                                 `json:"proof_surface_refs,omitempty"`
	EvidenceRefs                   []string                                 `json:"evidence_refs,omitempty"`
	DependencyProofSurfaceRefs     []string                                 `json:"dependency_proof_surface_refs,omitempty"`
	DependencyEvidenceRefs         []string                                 `json:"dependency_evidence_refs,omitempty"`
	DependencyEvidence             []ReferenceArchitectureEvidenceReference `json:"dependency_evidence,omitempty"`
	DependencyProjectionDisclaimer string                                   `json:"dependency_projection_disclaimer"`
	ProjectionDisclaimer           string                                   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValDNoOverclaim struct {
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
	PublicBadgeClaim           bool     `json:"public_badge_claim"`
	OfficialOSSAuthorityClaim  bool     `json:"official_oss_authority_claim"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValDCore struct {
	CurrentState                        string                                            `json:"current_state"`
	Point9State                         string                                            `json:"point_9_state"`
	DependencyState                     string                                            `json:"dependency_state"`
	SignalCorrectnessState              string                                            `json:"signal_correctness_state"`
	ReleaseFoundationState              string                                            `json:"release_foundation_state"`
	ReviewedIntelligenceState           string                                            `json:"reviewed_intelligence_state"`
	PropagationSafetyState              string                                            `json:"propagation_safety_state"`
	RemediationPRSafetyState            string                                            `json:"remediation_pr_safety_state"`
	EcosystemVisibilityConsistencyState string                                            `json:"ecosystem_visibility_consistency_state"`
	EvidenceQualityState                string                                            `json:"evidence_quality_state"`
	NoOverclaimState                    string                                            `json:"no_overclaim_state"`
	Dependency                          OSSTrustNetworkValDDependencySnapshot             `json:"dependency"`
	SignalCorrectness                   OSSTrustNetworkValDSignalCorrectness              `json:"signal_correctness"`
	ReleaseFoundation                   OSSTrustNetworkValDReleaseFoundation              `json:"release_foundation"`
	ReviewedIntelligence                OSSTrustNetworkValDReviewedIntelligence           `json:"reviewed_intelligence"`
	PropagationSafety                   OSSTrustNetworkValDPropagationSafety              `json:"propagation_safety"`
	RemediationPRSafety                 OSSTrustNetworkValDRemediationPRSafety            `json:"remediation_pr_safety"`
	EcosystemVisibilityConsistency      OSSTrustNetworkValDEcosystemVisibilityConsistency `json:"ecosystem_visibility_consistency"`
	EvidenceQuality                     OSSTrustNetworkValDEvidenceQuality                `json:"evidence_quality"`
	NoOverclaim                         OSSTrustNetworkValDNoOverclaim                    `json:"no_overclaim"`
	ProofSurfaceRefs                    []string                                          `json:"proof_surface_refs,omitempty"`
	EvidenceRefs                        []string                                          `json:"evidence_refs,omitempty"`
	BlockingReasons                     []string                                          `json:"blocking_reasons,omitempty"`
	WhyPoint9NotComplete                []string                                          `json:"why_point_9_not_complete,omitempty"`
	FinalReadinessSummary               []string                                          `json:"final_readiness_summary,omitempty"`
	ProjectionDisclaimer                string                                            `json:"projection_disclaimer"`
}

type ossTrustNetworkValDExpectedEvidenceMetadata struct {
	EvidenceType string
	Source       string
	Scope        string
}

func ossTrustNetworkValDProjectionDisclaimer() string {
	return "projection_only not_canonical_truth oss_trust_network_vald final_readiness_gate not_integrated_closure"
}

func ossTrustNetworkValDHasProjectionDisclaimer(value string) bool {
	normalized := strings.ToLower(strings.TrimSpace(value))
	return strings.Contains(normalized, "projection_only") &&
		strings.Contains(normalized, "not_canonical_truth") &&
		strings.Contains(normalized, "oss_trust_network_vald")
}

func ossTrustNetworkValDSignalLifecycleStates() []string {
	return []string{
		OSSTrustNetworkValDSignalLifecycleCandidate,
		OSSTrustNetworkValDSignalLifecycleReviewed,
		OSSTrustNetworkValDSignalLifecycleRejected,
		OSSTrustNetworkValDSignalLifecycleSuperseded,
		OSSTrustNetworkValDSignalLifecycleRevoked,
		OSSTrustNetworkValDSignalLifecycleUnsupported,
		OSSTrustNetworkValDSignalLifecycleUnknown,
	}
}

func OSSTrustNetworkValDProofSurfaceRefs() []string {
	return []string{
		"/v1/oss-trust-network/vald/status",
		"/v1/oss-trust-network/vald/proofs",
	}
}

func OSSTrustNetworkValDProofEvidenceRefs() []string {
	return []string{
		"evidence:ostn-vald-dependency-001",
		"evidence:ostn-vald-signal-correctness-001",
		"evidence:ostn-vald-release-foundation-001",
		"evidence:ostn-vald-reviewed-intelligence-001",
		"evidence:ostn-vald-propagation-safety-001",
		"evidence:ostn-vald-remediation-pr-safety-001",
		"evidence:ostn-vald-ecosystem-consistency-001",
		"evidence:ostn-vald-evidence-quality-001",
		"evidence:ostn-vald-no-overclaim-001",
		"evidence:ostn-vald-point9-governance-001",
	}
}

func ossTrustNetworkValDExpectedEvidenceMetadataByID() map[string]ossTrustNetworkValDExpectedEvidenceMetadata {
	return map[string]ossTrustNetworkValDExpectedEvidenceMetadata{
		"evidence:ostn-vald-dependency-001":            {EvidenceType: "dependency_state", Source: "oss-trust-network/vald/dependency", Scope: "valc_dependency"},
		"evidence:ostn-vald-signal-correctness-001":    {EvidenceType: "signal_correctness", Source: "oss-trust-network/vald/signal-correctness", Scope: "signal_correctness_gate"},
		"evidence:ostn-vald-release-foundation-001":    {EvidenceType: "release_foundation", Source: "oss-trust-network/vald/release-foundation", Scope: "release_foundation_gate"},
		"evidence:ostn-vald-reviewed-intelligence-001": {EvidenceType: "reviewed_intelligence", Source: "oss-trust-network/vald/reviewed-intelligence", Scope: "reviewed_intelligence_gate"},
		"evidence:ostn-vald-propagation-safety-001":    {EvidenceType: "propagation_safety", Source: "oss-trust-network/vald/propagation-safety", Scope: "propagation_safety_gate"},
		"evidence:ostn-vald-remediation-pr-safety-001": {EvidenceType: "remediation_pr_safety", Source: "oss-trust-network/vald/remediation-pr-safety", Scope: "remediation_pr_safety_gate"},
		"evidence:ostn-vald-ecosystem-consistency-001": {EvidenceType: "ecosystem_visibility_consistency", Source: "oss-trust-network/vald/ecosystem-consistency", Scope: "ecosystem_visibility_consistency_gate"},
		"evidence:ostn-vald-evidence-quality-001":      {EvidenceType: "evidence_quality", Source: "oss-trust-network/vald/evidence-quality", Scope: "evidence_quality_gate"},
		"evidence:ostn-vald-no-overclaim-001":          {EvidenceType: "no_overclaim", Source: "oss-trust-network/vald/no-overclaim", Scope: "no_overclaim_discipline"},
		"evidence:ostn-vald-point9-governance-001":     {EvidenceType: "state_governance", Source: "oss-trust-network/vald/point9-governance", Scope: "point9_governance"},
	}
}

func ossTrustNetworkValDEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:ostn-vald-dependency-001", EvidenceType: "dependency_state", Source: "oss-trust-network/vald/dependency", Timestamp: "2026-04-30T09:00:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "valc_dependency", Caveats: []string{"Val D depends on exact and active OSTN Val C only."}},
		{EvidenceID: "evidence:ostn-vald-signal-correctness-001", EvidenceType: "signal_correctness", Source: "oss-trust-network/vald/signal-correctness", Timestamp: "2026-04-30T09:01:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "signal_correctness_gate", Caveats: []string{"Signal correctness remains bounded and not canonical truth or enterprise approval authority."}},
		{EvidenceID: "evidence:ostn-vald-release-foundation-001", EvidenceType: "release_foundation", Source: "oss-trust-network/vald/release-foundation", Timestamp: "2026-04-30T09:02:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "release_foundation_gate", Caveats: []string{"Release, provenance, maintainer, registry, typo-warning, and drift gates remain readiness checks and not integrated closure."}},
		{EvidenceID: "evidence:ostn-vald-reviewed-intelligence-001", EvidenceType: "reviewed_intelligence", Source: "oss-trust-network/vald/reviewed-intelligence", Timestamp: "2026-04-30T09:03:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "reviewed_intelligence_gate", Caveats: []string{"Reviewed intelligence remains source-weighted, locally bounded, and non-canonical."}},
		{EvidenceID: "evidence:ostn-vald-propagation-safety-001", EvidenceType: "propagation_safety", Source: "oss-trust-network/vald/propagation-safety", Timestamp: "2026-04-30T09:04:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "propagation_safety_gate", Caveats: []string{"Network exchange remains bounded and cannot create global truth or global blocklists."}},
		{EvidenceID: "evidence:ostn-vald-remediation-pr-safety-001", EvidenceType: "remediation_pr_safety", Source: "oss-trust-network/vald/remediation-pr-safety", Timestamp: "2026-04-30T09:05:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "remediation_pr_safety_gate", Caveats: []string{"Remediation and proposal outputs remain advisory-only and free of hidden mutation, branch write, or PR creation."}},
		{EvidenceID: "evidence:ostn-vald-ecosystem-consistency-001", EvidenceType: "ecosystem_visibility_consistency", Source: "oss-trust-network/vald/ecosystem-consistency", Timestamp: "2026-04-30T09:06:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "ecosystem_visibility_consistency_gate", Caveats: []string{"Visibility, export, and local override surfaces must remain internally consistent and locally bounded."}},
		{EvidenceID: "evidence:ostn-vald-evidence-quality-001", EvidenceType: "evidence_quality", Source: "oss-trust-network/vald/evidence-quality", Timestamp: "2026-04-30T09:07:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "evidence_quality_gate", Caveats: []string{"Val D evidence must remain exact, fresh, related, and scope-correct."}},
		{EvidenceID: "evidence:ostn-vald-no-overclaim-001", EvidenceType: "no_overclaim", Source: "oss-trust-network/vald/no-overclaim", Timestamp: "2026-04-30T09:08:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim_discipline", Caveats: []string{"Val D cannot create promotional badge surfaces, official authority claims, or forbidden point pass semantics."}},
		{EvidenceID: "evidence:ostn-vald-point9-governance-001", EvidenceType: "state_governance", Source: "oss-trust-network/vald/point9-governance", Timestamp: "2026-04-30T09:09:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point9_governance", Caveats: []string{"Val D is a final readiness gate only; Val E remains required for integrated closure and any final pass semantics."}},
	}
}

func ossTrustNetworkValDCopyEvidence(items []ReferenceArchitectureEvidenceReference) []ReferenceArchitectureEvidenceReference {
	cloned := make([]ReferenceArchitectureEvidenceReference, 0, len(items))
	for _, item := range items {
		cloned = append(cloned, ReferenceArchitectureEvidenceReference{
			EvidenceID:     item.EvidenceID,
			EvidenceType:   item.EvidenceType,
			Source:         item.Source,
			Timestamp:      item.Timestamp,
			FreshnessState: item.FreshnessState,
			Scope:          item.Scope,
			Caveats:        append([]string{}, item.Caveats...),
		})
	}
	return cloned
}

func OSSTrustNetworkValDProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, refs []string) bool {
	if !containsExactTrimmedStringSet(refs, OSSTrustNetworkValDProofEvidenceRefs()...) {
		return false
	}
	expected := ossTrustNetworkValDExpectedEvidenceMetadataByID()
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
		expectedMetadata, exists := expected[id]
		if id == "" || !exists {
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

func ossTrustNetworkValDContainsForbiddenClaim(values ...string) bool {
	allowed := map[string]struct{}{
		"final ostn readiness gate":       {},
		"bounded oss trust signal":        {},
		"reviewed oss trust signal":       {},
		"candidate oss trust signal":      {},
		"source-weighted reviewed signal": {},
		"bounded reviewed exchange":       {},
		"local applicability context":     {},
		"package trust visibility":        {},
		"advisory remediation suggestion": {},
		"pr proposal descriptor":          {},
		"reviewer-required proposal":      {},
		"no hidden mutation path":         {},
		"not canonical truth":             {},
		"not formal certification":        {},
		"not production approval":         {},
		"not integrated closure":          {},
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
		"auto-remediated",
		"auto-merged",
		"production autopatch",
		"public badge",
		"official oss authority",
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

func OSSTrustNetworkValDDependencySnapshotModel() OSSTrustNetworkValDDependencySnapshot {
	model := ComputeOSSTrustNetworkValCCore(OSSTrustNetworkValCCoreModel())
	return OSSTrustNetworkValDDependencySnapshot{
		ValCCurrentState:               model.CurrentState,
		ValCPoint9State:                model.Point9State,
		ValCDependencyState:            model.DependencyState,
		ValCTrustVisibilityState:       model.TrustVisibilityState,
		ValCPackageTrustStatusState:    model.PackageTrustStatusState,
		ValCExportBoundaryState:        model.ExportBoundaryState,
		ValCRemediationSuggestionState: model.RemediationSuggestionState,
		ValCPRProposalState:            model.PRProposalState,
		ValCLocalOverrideState:         model.LocalOverrideState,
		ValCRemediationSafetyState:     model.RemediationSafetyState,
		ValCEcosystemConsistencyState:  model.EcosystemConsistencyState,
		ValCNoOverclaimState:           model.NoOverclaimState,
		ValCProofSurfaceRefs:           append([]string{}, model.ProofSurfaceRefs...),
		ValCEvidenceRefs:               append([]string{}, model.EvidenceRefs...),
		ValCProjectionDisclaimer:       model.ProjectionDisclaimer,
	}
}

func OSSTrustNetworkValDSignalCorrectnessModel() OSSTrustNetworkValDSignalCorrectness {
	return OSSTrustNetworkValDSignalCorrectness{
		SignalID:                 "ostn-vald-signal-correctness-001",
		SignalLifecycleState:     OSSTrustNetworkValDSignalLifecycleReviewed,
		ReviewState:              OSSTrustNetworkValBReviewStateReviewed,
		ReviewerDecisionState:    OSSTrustNetworkValBReviewerDecisionStateAccepted,
		SourceClass:              OSSTrustNetworkValBCandidateSourceClassMaintainer,
		SourceWeightClass:        OSSTrustNetworkValBSourceWeightClassBounded,
		FreshnessState:           IntelligenceCalibrationFreshnessFresh,
		LocalApplicabilityStatus: OSSTrustNetworkValBLocalApplicabilityStatusApplicable,
		PropagationState:         OSSTrustNetworkValBPropagationStateReviewedExchange,
		EvidenceRefs:             []string{"evidence:ostn-vald-signal-correctness-001"},
		Caveats:                  []string{"signal correctness remains bounded and non-canonical"},
		ProjectionDisclaimer:     ossTrustNetworkValDProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValDReleaseFoundationModel() OSSTrustNetworkValDReleaseFoundation {
	valA := ComputeOSSTrustNetworkValACore(OSSTrustNetworkValACoreModel())
	return OSSTrustNetworkValDReleaseFoundation{
		ReleaseTrustIntakeState:         valA.ReleaseTrustIntakeState,
		SigningSignalState:              valA.SigningSignalState,
		SigningVerificationState:        OSSTrustNetworkValDSigningVerificationVerified,
		SigningScoped:                   true,
		SigningEvidenceLinked:           true,
		SigningNonAuthoritative:         true,
		MaintainerAttestationState:      valA.MaintainerAttestationState,
		MaintainerIdentityBinding:       true,
		MaintainerKeyLinked:             true,
		MaintainerDelegationHandled:     true,
		MaintainerRotationHandled:       true,
		MaintainerRevocationHandled:     true,
		MaintainerCompromiseHandled:     true,
		ProvenanceMaterialState:         valA.ProvenanceMaterialState,
		ProvenanceVerificationState:     OSSTrustNetworkValDProvenanceVerificationVerified,
		ProvenanceReleaseArtifactScoped: true,
		ProvenanceEvidenceLinked:        true,
		RegistryDescriptorState:         valA.RegistryDescriptorState,
		RegistryDescriptorMode:          OSSTrustNetworkValDRegistryDescriptorModeOnly,
		RegistryMetadataState:           valA.RegistryMetadataState,
		RegistryMetadataNormalized:      true,
		RegistryMetadataFresh:           true,
		RegistryMetadataScoped:          true,
		RegistryMetadataEvidenceLinked:  true,
		TypoSquattingWarningState:       valA.TypoSquattingWarningState,
		TypoWarningCandidateBounded:     true,
		DriftSignalState:                valA.DriftSignalState,
		DriftSignalEvidenceLinked:       true,
		DriftSignalSourceWeighted:       true,
		DriftSignalScoped:               true,
		DriftSignalNonOverriding:        true,
		EvidenceRefs:                    []string{"evidence:ostn-vald-release-foundation-001"},
		Caveats:                         []string{"release foundation gate remains bounded readiness only"},
		ProjectionDisclaimer:            ossTrustNetworkValDProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValDReviewedIntelligenceModel() OSSTrustNetworkValDReviewedIntelligence {
	return OSSTrustNetworkValDReviewedIntelligence{
		CandidateIntakeState:        OSSTrustNetworkValBCandidateSignalIntakeStateActive,
		CandidateIntakeNormalized:   true,
		ReviewWorkflowState:         OSSTrustNetworkValBReviewWorkflowStateActive,
		ReviewState:                 OSSTrustNetworkValBReviewStateReviewed,
		ReviewerDecisionState:       OSSTrustNetworkValBReviewerDecisionStateAccepted,
		ReviewerRationale:           "reviewed signal accepted with explicit rationale",
		SharedVEXTriageState:        OSSTrustNetworkValBSharedVEXTriageStateActive,
		SharedVEXState:              OSSTrustNetworkValBSharedVEXStateReviewed,
		SourceWeightingState:        OSSTrustNetworkValBSourceWeightingStateActive,
		SourceClass:                 OSSTrustNetworkValBCandidateSourceClassVerifier,
		SourceWeightClass:           OSSTrustNetworkValBSourceWeightClassBounded,
		LocalApplicabilityGateState: OSSTrustNetworkValBLocalApplicabilityStateActive,
		LocalApplicabilityStatus:    OSSTrustNetworkValBLocalApplicabilityStatusApplicable,
		DisplayedAsApplicable:       true,
		ReviewerAuditabilityState:   OSSTrustNetworkValBReviewerAuditabilityStateActive,
		ReviewerRoleClass:           "oss_trust_reviewer",
		ReviewerTimestamp:           "2026-04-30T09:03:30Z",
		ReviewerEvidenceLinked:      true,
		FreshnessState:              IntelligenceCalibrationFreshnessFresh,
		EvidenceRefs:                []string{"evidence:ostn-vald-reviewed-intelligence-001"},
		Caveats:                     []string{"reviewed intelligence remains advisory and cannot rewrite canonical evidence"},
		ProjectionDisclaimer:        ossTrustNetworkValDProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValDPropagationSafetyModel() OSSTrustNetworkValDPropagationSafety {
	return OSSTrustNetworkValDPropagationSafety{
		PropagationState:         OSSTrustNetworkValBPropagationStateReviewedExchange,
		ReviewState:              OSSTrustNetworkValBReviewStateReviewed,
		LocalApplicabilityStatus: OSSTrustNetworkValBLocalApplicabilityStatusApplicable,
		SourceWeightClass:        OSSTrustNetworkValBSourceWeightClassBounded,
		FreshnessState:           IntelligenceCalibrationFreshnessFresh,
		SimilarityContextGating:  true,
		EvidenceRefs:             []string{"evidence:ostn-vald-propagation-safety-001"},
		Caveats:                  []string{"network exchange remains bounded and similarity-gated"},
		ProjectionDisclaimer:     ossTrustNetworkValDProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValDRemediationPRSafetyModel() OSSTrustNetworkValDRemediationPRSafety {
	return OSSTrustNetworkValDRemediationPRSafety{
		SuggestionClass:          OSSTrustNetworkValCSuggestionClassVersionUpgrade,
		ProposalState:            OSSTrustNetworkValCProposalStateProposalReady,
		RiskClass:                OSSTrustNetworkValCRiskClassMedium,
		CompatibilityNote:        "compatibility remains reviewer-validated and bounded",
		RiskNote:                 "bounded rollout risk remains visible",
		RollbackNote:             "rollback to refs/tags/v1.2.3 if validation fails",
		TestValidationNote:       "run unit, integration, and release validation before adoption",
		LocalApplicabilityNote:   "apply only after enterprise-local applicability review",
		Rationale:                "reviewed signal supports bounded upgrade guidance",
		EvidenceRefs:             []string{"evidence:ostn-vald-remediation-pr-safety-001"},
		Caveats:                  []string{"remediation and proposal outputs remain advisory only"},
		ProjectionDisclaimer:     ossTrustNetworkValDProjectionDisclaimer(),
		ReviewerRequired:         true,
		ProposalReviewerRequired: true,
		ProposalNoAutomerge:      true,
		ProposalNoHiddenMutation: true,
		ProposalAdvisoryOnly:     true,
	}
}

func OSSTrustNetworkValDEcosystemVisibilityConsistencyModel() OSSTrustNetworkValDEcosystemVisibilityConsistency {
	return OSSTrustNetworkValDEcosystemVisibilityConsistency{
		VisibilityState:                 OSSTrustNetworkValCVisibilityVisible,
		ReviewedSignalState:             OSSTrustNetworkValBReviewWorkflowStateActive,
		LocalApplicabilityState:         OSSTrustNetworkValBLocalApplicabilityStateActive,
		SourceWeightingState:            OSSTrustNetworkValBSourceWeightingStateActive,
		VisibilityFreshnessState:        IntelligenceCalibrationFreshnessFresh,
		PackageStatusClass:              OSSTrustNetworkValCPackageStatusReviewedSignalAvailable,
		PackageDisplayedAsReviewed:      true,
		ExportClass:                     OSSTrustNetworkValCExportClassEnterpriseCustomerView,
		LocalOverrideState:              OSSTrustNetworkValCOverrideStateNoOverride,
		LocalOverrideScope:              OSSTrustNetworkApplicabilityEnterpriseLocal,
		LocalOverrideRationale:          "no local override recorded for this release context",
		LocalOnlyBoundary:               true,
		ReviewedExchangePresentedActive: true,
		EvidenceRefs:                    []string{"evidence:ostn-vald-ecosystem-consistency-001"},
		Caveats:                         []string{"ecosystem visibility remains exact, local, and non-promotional"},
		ProjectionDisclaimer:            ossTrustNetworkValDProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValDEvidenceQualityModel() OSSTrustNetworkValDEvidenceQuality {
	valC := ComputeOSSTrustNetworkValCCore(OSSTrustNetworkValCCoreModel())
	return OSSTrustNetworkValDEvidenceQuality{
		Evidence:                       ossTrustNetworkValDCopyEvidence(ossTrustNetworkValDEvidence()),
		ProofSurfaceRefs:               OSSTrustNetworkValDProofSurfaceRefs(),
		EvidenceRefs:                   OSSTrustNetworkValDProofEvidenceRefs(),
		DependencyProofSurfaceRefs:     append([]string{}, valC.ProofSurfaceRefs...),
		DependencyEvidenceRefs:         append([]string{}, valC.EvidenceRefs...),
		DependencyEvidence:             ossTrustNetworkValDCopyEvidence(ossTrustNetworkValCEvidence()),
		DependencyProjectionDisclaimer: valC.ProjectionDisclaimer,
		ProjectionDisclaimer:           ossTrustNetworkValDProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValDNoOverclaimModel() OSSTrustNetworkValDNoOverclaim {
	return OSSTrustNetworkValDNoOverclaim{
		DisciplineID:         "oss_no_overclaim_vald",
		Version:              "v0",
		ObservedClaims:       []string{"final ostn readiness gate", "no hidden mutation path", "not integrated closure"},
		ProjectionDisclaimer: ossTrustNetworkValDProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValDCoreModel() OSSTrustNetworkValDCore {
	return OSSTrustNetworkValDCore{
		CurrentState:                   OSSTrustNetworkValDStateActive,
		Point9State:                    OSSTrustNetworkPoint9StateNotComplete,
		Dependency:                     OSSTrustNetworkValDDependencySnapshotModel(),
		SignalCorrectness:              OSSTrustNetworkValDSignalCorrectnessModel(),
		ReleaseFoundation:              OSSTrustNetworkValDReleaseFoundationModel(),
		ReviewedIntelligence:           OSSTrustNetworkValDReviewedIntelligenceModel(),
		PropagationSafety:              OSSTrustNetworkValDPropagationSafetyModel(),
		RemediationPRSafety:            OSSTrustNetworkValDRemediationPRSafetyModel(),
		EcosystemVisibilityConsistency: OSSTrustNetworkValDEcosystemVisibilityConsistencyModel(),
		EvidenceQuality:                OSSTrustNetworkValDEvidenceQualityModel(),
		NoOverclaim:                    OSSTrustNetworkValDNoOverclaimModel(),
		ProofSurfaceRefs:               OSSTrustNetworkValDProofSurfaceRefs(),
		EvidenceRefs:                   OSSTrustNetworkValDProofEvidenceRefs(),
		WhyPoint9NotComplete: []string{
			"Val D is a final OSTN readiness gate only and cannot complete Točka 9.",
			"Integrated closure and any final pass semantics remain reserved for Val E.",
			"Val D may report readiness over bounded OSTN signals, visibility, and remediation safety, but it does not create canonical truth, approval authority, or integrated closure.",
		},
		FinalReadinessSummary: []string{
			"Val D checks exact Val C dependency, signal correctness, release foundation, reviewed intelligence, propagation safety, remediation safety, ecosystem visibility consistency, evidence quality, and no-overclaim discipline.",
			"Val D can report final readiness for Val E handoff, but not PASS or integrated closure.",
		},
		ProjectionDisclaimer: ossTrustNetworkValDProjectionDisclaimer(),
	}
}

func EvaluateOSSTrustNetworkValDDependencyState(model OSSTrustNetworkValDDependencySnapshot) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ValCCurrentState,
		model.ValCPoint9State,
		model.ValCDependencyState,
		model.ValCTrustVisibilityState,
		model.ValCPackageTrustStatusState,
		model.ValCExportBoundaryState,
		model.ValCRemediationSuggestionState,
		model.ValCPRProposalState,
		model.ValCLocalOverrideState,
		model.ValCRemediationSafetyState,
		model.ValCEcosystemConsistencyState,
		model.ValCNoOverclaimState,
		model.ValCProjectionDisclaimer,
	) || len(model.ValCProofSurfaceRefs) == 0 || len(model.ValCEvidenceRefs) == 0 {
		return OSSTrustNetworkValDDependencyStateIncomplete
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(model.ValCProjectionDisclaimer) {
		return OSSTrustNetworkValDDependencyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.ValCProofSurfaceRefs, OSSTrustNetworkValCProofSurfaceRefs()...) ||
		!containsExactTrimmedStringSet(model.ValCEvidenceRefs, OSSTrustNetworkValCProofEvidenceRefs()...) ||
		!OSSTrustNetworkValCProofEvidenceQualityValid(ossTrustNetworkValCEvidence(), model.ValCEvidenceRefs) {
		return OSSTrustNetworkValDDependencyStateBlocked
	}
	if strings.TrimSpace(model.ValCCurrentState) != OSSTrustNetworkValCStateActive ||
		strings.TrimSpace(model.ValCPoint9State) != OSSTrustNetworkPoint9StateNotComplete ||
		strings.TrimSpace(model.ValCDependencyState) != OSSTrustNetworkValCDependencyStateActive ||
		strings.TrimSpace(model.ValCTrustVisibilityState) != OSSTrustNetworkValCTrustVisibilityStateActive ||
		strings.TrimSpace(model.ValCPackageTrustStatusState) != OSSTrustNetworkValCPackageTrustStatusStateActive ||
		strings.TrimSpace(model.ValCExportBoundaryState) != OSSTrustNetworkValCExportBoundaryStateActive ||
		strings.TrimSpace(model.ValCRemediationSuggestionState) != OSSTrustNetworkValCRemediationSuggestionStateActive ||
		strings.TrimSpace(model.ValCPRProposalState) != OSSTrustNetworkValCPRProposalStateActive ||
		strings.TrimSpace(model.ValCLocalOverrideState) != OSSTrustNetworkValCLocalOverrideStateActive ||
		strings.TrimSpace(model.ValCRemediationSafetyState) != OSSTrustNetworkValCRemediationSafetyStateActive ||
		strings.TrimSpace(model.ValCEcosystemConsistencyState) != OSSTrustNetworkValCEcosystemConsistencyStateActive ||
		strings.TrimSpace(model.ValCNoOverclaimState) != OSSTrustNetworkValCNoOverclaimStateActive {
		return OSSTrustNetworkValDDependencyStateBlocked
	}
	return OSSTrustNetworkValDDependencyStateActive
}

func EvaluateOSSTrustNetworkValDSignalCorrectnessState(model OSSTrustNetworkValDSignalCorrectness) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.SignalID,
		model.SignalLifecycleState,
		model.ReviewState,
		model.ReviewerDecisionState,
		model.SourceClass,
		model.SourceWeightClass,
		model.FreshnessState,
		model.LocalApplicabilityStatus,
		model.PropagationState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValDSignalCorrectnessStateIncomplete
	}
	if !ossTrustNetworkValDHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValDSignalCorrectnessStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vald-signal-correctness-001") ||
		!containsTrimmedString(ossTrustNetworkValDSignalLifecycleStates(), model.SignalLifecycleState) ||
		!containsTrimmedString(ossTrustNetworkValBReviewStates(), model.ReviewState) ||
		!containsTrimmedString(ossTrustNetworkValBReviewerDecisionStates(), model.ReviewerDecisionState) ||
		!containsTrimmedString(ossTrustNetworkValBCandidateSourceClasses(), model.SourceClass) ||
		!containsTrimmedString(ossTrustNetworkValBSourceWeightClasses(), model.SourceWeightClass) ||
		!containsTrimmedString(ossTrustNetworkValBLocalApplicabilityStates(), model.LocalApplicabilityStatus) ||
		!containsTrimmedString(ossTrustNetworkValBPropagationStates(), model.PropagationState) ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) ||
		model.CanonicalTruthClaim ||
		model.EnterpriseApprovalAuthorityClaim {
		return OSSTrustNetworkValDSignalCorrectnessStateBlocked
	}
	switch strings.TrimSpace(model.SignalLifecycleState) {
	case OSSTrustNetworkValDSignalLifecycleReviewed:
		if strings.TrimSpace(model.ReviewState) == OSSTrustNetworkValBReviewStateReviewed &&
			strings.TrimSpace(model.ReviewerDecisionState) == OSSTrustNetworkValBReviewerDecisionStateAccepted &&
			!model.CandidateDisplayedAsReviewed &&
			!model.RejectedUsableTrust &&
			!model.RevokedUsableTrust {
			return OSSTrustNetworkValDSignalCorrectnessStateActive
		}
		return OSSTrustNetworkValDSignalCorrectnessStateBlocked
	case OSSTrustNetworkValDSignalLifecycleCandidate:
		if model.CandidateDisplayedAsReviewed {
			return OSSTrustNetworkValDSignalCorrectnessStateBlocked
		}
		return OSSTrustNetworkValDSignalCorrectnessStatePartial
	case OSSTrustNetworkValDSignalLifecycleSuperseded:
		if strings.TrimSpace(model.ReplacementRef) == "" || model.SupersededUsableReviewedExchange {
			return OSSTrustNetworkValDSignalCorrectnessStateBlocked
		}
		return OSSTrustNetworkValDSignalCorrectnessStatePartial
	case OSSTrustNetworkValDSignalLifecycleRejected:
		if model.RejectedUsableTrust {
			return OSSTrustNetworkValDSignalCorrectnessStateBlocked
		}
		return OSSTrustNetworkValDSignalCorrectnessStateBlocked
	case OSSTrustNetworkValDSignalLifecycleRevoked:
		if model.RevokedUsableTrust {
			return OSSTrustNetworkValDSignalCorrectnessStateBlocked
		}
		return OSSTrustNetworkValDSignalCorrectnessStateBlocked
	default:
		return OSSTrustNetworkValDSignalCorrectnessStateBlocked
	}
}

func EvaluateOSSTrustNetworkValDReleaseFoundationState(model OSSTrustNetworkValDReleaseFoundation) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ReleaseTrustIntakeState,
		model.SigningSignalState,
		model.SigningVerificationState,
		model.MaintainerAttestationState,
		model.ProvenanceMaterialState,
		model.ProvenanceVerificationState,
		model.RegistryDescriptorState,
		model.RegistryDescriptorMode,
		model.RegistryMetadataState,
		model.TypoSquattingWarningState,
		model.DriftSignalState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValDReleaseFoundationStateIncomplete
	}
	if !ossTrustNetworkValDHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValDReleaseFoundationStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vald-release-foundation-001") {
		return OSSTrustNetworkValDReleaseFoundationStateBlocked
	}
	if strings.TrimSpace(model.ReleaseTrustIntakeState) != OSSTrustNetworkValAReleaseTrustIntakeStateActive ||
		strings.TrimSpace(model.SigningSignalState) != OSSTrustNetworkValASigningSignalStateActive ||
		strings.TrimSpace(model.SigningVerificationState) != OSSTrustNetworkValDSigningVerificationVerified ||
		!model.SigningScoped ||
		!model.SigningEvidenceLinked ||
		!model.SigningNonAuthoritative ||
		strings.TrimSpace(model.MaintainerAttestationState) != OSSTrustNetworkValAMaintainerAttestationStateActive ||
		!model.MaintainerIdentityBinding ||
		!model.MaintainerKeyLinked ||
		!model.MaintainerDelegationHandled ||
		!model.MaintainerRotationHandled ||
		!model.MaintainerRevocationHandled ||
		!model.MaintainerCompromiseHandled ||
		strings.TrimSpace(model.ProvenanceMaterialState) != OSSTrustNetworkValAProvenanceMaterialStateActive ||
		strings.TrimSpace(model.ProvenanceVerificationState) != OSSTrustNetworkValDProvenanceVerificationVerified ||
		!model.ProvenanceReleaseArtifactScoped ||
		!model.ProvenanceEvidenceLinked ||
		strings.TrimSpace(model.RegistryDescriptorState) != OSSTrustNetworkValARegistryDescriptorStateActive ||
		strings.TrimSpace(model.RegistryDescriptorMode) != OSSTrustNetworkValDRegistryDescriptorModeOnly ||
		model.LiveRegistryFetch ||
		strings.TrimSpace(model.RegistryMetadataState) != OSSTrustNetworkValARegistryMetadataStateActive ||
		!model.RegistryMetadataNormalized ||
		!model.RegistryMetadataFresh ||
		!model.RegistryMetadataScoped ||
		!model.RegistryMetadataEvidenceLinked ||
		model.RegistryMetadataCreatesReviewedTrust ||
		model.RegistryMetadataCreatesGlobalBlocklist ||
		strings.TrimSpace(model.TypoSquattingWarningState) != OSSTrustNetworkValATypoSquattingWarningStateActive ||
		!model.TypoWarningCandidateBounded ||
		model.TypoWarningAutoGlobalBlock ||
		model.TypoWarningCanonicalTruthClaim ||
		strings.TrimSpace(model.DriftSignalState) != OSSTrustNetworkValADriftSignalStateActive ||
		!model.DriftSignalEvidenceLinked ||
		!model.DriftSignalSourceWeighted ||
		!model.DriftSignalScoped ||
		!model.DriftSignalNonOverriding ||
		model.DriftOverridesLocalEnterprise {
		return OSSTrustNetworkValDReleaseFoundationStateBlocked
	}
	return OSSTrustNetworkValDReleaseFoundationStateActive
}

func EvaluateOSSTrustNetworkValDReviewedIntelligenceState(model OSSTrustNetworkValDReviewedIntelligence) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CandidateIntakeState,
		model.ReviewWorkflowState,
		model.ReviewState,
		model.ReviewerDecisionState,
		model.SharedVEXTriageState,
		model.SharedVEXState,
		model.SourceWeightingState,
		model.SourceClass,
		model.SourceWeightClass,
		model.LocalApplicabilityGateState,
		model.LocalApplicabilityStatus,
		model.ReviewerAuditabilityState,
		model.ReviewerRoleClass,
		model.ReviewerTimestamp,
		model.FreshnessState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValDReviewedIntelligenceStateIncomplete
	}
	if !ossTrustNetworkValDHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValDReviewedIntelligenceStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vald-reviewed-intelligence-001") ||
		!containsTrimmedString(ossTrustNetworkValBReviewStates(), model.ReviewState) ||
		!containsTrimmedString(ossTrustNetworkValBSharedVEXStates(), model.SharedVEXState) ||
		!containsTrimmedString(ossTrustNetworkValBCandidateSourceClasses(), model.SourceClass) ||
		!containsTrimmedString(ossTrustNetworkValBSourceWeightClasses(), model.SourceWeightClass) ||
		!containsTrimmedString(ossTrustNetworkValBLocalApplicabilityStates(), model.LocalApplicabilityStatus) ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) ||
		model.UniversalTrustScoreClaim ||
		model.IntegrityScoreClaim ||
		model.BadgeScoreClaim ||
		model.RewritesCanonicalEvidence ||
		model.RewritesLocalEnterpriseResult {
		return OSSTrustNetworkValDReviewedIntelligenceStateBlocked
	}
	if strings.TrimSpace(model.CandidateIntakeState) != OSSTrustNetworkValBCandidateSignalIntakeStateActive ||
		!model.CandidateIntakeNormalized ||
		strings.TrimSpace(model.ReviewWorkflowState) != OSSTrustNetworkValBReviewWorkflowStateActive ||
		strings.TrimSpace(model.ReviewState) != OSSTrustNetworkValBReviewStateReviewed ||
		strings.TrimSpace(model.ReviewerDecisionState) != OSSTrustNetworkValBReviewerDecisionStateAccepted ||
		strings.TrimSpace(model.ReviewerRationale) == "" ||
		strings.TrimSpace(model.SharedVEXTriageState) != OSSTrustNetworkValBSharedVEXTriageStateActive ||
		strings.TrimSpace(model.SourceWeightingState) != OSSTrustNetworkValBSourceWeightingStateActive ||
		strings.TrimSpace(model.LocalApplicabilityGateState) != OSSTrustNetworkValBLocalApplicabilityStateActive ||
		strings.TrimSpace(model.ReviewerAuditabilityState) != OSSTrustNetworkValBReviewerAuditabilityStateActive ||
		strings.TrimSpace(model.ReviewerRoleClass) == "" ||
		!model.ReviewerEvidenceLinked ||
		strings.TrimSpace(model.ReviewerTimestamp) == "" {
		return OSSTrustNetworkValDReviewedIntelligenceStateBlocked
	}
	if model.SharedVEXDisplayedAsReviewed && strings.TrimSpace(model.SharedVEXState) != OSSTrustNetworkValBSharedVEXStateReviewed {
		return OSSTrustNetworkValDReviewedIntelligenceStateBlocked
	}
	if strings.TrimSpace(model.SharedVEXState) != OSSTrustNetworkValBSharedVEXStateReviewed ||
		(strings.TrimSpace(model.LocalApplicabilityStatus) == OSSTrustNetworkValBLocalApplicabilityStatusUnknown && model.DisplayedAsApplicable) {
		return OSSTrustNetworkValDReviewedIntelligenceStateBlocked
	}
	return OSSTrustNetworkValDReviewedIntelligenceStateActive
}

func EvaluateOSSTrustNetworkValDPropagationSafetyState(model OSSTrustNetworkValDPropagationSafety) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.PropagationState,
		model.ReviewState,
		model.LocalApplicabilityStatus,
		model.SourceWeightClass,
		model.FreshnessState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValDPropagationSafetyStateIncomplete
	}
	if !ossTrustNetworkValDHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValDPropagationSafetyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vald-propagation-safety-001") ||
		!containsTrimmedString(ossTrustNetworkValBPropagationStates(), model.PropagationState) ||
		!containsTrimmedString(ossTrustNetworkValBReviewStates(), model.ReviewState) ||
		!containsTrimmedString(ossTrustNetworkValBLocalApplicabilityStates(), model.LocalApplicabilityStatus) ||
		!containsTrimmedString(ossTrustNetworkValBSourceWeightClasses(), model.SourceWeightClass) ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) ||
		model.AutomaticGlobalSpread ||
		model.GlobalBlocklistClaim ||
		model.EnterpriseOverride {
		return OSSTrustNetworkValDPropagationSafetyStateBlocked
	}
	switch strings.TrimSpace(model.PropagationState) {
	case OSSTrustNetworkValBPropagationStateReviewedExchange:
		if strings.TrimSpace(model.ReviewState) == OSSTrustNetworkValBReviewStateReviewed &&
			strings.TrimSpace(model.LocalApplicabilityStatus) == OSSTrustNetworkValBLocalApplicabilityStatusApplicable &&
			model.SimilarityContextGating {
			return OSSTrustNetworkValDPropagationSafetyStateActive
		}
		return OSSTrustNetworkValDPropagationSafetyStateBlocked
	case OSSTrustNetworkValBPropagationStateCandidateExchange:
		if model.CandidateDisplayedAsReviewed {
			return OSSTrustNetworkValDPropagationSafetyStateBlocked
		}
		return OSSTrustNetworkValDPropagationSafetyStatePartial
	case OSSTrustNetworkValBPropagationStateSuperseded:
		if strings.TrimSpace(model.ReplacementRef) == "" {
			return OSSTrustNetworkValDPropagationSafetyStateBlocked
		}
		return OSSTrustNetworkValDPropagationSafetyStatePartial
	default:
		return OSSTrustNetworkValDPropagationSafetyStateBlocked
	}
}

func EvaluateOSSTrustNetworkValDRemediationPRSafetyState(model OSSTrustNetworkValDRemediationPRSafety) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.SuggestionClass,
		model.ProposalState,
		model.RiskClass,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValDRemediationPRSafetyStateIncomplete
	}
	if !ossTrustNetworkValDHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValDRemediationPRSafetyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vald-remediation-pr-safety-001") ||
		!containsTrimmedString(ossTrustNetworkValCSuggestionClasses(), model.SuggestionClass) ||
		!containsTrimmedString(ossTrustNetworkValCRiskClasses(), model.RiskClass) ||
		!containsTrimmedString(ossTrustNetworkValCProposalStates(), model.ProposalState) ||
		model.DependencyMutationAttempt ||
		model.PolicyOverrideAttempt ||
		model.HiddenMutationPath ||
		model.ProductionApprovalClaim ||
		model.DeploymentApprovalClaim ||
		model.ProposalBranchWrite ||
		model.ProposalNetworkAction ||
		model.ProposalDependencyMutation ||
		model.ProposalPRCreation ||
		model.ProposalAutoMerge ||
		!model.ProposalAdvisoryOnly {
		return OSSTrustNetworkValDRemediationPRSafetyStateBlocked
	}
	if strings.TrimSpace(model.CompatibilityNote) == "" ||
		strings.TrimSpace(model.RiskNote) == "" ||
		strings.TrimSpace(model.RollbackNote) == "" ||
		strings.TrimSpace(model.TestValidationNote) == "" ||
		strings.TrimSpace(model.LocalApplicabilityNote) == "" {
		return OSSTrustNetworkValDRemediationPRSafetyStateBlocked
	}
	if strings.TrimSpace(model.SuggestionClass) == OSSTrustNetworkValCSuggestionClassNoAction &&
		(strings.TrimSpace(model.Rationale) == "" || model.NoActionHidesRisk) {
		return OSSTrustNetworkValDRemediationPRSafetyStateBlocked
	}
	if strings.TrimSpace(model.RiskClass) == OSSTrustNetworkValCRiskClassHigh && !model.ReviewerRequired {
		return OSSTrustNetworkValDRemediationPRSafetyStateBlocked
	}
	if strings.TrimSpace(model.ProposalState) != OSSTrustNetworkValCProposalStateProposalReady ||
		!model.ProposalReviewerRequired ||
		!model.ProposalNoAutomerge ||
		!model.ProposalNoHiddenMutation {
		if strings.TrimSpace(model.ProposalState) == OSSTrustNetworkValCProposalStateNeedsReview {
			return OSSTrustNetworkValDRemediationPRSafetyStatePartial
		}
		return OSSTrustNetworkValDRemediationPRSafetyStateBlocked
	}
	if strings.TrimSpace(model.SuggestionClass) == OSSTrustNetworkValCSuggestionClassUnsupported {
		return OSSTrustNetworkValDRemediationPRSafetyStatePartial
	}
	return OSSTrustNetworkValDRemediationPRSafetyStateActive
}

func EvaluateOSSTrustNetworkValDEcosystemVisibilityConsistencyState(model OSSTrustNetworkValDEcosystemVisibilityConsistency) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(
		model.VisibilityState,
		model.ReviewedSignalState,
		model.LocalApplicabilityState,
		model.SourceWeightingState,
		model.VisibilityFreshnessState,
		model.PackageStatusClass,
		model.ExportClass,
		model.LocalOverrideState,
		model.LocalOverrideScope,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValDEcosystemVisibilityConsistencyStateIncomplete
	}
	if !ossTrustNetworkValDHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValDEcosystemVisibilityConsistencyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vald-ecosystem-consistency-001") ||
		!containsTrimmedString(ossTrustNetworkValCVisibilityStates(), model.VisibilityState) ||
		!containsTrimmedString(ossTrustNetworkValCPackageStatusClasses(), model.PackageStatusClass) ||
		!containsTrimmedString(ossTrustNetworkValCExportClasses(), model.ExportClass) ||
		!containsTrimmedString(ossTrustNetworkValCLocalOverrideStates(), model.LocalOverrideState) ||
		!containsTrimmedString(ossTrustNetworkVal0ApplicabilityScopes(), model.LocalOverrideScope) ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.VisibilityFreshnessState) ||
		model.RewriteCanonicalEvidence ||
		model.SilentlySuppressReviewedNetworkIntelligence ||
		model.SharedSignalOverridesLocalDecision {
		return OSSTrustNetworkValDEcosystemVisibilityConsistencyStateBlocked
	}
	if model.CandidatePromotedToReviewed ||
		model.RejectedPromotedToActive ||
		model.RevokedPromotedToActive ||
		model.UnknownPromotedToActive {
		return OSSTrustNetworkValDEcosystemVisibilityConsistencyStateBlocked
	}
	if strings.TrimSpace(model.ExportClass) == OSSTrustNetworkValCExportClassPublicSummaryView &&
		(model.CanonicalInternalExposure || model.CertificationClaim || model.ApprovalClaim) {
		return OSSTrustNetworkValDEcosystemVisibilityConsistencyStateBlocked
	}
	switch strings.TrimSpace(model.VisibilityState) {
	case OSSTrustNetworkValCVisibilityVisible:
		if strings.TrimSpace(model.ReviewedSignalState) != OSSTrustNetworkValBReviewWorkflowStateActive ||
			strings.TrimSpace(model.LocalApplicabilityState) != OSSTrustNetworkValBLocalApplicabilityStateActive ||
			strings.TrimSpace(model.SourceWeightingState) != OSSTrustNetworkValBSourceWeightingStateActive ||
			strings.TrimSpace(model.PackageStatusClass) != OSSTrustNetworkValCPackageStatusReviewedSignalAvailable ||
			!model.PackageDisplayedAsReviewed ||
			strings.TrimSpace(model.LocalOverrideState) != OSSTrustNetworkValCOverrideStateNoOverride {
			return OSSTrustNetworkValDEcosystemVisibilityConsistencyStateBlocked
		}
		return OSSTrustNetworkValDEcosystemVisibilityConsistencyStateActive
	case OSSTrustNetworkValCVisibilityLimited:
		return OSSTrustNetworkValDEcosystemVisibilityConsistencyStatePartial
	default:
		if strings.TrimSpace(model.PackageStatusClass) == OSSTrustNetworkValCPackageStatusCandidateSignalAvailable ||
			strings.TrimSpace(model.PackageStatusClass) == OSSTrustNetworkValCPackageStatusLocalReviewNeeded {
			if model.PackageDisplayedAsReviewed {
				return OSSTrustNetworkValDEcosystemVisibilityConsistencyStateBlocked
			}
			return OSSTrustNetworkValDEcosystemVisibilityConsistencyStatePartial
		}
		return OSSTrustNetworkValDEcosystemVisibilityConsistencyStateBlocked
	}
}

func EvaluateOSSTrustNetworkValDEvidenceQualityState(model OSSTrustNetworkValDEvidenceQuality) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if len(model.Evidence) == 0 || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 ||
		len(model.DependencyProofSurfaceRefs) == 0 || len(model.DependencyEvidenceRefs) == 0 || len(model.DependencyEvidence) == 0 ||
		strings.TrimSpace(model.DependencyProjectionDisclaimer) == "" || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return OSSTrustNetworkValDEvidenceQualityStateIncomplete
	}
	if !ossTrustNetworkValDHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValDEvidenceQualityStateUnknown
	}
	if !ossTrustNetworkValCHasProjectionDisclaimer(model.DependencyProjectionDisclaimer) {
		return OSSTrustNetworkValDEvidenceQualityStateUnknown
	}
	if !containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValDProofSurfaceRefs()...) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, OSSTrustNetworkValDProofEvidenceRefs()...) ||
		!containsExactTrimmedStringSet(model.DependencyProofSurfaceRefs, OSSTrustNetworkValCProofSurfaceRefs()...) ||
		!containsExactTrimmedStringSet(model.DependencyEvidenceRefs, OSSTrustNetworkValCProofEvidenceRefs()...) ||
		!OSSTrustNetworkValDProofEvidenceQualityValid(model.Evidence, model.EvidenceRefs) ||
		!OSSTrustNetworkValCProofEvidenceQualityValid(model.DependencyEvidence, model.DependencyEvidenceRefs) {
		return OSSTrustNetworkValDEvidenceQualityStateBlocked
	}
	return OSSTrustNetworkValDEvidenceQualityStateActive
}

func EvaluateOSSTrustNetworkValDNoOverclaimState(model OSSTrustNetworkValDNoOverclaim) string {
	projectionDisclaimer := strings.TrimSpace(model.ProjectionDisclaimer)
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return OSSTrustNetworkValDNoOverclaimStateIncomplete
	}
	if !ossTrustNetworkValDHasProjectionDisclaimer(projectionDisclaimer) {
		return OSSTrustNetworkValDNoOverclaimStateUnknown
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
		model.PublicBadgeClaim ||
		model.OfficialOSSAuthorityClaim ||
		ossTrustNetworkValDContainsForbiddenClaim(model.ObservedClaims...) {
		return OSSTrustNetworkValDNoOverclaimStateBlocked
	}
	return OSSTrustNetworkValDNoOverclaimStateActive
}

func EvaluateOSSTrustNetworkValDState(model OSSTrustNetworkValDCore) string {
	if strings.TrimSpace(model.Point9State) == "" || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return OSSTrustNetworkValDStateIncomplete
	}
	if !ossTrustNetworkValDHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValDStateUnknown
	}
	if strings.TrimSpace(model.Point9State) != OSSTrustNetworkPoint9StateNotComplete ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValDProofSurfaceRefs()...) ||
		!OSSTrustNetworkValDProofEvidenceQualityValid(ossTrustNetworkValDEvidence(), model.EvidenceRefs) {
		return OSSTrustNetworkValDStateBlocked
	}
	states := []string{
		model.DependencyState,
		model.SignalCorrectnessState,
		model.ReleaseFoundationState,
		model.ReviewedIntelligenceState,
		model.PropagationSafetyState,
		model.RemediationPRSafetyState,
		model.EcosystemVisibilityConsistencyState,
		model.EvidenceQualityState,
		model.NoOverclaimState,
	}
	allActive := true
	for _, state := range states {
		if strings.TrimSpace(state) == "" {
			return OSSTrustNetworkValDStateIncomplete
		}
		if !strings.HasSuffix(strings.TrimSpace(state), "_active") {
			allActive = false
		}
	}
	if allActive {
		return OSSTrustNetworkValDStateActive
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_blocked") {
			return OSSTrustNetworkValDStateBlocked
		}
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_incomplete") {
			return OSSTrustNetworkValDStateIncomplete
		}
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_unknown") {
			return OSSTrustNetworkValDStateUnknown
		}
	}
	return OSSTrustNetworkValDStatePartial
}

func EvaluateOSSTrustNetworkValDPointsState(currentState string) string {
	_ = currentState
	return OSSTrustNetworkPoint9StateNotComplete
}

func EvaluateOSSTrustNetworkValDProofsState(model OSSTrustNetworkValDCore, limitations []string) string {
	baseState := strings.TrimSpace(model.CurrentState)
	if baseState == "" {
		baseState = OSSTrustNetworkValDStateUnknown
	}
	if !ossTrustNetworkValDHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValDProofSurfaceRefs()...) ||
		!OSSTrustNetworkValDProofEvidenceQualityValid(ossTrustNetworkValDEvidence(), model.EvidenceRefs) ||
		len(limitations) == 0 ||
		strings.TrimSpace(model.Point9State) != OSSTrustNetworkPoint9StateNotComplete {
		if baseState == OSSTrustNetworkValDStateActive {
			return OSSTrustNetworkValDStatePartial
		}
		return baseState
	}
	return baseState
}

func ossTrustNetworkValDBlockingReasons(model OSSTrustNetworkValDCore) []string {
	reasons := []string{}
	if model.DependencyState != OSSTrustNetworkValDDependencyStateActive {
		reasons = append(reasons, "OSTN Val C dependency is not exact, active, and evidence-safe.")
	}
	if model.SignalCorrectnessState != OSSTrustNetworkValDSignalCorrectnessStateActive {
		reasons = append(reasons, "Signal correctness is not exact, reviewed, fresh, and non-overclaiming.")
	}
	if model.ReleaseFoundationState != OSSTrustNetworkValDReleaseFoundationStateActive {
		reasons = append(reasons, "Release, provenance, maintainer, registry, typo-warning, or drift readiness is not exact and bounded.")
	}
	if model.ReviewedIntelligenceState != OSSTrustNetworkValDReviewedIntelligenceStateActive {
		reasons = append(reasons, "Reviewed intelligence or shared VEX readiness is not exact, reviewed, evidence-linked, and locally bounded.")
	}
	if model.PropagationSafetyState != OSSTrustNetworkValDPropagationSafetyStateActive {
		reasons = append(reasons, "Propagation or exchange safety is not exact, similarity-gated, and free of global spread or override behavior.")
	}
	if model.RemediationPRSafetyState != OSSTrustNetworkValDRemediationPRSafetyStateActive {
		reasons = append(reasons, "Remediation suggestion or PR proposal safety is not reviewer-gated, rollback-ready, and free of mutation paths.")
	}
	if model.EcosystemVisibilityConsistencyState != OSSTrustNetworkValDEcosystemVisibilityConsistencyStateActive {
		reasons = append(reasons, "Ecosystem visibility, export, package status, or local override consistency is not exact and bounded.")
	}
	if model.EvidenceQualityState != OSSTrustNetworkValDEvidenceQualityStateActive {
		reasons = append(reasons, "Val D evidence quality or exact reference discipline is not active.")
	}
	if model.NoOverclaimState != OSSTrustNetworkValDNoOverclaimStateActive {
		reasons = append(reasons, "Val D no-overclaim and no-global-truth guard is not active.")
	}
	return developerEcosystemValECollectText(reasons)
}

func ComputeOSSTrustNetworkValDCore(model OSSTrustNetworkValDCore) OSSTrustNetworkValDCore {
	model.DependencyState = EvaluateOSSTrustNetworkValDDependencyState(model.Dependency)
	model.SignalCorrectnessState = EvaluateOSSTrustNetworkValDSignalCorrectnessState(model.SignalCorrectness)
	model.ReleaseFoundationState = EvaluateOSSTrustNetworkValDReleaseFoundationState(model.ReleaseFoundation)
	model.ReviewedIntelligenceState = EvaluateOSSTrustNetworkValDReviewedIntelligenceState(model.ReviewedIntelligence)
	model.PropagationSafetyState = EvaluateOSSTrustNetworkValDPropagationSafetyState(model.PropagationSafety)
	model.RemediationPRSafetyState = EvaluateOSSTrustNetworkValDRemediationPRSafetyState(model.RemediationPRSafety)
	model.EcosystemVisibilityConsistencyState = EvaluateOSSTrustNetworkValDEcosystemVisibilityConsistencyState(model.EcosystemVisibilityConsistency)
	model.EvidenceQualityState = EvaluateOSSTrustNetworkValDEvidenceQualityState(model.EvidenceQuality)
	model.NoOverclaimState = EvaluateOSSTrustNetworkValDNoOverclaimState(model.NoOverclaim)
	model.Point9State = EvaluateOSSTrustNetworkValDPointsState(model.CurrentState)
	model.CurrentState = EvaluateOSSTrustNetworkValDState(model)
	model.BlockingReasons = ossTrustNetworkValDBlockingReasons(model)
	return model
}

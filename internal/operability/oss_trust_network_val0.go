package operability

import "strings"

const (
	OSSTrustNetworkPoint9StateNotComplete = "oss_trust_network_point_9_not_complete"

	OSSTrustNetworkVal0StateActive     = "oss_trust_network_val0_active"
	OSSTrustNetworkVal0StatePartial    = "oss_trust_network_val0_partial"
	OSSTrustNetworkVal0StateIncomplete = "oss_trust_network_val0_incomplete"
	OSSTrustNetworkVal0StateBlocked    = "oss_trust_network_val0_blocked"
	OSSTrustNetworkVal0StateUnknown    = "oss_trust_network_val0_unknown"

	OSSTrustNetworkVal0DependencyStateActive     = "oss_trust_network_val0_dependency_active"
	OSSTrustNetworkVal0DependencyStatePartial    = "oss_trust_network_val0_dependency_partial"
	OSSTrustNetworkVal0DependencyStateIncomplete = "oss_trust_network_val0_dependency_incomplete"
	OSSTrustNetworkVal0DependencyStateBlocked    = "oss_trust_network_val0_dependency_blocked"
	OSSTrustNetworkVal0DependencyStateUnknown    = "oss_trust_network_val0_dependency_unknown"

	OSSTrustNetworkVal0SignalContractStateActive     = "oss_trust_network_val0_signal_contract_active"
	OSSTrustNetworkVal0SignalContractStatePartial    = "oss_trust_network_val0_signal_contract_partial"
	OSSTrustNetworkVal0SignalContractStateIncomplete = "oss_trust_network_val0_signal_contract_incomplete"
	OSSTrustNetworkVal0SignalContractStateBlocked    = "oss_trust_network_val0_signal_contract_blocked"
	OSSTrustNetworkVal0SignalContractStateUnknown    = "oss_trust_network_val0_signal_contract_unknown"

	OSSTrustNetworkVal0TrustMarkingStateActive     = "oss_trust_network_val0_trust_marking_active"
	OSSTrustNetworkVal0TrustMarkingStatePartial    = "oss_trust_network_val0_trust_marking_partial"
	OSSTrustNetworkVal0TrustMarkingStateIncomplete = "oss_trust_network_val0_trust_marking_incomplete"
	OSSTrustNetworkVal0TrustMarkingStateBlocked    = "oss_trust_network_val0_trust_marking_blocked"
	OSSTrustNetworkVal0TrustMarkingStateUnknown    = "oss_trust_network_val0_trust_marking_unknown"

	OSSTrustNetworkVal0MaintainerIdentityStateActive     = "oss_trust_network_val0_maintainer_identity_active"
	OSSTrustNetworkVal0MaintainerIdentityStatePartial    = "oss_trust_network_val0_maintainer_identity_partial"
	OSSTrustNetworkVal0MaintainerIdentityStateIncomplete = "oss_trust_network_val0_maintainer_identity_incomplete"
	OSSTrustNetworkVal0MaintainerIdentityStateBlocked    = "oss_trust_network_val0_maintainer_identity_blocked"
	OSSTrustNetworkVal0MaintainerIdentityStateUnknown    = "oss_trust_network_val0_maintainer_identity_unknown"

	OSSTrustNetworkVal0RegistryFreshnessStateActive     = "oss_trust_network_val0_registry_freshness_active"
	OSSTrustNetworkVal0RegistryFreshnessStatePartial    = "oss_trust_network_val0_registry_freshness_partial"
	OSSTrustNetworkVal0RegistryFreshnessStateIncomplete = "oss_trust_network_val0_registry_freshness_incomplete"
	OSSTrustNetworkVal0RegistryFreshnessStateBlocked    = "oss_trust_network_val0_registry_freshness_blocked"
	OSSTrustNetworkVal0RegistryFreshnessStateUnknown    = "oss_trust_network_val0_registry_freshness_unknown"

	OSSTrustNetworkVal0SharedVEXStateActive     = "oss_trust_network_val0_shared_vex_active"
	OSSTrustNetworkVal0SharedVEXStatePartial    = "oss_trust_network_val0_shared_vex_partial"
	OSSTrustNetworkVal0SharedVEXStateIncomplete = "oss_trust_network_val0_shared_vex_incomplete"
	OSSTrustNetworkVal0SharedVEXStateBlocked    = "oss_trust_network_val0_shared_vex_blocked"
	OSSTrustNetworkVal0SharedVEXStateUnknown    = "oss_trust_network_val0_shared_vex_unknown"

	OSSTrustNetworkVal0PropagationStateActive     = "oss_trust_network_val0_propagation_active"
	OSSTrustNetworkVal0PropagationStatePartial    = "oss_trust_network_val0_propagation_partial"
	OSSTrustNetworkVal0PropagationStateIncomplete = "oss_trust_network_val0_propagation_incomplete"
	OSSTrustNetworkVal0PropagationStateBlocked    = "oss_trust_network_val0_propagation_blocked"
	OSSTrustNetworkVal0PropagationStateUnknown    = "oss_trust_network_val0_propagation_unknown"

	OSSTrustNetworkVal0LocalApplicabilityStateActive     = "oss_trust_network_val0_local_applicability_active"
	OSSTrustNetworkVal0LocalApplicabilityStatePartial    = "oss_trust_network_val0_local_applicability_partial"
	OSSTrustNetworkVal0LocalApplicabilityStateIncomplete = "oss_trust_network_val0_local_applicability_incomplete"
	OSSTrustNetworkVal0LocalApplicabilityStateBlocked    = "oss_trust_network_val0_local_applicability_blocked"
	OSSTrustNetworkVal0LocalApplicabilityStateUnknown    = "oss_trust_network_val0_local_applicability_unknown"

	OSSTrustNetworkVal0NoOverclaimStateActive     = "oss_trust_network_val0_no_overclaim_active"
	OSSTrustNetworkVal0NoOverclaimStatePartial    = "oss_trust_network_val0_no_overclaim_partial"
	OSSTrustNetworkVal0NoOverclaimStateIncomplete = "oss_trust_network_val0_no_overclaim_incomplete"
	OSSTrustNetworkVal0NoOverclaimStateBlocked    = "oss_trust_network_val0_no_overclaim_blocked"
	OSSTrustNetworkVal0NoOverclaimStateUnknown    = "oss_trust_network_val0_no_overclaim_unknown"

	OSSTrustNetworkReviewStateCandidate  = "candidate"
	OSSTrustNetworkReviewStateReviewed   = "reviewed"
	OSSTrustNetworkReviewStateRejected   = "rejected"
	OSSTrustNetworkReviewStateSuperseded = "superseded"
	OSSTrustNetworkReviewStateRevoked    = "revoked"

	OSSTrustNetworkSignalClassProvenanceVerified   = "provenance_verified"
	OSSTrustNetworkSignalClassMaintainerAttested   = "maintainer_attested"
	OSSTrustNetworkSignalClassReviewedTriage       = "reviewed_triage_available"
	OSSTrustNetworkSignalClassReleaseDriftWarning  = "release_drift_warning"
	OSSTrustNetworkSignalClassTypoSquattingWarning = "typo_squatting_warning"
	OSSTrustNetworkSignalClassUnsupportedSignal    = "unsupported_signal"
	OSSTrustNetworkSignalClassRevokedTrustMarker   = "revoked_trust_marker"

	OSSTrustNetworkSourceClassMaintainerAttestation = "maintainer_attestation"
	OSSTrustNetworkSourceClassRegistryMetadata      = "registry_metadata"
	OSSTrustNetworkSourceClassSharedVEXTriage       = "shared_vex_triage"
	OSSTrustNetworkSourceClassCommunityReport       = "community_report"
	OSSTrustNetworkSourceClassEnterpriseObservation = "enterprise_observation"
	OSSTrustNetworkSourceClassProvenanceMaterial    = "provenance_material"

	OSSTrustNetworkSourceWeightLow    = "low"
	OSSTrustNetworkSourceWeightMedium = "medium"
	OSSTrustNetworkSourceWeightHigh   = "high"

	OSSTrustNetworkConfidenceLow     = "low"
	OSSTrustNetworkConfidenceMedium  = "medium"
	OSSTrustNetworkConfidenceHigh    = "high"
	OSSTrustNetworkConfidenceBounded = "bounded"

	OSSTrustNetworkApplicabilityEnterpriseLocal = "enterprise_local"
	OSSTrustNetworkApplicabilityProjectScoped   = "project_scoped"

	OSSTrustNetworkPropagationLocalOnly        = "local_only"
	OSSTrustNetworkPropagationProjectFamily    = "project_family"
	OSSTrustNetworkPropagationReviewedExchange = "reviewed_exchange"
)

type OSSTrustNetworkVal0DependencySnapshot struct {
	CurrentState         string   `json:"current_state"`
	Point8State          string   `json:"point_8_state"`
	Point8PassAllowed    bool     `json:"point_8_pass_allowed"`
	Point8PassReason     string   `json:"point_8_pass_reason"`
	ClosureState         string   `json:"closure_state"`
	NoOverclaimState     string   `json:"no_overclaim_state"`
	ProofSurfaceRefs     []string `json:"proof_surface_refs,omitempty"`
	EvidenceRefs         []string `json:"evidence_refs,omitempty"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkVal0SignalContract struct {
	CurrentState               string   `json:"current_state"`
	ContractID                 string   `json:"contract_id"`
	Version                    string   `json:"version"`
	SupportedReviewStates      []string `json:"supported_review_states,omitempty"`
	SignalID                   string   `json:"signal_id"`
	PackageOrProjectIdentity   string   `json:"package_or_project_identity"`
	RegistryOrEcosystem        string   `json:"registry_or_ecosystem"`
	PackageVersionOrReleaseRef string   `json:"package_version_or_release_ref"`
	SignalClass                string   `json:"signal_class"`
	SourceClass                string   `json:"source_class"`
	SourceWeightClass          string   `json:"source_weight_class"`
	EvidenceRefs               []string `json:"evidence_refs,omitempty"`
	FreshnessState             string   `json:"freshness_state"`
	ConfidenceClass            string   `json:"confidence_class"`
	ReviewState                string   `json:"review_state"`
	PresentedAsReviewed        bool     `json:"presented_as_reviewed"`
	ApplicabilityScope         string   `json:"applicability_scope"`
	PropagationScope           string   `json:"propagation_scope"`
	Caveats                    []string `json:"caveats,omitempty"`
	BlockingReasons            []string `json:"blocking_reasons,omitempty"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkVal0TrustMarkingSemantics struct {
	CurrentState                   string   `json:"current_state"`
	DisciplineID                   string   `json:"discipline_id"`
	Version                        string   `json:"version"`
	AllowedTrustMarkingClasses     []string `json:"allowed_trust_marking_classes,omitempty"`
	GenericVerifiedBadgeClaim      bool     `json:"generic_verified_badge_claim"`
	BroadProjectSafetyBadgeClaim   bool     `json:"broad_project_safety_badge_claim"`
	IntegrityScoreClaim            bool     `json:"integrity_score_claim"`
	ScoreGreaterThanNinetyClaim    bool     `json:"score_greater_than_ninety_claim"`
	UniversalTrustScoreClaim       bool     `json:"universal_trust_score_claim"`
	SafePackageClaim               bool     `json:"safe_package_claim"`
	CertifiedPackageClaim          bool     `json:"certified_package_claim"`
	OfficialEcosystemApprovalClaim bool     `json:"official_ecosystem_approval_claim"`
	ProjectionDisclaimer           string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkVal0MaintainerIdentityDiscipline struct {
	CurrentState              string `json:"current_state"`
	DisciplineID              string `json:"discipline_id"`
	Version                   string `json:"version"`
	MaintainerIdentityBinding bool   `json:"maintainer_identity_binding"`
	KeyToMaintainerLinkage    bool   `json:"key_to_maintainer_linkage"`
	DelegatedSigningReviewed  bool   `json:"delegated_signing_reviewed"`
	KeyRotationHandled        bool   `json:"key_rotation_handled"`
	KeyRevocationHandled      bool   `json:"key_revocation_handled"`
	CompromiseHandlingVisible bool   `json:"compromise_handling_visible"`
	TrustContinuityBounded    bool   `json:"trust_continuity_bounded"`
	EnterpriseApprovalClaim   bool   `json:"enterprise_approval_claim"`
	CanonicalTruthClaim       bool   `json:"canonical_truth_claim"`
	ProjectionDisclaimer      string `json:"projection_disclaimer"`
}

type OSSTrustNetworkVal0RegistryFreshnessDiscipline struct {
	CurrentState                       string `json:"current_state"`
	DisciplineID                       string `json:"discipline_id"`
	Version                            string `json:"version"`
	FreshnessState                     string `json:"freshness_state"`
	SupportedRegistryMetadataBounded   bool   `json:"supported_registry_metadata_bounded"`
	UnsupportedRegistryStateExplicit   bool   `json:"unsupported_registry_state_explicit"`
	StaleRegistryBlocksActive          bool   `json:"stale_registry_blocks_active"`
	UnknownFreshnessFailsClosed        bool   `json:"unknown_freshness_fails_closed"`
	RegistryMetadataAloneReviewedTrust bool   `json:"registry_metadata_alone_reviewed_trust"`
	RegistryHeuristicsGlobalBlocklist  bool   `json:"registry_heuristics_global_blocklist"`
	ProjectionDisclaimer               string `json:"projection_disclaimer"`
}

type OSSTrustNetworkVal0SharedVEXTriageDiscipline struct {
	CurrentState              string `json:"current_state"`
	DisciplineID              string `json:"discipline_id"`
	Version                   string `json:"version"`
	ReviewState               string `json:"review_state"`
	PresentedAsReviewed       bool   `json:"presented_as_reviewed"`
	ReplacementRef            string `json:"replacement_ref"`
	RejectedPropagatedUsable  bool   `json:"rejected_propagated_usable"`
	RevokedUsable             bool   `json:"revoked_usable"`
	LocalApplicabilityBounded bool   `json:"local_applicability_bounded"`
	SharedIntelligenceOnly    bool   `json:"shared_intelligence_only"`
	ProjectionDisclaimer      string `json:"projection_disclaimer"`
}

type OSSTrustNetworkVal0PropagationDiscipline struct {
	CurrentState                          string   `json:"current_state"`
	DisciplineID                          string   `json:"discipline_id"`
	Version                               string   `json:"version"`
	ReviewState                           string   `json:"review_state"`
	SourceWeightClass                     string   `json:"source_weight_class"`
	FreshnessState                        string   `json:"freshness_state"`
	EvidenceRefs                          []string `json:"evidence_refs,omitempty"`
	ApplicabilityScope                    string   `json:"applicability_scope"`
	PropagationScope                      string   `json:"propagation_scope"`
	SimilarityContextGating               bool     `json:"similarity_context_gating"`
	Caveats                               []string `json:"caveats,omitempty"`
	AutomaticGlobalSpread                 bool     `json:"automatic_global_spread"`
	CandidateTreatedAsReviewed            bool     `json:"candidate_treated_as_reviewed"`
	RegistryHeuristicAsCanonicalProof     bool     `json:"registry_heuristic_as_canonical_proof"`
	CommunityReportAsFinalTruth           bool     `json:"community_report_as_final_truth"`
	OverridesLocalEnterpriseWithoutReview bool     `json:"overrides_local_enterprise_without_review"`
	CommunityCandidateCanonicalTruth      bool     `json:"community_candidate_canonical_truth"`
	ProjectionDisclaimer                  string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkVal0LocalApplicabilityDiscipline struct {
	CurrentState                      string `json:"current_state"`
	DisciplineID                      string `json:"discipline_id"`
	Version                           string `json:"version"`
	NetworkSignalAdvisoryOnly         bool   `json:"network_signal_advisory_only"`
	LocalApplicabilityExplicit        bool   `json:"local_applicability_explicit"`
	LocalOverrideVisible              bool   `json:"local_override_visible"`
	LocalOverrideEvidenceLinked       bool   `json:"local_override_evidence_linked"`
	CommunitySignalRewritesEnterprise bool   `json:"community_signal_rewrites_enterprise"`
	ProjectionDisclaimer              string `json:"projection_disclaimer"`
}

type OSSTrustNetworkVal0NoOverclaimDiscipline struct {
	CurrentState         string   `json:"current_state"`
	DisciplineID         string   `json:"discipline_id"`
	Version              string   `json:"version"`
	ObservedClaims       []string `json:"observed_claims,omitempty"`
	GlobalTruthClaim     bool     `json:"global_truth_claim"`
	CertifiedClaim       bool     `json:"certified_claim"`
	RegulatorApproved    bool     `json:"regulator_approved"`
	ProductionApproved   bool     `json:"production_approved"`
	LegalCertification   bool     `json:"legal_certification"`
	PatentCleared        bool     `json:"patent_cleared"`
	FTOPCleared          bool     `json:"fto_cleared"`
	ProjectionDisclaimer string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkVal0Foundation struct {
	CurrentState            string                                          `json:"current_state"`
	Point9State             string                                          `json:"point_9_state"`
	DependencyState         string                                          `json:"dependency_state"`
	SignalContractState     string                                          `json:"signal_contract_state"`
	TrustMarkingState       string                                          `json:"trust_marking_state"`
	MaintainerIdentityState string                                          `json:"maintainer_identity_state"`
	RegistryFreshnessState  string                                          `json:"registry_freshness_state"`
	SharedVEXState          string                                          `json:"shared_vex_state"`
	PropagationState        string                                          `json:"propagation_state"`
	LocalApplicabilityState string                                          `json:"local_applicability_state"`
	NoOverclaimState        string                                          `json:"no_overclaim_state"`
	Dependency              OSSTrustNetworkVal0DependencySnapshot           `json:"dependency"`
	SignalContract          OSSTrustNetworkVal0SignalContract               `json:"signal_contract"`
	TrustMarking            OSSTrustNetworkVal0TrustMarkingSemantics        `json:"trust_marking"`
	MaintainerIdentity      OSSTrustNetworkVal0MaintainerIdentityDiscipline `json:"maintainer_identity"`
	RegistryFreshness       OSSTrustNetworkVal0RegistryFreshnessDiscipline  `json:"registry_freshness"`
	SharedVEXTriage         OSSTrustNetworkVal0SharedVEXTriageDiscipline    `json:"shared_vex_triage"`
	Propagation             OSSTrustNetworkVal0PropagationDiscipline        `json:"propagation"`
	LocalApplicability      OSSTrustNetworkVal0LocalApplicabilityDiscipline `json:"local_applicability"`
	NoOverclaim             OSSTrustNetworkVal0NoOverclaimDiscipline        `json:"no_overclaim"`
	ProofSurfaceRefs        []string                                        `json:"proof_surface_refs,omitempty"`
	EvidenceRefs            []string                                        `json:"evidence_refs,omitempty"`
	BlockingReasons         []string                                        `json:"blocking_reasons,omitempty"`
	WhyPoint9NotComplete    []string                                        `json:"why_point_9_not_complete,omitempty"`
	ProjectionDisclaimer    string                                          `json:"projection_disclaimer"`
}

func ossTrustNetworkVal0ProjectionDisclaimer() string {
	return "projection_only not_canonical_truth oss_trust_network_val0 advisory_projection"
}

func ossTrustNetworkVal0ReviewStates() []string {
	return []string{
		OSSTrustNetworkReviewStateCandidate,
		OSSTrustNetworkReviewStateReviewed,
		OSSTrustNetworkReviewStateRejected,
		OSSTrustNetworkReviewStateSuperseded,
		OSSTrustNetworkReviewStateRevoked,
	}
}

func ossTrustNetworkVal0TrustMarkingClasses() []string {
	return []string{
		OSSTrustNetworkSignalClassProvenanceVerified,
		OSSTrustNetworkSignalClassMaintainerAttested,
		OSSTrustNetworkSignalClassReviewedTriage,
		OSSTrustNetworkSignalClassReleaseDriftWarning,
		OSSTrustNetworkSignalClassTypoSquattingWarning,
		OSSTrustNetworkSignalClassUnsupportedSignal,
		OSSTrustNetworkSignalClassRevokedTrustMarker,
	}
}

func ossTrustNetworkVal0SourceClasses() []string {
	return []string{
		OSSTrustNetworkSourceClassMaintainerAttestation,
		OSSTrustNetworkSourceClassRegistryMetadata,
		OSSTrustNetworkSourceClassSharedVEXTriage,
		OSSTrustNetworkSourceClassCommunityReport,
		OSSTrustNetworkSourceClassEnterpriseObservation,
		OSSTrustNetworkSourceClassProvenanceMaterial,
	}
}

func ossTrustNetworkVal0SourceWeightClasses() []string {
	return []string{
		OSSTrustNetworkSourceWeightLow,
		OSSTrustNetworkSourceWeightMedium,
		OSSTrustNetworkSourceWeightHigh,
	}
}

func ossTrustNetworkVal0ConfidenceClasses() []string {
	return []string{
		OSSTrustNetworkConfidenceLow,
		OSSTrustNetworkConfidenceMedium,
		OSSTrustNetworkConfidenceHigh,
		OSSTrustNetworkConfidenceBounded,
	}
}

func ossTrustNetworkVal0PropagationScopes() []string {
	return []string{
		OSSTrustNetworkPropagationLocalOnly,
		OSSTrustNetworkPropagationProjectFamily,
		OSSTrustNetworkPropagationReviewedExchange,
	}
}

func ossTrustNetworkVal0ApplicabilityScopes() []string {
	return []string{
		OSSTrustNetworkApplicabilityEnterpriseLocal,
		OSSTrustNetworkApplicabilityProjectScoped,
	}
}

func OSSTrustNetworkVal0ProofSurfaceRefs() []string {
	return []string{
		"/v1/developer-ecosystem/vale/closure",
		"/v1/developer-ecosystem/vale/proofs",
		"/v1/oss-trust-network/val0/status",
		"/v1/oss-trust-network/val0/proofs",
	}
}

func OSSTrustNetworkVal0ProofEvidenceRefs() []string {
	return []string{
		"evidence:ostn-val0-dependency-001",
		"evidence:ostn-val0-signal-contract-001",
		"evidence:ostn-val0-trust-marking-001",
		"evidence:ostn-val0-maintainer-identity-001",
		"evidence:ostn-val0-registry-freshness-001",
		"evidence:ostn-val0-shared-vex-001",
		"evidence:ostn-val0-propagation-001",
		"evidence:ostn-val0-local-applicability-001",
		"evidence:ostn-val0-no-overclaim-001",
		"evidence:ostn-val0-canonical-boundary-001",
		"evidence:ostn-val0-point9-governance-001",
	}
}

type ossTrustNetworkVal0ExpectedEvidenceMetadata struct {
	EvidenceType string
	Source       string
	Scope        string
}

func ossTrustNetworkVal0ExpectedEvidenceMetadataByID() map[string]ossTrustNetworkVal0ExpectedEvidenceMetadata {
	return map[string]ossTrustNetworkVal0ExpectedEvidenceMetadata{
		"evidence:ostn-val0-dependency-001":          {EvidenceType: "dependency_state", Source: "oss-trust-network/val0/dependency", Scope: "point8_dependency"},
		"evidence:ostn-val0-signal-contract-001":     {EvidenceType: "signal_contract", Source: "oss-trust-network/val0/signal-contract", Scope: "oss_signal_contract"},
		"evidence:ostn-val0-trust-marking-001":       {EvidenceType: "trust_marking", Source: "oss-trust-network/val0/trust-marking", Scope: "trust_marking_semantics"},
		"evidence:ostn-val0-maintainer-identity-001": {EvidenceType: "maintainer_identity", Source: "oss-trust-network/val0/maintainer-identity", Scope: "maintainer_identity_discipline"},
		"evidence:ostn-val0-registry-freshness-001":  {EvidenceType: "registry_freshness", Source: "oss-trust-network/val0/registry-freshness", Scope: "registry_freshness_discipline"},
		"evidence:ostn-val0-shared-vex-001":          {EvidenceType: "shared_vex", Source: "oss-trust-network/val0/shared-vex", Scope: "shared_vex_discipline"},
		"evidence:ostn-val0-propagation-001":         {EvidenceType: "propagation", Source: "oss-trust-network/val0/propagation", Scope: "propagation_discipline"},
		"evidence:ostn-val0-local-applicability-001": {EvidenceType: "local_applicability", Source: "oss-trust-network/val0/local-applicability", Scope: "local_enterprise_boundary"},
		"evidence:ostn-val0-no-overclaim-001":        {EvidenceType: "no_overclaim", Source: "oss-trust-network/val0/no-overclaim", Scope: "no_overclaim_discipline"},
		"evidence:ostn-val0-canonical-boundary-001":  {EvidenceType: "canonical_boundary", Source: "oss-trust-network/val0/canonical-boundary", Scope: "canonical_evidence_boundary"},
		"evidence:ostn-val0-point9-governance-001":   {EvidenceType: "state_governance", Source: "oss-trust-network/val0/point9-governance", Scope: "point9_governance"},
	}
}

func ossTrustNetworkVal0Evidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:ostn-val0-dependency-001", EvidenceType: "dependency_state", Source: "oss-trust-network/val0/dependency", Timestamp: "2026-04-29T08:40:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point8_dependency", Caveats: []string{"Točka 9 / Val 0 depends on accepted Točka 8 / Val E integrated closure only."}},
		{EvidenceID: "evidence:ostn-val0-signal-contract-001", EvidenceType: "signal_contract", Source: "oss-trust-network/val0/signal-contract", Timestamp: "2026-04-29T08:41:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "oss_signal_contract", Caveats: []string{"OSS trust signals remain bounded, reviewable, freshness-aware, and evidence-linked."}},
		{EvidenceID: "evidence:ostn-val0-trust-marking-001", EvidenceType: "trust_marking", Source: "oss-trust-network/val0/trust-marking", Timestamp: "2026-04-29T08:42:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "trust_marking_semantics", Caveats: []string{"Trust marking classes remain bounded and do not imply generic safety, certification, or approval."}},
		{EvidenceID: "evidence:ostn-val0-maintainer-identity-001", EvidenceType: "maintainer_identity", Source: "oss-trust-network/val0/maintainer-identity", Timestamp: "2026-04-29T08:43:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "maintainer_identity_discipline", Caveats: []string{"Maintainer attestations remain bounded signals and do not become enterprise approval or canonical truth."}},
		{EvidenceID: "evidence:ostn-val0-registry-freshness-001", EvidenceType: "registry_freshness", Source: "oss-trust-network/val0/registry-freshness", Timestamp: "2026-04-29T08:44:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "registry_freshness_discipline", Caveats: []string{"Registry metadata remains bounded evidence input and cannot create reviewed trust by itself."}},
		{EvidenceID: "evidence:ostn-val0-shared-vex-001", EvidenceType: "shared_vex", Source: "oss-trust-network/val0/shared-vex", Timestamp: "2026-04-29T08:45:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "shared_vex_discipline", Caveats: []string{"Shared VEX remains reviewed shared intelligence rather than global truth."}},
		{EvidenceID: "evidence:ostn-val0-propagation-001", EvidenceType: "propagation", Source: "oss-trust-network/val0/propagation", Timestamp: "2026-04-29T08:46:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "propagation_discipline", Caveats: []string{"Candidate or heuristic signals cannot auto-propagate as reviewed network truth."}},
		{EvidenceID: "evidence:ostn-val0-local-applicability-001", EvidenceType: "local_applicability", Source: "oss-trust-network/val0/local-applicability", Timestamp: "2026-04-29T08:47:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "local_enterprise_boundary", Caveats: []string{"OSS trust network signals remain advisory for local enterprise applicability unless explicitly governed otherwise."}},
		{EvidenceID: "evidence:ostn-val0-no-overclaim-001", EvidenceType: "no_overclaim", Source: "oss-trust-network/val0/no-overclaim", Timestamp: "2026-04-29T08:48:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim_discipline", Caveats: []string{"OSTN Val 0 cannot certify packages, create global truth, or return forbidden point pass claims."}},
		{EvidenceID: "evidence:ostn-val0-canonical-boundary-001", EvidenceType: "canonical_boundary", Source: "oss-trust-network/val0/canonical-boundary", Timestamp: "2026-04-29T08:49:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "canonical_evidence_boundary", Caveats: []string{"Canonical execution, audit, and evidence remain the only source of truth."}},
		{EvidenceID: "evidence:ostn-val0-point9-governance-001", EvidenceType: "state_governance", Source: "oss-trust-network/val0/point9-governance", Timestamp: "2026-04-29T08:50:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point9_governance", Caveats: []string{"point_9_state remains not complete and later waves are required for any final pass."}},
	}
}

func OSSTrustNetworkVal0ProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, refs []string) bool {
	if !containsExactTrimmedStringSet(refs, OSSTrustNetworkVal0ProofEvidenceRefs()...) {
		return false
	}
	expected := ossTrustNetworkVal0ExpectedEvidenceMetadataByID()
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

func ossTrustNetworkVal0HasProjectionDisclaimer(value string) bool {
	normalized := strings.ToLower(strings.TrimSpace(value))
	return strings.Contains(normalized, "projection_only") &&
		strings.Contains(normalized, "not_canonical_truth") &&
		strings.Contains(normalized, "oss_trust_network_val0")
}

func ossTrustNetworkVal0FreshnessStateIsExactlyFresh(value string) bool {
	switch strings.TrimSpace(value) {
	case IntelligenceCalibrationFreshnessFresh:
		return true
	case IntelligenceCalibrationFreshnessUnknown, IntelligenceCalibrationFreshnessStale, IntelligenceCalibrationFreshnessExpired, IntelligenceCalibrationFreshnessUnsupported:
		return false
	default:
		return false
	}
}

func ossTrustNetworkVal0ContainsForbiddenClaim(values ...string) bool {
	allowed := map[string]struct{}{
		"bounded oss trust signal":      {},
		"reviewed oss trust signal":     {},
		"candidate oss signal":          {},
		"source-weighted signal":        {},
		"evidence-linked trust marking": {},
		"provenance-aware signal":       {},
		"verifier-friendly oss signal":  {},
		"local applicability context":   {},
		"advisory network signal":       {},
		"reviewed shared intelligence":  {},
		"not a global truth layer":      {},
		"not formal certification":      {},
	}
	disallowed := []string{
		"changelock verified",
		"certified",
		"official certified",
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

func OSSTrustNetworkVal0DependencySnapshotModel() OSSTrustNetworkVal0DependencySnapshot {
	model := ComputeDeveloperEcosystemValEClosure(DeveloperEcosystemValEIntegratedClosureModel())
	return OSSTrustNetworkVal0DependencySnapshot{
		CurrentState:         model.CurrentState,
		Point8State:          model.Point8State,
		Point8PassAllowed:    model.Point8PassAllowed,
		Point8PassReason:     model.Point8PassReason,
		ClosureState:         model.ClosureState,
		NoOverclaimState:     model.NoOverclaimState,
		ProofSurfaceRefs:     append([]string{}, model.ProofSurfaceRefs...),
		EvidenceRefs:         append([]string{}, model.EvidenceRefs...),
		ProjectionDisclaimer: model.ProjectionDisclaimer,
	}
}

func OSSTrustNetworkVal0SignalContractModel() OSSTrustNetworkVal0SignalContract {
	return OSSTrustNetworkVal0SignalContract{
		CurrentState:               OSSTrustNetworkVal0SignalContractStateActive,
		ContractID:                 "oss_trust_signal_contract",
		Version:                    "v0",
		SupportedReviewStates:      ossTrustNetworkVal0ReviewStates(),
		SignalID:                   "ostn-signal-provenance-reviewed-001",
		PackageOrProjectIdentity:   "github.com/example/project",
		RegistryOrEcosystem:        "github",
		PackageVersionOrReleaseRef: "refs/tags/v1.2.3",
		SignalClass:                OSSTrustNetworkSignalClassReviewedTriage,
		SourceClass:                OSSTrustNetworkSourceClassSharedVEXTriage,
		SourceWeightClass:          OSSTrustNetworkSourceWeightMedium,
		EvidenceRefs:               []string{"evidence:ostn-val0-signal-contract-001"},
		FreshnessState:             IntelligenceCalibrationFreshnessFresh,
		ConfidenceClass:            OSSTrustNetworkConfidenceBounded,
		ReviewState:                OSSTrustNetworkReviewStateReviewed,
		PresentedAsReviewed:        true,
		ApplicabilityScope:         OSSTrustNetworkApplicabilityEnterpriseLocal,
		PropagationScope:           OSSTrustNetworkPropagationReviewedExchange,
		Caveats:                    []string{"reviewed shared intelligence only", "not a global truth layer"},
		BlockingReasons:            []string{},
		ProjectionDisclaimer:       ossTrustNetworkVal0ProjectionDisclaimer(),
	}
}

func OSSTrustNetworkVal0TrustMarkingModel() OSSTrustNetworkVal0TrustMarkingSemantics {
	return OSSTrustNetworkVal0TrustMarkingSemantics{
		CurrentState:               OSSTrustNetworkVal0TrustMarkingStateActive,
		DisciplineID:               "oss_trust_marking_semantics",
		Version:                    "v0",
		AllowedTrustMarkingClasses: ossTrustNetworkVal0TrustMarkingClasses(),
		ProjectionDisclaimer:       ossTrustNetworkVal0ProjectionDisclaimer(),
	}
}

func OSSTrustNetworkVal0MaintainerIdentityModel() OSSTrustNetworkVal0MaintainerIdentityDiscipline {
	return OSSTrustNetworkVal0MaintainerIdentityDiscipline{
		CurrentState:              OSSTrustNetworkVal0MaintainerIdentityStateActive,
		DisciplineID:              "oss_maintainer_identity",
		Version:                   "v0",
		MaintainerIdentityBinding: true,
		KeyToMaintainerLinkage:    true,
		DelegatedSigningReviewed:  true,
		KeyRotationHandled:        true,
		KeyRevocationHandled:      true,
		CompromiseHandlingVisible: true,
		TrustContinuityBounded:    true,
		ProjectionDisclaimer:      ossTrustNetworkVal0ProjectionDisclaimer(),
	}
}

func OSSTrustNetworkVal0RegistryFreshnessModel() OSSTrustNetworkVal0RegistryFreshnessDiscipline {
	return OSSTrustNetworkVal0RegistryFreshnessDiscipline{
		CurrentState:                     OSSTrustNetworkVal0RegistryFreshnessStateActive,
		DisciplineID:                     "oss_registry_freshness",
		Version:                          "v0",
		FreshnessState:                   IntelligenceCalibrationFreshnessFresh,
		SupportedRegistryMetadataBounded: true,
		UnsupportedRegistryStateExplicit: true,
		StaleRegistryBlocksActive:        true,
		UnknownFreshnessFailsClosed:      true,
		ProjectionDisclaimer:             ossTrustNetworkVal0ProjectionDisclaimer(),
	}
}

func OSSTrustNetworkVal0SharedVEXTriageModel() OSSTrustNetworkVal0SharedVEXTriageDiscipline {
	return OSSTrustNetworkVal0SharedVEXTriageDiscipline{
		CurrentState:              OSSTrustNetworkVal0SharedVEXStateActive,
		DisciplineID:              "oss_shared_vex_triage",
		Version:                   "v0",
		ReviewState:               OSSTrustNetworkReviewStateReviewed,
		PresentedAsReviewed:       true,
		LocalApplicabilityBounded: true,
		SharedIntelligenceOnly:    true,
		ProjectionDisclaimer:      ossTrustNetworkVal0ProjectionDisclaimer(),
	}
}

func OSSTrustNetworkVal0PropagationModel() OSSTrustNetworkVal0PropagationDiscipline {
	return OSSTrustNetworkVal0PropagationDiscipline{
		CurrentState:            OSSTrustNetworkVal0PropagationStateActive,
		DisciplineID:            "oss_signal_propagation",
		Version:                 "v0",
		ReviewState:             OSSTrustNetworkReviewStateReviewed,
		SourceWeightClass:       OSSTrustNetworkSourceWeightMedium,
		FreshnessState:          IntelligenceCalibrationFreshnessFresh,
		EvidenceRefs:            []string{"evidence:ostn-val0-propagation-001"},
		ApplicabilityScope:      OSSTrustNetworkApplicabilityEnterpriseLocal,
		PropagationScope:        OSSTrustNetworkPropagationReviewedExchange,
		SimilarityContextGating: true,
		Caveats:                 []string{"bounded reviewed exchange only"},
		ProjectionDisclaimer:    ossTrustNetworkVal0ProjectionDisclaimer(),
	}
}

func OSSTrustNetworkVal0LocalApplicabilityModel() OSSTrustNetworkVal0LocalApplicabilityDiscipline {
	return OSSTrustNetworkVal0LocalApplicabilityDiscipline{
		CurrentState:                OSSTrustNetworkVal0LocalApplicabilityStateActive,
		DisciplineID:                "oss_local_applicability",
		Version:                     "v0",
		NetworkSignalAdvisoryOnly:   true,
		LocalApplicabilityExplicit:  true,
		LocalOverrideVisible:        true,
		LocalOverrideEvidenceLinked: true,
		ProjectionDisclaimer:        ossTrustNetworkVal0ProjectionDisclaimer(),
	}
}

func OSSTrustNetworkVal0NoOverclaimModel() OSSTrustNetworkVal0NoOverclaimDiscipline {
	return OSSTrustNetworkVal0NoOverclaimDiscipline{
		CurrentState:         OSSTrustNetworkVal0NoOverclaimStateActive,
		DisciplineID:         "oss_no_overclaim",
		Version:              "v0",
		ObservedClaims:       []string{"bounded OSS trust signal", "reviewed shared intelligence", "advisory network signal"},
		ProjectionDisclaimer: ossTrustNetworkVal0ProjectionDisclaimer(),
	}
}

func OSSTrustNetworkVal0FoundationModel() OSSTrustNetworkVal0Foundation {
	return OSSTrustNetworkVal0Foundation{
		CurrentState:            OSSTrustNetworkVal0StateActive,
		Point9State:             OSSTrustNetworkPoint9StateNotComplete,
		DependencyState:         OSSTrustNetworkVal0DependencyStateActive,
		SignalContractState:     OSSTrustNetworkVal0SignalContractStateActive,
		TrustMarkingState:       OSSTrustNetworkVal0TrustMarkingStateActive,
		MaintainerIdentityState: OSSTrustNetworkVal0MaintainerIdentityStateActive,
		RegistryFreshnessState:  OSSTrustNetworkVal0RegistryFreshnessStateActive,
		SharedVEXState:          OSSTrustNetworkVal0SharedVEXStateActive,
		PropagationState:        OSSTrustNetworkVal0PropagationStateActive,
		LocalApplicabilityState: OSSTrustNetworkVal0LocalApplicabilityStateActive,
		NoOverclaimState:        OSSTrustNetworkVal0NoOverclaimStateActive,
		Dependency:              OSSTrustNetworkVal0DependencySnapshotModel(),
		SignalContract:          OSSTrustNetworkVal0SignalContractModel(),
		TrustMarking:            OSSTrustNetworkVal0TrustMarkingModel(),
		MaintainerIdentity:      OSSTrustNetworkVal0MaintainerIdentityModel(),
		RegistryFreshness:       OSSTrustNetworkVal0RegistryFreshnessModel(),
		SharedVEXTriage:         OSSTrustNetworkVal0SharedVEXTriageModel(),
		Propagation:             OSSTrustNetworkVal0PropagationModel(),
		LocalApplicability:      OSSTrustNetworkVal0LocalApplicabilityModel(),
		NoOverclaim:             OSSTrustNetworkVal0NoOverclaimModel(),
		ProofSurfaceRefs:        OSSTrustNetworkVal0ProofSurfaceRefs(),
		EvidenceRefs:            OSSTrustNetworkVal0ProofEvidenceRefs(),
		WhyPoint9NotComplete: []string{
			"Val 0 establishes the OSS trust discipline foundation only and cannot complete Točka 9.",
			"Registry connectors, signing integrations, reviewed propagation workflows, dashboards, and integrated closure remain for later waves.",
			"OSS trust network outputs remain advisory and cannot become canonical truth, certification, approval, or deployment authority here.",
		},
		ProjectionDisclaimer: ossTrustNetworkVal0ProjectionDisclaimer(),
	}
}

func EvaluateOSSTrustNetworkVal0DependencyState(model OSSTrustNetworkVal0DependencySnapshot) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.CurrentState,
		model.Point8State,
		model.ClosureState,
		model.NoOverclaimState,
		model.ProjectionDisclaimer,
	) || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 {
		return OSSTrustNetworkVal0DependencyStateIncomplete
	}
	if !developerEcosystemValEHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0DependencyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.ProofSurfaceRefs, DeveloperEcosystemValEProofSurfaceRefs()...) ||
		!DeveloperEcosystemValEProofEvidenceQualityValid(developerEcosystemValEEvidence(), model.EvidenceRefs) {
		return OSSTrustNetworkVal0DependencyStateBlocked
	}
	if model.Point8PassAllowed && strings.TrimSpace(model.Point8PassReason) != DeveloperEcosystemValEPoint8PassReasonAllowed {
		return OSSTrustNetworkVal0DependencyStateBlocked
	}
	if strings.TrimSpace(model.Point8State) != DeveloperEcosystemPoint8StatePass ||
		!model.Point8PassAllowed ||
		strings.TrimSpace(model.ClosureState) != DeveloperEcosystemValEClosureStateActive ||
		strings.TrimSpace(model.NoOverclaimState) != DeveloperEcosystemValENoOverclaimStateActive {
		switch strings.TrimSpace(model.CurrentState) {
		case DeveloperEcosystemValEStateIncomplete:
			return OSSTrustNetworkVal0DependencyStateIncomplete
		case DeveloperEcosystemValEStateUnknown:
			return OSSTrustNetworkVal0DependencyStateUnknown
		case DeveloperEcosystemValEStateBlocked:
			return OSSTrustNetworkVal0DependencyStateBlocked
		default:
			return OSSTrustNetworkVal0DependencyStatePartial
		}
	}
	if strings.TrimSpace(model.CurrentState) != DeveloperEcosystemValEStatePass {
		return OSSTrustNetworkVal0DependencyStatePartial
	}
	return OSSTrustNetworkVal0DependencyStateActive
}

func EvaluateOSSTrustNetworkVal0SignalContractState(model OSSTrustNetworkVal0SignalContract) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ContractID,
		model.Version,
		model.SignalID,
		model.PackageOrProjectIdentity,
		model.RegistryOrEcosystem,
		model.PackageVersionOrReleaseRef,
		model.SignalClass,
		model.SourceClass,
		model.SourceWeightClass,
		model.FreshnessState,
		model.ConfidenceClass,
		model.ReviewState,
		model.ApplicabilityScope,
		model.PropagationScope,
		model.ProjectionDisclaimer,
	) || len(model.SupportedReviewStates) == 0 || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkVal0SignalContractStateIncomplete
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0SignalContractStateUnknown
	}
	if !containsExactTrimmedStringSet(model.SupportedReviewStates, ossTrustNetworkVal0ReviewStates()...) {
		return OSSTrustNetworkVal0SignalContractStateBlocked
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-val0-signal-contract-001") {
		return OSSTrustNetworkVal0SignalContractStateBlocked
	}
	if !containsTrimmedString(ossTrustNetworkVal0TrustMarkingClasses(), model.SignalClass) ||
		!containsTrimmedString(ossTrustNetworkVal0SourceClasses(), model.SourceClass) ||
		!containsTrimmedString(ossTrustNetworkVal0SourceWeightClasses(), model.SourceWeightClass) ||
		!containsTrimmedString(ossTrustNetworkVal0ConfidenceClasses(), model.ConfidenceClass) ||
		!containsTrimmedString(ossTrustNetworkVal0ApplicabilityScopes(), model.ApplicabilityScope) ||
		!containsTrimmedString(ossTrustNetworkVal0PropagationScopes(), model.PropagationScope) {
		return OSSTrustNetworkVal0SignalContractStateBlocked
	}
	if !ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) {
		return OSSTrustNetworkVal0SignalContractStateBlocked
	}
	if !containsTrimmedString(ossTrustNetworkVal0ReviewStates(), model.ReviewState) {
		return OSSTrustNetworkVal0SignalContractStateBlocked
	}
	if strings.TrimSpace(model.ReviewState) == OSSTrustNetworkReviewStateCandidate && model.PresentedAsReviewed {
		return OSSTrustNetworkVal0SignalContractStateBlocked
	}
	return OSSTrustNetworkVal0SignalContractStateActive
}

func EvaluateOSSTrustNetworkVal0TrustMarkingState(model OSSTrustNetworkVal0TrustMarkingSemantics) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) ||
		len(model.AllowedTrustMarkingClasses) == 0 {
		return OSSTrustNetworkVal0TrustMarkingStateIncomplete
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0TrustMarkingStateUnknown
	}
	if !containsExactTrimmedStringSet(model.AllowedTrustMarkingClasses, ossTrustNetworkVal0TrustMarkingClasses()...) {
		return OSSTrustNetworkVal0TrustMarkingStateBlocked
	}
	if model.GenericVerifiedBadgeClaim ||
		model.BroadProjectSafetyBadgeClaim ||
		model.IntegrityScoreClaim ||
		model.ScoreGreaterThanNinetyClaim ||
		model.UniversalTrustScoreClaim ||
		model.SafePackageClaim ||
		model.CertifiedPackageClaim ||
		model.OfficialEcosystemApprovalClaim {
		return OSSTrustNetworkVal0TrustMarkingStateBlocked
	}
	return OSSTrustNetworkVal0TrustMarkingStateActive
}

func EvaluateOSSTrustNetworkVal0MaintainerIdentityState(model OSSTrustNetworkVal0MaintainerIdentityDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0MaintainerIdentityStateIncomplete
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0MaintainerIdentityStateUnknown
	}
	if model.EnterpriseApprovalClaim || model.CanonicalTruthClaim {
		return OSSTrustNetworkVal0MaintainerIdentityStateBlocked
	}
	if model.MaintainerIdentityBinding &&
		model.KeyToMaintainerLinkage &&
		model.DelegatedSigningReviewed &&
		model.KeyRotationHandled &&
		model.KeyRevocationHandled &&
		model.CompromiseHandlingVisible &&
		model.TrustContinuityBounded {
		return OSSTrustNetworkVal0MaintainerIdentityStateActive
	}
	return OSSTrustNetworkVal0MaintainerIdentityStateBlocked
}

func EvaluateOSSTrustNetworkVal0RegistryFreshnessState(model OSSTrustNetworkVal0RegistryFreshnessDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.FreshnessState, model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0RegistryFreshnessStateIncomplete
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0RegistryFreshnessStateUnknown
	}
	if model.RegistryMetadataAloneReviewedTrust || model.RegistryHeuristicsGlobalBlocklist {
		return OSSTrustNetworkVal0RegistryFreshnessStateBlocked
	}
	switch strings.TrimSpace(model.FreshnessState) {
	case IntelligenceCalibrationFreshnessFresh:
		if model.SupportedRegistryMetadataBounded &&
			model.UnsupportedRegistryStateExplicit &&
			model.StaleRegistryBlocksActive &&
			model.UnknownFreshnessFailsClosed {
			return OSSTrustNetworkVal0RegistryFreshnessStateActive
		}
		return OSSTrustNetworkVal0RegistryFreshnessStatePartial
	case IntelligenceCalibrationFreshnessUnsupported:
		if model.UnsupportedRegistryStateExplicit {
			return OSSTrustNetworkVal0RegistryFreshnessStatePartial
		}
		return OSSTrustNetworkVal0RegistryFreshnessStateBlocked
	case IntelligenceCalibrationFreshnessUnknown, IntelligenceCalibrationFreshnessStale, IntelligenceCalibrationFreshnessExpired:
		return OSSTrustNetworkVal0RegistryFreshnessStateBlocked
	default:
		return OSSTrustNetworkVal0RegistryFreshnessStateBlocked
	}
}

func EvaluateOSSTrustNetworkVal0SharedVEXState(model OSSTrustNetworkVal0SharedVEXTriageDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ReviewState, model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0SharedVEXStateIncomplete
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0SharedVEXStateUnknown
	}
	if !containsTrimmedString(ossTrustNetworkVal0ReviewStates(), model.ReviewState) {
		return OSSTrustNetworkVal0SharedVEXStateBlocked
	}
	if model.RejectedPropagatedUsable || model.RevokedUsable || !model.LocalApplicabilityBounded || !model.SharedIntelligenceOnly {
		return OSSTrustNetworkVal0SharedVEXStateBlocked
	}
	switch strings.TrimSpace(model.ReviewState) {
	case OSSTrustNetworkReviewStateReviewed:
		if model.PresentedAsReviewed {
			return OSSTrustNetworkVal0SharedVEXStateActive
		}
		return OSSTrustNetworkVal0SharedVEXStatePartial
	case OSSTrustNetworkReviewStateCandidate:
		if model.PresentedAsReviewed {
			return OSSTrustNetworkVal0SharedVEXStateBlocked
		}
		return OSSTrustNetworkVal0SharedVEXStatePartial
	case OSSTrustNetworkReviewStateRejected:
		return OSSTrustNetworkVal0SharedVEXStatePartial
	case OSSTrustNetworkReviewStateSuperseded:
		if strings.TrimSpace(model.ReplacementRef) == "" {
			return OSSTrustNetworkVal0SharedVEXStateBlocked
		}
		return OSSTrustNetworkVal0SharedVEXStatePartial
	case OSSTrustNetworkReviewStateRevoked:
		return OSSTrustNetworkVal0SharedVEXStateBlocked
	default:
		return OSSTrustNetworkVal0SharedVEXStateBlocked
	}
}

func EvaluateOSSTrustNetworkVal0PropagationState(model OSSTrustNetworkVal0PropagationDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.ReviewState,
		model.SourceWeightClass,
		model.FreshnessState,
		model.ApplicabilityScope,
		model.PropagationScope,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkVal0PropagationStateIncomplete
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0PropagationStateUnknown
	}
	if !containsTrimmedString(ossTrustNetworkVal0ReviewStates(), model.ReviewState) ||
		!containsTrimmedString(ossTrustNetworkVal0SourceWeightClasses(), model.SourceWeightClass) ||
		!containsTrimmedString(ossTrustNetworkVal0ApplicabilityScopes(), model.ApplicabilityScope) ||
		!containsTrimmedString(ossTrustNetworkVal0PropagationScopes(), model.PropagationScope) ||
		!containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-val0-propagation-001") {
		return OSSTrustNetworkVal0PropagationStateBlocked
	}
	if model.AutomaticGlobalSpread ||
		model.CandidateTreatedAsReviewed ||
		model.RegistryHeuristicAsCanonicalProof ||
		model.CommunityReportAsFinalTruth ||
		model.OverridesLocalEnterpriseWithoutReview ||
		model.CommunityCandidateCanonicalTruth {
		return OSSTrustNetworkVal0PropagationStateBlocked
	}
	if !ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) {
		return OSSTrustNetworkVal0PropagationStateBlocked
	}
	if !model.SimilarityContextGating {
		return OSSTrustNetworkVal0PropagationStateBlocked
	}
	if strings.TrimSpace(model.ReviewState) == OSSTrustNetworkReviewStateReviewed && strings.TrimSpace(model.FreshnessState) == IntelligenceCalibrationFreshnessFresh {
		return OSSTrustNetworkVal0PropagationStateActive
	}
	return OSSTrustNetworkVal0PropagationStatePartial
}

func EvaluateOSSTrustNetworkVal0LocalApplicabilityState(model OSSTrustNetworkVal0LocalApplicabilityDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0LocalApplicabilityStateIncomplete
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0LocalApplicabilityStateUnknown
	}
	if model.CommunitySignalRewritesEnterprise {
		return OSSTrustNetworkVal0LocalApplicabilityStateBlocked
	}
	if model.NetworkSignalAdvisoryOnly &&
		model.LocalApplicabilityExplicit &&
		model.LocalOverrideVisible &&
		model.LocalOverrideEvidenceLinked {
		return OSSTrustNetworkVal0LocalApplicabilityStateActive
	}
	return OSSTrustNetworkVal0LocalApplicabilityStatePartial
}

func EvaluateOSSTrustNetworkVal0NoOverclaimState(model OSSTrustNetworkVal0NoOverclaimDiscipline) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0NoOverclaimStateIncomplete
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0NoOverclaimStateUnknown
	}
	if model.GlobalTruthClaim ||
		model.CertifiedClaim ||
		model.RegulatorApproved ||
		model.ProductionApproved ||
		model.LegalCertification ||
		model.PatentCleared ||
		model.FTOPCleared ||
		ossTrustNetworkVal0ContainsForbiddenClaim(model.ObservedClaims...) {
		return OSSTrustNetworkVal0NoOverclaimStateBlocked
	}
	return OSSTrustNetworkVal0NoOverclaimStateActive
}

func EvaluateOSSTrustNetworkVal0State(model OSSTrustNetworkVal0Foundation) string {
	if strings.TrimSpace(model.Point9State) == "" || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return OSSTrustNetworkVal0StateIncomplete
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkVal0StateUnknown
	}
	if strings.TrimSpace(model.Point9State) != OSSTrustNetworkPoint9StateNotComplete ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkVal0ProofSurfaceRefs()...) ||
		!OSSTrustNetworkVal0ProofEvidenceQualityValid(ossTrustNetworkVal0Evidence(), model.EvidenceRefs) {
		return OSSTrustNetworkVal0StateBlocked
	}
	states := []string{
		model.DependencyState,
		model.SignalContractState,
		model.TrustMarkingState,
		model.MaintainerIdentityState,
		model.RegistryFreshnessState,
		model.SharedVEXState,
		model.PropagationState,
		model.LocalApplicabilityState,
		model.NoOverclaimState,
	}
	allActive := true
	for _, state := range states {
		if strings.TrimSpace(state) == "" {
			return OSSTrustNetworkVal0StateIncomplete
		}
		if !strings.HasSuffix(strings.TrimSpace(state), "_active") {
			allActive = false
		}
	}
	if allActive {
		return OSSTrustNetworkVal0StateActive
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_blocked") {
			return OSSTrustNetworkVal0StateBlocked
		}
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_incomplete") {
			return OSSTrustNetworkVal0StateIncomplete
		}
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_unknown") {
			return OSSTrustNetworkVal0StateUnknown
		}
	}
	return OSSTrustNetworkVal0StatePartial
}

func EvaluateOSSTrustNetworkPoint9State(currentState string) string {
	_ = currentState
	return OSSTrustNetworkPoint9StateNotComplete
}

func EvaluateOSSTrustNetworkVal0ProofsState(model OSSTrustNetworkVal0Foundation, limitations []string) string {
	baseState := strings.TrimSpace(model.CurrentState)
	if baseState == "" {
		baseState = OSSTrustNetworkVal0StateUnknown
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkVal0ProofSurfaceRefs()...) ||
		!OSSTrustNetworkVal0ProofEvidenceQualityValid(ossTrustNetworkVal0Evidence(), model.EvidenceRefs) ||
		len(limitations) == 0 ||
		strings.TrimSpace(model.Point9State) != OSSTrustNetworkPoint9StateNotComplete {
		if baseState == OSSTrustNetworkVal0StateActive {
			return OSSTrustNetworkVal0StatePartial
		}
		return baseState
	}
	return baseState
}

func ossTrustNetworkVal0BlockingReasons(model OSSTrustNetworkVal0Foundation) []string {
	reasons := []string{}
	if model.DependencyState != OSSTrustNetworkVal0DependencyStateActive {
		reasons = append(reasons, "Točka 8 / Val E integrated closure dependency is not active and exact.")
	}
	if model.SignalContractState != OSSTrustNetworkVal0SignalContractStateActive {
		reasons = append(reasons, "OSS signal contract discipline is not exact, review-safe, and evidence-linked.")
	}
	if model.TrustMarkingState != OSSTrustNetworkVal0TrustMarkingStateActive {
		reasons = append(reasons, "Trust marking semantics exceed bounded classes or imply broad verification.")
	}
	if model.MaintainerIdentityState != OSSTrustNetworkVal0MaintainerIdentityStateActive {
		reasons = append(reasons, "Maintainer identity lifecycle discipline is incomplete or overclaiming.")
	}
	if model.RegistryFreshnessState != OSSTrustNetworkVal0RegistryFreshnessStateActive {
		reasons = append(reasons, "Registry freshness or unsupported-state discipline is not active and fail-closed.")
	}
	if model.SharedVEXState != OSSTrustNetworkVal0SharedVEXStateActive {
		reasons = append(reasons, "Shared VEX / triage review discipline is not active and bounded.")
	}
	if model.PropagationState != OSSTrustNetworkVal0PropagationStateActive {
		reasons = append(reasons, "Propagation discipline is not active and bounded by review, freshness, and context.")
	}
	if model.LocalApplicabilityState != OSSTrustNetworkVal0LocalApplicabilityStateActive {
		reasons = append(reasons, "Local enterprise applicability boundary is not active and visible.")
	}
	if model.NoOverclaimState != OSSTrustNetworkVal0NoOverclaimStateActive {
		reasons = append(reasons, "OSTN no-overclaim and no-global-truth guard is not active.")
	}
	return developerEcosystemValECollectText(reasons)
}

func ComputeOSSTrustNetworkVal0Foundation(model OSSTrustNetworkVal0Foundation) OSSTrustNetworkVal0Foundation {
	model.DependencyState = EvaluateOSSTrustNetworkVal0DependencyState(model.Dependency)
	model.SignalContractState = EvaluateOSSTrustNetworkVal0SignalContractState(model.SignalContract)
	model.TrustMarkingState = EvaluateOSSTrustNetworkVal0TrustMarkingState(model.TrustMarking)
	model.MaintainerIdentityState = EvaluateOSSTrustNetworkVal0MaintainerIdentityState(model.MaintainerIdentity)
	model.RegistryFreshnessState = EvaluateOSSTrustNetworkVal0RegistryFreshnessState(model.RegistryFreshness)
	model.SharedVEXState = EvaluateOSSTrustNetworkVal0SharedVEXState(model.SharedVEXTriage)
	model.PropagationState = EvaluateOSSTrustNetworkVal0PropagationState(model.Propagation)
	model.LocalApplicabilityState = EvaluateOSSTrustNetworkVal0LocalApplicabilityState(model.LocalApplicability)
	model.NoOverclaimState = EvaluateOSSTrustNetworkVal0NoOverclaimState(model.NoOverclaim)
	model.Point9State = EvaluateOSSTrustNetworkPoint9State(model.CurrentState)
	model.CurrentState = EvaluateOSSTrustNetworkVal0State(model)
	model.Point9State = EvaluateOSSTrustNetworkPoint9State(model.CurrentState)
	model.BlockingReasons = ossTrustNetworkVal0BlockingReasons(model)

	model.SignalContract.CurrentState = model.SignalContractState
	model.TrustMarking.CurrentState = model.TrustMarkingState
	model.MaintainerIdentity.CurrentState = model.MaintainerIdentityState
	model.RegistryFreshness.CurrentState = model.RegistryFreshnessState
	model.SharedVEXTriage.CurrentState = model.SharedVEXState
	model.Propagation.CurrentState = model.PropagationState
	model.LocalApplicability.CurrentState = model.LocalApplicabilityState
	model.NoOverclaim.CurrentState = model.NoOverclaimState

	return model
}

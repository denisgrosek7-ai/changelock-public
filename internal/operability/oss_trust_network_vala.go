package operability

import "strings"

const (
	OSSTrustNetworkValAStateActive     = "oss_trust_network_vala_active"
	OSSTrustNetworkValAStatePartial    = "oss_trust_network_vala_partial"
	OSSTrustNetworkValAStateIncomplete = "oss_trust_network_vala_incomplete"
	OSSTrustNetworkValAStateBlocked    = "oss_trust_network_vala_blocked"
	OSSTrustNetworkValAStateUnknown    = "oss_trust_network_vala_unknown"

	OSSTrustNetworkValADependencyStateActive     = "oss_trust_network_vala_dependency_active"
	OSSTrustNetworkValADependencyStatePartial    = "oss_trust_network_vala_dependency_partial"
	OSSTrustNetworkValADependencyStateIncomplete = "oss_trust_network_vala_dependency_incomplete"
	OSSTrustNetworkValADependencyStateBlocked    = "oss_trust_network_vala_dependency_blocked"
	OSSTrustNetworkValADependencyStateUnknown    = "oss_trust_network_vala_dependency_unknown"

	OSSTrustNetworkValAReleaseTrustIntakeStateActive     = "oss_trust_network_vala_release_trust_intake_active"
	OSSTrustNetworkValAReleaseTrustIntakeStatePartial    = "oss_trust_network_vala_release_trust_intake_partial"
	OSSTrustNetworkValAReleaseTrustIntakeStateIncomplete = "oss_trust_network_vala_release_trust_intake_incomplete"
	OSSTrustNetworkValAReleaseTrustIntakeStateBlocked    = "oss_trust_network_vala_release_trust_intake_blocked"
	OSSTrustNetworkValAReleaseTrustIntakeStateUnknown    = "oss_trust_network_vala_release_trust_intake_unknown"

	OSSTrustNetworkValASigningSignalStateActive     = "oss_trust_network_vala_signing_signal_active"
	OSSTrustNetworkValASigningSignalStatePartial    = "oss_trust_network_vala_signing_signal_partial"
	OSSTrustNetworkValASigningSignalStateIncomplete = "oss_trust_network_vala_signing_signal_incomplete"
	OSSTrustNetworkValASigningSignalStateBlocked    = "oss_trust_network_vala_signing_signal_blocked"
	OSSTrustNetworkValASigningSignalStateUnknown    = "oss_trust_network_vala_signing_signal_unknown"

	OSSTrustNetworkValAMaintainerAttestationStateActive     = "oss_trust_network_vala_maintainer_attestation_active"
	OSSTrustNetworkValAMaintainerAttestationStatePartial    = "oss_trust_network_vala_maintainer_attestation_partial"
	OSSTrustNetworkValAMaintainerAttestationStateIncomplete = "oss_trust_network_vala_maintainer_attestation_incomplete"
	OSSTrustNetworkValAMaintainerAttestationStateBlocked    = "oss_trust_network_vala_maintainer_attestation_blocked"
	OSSTrustNetworkValAMaintainerAttestationStateUnknown    = "oss_trust_network_vala_maintainer_attestation_unknown"

	OSSTrustNetworkValAProvenanceMaterialStateActive     = "oss_trust_network_vala_provenance_material_active"
	OSSTrustNetworkValAProvenanceMaterialStatePartial    = "oss_trust_network_vala_provenance_material_partial"
	OSSTrustNetworkValAProvenanceMaterialStateIncomplete = "oss_trust_network_vala_provenance_material_incomplete"
	OSSTrustNetworkValAProvenanceMaterialStateBlocked    = "oss_trust_network_vala_provenance_material_blocked"
	OSSTrustNetworkValAProvenanceMaterialStateUnknown    = "oss_trust_network_vala_provenance_material_unknown"

	OSSTrustNetworkValARegistryDescriptorStateActive     = "oss_trust_network_vala_registry_descriptor_active"
	OSSTrustNetworkValARegistryDescriptorStatePartial    = "oss_trust_network_vala_registry_descriptor_partial"
	OSSTrustNetworkValARegistryDescriptorStateIncomplete = "oss_trust_network_vala_registry_descriptor_incomplete"
	OSSTrustNetworkValARegistryDescriptorStateBlocked    = "oss_trust_network_vala_registry_descriptor_blocked"
	OSSTrustNetworkValARegistryDescriptorStateUnknown    = "oss_trust_network_vala_registry_descriptor_unknown"

	OSSTrustNetworkValARegistryMetadataStateActive     = "oss_trust_network_vala_registry_metadata_active"
	OSSTrustNetworkValARegistryMetadataStatePartial    = "oss_trust_network_vala_registry_metadata_partial"
	OSSTrustNetworkValARegistryMetadataStateIncomplete = "oss_trust_network_vala_registry_metadata_incomplete"
	OSSTrustNetworkValARegistryMetadataStateBlocked    = "oss_trust_network_vala_registry_metadata_blocked"
	OSSTrustNetworkValARegistryMetadataStateUnknown    = "oss_trust_network_vala_registry_metadata_unknown"

	OSSTrustNetworkValATypoSquattingWarningStateActive     = "oss_trust_network_vala_typo_squatting_warning_active"
	OSSTrustNetworkValATypoSquattingWarningStatePartial    = "oss_trust_network_vala_typo_squatting_warning_partial"
	OSSTrustNetworkValATypoSquattingWarningStateIncomplete = "oss_trust_network_vala_typo_squatting_warning_incomplete"
	OSSTrustNetworkValATypoSquattingWarningStateBlocked    = "oss_trust_network_vala_typo_squatting_warning_blocked"
	OSSTrustNetworkValATypoSquattingWarningStateUnknown    = "oss_trust_network_vala_typo_squatting_warning_unknown"

	OSSTrustNetworkValADriftSignalStateActive     = "oss_trust_network_vala_drift_signal_active"
	OSSTrustNetworkValADriftSignalStatePartial    = "oss_trust_network_vala_drift_signal_partial"
	OSSTrustNetworkValADriftSignalStateIncomplete = "oss_trust_network_vala_drift_signal_incomplete"
	OSSTrustNetworkValADriftSignalStateBlocked    = "oss_trust_network_vala_drift_signal_blocked"
	OSSTrustNetworkValADriftSignalStateUnknown    = "oss_trust_network_vala_drift_signal_unknown"

	OSSTrustNetworkValANoOverclaimStateActive     = "oss_trust_network_vala_no_overclaim_active"
	OSSTrustNetworkValANoOverclaimStatePartial    = "oss_trust_network_vala_no_overclaim_partial"
	OSSTrustNetworkValANoOverclaimStateIncomplete = "oss_trust_network_vala_no_overclaim_incomplete"
	OSSTrustNetworkValANoOverclaimStateBlocked    = "oss_trust_network_vala_no_overclaim_blocked"
	OSSTrustNetworkValANoOverclaimStateUnknown    = "oss_trust_network_vala_no_overclaim_unknown"

	OSSTrustNetworkValASigningStatePresent     = "present"
	OSSTrustNetworkValASigningStateVerified    = "verified"
	OSSTrustNetworkValASigningStateMissing     = "missing"
	OSSTrustNetworkValASigningStateMismatch    = "mismatch"
	OSSTrustNetworkValASigningStateRevoked     = "revoked"
	OSSTrustNetworkValASigningStateUnsupported = "unsupported"
	OSSTrustNetworkValASigningStateUnknown     = "unknown"

	OSSTrustNetworkValAAttestationStateAttested    = "attested"
	OSSTrustNetworkValAAttestationStateMissing     = "missing"
	OSSTrustNetworkValAAttestationStateStale       = "stale"
	OSSTrustNetworkValAAttestationStateRevoked     = "revoked"
	OSSTrustNetworkValAAttestationStateDelegated   = "delegated"
	OSSTrustNetworkValAAttestationStateUnsupported = "unsupported"
	OSSTrustNetworkValAAttestationStateUnknown     = "unknown"

	OSSTrustNetworkValAProvenanceStateVerified          = "verified"
	OSSTrustNetworkValAProvenanceStatePresentUnverified = "present_unverified"
	OSSTrustNetworkValAProvenanceStateMissing           = "missing"
	OSSTrustNetworkValAProvenanceStateMismatch          = "mismatch"
	OSSTrustNetworkValAProvenanceStateStale             = "stale"
	OSSTrustNetworkValAProvenanceStateUnsupported       = "unsupported"
	OSSTrustNetworkValAProvenanceStateUnknown           = "unknown"

	OSSTrustNetworkValARegistryDescriptorNPM                = "npm"
	OSSTrustNetworkValARegistryDescriptorPyPI               = "pypi"
	OSSTrustNetworkValARegistryDescriptorMaven              = "maven"
	OSSTrustNetworkValARegistryDescriptorGitHubReleases     = "github_releases"
	OSSTrustNetworkValARegistryDescriptorGitLabReleases     = "gitlab_releases"
	OSSTrustNetworkValARegistryDescriptorGenericOSSRegistry = "generic_oss_registry"
	OSSTrustNetworkValARegistryDescriptorUnsupported        = "unsupported"

	OSSTrustNetworkValADriftClassMaintainer        = "maintainer_drift"
	OSSTrustNetworkValADriftClassProvenance        = "provenance_drift"
	OSSTrustNetworkValADriftClassSigning           = "signing_drift"
	OSSTrustNetworkValADriftClassRegistryMetadata  = "registry_metadata_drift"
	OSSTrustNetworkValADriftClassSuspiciousRelease = "suspicious_release_delta"

	OSSTrustNetworkValADriftStateCandidate   = "candidate"
	OSSTrustNetworkValADriftStateReviewed    = "reviewed"
	OSSTrustNetworkValADriftStateUnsupported = "unsupported"
	OSSTrustNetworkValADriftStateUnknown     = "unknown"
)

type OSSTrustNetworkValADependencySnapshot struct {
	Val0CurrentState         string   `json:"val0_current_state"`
	Val0Point9State          string   `json:"val0_point_9_state"`
	Val0DependencyState      string   `json:"val0_dependency_state"`
	Val0NoOverclaimState     string   `json:"val0_no_overclaim_state"`
	Val0Point8State          string   `json:"val0_point_8_state"`
	Val0Point8PassAllowed    bool     `json:"val0_point_8_pass_allowed"`
	Val0Point8PassReason     string   `json:"val0_point_8_pass_reason"`
	Val0Point8ClosureState   string   `json:"val0_point_8_closure_state"`
	Val0ProofSurfaceRefs     []string `json:"val0_proof_surface_refs,omitempty"`
	Val0EvidenceRefs         []string `json:"val0_evidence_refs,omitempty"`
	Val0ProjectionDisclaimer string   `json:"val0_projection_disclaimer"`
}

type OSSTrustNetworkValAReleaseTrustIntake struct {
	CurrentState               string   `json:"current_state"`
	ReleaseTrustProfileID      string   `json:"release_trust_profile_id"`
	PackageOrProjectIdentity   string   `json:"package_or_project_identity"`
	RegistryOrEcosystem        string   `json:"registry_or_ecosystem"`
	ReleaseRef                 string   `json:"release_ref"`
	ArtifactIdentity           string   `json:"artifact_identity"`
	SigningSignalState         string   `json:"signing_signal_state"`
	ProvenanceSignalState      string   `json:"provenance_signal_state"`
	MaintainerAttestationState string   `json:"maintainer_attestation_state"`
	RegistryMetadataState      string   `json:"registry_metadata_state"`
	EvidenceRefs               []string `json:"evidence_refs,omitempty"`
	FreshnessState             string   `json:"freshness_state"`
	ReviewState                string   `json:"review_state"`
	Caveats                    []string `json:"caveats,omitempty"`
	PackageSafetyClaim         bool     `json:"package_safety_claim"`
	CertifiedPackageClaim      bool     `json:"certified_package_claim"`
	ProductionApprovalClaim    bool     `json:"production_approval_claim"`
	GlobalApprovalClaim        bool     `json:"global_approval_claim"`
	ProjectionDisclaimer       string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValASigningSignal struct {
	CurrentState                  string   `json:"current_state"`
	DisciplineID                  string   `json:"discipline_id"`
	Version                       string   `json:"version"`
	SigningState                  string   `json:"signing_state"`
	ReleaseRef                    string   `json:"release_ref"`
	ArtifactIdentity              string   `json:"artifact_identity"`
	EvidenceRefs                  []string `json:"evidence_refs,omitempty"`
	Caveats                       []string `json:"caveats,omitempty"`
	VerifiedEvidenceLinked        bool     `json:"verified_evidence_linked"`
	ScopedReleaseArtifactIdentity bool     `json:"scoped_release_artifact_identity"`
	MaintainerApprovalClaim       bool     `json:"maintainer_approval_claim"`
	PackageCertificationClaim     bool     `json:"package_certification_claim"`
	CanonicalTruthClaim           bool     `json:"canonical_truth_claim"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValAMaintainerAttestation struct {
	CurrentState              string   `json:"current_state"`
	DisciplineID              string   `json:"discipline_id"`
	Version                   string   `json:"version"`
	AttestationState          string   `json:"attestation_state"`
	PackageOrProjectIdentity  string   `json:"package_or_project_identity"`
	ReleaseRef                string   `json:"release_ref"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	Caveats                   []string `json:"caveats,omitempty"`
	IdentityBindingVisible    bool     `json:"identity_binding_visible"`
	KeyLinkageVisible         bool     `json:"key_linkage_visible"`
	DelegationHandlingVisible bool     `json:"delegation_handling_visible"`
	DelegatedSigningReviewed  bool     `json:"delegated_signing_reviewed"`
	KeyRotationHandled        bool     `json:"key_rotation_handled"`
	KeyRevocationHandled      bool     `json:"key_revocation_handled"`
	CompromiseHandlingVisible bool     `json:"compromise_handling_visible"`
	OverridesLocalEnterprise  bool     `json:"overrides_local_enterprise"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValAProvenanceMaterial struct {
	CurrentState                  string   `json:"current_state"`
	DisciplineID                  string   `json:"discipline_id"`
	Version                       string   `json:"version"`
	ProvenanceState               string   `json:"provenance_state"`
	ReleaseRef                    string   `json:"release_ref"`
	ArtifactIdentity              string   `json:"artifact_identity"`
	EvidenceRefs                  []string `json:"evidence_refs,omitempty"`
	Caveats                       []string `json:"caveats,omitempty"`
	ScopedReleaseArtifactIdentity bool     `json:"scoped_release_artifact_identity"`
	RegulatorApprovalClaim        bool     `json:"regulator_approval_claim"`
	CertifiedPackageClaim         bool     `json:"certified_package_claim"`
	LegalClearanceClaim           bool     `json:"legal_clearance_claim"`
	ProductionApprovalClaim       bool     `json:"production_approval_claim"`
	ProjectionDisclaimer          string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValARegistryDescriptor struct {
	CurrentState                         string   `json:"current_state"`
	DisciplineID                         string   `json:"discipline_id"`
	Version                              string   `json:"version"`
	SupportedRegistryDescriptors         []string `json:"supported_registry_descriptors,omitempty"`
	RequestedRegistryDescriptor          string   `json:"requested_registry_descriptor"`
	UnsupportedRegistryExplicit          bool     `json:"unsupported_registry_explicit"`
	DescriptorOnlyBehavior               bool     `json:"descriptor_only_behavior"`
	LiveNetworkFetcher                   bool     `json:"live_network_fetcher"`
	RegistryMetadataCreatesReviewedTrust bool     `json:"registry_metadata_creates_reviewed_trust"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValARegistryMetadata struct {
	CurrentState                           string   `json:"current_state"`
	DisciplineID                           string   `json:"discipline_id"`
	Version                                string   `json:"version"`
	PackageName                            string   `json:"package_name"`
	Registry                               string   `json:"registry"`
	VersionOrRelease                       string   `json:"version_or_release"`
	PublisherOrMaintainerRef               string   `json:"publisher_or_maintainer_ref"`
	ReleaseTimestamp                       string   `json:"release_timestamp"`
	ArtifactRef                            string   `json:"artifact_ref"`
	SourceRepositoryRef                    string   `json:"source_repository_ref"`
	MetadataFreshness                      string   `json:"metadata_freshness"`
	EvidenceRefs                           []string `json:"evidence_refs,omitempty"`
	Caveats                                []string `json:"caveats,omitempty"`
	RegistryMetadataCreatesReviewedTrust   bool     `json:"registry_metadata_creates_reviewed_trust"`
	RegistryMetadataCreatesGlobalBlocklist bool     `json:"registry_metadata_creates_global_blocklist"`
	ProjectionDisclaimer                   string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValATypoSquattingWarning struct {
	CurrentState                 string   `json:"current_state"`
	DisciplineID                 string   `json:"discipline_id"`
	Version                      string   `json:"version"`
	WarningID                    string   `json:"warning_id"`
	SimilarityBasis              string   `json:"similarity_basis"`
	Scope                        string   `json:"scope"`
	ReviewState                  string   `json:"review_state"`
	EvidenceRefs                 []string `json:"evidence_refs,omitempty"`
	Caveats                      []string `json:"caveats,omitempty"`
	FalsePositiveHandlingVisible bool     `json:"false_positive_handling_visible"`
	AutomaticGlobalBlock         bool     `json:"automatic_global_block"`
	CanonicalTruthClaim          bool     `json:"canonical_truth_claim"`
	CandidatePromotedToReviewed  bool     `json:"candidate_promoted_to_reviewed"`
	ProjectionDisclaimer         string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValADriftSignal struct {
	CurrentState             string   `json:"current_state"`
	DisciplineID             string   `json:"discipline_id"`
	Version                  string   `json:"version"`
	SignalID                 string   `json:"signal_id"`
	DriftClass               string   `json:"drift_class"`
	DriftState               string   `json:"drift_state"`
	SourceWeightClass        string   `json:"source_weight_class"`
	ApplicabilityScope       string   `json:"applicability_scope"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	Caveats                  []string `json:"caveats,omitempty"`
	AutomaticGlobalBlock     bool     `json:"automatic_global_block"`
	GlobalTruthClaim         bool     `json:"global_truth_claim"`
	OverridesLocalEnterprise bool     `json:"overrides_local_enterprise"`
	ProjectionDisclaimer     string   `json:"projection_disclaimer"`
}

type OSSTrustNetworkValANoOverclaim struct {
	CurrentState               string   `json:"current_state"`
	DisciplineID               string   `json:"discipline_id"`
	Version                    string   `json:"version"`
	ObservedClaims             []string `json:"observed_claims,omitempty"`
	GlobalTruthClaim           bool     `json:"global_truth_claim"`
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

type OSSTrustNetworkValACore struct {
	CurrentState               string                                   `json:"current_state"`
	Point9State                string                                   `json:"point_9_state"`
	DependencyState            string                                   `json:"dependency_state"`
	ReleaseTrustIntakeState    string                                   `json:"release_trust_intake_state"`
	SigningSignalState         string                                   `json:"signing_signal_state"`
	MaintainerAttestationState string                                   `json:"maintainer_attestation_state"`
	ProvenanceMaterialState    string                                   `json:"provenance_material_state"`
	RegistryDescriptorState    string                                   `json:"registry_descriptor_state"`
	RegistryMetadataState      string                                   `json:"registry_metadata_state"`
	TypoSquattingWarningState  string                                   `json:"typo_squatting_warning_state"`
	DriftSignalState           string                                   `json:"drift_signal_state"`
	NoOverclaimState           string                                   `json:"no_overclaim_state"`
	Dependency                 OSSTrustNetworkValADependencySnapshot    `json:"dependency"`
	ReleaseTrustIntake         OSSTrustNetworkValAReleaseTrustIntake    `json:"release_trust_intake"`
	SigningSignal              OSSTrustNetworkValASigningSignal         `json:"signing_signal"`
	MaintainerAttestation      OSSTrustNetworkValAMaintainerAttestation `json:"maintainer_attestation"`
	ProvenanceMaterial         OSSTrustNetworkValAProvenanceMaterial    `json:"provenance_material"`
	RegistryDescriptor         OSSTrustNetworkValARegistryDescriptor    `json:"registry_descriptor"`
	RegistryMetadata           OSSTrustNetworkValARegistryMetadata      `json:"registry_metadata"`
	TypoSquattingWarning       OSSTrustNetworkValATypoSquattingWarning  `json:"typo_squatting_warning"`
	DriftSignal                OSSTrustNetworkValADriftSignal           `json:"drift_signal"`
	NoOverclaim                OSSTrustNetworkValANoOverclaim           `json:"no_overclaim"`
	ProofSurfaceRefs           []string                                 `json:"proof_surface_refs,omitempty"`
	EvidenceRefs               []string                                 `json:"evidence_refs,omitempty"`
	BlockingReasons            []string                                 `json:"blocking_reasons,omitempty"`
	WhyPoint9NotComplete       []string                                 `json:"why_point_9_not_complete,omitempty"`
	ProjectionDisclaimer       string                                   `json:"projection_disclaimer"`
}

type ossTrustNetworkValAExpectedEvidenceMetadata struct {
	EvidenceType string
	Source       string
	Scope        string
}

func ossTrustNetworkValAProjectionDisclaimer() string {
	return "projection_only not_canonical_truth oss_trust_network_vala advisory_projection release_trust_registry_core"
}

func ossTrustNetworkValAHasProjectionDisclaimer(value string) bool {
	normalized := strings.ToLower(strings.TrimSpace(value))
	return strings.Contains(normalized, "projection_only") &&
		strings.Contains(normalized, "not_canonical_truth") &&
		strings.Contains(normalized, "oss_trust_network_vala")
}

func ossTrustNetworkValAReviewStates() []string {
	return []string{
		OSSTrustNetworkReviewStateCandidate,
		OSSTrustNetworkReviewStateReviewed,
		OSSTrustNetworkReviewStateRejected,
		OSSTrustNetworkReviewStateSuperseded,
		OSSTrustNetworkReviewStateRevoked,
	}
}

func ossTrustNetworkValASigningStates() []string {
	return []string{
		OSSTrustNetworkValASigningStatePresent,
		OSSTrustNetworkValASigningStateVerified,
		OSSTrustNetworkValASigningStateMissing,
		OSSTrustNetworkValASigningStateMismatch,
		OSSTrustNetworkValASigningStateRevoked,
		OSSTrustNetworkValASigningStateUnsupported,
		OSSTrustNetworkValASigningStateUnknown,
	}
}

func ossTrustNetworkValAAttestationStates() []string {
	return []string{
		OSSTrustNetworkValAAttestationStateAttested,
		OSSTrustNetworkValAAttestationStateMissing,
		OSSTrustNetworkValAAttestationStateStale,
		OSSTrustNetworkValAAttestationStateRevoked,
		OSSTrustNetworkValAAttestationStateDelegated,
		OSSTrustNetworkValAAttestationStateUnsupported,
		OSSTrustNetworkValAAttestationStateUnknown,
	}
}

func ossTrustNetworkValAProvenanceStates() []string {
	return []string{
		OSSTrustNetworkValAProvenanceStateVerified,
		OSSTrustNetworkValAProvenanceStatePresentUnverified,
		OSSTrustNetworkValAProvenanceStateMissing,
		OSSTrustNetworkValAProvenanceStateMismatch,
		OSSTrustNetworkValAProvenanceStateStale,
		OSSTrustNetworkValAProvenanceStateUnsupported,
		OSSTrustNetworkValAProvenanceStateUnknown,
	}
}

func ossTrustNetworkValARegistryDescriptorClasses() []string {
	return []string{
		OSSTrustNetworkValARegistryDescriptorNPM,
		OSSTrustNetworkValARegistryDescriptorPyPI,
		OSSTrustNetworkValARegistryDescriptorMaven,
		OSSTrustNetworkValARegistryDescriptorGitHubReleases,
		OSSTrustNetworkValARegistryDescriptorGitLabReleases,
		OSSTrustNetworkValARegistryDescriptorGenericOSSRegistry,
	}
}

func ossTrustNetworkValADriftClasses() []string {
	return []string{
		OSSTrustNetworkValADriftClassMaintainer,
		OSSTrustNetworkValADriftClassProvenance,
		OSSTrustNetworkValADriftClassSigning,
		OSSTrustNetworkValADriftClassRegistryMetadata,
		OSSTrustNetworkValADriftClassSuspiciousRelease,
	}
}

func ossTrustNetworkValADriftStates() []string {
	return []string{
		OSSTrustNetworkValADriftStateCandidate,
		OSSTrustNetworkValADriftStateReviewed,
		OSSTrustNetworkValADriftStateUnsupported,
		OSSTrustNetworkValADriftStateUnknown,
	}
}

func OSSTrustNetworkValAProofSurfaceRefs() []string {
	return []string{
		"/v1/oss-trust-network/val0/status",
		"/v1/oss-trust-network/val0/proofs",
		"/v1/oss-trust-network/vala/status",
		"/v1/oss-trust-network/vala/proofs",
	}
}

func OSSTrustNetworkValAProofEvidenceRefs() []string {
	return []string{
		"evidence:ostn-vala-dependency-001",
		"evidence:ostn-vala-release-trust-intake-001",
		"evidence:ostn-vala-signing-signal-001",
		"evidence:ostn-vala-maintainer-attestation-001",
		"evidence:ostn-vala-provenance-material-001",
		"evidence:ostn-vala-registry-descriptor-001",
		"evidence:ostn-vala-registry-metadata-001",
		"evidence:ostn-vala-typo-warning-001",
		"evidence:ostn-vala-drift-signal-001",
		"evidence:ostn-vala-no-overclaim-001",
		"evidence:ostn-vala-canonical-boundary-001",
		"evidence:ostn-vala-point9-governance-001",
	}
}

func ossTrustNetworkValAExpectedEvidenceMetadataByID() map[string]ossTrustNetworkValAExpectedEvidenceMetadata {
	return map[string]ossTrustNetworkValAExpectedEvidenceMetadata{
		"evidence:ostn-vala-dependency-001":             {EvidenceType: "dependency_state", Source: "oss-trust-network/vala/dependency", Scope: "val0_dependency"},
		"evidence:ostn-vala-release-trust-intake-001":   {EvidenceType: "release_trust_intake", Source: "oss-trust-network/vala/release-trust-intake", Scope: "release_trust_intake"},
		"evidence:ostn-vala-signing-signal-001":         {EvidenceType: "signing_signal", Source: "oss-trust-network/vala/signing-signal", Scope: "signing_signal"},
		"evidence:ostn-vala-maintainer-attestation-001": {EvidenceType: "maintainer_attestation", Source: "oss-trust-network/vala/maintainer-attestation", Scope: "maintainer_attestation"},
		"evidence:ostn-vala-provenance-material-001":    {EvidenceType: "provenance_material", Source: "oss-trust-network/vala/provenance-material", Scope: "provenance_material"},
		"evidence:ostn-vala-registry-descriptor-001":    {EvidenceType: "registry_descriptor", Source: "oss-trust-network/vala/registry-descriptor", Scope: "registry_descriptor"},
		"evidence:ostn-vala-registry-metadata-001":      {EvidenceType: "registry_metadata", Source: "oss-trust-network/vala/registry-metadata", Scope: "registry_metadata_normalization"},
		"evidence:ostn-vala-typo-warning-001":           {EvidenceType: "typo_squatting_warning", Source: "oss-trust-network/vala/typo-warning", Scope: "typo_squatting_warning"},
		"evidence:ostn-vala-drift-signal-001":           {EvidenceType: "drift_signal", Source: "oss-trust-network/vala/drift-signal", Scope: "drift_signal"},
		"evidence:ostn-vala-no-overclaim-001":           {EvidenceType: "no_overclaim", Source: "oss-trust-network/vala/no-overclaim", Scope: "no_overclaim_discipline"},
		"evidence:ostn-vala-canonical-boundary-001":     {EvidenceType: "canonical_boundary", Source: "oss-trust-network/vala/canonical-boundary", Scope: "canonical_evidence_boundary"},
		"evidence:ostn-vala-point9-governance-001":      {EvidenceType: "state_governance", Source: "oss-trust-network/vala/point9-governance", Scope: "point9_governance"},
	}
}

func ossTrustNetworkValAEvidence() []ReferenceArchitectureEvidenceReference {
	return []ReferenceArchitectureEvidenceReference{
		{EvidenceID: "evidence:ostn-vala-dependency-001", EvidenceType: "dependency_state", Source: "oss-trust-network/vala/dependency", Timestamp: "2026-04-29T10:05:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "val0_dependency", Caveats: []string{"Val A depends on exact and active OSTN Val 0 discipline only."}},
		{EvidenceID: "evidence:ostn-vala-release-trust-intake-001", EvidenceType: "release_trust_intake", Source: "oss-trust-network/vala/release-trust-intake", Timestamp: "2026-04-29T10:06:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "release_trust_intake", Caveats: []string{"Release trust intake remains bounded, freshness-aware, and evidence-linked."}},
		{EvidenceID: "evidence:ostn-vala-signing-signal-001", EvidenceType: "signing_signal", Source: "oss-trust-network/vala/signing-signal", Timestamp: "2026-04-29T10:07:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "signing_signal", Caveats: []string{"Signing is a bounded release signal and not maintainer approval or certification."}},
		{EvidenceID: "evidence:ostn-vala-maintainer-attestation-001", EvidenceType: "maintainer_attestation", Source: "oss-trust-network/vala/maintainer-attestation", Timestamp: "2026-04-29T10:08:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "maintainer_attestation", Caveats: []string{"Maintainer attestation remains bounded and cannot override enterprise-local evidence."}},
		{EvidenceID: "evidence:ostn-vala-provenance-material-001", EvidenceType: "provenance_material", Source: "oss-trust-network/vala/provenance-material", Timestamp: "2026-04-29T10:09:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "provenance_material", Caveats: []string{"Provenance material is evidence material only and not universal truth."}},
		{EvidenceID: "evidence:ostn-vala-registry-descriptor-001", EvidenceType: "registry_descriptor", Source: "oss-trust-network/vala/registry-descriptor", Timestamp: "2026-04-29T10:10:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "registry_descriptor", Caveats: []string{"Registry descriptors remain descriptor-only and do not perform live trust fetches here."}},
		{EvidenceID: "evidence:ostn-vala-registry-metadata-001", EvidenceType: "registry_metadata", Source: "oss-trust-network/vala/registry-metadata", Timestamp: "2026-04-29T10:11:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "registry_metadata_normalization", Caveats: []string{"Registry metadata is normalized input only and cannot create reviewed trust by itself."}},
		{EvidenceID: "evidence:ostn-vala-typo-warning-001", EvidenceType: "typo_squatting_warning", Source: "oss-trust-network/vala/typo-warning", Timestamp: "2026-04-29T10:12:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "typo_squatting_warning", Caveats: []string{"Typo-squatting warnings remain candidate early warnings with visible false-positive handling."}},
		{EvidenceID: "evidence:ostn-vala-drift-signal-001", EvidenceType: "drift_signal", Source: "oss-trust-network/vala/drift-signal", Timestamp: "2026-04-29T10:13:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "drift_signal", Caveats: []string{"Drift signals remain source-weighted and scoped, not global truth or automatic global blocking."}},
		{EvidenceID: "evidence:ostn-vala-no-overclaim-001", EvidenceType: "no_overclaim", Source: "oss-trust-network/vala/no-overclaim", Timestamp: "2026-04-29T10:14:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "no_overclaim_discipline", Caveats: []string{"OSTN Val A cannot certify packages, create global truth, or expose forbidden point pass claims."}},
		{EvidenceID: "evidence:ostn-vala-canonical-boundary-001", EvidenceType: "canonical_boundary", Source: "oss-trust-network/vala/canonical-boundary", Timestamp: "2026-04-29T10:15:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "canonical_evidence_boundary", Caveats: []string{"Canonical execution, audit, and evidence remain the only source of truth."}},
		{EvidenceID: "evidence:ostn-vala-point9-governance-001", EvidenceType: "state_governance", Source: "oss-trust-network/vala/point9-governance", Timestamp: "2026-04-29T10:16:00Z", FreshnessState: IntelligenceCalibrationFreshnessFresh, Scope: "point9_governance", Caveats: []string{"point_9_state remains not complete and later OSTN waves are still required."}},
	}
}

func OSSTrustNetworkValAProofEvidenceQualityValid(evidence []ReferenceArchitectureEvidenceReference, refs []string) bool {
	if !containsExactTrimmedStringSet(refs, OSSTrustNetworkValAProofEvidenceRefs()...) {
		return false
	}
	expected := ossTrustNetworkValAExpectedEvidenceMetadataByID()
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

func ossTrustNetworkValAContainsForbiddenClaim(values ...string) bool {
	allowed := map[string]struct{}{
		"bounded release trust signal":     {},
		"evidence-linked signing signal":   {},
		"maintainer attestation signal":    {},
		"provenance-aware release signal":  {},
		"bounded registry descriptor":      {},
		"registry metadata input":          {},
		"candidate typo-squatting warning": {},
		"reviewed release trust signal":    {},
		"source-weighted drift signal":     {},
		"advisory oss network signal":      {},
		"not canonical truth":              {},
		"not formal certification":         {},
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

func OSSTrustNetworkValADependencySnapshotModel() OSSTrustNetworkValADependencySnapshot {
	model := ComputeOSSTrustNetworkVal0Foundation(OSSTrustNetworkVal0FoundationModel())
	return OSSTrustNetworkValADependencySnapshot{
		Val0CurrentState:         model.CurrentState,
		Val0Point9State:          model.Point9State,
		Val0DependencyState:      model.DependencyState,
		Val0NoOverclaimState:     model.NoOverclaimState,
		Val0Point8State:          model.Dependency.Point8State,
		Val0Point8PassAllowed:    model.Dependency.Point8PassAllowed,
		Val0Point8PassReason:     model.Dependency.Point8PassReason,
		Val0Point8ClosureState:   model.Dependency.ClosureState,
		Val0ProofSurfaceRefs:     append([]string{}, model.ProofSurfaceRefs...),
		Val0EvidenceRefs:         append([]string{}, model.EvidenceRefs...),
		Val0ProjectionDisclaimer: model.ProjectionDisclaimer,
	}
}

func OSSTrustNetworkValAReleaseTrustIntakeModel() OSSTrustNetworkValAReleaseTrustIntake {
	return OSSTrustNetworkValAReleaseTrustIntake{
		CurrentState:               OSSTrustNetworkValAReleaseTrustIntakeStateActive,
		ReleaseTrustProfileID:      "ostn-vala-release-trust-profile-001",
		PackageOrProjectIdentity:   "github.com/example/project",
		RegistryOrEcosystem:        OSSTrustNetworkValARegistryDescriptorGitHubReleases,
		ReleaseRef:                 "refs/tags/v1.2.3",
		ArtifactIdentity:           "sha256:artifact-example-v123",
		SigningSignalState:         OSSTrustNetworkValASigningSignalStateActive,
		ProvenanceSignalState:      OSSTrustNetworkValAProvenanceMaterialStateActive,
		MaintainerAttestationState: OSSTrustNetworkValAMaintainerAttestationStateActive,
		RegistryMetadataState:      OSSTrustNetworkValARegistryMetadataStateActive,
		EvidenceRefs:               []string{"evidence:ostn-vala-release-trust-intake-001"},
		FreshnessState:             IntelligenceCalibrationFreshnessFresh,
		ReviewState:                OSSTrustNetworkReviewStateReviewed,
		Caveats:                    []string{"reviewed release trust signal", "not canonical truth"},
		ProjectionDisclaimer:       ossTrustNetworkValAProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValASigningSignalModel() OSSTrustNetworkValASigningSignal {
	return OSSTrustNetworkValASigningSignal{
		CurrentState:                  OSSTrustNetworkValASigningSignalStateActive,
		DisciplineID:                  "oss_release_signing_signal",
		Version:                       "v0",
		SigningState:                  OSSTrustNetworkValASigningStateVerified,
		ReleaseRef:                    "refs/tags/v1.2.3",
		ArtifactIdentity:              "sha256:artifact-example-v123",
		EvidenceRefs:                  []string{"evidence:ostn-vala-signing-signal-001"},
		Caveats:                       []string{"evidence-linked signing signal"},
		VerifiedEvidenceLinked:        true,
		ScopedReleaseArtifactIdentity: true,
		ProjectionDisclaimer:          ossTrustNetworkValAProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValAMaintainerAttestationModel() OSSTrustNetworkValAMaintainerAttestation {
	return OSSTrustNetworkValAMaintainerAttestation{
		CurrentState:              OSSTrustNetworkValAMaintainerAttestationStateActive,
		DisciplineID:              "oss_maintainer_attestation_intake",
		Version:                   "v0",
		AttestationState:          OSSTrustNetworkValAAttestationStateAttested,
		PackageOrProjectIdentity:  "github.com/example/project",
		ReleaseRef:                "refs/tags/v1.2.3",
		EvidenceRefs:              []string{"evidence:ostn-vala-maintainer-attestation-001"},
		Caveats:                   []string{"maintainer attestation signal"},
		IdentityBindingVisible:    true,
		KeyLinkageVisible:         true,
		DelegationHandlingVisible: true,
		KeyRotationHandled:        true,
		KeyRevocationHandled:      true,
		CompromiseHandlingVisible: true,
		ProjectionDisclaimer:      ossTrustNetworkValAProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValAProvenanceMaterialModel() OSSTrustNetworkValAProvenanceMaterial {
	return OSSTrustNetworkValAProvenanceMaterial{
		CurrentState:                  OSSTrustNetworkValAProvenanceMaterialStateActive,
		DisciplineID:                  "oss_provenance_material",
		Version:                       "v0",
		ProvenanceState:               OSSTrustNetworkValAProvenanceStateVerified,
		ReleaseRef:                    "refs/tags/v1.2.3",
		ArtifactIdentity:              "sha256:artifact-example-v123",
		EvidenceRefs:                  []string{"evidence:ostn-vala-provenance-material-001"},
		Caveats:                       []string{"provenance-aware release signal"},
		ScopedReleaseArtifactIdentity: true,
		ProjectionDisclaimer:          ossTrustNetworkValAProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValARegistryDescriptorModel() OSSTrustNetworkValARegistryDescriptor {
	return OSSTrustNetworkValARegistryDescriptor{
		CurrentState:                 OSSTrustNetworkValARegistryDescriptorStateActive,
		DisciplineID:                 "oss_registry_connector_descriptor",
		Version:                      "v0",
		SupportedRegistryDescriptors: ossTrustNetworkValARegistryDescriptorClasses(),
		RequestedRegistryDescriptor:  OSSTrustNetworkValARegistryDescriptorGitHubReleases,
		DescriptorOnlyBehavior:       true,
		ProjectionDisclaimer:         ossTrustNetworkValAProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValARegistryMetadataModel() OSSTrustNetworkValARegistryMetadata {
	return OSSTrustNetworkValARegistryMetadata{
		CurrentState:             OSSTrustNetworkValARegistryMetadataStateActive,
		DisciplineID:             "oss_registry_metadata_normalization",
		Version:                  "v0",
		PackageName:              "example-project",
		Registry:                 OSSTrustNetworkValARegistryDescriptorGitHubReleases,
		VersionOrRelease:         "v1.2.3",
		PublisherOrMaintainerRef: "maintainer:example-project",
		ReleaseTimestamp:         "2026-04-28T18:00:00Z",
		ArtifactRef:              "sha256:artifact-example-v123",
		SourceRepositoryRef:      "https://github.com/example/project",
		MetadataFreshness:        IntelligenceCalibrationFreshnessFresh,
		EvidenceRefs:             []string{"evidence:ostn-vala-registry-metadata-001"},
		Caveats:                  []string{"registry metadata input"},
		ProjectionDisclaimer:     ossTrustNetworkValAProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValATypoSquattingWarningModel() OSSTrustNetworkValATypoSquattingWarning {
	return OSSTrustNetworkValATypoSquattingWarning{
		CurrentState:                 OSSTrustNetworkValATypoSquattingWarningStateActive,
		DisciplineID:                 "oss_typo_squatting_warning",
		Version:                      "v0",
		WarningID:                    "ostn-vala-typo-warning-001",
		SimilarityBasis:              "name_similarity_release_ref_distance",
		Scope:                        OSSTrustNetworkApplicabilityProjectScoped,
		ReviewState:                  OSSTrustNetworkReviewStateCandidate,
		EvidenceRefs:                 []string{"evidence:ostn-vala-typo-warning-001"},
		Caveats:                      []string{"candidate typo-squatting warning"},
		FalsePositiveHandlingVisible: true,
		ProjectionDisclaimer:         ossTrustNetworkValAProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValADriftSignalModel() OSSTrustNetworkValADriftSignal {
	return OSSTrustNetworkValADriftSignal{
		CurrentState:         OSSTrustNetworkValADriftSignalStateActive,
		DisciplineID:         "oss_release_drift_signal",
		Version:              "v0",
		SignalID:             "ostn-vala-drift-001",
		DriftClass:           OSSTrustNetworkValADriftClassSigning,
		DriftState:           OSSTrustNetworkValADriftStateCandidate,
		SourceWeightClass:    OSSTrustNetworkSourceWeightMedium,
		ApplicabilityScope:   OSSTrustNetworkApplicabilityEnterpriseLocal,
		EvidenceRefs:         []string{"evidence:ostn-vala-drift-signal-001"},
		Caveats:              []string{"source-weighted drift signal"},
		ProjectionDisclaimer: ossTrustNetworkValAProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValANoOverclaimModel() OSSTrustNetworkValANoOverclaim {
	return OSSTrustNetworkValANoOverclaim{
		CurrentState:         OSSTrustNetworkValANoOverclaimStateActive,
		DisciplineID:         "oss_release_trust_no_overclaim",
		Version:              "v0",
		ObservedClaims:       []string{"bounded release trust signal", "evidence-linked signing signal", "bounded registry descriptor", "candidate typo-squatting warning", "source-weighted drift signal", "not canonical truth"},
		ProjectionDisclaimer: ossTrustNetworkValAProjectionDisclaimer(),
	}
}

func OSSTrustNetworkValACoreModel() OSSTrustNetworkValACore {
	return OSSTrustNetworkValACore{
		CurrentState:               OSSTrustNetworkValAStateActive,
		Point9State:                OSSTrustNetworkPoint9StateNotComplete,
		DependencyState:            OSSTrustNetworkValADependencyStateActive,
		ReleaseTrustIntakeState:    OSSTrustNetworkValAReleaseTrustIntakeStateActive,
		SigningSignalState:         OSSTrustNetworkValASigningSignalStateActive,
		MaintainerAttestationState: OSSTrustNetworkValAMaintainerAttestationStateActive,
		ProvenanceMaterialState:    OSSTrustNetworkValAProvenanceMaterialStateActive,
		RegistryDescriptorState:    OSSTrustNetworkValARegistryDescriptorStateActive,
		RegistryMetadataState:      OSSTrustNetworkValARegistryMetadataStateActive,
		TypoSquattingWarningState:  OSSTrustNetworkValATypoSquattingWarningStateActive,
		DriftSignalState:           OSSTrustNetworkValADriftSignalStateActive,
		NoOverclaimState:           OSSTrustNetworkValANoOverclaimStateActive,
		Dependency:                 OSSTrustNetworkValADependencySnapshotModel(),
		ReleaseTrustIntake:         OSSTrustNetworkValAReleaseTrustIntakeModel(),
		SigningSignal:              OSSTrustNetworkValASigningSignalModel(),
		MaintainerAttestation:      OSSTrustNetworkValAMaintainerAttestationModel(),
		ProvenanceMaterial:         OSSTrustNetworkValAProvenanceMaterialModel(),
		RegistryDescriptor:         OSSTrustNetworkValARegistryDescriptorModel(),
		RegistryMetadata:           OSSTrustNetworkValARegistryMetadataModel(),
		TypoSquattingWarning:       OSSTrustNetworkValATypoSquattingWarningModel(),
		DriftSignal:                OSSTrustNetworkValADriftSignalModel(),
		NoOverclaim:                OSSTrustNetworkValANoOverclaimModel(),
		ProofSurfaceRefs:           OSSTrustNetworkValAProofSurfaceRefs(),
		EvidenceRefs:               OSSTrustNetworkValAProofEvidenceRefs(),
		WhyPoint9NotComplete: []string{
			"Val A provides bounded release trust and registry core only and cannot complete Točka 9.",
			"Shared reviewed intelligence, propagation workflows, dashboards, final gates, and integrated closure remain for later OSTN waves.",
			"Val A keeps release, maintainer, registry, provenance, typo-warning, and drift outputs advisory and not canonical truth.",
		},
		ProjectionDisclaimer: ossTrustNetworkValAProjectionDisclaimer(),
	}
}

func EvaluateOSSTrustNetworkValADependencyState(model OSSTrustNetworkValADependencySnapshot) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.Val0CurrentState,
		model.Val0Point9State,
		model.Val0DependencyState,
		model.Val0NoOverclaimState,
		model.Val0Point8State,
		model.Val0Point8ClosureState,
		model.Val0ProjectionDisclaimer,
	) || len(model.Val0ProofSurfaceRefs) == 0 || len(model.Val0EvidenceRefs) == 0 {
		return OSSTrustNetworkValADependencyStateIncomplete
	}
	if !ossTrustNetworkVal0HasProjectionDisclaimer(model.Val0ProjectionDisclaimer) {
		return OSSTrustNetworkValADependencyStateUnknown
	}
	if !containsExactTrimmedStringSet(model.Val0ProofSurfaceRefs, OSSTrustNetworkVal0ProofSurfaceRefs()...) ||
		!containsExactTrimmedStringSet(model.Val0EvidenceRefs, OSSTrustNetworkVal0ProofEvidenceRefs()...) ||
		!OSSTrustNetworkVal0ProofEvidenceQualityValid(ossTrustNetworkVal0Evidence(), model.Val0EvidenceRefs) {
		return OSSTrustNetworkValADependencyStateBlocked
	}
	if strings.TrimSpace(model.Val0CurrentState) != OSSTrustNetworkVal0StateActive ||
		strings.TrimSpace(model.Val0Point9State) != OSSTrustNetworkPoint9StateNotComplete ||
		strings.TrimSpace(model.Val0DependencyState) != OSSTrustNetworkVal0DependencyStateActive ||
		strings.TrimSpace(model.Val0NoOverclaimState) != OSSTrustNetworkVal0NoOverclaimStateActive ||
		strings.TrimSpace(model.Val0Point8State) != DeveloperEcosystemPoint8StatePass ||
		!model.Val0Point8PassAllowed ||
		strings.TrimSpace(model.Val0Point8PassReason) != DeveloperEcosystemValEPoint8PassReasonAllowed ||
		strings.TrimSpace(model.Val0Point8ClosureState) != DeveloperEcosystemValEClosureStateActive {
		return OSSTrustNetworkValADependencyStateBlocked
	}
	return OSSTrustNetworkValADependencyStateActive
}

func EvaluateOSSTrustNetworkValAReleaseTrustIntakeState(model OSSTrustNetworkValAReleaseTrustIntake) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.ReleaseTrustProfileID,
		model.PackageOrProjectIdentity,
		model.RegistryOrEcosystem,
		model.ReleaseRef,
		model.ArtifactIdentity,
		model.SigningSignalState,
		model.ProvenanceSignalState,
		model.MaintainerAttestationState,
		model.RegistryMetadataState,
		model.FreshnessState,
		model.ReviewState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValAReleaseTrustIntakeStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValAReleaseTrustIntakeStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vala-release-trust-intake-001") ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.FreshnessState) ||
		!containsTrimmedString(ossTrustNetworkValAReviewStates(), model.ReviewState) ||
		model.PackageSafetyClaim ||
		model.CertifiedPackageClaim ||
		model.ProductionApprovalClaim ||
		model.GlobalApprovalClaim {
		return OSSTrustNetworkValAReleaseTrustIntakeStateBlocked
	}
	if strings.TrimSpace(model.SigningSignalState) == OSSTrustNetworkValASigningSignalStateActive &&
		strings.TrimSpace(model.ProvenanceSignalState) == OSSTrustNetworkValAProvenanceMaterialStateActive &&
		strings.TrimSpace(model.MaintainerAttestationState) == OSSTrustNetworkValAMaintainerAttestationStateActive &&
		strings.TrimSpace(model.RegistryMetadataState) == OSSTrustNetworkValARegistryMetadataStateActive &&
		strings.TrimSpace(model.ReviewState) == OSSTrustNetworkReviewStateReviewed {
		return OSSTrustNetworkValAReleaseTrustIntakeStateActive
	}
	if strings.TrimSpace(model.ReviewState) == OSSTrustNetworkReviewStateCandidate {
		return OSSTrustNetworkValAReleaseTrustIntakeStatePartial
	}
	return OSSTrustNetworkValAReleaseTrustIntakeStatePartial
}

func EvaluateOSSTrustNetworkValASigningSignalState(model OSSTrustNetworkValASigningSignal) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.SigningState,
		model.ReleaseRef,
		model.ArtifactIdentity,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValASigningSignalStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValASigningSignalStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vala-signing-signal-001") ||
		!containsTrimmedString(ossTrustNetworkValASigningStates(), model.SigningState) ||
		model.MaintainerApprovalClaim ||
		model.PackageCertificationClaim ||
		model.CanonicalTruthClaim {
		return OSSTrustNetworkValASigningSignalStateBlocked
	}
	switch strings.TrimSpace(model.SigningState) {
	case OSSTrustNetworkValASigningStateVerified:
		if model.VerifiedEvidenceLinked && model.ScopedReleaseArtifactIdentity {
			return OSSTrustNetworkValASigningSignalStateActive
		}
		return OSSTrustNetworkValASigningSignalStateBlocked
	case OSSTrustNetworkValASigningStatePresent:
		return OSSTrustNetworkValASigningSignalStatePartial
	case OSSTrustNetworkValASigningStateMissing,
		OSSTrustNetworkValASigningStateMismatch,
		OSSTrustNetworkValASigningStateRevoked,
		OSSTrustNetworkValASigningStateUnsupported,
		OSSTrustNetworkValASigningStateUnknown:
		return OSSTrustNetworkValASigningSignalStateBlocked
	default:
		return OSSTrustNetworkValASigningSignalStateBlocked
	}
}

func EvaluateOSSTrustNetworkValAMaintainerAttestationState(model OSSTrustNetworkValAMaintainerAttestation) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.AttestationState,
		model.PackageOrProjectIdentity,
		model.ReleaseRef,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValAMaintainerAttestationStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValAMaintainerAttestationStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vala-maintainer-attestation-001") ||
		!containsTrimmedString(ossTrustNetworkValAAttestationStates(), model.AttestationState) ||
		model.OverridesLocalEnterprise {
		return OSSTrustNetworkValAMaintainerAttestationStateBlocked
	}
	lifecycleActive := model.IdentityBindingVisible &&
		model.KeyLinkageVisible &&
		model.DelegationHandlingVisible &&
		model.KeyRotationHandled &&
		model.KeyRevocationHandled &&
		model.CompromiseHandlingVisible
	switch strings.TrimSpace(model.AttestationState) {
	case OSSTrustNetworkValAAttestationStateAttested:
		if lifecycleActive {
			return OSSTrustNetworkValAMaintainerAttestationStateActive
		}
		return OSSTrustNetworkValAMaintainerAttestationStateBlocked
	case OSSTrustNetworkValAAttestationStateDelegated:
		if lifecycleActive && model.DelegatedSigningReviewed {
			return OSSTrustNetworkValAMaintainerAttestationStateActive
		}
		return OSSTrustNetworkValAMaintainerAttestationStateBlocked
	case OSSTrustNetworkValAAttestationStateMissing,
		OSSTrustNetworkValAAttestationStateStale,
		OSSTrustNetworkValAAttestationStateRevoked,
		OSSTrustNetworkValAAttestationStateUnsupported,
		OSSTrustNetworkValAAttestationStateUnknown:
		return OSSTrustNetworkValAMaintainerAttestationStateBlocked
	default:
		return OSSTrustNetworkValAMaintainerAttestationStateBlocked
	}
}

func EvaluateOSSTrustNetworkValAProvenanceMaterialState(model OSSTrustNetworkValAProvenanceMaterial) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.ProvenanceState,
		model.ReleaseRef,
		model.ArtifactIdentity,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValAProvenanceMaterialStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValAProvenanceMaterialStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vala-provenance-material-001") ||
		!containsTrimmedString(ossTrustNetworkValAProvenanceStates(), model.ProvenanceState) ||
		model.RegulatorApprovalClaim ||
		model.CertifiedPackageClaim ||
		model.LegalClearanceClaim ||
		model.ProductionApprovalClaim {
		return OSSTrustNetworkValAProvenanceMaterialStateBlocked
	}
	switch strings.TrimSpace(model.ProvenanceState) {
	case OSSTrustNetworkValAProvenanceStateVerified:
		if model.ScopedReleaseArtifactIdentity {
			return OSSTrustNetworkValAProvenanceMaterialStateActive
		}
		return OSSTrustNetworkValAProvenanceMaterialStateBlocked
	case OSSTrustNetworkValAProvenanceStatePresentUnverified:
		return OSSTrustNetworkValAProvenanceMaterialStatePartial
	case OSSTrustNetworkValAProvenanceStateMissing,
		OSSTrustNetworkValAProvenanceStateMismatch,
		OSSTrustNetworkValAProvenanceStateStale,
		OSSTrustNetworkValAProvenanceStateUnsupported,
		OSSTrustNetworkValAProvenanceStateUnknown:
		return OSSTrustNetworkValAProvenanceMaterialStateBlocked
	default:
		return OSSTrustNetworkValAProvenanceMaterialStateBlocked
	}
}

func EvaluateOSSTrustNetworkValARegistryDescriptorState(model OSSTrustNetworkValARegistryDescriptor) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.RequestedRegistryDescriptor, model.ProjectionDisclaimer) ||
		len(model.SupportedRegistryDescriptors) == 0 {
		return OSSTrustNetworkValARegistryDescriptorStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValARegistryDescriptorStateUnknown
	}
	if !containsExactTrimmedStringSet(model.SupportedRegistryDescriptors, ossTrustNetworkValARegistryDescriptorClasses()...) ||
		model.RegistryMetadataCreatesReviewedTrust ||
		!model.DescriptorOnlyBehavior ||
		model.LiveNetworkFetcher {
		return OSSTrustNetworkValARegistryDescriptorStateBlocked
	}
	if strings.TrimSpace(model.RequestedRegistryDescriptor) == OSSTrustNetworkValARegistryDescriptorUnsupported {
		if model.UnsupportedRegistryExplicit {
			return OSSTrustNetworkValARegistryDescriptorStatePartial
		}
		return OSSTrustNetworkValARegistryDescriptorStateBlocked
	}
	if !containsTrimmedString(ossTrustNetworkValARegistryDescriptorClasses(), model.RequestedRegistryDescriptor) {
		return OSSTrustNetworkValARegistryDescriptorStateBlocked
	}
	return OSSTrustNetworkValARegistryDescriptorStateActive
}

func EvaluateOSSTrustNetworkValARegistryMetadataState(model OSSTrustNetworkValARegistryMetadata) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.PackageName,
		model.Registry,
		model.VersionOrRelease,
		model.PublisherOrMaintainerRef,
		model.ReleaseTimestamp,
		model.ArtifactRef,
		model.SourceRepositoryRef,
		model.MetadataFreshness,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValARegistryMetadataStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValARegistryMetadataStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vala-registry-metadata-001") ||
		!containsTrimmedString(ossTrustNetworkValARegistryDescriptorClasses(), model.Registry) ||
		model.RegistryMetadataCreatesReviewedTrust ||
		model.RegistryMetadataCreatesGlobalBlocklist ||
		!ossTrustNetworkVal0FreshnessStateIsExactlyFresh(model.MetadataFreshness) {
		return OSSTrustNetworkValARegistryMetadataStateBlocked
	}
	if _, valid := referenceArchitectureVal0ParseTimestamp(model.ReleaseTimestamp); !valid {
		return OSSTrustNetworkValARegistryMetadataStateBlocked
	}
	return OSSTrustNetworkValARegistryMetadataStateActive
}

func EvaluateOSSTrustNetworkValATypoSquattingWarningState(model OSSTrustNetworkValATypoSquattingWarning) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.WarningID,
		model.SimilarityBasis,
		model.Scope,
		model.ReviewState,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValATypoSquattingWarningStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValATypoSquattingWarningStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vala-typo-warning-001") ||
		!containsTrimmedString(ossTrustNetworkValAReviewStates(), model.ReviewState) ||
		!containsTrimmedString(ossTrustNetworkVal0ApplicabilityScopes(), model.Scope) {
		return OSSTrustNetworkValATypoSquattingWarningStateBlocked
	}
	if !model.FalsePositiveHandlingVisible ||
		model.AutomaticGlobalBlock ||
		model.CanonicalTruthClaim ||
		model.CandidatePromotedToReviewed {
		return OSSTrustNetworkValATypoSquattingWarningStateBlocked
	}
	if strings.TrimSpace(model.ReviewState) == OSSTrustNetworkReviewStateCandidate {
		return OSSTrustNetworkValATypoSquattingWarningStateActive
	}
	return OSSTrustNetworkValATypoSquattingWarningStateBlocked
}

func EvaluateOSSTrustNetworkValADriftSignalState(model OSSTrustNetworkValADriftSignal) string {
	if !referenceArchitectureValBRequiredRefsPresent(
		model.DisciplineID,
		model.Version,
		model.SignalID,
		model.DriftClass,
		model.DriftState,
		model.SourceWeightClass,
		model.ApplicabilityScope,
		model.ProjectionDisclaimer,
	) || len(model.EvidenceRefs) == 0 || len(model.Caveats) == 0 {
		return OSSTrustNetworkValADriftSignalStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValADriftSignalStateUnknown
	}
	if !containsExactTrimmedStringSet(model.EvidenceRefs, "evidence:ostn-vala-drift-signal-001") ||
		!containsTrimmedString(ossTrustNetworkValADriftClasses(), model.DriftClass) ||
		!containsTrimmedString(ossTrustNetworkValADriftStates(), model.DriftState) ||
		!containsTrimmedString(ossTrustNetworkVal0SourceWeightClasses(), model.SourceWeightClass) ||
		!containsTrimmedString(ossTrustNetworkVal0ApplicabilityScopes(), model.ApplicabilityScope) {
		return OSSTrustNetworkValADriftSignalStateBlocked
	}
	if model.AutomaticGlobalBlock || model.GlobalTruthClaim || model.OverridesLocalEnterprise {
		return OSSTrustNetworkValADriftSignalStateBlocked
	}
	switch strings.TrimSpace(model.DriftState) {
	case OSSTrustNetworkValADriftStateCandidate, OSSTrustNetworkValADriftStateReviewed:
		return OSSTrustNetworkValADriftSignalStateActive
	case OSSTrustNetworkValADriftStateUnsupported, OSSTrustNetworkValADriftStateUnknown:
		return OSSTrustNetworkValADriftSignalStateBlocked
	default:
		return OSSTrustNetworkValADriftSignalStateBlocked
	}
}

func EvaluateOSSTrustNetworkValANoOverclaimState(model OSSTrustNetworkValANoOverclaim) string {
	if !referenceArchitectureValBRequiredRefsPresent(model.DisciplineID, model.Version, model.ProjectionDisclaimer) {
		return OSSTrustNetworkValANoOverclaimStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValANoOverclaimStateUnknown
	}
	if model.GlobalTruthClaim ||
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
		ossTrustNetworkValAContainsForbiddenClaim(model.ObservedClaims...) {
		return OSSTrustNetworkValANoOverclaimStateBlocked
	}
	return OSSTrustNetworkValANoOverclaimStateActive
}

func EvaluateOSSTrustNetworkValAState(model OSSTrustNetworkValACore) string {
	if strings.TrimSpace(model.Point9State) == "" || len(model.ProofSurfaceRefs) == 0 || len(model.EvidenceRefs) == 0 || strings.TrimSpace(model.ProjectionDisclaimer) == "" {
		return OSSTrustNetworkValAStateIncomplete
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) {
		return OSSTrustNetworkValAStateUnknown
	}
	if strings.TrimSpace(model.Point9State) != OSSTrustNetworkPoint9StateNotComplete ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValAProofSurfaceRefs()...) ||
		!OSSTrustNetworkValAProofEvidenceQualityValid(ossTrustNetworkValAEvidence(), model.EvidenceRefs) {
		return OSSTrustNetworkValAStateBlocked
	}
	states := []string{
		model.DependencyState,
		model.ReleaseTrustIntakeState,
		model.SigningSignalState,
		model.MaintainerAttestationState,
		model.ProvenanceMaterialState,
		model.RegistryDescriptorState,
		model.RegistryMetadataState,
		model.TypoSquattingWarningState,
		model.DriftSignalState,
		model.NoOverclaimState,
	}
	allActive := true
	for _, state := range states {
		if strings.TrimSpace(state) == "" {
			return OSSTrustNetworkValAStateIncomplete
		}
		if !strings.HasSuffix(strings.TrimSpace(state), "_active") {
			allActive = false
		}
	}
	if allActive {
		return OSSTrustNetworkValAStateActive
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_blocked") {
			return OSSTrustNetworkValAStateBlocked
		}
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_incomplete") {
			return OSSTrustNetworkValAStateIncomplete
		}
	}
	for _, state := range states {
		if strings.HasSuffix(strings.TrimSpace(state), "_unknown") {
			return OSSTrustNetworkValAStateUnknown
		}
	}
	return OSSTrustNetworkValAStatePartial
}

func EvaluateOSSTrustNetworkValAPointsState(currentState string) string {
	_ = currentState
	return OSSTrustNetworkPoint9StateNotComplete
}

func EvaluateOSSTrustNetworkValAProofsState(model OSSTrustNetworkValACore, limitations []string) string {
	baseState := strings.TrimSpace(model.CurrentState)
	if baseState == "" {
		baseState = OSSTrustNetworkValAStateUnknown
	}
	if !ossTrustNetworkValAHasProjectionDisclaimer(model.ProjectionDisclaimer) ||
		!containsExactTrimmedStringSet(model.ProofSurfaceRefs, OSSTrustNetworkValAProofSurfaceRefs()...) ||
		!OSSTrustNetworkValAProofEvidenceQualityValid(ossTrustNetworkValAEvidence(), model.EvidenceRefs) ||
		len(limitations) == 0 ||
		strings.TrimSpace(model.Point9State) != OSSTrustNetworkPoint9StateNotComplete {
		if baseState == OSSTrustNetworkValAStateActive {
			return OSSTrustNetworkValAStatePartial
		}
		return baseState
	}
	return baseState
}

func ossTrustNetworkValABlockingReasons(model OSSTrustNetworkValACore) []string {
	reasons := []string{}
	if model.DependencyState != OSSTrustNetworkValADependencyStateActive {
		reasons = append(reasons, "OSTN Val 0 dependency is not exact, active, and no-overclaim-safe.")
	}
	if model.ReleaseTrustIntakeState != OSSTrustNetworkValAReleaseTrustIntakeStateActive {
		reasons = append(reasons, "Release trust intake is not fully identified, fresh, reviewed, and bounded.")
	}
	if model.SigningSignalState != OSSTrustNetworkValASigningSignalStateActive {
		reasons = append(reasons, "Signing signal intake is not exact, verified, and evidence-linked.")
	}
	if model.MaintainerAttestationState != OSSTrustNetworkValAMaintainerAttestationStateActive {
		reasons = append(reasons, "Maintainer attestation intake is not lifecycle-safe and enterprise-bounded.")
	}
	if model.ProvenanceMaterialState != OSSTrustNetworkValAProvenanceMaterialStateActive {
		reasons = append(reasons, "Provenance material is not exact, scoped, and release-linked.")
	}
	if model.RegistryDescriptorState != OSSTrustNetworkValARegistryDescriptorStateActive {
		reasons = append(reasons, "Registry descriptor discipline is not exact, descriptor-only, and bounded.")
	}
	if model.RegistryMetadataState != OSSTrustNetworkValARegistryMetadataStateActive {
		reasons = append(reasons, "Registry metadata normalization is not exact, fresh, and evidence-linked.")
	}
	if model.TypoSquattingWarningState != OSSTrustNetworkValATypoSquattingWarningStateActive {
		reasons = append(reasons, "Typo-squatting warning discipline is not bounded as candidate early warning.")
	}
	if model.DriftSignalState != OSSTrustNetworkValADriftSignalStateActive {
		reasons = append(reasons, "Drift signal discipline is not evidence-linked, source-weighted, scoped, and bounded.")
	}
	if model.NoOverclaimState != OSSTrustNetworkValANoOverclaimStateActive {
		reasons = append(reasons, "Val A no-overclaim and no-global-truth guard is not active.")
	}
	return developerEcosystemValECollectText(reasons)
}

func ComputeOSSTrustNetworkValACore(model OSSTrustNetworkValACore) OSSTrustNetworkValACore {
	model.DependencyState = EvaluateOSSTrustNetworkValADependencyState(model.Dependency)
	model.SigningSignalState = EvaluateOSSTrustNetworkValASigningSignalState(model.SigningSignal)
	model.MaintainerAttestationState = EvaluateOSSTrustNetworkValAMaintainerAttestationState(model.MaintainerAttestation)
	model.ProvenanceMaterialState = EvaluateOSSTrustNetworkValAProvenanceMaterialState(model.ProvenanceMaterial)
	model.RegistryDescriptorState = EvaluateOSSTrustNetworkValARegistryDescriptorState(model.RegistryDescriptor)
	model.RegistryMetadataState = EvaluateOSSTrustNetworkValARegistryMetadataState(model.RegistryMetadata)
	model.TypoSquattingWarningState = EvaluateOSSTrustNetworkValATypoSquattingWarningState(model.TypoSquattingWarning)
	model.DriftSignalState = EvaluateOSSTrustNetworkValADriftSignalState(model.DriftSignal)
	model.NoOverclaimState = EvaluateOSSTrustNetworkValANoOverclaimState(model.NoOverclaim)

	model.ReleaseTrustIntake.SigningSignalState = model.SigningSignalState
	model.ReleaseTrustIntake.ProvenanceSignalState = model.ProvenanceMaterialState
	model.ReleaseTrustIntake.MaintainerAttestationState = model.MaintainerAttestationState
	model.ReleaseTrustIntake.RegistryMetadataState = model.RegistryMetadataState
	model.ReleaseTrustIntakeState = EvaluateOSSTrustNetworkValAReleaseTrustIntakeState(model.ReleaseTrustIntake)

	model.Point9State = EvaluateOSSTrustNetworkValAPointsState(model.CurrentState)
	model.CurrentState = EvaluateOSSTrustNetworkValAState(model)
	model.Point9State = EvaluateOSSTrustNetworkValAPointsState(model.CurrentState)
	model.BlockingReasons = ossTrustNetworkValABlockingReasons(model)

	model.ReleaseTrustIntake.CurrentState = model.ReleaseTrustIntakeState
	model.SigningSignal.CurrentState = model.SigningSignalState
	model.MaintainerAttestation.CurrentState = model.MaintainerAttestationState
	model.ProvenanceMaterial.CurrentState = model.ProvenanceMaterialState
	model.RegistryDescriptor.CurrentState = model.RegistryDescriptorState
	model.RegistryMetadata.CurrentState = model.RegistryMetadataState
	model.TypoSquattingWarning.CurrentState = model.TypoSquattingWarningState
	model.DriftSignal.CurrentState = model.DriftSignalState
	model.NoOverclaim.CurrentState = model.NoOverclaimState

	return model
}

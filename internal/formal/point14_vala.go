package formal

import (
	"sort"
	"strings"

	"github.com/denisgrosek/changelock/internal/operability"
)

const (
	Point14ValAStateActive         = "point14_vala_external_signal_validation_active"
	Point14ValAStateBlocked        = "point14_vala_external_signal_validation_blocked"
	Point14ValAStateReviewRequired = "point14_vala_external_signal_validation_review_required"
	Point14ValAStateIncomplete     = "point14_vala_external_signal_validation_incomplete"
)

const (
	point14ValAWaveID                   = "val_a"
	point14ValAProjectionDisclaimerBase = "projection_only not_canonical_truth point14_vala_external_signal_validation_gate"
	point14ValABlockedPassToken         = "point_14_pass"

	point14ValAValidationCandidateValidated      = "candidate_validated"
	point14ValAValidationCandidateBlocked        = "candidate_blocked"
	point14ValAValidationCandidateReviewRequired = "candidate_review_required"
	point14ValAValidationCandidateIncomplete     = "candidate_incomplete"
	point14ValAValidationCandidateUnsupported    = "candidate_unsupported"
	point14ValAValidationCandidateTampered       = "candidate_tampered"
	point14ValAValidationCandidateDuplicate      = "candidate_duplicate"
	point14ValAValidationCandidateUnrelated      = "candidate_unrelated"
	point14ValAValidationCandidateCrossTenant    = "candidate_cross_tenant"

	point14ValAEvidenceStateActive     = "active"
	point14ValAEvidenceStateStale      = "stale"
	point14ValAEvidenceStateRevoked    = "revoked"
	point14ValAEvidenceStateExpired    = "expired"
	point14ValAEvidenceStateSuperseded = "superseded"
	point14ValAEvidenceStateUnrelated  = "unrelated"
)

type Point14ValADependencySnapshot struct {
	Point14Val0CurrentState                          string                                          `json:"point14_val0_current_state"`
	Point14Val0DependencyState                       string                                          `json:"point14_val0_dependency_state"`
	Point14Val0ExternalSignalCandidateState          string                                          `json:"point14_val0_external_signal_candidate_state"`
	Point14Val0ExternalStakeholderAuthorityRoleState string                                          `json:"point14_val0_external_stakeholder_authority_role_state"`
	Point14Val0ExternalAuthorityConflictMatrixState  string                                          `json:"point14_val0_external_authority_conflict_matrix_state"`
	Point14Val0ExternalSignalDisputeLifecycleState   string                                          `json:"point14_val0_external_signal_dispute_lifecycle_state"`
	Point14Val0ExternalCorrectionRevocationState     string                                          `json:"point14_val0_external_correction_revocation_state"`
	Point14Val0ExternalVisibilityPublicationState    string                                          `json:"point14_val0_external_visibility_publication_state"`
	Point14Val0AgentEcosystemInputBoundaryState      string                                          `json:"point14_val0_agent_ecosystem_input_boundary_state"`
	Point14Val0NoExternalAuthorityGuardState         string                                          `json:"point14_val0_no_external_authority_guard_state"`
	Point14Val0NoOverclaimState                      string                                          `json:"point14_val0_no_overclaim_state"`
	Point14Val0PointID                               string                                          `json:"point14_val0_point_id"`
	Point14Val0WaveID                                string                                          `json:"point14_val0_wave_id"`
	Point14Val0ComputedFromUpstream                  bool                                            `json:"point14_val0_computed_from_upstream"`
	Point14Val0Merged                                bool                                            `json:"point14_val0_merged"`
	Point14Val0CIGreen                               bool                                            `json:"point14_val0_ci_green"`
	Point14Val0ReviewedOnMain                        bool                                            `json:"point14_val0_reviewed_on_main"`
	Point14PassSeen                                  bool                                            `json:"point14_pass_seen"`
	InheritedPoint13ValECurrentState                 string                                          `json:"inherited_point13_vale_current_state"`
	InheritedPoint13ValEPassClosureState             string                                          `json:"inherited_point13_vale_pass_closure_state"`
	InheritedPoint13ValEPassAllowed                  bool                                            `json:"inherited_point13_vale_pass_allowed"`
	InheritedPoint13ValEPassToken                    string                                          `json:"inherited_point13_vale_pass_token"`
	InheritedPoint12CurrentState                     string                                          `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState                  string                                          `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureState                 string                                          `json:"inherited_point12_pass_closure_state"`
	InheritedPoint12ReviewerResult                   string                                          `json:"inherited_point12_reviewer_result"`
	InheritedPoint11PublicationState                 string                                          `json:"inherited_point11_publication_state"`
	InheritedPoint11NoOverclaimState                 string                                          `json:"inherited_point11_no_overclaim_state"`
	InheritedPoint11FinalPassGateState               string                                          `json:"inherited_point11_final_pass_gate_state"`
	InheritedPoint10CurrentState                     string                                          `json:"inherited_point10_current_state"`
	InheritedPoint10NoOverclaimState                 string                                          `json:"inherited_point10_no_overclaim_state"`
	InheritedPoint10ProjectionState                  string                                          `json:"inherited_point10_projection_state"`
	InheritedPoint10PassRuleState                    string                                          `json:"inherited_point10_pass_rule_state"`
	InheritedTenantScope                             string                                          `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput                       bool                                            `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                              []string                                        `json:"review_prerequisites,omitempty"`
	Point14Val0                                      Point14Val0Foundation                           `json:"point14_val0"`
	Point13ValE                                      Point13ValEFoundation                           `json:"point13_vale"`
	Point12                                          Point12ValEFoundation                           `json:"point12"`
	Point11                                          Point11ValDFoundation                           `json:"point11"`
	Point10                                          operability.DeploymentMultiTenantValEFoundation `json:"point10"`
}

type NormalizedExternalSignal struct {
	NormalizedSignalID        string   `json:"normalized_signal_id"`
	OriginalSignalID          string   `json:"original_signal_id"`
	SourceIdentityRef         string   `json:"source_identity_ref"`
	SourceType                string   `json:"source_type"`
	SignalType                string   `json:"signal_type"`
	NormalizedPayloadHash     string   `json:"normalized_payload_hash"`
	TenantScope               string   `json:"tenant_scope"`
	GlobalScopeClassification string   `json:"global_scope_classification"`
	ArtifactRefs              []string `json:"artifact_refs,omitempty"`
	ArtifactScoped            bool     `json:"artifact_scoped"`
	ClaimRefs                 []string `json:"claim_refs,omitempty"`
	ClaimScoped               bool     `json:"claim_scoped"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	EvidenceLinked            bool     `json:"evidence_linked"`
	ReceivedAt                string   `json:"received_at"`
	NormalizationTimeSource   string   `json:"normalization_time_source"`
	ValidationStatus          string   `json:"validation_status"`
	PayloadNormalized         bool     `json:"payload_normalized"`
	CanonicalAuthority        bool     `json:"canonical_authority"`
	PassAllowed               bool     `json:"pass_allowed"`
	CorrectionPublished       bool     `json:"correction_published"`
	ProductionApproved        bool     `json:"production_approved"`
	PublicBadgeAllowed        bool     `json:"public_badge_allowed"`
}

type ExternalSignalSourceIdentity struct {
	SourceIdentityID          string `json:"source_identity_id"`
	SourceType                string `json:"source_type"`
	SourceRef                 string `json:"source_ref"`
	SourceAuthorityScope      string `json:"source_authority_scope"`
	ProvenanceRef             string `json:"provenance_ref"`
	CustodyRef                string `json:"custody_ref"`
	HashRef                   string `json:"hash_ref"`
	SignatureRef              string `json:"signature_ref"`
	RequireProvenance         bool   `json:"require_provenance"`
	RequireCustody            bool   `json:"require_custody"`
	RequireHash               bool   `json:"require_hash"`
	RequireSignature          bool   `json:"require_signature"`
	SourceSupported           bool   `json:"source_supported"`
	SourceTrusted             bool   `json:"source_trusted"`
	CanonicalAuthorityGranted bool   `json:"canonical_authority_granted"`
	PassAllowed               bool   `json:"pass_allowed"`
	OverrideCanonicalDecision bool   `json:"override_canonical_decision"`
	ApproveProduction         bool   `json:"approve_production"`
	CertifyCompliance         bool   `json:"certify_compliance"`
}

type ExternalSignalScopeBinding struct {
	ScopeBindingID            string   `json:"scope_binding_id"`
	ScopeClassification       string   `json:"scope_classification"`
	TenantScope               string   `json:"tenant_scope"`
	ReferencedTenantScope     string   `json:"referenced_tenant_scope"`
	GlobalScopeClassification string   `json:"global_scope_classification"`
	ArtifactScoped            bool     `json:"artifact_scoped"`
	ArtifactRefs              []string `json:"artifact_refs,omitempty"`
	ArtifactHashRefs          []string `json:"artifact_hash_refs,omitempty"`
	ArtifactHashesMatch       bool     `json:"artifact_hashes_match"`
	ClaimScoped               bool     `json:"claim_scoped"`
	ClaimRefs                 []string `json:"claim_refs,omitempty"`
	ClaimBindingExact         bool     `json:"claim_binding_exact"`
	SimilarNameOnly           bool     `json:"similar_name_only"`
	SimilarPathOnly           bool     `json:"similar_path_only"`
	SimilarPackageOnly        bool     `json:"similar_package_only"`
}

type ExternalSignalEvidenceBinding struct {
	EvidenceBindingID        string   `json:"evidence_binding_id"`
	TenantScope              string   `json:"tenant_scope"`
	EvidenceLinked           bool     `json:"evidence_linked"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	EvidenceHashRefs         []string `json:"evidence_hash_refs,omitempty"`
	EvidenceHashesMatch      bool     `json:"evidence_hashes_match"`
	EvidenceState            string   `json:"evidence_state"`
	ProvenanceRef            string   `json:"provenance_ref"`
	CustodyRef               string   `json:"custody_ref"`
	RequireProvenance        bool     `json:"require_provenance"`
	RequireCustody           bool     `json:"require_custody"`
	RequireHash              bool     `json:"require_hash"`
	CanonicalMutationAllowed bool     `json:"canonical_mutation_allowed"`
}

type ExternalSignalFreshnessAndTimestampBoundary struct {
	BoundaryID                 string `json:"boundary_id"`
	TenantScope                string `json:"tenant_scope"`
	ReceivedAt                 string `json:"received_at"`
	ReceivedTimeSource         string `json:"received_time_source"`
	SourceEventAt              string `json:"source_event_at"`
	SourceEventTimeSource      string `json:"source_event_time_source"`
	ValidatedAt                string `json:"validated_at"`
	ValidatedTimeSource        string `json:"validated_time_source"`
	ClientReportedAt           string `json:"client_reported_at"`
	SourceEventAdvisoryOnly    bool   `json:"source_event_advisory_only"`
	ClientTimeMetadataOnly     bool   `json:"client_time_metadata_only"`
	StaleSignal                bool   `json:"stale_signal"`
	GovernanceReviewPathExists bool   `json:"governance_review_path_exists"`
	AuthorityUpgradeRequested  bool   `json:"authority_upgrade_requested"`
}

type ExternalSignalDuplicateAndRelationGuard struct {
	GuardID                       string   `json:"guard_id"`
	TenantScope                   string   `json:"tenant_scope"`
	NormalizedSignalRef           string   `json:"normalized_signal_ref"`
	OriginalSignalRef             string   `json:"original_signal_ref"`
	DuplicateSignalRefs           []string `json:"duplicate_signal_refs,omitempty"`
	DuplicateIdentityKey          string   `json:"duplicate_identity_key"`
	ExpectedNormalizedPayloadHash string   `json:"expected_normalized_payload_hash"`
	ObservedNormalizedPayloadHash string   `json:"observed_normalized_payload_hash"`
	RelatedArtifactRefs           []string `json:"related_artifact_refs,omitempty"`
	RelatedClaimRefs              []string `json:"related_claim_refs,omitempty"`
	ArtifactRelationExact         bool     `json:"artifact_relation_exact"`
	ClaimRelationExact            bool     `json:"claim_relation_exact"`
	ConflictingDuplicate          bool     `json:"conflicting_duplicate"`
	SilentReplacementRequested    bool     `json:"silent_replacement_requested"`
}

type ExternalSignalTenantBoundaryGuard struct {
	GuardID                   string `json:"guard_id"`
	TenantScope               string `json:"tenant_scope"`
	SourceTenantScope         string `json:"source_tenant_scope"`
	ArtifactTenantScope       string `json:"artifact_tenant_scope"`
	ClaimTenantScope          string `json:"claim_tenant_scope"`
	EvidenceTenantScope       string `json:"evidence_tenant_scope"`
	ScopeClassification       string `json:"scope_classification"`
	ExplicitBoundedRuleExists bool   `json:"explicit_bounded_rule_exists"`
	TenantPrivateDataExposed  bool   `json:"tenant_private_data_exposed"`
}

type ExternalSignalValidationResult struct {
	ValidationResultID          string `json:"validation_result_id"`
	NormalizedSignalRef         string `json:"normalized_signal_ref"`
	SourceIdentityRef           string `json:"source_identity_ref"`
	ScopeBindingRef             string `json:"scope_binding_ref"`
	EvidenceBindingRef          string `json:"evidence_binding_ref"`
	FreshnessBoundaryRef        string `json:"freshness_boundary_ref"`
	DuplicateGuardRef           string `json:"duplicate_guard_ref"`
	TenantBoundaryRef           string `json:"tenant_boundary_ref"`
	NormalizedSignalState       string `json:"normalized_signal_state"`
	SourceIdentityState         string `json:"source_identity_state"`
	ScopeBindingState           string `json:"scope_binding_state"`
	EvidenceBindingState        string `json:"evidence_binding_state"`
	FreshnessTimestampState     string `json:"freshness_timestamp_state"`
	DuplicateRelationGuardState string `json:"duplicate_relation_guard_state"`
	TenantBoundaryGuardState    string `json:"tenant_boundary_guard_state"`
	NoExternalAuthorityState    string `json:"no_external_authority_state"`
	NoOverclaimState            string `json:"no_overclaim_state"`
	ValidationState             string `json:"validation_state"`
	CandidateUsable             bool   `json:"candidate_usable"`
	EmitsPass                   bool   `json:"emits_pass"`
	PublishesCorrection         bool   `json:"publishes_correction"`
	RevokesClaim                bool   `json:"revokes_claim"`
	OverridesCanonicalDecision  bool   `json:"overrides_canonical_decision"`
	ApprovesProduction          bool   `json:"approves_production"`
	CertifiesCompliance         bool   `json:"certifies_compliance"`
	CreatesPublicBadge          bool   `json:"creates_public_badge"`
}

type Point14ValANoExternalAuthorityValidationGuard struct {
	ObservedAuthorityMarkers  []string `json:"observed_authority_markers,omitempty"`
	CanonicalAuthorityGranted bool     `json:"canonical_authority_granted"`
	ProductionApprovalGranted bool     `json:"production_approval_granted"`
	CorrectionPublished       bool     `json:"correction_published"`
	PublicBadgeAllowed        bool     `json:"public_badge_allowed"`
	ExternalAuthorityAllowed  bool     `json:"external_authority_allowed"`
}

type Point14ValANoOverclaimValidationWording struct {
	ObservedNormalizationTexts           []string `json:"observed_normalization_texts,omitempty"`
	ObservedValidationTexts              []string `json:"observed_validation_texts,omitempty"`
	ObservedSourceIdentityTexts          []string `json:"observed_source_identity_texts,omitempty"`
	ObservedScopeBindingTexts            []string `json:"observed_scope_binding_texts,omitempty"`
	ObservedEvidenceBindingTexts         []string `json:"observed_evidence_binding_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point14ValAFoundation struct {
	CurrentState                 string                                        `json:"current_state"`
	BlockingReasons              []string                                      `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites          []string                                      `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer         string                                        `json:"projection_disclaimer"`
	DependencyState              string                                        `json:"dependency_state"`
	NormalizedSignalState        string                                        `json:"normalized_signal_state"`
	SourceIdentityState          string                                        `json:"source_identity_state"`
	ScopeBindingState            string                                        `json:"scope_binding_state"`
	EvidenceBindingState         string                                        `json:"evidence_binding_state"`
	FreshnessTimestampState      string                                        `json:"freshness_timestamp_state"`
	DuplicateRelationGuardState  string                                        `json:"duplicate_relation_guard_state"`
	TenantBoundaryGuardState     string                                        `json:"tenant_boundary_guard_state"`
	ValidationResultState        string                                        `json:"validation_result_state"`
	NoExternalAuthorityState     string                                        `json:"no_external_authority_state"`
	NoOverclaimState             string                                        `json:"no_overclaim_state"`
	Dependency                   Point14ValADependencySnapshot                 `json:"dependency"`
	NormalizedExternalSignal     NormalizedExternalSignal                      `json:"normalized_external_signal"`
	SourceIdentity               ExternalSignalSourceIdentity                  `json:"source_identity"`
	ScopeBinding                 ExternalSignalScopeBinding                    `json:"scope_binding"`
	EvidenceBinding              ExternalSignalEvidenceBinding                 `json:"evidence_binding"`
	FreshnessAndTimestamp        ExternalSignalFreshnessAndTimestampBoundary   `json:"freshness_and_timestamp"`
	DuplicateAndRelationGuard    ExternalSignalDuplicateAndRelationGuard       `json:"duplicate_and_relation_guard"`
	TenantBoundaryGuard          ExternalSignalTenantBoundaryGuard             `json:"tenant_boundary_guard"`
	ValidationResult             ExternalSignalValidationResult                `json:"validation_result"`
	NoExternalAuthorityGuard     Point14ValANoExternalAuthorityValidationGuard `json:"no_external_authority_guard"`
	NoOverclaimValidationWording Point14ValANoOverclaimValidationWording       `json:"no_overclaim_validation_wording"`
}

func point14ValAStates() []string {
	return []string{
		Point14ValAStateActive,
		Point14ValAStateBlocked,
		Point14ValAStateReviewRequired,
		Point14ValAStateIncomplete,
	}
}

func point14ValAStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point14ValAStates(), value)
}

func point14ValAValidationStates() []string {
	return []string{
		point14ValAValidationCandidateValidated,
		point14ValAValidationCandidateBlocked,
		point14ValAValidationCandidateReviewRequired,
		point14ValAValidationCandidateIncomplete,
		point14ValAValidationCandidateUnsupported,
		point14ValAValidationCandidateTampered,
		point14ValAValidationCandidateDuplicate,
		point14ValAValidationCandidateUnrelated,
		point14ValAValidationCandidateCrossTenant,
	}
}

func point14ValAValidationStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point14ValAValidationStates(), value)
}

func point14ValAEvidenceStates() []string {
	return []string{
		point14ValAEvidenceStateActive,
		point14ValAEvidenceStateStale,
		point14ValAEvidenceStateRevoked,
		point14ValAEvidenceStateExpired,
		point14ValAEvidenceStateSuperseded,
		point14ValAEvidenceStateUnrelated,
	}
}

func point14ValAEvidenceStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point14ValAEvidenceStates(), value)
}

func point14ValASourceAuthorityScopes() []string {
	return []string{
		"tenant_scoped_advisory_source",
		"global_advisory_source",
		"public_non_authoritative_source",
	}
}

func point14ValASourceAuthorityScopeValid(value string) bool {
	return point11Val0ContainsTrimmed(point14ValASourceAuthorityScopes(), value)
}

func point14ValAForbiddenAuthorityMarkers() []string {
	markers := append([]string{}, point14Val0ForbiddenAuthorityMarkers()...)
	return append(markers,
		point14ValABlockedPassToken,
		"canonical_authority_granted",
		"production_approval_granted",
	)
}

func point14ValASafeWording() []string {
	return []string{
		"normalized external evidence signal",
		"validated candidate signal",
		"advisory external signal",
		"evidence input pending governance review",
		"candidate validation does not grant authority",
		"no external PASS authority",
		"bounded external signal input",
	}
}

func point14ValAObservedTextContainsForbiddenWording(text string) bool {
	trimmed := strings.ToLower(strings.TrimSpace(text))
	if trimmed == "" {
		return false
	}
	for _, safe := range point14ValASafeWording() {
		if trimmed == strings.ToLower(strings.TrimSpace(safe)) {
			return false
		}
	}
	return point14Val0ContainsForbiddenWording(text)
}

func point14ValAObservedListContainsForbiddenWording(values []string) bool {
	for _, value := range values {
		if point14ValAObservedTextContainsForbiddenWording(value) {
			return true
		}
	}
	return false
}

func point14ValANormalizedSignalIDValid(value string) bool {
	return point14Val0RefValid(value, "normalized_signal_")
}

func point14ValASourceIdentityIDValid(value string) bool {
	return point14Val0RefValid(value, "source_identity_")
}

func point14ValAScopeBindingIDValid(value string) bool {
	return point14Val0RefValid(value, "scope_binding_")
}

func point14ValAEvidenceBindingIDValid(value string) bool {
	return point14Val0RefValid(value, "evidence_binding_")
}

func point14ValAValidationResultIDValid(value string) bool {
	return point14Val0RefValid(value, "validation_result_")
}

func point14ValAHashRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "hash_")
}

func point14ValADuplicateGuardIDValid(value string) bool {
	return point14Val0RefValid(value, "duplicate_guard_")
}

func point14ValATenantGuardIDValid(value string) bool {
	return point14Val0RefValid(value, "tenant_boundary_guard_")
}

func point14ValAValidationStateClassification(value string) (string, bool, bool) {
	switch strings.TrimSpace(value) {
	case point14ValAValidationCandidateBlocked,
		point14ValAValidationCandidateUnsupported,
		point14ValAValidationCandidateTampered,
		point14ValAValidationCandidateDuplicate,
		point14ValAValidationCandidateUnrelated,
		point14ValAValidationCandidateCrossTenant:
		return Point14ValAStateBlocked, true, true
	case point14ValAValidationCandidateReviewRequired:
		return Point14ValAStateReviewRequired, false, true
	case point14ValAValidationCandidateIncomplete:
		return Point14ValAStateIncomplete, false, true
	case point14ValAValidationCandidateValidated:
		return Point14ValAStateActive, false, true
	default:
		return Point14ValAStateBlocked, false, false
	}
}

func point14ValAResolvedValidationState(current, aggregate string) string {
	current = strings.TrimSpace(current)
	currentState, blockedLike, recognized := point14ValAValidationStateClassification(current)
	resolvedAggregate := point14ValAFoundationState(currentState, strings.TrimSpace(aggregate))
	switch resolvedAggregate {
	case Point14ValAStateBlocked:
		if recognized && blockedLike {
			return current
		}
		return point14ValAValidationCandidateBlocked
	case Point14ValAStateReviewRequired:
		if recognized && blockedLike {
			return current
		}
		return point14ValAValidationCandidateReviewRequired
	case Point14ValAStateIncomplete:
		if recognized && (blockedLike || current == point14ValAValidationCandidateReviewRequired) {
			return current
		}
		return point14ValAValidationCandidateIncomplete
	default:
		if recognized {
			return current
		}
		return point14ValAValidationCandidateBlocked
	}
}

func point14ValADependencySnapshotFromUpstream(
	val0 Point14Val0Foundation,
) Point14ValADependencySnapshot {
	return Point14ValADependencySnapshot{
		Point14Val0CurrentState:                          val0.CurrentState,
		Point14Val0DependencyState:                       val0.DependencyState,
		Point14Val0ExternalSignalCandidateState:          val0.ExternalSignalCandidateState,
		Point14Val0ExternalStakeholderAuthorityRoleState: val0.ExternalStakeholderAuthorityRoleState,
		Point14Val0ExternalAuthorityConflictMatrixState:  val0.ExternalAuthorityConflictMatrixState,
		Point14Val0ExternalSignalDisputeLifecycleState:   val0.ExternalSignalDisputeLifecycleState,
		Point14Val0ExternalCorrectionRevocationState:     val0.ExternalCorrectionRevocationState,
		Point14Val0ExternalVisibilityPublicationState:    val0.ExternalVisibilityPublicationState,
		Point14Val0AgentEcosystemInputBoundaryState:      val0.AgentEcosystemInputBoundaryState,
		Point14Val0NoExternalAuthorityGuardState:         val0.NoExternalAuthorityGuardState,
		Point14Val0NoOverclaimState:                      val0.NoOverclaimState,
		Point14Val0PointID:                               point14Val0PointID,
		Point14Val0WaveID:                                point14Val0WaveID,
		Point14Val0ComputedFromUpstream:                  val0.Dependency.SnapshotFromComputedOutput,
		Point14Val0Merged:                                true,
		Point14Val0CIGreen:                               true,
		Point14Val0ReviewedOnMain:                        true,
		Point14PassSeen:                                  false,
		InheritedPoint13ValECurrentState:                 val0.Dependency.Point13ValECurrentState,
		InheritedPoint13ValEPassClosureState:             val0.Dependency.Point13ValEPassClosureManifestState,
		InheritedPoint13ValEPassAllowed:                  val0.Dependency.Point13ValEPassAllowed,
		InheritedPoint13ValEPassToken:                    val0.Dependency.Point13ValEPassToken,
		InheritedPoint12CurrentState:                     val0.Dependency.InheritedPoint12CurrentState,
		InheritedPoint12DependencyState:                  val0.Dependency.InheritedPoint12DependencyState,
		InheritedPoint12PassClosureState:                 val0.Dependency.InheritedPoint12PassClosureState,
		InheritedPoint12ReviewerResult:                   val0.Dependency.InheritedPoint12ReviewerResult,
		InheritedPoint11PublicationState:                 val0.Dependency.InheritedPoint11PublicationState,
		InheritedPoint11NoOverclaimState:                 val0.Dependency.InheritedPoint11NoOverclaimState,
		InheritedPoint11FinalPassGateState:               val0.Dependency.InheritedPoint11FinalPassGateState,
		InheritedPoint10CurrentState:                     val0.Dependency.InheritedPoint10CurrentState,
		InheritedPoint10NoOverclaimState:                 val0.Dependency.InheritedPoint10NoOverclaimState,
		InheritedPoint10ProjectionState:                  val0.Dependency.InheritedPoint10ProjectionState,
		InheritedPoint10PassRuleState:                    val0.Dependency.InheritedPoint10PassRuleState,
		InheritedTenantScope:                             val0.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:                       true,
		Point14Val0:                                      val0,
		Point13ValE:                                      val0.Dependency.Point13ValE,
		Point12:                                          val0.Dependency.Point12,
		Point11:                                          val0.Dependency.Point11,
		Point10:                                          val0.Dependency.Point10,
	}
}

func point14ValADependencySnapshotModel() Point14ValADependencySnapshot {
	val0 := ComputePoint14Val0Foundation(Point14Val0FoundationModel())
	return point14ValADependencySnapshotFromUpstream(val0)
}

func EvaluatePoint14ValADependencyState(model Point14ValADependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point14Val0ComputedFromUpstream ||
		!model.Point14Val0Merged ||
		!model.Point14Val0CIGreen ||
		!model.Point14Val0ReviewedOnMain ||
		model.Point14PassSeen ||
		model.Point14Val0PointID != point14Val0PointID ||
		model.Point14Val0WaveID != point14Val0WaveID ||
		!point14Val0StateValid(model.Point14Val0CurrentState) ||
		!point14Val0StateValid(model.Point14Val0DependencyState) ||
		!point14Val0StateValid(model.Point14Val0ExternalSignalCandidateState) ||
		!point14Val0StateValid(model.Point14Val0ExternalStakeholderAuthorityRoleState) ||
		!point14Val0StateValid(model.Point14Val0ExternalAuthorityConflictMatrixState) ||
		!point14Val0StateValid(model.Point14Val0ExternalSignalDisputeLifecycleState) ||
		!point14Val0StateValid(model.Point14Val0ExternalCorrectionRevocationState) ||
		!point14Val0StateValid(model.Point14Val0ExternalVisibilityPublicationState) ||
		!point14Val0StateValid(model.Point14Val0AgentEcosystemInputBoundaryState) ||
		!point14Val0StateValid(model.Point14Val0NoExternalAuthorityGuardState) ||
		!point14Val0StateValid(model.Point14Val0NoOverclaimState) ||
		!point13ValEStateValid(model.InheritedPoint13ValECurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValEPassClosureState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		!point12ValEReviewerResultValid(model.InheritedPoint12ReviewerResult) ||
		model.InheritedPoint11PublicationState == "" ||
		model.InheritedPoint11NoOverclaimState == "" ||
		model.InheritedPoint11FinalPassGateState == "" ||
		model.InheritedPoint10CurrentState == "" ||
		model.InheritedPoint10NoOverclaimState == "" ||
		model.InheritedPoint10ProjectionState == "" ||
		model.InheritedPoint10PassRuleState == "" ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return Point14ValAStateBlocked
	}
	if model.Point14Val0CurrentState != model.Point14Val0.CurrentState ||
		model.Point14Val0DependencyState != model.Point14Val0.DependencyState ||
		model.Point14Val0ExternalSignalCandidateState != model.Point14Val0.ExternalSignalCandidateState ||
		model.Point14Val0ExternalStakeholderAuthorityRoleState != model.Point14Val0.ExternalStakeholderAuthorityRoleState ||
		model.Point14Val0ExternalAuthorityConflictMatrixState != model.Point14Val0.ExternalAuthorityConflictMatrixState ||
		model.Point14Val0ExternalSignalDisputeLifecycleState != model.Point14Val0.ExternalSignalDisputeLifecycleState ||
		model.Point14Val0ExternalCorrectionRevocationState != model.Point14Val0.ExternalCorrectionRevocationState ||
		model.Point14Val0ExternalVisibilityPublicationState != model.Point14Val0.ExternalVisibilityPublicationState ||
		model.Point14Val0AgentEcosystemInputBoundaryState != model.Point14Val0.AgentEcosystemInputBoundaryState ||
		model.Point14Val0NoExternalAuthorityGuardState != model.Point14Val0.NoExternalAuthorityGuardState ||
		model.Point14Val0NoOverclaimState != model.Point14Val0.NoOverclaimState ||
		model.Point14Val0ComputedFromUpstream != model.Point14Val0.Dependency.SnapshotFromComputedOutput ||
		model.InheritedPoint13ValECurrentState != model.Point14Val0.Dependency.Point13ValECurrentState ||
		model.InheritedPoint13ValEPassClosureState != model.Point14Val0.Dependency.Point13ValEPassClosureManifestState ||
		model.InheritedPoint13ValEPassAllowed != model.Point14Val0.Dependency.Point13ValEPassAllowed ||
		model.InheritedPoint13ValEPassToken != model.Point14Val0.Dependency.Point13ValEPassToken ||
		model.InheritedPoint12CurrentState != model.Point14Val0.Dependency.InheritedPoint12CurrentState ||
		model.InheritedPoint12DependencyState != model.Point14Val0.Dependency.InheritedPoint12DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point14Val0.Dependency.InheritedPoint12PassClosureState ||
		model.InheritedPoint12ReviewerResult != model.Point14Val0.Dependency.InheritedPoint12ReviewerResult ||
		model.InheritedPoint11PublicationState != model.Point14Val0.Dependency.InheritedPoint11PublicationState ||
		model.InheritedPoint11NoOverclaimState != model.Point14Val0.Dependency.InheritedPoint11NoOverclaimState ||
		model.InheritedPoint11FinalPassGateState != model.Point14Val0.Dependency.InheritedPoint11FinalPassGateState ||
		model.InheritedPoint10CurrentState != model.Point14Val0.Dependency.InheritedPoint10CurrentState ||
		model.InheritedPoint10NoOverclaimState != model.Point14Val0.Dependency.InheritedPoint10NoOverclaimState ||
		model.InheritedPoint10ProjectionState != model.Point14Val0.Dependency.InheritedPoint10ProjectionState ||
		model.InheritedPoint10PassRuleState != model.Point14Val0.Dependency.InheritedPoint10PassRuleState ||
		model.InheritedTenantScope != model.Point14Val0.Dependency.InheritedTenantScope ||
		model.InheritedPoint13ValECurrentState != model.Point13ValE.CurrentState ||
		model.InheritedPoint13ValEPassClosureState != model.Point13ValE.PassClosureManifestState ||
		model.InheritedPoint13ValEPassAllowed != model.Point13ValE.Point13PassAllowed ||
		model.InheritedPoint13ValEPassToken != model.Point13ValE.Point13PassToken ||
		model.InheritedPoint12CurrentState != model.Point12.CurrentState ||
		model.InheritedPoint12DependencyState != model.Point12.DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point12.PassClosureManifestState ||
		model.InheritedPoint12ReviewerResult != model.Point12.PassClosureManifest.ReviewerResult ||
		model.InheritedPoint11PublicationState != model.Point11.PublicationReviewState ||
		model.InheritedPoint11NoOverclaimState != model.Point11.NoOverclaimReviewState ||
		model.InheritedPoint11FinalPassGateState != model.Point11.FinalPassGateState ||
		model.InheritedPoint10CurrentState != model.Point10.Point10State ||
		model.InheritedPoint10NoOverclaimState != model.Point10.NoOverclaimState ||
		model.InheritedPoint10ProjectionState != model.Point10.ProjectionBoundaryState ||
		model.InheritedPoint10PassRuleState != model.Point10.Point10PassRuleState {
		return Point14ValAStateBlocked
	}
	if model.Point14Val0CurrentState != Point14Val0StateActive ||
		model.Point14Val0DependencyState != Point14Val0StateActive ||
		model.Point14Val0ExternalSignalCandidateState != Point14Val0StateActive ||
		model.Point14Val0ExternalStakeholderAuthorityRoleState != Point14Val0StateActive ||
		model.Point14Val0ExternalAuthorityConflictMatrixState != Point14Val0StateActive ||
		model.Point14Val0ExternalSignalDisputeLifecycleState != Point14Val0StateActive ||
		model.Point14Val0ExternalCorrectionRevocationState != Point14Val0StateActive ||
		model.Point14Val0ExternalVisibilityPublicationState != Point14Val0StateActive ||
		model.Point14Val0AgentEcosystemInputBoundaryState != Point14Val0StateActive ||
		model.Point14Val0NoExternalAuthorityGuardState != Point14Val0StateActive ||
		model.Point14Val0NoOverclaimState != Point14Val0StateActive ||
		model.InheritedPoint13ValECurrentState != Point13ValEStatePassConfirmed ||
		model.InheritedPoint13ValEPassClosureState != Point13ValEStateActive ||
		!model.InheritedPoint13ValEPassAllowed ||
		model.InheritedPoint13ValEPassToken != point13ValEPoint13PassToken ||
		model.InheritedPoint12CurrentState != Point12ValEStatePassConfirmed ||
		model.InheritedPoint12DependencyState != Point12ValEStateActive ||
		model.InheritedPoint12PassClosureState != Point12ValEStateActive ||
		model.InheritedPoint12ReviewerResult != point12ValEReviewerResultPassConfirmed ||
		model.InheritedPoint11PublicationState != Point11ValDPublicationReviewStateActive ||
		model.InheritedPoint11NoOverclaimState != Point11ValDNoOverclaimReviewStateActive ||
		model.InheritedPoint10CurrentState != operability.DeploymentMultiTenantPoint10StatePass ||
		model.InheritedPoint10NoOverclaimState != operability.DeploymentMultiTenantValENoOverclaimStateActive ||
		model.InheritedPoint10ProjectionState != operability.DeploymentMultiTenantValEProjectionBoundaryStateActive ||
		model.InheritedPoint10PassRuleState != operability.DeploymentMultiTenantValEPoint10PassRuleStateActive {
		return Point14ValAStateBlocked
	}
	return Point14ValAStateActive
}

func point14ValANormalizedExternalSignalModel(dependency Point14ValADependencySnapshot) NormalizedExternalSignal {
	return NormalizedExternalSignal{
		NormalizedSignalID:      "normalized_signal_point14_vala_001",
		OriginalSignalID:        "signal_point14_vala_original_001",
		SourceIdentityRef:       "source_identity_point14_vala_001",
		SourceType:              "vendor_advisory",
		SignalType:              "evidence_submission",
		NormalizedPayloadHash:   "hash_point14_vala_normalized_payload_001",
		TenantScope:             dependency.InheritedTenantScope,
		ArtifactRefs:            []string{"artifact_point14_vala_component_001"},
		ArtifactScoped:          true,
		ClaimRefs:               []string{"claim_point14_vala_001"},
		ClaimScoped:             true,
		EvidenceRefs:            []string{"evidence_point14_vala_001"},
		EvidenceLinked:          true,
		ReceivedAt:              "2026-05-05T23:05:00Z",
		NormalizationTimeSource: point14Val0TimeSourceServerUTC,
		ValidationStatus:        point14Val0ValidationValidated,
		PayloadNormalized:       true,
	}
}

func EvaluatePoint14ValANormalizedExternalSignalState(model NormalizedExternalSignal, dependency Point14ValADependencySnapshot) string {
	if !point14ValANormalizedSignalIDValid(model.NormalizedSignalID) ||
		!point14Val0SignalIDValid(model.OriginalSignalID) ||
		!point14ValASourceIdentityIDValid(model.SourceIdentityRef) ||
		!point14Val0ExactValueValid(model.SourceType, point14Val0SourceTypes()) ||
		!point14Val0ExactValueValid(model.SignalType, point14Val0SignalTypes()) ||
		!point14Val0HashRefValid(model.NormalizedPayloadHash) ||
		!point14Val0TimeSourceValid(model.NormalizationTimeSource) ||
		!point14Val0CanonicalTimeSourceValid(model.NormalizationTimeSource) ||
		!point14Val0ExactValueValid(model.ValidationStatus, point14Val0ValidationStatuses()) ||
		!point14Val0ParsedTimeOk(model.ReceivedAt) ||
		!model.PayloadNormalized {
		return Point14ValAStateBlocked
	}
	if model.TenantScope == "" && strings.TrimSpace(model.GlobalScopeClassification) == "" {
		return Point14ValAStateBlocked
	}
	if model.TenantScope != "" {
		if !point11Val0ScopeValid(model.TenantScope) || model.TenantScope != dependency.InheritedTenantScope {
			return Point14ValAStateBlocked
		}
		if strings.TrimSpace(model.GlobalScopeClassification) != "" {
			return Point14ValAStateBlocked
		}
	} else if !point11Val0ContainsTrimmed([]string{point14Val0ScopeGlobalAdvisory, point14Val0ScopePublicNonAuthorative}, model.GlobalScopeClassification) {
		return Point14ValAStateBlocked
	}
	if model.ArtifactScoped && !point14Val0RefListValid(model.ArtifactRefs, "artifact_") {
		return Point14ValAStateBlocked
	}
	if model.ClaimScoped && !point14Val0ClaimRefsValid(model.ClaimRefs) {
		return Point14ValAStateBlocked
	}
	if model.EvidenceLinked && !point14Val0EvidenceRefsValid(model.EvidenceRefs) {
		return Point14ValAStateBlocked
	}
	if model.CanonicalAuthority ||
		model.PassAllowed ||
		model.CorrectionPublished ||
		model.ProductionApproved ||
		model.PublicBadgeAllowed {
		return Point14ValAStateBlocked
	}
	switch strings.TrimSpace(model.ValidationStatus) {
	case point14Val0ValidationProvenancePending, point14Val0ValidationConflicting:
		return Point14ValAStateReviewRequired
	case point14Val0ValidationRevoked, point14Val0ValidationSuperseded, point14Val0ValidationRejected:
		return Point14ValAStateBlocked
	}
	return Point14ValAStateActive
}

func point14ValASourceIdentityModel() ExternalSignalSourceIdentity {
	return ExternalSignalSourceIdentity{
		SourceIdentityID:     "source_identity_point14_vala_001",
		SourceType:           "vendor_advisory",
		SourceRef:            "source_point14_vala_vendor_001",
		SourceAuthorityScope: "tenant_scoped_advisory_source",
		ProvenanceRef:        "provenance_point14_vala_001",
		CustodyRef:           "custody_point14_vala_001",
		HashRef:              "hash_point14_vala_source_001",
		SignatureRef:         "signature_point14_vala_001",
		RequireProvenance:    true,
		RequireCustody:       true,
		RequireHash:          true,
		RequireSignature:     true,
		SourceSupported:      true,
		SourceTrusted:        true,
	}
}

func EvaluatePoint14ValAExternalSignalSourceIdentityState(model ExternalSignalSourceIdentity) string {
	if !point14ValASourceIdentityIDValid(model.SourceIdentityID) ||
		!point14Val0ExactValueValid(model.SourceType, point14Val0SourceTypes()) ||
		!point14Val0SourceRefValid(model.SourceRef) ||
		!point14ValASourceAuthorityScopeValid(model.SourceAuthorityScope) {
		return Point14ValAStateBlocked
	}
	if model.RequireProvenance && !point14Val0ProvenanceRefValid(model.ProvenanceRef) {
		return Point14ValAStateBlocked
	}
	if model.RequireCustody && !point14Val0CustodyRefValid(model.CustodyRef) {
		return Point14ValAStateBlocked
	}
	if model.RequireHash && !point14Val0HashRefValid(model.HashRef) {
		return Point14ValAStateBlocked
	}
	if model.RequireSignature && !point14Val0SignatureRefValid(model.SignatureRef) {
		return Point14ValAStateBlocked
	}
	if model.CanonicalAuthorityGranted ||
		model.PassAllowed ||
		model.OverrideCanonicalDecision ||
		model.ApproveProduction ||
		model.CertifyCompliance {
		return Point14ValAStateBlocked
	}
	if !model.SourceSupported || !model.SourceTrusted {
		return Point14ValAStateReviewRequired
	}
	return Point14ValAStateActive
}

func point14ValAScopeBindingModel(dependency Point14ValADependencySnapshot) ExternalSignalScopeBinding {
	return ExternalSignalScopeBinding{
		ScopeBindingID:        "scope_binding_point14_vala_001",
		ScopeClassification:   point14Val0ScopeTenantScoped,
		TenantScope:           dependency.InheritedTenantScope,
		ReferencedTenantScope: dependency.InheritedTenantScope,
		ArtifactScoped:        true,
		ArtifactRefs:          []string{"artifact_point14_vala_component_001"},
		ArtifactHashRefs:      []string{"hash_point14_vala_artifact_001"},
		ArtifactHashesMatch:   true,
		ClaimScoped:           true,
		ClaimRefs:             []string{"claim_point14_vala_001"},
		ClaimBindingExact:     true,
	}
}

func EvaluatePoint14ValAExternalSignalScopeBindingState(model ExternalSignalScopeBinding, dependency Point14ValADependencySnapshot) string {
	if !point14ValAScopeBindingIDValid(model.ScopeBindingID) ||
		!point14Val0ExactValueValid(model.ScopeClassification, point14Val0ScopeClassifications()) {
		return Point14ValAStateBlocked
	}
	if strings.TrimSpace(model.ScopeClassification) == point14Val0ScopeTenantScoped {
		if !point11Val0ScopeValid(model.TenantScope) ||
			model.TenantScope != dependency.InheritedTenantScope ||
			strings.TrimSpace(model.GlobalScopeClassification) != "" {
			return Point14ValAStateBlocked
		}
	} else {
		if strings.TrimSpace(model.TenantScope) != "" ||
			!point11Val0ContainsTrimmed([]string{point14Val0ScopeGlobalAdvisory, point14Val0ScopePublicNonAuthorative}, model.GlobalScopeClassification) {
			return Point14ValAStateBlocked
		}
	}
	if model.ReferencedTenantScope != "" &&
		model.ReferencedTenantScope != dependency.InheritedTenantScope {
		return Point14ValAStateBlocked
	}
	if model.ArtifactScoped {
		if !point14Val0RefListValid(model.ArtifactRefs, "artifact_") || !point14ValAHashRefsValid(model.ArtifactHashRefs) || !model.ArtifactHashesMatch {
			return Point14ValAStateBlocked
		}
	}
	if model.ClaimScoped {
		if !point14Val0ClaimRefsValid(model.ClaimRefs) || !model.ClaimBindingExact {
			return Point14ValAStateBlocked
		}
	}
	if model.SimilarNameOnly || model.SimilarPathOnly || model.SimilarPackageOnly {
		return Point14ValAStateBlocked
	}
	return Point14ValAStateActive
}

func point14ValAEvidenceBindingModel(dependency Point14ValADependencySnapshot) ExternalSignalEvidenceBinding {
	return ExternalSignalEvidenceBinding{
		EvidenceBindingID:   "evidence_binding_point14_vala_001",
		TenantScope:         dependency.InheritedTenantScope,
		EvidenceLinked:      true,
		EvidenceRefs:        []string{"evidence_point14_vala_001"},
		EvidenceHashRefs:    []string{"hash_point14_vala_evidence_001"},
		EvidenceHashesMatch: true,
		EvidenceState:       point14ValAEvidenceStateActive,
		ProvenanceRef:       "provenance_point14_vala_001",
		CustodyRef:          "custody_point14_vala_001",
		RequireProvenance:   true,
		RequireCustody:      true,
		RequireHash:         true,
	}
}

func EvaluatePoint14ValAExternalSignalEvidenceBindingState(model ExternalSignalEvidenceBinding) string {
	if !point14ValAEvidenceBindingIDValid(model.EvidenceBindingID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point14ValAEvidenceStateValid(model.EvidenceState) {
		return Point14ValAStateBlocked
	}
	if model.EvidenceLinked && !point14Val0EvidenceRefsValid(model.EvidenceRefs) {
		return Point14ValAStateBlocked
	}
	if model.RequireHash && (!point14ValAHashRefsValid(model.EvidenceHashRefs) || !model.EvidenceHashesMatch) {
		return Point14ValAStateBlocked
	}
	if model.RequireProvenance && !point14Val0ProvenanceRefValid(model.ProvenanceRef) {
		return Point14ValAStateBlocked
	}
	if model.RequireCustody && !point14Val0CustodyRefValid(model.CustodyRef) {
		return Point14ValAStateBlocked
	}
	if model.CanonicalMutationAllowed {
		return Point14ValAStateBlocked
	}
	switch strings.TrimSpace(model.EvidenceState) {
	case point14ValAEvidenceStateStale, point14ValAEvidenceStateSuperseded:
		return Point14ValAStateReviewRequired
	case point14ValAEvidenceStateRevoked, point14ValAEvidenceStateExpired, point14ValAEvidenceStateUnrelated:
		return Point14ValAStateBlocked
	}
	return Point14ValAStateActive
}

func point14ValAFreshnessAndTimestampBoundaryModel(dependency Point14ValADependencySnapshot) ExternalSignalFreshnessAndTimestampBoundary {
	return ExternalSignalFreshnessAndTimestampBoundary{
		BoundaryID:                 "boundary_point14_vala_freshness_001",
		TenantScope:                dependency.InheritedTenantScope,
		ReceivedAt:                 "2026-05-05T23:05:00Z",
		ReceivedTimeSource:         point14Val0TimeSourceServerUTC,
		SourceEventAt:              "2026-05-05T22:50:00Z",
		SourceEventTimeSource:      point14Val0TimeSourceApprovedCustomerTime,
		ValidatedAt:                "2026-05-05T23:10:00Z",
		ValidatedTimeSource:        point14Val0TimeSourceServerUTC,
		SourceEventAdvisoryOnly:    true,
		ClientTimeMetadataOnly:     true,
		GovernanceReviewPathExists: false,
	}
}

func EvaluatePoint14ValAExternalSignalFreshnessAndTimestampBoundaryState(model ExternalSignalFreshnessAndTimestampBoundary) string {
	if !point14Val0BoundaryRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point14Val0ParsedTimeOk(model.ReceivedAt) ||
		!point14Val0TimeSourceValid(model.ReceivedTimeSource) ||
		!point14Val0CanonicalTimeSourceValid(model.ReceivedTimeSource) ||
		!model.SourceEventAdvisoryOnly ||
		model.AuthorityUpgradeRequested {
		return Point14ValAStateBlocked
	}
	receivedAt, _ := point14Val0ParsedTime(model.ReceivedAt)
	if strings.TrimSpace(model.SourceEventAt) != "" {
		if !point14Val0ParsedTimeOk(model.SourceEventAt) || !point14Val0TimeSourceValid(model.SourceEventTimeSource) {
			return Point14ValAStateBlocked
		}
		sourceEventAt, _ := point14Val0ParsedTime(model.SourceEventAt)
		if sourceEventAt.After(receivedAt) {
			return Point14ValAStateReviewRequired
		}
	}
	if strings.TrimSpace(model.ValidatedAt) != "" {
		if !point14Val0ParsedTimeOk(model.ValidatedAt) || !point14Val0TimeSourceValid(model.ValidatedTimeSource) || !point14Val0CanonicalTimeSourceValid(model.ValidatedTimeSource) {
			return Point14ValAStateBlocked
		}
		validatedAt, _ := point14Val0ParsedTime(model.ValidatedAt)
		if validatedAt.Before(receivedAt) {
			return Point14ValAStateReviewRequired
		}
	}
	if strings.TrimSpace(model.ClientReportedAt) != "" && !model.ClientTimeMetadataOnly {
		return Point14ValAStateBlocked
	}
	if model.StaleSignal {
		if model.GovernanceReviewPathExists {
			return Point14ValAStateReviewRequired
		}
		return Point14ValAStateBlocked
	}
	return Point14ValAStateActive
}

func point14ValADuplicateAndRelationGuardModel(dependency Point14ValADependencySnapshot) ExternalSignalDuplicateAndRelationGuard {
	return ExternalSignalDuplicateAndRelationGuard{
		GuardID:                       "duplicate_guard_point14_vala_001",
		TenantScope:                   dependency.InheritedTenantScope,
		NormalizedSignalRef:           "normalized_signal_point14_vala_001",
		OriginalSignalRef:             "signal_point14_vala_original_001",
		DuplicateIdentityKey:          "signal_identity_point14_vala_duplicate_001",
		ExpectedNormalizedPayloadHash: "hash_point14_vala_normalized_payload_001",
		ObservedNormalizedPayloadHash: "hash_point14_vala_normalized_payload_001",
		RelatedArtifactRefs:           []string{"artifact_point14_vala_component_001"},
		RelatedClaimRefs:              []string{"claim_point14_vala_001"},
		ArtifactRelationExact:         true,
		ClaimRelationExact:            true,
	}
}

func EvaluatePoint14ValAExternalSignalDuplicateAndRelationGuardState(model ExternalSignalDuplicateAndRelationGuard) string {
	if !point14ValADuplicateGuardIDValid(model.GuardID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point14ValANormalizedSignalIDValid(model.NormalizedSignalRef) ||
		!point14Val0SignalIDValid(model.OriginalSignalRef) ||
		!point14Val0IdentityKeyValid(model.DuplicateIdentityKey) ||
		!point14Val0HashRefValid(model.ExpectedNormalizedPayloadHash) ||
		!point14Val0HashRefValid(model.ObservedNormalizedPayloadHash) ||
		len(model.RelatedArtifactRefs) == 0 ||
		!point14Val0RefListValid(model.RelatedArtifactRefs, "artifact_") ||
		len(model.RelatedClaimRefs) == 0 ||
		!point14Val0ClaimRefsValid(model.RelatedClaimRefs) {
		return Point14ValAStateBlocked
	}
	if !model.ArtifactRelationExact || !model.ClaimRelationExact {
		return Point14ValAStateBlocked
	}
	if model.ExpectedNormalizedPayloadHash != model.ObservedNormalizedPayloadHash {
		return Point14ValAStateBlocked
	}
	if model.SilentReplacementRequested || model.ConflictingDuplicate {
		return Point14ValAStateBlocked
	}
	if len(model.DuplicateSignalRefs) > 0 {
		if !point14Val0SignalRefsValid(model.DuplicateSignalRefs) {
			return Point14ValAStateBlocked
		}
		return Point14ValAStateReviewRequired
	}
	return Point14ValAStateActive
}

func point14ValATenantBoundaryGuardModel(dependency Point14ValADependencySnapshot) ExternalSignalTenantBoundaryGuard {
	return ExternalSignalTenantBoundaryGuard{
		GuardID:             "tenant_boundary_guard_point14_vala_001",
		TenantScope:         dependency.InheritedTenantScope,
		SourceTenantScope:   dependency.InheritedTenantScope,
		ArtifactTenantScope: dependency.InheritedTenantScope,
		ClaimTenantScope:    dependency.InheritedTenantScope,
		EvidenceTenantScope: dependency.InheritedTenantScope,
		ScopeClassification: point14Val0ScopeTenantScoped,
	}
}

func EvaluatePoint14ValAExternalSignalTenantBoundaryGuardState(model ExternalSignalTenantBoundaryGuard) string {
	if !point14ValATenantGuardIDValid(model.GuardID) ||
		!point14Val0ExactValueValid(model.ScopeClassification, point14Val0ScopeClassifications()) ||
		model.TenantPrivateDataExposed {
		return Point14ValAStateBlocked
	}
	if strings.TrimSpace(model.ScopeClassification) == point14Val0ScopeTenantScoped {
		if !point11Val0ScopeValid(model.TenantScope) ||
			model.SourceTenantScope != model.TenantScope ||
			model.ArtifactTenantScope != model.TenantScope ||
			model.ClaimTenantScope != model.TenantScope ||
			model.EvidenceTenantScope != model.TenantScope {
			return Point14ValAStateBlocked
		}
		return Point14ValAStateActive
	}
	if !point11Val0ContainsTrimmed([]string{point14Val0ScopeGlobalAdvisory, point14Val0ScopePublicNonAuthorative}, model.ScopeClassification) {
		return Point14ValAStateBlocked
	}
	if !model.ExplicitBoundedRuleExists &&
		(strings.TrimSpace(model.ArtifactTenantScope) != "" ||
			strings.TrimSpace(model.ClaimTenantScope) != "" ||
			strings.TrimSpace(model.EvidenceTenantScope) != "") {
		return Point14ValAStateBlocked
	}
	return Point14ValAStateActive
}

func point14ValAValidationResultModel() ExternalSignalValidationResult {
	return ExternalSignalValidationResult{
		ValidationResultID:   "validation_result_point14_vala_001",
		NormalizedSignalRef:  "normalized_signal_point14_vala_001",
		SourceIdentityRef:    "source_identity_point14_vala_001",
		ScopeBindingRef:      "scope_binding_point14_vala_001",
		EvidenceBindingRef:   "evidence_binding_point14_vala_001",
		FreshnessBoundaryRef: "boundary_point14_vala_freshness_001",
		DuplicateGuardRef:    "duplicate_guard_point14_vala_001",
		TenantBoundaryRef:    "tenant_boundary_guard_point14_vala_001",
		ValidationState:      point14ValAValidationCandidateValidated,
		CandidateUsable:      true,
	}
}

func EvaluatePoint14ValAExternalSignalValidationResultState(model ExternalSignalValidationResult) string {
	if !point14ValAValidationResultIDValid(model.ValidationResultID) ||
		!point14ValANormalizedSignalIDValid(model.NormalizedSignalRef) ||
		!point14ValASourceIdentityIDValid(model.SourceIdentityRef) ||
		!point14ValAScopeBindingIDValid(model.ScopeBindingRef) ||
		!point14ValAEvidenceBindingIDValid(model.EvidenceBindingRef) ||
		!point14Val0BoundaryRefValid(model.FreshnessBoundaryRef) ||
		!point14ValADuplicateGuardIDValid(model.DuplicateGuardRef) ||
		!point14ValATenantGuardIDValid(model.TenantBoundaryRef) ||
		!point14ValAStateValid(model.NormalizedSignalState) ||
		!point14ValAStateValid(model.SourceIdentityState) ||
		!point14ValAStateValid(model.ScopeBindingState) ||
		!point14ValAStateValid(model.EvidenceBindingState) ||
		!point14ValAStateValid(model.FreshnessTimestampState) ||
		!point14ValAStateValid(model.DuplicateRelationGuardState) ||
		!point14ValAStateValid(model.TenantBoundaryGuardState) ||
		!point14ValAStateValid(model.NoExternalAuthorityState) ||
		!point14ValAStateValid(model.NoOverclaimState) ||
		!point14ValAValidationStateValid(model.ValidationState) {
		return Point14ValAStateBlocked
	}
	if model.EmitsPass ||
		model.PublishesCorrection ||
		model.RevokesClaim ||
		model.OverridesCanonicalDecision ||
		model.ApprovesProduction ||
		model.CertifiesCompliance ||
		model.CreatesPublicBadge {
		return Point14ValAStateBlocked
	}
	componentAggregate := point14ValAFoundationState(
		model.NormalizedSignalState,
		model.SourceIdentityState,
		model.ScopeBindingState,
		model.EvidenceBindingState,
		model.FreshnessTimestampState,
		model.DuplicateRelationGuardState,
		model.TenantBoundaryGuardState,
		model.NoExternalAuthorityState,
		model.NoOverclaimState,
	)
	if componentAggregate == Point14ValAStateBlocked {
		return Point14ValAStateBlocked
	}
	if componentAggregate == Point14ValAStateReviewRequired {
		if strings.TrimSpace(model.ValidationState) == point14ValAValidationCandidateValidated {
			return Point14ValAStateBlocked
		}
		return Point14ValAStateReviewRequired
	}
	if componentAggregate == Point14ValAStateIncomplete {
		if strings.TrimSpace(model.ValidationState) == point14ValAValidationCandidateValidated {
			return Point14ValAStateBlocked
		}
		return Point14ValAStateIncomplete
	}
	switch strings.TrimSpace(model.ValidationState) {
	case point14ValAValidationCandidateBlocked,
		point14ValAValidationCandidateUnsupported,
		point14ValAValidationCandidateTampered,
		point14ValAValidationCandidateDuplicate,
		point14ValAValidationCandidateUnrelated,
		point14ValAValidationCandidateCrossTenant:
		return Point14ValAStateBlocked
	case point14ValAValidationCandidateReviewRequired:
		return Point14ValAStateReviewRequired
	case point14ValAValidationCandidateIncomplete:
		return Point14ValAStateIncomplete
	}
	if !model.CandidateUsable {
		return Point14ValAStateBlocked
	}
	return Point14ValAStateActive
}

func point14ValANoExternalAuthorityValidationGuardModel() Point14ValANoExternalAuthorityValidationGuard {
	return Point14ValANoExternalAuthorityValidationGuard{}
}

func EvaluatePoint14ValANoExternalAuthorityValidationGuardState(model Point14ValANoExternalAuthorityValidationGuard) string {
	for _, marker := range model.ObservedAuthorityMarkers {
		if point11Val0ContainsTrimmed(point14ValAForbiddenAuthorityMarkers(), marker) {
			return Point14ValAStateBlocked
		}
	}
	if model.CanonicalAuthorityGranted ||
		model.ProductionApprovalGranted ||
		model.CorrectionPublished ||
		model.PublicBadgeAllowed ||
		model.ExternalAuthorityAllowed {
		return Point14ValAStateBlocked
	}
	return Point14ValAStateActive
}

func point14ValANoOverclaimValidationWordingModel() Point14ValANoOverclaimValidationWording {
	return Point14ValANoOverclaimValidationWording{
		ObservedNormalizationTexts:   []string{"normalized external evidence signal"},
		ObservedValidationTexts:      []string{"validated candidate signal", "candidate validation does not grant authority"},
		ObservedSourceIdentityTexts:  []string{"advisory external signal"},
		ObservedScopeBindingTexts:    []string{"bounded external signal input"},
		ObservedEvidenceBindingTexts: []string{"evidence input pending governance review", "no external PASS authority"},
		AllowedSafeWording:           point14ValASafeWording(),
		BlockedWording:               point14Val0ForbiddenWording(),
		ProjectionDisclaimer:         point14ValAProjectionDisclaimerBase,
	}
}

func EvaluatePoint14ValANoOverclaimValidationWordingState(model Point14ValANoOverclaimValidationWording) string {
	if model.ProjectionDisclaimer != point14ValAProjectionDisclaimerBase ||
		!point14Val0TextListValid(model.AllowedSafeWording) ||
		!point14Val0TextListValid(model.BlockedWording) {
		return Point14ValAStateBlocked
	}
	if point14ValAObservedListContainsForbiddenWording(model.ObservedNormalizationTexts) ||
		point14ValAObservedListContainsForbiddenWording(model.ObservedValidationTexts) ||
		point14ValAObservedListContainsForbiddenWording(model.ObservedSourceIdentityTexts) ||
		point14ValAObservedListContainsForbiddenWording(model.ObservedScopeBindingTexts) ||
		point14ValAObservedListContainsForbiddenWording(model.ObservedEvidenceBindingTexts) {
		return Point14ValAStateBlocked
	}
	if point14Val0ListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point14ValAStateBlocked
	}
	return Point14ValAStateActive
}

func point14ValAFoundationState(states ...string) string {
	hasReview := false
	hasIncomplete := false
	for _, state := range states {
		switch strings.TrimSpace(state) {
		case Point14ValAStateBlocked:
			return Point14ValAStateBlocked
		case Point14ValAStateReviewRequired:
			hasReview = true
		case Point14ValAStateIncomplete:
			hasIncomplete = true
		}
	}
	if hasReview {
		return Point14ValAStateReviewRequired
	}
	if hasIncomplete {
		return Point14ValAStateIncomplete
	}
	return Point14ValAStateActive
}

func point14ValABlockingReasons(model Point14ValAFoundation) []string {
	componentStates := map[string]string{
		"dependency":               model.DependencyState,
		"normalized_signal":        model.NormalizedSignalState,
		"source_identity":          model.SourceIdentityState,
		"scope_binding":            model.ScopeBindingState,
		"evidence_binding":         model.EvidenceBindingState,
		"freshness_timestamp":      model.FreshnessTimestampState,
		"duplicate_relation_guard": model.DuplicateRelationGuardState,
		"tenant_boundary_guard":    model.TenantBoundaryGuardState,
		"validation_result":        model.ValidationResultState,
		"no_external_authority":    model.NoExternalAuthorityState,
		"no_overclaim":             model.NoOverclaimState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if state == Point14ValAStateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point14ValAReviewPrerequisites(model Point14ValAFoundation) []string {
	componentStates := map[string]string{
		"dependency":               model.DependencyState,
		"normalized_signal":        model.NormalizedSignalState,
		"source_identity":          model.SourceIdentityState,
		"scope_binding":            model.ScopeBindingState,
		"evidence_binding":         model.EvidenceBindingState,
		"freshness_timestamp":      model.FreshnessTimestampState,
		"duplicate_relation_guard": model.DuplicateRelationGuardState,
		"tenant_boundary_guard":    model.TenantBoundaryGuardState,
		"validation_result":        model.ValidationResultState,
		"no_external_authority":    model.NoExternalAuthorityState,
		"no_overclaim":             model.NoOverclaimState,
	}
	prereqs := []string{}
	for name, state := range componentStates {
		if state == Point14ValAStateReviewRequired || state == Point14ValAStateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func point14ValAFoundationModelFromUpstream(val0 Point14Val0Foundation) Point14ValAFoundation {
	dependency := point14ValADependencySnapshotFromUpstream(val0)
	return Point14ValAFoundation{
		CurrentState:                 Point14ValAStateActive,
		ProjectionDisclaimer:         point14ValAProjectionDisclaimerBase,
		DependencyState:              Point14ValAStateActive,
		NormalizedSignalState:        Point14ValAStateActive,
		SourceIdentityState:          Point14ValAStateActive,
		ScopeBindingState:            Point14ValAStateActive,
		EvidenceBindingState:         Point14ValAStateActive,
		FreshnessTimestampState:      Point14ValAStateActive,
		DuplicateRelationGuardState:  Point14ValAStateActive,
		TenantBoundaryGuardState:     Point14ValAStateActive,
		ValidationResultState:        Point14ValAStateActive,
		NoExternalAuthorityState:     Point14ValAStateActive,
		NoOverclaimState:             Point14ValAStateActive,
		Dependency:                   dependency,
		NormalizedExternalSignal:     point14ValANormalizedExternalSignalModel(dependency),
		SourceIdentity:               point14ValASourceIdentityModel(),
		ScopeBinding:                 point14ValAScopeBindingModel(dependency),
		EvidenceBinding:              point14ValAEvidenceBindingModel(dependency),
		FreshnessAndTimestamp:        point14ValAFreshnessAndTimestampBoundaryModel(dependency),
		DuplicateAndRelationGuard:    point14ValADuplicateAndRelationGuardModel(dependency),
		TenantBoundaryGuard:          point14ValATenantBoundaryGuardModel(dependency),
		ValidationResult:             point14ValAValidationResultModel(),
		NoExternalAuthorityGuard:     point14ValANoExternalAuthorityValidationGuardModel(),
		NoOverclaimValidationWording: point14ValANoOverclaimValidationWordingModel(),
	}
}

func Point14ValAFoundationModel() Point14ValAFoundation {
	val0 := ComputePoint14Val0Foundation(Point14Val0FoundationModel())
	return point14ValAFoundationModelFromUpstream(val0)
}

func ComputePoint14ValAFoundation(model Point14ValAFoundation) Point14ValAFoundation {
	model.DependencyState = EvaluatePoint14ValADependencyState(model.Dependency)
	model.NormalizedSignalState = EvaluatePoint14ValANormalizedExternalSignalState(model.NormalizedExternalSignal, model.Dependency)
	model.SourceIdentityState = EvaluatePoint14ValAExternalSignalSourceIdentityState(model.SourceIdentity)
	model.ScopeBindingState = EvaluatePoint14ValAExternalSignalScopeBindingState(model.ScopeBinding, model.Dependency)
	model.EvidenceBindingState = EvaluatePoint14ValAExternalSignalEvidenceBindingState(model.EvidenceBinding)
	model.FreshnessTimestampState = EvaluatePoint14ValAExternalSignalFreshnessAndTimestampBoundaryState(model.FreshnessAndTimestamp)
	model.DuplicateRelationGuardState = EvaluatePoint14ValAExternalSignalDuplicateAndRelationGuardState(model.DuplicateAndRelationGuard)
	model.TenantBoundaryGuardState = EvaluatePoint14ValAExternalSignalTenantBoundaryGuardState(model.TenantBoundaryGuard)
	model.NoExternalAuthorityState = EvaluatePoint14ValANoExternalAuthorityValidationGuardState(model.NoExternalAuthorityGuard)
	model.NoOverclaimState = EvaluatePoint14ValANoOverclaimValidationWordingState(model.NoOverclaimValidationWording)

	model.ValidationResult.NormalizedSignalState = model.NormalizedSignalState
	model.ValidationResult.SourceIdentityState = model.SourceIdentityState
	model.ValidationResult.ScopeBindingState = model.ScopeBindingState
	model.ValidationResult.EvidenceBindingState = model.EvidenceBindingState
	model.ValidationResult.FreshnessTimestampState = model.FreshnessTimestampState
	model.ValidationResult.DuplicateRelationGuardState = model.DuplicateRelationGuardState
	model.ValidationResult.TenantBoundaryGuardState = model.TenantBoundaryGuardState
	model.ValidationResult.NoExternalAuthorityState = model.NoExternalAuthorityState
	model.ValidationResult.NoOverclaimState = model.NoOverclaimState
	model.ValidationResult.ValidationState = point14ValAResolvedValidationState(
		model.ValidationResult.ValidationState,
		point14ValAFoundationState(
			model.NormalizedSignalState,
			model.SourceIdentityState,
			model.ScopeBindingState,
			model.EvidenceBindingState,
			model.FreshnessTimestampState,
			model.DuplicateRelationGuardState,
			model.TenantBoundaryGuardState,
			model.NoExternalAuthorityState,
			model.NoOverclaimState,
		),
	)
	model.ValidationResult.CandidateUsable = strings.TrimSpace(model.ValidationResult.ValidationState) == point14ValAValidationCandidateValidated
	model.ValidationResultState = EvaluatePoint14ValAExternalSignalValidationResultState(model.ValidationResult)

	model.CurrentState = point14ValAFoundationState(
		model.DependencyState,
		model.NormalizedSignalState,
		model.SourceIdentityState,
		model.ScopeBindingState,
		model.EvidenceBindingState,
		model.FreshnessTimestampState,
		model.DuplicateRelationGuardState,
		model.TenantBoundaryGuardState,
		model.ValidationResultState,
		model.NoExternalAuthorityState,
		model.NoOverclaimState,
	)
	model.BlockingReasons = point14ValABlockingReasons(model)
	model.ReviewPrerequisites = point14ValAReviewPrerequisites(model)
	return model
}

package formal

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/operability"
)

const (
	Point14Val0StateActive         = "point14_val0_ecosystem_authority_foundation_active"
	Point14Val0StateBlocked        = "point14_val0_ecosystem_authority_foundation_blocked"
	Point14Val0StateReviewRequired = "point14_val0_ecosystem_authority_foundation_review_required"
	Point14Val0StateIncomplete     = "point14_val0_ecosystem_authority_foundation_incomplete"
)

const (
	point14Val0PointID                  = "point_14"
	point14Val0WaveID                   = "val_0"
	point14Val0ProjectionDisclaimerBase = "projection_only not_canonical_truth point14_val0_external_ecosystem_authority_foundation"
	point14Val0BlockedPassToken         = "point_14_pass"

	point14Val0ScopeTenantScoped         = "tenant_scoped"
	point14Val0ScopeGlobalAdvisory       = "global_advisory"
	point14Val0ScopePublicNonAuthorative = "public_non_authoritative"

	point14Val0VisibilityTenantPrivate        = "tenant_private"
	point14Val0VisibilityScopedCustomer       = "scoped_customer_visible"
	point14Val0VisibilityPublicNoticeLimited  = "public_notice_limited"
	point14Val0TimeSourceServerUTC            = "server_utc"
	point14Val0TimeSourceApprovedCustomerTime = "approved_customer_time_source"
	point14Val0TimeSourceClientLocal          = "client_local"

	point14Val0ValidationReceived          = "received"
	point14Val0ValidationProvenancePending = "provenance_pending"
	point14Val0ValidationValidated         = "validated_candidate"
	point14Val0ValidationConflicting       = "conflicting"
	point14Val0ValidationSuperseded        = "superseded"
	point14Val0ValidationRevoked           = "revoked"
	point14Val0ValidationRejected          = "rejected"

	point14Val0DisputeOpened         = "opened"
	point14Val0DisputeTriaged        = "triaged"
	point14Val0DisputeEvidenceNeeded = "evidence_required"
	point14Val0DisputeReviewRequired = "review_required"
	point14Val0DisputeCorrected      = "corrected"
	point14Val0DisputeSuperseded     = "superseded"
	point14Val0DisputeRejected       = "rejected"
	point14Val0DisputeRevoked        = "revoked"
	point14Val0DisputePublished      = "published_notice"
)

type Point14Val0DependencySnapshot struct {
	Point13ValECurrentState             string                                          `json:"point13_vale_current_state"`
	Point13ValEDependencyState          string                                          `json:"point13_vale_dependency_state"`
	Point13ValENoOverclaimState         string                                          `json:"point13_vale_no_overclaim_state"`
	Point13ValEAuthorityBoundaryState   string                                          `json:"point13_vale_authority_boundary_state"`
	Point13ValETimestampIntegrityState  string                                          `json:"point13_vale_timestamp_integrity_state"`
	Point13ValETwitterIsolationState    string                                          `json:"point13_vale_tenant_isolation_state"`
	Point13ValEEvidenceIntegrityState   string                                          `json:"point13_vale_evidence_integrity_state"`
	Point13ValEPassClosureManifestState string                                          `json:"point13_vale_pass_closure_manifest_state"`
	Point13ValEPassAllowed              bool                                            `json:"point13_vale_pass_allowed"`
	Point13ValEPassToken                string                                          `json:"point13_vale_pass_token"`
	Point13ValEPointID                  string                                          `json:"point13_vale_point_id"`
	Point13ValEWaveID                   string                                          `json:"point13_vale_wave_id"`
	Point13ValEComputedFromUpstream     bool                                            `json:"point13_vale_computed_from_upstream"`
	Point13ValEMerged                   bool                                            `json:"point13_vale_merged"`
	Point13ValECIGreen                  bool                                            `json:"point13_vale_ci_green"`
	Point13ValEReviewedOnMain           bool                                            `json:"point13_vale_reviewed_on_main"`
	Point14PassSeen                     bool                                            `json:"point14_pass_seen"`
	InheritedPoint12CurrentState        string                                          `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState     string                                          `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureState    string                                          `json:"inherited_point12_pass_closure_state"`
	InheritedPoint12ReviewerResult      string                                          `json:"inherited_point12_reviewer_result"`
	InheritedPoint11CurrentState        string                                          `json:"inherited_point11_current_state"`
	InheritedPoint11PublicationState    string                                          `json:"inherited_point11_publication_state"`
	InheritedPoint11NoOverclaimState    string                                          `json:"inherited_point11_no_overclaim_state"`
	InheritedPoint11FinalPassGateState  string                                          `json:"inherited_point11_final_pass_gate_state"`
	InheritedPoint10CurrentState        string                                          `json:"inherited_point10_current_state"`
	InheritedPoint10NoOverclaimState    string                                          `json:"inherited_point10_no_overclaim_state"`
	InheritedPoint10ProjectionState     string                                          `json:"inherited_point10_projection_state"`
	InheritedPoint10PassRuleState       string                                          `json:"inherited_point10_pass_rule_state"`
	InheritedTenantScope                string                                          `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput          bool                                            `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                 []string                                        `json:"review_prerequisites,omitempty"`
	Point13ValE                         Point13ValEFoundation                           `json:"point13_vale"`
	Point12                             Point12ValEFoundation                           `json:"point12"`
	Point11                             Point11ValDFoundation                           `json:"point11"`
	Point10                             operability.DeploymentMultiTenantValEFoundation `json:"point10"`
}

type ExternalEcosystemSignalCandidate struct {
	SignalID                  string   `json:"signal_id"`
	SourceType                string   `json:"source_type"`
	SourceRef                 string   `json:"source_ref"`
	SignalType                string   `json:"signal_type"`
	ScopeClassification       string   `json:"scope_classification"`
	TenantScope               string   `json:"tenant_scope"`
	ReferencedTenantScope     string   `json:"referenced_tenant_scope"`
	ArtifactRef               string   `json:"artifact_ref"`
	ArtifactBindingRequired   bool     `json:"artifact_binding_required"`
	ArtifactBindingConsistent bool     `json:"artifact_binding_consistent"`
	ClaimRefs                 []string `json:"claim_refs,omitempty"`
	ClaimBindingRequired      bool     `json:"claim_binding_required"`
	ClaimBindingConsistent    bool     `json:"claim_binding_consistent"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	SourceEventAt             string   `json:"source_event_at"`
	SourceEventTimeSource     string   `json:"source_event_time_source"`
	ReceivedAt                string   `json:"received_at"`
	ReceivedTimeSource        string   `json:"received_time_source"`
	ValidatedAt               string   `json:"validated_at"`
	ValidationStatus          string   `json:"validation_status"`
	CustodyRef                string   `json:"custody_ref"`
	HashRef                   string   `json:"hash_ref"`
	SignatureRef              string   `json:"signature_ref"`
	ProvenanceRef             string   `json:"provenance_ref"`
	RequireHash               bool     `json:"require_hash"`
	RequireSignature          bool     `json:"require_signature"`
	RequireProvenance         bool     `json:"require_provenance"`
	SignalIdentityKey         string   `json:"signal_identity_key"`
	DuplicateSignalRefs       []string `json:"duplicate_signal_refs,omitempty"`
	RevocationRef             string   `json:"revocation_ref"`
	SupersededByRef           string   `json:"superseded_by_ref"`
	CanonicalAuthority        bool     `json:"canonical_authority"`
	PassAllowed               bool     `json:"pass_allowed"`
	CanonicalMutationAllowed  bool     `json:"canonical_mutation_allowed"`
	ProductionMutationAllowed bool     `json:"production_mutation_allowed"`
	OverrideCanonicalDecision bool     `json:"override_canonical_decision"`
}

type ExternalStakeholderAuthorityRole struct {
	RoleID                       string   `json:"role_id"`
	RoleType                     string   `json:"role_type"`
	StakeholderRef               string   `json:"stakeholder_ref"`
	ScopeClassification          string   `json:"scope_classification"`
	TenantScope                  string   `json:"tenant_scope"`
	AllowedActionRefs            []string `json:"allowed_action_refs,omitempty"`
	GovernanceRequiredActionRefs []string `json:"governance_required_action_refs,omitempty"`
}

type ExternalAuthorityConflictMatrix struct {
	ConflictPresent                 bool     `json:"conflict_present"`
	ConflictID                      string   `json:"conflict_id"`
	ConflictType                    string   `json:"conflict_type"`
	SignalRefs                      []string `json:"signal_refs,omitempty"`
	RoleRefs                        []string `json:"role_refs,omitempty"`
	AffectedArtifactRefs            []string `json:"affected_artifact_refs,omitempty"`
	AffectedClaimRefs               []string `json:"affected_claim_refs,omitempty"`
	AffectedEvidenceRefs            []string `json:"affected_evidence_refs,omitempty"`
	GovernancePath                  string   `json:"governance_path"`
	AuditEventRef                   string   `json:"audit_event_ref"`
	TenantScope                     string   `json:"tenant_scope"`
	AffectedTenantScope             string   `json:"affected_tenant_scope"`
	ResolvedByGovernance            bool     `json:"resolved_by_governance"`
	GovernanceResolutionEventRef    string   `json:"governance_resolution_event_ref"`
	ConsensusResolutionRequested    bool     `json:"consensus_resolution_requested"`
	SilentActiveResolutionRequested bool     `json:"silent_active_resolution_requested"`
}

type ExternalSignalDisputeLifecycle struct {
	DisputeID                  string   `json:"dispute_id"`
	DisputeState               string   `json:"dispute_state"`
	TenantScope                string   `json:"tenant_scope"`
	AffectedTenantScope        string   `json:"affected_tenant_scope"`
	SignalRefs                 []string `json:"signal_refs,omitempty"`
	EvidenceRefs               []string `json:"evidence_refs,omitempty"`
	AuditEventRefs             []string `json:"audit_event_refs,omitempty"`
	OpenedAt                   string   `json:"opened_at"`
	CorrectedAt                string   `json:"corrected_at"`
	RevokedAt                  string   `json:"revoked_at"`
	PublishedAt                string   `json:"published_at"`
	TimestampSource            string   `json:"timestamp_source"`
	GovernanceEventRef         string   `json:"governance_event_ref"`
	VisibilityBoundaryRef      string   `json:"visibility_boundary_ref"`
	SupersessionRef            string   `json:"supersession_ref"`
	CanonicalMutationRequested bool     `json:"canonical_mutation_requested"`
	RejectedDeletesEvidence    bool     `json:"rejected_deletes_evidence"`
	PrivacyCheckPassed         bool     `json:"privacy_check_passed"`
}

type ExternalCorrectionRevocationBoundary struct {
	CorrectionRef                  string   `json:"correction_ref"`
	RevocationRef                  string   `json:"revocation_ref"`
	RevocationModeled              bool     `json:"revocation_modeled"`
	PublicPrivateVisibility        string   `json:"public_private_visibility"`
	TenantScope                    string   `json:"tenant_scope"`
	PrivacyBoundaryRef             string   `json:"privacy_boundary_ref"`
	ApproverRef                    string   `json:"approver_ref"`
	GovernanceEventRef             string   `json:"governance_event_ref"`
	AuditEventRef                  string   `json:"audit_event_ref"`
	LimitationRefs                 []string `json:"limitation_refs,omitempty"`
	ObservedTexts                  []string `json:"observed_texts,omitempty"`
	OpenedAt                       string   `json:"opened_at"`
	CorrectedAt                    string   `json:"corrected_at"`
	RevokedAt                      string   `json:"revoked_at"`
	PublishedAt                    string   `json:"published_at"`
	TimestampSource                string   `json:"timestamp_source"`
	LeaksTenantPrivateData         bool     `json:"leaks_tenant_private_data"`
	StrengthensClaim               bool     `json:"strengthens_claim"`
	CreatesForbiddenAuthorityWords bool     `json:"creates_forbidden_authority_words"`
	OmitsMeaningChangingLimitation bool     `json:"omits_meaning_changing_limitation"`
	HidesDecisiveMissingEvidence   bool     `json:"hides_decisive_missing_evidence"`
	ClientTimeBackdated            bool     `json:"client_time_backdated"`
}

type ExternalVisibilityPublicationBoundary struct {
	VisibilityID              string   `json:"visibility_id"`
	VisibilityScope           string   `json:"visibility_scope"`
	TenantScope               string   `json:"tenant_scope"`
	PublicOutputRefs          []string `json:"public_output_refs,omitempty"`
	PrivateScopeRefs          []string `json:"private_scope_refs,omitempty"`
	LimitationTexts           []string `json:"limitation_texts,omitempty"`
	ObservedTexts             []string `json:"observed_texts,omitempty"`
	AuditEventRef             string   `json:"audit_event_ref"`
	ContainsPrivateTenantData bool     `json:"contains_private_tenant_data"`
	OmitsLimitations          bool     `json:"omits_limitations"`
}

type AgentEcosystemInputBoundary struct {
	BoundaryID               string   `json:"boundary_id"`
	TenantScope              string   `json:"tenant_scope"`
	AgentInputRefs           []string `json:"agent_input_refs,omitempty"`
	EvidenceRefs             []string `json:"evidence_refs,omitempty"`
	AuditEventRef            string   `json:"audit_event_ref"`
	AdvisoryOnly             bool     `json:"advisory_only"`
	CanDecideDispute         bool     `json:"can_decide_dispute"`
	CanPublishCorrection     bool     `json:"can_publish_correction"`
	CanRevokeClaim           bool     `json:"can_revoke_claim"`
	CanOverrideConflict      bool     `json:"can_override_conflict"`
	CanEmitPass              bool     `json:"can_emit_pass"`
	CanEmitPublicAuthority   bool     `json:"can_emit_public_authority"`
	PassAllowed              bool     `json:"pass_allowed"`
	ApprovalGranted          bool     `json:"approval_granted"`
	ProductionApproved       bool     `json:"production_approved"`
	DeploymentApproved       bool     `json:"deployment_approved"`
	ExternalAuthorityAllowed bool     `json:"external_authority_allowed"`
}

type Point14Val0NoExternalAuthorityGuard struct {
	ObservedAuthorityMarkers []string `json:"observed_authority_markers,omitempty"`
}

type Point14Val0NoOverclaimEcosystemWording struct {
	ObservedSignalTexts                  []string `json:"observed_signal_texts,omitempty"`
	ObservedDisputeTexts                 []string `json:"observed_dispute_texts,omitempty"`
	ObservedCorrectionTexts              []string `json:"observed_correction_texts,omitempty"`
	ObservedPublicationTexts             []string `json:"observed_publication_texts,omitempty"`
	ObservedAgentTexts                   []string `json:"observed_agent_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point14Val0Foundation struct {
	CurrentState                          string                                 `json:"current_state"`
	BlockingReasons                       []string                               `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites                   []string                               `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer                  string                                 `json:"projection_disclaimer"`
	DependencyState                       string                                 `json:"dependency_state"`
	ExternalSignalCandidateState          string                                 `json:"external_signal_candidate_state"`
	ExternalStakeholderAuthorityRoleState string                                 `json:"external_stakeholder_authority_role_state"`
	ExternalAuthorityConflictMatrixState  string                                 `json:"external_authority_conflict_matrix_state"`
	ExternalSignalDisputeLifecycleState   string                                 `json:"external_signal_dispute_lifecycle_state"`
	ExternalCorrectionRevocationState     string                                 `json:"external_correction_revocation_state"`
	ExternalVisibilityPublicationState    string                                 `json:"external_visibility_publication_state"`
	AgentEcosystemInputBoundaryState      string                                 `json:"agent_ecosystem_input_boundary_state"`
	NoExternalAuthorityGuardState         string                                 `json:"no_external_authority_guard_state"`
	NoOverclaimState                      string                                 `json:"no_overclaim_state"`
	Dependency                            Point14Val0DependencySnapshot          `json:"dependency"`
	ExternalSignalCandidate               ExternalEcosystemSignalCandidate       `json:"external_signal_candidate"`
	ExternalStakeholderAuthorityRole      ExternalStakeholderAuthorityRole       `json:"external_stakeholder_authority_role"`
	ExternalAuthorityConflictMatrix       ExternalAuthorityConflictMatrix        `json:"external_authority_conflict_matrix"`
	ExternalSignalDisputeLifecycle        ExternalSignalDisputeLifecycle         `json:"external_signal_dispute_lifecycle"`
	ExternalCorrectionRevocationBoundary  ExternalCorrectionRevocationBoundary   `json:"external_correction_revocation_boundary"`
	ExternalVisibilityPublicationBoundary ExternalVisibilityPublicationBoundary  `json:"external_visibility_publication_boundary"`
	AgentEcosystemInputBoundary           AgentEcosystemInputBoundary            `json:"agent_ecosystem_input_boundary"`
	NoExternalAuthorityGuard              Point14Val0NoExternalAuthorityGuard    `json:"no_external_authority_guard"`
	NoOverclaimEcosystemWording           Point14Val0NoOverclaimEcosystemWording `json:"no_overclaim_ecosystem_wording"`
}

func point14Val0States() []string {
	return []string{
		Point14Val0StateActive,
		Point14Val0StateBlocked,
		Point14Val0StateReviewRequired,
		Point14Val0StateIncomplete,
	}
}

func point14Val0RawValueInSet(value string, allowed []string) bool {
	for _, candidate := range allowed {
		if value == candidate {
			return true
		}
	}
	return false
}

func point14Val0StateValid(value string) bool {
	return point14Val0RawValueInSet(value, point14Val0States())
}

func point14Val0SourceTypes() []string {
	return []string{
		"vendor_advisory",
		"scanner_finding",
		"vex_statement",
		"maintainer_notice",
		"auditor_note",
		"verifier_report",
		"partner_signal",
		"customer_admin_signal",
		"security_reviewer_signal",
		"external_research_report",
		"public_notice",
		"crowd_signal",
		"agent_recommendation_source",
	}
}

func point14Val0SignalTypes() []string {
	return []string{
		"vulnerability_assertion",
		"not_affected_assertion",
		"fixed_assertion",
		"disputed_assertion",
		"exploit_notice",
		"correction_notice",
		"revocation_notice",
		"suppression_request",
		"publication_notice",
		"scope_disagreement",
		"evidence_submission",
	}
}

func point14Val0ScopeClassifications() []string {
	return []string{
		point14Val0ScopeTenantScoped,
		point14Val0ScopeGlobalAdvisory,
		point14Val0ScopePublicNonAuthorative,
	}
}

func point14Val0ValidationStatuses() []string {
	return []string{
		point14Val0ValidationReceived,
		point14Val0ValidationProvenancePending,
		point14Val0ValidationValidated,
		point14Val0ValidationConflicting,
		point14Val0ValidationSuperseded,
		point14Val0ValidationRevoked,
		point14Val0ValidationRejected,
	}
}

func point14Val0RoleTypes() []string {
	return []string{
		"developer",
		"maintainer",
		"scanner",
		"vex_issuer",
		"verifier",
		"auditor",
		"partner",
		"customer_admin",
		"security_reviewer",
		"external_researcher",
		"agent_recommendation_source",
	}
}

func point14Val0ConflictTypes() []string {
	return []string{
		"vulnerable_vs_not_affected",
		"fixed_vs_not_fixed",
		"scanner_vs_maintainer",
		"vex_vs_scanner",
		"auditor_vs_canonical",
		"verifier_vs_canonical",
		"public_consensus_vs_canonical",
		"partner_vs_customer_admin",
		"duplicate_signal_disagreement",
		"provenance_conflict",
		"tenant_scope_conflict",
	}
}

func point14Val0GovernancePaths() []string {
	return []string{
		"canonical_review",
		"contradiction_review",
		"provenance_review",
		"tenant_scope_review",
		"revocation_review",
		"publication_review",
		"correction_review",
	}
}

func point14Val0VisibilityScopes() []string {
	return []string{
		point14Val0VisibilityTenantPrivate,
		point14Val0VisibilityScopedCustomer,
		point14Val0VisibilityPublicNoticeLimited,
	}
}

func point14Val0TimeSources() []string {
	return []string{
		point14Val0TimeSourceServerUTC,
		point14Val0TimeSourceApprovedCustomerTime,
		point14Val0TimeSourceClientLocal,
	}
}

func point14Val0BoundedRoleActions() []string {
	return []string{
		"submit_evidence",
		"submit_annotation",
		"submit_dispute",
		"submit_correction_request",
		"submit_revocation_request",
		"submit_publication_request",
		"submit_scope_feedback",
	}
}

func point14Val0GovernanceRequiredActions() []string {
	return []string{
		"resolve_conflict",
		"publish_correction_notice",
		"execute_revocation",
		"approve_public_notice",
		"approve_authoritative_correction",
	}
}

func point14Val0ForbiddenRoleActions() []string {
	return []string{
		"approve_production",
		"authorize_deployment",
		"emit_pass",
		"certify",
		"publish_public_badge",
		"override_canonical_evidence",
		"revoke_canonical_claim",
		"publish_authoritative_correction",
	}
}

func point14Val0ForbiddenAuthorityMarkers() []string {
	return []string{
		"external_pass",
		"public_pass",
		"crowd_pass",
		"vendor_certified",
		"scanner_certified",
		"auditor_certified",
		"verifier_approved",
		"partner_approved",
		"customer_admin_production_approved",
		"global_oss_authority",
		"public_badge_authority",
		"external_source_of_truth",
		"automatic_correction_published",
	}
}

func point14Val0ForbiddenWording() []string {
	return []string{
		"certified",
		"guaranteed secure",
		"regulator-approved",
		"compliance guaranteed",
		"production approved",
		"deployment approved",
		"public badge",
		"global truth",
		"official authority",
		"universal trust score",
		"crowd-approved",
		"vendor-certified",
		"auditor-certified",
		"external PASS",
		"scanner PASS",
		"automatic correction",
		"compliance proof",
		"legal proof",
		"financial guarantee",
	}
}

func point14Val0SafeWording() []string {
	return []string{
		"external evidence signal",
		"advisory external signal",
		"evidence input pending validation",
		"dispute requires governance review",
		"correction notice pending approval",
		"bounded ecosystem evidence input",
		"no external authority granted",
	}
}

func point14Val0AuthorityMarkerForbidden(marker string, forbidden []string) bool {
	normalizedMarker := point16Val0NormalizeObservedText(marker)
	if normalizedMarker == "" {
		return false
	}
	for _, forbiddenMarker := range forbidden {
		normalizedForbidden := point16Val0NormalizeObservedText(forbiddenMarker)
		if normalizedForbidden == "" {
			continue
		}
		if point16Val0NormalizedTextContainsNormalizedPhrase(normalizedMarker, normalizedForbidden) {
			return true
		}
		if point16Val0ObservedTextContainsNormalizedPhrase(marker, normalizedForbidden) {
			return true
		}
	}
	return false
}

func point14Val0AuthorityMarkersContainForbidden(markers []string, forbidden []string) bool {
	return point14Val0ListContainsForbiddenWordingFor(markers, nil, forbidden)
}

func point14Val0RefValid(value string, prefixes ...string) bool {
	if value == "" ||
		value != strings.TrimSpace(value) ||
		strings.ContainsAny(value, "\t\r\n") ||
		strings.Contains(value, " ") ||
		strings.Contains(value, "/") ||
		!point11Val0ASCIIVisibleValue(value) {
		return false
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(value, prefix) {
			return true
		}
	}
	return false
}

func point14Val0SignalIDValid(value string) bool {
	return point14Val0RefValid(value, "signal_")
}

func point14Val0SourceRefValid(value string) bool {
	return point14Val0RefValid(value, "source_")
}

func point14Val0ArtifactRefValid(value string) bool {
	return point14Val0RefValid(value, "artifact_")
}

func point14Val0ClaimRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "claim_")
}

func point14Val0EvidenceRefsValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	seen := map[string]struct{}{}
	for _, value := range values {
		if !point14Val0EvidenceRefValid(value) {
			return false
		}
		if _, exists := seen[value]; exists {
			return false
		}
		seen[value] = struct{}{}
	}
	return true
}

func point14Val0EvidenceRefValid(value string) bool {
	if !point11Val0IdentityValueValid(value) || strings.Contains(value, " ") || strings.Contains(value, "/") || !strings.HasPrefix(value, "evidence_") {
		return false
	}
	return !point11Val0ContainsTenantBoundaryBypass(value)
}

func point14Val0Point11FoundationActive(model Point11ValDFoundation) bool {
	computed := ComputePoint11ValDFoundation(model)
	if computed.CurrentState != Point11ValDStateActive ||
		computed.DependencyState != Point11ValDDependencyStateActive ||
		computed.IntegratedInvariantState != Point11ValDIntegratedInvariantStateActive ||
		computed.QualityMapState != Point11ValDQualityMapStateActive ||
		computed.PublicationReviewState != Point11ValDPublicationReviewStateActive ||
		computed.NoOverclaimReviewState != Point11ValDNoOverclaimReviewStateActive ||
		computed.CleanRoomIPReviewState != Point11ValDCleanRoomIPReviewStateActive ||
		computed.CLBClosureState != Point11ValDCLBClosureStateActive ||
		computed.PassClosureManifestState != Point11ValDPassClosureManifestStateActive ||
		computed.FinalPassGateState != Point11ValDFinalPassGateStateActive ||
		computed.Point11PassToken != point11ValDPoint11PassToken ||
		computed.PassClosureManifest.CurrentState != Point11ValDPassClosureManifestStateActive ||
		!computed.PassClosureManifest.Point11PassAllowed ||
		computed.PassClosureManifest.Point11PassToken != point11ValDPoint11PassToken ||
		computed.FinalPassGate.CurrentState != Point11ValDFinalPassGateStateActive ||
		!computed.FinalPassGate.Point11PassAllowed ||
		!computed.FinalPassGate.Point11PassEmitted ||
		computed.FinalPassGate.Point11PassToken != point11ValDPoint11PassToken {
		return false
	}
	return model.CurrentState == computed.CurrentState &&
		model.DependencyState == computed.DependencyState &&
		model.IntegratedInvariantState == computed.IntegratedInvariantState &&
		model.QualityMapState == computed.QualityMapState &&
		model.PublicationReviewState == computed.PublicationReviewState &&
		model.NoOverclaimReviewState == computed.NoOverclaimReviewState &&
		model.CleanRoomIPReviewState == computed.CleanRoomIPReviewState &&
		model.CLBClosureState == computed.CLBClosureState &&
		model.PassClosureManifestState == computed.PassClosureManifestState &&
		model.FinalPassGateState == computed.FinalPassGateState &&
		model.Point11PassToken == computed.Point11PassToken &&
		reflect.DeepEqual(model.PassClosureManifest, computed.PassClosureManifest) &&
		reflect.DeepEqual(model.FinalPassGate, computed.FinalPassGate)
}

func point14Val0FoundationComputedActive(model Point14Val0Foundation) bool {
	computed := ComputePoint14Val0Foundation(model)
	return computed.CurrentState == Point14Val0StateActive &&
		model.CurrentState == Point14Val0StateActive
}

func point14ValAFoundationComputedActive(model Point14ValAFoundation) bool {
	computed := ComputePoint14ValAFoundation(model)
	return computed.CurrentState == Point14ValAStateActive &&
		model.CurrentState == Point14ValAStateActive
}

func point14ValBFoundationComputedActive(model Point14ValBFoundation) bool {
	computed := ComputePoint14ValBFoundation(model)
	return computed.CurrentState == Point14ValBStateActive &&
		model.CurrentState == Point14ValBStateActive
}

func point14ValCFoundationComputedActive(model Point14ValCFoundation) bool {
	computed := ComputePoint14ValCFoundation(model)
	return computed.CurrentState == Point14ValCStateActive &&
		model.CurrentState == Point14ValCStateActive
}

func point14ValDFoundationComputedActive(model Point14ValDFoundation) bool {
	computed := ComputePoint14ValDFoundation(model)
	return computed.CurrentState == Point14ValDStateActive &&
		model.CurrentState == Point14ValDStateActive
}

func point14ValEFoundationComputedPassConfirmed(model Point14ValEFoundation) bool {
	computed := ComputePoint14ValEFoundation(model)
	return computed.CurrentState == Point14ValEStatePassConfirmed &&
		model.CurrentState == computed.CurrentState &&
		model.DependencyState == computed.DependencyState &&
		model.ExternalSignalValidationClosureState == computed.ExternalSignalValidationClosureState &&
		model.ConflictDisputeClosureState == computed.ConflictDisputeClosureState &&
		model.CorrectionPublicationClosureState == computed.CorrectionPublicationClosureState &&
		model.TimelineProjectionClosureState == computed.TimelineProjectionClosureState &&
		model.AuthorityBoundaryClosureState == computed.AuthorityBoundaryClosureState &&
		model.TenantPrivacyClosureState == computed.TenantPrivacyClosureState &&
		model.TimestampIntegrityClosureState == computed.TimestampIntegrityClosureState &&
		model.AgentAdvisoryClosureState == computed.AgentAdvisoryClosureState &&
		model.NoOverclaimFinalCheckState == computed.NoOverclaimFinalCheckState &&
		model.CLBFinalCheckState == computed.CLBFinalCheckState &&
		model.ClosureEvaluatorState == computed.ClosureEvaluatorState &&
		model.ClosureEvaluator.CurrentState == computed.ClosureEvaluator.CurrentState &&
		reflect.DeepEqual(model.ClosureEvaluator, computed.ClosureEvaluator) &&
		computed.PassClosureManifestState == Point14ValEStatePassConfirmed &&
		model.PassClosureManifestState == computed.PassClosureManifestState &&
		model.PassClosureManifest.CurrentState == computed.PassClosureManifest.CurrentState &&
		computed.Point14PassAllowed &&
		model.Point14PassAllowed == computed.Point14PassAllowed &&
		reflect.DeepEqual(model.PassClosureManifest, computed.PassClosureManifest) &&
		model.PassClosureManifest.Point14PassAllowed == computed.PassClosureManifest.Point14PassAllowed &&
		model.PassClosureManifest.Point14PassToken == computed.PassClosureManifest.Point14PassToken &&
		computed.Point14PassToken == point14Val0BlockedPassToken &&
		model.Point14PassToken == computed.Point14PassToken
}

func point14ValBDependencyChainComputedActive(model Point14ValBFoundation) bool {
	return point14ValBFoundationComputedActive(model) &&
		point14ValAFoundationComputedActive(model.Dependency.Point14ValA) &&
		point14Val0FoundationComputedActive(model.Dependency.Point14Val0)
}

func point14ValCDependencyChainComputedActive(model Point14ValCFoundation) bool {
	return point14ValCFoundationComputedActive(model) &&
		point14ValBDependencyChainComputedActive(model.Dependency.Point14ValB) &&
		point14ValAFoundationComputedActive(model.Dependency.Point14ValA) &&
		point14Val0FoundationComputedActive(model.Dependency.Point14Val0)
}

func point14ValDDependencyChainComputedActive(model Point14ValDFoundation) bool {
	return point14ValDFoundationComputedActive(model) &&
		point14ValCDependencyChainComputedActive(model.Dependency.Point14ValC) &&
		point14ValBDependencyChainComputedActive(model.Dependency.Point14ValB) &&
		point14ValAFoundationComputedActive(model.Dependency.Point14ValA) &&
		point14Val0FoundationComputedActive(model.Dependency.Point14Val0)
}

func point14ValCDependencyEmbeddedSnapshotCopiesExact(model Point14ValCDependencySnapshot) bool {
	return reflect.DeepEqual(model.Point14ValB.Dependency.Point14ValA, model.Point14ValA) &&
		reflect.DeepEqual(model.Point14ValB.Dependency.Point14Val0, model.Point14Val0) &&
		reflect.DeepEqual(model.Point14ValB.Dependency.Point14ValA.Dependency.Point14Val0, model.Point14Val0)
}

func point14ValCEmbeddedSnapshotCopiesExact(model Point14ValCFoundation) bool {
	return point14ValCDependencyEmbeddedSnapshotCopiesExact(model.Dependency)
}

func point14ValDDependencyEmbeddedSnapshotCopiesExact(model Point14ValDDependencySnapshot) bool {
	return reflect.DeepEqual(model.Point14ValB, model.Point14ValC.Dependency.Point14ValB) &&
		reflect.DeepEqual(model.Point14ValA, model.Point14ValC.Dependency.Point14ValA) &&
		reflect.DeepEqual(model.Point14Val0, model.Point14ValC.Dependency.Point14Val0) &&
		point14ValCEmbeddedSnapshotCopiesExact(model.Point14ValC)
}

func point14ValDFoundationEmbeddedSnapshotCopiesExact(model Point14ValDFoundation) bool {
	return point14ValDDependencyEmbeddedSnapshotCopiesExact(model.Dependency)
}

func point14Val0RoleRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "role_")
}

func point14Val0SignalRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "signal_")
}

func point14Val0AuditEventRefValid(value string) bool {
	return point14Val0RefValid(value, "audit_event_")
}

func point14Val0AuditEventRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "audit_event_")
}

func point14Val0ProvenanceRefValid(value string) bool {
	return point14Val0RefValid(value, "provenance_")
}

func point14Val0HashRefValid(value string) bool {
	return point14Val0RefValid(value, "hash_")
}

func point14Val0SignatureRefValid(value string) bool {
	return point14Val0RefValid(value, "signature_")
}

func point14Val0CustodyRefValid(value string) bool {
	return point14Val0RefValid(value, "custody_")
}

func point14Val0StakeholderRefValid(value string) bool {
	return point14Val0RefValid(value, "stakeholder_")
}

func point14Val0RoleIDValid(value string) bool {
	return point14Val0RefValid(value, "role_")
}

func point14Val0ConflictIDValid(value string) bool {
	return point14Val0RefValid(value, "conflict_")
}

func point14Val0DisputeIDValid(value string) bool {
	return point14Val0RefValid(value, "dispute_")
}

func point14Val0CorrectionRefValid(value string) bool {
	return point14Val0RefValid(value, "correction_")
}

func point14Val0RevocationRefValid(value string) bool {
	return point14Val0RefValid(value, "revocation_")
}

func point14Val0BoundaryRefValid(value string) bool {
	return point14Val0RefValid(value, "boundary_", "privacy_boundary_", "visibility_boundary_")
}

func point14Val0ApproverRefValid(value string) bool {
	return point14Val0RefValid(value, "approver_", "owner_")
}

func point14Val0GovernanceEventRefValid(value string) bool {
	return point14Val0RefValid(value, "governance_event_")
}

func point14Val0VisibilityIDValid(value string) bool {
	return point14Val0RefValid(value, "visibility_")
}

func point14Val0OutputRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "publication_output_")
}

func point14Val0PrivateScopeRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "tenant_scope_ref_", "private_scope_")
}

func point14Val0AgentInputRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "agent_input_")
}

func point14Val0IdentityKeyValid(value string) bool {
	return point14Val0RefValid(value, "signal_identity_")
}

func point14Val0RefListValid(values []string, prefixes ...string) bool {
	if len(values) == 0 {
		return false
	}
	for _, value := range values {
		if !point14Val0RefValid(value, prefixes...) {
			return false
		}
	}
	return true
}

func point14Val0TextListValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return false
		}
	}
	return true
}

func point14Val0ExactValueValid(value string, allowed []string) bool {
	return point14Val0RawValueInSet(value, allowed)
}

func point14Val0TimeSourceValid(value string) bool {
	return point14Val0ExactValueValid(value, point14Val0TimeSources())
}

func point14Val0CanonicalTimeSourceValid(value string) bool {
	return value == point14Val0TimeSourceServerUTC || value == point14Val0TimeSourceApprovedCustomerTime
}

func point14Val0ParsedTime(value string) (time.Time, bool) {
	if value == "" {
		return time.Time{}, false
	}
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, false
	}
	return parsed, true
}

func point14Val0ContainsForbiddenWording(text string) bool {
	return point14Val0ContainsForbiddenWordingFor(text, point14Val0SafeWording(), point14Val0ForbiddenWording())
}

func point14Val0ContainsForbiddenWordingFor(text string, safeWording []string, forbiddenWording []string) bool {
	normalized := point16Val0NormalizeObservedText(text)
	if normalized == "" {
		return false
	}
	for _, safe := range safeWording {
		if normalized == point16Val0NormalizeObservedText(safe) {
			return false
		}
	}
	for _, phrase := range forbiddenWording {
		normalizedPhrase := point16Val0NormalizeObservedText(phrase)
		if normalizedPhrase == "" {
			continue
		}
		if point16Val0NormalizedTextContainsNormalizedPhrase(normalized, normalizedPhrase) {
			return true
		}
		if point16Val0ObservedTextContainsNormalizedPhrase(text, normalizedPhrase) {
			return true
		}
	}
	return false
}

func point14Val0ListContainsForbiddenWording(values []string) bool {
	return point14Val0ListContainsForbiddenWordingFor(values, point14Val0SafeWording(), point14Val0ForbiddenWording())
}

func point14Val0ListContainsForbiddenWordingFor(values []string, safeWording []string, forbiddenWording []string) bool {
	corpusVariants := make([][]string, 0, len(values))
	corpusAllowed := make([]bool, 0, len(values))
	for _, value := range values {
		variants := point16Val0NormalizedObservedTextVariants(value)
		if len(variants) == 0 {
			continue
		}
		normalized := variants[0]
		allowed := false
		for _, safe := range safeWording {
			if normalized == point16Val0NormalizeObservedText(safe) {
				allowed = true
				break
			}
		}
		if !allowed && point14Val0ContainsForbiddenWordingFor(value, safeWording, forbiddenWording) {
			return true
		}
		corpusVariants = append(corpusVariants, variants)
		corpusAllowed = append(corpusAllowed, allowed)
	}
	if len(corpusVariants) == 0 {
		return false
	}
	for _, phrase := range forbiddenWording {
		if formalNoOverclaimForbiddenPhraseAcrossValueVariants(corpusVariants, corpusAllowed, point16Val0NormalizeObservedText(phrase)) {
			return true
		}
	}
	return false
}

func point14Val0SignalIdentityKey(model ExternalEcosystemSignalCandidate) string {
	claims := append([]string{}, model.ClaimRefs...)
	duplicates := append([]string{}, model.DuplicateSignalRefs...)
	sort.Strings(claims)
	sort.Strings(duplicates)
	parts := []string{
		model.SourceType,
		model.SourceRef,
		model.SignalType,
		model.ScopeClassification,
		model.ReferencedTenantScope,
		model.ArtifactRef,
		strings.Join(claims, ","),
		model.HashRef,
		model.SignatureRef,
		model.ProvenanceRef,
		strings.Join(duplicates, ","),
	}
	sum := sha256.Sum256([]byte(strings.Join(parts, "|")))
	return "signal_identity_" + hex.EncodeToString(sum[:])
}

func point14Val0DependencySnapshotFromUpstream(
	valE Point13ValEFoundation,
	point12 Point12ValEFoundation,
	point11 Point11ValDFoundation,
	point10 operability.DeploymentMultiTenantValEFoundation,
) Point14Val0DependencySnapshot {
	return Point14Val0DependencySnapshot{
		Point13ValECurrentState:             valE.CurrentState,
		Point13ValEDependencyState:          valE.DependencyState,
		Point13ValENoOverclaimState:         valE.NoOverclaimFinalCheckState,
		Point13ValEAuthorityBoundaryState:   valE.AuthorityBoundaryCheckState,
		Point13ValETimestampIntegrityState:  valE.TimestampIntegrityCheckState,
		Point13ValETwitterIsolationState:    valE.TenantIsolationCheckState,
		Point13ValEEvidenceIntegrityState:   valE.EvidenceIntegrityCheckState,
		Point13ValEPassClosureManifestState: valE.PassClosureManifestState,
		Point13ValEPassAllowed:              valE.Point13PassAllowed,
		Point13ValEPassToken:                valE.Point13PassToken,
		Point13ValEPointID:                  point13Val0PointID,
		Point13ValEWaveID:                   point13ValEWaveID,
		Point13ValEComputedFromUpstream:     valE.Dependency.SnapshotFromComputedOutput,
		Point13ValEMerged:                   true,
		Point13ValECIGreen:                  true,
		Point13ValEReviewedOnMain:           true,
		Point14PassSeen:                     false,
		InheritedPoint12CurrentState:        point12.CurrentState,
		InheritedPoint12DependencyState:     point12.DependencyState,
		InheritedPoint12PassClosureState:    point12.PassClosureManifestState,
		InheritedPoint12ReviewerResult:      point12.PassClosureManifest.ReviewerResult,
		InheritedPoint11CurrentState:        point11.CurrentState,
		InheritedPoint11PublicationState:    point11.PublicationReviewState,
		InheritedPoint11NoOverclaimState:    point11.NoOverclaimReviewState,
		InheritedPoint11FinalPassGateState:  point11.FinalPassGateState,
		InheritedPoint10CurrentState:        point10.Point10State,
		InheritedPoint10NoOverclaimState:    point10.NoOverclaimState,
		InheritedPoint10ProjectionState:     point10.ProjectionBoundaryState,
		InheritedPoint10PassRuleState:       point10.Point10PassRuleState,
		InheritedTenantScope:                valE.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:          true,
		Point13ValE:                         valE,
		Point12:                             point12,
		Point11:                             point11,
		Point10:                             point10,
	}
}

func point14Val0ActivePoint11Val0Foundation() Point11Val0Foundation {
	model := Point11Val0FoundationModel()
	valE := operability.ComputeDeploymentMultiTenantValEFoundation(operability.DeploymentMultiTenantValEFoundationModel())
	model.Dependency = SnapshotPoint11Val0DependencyFromComputedPoint10ValE(valE, Point11Val0Point10RepoReview{
		LatestValEClosurePatchPresent: true,
		Point10PassOutsideValE:        false,
		CIGreenVisible:                true,
		CIGreen:                       true,
		MergeStatusVisible:            true,
		MergeAccepted:                 true,
	})
	return ComputePoint11Val0Foundation(model)
}

func point14Val0ActivePoint11ValAFoundation() Point11ValAFoundation {
	model := Point11ValAFoundationModel()
	model.Dependency = SnapshotPoint11ValADependencyFromComputedVal0(point14Val0ActivePoint11Val0Foundation(), Point11ValAVal0ReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
	})
	return ComputePoint11ValAFoundation(model)
}

func point14Val0ActivePoint11ValBFoundation() Point11ValBFoundation {
	model := Point11ValBFoundationModel()
	model.Dependency = SnapshotPoint11ValBDependencyFromComputedValA(point14Val0ActivePoint11ValAFoundation(), Point11ValBValAReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
	})
	return ComputePoint11ValBFoundation(model)
}

func point14Val0ActivePoint11ValCFoundation() Point11ValCFoundation {
	model := Point11ValCFoundationModel()
	model.Dependency = SnapshotPoint11ValCDependencyFromComputedValB(point14Val0ActivePoint11ValBFoundation(), Point11ValCValBReviewContext{
		LocalReviewAllowsDependencyReviewRequired: true,
	})
	return ComputePoint11ValCFoundation(model)
}

func point14Val0ActivePoint11ValDFoundation() Point11ValDFoundation {
	model := Point11ValDFoundationModel()
	model.Val0Dependency = SnapshotPoint11ValDVal0DependencyFromComputed(point14Val0ActivePoint11Val0Foundation(), Point11ValDVal0ReviewContext{})
	model.ValADependency = SnapshotPoint11ValDValADependencyFromComputed(point14Val0ActivePoint11ValAFoundation(), Point11ValDValAReviewContext{})
	model.ValBDependency = SnapshotPoint11ValDValBDependencyFromComputed(point14Val0ActivePoint11ValBFoundation(), Point11ValDValBReviewContext{})
	model.ValCDependency = SnapshotPoint11ValDValCDependencyFromComputed(point14Val0ActivePoint11ValCFoundation(), Point11ValDValCReviewContext{})
	return ComputePoint11ValDFoundation(model)
}

func point14Val0DependencySnapshotModel() Point14Val0DependencySnapshot {
	return cachedFormalModel(&point14Val0DependencySnapshotModelOnce, &point14Val0DependencySnapshotModelCached, func() Point14Val0DependencySnapshot {
		valE := ComputePoint13ValEFoundation(Point13ValEFoundationModel())
		point12 := ComputePoint12ValEFoundation(Point12ValEFoundationModel())
		point11 := point14Val0ActivePoint11ValDFoundation()
		point10 := operability.ComputeDeploymentMultiTenantValEFoundation(operability.DeploymentMultiTenantValEFoundationModel())
		return point14Val0DependencySnapshotFromUpstream(valE, point12, point11, point10)
	})
}

func EvaluatePoint14Val0DependencyState(model Point14Val0DependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point13ValEComputedFromUpstream ||
		!model.Point13ValEMerged ||
		!model.Point13ValECIGreen ||
		!model.Point13ValEReviewedOnMain ||
		model.Point14PassSeen ||
		model.Point13ValEPointID != point13Val0PointID ||
		model.Point13ValEWaveID != point13ValEWaveID ||
		!point13ValEStateValid(model.Point13ValECurrentState) ||
		!point13ValEStateValid(model.Point13ValEDependencyState) ||
		!point13ValEStateValid(model.Point13ValENoOverclaimState) ||
		!point13ValEStateValid(model.Point13ValEAuthorityBoundaryState) ||
		!point13ValEStateValid(model.Point13ValETimestampIntegrityState) ||
		!point13ValEStateValid(model.Point13ValETwitterIsolationState) ||
		!point13ValEStateValid(model.Point13ValEEvidenceIntegrityState) ||
		!point13ValEStateValid(model.Point13ValEPassClosureManifestState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		!point12ValEReviewerResultValid(model.InheritedPoint12ReviewerResult) ||
		model.InheritedPoint11CurrentState == "" ||
		!point14Val0Point11FoundationActive(model.Point11) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return Point14Val0StateBlocked
	}
	if model.Point13ValECurrentState != model.Point13ValE.CurrentState ||
		model.Point13ValEDependencyState != model.Point13ValE.DependencyState ||
		model.Point13ValENoOverclaimState != model.Point13ValE.NoOverclaimFinalCheckState ||
		model.Point13ValEAuthorityBoundaryState != model.Point13ValE.AuthorityBoundaryCheckState ||
		model.Point13ValETimestampIntegrityState != model.Point13ValE.TimestampIntegrityCheckState ||
		model.Point13ValETwitterIsolationState != model.Point13ValE.TenantIsolationCheckState ||
		model.Point13ValEEvidenceIntegrityState != model.Point13ValE.EvidenceIntegrityCheckState ||
		model.Point13ValEPassClosureManifestState != model.Point13ValE.PassClosureManifestState ||
		model.Point13ValEPassAllowed != model.Point13ValE.Point13PassAllowed ||
		model.Point13ValEPassToken != model.Point13ValE.Point13PassToken ||
		model.Point13ValEComputedFromUpstream != model.Point13ValE.Dependency.SnapshotFromComputedOutput ||
		model.InheritedPoint12CurrentState != model.Point13ValE.Dependency.InheritedPoint12CurrentState ||
		model.InheritedPoint12DependencyState != model.Point13ValE.Dependency.InheritedPoint12DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point13ValE.Dependency.InheritedPoint12PassClosureState ||
		model.InheritedPoint12ReviewerResult != model.Point13ValE.Dependency.InheritedPoint12ReviewerResult ||
		model.InheritedTenantScope != model.Point13ValE.Dependency.InheritedTenantScope ||
		model.InheritedPoint12CurrentState != model.Point12.CurrentState ||
		model.InheritedPoint12DependencyState != model.Point12.DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point12.PassClosureManifestState ||
		model.InheritedPoint12ReviewerResult != model.Point12.PassClosureManifest.ReviewerResult ||
		model.InheritedPoint11CurrentState != model.Point11.CurrentState ||
		model.InheritedPoint11PublicationState != model.Point11.PublicationReviewState ||
		model.InheritedPoint11NoOverclaimState != model.Point11.NoOverclaimReviewState ||
		model.InheritedPoint11FinalPassGateState != model.Point11.FinalPassGateState ||
		model.InheritedPoint10CurrentState != model.Point10.Point10State ||
		model.InheritedPoint10NoOverclaimState != model.Point10.NoOverclaimState ||
		model.InheritedPoint10ProjectionState != model.Point10.ProjectionBoundaryState ||
		model.InheritedPoint10PassRuleState != model.Point10.Point10PassRuleState {
		return Point14Val0StateBlocked
	}
	if model.Point13ValECurrentState != Point13ValEStatePassConfirmed ||
		model.Point13ValEDependencyState != Point13ValEStateActive ||
		model.Point13ValENoOverclaimState != Point13ValEStateActive ||
		model.Point13ValEAuthorityBoundaryState != Point13ValEStateActive ||
		model.Point13ValETimestampIntegrityState != Point13ValEStateActive ||
		model.Point13ValETwitterIsolationState != Point13ValEStateActive ||
		model.Point13ValEEvidenceIntegrityState != Point13ValEStateActive ||
		model.Point13ValEPassClosureManifestState != Point13ValEStateActive ||
		!model.Point13ValEPassAllowed ||
		model.Point13ValEPassToken != point13ValEPoint13PassToken ||
		model.InheritedPoint12CurrentState != Point12ValEStatePassConfirmed ||
		model.InheritedPoint12DependencyState != Point12ValEStateActive ||
		model.InheritedPoint12PassClosureState != Point12ValEStateActive ||
		model.InheritedPoint12ReviewerResult != point12ValEReviewerResultPassConfirmed ||
		model.InheritedPoint11CurrentState != Point11ValDStateActive ||
		model.InheritedPoint11FinalPassGateState != Point11ValDFinalPassGateStateActive ||
		model.InheritedPoint11PublicationState != Point11ValDPublicationReviewStateActive ||
		model.InheritedPoint11NoOverclaimState != Point11ValDNoOverclaimReviewStateActive ||
		model.InheritedPoint10CurrentState != operability.DeploymentMultiTenantPoint10StatePass ||
		model.InheritedPoint10NoOverclaimState != operability.DeploymentMultiTenantValENoOverclaimStateActive ||
		model.InheritedPoint10ProjectionState != operability.DeploymentMultiTenantValEProjectionBoundaryStateActive ||
		model.InheritedPoint10PassRuleState != operability.DeploymentMultiTenantValEPoint10PassRuleStateActive {
		return Point14Val0StateBlocked
	}
	return Point14Val0StateActive
}

func point14Val0ExternalSignalCandidateModel(dependency Point14Val0DependencySnapshot) ExternalEcosystemSignalCandidate {
	model := ExternalEcosystemSignalCandidate{
		SignalID:                  "signal_point14_val0_vendor_001",
		SourceType:                "vendor_advisory",
		SourceRef:                 "source_point14_val0_vendor_001",
		SignalType:                "evidence_submission",
		ScopeClassification:       point14Val0ScopeTenantScoped,
		TenantScope:               dependency.InheritedTenantScope,
		ReferencedTenantScope:     dependency.InheritedTenantScope,
		ArtifactRef:               "artifact_point14_val0_component_001",
		ArtifactBindingRequired:   true,
		ArtifactBindingConsistent: true,
		ClaimRefs:                 []string{"claim_point14_val0_001"},
		ClaimBindingRequired:      true,
		ClaimBindingConsistent:    true,
		EvidenceRefs:              []string{"evidence_point14_val0_001"},
		SourceEventAt:             "2026-05-05T08:45:00Z",
		SourceEventTimeSource:     point14Val0TimeSourceApprovedCustomerTime,
		ReceivedAt:                "2026-05-05T09:15:00Z",
		ReceivedTimeSource:        point14Val0TimeSourceServerUTC,
		ValidationStatus:          point14Val0ValidationReceived,
		CustodyRef:                "custody_point14_val0_001",
		HashRef:                   "hash_point14_val0_001",
		SignatureRef:              "signature_point14_val0_001",
		ProvenanceRef:             "provenance_point14_val0_001",
		RequireHash:               true,
		RequireSignature:          true,
		RequireProvenance:         true,
		DuplicateSignalRefs:       nil,
		CanonicalAuthority:        false,
		PassAllowed:               false,
		CanonicalMutationAllowed:  false,
		ProductionMutationAllowed: false,
		OverrideCanonicalDecision: false,
	}
	model.SignalIdentityKey = point14Val0SignalIdentityKey(model)
	return model
}

func EvaluatePoint14Val0ExternalSignalCandidateState(model ExternalEcosystemSignalCandidate, dependency Point14Val0DependencySnapshot) string {
	if !point14Val0SignalIDValid(model.SignalID) ||
		!point14Val0ExactValueValid(model.SourceType, point14Val0SourceTypes()) ||
		!point14Val0SourceRefValid(model.SourceRef) ||
		!point14Val0ExactValueValid(model.SignalType, point14Val0SignalTypes()) ||
		!point14Val0ExactValueValid(model.ScopeClassification, point14Val0ScopeClassifications()) ||
		!point14Val0ExactValueValid(model.ValidationStatus, point14Val0ValidationStatuses()) ||
		!point14Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point14Val0TimeSourceValid(model.ReceivedTimeSource) ||
		!point14Val0CanonicalTimeSourceValid(model.ReceivedTimeSource) ||
		!point14Val0IdentityKeyValid(model.SignalIdentityKey) ||
		!point14Val0ParsedTimeOk(model.ReceivedAt) {
		return Point14Val0StateBlocked
	}
	if model.ArtifactBindingRequired && !point14Val0ArtifactRefValid(model.ArtifactRef) {
		return Point14Val0StateBlocked
	}
	if model.ClaimBindingRequired && !point14Val0ClaimRefsValid(model.ClaimRefs) {
		return Point14Val0StateBlocked
	}
	if !model.ArtifactBindingConsistent || !model.ClaimBindingConsistent {
		return Point14Val0StateBlocked
	}
	if model.ScopeClassification == point14Val0ScopeTenantScoped {
		if !point11Val0ScopeValid(model.TenantScope) {
			return Point14Val0StateBlocked
		}
	} else {
		if model.TenantScope != "" || model.ReferencedTenantScope != "" {
			return Point14Val0StateBlocked
		}
	}
	if model.ReferencedTenantScope != "" &&
		model.ReferencedTenantScope != dependency.InheritedTenantScope {
		return Point14Val0StateBlocked
	}
	if model.TenantScope != "" &&
		model.TenantScope != dependency.InheritedTenantScope {
		return Point14Val0StateBlocked
	}
	if model.RequireHash && !point14Val0HashRefValid(model.HashRef) {
		return Point14Val0StateBlocked
	}
	if model.RequireSignature && !point14Val0SignatureRefValid(model.SignatureRef) {
		return Point14Val0StateBlocked
	}
	if model.RequireProvenance && !point14Val0ProvenanceRefValid(model.ProvenanceRef) {
		return Point14Val0StateBlocked
	}
	if model.CustodyRef != "" && !point14Val0CustodyRefValid(model.CustodyRef) {
		return Point14Val0StateBlocked
	}
	if model.SignalIdentityKey != point14Val0SignalIdentityKey(model) ||
		len(model.DuplicateSignalRefs) > 0 {
		return Point14Val0StateBlocked
	}
	if model.CanonicalAuthority ||
		model.PassAllowed ||
		model.CanonicalMutationAllowed ||
		model.ProductionMutationAllowed ||
		model.OverrideCanonicalDecision {
		return Point14Val0StateBlocked
	}
	if strings.TrimSpace(model.ValidatedAt) != "" {
		validatedAt, ok := point14Val0ParsedTime(model.ValidatedAt)
		receivedAt, okReceived := point14Val0ParsedTime(model.ReceivedAt)
		if !ok || !okReceived || model.ValidationStatus != point14Val0ValidationValidated || validatedAt.Before(receivedAt) {
			return Point14Val0StateBlocked
		}
	} else if model.ValidationStatus == point14Val0ValidationValidated {
		return Point14Val0StateIncomplete
	}
	switch model.ValidationStatus {
	case point14Val0ValidationSuperseded:
		if !point14Val0SignalIDValid(model.SupersededByRef) {
			return Point14Val0StateBlocked
		}
		return Point14Val0StateBlocked
	case point14Val0ValidationRevoked:
		if !point14Val0RevocationRefValid(model.RevocationRef) {
			return Point14Val0StateBlocked
		}
		return Point14Val0StateBlocked
	case point14Val0ValidationRejected:
		return Point14Val0StateBlocked
	case point14Val0ValidationProvenancePending, point14Val0ValidationConflicting:
		return Point14Val0StateReviewRequired
	}
	if strings.TrimSpace(model.SourceEventAt) != "" {
		sourceEventAt, okSource := point14Val0ParsedTime(model.SourceEventAt)
		receivedAt, okReceived := point14Val0ParsedTime(model.ReceivedAt)
		if !okSource || !okReceived || !point14Val0TimeSourceValid(model.SourceEventTimeSource) {
			return Point14Val0StateBlocked
		}
		if model.SourceEventTimeSource == point14Val0TimeSourceClientLocal {
			return Point14Val0StateReviewRequired
		}
		if sourceEventAt.After(receivedAt) {
			return Point14Val0StateReviewRequired
		}
	}
	return Point14Val0StateActive
}

func point14Val0ExternalStakeholderAuthorityRoleModel(dependency Point14Val0DependencySnapshot) ExternalStakeholderAuthorityRole {
	return ExternalStakeholderAuthorityRole{
		RoleID:                       "role_point14_val0_external_researcher_001",
		RoleType:                     "external_researcher",
		StakeholderRef:               "stakeholder_point14_val0_external_researcher_001",
		ScopeClassification:          point14Val0ScopeTenantScoped,
		TenantScope:                  dependency.InheritedTenantScope,
		AllowedActionRefs:            []string{"submit_evidence", "submit_dispute"},
		GovernanceRequiredActionRefs: []string{"resolve_conflict", "publish_correction_notice"},
	}
}

func EvaluatePoint14Val0ExternalStakeholderAuthorityRoleState(model ExternalStakeholderAuthorityRole) string {
	if !point14Val0RoleIDValid(model.RoleID) ||
		!point14Val0ExactValueValid(model.RoleType, point14Val0RoleTypes()) ||
		!point14Val0StakeholderRefValid(model.StakeholderRef) ||
		!point14Val0ExactValueValid(model.ScopeClassification, point14Val0ScopeClassifications()) ||
		len(model.AllowedActionRefs) == 0 {
		return Point14Val0StateBlocked
	}
	if model.ScopeClassification == point14Val0ScopeTenantScoped {
		if !point11Val0ScopeValid(model.TenantScope) {
			return Point14Val0StateBlocked
		}
	} else if model.TenantScope != "" {
		return Point14Val0StateBlocked
	}
	for _, action := range model.AllowedActionRefs {
		if !point14Val0ExactValueValid(action, point14Val0BoundedRoleActions()) ||
			point14Val0ExactValueValid(action, point14Val0ForbiddenRoleActions()) {
			return Point14Val0StateBlocked
		}
	}
	for _, action := range model.GovernanceRequiredActionRefs {
		if !point14Val0ExactValueValid(action, point14Val0GovernanceRequiredActions()) {
			return Point14Val0StateBlocked
		}
	}
	return Point14Val0StateActive
}

func point14Val0ExternalAuthorityConflictMatrixModel(dependency Point14Val0DependencySnapshot) ExternalAuthorityConflictMatrix {
	return ExternalAuthorityConflictMatrix{
		ConflictPresent:     false,
		ConflictID:          "conflict_point14_val0_none_001",
		AuditEventRef:       "audit_event_point14_val0_conflict_001",
		TenantScope:         dependency.InheritedTenantScope,
		AffectedTenantScope: dependency.InheritedTenantScope,
	}
}

func EvaluatePoint14Val0ExternalAuthorityConflictMatrixState(model ExternalAuthorityConflictMatrix) string {
	if !point14Val0ConflictIDValid(model.ConflictID) ||
		!point14Val0AuditEventRefValid(model.AuditEventRef) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point11Val0ScopeValid(model.AffectedTenantScope) ||
		model.AffectedTenantScope != model.TenantScope {
		return Point14Val0StateBlocked
	}
	if !model.ConflictPresent {
		if model.ConflictType != "" ||
			len(model.SignalRefs) > 0 ||
			len(model.RoleRefs) > 0 ||
			len(model.AffectedArtifactRefs) > 0 ||
			len(model.AffectedClaimRefs) > 0 ||
			len(model.AffectedEvidenceRefs) > 0 ||
			model.GovernancePath != "" ||
			model.ResolvedByGovernance ||
			model.GovernanceResolutionEventRef != "" ||
			model.ConsensusResolutionRequested ||
			model.SilentActiveResolutionRequested {
			return Point14Val0StateBlocked
		}
		return Point14Val0StateActive
	}
	if !point14Val0ConflictIDValid(model.ConflictID) ||
		!point14Val0ExactValueValid(model.ConflictType, point14Val0ConflictTypes()) ||
		!point14Val0SignalRefsValid(model.SignalRefs) ||
		!point14Val0RoleRefsValid(model.RoleRefs) ||
		!point14Val0EvidenceRefsValid(model.AffectedEvidenceRefs) ||
		!point14Val0ExactValueValid(model.GovernancePath, point14Val0GovernancePaths()) {
		return Point14Val0StateBlocked
	}
	if len(model.AffectedArtifactRefs) > 0 && !point14Val0RefListValid(model.AffectedArtifactRefs, "artifact_") {
		return Point14Val0StateBlocked
	}
	if len(model.AffectedClaimRefs) > 0 && !point14Val0ClaimRefsValid(model.AffectedClaimRefs) {
		return Point14Val0StateBlocked
	}
	if model.ConsensusResolutionRequested || model.SilentActiveResolutionRequested {
		return Point14Val0StateBlocked
	}
	if model.ResolvedByGovernance {
		if !point14Val0GovernanceEventRefValid(model.GovernanceResolutionEventRef) {
			return Point14Val0StateBlocked
		}
		return Point14Val0StateActive
	}
	return Point14Val0StateReviewRequired
}

func point14Val0ExternalSignalDisputeLifecycleModel(dependency Point14Val0DependencySnapshot) ExternalSignalDisputeLifecycle {
	return ExternalSignalDisputeLifecycle{
		DisputeID:                  "dispute_point14_val0_opened_001",
		DisputeState:               point14Val0DisputeOpened,
		TenantScope:                dependency.InheritedTenantScope,
		AffectedTenantScope:        dependency.InheritedTenantScope,
		SignalRefs:                 []string{"signal_point14_val0_vendor_001"},
		EvidenceRefs:               []string{"evidence_point14_val0_001"},
		AuditEventRefs:             []string{"audit_event_point14_val0_dispute_001"},
		OpenedAt:                   "2026-05-05T09:20:00Z",
		TimestampSource:            point14Val0TimeSourceServerUTC,
		CanonicalMutationRequested: false,
		RejectedDeletesEvidence:    false,
		PrivacyCheckPassed:         true,
	}
}

func EvaluatePoint14Val0ExternalSignalDisputeLifecycleState(model ExternalSignalDisputeLifecycle) string {
	if !point14Val0DisputeIDValid(model.DisputeID) ||
		!point14Val0ExactValueValid(model.DisputeState, []string{
			point14Val0DisputeOpened,
			point14Val0DisputeTriaged,
			point14Val0DisputeEvidenceNeeded,
			point14Val0DisputeReviewRequired,
			point14Val0DisputeCorrected,
			point14Val0DisputeSuperseded,
			point14Val0DisputeRejected,
			point14Val0DisputeRevoked,
			point14Val0DisputePublished,
		}) ||
		!point14Val0SignalRefsValid(model.SignalRefs) ||
		!point14Val0AuditEventRefsValid(model.AuditEventRefs) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point14Val0CanonicalTimeSourceValid(model.TimestampSource) ||
		!point14Val0ParsedTimeOk(model.OpenedAt) {
		return Point14Val0StateBlocked
	}
	if model.AffectedTenantScope != "" &&
		model.AffectedTenantScope != model.TenantScope {
		return Point14Val0StateBlocked
	}
	if model.CanonicalMutationRequested {
		return Point14Val0StateBlocked
	}
	if model.DisputeState == point14Val0DisputeOpened || model.DisputeState == point14Val0DisputeTriaged {
		if !point14Val0EvidenceRefsValid(model.EvidenceRefs) {
			return Point14Val0StateBlocked
		}
		return Point14Val0StateActive
	}
	if model.DisputeState == point14Val0DisputeEvidenceNeeded {
		if !point14Val0EvidenceRefsValid(model.EvidenceRefs) {
			return Point14Val0StateIncomplete
		}
		return Point14Val0StateReviewRequired
	}
	if model.DisputeState == point14Val0DisputeReviewRequired {
		return Point14Val0StateReviewRequired
	}
	openedAt, _ := point14Val0ParsedTime(model.OpenedAt)
	if model.DisputeState == point14Val0DisputeCorrected {
		correctedAt, ok := point14Val0ParsedTime(model.CorrectedAt)
		if !ok || !point14Val0GovernanceEventRefValid(model.GovernanceEventRef) || correctedAt.Before(openedAt) {
			return Point14Val0StateBlocked
		}
		return Point14Val0StateActive
	}
	if model.DisputeState == point14Val0DisputeRevoked {
		revokedAt, ok := point14Val0ParsedTime(model.RevokedAt)
		if !ok || !point14Val0GovernanceEventRefValid(model.GovernanceEventRef) || revokedAt.Before(openedAt) {
			return Point14Val0StateBlocked
		}
		return Point14Val0StateActive
	}
	if model.DisputeState == point14Val0DisputeSuperseded {
		if !point14Val0DisputeIDValid(model.SupersessionRef) {
			return Point14Val0StateBlocked
		}
		return Point14Val0StateActive
	}
	if model.DisputeState == point14Val0DisputePublished {
		publishedAt, ok := point14Val0ParsedTime(model.PublishedAt)
		if !ok || !point14Val0BoundaryRefValid(model.VisibilityBoundaryRef) || !model.PrivacyCheckPassed || publishedAt.Before(openedAt) {
			return Point14Val0StateBlocked
		}
		return Point14Val0StateActive
	}
	if model.DisputeState == point14Val0DisputeRejected && model.RejectedDeletesEvidence {
		return Point14Val0StateBlocked
	}
	return Point14Val0StateActive
}

func point14Val0ExternalCorrectionRevocationBoundaryModel(dependency Point14Val0DependencySnapshot) ExternalCorrectionRevocationBoundary {
	return ExternalCorrectionRevocationBoundary{
		CorrectionRef:           "correction_point14_val0_001",
		PublicPrivateVisibility: point14Val0VisibilityTenantPrivate,
		TenantScope:             dependency.InheritedTenantScope,
		PrivacyBoundaryRef:      "privacy_boundary_point14_val0_001",
		ApproverRef:             "approver_point14_val0_001",
		GovernanceEventRef:      "governance_event_point14_val0_correction_001",
		AuditEventRef:           "audit_event_point14_val0_correction_001",
		LimitationRefs:          []string{"limitation_ref_point14_val0_001"},
		ObservedTexts:           []string{"correction notice pending approval"},
		OpenedAt:                "2026-05-05T09:20:00Z",
		CorrectedAt:             "2026-05-05T09:30:00Z",
		TimestampSource:         point14Val0TimeSourceServerUTC,
	}
}

func EvaluatePoint14Val0ExternalCorrectionRevocationBoundaryState(model ExternalCorrectionRevocationBoundary) string {
	if !point14Val0CorrectionRefValid(model.CorrectionRef) ||
		!point14Val0ExactValueValid(model.PublicPrivateVisibility, point14Val0VisibilityScopes()) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point14Val0BoundaryRefValid(model.PrivacyBoundaryRef) ||
		!point14Val0ApproverRefValid(model.ApproverRef) ||
		!point14Val0GovernanceEventRefValid(model.GovernanceEventRef) ||
		!point14Val0AuditEventRefValid(model.AuditEventRef) ||
		!point14Val0CanonicalTimeSourceValid(model.TimestampSource) ||
		!point14Val0ParsedTimeOk(model.OpenedAt) {
		return Point14Val0StateBlocked
	}
	if !point14Val0RefListValid(model.LimitationRefs, "limitation_ref_") {
		return Point14Val0StateBlocked
	}
	if model.RevocationModeled && !point14Val0RevocationRefValid(model.RevocationRef) {
		return Point14Val0StateBlocked
	}
	if model.LeaksTenantPrivateData ||
		model.StrengthensClaim ||
		model.CreatesForbiddenAuthorityWords ||
		model.OmitsMeaningChangingLimitation ||
		model.HidesDecisiveMissingEvidence ||
		model.ClientTimeBackdated ||
		point14Val0ListContainsForbiddenWording(model.ObservedTexts) {
		return Point14Val0StateBlocked
	}
	openedAt, _ := point14Val0ParsedTime(model.OpenedAt)
	for _, pair := range []struct {
		value string
		name  string
	}{
		{model.CorrectedAt, "corrected"},
		{model.RevokedAt, "revoked"},
		{model.PublishedAt, "published"},
	} {
		if strings.TrimSpace(pair.value) == "" {
			continue
		}
		parsed, ok := point14Val0ParsedTime(pair.value)
		if !ok || parsed.Before(openedAt) {
			return Point14Val0StateBlocked
		}
	}
	return Point14Val0StateActive
}

func point14Val0ExternalVisibilityPublicationBoundaryModel(dependency Point14Val0DependencySnapshot) ExternalVisibilityPublicationBoundary {
	return ExternalVisibilityPublicationBoundary{
		VisibilityID:     "visibility_point14_val0_001",
		VisibilityScope:  point14Val0VisibilityScopedCustomer,
		TenantScope:      dependency.InheritedTenantScope,
		PublicOutputRefs: []string{"publication_output_point14_val0_001"},
		PrivateScopeRefs: []string{"private_scope_point14_val0_001"},
		LimitationTexts:  []string{"bounded ecosystem evidence input"},
		ObservedTexts:    []string{"advisory external signal", "no external authority granted"},
		AuditEventRef:    "audit_event_point14_val0_visibility_001",
	}
}

func EvaluatePoint14Val0ExternalVisibilityPublicationBoundaryState(model ExternalVisibilityPublicationBoundary) string {
	if !point14Val0VisibilityIDValid(model.VisibilityID) ||
		!point14Val0ExactValueValid(model.VisibilityScope, point14Val0VisibilityScopes()) ||
		!point14Val0OutputRefsValid(model.PublicOutputRefs) ||
		!point14Val0PrivateScopeRefsValid(model.PrivateScopeRefs) ||
		!point14Val0TextListValid(model.LimitationTexts) ||
		!point14Val0AuditEventRefValid(model.AuditEventRef) {
		return Point14Val0StateBlocked
	}
	if model.VisibilityScope != point14Val0VisibilityPublicNoticeLimited && !point11Val0ScopeValid(model.TenantScope) {
		return Point14Val0StateBlocked
	}
	if model.ContainsPrivateTenantData || model.OmitsLimitations || point14Val0ListContainsForbiddenWording(model.ObservedTexts) {
		return Point14Val0StateBlocked
	}
	if model.VisibilityScope == point14Val0VisibilityPublicNoticeLimited && len(model.LimitationTexts) == 0 {
		return Point14Val0StateBlocked
	}
	return Point14Val0StateActive
}

func point14Val0AgentEcosystemInputBoundaryModel(dependency Point14Val0DependencySnapshot) AgentEcosystemInputBoundary {
	return AgentEcosystemInputBoundary{
		BoundaryID:               "boundary_point14_val0_agent_001",
		TenantScope:              dependency.InheritedTenantScope,
		AgentInputRefs:           []string{"agent_input_point14_val0_001"},
		EvidenceRefs:             []string{"evidence_point14_val0_001"},
		AuditEventRef:            "audit_event_point14_val0_agent_001",
		AdvisoryOnly:             true,
		CanDecideDispute:         false,
		CanPublishCorrection:     false,
		CanRevokeClaim:           false,
		CanOverrideConflict:      false,
		CanEmitPass:              false,
		CanEmitPublicAuthority:   false,
		PassAllowed:              false,
		ApprovalGranted:          false,
		ProductionApproved:       false,
		DeploymentApproved:       false,
		ExternalAuthorityAllowed: false,
	}
}

func EvaluatePoint14Val0AgentEcosystemInputBoundaryState(model AgentEcosystemInputBoundary) string {
	if !point14Val0BoundaryRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point14Val0AgentInputRefsValid(model.AgentInputRefs) ||
		!point14Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point14Val0AuditEventRefValid(model.AuditEventRef) {
		return Point14Val0StateBlocked
	}
	if !model.AdvisoryOnly ||
		model.CanDecideDispute ||
		model.CanPublishCorrection ||
		model.CanRevokeClaim ||
		model.CanOverrideConflict ||
		model.CanEmitPass ||
		model.CanEmitPublicAuthority ||
		model.PassAllowed ||
		model.ApprovalGranted ||
		model.ProductionApproved ||
		model.DeploymentApproved ||
		model.ExternalAuthorityAllowed {
		return Point14Val0StateBlocked
	}
	return Point14Val0StateActive
}

func point14Val0NoExternalAuthorityGuardModel() Point14Val0NoExternalAuthorityGuard {
	return Point14Val0NoExternalAuthorityGuard{
		ObservedAuthorityMarkers: nil,
	}
}

func EvaluatePoint14Val0NoExternalAuthorityGuardState(model Point14Val0NoExternalAuthorityGuard) string {
	if point14Val0AuthorityMarkersContainForbidden(model.ObservedAuthorityMarkers, point14Val0ForbiddenAuthorityMarkers()) {
		return Point14Val0StateBlocked
	}
	return Point14Val0StateActive
}

func point14Val0NoOverclaimEcosystemWordingModel() Point14Val0NoOverclaimEcosystemWording {
	return Point14Val0NoOverclaimEcosystemWording{
		ObservedSignalTexts:      []string{"external evidence signal", "evidence input pending validation"},
		ObservedDisputeTexts:     []string{"dispute requires governance review"},
		ObservedCorrectionTexts:  []string{"correction notice pending approval"},
		ObservedPublicationTexts: []string{"bounded ecosystem evidence input"},
		ObservedAgentTexts:       []string{"advisory external signal", "no external authority granted"},
		AllowedSafeWording:       point14Val0SafeWording(),
		BlockedWording:           point14Val0ForbiddenWording(),
		ProjectionDisclaimer:     point14Val0ProjectionDisclaimerBase,
	}
}

func EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model Point14Val0NoOverclaimEcosystemWording) string {
	if model.ProjectionDisclaimer != point14Val0ProjectionDisclaimerBase ||
		!point12Val0ExactStringSetMatch(model.AllowedSafeWording, point14Val0SafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point14Val0ForbiddenWording()) {
		return Point14Val0StateBlocked
	}
	observedTexts := append([]string{}, model.ObservedSignalTexts...)
	observedTexts = append(observedTexts, model.ObservedDisputeTexts...)
	observedTexts = append(observedTexts, model.ObservedCorrectionTexts...)
	observedTexts = append(observedTexts, model.ObservedPublicationTexts...)
	observedTexts = append(observedTexts, model.ObservedAgentTexts...)
	if point14Val0ListContainsForbiddenWording(observedTexts) {
		return Point14Val0StateBlocked
	}
	if point14Val0ListContainsForbiddenWording(model.InternalDiagnosticTexts) &&
		!model.InternalDiagnosticsClassifiedBlocked {
		return Point14Val0StateBlocked
	}
	return Point14Val0StateActive
}

func point14Val0FoundationState(states ...string) string {
	hasReview := false
	hasIncomplete := false
	for _, state := range states {
		switch state {
		case Point14Val0StateBlocked:
			return Point14Val0StateBlocked
		case Point14Val0StateReviewRequired:
			hasReview = true
		case Point14Val0StateIncomplete:
			hasIncomplete = true
		case Point14Val0StateActive:
			continue
		default:
			return Point14Val0StateBlocked
		}
	}
	if hasReview {
		return Point14Val0StateReviewRequired
	}
	if hasIncomplete {
		return Point14Val0StateIncomplete
	}
	return Point14Val0StateActive
}

func point14Val0BlockingReasons(model Point14Val0Foundation) []string {
	reasons := []string{}
	componentStates := map[string]string{
		"dependency":                     model.DependencyState,
		"external_signal_candidate":      model.ExternalSignalCandidateState,
		"external_stakeholder_role":      model.ExternalStakeholderAuthorityRoleState,
		"conflict_matrix":                model.ExternalAuthorityConflictMatrixState,
		"dispute_lifecycle":              model.ExternalSignalDisputeLifecycleState,
		"correction_revocation_boundary": model.ExternalCorrectionRevocationState,
		"visibility_publication":         model.ExternalVisibilityPublicationState,
		"agent_boundary":                 model.AgentEcosystemInputBoundaryState,
		"no_external_authority_guard":    model.NoExternalAuthorityGuardState,
		"no_overclaim":                   model.NoOverclaimState,
	}
	for name, state := range componentStates {
		if state == Point14Val0StateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point14Val0ReviewPrerequisites(model Point14Val0Foundation) []string {
	prereqs := []string{}
	componentStates := map[string]string{
		"dependency":                     model.DependencyState,
		"external_signal_candidate":      model.ExternalSignalCandidateState,
		"external_stakeholder_role":      model.ExternalStakeholderAuthorityRoleState,
		"conflict_matrix":                model.ExternalAuthorityConflictMatrixState,
		"dispute_lifecycle":              model.ExternalSignalDisputeLifecycleState,
		"correction_revocation_boundary": model.ExternalCorrectionRevocationState,
		"visibility_publication":         model.ExternalVisibilityPublicationState,
		"agent_boundary":                 model.AgentEcosystemInputBoundaryState,
		"no_external_authority_guard":    model.NoExternalAuthorityGuardState,
		"no_overclaim":                   model.NoOverclaimState,
	}
	for name, state := range componentStates {
		if state == Point14Val0StateReviewRequired || state == Point14Val0StateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func point14Val0ParsedTimeOk(value string) bool {
	_, ok := point14Val0ParsedTime(value)
	return ok
}

func point14Val0FoundationModelFromUpstream(
	valE Point13ValEFoundation,
	point12 Point12ValEFoundation,
	point11 Point11ValDFoundation,
	point10 operability.DeploymentMultiTenantValEFoundation,
) Point14Val0Foundation {
	dependency := point14Val0DependencySnapshotFromUpstream(valE, point12, point11, point10)
	return Point14Val0Foundation{
		CurrentState:                          Point14Val0StateActive,
		ProjectionDisclaimer:                  point14Val0ProjectionDisclaimerBase,
		DependencyState:                       Point14Val0StateActive,
		ExternalSignalCandidateState:          Point14Val0StateActive,
		ExternalStakeholderAuthorityRoleState: Point14Val0StateActive,
		ExternalAuthorityConflictMatrixState:  Point14Val0StateActive,
		ExternalSignalDisputeLifecycleState:   Point14Val0StateActive,
		ExternalCorrectionRevocationState:     Point14Val0StateActive,
		ExternalVisibilityPublicationState:    Point14Val0StateActive,
		AgentEcosystemInputBoundaryState:      Point14Val0StateActive,
		NoExternalAuthorityGuardState:         Point14Val0StateActive,
		NoOverclaimState:                      Point14Val0StateActive,
		Dependency:                            dependency,
		ExternalSignalCandidate:               point14Val0ExternalSignalCandidateModel(dependency),
		ExternalStakeholderAuthorityRole:      point14Val0ExternalStakeholderAuthorityRoleModel(dependency),
		ExternalAuthorityConflictMatrix:       point14Val0ExternalAuthorityConflictMatrixModel(dependency),
		ExternalSignalDisputeLifecycle:        point14Val0ExternalSignalDisputeLifecycleModel(dependency),
		ExternalCorrectionRevocationBoundary:  point14Val0ExternalCorrectionRevocationBoundaryModel(dependency),
		ExternalVisibilityPublicationBoundary: point14Val0ExternalVisibilityPublicationBoundaryModel(dependency),
		AgentEcosystemInputBoundary:           point14Val0AgentEcosystemInputBoundaryModel(dependency),
		NoExternalAuthorityGuard:              point14Val0NoExternalAuthorityGuardModel(),
		NoOverclaimEcosystemWording:           point14Val0NoOverclaimEcosystemWordingModel(),
	}
}

func Point14Val0FoundationModel() Point14Val0Foundation {
	return cachedFormalModel(&point14Val0FoundationModelOnce, &point14Val0FoundationModelCached, func() Point14Val0Foundation {
		valE := ComputePoint13ValEFoundation(Point13ValEFoundationModel())
		point12 := ComputePoint12ValEFoundation(Point12ValEFoundationModel())
		point11 := point14Val0ActivePoint11ValDFoundation()
		point10 := operability.ComputeDeploymentMultiTenantValEFoundation(operability.DeploymentMultiTenantValEFoundationModel())
		return point14Val0FoundationModelFromUpstream(valE, point12, point11, point10)
	})
}

func ComputePoint14Val0Foundation(model Point14Val0Foundation) Point14Val0Foundation {
	model.DependencyState = EvaluatePoint14Val0DependencyState(model.Dependency)
	model.ExternalSignalCandidateState = EvaluatePoint14Val0ExternalSignalCandidateState(model.ExternalSignalCandidate, model.Dependency)
	model.ExternalStakeholderAuthorityRoleState = EvaluatePoint14Val0ExternalStakeholderAuthorityRoleState(model.ExternalStakeholderAuthorityRole)
	model.ExternalAuthorityConflictMatrixState = EvaluatePoint14Val0ExternalAuthorityConflictMatrixState(model.ExternalAuthorityConflictMatrix)
	model.ExternalSignalDisputeLifecycleState = EvaluatePoint14Val0ExternalSignalDisputeLifecycleState(model.ExternalSignalDisputeLifecycle)
	model.ExternalCorrectionRevocationState = EvaluatePoint14Val0ExternalCorrectionRevocationBoundaryState(model.ExternalCorrectionRevocationBoundary)
	model.ExternalVisibilityPublicationState = EvaluatePoint14Val0ExternalVisibilityPublicationBoundaryState(model.ExternalVisibilityPublicationBoundary)
	model.AgentEcosystemInputBoundaryState = EvaluatePoint14Val0AgentEcosystemInputBoundaryState(model.AgentEcosystemInputBoundary)
	model.NoExternalAuthorityGuardState = EvaluatePoint14Val0NoExternalAuthorityGuardState(model.NoExternalAuthorityGuard)
	model.NoOverclaimState = EvaluatePoint14Val0NoOverclaimEcosystemWordingState(model.NoOverclaimEcosystemWording)
	model.CurrentState = point14Val0FoundationState(
		model.DependencyState,
		model.ExternalSignalCandidateState,
		model.ExternalStakeholderAuthorityRoleState,
		model.ExternalAuthorityConflictMatrixState,
		model.ExternalSignalDisputeLifecycleState,
		model.ExternalCorrectionRevocationState,
		model.ExternalVisibilityPublicationState,
		model.AgentEcosystemInputBoundaryState,
		model.NoExternalAuthorityGuardState,
		model.NoOverclaimState,
	)
	model.BlockingReasons = point14Val0BlockingReasons(model)
	model.ReviewPrerequisites = point14Val0ReviewPrerequisites(model)
	return model
}

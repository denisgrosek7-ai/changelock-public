package formal

import (
	"sort"
	"strings"

	"github.com/denisgrosek/changelock/internal/operability"
)

const (
	Point14ValBStateActive         = "point14_valb_external_conflict_dispute_triage_active"
	Point14ValBStateBlocked        = "point14_valb_external_conflict_dispute_triage_blocked"
	Point14ValBStateReviewRequired = "point14_valb_external_conflict_dispute_triage_review_required"
	Point14ValBStateIncomplete     = "point14_valb_external_conflict_dispute_triage_incomplete"
)

const (
	point14ValBWaveID                   = "val_b"
	point14ValBProjectionDisclaimerBase = "projection_only not_canonical_truth point14_valb_external_conflict_dispute_triage_gate"

	point14ValBTriageNoConflictDetected   = "no_conflict_detected"
	point14ValBTriageReviewRequired       = "conflict_review_required"
	point14ValBTriageEvidenceRequired     = "conflict_evidence_required"
	point14ValBTriageBlocked              = "conflict_blocked"
	point14ValBTriageUnsupported          = "conflict_unsupported"
	point14ValBTriageCrossTenant          = "conflict_cross_tenant"
	point14ValBTriagePrivacyBlocked       = "conflict_privacy_blocked"
	point14ValBTriageGovernanceEscalated  = "conflict_governance_escalated"
	point14ValBEscalationStateNotRequired = "not_required"
	point14ValBEscalationStateQueued      = "queued"
	point14ValBEscalationStateAssigned    = "assigned"
	point14ValBEscalationStateInReview    = "in_review"
	point14ValBEscalationStateCompleted   = "completed"
)

type Point14ValBDependencySnapshot struct {
	Point14ValACurrentState                                   string                                          `json:"point14_vala_current_state"`
	Point14ValADependencyState                                string                                          `json:"point14_vala_dependency_state"`
	Point14ValANormalizedSignalState                          string                                          `json:"point14_vala_normalized_signal_state"`
	Point14ValASourceIdentityState                            string                                          `json:"point14_vala_source_identity_state"`
	Point14ValAScopeBindingState                              string                                          `json:"point14_vala_scope_binding_state"`
	Point14ValAEvidenceBindingState                           string                                          `json:"point14_vala_evidence_binding_state"`
	Point14ValAFreshnessTimestampState                        string                                          `json:"point14_vala_freshness_timestamp_state"`
	Point14ValADuplicateRelationGuardState                    string                                          `json:"point14_vala_duplicate_relation_guard_state"`
	Point14ValATenantBoundaryGuardState                       string                                          `json:"point14_vala_tenant_boundary_guard_state"`
	Point14ValAValidationResultState                          string                                          `json:"point14_vala_validation_result_state"`
	Point14ValANoExternalAuthorityState                       string                                          `json:"point14_vala_no_external_authority_state"`
	Point14ValANoOverclaimState                               string                                          `json:"point14_vala_no_overclaim_state"`
	Point14ValAPointID                                        string                                          `json:"point14_vala_point_id"`
	Point14ValAWaveID                                         string                                          `json:"point14_vala_wave_id"`
	Point14ValAComputedFromUpstream                           bool                                            `json:"point14_vala_computed_from_upstream"`
	Point14ValAMerged                                         bool                                            `json:"point14_vala_merged"`
	Point14ValACIGreen                                        bool                                            `json:"point14_vala_ci_green"`
	Point14ValAReviewedOnMain                                 bool                                            `json:"point14_vala_reviewed_on_main"`
	Point14PassSeen                                           bool                                            `json:"point14_pass_seen"`
	InheritedPoint14Val0CurrentState                          string                                          `json:"inherited_point14_val0_current_state"`
	InheritedPoint14Val0DependencyState                       string                                          `json:"inherited_point14_val0_dependency_state"`
	InheritedPoint14Val0ExternalSignalCandidateState          string                                          `json:"inherited_point14_val0_external_signal_candidate_state"`
	InheritedPoint14Val0ExternalStakeholderAuthorityRoleState string                                          `json:"inherited_point14_val0_external_stakeholder_authority_role_state"`
	InheritedPoint14Val0ExternalAuthorityConflictMatrixState  string                                          `json:"inherited_point14_val0_external_authority_conflict_matrix_state"`
	InheritedPoint14Val0ExternalSignalDisputeLifecycleState   string                                          `json:"inherited_point14_val0_external_signal_dispute_lifecycle_state"`
	InheritedPoint14Val0ExternalCorrectionRevocationState     string                                          `json:"inherited_point14_val0_external_correction_revocation_state"`
	InheritedPoint14Val0ExternalVisibilityPublicationState    string                                          `json:"inherited_point14_val0_external_visibility_publication_state"`
	InheritedPoint14Val0AgentEcosystemInputBoundaryState      string                                          `json:"inherited_point14_val0_agent_ecosystem_input_boundary_state"`
	InheritedPoint14Val0NoExternalAuthorityGuardState         string                                          `json:"inherited_point14_val0_no_external_authority_guard_state"`
	InheritedPoint14Val0NoOverclaimState                      string                                          `json:"inherited_point14_val0_no_overclaim_state"`
	InheritedPoint13ValECurrentState                          string                                          `json:"inherited_point13_vale_current_state"`
	InheritedPoint13ValEPassClosureState                      string                                          `json:"inherited_point13_vale_pass_closure_state"`
	InheritedPoint13ValEPassAllowed                           bool                                            `json:"inherited_point13_vale_pass_allowed"`
	InheritedPoint13ValEPassToken                             string                                          `json:"inherited_point13_vale_pass_token"`
	InheritedPoint12CurrentState                              string                                          `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState                           string                                          `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureState                          string                                          `json:"inherited_point12_pass_closure_state"`
	InheritedPoint12ReviewerResult                            string                                          `json:"inherited_point12_reviewer_result"`
	InheritedPoint11PublicationState                          string                                          `json:"inherited_point11_publication_state"`
	InheritedPoint11NoOverclaimState                          string                                          `json:"inherited_point11_no_overclaim_state"`
	InheritedPoint11FinalPassGateState                        string                                          `json:"inherited_point11_final_pass_gate_state"`
	InheritedPoint10CurrentState                              string                                          `json:"inherited_point10_current_state"`
	InheritedPoint10NoOverclaimState                          string                                          `json:"inherited_point10_no_overclaim_state"`
	InheritedPoint10ProjectionState                           string                                          `json:"inherited_point10_projection_state"`
	InheritedPoint10PassRuleState                             string                                          `json:"inherited_point10_pass_rule_state"`
	InheritedTenantScope                                      string                                          `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput                                bool                                            `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                                       []string                                        `json:"review_prerequisites,omitempty"`
	Point14ValA                                               Point14ValAFoundation                           `json:"point14_vala"`
	Point14Val0                                               Point14Val0Foundation                           `json:"point14_val0"`
	Point13ValE                                               Point13ValEFoundation                           `json:"point13_vale"`
	Point12                                                   Point12ValEFoundation                           `json:"point12"`
	Point11                                                   Point11ValDFoundation                           `json:"point11"`
	Point10                                                   operability.DeploymentMultiTenantValEFoundation `json:"point10"`
}

type ExternalSignalConflictSet struct {
	ConflictSetID             string   `json:"conflict_set_id"`
	TenantScope               string   `json:"tenant_scope"`
	GlobalScopeClassification string   `json:"global_scope_classification"`
	NormalizedSignalRefs      []string `json:"normalized_signal_refs,omitempty"`
	ValidationResultRefs      []string `json:"validation_result_refs,omitempty"`
	StakeholderRoleRefs       []string `json:"stakeholder_role_refs,omitempty"`
	AffectedArtifactRefs      []string `json:"affected_artifact_refs,omitempty"`
	ArtifactScoped            bool     `json:"artifact_scoped"`
	AffectedClaimRefs         []string `json:"affected_claim_refs,omitempty"`
	ClaimScoped               bool     `json:"claim_scoped"`
	AffectedEvidenceRefs      []string `json:"affected_evidence_refs,omitempty"`
	ConflictType              string   `json:"conflict_type"`
	ConflictDetected          bool     `json:"conflict_detected"`
	ConflictSourceRefs        []string `json:"conflict_source_refs,omitempty"`
	CanonicalContextRef       string   `json:"canonical_context_ref"`
	ResolveToPass             bool     `json:"resolve_to_pass"`
	ResolveToCanonicalTruth   bool     `json:"resolve_to_canonical_truth"`
	ApproveProduction         bool     `json:"approve_production"`
	PublishCorrection         bool     `json:"publish_correction"`
	RevokeClaim               bool     `json:"revoke_claim"`
	CreatePublicBadge         bool     `json:"create_public_badge"`
}

type StakeholderSignalComparison struct {
	ComparisonID                      string   `json:"comparison_id"`
	StakeholderRoles                  []string `json:"stakeholder_roles,omitempty"`
	SignalRefs                        []string `json:"signal_refs,omitempty"`
	AuthorityScopeRefs                []string `json:"authority_scope_refs,omitempty"`
	CanonicalContextRef               string   `json:"canonical_context_ref"`
	ComparisonResult                  string   `json:"comparison_result"`
	ResolvesConflict                  bool     `json:"resolves_conflict"`
	EmitsPass                         bool     `json:"emits_pass"`
	CertifiesCompliance               bool     `json:"certifies_compliance"`
	ApprovesProduction                bool     `json:"approves_production"`
	PublishesPublicAuthority          bool     `json:"publishes_public_authority"`
	CrowdConsensusResolutionRequested bool     `json:"crowd_consensus_resolution_requested"`
}

type ExternalConflictTriageResult struct {
	TriageResultID             string `json:"triage_result_id"`
	ConflictSetRef             string `json:"conflict_set_ref"`
	StakeholderComparisonRef   string `json:"stakeholder_comparison_ref"`
	DisputeIntakeRef           string `json:"dispute_intake_ref"`
	EvidenceRequirementRef     string `json:"evidence_requirement_ref"`
	GovernanceEscalationRef    string `json:"governance_escalation_ref"`
	TenantPrivacyBoundaryRef   string `json:"tenant_privacy_boundary_ref"`
	AgentBoundaryRef           string `json:"agent_boundary_ref"`
	ConflictSetState           string `json:"conflict_set_state"`
	StakeholderComparisonState string `json:"stakeholder_comparison_state"`
	DisputeIntakeState         string `json:"dispute_intake_state"`
	EvidenceRequirementState   string `json:"evidence_requirement_state"`
	GovernanceEscalationState  string `json:"governance_escalation_state"`
	TenantPrivacyBoundaryState string `json:"tenant_privacy_boundary_state"`
	AgentBoundaryState         string `json:"agent_boundary_state"`
	NoExternalAuthorityState   string `json:"no_external_authority_state"`
	NoOverclaimState           string `json:"no_overclaim_state"`
	TriageState                string `json:"triage_state"`
	EmitsPass                  bool   `json:"emits_pass"`
	PublishesCorrection        bool   `json:"publishes_correction"`
	RevokesClaim               bool   `json:"revokes_claim"`
	OverridesCanonicalDecision bool   `json:"overrides_canonical_decision"`
	ApprovesProduction         bool   `json:"approves_production"`
	CertifiesCompliance        bool   `json:"certifies_compliance"`
	CreatesPublicBadge         bool   `json:"creates_public_badge"`
}

type DisputeIntakePacket struct {
	DisputeID                  string   `json:"dispute_id"`
	OpenedByRole               string   `json:"opened_by_role"`
	OpenedAt                   string   `json:"opened_at"`
	OpenedTimeSource           string   `json:"opened_time_source"`
	TenantScope                string   `json:"tenant_scope"`
	GlobalScopeClassification  string   `json:"global_scope_classification"`
	AffectedSignalRefs         []string `json:"affected_signal_refs,omitempty"`
	AffectedArtifactRefs       []string `json:"affected_artifact_refs,omitempty"`
	ArtifactScoped             bool     `json:"artifact_scoped"`
	AffectedClaimRefs          []string `json:"affected_claim_refs,omitempty"`
	ClaimScoped                bool     `json:"claim_scoped"`
	AffectedEvidenceRefs       []string `json:"affected_evidence_refs,omitempty"`
	DisputeReason              string   `json:"dispute_reason"`
	LifecycleState             string   `json:"lifecycle_state"`
	CanonicalMutationRequested bool     `json:"canonical_mutation_requested"`
	SourceEventAt              string   `json:"source_event_at"`
	SourceEventTimeSource      string   `json:"source_event_time_source"`
	ClosureAttempted           bool     `json:"closure_attempted"`
	ClosedAt                   string   `json:"closed_at"`
}

type DisputeEvidenceRequirementGate struct {
	EvidenceRequirementID         string   `json:"evidence_requirement_id"`
	RequiredEvidenceRefs          []string `json:"required_evidence_refs,omitempty"`
	RequiredEvidenceTypes         []string `json:"required_evidence_types,omitempty"`
	DecisiveEvidenceMissing       bool     `json:"decisive_evidence_missing"`
	EvidenceState                 string   `json:"evidence_state"`
	AgentRecommendationOnly       bool     `json:"agent_recommendation_only"`
	CrowdConsensusOnly            bool     `json:"crowd_consensus_only"`
	AuditorNoteOnly               bool     `json:"auditor_note_only"`
	ConflictTriageBypassRequested bool     `json:"conflict_triage_bypass_requested"`
}

type GovernanceEscalationPath struct {
	EscalationID        string `json:"escalation_id"`
	EscalationRequired  bool   `json:"escalation_required"`
	GovernanceEventRef  string `json:"governance_event_ref"`
	Owner               string `json:"owner"`
	ApproverRole        string `json:"approver_role"`
	Reason              string `json:"reason"`
	AuditRef            string `json:"audit_ref"`
	EscalationState     string `json:"escalation_state"`
	ApprovesProduction  bool   `json:"approves_production"`
	PublishesCorrection bool   `json:"publishes_correction"`
}

type TenantPrivacyConflictBoundary struct {
	BoundaryID                string   `json:"boundary_id"`
	TenantScope               string   `json:"tenant_scope"`
	VisibilityScope           string   `json:"visibility_scope"`
	PublicVisibilityRequested bool     `json:"public_visibility_requested"`
	RedactionRefs             []string `json:"redaction_refs,omitempty"`
	LimitationRefs            []string `json:"limitation_refs,omitempty"`
	LimitationVisible         bool     `json:"limitation_visible"`
	ObservedTexts             []string `json:"observed_texts,omitempty"`
	TenantPrivateDataExposed  bool     `json:"tenant_private_data_exposed"`
	StrengthensClaim          bool     `json:"strengthens_claim"`
}

type AgentDisputeRecommendationBoundary struct {
	BoundaryID                         string   `json:"boundary_id"`
	TenantScope                        string   `json:"tenant_scope"`
	AgentInputRefs                     []string `json:"agent_input_refs,omitempty"`
	EvidenceRefs                       []string `json:"evidence_refs,omitempty"`
	AuditEventRef                      string   `json:"audit_event_ref"`
	AdvisoryOnly                       bool     `json:"advisory_only"`
	CanResolveConflict                 bool     `json:"can_resolve_conflict"`
	CanPublishCorrection               bool     `json:"can_publish_correction"`
	CanRevokeClaim                     bool     `json:"can_revoke_claim"`
	CanOverrideGovernance              bool     `json:"can_override_governance"`
	CanSatisfyEvidenceRequirementAlone bool     `json:"can_satisfy_evidence_requirement_alone"`
	CanEmitPass                        bool     `json:"can_emit_pass"`
	CanEmitPublicAuthority             bool     `json:"can_emit_public_authority"`
	PassAllowed                        bool     `json:"pass_allowed"`
	ApprovalGranted                    bool     `json:"approval_granted"`
	ProductionApproved                 bool     `json:"production_approved"`
	ExternalAuthorityAllowed           bool     `json:"external_authority_allowed"`
}

type Point14ValBNoExternalAuthorityConflictGuard struct {
	ObservedAuthorityMarkers  []string `json:"observed_authority_markers,omitempty"`
	CanonicalAuthorityGranted bool     `json:"canonical_authority_granted"`
	ProductionApprovalGranted bool     `json:"production_approval_granted"`
	PublishesCorrection       bool     `json:"publishes_correction"`
	RevokesClaim              bool     `json:"revokes_claim"`
	PublicBadgeAllowed        bool     `json:"public_badge_allowed"`
	DisputeAutoResolved       bool     `json:"dispute_auto_resolved"`
	CrowdResolved             bool     `json:"crowd_resolved"`
	AgentResolved             bool     `json:"agent_resolved"`
	ExternalAuthorityAllowed  bool     `json:"external_authority_allowed"`
}

type Point14ValBNoOverclaimDisputeWording struct {
	ObservedConflictTexts                []string `json:"observed_conflict_texts,omitempty"`
	ObservedDisputeTexts                 []string `json:"observed_dispute_texts,omitempty"`
	ObservedEvidenceRequirementTexts     []string `json:"observed_evidence_requirement_texts,omitempty"`
	ObservedEscalationTexts              []string `json:"observed_escalation_texts,omitempty"`
	ObservedPrivacyTexts                 []string `json:"observed_privacy_texts,omitempty"`
	ObservedAgentTexts                   []string `json:"observed_agent_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point14ValBFoundation struct {
	CurrentState                       string                                      `json:"current_state"`
	BlockingReasons                    []string                                    `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites                []string                                    `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer               string                                      `json:"projection_disclaimer"`
	DependencyState                    string                                      `json:"dependency_state"`
	ConflictSetState                   string                                      `json:"conflict_set_state"`
	StakeholderComparisonState         string                                      `json:"stakeholder_comparison_state"`
	DisputeTriageResultState           string                                      `json:"dispute_triage_result_state"`
	DisputeIntakeState                 string                                      `json:"dispute_intake_state"`
	EvidenceRequirementState           string                                      `json:"evidence_requirement_state"`
	GovernanceEscalationState          string                                      `json:"governance_escalation_state"`
	TenantPrivacyBoundaryState         string                                      `json:"tenant_privacy_boundary_state"`
	AgentDisputeBoundaryState          string                                      `json:"agent_dispute_boundary_state"`
	NoExternalAuthorityState           string                                      `json:"no_external_authority_state"`
	NoOverclaimState                   string                                      `json:"no_overclaim_state"`
	Dependency                         Point14ValBDependencySnapshot               `json:"dependency"`
	ConflictSet                        ExternalSignalConflictSet                   `json:"conflict_set"`
	StakeholderComparison              StakeholderSignalComparison                 `json:"stakeholder_comparison"`
	DisputeTriageResult                ExternalConflictTriageResult                `json:"dispute_triage_result"`
	DisputeIntake                      DisputeIntakePacket                         `json:"dispute_intake"`
	EvidenceRequirementGate            DisputeEvidenceRequirementGate              `json:"evidence_requirement_gate"`
	GovernanceEscalationPath           GovernanceEscalationPath                    `json:"governance_escalation_path"`
	TenantPrivacyConflictBoundary      TenantPrivacyConflictBoundary               `json:"tenant_privacy_conflict_boundary"`
	AgentDisputeRecommendationBoundary AgentDisputeRecommendationBoundary          `json:"agent_dispute_recommendation_boundary"`
	NoExternalAuthorityConflictGuard   Point14ValBNoExternalAuthorityConflictGuard `json:"no_external_authority_conflict_guard"`
	NoOverclaimDisputeWording          Point14ValBNoOverclaimDisputeWording        `json:"no_overclaim_dispute_wording"`
}

func point14ValBStates() []string {
	return []string{
		Point14ValBStateActive,
		Point14ValBStateBlocked,
		Point14ValBStateReviewRequired,
		Point14ValBStateIncomplete,
	}
}

func point14ValBStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point14ValBStates(), value)
}

func point14ValBConflictTypes() []string {
	return []string{
		"scanner_vs_vex",
		"scanner_vs_maintainer",
		"vex_vs_canonical_evidence",
		"maintainer_vs_reviewer",
		"auditor_note_vs_insufficient_evidence",
		"verifier_note_vs_canonical_decision",
		"partner_signal_vs_customer_scope",
		"duplicate_external_signal_conflict",
		"stale_external_signal_conflict",
		"cross_tenant_signal_conflict",
		"public_claim_vs_private_evidence",
		"agent_recommendation_vs_governance",
	}
}

func point14ValBComparisonResults() []string {
	return []string{
		"consistent",
		"conflicting",
		"insufficient_context",
		"unsupported_role",
		"stale_signal",
		"cross_tenant_conflict",
		"review_required",
	}
}

func point14ValBTriageStates() []string {
	return []string{
		point14ValBTriageNoConflictDetected,
		point14ValBTriageReviewRequired,
		point14ValBTriageEvidenceRequired,
		point14ValBTriageBlocked,
		point14ValBTriageUnsupported,
		point14ValBTriageCrossTenant,
		point14ValBTriagePrivacyBlocked,
		point14ValBTriageGovernanceEscalated,
	}
}

func point14ValBDisputeLifecycleStates() []string {
	return []string{
		point14Val0DisputeOpened,
		point14Val0DisputeTriaged,
		point14Val0DisputeEvidenceNeeded,
		point14Val0DisputeReviewRequired,
		point14Val0DisputeRejected,
	}
}

func point14ValBEvidenceTypes() []string {
	return []string{
		"canonical_evidence",
		"artifact_binding",
		"claim_binding",
		"provenance_record",
		"custody_record",
		"hash_match",
		"stakeholder_review",
	}
}

func point14ValBEscalationStates() []string {
	return []string{
		point14ValBEscalationStateNotRequired,
		point14ValBEscalationStateQueued,
		point14ValBEscalationStateAssigned,
		point14ValBEscalationStateInReview,
		point14ValBEscalationStateCompleted,
	}
}

func point14ValBForbiddenAuthorityMarkers() []string {
	markers := append([]string{}, point14ValAForbiddenAuthorityMarkers()...)
	return append(markers,
		"scanner_pass",
		"dispute_auto_resolved",
		"crowd_resolved",
		"agent_resolved",
	)
}

func point14ValBSafeWording() []string {
	return []string{
		"external signal conflict detected",
		"dispute requires governance review",
		"evidence required before triage closure",
		"advisory conflict input",
		"no external authority granted",
		"no automatic correction publication",
		"bounded dispute triage state",
		"canonical evidence spine remains source of truth",
	}
}

func point14ValBForbiddenWording() []string {
	return append(
		append([]string{}, point14Val0ForbiddenWording()...),
		"dispute resolved by crowd",
		"dispute resolved by ai",
	)
}

func point14ValBObservedTextContainsForbiddenWording(text string) bool {
	trimmed := strings.ToLower(strings.TrimSpace(text))
	if trimmed == "" {
		return false
	}
	for _, safe := range point14ValBSafeWording() {
		if trimmed == strings.ToLower(strings.TrimSpace(safe)) {
			return false
		}
	}
	for _, phrase := range point14ValBForbiddenWording() {
		if strings.Contains(trimmed, strings.ToLower(strings.TrimSpace(phrase))) {
			return true
		}
	}
	return false
}

func point14ValBObservedListContainsForbiddenWording(values []string) bool {
	for _, value := range values {
		if point14ValBObservedTextContainsForbiddenWording(value) {
			return true
		}
	}
	return false
}

func point14ValBConflictSetIDValid(value string) bool {
	return point14Val0RefValid(value, "conflict_set_")
}

func point14ValBComparisonIDValid(value string) bool {
	return point14Val0RefValid(value, "comparison_")
}

func point14ValBTriageResultIDValid(value string) bool {
	return point14Val0RefValid(value, "triage_result_")
}

func point14ValBEvidenceRequirementIDValid(value string) bool {
	return point14Val0RefValid(value, "evidence_requirement_")
}

func point14ValBEscalationIDValid(value string) bool {
	return point14Val0RefValid(value, "escalation_")
}

func point14ValBCanonicalContextRefValid(value string) bool {
	return point14Val0RefValid(value, "canonical_context_")
}

func point14ValBValidationResultRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "validation_result_")
}

func point14ValBNormalizedSignalRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "normalized_signal_")
}

func point14ValBAuthorityScopeRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "authority_scope_")
}

func point14ValBRedactionRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "redaction_ref_")
}

func point14ValBLimitationRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "limitation_ref_")
}

func point14ValBExactValueValid(value string, allowed []string) bool {
	return point11Val0ContainsTrimmed(allowed, value)
}

func point14ValBEvidenceTypesValid(values []string) bool {
	if len(values) == 0 {
		return false
	}
	for _, value := range values {
		if !point14ValBExactValueValid(value, point14ValBEvidenceTypes()) {
			return false
		}
	}
	return true
}

func point14ValBTriageStateClassification(value string) (string, bool, bool) {
	switch strings.TrimSpace(value) {
	case point14ValBTriageBlocked,
		point14ValBTriageUnsupported,
		point14ValBTriageCrossTenant,
		point14ValBTriagePrivacyBlocked:
		return Point14ValBStateBlocked, true, true
	case point14ValBTriageReviewRequired,
		point14ValBTriageGovernanceEscalated:
		return Point14ValBStateReviewRequired, false, true
	case point14ValBTriageEvidenceRequired:
		return Point14ValBStateIncomplete, false, true
	case point14ValBTriageNoConflictDetected:
		return Point14ValBStateActive, false, true
	default:
		return Point14ValBStateBlocked, false, false
	}
}

func point14ValBResolvedTriageState(current, aggregate string) string {
	current = strings.TrimSpace(current)
	currentState, blockedLike, recognized := point14ValBTriageStateClassification(current)
	resolvedAggregate := point14ValBFoundationState(currentState, strings.TrimSpace(aggregate))
	switch resolvedAggregate {
	case Point14ValBStateBlocked:
		if recognized && blockedLike {
			return current
		}
		return point14ValBTriageBlocked
	case Point14ValBStateReviewRequired:
		if recognized && (blockedLike || current == point14ValBTriageReviewRequired || current == point14ValBTriageGovernanceEscalated) {
			return current
		}
		return point14ValBTriageReviewRequired
	case Point14ValBStateIncomplete:
		if recognized && (blockedLike || current == point14ValBTriageReviewRequired || current == point14ValBTriageGovernanceEscalated) {
			return current
		}
		return point14ValBTriageEvidenceRequired
	default:
		if recognized {
			return current
		}
		return point14ValBTriageBlocked
	}
}

func point14ValBDependencySnapshotFromUpstream(valA Point14ValAFoundation) Point14ValBDependencySnapshot {
	return Point14ValBDependencySnapshot{
		Point14ValACurrentState:                                   valA.CurrentState,
		Point14ValADependencyState:                                valA.DependencyState,
		Point14ValANormalizedSignalState:                          valA.NormalizedSignalState,
		Point14ValASourceIdentityState:                            valA.SourceIdentityState,
		Point14ValAScopeBindingState:                              valA.ScopeBindingState,
		Point14ValAEvidenceBindingState:                           valA.EvidenceBindingState,
		Point14ValAFreshnessTimestampState:                        valA.FreshnessTimestampState,
		Point14ValADuplicateRelationGuardState:                    valA.DuplicateRelationGuardState,
		Point14ValATenantBoundaryGuardState:                       valA.TenantBoundaryGuardState,
		Point14ValAValidationResultState:                          valA.ValidationResultState,
		Point14ValANoExternalAuthorityState:                       valA.NoExternalAuthorityState,
		Point14ValANoOverclaimState:                               valA.NoOverclaimState,
		Point14ValAPointID:                                        point14Val0PointID,
		Point14ValAWaveID:                                         point14ValAWaveID,
		Point14ValAComputedFromUpstream:                           valA.Dependency.SnapshotFromComputedOutput,
		Point14ValAMerged:                                         true,
		Point14ValACIGreen:                                        true,
		Point14ValAReviewedOnMain:                                 true,
		Point14PassSeen:                                           false,
		InheritedPoint14Val0CurrentState:                          valA.Dependency.Point14Val0CurrentState,
		InheritedPoint14Val0DependencyState:                       valA.Dependency.Point14Val0DependencyState,
		InheritedPoint14Val0ExternalSignalCandidateState:          valA.Dependency.Point14Val0ExternalSignalCandidateState,
		InheritedPoint14Val0ExternalStakeholderAuthorityRoleState: valA.Dependency.Point14Val0ExternalStakeholderAuthorityRoleState,
		InheritedPoint14Val0ExternalAuthorityConflictMatrixState:  valA.Dependency.Point14Val0ExternalAuthorityConflictMatrixState,
		InheritedPoint14Val0ExternalSignalDisputeLifecycleState:   valA.Dependency.Point14Val0ExternalSignalDisputeLifecycleState,
		InheritedPoint14Val0ExternalCorrectionRevocationState:     valA.Dependency.Point14Val0ExternalCorrectionRevocationState,
		InheritedPoint14Val0ExternalVisibilityPublicationState:    valA.Dependency.Point14Val0ExternalVisibilityPublicationState,
		InheritedPoint14Val0AgentEcosystemInputBoundaryState:      valA.Dependency.Point14Val0AgentEcosystemInputBoundaryState,
		InheritedPoint14Val0NoExternalAuthorityGuardState:         valA.Dependency.Point14Val0NoExternalAuthorityGuardState,
		InheritedPoint14Val0NoOverclaimState:                      valA.Dependency.Point14Val0NoOverclaimState,
		InheritedPoint13ValECurrentState:                          valA.Dependency.InheritedPoint13ValECurrentState,
		InheritedPoint13ValEPassClosureState:                      valA.Dependency.InheritedPoint13ValEPassClosureState,
		InheritedPoint13ValEPassAllowed:                           valA.Dependency.InheritedPoint13ValEPassAllowed,
		InheritedPoint13ValEPassToken:                             valA.Dependency.InheritedPoint13ValEPassToken,
		InheritedPoint12CurrentState:                              valA.Dependency.InheritedPoint12CurrentState,
		InheritedPoint12DependencyState:                           valA.Dependency.InheritedPoint12DependencyState,
		InheritedPoint12PassClosureState:                          valA.Dependency.InheritedPoint12PassClosureState,
		InheritedPoint12ReviewerResult:                            valA.Dependency.InheritedPoint12ReviewerResult,
		InheritedPoint11PublicationState:                          valA.Dependency.InheritedPoint11PublicationState,
		InheritedPoint11NoOverclaimState:                          valA.Dependency.InheritedPoint11NoOverclaimState,
		InheritedPoint11FinalPassGateState:                        valA.Dependency.InheritedPoint11FinalPassGateState,
		InheritedPoint10CurrentState:                              valA.Dependency.InheritedPoint10CurrentState,
		InheritedPoint10NoOverclaimState:                          valA.Dependency.InheritedPoint10NoOverclaimState,
		InheritedPoint10ProjectionState:                           valA.Dependency.InheritedPoint10ProjectionState,
		InheritedPoint10PassRuleState:                             valA.Dependency.InheritedPoint10PassRuleState,
		InheritedTenantScope:                                      valA.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:                                true,
		Point14ValA:                                               valA,
		Point14Val0:                                               valA.Dependency.Point14Val0,
		Point13ValE:                                               valA.Dependency.Point13ValE,
		Point12:                                                   valA.Dependency.Point12,
		Point11:                                                   valA.Dependency.Point11,
		Point10:                                                   valA.Dependency.Point10,
	}
}

func point14ValBDependencySnapshotModel() Point14ValBDependencySnapshot {
	valA := ComputePoint14ValAFoundation(Point14ValAFoundationModel())
	return point14ValBDependencySnapshotFromUpstream(valA)
}

func EvaluatePoint14ValBDependencyState(model Point14ValBDependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point14ValAComputedFromUpstream ||
		!model.Point14ValAMerged ||
		!model.Point14ValACIGreen ||
		!model.Point14ValAReviewedOnMain ||
		model.Point14PassSeen ||
		model.Point14ValAPointID != point14Val0PointID ||
		model.Point14ValAWaveID != point14ValAWaveID ||
		!point14ValAStateValid(model.Point14ValACurrentState) ||
		!point14ValAStateValid(model.Point14ValADependencyState) ||
		!point14ValAStateValid(model.Point14ValANormalizedSignalState) ||
		!point14ValAStateValid(model.Point14ValASourceIdentityState) ||
		!point14ValAStateValid(model.Point14ValAScopeBindingState) ||
		!point14ValAStateValid(model.Point14ValAEvidenceBindingState) ||
		!point14ValAStateValid(model.Point14ValAFreshnessTimestampState) ||
		!point14ValAStateValid(model.Point14ValADuplicateRelationGuardState) ||
		!point14ValAStateValid(model.Point14ValATenantBoundaryGuardState) ||
		!point14ValAStateValid(model.Point14ValAValidationResultState) ||
		!point14ValAStateValid(model.Point14ValANoExternalAuthorityState) ||
		!point14ValAStateValid(model.Point14ValANoOverclaimState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0CurrentState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0DependencyState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0ExternalSignalCandidateState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0ExternalStakeholderAuthorityRoleState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0ExternalAuthorityConflictMatrixState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0ExternalSignalDisputeLifecycleState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0ExternalCorrectionRevocationState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0ExternalVisibilityPublicationState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0AgentEcosystemInputBoundaryState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0NoExternalAuthorityGuardState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0NoOverclaimState) ||
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
		return Point14ValBStateBlocked
	}
	if model.Point14ValACurrentState != model.Point14ValA.CurrentState ||
		model.Point14ValADependencyState != model.Point14ValA.DependencyState ||
		model.Point14ValANormalizedSignalState != model.Point14ValA.NormalizedSignalState ||
		model.Point14ValASourceIdentityState != model.Point14ValA.SourceIdentityState ||
		model.Point14ValAScopeBindingState != model.Point14ValA.ScopeBindingState ||
		model.Point14ValAEvidenceBindingState != model.Point14ValA.EvidenceBindingState ||
		model.Point14ValAFreshnessTimestampState != model.Point14ValA.FreshnessTimestampState ||
		model.Point14ValADuplicateRelationGuardState != model.Point14ValA.DuplicateRelationGuardState ||
		model.Point14ValATenantBoundaryGuardState != model.Point14ValA.TenantBoundaryGuardState ||
		model.Point14ValAValidationResultState != model.Point14ValA.ValidationResultState ||
		model.Point14ValANoExternalAuthorityState != model.Point14ValA.NoExternalAuthorityState ||
		model.Point14ValANoOverclaimState != model.Point14ValA.NoOverclaimState ||
		model.Point14ValAComputedFromUpstream != model.Point14ValA.Dependency.SnapshotFromComputedOutput ||
		model.InheritedPoint14Val0CurrentState != model.Point14ValA.Dependency.Point14Val0CurrentState ||
		model.InheritedPoint14Val0DependencyState != model.Point14ValA.Dependency.Point14Val0DependencyState ||
		model.InheritedPoint14Val0ExternalSignalCandidateState != model.Point14ValA.Dependency.Point14Val0ExternalSignalCandidateState ||
		model.InheritedPoint14Val0ExternalStakeholderAuthorityRoleState != model.Point14ValA.Dependency.Point14Val0ExternalStakeholderAuthorityRoleState ||
		model.InheritedPoint14Val0ExternalAuthorityConflictMatrixState != model.Point14ValA.Dependency.Point14Val0ExternalAuthorityConflictMatrixState ||
		model.InheritedPoint14Val0ExternalSignalDisputeLifecycleState != model.Point14ValA.Dependency.Point14Val0ExternalSignalDisputeLifecycleState ||
		model.InheritedPoint14Val0ExternalCorrectionRevocationState != model.Point14ValA.Dependency.Point14Val0ExternalCorrectionRevocationState ||
		model.InheritedPoint14Val0ExternalVisibilityPublicationState != model.Point14ValA.Dependency.Point14Val0ExternalVisibilityPublicationState ||
		model.InheritedPoint14Val0AgentEcosystemInputBoundaryState != model.Point14ValA.Dependency.Point14Val0AgentEcosystemInputBoundaryState ||
		model.InheritedPoint14Val0NoExternalAuthorityGuardState != model.Point14ValA.Dependency.Point14Val0NoExternalAuthorityGuardState ||
		model.InheritedPoint14Val0NoOverclaimState != model.Point14ValA.Dependency.Point14Val0NoOverclaimState ||
		model.InheritedPoint13ValECurrentState != model.Point14ValA.Dependency.InheritedPoint13ValECurrentState ||
		model.InheritedPoint13ValEPassClosureState != model.Point14ValA.Dependency.InheritedPoint13ValEPassClosureState ||
		model.InheritedPoint13ValEPassAllowed != model.Point14ValA.Dependency.InheritedPoint13ValEPassAllowed ||
		model.InheritedPoint13ValEPassToken != model.Point14ValA.Dependency.InheritedPoint13ValEPassToken ||
		model.InheritedPoint12CurrentState != model.Point14ValA.Dependency.InheritedPoint12CurrentState ||
		model.InheritedPoint12DependencyState != model.Point14ValA.Dependency.InheritedPoint12DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point14ValA.Dependency.InheritedPoint12PassClosureState ||
		model.InheritedPoint12ReviewerResult != model.Point14ValA.Dependency.InheritedPoint12ReviewerResult ||
		model.InheritedPoint11PublicationState != model.Point14ValA.Dependency.InheritedPoint11PublicationState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValA.Dependency.InheritedPoint11NoOverclaimState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValA.Dependency.InheritedPoint11FinalPassGateState ||
		model.InheritedPoint10CurrentState != model.Point14ValA.Dependency.InheritedPoint10CurrentState ||
		model.InheritedPoint10NoOverclaimState != model.Point14ValA.Dependency.InheritedPoint10NoOverclaimState ||
		model.InheritedPoint10ProjectionState != model.Point14ValA.Dependency.InheritedPoint10ProjectionState ||
		model.InheritedPoint10PassRuleState != model.Point14ValA.Dependency.InheritedPoint10PassRuleState ||
		model.InheritedTenantScope != model.Point14ValA.Dependency.InheritedTenantScope ||
		model.InheritedPoint14Val0CurrentState != model.Point14Val0.CurrentState ||
		model.InheritedPoint14Val0DependencyState != model.Point14Val0.DependencyState ||
		model.InheritedPoint14Val0ExternalSignalCandidateState != model.Point14Val0.ExternalSignalCandidateState ||
		model.InheritedPoint14Val0ExternalStakeholderAuthorityRoleState != model.Point14Val0.ExternalStakeholderAuthorityRoleState ||
		model.InheritedPoint14Val0ExternalAuthorityConflictMatrixState != model.Point14Val0.ExternalAuthorityConflictMatrixState ||
		model.InheritedPoint14Val0ExternalSignalDisputeLifecycleState != model.Point14Val0.ExternalSignalDisputeLifecycleState ||
		model.InheritedPoint14Val0ExternalCorrectionRevocationState != model.Point14Val0.ExternalCorrectionRevocationState ||
		model.InheritedPoint14Val0ExternalVisibilityPublicationState != model.Point14Val0.ExternalVisibilityPublicationState ||
		model.InheritedPoint14Val0AgentEcosystemInputBoundaryState != model.Point14Val0.AgentEcosystemInputBoundaryState ||
		model.InheritedPoint14Val0NoExternalAuthorityGuardState != model.Point14Val0.NoExternalAuthorityGuardState ||
		model.InheritedPoint14Val0NoOverclaimState != model.Point14Val0.NoOverclaimState ||
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
		return Point14ValBStateBlocked
	}
	if model.Point14ValACurrentState != Point14ValAStateActive ||
		model.Point14ValADependencyState != Point14ValAStateActive ||
		model.Point14ValANormalizedSignalState != Point14ValAStateActive ||
		model.Point14ValASourceIdentityState != Point14ValAStateActive ||
		model.Point14ValAScopeBindingState != Point14ValAStateActive ||
		model.Point14ValAEvidenceBindingState != Point14ValAStateActive ||
		model.Point14ValAFreshnessTimestampState != Point14ValAStateActive ||
		model.Point14ValADuplicateRelationGuardState != Point14ValAStateActive ||
		model.Point14ValATenantBoundaryGuardState != Point14ValAStateActive ||
		model.Point14ValAValidationResultState != Point14ValAStateActive ||
		model.Point14ValANoExternalAuthorityState != Point14ValAStateActive ||
		model.Point14ValANoOverclaimState != Point14ValAStateActive ||
		model.InheritedPoint14Val0CurrentState != Point14Val0StateActive ||
		model.InheritedPoint14Val0DependencyState != Point14Val0StateActive ||
		model.InheritedPoint14Val0ExternalSignalCandidateState != Point14Val0StateActive ||
		model.InheritedPoint14Val0ExternalStakeholderAuthorityRoleState != Point14Val0StateActive ||
		model.InheritedPoint14Val0ExternalAuthorityConflictMatrixState != Point14Val0StateActive ||
		model.InheritedPoint14Val0ExternalSignalDisputeLifecycleState != Point14Val0StateActive ||
		model.InheritedPoint14Val0ExternalCorrectionRevocationState != Point14Val0StateActive ||
		model.InheritedPoint14Val0ExternalVisibilityPublicationState != Point14Val0StateActive ||
		model.InheritedPoint14Val0AgentEcosystemInputBoundaryState != Point14Val0StateActive ||
		model.InheritedPoint14Val0NoExternalAuthorityGuardState != Point14Val0StateActive ||
		model.InheritedPoint14Val0NoOverclaimState != Point14Val0StateActive ||
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
		return Point14ValBStateBlocked
	}
	return Point14ValBStateActive
}

func point14ValBConflictSetModel(dependency Point14ValBDependencySnapshot) ExternalSignalConflictSet {
	return ExternalSignalConflictSet{
		ConflictSetID:           "conflict_set_point14_valb_001",
		TenantScope:             dependency.InheritedTenantScope,
		NormalizedSignalRefs:    []string{"normalized_signal_point14_vala_001"},
		ValidationResultRefs:    []string{"validation_result_point14_vala_001"},
		StakeholderRoleRefs:     []string{"role_point14_valb_scanner_001", "role_point14_valb_vex_issuer_001"},
		AffectedArtifactRefs:    []string{"artifact_point14_valb_component_001"},
		ArtifactScoped:          true,
		AffectedClaimRefs:       []string{"claim_point14_valb_001"},
		ClaimScoped:             true,
		AffectedEvidenceRefs:    []string{"evidence_point14_valb_001"},
		ConflictType:            "scanner_vs_vex",
		ConflictDetected:        false,
		ConflictSourceRefs:      []string{"source_point14_valb_scanner_001", "source_point14_valb_vex_issuer_001"},
		CanonicalContextRef:     "canonical_context_point14_valb_001",
		ResolveToPass:           false,
		ResolveToCanonicalTruth: false,
		ApproveProduction:       false,
		PublishCorrection:       false,
		RevokeClaim:             false,
		CreatePublicBadge:       false,
	}
}

func EvaluatePoint14ValBExternalSignalConflictSetState(model ExternalSignalConflictSet, dependency Point14ValBDependencySnapshot) string {
	if !point14ValBConflictSetIDValid(model.ConflictSetID) ||
		!point14ValBNormalizedSignalRefsValid(model.NormalizedSignalRefs) ||
		!point14ValBValidationResultRefsValid(model.ValidationResultRefs) ||
		!point14Val0RoleRefsValid(model.StakeholderRoleRefs) ||
		!point14Val0EvidenceRefsValid(model.AffectedEvidenceRefs) ||
		!point14ValBExactValueValid(model.ConflictType, point14ValBConflictTypes()) ||
		!point14Val0RefListValid(model.ConflictSourceRefs, "source_") ||
		!point14ValBCanonicalContextRefValid(model.CanonicalContextRef) {
		return Point14ValBStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) == "" && strings.TrimSpace(model.GlobalScopeClassification) == "" {
		return Point14ValBStateBlocked
	}
	if model.TenantScope != "" {
		if !point11Val0ScopeValid(model.TenantScope) || model.TenantScope != dependency.InheritedTenantScope {
			return Point14ValBStateBlocked
		}
		if strings.TrimSpace(model.GlobalScopeClassification) != "" {
			return Point14ValBStateBlocked
		}
	} else if !point11Val0ContainsTrimmed([]string{point14Val0ScopeGlobalAdvisory, point14Val0ScopePublicNonAuthorative}, model.GlobalScopeClassification) {
		return Point14ValBStateBlocked
	}
	if model.ArtifactScoped && !point14Val0RefListValid(model.AffectedArtifactRefs, "artifact_") {
		return Point14ValBStateBlocked
	}
	if model.ClaimScoped && !point14Val0ClaimRefsValid(model.AffectedClaimRefs) {
		return Point14ValBStateBlocked
	}
	if model.ResolveToPass ||
		model.ResolveToCanonicalTruth ||
		model.ApproveProduction ||
		model.PublishCorrection ||
		model.RevokeClaim ||
		model.CreatePublicBadge {
		return Point14ValBStateBlocked
	}
	if !model.ConflictDetected {
		return Point14ValBStateActive
	}
	switch strings.TrimSpace(model.ConflictType) {
	case "cross_tenant_signal_conflict":
		return Point14ValBStateBlocked
	case "scanner_vs_vex",
		"scanner_vs_maintainer",
		"vex_vs_canonical_evidence",
		"maintainer_vs_reviewer",
		"auditor_note_vs_insufficient_evidence",
		"verifier_note_vs_canonical_decision",
		"partner_signal_vs_customer_scope",
		"duplicate_external_signal_conflict",
		"stale_external_signal_conflict",
		"public_claim_vs_private_evidence",
		"agent_recommendation_vs_governance":
		return Point14ValBStateReviewRequired
	default:
		return Point14ValBStateBlocked
	}
}

func point14ValBStakeholderComparisonModel() StakeholderSignalComparison {
	return StakeholderSignalComparison{
		ComparisonID:        "comparison_point14_valb_001",
		StakeholderRoles:    []string{"scanner", "vex_issuer"},
		SignalRefs:          []string{"signal_point14_valb_scanner_001", "signal_point14_valb_vex_issuer_001"},
		AuthorityScopeRefs:  []string{"authority_scope_point14_valb_scanner_001", "authority_scope_point14_valb_vex_issuer_001"},
		CanonicalContextRef: "canonical_context_point14_valb_001",
		ComparisonResult:    "consistent",
	}
}

func EvaluatePoint14ValBStakeholderSignalComparisonState(model StakeholderSignalComparison) string {
	if !point14ValBComparisonIDValid(model.ComparisonID) ||
		len(model.StakeholderRoles) == 0 ||
		!point14Val0SignalRefsValid(model.SignalRefs) ||
		!point14ValBAuthorityScopeRefsValid(model.AuthorityScopeRefs) ||
		!point14ValBCanonicalContextRefValid(model.CanonicalContextRef) ||
		!point14ValBExactValueValid(model.ComparisonResult, point14ValBComparisonResults()) {
		return Point14ValBStateBlocked
	}
	for _, role := range model.StakeholderRoles {
		if !point14Val0ExactValueValid(role, point14Val0RoleTypes()) {
			return Point14ValBStateBlocked
		}
	}
	if model.ResolvesConflict ||
		model.EmitsPass ||
		model.CertifiesCompliance ||
		model.ApprovesProduction ||
		model.PublishesPublicAuthority ||
		model.CrowdConsensusResolutionRequested {
		return Point14ValBStateBlocked
	}
	switch strings.TrimSpace(model.ComparisonResult) {
	case "consistent":
		return Point14ValBStateActive
	case "conflicting", "insufficient_context", "stale_signal", "review_required":
		return Point14ValBStateReviewRequired
	case "unsupported_role", "cross_tenant_conflict":
		return Point14ValBStateBlocked
	default:
		return Point14ValBStateBlocked
	}
}

func point14ValBDisputeTriageResultModel() ExternalConflictTriageResult {
	return ExternalConflictTriageResult{
		TriageResultID:           "triage_result_point14_valb_001",
		ConflictSetRef:           "conflict_set_point14_valb_001",
		StakeholderComparisonRef: "comparison_point14_valb_001",
		DisputeIntakeRef:         "dispute_point14_valb_001",
		EvidenceRequirementRef:   "evidence_requirement_point14_valb_001",
		GovernanceEscalationRef:  "escalation_point14_valb_001",
		TenantPrivacyBoundaryRef: "privacy_boundary_point14_valb_001",
		AgentBoundaryRef:         "boundary_point14_valb_agent_001",
		TriageState:              point14ValBTriageNoConflictDetected,
	}
}

func EvaluatePoint14ValBExternalConflictTriageResultState(model ExternalConflictTriageResult) string {
	if !point14ValBTriageResultIDValid(model.TriageResultID) ||
		!point14ValBConflictSetIDValid(model.ConflictSetRef) ||
		!point14ValBComparisonIDValid(model.StakeholderComparisonRef) ||
		!point14Val0DisputeIDValid(model.DisputeIntakeRef) ||
		!point14ValBEvidenceRequirementIDValid(model.EvidenceRequirementRef) ||
		!point14ValBEscalationIDValid(model.GovernanceEscalationRef) ||
		!point14Val0BoundaryRefValid(model.TenantPrivacyBoundaryRef) ||
		!point14Val0BoundaryRefValid(model.AgentBoundaryRef) ||
		!point14ValBStateValid(model.ConflictSetState) ||
		!point14ValBStateValid(model.StakeholderComparisonState) ||
		!point14ValBStateValid(model.DisputeIntakeState) ||
		!point14ValBStateValid(model.EvidenceRequirementState) ||
		!point14ValBStateValid(model.GovernanceEscalationState) ||
		!point14ValBStateValid(model.TenantPrivacyBoundaryState) ||
		!point14ValBStateValid(model.AgentBoundaryState) ||
		!point14ValBStateValid(model.NoExternalAuthorityState) ||
		!point14ValBStateValid(model.NoOverclaimState) ||
		!point14ValBExactValueValid(model.TriageState, point14ValBTriageStates()) {
		return Point14ValBStateBlocked
	}
	if model.EmitsPass ||
		model.PublishesCorrection ||
		model.RevokesClaim ||
		model.OverridesCanonicalDecision ||
		model.ApprovesProduction ||
		model.CertifiesCompliance ||
		model.CreatesPublicBadge {
		return Point14ValBStateBlocked
	}
	componentAggregate := point14ValBFoundationState(
		model.ConflictSetState,
		model.StakeholderComparisonState,
		model.DisputeIntakeState,
		model.EvidenceRequirementState,
		model.GovernanceEscalationState,
		model.TenantPrivacyBoundaryState,
		model.AgentBoundaryState,
		model.NoExternalAuthorityState,
		model.NoOverclaimState,
	)
	if componentAggregate == Point14ValBStateBlocked {
		return Point14ValBStateBlocked
	}
	if componentAggregate == Point14ValBStateReviewRequired {
		if strings.TrimSpace(model.TriageState) == point14ValBTriageNoConflictDetected || strings.TrimSpace(model.TriageState) == point14ValBTriageEvidenceRequired {
			return Point14ValBStateBlocked
		}
		return Point14ValBStateReviewRequired
	}
	if componentAggregate == Point14ValBStateIncomplete {
		if strings.TrimSpace(model.TriageState) == point14ValBTriageNoConflictDetected {
			return Point14ValBStateBlocked
		}
		return Point14ValBStateIncomplete
	}
	switch strings.TrimSpace(model.TriageState) {
	case point14ValBTriageBlocked, point14ValBTriageUnsupported, point14ValBTriageCrossTenant, point14ValBTriagePrivacyBlocked:
		return Point14ValBStateBlocked
	case point14ValBTriageReviewRequired, point14ValBTriageGovernanceEscalated:
		return Point14ValBStateReviewRequired
	case point14ValBTriageEvidenceRequired:
		return Point14ValBStateIncomplete
	}
	return Point14ValBStateActive
}

func point14ValBDisputeIntakeModel(dependency Point14ValBDependencySnapshot) DisputeIntakePacket {
	return DisputeIntakePacket{
		DisputeID:             "dispute_point14_valb_001",
		OpenedByRole:          "security_reviewer",
		OpenedAt:              "2026-05-06T07:30:00Z",
		OpenedTimeSource:      point14Val0TimeSourceServerUTC,
		TenantScope:           dependency.InheritedTenantScope,
		AffectedSignalRefs:    []string{"signal_point14_valb_scanner_001", "signal_point14_valb_vex_issuer_001"},
		AffectedArtifactRefs:  []string{"artifact_point14_valb_component_001"},
		ArtifactScoped:        true,
		AffectedClaimRefs:     []string{"claim_point14_valb_001"},
		ClaimScoped:           true,
		AffectedEvidenceRefs:  []string{"evidence_point14_valb_001"},
		DisputeReason:         "external signal conflict detected",
		LifecycleState:        point14Val0DisputeTriaged,
		SourceEventAt:         "2026-05-06T07:20:00Z",
		SourceEventTimeSource: point14Val0TimeSourceApprovedCustomerTime,
	}
}

func EvaluatePoint14ValBDisputeIntakePacketState(model DisputeIntakePacket, dependency Point14ValBDependencySnapshot) string {
	if !point14Val0DisputeIDValid(model.DisputeID) ||
		!point14Val0ExactValueValid(model.OpenedByRole, point14Val0RoleTypes()) ||
		!point14Val0ParsedTimeOk(model.OpenedAt) ||
		!point14Val0CanonicalTimeSourceValid(model.OpenedTimeSource) ||
		!point14Val0SignalRefsValid(model.AffectedSignalRefs) ||
		!point14ValBExactValueValid(model.LifecycleState, point14ValBDisputeLifecycleStates()) ||
		strings.TrimSpace(model.DisputeReason) == "" {
		return Point14ValBStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) == "" && strings.TrimSpace(model.GlobalScopeClassification) == "" {
		return Point14ValBStateBlocked
	}
	if model.TenantScope != "" {
		if !point11Val0ScopeValid(model.TenantScope) || model.TenantScope != dependency.InheritedTenantScope {
			return Point14ValBStateBlocked
		}
		if strings.TrimSpace(model.GlobalScopeClassification) != "" {
			return Point14ValBStateBlocked
		}
	} else if !point11Val0ContainsTrimmed([]string{point14Val0ScopeGlobalAdvisory, point14Val0ScopePublicNonAuthorative}, model.GlobalScopeClassification) {
		return Point14ValBStateBlocked
	}
	if model.ArtifactScoped && !point14Val0RefListValid(model.AffectedArtifactRefs, "artifact_") {
		return Point14ValBStateBlocked
	}
	if model.ClaimScoped && !point14Val0ClaimRefsValid(model.AffectedClaimRefs) {
		return Point14ValBStateBlocked
	}
	if !point14Val0EvidenceRefsValid(model.AffectedEvidenceRefs) || model.CanonicalMutationRequested {
		return Point14ValBStateBlocked
	}
	if strings.TrimSpace(model.SourceEventAt) != "" {
		if !point14Val0ParsedTimeOk(model.SourceEventAt) || !point14Val0TimeSourceValid(model.SourceEventTimeSource) {
			return Point14ValBStateBlocked
		}
		openedAt, _ := point14Val0ParsedTime(model.OpenedAt)
		sourceEventAt, _ := point14Val0ParsedTime(model.SourceEventAt)
		if sourceEventAt.After(openedAt) {
			return Point14ValBStateReviewRequired
		}
	}
	if model.ClosureAttempted {
		if !point14Val0ParsedTimeOk(model.ClosedAt) {
			return Point14ValBStateBlocked
		}
		openedAt, _ := point14Val0ParsedTime(model.OpenedAt)
		closedAt, _ := point14Val0ParsedTime(model.ClosedAt)
		if closedAt.Before(openedAt) {
			return Point14ValBStateBlocked
		}
		return Point14ValBStateBlocked
	}
	switch strings.TrimSpace(model.LifecycleState) {
	case point14Val0DisputeOpened, point14Val0DisputeTriaged, point14Val0DisputeRejected:
		return Point14ValBStateActive
	case point14Val0DisputeEvidenceNeeded:
		return Point14ValBStateIncomplete
	case point14Val0DisputeReviewRequired:
		return Point14ValBStateReviewRequired
	default:
		return Point14ValBStateBlocked
	}
}

func point14ValBEvidenceRequirementGateModel() DisputeEvidenceRequirementGate {
	return DisputeEvidenceRequirementGate{
		EvidenceRequirementID: "evidence_requirement_point14_valb_001",
		RequiredEvidenceRefs:  []string{"evidence_point14_valb_001"},
		RequiredEvidenceTypes: []string{"canonical_evidence", "hash_match"},
		EvidenceState:         point14ValAEvidenceStateActive,
	}
}

func EvaluatePoint14ValBDisputeEvidenceRequirementGateState(model DisputeEvidenceRequirementGate) string {
	if !point14ValBEvidenceRequirementIDValid(model.EvidenceRequirementID) ||
		!point14Val0EvidenceRefsValid(model.RequiredEvidenceRefs) ||
		!point14ValBEvidenceTypesValid(model.RequiredEvidenceTypes) ||
		!point14ValAEvidenceStateValid(model.EvidenceState) {
		return Point14ValBStateBlocked
	}
	if model.ConflictTriageBypassRequested ||
		model.AgentRecommendationOnly ||
		model.CrowdConsensusOnly ||
		model.AuditorNoteOnly {
		return Point14ValBStateBlocked
	}
	switch strings.TrimSpace(model.EvidenceState) {
	case point14ValAEvidenceStateRevoked, point14ValAEvidenceStateExpired, point14ValAEvidenceStateUnrelated:
		return Point14ValBStateBlocked
	case point14ValAEvidenceStateStale, point14ValAEvidenceStateSuperseded:
		return Point14ValBStateReviewRequired
	}
	if model.DecisiveEvidenceMissing {
		return Point14ValBStateIncomplete
	}
	return Point14ValBStateActive
}

func point14ValBGovernanceEscalationPathModel() GovernanceEscalationPath {
	return GovernanceEscalationPath{
		EscalationID:    "escalation_point14_valb_001",
		EscalationState: point14ValBEscalationStateNotRequired,
	}
}

func EvaluatePoint14ValBGovernanceEscalationPathState(model GovernanceEscalationPath) string {
	if !point14ValBEscalationIDValid(model.EscalationID) ||
		!point14ValBExactValueValid(model.EscalationState, point14ValBEscalationStates()) {
		return Point14ValBStateBlocked
	}
	if model.ApprovesProduction || model.PublishesCorrection {
		return Point14ValBStateBlocked
	}
	if !model.EscalationRequired {
		if strings.TrimSpace(model.EscalationState) != point14ValBEscalationStateNotRequired ||
			strings.TrimSpace(model.GovernanceEventRef) != "" ||
			strings.TrimSpace(model.Owner) != "" ||
			strings.TrimSpace(model.ApproverRole) != "" ||
			strings.TrimSpace(model.Reason) != "" ||
			strings.TrimSpace(model.AuditRef) != "" {
			return Point14ValBStateBlocked
		}
		return Point14ValBStateActive
	}
	if !point14Val0GovernanceEventRefValid(model.GovernanceEventRef) ||
		!point14Val0ApproverRefValid(model.Owner) ||
		!point14Val0ExactValueValid(model.ApproverRole, point14Val0RoleTypes()) ||
		strings.TrimSpace(model.Reason) == "" ||
		!point14Val0AuditEventRefValid(model.AuditRef) {
		return Point14ValBStateBlocked
	}
	if strings.TrimSpace(model.EscalationState) == point14ValBEscalationStateCompleted {
		return Point14ValBStateActive
	}
	return Point14ValBStateReviewRequired
}

func point14ValBTenantPrivacyConflictBoundaryModel(dependency Point14ValBDependencySnapshot) TenantPrivacyConflictBoundary {
	return TenantPrivacyConflictBoundary{
		BoundaryID:        "privacy_boundary_point14_valb_001",
		TenantScope:       dependency.InheritedTenantScope,
		VisibilityScope:   point14Val0VisibilityTenantPrivate,
		RedactionRefs:     []string{"redaction_ref_point14_valb_001"},
		LimitationRefs:    []string{"limitation_ref_point14_valb_001"},
		LimitationVisible: true,
		ObservedTexts:     []string{"bounded dispute triage state"},
	}
}

func EvaluatePoint14ValBTenantPrivacyConflictBoundaryState(model TenantPrivacyConflictBoundary, dependency Point14ValBDependencySnapshot) string {
	if !point14Val0BoundaryRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		model.TenantScope != dependency.InheritedTenantScope ||
		!point14Val0ExactValueValid(model.VisibilityScope, point14Val0VisibilityScopes()) ||
		!point14ValBRedactionRefsValid(model.RedactionRefs) ||
		!point14ValBLimitationRefsValid(model.LimitationRefs) ||
		!point14Val0TextListValid(model.ObservedTexts) {
		return Point14ValBStateBlocked
	}
	if model.TenantPrivateDataExposed || model.StrengthensClaim || !model.LimitationVisible {
		return Point14ValBStateBlocked
	}
	if model.PublicVisibilityRequested {
		if strings.TrimSpace(model.VisibilityScope) != point14Val0VisibilityPublicNoticeLimited {
			return Point14ValBStateBlocked
		}
		return Point14ValBStateReviewRequired
	}
	return Point14ValBStateActive
}

func point14ValBAgentDisputeRecommendationBoundaryModel(dependency Point14ValBDependencySnapshot) AgentDisputeRecommendationBoundary {
	return AgentDisputeRecommendationBoundary{
		BoundaryID:     "boundary_point14_valb_agent_001",
		TenantScope:    dependency.InheritedTenantScope,
		AgentInputRefs: []string{"agent_input_point14_valb_001"},
		EvidenceRefs:   []string{"evidence_point14_valb_001"},
		AuditEventRef:  "audit_event_point14_valb_agent_001",
		AdvisoryOnly:   true,
	}
}

func EvaluatePoint14ValBAgentDisputeRecommendationBoundaryState(model AgentDisputeRecommendationBoundary, dependency Point14ValBDependencySnapshot) string {
	if !point14Val0BoundaryRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		model.TenantScope != dependency.InheritedTenantScope ||
		!point14Val0AgentInputRefsValid(model.AgentInputRefs) ||
		!point14Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point14Val0AuditEventRefValid(model.AuditEventRef) {
		return Point14ValBStateBlocked
	}
	if !model.AdvisoryOnly ||
		model.CanResolveConflict ||
		model.CanPublishCorrection ||
		model.CanRevokeClaim ||
		model.CanOverrideGovernance ||
		model.CanSatisfyEvidenceRequirementAlone ||
		model.CanEmitPass ||
		model.CanEmitPublicAuthority ||
		model.PassAllowed ||
		model.ApprovalGranted ||
		model.ProductionApproved ||
		model.ExternalAuthorityAllowed {
		return Point14ValBStateBlocked
	}
	return Point14ValBStateActive
}

func point14ValBNoExternalAuthorityConflictGuardModel() Point14ValBNoExternalAuthorityConflictGuard {
	return Point14ValBNoExternalAuthorityConflictGuard{}
}

func EvaluatePoint14ValBNoExternalAuthorityConflictGuardState(model Point14ValBNoExternalAuthorityConflictGuard) string {
	for _, marker := range model.ObservedAuthorityMarkers {
		if point11Val0ContainsTrimmed(point14ValBForbiddenAuthorityMarkers(), marker) {
			return Point14ValBStateBlocked
		}
	}
	if model.CanonicalAuthorityGranted ||
		model.ProductionApprovalGranted ||
		model.PublishesCorrection ||
		model.RevokesClaim ||
		model.PublicBadgeAllowed ||
		model.DisputeAutoResolved ||
		model.CrowdResolved ||
		model.AgentResolved ||
		model.ExternalAuthorityAllowed {
		return Point14ValBStateBlocked
	}
	return Point14ValBStateActive
}

func point14ValBNoOverclaimDisputeWordingModel() Point14ValBNoOverclaimDisputeWording {
	return Point14ValBNoOverclaimDisputeWording{
		ObservedConflictTexts:            []string{"external signal conflict detected"},
		ObservedDisputeTexts:             []string{"dispute requires governance review"},
		ObservedEvidenceRequirementTexts: []string{"evidence required before triage closure"},
		ObservedEscalationTexts:          []string{"advisory conflict input"},
		ObservedPrivacyTexts:             []string{"bounded dispute triage state"},
		ObservedAgentTexts:               []string{"no external authority granted", "no automatic correction publication", "canonical evidence spine remains source of truth"},
		AllowedSafeWording:               point14ValBSafeWording(),
		BlockedWording:                   point14ValBForbiddenWording(),
		ProjectionDisclaimer:             point14ValBProjectionDisclaimerBase,
	}
}

func EvaluatePoint14ValBNoOverclaimDisputeWordingState(model Point14ValBNoOverclaimDisputeWording) string {
	if model.ProjectionDisclaimer != point14ValBProjectionDisclaimerBase ||
		!point14Val0TextListValid(model.AllowedSafeWording) ||
		!point14Val0TextListValid(model.BlockedWording) {
		return Point14ValBStateBlocked
	}
	if point14ValBObservedListContainsForbiddenWording(model.ObservedConflictTexts) ||
		point14ValBObservedListContainsForbiddenWording(model.ObservedDisputeTexts) ||
		point14ValBObservedListContainsForbiddenWording(model.ObservedEvidenceRequirementTexts) ||
		point14ValBObservedListContainsForbiddenWording(model.ObservedEscalationTexts) ||
		point14ValBObservedListContainsForbiddenWording(model.ObservedPrivacyTexts) ||
		point14ValBObservedListContainsForbiddenWording(model.ObservedAgentTexts) {
		return Point14ValBStateBlocked
	}
	if point14ValBObservedListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point14ValBStateBlocked
	}
	return Point14ValBStateActive
}

func point14ValBFoundationState(states ...string) string {
	hasReview := false
	hasIncomplete := false
	for _, state := range states {
		switch strings.TrimSpace(state) {
		case Point14ValBStateBlocked:
			return Point14ValBStateBlocked
		case Point14ValBStateReviewRequired:
			hasReview = true
		case Point14ValBStateIncomplete:
			hasIncomplete = true
		}
	}
	if hasReview {
		return Point14ValBStateReviewRequired
	}
	if hasIncomplete {
		return Point14ValBStateIncomplete
	}
	return Point14ValBStateActive
}

func point14ValBBlockingReasons(model Point14ValBFoundation) []string {
	componentStates := map[string]string{
		"dependency":              model.DependencyState,
		"conflict_set":            model.ConflictSetState,
		"stakeholder_comparison":  model.StakeholderComparisonState,
		"dispute_triage_result":   model.DisputeTriageResultState,
		"dispute_intake":          model.DisputeIntakeState,
		"evidence_requirement":    model.EvidenceRequirementState,
		"governance_escalation":   model.GovernanceEscalationState,
		"tenant_privacy_boundary": model.TenantPrivacyBoundaryState,
		"agent_dispute_boundary":  model.AgentDisputeBoundaryState,
		"no_external_authority":   model.NoExternalAuthorityState,
		"no_overclaim":            model.NoOverclaimState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if state == Point14ValBStateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point14ValBReviewPrerequisites(model Point14ValBFoundation) []string {
	componentStates := map[string]string{
		"dependency":              model.DependencyState,
		"conflict_set":            model.ConflictSetState,
		"stakeholder_comparison":  model.StakeholderComparisonState,
		"dispute_triage_result":   model.DisputeTriageResultState,
		"dispute_intake":          model.DisputeIntakeState,
		"evidence_requirement":    model.EvidenceRequirementState,
		"governance_escalation":   model.GovernanceEscalationState,
		"tenant_privacy_boundary": model.TenantPrivacyBoundaryState,
		"agent_dispute_boundary":  model.AgentDisputeBoundaryState,
		"no_external_authority":   model.NoExternalAuthorityState,
		"no_overclaim":            model.NoOverclaimState,
	}
	prereqs := []string{}
	for name, state := range componentStates {
		if state == Point14ValBStateReviewRequired || state == Point14ValBStateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func point14ValBFoundationModelFromUpstream(valA Point14ValAFoundation) Point14ValBFoundation {
	dependency := point14ValBDependencySnapshotFromUpstream(valA)
	return Point14ValBFoundation{
		CurrentState:                       Point14ValBStateActive,
		ProjectionDisclaimer:               point14ValBProjectionDisclaimerBase,
		DependencyState:                    Point14ValBStateActive,
		ConflictSetState:                   Point14ValBStateActive,
		StakeholderComparisonState:         Point14ValBStateActive,
		DisputeTriageResultState:           Point14ValBStateActive,
		DisputeIntakeState:                 Point14ValBStateActive,
		EvidenceRequirementState:           Point14ValBStateActive,
		GovernanceEscalationState:          Point14ValBStateActive,
		TenantPrivacyBoundaryState:         Point14ValBStateActive,
		AgentDisputeBoundaryState:          Point14ValBStateActive,
		NoExternalAuthorityState:           Point14ValBStateActive,
		NoOverclaimState:                   Point14ValBStateActive,
		Dependency:                         dependency,
		ConflictSet:                        point14ValBConflictSetModel(dependency),
		StakeholderComparison:              point14ValBStakeholderComparisonModel(),
		DisputeTriageResult:                point14ValBDisputeTriageResultModel(),
		DisputeIntake:                      point14ValBDisputeIntakeModel(dependency),
		EvidenceRequirementGate:            point14ValBEvidenceRequirementGateModel(),
		GovernanceEscalationPath:           point14ValBGovernanceEscalationPathModel(),
		TenantPrivacyConflictBoundary:      point14ValBTenantPrivacyConflictBoundaryModel(dependency),
		AgentDisputeRecommendationBoundary: point14ValBAgentDisputeRecommendationBoundaryModel(dependency),
		NoExternalAuthorityConflictGuard:   point14ValBNoExternalAuthorityConflictGuardModel(),
		NoOverclaimDisputeWording:          point14ValBNoOverclaimDisputeWordingModel(),
	}
}

func Point14ValBFoundationModel() Point14ValBFoundation {
	valA := ComputePoint14ValAFoundation(Point14ValAFoundationModel())
	return point14ValBFoundationModelFromUpstream(valA)
}

func ComputePoint14ValBFoundation(model Point14ValBFoundation) Point14ValBFoundation {
	model.DependencyState = EvaluatePoint14ValBDependencyState(model.Dependency)
	model.ConflictSetState = EvaluatePoint14ValBExternalSignalConflictSetState(model.ConflictSet, model.Dependency)
	model.StakeholderComparisonState = EvaluatePoint14ValBStakeholderSignalComparisonState(model.StakeholderComparison)
	model.DisputeIntakeState = EvaluatePoint14ValBDisputeIntakePacketState(model.DisputeIntake, model.Dependency)
	model.EvidenceRequirementState = EvaluatePoint14ValBDisputeEvidenceRequirementGateState(model.EvidenceRequirementGate)
	model.GovernanceEscalationState = EvaluatePoint14ValBGovernanceEscalationPathState(model.GovernanceEscalationPath)
	model.TenantPrivacyBoundaryState = EvaluatePoint14ValBTenantPrivacyConflictBoundaryState(model.TenantPrivacyConflictBoundary, model.Dependency)
	model.AgentDisputeBoundaryState = EvaluatePoint14ValBAgentDisputeRecommendationBoundaryState(model.AgentDisputeRecommendationBoundary, model.Dependency)
	model.NoExternalAuthorityState = EvaluatePoint14ValBNoExternalAuthorityConflictGuardState(model.NoExternalAuthorityConflictGuard)
	model.NoOverclaimState = EvaluatePoint14ValBNoOverclaimDisputeWordingState(model.NoOverclaimDisputeWording)

	model.DisputeTriageResult.ConflictSetState = model.ConflictSetState
	model.DisputeTriageResult.StakeholderComparisonState = model.StakeholderComparisonState
	model.DisputeTriageResult.DisputeIntakeState = model.DisputeIntakeState
	model.DisputeTriageResult.EvidenceRequirementState = model.EvidenceRequirementState
	model.DisputeTriageResult.GovernanceEscalationState = model.GovernanceEscalationState
	model.DisputeTriageResult.TenantPrivacyBoundaryState = model.TenantPrivacyBoundaryState
	model.DisputeTriageResult.AgentBoundaryState = model.AgentDisputeBoundaryState
	model.DisputeTriageResult.NoExternalAuthorityState = model.NoExternalAuthorityState
	model.DisputeTriageResult.NoOverclaimState = model.NoOverclaimState
	model.DisputeTriageResult.TriageState = point14ValBResolvedTriageState(
		model.DisputeTriageResult.TriageState,
		point14ValBFoundationState(
			model.ConflictSetState,
			model.StakeholderComparisonState,
			model.DisputeIntakeState,
			model.EvidenceRequirementState,
			model.GovernanceEscalationState,
			model.TenantPrivacyBoundaryState,
			model.AgentDisputeBoundaryState,
			model.NoExternalAuthorityState,
			model.NoOverclaimState,
		),
	)
	model.DisputeTriageResultState = EvaluatePoint14ValBExternalConflictTriageResultState(model.DisputeTriageResult)

	model.CurrentState = point14ValBFoundationState(
		model.DependencyState,
		model.ConflictSetState,
		model.StakeholderComparisonState,
		model.DisputeTriageResultState,
		model.DisputeIntakeState,
		model.EvidenceRequirementState,
		model.GovernanceEscalationState,
		model.TenantPrivacyBoundaryState,
		model.AgentDisputeBoundaryState,
		model.NoExternalAuthorityState,
		model.NoOverclaimState,
	)
	model.BlockingReasons = point14ValBBlockingReasons(model)
	model.ReviewPrerequisites = point14ValBReviewPrerequisites(model)
	return model
}

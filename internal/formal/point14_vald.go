package formal

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/denisgrosek/changelock/internal/operability"
)

const (
	Point14ValDStateActive         = "point14_vald_ecosystem_timeline_projection_active"
	Point14ValDStateBlocked        = "point14_vald_ecosystem_timeline_projection_blocked"
	Point14ValDStateReviewRequired = "point14_vald_ecosystem_timeline_projection_review_required"
	Point14ValDStateIncomplete     = "point14_vald_ecosystem_timeline_projection_incomplete"
)

const (
	point14ValDWaveID                     = "val_d"
	point14ValDProjectionDisclaimerBase   = "projection_only not_canonical_truth point14_vald_ecosystem_timeline_projection"
	point14ValDTimelineScopeEcosystem     = "ecosystem_timeline"
	point14ValDTimelineScopeSignals       = "signal_timeline"
	point14ValDTimelineScopeDisputes      = "dispute_timeline"
	point14ValDTimelineScopePublication   = "publication_history"
	point14ValDTimelineScopeGovernance    = "governance_history"
	point14ValDTimelineScopeAgentAdvisory = "agent_advisory_history"
	point14ValDFilterSignalRef            = "signal_ref"
	point14ValDFilterDisputeRef           = "dispute_ref"
	point14ValDFilterPublicationState     = "publication_state"
	point14ValDFilterTimelineState        = "timeline_state"
	point14ValDFilterStakeholderRole      = "stakeholder_role"
	point14ValDFilterVisibilityScope      = "visibility_scope"
	point14ValDFilterGovernanceTrace      = "governance_trace"
	point14ValDFilterTimeWindow           = "time_window"
	point14ValDFilterLimitationVisible    = "limitation_visible"
	point14ValDSignalEntryVisible         = "entry_visible"
	point14ValDSignalEntryReviewRequired  = "entry_review_required"
	point14ValDSignalEntryIncomplete      = "entry_incomplete"
	point14ValDSignalEntryBlocked         = "entry_blocked"
	point14ValDDisputeTimelineVisible     = "timeline_visible"
	point14ValDDisputeTimelineReview      = "timeline_review_required"
	point14ValDDisputeTimelineIncomplete  = "timeline_incomplete"
	point14ValDDisputeTimelineBlocked     = "timeline_blocked"
	point14ValDViewerPublic               = "public_viewer"
)

type Point14ValDDependencySnapshot struct {
	Point14ValCCurrentState                  string                                          `json:"point14_valc_current_state"`
	Point14ValCDependencyState               string                                          `json:"point14_valc_dependency_state"`
	Point14ValCCorrectionNoticeState         string                                          `json:"point14_valc_correction_notice_state"`
	Point14ValCRevocationRequestState        string                                          `json:"point14_valc_revocation_request_state"`
	Point14ValCSupersessionRecordState       string                                          `json:"point14_valc_supersession_record_state"`
	Point14ValCPublicationApprovalState      string                                          `json:"point14_valc_publication_approval_state"`
	Point14ValCVisibilityBoundaryState       string                                          `json:"point14_valc_visibility_boundary_state"`
	Point14ValCTenantPrivacyState            string                                          `json:"point14_valc_tenant_privacy_state"`
	Point14ValCRedactionLimitationState      string                                          `json:"point14_valc_redaction_limitation_state"`
	Point14ValCGovernanceTraceState          string                                          `json:"point14_valc_governance_trace_state"`
	Point14ValCAgentPublicationBoundaryState string                                          `json:"point14_valc_agent_publication_boundary_state"`
	Point14ValCNoExternalAuthorityState      string                                          `json:"point14_valc_no_external_authority_state"`
	Point14ValCNoOverclaimState              string                                          `json:"point14_valc_no_overclaim_state"`
	Point14ValCPointID                       string                                          `json:"point14_valc_point_id"`
	Point14ValCWaveID                        string                                          `json:"point14_valc_wave_id"`
	Point14ValCComputedFromUpstream          bool                                            `json:"point14_valc_computed_from_upstream"`
	Point14ValCMerged                        bool                                            `json:"point14_valc_merged"`
	Point14ValCCIGreen                       bool                                            `json:"point14_valc_ci_green"`
	Point14ValCReviewedOnMain                bool                                            `json:"point14_valc_reviewed_on_main"`
	Point14PassSeen                          bool                                            `json:"point14_pass_seen"`
	InheritedPoint14ValBCurrentState         string                                          `json:"inherited_point14_valb_current_state"`
	InheritedPoint14ValACurrentState         string                                          `json:"inherited_point14_vala_current_state"`
	InheritedPoint14Val0CurrentState         string                                          `json:"inherited_point14_val0_current_state"`
	InheritedPoint13ValECurrentState         string                                          `json:"inherited_point13_vale_current_state"`
	InheritedPoint13ValEPassClosureState     string                                          `json:"inherited_point13_vale_pass_closure_state"`
	InheritedPoint13ValEPassAllowed          bool                                            `json:"inherited_point13_vale_pass_allowed"`
	InheritedPoint13ValEPassToken            string                                          `json:"inherited_point13_vale_pass_token"`
	InheritedPoint12CurrentState             string                                          `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState          string                                          `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureState         string                                          `json:"inherited_point12_pass_closure_state"`
	InheritedPoint12ReviewerResult           string                                          `json:"inherited_point12_reviewer_result"`
	InheritedPoint11PublicationState         string                                          `json:"inherited_point11_publication_state"`
	InheritedPoint11NoOverclaimState         string                                          `json:"inherited_point11_no_overclaim_state"`
	InheritedPoint11FinalPassGateState       string                                          `json:"inherited_point11_final_pass_gate_state"`
	InheritedPoint10CurrentState             string                                          `json:"inherited_point10_current_state"`
	InheritedPoint10NoOverclaimState         string                                          `json:"inherited_point10_no_overclaim_state"`
	InheritedPoint10ProjectionState          string                                          `json:"inherited_point10_projection_state"`
	InheritedPoint10PassRuleState            string                                          `json:"inherited_point10_pass_rule_state"`
	InheritedTenantScope                     string                                          `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput               bool                                            `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                      []string                                        `json:"review_prerequisites,omitempty"`
	Point14ValC                              Point14ValCFoundation                           `json:"point14_valc"`
	Point14ValB                              Point14ValBFoundation                           `json:"point14_valb"`
	Point14ValA                              Point14ValAFoundation                           `json:"point14_vala"`
	Point14Val0                              Point14Val0Foundation                           `json:"point14_val0"`
	Point13ValE                              Point13ValEFoundation                           `json:"point13_vale"`
	Point12                                  Point12ValEFoundation                           `json:"point12"`
	Point11                                  Point11ValDFoundation                           `json:"point11"`
	Point10                                  operability.DeploymentMultiTenantValEFoundation `json:"point10"`
}

type ExternalEcosystemTimelineProjection struct {
	TimelineProjectionID           string   `json:"timeline_projection_id"`
	TenantScope                    string   `json:"tenant_scope"`
	GlobalScopeClassification      string   `json:"global_scope_classification"`
	SourceProjectionRefs           []string `json:"source_projection_refs,omitempty"`
	SignalEntryRefs                []string `json:"signal_entry_refs,omitempty"`
	DisputeEntryRefs               []string `json:"dispute_entry_refs,omitempty"`
	CorrectionPublicationEntryRefs []string `json:"correction_publication_entry_refs,omitempty"`
	GovernanceTraceRefs            []string `json:"governance_trace_refs,omitempty"`
	GeneratedAt                    string   `json:"generated_at"`
	GeneratedTimeSource            string   `json:"generated_time_source"`
	ReadOnly                       bool     `json:"read_only"`
	ProjectionOnly                 bool     `json:"projection_only"`
	MutatesCanonicalEvidence       bool     `json:"mutates_canonical_evidence"`
	MutatesSignalState             bool     `json:"mutates_signal_state"`
	ResolvesDispute                bool     `json:"resolves_dispute"`
	PublishesCorrection            bool     `json:"publishes_correction"`
	RevokesClaim                   bool     `json:"revokes_claim"`
	ApprovesProduction             bool     `json:"approves_production"`
	CertifiesCompliance            bool     `json:"certifies_compliance"`
	CreatesPublicBadge             bool     `json:"creates_public_badge"`
	EmitsPass                      bool     `json:"emits_pass"`
}

type ExternalSignalTimelineEntryProjection struct {
	TimelineEntryID                 string `json:"timeline_entry_id"`
	NormalizedSignalRef             string `json:"normalized_signal_ref"`
	ValidationResultRef             string `json:"validation_result_ref"`
	SourceIdentityRef               string `json:"source_identity_ref"`
	StakeholderRoleRef              string `json:"stakeholder_role_ref"`
	EventAt                         string `json:"event_at"`
	EventTimeSource                 string `json:"event_time_source"`
	ReceivedAt                      string `json:"received_at"`
	ReceivedTimeSource              string `json:"received_time_source"`
	TimelineState                   string `json:"timeline_state"`
	AdvisoryOnly                    bool   `json:"advisory_only"`
	AuthorityGranted                bool   `json:"authority_granted"`
	UpgradesSignalValidity          bool   `json:"upgrades_signal_validity"`
	SourceEventAsCanonicalAuthority bool   `json:"source_event_as_canonical_authority"`
}

type ExternalDisputeTimelineProjection struct {
	DisputeTimelineID                string   `json:"dispute_timeline_id"`
	DisputeRef                       string   `json:"dispute_ref"`
	ConflictSetRef                   string   `json:"conflict_set_ref"`
	TriageResultRef                  string   `json:"triage_result_ref"`
	LifecycleState                   string   `json:"lifecycle_state"`
	EvidenceRequirementRefs          []string `json:"evidence_requirement_refs,omitempty"`
	GovernanceEscalationRefs         []string `json:"governance_escalation_refs,omitempty"`
	TimelineState                    string   `json:"timeline_state"`
	ResolvesDispute                  bool     `json:"resolves_dispute"`
	MovesLifecycleToCorrected        bool     `json:"moves_lifecycle_to_corrected"`
	MovesLifecycleToRevoked          bool     `json:"moves_lifecycle_to_revoked"`
	MovesLifecycleToPublishedNotice  bool     `json:"moves_lifecycle_to_published_notice"`
	HidesEvidenceRequired            bool     `json:"hides_evidence_required"`
	ConvertsReviewIncompleteToActive bool     `json:"converts_review_incomplete_to_active"`
}

type CorrectionRevocationPublicationReadProjection struct {
	ReadProjectionID               string   `json:"read_projection_id"`
	CorrectionNoticeRefs           []string `json:"correction_notice_refs,omitempty"`
	RevocationRequestRefs          []string `json:"revocation_request_refs,omitempty"`
	SupersessionRecordRefs         []string `json:"supersession_record_refs,omitempty"`
	PublicationApprovalRefs        []string `json:"publication_approval_refs,omitempty"`
	VisibilityBoundaryRefs         []string `json:"visibility_boundary_refs,omitempty"`
	LimitationRefs                 []string `json:"limitation_refs,omitempty"`
	RedactionRefs                  []string `json:"redaction_refs,omitempty"`
	PublicationStateRefs           []string `json:"publication_state_refs,omitempty"`
	ReadOnly                       bool     `json:"read_only"`
	PublishesCorrection            bool     `json:"publishes_correction"`
	ExecutesRevocation             bool     `json:"executes_revocation"`
	SilentReplacesSupersededSignal bool     `json:"silent_replaces_superseded_signal"`
	OmitsLimitations               bool     `json:"omits_limitations"`
	HidesRedaction                 bool     `json:"hides_redaction"`
	ObservedReadTexts              []string `json:"observed_read_texts,omitempty"`
}

type GovernanceTraceReadProjection struct {
	GovernanceTraceProjectionID         string   `json:"governance_trace_projection_id"`
	GovernanceTraceRefs                 []string `json:"governance_trace_refs,omitempty"`
	OwnerRefs                           []string `json:"owner_refs,omitempty"`
	ApproverRoleRefs                    []string `json:"approver_role_refs,omitempty"`
	AuditRefs                           []string `json:"audit_refs,omitempty"`
	EvidenceRefs                        []string `json:"evidence_refs,omitempty"`
	DecisionReasonRefs                  []string `json:"decision_reason_refs,omitempty"`
	TimestampRefs                       []string `json:"timestamp_refs,omitempty"`
	ReadOnly                            bool     `json:"read_only"`
	ApprovesAnything                    bool     `json:"approves_anything"`
	MutatesGovernanceTrace              bool     `json:"mutates_governance_trace"`
	SatisfiesMissingGovernanceByDisplay bool     `json:"satisfies_missing_governance_by_display"`
}

type EcosystemTimelineQueryProjection struct {
	QueryProjectionID                 string   `json:"query_projection_id"`
	QueryScope                        string   `json:"query_scope"`
	AllowedFilters                    []string `json:"allowed_filters,omitempty"`
	ResultRefs                        []string `json:"result_refs,omitempty"`
	ReadOnly                          bool     `json:"read_only"`
	QueryIsProjectionOnly             bool     `json:"query_is_projection_only"`
	MutationRequested                 bool     `json:"mutation_requested"`
	WritesFiltersBack                 bool     `json:"writes_filters_back"`
	HidesDecisiveMissingEvidence      bool     `json:"hides_decisive_missing_evidence"`
	OmitsLimitationsWithoutDisclosure bool     `json:"omits_limitations_without_disclosure"`
	CrossTenantResults                bool     `json:"cross_tenant_results"`
}

type EcosystemTimelineAccessBoundary struct {
	AccessBoundaryID          string `json:"access_boundary_id"`
	ViewerRole                string `json:"viewer_role"`
	TenantScope               string `json:"tenant_scope"`
	GlobalScopeClassification string `json:"global_scope_classification"`
	AllowedViewScope          string `json:"allowed_view_scope"`
	AuditRef                  string `json:"audit_ref"`
	AccessTime                string `json:"access_time"`
	AccessTimeSource          string `json:"access_time_source"`
	AccessExpired             bool   `json:"access_expired"`
	AccessRevoked             bool   `json:"access_revoked"`
	CrossTenantAccess         bool   `json:"cross_tenant_access"`
	AuthorityGranted          bool   `json:"authority_granted"`
	TenantPrivateDataExposed  bool   `json:"tenant_private_data_exposed"`
}

type TenantPrivacyTimelineProjectionGuard struct {
	BoundaryID                  string   `json:"boundary_id"`
	TenantScope                 string   `json:"tenant_scope"`
	ProjectionTargetScope       string   `json:"projection_target_scope"`
	PublicPrivateClassification string   `json:"public_private_classification"`
	RedactionRefs               []string `json:"redaction_refs,omitempty"`
	LimitationRefs              []string `json:"limitation_refs,omitempty"`
	TenantPrivateDataExposed    bool     `json:"tenant_private_data_exposed"`
	CrossTenantProjection       bool     `json:"cross_tenant_projection"`
	ObservedSummaryTexts        []string `json:"observed_summary_texts,omitempty"`
	LimitationsVisible          bool     `json:"limitations_visible"`
	StrengthensClaim            bool     `json:"strengthens_claim"`
}

type AgentEcosystemTimelineProjection struct {
	ProjectionID              string   `json:"projection_id"`
	TenantScope               string   `json:"tenant_scope"`
	AgentInputRefs            []string `json:"agent_input_refs,omitempty"`
	AgentRecommendationRefs   []string `json:"agent_recommendation_refs,omitempty"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	AuditEventRef             string   `json:"audit_event_ref"`
	AdvisoryOnly              bool     `json:"advisory_only"`
	AgentApprovalFlags        bool     `json:"agent_approval_flags"`
	AgentAuthorityFlags       bool     `json:"agent_authority_flags"`
	CanResolveDispute         bool     `json:"can_resolve_dispute"`
	CanPublishCorrection      bool     `json:"can_publish_correction"`
	CanRevokeClaim            bool     `json:"can_revoke_claim"`
	CanSatisfyGovernanceTrace bool     `json:"can_satisfy_governance_trace"`
	PassAllowed               bool     `json:"pass_allowed"`
	ExternalAuthorityAllowed  bool     `json:"external_authority_allowed"`
}

type TimelineTimestampIntegrityGuard struct {
	GuardID                       string   `json:"guard_id"`
	TimelineEventRefs             []string `json:"timeline_event_refs,omitempty"`
	EventAt                       string   `json:"event_at"`
	EventTimeSource               string   `json:"event_time_source"`
	ReceivedAt                    string   `json:"received_at"`
	ReceivedTimeSource            string   `json:"received_time_source"`
	GeneratedAt                   string   `json:"generated_at"`
	GeneratedTimeSource           string   `json:"generated_time_source"`
	AccessTime                    string   `json:"access_time"`
	AccessTimeSource              string   `json:"access_time_source"`
	SourceEventAt                 string   `json:"source_event_at"`
	SourceEventTimeSource         string   `json:"source_event_time_source"`
	PublicationApprovedAt         string   `json:"publication_approved_at"`
	PublicationApprovedTimeSource string   `json:"publication_approved_time_source"`
	DisputeOpenedAt               string   `json:"dispute_opened_at"`
	DisputeOpenedTimeSource       string   `json:"dispute_opened_time_source"`
	AttemptsValidityUpgrade       bool     `json:"attempts_validity_upgrade"`
}

type Point14ValDNoMutationProjectionGuard struct {
	BoundaryID                 string `json:"boundary_id"`
	MutatesCanonicalEvidence   bool   `json:"mutates_canonical_evidence"`
	MutatesNormalizedSignal    bool   `json:"mutates_normalized_signal"`
	MutatesValidationResult    bool   `json:"mutates_validation_result"`
	MutatesDisputeLifecycle    bool   `json:"mutates_dispute_lifecycle"`
	MutatesCorrectionNotice    bool   `json:"mutates_correction_notice"`
	MutatesRevocationRequest   bool   `json:"mutates_revocation_request"`
	MutatesSupersessionRecord  bool   `json:"mutates_supersession_record"`
	MutatesPublicationApproval bool   `json:"mutates_publication_approval"`
	MutatesVisibilityBoundary  bool   `json:"mutates_visibility_boundary"`
	MutatesGovernanceTrace     bool   `json:"mutates_governance_trace"`
	PublishesCorrection        bool   `json:"publishes_correction"`
	ExecutesRevocation         bool   `json:"executes_revocation"`
	ResolvesDispute            bool   `json:"resolves_dispute"`
	ApprovesProduction         bool   `json:"approves_production"`
	CertifiesCompliance        bool   `json:"certifies_compliance"`
	CreatesPublicBadge         bool   `json:"creates_public_badge"`
	EmitsPass                  bool   `json:"emits_pass"`
}

type Point14ValDNoOverclaimTimelineWording struct {
	ObservedTimelineTexts                []string `json:"observed_timeline_texts,omitempty"`
	ObservedSignalTexts                  []string `json:"observed_signal_texts,omitempty"`
	ObservedDisputeTexts                 []string `json:"observed_dispute_texts,omitempty"`
	ObservedReadProjectionTexts          []string `json:"observed_read_projection_texts,omitempty"`
	ObservedGovernanceTexts              []string `json:"observed_governance_texts,omitempty"`
	ObservedQueryTexts                   []string `json:"observed_query_texts,omitempty"`
	ObservedAccessTexts                  []string `json:"observed_access_texts,omitempty"`
	ObservedPrivacyTexts                 []string `json:"observed_privacy_texts,omitempty"`
	ObservedAgentTexts                   []string `json:"observed_agent_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point14ValDFoundation struct {
	CurrentState                         string                                        `json:"current_state"`
	BlockingReasons                      []string                                      `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites                  []string                                      `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer                 string                                        `json:"projection_disclaimer"`
	DependencyState                      string                                        `json:"dependency_state"`
	TimelineProjectionState              string                                        `json:"timeline_projection_state"`
	SignalTimelineEntryState             string                                        `json:"signal_timeline_entry_state"`
	DisputeTimelineState                 string                                        `json:"dispute_timeline_state"`
	CorrectionReadProjectionState        string                                        `json:"correction_read_projection_state"`
	GovernanceTraceProjectionState       string                                        `json:"governance_trace_projection_state"`
	QueryProjectionState                 string                                        `json:"query_projection_state"`
	AccessBoundaryState                  string                                        `json:"access_boundary_state"`
	TenantPrivacyTimelineState           string                                        `json:"tenant_privacy_timeline_state"`
	AgentTimelineProjectionState         string                                        `json:"agent_timeline_projection_state"`
	TimestampIntegrityState              string                                        `json:"timestamp_integrity_state"`
	NoMutationProjectionGuardState       string                                        `json:"no_mutation_projection_guard_state"`
	NoOverclaimTimelineWordingState      string                                        `json:"no_overclaim_timeline_wording_state"`
	Dependency                           Point14ValDDependencySnapshot                 `json:"dependency"`
	TimelineProjection                   ExternalEcosystemTimelineProjection           `json:"timeline_projection"`
	SignalTimelineEntry                  ExternalSignalTimelineEntryProjection         `json:"signal_timeline_entry"`
	DisputeTimelineProjection            ExternalDisputeTimelineProjection             `json:"dispute_timeline_projection"`
	CorrectionReadProjection             CorrectionRevocationPublicationReadProjection `json:"correction_read_projection"`
	GovernanceTraceProjection            GovernanceTraceReadProjection                 `json:"governance_trace_projection"`
	QueryProjection                      EcosystemTimelineQueryProjection              `json:"query_projection"`
	AccessBoundary                       EcosystemTimelineAccessBoundary               `json:"access_boundary"`
	TenantPrivacyTimelineProjectionGuard TenantPrivacyTimelineProjectionGuard          `json:"tenant_privacy_timeline_projection_guard"`
	AgentTimelineProjection              AgentEcosystemTimelineProjection              `json:"agent_timeline_projection"`
	TimestampIntegrityGuard              TimelineTimestampIntegrityGuard               `json:"timestamp_integrity_guard"`
	NoMutationProjectionGuard            Point14ValDNoMutationProjectionGuard          `json:"no_mutation_projection_guard"`
	NoOverclaimTimelineWording           Point14ValDNoOverclaimTimelineWording         `json:"no_overclaim_timeline_wording"`
}

func point14ValDStates() []string {
	return []string{
		Point14ValDStateActive,
		Point14ValDStateBlocked,
		Point14ValDStateReviewRequired,
		Point14ValDStateIncomplete,
	}
}

func point14ValDStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point14ValDStates(), value)
}

func point14ValDQueryScopes() []string {
	return []string{
		point14ValDTimelineScopeEcosystem,
		point14ValDTimelineScopeSignals,
		point14ValDTimelineScopeDisputes,
		point14ValDTimelineScopePublication,
		point14ValDTimelineScopeGovernance,
		point14ValDTimelineScopeAgentAdvisory,
	}
}

func point14ValDAllowedFilters() []string {
	return []string{
		point14ValDFilterSignalRef,
		point14ValDFilterDisputeRef,
		point14ValDFilterPublicationState,
		point14ValDFilterTimelineState,
		point14ValDFilterStakeholderRole,
		point14ValDFilterVisibilityScope,
		point14ValDFilterGovernanceTrace,
		point14ValDFilterTimeWindow,
		point14ValDFilterLimitationVisible,
	}
}

func point14ValDSignalTimelineStates() []string {
	return []string{
		point14ValDSignalEntryVisible,
		point14ValDSignalEntryReviewRequired,
		point14ValDSignalEntryIncomplete,
		point14ValDSignalEntryBlocked,
	}
}

func point14ValDDisputeTimelineStates() []string {
	return []string{
		point14ValDDisputeTimelineVisible,
		point14ValDDisputeTimelineReview,
		point14ValDDisputeTimelineIncomplete,
		point14ValDDisputeTimelineBlocked,
	}
}

func point14ValDViewerRoles() []string {
	return append(append([]string{}, point14Val0RoleTypes()...), point14ValDViewerPublic)
}

func point14ValDTimelineProjectionIDValid(value string) bool {
	return point14Val0RefValid(value, "timeline_projection_")
}

func point14ValDSignalTimelineEntryIDValid(value string) bool {
	return point14Val0RefValid(value, "timeline_entry_")
}

func point14ValDDisputeTimelineIDValid(value string) bool {
	return point14Val0RefValid(value, "dispute_timeline_")
}

func point14ValDSourceProjectionRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "source_projection_")
}

func point14ValDSignalEntryRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "timeline_entry_")
}

func point14ValDDisputeEntryRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "dispute_timeline_")
}

func point14ValDCorrectionPublicationEntryRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "read_projection_")
}

func point14ValDGovernanceTraceRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "governance_trace_")
}

func point14ValDSourceIdentityRefValid(value string) bool {
	return point14Val0RefValid(value, "source_identity_")
}

func point14ValDValidationResultRefValid(value string) bool {
	return point14Val0RefValid(value, "validation_result_")
}

func point14ValDStakeholderRoleRefValid(value string) bool {
	return point14Val0RefValid(value, "stakeholder_role_")
}

func point14ValDReadProjectionIDValid(value string) bool {
	return point14Val0RefValid(value, "read_projection_")
}

func point14ValDPublicationStateRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "publication_state_ref_")
}

func point14ValDDecisionReasonRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "decision_reason_ref_")
}

func point14ValDTimestampRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "timestamp_ref_")
}

func point14ValDResultRefsValid(values []string) bool {
	return point14Val0RefListValid(values,
		"timeline_projection_",
		"timeline_entry_",
		"dispute_timeline_",
		"read_projection_",
		"governance_trace_projection_",
		"boundary_",
	)
}

func point14ValDQueryProjectionIDValid(value string) bool {
	return point14Val0RefValid(value, "query_projection_")
}

func point14ValDGovernanceTraceProjectionIDValid(value string) bool {
	return point14Val0RefValid(value, "governance_trace_projection_")
}

func point14ValDAgentTimelineProjectionIDValid(value string) bool {
	return point14Val0RefValid(value, "agent_timeline_projection_")
}

func point14ValDGuardIDValid(value string) bool {
	return point14Val0RefValid(value, "guard_", "boundary_")
}

func point14ValDForbiddenAuthorityMarkers() []string {
	markers := append([]string{}, point14ValCForbiddenAuthorityMarkers()...)
	return append(markers,
		"timeline_approved",
		"query_approved",
		"timeline_proves_truth",
	)
}

func point14ValDSafeWording() []string {
	return []string{
		"read-only ecosystem timeline",
		"bounded timeline projection",
		"dispute history requires governance context",
		"correction notice displayed as bounded record",
		"revocation request displayed as request only",
		"supersession record preserves prior context",
		"publication boundary shown without authority grant",
		"no public badge authority granted",
		"canonical evidence spine remains source of truth",
		"timeline does not resolve disputes",
		"limitations remain visible",
	}
}

func point14ValDForbiddenWording() []string {
	return append(
		append([]string{}, point14ValCForbiddenWording()...),
		"dispute resolved by timeline",
		"timeline proves truth",
		"query approved",
		"timeline approved",
	)
}

func point14ValDObservedTextContainsForbiddenWording(text string) bool {
	trimmed := strings.ToLower(strings.TrimSpace(text))
	if trimmed == "" {
		return false
	}
	for _, safe := range point14ValDSafeWording() {
		if trimmed == strings.ToLower(strings.TrimSpace(safe)) {
			return false
		}
	}
	for _, phrase := range point14ValDForbiddenWording() {
		if strings.Contains(trimmed, strings.ToLower(strings.TrimSpace(phrase))) {
			return true
		}
	}
	return false
}

func point14ValDObservedListContainsForbiddenWording(values []string) bool {
	for _, value := range values {
		if point14ValDObservedTextContainsForbiddenWording(value) {
			return true
		}
	}
	return false
}

func point14ValDExactValueValid(value string, allowed []string) bool {
	return point11Val0ContainsTrimmed(allowed, value)
}

func point14ValDValCPayloadContainsPointPass(valC Point14ValCFoundation) bool {
	payload, err := json.Marshal(valC)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point14ValABlockedPassToken)
}

func point14ValDDependencySnapshotFromUpstream(valC Point14ValCFoundation) Point14ValDDependencySnapshot {
	return Point14ValDDependencySnapshot{
		Point14ValCCurrentState:                  valC.CurrentState,
		Point14ValCDependencyState:               valC.DependencyState,
		Point14ValCCorrectionNoticeState:         valC.CorrectionNoticeState,
		Point14ValCRevocationRequestState:        valC.RevocationRequestState,
		Point14ValCSupersessionRecordState:       valC.SupersessionRecordState,
		Point14ValCPublicationApprovalState:      valC.PublicationApprovalState,
		Point14ValCVisibilityBoundaryState:       valC.VisibilityBoundaryState,
		Point14ValCTenantPrivacyState:            valC.TenantPrivacyState,
		Point14ValCRedactionLimitationState:      valC.RedactionLimitationState,
		Point14ValCGovernanceTraceState:          valC.GovernanceTraceState,
		Point14ValCAgentPublicationBoundaryState: valC.AgentPublicationBoundaryState,
		Point14ValCNoExternalAuthorityState:      valC.NoExternalAuthorityState,
		Point14ValCNoOverclaimState:              valC.NoOverclaimState,
		Point14ValCPointID:                       point14Val0PointID,
		Point14ValCWaveID:                        point14ValCWaveID,
		Point14ValCComputedFromUpstream:          valC.Dependency.SnapshotFromComputedOutput,
		Point14ValCMerged:                        true,
		Point14ValCCIGreen:                       true,
		Point14ValCReviewedOnMain:                true,
		Point14PassSeen:                          point14ValDValCPayloadContainsPointPass(valC),
		InheritedPoint14ValBCurrentState:         valC.Dependency.Point14ValBCurrentState,
		InheritedPoint14ValACurrentState:         valC.Dependency.InheritedPoint14ValACurrentState,
		InheritedPoint14Val0CurrentState:         valC.Dependency.InheritedPoint14Val0CurrentState,
		InheritedPoint13ValECurrentState:         valC.Dependency.InheritedPoint13ValECurrentState,
		InheritedPoint13ValEPassClosureState:     valC.Dependency.InheritedPoint13ValEPassClosureState,
		InheritedPoint13ValEPassAllowed:          valC.Dependency.InheritedPoint13ValEPassAllowed,
		InheritedPoint13ValEPassToken:            valC.Dependency.InheritedPoint13ValEPassToken,
		InheritedPoint12CurrentState:             valC.Dependency.InheritedPoint12CurrentState,
		InheritedPoint12DependencyState:          valC.Dependency.InheritedPoint12DependencyState,
		InheritedPoint12PassClosureState:         valC.Dependency.InheritedPoint12PassClosureState,
		InheritedPoint12ReviewerResult:           valC.Dependency.InheritedPoint12ReviewerResult,
		InheritedPoint11PublicationState:         valC.Dependency.InheritedPoint11PublicationState,
		InheritedPoint11NoOverclaimState:         valC.Dependency.InheritedPoint11NoOverclaimState,
		InheritedPoint11FinalPassGateState:       valC.Dependency.InheritedPoint11FinalPassGateState,
		InheritedPoint10CurrentState:             valC.Dependency.InheritedPoint10CurrentState,
		InheritedPoint10NoOverclaimState:         valC.Dependency.InheritedPoint10NoOverclaimState,
		InheritedPoint10ProjectionState:          valC.Dependency.InheritedPoint10ProjectionState,
		InheritedPoint10PassRuleState:            valC.Dependency.InheritedPoint10PassRuleState,
		InheritedTenantScope:                     valC.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:               true,
		ReviewPrerequisites:                      append([]string{}, valC.ReviewPrerequisites...),
		Point14ValC:                              valC,
		Point14ValB:                              valC.Dependency.Point14ValB,
		Point14ValA:                              valC.Dependency.Point14ValA,
		Point14Val0:                              valC.Dependency.Point14Val0,
		Point13ValE:                              valC.Dependency.Point13ValE,
		Point12:                                  valC.Dependency.Point12,
		Point11:                                  valC.Dependency.Point11,
		Point10:                                  valC.Dependency.Point10,
	}
}

func point14ValDDependencySnapshotModel() Point14ValDDependencySnapshot {
	valC := ComputePoint14ValCFoundation(Point14ValCFoundationModel())
	return point14ValDDependencySnapshotFromUpstream(valC)
}

func EvaluatePoint14ValDDependencyState(model Point14ValDDependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point14ValCComputedFromUpstream ||
		!model.Point14ValCMerged ||
		!model.Point14ValCCIGreen ||
		!model.Point14ValCReviewedOnMain ||
		model.Point14PassSeen ||
		strings.TrimSpace(model.Point14ValCPointID) != point14Val0PointID ||
		strings.TrimSpace(model.Point14ValCWaveID) != point14ValCWaveID ||
		!point14ValCStateValid(model.Point14ValCCurrentState) ||
		!point14ValCStateValid(model.Point14ValCDependencyState) ||
		!point14ValCStateValid(model.Point14ValCCorrectionNoticeState) ||
		!point14ValCStateValid(model.Point14ValCRevocationRequestState) ||
		!point14ValCStateValid(model.Point14ValCSupersessionRecordState) ||
		!point14ValCStateValid(model.Point14ValCPublicationApprovalState) ||
		!point14ValCStateValid(model.Point14ValCVisibilityBoundaryState) ||
		!point14ValCStateValid(model.Point14ValCTenantPrivacyState) ||
		!point14ValCStateValid(model.Point14ValCRedactionLimitationState) ||
		!point14ValCStateValid(model.Point14ValCGovernanceTraceState) ||
		!point14ValCStateValid(model.Point14ValCAgentPublicationBoundaryState) ||
		!point14ValCStateValid(model.Point14ValCNoExternalAuthorityState) ||
		!point14ValCStateValid(model.Point14ValCNoOverclaimState) ||
		!point14ValBStateValid(model.InheritedPoint14ValBCurrentState) ||
		!point14ValAStateValid(model.InheritedPoint14ValACurrentState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0CurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValECurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValEPassClosureState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		!point12ValEReviewerResultValid(model.InheritedPoint12ReviewerResult) ||
		strings.TrimSpace(model.InheritedPoint11PublicationState) == "" ||
		strings.TrimSpace(model.InheritedPoint11NoOverclaimState) == "" ||
		strings.TrimSpace(model.InheritedPoint11FinalPassGateState) == "" ||
		strings.TrimSpace(model.InheritedPoint10CurrentState) == "" ||
		strings.TrimSpace(model.InheritedPoint10NoOverclaimState) == "" ||
		strings.TrimSpace(model.InheritedPoint10ProjectionState) == "" ||
		strings.TrimSpace(model.InheritedPoint10PassRuleState) == "" ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return Point14ValDStateBlocked
	}
	if strings.TrimSpace(model.Point14ValCCurrentState) != strings.TrimSpace(model.Point14ValC.CurrentState) ||
		strings.TrimSpace(model.Point14ValCDependencyState) != strings.TrimSpace(model.Point14ValC.DependencyState) ||
		strings.TrimSpace(model.Point14ValCCorrectionNoticeState) != strings.TrimSpace(model.Point14ValC.CorrectionNoticeState) ||
		strings.TrimSpace(model.Point14ValCRevocationRequestState) != strings.TrimSpace(model.Point14ValC.RevocationRequestState) ||
		strings.TrimSpace(model.Point14ValCSupersessionRecordState) != strings.TrimSpace(model.Point14ValC.SupersessionRecordState) ||
		strings.TrimSpace(model.Point14ValCPublicationApprovalState) != strings.TrimSpace(model.Point14ValC.PublicationApprovalState) ||
		strings.TrimSpace(model.Point14ValCVisibilityBoundaryState) != strings.TrimSpace(model.Point14ValC.VisibilityBoundaryState) ||
		strings.TrimSpace(model.Point14ValCTenantPrivacyState) != strings.TrimSpace(model.Point14ValC.TenantPrivacyState) ||
		strings.TrimSpace(model.Point14ValCRedactionLimitationState) != strings.TrimSpace(model.Point14ValC.RedactionLimitationState) ||
		strings.TrimSpace(model.Point14ValCGovernanceTraceState) != strings.TrimSpace(model.Point14ValC.GovernanceTraceState) ||
		strings.TrimSpace(model.Point14ValCAgentPublicationBoundaryState) != strings.TrimSpace(model.Point14ValC.AgentPublicationBoundaryState) ||
		strings.TrimSpace(model.Point14ValCNoExternalAuthorityState) != strings.TrimSpace(model.Point14ValC.NoExternalAuthorityState) ||
		strings.TrimSpace(model.Point14ValCNoOverclaimState) != strings.TrimSpace(model.Point14ValC.NoOverclaimState) ||
		model.Point14ValCComputedFromUpstream != model.Point14ValC.Dependency.SnapshotFromComputedOutput ||
		strings.TrimSpace(model.InheritedPoint14ValBCurrentState) != strings.TrimSpace(model.Point14ValC.Dependency.Point14ValBCurrentState) ||
		strings.TrimSpace(model.InheritedPoint14ValACurrentState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint14ValACurrentState) ||
		strings.TrimSpace(model.InheritedPoint14Val0CurrentState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint14Val0CurrentState) ||
		strings.TrimSpace(model.InheritedPoint13ValECurrentState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint13ValECurrentState) ||
		strings.TrimSpace(model.InheritedPoint13ValEPassClosureState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint13ValEPassClosureState) ||
		model.InheritedPoint13ValEPassAllowed != model.Point14ValC.Dependency.InheritedPoint13ValEPassAllowed ||
		strings.TrimSpace(model.InheritedPoint13ValEPassToken) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint13ValEPassToken) ||
		strings.TrimSpace(model.InheritedPoint12CurrentState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint12CurrentState) ||
		strings.TrimSpace(model.InheritedPoint12DependencyState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint12DependencyState) ||
		strings.TrimSpace(model.InheritedPoint12PassClosureState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint12PassClosureState) ||
		strings.TrimSpace(model.InheritedPoint12ReviewerResult) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint12ReviewerResult) ||
		strings.TrimSpace(model.InheritedPoint11PublicationState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint11PublicationState) ||
		strings.TrimSpace(model.InheritedPoint11NoOverclaimState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint11NoOverclaimState) ||
		strings.TrimSpace(model.InheritedPoint11FinalPassGateState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint11FinalPassGateState) ||
		strings.TrimSpace(model.InheritedPoint10CurrentState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint10CurrentState) ||
		strings.TrimSpace(model.InheritedPoint10NoOverclaimState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint10NoOverclaimState) ||
		strings.TrimSpace(model.InheritedPoint10ProjectionState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint10ProjectionState) ||
		strings.TrimSpace(model.InheritedPoint10PassRuleState) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedPoint10PassRuleState) ||
		strings.TrimSpace(model.InheritedTenantScope) != strings.TrimSpace(model.Point14ValC.Dependency.InheritedTenantScope) {
		return Point14ValDStateBlocked
	}
	if strings.TrimSpace(model.Point14ValCCurrentState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCDependencyState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCCorrectionNoticeState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCRevocationRequestState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCSupersessionRecordState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCPublicationApprovalState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCVisibilityBoundaryState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCTenantPrivacyState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCRedactionLimitationState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCGovernanceTraceState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCAgentPublicationBoundaryState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCNoExternalAuthorityState) != Point14ValCStateActive ||
		strings.TrimSpace(model.Point14ValCNoOverclaimState) != Point14ValCStateActive ||
		strings.TrimSpace(model.InheritedPoint14ValBCurrentState) != Point14ValBStateActive ||
		strings.TrimSpace(model.InheritedPoint14ValACurrentState) != Point14ValAStateActive ||
		strings.TrimSpace(model.InheritedPoint14Val0CurrentState) != Point14Val0StateActive ||
		strings.TrimSpace(model.InheritedPoint13ValECurrentState) != Point13ValEStatePassConfirmed ||
		strings.TrimSpace(model.InheritedPoint13ValEPassClosureState) != Point13ValEStateActive ||
		!model.InheritedPoint13ValEPassAllowed ||
		strings.TrimSpace(model.InheritedPoint13ValEPassToken) != point13ValEPoint13PassToken ||
		strings.TrimSpace(model.InheritedPoint12CurrentState) != Point12ValEStatePassConfirmed ||
		strings.TrimSpace(model.InheritedPoint12DependencyState) != Point12ValEStateActive ||
		strings.TrimSpace(model.InheritedPoint12PassClosureState) != Point12ValEStateActive ||
		strings.TrimSpace(model.InheritedPoint12ReviewerResult) != point12ValEReviewerResultPassConfirmed ||
		strings.TrimSpace(model.InheritedPoint11PublicationState) != Point11ValDPublicationReviewStateActive ||
		strings.TrimSpace(model.InheritedPoint11NoOverclaimState) != Point11ValDNoOverclaimReviewStateActive ||
		strings.TrimSpace(model.InheritedPoint11FinalPassGateState) != Point11ValDFinalPassGateStateActive ||
		strings.TrimSpace(model.InheritedPoint10CurrentState) != operability.DeploymentMultiTenantPoint10StatePass ||
		strings.TrimSpace(model.InheritedPoint10NoOverclaimState) != operability.DeploymentMultiTenantValENoOverclaimStateActive ||
		strings.TrimSpace(model.InheritedPoint10ProjectionState) != operability.DeploymentMultiTenantValEProjectionBoundaryStateActive ||
		strings.TrimSpace(model.InheritedPoint10PassRuleState) != operability.DeploymentMultiTenantValEPoint10PassRuleStateActive {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDTimelineProjectionModel(dependency Point14ValDDependencySnapshot) ExternalEcosystemTimelineProjection {
	return ExternalEcosystemTimelineProjection{
		TimelineProjectionID:           "timeline_projection_point14_vald_001",
		TenantScope:                    dependency.InheritedTenantScope,
		SourceProjectionRefs:           []string{"source_projection_point14_vald_001"},
		SignalEntryRefs:                []string{"timeline_entry_point14_vald_001"},
		DisputeEntryRefs:               []string{"dispute_timeline_point14_vald_001"},
		CorrectionPublicationEntryRefs: []string{"read_projection_point14_vald_001"},
		GovernanceTraceRefs:            []string{"governance_trace_projection_point14_vald_001"},
		GeneratedAt:                    "2026-05-06T09:10:00Z",
		GeneratedTimeSource:            point14Val0TimeSourceServerUTC,
		ReadOnly:                       true,
		ProjectionOnly:                 true,
	}
}

func EvaluatePoint14ValDTimelineProjectionState(model ExternalEcosystemTimelineProjection) string {
	if !point14ValDTimelineProjectionIDValid(model.TimelineProjectionID) ||
		(len(strings.TrimSpace(model.TenantScope)) == 0 && len(strings.TrimSpace(model.GlobalScopeClassification)) == 0) ||
		!point14ValDSourceProjectionRefsValid(model.SourceProjectionRefs) ||
		!point14ValDSignalEntryRefsValid(model.SignalEntryRefs) ||
		(len(model.DisputeEntryRefs) > 0 && !point14ValDDisputeEntryRefsValid(model.DisputeEntryRefs)) ||
		(len(model.CorrectionPublicationEntryRefs) > 0 && !point14ValDCorrectionPublicationEntryRefsValid(model.CorrectionPublicationEntryRefs)) ||
		!point14ValDGovernanceTraceRefsValid(model.GovernanceTraceRefs) ||
		!point14Val0ParsedTimeOk(model.GeneratedAt) ||
		!point14Val0CanonicalTimeSourceValid(model.GeneratedTimeSource) ||
		!model.ReadOnly ||
		!model.ProjectionOnly {
		return Point14ValDStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != "" {
		if !point11Val0ScopeValid(model.TenantScope) || strings.TrimSpace(model.GlobalScopeClassification) != "" {
			return Point14ValDStateBlocked
		}
	} else if !point11Val0ContainsTrimmed([]string{point14Val0ScopeGlobalAdvisory, point14Val0ScopePublicNonAuthorative}, model.GlobalScopeClassification) {
		return Point14ValDStateBlocked
	}
	if model.MutatesCanonicalEvidence ||
		model.MutatesSignalState ||
		model.ResolvesDispute ||
		model.PublishesCorrection ||
		model.RevokesClaim ||
		model.ApprovesProduction ||
		model.CertifiesCompliance ||
		model.CreatesPublicBadge ||
		model.EmitsPass {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDSignalTimelineEntryModel(dependency Point14ValDDependencySnapshot) ExternalSignalTimelineEntryProjection {
	return ExternalSignalTimelineEntryProjection{
		TimelineEntryID:     "timeline_entry_point14_vald_001",
		NormalizedSignalRef: "normalized_signal_point14_vala_001",
		ValidationResultRef: "validation_result_point14_vala_001",
		SourceIdentityRef:   "source_identity_point14_vala_001",
		StakeholderRoleRef:  "stakeholder_role_point14_valb_001",
		EventAt:             "2026-05-06T09:00:00Z",
		EventTimeSource:     point14Val0TimeSourceServerUTC,
		ReceivedAt:          "2026-05-06T09:01:00Z",
		ReceivedTimeSource:  point14Val0TimeSourceServerUTC,
		TimelineState:       point14ValDSignalEntryVisible,
		AdvisoryOnly:        true,
	}
}

func EvaluatePoint14ValDSignalTimelineEntryState(model ExternalSignalTimelineEntryProjection) string {
	if !point14ValDSignalTimelineEntryIDValid(model.TimelineEntryID) ||
		!point14Val0RefValid(model.NormalizedSignalRef, "normalized_signal_", "signal_") ||
		!point14ValDValidationResultRefValid(model.ValidationResultRef) ||
		!point14ValDSourceIdentityRefValid(model.SourceIdentityRef) ||
		!point14ValDStakeholderRoleRefValid(model.StakeholderRoleRef) ||
		!point14Val0ParsedTimeOk(model.EventAt) ||
		!point14Val0CanonicalTimeSourceValid(model.EventTimeSource) ||
		!point14Val0ParsedTimeOk(model.ReceivedAt) ||
		!point14Val0CanonicalTimeSourceValid(model.ReceivedTimeSource) ||
		!point14ValDExactValueValid(model.TimelineState, point14ValDSignalTimelineStates()) {
		return Point14ValDStateBlocked
	}
	eventAt, _ := point14Val0ParsedTime(model.EventAt)
	receivedAt, _ := point14Val0ParsedTime(model.ReceivedAt)
	if eventAt.After(receivedAt) {
		return Point14ValDStateReviewRequired
	}
	if !model.AdvisoryOnly || model.AuthorityGranted || model.UpgradesSignalValidity || model.SourceEventAsCanonicalAuthority {
		return Point14ValDStateBlocked
	}
	switch strings.TrimSpace(model.TimelineState) {
	case point14ValDSignalEntryBlocked:
		return Point14ValDStateBlocked
	case point14ValDSignalEntryReviewRequired:
		return Point14ValDStateReviewRequired
	case point14ValDSignalEntryIncomplete:
		return Point14ValDStateIncomplete
	default:
		return Point14ValDStateActive
	}
}

func point14ValDDisputeTimelineProjectionModel() ExternalDisputeTimelineProjection {
	return ExternalDisputeTimelineProjection{
		DisputeTimelineID:        "dispute_timeline_point14_vald_001",
		DisputeRef:               "dispute_point14_valb_001",
		ConflictSetRef:           "conflict_set_point14_valb_001",
		TriageResultRef:          "triage_result_point14_valb_001",
		LifecycleState:           point14Val0DisputeTriaged,
		EvidenceRequirementRefs:  []string{"evidence_requirement_point14_valb_001"},
		GovernanceEscalationRefs: []string{"escalation_point14_valb_001"},
		TimelineState:            point14ValDDisputeTimelineVisible,
	}
}

func EvaluatePoint14ValDDisputeTimelineProjectionState(model ExternalDisputeTimelineProjection) string {
	if !point14ValDDisputeTimelineIDValid(model.DisputeTimelineID) ||
		!point14Val0DisputeIDValid(model.DisputeRef) ||
		!point14ValBConflictSetIDValid(model.ConflictSetRef) ||
		!point14ValBTriageResultIDValid(model.TriageResultRef) ||
		!point14ValDExactValueValid(model.LifecycleState, point14ValBDisputeLifecycleStates()) ||
		!point14ValDExactValueValid(model.TimelineState, point14ValDDisputeTimelineStates()) {
		return Point14ValDStateBlocked
	}
	if strings.TrimSpace(model.LifecycleState) == point14Val0DisputeEvidenceNeeded && !point14Val0RefListValid(model.EvidenceRequirementRefs, "evidence_requirement_") {
		return Point14ValDStateBlocked
	}
	if strings.TrimSpace(model.LifecycleState) == point14Val0DisputeReviewRequired && len(model.GovernanceEscalationRefs) > 0 && !point14Val0RefListValid(model.GovernanceEscalationRefs, "escalation_") {
		return Point14ValDStateBlocked
	}
	if model.ResolvesDispute ||
		model.MovesLifecycleToCorrected ||
		model.MovesLifecycleToRevoked ||
		model.MovesLifecycleToPublishedNotice ||
		model.HidesEvidenceRequired ||
		model.ConvertsReviewIncompleteToActive {
		return Point14ValDStateBlocked
	}
	if strings.TrimSpace(model.LifecycleState) == point14Val0DisputeEvidenceNeeded && strings.TrimSpace(model.TimelineState) == point14ValDDisputeTimelineVisible {
		return Point14ValDStateIncomplete
	}
	switch strings.TrimSpace(model.TimelineState) {
	case point14ValDDisputeTimelineBlocked:
		return Point14ValDStateBlocked
	case point14ValDDisputeTimelineReview:
		return Point14ValDStateReviewRequired
	case point14ValDDisputeTimelineIncomplete:
		return Point14ValDStateIncomplete
	default:
		return Point14ValDStateActive
	}
}

func point14ValDCorrectionReadProjectionModel() CorrectionRevocationPublicationReadProjection {
	return CorrectionRevocationPublicationReadProjection{
		ReadProjectionID:        "read_projection_point14_vald_001",
		CorrectionNoticeRefs:    []string{"correction_notice_point14_valc_001"},
		RevocationRequestRefs:   []string{"revocation_request_point14_valc_001"},
		SupersessionRecordRefs:  []string{"supersession_record_point14_valc_001"},
		PublicationApprovalRefs: []string{"publication_approval_point14_valc_001"},
		VisibilityBoundaryRefs:  []string{"boundary_point14_valc_visibility_001"},
		LimitationRefs:          []string{"limitation_ref_point14_valc_001"},
		RedactionRefs:           []string{"redaction_ref_point14_valc_001"},
		PublicationStateRefs:    []string{"publication_state_ref_point14_valc_001"},
		ReadOnly:                true,
		ObservedReadTexts: []string{
			"correction notice displayed as bounded record",
			"revocation request displayed as request only",
			"supersession record preserves prior context",
			"publication boundary shown without authority grant",
		},
	}
}

func EvaluatePoint14ValDCorrectionReadProjectionState(model CorrectionRevocationPublicationReadProjection) string {
	if !point14ValDReadProjectionIDValid(model.ReadProjectionID) ||
		!model.ReadOnly ||
		!point14Val0RefListValid(model.VisibilityBoundaryRefs, "boundary_") ||
		!point14ValDPublicationStateRefsValid(model.PublicationStateRefs) {
		return Point14ValDStateBlocked
	}
	hasScope := false
	if len(model.CorrectionNoticeRefs) > 0 {
		hasScope = true
		if !point14Val0RefListValid(model.CorrectionNoticeRefs, "correction_notice_") {
			return Point14ValDStateBlocked
		}
	}
	if len(model.RevocationRequestRefs) > 0 {
		hasScope = true
		if !point14Val0RefListValid(model.RevocationRequestRefs, "revocation_request_") {
			return Point14ValDStateBlocked
		}
	}
	if len(model.SupersessionRecordRefs) > 0 {
		hasScope = true
		if !point14Val0RefListValid(model.SupersessionRecordRefs, "supersession_record_") {
			return Point14ValDStateBlocked
		}
	}
	if len(model.PublicationApprovalRefs) > 0 {
		hasScope = true
		if !point14Val0RefListValid(model.PublicationApprovalRefs, "publication_approval_") {
			return Point14ValDStateBlocked
		}
	}
	if !hasScope {
		return Point14ValDStateBlocked
	}
	if (len(model.LimitationRefs) == 0 && len(model.RedactionRefs) > 0) ||
		(len(model.LimitationRefs) > 0 && !point14ValCLimitationRefsValid(model.LimitationRefs)) ||
		(len(model.RedactionRefs) > 0 && !point14ValCRedactionRefsValid(model.RedactionRefs)) {
		return Point14ValDStateBlocked
	}
	if model.PublishesCorrection ||
		model.ExecutesRevocation ||
		model.SilentReplacesSupersededSignal ||
		model.OmitsLimitations ||
		model.HidesRedaction ||
		point14ValDObservedListContainsForbiddenWording(model.ObservedReadTexts) {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDGovernanceTraceReadProjectionModel() GovernanceTraceReadProjection {
	return GovernanceTraceReadProjection{
		GovernanceTraceProjectionID: "governance_trace_projection_point14_vald_001",
		GovernanceTraceRefs:         []string{"governance_trace_point14_valc_001"},
		OwnerRefs:                   []string{"approver_point14_valc_owner_001"},
		ApproverRoleRefs:            []string{"role_ref_point14_vald_security_reviewer_001"},
		AuditRefs:                   []string{"audit_event_point14_valc_governance_001"},
		EvidenceRefs:                []string{"evidence_point14_valc_001"},
		DecisionReasonRefs:          []string{"decision_reason_ref_point14_vald_001"},
		TimestampRefs:               []string{"timestamp_ref_point14_vald_001"},
		ReadOnly:                    true,
	}
}

func EvaluatePoint14ValDGovernanceTraceProjectionState(model GovernanceTraceReadProjection) string {
	if !point14ValDGovernanceTraceProjectionIDValid(model.GovernanceTraceProjectionID) ||
		!point14ValDGovernanceTraceRefsValid(model.GovernanceTraceRefs) ||
		!point14Val0RefListValid(model.OwnerRefs, "approver_") ||
		!point14Val0RefListValid(model.ApproverRoleRefs, "role_ref_") ||
		!point14Val0AuditEventRefsValid(model.AuditRefs) ||
		!point14Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point14ValDDecisionReasonRefsValid(model.DecisionReasonRefs) ||
		!point14ValDTimestampRefsValid(model.TimestampRefs) ||
		!model.ReadOnly {
		return Point14ValDStateBlocked
	}
	if model.ApprovesAnything || model.MutatesGovernanceTrace || model.SatisfiesMissingGovernanceByDisplay {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDQueryProjectionModel() EcosystemTimelineQueryProjection {
	return EcosystemTimelineQueryProjection{
		QueryProjectionID:     "query_projection_point14_vald_001",
		QueryScope:            point14ValDTimelineScopeEcosystem,
		AllowedFilters:        []string{point14ValDFilterSignalRef, point14ValDFilterTimelineState, point14ValDFilterGovernanceTrace},
		ResultRefs:            []string{"timeline_projection_point14_vald_001", "dispute_timeline_point14_vald_001", "read_projection_point14_vald_001"},
		ReadOnly:              true,
		QueryIsProjectionOnly: true,
	}
}

func EvaluatePoint14ValDQueryProjectionState(model EcosystemTimelineQueryProjection) string {
	if !point14ValDQueryProjectionIDValid(model.QueryProjectionID) ||
		!point14ValDExactValueValid(model.QueryScope, point14ValDQueryScopes()) ||
		!point14Val0TextListValid(model.AllowedFilters) ||
		!point14ValDResultRefsValid(model.ResultRefs) ||
		!model.ReadOnly ||
		!model.QueryIsProjectionOnly {
		return Point14ValDStateBlocked
	}
	for _, filter := range model.AllowedFilters {
		if !point14ValDExactValueValid(filter, point14ValDAllowedFilters()) {
			return Point14ValDStateBlocked
		}
	}
	if model.MutationRequested ||
		model.WritesFiltersBack ||
		model.HidesDecisiveMissingEvidence ||
		model.OmitsLimitationsWithoutDisclosure ||
		model.CrossTenantResults {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDAccessBoundaryModel(dependency Point14ValDDependencySnapshot) EcosystemTimelineAccessBoundary {
	return EcosystemTimelineAccessBoundary{
		AccessBoundaryID: "boundary_point14_vald_access_001",
		ViewerRole:       "customer_admin",
		TenantScope:      dependency.InheritedTenantScope,
		AllowedViewScope: point14ValCVisibilityCustomerBounded,
		AuditRef:         "audit_event_point14_vald_access_001",
		AccessTime:       "2026-05-06T09:11:00Z",
		AccessTimeSource: point14Val0TimeSourceServerUTC,
	}
}

func EvaluatePoint14ValDAccessBoundaryState(model EcosystemTimelineAccessBoundary, dependency Point14ValDDependencySnapshot) string {
	if !point14Val0BoundaryRefValid(model.AccessBoundaryID) ||
		!point14ValDExactValueValid(model.ViewerRole, point14ValDViewerRoles()) ||
		(strings.TrimSpace(model.TenantScope) == "" && strings.TrimSpace(model.GlobalScopeClassification) == "") ||
		!point14ValDExactValueValid(model.AllowedViewScope, point14ValCVisibilityClassifications()) ||
		!point14Val0AuditEventRefValid(model.AuditRef) ||
		!point14Val0ParsedTimeOk(model.AccessTime) ||
		!point14Val0CanonicalTimeSourceValid(model.AccessTimeSource) {
		return Point14ValDStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) != "" {
		if !point11Val0ScopeValid(model.TenantScope) || strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) || strings.TrimSpace(model.GlobalScopeClassification) != "" {
			return Point14ValDStateBlocked
		}
	} else if !point11Val0ContainsTrimmed([]string{point14Val0ScopeGlobalAdvisory, point14Val0ScopePublicNonAuthorative}, model.GlobalScopeClassification) {
		return Point14ValDStateBlocked
	}
	if model.AccessExpired || model.AccessRevoked || model.CrossTenantAccess || model.AuthorityGranted || model.TenantPrivateDataExposed {
		return Point14ValDStateBlocked
	}
	if strings.TrimSpace(model.ViewerRole) == point14ValDViewerPublic && strings.TrimSpace(model.AllowedViewScope) == point14ValCVisibilityPrivateTenantOnly {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDTenantPrivacyTimelineProjectionGuardModel(dependency Point14ValDDependencySnapshot) TenantPrivacyTimelineProjectionGuard {
	return TenantPrivacyTimelineProjectionGuard{
		BoundaryID:                  "boundary_point14_vald_privacy_001",
		TenantScope:                 dependency.InheritedTenantScope,
		ProjectionTargetScope:       point14ValCVisibilityPrivateTenantOnly,
		PublicPrivateClassification: point14ValCPublicationBoundaryPrivate,
		RedactionRefs:               []string{"redaction_ref_point14_valc_001"},
		LimitationRefs:              []string{"limitation_ref_point14_valc_001"},
		ObservedSummaryTexts:        []string{"limitations remain visible"},
		LimitationsVisible:          true,
	}
}

func EvaluatePoint14ValDTenantPrivacyTimelineProjectionGuardState(model TenantPrivacyTimelineProjectionGuard, dependency Point14ValDDependencySnapshot) string {
	if !point14Val0BoundaryRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		!point14ValDExactValueValid(model.ProjectionTargetScope, point14ValCVisibilityClassifications()) ||
		!point14ValDExactValueValid(model.PublicPrivateClassification, point14ValCPublicPrivateBoundaries()) ||
		!point14ValCPublicationBoundaryPairValid(model.ProjectionTargetScope, model.PublicPrivateClassification) ||
		!point14ValCLimitationRefsValid(model.LimitationRefs) ||
		!point14ValCRedactionRefsValid(model.RedactionRefs) {
		return Point14ValDStateBlocked
	}
	if model.TenantPrivateDataExposed || model.CrossTenantProjection || !model.LimitationsVisible || model.StrengthensClaim || point14ValDObservedListContainsForbiddenWording(model.ObservedSummaryTexts) {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDAgentTimelineProjectionModel(dependency Point14ValDDependencySnapshot) AgentEcosystemTimelineProjection {
	return AgentEcosystemTimelineProjection{
		ProjectionID:            "agent_timeline_projection_point14_vald_001",
		TenantScope:             dependency.InheritedTenantScope,
		AgentInputRefs:          []string{"agent_input_point14_vald_001"},
		AgentRecommendationRefs: []string{"agent_recommendation_point14_vald_001"},
		EvidenceRefs:            []string{"evidence_point14_valc_001"},
		AuditEventRef:           "audit_event_point14_vald_agent_001",
		AdvisoryOnly:            true,
	}
}

func EvaluatePoint14ValDAgentTimelineProjectionState(model AgentEcosystemTimelineProjection, dependency Point14ValDDependencySnapshot) string {
	if !point14ValDAgentTimelineProjectionIDValid(model.ProjectionID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(dependency.InheritedTenantScope) ||
		(len(model.AgentInputRefs) > 0 && !point14Val0AgentInputRefsValid(model.AgentInputRefs)) ||
		(len(model.AgentRecommendationRefs) > 0 && !point14Val0RefListValid(model.AgentRecommendationRefs, "agent_recommendation_")) ||
		!point14Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point14Val0AuditEventRefValid(model.AuditEventRef) {
		return Point14ValDStateBlocked
	}
	if !model.AdvisoryOnly ||
		model.AgentApprovalFlags ||
		model.AgentAuthorityFlags ||
		model.CanResolveDispute ||
		model.CanPublishCorrection ||
		model.CanRevokeClaim ||
		model.CanSatisfyGovernanceTrace ||
		model.PassAllowed ||
		model.ExternalAuthorityAllowed {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDTimestampIntegrityGuardModel() TimelineTimestampIntegrityGuard {
	return TimelineTimestampIntegrityGuard{
		GuardID:                       "guard_point14_vald_timestamp_001",
		TimelineEventRefs:             []string{"timeline_event_point14_vald_001"},
		EventAt:                       "2026-05-06T09:00:00Z",
		EventTimeSource:               point14Val0TimeSourceServerUTC,
		ReceivedAt:                    "2026-05-06T09:01:00Z",
		ReceivedTimeSource:            point14Val0TimeSourceServerUTC,
		GeneratedAt:                   "2026-05-06T09:10:00Z",
		GeneratedTimeSource:           point14Val0TimeSourceServerUTC,
		AccessTime:                    "2026-05-06T09:11:00Z",
		AccessTimeSource:              point14Val0TimeSourceServerUTC,
		SourceEventAt:                 "2026-05-06T08:59:30Z",
		SourceEventTimeSource:         point14Val0TimeSourceClientLocal,
		PublicationApprovedAt:         "2026-05-06T09:05:00Z",
		PublicationApprovedTimeSource: point14Val0TimeSourceServerUTC,
		DisputeOpenedAt:               "2026-05-06T09:00:00Z",
		DisputeOpenedTimeSource:       point14Val0TimeSourceServerUTC,
	}
}

func EvaluatePoint14ValDTimestampIntegrityGuardState(model TimelineTimestampIntegrityGuard) string {
	if !point14ValDGuardIDValid(model.GuardID) ||
		!point14Val0RefListValid(model.TimelineEventRefs, "timeline_event_") ||
		!point14Val0ParsedTimeOk(model.EventAt) ||
		!point14Val0CanonicalTimeSourceValid(model.EventTimeSource) ||
		!point14Val0ParsedTimeOk(model.GeneratedAt) ||
		!point14Val0CanonicalTimeSourceValid(model.GeneratedTimeSource) ||
		!point14Val0ParsedTimeOk(model.AccessTime) ||
		!point14Val0CanonicalTimeSourceValid(model.AccessTimeSource) {
		return Point14ValDStateBlocked
	}
	if strings.TrimSpace(model.ReceivedAt) != "" {
		if !point14Val0ParsedTimeOk(model.ReceivedAt) || !point14Val0CanonicalTimeSourceValid(model.ReceivedTimeSource) {
			return Point14ValDStateBlocked
		}
	}
	if strings.TrimSpace(model.SourceEventAt) != "" {
		if !point14Val0ParsedTimeOk(model.SourceEventAt) || !point14Val0TimeSourceValid(model.SourceEventTimeSource) {
			return Point14ValDStateBlocked
		}
	}
	if strings.TrimSpace(model.PublicationApprovedAt) != "" {
		if !point14Val0ParsedTimeOk(model.PublicationApprovedAt) || !point14Val0CanonicalTimeSourceValid(model.PublicationApprovedTimeSource) {
			return Point14ValDStateBlocked
		}
	}
	if strings.TrimSpace(model.DisputeOpenedAt) != "" {
		if !point14Val0ParsedTimeOk(model.DisputeOpenedAt) || !point14Val0CanonicalTimeSourceValid(model.DisputeOpenedTimeSource) {
			return Point14ValDStateBlocked
		}
	}
	eventAt, _ := point14Val0ParsedTime(model.EventAt)
	generatedAt, _ := point14Val0ParsedTime(model.GeneratedAt)
	accessAt, _ := point14Val0ParsedTime(model.AccessTime)
	if eventAt.After(generatedAt) || generatedAt.After(accessAt) {
		return Point14ValDStateReviewRequired
	}
	if strings.TrimSpace(model.ReceivedAt) != "" {
		receivedAt, _ := point14Val0ParsedTime(model.ReceivedAt)
		if eventAt.After(receivedAt) || receivedAt.After(generatedAt) {
			return Point14ValDStateReviewRequired
		}
	}
	if strings.TrimSpace(model.PublicationApprovedAt) != "" && strings.TrimSpace(model.DisputeOpenedAt) != "" {
		publicationApprovedAt, _ := point14Val0ParsedTime(model.PublicationApprovedAt)
		disputeOpenedAt, _ := point14Val0ParsedTime(model.DisputeOpenedAt)
		if publicationApprovedAt.Before(disputeOpenedAt) {
			return Point14ValDStateReviewRequired
		}
	}
	if model.AttemptsValidityUpgrade {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDNoMutationProjectionGuardModel() Point14ValDNoMutationProjectionGuard {
	return Point14ValDNoMutationProjectionGuard{
		BoundaryID: "boundary_point14_vald_nomutation_001",
	}
}

func EvaluatePoint14ValDNoMutationProjectionGuardState(model Point14ValDNoMutationProjectionGuard) string {
	if !point14ValDGuardIDValid(model.BoundaryID) {
		return Point14ValDStateBlocked
	}
	if model.MutatesCanonicalEvidence ||
		model.MutatesNormalizedSignal ||
		model.MutatesValidationResult ||
		model.MutatesDisputeLifecycle ||
		model.MutatesCorrectionNotice ||
		model.MutatesRevocationRequest ||
		model.MutatesSupersessionRecord ||
		model.MutatesPublicationApproval ||
		model.MutatesVisibilityBoundary ||
		model.MutatesGovernanceTrace ||
		model.PublishesCorrection ||
		model.ExecutesRevocation ||
		model.ResolvesDispute ||
		model.ApprovesProduction ||
		model.CertifiesCompliance ||
		model.CreatesPublicBadge ||
		model.EmitsPass {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDNoOverclaimTimelineWordingModel() Point14ValDNoOverclaimTimelineWording {
	return Point14ValDNoOverclaimTimelineWording{
		ObservedTimelineTexts: []string{
			"read-only ecosystem timeline",
			"bounded timeline projection",
		},
		ObservedSignalTexts: []string{
			"canonical evidence spine remains source of truth",
		},
		ObservedDisputeTexts: []string{
			"dispute history requires governance context",
			"timeline does not resolve disputes",
		},
		ObservedReadProjectionTexts: []string{
			"correction notice displayed as bounded record",
			"revocation request displayed as request only",
			"supersession record preserves prior context",
			"publication boundary shown without authority grant",
		},
		ObservedGovernanceTexts: []string{
			"limitations remain visible",
		},
		ObservedQueryTexts: []string{
			"bounded timeline projection",
		},
		ObservedAccessTexts: []string{
			"no public badge authority granted",
		},
		ObservedPrivacyTexts: []string{
			"limitations remain visible",
		},
		ObservedAgentTexts: []string{
			"timeline does not resolve disputes",
		},
		AllowedSafeWording:   point14ValDSafeWording(),
		BlockedWording:       point14ValDForbiddenWording(),
		ProjectionDisclaimer: point14ValDProjectionDisclaimerBase,
	}
}

func EvaluatePoint14ValDNoOverclaimTimelineWordingState(model Point14ValDNoOverclaimTimelineWording) string {
	if strings.TrimSpace(model.ProjectionDisclaimer) != point14ValDProjectionDisclaimerBase ||
		!point14Val0TextListValid(model.AllowedSafeWording) ||
		!point14Val0TextListValid(model.BlockedWording) {
		return Point14ValDStateBlocked
	}
	if point14ValDObservedListContainsForbiddenWording(model.ObservedTimelineTexts) ||
		point14ValDObservedListContainsForbiddenWording(model.ObservedSignalTexts) ||
		point14ValDObservedListContainsForbiddenWording(model.ObservedDisputeTexts) ||
		point14ValDObservedListContainsForbiddenWording(model.ObservedReadProjectionTexts) ||
		point14ValDObservedListContainsForbiddenWording(model.ObservedGovernanceTexts) ||
		point14ValDObservedListContainsForbiddenWording(model.ObservedQueryTexts) ||
		point14ValDObservedListContainsForbiddenWording(model.ObservedAccessTexts) ||
		point14ValDObservedListContainsForbiddenWording(model.ObservedPrivacyTexts) ||
		point14ValDObservedListContainsForbiddenWording(model.ObservedAgentTexts) {
		return Point14ValDStateBlocked
	}
	if point14ValDObservedListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point14ValDStateBlocked
	}
	return Point14ValDStateActive
}

func point14ValDFoundationState(states ...string) string {
	hasReview := false
	hasIncomplete := false
	for _, state := range states {
		switch strings.TrimSpace(state) {
		case Point14ValDStateBlocked:
			return Point14ValDStateBlocked
		case Point14ValDStateReviewRequired:
			hasReview = true
		case Point14ValDStateIncomplete:
			hasIncomplete = true
		}
	}
	if hasReview {
		return Point14ValDStateReviewRequired
	}
	if hasIncomplete {
		return Point14ValDStateIncomplete
	}
	return Point14ValDStateActive
}

func point14ValDBlockingReasons(model Point14ValDFoundation) []string {
	componentStates := map[string]string{
		"dependency":                  model.DependencyState,
		"timeline_projection":         model.TimelineProjectionState,
		"signal_timeline_entry":       model.SignalTimelineEntryState,
		"dispute_timeline":            model.DisputeTimelineState,
		"correction_read_projection":  model.CorrectionReadProjectionState,
		"governance_trace_projection": model.GovernanceTraceProjectionState,
		"query_projection":            model.QueryProjectionState,
		"access_boundary":             model.AccessBoundaryState,
		"tenant_privacy_timeline":     model.TenantPrivacyTimelineState,
		"agent_timeline_projection":   model.AgentTimelineProjectionState,
		"timestamp_integrity":         model.TimestampIntegrityState,
		"no_mutation_projection":      model.NoMutationProjectionGuardState,
		"no_overclaim":                model.NoOverclaimTimelineWordingState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if state == Point14ValDStateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point14ValDReviewPrerequisites(model Point14ValDFoundation) []string {
	componentStates := map[string]string{
		"dependency":                  model.DependencyState,
		"timeline_projection":         model.TimelineProjectionState,
		"signal_timeline_entry":       model.SignalTimelineEntryState,
		"dispute_timeline":            model.DisputeTimelineState,
		"correction_read_projection":  model.CorrectionReadProjectionState,
		"governance_trace_projection": model.GovernanceTraceProjectionState,
		"query_projection":            model.QueryProjectionState,
		"access_boundary":             model.AccessBoundaryState,
		"tenant_privacy_timeline":     model.TenantPrivacyTimelineState,
		"agent_timeline_projection":   model.AgentTimelineProjectionState,
		"timestamp_integrity":         model.TimestampIntegrityState,
		"no_mutation_projection":      model.NoMutationProjectionGuardState,
		"no_overclaim":                model.NoOverclaimTimelineWordingState,
	}
	prereqs := []string{}
	for name, state := range componentStates {
		if state == Point14ValDStateReviewRequired || state == Point14ValDStateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func point14ValDFoundationModelFromUpstream(valC Point14ValCFoundation) Point14ValDFoundation {
	dependency := point14ValDDependencySnapshotFromUpstream(valC)
	return Point14ValDFoundation{
		CurrentState:                         Point14ValDStateActive,
		ProjectionDisclaimer:                 point14ValDProjectionDisclaimerBase,
		DependencyState:                      Point14ValDStateActive,
		TimelineProjectionState:              Point14ValDStateActive,
		SignalTimelineEntryState:             Point14ValDStateActive,
		DisputeTimelineState:                 Point14ValDStateActive,
		CorrectionReadProjectionState:        Point14ValDStateActive,
		GovernanceTraceProjectionState:       Point14ValDStateActive,
		QueryProjectionState:                 Point14ValDStateActive,
		AccessBoundaryState:                  Point14ValDStateActive,
		TenantPrivacyTimelineState:           Point14ValDStateActive,
		AgentTimelineProjectionState:         Point14ValDStateActive,
		TimestampIntegrityState:              Point14ValDStateActive,
		NoMutationProjectionGuardState:       Point14ValDStateActive,
		NoOverclaimTimelineWordingState:      Point14ValDStateActive,
		Dependency:                           dependency,
		TimelineProjection:                   point14ValDTimelineProjectionModel(dependency),
		SignalTimelineEntry:                  point14ValDSignalTimelineEntryModel(dependency),
		DisputeTimelineProjection:            point14ValDDisputeTimelineProjectionModel(),
		CorrectionReadProjection:             point14ValDCorrectionReadProjectionModel(),
		GovernanceTraceProjection:            point14ValDGovernanceTraceReadProjectionModel(),
		QueryProjection:                      point14ValDQueryProjectionModel(),
		AccessBoundary:                       point14ValDAccessBoundaryModel(dependency),
		TenantPrivacyTimelineProjectionGuard: point14ValDTenantPrivacyTimelineProjectionGuardModel(dependency),
		AgentTimelineProjection:              point14ValDAgentTimelineProjectionModel(dependency),
		TimestampIntegrityGuard:              point14ValDTimestampIntegrityGuardModel(),
		NoMutationProjectionGuard:            point14ValDNoMutationProjectionGuardModel(),
		NoOverclaimTimelineWording:           point14ValDNoOverclaimTimelineWordingModel(),
	}
}

func Point14ValDFoundationModel() Point14ValDFoundation {
	valC := ComputePoint14ValCFoundation(Point14ValCFoundationModel())
	return point14ValDFoundationModelFromUpstream(valC)
}

func ComputePoint14ValDFoundation(model Point14ValDFoundation) Point14ValDFoundation {
	model.DependencyState = EvaluatePoint14ValDDependencyState(model.Dependency)
	model.TimelineProjectionState = EvaluatePoint14ValDTimelineProjectionState(model.TimelineProjection)
	model.SignalTimelineEntryState = EvaluatePoint14ValDSignalTimelineEntryState(model.SignalTimelineEntry)
	model.DisputeTimelineState = EvaluatePoint14ValDDisputeTimelineProjectionState(model.DisputeTimelineProjection)
	model.CorrectionReadProjectionState = EvaluatePoint14ValDCorrectionReadProjectionState(model.CorrectionReadProjection)
	model.GovernanceTraceProjectionState = EvaluatePoint14ValDGovernanceTraceProjectionState(model.GovernanceTraceProjection)
	model.QueryProjectionState = EvaluatePoint14ValDQueryProjectionState(model.QueryProjection)
	model.AccessBoundaryState = EvaluatePoint14ValDAccessBoundaryState(model.AccessBoundary, model.Dependency)
	model.TenantPrivacyTimelineState = EvaluatePoint14ValDTenantPrivacyTimelineProjectionGuardState(model.TenantPrivacyTimelineProjectionGuard, model.Dependency)
	model.AgentTimelineProjectionState = EvaluatePoint14ValDAgentTimelineProjectionState(model.AgentTimelineProjection, model.Dependency)
	model.TimestampIntegrityState = EvaluatePoint14ValDTimestampIntegrityGuardState(model.TimestampIntegrityGuard)
	model.NoMutationProjectionGuardState = EvaluatePoint14ValDNoMutationProjectionGuardState(model.NoMutationProjectionGuard)
	model.NoOverclaimTimelineWordingState = EvaluatePoint14ValDNoOverclaimTimelineWordingState(model.NoOverclaimTimelineWording)
	model.CurrentState = point14ValDFoundationState(
		model.DependencyState,
		model.TimelineProjectionState,
		model.SignalTimelineEntryState,
		model.DisputeTimelineState,
		model.CorrectionReadProjectionState,
		model.GovernanceTraceProjectionState,
		model.QueryProjectionState,
		model.AccessBoundaryState,
		model.TenantPrivacyTimelineState,
		model.AgentTimelineProjectionState,
		model.TimestampIntegrityState,
		model.NoMutationProjectionGuardState,
		model.NoOverclaimTimelineWordingState,
	)
	model.BlockingReasons = point14ValDBlockingReasons(model)
	model.ReviewPrerequisites = point14ValDReviewPrerequisites(model)
	return model
}

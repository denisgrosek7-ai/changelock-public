package formal

import (
	"sort"
	"strings"

	"github.com/denisgrosek/changelock/internal/operability"
)

const (
	Point14ValCStateActive         = "point14_valc_external_correction_publication_active"
	Point14ValCStateBlocked        = "point14_valc_external_correction_publication_blocked"
	Point14ValCStateReviewRequired = "point14_valc_external_correction_publication_review_required"
	Point14ValCStateIncomplete     = "point14_valc_external_correction_publication_incomplete"
)

const (
	point14ValCWaveID                       = "val_c"
	point14ValCPublicationDisclaimerBase    = "projection_only not_canonical_truth point14_valc_external_correction_publication_boundary"
	point14ValCCorrectionDraft              = "correction_draft"
	point14ValCCorrectionReviewRequired     = "correction_review_required"
	point14ValCCorrectionEvidenceRequired   = "correction_evidence_required"
	point14ValCCorrectionApprovedBounded    = "correction_approved_for_bounded_publication"
	point14ValCCorrectionBlocked            = "correction_blocked"
	point14ValCRevocationRequested          = "revocation_requested"
	point14ValCRevocationReviewRequired     = "revocation_review_required"
	point14ValCRevocationEvidenceRequired   = "revocation_evidence_required"
	point14ValCRevocationApprovedGovernance = "revocation_approved_for_governance_action"
	point14ValCRevocationBlocked            = "revocation_blocked"
	point14ValCSupersessionRecorded         = "supersession_recorded"
	point14ValCSupersessionReviewRequired   = "supersession_review_required"
	point14ValCSupersessionEvidenceRequired = "supersession_evidence_required"
	point14ValCSupersessionBlocked          = "supersession_blocked"
	point14ValCPublicationNotRequested      = "publication_not_requested"
	point14ValCPublicationReviewRequired    = "publication_review_required"
	point14ValCPublicationApprovedBounded   = "publication_approved_bounded"
	point14ValCPublicationBlocked           = "publication_blocked"
	point14ValCPublicationPrivateOnly       = "publication_private_only"
	point14ValCVisibilityPrivateTenantOnly  = "private_tenant_only"
	point14ValCVisibilityCustomerBounded    = "customer_shared_bounded"
	point14ValCVisibilityAuditorBounded     = "auditor_shared_bounded"
	point14ValCVisibilityPublicBounded      = "public_notice_bounded"
	point14ValCVisibilityBlocked            = "publication_blocked"
	point14ValCPublicationBoundaryPrivate   = "tenant_private"
	point14ValCPublicationBoundaryCustomer  = "bounded_customer_shared"
	point14ValCPublicationBoundaryAuditor   = "bounded_auditor_shared"
	point14ValCPublicationBoundaryPublic    = "bounded_public_notice"
)

type Point14ValCDependencySnapshot struct {
	Point14ValBCurrentState               string                                          `json:"point14_valb_current_state"`
	Point14ValBDependencyState            string                                          `json:"point14_valb_dependency_state"`
	Point14ValBConflictSetState           string                                          `json:"point14_valb_conflict_set_state"`
	Point14ValBStakeholderComparisonState string                                          `json:"point14_valb_stakeholder_comparison_state"`
	Point14ValBDisputeTriageResultState   string                                          `json:"point14_valb_dispute_triage_result_state"`
	Point14ValBDisputeIntakeState         string                                          `json:"point14_valb_dispute_intake_state"`
	Point14ValBEvidenceRequirementState   string                                          `json:"point14_valb_evidence_requirement_state"`
	Point14ValBGovernanceEscalationState  string                                          `json:"point14_valb_governance_escalation_state"`
	Point14ValBTenantPrivacyBoundaryState string                                          `json:"point14_valb_tenant_privacy_boundary_state"`
	Point14ValBAgentDisputeBoundaryState  string                                          `json:"point14_valb_agent_dispute_boundary_state"`
	Point14ValBNoExternalAuthorityState   string                                          `json:"point14_valb_no_external_authority_state"`
	Point14ValBNoOverclaimState           string                                          `json:"point14_valb_no_overclaim_state"`
	Point14ValBPointID                    string                                          `json:"point14_valb_point_id"`
	Point14ValBWaveID                     string                                          `json:"point14_valb_wave_id"`
	Point14ValBComputedFromUpstream       bool                                            `json:"point14_valb_computed_from_upstream"`
	Point14ValBMerged                     bool                                            `json:"point14_valb_merged"`
	Point14ValBCIGreen                    bool                                            `json:"point14_valb_ci_green"`
	Point14ValBReviewedOnMain             bool                                            `json:"point14_valb_reviewed_on_main"`
	Point14PassSeen                       bool                                            `json:"point14_pass_seen"`
	InheritedPoint14ValACurrentState      string                                          `json:"inherited_point14_vala_current_state"`
	InheritedPoint14Val0CurrentState      string                                          `json:"inherited_point14_val0_current_state"`
	InheritedPoint13ValECurrentState      string                                          `json:"inherited_point13_vale_current_state"`
	InheritedPoint13ValEPassClosureState  string                                          `json:"inherited_point13_vale_pass_closure_state"`
	InheritedPoint13ValEPassAllowed       bool                                            `json:"inherited_point13_vale_pass_allowed"`
	InheritedPoint13ValEPassToken         string                                          `json:"inherited_point13_vale_pass_token"`
	InheritedPoint12CurrentState          string                                          `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState       string                                          `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureState      string                                          `json:"inherited_point12_pass_closure_state"`
	InheritedPoint12ReviewerResult        string                                          `json:"inherited_point12_reviewer_result"`
	InheritedPoint11PublicationState      string                                          `json:"inherited_point11_publication_state"`
	InheritedPoint11NoOverclaimState      string                                          `json:"inherited_point11_no_overclaim_state"`
	InheritedPoint11FinalPassGateState    string                                          `json:"inherited_point11_final_pass_gate_state"`
	InheritedPoint10CurrentState          string                                          `json:"inherited_point10_current_state"`
	InheritedPoint10NoOverclaimState      string                                          `json:"inherited_point10_no_overclaim_state"`
	InheritedPoint10ProjectionState       string                                          `json:"inherited_point10_projection_state"`
	InheritedPoint10PassRuleState         string                                          `json:"inherited_point10_pass_rule_state"`
	InheritedTenantScope                  string                                          `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput            bool                                            `json:"snapshot_from_computed_output"`
	Point14ValB                           Point14ValBFoundation                           `json:"point14_valb"`
	Point14ValA                           Point14ValAFoundation                           `json:"point14_vala"`
	Point14Val0                           Point14Val0Foundation                           `json:"point14_val0"`
	Point13ValE                           Point13ValEFoundation                           `json:"point13_vale"`
	Point12                               Point12ValEFoundation                           `json:"point12"`
	Point11                               Point11ValDFoundation                           `json:"point11"`
	Point10                               operability.DeploymentMultiTenantValEFoundation `json:"point10"`
}

type ExternalCorrectionNoticeBoundary struct {
	CorrectionNoticeID         string   `json:"correction_notice_id"`
	DisputeRef                 string   `json:"dispute_ref"`
	ConflictSetRef             string   `json:"conflict_set_ref"`
	CorrectedSignalRefs        []string `json:"corrected_signal_refs,omitempty"`
	AffectedArtifactRefs       []string `json:"affected_artifact_refs,omitempty"`
	ArtifactScoped             bool     `json:"artifact_scoped"`
	AffectedClaimRefs          []string `json:"affected_claim_refs,omitempty"`
	ClaimScoped                bool     `json:"claim_scoped"`
	AffectedEvidenceRefs       []string `json:"affected_evidence_refs,omitempty"`
	CorrectionReason           string   `json:"correction_reason"`
	CorrectionLimitations      []string `json:"correction_limitations,omitempty"`
	GovernanceEventRef         string   `json:"governance_event_ref"`
	AuditRef                   string   `json:"audit_ref"`
	Owner                      string   `json:"owner"`
	CorrectionState            string   `json:"correction_state"`
	CanonicalMutationRequested bool     `json:"canonical_mutation_requested"`
	OverridesCanonicalDecision bool     `json:"overrides_canonical_decision"`
	CertifiesCompliance        bool     `json:"certifies_compliance"`
	ApprovesProduction         bool     `json:"approves_production"`
	CreatesPublicBadge         bool     `json:"creates_public_badge"`
	EmitsPass                  bool     `json:"emits_pass"`
	StrengthensClaim           bool     `json:"strengthens_claim"`
}

type ExternalRevocationRequestBoundary struct {
	RevocationRequestID        string   `json:"revocation_request_id"`
	DisputeRef                 string   `json:"dispute_ref"`
	TargetClaimRefs            []string `json:"target_claim_refs,omitempty"`
	TargetSignalRefs           []string `json:"target_signal_refs,omitempty"`
	RevocationReason           string   `json:"revocation_reason"`
	EvidenceRefs               []string `json:"evidence_refs,omitempty"`
	GovernanceEventRef         string   `json:"governance_event_ref"`
	AuditRef                   string   `json:"audit_ref"`
	Owner                      string   `json:"owner"`
	RevocationState            string   `json:"revocation_state"`
	CanonicalMutationRequested bool     `json:"canonical_mutation_requested"`
	ExternalAuthorityGranted   bool     `json:"external_authority_granted"`
	PublicBadgeAllowed         bool     `json:"public_badge_allowed"`
	PassAllowed                bool     `json:"pass_allowed"`
}

type ExternalSupersessionRecordBoundary struct {
	SupersessionRecordID      string   `json:"supersession_record_id"`
	PreviousSignalRef         string   `json:"previous_signal_ref"`
	ReplacementSignalRef      string   `json:"replacement_signal_ref"`
	SupersessionReason        string   `json:"supersession_reason"`
	EvidenceRefs              []string `json:"evidence_refs,omitempty"`
	GovernanceEventRef        string   `json:"governance_event_ref"`
	AuditRef                  string   `json:"audit_ref"`
	Owner                     string   `json:"owner"`
	SupersessionState         string   `json:"supersession_state"`
	SilentReplacement         bool     `json:"silent_replacement"`
	HidesPreviousEvidence     bool     `json:"hides_previous_evidence"`
	DeletesHistory            bool     `json:"deletes_history"`
	ReplacementHashRecomputed bool     `json:"replacement_hash_recomputed"`
}

type ExternalPublicationApprovalBoundary struct {
	PublicationApprovalID      string `json:"publication_approval_id"`
	CorrectionNoticeRef        string `json:"correction_notice_ref"`
	RevocationRequestRef       string `json:"revocation_request_ref"`
	SupersessionRecordRef      string `json:"supersession_record_ref"`
	ApproverRole               string `json:"approver_role"`
	ApproverRef                string `json:"approver_ref"`
	ApprovalReason             string `json:"approval_reason"`
	ApprovalScope              string `json:"approval_scope"`
	AuditRef                   string `json:"audit_ref"`
	GovernanceEventRef         string `json:"governance_event_ref"`
	RequestedAt                string `json:"requested_at"`
	RequestedTimeSource        string `json:"requested_time_source"`
	ApprovedAt                 string `json:"approved_at"`
	ApprovedTimeSource         string `json:"approved_time_source"`
	RecordedAt                 string `json:"recorded_at"`
	RecordedTimeSource         string `json:"recorded_time_source"`
	PublicationState           string `json:"publication_state"`
	ApprovesProduction         bool   `json:"approves_production"`
	Certifies                  bool   `json:"certifies"`
	CreatesPass                bool   `json:"creates_pass"`
	CreatesPublicBadge         bool   `json:"creates_public_badge"`
	AutomaticApprovalRequested bool   `json:"automatic_approval_requested"`
	GlobalTruthRequested       bool   `json:"global_truth_requested"`
}

type ExternalPublicationVisibilityBoundary struct {
	VisibilityBoundaryID                    string   `json:"visibility_boundary_id"`
	VisibilityClassification                string   `json:"visibility_classification"`
	TenantScope                             string   `json:"tenant_scope"`
	GlobalScopeClassification               string   `json:"global_scope_classification"`
	PublicPrivateBoundary                   string   `json:"public_private_boundary"`
	LimitationRefs                          []string `json:"limitation_refs,omitempty"`
	RedactionRefs                           []string `json:"redaction_refs,omitempty"`
	PublicationTextRefs                     []string `json:"publication_text_refs,omitempty"`
	TenantPrivateDataExposed                bool     `json:"tenant_private_data_exposed"`
	StrengthensClaim                        bool     `json:"strengthens_claim"`
	ImpliesPublicAuthority                  bool     `json:"implies_public_authority"`
	MeaningChangingPrivateLimitationOmitted bool     `json:"meaning_changing_private_limitation_omitted"`
}

type TenantPrivacyPublicationGuard struct {
	BoundaryID                  string   `json:"boundary_id"`
	TenantScope                 string   `json:"tenant_scope"`
	PublicationTargetScope      string   `json:"publication_target_scope"`
	GlobalScopeClassification   string   `json:"global_scope_classification"`
	PublicPrivateClassification string   `json:"public_private_classification"`
	RedactionRefs               []string `json:"redaction_refs,omitempty"`
	LimitationRefs              []string `json:"limitation_refs,omitempty"`
	TenantPrivateDataExposed    bool     `json:"tenant_private_data_exposed"`
	CrossTenantPublication      bool     `json:"cross_tenant_publication"`
	LimitationsVisible          bool     `json:"limitations_visible"`
}

type CorrectionRedactionLimitationGuard struct {
	GuardID                       string   `json:"guard_id"`
	Redacted                      bool     `json:"redacted"`
	RedactionManifestRef          string   `json:"redaction_manifest_ref"`
	RedactionRefs                 []string `json:"redaction_refs,omitempty"`
	LimitationRefs                []string `json:"limitation_refs,omitempty"`
	DecisiveMissingEvidenceHidden bool     `json:"decisive_missing_evidence_hidden"`
	RedactionStrengthensClaim     bool     `json:"redaction_strengthens_claim"`
	LimitationOmitted             bool     `json:"limitation_omitted"`
	SurvivingTextMisleading       bool     `json:"surviving_text_misleading"`
	SourcePublicationState        string   `json:"source_publication_state"`
	ResolvedPublicationState      string   `json:"resolved_publication_state"`
}

type CorrectionPublicationGovernanceTrace struct {
	GovernanceTraceID       string   `json:"governance_trace_id"`
	GovernanceEventRef      string   `json:"governance_event_ref"`
	Owner                   string   `json:"owner"`
	ApproverRole            string   `json:"approver_role"`
	AuditRef                string   `json:"audit_ref"`
	EvidenceRefs            []string `json:"evidence_refs,omitempty"`
	DecisionReason          string   `json:"decision_reason"`
	Timestamp               string   `json:"timestamp"`
	TimeSource              string   `json:"time_source"`
	DisputeOpenedAt         string   `json:"dispute_opened_at"`
	DisputeOpenedTimeSource string   `json:"dispute_opened_time_source"`
	ApprovesProduction      bool     `json:"approves_production"`
	CertifiesCompliance     bool     `json:"certifies_compliance"`
}

type AgentCorrectionPublicationBoundary struct {
	BoundaryID                     string   `json:"boundary_id"`
	TenantScope                    string   `json:"tenant_scope"`
	AgentInputRefs                 []string `json:"agent_input_refs,omitempty"`
	EvidenceRefs                   []string `json:"evidence_refs,omitempty"`
	AuditEventRef                  string   `json:"audit_event_ref"`
	AdvisoryOnly                   bool     `json:"advisory_only"`
	CanApproveCorrection           bool     `json:"can_approve_correction"`
	CanApproveRevocation           bool     `json:"can_approve_revocation"`
	CanApprovePublication          bool     `json:"can_approve_publication"`
	CanPublishNotice               bool     `json:"can_publish_notice"`
	CanSatisfyGovernanceTraceAlone bool     `json:"can_satisfy_governance_trace_alone"`
	CanEmitPass                    bool     `json:"can_emit_pass"`
	CanEmitPublicAuthority         bool     `json:"can_emit_public_authority"`
	ApprovalGranted                bool     `json:"approval_granted"`
	ProductionApproved             bool     `json:"production_approved"`
	ExternalAuthorityAllowed       bool     `json:"external_authority_allowed"`
}

type Point14ValCNoExternalAuthorityPublicationGuard struct {
	ObservedAuthorityMarkers   []string `json:"observed_authority_markers,omitempty"`
	CanonicalAuthorityGranted  bool     `json:"canonical_authority_granted"`
	ProductionApprovalGranted  bool     `json:"production_approval_granted"`
	PublishesCorrection        bool     `json:"publishes_correction"`
	RevokesClaim               bool     `json:"revokes_claim"`
	OverridesCanonicalDecision bool     `json:"overrides_canonical_decision"`
	CorrectionPublished        bool     `json:"correction_published"`
	PublicationApproved        bool     `json:"publication_approved"`
	PublicBadgeAllowed         bool     `json:"public_badge_allowed"`
	CorrectionAutoPublished    bool     `json:"correction_auto_published"`
	RevocationAutoExecuted     bool     `json:"revocation_auto_executed"`
	PublicationAutoApproved    bool     `json:"publication_auto_approved"`
	CrowdApprovedPublication   bool     `json:"crowd_approved_publication"`
	AgentApprovedPublication   bool     `json:"agent_approved_publication"`
	PassAllowed                bool     `json:"pass_allowed"`
	ExternalAuthorityAllowed   bool     `json:"external_authority_allowed"`
}

type Point14ValCNoOverclaimPublicationWording struct {
	ObservedCorrectionTexts              []string `json:"observed_correction_texts,omitempty"`
	ObservedRevocationTexts              []string `json:"observed_revocation_texts,omitempty"`
	ObservedSupersessionTexts            []string `json:"observed_supersession_texts,omitempty"`
	ObservedPublicationTexts             []string `json:"observed_publication_texts,omitempty"`
	ObservedVisibilityTexts              []string `json:"observed_visibility_texts,omitempty"`
	ObservedPrivacyTexts                 []string `json:"observed_privacy_texts,omitempty"`
	ObservedGovernanceTexts              []string `json:"observed_governance_texts,omitempty"`
	ObservedAgentTexts                   []string `json:"observed_agent_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point14ValCFoundation struct {
	CurrentState                  string                                         `json:"current_state"`
	BlockingReasons               []string                                       `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites           []string                                       `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer          string                                         `json:"projection_disclaimer"`
	DependencyState               string                                         `json:"dependency_state"`
	CorrectionNoticeState         string                                         `json:"correction_notice_state"`
	RevocationRequestState        string                                         `json:"revocation_request_state"`
	SupersessionRecordState       string                                         `json:"supersession_record_state"`
	PublicationApprovalState      string                                         `json:"publication_approval_state"`
	VisibilityBoundaryState       string                                         `json:"visibility_boundary_state"`
	TenantPrivacyState            string                                         `json:"tenant_privacy_state"`
	RedactionLimitationState      string                                         `json:"redaction_limitation_state"`
	GovernanceTraceState          string                                         `json:"governance_trace_state"`
	AgentPublicationBoundaryState string                                         `json:"agent_publication_boundary_state"`
	NoExternalAuthorityState      string                                         `json:"no_external_authority_state"`
	NoOverclaimState              string                                         `json:"no_overclaim_state"`
	Dependency                    Point14ValCDependencySnapshot                  `json:"dependency"`
	CorrectionNoticeBoundary      ExternalCorrectionNoticeBoundary               `json:"correction_notice_boundary"`
	RevocationRequestBoundary     ExternalRevocationRequestBoundary              `json:"revocation_request_boundary"`
	SupersessionRecordBoundary    ExternalSupersessionRecordBoundary             `json:"supersession_record_boundary"`
	PublicationApprovalBoundary   ExternalPublicationApprovalBoundary            `json:"publication_approval_boundary"`
	PublicationVisibilityBoundary ExternalPublicationVisibilityBoundary          `json:"publication_visibility_boundary"`
	TenantPrivacyPublicationGuard TenantPrivacyPublicationGuard                  `json:"tenant_privacy_publication_guard"`
	CorrectionRedactionGuard      CorrectionRedactionLimitationGuard             `json:"correction_redaction_guard"`
	GovernanceTrace               CorrectionPublicationGovernanceTrace           `json:"governance_trace"`
	AgentCorrectionBoundary       AgentCorrectionPublicationBoundary             `json:"agent_correction_boundary"`
	NoExternalAuthorityGuard      Point14ValCNoExternalAuthorityPublicationGuard `json:"no_external_authority_guard"`
	NoOverclaimPublicationWording Point14ValCNoOverclaimPublicationWording       `json:"no_overclaim_publication_wording"`
}

func point14ValCStates() []string {
	return []string{
		Point14ValCStateActive,
		Point14ValCStateBlocked,
		Point14ValCStateReviewRequired,
		Point14ValCStateIncomplete,
	}
}

func point14ValCStateValid(value string) bool {
	return point11Val0ContainsTrimmed(point14ValCStates(), value)
}

func point14ValCCorrectionStates() []string {
	return []string{
		point14ValCCorrectionDraft,
		point14ValCCorrectionReviewRequired,
		point14ValCCorrectionEvidenceRequired,
		point14ValCCorrectionApprovedBounded,
		point14ValCCorrectionBlocked,
	}
}

func point14ValCRevocationStates() []string {
	return []string{
		point14ValCRevocationRequested,
		point14ValCRevocationReviewRequired,
		point14ValCRevocationEvidenceRequired,
		point14ValCRevocationApprovedGovernance,
		point14ValCRevocationBlocked,
	}
}

func point14ValCSupersessionStates() []string {
	return []string{
		point14ValCSupersessionRecorded,
		point14ValCSupersessionReviewRequired,
		point14ValCSupersessionEvidenceRequired,
		point14ValCSupersessionBlocked,
	}
}

func point14ValCPublicationStates() []string {
	return []string{
		point14ValCPublicationNotRequested,
		point14ValCPublicationReviewRequired,
		point14ValCPublicationApprovedBounded,
		point14ValCPublicationBlocked,
		point14ValCPublicationPrivateOnly,
	}
}

func point14ValCVisibilityClassifications() []string {
	return []string{
		point14ValCVisibilityPrivateTenantOnly,
		point14ValCVisibilityCustomerBounded,
		point14ValCVisibilityAuditorBounded,
		point14ValCVisibilityPublicBounded,
		point14ValCVisibilityBlocked,
	}
}

func point14ValCPublicPrivateBoundaries() []string {
	return []string{
		point14ValCPublicationBoundaryPrivate,
		point14ValCPublicationBoundaryCustomer,
		point14ValCPublicationBoundaryAuditor,
		point14ValCPublicationBoundaryPublic,
	}
}

func point14ValCApprovalScopes() []string {
	return []string{
		point14ValCPublicationBoundaryPrivate,
		point14ValCPublicationBoundaryCustomer,
		point14ValCPublicationBoundaryAuditor,
		point14ValCPublicationBoundaryPublic,
	}
}

func point14ValCPublicationBoundaryPairValid(scope, boundary string) bool {
	switch strings.TrimSpace(scope) {
	case point14ValCVisibilityPrivateTenantOnly:
		return strings.TrimSpace(boundary) == point14ValCPublicationBoundaryPrivate
	case point14ValCVisibilityCustomerBounded:
		return strings.TrimSpace(boundary) == point14ValCPublicationBoundaryCustomer
	case point14ValCVisibilityAuditorBounded:
		return strings.TrimSpace(boundary) == point14ValCPublicationBoundaryAuditor
	case point14ValCVisibilityPublicBounded:
		return strings.TrimSpace(boundary) == point14ValCPublicationBoundaryPublic
	default:
		return false
	}
}

func point14ValCForbiddenAuthorityMarkers() []string {
	markers := append([]string{}, point14ValBForbiddenAuthorityMarkers()...)
	return append(markers,
		"correction_auto_published",
		"correction_certified",
		"correction_production_approved",
		"revocation_auto_executed",
		"publication_auto_approved",
		"publication_production_approved",
		"publication_certified",
		"publication_public_badge",
		"publication_global_truth",
		"crowd_approved_publication",
		"agent_approved_publication",
	)
}

func point14ValCSafeWording() []string {
	return []string{
		"bounded correction notice",
		"correction pending governance approval",
		"revocation request pending governance review",
		"supersession record preserves prior context",
		"bounded public notice",
		"private tenant-scoped correction",
		"publication approval is not production approval",
		"no public badge authority granted",
		"canonical evidence spine remains source of truth",
		"limitations remain visible",
	}
}

func point14ValCForbiddenWording() []string {
	return append(
		append([]string{}, point14Val0ForbiddenWording()...),
		"correction certified",
		"correction proves compliance",
		"publication proves safety",
		"revoked by external authority",
		"ai-approved correction",
		"crowd-approved correction",
	)
}

func point14ValCObservedTextContainsForbiddenWording(text string) bool {
	trimmed := strings.ToLower(strings.TrimSpace(text))
	if trimmed == "" {
		return false
	}
	for _, safe := range point14ValCSafeWording() {
		if trimmed == strings.ToLower(strings.TrimSpace(safe)) {
			return false
		}
	}
	for _, phrase := range point14ValCForbiddenWording() {
		if strings.Contains(trimmed, strings.ToLower(strings.TrimSpace(phrase))) {
			return true
		}
	}
	return false
}

func point14ValCObservedListContainsForbiddenWording(values []string) bool {
	for _, value := range values {
		if point14ValCObservedTextContainsForbiddenWording(value) {
			return true
		}
	}
	return false
}

func point14ValCCorrectionNoticeIDValid(value string) bool {
	return point14Val0RefValid(value, "correction_notice_")
}

func point14ValCRevocationRequestIDValid(value string) bool {
	return point14Val0RefValid(value, "revocation_request_")
}

func point14ValCSupersessionRecordIDValid(value string) bool {
	return point14Val0RefValid(value, "supersession_record_")
}

func point14ValCPublicationApprovalIDValid(value string) bool {
	return point14Val0RefValid(value, "publication_approval_")
}

func point14ValCGovernanceTraceIDValid(value string) bool {
	return point14Val0RefValid(value, "governance_trace_")
}

func point14ValCRedactionManifestRefValid(value string) bool {
	return point14Val0RefValid(value, "redaction_manifest_")
}

func point14ValCLimitationRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "limitation_ref_")
}

func point14ValCRedactionRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "redaction_ref_")
}

func point14ValCPublicationTextRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "publication_text_")
}

func point14ValCCorrectedSignalRefsValid(values []string) bool {
	return point14Val0RefListValid(values, "normalized_signal_", "signal_")
}

func point14ValCExactValueValid(value string, allowed []string) bool {
	return point11Val0ContainsTrimmed(allowed, value)
}

func point14ValCDependencySnapshotFromUpstream(valB Point14ValBFoundation) Point14ValCDependencySnapshot {
	return Point14ValCDependencySnapshot{
		Point14ValBCurrentState:               valB.CurrentState,
		Point14ValBDependencyState:            valB.DependencyState,
		Point14ValBConflictSetState:           valB.ConflictSetState,
		Point14ValBStakeholderComparisonState: valB.StakeholderComparisonState,
		Point14ValBDisputeTriageResultState:   valB.DisputeTriageResultState,
		Point14ValBDisputeIntakeState:         valB.DisputeIntakeState,
		Point14ValBEvidenceRequirementState:   valB.EvidenceRequirementState,
		Point14ValBGovernanceEscalationState:  valB.GovernanceEscalationState,
		Point14ValBTenantPrivacyBoundaryState: valB.TenantPrivacyBoundaryState,
		Point14ValBAgentDisputeBoundaryState:  valB.AgentDisputeBoundaryState,
		Point14ValBNoExternalAuthorityState:   valB.NoExternalAuthorityState,
		Point14ValBNoOverclaimState:           valB.NoOverclaimState,
		Point14ValBPointID:                    point14Val0PointID,
		Point14ValBWaveID:                     point14ValBWaveID,
		Point14ValBComputedFromUpstream:       valB.Dependency.SnapshotFromComputedOutput,
		Point14ValBMerged:                     true,
		Point14ValBCIGreen:                    true,
		Point14ValBReviewedOnMain:             true,
		Point14PassSeen:                       false,
		InheritedPoint14ValACurrentState:      valB.Dependency.Point14ValACurrentState,
		InheritedPoint14Val0CurrentState:      valB.Dependency.InheritedPoint14Val0CurrentState,
		InheritedPoint13ValECurrentState:      valB.Dependency.InheritedPoint13ValECurrentState,
		InheritedPoint13ValEPassClosureState:  valB.Dependency.InheritedPoint13ValEPassClosureState,
		InheritedPoint13ValEPassAllowed:       valB.Dependency.InheritedPoint13ValEPassAllowed,
		InheritedPoint13ValEPassToken:         valB.Dependency.InheritedPoint13ValEPassToken,
		InheritedPoint12CurrentState:          valB.Dependency.InheritedPoint12CurrentState,
		InheritedPoint12DependencyState:       valB.Dependency.InheritedPoint12DependencyState,
		InheritedPoint12PassClosureState:      valB.Dependency.InheritedPoint12PassClosureState,
		InheritedPoint12ReviewerResult:        valB.Dependency.InheritedPoint12ReviewerResult,
		InheritedPoint11PublicationState:      valB.Dependency.InheritedPoint11PublicationState,
		InheritedPoint11NoOverclaimState:      valB.Dependency.InheritedPoint11NoOverclaimState,
		InheritedPoint11FinalPassGateState:    valB.Dependency.InheritedPoint11FinalPassGateState,
		InheritedPoint10CurrentState:          valB.Dependency.InheritedPoint10CurrentState,
		InheritedPoint10NoOverclaimState:      valB.Dependency.InheritedPoint10NoOverclaimState,
		InheritedPoint10ProjectionState:       valB.Dependency.InheritedPoint10ProjectionState,
		InheritedPoint10PassRuleState:         valB.Dependency.InheritedPoint10PassRuleState,
		InheritedTenantScope:                  valB.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:            true,
		Point14ValB:                           valB,
		Point14ValA:                           valB.Dependency.Point14ValA,
		Point14Val0:                           valB.Dependency.Point14Val0,
		Point13ValE:                           valB.Dependency.Point13ValE,
		Point12:                               valB.Dependency.Point12,
		Point11:                               valB.Dependency.Point11,
		Point10:                               valB.Dependency.Point10,
	}
}

func point14ValCDependencySnapshotModel() Point14ValCDependencySnapshot {
	return cachedFormalModel(&point14ValCDependencySnapshotModelOnce, &point14ValCDependencySnapshotModelCached, func() Point14ValCDependencySnapshot {
		valB := ComputePoint14ValBFoundation(Point14ValBFoundationModel())
		return point14ValCDependencySnapshotFromUpstream(valB)
	})
}

func EvaluatePoint14ValCDependencyState(model Point14ValCDependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point14ValBComputedFromUpstream ||
		!model.Point14ValBMerged ||
		!model.Point14ValBCIGreen ||
		!model.Point14ValBReviewedOnMain ||
		model.Point14PassSeen ||
		model.Point14ValBPointID != point14Val0PointID ||
		model.Point14ValBWaveID != point14ValBWaveID ||
		!point14ValBStateValid(model.Point14ValBCurrentState) ||
		!point14ValBStateValid(model.Point14ValBDependencyState) ||
		!point14ValBStateValid(model.Point14ValBConflictSetState) ||
		!point14ValBStateValid(model.Point14ValBStakeholderComparisonState) ||
		!point14ValBStateValid(model.Point14ValBDisputeTriageResultState) ||
		!point14ValBStateValid(model.Point14ValBDisputeIntakeState) ||
		!point14ValBStateValid(model.Point14ValBEvidenceRequirementState) ||
		!point14ValBStateValid(model.Point14ValBGovernanceEscalationState) ||
		!point14ValBStateValid(model.Point14ValBTenantPrivacyBoundaryState) ||
		!point14ValBStateValid(model.Point14ValBAgentDisputeBoundaryState) ||
		!point14ValBStateValid(model.Point14ValBNoExternalAuthorityState) ||
		!point14ValBStateValid(model.Point14ValBNoOverclaimState) ||
		!point14ValAStateValid(model.InheritedPoint14ValACurrentState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0CurrentState) ||
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
		return Point14ValCStateBlocked
	}
	if model.Point14ValBCurrentState != model.Point14ValB.CurrentState ||
		model.Point14ValBDependencyState != model.Point14ValB.DependencyState ||
		model.Point14ValBConflictSetState != model.Point14ValB.ConflictSetState ||
		model.Point14ValBStakeholderComparisonState != model.Point14ValB.StakeholderComparisonState ||
		model.Point14ValBDisputeTriageResultState != model.Point14ValB.DisputeTriageResultState ||
		model.Point14ValBDisputeIntakeState != model.Point14ValB.DisputeIntakeState ||
		model.Point14ValBEvidenceRequirementState != model.Point14ValB.EvidenceRequirementState ||
		model.Point14ValBGovernanceEscalationState != model.Point14ValB.GovernanceEscalationState ||
		model.Point14ValBTenantPrivacyBoundaryState != model.Point14ValB.TenantPrivacyBoundaryState ||
		model.Point14ValBAgentDisputeBoundaryState != model.Point14ValB.AgentDisputeBoundaryState ||
		model.Point14ValBNoExternalAuthorityState != model.Point14ValB.NoExternalAuthorityState ||
		model.Point14ValBNoOverclaimState != model.Point14ValB.NoOverclaimState ||
		model.Point14ValBComputedFromUpstream != model.Point14ValB.Dependency.SnapshotFromComputedOutput ||
		model.InheritedPoint14ValACurrentState != model.Point14ValB.Dependency.Point14ValACurrentState ||
		model.InheritedPoint14Val0CurrentState != model.Point14ValB.Dependency.InheritedPoint14Val0CurrentState ||
		model.InheritedPoint13ValECurrentState != model.Point14ValB.Dependency.InheritedPoint13ValECurrentState ||
		model.InheritedPoint13ValEPassClosureState != model.Point14ValB.Dependency.InheritedPoint13ValEPassClosureState ||
		model.InheritedPoint13ValEPassAllowed != model.Point14ValB.Dependency.InheritedPoint13ValEPassAllowed ||
		model.InheritedPoint13ValEPassToken != model.Point14ValB.Dependency.InheritedPoint13ValEPassToken ||
		model.InheritedPoint12CurrentState != model.Point14ValB.Dependency.InheritedPoint12CurrentState ||
		model.InheritedPoint12DependencyState != model.Point14ValB.Dependency.InheritedPoint12DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point14ValB.Dependency.InheritedPoint12PassClosureState ||
		model.InheritedPoint12ReviewerResult != model.Point14ValB.Dependency.InheritedPoint12ReviewerResult ||
		model.InheritedPoint11PublicationState != model.Point14ValB.Dependency.InheritedPoint11PublicationState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValB.Dependency.InheritedPoint11NoOverclaimState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValB.Dependency.InheritedPoint11FinalPassGateState ||
		model.InheritedPoint10CurrentState != model.Point14ValB.Dependency.InheritedPoint10CurrentState ||
		model.InheritedPoint10NoOverclaimState != model.Point14ValB.Dependency.InheritedPoint10NoOverclaimState ||
		model.InheritedPoint10ProjectionState != model.Point14ValB.Dependency.InheritedPoint10ProjectionState ||
		model.InheritedPoint10PassRuleState != model.Point14ValB.Dependency.InheritedPoint10PassRuleState ||
		model.InheritedTenantScope != model.Point14ValB.Dependency.InheritedTenantScope {
		return Point14ValCStateBlocked
	}
	if model.Point14ValBCurrentState != Point14ValBStateActive ||
		model.Point14ValBDependencyState != Point14ValBStateActive ||
		model.Point14ValBConflictSetState != Point14ValBStateActive ||
		model.Point14ValBStakeholderComparisonState != Point14ValBStateActive ||
		model.Point14ValBDisputeTriageResultState != Point14ValBStateActive ||
		model.Point14ValBDisputeIntakeState != Point14ValBStateActive ||
		model.Point14ValBEvidenceRequirementState != Point14ValBStateActive ||
		model.Point14ValBGovernanceEscalationState != Point14ValBStateActive ||
		model.Point14ValBTenantPrivacyBoundaryState != Point14ValBStateActive ||
		model.Point14ValBAgentDisputeBoundaryState != Point14ValBStateActive ||
		model.Point14ValBNoExternalAuthorityState != Point14ValBStateActive ||
		model.Point14ValBNoOverclaimState != Point14ValBStateActive ||
		model.InheritedPoint14ValACurrentState != Point14ValAStateActive ||
		model.InheritedPoint14Val0CurrentState != Point14Val0StateActive ||
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
		return Point14ValCStateBlocked
	}
	return Point14ValCStateActive
}

func point14ValCCorrectionNoticeBoundaryModel() ExternalCorrectionNoticeBoundary {
	return ExternalCorrectionNoticeBoundary{
		CorrectionNoticeID:    "correction_notice_point14_valc_001",
		DisputeRef:            "dispute_point14_valb_001",
		ConflictSetRef:        "conflict_set_point14_valb_001",
		CorrectedSignalRefs:   []string{"normalized_signal_point14_vala_001"},
		AffectedArtifactRefs:  []string{"artifact_point14_valc_component_001"},
		ArtifactScoped:        true,
		AffectedClaimRefs:     []string{"claim_point14_valc_001"},
		ClaimScoped:           true,
		AffectedEvidenceRefs:  []string{"evidence_point14_valc_001"},
		CorrectionReason:      "bounded correction notice",
		CorrectionLimitations: []string{"limitation_ref_point14_valc_001"},
		GovernanceEventRef:    "governance_event_point14_valc_001",
		AuditRef:              "audit_event_point14_valc_001",
		Owner:                 "approver_point14_valc_owner_001",
		CorrectionState:       point14ValCCorrectionApprovedBounded,
	}
}

func EvaluatePoint14ValCCorrectionNoticeBoundaryState(model ExternalCorrectionNoticeBoundary) string {
	if !point14ValCCorrectionNoticeIDValid(model.CorrectionNoticeID) ||
		!point14Val0DisputeIDValid(model.DisputeRef) ||
		!point14ValBConflictSetIDValid(model.ConflictSetRef) ||
		!point14ValCCorrectedSignalRefsValid(model.CorrectedSignalRefs) ||
		!point14Val0EvidenceRefsValid(model.AffectedEvidenceRefs) ||
		strings.TrimSpace(model.CorrectionReason) == "" ||
		!point14ValCLimitationRefsValid(model.CorrectionLimitations) ||
		!point14Val0GovernanceEventRefValid(model.GovernanceEventRef) ||
		!point14Val0AuditEventRefValid(model.AuditRef) ||
		!point14Val0ApproverRefValid(model.Owner) ||
		!point14ValCExactValueValid(model.CorrectionState, point14ValCCorrectionStates()) {
		return Point14ValCStateBlocked
	}
	if model.ArtifactScoped && !point14Val0RefListValid(model.AffectedArtifactRefs, "artifact_") {
		return Point14ValCStateBlocked
	}
	if model.ClaimScoped && !point14Val0ClaimRefsValid(model.AffectedClaimRefs) {
		return Point14ValCStateBlocked
	}
	if model.CanonicalMutationRequested ||
		model.OverridesCanonicalDecision ||
		model.CertifiesCompliance ||
		model.ApprovesProduction ||
		model.CreatesPublicBadge ||
		model.EmitsPass ||
		model.StrengthensClaim {
		return Point14ValCStateBlocked
	}
	switch strings.TrimSpace(model.CorrectionState) {
	case point14ValCCorrectionBlocked:
		return Point14ValCStateBlocked
	case point14ValCCorrectionReviewRequired, point14ValCCorrectionDraft:
		return Point14ValCStateReviewRequired
	case point14ValCCorrectionEvidenceRequired:
		return Point14ValCStateIncomplete
	default:
		return Point14ValCStateActive
	}
}

func point14ValCRevocationRequestBoundaryModel() ExternalRevocationRequestBoundary {
	return ExternalRevocationRequestBoundary{
		RevocationRequestID: "revocation_request_point14_valc_001",
		DisputeRef:          "dispute_point14_valb_001",
		TargetClaimRefs:     []string{"claim_point14_valc_001"},
		RevocationReason:    "revocation request pending governance review",
		EvidenceRefs:        []string{"evidence_point14_valc_001"},
		AuditRef:            "audit_event_point14_valc_revocation_001",
		Owner:               "approver_point14_valc_owner_001",
		RevocationState:     point14ValCRevocationRequested,
	}
}

func EvaluatePoint14ValCRevocationRequestBoundaryState(model ExternalRevocationRequestBoundary) string {
	if !point14ValCRevocationRequestIDValid(model.RevocationRequestID) ||
		!point14Val0DisputeIDValid(model.DisputeRef) ||
		(len(model.TargetClaimRefs) == 0 && len(model.TargetSignalRefs) == 0) ||
		strings.TrimSpace(model.RevocationReason) == "" ||
		!point14Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point14Val0AuditEventRefValid(model.AuditRef) ||
		!point14Val0ApproverRefValid(model.Owner) ||
		!point14ValCExactValueValid(model.RevocationState, point14ValCRevocationStates()) {
		return Point14ValCStateBlocked
	}
	if len(model.TargetClaimRefs) > 0 && !point14Val0ClaimRefsValid(model.TargetClaimRefs) {
		return Point14ValCStateBlocked
	}
	if len(model.TargetSignalRefs) > 0 && !point14ValCCorrectedSignalRefsValid(model.TargetSignalRefs) {
		return Point14ValCStateBlocked
	}
	if model.CanonicalMutationRequested || model.ExternalAuthorityGranted || model.PublicBadgeAllowed || model.PassAllowed {
		return Point14ValCStateBlocked
	}
	if strings.TrimSpace(model.RevocationState) == point14ValCRevocationApprovedGovernance &&
		!point14Val0GovernanceEventRefValid(model.GovernanceEventRef) {
		return Point14ValCStateBlocked
	}
	if strings.TrimSpace(model.RevocationState) == point14ValCRevocationBlocked {
		return Point14ValCStateBlocked
	}
	if strings.TrimSpace(model.RevocationState) == point14ValCRevocationReviewRequired {
		return Point14ValCStateReviewRequired
	}
	if strings.TrimSpace(model.RevocationState) == point14ValCRevocationEvidenceRequired {
		return Point14ValCStateIncomplete
	}
	return Point14ValCStateActive
}

func point14ValCSupersessionRecordBoundaryModel() ExternalSupersessionRecordBoundary {
	return ExternalSupersessionRecordBoundary{
		SupersessionRecordID: "supersession_record_point14_valc_001",
		PreviousSignalRef:    "normalized_signal_point14_vala_001",
		ReplacementSignalRef: "normalized_signal_point14_valc_001",
		SupersessionReason:   "supersession record preserves prior context",
		EvidenceRefs:         []string{"evidence_point14_valc_001"},
		GovernanceEventRef:   "governance_event_point14_valc_001",
		AuditRef:             "audit_event_point14_valc_supersession_001",
		Owner:                "approver_point14_valc_owner_001",
		SupersessionState:    point14ValCSupersessionRecorded,
	}
}

func EvaluatePoint14ValCSupersessionRecordBoundaryState(model ExternalSupersessionRecordBoundary) string {
	if !point14ValCSupersessionRecordIDValid(model.SupersessionRecordID) ||
		!point14Val0RefValid(model.PreviousSignalRef, "normalized_signal_", "signal_") ||
		!point14Val0RefValid(model.ReplacementSignalRef, "normalized_signal_", "signal_") ||
		strings.TrimSpace(model.SupersessionReason) == "" ||
		!point14Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point14Val0GovernanceEventRefValid(model.GovernanceEventRef) ||
		!point14Val0AuditEventRefValid(model.AuditRef) ||
		!point14Val0ApproverRefValid(model.Owner) ||
		!point14ValCExactValueValid(model.SupersessionState, point14ValCSupersessionStates()) {
		return Point14ValCStateBlocked
	}
	if model.SilentReplacement || model.HidesPreviousEvidence || model.DeletesHistory || model.ReplacementHashRecomputed {
		return Point14ValCStateBlocked
	}
	switch strings.TrimSpace(model.SupersessionState) {
	case point14ValCSupersessionBlocked:
		return Point14ValCStateBlocked
	case point14ValCSupersessionReviewRequired:
		return Point14ValCStateReviewRequired
	case point14ValCSupersessionEvidenceRequired:
		return Point14ValCStateIncomplete
	default:
		return Point14ValCStateActive
	}
}

func point14ValCPublicationApprovalBoundaryModel() ExternalPublicationApprovalBoundary {
	return ExternalPublicationApprovalBoundary{
		PublicationApprovalID: "publication_approval_point14_valc_001",
		CorrectionNoticeRef:   "correction_notice_point14_valc_001",
		ApproverRole:          "security_reviewer",
		ApproverRef:           "approver_point14_valc_publication_001",
		ApprovalReason:        "publication approval is not production approval",
		ApprovalScope:         point14ValCPublicationBoundaryPrivate,
		AuditRef:              "audit_event_point14_valc_publication_001",
		GovernanceEventRef:    "governance_event_point14_valc_publication_001",
		RequestedAt:           "2026-05-06T08:00:00Z",
		RequestedTimeSource:   point14Val0TimeSourceServerUTC,
		ApprovedAt:            "2026-05-06T08:05:00Z",
		ApprovedTimeSource:    point14Val0TimeSourceServerUTC,
		RecordedAt:            "2026-05-06T08:06:00Z",
		RecordedTimeSource:    point14Val0TimeSourceServerUTC,
		PublicationState:      point14ValCPublicationPrivateOnly,
	}
}

func EvaluatePoint14ValCPublicationApprovalBoundaryState(model ExternalPublicationApprovalBoundary) string {
	if !point14ValCPublicationApprovalIDValid(model.PublicationApprovalID) ||
		!point14Val0ExactValueValid(model.ApproverRole, point14Val0RoleTypes()) ||
		!point14Val0ApproverRefValid(model.ApproverRef) ||
		strings.TrimSpace(model.ApprovalReason) == "" ||
		!point14ValCExactValueValid(model.ApprovalScope, point14ValCApprovalScopes()) ||
		!point14Val0AuditEventRefValid(model.AuditRef) ||
		!point14Val0GovernanceEventRefValid(model.GovernanceEventRef) ||
		!point14Val0ParsedTimeOk(model.RequestedAt) ||
		!point14Val0CanonicalTimeSourceValid(model.RequestedTimeSource) ||
		!point14Val0ParsedTimeOk(model.ApprovedAt) ||
		!point14Val0CanonicalTimeSourceValid(model.ApprovedTimeSource) ||
		!point14Val0ParsedTimeOk(model.RecordedAt) ||
		!point14Val0CanonicalTimeSourceValid(model.RecordedTimeSource) ||
		!point14ValCExactValueValid(model.PublicationState, point14ValCPublicationStates()) {
		return Point14ValCStateBlocked
	}
	refCount := 0
	if point14ValCCorrectionNoticeIDValid(model.CorrectionNoticeRef) {
		refCount++
	}
	if point14ValCRevocationRequestIDValid(model.RevocationRequestRef) {
		refCount++
	}
	if point14ValCSupersessionRecordIDValid(model.SupersessionRecordRef) {
		refCount++
	}
	if refCount != 1 {
		return Point14ValCStateBlocked
	}
	requestedAt, _ := point14Val0ParsedTime(model.RequestedAt)
	approvedAt, _ := point14Val0ParsedTime(model.ApprovedAt)
	recordedAt, _ := point14Val0ParsedTime(model.RecordedAt)
	if approvedAt.Before(requestedAt) {
		return Point14ValCStateReviewRequired
	}
	if approvedAt.After(recordedAt) {
		return Point14ValCStateReviewRequired
	}
	if model.ApprovesProduction ||
		model.Certifies ||
		model.CreatesPass ||
		model.CreatesPublicBadge ||
		model.AutomaticApprovalRequested ||
		model.GlobalTruthRequested {
		return Point14ValCStateBlocked
	}
	switch strings.TrimSpace(model.PublicationState) {
	case point14ValCPublicationBlocked:
		return Point14ValCStateBlocked
	case point14ValCPublicationReviewRequired:
		return Point14ValCStateReviewRequired
	default:
		return Point14ValCStateActive
	}
}

func point14ValCPublicationVisibilityBoundaryModel(dependency Point14ValCDependencySnapshot) ExternalPublicationVisibilityBoundary {
	return ExternalPublicationVisibilityBoundary{
		VisibilityBoundaryID:     "boundary_point14_valc_visibility_001",
		VisibilityClassification: point14ValCVisibilityPrivateTenantOnly,
		TenantScope:              dependency.InheritedTenantScope,
		PublicPrivateBoundary:    point14ValCPublicationBoundaryPrivate,
		LimitationRefs:           []string{"limitation_ref_point14_valc_001"},
		PublicationTextRefs:      []string{"publication_text_point14_valc_001"},
	}
}

func EvaluatePoint14ValCPublicationVisibilityBoundaryState(model ExternalPublicationVisibilityBoundary, dependency Point14ValCDependencySnapshot) string {
	if !point14Val0BoundaryRefValid(model.VisibilityBoundaryID) ||
		!point14ValCExactValueValid(model.VisibilityClassification, point14ValCVisibilityClassifications()) ||
		!point14ValCExactValueValid(model.PublicPrivateBoundary, point14ValCPublicPrivateBoundaries()) ||
		!point14ValCPublicationTextRefsValid(model.PublicationTextRefs) ||
		!point14ValCPublicationBoundaryPairValid(model.VisibilityClassification, model.PublicPrivateBoundary) {
		return Point14ValCStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) == "" && strings.TrimSpace(model.GlobalScopeClassification) == "" {
		return Point14ValCStateBlocked
	}
	if model.TenantScope != "" {
		if !point11Val0ScopeValid(model.TenantScope) || model.TenantScope != dependency.InheritedTenantScope {
			return Point14ValCStateBlocked
		}
		if strings.TrimSpace(model.GlobalScopeClassification) != "" {
			return Point14ValCStateBlocked
		}
	} else if !point11Val0ContainsTrimmed([]string{point14Val0ScopeGlobalAdvisory, point14Val0ScopePublicNonAuthorative}, model.GlobalScopeClassification) {
		return Point14ValCStateBlocked
	}
	if model.TenantPrivateDataExposed || model.StrengthensClaim || model.ImpliesPublicAuthority || model.MeaningChangingPrivateLimitationOmitted {
		return Point14ValCStateBlocked
	}
	if strings.TrimSpace(model.VisibilityClassification) == point14ValCVisibilityBlocked {
		return Point14ValCStateBlocked
	}
	if strings.TrimSpace(model.VisibilityClassification) == point14ValCVisibilityPublicBounded {
		if !point14ValCLimitationRefsValid(model.LimitationRefs) || !point14ValCRedactionRefsValid(model.RedactionRefs) {
			return Point14ValCStateBlocked
		}
	}
	if strings.TrimSpace(model.VisibilityClassification) == point14ValCVisibilityCustomerBounded ||
		strings.TrimSpace(model.VisibilityClassification) == point14ValCVisibilityAuditorBounded {
		if !point14ValCLimitationRefsValid(model.LimitationRefs) {
			return Point14ValCStateBlocked
		}
	}
	return Point14ValCStateActive
}

func point14ValCTenantPrivacyPublicationGuardModel(dependency Point14ValCDependencySnapshot) TenantPrivacyPublicationGuard {
	return TenantPrivacyPublicationGuard{
		BoundaryID:                  "boundary_point14_valc_privacy_001",
		TenantScope:                 dependency.InheritedTenantScope,
		PublicationTargetScope:      point14ValCVisibilityPrivateTenantOnly,
		PublicPrivateClassification: point14ValCPublicationBoundaryPrivate,
		RedactionRefs:               []string{"redaction_ref_point14_valc_001"},
		LimitationRefs:              []string{"limitation_ref_point14_valc_001"},
		LimitationsVisible:          true,
	}
}

func EvaluatePoint14ValCTenantPrivacyPublicationGuardState(model TenantPrivacyPublicationGuard, dependency Point14ValCDependencySnapshot) string {
	if !point14Val0BoundaryRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		model.TenantScope != dependency.InheritedTenantScope ||
		!point14ValCExactValueValid(model.PublicationTargetScope, point14ValCVisibilityClassifications()) ||
		!point14ValCExactValueValid(model.PublicPrivateClassification, point14ValCPublicPrivateBoundaries()) ||
		!point14ValCPublicationBoundaryPairValid(model.PublicationTargetScope, model.PublicPrivateClassification) ||
		!point14ValCRedactionRefsValid(model.RedactionRefs) ||
		!point14ValCLimitationRefsValid(model.LimitationRefs) {
		return Point14ValCStateBlocked
	}
	if model.TenantPrivateDataExposed || model.CrossTenantPublication || !model.LimitationsVisible {
		return Point14ValCStateBlocked
	}
	if strings.TrimSpace(model.GlobalScopeClassification) != "" &&
		!point11Val0ContainsTrimmed([]string{point14Val0ScopeGlobalAdvisory, point14Val0ScopePublicNonAuthorative}, model.GlobalScopeClassification) {
		return Point14ValCStateBlocked
	}
	if strings.TrimSpace(model.PublicationTargetScope) == point14ValCVisibilityPublicBounded && len(model.RedactionRefs) == 0 {
		return Point14ValCStateBlocked
	}
	return Point14ValCStateActive
}

func point14ValCCorrectionRedactionLimitationGuardModel() CorrectionRedactionLimitationGuard {
	return CorrectionRedactionLimitationGuard{
		GuardID:                  "guard_point14_valc_redaction_001",
		Redacted:                 true,
		RedactionManifestRef:     "redaction_manifest_point14_valc_001",
		RedactionRefs:            []string{"redaction_ref_point14_valc_001"},
		LimitationRefs:           []string{"limitation_ref_point14_valc_001"},
		SourcePublicationState:   Point14ValCStateActive,
		ResolvedPublicationState: Point14ValCStateActive,
	}
}

func EvaluatePoint14ValCCorrectionRedactionLimitationGuardState(model CorrectionRedactionLimitationGuard) string {
	if !point14Val0RefValid(model.GuardID, "guard_") ||
		!point14ValCLimitationRefsValid(model.LimitationRefs) ||
		!point14ValCStateValid(model.SourcePublicationState) ||
		!point14ValCStateValid(model.ResolvedPublicationState) {
		return Point14ValCStateBlocked
	}
	if model.Redacted {
		if !point14ValCRedactionManifestRefValid(model.RedactionManifestRef) || !point14ValCRedactionRefsValid(model.RedactionRefs) {
			return Point14ValCStateBlocked
		}
	}
	if model.DecisiveMissingEvidenceHidden || model.RedactionStrengthensClaim || model.LimitationOmitted || model.SurvivingTextMisleading {
		return Point14ValCStateBlocked
	}
	if (strings.TrimSpace(model.SourcePublicationState) == Point14ValCStateReviewRequired ||
		strings.TrimSpace(model.SourcePublicationState) == Point14ValCStateIncomplete) &&
		strings.TrimSpace(model.ResolvedPublicationState) == Point14ValCStateActive {
		return Point14ValCStateBlocked
	}
	return Point14ValCStateActive
}

func point14ValCGovernanceTraceModel() CorrectionPublicationGovernanceTrace {
	return CorrectionPublicationGovernanceTrace{
		GovernanceTraceID:       "governance_trace_point14_valc_001",
		GovernanceEventRef:      "governance_event_point14_valc_001",
		Owner:                   "approver_point14_valc_owner_001",
		ApproverRole:            "security_reviewer",
		AuditRef:                "audit_event_point14_valc_governance_001",
		EvidenceRefs:            []string{"evidence_point14_valc_001"},
		DecisionReason:          "correction pending governance approval",
		Timestamp:               "2026-05-06T08:04:00Z",
		TimeSource:              point14Val0TimeSourceServerUTC,
		DisputeOpenedAt:         "2026-05-06T08:00:00Z",
		DisputeOpenedTimeSource: point14Val0TimeSourceServerUTC,
	}
}

func EvaluatePoint14ValCGovernanceTraceState(model CorrectionPublicationGovernanceTrace) string {
	if !point14ValCGovernanceTraceIDValid(model.GovernanceTraceID) ||
		!point14Val0GovernanceEventRefValid(model.GovernanceEventRef) ||
		!point14Val0ApproverRefValid(model.Owner) ||
		!point14Val0ExactValueValid(model.ApproverRole, point14Val0RoleTypes()) ||
		!point14Val0AuditEventRefValid(model.AuditRef) ||
		!point14Val0EvidenceRefsValid(model.EvidenceRefs) ||
		strings.TrimSpace(model.DecisionReason) == "" ||
		!point14Val0ParsedTimeOk(model.Timestamp) ||
		!point14Val0CanonicalTimeSourceValid(model.TimeSource) {
		return Point14ValCStateBlocked
	}
	if strings.TrimSpace(model.DisputeOpenedAt) != "" {
		if !point14Val0ParsedTimeOk(model.DisputeOpenedAt) || !point14Val0CanonicalTimeSourceValid(model.DisputeOpenedTimeSource) {
			return Point14ValCStateBlocked
		}
		timestamp, _ := point14Val0ParsedTime(model.Timestamp)
		disputeOpenedAt, _ := point14Val0ParsedTime(model.DisputeOpenedAt)
		if timestamp.Before(disputeOpenedAt) {
			return Point14ValCStateReviewRequired
		}
	}
	if model.ApprovesProduction || model.CertifiesCompliance {
		return Point14ValCStateBlocked
	}
	return Point14ValCStateActive
}

func point14ValCAgentCorrectionPublicationBoundaryModel(dependency Point14ValCDependencySnapshot) AgentCorrectionPublicationBoundary {
	return AgentCorrectionPublicationBoundary{
		BoundaryID:     "boundary_point14_valc_agent_001",
		TenantScope:    dependency.InheritedTenantScope,
		AgentInputRefs: []string{"agent_input_point14_valc_001"},
		EvidenceRefs:   []string{"evidence_point14_valc_001"},
		AuditEventRef:  "audit_event_point14_valc_agent_001",
		AdvisoryOnly:   true,
	}
}

func EvaluatePoint14ValCAgentCorrectionPublicationBoundaryState(model AgentCorrectionPublicationBoundary, dependency Point14ValCDependencySnapshot) string {
	if !point14Val0BoundaryRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		model.TenantScope != dependency.InheritedTenantScope ||
		!point14Val0AgentInputRefsValid(model.AgentInputRefs) ||
		!point14Val0EvidenceRefsValid(model.EvidenceRefs) ||
		!point14Val0AuditEventRefValid(model.AuditEventRef) {
		return Point14ValCStateBlocked
	}
	if !model.AdvisoryOnly ||
		model.CanApproveCorrection ||
		model.CanApproveRevocation ||
		model.CanApprovePublication ||
		model.CanPublishNotice ||
		model.CanSatisfyGovernanceTraceAlone ||
		model.CanEmitPass ||
		model.CanEmitPublicAuthority ||
		model.ApprovalGranted ||
		model.ProductionApproved ||
		model.ExternalAuthorityAllowed {
		return Point14ValCStateBlocked
	}
	return Point14ValCStateActive
}

func point14ValCNoExternalAuthorityPublicationGuardModel() Point14ValCNoExternalAuthorityPublicationGuard {
	return Point14ValCNoExternalAuthorityPublicationGuard{}
}

func EvaluatePoint14ValCNoExternalAuthorityPublicationGuardState(model Point14ValCNoExternalAuthorityPublicationGuard) string {
	for _, marker := range model.ObservedAuthorityMarkers {
		if point11Val0ContainsTrimmed(point14ValCForbiddenAuthorityMarkers(), marker) {
			return Point14ValCStateBlocked
		}
	}
	if model.CanonicalAuthorityGranted ||
		model.ProductionApprovalGranted ||
		model.PublishesCorrection ||
		model.RevokesClaim ||
		model.OverridesCanonicalDecision ||
		model.CorrectionPublished ||
		model.PublicationApproved ||
		model.PublicBadgeAllowed ||
		model.CorrectionAutoPublished ||
		model.RevocationAutoExecuted ||
		model.PublicationAutoApproved ||
		model.CrowdApprovedPublication ||
		model.AgentApprovedPublication ||
		model.PassAllowed ||
		model.ExternalAuthorityAllowed {
		return Point14ValCStateBlocked
	}
	return Point14ValCStateActive
}

func point14ValCNoOverclaimPublicationWordingModel() Point14ValCNoOverclaimPublicationWording {
	return Point14ValCNoOverclaimPublicationWording{
		ObservedCorrectionTexts:   []string{"bounded correction notice", "correction pending governance approval"},
		ObservedRevocationTexts:   []string{"revocation request pending governance review"},
		ObservedSupersessionTexts: []string{"supersession record preserves prior context"},
		ObservedPublicationTexts:  []string{"bounded public notice", "publication approval is not production approval"},
		ObservedVisibilityTexts:   []string{"private tenant-scoped correction"},
		ObservedPrivacyTexts:      []string{"limitations remain visible"},
		ObservedGovernanceTexts:   []string{"canonical evidence spine remains source of truth"},
		ObservedAgentTexts:        []string{"no public badge authority granted"},
		AllowedSafeWording:        point14ValCSafeWording(),
		BlockedWording:            point14ValCForbiddenWording(),
		ProjectionDisclaimer:      point14ValCPublicationDisclaimerBase,
	}
}

func EvaluatePoint14ValCNoOverclaimPublicationWordingState(model Point14ValCNoOverclaimPublicationWording) string {
	if model.ProjectionDisclaimer != point14ValCPublicationDisclaimerBase ||
		!point14Val0TextListValid(model.AllowedSafeWording) ||
		!point14Val0TextListValid(model.BlockedWording) {
		return Point14ValCStateBlocked
	}
	if point14ValCObservedListContainsForbiddenWording(model.ObservedCorrectionTexts) ||
		point14ValCObservedListContainsForbiddenWording(model.ObservedRevocationTexts) ||
		point14ValCObservedListContainsForbiddenWording(model.ObservedSupersessionTexts) ||
		point14ValCObservedListContainsForbiddenWording(model.ObservedPublicationTexts) ||
		point14ValCObservedListContainsForbiddenWording(model.ObservedVisibilityTexts) ||
		point14ValCObservedListContainsForbiddenWording(model.ObservedPrivacyTexts) ||
		point14ValCObservedListContainsForbiddenWording(model.ObservedGovernanceTexts) ||
		point14ValCObservedListContainsForbiddenWording(model.ObservedAgentTexts) {
		return Point14ValCStateBlocked
	}
	if point14ValCObservedListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point14ValCStateBlocked
	}
	return Point14ValCStateActive
}

func point14ValCFoundationState(states ...string) string {
	hasReview := false
	hasIncomplete := false
	for _, state := range states {
		switch strings.TrimSpace(state) {
		case Point14ValCStateBlocked:
			return Point14ValCStateBlocked
		case Point14ValCStateReviewRequired:
			hasReview = true
		case Point14ValCStateIncomplete:
			hasIncomplete = true
		}
	}
	if hasReview {
		return Point14ValCStateReviewRequired
	}
	if hasIncomplete {
		return Point14ValCStateIncomplete
	}
	return Point14ValCStateActive
}

func point14ValCBlockingReasons(model Point14ValCFoundation) []string {
	componentStates := map[string]string{
		"dependency":                 model.DependencyState,
		"correction_notice":          model.CorrectionNoticeState,
		"revocation_request":         model.RevocationRequestState,
		"supersession_record":        model.SupersessionRecordState,
		"publication_approval":       model.PublicationApprovalState,
		"visibility_boundary":        model.VisibilityBoundaryState,
		"tenant_privacy":             model.TenantPrivacyState,
		"redaction_limitation":       model.RedactionLimitationState,
		"governance_trace":           model.GovernanceTraceState,
		"agent_publication_boundary": model.AgentPublicationBoundaryState,
		"no_external_authority":      model.NoExternalAuthorityState,
		"no_overclaim":               model.NoOverclaimState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if state == Point14ValCStateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point14ValCReviewPrerequisites(model Point14ValCFoundation) []string {
	componentStates := map[string]string{
		"dependency":                 model.DependencyState,
		"correction_notice":          model.CorrectionNoticeState,
		"revocation_request":         model.RevocationRequestState,
		"supersession_record":        model.SupersessionRecordState,
		"publication_approval":       model.PublicationApprovalState,
		"visibility_boundary":        model.VisibilityBoundaryState,
		"tenant_privacy":             model.TenantPrivacyState,
		"redaction_limitation":       model.RedactionLimitationState,
		"governance_trace":           model.GovernanceTraceState,
		"agent_publication_boundary": model.AgentPublicationBoundaryState,
		"no_external_authority":      model.NoExternalAuthorityState,
		"no_overclaim":               model.NoOverclaimState,
	}
	prereqs := []string{}
	for name, state := range componentStates {
		if state == Point14ValCStateReviewRequired || state == Point14ValCStateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func point14ValCFoundationModelFromUpstream(valB Point14ValBFoundation) Point14ValCFoundation {
	dependency := point14ValCDependencySnapshotFromUpstream(valB)
	return Point14ValCFoundation{
		CurrentState:                  Point14ValCStateActive,
		ProjectionDisclaimer:          point14ValCPublicationDisclaimerBase,
		DependencyState:               Point14ValCStateActive,
		CorrectionNoticeState:         Point14ValCStateActive,
		RevocationRequestState:        Point14ValCStateActive,
		SupersessionRecordState:       Point14ValCStateActive,
		PublicationApprovalState:      Point14ValCStateActive,
		VisibilityBoundaryState:       Point14ValCStateActive,
		TenantPrivacyState:            Point14ValCStateActive,
		RedactionLimitationState:      Point14ValCStateActive,
		GovernanceTraceState:          Point14ValCStateActive,
		AgentPublicationBoundaryState: Point14ValCStateActive,
		NoExternalAuthorityState:      Point14ValCStateActive,
		NoOverclaimState:              Point14ValCStateActive,
		Dependency:                    dependency,
		CorrectionNoticeBoundary:      point14ValCCorrectionNoticeBoundaryModel(),
		RevocationRequestBoundary:     point14ValCRevocationRequestBoundaryModel(),
		SupersessionRecordBoundary:    point14ValCSupersessionRecordBoundaryModel(),
		PublicationApprovalBoundary:   point14ValCPublicationApprovalBoundaryModel(),
		PublicationVisibilityBoundary: point14ValCPublicationVisibilityBoundaryModel(dependency),
		TenantPrivacyPublicationGuard: point14ValCTenantPrivacyPublicationGuardModel(dependency),
		CorrectionRedactionGuard:      point14ValCCorrectionRedactionLimitationGuardModel(),
		GovernanceTrace:               point14ValCGovernanceTraceModel(),
		AgentCorrectionBoundary:       point14ValCAgentCorrectionPublicationBoundaryModel(dependency),
		NoExternalAuthorityGuard:      point14ValCNoExternalAuthorityPublicationGuardModel(),
		NoOverclaimPublicationWording: point14ValCNoOverclaimPublicationWordingModel(),
	}
}

func Point14ValCFoundationModel() Point14ValCFoundation {
	return cachedFormalModel(&point14ValCFoundationModelOnce, &point14ValCFoundationModelCached, func() Point14ValCFoundation {
		valB := ComputePoint14ValBFoundation(Point14ValBFoundationModel())
		return point14ValCFoundationModelFromUpstream(valB)
	})
}

func ComputePoint14ValCFoundation(model Point14ValCFoundation) Point14ValCFoundation {
	model.DependencyState = EvaluatePoint14ValCDependencyState(model.Dependency)
	model.CorrectionNoticeState = EvaluatePoint14ValCCorrectionNoticeBoundaryState(model.CorrectionNoticeBoundary)
	model.RevocationRequestState = EvaluatePoint14ValCRevocationRequestBoundaryState(model.RevocationRequestBoundary)
	model.SupersessionRecordState = EvaluatePoint14ValCSupersessionRecordBoundaryState(model.SupersessionRecordBoundary)
	model.PublicationApprovalState = EvaluatePoint14ValCPublicationApprovalBoundaryState(model.PublicationApprovalBoundary)
	model.VisibilityBoundaryState = EvaluatePoint14ValCPublicationVisibilityBoundaryState(model.PublicationVisibilityBoundary, model.Dependency)
	model.TenantPrivacyState = EvaluatePoint14ValCTenantPrivacyPublicationGuardState(model.TenantPrivacyPublicationGuard, model.Dependency)
	model.RedactionLimitationState = EvaluatePoint14ValCCorrectionRedactionLimitationGuardState(model.CorrectionRedactionGuard)
	model.GovernanceTraceState = EvaluatePoint14ValCGovernanceTraceState(model.GovernanceTrace)
	model.AgentPublicationBoundaryState = EvaluatePoint14ValCAgentCorrectionPublicationBoundaryState(model.AgentCorrectionBoundary, model.Dependency)
	model.NoExternalAuthorityState = EvaluatePoint14ValCNoExternalAuthorityPublicationGuardState(model.NoExternalAuthorityGuard)
	model.NoOverclaimState = EvaluatePoint14ValCNoOverclaimPublicationWordingState(model.NoOverclaimPublicationWording)
	model.CurrentState = point14ValCFoundationState(
		model.DependencyState,
		model.CorrectionNoticeState,
		model.RevocationRequestState,
		model.SupersessionRecordState,
		model.PublicationApprovalState,
		model.VisibilityBoundaryState,
		model.TenantPrivacyState,
		model.RedactionLimitationState,
		model.GovernanceTraceState,
		model.AgentPublicationBoundaryState,
		model.NoExternalAuthorityState,
		model.NoOverclaimState,
	)
	model.BlockingReasons = point14ValCBlockingReasons(model)
	model.ReviewPrerequisites = point14ValCReviewPrerequisites(model)
	return model
}

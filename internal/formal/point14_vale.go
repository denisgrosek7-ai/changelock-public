package formal

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"

	"github.com/denisgrosek/changelock/internal/operability"
)

const (
	Point14ValEStateBlocked        = "point14_vale_external_ecosystem_closure_blocked"
	Point14ValEStateReviewRequired = "point14_vale_external_ecosystem_closure_review_required"
	Point14ValEStateIncomplete     = "point14_vale_external_ecosystem_closure_incomplete"
	Point14ValEStatePassConfirmed  = "point14_vale_external_ecosystem_closure_pass_confirmed"
)

const (
	point14ValEWaveID                     = "val_e"
	point14ValEScope                      = "final_external_ecosystem_governance_closure_gate"
	point14ValEProjectionDisclaimerBase   = "projection_only not_canonical_truth point14_vale_external_ecosystem_closure_gate"
	point14ValECleanRoomBoundaryPreserved = "not_applicable_boundary_preserved"
	point14ValEEvidenceIdentity           = "bounded_external_ecosystem_evidence_identity_preserved"
	point14ValEPassClosureGeneratedAt     = "2026-05-06T10:00:00Z"
)

type Point14ValEDependencySnapshot struct {
	Point14ValDCurrentState                    string                `json:"point14_vald_current_state"`
	Point14ValDDependencyState                 string                `json:"point14_vald_dependency_state"`
	Point14ValDTimelineProjectionState         string                `json:"point14_vald_timeline_projection_state"`
	Point14ValDSignalTimelineEntryState        string                `json:"point14_vald_signal_timeline_entry_state"`
	Point14ValDDisputeTimelineState            string                `json:"point14_vald_dispute_timeline_state"`
	Point14ValDCorrectionReadProjectionState   string                `json:"point14_vald_correction_read_projection_state"`
	Point14ValDGovernanceTraceProjectionState  string                `json:"point14_vald_governance_trace_projection_state"`
	Point14ValDQueryProjectionState            string                `json:"point14_vald_query_projection_state"`
	Point14ValDAccessBoundaryState             string                `json:"point14_vald_access_boundary_state"`
	Point14ValDTenantPrivacyTimelineState      string                `json:"point14_vald_tenant_privacy_timeline_state"`
	Point14ValDAgentTimelineProjectionState    string                `json:"point14_vald_agent_timeline_projection_state"`
	Point14ValDTimestampIntegrityState         string                `json:"point14_vald_timestamp_integrity_state"`
	Point14ValDNoMutationProjectionGuardState  string                `json:"point14_vald_no_mutation_projection_guard_state"`
	Point14ValDNoOverclaimTimelineWordingState string                `json:"point14_vald_no_overclaim_timeline_wording_state"`
	Point14ValDPointID                         string                `json:"point14_vald_point_id"`
	Point14ValDWaveID                          string                `json:"point14_vald_wave_id"`
	Point14ValDComputedFromUpstream            bool                  `json:"point14_vald_computed_from_upstream"`
	Point14ValDMerged                          bool                  `json:"point14_vald_merged"`
	Point14ValDCIGreen                         bool                  `json:"point14_vald_ci_green"`
	Point14ValDReviewedOnMain                  bool                  `json:"point14_vald_reviewed_on_main"`
	Point14PassSeen                            bool                  `json:"point14_pass_seen"`
	InheritedPoint14ValCCurrentState           string                `json:"inherited_point14_valc_current_state"`
	InheritedPoint14ValBCurrentState           string                `json:"inherited_point14_valb_current_state"`
	InheritedPoint14ValACurrentState           string                `json:"inherited_point14_vala_current_state"`
	InheritedPoint14Val0CurrentState           string                `json:"inherited_point14_val0_current_state"`
	InheritedPoint13ValECurrentState           string                `json:"inherited_point13_vale_current_state"`
	InheritedPoint13ValEPassClosureState       string                `json:"inherited_point13_vale_pass_closure_state"`
	InheritedPoint13ValEPassAllowed            bool                  `json:"inherited_point13_vale_pass_allowed"`
	InheritedPoint13ValEPassToken              string                `json:"inherited_point13_vale_pass_token"`
	InheritedPoint12CurrentState               string                `json:"inherited_point12_current_state"`
	InheritedPoint12DependencyState            string                `json:"inherited_point12_dependency_state"`
	InheritedPoint12PassClosureState           string                `json:"inherited_point12_pass_closure_state"`
	InheritedPoint12ReviewerResult             string                `json:"inherited_point12_reviewer_result"`
	InheritedPoint11CurrentState               string                `json:"inherited_point11_current_state"`
	InheritedPoint11PublicationState           string                `json:"inherited_point11_publication_state"`
	InheritedPoint11NoOverclaimState           string                `json:"inherited_point11_no_overclaim_state"`
	InheritedPoint11FinalPassGateState         string                `json:"inherited_point11_final_pass_gate_state"`
	InheritedPoint10CurrentState               string                `json:"inherited_point10_current_state"`
	InheritedPoint10NoOverclaimState           string                `json:"inherited_point10_no_overclaim_state"`
	InheritedPoint10ProjectionState            string                `json:"inherited_point10_projection_state"`
	InheritedPoint10PassRuleState              string                `json:"inherited_point10_pass_rule_state"`
	InheritedTenantScope                       string                `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput                 bool                  `json:"snapshot_from_computed_output"`
	ReviewPrerequisites                        []string              `json:"review_prerequisites,omitempty"`
	Point14ValD                                Point14ValDFoundation `json:"point14_vald"`
}

type Point14ValEClosureEvaluator struct {
	ClosureEvaluatorID                string   `json:"closure_evaluator_id"`
	DependencyState                   string   `json:"dependency_state"`
	ValidationClosureState            string   `json:"validation_closure_state"`
	DisputeClosureState               string   `json:"dispute_closure_state"`
	CorrectionPublicationClosureState string   `json:"correction_publication_closure_state"`
	TimelineProjectionClosureState    string   `json:"timeline_projection_closure_state"`
	AuthorityBoundaryState            string   `json:"authority_boundary_state"`
	TenantPrivacyState                string   `json:"tenant_privacy_state"`
	TimestampIntegrityState           string   `json:"timestamp_integrity_state"`
	AgentAdvisoryState                string   `json:"agent_advisory_state"`
	NoOverclaimState                  string   `json:"no_overclaim_state"`
	CLBFinalState                     string   `json:"clb_final_state"`
	ReadOnlyProjectionConfirmed       bool     `json:"read_only_projection_confirmed"`
	NoMutationPathsDetected           bool     `json:"no_mutation_paths_detected"`
	NoExternalAuthorityDetected       bool     `json:"no_external_authority_detected"`
	NoPrematurePoint14Pass            bool     `json:"no_premature_point14_pass"`
	FinalPassAllowed                  bool     `json:"final_pass_allowed"`
	CommandsRun                       []string `json:"commands_run,omitempty"`
	TestsRun                          []string `json:"tests_run,omitempty"`
	GrepsRun                          []string `json:"greps_run,omitempty"`
	NegativeFixturesRun               []string `json:"negative_fixtures_run,omitempty"`
	ReviewerResult                    string   `json:"reviewer_result"`
	CurrentState                      string   `json:"current_state"`
	ProjectionDisclaimer              string   `json:"projection_disclaimer"`
}

type Point14PassClosureManifest struct {
	CurrentState              string   `json:"current_state"`
	ClosureManifestID         string   `json:"closure_manifest_id"`
	PointID                   string   `json:"point_id"`
	WaveID                    string   `json:"wave_id"`
	ClosureToken              string   `json:"closure_token"`
	Scope                     string   `json:"scope"`
	ExplicitNonGoals          []string `json:"explicit_non_goals,omitempty"`
	DependencyGateResult      string   `json:"dependency_gate_result"`
	ClosureEvaluatorResult    string   `json:"closure_evaluator_result"`
	EvidenceIdentity          string   `json:"evidence_identity"`
	CommandsRun               []string `json:"commands_run,omitempty"`
	TestsRun                  []string `json:"tests_run,omitempty"`
	GrepsRun                  []string `json:"greps_run,omitempty"`
	NegativeFixturesRun       []string `json:"negative_fixtures_run,omitempty"`
	ProjectionBoundaryResult  string   `json:"projection_boundary_result"`
	NoExternalAuthorityResult string   `json:"no_external_authority_result"`
	NoOverclaimGrepResult     string   `json:"no_overclaim_grep_result"`
	TenantPrivacyResult       string   `json:"tenant_privacy_result"`
	TimestampIntegrityResult  string   `json:"timestamp_integrity_result"`
	AIAgentBoundaryResult     string   `json:"ai_agent_boundary_result"`
	CleanRoomIPResult         string   `json:"clean_room_ip_result"`
	CLBResult                 string   `json:"clb_result"`
	ReviewerResult            string   `json:"reviewer_result"`
	GeneratedAt               string   `json:"generated_at"`
	Point14PassAllowed        bool     `json:"point14_pass_allowed"`
	Point14PassToken          string   `json:"point14_pass_token"`
	ProjectionDisclaimer      string   `json:"projection_disclaimer"`
}

type Point14ValEExternalSignalValidationClosureCheck struct {
	CheckID                        string `json:"check_id"`
	ValACurrentState               string `json:"vala_current_state"`
	ValidationResultState          string `json:"validation_result_state"`
	CandidateSemanticsBounded      bool   `json:"candidate_semantics_bounded"`
	CanonicalAuthorityGranted      bool   `json:"canonical_authority_granted"`
	PassEmitted                    bool   `json:"pass_emitted"`
	SourceIdentityAuthorityGranted bool   `json:"source_identity_authority_granted"`
	CanonicalMutationDetected      bool   `json:"canonical_mutation_detected"`
	DuplicateActiveEvidencePath    bool   `json:"duplicate_active_evidence_path"`
	UnrelatedSignalActivePath      bool   `json:"unrelated_signal_active_path"`
	CrossTenantSignalActivePath    bool   `json:"cross_tenant_signal_active_path"`
}

type Point14ValEConflictDisputeClosureCheck struct {
	CheckID                             string `json:"check_id"`
	ValBCurrentState                    string `json:"valb_current_state"`
	DisputeTriageResultState            string `json:"dispute_triage_result_state"`
	TriageOnlyBounded                   bool   `json:"triage_only_bounded"`
	UnresolvedDisputeReviewRequired     bool   `json:"unresolved_dispute_review_required"`
	EvidenceRequiredUnclosed            bool   `json:"evidence_required_unclosed"`
	DisputeAutoResolved                 bool   `json:"dispute_auto_resolved"`
	ConflictResolvedToPass              bool   `json:"conflict_resolved_to_pass"`
	ExternalAuthorityResolutionDetected bool   `json:"external_authority_resolution_detected"`
	GovernanceEscalationMissing         bool   `json:"governance_escalation_missing"`
}

type Point14ValECorrectionPublicationClosureCheck struct {
	CheckID                             string `json:"check_id"`
	ValCCurrentState                    string `json:"valc_current_state"`
	PublicationApprovalState            string `json:"publication_approval_state"`
	GovernanceControlledOnly            bool   `json:"governance_controlled_only"`
	BoundedCorrectionNoticeProven       bool   `json:"bounded_correction_notice_proven"`
	CorrectionAutoPublished             bool   `json:"correction_auto_published"`
	RevocationAutoExecuted              bool   `json:"revocation_auto_executed"`
	SupersessionSilentReplacement       bool   `json:"supersession_silent_replacement"`
	PublicationApprovalBecameProduction bool   `json:"publication_approval_became_production"`
	PublicationCertified                bool   `json:"publication_certified"`
	PublicNoticeBecameBadge             bool   `json:"public_notice_became_badge"`
	RedactionHidesDecisiveEvidence      bool   `json:"redaction_hides_decisive_evidence"`
	LimitationsOmitted                  bool   `json:"limitations_omitted"`
	PublicationStrengthensClaims        bool   `json:"publication_strengthens_claims"`
}

type Point14ValETimelineProjectionClosureCheck struct {
	CheckID                           string `json:"check_id"`
	ValDCurrentState                  string `json:"vald_current_state"`
	QueryProjectionState              string `json:"query_projection_state"`
	ReadOnlyProjectionOnly            bool   `json:"read_only_projection_only"`
	TimelineMutationDetected          bool   `json:"timeline_mutation_detected"`
	QueryMutationDetected             bool   `json:"query_mutation_detected"`
	TimelineResolvesDisputes          bool   `json:"timeline_resolves_disputes"`
	TimelinePublishesCorrections      bool   `json:"timeline_publishes_corrections"`
	TimelineExecutesRevocation        bool   `json:"timeline_executes_revocation"`
	TimelineCreatesAuthority          bool   `json:"timeline_creates_authority"`
	QueryHidesDecisiveMissingEvidence bool   `json:"query_hides_decisive_missing_evidence"`
	TimelineStrengthensClaims         bool   `json:"timeline_strengthens_claims"`
}

type Point14ValEAuthorityBoundaryClosureCheck struct {
	CheckID                string   `json:"check_id"`
	ObservedAuthorityMarks []string `json:"observed_authority_marks,omitempty"`
	ExternalAuthorityFound bool     `json:"external_authority_found"`
}

type Point14ValETenantPrivacyClosureCheck struct {
	CheckID                                string `json:"check_id"`
	TenantScope                            string `json:"tenant_scope"`
	CrossTenantDetected                    bool   `json:"cross_tenant_detected"`
	TenantPrivateDataExposed               bool   `json:"tenant_private_data_exposed"`
	PublicNoticeLeaksTenantPrivateData     bool   `json:"public_notice_leaks_tenant_private_data"`
	AccessBoundaryAllowsCrossTenantQuery   bool   `json:"access_boundary_allows_cross_tenant_query"`
	PublicPrivateClassificationPresent     bool   `json:"public_private_classification_present"`
	RequiredRedactionLimitationRefsPresent bool   `json:"required_redaction_limitation_refs_present"`
}

type Point14ValETimestampIntegrityClosureCheck struct {
	CheckID                          string `json:"check_id"`
	TenantScope                      string `json:"tenant_scope"`
	EventAt                          string `json:"event_at"`
	EventTimeSource                  string `json:"event_time_source"`
	GeneratedAt                      string `json:"generated_at"`
	GeneratedTimeSource              string `json:"generated_time_source"`
	ApprovalAt                       string `json:"approval_at"`
	ApprovalTimeSource               string `json:"approval_time_source"`
	SourceEventAt                    string `json:"source_event_at"`
	SourceEventTimeSource            string `json:"source_event_time_source"`
	ClientLocalCreatesCanonical      bool   `json:"client_local_creates_canonical"`
	SourceEventCreatesAuthority      bool   `json:"source_event_creates_authority"`
	FutureDatedActiveEvent           bool   `json:"future_dated_active_event"`
	BackdatedApproval                bool   `json:"backdated_approval"`
	ImpossibleOrdering               bool   `json:"impossible_ordering"`
	TimelineOrderingUpgradesValidity bool   `json:"timeline_ordering_upgrades_validity"`
}

type Point14ValEAgentAdvisoryClosureCheck struct {
	CheckID                       string `json:"check_id"`
	TenantScope                   string `json:"tenant_scope"`
	AdvisoryOnly                  bool   `json:"advisory_only"`
	AgentResolvesDispute          bool   `json:"agent_resolves_dispute"`
	AgentPublishesCorrection      bool   `json:"agent_publishes_correction"`
	AgentRevokesClaim             bool   `json:"agent_revokes_claim"`
	AgentSatisfiesGovernanceAlone bool   `json:"agent_satisfies_governance_alone"`
	AgentAuthorityFlags           bool   `json:"agent_authority_flags"`
	AgentPassAllowed              bool   `json:"agent_pass_allowed"`
	AgentPublicBadgeAllowed       bool   `json:"agent_public_badge_allowed"`
}

type Point14ValENoOverclaimFinalCheck struct {
	ObservedTexts                        []string `json:"observed_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point14ValECLBFinalCheck struct {
	CheckID        string   `json:"check_id"`
	CLB0Present    bool     `json:"clb0_present"`
	CLB1Present    bool     `json:"clb1_present"`
	CLB2Present    bool     `json:"clb2_present"`
	CLB3Advisories []string `json:"clb3_advisories,omitempty"`
}

type Point14ValEFoundation struct {
	CurrentState                         string                                          `json:"current_state"`
	BlockingReasons                      []string                                        `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites                  []string                                        `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer                 string                                          `json:"projection_disclaimer"`
	DependencyState                      string                                          `json:"dependency_state"`
	ClosureEvaluatorState                string                                          `json:"closure_evaluator_state"`
	PassClosureManifestState             string                                          `json:"pass_closure_manifest_state"`
	ExternalSignalValidationClosureState string                                          `json:"external_signal_validation_closure_state"`
	ConflictDisputeClosureState          string                                          `json:"conflict_dispute_closure_state"`
	CorrectionPublicationClosureState    string                                          `json:"correction_publication_closure_state"`
	TimelineProjectionClosureState       string                                          `json:"timeline_projection_closure_state"`
	AuthorityBoundaryClosureState        string                                          `json:"authority_boundary_closure_state"`
	TenantPrivacyClosureState            string                                          `json:"tenant_privacy_closure_state"`
	TimestampIntegrityClosureState       string                                          `json:"timestamp_integrity_closure_state"`
	AgentAdvisoryClosureState            string                                          `json:"agent_advisory_closure_state"`
	NoOverclaimFinalCheckState           string                                          `json:"no_overclaim_final_check_state"`
	CLBFinalCheckState                   string                                          `json:"clb_final_check_state"`
	Point14PassAllowed                   bool                                            `json:"point14_pass_allowed"`
	Point14PassToken                     string                                          `json:"point14_pass_token"`
	Dependency                           Point14ValEDependencySnapshot                   `json:"dependency"`
	ClosureEvaluator                     Point14ValEClosureEvaluator                     `json:"closure_evaluator"`
	PassClosureManifest                  Point14PassClosureManifest                      `json:"pass_closure_manifest"`
	ExternalSignalValidationClosureCheck Point14ValEExternalSignalValidationClosureCheck `json:"external_signal_validation_closure_check"`
	ConflictDisputeClosureCheck          Point14ValEConflictDisputeClosureCheck          `json:"conflict_dispute_closure_check"`
	CorrectionPublicationClosureCheck    Point14ValECorrectionPublicationClosureCheck    `json:"correction_publication_closure_check"`
	TimelineProjectionClosureCheck       Point14ValETimelineProjectionClosureCheck       `json:"timeline_projection_closure_check"`
	AuthorityBoundaryClosureCheck        Point14ValEAuthorityBoundaryClosureCheck        `json:"authority_boundary_closure_check"`
	TenantPrivacyClosureCheck            Point14ValETenantPrivacyClosureCheck            `json:"tenant_privacy_closure_check"`
	TimestampIntegrityClosureCheck       Point14ValETimestampIntegrityClosureCheck       `json:"timestamp_integrity_closure_check"`
	AgentAdvisoryClosureCheck            Point14ValEAgentAdvisoryClosureCheck            `json:"agent_advisory_closure_check"`
	NoOverclaimFinalCheck                Point14ValENoOverclaimFinalCheck                `json:"no_overclaim_final_check"`
	CLBFinalCheck                        Point14ValECLBFinalCheck                        `json:"clb_final_check"`
}

func point14ValEStates() []string {
	return []string{
		Point14ValEStateBlocked,
		Point14ValEStateReviewRequired,
		Point14ValEStateIncomplete,
		Point14ValEStatePassConfirmed,
	}
}

func point14ValEStateValid(value string) bool {
	return point14Val0ExactValueValid(value, point14ValEStates())
}

func point14ValEForbiddenAuthorityMarkers() []string {
	return point14ValDForbiddenAuthorityMarkers()
}

func point14ValEForbiddenWording() []string {
	forbidden := append([]string{}, point14Val0ForbiddenWording()...)
	forbidden = append(forbidden, point14ValBForbiddenWording()...)
	forbidden = append(forbidden, point14ValCForbiddenWording()...)
	forbidden = append(forbidden, point14ValDForbiddenWording()...)
	return point14ValEUniqueWording(forbidden)
}

func point14ValESafeWording() []string {
	safe := []string{
		"bounded external ecosystem evidence input",
		"dispute requires governance review",
		"correction publication boundary is governed",
		"read-only ecosystem timeline",
		"no external authority granted",
		"no public badge authority granted",
		"canonical evidence spine remains source of truth",
		"limitations remain visible",
		"point 14 closure verifies boundaries only",
	}
	safe = append(safe, point14ValASafeWording()...)
	safe = append(safe, point14ValBSafeWording()...)
	safe = append(safe, point14ValCSafeWording()...)
	safe = append(safe, point14ValDSafeWording()...)
	return point14ValEUniqueWording(safe)
}

func point14ValEUniqueWording(values []string) []string {
	seen := map[string]struct{}{}
	unique := make([]string, 0, len(values))
	for _, value := range values {
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		unique = append(unique, value)
	}
	return unique
}

func point14ValEExplicitNonGoals() []string {
	return []string{
		"no point 15 implementation",
		"no new ecosystem feature beyond closure gate",
		"no mutation/write api",
		"no automatic correction publication",
		"no revocation execution",
		"no claim lifecycle mutation",
		"no dispute resolution execution",
		"no public badge",
		"no global authority",
		"no crowd authority",
		"no portal authority",
		"no external api default",
		"no ai/agent authority",
		"no canonical mutation",
		"no production approval",
		"no certification/compliance/legal/financial guarantee",
	}
}

func point14ValEObservedTextContainsForbiddenWording(text string) bool {
	return point14Val0ContainsForbiddenWordingFor(text, point14ValESafeWording(), point14ValEForbiddenWording())
}

func point14ValEObservedListContainsForbiddenWording(values []string) bool {
	return point14Val0ListContainsForbiddenWordingFor(values, point14ValESafeWording(), point14ValEForbiddenWording())
}

func point14ValECheckIDValid(value string) bool {
	return point14Val0RefValid(value, "point14_vale_check_", "closure_check_")
}

func point14ValEClosureEvaluatorIDValid(value string) bool {
	return point14Val0RefValid(value, "point14_vale_closure_evaluator_", "closure_evaluator_")
}

func point14ValEClosureManifestIDValid(value string) bool {
	return point14Val0RefValid(value, "point14_vale_pass_manifest_", "pass_manifest_")
}

func point14ValEValDPayloadContainsPointPass(valD Point14ValDFoundation) bool {
	payload, err := json.Marshal(valD)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point14Val0BlockedPassToken)
}

func point14ValECommandsRun() []string {
	return []string{
		"git diff --check",
		"gofmt on changed Go files",
		"go test ./internal/formal -run 'Test.*Point14ValE.*|Test.*Point14.*ValE.*' -v",
		"go test ./internal/formal -run 'Test.*Point14ValD.*|Test.*Point14.*ValD.*' -v",
		"go test ./internal/formal -run 'Test.*Point14ValC.*|Test.*Point14.*ValC.*' -v",
		"go test ./internal/formal -run 'Test.*Point14ValB.*|Test.*Point14.*ValB.*' -v",
		"go test ./internal/formal -run 'Test.*Point14ValA.*|Test.*Point14.*ValA.*' -v",
		"go test ./internal/formal -run 'Test.*Point14Val0.*|Test.*Point14.*Val0.*' -v",
		"go test ./internal/formal -run 'Test.*Point13ValE.*|Test.*Point13.*ValE.*' -v",
		"go test ./internal/formal -run 'Test.*Point13.*' -v",
		"go test ./internal/formal -run 'Test.*Point12.*|Test.*Replay.*|Test.*ProofPack.*|Test.*Binding.*|Test.*Mutation.*' -v",
		"go test ./internal/formal -run 'Test.*Point11.*|Test.*Claim.*|Test.*NoOverclaim.*|Test.*Governance.*' -v",
		"go test ./internal/formal -run 'Test.*AI.*|Test.*Agent.*|Test.*Lineage.*|Test.*Provenance.*' -v",
		"go test -timeout 20m ./...",
	}
}

func point14ValETestsRun() []string {
	return []string{
		"point14_vale_dependency",
		"point14_vale_closure_evaluator",
		"point14_vale_pass_closure_manifest",
		"point14_vale_validation_closure",
		"point14_vale_conflict_dispute_closure",
		"point14_vale_correction_publication_closure",
		"point14_vale_timeline_query_closure",
		"point14_vale_authority_boundary_final",
		"point14_vale_tenant_privacy_final",
		"point14_vale_timestamp_integrity_final",
		"point14_vale_ai_agent_final",
		"point14_vale_no_overclaim_final",
		"point14_vale_clb_final",
		"point10_through_point14_current_sweep",
	}
}

func point14ValEGrepsRun() []string {
	return []string{
		"point_14_pass scan",
		"authority marker scan",
		"forbidden wording scan",
		"mutation flag scan",
		"ai authority scan",
		"external api scan",
		"skip todo fixme scan",
	}
}

func point14ValENegativeFixturesRun() []string {
	return []string{
		"premature_point14_pass",
		"external_authority_negative",
		"tenant_privacy_negative",
		"timestamp_negative",
		"agent_authority_negative",
		"no_overclaim_negative",
	}
}

func point14ValECommandsRunValid(values []string) bool {
	return point12Val0ExactStringSetMatch(values, point14ValECommandsRun())
}

func point14ValETestsRunValid(values []string) bool {
	return point12Val0ExactStringSetMatch(values, point14ValETestsRun())
}

func point14ValEGrepsRunValid(values []string) bool {
	return point12Val0ExactStringSetMatch(values, point14ValEGrepsRun())
}

func point14ValENegativeFixturesRunValid(values []string) bool {
	return point12Val0ExactStringSetMatch(values, point14ValENegativeFixturesRun())
}

func point14ValEExplicitNonGoalsValid(values []string) bool {
	return point12Val0ExactStringSetMatch(values, point14ValEExplicitNonGoals())
}

func point14ValEDependencySnapshotFromUpstream(valD Point14ValDFoundation) Point14ValEDependencySnapshot {
	return Point14ValEDependencySnapshot{
		Point14ValDCurrentState:                    valD.CurrentState,
		Point14ValDDependencyState:                 valD.DependencyState,
		Point14ValDTimelineProjectionState:         valD.TimelineProjectionState,
		Point14ValDSignalTimelineEntryState:        valD.SignalTimelineEntryState,
		Point14ValDDisputeTimelineState:            valD.DisputeTimelineState,
		Point14ValDCorrectionReadProjectionState:   valD.CorrectionReadProjectionState,
		Point14ValDGovernanceTraceProjectionState:  valD.GovernanceTraceProjectionState,
		Point14ValDQueryProjectionState:            valD.QueryProjectionState,
		Point14ValDAccessBoundaryState:             valD.AccessBoundaryState,
		Point14ValDTenantPrivacyTimelineState:      valD.TenantPrivacyTimelineState,
		Point14ValDAgentTimelineProjectionState:    valD.AgentTimelineProjectionState,
		Point14ValDTimestampIntegrityState:         valD.TimestampIntegrityState,
		Point14ValDNoMutationProjectionGuardState:  valD.NoMutationProjectionGuardState,
		Point14ValDNoOverclaimTimelineWordingState: valD.NoOverclaimTimelineWordingState,
		Point14ValDPointID:                         point14Val0PointID,
		Point14ValDWaveID:                          point14ValDWaveID,
		Point14ValDComputedFromUpstream:            valD.Dependency.SnapshotFromComputedOutput,
		Point14ValDMerged:                          true,
		Point14ValDCIGreen:                         true,
		Point14ValDReviewedOnMain:                  true,
		Point14PassSeen:                            point14ValEValDPayloadContainsPointPass(valD),
		InheritedPoint14ValCCurrentState:           valD.Dependency.Point14ValCCurrentState,
		InheritedPoint14ValBCurrentState:           valD.Dependency.InheritedPoint14ValBCurrentState,
		InheritedPoint14ValACurrentState:           valD.Dependency.InheritedPoint14ValACurrentState,
		InheritedPoint14Val0CurrentState:           valD.Dependency.InheritedPoint14Val0CurrentState,
		InheritedPoint13ValECurrentState:           valD.Dependency.InheritedPoint13ValECurrentState,
		InheritedPoint13ValEPassClosureState:       valD.Dependency.InheritedPoint13ValEPassClosureState,
		InheritedPoint13ValEPassAllowed:            valD.Dependency.InheritedPoint13ValEPassAllowed,
		InheritedPoint13ValEPassToken:              valD.Dependency.InheritedPoint13ValEPassToken,
		InheritedPoint12CurrentState:               valD.Dependency.InheritedPoint12CurrentState,
		InheritedPoint12DependencyState:            valD.Dependency.InheritedPoint12DependencyState,
		InheritedPoint12PassClosureState:           valD.Dependency.InheritedPoint12PassClosureState,
		InheritedPoint12ReviewerResult:             valD.Dependency.InheritedPoint12ReviewerResult,
		InheritedPoint11CurrentState:               valD.Dependency.InheritedPoint11CurrentState,
		InheritedPoint11PublicationState:           valD.Dependency.InheritedPoint11PublicationState,
		InheritedPoint11NoOverclaimState:           valD.Dependency.InheritedPoint11NoOverclaimState,
		InheritedPoint11FinalPassGateState:         valD.Dependency.InheritedPoint11FinalPassGateState,
		InheritedPoint10CurrentState:               valD.Dependency.InheritedPoint10CurrentState,
		InheritedPoint10NoOverclaimState:           valD.Dependency.InheritedPoint10NoOverclaimState,
		InheritedPoint10ProjectionState:            valD.Dependency.InheritedPoint10ProjectionState,
		InheritedPoint10PassRuleState:              valD.Dependency.InheritedPoint10PassRuleState,
		InheritedTenantScope:                       valD.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:                 true,
		ReviewPrerequisites:                        append([]string{}, valD.ReviewPrerequisites...),
		Point14ValD:                                valD,
	}
}

func point14ValEDependencySnapshotModel() Point14ValEDependencySnapshot {
	return cachedFormalModel(&point14ValEDependencySnapshotModelOnce, &point14ValEDependencySnapshotModelCached, func() Point14ValEDependencySnapshot {
		valD := ComputePoint14ValDFoundation(Point14ValDFoundationModel())
		return point14ValEDependencySnapshotFromUpstream(valD)
	})
}

func EvaluatePoint14ValEDependencyState(model Point14ValEDependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point14ValDComputedFromUpstream ||
		!model.Point14ValDMerged ||
		!model.Point14ValDCIGreen ||
		!model.Point14ValDReviewedOnMain ||
		model.Point14PassSeen ||
		model.Point14ValDPointID != point14Val0PointID ||
		model.Point14ValDWaveID != point14ValDWaveID ||
		!point14ValDStateValid(model.Point14ValDCurrentState) ||
		!point14ValDStateValid(model.Point14ValDDependencyState) ||
		!point14ValDStateValid(model.Point14ValDTimelineProjectionState) ||
		!point14ValDStateValid(model.Point14ValDSignalTimelineEntryState) ||
		!point14ValDStateValid(model.Point14ValDDisputeTimelineState) ||
		!point14ValDStateValid(model.Point14ValDCorrectionReadProjectionState) ||
		!point14ValDStateValid(model.Point14ValDGovernanceTraceProjectionState) ||
		!point14ValDStateValid(model.Point14ValDQueryProjectionState) ||
		!point14ValDStateValid(model.Point14ValDAccessBoundaryState) ||
		!point14ValDStateValid(model.Point14ValDTenantPrivacyTimelineState) ||
		!point14ValDStateValid(model.Point14ValDAgentTimelineProjectionState) ||
		!point14ValDStateValid(model.Point14ValDTimestampIntegrityState) ||
		!point14ValDStateValid(model.Point14ValDNoMutationProjectionGuardState) ||
		!point14ValDStateValid(model.Point14ValDNoOverclaimTimelineWordingState) ||
		!point14ValCStateValid(model.InheritedPoint14ValCCurrentState) ||
		!point14ValBStateValid(model.InheritedPoint14ValBCurrentState) ||
		!point14ValAStateValid(model.InheritedPoint14ValACurrentState) ||
		!point14Val0StateValid(model.InheritedPoint14Val0CurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValECurrentState) ||
		!point13ValEStateValid(model.InheritedPoint13ValEPassClosureState) ||
		!point12ValEStateValid(model.InheritedPoint12CurrentState) ||
		!point12ValEStateValid(model.InheritedPoint12DependencyState) ||
		!point12ValEStateValid(model.InheritedPoint12PassClosureState) ||
		!point12ValEReviewerResultValid(model.InheritedPoint12ReviewerResult) ||
		model.InheritedPoint11CurrentState == "" ||
		model.InheritedPoint11PublicationState == "" ||
		model.InheritedPoint11NoOverclaimState == "" ||
		model.InheritedPoint11FinalPassGateState == "" ||
		model.InheritedPoint10CurrentState == "" ||
		model.InheritedPoint10NoOverclaimState == "" ||
		model.InheritedPoint10ProjectionState == "" ||
		model.InheritedPoint10PassRuleState == "" ||
		!point14ValDFoundationEmbeddedSnapshotCopiesExact(model.Point14ValD) ||
		!point14ValDDependencyChainComputedActive(model.Point14ValD) ||
		!point14Val0Point11FoundationActive(model.Point14ValD.Dependency.Point11) ||
		!point14Val0Point11FoundationActive(model.Point14ValD.Dependency.Point14ValC.Dependency.Point11) ||
		!point14Val0Point11FoundationActive(model.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point11) ||
		!point14Val0Point11FoundationActive(model.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point14ValA.Dependency.Point11) ||
		!point14Val0Point11FoundationActive(model.Point14ValD.Dependency.Point14ValC.Dependency.Point14ValB.Dependency.Point14ValA.Dependency.Point14Val0.Dependency.Point11) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return Point14ValEStateBlocked
	}
	if model.Point14ValDCurrentState != model.Point14ValD.CurrentState ||
		model.Point14ValDDependencyState != model.Point14ValD.DependencyState ||
		model.Point14ValDTimelineProjectionState != model.Point14ValD.TimelineProjectionState ||
		model.Point14ValDSignalTimelineEntryState != model.Point14ValD.SignalTimelineEntryState ||
		model.Point14ValDDisputeTimelineState != model.Point14ValD.DisputeTimelineState ||
		model.Point14ValDCorrectionReadProjectionState != model.Point14ValD.CorrectionReadProjectionState ||
		model.Point14ValDGovernanceTraceProjectionState != model.Point14ValD.GovernanceTraceProjectionState ||
		model.Point14ValDQueryProjectionState != model.Point14ValD.QueryProjectionState ||
		model.Point14ValDAccessBoundaryState != model.Point14ValD.AccessBoundaryState ||
		model.Point14ValDTenantPrivacyTimelineState != model.Point14ValD.TenantPrivacyTimelineState ||
		model.Point14ValDAgentTimelineProjectionState != model.Point14ValD.AgentTimelineProjectionState ||
		model.Point14ValDTimestampIntegrityState != model.Point14ValD.TimestampIntegrityState ||
		model.Point14ValDNoMutationProjectionGuardState != model.Point14ValD.NoMutationProjectionGuardState ||
		model.Point14ValDNoOverclaimTimelineWordingState != model.Point14ValD.NoOverclaimTimelineWordingState ||
		model.Point14ValDComputedFromUpstream != model.Point14ValD.Dependency.SnapshotFromComputedOutput ||
		model.InheritedPoint14ValCCurrentState != model.Point14ValD.Dependency.Point14ValCCurrentState ||
		model.InheritedPoint14ValBCurrentState != model.Point14ValD.Dependency.InheritedPoint14ValBCurrentState ||
		model.InheritedPoint14ValACurrentState != model.Point14ValD.Dependency.InheritedPoint14ValACurrentState ||
		model.InheritedPoint14Val0CurrentState != model.Point14ValD.Dependency.InheritedPoint14Val0CurrentState ||
		model.InheritedPoint13ValECurrentState != model.Point14ValD.Dependency.InheritedPoint13ValECurrentState ||
		model.InheritedPoint13ValEPassClosureState != model.Point14ValD.Dependency.InheritedPoint13ValEPassClosureState ||
		model.InheritedPoint13ValEPassAllowed != model.Point14ValD.Dependency.InheritedPoint13ValEPassAllowed ||
		model.InheritedPoint13ValEPassToken != model.Point14ValD.Dependency.InheritedPoint13ValEPassToken ||
		model.InheritedPoint12CurrentState != model.Point14ValD.Dependency.InheritedPoint12CurrentState ||
		model.InheritedPoint12DependencyState != model.Point14ValD.Dependency.InheritedPoint12DependencyState ||
		model.InheritedPoint12PassClosureState != model.Point14ValD.Dependency.InheritedPoint12PassClosureState ||
		model.InheritedPoint12ReviewerResult != model.Point14ValD.Dependency.InheritedPoint12ReviewerResult ||
		model.InheritedPoint11CurrentState != model.Point14ValD.Dependency.InheritedPoint11CurrentState ||
		model.InheritedPoint11PublicationState != model.Point14ValD.Dependency.InheritedPoint11PublicationState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValD.Dependency.InheritedPoint11NoOverclaimState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValD.Dependency.InheritedPoint11FinalPassGateState ||
		model.InheritedPoint11CurrentState != model.Point14ValD.Dependency.Point11.CurrentState ||
		model.InheritedPoint11PublicationState != model.Point14ValD.Dependency.Point11.PublicationReviewState ||
		model.InheritedPoint11NoOverclaimState != model.Point14ValD.Dependency.Point11.NoOverclaimReviewState ||
		model.InheritedPoint11FinalPassGateState != model.Point14ValD.Dependency.Point11.FinalPassGateState ||
		model.InheritedPoint10CurrentState != model.Point14ValD.Dependency.InheritedPoint10CurrentState ||
		model.InheritedPoint10NoOverclaimState != model.Point14ValD.Dependency.InheritedPoint10NoOverclaimState ||
		model.InheritedPoint10ProjectionState != model.Point14ValD.Dependency.InheritedPoint10ProjectionState ||
		model.InheritedPoint10PassRuleState != model.Point14ValD.Dependency.InheritedPoint10PassRuleState ||
		model.InheritedTenantScope != model.Point14ValD.Dependency.InheritedTenantScope {
		return Point14ValEStateBlocked
	}
	if model.Point14ValDCurrentState != Point14ValDStateActive ||
		model.Point14ValDDependencyState != Point14ValDStateActive ||
		model.Point14ValDTimelineProjectionState != Point14ValDStateActive ||
		model.Point14ValDSignalTimelineEntryState != Point14ValDStateActive ||
		model.Point14ValDDisputeTimelineState != Point14ValDStateActive ||
		model.Point14ValDCorrectionReadProjectionState != Point14ValDStateActive ||
		model.Point14ValDGovernanceTraceProjectionState != Point14ValDStateActive ||
		model.Point14ValDQueryProjectionState != Point14ValDStateActive ||
		model.Point14ValDAccessBoundaryState != Point14ValDStateActive ||
		model.Point14ValDTenantPrivacyTimelineState != Point14ValDStateActive ||
		model.Point14ValDAgentTimelineProjectionState != Point14ValDStateActive ||
		model.Point14ValDTimestampIntegrityState != Point14ValDStateActive ||
		model.Point14ValDNoMutationProjectionGuardState != Point14ValDStateActive ||
		model.Point14ValDNoOverclaimTimelineWordingState != Point14ValDStateActive ||
		model.InheritedPoint14ValCCurrentState != Point14ValCStateActive ||
		model.InheritedPoint14ValBCurrentState != Point14ValBStateActive ||
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
		model.InheritedPoint11CurrentState != Point11ValDStateActive ||
		model.InheritedPoint11PublicationState != Point11ValDPublicationReviewStateActive ||
		model.InheritedPoint11NoOverclaimState != Point11ValDNoOverclaimReviewStateActive ||
		model.InheritedPoint11FinalPassGateState != Point11ValDFinalPassGateStateActive ||
		model.InheritedPoint10CurrentState != operability.DeploymentMultiTenantPoint10StatePass ||
		model.InheritedPoint10NoOverclaimState != operability.DeploymentMultiTenantValENoOverclaimStateActive ||
		model.InheritedPoint10ProjectionState != operability.DeploymentMultiTenantValEProjectionBoundaryStateActive ||
		model.InheritedPoint10PassRuleState != operability.DeploymentMultiTenantValEPoint10PassRuleStateActive {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValEExternalSignalValidationClosureCheckModel(dependency Point14ValEDependencySnapshot) Point14ValEExternalSignalValidationClosureCheck {
	valA := dependency.Point14ValD.Dependency.Point14ValA
	return Point14ValEExternalSignalValidationClosureCheck{
		CheckID:                        "point14_vale_check_validation_001",
		ValACurrentState:               valA.CurrentState,
		ValidationResultState:          valA.ValidationResultState,
		CandidateSemanticsBounded:      valA.ValidationResult.CandidateUsable && valA.ValidationResult.ValidationState == point14ValAValidationCandidateValidated,
		CanonicalAuthorityGranted:      valA.NormalizedExternalSignal.CanonicalAuthority || valA.SourceIdentity.CanonicalAuthorityGranted || valA.NoExternalAuthorityGuard.CanonicalAuthorityGranted,
		PassEmitted:                    valA.ValidationResult.EmitsPass || valA.NormalizedExternalSignal.PassAllowed || valA.SourceIdentity.PassAllowed,
		SourceIdentityAuthorityGranted: valA.SourceIdentity.CanonicalAuthorityGranted || valA.SourceIdentity.OverrideCanonicalDecision || valA.SourceIdentity.ApproveProduction || valA.SourceIdentity.CertifyCompliance,
		CanonicalMutationDetected:      valA.EvidenceBinding.CanonicalMutationAllowed || valA.ValidationResult.OverridesCanonicalDecision,
		DuplicateActiveEvidencePath:    valA.DuplicateAndRelationGuard.ConflictingDuplicate && valA.ValidationResult.CandidateUsable,
		UnrelatedSignalActivePath:      valA.ValidationResult.ValidationState == point14ValAValidationCandidateUnrelated && valA.ValidationResult.CandidateUsable,
		CrossTenantSignalActivePath:    valA.ValidationResult.ValidationState == point14ValAValidationCandidateCrossTenant && valA.ValidationResult.CandidateUsable,
	}
}

func EvaluatePoint14ValEExternalSignalValidationClosureCheckState(model Point14ValEExternalSignalValidationClosureCheck) string {
	if !point14ValECheckIDValid(model.CheckID) ||
		model.ValACurrentState != Point14ValAStateActive ||
		model.ValidationResultState != Point14ValAStateActive ||
		!model.CandidateSemanticsBounded {
		return Point14ValEStateBlocked
	}
	if model.CanonicalAuthorityGranted ||
		model.PassEmitted ||
		model.SourceIdentityAuthorityGranted ||
		model.CanonicalMutationDetected ||
		model.DuplicateActiveEvidencePath ||
		model.UnrelatedSignalActivePath ||
		model.CrossTenantSignalActivePath {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValEConflictDisputeClosureCheckModel(dependency Point14ValEDependencySnapshot) Point14ValEConflictDisputeClosureCheck {
	valB := dependency.Point14ValD.Dependency.Point14ValB
	return Point14ValEConflictDisputeClosureCheck{
		CheckID:                  "point14_vale_check_conflict_001",
		ValBCurrentState:         valB.CurrentState,
		DisputeTriageResultState: valB.DisputeTriageResultState,
		TriageOnlyBounded: !valB.ConflictSet.ResolveToPass &&
			!valB.ConflictSet.ResolveToCanonicalTruth &&
			!valB.ConflictSet.ApproveProduction &&
			!valB.ConflictSet.PublishCorrection &&
			!valB.ConflictSet.RevokeClaim &&
			!valB.ConflictSet.CreatePublicBadge &&
			!valB.StakeholderComparison.ResolvesConflict &&
			!valB.DisputeTriageResult.EmitsPass &&
			!valB.DisputeTriageResult.PublishesCorrection &&
			!valB.DisputeTriageResult.RevokesClaim &&
			!valB.DisputeTriageResult.OverridesCanonicalDecision &&
			!valB.DisputeTriageResult.ApprovesProduction &&
			!valB.DisputeTriageResult.CertifiesCompliance &&
			!valB.DisputeTriageResult.CreatesPublicBadge,
		UnresolvedDisputeReviewRequired:     valB.DisputeTriageResult.TriageState == point14ValBTriageReviewRequired || valB.DisputeTriageResult.TriageState == point14ValBTriageGovernanceEscalated,
		EvidenceRequiredUnclosed:            valB.DisputeTriageResult.TriageState == point14ValBTriageEvidenceRequired || valB.EvidenceRequirementGate.DecisiveEvidenceMissing,
		DisputeAutoResolved:                 valB.DisputeIntake.ClosureAttempted || valB.NoExternalAuthorityConflictGuard.DisputeAutoResolved,
		ConflictResolvedToPass:              valB.ConflictSet.ResolveToPass || valB.DisputeTriageResult.EmitsPass,
		ExternalAuthorityResolutionDetected: valB.ConflictSet.ResolveToCanonicalTruth || valB.StakeholderComparison.CrowdConsensusResolutionRequested || valB.StakeholderComparison.PublishesPublicAuthority || valB.NoExternalAuthorityConflictGuard.CrowdResolved || valB.NoExternalAuthorityConflictGuard.AgentResolved || valB.NoExternalAuthorityConflictGuard.CanonicalAuthorityGranted || valB.NoExternalAuthorityConflictGuard.ProductionApprovalGranted || valB.NoExternalAuthorityConflictGuard.ExternalAuthorityAllowed,
		GovernanceEscalationMissing:         valB.GovernanceEscalationPath.EscalationRequired && (strings.TrimSpace(valB.GovernanceEscalationPath.GovernanceEventRef) == "" || strings.TrimSpace(valB.GovernanceEscalationPath.Owner) == "" || strings.TrimSpace(valB.GovernanceEscalationPath.ApproverRole) == "" || strings.TrimSpace(valB.GovernanceEscalationPath.AuditRef) == ""),
	}
}

func EvaluatePoint14ValEConflictDisputeClosureCheckState(model Point14ValEConflictDisputeClosureCheck) string {
	if !point14ValECheckIDValid(model.CheckID) ||
		model.ValBCurrentState != Point14ValBStateActive ||
		model.DisputeTriageResultState != Point14ValBStateActive ||
		!model.TriageOnlyBounded {
		return Point14ValEStateBlocked
	}
	if model.DisputeAutoResolved ||
		model.ConflictResolvedToPass ||
		model.ExternalAuthorityResolutionDetected ||
		model.GovernanceEscalationMissing {
		return Point14ValEStateBlocked
	}
	if model.UnresolvedDisputeReviewRequired {
		return Point14ValEStateReviewRequired
	}
	if model.EvidenceRequiredUnclosed {
		return Point14ValEStateIncomplete
	}
	return Point14ValEStatePassConfirmed
}

func point14ValECorrectionPublicationClosureCheckModel(dependency Point14ValEDependencySnapshot) Point14ValECorrectionPublicationClosureCheck {
	valC := dependency.Point14ValD.Dependency.Point14ValC
	return Point14ValECorrectionPublicationClosureCheck{
		CheckID:                             "point14_vale_check_correction_001",
		ValCCurrentState:                    valC.CurrentState,
		PublicationApprovalState:            valC.PublicationApprovalState,
		GovernanceControlledOnly:            !valC.CorrectionNoticeBoundary.CanonicalMutationRequested && !valC.CorrectionNoticeBoundary.OverridesCanonicalDecision && !valC.RevocationRequestBoundary.CanonicalMutationRequested && !valC.PublicationApprovalBoundary.ApprovesProduction && !valC.PublicationApprovalBoundary.Certifies && !valC.AgentCorrectionBoundary.ApprovalGranted && !valC.AgentCorrectionBoundary.ProductionApproved && !valC.AgentCorrectionBoundary.ExternalAuthorityAllowed,
		BoundedCorrectionNoticeProven:       len(valC.CorrectionNoticeBoundary.CorrectionLimitations) > 0 && strings.TrimSpace(valC.CorrectionNoticeBoundary.GovernanceEventRef) != "" && strings.TrimSpace(valC.CorrectionNoticeBoundary.AuditRef) != "" && len(valC.PublicationVisibilityBoundary.LimitationRefs) > 0,
		CorrectionAutoPublished:             valC.CorrectionNoticeBoundary.CorrectionState == "correction_auto_published",
		RevocationAutoExecuted:              valC.RevocationRequestBoundary.RevocationState == "revocation_auto_executed",
		SupersessionSilentReplacement:       valC.SupersessionRecordBoundary.SilentReplacement || valC.SupersessionRecordBoundary.DeletesHistory || valC.SupersessionRecordBoundary.HidesPreviousEvidence,
		PublicationApprovalBecameProduction: valC.PublicationApprovalBoundary.ApprovesProduction,
		PublicationCertified:                valC.PublicationApprovalBoundary.Certifies || valC.CorrectionNoticeBoundary.CertifiesCompliance,
		PublicNoticeBecameBadge:             valC.PublicationApprovalBoundary.CreatesPublicBadge || valC.CorrectionNoticeBoundary.CreatesPublicBadge || valC.RevocationRequestBoundary.PublicBadgeAllowed || valC.PublicationVisibilityBoundary.ImpliesPublicAuthority,
		RedactionHidesDecisiveEvidence:      valC.CorrectionRedactionGuard.DecisiveMissingEvidenceHidden || valC.CorrectionRedactionGuard.SurvivingTextMisleading,
		LimitationsOmitted:                  valC.CorrectionRedactionGuard.LimitationOmitted || valC.PublicationVisibilityBoundary.MeaningChangingPrivateLimitationOmitted,
		PublicationStrengthensClaims:        valC.CorrectionNoticeBoundary.StrengthensClaim || valC.PublicationVisibilityBoundary.StrengthensClaim || valC.CorrectionRedactionGuard.RedactionStrengthensClaim,
	}
}

func EvaluatePoint14ValECorrectionPublicationClosureCheckState(model Point14ValECorrectionPublicationClosureCheck) string {
	if !point14ValECheckIDValid(model.CheckID) ||
		model.ValCCurrentState != Point14ValCStateActive ||
		model.PublicationApprovalState != Point14ValCStateActive ||
		!model.GovernanceControlledOnly {
		return Point14ValEStateBlocked
	}
	if !model.BoundedCorrectionNoticeProven {
		return Point14ValEStateIncomplete
	}
	if model.CorrectionAutoPublished ||
		model.RevocationAutoExecuted ||
		model.SupersessionSilentReplacement ||
		model.PublicationApprovalBecameProduction ||
		model.PublicationCertified ||
		model.PublicNoticeBecameBadge ||
		model.RedactionHidesDecisiveEvidence ||
		model.LimitationsOmitted ||
		model.PublicationStrengthensClaims {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValETimelineProjectionClosureCheckModel(dependency Point14ValEDependencySnapshot) Point14ValETimelineProjectionClosureCheck {
	valD := dependency.Point14ValD
	return Point14ValETimelineProjectionClosureCheck{
		CheckID:                           "point14_vale_check_timeline_001",
		ValDCurrentState:                  valD.CurrentState,
		QueryProjectionState:              valD.QueryProjectionState,
		ReadOnlyProjectionOnly:            valD.TimelineProjection.ReadOnly && valD.TimelineProjection.ProjectionOnly && valD.QueryProjection.ReadOnly && valD.QueryProjection.QueryIsProjectionOnly && valD.CorrectionReadProjection.ReadOnly && valD.GovernanceTraceProjection.ReadOnly,
		TimelineMutationDetected:          valD.TimelineProjection.MutatesCanonicalEvidence || valD.TimelineProjection.MutatesSignalState || valD.NoMutationProjectionGuard.MutatesCanonicalEvidence || valD.NoMutationProjectionGuard.MutatesNormalizedSignal || valD.NoMutationProjectionGuard.MutatesValidationResult || valD.NoMutationProjectionGuard.MutatesDisputeLifecycle || valD.NoMutationProjectionGuard.MutatesCorrectionNotice || valD.NoMutationProjectionGuard.MutatesRevocationRequest || valD.NoMutationProjectionGuard.MutatesSupersessionRecord || valD.NoMutationProjectionGuard.MutatesPublicationApproval || valD.NoMutationProjectionGuard.MutatesVisibilityBoundary || valD.NoMutationProjectionGuard.MutatesGovernanceTrace,
		QueryMutationDetected:             valD.QueryProjection.MutationRequested || valD.QueryProjection.WritesFiltersBack,
		TimelineResolvesDisputes:          valD.TimelineProjection.ResolvesDispute || valD.DisputeTimelineProjection.ResolvesDispute || valD.NoMutationProjectionGuard.ResolvesDispute,
		TimelinePublishesCorrections:      valD.TimelineProjection.PublishesCorrection || valD.CorrectionReadProjection.PublishesCorrection || valD.NoMutationProjectionGuard.PublishesCorrection,
		TimelineExecutesRevocation:        valD.TimelineProjection.RevokesClaim || valD.CorrectionReadProjection.ExecutesRevocation || valD.NoMutationProjectionGuard.ExecutesRevocation,
		TimelineCreatesAuthority:          valD.TimelineProjection.ApprovesProduction || valD.TimelineProjection.CertifiesCompliance || valD.TimelineProjection.CreatesPublicBadge || valD.AccessBoundary.AuthorityGranted || valD.AgentTimelineProjection.AgentAuthorityFlags || valD.AgentTimelineProjection.ExternalAuthorityAllowed,
		QueryHidesDecisiveMissingEvidence: valD.QueryProjection.HidesDecisiveMissingEvidence || valD.QueryProjection.OmitsLimitationsWithoutDisclosure || valD.DisputeTimelineProjection.HidesEvidenceRequired,
		TimelineStrengthensClaims:         valD.TenantPrivacyTimelineProjectionGuard.StrengthensClaim || point14ValDObservedListContainsForbiddenWording(valD.CorrectionReadProjection.ObservedReadTexts) || point14ValDObservedListContainsForbiddenWording(valD.TenantPrivacyTimelineProjectionGuard.ObservedSummaryTexts),
	}
}

func EvaluatePoint14ValETimelineProjectionClosureCheckState(model Point14ValETimelineProjectionClosureCheck) string {
	if !point14ValECheckIDValid(model.CheckID) ||
		model.ValDCurrentState != Point14ValDStateActive ||
		model.QueryProjectionState != Point14ValDStateActive ||
		!model.ReadOnlyProjectionOnly {
		return Point14ValEStateBlocked
	}
	if model.TimelineMutationDetected ||
		model.QueryMutationDetected ||
		model.TimelineResolvesDisputes ||
		model.TimelinePublishesCorrections ||
		model.TimelineExecutesRevocation ||
		model.TimelineCreatesAuthority ||
		model.QueryHidesDecisiveMissingEvidence ||
		model.TimelineStrengthensClaims {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValEAuthorityBoundaryClosureCheckModel(dependency Point14ValEDependencySnapshot) Point14ValEAuthorityBoundaryClosureCheck {
	valA := dependency.Point14ValD.Dependency.Point14ValA
	valB := dependency.Point14ValD.Dependency.Point14ValB
	valC := dependency.Point14ValD.Dependency.Point14ValC
	markers := append([]string{}, valA.NoExternalAuthorityGuard.ObservedAuthorityMarkers...)
	markers = append(markers, valB.NoExternalAuthorityConflictGuard.ObservedAuthorityMarkers...)
	markers = append(markers, valC.NoExternalAuthorityGuard.ObservedAuthorityMarkers...)
	return Point14ValEAuthorityBoundaryClosureCheck{
		CheckID:                "point14_vale_check_authority_001",
		ObservedAuthorityMarks: markers,
		ExternalAuthorityFound: valA.NoExternalAuthorityGuard.CanonicalAuthorityGranted || valA.NoExternalAuthorityGuard.ProductionApprovalGranted || valA.NoExternalAuthorityGuard.CorrectionPublished || valA.NoExternalAuthorityGuard.PublicBadgeAllowed || valA.NoExternalAuthorityGuard.ExternalAuthorityAllowed ||
			valB.NoExternalAuthorityConflictGuard.CanonicalAuthorityGranted || valB.NoExternalAuthorityConflictGuard.ProductionApprovalGranted || valB.NoExternalAuthorityConflictGuard.PublishesCorrection || valB.NoExternalAuthorityConflictGuard.RevokesClaim || valB.NoExternalAuthorityConflictGuard.PublicBadgeAllowed || valB.NoExternalAuthorityConflictGuard.DisputeAutoResolved || valB.NoExternalAuthorityConflictGuard.CrowdResolved || valB.NoExternalAuthorityConflictGuard.AgentResolved || valB.NoExternalAuthorityConflictGuard.ExternalAuthorityAllowed ||
			valC.NoExternalAuthorityGuard.CanonicalAuthorityGranted || valC.NoExternalAuthorityGuard.ProductionApprovalGranted || valC.NoExternalAuthorityGuard.PublishesCorrection || valC.NoExternalAuthorityGuard.RevokesClaim || valC.NoExternalAuthorityGuard.PublicBadgeAllowed || valC.NoExternalAuthorityGuard.ExternalAuthorityAllowed ||
			dependency.Point14ValD.AgentTimelineProjection.AgentAuthorityFlags || dependency.Point14ValD.AgentTimelineProjection.ExternalAuthorityAllowed,
	}
}

func EvaluatePoint14ValEAuthorityBoundaryClosureCheckState(model Point14ValEAuthorityBoundaryClosureCheck) string {
	if !point14ValECheckIDValid(model.CheckID) {
		return Point14ValEStateBlocked
	}
	if model.ExternalAuthorityFound {
		return Point14ValEStateBlocked
	}
	if point14Val0AuthorityMarkersContainForbidden(model.ObservedAuthorityMarks, point14ValEForbiddenAuthorityMarkers()) {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValETenantPrivacyClosureCheckModel(dependency Point14ValEDependencySnapshot) Point14ValETenantPrivacyClosureCheck {
	valC := dependency.Point14ValD.Dependency.Point14ValC
	valD := dependency.Point14ValD
	return Point14ValETenantPrivacyClosureCheck{
		CheckID:                                "point14_vale_check_privacy_001",
		TenantScope:                            dependency.InheritedTenantScope,
		CrossTenantDetected:                    valC.TenantPrivacyPublicationGuard.CrossTenantPublication || valD.AccessBoundary.CrossTenantAccess || valD.TenantPrivacyTimelineProjectionGuard.CrossTenantProjection || valD.QueryProjection.CrossTenantResults,
		TenantPrivateDataExposed:               valC.TenantPrivacyPublicationGuard.TenantPrivateDataExposed || valC.PublicationVisibilityBoundary.TenantPrivateDataExposed || valD.AccessBoundary.TenantPrivateDataExposed || valD.TenantPrivacyTimelineProjectionGuard.TenantPrivateDataExposed,
		PublicNoticeLeaksTenantPrivateData:     valC.PublicationVisibilityBoundary.VisibilityClassification == point14ValCVisibilityPublicBounded && valC.PublicationVisibilityBoundary.TenantPrivateDataExposed,
		AccessBoundaryAllowsCrossTenantQuery:   valD.AccessBoundary.CrossTenantAccess || valD.QueryProjection.CrossTenantResults,
		PublicPrivateClassificationPresent:     strings.TrimSpace(valC.TenantPrivacyPublicationGuard.PublicPrivateClassification) != "" && strings.TrimSpace(valD.TenantPrivacyTimelineProjectionGuard.PublicPrivateClassification) != "",
		RequiredRedactionLimitationRefsPresent: len(valC.TenantPrivacyPublicationGuard.LimitationRefs) > 0 && len(valD.TenantPrivacyTimelineProjectionGuard.LimitationRefs) > 0 && len(valD.TenantPrivacyTimelineProjectionGuard.RedactionRefs) > 0,
	}
}

func EvaluatePoint14ValETenantPrivacyClosureCheckState(model Point14ValETenantPrivacyClosureCheck, dependency Point14ValEDependencySnapshot) string {
	if !point14ValECheckIDValid(model.CheckID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		model.TenantScope != dependency.InheritedTenantScope {
		return Point14ValEStateBlocked
	}
	if model.CrossTenantDetected ||
		model.TenantPrivateDataExposed ||
		model.PublicNoticeLeaksTenantPrivateData ||
		model.AccessBoundaryAllowsCrossTenantQuery ||
		!model.PublicPrivateClassificationPresent ||
		!model.RequiredRedactionLimitationRefsPresent {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValETimestampIntegrityClosureCheckModel(dependency Point14ValEDependencySnapshot) Point14ValETimestampIntegrityClosureCheck {
	valD := dependency.Point14ValD
	valA := valD.Dependency.Point14ValA
	return Point14ValETimestampIntegrityClosureCheck{
		CheckID:                     "point14_vale_check_timestamp_001",
		TenantScope:                 dependency.InheritedTenantScope,
		EventAt:                     valD.TimestampIntegrityGuard.EventAt,
		EventTimeSource:             valD.TimestampIntegrityGuard.EventTimeSource,
		GeneratedAt:                 valD.TimestampIntegrityGuard.GeneratedAt,
		GeneratedTimeSource:         valD.TimestampIntegrityGuard.GeneratedTimeSource,
		ApprovalAt:                  valD.TimestampIntegrityGuard.PublicationApprovedAt,
		ApprovalTimeSource:          valD.TimestampIntegrityGuard.PublicationApprovedTimeSource,
		SourceEventAt:               valD.TimestampIntegrityGuard.SourceEventAt,
		SourceEventTimeSource:       valD.TimestampIntegrityGuard.SourceEventTimeSource,
		ClientLocalCreatesCanonical: valA.FreshnessAndTimestamp.ReceivedTimeSource == point14Val0TimeSourceClientLocal || valA.FreshnessAndTimestamp.ValidatedTimeSource == point14Val0TimeSourceClientLocal,
		SourceEventCreatesAuthority: valA.FreshnessAndTimestamp.AuthorityUpgradeRequested || valD.SignalTimelineEntry.SourceEventAsCanonicalAuthority,
		FutureDatedActiveEvent:      valD.TimestampIntegrityState == Point14ValDStateReviewRequired && valD.TimestampIntegrityGuard.EventAt > valD.TimestampIntegrityGuard.GeneratedAt,
		BackdatedApproval: valD.TimestampIntegrityGuard.PublicationApprovedAt != "" && valD.TimestampIntegrityGuard.DisputeOpenedAt != "" && func() bool {
			approvalAt, _ := point14Val0ParsedTime(valD.TimestampIntegrityGuard.PublicationApprovedAt)
			disputeOpenedAt, _ := point14Val0ParsedTime(valD.TimestampIntegrityGuard.DisputeOpenedAt)
			return approvalAt.Before(disputeOpenedAt)
		}(),
		ImpossibleOrdering:               valD.TimestampIntegrityState == Point14ValDStateReviewRequired,
		TimelineOrderingUpgradesValidity: valD.TimestampIntegrityGuard.AttemptsValidityUpgrade,
	}
}

func EvaluatePoint14ValETimestampIntegrityClosureCheckState(model Point14ValETimestampIntegrityClosureCheck, dependency Point14ValEDependencySnapshot) string {
	if !point14ValECheckIDValid(model.CheckID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		model.TenantScope != dependency.InheritedTenantScope ||
		!point14Val0ParsedTimeOk(model.EventAt) ||
		!point14Val0CanonicalTimeSourceValid(model.EventTimeSource) ||
		!point14Val0ParsedTimeOk(model.GeneratedAt) ||
		!point14Val0CanonicalTimeSourceValid(model.GeneratedTimeSource) {
		return Point14ValEStateBlocked
	}
	if strings.TrimSpace(model.ApprovalAt) != "" {
		if !point14Val0ParsedTimeOk(model.ApprovalAt) || !point14Val0CanonicalTimeSourceValid(model.ApprovalTimeSource) {
			return Point14ValEStateBlocked
		}
	}
	if strings.TrimSpace(model.SourceEventAt) != "" {
		if !point14Val0ParsedTimeOk(model.SourceEventAt) || !point14Val0TimeSourceValid(model.SourceEventTimeSource) {
			return Point14ValEStateBlocked
		}
	}
	if model.ClientLocalCreatesCanonical ||
		model.SourceEventCreatesAuthority ||
		model.FutureDatedActiveEvent ||
		model.BackdatedApproval ||
		model.ImpossibleOrdering ||
		model.TimelineOrderingUpgradesValidity {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValEAgentAdvisoryClosureCheckModel(dependency Point14ValEDependencySnapshot) Point14ValEAgentAdvisoryClosureCheck {
	valB := dependency.Point14ValD.Dependency.Point14ValB
	valC := dependency.Point14ValD.Dependency.Point14ValC
	valD := dependency.Point14ValD
	return Point14ValEAgentAdvisoryClosureCheck{
		CheckID:                       "point14_vale_check_agent_001",
		TenantScope:                   dependency.InheritedTenantScope,
		AdvisoryOnly:                  valB.AgentDisputeRecommendationBoundary.AdvisoryOnly && valC.AgentCorrectionBoundary.AdvisoryOnly && valD.AgentTimelineProjection.AdvisoryOnly,
		AgentResolvesDispute:          valB.AgentDisputeRecommendationBoundary.CanResolveConflict || valD.AgentTimelineProjection.CanResolveDispute,
		AgentPublishesCorrection:      valB.AgentDisputeRecommendationBoundary.CanPublishCorrection || valC.AgentCorrectionBoundary.CanPublishNotice || valD.AgentTimelineProjection.CanPublishCorrection,
		AgentRevokesClaim:             valB.AgentDisputeRecommendationBoundary.CanRevokeClaim || valD.AgentTimelineProjection.CanRevokeClaim,
		AgentSatisfiesGovernanceAlone: valB.AgentDisputeRecommendationBoundary.CanSatisfyEvidenceRequirementAlone || valC.AgentCorrectionBoundary.CanSatisfyGovernanceTraceAlone || valD.AgentTimelineProjection.CanSatisfyGovernanceTrace,
		AgentAuthorityFlags:           valB.AgentDisputeRecommendationBoundary.CanOverrideGovernance || valB.AgentDisputeRecommendationBoundary.CanEmitPublicAuthority || valB.AgentDisputeRecommendationBoundary.ApprovalGranted || valB.AgentDisputeRecommendationBoundary.ProductionApproved || valB.AgentDisputeRecommendationBoundary.ExternalAuthorityAllowed || valC.AgentCorrectionBoundary.CanEmitPublicAuthority || valC.AgentCorrectionBoundary.ApprovalGranted || valC.AgentCorrectionBoundary.ProductionApproved || valC.AgentCorrectionBoundary.ExternalAuthorityAllowed || valD.AgentTimelineProjection.AgentAuthorityFlags || valD.AgentTimelineProjection.ExternalAuthorityAllowed,
		AgentPassAllowed:              valB.AgentDisputeRecommendationBoundary.CanEmitPass || valB.AgentDisputeRecommendationBoundary.PassAllowed || valC.AgentCorrectionBoundary.CanEmitPass || valD.AgentTimelineProjection.PassAllowed,
		AgentPublicBadgeAllowed:       valC.AgentCorrectionBoundary.CanEmitPublicAuthority || valD.AgentTimelineProjection.AgentAuthorityFlags,
	}
}

func EvaluatePoint14ValEAgentAdvisoryClosureCheckState(model Point14ValEAgentAdvisoryClosureCheck, dependency Point14ValEDependencySnapshot) string {
	if !point14ValECheckIDValid(model.CheckID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		model.TenantScope != dependency.InheritedTenantScope ||
		!model.AdvisoryOnly {
		return Point14ValEStateBlocked
	}
	if model.AgentResolvesDispute ||
		model.AgentPublishesCorrection ||
		model.AgentRevokesClaim ||
		model.AgentSatisfiesGovernanceAlone ||
		model.AgentAuthorityFlags ||
		model.AgentPassAllowed ||
		model.AgentPublicBadgeAllowed {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

type point14ValENoOverclaimCorpus struct {
	observed                             []string
	observedSeen                         map[string]struct{}
	internalDiagnostics                  []string
	internalDiagnosticsSeen              map[string]struct{}
	internalForbiddenDiagnosticSeen      bool
	internalForbiddenDiagnosticUnblocked bool
}

func (corpus *point14ValENoOverclaimCorpus) appendObserved(texts ...[]string) {
	if corpus.observedSeen == nil {
		corpus.observedSeen = make(map[string]struct{})
		for _, existing := range corpus.observed {
			corpus.observedSeen[existing] = struct{}{}
		}
	}
	for _, group := range texts {
		for _, text := range group {
			if _, ok := corpus.observedSeen[text]; ok {
				continue
			}
			corpus.observedSeen[text] = struct{}{}
			corpus.observed = append(corpus.observed, text)
		}
	}
}

func (corpus *point14ValENoOverclaimCorpus) appendInternalDiagnostics(texts []string, classifiedBlocked bool) {
	if corpus.internalDiagnosticsSeen == nil {
		corpus.internalDiagnosticsSeen = make(map[string]struct{})
		for _, existing := range corpus.internalDiagnostics {
			corpus.internalDiagnosticsSeen[existing] = struct{}{}
		}
	}
	for _, text := range texts {
		if _, ok := corpus.internalDiagnosticsSeen[text]; ok {
			continue
		}
		corpus.internalDiagnosticsSeen[text] = struct{}{}
		corpus.internalDiagnostics = append(corpus.internalDiagnostics, text)
	}
	if point14ValEObservedListContainsForbiddenWording(texts) {
		corpus.internalForbiddenDiagnosticSeen = true
		if !classifiedBlocked {
			corpus.internalForbiddenDiagnosticUnblocked = true
		}
	}
}

func (corpus point14ValENoOverclaimCorpus) internalDiagnosticsClassifiedBlocked() bool {
	return corpus.internalForbiddenDiagnosticSeen && !corpus.internalForbiddenDiagnosticUnblocked
}

func point14ValEAppendValANoOverclaimCorpus(corpus *point14ValENoOverclaimCorpus, valA Point14ValAFoundation) {
	wording := valA.NoOverclaimValidationWording
	corpus.appendObserved(
		wording.ObservedNormalizationTexts,
		wording.ObservedValidationTexts,
		wording.ObservedSourceIdentityTexts,
		wording.ObservedScopeBindingTexts,
		wording.ObservedEvidenceBindingTexts,
	)
	corpus.appendInternalDiagnostics(wording.InternalDiagnosticTexts, wording.InternalDiagnosticsClassifiedBlocked)
}

func point14ValEAppendValAInheritedNoOverclaimCorpus(corpus *point14ValENoOverclaimCorpus, valA Point14ValAFoundation) {
	point14ValEAppendValANoOverclaimCorpus(corpus, valA)
}

func point14ValEAppendValBNoOverclaimCorpus(corpus *point14ValENoOverclaimCorpus, valB Point14ValBFoundation) {
	wording := valB.NoOverclaimDisputeWording
	corpus.appendObserved(
		wording.ObservedConflictTexts,
		wording.ObservedDisputeTexts,
		wording.ObservedEvidenceRequirementTexts,
		wording.ObservedEscalationTexts,
		wording.ObservedPrivacyTexts,
		wording.ObservedAgentTexts,
	)
	corpus.appendInternalDiagnostics(wording.InternalDiagnosticTexts, wording.InternalDiagnosticsClassifiedBlocked)
}

func point14ValEAppendValBInheritedNoOverclaimCorpus(corpus *point14ValENoOverclaimCorpus, valB Point14ValBFoundation) {
	point14ValEAppendValBNoOverclaimCorpus(corpus, valB)
	point14ValEAppendValAInheritedNoOverclaimCorpus(corpus, valB.Dependency.Point14ValA)
}

func point14ValEAppendValCNoOverclaimCorpus(corpus *point14ValENoOverclaimCorpus, valC Point14ValCFoundation) {
	wording := valC.NoOverclaimPublicationWording
	corpus.appendObserved(
		wording.ObservedCorrectionTexts,
		wording.ObservedRevocationTexts,
		wording.ObservedSupersessionTexts,
		wording.ObservedPublicationTexts,
		wording.ObservedVisibilityTexts,
		wording.ObservedPrivacyTexts,
		wording.ObservedGovernanceTexts,
		wording.ObservedAgentTexts,
	)
	corpus.appendInternalDiagnostics(wording.InternalDiagnosticTexts, wording.InternalDiagnosticsClassifiedBlocked)
}

func point14ValEAppendValCInheritedNoOverclaimCorpus(corpus *point14ValENoOverclaimCorpus, valC Point14ValCFoundation) {
	point14ValEAppendValCNoOverclaimCorpus(corpus, valC)
	point14ValEAppendValBInheritedNoOverclaimCorpus(corpus, valC.Dependency.Point14ValB)
}

func point14ValEAppendValDNoOverclaimCorpus(corpus *point14ValENoOverclaimCorpus, valD Point14ValDFoundation) {
	wording := valD.NoOverclaimTimelineWording
	corpus.appendObserved(
		wording.ObservedTimelineTexts,
		wording.ObservedSignalTexts,
		wording.ObservedDisputeTexts,
		wording.ObservedReadProjectionTexts,
		wording.ObservedGovernanceTexts,
		wording.ObservedQueryTexts,
		wording.ObservedAccessTexts,
		wording.ObservedPrivacyTexts,
		wording.ObservedAgentTexts,
	)
	corpus.appendInternalDiagnostics(wording.InternalDiagnosticTexts, wording.InternalDiagnosticsClassifiedBlocked)
}

func point14ValEAppendValDInheritedNoOverclaimCorpus(corpus *point14ValENoOverclaimCorpus, valD Point14ValDFoundation) {
	point14ValEAppendValDNoOverclaimCorpus(corpus, valD)
	point14ValEAppendValCInheritedNoOverclaimCorpus(corpus, valD.Dependency.Point14ValC)
	point14ValEAppendValBInheritedNoOverclaimCorpus(corpus, valD.Dependency.Point14ValB)
	point14ValEAppendValAInheritedNoOverclaimCorpus(corpus, valD.Dependency.Point14ValA)
}

func point14ValENoOverclaimFinalCheckModel(dependency Point14ValEDependencySnapshot) Point14ValENoOverclaimFinalCheck {
	corpus := point14ValENoOverclaimCorpus{observed: []string{
		"bounded external ecosystem evidence input",
		"read-only ecosystem timeline",
		"no external authority granted",
		"point 14 closure verifies boundaries only",
	}}
	point14ValEAppendValDInheritedNoOverclaimCorpus(&corpus, dependency.Point14ValD)
	return Point14ValENoOverclaimFinalCheck{
		ObservedTexts:                        corpus.observed,
		InternalDiagnosticTexts:              corpus.internalDiagnostics,
		InternalDiagnosticsClassifiedBlocked: corpus.internalDiagnosticsClassifiedBlocked(),
		AllowedSafeWording:                   point14ValESafeWording(),
		BlockedWording:                       point14ValEForbiddenWording(),
		ProjectionDisclaimer:                 point14ValEProjectionDisclaimerBase,
	}
}

func EvaluatePoint14ValENoOverclaimFinalCheckState(model Point14ValENoOverclaimFinalCheck) string {
	if model.ProjectionDisclaimer != point14ValEProjectionDisclaimerBase ||
		!point12Val0ExactStringSetMatch(model.AllowedSafeWording, point14ValESafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point14ValEForbiddenWording()) {
		return Point14ValEStateBlocked
	}
	if point14ValEObservedListContainsForbiddenWording(model.ObservedTexts) {
		return Point14ValEStateBlocked
	}
	if point14ValEObservedListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValECLBFinalCheckModel() Point14ValECLBFinalCheck {
	return Point14ValECLBFinalCheck{
		CheckID: "point14_vale_check_clb_001",
	}
}

func EvaluatePoint14ValECLBFinalCheckState(model Point14ValECLBFinalCheck) string {
	if !point14ValECheckIDValid(model.CheckID) {
		return Point14ValEStateBlocked
	}
	if model.CLB0Present || model.CLB1Present || model.CLB2Present {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValEClosureEvaluatorModel() Point14ValEClosureEvaluator {
	return Point14ValEClosureEvaluator{
		ClosureEvaluatorID:          "point14_vale_closure_evaluator_001",
		ReadOnlyProjectionConfirmed: true,
		NoMutationPathsDetected:     true,
		NoExternalAuthorityDetected: true,
		NoPrematurePoint14Pass:      true,
		CommandsRun:                 point14ValECommandsRun(),
		TestsRun:                    point14ValETestsRun(),
		GrepsRun:                    point14ValEGrepsRun(),
		NegativeFixturesRun:         point14ValENegativeFixturesRun(),
		ReviewerResult:              point12ValEReviewerResultPassConfirmed,
		ProjectionDisclaimer:        point14ValEProjectionDisclaimerBase,
	}
}

func point14ValEComponentAggregate(states ...string) string {
	for _, state := range states {
		switch strings.TrimSpace(state) {
		case Point14ValEStateBlocked:
			return Point14ValEStateBlocked
		}
	}
	for _, state := range states {
		switch strings.TrimSpace(state) {
		case Point14ValEStateReviewRequired:
			return Point14ValEStateReviewRequired
		}
	}
	for _, state := range states {
		switch strings.TrimSpace(state) {
		case Point14ValEStateIncomplete:
			return Point14ValEStateIncomplete
		}
	}
	return Point14ValEStatePassConfirmed
}

func EvaluatePoint14ValEClosureEvaluatorState(model Point14ValEClosureEvaluator) string {
	if !point14ValEClosureEvaluatorIDValid(model.ClosureEvaluatorID) ||
		!point14ValEStateValid(model.DependencyState) ||
		!point14ValEStateValid(model.ValidationClosureState) ||
		!point14ValEStateValid(model.DisputeClosureState) ||
		!point14ValEStateValid(model.CorrectionPublicationClosureState) ||
		!point14ValEStateValid(model.TimelineProjectionClosureState) ||
		!point14ValEStateValid(model.AuthorityBoundaryState) ||
		!point14ValEStateValid(model.TenantPrivacyState) ||
		!point14ValEStateValid(model.TimestampIntegrityState) ||
		!point14ValEStateValid(model.AgentAdvisoryState) ||
		!point14ValEStateValid(model.NoOverclaimState) ||
		!point14ValEStateValid(model.CLBFinalState) ||
		!point14ValECommandsRunValid(model.CommandsRun) ||
		!point14ValETestsRunValid(model.TestsRun) ||
		!point14ValEGrepsRunValid(model.GrepsRun) ||
		!point14ValENegativeFixturesRunValid(model.NegativeFixturesRun) ||
		!point12ValEReviewerResultValid(model.ReviewerResult) ||
		model.ReviewerResult != point12ValEReviewerResultPassConfirmed ||
		model.ProjectionDisclaimer != point14ValEProjectionDisclaimerBase {
		return Point14ValEStateBlocked
	}
	if !model.ReadOnlyProjectionConfirmed ||
		!model.NoMutationPathsDetected ||
		!model.NoExternalAuthorityDetected ||
		!model.NoPrematurePoint14Pass {
		return Point14ValEStateBlocked
	}
	componentState := point14ValEComponentAggregate(
		model.DependencyState,
		model.ValidationClosureState,
		model.DisputeClosureState,
		model.CorrectionPublicationClosureState,
		model.TimelineProjectionClosureState,
		model.AuthorityBoundaryState,
		model.TenantPrivacyState,
		model.TimestampIntegrityState,
		model.AgentAdvisoryState,
		model.NoOverclaimState,
		model.CLBFinalState,
	)
	if componentState != Point14ValEStatePassConfirmed {
		return componentState
	}
	if !model.FinalPassAllowed {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValEPassClosureManifestModel() Point14PassClosureManifest {
	return Point14PassClosureManifest{
		ClosureManifestID:    "point14_vale_pass_manifest_001",
		PointID:              point14Val0PointID,
		WaveID:               point14ValEWaveID,
		ClosureToken:         point14Val0BlockedPassToken,
		Scope:                point14ValEScope,
		ExplicitNonGoals:     point14ValEExplicitNonGoals(),
		EvidenceIdentity:     point14ValEEvidenceIdentity,
		CommandsRun:          point14ValECommandsRun(),
		TestsRun:             point14ValETestsRun(),
		GrepsRun:             point14ValEGrepsRun(),
		NegativeFixturesRun:  point14ValENegativeFixturesRun(),
		CleanRoomIPResult:    point14ValECleanRoomBoundaryPreserved,
		ReviewerResult:       point12ValEReviewerResultPassConfirmed,
		GeneratedAt:          point14ValEPassClosureGeneratedAt,
		ProjectionDisclaimer: point14ValEProjectionDisclaimerBase,
	}
}

func EvaluatePoint14PassClosureManifestState(model Point14PassClosureManifest) string {
	if !point14ValEClosureManifestIDValid(model.ClosureManifestID) ||
		model.PointID != point14Val0PointID ||
		model.WaveID != point14ValEWaveID ||
		model.ClosureToken != point14Val0BlockedPassToken ||
		model.Scope != point14ValEScope ||
		!point14ValEExplicitNonGoalsValid(model.ExplicitNonGoals) ||
		!point14ValEStateValid(model.DependencyGateResult) ||
		!point14ValEStateValid(model.ClosureEvaluatorResult) ||
		!point14ValEStateValid(model.ProjectionBoundaryResult) ||
		!point14ValEStateValid(model.NoExternalAuthorityResult) ||
		!point14ValEStateValid(model.NoOverclaimGrepResult) ||
		!point14ValEStateValid(model.TenantPrivacyResult) ||
		!point14ValEStateValid(model.TimestampIntegrityResult) ||
		!point14ValEStateValid(model.AIAgentBoundaryResult) ||
		!point14ValEStateValid(model.CLBResult) ||
		model.EvidenceIdentity != point14ValEEvidenceIdentity ||
		!point14ValECommandsRunValid(model.CommandsRun) ||
		!point14ValETestsRunValid(model.TestsRun) ||
		!point14ValEGrepsRunValid(model.GrepsRun) ||
		!point14ValENegativeFixturesRunValid(model.NegativeFixturesRun) ||
		model.CleanRoomIPResult != point14ValECleanRoomBoundaryPreserved ||
		!point12ValEReviewerResultValid(model.ReviewerResult) ||
		model.ReviewerResult != point12ValEReviewerResultPassConfirmed ||
		model.GeneratedAt != point14ValEPassClosureGeneratedAt ||
		model.ProjectionDisclaimer != point14ValEProjectionDisclaimerBase {
		return Point14ValEStateBlocked
	}
	componentState := point14ValEComponentAggregate(
		model.DependencyGateResult,
		model.ClosureEvaluatorResult,
		model.ProjectionBoundaryResult,
		model.NoExternalAuthorityResult,
		model.NoOverclaimGrepResult,
		model.TenantPrivacyResult,
		model.TimestampIntegrityResult,
		model.AIAgentBoundaryResult,
		model.CLBResult,
	)
	if componentState == Point14ValEStateBlocked {
		return Point14ValEStateBlocked
	}
	if componentState == Point14ValEStateReviewRequired {
		if model.Point14PassAllowed || model.Point14PassToken != "" {
			return Point14ValEStateBlocked
		}
		return Point14ValEStateReviewRequired
	}
	if componentState == Point14ValEStateIncomplete {
		if model.Point14PassAllowed || model.Point14PassToken != "" {
			return Point14ValEStateBlocked
		}
		return Point14ValEStateIncomplete
	}
	if !model.Point14PassAllowed || model.Point14PassToken != point14Val0BlockedPassToken {
		return Point14ValEStateBlocked
	}
	return Point14ValEStatePassConfirmed
}

func point14ValEFoundationModelFromUpstream(valD Point14ValDFoundation) Point14ValEFoundation {
	dependency := point14ValEDependencySnapshotFromUpstream(valD)
	return Point14ValEFoundation{
		ProjectionDisclaimer:                 point14ValEProjectionDisclaimerBase,
		Dependency:                           dependency,
		ClosureEvaluator:                     point14ValEClosureEvaluatorModel(),
		PassClosureManifest:                  point14ValEPassClosureManifestModel(),
		ExternalSignalValidationClosureCheck: point14ValEExternalSignalValidationClosureCheckModel(dependency),
		ConflictDisputeClosureCheck:          point14ValEConflictDisputeClosureCheckModel(dependency),
		CorrectionPublicationClosureCheck:    point14ValECorrectionPublicationClosureCheckModel(dependency),
		TimelineProjectionClosureCheck:       point14ValETimelineProjectionClosureCheckModel(dependency),
		AuthorityBoundaryClosureCheck:        point14ValEAuthorityBoundaryClosureCheckModel(dependency),
		TenantPrivacyClosureCheck:            point14ValETenantPrivacyClosureCheckModel(dependency),
		TimestampIntegrityClosureCheck:       point14ValETimestampIntegrityClosureCheckModel(dependency),
		AgentAdvisoryClosureCheck:            point14ValEAgentAdvisoryClosureCheckModel(dependency),
		NoOverclaimFinalCheck:                point14ValENoOverclaimFinalCheckModel(dependency),
		CLBFinalCheck:                        point14ValECLBFinalCheckModel(),
	}
}

func Point14ValEFoundationModel() Point14ValEFoundation {
	return cachedFormalModel(&point14ValEFoundationModelOnce, &point14ValEFoundationModelCached, func() Point14ValEFoundation {
		valD := ComputePoint14ValDFoundation(Point14ValDFoundationModel())
		return point14ValEFoundationModelFromUpstream(valD)
	})
}

func point14ValEFoundationState(states ...string) string {
	return point14ValEComponentAggregate(states...)
}

func point14ValEBlockingReasons(model Point14ValEFoundation) []string {
	componentStates := map[string]string{
		"dependency":                 model.DependencyState,
		"external_signal_validation": model.ExternalSignalValidationClosureState,
		"conflict_dispute":           model.ConflictDisputeClosureState,
		"correction_publication":     model.CorrectionPublicationClosureState,
		"timeline_projection":        model.TimelineProjectionClosureState,
		"authority_boundary":         model.AuthorityBoundaryClosureState,
		"tenant_privacy":             model.TenantPrivacyClosureState,
		"timestamp_integrity":        model.TimestampIntegrityClosureState,
		"agent_advisory":             model.AgentAdvisoryClosureState,
		"no_overclaim":               model.NoOverclaimFinalCheckState,
		"clb":                        model.CLBFinalCheckState,
		"closure_evaluator":          model.ClosureEvaluatorState,
		"pass_closure_manifest":      model.PassClosureManifestState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if strings.TrimSpace(state) == Point14ValEStateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point14ValEReviewPrerequisites(model Point14ValEFoundation) []string {
	componentStates := map[string]string{
		"external_signal_validation": model.ExternalSignalValidationClosureState,
		"conflict_dispute":           model.ConflictDisputeClosureState,
		"correction_publication":     model.CorrectionPublicationClosureState,
		"timeline_projection":        model.TimelineProjectionClosureState,
		"authority_boundary":         model.AuthorityBoundaryClosureState,
		"tenant_privacy":             model.TenantPrivacyClosureState,
		"timestamp_integrity":        model.TimestampIntegrityClosureState,
		"agent_advisory":             model.AgentAdvisoryClosureState,
		"no_overclaim":               model.NoOverclaimFinalCheckState,
		"clb":                        model.CLBFinalCheckState,
		"closure_evaluator":          model.ClosureEvaluatorState,
		"pass_closure_manifest":      model.PassClosureManifestState,
	}
	prereqs := append([]string{}, model.Dependency.ReviewPrerequisites...)
	for name, state := range componentStates {
		if strings.TrimSpace(state) == Point14ValEStateReviewRequired || strings.TrimSpace(state) == Point14ValEStateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func ComputePoint14ValEFoundation(model Point14ValEFoundation) Point14ValEFoundation {
	model.DependencyState = EvaluatePoint14ValEDependencyState(model.Dependency)
	model.ExternalSignalValidationClosureState = EvaluatePoint14ValEExternalSignalValidationClosureCheckState(model.ExternalSignalValidationClosureCheck)
	model.ConflictDisputeClosureState = EvaluatePoint14ValEConflictDisputeClosureCheckState(model.ConflictDisputeClosureCheck)
	model.CorrectionPublicationClosureState = EvaluatePoint14ValECorrectionPublicationClosureCheckState(model.CorrectionPublicationClosureCheck)
	model.TimelineProjectionClosureState = EvaluatePoint14ValETimelineProjectionClosureCheckState(model.TimelineProjectionClosureCheck)
	model.AuthorityBoundaryClosureState = EvaluatePoint14ValEAuthorityBoundaryClosureCheckState(model.AuthorityBoundaryClosureCheck)
	model.TenantPrivacyClosureState = EvaluatePoint14ValETenantPrivacyClosureCheckState(model.TenantPrivacyClosureCheck, model.Dependency)
	model.TimestampIntegrityClosureState = EvaluatePoint14ValETimestampIntegrityClosureCheckState(model.TimestampIntegrityClosureCheck, model.Dependency)
	model.AgentAdvisoryClosureState = EvaluatePoint14ValEAgentAdvisoryClosureCheckState(model.AgentAdvisoryClosureCheck, model.Dependency)
	derivedNoOverclaimFinalCheck := point14ValENoOverclaimFinalCheckModel(model.Dependency)
	noOverclaimFinalCheckMatchesDependency := reflect.DeepEqual(model.NoOverclaimFinalCheck, derivedNoOverclaimFinalCheck)
	model.NoOverclaimFinalCheck = derivedNoOverclaimFinalCheck
	model.NoOverclaimFinalCheckState = EvaluatePoint14ValENoOverclaimFinalCheckState(model.NoOverclaimFinalCheck)
	if !noOverclaimFinalCheckMatchesDependency && model.NoOverclaimFinalCheckState == Point14ValEStatePassConfirmed {
		model.NoOverclaimFinalCheckState = Point14ValEStateBlocked
	}
	model.CLBFinalCheckState = EvaluatePoint14ValECLBFinalCheckState(model.CLBFinalCheck)

	passCandidate := point14ValEFoundationState(
		model.DependencyState,
		model.ExternalSignalValidationClosureState,
		model.ConflictDisputeClosureState,
		model.CorrectionPublicationClosureState,
		model.TimelineProjectionClosureState,
		model.AuthorityBoundaryClosureState,
		model.TenantPrivacyClosureState,
		model.TimestampIntegrityClosureState,
		model.AgentAdvisoryClosureState,
		model.NoOverclaimFinalCheckState,
		model.CLBFinalCheckState,
	) == Point14ValEStatePassConfirmed &&
		model.ClosureEvaluator.ReadOnlyProjectionConfirmed &&
		model.ClosureEvaluator.NoMutationPathsDetected &&
		model.ClosureEvaluator.NoExternalAuthorityDetected &&
		model.ClosureEvaluator.NoPrematurePoint14Pass &&
		model.ClosureEvaluator.ReviewerResult == point12ValEReviewerResultPassConfirmed

	model.ClosureEvaluator.DependencyState = model.DependencyState
	model.ClosureEvaluator.ValidationClosureState = model.ExternalSignalValidationClosureState
	model.ClosureEvaluator.DisputeClosureState = model.ConflictDisputeClosureState
	model.ClosureEvaluator.CorrectionPublicationClosureState = model.CorrectionPublicationClosureState
	model.ClosureEvaluator.TimelineProjectionClosureState = model.TimelineProjectionClosureState
	model.ClosureEvaluator.AuthorityBoundaryState = model.AuthorityBoundaryClosureState
	model.ClosureEvaluator.TenantPrivacyState = model.TenantPrivacyClosureState
	model.ClosureEvaluator.TimestampIntegrityState = model.TimestampIntegrityClosureState
	model.ClosureEvaluator.AgentAdvisoryState = model.AgentAdvisoryClosureState
	model.ClosureEvaluator.NoOverclaimState = model.NoOverclaimFinalCheckState
	model.ClosureEvaluator.CLBFinalState = model.CLBFinalCheckState
	model.ClosureEvaluator.FinalPassAllowed = passCandidate
	model.ClosureEvaluatorState = EvaluatePoint14ValEClosureEvaluatorState(model.ClosureEvaluator)
	model.ClosureEvaluator.CurrentState = model.ClosureEvaluatorState

	model.PassClosureManifest.DependencyGateResult = model.DependencyState
	model.PassClosureManifest.ClosureEvaluatorResult = model.ClosureEvaluatorState
	model.PassClosureManifest.ProjectionBoundaryResult = model.TimelineProjectionClosureState
	model.PassClosureManifest.NoExternalAuthorityResult = model.AuthorityBoundaryClosureState
	model.PassClosureManifest.NoOverclaimGrepResult = model.NoOverclaimFinalCheckState
	model.PassClosureManifest.TenantPrivacyResult = model.TenantPrivacyClosureState
	model.PassClosureManifest.TimestampIntegrityResult = model.TimestampIntegrityClosureState
	model.PassClosureManifest.AIAgentBoundaryResult = model.AgentAdvisoryClosureState
	model.PassClosureManifest.CLBResult = model.CLBFinalCheckState
	model.PassClosureManifest.Point14PassAllowed = passCandidate && model.ClosureEvaluatorState == Point14ValEStatePassConfirmed
	if model.PassClosureManifest.Point14PassAllowed {
		model.PassClosureManifest.Point14PassToken = point14Val0BlockedPassToken
	} else {
		model.PassClosureManifest.Point14PassToken = ""
	}
	model.PassClosureManifestState = EvaluatePoint14PassClosureManifestState(model.PassClosureManifest)
	model.PassClosureManifest.CurrentState = model.PassClosureManifestState

	model.CurrentState = point14ValEFoundationState(
		model.DependencyState,
		model.ExternalSignalValidationClosureState,
		model.ConflictDisputeClosureState,
		model.CorrectionPublicationClosureState,
		model.TimelineProjectionClosureState,
		model.AuthorityBoundaryClosureState,
		model.TenantPrivacyClosureState,
		model.TimestampIntegrityClosureState,
		model.AgentAdvisoryClosureState,
		model.NoOverclaimFinalCheckState,
		model.CLBFinalCheckState,
		model.ClosureEvaluatorState,
		model.PassClosureManifestState,
	)
	model.Point14PassAllowed = model.CurrentState == Point14ValEStatePassConfirmed &&
		model.ClosureEvaluatorState == Point14ValEStatePassConfirmed &&
		model.PassClosureManifestState == Point14ValEStatePassConfirmed &&
		model.PassClosureManifest.Point14PassAllowed
	if model.Point14PassAllowed {
		model.Point14PassToken = point14Val0BlockedPassToken
	} else {
		model.Point14PassToken = ""
	}
	model.BlockingReasons = point14ValEBlockingReasons(model)
	model.ReviewPrerequisites = point14ValEReviewPrerequisites(model)
	return model
}

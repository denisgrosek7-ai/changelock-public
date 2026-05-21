package formal

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"
)

const (
	Point15ValEStatePassConfirmed  = "point15_vale_continuous_verification_closure_pass_confirmed"
	Point15ValEStateBlocked        = "point15_vale_continuous_verification_closure_blocked"
	Point15ValEStateReviewRequired = "point15_vale_continuous_verification_closure_review_required"
	Point15ValEStateIncomplete     = "point15_vale_continuous_verification_closure_incomplete"
)

const (
	point15ValEWaveID               = "val_e"
	point15ValEScope                = "final_continuous_verification_closure_gate"
	point15ValEClosureDisclaimer    = "final_continuous_verification_closure_only no_new_feature_semantics point15_vale"
	point15ValECleanRoomIPPreserved = "not_applicable_boundary_preserved"
)

type Point15ValEDependencySnapshot struct {
	Point15ValDCurrentState           string                                   `json:"point15_vald_current_state"`
	Point15ValDDependencyState        string                                   `json:"point15_vald_dependency_state"`
	Point15ValDTimelineState          string                                   `json:"point15_vald_timeline_state"`
	Point15ValDDashboardState         string                                   `json:"point15_vald_dashboard_state"`
	Point15ValDQueryState             string                                   `json:"point15_vald_query_state"`
	Point15ValDEvidenceDetailState    string                                   `json:"point15_vald_evidence_detail_state"`
	Point15ValDRevalidationState      string                                   `json:"point15_vald_revalidation_detail_state"`
	Point15ValDEnforcementState       string                                   `json:"point15_vald_enforcement_detail_state"`
	Point15ValDReplayHistoryState     string                                   `json:"point15_vald_replay_proof_history_state"`
	Point15ValDAccessTenantState      string                                   `json:"point15_vald_access_tenant_state"`
	Point15ValDTimestampDisplayState  string                                   `json:"point15_vald_timestamp_display_state"`
	Point15ValDNoMutationState        string                                   `json:"point15_vald_no_mutation_state"`
	Point15ValDAuthorityBoundaryState string                                   `json:"point15_vald_authority_boundary_state"`
	Point15ValDNoOverclaimState       string                                   `json:"point15_vald_no_overclaim_state"`
	Point15ValDComputedFromUpstream   bool                                     `json:"point15_vald_computed_from_upstream"`
	Point15ValDMerged                 bool                                     `json:"point15_vald_merged"`
	Point15ValDCIGreen                bool                                     `json:"point15_vald_ci_green"`
	Point15ValDReviewedOnMain         bool                                     `json:"point15_vald_reviewed_on_main"`
	Point15PassSeen                   bool                                     `json:"point15_pass_seen"`
	InheritedPoint15ValCCurrentState  string                                   `json:"inherited_point15_valc_current_state"`
	InheritedPoint15ValBCurrentState  string                                   `json:"inherited_point15_valb_current_state"`
	InheritedPoint15ValACurrentState  string                                   `json:"inherited_point15_vala_current_state"`
	InheritedPoint15Val0CurrentState  string                                   `json:"inherited_point15_val0_current_state"`
	InheritedPoint14ValECurrentState  string                                   `json:"inherited_point14_vale_current_state"`
	InheritedTenantScope              string                                   `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput        bool                                     `json:"snapshot_from_computed_output"`
	ReviewPrerequisites               []string                                 `json:"review_prerequisites,omitempty"`
	Point15ValD                       Point15ValDAssuranceProjectionFoundation `json:"point15_vald"`
}

type Point15ValEClosureEvaluator struct {
	ClosureEvaluatorID          string   `json:"closure_evaluator_id"`
	DependencyState             string   `json:"dependency_state"`
	FreshnessTaxonomyState      string   `json:"freshness_taxonomy_state"`
	DowngradeTriggerState       string   `json:"downgrade_trigger_state"`
	ScheduledRevalidationState  string   `json:"scheduled_revalidation_state"`
	EnforcementBoundaryState    string   `json:"enforcement_boundary_state"`
	ProjectionBoundaryState     string   `json:"projection_boundary_state"`
	ReplayProofHistoryState     string   `json:"replay_proof_history_state"`
	TenantPrivacyState          string   `json:"tenant_privacy_state"`
	TimestampIntegrityState     string   `json:"timestamp_integrity_state"`
	AuthorityBoundaryState      string   `json:"authority_boundary_state"`
	NoMutationState             string   `json:"no_mutation_state"`
	NoOverclaimState            string   `json:"no_overclaim_state"`
	CLBFinalState               string   `json:"clb_final_state"`
	ReadOnlyProjectionConfirmed bool     `json:"read_only_projection_confirmed"`
	NoMutationPathsDetected     bool     `json:"no_mutation_paths_detected"`
	NoExternalAuthorityDetected bool     `json:"no_external_authority_detected"`
	ReplayableManifestReady     bool     `json:"replayable_manifest_ready"`
	NoPrematurePoint15Pass      bool     `json:"no_premature_point15_pass"`
	FinalPassAllowed            bool     `json:"final_pass_allowed"`
	CommandsRun                 []string `json:"commands_run,omitempty"`
	TestsRun                    []string `json:"tests_run,omitempty"`
	GrepsRun                    []string `json:"greps_run,omitempty"`
	NegativeFixturesRun         []string `json:"negative_fixtures_run,omitempty"`
	ReviewerResult              string   `json:"reviewer_result"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
}

type Point15PassClosureManifest struct {
	CurrentState                string   `json:"current_state"`
	ClosureManifestID           string   `json:"closure_manifest_id"`
	PointID                     string   `json:"point_id"`
	WaveID                      string   `json:"wave_id"`
	ClosureToken                string   `json:"closure_token"`
	Scope                       string   `json:"scope"`
	ExplicitNonGoals            []string `json:"explicit_non_goals,omitempty"`
	DependencyGateResult        string   `json:"dependency_gate_result"`
	FreshnessTaxonomyResult     string   `json:"freshness_taxonomy_result"`
	DowngradeTriggerResult      string   `json:"downgrade_trigger_result"`
	ScheduledRevalidationResult string   `json:"scheduled_revalidation_result"`
	EnforcementBoundaryResult   string   `json:"enforcement_boundary_result"`
	ProjectionBoundaryResult    string   `json:"projection_boundary_result"`
	ReplayProofHistoryResult    string   `json:"replay_proof_history_result"`
	TenantPrivacyResult         string   `json:"tenant_privacy_result"`
	EvidenceID                  string   `json:"evidence_id"`
	EvidenceIdentity            string   `json:"evidence_identity"`
	EvidenceHash                string   `json:"evidence_hash"`
	PolicyVersion               string   `json:"policy_version"`
	EngineVersion               string   `json:"engine_version"`
	SchemaVersion               string   `json:"schema_version"`
	TenantScope                 string   `json:"tenant_scope"`
	TimestampIntegrityResult    string   `json:"timestamp_integrity_result"`
	AuthorityBoundaryResult     string   `json:"authority_boundary_result"`
	NoMutationResult            string   `json:"no_mutation_result"`
	NoOverclaimResult           string   `json:"no_overclaim_result"`
	CleanRoomIPResult           string   `json:"clean_room_ip_result"`
	CommandsRun                 []string `json:"commands_run,omitempty"`
	TestsRun                    []string `json:"tests_run,omitempty"`
	GrepsRun                    []string `json:"greps_run,omitempty"`
	NegativeFixturesRun         []string `json:"negative_fixtures_run,omitempty"`
	CLBResult                   string   `json:"clb_result"`
	ReviewerResult              string   `json:"reviewer_result"`
	GeneratedAt                 string   `json:"generated_at"`
	Point15PassAllowed          bool     `json:"point15_pass_allowed"`
	Point15PassToken            string   `json:"point15_pass_token"`
	ProjectionDisclaimer        string   `json:"projection_disclaimer"`
}

type Point15ValEFreshnessTaxonomyClosureCheck struct {
	CheckID                       string `json:"check_id"`
	FreshnessTaxonomyState        string `json:"freshness_taxonomy_state"`
	DowngradeTaxonomyState        string `json:"downgrade_taxonomy_state"`
	EvidenceContextState          string `json:"evidence_context_state"`
	FreshnessStatus               string `json:"freshness_status"`
	DowngradeOutcome              string `json:"downgrade_outcome"`
	MappedState                   string `json:"mapped_state"`
	SupersessionLineageRef        string `json:"supersession_lineage_ref"`
	DriftIsDecisive               bool   `json:"drift_is_decisive"`
	MissingFreshnessProofDecisive bool   `json:"missing_freshness_proof_decisive"`
	FreshnessProofPresent         bool   `json:"freshness_proof_present"`
	RetainsPass                   bool   `json:"retains_pass"`
	RetainsActiveClosure          bool   `json:"retains_active_closure"`
	EvidenceID                    string `json:"evidence_id"`
	EvidenceHash                  string `json:"evidence_hash"`
	PolicyVersion                 string `json:"policy_version"`
	EngineVersion                 string `json:"engine_version"`
	SchemaVersion                 string `json:"schema_version"`
	TenantScope                   string `json:"tenant_scope"`
}

type Point15ValEDowngradeTriggerClosureCheck struct {
	CheckID                string `json:"check_id"`
	TriggerTableState      string `json:"trigger_table_state"`
	TriggerState           string `json:"trigger_state"`
	ReasonState            string `json:"reason_state"`
	DecisionState          string `json:"decision_state"`
	TriggerDetected        bool   `json:"trigger_detected"`
	TriggerType            string `json:"trigger_type"`
	TriggerIsDecisive      bool   `json:"trigger_is_decisive"`
	SupersessionLineageRef string `json:"supersession_lineage_ref"`
	TargetState            string `json:"target_state"`
	TargetDowngradeOutcome string `json:"target_downgrade_outcome"`
	RetainsPass            bool   `json:"retains_pass"`
	RetainsActiveClosure   bool   `json:"retains_active_closure"`
}

type Point15ValEScheduledRevalidationClosureCheck struct {
	CheckID                   string `json:"check_id"`
	ScheduleState             string `json:"schedule_state"`
	RunState                  string `json:"run_state"`
	RetryBudgetState          string `json:"retry_budget_state"`
	TenantThrottleState       string `json:"tenant_throttle_state"`
	DowngradeBindingState     string `json:"downgrade_binding_state"`
	TimestampDisciplineState  string `json:"timestamp_discipline_state"`
	AuthorityBoundaryState    string `json:"authority_boundary_state"`
	ScheduledStatus           string `json:"scheduled_status"`
	RunResult                 string `json:"run_result"`
	RetryBudgetStatus         string `json:"retry_budget_status"`
	ThrottleStatus            string `json:"throttle_status"`
	TriggerType               string `json:"trigger_type"`
	TargetState               string `json:"target_state"`
	RunEvidenceHashMatches    bool   `json:"run_evidence_hash_matches"`
	ExactBindingConfirmed     bool   `json:"exact_binding_confirmed"`
	SchedulerAuthorityGranted bool   `json:"scheduler_authority_granted"`
	RetainsPass               bool   `json:"retains_pass"`
	RetainsActiveClosure      bool   `json:"retains_active_closure"`
}

type Point15ValEEnforcementClosureCheck struct {
	CheckID                      string `json:"check_id"`
	EnforcementActionState       string `json:"enforcement_action_state"`
	EvidenceLifecycleState       string `json:"evidence_lifecycle_state"`
	RevocationState              string `json:"revocation_state"`
	ExpiryState                  string `json:"expiry_state"`
	SupersessionState            string `json:"supersession_state"`
	ReplayProofHistoryState      string `json:"replay_proof_history_state"`
	AuthorityBoundaryState       string `json:"authority_boundary_state"`
	EnforcementAction            string `json:"enforcement_action"`
	EnforcementReason            string `json:"enforcement_reason"`
	TargetState                  string `json:"target_state"`
	LifecycleStatus              string `json:"lifecycle_status"`
	HistoryPreserved             bool   `json:"history_preserved"`
	CanonicalMutationAttempted   bool   `json:"canonical_mutation_attempted"`
	ProductionMutationAllowed    bool   `json:"production_mutation_allowed"`
	EvidenceDeletionDetected     bool   `json:"evidence_deletion_detected"`
	SilentReplacementDetected    bool   `json:"silent_replacement_detected"`
	AutomaticPublicationDetected bool   `json:"automatic_publication_detected"`
	RevocationExecutionDetected  bool   `json:"revocation_execution_detected"`
}

type Point15ValEProjectionClosureCheck struct {
	CheckID                 string `json:"check_id"`
	TimelineState           string `json:"timeline_state"`
	DashboardState          string `json:"dashboard_state"`
	QueryState              string `json:"query_state"`
	EvidenceDetailState     string `json:"evidence_detail_state"`
	RevalidationDetailState string `json:"revalidation_detail_state"`
	EnforcementDetailState  string `json:"enforcement_detail_state"`
	ReplayProofHistoryState string `json:"replay_proof_history_state"`
	AccessTenantState       string `json:"access_tenant_state"`
	TimestampDisplayState   string `json:"timestamp_display_state"`
	NoMutationState         string `json:"no_mutation_state"`
	AuthorityBoundaryState  string `json:"authority_boundary_state"`
	NoOverclaimState        string `json:"no_overclaim_state"`
	DisplayOnly             bool   `json:"display_only"`
	MutatesState            bool   `json:"mutates_state"`
	ApprovesPass            bool   `json:"approves_pass"`
	PerformsEnforcement     bool   `json:"performs_enforcement"`
	Publishes               bool   `json:"publishes"`
	Revokes                 bool   `json:"revokes"`
	RestoresActive          bool   `json:"restores_active"`
	HidesDecisiveEvidence   bool   `json:"hides_decisive_evidence"`
	StrengthensClaims       bool   `json:"strengthens_claims"`
}

type Point15ValEReplayProofHistoryClosureCheck struct {
	CheckID                 string `json:"check_id"`
	PriorStateVisible       bool   `json:"prior_state_visible"`
	CurrentStateVisible     bool   `json:"current_state_visible"`
	BlockedReasonVisible    bool   `json:"blocked_reason_visible"`
	DecisiveEvidenceVisible bool   `json:"decisive_evidence_visible"`
	HashBindingVisible      bool   `json:"hash_binding_visible"`
	ReplayRef               string `json:"replay_ref"`
	ProofPackRef            string `json:"proof_pack_ref"`
	ProofHistoryRef         string `json:"proof_history_ref"`
	ProofHistoryHidden      bool   `json:"proof_history_hidden"`
}

type Point15ValETenantPrivacyClosureCheck struct {
	CheckID                        string `json:"check_id"`
	TenantScope                    string `json:"tenant_scope"`
	CrossTenantProofDetected       bool   `json:"cross_tenant_proof_detected"`
	CrossTenantScheduleRunDetected bool   `json:"cross_tenant_schedule_run_detected"`
	CrossTenantEnforcementDetected bool   `json:"cross_tenant_enforcement_detected"`
	CrossTenantProjectionDetected  bool   `json:"cross_tenant_projection_detected"`
	TenantPrivateDataExposed       bool   `json:"tenant_private_data_exposed"`
	UnsafeRedactionStateDetected   bool   `json:"unsafe_redaction_state_detected"`
	RedactionHidesDecisiveEvidence bool   `json:"redaction_hides_decisive_evidence"`
}

type Point15ValETimestampIntegrityClosureCheck struct {
	CheckID                         string `json:"check_id"`
	TenantScope                     string `json:"tenant_scope"`
	ScheduledStatus                 string `json:"scheduled_status"`
	RunResult                       string `json:"run_result"`
	RevalidationRequired            bool   `json:"revalidation_required"`
	FreshnessEvaluatedAt            string `json:"freshness_evaluated_at"`
	FreshnessEvaluatedTimeSource    string `json:"freshness_evaluated_time_source"`
	FreshnessValidatedAt            string `json:"freshness_validated_at"`
	FreshnessValidatedTimeSource    string `json:"freshness_validated_time_source"`
	RevalidationDueAt               string `json:"revalidation_due_at"`
	RevalidationDueTimeSource       string `json:"revalidation_due_time_source"`
	RevalidationCompletedAt         string `json:"revalidation_completed_at"`
	RevalidationCompletedTimeSource string `json:"revalidation_completed_time_source"`
	EnforcementEnforcedAt           string `json:"enforcement_enforced_at"`
	EnforcementEnforcedTimeSource   string `json:"enforcement_enforced_time_source"`
	ProjectionDisplayedAt           string `json:"projection_displayed_at"`
	ProjectionDisplayedTimeSource   string `json:"projection_displayed_time_source"`
	ReferenceNow                    string `json:"reference_now"`
	ReferenceNowTimeSource          string `json:"reference_now_time_source"`
	SourceEventAt                   string `json:"source_event_at"`
	SourceEventTimeSource           string `json:"source_event_time_source"`
	ClientLocalCreatesCanonical     bool   `json:"client_local_creates_canonical"`
	SourceEventCreatesCanonical     bool   `json:"source_event_creates_canonical"`
}

type Point15ValEAuthorityBoundaryClosureCheck struct {
	CheckID                  string `json:"check_id"`
	TenantScope              string `json:"tenant_scope"`
	FormalCoreOnly           bool   `json:"formal_core_only"`
	SchedulerPassAllowed     bool   `json:"scheduler_pass_allowed"`
	ConnectorPassAllowed     bool   `json:"connector_pass_allowed"`
	DashboardPassAllowed     bool   `json:"dashboard_pass_allowed"`
	TimelineAuthorityAllowed bool   `json:"timeline_authority_allowed"`
	QueryMutationAllowed     bool   `json:"query_mutation_allowed"`
	PortalMutationAllowed    bool   `json:"portal_mutation_allowed"`
	CustomerMutationAllowed  bool   `json:"customer_mutation_allowed"`
	AuditorMutationAllowed   bool   `json:"auditor_mutation_allowed"`
	AgentPassAllowed         bool   `json:"agent_pass_allowed"`
	AIPassAllowed            bool   `json:"ai_pass_allowed"`
	ExternalAuthorityAllowed bool   `json:"external_authority_allowed"`
	PublicBadgeAllowed       bool   `json:"public_badge_allowed"`
}

type Point15ValENoMutationClosureCheck struct {
	CheckID                       string `json:"check_id"`
	CanonicalMutationDetected     bool   `json:"canonical_mutation_detected"`
	ProductionMutationDetected    bool   `json:"production_mutation_detected"`
	EvidenceDeletionDetected      bool   `json:"evidence_deletion_detected"`
	HistoryHidingDetected         bool   `json:"history_hiding_detected"`
	RevocationExecutionDetected   bool   `json:"revocation_execution_detected"`
	AutomaticPublicationDetected  bool   `json:"automatic_publication_detected"`
	SilentSupersessionReplacement bool   `json:"silent_supersession_replacement"`
	RetryBudgetResetByNonCore     bool   `json:"retry_budget_reset_by_non_core"`
	PassRestorationDetected       bool   `json:"pass_restoration_detected"`
}

type Point15ValENoOverclaimFinalCheck struct {
	ObservedTexts                        []string `json:"observed_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

type Point15ValECLBFinalCheck struct {
	CheckID        string   `json:"check_id"`
	CLB0Present    bool     `json:"clb0_present"`
	CLB1Present    bool     `json:"clb1_present"`
	CLB2Present    bool     `json:"clb2_present"`
	CLB3Advisories []string `json:"clb3_advisories,omitempty"`
}

type Point15ValEContinuousVerificationClosureFoundation struct {
	CurrentState                      string                                       `json:"current_state"`
	BlockingReasons                   []string                                     `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites               []string                                     `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer              string                                       `json:"projection_disclaimer"`
	DependencyState                   string                                       `json:"dependency_state"`
	ClosureEvaluatorState             string                                       `json:"closure_evaluator_state"`
	PassClosureManifestState          string                                       `json:"pass_closure_manifest_state"`
	FreshnessTaxonomyClosureState     string                                       `json:"freshness_taxonomy_closure_state"`
	DowngradeTriggerClosureState      string                                       `json:"downgrade_trigger_closure_state"`
	ScheduledRevalidationClosureState string                                       `json:"scheduled_revalidation_closure_state"`
	EnforcementClosureState           string                                       `json:"enforcement_closure_state"`
	ProjectionClosureState            string                                       `json:"projection_closure_state"`
	ReplayProofHistoryClosureState    string                                       `json:"replay_proof_history_closure_state"`
	TenantPrivacyClosureState         string                                       `json:"tenant_privacy_closure_state"`
	TimestampIntegrityClosureState    string                                       `json:"timestamp_integrity_closure_state"`
	AuthorityBoundaryClosureState     string                                       `json:"authority_boundary_closure_state"`
	NoMutationClosureState            string                                       `json:"no_mutation_closure_state"`
	NoOverclaimFinalCheckState        string                                       `json:"no_overclaim_final_check_state"`
	CLBFinalCheckState                string                                       `json:"clb_final_check_state"`
	Dependency                        Point15ValEDependencySnapshot                `json:"dependency"`
	ClosureEvaluator                  Point15ValEClosureEvaluator                  `json:"closure_evaluator"`
	PassClosureManifest               Point15PassClosureManifest                   `json:"pass_closure_manifest"`
	FreshnessTaxonomyClosureCheck     Point15ValEFreshnessTaxonomyClosureCheck     `json:"freshness_taxonomy_closure_check"`
	DowngradeTriggerClosureCheck      Point15ValEDowngradeTriggerClosureCheck      `json:"downgrade_trigger_closure_check"`
	ScheduledRevalidationClosureCheck Point15ValEScheduledRevalidationClosureCheck `json:"scheduled_revalidation_closure_check"`
	EnforcementClosureCheck           Point15ValEEnforcementClosureCheck           `json:"enforcement_closure_check"`
	ProjectionClosureCheck            Point15ValEProjectionClosureCheck            `json:"projection_closure_check"`
	ReplayProofHistoryClosureCheck    Point15ValEReplayProofHistoryClosureCheck    `json:"replay_proof_history_closure_check"`
	TenantPrivacyClosureCheck         Point15ValETenantPrivacyClosureCheck         `json:"tenant_privacy_closure_check"`
	TimestampIntegrityClosureCheck    Point15ValETimestampIntegrityClosureCheck    `json:"timestamp_integrity_closure_check"`
	AuthorityBoundaryClosureCheck     Point15ValEAuthorityBoundaryClosureCheck     `json:"authority_boundary_closure_check"`
	NoMutationClosureCheck            Point15ValENoMutationClosureCheck            `json:"no_mutation_closure_check"`
	NoOverclaimFinalCheck             Point15ValENoOverclaimFinalCheck             `json:"no_overclaim_final_check"`
	CLBFinalCheck                     Point15ValECLBFinalCheck                     `json:"clb_final_check"`
}

func point15ValEStates() []string {
	return []string{
		Point15ValEStatePassConfirmed,
		Point15ValEStateBlocked,
		Point15ValEStateReviewRequired,
		Point15ValEStateIncomplete,
	}
}

func point15ValEStateValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValEStates())
}

func formalRawExactNonEmpty(value string) bool {
	if value == "" || value != strings.TrimSpace(value) || strings.ContainsAny(value, "\t\r\n") {
		return false
	}
	for _, r := range value {
		if unicode.IsControl(r) || unicode.Is(unicode.Cf, r) || (unicode.IsSpace(r) && r != ' ') {
			return false
		}
	}
	return true
}

func formalRawExactValid(value string, valid func(string) bool) bool {
	return formalRawExactNonEmpty(value) && valid(value)
}

func point15ValERefValid(value string) bool {
	return point14Val0RefValid(value,
		"point15_vale_",
		"point15_vald_",
		"point15_valc_",
		"point15_valb_",
		"point15_vala_",
		"point15_val0_",
		"timeline_",
		"dashboard_",
		"query_",
		"replay_",
		"proof_",
		"proof_pack_",
		"history_",
		"schedule_",
		"run_",
		"budget_",
		"throttle_",
		"binding_",
		"enforcement_",
		"decision_",
		"authority_boundary_",
		"timestamp_",
		"closure_",
		"manifest_",
	)
}

func point15ValECheckIDValid(value string) bool {
	return formalRawExactTokenValid(value, func(candidate string) bool {
		return point14Val0RefValid(candidate, "point15_vale_check_", "closure_check_")
	})
}

func point15ValEClosureEvaluatorIDValid(value string) bool {
	return formalRawExactTokenValid(value, func(candidate string) bool {
		return point14Val0RefValid(candidate, "point15_vale_closure_evaluator_", "closure_evaluator_")
	})
}

func point15ValEClosureManifestIDValid(value string) bool {
	return point14Val0RefValid(value, "point15_vale_pass_manifest_", "pass_manifest_")
}

func point15ValEForbiddenWording() []string {
	return append([]string{
		"continuous assurance guaranteed",
		"automatically verified forever",
		"always fresh",
		"certified secure",
		"guaranteed secure",
		"compliance guaranteed",
		"regulator-approved",
		"regulator approved",
		"production approved",
		"deployment approved",
		"legal proof",
		"financial guarantee",
		"official authority",
		"global truth",
		"public badge",
	}, inheritedDeploymentReadinessOverclaimClaims()...)
}

func point15ValESafeWording() []string {
	return []string{
		"final continuous verification closure gate",
		"bounded final closure manifest",
		"projection remains read-only",
		"tenant-safe replayable closure evidence",
		"limitations remain visible",
		"point 15 closure verifies boundaries only",
	}
}

func point15ValEExplicitNonGoals() []string {
	return []string{
		"no point 16 implementation",
		"no workspace consolidation",
		"no demo pilot enterprise structure",
		"no new feature beyond final closure gate",
		"no mutation write api",
		"no actual dashboard ui",
		"no portal customer auditor product implementation",
		"no runtime scheduler daemon",
		"no background job execution",
		"no connector write mutation",
		"no revocation execution side effect",
		"no automatic publication",
		"no evidence deletion hiding",
		"no ai agent authority",
		"no external api default",
		"no canonical mutation",
		"no production mutation",
		"no production approval",
		"no certification compliance legal financial guarantee",
		"no public badge global truth official authority wording",
	}
}

func point15ValEObservedTextContainsForbiddenWording(text string) bool {
	return point15ValEObservedTextContainsForbiddenWordingWithNormalizer(text, point15ValENormalizedObservedText)
}

func point15ValEObservedTextContainsForbiddenWordingWithNormalizer(text string, normalize func(string) string) bool {
	normalized := normalize(text)
	if normalized == "" {
		return false
	}
	for _, safe := range point15ValESafeWording() {
		if normalized == normalize(safe) {
			return false
		}
	}
	for _, forbidden := range point15ValEForbiddenWording() {
		if formalNoOverclaimContainsForbidden(normalized, normalize(forbidden)) {
			return true
		}
	}
	return false
}

func point15ValENormalizedObservedText(text string) string {
	return formalNoOverclaimNormalizePublicText(text)
}

func point15ValENormalizedInternalText(text string) string {
	return formalNoOverclaimNormalizeText(text)
}

func point15ValEObservedTextAllowedSafe(normalized string) bool {
	return point15ValEObservedTextAllowedSafeWithNormalizer(normalized, point15ValENormalizedObservedText)
}

func point15ValEObservedTextAllowedSafeWithNormalizer(normalized string, normalize func(string) string) bool {
	for _, safe := range point15ValESafeWording() {
		if normalized == normalize(safe) {
			return true
		}
	}
	return false
}

func point15ValEObservedListContainsForbiddenWording(values []string) bool {
	return point15ValEListContainsForbiddenWording(values, point15ValENormalizedObservedText)
}

func point15ValEInternalListContainsForbiddenWording(values []string) bool {
	return point15ValEListContainsForbiddenWording(values, point15ValENormalizedInternalText)
}

func point15ValEListContainsForbiddenWording(values []string, normalize func(string) string) bool {
	return point15ObservedListContainsForbiddenWordingWithNormalizer(values, point15ValESafeWording(), point15ValEForbiddenWording(), normalize)
}

func point15ValECommandsRun() []string {
	return []string{
		"git diff --check",
		"gofmt on changed Go files",
		"go test ./internal/formal -run 'Test.*Point15ValE.*|Test.*Point15.*ValE.*' -v",
		"go test ./internal/formal -run 'Test.*Point15ValD.*|Test.*Point15.*ValD.*' -v",
		"go test ./internal/formal -run 'Test.*Point15ValC.*|Test.*Point15.*ValC.*' -v",
		"go test ./internal/formal -run 'Test.*Point15ValB.*|Test.*Point15.*ValB.*' -v",
		"go test ./internal/formal -run 'Test.*Point15ValA.*|Test.*Point15.*ValA.*' -v",
		"go test ./internal/formal -run 'Test.*Point15Val0.*|Test.*Point15.*Val0.*' -v",
		"go test ./internal/formal -run 'TestPoint10ThroughPoint15ValDCurrentSweep|TestPoint10ThroughPoint15ValCCurrentSweep|TestPoint10ThroughPoint15ValBCurrentSweep|TestPoint10ThroughPoint15ValACurrentSweep|TestPoint10ThroughPoint15Val0CurrentSweep|TestPoint10ThroughPoint14CurrentSweep' -v",
		"go test ./internal/formal -run 'TestPoint15Val0CachedHelperIsolation|TestPoint15ValACachedHelperIsolation|TestPoint15ValBCachedHelperIsolation|TestPoint15ValCCachedHelperIsolation' -v",
		"go test ./internal/formal -run 'Test.*Point14ValE.*|Test.*Point14.*ValE.*' -v",
		"go test ./internal/formal -run 'Test.*Point14.*' -v",
		"go test ./internal/formal -run 'Test.*Point13.*' -v",
		"go test ./internal/formal -run 'Test.*Point12.*|Test.*Replay.*|Test.*ProofPack.*|Test.*Binding.*|Test.*Mutation.*' -v",
		"go test ./internal/formal -run 'Test.*Point11.*|Test.*Claim.*|Test.*NoOverclaim.*|Test.*Governance.*' -v",
		"go test ./internal/formal -run 'Test.*AI.*|Test.*Agent.*|Test.*Lineage.*|Test.*Provenance.*' -v",
		"go test -timeout 20m ./...",
	}
}

func point15ValETestsRun() []string {
	return []string{
		"point15_vale_dependency",
		"point15_vale_pass_closure_manifest",
		"point15_vale_freshness_taxonomy_closure",
		"point15_vale_downgrade_trigger_closure",
		"point15_vale_scheduled_revalidation_closure",
		"point15_vale_enforcement_closure",
		"point15_vale_projection_closure",
		"point15_vale_replay_proof_history_closure",
		"point15_vale_tenant_privacy_closure",
		"point15_vale_timestamp_integrity_closure",
		"point15_vale_authority_boundary_closure",
		"point15_vale_no_mutation_closure",
		"point15_vale_no_overclaim_final",
		"point15_vale_clb_final",
		"point10_through_point15_vald_current_sweep",
		"point15_helper_cache_isolation",
	}
}

func point15ValEGrepsRun() []string {
	return []string{
		"point_15_pass scan",
		"forbidden wording scan",
		"ai authority scan",
		"authority mutation marker scan",
		"external api scan",
		"skip todo fixme scan",
	}
}

func point15ValENegativeFixturesRun() []string {
	return []string{
		"premature_point15_pass",
		"freshness_negative",
		"downgrade_negative",
		"revalidation_negative",
		"projection_negative",
		"timestamp_negative",
		"authority_negative",
		"clb_negative",
	}
}

func point15ValECommandsRunValid(values []string) bool {
	return point12Val0ExactStringSetMatch(values, point15ValECommandsRun())
}

func point15ValETestsRunValid(values []string) bool {
	return point12Val0ExactStringSetMatch(values, point15ValETestsRun())
}

func point15ValEGrepsRunValid(values []string) bool {
	return point12Val0ExactStringSetMatch(values, point15ValEGrepsRun())
}

func point15ValENegativeFixturesRunValid(values []string) bool {
	return point12Val0ExactStringSetMatch(values, point15ValENegativeFixturesRun())
}

func point15ValEExplicitNonGoalsValid(values []string) bool {
	return point12Val0ExactStringSetMatch(values, point15ValEExplicitNonGoals())
}

func point15ValETargetStateToWaveState(state string) string {
	if !formalRawExactValid(state, point15Val0StateValid) {
		return Point15ValEStateBlocked
	}
	switch state {
	case Point15Val0StateBlocked:
		return Point15ValEStateBlocked
	case Point15Val0StateReviewRequired:
		return Point15ValEStateReviewRequired
	case Point15Val0StateIncomplete:
		return Point15ValEStateIncomplete
	default:
		return Point15ValEStatePassConfirmed
	}
}

func point15ValEComponentAggregate(states ...string) string {
	if len(states) == 0 {
		return Point15ValEStateBlocked
	}
	for _, state := range states {
		if !formalRawExactValid(state, point15ValEStateValid) || state == Point15ValEStateBlocked {
			return Point15ValEStateBlocked
		}
	}
	for _, state := range states {
		if state == Point15ValEStateReviewRequired {
			return Point15ValEStateReviewRequired
		}
	}
	for _, state := range states {
		if state == Point15ValEStateIncomplete {
			return Point15ValEStateIncomplete
		}
	}
	return Point15ValEStatePassConfirmed
}

func point15ValEValDPayloadContainsPoint15Pass(valD Point15ValDAssuranceProjectionFoundation) bool {
	payload, err := json.Marshal(valD)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point15Val0BlockedPassToken)
}

func point15ValEDependencySnapshotFromUpstream(valD Point15ValDAssuranceProjectionFoundation) Point15ValEDependencySnapshot {
	return Point15ValEDependencySnapshot{
		Point15ValDCurrentState:           valD.CurrentState,
		Point15ValDDependencyState:        valD.DependencyState,
		Point15ValDTimelineState:          valD.TimelineState,
		Point15ValDDashboardState:         valD.DashboardState,
		Point15ValDQueryState:             valD.QueryState,
		Point15ValDEvidenceDetailState:    valD.EvidenceDetailState,
		Point15ValDRevalidationState:      valD.RevalidationDetailState,
		Point15ValDEnforcementState:       valD.EnforcementDetailState,
		Point15ValDReplayHistoryState:     valD.ReplayProofHistoryState,
		Point15ValDAccessTenantState:      valD.AccessTenantState,
		Point15ValDTimestampDisplayState:  valD.TimestampDisplayState,
		Point15ValDNoMutationState:        valD.NoMutationState,
		Point15ValDAuthorityBoundaryState: valD.AuthorityBoundaryState,
		Point15ValDNoOverclaimState:       valD.NoOverclaimState,
		Point15ValDComputedFromUpstream:   valD.Dependency.SnapshotFromComputedOutput,
		Point15ValDMerged:                 true,
		Point15ValDCIGreen:                true,
		Point15ValDReviewedOnMain:         true,
		Point15PassSeen:                   point15ValEValDPayloadContainsPoint15Pass(valD),
		InheritedPoint15ValCCurrentState:  valD.Dependency.Point15ValCCurrentState,
		InheritedPoint15ValBCurrentState:  valD.Dependency.InheritedPoint15ValBCurrentState,
		InheritedPoint15ValACurrentState:  valD.Dependency.InheritedPoint15ValACurrentState,
		InheritedPoint15Val0CurrentState:  valD.Dependency.InheritedPoint15Val0CurrentState,
		InheritedPoint14ValECurrentState:  valD.Dependency.InheritedPoint14ValECurrentState,
		InheritedTenantScope:              valD.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:        true,
		ReviewPrerequisites:               append([]string{}, valD.ReviewPrerequisites...),
		Point15ValD:                       valD,
	}
}

func point15ValEDependencySnapshotModel() Point15ValEDependencySnapshot {
	valD := ComputePoint15ValDAssuranceProjectionFoundation(Point15ValDFoundationModel())
	return point15ValEDependencySnapshotFromUpstream(valD)
}

func EvaluatePoint15ValEDependencyState(model Point15ValEDependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point15ValDComputedFromUpstream ||
		!model.Point15ValDMerged ||
		!model.Point15ValDCIGreen ||
		!model.Point15ValDReviewedOnMain ||
		model.Point15PassSeen ||
		point15ValEValDPayloadContainsPoint15Pass(model.Point15ValD) ||
		!formalRawExactValid(model.Point15ValDCurrentState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDDependencyState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDTimelineState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDDashboardState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDQueryState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDEvidenceDetailState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDRevalidationState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDEnforcementState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDReplayHistoryState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDAccessTenantState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDTimestampDisplayState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDNoMutationState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDAuthorityBoundaryState, point15ValDStateValid) ||
		!formalRawExactValid(model.Point15ValDNoOverclaimState, point15ValDStateValid) ||
		!formalRawExactValid(model.InheritedPoint15ValCCurrentState, point15ValCStateValid) ||
		!formalRawExactValid(model.InheritedPoint15ValBCurrentState, point15ValBStateValid) ||
		!formalRawExactValid(model.InheritedPoint15ValACurrentState, point15ValAStateValid) ||
		!formalRawExactValid(model.InheritedPoint15Val0CurrentState, point15Val0StateValid) ||
		!formalRawExactValid(model.InheritedPoint14ValECurrentState, point14ValEStateValid) ||
		!formalRawExactValid(model.InheritedTenantScope, point11Val0ScopeValid) {
		return Point15ValEStateBlocked
	}
	if model.Point15ValDCurrentState != model.Point15ValD.CurrentState ||
		model.Point15ValDDependencyState != model.Point15ValD.DependencyState ||
		model.Point15ValDTimelineState != model.Point15ValD.TimelineState ||
		model.Point15ValDDashboardState != model.Point15ValD.DashboardState ||
		model.Point15ValDQueryState != model.Point15ValD.QueryState ||
		model.Point15ValDEvidenceDetailState != model.Point15ValD.EvidenceDetailState ||
		model.Point15ValDRevalidationState != model.Point15ValD.RevalidationDetailState ||
		model.Point15ValDEnforcementState != model.Point15ValD.EnforcementDetailState ||
		model.Point15ValDReplayHistoryState != model.Point15ValD.ReplayProofHistoryState ||
		model.Point15ValDAccessTenantState != model.Point15ValD.AccessTenantState ||
		model.Point15ValDTimestampDisplayState != model.Point15ValD.TimestampDisplayState ||
		model.Point15ValDNoMutationState != model.Point15ValD.NoMutationState ||
		model.Point15ValDAuthorityBoundaryState != model.Point15ValD.AuthorityBoundaryState ||
		model.Point15ValDNoOverclaimState != model.Point15ValD.NoOverclaimState ||
		model.Point15ValDComputedFromUpstream != model.Point15ValD.Dependency.SnapshotFromComputedOutput ||
		model.InheritedPoint15ValCCurrentState != model.Point15ValD.Dependency.Point15ValCCurrentState ||
		model.InheritedPoint15ValBCurrentState != model.Point15ValD.Dependency.InheritedPoint15ValBCurrentState ||
		model.InheritedPoint15ValACurrentState != model.Point15ValD.Dependency.InheritedPoint15ValACurrentState ||
		model.InheritedPoint15Val0CurrentState != model.Point15ValD.Dependency.InheritedPoint15Val0CurrentState ||
		model.InheritedPoint14ValECurrentState != model.Point15ValD.Dependency.InheritedPoint14ValECurrentState ||
		model.InheritedTenantScope != model.Point15ValD.Dependency.InheritedTenantScope {
		return Point15ValEStateBlocked
	}
	recomputedValD := ComputePoint15ValDAssuranceProjectionFoundation(model.Point15ValD)
	if recomputedValD.CurrentState != model.Point15ValD.CurrentState ||
		recomputedValD.DependencyState != model.Point15ValD.DependencyState ||
		recomputedValD.TimelineState != model.Point15ValD.TimelineState ||
		recomputedValD.DashboardState != model.Point15ValD.DashboardState ||
		recomputedValD.QueryState != model.Point15ValD.QueryState ||
		recomputedValD.EvidenceDetailState != model.Point15ValD.EvidenceDetailState ||
		recomputedValD.RevalidationDetailState != model.Point15ValD.RevalidationDetailState ||
		recomputedValD.EnforcementDetailState != model.Point15ValD.EnforcementDetailState ||
		recomputedValD.ReplayProofHistoryState != model.Point15ValD.ReplayProofHistoryState ||
		recomputedValD.AccessTenantState != model.Point15ValD.AccessTenantState ||
		recomputedValD.TimestampDisplayState != model.Point15ValD.TimestampDisplayState ||
		recomputedValD.NoMutationState != model.Point15ValD.NoMutationState ||
		recomputedValD.AuthorityBoundaryState != model.Point15ValD.AuthorityBoundaryState ||
		recomputedValD.NoOverclaimState != model.Point15ValD.NoOverclaimState ||
		recomputedValD.CurrentState != model.Point15ValDCurrentState ||
		recomputedValD.DependencyState != model.Point15ValDDependencyState ||
		recomputedValD.TimelineState != model.Point15ValDTimelineState ||
		recomputedValD.DashboardState != model.Point15ValDDashboardState ||
		recomputedValD.QueryState != model.Point15ValDQueryState ||
		recomputedValD.EvidenceDetailState != model.Point15ValDEvidenceDetailState ||
		recomputedValD.RevalidationDetailState != model.Point15ValDRevalidationState ||
		recomputedValD.EnforcementDetailState != model.Point15ValDEnforcementState ||
		recomputedValD.ReplayProofHistoryState != model.Point15ValDReplayHistoryState ||
		recomputedValD.AccessTenantState != model.Point15ValDAccessTenantState ||
		recomputedValD.TimestampDisplayState != model.Point15ValDTimestampDisplayState ||
		recomputedValD.NoMutationState != model.Point15ValDNoMutationState ||
		recomputedValD.AuthorityBoundaryState != model.Point15ValDAuthorityBoundaryState ||
		recomputedValD.NoOverclaimState != model.Point15ValDNoOverclaimState {
		return Point15ValEStateBlocked
	}
	if EvaluatePoint15ValDNoOverclaimGuardState(model.Point15ValD.NoOverclaimGuard) != Point15ValDStateActive {
		return Point15ValEStateBlocked
	}
	if !point15ValDEmbeddedStateChainActive(model.Point15ValD) {
		return Point15ValEStateBlocked
	}
	if model.Point15ValDCurrentState != Point15ValDStateActive ||
		model.Point15ValDDependencyState != Point15ValDStateActive ||
		model.Point15ValDTimelineState != Point15ValDStateActive ||
		model.Point15ValDDashboardState != Point15ValDStateActive ||
		model.Point15ValDQueryState != Point15ValDStateActive ||
		model.Point15ValDEvidenceDetailState != Point15ValDStateActive ||
		model.Point15ValDRevalidationState != Point15ValDStateActive ||
		model.Point15ValDEnforcementState != Point15ValDStateActive ||
		model.Point15ValDReplayHistoryState != Point15ValDStateActive ||
		model.Point15ValDAccessTenantState != Point15ValDStateActive ||
		model.Point15ValDTimestampDisplayState != Point15ValDStateActive ||
		model.Point15ValDNoMutationState != Point15ValDStateActive ||
		model.Point15ValDAuthorityBoundaryState != Point15ValDStateActive ||
		model.Point15ValDNoOverclaimState != Point15ValDStateActive ||
		model.InheritedPoint15ValCCurrentState != Point15ValCStateActive ||
		model.InheritedPoint15ValBCurrentState != Point15ValBStateActive ||
		model.InheritedPoint15ValACurrentState != Point15ValAStateActive ||
		model.InheritedPoint15Val0CurrentState != Point15Val0StateActive ||
		model.InheritedPoint14ValECurrentState != Point14ValEStatePassConfirmed {
		return Point15ValEStateBlocked
	}
	return Point15ValEStatePassConfirmed
}

func point15ValEFreshnessTaxonomyClosureCheckModel(dependency Point15ValEDependencySnapshot) Point15ValEFreshnessTaxonomyClosureCheck {
	val0 := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0
	return Point15ValEFreshnessTaxonomyClosureCheck{
		CheckID:                       "point15_vale_check_freshness_001",
		FreshnessTaxonomyState:        val0.FreshnessTaxonomyState,
		DowngradeTaxonomyState:        val0.DowngradeTaxonomyState,
		EvidenceContextState:          val0.EvidenceContextState,
		FreshnessStatus:               val0.FreshnessTaxonomy.FreshnessStatus,
		DowngradeOutcome:              val0.DowngradeTaxonomy.DowngradeOutcome,
		MappedState:                   val0.FreshnessTaxonomy.MappedState,
		SupersessionLineageRef:        val0.DowngradeTaxonomy.SupersessionLineageRef,
		DriftIsDecisive:               val0.DowngradeTaxonomy.DriftIsDecisive,
		MissingFreshnessProofDecisive: val0.DowngradeTaxonomy.MissingFreshnessProofDecisive,
		FreshnessProofPresent:         val0.DowngradeTaxonomy.FreshnessProofPresent,
		RetainsPass:                   val0.DowngradeTaxonomy.RetainsPass,
		RetainsActiveClosure:          val0.DowngradeTaxonomy.RetainsActiveClosure,
		EvidenceID:                    val0.EvidenceContext.EvidenceID,
		EvidenceHash:                  val0.EvidenceContext.EvidenceHash,
		PolicyVersion:                 val0.EvidenceContext.PolicyVersion,
		EngineVersion:                 val0.EvidenceContext.EngineVersion,
		SchemaVersion:                 val0.EvidenceContext.SchemaVersion,
		TenantScope:                   val0.EvidenceContext.TenantScope,
	}
}

func EvaluatePoint15ValEFreshnessTaxonomyClosureCheckState(model Point15ValEFreshnessTaxonomyClosureCheck) string {
	if !point15ValECheckIDValid(model.CheckID) ||
		model.FreshnessTaxonomyState != Point15Val0StateActive ||
		model.DowngradeTaxonomyState != Point15Val0StateActive ||
		model.EvidenceContextState != Point15Val0StateActive ||
		!formalRawExactValid(model.FreshnessStatus, point15Val0FreshnessStatusValid) ||
		!formalRawExactValid(model.DowngradeOutcome, point15Val0DowngradeOutcomeValid) ||
		!formalRawExactValid(model.MappedState, point15Val0StateValid) ||
		!formalRawExactNonEmpty(model.EvidenceID) ||
		!formalRawExactNonEmpty(model.EvidenceHash) ||
		!formalRawExactNonEmpty(model.PolicyVersion) ||
		!formalRawExactNonEmpty(model.EngineVersion) ||
		!formalRawExactNonEmpty(model.SchemaVersion) ||
		!formalRawExactValid(model.TenantScope, point11Val0ScopeValid) {
		return Point15ValEStateBlocked
	}
	downgrade := Point15Val0DowngradeTaxonomy{
		FreshnessStatus:               model.FreshnessStatus,
		SupersessionLineageRef:        model.SupersessionLineageRef,
		DriftIsDecisive:               model.DriftIsDecisive,
		MissingFreshnessProofDecisive: model.MissingFreshnessProofDecisive,
		FreshnessProofPresent:         model.FreshnessProofPresent,
	}
	expectedOutcome := point15Val0ExpectedDowngradeOutcome(downgrade)
	expectedState := point15Val0ExpectedDowngradeState(downgrade)
	if model.DowngradeOutcome != expectedOutcome || model.MappedState != expectedState {
		return Point15ValEStateBlocked
	}
	if model.FreshnessStatus != point15Val0FreshnessFresh && (model.RetainsPass || model.RetainsActiveClosure) {
		return Point15ValEStateBlocked
	}
	switch model.FreshnessStatus {
	case point15Val0FreshnessFresh:
		if model.DowngradeOutcome != point15Val0DowngradeRetainActive || model.MappedState != Point15Val0StateActive {
			return Point15ValEStateBlocked
		}
		return Point15ValEStatePassConfirmed
	case point15Val0FreshnessStale:
		return Point15ValEStateReviewRequired
	case point15Val0FreshnessExpired, point15Val0FreshnessRevoked, point15Val0FreshnessUnsupported, point15Val0FreshnessTampered:
		return Point15ValEStateBlocked
	case point15Val0FreshnessSuperseded:
		if !formalRawExactNonEmpty(model.SupersessionLineageRef) {
			return Point15ValEStateBlocked
		}
		return Point15ValEStateReviewRequired
	case point15Val0FreshnessDrifted:
		if model.DriftIsDecisive {
			return Point15ValEStateBlocked
		}
		return Point15ValEStateReviewRequired
	case point15Val0FreshnessMissing:
		if model.MissingFreshnessProofDecisive || !model.FreshnessProofPresent {
			return Point15ValEStateBlocked
		}
		return Point15ValEStateIncomplete
	default:
		return Point15ValEStateBlocked
	}
}

func point15ValEDowngradeTriggerClosureCheckModel(dependency Point15ValEDependencySnapshot) Point15ValEDowngradeTriggerClosureCheck {
	valA := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA
	return Point15ValEDowngradeTriggerClosureCheck{
		CheckID:                "point15_vale_check_downgrade_001",
		TriggerTableState:      valA.TriggerTableState,
		TriggerState:           valA.TriggerState,
		ReasonState:            valA.ReasonState,
		DecisionState:          valA.DecisionState,
		TriggerDetected:        valA.Trigger.TriggerDetected,
		TriggerType:            valA.Trigger.TriggerType,
		TriggerIsDecisive:      valA.Trigger.TriggerIsDecisive,
		SupersessionLineageRef: valA.Trigger.SupersessionLineageRef,
		TargetState:            valA.Decision.TargetState,
		TargetDowngradeOutcome: valA.Decision.TargetDowngradeOutcome,
		RetainsPass:            valA.Decision.RetainsPass,
		RetainsActiveClosure:   valA.Decision.RetainsActiveClosure,
	}
}

func EvaluatePoint15ValEDowngradeTriggerClosureCheckState(model Point15ValEDowngradeTriggerClosureCheck) string {
	if !point15ValECheckIDValid(model.CheckID) ||
		model.TriggerTableState != Point15ValAStateActive ||
		model.TriggerState != Point15ValAStateActive ||
		model.ReasonState != Point15ValAStateActive ||
		model.DecisionState != Point15ValAStateActive {
		return Point15ValEStateBlocked
	}
	if !model.TriggerDetected {
		if model.TriggerType != "" || model.TargetState != Point15Val0StateActive || model.TargetDowngradeOutcome != point15Val0DowngradeRetainActive || model.RetainsPass || !model.RetainsActiveClosure {
			return Point15ValEStateBlocked
		}
		return Point15ValEStatePassConfirmed
	}
	if !formalRawExactValid(model.TriggerType, point15ValATriggerValid) || !formalRawExactValid(model.TargetState, point15Val0StateValid) || !formalRawExactValid(model.TargetDowngradeOutcome, point15Val0DowngradeOutcomeValid) || model.RetainsPass || model.RetainsActiveClosure {
		return Point15ValEStateBlocked
	}
	expectedState := point15ValATriggerExpectedState(model.TriggerType, model.TriggerIsDecisive, model.SupersessionLineageRef != "")
	expectedOutcome := point15ValATriggerExpectedOutcome(model.TriggerType, model.TriggerIsDecisive, model.SupersessionLineageRef != "")
	if model.TargetState != expectedState || model.TargetDowngradeOutcome != expectedOutcome {
		return Point15ValEStateBlocked
	}
	return point15ValETargetStateToWaveState(model.TargetState)
}

func point15ValEScheduledRevalidationClosureCheckModel(dependency Point15ValEDependencySnapshot) Point15ValEScheduledRevalidationClosureCheck {
	valB := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB
	exactBindingConfirmed := valB.DowngradeBinding.ScheduleRef == valB.Schedule.ScheduleID &&
		valB.DowngradeBinding.RetryBudgetRef == valB.RetryBudget.BudgetID &&
		valB.DowngradeBinding.ThrottleRef == valB.TenantThrottle.ThrottleID &&
		valB.DowngradeBinding.ScheduleStatus == valB.Schedule.ScheduledStatus &&
		valB.DowngradeBinding.RunResult == valB.Run.RunResult &&
		valB.DowngradeBinding.RetryBudgetStatus == valB.RetryBudget.RetryBudgetStatus &&
		valB.DowngradeBinding.ThrottleStatus == valB.TenantThrottle.ThrottleStatus
	if valB.Run.RunResult == point15ValBRunNotRun {
		exactBindingConfirmed = exactBindingConfirmed && valB.DowngradeBinding.RunRef == ""
	} else {
		exactBindingConfirmed = exactBindingConfirmed && valB.DowngradeBinding.RunRef == valB.Run.RunID
	}
	return Point15ValEScheduledRevalidationClosureCheck{
		CheckID:                   "point15_vale_check_revalidation_001",
		ScheduleState:             valB.ScheduleState,
		RunState:                  valB.RunState,
		RetryBudgetState:          valB.RetryBudgetState,
		TenantThrottleState:       valB.TenantThrottleState,
		DowngradeBindingState:     valB.DowngradeBindingState,
		TimestampDisciplineState:  valB.TimestampDisciplineState,
		AuthorityBoundaryState:    valB.AuthorityBoundaryState,
		ScheduledStatus:           valB.Schedule.ScheduledStatus,
		RunResult:                 valB.Run.RunResult,
		RetryBudgetStatus:         valB.RetryBudget.RetryBudgetStatus,
		ThrottleStatus:            valB.TenantThrottle.ThrottleStatus,
		TriggerType:               valB.DowngradeBinding.TriggerType,
		TargetState:               valB.DowngradeBinding.TargetState,
		RunEvidenceHashMatches:    valB.DowngradeBinding.RunEvidenceHashMatches,
		ExactBindingConfirmed:     exactBindingConfirmed,
		SchedulerAuthorityGranted: valB.AuthorityBoundary.SchedulerMarksEvidenceFresh || valB.AuthorityBoundary.SchedulerCreatesRevalidationTruth || valB.AuthorityBoundary.PassAllowed,
		RetainsPass:               valB.DowngradeBinding.RetainsPass,
		RetainsActiveClosure:      valB.DowngradeBinding.RetainsActiveClosure,
	}
}

func EvaluatePoint15ValEScheduledRevalidationClosureCheckState(model Point15ValEScheduledRevalidationClosureCheck) string {
	if !point15ValECheckIDValid(model.CheckID) ||
		model.ScheduleState != Point15ValBStateActive ||
		model.RunState != Point15ValBStateActive ||
		model.RetryBudgetState != Point15ValBStateActive ||
		model.TenantThrottleState != Point15ValBStateActive ||
		model.DowngradeBindingState != Point15ValBStateActive ||
		model.TimestampDisciplineState != Point15ValBStateActive ||
		model.AuthorityBoundaryState != Point15ValBStateActive ||
		!formalRawExactValid(model.ScheduledStatus, point15ValBScheduleStatusValid) ||
		!formalRawExactValid(model.RunResult, point15ValBRunResultValid) ||
		!formalRawExactValid(model.RetryBudgetStatus, point15ValBRetryStatusValid) ||
		!formalRawExactValid(model.ThrottleStatus, point15ValBThrottleStatusValid) ||
		!formalRawExactValid(model.TargetState, point15Val0StateValid) {
		return Point15ValEStateBlocked
	}
	if !model.ExactBindingConfirmed || model.SchedulerAuthorityGranted || model.RetainsPass {
		return Point15ValEStateBlocked
	}
	if model.ScheduledStatus != point15ValBScheduleCompleted &&
		model.ScheduledStatus != point15ValBScheduleScheduled &&
		model.ScheduledStatus != point15ValBScheduleNotRequired &&
		model.RetainsActiveClosure {
		return Point15ValEStateBlocked
	}
	if model.RunResult == point15ValBRunCompletedClean || model.RunResult == point15ValBRunCompletedWithDowngrade {
		if !model.RunEvidenceHashMatches {
			return Point15ValEStateBlocked
		}
	}
	if model.TargetState != Point15Val0StateActive {
		return point15ValETargetStateToWaveState(model.TargetState)
	}
	switch model.ScheduledStatus {
	case point15ValBScheduleMissed, point15ValBScheduleOverdue, point15ValBScheduleFailed, point15ValBScheduleRetryExhausted, point15ValBScheduleBlocked:
		return Point15ValEStateBlocked
	case point15ValBScheduleDue, point15ValBScheduleRetryPending, point15ValBScheduleThrottled, point15ValBScheduleRunning:
		return Point15ValEStateReviewRequired
	}
	switch model.RunResult {
	case point15ValBRunFailed, point15ValBRunUnauthorized, point15ValBRunTenantMismatch, point15ValBRunTampered:
		return Point15ValEStateBlocked
	case point15ValBRunMissed, point15ValBRunTimeout, point15ValBRunThrottled:
		return Point15ValEStateReviewRequired
	}
	switch model.RetryBudgetStatus {
	case point15ValBRetryExhausted, point15ValBRetryBlocked:
		return Point15ValEStateBlocked
	}
	switch model.ThrottleStatus {
	case point15ValBThrottleBlocked, point15ValBThrottleCrossTenantBlocked:
		return Point15ValEStateBlocked
	case point15ValBThrottleReviewRequired:
		return Point15ValEStateReviewRequired
	}
	return Point15ValEStatePassConfirmed
}

func point15ValEEnforcementClosureCheckModel(dependency Point15ValEDependencySnapshot) Point15ValEEnforcementClosureCheck {
	valC := dependency.Point15ValD.Dependency.Point15ValC
	historyPreserved := valC.EvidenceLifecycle.HistoryPreserved && valC.RevocationBoundary.HistoryPreserved && valC.ExpiryBoundary.ExpiryHistoryPreserved && valC.SupersessionBoundary.HistoryPreserved
	return Point15ValEEnforcementClosureCheck{
		CheckID:                      "point15_vale_check_enforcement_001",
		EnforcementActionState:       valC.EnforcementActionState,
		EvidenceLifecycleState:       valC.EvidenceLifecycleState,
		RevocationState:              valC.RevocationBoundaryState,
		ExpiryState:                  valC.ExpiryBoundaryState,
		SupersessionState:            valC.SupersessionState,
		ReplayProofHistoryState:      valC.ReplayProofHistoryState,
		AuthorityBoundaryState:       valC.AuthorityBoundaryState,
		EnforcementAction:            valC.EnforcementAction.EnforcementAction,
		EnforcementReason:            valC.EnforcementAction.EnforcementReason,
		TargetState:                  valC.EnforcementAction.TargetState,
		LifecycleStatus:              valC.EvidenceLifecycle.LifecycleStatus,
		HistoryPreserved:             historyPreserved,
		CanonicalMutationAttempted:   valC.EvidenceLifecycle.CanonicalMutationAttempted,
		ProductionMutationAllowed:    valC.AuthorityBoundary.ProductionMutationAllowed,
		EvidenceDeletionDetected:     !valC.ExpiryBoundary.ExpiryHistoryPreserved || !valC.RevocationBoundary.HistoryPreserved || !valC.SupersessionBoundary.HistoryPreserved,
		SilentReplacementDetected:    valC.SupersessionBoundary.SilentReplacementDetected,
		AutomaticPublicationDetected: valC.RevocationBoundary.AutoPublished || valC.SupersessionBoundary.AutoPublished || valC.AuthorityBoundary.AutomaticPublicationAllowed,
		RevocationExecutionDetected:  valC.RevocationBoundary.AutoRevoked || valC.AuthorityBoundary.RevocationExecutionSideEffectAllowed,
	}
}

func EvaluatePoint15ValEEnforcementClosureCheckState(model Point15ValEEnforcementClosureCheck) string {
	if !point15ValECheckIDValid(model.CheckID) ||
		model.EnforcementActionState != Point15ValCStateActive ||
		model.EvidenceLifecycleState != Point15ValCStateActive ||
		model.RevocationState != Point15ValCStateActive ||
		model.ExpiryState != Point15ValCStateActive ||
		model.SupersessionState != Point15ValCStateActive ||
		model.ReplayProofHistoryState != Point15ValCStateActive ||
		model.AuthorityBoundaryState != Point15ValCStateActive ||
		!formalRawExactValid(model.EnforcementAction, point15ValCActionValid) ||
		!formalRawExactValid(model.LifecycleStatus, point15ValCLifecycleStatusValid) ||
		!formalRawExactValid(model.TargetState, point15Val0StateValid) {
		return Point15ValEStateBlocked
	}
	if model.CanonicalMutationAttempted || model.ProductionMutationAllowed || model.EvidenceDeletionDetected || model.SilentReplacementDetected || model.AutomaticPublicationDetected || model.RevocationExecutionDetected || !model.HistoryPreserved {
		return Point15ValEStateBlocked
	}
	if model.EnforcementAction == point15ValCActionNone {
		return Point15ValEStatePassConfirmed
	}
	if !formalRawExactValid(model.EnforcementReason, point15ValCReasonValid) {
		return Point15ValEStateBlocked
	}
	return point15ValETargetStateToWaveState(model.TargetState)
}

func point15ValEValDProjectionContractDisplayOnly(valD Point15ValDAssuranceProjectionFoundation) bool {
	return valD.Timeline.ProjectionMode == point15ValDModeTimeline &&
		(valD.Timeline.ProjectionAction == point15ValDActionDisplayOnly || valD.Timeline.ProjectionAction == point15ValDActionExplainOnly) &&
		point15ValEVisibilityIsNonPublic(valD.Timeline.Visibility) &&
		valD.Dashboard.ProjectionMode == point15ValDModeDashboardSummary &&
		valD.Dashboard.ProjectionAction == point15ValDActionDisplayOnly &&
		point15ValEVisibilityIsNonPublic(valD.Dashboard.Visibility) &&
		valD.Query.ProjectionMode == point15ValDModeQueryResult &&
		(valD.Query.ProjectionAction == point15ValDActionFilterOnly || valD.Query.ProjectionAction == point15ValDActionSortOnly || valD.Query.ProjectionAction == point15ValDActionExplainOnly) &&
		point15ValEVisibilityIsNonPublic(valD.Query.Visibility) &&
		valD.EvidenceDetail.ProjectionMode == point15ValDModeEvidenceDetail &&
		(valD.EvidenceDetail.ProjectionAction == point15ValDActionDisplayOnly || valD.EvidenceDetail.ProjectionAction == point15ValDActionExplainOnly) &&
		point15ValEVisibilityIsNonPublic(valD.EvidenceDetail.Visibility) &&
		valD.RevalidationDetail.ProjectionMode == point15ValDModeRevalidationDetail &&
		(valD.RevalidationDetail.ProjectionAction == point15ValDActionDisplayOnly || valD.RevalidationDetail.ProjectionAction == point15ValDActionExplainOnly) &&
		point15ValEVisibilityIsNonPublic(valD.RevalidationDetail.Visibility) &&
		valD.EnforcementDetail.ProjectionMode == point15ValDModeEnforcementDetail &&
		(valD.EnforcementDetail.ProjectionAction == point15ValDActionDisplayOnly || valD.EnforcementDetail.ProjectionAction == point15ValDActionExplainOnly) &&
		point15ValEVisibilityIsNonPublic(valD.EnforcementDetail.Visibility) &&
		(valD.ReplayProofHistory.ProjectionMode == point15ValDModeReplayDetail || valD.ReplayProofHistory.ProjectionMode == point15ValDModeExportPreview) &&
		(valD.ReplayProofHistory.ProjectionAction == point15ValDActionDisplayOnly || valD.ReplayProofHistory.ProjectionAction == point15ValDActionExplainOnly || valD.ReplayProofHistory.ProjectionAction == point15ValDActionExportPreviewOnly) &&
		point15ValEVisibilityIsNonPublic(valD.ReplayProofHistory.Visibility) &&
		point15ValEVisibilityIsNonPublic(valD.AccessTenantPrivacy.Visibility) &&
		valD.TimestampDisplayDiscipline.ProjectionMode == point15ValDModeTimeline
}

func point15ValEVisibilityIsNonPublic(value string) bool {
	return point15ValDVisibilityValid(value) && value != point15ValDVisibilityPublicBlocked
}

func point15ValEProjectionClosureCheckModel(dependency Point15ValEDependencySnapshot) Point15ValEProjectionClosureCheck {
	valD := dependency.Point15ValD
	displayOnly := point15ValEValDProjectionContractDisplayOnly(valD)
	return Point15ValEProjectionClosureCheck{
		CheckID:                 "point15_vale_check_projection_001",
		TimelineState:           valD.TimelineState,
		DashboardState:          valD.DashboardState,
		QueryState:              valD.QueryState,
		EvidenceDetailState:     valD.EvidenceDetailState,
		RevalidationDetailState: valD.RevalidationDetailState,
		EnforcementDetailState:  valD.EnforcementDetailState,
		ReplayProofHistoryState: valD.ReplayProofHistoryState,
		AccessTenantState:       valD.AccessTenantState,
		TimestampDisplayState:   valD.TimestampDisplayState,
		NoMutationState:         valD.NoMutationState,
		AuthorityBoundaryState:  valD.AuthorityBoundaryState,
		NoOverclaimState:        valD.NoOverclaimState,
		DisplayOnly:             displayOnly,
		MutatesState:            valD.Query.QueryMutationAttempted || valD.NoMutationGuard.EvidenceMutationAttempted || valD.NoMutationGuard.LifecycleMutationAttempted || valD.NoMutationGuard.EnforcementMutationAttempted || valD.AccessTenantPrivacy.ProjectionStateMutated || valD.RevalidationDetail.ScheduleMutationAttempted || valD.RevalidationDetail.RetryBudgetResetAttempted,
		ApprovesPass:            valD.AuthorityBoundary.PassAllowed || valD.AuthorityBoundary.DashboardApprovesPass,
		PerformsEnforcement:     valD.EnforcementDetail.PerformsEnforcement || valD.AuthorityBoundary.QueryEnforcesState,
		Publishes:               valD.AuthorityBoundary.ExportPreviewPublishes || valD.EnforcementDetail.AutoPublishes,
		Revokes:                 valD.EnforcementDetail.AutoRevokes,
		RestoresActive:          valD.Dashboard.RestoresActiveClosure || valD.RevalidationDetail.RestoresActiveClosure || valD.NoMutationGuard.PassRestoreAttempted,
		HidesDecisiveEvidence:   !valD.Timeline.DecisiveEvidenceVisible || !valD.Query.DecisiveEvidenceVisible || !valD.ReplayProofHistory.DecisiveEvidenceVisible || valD.AccessTenantPrivacy.DecisiveFailureHidden || valD.ReplayProofHistory.ProofHistoryHidden,
		StrengthensClaims:       valD.Query.StrengthensClaims || point15ValDObservedListContainsForbiddenWording(valD.NoOverclaimGuard.ObservedTexts),
	}
}

func EvaluatePoint15ValEProjectionClosureCheckState(model Point15ValEProjectionClosureCheck) string {
	if !point15ValECheckIDValid(model.CheckID) ||
		model.TimelineState != Point15ValDStateActive ||
		model.DashboardState != Point15ValDStateActive ||
		model.QueryState != Point15ValDStateActive ||
		model.EvidenceDetailState != Point15ValDStateActive ||
		model.RevalidationDetailState != Point15ValDStateActive ||
		model.EnforcementDetailState != Point15ValDStateActive ||
		model.ReplayProofHistoryState != Point15ValDStateActive ||
		model.AccessTenantState != Point15ValDStateActive ||
		model.TimestampDisplayState != Point15ValDStateActive ||
		model.NoMutationState != Point15ValDStateActive ||
		model.AuthorityBoundaryState != Point15ValDStateActive ||
		model.NoOverclaimState != Point15ValDStateActive ||
		!model.DisplayOnly {
		return Point15ValEStateBlocked
	}
	if model.MutatesState || model.ApprovesPass || model.PerformsEnforcement || model.Publishes || model.Revokes || model.RestoresActive || model.HidesDecisiveEvidence || model.StrengthensClaims {
		return Point15ValEStateBlocked
	}
	return Point15ValEStatePassConfirmed
}

func point15ValEReplayProofHistoryClosureCheckModel(dependency Point15ValEDependencySnapshot) Point15ValEReplayProofHistoryClosureCheck {
	valD := dependency.Point15ValD.ReplayProofHistory
	return Point15ValEReplayProofHistoryClosureCheck{
		CheckID:                 "point15_vale_check_replay_001",
		PriorStateVisible:       valD.PriorStateVisible,
		CurrentStateVisible:     valD.CurrentStateVisible,
		BlockedReasonVisible:    valD.BlockedReasonVisible,
		DecisiveEvidenceVisible: valD.DecisiveEvidenceVisible,
		HashBindingVisible:      valD.HashBindingVisible,
		ReplayRef:               valD.ReplayRef,
		ProofPackRef:            valD.ProofPackRef,
		ProofHistoryRef:         valD.ProofHistoryRef,
		ProofHistoryHidden:      valD.ProofHistoryHidden,
	}
}

func EvaluatePoint15ValEReplayProofHistoryClosureCheckState(model Point15ValEReplayProofHistoryClosureCheck) string {
	if !point15ValECheckIDValid(model.CheckID) ||
		!point15ValDRefValid(model.ReplayRef) ||
		!point15ValDRefValid(model.ProofPackRef) ||
		!point15ValDRefValid(model.ProofHistoryRef) {
		return Point15ValEStateBlocked
	}
	if !model.PriorStateVisible || !model.CurrentStateVisible || !model.BlockedReasonVisible || !model.DecisiveEvidenceVisible || !model.HashBindingVisible || model.ProofHistoryHidden {
		return Point15ValEStateBlocked
	}
	return Point15ValEStatePassConfirmed
}

func point15ValETenantPrivacyClosureCheckModel(dependency Point15ValEDependencySnapshot) Point15ValETenantPrivacyClosureCheck {
	val0 := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0
	valB := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB
	valC := dependency.Point15ValD.Dependency.Point15ValC
	valD := dependency.Point15ValD
	tenant := dependency.InheritedTenantScope
	queryRedactionStateValid := point15ValDRedactionStateValid(valD.Query.RedactionState)
	return Point15ValETenantPrivacyClosureCheck{
		CheckID:                        "point15_vale_check_tenant_001",
		TenantScope:                    tenant,
		CrossTenantProofDetected:       val0.EvidenceContext.ReferencedTenantScope != "" && val0.EvidenceContext.ReferencedTenantScope != tenant,
		CrossTenantScheduleRunDetected: valB.TenantThrottle.CrossTenantDetected || valB.Schedule.TenantScope != tenant || valB.Run.TenantScope != tenant,
		CrossTenantEnforcementDetected: valC.TenantBoundary.CrossTenantDetected || valC.TenantBoundary.EnforcementResultTenantScope != tenant || valC.TenantBoundary.ReferencedTenantScope != tenant,
		CrossTenantProjectionDetected:  valD.AccessTenantPrivacy.CrossTenantDetected || valD.Query.CrossTenantQuery || valD.AccessTenantPrivacy.ViewerScope != tenant,
		TenantPrivateDataExposed:       valD.AccessTenantPrivacy.TenantPrivateDataExposed,
		UnsafeRedactionStateDetected:   !queryRedactionStateValid,
		RedactionHidesDecisiveEvidence: valD.AccessTenantPrivacy.DecisiveFailureHidden || (queryRedactionStateValid && valD.Query.RedactionState != point15ValDRedactionNone && !valD.Query.DecisiveEvidenceVisible),
	}
}

func EvaluatePoint15ValETenantPrivacyClosureCheckState(model Point15ValETenantPrivacyClosureCheck) string {
	if !point15ValECheckIDValid(model.CheckID) || !formalRawExactValid(model.TenantScope, point11Val0ScopeValid) {
		return Point15ValEStateBlocked
	}
	if model.CrossTenantProofDetected || model.CrossTenantScheduleRunDetected || model.CrossTenantEnforcementDetected || model.CrossTenantProjectionDetected || model.TenantPrivateDataExposed || model.UnsafeRedactionStateDetected {
		return Point15ValEStateBlocked
	}
	if model.RedactionHidesDecisiveEvidence {
		return Point15ValEStateReviewRequired
	}
	return Point15ValEStatePassConfirmed
}

func point15ValETimestampIntegrityClosureCheckModel(dependency Point15ValEDependencySnapshot) Point15ValETimestampIntegrityClosureCheck {
	val0 := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0
	valB := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB
	valC := dependency.Point15ValD.Dependency.Point15ValC
	valD := dependency.Point15ValD
	enforcedAt := valC.TimestampDiscipline.EnforcedAt
	enforcedAtTimeSource := valC.TimestampDiscipline.EnforcedAtTimeSource
	if enforcedAt == "" {
		enforcedAt = valC.TimestampDiscipline.EvaluatedAt
		enforcedAtTimeSource = valC.TimestampDiscipline.EvaluatedAtTimeSource
	}
	return Point15ValETimestampIntegrityClosureCheck{
		CheckID:                         "point15_vale_check_timestamp_001",
		TenantScope:                     dependency.InheritedTenantScope,
		ScheduledStatus:                 valB.Schedule.ScheduledStatus,
		RunResult:                       valB.Run.RunResult,
		RevalidationRequired:            valB.Schedule.Required,
		FreshnessEvaluatedAt:            val0.TimestampDiscipline.EvaluatedAt,
		FreshnessEvaluatedTimeSource:    val0.TimestampDiscipline.EvaluatedTimeSource,
		FreshnessValidatedAt:            val0.TimestampDiscipline.ValidatedAt,
		FreshnessValidatedTimeSource:    val0.TimestampDiscipline.ValidatedTimeSource,
		RevalidationDueAt:               valB.TimestampDiscipline.DueAt,
		RevalidationDueTimeSource:       valB.TimestampDiscipline.DueAtTimeSource,
		RevalidationCompletedAt:         valB.TimestampDiscipline.CompletedAt,
		RevalidationCompletedTimeSource: valB.TimestampDiscipline.CompletedAtTimeSource,
		EnforcementEnforcedAt:           enforcedAt,
		EnforcementEnforcedTimeSource:   enforcedAtTimeSource,
		ProjectionDisplayedAt:           valD.TimestampDisplayDiscipline.DisplayedAt,
		ProjectionDisplayedTimeSource:   valD.TimestampDisplayDiscipline.TimeSource,
		ReferenceNow:                    valD.TimestampDisplayDiscipline.ReferenceNow,
		ReferenceNowTimeSource:          valD.TimestampDisplayDiscipline.TimeSource,
		SourceEventAt:                   valD.TimestampDisplayDiscipline.SourceEventAt,
		SourceEventTimeSource:           valC.TimestampDiscipline.SourceEventTimeSource,
		ClientLocalCreatesCanonical:     val0.TimestampDiscipline.ClientLocalCreatesCanonical || valB.TimestampDiscipline.ClientLocalCreatesCanonical || valC.TimestampDiscipline.ClientLocalCreatesCanonical || valD.TimestampDisplayDiscipline.ClientLocalCreatesCanonical,
		SourceEventCreatesCanonical:     val0.TimestampDiscipline.SourceEventCreatesCanonical || valB.TimestampDiscipline.SourceEventCreatesCanonical || valC.TimestampDiscipline.SourceEventCreatesCanonical || valD.TimestampDisplayDiscipline.SourceEventCreatesCanonical,
	}
}

func point15ValEMissingTimestampRequired(model Point15ValETimestampIntegrityClosureCheck) bool {
	if model.FreshnessEvaluatedAt == "" ||
		model.FreshnessValidatedAt == "" ||
		model.EnforcementEnforcedAt == "" ||
		model.ProjectionDisplayedAt == "" ||
		model.ReferenceNow == "" {
		return true
	}
	if model.RevalidationRequired && model.RevalidationDueAt == "" {
		return true
	}
	if model.ScheduledStatus == point15ValBScheduleCompleted &&
		(model.RunResult == point15ValBRunCompletedClean || model.RunResult == point15ValBRunCompletedWithDowngrade) &&
		model.RevalidationCompletedAt == "" {
		return true
	}
	return false
}

func EvaluatePoint15ValETimestampIntegrityClosureCheckState(model Point15ValETimestampIntegrityClosureCheck) string {
	if !point15ValECheckIDValid(model.CheckID) ||
		!formalRawExactValid(model.TenantScope, point11Val0ScopeValid) ||
		!formalRawExactValid(model.ScheduledStatus, point15ValBScheduleStatusValid) ||
		!formalRawExactValid(model.RunResult, point15ValBRunResultValid) {
		return Point15ValEStateBlocked
	}
	if point15ValEMissingTimestampRequired(model) {
		return Point15ValEStateIncomplete
	}
	if model.ClientLocalCreatesCanonical || model.SourceEventCreatesCanonical {
		return Point15ValEStateBlocked
	}
	requiredPairs := [][2]string{
		{model.FreshnessEvaluatedAt, model.FreshnessEvaluatedTimeSource},
		{model.FreshnessValidatedAt, model.FreshnessValidatedTimeSource},
		{model.EnforcementEnforcedAt, model.EnforcementEnforcedTimeSource},
		{model.ProjectionDisplayedAt, model.ProjectionDisplayedTimeSource},
		{model.ReferenceNow, model.ReferenceNowTimeSource},
	}
	if model.RevalidationRequired {
		requiredPairs = append(requiredPairs, [2]string{model.RevalidationDueAt, model.RevalidationDueTimeSource})
	}
	if model.RevalidationCompletedAt != "" {
		requiredPairs = append(requiredPairs, [2]string{model.RevalidationCompletedAt, model.RevalidationCompletedTimeSource})
	}
	for _, pair := range requiredPairs {
		if !point15ValERawCanonicalTimeValid(pair[0]) || !formalRawExactValid(pair[1], point14Val0CanonicalTimeSourceValid) {
			return Point15ValEStateBlocked
		}
	}
	if model.SourceEventAt != "" {
		if !point15ValERawCanonicalTimeValid(model.SourceEventAt) || !formalRawExactValid(model.SourceEventTimeSource, point14Val0TimeSourceValid) {
			return Point15ValEStateBlocked
		}
	}
	freshEval, _ := point14Val0ParsedTime(model.FreshnessEvaluatedAt)
	freshValidated, _ := point14Val0ParsedTime(model.FreshnessValidatedAt)
	enforcedAt, _ := point14Val0ParsedTime(model.EnforcementEnforcedAt)
	displayedAt, _ := point14Val0ParsedTime(model.ProjectionDisplayedAt)
	referenceNow, _ := point14Val0ParsedTime(model.ReferenceNow)
	if freshValidated.Before(freshEval) || enforcedAt.Before(freshValidated) || displayedAt.Before(enforcedAt) {
		return Point15ValEStateReviewRequired
	}
	if enforcedAt.After(referenceNow) || displayedAt.After(referenceNow) || freshEval.After(referenceNow) {
		return Point15ValEStateReviewRequired
	}
	if model.RevalidationDueAt != "" {
		dueAt, _ := point14Val0ParsedTime(model.RevalidationDueAt)
		if dueAt.After(referenceNow) && (model.ScheduledStatus == point15ValBScheduleDue || model.ScheduledStatus == point15ValBScheduleOverdue) {
			return Point15ValEStateBlocked
		}
	}
	if model.RevalidationCompletedAt != "" {
		completedAt, _ := point14Val0ParsedTime(model.RevalidationCompletedAt)
		if completedAt.Before(freshValidated) {
			return Point15ValEStateReviewRequired
		}
	}
	return Point15ValEStatePassConfirmed
}

func point15ValEAuthorityBoundaryClosureCheckModel(dependency Point15ValEDependencySnapshot) Point15ValEAuthorityBoundaryClosureCheck {
	val0 := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0
	valA := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA
	valB := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB
	valC := dependency.Point15ValD.Dependency.Point15ValC
	valD := dependency.Point15ValD
	formalCoreOnly := val0.AuthorityBoundary.ExternalSourceInputOnly &&
		val0.AuthorityBoundary.AgentRecommendationAdvisoryOnly &&
		valA.AuthorityBoundary.ExternalSourceInputOnly &&
		valA.AuthorityBoundary.FormalEvaluatorOnly &&
		valB.AuthorityBoundary.ExternalSourceInputOnly &&
		valB.AuthorityBoundary.FormalEvaluatorOnly &&
		valB.AuthorityBoundary.AgentRecommendationAdvisoryOnly &&
		valC.AuthorityBoundary.ExternalSourceInputOnly &&
		valC.AuthorityBoundary.FormalEvaluatorOnly &&
		valC.AuthorityBoundary.AgentRecommendationAdvisoryOnly &&
		valD.AuthorityBoundary.FormalCoreOnly
	return Point15ValEAuthorityBoundaryClosureCheck{
		CheckID:                  "point15_vale_check_authority_001",
		TenantScope:              dependency.InheritedTenantScope,
		FormalCoreOnly:           formalCoreOnly,
		SchedulerPassAllowed:     val0.AuthorityBoundary.SchedulerPassAllowed || valB.AuthorityBoundary.SchedulerMarksEvidenceFresh || valB.AuthorityBoundary.SchedulerCreatesRevalidationTruth || valC.AuthorityBoundary.SchedulerEnforcesBoundary,
		ConnectorPassAllowed:     val0.AuthorityBoundary.ConnectorFreshnessAuthorityAllowed || valA.AuthorityBoundary.ConnectorRestoresActiveClosure || valA.AuthorityBoundary.ConnectorMarksEvidenceFresh || valA.AuthorityBoundary.ConnectorOverridesTerminalStatus || valB.AuthorityBoundary.ConnectorRestoresActiveClosure || valC.AuthorityBoundary.ConnectorRestoresActiveClosure || valD.AuthorityBoundary.ConnectorAuthorityGranted,
		DashboardPassAllowed:     val0.AuthorityBoundary.DashboardFreshnessAllowed || valB.AuthorityBoundary.DashboardSuppressesOverdueStatus || valC.AuthorityBoundary.DashboardSuppressesEnforcement || valD.AuthorityBoundary.DashboardApprovesPass,
		TimelineAuthorityAllowed: valD.AuthorityBoundary.TimelineCreatesAuthority,
		QueryMutationAllowed:     valD.AuthorityBoundary.QueryEnforcesState,
		PortalMutationAllowed:    val0.AuthorityBoundary.PortalProjectionMutatesFreshness || valA.AuthorityBoundary.PortalProjectionMutatesDowngrade || valB.AuthorityBoundary.PortalProjectionMutatesRevalidation || valC.AuthorityBoundary.PortalProjectionMutatesEnforcement || valD.AuthorityBoundary.PortalMutationAttempted,
		CustomerMutationAllowed:  val0.AuthorityBoundary.CustomerProjectionMutatesFreshness || valA.AuthorityBoundary.CustomerProjectionMutatesDowngrade || valB.AuthorityBoundary.CustomerProjectionMutatesRevalidation || valC.AuthorityBoundary.CustomerProjectionMutatesEnforcement || valD.AuthorityBoundary.CustomerMutationAttempted,
		AuditorMutationAllowed:   val0.AuthorityBoundary.AuditorProjectionMutatesFreshness || valA.AuthorityBoundary.AuditorProjectionMutatesDowngrade || valB.AuthorityBoundary.AuditorProjectionMutatesRevalidation || valC.AuthorityBoundary.AuditorProjectionMutatesEnforcement || valD.AuthorityBoundary.AuditorMutationAttempted,
		AgentPassAllowed:         val0.AuthorityBoundary.AgentFreshnessAllowed || valB.AuthorityBoundary.AgentSatisfiesRevalidation || valC.AuthorityBoundary.AgentSatisfiesEnforcement || valD.AuthorityBoundary.AgentAuthorityGranted,
		AIPassAllowed:            false,
		ExternalAuthorityAllowed: val0.AuthorityBoundary.FreshnessAuthorityAllowed || valA.AuthorityBoundary.PassAllowed || valB.AuthorityBoundary.PassAllowed || valC.AuthorityBoundary.PassAllowed,
		PublicBadgeAllowed:       false,
	}
}

func EvaluatePoint15ValEAuthorityBoundaryClosureCheckState(model Point15ValEAuthorityBoundaryClosureCheck) string {
	if !point15ValECheckIDValid(model.CheckID) || !formalRawExactValid(model.TenantScope, point11Val0ScopeValid) || !model.FormalCoreOnly {
		return Point15ValEStateBlocked
	}
	if model.SchedulerPassAllowed ||
		model.ConnectorPassAllowed ||
		model.DashboardPassAllowed ||
		model.TimelineAuthorityAllowed ||
		model.QueryMutationAllowed ||
		model.PortalMutationAllowed ||
		model.CustomerMutationAllowed ||
		model.AuditorMutationAllowed ||
		model.AgentPassAllowed ||
		model.AIPassAllowed ||
		model.ExternalAuthorityAllowed ||
		model.PublicBadgeAllowed {
		return Point15ValEStateBlocked
	}
	return Point15ValEStatePassConfirmed
}

func point15ValENoMutationClosureCheckModel(dependency Point15ValEDependencySnapshot) Point15ValENoMutationClosureCheck {
	val0 := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0
	valA := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA
	valB := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB
	valC := dependency.Point15ValD.Dependency.Point15ValC
	valD := dependency.Point15ValD
	return Point15ValENoMutationClosureCheck{
		CheckID:                       "point15_vale_check_mutation_001",
		CanonicalMutationDetected:     val0.AuthorityBoundary.CanonicalMutationAllowed || valA.AuthorityBoundary.CanonicalMutationAllowed || valB.AuthorityBoundary.CanonicalMutationAllowed || valC.AuthorityBoundary.CanonicalMutationAllowed || valD.AuthorityBoundary.CanonicalMutationAllowed || valC.EvidenceLifecycle.CanonicalMutationAttempted,
		ProductionMutationDetected:    val0.AuthorityBoundary.ProductionMutationAllowed || valA.AuthorityBoundary.ProductionMutationAllowed || valB.AuthorityBoundary.ProductionMutationAllowed || valC.AuthorityBoundary.ProductionMutationAllowed || valD.AuthorityBoundary.ProductionMutationAllowed,
		EvidenceDeletionDetected:      !valC.RevocationBoundary.HistoryPreserved || !valC.ExpiryBoundary.ExpiryHistoryPreserved || !valC.SupersessionBoundary.HistoryPreserved || valD.EnforcementDetail.DeletesEvidence || valD.NoMutationGuard.ExpiryDeletionAttempted,
		HistoryHidingDetected:         !valC.ReplayProofHistory.PriorStateVisible || !valC.ReplayProofHistory.CurrentStateVisible || !valC.ReplayProofHistory.DecisiveEvidenceVisible || !valC.ReplayProofHistory.BlockedReasonVisible || valD.ReplayProofHistory.ProofHistoryHidden,
		RevocationExecutionDetected:   valC.RevocationBoundary.AutoRevoked || valC.AuthorityBoundary.RevocationExecutionSideEffectAllowed || valD.NoMutationGuard.RevocationExecutionAttempted,
		AutomaticPublicationDetected:  valC.RevocationBoundary.AutoPublished || valC.SupersessionBoundary.AutoPublished || valC.AuthorityBoundary.AutomaticPublicationAllowed || valD.EnforcementDetail.AutoPublishes,
		SilentSupersessionReplacement: valC.SupersessionBoundary.SilentReplacementDetected || valD.EnforcementDetail.SilentReplacement || valD.NoMutationGuard.SupersessionReplacementAttempted,
		RetryBudgetResetByNonCore:     valB.AuthorityBoundary.RetryBudgetResetAllowed || valD.RevalidationDetail.RetryBudgetResetAttempted || valD.NoMutationGuard.ScheduleRetryMutationAttempted,
		PassRestorationDetected:       val0.AuthorityBoundary.PassAllowed || valA.AuthorityBoundary.PassAllowed || valB.AuthorityBoundary.PassAllowed || valC.AuthorityBoundary.PassAllowed || valD.AuthorityBoundary.PassAllowed || valD.NoMutationGuard.PassRestoreAttempted || valD.Dashboard.RestoresActiveClosure || valD.RevalidationDetail.RestoresActiveClosure,
	}
}

func EvaluatePoint15ValENoMutationClosureCheckState(model Point15ValENoMutationClosureCheck) string {
	if !point15ValECheckIDValid(model.CheckID) {
		return Point15ValEStateBlocked
	}
	if model.CanonicalMutationDetected ||
		model.ProductionMutationDetected ||
		model.EvidenceDeletionDetected ||
		model.HistoryHidingDetected ||
		model.RevocationExecutionDetected ||
		model.AutomaticPublicationDetected ||
		model.SilentSupersessionReplacement ||
		model.RetryBudgetResetByNonCore ||
		model.PassRestorationDetected {
		return Point15ValEStateBlocked
	}
	return Point15ValEStatePassConfirmed
}

func point15ValENoOverclaimFinalCheckModel(dependency Point15ValEDependencySnapshot) Point15ValENoOverclaimFinalCheck {
	val0 := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0
	valA := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA
	valB := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB
	valC := dependency.Point15ValD.Dependency.Point15ValC
	valD := dependency.Point15ValD
	if !point15ValDEmbeddedNoOverclaimChainActive(valD) {
		return Point15ValENoOverclaimFinalCheck{
			ObservedTexts:                        []string{"point15 nested no-overclaim guard blocked"},
			InternalDiagnosticTexts:              []string{"nested_no_overclaim_guard_blocked"},
			InternalDiagnosticsClassifiedBlocked: false,
			AllowedSafeWording:                   point15ValESafeWording(),
			BlockedWording:                       point15ValEForbiddenWording(),
			ProjectionDisclaimer:                 "nested_no_overclaim_guard_blocked",
		}
	}
	observed := []string{
		"final continuous verification closure gate",
		"projection remains read-only",
		"point 15 closure verifies boundaries only",
	}
	observed = append(observed, val0.NoOverclaimGuard.ObservedTexts...)
	observed = append(observed, valA.NoOverclaimGuard.ObservedTexts...)
	observed = append(observed, valB.NoOverclaimGuard.ObservedTexts...)
	observed = append(observed, valC.NoOverclaimGuard.ObservedTexts...)
	observed = append(observed, valD.NoOverclaimGuard.ObservedTexts...)
	internalDiagnostics := append([]string{}, val0.NoOverclaimGuard.InternalDiagnosticTexts...)
	internalDiagnostics = append(internalDiagnostics, valA.NoOverclaimGuard.InternalDiagnosticTexts...)
	internalDiagnostics = append(internalDiagnostics, valB.NoOverclaimGuard.InternalDiagnosticTexts...)
	internalDiagnostics = append(internalDiagnostics, valC.NoOverclaimGuard.InternalDiagnosticTexts...)
	internalDiagnostics = append(internalDiagnostics, valD.NoOverclaimGuard.InternalDiagnosticTexts...)
	return Point15ValENoOverclaimFinalCheck{
		ObservedTexts:                        observed,
		InternalDiagnosticTexts:              internalDiagnostics,
		InternalDiagnosticsClassifiedBlocked: val0.NoOverclaimGuard.InternalDiagnosticsClassifiedBlocked && valA.NoOverclaimGuard.InternalDiagnosticsClassifiedBlocked && valB.NoOverclaimGuard.InternalDiagnosticsClassifiedBlocked && valC.NoOverclaimGuard.InternalDiagnosticsClassifiedBlocked && valD.NoOverclaimGuard.InternalDiagnosticsClassifiedBlocked,
		AllowedSafeWording:                   point15ValESafeWording(),
		BlockedWording:                       point15ValEForbiddenWording(),
		ProjectionDisclaimer:                 point15ValEClosureDisclaimer,
	}
}

func EvaluatePoint15ValENoOverclaimFinalCheckState(model Point15ValENoOverclaimFinalCheck) string {
	if model.ProjectionDisclaimer != point15ValEClosureDisclaimer ||
		!point12Val0ExactStringSetMatch(model.AllowedSafeWording, point15ValESafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point15ValEForbiddenWording()) {
		return Point15ValEStateBlocked
	}
	if point15ValEObservedListContainsForbiddenWording(model.ObservedTexts) {
		return Point15ValEStateBlocked
	}
	if point15ValEInternalListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point15ValEStateBlocked
	}
	return Point15ValEStatePassConfirmed
}

func point15ValECLBFinalCheckModel() Point15ValECLBFinalCheck {
	return Point15ValECLBFinalCheck{
		CheckID: "point15_vale_check_clb_001",
	}
}

func EvaluatePoint15ValECLBFinalCheckState(model Point15ValECLBFinalCheck) string {
	if !point15ValECheckIDValid(model.CheckID) {
		return Point15ValEStateBlocked
	}
	if model.CLB0Present || model.CLB1Present || model.CLB2Present {
		return Point15ValEStateBlocked
	}
	return Point15ValEStatePassConfirmed
}

func point15ValEClosureEvaluatorModel() Point15ValEClosureEvaluator {
	return Point15ValEClosureEvaluator{
		ClosureEvaluatorID:          "point15_vale_closure_evaluator_001",
		ReadOnlyProjectionConfirmed: true,
		NoMutationPathsDetected:     true,
		NoExternalAuthorityDetected: true,
		ReplayableManifestReady:     true,
		NoPrematurePoint15Pass:      true,
		CommandsRun:                 point15ValECommandsRun(),
		TestsRun:                    point15ValETestsRun(),
		GrepsRun:                    point15ValEGrepsRun(),
		NegativeFixturesRun:         point15ValENegativeFixturesRun(),
		ReviewerResult:              point12ValEReviewerResultPassConfirmed,
		ProjectionDisclaimer:        point15ValEClosureDisclaimer,
	}
}

func EvaluatePoint15ValEClosureEvaluatorState(model Point15ValEClosureEvaluator) string {
	if !point15ValEClosureEvaluatorIDValid(model.ClosureEvaluatorID) ||
		!point15ValEStateValid(model.DependencyState) ||
		!point15ValEStateValid(model.FreshnessTaxonomyState) ||
		!point15ValEStateValid(model.DowngradeTriggerState) ||
		!point15ValEStateValid(model.ScheduledRevalidationState) ||
		!point15ValEStateValid(model.EnforcementBoundaryState) ||
		!point15ValEStateValid(model.ProjectionBoundaryState) ||
		!point15ValEStateValid(model.ReplayProofHistoryState) ||
		!point15ValEStateValid(model.TenantPrivacyState) ||
		!point15ValEStateValid(model.TimestampIntegrityState) ||
		!point15ValEStateValid(model.AuthorityBoundaryState) ||
		!point15ValEStateValid(model.NoMutationState) ||
		!point15ValEStateValid(model.NoOverclaimState) ||
		!point15ValEStateValid(model.CLBFinalState) ||
		!point15ValECommandsRunValid(model.CommandsRun) ||
		!point15ValETestsRunValid(model.TestsRun) ||
		!point15ValEGrepsRunValid(model.GrepsRun) ||
		!point15ValENegativeFixturesRunValid(model.NegativeFixturesRun) ||
		!point12ValEReviewerResultValid(model.ReviewerResult) ||
		model.ReviewerResult != point12ValEReviewerResultPassConfirmed ||
		model.ProjectionDisclaimer != point15ValEClosureDisclaimer {
		return Point15ValEStateBlocked
	}
	if !model.ReadOnlyProjectionConfirmed || !model.NoMutationPathsDetected || !model.NoExternalAuthorityDetected || !model.ReplayableManifestReady || !model.NoPrematurePoint15Pass {
		return Point15ValEStateBlocked
	}
	componentState := point15ValEComponentAggregate(
		model.DependencyState,
		model.FreshnessTaxonomyState,
		model.DowngradeTriggerState,
		model.ScheduledRevalidationState,
		model.EnforcementBoundaryState,
		model.ProjectionBoundaryState,
		model.ReplayProofHistoryState,
		model.TenantPrivacyState,
		model.TimestampIntegrityState,
		model.AuthorityBoundaryState,
		model.NoMutationState,
		model.NoOverclaimState,
		model.CLBFinalState,
	)
	if componentState != Point15ValEStatePassConfirmed {
		return componentState
	}
	if !model.FinalPassAllowed {
		return Point15ValEStateBlocked
	}
	return Point15ValEStatePassConfirmed
}

func point15ValEEvidenceIdentity(dependency Point15ValEDependencySnapshot) string {
	val0 := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0
	return point15ValEManifestEvidenceIdentity(
		val0.EvidenceContext.EvidenceID,
		val0.EvidenceContext.EvidenceHash,
		val0.EvidenceContext.PolicyVersion,
		val0.EvidenceContext.EngineVersion,
		val0.EvidenceContext.SchemaVersion,
		dependency.InheritedTenantScope,
	)
}

func point15ValEManifestEvidenceIdentity(evidenceID, evidenceHash, policyVersion, engineVersion, schemaVersion, tenantScope string) string {
	return fmt.Sprintf("evidence_id=%s evidence_hash=%s policy=%s engine=%s schema=%s tenant=%s",
		evidenceID,
		evidenceHash,
		policyVersion,
		engineVersion,
		schemaVersion,
		tenantScope,
	)
}

func point15ValEPassClosureManifestModel(dependency Point15ValEDependencySnapshot) Point15PassClosureManifest {
	val0 := dependency.Point15ValD.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0
	return Point15PassClosureManifest{
		ClosureManifestID:    "point15_vale_pass_manifest_001",
		PointID:              point15Val0PointID,
		WaveID:               point15ValEWaveID,
		ClosureToken:         point15Val0BlockedPassToken,
		Scope:                point15ValEScope,
		ExplicitNonGoals:     point15ValEExplicitNonGoals(),
		EvidenceIdentity:     point15ValEEvidenceIdentity(dependency),
		EvidenceHash:         val0.EvidenceContext.EvidenceHash,
		PolicyVersion:        val0.EvidenceContext.PolicyVersion,
		EngineVersion:        val0.EvidenceContext.EngineVersion,
		SchemaVersion:        val0.EvidenceContext.SchemaVersion,
		TenantScope:          dependency.InheritedTenantScope,
		CommandsRun:          point15ValECommandsRun(),
		TestsRun:             point15ValETestsRun(),
		GrepsRun:             point15ValEGrepsRun(),
		NegativeFixturesRun:  point15ValENegativeFixturesRun(),
		CleanRoomIPResult:    point15ValECleanRoomIPPreserved,
		ReviewerResult:       point12ValEReviewerResultPassConfirmed,
		EvidenceID:           val0.EvidenceContext.EvidenceID,
		GeneratedAt:          dependency.Point15ValD.TimestampDisplayDiscipline.ReferenceNow,
		ProjectionDisclaimer: point15ValEClosureDisclaimer,
	}
}

func point15ValEManifestEvidenceIdentityMatchesFields(identity, evidenceID, evidenceHash, policyVersion, engineVersion, schemaVersion, tenantScope string) bool {
	if !formalRawExactNonEmpty(identity) {
		return false
	}
	parts := strings.Split(identity, " ")
	if len(parts) != 6 {
		return false
	}
	values := map[string]string{}
	expectedKeys := []string{"evidence_id", "evidence_hash", "policy", "engine", "schema", "tenant"}
	for idx, part := range parts {
		key, value, ok := strings.Cut(part, "=")
		if !ok || key != expectedKeys[idx] || !formalRawExactNonEmpty(value) {
			return false
		}
		if _, exists := values[key]; exists {
			return false
		}
		values[key] = value
	}
	return values["evidence_id"] == evidenceID &&
		values["evidence_hash"] == evidenceHash &&
		values["policy"] == policyVersion &&
		values["engine"] == engineVersion &&
		values["schema"] == schemaVersion &&
		values["tenant"] == tenantScope
}

func EvaluatePoint15PassClosureManifestState(model Point15PassClosureManifest) string {
	if !formalRawExactValid(model.ClosureManifestID, point15ValEClosureManifestIDValid) ||
		model.PointID != point15Val0PointID ||
		model.WaveID != point15ValEWaveID ||
		model.ClosureToken != point15Val0BlockedPassToken ||
		model.Scope != point15ValEScope ||
		!point15ValEExplicitNonGoalsValid(model.ExplicitNonGoals) ||
		!formalRawExactValid(model.DependencyGateResult, point15ValEStateValid) ||
		!formalRawExactValid(model.FreshnessTaxonomyResult, point15ValEStateValid) ||
		!formalRawExactValid(model.DowngradeTriggerResult, point15ValEStateValid) ||
		!formalRawExactValid(model.ScheduledRevalidationResult, point15ValEStateValid) ||
		!formalRawExactValid(model.EnforcementBoundaryResult, point15ValEStateValid) ||
		!formalRawExactValid(model.ProjectionBoundaryResult, point15ValEStateValid) ||
		!formalRawExactValid(model.ReplayProofHistoryResult, point15ValEStateValid) ||
		!formalRawExactValid(model.TenantPrivacyResult, point15ValEStateValid) ||
		!formalRawExactValid(model.TimestampIntegrityResult, point15ValEStateValid) ||
		!formalRawExactValid(model.AuthorityBoundaryResult, point15ValEStateValid) ||
		!formalRawExactValid(model.NoMutationResult, point15ValEStateValid) ||
		!formalRawExactValid(model.NoOverclaimResult, point15ValEStateValid) ||
		!formalRawExactValid(model.CLBResult, point15ValEStateValid) ||
		!formalRawExactNonEmpty(model.EvidenceID) ||
		!formalRawExactNonEmpty(model.EvidenceIdentity) ||
		!formalRawExactNonEmpty(model.EvidenceHash) ||
		!formalRawExactNonEmpty(model.PolicyVersion) ||
		!formalRawExactNonEmpty(model.EngineVersion) ||
		!formalRawExactNonEmpty(model.SchemaVersion) ||
		!formalRawExactValid(model.TenantScope, point11Val0ScopeValid) ||
		!point15ValEManifestEvidenceIdentityMatchesFields(model.EvidenceIdentity, model.EvidenceID, model.EvidenceHash, model.PolicyVersion, model.EngineVersion, model.SchemaVersion, model.TenantScope) ||
		!point15ValECommandsRunValid(model.CommandsRun) ||
		!point15ValETestsRunValid(model.TestsRun) ||
		!point15ValEGrepsRunValid(model.GrepsRun) ||
		!point15ValENegativeFixturesRunValid(model.NegativeFixturesRun) ||
		model.CleanRoomIPResult != point15ValECleanRoomIPPreserved ||
		!formalRawExactValid(model.ReviewerResult, point12ValEReviewerResultValid) ||
		model.ReviewerResult != point12ValEReviewerResultPassConfirmed ||
		!point15ValERawCanonicalTimeValid(model.GeneratedAt) ||
		model.ProjectionDisclaimer != point15ValEClosureDisclaimer {
		return Point15ValEStateBlocked
	}
	componentState := point15ValEComponentAggregate(
		model.DependencyGateResult,
		model.FreshnessTaxonomyResult,
		model.DowngradeTriggerResult,
		model.ScheduledRevalidationResult,
		model.EnforcementBoundaryResult,
		model.ProjectionBoundaryResult,
		model.ReplayProofHistoryResult,
		model.TenantPrivacyResult,
		model.TimestampIntegrityResult,
		model.AuthorityBoundaryResult,
		model.NoMutationResult,
		model.NoOverclaimResult,
		model.CLBResult,
	)
	if componentState == Point15ValEStateBlocked {
		return Point15ValEStateBlocked
	}
	if componentState == Point15ValEStateReviewRequired {
		if model.Point15PassAllowed || model.Point15PassToken != "" {
			return Point15ValEStateBlocked
		}
		return Point15ValEStateReviewRequired
	}
	if componentState == Point15ValEStateIncomplete {
		if model.Point15PassAllowed || model.Point15PassToken != "" {
			return Point15ValEStateBlocked
		}
		return Point15ValEStateIncomplete
	}
	if !model.Point15PassAllowed || model.Point15PassToken != point15Val0BlockedPassToken {
		return Point15ValEStateBlocked
	}
	return Point15ValEStatePassConfirmed
}

func point15ValERawCanonicalTimeValid(value string) bool {
	if !formalRawExactNonEmpty(value) {
		return false
	}
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return false
	}
	return parsed.UTC().Format(time.RFC3339) == value
}

func point15ValEPassManifestGeneratedAtBound(manifest Point15PassClosureManifest, dependency Point15ValEDependencySnapshot) bool {
	expected := dependency.Point15ValD.TimestampDisplayDiscipline.ReferenceNow
	return expected != "" &&
		manifest.GeneratedAt == expected &&
		point15ValERawCanonicalTimeValid(manifest.GeneratedAt)
}

func point15ValEFoundationModelFromUpstream(valD Point15ValDAssuranceProjectionFoundation) Point15ValEContinuousVerificationClosureFoundation {
	dependency := point15ValEDependencySnapshotFromUpstream(valD)
	return Point15ValEContinuousVerificationClosureFoundation{
		ProjectionDisclaimer:              point15ValEClosureDisclaimer,
		Dependency:                        dependency,
		ClosureEvaluator:                  point15ValEClosureEvaluatorModel(),
		PassClosureManifest:               point15ValEPassClosureManifestModel(dependency),
		FreshnessTaxonomyClosureCheck:     point15ValEFreshnessTaxonomyClosureCheckModel(dependency),
		DowngradeTriggerClosureCheck:      point15ValEDowngradeTriggerClosureCheckModel(dependency),
		ScheduledRevalidationClosureCheck: point15ValEScheduledRevalidationClosureCheckModel(dependency),
		EnforcementClosureCheck:           point15ValEEnforcementClosureCheckModel(dependency),
		ProjectionClosureCheck:            point15ValEProjectionClosureCheckModel(dependency),
		ReplayProofHistoryClosureCheck:    point15ValEReplayProofHistoryClosureCheckModel(dependency),
		TenantPrivacyClosureCheck:         point15ValETenantPrivacyClosureCheckModel(dependency),
		TimestampIntegrityClosureCheck:    point15ValETimestampIntegrityClosureCheckModel(dependency),
		AuthorityBoundaryClosureCheck:     point15ValEAuthorityBoundaryClosureCheckModel(dependency),
		NoMutationClosureCheck:            point15ValENoMutationClosureCheckModel(dependency),
		NoOverclaimFinalCheck:             point15ValENoOverclaimFinalCheckModel(dependency),
		CLBFinalCheck:                     point15ValECLBFinalCheckModel(),
	}
}

func Point15ValEFoundationModel() Point15ValEContinuousVerificationClosureFoundation {
	valD := ComputePoint15ValDAssuranceProjectionFoundation(Point15ValDFoundationModel())
	return point15ValEFoundationModelFromUpstream(valD)
}

func point15ValEBlockingReasons(model Point15ValEContinuousVerificationClosureFoundation) []string {
	componentStates := map[string]string{
		"dependency":             model.DependencyState,
		"freshness_taxonomy":     model.FreshnessTaxonomyClosureState,
		"downgrade_trigger":      model.DowngradeTriggerClosureState,
		"scheduled_revalidation": model.ScheduledRevalidationClosureState,
		"enforcement":            model.EnforcementClosureState,
		"projection":             model.ProjectionClosureState,
		"replay_proof_history":   model.ReplayProofHistoryClosureState,
		"tenant_privacy":         model.TenantPrivacyClosureState,
		"timestamp_integrity":    model.TimestampIntegrityClosureState,
		"authority_boundary":     model.AuthorityBoundaryClosureState,
		"no_mutation":            model.NoMutationClosureState,
		"no_overclaim":           model.NoOverclaimFinalCheckState,
		"clb":                    model.CLBFinalCheckState,
		"closure_evaluator":      model.ClosureEvaluatorState,
		"pass_closure_manifest":  model.PassClosureManifestState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if !formalRawExactValid(state, point15ValEStateValid) || state == Point15ValEStateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point15ValEReviewPrerequisites(model Point15ValEContinuousVerificationClosureFoundation) []string {
	componentStates := map[string]string{
		"freshness_taxonomy":     model.FreshnessTaxonomyClosureState,
		"downgrade_trigger":      model.DowngradeTriggerClosureState,
		"scheduled_revalidation": model.ScheduledRevalidationClosureState,
		"enforcement":            model.EnforcementClosureState,
		"projection":             model.ProjectionClosureState,
		"replay_proof_history":   model.ReplayProofHistoryClosureState,
		"tenant_privacy":         model.TenantPrivacyClosureState,
		"timestamp_integrity":    model.TimestampIntegrityClosureState,
		"authority_boundary":     model.AuthorityBoundaryClosureState,
		"no_mutation":            model.NoMutationClosureState,
		"no_overclaim":           model.NoOverclaimFinalCheckState,
		"clb":                    model.CLBFinalCheckState,
	}
	prereqs := append([]string{}, model.Dependency.ReviewPrerequisites...)
	for name, state := range componentStates {
		if state == Point15ValEStateReviewRequired || state == Point15ValEStateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func ComputePoint15ValEFoundation(model Point15ValEContinuousVerificationClosureFoundation) Point15ValEContinuousVerificationClosureFoundation {
	model.DependencyState = EvaluatePoint15ValEDependencyState(model.Dependency)
	model.FreshnessTaxonomyClosureState = EvaluatePoint15ValEFreshnessTaxonomyClosureCheckState(model.FreshnessTaxonomyClosureCheck)
	model.DowngradeTriggerClosureState = EvaluatePoint15ValEDowngradeTriggerClosureCheckState(model.DowngradeTriggerClosureCheck)
	model.ScheduledRevalidationClosureState = EvaluatePoint15ValEScheduledRevalidationClosureCheckState(model.ScheduledRevalidationClosureCheck)
	model.EnforcementClosureState = EvaluatePoint15ValEEnforcementClosureCheckState(model.EnforcementClosureCheck)
	model.ProjectionClosureState = EvaluatePoint15ValEProjectionClosureCheckState(model.ProjectionClosureCheck)
	model.ReplayProofHistoryClosureState = EvaluatePoint15ValEReplayProofHistoryClosureCheckState(model.ReplayProofHistoryClosureCheck)
	model.TenantPrivacyClosureState = EvaluatePoint15ValETenantPrivacyClosureCheckState(model.TenantPrivacyClosureCheck)
	model.TimestampIntegrityClosureState = EvaluatePoint15ValETimestampIntegrityClosureCheckState(model.TimestampIntegrityClosureCheck)
	model.AuthorityBoundaryClosureState = EvaluatePoint15ValEAuthorityBoundaryClosureCheckState(model.AuthorityBoundaryClosureCheck)
	model.NoMutationClosureState = EvaluatePoint15ValENoMutationClosureCheckState(model.NoMutationClosureCheck)
	model.NoOverclaimFinalCheckState = EvaluatePoint15ValENoOverclaimFinalCheckState(model.NoOverclaimFinalCheck)
	model.CLBFinalCheckState = EvaluatePoint15ValECLBFinalCheckState(model.CLBFinalCheck)

	model.FreshnessTaxonomyClosureState = point15ValEComponentAggregate(
		model.FreshnessTaxonomyClosureState,
		EvaluatePoint15ValEFreshnessTaxonomyClosureCheckState(point15ValEFreshnessTaxonomyClosureCheckModel(model.Dependency)),
	)
	model.DowngradeTriggerClosureState = point15ValEComponentAggregate(
		model.DowngradeTriggerClosureState,
		EvaluatePoint15ValEDowngradeTriggerClosureCheckState(point15ValEDowngradeTriggerClosureCheckModel(model.Dependency)),
	)
	model.ScheduledRevalidationClosureState = point15ValEComponentAggregate(
		model.ScheduledRevalidationClosureState,
		EvaluatePoint15ValEScheduledRevalidationClosureCheckState(point15ValEScheduledRevalidationClosureCheckModel(model.Dependency)),
	)
	model.EnforcementClosureState = point15ValEComponentAggregate(
		model.EnforcementClosureState,
		EvaluatePoint15ValEEnforcementClosureCheckState(point15ValEEnforcementClosureCheckModel(model.Dependency)),
	)
	model.ProjectionClosureState = point15ValEComponentAggregate(
		model.ProjectionClosureState,
		EvaluatePoint15ValEProjectionClosureCheckState(point15ValEProjectionClosureCheckModel(model.Dependency)),
	)
	model.ReplayProofHistoryClosureState = point15ValEComponentAggregate(
		model.ReplayProofHistoryClosureState,
		EvaluatePoint15ValEReplayProofHistoryClosureCheckState(point15ValEReplayProofHistoryClosureCheckModel(model.Dependency)),
	)
	model.TenantPrivacyClosureState = point15ValEComponentAggregate(
		model.TenantPrivacyClosureState,
		EvaluatePoint15ValETenantPrivacyClosureCheckState(point15ValETenantPrivacyClosureCheckModel(model.Dependency)),
	)
	model.TimestampIntegrityClosureState = point15ValEComponentAggregate(
		model.TimestampIntegrityClosureState,
		EvaluatePoint15ValETimestampIntegrityClosureCheckState(point15ValETimestampIntegrityClosureCheckModel(model.Dependency)),
	)
	model.AuthorityBoundaryClosureState = point15ValEComponentAggregate(
		model.AuthorityBoundaryClosureState,
		EvaluatePoint15ValEAuthorityBoundaryClosureCheckState(point15ValEAuthorityBoundaryClosureCheckModel(model.Dependency)),
	)
	model.NoMutationClosureState = point15ValEComponentAggregate(
		model.NoMutationClosureState,
		EvaluatePoint15ValENoMutationClosureCheckState(point15ValENoMutationClosureCheckModel(model.Dependency)),
	)
	model.NoOverclaimFinalCheckState = point15ValEComponentAggregate(
		model.NoOverclaimFinalCheckState,
		EvaluatePoint15ValENoOverclaimFinalCheckState(point15ValENoOverclaimFinalCheckModel(model.Dependency)),
	)

	valD := model.Dependency.Point15ValD
	valC := valD.Dependency.Point15ValC
	valB := valC.Dependency.Point15ValB
	valA := valB.Dependency.Point15ValA
	val0 := valA.Dependency.Point15Val0
	expectedTenant := model.Dependency.InheritedTenantScope
	expectedEvidenceID := val0.EvidenceContext.EvidenceID
	expectedEvidenceHash := val0.EvidenceContext.EvidenceHash
	expectedPolicy := val0.EvidenceContext.PolicyVersion
	expectedEngine := val0.EvidenceContext.EngineVersion
	expectedSchema := val0.EvidenceContext.SchemaVersion
	expectedPassManifestIdentity := point15ValEEvidenceIdentity(model.Dependency)
	passManifestDependencyBound := expectedTenant != "" &&
		expectedEvidenceID != "" &&
		model.PassClosureManifest.EvidenceID == expectedEvidenceID &&
		model.PassClosureManifest.EvidenceIdentity == expectedPassManifestIdentity &&
		model.PassClosureManifest.EvidenceHash == expectedEvidenceHash &&
		model.PassClosureManifest.PolicyVersion == expectedPolicy &&
		model.PassClosureManifest.EngineVersion == expectedEngine &&
		model.PassClosureManifest.SchemaVersion == expectedSchema &&
		model.PassClosureManifest.TenantScope == expectedTenant

	if !point15ValEValDProjectionContractDisplayOnly(valD) {
		model.ProjectionClosureState = Point15ValEStateBlocked
	}

	if expectedTenant == "" ||
		model.FreshnessTaxonomyClosureCheck.TenantScope != expectedTenant ||
		model.TenantPrivacyClosureCheck.TenantScope != expectedTenant ||
		model.TimestampIntegrityClosureCheck.TenantScope != expectedTenant ||
		model.AuthorityBoundaryClosureCheck.TenantScope != expectedTenant ||
		model.PassClosureManifest.TenantScope != expectedTenant {
		model.FreshnessTaxonomyClosureState = Point15ValEStateBlocked
		model.TenantPrivacyClosureState = Point15ValEStateBlocked
		model.TimestampIntegrityClosureState = Point15ValEStateBlocked
		model.AuthorityBoundaryClosureState = Point15ValEStateBlocked
		model.PassClosureManifestState = Point15ValEStateBlocked
	}
	if expectedEvidenceID == "" ||
		model.FreshnessTaxonomyClosureCheck.EvidenceID != expectedEvidenceID ||
		model.PassClosureManifest.EvidenceID != expectedEvidenceID ||
		model.PassClosureManifest.EvidenceIdentity != expectedPassManifestIdentity ||
		model.PassClosureManifest.EvidenceHash != expectedEvidenceHash ||
		model.PassClosureManifest.PolicyVersion != expectedPolicy ||
		model.PassClosureManifest.EngineVersion != expectedEngine ||
		model.PassClosureManifest.SchemaVersion != expectedSchema {
		model.FreshnessTaxonomyClosureState = Point15ValEStateBlocked
		model.PassClosureManifestState = Point15ValEStateBlocked
	}

	passCandidate := point15ValEComponentAggregate(
		model.DependencyState,
		model.FreshnessTaxonomyClosureState,
		model.DowngradeTriggerClosureState,
		model.ScheduledRevalidationClosureState,
		model.EnforcementClosureState,
		model.ProjectionClosureState,
		model.ReplayProofHistoryClosureState,
		model.TenantPrivacyClosureState,
		model.TimestampIntegrityClosureState,
		model.AuthorityBoundaryClosureState,
		model.NoMutationClosureState,
		model.NoOverclaimFinalCheckState,
		model.CLBFinalCheckState,
	) == Point15ValEStatePassConfirmed

	model.ClosureEvaluator.DependencyState = model.DependencyState
	model.ClosureEvaluator.FreshnessTaxonomyState = model.FreshnessTaxonomyClosureState
	model.ClosureEvaluator.DowngradeTriggerState = model.DowngradeTriggerClosureState
	model.ClosureEvaluator.ScheduledRevalidationState = model.ScheduledRevalidationClosureState
	model.ClosureEvaluator.EnforcementBoundaryState = model.EnforcementClosureState
	model.ClosureEvaluator.ProjectionBoundaryState = model.ProjectionClosureState
	model.ClosureEvaluator.ReplayProofHistoryState = model.ReplayProofHistoryClosureState
	model.ClosureEvaluator.TenantPrivacyState = model.TenantPrivacyClosureState
	model.ClosureEvaluator.TimestampIntegrityState = model.TimestampIntegrityClosureState
	model.ClosureEvaluator.AuthorityBoundaryState = model.AuthorityBoundaryClosureState
	model.ClosureEvaluator.NoMutationState = model.NoMutationClosureState
	model.ClosureEvaluator.NoOverclaimState = model.NoOverclaimFinalCheckState
	model.ClosureEvaluator.CLBFinalState = model.CLBFinalCheckState
	model.ClosureEvaluator.ReadOnlyProjectionConfirmed = model.ProjectionClosureState != Point15ValEStateBlocked
	model.ClosureEvaluator.NoMutationPathsDetected = model.NoMutationClosureState != Point15ValEStateBlocked
	model.ClosureEvaluator.NoExternalAuthorityDetected = model.AuthorityBoundaryClosureState != Point15ValEStateBlocked
	model.ClosureEvaluator.ReplayableManifestReady = model.ReplayProofHistoryClosureState != Point15ValEStateBlocked && model.TimestampIntegrityClosureState != Point15ValEStateBlocked
	model.ClosureEvaluator.NoPrematurePoint15Pass = !model.Dependency.Point15PassSeen
	model.ClosureEvaluator.FinalPassAllowed = passCandidate
	model.ClosureEvaluatorState = EvaluatePoint15ValEClosureEvaluatorState(model.ClosureEvaluator)

	model.PassClosureManifest.DependencyGateResult = model.DependencyState
	model.PassClosureManifest.FreshnessTaxonomyResult = model.FreshnessTaxonomyClosureState
	model.PassClosureManifest.DowngradeTriggerResult = model.DowngradeTriggerClosureState
	model.PassClosureManifest.ScheduledRevalidationResult = model.ScheduledRevalidationClosureState
	model.PassClosureManifest.EnforcementBoundaryResult = model.EnforcementClosureState
	model.PassClosureManifest.ProjectionBoundaryResult = model.ProjectionClosureState
	model.PassClosureManifest.ReplayProofHistoryResult = model.ReplayProofHistoryClosureState
	model.PassClosureManifest.TenantPrivacyResult = model.TenantPrivacyClosureState
	model.PassClosureManifest.TimestampIntegrityResult = model.TimestampIntegrityClosureState
	model.PassClosureManifest.AuthorityBoundaryResult = model.AuthorityBoundaryClosureState
	model.PassClosureManifest.NoMutationResult = model.NoMutationClosureState
	model.PassClosureManifest.NoOverclaimResult = model.NoOverclaimFinalCheckState
	model.PassClosureManifest.CLBResult = model.CLBFinalCheckState
	model.PassClosureManifest.Point15PassAllowed = model.ClosureEvaluatorState == Point15ValEStatePassConfirmed
	if model.PassClosureManifest.Point15PassAllowed {
		model.PassClosureManifest.Point15PassToken = point15Val0BlockedPassToken
	} else {
		model.PassClosureManifest.Point15PassToken = ""
	}
	model.PassClosureManifestState = EvaluatePoint15PassClosureManifestState(model.PassClosureManifest)
	generatedAtBound := point15ValEPassManifestGeneratedAtBound(model.PassClosureManifest, model.Dependency)
	if !generatedAtBound || !passManifestDependencyBound {
		model.PassClosureManifestState = Point15ValEStateBlocked
	}

	model.CurrentState = point15ValEComponentAggregate(
		model.DependencyState,
		model.FreshnessTaxonomyClosureState,
		model.DowngradeTriggerClosureState,
		model.ScheduledRevalidationClosureState,
		model.EnforcementClosureState,
		model.ProjectionClosureState,
		model.ReplayProofHistoryClosureState,
		model.TenantPrivacyClosureState,
		model.TimestampIntegrityClosureState,
		model.AuthorityBoundaryClosureState,
		model.NoMutationClosureState,
		model.NoOverclaimFinalCheckState,
		model.CLBFinalCheckState,
		model.ClosureEvaluatorState,
		model.PassClosureManifestState,
	)
	if model.CurrentState != Point15ValEStatePassConfirmed {
		model.PassClosureManifest.Point15PassAllowed = false
		model.PassClosureManifest.Point15PassToken = ""
		model.PassClosureManifestState = EvaluatePoint15PassClosureManifestState(model.PassClosureManifest)
		if !generatedAtBound || !passManifestDependencyBound {
			model.PassClosureManifestState = Point15ValEStateBlocked
		}
		model.CurrentState = point15ValEComponentAggregate(
			model.DependencyState,
			model.FreshnessTaxonomyClosureState,
			model.DowngradeTriggerClosureState,
			model.ScheduledRevalidationClosureState,
			model.EnforcementClosureState,
			model.ProjectionClosureState,
			model.ReplayProofHistoryClosureState,
			model.TenantPrivacyClosureState,
			model.TimestampIntegrityClosureState,
			model.AuthorityBoundaryClosureState,
			model.NoMutationClosureState,
			model.NoOverclaimFinalCheckState,
			model.CLBFinalCheckState,
			model.ClosureEvaluatorState,
			model.PassClosureManifestState,
		)
	}
	model.BlockingReasons = point15ValEBlockingReasons(model)
	model.ReviewPrerequisites = point15ValEReviewPrerequisites(model)
	return model
}

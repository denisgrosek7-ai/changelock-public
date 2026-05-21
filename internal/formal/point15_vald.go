package formal

import (
	"encoding/json"
	"sort"
	"strings"
)

const (
	Point15ValDStateActive         = "point15_vald_assurance_projection_active"
	Point15ValDStateBlocked        = "point15_vald_assurance_projection_blocked"
	Point15ValDStateReviewRequired = "point15_vald_assurance_projection_review_required"
	Point15ValDStateIncomplete     = "point15_vald_assurance_projection_incomplete"
)

const (
	point15ValDWaveID                  = "val_d"
	point15ValDProjectionDisclaimer    = "formal_assurance_projection_only no_mutation_or_pass_restore point15_vald"
	point15ValDBlockedPassToken        = "point_15_pass"
	point15ValDModeTimeline            = "timeline"
	point15ValDModeDashboardSummary    = "dashboard_summary"
	point15ValDModeEvidenceDetail      = "evidence_detail"
	point15ValDModeRevalidationDetail  = "revalidation_detail"
	point15ValDModeEnforcementDetail   = "enforcement_detail"
	point15ValDModeReplayDetail        = "replay_detail"
	point15ValDModeExportPreview       = "export_preview"
	point15ValDModeQueryResult         = "query_result"
	point15ValDVisibilityInternalOnly  = "internal_only"
	point15ValDVisibilityTenantScoped  = "tenant_scoped"
	point15ValDVisibilityAuditorRead   = "auditor_read_only"
	point15ValDVisibilityCustomerRead  = "customer_read_only"
	point15ValDVisibilityPublicBlocked = "public_forbidden"
	point15ValDActionDisplayOnly       = "display_only"
	point15ValDActionFilterOnly        = "filter_only"
	point15ValDActionSortOnly          = "sort_only"
	point15ValDActionExplainOnly       = "explain_only"
	point15ValDActionExportPreviewOnly = "export_preview_only"
	point15ValDQueryComplete           = "complete"
	point15ValDQueryPartial            = "partial"
	point15ValDQueryRedacted           = "redacted"
	point15ValDQueryBlocked            = "blocked"
	point15ValDQueryReviewRequired     = "review_required"
	point15ValDQueryIncomplete         = "incomplete"
	point15ValDEventFreshness          = "freshness_transition"
	point15ValDEventRevalidation       = "revalidation_transition"
	point15ValDEventEnforcement        = "enforcement_transition"
	point15ValDEventRevocation         = "revocation_recorded"
	point15ValDEventExpiry             = "expiry_recorded"
	point15ValDEventSupersession       = "supersession_recorded"
	point15ValDEventReplay             = "replay_history_visible"
	point15ValDSummaryScopeTenant      = "tenant_assurance_summary"
	point15ValDQueryScopeHistory       = "tenant_assurance_history"
	point15ValDQueryScopeEvidence      = "tenant_evidence_projection"
	point15ValDQueryScopeRevalidation  = "tenant_revalidation_projection"
	point15ValDQueryScopeEnforcement   = "tenant_enforcement_projection"
	point15ValDSortEventDesc           = "event_at_desc"
	point15ValDSortEventAsc            = "event_at_asc"
	point15ValDSortStateSeverity       = "state_severity"
	point15ValDSortEvidenceID          = "evidence_id"
	point15ValDFilterTenantScope       = "tenant_scope"
	point15ValDFilterEvidenceID        = "evidence_id"
	point15ValDFilterCurrentState      = "current_state"
	point15ValDFilterDowngradeReason   = "downgrade_reason"
	point15ValDFilterEnforcementReason = "enforcement_reason"
	point15ValDFilterFreshnessStatus   = "freshness_status"
	point15ValDFilterEventType         = "event_type"
	point15ValDRedactionNone           = "none"
	point15ValDRedactionLimited        = "limited"
	point15ValDRedactionTenantPrivate  = "tenant_private_redacted"
	point15ValDRedactionBlockedPrivate = "blocked_private_exposure"
)

type Point15ValDDependencySnapshot struct {
	Point15ValCCurrentState           string                                   `json:"point15_valc_current_state"`
	Point15ValCDependencyState        string                                   `json:"point15_valc_dependency_state"`
	Point15ValCEnforcementActionState string                                   `json:"point15_valc_enforcement_action_state"`
	Point15ValCEvidenceLifecycleState string                                   `json:"point15_valc_evidence_lifecycle_state"`
	Point15ValCRevocationState        string                                   `json:"point15_valc_revocation_state"`
	Point15ValCExpiryState            string                                   `json:"point15_valc_expiry_state"`
	Point15ValCSupersessionState      string                                   `json:"point15_valc_supersession_state"`
	Point15ValCReplayHistoryState     string                                   `json:"point15_valc_replay_history_state"`
	Point15ValCTimestampState         string                                   `json:"point15_valc_timestamp_state"`
	Point15ValCAuthorityState         string                                   `json:"point15_valc_authority_state"`
	Point15ValCTenantState            string                                   `json:"point15_valc_tenant_state"`
	Point15ValCNoOverclaimState       string                                   `json:"point15_valc_no_overclaim_state"`
	Point15ValCComputedFromUpstream   bool                                     `json:"point15_valc_computed_from_upstream"`
	Point15ValCMerged                 bool                                     `json:"point15_valc_merged"`
	Point15ValCCIGreen                bool                                     `json:"point15_valc_ci_green"`
	Point15ValCReviewedOnMain         bool                                     `json:"point15_valc_reviewed_on_main"`
	Point15PassSeen                   bool                                     `json:"point15_pass_seen"`
	InheritedPoint15ValBCurrentState  string                                   `json:"inherited_point15_valb_current_state"`
	InheritedPoint15ValACurrentState  string                                   `json:"inherited_point15_vala_current_state"`
	InheritedPoint15Val0CurrentState  string                                   `json:"inherited_point15_val0_current_state"`
	InheritedPoint14ValECurrentState  string                                   `json:"inherited_point14_vale_current_state"`
	InheritedTenantScope              string                                   `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput        bool                                     `json:"snapshot_from_computed_output"`
	ReviewPrerequisites               []string                                 `json:"review_prerequisites,omitempty"`
	Point15ValC                       Point15ValCEnforcementBoundaryFoundation `json:"point15_valc"`
}

type Point15ValDAssuranceProjectionFoundation struct {
	CurrentState               string                                  `json:"current_state"`
	BlockingReasons            []string                                `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites        []string                                `json:"review_prerequisites,omitempty"`
	ProjectionDisclaimer       string                                  `json:"projection_disclaimer"`
	DependencyState            string                                  `json:"dependency_state"`
	TimelineState              string                                  `json:"timeline_state"`
	DashboardState             string                                  `json:"dashboard_state"`
	QueryState                 string                                  `json:"query_state"`
	EvidenceDetailState        string                                  `json:"evidence_detail_state"`
	RevalidationDetailState    string                                  `json:"revalidation_detail_state"`
	EnforcementDetailState     string                                  `json:"enforcement_detail_state"`
	ReplayProofHistoryState    string                                  `json:"replay_proof_history_state"`
	AccessTenantState          string                                  `json:"access_tenant_state"`
	TimestampDisplayState      string                                  `json:"timestamp_display_state"`
	NoMutationState            string                                  `json:"no_mutation_state"`
	AuthorityBoundaryState     string                                  `json:"authority_boundary_state"`
	NoOverclaimState           string                                  `json:"no_overclaim_state"`
	Dependency                 Point15ValDDependencySnapshot           `json:"dependency"`
	Timeline                   Point15ValDAssuranceTimelineEntry       `json:"timeline"`
	Dashboard                  Point15ValDDashboardSummary             `json:"dashboard"`
	Query                      Point15ValDQueryProjection              `json:"query"`
	EvidenceDetail             Point15ValDEvidenceDetailProjection     `json:"evidence_detail"`
	RevalidationDetail         Point15ValDRevalidationDetailProjection `json:"revalidation_detail"`
	EnforcementDetail          Point15ValDEnforcementDetailProjection  `json:"enforcement_detail"`
	ReplayProofHistory         Point15ValDReplayProofHistoryProjection `json:"replay_proof_history"`
	AccessTenantPrivacy        Point15ValDAccessTenantPrivacyBoundary  `json:"access_tenant_privacy"`
	TimestampDisplayDiscipline Point15ValDTimestampDisplayDiscipline   `json:"timestamp_display"`
	NoMutationGuard            Point15ValDNoMutationProjectionGuard    `json:"no_mutation_guard"`
	AuthorityBoundary          Point15ValDAuthorityBoundary            `json:"authority_boundary"`
	NoOverclaimGuard           Point15ValDNoOverclaimGuard             `json:"no_overclaim_guard"`
}

type Point15ValDAssuranceTimelineEntry struct {
	TimelineID              string `json:"timeline_id"`
	EntryID                 string `json:"entry_id"`
	ProjectionMode          string `json:"projection_mode"`
	ProjectionAction        string `json:"projection_action"`
	Visibility              string `json:"visibility"`
	TenantScope             string `json:"tenant_scope"`
	EvidenceID              string `json:"evidence_id"`
	EventType               string `json:"event_type"`
	PriorState              string `json:"prior_state"`
	CurrentState            string `json:"current_state"`
	DowngradeReason         string `json:"downgrade_reason"`
	EnforcementReason       string `json:"enforcement_reason"`
	SourceValCRef           string `json:"source_valc_ref"`
	ReplayRef               string `json:"replay_ref"`
	ProofHistoryRef         string `json:"proof_history_ref"`
	EventAt                 string `json:"event_at"`
	DisplayedAt             string `json:"displayed_at"`
	TimeSource              string `json:"time_source"`
	DecisiveEvidenceVisible bool   `json:"decisive_evidence_visible"`
	BlockedReasonVisible    bool   `json:"blocked_reason_visible"`
	PriorStateVisible       bool   `json:"prior_state_visible"`
	CurrentStateVisible     bool   `json:"current_state_visible"`
	PriorPassVisible        bool   `json:"prior_pass_visible"`
	LaterDowngradeVisible   bool   `json:"later_downgrade_visible"`
	TimelineCreatesValidity bool   `json:"timeline_creates_validity"`
}

type Point15ValDDashboardSummary struct {
	DashboardID                   string `json:"dashboard_id"`
	ProjectionMode                string `json:"projection_mode"`
	ProjectionAction              string `json:"projection_action"`
	Visibility                    string `json:"visibility"`
	TenantScope                   string `json:"tenant_scope"`
	SummaryScope                  string `json:"summary_scope"`
	EvidenceCount                 int    `json:"evidence_count"`
	ActiveCount                   int    `json:"active_count"`
	BlockedCount                  int    `json:"blocked_count"`
	ReviewRequiredCount           int    `json:"review_required_count"`
	IncompleteCount               int    `json:"incomplete_count"`
	StaleCount                    int    `json:"stale_count"`
	ExpiredCount                  int    `json:"expired_count"`
	RevokedCount                  int    `json:"revoked_count"`
	EnforcementCount              int    `json:"enforcement_count"`
	HiddenBlockedCount            bool   `json:"hidden_blocked_count"`
	HiddenReviewRequiredCount     bool   `json:"hidden_review_required_count"`
	HiddenIncompleteCount         bool   `json:"hidden_incomplete_count"`
	RestoresActiveClosure         bool   `json:"restores_active_closure"`
	ActiveCountIncludesDisallowed bool   `json:"active_count_includes_disallowed"`
}

type Point15ValDQueryProjection struct {
	QueryID                 string   `json:"query_id"`
	ProjectionMode          string   `json:"projection_mode"`
	ProjectionAction        string   `json:"projection_action"`
	Visibility              string   `json:"visibility"`
	TenantScope             string   `json:"tenant_scope"`
	ViewerScope             string   `json:"viewer_scope"`
	RequestedScope          string   `json:"requested_scope"`
	Filters                 []string `json:"filters,omitempty"`
	SortOrder               string   `json:"sort_order"`
	ResultState             string   `json:"result_state"`
	ResultRefs              []string `json:"result_refs,omitempty"`
	RedactionState          string   `json:"redaction_state"`
	DecisiveEvidenceVisible bool     `json:"decisive_evidence_visible"`
	LimitationsVisible      bool     `json:"limitations_visible"`
	QueryMutationAttempted  bool     `json:"query_mutation_attempted"`
	StrengthensClaims       bool     `json:"strengthens_claims"`
	CrossTenantQuery        bool     `json:"cross_tenant_query"`
}

type Point15ValDEvidenceDetailProjection struct {
	ProjectionMode              string `json:"projection_mode"`
	ProjectionAction            string `json:"projection_action"`
	Visibility                  string `json:"visibility"`
	EvidenceID                  string `json:"evidence_id"`
	TenantScope                 string `json:"tenant_scope"`
	EvidenceHash                string `json:"evidence_hash"`
	PolicyVersion               string `json:"policy_version"`
	EngineVersion               string `json:"engine_version"`
	SchemaVersion               string `json:"schema_version"`
	FreshnessStatus             string `json:"freshness_status"`
	DowngradeOutcome            string `json:"downgrade_outcome"`
	LifecycleStatus             string `json:"lifecycle_status"`
	EnforcementStatus           string `json:"enforcement_status"`
	LimitationsVisible          bool   `json:"limitations_visible"`
	IdentityDerivedFromNameOnly bool   `json:"identity_derived_from_name_only"`
}

type Point15ValDRevalidationDetailProjection struct {
	ProjectionMode            string `json:"projection_mode"`
	ProjectionAction          string `json:"projection_action"`
	Visibility                string `json:"visibility"`
	ScheduleRef               string `json:"schedule_ref"`
	RunRef                    string `json:"run_ref"`
	RetryBudgetRef            string `json:"retry_budget_ref"`
	TenantThrottleRef         string `json:"tenant_throttle_ref"`
	ScheduledStatus           string `json:"scheduled_status"`
	RunResult                 string `json:"run_result"`
	RetryStatus               string `json:"retry_status"`
	ThrottleStatus            string `json:"throttle_status"`
	ScheduleMutationAttempted bool   `json:"schedule_mutation_attempted"`
	RetryTriggered            bool   `json:"retry_triggered"`
	RetryBudgetResetAttempted bool   `json:"retry_budget_reset_attempted"`
	MarksFresh                bool   `json:"marks_fresh"`
	RestoresActiveClosure     bool   `json:"restores_active_closure"`
}

type Point15ValDEnforcementDetailProjection struct {
	ProjectionMode       string `json:"projection_mode"`
	ProjectionAction     string `json:"projection_action"`
	Visibility           string `json:"visibility"`
	EnforcementRef       string `json:"enforcement_ref"`
	EnforcementAction    string `json:"enforcement_action"`
	EnforcementReason    string `json:"enforcement_reason"`
	ReasonDecisive       bool   `json:"reason_decisive"`
	TargetState          string `json:"target_state"`
	PriorState           string `json:"prior_state"`
	CurrentState         string `json:"current_state"`
	HistoryPreserved     bool   `json:"history_preserved"`
	BlockedReasonVisible bool   `json:"blocked_reason_visible"`
	PerformsEnforcement  bool   `json:"performs_enforcement"`
	AutoRevokes          bool   `json:"auto_revokes"`
	AutoPublishes        bool   `json:"auto_publishes"`
	DeletesEvidence      bool   `json:"deletes_evidence"`
	SilentReplacement    bool   `json:"silent_replacement"`
}

type Point15ValDReplayProofHistoryProjection struct {
	ProjectionMode          string `json:"projection_mode"`
	ProjectionAction        string `json:"projection_action"`
	Visibility              string `json:"visibility"`
	ReplayRef               string `json:"replay_ref"`
	ProofPackRef            string `json:"proof_pack_ref"`
	ProofHistoryRef         string `json:"proof_history_ref"`
	PriorStateVisible       bool   `json:"prior_state_visible"`
	CurrentStateVisible     bool   `json:"current_state_visible"`
	DecisiveEvidenceVisible bool   `json:"decisive_evidence_visible"`
	BlockedReasonVisible    bool   `json:"blocked_reason_visible"`
	HashBindingVisible      bool   `json:"hash_binding_visible"`
	ProofHistoryHidden      bool   `json:"proof_history_hidden"`
}

type Point15ValDAccessTenantPrivacyBoundary struct {
	BoundaryID               string `json:"boundary_id"`
	TenantScope              string `json:"tenant_scope"`
	ViewerScope              string `json:"viewer_scope"`
	Visibility               string `json:"visibility"`
	RedactionState           string `json:"redaction_state"`
	TenantPrivateDataExposed bool   `json:"tenant_private_data_exposed"`
	CrossTenantDetected      bool   `json:"cross_tenant_detected"`
	DecisiveFailureHidden    bool   `json:"decisive_failure_hidden"`
	ProjectionStateMutated   bool   `json:"projection_state_mutated"`
}

type Point15ValDTimestampDisplayDiscipline struct {
	DisciplineID                 string `json:"discipline_id"`
	ProjectionMode               string `json:"projection_mode"`
	TenantScope                  string `json:"tenant_scope"`
	EventAt                      string `json:"event_at"`
	DisplayedAt                  string `json:"displayed_at"`
	SourceEventAt                string `json:"source_event_at"`
	ReceivedAt                   string `json:"received_at"`
	ValidatedAt                  string `json:"validated_at"`
	EnforcedAt                   string `json:"enforced_at"`
	ReferenceNow                 string `json:"reference_now"`
	TimeSource                   string `json:"time_source"`
	ClientLocalCreatesCanonical  bool   `json:"client_local_creates_canonical"`
	SourceEventCreatesCanonical  bool   `json:"source_event_creates_canonical"`
	CanonicalOrderingFromDisplay bool   `json:"canonical_ordering_from_display"`
}

type Point15ValDNoMutationProjectionGuard struct {
	GuardID                          string `json:"guard_id"`
	EvidenceMutationAttempted        bool   `json:"evidence_mutation_attempted"`
	LifecycleMutationAttempted       bool   `json:"lifecycle_mutation_attempted"`
	EnforcementMutationAttempted     bool   `json:"enforcement_mutation_attempted"`
	RevocationExecutionAttempted     bool   `json:"revocation_execution_attempted"`
	ExpiryDeletionAttempted          bool   `json:"expiry_deletion_attempted"`
	SupersessionReplacementAttempted bool   `json:"supersession_replacement_attempted"`
	ScheduleRetryMutationAttempted   bool   `json:"schedule_retry_mutation_attempted"`
	PassRestoreAttempted             bool   `json:"pass_restore_attempted"`
}

type Point15ValDAuthorityBoundary struct {
	BoundaryID                string `json:"boundary_id"`
	TenantScope               string `json:"tenant_scope"`
	FormalCoreOnly            bool   `json:"formal_core_only"`
	DashboardApprovesPass     bool   `json:"dashboard_approves_pass"`
	TimelineCreatesAuthority  bool   `json:"timeline_creates_authority"`
	QueryEnforcesState        bool   `json:"query_enforces_state"`
	ExportPreviewPublishes    bool   `json:"export_preview_publishes"`
	PortalMutationAttempted   bool   `json:"portal_mutation_attempted"`
	CustomerMutationAttempted bool   `json:"customer_mutation_attempted"`
	AuditorMutationAttempted  bool   `json:"auditor_mutation_attempted"`
	ConnectorAuthorityGranted bool   `json:"connector_authority_granted"`
	SchedulerAuthorityGranted bool   `json:"scheduler_authority_granted"`
	AgentAuthorityGranted     bool   `json:"agent_authority_granted"`
	CanonicalMutationAllowed  bool   `json:"canonical_mutation_allowed"`
	ProductionMutationAllowed bool   `json:"production_mutation_allowed"`
	PassAllowed               bool   `json:"pass_allowed"`
}

type Point15ValDNoOverclaimGuard struct {
	ObservedTexts                        []string `json:"observed_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ProjectionDisclaimer                 string   `json:"projection_disclaimer"`
}

func point15ValDStates() []string {
	return []string{
		Point15ValDStateActive,
		Point15ValDStateBlocked,
		Point15ValDStateReviewRequired,
		Point15ValDStateIncomplete,
	}
}

func point15ValDStateValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValDStates())
}

func point15ValDModes() []string {
	return []string{
		point15ValDModeTimeline,
		point15ValDModeDashboardSummary,
		point15ValDModeEvidenceDetail,
		point15ValDModeRevalidationDetail,
		point15ValDModeEnforcementDetail,
		point15ValDModeReplayDetail,
		point15ValDModeExportPreview,
		point15ValDModeQueryResult,
	}
}

func point15ValDRawExactValueValid(value string, allowedValues []string) bool {
	return formalRawExactValid(value, func(candidate string) bool {
		for _, allowed := range allowedValues {
			if candidate == allowed {
				return true
			}
		}
		return false
	})
}

func point15ValDRawVal0StateValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15Val0States())
}

func point15ValDRawFreshnessStatusValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15Val0FreshnessStatuses())
}

func point15ValDRawDowngradeOutcomeValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15Val0DowngradeOutcomes())
}

func point15ValDRawScheduleStatusValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValBScheduleStatuses())
}

func point15ValDRawRunResultValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValBRunResults())
}

func point15ValDRawRetryStatusValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValBRetryStatuses())
}

func point15ValDRawThrottleStatusValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValBThrottleStatuses())
}

func point15ValDRawActionValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValCActions())
}

func point15ValDRawReasonValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValCReasons())
}

func point15ValDRawLifecycleStatusValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValCLifecycleStatuses())
}

func point15ValDRawTriggerValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValATriggers())
}

func point15ValDModeValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValDModes())
}

func point15ValDVisibilities() []string {
	return []string{
		point15ValDVisibilityInternalOnly,
		point15ValDVisibilityTenantScoped,
		point15ValDVisibilityAuditorRead,
		point15ValDVisibilityCustomerRead,
		point15ValDVisibilityPublicBlocked,
	}
}

func point15ValDVisibilityValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValDVisibilities())
}

func point15ValDActions() []string {
	return []string{
		point15ValDActionDisplayOnly,
		point15ValDActionFilterOnly,
		point15ValDActionSortOnly,
		point15ValDActionExplainOnly,
		point15ValDActionExportPreviewOnly,
	}
}

func point15ValDActionValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValDActions())
}

func point15ValDQueryResultStates() []string {
	return []string{
		point15ValDQueryComplete,
		point15ValDQueryPartial,
		point15ValDQueryRedacted,
		point15ValDQueryBlocked,
		point15ValDQueryReviewRequired,
		point15ValDQueryIncomplete,
	}
}

func point15ValDQueryResultStateValid(value string) bool {
	return formalRawExactValid(value, func(candidate string) bool {
		for _, allowed := range point15ValDQueryResultStates() {
			if candidate == allowed {
				return true
			}
		}
		return false
	})
}

func point15ValDEventTypes() []string {
	return []string{
		point15ValDEventFreshness,
		point15ValDEventRevalidation,
		point15ValDEventEnforcement,
		point15ValDEventRevocation,
		point15ValDEventExpiry,
		point15ValDEventSupersession,
		point15ValDEventReplay,
	}
}

func point15ValDEventTypeValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValDEventTypes())
}

func point15ValDQueryScopes() []string {
	return []string{
		point15ValDQueryScopeHistory,
		point15ValDQueryScopeEvidence,
		point15ValDQueryScopeRevalidation,
		point15ValDQueryScopeEnforcement,
	}
}

func point15ValDQueryScopeValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValDQueryScopes())
}

func point15ValDSortOrders() []string {
	return []string{
		point15ValDSortEventDesc,
		point15ValDSortEventAsc,
		point15ValDSortStateSeverity,
		point15ValDSortEvidenceID,
	}
}

func point15ValDSortOrderValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValDSortOrders())
}

func point15ValDQueryFilters() []string {
	return []string{
		point15ValDFilterTenantScope,
		point15ValDFilterEvidenceID,
		point15ValDFilterCurrentState,
		point15ValDFilterDowngradeReason,
		point15ValDFilterEnforcementReason,
		point15ValDFilterFreshnessStatus,
		point15ValDFilterEventType,
	}
}

func point15ValDFilterValid(value string) bool {
	return point15ValDRawExactValueValid(value, point15ValDQueryFilters())
}

func point15ValDRedactionStates() []string {
	return []string{
		point15ValDRedactionNone,
		point15ValDRedactionLimited,
		point15ValDRedactionTenantPrivate,
		point15ValDRedactionBlockedPrivate,
	}
}

func point15ValDRedactionStateValid(value string) bool {
	return formalRawExactValid(value, func(candidate string) bool {
		for _, allowed := range point15ValDRedactionStates() {
			if candidate == allowed {
				return true
			}
		}
		return false
	})
}

func point15ValDRefValid(value string) bool {
	return formalRawExactTokenValid(value, func(candidate string) bool {
		return point14Val0RefValid(candidate,
			"point15_vald_",
			"point15_valc_",
			"point15_valb_",
			"point15_vala_",
			"point15_val0_",
			"timeline_",
			"dashboard_",
			"query_",
			"evidence_",
			"schedule_",
			"run_",
			"budget_",
			"retry_budget_",
			"throttle_",
			"binding_",
			"decision_",
			"lineage_",
			"history_",
			"replay_",
			"proof_",
			"proof_pack_",
			"enforcement_",
			"revocation_",
			"revocation_source_",
			"expiry_",
			"supersession_",
			"authority_boundary_",
			"access_",
			"tenant_",
			"timestamp_",
			"no_mutation_",
		)
	})
}

func point15ValDRawScopeValid(value string) bool {
	return formalRawExactValid(value, point11Val0ScopeValid)
}

func point15ValDOptionalRawTimestampValid(value string) bool {
	return value == "" || point12Val0RawTimestampValid(value)
}

func point15ValDForbiddenWording() []string {
	return append([]string{
		"continuous assurance guaranteed",
		"automatically verified forever",
		"always fresh",
		"production approved",
		"deployment approved",
		"compliance guaranteed",
		"legal proof",
		"financial guarantee",
		"official authority",
		"global truth",
		"public badge",
		"guaranteed secure",
		"certified secure",
	}, inheritedDeploymentReadinessOverclaimClaims()...)
}

func point15ValDSafeWording() []string {
	return []string{
		"read-only assurance projection",
		"display-only timeline",
		"projection sourced from formal enforcement result",
		"tenant-scoped read-only evidence view",
		"limitations remain visible",
		"bounded assurance history",
	}
}

func point15ValDObservedTextContainsForbiddenWording(text string) bool {
	return point15ValDObservedTextContainsForbiddenWordingWithNormalizer(text, point15ValDNormalizedObservedText)
}

func point15ValDInternalTextContainsForbiddenWording(text string) bool {
	return point15ValDObservedTextContainsForbiddenWordingWithNormalizer(text, point15ValDNormalizedInternalText)
}

func point15ValDObservedTextContainsForbiddenWordingWithNormalizer(text string, normalize func(string) string) bool {
	normalized := normalize(text)
	if normalized == "" {
		return false
	}
	if point15ValDObservedTextAllowedSafeWithNormalizer(normalized, normalize) {
		return false
	}
	for _, forbidden := range point15ValDForbiddenWording() {
		if formalNoOverclaimContainsForbidden(normalized, normalize(forbidden)) {
			return true
		}
	}
	return false
}

func point15ValDNormalizedObservedText(text string) string {
	return formalNoOverclaimNormalizePublicText(text)
}

func point15ValDNormalizedInternalText(text string) string {
	return formalNoOverclaimNormalizeText(text)
}

func point15ValDObservedTextAllowedSafe(normalized string) bool {
	return point15ValDObservedTextAllowedSafeWithNormalizer(normalized, point15ValDNormalizedObservedText)
}

func point15ValDObservedTextAllowedSafeWithNormalizer(normalized string, normalize func(string) string) bool {
	for _, safe := range point15ValDSafeWording() {
		if normalized == normalize(safe) {
			return true
		}
	}
	return false
}

func point15ValDObservedListContainsForbiddenWording(values []string) bool {
	return point15ValDListContainsForbiddenWording(values, point15ValDNormalizedObservedText)
}

func point15ValDInternalListContainsForbiddenWording(values []string) bool {
	return point15ValDListContainsForbiddenWording(values, point15ValDNormalizedInternalText)
}

func point15ValDListContainsForbiddenWording(values []string, normalize func(string) string) bool {
	return point15ObservedListContainsForbiddenWordingWithNormalizer(values, point15ValDSafeWording(), point15ValDForbiddenWording(), normalize)
}

func point15ValDValCPayloadContainsPoint15Pass(valC Point15ValCEnforcementBoundaryFoundation) bool {
	payload, err := json.Marshal(valC)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point15ValDBlockedPassToken)
}

func point15ValDQueryResultWaveState(state string) string {
	switch state {
	case point15ValDQueryBlocked:
		return Point15ValDStateBlocked
	case point15ValDQueryReviewRequired:
		return Point15ValDStateReviewRequired
	case point15ValDQueryIncomplete:
		return Point15ValDStateIncomplete
	default:
		return Point15ValDStateActive
	}
}

func point15ValDDependencyModelFromUpstream(valC Point15ValCEnforcementBoundaryFoundation) Point15ValDDependencySnapshot {
	return Point15ValDDependencySnapshot{
		Point15ValCCurrentState:           valC.CurrentState,
		Point15ValCDependencyState:        valC.DependencyState,
		Point15ValCEnforcementActionState: valC.EnforcementActionState,
		Point15ValCEvidenceLifecycleState: valC.EvidenceLifecycleState,
		Point15ValCRevocationState:        valC.RevocationBoundaryState,
		Point15ValCExpiryState:            valC.ExpiryBoundaryState,
		Point15ValCSupersessionState:      valC.SupersessionState,
		Point15ValCReplayHistoryState:     valC.ReplayProofHistoryState,
		Point15ValCTimestampState:         valC.TimestampDisciplineState,
		Point15ValCAuthorityState:         valC.AuthorityBoundaryState,
		Point15ValCTenantState:            valC.TenantBoundaryState,
		Point15ValCNoOverclaimState:       valC.NoOverclaimState,
		Point15ValCComputedFromUpstream:   valC.Dependency.SnapshotFromComputedOutput,
		Point15ValCMerged:                 true,
		Point15ValCCIGreen:                true,
		Point15ValCReviewedOnMain:         true,
		Point15PassSeen:                   point15ValDValCPayloadContainsPoint15Pass(valC),
		InheritedPoint15ValBCurrentState:  valC.Dependency.Point15ValB.CurrentState,
		InheritedPoint15ValACurrentState:  valC.Dependency.Point15ValB.Dependency.Point15ValA.CurrentState,
		InheritedPoint15Val0CurrentState:  valC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.CurrentState,
		InheritedPoint14ValECurrentState:  valC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.Dependency.Point14ValE.CurrentState,
		InheritedTenantScope:              valC.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:        true,
		ReviewPrerequisites:               []string{},
		Point15ValC:                       valC,
	}
}

func point15ValDDependencyModel() Point15ValDDependencySnapshot {
	valC := ComputePoint15ValCEnforcementBoundaryFoundation(Point15ValCFoundationModel())
	return point15ValDDependencyModelFromUpstream(valC)
}

func point15ValCEmbeddedNoOverclaimChainActive(valC Point15ValCEnforcementBoundaryFoundation) bool {
	valB := valC.Dependency.Point15ValB
	valA := valB.Dependency.Point15ValA
	val0 := valA.Dependency.Point15Val0
	if val0.FreshnessDisclaimer != point15Val0FreshnessDisclaimer ||
		valA.TriggerDisclaimer != point15ValATriggerDisclaimer ||
		valB.RevalidationDisclaimer != point15ValBRevalidationDisclaimer ||
		valC.EnforcementDisclaimer != point15ValCEnforcementDisclaimer {
		return false
	}
	if val0.NoOverclaimState != Point15Val0StateActive ||
		EvaluatePoint15Val0NoOverclaimGuardState(val0.NoOverclaimGuard) != Point15Val0StateActive {
		return false
	}
	if valA.NoOverclaimState != Point15ValAStateActive ||
		EvaluatePoint15ValANoOverclaimGuardState(valA.NoOverclaimGuard) != Point15ValAStateActive {
		return false
	}
	if valB.NoOverclaimState != Point15ValBStateActive ||
		EvaluatePoint15ValBNoOverclaimGuardState(valB.NoOverclaimGuard) != Point15ValBStateActive {
		return false
	}
	if valC.NoOverclaimState != Point15ValCStateActive ||
		EvaluatePoint15ValCNoOverclaimGuardState(valC.NoOverclaimGuard) != Point15ValCStateActive {
		return false
	}
	return true
}

func point15Val0EmbeddedStateChainActive(val0 Point15Val0FreshnessDisciplineFoundation) bool {
	dependencyState := Point15Val0StateBlocked
	if point15Val0EmbeddedDependencySnapshotActive(val0.Dependency) {
		dependencyState = Point15Val0StateActive
	}
	freshnessTaxonomyState := EvaluatePoint15Val0EvidenceFreshnessTaxonomyState(val0.FreshnessTaxonomy)
	downgradeTaxonomyState := EvaluatePoint15Val0DowngradeTaxonomyState(val0.DowngradeTaxonomy)
	evidenceContextState := EvaluatePoint15Val0FreshnessEvidenceContextState(val0.EvidenceContext)
	tenantBoundaryState := EvaluatePoint15Val0TenantBoundaryState(val0.EvidenceContext)
	timestampDisciplineState := EvaluatePoint15Val0TimestampDisciplineState(val0.TimestampDiscipline)
	authorityBoundaryState := EvaluatePoint15Val0AuthorityBoundaryState(val0.AuthorityBoundary)
	noOverclaimState := EvaluatePoint15Val0NoOverclaimGuardState(val0.NoOverclaimGuard)
	if val0.FreshnessDisclaimer != point15Val0FreshnessDisclaimer {
		noOverclaimState = Point15Val0StateBlocked
	}

	expectedTenantScope := val0.Dependency.InheritedTenantScope
	if val0.EvidenceContext.FreshnessStatus != val0.FreshnessTaxonomy.FreshnessStatus {
		evidenceContextState = Point15Val0StateBlocked
	}
	if val0.DowngradeTaxonomy.FreshnessStatus != val0.FreshnessTaxonomy.FreshnessStatus {
		downgradeTaxonomyState = Point15Val0StateBlocked
	}
	if val0.TimestampDiscipline.FreshnessStatus != val0.FreshnessTaxonomy.FreshnessStatus {
		timestampDisciplineState = Point15Val0StateBlocked
	}
	if val0.EvidenceContext.DowngradeOutcome != val0.DowngradeTaxonomy.DowngradeOutcome {
		evidenceContextState = Point15Val0StateBlocked
	}
	if expectedTenantScope == "" ||
		val0.EvidenceContext.TenantScope != expectedTenantScope ||
		(val0.EvidenceContext.ReferencedTenantScope != "" && val0.EvidenceContext.ReferencedTenantScope != expectedTenantScope) {
		evidenceContextState = Point15Val0StateBlocked
		tenantBoundaryState = Point15Val0StateBlocked
	}
	if expectedTenantScope == "" || val0.TimestampDiscipline.TenantScope != expectedTenantScope {
		timestampDisciplineState = Point15Val0StateBlocked
	}
	if expectedTenantScope == "" || val0.AuthorityBoundary.TenantScope != expectedTenantScope {
		authorityBoundaryState = Point15Val0StateBlocked
	}

	return val0.CurrentState == Point15Val0StateActive &&
		val0.DependencyState == dependencyState &&
		val0.FreshnessTaxonomyState == freshnessTaxonomyState &&
		val0.DowngradeTaxonomyState == downgradeTaxonomyState &&
		val0.EvidenceContextState == evidenceContextState &&
		val0.TenantBoundaryState == tenantBoundaryState &&
		val0.TimestampDisciplineState == timestampDisciplineState &&
		val0.AuthorityBoundaryState == authorityBoundaryState &&
		val0.NoOverclaimState == noOverclaimState &&
		dependencyState == Point15Val0StateActive &&
		freshnessTaxonomyState == Point15Val0StateActive &&
		downgradeTaxonomyState == Point15Val0StateActive &&
		evidenceContextState == Point15Val0StateActive &&
		tenantBoundaryState == Point15Val0StateActive &&
		timestampDisciplineState == Point15Val0StateActive &&
		authorityBoundaryState == Point15Val0StateActive &&
		noOverclaimState == Point15Val0StateActive
}

func point15ValAEmbeddedStateChainActive(valA Point15ValADowngradeTriggerFoundation) bool {
	recomputed := ComputePoint15ValADowngradeTriggerFoundation(valA)
	return point15Val0EmbeddedStateChainActive(valA.Dependency.Point15Val0) &&
		valA.CurrentState == recomputed.CurrentState &&
		valA.DependencyState == recomputed.DependencyState &&
		valA.TriggerTableState == recomputed.TriggerTableState &&
		valA.TriggerState == recomputed.TriggerState &&
		valA.ReasonState == recomputed.ReasonState &&
		valA.DecisionState == recomputed.DecisionState &&
		valA.AuthorityBoundaryState == recomputed.AuthorityBoundaryState &&
		valA.NoOverclaimState == recomputed.NoOverclaimState &&
		recomputed.CurrentState == Point15ValAStateActive &&
		recomputed.DependencyState == Point15ValAStateActive &&
		recomputed.TriggerTableState == Point15ValAStateActive &&
		recomputed.TriggerState == Point15ValAStateActive &&
		recomputed.ReasonState == Point15ValAStateActive &&
		recomputed.DecisionState == Point15ValAStateActive &&
		recomputed.AuthorityBoundaryState == Point15ValAStateActive &&
		recomputed.NoOverclaimState == Point15ValAStateActive
}

func point15ValBEmbeddedStateChainActive(valB Point15ValBScheduledRevalidationFoundation) bool {
	recomputed := ComputePoint15ValBScheduledRevalidationFoundation(valB)
	return point15ValAEmbeddedStateChainActive(valB.Dependency.Point15ValA) &&
		valB.CurrentState == recomputed.CurrentState &&
		valB.DependencyState == recomputed.DependencyState &&
		valB.ScheduleState == recomputed.ScheduleState &&
		valB.RunState == recomputed.RunState &&
		valB.RetryBudgetState == recomputed.RetryBudgetState &&
		valB.TenantThrottleState == recomputed.TenantThrottleState &&
		valB.DowngradeBindingState == recomputed.DowngradeBindingState &&
		valB.TimestampDisciplineState == recomputed.TimestampDisciplineState &&
		valB.AuthorityBoundaryState == recomputed.AuthorityBoundaryState &&
		valB.NoOverclaimState == recomputed.NoOverclaimState &&
		recomputed.CurrentState == Point15ValBStateActive &&
		recomputed.DependencyState == Point15ValBStateActive &&
		recomputed.ScheduleState == Point15ValBStateActive &&
		recomputed.RunState == Point15ValBStateActive &&
		recomputed.RetryBudgetState == Point15ValBStateActive &&
		recomputed.TenantThrottleState == Point15ValBStateActive &&
		recomputed.DowngradeBindingState == Point15ValBStateActive &&
		recomputed.TimestampDisciplineState == Point15ValBStateActive &&
		recomputed.AuthorityBoundaryState == Point15ValBStateActive &&
		recomputed.NoOverclaimState == Point15ValBStateActive
}

func point15ValCEmbeddedStateChainActive(valC Point15ValCEnforcementBoundaryFoundation) bool {
	recomputed := ComputePoint15ValCEnforcementBoundaryFoundation(valC)
	return point15ValBEmbeddedStateChainActive(valC.Dependency.Point15ValB) &&
		valC.CurrentState == recomputed.CurrentState &&
		valC.DependencyState == recomputed.DependencyState &&
		valC.EnforcementActionState == recomputed.EnforcementActionState &&
		valC.EvidenceLifecycleState == recomputed.EvidenceLifecycleState &&
		valC.RevocationBoundaryState == recomputed.RevocationBoundaryState &&
		valC.ExpiryBoundaryState == recomputed.ExpiryBoundaryState &&
		valC.SupersessionState == recomputed.SupersessionState &&
		valC.ReplayProofHistoryState == recomputed.ReplayProofHistoryState &&
		valC.TimestampDisciplineState == recomputed.TimestampDisciplineState &&
		valC.AuthorityBoundaryState == recomputed.AuthorityBoundaryState &&
		valC.TenantBoundaryState == recomputed.TenantBoundaryState &&
		valC.NoOverclaimState == recomputed.NoOverclaimState &&
		recomputed.CurrentState == Point15ValCStateActive &&
		recomputed.DependencyState == Point15ValCStateActive &&
		recomputed.EnforcementActionState == Point15ValCStateActive &&
		recomputed.EvidenceLifecycleState == Point15ValCStateActive &&
		recomputed.RevocationBoundaryState == Point15ValCStateActive &&
		recomputed.ExpiryBoundaryState == Point15ValCStateActive &&
		recomputed.SupersessionState == Point15ValCStateActive &&
		recomputed.ReplayProofHistoryState == Point15ValCStateActive &&
		recomputed.TimestampDisciplineState == Point15ValCStateActive &&
		recomputed.AuthorityBoundaryState == Point15ValCStateActive &&
		recomputed.TenantBoundaryState == Point15ValCStateActive &&
		recomputed.NoOverclaimState == Point15ValCStateActive &&
		point15ValCEmbeddedNoOverclaimChainActive(recomputed)
}

func point15ValDEmbeddedStateChainActive(valD Point15ValDAssuranceProjectionFoundation) bool {
	recomputed := ComputePoint15ValDAssuranceProjectionFoundation(valD)
	return valD.ProjectionDisclaimer == point15ValDProjectionDisclaimer &&
		point15ValCEmbeddedStateChainActive(valD.Dependency.Point15ValC) &&
		valD.CurrentState == recomputed.CurrentState &&
		valD.DependencyState == recomputed.DependencyState &&
		valD.TimelineState == recomputed.TimelineState &&
		valD.DashboardState == recomputed.DashboardState &&
		valD.QueryState == recomputed.QueryState &&
		valD.EvidenceDetailState == recomputed.EvidenceDetailState &&
		valD.RevalidationDetailState == recomputed.RevalidationDetailState &&
		valD.EnforcementDetailState == recomputed.EnforcementDetailState &&
		valD.ReplayProofHistoryState == recomputed.ReplayProofHistoryState &&
		valD.AccessTenantState == recomputed.AccessTenantState &&
		valD.TimestampDisplayState == recomputed.TimestampDisplayState &&
		valD.NoMutationState == recomputed.NoMutationState &&
		valD.AuthorityBoundaryState == recomputed.AuthorityBoundaryState &&
		valD.NoOverclaimState == recomputed.NoOverclaimState &&
		recomputed.CurrentState == Point15ValDStateActive &&
		recomputed.DependencyState == Point15ValDStateActive &&
		recomputed.TimelineState == Point15ValDStateActive &&
		recomputed.DashboardState == Point15ValDStateActive &&
		recomputed.QueryState == Point15ValDStateActive &&
		recomputed.EvidenceDetailState == Point15ValDStateActive &&
		recomputed.RevalidationDetailState == Point15ValDStateActive &&
		recomputed.EnforcementDetailState == Point15ValDStateActive &&
		recomputed.ReplayProofHistoryState == Point15ValDStateActive &&
		recomputed.AccessTenantState == Point15ValDStateActive &&
		recomputed.TimestampDisplayState == Point15ValDStateActive &&
		recomputed.NoMutationState == Point15ValDStateActive &&
		recomputed.AuthorityBoundaryState == Point15ValDStateActive &&
		recomputed.NoOverclaimState == Point15ValDStateActive
}

func point15ValDEmbeddedNoOverclaimChainActive(valD Point15ValDAssuranceProjectionFoundation) bool {
	return valD.ProjectionDisclaimer == point15ValDProjectionDisclaimer &&
		valD.NoOverclaimState == Point15ValDStateActive &&
		EvaluatePoint15ValDNoOverclaimGuardState(valD.NoOverclaimGuard) == Point15ValDStateActive &&
		point15ValCEmbeddedNoOverclaimChainActive(valD.Dependency.Point15ValC)
}

func EvaluatePoint15ValDDependencyState(model Point15ValDDependencySnapshot) string {
	if model.Point15ValCCurrentState != Point15ValCStateActive ||
		model.Point15ValCDependencyState != Point15ValCStateActive ||
		model.Point15ValCEnforcementActionState != Point15ValCStateActive ||
		model.Point15ValCEvidenceLifecycleState != Point15ValCStateActive ||
		model.Point15ValCRevocationState != Point15ValCStateActive ||
		model.Point15ValCExpiryState != Point15ValCStateActive ||
		model.Point15ValCSupersessionState != Point15ValCStateActive ||
		model.Point15ValCReplayHistoryState != Point15ValCStateActive ||
		model.Point15ValCTimestampState != Point15ValCStateActive ||
		model.Point15ValCAuthorityState != Point15ValCStateActive ||
		model.Point15ValCTenantState != Point15ValCStateActive ||
		model.Point15ValCNoOverclaimState != Point15ValCStateActive {
		return Point15ValDStateBlocked
	}
	if model.Point15ValCCurrentState != model.Point15ValC.CurrentState ||
		model.Point15ValCDependencyState != model.Point15ValC.DependencyState ||
		model.Point15ValCEnforcementActionState != model.Point15ValC.EnforcementActionState ||
		model.Point15ValCEvidenceLifecycleState != model.Point15ValC.EvidenceLifecycleState ||
		model.Point15ValCRevocationState != model.Point15ValC.RevocationBoundaryState ||
		model.Point15ValCExpiryState != model.Point15ValC.ExpiryBoundaryState ||
		model.Point15ValCSupersessionState != model.Point15ValC.SupersessionState ||
		model.Point15ValCReplayHistoryState != model.Point15ValC.ReplayProofHistoryState ||
		model.Point15ValCTimestampState != model.Point15ValC.TimestampDisciplineState ||
		model.Point15ValCAuthorityState != model.Point15ValC.AuthorityBoundaryState ||
		model.Point15ValCTenantState != model.Point15ValC.TenantBoundaryState ||
		model.Point15ValCNoOverclaimState != model.Point15ValC.NoOverclaimState ||
		model.Point15ValCComputedFromUpstream != model.Point15ValC.Dependency.SnapshotFromComputedOutput ||
		model.InheritedPoint15ValBCurrentState != model.Point15ValC.Dependency.Point15ValB.CurrentState ||
		model.InheritedPoint15ValACurrentState != model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.CurrentState ||
		model.InheritedPoint15Val0CurrentState != model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.CurrentState ||
		model.InheritedPoint14ValECurrentState != model.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.Dependency.Point14ValE.CurrentState ||
		model.InheritedTenantScope != model.Point15ValC.Dependency.InheritedTenantScope {
		return Point15ValDStateBlocked
	}
	if !model.Point15ValCComputedFromUpstream || !model.Point15ValCMerged || !model.Point15ValCCIGreen || !model.Point15ValCReviewedOnMain || !model.SnapshotFromComputedOutput {
		return Point15ValDStateBlocked
	}
	if model.Point15PassSeen || point15ValDValCPayloadContainsPoint15Pass(model.Point15ValC) {
		return Point15ValDStateBlocked
	}
	if !point15ValCEmbeddedStateChainActive(model.Point15ValC) {
		return Point15ValDStateBlocked
	}
	if model.InheritedPoint15ValBCurrentState != Point15ValBStateActive ||
		model.InheritedPoint15ValACurrentState != Point15ValAStateActive ||
		model.InheritedPoint15Val0CurrentState != Point15Val0StateActive ||
		model.InheritedPoint14ValECurrentState != Point14ValEStatePassConfirmed {
		return Point15ValDStateBlocked
	}
	if !formalRawExactValid(model.InheritedTenantScope, point11Val0ScopeValid) {
		return Point15ValDStateBlocked
	}
	return Point15ValDStateActive
}

func point15ValDAssuranceTimelineModel(dependency Point15ValDDependencySnapshot) Point15ValDAssuranceTimelineEntry {
	eventAt := dependency.Point15ValC.TimestampDiscipline.EvaluatedAt
	if eventAt == "" {
		eventAt = dependency.Point15ValC.TimestampDiscipline.ReferenceNow
	}
	return Point15ValDAssuranceTimelineEntry{
		TimelineID:              "timeline_point15_vald_001",
		EntryID:                 "timeline_entry_point15_vald_001",
		ProjectionMode:          point15ValDModeTimeline,
		ProjectionAction:        point15ValDActionDisplayOnly,
		Visibility:              point15ValDVisibilityTenantScoped,
		TenantScope:             dependency.InheritedTenantScope,
		EvidenceID:              dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.EvidenceID,
		EventType:               point15ValDEventEnforcement,
		PriorState:              Point15Val0StateActive,
		CurrentState:            dependency.Point15ValC.EnforcementAction.TargetState,
		DowngradeReason:         dependency.Point15ValC.Dependency.Point15ValB.DowngradeBinding.TriggerType,
		EnforcementReason:       dependency.Point15ValC.EnforcementAction.EnforcementReason,
		SourceValCRef:           dependency.Point15ValC.EnforcementAction.EnforcementID,
		ReplayRef:               dependency.Point15ValC.ReplayProofHistory.ReplayRef,
		ProofHistoryRef:         dependency.Point15ValC.ReplayProofHistory.ProofHistoryRef,
		EventAt:                 eventAt,
		DisplayedAt:             dependency.Point15ValC.TimestampDiscipline.ReferenceNow,
		TimeSource:              dependency.Point15ValC.TimestampDiscipline.ReferenceNowTimeSource,
		DecisiveEvidenceVisible: true,
		BlockedReasonVisible:    true,
		PriorStateVisible:       true,
		CurrentStateVisible:     true,
		PriorPassVisible:        true,
		LaterDowngradeVisible:   true,
		TimelineCreatesValidity: false,
	}
}

func EvaluatePoint15ValDAssuranceTimelineState(model Point15ValDAssuranceTimelineEntry) string {
	if !point15ValDRefValid(model.TimelineID) ||
		!point15ValDRefValid(model.EntryID) ||
		!point15ValDModeValid(model.ProjectionMode) ||
		!point15ValDActionValid(model.ProjectionAction) ||
		!point15ValDVisibilityValid(model.Visibility) ||
		!point15ValDRawScopeValid(model.TenantScope) ||
		!formalRawExactNonEmpty(model.EvidenceID) ||
		!point15ValDEventTypeValid(model.EventType) ||
		!point15ValDRawVal0StateValid(model.PriorState) ||
		!point15ValDRawVal0StateValid(model.CurrentState) ||
		!point15ValDRefValid(model.SourceValCRef) ||
		!point15ValDRefValid(model.ReplayRef) ||
		!point15ValDRefValid(model.ProofHistoryRef) ||
		!point12Val0RawTimestampValid(model.EventAt) ||
		!point12Val0RawTimestampValid(model.DisplayedAt) ||
		!point14Val0CanonicalTimeSourceValid(model.TimeSource) {
		return Point15ValDStateBlocked
	}
	if model.ProjectionMode != point15ValDModeTimeline ||
		(model.ProjectionAction != point15ValDActionDisplayOnly && model.ProjectionAction != point15ValDActionExplainOnly) {
		return Point15ValDStateBlocked
	}
	displayedAt, _ := point14Val0ParsedTime(model.DisplayedAt)
	eventAt, _ := point14Val0ParsedTime(model.EventAt)
	if displayedAt.Before(eventAt) {
		return Point15ValDStateBlocked
	}
	if model.TimelineCreatesValidity {
		return Point15ValDStateBlocked
	}
	if !model.PriorStateVisible || !model.CurrentStateVisible || !model.PriorPassVisible || !model.LaterDowngradeVisible {
		return Point15ValDStateBlocked
	}
	if !model.DecisiveEvidenceVisible {
		if model.CurrentState == Point15Val0StateReviewRequired {
			return Point15ValDStateReviewRequired
		}
		return Point15ValDStateBlocked
	}
	if model.CurrentState == Point15Val0StateBlocked && !model.BlockedReasonVisible {
		return Point15ValDStateBlocked
	}
	if model.CurrentState != Point15Val0StateActive &&
		model.DowngradeReason == "" &&
		model.EnforcementReason == "" {
		return Point15ValDStateBlocked
	}
	if model.DowngradeReason != "" && !point15ValDRawTriggerValid(model.DowngradeReason) {
		return Point15ValDStateBlocked
	}
	if model.EnforcementReason != "" && !point15ValDRawReasonValid(model.EnforcementReason) {
		return Point15ValDStateBlocked
	}
	return Point15ValDStateActive
}

func point15ValDDashboardSummaryModel(dependency Point15ValDDependencySnapshot) Point15ValDDashboardSummary {
	return Point15ValDDashboardSummary{
		DashboardID:                   "dashboard_point15_vald_001",
		ProjectionMode:                point15ValDModeDashboardSummary,
		ProjectionAction:              point15ValDActionDisplayOnly,
		Visibility:                    point15ValDVisibilityTenantScoped,
		TenantScope:                   dependency.InheritedTenantScope,
		SummaryScope:                  point15ValDSummaryScopeTenant,
		EvidenceCount:                 1,
		ActiveCount:                   1,
		BlockedCount:                  0,
		ReviewRequiredCount:           0,
		IncompleteCount:               0,
		StaleCount:                    0,
		ExpiredCount:                  0,
		RevokedCount:                  0,
		EnforcementCount:              0,
		HiddenBlockedCount:            false,
		HiddenReviewRequiredCount:     false,
		HiddenIncompleteCount:         false,
		RestoresActiveClosure:         false,
		ActiveCountIncludesDisallowed: false,
	}
}

func EvaluatePoint15ValDDashboardSummaryState(model Point15ValDDashboardSummary) string {
	if !point15ValDRefValid(model.DashboardID) ||
		model.ProjectionMode != point15ValDModeDashboardSummary ||
		model.ProjectionAction != point15ValDActionDisplayOnly ||
		!point15ValDVisibilityValid(model.Visibility) ||
		model.SummaryScope != point15ValDSummaryScopeTenant {
		return Point15ValDStateBlocked
	}
	if model.TenantScope == "" {
		return Point15ValDStateIncomplete
	}
	if !point15ValDRawScopeValid(model.TenantScope) {
		return Point15ValDStateBlocked
	}
	if model.EvidenceCount < 0 || model.ActiveCount < 0 || model.BlockedCount < 0 || model.ReviewRequiredCount < 0 || model.IncompleteCount < 0 || model.StaleCount < 0 || model.ExpiredCount < 0 || model.RevokedCount < 0 || model.EnforcementCount < 0 {
		return Point15ValDStateBlocked
	}
	if model.ActiveCount+model.BlockedCount+model.ReviewRequiredCount+model.IncompleteCount != model.EvidenceCount {
		return Point15ValDStateBlocked
	}
	if model.HiddenBlockedCount || model.HiddenReviewRequiredCount || model.HiddenIncompleteCount || model.RestoresActiveClosure || model.ActiveCountIncludesDisallowed {
		return Point15ValDStateBlocked
	}
	return Point15ValDStateActive
}

func point15ValDQueryProjectionModel(dependency Point15ValDDependencySnapshot) Point15ValDQueryProjection {
	evidenceID := dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.EvidenceID
	enforcementRef := dependency.Point15ValC.EnforcementAction.EnforcementID
	replayRef := dependency.Point15ValC.ReplayProofHistory.ReplayRef
	return Point15ValDQueryProjection{
		QueryID:                 "query_point15_vald_001",
		ProjectionMode:          point15ValDModeQueryResult,
		ProjectionAction:        point15ValDActionFilterOnly,
		Visibility:              point15ValDVisibilityTenantScoped,
		TenantScope:             dependency.InheritedTenantScope,
		ViewerScope:             dependency.InheritedTenantScope,
		RequestedScope:          point15ValDQueryScopeHistory,
		Filters:                 []string{point15ValDFilterTenantScope, point15ValDFilterCurrentState},
		SortOrder:               point15ValDSortEventDesc,
		ResultState:             point15ValDQueryComplete,
		ResultRefs:              []string{evidenceID, enforcementRef, replayRef},
		RedactionState:          point15ValDRedactionNone,
		DecisiveEvidenceVisible: true,
		LimitationsVisible:      true,
		QueryMutationAttempted:  false,
		StrengthensClaims:       false,
		CrossTenantQuery:        false,
	}
}

func EvaluatePoint15ValDQueryProjectionState(model Point15ValDQueryProjection) string {
	if !point15ValDRefValid(model.QueryID) ||
		model.ProjectionMode != point15ValDModeQueryResult ||
		!point15ValDActionValid(model.ProjectionAction) ||
		!point15ValDVisibilityValid(model.Visibility) ||
		!point15ValDRawScopeValid(model.TenantScope) ||
		!point15ValDRawScopeValid(model.ViewerScope) ||
		!point15ValDQueryScopeValid(model.RequestedScope) ||
		!point15ValDSortOrderValid(model.SortOrder) ||
		!point15ValDQueryResultStateValid(model.ResultState) ||
		!point15ValDRedactionStateValid(model.RedactionState) {
		return Point15ValDStateBlocked
	}
	if model.ProjectionAction != point15ValDActionFilterOnly &&
		model.ProjectionAction != point15ValDActionSortOnly &&
		model.ProjectionAction != point15ValDActionExplainOnly {
		return Point15ValDStateBlocked
	}
	for _, filter := range model.Filters {
		if !point15ValDFilterValid(filter) {
			return Point15ValDStateBlocked
		}
	}
	for _, ref := range model.ResultRefs {
		if !formalRawExactNonEmpty(ref) || !point15ValDRefValid(ref) {
			return Point15ValDStateBlocked
		}
	}
	if model.QueryMutationAttempted || model.StrengthensClaims || model.CrossTenantQuery || model.Visibility == point15ValDVisibilityPublicBlocked {
		return Point15ValDStateBlocked
	}
	if model.ViewerScope != model.TenantScope {
		return Point15ValDStateBlocked
	}
	if !model.LimitationsVisible && (model.ResultState == point15ValDQueryRedacted || model.RedactionState != point15ValDRedactionNone) {
		return Point15ValDStateReviewRequired
	}
	if !model.DecisiveEvidenceVisible {
		if model.ResultState == point15ValDQueryRedacted || model.ResultState == point15ValDQueryReviewRequired {
			return Point15ValDStateReviewRequired
		}
		return Point15ValDStateBlocked
	}
	return point15ValDQueryResultWaveState(model.ResultState)
}

func point15ValDEvidenceDetailModel(dependency Point15ValDDependencySnapshot) Point15ValDEvidenceDetailProjection {
	val0 := dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0
	return Point15ValDEvidenceDetailProjection{
		ProjectionMode:              point15ValDModeEvidenceDetail,
		ProjectionAction:            point15ValDActionDisplayOnly,
		Visibility:                  point15ValDVisibilityTenantScoped,
		EvidenceID:                  val0.EvidenceContext.EvidenceID,
		TenantScope:                 dependency.InheritedTenantScope,
		EvidenceHash:                val0.EvidenceContext.EvidenceHash,
		PolicyVersion:               val0.EvidenceContext.PolicyVersion,
		EngineVersion:               val0.EvidenceContext.EngineVersion,
		SchemaVersion:               val0.EvidenceContext.SchemaVersion,
		FreshnessStatus:             val0.FreshnessTaxonomy.FreshnessStatus,
		DowngradeOutcome:            val0.DowngradeTaxonomy.DowngradeOutcome,
		LifecycleStatus:             dependency.Point15ValC.EvidenceLifecycle.LifecycleStatus,
		EnforcementStatus:           dependency.Point15ValC.EnforcementAction.TargetState,
		LimitationsVisible:          true,
		IdentityDerivedFromNameOnly: false,
	}
}

func EvaluatePoint15ValDEvidenceDetailProjectionState(model Point15ValDEvidenceDetailProjection) string {
	if model.ProjectionMode != point15ValDModeEvidenceDetail ||
		(model.ProjectionAction != point15ValDActionDisplayOnly && model.ProjectionAction != point15ValDActionExplainOnly) ||
		!point15ValDVisibilityValid(model.Visibility) ||
		!formalRawExactNonEmpty(model.EvidenceID) ||
		!point15ValDRawScopeValid(model.TenantScope) ||
		!formalRawExactNonEmpty(model.EvidenceHash) ||
		!formalRawExactNonEmpty(model.PolicyVersion) ||
		!formalRawExactNonEmpty(model.EngineVersion) ||
		!formalRawExactNonEmpty(model.SchemaVersion) ||
		!point15ValDRawFreshnessStatusValid(model.FreshnessStatus) ||
		!point15ValDRawDowngradeOutcomeValid(model.DowngradeOutcome) ||
		!point15ValDRawLifecycleStatusValid(model.LifecycleStatus) ||
		!point15ValDRawVal0StateValid(model.EnforcementStatus) {
		return Point15ValDStateBlocked
	}
	if model.IdentityDerivedFromNameOnly {
		return Point15ValDStateBlocked
	}
	if !model.LimitationsVisible {
		return Point15ValDStateReviewRequired
	}
	return Point15ValDStateActive
}

func point15ValDRevalidationDetailModel(dependency Point15ValDDependencySnapshot) Point15ValDRevalidationDetailProjection {
	valB := dependency.Point15ValC.Dependency.Point15ValB
	return Point15ValDRevalidationDetailProjection{
		ProjectionMode:            point15ValDModeRevalidationDetail,
		ProjectionAction:          point15ValDActionDisplayOnly,
		Visibility:                point15ValDVisibilityTenantScoped,
		ScheduleRef:               valB.Schedule.ScheduleID,
		RunRef:                    valB.Run.RunID,
		RetryBudgetRef:            valB.RetryBudget.BudgetID,
		TenantThrottleRef:         valB.TenantThrottle.ThrottleID,
		ScheduledStatus:           valB.Schedule.ScheduledStatus,
		RunResult:                 valB.Run.RunResult,
		RetryStatus:               valB.RetryBudget.RetryBudgetStatus,
		ThrottleStatus:            valB.TenantThrottle.ThrottleStatus,
		ScheduleMutationAttempted: false,
		RetryTriggered:            false,
		RetryBudgetResetAttempted: false,
		MarksFresh:                false,
		RestoresActiveClosure:     false,
	}
}

func EvaluatePoint15ValDRevalidationDetailProjectionState(model Point15ValDRevalidationDetailProjection) string {
	if model.ProjectionMode != point15ValDModeRevalidationDetail ||
		(model.ProjectionAction != point15ValDActionDisplayOnly && model.ProjectionAction != point15ValDActionExplainOnly) ||
		!point15ValDVisibilityValid(model.Visibility) ||
		!point15ValDRefValid(model.ScheduleRef) ||
		(model.RunRef != "" && !point15ValDRefValid(model.RunRef)) ||
		!point15ValDRefValid(model.RetryBudgetRef) ||
		!point15ValDRefValid(model.TenantThrottleRef) ||
		!point15ValDRawScheduleStatusValid(model.ScheduledStatus) ||
		!point15ValDRawRunResultValid(model.RunResult) ||
		!point15ValDRawRetryStatusValid(model.RetryStatus) ||
		!point15ValDRawThrottleStatusValid(model.ThrottleStatus) {
		return Point15ValDStateBlocked
	}
	if model.ScheduleMutationAttempted || model.RetryTriggered || model.RetryBudgetResetAttempted || model.MarksFresh || model.RestoresActiveClosure {
		return Point15ValDStateBlocked
	}
	return Point15ValDStateActive
}

func point15ValDEnforcementDetailModel(dependency Point15ValDDependencySnapshot) Point15ValDEnforcementDetailProjection {
	return Point15ValDEnforcementDetailProjection{
		ProjectionMode:       point15ValDModeEnforcementDetail,
		ProjectionAction:     point15ValDActionDisplayOnly,
		Visibility:           point15ValDVisibilityTenantScoped,
		EnforcementRef:       dependency.Point15ValC.EnforcementAction.EnforcementID,
		EnforcementAction:    dependency.Point15ValC.EnforcementAction.EnforcementAction,
		EnforcementReason:    dependency.Point15ValC.EnforcementAction.EnforcementReason,
		ReasonDecisive:       dependency.Point15ValC.EnforcementAction.ReasonDecisive,
		TargetState:          dependency.Point15ValC.EnforcementAction.TargetState,
		PriorState:           Point15Val0StateActive,
		CurrentState:         dependency.Point15ValC.EnforcementAction.TargetState,
		HistoryPreserved:     true,
		BlockedReasonVisible: true,
		PerformsEnforcement:  false,
		AutoRevokes:          false,
		AutoPublishes:        false,
		DeletesEvidence:      false,
		SilentReplacement:    false,
	}
}

func EvaluatePoint15ValDEnforcementDetailProjectionState(model Point15ValDEnforcementDetailProjection) string {
	if model.ProjectionMode != point15ValDModeEnforcementDetail ||
		(model.ProjectionAction != point15ValDActionDisplayOnly && model.ProjectionAction != point15ValDActionExplainOnly) ||
		!point15ValDVisibilityValid(model.Visibility) ||
		!point15ValDRefValid(model.EnforcementRef) ||
		!point15ValDRawActionValid(model.EnforcementAction) ||
		!point15ValDRawVal0StateValid(model.TargetState) ||
		!point15ValDRawVal0StateValid(model.PriorState) ||
		!point15ValDRawVal0StateValid(model.CurrentState) {
		return Point15ValDStateBlocked
	}
	if model.EnforcementAction == point15ValCActionNone {
		if model.EnforcementReason != "" || model.TargetState != Point15Val0StateActive || model.CurrentState != Point15Val0StateActive {
			return Point15ValDStateBlocked
		}
	} else {
		if !point15ValDRawReasonValid(model.EnforcementReason) {
			return Point15ValDStateBlocked
		}
		expectedAction, expectedState, _ := point15ValCExpectedAction(model.EnforcementReason, model.ReasonDecisive)
		if model.EnforcementAction != expectedAction || model.TargetState != expectedState {
			return Point15ValDStateBlocked
		}
	}
	if model.PerformsEnforcement || model.AutoRevokes || model.AutoPublishes || model.DeletesEvidence || model.SilentReplacement || !model.HistoryPreserved {
		return Point15ValDStateBlocked
	}
	if model.CurrentState == Point15Val0StateBlocked && !model.BlockedReasonVisible {
		return Point15ValDStateBlocked
	}
	return Point15ValDStateActive
}

func point15ValDReplayProofHistoryModel(dependency Point15ValDDependencySnapshot) Point15ValDReplayProofHistoryProjection {
	return Point15ValDReplayProofHistoryProjection{
		ProjectionMode:          point15ValDModeReplayDetail,
		ProjectionAction:        point15ValDActionDisplayOnly,
		Visibility:              point15ValDVisibilityTenantScoped,
		ReplayRef:               dependency.Point15ValC.ReplayProofHistory.ReplayRef,
		ProofPackRef:            dependency.Point15ValC.ReplayProofHistory.ProofPackRef,
		ProofHistoryRef:         dependency.Point15ValC.ReplayProofHistory.ProofHistoryRef,
		PriorStateVisible:       true,
		CurrentStateVisible:     true,
		DecisiveEvidenceVisible: true,
		BlockedReasonVisible:    true,
		HashBindingVisible:      true,
		ProofHistoryHidden:      false,
	}
}

func EvaluatePoint15ValDReplayProofHistoryProjectionState(model Point15ValDReplayProofHistoryProjection) string {
	if !point15ValDModeValid(model.ProjectionMode) ||
		!point15ValDActionValid(model.ProjectionAction) ||
		!point15ValDVisibilityValid(model.Visibility) ||
		!point15ValDRefValid(model.ReplayRef) ||
		!point15ValDRefValid(model.ProofPackRef) ||
		!point15ValDRefValid(model.ProofHistoryRef) {
		return Point15ValDStateBlocked
	}
	if model.ProjectionMode != point15ValDModeReplayDetail && model.ProjectionMode != point15ValDModeExportPreview {
		return Point15ValDStateBlocked
	}
	if model.ProjectionMode == point15ValDModeReplayDetail &&
		model.ProjectionAction != point15ValDActionDisplayOnly &&
		model.ProjectionAction != point15ValDActionExplainOnly {
		return Point15ValDStateBlocked
	}
	if model.ProjectionMode == point15ValDModeExportPreview && model.ProjectionAction != point15ValDActionExportPreviewOnly {
		return Point15ValDStateBlocked
	}
	if !model.PriorStateVisible || !model.CurrentStateVisible || !model.DecisiveEvidenceVisible || !model.BlockedReasonVisible || !model.HashBindingVisible || model.ProofHistoryHidden {
		return Point15ValDStateBlocked
	}
	return Point15ValDStateActive
}

func point15ValDAccessTenantPrivacyModel(dependency Point15ValDDependencySnapshot) Point15ValDAccessTenantPrivacyBoundary {
	return Point15ValDAccessTenantPrivacyBoundary{
		BoundaryID:               "access_tenant_point15_vald_001",
		TenantScope:              dependency.InheritedTenantScope,
		ViewerScope:              dependency.InheritedTenantScope,
		Visibility:               point15ValDVisibilityTenantScoped,
		RedactionState:           point15ValDRedactionNone,
		TenantPrivateDataExposed: false,
		CrossTenantDetected:      false,
		DecisiveFailureHidden:    false,
		ProjectionStateMutated:   false,
	}
}

func EvaluatePoint15ValDAccessTenantPrivacyBoundaryState(model Point15ValDAccessTenantPrivacyBoundary) string {
	if !point15ValDRefValid(model.BoundaryID) || !point15ValDVisibilityValid(model.Visibility) || !point15ValDRedactionStateValid(model.RedactionState) {
		return Point15ValDStateBlocked
	}
	if model.TenantScope == "" || model.ViewerScope == "" {
		return Point15ValDStateIncomplete
	}
	if !point15ValDRawScopeValid(model.TenantScope) || !point15ValDRawScopeValid(model.ViewerScope) {
		return Point15ValDStateBlocked
	}
	if model.CrossTenantDetected || model.ViewerScope != model.TenantScope || model.TenantPrivateDataExposed || model.Visibility == point15ValDVisibilityPublicBlocked || model.ProjectionStateMutated {
		return Point15ValDStateBlocked
	}
	if model.DecisiveFailureHidden {
		return Point15ValDStateReviewRequired
	}
	return Point15ValDStateActive
}

func point15ValDTimestampDisplayModel(dependency Point15ValDDependencySnapshot) Point15ValDTimestampDisplayDiscipline {
	valC := dependency.Point15ValC.TimestampDiscipline
	eventAt := valC.EnforcedAt
	if eventAt == "" {
		eventAt = valC.EvaluatedAt
	}
	return Point15ValDTimestampDisplayDiscipline{
		DisciplineID:                 "timestamp_display_point15_vald_001",
		ProjectionMode:               point15ValDModeTimeline,
		TenantScope:                  dependency.InheritedTenantScope,
		EventAt:                      eventAt,
		DisplayedAt:                  valC.ReferenceNow,
		SourceEventAt:                valC.SourceEventAt,
		ReceivedAt:                   valC.ReceivedAt,
		ValidatedAt:                  valC.ValidatedAt,
		EnforcedAt:                   valC.EnforcedAt,
		ReferenceNow:                 valC.ReferenceNow,
		TimeSource:                   valC.ReferenceNowTimeSource,
		ClientLocalCreatesCanonical:  false,
		SourceEventCreatesCanonical:  false,
		CanonicalOrderingFromDisplay: false,
	}
}

func EvaluatePoint15ValDTimestampDisplayDisciplineState(model Point15ValDTimestampDisplayDiscipline) string {
	if !point15ValDRefValid(model.DisciplineID) ||
		model.ProjectionMode != point15ValDModeTimeline ||
		!point15ValDRawScopeValid(model.TenantScope) ||
		!point12Val0RawTimestampValid(model.EventAt) ||
		!point12Val0RawTimestampValid(model.DisplayedAt) ||
		!point12Val0RawTimestampValid(model.ReferenceNow) ||
		!point14Val0CanonicalTimeSourceValid(model.TimeSource) {
		return Point15ValDStateBlocked
	}
	if model.ClientLocalCreatesCanonical || model.SourceEventCreatesCanonical || model.CanonicalOrderingFromDisplay {
		return Point15ValDStateBlocked
	}
	if !point15ValDOptionalRawTimestampValid(model.SourceEventAt) ||
		!point15ValDOptionalRawTimestampValid(model.ReceivedAt) ||
		!point15ValDOptionalRawTimestampValid(model.ValidatedAt) ||
		!point15ValDOptionalRawTimestampValid(model.EnforcedAt) {
		return Point15ValDStateBlocked
	}
	eventAt, _ := point14Val0ParsedTime(model.EventAt)
	displayedAt, _ := point14Val0ParsedTime(model.DisplayedAt)
	referenceNow, _ := point14Val0ParsedTime(model.ReferenceNow)
	if displayedAt.Before(eventAt) {
		return Point15ValDStateBlocked
	}
	if model.SourceEventAt != "" && model.EventAt == "" {
		return Point15ValDStateBlocked
	}
	if eventAt.After(referenceNow) {
		return Point15ValDStateReviewRequired
	}
	if model.EnforcedAt != "" && model.ValidatedAt != "" {
		enforcedAt, _ := point14Val0ParsedTime(model.EnforcedAt)
		validatedAt, _ := point14Val0ParsedTime(model.ValidatedAt)
		if enforcedAt.Before(validatedAt) {
			return Point15ValDStateReviewRequired
		}
	}
	if model.ValidatedAt != "" && model.ReceivedAt != "" {
		validatedAt, _ := point14Val0ParsedTime(model.ValidatedAt)
		receivedAt, _ := point14Val0ParsedTime(model.ReceivedAt)
		if validatedAt.Before(receivedAt) {
			return Point15ValDStateReviewRequired
		}
	}
	return Point15ValDStateActive
}

func point15ValDNoMutationGuardModel() Point15ValDNoMutationProjectionGuard {
	return Point15ValDNoMutationProjectionGuard{
		GuardID: "no_mutation_point15_vald_001",
	}
}

func EvaluatePoint15ValDNoMutationProjectionGuardState(model Point15ValDNoMutationProjectionGuard) string {
	if !point15ValDRefValid(model.GuardID) {
		return Point15ValDStateBlocked
	}
	if model.EvidenceMutationAttempted ||
		model.LifecycleMutationAttempted ||
		model.EnforcementMutationAttempted ||
		model.RevocationExecutionAttempted ||
		model.ExpiryDeletionAttempted ||
		model.SupersessionReplacementAttempted ||
		model.ScheduleRetryMutationAttempted ||
		model.PassRestoreAttempted {
		return Point15ValDStateBlocked
	}
	return Point15ValDStateActive
}

func point15ValDAuthorityBoundaryModel(dependency Point15ValDDependencySnapshot) Point15ValDAuthorityBoundary {
	return Point15ValDAuthorityBoundary{
		BoundaryID:                "authority_boundary_point15_vald_001",
		TenantScope:               dependency.InheritedTenantScope,
		FormalCoreOnly:            true,
		DashboardApprovesPass:     false,
		TimelineCreatesAuthority:  false,
		QueryEnforcesState:        false,
		ExportPreviewPublishes:    false,
		PortalMutationAttempted:   false,
		CustomerMutationAttempted: false,
		AuditorMutationAttempted:  false,
		ConnectorAuthorityGranted: false,
		SchedulerAuthorityGranted: false,
		AgentAuthorityGranted:     false,
		CanonicalMutationAllowed:  false,
		ProductionMutationAllowed: false,
		PassAllowed:               false,
	}
}

func EvaluatePoint15ValDAuthorityBoundaryState(model Point15ValDAuthorityBoundary) string {
	if !point15ValDRefValid(model.BoundaryID) || !point15ValDRawScopeValid(model.TenantScope) || !model.FormalCoreOnly {
		return Point15ValDStateBlocked
	}
	if model.DashboardApprovesPass ||
		model.TimelineCreatesAuthority ||
		model.QueryEnforcesState ||
		model.ExportPreviewPublishes ||
		model.PortalMutationAttempted ||
		model.CustomerMutationAttempted ||
		model.AuditorMutationAttempted ||
		model.ConnectorAuthorityGranted ||
		model.SchedulerAuthorityGranted ||
		model.AgentAuthorityGranted ||
		model.CanonicalMutationAllowed ||
		model.ProductionMutationAllowed ||
		model.PassAllowed {
		return Point15ValDStateBlocked
	}
	return Point15ValDStateActive
}

func point15ValDNoOverclaimGuardModel() Point15ValDNoOverclaimGuard {
	return Point15ValDNoOverclaimGuard{
		AllowedSafeWording:   point15ValDSafeWording(),
		BlockedWording:       point15ValDForbiddenWording(),
		ProjectionDisclaimer: point15ValDProjectionDisclaimer,
	}
}

func EvaluatePoint15ValDNoOverclaimGuardState(model Point15ValDNoOverclaimGuard) string {
	if model.ProjectionDisclaimer != point15ValDProjectionDisclaimer ||
		!point12Val0ExactStringSetMatch(model.AllowedSafeWording, point15ValDSafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point15ValDForbiddenWording()) {
		return Point15ValDStateBlocked
	}
	if point15ValDObservedListContainsForbiddenWording(model.ObservedTexts) {
		return Point15ValDStateBlocked
	}
	if point15ValDInternalListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point15ValDStateBlocked
	}
	return Point15ValDStateActive
}

func point15ValDFoundationModelFromUpstream(valC Point15ValCEnforcementBoundaryFoundation) Point15ValDAssuranceProjectionFoundation {
	dependency := point15ValDDependencyModelFromUpstream(valC)
	return Point15ValDAssuranceProjectionFoundation{
		ProjectionDisclaimer:       point15ValDProjectionDisclaimer,
		Dependency:                 dependency,
		Timeline:                   point15ValDAssuranceTimelineModel(dependency),
		Dashboard:                  point15ValDDashboardSummaryModel(dependency),
		Query:                      point15ValDQueryProjectionModel(dependency),
		EvidenceDetail:             point15ValDEvidenceDetailModel(dependency),
		RevalidationDetail:         point15ValDRevalidationDetailModel(dependency),
		EnforcementDetail:          point15ValDEnforcementDetailModel(dependency),
		ReplayProofHistory:         point15ValDReplayProofHistoryModel(dependency),
		AccessTenantPrivacy:        point15ValDAccessTenantPrivacyModel(dependency),
		TimestampDisplayDiscipline: point15ValDTimestampDisplayModel(dependency),
		NoMutationGuard:            point15ValDNoMutationGuardModel(),
		AuthorityBoundary:          point15ValDAuthorityBoundaryModel(dependency),
		NoOverclaimGuard:           point15ValDNoOverclaimGuardModel(),
	}
}

func Point15ValDFoundationModel() Point15ValDAssuranceProjectionFoundation {
	valC := ComputePoint15ValCEnforcementBoundaryFoundation(Point15ValCFoundationModel())
	return point15ValDFoundationModelFromUpstream(valC)
}

func point15ValDAggregate(states ...string) string {
	if len(states) == 0 {
		return Point15ValDStateBlocked
	}
	for _, state := range states {
		switch state {
		case Point15ValDStateActive, Point15ValDStateBlocked, Point15ValDStateReviewRequired, Point15ValDStateIncomplete:
		default:
			return Point15ValDStateBlocked
		}
	}
	for _, state := range states {
		if state == Point15ValDStateBlocked {
			return Point15ValDStateBlocked
		}
	}
	for _, state := range states {
		if state == Point15ValDStateReviewRequired {
			return Point15ValDStateReviewRequired
		}
	}
	for _, state := range states {
		if state == Point15ValDStateIncomplete {
			return Point15ValDStateIncomplete
		}
	}
	return Point15ValDStateActive
}

func point15ValDBlockingReasons(model Point15ValDAssuranceProjectionFoundation) []string {
	componentStates := map[string]string{
		"dependency":           model.DependencyState,
		"timeline":             model.TimelineState,
		"dashboard":            model.DashboardState,
		"query":                model.QueryState,
		"evidence_detail":      model.EvidenceDetailState,
		"revalidation_detail":  model.RevalidationDetailState,
		"enforcement_detail":   model.EnforcementDetailState,
		"replay_proof_history": model.ReplayProofHistoryState,
		"access_tenant":        model.AccessTenantState,
		"timestamp_display":    model.TimestampDisplayState,
		"no_mutation":          model.NoMutationState,
		"authority_boundary":   model.AuthorityBoundaryState,
		"no_overclaim":         model.NoOverclaimState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		switch state {
		case Point15ValDStateBlocked:
			reasons = append(reasons, name)
		case Point15ValDStateActive, Point15ValDStateReviewRequired, Point15ValDStateIncomplete:
		default:
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point15ValDReviewPrerequisites(model Point15ValDAssuranceProjectionFoundation) []string {
	componentStates := map[string]string{
		"timeline":             model.TimelineState,
		"dashboard":            model.DashboardState,
		"query":                model.QueryState,
		"evidence_detail":      model.EvidenceDetailState,
		"revalidation_detail":  model.RevalidationDetailState,
		"enforcement_detail":   model.EnforcementDetailState,
		"replay_proof_history": model.ReplayProofHistoryState,
		"access_tenant":        model.AccessTenantState,
		"timestamp_display":    model.TimestampDisplayState,
		"no_mutation":          model.NoMutationState,
		"authority_boundary":   model.AuthorityBoundaryState,
		"no_overclaim":         model.NoOverclaimState,
	}
	prereqs := append([]string{}, model.Dependency.ReviewPrerequisites...)
	for name, state := range componentStates {
		if state == Point15ValDStateReviewRequired || state == Point15ValDStateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func ComputePoint15ValDAssuranceProjectionFoundation(model Point15ValDAssuranceProjectionFoundation) Point15ValDAssuranceProjectionFoundation {
	model.DependencyState = EvaluatePoint15ValDDependencyState(model.Dependency)
	model.TimelineState = EvaluatePoint15ValDAssuranceTimelineState(model.Timeline)
	model.DashboardState = EvaluatePoint15ValDDashboardSummaryState(model.Dashboard)
	model.QueryState = EvaluatePoint15ValDQueryProjectionState(model.Query)
	model.EvidenceDetailState = EvaluatePoint15ValDEvidenceDetailProjectionState(model.EvidenceDetail)
	model.RevalidationDetailState = EvaluatePoint15ValDRevalidationDetailProjectionState(model.RevalidationDetail)
	model.EnforcementDetailState = EvaluatePoint15ValDEnforcementDetailProjectionState(model.EnforcementDetail)
	model.ReplayProofHistoryState = EvaluatePoint15ValDReplayProofHistoryProjectionState(model.ReplayProofHistory)
	model.AccessTenantState = EvaluatePoint15ValDAccessTenantPrivacyBoundaryState(model.AccessTenantPrivacy)
	model.TimestampDisplayState = EvaluatePoint15ValDTimestampDisplayDisciplineState(model.TimestampDisplayDiscipline)
	model.NoMutationState = EvaluatePoint15ValDNoMutationProjectionGuardState(model.NoMutationGuard)
	model.AuthorityBoundaryState = EvaluatePoint15ValDAuthorityBoundaryState(model.AuthorityBoundary)
	model.NoOverclaimState = EvaluatePoint15ValDNoOverclaimGuardState(model.NoOverclaimGuard)
	if model.ProjectionDisclaimer != point15ValDProjectionDisclaimer {
		model.NoOverclaimState = Point15ValDStateBlocked
	}

	expectedTenant := model.Dependency.InheritedTenantScope
	val0 := model.Dependency.Point15ValC.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0
	expectedEvidenceID := val0.EvidenceContext.EvidenceID
	expectedHash := val0.EvidenceContext.EvidenceHash
	expectedPolicy := val0.EvidenceContext.PolicyVersion
	expectedEngine := val0.EvidenceContext.EngineVersion
	expectedSchema := val0.EvidenceContext.SchemaVersion
	expectedFreshness := val0.FreshnessTaxonomy.FreshnessStatus
	expectedDowngrade := val0.DowngradeTaxonomy.DowngradeOutcome
	expectedLifecycle := model.Dependency.Point15ValC.EvidenceLifecycle.LifecycleStatus
	expectedEnforcementRef := model.Dependency.Point15ValC.EnforcementAction.EnforcementID
	expectedEnforcementAction := model.Dependency.Point15ValC.EnforcementAction.EnforcementAction
	expectedEnforcementReason := model.Dependency.Point15ValC.EnforcementAction.EnforcementReason
	expectedTargetState := model.Dependency.Point15ValC.EnforcementAction.TargetState
	expectedReplayRef := model.Dependency.Point15ValC.ReplayProofHistory.ReplayRef
	expectedProofHistoryRef := model.Dependency.Point15ValC.ReplayProofHistory.ProofHistoryRef
	expectedProofPackRef := model.Dependency.Point15ValC.ReplayProofHistory.ProofPackRef
	expectedScheduleRef := model.Dependency.Point15ValC.Dependency.Point15ValB.Schedule.ScheduleID
	expectedRunRef := model.Dependency.Point15ValC.Dependency.Point15ValB.Run.RunID
	expectedRetryRef := model.Dependency.Point15ValC.Dependency.Point15ValB.RetryBudget.BudgetID
	expectedThrottleRef := model.Dependency.Point15ValC.Dependency.Point15ValB.TenantThrottle.ThrottleID
	expectedScheduledStatus := model.Dependency.Point15ValC.Dependency.Point15ValB.Schedule.ScheduledStatus
	expectedRunResult := model.Dependency.Point15ValC.Dependency.Point15ValB.Run.RunResult
	expectedRetryStatus := model.Dependency.Point15ValC.Dependency.Point15ValB.RetryBudget.RetryBudgetStatus
	expectedThrottleStatus := model.Dependency.Point15ValC.Dependency.Point15ValB.TenantThrottle.ThrottleStatus

	if !formalRawExactValid(expectedTenant, point11Val0ScopeValid) ||
		model.Timeline.TenantScope != expectedTenant ||
		model.Dashboard.TenantScope != expectedTenant ||
		model.Query.TenantScope != expectedTenant ||
		model.EvidenceDetail.TenantScope != expectedTenant ||
		model.AccessTenantPrivacy.TenantScope != expectedTenant ||
		model.TimestampDisplayDiscipline.TenantScope != expectedTenant ||
		model.AuthorityBoundary.TenantScope != expectedTenant {
		model.TimelineState = Point15ValDStateBlocked
		model.DashboardState = Point15ValDStateBlocked
		model.QueryState = Point15ValDStateBlocked
		model.EvidenceDetailState = Point15ValDStateBlocked
		model.AccessTenantState = Point15ValDStateBlocked
		model.TimestampDisplayState = Point15ValDStateBlocked
		model.AuthorityBoundaryState = Point15ValDStateBlocked
	}
	if !formalRawExactNonEmpty(expectedEvidenceID) ||
		model.Timeline.EvidenceID != expectedEvidenceID ||
		model.EvidenceDetail.EvidenceID != expectedEvidenceID {
		model.TimelineState = Point15ValDStateBlocked
		model.EvidenceDetailState = Point15ValDStateBlocked
		model.QueryState = Point15ValDStateBlocked
	}
	if model.EvidenceDetail.EvidenceHash != expectedHash ||
		model.EvidenceDetail.PolicyVersion != expectedPolicy ||
		model.EvidenceDetail.EngineVersion != expectedEngine ||
		model.EvidenceDetail.SchemaVersion != expectedSchema ||
		model.EvidenceDetail.FreshnessStatus != expectedFreshness ||
		model.EvidenceDetail.DowngradeOutcome != expectedDowngrade ||
		model.EvidenceDetail.LifecycleStatus != expectedLifecycle ||
		model.EvidenceDetail.EnforcementStatus != expectedTargetState {
		model.EvidenceDetailState = Point15ValDStateBlocked
	}
	if model.Timeline.SourceValCRef != expectedEnforcementRef ||
		model.Timeline.ReplayRef != expectedReplayRef ||
		model.Timeline.ProofHistoryRef != expectedProofHistoryRef {
		model.TimelineState = Point15ValDStateBlocked
	}
	if model.EnforcementDetail.EnforcementRef != expectedEnforcementRef ||
		model.EnforcementDetail.EnforcementAction != expectedEnforcementAction ||
		model.EnforcementDetail.EnforcementReason != expectedEnforcementReason ||
		model.EnforcementDetail.TargetState != expectedTargetState ||
		model.EnforcementDetail.CurrentState != expectedTargetState ||
		model.EnforcementDetail.CurrentState != model.Timeline.CurrentState ||
		model.EnforcementDetail.PriorState != model.Timeline.PriorState {
		model.EnforcementDetailState = Point15ValDStateBlocked
		model.TimelineState = Point15ValDStateBlocked
	}
	if model.ReplayProofHistory.ReplayRef != expectedReplayRef ||
		model.ReplayProofHistory.ProofHistoryRef != expectedProofHistoryRef ||
		model.ReplayProofHistory.ProofPackRef != expectedProofPackRef {
		model.ReplayProofHistoryState = Point15ValDStateBlocked
	}
	if !point12Val0ExactStringSetMatch(model.Query.ResultRefs, []string{expectedEvidenceID, expectedEnforcementRef, expectedReplayRef}) {
		model.QueryState = Point15ValDStateBlocked
	}
	if model.RevalidationDetail.ScheduleRef != expectedScheduleRef ||
		model.RevalidationDetail.RunRef != expectedRunRef ||
		model.RevalidationDetail.RetryBudgetRef != expectedRetryRef ||
		model.RevalidationDetail.TenantThrottleRef != expectedThrottleRef ||
		model.RevalidationDetail.ScheduledStatus != expectedScheduledStatus ||
		model.RevalidationDetail.RunResult != expectedRunResult ||
		model.RevalidationDetail.RetryStatus != expectedRetryStatus ||
		model.RevalidationDetail.ThrottleStatus != expectedThrottleStatus {
		model.RevalidationDetailState = Point15ValDStateBlocked
	}

	expectedActiveCount := 0
	expectedBlockedCount := 0
	expectedReviewCount := 0
	expectedIncompleteCount := 0
	switch model.EvidenceDetail.EnforcementStatus {
	case Point15Val0StateActive:
		expectedActiveCount = 1
	case Point15Val0StateBlocked:
		expectedBlockedCount = 1
	case Point15Val0StateReviewRequired:
		expectedReviewCount = 1
	case Point15Val0StateIncomplete:
		expectedIncompleteCount = 1
	}
	expectedStaleCount := 0
	if model.EvidenceDetail.FreshnessStatus == point15Val0FreshnessStale || model.EvidenceDetail.LifecycleStatus == point15ValCLifecycleStale {
		expectedStaleCount = 1
	}
	expectedExpiredCount := 0
	if model.EvidenceDetail.LifecycleStatus == point15ValCLifecycleExpired {
		expectedExpiredCount = 1
	}
	expectedRevokedCount := 0
	if model.EvidenceDetail.LifecycleStatus == point15ValCLifecycleRevoked {
		expectedRevokedCount = 1
	}
	expectedEnforcementCount := 0
	if model.EnforcementDetail.EnforcementAction != point15ValCActionNone {
		expectedEnforcementCount = 1
	}
	if model.Dashboard.EvidenceCount != 1 ||
		model.Dashboard.ActiveCount != expectedActiveCount ||
		model.Dashboard.BlockedCount != expectedBlockedCount ||
		model.Dashboard.ReviewRequiredCount != expectedReviewCount ||
		model.Dashboard.IncompleteCount != expectedIncompleteCount ||
		model.Dashboard.StaleCount != expectedStaleCount ||
		model.Dashboard.ExpiredCount != expectedExpiredCount ||
		model.Dashboard.RevokedCount != expectedRevokedCount ||
		model.Dashboard.EnforcementCount != expectedEnforcementCount {
		model.DashboardState = Point15ValDStateBlocked
	}

	model.CurrentState = point15ValDAggregate(
		model.DependencyState,
		model.TimelineState,
		model.DashboardState,
		model.QueryState,
		model.EvidenceDetailState,
		model.RevalidationDetailState,
		model.EnforcementDetailState,
		model.ReplayProofHistoryState,
		model.AccessTenantState,
		model.TimestampDisplayState,
		model.NoMutationState,
		model.AuthorityBoundaryState,
		model.NoOverclaimState,
	)
	model.BlockingReasons = point15ValDBlockingReasons(model)
	model.ReviewPrerequisites = point15ValDReviewPrerequisites(model)
	return model
}

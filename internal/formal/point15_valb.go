package formal

import (
	"encoding/json"
	"sort"
	"strings"
)

const (
	Point15ValBStateActive         = "point15_valb_scheduled_revalidation_active"
	Point15ValBStateBlocked        = "point15_valb_scheduled_revalidation_blocked"
	Point15ValBStateReviewRequired = "point15_valb_scheduled_revalidation_review_required"
	Point15ValBStateIncomplete     = "point15_valb_scheduled_revalidation_incomplete"
)

const (
	point15ValBWaveID                     = "val_b"
	point15ValBRevalidationDisclaimer     = "formal_scheduled_revalidation_gate no_silent_pass_retention point15_valb"
	point15ValBBlockedPassToken           = "point_15_pass"
	point15ValBScheduleNotRequired        = "not_required"
	point15ValBScheduleScheduled          = "scheduled"
	point15ValBScheduleDue                = "due"
	point15ValBScheduleOverdue            = "overdue"
	point15ValBScheduleRunning            = "running"
	point15ValBScheduleCompleted          = "completed"
	point15ValBScheduleMissed             = "missed"
	point15ValBScheduleFailed             = "failed"
	point15ValBScheduleRetryPending       = "retry_pending"
	point15ValBScheduleRetryExhausted     = "retry_exhausted"
	point15ValBScheduleThrottled          = "throttled"
	point15ValBScheduleBlocked            = "blocked"
	point15ValBRunNotRun                  = "not_run"
	point15ValBRunCompletedClean          = "completed_clean"
	point15ValBRunCompletedWithDowngrade  = "completed_with_downgrade"
	point15ValBRunFailed                  = "failed"
	point15ValBRunMissed                  = "missed"
	point15ValBRunUnauthorized            = "unauthorized"
	point15ValBRunTenantMismatch          = "tenant_mismatch"
	point15ValBRunTimeout                 = "timeout"
	point15ValBRunThrottled               = "throttled"
	point15ValBRunTampered                = "tampered"
	point15ValBRetryNotApplicable         = "not_applicable"
	point15ValBRetryAvailable             = "available"
	point15ValBRetryExhausted             = "exhausted"
	point15ValBRetryBlocked               = "blocked"
	point15ValBRetryReasonNone            = ""
	point15ValBRetryReasonWindow          = "transient_failure_retry_window"
	point15ValBRetryReasonManualReview    = "manual_review_after_exhaustion"
	point15ValBRetryReasonTerminal        = "terminal_exhaustion"
	point15ValBRetryReasonPolicyBlocked   = "policy_blocked_retry_budget"
	point15ValBThrottleNotApplicable      = "not_applicable"
	point15ValBThrottleWithinLimit        = "within_limit"
	point15ValBThrottleReviewRequired     = "throttled_review_required"
	point15ValBThrottleBlocked            = "throttled_blocked"
	point15ValBThrottleCrossTenantBlocked = "cross_tenant_blocked"
)

type Point15ValBDependencySnapshot struct {
	Point15ValACurrentState          string                                `json:"point15_vala_current_state"`
	Point15ValADependencyState       string                                `json:"point15_vala_dependency_state"`
	Point15ValATriggerTableState     string                                `json:"point15_vala_trigger_table_state"`
	Point15ValATriggerState          string                                `json:"point15_vala_trigger_state"`
	Point15ValAReasonState           string                                `json:"point15_vala_reason_state"`
	Point15ValADecisionState         string                                `json:"point15_vala_decision_state"`
	Point15ValAAuthorityState        string                                `json:"point15_vala_authority_state"`
	Point15ValANoOverclaimState      string                                `json:"point15_vala_no_overclaim_state"`
	Point15ValAComputedFromUpstream  bool                                  `json:"point15_vala_computed_from_upstream"`
	Point15ValAMerged                bool                                  `json:"point15_vala_merged"`
	Point15ValACIGreen               bool                                  `json:"point15_vala_ci_green"`
	Point15ValAReviewedOnMain        bool                                  `json:"point15_vala_reviewed_on_main"`
	Point15PassSeen                  bool                                  `json:"point15_pass_seen"`
	InheritedPoint15Val0CurrentState string                                `json:"inherited_point15_val0_current_state"`
	InheritedPoint14ValECurrentState string                                `json:"inherited_point14_vale_current_state"`
	InheritedTenantScope             string                                `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput       bool                                  `json:"snapshot_from_computed_output"`
	ReviewPrerequisites              []string                              `json:"review_prerequisites,omitempty"`
	Point15ValA                      Point15ValADowngradeTriggerFoundation `json:"point15_vala"`
}

type Point15ValBScheduledRevalidationFoundation struct {
	CurrentState             string                          `json:"current_state"`
	BlockingReasons          []string                        `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites      []string                        `json:"review_prerequisites,omitempty"`
	RevalidationDisclaimer   string                          `json:"revalidation_disclaimer"`
	DependencyState          string                          `json:"dependency_state"`
	ScheduleState            string                          `json:"schedule_state"`
	RunState                 string                          `json:"run_state"`
	RetryBudgetState         string                          `json:"retry_budget_state"`
	TenantThrottleState      string                          `json:"tenant_throttle_state"`
	DowngradeBindingState    string                          `json:"downgrade_binding_state"`
	TimestampDisciplineState string                          `json:"timestamp_discipline_state"`
	AuthorityBoundaryState   string                          `json:"authority_boundary_state"`
	NoOverclaimState         string                          `json:"no_overclaim_state"`
	Dependency               Point15ValBDependencySnapshot   `json:"dependency"`
	Schedule                 Point15ValBRevalidationSchedule `json:"schedule"`
	Run                      Point15ValBRevalidationRun      `json:"run"`
	RetryBudget              Point15ValBRetryBudget          `json:"retry_budget"`
	TenantThrottle           Point15ValBTenantThrottle       `json:"tenant_throttle"`
	DowngradeBinding         Point15ValBDowngradeBinding     `json:"downgrade_binding"`
	TimestampDiscipline      Point15ValBTimestampDiscipline  `json:"timestamp_discipline"`
	AuthorityBoundary        Point15ValBAuthorityBoundary    `json:"authority_boundary"`
	NoOverclaimGuard         Point15ValBNoOverclaimGuard     `json:"no_overclaim_guard"`
}

type Point15ValBRevalidationSchedule struct {
	ScheduleID          string `json:"schedule_id"`
	EvidenceID          string `json:"evidence_id"`
	TenantScope         string `json:"tenant_scope"`
	Required            bool   `json:"required"`
	ScheduledStatus     string `json:"scheduled_status"`
	ScheduledAt         string `json:"scheduled_at"`
	RevalidationDueAt   string `json:"revalidation_due_at"`
	LastCompletedAt     string `json:"last_completed_at"`
	NextRetryAt         string `json:"next_retry_at"`
	SchedulerTimeSource string `json:"scheduler_time_source"`
	PolicyVersion       string `json:"policy_version"`
	EngineVersion       string `json:"engine_version"`
	SchemaVersion       string `json:"schema_version"`
}

type Point15ValBRevalidationRun struct {
	RunID               string `json:"run_id"`
	ScheduleRef         string `json:"schedule_ref"`
	EvidenceID          string `json:"evidence_id"`
	TenantScope         string `json:"tenant_scope"`
	StartedAt           string `json:"started_at"`
	CompletedAt         string `json:"completed_at"`
	RunResult           string `json:"run_result"`
	RunEvidenceHash     string `json:"run_evidence_hash"`
	ConnectorResultRef  string `json:"connector_result_ref"`
	DowngradeTriggerRef string `json:"downgrade_trigger_ref"`
	PolicyVersion       string `json:"policy_version"`
	EngineVersion       string `json:"engine_version"`
	SchemaVersion       string `json:"schema_version"`
}

type Point15ValBRetryBudget struct {
	BudgetID            string `json:"budget_id"`
	ScheduleRef         string `json:"schedule_ref"`
	MaxRetries          int    `json:"max_retries"`
	AttemptsUsed        int    `json:"attempts_used"`
	RetryBudgetStatus   string `json:"retry_budget_status"`
	RetryReason         string `json:"retry_reason"`
	NextRetryAt         string `json:"next_retry_at"`
	NextRetryTimeSource string `json:"next_retry_time_source"`
	SelfResetDetected   bool   `json:"self_reset_detected"`
}

type Point15ValBTenantThrottle struct {
	ThrottleID             string `json:"throttle_id"`
	TenantScope            string `json:"tenant_scope"`
	RequestedRevalidations int    `json:"requested_revalidations"`
	AllowedRevalidations   int    `json:"allowed_revalidations"`
	ThrottleStatus         string `json:"throttle_status"`
	CrossTenantDetected    bool   `json:"cross_tenant_detected"`
}

type Point15ValBDowngradeBinding struct {
	BindingID              string `json:"binding_id"`
	ScheduleRef            string `json:"schedule_ref"`
	RunRef                 string `json:"run_ref"`
	RetryBudgetRef         string `json:"retry_budget_ref"`
	ThrottleRef            string `json:"throttle_ref"`
	RequiredRevalidation   bool   `json:"required_revalidation"`
	ScheduleStatus         string `json:"schedule_status"`
	RunResult              string `json:"run_result"`
	RetryBudgetStatus      string `json:"retry_budget_status"`
	ThrottleStatus         string `json:"throttle_status"`
	LastCompletedAt        string `json:"last_completed_at"`
	RunEvidenceHashMatches bool   `json:"run_evidence_hash_matches"`
	TriggerType            string `json:"trigger_type"`
	TriggerIsDecisive      bool   `json:"trigger_is_decisive"`
	TargetState            string `json:"target_state"`
	TargetDowngradeOutcome string `json:"target_downgrade_outcome"`
	RetainsPass            bool   `json:"retains_pass"`
	RetainsActiveClosure   bool   `json:"retains_active_closure"`
}

type Point15ValBTimestampDiscipline struct {
	DisciplineID                string `json:"discipline_id"`
	TenantScope                 string `json:"tenant_scope"`
	ScheduledStatus             string `json:"scheduled_status"`
	ScheduledAt                 string `json:"scheduled_at"`
	ScheduledAtTimeSource       string `json:"scheduled_at_time_source"`
	DueAt                       string `json:"due_at"`
	DueAtTimeSource             string `json:"due_at_time_source"`
	StartedAt                   string `json:"started_at"`
	StartedAtTimeSource         string `json:"started_at_time_source"`
	CompletedAt                 string `json:"completed_at"`
	CompletedAtTimeSource       string `json:"completed_at_time_source"`
	NextRetryAt                 string `json:"next_retry_at"`
	NextRetryAtTimeSource       string `json:"next_retry_at_time_source"`
	ReferenceNow                string `json:"reference_now"`
	ReferenceNowTimeSource      string `json:"reference_now_time_source"`
	SourceEventAt               string `json:"source_event_at"`
	SourceEventTimeSource       string `json:"source_event_time_source"`
	ClientLocalCreatesCanonical bool   `json:"client_local_creates_canonical"`
	SourceEventCreatesCanonical bool   `json:"source_event_creates_canonical"`
}

type Point15ValBAuthorityBoundary struct {
	BoundaryID                            string `json:"boundary_id"`
	TenantScope                           string `json:"tenant_scope"`
	ExternalSourceInputOnly               bool   `json:"external_source_input_only"`
	FormalEvaluatorOnly                   bool   `json:"formal_evaluator_only"`
	AgentRecommendationAdvisoryOnly       bool   `json:"agent_recommendation_advisory_only"`
	SchedulerMarksEvidenceFresh           bool   `json:"scheduler_marks_evidence_fresh"`
	SchedulerCreatesRevalidationTruth     bool   `json:"scheduler_creates_revalidation_truth"`
	ConnectorRestoresActiveClosure        bool   `json:"connector_restores_active_closure"`
	DashboardSuppressesOverdueStatus      bool   `json:"dashboard_suppresses_overdue_status"`
	PortalProjectionMutatesRevalidation   bool   `json:"portal_projection_mutates_revalidation"`
	CustomerProjectionMutatesRevalidation bool   `json:"customer_projection_mutates_revalidation"`
	AuditorProjectionMutatesRevalidation  bool   `json:"auditor_projection_mutates_revalidation"`
	AgentSatisfiesRevalidation            bool   `json:"agent_satisfies_revalidation"`
	RetryBudgetResetAllowed               bool   `json:"retry_budget_reset_allowed"`
	CanonicalMutationAllowed              bool   `json:"canonical_mutation_allowed"`
	ProductionMutationAllowed             bool   `json:"production_mutation_allowed"`
	PassAllowed                           bool   `json:"pass_allowed"`
}

type Point15ValBNoOverclaimGuard struct {
	ObservedTexts                        []string `json:"observed_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	RevalidationDisclaimer               string   `json:"revalidation_disclaimer"`
}

func point15ValBStates() []string {
	return []string{
		Point15ValBStateActive,
		Point15ValBStateBlocked,
		Point15ValBStateReviewRequired,
		Point15ValBStateIncomplete,
	}
}

func point15ValBStateValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValBStates())
}

func point15ValBScheduleStatuses() []string {
	return []string{
		point15ValBScheduleNotRequired,
		point15ValBScheduleScheduled,
		point15ValBScheduleDue,
		point15ValBScheduleOverdue,
		point15ValBScheduleRunning,
		point15ValBScheduleCompleted,
		point15ValBScheduleMissed,
		point15ValBScheduleFailed,
		point15ValBScheduleRetryPending,
		point15ValBScheduleRetryExhausted,
		point15ValBScheduleThrottled,
		point15ValBScheduleBlocked,
	}
}

func point15ValBScheduleStatusValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValBScheduleStatuses())
}

func point15ValBRunResults() []string {
	return []string{
		point15ValBRunNotRun,
		point15ValBRunCompletedClean,
		point15ValBRunCompletedWithDowngrade,
		point15ValBRunFailed,
		point15ValBRunMissed,
		point15ValBRunUnauthorized,
		point15ValBRunTenantMismatch,
		point15ValBRunTimeout,
		point15ValBRunThrottled,
		point15ValBRunTampered,
	}
}

func point15ValBRunResultValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValBRunResults())
}

func point15ValBRetryStatuses() []string {
	return []string{
		point15ValBRetryNotApplicable,
		point15ValBRetryAvailable,
		point15ValBRetryExhausted,
		point15ValBRetryBlocked,
	}
}

func point15ValBRetryStatusValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValBRetryStatuses())
}

func point15ValBRetryReasons() []string {
	return []string{
		point15ValBRetryReasonNone,
		point15ValBRetryReasonWindow,
		point15ValBRetryReasonManualReview,
		point15ValBRetryReasonTerminal,
		point15ValBRetryReasonPolicyBlocked,
	}
}

func point15ValBRetryReasonValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValBRetryReasons())
}

func point15ValBThrottleStatuses() []string {
	return []string{
		point15ValBThrottleNotApplicable,
		point15ValBThrottleWithinLimit,
		point15ValBThrottleReviewRequired,
		point15ValBThrottleBlocked,
		point15ValBThrottleCrossTenantBlocked,
	}
}

func point15ValBThrottleStatusValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValBThrottleStatuses())
}

func point15ValBDependencyRefValid(value string) bool {
	return point14Val0RefValid(value, "point15_valb_", "schedule_", "run_", "retry_", "throttle_", "binding_", "authority_", "timestamp_")
}

func point15ValBConnectorResultRefValid(value string) bool {
	return point14Val0RefValid(value, "connector_result_")
}

func point15ValBForbiddenWording() []string {
	return []string{
		"continuous assurance guaranteed",
		"automatically verified forever",
		"always fresh",
		"certified secure",
		"compliance guaranteed",
		"production approved",
		"deployment approved",
		"legal proof",
		"financial guarantee",
		"official authority",
		"global truth",
		"public badge",
		"regulator-approved",
		"regulator approved",
	}
}

func point15ValBSafeWording() []string {
	return []string{
		"scheduled revalidation requires formal evaluator",
		"revalidation overdue requires downgrade review",
		"bounded revalidation decision",
		"retry budget status is evaluator bound",
		"trigger mapped by formal evaluator only",
	}
}

func point15ValBObservedTextContainsForbiddenWording(text string) bool {
	trimmed := strings.TrimSpace(strings.ToLower(text))
	if trimmed == "" {
		return false
	}
	for _, safe := range point15ValBSafeWording() {
		if trimmed == strings.ToLower(strings.TrimSpace(safe)) {
			return false
		}
	}
	for _, forbidden := range point15ValBForbiddenWording() {
		if strings.Contains(trimmed, strings.ToLower(strings.TrimSpace(forbidden))) {
			return true
		}
	}
	return false
}

func point15ValBObservedListContainsForbiddenWording(values []string) bool {
	for _, value := range values {
		if point15ValBObservedTextContainsForbiddenWording(value) {
			return true
		}
	}
	return false
}

func point15ValBValAPayloadContainsPoint15Pass(valA Point15ValADowngradeTriggerFoundation) bool {
	payload, err := json.Marshal(valA)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point15ValBBlockedPassToken)
}

func point15ValBTargetStateToWaveState(state string) string {
	switch strings.TrimSpace(state) {
	case Point15Val0StateActive:
		return Point15ValBStateActive
	case Point15Val0StateReviewRequired:
		return Point15ValBStateReviewRequired
	case Point15Val0StateIncomplete:
		return Point15ValBStateIncomplete
	case Point15Val0StateBlocked:
		return Point15ValBStateBlocked
	default:
		return Point15ValBStateBlocked
	}
}

func point15ValBDependencySnapshotFromUpstream(valA Point15ValADowngradeTriggerFoundation) Point15ValBDependencySnapshot {
	return Point15ValBDependencySnapshot{
		Point15ValACurrentState:          valA.CurrentState,
		Point15ValADependencyState:       valA.DependencyState,
		Point15ValATriggerTableState:     valA.TriggerTableState,
		Point15ValATriggerState:          valA.TriggerState,
		Point15ValAReasonState:           valA.ReasonState,
		Point15ValADecisionState:         valA.DecisionState,
		Point15ValAAuthorityState:        valA.AuthorityBoundaryState,
		Point15ValANoOverclaimState:      valA.NoOverclaimState,
		Point15ValAComputedFromUpstream:  valA.Dependency.SnapshotFromComputedOutput,
		Point15ValAMerged:                true,
		Point15ValACIGreen:               true,
		Point15ValAReviewedOnMain:        true,
		Point15PassSeen:                  point15ValBValAPayloadContainsPoint15Pass(valA),
		InheritedPoint15Val0CurrentState: valA.Dependency.Point15Val0CurrentState,
		InheritedPoint14ValECurrentState: valA.Dependency.InheritedPoint14ValECurrentState,
		InheritedTenantScope:             valA.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:       true,
		ReviewPrerequisites:              append([]string{}, valA.ReviewPrerequisites...),
		Point15ValA:                      valA,
	}
}

func point15ValBDependencySnapshotModel() Point15ValBDependencySnapshot {
	valA := ComputePoint15ValADowngradeTriggerFoundation(Point15ValAFoundationModel())
	return point15ValBDependencySnapshotFromUpstream(valA)
}

func EvaluatePoint15ValBDependencyState(model Point15ValBDependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point15ValAComputedFromUpstream ||
		!model.Point15ValAMerged ||
		!model.Point15ValACIGreen ||
		!model.Point15ValAReviewedOnMain ||
		model.Point15PassSeen ||
		!point15ValAStateValid(model.Point15ValACurrentState) ||
		!point15ValAStateValid(model.Point15ValADependencyState) ||
		!point15ValAStateValid(model.Point15ValATriggerTableState) ||
		!point15ValAStateValid(model.Point15ValATriggerState) ||
		!point15ValAStateValid(model.Point15ValAReasonState) ||
		!point15ValAStateValid(model.Point15ValADecisionState) ||
		!point15ValAStateValid(model.Point15ValAAuthorityState) ||
		!point15ValAStateValid(model.Point15ValANoOverclaimState) ||
		!point15Val0StateValid(model.InheritedPoint15Val0CurrentState) ||
		!point14ValEStateValid(model.InheritedPoint14ValECurrentState) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.Point15ValACurrentState) != strings.TrimSpace(model.Point15ValA.CurrentState) ||
		strings.TrimSpace(model.Point15ValADependencyState) != strings.TrimSpace(model.Point15ValA.DependencyState) ||
		strings.TrimSpace(model.Point15ValATriggerTableState) != strings.TrimSpace(model.Point15ValA.TriggerTableState) ||
		strings.TrimSpace(model.Point15ValATriggerState) != strings.TrimSpace(model.Point15ValA.TriggerState) ||
		strings.TrimSpace(model.Point15ValAReasonState) != strings.TrimSpace(model.Point15ValA.ReasonState) ||
		strings.TrimSpace(model.Point15ValADecisionState) != strings.TrimSpace(model.Point15ValA.DecisionState) ||
		strings.TrimSpace(model.Point15ValAAuthorityState) != strings.TrimSpace(model.Point15ValA.AuthorityBoundaryState) ||
		strings.TrimSpace(model.Point15ValANoOverclaimState) != strings.TrimSpace(model.Point15ValA.NoOverclaimState) ||
		model.Point15ValAComputedFromUpstream != model.Point15ValA.Dependency.SnapshotFromComputedOutput ||
		strings.TrimSpace(model.InheritedPoint15Val0CurrentState) != strings.TrimSpace(model.Point15ValA.Dependency.Point15Val0CurrentState) ||
		strings.TrimSpace(model.InheritedPoint14ValECurrentState) != strings.TrimSpace(model.Point15ValA.Dependency.InheritedPoint14ValECurrentState) ||
		strings.TrimSpace(model.InheritedTenantScope) != strings.TrimSpace(model.Point15ValA.Dependency.InheritedTenantScope) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.Point15ValACurrentState) != Point15ValAStateActive ||
		strings.TrimSpace(model.Point15ValADependencyState) != Point15ValAStateActive ||
		strings.TrimSpace(model.Point15ValATriggerTableState) != Point15ValAStateActive ||
		strings.TrimSpace(model.Point15ValATriggerState) != Point15ValAStateActive ||
		strings.TrimSpace(model.Point15ValAReasonState) != Point15ValAStateActive ||
		strings.TrimSpace(model.Point15ValADecisionState) != Point15ValAStateActive ||
		strings.TrimSpace(model.Point15ValAAuthorityState) != Point15ValAStateActive ||
		strings.TrimSpace(model.Point15ValANoOverclaimState) != Point15ValAStateActive ||
		strings.TrimSpace(model.InheritedPoint15Val0CurrentState) != Point15Val0StateActive ||
		strings.TrimSpace(model.InheritedPoint14ValECurrentState) != Point14ValEStatePassConfirmed {
		return Point15ValBStateBlocked
	}
	return Point15ValBStateActive
}

func point15ValBExpectedScheduleState(required bool, status string) string {
	switch strings.TrimSpace(status) {
	case point15ValBScheduleNotRequired:
		if required {
			return Point15Val0StateBlocked
		}
		return Point15Val0StateActive
	case point15ValBScheduleScheduled, point15ValBScheduleCompleted:
		return Point15Val0StateActive
	case point15ValBScheduleDue, point15ValBScheduleOverdue, point15ValBScheduleRunning, point15ValBScheduleMissed, point15ValBScheduleFailed, point15ValBScheduleRetryPending, point15ValBScheduleThrottled:
		return Point15Val0StateReviewRequired
	case point15ValBScheduleRetryExhausted, point15ValBScheduleBlocked:
		return Point15Val0StateBlocked
	default:
		return Point15Val0StateBlocked
	}
}

func point15ValBExpectedRunState(result string) string {
	switch strings.TrimSpace(result) {
	case point15ValBRunNotRun, point15ValBRunCompletedClean:
		return Point15Val0StateActive
	case point15ValBRunCompletedWithDowngrade, point15ValBRunFailed, point15ValBRunMissed, point15ValBRunTimeout, point15ValBRunThrottled:
		return Point15Val0StateReviewRequired
	case point15ValBRunUnauthorized, point15ValBRunTenantMismatch, point15ValBRunTampered:
		return Point15Val0StateBlocked
	default:
		return Point15Val0StateBlocked
	}
}

func point15ValBExpectedRetryState(status, reason string) string {
	switch strings.TrimSpace(status) {
	case point15ValBRetryNotApplicable, point15ValBRetryAvailable:
		return Point15Val0StateActive
	case point15ValBRetryExhausted:
		if strings.TrimSpace(reason) == point15ValBRetryReasonManualReview {
			return Point15Val0StateReviewRequired
		}
		return Point15Val0StateBlocked
	case point15ValBRetryBlocked:
		return Point15Val0StateBlocked
	default:
		return Point15Val0StateBlocked
	}
}

func point15ValBExpectedThrottleState(status string) string {
	switch strings.TrimSpace(status) {
	case point15ValBThrottleNotApplicable, point15ValBThrottleWithinLimit:
		return Point15Val0StateActive
	case point15ValBThrottleReviewRequired:
		return Point15Val0StateReviewRequired
	case point15ValBThrottleBlocked, point15ValBThrottleCrossTenantBlocked:
		return Point15Val0StateBlocked
	default:
		return Point15Val0StateBlocked
	}
}

func point15ValBExpectedBinding(model Point15ValBDowngradeBinding) (string, string, string) {
	if !model.RequiredRevalidation {
		return "", point15Val0DowngradeRetainActive, Point15Val0StateActive
	}
	if !model.RunEvidenceHashMatches && (strings.TrimSpace(model.RunResult) == point15ValBRunCompletedClean || strings.TrimSpace(model.RunResult) == point15ValBRunCompletedWithDowngrade) {
		return point15ValATriggerHash, point15Val0DowngradeBlocked, Point15Val0StateBlocked
	}
	switch strings.TrimSpace(model.RunResult) {
	case point15ValBRunUnauthorized:
		return point15ValATriggerConnAuth, point15Val0DowngradeBlocked, Point15Val0StateBlocked
	case point15ValBRunTenantMismatch:
		return point15ValATriggerConnTenant, point15Val0DowngradeBlocked, Point15Val0StateBlocked
	case point15ValBRunTimeout:
		return point15ValATriggerConnTimeout, point15Val0DowngradeReview, Point15Val0StateReviewRequired
	case point15ValBRunTampered:
		return point15ValATriggerTampered, point15Val0DowngradeBlocked, Point15Val0StateBlocked
	case point15ValBRunFailed, point15ValBRunThrottled:
		return point15ValATriggerConnFail, point15Val0DowngradeReview, Point15Val0StateReviewRequired
	case point15ValBRunCompletedWithDowngrade:
		trigger := strings.TrimSpace(model.TriggerType)
		if trigger == "" {
			return "", "", Point15Val0StateBlocked
		}
		expectedOutcome := point15ValATriggerExpectedOutcome(trigger, model.TriggerIsDecisive, false)
		expectedState := point15ValATriggerExpectedState(trigger, model.TriggerIsDecisive, false)
		return trigger, expectedOutcome, expectedState
	}
	if strings.TrimSpace(model.ThrottleStatus) == point15ValBThrottleCrossTenantBlocked {
		return point15ValATriggerConnTenant, point15Val0DowngradeBlocked, Point15Val0StateBlocked
	}
	if strings.TrimSpace(model.ThrottleStatus) == point15ValBThrottleBlocked {
		if strings.TrimSpace(model.LastCompletedAt) == "" {
			return point15ValATriggerMissing, point15Val0DowngradeBlocked, Point15Val0StateBlocked
		}
		return point15ValATriggerStale, point15Val0DowngradeReview, Point15Val0StateReviewRequired
	}
	if strings.TrimSpace(model.ScheduleStatus) == point15ValBScheduleDue ||
		strings.TrimSpace(model.ScheduleStatus) == point15ValBScheduleOverdue ||
		strings.TrimSpace(model.ScheduleStatus) == point15ValBScheduleRunning ||
		strings.TrimSpace(model.ScheduleStatus) == point15ValBScheduleMissed ||
		strings.TrimSpace(model.ScheduleStatus) == point15ValBScheduleFailed ||
		strings.TrimSpace(model.ScheduleStatus) == point15ValBScheduleRetryPending ||
		strings.TrimSpace(model.ScheduleStatus) == point15ValBScheduleRetryExhausted ||
		strings.TrimSpace(model.ScheduleStatus) == point15ValBScheduleThrottled ||
		strings.TrimSpace(model.ScheduleStatus) == point15ValBScheduleBlocked ||
		strings.TrimSpace(model.RetryBudgetStatus) == point15ValBRetryExhausted ||
		strings.TrimSpace(model.RetryBudgetStatus) == point15ValBRetryBlocked ||
		strings.TrimSpace(model.ThrottleStatus) == point15ValBThrottleReviewRequired ||
		strings.TrimSpace(model.RunResult) == point15ValBRunMissed {
		if strings.TrimSpace(model.LastCompletedAt) == "" {
			if strings.TrimSpace(model.ScheduleStatus) == point15ValBScheduleRetryExhausted ||
				strings.TrimSpace(model.RetryBudgetStatus) == point15ValBRetryExhausted ||
				strings.TrimSpace(model.RetryBudgetStatus) == point15ValBRetryBlocked {
				return point15ValATriggerMissing, point15Val0DowngradeBlocked, Point15Val0StateBlocked
			}
			return point15ValATriggerMissing, point15Val0DowngradeIncomplete, Point15Val0StateIncomplete
		}
		return point15ValATriggerStale, point15Val0DowngradeReview, Point15Val0StateReviewRequired
	}
	return "", point15Val0DowngradeRetainActive, Point15Val0StateActive
}

func point15ValBRevalidationScheduleModel(dependency Point15ValBDependencySnapshot) Point15ValBRevalidationSchedule {
	val0 := dependency.Point15ValA.Dependency.Point15Val0
	return Point15ValBRevalidationSchedule{
		ScheduleID:          "schedule_point15_valb_001",
		EvidenceID:          val0.EvidenceContext.EvidenceID,
		TenantScope:         dependency.InheritedTenantScope,
		Required:            true,
		ScheduledStatus:     point15ValBScheduleScheduled,
		ScheduledAt:         "2026-05-07T09:00:00Z",
		RevalidationDueAt:   "2026-05-07T10:00:00Z",
		LastCompletedAt:     "2026-05-07T08:00:00Z",
		SchedulerTimeSource: point14Val0TimeSourceServerUTC,
		PolicyVersion:       val0.EvidenceContext.PolicyVersion,
		EngineVersion:       val0.EvidenceContext.EngineVersion,
		SchemaVersion:       val0.EvidenceContext.SchemaVersion,
	}
}

func EvaluatePoint15ValBRevalidationScheduleState(model Point15ValBRevalidationSchedule) string {
	if !point15ValBDependencyRefValid(model.ScheduleID) ||
		!point15ValBScheduleStatusValid(model.ScheduledStatus) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		(strings.TrimSpace(model.SchedulerTimeSource) != "" && !point14Val0CanonicalTimeSourceValid(model.SchedulerTimeSource)) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.EvidenceID) == "" || strings.TrimSpace(model.TenantScope) == "" || strings.TrimSpace(model.PolicyVersion) == "" || strings.TrimSpace(model.EngineVersion) == "" || strings.TrimSpace(model.SchemaVersion) == "" {
		return Point15ValBStateIncomplete
	}
	if !point15Val0EvidenceIDValid(model.EvidenceID) ||
		!point12Val0VersionIdentityValid(model.PolicyVersion) ||
		!point12Val0VersionIdentityValid(model.EngineVersion) ||
		!point12Val0VersionIdentityValid(model.SchemaVersion) {
		return Point15ValBStateBlocked
	}
	status := strings.TrimSpace(model.ScheduledStatus)
	if !model.Required {
		if status != point15ValBScheduleNotRequired {
			return Point15ValBStateBlocked
		}
		if strings.TrimSpace(model.RevalidationDueAt) != "" && !point14Val0ParsedTimeOk(model.RevalidationDueAt) {
			return Point15ValBStateBlocked
		}
		return Point15ValBStateActive
	}
	if status == point15ValBScheduleNotRequired {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.ScheduledAt) == "" || strings.TrimSpace(model.RevalidationDueAt) == "" || strings.TrimSpace(model.SchedulerTimeSource) == "" {
		return Point15ValBStateIncomplete
	}
	if !point14Val0ParsedTimeOk(model.ScheduledAt) || !point14Val0ParsedTimeOk(model.RevalidationDueAt) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.LastCompletedAt) != "" && !point14Val0ParsedTimeOk(model.LastCompletedAt) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.NextRetryAt) != "" && !point14Val0ParsedTimeOk(model.NextRetryAt) {
		return Point15ValBStateBlocked
	}
	if status == point15ValBScheduleCompleted && strings.TrimSpace(model.LastCompletedAt) == "" {
		return Point15ValBStateIncomplete
	}
	return point15ValBTargetStateToWaveState(point15ValBExpectedScheduleState(model.Required, model.ScheduledStatus))
}

func point15ValBRevalidationRunModel(schedule Point15ValBRevalidationSchedule, dependency Point15ValBDependencySnapshot) Point15ValBRevalidationRun {
	val0 := dependency.Point15ValA.Dependency.Point15Val0
	return Point15ValBRevalidationRun{
		RunResult:     point15ValBRunNotRun,
		EvidenceID:    val0.EvidenceContext.EvidenceID,
		TenantScope:   dependency.InheritedTenantScope,
		PolicyVersion: val0.EvidenceContext.PolicyVersion,
		EngineVersion: val0.EvidenceContext.EngineVersion,
		SchemaVersion: val0.EvidenceContext.SchemaVersion,
		ScheduleRef:   schedule.ScheduleID,
	}
}

func EvaluatePoint15ValBRevalidationRunState(model Point15ValBRevalidationRun) string {
	if !point15ValBRunResultValid(model.RunResult) ||
		(strings.TrimSpace(model.TenantScope) != "" && !point11Val0ScopeValid(model.TenantScope)) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.RunResult) == point15ValBRunNotRun {
		if strings.TrimSpace(model.RunID) != "" ||
			strings.TrimSpace(model.StartedAt) != "" ||
			strings.TrimSpace(model.CompletedAt) != "" ||
			strings.TrimSpace(model.RunEvidenceHash) != "" ||
			strings.TrimSpace(model.ConnectorResultRef) != "" ||
			strings.TrimSpace(model.DowngradeTriggerRef) != "" {
			return Point15ValBStateBlocked
		}
		return Point15ValBStateActive
	}
	if strings.TrimSpace(model.RunID) == "" ||
		strings.TrimSpace(model.ScheduleRef) == "" ||
		strings.TrimSpace(model.EvidenceID) == "" ||
		strings.TrimSpace(model.TenantScope) == "" ||
		strings.TrimSpace(model.StartedAt) == "" ||
		strings.TrimSpace(model.ConnectorResultRef) == "" {
		return Point15ValBStateIncomplete
	}
	if !point15ValBDependencyRefValid(model.RunID) ||
		!point15ValBDependencyRefValid(model.ScheduleRef) ||
		!point15Val0EvidenceIDValid(model.EvidenceID) ||
		!point15ValBConnectorResultRefValid(model.ConnectorResultRef) ||
		(strings.TrimSpace(model.DowngradeTriggerRef) != "" && !point15ValADependencyRefValid(model.DowngradeTriggerRef)) {
		return Point15ValBStateBlocked
	}
	if !point14Val0ParsedTimeOk(model.StartedAt) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.CompletedAt) == "" {
		return Point15ValBStateIncomplete
	}
	if !point14Val0ParsedTimeOk(model.CompletedAt) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.RunEvidenceHash) == "" {
		return Point15ValBStateIncomplete
	}
	if !point14Val0HashRefValid(model.RunEvidenceHash) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.PolicyVersion) == "" || strings.TrimSpace(model.EngineVersion) == "" || strings.TrimSpace(model.SchemaVersion) == "" {
		return Point15ValBStateIncomplete
	}
	if !point12Val0VersionIdentityValid(model.PolicyVersion) ||
		!point12Val0VersionIdentityValid(model.EngineVersion) ||
		!point12Val0VersionIdentityValid(model.SchemaVersion) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.RunResult) == point15ValBRunCompletedWithDowngrade && strings.TrimSpace(model.DowngradeTriggerRef) == "" {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.RunResult) == point15ValBRunCompletedClean && strings.TrimSpace(model.DowngradeTriggerRef) != "" {
		return Point15ValBStateBlocked
	}
	return point15ValBTargetStateToWaveState(point15ValBExpectedRunState(model.RunResult))
}

func point15ValBRetryBudgetModel(schedule Point15ValBRevalidationSchedule) Point15ValBRetryBudget {
	return Point15ValBRetryBudget{
		BudgetID:          "retry_budget_point15_valb_001",
		ScheduleRef:       schedule.ScheduleID,
		MaxRetries:        3,
		AttemptsUsed:      0,
		RetryBudgetStatus: point15ValBRetryAvailable,
	}
}

func EvaluatePoint15ValBRetryBudgetState(model Point15ValBRetryBudget) string {
	if !point15ValBDependencyRefValid(model.BudgetID) ||
		!point15ValBDependencyRefValid(model.ScheduleRef) ||
		!point15ValBRetryStatusValid(model.RetryBudgetStatus) ||
		!point15ValBRetryReasonValid(model.RetryReason) ||
		model.MaxRetries < 0 ||
		model.AttemptsUsed < 0 ||
		model.AttemptsUsed > model.MaxRetries ||
		model.SelfResetDetected {
		return Point15ValBStateBlocked
	}
	status := strings.TrimSpace(model.RetryBudgetStatus)
	if status == point15ValBRetryNotApplicable {
		if model.MaxRetries != 0 || model.AttemptsUsed != 0 || strings.TrimSpace(model.NextRetryAt) != "" || strings.TrimSpace(model.NextRetryTimeSource) != "" {
			return Point15ValBStateBlocked
		}
		return Point15ValBStateActive
	}
	if strings.TrimSpace(model.NextRetryAt) != "" {
		if !point14Val0ParsedTimeOk(model.NextRetryAt) || !point14Val0CanonicalTimeSourceValid(model.NextRetryTimeSource) {
			return Point15ValBStateBlocked
		}
	}
	switch status {
	case point15ValBRetryAvailable:
		if model.AttemptsUsed == model.MaxRetries && model.MaxRetries > 0 {
			return Point15ValBStateBlocked
		}
		return Point15ValBStateActive
	case point15ValBRetryExhausted:
		return point15ValBTargetStateToWaveState(point15ValBExpectedRetryState(model.RetryBudgetStatus, model.RetryReason))
	case point15ValBRetryBlocked:
		return Point15ValBStateBlocked
	default:
		return Point15ValBStateBlocked
	}
}

func point15ValBTenantThrottleModel(dependency Point15ValBDependencySnapshot) Point15ValBTenantThrottle {
	return Point15ValBTenantThrottle{
		ThrottleID:             "throttle_point15_valb_001",
		TenantScope:            dependency.InheritedTenantScope,
		RequestedRevalidations: 1,
		AllowedRevalidations:   10,
		ThrottleStatus:         point15ValBThrottleWithinLimit,
	}
}

func EvaluatePoint15ValBTenantThrottleState(model Point15ValBTenantThrottle) string {
	if !point15ValBDependencyRefValid(model.ThrottleID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point15ValBThrottleStatusValid(model.ThrottleStatus) ||
		model.RequestedRevalidations < 0 ||
		model.AllowedRevalidations < 0 {
		return Point15ValBStateBlocked
	}
	if model.CrossTenantDetected || strings.TrimSpace(model.ThrottleStatus) == point15ValBThrottleCrossTenantBlocked {
		return Point15ValBStateBlocked
	}
	switch strings.TrimSpace(model.ThrottleStatus) {
	case point15ValBThrottleNotApplicable:
		if model.RequestedRevalidations != 0 || model.AllowedRevalidations != 0 {
			return Point15ValBStateBlocked
		}
		return Point15ValBStateActive
	case point15ValBThrottleWithinLimit:
		if model.RequestedRevalidations > model.AllowedRevalidations {
			return Point15ValBStateBlocked
		}
		return Point15ValBStateActive
	case point15ValBThrottleReviewRequired:
		if model.RequestedRevalidations <= model.AllowedRevalidations {
			return Point15ValBStateBlocked
		}
		return Point15ValBStateReviewRequired
	case point15ValBThrottleBlocked:
		if model.RequestedRevalidations <= model.AllowedRevalidations {
			return Point15ValBStateBlocked
		}
		return Point15ValBStateBlocked
	default:
		return Point15ValBStateBlocked
	}
}

func point15ValBDowngradeBindingModel(schedule Point15ValBRevalidationSchedule, run Point15ValBRevalidationRun, retry Point15ValBRetryBudget, throttle Point15ValBTenantThrottle) Point15ValBDowngradeBinding {
	runRef := ""
	if strings.TrimSpace(run.RunResult) != point15ValBRunNotRun {
		runRef = run.RunID
	}
	return Point15ValBDowngradeBinding{
		BindingID:              "binding_point15_valb_001",
		ScheduleRef:            schedule.ScheduleID,
		RunRef:                 runRef,
		RetryBudgetRef:         retry.BudgetID,
		ThrottleRef:            throttle.ThrottleID,
		RequiredRevalidation:   schedule.Required,
		ScheduleStatus:         schedule.ScheduledStatus,
		RunResult:              run.RunResult,
		RetryBudgetStatus:      retry.RetryBudgetStatus,
		ThrottleStatus:         throttle.ThrottleStatus,
		LastCompletedAt:        schedule.LastCompletedAt,
		RunEvidenceHashMatches: true,
		TargetState:            Point15Val0StateActive,
		TargetDowngradeOutcome: point15Val0DowngradeRetainActive,
		RetainsActiveClosure:   true,
	}
}

func EvaluatePoint15ValBDowngradeBindingState(model Point15ValBDowngradeBinding) string {
	if !point15ValBDependencyRefValid(model.BindingID) ||
		!point15ValBDependencyRefValid(model.ScheduleRef) ||
		!point15ValBDependencyRefValid(model.RetryBudgetRef) ||
		!point15ValBDependencyRefValid(model.ThrottleRef) ||
		!point15ValBScheduleStatusValid(model.ScheduleStatus) ||
		!point15ValBRunResultValid(model.RunResult) ||
		!point15ValBRetryStatusValid(model.RetryBudgetStatus) ||
		!point15ValBThrottleStatusValid(model.ThrottleStatus) ||
		!point15Val0StateValid(model.TargetState) ||
		!point15Val0DowngradeOutcomeValid(model.TargetDowngradeOutcome) ||
		(strings.TrimSpace(model.TriggerType) != "" && !point15ValATriggerValid(model.TriggerType)) {
		return Point15ValBStateBlocked
	}
	runResult := strings.TrimSpace(model.RunResult)
	runRef := strings.TrimSpace(model.RunRef)
	if runResult == point15ValBRunNotRun {
		if runRef != "" {
			return Point15ValBStateBlocked
		}
	} else if !point15ValBDependencyRefValid(model.RunRef) {
		return Point15ValBStateBlocked
	}
	expectedTrigger, expectedOutcome, expectedState := point15ValBExpectedBinding(model)
	if expectedOutcome == "" {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.TargetState) != expectedState ||
		strings.TrimSpace(model.TargetDowngradeOutcome) != expectedOutcome {
		return Point15ValBStateBlocked
	}
	if expectedTrigger == "" {
		if strings.TrimSpace(model.TriggerType) != "" || model.RetainsPass || !model.RetainsActiveClosure {
			return Point15ValBStateBlocked
		}
		return Point15ValBStateActive
	}
	if strings.TrimSpace(model.TriggerType) != expectedTrigger || model.RetainsPass || model.RetainsActiveClosure {
		return Point15ValBStateBlocked
	}
	return point15ValBTargetStateToWaveState(model.TargetState)
}

func point15ValBTimestampDisciplineModel(schedule Point15ValBRevalidationSchedule, retry Point15ValBRetryBudget, dependency Point15ValBDependencySnapshot) Point15ValBTimestampDiscipline {
	return Point15ValBTimestampDiscipline{
		DisciplineID:           "timestamp_discipline_point15_valb_001",
		TenantScope:            dependency.InheritedTenantScope,
		ScheduledStatus:        schedule.ScheduledStatus,
		ScheduledAt:            schedule.ScheduledAt,
		ScheduledAtTimeSource:  point14Val0TimeSourceServerUTC,
		DueAt:                  schedule.RevalidationDueAt,
		DueAtTimeSource:        point14Val0TimeSourceServerUTC,
		NextRetryAt:            retry.NextRetryAt,
		NextRetryAtTimeSource:  retry.NextRetryTimeSource,
		ReferenceNow:           "2026-05-07T09:30:00Z",
		ReferenceNowTimeSource: point14Val0TimeSourceServerUTC,
		SourceEventAt:          "2026-05-07T08:55:00Z",
		SourceEventTimeSource:  point14Val0TimeSourceApprovedCustomerTime,
	}
}

func EvaluatePoint15ValBTimestampDisciplineState(model Point15ValBTimestampDiscipline) string {
	if !point15ValBDependencyRefValid(model.DisciplineID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point15ValBScheduleStatusValid(model.ScheduledStatus) ||
		!point14Val0ParsedTimeOk(model.ReferenceNow) ||
		!point14Val0CanonicalTimeSourceValid(model.ReferenceNowTimeSource) {
		return Point15ValBStateBlocked
	}
	if model.ClientLocalCreatesCanonical || model.SourceEventCreatesCanonical {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.ScheduledAt) == "" || strings.TrimSpace(model.ScheduledAtTimeSource) == "" {
		return Point15ValBStateIncomplete
	}
	if !point14Val0ParsedTimeOk(model.ScheduledAt) || !point14Val0CanonicalTimeSourceValid(model.ScheduledAtTimeSource) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.DueAt) == "" {
		return Point15ValBStateIncomplete
	}
	if !point14Val0ParsedTimeOk(model.DueAt) || !point14Val0CanonicalTimeSourceValid(model.DueAtTimeSource) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.StartedAt) != "" {
		if !point14Val0ParsedTimeOk(model.StartedAt) || !point14Val0CanonicalTimeSourceValid(model.StartedAtTimeSource) {
			return Point15ValBStateBlocked
		}
	}
	if strings.TrimSpace(model.CompletedAt) != "" {
		if !point14Val0ParsedTimeOk(model.CompletedAt) || !point14Val0CanonicalTimeSourceValid(model.CompletedAtTimeSource) {
			return Point15ValBStateBlocked
		}
	}
	if strings.TrimSpace(model.NextRetryAt) != "" {
		if !point14Val0ParsedTimeOk(model.NextRetryAt) || !point14Val0CanonicalTimeSourceValid(model.NextRetryAtTimeSource) {
			return Point15ValBStateBlocked
		}
	}
	if strings.TrimSpace(model.SourceEventAt) != "" {
		if !point14Val0ParsedTimeOk(model.SourceEventAt) || !point14Val0TimeSourceValid(model.SourceEventTimeSource) {
			return Point15ValBStateBlocked
		}
	}
	referenceNow, _ := point14Val0ParsedTime(model.ReferenceNow)
	scheduledAt, _ := point14Val0ParsedTime(model.ScheduledAt)
	dueAt, _ := point14Val0ParsedTime(model.DueAt)
	if scheduledAt.After(dueAt) || scheduledAt.After(referenceNow) {
		return Point15ValBStateBlocked
	}
	status := strings.TrimSpace(model.ScheduledStatus)
	if status == point15ValBScheduleScheduled && !dueAt.After(referenceNow) {
		return Point15ValBStateBlocked
	}
	if (status == point15ValBScheduleDue ||
		status == point15ValBScheduleOverdue ||
		status == point15ValBScheduleRunning ||
		status == point15ValBScheduleMissed ||
		status == point15ValBScheduleFailed ||
		status == point15ValBScheduleRetryPending ||
		status == point15ValBScheduleRetryExhausted ||
		status == point15ValBScheduleThrottled ||
		status == point15ValBScheduleBlocked) && dueAt.After(referenceNow) {
		return Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.NextRetryAt) != "" {
		nextRetryAt, _ := point14Val0ParsedTime(model.NextRetryAt)
		if dueAt.After(nextRetryAt) {
			return Point15ValBStateBlocked
		}
	}
	if strings.TrimSpace(model.StartedAt) != "" {
		startedAt, _ := point14Val0ParsedTime(model.StartedAt)
		if startedAt.Before(scheduledAt) {
			return Point15ValBStateReviewRequired
		}
		if startedAt.After(referenceNow) {
			return Point15ValBStateBlocked
		}
		if strings.TrimSpace(model.CompletedAt) != "" {
			completedAt, _ := point14Val0ParsedTime(model.CompletedAt)
			if completedAt.Before(startedAt) {
				return Point15ValBStateBlocked
			}
			if completedAt.After(referenceNow) {
				return Point15ValBStateBlocked
			}
			if completedAt.Before(scheduledAt) {
				return Point15ValBStateReviewRequired
			}
		}
	}
	return Point15ValBStateActive
}

func point15ValBAuthorityBoundaryModel(dependency Point15ValBDependencySnapshot) Point15ValBAuthorityBoundary {
	return Point15ValBAuthorityBoundary{
		BoundaryID:                      "authority_boundary_point15_valb_001",
		TenantScope:                     dependency.InheritedTenantScope,
		ExternalSourceInputOnly:         true,
		FormalEvaluatorOnly:             true,
		AgentRecommendationAdvisoryOnly: true,
	}
}

func EvaluatePoint15ValBAuthorityBoundaryState(model Point15ValBAuthorityBoundary) string {
	if !point15ValBDependencyRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!model.ExternalSourceInputOnly ||
		!model.FormalEvaluatorOnly ||
		!model.AgentRecommendationAdvisoryOnly {
		return Point15ValBStateBlocked
	}
	if model.SchedulerMarksEvidenceFresh ||
		model.SchedulerCreatesRevalidationTruth ||
		model.ConnectorRestoresActiveClosure ||
		model.DashboardSuppressesOverdueStatus ||
		model.PortalProjectionMutatesRevalidation ||
		model.CustomerProjectionMutatesRevalidation ||
		model.AuditorProjectionMutatesRevalidation ||
		model.AgentSatisfiesRevalidation ||
		model.RetryBudgetResetAllowed ||
		model.CanonicalMutationAllowed ||
		model.ProductionMutationAllowed ||
		model.PassAllowed {
		return Point15ValBStateBlocked
	}
	return Point15ValBStateActive
}

func point15ValBNoOverclaimGuardModel() Point15ValBNoOverclaimGuard {
	return Point15ValBNoOverclaimGuard{
		ObservedTexts: []string{
			"scheduled revalidation requires formal evaluator",
			"bounded revalidation decision",
		},
		AllowedSafeWording:     point15ValBSafeWording(),
		BlockedWording:         point15ValBForbiddenWording(),
		RevalidationDisclaimer: point15ValBRevalidationDisclaimer,
	}
}

func EvaluatePoint15ValBNoOverclaimGuardState(model Point15ValBNoOverclaimGuard) string {
	if strings.TrimSpace(model.RevalidationDisclaimer) != point15ValBRevalidationDisclaimer ||
		!point12Val0ExactStringSetMatch(model.AllowedSafeWording, point15ValBSafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point15ValBForbiddenWording()) {
		return Point15ValBStateBlocked
	}
	if point15ValBObservedListContainsForbiddenWording(model.ObservedTexts) {
		return Point15ValBStateBlocked
	}
	if point15ValBObservedListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point15ValBStateBlocked
	}
	return Point15ValBStateActive
}

func point15ValBFoundationModelFromUpstream(valA Point15ValADowngradeTriggerFoundation) Point15ValBScheduledRevalidationFoundation {
	dependency := point15ValBDependencySnapshotFromUpstream(valA)
	schedule := point15ValBRevalidationScheduleModel(dependency)
	run := point15ValBRevalidationRunModel(schedule, dependency)
	retry := point15ValBRetryBudgetModel(schedule)
	throttle := point15ValBTenantThrottleModel(dependency)
	return Point15ValBScheduledRevalidationFoundation{
		RevalidationDisclaimer: point15ValBRevalidationDisclaimer,
		Dependency:             dependency,
		Schedule:               schedule,
		Run:                    run,
		RetryBudget:            retry,
		TenantThrottle:         throttle,
		DowngradeBinding:       point15ValBDowngradeBindingModel(schedule, run, retry, throttle),
		TimestampDiscipline:    point15ValBTimestampDisciplineModel(schedule, retry, dependency),
		AuthorityBoundary:      point15ValBAuthorityBoundaryModel(dependency),
		NoOverclaimGuard:       point15ValBNoOverclaimGuardModel(),
	}
}

func Point15ValBFoundationModel() Point15ValBScheduledRevalidationFoundation {
	valA := ComputePoint15ValADowngradeTriggerFoundation(Point15ValAFoundationModel())
	return point15ValBFoundationModelFromUpstream(valA)
}

func point15ValBAggregate(states ...string) string {
	for _, state := range states {
		if strings.TrimSpace(state) == Point15ValBStateBlocked {
			return Point15ValBStateBlocked
		}
	}
	for _, state := range states {
		if strings.TrimSpace(state) == Point15ValBStateReviewRequired {
			return Point15ValBStateReviewRequired
		}
	}
	for _, state := range states {
		if strings.TrimSpace(state) == Point15ValBStateIncomplete {
			return Point15ValBStateIncomplete
		}
	}
	return Point15ValBStateActive
}

func point15ValBBlockingReasons(model Point15ValBScheduledRevalidationFoundation) []string {
	componentStates := map[string]string{
		"dependency":           model.DependencyState,
		"schedule":             model.ScheduleState,
		"run":                  model.RunState,
		"retry_budget":         model.RetryBudgetState,
		"tenant_throttle":      model.TenantThrottleState,
		"downgrade_binding":    model.DowngradeBindingState,
		"timestamp_discipline": model.TimestampDisciplineState,
		"authority_boundary":   model.AuthorityBoundaryState,
		"no_overclaim":         model.NoOverclaimState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if strings.TrimSpace(state) == Point15ValBStateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point15ValBReviewPrerequisites(model Point15ValBScheduledRevalidationFoundation) []string {
	componentStates := map[string]string{
		"schedule":             model.ScheduleState,
		"run":                  model.RunState,
		"retry_budget":         model.RetryBudgetState,
		"tenant_throttle":      model.TenantThrottleState,
		"downgrade_binding":    model.DowngradeBindingState,
		"timestamp_discipline": model.TimestampDisciplineState,
		"authority_boundary":   model.AuthorityBoundaryState,
		"no_overclaim":         model.NoOverclaimState,
	}
	prereqs := append([]string{}, model.Dependency.ReviewPrerequisites...)
	for name, state := range componentStates {
		if strings.TrimSpace(state) == Point15ValBStateReviewRequired || strings.TrimSpace(state) == Point15ValBStateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func ComputePoint15ValBScheduledRevalidationFoundation(model Point15ValBScheduledRevalidationFoundation) Point15ValBScheduledRevalidationFoundation {
	model.DependencyState = EvaluatePoint15ValBDependencyState(model.Dependency)
	model.ScheduleState = EvaluatePoint15ValBRevalidationScheduleState(model.Schedule)
	model.RunState = EvaluatePoint15ValBRevalidationRunState(model.Run)
	model.RetryBudgetState = EvaluatePoint15ValBRetryBudgetState(model.RetryBudget)
	model.TenantThrottleState = EvaluatePoint15ValBTenantThrottleState(model.TenantThrottle)
	model.DowngradeBindingState = EvaluatePoint15ValBDowngradeBindingState(model.DowngradeBinding)
	model.TimestampDisciplineState = EvaluatePoint15ValBTimestampDisciplineState(model.TimestampDiscipline)
	model.AuthorityBoundaryState = EvaluatePoint15ValBAuthorityBoundaryState(model.AuthorityBoundary)
	model.NoOverclaimState = EvaluatePoint15ValBNoOverclaimGuardState(model.NoOverclaimGuard)

	expectedTenant := strings.TrimSpace(model.Dependency.InheritedTenantScope)
	val0 := model.Dependency.Point15ValA.Dependency.Point15Val0
	expectedEvidenceID := strings.TrimSpace(val0.EvidenceContext.EvidenceID)
	expectedHash := strings.TrimSpace(val0.EvidenceContext.EvidenceHash)
	expectedPolicy := strings.TrimSpace(val0.EvidenceContext.PolicyVersion)
	expectedEngine := strings.TrimSpace(val0.EvidenceContext.EngineVersion)
	expectedSchema := strings.TrimSpace(val0.EvidenceContext.SchemaVersion)

	if strings.TrimSpace(model.Schedule.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.Run.TenantScope) != "" && strings.TrimSpace(model.Run.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.TenantThrottle.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.TimestampDiscipline.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.AuthorityBoundary.TenantScope) != expectedTenant {
		model.ScheduleState = Point15ValBStateBlocked
		model.RunState = Point15ValBStateBlocked
		model.TenantThrottleState = Point15ValBStateBlocked
		model.TimestampDisciplineState = Point15ValBStateBlocked
		model.AuthorityBoundaryState = Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.Schedule.EvidenceID) != expectedEvidenceID ||
		(strings.TrimSpace(model.Run.EvidenceID) != "" && strings.TrimSpace(model.Run.EvidenceID) != expectedEvidenceID) {
		model.ScheduleState = Point15ValBStateBlocked
		model.RunState = Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.Schedule.PolicyVersion) != expectedPolicy ||
		strings.TrimSpace(model.Schedule.EngineVersion) != expectedEngine ||
		strings.TrimSpace(model.Schedule.SchemaVersion) != expectedSchema ||
		(strings.TrimSpace(model.Run.PolicyVersion) != "" && strings.TrimSpace(model.Run.PolicyVersion) != expectedPolicy) ||
		(strings.TrimSpace(model.Run.EngineVersion) != "" && strings.TrimSpace(model.Run.EngineVersion) != expectedEngine) ||
		(strings.TrimSpace(model.Run.SchemaVersion) != "" && strings.TrimSpace(model.Run.SchemaVersion) != expectedSchema) {
		model.ScheduleState = Point15ValBStateBlocked
		model.RunState = Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.Run.RunResult) == point15ValBRunCompletedClean || strings.TrimSpace(model.Run.RunResult) == point15ValBRunCompletedWithDowngrade {
		if strings.TrimSpace(model.Run.RunEvidenceHash) != expectedHash {
			model.RunState = Point15ValBStateBlocked
			model.DowngradeBindingState = Point15ValBStateBlocked
		}
	}
	if strings.TrimSpace(model.Schedule.ScheduleID) != strings.TrimSpace(model.Run.ScheduleRef) && strings.TrimSpace(model.Run.RunResult) != point15ValBRunNotRun {
		model.RunState = Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.Schedule.ScheduleID) != strings.TrimSpace(model.RetryBudget.ScheduleRef) {
		model.RetryBudgetState = Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.Schedule.ScheduleID) != strings.TrimSpace(model.DowngradeBinding.ScheduleRef) ||
		(strings.TrimSpace(model.Run.RunResult) == point15ValBRunNotRun && strings.TrimSpace(model.DowngradeBinding.RunRef) != "") ||
		(strings.TrimSpace(model.Run.RunResult) != point15ValBRunNotRun && strings.TrimSpace(model.DowngradeBinding.RunRef) != strings.TrimSpace(model.Run.RunID)) ||
		strings.TrimSpace(model.DowngradeBinding.RetryBudgetRef) != strings.TrimSpace(model.RetryBudget.BudgetID) ||
		strings.TrimSpace(model.DowngradeBinding.ThrottleRef) != strings.TrimSpace(model.TenantThrottle.ThrottleID) ||
		model.DowngradeBinding.RequiredRevalidation != model.Schedule.Required ||
		strings.TrimSpace(model.DowngradeBinding.ScheduleStatus) != strings.TrimSpace(model.Schedule.ScheduledStatus) ||
		strings.TrimSpace(model.DowngradeBinding.RunResult) != strings.TrimSpace(model.Run.RunResult) ||
		strings.TrimSpace(model.DowngradeBinding.RetryBudgetStatus) != strings.TrimSpace(model.RetryBudget.RetryBudgetStatus) ||
		strings.TrimSpace(model.DowngradeBinding.ThrottleStatus) != strings.TrimSpace(model.TenantThrottle.ThrottleStatus) ||
		strings.TrimSpace(model.DowngradeBinding.LastCompletedAt) != strings.TrimSpace(model.Schedule.LastCompletedAt) {
		model.DowngradeBindingState = Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.Schedule.ScheduledStatus) == point15ValBScheduleCompleted &&
		strings.TrimSpace(model.Run.RunResult) != point15ValBRunCompletedClean &&
		strings.TrimSpace(model.Run.RunResult) != point15ValBRunCompletedWithDowngrade {
		model.ScheduleState = Point15ValBStateBlocked
		model.RunState = Point15ValBStateBlocked
	}
	if (strings.TrimSpace(model.Run.RunResult) == point15ValBRunCompletedClean || strings.TrimSpace(model.Run.RunResult) == point15ValBRunCompletedWithDowngrade) &&
		strings.TrimSpace(model.Schedule.ScheduledStatus) != point15ValBScheduleCompleted {
		model.RunState = Point15ValBStateBlocked
		model.ScheduleState = Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.TimestampDiscipline.ScheduledStatus) != strings.TrimSpace(model.Schedule.ScheduledStatus) ||
		strings.TrimSpace(model.TimestampDiscipline.ScheduledAt) != strings.TrimSpace(model.Schedule.ScheduledAt) ||
		strings.TrimSpace(model.TimestampDiscipline.DueAt) != strings.TrimSpace(model.Schedule.RevalidationDueAt) ||
		strings.TrimSpace(model.TimestampDiscipline.NextRetryAt) != strings.TrimSpace(model.RetryBudget.NextRetryAt) ||
		strings.TrimSpace(model.TimestampDiscipline.StartedAt) != strings.TrimSpace(model.Run.StartedAt) ||
		strings.TrimSpace(model.TimestampDiscipline.CompletedAt) != strings.TrimSpace(model.Run.CompletedAt) {
		model.TimestampDisciplineState = Point15ValBStateBlocked
	}
	if strings.TrimSpace(model.Schedule.ScheduledStatus) == point15ValBScheduleDue ||
		strings.TrimSpace(model.Schedule.ScheduledStatus) == point15ValBScheduleOverdue ||
		strings.TrimSpace(model.Schedule.ScheduledStatus) == point15ValBScheduleMissed ||
		strings.TrimSpace(model.Schedule.ScheduledStatus) == point15ValBScheduleRetryPending ||
		strings.TrimSpace(model.Schedule.ScheduledStatus) == point15ValBScheduleRetryExhausted ||
		strings.TrimSpace(model.Schedule.ScheduledStatus) == point15ValBScheduleThrottled ||
		strings.TrimSpace(model.Schedule.ScheduledStatus) == point15ValBScheduleFailed ||
		strings.TrimSpace(model.RetryBudget.RetryBudgetStatus) == point15ValBRetryExhausted ||
		strings.TrimSpace(model.TenantThrottle.ThrottleStatus) == point15ValBThrottleReviewRequired ||
		strings.TrimSpace(model.TenantThrottle.ThrottleStatus) == point15ValBThrottleBlocked ||
		strings.TrimSpace(model.TenantThrottle.ThrottleStatus) == point15ValBThrottleCrossTenantBlocked {
		if strings.TrimSpace(model.DowngradeBinding.TargetState) == Point15Val0StateActive ||
			strings.TrimSpace(model.DowngradeBinding.TargetDowngradeOutcome) == point15Val0DowngradeRetainActive ||
			model.DowngradeBinding.RetainsActiveClosure {
			model.DowngradeBindingState = Point15ValBStateBlocked
		}
	}

	model.CurrentState = point15ValBAggregate(
		model.DependencyState,
		model.ScheduleState,
		model.RunState,
		model.RetryBudgetState,
		model.TenantThrottleState,
		model.DowngradeBindingState,
		model.TimestampDisciplineState,
		model.AuthorityBoundaryState,
		model.NoOverclaimState,
	)
	model.BlockingReasons = point15ValBBlockingReasons(model)
	model.ReviewPrerequisites = point15ValBReviewPrerequisites(model)
	return model
}

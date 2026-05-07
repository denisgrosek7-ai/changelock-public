package formal

import (
	"encoding/json"
	"sort"
	"strings"
)

const (
	Point15ValCStateActive         = "point15_valc_enforcement_boundary_active"
	Point15ValCStateBlocked        = "point15_valc_enforcement_boundary_blocked"
	Point15ValCStateReviewRequired = "point15_valc_enforcement_boundary_review_required"
	Point15ValCStateIncomplete     = "point15_valc_enforcement_boundary_incomplete"
)

const (
	point15ValCWaveID                      = "val_c"
	point15ValCEnforcementDisclaimer       = "formal_enforcement_boundary_only no_silent_pass_retention point15_valc"
	point15ValCBlockedPassToken            = "point_15_pass"
	point15ValCActionNone                  = "no_action_required"
	point15ValCActionReview                = "enforce_review_required"
	point15ValCActionBlocked               = "enforce_blocked"
	point15ValCActionIncomplete            = "enforce_incomplete"
	point15ValCActionQuarantine            = "quarantine_evidence"
	point15ValCActionPreserveBlock         = "preserve_history_and_block"
	point15ValCActionPreserveReview        = "preserve_history_and_review"
	point15ValCActionGovernanceReview      = "require_governance_review"
	point15ValCReasonExpired               = "expired_evidence"
	point15ValCReasonRevoked               = "revoked_evidence"
	point15ValCReasonSupersededNoLineage   = "superseded_without_lineage"
	point15ValCReasonSupersededWithLineage = "superseded_with_lineage_review"
	point15ValCReasonStale                 = "stale_evidence"
	point15ValCReasonDecisiveDrift         = "decisive_drift"
	point15ValCReasonNonDecisiveDrift      = "non_decisive_drift_review"
	point15ValCReasonMissing               = "missing_freshness_proof"
	point15ValCReasonUnsupported           = "unsupported_freshness_status"
	point15ValCReasonTampered              = "tampered_freshness_proof"
	point15ValCReasonHashMismatch          = "evidence_hash_mismatch"
	point15ValCReasonRevalidationMissed    = "revalidation_missed"
	point15ValCReasonRevalidationOverdue   = "revalidation_overdue"
	point15ValCReasonRevalidationFailed    = "revalidation_failed"
	point15ValCReasonRetryExhausted        = "retry_budget_exhausted"
	point15ValCReasonConnectorAuth         = "connector_unauthorized"
	point15ValCReasonConnectorTimeout      = "connector_timeout"
	point15ValCReasonConnectorTenant       = "connector_tenant_mismatch"
	point15ValCReasonTenantMismatch        = "tenant_scope_mismatch"
	point15ValCReasonTimestampUntrusted    = "timestamp_untrusted"
	point15ValCReasonImpossibleOrdering    = "impossible_ordering"
	point15ValCLifecycleActive             = "active_bound"
	point15ValCLifecycleExpired            = "expired"
	point15ValCLifecycleRevoked            = "revoked"
	point15ValCLifecycleSuperseded         = "superseded"
	point15ValCLifecycleStale              = "stale"
	point15ValCLifecycleDrifted            = "drifted"
	point15ValCLifecycleTampered           = "tampered"
	point15ValCLifecycleHashMismatch       = "hash_mismatch"
	point15ValCLifecycleRevalidationFailed = "revalidation_failed"
)

type Point15ValCDependencySnapshot struct {
	Point15ValBCurrentState          string                                     `json:"point15_valb_current_state"`
	Point15ValBDependencyState       string                                     `json:"point15_valb_dependency_state"`
	Point15ValBScheduleState         string                                     `json:"point15_valb_schedule_state"`
	Point15ValBRunState              string                                     `json:"point15_valb_run_state"`
	Point15ValBRetryBudgetState      string                                     `json:"point15_valb_retry_budget_state"`
	Point15ValBTenantThrottleState   string                                     `json:"point15_valb_tenant_throttle_state"`
	Point15ValBDowngradeBindingState string                                     `json:"point15_valb_downgrade_binding_state"`
	Point15ValBTimestampState        string                                     `json:"point15_valb_timestamp_state"`
	Point15ValBAuthorityState        string                                     `json:"point15_valb_authority_state"`
	Point15ValBNoOverclaimState      string                                     `json:"point15_valb_no_overclaim_state"`
	Point15ValBComputedFromUpstream  bool                                       `json:"point15_valb_computed_from_upstream"`
	Point15ValBMerged                bool                                       `json:"point15_valb_merged"`
	Point15ValBCIGreen               bool                                       `json:"point15_valb_ci_green"`
	Point15ValBReviewedOnMain        bool                                       `json:"point15_valb_reviewed_on_main"`
	Point15PassSeen                  bool                                       `json:"point15_pass_seen"`
	InheritedPoint15ValACurrentState string                                     `json:"inherited_point15_vala_current_state"`
	InheritedPoint15Val0CurrentState string                                     `json:"inherited_point15_val0_current_state"`
	InheritedPoint14ValECurrentState string                                     `json:"inherited_point14_vale_current_state"`
	InheritedTenantScope             string                                     `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput       bool                                       `json:"snapshot_from_computed_output"`
	ReviewPrerequisites              []string                                   `json:"review_prerequisites,omitempty"`
	Point15ValB                      Point15ValBScheduledRevalidationFoundation `json:"point15_valb"`
}

type Point15ValCEnforcementBoundaryFoundation struct {
	CurrentState             string                                `json:"current_state"`
	BlockingReasons          []string                              `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites      []string                              `json:"review_prerequisites,omitempty"`
	EnforcementDisclaimer    string                                `json:"enforcement_disclaimer"`
	DependencyState          string                                `json:"dependency_state"`
	EnforcementActionState   string                                `json:"enforcement_action_state"`
	EvidenceLifecycleState   string                                `json:"evidence_lifecycle_state"`
	RevocationBoundaryState  string                                `json:"revocation_boundary_state"`
	ExpiryBoundaryState      string                                `json:"expiry_boundary_state"`
	SupersessionState        string                                `json:"supersession_boundary_state"`
	ReplayProofHistoryState  string                                `json:"replay_proof_history_state"`
	TimestampDisciplineState string                                `json:"timestamp_discipline_state"`
	AuthorityBoundaryState   string                                `json:"authority_boundary_state"`
	TenantBoundaryState      string                                `json:"tenant_boundary_state"`
	NoOverclaimState         string                                `json:"no_overclaim_state"`
	Dependency               Point15ValCDependencySnapshot         `json:"dependency"`
	EnforcementAction        Point15ValCEnforcementAction          `json:"enforcement_action"`
	EvidenceLifecycle        Point15ValCEvidenceLifecycleBoundary  `json:"evidence_lifecycle"`
	RevocationBoundary       Point15ValCRevocationBoundary         `json:"revocation_boundary"`
	ExpiryBoundary           Point15ValCExpiryBoundary             `json:"expiry_boundary"`
	SupersessionBoundary     Point15ValCSupersessionBoundary       `json:"supersession_boundary"`
	ReplayProofHistory       Point15ValCReplayProofHistoryBoundary `json:"replay_proof_history"`
	TimestampDiscipline      Point15ValCTimestampDiscipline        `json:"timestamp_discipline"`
	TenantBoundary           Point15ValCTenantBoundary             `json:"tenant_boundary"`
	AuthorityBoundary        Point15ValCAuthorityBoundary          `json:"authority_boundary"`
	NoOverclaimGuard         Point15ValCNoOverclaimGuard           `json:"no_overclaim_guard"`
}

type Point15ValCEnforcementAction struct {
	EnforcementID         string `json:"enforcement_id"`
	EvidenceID            string `json:"evidence_id"`
	TenantScope           string `json:"tenant_scope"`
	EnforcementAction     string `json:"enforcement_action"`
	EnforcementReason     string `json:"enforcement_reason"`
	ReasonDecisive        bool   `json:"reason_decisive"`
	TargetState           string `json:"target_state"`
	DowngradeOutcome      string `json:"downgrade_outcome"`
	SourceValBTriggerRef  string `json:"source_valb_trigger_ref"`
	SourceValADecisionRef string `json:"source_vala_decision_ref"`
	RetainsPass           bool   `json:"retains_pass"`
	RetainsActiveClosure  bool   `json:"retains_active_closure"`
}

type Point15ValCEvidenceLifecycleBoundary struct {
	LifecycleID                string `json:"lifecycle_id"`
	EvidenceID                 string `json:"evidence_id"`
	TenantScope                string `json:"tenant_scope"`
	PreviousEvidenceRef        string `json:"previous_evidence_ref"`
	ReplacementEvidenceRef     string `json:"replacement_evidence_ref"`
	LifecycleStatus            string `json:"lifecycle_status"`
	LifecycleReason            string `json:"lifecycle_reason"`
	ReasonDecisive             bool   `json:"reason_decisive"`
	HistoryPreserved           bool   `json:"history_preserved"`
	LineageRef                 string `json:"lineage_ref"`
	CanonicalMutationAttempted bool   `json:"canonical_mutation_attempted"`
}

type Point15ValCRevocationBoundary struct {
	RevocationID             string `json:"revocation_id"`
	EvidenceID               string `json:"evidence_id"`
	TenantScope              string `json:"tenant_scope"`
	RevocationPresent        bool   `json:"revocation_present"`
	RevocationSourceRef      string `json:"revocation_source_ref"`
	RevocationReceivedAt     string `json:"revocation_received_at"`
	RevocationValidatedAt    string `json:"revocation_validated_at"`
	RevocationTimeSource     string `json:"revocation_time_source"`
	GovernanceReviewRequired bool   `json:"governance_review_required"`
	AutoRevoked              bool   `json:"auto_revoked"`
	AutoPublished            bool   `json:"auto_published"`
	HistoryPreserved         bool   `json:"history_preserved"`
	SourceAuthorityGranted   bool   `json:"source_authority_granted"`
}

type Point15ValCExpiryBoundary struct {
	ExpiryID               string `json:"expiry_id"`
	EvidenceID             string `json:"evidence_id"`
	TenantScope            string `json:"tenant_scope"`
	ExpiringEvidence       bool   `json:"expiring_evidence"`
	ExpiresAt              string `json:"expires_at"`
	EvaluatedAt            string `json:"evaluated_at"`
	ExpiryTimeSource       string `json:"expiry_time_source"`
	ExpiryEnforced         bool   `json:"expiry_enforced"`
	ExpiryHistoryPreserved bool   `json:"expiry_history_preserved"`
}

type Point15ValCSupersessionBoundary struct {
	SupersessionID            string `json:"supersession_id"`
	TenantScope               string `json:"tenant_scope"`
	SupersessionPresent       bool   `json:"supersession_present"`
	OldEvidenceRef            string `json:"old_evidence_ref"`
	NewEvidenceRef            string `json:"new_evidence_ref"`
	LineageRef                string `json:"lineage_ref"`
	ReplacementHash           string `json:"replacement_hash"`
	PriorHash                 string `json:"prior_hash"`
	HistoryPreserved          bool   `json:"history_preserved"`
	SilentReplacementDetected bool   `json:"silent_replacement_detected"`
	AutoPublished             bool   `json:"auto_published"`
	AutoApproved              bool   `json:"auto_approved"`
}

type Point15ValCReplayProofHistoryBoundary struct {
	HistoryID                   string `json:"history_id"`
	ProofHistoryRef             string `json:"proof_history_ref"`
	ReplayRef                   string `json:"replay_ref"`
	ProofPackRef                string `json:"proof_pack_ref"`
	DecisiveEvidenceVisible     bool   `json:"decisive_evidence_visible"`
	BlockedReasonVisible        bool   `json:"blocked_reason_visible"`
	PriorStateVisible           bool   `json:"prior_state_visible"`
	CurrentStateVisible         bool   `json:"current_state_visible"`
	ProjectionStrengthensClaims bool   `json:"projection_strengthens_claims"`
}

type Point15ValCTimestampDiscipline struct {
	DisciplineID                string `json:"discipline_id"`
	EvidenceID                  string `json:"evidence_id"`
	TenantScope                 string `json:"tenant_scope"`
	RevocationPresent           bool   `json:"revocation_present"`
	ExpiryEnforced              bool   `json:"expiry_enforced"`
	SupersessionPresent         bool   `json:"supersession_present"`
	SourceEventAt               string `json:"source_event_at"`
	SourceEventTimeSource       string `json:"source_event_time_source"`
	ReceivedAt                  string `json:"received_at"`
	ReceivedAtTimeSource        string `json:"received_at_time_source"`
	ValidatedAt                 string `json:"validated_at"`
	ValidatedAtTimeSource       string `json:"validated_at_time_source"`
	EnforcedAt                  string `json:"enforced_at"`
	EnforcedAtTimeSource        string `json:"enforced_at_time_source"`
	EvaluatedAt                 string `json:"evaluated_at"`
	EvaluatedAtTimeSource       string `json:"evaluated_at_time_source"`
	ReferenceNow                string `json:"reference_now"`
	ReferenceNowTimeSource      string `json:"reference_now_time_source"`
	ClientLocalCreatesCanonical bool   `json:"client_local_creates_canonical"`
	SourceEventCreatesCanonical bool   `json:"source_event_creates_canonical"`
}

type Point15ValCTenantBoundary struct {
	BoundaryID                   string `json:"boundary_id"`
	TenantScope                  string `json:"tenant_scope"`
	ReferencedTenantScope        string `json:"referenced_tenant_scope"`
	EnforcementResultTenantScope string `json:"enforcement_result_tenant_scope"`
	CrossTenantDetected          bool   `json:"cross_tenant_detected"`
}

type Point15ValCAuthorityBoundary struct {
	BoundaryID                           string `json:"boundary_id"`
	TenantScope                          string `json:"tenant_scope"`
	ExternalSourceInputOnly              bool   `json:"external_source_input_only"`
	FormalEvaluatorOnly                  bool   `json:"formal_evaluator_only"`
	AgentRecommendationAdvisoryOnly      bool   `json:"agent_recommendation_advisory_only"`
	SchedulerEnforcesBoundary            bool   `json:"scheduler_enforces_boundary"`
	ConnectorRestoresActiveClosure       bool   `json:"connector_restores_active_closure"`
	DashboardSuppressesEnforcement       bool   `json:"dashboard_suppresses_enforcement"`
	PortalProjectionMutatesEnforcement   bool   `json:"portal_projection_mutates_enforcement"`
	CustomerProjectionMutatesEnforcement bool   `json:"customer_projection_mutates_enforcement"`
	AuditorProjectionMutatesEnforcement  bool   `json:"auditor_projection_mutates_enforcement"`
	AgentSatisfiesEnforcement            bool   `json:"agent_satisfies_enforcement"`
	RevocationExecutionSideEffectAllowed bool   `json:"revocation_execution_side_effect_allowed"`
	AutomaticPublicationAllowed          bool   `json:"automatic_publication_allowed"`
	CanonicalMutationAllowed             bool   `json:"canonical_mutation_allowed"`
	ProductionMutationAllowed            bool   `json:"production_mutation_allowed"`
	PassAllowed                          bool   `json:"pass_allowed"`
}

type Point15ValCNoOverclaimGuard struct {
	ObservedTexts                        []string `json:"observed_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	EnforcementDisclaimer                string   `json:"enforcement_disclaimer"`
}

func point15ValCStates() []string {
	return []string{
		Point15ValCStateActive,
		Point15ValCStateBlocked,
		Point15ValCStateReviewRequired,
		Point15ValCStateIncomplete,
	}
}

func point15ValCStateValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValCStates())
}

func point15ValCActions() []string {
	return []string{
		point15ValCActionNone,
		point15ValCActionReview,
		point15ValCActionBlocked,
		point15ValCActionIncomplete,
		point15ValCActionQuarantine,
		point15ValCActionPreserveBlock,
		point15ValCActionPreserveReview,
		point15ValCActionGovernanceReview,
	}
}

func point15ValCActionValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValCActions())
}

func point15ValCReasons() []string {
	return []string{
		point15ValCReasonExpired,
		point15ValCReasonRevoked,
		point15ValCReasonSupersededNoLineage,
		point15ValCReasonSupersededWithLineage,
		point15ValCReasonStale,
		point15ValCReasonDecisiveDrift,
		point15ValCReasonNonDecisiveDrift,
		point15ValCReasonMissing,
		point15ValCReasonUnsupported,
		point15ValCReasonTampered,
		point15ValCReasonHashMismatch,
		point15ValCReasonRevalidationMissed,
		point15ValCReasonRevalidationOverdue,
		point15ValCReasonRevalidationFailed,
		point15ValCReasonRetryExhausted,
		point15ValCReasonConnectorAuth,
		point15ValCReasonConnectorTimeout,
		point15ValCReasonConnectorTenant,
		point15ValCReasonTenantMismatch,
		point15ValCReasonTimestampUntrusted,
		point15ValCReasonImpossibleOrdering,
	}
}

func point15ValCReasonValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValCReasons())
}

func point15ValCLifecycleStatuses() []string {
	return []string{
		point15ValCLifecycleActive,
		point15ValCLifecycleExpired,
		point15ValCLifecycleRevoked,
		point15ValCLifecycleSuperseded,
		point15ValCLifecycleStale,
		point15ValCLifecycleDrifted,
		point15ValCLifecycleTampered,
		point15ValCLifecycleHashMismatch,
		point15ValCLifecycleRevalidationFailed,
	}
}

func point15ValCLifecycleStatusValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValCLifecycleStatuses())
}

func point15ValCDependencyRefValid(value string) bool {
	return point14Val0RefValid(value,
		"point15_valc_",
		"point15_valb_",
		"point15_vala_",
		"point15_val0_",
		"binding_",
		"decision_",
		"lineage_",
		"history_",
		"history_boundary_",
		"replay_",
		"proof_",
		"proof_pack_",
		"revocation_",
		"revocation_source_",
		"expiry_",
		"supersession_",
		"enforcement_",
		"lifecycle_",
		"timestamp_discipline_",
		"tenant_boundary_",
		"authority_boundary_",
		"boundary_",
	)
}

func point15ValCForbiddenWording() []string {
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
		"guaranteed secure",
	}
}

func point15ValCSafeWording() []string {
	return []string{
		"downgrade trigger detected",
		"freshness requires review",
		"evidence support available for review",
		"bounded freshness decision",
		"trigger mapped by formal evaluator only",
		"enforcement preserves proof history",
	}
}

func point15ValCObservedTextContainsForbiddenWording(text string) bool {
	trimmed := strings.TrimSpace(strings.ToLower(text))
	if trimmed == "" {
		return false
	}
	for _, safe := range point15ValCSafeWording() {
		if trimmed == strings.ToLower(strings.TrimSpace(safe)) {
			return false
		}
	}
	for _, forbidden := range point15ValCForbiddenWording() {
		if strings.Contains(trimmed, strings.ToLower(strings.TrimSpace(forbidden))) {
			return true
		}
	}
	return false
}

func point15ValCObservedListContainsForbiddenWording(values []string) bool {
	for _, value := range values {
		if point15ValCObservedTextContainsForbiddenWording(value) {
			return true
		}
	}
	return false
}

func point15ValCValBPayloadContainsPoint15Pass(valB Point15ValBScheduledRevalidationFoundation) bool {
	payload, err := json.Marshal(valB)
	if err != nil {
		return true
	}
	text := strings.ToLower(string(payload))
	return strings.Contains(text, `"point_15_pass"`) || strings.Contains(text, `"point15_pass"`)
}

func point15ValCExpectedAction(reason string, decisive bool) (string, string, string) {
	switch strings.TrimSpace(reason) {
	case "":
		return point15ValCActionNone, Point15Val0StateActive, point15Val0DowngradeRetainActive
	case point15ValCReasonExpired, point15ValCReasonRevoked, point15ValCReasonUnsupported, point15ValCReasonRetryExhausted,
		point15ValCReasonConnectorAuth, point15ValCReasonConnectorTenant, point15ValCReasonTenantMismatch,
		point15ValCReasonTimestampUntrusted, point15ValCReasonImpossibleOrdering, point15ValCReasonDecisiveDrift:
		return point15ValCActionBlocked, Point15Val0StateBlocked, point15Val0DowngradeBlocked
	case point15ValCReasonStale, point15ValCReasonNonDecisiveDrift, point15ValCReasonRevalidationMissed,
		point15ValCReasonRevalidationOverdue, point15ValCReasonRevalidationFailed, point15ValCReasonConnectorTimeout:
		return point15ValCActionReview, Point15Val0StateReviewRequired, point15Val0DowngradeReview
	case point15ValCReasonMissing:
		if decisive {
			return point15ValCActionBlocked, Point15Val0StateBlocked, point15Val0DowngradeBlocked
		}
		return point15ValCActionIncomplete, Point15Val0StateIncomplete, point15Val0DowngradeIncomplete
	case point15ValCReasonTampered, point15ValCReasonHashMismatch:
		return point15ValCActionQuarantine, Point15Val0StateBlocked, point15Val0DowngradeBlocked
	case point15ValCReasonSupersededNoLineage:
		return point15ValCActionPreserveBlock, Point15Val0StateBlocked, point15Val0DowngradeBlocked
	case point15ValCReasonSupersededWithLineage:
		return point15ValCActionPreserveReview, Point15Val0StateReviewRequired, point15Val0DowngradeReview
	default:
		return "", "", ""
	}
}

func point15ValCTargetStateToWaveState(target string) string {
	switch strings.TrimSpace(target) {
	case Point15Val0StateActive:
		return Point15ValCStateActive
	case Point15Val0StateBlocked:
		return Point15ValCStateBlocked
	case Point15Val0StateReviewRequired:
		return Point15ValCStateReviewRequired
	case Point15Val0StateIncomplete:
		return Point15ValCStateIncomplete
	default:
		return Point15ValCStateBlocked
	}
}

func point15ValCDependencyModel() Point15ValCDependencySnapshot {
	valB := ComputePoint15ValBScheduledRevalidationFoundation(Point15ValBFoundationModel())
	return Point15ValCDependencySnapshot{
		Point15ValBCurrentState:          valB.CurrentState,
		Point15ValBDependencyState:       valB.DependencyState,
		Point15ValBScheduleState:         valB.ScheduleState,
		Point15ValBRunState:              valB.RunState,
		Point15ValBRetryBudgetState:      valB.RetryBudgetState,
		Point15ValBTenantThrottleState:   valB.TenantThrottleState,
		Point15ValBDowngradeBindingState: valB.DowngradeBindingState,
		Point15ValBTimestampState:        valB.TimestampDisciplineState,
		Point15ValBAuthorityState:        valB.AuthorityBoundaryState,
		Point15ValBNoOverclaimState:      valB.NoOverclaimState,
		Point15ValBComputedFromUpstream:  valB.Dependency.SnapshotFromComputedOutput,
		Point15ValBMerged:                true,
		Point15ValBCIGreen:               true,
		Point15ValBReviewedOnMain:        true,
		Point15PassSeen:                  false,
		InheritedPoint15ValACurrentState: valB.Dependency.Point15ValA.CurrentState,
		InheritedPoint15Val0CurrentState: valB.Dependency.Point15ValA.Dependency.Point15Val0.CurrentState,
		InheritedPoint14ValECurrentState: valB.Dependency.Point15ValA.Dependency.Point15Val0.Dependency.Point14ValE.CurrentState,
		InheritedTenantScope:             valB.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:       true,
		Point15ValB:                      valB,
	}
}

func EvaluatePoint15ValCDependencyState(model Point15ValCDependencySnapshot) string {
	if !point15ValBStateValid(model.Point15ValBCurrentState) ||
		!point15ValBStateValid(model.Point15ValBDependencyState) ||
		!point15ValBStateValid(model.Point15ValBScheduleState) ||
		!point15ValBStateValid(model.Point15ValBRunState) ||
		!point15ValBStateValid(model.Point15ValBRetryBudgetState) ||
		!point15ValBStateValid(model.Point15ValBTenantThrottleState) ||
		!point15ValBStateValid(model.Point15ValBDowngradeBindingState) ||
		!point15ValBStateValid(model.Point15ValBTimestampState) ||
		!point15ValBStateValid(model.Point15ValBAuthorityState) ||
		!point15ValBStateValid(model.Point15ValBNoOverclaimState) ||
		!point15ValAStateValid(model.InheritedPoint15ValACurrentState) ||
		!point15Val0StateValid(model.InheritedPoint15Val0CurrentState) ||
		strings.TrimSpace(model.InheritedPoint14ValECurrentState) != Point14ValEStatePassConfirmed ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return Point15ValCStateBlocked
	}
	if model.Point15PassSeen || point15ValCValBPayloadContainsPoint15Pass(model.Point15ValB) {
		return Point15ValCStateBlocked
	}
	if !model.Point15ValBMerged || !model.Point15ValBCIGreen || !model.Point15ValBReviewedOnMain ||
		!model.Point15ValBComputedFromUpstream || !model.SnapshotFromComputedOutput {
		return Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.Point15ValBCurrentState) != Point15ValBStateActive ||
		strings.TrimSpace(model.Point15ValBDependencyState) != Point15ValBStateActive ||
		strings.TrimSpace(model.Point15ValBScheduleState) != Point15ValBStateActive ||
		strings.TrimSpace(model.Point15ValBRunState) != Point15ValBStateActive ||
		strings.TrimSpace(model.Point15ValBRetryBudgetState) != Point15ValBStateActive ||
		strings.TrimSpace(model.Point15ValBTenantThrottleState) != Point15ValBStateActive ||
		strings.TrimSpace(model.Point15ValBDowngradeBindingState) != Point15ValBStateActive ||
		strings.TrimSpace(model.Point15ValBTimestampState) != Point15ValBStateActive ||
		strings.TrimSpace(model.Point15ValBAuthorityState) != Point15ValBStateActive ||
		strings.TrimSpace(model.Point15ValBNoOverclaimState) != Point15ValBStateActive ||
		strings.TrimSpace(model.InheritedPoint15ValACurrentState) != Point15ValAStateActive ||
		strings.TrimSpace(model.InheritedPoint15Val0CurrentState) != Point15Val0StateActive {
		return Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.Point15ValB.CurrentState) != strings.TrimSpace(model.Point15ValBCurrentState) ||
		strings.TrimSpace(model.Point15ValB.DependencyState) != strings.TrimSpace(model.Point15ValBDependencyState) ||
		strings.TrimSpace(model.Point15ValB.ScheduleState) != strings.TrimSpace(model.Point15ValBScheduleState) ||
		strings.TrimSpace(model.Point15ValB.RunState) != strings.TrimSpace(model.Point15ValBRunState) ||
		strings.TrimSpace(model.Point15ValB.RetryBudgetState) != strings.TrimSpace(model.Point15ValBRetryBudgetState) ||
		strings.TrimSpace(model.Point15ValB.TenantThrottleState) != strings.TrimSpace(model.Point15ValBTenantThrottleState) ||
		strings.TrimSpace(model.Point15ValB.DowngradeBindingState) != strings.TrimSpace(model.Point15ValBDowngradeBindingState) ||
		strings.TrimSpace(model.Point15ValB.TimestampDisciplineState) != strings.TrimSpace(model.Point15ValBTimestampState) ||
		strings.TrimSpace(model.Point15ValB.AuthorityBoundaryState) != strings.TrimSpace(model.Point15ValBAuthorityState) ||
		strings.TrimSpace(model.Point15ValB.NoOverclaimState) != strings.TrimSpace(model.Point15ValBNoOverclaimState) ||
		strings.TrimSpace(model.Point15ValB.Dependency.InheritedTenantScope) != strings.TrimSpace(model.InheritedTenantScope) ||
		model.Point15ValBComputedFromUpstream != model.Point15ValB.Dependency.SnapshotFromComputedOutput {
		return Point15ValCStateBlocked
	}
	return Point15ValCStateActive
}

func point15ValCEnforcementActionModel(dependency Point15ValCDependencySnapshot) Point15ValCEnforcementAction {
	evidenceID := dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.EvidenceID
	return Point15ValCEnforcementAction{
		EnforcementID:         "enforcement_point15_valc_001",
		EvidenceID:            evidenceID,
		TenantScope:           dependency.InheritedTenantScope,
		EnforcementAction:     point15ValCActionNone,
		TargetState:           Point15Val0StateActive,
		DowngradeOutcome:      point15Val0DowngradeRetainActive,
		SourceValBTriggerRef:  dependency.Point15ValB.DowngradeBinding.BindingID,
		SourceValADecisionRef: dependency.Point15ValB.Dependency.Point15ValA.Decision.DecisionID,
		RetainsPass:           false,
		RetainsActiveClosure:  true,
	}
}

func EvaluatePoint15ValCEnforcementActionState(model Point15ValCEnforcementAction) string {
	if !point15ValCDependencyRefValid(model.EnforcementID) ||
		!point15ValCDependencyRefValid(model.SourceValBTriggerRef) ||
		!point15ValCDependencyRefValid(model.SourceValADecisionRef) ||
		!point14Val0RefValid(model.EvidenceID, "evidence_") ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point15ValCActionValid(model.EnforcementAction) ||
		!point15Val0StateValid(model.TargetState) ||
		!point15Val0DowngradeOutcomeValid(model.DowngradeOutcome) {
		return Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.EnforcementAction) == point15ValCActionNone {
		if strings.TrimSpace(model.EnforcementReason) != "" ||
			model.ReasonDecisive ||
			strings.TrimSpace(model.TargetState) != Point15Val0StateActive ||
			strings.TrimSpace(model.DowngradeOutcome) != point15Val0DowngradeRetainActive ||
			model.RetainsPass ||
			!model.RetainsActiveClosure {
			return Point15ValCStateBlocked
		}
		return Point15ValCStateActive
	}
	if !point15ValCReasonValid(model.EnforcementReason) {
		return Point15ValCStateBlocked
	}
	expectedAction, expectedState, expectedOutcome := point15ValCExpectedAction(model.EnforcementReason, model.ReasonDecisive)
	if expectedAction == "" ||
		strings.TrimSpace(model.EnforcementAction) != expectedAction ||
		strings.TrimSpace(model.TargetState) != expectedState ||
		strings.TrimSpace(model.DowngradeOutcome) != expectedOutcome ||
		model.RetainsPass ||
		model.RetainsActiveClosure {
		return Point15ValCStateBlocked
	}
	return point15ValCTargetStateToWaveState(model.TargetState)
}

func point15ValCEvidenceLifecycleStatusesReasonValid(status, reason string) bool {
	switch strings.TrimSpace(status) {
	case point15ValCLifecycleActive:
		return strings.TrimSpace(reason) == ""
	case point15ValCLifecycleExpired:
		return strings.TrimSpace(reason) == point15ValCReasonExpired
	case point15ValCLifecycleRevoked:
		return strings.TrimSpace(reason) == point15ValCReasonRevoked
	case point15ValCLifecycleSuperseded:
		return strings.TrimSpace(reason) == point15ValCReasonSupersededNoLineage || strings.TrimSpace(reason) == point15ValCReasonSupersededWithLineage
	case point15ValCLifecycleStale:
		return strings.TrimSpace(reason) == point15ValCReasonStale || strings.TrimSpace(reason) == point15ValCReasonRevalidationMissed || strings.TrimSpace(reason) == point15ValCReasonRevalidationOverdue || strings.TrimSpace(reason) == point15ValCReasonRevalidationFailed
	case point15ValCLifecycleDrifted:
		return strings.TrimSpace(reason) == point15ValCReasonDecisiveDrift || strings.TrimSpace(reason) == point15ValCReasonNonDecisiveDrift
	case point15ValCLifecycleTampered:
		return strings.TrimSpace(reason) == point15ValCReasonTampered
	case point15ValCLifecycleHashMismatch:
		return strings.TrimSpace(reason) == point15ValCReasonHashMismatch
	case point15ValCLifecycleRevalidationFailed:
		return strings.TrimSpace(reason) == point15ValCReasonRevalidationFailed || strings.TrimSpace(reason) == point15ValCReasonRetryExhausted
	default:
		return false
	}
}

func point15ValCEvidenceLifecycleBoundaryModel(dependency Point15ValCDependencySnapshot) Point15ValCEvidenceLifecycleBoundary {
	return Point15ValCEvidenceLifecycleBoundary{
		LifecycleID:      "lifecycle_point15_valc_001",
		EvidenceID:       dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.EvidenceID,
		TenantScope:      dependency.InheritedTenantScope,
		LifecycleStatus:  point15ValCLifecycleActive,
		HistoryPreserved: true,
	}
}

func EvaluatePoint15ValCEvidenceLifecycleBoundaryState(model Point15ValCEvidenceLifecycleBoundary) string {
	if !point15ValCDependencyRefValid(model.LifecycleID) ||
		!point14Val0RefValid(model.EvidenceID, "evidence_") ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!point15ValCLifecycleStatusValid(model.LifecycleStatus) ||
		(strings.TrimSpace(model.PreviousEvidenceRef) != "" && !point14Val0RefValid(model.PreviousEvidenceRef, "evidence_")) ||
		(strings.TrimSpace(model.ReplacementEvidenceRef) != "" && !point14Val0RefValid(model.ReplacementEvidenceRef, "evidence_")) ||
		(strings.TrimSpace(model.LineageRef) != "" && !point15ValCDependencyRefValid(model.LineageRef)) ||
		!point15ValCEvidenceLifecycleStatusesReasonValid(model.LifecycleStatus, model.LifecycleReason) {
		return Point15ValCStateBlocked
	}
	if !model.HistoryPreserved || model.CanonicalMutationAttempted {
		return Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.PreviousEvidenceRef) != "" && strings.TrimSpace(model.PreviousEvidenceRef) == strings.TrimSpace(model.ReplacementEvidenceRef) {
		return Point15ValCStateBlocked
	}
	switch strings.TrimSpace(model.LifecycleStatus) {
	case point15ValCLifecycleActive:
		if strings.TrimSpace(model.PreviousEvidenceRef) != "" || strings.TrimSpace(model.ReplacementEvidenceRef) != "" || strings.TrimSpace(model.LineageRef) != "" || model.ReasonDecisive {
			return Point15ValCStateBlocked
		}
		return Point15ValCStateActive
	case point15ValCLifecycleExpired, point15ValCLifecycleRevoked, point15ValCLifecycleTampered, point15ValCLifecycleHashMismatch:
		return Point15ValCStateBlocked
	case point15ValCLifecycleStale:
		return Point15ValCStateReviewRequired
	case point15ValCLifecycleDrifted:
		if strings.TrimSpace(model.LifecycleReason) == point15ValCReasonDecisiveDrift || model.ReasonDecisive {
			return Point15ValCStateBlocked
		}
		return Point15ValCStateReviewRequired
	case point15ValCLifecycleSuperseded:
		if strings.TrimSpace(model.PreviousEvidenceRef) == "" || strings.TrimSpace(model.ReplacementEvidenceRef) == "" {
			return Point15ValCStateBlocked
		}
		if strings.TrimSpace(model.LifecycleReason) == point15ValCReasonSupersededWithLineage {
			if strings.TrimSpace(model.LineageRef) == "" {
				return Point15ValCStateBlocked
			}
			return Point15ValCStateReviewRequired
		}
		if strings.TrimSpace(model.LineageRef) != "" {
			return Point15ValCStateBlocked
		}
		return Point15ValCStateBlocked
	case point15ValCLifecycleRevalidationFailed:
		if strings.TrimSpace(model.LifecycleReason) == point15ValCReasonRetryExhausted {
			return Point15ValCStateBlocked
		}
		return Point15ValCStateReviewRequired
	default:
		return Point15ValCStateBlocked
	}
}

func point15ValCRevocationBoundaryModel(dependency Point15ValCDependencySnapshot) Point15ValCRevocationBoundary {
	return Point15ValCRevocationBoundary{
		RevocationID:     "revocation_boundary_point15_valc_001",
		EvidenceID:       dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.EvidenceID,
		TenantScope:      dependency.InheritedTenantScope,
		HistoryPreserved: true,
	}
}

func EvaluatePoint15ValCRevocationBoundaryState(model Point15ValCRevocationBoundary) string {
	if !point15ValCDependencyRefValid(model.RevocationID) ||
		!point14Val0RefValid(model.EvidenceID, "evidence_") ||
		!point11Val0ScopeValid(model.TenantScope) {
		return Point15ValCStateBlocked
	}
	if model.SourceAuthorityGranted || model.AutoRevoked || model.AutoPublished || !model.HistoryPreserved {
		return Point15ValCStateBlocked
	}
	if !model.RevocationPresent {
		if strings.TrimSpace(model.RevocationSourceRef) != "" || strings.TrimSpace(model.RevocationReceivedAt) != "" || strings.TrimSpace(model.RevocationValidatedAt) != "" || strings.TrimSpace(model.RevocationTimeSource) != "" || model.GovernanceReviewRequired {
			return Point15ValCStateBlocked
		}
		return Point15ValCStateActive
	}
	if !point15ValCDependencyRefValid(model.RevocationSourceRef) {
		return Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.RevocationReceivedAt) == "" || strings.TrimSpace(model.RevocationValidatedAt) == "" {
		return Point15ValCStateIncomplete
	}
	if !point14Val0CanonicalTimeSourceValid(model.RevocationTimeSource) {
		return Point15ValCStateBlocked
	}
	return Point15ValCStateBlocked
}

func point15ValCExpiryBoundaryModel(dependency Point15ValCDependencySnapshot) Point15ValCExpiryBoundary {
	return Point15ValCExpiryBoundary{
		ExpiryID:               "expiry_boundary_point15_valc_001",
		EvidenceID:             dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.EvidenceID,
		TenantScope:            dependency.InheritedTenantScope,
		ExpiringEvidence:       true,
		ExpiresAt:              "2026-05-08T10:00:00Z",
		EvaluatedAt:            "2026-05-07T10:00:00Z",
		ExpiryTimeSource:       point14Val0TimeSourceServerUTC,
		ExpiryEnforced:         false,
		ExpiryHistoryPreserved: true,
	}
}

func EvaluatePoint15ValCExpiryBoundaryState(model Point15ValCExpiryBoundary) string {
	if !point15ValCDependencyRefValid(model.ExpiryID) ||
		!point14Val0RefValid(model.EvidenceID, "evidence_") ||
		!point11Val0ScopeValid(model.TenantScope) {
		return Point15ValCStateBlocked
	}
	if !model.ExpiryHistoryPreserved {
		return Point15ValCStateBlocked
	}
	if !model.ExpiringEvidence {
		if model.ExpiryEnforced {
			return Point15ValCStateBlocked
		}
		return Point15ValCStateActive
	}
	if strings.TrimSpace(model.ExpiresAt) == "" || strings.TrimSpace(model.EvaluatedAt) == "" {
		return Point15ValCStateIncomplete
	}
	if !point14Val0CanonicalTimeSourceValid(model.ExpiryTimeSource) {
		return Point15ValCStateBlocked
	}
	expiresAt, okExpires := point14Val0ParsedTime(model.ExpiresAt)
	evaluatedAt, okEvaluated := point14Val0ParsedTime(model.EvaluatedAt)
	if !okExpires || !okEvaluated {
		return Point15ValCStateBlocked
	}
	if !expiresAt.After(evaluatedAt) {
		return Point15ValCStateBlocked
	}
	if model.ExpiryEnforced {
		return Point15ValCStateBlocked
	}
	return Point15ValCStateActive
}

func point15ValCSupersessionBoundaryModel(dependency Point15ValCDependencySnapshot) Point15ValCSupersessionBoundary {
	return Point15ValCSupersessionBoundary{
		SupersessionID:   "supersession_boundary_point15_valc_001",
		TenantScope:      dependency.InheritedTenantScope,
		HistoryPreserved: true,
	}
}

func EvaluatePoint15ValCSupersessionBoundaryState(model Point15ValCSupersessionBoundary) string {
	if !point15ValCDependencyRefValid(model.SupersessionID) || !point11Val0ScopeValid(model.TenantScope) {
		return Point15ValCStateBlocked
	}
	if model.SilentReplacementDetected || !model.HistoryPreserved || model.AutoPublished || model.AutoApproved {
		return Point15ValCStateBlocked
	}
	if !model.SupersessionPresent {
		if strings.TrimSpace(model.OldEvidenceRef) != "" || strings.TrimSpace(model.NewEvidenceRef) != "" || strings.TrimSpace(model.LineageRef) != "" || strings.TrimSpace(model.ReplacementHash) != "" || strings.TrimSpace(model.PriorHash) != "" {
			return Point15ValCStateBlocked
		}
		return Point15ValCStateActive
	}
	if !point14Val0RefValid(model.OldEvidenceRef, "evidence_") ||
		!point14Val0RefValid(model.NewEvidenceRef, "evidence_") ||
		strings.TrimSpace(model.OldEvidenceRef) == strings.TrimSpace(model.NewEvidenceRef) ||
		strings.TrimSpace(model.ReplacementHash) == "" ||
		!point14Val0HashRefValid(model.ReplacementHash) ||
		strings.TrimSpace(model.PriorHash) == "" ||
		!point14Val0HashRefValid(model.PriorHash) {
		return Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.ReplacementHash) == strings.TrimSpace(model.PriorHash) {
		return Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.LineageRef) == "" || !point15ValCDependencyRefValid(model.LineageRef) {
		return Point15ValCStateBlocked
	}
	return Point15ValCStateReviewRequired
}

func point15ValCReplayProofHistoryBoundaryModel() Point15ValCReplayProofHistoryBoundary {
	return Point15ValCReplayProofHistoryBoundary{
		HistoryID:               "history_boundary_point15_valc_001",
		ProofHistoryRef:         "proof_history_point15_valc_001",
		ReplayRef:               "replay_point15_valc_001",
		ProofPackRef:            "proof_pack_point15_valc_001",
		DecisiveEvidenceVisible: true,
		BlockedReasonVisible:    true,
		PriorStateVisible:       true,
		CurrentStateVisible:     true,
	}
}

func EvaluatePoint15ValCReplayProofHistoryBoundaryState(model Point15ValCReplayProofHistoryBoundary) string {
	if !point15ValCDependencyRefValid(model.HistoryID) ||
		!point15ValCDependencyRefValid(model.ProofHistoryRef) ||
		!point15ValCDependencyRefValid(model.ReplayRef) ||
		!point15ValCDependencyRefValid(model.ProofPackRef) {
		return Point15ValCStateBlocked
	}
	if !model.DecisiveEvidenceVisible || !model.BlockedReasonVisible || !model.PriorStateVisible || !model.CurrentStateVisible || model.ProjectionStrengthensClaims {
		return Point15ValCStateBlocked
	}
	return Point15ValCStateActive
}

func point15ValCTimestampDisciplineModel(dependency Point15ValCDependencySnapshot) Point15ValCTimestampDiscipline {
	return Point15ValCTimestampDiscipline{
		DisciplineID:           "timestamp_discipline_point15_valc_001",
		EvidenceID:             dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.EvidenceID,
		TenantScope:            dependency.InheritedTenantScope,
		ReceivedAt:             "2026-05-07T09:00:00Z",
		ReceivedAtTimeSource:   point14Val0TimeSourceServerUTC,
		ValidatedAt:            "2026-05-07T09:05:00Z",
		ValidatedAtTimeSource:  point14Val0TimeSourceServerUTC,
		EvaluatedAt:            "2026-05-07T09:10:00Z",
		EvaluatedAtTimeSource:  point14Val0TimeSourceServerUTC,
		ReferenceNow:           "2026-05-07T09:15:00Z",
		ReferenceNowTimeSource: point14Val0TimeSourceServerUTC,
	}
}

func EvaluatePoint15ValCTimestampDisciplineState(model Point15ValCTimestampDiscipline) string {
	if !point15ValCDependencyRefValid(model.DisciplineID) ||
		!point14Val0RefValid(model.EvidenceID, "evidence_") ||
		!point11Val0ScopeValid(model.TenantScope) {
		return Point15ValCStateBlocked
	}
	if model.ClientLocalCreatesCanonical || model.SourceEventCreatesCanonical {
		return Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.ReceivedAt) != "" && !point14Val0CanonicalTimeSourceValid(model.ReceivedAtTimeSource) ||
		strings.TrimSpace(model.ValidatedAt) != "" && !point14Val0CanonicalTimeSourceValid(model.ValidatedAtTimeSource) ||
		strings.TrimSpace(model.EnforcedAt) != "" && !point14Val0CanonicalTimeSourceValid(model.EnforcedAtTimeSource) ||
		strings.TrimSpace(model.EvaluatedAt) != "" && !point14Val0CanonicalTimeSourceValid(model.EvaluatedAtTimeSource) ||
		strings.TrimSpace(model.ReferenceNow) != "" && !point14Val0CanonicalTimeSourceValid(model.ReferenceNowTimeSource) {
		return Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.SourceEventAt) != "" {
		if !point14Val0ParsedTimeOk(model.SourceEventAt) || !point14Val0CanonicalTimeSourceValid(model.SourceEventTimeSource) {
			return Point15ValCStateBlocked
		}
	}
	if (model.RevocationPresent || model.ExpiryEnforced || model.SupersessionPresent) && strings.TrimSpace(model.EnforcedAt) == "" {
		return Point15ValCStateIncomplete
	}
	if model.RevocationPresent && (strings.TrimSpace(model.ReceivedAt) == "" || strings.TrimSpace(model.ValidatedAt) == "") {
		return Point15ValCStateIncomplete
	}
	if (model.RevocationPresent || model.ExpiryEnforced || model.SupersessionPresent) &&
		strings.TrimSpace(model.SourceEventAt) != "" &&
		strings.TrimSpace(model.ReceivedAt) == "" &&
		strings.TrimSpace(model.ValidatedAt) == "" &&
		strings.TrimSpace(model.EnforcedAt) == "" {
		return Point15ValCStateBlocked
	}
	referenceNow, okNow := point14Val0ParsedTime(model.ReferenceNow)
	enforcedAt, okEnforced := point14Val0ParsedTime(model.EnforcedAt)
	validatedAt, okValidated := point14Val0ParsedTime(model.ValidatedAt)
	receivedAt, okReceived := point14Val0ParsedTime(model.ReceivedAt)
	evaluatedAt, okEvaluated := point14Val0ParsedTime(model.EvaluatedAt)
	if okEnforced && okNow && enforcedAt.After(referenceNow) {
		return Point15ValCStateBlocked
	}
	if okEnforced && okValidated && enforcedAt.Before(validatedAt) {
		return Point15ValCStateBlocked
	}
	if okValidated && okReceived && validatedAt.Before(receivedAt) {
		return Point15ValCStateReviewRequired
	}
	if okEnforced && okEvaluated && enforcedAt.Before(evaluatedAt) {
		return Point15ValCStateReviewRequired
	}
	return Point15ValCStateActive
}

func point15ValCTenantBoundaryModel(dependency Point15ValCDependencySnapshot) Point15ValCTenantBoundary {
	return Point15ValCTenantBoundary{
		BoundaryID:                   "tenant_boundary_point15_valc_001",
		TenantScope:                  dependency.InheritedTenantScope,
		ReferencedTenantScope:        dependency.InheritedTenantScope,
		EnforcementResultTenantScope: dependency.InheritedTenantScope,
	}
}

func EvaluatePoint15ValCTenantBoundaryState(model Point15ValCTenantBoundary) string {
	if !point15ValCDependencyRefValid(model.BoundaryID) {
		return Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.TenantScope) == "" || strings.TrimSpace(model.ReferencedTenantScope) == "" || strings.TrimSpace(model.EnforcementResultTenantScope) == "" {
		return Point15ValCStateIncomplete
	}
	if !point11Val0ScopeValid(model.TenantScope) ||
		!point11Val0ScopeValid(model.ReferencedTenantScope) ||
		!point11Val0ScopeValid(model.EnforcementResultTenantScope) {
		return Point15ValCStateBlocked
	}
	if model.CrossTenantDetected ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(model.ReferencedTenantScope) ||
		strings.TrimSpace(model.TenantScope) != strings.TrimSpace(model.EnforcementResultTenantScope) {
		return Point15ValCStateBlocked
	}
	return Point15ValCStateActive
}

func point15ValCAuthorityBoundaryModel(dependency Point15ValCDependencySnapshot) Point15ValCAuthorityBoundary {
	return Point15ValCAuthorityBoundary{
		BoundaryID:                      "authority_boundary_point15_valc_001",
		TenantScope:                     dependency.InheritedTenantScope,
		ExternalSourceInputOnly:         true,
		FormalEvaluatorOnly:             true,
		AgentRecommendationAdvisoryOnly: true,
	}
}

func EvaluatePoint15ValCAuthorityBoundaryState(model Point15ValCAuthorityBoundary) string {
	if !point15ValCDependencyRefValid(model.BoundaryID) || !point11Val0ScopeValid(model.TenantScope) {
		return Point15ValCStateBlocked
	}
	if !model.ExternalSourceInputOnly || !model.FormalEvaluatorOnly || !model.AgentRecommendationAdvisoryOnly ||
		model.SchedulerEnforcesBoundary ||
		model.ConnectorRestoresActiveClosure ||
		model.DashboardSuppressesEnforcement ||
		model.PortalProjectionMutatesEnforcement ||
		model.CustomerProjectionMutatesEnforcement ||
		model.AuditorProjectionMutatesEnforcement ||
		model.AgentSatisfiesEnforcement ||
		model.RevocationExecutionSideEffectAllowed ||
		model.AutomaticPublicationAllowed ||
		model.CanonicalMutationAllowed ||
		model.ProductionMutationAllowed ||
		model.PassAllowed {
		return Point15ValCStateBlocked
	}
	return Point15ValCStateActive
}

func point15ValCNoOverclaimGuardModel() Point15ValCNoOverclaimGuard {
	return Point15ValCNoOverclaimGuard{
		AllowedSafeWording:    point15ValCSafeWording(),
		BlockedWording:        point15ValCForbiddenWording(),
		EnforcementDisclaimer: point15ValCEnforcementDisclaimer,
	}
}

func EvaluatePoint15ValCNoOverclaimGuardState(model Point15ValCNoOverclaimGuard) string {
	if strings.TrimSpace(model.EnforcementDisclaimer) != point15ValCEnforcementDisclaimer {
		return Point15ValCStateBlocked
	}
	if !point12Val0ExactStringSetMatch(model.AllowedSafeWording, point15ValCSafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point15ValCForbiddenWording()) {
		return Point15ValCStateBlocked
	}
	if point15ValCObservedListContainsForbiddenWording(model.ObservedTexts) {
		return Point15ValCStateBlocked
	}
	if point15ValCObservedListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point15ValCStateBlocked
	}
	return Point15ValCStateActive
}

func point15ValCFoundationModelFromUpstream(valB Point15ValBScheduledRevalidationFoundation) Point15ValCEnforcementBoundaryFoundation {
	dependency := point15ValCDependencyModel()
	dependency.Point15ValB = valB
	dependency.Point15ValBCurrentState = valB.CurrentState
	dependency.Point15ValBDependencyState = valB.DependencyState
	dependency.Point15ValBScheduleState = valB.ScheduleState
	dependency.Point15ValBRunState = valB.RunState
	dependency.Point15ValBRetryBudgetState = valB.RetryBudgetState
	dependency.Point15ValBTenantThrottleState = valB.TenantThrottleState
	dependency.Point15ValBDowngradeBindingState = valB.DowngradeBindingState
	dependency.Point15ValBTimestampState = valB.TimestampDisciplineState
	dependency.Point15ValBAuthorityState = valB.AuthorityBoundaryState
	dependency.Point15ValBNoOverclaimState = valB.NoOverclaimState
	dependency.Point15ValBComputedFromUpstream = valB.Dependency.SnapshotFromComputedOutput
	dependency.InheritedPoint15ValACurrentState = valB.Dependency.Point15ValA.CurrentState
	dependency.InheritedPoint15Val0CurrentState = valB.Dependency.Point15ValA.Dependency.Point15Val0.CurrentState
	dependency.InheritedPoint14ValECurrentState = valB.Dependency.Point15ValA.Dependency.Point15Val0.Dependency.Point14ValE.CurrentState
	dependency.InheritedTenantScope = valB.Dependency.InheritedTenantScope

	action := point15ValCEnforcementActionModel(dependency)
	lifecycle := point15ValCEvidenceLifecycleBoundaryModel(dependency)
	revocation := point15ValCRevocationBoundaryModel(dependency)
	expiry := point15ValCExpiryBoundaryModel(dependency)
	supersession := point15ValCSupersessionBoundaryModel(dependency)
	timestamp := point15ValCTimestampDisciplineModel(dependency)

	return Point15ValCEnforcementBoundaryFoundation{
		EnforcementDisclaimer: point15ValCEnforcementDisclaimer,
		Dependency:            dependency,
		EnforcementAction:     action,
		EvidenceLifecycle:     lifecycle,
		RevocationBoundary:    revocation,
		ExpiryBoundary:        expiry,
		SupersessionBoundary:  supersession,
		ReplayProofHistory:    point15ValCReplayProofHistoryBoundaryModel(),
		TimestampDiscipline:   timestamp,
		TenantBoundary:        point15ValCTenantBoundaryModel(dependency),
		AuthorityBoundary:     point15ValCAuthorityBoundaryModel(dependency),
		NoOverclaimGuard:      point15ValCNoOverclaimGuardModel(),
	}
}

func Point15ValCFoundationModel() Point15ValCEnforcementBoundaryFoundation {
	valB := ComputePoint15ValBScheduledRevalidationFoundation(Point15ValBFoundationModel())
	return point15ValCFoundationModelFromUpstream(valB)
}

func point15ValCAggregate(states ...string) string {
	for _, state := range states {
		if strings.TrimSpace(state) == Point15ValCStateBlocked {
			return Point15ValCStateBlocked
		}
	}
	for _, state := range states {
		if strings.TrimSpace(state) == Point15ValCStateReviewRequired {
			return Point15ValCStateReviewRequired
		}
	}
	for _, state := range states {
		if strings.TrimSpace(state) == Point15ValCStateIncomplete {
			return Point15ValCStateIncomplete
		}
	}
	return Point15ValCStateActive
}

func point15ValCBlockingReasons(model Point15ValCEnforcementBoundaryFoundation) []string {
	componentStates := map[string]string{
		"dependency":            model.DependencyState,
		"enforcement_action":    model.EnforcementActionState,
		"evidence_lifecycle":    model.EvidenceLifecycleState,
		"revocation_boundary":   model.RevocationBoundaryState,
		"expiry_boundary":       model.ExpiryBoundaryState,
		"supersession_boundary": model.SupersessionState,
		"replay_history":        model.ReplayProofHistoryState,
		"timestamp_discipline":  model.TimestampDisciplineState,
		"authority_boundary":    model.AuthorityBoundaryState,
		"tenant_boundary":       model.TenantBoundaryState,
		"no_overclaim":          model.NoOverclaimState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if strings.TrimSpace(state) == Point15ValCStateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point15ValCReviewPrerequisites(model Point15ValCEnforcementBoundaryFoundation) []string {
	componentStates := map[string]string{
		"enforcement_action":    model.EnforcementActionState,
		"evidence_lifecycle":    model.EvidenceLifecycleState,
		"revocation_boundary":   model.RevocationBoundaryState,
		"expiry_boundary":       model.ExpiryBoundaryState,
		"supersession_boundary": model.SupersessionState,
		"replay_history":        model.ReplayProofHistoryState,
		"timestamp_discipline":  model.TimestampDisciplineState,
		"authority_boundary":    model.AuthorityBoundaryState,
		"tenant_boundary":       model.TenantBoundaryState,
		"no_overclaim":          model.NoOverclaimState,
	}
	prereqs := append([]string{}, model.Dependency.ReviewPrerequisites...)
	for name, state := range componentStates {
		if strings.TrimSpace(state) == Point15ValCStateReviewRequired || strings.TrimSpace(state) == Point15ValCStateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func ComputePoint15ValCEnforcementBoundaryFoundation(model Point15ValCEnforcementBoundaryFoundation) Point15ValCEnforcementBoundaryFoundation {
	model.DependencyState = EvaluatePoint15ValCDependencyState(model.Dependency)
	model.EnforcementActionState = EvaluatePoint15ValCEnforcementActionState(model.EnforcementAction)
	model.EvidenceLifecycleState = EvaluatePoint15ValCEvidenceLifecycleBoundaryState(model.EvidenceLifecycle)
	model.RevocationBoundaryState = EvaluatePoint15ValCRevocationBoundaryState(model.RevocationBoundary)
	model.ExpiryBoundaryState = EvaluatePoint15ValCExpiryBoundaryState(model.ExpiryBoundary)
	model.SupersessionState = EvaluatePoint15ValCSupersessionBoundaryState(model.SupersessionBoundary)
	model.ReplayProofHistoryState = EvaluatePoint15ValCReplayProofHistoryBoundaryState(model.ReplayProofHistory)
	model.TimestampDisciplineState = EvaluatePoint15ValCTimestampDisciplineState(model.TimestampDiscipline)
	model.TenantBoundaryState = EvaluatePoint15ValCTenantBoundaryState(model.TenantBoundary)
	model.AuthorityBoundaryState = EvaluatePoint15ValCAuthorityBoundaryState(model.AuthorityBoundary)
	model.NoOverclaimState = EvaluatePoint15ValCNoOverclaimGuardState(model.NoOverclaimGuard)

	expectedTenant := strings.TrimSpace(model.Dependency.InheritedTenantScope)
	expectedEvidenceID := strings.TrimSpace(model.Dependency.Point15ValB.Dependency.Point15ValA.Dependency.Point15Val0.EvidenceContext.EvidenceID)
	expectedBindingRef := strings.TrimSpace(model.Dependency.Point15ValB.DowngradeBinding.BindingID)
	expectedDecisionRef := strings.TrimSpace(model.Dependency.Point15ValB.Dependency.Point15ValA.Decision.DecisionID)

	if expectedTenant == "" ||
		strings.TrimSpace(model.EnforcementAction.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.EvidenceLifecycle.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.RevocationBoundary.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.ExpiryBoundary.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.SupersessionBoundary.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.TimestampDiscipline.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.TenantBoundary.TenantScope) != expectedTenant ||
		strings.TrimSpace(model.AuthorityBoundary.TenantScope) != expectedTenant {
		model.EnforcementActionState = Point15ValCStateBlocked
		model.EvidenceLifecycleState = Point15ValCStateBlocked
		model.RevocationBoundaryState = Point15ValCStateBlocked
		model.ExpiryBoundaryState = Point15ValCStateBlocked
		model.SupersessionState = Point15ValCStateBlocked
		model.TimestampDisciplineState = Point15ValCStateBlocked
		model.TenantBoundaryState = Point15ValCStateBlocked
		model.AuthorityBoundaryState = Point15ValCStateBlocked
	}
	if expectedEvidenceID == "" ||
		strings.TrimSpace(model.EnforcementAction.EvidenceID) != expectedEvidenceID ||
		strings.TrimSpace(model.EvidenceLifecycle.EvidenceID) != expectedEvidenceID ||
		strings.TrimSpace(model.RevocationBoundary.EvidenceID) != expectedEvidenceID ||
		strings.TrimSpace(model.ExpiryBoundary.EvidenceID) != expectedEvidenceID ||
		strings.TrimSpace(model.TimestampDiscipline.EvidenceID) != expectedEvidenceID {
		model.EnforcementActionState = Point15ValCStateBlocked
		model.EvidenceLifecycleState = Point15ValCStateBlocked
		model.RevocationBoundaryState = Point15ValCStateBlocked
		model.ExpiryBoundaryState = Point15ValCStateBlocked
		model.TimestampDisciplineState = Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.EnforcementAction.SourceValBTriggerRef) != expectedBindingRef ||
		strings.TrimSpace(model.EnforcementAction.SourceValADecisionRef) != expectedDecisionRef {
		model.EnforcementActionState = Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.TenantBoundary.ReferencedTenantScope) != expectedTenant ||
		strings.TrimSpace(model.TenantBoundary.EnforcementResultTenantScope) != expectedTenant {
		model.TenantBoundaryState = Point15ValCStateBlocked
	}
	if strings.TrimSpace(model.EnforcementAction.TargetState) == Point15Val0StateActive {
		if strings.TrimSpace(model.EvidenceLifecycle.LifecycleStatus) != point15ValCLifecycleActive ||
			model.RevocationBoundary.RevocationPresent ||
			model.ExpiryBoundary.ExpiryEnforced ||
			model.SupersessionBoundary.SupersessionPresent {
			model.EnforcementActionState = Point15ValCStateBlocked
		}
	}
	model.CurrentState = point15ValCAggregate(
		model.DependencyState,
		model.EnforcementActionState,
		model.EvidenceLifecycleState,
		model.RevocationBoundaryState,
		model.ExpiryBoundaryState,
		model.SupersessionState,
		model.ReplayProofHistoryState,
		model.TimestampDisciplineState,
		model.AuthorityBoundaryState,
		model.TenantBoundaryState,
		model.NoOverclaimState,
	)
	model.BlockingReasons = point15ValCBlockingReasons(model)
	model.ReviewPrerequisites = point15ValCReviewPrerequisites(model)
	return model
}

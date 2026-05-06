package formal

import (
	"encoding/json"
	"sort"
	"strings"
)

const (
	Point15ValAStateActive         = "point15_vala_downgrade_trigger_table_active"
	Point15ValAStateBlocked        = "point15_vala_downgrade_trigger_table_blocked"
	Point15ValAStateReviewRequired = "point15_vala_downgrade_trigger_table_review_required"
	Point15ValAStateIncomplete     = "point15_vala_downgrade_trigger_table_incomplete"
)

const (
	point15ValAWaveID             = "val_a"
	point15ValATriggerDisclaimer  = "formal_trigger_table_only no_silent_pass_retention point15_vala"
	point15ValABlockedPassToken   = "point_15_pass"
	point15ValATriggerExpired     = "expired_evidence"
	point15ValATriggerRevoked     = "revoked_signal"
	point15ValATriggerStale       = "stale_evidence"
	point15ValATriggerSuperseded  = "superseded_evidence"
	point15ValATriggerPolicyDrift = "policy_drift"
	point15ValATriggerArtifact    = "artifact_drift"
	point15ValATriggerVerifier    = "verifier_drift"
	point15ValATriggerSchema      = "schema_drift"
	point15ValATriggerEngine      = "engine_drift"
	point15ValATriggerConnFail    = "connector_failure"
	point15ValATriggerConnTimeout = "connector_timeout"
	point15ValATriggerConnAuth    = "connector_unauthorized"
	point15ValATriggerConnTenant  = "connector_tenant_mismatch"
	point15ValATriggerHash        = "evidence_hash_mismatch"
	point15ValATriggerMissing     = "missing_freshness_proof"
	point15ValATriggerUnsupported = "unsupported_freshness_status"
	point15ValATriggerTampered    = "tampered_freshness_proof"
)

type Point15ValADependencySnapshot struct {
	Point15Val0CurrentState           string                                   `json:"point15_val0_current_state"`
	Point15Val0DependencyState        string                                   `json:"point15_val0_dependency_state"`
	Point15Val0FreshnessTaxonomyState string                                   `json:"point15_val0_freshness_taxonomy_state"`
	Point15Val0DowngradeTaxonomyState string                                   `json:"point15_val0_downgrade_taxonomy_state"`
	Point15Val0EvidenceContextState   string                                   `json:"point15_val0_evidence_context_state"`
	Point15Val0TenantBoundaryState    string                                   `json:"point15_val0_tenant_boundary_state"`
	Point15Val0TimestampState         string                                   `json:"point15_val0_timestamp_state"`
	Point15Val0AuthorityState         string                                   `json:"point15_val0_authority_state"`
	Point15Val0NoOverclaimState       string                                   `json:"point15_val0_no_overclaim_state"`
	Point15Val0ComputedFromUpstream   bool                                     `json:"point15_val0_computed_from_upstream"`
	Point15Val0Merged                 bool                                     `json:"point15_val0_merged"`
	Point15Val0CIGreen                bool                                     `json:"point15_val0_ci_green"`
	Point15Val0ReviewedOnMain         bool                                     `json:"point15_val0_reviewed_on_main"`
	Point15PassSeen                   bool                                     `json:"point15_pass_seen"`
	InheritedPoint14ValECurrentState  string                                   `json:"inherited_point14_vale_current_state"`
	InheritedTenantScope              string                                   `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput        bool                                     `json:"snapshot_from_computed_output"`
	ReviewPrerequisites               []string                                 `json:"review_prerequisites,omitempty"`
	Point15Val0                       Point15Val0FreshnessDisciplineFoundation `json:"point15_val0"`
}

type Point15ValADowngradeTriggerTable struct {
	TableID                    string   `json:"table_id"`
	AllowedTriggers            []string `json:"allowed_triggers,omitempty"`
	CurrentTriggerDetected     bool     `json:"current_trigger_detected"`
	CurrentTriggerRef          string   `json:"current_trigger_ref"`
	CurrentReasonRef           string   `json:"current_reason_ref"`
	CurrentDecisionRef         string   `json:"current_decision_ref"`
	CurrentTriggerType         string   `json:"current_trigger_type"`
	CurrentTargetState         string   `json:"current_target_state"`
	CurrentDowngradeOutcome    string   `json:"current_downgrade_outcome"`
	FormalEvaluatorOnly        bool     `json:"formal_evaluator_only"`
	SilentPassRetentionAllowed bool     `json:"silent_pass_retention_allowed"`
}

type Point15ValADowngradeTrigger struct {
	TriggerID               string `json:"trigger_id"`
	TriggerDetected         bool   `json:"trigger_detected"`
	TriggerType             string `json:"trigger_type"`
	ObservedFreshnessStatus string `json:"observed_freshness_status"`
	TriggerIsDecisive       bool   `json:"trigger_is_decisive"`
	SupersessionLineageRef  string `json:"supersession_lineage_ref"`
	TargetState             string `json:"target_state"`
	TargetDowngradeOutcome  string `json:"target_downgrade_outcome"`
	RetainsPass             bool   `json:"retains_pass"`
	RetainsActiveClosure    bool   `json:"retains_active_closure"`
}

type Point15ValADowngradeReason struct {
	ReasonID                string `json:"reason_id"`
	TriggerType             string `json:"trigger_type"`
	ReasonCode              string `json:"reason_code"`
	ObservedFreshnessStatus string `json:"observed_freshness_status"`
	Decisive                bool   `json:"decisive"`
	SupersessionLineageRef  string `json:"supersession_lineage_ref"`
	TargetState             string `json:"target_state"`
	TargetDowngradeOutcome  string `json:"target_downgrade_outcome"`
}

type Point15ValADowngradeDecision struct {
	DecisionID             string `json:"decision_id"`
	TriggerRef             string `json:"trigger_ref"`
	ReasonRef              string `json:"reason_ref"`
	TriggerDetected        bool   `json:"trigger_detected"`
	TriggerType            string `json:"trigger_type"`
	TargetState            string `json:"target_state"`
	TargetDowngradeOutcome string `json:"target_downgrade_outcome"`
	RetainsPass            bool   `json:"retains_pass"`
	RetainsActiveClosure   bool   `json:"retains_active_closure"`
	FormalEvaluatorOnly    bool   `json:"formal_evaluator_only"`
}

type Point15ValAAuthorityBoundary struct {
	BoundaryID                         string `json:"boundary_id"`
	TenantScope                        string `json:"tenant_scope"`
	ExternalSourceInputOnly            bool   `json:"external_source_input_only"`
	FormalEvaluatorOnly                bool   `json:"formal_evaluator_only"`
	SchedulerMapsTriggerToDowngrade    bool   `json:"scheduler_maps_trigger_to_downgrade"`
	DashboardMapsTriggerToDowngrade    bool   `json:"dashboard_maps_trigger_to_downgrade"`
	ConnectorMapsTriggerToDowngrade    bool   `json:"connector_maps_trigger_to_downgrade"`
	AgentMapsTriggerToDowngrade        bool   `json:"agent_maps_trigger_to_downgrade"`
	CustomerProjectionMutatesDowngrade bool   `json:"customer_projection_mutates_downgrade"`
	AuditorProjectionMutatesDowngrade  bool   `json:"auditor_projection_mutates_downgrade"`
	PortalProjectionMutatesDowngrade   bool   `json:"portal_projection_mutates_downgrade"`
	ConnectorSuppressesFailure         bool   `json:"connector_suppresses_failure"`
	ConnectorRestoresActiveClosure     bool   `json:"connector_restores_active_closure"`
	ConnectorMarksEvidenceFresh        bool   `json:"connector_marks_evidence_fresh"`
	ConnectorOverridesTerminalStatus   bool   `json:"connector_overrides_terminal_status"`
	CanonicalMutationAllowed           bool   `json:"canonical_mutation_allowed"`
	ProductionMutationAllowed          bool   `json:"production_mutation_allowed"`
	PassAllowed                        bool   `json:"pass_allowed"`
}

type Point15ValANoOverclaimGuard struct {
	ObservedTexts                        []string `json:"observed_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	TriggerDisclaimer                    string   `json:"trigger_disclaimer"`
}

type Point15ValADowngradeTriggerFoundation struct {
	CurrentState           string                           `json:"current_state"`
	BlockingReasons        []string                         `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites    []string                         `json:"review_prerequisites,omitempty"`
	TriggerDisclaimer      string                           `json:"trigger_disclaimer"`
	DependencyState        string                           `json:"dependency_state"`
	TriggerTableState      string                           `json:"trigger_table_state"`
	TriggerState           string                           `json:"trigger_state"`
	ReasonState            string                           `json:"reason_state"`
	DecisionState          string                           `json:"decision_state"`
	AuthorityBoundaryState string                           `json:"authority_boundary_state"`
	NoOverclaimState       string                           `json:"no_overclaim_state"`
	Dependency             Point15ValADependencySnapshot    `json:"dependency"`
	TriggerTable           Point15ValADowngradeTriggerTable `json:"trigger_table"`
	Trigger                Point15ValADowngradeTrigger      `json:"trigger"`
	Reason                 Point15ValADowngradeReason       `json:"reason"`
	Decision               Point15ValADowngradeDecision     `json:"decision"`
	AuthorityBoundary      Point15ValAAuthorityBoundary     `json:"authority_boundary"`
	NoOverclaimGuard       Point15ValANoOverclaimGuard      `json:"no_overclaim_guard"`
}

func point15ValAStates() []string {
	return []string{
		Point15ValAStateActive,
		Point15ValAStateBlocked,
		Point15ValAStateReviewRequired,
		Point15ValAStateIncomplete,
	}
}

func point15ValAStateValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValAStates())
}

func point15ValATriggers() []string {
	return []string{
		point15ValATriggerArtifact,
		point15ValATriggerConnAuth,
		point15ValATriggerConnFail,
		point15ValATriggerConnTenant,
		point15ValATriggerConnTimeout,
		point15ValATriggerEngine,
		point15ValATriggerExpired,
		point15ValATriggerHash,
		point15ValATriggerMissing,
		point15ValATriggerPolicyDrift,
		point15ValATriggerRevoked,
		point15ValATriggerSchema,
		point15ValATriggerStale,
		point15ValATriggerSuperseded,
		point15ValATriggerTampered,
		point15ValATriggerUnsupported,
		point15ValATriggerVerifier,
	}
}

func point15ValATriggerValid(value string) bool {
	return point14Val0ExactValueValid(value, point15ValATriggers())
}

func point15ValADependencyRefValid(value string) bool {
	return point14Val0RefValid(value, "point15_vala_", "trigger_", "reason_", "decision_", "authority_", "table_")
}

func point15ValAForbiddenWording() []string {
	return []string{
		"continuous assurance guaranteed",
		"always fresh",
		"automatically verified forever",
		"production approved",
		"compliance guaranteed",
		"regulator approved",
		"regulator-approved",
		"certified secure",
		"autonomous assurance pass",
		"public badge",
		"global truth",
	}
}

func point15ValASafeWording() []string {
	return []string{
		"downgrade trigger detected",
		"freshness requires review",
		"evidence support available for review",
		"bounded freshness decision",
		"trigger mapped by formal evaluator only",
	}
}

func point15ValAObservedTextContainsForbiddenWording(text string) bool {
	trimmed := strings.TrimSpace(strings.ToLower(text))
	if trimmed == "" {
		return false
	}
	for _, safe := range point15ValASafeWording() {
		if trimmed == strings.ToLower(strings.TrimSpace(safe)) {
			return false
		}
	}
	for _, forbidden := range point15ValAForbiddenWording() {
		if strings.Contains(trimmed, strings.ToLower(strings.TrimSpace(forbidden))) {
			return true
		}
	}
	return false
}

func point15ValAObservedListContainsForbiddenWording(values []string) bool {
	for _, value := range values {
		if point15ValAObservedTextContainsForbiddenWording(value) {
			return true
		}
	}
	return false
}

func point15ValAVal0PayloadContainsPoint15Pass(val0 Point15Val0FreshnessDisciplineFoundation) bool {
	payload, err := json.Marshal(val0)
	if err != nil {
		return true
	}
	return strings.Contains(string(payload), point15ValABlockedPassToken)
}

func point15ValATriggerObservedFreshnessStatus(triggerType string) string {
	switch strings.TrimSpace(triggerType) {
	case point15ValATriggerExpired:
		return point15Val0FreshnessExpired
	case point15ValATriggerRevoked:
		return point15Val0FreshnessRevoked
	case point15ValATriggerStale:
		return point15Val0FreshnessStale
	case point15ValATriggerSuperseded:
		return point15Val0FreshnessSuperseded
	case point15ValATriggerPolicyDrift, point15ValATriggerArtifact, point15ValATriggerVerifier, point15ValATriggerSchema, point15ValATriggerEngine:
		return point15Val0FreshnessDrifted
	case point15ValATriggerConnFail, point15ValATriggerConnTimeout, point15ValATriggerConnAuth, point15ValATriggerConnTenant, point15ValATriggerMissing:
		return point15Val0FreshnessMissing
	case point15ValATriggerHash, point15ValATriggerTampered:
		return point15Val0FreshnessTampered
	case point15ValATriggerUnsupported:
		return point15Val0FreshnessUnsupported
	default:
		return point15Val0FreshnessFresh
	}
}

func point15ValATriggerExpectedOutcome(triggerType string, decisive bool, lineageValid bool) string {
	switch strings.TrimSpace(triggerType) {
	case "":
		return point15Val0DowngradeRetainActive
	case point15ValATriggerExpired, point15ValATriggerRevoked, point15ValATriggerUnsupported, point15ValATriggerTampered, point15ValATriggerHash, point15ValATriggerConnAuth, point15ValATriggerConnTenant:
		return point15Val0DowngradeBlocked
	case point15ValATriggerStale, point15ValATriggerConnFail, point15ValATriggerConnTimeout:
		return point15Val0DowngradeReview
	case point15ValATriggerSuperseded:
		if lineageValid {
			return point15Val0DowngradeReview
		}
		return point15Val0DowngradeBlocked
	case point15ValATriggerPolicyDrift, point15ValATriggerArtifact, point15ValATriggerVerifier, point15ValATriggerSchema, point15ValATriggerEngine:
		if decisive {
			return point15Val0DowngradeBlocked
		}
		return point15Val0DowngradeReview
	case point15ValATriggerMissing:
		if decisive {
			return point15Val0DowngradeBlocked
		}
		return point15Val0DowngradeIncomplete
	default:
		return ""
	}
}

func point15ValATriggerExpectedState(triggerType string, decisive bool, lineageValid bool) string {
	switch point15ValATriggerExpectedOutcome(triggerType, decisive, lineageValid) {
	case point15Val0DowngradeRetainActive:
		return Point15Val0StateActive
	case point15Val0DowngradeReview:
		return Point15Val0StateReviewRequired
	case point15Val0DowngradeIncomplete:
		return Point15Val0StateIncomplete
	case point15Val0DowngradeBlocked:
		return Point15Val0StateBlocked
	default:
		return Point15Val0StateBlocked
	}
}

func point15ValAExpectedReasonCode(triggerType string, lineageValid bool) string {
	switch strings.TrimSpace(triggerType) {
	case "":
		return ""
	case point15ValATriggerExpired:
		return "evidence_validity_window_expired"
	case point15ValATriggerRevoked:
		return "signal_revoked_by_bounded_input"
	case point15ValATriggerStale:
		return "evidence_stale_requires_review"
	case point15ValATriggerSuperseded:
		if lineageValid {
			return "evidence_superseded_requires_lineage_review"
		}
		return "evidence_superseded_lineage_missing"
	case point15ValATriggerPolicyDrift:
		return "policy_version_drift_detected"
	case point15ValATriggerArtifact:
		return "artifact_identity_drift_detected"
	case point15ValATriggerVerifier:
		return "verifier_version_drift_detected"
	case point15ValATriggerSchema:
		return "schema_version_drift_detected"
	case point15ValATriggerEngine:
		return "engine_version_drift_detected"
	case point15ValATriggerConnFail:
		return "connector_freshness_signal_unavailable"
	case point15ValATriggerConnTimeout:
		return "connector_freshness_request_timed_out"
	case point15ValATriggerConnAuth:
		return "connector_freshness_access_unauthorized"
	case point15ValATriggerConnTenant:
		return "connector_freshness_tenant_mismatch"
	case point15ValATriggerHash:
		return "evidence_hash_mismatch_detected"
	case point15ValATriggerMissing:
		return "freshness_proof_missing"
	case point15ValATriggerUnsupported:
		return "freshness_status_unsupported"
	case point15ValATriggerTampered:
		return "freshness_proof_tampered"
	default:
		return ""
	}
}

func point15ValATargetStateToWaveState(state string) string {
	switch strings.TrimSpace(state) {
	case Point15Val0StateActive:
		return Point15ValAStateActive
	case Point15Val0StateReviewRequired:
		return Point15ValAStateReviewRequired
	case Point15Val0StateIncomplete:
		return Point15ValAStateIncomplete
	case Point15Val0StateBlocked:
		return Point15ValAStateBlocked
	default:
		return Point15ValAStateBlocked
	}
}

func point15ValADependencySnapshotFromUpstream(val0 Point15Val0FreshnessDisciplineFoundation) Point15ValADependencySnapshot {
	return Point15ValADependencySnapshot{
		Point15Val0CurrentState:           val0.CurrentState,
		Point15Val0DependencyState:        val0.DependencyState,
		Point15Val0FreshnessTaxonomyState: val0.FreshnessTaxonomyState,
		Point15Val0DowngradeTaxonomyState: val0.DowngradeTaxonomyState,
		Point15Val0EvidenceContextState:   val0.EvidenceContextState,
		Point15Val0TenantBoundaryState:    val0.TenantBoundaryState,
		Point15Val0TimestampState:         val0.TimestampDisciplineState,
		Point15Val0AuthorityState:         val0.AuthorityBoundaryState,
		Point15Val0NoOverclaimState:       val0.NoOverclaimState,
		Point15Val0ComputedFromUpstream:   val0.Dependency.SnapshotFromComputedOutput,
		Point15Val0Merged:                 true,
		Point15Val0CIGreen:                true,
		Point15Val0ReviewedOnMain:         true,
		Point15PassSeen:                   point15ValAVal0PayloadContainsPoint15Pass(val0),
		InheritedPoint14ValECurrentState:  val0.Dependency.Point14ValECurrentState,
		InheritedTenantScope:              val0.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:        true,
		ReviewPrerequisites:               append([]string{}, val0.ReviewPrerequisites...),
		Point15Val0:                       val0,
	}
}

func point15ValADependencySnapshotModel() Point15ValADependencySnapshot {
	val0 := ComputePoint15Val0FreshnessDisciplineFoundation(Point15Val0FoundationModel())
	return point15ValADependencySnapshotFromUpstream(val0)
}

func EvaluatePoint15ValADependencyState(model Point15ValADependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point15Val0ComputedFromUpstream ||
		!model.Point15Val0Merged ||
		!model.Point15Val0CIGreen ||
		!model.Point15Val0ReviewedOnMain ||
		model.Point15PassSeen ||
		!point15Val0StateValid(model.Point15Val0CurrentState) ||
		!point15Val0StateValid(model.Point15Val0DependencyState) ||
		!point15Val0StateValid(model.Point15Val0FreshnessTaxonomyState) ||
		!point15Val0StateValid(model.Point15Val0DowngradeTaxonomyState) ||
		!point15Val0StateValid(model.Point15Val0EvidenceContextState) ||
		!point15Val0StateValid(model.Point15Val0TenantBoundaryState) ||
		!point15Val0StateValid(model.Point15Val0TimestampState) ||
		!point15Val0StateValid(model.Point15Val0AuthorityState) ||
		!point15Val0StateValid(model.Point15Val0NoOverclaimState) ||
		!point14ValEStateValid(model.InheritedPoint14ValECurrentState) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return Point15ValAStateBlocked
	}
	if strings.TrimSpace(model.Point15Val0CurrentState) != strings.TrimSpace(model.Point15Val0.CurrentState) ||
		strings.TrimSpace(model.Point15Val0DependencyState) != strings.TrimSpace(model.Point15Val0.DependencyState) ||
		strings.TrimSpace(model.Point15Val0FreshnessTaxonomyState) != strings.TrimSpace(model.Point15Val0.FreshnessTaxonomyState) ||
		strings.TrimSpace(model.Point15Val0DowngradeTaxonomyState) != strings.TrimSpace(model.Point15Val0.DowngradeTaxonomyState) ||
		strings.TrimSpace(model.Point15Val0EvidenceContextState) != strings.TrimSpace(model.Point15Val0.EvidenceContextState) ||
		strings.TrimSpace(model.Point15Val0TenantBoundaryState) != strings.TrimSpace(model.Point15Val0.TenantBoundaryState) ||
		strings.TrimSpace(model.Point15Val0TimestampState) != strings.TrimSpace(model.Point15Val0.TimestampDisciplineState) ||
		strings.TrimSpace(model.Point15Val0AuthorityState) != strings.TrimSpace(model.Point15Val0.AuthorityBoundaryState) ||
		strings.TrimSpace(model.Point15Val0NoOverclaimState) != strings.TrimSpace(model.Point15Val0.NoOverclaimState) ||
		model.Point15Val0ComputedFromUpstream != model.Point15Val0.Dependency.SnapshotFromComputedOutput ||
		strings.TrimSpace(model.InheritedPoint14ValECurrentState) != strings.TrimSpace(model.Point15Val0.Dependency.Point14ValECurrentState) ||
		strings.TrimSpace(model.InheritedTenantScope) != strings.TrimSpace(model.Point15Val0.Dependency.InheritedTenantScope) {
		return Point15ValAStateBlocked
	}
	if strings.TrimSpace(model.Point15Val0CurrentState) != Point15Val0StateActive ||
		strings.TrimSpace(model.Point15Val0DependencyState) != Point15Val0StateActive ||
		strings.TrimSpace(model.Point15Val0FreshnessTaxonomyState) != Point15Val0StateActive ||
		strings.TrimSpace(model.Point15Val0DowngradeTaxonomyState) != Point15Val0StateActive ||
		strings.TrimSpace(model.Point15Val0EvidenceContextState) != Point15Val0StateActive ||
		strings.TrimSpace(model.Point15Val0TenantBoundaryState) != Point15Val0StateActive ||
		strings.TrimSpace(model.Point15Val0TimestampState) != Point15Val0StateActive ||
		strings.TrimSpace(model.Point15Val0AuthorityState) != Point15Val0StateActive ||
		strings.TrimSpace(model.Point15Val0NoOverclaimState) != Point15Val0StateActive ||
		strings.TrimSpace(model.InheritedPoint14ValECurrentState) != Point14ValEStatePassConfirmed {
		return Point15ValAStateBlocked
	}
	return Point15ValAStateActive
}

func point15ValADowngradeTriggerTableModel() Point15ValADowngradeTriggerTable {
	return Point15ValADowngradeTriggerTable{
		TableID:                 "point15_vala_trigger_table_001",
		AllowedTriggers:         point15ValATriggers(),
		CurrentTargetState:      Point15Val0StateActive,
		CurrentDowngradeOutcome: point15Val0DowngradeRetainActive,
		FormalEvaluatorOnly:     true,
	}
}

func EvaluatePoint15ValADowngradeTriggerTableState(model Point15ValADowngradeTriggerTable) string {
	if !point15ValADependencyRefValid(model.TableID) ||
		!point12Val0ExactStringSetMatch(model.AllowedTriggers, point15ValATriggers()) ||
		!point15Val0StateValid(model.CurrentTargetState) ||
		!point15Val0DowngradeOutcomeValid(model.CurrentDowngradeOutcome) ||
		!model.FormalEvaluatorOnly ||
		model.SilentPassRetentionAllowed {
		return Point15ValAStateBlocked
	}
	if model.CurrentTriggerDetected {
		if !point15ValATriggerValid(model.CurrentTriggerType) ||
			!point15ValADependencyRefValid(model.CurrentTriggerRef) ||
			!point15ValADependencyRefValid(model.CurrentReasonRef) ||
			!point15ValADependencyRefValid(model.CurrentDecisionRef) ||
			strings.TrimSpace(model.CurrentTargetState) == Point15Val0StateActive ||
			strings.TrimSpace(model.CurrentDowngradeOutcome) == point15Val0DowngradeRetainActive {
			return Point15ValAStateBlocked
		}
		return point15ValATargetStateToWaveState(model.CurrentTargetState)
	}
	if strings.TrimSpace(model.CurrentTriggerRef) != "" ||
		strings.TrimSpace(model.CurrentReasonRef) != "" ||
		strings.TrimSpace(model.CurrentDecisionRef) != "" ||
		strings.TrimSpace(model.CurrentTriggerType) != "" ||
		strings.TrimSpace(model.CurrentTargetState) != Point15Val0StateActive ||
		strings.TrimSpace(model.CurrentDowngradeOutcome) != point15Val0DowngradeRetainActive {
		return Point15ValAStateBlocked
	}
	return Point15ValAStateActive
}

func point15ValADowngradeTriggerModel() Point15ValADowngradeTrigger {
	return Point15ValADowngradeTrigger{
		TriggerID:               "trigger_point15_vala_001",
		ObservedFreshnessStatus: point15Val0FreshnessFresh,
		TargetState:             Point15Val0StateActive,
		TargetDowngradeOutcome:  point15Val0DowngradeRetainActive,
		RetainsActiveClosure:    true,
	}
}

func EvaluatePoint15ValADowngradeTriggerState(model Point15ValADowngradeTrigger) string {
	if !point15ValADependencyRefValid(model.TriggerID) ||
		!point15Val0StateValid(model.TargetState) ||
		!point15Val0DowngradeOutcomeValid(model.TargetDowngradeOutcome) ||
		(strings.TrimSpace(model.ObservedFreshnessStatus) != "" && !point15Val0FreshnessStatusValid(model.ObservedFreshnessStatus)) {
		return Point15ValAStateBlocked
	}
	if model.TriggerDetected {
		if !point15ValATriggerValid(model.TriggerType) {
			return Point15ValAStateBlocked
		}
		lineageValid := point15Val0LineageRefValid(model.SupersessionLineageRef)
		expectedOutcome := point15ValATriggerExpectedOutcome(model.TriggerType, model.TriggerIsDecisive, lineageValid)
		expectedState := point15ValATriggerExpectedState(model.TriggerType, model.TriggerIsDecisive, lineageValid)
		expectedFreshness := point15ValATriggerObservedFreshnessStatus(model.TriggerType)
		if expectedOutcome == "" ||
			strings.TrimSpace(model.ObservedFreshnessStatus) != expectedFreshness ||
			strings.TrimSpace(model.TargetState) != expectedState ||
			strings.TrimSpace(model.TargetDowngradeOutcome) != expectedOutcome ||
			model.RetainsPass ||
			model.RetainsActiveClosure {
			return Point15ValAStateBlocked
		}
		return point15ValATargetStateToWaveState(model.TargetState)
	}
	if strings.TrimSpace(model.TriggerType) != "" ||
		model.TriggerIsDecisive ||
		strings.TrimSpace(model.SupersessionLineageRef) != "" ||
		strings.TrimSpace(model.ObservedFreshnessStatus) != point15Val0FreshnessFresh ||
		strings.TrimSpace(model.TargetState) != Point15Val0StateActive ||
		strings.TrimSpace(model.TargetDowngradeOutcome) != point15Val0DowngradeRetainActive ||
		model.RetainsPass ||
		!model.RetainsActiveClosure {
		return Point15ValAStateBlocked
	}
	return Point15ValAStateActive
}

func point15ValADowngradeReasonModel() Point15ValADowngradeReason {
	return Point15ValADowngradeReason{
		ReasonID:                "reason_point15_vala_001",
		ObservedFreshnessStatus: point15Val0FreshnessFresh,
		TargetState:             Point15Val0StateActive,
		TargetDowngradeOutcome:  point15Val0DowngradeRetainActive,
	}
}

func EvaluatePoint15ValADowngradeReasonState(model Point15ValADowngradeReason) string {
	if !point15ValADependencyRefValid(model.ReasonID) ||
		!point15Val0StateValid(model.TargetState) ||
		!point15Val0DowngradeOutcomeValid(model.TargetDowngradeOutcome) ||
		(strings.TrimSpace(model.ObservedFreshnessStatus) != "" && !point15Val0FreshnessStatusValid(model.ObservedFreshnessStatus)) {
		return Point15ValAStateBlocked
	}
	if strings.TrimSpace(model.TriggerType) == "" {
		if strings.TrimSpace(model.ReasonCode) != "" ||
			model.Decisive ||
			strings.TrimSpace(model.SupersessionLineageRef) != "" ||
			strings.TrimSpace(model.ObservedFreshnessStatus) != point15Val0FreshnessFresh ||
			strings.TrimSpace(model.TargetState) != Point15Val0StateActive ||
			strings.TrimSpace(model.TargetDowngradeOutcome) != point15Val0DowngradeRetainActive {
			return Point15ValAStateBlocked
		}
		return Point15ValAStateActive
	}
	if !point15ValATriggerValid(model.TriggerType) {
		return Point15ValAStateBlocked
	}
	lineageValid := point15Val0LineageRefValid(model.SupersessionLineageRef)
	expectedReason := point15ValAExpectedReasonCode(model.TriggerType, lineageValid)
	expectedOutcome := point15ValATriggerExpectedOutcome(model.TriggerType, model.Decisive, lineageValid)
	expectedState := point15ValATriggerExpectedState(model.TriggerType, model.Decisive, lineageValid)
	expectedFreshness := point15ValATriggerObservedFreshnessStatus(model.TriggerType)
	if strings.TrimSpace(model.ReasonCode) != expectedReason ||
		strings.TrimSpace(model.ObservedFreshnessStatus) != expectedFreshness ||
		strings.TrimSpace(model.TargetState) != expectedState ||
		strings.TrimSpace(model.TargetDowngradeOutcome) != expectedOutcome {
		return Point15ValAStateBlocked
	}
	return point15ValATargetStateToWaveState(model.TargetState)
}

func point15ValADowngradeDecisionModel() Point15ValADowngradeDecision {
	return Point15ValADowngradeDecision{
		DecisionID:             "decision_point15_vala_001",
		TargetState:            Point15Val0StateActive,
		TargetDowngradeOutcome: point15Val0DowngradeRetainActive,
		RetainsActiveClosure:   true,
		FormalEvaluatorOnly:    true,
	}
}

func EvaluatePoint15ValADowngradeDecisionState(model Point15ValADowngradeDecision) string {
	if !point15ValADependencyRefValid(model.DecisionID) ||
		!point15Val0StateValid(model.TargetState) ||
		!point15Val0DowngradeOutcomeValid(model.TargetDowngradeOutcome) ||
		!model.FormalEvaluatorOnly {
		return Point15ValAStateBlocked
	}
	if model.TriggerDetected {
		if !point15ValATriggerValid(model.TriggerType) ||
			!point15ValADependencyRefValid(model.TriggerRef) ||
			!point15ValADependencyRefValid(model.ReasonRef) ||
			model.RetainsPass ||
			model.RetainsActiveClosure {
			return Point15ValAStateBlocked
		}
		if strings.TrimSpace(model.TargetState) == Point15Val0StateActive ||
			strings.TrimSpace(model.TargetDowngradeOutcome) == point15Val0DowngradeRetainActive {
			return Point15ValAStateBlocked
		}
		return point15ValATargetStateToWaveState(model.TargetState)
	}
	if strings.TrimSpace(model.TriggerRef) != "" ||
		strings.TrimSpace(model.ReasonRef) != "" ||
		strings.TrimSpace(model.TriggerType) != "" ||
		strings.TrimSpace(model.TargetState) != Point15Val0StateActive ||
		strings.TrimSpace(model.TargetDowngradeOutcome) != point15Val0DowngradeRetainActive ||
		model.RetainsPass ||
		!model.RetainsActiveClosure {
		return Point15ValAStateBlocked
	}
	return Point15ValAStateActive
}

func point15ValAAuthorityBoundaryModel(dependency Point15ValADependencySnapshot) Point15ValAAuthorityBoundary {
	return Point15ValAAuthorityBoundary{
		BoundaryID:              "authority_boundary_point15_vala_001",
		TenantScope:             dependency.InheritedTenantScope,
		ExternalSourceInputOnly: true,
		FormalEvaluatorOnly:     true,
	}
}

func EvaluatePoint15ValAAuthorityBoundaryState(model Point15ValAAuthorityBoundary) string {
	if !point15ValADependencyRefValid(model.BoundaryID) ||
		!point11Val0ScopeValid(model.TenantScope) ||
		!model.ExternalSourceInputOnly ||
		!model.FormalEvaluatorOnly {
		return Point15ValAStateBlocked
	}
	if model.SchedulerMapsTriggerToDowngrade ||
		model.DashboardMapsTriggerToDowngrade ||
		model.ConnectorMapsTriggerToDowngrade ||
		model.AgentMapsTriggerToDowngrade ||
		model.CustomerProjectionMutatesDowngrade ||
		model.AuditorProjectionMutatesDowngrade ||
		model.PortalProjectionMutatesDowngrade ||
		model.ConnectorSuppressesFailure ||
		model.ConnectorRestoresActiveClosure ||
		model.ConnectorMarksEvidenceFresh ||
		model.ConnectorOverridesTerminalStatus ||
		model.CanonicalMutationAllowed ||
		model.ProductionMutationAllowed ||
		model.PassAllowed {
		return Point15ValAStateBlocked
	}
	return Point15ValAStateActive
}

func point15ValANoOverclaimGuardModel() Point15ValANoOverclaimGuard {
	return Point15ValANoOverclaimGuard{
		ObservedTexts: []string{
			"downgrade trigger detected",
			"trigger mapped by formal evaluator only",
		},
		AllowedSafeWording: point15ValASafeWording(),
		BlockedWording:     point15ValAForbiddenWording(),
		TriggerDisclaimer:  point15ValATriggerDisclaimer,
	}
}

func EvaluatePoint15ValANoOverclaimGuardState(model Point15ValANoOverclaimGuard) string {
	if strings.TrimSpace(model.TriggerDisclaimer) != point15ValATriggerDisclaimer ||
		!point12Val0ExactStringSetMatch(model.AllowedSafeWording, point15ValASafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point15ValAForbiddenWording()) {
		return Point15ValAStateBlocked
	}
	if point15ValAObservedListContainsForbiddenWording(model.ObservedTexts) {
		return Point15ValAStateBlocked
	}
	if point15ValAObservedListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point15ValAStateBlocked
	}
	return Point15ValAStateActive
}

func point15ValAFoundationModelFromUpstream(val0 Point15Val0FreshnessDisciplineFoundation) Point15ValADowngradeTriggerFoundation {
	dependency := point15ValADependencySnapshotFromUpstream(val0)
	return Point15ValADowngradeTriggerFoundation{
		TriggerDisclaimer: point15ValATriggerDisclaimer,
		Dependency:        dependency,
		TriggerTable:      point15ValADowngradeTriggerTableModel(),
		Trigger:           point15ValADowngradeTriggerModel(),
		Reason:            point15ValADowngradeReasonModel(),
		Decision:          point15ValADowngradeDecisionModel(),
		AuthorityBoundary: point15ValAAuthorityBoundaryModel(dependency),
		NoOverclaimGuard:  point15ValANoOverclaimGuardModel(),
	}
}

func Point15ValAFoundationModel() Point15ValADowngradeTriggerFoundation {
	val0 := ComputePoint15Val0FreshnessDisciplineFoundation(Point15Val0FoundationModel())
	return point15ValAFoundationModelFromUpstream(val0)
}

func point15ValAAggregate(states ...string) string {
	for _, state := range states {
		if strings.TrimSpace(state) == Point15ValAStateBlocked {
			return Point15ValAStateBlocked
		}
	}
	for _, state := range states {
		if strings.TrimSpace(state) == Point15ValAStateReviewRequired {
			return Point15ValAStateReviewRequired
		}
	}
	for _, state := range states {
		if strings.TrimSpace(state) == Point15ValAStateIncomplete {
			return Point15ValAStateIncomplete
		}
	}
	return Point15ValAStateActive
}

func point15ValABlockingReasons(model Point15ValADowngradeTriggerFoundation) []string {
	componentStates := map[string]string{
		"dependency":         model.DependencyState,
		"trigger_table":      model.TriggerTableState,
		"trigger":            model.TriggerState,
		"reason":             model.ReasonState,
		"decision":           model.DecisionState,
		"authority_boundary": model.AuthorityBoundaryState,
		"no_overclaim":       model.NoOverclaimState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if strings.TrimSpace(state) == Point15ValAStateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point15ValAReviewPrerequisites(model Point15ValADowngradeTriggerFoundation) []string {
	componentStates := map[string]string{
		"trigger_table":      model.TriggerTableState,
		"trigger":            model.TriggerState,
		"reason":             model.ReasonState,
		"decision":           model.DecisionState,
		"authority_boundary": model.AuthorityBoundaryState,
		"no_overclaim":       model.NoOverclaimState,
	}
	prereqs := append([]string{}, model.Dependency.ReviewPrerequisites...)
	for name, state := range componentStates {
		if strings.TrimSpace(state) == Point15ValAStateReviewRequired || strings.TrimSpace(state) == Point15ValAStateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func ComputePoint15ValADowngradeTriggerFoundation(model Point15ValADowngradeTriggerFoundation) Point15ValADowngradeTriggerFoundation {
	model.DependencyState = EvaluatePoint15ValADependencyState(model.Dependency)
	model.TriggerTableState = EvaluatePoint15ValADowngradeTriggerTableState(model.TriggerTable)
	model.TriggerState = EvaluatePoint15ValADowngradeTriggerState(model.Trigger)
	model.ReasonState = EvaluatePoint15ValADowngradeReasonState(model.Reason)
	model.DecisionState = EvaluatePoint15ValADowngradeDecisionState(model.Decision)
	model.AuthorityBoundaryState = EvaluatePoint15ValAAuthorityBoundaryState(model.AuthorityBoundary)
	model.NoOverclaimState = EvaluatePoint15ValANoOverclaimGuardState(model.NoOverclaimGuard)

	if strings.TrimSpace(model.AuthorityBoundary.TenantScope) != strings.TrimSpace(model.Dependency.InheritedTenantScope) {
		model.AuthorityBoundaryState = Point15ValAStateBlocked
	}
	if model.TriggerTable.CurrentTriggerDetected != model.Trigger.TriggerDetected ||
		model.Trigger.TriggerDetected != model.Decision.TriggerDetected {
		model.TriggerTableState = Point15ValAStateBlocked
		model.DecisionState = Point15ValAStateBlocked
	}
	if model.TriggerTable.CurrentTriggerDetected {
		if strings.TrimSpace(model.TriggerTable.CurrentTriggerRef) != strings.TrimSpace(model.Trigger.TriggerID) ||
			strings.TrimSpace(model.TriggerTable.CurrentReasonRef) != strings.TrimSpace(model.Reason.ReasonID) ||
			strings.TrimSpace(model.TriggerTable.CurrentDecisionRef) != strings.TrimSpace(model.Decision.DecisionID) ||
			strings.TrimSpace(model.TriggerTable.CurrentTriggerType) != strings.TrimSpace(model.Trigger.TriggerType) ||
			strings.TrimSpace(model.Trigger.TriggerType) != strings.TrimSpace(model.Reason.TriggerType) ||
			strings.TrimSpace(model.Reason.TriggerType) != strings.TrimSpace(model.Decision.TriggerType) {
			model.TriggerTableState = Point15ValAStateBlocked
			model.ReasonState = Point15ValAStateBlocked
			model.DecisionState = Point15ValAStateBlocked
		}
	}
	if strings.TrimSpace(model.TriggerTable.CurrentTargetState) != strings.TrimSpace(model.Trigger.TargetState) ||
		strings.TrimSpace(model.TriggerTable.CurrentTargetState) != strings.TrimSpace(model.Reason.TargetState) ||
		strings.TrimSpace(model.TriggerTable.CurrentTargetState) != strings.TrimSpace(model.Decision.TargetState) ||
		strings.TrimSpace(model.TriggerTable.CurrentDowngradeOutcome) != strings.TrimSpace(model.Trigger.TargetDowngradeOutcome) ||
		strings.TrimSpace(model.TriggerTable.CurrentDowngradeOutcome) != strings.TrimSpace(model.Reason.TargetDowngradeOutcome) ||
		strings.TrimSpace(model.TriggerTable.CurrentDowngradeOutcome) != strings.TrimSpace(model.Decision.TargetDowngradeOutcome) {
		model.TriggerTableState = Point15ValAStateBlocked
		model.ReasonState = Point15ValAStateBlocked
		model.DecisionState = Point15ValAStateBlocked
	}

	model.CurrentState = point15ValAAggregate(
		model.DependencyState,
		model.TriggerTableState,
		model.TriggerState,
		model.ReasonState,
		model.DecisionState,
		model.AuthorityBoundaryState,
		model.NoOverclaimState,
	)
	model.BlockingReasons = point15ValABlockingReasons(model)
	model.ReviewPrerequisites = point15ValAReviewPrerequisites(model)
	return model
}

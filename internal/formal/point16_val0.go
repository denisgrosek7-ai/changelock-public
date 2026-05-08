package formal

import (
	"sort"
	"strings"
)

const (
	Point16Val0StateActive         = "point16_val0_historical_replay_governance_active"
	Point16Val0StateBlocked        = "point16_val0_historical_replay_governance_blocked"
	Point16Val0StateReviewRequired = "point16_val0_historical_replay_governance_review_required"
	Point16Val0StateIncomplete     = "point16_val0_historical_replay_governance_incomplete"
)

const (
	point16Val0PointID          = "point_16"
	point16Val0WaveID           = "val_0"
	point16Val0Scope            = "continuous_trust_evolution_and_historical_replay_governance_foundation"
	point16Val0ReplayDisclaimer = "historical_replay_governance_foundation no_final_point16_closure point16_val0"

	point16Val0ReplayBound                  = "replay_bound"
	point16Val0HistoricalContextComplete    = "historical_context_complete"
	point16Val0OriginalContextMissing       = "original_context_missing"
	point16Val0OriginalContextIncomplete    = "original_context_incomplete"
	point16Val0OriginalEvidenceMissing      = "original_evidence_missing"
	point16Val0EvidenceHashMismatch         = "evidence_hash_mismatch"
	point16Val0PolicyHashMismatch           = "policy_hash_mismatch"
	point16Val0EngineHashMismatch           = "engine_hash_mismatch"
	point16Val0TenantScopeMismatch          = "tenant_scope_mismatch"
	point16Val0ArtifactScopeMismatch        = "artifact_scope_mismatch"
	point16Val0ClaimScopeMismatch           = "claim_scope_mismatch"
	point16Val0GovernanceScopeMismatch      = "governance_scope_mismatch"
	point16Val0TimestampMismatch            = "timestamp_mismatch"
	point16Val0TimestampUnsafe              = "timestamp_unsafe"
	point16Val0LineageMissing               = "lineage_missing"
	point16Val0UnsupportedReplay            = "unsupported_replay"
	point16Val0TamperedHistory              = "tampered_history"
	point16Val0CurrentPolicySubstitution    = "current_policy_substitution_attempt"
	point16Val0CurrentEngineSubstitution    = "current_engine_substitution_attempt"
	point16Val0CurrentEvidenceSubstitution  = "current_evidence_substitution_attempt"
	point16Val0CurrentTimeSubstitution      = "current_time_substitution_attempt"
	point16Val0ReplayTaxonomyReviewRequired = "review_required"
	point16Val0ReplayTaxonomyBlocked        = "blocked"
	point16Val0ReplayTaxonomyIncomplete     = "incomplete"
)

type Point16Val0DependencySnapshot struct {
	Point15ValECurrentState          string                                             `json:"point15_vale_current_state"`
	Point15ValEDependencyState       string                                             `json:"point15_vale_dependency_state"`
	Point15ValEClosureEvaluatorState string                                             `json:"point15_vale_closure_evaluator_state"`
	Point15ValEPassClosureState      string                                             `json:"point15_vale_pass_closure_manifest_state"`
	Point15ValEComputedFromUpstream  bool                                               `json:"point15_vale_computed_from_upstream"`
	Point15ValEMerged                bool                                               `json:"point15_vale_merged"`
	Point15ValECIGreen               bool                                               `json:"point15_vale_ci_green"`
	Point15ValEReviewedOnMain        bool                                               `json:"point15_vale_reviewed_on_main"`
	Point15PassAllowed               bool                                               `json:"point15_pass_allowed"`
	Point15PassToken                 string                                             `json:"point15_pass_token"`
	Point15PassManifestPointID       string                                             `json:"point15_pass_manifest_point_id"`
	Point15PassManifestWaveID        string                                             `json:"point15_pass_manifest_wave_id"`
	Point15PassManifestClosureToken  string                                             `json:"point15_pass_manifest_closure_token"`
	InheritedTenantScope             string                                             `json:"inherited_tenant_scope"`
	SnapshotFromComputedOutput       bool                                               `json:"snapshot_from_computed_output"`
	ReviewPrerequisites              []string                                           `json:"review_prerequisites,omitempty"`
	Point15ValE                      Point15ValEContinuousVerificationClosureFoundation `json:"point15_vale"`
}

type Point16Val0HistoricalReplayContext struct {
	ContextID                   string `json:"context_id"`
	OriginalEvidenceID          string `json:"original_evidence_id"`
	OriginalEvidenceHash        string `json:"original_evidence_hash"`
	OriginalPolicyID            string `json:"original_policy_id"`
	OriginalPolicyVersion       string `json:"original_policy_version"`
	OriginalPolicyHash          string `json:"original_policy_hash"`
	OriginalEngineID            string `json:"original_engine_id"`
	OriginalEngineVersion       string `json:"original_engine_version"`
	OriginalEngineHash          string `json:"original_engine_hash"`
	OriginalTenantScope         string `json:"original_tenant_scope"`
	OriginalArtifactScope       string `json:"original_artifact_scope"`
	OriginalClaimScope          string `json:"original_claim_scope"`
	OriginalGovernanceScope     string `json:"original_governance_scope"`
	OriginalDecisionAt          string `json:"original_decision_at"`
	OriginalDecisionTimeSource  string `json:"original_decision_time_source"`
	OriginalEvaluatedAt         string `json:"original_evaluated_at"`
	OriginalEvaluatedTimeSource string `json:"original_evaluated_time_source"`
	ReplayAt                    string `json:"replay_at"`
	ReplayTimeSource            string `json:"replay_time_source"`
	LineageRef                  string `json:"lineage_ref"`
}

type Point16Val0OriginalDecisionBinding struct {
	BindingID                    string `json:"binding_id"`
	HistoricalReplayContextRef   string `json:"historical_replay_context_ref"`
	OriginalDecisionID           string `json:"original_decision_id"`
	OriginalDecisionHash         string `json:"original_decision_hash"`
	OriginalEvidenceID           string `json:"original_evidence_id"`
	OriginalEvidenceHash         string `json:"original_evidence_hash"`
	OriginalPolicyID             string `json:"original_policy_id"`
	OriginalPolicyVersion        string `json:"original_policy_version"`
	OriginalPolicyHash           string `json:"original_policy_hash"`
	OriginalEngineID             string `json:"original_engine_id"`
	OriginalEngineVersion        string `json:"original_engine_version"`
	OriginalEngineHash           string `json:"original_engine_hash"`
	OriginalTenantScope          string `json:"original_tenant_scope"`
	OriginalArtifactScope        string `json:"original_artifact_scope"`
	OriginalClaimScope           string `json:"original_claim_scope"`
	OriginalGovernanceScope      string `json:"original_governance_scope"`
	OriginalDecisionAt           string `json:"original_decision_at"`
	OriginalEvaluatedAt          string `json:"original_evaluated_at"`
	CurrentPolicyVersion         string `json:"current_policy_version"`
	CurrentPolicyHash            string `json:"current_policy_hash"`
	CurrentEngineVersion         string `json:"current_engine_version"`
	CurrentEngineHash            string `json:"current_engine_hash"`
	CurrentEvidenceHash          string `json:"current_evidence_hash"`
	CurrentTenantScope           string `json:"current_tenant_scope"`
	CurrentArtifactScope         string `json:"current_artifact_scope"`
	CurrentClaimScope            string `json:"current_claim_scope"`
	CurrentGovernanceScope       string `json:"current_governance_scope"`
	CurrentContextComparisonOnly bool   `json:"current_context_comparison_only"`
	LineageRef                   string `json:"lineage_ref"`
}

type Point16Val0ReplayTaxonomy struct {
	TaxonomyID                           string   `json:"taxonomy_id"`
	ReplayStatus                         string   `json:"replay_status"`
	AllowedStatuses                      []string `json:"allowed_statuses,omitempty"`
	OriginalContextComplete              bool     `json:"original_context_complete"`
	OriginalEvidencePresent              bool     `json:"original_evidence_present"`
	OriginalPolicyPresent                bool     `json:"original_policy_present"`
	OriginalEnginePresent                bool     `json:"original_engine_present"`
	EvidenceHashMatches                  bool     `json:"evidence_hash_matches"`
	PolicyHashMatches                    bool     `json:"policy_hash_matches"`
	EngineHashMatches                    bool     `json:"engine_hash_matches"`
	TenantScopeMatches                   bool     `json:"tenant_scope_matches"`
	ArtifactScopeMatches                 bool     `json:"artifact_scope_matches"`
	ClaimScopeMatches                    bool     `json:"claim_scope_matches"`
	GovernanceScopeMatches               bool     `json:"governance_scope_matches"`
	TimestampConsistent                  bool     `json:"timestamp_consistent"`
	TimestampSafe                        bool     `json:"timestamp_safe"`
	LineagePresent                       bool     `json:"lineage_present"`
	ReplaySupported                      bool     `json:"replay_supported"`
	HistoryTampered                      bool     `json:"history_tampered"`
	CurrentPolicySubstitutionAttempted   bool     `json:"current_policy_substitution_attempted"`
	CurrentEngineSubstitutionAttempted   bool     `json:"current_engine_substitution_attempted"`
	CurrentEvidenceSubstitutionAttempted bool     `json:"current_evidence_substitution_attempted"`
	CurrentTimeSubstitutionAttempted     bool     `json:"current_time_substitution_attempted"`
}

type Point16Val0CurrentSubstitutionGuard struct {
	GuardID                        string `json:"guard_id"`
	OriginalPolicyHash             string `json:"original_policy_hash"`
	CurrentPolicyHash              string `json:"current_policy_hash"`
	OriginalEngineHash             string `json:"original_engine_hash"`
	CurrentEngineHash              string `json:"current_engine_hash"`
	OriginalEvidenceHash           string `json:"original_evidence_hash"`
	CurrentEvidenceHash            string `json:"current_evidence_hash"`
	OriginalDecisionAt             string `json:"original_decision_at"`
	ReplayAt                       string `json:"replay_at"`
	OriginalTenantScope            string `json:"original_tenant_scope"`
	CurrentTenantScope             string `json:"current_tenant_scope"`
	OriginalClaimScope             string `json:"original_claim_scope"`
	CurrentClaimScope              string `json:"current_claim_scope"`
	OriginalGovernanceScope        string `json:"original_governance_scope"`
	CurrentGovernanceScope         string `json:"current_governance_scope"`
	CurrentPolicyAuthoritative     bool   `json:"current_policy_authoritative"`
	CurrentEngineAuthoritative     bool   `json:"current_engine_authoritative"`
	CurrentEvidenceAuthoritative   bool   `json:"current_evidence_authoritative"`
	CurrentTimestampAuthoritative  bool   `json:"current_timestamp_authoritative"`
	CurrentTenantAuthoritative     bool   `json:"current_tenant_authoritative"`
	CurrentClaimAuthoritative      bool   `json:"current_claim_authoritative"`
	CurrentGovernanceAuthoritative bool   `json:"current_governance_authoritative"`
}

type Point16Val0ReplayReadinessEvaluation struct {
	EvaluationID                  string `json:"evaluation_id"`
	DependencyState               string `json:"dependency_state"`
	HistoricalReplayContextState  string `json:"historical_replay_context_state"`
	OriginalDecisionBindingState  string `json:"original_decision_binding_state"`
	ReplayTaxonomyState           string `json:"replay_taxonomy_state"`
	CurrentSubstitutionGuardState string `json:"current_substitution_guard_state"`
	NoOverclaimState              string `json:"no_overclaim_state"`
	OriginalContextExactBound     bool   `json:"original_context_exact_bound"`
	TenantSafe                    bool   `json:"tenant_safe"`
	TimestampSafe                 bool   `json:"timestamp_safe"`
	HashSafe                      bool   `json:"hash_safe"`
	GovernanceSafe                bool   `json:"governance_safe"`
	LineageSafe                   bool   `json:"lineage_safe"`
	ReplaySupported               bool   `json:"replay_supported"`
	NoMutationPathsDetected       bool   `json:"no_mutation_paths_detected"`
	NoPublicationPathDetected     bool   `json:"no_publication_path_detected"`
	NoRevocationExecutionDetected bool   `json:"no_revocation_execution_detected"`
	NoEvidenceDeletionDetected    bool   `json:"no_evidence_deletion_detected"`
	NoExternalAPIDefaultDetected  bool   `json:"no_external_api_default_detected"`
	NoConnectorMutationDetected   bool   `json:"no_connector_mutation_detected"`
	NoSchedulerPathDetected       bool   `json:"no_scheduler_path_detected"`
	NoCustomerAuthorityDetected   bool   `json:"no_customer_authority_detected"`
	NoAuditorAuthorityDetected    bool   `json:"no_auditor_authority_detected"`
	NoPortalAuthorityDetected     bool   `json:"no_portal_authority_detected"`
	NoAIAgentAuthorityDetected    bool   `json:"no_ai_agent_authority_detected"`
	NoPrematureFinalTokenDetected bool   `json:"no_premature_final_token_detected"`
}

type Point16Val0NoOverclaimBaseline struct {
	ObservedTexts                        []string `json:"observed_texts,omitempty"`
	InternalDiagnosticTexts              []string `json:"internal_diagnostic_texts,omitempty"`
	InternalDiagnosticsClassifiedBlocked bool     `json:"internal_diagnostics_classified_blocked"`
	AllowedSafeWording                   []string `json:"allowed_safe_wording,omitempty"`
	BlockedWording                       []string `json:"blocked_wording,omitempty"`
	ReplayDisclaimer                     string   `json:"replay_disclaimer"`
}

type Point16Val0Foundation struct {
	CurrentState                  string                               `json:"current_state"`
	BlockingReasons               []string                             `json:"blocking_reasons,omitempty"`
	ReviewPrerequisites           []string                             `json:"review_prerequisites,omitempty"`
	ReplayDisclaimer              string                               `json:"replay_disclaimer"`
	DependencyState               string                               `json:"dependency_state"`
	HistoricalReplayContextState  string                               `json:"historical_replay_context_state"`
	OriginalDecisionBindingState  string                               `json:"original_decision_binding_state"`
	ReplayTaxonomyState           string                               `json:"replay_taxonomy_state"`
	CurrentSubstitutionGuardState string                               `json:"current_substitution_guard_state"`
	ReplayReadinessState          string                               `json:"replay_readiness_state"`
	NoOverclaimState              string                               `json:"no_overclaim_state"`
	Dependency                    Point16Val0DependencySnapshot        `json:"dependency"`
	HistoricalReplayContext       Point16Val0HistoricalReplayContext   `json:"historical_replay_context"`
	OriginalDecisionBinding       Point16Val0OriginalDecisionBinding   `json:"original_decision_binding"`
	ReplayTaxonomy                Point16Val0ReplayTaxonomy            `json:"replay_taxonomy"`
	CurrentSubstitutionGuard      Point16Val0CurrentSubstitutionGuard  `json:"current_substitution_guard"`
	ReplayReadinessEvaluation     Point16Val0ReplayReadinessEvaluation `json:"replay_readiness_evaluation"`
	NoOverclaimBaseline           Point16Val0NoOverclaimBaseline       `json:"no_overclaim_baseline"`
}

func point16Val0States() []string {
	return []string{
		Point16Val0StateActive,
		Point16Val0StateBlocked,
		Point16Val0StateReviewRequired,
		Point16Val0StateIncomplete,
	}
}

func point16Val0StateValid(value string) bool {
	return point14Val0ExactValueValid(value, point16Val0States())
}

func point16Val0ReplayStatuses() []string {
	return []string{
		point16Val0ReplayBound,
		point16Val0HistoricalContextComplete,
		point16Val0OriginalContextMissing,
		point16Val0OriginalContextIncomplete,
		point16Val0OriginalEvidenceMissing,
		point16Val0EvidenceHashMismatch,
		point16Val0PolicyHashMismatch,
		point16Val0EngineHashMismatch,
		point16Val0TenantScopeMismatch,
		point16Val0ArtifactScopeMismatch,
		point16Val0ClaimScopeMismatch,
		point16Val0GovernanceScopeMismatch,
		point16Val0TimestampMismatch,
		point16Val0TimestampUnsafe,
		point16Val0LineageMissing,
		point16Val0UnsupportedReplay,
		point16Val0TamperedHistory,
		point16Val0CurrentPolicySubstitution,
		point16Val0CurrentEngineSubstitution,
		point16Val0CurrentEvidenceSubstitution,
		point16Val0CurrentTimeSubstitution,
		point16Val0ReplayTaxonomyReviewRequired,
		point16Val0ReplayTaxonomyBlocked,
		point16Val0ReplayTaxonomyIncomplete,
	}
}

func point16Val0ReplayStatusValid(value string) bool {
	return point14Val0ExactValueValid(value, point16Val0ReplayStatuses())
}

func point16Val0ForbiddenWording() []string {
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
		"legal proof",
		"financial guarantee",
	}
}

func point16Val0SafeWording() []string {
	return []string{
		"historical replay remains bound to original decision context",
		"current context comparison cannot replace original context",
		"unsupported or incomplete replay fails closed",
		"point 16 val 0 verifies replay governance foundation only",
	}
}

func point16Val0ObservedTextContainsForbiddenWording(text string) bool {
	trimmed := strings.TrimSpace(strings.ToLower(text))
	if trimmed == "" {
		return false
	}
	for _, phrase := range point16Val0ForbiddenWording() {
		if strings.Contains(trimmed, strings.ToLower(phrase)) {
			return true
		}
	}
	return false
}

func point16Val0ObservedListContainsForbiddenWording(values []string) bool {
	for _, value := range values {
		if point16Val0ObservedTextContainsForbiddenWording(value) {
			return true
		}
	}
	return false
}

func point16Val0DependencyRefValid(value string) bool {
	return point14Val0RefValid(value, "point16_val0_", "historical_replay_", "original_decision_", "replay_", "substitution_", "readiness_")
}

func point16Val0EvidenceIdentityValid(value string) bool {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return false
	}
	if point15Val0EvidenceIDValid(trimmed) {
		return true
	}
	return strings.Contains(trimmed, "evidence_id=") &&
		strings.Contains(trimmed, "evidence_hash=") &&
		strings.Contains(trimmed, "policy=") &&
		strings.Contains(trimmed, "engine=") &&
		strings.Contains(trimmed, "tenant=")
}

func point16Val0HistoricalEvidenceHashValid(value string) bool {
	return point12Val0HashValid(value) || point14Val0RefValid(value, "hash_", "evidence_hash_")
}

func point16Val0ContextRefValid(value string) bool {
	return point14Val0RefValid(value, "historical_replay_context_", "replay_context_")
}

func point16Val0DecisionIDValid(value string) bool {
	return point14Val0RefValid(value, "decision_")
}

func point16Val0PolicyIDValid(value string) bool {
	return point14Val0RefValid(value, "policy_")
}

func point16Val0EngineIDValid(value string) bool {
	return point14Val0RefValid(value, "engine_")
}

func point16Val0ArtifactScopeValid(value string) bool {
	return point14Val0RefValid(value, "artifact_", "artifact_scope_")
}

func point16Val0ClaimScopeValid(value string) bool {
	return point14Val0RefValid(value, "claim_", "claim_scope_")
}

func point16Val0GovernanceScopeValid(value string) bool {
	return point14Val0RefValid(value, "governance_", "governance_scope_")
}

func point16Val0LineageRefValid(value string) bool {
	return point14Val0RefValid(value, "lineage_", "replay_lineage_")
}

func point16Val0ReplayStatusState(status string) string {
	switch strings.TrimSpace(status) {
	case point16Val0ReplayBound, point16Val0HistoricalContextComplete:
		return Point16Val0StateActive
	case point16Val0PolicyHashMismatch,
		point16Val0EngineHashMismatch,
		point16Val0ArtifactScopeMismatch,
		point16Val0ClaimScopeMismatch,
		point16Val0GovernanceScopeMismatch,
		point16Val0TimestampMismatch,
		point16Val0ReplayTaxonomyReviewRequired:
		return Point16Val0StateReviewRequired
	case point16Val0OriginalContextMissing,
		point16Val0OriginalContextIncomplete,
		point16Val0OriginalEvidenceMissing,
		point16Val0LineageMissing,
		point16Val0ReplayTaxonomyIncomplete:
		return Point16Val0StateIncomplete
	default:
		return Point16Val0StateBlocked
	}
}

func point16Val0ExpectedReplayStatus(model Point16Val0ReplayTaxonomy) string {
	if model.HistoryTampered {
		return point16Val0TamperedHistory
	}
	if !model.ReplaySupported {
		return point16Val0UnsupportedReplay
	}
	if model.CurrentPolicySubstitutionAttempted {
		return point16Val0CurrentPolicySubstitution
	}
	if model.CurrentEngineSubstitutionAttempted {
		return point16Val0CurrentEngineSubstitution
	}
	if model.CurrentEvidenceSubstitutionAttempted {
		return point16Val0CurrentEvidenceSubstitution
	}
	if model.CurrentTimeSubstitutionAttempted {
		return point16Val0CurrentTimeSubstitution
	}
	if !model.OriginalEvidencePresent {
		return point16Val0OriginalEvidenceMissing
	}
	if !model.OriginalPolicyPresent && !model.OriginalEnginePresent {
		return point16Val0OriginalContextMissing
	}
	if !model.OriginalContextComplete || !model.OriginalPolicyPresent || !model.OriginalEnginePresent {
		return point16Val0OriginalContextIncomplete
	}
	if !model.TimestampSafe {
		return point16Val0TimestampUnsafe
	}
	if !model.TimestampConsistent {
		return point16Val0TimestampMismatch
	}
	if !model.EvidenceHashMatches {
		return point16Val0EvidenceHashMismatch
	}
	if !model.TenantScopeMatches {
		return point16Val0TenantScopeMismatch
	}
	if !model.PolicyHashMatches {
		return point16Val0PolicyHashMismatch
	}
	if !model.EngineHashMatches {
		return point16Val0EngineHashMismatch
	}
	if !model.ArtifactScopeMatches {
		return point16Val0ArtifactScopeMismatch
	}
	if !model.ClaimScopeMatches {
		return point16Val0ClaimScopeMismatch
	}
	if !model.GovernanceScopeMatches {
		return point16Val0GovernanceScopeMismatch
	}
	if !model.LineagePresent {
		return point16Val0LineageMissing
	}
	if model.OriginalContextComplete {
		return point16Val0ReplayBound
	}
	return point16Val0HistoricalContextComplete
}

func point16Val0DependencySnapshotFromUpstream(valE Point15ValEContinuousVerificationClosureFoundation) Point16Val0DependencySnapshot {
	return Point16Val0DependencySnapshot{
		Point15ValECurrentState:          valE.CurrentState,
		Point15ValEDependencyState:       valE.DependencyState,
		Point15ValEClosureEvaluatorState: valE.ClosureEvaluatorState,
		Point15ValEPassClosureState:      valE.PassClosureManifestState,
		Point15ValEComputedFromUpstream:  valE.Dependency.SnapshotFromComputedOutput,
		Point15ValEMerged:                true,
		Point15ValECIGreen:               true,
		Point15ValEReviewedOnMain:        true,
		Point15PassAllowed:               valE.PassClosureManifest.Point15PassAllowed,
		Point15PassToken:                 valE.PassClosureManifest.Point15PassToken,
		Point15PassManifestPointID:       valE.PassClosureManifest.PointID,
		Point15PassManifestWaveID:        valE.PassClosureManifest.WaveID,
		Point15PassManifestClosureToken:  valE.PassClosureManifest.ClosureToken,
		InheritedTenantScope:             valE.Dependency.InheritedTenantScope,
		SnapshotFromComputedOutput:       true,
		ReviewPrerequisites:              append([]string{}, valE.ReviewPrerequisites...),
		Point15ValE:                      valE,
	}
}

func point16Val0DependencySnapshotModel() Point16Val0DependencySnapshot {
	valE := ComputePoint15ValEFoundation(Point15ValEFoundationModel())
	return point16Val0DependencySnapshotFromUpstream(valE)
}

func EvaluatePoint16Val0DependencyState(model Point16Val0DependencySnapshot) string {
	if !model.SnapshotFromComputedOutput ||
		!model.Point15ValEComputedFromUpstream ||
		!model.Point15ValEMerged ||
		!model.Point15ValECIGreen ||
		!model.Point15ValEReviewedOnMain ||
		!point15ValEStateValid(model.Point15ValECurrentState) ||
		!point15ValEStateValid(model.Point15ValEDependencyState) ||
		!point15ValEStateValid(model.Point15ValEClosureEvaluatorState) ||
		!point15ValEStateValid(model.Point15ValEPassClosureState) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) {
		return Point16Val0StateBlocked
	}
	if strings.TrimSpace(model.Point15ValECurrentState) != strings.TrimSpace(model.Point15ValE.CurrentState) ||
		strings.TrimSpace(model.Point15ValEDependencyState) != strings.TrimSpace(model.Point15ValE.DependencyState) ||
		strings.TrimSpace(model.Point15ValEClosureEvaluatorState) != strings.TrimSpace(model.Point15ValE.ClosureEvaluatorState) ||
		strings.TrimSpace(model.Point15ValEPassClosureState) != strings.TrimSpace(model.Point15ValE.PassClosureManifestState) ||
		model.Point15ValEComputedFromUpstream != model.Point15ValE.Dependency.SnapshotFromComputedOutput ||
		model.Point15PassAllowed != model.Point15ValE.PassClosureManifest.Point15PassAllowed ||
		strings.TrimSpace(model.Point15PassToken) != strings.TrimSpace(model.Point15ValE.PassClosureManifest.Point15PassToken) ||
		strings.TrimSpace(model.Point15PassManifestPointID) != strings.TrimSpace(model.Point15ValE.PassClosureManifest.PointID) ||
		strings.TrimSpace(model.Point15PassManifestWaveID) != strings.TrimSpace(model.Point15ValE.PassClosureManifest.WaveID) ||
		strings.TrimSpace(model.Point15PassManifestClosureToken) != strings.TrimSpace(model.Point15ValE.PassClosureManifest.ClosureToken) ||
		strings.TrimSpace(model.InheritedTenantScope) != strings.TrimSpace(model.Point15ValE.Dependency.InheritedTenantScope) {
		return Point16Val0StateBlocked
	}
	if strings.TrimSpace(model.Point15ValECurrentState) != Point15ValEStatePassConfirmed ||
		strings.TrimSpace(model.Point15ValEDependencyState) != Point15ValEStatePassConfirmed ||
		strings.TrimSpace(model.Point15ValEClosureEvaluatorState) != Point15ValEStatePassConfirmed ||
		strings.TrimSpace(model.Point15ValEPassClosureState) != Point15ValEStatePassConfirmed ||
		!model.Point15PassAllowed ||
		strings.TrimSpace(model.Point15PassToken) != point15Val0BlockedPassToken ||
		strings.TrimSpace(model.Point15PassManifestPointID) != point15Val0PointID ||
		strings.TrimSpace(model.Point15PassManifestWaveID) != point15ValEWaveID ||
		strings.TrimSpace(model.Point15PassManifestClosureToken) != point15Val0BlockedPassToken {
		return Point16Val0StateBlocked
	}
	return Point16Val0StateActive
}

func point16Val0HistoricalReplayContextModel(dependency Point16Val0DependencySnapshot) Point16Val0HistoricalReplayContext {
	return Point16Val0HistoricalReplayContext{
		ContextID:                   "historical_replay_context_point16_val0_001",
		OriginalEvidenceID:          dependency.Point15ValE.PassClosureManifest.EvidenceIdentity,
		OriginalEvidenceHash:        dependency.Point15ValE.PassClosureManifest.EvidenceHash,
		OriginalPolicyID:            "policy_historical_replay_001",
		OriginalPolicyVersion:       dependency.Point15ValE.PassClosureManifest.PolicyVersion,
		OriginalPolicyHash:          "sha256:1111111111111111111111111111111111111111111111111111111111111111",
		OriginalEngineID:            "engine_historical_replay_001",
		OriginalEngineVersion:       dependency.Point15ValE.PassClosureManifest.EngineVersion,
		OriginalEngineHash:          "sha256:2222222222222222222222222222222222222222222222222222222222222222",
		OriginalTenantScope:         dependency.InheritedTenantScope,
		OriginalArtifactScope:       "artifact_scope_historical_replay_001",
		OriginalClaimScope:          "claim_scope_historical_replay_001",
		OriginalGovernanceScope:     "governance_scope_historical_replay_001",
		OriginalDecisionAt:          "2026-05-08T09:00:00Z",
		OriginalDecisionTimeSource:  point14Val0TimeSourceServerUTC,
		OriginalEvaluatedAt:         "2026-05-08T09:05:00Z",
		OriginalEvaluatedTimeSource: point14Val0TimeSourceServerUTC,
		ReplayAt:                    "2026-05-08T09:10:00Z",
		ReplayTimeSource:            point14Val0TimeSourceServerUTC,
		LineageRef:                  "lineage_point16_val0_replay_001",
	}
}

func EvaluatePoint16Val0HistoricalReplayContextState(model Point16Val0HistoricalReplayContext) string {
	if !point16Val0ContextRefValid(model.ContextID) {
		return Point16Val0StateBlocked
	}
	requiredMissing := false
	for _, value := range []string{
		model.OriginalEvidenceID,
		model.OriginalEvidenceHash,
		model.OriginalPolicyID,
		model.OriginalPolicyVersion,
		model.OriginalPolicyHash,
		model.OriginalEngineID,
		model.OriginalEngineVersion,
		model.OriginalEngineHash,
		model.OriginalTenantScope,
		model.OriginalArtifactScope,
		model.OriginalClaimScope,
		model.OriginalGovernanceScope,
		model.OriginalDecisionAt,
		model.OriginalEvaluatedAt,
		model.ReplayAt,
		model.LineageRef,
	} {
		if strings.TrimSpace(value) == "" {
			requiredMissing = true
		}
	}
	if requiredMissing {
		return Point16Val0StateIncomplete
	}
	if !point16Val0EvidenceIdentityValid(model.OriginalEvidenceID) ||
		!point16Val0HistoricalEvidenceHashValid(model.OriginalEvidenceHash) ||
		!point16Val0PolicyIDValid(model.OriginalPolicyID) ||
		!point12Val0VersionIdentityValid(model.OriginalPolicyVersion) ||
		!point12Val0HashValid(model.OriginalPolicyHash) ||
		!point16Val0EngineIDValid(model.OriginalEngineID) ||
		!point12Val0VersionIdentityValid(model.OriginalEngineVersion) ||
		!point12Val0HashValid(model.OriginalEngineHash) ||
		!point11Val0ScopeValid(model.OriginalTenantScope) ||
		!point16Val0ArtifactScopeValid(model.OriginalArtifactScope) ||
		!point16Val0ClaimScopeValid(model.OriginalClaimScope) ||
		!point16Val0GovernanceScopeValid(model.OriginalGovernanceScope) ||
		!point14Val0ParsedTimeOk(model.OriginalDecisionAt) ||
		!point14Val0ParsedTimeOk(model.OriginalEvaluatedAt) ||
		!point14Val0ParsedTimeOk(model.ReplayAt) ||
		!point14Val0CanonicalTimeSourceValid(model.OriginalDecisionTimeSource) ||
		!point14Val0CanonicalTimeSourceValid(model.OriginalEvaluatedTimeSource) ||
		!point14Val0CanonicalTimeSourceValid(model.ReplayTimeSource) ||
		!point16Val0LineageRefValid(model.LineageRef) {
		return Point16Val0StateBlocked
	}
	decisionAt, _ := point14Val0ParsedTime(model.OriginalDecisionAt)
	evaluatedAt, _ := point14Val0ParsedTime(model.OriginalEvaluatedAt)
	replayAt, _ := point14Val0ParsedTime(model.ReplayAt)
	if decisionAt.After(evaluatedAt) || evaluatedAt.After(replayAt) || decisionAt.After(replayAt) {
		return Point16Val0StateBlocked
	}
	return Point16Val0StateActive
}

func point16Val0OriginalDecisionBindingModel(context Point16Val0HistoricalReplayContext) Point16Val0OriginalDecisionBinding {
	return Point16Val0OriginalDecisionBinding{
		BindingID:                    "original_decision_binding_point16_val0_001",
		HistoricalReplayContextRef:   context.ContextID,
		OriginalDecisionID:           "decision_point16_val0_original_001",
		OriginalDecisionHash:         "sha256:3333333333333333333333333333333333333333333333333333333333333333",
		OriginalEvidenceID:           context.OriginalEvidenceID,
		OriginalEvidenceHash:         context.OriginalEvidenceHash,
		OriginalPolicyID:             context.OriginalPolicyID,
		OriginalPolicyVersion:        context.OriginalPolicyVersion,
		OriginalPolicyHash:           context.OriginalPolicyHash,
		OriginalEngineID:             context.OriginalEngineID,
		OriginalEngineVersion:        context.OriginalEngineVersion,
		OriginalEngineHash:           context.OriginalEngineHash,
		OriginalTenantScope:          context.OriginalTenantScope,
		OriginalArtifactScope:        context.OriginalArtifactScope,
		OriginalClaimScope:           context.OriginalClaimScope,
		OriginalGovernanceScope:      context.OriginalGovernanceScope,
		OriginalDecisionAt:           context.OriginalDecisionAt,
		OriginalEvaluatedAt:          context.OriginalEvaluatedAt,
		CurrentPolicyVersion:         context.OriginalPolicyVersion,
		CurrentPolicyHash:            context.OriginalPolicyHash,
		CurrentEngineVersion:         context.OriginalEngineVersion,
		CurrentEngineHash:            context.OriginalEngineHash,
		CurrentEvidenceHash:          context.OriginalEvidenceHash,
		CurrentTenantScope:           context.OriginalTenantScope,
		CurrentArtifactScope:         context.OriginalArtifactScope,
		CurrentClaimScope:            context.OriginalClaimScope,
		CurrentGovernanceScope:       context.OriginalGovernanceScope,
		CurrentContextComparisonOnly: true,
		LineageRef:                   context.LineageRef,
	}
}

func EvaluatePoint16Val0OriginalDecisionBindingState(model Point16Val0OriginalDecisionBinding) string {
	if !point16Val0DependencyRefValid(model.BindingID) ||
		!point16Val0ContextRefValid(model.HistoricalReplayContextRef) ||
		!point16Val0DecisionIDValid(model.OriginalDecisionID) ||
		!point12Val0HashValid(model.OriginalDecisionHash) {
		return Point16Val0StateBlocked
	}
	requiredMissing := false
	for _, value := range []string{
		model.OriginalEvidenceID,
		model.OriginalEvidenceHash,
		model.OriginalPolicyID,
		model.OriginalPolicyVersion,
		model.OriginalPolicyHash,
		model.OriginalEngineID,
		model.OriginalEngineVersion,
		model.OriginalEngineHash,
		model.OriginalTenantScope,
		model.OriginalArtifactScope,
		model.OriginalClaimScope,
		model.OriginalGovernanceScope,
		model.OriginalDecisionAt,
		model.OriginalEvaluatedAt,
		model.LineageRef,
	} {
		if strings.TrimSpace(value) == "" {
			requiredMissing = true
		}
	}
	if requiredMissing {
		return Point16Val0StateIncomplete
	}
	if !point16Val0EvidenceIdentityValid(model.OriginalEvidenceID) ||
		!point16Val0HistoricalEvidenceHashValid(model.OriginalEvidenceHash) ||
		!point16Val0PolicyIDValid(model.OriginalPolicyID) ||
		!point12Val0VersionIdentityValid(model.OriginalPolicyVersion) ||
		!point12Val0HashValid(model.OriginalPolicyHash) ||
		!point16Val0EngineIDValid(model.OriginalEngineID) ||
		!point12Val0VersionIdentityValid(model.OriginalEngineVersion) ||
		!point12Val0HashValid(model.OriginalEngineHash) ||
		!point11Val0ScopeValid(model.OriginalTenantScope) ||
		!point16Val0ArtifactScopeValid(model.OriginalArtifactScope) ||
		!point16Val0ClaimScopeValid(model.OriginalClaimScope) ||
		!point16Val0GovernanceScopeValid(model.OriginalGovernanceScope) ||
		!point14Val0ParsedTimeOk(model.OriginalDecisionAt) ||
		!point14Val0ParsedTimeOk(model.OriginalEvaluatedAt) ||
		!point16Val0LineageRefValid(model.LineageRef) {
		return Point16Val0StateBlocked
	}
	if !model.CurrentContextComparisonOnly {
		return Point16Val0StateBlocked
	}
	return Point16Val0StateActive
}

func point16Val0ReplayTaxonomyModel(context Point16Val0HistoricalReplayContext, binding Point16Val0OriginalDecisionBinding) Point16Val0ReplayTaxonomy {
	return Point16Val0ReplayTaxonomy{
		TaxonomyID:              "replay_taxonomy_point16_val0_001",
		ReplayStatus:            point16Val0ReplayBound,
		AllowedStatuses:         point16Val0ReplayStatuses(),
		OriginalContextComplete: true,
		OriginalEvidencePresent: true,
		OriginalPolicyPresent:   true,
		OriginalEnginePresent:   true,
		EvidenceHashMatches:     context.OriginalEvidenceHash == binding.CurrentEvidenceHash,
		PolicyHashMatches:       context.OriginalPolicyHash == binding.CurrentPolicyHash,
		EngineHashMatches:       context.OriginalEngineHash == binding.CurrentEngineHash,
		TenantScopeMatches:      context.OriginalTenantScope == binding.CurrentTenantScope,
		ArtifactScopeMatches:    context.OriginalArtifactScope == binding.CurrentArtifactScope,
		ClaimScopeMatches:       context.OriginalClaimScope == binding.CurrentClaimScope,
		GovernanceScopeMatches:  context.OriginalGovernanceScope == binding.CurrentGovernanceScope,
		TimestampConsistent:     strings.TrimSpace(context.OriginalDecisionAt) == strings.TrimSpace(binding.OriginalDecisionAt) && strings.TrimSpace(context.OriginalEvaluatedAt) == strings.TrimSpace(binding.OriginalEvaluatedAt),
		TimestampSafe:           true,
		LineagePresent:          point16Val0LineageRefValid(binding.LineageRef),
		ReplaySupported:         true,
	}
}

func EvaluatePoint16Val0ReplayTaxonomyState(model Point16Val0ReplayTaxonomy) string {
	if !point16Val0DependencyRefValid(model.TaxonomyID) ||
		!point16Val0ReplayStatusValid(model.ReplayStatus) ||
		!point12Val0ExactStringSetMatch(model.AllowedStatuses, point16Val0ReplayStatuses()) {
		return Point16Val0StateBlocked
	}
	expectedStatus := point16Val0ExpectedReplayStatus(model)
	if strings.TrimSpace(model.ReplayStatus) != expectedStatus {
		return Point16Val0StateBlocked
	}
	return point16Val0ReplayStatusState(expectedStatus)
}

func point16Val0CurrentSubstitutionGuardModel(context Point16Val0HistoricalReplayContext, binding Point16Val0OriginalDecisionBinding) Point16Val0CurrentSubstitutionGuard {
	return Point16Val0CurrentSubstitutionGuard{
		GuardID:                 "substitution_guard_point16_val0_001",
		OriginalPolicyHash:      context.OriginalPolicyHash,
		CurrentPolicyHash:       binding.CurrentPolicyHash,
		OriginalEngineHash:      context.OriginalEngineHash,
		CurrentEngineHash:       binding.CurrentEngineHash,
		OriginalEvidenceHash:    context.OriginalEvidenceHash,
		CurrentEvidenceHash:     binding.CurrentEvidenceHash,
		OriginalDecisionAt:      context.OriginalDecisionAt,
		ReplayAt:                context.ReplayAt,
		OriginalTenantScope:     context.OriginalTenantScope,
		CurrentTenantScope:      binding.CurrentTenantScope,
		OriginalClaimScope:      context.OriginalClaimScope,
		CurrentClaimScope:       binding.CurrentClaimScope,
		OriginalGovernanceScope: context.OriginalGovernanceScope,
		CurrentGovernanceScope:  binding.CurrentGovernanceScope,
	}
}

func EvaluatePoint16Val0CurrentSubstitutionGuardState(model Point16Val0CurrentSubstitutionGuard) string {
	if !point16Val0DependencyRefValid(model.GuardID) ||
		!point12Val0HashValid(model.OriginalPolicyHash) ||
		!point12Val0HashValid(model.CurrentPolicyHash) ||
		!point12Val0HashValid(model.OriginalEngineHash) ||
		!point12Val0HashValid(model.CurrentEngineHash) ||
		!point16Val0HistoricalEvidenceHashValid(model.OriginalEvidenceHash) ||
		!point16Val0HistoricalEvidenceHashValid(model.CurrentEvidenceHash) ||
		!point14Val0ParsedTimeOk(model.OriginalDecisionAt) ||
		!point14Val0ParsedTimeOk(model.ReplayAt) ||
		!point11Val0ScopeValid(model.OriginalTenantScope) ||
		!point11Val0ScopeValid(model.CurrentTenantScope) ||
		!point16Val0ClaimScopeValid(model.OriginalClaimScope) ||
		!point16Val0ClaimScopeValid(model.CurrentClaimScope) ||
		!point16Val0GovernanceScopeValid(model.OriginalGovernanceScope) ||
		!point16Val0GovernanceScopeValid(model.CurrentGovernanceScope) {
		return Point16Val0StateBlocked
	}
	if model.CurrentPolicyAuthoritative ||
		model.CurrentEngineAuthoritative ||
		model.CurrentEvidenceAuthoritative ||
		model.CurrentTimestampAuthoritative ||
		model.CurrentTenantAuthoritative ||
		model.CurrentClaimAuthoritative ||
		model.CurrentGovernanceAuthoritative {
		return Point16Val0StateBlocked
	}
	return Point16Val0StateActive
}

func point16Val0ReplayReadinessEvaluationModel() Point16Val0ReplayReadinessEvaluation {
	return Point16Val0ReplayReadinessEvaluation{
		EvaluationID:                  "readiness_evaluation_point16_val0_001",
		OriginalContextExactBound:     true,
		TenantSafe:                    true,
		TimestampSafe:                 true,
		HashSafe:                      true,
		GovernanceSafe:                true,
		LineageSafe:                   true,
		ReplaySupported:               true,
		NoMutationPathsDetected:       true,
		NoPublicationPathDetected:     true,
		NoRevocationExecutionDetected: true,
		NoEvidenceDeletionDetected:    true,
		NoExternalAPIDefaultDetected:  true,
		NoConnectorMutationDetected:   true,
		NoSchedulerPathDetected:       true,
		NoCustomerAuthorityDetected:   true,
		NoAuditorAuthorityDetected:    true,
		NoPortalAuthorityDetected:     true,
		NoAIAgentAuthorityDetected:    true,
		NoPrematureFinalTokenDetected: true,
	}
}

func EvaluatePoint16Val0ReplayReadinessState(model Point16Val0ReplayReadinessEvaluation) string {
	if !point16Val0DependencyRefValid(model.EvaluationID) ||
		!point16Val0StateValid(model.DependencyState) ||
		!point16Val0StateValid(model.HistoricalReplayContextState) ||
		!point16Val0StateValid(model.OriginalDecisionBindingState) ||
		!point16Val0StateValid(model.ReplayTaxonomyState) ||
		!point16Val0StateValid(model.CurrentSubstitutionGuardState) ||
		!point16Val0StateValid(model.NoOverclaimState) {
		return Point16Val0StateBlocked
	}
	componentState := point16Val0Aggregate(
		model.DependencyState,
		model.HistoricalReplayContextState,
		model.OriginalDecisionBindingState,
		model.ReplayTaxonomyState,
		model.CurrentSubstitutionGuardState,
		model.NoOverclaimState,
	)
	if componentState == Point16Val0StateBlocked {
		return Point16Val0StateBlocked
	}
	if !model.OriginalContextExactBound ||
		!model.TenantSafe ||
		!model.TimestampSafe ||
		!model.HashSafe ||
		!model.GovernanceSafe ||
		!model.LineageSafe ||
		!model.ReplaySupported ||
		!model.NoMutationPathsDetected ||
		!model.NoPublicationPathDetected ||
		!model.NoRevocationExecutionDetected ||
		!model.NoEvidenceDeletionDetected ||
		!model.NoExternalAPIDefaultDetected ||
		!model.NoConnectorMutationDetected ||
		!model.NoSchedulerPathDetected ||
		!model.NoCustomerAuthorityDetected ||
		!model.NoAuditorAuthorityDetected ||
		!model.NoPortalAuthorityDetected ||
		!model.NoAIAgentAuthorityDetected ||
		!model.NoPrematureFinalTokenDetected {
		return Point16Val0StateBlocked
	}
	return componentState
}

func point16Val0NoOverclaimBaselineModel() Point16Val0NoOverclaimBaseline {
	return Point16Val0NoOverclaimBaseline{
		ObservedTexts: []string{
			"historical replay remains bound to original decision context",
			"point 16 val 0 verifies replay governance foundation only",
		},
		AllowedSafeWording: point16Val0SafeWording(),
		BlockedWording:     point16Val0ForbiddenWording(),
		ReplayDisclaimer:   point16Val0ReplayDisclaimer,
	}
}

func EvaluatePoint16Val0NoOverclaimBaselineState(model Point16Val0NoOverclaimBaseline) string {
	if strings.TrimSpace(model.ReplayDisclaimer) != point16Val0ReplayDisclaimer ||
		!point12Val0ExactStringSetMatch(model.AllowedSafeWording, point16Val0SafeWording()) ||
		!point12Val0ExactStringSetMatch(model.BlockedWording, point16Val0ForbiddenWording()) {
		return Point16Val0StateBlocked
	}
	if point16Val0ObservedListContainsForbiddenWording(model.ObservedTexts) {
		return Point16Val0StateBlocked
	}
	if point16Val0ObservedListContainsForbiddenWording(model.InternalDiagnosticTexts) && !model.InternalDiagnosticsClassifiedBlocked {
		return Point16Val0StateBlocked
	}
	return Point16Val0StateActive
}

func point16Val0FoundationModelFromUpstream(valE Point15ValEContinuousVerificationClosureFoundation) Point16Val0Foundation {
	dependency := point16Val0DependencySnapshotFromUpstream(valE)
	context := point16Val0HistoricalReplayContextModel(dependency)
	binding := point16Val0OriginalDecisionBindingModel(context)
	return Point16Val0Foundation{
		ReplayDisclaimer:          point16Val0ReplayDisclaimer,
		Dependency:                dependency,
		HistoricalReplayContext:   context,
		OriginalDecisionBinding:   binding,
		ReplayTaxonomy:            point16Val0ReplayTaxonomyModel(context, binding),
		CurrentSubstitutionGuard:  point16Val0CurrentSubstitutionGuardModel(context, binding),
		ReplayReadinessEvaluation: point16Val0ReplayReadinessEvaluationModel(),
		NoOverclaimBaseline:       point16Val0NoOverclaimBaselineModel(),
	}
}

func Point16Val0FoundationModel() Point16Val0Foundation {
	valE := ComputePoint15ValEFoundation(Point15ValEFoundationModel())
	return point16Val0FoundationModelFromUpstream(valE)
}

func point16Val0Aggregate(states ...string) string {
	for _, state := range states {
		if strings.TrimSpace(state) == Point16Val0StateBlocked {
			return Point16Val0StateBlocked
		}
	}
	for _, state := range states {
		if strings.TrimSpace(state) == Point16Val0StateReviewRequired {
			return Point16Val0StateReviewRequired
		}
	}
	for _, state := range states {
		if strings.TrimSpace(state) == Point16Val0StateIncomplete {
			return Point16Val0StateIncomplete
		}
	}
	return Point16Val0StateActive
}

func point16Val0BlockingReasons(model Point16Val0Foundation) []string {
	componentStates := map[string]string{
		"dependency":                  model.DependencyState,
		"historical_replay_context":   model.HistoricalReplayContextState,
		"original_decision_binding":   model.OriginalDecisionBindingState,
		"replay_taxonomy":             model.ReplayTaxonomyState,
		"current_substitution_guard":  model.CurrentSubstitutionGuardState,
		"replay_readiness_evaluation": model.ReplayReadinessState,
		"no_overclaim":                model.NoOverclaimState,
	}
	reasons := []string{}
	for name, state := range componentStates {
		if strings.TrimSpace(state) == Point16Val0StateBlocked {
			reasons = append(reasons, name)
		}
	}
	sort.Strings(reasons)
	return reasons
}

func point16Val0ReviewPrerequisites(model Point16Val0Foundation) []string {
	componentStates := map[string]string{
		"historical_replay_context":  model.HistoricalReplayContextState,
		"original_decision_binding":  model.OriginalDecisionBindingState,
		"replay_taxonomy":            model.ReplayTaxonomyState,
		"current_substitution_guard": model.CurrentSubstitutionGuardState,
		"replay_readiness":           model.ReplayReadinessState,
		"no_overclaim":               model.NoOverclaimState,
	}
	prereqs := append([]string{}, model.Dependency.ReviewPrerequisites...)
	for name, state := range componentStates {
		if strings.TrimSpace(state) == Point16Val0StateReviewRequired || strings.TrimSpace(state) == Point16Val0StateIncomplete {
			prereqs = append(prereqs, name)
		}
	}
	sort.Strings(prereqs)
	return prereqs
}

func ComputePoint16Val0Foundation(model Point16Val0Foundation) Point16Val0Foundation {
	model.DependencyState = EvaluatePoint16Val0DependencyState(model.Dependency)
	model.HistoricalReplayContextState = EvaluatePoint16Val0HistoricalReplayContextState(model.HistoricalReplayContext)
	model.OriginalDecisionBindingState = EvaluatePoint16Val0OriginalDecisionBindingState(model.OriginalDecisionBinding)

	expectedEvidenceHashMatches := strings.TrimSpace(model.HistoricalReplayContext.OriginalEvidenceHash) == strings.TrimSpace(model.OriginalDecisionBinding.CurrentEvidenceHash) &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEvidenceHash) == strings.TrimSpace(model.CurrentSubstitutionGuard.CurrentEvidenceHash)
	expectedPolicyHashMatches := strings.TrimSpace(model.HistoricalReplayContext.OriginalPolicyHash) == strings.TrimSpace(model.OriginalDecisionBinding.CurrentPolicyHash) &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalPolicyHash) == strings.TrimSpace(model.CurrentSubstitutionGuard.CurrentPolicyHash)
	expectedEngineHashMatches := strings.TrimSpace(model.HistoricalReplayContext.OriginalEngineHash) == strings.TrimSpace(model.OriginalDecisionBinding.CurrentEngineHash) &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEngineHash) == strings.TrimSpace(model.CurrentSubstitutionGuard.CurrentEngineHash)
	expectedTenantScopeMatches := strings.TrimSpace(model.HistoricalReplayContext.OriginalTenantScope) == strings.TrimSpace(model.OriginalDecisionBinding.CurrentTenantScope) &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalTenantScope) == strings.TrimSpace(model.CurrentSubstitutionGuard.CurrentTenantScope)
	expectedArtifactScopeMatches := strings.TrimSpace(model.HistoricalReplayContext.OriginalArtifactScope) == strings.TrimSpace(model.OriginalDecisionBinding.CurrentArtifactScope)
	expectedClaimScopeMatches := strings.TrimSpace(model.HistoricalReplayContext.OriginalClaimScope) == strings.TrimSpace(model.OriginalDecisionBinding.CurrentClaimScope) &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalClaimScope) == strings.TrimSpace(model.CurrentSubstitutionGuard.CurrentClaimScope)
	expectedGovernanceScopeMatches := strings.TrimSpace(model.HistoricalReplayContext.OriginalGovernanceScope) == strings.TrimSpace(model.OriginalDecisionBinding.CurrentGovernanceScope) &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalGovernanceScope) == strings.TrimSpace(model.CurrentSubstitutionGuard.CurrentGovernanceScope)
	expectedTimestampConsistent := strings.TrimSpace(model.HistoricalReplayContext.OriginalDecisionAt) == strings.TrimSpace(model.OriginalDecisionBinding.OriginalDecisionAt) &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalDecisionAt) == strings.TrimSpace(model.CurrentSubstitutionGuard.OriginalDecisionAt) &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEvaluatedAt) == strings.TrimSpace(model.OriginalDecisionBinding.OriginalEvaluatedAt) &&
		strings.TrimSpace(model.HistoricalReplayContext.ReplayAt) == strings.TrimSpace(model.CurrentSubstitutionGuard.ReplayAt)
	expectedOriginalContextComplete := strings.TrimSpace(model.HistoricalReplayContext.OriginalEvidenceID) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEvidenceHash) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalPolicyID) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalPolicyVersion) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalPolicyHash) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEngineID) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEngineVersion) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEngineHash) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalTenantScope) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalArtifactScope) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalClaimScope) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalGovernanceScope) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalDecisionAt) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEvaluatedAt) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.ReplayAt) != ""
	if model.ReplayTaxonomy.EvidenceHashMatches != expectedEvidenceHashMatches ||
		model.ReplayTaxonomy.PolicyHashMatches != expectedPolicyHashMatches ||
		model.ReplayTaxonomy.EngineHashMatches != expectedEngineHashMatches ||
		model.ReplayTaxonomy.TenantScopeMatches != expectedTenantScopeMatches ||
		model.ReplayTaxonomy.ArtifactScopeMatches != expectedArtifactScopeMatches ||
		model.ReplayTaxonomy.ClaimScopeMatches != expectedClaimScopeMatches ||
		model.ReplayTaxonomy.GovernanceScopeMatches != expectedGovernanceScopeMatches ||
		model.ReplayTaxonomy.TimestampConsistent != expectedTimestampConsistent ||
		model.ReplayTaxonomy.OriginalContextComplete != expectedOriginalContextComplete ||
		model.ReplayTaxonomy.OriginalEvidencePresent != (strings.TrimSpace(model.HistoricalReplayContext.OriginalEvidenceID) != "" && strings.TrimSpace(model.HistoricalReplayContext.OriginalEvidenceHash) != "") ||
		model.ReplayTaxonomy.OriginalPolicyPresent != (strings.TrimSpace(model.HistoricalReplayContext.OriginalPolicyID) != "" && strings.TrimSpace(model.HistoricalReplayContext.OriginalPolicyVersion) != "" && strings.TrimSpace(model.HistoricalReplayContext.OriginalPolicyHash) != "") ||
		model.ReplayTaxonomy.OriginalEnginePresent != (strings.TrimSpace(model.HistoricalReplayContext.OriginalEngineID) != "" && strings.TrimSpace(model.HistoricalReplayContext.OriginalEngineVersion) != "" && strings.TrimSpace(model.HistoricalReplayContext.OriginalEngineHash) != "") ||
		model.ReplayTaxonomy.LineagePresent != point16Val0LineageRefValid(model.OriginalDecisionBinding.LineageRef) {
		model.ReplayTaxonomyState = Point16Val0StateBlocked
	} else {
		model.ReplayTaxonomyState = EvaluatePoint16Val0ReplayTaxonomyState(model.ReplayTaxonomy)
	}

	model.CurrentSubstitutionGuardState = EvaluatePoint16Val0CurrentSubstitutionGuardState(model.CurrentSubstitutionGuard)
	model.NoOverclaimState = EvaluatePoint16Val0NoOverclaimBaselineState(model.NoOverclaimBaseline)

	expectedTenant := strings.TrimSpace(model.Dependency.InheritedTenantScope)
	expectedEvidenceID := strings.TrimSpace(model.Dependency.Point15ValE.PassClosureManifest.EvidenceIdentity)
	expectedEvidenceHash := strings.TrimSpace(model.Dependency.Point15ValE.PassClosureManifest.EvidenceHash)
	expectedPolicyVersion := strings.TrimSpace(model.Dependency.Point15ValE.PassClosureManifest.PolicyVersion)
	expectedEngineVersion := strings.TrimSpace(model.Dependency.Point15ValE.PassClosureManifest.EngineVersion)

	if expectedTenant == "" ||
		strings.TrimSpace(model.HistoricalReplayContext.OriginalTenantScope) != expectedTenant ||
		strings.TrimSpace(model.OriginalDecisionBinding.OriginalTenantScope) != expectedTenant ||
		strings.TrimSpace(model.CurrentSubstitutionGuard.OriginalTenantScope) != expectedTenant {
		model.HistoricalReplayContextState = Point16Val0StateBlocked
		model.OriginalDecisionBindingState = Point16Val0StateBlocked
		model.CurrentSubstitutionGuardState = Point16Val0StateBlocked
	}
	if expectedEvidenceID == "" ||
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEvidenceID) != expectedEvidenceID ||
		strings.TrimSpace(model.OriginalDecisionBinding.OriginalEvidenceID) != expectedEvidenceID ||
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEvidenceHash) != expectedEvidenceHash ||
		strings.TrimSpace(model.OriginalDecisionBinding.OriginalEvidenceHash) != expectedEvidenceHash ||
		strings.TrimSpace(model.CurrentSubstitutionGuard.OriginalEvidenceHash) != expectedEvidenceHash ||
		strings.TrimSpace(model.HistoricalReplayContext.OriginalPolicyVersion) != expectedPolicyVersion ||
		strings.TrimSpace(model.OriginalDecisionBinding.OriginalPolicyVersion) != expectedPolicyVersion ||
		strings.TrimSpace(model.HistoricalReplayContext.OriginalEngineVersion) != expectedEngineVersion ||
		strings.TrimSpace(model.OriginalDecisionBinding.OriginalEngineVersion) != expectedEngineVersion {
		model.HistoricalReplayContextState = Point16Val0StateBlocked
		model.OriginalDecisionBindingState = Point16Val0StateBlocked
		model.CurrentSubstitutionGuardState = Point16Val0StateBlocked
		model.ReplayTaxonomyState = Point16Val0StateBlocked
	}
	if strings.TrimSpace(model.HistoricalReplayContext.OriginalDecisionAt) == "" ||
		strings.TrimSpace(model.OriginalDecisionBinding.OriginalDecisionAt) == "" ||
		strings.TrimSpace(model.CurrentSubstitutionGuard.OriginalDecisionAt) == "" ||
		strings.TrimSpace(model.HistoricalReplayContext.OriginalDecisionAt) != strings.TrimSpace(model.OriginalDecisionBinding.OriginalDecisionAt) ||
		strings.TrimSpace(model.HistoricalReplayContext.OriginalDecisionAt) != strings.TrimSpace(model.CurrentSubstitutionGuard.OriginalDecisionAt) {
		model.OriginalDecisionBindingState = Point16Val0StateBlocked
		model.CurrentSubstitutionGuardState = Point16Val0StateBlocked
		model.ReplayTaxonomyState = Point16Val0StateBlocked
	}

	model.ReplayReadinessEvaluation.DependencyState = model.DependencyState
	model.ReplayReadinessEvaluation.HistoricalReplayContextState = model.HistoricalReplayContextState
	model.ReplayReadinessEvaluation.OriginalDecisionBindingState = model.OriginalDecisionBindingState
	model.ReplayReadinessEvaluation.ReplayTaxonomyState = model.ReplayTaxonomyState
	model.ReplayReadinessEvaluation.CurrentSubstitutionGuardState = model.CurrentSubstitutionGuardState
	model.ReplayReadinessEvaluation.NoOverclaimState = model.NoOverclaimState
	model.ReplayReadinessEvaluation.OriginalContextExactBound = expectedOriginalContextComplete
	model.ReplayReadinessEvaluation.TenantSafe = expectedTenantScopeMatches && expectedTenant != ""
	model.ReplayReadinessEvaluation.TimestampSafe = model.ReplayTaxonomy.TimestampSafe && expectedTimestampConsistent
	model.ReplayReadinessEvaluation.HashSafe = expectedEvidenceHashMatches && expectedPolicyHashMatches && expectedEngineHashMatches
	model.ReplayReadinessEvaluation.GovernanceSafe = expectedGovernanceScopeMatches
	model.ReplayReadinessEvaluation.LineageSafe = point16Val0LineageRefValid(model.OriginalDecisionBinding.LineageRef)
	model.ReplayReadinessEvaluation.ReplaySupported = model.ReplayTaxonomy.ReplaySupported
	model.ReplayReadinessState = EvaluatePoint16Val0ReplayReadinessState(model.ReplayReadinessEvaluation)

	model.CurrentState = point16Val0Aggregate(
		model.DependencyState,
		model.HistoricalReplayContextState,
		model.OriginalDecisionBindingState,
		model.ReplayTaxonomyState,
		model.CurrentSubstitutionGuardState,
		model.ReplayReadinessState,
		model.NoOverclaimState,
	)
	model.BlockingReasons = point16Val0BlockingReasons(model)
	model.ReviewPrerequisites = point16Val0ReviewPrerequisites(model)
	return model
}

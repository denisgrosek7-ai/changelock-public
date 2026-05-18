package formal

import (
	"sort"
	"strings"
	"sync"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

const (
	Point16Val0StateActive         = "point16_val0_historical_replay_governance_active"
	Point16Val0StateBlocked        = "point16_val0_historical_replay_governance_blocked"
	Point16Val0StateReviewRequired = "point16_val0_historical_replay_governance_review_required"
	Point16Val0StateIncomplete     = "point16_val0_historical_replay_governance_incomplete"
)

const (
	point16Val0PointID                 = "point_16"
	point16Val0WaveID                  = "val_0"
	point16Val0Scope                   = "continuous_trust_evolution_and_historical_replay_governance_foundation"
	point16Val0ReplayDisclaimer        = "historical_replay_governance_foundation no_final_point16_closure point16_val0"
	point16Val0ContextID               = "historical_replay_context_point16_val0_001"
	point16Val0BindingID               = "original_decision_binding_point16_val0_001"
	point16Val0TaxonomyID              = "replay_taxonomy_point16_val0_001"
	point16Val0GuardID                 = "substitution_guard_point16_val0_001"
	point16Val0EvaluationID            = "readiness_evaluation_point16_val0_001"
	point16Val0OriginalDecisionID      = "decision_point16_val0_original_001"
	point16Val0OriginalDecisionHash    = "sha256:3333333333333333333333333333333333333333333333333333333333333333"
	point16Val0LineageRef              = "lineage_point16_val0_replay_001"
	point16Val0OriginalDecisionAt      = "2026-05-08T09:00:00Z"
	point16Val0OriginalEvaluatedAt     = "2026-05-08T09:05:00Z"
	point16Val0ReplayAt                = "2026-05-08T09:10:00Z"
	point16Val0OriginalPolicyID        = "policy_historical_replay_001"
	point16Val0OriginalPolicyHash      = "sha256:1111111111111111111111111111111111111111111111111111111111111111"
	point16Val0OriginalEngineID        = "engine_historical_replay_001"
	point16Val0OriginalEngineHash      = "sha256:2222222222222222222222222222222222222222222222222222222222222222"
	point16Val0OriginalArtifactScope   = "artifact_scope_historical_replay_001"
	point16Val0OriginalClaimScope      = "claim_scope_historical_replay_001"
	point16Val0OriginalGovernanceScope = "governance_scope_historical_replay_001"

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

var (
	point16Val0CanonicalPoint15ValEOnce sync.Once
	point16Val0CanonicalPoint15ValE     Point15ValEContinuousVerificationClosureFoundation
	point16Val0CanonicalDependency      Point16Val0DependencySnapshot
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
	Point15PassManifestScope         string                                             `json:"point15_pass_manifest_scope"`
	Point15PassManifestEvidenceID    string                                             `json:"point15_pass_manifest_evidence_identity"`
	Point15PassManifestEvidenceHash  string                                             `json:"point15_pass_manifest_evidence_hash"`
	Point15PassManifestPolicyVersion string                                             `json:"point15_pass_manifest_policy_version"`
	Point15PassManifestEngineVersion string                                             `json:"point15_pass_manifest_engine_version"`
	Point15PassManifestSchemaVersion string                                             `json:"point15_pass_manifest_schema_version"`
	Point15PassManifestGeneratedAt   string                                             `json:"point15_pass_manifest_generated_at"`
	Point15PassManifestTenantScope   string                                             `json:"point15_pass_manifest_tenant_scope"`
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
	OriginalDecisionID          string `json:"original_decision_id"`
	OriginalDecisionHash        string `json:"original_decision_hash"`
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
	OriginalPolicyVersion          string `json:"original_policy_version"`
	CurrentPolicyVersion           string `json:"current_policy_version"`
	OriginalPolicyHash             string `json:"original_policy_hash"`
	CurrentPolicyHash              string `json:"current_policy_hash"`
	OriginalEngineVersion          string `json:"original_engine_version"`
	CurrentEngineVersion           string `json:"current_engine_version"`
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

func point16Val0ExactAllowedValueValid(value string, allowed []string) bool {
	for _, candidate := range allowed {
		if value == candidate {
			return true
		}
	}
	return false
}

func point16Val0StateValid(value string) bool {
	return point16Val0ExactAllowedValueValid(value, point16Val0States())
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
	return point16Val0ExactAllowedValueValid(value, point16Val0ReplayStatuses())
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
	normalized := point16Val0NormalizeObservedText(text)
	if normalized == "" {
		return false
	}
	for _, phrase := range point16Val0ForbiddenWording() {
		normalizedPhrase := point16Val0NormalizeObservedText(phrase)
		if normalizedPhrase == "" {
			continue
		}
		if point16Val0ObservedTextContainsNormalizedPhrase(text, normalizedPhrase) {
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
	if len(values) > 1 && point16Val0ObservedTextContainsForbiddenWording(strings.Join(values, " ")) {
		return true
	}
	return false
}

func point16Val0CrossObservedDiagnosticForbiddenWording(observed, diagnostics []string) bool {
	if len(observed) == 0 || len(diagnostics) == 0 {
		return false
	}
	observedCorpusVariants := point16Val0NormalizedObservedTextVariants(strings.Join(observed, " "))
	diagnosticsCorpusVariants := point16Val0NormalizedObservedTextVariants(strings.Join(diagnostics, " "))
	combinedCorpusVariants := point16Val0NormalizedObservedTextVariants(strings.Join(append(append([]string{}, observed...), diagnostics...), " "))
	if len(observedCorpusVariants) == 0 || len(diagnosticsCorpusVariants) == 0 || len(combinedCorpusVariants) == 0 {
		return false
	}
	for _, phrase := range point16Val0ForbiddenWording() {
		normalizedPhrase := point16Val0NormalizeObservedText(phrase)
		if normalizedPhrase == "" {
			continue
		}
		if !point16Val0AnyNormalizedVariantContainsNormalizedPhrase(combinedCorpusVariants, normalizedPhrase) {
			continue
		}
		if point16Val0AnyNormalizedVariantContainsNormalizedPhrase(observedCorpusVariants, normalizedPhrase) {
			continue
		}
		if point16Val0AnyNormalizedVariantContainsNormalizedPhrase(diagnosticsCorpusVariants, normalizedPhrase) {
			continue
		}
		return true
	}
	return false
}

func point16Val0ObservedTextContainsNormalizedPhrase(text, normalizedPhrase string) bool {
	for _, normalizedText := range point16Val0NormalizedObservedTextVariants(text) {
		if point16Val0NormalizedTextContainsNormalizedPhrase(normalizedText, normalizedPhrase) {
			return true
		}
	}
	return false
}

func point16Val0AnyNormalizedVariantContainsNormalizedPhrase(normalizedTexts []string, normalizedPhrase string) bool {
	for _, normalizedText := range normalizedTexts {
		if point16Val0NormalizedTextContainsNormalizedPhrase(normalizedText, normalizedPhrase) {
			return true
		}
	}
	return false
}

func point16Val0NormalizedTextContainsNormalizedPhrase(normalizedText, normalizedPhrase string) bool {
	if normalizedText == "" || normalizedPhrase == "" {
		return false
	}
	if strings.Contains(normalizedText, normalizedPhrase) {
		return true
	}
	compactText := strings.ReplaceAll(normalizedText, " ", "")
	compactPhrase := strings.ReplaceAll(normalizedPhrase, " ", "")
	return compactPhrase != "" && strings.Contains(compactText, compactPhrase)
}

func point16Val0NormalizedObservedTextVariants(text string) []string {
	primary := point16Val0NormalizeObservedText(text)
	if primary == "" {
		return nil
	}
	variants := []string{primary}
	if strings.ContainsRune(text, 'ſ') {
		longSAsF := strings.Map(func(r rune) rune {
			if r == 'ſ' {
				return 'f'
			}
			return r
		}, text)
		alternate := point16Val0NormalizeObservedText(longSAsF)
		if alternate != "" && alternate != primary {
			variants = append(variants, alternate)
		}
	}
	return variants
}

func point16Val0NormalizeObservedText(text string) string {
	trimmed := strings.TrimSpace(strings.Map(func(r rune) rune {
		switch r {
		case 'ſ':
			return 's'
		default:
			return r
		}
	}, text))
	trimmed = norm.NFKC.String(trimmed)
	if trimmed == "" {
		return ""
	}

	var builder strings.Builder
	lastWasSpace := true
	for _, r := range norm.NFD.String(trimmed) {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		r = unicode.ToLower(point16Val0FoldConfusableRune(r))
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(r)
			lastWasSpace = false
			continue
		}
		if !lastWasSpace {
			builder.WriteByte(' ')
			lastWasSpace = true
		}
	}
	return strings.TrimSpace(builder.String())
}

func point16Val0FoldConfusableRune(r rune) rune {
	switch r {
	case 'а', 'А', 'α', 'Α':
		return 'a'
	case 'ɑ', 'ɐ', 'ɒ', '⍺':
		return 'a'
	case 'ᴀ':
		return 'a'
	case 'е', 'Е', 'ε', 'Ε':
		return 'e'
	case 'ᴇ':
		return 'e'
	case 'ɛ', 'Ɛ':
		return 'e'
	case 'і', 'І', 'ı', 'ι', 'Ι':
		return 'i'
	case 'о', 'О', 'ο', 'Ο', 'օ', 'Օ', 'ɔ':
		return 'o'
	case 'ᴏ':
		return 'o'
	case 'р', 'Р', 'ρ', 'Ρ':
		return 'p'
	case 'ᴘ':
		return 'p'
	case 'ɹ':
		return 'r'
	case 'ʀ':
		return 'r'
	case 'с', 'С':
		return 'c'
	case 'ᴄ':
		return 'c'
	case 'ƈ', 'ȼ':
		return 'c'
	case 'ʙ':
		return 'b'
	case 'ƒ':
		return 'f'
	case 'ꜰ':
		return 'f'
	case 'ᴅ':
		return 'd'
	case 'υ', 'Υ', 'ύ', 'ϋ':
		return 'u'
	case 'ᴜ':
		return 'u'
	case 'ᴠ':
		return 'v'
	case 'у', 'У':
		return 'y'
	case 'х', 'Х', 'χ', 'Χ':
		return 'x'
	case 'к', 'К', 'κ', 'Κ':
		return 'k'
	case 'м', 'М':
		return 'm'
	case 'ſ':
		return 's'
	case 'т', 'Т', 'τ', 'Τ':
		return 't'
	case 'в', 'В':
		return 'b'
	case 'н', 'Н':
		return 'h'
	case 'ɴ':
		return 'n'
	case 'ј', 'Ј':
		return 'j'
	case 'ԁ', 'Ԁ', 'δ', 'Δ':
		return 'd'
	case 'ɡ', 'ԍ', 'ց', 'ǥ':
		return 'g'
	case 'ɢ':
		return 'g'
	case 'һ', 'Η', 'հ':
		return 'h'
	case 'ʜ':
		return 'h'
	case 'ⅼ', 'ӏ', 'Ӏ', 'ǀ':
		return 'l'
	case 'ɩ', 'ɪ', 'ɭ', 'ɫ', 'ł', 'ƚ', 'ḷ':
		return 'l'
	case 'ʟ':
		return 'l'
	case 'ѕ', 'Ѕ':
		return 's'
	case 'ꜱ':
		return 's'
	case 'ᴛ':
		return 't'
	case 'Ь', 'ь', 'β', 'Β':
		return 'b'
	case 'ɗ', 'ḍ', 'đ', 'ð':
		return 'd'
	case '0':
		return 'o'
	case '1':
		return 'i'
	case '3':
		return 'e'
	case '4':
		return 'a'
	case '5':
		return 's'
	case '7':
		return 't'
	case '8':
		return 'a'
	case '9':
		return 'g'
	default:
		return r
	}
}

func point16Val0ExactMatch(left, right string) bool {
	return left != "" && right != "" && left == right
}

func point16Val0ExactRawStringSetMatch(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	seen := map[string]int{}
	for _, value := range left {
		seen[value]++
	}
	for _, value := range right {
		if seen[value] == 0 {
			return false
		}
		seen[value]--
	}
	for _, count := range seen {
		if count != 0 {
			return false
		}
	}
	return true
}

func point16Val0ExactVersionBindingValid(left, right string) bool {
	return left != "" &&
		right != "" &&
		strings.TrimSpace(left) == left &&
		strings.TrimSpace(right) == right &&
		point12Val0VersionIdentityValid(left) &&
		point12Val0VersionIdentityValid(right) &&
		left == right
}

func point16Val0CanonicalizedValidatedBindingValid(left, right string, validator func(string) bool) bool {
	left = strings.TrimSpace(left)
	right = strings.TrimSpace(right)
	return left != "" &&
		right != "" &&
		validator(left) &&
		validator(right) &&
		left == right
}

func point16Val0ExactValidatedBindingValid(left, right string, validator func(string) bool) bool {
	return left != "" &&
		right != "" &&
		strings.TrimSpace(left) == left &&
		strings.TrimSpace(right) == right &&
		validator(left) &&
		validator(right) &&
		left == right
}

func point16Val0CanonicalizedHistoricalEvidenceHashBindingValid(left, right string) bool {
	return point16Val0CanonicalizedValidatedBindingValid(left, right, point16Val0HistoricalEvidenceHashValid)
}

func point16Val0ExactHistoricalEvidenceHashBindingValid(left, right string) bool {
	return point16Val0ExactValidatedBindingValid(left, right, point16Val0HistoricalEvidenceHashValid)
}

func point16Val0ExactHashBindingValid(left, right string) bool {
	return point16Val0ExactValidatedBindingValid(left, right, point12Val0HashValid)
}

func point16Val0ExactTenantScopeBindingValid(left, right string) bool {
	return point16Val0ExactValidatedBindingValid(left, right, point11Val0ScopeValid)
}

func point16Val0CanonicalizedTenantScopeBindingValid(left, right string) bool {
	return point16Val0CanonicalizedValidatedBindingValid(left, right, point11Val0ScopeValid)
}

func point16Val0CanonicalizedVersionBindingValid(left, right string) bool {
	return point16Val0CanonicalizedValidatedBindingValid(left, right, point12Val0VersionIdentityValid)
}

func point16Val0ExactArtifactScopeBindingValid(left, right string) bool {
	return point16Val0ExactValidatedBindingValid(left, right, point16Val0ArtifactScopeValid)
}

func point16Val0ExactClaimScopeBindingValid(left, right string) bool {
	return point16Val0ExactValidatedBindingValid(left, right, point16Val0ClaimScopeValid)
}

func point16Val0ExactGovernanceScopeBindingValid(left, right string) bool {
	return point16Val0ExactValidatedBindingValid(left, right, point16Val0GovernanceScopeValid)
}

func point16Val0ExactTimestampBindingValid(left, right string) bool {
	return left != "" &&
		right != "" &&
		strings.TrimSpace(left) == left &&
		strings.TrimSpace(right) == right &&
		point14Val0ParsedTimeOk(left) &&
		point14Val0ParsedTimeOk(right) &&
		left == right
}

func point16Val0ExactHistoricalTimeSourceValid(value string) bool {
	return value == point14Val0TimeSourceServerUTC
}

func point16Val0ReplayProvenanceExactBound(context Point16Val0HistoricalReplayContext, binding Point16Val0OriginalDecisionBinding) bool {
	return point16Val0ExactMatch(context.ContextID, binding.HistoricalReplayContextRef) &&
		point16Val0ExactMatch(context.OriginalDecisionID, binding.OriginalDecisionID) &&
		point16Val0ExactMatch(context.OriginalDecisionHash, binding.OriginalDecisionHash) &&
		point16Val0ExactMatch(context.LineageRef, binding.LineageRef)
}

func point16Val0DependencyRefValid(value string) bool {
	return point14Val0RefValid(value, "point16_val0_", "historical_replay_", "original_decision_", "replay_", "substitution_", "readiness_")
}

func point16Val0CanonicalDependencySnapshot() Point16Val0DependencySnapshot {
	point16Val0CanonicalPoint15ValEOnce.Do(func() {
		point16Val0CanonicalPoint15ValE = ComputePoint15ValEFoundation(Point15ValEFoundationModel())
		point16Val0CanonicalDependency = point16Val0DependencySnapshotFromUpstream(point16Val0CanonicalPoint15ValE)
	})
	return point16Val0CanonicalDependency
}

func point16Val0ExactInternalIDValid(value, expected string) bool {
	return value == expected
}

func point16Val0ExactOriginalDecisionIDValid(value string) bool {
	return value == point16Val0OriginalDecisionID
}

func point16Val0ExactOriginalDecisionHashValid(value string) bool {
	return value == point16Val0OriginalDecisionHash
}

func point16Val0ExactLineageRefValid(value string) bool {
	return value == point16Val0LineageRef
}

func point16Val0ExactOriginalPolicyIDValid(value string) bool {
	return value == point16Val0OriginalPolicyID
}

func point16Val0ExactOriginalPolicyHashValid(value string) bool {
	return value == point16Val0OriginalPolicyHash
}

func point16Val0ExactOriginalEngineIDValid(value string) bool {
	return value == point16Val0OriginalEngineID
}

func point16Val0ExactOriginalEngineHashValid(value string) bool {
	return value == point16Val0OriginalEngineHash
}

func point16Val0ExactOriginalArtifactScopeValid(value string) bool {
	return value == point16Val0OriginalArtifactScope
}

func point16Val0ExactOriginalClaimScopeValid(value string) bool {
	return value == point16Val0OriginalClaimScope
}

func point16Val0ExactOriginalGovernanceScopeValid(value string) bool {
	return value == point16Val0OriginalGovernanceScope
}

func point16Val0ExactOriginalDecisionAtValid(value string) bool {
	return value == point16Val0OriginalDecisionAt
}

func point16Val0ExactOriginalEvaluatedAtValid(value string) bool {
	return value == point16Val0OriginalEvaluatedAt
}

func point16Val0ExactReplayAtValid(value string) bool {
	return value == point16Val0ReplayAt
}

func point16Val0EvidenceIdentityValid(value string) bool {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" || trimmed != value {
		return false
	}

	fields := strings.Fields(trimmed)
	if len(fields) != 6 {
		return false
	}
	if strings.Join(fields, " ") != value {
		return false
	}

	validators := []struct {
		prefix    string
		validator func(string) bool
	}{
		{prefix: "evidence_id=", validator: point15Val0EvidenceIDValid},
		{prefix: "evidence_hash=", validator: point16Val0HistoricalEvidenceHashValid},
		{prefix: "policy=", validator: point12Val0VersionIdentityValid},
		{prefix: "engine=", validator: point12Val0VersionIdentityValid},
		{prefix: "schema=", validator: point12Val0VersionIdentityValid},
		{prefix: "tenant=", validator: point11Val0ScopeValid},
	}
	for idx, validator := range validators {
		field := fields[idx]
		if !strings.HasPrefix(field, validator.prefix) {
			return false
		}
		valuePart := strings.TrimSpace(strings.TrimPrefix(field, validator.prefix))
		if valuePart == "" || !validator.validator(valuePart) {
			return false
		}
	}
	return true
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
	switch status {
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

func point16Val0AggregateFoundationStates(states ...string) string {
	hasIncomplete := false
	hasReview := false
	for _, state := range states {
		switch state {
		case Point16Val0StateBlocked:
			return Point16Val0StateBlocked
		case Point16Val0StateReviewRequired:
			hasReview = true
		case Point16Val0StateIncomplete:
			hasIncomplete = true
		case Point16Val0StateActive:
			continue
		default:
			return Point16Val0StateBlocked
		}
	}
	if hasReview {
		return Point16Val0StateReviewRequired
	}
	if hasIncomplete {
		return Point16Val0StateIncomplete
	}
	return Point16Val0StateActive
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
		Point15PassManifestScope:         valE.PassClosureManifest.Scope,
		Point15PassManifestEvidenceID:    valE.PassClosureManifest.EvidenceIdentity,
		Point15PassManifestEvidenceHash:  valE.PassClosureManifest.EvidenceHash,
		Point15PassManifestPolicyVersion: valE.PassClosureManifest.PolicyVersion,
		Point15PassManifestEngineVersion: valE.PassClosureManifest.EngineVersion,
		Point15PassManifestSchemaVersion: valE.PassClosureManifest.SchemaVersion,
		Point15PassManifestGeneratedAt:   valE.PassClosureManifest.GeneratedAt,
		Point15PassManifestTenantScope:   valE.PassClosureManifest.TenantScope,
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
	expected := point16Val0CanonicalDependencySnapshot()
	if !model.SnapshotFromComputedOutput ||
		!model.Point15ValEComputedFromUpstream ||
		!model.Point15ValEMerged ||
		!model.Point15ValECIGreen ||
		!model.Point15ValEReviewedOnMain ||
		!point15ValEStateValid(model.Point15ValECurrentState) ||
		!point15ValEStateValid(model.Point15ValEDependencyState) ||
		!point15ValEStateValid(model.Point15ValEClosureEvaluatorState) ||
		!point15ValEStateValid(model.Point15ValEPassClosureState) ||
		!point14Val0ParsedTimeOk(model.Point15PassManifestGeneratedAt) ||
		!point11Val0ScopeValid(model.InheritedTenantScope) ||
		strings.TrimSpace(model.Point15ValECurrentState) != model.Point15ValECurrentState ||
		strings.TrimSpace(model.Point15ValEDependencyState) != model.Point15ValEDependencyState ||
		strings.TrimSpace(model.Point15ValEClosureEvaluatorState) != model.Point15ValEClosureEvaluatorState ||
		strings.TrimSpace(model.Point15ValEPassClosureState) != model.Point15ValEPassClosureState ||
		strings.TrimSpace(model.Point15PassManifestScope) != model.Point15PassManifestScope ||
		strings.TrimSpace(model.Point15PassManifestGeneratedAt) != model.Point15PassManifestGeneratedAt {
		return Point16Val0StateBlocked
	}
	if !point16Val0ExactMatch(model.Point15ValECurrentState, model.Point15ValE.CurrentState) ||
		!point16Val0ExactMatch(model.Point15ValEDependencyState, model.Point15ValE.DependencyState) ||
		!point16Val0ExactMatch(model.Point15ValEClosureEvaluatorState, model.Point15ValE.ClosureEvaluatorState) ||
		!point16Val0ExactMatch(model.Point15ValEPassClosureState, model.Point15ValE.PassClosureManifestState) ||
		model.Point15ValEComputedFromUpstream != model.Point15ValE.Dependency.SnapshotFromComputedOutput ||
		model.Point15PassAllowed != model.Point15ValE.PassClosureManifest.Point15PassAllowed ||
		!point16Val0ExactMatch(model.Point15PassToken, model.Point15ValE.PassClosureManifest.Point15PassToken) ||
		!point16Val0ExactMatch(model.Point15PassManifestPointID, model.Point15ValE.PassClosureManifest.PointID) ||
		!point16Val0ExactMatch(model.Point15PassManifestWaveID, model.Point15ValE.PassClosureManifest.WaveID) ||
		!point16Val0ExactMatch(model.Point15PassManifestClosureToken, model.Point15ValE.PassClosureManifest.ClosureToken) ||
		!point16Val0ExactMatch(model.Point15PassManifestScope, model.Point15ValE.PassClosureManifest.Scope) ||
		!point16Val0ExactValidatedBindingValid(model.Point15PassManifestEvidenceID, model.Point15ValE.PassClosureManifest.EvidenceIdentity, point16Val0EvidenceIdentityValid) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(model.Point15PassManifestEvidenceHash, model.Point15ValE.PassClosureManifest.EvidenceHash) ||
		!point16Val0ExactVersionBindingValid(model.Point15PassManifestPolicyVersion, model.Point15ValE.PassClosureManifest.PolicyVersion) ||
		!point16Val0ExactVersionBindingValid(model.Point15PassManifestEngineVersion, model.Point15ValE.PassClosureManifest.EngineVersion) ||
		!point16Val0ExactVersionBindingValid(model.Point15PassManifestSchemaVersion, model.Point15ValE.PassClosureManifest.SchemaVersion) ||
		!point16Val0ExactMatch(model.Point15PassManifestGeneratedAt, model.Point15ValE.PassClosureManifest.GeneratedAt) ||
		!point16Val0ExactTenantScopeBindingValid(model.Point15PassManifestTenantScope, model.Point15ValE.PassClosureManifest.TenantScope) ||
		!point16Val0ExactTenantScopeBindingValid(model.InheritedTenantScope, model.Point15ValE.Dependency.InheritedTenantScope) {
		return Point16Val0StateBlocked
	}
	if !point16Val0ExactMatch(model.Point15ValECurrentState, expected.Point15ValECurrentState) ||
		!point16Val0ExactMatch(model.Point15ValEDependencyState, expected.Point15ValEDependencyState) ||
		!point16Val0ExactMatch(model.Point15ValEClosureEvaluatorState, expected.Point15ValEClosureEvaluatorState) ||
		!point16Val0ExactMatch(model.Point15ValEPassClosureState, expected.Point15ValEPassClosureState) ||
		model.Point15PassAllowed != expected.Point15PassAllowed ||
		!point16Val0ExactMatch(model.Point15PassToken, expected.Point15PassToken) ||
		!point16Val0ExactMatch(model.Point15PassManifestPointID, expected.Point15PassManifestPointID) ||
		!point16Val0ExactMatch(model.Point15PassManifestWaveID, expected.Point15PassManifestWaveID) ||
		!point16Val0ExactMatch(model.Point15PassManifestClosureToken, expected.Point15PassManifestClosureToken) ||
		!point16Val0ExactMatch(model.Point15PassManifestScope, expected.Point15PassManifestScope) ||
		!point16Val0ExactMatch(model.Point15ValE.PassClosureManifest.Scope, expected.Point15ValE.PassClosureManifest.Scope) ||
		!point16Val0ExactValidatedBindingValid(model.Point15PassManifestEvidenceID, expected.Point15PassManifestEvidenceID, point16Val0EvidenceIdentityValid) ||
		!point16Val0ExactValidatedBindingValid(model.Point15ValE.PassClosureManifest.EvidenceIdentity, expected.Point15ValE.PassClosureManifest.EvidenceIdentity, point16Val0EvidenceIdentityValid) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(model.Point15PassManifestEvidenceHash, expected.Point15PassManifestEvidenceHash) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(model.Point15ValE.PassClosureManifest.EvidenceHash, expected.Point15ValE.PassClosureManifest.EvidenceHash) ||
		!point16Val0ExactVersionBindingValid(model.Point15PassManifestPolicyVersion, expected.Point15PassManifestPolicyVersion) ||
		!point16Val0ExactVersionBindingValid(model.Point15ValE.PassClosureManifest.PolicyVersion, expected.Point15ValE.PassClosureManifest.PolicyVersion) ||
		!point16Val0ExactVersionBindingValid(model.Point15PassManifestEngineVersion, expected.Point15PassManifestEngineVersion) ||
		!point16Val0ExactVersionBindingValid(model.Point15ValE.PassClosureManifest.EngineVersion, expected.Point15ValE.PassClosureManifest.EngineVersion) ||
		!point16Val0ExactVersionBindingValid(model.Point15PassManifestSchemaVersion, expected.Point15PassManifestSchemaVersion) ||
		!point16Val0ExactVersionBindingValid(model.Point15ValE.PassClosureManifest.SchemaVersion, expected.Point15ValE.PassClosureManifest.SchemaVersion) ||
		!point16Val0ExactMatch(model.Point15PassManifestGeneratedAt, expected.Point15PassManifestGeneratedAt) ||
		!point16Val0ExactMatch(model.Point15ValE.PassClosureManifest.GeneratedAt, expected.Point15ValE.PassClosureManifest.GeneratedAt) ||
		!point16Val0ExactTenantScopeBindingValid(model.Point15PassManifestTenantScope, expected.Point15PassManifestTenantScope) ||
		!point16Val0ExactTenantScopeBindingValid(model.Point15ValE.PassClosureManifest.TenantScope, expected.Point15ValE.PassClosureManifest.TenantScope) ||
		!point16Val0ExactTenantScopeBindingValid(model.InheritedTenantScope, expected.InheritedTenantScope) ||
		!point16Val0ExactTenantScopeBindingValid(model.Point15ValE.Dependency.InheritedTenantScope, expected.Point15ValE.Dependency.InheritedTenantScope) {
		return Point16Val0StateBlocked
	}
	if model.Point15ValECurrentState != Point15ValEStatePassConfirmed ||
		model.Point15ValEDependencyState != Point15ValEStatePassConfirmed ||
		model.Point15ValEClosureEvaluatorState != Point15ValEStatePassConfirmed ||
		model.Point15ValEPassClosureState != Point15ValEStatePassConfirmed ||
		!model.Point15PassAllowed ||
		model.Point15PassToken != point15Val0BlockedPassToken ||
		model.Point15PassManifestPointID != point15Val0PointID ||
		model.Point15PassManifestWaveID != point15ValEWaveID ||
		model.Point15PassManifestClosureToken != point15Val0BlockedPassToken ||
		model.Point15PassManifestScope != point15ValEScope ||
		!point16Val0EvidenceIdentityValid(model.Point15PassManifestEvidenceID) ||
		!point16Val0HistoricalEvidenceHashValid(model.Point15PassManifestEvidenceHash) ||
		!point12Val0VersionIdentityValid(model.Point15PassManifestPolicyVersion) ||
		!point12Val0VersionIdentityValid(model.Point15PassManifestEngineVersion) ||
		!point12Val0VersionIdentityValid(model.Point15PassManifestSchemaVersion) ||
		!point14Val0ParsedTimeOk(model.Point15PassManifestGeneratedAt) ||
		!point11Val0ScopeValid(model.Point15PassManifestTenantScope) ||
		strings.TrimSpace(model.InheritedTenantScope) != model.InheritedTenantScope {
		return Point16Val0StateBlocked
	}
	return Point16Val0StateActive
}

func point16Val0HistoricalReplayContextModel(dependency Point16Val0DependencySnapshot) Point16Val0HistoricalReplayContext {
	return Point16Val0HistoricalReplayContext{
		ContextID:                   point16Val0ContextID,
		OriginalEvidenceID:          dependency.Point15PassManifestEvidenceID,
		OriginalEvidenceHash:        dependency.Point15PassManifestEvidenceHash,
		OriginalPolicyID:            point16Val0OriginalPolicyID,
		OriginalPolicyVersion:       dependency.Point15PassManifestPolicyVersion,
		OriginalPolicyHash:          point16Val0OriginalPolicyHash,
		OriginalEngineID:            point16Val0OriginalEngineID,
		OriginalEngineVersion:       dependency.Point15PassManifestEngineVersion,
		OriginalEngineHash:          point16Val0OriginalEngineHash,
		OriginalTenantScope:         dependency.InheritedTenantScope,
		OriginalArtifactScope:       point16Val0OriginalArtifactScope,
		OriginalClaimScope:          point16Val0OriginalClaimScope,
		OriginalGovernanceScope:     point16Val0OriginalGovernanceScope,
		OriginalDecisionID:          point16Val0OriginalDecisionID,
		OriginalDecisionHash:        point16Val0OriginalDecisionHash,
		OriginalDecisionAt:          point16Val0OriginalDecisionAt,
		OriginalDecisionTimeSource:  point14Val0TimeSourceServerUTC,
		OriginalEvaluatedAt:         point16Val0OriginalEvaluatedAt,
		OriginalEvaluatedTimeSource: point14Val0TimeSourceServerUTC,
		ReplayAt:                    point16Val0ReplayAt,
		ReplayTimeSource:            point14Val0TimeSourceServerUTC,
		LineageRef:                  point16Val0LineageRef,
	}
}

func EvaluatePoint16Val0HistoricalReplayContextState(model Point16Val0HistoricalReplayContext) string {
	expectedDependency := point16Val0CanonicalDependencySnapshot()
	if !point16Val0ContextRefValid(model.ContextID) || !point16Val0ExactInternalIDValid(model.ContextID, point16Val0ContextID) {
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
		model.OriginalDecisionID,
		model.OriginalDecisionHash,
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
		!point16Val0ExactValidatedBindingValid(model.OriginalEvidenceID, expectedDependency.Point15PassManifestEvidenceID, point16Val0EvidenceIdentityValid) ||
		!point16Val0HistoricalEvidenceHashValid(model.OriginalEvidenceHash) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(model.OriginalEvidenceHash, expectedDependency.Point15PassManifestEvidenceHash) ||
		!point16Val0PolicyIDValid(model.OriginalPolicyID) ||
		!point16Val0ExactOriginalPolicyIDValid(model.OriginalPolicyID) ||
		!point12Val0VersionIdentityValid(model.OriginalPolicyVersion) ||
		!point16Val0ExactVersionBindingValid(model.OriginalPolicyVersion, expectedDependency.Point15PassManifestPolicyVersion) ||
		!point12Val0HashValid(model.OriginalPolicyHash) ||
		!point16Val0ExactOriginalPolicyHashValid(model.OriginalPolicyHash) ||
		!point16Val0EngineIDValid(model.OriginalEngineID) ||
		!point16Val0ExactOriginalEngineIDValid(model.OriginalEngineID) ||
		!point12Val0VersionIdentityValid(model.OriginalEngineVersion) ||
		!point16Val0ExactVersionBindingValid(model.OriginalEngineVersion, expectedDependency.Point15PassManifestEngineVersion) ||
		!point12Val0HashValid(model.OriginalEngineHash) ||
		!point16Val0ExactOriginalEngineHashValid(model.OriginalEngineHash) ||
		!point11Val0ScopeValid(model.OriginalTenantScope) ||
		!point16Val0ExactTenantScopeBindingValid(model.OriginalTenantScope, expectedDependency.InheritedTenantScope) ||
		!point16Val0ArtifactScopeValid(model.OriginalArtifactScope) ||
		!point16Val0ExactOriginalArtifactScopeValid(model.OriginalArtifactScope) ||
		!point16Val0ClaimScopeValid(model.OriginalClaimScope) ||
		!point16Val0ExactOriginalClaimScopeValid(model.OriginalClaimScope) ||
		!point16Val0GovernanceScopeValid(model.OriginalGovernanceScope) ||
		!point16Val0ExactOriginalGovernanceScopeValid(model.OriginalGovernanceScope) ||
		!point16Val0DecisionIDValid(model.OriginalDecisionID) ||
		!point16Val0ExactOriginalDecisionIDValid(model.OriginalDecisionID) ||
		!point12Val0HashValid(model.OriginalDecisionHash) ||
		!point16Val0ExactOriginalDecisionHashValid(model.OriginalDecisionHash) ||
		!point14Val0ParsedTimeOk(model.OriginalDecisionAt) ||
		!point16Val0ExactOriginalDecisionAtValid(model.OriginalDecisionAt) ||
		!point14Val0ParsedTimeOk(model.OriginalEvaluatedAt) ||
		!point16Val0ExactOriginalEvaluatedAtValid(model.OriginalEvaluatedAt) ||
		!point14Val0ParsedTimeOk(model.ReplayAt) ||
		!point16Val0ExactReplayAtValid(model.ReplayAt) ||
		!point16Val0ExactHistoricalTimeSourceValid(model.OriginalDecisionTimeSource) ||
		!point16Val0ExactHistoricalTimeSourceValid(model.OriginalEvaluatedTimeSource) ||
		!point16Val0ExactHistoricalTimeSourceValid(model.ReplayTimeSource) ||
		!point16Val0LineageRefValid(model.LineageRef) ||
		!point16Val0ExactLineageRefValid(model.LineageRef) {
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
		BindingID:                    point16Val0BindingID,
		HistoricalReplayContextRef:   context.ContextID,
		OriginalDecisionID:           context.OriginalDecisionID,
		OriginalDecisionHash:         context.OriginalDecisionHash,
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
	expectedDependency := point16Val0CanonicalDependencySnapshot()
	if !point16Val0DependencyRefValid(model.BindingID) ||
		!point16Val0ExactInternalIDValid(model.BindingID, point16Val0BindingID) ||
		!point16Val0ContextRefValid(model.HistoricalReplayContextRef) ||
		!point16Val0ExactInternalIDValid(model.HistoricalReplayContextRef, point16Val0ContextID) ||
		!point16Val0DecisionIDValid(model.OriginalDecisionID) ||
		!point16Val0ExactOriginalDecisionIDValid(model.OriginalDecisionID) ||
		!point12Val0HashValid(model.OriginalDecisionHash) ||
		!point16Val0ExactOriginalDecisionHashValid(model.OriginalDecisionHash) {
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
		!point16Val0ExactValidatedBindingValid(model.OriginalEvidenceID, expectedDependency.Point15PassManifestEvidenceID, point16Val0EvidenceIdentityValid) ||
		!point16Val0HistoricalEvidenceHashValid(model.OriginalEvidenceHash) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(model.OriginalEvidenceHash, expectedDependency.Point15PassManifestEvidenceHash) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(model.OriginalEvidenceHash, model.CurrentEvidenceHash) ||
		!point16Val0PolicyIDValid(model.OriginalPolicyID) ||
		!point16Val0ExactOriginalPolicyIDValid(model.OriginalPolicyID) ||
		!point16Val0ExactVersionBindingValid(model.OriginalPolicyVersion, model.CurrentPolicyVersion) ||
		!point16Val0ExactVersionBindingValid(model.OriginalPolicyVersion, expectedDependency.Point15PassManifestPolicyVersion) ||
		!point16Val0ExactHashBindingValid(model.OriginalPolicyHash, model.CurrentPolicyHash) ||
		!point12Val0HashValid(model.OriginalPolicyHash) ||
		!point16Val0ExactOriginalPolicyHashValid(model.OriginalPolicyHash) ||
		!point16Val0EngineIDValid(model.OriginalEngineID) ||
		!point16Val0ExactOriginalEngineIDValid(model.OriginalEngineID) ||
		!point16Val0ExactVersionBindingValid(model.OriginalEngineVersion, model.CurrentEngineVersion) ||
		!point16Val0ExactVersionBindingValid(model.OriginalEngineVersion, expectedDependency.Point15PassManifestEngineVersion) ||
		!point16Val0ExactHashBindingValid(model.OriginalEngineHash, model.CurrentEngineHash) ||
		!point12Val0HashValid(model.OriginalEngineHash) ||
		!point16Val0ExactOriginalEngineHashValid(model.OriginalEngineHash) ||
		!point11Val0ScopeValid(model.OriginalTenantScope) ||
		!point16Val0ExactTenantScopeBindingValid(model.OriginalTenantScope, expectedDependency.InheritedTenantScope) ||
		!point16Val0ExactTenantScopeBindingValid(model.OriginalTenantScope, model.CurrentTenantScope) ||
		!point16Val0ArtifactScopeValid(model.OriginalArtifactScope) ||
		!point16Val0ExactArtifactScopeBindingValid(model.OriginalArtifactScope, model.CurrentArtifactScope) ||
		!point16Val0ExactOriginalArtifactScopeValid(model.OriginalArtifactScope) ||
		!point16Val0ClaimScopeValid(model.OriginalClaimScope) ||
		!point16Val0ExactClaimScopeBindingValid(model.OriginalClaimScope, model.CurrentClaimScope) ||
		!point16Val0ExactOriginalClaimScopeValid(model.OriginalClaimScope) ||
		!point16Val0GovernanceScopeValid(model.OriginalGovernanceScope) ||
		!point16Val0ExactGovernanceScopeBindingValid(model.OriginalGovernanceScope, model.CurrentGovernanceScope) ||
		!point16Val0ExactOriginalGovernanceScopeValid(model.OriginalGovernanceScope) ||
		!point14Val0ParsedTimeOk(model.OriginalDecisionAt) ||
		!point16Val0ExactOriginalDecisionAtValid(model.OriginalDecisionAt) ||
		!point14Val0ParsedTimeOk(model.OriginalEvaluatedAt) ||
		!point16Val0ExactOriginalEvaluatedAtValid(model.OriginalEvaluatedAt) ||
		!point16Val0LineageRefValid(model.LineageRef) ||
		!point16Val0ExactLineageRefValid(model.LineageRef) {
		return Point16Val0StateBlocked
	}
	if !model.CurrentContextComparisonOnly {
		return Point16Val0StateBlocked
	}
	return Point16Val0StateActive
}

func point16Val0ReplayTaxonomyModel(context Point16Val0HistoricalReplayContext, binding Point16Val0OriginalDecisionBinding) Point16Val0ReplayTaxonomy {
	return Point16Val0ReplayTaxonomy{
		TaxonomyID:              point16Val0TaxonomyID,
		ReplayStatus:            point16Val0ReplayBound,
		AllowedStatuses:         point16Val0ReplayStatuses(),
		OriginalContextComplete: true,
		OriginalEvidencePresent: true,
		OriginalPolicyPresent:   true,
		OriginalEnginePresent:   true,
		EvidenceHashMatches:     point16Val0ExactHistoricalEvidenceHashBindingValid(context.OriginalEvidenceHash, binding.CurrentEvidenceHash),
		PolicyHashMatches:       point16Val0ExactHashBindingValid(context.OriginalPolicyHash, binding.CurrentPolicyHash),
		EngineHashMatches:       point16Val0ExactHashBindingValid(context.OriginalEngineHash, binding.CurrentEngineHash),
		TenantScopeMatches:      point16Val0ExactTenantScopeBindingValid(context.OriginalTenantScope, binding.CurrentTenantScope),
		ArtifactScopeMatches:    point16Val0ExactArtifactScopeBindingValid(context.OriginalArtifactScope, binding.CurrentArtifactScope),
		ClaimScopeMatches:       point16Val0ExactClaimScopeBindingValid(context.OriginalClaimScope, binding.CurrentClaimScope),
		GovernanceScopeMatches:  point16Val0ExactGovernanceScopeBindingValid(context.OriginalGovernanceScope, binding.CurrentGovernanceScope),
		TimestampConsistent:     strings.TrimSpace(context.OriginalDecisionAt) == strings.TrimSpace(binding.OriginalDecisionAt) && strings.TrimSpace(context.OriginalEvaluatedAt) == strings.TrimSpace(binding.OriginalEvaluatedAt),
		TimestampSafe:           true,
		LineagePresent:          point16Val0LineageRefValid(context.LineageRef) && point16Val0LineageRefValid(binding.LineageRef) && point16Val0ExactMatch(context.LineageRef, binding.LineageRef),
		ReplaySupported:         true,
	}
}

func EvaluatePoint16Val0ReplayTaxonomyState(model Point16Val0ReplayTaxonomy) string {
	if !point16Val0DependencyRefValid(model.TaxonomyID) ||
		!point16Val0ExactInternalIDValid(model.TaxonomyID, point16Val0TaxonomyID) ||
		!point16Val0ReplayStatusValid(model.ReplayStatus) ||
		!point16Val0ExactRawStringSetMatch(model.AllowedStatuses, point16Val0ReplayStatuses()) {
		return Point16Val0StateBlocked
	}
	expectedStatus := point16Val0ExpectedReplayStatus(model)
	if model.ReplayStatus != expectedStatus {
		return Point16Val0StateBlocked
	}
	return point16Val0ReplayStatusState(expectedStatus)
}

func point16Val0CurrentSubstitutionGuardModel(context Point16Val0HistoricalReplayContext, binding Point16Val0OriginalDecisionBinding) Point16Val0CurrentSubstitutionGuard {
	return Point16Val0CurrentSubstitutionGuard{
		GuardID:                 point16Val0GuardID,
		OriginalPolicyVersion:   context.OriginalPolicyVersion,
		CurrentPolicyVersion:    binding.CurrentPolicyVersion,
		OriginalPolicyHash:      context.OriginalPolicyHash,
		CurrentPolicyHash:       binding.CurrentPolicyHash,
		OriginalEngineVersion:   context.OriginalEngineVersion,
		CurrentEngineVersion:    binding.CurrentEngineVersion,
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
	expectedDependency := point16Val0CanonicalDependencySnapshot()
	if !point16Val0DependencyRefValid(model.GuardID) ||
		!point16Val0ExactInternalIDValid(model.GuardID, point16Val0GuardID) ||
		!point16Val0ExactVersionBindingValid(model.OriginalPolicyVersion, expectedDependency.Point15PassManifestPolicyVersion) ||
		!point16Val0ExactVersionBindingValid(model.OriginalPolicyVersion, model.CurrentPolicyVersion) ||
		!point12Val0HashValid(model.OriginalPolicyHash) ||
		!point16Val0ExactOriginalPolicyHashValid(model.OriginalPolicyHash) ||
		!point12Val0HashValid(model.CurrentPolicyHash) ||
		!point16Val0ExactOriginalPolicyHashValid(model.CurrentPolicyHash) ||
		!point16Val0ExactVersionBindingValid(model.OriginalEngineVersion, expectedDependency.Point15PassManifestEngineVersion) ||
		!point16Val0ExactVersionBindingValid(model.OriginalEngineVersion, model.CurrentEngineVersion) ||
		!point12Val0HashValid(model.OriginalEngineHash) ||
		!point16Val0ExactOriginalEngineHashValid(model.OriginalEngineHash) ||
		!point12Val0HashValid(model.CurrentEngineHash) ||
		!point16Val0ExactOriginalEngineHashValid(model.CurrentEngineHash) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(model.OriginalEvidenceHash, expectedDependency.Point15PassManifestEvidenceHash) ||
		!point16Val0HistoricalEvidenceHashValid(model.OriginalEvidenceHash) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(model.OriginalEvidenceHash, model.CurrentEvidenceHash) ||
		!point14Val0ParsedTimeOk(model.OriginalDecisionAt) ||
		!point16Val0ExactOriginalDecisionAtValid(model.OriginalDecisionAt) ||
		!point14Val0ParsedTimeOk(model.ReplayAt) ||
		!point16Val0ExactReplayAtValid(model.ReplayAt) ||
		!point16Val0ExactTenantScopeBindingValid(model.OriginalTenantScope, expectedDependency.InheritedTenantScope) ||
		!point11Val0ScopeValid(model.OriginalTenantScope) ||
		!point16Val0ExactTenantScopeBindingValid(model.OriginalTenantScope, model.CurrentTenantScope) ||
		!point16Val0ClaimScopeValid(model.OriginalClaimScope) ||
		!point16Val0ExactOriginalClaimScopeValid(model.OriginalClaimScope) ||
		!point16Val0ClaimScopeValid(model.CurrentClaimScope) ||
		!point16Val0ExactOriginalClaimScopeValid(model.CurrentClaimScope) ||
		!point16Val0GovernanceScopeValid(model.OriginalGovernanceScope) ||
		!point16Val0ExactOriginalGovernanceScopeValid(model.OriginalGovernanceScope) ||
		!point16Val0GovernanceScopeValid(model.CurrentGovernanceScope) ||
		!point16Val0ExactOriginalGovernanceScopeValid(model.CurrentGovernanceScope) {
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
		EvaluationID:                  point16Val0EvaluationID,
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
		!point16Val0ExactInternalIDValid(model.EvaluationID, point16Val0EvaluationID) ||
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
	if model.ReplayDisclaimer != point16Val0ReplayDisclaimer ||
		!point16Val0ExactRawStringSetMatch(model.AllowedSafeWording, point16Val0SafeWording()) ||
		!point16Val0ExactRawStringSetMatch(model.BlockedWording, point16Val0ForbiddenWording()) {
		return Point16Val0StateBlocked
	}
	if point16Val0CrossObservedDiagnosticForbiddenWording(model.ObservedTexts, model.InternalDiagnosticTexts) {
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
		if state == Point16Val0StateBlocked {
			return Point16Val0StateBlocked
		}
	}
	for _, state := range states {
		if state == Point16Val0StateReviewRequired {
			return Point16Val0StateReviewRequired
		}
	}
	for _, state := range states {
		if state == Point16Val0StateIncomplete {
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
		if state == Point16Val0StateBlocked {
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
		if state == Point16Val0StateReviewRequired || state == Point16Val0StateIncomplete {
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

	expectedEvidenceHashMatches := point16Val0ExactHistoricalEvidenceHashBindingValid(model.HistoricalReplayContext.OriginalEvidenceHash, model.OriginalDecisionBinding.CurrentEvidenceHash) &&
		point16Val0ExactHistoricalEvidenceHashBindingValid(model.HistoricalReplayContext.OriginalEvidenceHash, model.CurrentSubstitutionGuard.CurrentEvidenceHash)
	expectedPolicyVersionMatches := point16Val0ExactVersionBindingValid(model.HistoricalReplayContext.OriginalPolicyVersion, model.OriginalDecisionBinding.CurrentPolicyVersion) &&
		point16Val0ExactVersionBindingValid(model.HistoricalReplayContext.OriginalPolicyVersion, model.CurrentSubstitutionGuard.CurrentPolicyVersion)
	expectedPolicyHashMatches := point16Val0ExactHashBindingValid(model.HistoricalReplayContext.OriginalPolicyHash, model.OriginalDecisionBinding.CurrentPolicyHash) &&
		point16Val0ExactHashBindingValid(model.HistoricalReplayContext.OriginalPolicyHash, model.CurrentSubstitutionGuard.CurrentPolicyHash)
	expectedEngineVersionMatches := point16Val0ExactVersionBindingValid(model.HistoricalReplayContext.OriginalEngineVersion, model.OriginalDecisionBinding.CurrentEngineVersion) &&
		point16Val0ExactVersionBindingValid(model.HistoricalReplayContext.OriginalEngineVersion, model.CurrentSubstitutionGuard.CurrentEngineVersion)
	expectedEngineHashMatches := point16Val0ExactHashBindingValid(model.HistoricalReplayContext.OriginalEngineHash, model.OriginalDecisionBinding.CurrentEngineHash) &&
		point16Val0ExactHashBindingValid(model.HistoricalReplayContext.OriginalEngineHash, model.CurrentSubstitutionGuard.CurrentEngineHash)
	expectedTenantScopeMatches := point16Val0ExactTenantScopeBindingValid(model.HistoricalReplayContext.OriginalTenantScope, model.OriginalDecisionBinding.CurrentTenantScope) &&
		point16Val0ExactTenantScopeBindingValid(model.HistoricalReplayContext.OriginalTenantScope, model.CurrentSubstitutionGuard.CurrentTenantScope)
	expectedArtifactScopeMatches := point16Val0ExactArtifactScopeBindingValid(model.HistoricalReplayContext.OriginalArtifactScope, model.OriginalDecisionBinding.CurrentArtifactScope)
	expectedClaimScopeMatches := point16Val0ExactClaimScopeBindingValid(model.HistoricalReplayContext.OriginalClaimScope, model.OriginalDecisionBinding.CurrentClaimScope) &&
		point16Val0ExactClaimScopeBindingValid(model.HistoricalReplayContext.OriginalClaimScope, model.CurrentSubstitutionGuard.CurrentClaimScope)
	expectedGovernanceScopeMatches := point16Val0ExactGovernanceScopeBindingValid(model.HistoricalReplayContext.OriginalGovernanceScope, model.OriginalDecisionBinding.CurrentGovernanceScope) &&
		point16Val0ExactGovernanceScopeBindingValid(model.HistoricalReplayContext.OriginalGovernanceScope, model.CurrentSubstitutionGuard.CurrentGovernanceScope)
	expectedReplayProvenanceExactBound := point16Val0ReplayProvenanceExactBound(model.HistoricalReplayContext, model.OriginalDecisionBinding)
	expectedHistoricalTimeSourcesSafe := point16Val0ExactHistoricalTimeSourceValid(model.HistoricalReplayContext.OriginalDecisionTimeSource) &&
		point16Val0ExactHistoricalTimeSourceValid(model.HistoricalReplayContext.OriginalEvaluatedTimeSource) &&
		point16Val0ExactHistoricalTimeSourceValid(model.HistoricalReplayContext.ReplayTimeSource)
	expectedTimestampConsistent := point16Val0ExactTimestampBindingValid(model.HistoricalReplayContext.OriginalDecisionAt, model.OriginalDecisionBinding.OriginalDecisionAt) &&
		point16Val0ExactTimestampBindingValid(model.HistoricalReplayContext.OriginalDecisionAt, model.CurrentSubstitutionGuard.OriginalDecisionAt) &&
		point16Val0ExactTimestampBindingValid(model.HistoricalReplayContext.OriginalEvaluatedAt, model.OriginalDecisionBinding.OriginalEvaluatedAt) &&
		point16Val0ExactTimestampBindingValid(model.HistoricalReplayContext.ReplayAt, model.CurrentSubstitutionGuard.ReplayAt) &&
		expectedHistoricalTimeSourcesSafe
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
		strings.TrimSpace(model.HistoricalReplayContext.OriginalDecisionID) != "" &&
		strings.TrimSpace(model.HistoricalReplayContext.OriginalDecisionHash) != "" &&
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
		model.ReplayTaxonomy.LineagePresent != (point16Val0LineageRefValid(model.HistoricalReplayContext.LineageRef) && point16Val0LineageRefValid(model.OriginalDecisionBinding.LineageRef) && point16Val0ExactMatch(model.HistoricalReplayContext.LineageRef, model.OriginalDecisionBinding.LineageRef)) {
		model.ReplayTaxonomyState = Point16Val0StateBlocked
	} else {
		model.ReplayTaxonomyState = EvaluatePoint16Val0ReplayTaxonomyState(model.ReplayTaxonomy)
	}

	model.CurrentSubstitutionGuardState = EvaluatePoint16Val0CurrentSubstitutionGuardState(model.CurrentSubstitutionGuard)
	model.NoOverclaimState = EvaluatePoint16Val0NoOverclaimBaselineState(model.NoOverclaimBaseline)

	expectedTenant := model.Dependency.InheritedTenantScope
	expectedEvidenceID := model.Dependency.Point15PassManifestEvidenceID
	expectedEvidenceHash := model.Dependency.Point15PassManifestEvidenceHash
	expectedPolicyVersion := model.Dependency.Point15PassManifestPolicyVersion
	expectedEngineVersion := model.Dependency.Point15PassManifestEngineVersion

	if !point16Val0ExactTenantScopeBindingValid(expectedTenant, model.HistoricalReplayContext.OriginalTenantScope) ||
		!point16Val0ExactTenantScopeBindingValid(expectedTenant, model.OriginalDecisionBinding.OriginalTenantScope) ||
		!point16Val0ExactTenantScopeBindingValid(expectedTenant, model.CurrentSubstitutionGuard.OriginalTenantScope) {
		model.HistoricalReplayContextState = Point16Val0StateBlocked
		model.OriginalDecisionBindingState = Point16Val0StateBlocked
		model.CurrentSubstitutionGuardState = Point16Val0StateBlocked
	}
	if !point16Val0ExactValidatedBindingValid(expectedEvidenceID, model.HistoricalReplayContext.OriginalEvidenceID, point16Val0EvidenceIdentityValid) ||
		!point16Val0ExactValidatedBindingValid(expectedEvidenceID, model.OriginalDecisionBinding.OriginalEvidenceID, point16Val0EvidenceIdentityValid) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(expectedEvidenceHash, model.HistoricalReplayContext.OriginalEvidenceHash) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(expectedEvidenceHash, model.OriginalDecisionBinding.OriginalEvidenceHash) ||
		!point16Val0ExactHistoricalEvidenceHashBindingValid(expectedEvidenceHash, model.CurrentSubstitutionGuard.OriginalEvidenceHash) ||
		!point16Val0ExactVersionBindingValid(expectedPolicyVersion, model.HistoricalReplayContext.OriginalPolicyVersion) ||
		!point16Val0ExactVersionBindingValid(expectedPolicyVersion, model.OriginalDecisionBinding.OriginalPolicyVersion) ||
		!point16Val0ExactVersionBindingValid(expectedPolicyVersion, model.OriginalDecisionBinding.CurrentPolicyVersion) ||
		!point16Val0ExactVersionBindingValid(expectedPolicyVersion, model.CurrentSubstitutionGuard.OriginalPolicyVersion) ||
		!point16Val0ExactVersionBindingValid(expectedPolicyVersion, model.CurrentSubstitutionGuard.CurrentPolicyVersion) ||
		!point16Val0ExactVersionBindingValid(expectedEngineVersion, model.HistoricalReplayContext.OriginalEngineVersion) ||
		!point16Val0ExactVersionBindingValid(expectedEngineVersion, model.OriginalDecisionBinding.OriginalEngineVersion) ||
		!point16Val0ExactVersionBindingValid(expectedEngineVersion, model.OriginalDecisionBinding.CurrentEngineVersion) ||
		!point16Val0ExactVersionBindingValid(expectedEngineVersion, model.CurrentSubstitutionGuard.OriginalEngineVersion) ||
		!point16Val0ExactVersionBindingValid(expectedEngineVersion, model.CurrentSubstitutionGuard.CurrentEngineVersion) {
		model.HistoricalReplayContextState = Point16Val0StateBlocked
		model.OriginalDecisionBindingState = Point16Val0StateBlocked
		model.CurrentSubstitutionGuardState = Point16Val0StateBlocked
		model.ReplayTaxonomyState = Point16Val0StateBlocked
	}
	if !expectedPolicyVersionMatches || !expectedEngineVersionMatches {
		model.OriginalDecisionBindingState = Point16Val0StateBlocked
		model.CurrentSubstitutionGuardState = Point16Val0StateBlocked
		model.ReplayTaxonomyState = Point16Val0StateBlocked
	}
	if strings.TrimSpace(model.HistoricalReplayContext.OriginalDecisionAt) == "" ||
		strings.TrimSpace(model.OriginalDecisionBinding.OriginalDecisionAt) == "" ||
		strings.TrimSpace(model.CurrentSubstitutionGuard.OriginalDecisionAt) == "" ||
		!expectedHistoricalTimeSourcesSafe ||
		!point16Val0ExactTimestampBindingValid(model.HistoricalReplayContext.OriginalDecisionAt, model.OriginalDecisionBinding.OriginalDecisionAt) ||
		!point16Val0ExactTimestampBindingValid(model.HistoricalReplayContext.OriginalDecisionAt, model.CurrentSubstitutionGuard.OriginalDecisionAt) {
		model.OriginalDecisionBindingState = Point16Val0StateBlocked
		model.HistoricalReplayContextState = Point16Val0StateBlocked
		model.CurrentSubstitutionGuardState = Point16Val0StateBlocked
		model.ReplayTaxonomyState = Point16Val0StateBlocked
	}
	if !expectedReplayProvenanceExactBound {
		model.OriginalDecisionBindingState = Point16Val0StateBlocked
		model.ReplayTaxonomyState = Point16Val0StateBlocked
	}
	model.ReplayTaxonomyState = point16Val0AggregateFoundationStates(
		model.ReplayTaxonomyState,
		model.HistoricalReplayContextState,
		model.OriginalDecisionBindingState,
	)

	model.ReplayReadinessEvaluation.DependencyState = model.DependencyState
	model.ReplayReadinessEvaluation.HistoricalReplayContextState = model.HistoricalReplayContextState
	model.ReplayReadinessEvaluation.OriginalDecisionBindingState = model.OriginalDecisionBindingState
	model.ReplayReadinessEvaluation.ReplayTaxonomyState = model.ReplayTaxonomyState
	model.ReplayReadinessEvaluation.CurrentSubstitutionGuardState = model.CurrentSubstitutionGuardState
	model.ReplayReadinessEvaluation.NoOverclaimState = model.NoOverclaimState
	model.ReplayReadinessEvaluation.OriginalContextExactBound = expectedOriginalContextComplete && expectedReplayProvenanceExactBound
	model.ReplayReadinessEvaluation.TenantSafe = expectedTenantScopeMatches && expectedTenant != ""
	model.ReplayReadinessEvaluation.TimestampSafe = model.ReplayTaxonomy.TimestampSafe && expectedTimestampConsistent
	model.ReplayReadinessEvaluation.HashSafe = expectedEvidenceHashMatches && expectedPolicyHashMatches && expectedEngineHashMatches
	model.ReplayReadinessEvaluation.GovernanceSafe = expectedGovernanceScopeMatches
	model.ReplayReadinessEvaluation.LineageSafe = point16Val0LineageRefValid(model.HistoricalReplayContext.LineageRef) &&
		point16Val0LineageRefValid(model.OriginalDecisionBinding.LineageRef) &&
		point16Val0ExactMatch(model.HistoricalReplayContext.LineageRef, model.OriginalDecisionBinding.LineageRef)
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

package guidance

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	ModeDisabled      = "disabled"
	ModeLocalTemplate = "local-template"

	CategoryPolicy        = "policy"
	CategoryVulnerability = "vulnerability"
	CategorySigning       = "signing"
	CategoryRuntime       = "runtime"
	CategoryException     = "exception"
	CategoryArtifact      = "artifact"
	CategoryContext       = "context"
	CategoryBreakGlass    = "break_glass"
	CategoryScorecard     = "scorecard"
	CategoryShiftLeft     = "shift_left"

	PriorityCritical = "critical"
	PriorityHigh     = "high"
	PriorityMedium   = "medium"
	PriorityLow      = "low"

	ConfidenceHigh    = "high"
	ConfidenceMedium  = "medium"
	ConfidenceLow     = "low"
	ConfidenceLimited = "limited"

	ReasonGuidanceVulnerabilityActionable = "guidance_vulnerability_actionable"
	ReasonGuidanceSignerGovernanceGap     = "guidance_signer_governance_gap"
	ReasonGuidanceRuntimeContainment      = "guidance_runtime_containment"
	ReasonGuidanceBreakGlassActive        = "guidance_break_glass_active"
	ReasonGuidancePolicyFix               = "guidance_policy_fix"
	ReasonGuidanceArtifactIntegrityGap    = "guidance_artifact_integrity_gap"
	ReasonGuidanceVEXDraftCandidate       = "guidance_vex_draft_candidate"
	ReasonGuidanceMissingContext          = "guidance_missing_context"

	templateVersion = "8l-local-guidance-v1"
)

var (
	bearerTokenPattern = regexp.MustCompile(`(?i)bearer\s+[a-z0-9._\-]+`)
	secretPairPattern  = regexp.MustCompile(`(?i)(token|secret|password|apikey|api_key)\s*[:=]\s*[^,\s]+`)
)

type Config struct {
	Mode            string
	MaxItems        int
	IncludeDocs     bool
	RedactSensitive bool
}

type Scope struct {
	ScopeType   string
	ScopeRef    string
	TenantID    string
	ClusterID   string
	Environment string
	Repository  string
}

type InputFact struct {
	ID                 string            `json:"id"`
	Category           string            `json:"category"`
	SourceComponent    string            `json:"source_component,omitempty"`
	RelatedReasonCodes []string          `json:"related_reason_codes,omitempty"`
	FindingRefs        []string          `json:"finding_refs,omitempty"`
	EvidenceRefs       []string          `json:"evidence_refs,omitempty"`
	DocsRefs           []string          `json:"docs_refs,omitempty"`
	ScopeType          string            `json:"scope_type,omitempty"`
	ScopeRef           string            `json:"scope_ref,omitempty"`
	TenantID           string            `json:"tenant_id,omitempty"`
	ClusterID          string            `json:"cluster_id,omitempty"`
	Environment        string            `json:"environment,omitempty"`
	Repository         string            `json:"repository,omitempty"`
	Severity           string            `json:"severity,omitempty"`
	Summary            string            `json:"summary,omitempty"`
	Detail             string            `json:"detail,omitempty"`
	Metadata           map[string]string `json:"metadata,omitempty"`
	Blocking           bool              `json:"blocking"`
	Deterministic      bool              `json:"deterministic"`
}

type Grouping struct {
	Key                 string `json:"key"`
	Label               string `json:"label"`
	Category            string `json:"category"`
	FindingCount        int    `json:"finding_count"`
	Priority            string `json:"priority"`
	ContextualRiskScore int    `json:"contextual_risk_score"`
	Heuristic           bool   `json:"heuristic"`
}

type VEXDraftSuggestion struct {
	ID                 string   `json:"id"`
	CandidateStatus    string   `json:"candidate_status"`
	Justification      string   `json:"justification"`
	ImpactStatement    string   `json:"impact_statement"`
	MissingEvidence    []string `json:"missing_evidence,omitempty"`
	Confidence         string   `json:"confidence"`
	ConfidenceBasis    string   `json:"confidence_basis,omitempty"`
	AdvisoryOnly       bool     `json:"advisory_only"`
	RequiresHumanReview bool    `json:"requires_human_review"`
	DocsRefs           []string `json:"docs_refs,omitempty"`
}

type BreakGlassGuidance struct {
	ScopeExplanation      string   `json:"scope_explanation"`
	NarrowerAlternative   string   `json:"narrower_alternative,omitempty"`
	CleanupReminders      []string `json:"cleanup_reminders,omitempty"`
	ProposedContainment   []string `json:"proposed_containment_steps,omitempty"`
	Confidence            string   `json:"confidence"`
	ConfidenceBasis       string   `json:"confidence_basis,omitempty"`
	AdvisoryOnly          bool     `json:"advisory_only"`
	RequiresHumanReview   bool     `json:"requires_human_review"`
	DocsRefs              []string `json:"docs_refs,omitempty"`
}

type Item struct {
	ID                    string               `json:"id"`
	Category              string               `json:"category"`
	SourceComponent       string               `json:"source_component,omitempty"`
	Grouping              Grouping             `json:"grouping"`
	RelatedReasonCodes    []string             `json:"related_reason_codes,omitempty"`
	FindingRefs           []string             `json:"finding_refs,omitempty"`
	EvidenceRefs          []string             `json:"evidence_refs,omitempty"`
	DocsRefs              []string             `json:"docs_refs,omitempty"`
	ScopeType             string               `json:"scope_type,omitempty"`
	ScopeRef              string               `json:"scope_ref,omitempty"`
	TenantID              string               `json:"tenant_id,omitempty"`
	ClusterID             string               `json:"cluster_id,omitempty"`
	Environment           string               `json:"environment,omitempty"`
	Repository            string               `json:"repository,omitempty"`
	Severity              string               `json:"severity,omitempty"`
	Priority              string               `json:"priority"`
	Confidence            string               `json:"confidence"`
	ConfidenceBasis       string               `json:"confidence_basis,omitempty"`
	Explanation           string               `json:"explanation,omitempty"`
	RecommendationSummary string               `json:"recommendation_summary,omitempty"`
	RecommendationSteps   []string             `json:"recommendation_steps,omitempty"`
	SaferAlternative      string               `json:"safer_alternative,omitempty"`
	ImpactSummary         string               `json:"impact_summary,omitempty"`
	DataLimitations       []string             `json:"data_limitations,omitempty"`
	AdvisoryOnly          bool                 `json:"advisory_only"`
	RequiresHumanReview   bool                 `json:"requires_human_review"`
	GeneratedAt           time.Time            `json:"generated_at"`
	GeneratedBy           string               `json:"generated_by"`
	TemplateVersion       string               `json:"template_version,omitempty"`
	Heuristic             bool                 `json:"heuristic"`
	VEXDraft              *VEXDraftSuggestion  `json:"vex_draft,omitempty"`
	BreakGlassGuidance    *BreakGlassGuidance  `json:"break_glass_guidance,omitempty"`
}

type Summary struct {
	TotalItems        int            `json:"total_items"`
	CountsByCategory  map[string]int `json:"counts_by_category,omitempty"`
	CountsByPriority  map[string]int `json:"counts_by_priority,omitempty"`
	GuidanceMode      string         `json:"guidance_mode"`
	AIEnabled         bool           `json:"ai_enabled"`
	DeterministicOnly bool           `json:"deterministic_only"`
	Limitations       []string       `json:"limitations,omitempty"`
}

type Response struct {
	GeneratedAt time.Time `json:"generated_at"`
	ScopeType   string    `json:"scope_type,omitempty"`
	ScopeRef    string    `json:"scope_ref,omitempty"`
	TenantID    string    `json:"tenant_id,omitempty"`
	ClusterID   string    `json:"cluster_id,omitempty"`
	Environment string    `json:"environment,omitempty"`
	Repository  string    `json:"repository,omitempty"`
	Items       []Item    `json:"items"`
	Summary     Summary   `json:"summary"`
}

func ParseConfig(getenv func(string) string) (Config, error) {
	if getenv == nil {
		getenv = os.Getenv
	}
	mode := strings.ToLower(strings.TrimSpace(firstNonEmpty(getenv("CHANGELOCK_AI_GUIDANCE_MODE"), ModeDisabled)))
	switch mode {
	case ModeDisabled, ModeLocalTemplate:
	default:
		return Config{}, fmt.Errorf("invalid CHANGELOCK_AI_GUIDANCE_MODE")
	}
	maxItems := parseInt(firstNonEmpty(getenv("CHANGELOCK_AI_GUIDANCE_MAX_ITEMS"), "12"), 12)
	if maxItems <= 0 || maxItems > 100 {
		return Config{}, fmt.Errorf("invalid CHANGELOCK_AI_GUIDANCE_MAX_ITEMS")
	}
	includeDocs, err := parseBool(firstNonEmpty(getenv("CHANGELOCK_AI_GUIDANCE_INCLUDE_DOC_LINKS"), "true"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid CHANGELOCK_AI_GUIDANCE_INCLUDE_DOC_LINKS")
	}
	redactSensitive, err := parseBool(firstNonEmpty(getenv("CHANGELOCK_AI_GUIDANCE_REDACT_SENSITIVE"), "true"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid CHANGELOCK_AI_GUIDANCE_REDACT_SENSITIVE")
	}
	return Config{
		Mode:            mode,
		MaxItems:        maxItems,
		IncludeDocs:     includeDocs,
		RedactSensitive: redactSensitive,
	}, nil
}

func Build(scope Scope, facts []InputFact, config Config, now time.Time) Response {
	if now.IsZero() {
		now = time.Now().UTC()
	}
	scope = normalizeScope(scope)
	config = normalizeConfig(config)
	groups := map[string][]InputFact{}
	for _, fact := range facts {
		normalized := normalizeFact(fact, scope, config)
		if normalized.ID == "" {
			continue
		}
		key := groupingKey(normalized)
		groups[key] = append(groups[key], normalized)
	}
	items := make([]Item, 0, len(groups))
	for key, grouped := range groups {
		item := buildItem(key, grouped, config, now)
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		if priorityRank(items[i].Priority) == priorityRank(items[j].Priority) {
			if items[i].Grouping.ContextualRiskScore == items[j].Grouping.ContextualRiskScore {
				return items[i].ID < items[j].ID
			}
			return items[i].Grouping.ContextualRiskScore > items[j].Grouping.ContextualRiskScore
		}
		return priorityRank(items[i].Priority) > priorityRank(items[j].Priority)
	})
	if len(items) > config.MaxItems {
		items = items[:config.MaxItems]
	}
	return Response{
		GeneratedAt: now,
		ScopeType:   scope.ScopeType,
		ScopeRef:    scope.ScopeRef,
		TenantID:    scope.TenantID,
		ClusterID:   scope.ClusterID,
		Environment: scope.Environment,
		Repository:  scope.Repository,
		Items:       items,
		Summary:     summarize(items, config),
	}
}

func normalizeScope(scope Scope) Scope {
	scope.ScopeType = strings.TrimSpace(firstNonEmpty(scope.ScopeType, "global"))
	scope.ScopeRef = strings.TrimSpace(firstNonEmpty(scope.ScopeRef, scope.ScopeType+":default"))
	scope.TenantID = strings.TrimSpace(scope.TenantID)
	scope.ClusterID = strings.TrimSpace(scope.ClusterID)
	scope.Environment = strings.TrimSpace(scope.Environment)
	scope.Repository = strings.TrimSpace(scope.Repository)
	return scope
}

func normalizeConfig(config Config) Config {
	switch config.Mode {
	case ModeDisabled, ModeLocalTemplate:
	default:
		config.Mode = ModeDisabled
	}
	if config.MaxItems <= 0 {
		config.MaxItems = 12
	}
	return config
}

func normalizeFact(fact InputFact, scope Scope, config Config) InputFact {
	fact.ID = strings.TrimSpace(firstNonEmpty(fact.ID, stableID(scope.ScopeRef, firstNonEmpty(fact.Category, CategoryContext), strings.Join(fact.RelatedReasonCodes, "|"), fact.Summary)))
	fact.Category = normalizeCategory(fact.Category)
	fact.SourceComponent = strings.TrimSpace(fact.SourceComponent)
	fact.ScopeType = strings.TrimSpace(firstNonEmpty(fact.ScopeType, scope.ScopeType))
	fact.ScopeRef = strings.TrimSpace(firstNonEmpty(fact.ScopeRef, scope.ScopeRef))
	fact.TenantID = strings.TrimSpace(firstNonEmpty(fact.TenantID, scope.TenantID))
	fact.ClusterID = strings.TrimSpace(firstNonEmpty(fact.ClusterID, scope.ClusterID))
	fact.Environment = strings.TrimSpace(firstNonEmpty(fact.Environment, scope.Environment))
	fact.Repository = strings.TrimSpace(firstNonEmpty(fact.Repository, scope.Repository))
	fact.Severity = normalizeSeverity(fact.Severity)
	fact.RelatedReasonCodes = uniqueStrings(fact.RelatedReasonCodes)
	fact.FindingRefs = uniqueStrings(fact.FindingRefs)
	fact.EvidenceRefs = uniqueStrings(fact.EvidenceRefs)
	fact.DocsRefs = uniqueStrings(fact.DocsRefs)
	fact.Summary = strings.TrimSpace(fact.Summary)
	fact.Detail = strings.TrimSpace(fact.Detail)
	if config.RedactSensitive {
		fact.Summary = redact(fact.Summary)
		fact.Detail = redact(fact.Detail)
	}
	if !config.IncludeDocs {
		fact.DocsRefs = nil
	}
	if fact.Metadata == nil {
		fact.Metadata = map[string]string{}
	}
	if config.RedactSensitive {
		sanitized := map[string]string{}
		for key, value := range fact.Metadata {
			sanitized[strings.TrimSpace(key)] = redact(strings.TrimSpace(value))
		}
		fact.Metadata = sanitized
	}
	return fact
}

func buildItem(key string, facts []InputFact, config Config, now time.Time) Item {
	dominant := dominantFact(facts)
	riskScore := contextualRiskScore(facts)
	priority := priorityForRisk(riskScore)
	confidence, confidenceBasis := confidenceForFacts(facts)
	recommendationSummary, recommendationSteps, saferAlternative, impactSummary, limitations := recommendationForFacts(facts, dominant)

	item := Item{
		ID:                    stableID(key, dominant.ScopeRef, dominant.Category),
		Category:              dominant.Category,
		SourceComponent:       dominant.SourceComponent,
		Grouping:              buildGrouping(key, facts, dominant.Category, priority, riskScore),
		RelatedReasonCodes:    collectStrings(facts, func(f InputFact) []string { return f.RelatedReasonCodes }),
		FindingRefs:           collectStrings(facts, func(f InputFact) []string { return f.FindingRefs }),
		EvidenceRefs:          collectStrings(facts, func(f InputFact) []string { return f.EvidenceRefs }),
		DocsRefs:              collectStrings(facts, func(f InputFact) []string { return f.DocsRefs }),
		ScopeType:             dominant.ScopeType,
		ScopeRef:              dominant.ScopeRef,
		TenantID:              dominant.TenantID,
		ClusterID:             dominant.ClusterID,
		Environment:           dominant.Environment,
		Repository:            dominant.Repository,
		Severity:              dominant.Severity,
		Priority:              priority,
		Confidence:            confidence,
		ConfidenceBasis:       confidenceBasis,
		RecommendationSummary: recommendationSummary,
		RecommendationSteps:   recommendationSteps,
		SaferAlternative:      saferAlternative,
		ImpactSummary:         impactSummary,
		DataLimitations:       uniqueStrings(limitations),
		AdvisoryOnly:          true,
		RequiresHumanReview:   requiresHumanReview(facts),
		GeneratedAt:           now,
		GeneratedBy:           generationLabel(config.Mode),
		TemplateVersion:       templateVersion,
		Heuristic:             true,
	}
	if config.Mode == ModeLocalTemplate {
		item.Explanation = explanationForFacts(facts, item)
	}
	if draft := buildVEXDraft(facts, item); draft != nil {
		item.VEXDraft = draft
	}
	if guidance := buildBreakGlassGuidance(facts, item); guidance != nil {
		item.BreakGlassGuidance = guidance
	}
	if item.Explanation == "" {
		item.Explanation = defaultExplanation(facts, item)
	}
	return item
}

func buildGrouping(key string, facts []InputFact, category, priority string, riskScore int) Grouping {
	return Grouping{
		Key:                 key,
		Label:               groupingLabel(category, facts),
		Category:            category,
		FindingCount:        len(facts),
		Priority:            priority,
		ContextualRiskScore: riskScore,
		Heuristic:           true,
	}
}

func dominantFact(facts []InputFact) InputFact {
	if len(facts) == 0 {
		return InputFact{Category: CategoryContext, Severity: "low"}
	}
	sort.SliceStable(facts, func(i, j int) bool {
		if severityRank(facts[i].Severity) == severityRank(facts[j].Severity) {
			if facts[i].Blocking == facts[j].Blocking {
				return facts[i].ID < facts[j].ID
			}
			return facts[i].Blocking
		}
		return severityRank(facts[i].Severity) > severityRank(facts[j].Severity)
	})
	return facts[0]
}

func contextualRiskScore(facts []InputFact) int {
	score := 0
	for _, fact := range facts {
		current := severityBase(fact.Severity)
		if fact.Blocking {
			current += 10
		}
		for _, reason := range fact.RelatedReasonCodes {
			switch {
			case strings.Contains(reason, "critical"), strings.Contains(reason, "quarantine"), strings.Contains(reason, "unauthorized"), strings.Contains(reason, "signer_policy_missing"):
				current += 10
			case strings.Contains(reason, "threshold_breached"), strings.Contains(reason, "actionable"), strings.Contains(reason, "violation"), strings.Contains(reason, "gap"):
				current += 6
			case strings.Contains(reason, "unknown"), strings.Contains(reason, "unavailable"), strings.Contains(reason, "partial"):
				current -= 4
			}
		}
		if value := parseInt(fact.Metadata["actionable_count"], 0); value > 0 {
			current += min(value*3, 12)
		}
		if value := parseInt(fact.Metadata["stale_exception_count"], 0); value > 0 {
			current += min(value*2, 10)
		}
		if value := parseInt(fact.Metadata["quarantined_count"], 0); value > 0 {
			current += min(value*5, 15)
		}
		score = max(score, clamp(current, 5, 100))
	}
	return score
}

func priorityForRisk(score int) string {
	switch {
	case score >= 85:
		return PriorityCritical
	case score >= 70:
		return PriorityHigh
	case score >= 50:
		return PriorityMedium
	default:
		return PriorityLow
	}
}

func confidenceForFacts(facts []InputFact) (string, string) {
	unknownSignals := 0
	evidenceRefs := 0
	deterministicCount := 0
	for _, fact := range facts {
		evidenceRefs += len(fact.EvidenceRefs)
		if fact.Deterministic {
			deterministicCount++
		}
		for _, reason := range fact.RelatedReasonCodes {
			if strings.Contains(reason, "unknown") || strings.Contains(reason, "unavailable") || strings.Contains(reason, "missing_context") {
				unknownSignals++
			}
		}
	}
	switch {
	case deterministicCount == len(facts) && evidenceRefs > 0 && unknownSignals == 0:
		return ConfidenceHigh, "Derived directly from deterministic findings with linked evidence references."
	case unknownSignals > 0:
		return ConfidenceLimited, "Underlying findings include unknown or unavailable context; guidance stays conservative."
	case evidenceRefs > 0 && unknownSignals <= 1:
		return ConfidenceMedium, "Derived from deterministic findings, but some context remains heuristic or grouped."
	default:
		return ConfidenceLow, "Guidance is based on partial deterministic context without enough corroborating evidence references."
	}
}

func recommendationForFacts(facts []InputFact, dominant InputFact) (string, []string, string, string, []string) {
	reason := firstReasonCode(facts)
	if hasReason(facts, ReasonGuidanceMissingContext) {
		return "Gather the missing deterministic context before broadening policy, VEX, or exception scope.", []string{
			"Re-run the underlying deterministic check and confirm API-assisted or runtime context is reachable.",
			"Do not treat unavailable context as evidence that the issue is safe to suppress.",
			"Use the linked deterministic reason codes as the source of truth until context collection succeeds.",
		}, "Prefer waiting for complete deterministic context over approving a broader exception or trust change.", "Incomplete context can hide real deploy, runtime, or vulnerability risk.", []string{"Current guidance is intentionally conservative because one or more supporting context sources were unavailable."}
	}
	switch {
	case dominant.Category == CategoryVulnerability || strings.Contains(reason, "vulnerability"):
		actionable := parseInt(dominant.Metadata["actionable_count"], 0)
		resolved := parseInt(dominant.Metadata["resolved_by_vex_count"], 0)
		limitations := []string{"Runtime reachability is not proven unless explicit runtime exposure evidence is present."}
		summary := "Triage the net actionable vulnerability set first, then decide whether remediation or a scoped VEX draft is justified."
		steps := []string{
			"Confirm the affected digest, package, and scope before changing policy or exceptions.",
			"Prefer updating the vulnerable package or image digest before considering a compensating statement.",
			"If the finding may be non-actionable, generate a draft VEX explanation and attach supporting evidence for human review.",
		}
		if actionable == 0 && resolved > 0 {
			summary = "Keep the existing VEX-backed posture under review and confirm that the supporting evidence still matches the deployed digest."
			steps = []string{
				"Review the active VEX statements for scope drift, expiry, and supporting evidence freshness.",
				"Confirm that the current digest still matches the resolved findings.",
				"Do not suppress future findings without a new scoped VEX review.",
			}
		}
		return summary, steps, "Prefer digest-scoped remediation or a tightly scoped VEX statement over broad risk acceptance.", "Unresolved vulnerability posture can affect deploy approval, scorecard grade, and runtime containment decisions.", limitations
	case dominant.Category == CategorySigning || strings.Contains(reason, "signer"):
		return "Restore explicit signer governance before widening trust or ignoring signer anomalies.", []string{
			"Review the observed signer identity, repository, workflow, and ref against the allowed policy.",
			"Prefer adding or tightening explicit signer policy over bypassing identity checks.",
			"If a signer was compromised or retired, set a distrust cutoff instead of relying on certificate revocation semantics.",
		}, "Prefer exact issuer/repository/workflow policy alignment over broad repository-based trust.", "Unauthorized or unknown signers can invalidate deploy-time trust and may trigger runtime containment.", []string{"Workflow drift findings are advisory unless signer enforcement is configured to enforce."}
	case dominant.Category == CategoryRuntime || strings.Contains(reason, "runtime") || strings.Contains(reason, "quarantine"):
		return "Treat runtime containment as a signal to reconcile desired state or investigate a repeated drift source.", []string{
			"Confirm whether the parent controller spec or pod runtime state diverged from the approved desired state.",
			"Prefer patching back to approved state over broad restarts or disabling closed-loop controls.",
			"If containment is active, keep the workload scoped and investigate recurrence before releasing quarantine.",
		}, "Prefer spec correction and bounded quarantine over disabling runtime controls globally.", "Runtime failures or quarantine can indicate either malicious drift or repeated reconciliation instability.", []string{"Network quarantine guidance is only a containment intent unless the cluster enforces NetworkPolicy."}
	case dominant.Category == CategoryException || dominant.Category == CategoryBreakGlass || strings.Contains(reason, "exception"):
		return "Reduce exception scope and duration before extending or reusing emergency access.", []string{
			"Confirm the exception still matches the intended tenant, repo, digest, CVE, or namespace scope.",
			"Prefer the narrowest digest, CVE, and time-bounded scope that still addresses the incident.",
			"Plan explicit cleanup and post-incident review before the exception expires.",
		}, "Prefer a narrow, time-bounded exception plus remediation plan over a broad break-glass extension.", "Exception sprawl increases audit debt and can lower hardening posture even when approvals are formally recorded.", []string{"Guidance does not approve or extend the exception; it only summarizes safer review options."}
	case dominant.Category == CategoryPolicy || strings.Contains(reason, "manifest_policy_violation") || strings.Contains(reason, "policy"):
		return "Fix the manifest or policy mismatch directly instead of normalizing the exception path.", []string{
			"Review the denying rule, resource scope, and fix hint from the deterministic diagnostic.",
			"Prefer policy-compliant manifest changes, such as digest pinning or reduced privileges, before asking for a bypass.",
			"Re-run preflight after the smallest safe change and compare the deterministic reason codes.",
		}, "Prefer a narrow manifest fix over policy relaxation or a break-glass request.", "Policy failures usually indicate a bounded configuration issue that can be fixed earlier than deploy time.", []string{"Line precision is file-scoped when the underlying deterministic check cannot prove a narrower range."}
	default:
		return "Review the deterministic findings and linked evidence before making a trust or deployment change.", []string{
			"Inspect the linked reason codes and evidence references first.",
			"Prefer the smallest explicit corrective action that resolves the measured finding.",
			"If context is missing, gather it before broadening policy or exceptions.",
		}, "Prefer explicit scoped remediation over broad trust expansion.", "The current context is incomplete, so recommendations stay conservative.", []string{"Some findings were grouped heuristically for operator readability."}
	}
}

func explanationForFacts(facts []InputFact, item Item) string {
	reason := firstReasonCode(facts)
	if hasReason(facts, ReasonGuidanceMissingContext) {
		return fmt.Sprintf("This guidance stays conservative for %s because one or more deterministic context sources were unavailable or incomplete.", joinScope(item))
	}
	switch item.Category {
	case CategoryVulnerability:
		return fmt.Sprintf("This guidance groups vulnerability posture signals around %s. The current priority is %s because deterministic findings still show unresolved exposure after VEX-aware evaluation.", joinScope(item), item.Priority)
	case CategorySigning:
		return fmt.Sprintf("This guidance groups signer identity evidence for %s. The recommendation stays conservative because signer authorization remains a trust anchor rather than a heuristic signal.", joinScope(item))
	case CategoryRuntime:
		return fmt.Sprintf("This guidance groups runtime reconciliation and containment signals for %s. The suggested actions favor restoring approved state and keeping quarantine bounded.", joinScope(item))
	case CategoryException, CategoryBreakGlass:
		return fmt.Sprintf("This guidance explains active exception posture for %s. It focuses on narrowing scope, cleanup reminders, and post-incident review rather than approval logic.", joinScope(item))
	case CategoryPolicy:
		return fmt.Sprintf("This guidance summarizes policy and shift-left diagnostics for %s. It keeps deterministic denial reasons authoritative and only adds bounded remediation context.", joinScope(item))
	default:
		if reason != "" {
			return fmt.Sprintf("This guidance is a bounded interpretation of deterministic ChangeLock findings with reason code %s.", reason)
		}
		return "This guidance is a bounded interpretation of deterministic ChangeLock findings."
	}
}

func hasReason(facts []InputFact, reason string) bool {
	for _, fact := range facts {
		for _, code := range fact.RelatedReasonCodes {
			if strings.TrimSpace(code) == strings.TrimSpace(reason) {
				return true
			}
		}
	}
	return false
}

func defaultExplanation(facts []InputFact, item Item) string {
	reason := firstReasonCode(facts)
	if reason == "" {
		return "Deterministic findings were grouped to provide bounded, advisory context."
	}
	return fmt.Sprintf("Deterministic findings with reason code %s were grouped to provide bounded, advisory context.", reason)
}

func buildVEXDraft(facts []InputFact, item Item) *VEXDraftSuggestion {
	if item.Category != CategoryVulnerability {
		return nil
	}
	actionable := 0
	resolvedByVEX := 0
	for _, fact := range facts {
		actionable = max(actionable, parseInt(fact.Metadata["actionable_count"], 0))
		resolvedByVEX = max(resolvedByVEX, parseInt(fact.Metadata["resolved_by_vex_count"], 0))
	}
	if actionable == 0 || resolvedByVEX > 0 {
		return nil
	}
	return &VEXDraftSuggestion{
		ID:                  stableID(item.ID, "vex-draft"),
		CandidateStatus:     "under_investigation",
		Justification:       "The current deterministic context does not prove non-exploitability. Review package scope, runtime exposure, and any mitigating controls before choosing a final VEX status.",
		ImpactStatement:     "The finding remains actionable until supporting evidence justifies a narrower disposition.",
		MissingEvidence:     []string{"runtime reachability proof", "product-specific mitigation evidence", "package usage confirmation"},
		Confidence:          ConfidenceLimited,
		ConfidenceBasis:     "A conservative draft is suggested because the repository can prove the finding exists but not that it is unreachable or not affected.",
		AdvisoryOnly:        true,
		RequiresHumanReview: true,
		DocsRefs:            []string{"docs/vex-exploitability-ops.md"},
	}
}

func buildBreakGlassGuidance(facts []InputFact, item Item) *BreakGlassGuidance {
	if item.Category != CategoryException && item.Category != CategoryBreakGlass {
		return nil
	}
	exceptionType := ""
	activeCount := 0
	for _, fact := range facts {
		exceptionType = firstNonEmpty(exceptionType, fact.Metadata["exception_type"])
		activeCount = max(activeCount, parseInt(fact.Metadata["active_exception_count"], 0))
	}
	scope := joinScope(item)
	scopeExplanation := fmt.Sprintf("Current exception guidance applies to %s.", scope)
	if exceptionType != "" {
		scopeExplanation = fmt.Sprintf("Current %s exception guidance applies to %s.", exceptionType, scope)
	}
	return &BreakGlassGuidance{
		ScopeExplanation:    scopeExplanation,
		NarrowerAlternative: "Prefer digest-, CVE-, or namespace-scoped temporary access over broader repository- or environment-wide emergency exceptions.",
		CleanupReminders: []string{
			"Confirm the expiry still matches the real incident window.",
			"Revoke or replace the exception once the bounded remediation lands.",
			"Capture post-incident review notes while the evidence is still fresh.",
		},
		ProposedContainment: []string{
			"Keep the exception time-bounded and tied to the smallest affected scope.",
			"Prefer runtime containment or signer/VEX review over broad allow rules when the incident allows it.",
		},
		Confidence:          ternary(activeCount > 0, ConfidenceMedium, ConfidenceLimited),
		ConfidenceBasis:     "Exception scope and lifecycle are explicit, but the safest narrower alternative depends on the incident details the repository cannot infer automatically.",
		AdvisoryOnly:        true,
		RequiresHumanReview: true,
		DocsRefs:            []string{"docs/auth-rbac.md", "docs/runtime-closed-loop-hardening.md"},
	}
}

func summarize(items []Item, config Config) Summary {
	summary := Summary{
		TotalItems:        len(items),
		CountsByCategory:  map[string]int{},
		CountsByPriority:  map[string]int{},
		GuidanceMode:      config.Mode,
		AIEnabled:         config.Mode != ModeDisabled,
		DeterministicOnly: config.Mode == ModeDisabled,
		Limitations: []string{
			"Deterministic findings remain authoritative; guidance is advisory only.",
			"Missing runtime reachability or business context is surfaced as uncertainty rather than inferred certainty.",
		},
	}
	for _, item := range items {
		summary.CountsByCategory[item.Category]++
		summary.CountsByPriority[item.Priority]++
	}
	if config.Mode == ModeDisabled {
		summary.Limitations = append(summary.Limitations, "AI-specific narrative enrichment is disabled; output is generated from bounded deterministic mappings.")
	}
	return summary
}

func normalizeCategory(value string) string {
	switch strings.TrimSpace(strings.ToLower(value)) {
	case CategoryPolicy, CategoryVulnerability, CategorySigning, CategoryRuntime, CategoryException, CategoryArtifact, CategoryContext, CategoryBreakGlass, CategoryScorecard, CategoryShiftLeft:
		return strings.TrimSpace(strings.ToLower(value))
	default:
		return CategoryContext
	}
}

func groupingKey(fact InputFact) string {
	primaryReason := firstNonEmpty(firstReasonCode([]InputFact{fact}), fact.Category)
	return strings.Join([]string{fact.ScopeRef, fact.Category, primaryReason}, "|")
}

func groupingLabel(category string, facts []InputFact) string {
	switch category {
	case CategoryVulnerability:
		return "Contextual vulnerability triage"
	case CategorySigning:
		return "Signer identity governance"
	case CategoryRuntime:
		return "Runtime containment and reconciliation"
	case CategoryException, CategoryBreakGlass:
		return "Break-glass and exception hygiene"
	case CategoryPolicy:
		return "Policy and shift-left remediation"
	case CategoryArtifact:
		return "Artifact integrity and evidence"
	default:
		if len(facts) > 0 && facts[0].Category == CategoryShiftLeft {
			return "Developer guidance"
		}
		return "Contextual guidance"
	}
}

func firstReasonCode(facts []InputFact) string {
	for _, fact := range facts {
		for _, reason := range fact.RelatedReasonCodes {
			if strings.TrimSpace(reason) != "" {
				return strings.TrimSpace(reason)
			}
		}
	}
	return ""
}

func collectStrings[T any](items []T, fn func(T) []string) []string {
	combined := make([]string, 0)
	for _, item := range items {
		combined = append(combined, fn(item)...)
	}
	return uniqueStrings(combined)
}

func uniqueStrings(values []string) []string {
	seen := map[string]struct{}{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		result = append(result, trimmed)
	}
	sort.Strings(result)
	return result
}

func priorityRank(priority string) int {
	switch strings.TrimSpace(strings.ToLower(priority)) {
	case PriorityCritical:
		return 4
	case PriorityHigh:
		return 3
	case PriorityMedium:
		return 2
	default:
		return 1
	}
}

func normalizeSeverity(value string) string {
	switch strings.TrimSpace(strings.ToLower(value)) {
	case "critical", "high", "medium", "warning", "error", "low", "note":
		return strings.TrimSpace(strings.ToLower(value))
	default:
		return "low"
	}
}

func severityRank(severity string) int {
	switch normalizeSeverity(severity) {
	case "critical":
		return 6
	case "error":
		return 5
	case "high":
		return 4
	case "medium", "warning":
		return 3
	case "low":
		return 2
	default:
		return 1
	}
}

func severityBase(severity string) int {
	switch normalizeSeverity(severity) {
	case "critical":
		return 88
	case "error", "high":
		return 75
	case "medium", "warning":
		return 58
	case "low":
		return 40
	default:
		return 25
	}
}

func requiresHumanReview(facts []InputFact) bool {
	for _, fact := range facts {
		for _, reason := range fact.RelatedReasonCodes {
			if strings.Contains(reason, "exception") || strings.Contains(reason, "break_glass") || strings.Contains(reason, "vulnerability") || strings.Contains(reason, "signer") || strings.Contains(reason, "runtime") {
				return true
			}
		}
		if fact.Blocking {
			return true
		}
	}
	return false
}

func redact(value string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return ""
	}
	trimmed = bearerTokenPattern.ReplaceAllString(trimmed, "Bearer [redacted]")
	trimmed = secretPairPattern.ReplaceAllStringFunc(trimmed, func(match string) string {
		parts := strings.SplitN(match, "=", 2)
		if len(parts) == 2 {
			return strings.TrimSpace(parts[0]) + "=[redacted]"
		}
		parts = strings.SplitN(match, ":", 2)
		if len(parts) == 2 {
			return strings.TrimSpace(parts[0]) + ": [redacted]"
		}
		return "[redacted]"
	})
	return trimmed
}

func generationLabel(mode string) string {
	if mode == ModeLocalTemplate {
		return "local-template-guidance"
	}
	return "deterministic-guidance"
}

func joinScope(item Item) string {
	switch {
	case item.Repository != "":
		return item.Repository
	case item.ScopeRef != "":
		return item.ScopeRef
	default:
		return item.ScopeType
	}
}

func stableID(parts ...string) string {
	hash := sha1.Sum([]byte(strings.Join(parts, "|")))
	return hex.EncodeToString(hash[:])[:12]
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func parseBool(value string) (bool, error) {
	switch strings.TrimSpace(strings.ToLower(value)) {
	case "", "false", "0", "no":
		return false, nil
	case "true", "1", "yes":
		return true, nil
	default:
		return false, fmt.Errorf("invalid boolean")
	}
}

func parseInt(value string, fallback int) int {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return fallback
	}
	return parsed
}

func clamp(value, minValue, maxValue int) int {
	if value < minValue {
		return minValue
	}
	if value > maxValue {
		return maxValue
	}
	return value
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}

func ternary[T any](condition bool, left, right T) T {
	if condition {
		return left
	}
	return right
}

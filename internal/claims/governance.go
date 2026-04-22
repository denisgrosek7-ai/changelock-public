package claims

import (
	"strings"
	"time"
)

const (
	ScopePublic   = "public"
	ScopePartner  = "partner"
	ScopeAuditor  = "auditor"
	ScopeInternal = "internal"

	StateReady   = "claim_ready"
	StateStale   = "claim_stale"
	StateBlocked = "claim_blocked"

	FreshnessFresh   = "fresh"
	FreshnessStale   = "stale"
	FreshnessExpired = "expired"

	VisibilityPublicSafe  = "public_safe"
	VisibilityPartnerSafe = "partner_safe"
	VisibilityAuditorOnly = "auditor_only"
)

type Input struct {
	ClaimID                  string    `json:"claim_id"`
	ClaimClass               string    `json:"claim_class"`
	Scope                    string    `json:"scope"`
	PublicationClass         string    `json:"publication_class,omitempty"`
	VerifiedAt               time.Time `json:"verified_at,omitempty"`
	ValidUntil               time.Time `json:"valid_until,omitempty"`
	EvidenceRefs             []string  `json:"evidence_refs,omitempty"`
	ProofRefs                []string  `json:"proof_refs,omitempty"`
	MethodologyRefs          []string  `json:"methodology_refs,omitempty"`
	VerifierRefs             []string  `json:"verifier_refs,omitempty"`
	MethodologyRef           string    `json:"methodology_ref,omitempty"`
	SupportsIndependentCheck bool      `json:"supports_independent_check"`
	PartnerVisibleOnly       bool      `json:"partner_visible_only"`
	AuditorVisibleOnly       bool      `json:"auditor_visible_only"`
	ContainsSensitiveData    bool      `json:"contains_sensitive_data"`
	BadgeState               string    `json:"badge_state,omitempty"`
}

type Decision struct {
	ClaimID               string    `json:"claim_id"`
	ClaimClass            string    `json:"claim_class"`
	PublicationClass      string    `json:"publication_class,omitempty"`
	CurrentState          string    `json:"current_state"`
	FreshnessState        string    `json:"freshness_state"`
	VisibilityState       string    `json:"visibility_state"`
	AllowedScopes         []string  `json:"allowed_scopes,omitempty"`
	TraceRefs             []string  `json:"trace_refs,omitempty"`
	RequiredPreconditions []string  `json:"required_preconditions,omitempty"`
	VerifiedAt            time.Time `json:"verified_at,omitempty"`
	ValidUntil            time.Time `json:"valid_until,omitempty"`
	ReasonCodes           []string  `json:"reason_codes,omitempty"`
	Limitations           []string  `json:"limitations,omitempty"`
}

func Evaluate(input Input, asOf time.Time) Decision {
	if asOf.IsZero() {
		asOf = time.Now().UTC()
	}
	input.ClaimID = strings.TrimSpace(input.ClaimID)
	input.ClaimClass = strings.TrimSpace(input.ClaimClass)
	input.Scope = strings.TrimSpace(input.Scope)
	input.PublicationClass = strings.TrimSpace(input.PublicationClass)
	input.MethodologyRef = strings.TrimSpace(input.MethodologyRef)
	input.BadgeState = strings.TrimSpace(input.BadgeState)

	reasons := make([]string, 0, 6)
	preconditions := make([]string, 0, 4)
	allowedScopes := []string{}
	visibility := VisibilityPublicSafe
	if input.ContainsSensitiveData {
		visibility = VisibilityAuditorOnly
		allowedScopes = []string{ScopeAuditor, ScopeInternal}
		reasons = append(reasons, "customer_sensitive_data_not_public")
	} else if input.AuditorVisibleOnly {
		visibility = VisibilityAuditorOnly
		allowedScopes = []string{ScopeAuditor, ScopeInternal}
		reasons = append(reasons, "auditor_scope_required")
	} else if input.PartnerVisibleOnly {
		visibility = VisibilityPartnerSafe
		allowedScopes = []string{ScopePartner, ScopeAuditor, ScopeInternal}
		reasons = append(reasons, "partner_scope_required")
	} else {
		allowedScopes = []string{ScopePublic, ScopePartner, ScopeAuditor, ScopeInternal}
	}

	freshness := FreshnessFresh
	if !input.ValidUntil.IsZero() {
		switch {
		case asOf.After(input.ValidUntil):
			freshness = FreshnessExpired
		case asOf.After(input.ValidUntil.Add(-14 * 24 * time.Hour)):
			freshness = FreshnessStale
		}
	}

	currentState := StateReady
	if len(input.EvidenceRefs) == 0 && len(input.ProofRefs) == 0 && len(input.VerifierRefs) == 0 {
		currentState = StateBlocked
		reasons = append(reasons, "missing_required_evidence")
	}
	if input.ClaimClass == "verification_claim" || input.ClaimClass == "conformance_claim" || input.ClaimClass == "auditor_ready_claim" {
		preconditions = append(preconditions, "independent_verification_required")
	}
	if input.ClaimClass == "verification_claim" && !input.SupportsIndependentCheck {
		currentState = StateBlocked
		reasons = append(reasons, "independent_verification_not_available")
	}
	if (input.ClaimClass == "benchmark_claim" || input.ClaimClass == "conformance_claim") && input.MethodologyRef == "" && len(input.MethodologyRefs) == 0 {
		currentState = StateBlocked
		reasons = append(reasons, "benchmark_methodology_missing")
	}
	if input.ClaimClass == "benchmark_claim" || input.ClaimClass == "conformance_claim" {
		preconditions = append(preconditions, "methodology_reference_required")
	}
	if input.ClaimClass == "trust_badge_claim" && input.BadgeState != "active" {
		currentState = StateBlocked
		reasons = append(reasons, "trust_badge_not_active")
	}
	if input.ContainsSensitiveData && input.Scope == ScopePublic {
		currentState = StateBlocked
		reasons = append(reasons, "public_scope_redacted")
	}
	if input.Scope != "" && !containsScope(allowedScopes, input.Scope) {
		currentState = StateBlocked
		reasons = append(reasons, "requested_scope_not_permitted")
	}
	if currentState == StateReady && (freshness == FreshnessStale || freshness == FreshnessExpired) {
		currentState = StateStale
		reasons = append(reasons, "claim_freshness_window_elapsed")
	}
	if visibility == VisibilityAuditorOnly && input.Scope == ScopePublic {
		currentState = StateBlocked
	}
	if visibility == VisibilityPublicSafe && input.Scope == ScopePartner {
		visibility = VisibilityPartnerSafe
	}
	if currentState == StateReady {
		reasons = append(reasons, "claim_bounded_by_published_evidence")
	}

	return Decision{
		ClaimID:               input.ClaimID,
		ClaimClass:            input.ClaimClass,
		PublicationClass:      firstNonEmpty(input.PublicationClass, input.ClaimClass),
		CurrentState:          currentState,
		FreshnessState:        freshness,
		VisibilityState:       visibility,
		AllowedScopes:         uniqueStrings(allowedScopes),
		TraceRefs:             uniqueStrings(append(append(append([]string{}, input.EvidenceRefs...), input.ProofRefs...), append(input.MethodologyRefs, input.VerifierRefs...)...)),
		RequiredPreconditions: uniqueStrings(preconditions),
		VerifiedAt:            input.VerifiedAt.UTC(),
		ValidUntil:            input.ValidUntil.UTC(),
		ReasonCodes:           uniqueStrings(reasons),
		Limitations: []string{
			"Claims governance remains bounded by published evidence, freshness, and disclosure scope.",
			"Public claims do not become certification or universal-security statements through this decision model.",
		},
	}
}

func AllowsScope(decision Decision, scope string) bool {
	scope = strings.TrimSpace(scope)
	for _, candidate := range decision.AllowedScopes {
		if candidate == scope {
			return true
		}
	}
	return false
}

func uniqueStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]struct{}{}
	items := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		items = append(items, value)
	}
	if len(items) == 0 {
		return nil
	}
	return items
}

func containsScope(values []string, scope string) bool {
	scope = strings.TrimSpace(scope)
	for _, value := range values {
		if value == scope {
			return true
		}
	}
	return false
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			return value
		}
	}
	return ""
}

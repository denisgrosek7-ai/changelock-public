package guidance

import (
	"strings"
	"testing"
	"time"
)

func TestBuildGroupsAndPrioritizesFindings(t *testing.T) {
	now := time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC)
	response := Build(Scope{ScopeType: "repository", ScopeRef: "demo/app", Repository: "demo/app"}, []InputFact{
		{
			ID:                 "vuln-1",
			Category:           CategoryVulnerability,
			Severity:           "critical",
			Summary:            "Threshold breached",
			RelatedReasonCodes: []string{"vulnerability_threshold_breached"},
			EvidenceRefs:       []string{"report:/v1/vulnerabilities/net"},
			Metadata: map[string]string{
				"actionable_count": "3",
			},
			Deterministic: true,
			Blocking:      true,
		},
		{
			ID:                 "signer-1",
			Category:           CategorySigning,
			Severity:           "high",
			Summary:            "Unauthorized signer observed",
			RelatedReasonCodes: []string{"signer_identity_findings_active"},
			EvidenceRefs:       []string{"status:/v1/signing-identities/findings"},
			Deterministic:      true,
		},
	}, Config{Mode: ModeLocalTemplate, MaxItems: 10, IncludeDocs: true, RedactSensitive: true}, now)

	if len(response.Items) != 2 {
		t.Fatalf("expected 2 grouped items, got %d", len(response.Items))
	}
	if response.Items[0].Category != CategoryVulnerability {
		t.Fatalf("expected vulnerability item first, got %#v", response.Items[0])
	}
	if response.Items[0].Priority != PriorityCritical {
		t.Fatalf("expected critical priority, got %#v", response.Items[0])
	}
	if response.Items[0].Confidence != ConfidenceHigh {
		t.Fatalf("expected high confidence, got %#v", response.Items[0])
	}
}

func TestBuildKeepsUnknownContextLimited(t *testing.T) {
	now := time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC)
	response := Build(Scope{ScopeType: "tenant", ScopeRef: "acme"}, []InputFact{
		{
			ID:                 "vuln-unknown",
			Category:           CategoryVulnerability,
			Severity:           "medium",
			Summary:            "VEX context unavailable",
			RelatedReasonCodes: []string{"remote_vex_context_unavailable"},
			Deterministic:      true,
		},
	}, Config{Mode: ModeDisabled, MaxItems: 10, IncludeDocs: true, RedactSensitive: true}, now)

	if len(response.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(response.Items))
	}
	if response.Items[0].Confidence != ConfidenceLimited {
		t.Fatalf("expected limited confidence, got %#v", response.Items[0])
	}
	if !strings.Contains(response.Items[0].Explanation, "reason code") {
		t.Fatalf("expected deterministic explanation fallback, got %q", response.Items[0].Explanation)
	}
	if !response.Summary.DeterministicOnly {
		t.Fatalf("expected deterministic-only summary, got %#v", response.Summary)
	}
}

func TestBuildProducesVEXDraftSuggestion(t *testing.T) {
	now := time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC)
	response := Build(Scope{ScopeType: "repository", ScopeRef: "demo/app"}, []InputFact{
		{
			ID:                 "vuln-1",
			Category:           CategoryVulnerability,
			Severity:           "high",
			Summary:            "Actionable findings remain",
			RelatedReasonCodes: []string{"vulnerability_posture_actionable"},
			Metadata: map[string]string{
				"actionable_count": "2",
			},
			EvidenceRefs:  []string{"report:/v1/vulnerabilities/net"},
			Deterministic: true,
			Blocking:      true,
		},
	}, Config{Mode: ModeLocalTemplate, MaxItems: 10, IncludeDocs: true, RedactSensitive: true}, now)

	if response.Items[0].VEXDraft == nil {
		t.Fatalf("expected vex draft suggestion, got %#v", response.Items[0])
	}
	if response.Items[0].VEXDraft.CandidateStatus != "under_investigation" {
		t.Fatalf("unexpected draft %#v", response.Items[0].VEXDraft)
	}
}

func TestBuildProducesBreakGlassGuidance(t *testing.T) {
	now := time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC)
	response := Build(Scope{ScopeType: "repository", ScopeRef: "demo/app"}, []InputFact{
		{
			ID:       "ex-1",
			Category: CategoryBreakGlass,
			Severity: "medium",
			Summary:  "Active break-glass exception",
			RelatedReasonCodes: []string{
				"stale_exception_active",
			},
			Metadata: map[string]string{
				"exception_type":         "BREAK_GLASS",
				"active_exception_count": "1",
			},
			Deterministic: true,
		},
	}, Config{Mode: ModeLocalTemplate, MaxItems: 10, IncludeDocs: true, RedactSensitive: true}, now)

	if response.Items[0].BreakGlassGuidance == nil {
		t.Fatalf("expected break-glass guidance, got %#v", response.Items[0])
	}
	if !strings.Contains(response.Items[0].BreakGlassGuidance.ScopeExplanation, "BREAK_GLASS") {
		t.Fatalf("unexpected break-glass guidance %#v", response.Items[0].BreakGlassGuidance)
	}
}

func TestBuildRedactsSensitiveContext(t *testing.T) {
	now := time.Date(2026, 4, 17, 12, 0, 0, 0, time.UTC)
	response := Build(Scope{ScopeType: "repository", ScopeRef: "demo/app"}, []InputFact{
		{
			ID:                 "sensitive",
			Category:           CategoryPolicy,
			Severity:           "high",
			Summary:            "Authorization: Bearer secret-token",
			Detail:             "token=abc123",
			RelatedReasonCodes: []string{"manifest_policy_violation"},
			Deterministic:      true,
		},
	}, Config{Mode: ModeLocalTemplate, MaxItems: 10, IncludeDocs: true, RedactSensitive: true}, now)

	item := response.Items[0]
	if strings.Contains(item.Explanation, "secret-token") || strings.Contains(item.RecommendationSummary, "abc123") {
		t.Fatalf("expected redaction, got %#v", item)
	}
}

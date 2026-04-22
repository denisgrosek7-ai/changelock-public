package claims

import (
	"testing"
	"time"
)

func TestEvaluateBlocksPublicSensitiveClaim(t *testing.T) {
	decision := Evaluate(Input{
		ClaimID:               "public-sensitive",
		ClaimClass:            "verification_claim",
		Scope:                 ScopePublic,
		VerifiedAt:            time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC),
		ValidUntil:            time.Date(2026, time.July, 1, 10, 0, 0, 0, time.UTC),
		EvidenceRefs:          []string{"/v1/public/transparency/anchor"},
		ContainsSensitiveData: true,
	}, time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC))

	if decision.CurrentState != StateBlocked {
		t.Fatalf("expected blocked decision, got %#v", decision)
	}
	if AllowsScope(decision, ScopePublic) {
		t.Fatalf("expected public scope to stay blocked, got %#v", decision)
	}
}

func TestEvaluateRequiresIndependentVerificationForVerificationClaim(t *testing.T) {
	decision := Evaluate(Input{
		ClaimID:    "verification",
		ClaimClass: "verification_claim",
		Scope:      ScopePublic,
		VerifiedAt: time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC),
		ValidUntil: time.Date(2026, time.July, 1, 10, 0, 0, 0, time.UTC),
		EvidenceRefs: []string{
			"/v1/public/verifier/reference-pack",
		},
		SupportsIndependentCheck: false,
	}, time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC))

	if decision.CurrentState != StateBlocked {
		t.Fatalf("expected blocked verification claim, got %#v", decision)
	}
}

func TestEvaluateBlocksBenchmarkClaimWithoutMethodology(t *testing.T) {
	decision := Evaluate(Input{
		ClaimID:    "benchmark",
		ClaimClass: "benchmark_claim",
		Scope:      ScopePublic,
		VerifiedAt: time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC),
		ValidUntil: time.Date(2026, time.July, 1, 10, 0, 0, 0, time.UTC),
		EvidenceRefs: []string{
			"/v1/public/benchmarks/packs",
		},
		SupportsIndependentCheck: true,
	}, time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC))

	if decision.CurrentState != StateBlocked {
		t.Fatalf("expected blocked benchmark claim, got %#v", decision)
	}
	if !containsReason(decision.ReasonCodes, "benchmark_methodology_missing") {
		t.Fatalf("expected methodology blocker, got %#v", decision)
	}
}

func TestEvaluateKeepsPartnerAndAuditorScopesSeparated(t *testing.T) {
	partnerDecision := Evaluate(Input{
		ClaimID:                  "partner",
		ClaimClass:               "verification_claim",
		Scope:                    ScopePartner,
		PartnerVisibleOnly:       true,
		VerifiedAt:               time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC),
		ValidUntil:               time.Date(2026, time.July, 1, 10, 0, 0, 0, time.UTC),
		EvidenceRefs:             []string{"/v1/public/specs/federation-proof-exchange"},
		SupportsIndependentCheck: true,
	}, time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC))
	if AllowsScope(partnerDecision, ScopePublic) {
		t.Fatalf("expected partner-only claim to stay out of public scope, got %#v", partnerDecision)
	}

	auditorDecision := Evaluate(Input{
		ClaimID:                  "auditor",
		ClaimClass:               "auditor_ready_claim",
		Scope:                    ScopeAuditor,
		AuditorVisibleOnly:       true,
		VerifiedAt:               time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC),
		ValidUntil:               time.Date(2026, time.July, 1, 10, 0, 0, 0, time.UTC),
		EvidenceRefs:             []string{"/v1/public/auditor/workflows"},
		SupportsIndependentCheck: true,
	}, time.Date(2026, time.April, 22, 10, 0, 0, 0, time.UTC))
	if AllowsScope(auditorDecision, ScopePartner) || AllowsScope(auditorDecision, ScopePublic) {
		t.Fatalf("expected auditor-only claim to stay scoped, got %#v", auditorDecision)
	}
}

func containsReason(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

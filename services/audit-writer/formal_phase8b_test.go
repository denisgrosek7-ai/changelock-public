package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	formalcore "github.com/denisgrosek/changelock/internal/formal"
)

func TestPhase8ConsensusReviewAndPolicySuggestionsHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/governance/consensus-review", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 consensus review 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var consensus phase8ConsensusReviewResponse
	if err := json.NewDecoder(rec.Body).Decode(&consensus); err != nil {
		t.Fatalf("decode phase8 consensus review: %v", err)
	}
	if consensus.CurrentState != phase8ConsensusReviewStateActive {
		t.Fatalf("expected active phase8 consensus slice, got %#v", consensus)
	}
	if consensus.ConsensusState != "consensus_assisted_review_visible" || consensus.MinorityState != "weighted_disagreement_and_minority_report_visible" {
		t.Fatalf("expected bounded consensus states, got %#v", consensus)
	}
	if len(consensus.Items) == 0 || consensus.Items[0].QuorumThreshold == "" || consensus.Items[0].AbstainState != "abstain_visible" {
		t.Fatalf("expected populated consensus items, got %#v", consensus)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/governance/policy-suggestions", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 policy suggestions 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var suggestions phase8PolicySuggestionsResponse
	if err := json.NewDecoder(rec.Body).Decode(&suggestions); err != nil {
		t.Fatalf("decode phase8 policy suggestions: %v", err)
	}
	if suggestions.CurrentState != phase8PolicySuggestionsStateActive {
		t.Fatalf("expected active phase8 policy suggestions slice, got %#v", suggestions)
	}
	if suggestions.ApprovalBoundary != "advisory_until_formally_approved" {
		t.Fatalf("expected advisory approval boundary, got %#v", suggestions)
	}
	if len(suggestions.Suggestions) == 0 || suggestions.Suggestions[0].BlastRadiusEstimate == "" || suggestions.Suggestions[0].RollbackFeasibility == "" {
		t.Fatalf("expected bounded policy suggestion fields, got %#v", suggestions)
	}
	for _, suggestion := range suggestions.Suggestions {
		if !containsString(suggestion.ForbiddenActions, "automatic_profile_activation") && !containsString(suggestion.ForbiddenActions, "automatic_certification_release") {
			t.Fatalf("expected explicit forbidden automatic actions, got %#v", suggestion)
		}
	}
}

func TestPhase8AuthorityRoutingAndAIGuardrailsHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/governance/authority-routing", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 authority routing 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var routing phase8AuthorityRoutingResponse
	if err := json.NewDecoder(rec.Body).Decode(&routing); err != nil {
		t.Fatalf("decode phase8 authority routing: %v", err)
	}
	if routing.CurrentState != phase8AuthorityRoutingStateActive {
		t.Fatalf("expected active phase8 authority routing slice, got %#v", routing)
	}
	if routing.RoutingState != "formal_authority_routing_active" || routing.AlternateApproverPath == "" || routing.DeadlockResolutionPath == "" || routing.EmergencySuspensionPath == "" {
		t.Fatalf("expected bounded routing paths, got %#v", routing)
	}
	if !containsString(routing.ConstitutionalBoundaries, "anti_capture_rule") || !containsString(routing.ConstitutionalBoundaries, "no_single_hidden_model_override") {
		t.Fatalf("expected constitutional boundaries, got %#v", routing.ConstitutionalBoundaries)
	}
	if !containsString(routing.NonDelegableActions, "reactivate_challenged_artifact") {
		t.Fatalf("expected non-delegable action visibility, got %#v", routing.NonDelegableActions)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/governance/ai-guardrails", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 ai guardrails 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var guardrails phase8AIGuardrailsResponse
	if err := json.NewDecoder(rec.Body).Decode(&guardrails); err != nil {
		t.Fatalf("decode phase8 ai guardrails: %v", err)
	}
	if guardrails.CurrentState != phase8AIGuardrailsStateActive {
		t.Fatalf("expected active phase8 ai guardrail slice, got %#v", guardrails)
	}
	if guardrails.EscalationState != "human_escalation_threshold_visible" || guardrails.ConfidenceState != "confidence_floor_visible" {
		t.Fatalf("expected ai escalation and confidence states, got %#v", guardrails)
	}
	if !hasForbiddenRecommendationClass(guardrails.Guardrails, "legal_verdict") || !hasForbiddenRecommendationClass(guardrails.Guardrails, "non_delegable_action_execution") {
		t.Fatalf("expected prohibited recommendation classes, got %#v", guardrails.Guardrails)
	}
}

func TestPhase8ModelRiskHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/governance/model-risk", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 model risk 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase8ModelRiskResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode phase8 model risk: %v", err)
	}
	if response.CurrentState != phase8ModelRiskStateActive {
		t.Fatalf("expected active phase8 model risk slice, got %#v", response)
	}
	if response.ModelRiskState != "formal_model_risk_baseline_active" || response.DependencyRegistryState != "institutional_dependency_registry_visible" {
		t.Fatalf("expected bounded model risk states, got %#v", response)
	}
	if len(response.ModelRiskContracts) == 0 || response.ModelRiskContracts[0].RollbackPath == "" || response.ModelRiskContracts[0].ChallengerModelReview == "" {
		t.Fatalf("expected rollback and challenger review visibility, got %#v", response.ModelRiskContracts)
	}
	if !hasDependencyClass(response.Dependencies, "regulatory_framework") || !hasDependencyClass(response.Dependencies, "certification_scheme") {
		t.Fatalf("expected dependency registry coverage, got %#v", response.Dependencies)
	}
	if containsString(response.ModelRiskContracts[0].ForbiddenUse, formalcore.ClaimClassRegulatorSafeDisclosure) {
		t.Fatalf("unexpected claim class leak into forbidden-use semantics, got %#v", response.ModelRiskContracts[0].ForbiddenUse)
	}
}

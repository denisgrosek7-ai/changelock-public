package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	formalcore "github.com/denisgrosek/changelock/internal/formal"
)

func TestPhase8FinalSummaryHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/final-summary", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 final summary 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase8FinalSummaryResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode phase8 final summary: %v", err)
	}
	if response.CurrentState != phase8FinalizationStateReady {
		t.Fatalf("expected ready phase8 finalization pack, got %#v", response)
	}
	if response.Phase8CoreState != formalcore.Phase8StateActive {
		t.Fatalf("expected active phase8 core state, got %#v", response)
	}
	if response.LegalAndRegulatoryClaimReview.CurrentState != phase8FinalReviewStateActive ||
		response.UsePermissionAndStandardOfProofReview.CurrentState != phase8FinalReviewStateActive ||
		response.CustodyRedactionAndLegalHoldReview.CurrentState != phase8FinalReviewStateActive ||
		response.ComplianceCodificationReview.CurrentState != phase8FinalReviewStateActive ||
		response.CertificationWorkflowReview.CurrentState != phase8FinalReviewStateActive ||
		response.ConsensusGovernanceReview.CurrentState != phase8FinalReviewStateActive ||
		response.AIGuardrailReview.CurrentState != phase8FinalReviewStateActive ||
		response.InsurerEvidenceReview.CurrentState != phase8FinalReviewStateActive ||
		response.FormalAuthorityBoundaryReview.CurrentState != phase8FinalReviewStateActive {
		t.Fatalf("expected all final review sections active, got %#v", response)
	}
	if !containsString(response.DeferredInstitutionalScope, "insurer_integration_program") ||
		!containsString(response.DeferredInstitutionalScope, "wider_federated_governance") ||
		!containsString(response.DeferredInstitutionalScope, "advanced_institutional_disclosure_programs") {
		t.Fatalf("expected remaining deferred scope to stay visible, got %#v", response.DeferredInstitutionalScope)
	}
	if containsString(response.DeferredInstitutionalScope, "risk_quantification_baseline") ||
		containsString(response.DeferredInstitutionalScope, "insurance_facing_evidence_exports") {
		t.Fatalf("expected implemented 8C institutional surfaces to be excluded from remaining deferred scope, got %#v", response.DeferredInstitutionalScope)
	}
	if !containsString(response.FormalAuthorityBoundaryReview.DocRefs, "docs/formal-phase8-final.md") {
		t.Fatalf("expected final doc ref, got %#v", response.FormalAuthorityBoundaryReview.DocRefs)
	}
}

func TestPhase8FinalizationStateRequiresActiveCoreAndAllSections(t *testing.T) {
	activeSection := phase8FinalReviewSection{CurrentState: phase8FinalReviewStateActive}
	if got := phase8FinalizationState(formalcore.Phase8StateActive, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection); got != phase8FinalizationStateReady {
		t.Fatalf("expected ready finalization state, got %q", got)
	}
	if got := phase8FinalizationState(formalcore.Phase8StateSubstantial, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection); got != phase8FinalizationStateIncomplete {
		t.Fatalf("expected incomplete finalization for non-active core state, got %q", got)
	}
	partialSection := phase8FinalReviewSection{CurrentState: phase8FinalReviewStatePartial}
	if got := phase8FinalizationState(formalcore.Phase8StateActive, activeSection, partialSection, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection); got != phase8FinalizationStateSubstantial {
		t.Fatalf("expected substantial finalization when one section is partial, got %q", got)
	}
	incompleteSection := phase8FinalReviewSection{CurrentState: phase8FinalReviewStateIncomplete}
	if got := phase8FinalizationState(formalcore.Phase8StateActive, activeSection, incompleteSection, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection, activeSection); got != phase8FinalizationStateIncomplete {
		t.Fatalf("expected incomplete finalization when one section is incomplete, got %q", got)
	}
}

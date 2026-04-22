package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	ecosystemcore "github.com/denisgrosek/changelock/internal/ecosystem"
)

func TestPhase7FinalSummaryHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/ecosystem/phase7/final-summary", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase7 final summary 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response phase7FinalSummaryResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode final summary: %v", err)
	}
	if response.CurrentState != phase7FinalizationStateReady {
		t.Fatalf("expected ready phase7 finalization pack, got %#v", response)
	}
	if response.Phase7CoreState != ecosystemcore.Phase7StateActive {
		t.Fatalf("expected active phase7 core state, got %#v", response)
	}
	if response.DeveloperPresenceReview.CurrentState != phase7FinalReviewStateActive ||
		response.OSSBoundaryReview.CurrentState != phase7FinalReviewStateActive ||
		response.DistributionBoundaryReview.CurrentState != phase7FinalReviewStateActive ||
		response.ContractAlignment.CurrentState != phase7FinalReviewStateActive ||
		response.DocsAndProofs.CurrentState != phase7FinalReviewStateActive {
		t.Fatalf("expected all final review sections active, got %#v", response)
	}
	if !containsString(response.DeferredScope, "automated_pr_discipline") || !containsString(response.DeferredScope, "integrity_as_a_service_package") {
		t.Fatalf("expected deferred scope to remain visible, got %#v", response.DeferredScope)
	}
	if !containsString(response.DocsAndProofs.DocRefs, "docs/ecosystem-phase7-final.md") {
		t.Fatalf("expected final doc ref, got %#v", response.DocsAndProofs.DocRefs)
	}
}

func TestPhase7FinalizationStateRequiresActiveCoreAndAllSections(t *testing.T) {
	activeSection := phase7FinalReviewSection{CurrentState: phase7FinalReviewStateActive}
	if got := phase7FinalizationState(ecosystemcore.Phase7StateActive, activeSection, activeSection, activeSection, activeSection, activeSection); got != phase7FinalizationStateReady {
		t.Fatalf("expected ready finalization state, got %q", got)
	}
	if got := phase7FinalizationState(ecosystemcore.Phase7StateSubstantial, activeSection, activeSection, activeSection, activeSection, activeSection); got != phase7FinalizationStateIncomplete {
		t.Fatalf("expected incomplete finalization for non-active core state, got %q", got)
	}
	partialSection := phase7FinalReviewSection{CurrentState: phase7FinalReviewStatePartial}
	if got := phase7FinalizationState(ecosystemcore.Phase7StateActive, activeSection, partialSection, activeSection, activeSection, activeSection); got != phase7FinalizationStateSubstantial {
		t.Fatalf("expected substantial finalization when one section is partial, got %q", got)
	}
	incompleteSection := phase7FinalReviewSection{CurrentState: phase7FinalReviewStateIncomplete}
	if got := phase7FinalizationState(ecosystemcore.Phase7StateActive, activeSection, incompleteSection, activeSection, activeSection, activeSection); got != phase7FinalizationStateIncomplete {
		t.Fatalf("expected incomplete finalization when one section is incomplete, got %q", got)
	}
}

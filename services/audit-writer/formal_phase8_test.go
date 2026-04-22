package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	formalcore "github.com/denisgrosek/changelock/internal/formal"
)

func TestPhase8EntryGateAndContractsHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/entry-gate", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 entry gate 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var entry phase8EntryGateResponse
	if err := json.NewDecoder(rec.Body).Decode(&entry); err != nil {
		t.Fatalf("decode phase8 entry gate: %v", err)
	}
	if entry.CurrentState != formalcore.EntryGateStateReady {
		t.Fatalf("expected ready phase8 entry gate, got %#v", entry)
	}
	if !containsString(entry.DeferredScope, "risk_quantification_baseline") {
		t.Fatalf("expected deferred institutional scope to remain visible, got %#v", entry.DeferredScope)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/contracts", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 contracts 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var contracts phase8ContractsResponse
	if err := json.NewDecoder(rec.Body).Decode(&contracts); err != nil {
		t.Fatalf("decode phase8 contracts: %v", err)
	}
	if contracts.CurrentState != formalcore.FoundationStateActive {
		t.Fatalf("expected active phase8 contract foundation, got %#v", contracts)
	}
	if contracts.Coverage.ClaimClasses == 0 || contracts.Coverage.AuthorityControls == 0 {
		t.Fatalf("expected populated coverage, got %#v", contracts.Coverage)
	}
}

func TestPhase8FormalDisciplineAndComplianceHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/formal-discipline", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 formal discipline 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var formal phase8FormalDisciplineResponse
	if err := json.NewDecoder(rec.Body).Decode(&formal); err != nil {
		t.Fatalf("decode phase8 formal discipline: %v", err)
	}
	if formal.CurrentState != formalcore.FormalDisciplineStateActive {
		t.Fatalf("expected active phase8 formal discipline, got %#v", formal)
	}
	if formal.ClaimTaxonomyState != "formal_claim_taxonomy_active" || formal.EvidenceCustodyState != "formal_evidence_custody_active" {
		t.Fatalf("expected active formal states, got %#v", formal)
	}
	if len(formal.ClaimClasses) == 0 || len(formal.UsePermissionRules) == 0 || len(formal.StandardOfProofClasses) == 0 {
		t.Fatalf("expected populated formal contracts, got %#v", formal)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/compliance-codification", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 compliance codification 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var compliance phase8ComplianceCodificationResponse
	if err := json.NewDecoder(rec.Body).Decode(&compliance); err != nil {
		t.Fatalf("decode phase8 compliance codification: %v", err)
	}
	if compliance.CurrentState != formalcore.ComplianceCodificationStateActive {
		t.Fatalf("expected active phase8 compliance codification, got %#v", compliance)
	}
	if compliance.PolicyAsLawState != "policy_as_law_profile_active" || compliance.VerifierSurfaceState != "regulator_and_certifier_safe_surface_active" {
		t.Fatalf("expected active compliance codification states, got %#v", compliance)
	}
}

func TestPhase8GovernedAutonomyAndProofsHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/governed-autonomy", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 governed autonomy 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var governance phase8GovernedAutonomyResponse
	if err := json.NewDecoder(rec.Body).Decode(&governance); err != nil {
		t.Fatalf("decode phase8 governed autonomy: %v", err)
	}
	if governance.CurrentState != formalcore.GovernedAutonomyStateActive {
		t.Fatalf("expected active phase8 governed autonomy, got %#v", governance)
	}
	if !containsString(governance.DeferredInstitutionalScope, "insurance_facing_evidence_exports") {
		t.Fatalf("expected institutional expansion deferment, got %#v", governance.DeferredInstitutionalScope)
	}
	if governance.AuthorityControlState != "non_delegable_authority_controls_active" || governance.AIGuardrailState != "formal_ai_guardrails_active" {
		t.Fatalf("expected active authority and ai guardrails, got %#v", governance)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/proofs", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs phase8ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode phase8 proofs: %v", err)
	}
	if proofs.CurrentState != formalcore.Phase8StateActive {
		t.Fatalf("expected active phase8 proofs gate, got %#v", proofs)
	}
	if proofs.CoverageScope != phase8CoverageScopeCorePass {
		t.Fatalf("expected core-pass coverage scope, got %#v", proofs)
	}
	if !containsString(proofs.DeferredInstitutionalScope, "actuarial_benchmark_discipline") {
		t.Fatalf("expected deferred institutional expansion to stay visible, got %#v", proofs.DeferredInstitutionalScope)
	}
}

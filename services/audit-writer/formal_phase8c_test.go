package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	formalcore "github.com/denisgrosek/changelock/internal/formal"
)

func TestPhase8RiskQuantificationAndInsuranceExportsHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/institutional/risk-quantification", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 risk quantification 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var quant phase8RiskQuantificationResponse
	if err := json.NewDecoder(rec.Body).Decode(&quant); err != nil {
		t.Fatalf("decode phase8 risk quantification: %v", err)
	}
	if quant.CurrentState != phase8RiskQuantificationStateActive {
		t.Fatalf("expected active phase8 risk quantification slice, got %#v", quant)
	}
	if quant.CalibrationClass != "integration_ready_not_pricing_promise" || quant.PremiumModelBoundary != "never_automatic_pricing_promise" {
		t.Fatalf("expected bounded calibration and premium boundary, got %#v", quant)
	}
	if !containsString(quant.ForbiddenUseNotes, "Do not use as an automatic pricing engine output.") {
		t.Fatalf("expected automatic pricing prohibition, got %#v", quant.ForbiddenUseNotes)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/institutional/insurance-exports", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 insurance exports 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var exports phase8InsuranceExportsResponse
	if err := json.NewDecoder(rec.Body).Decode(&exports); err != nil {
		t.Fatalf("decode phase8 insurance exports: %v", err)
	}
	if exports.CurrentState != phase8InsuranceExportsStateActive {
		t.Fatalf("expected active phase8 insurance export slice, got %#v", exports)
	}
	if exports.DisclosureScopeState != "insurer_scoped_export_only" || exports.ReleaseLifecycleState != "release_and_withdrawal_lifecycle_visible" {
		t.Fatalf("expected bounded insurer disclosure states, got %#v", exports)
	}
	if len(exports.Exports) == 0 {
		t.Fatalf("expected insurer export items, got %#v", exports)
	}
	item := exports.Exports[0]
	if item.Audience != formalcore.AudienceInsurer || item.ClaimClass != formalcore.ClaimClassInsurerFacingRiskInput {
		t.Fatalf("expected insurer audience and claim class, got %#v", item)
	}
	if item.CanCitePublicly {
		t.Fatalf("expected insurer export to remain non-public, got %#v", item)
	}
	if !item.ReleaseApprovalRequired || !containsString(item.DisclosureScopeProfile, "insurer_scoped") {
		t.Fatalf("expected release approval and insurer scope, got %#v", item)
	}
	if item.AdverseDecisionExplanation.ChallengePath == "" || item.AdverseDecisionExplanation.SufficiencyClass == "" {
		t.Fatalf("expected adverse decision explanation support, got %#v", item.AdverseDecisionExplanation)
	}
}

func TestPhase8IncidentAttributionAndActuarialBenchmarksHandlers(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/institutional/incident-attribution", nil)
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 incident attribution 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var attribution phase8IncidentAttributionResponse
	if err := json.NewDecoder(rec.Body).Decode(&attribution); err != nil {
		t.Fatalf("decode phase8 incident attribution: %v", err)
	}
	if attribution.CurrentState != phase8IncidentAttributionStateActive {
		t.Fatalf("expected active phase8 incident attribution slice, got %#v", attribution)
	}
	if attribution.NonLegalConclusionState != "non_legal_conclusion_marker_active" || attribution.AmbiguityHandlingState != "unresolved_ambiguity_visible" {
		t.Fatalf("expected non-legal and ambiguity states, got %#v", attribution)
	}
	if len(attribution.Items) == 0 {
		t.Fatalf("expected attribution support items, got %#v", attribution)
	}
	for _, item := range attribution.Items {
		if item.NonLegalConclusionMarker != "support_only_not_legal_conclusion" || item.UnresolvedAmbiguityNote == "" {
			t.Fatalf("expected bounded non-legal attribution support, got %#v", item)
		}
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/formal/phase8/institutional/actuarial-benchmarks", nil)
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected phase8 actuarial benchmarks 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var benchmarks phase8ActuarialBenchmarksResponse
	if err := json.NewDecoder(rec.Body).Decode(&benchmarks); err != nil {
		t.Fatalf("decode phase8 actuarial benchmarks: %v", err)
	}
	if benchmarks.CurrentState != phase8ActuarialBenchmarksStateActive {
		t.Fatalf("expected active phase8 actuarial benchmark slice, got %#v", benchmarks)
	}
	if benchmarks.PrivacyBoundaryState != "aggregate_only_and_reidentification_guarded" || benchmarks.PublicationBoundaryState != "withdrawal_trigger_visible" {
		t.Fatalf("expected bounded actuarial benchmark states, got %#v", benchmarks)
	}
	if len(benchmarks.Items) == 0 {
		t.Fatalf("expected actuarial benchmark items, got %#v", benchmarks)
	}
	benchmark := benchmarks.Items[0]
	if benchmark.MinimumCohortSize < 50 || benchmark.AggregationScope != "aggregate_only_cross_tenant_safe_band" {
		t.Fatalf("expected aggregate-only benchmark discipline, got %#v", benchmark)
	}
	if !containsString(benchmark.ForbiddenUse, "tenant_level_pricing") || !containsString(benchmark.ForbiddenUse, "raw_subject_disclosure") {
		t.Fatalf("expected forbidden tenant-level pricing and raw disclosure, got %#v", benchmark.ForbiddenUse)
	}
}

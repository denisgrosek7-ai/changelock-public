package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestIntelligenceCalibrationValAFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/intelligence/calibration/vala/reachability-aggregation?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValAAggregationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode aggregation: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValAAggregationStateActive || response.Model.AggregationID == "" {
					t.Fatalf("unexpected aggregation response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vala/exploitability-calibration?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValAExploitabilityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode exploitability: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValAExploitabilityStateActive || !response.Model.AdvisoryOnly {
					t.Fatalf("unexpected exploitability response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vala/downgrade-escalation?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValADecisionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode decision: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValADecisionStateActive || response.Model.RollbackRef == "" {
					t.Fatalf("unexpected decision response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vala/cavi-tuning?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValACAVIResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode cavi: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValACAVIStateActive || response.Model.CAVIProfileID == "" {
					t.Fatalf("unexpected cavi response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vala/vex-candidates?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValAVEXResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode vex candidate: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValAVEXCandidateStateActive || response.Model.PublicationAllowed {
					t.Fatalf("unexpected vex response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vala/vex-sufficiency?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValASufficiencyResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode sufficiency: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValAVEXSufficiencyStateActive || response.Model.SufficiencyCheckID == "" {
					t.Fatalf("unexpected sufficiency response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vala/explanations?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValAExplanationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode explanations: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValAExplanationStateActive || !response.Model.DistinguishesNotEvidencedFromSafe {
					t.Fatalf("unexpected explanation response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vala/confidence-outcomes?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValAOutcomeResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode outcomes: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValAConfidenceOutcomeStateActive || response.Model.OutcomeID == "" {
					t.Fatalf("unexpected outcome response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vala/publication-guardrail?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValAGuardrailResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode guardrail: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValAPublicationGuardrailStateActive || !response.Model.FinalClaimBlocked {
					t.Fatalf("unexpected guardrail response %#v", response)
				}
			},
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, tc.path, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200 for %s, got %d: %s", tc.path, rec.Code, rec.Body.String())
		}
		tc.decode(t, rec)
	}
}

func TestIntelligenceCalibrationValAProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/intelligence/calibration/vala/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response intelligenceCalibrationValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.IntelligenceCalibrationValAStateActive {
		t.Fatalf("expected active proofs state, got %#v", response)
	}
	if response.Val0DependencyState != operability.IntelligenceCalibrationVal0StateActive || response.Val0FoundationState != operability.IntelligenceCalibrationVal0StateActive {
		t.Fatalf("expected active Val 0 dependency states, got %#v", response)
	}
	if response.ValAState != operability.IntelligenceCalibrationValAStateActive {
		t.Fatalf("expected active Val A state, got %#v", response)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StateNotComplete {
		t.Fatalf("expected point 5 to remain not complete, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection-only disclaimer, got %#v", response)
	}
	if len(response.WhyPoint5NotPass) == 0 || len(response.SurfaceRefs) < 10 || len(response.EvidenceRefs) < 8 || len(response.Limitations) == 0 {
		t.Fatalf("expected non-empty why_not_pass, refs, and limitations, got %#v", response)
	}
}

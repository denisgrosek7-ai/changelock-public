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

func TestIntelligenceCalibrationValBFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/intelligence/calibration/valb/behavioral-baseline?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValBBaselineResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode behavioral baseline: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValBBehavioralBaselineStateActive || response.Model.BaselineID == "" {
					t.Fatalf("unexpected behavioral baseline response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valb/learning-mode-runtime?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValBLearningResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode learning runtime: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValBLearningRuntimeStateActive || !response.Model.OutputReviewRequired {
					t.Fatalf("unexpected learning runtime response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valb/anomaly-thresholds?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValBThresholdResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode threshold calibration: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValBThresholdStateActive || response.Model.BaselineRef == "" {
					t.Fatalf("unexpected threshold response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valb/drift-sensitivity?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValBDriftResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode drift sensitivity: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValBDriftStateActive || response.Model.DriftProfileID == "" {
					t.Fatalf("unexpected drift response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valb/criticality-weighting?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValBWeightingResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode criticality weighting: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValBWeightingStateActive || response.Model.WeightingProfileID == "" {
					t.Fatalf("unexpected weighting response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valb/baseline-freshness?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValBFreshnessResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode baseline freshness: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValBBaselineFreshnessStateActive || response.Model.BaselineRef == "" {
					t.Fatalf("unexpected freshness response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valb/baseline-adoption?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValBAdoptionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode baseline adoption: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValBBaselineAdoptionStateActive || response.Model.RollbackRef == "" {
					t.Fatalf("unexpected adoption response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valb/explanations?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValBExplanationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode behavioral explanation: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValBExplanationStateActive || !response.Model.ReviewerRequired {
					t.Fatalf("unexpected explanation response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valb/safety-guardrails?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValBGuardrailResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode behavioral guardrail: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValBGuardrailStateActive || !response.Model.AutoSuppressionBlocked {
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

func TestIntelligenceCalibrationValBProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/intelligence/calibration/valb/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response intelligenceCalibrationValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.IntelligenceCalibrationValBStateActive {
		t.Fatalf("expected active Val B proofs state, got %#v", response)
	}
	if response.Val0DependencyState != operability.IntelligenceCalibrationVal0StateActive || response.Val0FoundationState != operability.IntelligenceCalibrationVal0StateActive {
		t.Fatalf("expected active Val 0 dependency states, got %#v", response)
	}
	if response.ValADependencyState != operability.IntelligenceCalibrationValAStateActive || response.ValAReachabilityVEXState != operability.IntelligenceCalibrationValAStateActive {
		t.Fatalf("expected active Val A dependency states, got %#v", response)
	}
	if response.ValBState != operability.IntelligenceCalibrationValBStateActive {
		t.Fatalf("expected active Val B state, got %#v", response)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StateNotComplete {
		t.Fatalf("expected point 5 to remain not complete, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection-only disclaimer, got %#v", response)
	}
	if len(response.WhyPoint5NotPass) == 0 || len(response.SurfaceRefs) < 10 || len(response.EvidenceRefs) < 9 || len(response.Limitations) == 0 {
		t.Fatalf("expected non-empty why_not_pass, refs, and limitations, got %#v", response)
	}
}

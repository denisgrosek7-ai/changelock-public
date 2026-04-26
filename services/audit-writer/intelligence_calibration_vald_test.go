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

func TestIntelligenceCalibrationValDFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/intelligence/calibration/vald/simulation-harness?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDSimulationHarnessResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode simulation harness: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDSimulationHarnessStateActive || response.Model.HarnessID == "" {
					t.Fatalf("unexpected simulation harness response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vald/scenario-library?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDScenarioLibraryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode scenario library: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDScenarioLibraryStateActive || response.Model.LibraryID == "" {
					t.Fatalf("unexpected scenario library response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vald/missed-detection-analysis?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDMissedDetectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode missed detection analysis: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDMissedDetectionStateActive || response.Model.AnalysisID == "" {
					t.Fatalf("unexpected missed detection response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vald/fp-fn-balance?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDFPFNBalanceResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode fp-fn balance: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDFPFNBalanceStateActive || response.Model.ReviewID == "" {
					t.Fatalf("unexpected fp-fn balance response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vald/confidence-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDConfidenceReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode confidence review: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDConfidenceReviewStateActive || response.Model.ReviewID == "" {
					t.Fatalf("unexpected confidence review response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vald/vex-quality?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDVEXQualityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode vex quality: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDVEXQualityStateActive || response.Model.ReviewID == "" {
					t.Fatalf("unexpected vex quality response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vald/reachability-quality?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDReachabilityQualityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode reachability quality: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDReachabilityQualityStateActive || response.Model.ReviewID == "" {
					t.Fatalf("unexpected reachability quality response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vald/behavioral-quality?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDBehavioralQualityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode behavioral quality: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDBehavioralQualityStateActive || response.Model.ReviewID == "" {
					t.Fatalf("unexpected behavioral quality response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vald/federated-quality?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDFederatedQualityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode federated quality: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDFederatedQualityStateActive || response.Model.ReviewID == "" {
					t.Fatalf("unexpected federated quality response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vald/simulation-coverage?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDSimulationCoverageResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode simulation coverage: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDSimulationCoverageStateActive || response.Model.CoverageID == "" {
					t.Fatalf("unexpected simulation coverage response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/vald/quality-scoreboard?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValDQualityScoreboardResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode quality scoreboard: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValDQualityScoreboardStateActive || response.Model.ScoreboardID == "" {
					t.Fatalf("unexpected quality scoreboard response %#v", response)
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

func TestIntelligenceCalibrationValDProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/intelligence/calibration/vald/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response intelligenceCalibrationValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val D proofs: %v", err)
	}
	if response.CurrentState != operability.IntelligenceCalibrationValDStateActive {
		t.Fatalf("expected active Val D proofs state, got %#v", response)
	}
	if response.ValDState != operability.IntelligenceCalibrationValDStateActive {
		t.Fatalf("expected active Val D state in proofs, got %#v", response)
	}
	if response.Val0DependencyState != operability.IntelligenceCalibrationVal0StateActive || response.ValADependencyState != operability.IntelligenceCalibrationValAStateActive || response.ValBDependencyState != operability.IntelligenceCalibrationValBStateActive || response.ValCDependencyState != operability.IntelligenceCalibrationValCStateActive {
		t.Fatalf("expected active dependency states in Val D proofs, got %#v", response)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StateNotComplete {
		t.Fatalf("expected point 5 to remain not_complete, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected Val D proofs disclaimer to stay projection-only, got %#v", response)
	}
}

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

func TestIntelligenceCalibrationVal0FoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/intelligence/calibration/val0/datasets?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0DatasetResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode datasets: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0DatasetStateActive || response.Model.DatasetVersion == "" {
					t.Fatalf("unexpected dataset response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/confidence-model?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0ConfidenceResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode confidence: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0ConfidenceStateActive || !response.Model.AdvisoryOnly {
					t.Fatalf("unexpected confidence response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/output-lifecycle?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0LifecycleResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode lifecycle: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0LifecycleStateActive || len(response.Model.Items) != 8 {
					t.Fatalf("unexpected lifecycle response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/reachability-taxonomy?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0ReachabilityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode reachability: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0ReachabilityStateActive || response.Model.Example.ReachabilityClass == "" {
					t.Fatalf("unexpected reachability response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/vex-candidates?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0VEXResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode vex: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0VEXStateActive || response.Model.Example.CandidateID == "" {
					t.Fatalf("unexpected vex response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/feedback?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0FeedbackResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode feedback: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0FeedbackStateActive || len(response.Model.Items) != 8 {
					t.Fatalf("unexpected feedback response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/learning-mode?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0LearningModeResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode learning mode: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0LearningModeStateActive || !response.Model.OutputReviewRequired {
					t.Fatalf("unexpected learning mode response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/suppression-safety?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0SuppressionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode suppression: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0SuppressionStateActive || response.Model.RollbackRef == "" {
					t.Fatalf("unexpected suppression response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/federated-boundary?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0FederatedResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode federated boundary: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0FederatedBoundaryStateActive || response.Model.PropagationAllowed {
					t.Fatalf("unexpected federated response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/provenance?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0ProvenanceResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode provenance: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0ProvenanceStateActive || response.Model.ApprovalState != operability.IntelligenceCalibrationApprovalApproved {
					t.Fatalf("unexpected provenance response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/freshness-expiry?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0FreshnessResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode freshness: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0FreshnessStateActive || len(response.Model.Items) != 8 {
					t.Fatalf("unexpected freshness response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/metrics?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0MetricsResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode metrics: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0MetricsStateActive || len(response.Model.Items) != 11 {
					t.Fatalf("unexpected metrics response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/fp-fn-discipline?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0FPFNResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode fp/fn: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0FPFNStateActive || !response.Model.CriticalLowSignalReviewRequired {
					t.Fatalf("unexpected fp/fn response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/val0/rollback?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationVal0RollbackResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode rollback: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationVal0RollbackStateActive || !response.Model.RollbackAvailable {
					t.Fatalf("unexpected rollback response %#v", response)
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

func TestIntelligenceCalibrationVal0ProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/intelligence/calibration/val0/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response intelligenceCalibrationVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if response.CurrentState != operability.IntelligenceCalibrationVal0StateActive {
		t.Fatalf("expected active proofs state, got %#v", response)
	}
	if response.Val0State != operability.IntelligenceCalibrationVal0StateActive {
		t.Fatalf("expected active Val 0 foundation state, got %#v", response)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StateNotComplete {
		t.Fatalf("expected point 5 to remain not complete, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected projection-only disclaimer, got %#v", response)
	}
	if len(response.WhyPoint5NotPass) == 0 || len(response.SurfaceRefs) < 15 || len(response.EvidenceRefs) < 8 || len(response.Limitations) == 0 {
		t.Fatalf("expected non-empty why_not_pass, refs, and limitations, got %#v", response)
	}
}

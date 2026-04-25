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

func TestIntelligenceCalibrationValCFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/intelligence/calibration/valc/feedback-intake?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCFeedbackIntakeResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode feedback intake: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCFeedbackIntakeStateActive || response.Model.FeedbackID == "" {
					t.Fatalf("unexpected feedback intake response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valc/feedback-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCReviewCockpitResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode feedback review: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCReviewCockpitStateActive || response.Model.ReviewQueueID == "" {
					t.Fatalf("unexpected feedback review response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valc/tuning-proposals?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCTuningProposalResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode tuning proposals: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCTuningProposalStateActive || response.Model.ProposalID == "" {
					t.Fatalf("unexpected tuning proposal response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valc/suppression-safety?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCSuppressionSafetyResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode suppression safety: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCSuppressionSafetyStateActive || response.Model.SuppressionCandidateID == "" {
					t.Fatalf("unexpected suppression safety response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valc/suppression-rollback?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCSuppressionRollbackResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode suppression rollback: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCSuppressionRollbackStateActive || response.Model.RollbackID == "" {
					t.Fatalf("unexpected suppression rollback response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valc/local-change-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCLocalChangeReviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode local change review: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCLocalChangeReviewStateActive || response.Model.ChangeReviewID == "" {
					t.Fatalf("unexpected local change review response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valc/federated-weighting?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCFederatedWeightingResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode federated weighting: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCFederatedWeightingStateActive || response.Model.FederatedSignalID == "" {
					t.Fatalf("unexpected federated weighting response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valc/similarity-gating?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCSimilarityGatingResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode similarity gating: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCSimilarityGatingStateActive || response.Model.SimilarityProfileID == "" {
					t.Fatalf("unexpected similarity gating response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valc/local-override?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCLocalOverrideResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode local override: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCLocalOverrideStateActive || response.Model.OverridePolicyID == "" {
					t.Fatalf("unexpected local override response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valc/propagation-policy?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCPropagationPolicyResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode propagation policy: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCPropagationPolicyStateActive || response.Model.PropagationPolicyID == "" {
					t.Fatalf("unexpected propagation policy response %#v", response)
				}
			},
		},
		{
			path: "/v1/intelligence/calibration/valc/explanations?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response intelligenceCalibrationValCExplanationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode explanations: %v", err)
				}
				if response.CurrentState != operability.IntelligenceCalibrationValCExplanationStateActive || !response.Model.ReviewerRequired {
					t.Fatalf("unexpected explanation response %#v", response)
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

func TestIntelligenceCalibrationValCProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/intelligence/calibration/valc/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response intelligenceCalibrationValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val C proofs: %v", err)
	}
	if response.CurrentState != operability.IntelligenceCalibrationValCStateActive {
		t.Fatalf("expected active Val C proofs state, got %#v", response)
	}
	if response.ValCState != operability.IntelligenceCalibrationValCStateActive {
		t.Fatalf("expected active Val C state in proofs, got %#v", response)
	}
	if response.Val0DependencyState != operability.IntelligenceCalibrationVal0StateActive || response.ValADependencyState != operability.IntelligenceCalibrationValAStateActive || response.ValBDependencyState != operability.IntelligenceCalibrationValBStateActive {
		t.Fatalf("expected active dependency states in Val C proofs, got %#v", response)
	}
	if response.Point5State != operability.IntelligenceCalibrationPoint5StateNotComplete {
		t.Fatalf("expected point 5 to remain not_complete, got %#v", response)
	}
	if !strings.Contains(response.ProjectionDisclaimer, "projection_only") || !strings.Contains(response.ProjectionDisclaimer, "not_canonical_truth") {
		t.Fatalf("expected Val C proofs disclaimer to stay projection-only, got %#v", response)
	}
}

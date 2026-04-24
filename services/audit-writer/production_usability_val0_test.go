package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestProductionUsabilityVal0FoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/production/usability-operability-recovery/val0/config-integrity?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityVal0ConfigIntegrityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode config integrity: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityVal0ConfigIntegrityStateActive || response.Model.CurrentValidationResult != operability.ProductionUsabilityValidationValid {
					t.Fatalf("unexpected config integrity response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/val0/explainability-contract?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityVal0ExplainabilityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode explainability contract: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityVal0ExplainabilityStateActive || len(response.Model.RequiredFields) < 13 {
					t.Fatalf("unexpected explainability response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/val0/status-model?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityVal0StatusModelResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode status model: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityVal0StatusModelStateActive || len(response.Model.Items) != 6 {
					t.Fatalf("unexpected status model %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/val0/operation-contracts?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityVal0OperationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode operation contracts: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityVal0OperationContractStateActive || len(response.Model.Items) != 5 {
					t.Fatalf("unexpected operation contract %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/val0/decision-quality?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityVal0DecisionQualityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode decision quality: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityVal0DecisionQualityStateActive || !response.Model.AdvisoryOnly {
					t.Fatalf("unexpected decision quality %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/val0/notification-taxonomy?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityVal0NotificationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode notification taxonomy: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityVal0NotificationStateActive || len(response.Model.AcknowledgementStates) != 5 {
					t.Fatalf("unexpected notification taxonomy %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/val0/permission-redaction?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityVal0PermissionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode permission redaction: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityVal0PermissionRedactionStateActive || len(response.Model.Items) != 5 {
					t.Fatalf("unexpected permission redaction %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/val0/recovery-contract?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityVal0RecoveryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode recovery contract: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityVal0RecoveryStateActive || len(response.Model.RemediationClasses) != 8 {
					t.Fatalf("unexpected recovery contract %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/val0/action-modes?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityVal0ActionModesResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode action modes: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityVal0ActionModeStateActive || len(response.Model.Items) != 7 {
					t.Fatalf("unexpected action modes %#v", response)
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

func TestProductionUsabilityVal0ProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/val0/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val 0 proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val 0 proofs: %v", err)
	}
	if response.CurrentState != operability.ProductionUsabilityVal0StateActive {
		t.Fatalf("expected active Val 0 proofs state, got %#v", response)
	}
	if response.Val0State != operability.ProductionUsabilityVal0StateActive {
		t.Fatalf("expected active Val 0 foundation state, got %#v", response)
	}
	if response.Point4State != operability.ProductionUsabilityPoint4StateNotComplete {
		t.Fatalf("expected point 4 to remain not complete, got %#v", response)
	}
	if len(response.WhyPoint4NotPass) == 0 || len(response.SurfaceRefs) < 10 || len(response.EvidenceRefs) < 6 {
		t.Fatalf("expected non-empty why_not_pass, surface refs, and evidence refs, got %#v", response)
	}
}

func TestProductionUsabilityVal0ProofsStayInactiveWithoutPoint3(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/val0/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val 0 proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode inactive Val 0 proofs: %v", err)
	}
	if response.CurrentState == operability.ProductionUsabilityVal0StateActive {
		t.Fatalf("expected inactive Val 0 proofs without point 3 dependency, got %#v", response)
	}
	if response.Point3DependencyState == "" || response.Point3DependencyState == operability.ProductionUsabilityVal0StateActive {
		t.Fatalf("expected explicit non-active point 3 dependency, got %#v", response)
	}
}

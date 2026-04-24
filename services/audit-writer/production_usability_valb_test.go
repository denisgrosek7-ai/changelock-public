package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestProductionUsabilityValBFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/production/usability-operability-recovery/valb/ui-data-resilience?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValBUIDataResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode UI data resilience: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValBUIDataResilienceStateActive || len(response.Model.Items) != 6 {
					t.Fatalf("unexpected UI data resilience response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valb/windowing?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValBWindowingResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode windowing: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValBWindowingStateActive || !response.Model.PartialResult {
					t.Fatalf("unexpected windowing response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valb/result-semantics?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValBResultSemanticsResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode result semantics: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValBResultSemanticsStateActive || len(response.Model.Items) != 6 {
					t.Fatalf("unexpected result semantics response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valb/command-center-tasks?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValBCommandCenterResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode command center tasks: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValBCommandCenterStateActive || len(response.Model.Items) != 3 {
					t.Fatalf("unexpected command center response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valb/noise-budget?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValBNoiseBudgetResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode noise budget: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValBNoiseBudgetStateActive || len(response.Model.Items) != 3 {
					t.Fatalf("unexpected noise budget response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valb/api-protection?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValBAPIProtectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode API protection: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValBAPIProtectionStateActive || len(response.Model.Items) != 7 {
					t.Fatalf("unexpected API protection response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valb/cli-resilience?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValBCLIResilienceResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode CLI resilience: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValBCLIResilienceStateActive || len(response.Model.Items) != 7 {
					t.Fatalf("unexpected CLI resilience response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valb/scale-envelope?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValBScaleEnvelopeResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode scale envelope: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValBScaleEnvelopeStateActive || response.Model.ClaimsLatencyGuarantee {
					t.Fatalf("unexpected scale envelope response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valb/action-mode-enforcement?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValBActionModeResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode action mode enforcement: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValBActionModeEnforcementStateActive || len(response.Model.Items) != 7 {
					t.Fatalf("unexpected action mode enforcement response %#v", response)
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

func TestProductionUsabilityValBProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/valb/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val B proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val B proofs: %v", err)
	}
	if response.CurrentState != operability.ProductionUsabilityValBStateActive {
		t.Fatalf("expected active Val B proofs state, got %#v", response)
	}
	if response.Val0FoundationState != operability.ProductionUsabilityVal0StateActive || response.ValACoreState != operability.ProductionUsabilityValAStateActive {
		t.Fatalf("expected active Val 0 and Val A dependencies, got %#v", response)
	}
	if response.ValBState != operability.ProductionUsabilityValBStateActive {
		t.Fatalf("expected active Val B resilience state, got %#v", response)
	}
	if response.Point4State != operability.ProductionUsabilityPoint4StateNotComplete {
		t.Fatalf("expected point 4 to remain not complete, got %#v", response)
	}
	if len(response.WhyPoint4NotPass) == 0 || len(response.SurfaceRefs) < 10 || len(response.EvidenceRefs) < 8 {
		t.Fatalf("expected why_not_pass, surface refs, and evidence refs, got %#v", response)
	}
}

func TestProductionUsabilityValBProofsStayInactiveWithoutDependencies(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/valb/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val B proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode inactive Val B proofs: %v", err)
	}
	if response.CurrentState == operability.ProductionUsabilityValBStateActive {
		t.Fatalf("expected inactive Val B proofs without dependencies, got %#v", response)
	}
	if response.Val0DependencyState == "" || response.ValADependencyState == "" {
		t.Fatalf("expected explicit dependency states, got %#v", response)
	}
}

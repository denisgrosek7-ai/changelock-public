package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestProductionUsabilityValAFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/production/usability-operability-recovery/vala/config-factory?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValAConfigFactoryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode config factory: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValAConfigFactoryStateActive || response.Model.CurrentValidationResult != operability.ProductionUsabilityValidationValid {
					t.Fatalf("unexpected config factory response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vala/bootstrap-validation?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValABootstrapResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode bootstrap validation: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValABootstrapValidationStateActive || response.Model.BootstrapDisposition != operability.ProductionUsabilityBootstrapAllowed {
					t.Fatalf("unexpected bootstrap validation response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vala/policy-schema?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValAPolicySchemaResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode policy schema: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValAPolicySchemaStateActive || response.Model.CurrentValidationResult != operability.ProductionUsabilityValidationValid {
					t.Fatalf("unexpected policy schema response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vala/effective-config?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValAEffectiveConfigResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode effective config: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValAEffectiveConfigStateActive || len(response.Model.RedactedFields) == 0 {
					t.Fatalf("unexpected effective config response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vala/rejections?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValARejectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode rejections: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValARejectionLayerStateActive || len(response.Model.RequiredFields) < 13 {
					t.Fatalf("unexpected rejections response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vala/policy-dry-run?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValADryRunResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode policy dry-run: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValADryRunStateActive || response.Model.MutatesCanonicalState {
					t.Fatalf("unexpected dry-run response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vala/explain?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValAExplainResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode explain outputs: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValAExplainStateActive || len(response.Model.Items) != 5 {
					t.Fatalf("unexpected explain output response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vala/recovery-guidance?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValARecoveryGuidanceResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode recovery guidance: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValARecoveryGuidanceStateActive || len(response.Model.Items) != 9 {
					t.Fatalf("unexpected recovery guidance response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vala/first-run-bootstrap?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValAFirstRunResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode first-run bootstrap: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValAFirstRunStateActive || response.Model.AutoEnablesProduction {
					t.Fatalf("unexpected first-run bootstrap response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/vala/upgrade-impact-preview?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValAUpgradePreviewResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode upgrade impact preview: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValAUpgradePreviewStateActive || !response.Model.RollbackAppearsPossible {
					t.Fatalf("unexpected upgrade impact preview response %#v", response)
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

func TestProductionUsabilityValAProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/vala/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val A proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val A proofs: %v", err)
	}
	if response.CurrentState != operability.ProductionUsabilityValAStateActive {
		t.Fatalf("expected active Val A proofs state, got %#v", response)
	}
	if response.Val0FoundationState != operability.ProductionUsabilityVal0StateActive {
		t.Fatalf("expected active Val 0 dependency, got %#v", response)
	}
	if response.ValAState != operability.ProductionUsabilityValAStateActive {
		t.Fatalf("expected active Val A core state, got %#v", response)
	}
	if response.Point4State != operability.ProductionUsabilityPoint4StateNotComplete {
		t.Fatalf("expected point 4 to remain not complete, got %#v", response)
	}
	if len(response.WhyPoint4NotPass) == 0 || len(response.SurfaceRefs) < 11 || len(response.EvidenceRefs) < 7 {
		t.Fatalf("expected why_not_pass, surface refs, and evidence refs, got %#v", response)
	}
}

func TestProductionUsabilityValAProofsStayInactiveWithoutVal0(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/vala/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val A proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode inactive Val A proofs: %v", err)
	}
	if response.CurrentState == operability.ProductionUsabilityValAStateActive {
		t.Fatalf("expected inactive Val A proofs without Val 0 dependency, got %#v", response)
	}
	if response.Val0DependencyState == "" || response.Val0FoundationState == operability.ProductionUsabilityVal0StateActive {
		t.Fatalf("expected explicit non-active Val 0 dependency, got %#v", response)
	}
}

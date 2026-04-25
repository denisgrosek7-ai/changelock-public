package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	operability "github.com/denisgrosek/changelock/internal/operability"
)

func TestProductionUsabilityValCFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/production/usability-operability-recovery/valc/readiness?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValCReadinessResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode readiness: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValCReadinessStateActive || len(response.Model.Items) != 3 {
					t.Fatalf("unexpected readiness response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valc/guided-readiness?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValCGuidedReadinessResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode guided readiness: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValCGuidedReadinessStateActive || !response.Model.GoLiveAllowed {
					t.Fatalf("unexpected guided readiness response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valc/support-bundle?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValCSupportBundleResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode support bundle: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValCSupportBundleStateActive || !response.Model.ManifestPresent {
					t.Fatalf("unexpected support bundle response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valc/diagnostics?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValCDiagnosticsResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode diagnostics: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValCDiagnosticsStateActive || len(response.Model.Sections) != 5 {
					t.Fatalf("unexpected diagnostics response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valc/health-snapshot?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValCHealthSnapshotResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode health snapshot: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValCHealthSnapshotStateActive || len(response.Model.ComponentStates) != 4 {
					t.Fatalf("unexpected health snapshot response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valc/recovery-playbooks?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValCRecoveryPlaybookResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode recovery playbooks: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValCRecoveryPlaybookStateActive || len(response.Model.Items) != 9 {
					t.Fatalf("unexpected recovery playbooks response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valc/upgrade-rollback-advisory?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValCUpgradeAdvisoryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode upgrade advisory: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValCUpgradeAdvisoryStateActive || !response.Model.RollbackAvailable {
					t.Fatalf("unexpected upgrade advisory response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valc/permission-support-flows?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValCPermissionSupportResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode permission support flows: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValCPermissionSupportStateActive || len(response.Model.Items) != 5 {
					t.Fatalf("unexpected permission support flow response %#v", response)
				}
			},
		},
		{
			path: "/v1/production/usability-operability-recovery/valc/redaction-export-safety?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response productionUsabilityValCExportSafetyResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode export safety: %v", err)
				}
				if response.CurrentState != operability.ProductionUsabilityValCExportSafetyStateActive || !response.Model.AuditorSafe {
					t.Fatalf("unexpected export safety response %#v", response)
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

func TestProductionUsabilityValCProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/valc/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val C proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val C proofs: %v", err)
	}
	if response.CurrentState != operability.ProductionUsabilityValCStateActive {
		t.Fatalf("expected active Val C proofs state, got %#v", response)
	}
	if response.Val0FoundationState != operability.ProductionUsabilityVal0StateActive || response.ValACoreState != operability.ProductionUsabilityValAStateActive || response.ValBResilienceState != operability.ProductionUsabilityValBStateActive {
		t.Fatalf("expected active Val 0/Val A/Val B dependencies, got %#v", response)
	}
	if response.ValCState != operability.ProductionUsabilityValCStateActive {
		t.Fatalf("expected active Val C supportability state, got %#v", response)
	}
	if response.Point4State != operability.ProductionUsabilityPoint4StateNotComplete {
		t.Fatalf("expected point 4 to remain not complete, got %#v", response)
	}
	if len(response.WhyPoint4NotPass) == 0 || len(response.SurfaceRefs) < 11 || len(response.EvidenceRefs) < 9 {
		t.Fatalf("expected why_not_pass, surface refs, and evidence refs, got %#v", response)
	}
}

func TestProductionUsabilityValCProofsStayInactiveWithoutDependencies(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/production/usability-operability-recovery/valc/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val C proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response productionUsabilityValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode inactive Val C proofs: %v", err)
	}
	if response.CurrentState == operability.ProductionUsabilityValCStateActive {
		t.Fatalf("expected inactive Val C proofs without dependencies, got %#v", response)
	}
	if response.Val0DependencyState == "" || response.ValADependencyState == "" || response.ValBDependencyState == "" {
		t.Fatalf("expected explicit dependency states, got %#v", response)
	}
}

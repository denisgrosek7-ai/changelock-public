package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/workflow"
)

func TestEnterpriseWorkflowAuthorityValCFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/enterprise/workflow-authority/valc/closure-validation-enforcement?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValCClosureValidationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode closure validation enforcement: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValCClosureValidationEnforcementStateActive || len(response.Model.RequiredChecks) < 7 {
					t.Fatalf("unexpected closure validation enforcement %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/valc/workflow-ledger?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValCWorkflowLedgerResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode workflow ledger: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValCWorkflowLedgerStateActive || !response.Model.AppendOnly || !response.Model.SignedEntries {
					t.Fatalf("unexpected workflow ledger %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/valc/stale-reopen-handling?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValCStaleReopenResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode stale/reopen handling: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValCStaleReopenHandlingStateActive || len(response.Model.ReopenTriggers) < 4 {
					t.Fatalf("unexpected stale/reopen handling %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/valc/rollback-linkage?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValCRollbackLinkageResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode rollback linkage: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValCRollbackLinkageStateActive || len(response.Model.RequiredRefs) < 5 {
					t.Fatalf("unexpected rollback linkage %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/valc/governance-mapping?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValCGovernanceMappingResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode governance mapping: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValCGovernanceMappingStateActive || len(response.Model.RequiredMappings) < 5 {
					t.Fatalf("unexpected governance mapping %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/valc/replay-recovery-hardening?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValCReplayRecoveryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode replay/recovery hardening: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValCReplayRecoveryHardeningStateActive || len(response.Model.RequiredSources) < 3 {
					t.Fatalf("unexpected replay/recovery hardening %#v", response)
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

func TestEnterpriseWorkflowAuthorityValCProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/enterprise/workflow-authority/valc/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val C proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response enterpriseWorkflowAuthorityValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val C proofs: %v", err)
	}
	if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValCStateActive {
		t.Fatalf("expected active Val C proofs, got %#v", response)
	}
	if response.ValBState != workflow.EnterpriseWorkflowAuthorityValBStateActive || response.Phase4State != phase4ProofStateActive {
		t.Fatalf("expected active dependencies, got %#v", response)
	}
	if len(response.SurfaceRefs) < 10 || len(response.DeferredScope) == 0 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected surface refs, deferred scope, and integration summary, got %#v", response)
	}
}

func TestEnterpriseWorkflowAuthorityValCProofsStayInactiveWithoutValB(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/enterprise/workflow-authority/valc/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val C proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response enterpriseWorkflowAuthorityValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode inactive Val C proofs: %v", err)
	}
	if response.CurrentState == workflow.EnterpriseWorkflowAuthorityValCStateActive {
		t.Fatalf("expected inactive Val C proofs without Val B baseline, got %#v", response)
	}
	if response.ValBState == workflow.EnterpriseWorkflowAuthorityValBStateActive {
		t.Fatalf("expected inactive Val B dependency, got %#v", response)
	}
}

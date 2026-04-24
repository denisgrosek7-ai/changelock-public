package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/workflow"
)

func TestEnterpriseWorkflowAuthorityValDFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/enterprise/workflow-authority/vald/connector-correctness-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValDConnectorCorrectnessResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode connector correctness review: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValDConnectorCorrectnessReviewStateActive || len(response.Model.RequiredConnectors) != 3 {
					t.Fatalf("unexpected connector correctness review %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vald/approval-boundary-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValDApprovalBoundaryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode approval boundary review: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValDApprovalBoundaryReviewStateActive || len(response.Model.RequiredActionClasses) != 5 {
					t.Fatalf("unexpected approval boundary review %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vald/exception-expiry-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValDExceptionExpiryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode exception expiry review: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValDExceptionExpiryReviewStateActive || len(response.Model.RequiredLifecycleStages) != 8 {
					t.Fatalf("unexpected exception expiry review %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vald/closure-correctness-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValDClosureCorrectnessResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode closure correctness review: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValDClosureCorrectnessReviewStateActive || len(response.Model.RequiredChecks) < 5 {
					t.Fatalf("unexpected closure correctness review %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vald/reconciliation-conflict-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValDReconciliationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode reconciliation conflict review: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValDReconciliationConflictReviewStateActive || len(response.Model.RequiredSignals) < 5 {
					t.Fatalf("unexpected reconciliation conflict review %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vald/workflow-ledger-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValDWorkflowLedgerResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode workflow ledger review: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValDWorkflowLedgerReviewStateActive || len(response.Model.RequiredRecordTypes) < 7 {
					t.Fatalf("unexpected workflow ledger review %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vald/governance-traceability-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValDGovernanceTraceabilityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode governance traceability review: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValDGovernanceTraceabilityReviewStateActive || len(response.Model.RequiredDecisionClasses) < 6 {
					t.Fatalf("unexpected governance traceability review %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vald/reopen-rollback-review?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValDReopenRollbackResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode reopen/rollback review: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValDReopenRollbackReviewStateActive || len(response.Model.RequiredRollbackStates) != 4 {
					t.Fatalf("unexpected reopen/rollback review %#v", response)
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

func TestEnterpriseWorkflowAuthorityValDProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/enterprise/workflow-authority/vald/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val D proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response enterpriseWorkflowAuthorityValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val D proofs: %v", err)
	}
	if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValDStateActive {
		t.Fatalf("expected active Val D proofs, got %#v", response)
	}
	if response.ValCState != workflow.EnterpriseWorkflowAuthorityValCStateActive || response.Phase4State != phase4ProofStateActive {
		t.Fatalf("expected active dependencies, got %#v", response)
	}
	if len(response.SurfaceRefs) < 12 || len(response.IntegrationSummary) == 0 || response.DeferredScope != nil {
		t.Fatalf("expected surface refs, integration summary, and no deferred scope, got %#v", response)
	}
}

func TestEnterpriseWorkflowAuthorityValDProofsStayInactiveWithoutValC(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/enterprise/workflow-authority/vald/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val D proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response enterpriseWorkflowAuthorityValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode inactive Val D proofs: %v", err)
	}
	if response.CurrentState == workflow.EnterpriseWorkflowAuthorityValDStateActive {
		t.Fatalf("expected inactive Val D proofs without Val C baseline, got %#v", response)
	}
	if response.ValCState == workflow.EnterpriseWorkflowAuthorityValCStateActive {
		t.Fatalf("expected inactive Val C dependency, got %#v", response)
	}
}

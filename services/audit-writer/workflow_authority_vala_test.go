package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/workflow"
)

func TestEnterpriseWorkflowAuthorityValAFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/enterprise/workflow-authority/vala/event-orchestration?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValAEventOrchestrationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode event orchestration: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValAEventOrchestrationStateActive || len(response.Model.EventClasses) < 7 {
					t.Fatalf("unexpected event orchestration %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vala/lifecycle-connectors?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValALifecycleConnectorsResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode lifecycle connectors: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValALifecycleConnectorsStateActive || len(response.Items) != 3 {
					t.Fatalf("unexpected lifecycle connectors %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vala/evidence-bundle-injection?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValAEvidenceBundleResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode evidence bundle injection: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValAEvidenceBundleInjectionStateActive || len(response.Items) != 3 {
					t.Fatalf("unexpected evidence bundle injection %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vala/ticket-change-projection?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValAProjectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode ticket/change projection: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValATicketChangeProjectionStateActive || len(response.Items) != 3 {
					t.Fatalf("unexpected ticket/change projection %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vala/reconciliation-baseline?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValAReconciliationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode reconciliation baseline: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValAReconciliationBaselineStateActive || len(response.Items) != 3 {
					t.Fatalf("unexpected reconciliation baseline %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/vala/idempotent-mutation-discipline?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValAIdempotentResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode idempotent mutation discipline: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValAIdempotentMutationStateActive || len(response.Items) != 3 {
					t.Fatalf("unexpected idempotent mutation discipline %#v", response)
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

func TestEnterpriseWorkflowAuthorityValAProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/enterprise/workflow-authority/vala/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val A proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response enterpriseWorkflowAuthorityValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val A proofs: %v", err)
	}
	if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValAStateActive {
		t.Fatalf("expected active Val A proofs, got %#v", response)
	}
	if response.Val0State != workflow.EnterpriseWorkflowAuthorityVal0StateActive || response.Phase4State != phase4ProofStateActive {
		t.Fatalf("expected active dependencies, got %#v", response)
	}
	if len(response.SurfaceRefs) < 8 || len(response.DeferredScope) == 0 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected surface refs, deferred scope, and integration summary, got %#v", response)
	}
}

func TestEnterpriseWorkflowAuthorityValAProofsStayInactiveWithoutVal0(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/enterprise/workflow-authority/vala/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val A proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response enterpriseWorkflowAuthorityValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode inactive Val A proofs: %v", err)
	}
	if response.CurrentState == workflow.EnterpriseWorkflowAuthorityValAStateActive {
		t.Fatalf("expected inactive Val A proofs without Val 0 baseline, got %#v", response)
	}
	if response.Val0State == workflow.EnterpriseWorkflowAuthorityVal0StateActive {
		t.Fatalf("expected inactive Val 0 dependency, got %#v", response)
	}
}

func TestEnterpriseWorkflowAuthorityValAProofsStateStaysInactiveWithoutFullConnectorSet(t *testing.T) {
	eventOrchestration := workflow.EnterpriseWorkflowAuthorityValAEventOrchestration()
	lifecycleConnectors := workflow.EnterpriseWorkflowAuthorityValALifecycleConnectors()[:2]
	evidenceBundle := workflow.EnterpriseWorkflowAuthorityValAEvidenceBundleInjection()
	projection := workflow.EnterpriseWorkflowAuthorityValATicketChangeProjection()
	reconciliation := workflow.EnterpriseWorkflowAuthorityValAReconciliationBaseline()
	idempotent := workflow.EnterpriseWorkflowAuthorityValAIdempotentMutationDiscipline()

	got := enterpriseWorkflowAuthorityValAProofsCurrentState(
		workflow.EnterpriseWorkflowAuthorityVal0StateActive,
		eventOrchestration,
		lifecycleConnectors,
		evidenceBundle,
		projection,
		reconciliation,
		idempotent,
	)
	if got == workflow.EnterpriseWorkflowAuthorityValAStateActive {
		t.Fatalf("expected non-active Val A proofs state without full jira/servicenow/github connector coverage, got %q", got)
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/workflow"
)

func TestEnterpriseWorkflowAuthorityVal0FoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/enterprise/workflow-authority/val0/authority-boundaries?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityVal0BoundaryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode authority boundaries: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityVal0BoundaryStateActive || len(response.Items) != 6 {
					t.Fatalf("unexpected authority boundaries %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/val0/state-machine?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityVal0StateMachineResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode state machine: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityVal0StateMachineStateActive || len(response.Model.States) < 14 {
					t.Fatalf("unexpected state machine %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/val0/external-projection-rules?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityVal0ProjectionResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode projection rules: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityVal0ProjectionStateActive || len(response.Items) != 3 {
					t.Fatalf("unexpected projection rules %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/val0/approval-contract?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityVal0ApprovalContractResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode approval contract: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityVal0ApprovalContractStateActive || len(response.Model.AntiReplayRules) == 0 {
					t.Fatalf("unexpected approval contract %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/val0/exception-lifecycle?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityVal0ExceptionLifecycleResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode exception lifecycle: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityVal0ExceptionLifecycleStateActive || len(response.Items) != 8 {
					t.Fatalf("unexpected exception lifecycle %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/val0/closure-validation?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityVal0ClosureValidationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode closure validation: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityVal0ClosureValidationStateActive || len(response.Model.RequiredChecks) == 0 {
					t.Fatalf("unexpected closure validation %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/val0/separation-of-duties?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityVal0SeparationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode separation-of-duties: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityVal0SeparationOfDutiesStateActive || len(response.Items) != 4 {
					t.Fatalf("unexpected separation-of-duties %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/val0/time-authority?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityVal0TimeAuthorityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode time authority: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityVal0TimeAuthorityStateActive || response.Model.ClockSkewTolerance == "" {
					t.Fatalf("unexpected time authority %#v", response)
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

func TestEnterpriseWorkflowAuthorityVal0ProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/enterprise/workflow-authority/val0/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected val0 proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response enterpriseWorkflowAuthorityVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode val0 proofs: %v", err)
	}
	if response.CurrentState != workflow.EnterpriseWorkflowAuthorityVal0StateActive {
		t.Fatalf("expected active val0 proofs, got %#v", response)
	}
	if response.Phase4State != phase4ProofStateActive {
		t.Fatalf("expected active phase4 dependency, got %#v", response)
	}
	if len(response.DeferredScope) == 0 || len(response.SurfaceRefs) < 9 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected deferred scope and integration summary, got %#v", response)
	}
}

func TestEnterpriseWorkflowAuthorityVal0ProofsStayInactiveWithoutPhase4(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/enterprise/workflow-authority/val0/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected val0 proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response enterpriseWorkflowAuthorityVal0ProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode inactive val0 proofs: %v", err)
	}
	if response.CurrentState == workflow.EnterpriseWorkflowAuthorityVal0StateActive {
		t.Fatalf("expected inactive val0 proofs without phase4 baseline, got %#v", response)
	}
	if response.Phase4State == phase4ProofStateActive {
		t.Fatalf("expected inactive phase4 dependency, got %#v", response)
	}
}

func seedEnterprisePhase4AuthorityArtifacts(t *testing.T, handler http.Handler) {
	t.Helper()

	requests := []struct {
		method string
		path   string
		body   string
		token  string
		want   int
	}{
		{
			method: http.MethodPost,
			path:   "/v1/enterprise/workflow/lifecycle?tenant_id=acme&environment=prod&repo=github.com/acme/api",
			body: `{
			  "input":{
			    "workflow_id":"wf-point3-100",
			    "artifact_type":"finding",
			    "subject_ref":"cluster-a/acme-prod/Deployment/api",
			    "severity":"critical",
			    "requested_state":"resolved",
			    "validation_required":true,
			    "validation_state":"verified",
			    "owners":{"finding_owner":"team-api","remediation_owner":"team-api","approver":"security-admin","compliance_owner":"compliance-lead"},
			    "evidence_refs":["event://deploy-gate/1","report:/v1/runtime/phase2/proofs"]
			  }
			}`,
			token: "Bearer operator-demo-token",
			want:  http.StatusCreated,
		},
		{
			method: http.MethodPost,
			path:   "/v1/enterprise/workflow/connectors/reconcile?tenant_id=acme&environment=prod&repo=github.com/acme/api",
			body: `{
			  "input":{
			    "workflow_id":"wf-point3-100",
			    "subject_ref":"cluster-a/acme-prod/Deployment/api",
			    "connector_system":"jira",
			    "connector_ref":"JIRA-POINT3-100",
			    "object_type":"ticket",
			    "internal_state":"resolved",
			    "external_state":"closed",
			    "validation_state":"verified",
			    "health":{"current_state":"healthy"},
			    "evidence_refs":["ticket://JIRA-POINT3-100"]
			  }
			}`,
			token: "Bearer operator-demo-token",
			want:  http.StatusCreated,
		},
		{
			method: http.MethodPost,
			path:   "/v1/enterprise/partner-trust/intake?tenant_id=acme&environment=prod&repo=github.com/acme/api&partner_id=vendor-a",
			body: `{
			  "input":{
			    "partner_id":"vendor-a",
			    "organization":"Vendor A",
			    "trust_domain":"suppliers.acme",
			    "handoff_ref":"handoff-point3-100",
			    "verification_status":"verified",
			    "freshness_state":"fresh",
			    "policy_compatibility":"compatible",
			    "incident_disclosure_status":"shared",
			    "partner_visible_evidence":["sealed://proof/point3-100"],
			    "evidence_refs":["handoff://point3-100"]
			  }
			}`,
			token: "Bearer security-admin-demo-token",
			want:  http.StatusCreated,
		},
		{
			method: http.MethodPost,
			path:   "/v1/enterprise/governance/compliance-mapping?tenant_id=acme&environment=prod&repo=github.com/acme/api",
			body: `{
			  "input":{
			    "subject_ref":"cluster-a/acme-prod/Deployment/api",
			    "control_family":"soc2.cc7",
			    "control_id":"CC7.2",
			    "coverage_state":"full",
			    "freshness_state":"fresh",
			    "evidence_refs":["event://deploy-gate/1"],
			    "technical_event_refs":["deploy_gate_decision","runtime_phase2_proofs"]
			  }
			}`,
			token: "Bearer security-admin-demo-token",
			want:  http.StatusCreated,
		},
		{
			method: http.MethodPost,
			path:   "/v1/enterprise/governance/policy-drift?tenant_id=acme&environment=prod&repo=github.com/acme/api",
			body: `{
			  "input":{
			    "subject_ref":"cluster-a/acme-prod/Deployment/api",
			    "actor":"security-admin",
			    "previous_mode":"deny",
			    "current_mode":"exception",
			    "change_reason":"temporary maintenance window",
			    "exception_id":"exc-point3-100",
			    "impacted_controls":["CC7.2"],
			    "evidence_refs":["exception://exc-point3-100"]
			  }
			}`,
			token: "Bearer security-admin-demo-token",
			want:  http.StatusCreated,
		},
		{
			method: http.MethodPost,
			path:   "/v1/enterprise/governance/executive-report?tenant_id=acme&environment=prod&repo=github.com/acme/api",
			body:   `{"scope_ref":"tenant:acme"}`,
			token:  "Bearer security-admin-demo-token",
			want:   http.StatusOK,
		},
	}

	for _, tc := range requests {
		req := httptest.NewRequest(tc.method, tc.path, bytes.NewBufferString(tc.body))
		req.Header.Set("Authorization", tc.token)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != tc.want {
			t.Fatalf("expected %d for %s, got %d: %s", tc.want, tc.path, rec.Code, rec.Body.String())
		}
	}
}

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/workflow"
)

func TestEnterpriseWorkflowAuthorityValBFoundationHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	testCases := []struct {
		path   string
		decode func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			path: "/v1/enterprise/workflow-authority/valb/signed-authorizations?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValBSignedAuthorizationsResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode signed authorizations: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValBSignedAuthorizationsStateActive || len(response.Model.RequiredArtifactFields) < 13 {
					t.Fatalf("unexpected signed authorizations %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/valb/break-glass-flow?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValBBreakGlassResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode break-glass flow: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValBBreakGlassStateActive || !response.Model.DualControlRequired {
					t.Fatalf("unexpected break-glass flow %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/valb/managed-exception-registry?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValBExceptionRegistryResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode managed exception registry: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValBManagedExceptionRegistryStateActive || len(response.Model.LifecycleStages) != 8 {
					t.Fatalf("unexpected managed exception registry %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/valb/expiry-revocation-enforcement?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValBExpiryRevocationResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode expiry revocation enforcement: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValBExpiryRevocationStateActive || response.Model.ClockSkewTolerance == "" {
					t.Fatalf("unexpected expiry revocation enforcement %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/valb/anti-replay-protection?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValBAntiReplayResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode anti-replay protection: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValBAntiReplayStateActive || len(response.Model.NonceOrJTIFields) < 3 {
					t.Fatalf("unexpected anti-replay protection %#v", response)
				}
			},
		},
		{
			path: "/v1/enterprise/workflow-authority/valb/approval-traceability?tenant_id=acme&environment=prod",
			decode: func(t *testing.T, rec *httptest.ResponseRecorder) {
				var response enterpriseWorkflowAuthorityValBApprovalTraceabilityResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("decode approval traceability: %v", err)
				}
				if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValBApprovalTraceabilityStateActive || len(response.Model.RequiredTraceFields) < 9 {
					t.Fatalf("unexpected approval traceability %#v", response)
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

func TestEnterpriseWorkflowAuthorityValBProofsHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedEnterprisePhase4AuthorityArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/enterprise/workflow-authority/valb/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val B proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response enterpriseWorkflowAuthorityValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode Val B proofs: %v", err)
	}
	if response.CurrentState != workflow.EnterpriseWorkflowAuthorityValBStateActive {
		t.Fatalf("expected active Val B proofs, got %#v", response)
	}
	if response.ValAState != workflow.EnterpriseWorkflowAuthorityValAStateActive || response.Phase4State != phase4ProofStateActive {
		t.Fatalf("expected active dependencies, got %#v", response)
	}
	if len(response.SurfaceRefs) < 10 || len(response.DeferredScope) == 0 || len(response.IntegrationSummary) == 0 {
		t.Fatalf("expected surface refs, deferred scope, and integration summary, got %#v", response)
	}
}

func TestEnterpriseWorkflowAuthorityValBProofsStayInactiveWithoutValA(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/enterprise/workflow-authority/valb/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected Val B proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response enterpriseWorkflowAuthorityValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode inactive Val B proofs: %v", err)
	}
	if response.CurrentState == workflow.EnterpriseWorkflowAuthorityValBStateActive {
		t.Fatalf("expected inactive Val B proofs without Val A baseline, got %#v", response)
	}
	if response.ValAState == workflow.EnterpriseWorkflowAuthorityValAStateActive {
		t.Fatalf("expected inactive Val A dependency, got %#v", response)
	}
}

func TestEnterpriseWorkflowAuthorityValBProofsStateStaysInactiveWithoutFullSignedAuthorizationActionClassCoverage(t *testing.T) {
	signedAuthorizations := workflow.EnterpriseWorkflowAuthorityValBSignedAuthorizations()
	signedAuthorizations.SupportedActionClasses = []string{
		workflow.WorkflowAuthorityActionApprovalRequired,
		workflow.WorkflowAuthoritySensitiveActionBreakGlass,
		workflow.WorkflowAuthoritySensitiveActionBroadScopeOverride,
		workflow.WorkflowAuthoritySensitiveActionLongLivedException,
	}
	breakGlass := workflow.EnterpriseWorkflowAuthorityValBBreakGlassFlow()
	exceptionRegistry := workflow.EnterpriseWorkflowAuthorityValBManagedExceptionRegistry()
	expiryRevocation := workflow.EnterpriseWorkflowAuthorityValBExpiryRevocationEnforcement()
	antiReplay := workflow.EnterpriseWorkflowAuthorityValBAntiReplayProtection()
	traceability := workflow.EnterpriseWorkflowAuthorityValBApprovalTraceability()

	got := enterpriseWorkflowAuthorityValBProofsCurrentState(
		workflow.EnterpriseWorkflowAuthorityValAStateActive,
		signedAuthorizations,
		breakGlass,
		exceptionRegistry,
		expiryRevocation,
		antiReplay,
		traceability,
	)
	if got == workflow.EnterpriseWorkflowAuthorityValBStateActive {
		t.Fatalf("expected non-active Val B proofs state without full signed-authorization action-class coverage, got %q", got)
	}
}

func TestEnterpriseWorkflowAuthorityValBProofsStateStaysInactiveWithoutFullAntiReplayTokenCoverage(t *testing.T) {
	signedAuthorizations := workflow.EnterpriseWorkflowAuthorityValBSignedAuthorizations()
	breakGlass := workflow.EnterpriseWorkflowAuthorityValBBreakGlassFlow()
	exceptionRegistry := workflow.EnterpriseWorkflowAuthorityValBManagedExceptionRegistry()
	expiryRevocation := workflow.EnterpriseWorkflowAuthorityValBExpiryRevocationEnforcement()
	antiReplay := workflow.EnterpriseWorkflowAuthorityValBAntiReplayProtection()
	antiReplay.TokenTypes = []string{
		"signed_authorization_artifact",
		"break_glass_authorization",
	}
	traceability := workflow.EnterpriseWorkflowAuthorityValBApprovalTraceability()

	got := enterpriseWorkflowAuthorityValBProofsCurrentState(
		workflow.EnterpriseWorkflowAuthorityValAStateActive,
		signedAuthorizations,
		breakGlass,
		exceptionRegistry,
		expiryRevocation,
		antiReplay,
		traceability,
	)
	if got == workflow.EnterpriseWorkflowAuthorityValBStateActive {
		t.Fatalf("expected non-active Val B proofs state without full anti-replay token coverage, got %q", got)
	}
}

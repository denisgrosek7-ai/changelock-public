package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/signingidentity"
)

func TestSigningIdentityPolicyMutationRequiresSecurityAdmin(t *testing.T) {
	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", authConfig)

	req := httptest.NewRequest(http.MethodPost, "/v1/signing-identities/policies", bytes.NewBufferString(`{
	  "issuer":"https://token.actions.githubusercontent.com",
	  "signer_identity":"https://github.com/my-org/acme-app/.github/workflows/build.yml@refs/heads/main",
	  "repository":"my-org/acme-app"
	}`))
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestSigningIdentityPolicyCreateAndEvaluate(t *testing.T) {
	t.Setenv("CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT", signingidentity.EnforcementEnforce)
	t.Setenv("CHANGELOCK_SIGNER_IDENTITY_WORKFLOWS_DIR", t.TempDir())
	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", authConfig)

	createReq := httptest.NewRequest(http.MethodPost, "/v1/signing-identities/policies", bytes.NewBufferString(`{
	  "issuer":"https://token.actions.githubusercontent.com",
	  "signer_identity":"https://github.com/my-org/acme-app/.github/workflows/build.yml@refs/heads/main",
	  "subject":"repo:my-org/acme-app",
	  "repository":"my-org/acme-app",
	  "workflow":".github/workflows/build.yml",
	  "ref":"refs/heads/main"
	}`))
	createReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()
	handler.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", createRec.Code, createRec.Body.String())
	}

	evalReq := httptest.NewRequest(http.MethodPost, "/v1/signing-identities/evaluate", bytes.NewBufferString(`{
	  "issuer":"https://token.actions.githubusercontent.com",
	  "signer_identity":"https://github.com/my-org/acme-app/.github/workflows/build.yml@refs/heads/main",
	  "subject":"repo:my-org/acme-app",
	  "repository":"my-org/acme-app",
	  "workflow":".github/workflows/build.yml",
	  "ref":"refs/heads/main",
	  "transparency_state":"verified"
	}`))
	evalReq.Header.Set("Authorization", "Bearer service-internal-demo-token")
	evalReq.Header.Set("Content-Type", "application/json")
	evalRec := httptest.NewRecorder()
	handler.ServeHTTP(evalRec, evalReq)
	if evalRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", evalRec.Code, evalRec.Body.String())
	}

	var decision signingidentity.Decision
	if err := json.NewDecoder(evalRec.Body).Decode(&decision); err != nil {
		t.Fatalf("decode evaluate response: %v", err)
	}
	if decision.ReasonCode != signingidentity.ReasonAuthorized || decision.Authorized != signingidentity.AuthorizationAuthorized {
		t.Fatalf("unexpected decision %#v", decision)
	}
}

func TestSigningIdentityObservationsRespectTenantScopedOIDC(t *testing.T) {
	t.Setenv("CHANGELOCK_SIGNER_IDENTITY_WORKFLOWS_DIR", t.TempDir())
	store := audit.NewMemoryStore()
	mustIngestEnterpriseEvent(t, store, audit.Event{
		Component:   "deploy-gate",
		EventType:   audit.EventTypePolicyDecision,
		Decision:    audit.DecisionAllow,
		TenantID:    "acme",
		Environment: "prod",
		Repo:        "my-org/acme-app",
		Digest:      "sha256:acme",
		Evidence: &audit.Evidence{
			Artifact: &audit.ArtifactEvidence{
				SignerIdentity: "https://github.com/my-org/acme-app/.github/workflows/build.yml@refs/heads/main",
				Issuer:         "https://token.actions.githubusercontent.com",
				Subject:        "repo:my-org/acme-app",
				Repository:     "my-org/acme-app",
				Workflow:       ".github/workflows/build.yml",
				Ref:            "refs/heads/main",
				Digest:         "sha256:acme",
			},
			VerificationState: "verified",
		},
	})
	mustIngestEnterpriseEvent(t, store, audit.Event{
		Component:   "deploy-gate",
		EventType:   audit.EventTypePolicyDecision,
		Decision:    audit.DecisionAllow,
		TenantID:    "globex",
		Environment: "prod",
		Repo:        "my-org/globex-app",
		Digest:      "sha256:globex",
		Evidence: &audit.Evidence{
			Artifact: &audit.ArtifactEvidence{
				SignerIdentity: "https://github.com/my-org/globex-app/.github/workflows/build.yml@refs/heads/main",
				Issuer:         "https://token.actions.githubusercontent.com",
				Subject:        "repo:my-org/globex-app",
				Repository:     "my-org/globex-app",
				Workflow:       ".github/workflows/build.yml",
				Ref:            "refs/heads/main",
				Digest:         "sha256:globex",
			},
			VerificationState: "verified",
		},
	})

	cfg, signer := newOIDCHandlerConfig(t, true, true)
	handler := newHandlerWithAuth(store, "memory", cfg)
	authHeader := "Bearer " + signer.token(t, map[string]any{
		"sub":       "viewer@example.com",
		"groups":    []string{"changelock-viewers"},
		"tenant_id": "acme",
	})

	req := httptest.NewRequest(http.MethodGet, "/v1/signing-identities?limit=10", nil)
	req.Header.Set("Authorization", authHeader)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response signingIdentityObservationsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode observations response: %v", err)
	}
	if len(response.Items) != 1 || response.Items[0].TenantID != "acme" {
		t.Fatalf("unexpected tenant-scoped observations %#v", response)
	}

	rejectReq := httptest.NewRequest(http.MethodGet, "/v1/signing-identities?tenant_id=globex", nil)
	rejectReq.Header.Set("Authorization", authHeader)
	rejectRec := httptest.NewRecorder()
	handler.ServeHTTP(rejectRec, rejectReq)
	if rejectRec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", rejectRec.Code, rejectRec.Body.String())
	}
}

func TestSigningIdentityStatusAndFindingsSurfaceUnauthorizedObservation(t *testing.T) {
	t.Setenv("CHANGELOCK_SIGNER_IDENTITY_ENFORCEMENT", signingidentity.EnforcementMonitor)
	t.Setenv("CHANGELOCK_SIGNER_IDENTITY_WORKFLOWS_DIR", t.TempDir())
	store := audit.NewMemoryStore()
	if _, err := store.Ingest(context.Background(), audit.Event{
		Component:   "deploy-gate",
		EventType:   audit.EventTypePolicyDecision,
		Decision:    audit.DecisionAllow,
		TenantID:    "acme",
		Environment: "prod",
		Repo:        "my-org/acme-app",
		Digest:      "sha256:acme",
		Timestamp:   time.Now().UTC(),
		Evidence: &audit.Evidence{
			Artifact: &audit.ArtifactEvidence{
				SignerIdentity: "https://github.com/my-org/acme-app/.github/workflows/rogue.yml@refs/heads/main",
				Issuer:         "https://token.actions.githubusercontent.com",
				Subject:        "repo:my-org/acme-app",
				Repository:     "my-org/acme-app",
				Workflow:       ".github/workflows/rogue.yml",
				Ref:            "refs/heads/main",
				Digest:         "sha256:acme",
			},
			VerificationState: "verified",
		},
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	handler := newHandler(store, "memory")

	statusReq := httptest.NewRequest(http.MethodGet, "/v1/signing-identities/status?tenant_id=acme", nil)
	statusRec := httptest.NewRecorder()
	handler.ServeHTTP(statusRec, statusReq)
	if statusRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", statusRec.Code, statusRec.Body.String())
	}

	var status signingIdentityStatusResponse
	if err := json.NewDecoder(statusRec.Body).Decode(&status); err != nil {
		t.Fatalf("decode status response: %v", err)
	}
	if status.Status.Unknown == 0 {
		t.Fatalf("expected unknown identity count, got %#v", status.Status)
	}

	findingsReq := httptest.NewRequest(http.MethodGet, "/v1/signing-identities/findings?tenant_id=acme", nil)
	findingsRec := httptest.NewRecorder()
	handler.ServeHTTP(findingsRec, findingsReq)
	if findingsRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", findingsRec.Code, findingsRec.Body.String())
	}

	var findings signingIdentityFindingsResponse
	if err := json.NewDecoder(findingsRec.Body).Decode(&findings); err != nil {
		t.Fatalf("decode findings response: %v", err)
	}
	if len(findings.Items) == 0 || findings.Items[0].Type != signingidentity.FindingUnauthorizedIdentity {
		t.Fatalf("unexpected findings %#v", findings)
	}
}

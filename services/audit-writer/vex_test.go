package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

func TestVEXStatementsRequireSecurityAdminForMutation(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	handler := newHandlerWithDeps(store, "memory", authConfig, nil, newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}), nil)

	viewerCreateReq := httptest.NewRequest(http.MethodPost, "/v1/vex", bytes.NewBufferString(`{
	  "vulnerability_id":"CVE-2026-6100",
	  "scope":{"image_digest":"sha256:vex-auth"},
	  "status":"not_affected",
	  "justification":"viewer should not be able to mutate"
	}`))
	viewerCreateReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	viewerCreateReq.Header.Set("Content-Type", "application/json")
	viewerCreateRec := httptest.NewRecorder()
	handler.ServeHTTP(viewerCreateRec, viewerCreateReq)
	if viewerCreateRec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", viewerCreateRec.Code, viewerCreateRec.Body.String())
	}

	adminCreateReq := httptest.NewRequest(http.MethodPost, "/v1/vex", bytes.NewBufferString(`{
	  "vulnerability_id":"CVE-2026-6100",
	  "scope":{"image_digest":"sha256:vex-auth","package_name":"openssl"},
	  "status":"not_affected",
	  "justification":"component is not reachable",
	  "action_statement":"monitor upstream fix"
	}`))
	adminCreateReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	adminCreateReq.Header.Set("Content-Type", "application/json")
	adminCreateRec := httptest.NewRecorder()
	handler.ServeHTTP(adminCreateRec, adminCreateReq)
	if adminCreateRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", adminCreateRec.Code, adminCreateRec.Body.String())
	}

	var created vexStatementResponse
	if err := json.NewDecoder(adminCreateRec.Body).Decode(&created); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if created.Statement.Status != internalvex.StatusNotAffected {
		t.Fatalf("unexpected created statement %#v", created.Statement)
	}

	listReq := httptest.NewRequest(http.MethodGet, "/v1/vex?image_digest=sha256:vex-auth&vulnerability_id=CVE-2026-6100", nil)
	listReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	listRec := httptest.NewRecorder()
	handler.ServeHTTP(listRec, listReq)
	if listRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", listRec.Code, listRec.Body.String())
	}

	var listed vexStatementsResponse
	if err := json.NewDecoder(listRec.Body).Decode(&listed); err != nil {
		t.Fatalf("decode list response: %v", err)
	}
	if len(listed.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %#v", listed.Statements)
	}

	revokeReq := httptest.NewRequest(http.MethodPost, "/v1/vex/"+strconv.FormatInt(created.Statement.ID, 10)+"/revoke", nil)
	revokeReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	revokeRec := httptest.NewRecorder()
	handler.ServeHTTP(revokeRec, revokeReq)
	if revokeRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", revokeRec.Code, revokeRec.Body.String())
	}
}

func TestVulnerabilityNetResponseShowsRawResolvedAndActionableCounts(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	handler := newHandlerWithDeps(store, "memory", authConfig, nil, newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}), nil)

	if _, err := store.RecordVulnerabilityScan(context.Background(), audit.VulnerabilityScanRequest{
		ImageDigest: "sha256:vex-net",
		ImageRef:    "ghcr.io/example/api@sha256:vex-net",
		Scanner:     "trivy",
		StartedAt:   time.Now().UTC(),
		CompletedAt: ptrTimeMain(time.Now().UTC()),
		Status:      audit.VulnerabilityScanStatusCompleted,
		Findings: []audit.VulnerabilityFindingInput{
			{CVEID: "CVE-2026-7001", Severity: "HIGH", PackageName: "openssl", PackageVersion: "3.0.14-r0"},
			{CVEID: "CVE-2026-7002", Severity: "HIGH", PackageName: "glibc", PackageVersion: "2.39-r0"},
			{CVEID: "CVE-2026-7003", Severity: "LOW", PackageName: "busybox", PackageVersion: "1.36.1-r0"},
		},
	}); err != nil {
		t.Fatalf("RecordVulnerabilityScan() error = %v", err)
	}

	if _, err := store.CreateVEXStatement(context.Background(), internalvex.CreateRequest{
		VulnerabilityID: "CVE-2026-7001",
		Scope:           internalvex.Scope{ImageDigest: "sha256:vex-net", PackageName: "openssl"},
		Status:          internalvex.StatusNotAffected,
		Justification:   "openssl code path is not reachable in this service",
	}, "demo-admin"); err != nil {
		t.Fatalf("CreateVEXStatement() error = %v", err)
	}
	if _, err := store.CreateVEXStatement(context.Background(), internalvex.CreateRequest{
		VulnerabilityID: "CVE-2026-7002",
		Scope:           internalvex.Scope{ImageDigest: "sha256:vex-net", PackageName: "glibc"},
		Status:          internalvex.StatusUnderInvestigation,
		Justification:   "runtime exploitability is still under review",
	}, "demo-admin"); err != nil {
		t.Fatalf("CreateVEXStatement() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/vulnerabilities/net?image_digest=sha256:vex-net&severity_threshold=HIGH", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response audit.VulnerabilityNetResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode net response: %v", err)
	}
	if response.RawCount != 3 || response.ResolvedByVEXCount != 1 || response.ActionableCount != 2 {
		t.Fatalf("unexpected counts %#v", response)
	}
	if response.UnderInvestigationCount != 1 {
		t.Fatalf("expected one under investigation finding, got %#v", response)
	}
	if !response.ThresholdBreached {
		t.Fatalf("expected threshold breach, got %#v", response)
	}
	if len(response.Findings) != 2 {
		t.Fatalf("expected actionable findings only, got %#v", response.Findings)
	}
}

func TestVEXStatusFiltersByImageDigestScope(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	handler := newHandlerWithDeps(store, "memory", authConfig, nil, newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}), nil)

	if _, err := store.CreateVEXStatement(context.Background(), internalvex.CreateRequest{
		VulnerabilityID: "CVE-2026-7100",
		Scope:           internalvex.Scope{ImageDigest: "sha256:one"},
		Status:          internalvex.StatusNotAffected,
		Justification:   "digest one is not affected",
	}, "demo-admin"); err != nil {
		t.Fatalf("CreateVEXStatement(one) error = %v", err)
	}
	if _, err := store.CreateVEXStatement(context.Background(), internalvex.CreateRequest{
		VulnerabilityID: "CVE-2026-7100",
		Scope:           internalvex.Scope{ImageDigest: "sha256:two"},
		Status:          internalvex.StatusAffected,
		Justification:   "digest two remains affected",
	}, "demo-admin"); err != nil {
		t.Fatalf("CreateVEXStatement(two) error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/vex?image_digest=sha256:one", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response vexStatementsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode vex statements: %v", err)
	}
	if len(response.Statements) != 1 || response.Statements[0].Scope.ImageDigest != "sha256:one" {
		t.Fatalf("expected scoped digest response, got %#v", response)
	}
}

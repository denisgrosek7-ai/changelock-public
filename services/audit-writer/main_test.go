package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	internalvulnops "github.com/denisgrosek/changelock/internal/vulnops"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeDisabled)
	_ = os.Unsetenv("CHANGELOCK_AUTH_TOKENS_JSON")
	_ = os.Unsetenv("CHANGELOCK_INTERNAL_SERVICE_TOKEN")
	os.Exit(m.Run())
}

func testAuthTokensJSON() string {
	return `[
	  {"token":"viewer-demo-token","subject":"demo-viewer","role":"viewer","token_id":"viewer-demo"},
	  {"token":"operator-demo-token","subject":"demo-operator","role":"operator","token_id":"operator-demo"},
	  {"token":"security-admin-demo-token","subject":"demo-admin","role":"security_admin","token_id":"security-admin-demo"},
	  {"token":"service-internal-demo-token","subject":"policy-engine","role":"service_internal","token_id":"service-internal-demo"}
	]`
}

func postgresReportsTestDSN() string {
	for _, candidate := range []string{
		strings.TrimSpace(os.Getenv("CHANGELOCK_POSTGRES_TEST_DSN")),
		strings.TrimSpace(os.Getenv("CHANGELOCK_POSTGRES_DSN")),
		"postgres://changelock:changelock@127.0.0.1:5433/changelock?sslmode=disable",
	} {
		if candidate != "" {
			return candidate
		}
	}
	return ""
}

func newPostgresReportsTestStore(t *testing.T) *audit.PostgresStore {
	t.Helper()

	dsn := postgresReportsTestDSN()
	if dsn == "" {
		t.Skip("Postgres reports test DSN is not configured")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Skipf("postgres unavailable for reports integration test: %v", err)
	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil {
		t.Skipf("postgres unavailable for reports integration test: %v", err)
	}

	store, err := audit.NewPostgresStore(ctx, dsn)
	if err != nil {
		t.Fatalf("NewPostgresStore() error = %v", err)
	}
	if err := store.Migrate(ctx); err != nil {
		store.Close()
		t.Fatalf("Migrate() error = %v", err)
	}
	t.Cleanup(store.Close)
	return store
}

func TestIngestStoresEvent(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")

	body := bytes.NewBufferString(`{"component":"deploy-gate","event_type":"deploy_gate_decision","decision":"DENY","reasons":["workflow mismatch"]}`)
	req := httptest.NewRequest(http.MethodPost, "/v1/ingest", body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-Id", "req-123")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", rec.Code, rec.Body.String())
	}

	events, err := store.ListEvents(req.Context(), audit.EventFilter{Limit: 10})
	if err != nil {
		t.Fatalf("ListEvents() error = %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
	if events[0].RequestID != "req-123" {
		t.Fatalf("expected request id from header, got %#v", events[0])
	}
}

func TestIngestRejectsInvalidEvent(t *testing.T) {
	handler := newHandler(audit.NewMemoryStore(), "memory")

	req := httptest.NewRequest(http.MethodPost, "/v1/ingest", bytes.NewBufferString(`{"event_type":"policy_decision","decision":"DENY"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestHealthAndReadyExposeProcessVsStoreState(t *testing.T) {
	handler := newHandler(audit.NewMemoryStore(), "memory")

	healthReq := httptest.NewRequest(http.MethodGet, "/health", nil)
	healthRec := httptest.NewRecorder()
	handler.ServeHTTP(healthRec, healthReq)
	if healthRec.Code != http.StatusOK {
		t.Fatalf("expected health 200, got %d: %s", healthRec.Code, healthRec.Body.String())
	}

	readyReq := httptest.NewRequest(http.MethodGet, "/ready", nil)
	readyRec := httptest.NewRecorder()
	handler.ServeHTTP(readyRec, readyReq)
	if readyRec.Code != http.StatusOK {
		t.Fatalf("expected ready 200, got %d: %s", readyRec.Code, readyRec.Body.String())
	}
}

func TestReportsEventsSupportsFilters(t *testing.T) {
	store := audit.NewMemoryStore()
	mustIngest := func(event audit.Event) {
		t.Helper()
		if _, err := store.Ingest(t.Context(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}
	mustIngest(audit.Event{Component: "deploy-gate", EventType: audit.EventTypeDeployGateDecision, Decision: audit.DecisionDeny, TenantID: "acme"})
	mustIngest(audit.Event{Component: "runtime-agent", EventType: audit.EventTypeRuntimeDriftResult, Decision: audit.DecisionAllow, TenantID: "globex"})

	req := httptest.NewRequest(http.MethodGet, "/v1/reports/events?tenant_id=acme&decision=DENY", nil)
	rec := httptest.NewRecorder()
	newHandler(store, "memory").ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response eventsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(response.Events) != 1 || response.Events[0].TenantID != "acme" {
		t.Fatalf("unexpected response %#v", response)
	}
}

func TestReportsSummaryReturnsCounts(t *testing.T) {
	store := audit.NewMemoryStore()
	mustIngest := func(event audit.Event) {
		t.Helper()
		if _, err := store.Ingest(t.Context(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}
	mustIngest(audit.Event{Component: "deploy-gate", EventType: audit.EventTypeDeployGateDecision, Decision: audit.DecisionAllow, TenantID: "acme"})
	mustIngest(audit.Event{Component: "deploy-gate", EventType: audit.EventTypeDeployGateDecision, Decision: audit.DecisionDeny, TenantID: "acme", Reasons: []string{"workflow mismatch"}})

	req := httptest.NewRequest(http.MethodGet, "/v1/reports/summary?tenant_id=acme", nil)
	rec := httptest.NewRecorder()
	newHandler(store, "memory").ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var summary audit.Summary
	if err := json.NewDecoder(rec.Body).Decode(&summary); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if summary.TotalEvents != 2 || summary.TotalAllow != 1 || summary.TotalDeny != 1 {
		t.Fatalf("unexpected summary %#v", summary)
	}
}

func TestPostgresReportsRoundTripPreservesRawEventAndSummary(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeDisabled)

	store := newPostgresReportsTestStore(t)
	handler := newHandler(store, "postgres")
	tenantID := fmt.Sprintf("pgtest-%d", time.Now().UnixNano())

	ingestEvent := func(body string) {
		t.Helper()
		req := httptest.NewRequest(http.MethodPost, "/v1/ingest", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusCreated {
			t.Fatalf("expected 201, got %d: %s", rec.Code, rec.Body.String())
		}
	}

	ingestEvent(fmt.Sprintf(`{
	  "request_id":"req-postgres-deny",
	  "component":"deploy-gate",
	  "event_type":"deploy_gate_decision",
	  "tenant_id":%q,
	  "repo":"my-org/acme-app",
	  "environment":"prod",
	  "decision":"DENY",
	  "reasons":["workflow mismatch"],
	  "policy_version":"tenant-acme-v1",
	  "evidence":{"artifact":{"repository":"my-org/acme-app","digest":"sha256:abc123"}}
	}`, tenantID))
	ingestEvent(fmt.Sprintf(`{
	  "request_id":"req-postgres-drift",
	  "component":"runtime-agent",
	  "event_type":"runtime_drift_result",
	  "tenant_id":%q,
	  "environment":"prod",
	  "decision":"DENY",
	  "drift_result":"image_drift",
	  "reasons":["image drift"]
	}`, tenantID))

	eventsReq := httptest.NewRequest(http.MethodGet, "/v1/reports/events?tenant_id="+tenantID+"&limit=10", nil)
	eventsRec := httptest.NewRecorder()
	handler.ServeHTTP(eventsRec, eventsReq)
	if eventsRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", eventsRec.Code, eventsRec.Body.String())
	}

	var response eventsResponse
	if err := json.NewDecoder(eventsRec.Body).Decode(&response); err != nil {
		t.Fatalf("decode events response: %v", err)
	}
	if len(response.Events) != 2 {
		t.Fatalf("expected 2 events, got %#v", response)
	}

	var deployRecord *audit.StoredEvent
	for i := range response.Events {
		if response.Events[i].RequestID == "req-postgres-deny" {
			deployRecord = &response.Events[i]
			break
		}
	}
	if deployRecord == nil {
		t.Fatalf("expected deploy-gate event in %#v", response.Events)
	}
	if len(deployRecord.RawEvent) == 0 || !bytes.Contains(deployRecord.RawEvent, []byte(`"request_id":"req-postgres-deny"`)) {
		t.Fatalf("expected raw_event preservation, got %#v", deployRecord)
	}
	if deployRecord.Evidence == nil || deployRecord.Evidence.Artifact == nil || deployRecord.Evidence.Artifact.Digest != "sha256:abc123" {
		t.Fatalf("expected artifact evidence roundtrip, got %#v", deployRecord)
	}

	summaryReq := httptest.NewRequest(http.MethodGet, "/v1/reports/summary?tenant_id="+tenantID, nil)
	summaryRec := httptest.NewRecorder()
	handler.ServeHTTP(summaryRec, summaryReq)
	if summaryRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", summaryRec.Code, summaryRec.Body.String())
	}

	var summary audit.Summary
	if err := json.NewDecoder(summaryRec.Body).Decode(&summary); err != nil {
		t.Fatalf("decode summary response: %v", err)
	}
	if summary.TotalEvents != 2 || summary.TotalDeny != 2 {
		t.Fatalf("unexpected summary %#v", summary)
	}
	if summary.CountsByEventType[audit.EventTypeDeployGateDecision] != 1 || summary.CountsByEventType[audit.EventTypeRuntimeDriftResult] != 1 {
		t.Fatalf("unexpected event-type counts %#v", summary.CountsByEventType)
	}
	if summary.RecentRuntimeDriftDeny != 1 {
		t.Fatalf("expected runtime drift deny count, got %#v", summary)
	}
}

func TestRuntimeDriftEndpointFiltersEventType(t *testing.T) {
	store := audit.NewMemoryStore()
	if _, err := store.Ingest(t.Context(), audit.Event{
		Component:   "runtime-agent",
		EventType:   audit.EventTypeRuntimeDriftResult,
		Decision:    audit.DecisionDeny,
		TenantID:    "acme",
		DriftResult: "image_drift",
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}
	if _, err := store.Ingest(t.Context(), audit.Event{
		Component: "deploy-gate",
		EventType: audit.EventTypeDeployGateDecision,
		Decision:  audit.DecisionAllow,
		TenantID:  "acme",
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/reports/runtime-drift?tenant_id=acme", nil)
	rec := httptest.NewRecorder()
	newHandler(store, "memory").ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response eventsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(response.Events) != 1 || response.Events[0].EventType != audit.EventTypeRuntimeDriftResult {
		t.Fatalf("unexpected response %#v", response)
	}
}

func TestReportsSetNoStoreHeaders(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/reports/summary", nil)
	rec := httptest.NewRecorder()

	newHandler(audit.NewMemoryStore(), "memory").ServeHTTP(rec, req)

	if got := rec.Header().Get("Cache-Control"); got != "no-store, max-age=0" {
		t.Fatalf("expected no-store cache header, got %q", got)
	}
	if got := rec.Header().Get("X-Content-Type-Options"); got != "nosniff" {
		t.Fatalf("expected nosniff header, got %q", got)
	}
}

func TestCORSAllowsConfiguredOrigin(t *testing.T) {
	t.Setenv("CHANGELOCK_CORS_ALLOW_ORIGINS", "http://localhost:5173")

	req := httptest.NewRequest(http.MethodOptions, "/v1/reports/events", nil)
	req.Header.Set("Origin", "http://localhost:5173")
	req.Header.Set("Access-Control-Request-Method", http.MethodGet)
	rec := httptest.NewRecorder()

	newHandler(audit.NewMemoryStore(), "memory").ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", rec.Code)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:5173" {
		t.Fatalf("expected allow origin header, got %q", got)
	}
	if got := rec.Header().Get("Access-Control-Allow-Headers"); got != "Authorization, Content-Type, X-Request-Id" {
		t.Fatalf("expected authorization header support, got %q", got)
	}
}

func TestCORSRejectsUnknownOriginPreflight(t *testing.T) {
	t.Setenv("CHANGELOCK_CORS_ALLOW_ORIGINS", "http://localhost:5173")

	req := httptest.NewRequest(http.MethodOptions, "/v1/reports/events", nil)
	req.Header.Set("Origin", "http://evil.example")
	req.Header.Set("Access-Control-Request-Method", http.MethodGet)
	rec := httptest.NewRecorder()

	newHandler(audit.NewMemoryStore(), "memory").ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", rec.Code)
	}
}

func TestAllowedOriginsFromEnvDefaults(t *testing.T) {
	old := os.Getenv("CHANGELOCK_CORS_ALLOW_ORIGINS")
	t.Cleanup(func() {
		_ = os.Setenv("CHANGELOCK_CORS_ALLOW_ORIGINS", old)
	})
	_ = os.Unsetenv("CHANGELOCK_CORS_ALLOW_ORIGINS")

	origins := allowedOriginsFromEnv()
	if _, ok := origins["http://127.0.0.1:5173"]; !ok {
		t.Fatalf("expected default vite origin")
	}
	if _, ok := origins["http://127.0.0.1:3000"]; !ok {
		t.Fatalf("expected default docker ui origin")
	}
}

func TestExceptionsLifecycleEndpoints(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")

	createBody := bytes.NewBufferString(`{
	  "exception_id":"EX-2026-001",
	  "exception_type":"BREAK_GLASS",
	  "tenant_id":"acme",
	  "environment":"prod",
	  "namespace":"acme-prod",
	  "reason":"P0 production fix",
	  "ticket_id":"INC-1234",
	  "approved_by":"oncall@example.com",
	  "ttl_hours":2
	}`)
	createReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions", createBody)
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()
	handler.ServeHTTP(createRec, createReq)

	if createRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", createRec.Code, createRec.Body.String())
	}

	listReq := httptest.NewRequest(http.MethodGet, "/v1/exceptions?active=true&environment=prod", nil)
	listRec := httptest.NewRecorder()
	handler.ServeHTTP(listRec, listReq)

	var listed exceptionsResponse
	if err := json.NewDecoder(listRec.Body).Decode(&listed); err != nil {
		t.Fatalf("decode list response: %v", err)
	}
	if len(listed.Exceptions) != 1 || listed.Exceptions[0].ExceptionID != "EX-2026-001" {
		t.Fatalf("unexpected exceptions %#v", listed)
	}

	validateReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/validate", bytes.NewBufferString(`{
	  "exception_id":"EX-2026-001",
	  "tenant_id":"acme",
	  "environment":"prod",
	  "namespace":"acme-prod"
	}`))
	validateReq.Header.Set("Content-Type", "application/json")
	validateRec := httptest.NewRecorder()
	handler.ServeHTTP(validateRec, validateReq)

	var validation audit.ExceptionValidationResult
	if err := json.NewDecoder(validateRec.Body).Decode(&validation); err != nil {
		t.Fatalf("decode validation response: %v", err)
	}
	if !validation.Valid || validation.Exception == nil || validation.Exception.ExceptionID != "EX-2026-001" {
		t.Fatalf("unexpected validation %#v", validation)
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/v1/exceptions/EX-2026-001", nil)
	deleteRec := httptest.NewRecorder()
	handler.ServeHTTP(deleteRec, deleteReq)

	if deleteRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", deleteRec.Code, deleteRec.Body.String())
	}

	reportReq := httptest.NewRequest(http.MethodGet, "/v1/reports/exceptions?environment=prod", nil)
	reportRec := httptest.NewRecorder()
	handler.ServeHTTP(reportRec, reportReq)

	var report audit.ExceptionReport
	if err := json.NewDecoder(reportRec.Body).Decode(&report); err != nil {
		t.Fatalf("decode exception report: %v", err)
	}
	if len(report.RecentInactive) != 1 {
		t.Fatalf("expected revoked exception in report, got %#v", report)
	}

	events, err := store.ListEvents(t.Context(), audit.EventFilter{Limit: 10})
	if err != nil {
		t.Fatalf("ListEvents() error = %v", err)
	}
	foundCreated := false
	foundRevoked := false
	for _, event := range events {
		if event.EventType == audit.EventTypeExceptionApproved && event.ExceptionID == "EX-2026-001" {
			foundCreated = true
		}
		if event.EventType == audit.EventTypeExceptionRevoked && event.ExceptionID == "EX-2026-001" {
			foundRevoked = true
		}
	}
	if !foundCreated || !foundRevoked {
		t.Fatalf("expected lifecycle audit events, got %#v", events)
	}
}

func TestValidateExceptionEndpointReturnsInvalidResult(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")
	if _, err := store.CreateException(t.Context(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-2026-002",
		ExceptionType: audit.ExceptionTypeDigestBypass,
		ImageDigest:   "sha256:abc123",
		Reason:        "digest bypass",
		TicketID:      "INC-2000",
		ApprovedBy:    "oncall@example.com",
		TTLHours:      1,
	}); err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/v1/exceptions/validate", bytes.NewBufferString(`{
	  "exception_id":"EX-2026-002",
	  "image_digest":"sha256:def456"
	}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var validation audit.ExceptionValidationResult
	if err := json.NewDecoder(rec.Body).Decode(&validation); err != nil {
		t.Fatalf("decode validation response: %v", err)
	}
	if validation.Valid || validation.Reason == "" {
		t.Fatalf("expected invalid validation result, got %#v", validation)
	}
}

func TestExceptionsReportEndpointFiltersByCVEID(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")
	exception, err := store.CreateException(t.Context(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-CVE-REPORT",
		ExceptionType: audit.ExceptionTypeCVEWhitelist,
		TenantID:      "acme",
		Environment:   "prod",
		Repo:          "my-org/acme-app",
		CVEID:         "CVE-2026-7777",
		Reason:        "temporary waiver",
		TicketID:      "SEC-7777",
		ApprovedBy:    "security@example.com",
		TTLHours:      1,
	})
	if err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}
	if _, err := store.Ingest(t.Context(), audit.Event{
		Component:           "policy-engine",
		EventType:           audit.EventTypeExceptionUsed,
		Decision:            audit.DecisionAllow,
		TenantID:            "acme",
		Environment:         "prod",
		Repo:                "my-org/acme-app",
		Digest:              "sha256:abc123",
		CVEID:               "CVE-2026-7777",
		IsException:         true,
		ExceptionID:         exception.ExceptionID,
		ExceptionType:       exception.ExceptionType,
		ExceptionReason:     exception.Reason,
		ExceptionTicketID:   exception.TicketID,
		ExceptionApprovedBy: exception.ApprovedBy,
		ExceptionExpiresAt:  &exception.ExpiresAt,
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/reports/exceptions?cve_id=CVE-2026-7777", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var report audit.ExceptionReport
	if err := json.NewDecoder(rec.Body).Decode(&report); err != nil {
		t.Fatalf("decode exception report: %v", err)
	}
	if len(report.Active) != 1 || report.Active[0].ExceptionID != "EX-CVE-REPORT" {
		t.Fatalf("unexpected active exceptions %#v", report.Active)
	}
	if len(report.RecentUsed) != 1 || report.RecentUsed[0].ExceptionID != "EX-CVE-REPORT" {
		t.Fatalf("unexpected used events %#v", report.RecentUsed)
	}
}

func TestAuthDisabledModeAllowsProtectedEndpoints(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeDisabled)
	store := audit.NewMemoryStore()
	if _, err := store.Ingest(t.Context(), audit.Event{
		Component: "deploy-gate",
		EventType: audit.EventTypeDeployGateDecision,
		Decision:  audit.DecisionAllow,
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/reports/events", nil)
	rec := httptest.NewRecorder()
	newHandler(store, "memory").ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestProtectedEndpointsRequireBearerToken(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	req := httptest.NewRequest(http.MethodGet, "/v1/reports/events", nil)
	rec := httptest.NewRecorder()
	newHandler(audit.NewMemoryStore(), "memory").ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestProtectedEndpointsRejectMalformedAndInvalidTokens(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())
	handler := newHandler(audit.NewMemoryStore(), "memory")

	tests := []struct {
		name   string
		header string
	}{
		{name: "malformed", header: "Token nope"},
		{name: "invalid", header: "Bearer nope"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/v1/reports/events", nil)
			req.Header.Set("Authorization", tc.header)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			if rec.Code != http.StatusUnauthorized {
				t.Fatalf("expected 401, got %d: %s", rec.Code, rec.Body.String())
			}
		})
	}
}

func TestViewerCanReadReportsAndExceptionsButCannotMutate(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	if _, err := store.Ingest(t.Context(), audit.Event{
		Component: "deploy-gate",
		EventType: audit.EventTypeDeployGateDecision,
		Decision:  audit.DecisionDeny,
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}
	handler := newHandler(store, "memory")

	readReq := httptest.NewRequest(http.MethodGet, "/v1/reports/events", nil)
	readReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	readRec := httptest.NewRecorder()
	handler.ServeHTTP(readRec, readReq)
	if readRec.Code != http.StatusOK {
		t.Fatalf("expected report read 200, got %d: %s", readRec.Code, readRec.Body.String())
	}

	listReq := httptest.NewRequest(http.MethodGet, "/v1/exceptions", nil)
	listReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	listRec := httptest.NewRecorder()
	handler.ServeHTTP(listRec, listReq)
	if listRec.Code != http.StatusOK {
		t.Fatalf("expected exception list 200, got %d: %s", listRec.Code, listRec.Body.String())
	}

	createReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions", bytes.NewBufferString(`{
	  "exception_id":"EX-2026-010",
	  "exception_type":"BREAK_GLASS",
	  "reason":"viewer should not create",
	  "ticket_id":"INC-10",
	  "approved_by":"security@example.com",
	  "ttl_hours":1
	}`))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	createRec := httptest.NewRecorder()
	handler.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", createRec.Code, createRec.Body.String())
	}
}

func TestOperatorCannotCreateOrRevokeExceptions(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	exception, err := store.CreateException(t.Context(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-2026-OP",
		ExceptionType: audit.ExceptionTypeBreakGlass,
		Environment:   "prod",
		Reason:        "operator test",
		TicketID:      "INC-OP",
		ApprovedBy:    "security@example.com",
		TTLHours:      1,
	})
	if err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}
	handler := newHandler(store, "memory")

	createReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions", bytes.NewBufferString(`{
	  "exception_id":"EX-2026-011",
	  "exception_type":"BREAK_GLASS",
	  "reason":"operator should not create",
	  "ticket_id":"INC-11",
	  "approved_by":"security@example.com",
	  "ttl_hours":1
	}`))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("Authorization", "Bearer operator-demo-token")
	createRec := httptest.NewRecorder()
	handler.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", createRec.Code, createRec.Body.String())
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/v1/exceptions/"+exception.ExceptionID, nil)
	deleteReq.Header.Set("Authorization", "Bearer operator-demo-token")
	deleteRec := httptest.NewRecorder()
	handler.ServeHTTP(deleteRec, deleteReq)
	if deleteRec.Code != http.StatusForbidden {
		t.Fatalf("expected revoke 403, got %d: %s", deleteRec.Code, deleteRec.Body.String())
	}
}

func TestSecurityAdminCanCreateAndRevokeExceptions(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())
	handler := newHandler(audit.NewMemoryStore(), "memory")

	createReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions", bytes.NewBufferString(`{
	  "exception_id":"EX-2026-ADMIN",
	  "exception_type":"BREAK_GLASS",
	  "tenant_id":"acme",
	  "reason":"admin create",
	  "ticket_id":"INC-ADMIN",
	  "approved_by":"security@example.com",
	  "ttl_hours":1
	}`))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	createRec := httptest.NewRecorder()
	handler.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", createRec.Code, createRec.Body.String())
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/v1/exceptions/EX-2026-ADMIN", nil)
	deleteReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	deleteRec := httptest.NewRecorder()
	handler.ServeHTTP(deleteRec, deleteReq)
	if deleteRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", deleteRec.Code, deleteRec.Body.String())
	}
}

func TestOperatorCanRequestButCannotApproveOrRejectExceptions(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")

	requestReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/request", bytes.NewBufferString(`{
	  "exception_id":"EX-2026-PENDING",
	  "exception_type":"BREAK_GLASS",
	  "tenant_id":"acme",
	  "environment":"prod",
	  "namespace":"acme-prod",
	  "reason":"operator request",
	  "ticket_id":"INC-PENDING",
	  "ttl_hours":1
	}`))
	requestReq.Header.Set("Content-Type", "application/json")
	requestReq.Header.Set("Authorization", "Bearer operator-demo-token")
	requestRec := httptest.NewRecorder()
	handler.ServeHTTP(requestRec, requestReq)
	if requestRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", requestRec.Code, requestRec.Body.String())
	}

	approveReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/EX-2026-PENDING/approve", bytes.NewBufferString(`{"reason":"approve"}`))
	approveReq.Header.Set("Content-Type", "application/json")
	approveReq.Header.Set("Authorization", "Bearer operator-demo-token")
	approveRec := httptest.NewRecorder()
	handler.ServeHTTP(approveRec, approveReq)
	if approveRec.Code != http.StatusForbidden {
		t.Fatalf("expected approve 403, got %d: %s", approveRec.Code, approveRec.Body.String())
	}

	rejectReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/EX-2026-PENDING/reject", bytes.NewBufferString(`{"reason":"reject"}`))
	rejectReq.Header.Set("Content-Type", "application/json")
	rejectReq.Header.Set("Authorization", "Bearer operator-demo-token")
	rejectRec := httptest.NewRecorder()
	handler.ServeHTTP(rejectRec, rejectReq)
	if rejectRec.Code != http.StatusForbidden {
		t.Fatalf("expected reject 403, got %d: %s", rejectRec.Code, rejectRec.Body.String())
	}
}

func TestSecurityAdminCanApproveAndRejectRequestedExceptions(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")

	requestReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/request", bytes.NewBufferString(`{
	  "exception_id":"EX-2026-PENDING-ADMIN",
	  "exception_type":"BREAK_GLASS",
	  "tenant_id":"acme",
	  "environment":"prod",
	  "namespace":"acme-prod",
	  "reason":"needs approval",
	  "ticket_id":"INC-PENDING-ADMIN",
	  "ttl_hours":1
	}`))
	requestReq.Header.Set("Content-Type", "application/json")
	requestReq.Header.Set("Authorization", "Bearer operator-demo-token")
	requestRec := httptest.NewRecorder()
	handler.ServeHTTP(requestRec, requestReq)
	if requestRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", requestRec.Code, requestRec.Body.String())
	}

	approveReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/EX-2026-PENDING-ADMIN/approve", bytes.NewBufferString(`{"reason":"approved for incident"}`))
	approveReq.Header.Set("Content-Type", "application/json")
	approveReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	approveRec := httptest.NewRecorder()
	handler.ServeHTTP(approveRec, approveReq)
	if approveRec.Code != http.StatusOK {
		t.Fatalf("expected approve 200, got %d: %s", approveRec.Code, approveRec.Body.String())
	}

	var approvedResponse exceptionActionResponse
	if err := json.NewDecoder(approveRec.Body).Decode(&approvedResponse); err != nil {
		t.Fatalf("decode approve response: %v", err)
	}
	if approvedResponse.Exception.Status != audit.ExceptionStatusApproved {
		t.Fatalf("unexpected approved response %#v", approvedResponse)
	}

	rejectCreateReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/request", bytes.NewBufferString(`{
	  "exception_id":"EX-2026-PENDING-REJECT",
	  "exception_type":"BREAK_GLASS",
	  "tenant_id":"acme",
	  "environment":"prod",
	  "namespace":"acme-prod",
	  "reason":"needs rejection",
	  "ticket_id":"INC-PENDING-REJECT",
	  "ttl_hours":1
	}`))
	rejectCreateReq.Header.Set("Content-Type", "application/json")
	rejectCreateReq.Header.Set("Authorization", "Bearer operator-demo-token")
	rejectCreateRec := httptest.NewRecorder()
	handler.ServeHTTP(rejectCreateRec, rejectCreateReq)
	if rejectCreateRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", rejectCreateRec.Code, rejectCreateRec.Body.String())
	}

	rejectReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/EX-2026-PENDING-REJECT/reject", bytes.NewBufferString(`{"reason":"missing evidence"}`))
	rejectReq.Header.Set("Content-Type", "application/json")
	rejectReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	rejectRec := httptest.NewRecorder()
	handler.ServeHTTP(rejectRec, rejectReq)
	if rejectRec.Code != http.StatusOK {
		t.Fatalf("expected reject 200, got %d: %s", rejectRec.Code, rejectRec.Body.String())
	}

	var rejectedResponse exceptionActionResponse
	if err := json.NewDecoder(rejectRec.Body).Decode(&rejectedResponse); err != nil {
		t.Fatalf("decode reject response: %v", err)
	}
	if rejectedResponse.Exception.Status != audit.ExceptionStatusRejected || rejectedResponse.Exception.RejectionReason != "missing evidence" {
		t.Fatalf("unexpected rejected response %#v", rejectedResponse)
	}
}

func TestServiceInternalCanValidateButCannotCreateExceptions(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	if _, err := store.CreateException(t.Context(), audit.ExceptionCreateRequest{
		ExceptionID:   "EX-2026-SVC",
		ExceptionType: audit.ExceptionTypeBreakGlass,
		Environment:   "prod",
		Reason:        "service validation",
		TicketID:      "INC-SVC",
		ApprovedBy:    "security@example.com",
		TTLHours:      1,
	}); err != nil {
		t.Fatalf("CreateException() error = %v", err)
	}
	handler := newHandler(store, "memory")

	validateReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/validate", bytes.NewBufferString(`{
	  "exception_id":"EX-2026-SVC"
	}`))
	validateReq.Header.Set("Content-Type", "application/json")
	validateReq.Header.Set("Authorization", "Bearer service-internal-demo-token")
	validateRec := httptest.NewRecorder()
	handler.ServeHTTP(validateRec, validateReq)
	if validateRec.Code != http.StatusOK {
		t.Fatalf("expected validation 200, got %d: %s", validateRec.Code, validateRec.Body.String())
	}

	createReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions", bytes.NewBufferString(`{
	  "exception_id":"EX-2026-SVC-CREATE",
	  "exception_type":"BREAK_GLASS",
	  "reason":"service should not create",
	  "ticket_id":"INC-SVC-CREATE",
	  "approved_by":"security@example.com",
	  "ttl_hours":1
	}`))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("Authorization", "Bearer service-internal-demo-token")
	createRec := httptest.NewRecorder()
	handler.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusForbidden {
		t.Fatalf("expected create 403, got %d: %s", createRec.Code, createRec.Body.String())
	}
}

func TestViewerCanReadAnalyticsButCannotRequestExceptions(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	if _, err := store.Ingest(t.Context(), audit.Event{
		Component:   "deploy-gate",
		EventType:   audit.EventTypeDeployGateDecision,
		Decision:    audit.DecisionDeny,
		TenantID:    "acme",
		Environment: "prod",
		Repo:        "my-org/acme-app",
		Reasons:     []string{"workflow mismatch"},
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}
	handler := newHandler(store, "memory")

	for _, path := range []string{
		"/v1/analytics/trends",
		"/v1/analytics/top-violators",
		"/v1/analytics/drift-stats",
	} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected analytics read 200 for %s, got %d: %s", path, rec.Code, rec.Body.String())
		}
	}

	requestReq := httptest.NewRequest(http.MethodPost, "/v1/exceptions/request", bytes.NewBufferString(`{
	  "exception_id":"EX-VIEWER-REQUEST",
	  "exception_type":"BREAK_GLASS",
	  "reason":"viewer cannot request",
	  "ticket_id":"INC-VIEWER",
	  "ttl_hours":1
	}`))
	requestReq.Header.Set("Content-Type", "application/json")
	requestReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	requestRec := httptest.NewRecorder()
	handler.ServeHTTP(requestRec, requestReq)
	if requestRec.Code != http.StatusForbidden {
		t.Fatalf("expected request 403, got %d: %s", requestRec.Code, requestRec.Body.String())
	}
}

func TestAuthMeReturnsCurrentPrincipal(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())
	handler := newHandler(audit.NewMemoryStore(), "memory")

	req := httptest.NewRequest(http.MethodGet, "/v1/auth/me", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response authInfoResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode auth me: %v", err)
	}
	if !response.Authenticated || response.AuthMode != auth.ModeStaticToken || response.Role != auth.RoleViewer || response.Subject != "demo-viewer" {
		t.Fatalf("unexpected auth me response %#v", response)
	}
}

func TestAuthMeRejectsServiceInternalRole(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())
	handler := newHandler(audit.NewMemoryStore(), "memory")

	req := httptest.NewRequest(http.MethodGet, "/v1/auth/me", nil)
	req.Header.Set("Authorization", "Bearer service-internal-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestLoadAuthConfigFromEnvRejectsInvalidConfig(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", "bogus")
	if _, err := loadAuthConfigFromEnv(); err == nil {
		t.Fatal("expected invalid mode error")
	}

	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", `[{"token":"dup","subject":"viewer","role":"viewer","token_id":"dup"},{"token":"other","subject":"admin","role":"security_admin","token_id":"dup"}]`)
	if _, err := loadAuthConfigFromEnv(); err == nil {
		t.Fatal("expected duplicate token_id error")
	}
}

func TestViewerCanReadSBOMAndVulnerabilityViewsButCannotMutate(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	if _, err := store.IngestSBOM(t.Context(), audit.SBOMIngestRequest{
		ImageDigest: "sha256:viewer-sbom",
		ImageRef:    "ghcr.io/example/viewer:1.0.0",
		SBOMFormat:  audit.SBOMFormatSPDXJSON,
		SBOM: []byte(`{
		  "packages": [{"name":"openssl","versionInfo":"3.0.14-r0","externalRefs":[{"referenceType":"purl","referenceLocator":"pkg:apk/alpine/openssl@3.0.14-r0"}]}]
		}`),
	}); err != nil {
		t.Fatalf("IngestSBOM() error = %v", err)
	}
	if _, err := store.RecordVulnerabilityScan(t.Context(), audit.VulnerabilityScanRequest{
		ImageDigest: "sha256:viewer-sbom",
		ImageRef:    "ghcr.io/example/viewer:1.0.0",
		Scanner:     "trivy",
		StartedAt:   time.Now().UTC(),
		CompletedAt: ptrTimeMain(time.Now().UTC()),
		Status:      audit.VulnerabilityScanStatusCompleted,
		Findings: []audit.VulnerabilityFindingInput{{
			CVEID:          "CVE-2026-4444",
			Severity:       "HIGH",
			PackageName:    "openssl",
			PackageVersion: "3.0.14-r0",
		}},
	}); err != nil {
		t.Fatalf("RecordVulnerabilityScan() error = %v", err)
	}
	handler := newHandler(store, "memory")

	for _, path := range []string{
		"/v1/sbom/components/search?component_name=openssl",
		"/v1/vulnerabilities/active?component_name=openssl",
		"/v1/vulnerabilities/blast-radius?cve_id=CVE-2026-4444",
		"/v1/vulnerabilities/timeline?image_digest=sha256:viewer-sbom&cve_id=CVE-2026-4444",
	} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200 for %s, got %d: %s", path, rec.Code, rec.Body.String())
		}
	}

	createReq := httptest.NewRequest(http.MethodPost, "/v1/vulnerabilities/decisions", bytes.NewBufferString(`{
	  "image_digest":"sha256:viewer-sbom",
	  "cve_id":"CVE-2026-4444",
	  "decision":"NOT_AFFECTED",
	  "justification":"viewer should not mutate"
	}`))
	createReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()
	handler.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", createRec.Code, createRec.Body.String())
	}
}

func TestSecurityAdminCanIngestSBOMAndManageVulnerabilityDecisions(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")

	ingestReq := httptest.NewRequest(http.MethodPost, "/v1/sbom/ingest", bytes.NewBufferString(`{
	  "image_digest":"sha256:sbom-admin",
	  "image_ref":"ghcr.io/example/admin:1.0.0",
	  "sbom_format":"spdx-json",
	  "sbom":{"packages":[{"name":"openssl","versionInfo":"3.0.14-r0"}]}
	}`))
	ingestReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	ingestReq.Header.Set("Content-Type", "application/json")
	ingestRec := httptest.NewRecorder()
	handler.ServeHTTP(ingestRec, ingestReq)
	if ingestRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", ingestRec.Code, ingestRec.Body.String())
	}

	if _, err := store.RecordVulnerabilityScan(t.Context(), audit.VulnerabilityScanRequest{
		ImageDigest: "sha256:sbom-admin",
		ImageRef:    "ghcr.io/example/admin:1.0.0",
		Scanner:     "trivy",
		StartedAt:   time.Now().UTC(),
		CompletedAt: ptrTimeMain(time.Now().UTC()),
		Status:      audit.VulnerabilityScanStatusCompleted,
		Findings: []audit.VulnerabilityFindingInput{{
			CVEID:          "CVE-2026-5555",
			Severity:       "MEDIUM",
			PackageName:    "openssl",
			PackageVersion: "3.0.14-r0",
		}},
	}); err != nil {
		t.Fatalf("RecordVulnerabilityScan() error = %v", err)
	}

	createDecisionReq := httptest.NewRequest(http.MethodPost, "/v1/vulnerabilities/decisions", bytes.NewBufferString(`{
	  "image_digest":"sha256:sbom-admin",
	  "cve_id":"CVE-2026-5555",
	  "decision":"ACCEPTED_RISK",
	  "justification":"accepted for maintenance window",
	  "ttl_hours":2
	}`))
	createDecisionReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	createDecisionReq.Header.Set("Content-Type", "application/json")
	createDecisionRec := httptest.NewRecorder()
	handler.ServeHTTP(createDecisionRec, createDecisionReq)
	if createDecisionRec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d: %s", createDecisionRec.Code, createDecisionRec.Body.String())
	}

	var created vulnerabilityDecisionActionResponse
	if err := json.NewDecoder(createDecisionRec.Body).Decode(&created); err != nil {
		t.Fatalf("decode decision response: %v", err)
	}
	if created.Decision.Decision != audit.VulnerabilityDecisionAcceptedRisk {
		t.Fatalf("unexpected decision %#v", created)
	}

	deactivateReq := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vulnerabilities/decisions/%d/deactivate", created.Decision.ID), nil)
	deactivateReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	deactivateRec := httptest.NewRecorder()
	handler.ServeHTTP(deactivateRec, deactivateReq)
	if deactivateRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", deactivateRec.Code, deactivateRec.Body.String())
	}
}

func TestServiceInternalCanTriggerRescanButCannotWriteHumanDecisionRoutes(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	if _, err := store.Ingest(t.Context(), audit.Event{
		Component:   "runtime-agent",
		EventType:   audit.EventTypeRuntimeDriftResult,
		Decision:    audit.DecisionAllow,
		TenantID:    "acme",
		Environment: "prod",
		Namespace:   "acme-prod",
		Workload:    "checkout",
		Repo:        "my-org/checkout",
		Image:       "ghcr.io/example/checkout:1.0.0",
		Digest:      "sha256:rescan1",
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	handler := newHandlerWithDeps(store, "memory", authConfig, &vulnOpsRuntime{
		config: internalvulnops.Config{
			Enabled:           true,
			SBOMIngestEnabled: true,
			ScanInterval:      time.Hour,
			Scanner:           internalvulnops.ScannerTrivy,
		},
		scanner: fakeScanner{
			result: internalvulnops.Result{
				ImageDigest: "sha256:rescan1",
				ImageRef:    "ghcr.io/example/checkout:1.0.0",
				Scanner:     internalvulnops.ScannerTrivy,
				StartedAt:   time.Now().UTC(),
				CompletedAt: time.Now().UTC(),
				Status:      audit.VulnerabilityScanStatusCompleted,
				Summary:     []byte(`{"critical":1,"high":0,"medium":0,"low":0,"unknown":0,"total":1}`),
				Findings: []audit.VulnerabilityFindingInput{{
					CVEID:          "CVE-2026-7777",
					Severity:       "CRITICAL",
					PackageName:    "glibc",
					PackageVersion: "2.39-r0",
				}},
			},
		},
	})

	rescanReq := httptest.NewRequest(http.MethodPost, "/v1/vulnerabilities/rescan", bytes.NewBufferString(`{"image_digest":"sha256:rescan1"}`))
	rescanReq.Header.Set("Authorization", "Bearer service-internal-demo-token")
	rescanReq.Header.Set("Content-Type", "application/json")
	rescanRec := httptest.NewRecorder()
	handler.ServeHTTP(rescanRec, rescanReq)
	if rescanRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rescanRec.Code, rescanRec.Body.String())
	}

	findings, err := store.ListActiveVulnerabilities(t.Context(), audit.VulnerabilityActiveFilter{ImageDigest: "sha256:rescan1", Limit: 10})
	if err != nil {
		t.Fatalf("ListActiveVulnerabilities() error = %v", err)
	}
	if len(findings) != 1 || findings[0].CVEID != "CVE-2026-7777" {
		t.Fatalf("unexpected findings after rescan %#v", findings)
	}

	decisionReq := httptest.NewRequest(http.MethodPost, "/v1/vulnerabilities/decisions", bytes.NewBufferString(`{
	  "image_digest":"sha256:rescan1",
	  "cve_id":"CVE-2026-7777",
	  "decision":"FIX_REQUIRED",
	  "justification":"service should not write decisions"
	}`))
	decisionReq.Header.Set("Authorization", "Bearer service-internal-demo-token")
	decisionReq.Header.Set("Content-Type", "application/json")
	decisionRec := httptest.NewRecorder()
	handler.ServeHTTP(decisionRec, decisionReq)
	if decisionRec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", decisionRec.Code, decisionRec.Body.String())
	}
}

type fakeScanner struct {
	result internalvulnops.Result
	err    error
}

func (s fakeScanner) ScanDigest(_ context.Context, _ audit.ActiveDigestRef) (internalvulnops.Result, error) {
	return s.result, s.err
}

func ptrTimeMain(value time.Time) *time.Time {
	return &value
}

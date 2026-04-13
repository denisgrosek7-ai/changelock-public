package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
)

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

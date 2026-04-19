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

func TestIncidentsEndpointBuildsCaseView(t *testing.T) {
	store := audit.NewMemoryStore()
	mustIngest := func(event audit.Event) {
		t.Helper()
		if _, err := store.Ingest(t.Context(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}

	mustIngest(audit.Event{
		RequestID:      "req-incident-1",
		Component:      "deploy-gate",
		EventType:      audit.EventTypeDeployGateDecision,
		Decision:       audit.DecisionDeny,
		TenantID:       "acme",
		Repo:           "repo-a",
		Environment:    "prod",
		Workload:       "api",
		Image:          "ghcr.io/acme/api@sha256:1111",
		Digest:         "sha256:1111",
		Reasons:        []string{"workflow mismatch"},
		PolicyBundleID: "bundle-a",
		CVEID:          "CVE-2026-1111",
	})
	mustIngest(audit.Event{
		RequestID:      "req-incident-2",
		Component:      "policy-engine",
		EventType:      audit.EventTypePolicyDecision,
		Decision:       audit.DecisionDeny,
		TenantID:       "acme",
		Repo:           "repo-a",
		Environment:    "prod",
		Namespace:      "prod-acme",
		Workload:       "api",
		Image:          "ghcr.io/acme/api@sha256:1111",
		Digest:         "sha256:1111",
		Reasons:        []string{"workflow mismatch"},
		PolicyBundleID: "bundle-a",
		ExceptionID:    "EX-2026-INCIDENT",
		VerifierSummary: &audit.VerifierSummary{
			SignatureValid:   false,
			AttestationValid: true,
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme&limit=25", nil)
	rec := httptest.NewRecorder()
	newHandler(store, "memory").ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response incidentsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(response.Incidents) != 1 {
		t.Fatalf("expected 1 incident, got %#v", response)
	}

	incident := response.Incidents[0]
	if !strings.HasPrefix(incident.ID, "INC-") || incident.IdentityKey == "" {
		t.Fatalf("unexpected incident id %#v", incident)
	}
	if incident.CategoryKey != "workflow-governance" || incident.State != incidentStateOpen {
		t.Fatalf("expected stable workflow incident metadata, got %#v", incident)
	}
	if incident.EventCount != 2 || incident.DenyCount != 2 {
		t.Fatalf("unexpected incident counts %#v", incident)
	}
	if len(incident.RemediationChecklist) == 0 || len(incident.Timeline) < 3 {
		t.Fatalf("expected checklist and timeline, got %#v", incident)
	}
	if len(incident.EvidencePack.RequestIDs) != 2 || len(incident.Events) != 2 {
		t.Fatalf("expected request IDs and events, got %#v", incident)
	}
	if len(incident.GovernanceImpacts) == 0 {
		t.Fatalf("expected governance impacts, got %#v", incident)
	}

	detailReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incident.ID+"?tenant_id=acme", nil)
	detailRec := httptest.NewRecorder()
	newHandler(store, "memory").ServeHTTP(detailRec, detailReq)
	if detailRec.Code != http.StatusOK {
		t.Fatalf("expected detail 200, got %d: %s", detailRec.Code, detailRec.Body.String())
	}
}

func TestIncidentLifecycleEndpointsPersistAndValidate(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")

	if _, err := store.Ingest(t.Context(), audit.Event{
		RequestID:   "req-incident-lifecycle-1",
		Component:   "deploy-gate",
		EventType:   audit.EventTypeDeployGateDecision,
		Decision:    audit.DecisionDeny,
		TenantID:    "acme",
		Repo:        "repo-lifecycle",
		Environment: "prod",
		Workload:    "api",
		Reasons:     []string{"workflow mismatch"},
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	listReq := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme", nil)
	listReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	listRec := httptest.NewRecorder()
	handler.ServeHTTP(listRec, listReq)
	if listRec.Code != http.StatusOK {
		t.Fatalf("expected list 200, got %d: %s", listRec.Code, listRec.Body.String())
	}

	var list incidentsResponse
	if err := json.NewDecoder(listRec.Body).Decode(&list); err != nil {
		t.Fatalf("decode incidents: %v", err)
	}
	if len(list.Incidents) != 1 {
		t.Fatalf("expected 1 incident, got %#v", list)
	}
	incidentID := list.Incidents[0].ID

	viewerAssignReq := httptest.NewRequest(http.MethodPost, "/v1/incidents/"+incidentID+"/assign?tenant_id=acme", bytes.NewBufferString(`{"owner":"secops","reason":"viewer denied"}`))
	viewerAssignReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	viewerAssignReq.Header.Set("Content-Type", "application/json")
	viewerAssignRec := httptest.NewRecorder()
	handler.ServeHTTP(viewerAssignRec, viewerAssignReq)
	if viewerAssignRec.Code != http.StatusForbidden {
		t.Fatalf("expected viewer assign 403, got %d: %s", viewerAssignRec.Code, viewerAssignRec.Body.String())
	}

	assignReq := httptest.NewRequest(http.MethodPost, "/v1/incidents/"+incidentID+"/assign?tenant_id=acme", bytes.NewBufferString(`{"owner":"secops","reason":"workflow drift ownership"}`))
	assignReq.Header.Set("Authorization", "Bearer operator-demo-token")
	assignReq.Header.Set("Content-Type", "application/json")
	assignRec := httptest.NewRecorder()
	handler.ServeHTTP(assignRec, assignReq)
	if assignRec.Code != http.StatusOK {
		t.Fatalf("expected assign 200, got %d: %s", assignRec.Code, assignRec.Body.String())
	}

	var assigned investigationIncident
	if err := json.NewDecoder(assignRec.Body).Decode(&assigned); err != nil {
		t.Fatalf("decode assigned incident: %v", err)
	}
	if assigned.Owner != "secops" || assigned.Assignment.Owner != "secops" {
		t.Fatalf("expected owner assignment, got %#v", assigned)
	}

	ackReq := httptest.NewRequest(http.MethodPost, "/v1/incidents/"+incidentID+"/acknowledge?tenant_id=acme", bytes.NewBufferString(`{"summary":"triage started"}`))
	ackReq.Header.Set("Authorization", "Bearer operator-demo-token")
	ackReq.Header.Set("Content-Type", "application/json")
	ackRec := httptest.NewRecorder()
	handler.ServeHTTP(ackRec, ackReq)
	if ackRec.Code != http.StatusOK {
		t.Fatalf("expected acknowledge 200, got %d: %s", ackRec.Code, ackRec.Body.String())
	}

	var acknowledged investigationIncident
	if err := json.NewDecoder(ackRec.Body).Decode(&acknowledged); err != nil {
		t.Fatalf("decode acknowledged incident: %v", err)
	}
	if acknowledged.State != incidentStateAcknowledged {
		t.Fatalf("expected acknowledged state, got %#v", acknowledged)
	}

	watchReq := httptest.NewRequest(http.MethodPost, "/v1/incidents/"+incidentID+"/watch?tenant_id=acme", bytes.NewBufferString(`{"summary":"watch for repeat signal"}`))
	watchReq.Header.Set("Authorization", "Bearer operator-demo-token")
	watchReq.Header.Set("Content-Type", "application/json")
	watchRec := httptest.NewRecorder()
	handler.ServeHTTP(watchRec, watchReq)
	if watchRec.Code != http.StatusOK {
		t.Fatalf("expected watch 200, got %d: %s", watchRec.Code, watchRec.Body.String())
	}

	var watching investigationIncident
	if err := json.NewDecoder(watchRec.Body).Decode(&watching); err != nil {
		t.Fatalf("decode watching incident: %v", err)
	}
	if watching.State != incidentStateWatching {
		t.Fatalf("expected watching state, got %#v", watching)
	}

	noteReq := httptest.NewRequest(http.MethodPost, "/v1/incidents/"+incidentID+"/notes?tenant_id=acme", bytes.NewBufferString(`{"note":"waiting for signer policy owner confirmation"}`))
	noteReq.Header.Set("Authorization", "Bearer operator-demo-token")
	noteReq.Header.Set("Content-Type", "application/json")
	noteRec := httptest.NewRecorder()
	handler.ServeHTTP(noteRec, noteReq)
	if noteRec.Code != http.StatusOK {
		t.Fatalf("expected note 200, got %d: %s", noteRec.Code, noteRec.Body.String())
	}

	resolveForbiddenReq := httptest.NewRequest(http.MethodPost, "/v1/incidents/"+incidentID+"/resolve?tenant_id=acme", bytes.NewBufferString(`{"resolution_type":"fixed","resolution_summary":"policy updated"}`))
	resolveForbiddenReq.Header.Set("Authorization", "Bearer operator-demo-token")
	resolveForbiddenReq.Header.Set("Content-Type", "application/json")
	resolveForbiddenRec := httptest.NewRecorder()
	handler.ServeHTTP(resolveForbiddenRec, resolveForbiddenReq)
	if resolveForbiddenRec.Code != http.StatusForbidden {
		t.Fatalf("expected operator resolve 403, got %d: %s", resolveForbiddenRec.Code, resolveForbiddenRec.Body.String())
	}

	resolveReq := httptest.NewRequest(http.MethodPost, "/v1/incidents/"+incidentID+"/resolve?tenant_id=acme", bytes.NewBufferString(`{
	  "resolution_type":"fixed",
	  "resolution_summary":"workflow trust rule updated",
	  "resolution_details":"trusted workflow ref rotated to the new release workflow",
	  "follow_up_required":true,
	  "resolution_refs":["decision:sha256:test"]
	}`))
	resolveReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	resolveReq.Header.Set("Content-Type", "application/json")
	resolveRec := httptest.NewRecorder()
	handler.ServeHTTP(resolveRec, resolveReq)
	if resolveRec.Code != http.StatusOK {
		t.Fatalf("expected resolve 200, got %d: %s", resolveRec.Code, resolveRec.Body.String())
	}

	var resolved investigationIncident
	if err := json.NewDecoder(resolveRec.Body).Decode(&resolved); err != nil {
		t.Fatalf("decode resolved incident: %v", err)
	}
	if resolved.State != incidentStateResolved || resolved.Resolution.Type != "fixed" {
		t.Fatalf("expected resolved incident, got %#v", resolved)
	}
	if !resolved.Resolution.FollowUpRequired || resolved.NewActivityDetected {
		t.Fatalf("expected structured resolution and no reopen signal, got %#v", resolved)
	}

	persistedHandler := newHandler(store, "memory")
	persistedReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"?tenant_id=acme", nil)
	persistedReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	persistedRec := httptest.NewRecorder()
	persistedHandler.ServeHTTP(persistedRec, persistedReq)
	if persistedRec.Code != http.StatusOK {
		t.Fatalf("expected persisted detail 200, got %d: %s", persistedRec.Code, persistedRec.Body.String())
	}
	var persisted investigationIncident
	if err := json.NewDecoder(persistedRec.Body).Decode(&persisted); err != nil {
		t.Fatalf("decode persisted incident: %v", err)
	}
	if persisted.State != incidentStateResolved || len(persisted.Notes) == 0 || len(persisted.History) == 0 {
		t.Fatalf("expected persisted overlay after handler rebuild, got %#v", persisted)
	}

	reopenReq := httptest.NewRequest(http.MethodPost, "/v1/incidents/"+incidentID+"/reopen?tenant_id=acme", bytes.NewBufferString(`{"reason":"new deploy signal requires re-review"}`))
	reopenReq.Header.Set("Authorization", "Bearer operator-demo-token")
	reopenReq.Header.Set("Content-Type", "application/json")
	reopenRec := httptest.NewRecorder()
	handler.ServeHTTP(reopenRec, reopenReq)
	if reopenRec.Code != http.StatusOK {
		t.Fatalf("expected reopen 200, got %d: %s", reopenRec.Code, reopenRec.Body.String())
	}

	var reopened investigationIncident
	if err := json.NewDecoder(reopenRec.Body).Decode(&reopened); err != nil {
		t.Fatalf("decode reopened incident: %v", err)
	}
	if reopened.State != incidentStateReopened {
		t.Fatalf("expected reopened state, got %#v", reopened)
	}

	timelineReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"/timeline?tenant_id=acme", nil)
	timelineReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	timelineRec := httptest.NewRecorder()
	handler.ServeHTTP(timelineRec, timelineReq)
	if timelineRec.Code != http.StatusOK {
		t.Fatalf("expected timeline 200, got %d: %s", timelineRec.Code, timelineRec.Body.String())
	}

	var timeline incidentTimelineResponse
	if err := json.NewDecoder(timelineRec.Body).Decode(&timeline); err != nil {
		t.Fatalf("decode timeline: %v", err)
	}
	if len(timeline.Timeline) < 5 {
		t.Fatalf("expected lifecycle timeline entries, got %#v", timeline)
	}

	historyReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"/history?tenant_id=acme", nil)
	historyReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	historyRec := httptest.NewRecorder()
	handler.ServeHTTP(historyRec, historyReq)
	if historyRec.Code != http.StatusOK {
		t.Fatalf("expected history 200, got %d: %s", historyRec.Code, historyRec.Body.String())
	}
	var history incidentHistoryResponse
	if err := json.NewDecoder(historyRec.Body).Decode(&history); err != nil {
		t.Fatalf("decode history: %v", err)
	}
	if len(history.History) < 5 {
		t.Fatalf("expected lifecycle history entries, got %#v", history)
	}
}

func TestResolvedIncidentShowsNewActivityWithoutImplicitReopen(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")
	baseTime := time.Now().UTC()

	if _, err := store.Ingest(t.Context(), audit.Event{
		RequestID:   "req-incident-new-activity-1",
		Timestamp:   baseTime,
		Component:   "deploy-gate",
		EventType:   audit.EventTypeDeployGateDecision,
		Decision:    audit.DecisionDeny,
		TenantID:    "acme",
		Repo:        "repo-new-activity",
		Environment: "prod",
		Workload:    "api",
		Reasons:     []string{"workflow mismatch"},
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	listReq := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme", nil)
	listReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	listRec := httptest.NewRecorder()
	handler.ServeHTTP(listRec, listReq)
	if listRec.Code != http.StatusOK {
		t.Fatalf("expected list 200, got %d: %s", listRec.Code, listRec.Body.String())
	}

	var list incidentsResponse
	if err := json.NewDecoder(listRec.Body).Decode(&list); err != nil {
		t.Fatalf("decode incidents: %v", err)
	}
	if len(list.Incidents) != 1 {
		t.Fatalf("expected 1 incident, got %#v", list)
	}
	incidentID := list.Incidents[0].ID

	resolveReq := httptest.NewRequest(http.MethodPost, "/v1/incidents/"+incidentID+"/resolve?tenant_id=acme", bytes.NewBufferString(`{
	  "resolution_type":"fixed",
	  "resolution_summary":"initial policy update"
	}`))
	resolveReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	resolveReq.Header.Set("Content-Type", "application/json")
	resolveRec := httptest.NewRecorder()
	handler.ServeHTTP(resolveRec, resolveReq)
	if resolveRec.Code != http.StatusOK {
		t.Fatalf("expected resolve 200, got %d: %s", resolveRec.Code, resolveRec.Body.String())
	}

	if _, err := store.Ingest(t.Context(), audit.Event{
		RequestID:   "req-incident-new-activity-2",
		Timestamp:   baseTime.Add(time.Minute),
		Component:   "deploy-gate",
		EventType:   audit.EventTypeDeployGateDecision,
		Decision:    audit.DecisionDeny,
		TenantID:    "acme",
		Repo:        "repo-new-activity",
		Environment: "prod",
		Workload:    "api",
		Reasons:     []string{"workflow mismatch"},
	}); err != nil {
		t.Fatalf("Ingest() second event error = %v", err)
	}

	detailReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"?tenant_id=acme", nil)
	detailReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	detailRec := httptest.NewRecorder()
	handler.ServeHTTP(detailRec, detailReq)
	if detailRec.Code != http.StatusOK {
		t.Fatalf("expected detail 200, got %d: %s", detailRec.Code, detailRec.Body.String())
	}

	var detail investigationIncident
	if err := json.NewDecoder(detailRec.Body).Decode(&detail); err != nil {
		t.Fatalf("decode incident detail: %v", err)
	}
	if detail.State != incidentStateResolved || !detail.NewActivityDetected {
		t.Fatalf("expected resolved incident with new activity flag, got %#v", detail)
	}
}

func TestIncidentExportAndMetricDrillDown(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")

	mustIngest := func(event audit.Event) {
		t.Helper()
		if _, err := store.Ingest(t.Context(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}

	mustIngest(audit.Event{
		RequestID:      "req-export-1",
		Component:      "deploy-gate",
		EventType:      audit.EventTypeDeployGateDecision,
		Decision:       audit.DecisionDeny,
		TenantID:       "acme",
		Repo:           "repo-export",
		Environment:    "prod",
		Workload:       "api",
		Image:          "ghcr.io/acme/api@sha256:aaaa",
		Digest:         "sha256:aaaa",
		Reasons:        []string{"workflow mismatch"},
		PolicyBundleID: "bundle-export",
	})
	mustIngest(audit.Event{
		RequestID:   "req-export-2",
		Component:   "policy-engine",
		EventType:   audit.EventTypePolicyDecision,
		Decision:    audit.DecisionDeny,
		TenantID:    "acme",
		Repo:        "repo-export",
		Environment: "prod",
		Namespace:   "prod-acme",
		Workload:    "api",
		Reasons:     []string{"workflow mismatch"},
	})

	listReq := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme", nil)
	listRec := httptest.NewRecorder()
	handler.ServeHTTP(listRec, listReq)
	if listRec.Code != http.StatusOK {
		t.Fatalf("expected list 200, got %d: %s", listRec.Code, listRec.Body.String())
	}

	var list incidentsResponse
	if err := json.NewDecoder(listRec.Body).Decode(&list); err != nil {
		t.Fatalf("decode incidents: %v", err)
	}
	if len(list.Incidents) != 1 {
		t.Fatalf("expected 1 incident, got %#v", list)
	}
	incidentID := list.Incidents[0].ID

	ackReq := httptest.NewRequest(http.MethodPost, "/v1/incidents/"+incidentID+"/acknowledge?tenant_id=acme", bytes.NewBufferString(`{"summary":"triage started"}`))
	ackReq.Header.Set("Content-Type", "application/json")
	ackRec := httptest.NewRecorder()
	handler.ServeHTTP(ackRec, ackReq)
	if ackRec.Code != http.StatusOK {
		t.Fatalf("expected acknowledge 200, got %d: %s", ackRec.Code, ackRec.Body.String())
	}

	detailReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"?tenant_id=acme", nil)
	detailRec := httptest.NewRecorder()
	handler.ServeHTTP(detailRec, detailReq)
	if detailRec.Code != http.StatusOK {
		t.Fatalf("expected detail 200, got %d: %s", detailRec.Code, detailRec.Body.String())
	}

	var detail investigationIncident
	if err := json.NewDecoder(detailRec.Body).Decode(&detail); err != nil {
		t.Fatalf("decode detail: %v", err)
	}
	if len(detail.MetricLinks) == 0 || detail.MetricLinks[0].MetricKey != "workflow-governance" {
		t.Fatalf("expected metric links on incident detail, got %#v", detail.MetricLinks)
	}

	exportReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"/export?tenant_id=acme", nil)
	exportRec := httptest.NewRecorder()
	handler.ServeHTTP(exportRec, exportReq)
	if exportRec.Code != http.StatusOK {
		t.Fatalf("expected export 200, got %d: %s", exportRec.Code, exportRec.Body.String())
	}

	var exportPayload incidentExportResponse
	if err := json.NewDecoder(exportRec.Body).Decode(&exportPayload); err != nil {
		t.Fatalf("decode export payload: %v", err)
	}
	if exportPayload.IncidentID != incidentID || exportPayload.State != incidentStateAcknowledged {
		t.Fatalf("unexpected export payload %#v", exportPayload)
	}
	if len(exportPayload.MetricLinks) == 0 || len(exportPayload.RelatedEventRefs) != 2 {
		t.Fatalf("expected metric links and event refs in export payload, got %#v", exportPayload)
	}
	if len(exportPayload.History) == 0 || !containsString(exportPayload.ScorecardRefs, "workflow-governance") {
		t.Fatalf("expected history and scorecard refs in export payload, got %#v", exportPayload)
	}

	auditorReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"/export?tenant_id=acme&audience=auditor_safe", nil)
	auditorRec := httptest.NewRecorder()
	handler.ServeHTTP(auditorRec, auditorReq)
	if auditorRec.Code != http.StatusOK {
		t.Fatalf("expected auditor export 200, got %d: %s", auditorRec.Code, auditorRec.Body.String())
	}

	var auditorPayload incidentExportResponse
	if err := json.NewDecoder(auditorRec.Body).Decode(&auditorPayload); err != nil {
		t.Fatalf("decode auditor export: %v", err)
	}
	if !auditorPayload.Redacted || auditorPayload.Audience != incidentAudienceAuditorSafe {
		t.Fatalf("expected redacted auditor payload, got %#v", auditorPayload)
	}
	if auditorPayload.Repository == "repo-export" || auditorPayload.Owner != "" || len(auditorPayload.Notes) != 0 {
		t.Fatalf("expected masked repo/owner and removed notes, got %#v", auditorPayload)
	}
	if len(auditorPayload.RedactionSummary) == 0 || len(auditorPayload.RelatedEventRefs) == 0 {
		t.Fatalf("expected redaction summary and masked event refs, got %#v", auditorPayload)
	}

	customerReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"/export?tenant_id=acme&audience=customer_safe", nil)
	customerRec := httptest.NewRecorder()
	handler.ServeHTTP(customerRec, customerReq)
	if customerRec.Code != http.StatusOK {
		t.Fatalf("expected customer export 200, got %d: %s", customerRec.Code, customerRec.Body.String())
	}

	var customerPayload incidentExportResponse
	if err := json.NewDecoder(customerRec.Body).Decode(&customerPayload); err != nil {
		t.Fatalf("decode customer export: %v", err)
	}
	if !customerPayload.Redacted || customerPayload.Audience != incidentAudienceCustomerSafe {
		t.Fatalf("expected redacted customer payload, got %#v", customerPayload)
	}
	if customerPayload.Repository != "" || customerPayload.ScopeRef != "" || len(customerPayload.RelatedEventRefs) != 0 {
		t.Fatalf("expected customer-safe scope and event refs to be stripped, got %#v", customerPayload)
	}
	if len(customerPayload.GuidanceRefs) != 0 || len(customerPayload.RedactionSummary) == 0 {
		t.Fatalf("expected customer-safe guidance stripping and redaction summary, got %#v", customerPayload)
	}

	metricReq := httptest.NewRequest(http.MethodGet, "/v1/scorecard/metrics/workflow-governance/incidents?tenant_id=acme", nil)
	metricRec := httptest.NewRecorder()
	handler.ServeHTTP(metricRec, metricReq)
	if metricRec.Code != http.StatusOK {
		t.Fatalf("expected metric drill-down 200, got %d: %s", metricRec.Code, metricRec.Body.String())
	}

	var metricResponse metricIncidentsResponse
	if err := json.NewDecoder(metricRec.Body).Decode(&metricResponse); err != nil {
		t.Fatalf("decode metric drill-down: %v", err)
	}
	if metricResponse.MetricKey != "workflow-governance" || metricResponse.MetricLabel == "" {
		t.Fatalf("unexpected metric drill-down header %#v", metricResponse)
	}
	if len(metricResponse.Incidents) != 1 || metricResponse.Incidents[0].ID != incidentID {
		t.Fatalf("expected linked incident in metric drill-down, got %#v", metricResponse)
	}
	if len(metricResponse.Limitations) == 0 {
		t.Fatalf("expected drill-down limitations, got %#v", metricResponse)
	}

	incidentDefenseReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"/defense-gaps?tenant_id=acme", nil)
	incidentDefenseRec := httptest.NewRecorder()
	handler.ServeHTTP(incidentDefenseRec, incidentDefenseReq)
	if incidentDefenseRec.Code != http.StatusOK {
		t.Fatalf("expected incident defense-gap 200, got %d: %s", incidentDefenseRec.Code, incidentDefenseRec.Body.String())
	}

	var incidentAssessment defenseGapAssessment
	if err := json.NewDecoder(incidentDefenseRec.Body).Decode(&incidentAssessment); err != nil {
		t.Fatalf("decode incident defense-gap assessment: %v", err)
	}
	if incidentAssessment.SubjectType != "incident" || incidentAssessment.SubjectRef != incidentID || len(incidentAssessment.DefenseGaps) == 0 {
		t.Fatalf("unexpected incident defense-gap assessment %#v", incidentAssessment)
	}
	if !incidentAssessment.AdvisoryOnly || len(incidentAssessment.Limitations) == 0 {
		t.Fatalf("expected advisory-only incident assessment with limitations, got %#v", incidentAssessment)
	}

	metricDefenseReq := httptest.NewRequest(http.MethodGet, "/v1/scorecard/metrics/workflow-governance/defense-gaps?tenant_id=acme", nil)
	metricDefenseRec := httptest.NewRecorder()
	handler.ServeHTTP(metricDefenseRec, metricDefenseReq)
	if metricDefenseRec.Code != http.StatusOK {
		t.Fatalf("expected metric defense-gap 200, got %d: %s", metricDefenseRec.Code, metricDefenseRec.Body.String())
	}

	var metricAssessment defenseGapAssessment
	if err := json.NewDecoder(metricDefenseRec.Body).Decode(&metricAssessment); err != nil {
		t.Fatalf("decode metric defense-gap assessment: %v", err)
	}
	if metricAssessment.SubjectType != "metric" || metricAssessment.SubjectRef != "workflow-governance" || len(metricAssessment.DefenseGaps) == 0 {
		t.Fatalf("unexpected metric defense-gap assessment %#v", metricAssessment)
	}
	if !containsString(metricAssessment.DefenseGaps[0].RelatedIncidentRefs, incidentID) {
		t.Fatalf("expected metric defense-gap to reference incident %s, got %#v", incidentID, metricAssessment)
	}

	aiDefenseReq := httptest.NewRequest(http.MethodGet, "/v1/ai/defense-gap-assessments?tenant_id=acme&incident_id="+incidentID, nil)
	aiDefenseRec := httptest.NewRecorder()
	handler.ServeHTTP(aiDefenseRec, aiDefenseReq)
	if aiDefenseRec.Code != http.StatusOK {
		t.Fatalf("expected ai defense-gap 200, got %d: %s", aiDefenseRec.Code, aiDefenseRec.Body.String())
	}

	incidentReplayReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"/policy-replay?tenant_id=acme", nil)
	incidentReplayRec := httptest.NewRecorder()
	handler.ServeHTTP(incidentReplayRec, incidentReplayReq)
	if incidentReplayRec.Code != http.StatusOK {
		t.Fatalf("expected incident replay 200, got %d: %s", incidentReplayRec.Code, incidentReplayRec.Body.String())
	}

	var incidentReplay policyReplayAssessment
	if err := json.NewDecoder(incidentReplayRec.Body).Decode(&incidentReplay); err != nil {
		t.Fatalf("decode incident replay: %v", err)
	}
	if incidentReplay.SubjectType != "incident" || incidentReplay.SubjectRef != incidentID || len(incidentReplay.ReplayResults) == 0 || len(incidentReplay.CoverageGaps) == 0 {
		t.Fatalf("unexpected incident replay %#v", incidentReplay)
	}
	if !incidentReplay.AdvisoryOnly || !incidentReplay.ShadowMode {
		t.Fatalf("expected advisory shadow-mode replay, got %#v", incidentReplay)
	}

	metricReplayReq := httptest.NewRequest(http.MethodGet, "/v1/scorecard/metrics/workflow-governance/policy-replay?tenant_id=acme", nil)
	metricReplayRec := httptest.NewRecorder()
	handler.ServeHTTP(metricReplayRec, metricReplayReq)
	if metricReplayRec.Code != http.StatusOK {
		t.Fatalf("expected metric replay 200, got %d: %s", metricReplayRec.Code, metricReplayRec.Body.String())
	}

	var metricReplay policyReplayAssessment
	if err := json.NewDecoder(metricReplayRec.Body).Decode(&metricReplay); err != nil {
		t.Fatalf("decode metric replay: %v", err)
	}
	if metricReplay.SubjectType != "metric" || metricReplay.SubjectRef != "workflow-governance" || len(metricReplay.ReplayResults) == 0 {
		t.Fatalf("unexpected metric replay %#v", metricReplay)
	}

	scopeReplayReq := httptest.NewRequest(http.MethodGet, "/v1/ai/policy-replay?tenant_id=acme", nil)
	scopeReplayRec := httptest.NewRecorder()
	handler.ServeHTTP(scopeReplayRec, scopeReplayReq)
	if scopeReplayRec.Code != http.StatusOK {
		t.Fatalf("expected scope replay 200, got %d: %s", scopeReplayRec.Code, scopeReplayRec.Body.String())
	}

	var scopeReplay policyReplayAssessment
	if err := json.NewDecoder(scopeReplayRec.Body).Decode(&scopeReplay); err != nil {
		t.Fatalf("decode scope replay: %v", err)
	}
	if scopeReplay.SubjectType != "scope" || len(scopeReplay.CoverageGaps) == 0 || scopeReplay.BlastRadius.IncidentCount == 0 {
		t.Fatalf("unexpected scope replay %#v", scopeReplay)
	}

	systemicReq := httptest.NewRequest(http.MethodGet, "/v1/ai/systemic-weaknesses?tenant_id=acme", nil)
	systemicRec := httptest.NewRecorder()
	handler.ServeHTTP(systemicRec, systemicReq)
	if systemicRec.Code != http.StatusOK {
		t.Fatalf("expected systemic weaknesses 200, got %d: %s", systemicRec.Code, systemicRec.Body.String())
	}

	var systemicResponse systemicWeaknessResponse
	if err := json.NewDecoder(systemicRec.Body).Decode(&systemicResponse); err != nil {
		t.Fatalf("decode systemic weaknesses: %v", err)
	}
	if !systemicResponse.AdvisoryOnly || len(systemicResponse.Weaknesses) == 0 {
		t.Fatalf("unexpected systemic weakness response %#v", systemicResponse)
	}

	metricSystemicReq := httptest.NewRequest(http.MethodGet, "/v1/scorecard/metrics/workflow-governance/systemic-weaknesses?tenant_id=acme", nil)
	metricSystemicRec := httptest.NewRecorder()
	handler.ServeHTTP(metricSystemicRec, metricSystemicReq)
	if metricSystemicRec.Code != http.StatusOK {
		t.Fatalf("expected metric systemic weaknesses 200, got %d: %s", metricSystemicRec.Code, metricSystemicRec.Body.String())
	}
}

func TestIncidentPackageEndpointBuildsDerivedBundle(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")

	mustIngest := func(event audit.Event) {
		t.Helper()
		if _, err := store.Ingest(t.Context(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}

	mustIngest(audit.Event{
		RequestID:      "req-package-1",
		Component:      "deploy-gate",
		EventType:      audit.EventTypeDeployGateDecision,
		Decision:       audit.DecisionDeny,
		TenantID:       "acme",
		Repo:           "repo-package-a",
		Environment:    "prod",
		Workload:       "api",
		Digest:         "sha256:package-a",
		Reasons:        []string{"workflow mismatch"},
		PolicyBundleID: "bundle-package-a",
	})
	mustIngest(audit.Event{
		RequestID:   "req-package-2",
		Component:   "runtime-agent",
		EventType:   audit.EventTypeRuntimeDriftResult,
		Decision:    audit.DecisionDeny,
		TenantID:    "acme",
		Repo:        "repo-package-b",
		Environment: "prod",
		Workload:    "worker",
		Image:       "ghcr.io/acme/worker@sha256:bbbb",
		Digest:      "sha256:bbbb",
		DriftResult: "image_drift",
		Reasons:     []string{"image drift"},
	})

	listReq := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme", nil)
	listRec := httptest.NewRecorder()
	handler.ServeHTTP(listRec, listReq)
	if listRec.Code != http.StatusOK {
		t.Fatalf("expected incident list 200, got %d: %s", listRec.Code, listRec.Body.String())
	}

	var list incidentsResponse
	if err := json.NewDecoder(listRec.Body).Decode(&list); err != nil {
		t.Fatalf("decode incident list: %v", err)
	}
	if len(list.Incidents) != 2 {
		t.Fatalf("expected 2 incidents, got %#v", list)
	}

	packageReq := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"/v1/incidents/package?tenant_id=acme&audience=auditor_safe&incident_id=%s&incident_id=%s",
			list.Incidents[0].ID,
			list.Incidents[1].ID,
		),
		nil,
	)
	packageRec := httptest.NewRecorder()
	handler.ServeHTTP(packageRec, packageReq)
	if packageRec.Code != http.StatusOK {
		t.Fatalf("expected package 200, got %d: %s", packageRec.Code, packageRec.Body.String())
	}

	var packageResponse incidentPackageResponse
	if err := json.NewDecoder(packageRec.Body).Decode(&packageResponse); err != nil {
		t.Fatalf("decode package response: %v", err)
	}
	if packageResponse.SelectionMode != "explicit" || packageResponse.Audience != incidentAudienceAuditorSafe || !packageResponse.Redacted {
		t.Fatalf("unexpected package header %#v", packageResponse)
	}
	if packageResponse.IncidentCount != 2 || len(packageResponse.Incidents) != 2 || len(packageResponse.IncidentRefs) != 2 {
		t.Fatalf("expected two included incidents, got %#v", packageResponse)
	}
	totalSeverityCount := 0
	for _, count := range packageResponse.Aggregate.BySeverity {
		totalSeverityCount += count
	}
	if totalSeverityCount != 2 {
		t.Fatalf("expected severity aggregate counts for both incidents, got %#v", packageResponse.Aggregate)
	}
	if len(packageResponse.RedactionSummary) == 0 || len(packageResponse.Limitations) == 0 {
		t.Fatalf("expected redaction summary and limitations, got %#v", packageResponse)
	}

	queryReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/package?tenant_id=acme&audience=internal", nil)
	queryRec := httptest.NewRecorder()
	handler.ServeHTTP(queryRec, queryReq)
	if queryRec.Code != http.StatusOK {
		t.Fatalf("expected query-derived package 200, got %d: %s", queryRec.Code, queryRec.Body.String())
	}

	var queryResponse incidentPackageResponse
	if err := json.NewDecoder(queryRec.Body).Decode(&queryResponse); err != nil {
		t.Fatalf("decode query-derived package: %v", err)
	}
	if queryResponse.SelectionMode != "query_derived" || queryResponse.IncidentCount != 2 {
		t.Fatalf("unexpected query-derived package %#v", queryResponse)
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

	incidentsReq := httptest.NewRequest(http.MethodGet, "/v1/incidents", nil)
	incidentsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	incidentsRec := httptest.NewRecorder()
	handler.ServeHTTP(incidentsRec, incidentsReq)
	if incidentsRec.Code != http.StatusOK {
		t.Fatalf("expected incidents read 200, got %d: %s", incidentsRec.Code, incidentsRec.Body.String())
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

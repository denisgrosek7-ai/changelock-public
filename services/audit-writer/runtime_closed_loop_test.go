package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
)

func TestRuntimeClosedLoopEndpoints(t *testing.T) {
	store := audit.NewMemoryStore()
	_, err := store.Ingest(t.Context(), audit.Event{
		RequestID:             "rt-1",
		Component:             "runtime-agent",
		EventType:             audit.EventTypeRuntimeActiveStateObserved,
		TenantID:              "acme",
		ClusterID:             "prod-eu",
		Namespace:             "acme-prod",
		WorkloadKind:          "Deployment",
		Workload:              "booking-api",
		Digest:                "sha256:active",
		Decision:              audit.DecisionDeny,
		ReconciliationStatus:  "quarantined",
		QuarantineType:        "vex",
		QuarantineReason:      "net actionable critical vulnerability requires containment",
		Timestamp:             time.Now().UTC(),
	})
	if err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	handler := newHandler(store, "memory")

	activeReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/active-state?tenant_id=acme", nil)
	activeRec := httptest.NewRecorder()
	handler.ServeHTTP(activeRec, activeReq)
	if activeRec.Code != http.StatusOK {
		t.Fatalf("expected 200 for active-state, got %d: %s", activeRec.Code, activeRec.Body.String())
	}

	var active struct {
		Items []audit.RuntimeActiveStateView `json:"items"`
	}
	if err := json.NewDecoder(activeRec.Body).Decode(&active); err != nil {
		t.Fatalf("decode active-state: %v", err)
	}
	if len(active.Items) != 1 || active.Items[0].QuarantineType != "vex" {
		t.Fatalf("unexpected active-state response %#v", active)
	}

	statusReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/closed-loop/status?tenant_id=acme", nil)
	statusRec := httptest.NewRecorder()
	handler.ServeHTTP(statusRec, statusReq)
	if statusRec.Code != http.StatusOK {
		t.Fatalf("expected 200 for closed-loop status, got %d: %s", statusRec.Code, statusRec.Body.String())
	}

	var status audit.RuntimeClosedLoopStatus
	if err := json.NewDecoder(statusRec.Body).Decode(&status); err != nil {
		t.Fatalf("decode closed-loop status: %v", err)
	}
	if status.Quarantined != 1 || status.CountsByQuarantine["vex"] != 1 {
		t.Fatalf("unexpected closed-loop status %#v", status)
	}

	quarantineReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/quarantine?tenant_id=acme", nil)
	quarantineRec := httptest.NewRecorder()
	handler.ServeHTTP(quarantineRec, quarantineReq)
	if quarantineRec.Code != http.StatusOK {
		t.Fatalf("expected 200 for quarantine list, got %d: %s", quarantineRec.Code, quarantineRec.Body.String())
	}
}

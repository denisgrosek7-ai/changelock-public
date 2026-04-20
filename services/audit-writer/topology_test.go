package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
)

func TestTopologyGraphBuildsDeclaredObservedAndEffectiveViews(t *testing.T) {
	handler := topologyTestHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/topology/graph?tenant_id=acme&environment=prod&limit=10", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected topology graph 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response topologyGraphResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode topology graph: %v", err)
	}
	if response.Summary.PublicEntryNodes == 0 || response.Summary.EffectiveEdges == 0 {
		t.Fatalf("expected topology graph summary with public entry and effective edges, got %#v", response.Summary)
	}
	if len(response.DeclaredGraph.Edges) == 0 || len(response.ObservedGraph.Edges) == 0 || len(response.EffectiveGraph.Nodes) == 0 {
		t.Fatalf("expected declared, observed, and effective graph data, got %#v", response)
	}
}

func TestIncidentBlastRadiusShowsCriticalReachAndContainmentOptions(t *testing.T) {
	handler := topologyTestHandler(t)
	incidentID := fetchIncidentForWorkload(t, handler, "edge-gateway")

	req := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"/blast-radius?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected incident blast radius 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response topologyBlastRadiusResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode incident blast radius: %v", err)
	}
	if response.SubjectType != "incident" || response.BlastRadiusScore == 0 {
		t.Fatalf("expected incident blast radius payload, got %#v", response)
	}
	if response.CriticalReachCount == 0 {
		t.Fatalf("expected critical downstream reach from topology graph, got %#v", response)
	}
	if len(response.ContainmentOptions) == 0 || response.PrimaryAffectedNode == nil {
		t.Fatalf("expected containment options and primary node, got %#v", response)
	}
}

func TestTopologyDeltaAndQuarantineSimulationAreExplainable(t *testing.T) {
	handler := topologyTestHandler(t)

	deltaReq := httptest.NewRequest(http.MethodGet, "/v1/topology/delta?tenant_id=acme&environment=prod&window=7d&limit=5", nil)
	deltaReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	deltaRec := httptest.NewRecorder()
	handler.ServeHTTP(deltaRec, deltaReq)
	if deltaRec.Code != http.StatusOK {
		t.Fatalf("expected topology delta 200, got %d: %s", deltaRec.Code, deltaRec.Body.String())
	}

	var delta topologyDeltaResponse
	if err := json.NewDecoder(deltaRec.Body).Decode(&delta); err != nil {
		t.Fatalf("decode topology delta: %v", err)
	}
	if len(delta.Items) == 0 {
		t.Fatalf("expected topology delta items, got %#v", delta)
	}
	if len(delta.Items[0].DriftSignals) == 0 {
		t.Fatalf("expected explainable drift signals, got %#v", delta.Items[0])
	}

	body := bytes.NewBufferString(`{"service":"edge-gateway"}`)
	simReq := httptest.NewRequest(http.MethodPost, "/v1/topology/quarantine-simulation?tenant_id=acme&environment=prod", body)
	simReq.Header.Set("Authorization", "Bearer operator-demo-token")
	simReq.Header.Set("Content-Type", "application/json")
	simRec := httptest.NewRecorder()
	handler.ServeHTTP(simRec, simReq)
	if simRec.Code != http.StatusOK {
		t.Fatalf("expected topology quarantine simulation 200, got %d: %s", simRec.Code, simRec.Body.String())
	}

	var simulation topologyQuarantineSimulationResponse
	if err := json.NewDecoder(simRec.Body).Decode(&simulation); err != nil {
		t.Fatalf("decode topology simulation: %v", err)
	}
	if !simulation.ApprovalRequired || simulation.Reduction == 0 || len(simulation.Options) == 0 {
		t.Fatalf("expected approval-based topology simulation with reduction, got %#v", simulation)
	}
}

func topologyTestHandler(t *testing.T) http.Handler {
	t.Helper()
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	store := audit.NewMemoryStore()
	now := time.Now().UTC()
	events := []audit.Event{
		{
			RequestID:      "topology-baseline-edge",
			Timestamp:      now.Add(-10 * 24 * time.Hour),
			Component:      "deploy-gate",
			EventType:      audit.EventTypeDeployGateDecision,
			Decision:       audit.DecisionAllow,
			TenantID:       "acme",
			Repo:           "acme/platform-edge",
			Environment:    "prod",
			Namespace:      "acme-prod",
			Workload:       "edge-gateway",
			ServiceAccount: "edge-sa",
			Digest:         "sha256:edge-v1",
			Reasons:        []string{"approved release"},
		},
		{
			RequestID:      "topology-current-edge",
			Timestamp:      now.Add(-2 * 24 * time.Hour),
			Component:      "deploy-gate",
			EventType:      audit.EventTypeDeployGateDecision,
			Decision:       audit.DecisionDeny,
			TenantID:       "acme",
			Repo:           "acme/platform-edge",
			Environment:    "prod",
			Namespace:      "acme-prod",
			Workload:       "edge-gateway",
			ServiceAccount: "edge-sa",
			Digest:         "sha256:edge-v2",
			Reasons:        []string{"ingress route widened", "workflow mismatch"},
		},
		{
			RequestID:            "topology-auth",
			Timestamp:            now.Add(-36 * time.Hour),
			Component:            "runtime-agent",
			EventType:            audit.EventTypeRuntimeActiveStateObserved,
			Decision:             audit.DecisionDeny,
			TenantID:             "acme",
			Repo:                 "acme/platform-auth",
			Environment:          "prod",
			Namespace:            "acme-prod",
			Workload:             "auth-api",
			ServiceAccount:       "edge-sa",
			Digest:               "sha256:auth-v2",
			DriftResult:          "service_account_drift",
			DriftClasses:         []string{"service_account_drift"},
			Reasons:              []string{"service account drift"},
			ReconciliationStatus: "drift_detected",
		},
		{
			RequestID:            "topology-db",
			Timestamp:            now.Add(-30 * time.Hour),
			Component:            "runtime-agent",
			EventType:            audit.EventTypeRuntimeActiveStateObserved,
			Decision:             audit.DecisionAllow,
			TenantID:             "acme",
			Repo:                 "acme/platform-billing",
			Environment:          "prod",
			Namespace:            "acme-prod",
			Workload:             "billing-db",
			ServiceAccount:       "db-sa",
			Digest:               "sha256:billing-v1",
			ReconciliationStatus: "in_sync",
			Reasons:              []string{"observed healthy"},
		},
		{
			RequestID:      "topology-stage-release",
			Timestamp:      now.Add(-20 * time.Hour),
			Component:      "deploy-gate",
			EventType:      audit.EventTypeDeployGateDecision,
			Decision:       audit.DecisionAllow,
			TenantID:       "acme",
			Repo:           "acme/platform-edge",
			Environment:    "stage",
			Namespace:      "acme-stage",
			Workload:       "edge-gateway",
			ServiceAccount: "edge-sa",
			Digest:         "sha256:edge-v2",
			Reasons:        []string{"stage rollout"},
		},
	}
	for _, event := range events {
		if _, err := store.Ingest(context.Background(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}
	return newHandlerWithAuth(store, "memory", authConfig)
}

func fetchIncidentForWorkload(t *testing.T, handler http.Handler, workload string) string {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected incidents 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var response incidentsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode incidents: %v", err)
	}
	for _, incident := range response.Incidents {
		for _, affected := range incident.AffectedWorkloads {
			if strings.TrimSpace(affected) == workload {
				return incident.ID
			}
		}
	}
	t.Fatalf("incident for workload %q not found in %#v", workload, response.Incidents)
	return ""
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
)

func TestCommandCenterNotificationsGroupWorkflowAndIntelligenceSignals(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	seedPhase5CommandCenterArtifacts(t, handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/command-center/notifications?tenant_id=acme&environment=prod&repo=github.com/acme/api&limit=10", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected notifications 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response commandCenterNotificationsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode notifications: %v", err)
	}
	if response.SchemaVersion != commandNotificationsSchemaVersion {
		t.Fatalf("expected schema-versioned notifications, got %#v", response)
	}
	if len(response.Items) == 0 {
		t.Fatalf("expected grouped notifications, got %#v", response)
	}
	if len(response.CountsByState) == 0 {
		t.Fatalf("expected notification state counts, got %#v", response)
	}

	var sawWorkflow, sawIntelligence bool
	for _, item := range response.Items {
		if item.SchemaVersion != commandNotificationSchemaVersion {
			t.Fatalf("expected schema-versioned notification item, got %#v", item)
		}
		switch item.LifecyclePhase {
		case "workflow":
			sawWorkflow = true
			if item.OwnerHint != "team-api" {
				t.Fatalf("expected workflow notification owner hint, got %#v", item)
			}
		case "intelligence":
			sawIntelligence = true
		}
	}
	if !sawWorkflow || !sawIntelligence {
		t.Fatalf("expected notifications to include intelligence and workflow buckets, got %#v", response.Items)
	}

	filteredReq := httptest.NewRequest(http.MethodGet, "/v1/command-center/notifications?tenant_id=acme&environment=prod&repo=github.com/acme/api&limit=10&lifecycle_phase=workflow", nil)
	filteredReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	filteredRec := httptest.NewRecorder()
	handler.ServeHTTP(filteredRec, filteredReq)
	if filteredRec.Code != http.StatusOK {
		t.Fatalf("expected filtered notifications 200, got %d: %s", filteredRec.Code, filteredRec.Body.String())
	}
	var filtered commandCenterNotificationsResponse
	if err := json.NewDecoder(filteredRec.Body).Decode(&filtered); err != nil {
		t.Fatalf("decode filtered notifications: %v", err)
	}
	for _, item := range filtered.Items {
		if item.LifecyclePhase != "workflow" {
			t.Fatalf("expected only workflow notifications after lifecycle filter, got %#v", filtered.Items)
		}
	}
}

func seedPhase5CommandCenterArtifacts(t *testing.T, handler http.Handler) {
	t.Helper()

	vulnerabilityReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/vulnerability-relevance?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "vulnerability_id":"CVE-2026-2000",
	    "image_digest":"sha256:api",
	    "package_name":"openssl",
	    "severity":"critical",
	    "reachability":{"current_state":"observed_reachable","confidence_score":92,"evidence_refs":["runtime:callgraph"]},
	    "exploitability":{"epss":0.92,"external_exposure":true,"local_confidence":88}
	  }
	}`))
	vulnerabilityReq.Header.Set("Authorization", "Bearer operator-demo-token")
	vulnerabilityReq.Header.Set("Content-Type", "application/json")
	vulnerabilityRec := httptest.NewRecorder()
	handler.ServeHTTP(vulnerabilityRec, vulnerabilityReq)
	if vulnerabilityRec.Code != http.StatusCreated {
		t.Fatalf("expected vulnerability relevance 201, got %d: %s", vulnerabilityRec.Code, vulnerabilityRec.Body.String())
	}

	workflowReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/workflow/lifecycle?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "workflow_id":"wf-phase5",
	    "artifact_type":"finding",
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "severity":"critical",
	    "requested_state":"resolved",
	    "validation_required":true,
	    "validation_state":"pending",
	    "owners":{"finding_owner":"team-api","remediation_owner":"team-api","approver":"security-admin"},
	    "evidence_refs":["event://deploy-gate/1"]
	  }
	}`))
	workflowReq.Header.Set("Authorization", "Bearer operator-demo-token")
	workflowReq.Header.Set("Content-Type", "application/json")
	workflowRec := httptest.NewRecorder()
	handler.ServeHTTP(workflowRec, workflowReq)
	if workflowRec.Code != http.StatusCreated {
		t.Fatalf("expected workflow 201, got %d: %s", workflowRec.Code, workflowRec.Body.String())
	}
}

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
)

func TestIdentityFabricHandler(t *testing.T) {
	cfg, signer := newOIDCHandlerConfig(t, true, true)
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", cfg)

	req := httptest.NewRequest(http.MethodGet, "/v1/integrations/identity-fabric", nil)
	req.Header.Set("Authorization", "Bearer "+signer.token(t, map[string]any{
		"sub":       "viewer@example.com",
		"email":     "viewer@example.com",
		"groups":    []string{"changelock-viewers"},
		"tenant_id": "acme",
	}))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected identity fabric 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response identityFabricResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode identity fabric response: %v", err)
	}
	if response.SchemaVersion != identityFabricSchemaVersion {
		t.Fatalf("expected identity fabric schema version, got %#v", response)
	}
	if response.CurrentActor.Role != auth.RoleViewer || response.AuthModel.Mode != auth.ModeOIDCJWT {
		t.Fatalf("expected OIDC-backed identity response, got %#v", response)
	}
	if len(response.TenantToBusinessRoleMapping) == 0 || len(response.ApproverClasses) == 0 {
		t.Fatalf("expected business-role mapping and approver classes, got %#v", response)
	}
}

func TestITSMLifecycleHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/integrations/itsm-lifecycle?tenant_id=acme&environment=prod&limit=10", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected itsm lifecycle 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response itsmLifecycleResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode itsm lifecycle response: %v", err)
	}
	if response.SchemaVersion != itsmLifecycleSchemaVersion {
		t.Fatalf("expected itsm lifecycle schema version, got %#v", response)
	}
	if len(response.Systems) == 0 || response.Systems[0].WriteMode != "draft_before_write_only" {
		t.Fatalf("expected bounded ITSM draft contract, got %#v", response.Systems)
	}
	if response.ScopeSummary.IncidentCount == 0 || response.ScopeSummary.RecommendationCount == 0 {
		t.Fatalf("expected incidents and recommendations in scope summary, got %#v", response.ScopeSummary)
	}
	if len(response.TicketClasses) < 3 || len(response.StateSyncRules) == 0 || len(response.OperatorOverrides) == 0 {
		t.Fatalf("expected richer ITSM lifecycle semantics, got %#v", response)
	}

	incidentID := fetchFirstIncidentID(t, fixture.handler)
	flowsReq := httptest.NewRequest(http.MethodGet, "/v1/integrations/itsm-lifecycle/flows?tenant_id=acme&environment=prod&incident_id="+incidentID, nil)
	flowsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	flowsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(flowsRec, flowsReq)
	if flowsRec.Code != http.StatusOK {
		t.Fatalf("expected itsm lifecycle flows 200, got %d: %s", flowsRec.Code, flowsRec.Body.String())
	}

	var flows itsmLifecycleFlowsResponse
	if err := json.NewDecoder(flowsRec.Body).Decode(&flows); err != nil {
		t.Fatalf("decode itsm lifecycle flows: %v", err)
	}
	if flows.SchemaVersion != itsmLifecycleFlowsSchemaVersion || len(flows.Items) < 2 {
		t.Fatalf("expected evidence-backed ITSM flows, got %#v", flows)
	}
	if !hasITSMTicketClass(flows.Items, "incident") || !hasITSMTicketClass(flows.Items, "remediation") {
		t.Fatalf("expected incident and remediation flow items, got %#v", flows.Items)
	}
}

func TestSIEMSyncHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	getReq := httptest.NewRequest(http.MethodGet, "/v1/integrations/siem-sync?tenant_id=acme", nil)
	getReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	getRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(getRec, getReq)
	if getRec.Code != http.StatusOK {
		t.Fatalf("expected siem sync GET 200, got %d: %s", getRec.Code, getRec.Body.String())
	}

	var summary siemSyncResponse
	if err := json.NewDecoder(getRec.Body).Decode(&summary); err != nil {
		t.Fatalf("decode siem sync response: %v", err)
	}
	if summary.SchemaVersion != siemSyncSchemaVersion || summary.InboundEvaluateEndpoint == "" {
		t.Fatalf("expected bounded siem sync contract, got %#v", summary)
	}

	postReq := httptest.NewRequest(http.MethodPost, "/v1/integrations/siem-sync/evaluate?tenant_id=acme", strings.NewReader(`{
	  "source_system":"splunk",
	  "source_trust":"trusted",
	  "signal_type":"response_hint",
	  "severity":"critical",
	  "correlation_id":"INC-4A-1",
	  "hinted_action":"security review"
	}`))
	postReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	postReq.Header.Set("Content-Type", "application/json")
	postRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(postRec, postReq)
	if postRec.Code != http.StatusOK {
		t.Fatalf("expected siem evaluation 200, got %d: %s", postRec.Code, postRec.Body.String())
	}

	var evaluation siemSignalEvaluationResponse
	if err := json.NewDecoder(postRec.Body).Decode(&evaluation); err != nil {
		t.Fatalf("decode siem evaluation response: %v", err)
	}
	if evaluation.SchemaVersion != siemSyncEvaluationSchemaVersion || evaluation.CorrelationID != "INC-4A-1" {
		t.Fatalf("expected correlation-preserving siem evaluation, got %#v", evaluation)
	}
	if evaluation.ActionabilityState != "review_required" || evaluation.MappedRecommendation != "create_security_review" {
		t.Fatalf("expected approval-gated security review mapping, got %#v", evaluation)
	}
	if evaluation.SourceTrustLabel != "TRUSTED" || evaluation.SafetyLimitRef == "" || evaluation.MappedWorkflowState == "" {
		t.Fatalf("expected explicit trust labeling and safety metadata, got %#v", evaluation)
	}
}

func TestIncidentCollaborationHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)
	incidentID := fetchFirstIncidentID(t, fixture.handler)

	req := httptest.NewRequest(http.MethodGet, "/v1/incidents/collaboration?tenant_id=acme&environment=prod&incident_id="+incidentID, nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected incident collaboration 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response incidentCollaborationResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode incident collaboration response: %v", err)
	}
	if response.SchemaVersion != incidentCollaborationSchemaVersion || response.IncidentRef == "" {
		t.Fatalf("expected incident collaboration schema and incident ref, got %#v", response)
	}
	if len(response.ExportVariants) != 3 || len(response.LinkedEvidenceRefs) == 0 {
		t.Fatalf("expected export variants and evidence linkage, got %#v", response)
	}
	if len(response.SharedContextModel) == 0 || response.VerificationAfterRemediation.CurrentState == "" || len(response.AudienceExportDiscipline) == 0 {
		t.Fatalf("expected collaboration lifecycle and audience discipline, got %#v", response)
	}
}

func TestIntegrationSafetyHandler(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	syncRuntime := newSyncRuntime(syncConfig{
		Mode:         audit.SyncModeSpoke,
		ClusterID:    "local",
		HubURL:       "https://hub.example.com",
		Token:        "sync-token",
		PollInterval: time.Minute,
		FailMode:     audit.SyncFailModeLastKnownGood,
	})
	syncRuntime.markFailure(errors.New("hub unavailable"), true)
	handler := newHandlerWithRuntimes(audit.NewMemoryStore(), "memory", authConfig, nil, syncRuntime)

	req := httptest.NewRequest(http.MethodGet, "/v1/integrations/safety?tenant_id=acme", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected integration safety 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response integrationSafetyResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode integration safety response: %v", err)
	}
	if response.SchemaVersion != integrationSafetySchemaVersion || !response.NoNewTruthLayer {
		t.Fatalf("expected integration safety schema and bounded truth posture, got %#v", response)
	}
	clusterSync := findIntegrationConnector(t, response.Connectors, "cluster_sync")
	if clusterSync.CurrentState != audit.SyncHealthStale {
		t.Fatalf("expected stale cluster sync connector state, got %#v", clusterSync)
	}

	healthReq := httptest.NewRequest(http.MethodGet, "/v1/integrations/safety/health?tenant_id=acme", nil)
	healthReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	healthRec := httptest.NewRecorder()
	handler.ServeHTTP(healthRec, healthReq)
	if healthRec.Code != http.StatusOK {
		t.Fatalf("expected integration safety health 200, got %d: %s", healthRec.Code, healthRec.Body.String())
	}

	var health integrationSafetyHealthResponse
	if err := json.NewDecoder(healthRec.Body).Decode(&health); err != nil {
		t.Fatalf("decode integration safety health response: %v", err)
	}
	if health.SchemaVersion != integrationSafetyHealthSchema {
		t.Fatalf("expected safety health schema, got %#v", health)
	}
	clusterHealth := findIntegrationHealthConnector(t, health.Connectors, "cluster_sync")
	if clusterHealth.HealthState != audit.SyncHealthStale || !clusterHealth.ReplaySafe {
		t.Fatalf("expected explicit stale connector health contract, got %#v", clusterHealth)
	}
}

func findIntegrationConnector(t *testing.T, connectors []integrationConnectorSafety, connectorID string) integrationConnectorSafety {
	t.Helper()
	for _, connector := range connectors {
		if connector.ConnectorID == connectorID {
			return connector
		}
	}
	t.Fatalf("expected integration connector %q, got %#v", connectorID, connectors)
	return integrationConnectorSafety{}
}

func findIntegrationHealthConnector(t *testing.T, connectors []integrationConnectorHealth, connectorID string) integrationConnectorHealth {
	t.Helper()
	for _, connector := range connectors {
		if connector.ConnectorID == connectorID {
			return connector
		}
	}
	t.Fatalf("expected integration health connector %q, got %#v", connectorID, connectors)
	return integrationConnectorHealth{}
}

func fetchFirstIncidentID(t *testing.T, handler http.Handler) string {
	t.Helper()
	incidentReq := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme&environment=prod&limit=5", nil)
	incidentReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	incidentRec := httptest.NewRecorder()
	handler.ServeHTTP(incidentRec, incidentReq)
	if incidentRec.Code != http.StatusOK {
		t.Fatalf("expected incidents 200, got %d: %s", incidentRec.Code, incidentRec.Body.String())
	}

	var incidents incidentsResponse
	if err := json.NewDecoder(incidentRec.Body).Decode(&incidents); err != nil {
		t.Fatalf("decode incidents response: %v", err)
	}
	if len(incidents.Incidents) == 0 {
		t.Fatal("expected at least one incident in fixture scope")
	}
	return incidents.Incidents[0].ID
}

func hasITSMTicketClass(items []itsmLifecycleFlowItem, ticketClass string) bool {
	for _, item := range items {
		if item.TicketClass == ticketClass {
			return true
		}
	}
	return false
}

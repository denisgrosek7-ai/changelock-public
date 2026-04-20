package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntegrationIdentityAndTicketCatalog(t *testing.T) {
	fixture := forensicsTestFixture(t)

	identityReq := httptest.NewRequest(http.MethodGet, "/v1/integrations/identity?tenant_id=acme&environment=prod", nil)
	identityReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	identityRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(identityRec, identityReq)
	if identityRec.Code != http.StatusOK {
		t.Fatalf("expected identity 200, got %d: %s", identityRec.Code, identityRec.Body.String())
	}
	var identity integrationIdentityResponse
	if err := json.NewDecoder(identityRec.Body).Decode(&identity); err != nil {
		t.Fatalf("decode identity response: %v", err)
	}
	if identity.SchemaVersion != integrationIdentitySchemaVersion {
		t.Fatalf("expected schema-versioned identity response, got %#v", identity)
	}
	if !identity.CurrentActor.Authenticated || identity.CurrentActor.Role != "viewer" {
		t.Fatalf("expected current actor attribution, got %#v", identity.CurrentActor)
	}
	if identity.AuthModel.IdentityProvider != "static_token" || !identity.AuthModel.AuditActorAttribution || !identity.AuthModel.ApprovalActorAttribution {
		t.Fatalf("expected auth model description, got %#v", identity.AuthModel)
	}

	catalogReq := httptest.NewRequest(http.MethodGet, "/v1/integrations/tickets/catalog?tenant_id=acme&environment=prod", nil)
	catalogReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	catalogRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(catalogRec, catalogReq)
	if catalogRec.Code != http.StatusOK {
		t.Fatalf("expected ticket catalog 200, got %d: %s", catalogRec.Code, catalogRec.Body.String())
	}
	var catalog integrationTicketCatalogResponse
	if err := json.NewDecoder(catalogRec.Body).Decode(&catalog); err != nil {
		t.Fatalf("decode ticket catalog: %v", err)
	}
	if catalog.SchemaVersion != integrationTicketCatalogSchemaVersion || len(catalog.Systems) != 2 {
		t.Fatalf("expected schema-versioned ticket catalog, got %#v", catalog)
	}
}

func TestIntegrationTicketPrepareProducesEvidenceBackedDraft(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)

	incidentsReq := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme&environment=prod&limit=20", nil)
	incidentsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	incidentsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(incidentsRec, incidentsReq)
	if incidentsRec.Code != http.StatusOK {
		t.Fatalf("expected incidents 200, got %d: %s", incidentsRec.Code, incidentsRec.Body.String())
	}
	var incidents incidentsResponse
	if err := json.NewDecoder(incidentsRec.Body).Decode(&incidents); err != nil {
		t.Fatalf("decode incidents: %v", err)
	}
	if len(incidents.Incidents) == 0 {
		t.Fatal("expected fixture incidents")
	}

	recommendationsReq := httptest.NewRequest(http.MethodGet, "/v1/recommendations?tenant_id=acme&environment=prod&limit=50", nil)
	recommendationsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	recommendationsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(recommendationsRec, recommendationsReq)
	if recommendationsRec.Code != http.StatusOK {
		t.Fatalf("expected recommendations 200, got %d: %s", recommendationsRec.Code, recommendationsRec.Body.String())
	}
	var recommendationList recommendationListResponse
	if err := json.NewDecoder(recommendationsRec.Body).Decode(&recommendationList); err != nil {
		t.Fatalf("decode recommendations: %v", err)
	}
	if len(recommendationList.Recommendations) == 0 {
		t.Fatal("expected recommendation items")
	}

	validationReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/validation/execute?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"scenario_ids":["safe_release_positive","unsigned_image_block"]}`),
	)
	validationReq.Header.Set("Authorization", "Bearer operator-demo-token")
	validationReq.Header.Set("Content-Type", "application/json")
	validationRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(validationRec, validationReq)
	if validationRec.Code != http.StatusOK {
		t.Fatalf("expected validation execute 200, got %d: %s", validationRec.Code, validationRec.Body.String())
	}

	sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)

	requestBody := bytes.NewBufferString(`{"system":"jira","incident_id":"` + incidents.Incidents[0].ID + `","recommendation_id":"` + recommendationList.Recommendations[0].RecommendationID + `"}`)
	req := httptest.NewRequest(http.MethodPost, "/v1/integrations/tickets/prepare?tenant_id=acme&environment=prod", requestBody)
	req.Header.Set("Authorization", "Bearer operator-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected ticket prepare 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var draft integrationTicketDraftResponse
	if err := json.NewDecoder(rec.Body).Decode(&draft); err != nil {
		t.Fatalf("decode ticket draft: %v", err)
	}
	if draft.SchemaVersion != integrationTicketDraftSchemaVersion {
		t.Fatalf("expected schema-versioned ticket draft, got %#v", draft)
	}
	if draft.System != integrationSystemJira || draft.IncidentRef == "" || draft.RecommendationRef == "" {
		t.Fatalf("expected linked incident and recommendation refs, got %#v", draft)
	}
	if len(draft.EvidenceRefs) == 0 || draft.DeepLinks["incident"] == "" || draft.DeepLinks["recommendation"] == "" {
		t.Fatalf("expected evidence-backed ticket draft with deeplinks, got %#v", draft)
	}
	if draft.DeepLinks["incident_export"] == "" {
		t.Fatalf("expected incident export deeplink on ticket draft, got %#v", draft)
	}
	if len(draft.ValidationRefs) == 0 {
		t.Fatalf("expected ticket draft to carry validation linkage, got %#v", draft)
	}
	if len(draft.HandoffRefs) == 0 {
		t.Fatalf("expected ticket draft to carry handoff linkage, got %#v", draft)
	}
}

func TestIntegrationSIEMAndEvidenceExport(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	sealed := sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)

	validationReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/validation/execute?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"scenario_ids":["safe_release_positive","unsigned_image_block"]}`),
	)
	validationReq.Header.Set("Authorization", "Bearer operator-demo-token")
	validationReq.Header.Set("Content-Type", "application/json")
	validationRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(validationRec, validationReq)
	if validationRec.Code != http.StatusOK {
		t.Fatalf("expected validation execute 200, got %d: %s", validationRec.Code, validationRec.Body.String())
	}
	var validationRun validationExecutionRun
	if err := json.NewDecoder(validationRec.Body).Decode(&validationRun); err != nil {
		t.Fatalf("decode validation run: %v", err)
	}

	incidentsReq := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme&environment=prod&limit=20", nil)
	incidentsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	incidentsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(incidentsRec, incidentsReq)
	if incidentsRec.Code != http.StatusOK {
		t.Fatalf("expected incidents 200, got %d: %s", incidentsRec.Code, incidentsRec.Body.String())
	}
	var incidents incidentsResponse
	if err := json.NewDecoder(incidentsRec.Body).Decode(&incidents); err != nil {
		t.Fatalf("decode incidents: %v", err)
	}
	if len(incidents.Incidents) == 0 {
		t.Fatal("expected incident fixture data")
	}

	siemReq := httptest.NewRequest(http.MethodGet, "/v1/integrations/siem/export?tenant_id=acme&environment=prod&limit=20", nil)
	siemReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	siemRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(siemRec, siemReq)
	if siemRec.Code != http.StatusOK {
		t.Fatalf("expected SIEM export 200, got %d: %s", siemRec.Code, siemRec.Body.String())
	}
	var siem integrationSIEMExportResponse
	if err := json.NewDecoder(siemRec.Body).Decode(&siem); err != nil {
		t.Fatalf("decode SIEM export: %v", err)
	}
	if siem.SchemaVersion != integrationSIEMExportSchemaVersion || len(siem.Items) == 0 {
		t.Fatalf("expected schema-versioned SIEM export items, got %#v", siem)
	}
	if siem.Items[0].EventType == "" || siem.Items[0].SourceComponent == "" || siem.Items[0].Severity == "" {
		t.Fatalf("expected normalized SIEM item, got %#v", siem.Items[0])
	}

	incidentExportReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/integrations/evidence/export?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"incident_id":"`+incidents.Incidents[0].ID+`"}`),
	)
	incidentExportReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	incidentExportReq.Header.Set("Content-Type", "application/json")
	incidentExportRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(incidentExportRec, incidentExportReq)
	if incidentExportRec.Code != http.StatusOK {
		t.Fatalf("expected incident evidence export 200, got %d: %s", incidentExportRec.Code, incidentExportRec.Body.String())
	}
	var incidentExport integrationEvidenceExportResponse
	if err := json.NewDecoder(incidentExportRec.Body).Decode(&incidentExport); err != nil {
		t.Fatalf("decode incident evidence export: %v", err)
	}
	if incidentExport.SchemaVersion != integrationEvidenceSchemaVersion || len(incidentExport.Items) == 0 {
		t.Fatalf("expected schema-versioned incident evidence export, got %#v", incidentExport)
	}
	if incidentExport.Items[0].URI == "" {
		t.Fatalf("expected routable evidence export entries, got %#v", incidentExport.Items[0])
	}

	packageExportReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/integrations/evidence/export?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"package_id":"`+sealed.PackageID+`"}`),
	)
	packageExportReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	packageExportReq.Header.Set("Content-Type", "application/json")
	packageExportRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(packageExportRec, packageExportReq)
	if packageExportRec.Code != http.StatusOK {
		t.Fatalf("expected package evidence export 200, got %d: %s", packageExportRec.Code, packageExportRec.Body.String())
	}
	var packageExport integrationEvidenceExportResponse
	if err := json.NewDecoder(packageExportRec.Body).Decode(&packageExport); err != nil {
		t.Fatalf("decode package evidence export: %v", err)
	}
	if len(packageExport.Items) == 0 || !packageExport.Items[0].Sealed {
		t.Fatalf("expected sealed handoff evidence export entries, got %#v", packageExport)
	}

	validationExportReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/integrations/evidence/export?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"validation_run_id":"`+validationRun.RunID+`"}`),
	)
	validationExportReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	validationExportReq.Header.Set("Content-Type", "application/json")
	validationExportRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(validationExportRec, validationExportReq)
	if validationExportRec.Code != http.StatusOK {
		t.Fatalf("expected validation evidence export 200, got %d: %s", validationExportRec.Code, validationExportRec.Body.String())
	}
	var validationExport integrationEvidenceExportResponse
	if err := json.NewDecoder(validationExportRec.Body).Decode(&validationExport); err != nil {
		t.Fatalf("decode validation evidence export: %v", err)
	}
	if len(validationExport.Items) == 0 || validationExport.Items[0].ItemType != "validation_certificate" {
		t.Fatalf("expected validation certificate export entry, got %#v", validationExport)
	}
}

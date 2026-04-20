package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
)

func TestRecommendationsListSupportsIncidentAndPackageScopes(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	handler, incidentIDs := recommendationTestHandler(t)

	incidentReq := httptest.NewRequest(http.MethodGet, "/v1/recommendations?tenant_id=acme&incident_id="+incidentIDs[0]+"&limit=3", nil)
	incidentReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	incidentRec := httptest.NewRecorder()
	handler.ServeHTTP(incidentRec, incidentReq)
	if incidentRec.Code != http.StatusOK {
		t.Fatalf("expected incident recommendations 200, got %d: %s", incidentRec.Code, incidentRec.Body.String())
	}

	var incidentResponse recommendationListResponse
	if err := json.NewDecoder(incidentRec.Body).Decode(&incidentResponse); err != nil {
		t.Fatalf("decode incident recommendations: %v", err)
	}
	if len(incidentResponse.Recommendations) == 0 {
		t.Fatal("expected at least one incident recommendation")
	}
	if len(incidentResponse.Recommendations) > 3 {
		t.Fatalf("expected limit=3 to be respected, got %d items", len(incidentResponse.Recommendations))
	}
	item := incidentResponse.Recommendations[0]
	if item.SourceType != "incident" || item.SubjectType != "incident" {
		t.Fatalf("expected incident recommendation, got %#v", item)
	}
	if item.ActionTemplate.TemplateID == "" || item.RecommendedAction == "" || len(item.EvidenceRefs) == 0 || len(item.ReadbackRefs) == 0 {
		t.Fatalf("expected evidence-backed incident recommendation, got %#v", item)
	}
	if !item.AdvisoryOnly {
		t.Fatalf("expected advisory-only recommendation, got %#v", item)
	}

	packageReq := httptest.NewRequest(http.MethodGet, "/v1/recommendations?tenant_id=acme&source_type=package&package_incident_id="+incidentIDs[0]+"&package_incident_id="+incidentIDs[1], nil)
	packageReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	packageRec := httptest.NewRecorder()
	handler.ServeHTTP(packageRec, packageReq)
	if packageRec.Code != http.StatusOK {
		t.Fatalf("expected package recommendations 200, got %d: %s", packageRec.Code, packageRec.Body.String())
	}

	var packageResponse recommendationListResponse
	if err := json.NewDecoder(packageRec.Body).Decode(&packageResponse); err != nil {
		t.Fatalf("decode package recommendations: %v", err)
	}
	if len(packageResponse.Recommendations) != 1 {
		t.Fatalf("expected exactly one package recommendation, got %#v", packageResponse)
	}
	pkg := packageResponse.Recommendations[0]
	if pkg.SourceType != "package" || len(pkg.RelatedIncidentRefs) != 2 {
		t.Fatalf("expected package recommendation with both incidents linked, got %#v", pkg)
	}
	if pkg.ActionTemplate.TemplateID != "create_ticket" {
		t.Fatalf("expected package workflow template, got %#v", pkg.ActionTemplate)
	}
}

func TestRecommendationOverlaySupportsAssignmentAndComment(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	handler, incidentIDs := recommendationTestHandler(t)
	recommendationItem := fetchRecommendationBySource(t, handler, "/v1/recommendations?tenant_id=acme&incident_id="+incidentIDs[0], "incident")

	assignReq := httptest.NewRequest(http.MethodPost, "/v1/recommendations/"+recommendationItem.RecommendationID+"/assign?tenant_id=acme", bytes.NewBufferString(`{"owner":"secops","reason":"route remediation owner"}`))
	assignReq.Header.Set("Authorization", "Bearer operator-demo-token")
	assignReq.Header.Set("Content-Type", "application/json")
	assignRec := httptest.NewRecorder()
	handler.ServeHTTP(assignRec, assignReq)
	if assignRec.Code != http.StatusOK {
		t.Fatalf("expected assign 200, got %d: %s", assignRec.Code, assignRec.Body.String())
	}

	commentReq := httptest.NewRequest(http.MethodPost, "/v1/recommendations/"+recommendationItem.RecommendationID+"/comment?tenant_id=acme", bytes.NewBufferString(`{"comment":"validate fix path before widening any exception scope"}`))
	commentReq.Header.Set("Authorization", "Bearer operator-demo-token")
	commentReq.Header.Set("Content-Type", "application/json")
	commentRec := httptest.NewRecorder()
	handler.ServeHTTP(commentRec, commentReq)
	if commentRec.Code != http.StatusOK {
		t.Fatalf("expected comment 200, got %d: %s", commentRec.Code, commentRec.Body.String())
	}

	getReq := httptest.NewRequest(http.MethodGet, "/v1/recommendations/"+recommendationItem.RecommendationID+"?tenant_id=acme", nil)
	getReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	getRec := httptest.NewRecorder()
	handler.ServeHTTP(getRec, getReq)
	if getRec.Code != http.StatusOK {
		t.Fatalf("expected get recommendation 200, got %d: %s", getRec.Code, getRec.Body.String())
	}

	var updated recommendation
	if err := json.NewDecoder(getRec.Body).Decode(&updated); err != nil {
		t.Fatalf("decode updated recommendation: %v", err)
	}
	if updated.Owner != "secops" {
		t.Fatalf("expected recommendation owner to persist, got %#v", updated)
	}
	if len(updated.Comments) == 0 || updated.Comments[0].Comment == "" {
		t.Fatalf("expected recommendation comment to persist, got %#v", updated.Comments)
	}
	if len(updated.History) < 2 {
		t.Fatalf("expected overlay history entries for assign/comment, got %#v", updated.History)
	}
}

func TestRecommendationApprovalGuardAndVerification(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	handler, _ := recommendationTestHandler(t)
	recommendationItem := fetchRecommendationBySource(t, handler, "/v1/recommendations?tenant_id=acme&source_type=systemic_weakness", "systemic_weakness")
	if recommendationItem.ApprovalMode != recommendationApprovalHumanReview {
		t.Fatalf("expected approval-required systemic recommendation, got %#v", recommendationItem)
	}

	executeReq := httptest.NewRequest(http.MethodPost, "/v1/recommendations/"+recommendationItem.RecommendationID+"/execute?tenant_id=acme", bytes.NewBufferString(`{"template_id":"`+recommendationItem.ActionTemplate.TemplateID+`"}`))
	executeReq.Header.Set("Authorization", "Bearer operator-demo-token")
	executeReq.Header.Set("Content-Type", "application/json")
	executeRec := httptest.NewRecorder()
	handler.ServeHTTP(executeRec, executeReq)
	if executeRec.Code != http.StatusConflict {
		t.Fatalf("expected approval-required execute to return 409, got %d: %s", executeRec.Code, executeRec.Body.String())
	}

	acceptReq := httptest.NewRequest(http.MethodPost, "/v1/recommendations/"+recommendationItem.RecommendationID+"/accept?tenant_id=acme", nil)
	acceptReq.Header.Set("Authorization", "Bearer operator-demo-token")
	acceptRec := httptest.NewRecorder()
	handler.ServeHTTP(acceptRec, acceptReq)
	if acceptRec.Code != http.StatusOK {
		t.Fatalf("expected accept 200, got %d: %s", acceptRec.Code, acceptRec.Body.String())
	}

	approvalReq := httptest.NewRequest(http.MethodPost, "/v1/recommendations/"+recommendationItem.RecommendationID+"/approval-request?tenant_id=acme", bytes.NewBufferString(`{"summary":"security review requested"}`))
	approvalReq.Header.Set("Authorization", "Bearer operator-demo-token")
	approvalReq.Header.Set("Content-Type", "application/json")
	approvalRec := httptest.NewRecorder()
	handler.ServeHTTP(approvalRec, approvalReq)
	if approvalRec.Code != http.StatusOK {
		t.Fatalf("expected approval request 200, got %d: %s", approvalRec.Code, approvalRec.Body.String())
	}

	executeReq = httptest.NewRequest(http.MethodPost, "/v1/recommendations/"+recommendationItem.RecommendationID+"/execute?tenant_id=acme", bytes.NewBufferString(`{"template_id":"`+recommendationItem.ActionTemplate.TemplateID+`"}`))
	executeReq.Header.Set("Authorization", "Bearer operator-demo-token")
	executeReq.Header.Set("Content-Type", "application/json")
	executeRec = httptest.NewRecorder()
	handler.ServeHTTP(executeRec, executeReq)
	if executeRec.Code != http.StatusOK {
		t.Fatalf("expected execute 200 after acceptance, got %d: %s", executeRec.Code, executeRec.Body.String())
	}

	verifyReq := httptest.NewRequest(http.MethodPost, "/v1/recommendations/"+recommendationItem.RecommendationID+"/verify?tenant_id=acme", nil)
	verifyReq.Header.Set("Authorization", "Bearer operator-demo-token")
	verifyRec := httptest.NewRecorder()
	handler.ServeHTTP(verifyRec, verifyReq)
	if verifyRec.Code != http.StatusOK {
		t.Fatalf("expected verify 200, got %d: %s", verifyRec.Code, verifyRec.Body.String())
	}

	var verified recommendation
	if err := json.NewDecoder(verifyRec.Body).Decode(&verified); err != nil {
		t.Fatalf("decode verified recommendation: %v", err)
	}
	if verified.Status != recommendationStatusExecutedNoEffect || verified.Outcome.Status != recommendationStatusExecutedNoEffect {
		t.Fatalf("expected systemic recommendation verification to remain advisory no-effect in current scope, got %#v", verified.Outcome)
	}
	if len(verified.History) < 3 {
		t.Fatalf("expected workflow history after accept/approval/execute/verify, got %#v", verified.History)
	}
}

func recommendationTestHandler(t *testing.T) (http.Handler, []string) {
	t.Helper()

	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}

	store := audit.NewMemoryStore()
	mustIngest := func(event audit.Event) {
		t.Helper()
		if _, err := store.Ingest(context.Background(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}

	mustIngest(audit.Event{
		RequestID:      "req-recommendation-1",
		Component:      "deploy-gate",
		EventType:      audit.EventTypeDeployGateDecision,
		Decision:       audit.DecisionDeny,
		TenantID:       "acme",
		Repo:           "repo-package-a",
		Environment:    "prod",
		Workload:       "api",
		Digest:         "sha256:recommendation-a",
		Reasons:        []string{"workflow mismatch", "signature verification failed"},
		PolicyBundleID: "bundle-recommendation-a",
	})
	mustIngest(audit.Event{
		RequestID:   "req-recommendation-2",
		Component:   "runtime-agent",
		EventType:   audit.EventTypeRuntimeDriftResult,
		Decision:    audit.DecisionDeny,
		TenantID:    "acme",
		Repo:        "repo-package-b",
		Environment: "prod",
		Workload:    "worker",
		Image:       "ghcr.io/acme/worker@sha256:recommendation-b",
		Digest:      "sha256:recommendation-b",
		DriftResult: "image_drift",
		Reasons:     []string{"image drift"},
	})

	handler := newHandlerWithAuth(store, "memory", authConfig)

	listReq := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme", nil)
	listReq.Header.Set("Authorization", "Bearer viewer-demo-token")
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

	return handler, []string{list.Incidents[0].ID, list.Incidents[1].ID}
}

func fetchRecommendationBySource(t *testing.T, handler http.Handler, path string, sourceType string) recommendation {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, path, nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected recommendations 200 for %s, got %d: %s", path, rec.Code, rec.Body.String())
	}

	var response recommendationListResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode recommendations for %s: %v", path, err)
	}
	for _, item := range response.Recommendations {
		if item.SourceType == sourceType {
			return item
		}
	}
	t.Fatalf("expected recommendation with source_type %q, got %#v", sourceType, response.Recommendations)
	return recommendation{}
}

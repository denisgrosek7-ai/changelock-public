package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
)

type defenseGapReadbackPayload struct {
	SchemaVersion      string                   `json:"schema_version"`
	ProjectionAudience string                   `json:"projection_audience"`
	EvidenceEnvelope   decisionEvidenceEnvelope `json:"evidence_envelope"`
	Payload            defenseGapAssessment     `json:"payload"`
	TopologyContext    *readbackTopologyContext `json:"topology_context"`
}

type policyReplayReadbackPayload struct {
	SchemaVersion      string                   `json:"schema_version"`
	ProjectionAudience string                   `json:"projection_audience"`
	EvidenceEnvelope   decisionEvidenceEnvelope `json:"evidence_envelope"`
	Payload            policyReplayAssessment   `json:"payload"`
	TopologyContext    *readbackTopologyContext `json:"topology_context"`
}

type systemicWeaknessReadbackPayload struct {
	SchemaVersion      string                   `json:"schema_version"`
	ProjectionAudience string                   `json:"projection_audience"`
	EvidenceEnvelope   decisionEvidenceEnvelope `json:"evidence_envelope"`
	Payload            systemicWeakness         `json:"payload"`
	TopologyContext    *readbackTopologyContext `json:"topology_context"`
}

func TestReadbackEndpointsExposeStableEvidenceEnvelope(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeDisabled)
	t.Setenv("CHANGELOCK_READBACK_GRANT_SECRET", "readback-secret")

	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")

	mustIngest := func(event audit.Event) {
		t.Helper()
		if _, err := store.Ingest(t.Context(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}

	mustIngest(audit.Event{
		RequestID:      "req-readback-deploy",
		Component:      "deploy-gate",
		EventType:      audit.EventTypeDeployGateDecision,
		Decision:       audit.DecisionDeny,
		TenantID:       "acme",
		Repo:           "repo-readback-a",
		Environment:    "prod",
		Namespace:      "acme-prod",
		Workload:       "api",
		ServiceAccount: "shared-sa",
		Digest:         "sha256:readback-a",
		Reasons:        []string{"workflow mismatch", "signature verification failed"},
		PolicyBundleID: "bundle-readback-a",
	})
	mustIngest(audit.Event{
		RequestID:   "req-readback-runtime",
		Component:   "runtime-agent",
		EventType:   audit.EventTypeRuntimeDriftResult,
		Decision:    audit.DecisionDeny,
		TenantID:    "acme",
		Repo:        "repo-readback-b",
		Environment: "prod",
		Workload:    "worker",
		Image:       "ghcr.io/acme/worker@sha256:readback-b",
		Digest:      "sha256:readback-b",
		DriftResult: "image_drift",
		Reasons:     []string{"image drift"},
	})
	mustIngest(audit.Event{
		RequestID:            "req-readback-auth",
		Component:            "runtime-agent",
		EventType:            audit.EventTypeRuntimeActiveStateObserved,
		Decision:             audit.DecisionDeny,
		TenantID:             "acme",
		Repo:                 "repo-readback-auth",
		Environment:          "prod",
		Namespace:            "acme-prod",
		Workload:             "auth-api",
		ServiceAccount:       "shared-sa",
		Digest:               "sha256:readback-auth",
		DriftResult:          "service_account_drift",
		DriftClasses:         []string{"service_account_drift"},
		Reasons:              []string{"service account drift"},
		ReconciliationStatus: "drift_detected",
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
	if len(list.Incidents) == 0 {
		t.Fatal("expected at least one incident for readback tests")
	}
	incidentID := list.Incidents[0].ID
	for _, incident := range list.Incidents {
		if containsString(incident.AffectedWorkloads, "api") {
			incidentID = incident.ID
			break
		}
	}

	defenseReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+url.PathEscape(incidentID)+"/defense-gaps?tenant_id=acme", nil)
	defenseRec := httptest.NewRecorder()
	handler.ServeHTTP(defenseRec, defenseReq)
	if defenseRec.Code != http.StatusOK {
		t.Fatalf("expected defense-gap 200, got %d: %s", defenseRec.Code, defenseRec.Body.String())
	}

	var defense defenseGapAssessment
	if err := json.NewDecoder(defenseRec.Body).Decode(&defense); err != nil {
		t.Fatalf("decode defense gap assessment: %v", err)
	}
	if defense.Readback.ResourceID == "" || defense.Readback.EvidenceHash == "" {
		t.Fatalf("expected defense-gap readback ref, got %#v", defense.Readback)
	}

	readbackPath := "/v1/readback/defense-gap/" + url.PathEscape(defense.Readback.ResourceID)
	readbackReq := httptest.NewRequest(http.MethodGet, readbackPath, nil)
	readbackRec := httptest.NewRecorder()
	handler.ServeHTTP(readbackRec, readbackReq)
	if readbackRec.Code != http.StatusOK {
		t.Fatalf("expected defense-gap readback 200, got %d: %s", readbackRec.Code, readbackRec.Body.String())
	}

	var defenseReadback defenseGapReadbackPayload
	if err := json.NewDecoder(readbackRec.Body).Decode(&defenseReadback); err != nil {
		t.Fatalf("decode defense-gap readback: %v", err)
	}
	if defenseReadback.EvidenceEnvelope.EvidenceHash != defense.Readback.EvidenceHash {
		t.Fatalf("expected stable defense-gap evidence hash, got %q want %q", defenseReadback.EvidenceEnvelope.EvidenceHash, defense.Readback.EvidenceHash)
	}
	if defenseReadback.SchemaVersion != readbackResponseSchemaVersion || defenseReadback.EvidenceEnvelope.SchemaVersion != readbackEnvelopeSchemaVersion {
		t.Fatalf("expected schema-versioned defense-gap readback, got %#v", defenseReadback)
	}
	if defenseReadback.TopologyContext == nil || defenseReadback.TopologyContext.PrimaryAffectedNode == nil {
		t.Fatalf("expected topology context on defense-gap readback, got %#v", defenseReadback.TopologyContext)
	}
	if defenseReadback.TopologyContext.BlastRadiusScore == 0 || len(defenseReadback.TopologyContext.TopRiskPathSummaries) == 0 {
		t.Fatalf("expected explainable topology summary on defense-gap readback, got %#v", defenseReadback.TopologyContext)
	}

	readbackRecAgain := httptest.NewRecorder()
	handler.ServeHTTP(readbackRecAgain, httptest.NewRequest(http.MethodGet, readbackPath, nil))
	if readbackRecAgain.Code != http.StatusOK {
		t.Fatalf("expected repeated defense-gap readback 200, got %d: %s", readbackRecAgain.Code, readbackRecAgain.Body.String())
	}
	var defenseReadbackAgain defenseGapReadbackPayload
	if err := json.NewDecoder(readbackRecAgain.Body).Decode(&defenseReadbackAgain); err != nil {
		t.Fatalf("decode repeated defense-gap readback: %v", err)
	}
	if defenseReadbackAgain.EvidenceEnvelope.EvidenceHash != defenseReadback.EvidenceEnvelope.EvidenceHash {
		t.Fatalf("expected repeated defense-gap readback hash to stay stable, got %q then %q", defenseReadback.EvidenceEnvelope.EvidenceHash, defenseReadbackAgain.EvidenceEnvelope.EvidenceHash)
	}
	if string(canonicalBytes(defenseReadback)) != string(canonicalBytes(defenseReadbackAgain)) {
		t.Fatalf("expected repeated defense-gap readback payload to stay deterministic")
	}

	replayReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+url.PathEscape(incidentID)+"/policy-replay?tenant_id=acme", nil)
	replayRec := httptest.NewRecorder()
	handler.ServeHTTP(replayRec, replayReq)
	if replayRec.Code != http.StatusOK {
		t.Fatalf("expected policy replay 200, got %d: %s", replayRec.Code, replayRec.Body.String())
	}

	var replay policyReplayAssessment
	if err := json.NewDecoder(replayRec.Body).Decode(&replay); err != nil {
		t.Fatalf("decode policy replay assessment: %v", err)
	}
	if replay.Readback.ResourceID == "" || replay.Readback.EvidenceHash == "" {
		t.Fatalf("expected policy replay readback ref, got %#v", replay.Readback)
	}

	replayReadbackReq := httptest.NewRequest(http.MethodGet, "/v1/readback/policy-replay/"+url.PathEscape(replay.Readback.ResourceID), nil)
	replayReadbackRec := httptest.NewRecorder()
	handler.ServeHTTP(replayReadbackRec, replayReadbackReq)
	if replayReadbackRec.Code != http.StatusOK {
		t.Fatalf("expected policy replay readback 200, got %d: %s", replayReadbackRec.Code, replayReadbackRec.Body.String())
	}

	var replayReadback policyReplayReadbackPayload
	if err := json.NewDecoder(replayReadbackRec.Body).Decode(&replayReadback); err != nil {
		t.Fatalf("decode policy replay readback: %v", err)
	}
	if replayReadback.EvidenceEnvelope.EvidenceHash != replay.Readback.EvidenceHash {
		t.Fatalf("expected stable policy replay evidence hash, got %q want %q", replayReadback.EvidenceEnvelope.EvidenceHash, replay.Readback.EvidenceHash)
	}
	if replayReadback.SchemaVersion != readbackResponseSchemaVersion {
		t.Fatalf("expected schema-versioned policy replay readback, got %#v", replayReadback)
	}

	weaknessReq := httptest.NewRequest(http.MethodGet, "/v1/ai/systemic-weaknesses?tenant_id=acme", nil)
	weaknessRec := httptest.NewRecorder()
	handler.ServeHTTP(weaknessRec, weaknessReq)
	if weaknessRec.Code != http.StatusOK {
		t.Fatalf("expected systemic weaknesses 200, got %d: %s", weaknessRec.Code, weaknessRec.Body.String())
	}

	var weaknesses systemicWeaknessResponse
	if err := json.NewDecoder(weaknessRec.Body).Decode(&weaknesses); err != nil {
		t.Fatalf("decode systemic weaknesses: %v", err)
	}
	if len(weaknesses.Weaknesses) == 0 || weaknesses.Weaknesses[0].Readback.ResourceID == "" {
		t.Fatalf("expected systemic weakness readback ref, got %#v", weaknesses)
	}

	weaknessReadbackReq := httptest.NewRequest(http.MethodGet, "/v1/readback/systemic-weakness/"+url.PathEscape(weaknesses.Weaknesses[0].Readback.ResourceID), nil)
	weaknessReadbackRec := httptest.NewRecorder()
	handler.ServeHTTP(weaknessReadbackRec, weaknessReadbackReq)
	if weaknessReadbackRec.Code != http.StatusOK {
		t.Fatalf("expected systemic weakness readback 200, got %d: %s", weaknessReadbackRec.Code, weaknessReadbackRec.Body.String())
	}

	var weaknessReadback systemicWeaknessReadbackPayload
	if err := json.NewDecoder(weaknessReadbackRec.Body).Decode(&weaknessReadback); err != nil {
		t.Fatalf("decode systemic weakness readback: %v", err)
	}
	if weaknessReadback.EvidenceEnvelope.EvidenceHash != weaknesses.Weaknesses[0].Readback.EvidenceHash {
		t.Fatalf("expected stable systemic weakness evidence hash, got %q want %q", weaknessReadback.EvidenceEnvelope.EvidenceHash, weaknesses.Weaknesses[0].Readback.EvidenceHash)
	}
	if weaknessReadback.SchemaVersion != readbackResponseSchemaVersion {
		t.Fatalf("expected schema-versioned systemic weakness readback, got %#v", weaknessReadback)
	}
}

func TestReadbackShareGrantsEnforceProjectionTamperAndExpiry(t *testing.T) {
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeDisabled)
	t.Setenv("CHANGELOCK_READBACK_GRANT_SECRET", "readback-secret")

	store := audit.NewMemoryStore()
	handler := newHandler(store, "memory")

	if _, err := store.Ingest(t.Context(), audit.Event{
		RequestID:      "req-share-1",
		Component:      "deploy-gate",
		EventType:      audit.EventTypeDeployGateDecision,
		Decision:       audit.DecisionDeny,
		TenantID:       "acme",
		Repo:           "repo-share-a",
		Environment:    "prod",
		Workload:       "api",
		Digest:         "sha256:share-a",
		Reasons:        []string{"workflow mismatch", "signature verification failed"},
		PolicyBundleID: "bundle-share-a",
	}); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}

	defenseReq := httptest.NewRequest(http.MethodGet, "/v1/ai/defense-gap-assessments?tenant_id=acme&incident_id=INC-0001", nil)
	defenseRec := httptest.NewRecorder()
	handler.ServeHTTP(defenseRec, defenseReq)
	if defenseRec.Code != http.StatusOK {
		listReq := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme", nil)
		listRec := httptest.NewRecorder()
		handler.ServeHTTP(listRec, listReq)
		var list incidentsResponse
		_ = json.NewDecoder(listRec.Body).Decode(&list)
		if len(list.Incidents) == 0 {
			t.Fatalf("expected at least one incident, got %d", len(list.Incidents))
		}
		defenseReq = httptest.NewRequest(http.MethodGet, "/v1/ai/defense-gap-assessments?tenant_id=acme&incident_id="+url.QueryEscape(list.Incidents[0].ID), nil)
		defenseRec = httptest.NewRecorder()
		handler.ServeHTTP(defenseRec, defenseReq)
	}
	if defenseRec.Code != http.StatusOK {
		t.Fatalf("expected defense-gap assessment 200, got %d: %s", defenseRec.Code, defenseRec.Body.String())
	}

	var defense defenseGapAssessment
	if err := json.NewDecoder(defenseRec.Body).Decode(&defense); err != nil {
		t.Fatalf("decode defense-gap assessment: %v", err)
	}

	body, err := json.Marshal(readbackGrantRequest{
		ResourceType:     defense.Readback.ResourceType,
		ResourceID:       defense.Readback.ResourceID,
		Audience:         incidentAudienceAuditorSafe,
		ExpiresInMinutes: 30,
		Purpose:          "test-share",
	})
	if err != nil {
		t.Fatalf("marshal share grant request: %v", err)
	}
	grantReq := httptest.NewRequest(http.MethodPost, "/v1/readback/grants", bytes.NewReader(body))
	grantReq.Header.Set("Content-Type", "application/json")
	grantRec := httptest.NewRecorder()
	handler.ServeHTTP(grantRec, grantReq)
	if grantRec.Code != http.StatusCreated {
		t.Fatalf("expected readback grant 201, got %d: %s", grantRec.Code, grantRec.Body.String())
	}

	var grant readbackGrantResponse
	if err := json.NewDecoder(grantRec.Body).Decode(&grant); err != nil {
		t.Fatalf("decode readback grant response: %v", err)
	}
	if grant.ShareURL == "" || grant.GrantID == "" {
		t.Fatalf("expected populated share grant response, got %#v", grant)
	}

	shareReq := httptest.NewRequest(http.MethodGet, grant.ShareURL, nil)
	shareRec := httptest.NewRecorder()
	handler.ServeHTTP(shareRec, shareReq)
	if shareRec.Code != http.StatusOK {
		t.Fatalf("expected share readback 200, got %d: %s", shareRec.Code, shareRec.Body.String())
	}

	var share defenseGapReadbackPayload
	if err := json.NewDecoder(shareRec.Body).Decode(&share); err != nil {
		t.Fatalf("decode share readback: %v", err)
	}
	if share.ProjectionAudience != incidentAudienceAuditorSafe {
		t.Fatalf("expected auditor-safe projection, got %#v", share)
	}
	if len(share.Payload.DefenseGaps) > 0 && len(share.Payload.DefenseGaps[0].EvidenceRefs) != 0 {
		t.Fatalf("expected auditor-safe share to redact evidence refs, got %#v", share.Payload.DefenseGaps[0])
	}

	customerBody, err := json.Marshal(readbackGrantRequest{
		ResourceType:     defense.Readback.ResourceType,
		ResourceID:       defense.Readback.ResourceID,
		Audience:         incidentAudienceCustomerSafe,
		ExpiresInMinutes: 30,
		Purpose:          "test-customer-share",
	})
	if err != nil {
		t.Fatalf("marshal customer-safe share grant request: %v", err)
	}
	customerGrantReq := httptest.NewRequest(http.MethodPost, "/v1/readback/grants", bytes.NewReader(customerBody))
	customerGrantReq.Header.Set("Content-Type", "application/json")
	customerGrantRec := httptest.NewRecorder()
	handler.ServeHTTP(customerGrantRec, customerGrantReq)
	if customerGrantRec.Code != http.StatusCreated {
		t.Fatalf("expected customer-safe readback grant 201, got %d: %s", customerGrantRec.Code, customerGrantRec.Body.String())
	}

	var customerGrant readbackGrantResponse
	if err := json.NewDecoder(customerGrantRec.Body).Decode(&customerGrant); err != nil {
		t.Fatalf("decode customer-safe readback grant response: %v", err)
	}
	if customerGrant.Audience != incidentAudienceCustomerSafe {
		t.Fatalf("expected customer-safe grant audience, got %#v", customerGrant)
	}

	customerShareReq := httptest.NewRequest(http.MethodGet, customerGrant.ShareURL, nil)
	customerShareRec := httptest.NewRecorder()
	handler.ServeHTTP(customerShareRec, customerShareReq)
	if customerShareRec.Code != http.StatusOK {
		t.Fatalf("expected customer-safe share readback 200, got %d: %s", customerShareRec.Code, customerShareRec.Body.String())
	}

	var customerShare defenseGapReadbackPayload
	if err := json.NewDecoder(customerShareRec.Body).Decode(&customerShare); err != nil {
		t.Fatalf("decode customer-safe share readback: %v", err)
	}
	if customerShare.ProjectionAudience != incidentAudienceCustomerSafe {
		t.Fatalf("expected customer-safe projection, got %#v", customerShare)
	}

	tamperedReq := httptest.NewRequest(http.MethodGet, grant.ShareURL+"tampered", nil)
	tamperedRec := httptest.NewRecorder()
	handler.ServeHTTP(tamperedRec, tamperedReq)
	if tamperedRec.Code != http.StatusUnauthorized {
		t.Fatalf("expected tampered grant 401, got %d: %s", tamperedRec.Code, tamperedRec.Body.String())
	}

	expiredToken, err := signReadbackGrant(signedReadbackGrant{
		Version:              readbackEnvelopeSchemaVersion,
		ResourceType:         defense.Readback.ResourceType,
		ResourceID:           defense.Readback.ResourceID,
		Audience:             incidentAudienceAuditorSafe,
		ExpectedEvidenceHash: defense.Readback.EvidenceHash,
		ExpiresAt:            time.Now().UTC().Add(-1 * time.Minute),
		Purpose:              "expired-test",
	}, readbackGrantSecret())
	if err != nil {
		t.Fatalf("sign expired readback grant: %v", err)
	}

	expiredReq := httptest.NewRequest(http.MethodGet, "/s/"+expiredToken, nil)
	expiredRec := httptest.NewRecorder()
	handler.ServeHTTP(expiredRec, expiredReq)
	if expiredRec.Code != http.StatusUnauthorized {
		t.Fatalf("expected expired grant 401, got %d: %s", expiredRec.Code, expiredRec.Body.String())
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWave2CProductionHardeningSurfaces(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")
	t.Setenv("CHANGELOCK_FEDERATION_SIGNING_SEED", "federation-seed")

	fixture := forensicsTestFixture(t)

	sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)
	registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-wave-2c",
		Organization:      "bounded-partner",
		Region:            "eu-central",
		TrustDomain:       "partner",
		PublicKeys:        []string{"peer-wave-2c-pub"},
		Capabilities:      []string{"sealed_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
	})

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

	findingsReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/findings?tenant_id=acme&environment=prod&limit=20", nil)
	findingsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	findingsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(findingsRec, findingsReq)
	if findingsRec.Code != http.StatusOK {
		t.Fatalf("expected runtime findings 200, got %d: %s", findingsRec.Code, findingsRec.Body.String())
	}
	var findings runtimeFindingsResponse
	if err := json.NewDecoder(findingsRec.Body).Decode(&findings); err != nil {
		t.Fatalf("decode runtime findings: %v", err)
	}
	binaryFinding := findRuntimeFinding(t, findings.Items, runtimeFindingUnknownBinaryExec, "edge-gateway")
	hardeningReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/hardening/quarantine?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-2C-1"}`),
	)
	hardeningReq.Header.Set("Authorization", "Bearer operator-demo-token")
	hardeningReq.Header.Set("Content-Type", "application/json")
	hardeningRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(hardeningRec, hardeningReq)
	if hardeningRec.Code != http.StatusOK {
		t.Fatalf("expected hardening quarantine 200, got %d: %s", hardeningRec.Code, hardeningRec.Body.String())
	}

	t.Run("handoff quality gates", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/handoff/quality-gates?tenant_id=acme&environment=prod", nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		fixture.handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected handoff quality gates 200, got %d: %s", rec.Code, rec.Body.String())
		}
		var response handoffQualityGatesResponse
		if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
			t.Fatalf("decode handoff quality gates: %v", err)
		}
		if response.SchemaVersion != handoffQualitySchemaVersion || len(response.Gates) == 0 {
			t.Fatalf("expected schema-versioned handoff quality gates, got %#v", response)
		}
		if !response.OfflineVerifySupported || response.ActiveSignerBackend == "" {
			t.Fatalf("expected production handoff model metadata, got %#v", response)
		}
	})

	t.Run("federation resilience", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/federation/resilience?tenant_id=acme&environment=prod", nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		fixture.handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected federation resilience 200, got %d: %s", rec.Code, rec.Body.String())
		}
		var response federationResilienceResponse
		if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
			t.Fatalf("decode federation resilience: %v", err)
		}
		if response.SchemaVersion != federationResilienceSchemaVersion || len(response.Peers) == 0 {
			t.Fatalf("expected schema-versioned federation resilience response, got %#v", response)
		}
	})

	t.Run("validation readiness", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/validation/readiness?tenant_id=acme&environment=prod", nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		fixture.handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected validation readiness 200, got %d: %s", rec.Code, rec.Body.String())
		}
		var response validationReadinessResponse
		if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
			t.Fatalf("decode validation readiness: %v", err)
		}
		if response.SchemaVersion != validationReadinessSchemaVersion || response.ResourceQuota.MaxScenariosPerRun != validationMaxScenarioSelection {
			t.Fatalf("expected schema-versioned validation readiness, got %#v", response)
		}
	})

	t.Run("self audit", func(t *testing.T) {
		summaryReq := httptest.NewRequest(http.MethodGet, "/v1/self-audit/summary", nil)
		summaryReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
		summaryRec := httptest.NewRecorder()
		fixture.handler.ServeHTTP(summaryRec, summaryReq)
		if summaryRec.Code != http.StatusOK {
			t.Fatalf("expected self-audit summary 200, got %d: %s", summaryRec.Code, summaryRec.Body.String())
		}
		var summary selfAuditSummaryResponse
		if err := json.NewDecoder(summaryRec.Body).Decode(&summary); err != nil {
			t.Fatalf("decode self-audit summary: %v", err)
		}
		if summary.SchemaVersion != selfAuditSummarySchemaVersion {
			t.Fatalf("expected schema-versioned self-audit summary, got %#v", summary)
		}
		if summary.CountsByCategory["signing_change"] == 0 || summary.CountsByCategory["validation_change"] == 0 || summary.CountsByCategory["federation_change"] == 0 {
			t.Fatalf("expected self-audit summary to track signing, validation, and federation changes, got %#v", summary.CountsByCategory)
		}

		eventsReq := httptest.NewRequest(http.MethodGet, "/v1/self-audit/events", nil)
		eventsReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
		eventsRec := httptest.NewRecorder()
		fixture.handler.ServeHTTP(eventsRec, eventsReq)
		if eventsRec.Code != http.StatusOK {
			t.Fatalf("expected self-audit events 200, got %d: %s", eventsRec.Code, eventsRec.Body.String())
		}
		var events selfAuditEventsResponse
		if err := json.NewDecoder(eventsRec.Body).Decode(&events); err != nil {
			t.Fatalf("decode self-audit events: %v", err)
		}
		if events.SchemaVersion != selfAuditEventsSchemaVersion || len(events.Events) == 0 {
			t.Fatalf("expected schema-versioned self-audit events, got %#v", events)
		}
	})
}

func TestFederationResilienceGuardsProofTraffic(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")
	t.Setenv("CHANGELOCK_FEDERATION_SIGNING_SEED", "federation-seed")

	fixture := forensicsTestFixture(t)
	sealed := sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)
	staleTime := time.Now().UTC().Add(-48 * time.Hour)
	registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-stale",
		Organization:      "stale-partner",
		Region:            "eu-central",
		TrustDomain:       "partner",
		PublicKeys:        []string{"peer-stale-pub"},
		Capabilities:      []string{"sealed_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
		LastSeen:          &staleTime,
	})
	staleReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/federation/proof-request",
		bytes.NewBufferString(`{"peer_id":"peer-stale","package_id":"`+sealed.PackageID+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`),
	)
	staleReq.Header.Set("Authorization", "Bearer operator-demo-token")
	staleReq.Header.Set("Content-Type", "application/json")
	staleRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(staleRec, staleReq)
	if staleRec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected stale peer proof request to open circuit with 503, got %d: %s", staleRec.Code, staleRec.Body.String())
	}

	registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-active",
		Organization:      "active-partner",
		Region:            "eu-central",
		TrustDomain:       "partner",
		PublicKeys:        []string{"peer-active-pub"},
		Capabilities:      []string{"sealed_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
	})

	for i := 0; i < federationResilienceMaxRequestsPerHour; i++ {
		req := httptest.NewRequest(
			http.MethodPost,
			"/v1/federation/proof-request",
			bytes.NewBufferString(`{"peer_id":"peer-active","package_id":"`+sealed.PackageID+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`),
		)
		req.Header.Set("Authorization", "Bearer operator-demo-token")
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		fixture.handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected bounded proof request %d to succeed, got %d: %s", i, rec.Code, rec.Body.String())
		}
	}

	limitedReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/federation/proof-request",
		bytes.NewBufferString(`{"peer_id":"peer-active","package_id":"`+sealed.PackageID+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`),
	)
	limitedReq.Header.Set("Authorization", "Bearer operator-demo-token")
	limitedReq.Header.Set("Content-Type", "application/json")
	limitedRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(limitedRec, limitedReq)
	if limitedRec.Code != http.StatusTooManyRequests {
		t.Fatalf("expected active peer proof request to rate-limit with 429, got %d: %s", limitedRec.Code, limitedRec.Body.String())
	}
}

func TestValidationReadinessEnforcesScenarioQuota(t *testing.T) {
	fixture := forensicsTestFixture(t)
	scenarios := validationScenarioRegistry()
	if len(scenarios) <= validationMaxScenarioSelection {
		t.Fatalf("expected scenario registry to exceed quota threshold, got %d", len(scenarios))
	}
	ids := make([]string, 0, validationMaxScenarioSelection+1)
	for _, scenario := range scenarios[:validationMaxScenarioSelection+1] {
		ids = append(ids, scenario.ScenarioID)
	}
	body, err := json.Marshal(map[string]any{"scenario_ids": ids})
	if err != nil {
		t.Fatalf("marshal validation quota request: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/v1/validation/execute?tenant_id=acme&environment=prod", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer operator-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected validation quota rejection 400, got %d: %s", rec.Code, rec.Body.String())
	}
}

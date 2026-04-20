package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
)

func TestFederationProofExchangeDecisionAndPolicyState(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")
	t.Setenv("CHANGELOCK_FEDERATION_SIGNING_SEED", "federation-seed")

	fixture := forensicsTestFixture(t)

	peer := registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-eu",
		Organization:      "Acme Europe",
		Region:            "eu-west",
		Cluster:           "eu-west-1",
		TrustDomain:       "acme.internal",
		Endpoint:          "https://peer-eu.example.invalid",
		PublicKeys:        []string{"pub-eu-1"},
		Capabilities:      []string{"sealed_handoff", "policy_sync", "forensics_handoff"},
		PolicyRole:        federationPolicyRoleLeader,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe, incidentAudienceCustomerSafe},
	})

	sealed := sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)
	if sealed.Bundle.ManifestHash == "" {
		t.Fatalf("expected sealed handoff manifest hash, got %#v", sealed)
	}
	bundleBytes := downloadHandoffBundleForTest(t, fixture.handler, sealed.PackageID)
	bundleBase64 := base64.StdEncoding.EncodeToString(bundleBytes)

	requestReq := httptest.NewRequest(http.MethodPost, "/v1/federation/proof-request", bytes.NewBufferString(`{"peer_id":"`+peer.PeerID+`","package_id":"`+sealed.PackageID+`","subject_type":"package","subject_ref":"`+sealed.PackageID+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`))
	requestReq.Header.Set("Authorization", "Bearer operator-demo-token")
	requestReq.Header.Set("Content-Type", "application/json")
	requestRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(requestRec, requestReq)
	if requestRec.Code != http.StatusOK {
		t.Fatalf("expected federation proof request 200, got %d: %s", requestRec.Code, requestRec.Body.String())
	}

	var exchange federationProofExchangeResult
	if err := json.NewDecoder(requestRec.Body).Decode(&exchange); err != nil {
		t.Fatalf("decode federation proof exchange: %v", err)
	}
	if exchange.Response.ManifestHash != sealed.Bundle.ManifestHash || len(exchange.Response.ReadbackRefs) == 0 {
		t.Fatalf("expected readback lineage and manifest hash in proof response, got %#v", exchange)
	}
	if exchange.SchemaVersion != federationProofExchangeSchemaVersion || exchange.Response.SchemaVersion != federationProofResponseSchemaVersion {
		t.Fatalf("expected schema-versioned federation proof exchange, got %#v", exchange)
	}

	verifyReq := httptest.NewRequest(http.MethodPost, "/v1/federation/proof-verify", bytes.NewBufferString(`{"peer_id":"`+peer.PeerID+`","bundle_base64":"`+bundleBase64+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`))
	verifyReq.Header.Set("Authorization", "Bearer operator-demo-token")
	verifyReq.Header.Set("Content-Type", "application/json")
	verifyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(verifyRec, verifyReq)
	if verifyRec.Code != http.StatusOK {
		t.Fatalf("expected federation proof verify 200, got %d: %s", verifyRec.Code, verifyRec.Body.String())
	}

	var accepted federationProofVerifyResult
	if err := json.NewDecoder(verifyRec.Body).Decode(&accepted); err != nil {
		t.Fatalf("decode accepted federation verify: %v", err)
	}
	if accepted.Decision.Decision != federationDecisionAccepted {
		t.Fatalf("expected accepted federated trust decision, got %#v", accepted.Decision)
	}
	if accepted.SchemaVersion != federationProofVerifySchemaVersion || accepted.Decision.SchemaVersion != federationTrustDecisionSchemaVersion {
		t.Fatalf("expected schema-versioned federation verify response, got %#v", accepted)
	}

	syncReq := httptest.NewRequest(http.MethodPost, "/v1/federation/policy-sync", bytes.NewBufferString(`{"leader_peer":"`+peer.PeerID+`","global_policy_root":"sha256:global-root","local_policy_root":"sha256:regional-root","local_overrides":["regional_compliance_hold"],"inherited_rules":["base:trusted-remote-proof"],"remote_policy_version":"2026.04"}`))
	syncReq.Header.Set("Authorization", "Bearer operator-demo-token")
	syncReq.Header.Set("Content-Type", "application/json")
	syncRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(syncRec, syncReq)
	if syncRec.Code != http.StatusOK {
		t.Fatalf("expected federation policy sync 200, got %d: %s", syncRec.Code, syncRec.Body.String())
	}

	var state policyFederationState
	if err := json.NewDecoder(syncRec.Body).Decode(&state); err != nil {
		t.Fatalf("decode policy sync state: %v", err)
	}
	if state.SyncStatus != federationSyncStatusSyncedWithOverrides || len(state.LocalOverrides) == 0 {
		t.Fatalf("expected synced-with-overrides state, got %#v", state)
	}

	verifyReq = httptest.NewRequest(http.MethodPost, "/v1/federation/proof-verify", bytes.NewBufferString(`{"peer_id":"`+peer.PeerID+`","bundle_base64":"`+bundleBase64+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`))
	verifyReq.Header.Set("Authorization", "Bearer operator-demo-token")
	verifyReq.Header.Set("Content-Type", "application/json")
	verifyRec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(verifyRec, verifyReq)
	if verifyRec.Code != http.StatusOK {
		t.Fatalf("expected second federation proof verify 200, got %d: %s", verifyRec.Code, verifyRec.Body.String())
	}
	if err := json.NewDecoder(verifyRec.Body).Decode(&accepted); err != nil {
		t.Fatalf("decode override federation verify: %v", err)
	}
	if accepted.Decision.Decision != federationDecisionAcceptedWithOverrides {
		t.Fatalf("expected accepted_with_local_overrides, got %#v", accepted.Decision)
	}

	oldLastSeen := time.Now().UTC().Add(-7 * time.Hour).Format(time.RFC3339)
	stalePeer := registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-stale",
		Organization:      "Acme Supplier",
		Region:            "us-east",
		Cluster:           "us-east-1",
		TrustDomain:       "supplier.example",
		Endpoint:          "https://peer-stale.example.invalid",
		PublicKeys:        []string{"pub-stale-1"},
		Capabilities:      []string{"sealed_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
		LastSeen:          mustParseTimeRFC3339(t, oldLastSeen),
	})

	staleVerifyReq := httptest.NewRequest(http.MethodPost, "/v1/federation/proof-verify", bytes.NewBufferString(`{"peer_id":"`+stalePeer.PeerID+`","bundle_base64":"`+bundleBase64+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`))
	staleVerifyReq.Header.Set("Authorization", "Bearer operator-demo-token")
	staleVerifyReq.Header.Set("Content-Type", "application/json")
	staleVerifyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(staleVerifyRec, staleVerifyReq)
	if staleVerifyRec.Code != http.StatusOK {
		t.Fatalf("expected stale federation proof verify 200, got %d: %s", staleVerifyRec.Code, staleVerifyRec.Body.String())
	}
	var staleDecision federationProofVerifyResult
	if err := json.NewDecoder(staleVerifyRec.Body).Decode(&staleDecision); err != nil {
		t.Fatalf("decode stale federation verify: %v", err)
	}
	if staleDecision.Decision.Decision != federationDecisionRejectedStale {
		t.Fatalf("expected rejected_stale, got %#v", staleDecision.Decision)
	}

	conflictPeer := registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-b2b",
		Organization:      "Supplier Co",
		Region:            "eu-central",
		Cluster:           "supplier-cluster",
		TrustDomain:       "supplier.b2b",
		Endpoint:          "https://supplier.example.invalid",
		PublicKeys:        []string{"pub-supplier-1"},
		Capabilities:      []string{"sealed_handoff", "supplier_proof"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceCustomerSafe},
	})

	conflictVerifyReq := httptest.NewRequest(http.MethodPost, "/v1/federation/proof-verify", bytes.NewBufferString(`{"peer_id":"`+conflictPeer.PeerID+`","bundle_base64":"`+bundleBase64+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`))
	conflictVerifyReq.Header.Set("Authorization", "Bearer operator-demo-token")
	conflictVerifyReq.Header.Set("Content-Type", "application/json")
	conflictVerifyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(conflictVerifyRec, conflictVerifyReq)
	if conflictVerifyRec.Code != http.StatusOK {
		t.Fatalf("expected conflict federation proof verify 200, got %d: %s", conflictVerifyRec.Code, conflictVerifyRec.Body.String())
	}
	var conflictDecision federationProofVerifyResult
	if err := json.NewDecoder(conflictVerifyRec.Body).Decode(&conflictDecision); err != nil {
		t.Fatalf("decode conflict federation verify: %v", err)
	}
	if conflictDecision.Decision.Decision != federationDecisionRejectedPolicyConflict {
		t.Fatalf("expected rejected_policy_conflict, got %#v", conflictDecision.Decision)
	}

	historyReq := httptest.NewRequest(http.MethodGet, "/v1/federation/proof-history", nil)
	historyReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	historyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(historyRec, historyReq)
	if historyRec.Code != http.StatusOK {
		t.Fatalf("expected federation proof history 200, got %d: %s", historyRec.Code, historyRec.Body.String())
	}
	var history federationProofHistoryResponse
	if err := json.NewDecoder(historyRec.Body).Decode(&history); err != nil {
		t.Fatalf("decode federation history: %v", err)
	}
	if len(history.Items) < 3 {
		t.Fatalf("expected recorded federation history items, got %#v", history)
	}
	if history.SchemaVersion != federationProofHistorySchemaVersion {
		t.Fatalf("expected schema-versioned federation history, got %#v", history)
	}

	historyRecRepeat := httptest.NewRecorder()
	fixture.handler.ServeHTTP(historyRecRepeat, historyReq)
	if historyRecRepeat.Code != http.StatusOK {
		t.Fatalf("expected repeated federation proof history 200, got %d: %s", historyRecRepeat.Code, historyRecRepeat.Body.String())
	}
	var historyRepeat federationProofHistoryResponse
	if err := json.NewDecoder(historyRecRepeat.Body).Decode(&historyRepeat); err != nil {
		t.Fatalf("decode repeated federation history: %v", err)
	}
	if string(canonicalJSONMust(history)) != string(canonicalJSONMust(historyRepeat)) {
		t.Fatalf("expected federation proof history to stay deterministic for the same input")
	}

	anchorsReq := httptest.NewRequest(http.MethodGet, "/v1/federation/anchors", nil)
	anchorsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	anchorsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(anchorsRec, anchorsReq)
	if anchorsRec.Code != http.StatusOK {
		t.Fatalf("expected federation anchors 200, got %d: %s", anchorsRec.Code, anchorsRec.Body.String())
	}
	var anchors federationAnchorsResponse
	if err := json.NewDecoder(anchorsRec.Body).Decode(&anchors); err != nil {
		t.Fatalf("decode federation anchors: %v", err)
	}
	if len(anchors.Items) == 0 || anchors.Items[0].AuditRootHash == "" {
		t.Fatalf("expected local federation anchors, got %#v", anchors)
	}
	if anchors.SchemaVersion != federationAnchorsSchemaVersion {
		t.Fatalf("expected schema-versioned federation anchors, got %#v", anchors)
	}

	anchorsRecRepeat := httptest.NewRecorder()
	fixture.handler.ServeHTTP(anchorsRecRepeat, anchorsReq)
	if anchorsRecRepeat.Code != http.StatusOK {
		t.Fatalf("expected repeated federation anchors 200, got %d: %s", anchorsRecRepeat.Code, anchorsRecRepeat.Body.String())
	}
	var anchorsRepeat federationAnchorsResponse
	if err := json.NewDecoder(anchorsRecRepeat.Body).Decode(&anchorsRepeat); err != nil {
		t.Fatalf("decode repeated federation anchors: %v", err)
	}
	if string(canonicalJSONMust(anchors)) != string(canonicalJSONMust(anchorsRepeat)) {
		t.Fatalf("expected federation anchors to stay deterministic for the same input")
	}

	globalReq := httptest.NewRequest(http.MethodGet, "/v1/federation/global-view", nil)
	globalReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	globalRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(globalRec, globalReq)
	if globalRec.Code != http.StatusOK {
		t.Fatalf("expected federation global view 200, got %d: %s", globalRec.Code, globalRec.Body.String())
	}
	var global federationGlobalView
	if err := json.NewDecoder(globalRec.Body).Decode(&global); err != nil {
		t.Fatalf("decode federation global view: %v", err)
	}
	if global.VerifiedArtifactsReused == 0 || len(global.StalePeers) == 0 {
		t.Fatalf("expected reused artifacts and stale peer signal in global view, got %#v", global)
	}
	if global.SchemaVersion != federationGlobalViewSchemaVersion {
		t.Fatalf("expected schema-versioned federation global view, got %#v", global)
	}
}

func TestFederationRecommendations(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")
	t.Setenv("CHANGELOCK_FEDERATION_SIGNING_SEED", "federation-seed")

	fixture := forensicsTestFixture(t)
	registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-stale",
		Organization:      "Supplier",
		Region:            "us-east",
		Cluster:           "cluster-stale",
		TrustDomain:       "supplier.example",
		Endpoint:          "https://supplier.example.invalid",
		PublicKeys:        []string{"pub-stale-1"},
		Capabilities:      []string{"sealed_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
		LastSeen:          timePointer(time.Now().UTC().Add(-8 * time.Hour)),
	})

	req := httptest.NewRequest(http.MethodGet, "/v1/recommendations?tenant_id=acme&environment=prod&source_type=federation_signal", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected federation recommendations 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var response recommendationListResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode federation recommendations: %v", err)
	}
	if len(response.Recommendations) == 0 {
		t.Fatalf("expected federation recommendation candidate, got %#v", response)
	}
	if response.Recommendations[0].SourceType != "federation_signal" || len(response.Recommendations[0].VerificationPlan) == 0 {
		t.Fatalf("expected federation signal recommendation, got %#v", response.Recommendations[0])
	}
}

func registerFederationPeerForTest(t testing.TB, handler http.Handler, request federationPeerRequest) federationPeer {
	t.Helper()
	body, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal federation peer request: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/v1/federation/peers", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer operator-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected federation peer registration 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var peer federationPeer
	if err := json.NewDecoder(rec.Body).Decode(&peer); err != nil {
		t.Fatalf("decode federation peer: %v", err)
	}
	return peer
}

func sealFederationHandoffForTest(t testing.TB, handler http.Handler, audience string) handoffSealResponse {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/v1/handoff/seal?tenant_id=acme&environment=prod", bytes.NewBufferString(`{"audience":"`+audience+`","include_forensics":true,"co_sign_mode":"system_only"}`))
	req.Header.Set("Authorization", "Bearer operator-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected handoff seal 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var response handoffSealResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode handoff seal: %v", err)
	}
	return response
}

func downloadHandoffBundleForTest(t testing.TB, handler http.Handler, packageID string) []byte {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, "/v1/handoff/"+packageID+"/download", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected handoff download 200, got %d: %s", rec.Code, rec.Body.String())
	}
	return rec.Body.Bytes()
}

func mustParseTimeRFC3339(t *testing.T, value string) *time.Time {
	t.Helper()
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		t.Fatalf("parse RFC3339 time: %v", err)
	}
	return &parsed
}

func federationTestHandler(t *testing.T) http.Handler {
	t.Helper()
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())
	store := audit.NewMemoryStore()
	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	return newHandlerWithAuth(store, "memory", authConfig)
}

func seedFederationEvent(t *testing.T, store audit.Store, event audit.Event) {
	t.Helper()
	if _, err := store.Ingest(context.Background(), event); err != nil {
		t.Fatalf("Ingest() error = %v", err)
	}
}

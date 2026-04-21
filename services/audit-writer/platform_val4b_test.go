package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestB2BSupplierOnboardingHandler(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	peer := registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "supplier-west-4b",
		Organization:      "Supplier West",
		Region:            "eu-west",
		Cluster:           "supplier-west-1",
		TrustDomain:       "supplier-west.example",
		Endpoint:          "https://supplier-west.example.invalid",
		PublicKeys:        []string{"pub-supplier-west-1"},
		Capabilities:      []string{"sealed_handoff", "supplier_proof"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe, incidentAudienceCustomerSafe},
	})

	req := httptest.NewRequest(http.MethodGet, "/v1/b2b/suppliers/onboarding?tenant_id=acme", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected supplier onboarding 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response b2bSupplierOnboardingResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode supplier onboarding response: %v", err)
	}
	if response.SchemaVersion != b2bSupplierOnboardingSchema {
		t.Fatalf("expected supplier onboarding schema, got %#v", response)
	}
	if len(response.LocalPolicyOverview) == 0 || !containsString(response.AcceptedProofFormats, federationProofTypeHandoff) {
		t.Fatalf("expected bounded onboarding policy overview, got %#v", response)
	}
	item := findB2BSupplierOnboardingItem(t, response.Items, peer.PeerID)
	if item.PolicyRole != federationPolicyRoleSupplier || len(item.AcceptedProofFormats) == 0 {
		t.Fatalf("expected supplier onboarding item with accepted proof formats, got %#v", item)
	}
	if len(item.LocalAdmissibilityPolicy) == 0 || len(item.RevocationAndDistrust) == 0 {
		t.Fatalf("expected local admissibility and distrust semantics, got %#v", item)
	}
}

func TestB2BSealedProofAcceptanceHandler(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")
	t.Setenv("CHANGELOCK_FEDERATION_SIGNING_SEED", "federation-seed")

	fixture := forensicsTestFixture(t)
	peer := registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-proof-4b",
		Organization:      "Proof Partner",
		Region:            "eu-central",
		Cluster:           "proof-eu-1",
		TrustDomain:       "proof.partner.example",
		Endpoint:          "https://proof-partner.example.invalid",
		PublicKeys:        []string{"pub-proof-1"},
		Capabilities:      []string{"sealed_handoff", "forensics_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
	})
	sealed := sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)
	bundleBase64 := base64.StdEncoding.EncodeToString(downloadHandoffBundleForTest(t, fixture.handler, sealed.PackageID))

	getReq := httptest.NewRequest(http.MethodGet, "/v1/b2b/sealed-proof/acceptance?tenant_id=acme", nil)
	getReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	getRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(getRec, getReq)
	if getRec.Code != http.StatusOK {
		t.Fatalf("expected proof acceptance contract 200, got %d: %s", getRec.Code, getRec.Body.String())
	}

	var contract b2bSealedProofAcceptanceResponse
	if err := json.NewDecoder(getRec.Body).Decode(&contract); err != nil {
		t.Fatalf("decode proof acceptance contract: %v", err)
	}
	if contract.SchemaVersion != b2bProofAcceptanceSchema || !contract.OfflineVerifySupported {
		t.Fatalf("expected bounded proof acceptance contract, got %#v", contract)
	}

	postReq := httptest.NewRequest(http.MethodPost, "/v1/b2b/sealed-proof/acceptance", bytes.NewBufferString(`{"peer_id":"`+peer.PeerID+`","bundle_base64":"`+bundleBase64+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`))
	postReq.Header.Set("Authorization", "Bearer operator-demo-token")
	postReq.Header.Set("Content-Type", "application/json")
	postRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(postRec, postReq)
	if postRec.Code != http.StatusOK {
		t.Fatalf("expected proof acceptance evaluation 200, got %d: %s", postRec.Code, postRec.Body.String())
	}

	var evaluation b2bSealedProofAcceptanceEvaluation
	if err := json.NewDecoder(postRec.Body).Decode(&evaluation); err != nil {
		t.Fatalf("decode proof acceptance evaluation: %v", err)
	}
	if evaluation.SchemaVersion != b2bProofAcceptanceEvalSchema || evaluation.PeerID != peer.PeerID {
		t.Fatalf("expected proof acceptance evaluation schema and peer, got %#v", evaluation)
	}
	if evaluation.LocalDecision.Decision != federationDecisionAccepted || evaluation.LocalVerification.OverallStatus == "" {
		t.Fatalf("expected locally accepted partner proof with verification result, got %#v", evaluation)
	}
	if len(evaluation.AcceptanceNarrative) == 0 || !evaluation.OfflineVerifySupported {
		t.Fatalf("expected bounded local verification narrative, got %#v", evaluation)
	}
}

func TestB2BDisclosureProfilesHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/b2b/disclosure-profiles?tenant_id=acme", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected disclosure profiles 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response b2bDisclosureProfilesResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode disclosure profiles: %v", err)
	}
	if response.SchemaVersion != b2bDisclosureProfilesSchema || len(response.Profiles) < 3 {
		t.Fatalf("expected disclosure profile contract, got %#v", response)
	}
	if !hasB2BDisclosureProfile(response.Profiles, "sealed_proof_only") || !hasB2BDisclosureProfile(response.Profiles, incidentAudienceCustomerSafe) {
		t.Fatalf("expected sealed-proof and customer-safe profiles, got %#v", response.Profiles)
	}
	if len(response.SelectiveDisclosurePolicy) == 0 {
		t.Fatalf("expected selective disclosure policy semantics, got %#v", response)
	}
}

func TestB2BCustomerBundleHandler(t *testing.T) {
	t.Setenv("CHANGELOCK_TRUST_PUBLICATION_MODE", "preview")
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	seedTrustScorecardStore(t, fixture.store)
	sealed := sealFederationHandoffForTest(t, fixture.handler, incidentAudienceCustomerSafe)

	req := httptest.NewRequest(http.MethodGet, "/v1/b2b/customer-bundles?tenant_id=acme&repo=demo/app&package_id="+sealed.PackageID, nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected customer bundle 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response b2bCustomerTrustBundleResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode customer bundle: %v", err)
	}
	if response.SchemaVersion != b2bCustomerBundleSchema || response.ProfileID != incidentAudienceCustomerSafe {
		t.Fatalf("expected customer bundle schema and profile, got %#v", response)
	}
	if response.PublicView == nil || len(response.BoundedTrustIndicators) == 0 {
		t.Fatalf("expected published trust indicators in customer bundle, got %#v", response)
	}
	if !containsString(response.MachineVerifiablePaths, "/v1/trust/published") || response.SealedProofVerificationPath == "" {
		t.Fatalf("expected machine-verifiable trust and proof paths, got %#v", response)
	}
}

func TestB2BConsortiumReadinessHandler(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-leader-4b",
		Organization:      "Acme Europe",
		Region:            "eu-west",
		Cluster:           "eu-west-1",
		TrustDomain:       "acme.internal",
		Endpoint:          "https://peer-leader-4b.example.invalid",
		PublicKeys:        []string{"pub-leader-4b"},
		Capabilities:      []string{"sealed_handoff", "policy_sync"},
		PolicyRole:        federationPolicyRoleLeader,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe, incidentAudienceCustomerSafe},
	})
	registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-stale-4b",
		Organization:      "Supplier Stale",
		Region:            "us-east",
		Cluster:           "supplier-us-1",
		TrustDomain:       "supplier.example",
		Endpoint:          "https://peer-stale-4b.example.invalid",
		PublicKeys:        []string{"pub-stale-4b"},
		Capabilities:      []string{"sealed_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
		LastSeen:          mustParseTimeRFC3339(t, time.Now().UTC().Add(-8*time.Hour).Format(time.RFC3339)),
	})

	req := httptest.NewRequest(http.MethodGet, "/v1/b2b/consortium-readiness?tenant_id=acme", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected consortium readiness 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response b2bConsortiumReadinessResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode consortium readiness: %v", err)
	}
	if response.SchemaVersion != b2bConsortiumReadinessSchema || response.ReadinessState != "consortium_degraded" {
		t.Fatalf("expected degraded consortium readiness contract, got %#v", response)
	}
	if response.PeerSummary.TotalPeers < 2 || response.PeerSummary.StalePeers == 0 {
		t.Fatalf("expected active and stale peers in consortium summary, got %#v", response.PeerSummary)
	}
	if len(response.SharedTrustAnchorReadiness) == 0 || len(response.LocalOverrideAndDistrust) == 0 {
		t.Fatalf("expected bounded readiness and local override semantics, got %#v", response)
	}
}

func findB2BSupplierOnboardingItem(t *testing.T, items []b2bSupplierOnboardingItem, peerID string) b2bSupplierOnboardingItem {
	t.Helper()
	for _, item := range items {
		if item.PeerID == peerID {
			return item
		}
	}
	t.Fatalf("expected supplier onboarding item %q, got %#v", peerID, items)
	return b2bSupplierOnboardingItem{}
}

func hasB2BDisclosureProfile(profiles []b2bDisclosureProfile, profileID string) bool {
	for _, profile := range profiles {
		if profile.ProfileID == profileID {
			return true
		}
	}
	return false
}

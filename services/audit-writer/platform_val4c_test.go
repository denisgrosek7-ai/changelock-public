package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTrustHubGovernanceHandler(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-governance-4c",
		Organization:      "Governance Partner",
		Region:            "eu-west",
		Cluster:           "governance-eu-1",
		TrustDomain:       "governance.partner.example",
		Endpoint:          "https://governance-partner.example.invalid",
		PublicKeys:        []string{"pub-governance-1"},
		Capabilities:      []string{"sealed_handoff", "policy_sync"},
		PolicyRole:        federationPolicyRoleLeader,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe, incidentAudienceCustomerSafe},
	})
	executeStrictValidationRunForTest(t, fixture.handler, "/v1/validation/execute?tenant_id=acme&environment=prod&service=edge-gateway")

	req := httptest.NewRequest(http.MethodGet, "/v1/trust-hub/governance?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected trust hub governance 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response trustHubGovernanceResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode trust hub governance response: %v", err)
	}
	if response.SchemaVersion != trustHubGovernanceSchema || !response.NoNewTruthLayer {
		t.Fatalf("expected governance schema and no-new-truth-layer guard, got %#v", response)
	}
	if len(response.GovernanceRuleCatalog) < 4 || len(response.StandardsMappings) == 0 {
		t.Fatalf("expected governance catalog and standards mapping, got %#v", response)
	}
	runtimeRule := findTrustHubGovernanceRule(t, response.GovernanceRuleCatalog, "runtime_response_governance")
	partnerRule := findTrustHubGovernanceRule(t, response.GovernanceRuleCatalog, "partner_proof_governance")
	if runtimeRule.CurrentState == "" || len(runtimeRule.EvidenceRefs) == 0 {
		t.Fatalf("expected runtime governance rule with evidence refs, got %#v", runtimeRule)
	}
	if partnerRule.CurrentState == "" || len(partnerRule.MappedSurfaces) == 0 {
		t.Fatalf("expected partner governance rule with mapped surfaces, got %#v", partnerRule)
	}
}

func TestTrustHubAnalyticsHandler(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-stale-4c",
		Organization:      "Stale Partner",
		Region:            "us-east",
		Cluster:           "stale-us-1",
		TrustDomain:       "stale.partner.example",
		Endpoint:          "https://stale-partner.example.invalid",
		PublicKeys:        []string{"pub-stale-4c"},
		Capabilities:      []string{"sealed_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
		LastSeen:          mustParseTimeRFC3339(t, "2026-04-20T00:00:00Z"),
	})
	executeStrictValidationRunForTest(t, fixture.handler, "/v1/validation/execute?tenant_id=acme&environment=prod&service=edge-gateway")

	req := httptest.NewRequest(http.MethodGet, "/v1/trust-hub/analytics?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected trust hub analytics 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response trustHubAnalyticsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode trust hub analytics response: %v", err)
	}
	if response.SchemaVersion != trustHubAnalyticsSchema {
		t.Fatalf("expected analytics schema, got %#v", response)
	}
	if response.InternalPosture.Score <= 0 || response.InternalPosture.Grade == "" {
		t.Fatalf("expected internal posture score, got %#v", response.InternalPosture)
	}
	if len(response.PartnerPostures) == 0 || len(response.TrustHealthIndicators) < 4 {
		t.Fatalf("expected partner posture scores and trust indicators, got %#v", response)
	}
	if response.ShieldHealth.Band == "" || len(response.RiskTrends) == 0 {
		t.Fatalf("expected shield health and risk trends, got %#v", response)
	}
}

func TestTrustHubScopeClearanceHandler(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	executeStrictValidationRunForTest(t, fixture.handler, "/v1/validation/execute?tenant_id=acme&environment=prod&service=edge-gateway")
	sealed := sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)

	req := httptest.NewRequest(http.MethodGet, "/v1/trust-hub/clearance?tenant_id=acme&environment=prod&package_id="+sealed.PackageID, nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected scope clearance 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response trustHubClearanceResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode scope clearance response: %v", err)
	}
	if response.SchemaVersion != trustHubClearanceSchema || response.Subject.SubjectType != "scope" {
		t.Fatalf("expected scope clearance schema and subject, got %#v", response)
	}
	if response.CurrentState == "" || response.ClearanceLevel == "" || len(response.SupportingSignals) < 3 {
		t.Fatalf("expected bounded scope clearance contract, got %#v", response)
	}
	if findTrustHubClearanceSignal(t, response.SupportingSignals, "validation_certificate").CurrentState == "" {
		t.Fatalf("expected validation clearance signal, got %#v", response.SupportingSignals)
	}
	if findTrustHubClearanceSignal(t, response.SupportingSignals, "handoff_verification").CurrentState != handoffVerificationValid {
		t.Fatalf("expected valid handoff verification signal, got %#v", response.SupportingSignals)
	}
}

func TestTrustHubPartnerClearanceHandler(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	peer := registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-clearance-4c",
		Organization:      "Partner Clearance Co",
		Region:            "eu-central",
		Cluster:           "clearance-eu-1",
		TrustDomain:       "clearance.partner.example",
		Endpoint:          "https://clearance-partner.example.invalid",
		PublicKeys:        []string{"pub-clearance-1"},
		Capabilities:      []string{"sealed_handoff", "supplier_proof"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
	})
	sealed := sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)

	req := httptest.NewRequest(http.MethodGet, "/v1/trust-hub/clearance?tenant_id=acme&environment=prod&peer_id="+peer.PeerID+"&package_id="+sealed.PackageID, nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected partner clearance 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response trustHubClearanceResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode partner clearance response: %v", err)
	}
	if response.SchemaVersion != trustHubClearanceSchema || response.Subject.SubjectType != "partner" || response.Subject.SubjectRef != peer.PeerID {
		t.Fatalf("expected partner clearance subject, got %#v", response)
	}
	if response.CurrentState == "" || len(response.SupportingSignals) < 2 {
		t.Fatalf("expected partner clearance signals, got %#v", response)
	}
	if findTrustHubClearanceSignal(t, response.SupportingSignals, "partner_peer_freshness").CurrentState != federationPeerStatusActive {
		t.Fatalf("expected active peer freshness signal, got %#v", response.SupportingSignals)
	}
}

func TestTrustHubBoundariesHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/trust-hub/boundaries?tenant_id=acme", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected trust hub boundaries 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response trustHubBoundariesResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode trust hub boundaries response: %v", err)
	}
	if response.SchemaVersion != trustHubBoundariesSchema || !response.NoNewTruthLayer {
		t.Fatalf("expected trust hub boundaries schema and guardrail, got %#v", response)
	}
	if len(response.Authorizes) == 0 || len(response.RecommendOnly) == 0 || len(response.ExternalBoundaries) == 0 {
		t.Fatalf("expected bounded trust hub boundary groups, got %#v", response)
	}
}

func findTrustHubGovernanceRule(t *testing.T, items []trustHubGovernanceRule, ruleID string) trustHubGovernanceRule {
	t.Helper()
	for _, item := range items {
		if item.RuleID == ruleID {
			return item
		}
	}
	t.Fatalf("expected governance rule %q, got %#v", ruleID, items)
	return trustHubGovernanceRule{}
}

func findTrustHubClearanceSignal(t *testing.T, items []trustHubClearanceSignal, signalID string) trustHubClearanceSignal {
	t.Helper()
	for _, item := range items {
		if item.SignalID == signalID {
			return item
		}
	}
	t.Fatalf("expected clearance signal %q, got %#v", signalID, items)
	return trustHubClearanceSignal{}
}

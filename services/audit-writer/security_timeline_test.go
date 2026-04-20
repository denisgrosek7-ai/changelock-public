package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecurityTimelineAggregatesCommandCenterSignals(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")
	t.Setenv("CHANGELOCK_FEDERATION_SIGNING_SEED", "federation-seed")

	fixture := forensicsTestFixture(t)

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
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-2A-1"}`),
	)
	hardeningReq.Header.Set("Authorization", "Bearer operator-demo-token")
	hardeningReq.Header.Set("Content-Type", "application/json")
	hardeningRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(hardeningRec, hardeningReq)
	if hardeningRec.Code != http.StatusOK {
		t.Fatalf("expected hardening quarantine 200, got %d: %s", hardeningRec.Code, hardeningRec.Body.String())
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
		t.Fatalf("expected strict validation execute 200, got %d: %s", validationRec.Code, validationRec.Body.String())
	}

	sealed := sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)
	peer := registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-review",
		Organization:      "trusted-reviewer",
		Region:            "eu-central",
		TrustDomain:       "partner",
		PublicKeys:        []string{"peer-review-pub"},
		Capabilities:      []string{"sealed_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
	})
	bundle := downloadHandoffBundleForTest(t, fixture.handler, sealed.PackageID)
	verifyReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/federation/proof-verify",
		bytes.NewBufferString(`{"peer_id":"`+peer.PeerID+`","bundle_base64":"`+base64.StdEncoding.EncodeToString(bundle)+`","requested_scope":{"tenant_id":"acme","environment":"prod","audience":"auditor_safe"}}`),
	)
	verifyReq.Header.Set("Authorization", "Bearer operator-demo-token")
	verifyReq.Header.Set("Content-Type", "application/json")
	verifyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(verifyRec, verifyReq)
	if verifyRec.Code != http.StatusOK {
		t.Fatalf("expected federation proof verify 200, got %d: %s", verifyRec.Code, verifyRec.Body.String())
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/command-center/timeline?tenant_id=acme&environment=prod&limit=20", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected security timeline 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response securityTimelineResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode security timeline: %v", err)
	}
	if response.SchemaVersion != securityTimelineSchemaVersion {
		t.Fatalf("expected schema-versioned security timeline, got %#v", response)
	}
	if len(response.Entries) == 0 {
		t.Fatalf("expected timeline entries, got %#v", response)
	}
	if len(response.CountsBySource) == 0 || len(response.CountsBySeverity) == 0 {
		t.Fatalf("expected timeline source and severity counts, got %#v", response)
	}

	var sawRuntime, sawValidation, sawHandoff, sawFederation bool
	for _, entry := range response.Entries {
		if entry.SchemaVersion != securityTimelineEntrySchemaVersion {
			t.Fatalf("expected schema-versioned timeline entry, got %#v", entry)
		}
		if entry.Title == "" || entry.SubjectLabel == "" || entry.DrilldownTab == "" {
			t.Fatalf("expected titled, routed timeline entry, got %#v", entry)
		}
		if entry.DrilldownTargetKind == "" || entry.DrilldownTargetRef == "" {
			t.Fatalf("expected exact drilldown target on timeline entry, got %#v", entry)
		}
		if len(entry.PersonaHints) == 0 {
			t.Fatalf("expected persona hints on timeline entry, got %#v", entry)
		}
		switch entry.SourceSubsystem {
		case "runtime", "hardening":
			sawRuntime = true
		case "validation":
			sawValidation = true
		case "handoff":
			sawHandoff = true
		case "federation":
			sawFederation = true
		}
	}
	if !sawRuntime || !sawValidation || !sawHandoff || !sawFederation {
		t.Fatalf("expected unified timeline to include runtime, validation, handoff, and federation signals, got %#v", response.CountsBySource)
	}
}

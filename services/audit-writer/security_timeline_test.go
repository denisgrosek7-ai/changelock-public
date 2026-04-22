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

	vulnerabilityReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/vulnerability-relevance?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/edge-gateway",
	    "vulnerability_id":"CVE-2026-3000",
	    "image_digest":"sha256:edge",
	    "package_name":"openssl",
	    "severity":"critical",
	    "reachability":{"current_state":"observed_reachable","confidence_score":90,"evidence_refs":["runtime:callgraph"]},
	    "exploitability":{"epss":0.88,"external_exposure":true,"local_confidence":82}
	  }
	}`))
	vulnerabilityReq.Header.Set("Authorization", "Bearer operator-demo-token")
	vulnerabilityReq.Header.Set("Content-Type", "application/json")
	vulnerabilityRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(vulnerabilityRec, vulnerabilityReq)
	if vulnerabilityRec.Code != http.StatusCreated {
		t.Fatalf("expected vulnerability relevance 201, got %d: %s", vulnerabilityRec.Code, vulnerabilityRec.Body.String())
	}

	workflowReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/workflow/lifecycle?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "workflow_id":"wf-edge",
	    "artifact_type":"finding",
	    "subject_ref":"cluster-a/acme-prod/Deployment/edge-gateway",
	    "severity":"critical",
	    "requested_state":"resolved",
	    "validation_required":true,
	    "validation_state":"pending",
	    "owners":{"finding_owner":"team-edge","remediation_owner":"team-edge","approver":"security-admin"},
	    "evidence_refs":["event://deploy-gate/1"]
	  }
	}`))
	workflowReq.Header.Set("Authorization", "Bearer operator-demo-token")
	workflowReq.Header.Set("Content-Type", "application/json")
	workflowRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(workflowRec, workflowReq)
	if workflowRec.Code != http.StatusCreated {
		t.Fatalf("expected workflow 201, got %d: %s", workflowRec.Code, workflowRec.Body.String())
	}

	partnerReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/partner-trust/intake?tenant_id=acme&environment=prod&repo=github.com/acme/api&partner_id=vendor-edge", bytes.NewBufferString(`{
	  "input":{
	    "partner_id":"vendor-edge",
	    "organization":"Vendor Edge",
	    "trust_domain":"suppliers.acme",
	    "handoff_ref":"handoff-edge",
	    "verification_status":"verified",
	    "freshness_state":"fresh",
	    "policy_compatibility":"compatible",
	    "partner_visible_evidence":["sealed://proof/edge"],
	    "evidence_refs":["handoff://edge"]
	  }
	}`))
	partnerReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	partnerReq.Header.Set("Content-Type", "application/json")
	partnerRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(partnerRec, partnerReq)
	if partnerRec.Code != http.StatusCreated {
		t.Fatalf("expected partner intake 201, got %d: %s", partnerRec.Code, partnerRec.Body.String())
	}

	complianceReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/governance/compliance-mapping?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/edge-gateway",
	    "control_family":"soc2.cc7",
	    "control_id":"CC7.2",
	    "coverage_state":"partial",
	    "freshness_state":"fresh",
	    "evidence_refs":["event://deploy-gate/1"]
	  }
	}`))
	complianceReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	complianceReq.Header.Set("Content-Type", "application/json")
	complianceRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(complianceRec, complianceReq)
	if complianceRec.Code != http.StatusCreated {
		t.Fatalf("expected compliance mapping 201, got %d: %s", complianceRec.Code, complianceRec.Body.String())
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

	var sawRuntime, sawValidation, sawHandoff, sawFederation, sawIntelligence, sawWorkflow, sawPartner, sawGovernance bool
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
		case "intelligence":
			sawIntelligence = true
		case "workflow":
			sawWorkflow = true
		case "partner":
			sawPartner = true
		case "governance":
			sawGovernance = true
		}
	}
	for _, lifecycle := range []string{"runtime", "validation", "intelligence", "workflow", "partner", "governance"} {
		if response.CountsByLifecycle[lifecycle] == 0 {
			t.Fatalf("expected lifecycle count for %s, got %#v", lifecycle, response.CountsByLifecycle)
		}
	}
	if !sawRuntime || !sawValidation || !sawHandoff || !sawFederation || !sawIntelligence || !sawWorkflow || !sawPartner || !sawGovernance {
		t.Fatalf("expected unified timeline to include runtime, validation, handoff, federation, intelligence, workflow, partner, and governance signals, got %#v", response.CountsBySource)
	}

	filteredReq := httptest.NewRequest(http.MethodGet, "/v1/command-center/timeline?tenant_id=acme&environment=prod&limit=20&lifecycle_phase=intelligence", nil)
	filteredReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	filteredRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(filteredRec, filteredReq)
	if filteredRec.Code != http.StatusOK {
		t.Fatalf("expected filtered security timeline 200, got %d: %s", filteredRec.Code, filteredRec.Body.String())
	}
	var filtered securityTimelineResponse
	if err := json.NewDecoder(filteredRec.Body).Decode(&filtered); err != nil {
		t.Fatalf("decode filtered security timeline: %v", err)
	}
	for _, entry := range filtered.Entries {
		if entry.LifecyclePhase != "intelligence" {
			t.Fatalf("expected only intelligence timeline entries after filter, got %#v", filtered.Entries)
		}
	}
}

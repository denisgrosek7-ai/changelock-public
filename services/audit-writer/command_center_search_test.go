package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommandCenterSearchReturnsExactFocusTargets(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")
	t.Setenv("CHANGELOCK_FEDERATION_SIGNING_SEED", "federation-seed")

	fixture := forensicsTestFixture(t)

	incidentReq := httptest.NewRequest(http.MethodGet, "/v1/incidents?tenant_id=acme&environment=prod&limit=20", nil)
	incidentReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	incidentRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(incidentRec, incidentReq)
	if incidentRec.Code != http.StatusOK {
		t.Fatalf("expected incidents 200, got %d: %s", incidentRec.Code, incidentRec.Body.String())
	}
	var incidentList incidentsResponse
	if err := json.NewDecoder(incidentRec.Body).Decode(&incidentList); err != nil {
		t.Fatalf("decode incidents: %v", err)
	}
	if len(incidentList.Incidents) == 0 {
		t.Fatal("expected fixture incidents")
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
	if len(findings.Items) == 0 {
		t.Fatal("expected runtime findings in fixture")
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
	var validationRun validationExecutionRun
	if err := json.NewDecoder(validationRec.Body).Decode(&validationRun); err != nil {
		t.Fatalf("decode validation run: %v", err)
	}

	sealed := sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)
	peer := registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-command-center",
		Organization:      "trusted-reviewer",
		Region:            "eu-central",
		TrustDomain:       "partner",
		PublicKeys:        []string{"peer-command-center-pub"},
		Capabilities:      []string{"sealed_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
	})

	vulnerabilityReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/vulnerability-relevance?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/edge-gateway",
	    "vulnerability_id":"CVE-2026-4000",
	    "image_digest":"sha256:edge",
	    "package_name":"openssl",
	    "severity":"critical",
	    "reachability":{"current_state":"observed_reachable","confidence_score":91,"evidence_refs":["runtime:callgraph"]},
	    "exploitability":{"epss":0.9,"external_exposure":true,"local_confidence":84}
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
	    "workflow_id":"wf-search-edge",
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

	partnerReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/partner-trust/intake?tenant_id=acme&environment=prod&repo=github.com/acme/api&partner_id=vendor-search", bytes.NewBufferString(`{
	  "input":{
	    "partner_id":"vendor-search",
	    "organization":"Vendor Search",
	    "trust_domain":"suppliers.acme",
	    "handoff_ref":"handoff-search",
	    "verification_status":"verified",
	    "freshness_state":"fresh",
	    "policy_compatibility":"compatible",
	    "partner_visible_evidence":["sealed://proof/search"],
	    "evidence_refs":["handoff://search"]
	  }
	}`))
	partnerReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	partnerReq.Header.Set("Content-Type", "application/json")
	partnerRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(partnerRec, partnerReq)
	if partnerRec.Code != http.StatusCreated {
		t.Fatalf("expected partner intake 201, got %d: %s", partnerRec.Code, partnerRec.Body.String())
	}

	checkSearch := func(query string, wantKind string, extra func(result commandCenterSearchResult)) {
		req := httptest.NewRequest(http.MethodGet, "/v1/command-center/search?tenant_id=acme&environment=prod&limit=10&q="+query, nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		fixture.handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected command-center search 200, got %d: %s", rec.Code, rec.Body.String())
		}
		var response commandCenterSearchResponse
		if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
			t.Fatalf("decode command-center search: %v", err)
		}
		if response.SchemaVersion != commandSearchResponseSchemaVersion || response.Query == "" {
			t.Fatalf("expected schema-versioned command-center search response, got %#v", response)
		}
		for _, result := range response.Results {
			if result.Target.Kind == wantKind {
				if result.SchemaVersion != commandSearchResultSchemaVersion {
					t.Fatalf("expected schema-versioned search result, got %#v", result)
				}
				if result.Target.Ref == "" || result.Target.Tab == "" {
					t.Fatalf("expected exact focus target on search result, got %#v", result)
				}
				if extra != nil {
					extra(result)
				}
				return
			}
		}
		t.Fatalf("expected search query %q to return target kind %s, got %#v", query, wantKind, response.Results)
	}

	checkSearch(incidentList.Incidents[0].ID, "incident", func(result commandCenterSearchResult) {
		if result.Target.Tab != "events" || result.IncidentRef != incidentList.Incidents[0].ID {
			t.Fatalf("expected incident search result to open exact incident focus, got %#v", result)
		}
	})
	checkSearch(findings.Items[0].FindingID, "runtime_finding", func(result commandCenterSearchResult) {
		if result.Target.Tab != "runtime" || result.Target.SecondaryRef == "" {
			t.Fatalf("expected runtime finding search result to carry exact runtime focus, got %#v", result)
		}
	})
	checkSearch(validationRun.RunID, "validation_run", func(result commandCenterSearchResult) {
		if result.Target.Tab != "validation" {
			t.Fatalf("expected validation run result to route to validation, got %#v", result)
		}
	})
	checkSearch(peer.PeerID, "federation_peer", func(result commandCenterSearchResult) {
		if result.Target.Tab != "federation" {
			t.Fatalf("expected federation peer result to route to federation, got %#v", result)
		}
	})
	checkSearch(sealed.PackageID, "handoff_package", func(result commandCenterSearchResult) {
		if result.Target.ResourceURI == "" || result.Target.Ref != sealed.PackageID {
			t.Fatalf("expected handoff package result to preserve exact package lookup metadata, got %#v", result)
		}
	})
	checkSearch("CVE-2026-4000", "vulnerability_relevance", func(result commandCenterSearchResult) {
		if result.Target.Tab != "vulnerabilities" || result.Target.SecondaryRef != "CVE-2026-4000" {
			t.Fatalf("expected vulnerability relevance result to route to vulnerabilities, got %#v", result)
		}
	})
	checkSearch("wf-search-edge", "workflow_record", func(result commandCenterSearchResult) {
		if result.Target.Tab != "exceptions" || result.Target.Ref != "wf-search-edge" {
			t.Fatalf("expected workflow record result to route to workflow view, got %#v", result)
		}
	})
	checkSearch("vendor-search", "partner_trust", func(result commandCenterSearchResult) {
		if result.Target.Tab != "federation" || result.Target.Ref != "vendor-search" {
			t.Fatalf("expected partner trust result to route to federation, got %#v", result)
		}
	})
}

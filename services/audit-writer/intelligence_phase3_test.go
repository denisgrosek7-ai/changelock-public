package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
)

func TestPhase3IntelligenceFlowAndProofs(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	vulnerabilityReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/vulnerability-relevance?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "vulnerability_id":"CVE-2026-0001",
	    "image_digest":"sha256:abc",
	    "package_name":"openssl",
	    "severity":"critical",
	    "reachability":{
	      "current_state":"observed_reachable",
	      "observed_calls":12,
	      "loaded_artifacts":1,
	      "exposure_class":"internet",
	      "confidence_score":92,
	      "evidence_refs":["runtime:callgraph","runtime:exposure-map"]
	    },
	    "exploitability":{
	      "epss":0.93,
	      "external_exposure":true,
	      "substrate_state":"runtime_truth_bound",
	      "local_confidence":88
	    },
	    "additional_evidence":["report:/v1/runtime/phase2/proofs"]
	  }
	}`))
	vulnerabilityReq.Header.Set("Authorization", "Bearer operator-demo-token")
	vulnerabilityReq.Header.Set("Content-Type", "application/json")
	vulnerabilityRec := httptest.NewRecorder()
	handler.ServeHTTP(vulnerabilityRec, vulnerabilityReq)
	if vulnerabilityRec.Code != http.StatusCreated {
		t.Fatalf("expected vulnerability relevance 201, got %d: %s", vulnerabilityRec.Code, vulnerabilityRec.Body.String())
	}
	var vulnerabilityResponse phase3VulnerabilityEvaluateResponse
	if err := json.NewDecoder(vulnerabilityRec.Body).Decode(&vulnerabilityResponse); err != nil {
		t.Fatalf("decode vulnerability response: %v", err)
	}
	if vulnerabilityResponse.Verdict.VEXCandidate == nil {
		t.Fatalf("expected bounded VEX candidate, got %#v", vulnerabilityResponse.Verdict)
	}

	supplyChainReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/supply-chain/patterns?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "package_name":"openssl",
	    "package_version":"3.0.0-mal",
	    "publisher":"mallory",
	    "previous_publisher":"openssl-foundation",
	    "publish_delta_hours":2,
	    "provenance_consistent":false,
	    "signing_consistent":false,
	    "baseline":{
	      "workload_class":"api",
	      "language_ecosystem":"go",
	      "expected_behaviors":["http_server","db_client"],
	      "expected_publish_cadence_hours":72
	    },
	    "runtime_behaviors":["ssh_spawn"],
	    "federated_signals":[
	      {"source_peer_id":"cluster-b","signal_type":"maintainer_drift","suspicion_level":"high","source_confidence":92,"reason_codes":["maintainer_drift_observed"]},
	      {"source_peer_id":"cluster-c","signal_type":"publish_anomaly","suspicion_level":"high","source_confidence":90,"reason_codes":["sudden_publish_cadence_anomaly"]}
	    ],
	    "evidence_refs":["report:/v1/sbom/components"]
	  }
	}`))
	supplyChainReq.Header.Set("Authorization", "Bearer operator-demo-token")
	supplyChainReq.Header.Set("Content-Type", "application/json")
	supplyChainRec := httptest.NewRecorder()
	handler.ServeHTTP(supplyChainRec, supplyChainReq)
	if supplyChainRec.Code != http.StatusCreated {
		t.Fatalf("expected supply-chain pattern 201, got %d: %s", supplyChainRec.Code, supplyChainRec.Body.String())
	}

	strategicReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/strategic/simulate?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "subject_ref":"cluster-a/acme-prod/Deployment/api",
	  "candidate_action":"prioritize_patch_and_validate",
	  "delay_days":7,
	  "effort_band":"moderate",
	  "blast_radius_score":75
	}`))
	strategicReq.Header.Set("Authorization", "Bearer operator-demo-token")
	strategicReq.Header.Set("Content-Type", "application/json")
	strategicRec := httptest.NewRecorder()
	handler.ServeHTTP(strategicRec, strategicReq)
	if strategicRec.Code != http.StatusOK {
		t.Fatalf("expected strategic simulate 200, got %d: %s", strategicRec.Code, strategicRec.Body.String())
	}

	queryReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/strategic/query?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "query":"Which issue should I prioritize for the api workload?",
	  "subject_ref":"cluster-a/acme-prod/Deployment/api"
	}`))
	queryReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	queryReq.Header.Set("Content-Type", "application/json")
	queryRec := httptest.NewRecorder()
	handler.ServeHTTP(queryRec, queryReq)
	if queryRec.Code != http.StatusOK {
		t.Fatalf("expected grounded query 200, got %d: %s", queryRec.Code, queryRec.Body.String())
	}
	var queryResponse phase3StrategicQueryResponse
	if err := json.NewDecoder(queryRec.Body).Decode(&queryResponse); err != nil {
		t.Fatalf("decode query response: %v", err)
	}
	if !queryResponse.Response.AdvisoryOnly || queryResponse.Response.RetrievalMode != "bounded_evidence_retrieval" {
		t.Fatalf("expected advisory grounded query response, got %#v", queryResponse.Response)
	}
	if queryResponse.Response.Scope.SubjectRef != "cluster-a|acme-prod|Deployment|api" || queryResponse.Response.Scope.Repo != "github.com/acme/api" {
		t.Fatalf("expected bounded query scope metadata, got %#v", queryResponse.Response.Scope)
	}

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/intelligence/phase3/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api", nil)
	proofsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	proofsRec := httptest.NewRecorder()
	handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}
	var proofs phase3ProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if proofs.CurrentState != phase3ProofStateActive {
		t.Fatalf("expected active phase3 proofs, got %#v", proofs)
	}
}

func TestPhase3ProofsRequireAllEvidenceTypes(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	vulnerabilityReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/vulnerability-relevance?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "vulnerability_id":"CVE-2026-0009",
	    "package_name":"libxml2",
	    "severity":"medium",
	    "reachability":{
	      "current_state":"not_observed_reachable",
	      "confidence_score":82,
	      "evidence_refs":["runtime:library-load"]
	    },
	    "exploitability":{
	      "epss":0.14,
	      "local_confidence":63
	    }
	  }
	}`))
	vulnerabilityReq.Header.Set("Authorization", "Bearer operator-demo-token")
	vulnerabilityReq.Header.Set("Content-Type", "application/json")
	vulnerabilityRec := httptest.NewRecorder()
	handler.ServeHTTP(vulnerabilityRec, vulnerabilityReq)
	if vulnerabilityRec.Code != http.StatusCreated {
		t.Fatalf("expected vulnerability relevance 201, got %d: %s", vulnerabilityRec.Code, vulnerabilityRec.Body.String())
	}

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/intelligence/phase3/proofs?tenant_id=acme&environment=prod", nil)
	proofsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	proofsRec := httptest.NewRecorder()
	handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}
	var proofs phase3ProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if proofs.CurrentState != phase3ProofStateIncomplete {
		t.Fatalf("expected incomplete phase3 proofs without supply-chain/strategic/query evidence, got %#v", proofs)
	}
}

func TestPhase3ProofsIgnoreCrossScopeQueryArtifacts(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	vulnerabilityReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/vulnerability-relevance?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "vulnerability_id":"CVE-2026-0100",
	    "image_digest":"sha256:api",
	    "package_name":"openssl",
	    "severity":"critical",
	    "reachability":{"current_state":"observed_reachable","confidence_score":90,"evidence_refs":["runtime:callgraph"]},
	    "exploitability":{"epss":0.91,"external_exposure":true,"local_confidence":88}
	  }
	}`))
	vulnerabilityReq.Header.Set("Authorization", "Bearer operator-demo-token")
	vulnerabilityReq.Header.Set("Content-Type", "application/json")
	vulnerabilityRec := httptest.NewRecorder()
	handler.ServeHTTP(vulnerabilityRec, vulnerabilityReq)
	if vulnerabilityRec.Code != http.StatusCreated {
		t.Fatalf("expected vulnerability relevance 201, got %d: %s", vulnerabilityRec.Code, vulnerabilityRec.Body.String())
	}

	supplyChainReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/supply-chain/patterns?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "package_name":"openssl",
	    "package_version":"3.0.0-mal",
	    "publisher":"mallory",
	    "previous_publisher":"openssl-foundation",
	    "provenance_consistent":false,
	    "signing_consistent":false,
	    "baseline":{"workload_class":"api","language_ecosystem":"go","expected_behaviors":["http_server"]},
	    "runtime_behaviors":["ssh_spawn"],
	    "evidence_refs":["report:/v1/sbom/components"]
	  }
	}`))
	supplyChainReq.Header.Set("Authorization", "Bearer operator-demo-token")
	supplyChainReq.Header.Set("Content-Type", "application/json")
	supplyChainRec := httptest.NewRecorder()
	handler.ServeHTTP(supplyChainRec, supplyChainReq)
	if supplyChainRec.Code != http.StatusCreated {
		t.Fatalf("expected supply-chain pattern 201, got %d: %s", supplyChainRec.Code, supplyChainRec.Body.String())
	}

	strategicReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/strategic/simulate?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "subject_ref":"cluster-a/acme-prod/Deployment/api",
	  "candidate_action":"prioritize_patch_and_validate",
	  "delay_days":7,
	  "effort_band":"moderate",
	  "blast_radius_score":75
	}`))
	strategicReq.Header.Set("Authorization", "Bearer operator-demo-token")
	strategicReq.Header.Set("Content-Type", "application/json")
	strategicRec := httptest.NewRecorder()
	handler.ServeHTTP(strategicRec, strategicReq)
	if strategicRec.Code != http.StatusOK {
		t.Fatalf("expected strategic simulate 200, got %d: %s", strategicRec.Code, strategicRec.Body.String())
	}

	crossScopeQueryReq := httptest.NewRequest(http.MethodPost, "/v1/intelligence/strategic/query?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "query":"What should I prioritize for the worker workload?",
	  "subject_ref":"cluster-a/acme-prod/Deployment/worker"
	}`))
	crossScopeQueryReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	crossScopeQueryReq.Header.Set("Content-Type", "application/json")
	crossScopeQueryRec := httptest.NewRecorder()
	handler.ServeHTTP(crossScopeQueryRec, crossScopeQueryReq)
	if crossScopeQueryRec.Code != http.StatusOK {
		t.Fatalf("expected grounded query 200, got %d: %s", crossScopeQueryRec.Code, crossScopeQueryRec.Body.String())
	}

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/intelligence/phase3/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	proofsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	proofsRec := httptest.NewRecorder()
	handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}
	var proofs phase3ProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if proofs.CurrentState != phase3ProofStateIncomplete {
		t.Fatalf("expected incomplete proofs when only cross-scope query exists, got %#v", proofs)
	}
}

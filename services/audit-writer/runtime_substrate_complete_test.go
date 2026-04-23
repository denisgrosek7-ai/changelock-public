package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

func TestRuntimeSubstratePoint1CompleteHandlerFailsClosedWithoutActiveWaves(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/complete?tenant_id=acme&environment=prod&limit=20", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected point 1 complete 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response runtimeSubstratePoint1CompleteResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode point 1 complete: %v", err)
	}
	if response.CurrentState == runtimesubstrate.RuntimeSubstratePoint1StateActive {
		t.Fatalf("expected non-active point 1 state without seeded runtime chain, got %#v", response)
	}
	if len(response.DeferredScope) == 0 {
		t.Fatalf("expected explicit deferred scope, got %#v", response)
	}
}

func TestRuntimeSubstratePoint1CompleteHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)
	seedRuntimeSubstratePoint1Chain(t, fixture)

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/complete?tenant_id=acme&environment=prod&limit=50", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected point 1 complete 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response runtimeSubstratePoint1CompleteResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode point 1 complete: %v", err)
	}
	if response.CurrentState != runtimesubstrate.RuntimeSubstratePoint1StateActive {
		t.Fatalf("expected active point 1 closure, got %#v", response)
	}
	if response.Point1State != runtimesubstrate.RuntimeSubstratePoint1StateActive {
		t.Fatalf("expected active point_1_state, got %#v", response)
	}
	if response.ValAState != runtimesubstrate.RuntimeSubstrateValAStateActive ||
		response.ValBState != runtimesubstrate.RuntimeSubstrateValBStateActive ||
		response.ValCState != runtimesubstrate.RuntimeSubstrateValCStateActive ||
		response.ValDState != runtimesubstrate.RuntimeSubstrateValDStateActive ||
		response.ValEState != runtimesubstrate.RuntimeSubstrateValEStateActive {
		t.Fatalf("expected all val states active, got %#v", response)
	}
	if !hasRuntimeSubstrateString(response.SurfaceRefs, "/v1/runtime/substrate-depth/complete") {
		t.Fatalf("expected complete route ref, got %#v", response.SurfaceRefs)
	}
	if !hasRuntimeSubstrateString(response.SurfaceRefs, "/v1/runtime/substrate-depth/vale/proofs") {
		t.Fatalf("expected val e proofs route ref, got %#v", response.SurfaceRefs)
	}
	if !hasRuntimeSubstrateString(response.DocumentationRefs, "docs/runtime-substrate-depth-complete.md") {
		t.Fatalf("expected closure documentation ref, got %#v", response.DocumentationRefs)
	}
	if len(response.DeferredScope) == 0 || len(response.IntegrationSummary) == 0 || len(response.EvidenceRefs) == 0 {
		t.Fatalf("expected deferred scope, integration summary, and evidence refs, got %#v", response)
	}
}

func seedRuntimeSubstratePoint1Chain(t *testing.T, fixture forensicsFixtureData) {
	t.Helper()

	for _, event := range []runtimesubstrate.RuntimeSubstrateObservedEvent{
		runtimeSubstrateValBExpectedExecEvent(),
		runtimeSubstrateValBLowRiskExecEvent(),
		runtimeSubstrateValBHardMismatchExecEvent(),
		runtimeSubstrateProcessStaleEvent(),
		runtimeSubstrateFilePartialEvent(),
		runtimeSubstrateNetworkUnsupportedEvent(),
	} {
		postRuntimeSubstrateObservation(t, fixture.handler, event)
	}

	seedRuntimeSubstrateValBArtifactEvidence(t, fixture.store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "api", "sha256:111", "https://github.com/example/api/.github/workflows/release.yml@refs/heads/main", "sha256:111", "https://github.com/example/api", "https://slsa.dev/provenance/v1")
	seedRuntimeSubstrateValBArtifactEvidence(t, fixture.store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "worker", "", "https://github.com/example/worker/.github/workflows/release.yml@refs/heads/main", "", "https://github.com/example/worker", "")
	seedRuntimeSubstrateValBArtifactEvidence(t, fixture.store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "rogue", "sha256:trusted", "https://github.com/example/rogue/.github/workflows/release.yml@refs/heads/main", "sha256:trusted", "https://github.com/example/rogue", "https://slsa.dev/provenance/v1")

	seedRuntimeSubstrateValBAttestation(t, fixture.store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "api", runtimeSnapshotVerifiedAttestation("api", "sha256:111"))
	seedRuntimeSubstrateValBAttestation(t, fixture.store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "worker", runtimeSnapshotDegradedAttestation("worker", "sha256:worker"))
	seedRuntimeSubstrateValBAttestation(t, fixture.store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "rogue", runtimeSnapshotMismatchAttestation("rogue", "sha256:rogue"))

	findingsReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/findings?tenant_id=acme&environment=prod&limit=20", nil)
	findingsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	findingsRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(findingsRec, findingsReq)
	if findingsRec.Code != http.StatusOK {
		t.Fatalf("expected runtime findings 200, got %d: %s", findingsRec.Code, findingsRec.Body.String())
	}
	var findings runtimeFindingsResponse
	if err := json.NewDecoder(findingsRec.Body).Decode(&findings); err != nil {
		t.Fatalf("decode findings: %v", err)
	}
	binaryFinding := findRuntimeFinding(t, findings.Items, runtimeFindingUnknownBinaryExec, "edge-gateway")

	postValCDecisionRequest(t, fixture.handler, "/v1/runtime/forensic-snapshot?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`"}`)
	postValCDecisionRequest(t, fixture.handler, "/v1/runtime/quarantine?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-COMPLETE-1"}`)
	postValCDecisionRequest(t, fixture.handler, "/v1/runtime/restart-trusted?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-COMPLETE-2"}`)
	postValCDecisionRequest(t, fixture.handler, "/v1/hardening/quarantine?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-COMPLETE-3"}`)
	seedRuntimeSubstrateValCPreventExecution(t, fixture.store)
}

func hasRuntimeSubstrateString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

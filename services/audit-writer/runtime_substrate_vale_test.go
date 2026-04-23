package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

func TestRuntimeSubstrateValELatencyPackHandler(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vale/latency-pack?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale latency pack 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response runtimeSubstrateValELatencyPackResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode vale latency pack: %v", err)
	}
	if response.CurrentState != runtimesubstrate.RuntimeSubstrateValELatencyPackStateActive {
		t.Fatalf("expected active vale latency pack, got %#v", response)
	}
	if len(response.Items) != 5 {
		t.Fatalf("expected 5 latency pack items, got %#v", response.Items)
	}
	if !hasValELatencyMeasurement(response.Items, runtimesubstrate.RuntimeExecutionClassStandardNode) {
		t.Fatalf("expected measured standard-node latency record, got %#v", response.Items)
	}
}

func TestRuntimeSubstrateValEHandlersAndProofs(t *testing.T) {
	fixture := forensicsTestFixture(t)

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
	postValCDecisionRequest(t, fixture.handler, "/v1/runtime/quarantine?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-VALE-1"}`)
	postValCDecisionRequest(t, fixture.handler, "/v1/runtime/restart-trusted?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-VALE-2"}`)
	postValCDecisionRequest(t, fixture.handler, "/v1/hardening/quarantine?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-VALE-3"}`)
	seedRuntimeSubstrateValCPreventExecution(t, fixture.store)

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vale/false-positive-budget?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale false-positive budget 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var falsePositive runtimeSubstrateValEFalsePositiveBudgetResponse
	if err := json.NewDecoder(rec.Body).Decode(&falsePositive); err != nil {
		t.Fatalf("decode vale false-positive budget: %v", err)
	}
	if falsePositive.CurrentState != runtimesubstrate.RuntimeSubstrateValEFalsePositiveBudgetStateActive {
		t.Fatalf("expected active false-positive budget, got %#v", falsePositive)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vale/replayable-benchmark-pack?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale replayable benchmark pack 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var replayable runtimeSubstrateValEReplayableBenchmarkPackResponse
	if err := json.NewDecoder(rec.Body).Decode(&replayable); err != nil {
		t.Fatalf("decode vale replayable benchmark pack: %v", err)
	}
	if replayable.CurrentState != runtimesubstrate.RuntimeSubstrateValEReplayableBenchmarkPackStateActive {
		t.Fatalf("expected active replayable benchmark pack, got %#v", replayable)
	}
	if len(replayable.Items) != 3 {
		t.Fatalf("expected 3 replayable benchmark packs, got %#v", replayable.Items)
	}
	if !hasValEReplayableProfile(replayable.Items, "stress") {
		t.Fatalf("expected stress replayable profile, got %#v", replayable.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vale/performance-gate?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale performance gate 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var performanceGate runtimeSubstrateValEPerformanceGateResponse
	if err := json.NewDecoder(rec.Body).Decode(&performanceGate); err != nil {
		t.Fatalf("decode vale performance gate: %v", err)
	}
	if performanceGate.CurrentState != runtimesubstrate.RuntimeSubstrateValEPerformanceGateStateActive {
		t.Fatalf("expected active performance gate, got %#v", performanceGate)
	}
	if len(performanceGate.Items) != 3 || len(performanceGate.BenchmarkEvaluations) != 3 {
		t.Fatalf("expected 3 performance gate items and evaluations, got %#v %#v", performanceGate.Items, performanceGate.BenchmarkEvaluations)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vale/proofs?tenant_id=acme&environment=prod&limit=50", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vale proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var proofs runtimeSubstrateValEProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode vale proofs: %v", err)
	}
	if proofs.CurrentState != runtimesubstrate.RuntimeSubstrateValEStateActive {
		t.Fatalf("expected active vale proofs, got %#v", proofs)
	}
	if proofs.ValDState != runtimesubstrate.RuntimeSubstrateValDStateActive {
		t.Fatalf("expected active vald dependency, got %#v", proofs)
	}
	if proofs.LatencyPackState != runtimesubstrate.RuntimeSubstrateValELatencyPackStateActive {
		t.Fatalf("expected active latency pack state, got %#v", proofs)
	}
	if proofs.FalsePositiveBudgetState != runtimesubstrate.RuntimeSubstrateValEFalsePositiveBudgetStateActive {
		t.Fatalf("expected active false-positive state, got %#v", proofs)
	}
	if proofs.ReplayableBenchmarkState != runtimesubstrate.RuntimeSubstrateValEReplayableBenchmarkPackStateActive {
		t.Fatalf("expected active replayable benchmark state, got %#v", proofs)
	}
	if proofs.PerformanceGateState != runtimesubstrate.RuntimeSubstrateValEPerformanceGateStateActive {
		t.Fatalf("expected active performance gate state, got %#v", proofs)
	}
	if len(proofs.RemainingDeferredScope) != 0 {
		t.Fatalf("expected no remaining deferred scope after val e, got %#v", proofs.RemainingDeferredScope)
	}
}

func hasValELatencyMeasurement(items []runtimesubstrate.RuntimeSubstrateExecutionClassLatencyPackItem, classID string) bool {
	for _, item := range items {
		if item.ExecutionClass != classID {
			continue
		}
		return item.MeasurementSource != "" &&
			!item.MeasuredAt.IsZero() &&
			item.CaptureP99Micros > 0 &&
			item.CorrelationP99Micros > 0 &&
			item.EnforcementDecisionP99Micros > 0
	}
	return false
}

func hasValEReplayableProfile(items []runtimesubstrate.RuntimeSubstrateReplayableBenchmarkPackItem, profileID string) bool {
	for _, item := range items {
		if item.ProfileID == profileID && item.Replayable && len(item.CommandHints) > 0 {
			return true
		}
	}
	return false
}

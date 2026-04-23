package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

func TestRuntimeSubstrateValDExecutionClassMatrixHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vald/execution-class-matrix?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald execution class matrix 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response runtimeSubstrateValDExecutionClassMatrixResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode vald execution class matrix: %v", err)
	}
	if response.CurrentState != runtimesubstrate.RuntimeSubstrateValDExecutionClassMatrixStateActive {
		t.Fatalf("expected active execution class matrix, got %#v", response)
	}
	if len(response.Items) != 5 {
		t.Fatalf("expected 5 execution classes, got %#v", response.Items)
	}
	if !hasValDExecutionClass(response.Items, runtimesubstrate.RuntimeExecutionClassOfflineAirgappedNode) {
		t.Fatalf("expected offline airgapped execution class, got %#v", response.Items)
	}
}

func TestRuntimeSubstrateValDHandlersAndProofs(t *testing.T) {
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
	postValCDecisionRequest(t, fixture.handler, "/v1/runtime/quarantine?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-VALD-1"}`)
	postValCDecisionRequest(t, fixture.handler, "/v1/runtime/restart-trusted?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-VALD-2"}`)
	postValCDecisionRequest(t, fixture.handler, "/v1/hardening/quarantine?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-VALD-3"}`)
	seedRuntimeSubstrateValCPreventExecution(t, fixture.store)

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vald/signal-coverage?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald signal coverage 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var signalCoverage runtimeSubstrateValDSignalCoverageResponse
	if err := json.NewDecoder(rec.Body).Decode(&signalCoverage); err != nil {
		t.Fatalf("decode vald signal coverage: %v", err)
	}
	if signalCoverage.CurrentState != runtimesubstrate.RuntimeSubstrateValDSignalCoverageStateActive {
		t.Fatalf("expected active signal coverage, got %#v", signalCoverage)
	}
	if !hasValDUnsupportedFamily(signalCoverage.Items, runtimesubstrate.RuntimeExecutionClassOfflineAirgappedNode, runtimesubstrate.RuntimeSubstrateEventFamilyNetworkActivity) {
		t.Fatalf("expected offline airgapped network family to stay unsupported, got %#v", signalCoverage.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vald/enforcement-availability?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald enforcement availability 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var enforcement runtimeSubstrateValDEnforcementAvailabilityResponse
	if err := json.NewDecoder(rec.Body).Decode(&enforcement); err != nil {
		t.Fatalf("decode vald enforcement availability: %v", err)
	}
	if enforcement.CurrentState != runtimesubstrate.RuntimeSubstrateValDEnforcementAvailabilityStateActive {
		t.Fatalf("expected active enforcement availability, got %#v", enforcement)
	}
	if !hasValDUnsupportedAction(enforcement.Items, runtimesubstrate.RuntimeExecutionClassOfflineAirgappedNode, "runtime.apply_network_isolation") {
		t.Fatalf("expected offline airgapped runtime.apply_network_isolation to stay unsupported, got %#v", enforcement.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vald/overhead-visibility?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald overhead visibility 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var overhead runtimeSubstrateValDOverheadVisibilityResponse
	if err := json.NewDecoder(rec.Body).Decode(&overhead); err != nil {
		t.Fatalf("decode vald overhead visibility: %v", err)
	}
	if overhead.CurrentState != runtimesubstrate.RuntimeSubstrateValDOverheadVisibilityStateActive {
		t.Fatalf("expected active overhead visibility, got %#v", overhead)
	}
	if !hasValDMeasurementStatus(overhead.Items, runtimesubstrate.RuntimeExecutionClassStandardNode, "class_specific_measurement_verified") {
		t.Fatalf("expected standard node measurement status, got %#v", overhead.Items)
	}
	if !hasValDMeasuredRecord(overhead.Items, runtimesubstrate.RuntimeExecutionClassOfflineAirgappedNode) {
		t.Fatalf("expected measured overhead record for offline airgapped node, got %#v", overhead.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vald/proofs?tenant_id=acme&environment=prod&limit=50", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vald proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs runtimeSubstrateValDProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode vald proofs: %v", err)
	}
	if proofs.CurrentState != runtimesubstrate.RuntimeSubstrateValDStateActive {
		t.Fatalf("expected active vald proofs, got %#v", proofs)
	}
	if proofs.ValCState != runtimesubstrate.RuntimeSubstrateValCStateActive {
		t.Fatalf("expected active valc dependency, got %#v", proofs)
	}
	if proofs.SignalCoverageState != runtimesubstrate.RuntimeSubstrateValDSignalCoverageStateActive {
		t.Fatalf("expected active signal coverage state, got %#v", proofs)
	}
	if proofs.EnforcementAvailabilityState != runtimesubstrate.RuntimeSubstrateValDEnforcementAvailabilityStateActive {
		t.Fatalf("expected active enforcement availability state, got %#v", proofs)
	}
	if proofs.OverheadVisibilityState != runtimesubstrate.RuntimeSubstrateValDOverheadVisibilityStateActive {
		t.Fatalf("expected active overhead visibility state, got %#v", proofs)
	}
	if !hasValDExecutionClass(proofs.ExecutionClassMatrix, runtimesubstrate.RuntimeExecutionClassConfidentialCapableNode) {
		t.Fatalf("expected confidential capable class in proofs, got %#v", proofs.ExecutionClassMatrix)
	}
	if !hasValDMeasuredRecord(proofs.OverheadVisibility, runtimesubstrate.RuntimeExecutionClassConfidentialCapableNode) {
		t.Fatalf("expected measured overhead evidence in proofs, got %#v", proofs.OverheadVisibility)
	}
}

func hasValDExecutionClass(items []runtimesubstrate.RuntimeSubstrateExecutionClassMatrixItem, classID string) bool {
	for _, item := range items {
		if item.ExecutionClass == classID {
			return true
		}
	}
	return false
}

func hasValDUnsupportedFamily(items []runtimesubstrate.RuntimeSubstrateExecutionClassSignalCoverageItem, classID, family string) bool {
	for _, item := range items {
		if item.ExecutionClass == classID && containsString(item.UnsupportedFamilies, family) {
			return true
		}
	}
	return false
}

func hasValDUnsupportedAction(items []runtimesubstrate.RuntimeSubstrateExecutionClassEnforcementAvailabilityItem, classID, actionID string) bool {
	for _, item := range items {
		if item.ExecutionClass == classID && containsString(item.UnsupportedActions, actionID) {
			return true
		}
	}
	return false
}

func hasValDMeasurementStatus(items []runtimesubstrate.RuntimeSubstrateExecutionClassOverheadVisibilityItem, classID, status string) bool {
	for _, item := range items {
		if item.ExecutionClass == classID && item.MeasurementStatus == status {
			return true
		}
	}
	return false
}

func hasValDMeasuredRecord(items []runtimesubstrate.RuntimeSubstrateExecutionClassOverheadVisibilityItem, classID string) bool {
	for _, item := range items {
		if item.ExecutionClass != classID {
			continue
		}
		return item.MeasurementBasis != "" &&
			!item.MeasuredAt.IsZero() &&
			item.MeasurementSource != "" &&
			len(item.EvidenceRefs) > 0 &&
			(item.ObservedCPUOverheadMillicores > 0 ||
				item.ObservedMemoryOverheadMiB > 0 ||
				item.ObservedCaptureLatencyMicros > 0 ||
				item.ObservedCorrelationLatencyMicros > 0)
	}
	return false
}

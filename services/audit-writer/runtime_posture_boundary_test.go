package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRuntimePostureLinkageContract(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/posture-linkage?tenant_id=acme&environment=prod&limit=10", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected runtime posture linkage 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response runtimePostureLinkageResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode runtime posture linkage: %v", err)
	}
	if response.SchemaVersion != runtimePostureLinkageSchemaVersion {
		t.Fatalf("expected runtime posture linkage schema version, got %#v", response)
	}
	if response.Summary.TotalSubjects == 0 || response.Summary.RuntimeModuleReady == 0 {
		t.Fatalf("expected non-empty posture linkage summary, got %#v", response.Summary)
	}
	if len(response.Semantics.SchedulingDecisionModel) == 0 || len(response.Semantics.MismatchModel) == 0 {
		t.Fatalf("expected explicit posture linkage semantics, got %#v", response.Semantics)
	}
	if !hasPostureLinkageMismatchSemantics(response.Semantics.MismatchModel, runtimeMismatchAttestation) {
		t.Fatalf("expected attestation-aware mismatch semantics, got %#v", response.Semantics.MismatchModel)
	}

	edge := findRuntimePosture(t, response.Items, "edge-gateway")
	if edge.SchedulingGuidance.Decision != runtimeSchedulingIsolatedReview {
		t.Fatalf("expected isolated review scheduling posture for edge-gateway, got %#v", edge)
	}
	if response.Summary.MismatchCounts[runtimeMismatchCriticalFindings] == 0 {
		t.Fatalf("expected critical finding mismatch summary, got %#v", response.Summary)
	}
}

func TestRuntimeBoundaryDisciplineContract(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/boundaries?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected runtime boundaries 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response runtimeBoundaryDisciplineResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode runtime boundaries: %v", err)
	}
	if response.SchemaVersion != runtimeBoundaryDisciplineSchema {
		t.Fatalf("expected runtime boundary discipline schema version, got %#v", response)
	}
	if response.SignalPath.CurrentPathModel == "" || !containsString(response.SignalPath.UnsupportedClaims, "zero-latency monitoring") {
		t.Fatalf("expected bounded signal path and unsupported claim list, got %#v", response.SignalPath)
	}

	preExecution := findRuntimeBoundaryPhase(t, response.EnforcementPhases, runtimeBoundaryPhasePreExecution)
	if len(preExecution.SupportedRulePacks) == 0 || !containsSubstringLocal(preExecution.Limitations, "not universal process blocking") {
		t.Fatalf("expected bounded pre-execution guidance semantics, got %#v", preExecution)
	}

	memoryBoundary := findRuntimeCoverageBoundary(t, response.CoverageBoundaries, runtimeBoundaryCoverageMemoryFileless)
	if memoryBoundary.CoverageState != "bounded_anomaly_signal" || !containsSubstringLocal(memoryBoundary.Limitations, "universal real-time memory scanning") {
		t.Fatalf("expected bounded memory/fileless coverage semantics, got %#v", memoryBoundary)
	}

	if response.OverheadCeiling.MeasurementStatus != "starting_points_only_not_benchmark_guarantee" || len(response.OverheadCeiling.StartingPointRefs) == 0 {
		t.Fatalf("expected explicit overhead ceiling discipline, got %#v", response.OverheadCeiling)
	}
}

func hasPostureLinkageMismatchSemantics(items []runtimePostureLinkageMismatchSemantics, code string) bool {
	for _, item := range items {
		if item.Code == code {
			return true
		}
	}
	return false
}

func findRuntimeBoundaryPhase(t *testing.T, items []runtimeBoundaryEnforcementPhase, phase string) runtimeBoundaryEnforcementPhase {
	t.Helper()
	for _, item := range items {
		if item.Phase == phase {
			return item
		}
	}
	t.Fatalf("expected runtime boundary phase %q, got %#v", phase, items)
	return runtimeBoundaryEnforcementPhase{}
}

func findRuntimeCoverageBoundary(t *testing.T, items []runtimeBoundaryCoverage, boundaryID string) runtimeBoundaryCoverage {
	t.Helper()
	for _, item := range items {
		if item.BoundaryID == boundaryID {
			return item
		}
	}
	t.Fatalf("expected runtime coverage boundary %q, got %#v", boundaryID, items)
	return runtimeBoundaryCoverage{}
}

func containsSubstringLocal(values []string, needle string) bool {
	needle = strings.TrimSpace(needle)
	for _, value := range values {
		if strings.Contains(value, needle) {
			return true
		}
	}
	return false
}

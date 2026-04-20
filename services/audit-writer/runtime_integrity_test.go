package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
)

func TestRuntimeIntegrityStateFindingsAndEnforcement(t *testing.T) {
	fixture := forensicsTestFixture(t)

	integrityReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/integrity?tenant_id=acme&environment=prod&limit=10", nil)
	integrityReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	integrityRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(integrityRec, integrityReq)
	if integrityRec.Code != http.StatusOK {
		t.Fatalf("expected runtime integrity 200, got %d: %s", integrityRec.Code, integrityRec.Body.String())
	}

	var integrity runtimeIntegrityListResponse
	if err := json.NewDecoder(integrityRec.Body).Decode(&integrity); err != nil {
		t.Fatalf("decode runtime integrity: %v", err)
	}
	edgeState := findRuntimeIntegrityState(t, integrity.Items, "edge-gateway")
	if edgeState.CurrentSandboxClass != runtimeSandboxClassIsolatedReview {
		t.Fatalf("expected edge-gateway to be routed into isolated review sandbox, got %#v", edgeState)
	}
	if edgeState.SBOMVerification.Status != runtimeSBOMStatusVerified {
		t.Fatalf("expected verified runtime-to-SBOM result, got %#v", edgeState.SBOMVerification)
	}
	if edgeState.DriftLevel != runtimeDriftLevelCritical {
		t.Fatalf("expected critical drift level for runtime binary exec, got %#v", edgeState)
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
	binaryFinding := findRuntimeFinding(t, findings.Items, runtimeFindingUnknownBinaryExec, "edge-gateway")
	if binaryFinding.RecommendedAction != runtimeActionApplyNetworkIsolation {
		t.Fatalf("expected network isolation recommendation for unknown binary exec, got %#v", binaryFinding)
	}
	if binaryFinding.ForensicContextURI == "" || len(binaryFinding.ReadbackRefs) == 0 {
		t.Fatalf("expected forensic and readback lineage on runtime finding, got %#v", binaryFinding)
	}
	outboundFinding := findRuntimeFinding(t, findings.Items, runtimeFindingOutboundDrift, "edge-gateway")
	if outboundFinding.Severity != "medium" {
		t.Fatalf("expected medium outbound drift severity, got %#v", outboundFinding)
	}

	profileReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/profiles/"+url.PathEscape(binaryFinding.SubjectRef)+"?tenant_id=acme&environment=prod", nil)
	profileReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	profileRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(profileRec, profileReq)
	if profileRec.Code != http.StatusOK {
		t.Fatalf("expected runtime profile 200, got %d: %s", profileRec.Code, profileRec.Body.String())
	}

	var profile runtimeIntegrityProfile
	if err := json.NewDecoder(profileRec.Body).Decode(&profile); err != nil {
		t.Fatalf("decode runtime profile: %v", err)
	}
	if !containsString(profile.ExpectedSigners, "signer-new") || !containsString(profile.AllowedExecPaths, "/app/*") {
		t.Fatalf("expected explainable signer and exec path profile, got %#v", profile)
	}

	evaluateReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/runtime/enforcement/evaluate?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`"}`),
	)
	evaluateReq.Header.Set("Authorization", "Bearer operator-demo-token")
	evaluateReq.Header.Set("Content-Type", "application/json")
	evaluateRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(evaluateRec, evaluateReq)
	if evaluateRec.Code != http.StatusOK {
		t.Fatalf("expected runtime enforcement evaluate 200, got %d: %s", evaluateRec.Code, evaluateRec.Body.String())
	}

	var evaluated runtimeEnforcementDecision
	if err := json.NewDecoder(evaluateRec.Body).Decode(&evaluated); err != nil {
		t.Fatalf("decode runtime enforcement evaluate: %v", err)
	}
	if evaluated.Action != runtimeActionApplyNetworkIsolation || evaluated.ApprovalMode != recommendationApprovalHumanReview || evaluated.Executed {
		t.Fatalf("expected approval-gated network isolation evaluation, got %#v", evaluated)
	}
	if evaluated.TopologyContext == nil {
		t.Fatalf("expected topology-aware containment evaluation, got %#v", evaluated)
	}

	quarantinePendingReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/runtime/quarantine?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`"}`),
	)
	quarantinePendingReq.Header.Set("Authorization", "Bearer operator-demo-token")
	quarantinePendingReq.Header.Set("Content-Type", "application/json")
	quarantinePendingRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(quarantinePendingRec, quarantinePendingReq)
	if quarantinePendingRec.Code != http.StatusOK {
		t.Fatalf("expected runtime quarantine evaluate 200, got %d: %s", quarantinePendingRec.Code, quarantinePendingRec.Body.String())
	}

	var pending runtimeEnforcementDecision
	if err := json.NewDecoder(quarantinePendingRec.Body).Decode(&pending); err != nil {
		t.Fatalf("decode approval-pending runtime quarantine: %v", err)
	}
	if pending.Executed || pending.ExecutionResult != "approval_pending" {
		t.Fatalf("expected approval pending quarantine, got %#v", pending)
	}

	quarantineReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/runtime/quarantine?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-9i-1"}`),
	)
	quarantineReq.Header.Set("Authorization", "Bearer operator-demo-token")
	quarantineReq.Header.Set("Content-Type", "application/json")
	quarantineRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(quarantineRec, quarantineReq)
	if quarantineRec.Code != http.StatusOK {
		t.Fatalf("expected runtime quarantine execute 200, got %d: %s", quarantineRec.Code, quarantineRec.Body.String())
	}

	var quarantined runtimeEnforcementDecision
	if err := json.NewDecoder(quarantineRec.Body).Decode(&quarantined); err != nil {
		t.Fatalf("decode runtime quarantine execute: %v", err)
	}
	if !quarantined.Executed || quarantined.ExecutionResult != "network_isolation_applied" {
		t.Fatalf("expected network isolation execution result, got %#v", quarantined)
	}

	forensicReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/runtime/forensic-snapshot?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`"}`),
	)
	forensicReq.Header.Set("Authorization", "Bearer operator-demo-token")
	forensicReq.Header.Set("Content-Type", "application/json")
	forensicRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(forensicRec, forensicReq)
	if forensicRec.Code != http.StatusOK {
		t.Fatalf("expected runtime forensic snapshot 200, got %d: %s", forensicRec.Code, forensicRec.Body.String())
	}

	var snapshotDecision runtimeEnforcementDecision
	if err := json.NewDecoder(forensicRec.Body).Decode(&snapshotDecision); err != nil {
		t.Fatalf("decode runtime forensic snapshot: %v", err)
	}
	if !snapshotDecision.Executed || snapshotDecision.ForensicContextURI == "" {
		t.Fatalf("expected forensic snapshot request with lineage URI, got %#v", snapshotDecision)
	}

	enforcementReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/enforcement?tenant_id=acme&environment=prod&limit=10", nil)
	enforcementReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	enforcementRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(enforcementRec, enforcementReq)
	if enforcementRec.Code != http.StatusOK {
		t.Fatalf("expected runtime enforcement history 200, got %d: %s", enforcementRec.Code, enforcementRec.Body.String())
	}

	var history runtimeEnforcementListResponse
	if err := json.NewDecoder(enforcementRec.Body).Decode(&history); err != nil {
		t.Fatalf("decode runtime enforcement history: %v", err)
	}
	if !containsRuntimeAction(history.Items, runtimeActionApplyNetworkIsolation) || !containsRuntimeAction(history.Items, runtimeActionCaptureForensics) {
		t.Fatalf("expected enforcement history to retain quarantine and forensic actions, got %#v", history)
	}
}

func TestRuntimeReadbackRecommendationsAndHandoffRuntimeContext(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	incidentID := fetchIncidentForWorkload(t, fixture.handler, "edge-gateway")

	readbackReq := httptest.NewRequest(http.MethodGet, "/v1/incidents/"+incidentID+"/policy-replay?tenant_id=acme&environment=prod", nil)
	readbackReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	readbackRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(readbackRec, readbackReq)
	if readbackRec.Code != http.StatusOK {
		t.Fatalf("expected incident policy replay 200, got %d: %s", readbackRec.Code, readbackRec.Body.String())
	}

	var assessment policyReplayAssessment
	if err := json.NewDecoder(readbackRec.Body).Decode(&assessment); err != nil {
		t.Fatalf("decode policy replay assessment: %v", err)
	}
	if assessment.Readback.ResourceID == "" {
		t.Fatalf("expected readback ref from policy replay, got %#v", assessment)
	}

	runtimeReadbackReq := httptest.NewRequest(http.MethodGet, "/v1/readback/policy-replay/"+assessment.Readback.ResourceID+"/runtime-context", nil)
	runtimeReadbackReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	runtimeReadbackRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(runtimeReadbackRec, runtimeReadbackReq)
	if runtimeReadbackRec.Code != http.StatusOK {
		t.Fatalf("expected runtime readback context 200, got %d: %s", runtimeReadbackRec.Code, runtimeReadbackRec.Body.String())
	}

	var runtimeReadback readbackRuntimeResponse
	if err := json.NewDecoder(runtimeReadbackRec.Body).Decode(&runtimeReadback); err != nil {
		t.Fatalf("decode runtime readback context: %v", err)
	}
	if runtimeReadback.RuntimeContextURI == "" || len(runtimeReadback.Workloads) == 0 || len(runtimeReadback.Findings) == 0 {
		t.Fatalf("expected workload and finding runtime context for readback scope, got %#v", runtimeReadback)
	}

	recommendationReq := httptest.NewRequest(http.MethodGet, "/v1/recommendations?tenant_id=acme&environment=prod&source_type=runtime_signal", nil)
	recommendationReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	recommendationRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(recommendationRec, recommendationReq)
	if recommendationRec.Code != http.StatusOK {
		t.Fatalf("expected runtime recommendations 200, got %d: %s", recommendationRec.Code, recommendationRec.Body.String())
	}

	var recommendations recommendationListResponse
	if err := json.NewDecoder(recommendationRec.Body).Decode(&recommendations); err != nil {
		t.Fatalf("decode runtime recommendations: %v", err)
	}
	if len(recommendations.Recommendations) == 0 {
		t.Fatalf("expected at least one runtime recommendation, got %#v", recommendations)
	}
	item := recommendations.Recommendations[0]
	if item.SourceType != "runtime_signal" || item.ApprovalMode == "" || len(item.VerificationPlan) == 0 {
		t.Fatalf("expected runtime recommendation workflow metadata, got %#v", item)
	}
	if !strings.Contains(strings.ToLower(item.Rationale), "runtime") && !strings.Contains(strings.ToLower(item.Rationale), "containment") {
		t.Fatalf("expected runtime recommendation rationale, got %#v", item.Rationale)
	}

	sealReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/handoff/seal?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"audience":"internal","include_runtime":true}`),
	)
	sealReq.Header.Set("Authorization", "Bearer operator-demo-token")
	sealReq.Header.Set("Content-Type", "application/json")
	sealRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(sealRec, sealReq)
	if sealRec.Code != http.StatusOK {
		t.Fatalf("expected runtime-inclusive handoff seal 200, got %d: %s", sealRec.Code, sealRec.Body.String())
	}

	var sealed handoffSealResponse
	if err := json.NewDecoder(sealRec.Body).Decode(&sealed); err != nil {
		t.Fatalf("decode runtime-inclusive handoff: %v", err)
	}
	if !handoffManifestHasArtifact(sealed.Manifest, "evidence/runtime_context.json") {
		t.Fatalf("expected runtime context artifact in sealed handoff, got %#v", sealed.Manifest.Artifacts)
	}
}

func TestRuntimeIntegrityMarksTelemetryGapAsUnverifiable(t *testing.T) {
	fixture := forensicsTestFixture(t)

	if _, err := fixture.store.Ingest(context.Background(), audit.Event{
		RequestID:      "runtime-gap-observation",
		Timestamp:      fixture.currentTimestamp.Add(-15 * time.Minute),
		Component:      "runtime-agent",
		EventType:      audit.EventTypeRuntimeObservationRecorded,
		Decision:       audit.DecisionAllow,
		TenantID:       "acme",
		ClusterID:      "local",
		Repo:           "acme/platform-gap",
		Environment:    "prod",
		Namespace:      "acme-prod",
		WorkloadKind:   "Deployment",
		Workload:       "telemetry-gap",
		ServiceAccount: "gap-sa",
		Digest:         "sha256:gap-v1",
		Reasons:        []string{"runtime observation without corresponding desired or active digest pair"},
		RuntimeIntegrity: canonicalJSONMust(runtimeIntegrityEventPayload{
			Observation: &runtimeObservationPayload{
				Node:        "node-gap",
				Pod:         "telemetry-gap-v1-0",
				ContainerID: "ctr-gap",
				EventType:   "binary_exec",
				Confidence:  runtimeConfidenceMedium,
				EventPayload: map[string]any{
					"binary_path": "/tmp/tool",
				},
			},
		}),
	}); err != nil {
		t.Fatalf("ingest telemetry-gap observation: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/integrity?tenant_id=acme&environment=prod&workload=telemetry-gap", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected runtime integrity 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var integrity runtimeIntegrityListResponse
	if err := json.NewDecoder(rec.Body).Decode(&integrity); err != nil {
		t.Fatalf("decode runtime integrity gap response: %v", err)
	}
	state := findRuntimeIntegrityState(t, integrity.Items, "telemetry-gap")
	if state.SBOMVerification.Status != runtimeSBOMStatusUnverifiable {
		t.Fatalf("expected unverifiable SBOM status for telemetry gap, got %#v", state.SBOMVerification)
	}
	if !strings.Contains(strings.ToLower(strings.Join(state.SBOMVerification.Limitations, " ")), "unverifiable") {
		t.Fatalf("expected explicit telemetry gap limitation, got %#v", state.SBOMVerification.Limitations)
	}
}

func findRuntimeIntegrityState(t *testing.T, items []runtimeIntegrityState, workload string) runtimeIntegrityState {
	t.Helper()
	for _, item := range items {
		if strings.Contains(item.SubjectRef, workload) {
			return item
		}
	}
	t.Fatalf("runtime integrity state for %s not found", workload)
	return runtimeIntegrityState{}
}

func findRuntimeFinding(t *testing.T, items []runtimeIntegrityFinding, findingType, workload string) runtimeIntegrityFinding {
	t.Helper()
	for _, item := range items {
		if item.FindingType == findingType && strings.Contains(item.SubjectRef, workload) {
			return item
		}
	}
	t.Fatalf("runtime finding %s for %s not found", findingType, workload)
	return runtimeIntegrityFinding{}
}

func containsRuntimeAction(items []runtimeEnforcementDecision, action string) bool {
	for _, item := range items {
		if item.Action == action {
			return true
		}
	}
	return false
}

func handoffManifestHasArtifact(manifest sealedManifest, path string) bool {
	for _, artifact := range manifest.Artifacts {
		if artifact.Path == path {
			return true
		}
	}
	return false
}

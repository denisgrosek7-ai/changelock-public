package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	internalvex "github.com/denisgrosek/changelock/internal/vex"
)

func TestForensicsStateDeltaTimelineReplayAndFlashback(t *testing.T) {
	fixture := forensicsTestFixture(t)

	stateReq := httptest.NewRequest(http.MethodGet, "/v1/forensics/state?tenant_id=acme&environment=prod&service=edge-gateway&timestamp="+fixture.historicalTimestamp.Format(time.RFC3339), nil)
	stateReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	stateRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(stateRec, stateReq)
	if stateRec.Code != http.StatusOK {
		t.Fatalf("expected forensic state 200, got %d: %s", stateRec.Code, stateRec.Body.String())
	}

	var state pointInTimeState
	if err := json.NewDecoder(stateRec.Body).Decode(&state); err != nil {
		t.Fatalf("decode forensic state: %v", err)
	}
	if state.Mode != forensicsModeHistoricalReconstruction || state.PolicyContext.PolicyBundleHash == "" {
		t.Fatalf("expected historical reconstruction with policy context, got %#v", state)
	}
	if state.SchemaVersion != forensicsStateSchemaVersion {
		t.Fatalf("expected schema-versioned forensic state, got %#v", state)
	}
	if len(state.VulnerabilityContext.UnknownLaterDisclosedRefs) == 0 {
		t.Fatalf("expected later disclosures to stay separate from historical known-state, got %#v", state.VulnerabilityContext)
	}
	if state.TopologyContext == nil || !state.TopologyContext.AdvisoryOnly {
		t.Fatalf("expected advisory topology context in forensic state, got %#v", state.TopologyContext)
	}

	deltaReq := httptest.NewRequest(http.MethodGet, "/v1/forensics/delta?tenant_id=acme&environment=prod&service=edge-gateway&t1="+fixture.historicalTimestamp.Format(time.RFC3339)+"&t2="+fixture.currentTimestamp.Format(time.RFC3339), nil)
	deltaReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	deltaRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(deltaRec, deltaReq)
	if deltaRec.Code != http.StatusOK {
		t.Fatalf("expected forensic delta 200, got %d: %s", deltaRec.Code, deltaRec.Body.String())
	}

	var delta timeDeltaResult
	if err := json.NewDecoder(deltaRec.Body).Decode(&delta); err != nil {
		t.Fatalf("decode forensic delta: %v", err)
	}
	if delta.Mode != forensicsModeTimeDelta || len(delta.TopologyDelta) == 0 {
		t.Fatalf("expected time delta with topology drift, got %#v", delta)
	}
	if delta.SchemaVersion != forensicsDeltaSchemaVersion {
		t.Fatalf("expected schema-versioned forensic delta, got %#v", delta)
	}
	if len(delta.PolicyDelta.Modified) == 0 && len(delta.PolicyDelta.Added) == 0 {
		t.Fatalf("expected policy delta in forensic compare, got %#v", delta.PolicyDelta)
	}

	timelineReq := httptest.NewRequest(http.MethodGet, "/v1/forensics/timeline?tenant_id=acme&environment=prod&t1="+fixture.historicalTimestamp.Format(time.RFC3339)+"&t2="+fixture.currentTimestamp.Format(time.RFC3339), nil)
	timelineReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	timelineRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(timelineRec, timelineReq)
	if timelineRec.Code != http.StatusOK {
		t.Fatalf("expected forensic timeline 200, got %d: %s", timelineRec.Code, timelineRec.Body.String())
	}

	var timeline forensicTimelineResponse
	if err := json.NewDecoder(timelineRec.Body).Decode(&timeline); err != nil {
		t.Fatalf("decode forensic timeline: %v", err)
	}
	if len(timeline.Markers) == 0 {
		t.Fatalf("expected evidence-backed timeline markers, got %#v", timeline)
	}
	if timeline.SchemaVersion != forensicsTimelineSchemaVersion {
		t.Fatalf("expected schema-versioned forensic timeline, got %#v", timeline)
	}

	replayReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/forensics/replay?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"service":"edge-gateway","timestamp":"`+fixture.historicalTimestamp.Format(time.RFC3339)+`","replay_mode":"modern_full_stack_replay"}`),
	)
	replayReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	replayReq.Header.Set("Content-Type", "application/json")
	replayRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(replayRec, replayReq)
	if replayRec.Code != http.StatusOK {
		t.Fatalf("expected forensic replay 200, got %d: %s", replayRec.Code, replayRec.Body.String())
	}

	var replay forensicReplayResponse
	if err := json.NewDecoder(replayRec.Body).Decode(&replay); err != nil {
		t.Fatalf("decode forensic replay: %v", err)
	}
	if replay.Mode != forensicsModeCounterfactualReplay || !replay.Counterfactual {
		t.Fatalf("expected counterfactual replay response, got %#v", replay)
	}
	if replay.SchemaVersion != forensicsReplaySchemaVersion {
		t.Fatalf("expected schema-versioned forensic replay, got %#v", replay)
	}
	if replay.VerdictDelta == "no_change" {
		t.Fatalf("expected replay verdict delta from historical state, got %#v", replay)
	}
	if !strings.Contains(strings.ToLower(strings.Join(replay.Limitations, " ")), "counterfactual") {
		t.Fatalf("expected replay limitations to declare simulation semantics, got %#v", replay.Limitations)
	}

	flashbackReq := httptest.NewRequest(http.MethodGet, "/v1/forensics/vex-flashback?tenant_id=acme&environment=prod&image_digest=sha256:edge-vex&timestamp="+fixture.flashbackTimestamp.Format(time.RFC3339), nil)
	flashbackReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	flashbackRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(flashbackRec, flashbackReq)
	if flashbackRec.Code != http.StatusOK {
		t.Fatalf("expected vex flashback 200, got %d: %s", flashbackRec.Code, flashbackRec.Body.String())
	}

	var flashback vexFlashbackResponse
	if err := json.NewDecoder(flashbackRec.Body).Decode(&flashback); err != nil {
		t.Fatalf("decode vex flashback: %v", err)
	}
	if len(flashback.HistoricalVulnerabilityState) == 0 || len(flashback.VEXFlashback) == 0 {
		t.Fatalf("expected historical vuln state and active vex flashback, got %#v", flashback)
	}
	if flashback.SchemaVersion != forensicsVEXFlashbackSchemaVersion {
		t.Fatalf("expected schema-versioned vex flashback, got %#v", flashback)
	}
	if len(flashback.DisclosedAfterTRefs) == 0 {
		t.Fatalf("expected disclosed-after-T refs in vex flashback, got %#v", flashback)
	}
}

func TestForensicsStateAndDeltaAreDeterministicForSameInput(t *testing.T) {
	fixture := forensicsTestFixture(t)

	buildState := func() pointInTimeState {
		req := httptest.NewRequest(http.MethodGet, "/v1/forensics/state?tenant_id=acme&environment=prod&service=edge-gateway&timestamp="+fixture.historicalTimestamp.Format(time.RFC3339), nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		fixture.handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected forensic state 200, got %d: %s", rec.Code, rec.Body.String())
		}
		var state pointInTimeState
		if err := json.NewDecoder(rec.Body).Decode(&state); err != nil {
			t.Fatalf("decode forensic state: %v", err)
		}
		return state
	}
	firstState := buildState()
	secondState := buildState()
	if string(canonicalJSONMust(firstState)) != string(canonicalJSONMust(secondState)) {
		t.Fatalf("expected point-in-time forensic state to stay deterministic for the same input")
	}

	buildDelta := func() timeDeltaResult {
		req := httptest.NewRequest(http.MethodGet, "/v1/forensics/delta?tenant_id=acme&environment=prod&service=edge-gateway&t1="+fixture.historicalTimestamp.Format(time.RFC3339)+"&t2="+fixture.currentTimestamp.Format(time.RFC3339), nil)
		req.Header.Set("Authorization", "Bearer viewer-demo-token")
		rec := httptest.NewRecorder()
		fixture.handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected forensic delta 200, got %d: %s", rec.Code, rec.Body.String())
		}
		var delta timeDeltaResult
		if err := json.NewDecoder(rec.Body).Decode(&delta); err != nil {
			t.Fatalf("decode forensic delta: %v", err)
		}
		return delta
	}
	firstDelta := buildDelta()
	secondDelta := buildDelta()
	if string(canonicalJSONMust(firstDelta)) != string(canonicalJSONMust(secondDelta)) {
		t.Fatalf("expected forensic delta to stay deterministic for the same input")
	}
}

func TestReadbackForensicContextAndForensicRecommendations(t *testing.T) {
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
		t.Fatalf("expected policy replay readback ref, got %#v", assessment)
	}

	forensicReadbackReq := httptest.NewRequest(http.MethodGet, "/v1/readback/policy-replay/"+assessment.Readback.ResourceID+"/forensic-context", nil)
	forensicReadbackReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	forensicReadbackRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(forensicReadbackRec, forensicReadbackReq)
	if forensicReadbackRec.Code != http.StatusOK {
		t.Fatalf("expected readback forensic context 200, got %d: %s", forensicReadbackRec.Code, forensicReadbackRec.Body.String())
	}

	var forensicReadback readbackForensicResponse
	if err := json.NewDecoder(forensicReadbackRec.Body).Decode(&forensicReadback); err != nil {
		t.Fatalf("decode readback forensic context: %v", err)
	}
	if forensicReadback.PointInTimeState.Mode != forensicsModeHistoricalReconstruction || forensicReadback.ForensicContextURI == "" {
		t.Fatalf("expected historical forensic context from readback, got %#v", forensicReadback)
	}

	recommendationReq := httptest.NewRequest(http.MethodGet, "/v1/recommendations?tenant_id=acme&environment=prod&source_type=forensic_signal&service=edge-gateway", nil)
	recommendationReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	recommendationRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(recommendationRec, recommendationReq)
	if recommendationRec.Code != http.StatusOK {
		t.Fatalf("expected forensic recommendations 200, got %d: %s", recommendationRec.Code, recommendationRec.Body.String())
	}

	var recommendations recommendationListResponse
	if err := json.NewDecoder(recommendationRec.Body).Decode(&recommendations); err != nil {
		t.Fatalf("decode forensic recommendations: %v", err)
	}
	if len(recommendations.Recommendations) == 0 {
		t.Fatalf("expected forensic recommendation candidate, got %#v", recommendations)
	}
	item := recommendations.Recommendations[0]
	if item.SourceType != "forensic_signal" || item.ApprovalMode == "" || len(item.VerificationPlan) == 0 {
		t.Fatalf("expected forensic recommendation with workflow metadata, got %#v", item)
	}
	if !strings.Contains(strings.ToLower(item.Rationale), "historical verdict") {
		t.Fatalf("expected forensic rationale to explain replay delta, got %#v", item.Rationale)
	}
}

func TestForensicsStateReturnsLimitationsForEmptyScope(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/forensics/state?tenant_id=acme&environment=prod&service=missing-service&timestamp="+fixture.historicalTimestamp.Format(time.RFC3339), nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected forensic state 200 for empty scope, got %d: %s", rec.Code, rec.Body.String())
	}

	var state pointInTimeState
	if err := json.NewDecoder(rec.Body).Decode(&state); err != nil {
		t.Fatalf("decode empty-scope forensic state: %v", err)
	}
	if state.SchemaVersion != forensicsStateSchemaVersion {
		t.Fatalf("expected schema-versioned empty forensic state, got %#v", state)
	}
	if len(state.EvidenceRefs) != 0 {
		t.Fatalf("expected no evidence refs for empty scope, got %#v", state.EvidenceRefs)
	}
	if len(state.Limitations) == 0 {
		t.Fatalf("expected bounded reconstruction limitations for empty scope, got %#v", state)
	}
}

type forensicsFixtureData struct {
	handler             http.Handler
	store               audit.Store
	historicalTimestamp time.Time
	currentTimestamp    time.Time
	flashbackTimestamp  time.Time
}

func forensicsTestFixture(t testing.TB) forensicsFixtureData {
	t.Helper()
	t.Setenv("CHANGELOCK_AUTH_MODE", auth.ModeStaticToken)
	t.Setenv("CHANGELOCK_AUTH_TOKENS_JSON", testAuthTokensJSON())

	authConfig, err := auth.ParseConfig(auth.ModeStaticToken, testAuthTokensJSON())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}
	store := audit.NewMemoryStore()
	now := time.Now().UTC()
	historicalTimestamp := now.Add(-7 * 24 * time.Hour)
	currentTimestamp := now
	flashbackTimestamp := now.Add(6 * time.Hour)

	events := []audit.Event{
		{
			RequestID:        "forensics-edge-baseline",
			Timestamp:        now.Add(-14 * 24 * time.Hour),
			Component:        "deploy-gate",
			EventType:        audit.EventTypeDeployGateDecision,
			Decision:         audit.DecisionAllow,
			TenantID:         "acme",
			Repo:             "acme/platform-edge",
			Environment:      "prod",
			Namespace:        "acme-prod",
			Workload:         "edge-gateway",
			ServiceAccount:   "edge-sa",
			Digest:           "sha256:edge-v1",
			PolicyBundleHash: "bundle-old",
			PolicyVersion:    "2026.03.1",
			Reasons:          []string{"approved release"},
			Evidence: &audit.Evidence{
				Artifact: &audit.ArtifactEvidence{
					SignerIdentity: "signer-old",
					Issuer:         "root-a",
					SBOMHash:       "sbom-edge-v1",
				},
			},
		},
		{
			RequestID:        "forensics-edge-current",
			Timestamp:        now.Add(-36 * time.Hour),
			Component:        "deploy-gate",
			EventType:        audit.EventTypeDeployGateDecision,
			Decision:         audit.DecisionDeny,
			TenantID:         "acme",
			Repo:             "acme/platform-edge",
			Environment:      "prod",
			Namespace:        "acme-prod",
			Workload:         "edge-gateway",
			ServiceAccount:   "edge-sa",
			Digest:           "sha256:edge-v2",
			CVEID:            "CVE-2026-9000",
			PolicyBundleHash: "bundle-new",
			PolicyVersion:    "2026.04.2",
			Reasons:          []string{"workflow mismatch", "policy bundle tightened"},
			Evidence: &audit.Evidence{
				Artifact: &audit.ArtifactEvidence{
					SignerIdentity: "signer-new",
					Issuer:         "root-b",
					SBOMHash:       "sbom-edge-v2",
					VulnerabilitySummary: &audit.VulnerabilitySummary{
						Critical: 1,
						Total:    1,
					},
				},
			},
		},
		{
			RequestID:                "runtime-edge-desired",
			Timestamp:                now.Add(-48 * time.Hour),
			Component:                "runtime-agent",
			EventType:                audit.EventTypeRuntimeDesiredStateRecorded,
			Decision:                 audit.DecisionAllow,
			TenantID:                 "acme",
			ClusterID:                "local",
			Repo:                     "acme/platform-edge",
			Environment:              "prod",
			Namespace:                "acme-prod",
			WorkloadKind:             "Deployment",
			Workload:                 "edge-gateway",
			ServiceAccount:           "edge-sa",
			Digest:                   "sha256:edge-v2",
			DesiredStateSourceRef:    "deploy:edge-gateway:v2",
			DesiredStateApprovalID:   "runtime-approval-edge",
			DesiredStateVerification: "verified",
			Reasons:                  []string{"runtime desired state approved"},
			Evidence: &audit.Evidence{
				Artifact: &audit.ArtifactEvidence{
					SignerIdentity:           "signer-new",
					Issuer:                   "root-b",
					AttestationPredicate:     "https://slsa.dev/provenance/v1",
					AttestationSubjectDigest: "sha256:edge-v2",
					SBOMHash:                 "sbom-edge-v2",
				},
				Runtime: &audit.RuntimeEvidence{
					ApprovedDigest:         "sha256:edge-v2",
					ExpectedConfigHash:     "cfg-edge-v2",
					ServiceAccountExpected: "edge-sa",
					ApprovedContainers: []audit.RuntimeApprovedContainer{
						{
							Name:           "edge-gateway",
							ApprovedDigest: "sha256:edge-v2",
							Runtime: audit.RuntimeSecurityConstraints{
								RunAsNonRoot:           true,
								ReadOnlyRootFilesystem: true,
								DropAllCapabilities:    true,
								SeccompRuntimeDefault:  true,
								DenyPrivileged:         true,
							},
						},
					},
				},
			},
		},
		{
			RequestID:                "runtime-edge-active",
			Timestamp:                now.Add(-12 * time.Hour),
			Component:                "runtime-agent",
			EventType:                audit.EventTypeRuntimeActiveStateObserved,
			Decision:                 audit.DecisionAllow,
			TenantID:                 "acme",
			ClusterID:                "local",
			Repo:                     "acme/platform-edge",
			Environment:              "prod",
			Namespace:                "acme-prod",
			WorkloadKind:             "Deployment",
			Workload:                 "edge-gateway",
			ServiceAccount:           "edge-sa",
			Digest:                   "sha256:edge-v2",
			ReconciliationStatus:     "in_sync",
			DesiredStateSourceRef:    "deploy:edge-gateway:v2",
			DesiredStateApprovalID:   "runtime-approval-edge",
			DesiredStateVerification: "verified",
			Reasons:                  []string{"observed healthy runtime state"},
			Evidence: &audit.Evidence{
				Artifact: &audit.ArtifactEvidence{
					SignerIdentity:           "signer-new",
					Issuer:                   "root-b",
					AttestationPredicate:     "https://slsa.dev/provenance/v1",
					AttestationSubjectDigest: "sha256:edge-v2",
					SBOMHash:                 "sbom-edge-v2",
				},
				Runtime: &audit.RuntimeEvidence{
					ApprovedDigest:         "sha256:edge-v2",
					RunningDigest:          "sha256:edge-v2",
					ExpectedConfigHash:     "cfg-edge-v2",
					ActualConfigHash:       "cfg-edge-v2",
					ServiceAccountExpected: "edge-sa",
					ServiceAccountObserved: "edge-sa",
				},
			},
		},
		{
			RequestID:      "runtime-edge-outbound",
			Timestamp:      now.Add(-10 * time.Hour),
			Component:      "runtime-agent",
			EventType:      audit.EventTypeRuntimeObservationRecorded,
			Decision:       audit.DecisionAllow,
			TenantID:       "acme",
			ClusterID:      "local",
			Repo:           "acme/platform-edge",
			Environment:    "prod",
			Namespace:      "acme-prod",
			WorkloadKind:   "Deployment",
			Workload:       "edge-gateway",
			ServiceAccount: "edge-sa",
			Digest:         "sha256:edge-v2",
			Reasons:        []string{"new outbound destination detected"},
			RuntimeIntegrity: canonicalJSONMust(runtimeIntegrityEventPayload{
				Observation: &runtimeObservationPayload{
					Node:        "node-a",
					Pod:         "edge-gateway-v2-7d9d5",
					ContainerID: "ctr-edge-v2",
					EventType:   "outbound_connect",
					Confidence:  runtimeConfidenceMedium,
					EventPayload: map[string]any{
						"destination":      "198.51.100.42:443",
						"destination_type": "external_ip",
					},
					ProfileHint: &runtimeProfileHint{
						AllowedBinaries:        []string{"/app/edge-gateway"},
						AllowedExecPaths:       []string{"/app/*"},
						AllowedLibraryPatterns: []string{"libssl.so.3", "libcrypto.so.3"},
						AllowedNetworkPatterns: []string{"service:auth-api"},
					},
				},
			}),
		},
		{
			RequestID:      "runtime-edge-binary",
			Timestamp:      now.Add(-8 * time.Hour),
			Component:      "runtime-agent",
			EventType:      audit.EventTypeRuntimeObservationRecorded,
			Decision:       audit.DecisionDeny,
			TenantID:       "acme",
			ClusterID:      "local",
			Repo:           "acme/platform-edge",
			Environment:    "prod",
			Namespace:      "acme-prod",
			WorkloadKind:   "Deployment",
			Workload:       "edge-gateway",
			ServiceAccount: "edge-sa",
			Digest:         "sha256:edge-v2",
			Reasons:        []string{"unknown binary exec at /tmp/nc"},
			RuntimeIntegrity: canonicalJSONMust(runtimeIntegrityEventPayload{
				Observation: &runtimeObservationPayload{
					Node:        "node-a",
					Pod:         "edge-gateway-v2-7d9d5",
					ContainerID: "ctr-edge-v2",
					EventType:   "binary_exec",
					Confidence:  runtimeConfidenceHigh,
					EventPayload: map[string]any{
						"binary_path": "/tmp/nc",
						"argv":        []string{"/tmp/nc", "-zv", "198.51.100.42", "443"},
					},
				},
			}),
		},
		{
			RequestID:      "runtime-edge-sbom",
			Timestamp:      now.Add(-7 * time.Hour),
			Component:      "runtime-agent",
			EventType:      audit.EventTypeRuntimeSBOMVerificationRecorded,
			Decision:       audit.DecisionAllow,
			TenantID:       "acme",
			ClusterID:      "local",
			Repo:           "acme/platform-edge",
			Environment:    "prod",
			Namespace:      "acme-prod",
			WorkloadKind:   "Deployment",
			Workload:       "edge-gateway",
			ServiceAccount: "edge-sa",
			Digest:         "sha256:edge-v2",
			Reasons:        []string{"runtime sbom verification recorded"},
			RuntimeIntegrity: canonicalJSONMust(runtimeIntegrityEventPayload{
				SBOMVerification: &runtimeSBOMVerificationPayload{
					Status:              runtimeSBOMStatusVerified,
					MatchedArtifacts:    []string{"sha256:edge-v2"},
					ObservedLibraryRefs: []string{"libssl.so.3", "libcrypto.so.3"},
					Limitations:         []string{"Library verification is limited to observed loaded-state evidence in the current runtime window."},
				},
			}),
		},
		{
			RequestID:            "forensics-auth-runtime",
			Timestamp:            now.Add(-30 * time.Hour),
			Component:            "runtime-agent",
			EventType:            audit.EventTypeRuntimeActiveStateObserved,
			Decision:             audit.DecisionDeny,
			TenantID:             "acme",
			Repo:                 "acme/platform-auth",
			Environment:          "prod",
			Namespace:            "acme-prod",
			Workload:             "auth-api",
			ServiceAccount:       "edge-sa",
			Digest:               "sha256:auth-v2",
			DriftResult:          "service_account_drift",
			DriftClasses:         []string{"service_account_drift"},
			Reasons:              []string{"service account drift"},
			ReconciliationStatus: "drift_detected",
		},
		{
			RequestID:            "forensics-db-runtime",
			Timestamp:            now.Add(-28 * time.Hour),
			Component:            "runtime-agent",
			EventType:            audit.EventTypeRuntimeActiveStateObserved,
			Decision:             audit.DecisionAllow,
			TenantID:             "acme",
			Repo:                 "acme/platform-billing",
			Environment:          "prod",
			Namespace:            "acme-prod",
			Workload:             "billing-db",
			ServiceAccount:       "db-sa",
			Digest:               "sha256:billing-v1",
			ReconciliationStatus: "in_sync",
			Reasons:              []string{"observed healthy"},
		},
		{
			RequestID:      "forensics-vuln-known",
			Timestamp:      now.Add(2 * time.Hour),
			Component:      "deploy-gate",
			EventType:      audit.EventTypeDeployGateDecision,
			Decision:       audit.DecisionAllow,
			TenantID:       "acme",
			Repo:           "acme/platform-edge",
			Environment:    "prod",
			Namespace:      "acme-prod",
			Workload:       "edge-gateway",
			ServiceAccount: "edge-sa",
			Digest:         "sha256:edge-vex",
			CVEID:          "CVE-2026-1111",
			Reasons:        []string{"historical vulnerability evidence"},
			Evidence: &audit.Evidence{
				Artifact: &audit.ArtifactEvidence{
					SignerIdentity: "signer-vex",
					Issuer:         "root-a",
					SBOMHash:       "sbom-edge-vex",
					VulnerabilitySummary: &audit.VulnerabilitySummary{
						High:  1,
						Total: 1,
					},
				},
			},
		},
		{
			RequestID:      "forensics-vuln-later",
			Timestamp:      now.Add(12 * time.Hour),
			Component:      "deploy-gate",
			EventType:      audit.EventTypeDeployGateDecision,
			Decision:       audit.DecisionDeny,
			TenantID:       "acme",
			Repo:           "acme/platform-edge",
			Environment:    "prod",
			Namespace:      "acme-prod",
			Workload:       "edge-gateway",
			ServiceAccount: "edge-sa",
			Digest:         "sha256:edge-vex",
			CVEID:          "CVE-2026-2222",
			Reasons:        []string{"later disclosed vulnerability"},
			Evidence: &audit.Evidence{
				Artifact: &audit.ArtifactEvidence{
					SignerIdentity: "signer-vex",
					Issuer:         "root-a",
					SBOMHash:       "sbom-edge-vex",
					VulnerabilitySummary: &audit.VulnerabilitySummary{
						Critical: 1,
						Total:    1,
					},
				},
			},
		},
		{
			RequestID:   "forensics-vex-audit",
			Timestamp:   now.Add(3 * time.Hour),
			Component:   "vex-manager",
			EventType:   audit.EventTypeVEXStatementRecorded,
			Decision:    audit.DecisionAllow,
			TenantID:    "acme",
			Repo:        "acme/platform-edge",
			Environment: "prod",
			Namespace:   "acme-prod",
			Workload:    "edge-gateway",
			Digest:      "sha256:edge-vex",
			CVEID:       "CVE-2026-1111",
			Reasons:     []string{"vex statement recorded"},
		},
	}
	for _, event := range events {
		if _, err := store.Ingest(context.Background(), event); err != nil {
			t.Fatalf("Ingest() error = %v", err)
		}
	}

	if _, err := store.CreateVEXStatement(context.Background(), internalvex.CreateRequest{
		VulnerabilityID: "CVE-2026-1111",
		Status:          internalvex.StatusNotAffected,
		Justification:   "component not on vulnerable code path",
		SourceRef:       "urn:test:vex-flashback",
		Scope: internalvex.Scope{
			ImageDigest: "sha256:edge-vex",
			Repo:        "acme/platform-edge",
			Workload:    "edge-gateway",
			TenantID:    "acme",
			Environment: "prod",
			Namespace:   "acme-prod",
		},
	}, "test-suite"); err != nil {
		t.Fatalf("CreateVEXStatement() error = %v", err)
	}

	return forensicsFixtureData{
		handler:             newHandlerWithAuth(store, "memory", authConfig),
		store:               store,
		historicalTimestamp: historicalTimestamp,
		currentTimestamp:    currentTimestamp,
		flashbackTimestamp:  flashbackTimestamp,
	}
}

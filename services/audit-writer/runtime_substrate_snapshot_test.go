package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	attestationruntime "github.com/denisgrosek/changelock/internal/attestation"
	"github.com/denisgrosek/changelock/internal/audit"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

func TestRuntimeSnapshotIgnoresForeignSubjectCreationAndLimitCrowding(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	postRuntimeSubstrateObservation(t, handler, runtimeSubstrateValBExpectedExecEvent())
	postRuntimeSubstrateObservation(t, handler, runtimeSubstrateValBLowRiskExecEvent())

	for i := 0; i < 5; i++ {
		seedRuntimeSnapshotObservationRecord(t, store, "foreign-controller", runtimeSnapshotGhostEvent(fmt.Sprintf("ghost-%d", i)))
	}

	srv := server{store: store, requestTimeout: time.Second}
	snapshot, err := srv.buildRuntimeSnapshot(context.Background(), runtimeIntegrityFilter{
		event:       audit.EventFilter{TenantID: "acme", Environment: "prod", Limit: 200},
		TenantID:    "acme",
		Environment: "prod",
		Limit:       1,
	})
	if err != nil {
		t.Fatalf("build snapshot: %v", err)
	}

	if len(snapshot.subjects) != 2 {
		t.Fatalf("expected only runtime subjects to exist, got %#v", snapshot.subjects)
	}
	items := snapshot.sortedSubjects()
	if len(items) != 1 {
		t.Fatalf("expected sortedSubjects limit=1, got %#v", items)
	}
	if items[0].Workload != "api" && items[0].Workload != "worker" {
		t.Fatalf("expected runtime workload under limit, got %#v", items[0])
	}
}

func TestRuntimeSnapshotBlocksForeignMutationAndSeparatesProvenanceWhitelist(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	postRuntimeSubstrateObservation(t, handler, runtimeSubstrateValBExpectedExecEvent())

	foreign := runtimeSubstrateValBExpectedExecEvent()
	foreign.Workload.Workload = "api"
	foreign.Workload.PolicySubject = runtimeSubjectRef("cluster-a", "acme-prod", "Deployment", "api")
	foreign.Process.ProcessName = "foreign-mutator"
	foreign.Process.BinaryDigest = "sha256:foreign-runtime"
	seedRuntimeSnapshotObservationRecord(t, store, "foreign-controller", foreign)

	seedRuntimeSnapshotArtifactEvidence(t, store, "foreign-verifier", "cluster-a", "acme", "prod", "acme-prod", "Deployment", "api", "sha256:foreign", "foreign-signer", "sha256:foreign", "https://example.invalid/foreign", "predicate/foreign")
	seedRuntimeSnapshotArtifactEvidence(t, store, "attestation-verifier", "cluster-a", "acme", "prod", "acme-prod", "Deployment", "api", "sha256:111", "trusted-signer", "sha256:111", "https://github.com/example/api", "https://slsa.dev/provenance/v1")
	seedRuntimeSubstrateValBAttestation(t, store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "api", runtimeSnapshotVerifiedAttestation("api", "sha256:111"))

	srv := server{store: store, requestTimeout: time.Second}
	snapshot, err := srv.buildRuntimeSnapshot(context.Background(), runtimeIntegrityFilter{
		event:       audit.EventFilter{TenantID: "acme", Environment: "prod", Limit: 200},
		TenantID:    "acme",
		Environment: "prod",
		Limit:       10,
	})
	if err != nil {
		t.Fatalf("build snapshot: %v", err)
	}

	subject := snapshot.snapshotSubject(runtimeSubjectRef("cluster-a", "acme-prod", "Deployment", "api"))
	if subject == nil {
		t.Fatal("expected api subject")
	}
	if len(subject.Observations) != 1 {
		t.Fatalf("expected foreign runtime record not to append observations, got %#v", subject.Observations)
	}
	if containsString(subject.ExpectedSigners, "foreign-signer") {
		t.Fatalf("expected foreign signer to be ignored, got %#v", subject.ExpectedSigners)
	}
	if !containsString(subject.ExpectedSigners, "trusted-signer") {
		t.Fatalf("expected whitelisted signer to enrich subject, got %#v", subject.ExpectedSigners)
	}
	if containsString(subject.ArtifactDigests, "sha256:foreign") {
		t.Fatalf("expected foreign artifact digest to be ignored, got %#v", subject.ArtifactDigests)
	}
	if !containsString(subject.ArtifactDigests, "sha256:111") {
		t.Fatalf("expected whitelisted artifact digest to be present, got %#v", subject.ArtifactDigests)
	}
	if subject.LatestAttestation == nil || subject.LatestAttestation.CurrentState == "" {
		t.Fatalf("expected whitelisted phase2 attestation to enrich subject, got %#v", subject.LatestAttestation)
	}
}

func TestRuntimeSnapshotRetainsRuntimeAuthoritativeProvenanceInputs(t *testing.T) {
	store := audit.NewMemoryStore()
	now := time.Date(2026, 4, 22, 12, 0, 0, 0, time.UTC)
	for _, event := range []audit.Event{
		{
			RequestID:                "runtime-api-desired",
			Timestamp:                now.Add(-20 * time.Minute),
			Component:                "runtime-agent",
			EventType:                audit.EventTypeRuntimeDesiredStateRecorded,
			Decision:                 audit.DecisionAllow,
			ClusterID:                "cluster-a",
			TenantID:                 "acme",
			Environment:              "prod",
			Namespace:                "acme-prod",
			WorkloadKind:             "Deployment",
			Workload:                 "api",
			Repo:                     "acme/platform-api",
			Digest:                   "sha256:api-v1",
			DesiredStateSourceRef:    "deploy:api:v1",
			DesiredStateApprovalID:   "approval-api-v1",
			DesiredStateVerification: "verified",
			Evidence: &audit.Evidence{
				Artifact: &audit.ArtifactEvidence{
					SignerIdentity:           "runtime-signer",
					AttestationSubjectDigest: "sha256:api-v1",
					AttestationPredicate:     "https://slsa.dev/provenance/v1",
					SBOMHash:                 "sbom-api-v1",
				},
			},
		},
		{
			RequestID:    "runtime-api-observation",
			Timestamp:    now.Add(-10 * time.Minute),
			Component:    "runtime-agent",
			EventType:    audit.EventTypeRuntimeObservationRecorded,
			Decision:     audit.DecisionAllow,
			ClusterID:    "cluster-a",
			TenantID:     "acme",
			Environment:  "prod",
			Namespace:    "acme-prod",
			WorkloadKind: "Deployment",
			Workload:     "api",
			RuntimeIntegrity: canonicalJSONMust(runtimeIntegrityEventPayload{
				Observation: &runtimeObservationPayload{
					Node:         "node-a",
					Pod:          "api-0",
					EventType:    "exec_lifecycle",
					EventPayload: runtimeSubstrateValAEventPayload(runtimeSubstrateValBExpectedExecEvent()),
					Confidence:   runtimeConfidenceHigh,
				},
			}),
		},
	} {
		if _, err := store.Ingest(context.Background(), event); err != nil {
			t.Fatalf("seed runtime authoritative record: %v", err)
		}
	}

	srv := server{store: store, requestTimeout: time.Second}
	snapshot, err := srv.buildRuntimeSnapshot(context.Background(), runtimeIntegrityFilter{
		event:       audit.EventFilter{TenantID: "acme", Environment: "prod", Limit: 100},
		TenantID:    "acme",
		Environment: "prod",
		Limit:       10,
	})
	if err != nil {
		t.Fatalf("build snapshot: %v", err)
	}

	subject := snapshot.snapshotSubject(runtimeSubjectRef("cluster-a", "acme-prod", "Deployment", "api"))
	if subject == nil {
		t.Fatal("expected runtime-authoritative subject")
	}
	if !containsString(subject.ExpectedSigners, "runtime-signer") {
		t.Fatalf("expected runtime-agent signer to be retained, got %#v", subject.ExpectedSigners)
	}
	if !containsString(subject.AttestationSubjectDigests, "sha256:api-v1") {
		t.Fatalf("expected runtime-agent attestation subject digest to be retained, got %#v", subject.AttestationSubjectDigests)
	}
	if !containsString(mapKeys(subject.TrustInputs), "attestation_provenance") {
		t.Fatalf("expected runtime-agent attestation provenance trust input, got %#v", mapKeys(subject.TrustInputs))
	}
}

func TestRuntimeSubstrateValAAndValBStayActiveWithForeignSnapshotRecordsPresent(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	for _, event := range []runtimesubstrate.RuntimeSubstrateObservedEvent{
		runtimeSubstrateValBExpectedExecEvent(),
		runtimeSubstrateValBLowRiskExecEvent(),
		runtimeSubstrateValBHardMismatchExecEvent(),
		runtimeSubstrateProcessStaleEvent(),
		runtimeSubstrateFilePartialEvent(),
		runtimeSubstrateNetworkUnsupportedEvent(),
	} {
		postRuntimeSubstrateObservation(t, handler, event)
	}

	seedRuntimeSubstrateValBArtifactEvidence(t, store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "api", "sha256:111", "https://github.com/example/api/.github/workflows/release.yml@refs/heads/main", "sha256:111", "https://github.com/example/api", "https://slsa.dev/provenance/v1")
	seedRuntimeSubstrateValBArtifactEvidence(t, store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "worker", "", "https://github.com/example/worker/.github/workflows/release.yml@refs/heads/main", "", "https://github.com/example/worker", "")
	seedRuntimeSubstrateValBArtifactEvidence(t, store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "rogue", "sha256:trusted", "https://github.com/example/rogue/.github/workflows/release.yml@refs/heads/main", "sha256:trusted", "https://github.com/example/rogue", "https://slsa.dev/provenance/v1")

	seedRuntimeSubstrateValBAttestation(t, store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "api", runtimeSnapshotVerifiedAttestation("api", "sha256:111"))
	seedRuntimeSubstrateValBAttestation(t, store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "worker", runtimeSnapshotDegradedAttestation("worker", "sha256:worker"))
	seedRuntimeSubstrateValBAttestation(t, store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "rogue", runtimeSnapshotMismatchAttestation("rogue", "sha256:rogue"))

	for i := 0; i < 8; i++ {
		seedRuntimeSnapshotObservationRecord(t, store, "foreign-controller", runtimeSnapshotGhostEvent(fmt.Sprintf("foreign-%d", i)))
		seedRuntimeSnapshotArtifactEvidence(t, store, "foreign-verifier", "cluster-a", "acme", "prod", "acme-prod", "Deployment", fmt.Sprintf("foreign-%d", i), fmt.Sprintf("sha256:foreign-%d", i), "foreign-signer", fmt.Sprintf("sha256:foreign-%d", i), "https://example.invalid/foreign", "predicate/foreign")
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vala/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vala proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var valA runtimeSubstrateValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&valA); err != nil {
		t.Fatalf("decode vala proofs: %v", err)
	}
	if valA.CurrentState != runtimesubstrate.RuntimeSubstrateValAStateActive {
		t.Fatalf("expected active vala proofs with foreign records ignored, got %#v", valA)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valb/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valb proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var valB runtimeSubstrateValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&valB); err != nil {
		t.Fatalf("decode valb proofs: %v", err)
	}
	if valB.CurrentState != runtimesubstrate.RuntimeSubstrateValBStateActive {
		t.Fatalf("expected active valb proofs with foreign records ignored, got %#v", valB)
	}
	for _, item := range valB.ProcessImageItems {
		if item.Workload.Workload == "foreign-0" {
			t.Fatalf("expected foreign workload to be excluded from valb process image items, got %#v", valB.ProcessImageItems)
		}
	}
}

func seedRuntimeSnapshotObservationRecord(t *testing.T, store audit.Store, component string, event runtimesubstrate.RuntimeSubstrateObservedEvent) {
	t.Helper()
	payload, err := canonicalJSON(runtimeIntegrityEventPayload{
		Observation: &runtimeObservationPayload{
			Node:         event.Node.NodeID,
			Pod:          event.Workload.PodUID,
			EventType:    event.EventFamily,
			EventPayload: runtimeSubstrateValAEventPayload(event),
			Confidence:   runtimeSubstrateConfidenceToRuntimeConfidence(event.AttributionConfidence),
		},
	})
	if err != nil {
		t.Fatalf("canonical observation payload: %v", err)
	}
	_, err = store.Ingest(context.Background(), audit.Event{
		RequestID:        audit.NewRequestID(),
		Timestamp:        event.ObservedAt.UTC(),
		Component:        component,
		EventType:        audit.EventTypeRuntimeObservationRecorded,
		Decision:         audit.DecisionAllow,
		ClusterID:        event.Workload.ClusterID,
		TenantID:         "acme",
		Environment:      "prod",
		Namespace:        event.Workload.Namespace,
		WorkloadKind:     event.Workload.WorkloadKind,
		Workload:         event.Workload.Workload,
		RuntimeIntegrity: payload,
	})
	if err != nil {
		t.Fatalf("seed observation record: %v", err)
	}
}

func seedRuntimeSnapshotArtifactEvidence(t *testing.T, store audit.Store, component, clusterID, tenantID, environment, namespace, workloadKind, workload, digest, signer, attestationSubjectDigest, repo, predicate string) {
	t.Helper()
	_, err := store.Ingest(context.Background(), audit.Event{
		RequestID:    audit.NewRequestID(),
		Timestamp:    time.Date(2026, 4, 22, 10, 5, 0, 0, time.UTC),
		Component:    component,
		EventType:    audit.EventTypeArtifactVerificationResult,
		ClusterID:    clusterID,
		TenantID:     tenantID,
		Environment:  environment,
		Namespace:    namespace,
		WorkloadKind: workloadKind,
		Workload:     workload,
		Repo:         repo,
		Digest:       digest,
		Decision:     audit.DecisionAllow,
		Evidence: &audit.Evidence{
			Artifact: &audit.ArtifactEvidence{
				Digest:                   digest,
				SignerIdentity:           signer,
				AttestationSubjectDigest: attestationSubjectDigest,
				Repository:               repo,
				AttestationPredicate:     predicate,
			},
		},
	})
	if err != nil {
		t.Fatalf("seed artifact evidence: %v", err)
	}
}

func runtimeSnapshotGhostEvent(workload string) runtimesubstrate.RuntimeSubstrateObservedEvent {
	event := runtimeSubstrateValBExpectedExecEvent()
	event.Workload.Workload = workload
	event.Workload.PolicySubject = runtimeSubjectRef(event.Workload.ClusterID, event.Workload.Namespace, event.Workload.WorkloadKind, workload)
	event.Process.ProcessName = "ghost-" + workload
	event.Process.BinaryDigest = "sha256:" + workload
	return event
}

func runtimeSnapshotVerifiedAttestation(workload, measurement string) attestationruntime.VerificationResult {
	return attestationruntime.VerificationResult{
		SchemaVersion: attestationruntime.SchemaVersion,
		SubjectRef:    runtimeSubjectRef("cluster-a", "acme-prod", "Deployment", workload),
		Provider:      "tdx",
		QuoteType:     "tdx_quote",
		Measurement:   measurement,
		CurrentState:  attestationruntime.VerdictVerified,
		VerifiedAt:    time.Date(2026, 4, 22, 9, 2, 0, 0, time.UTC),
	}
}

func runtimeSnapshotDegradedAttestation(workload, measurement string) attestationruntime.VerificationResult {
	return attestationruntime.VerificationResult{
		SchemaVersion: attestationruntime.SchemaVersion,
		SubjectRef:    runtimeSubjectRef("cluster-a", "acme-prod", "Deployment", workload),
		Provider:      "tdx",
		QuoteType:     "tdx_quote",
		Measurement:   measurement,
		CurrentState:  attestationruntime.VerdictDegraded,
		VerifiedAt:    time.Date(2026, 4, 22, 9, 3, 0, 0, time.UTC),
	}
}

func runtimeSnapshotMismatchAttestation(workload, measurement string) attestationruntime.VerificationResult {
	return attestationruntime.VerificationResult{
		SchemaVersion: attestationruntime.SchemaVersion,
		SubjectRef:    runtimeSubjectRef("cluster-a", "acme-prod", "Deployment", workload),
		Provider:      "tdx",
		QuoteType:     "tdx_quote",
		Measurement:   measurement,
		CurrentState:  attestationruntime.VerdictMismatch,
		VerifiedAt:    time.Date(2026, 4, 22, 9, 4, 0, 0, time.UTC),
	}
}

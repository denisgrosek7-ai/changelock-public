package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	attestationruntime "github.com/denisgrosek/changelock/internal/attestation"
	"github.com/denisgrosek/changelock/internal/audit"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

func TestRuntimeSubstrateValBCorrelationModelHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valb/correlation-model?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected correlation model 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response runtimeSubstrateValBCorrelationModelResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode correlation model: %v", err)
	}
	if response.CurrentState != runtimesubstrate.RuntimeSubstrateValBCorrelationModelStateActive {
		t.Fatalf("expected active correlation model, got %#v", response)
	}
	if !containsString(response.Model.DriftClasses, runtimesubstrate.RuntimeSubstrateDriftHardMismatch) {
		t.Fatalf("expected hard mismatch drift class, got %#v", response.Model.DriftClasses)
	}
}

func TestRuntimeSubstrateValBHandlersAndProofs(t *testing.T) {
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

	seedRuntimeSubstrateValBAttestation(t, store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "api", attestationruntime.VerificationResult{
		SchemaVersion: attestationruntime.SchemaVersion,
		SubjectRef:    runtimeSubjectRef("cluster-a", "acme-prod", "Deployment", "api"),
		Provider:      "tdx",
		QuoteType:     "tdx_quote",
		Measurement:   "sha256:111",
		CurrentState:  attestationruntime.VerdictVerified,
		VerifiedAt:    time.Date(2026, 4, 22, 9, 2, 0, 0, time.UTC),
	})
	seedRuntimeSubstrateValBAttestation(t, store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "worker", attestationruntime.VerificationResult{
		SchemaVersion: attestationruntime.SchemaVersion,
		SubjectRef:    runtimeSubjectRef("cluster-a", "acme-prod", "Deployment", "worker"),
		Provider:      "tdx",
		QuoteType:     "tdx_quote",
		Measurement:   "sha256:worker",
		CurrentState:  attestationruntime.VerdictDegraded,
		VerifiedAt:    time.Date(2026, 4, 22, 9, 3, 0, 0, time.UTC),
	})
	seedRuntimeSubstrateValBAttestation(t, store, "cluster-a", "acme", "prod", "acme-prod", "Deployment", "rogue", attestationruntime.VerificationResult{
		SchemaVersion: attestationruntime.SchemaVersion,
		SubjectRef:    runtimeSubjectRef("cluster-a", "acme-prod", "Deployment", "rogue"),
		Provider:      "tdx",
		QuoteType:     "tdx_quote",
		Measurement:   "sha256:rogue",
		CurrentState:  attestationruntime.VerdictMismatch,
		VerifiedAt:    time.Date(2026, 4, 22, 9, 4, 0, 0, time.UTC),
	})

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valb/process-image-linkage?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected process-image linkage 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var process runtimeSubstrateValBProcessImageResponse
	if err := json.NewDecoder(rec.Body).Decode(&process); err != nil {
		t.Fatalf("decode process-image linkage: %v", err)
	}
	if process.CurrentState != runtimesubstrate.RuntimeSubstrateValBProcessImageStateActive {
		t.Fatalf("expected active process-image state, got %#v", process)
	}
	if !hasProcessImageDrift(process.Items, "api", runtimesubstrate.RuntimeSubstrateDriftExpected) {
		t.Fatalf("expected api expected drift, got %#v", process.Items)
	}
	if !hasProcessImageDrift(process.Items, "worker", runtimesubstrate.RuntimeSubstrateDriftLowRisk) {
		t.Fatalf("expected worker low-risk drift, got %#v", process.Items)
	}
	if !hasProcessImageDrift(process.Items, "rogue", runtimesubstrate.RuntimeSubstrateDriftHardMismatch) {
		t.Fatalf("expected rogue hard mismatch drift, got %#v", process.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valb/provenance-linkage?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected provenance linkage 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var provenance runtimeSubstrateValBProvenanceResponse
	if err := json.NewDecoder(rec.Body).Decode(&provenance); err != nil {
		t.Fatalf("decode provenance linkage: %v", err)
	}
	if provenance.CurrentState != runtimesubstrate.RuntimeSubstrateValBProvenanceStateActive {
		t.Fatalf("expected active provenance state, got %#v", provenance)
	}
	if !hasProvenanceDrift(provenance.Items, "api", runtimesubstrate.RuntimeSubstrateDriftExpected) {
		t.Fatalf("expected api provenance expected drift, got %#v", provenance.Items)
	}
	if !hasProvenanceDrift(provenance.Items, "worker", runtimesubstrate.RuntimeSubstrateDriftLowRisk) {
		t.Fatalf("expected worker provenance low-risk drift, got %#v", provenance.Items)
	}
	if !hasProvenanceDrift(provenance.Items, "rogue", runtimesubstrate.RuntimeSubstrateDriftHardMismatch) {
		t.Fatalf("expected rogue provenance hard mismatch drift, got %#v", provenance.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valb/drift-catalog?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected drift catalog 200, got %d: %s", rec.Body.Len(), rec.Body.String())
	}

	var drift runtimeSubstrateValBDriftCatalogResponse
	if err := json.NewDecoder(rec.Body).Decode(&drift); err != nil {
		t.Fatalf("decode drift catalog: %v", err)
	}
	if drift.CurrentState != runtimesubstrate.RuntimeSubstrateValBDriftCatalogStateActive {
		t.Fatalf("expected active drift catalog, got %#v", drift)
	}
	if !hasDriftSummary(drift.Items, "rogue", runtimesubstrate.RuntimeSubstrateDriftHardMismatch) {
		t.Fatalf("expected rogue hard mismatch in drift catalog, got %#v", drift.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valb/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valb proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs runtimeSubstrateValBProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode valb proofs: %v", err)
	}
	if proofs.CurrentState != runtimesubstrate.RuntimeSubstrateValBStateActive {
		t.Fatalf("expected active valb proofs, got %#v", proofs)
	}
	if proofs.ValAState != runtimesubstrate.RuntimeSubstrateValAStateActive {
		t.Fatalf("expected active val a dependency, got %#v", proofs)
	}
}

func postRuntimeSubstrateObservation(t *testing.T, handler http.Handler, event runtimesubstrate.RuntimeSubstrateObservedEvent) {
	t.Helper()
	body, err := json.Marshal(runtimeSubstrateValAObservationWriteRequest{Event: event})
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/v1/runtime/substrate-depth/vala/observability?tenant_id=acme&environment=prod", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer operator-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected observability ingest 201, got %d: %s", rec.Code, rec.Body.String())
	}
}

func seedRuntimeSubstrateValBArtifactEvidence(t *testing.T, store audit.Store, clusterID, tenantID, environment, namespace, workloadKind, workload, digest, signer, attestationSubjectDigest, repo, predicate string) {
	t.Helper()
	_, err := store.Ingest(context.Background(), audit.Event{
		RequestID:    audit.NewRequestID(),
		Timestamp:    time.Date(2026, 4, 22, 9, 5, 0, 0, time.UTC),
		Component:    "attestation-verifier",
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

func seedRuntimeSubstrateValBAttestation(t *testing.T, store audit.Store, clusterID, tenantID, environment, namespace, workloadKind, workload string, result attestationruntime.VerificationResult) {
	t.Helper()
	payload, err := canonicalJSON(runtimePhase2EventPayload{Attestation: &result})
	if err != nil {
		t.Fatalf("canonical attestation payload: %v", err)
	}
	_, err = store.Ingest(context.Background(), audit.Event{
		RequestID:        audit.NewRequestID(),
		Timestamp:        result.VerifiedAt,
		Component:        runtimePhase2Component,
		EventType:        audit.EventTypeRuntimeAttestationVerified,
		ClusterID:        clusterID,
		TenantID:         tenantID,
		Environment:      environment,
		Namespace:        namespace,
		WorkloadKind:     workloadKind,
		Workload:         workload,
		Decision:         audit.DecisionAllow,
		RuntimeIntegrity: payload,
	})
	if err != nil {
		t.Fatalf("seed phase2 attestation: %v", err)
	}
}

func hasProcessImageDrift(items []runtimesubstrate.RuntimeSubstrateProcessImageLinkage, workload, drift string) bool {
	for _, item := range items {
		if item.Workload.Workload == workload && item.DriftClass == drift {
			return true
		}
	}
	return false
}

func hasProvenanceDrift(items []runtimesubstrate.RuntimeSubstrateProvenanceLinkage, workload, drift string) bool {
	for _, item := range items {
		_, _, _, actualWorkload, err := parseRuntimeSubjectRef(item.SubjectRef)
		if err == nil && actualWorkload == workload && item.DriftClass == drift {
			return true
		}
	}
	return false
}

func hasDriftSummary(items []runtimesubstrate.RuntimeSubstrateDriftRecord, workload, drift string) bool {
	for _, item := range items {
		_, _, _, actualWorkload, err := parseRuntimeSubjectRef(item.SubjectRef)
		if err == nil && actualWorkload == workload && item.DriftClass == drift {
			return true
		}
	}
	return false
}

func runtimeSubstrateValBExpectedExecEvent() runtimesubstrate.RuntimeSubstrateObservedEvent {
	event := runtimeSubstrateExecObservedEvent()
	event.Process.BinaryDigest = "sha256:111"
	event.Workload.ImageDigest = "sha256:111"
	event.Workload.Workload = "api"
	event.Workload.PodUID = "pod-api-0"
	event.Reasons = append(event.Reasons, "binary_digest_captured")
	return event
}

func runtimeSubstrateValBLowRiskExecEvent() runtimesubstrate.RuntimeSubstrateObservedEvent {
	return runtimesubstrate.RuntimeSubstrateObservedEvent{
		EventFamily:           runtimesubstrate.RuntimeSubstrateEventFamilyExecLifecycle,
		CurrentState:          runtimesubstrate.RuntimeSubstrateEventStateObserved,
		Process:               runtimesubstrate.ProcessIdentity{ProcessName: "worker", ProcessPath: "/usr/local/bin/worker", PID: 2210, CgroupID: "cg-worker", NamespaceID: "ns-acme-prod", LineageRef: "init->containerd-shim->worker"},
		ParentPID:             1024,
		ThreadID:              2210,
		ContainerRuntime:      "containerd",
		NamespaceMode:         "kubernetes_pod_namespace",
		Workload:              runtimesubstrate.WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "worker", PodUID: "pod-worker-0", ImageDigest: "sha256:worker-image"},
		Node:                  runtimesubstrate.NodeIdentity{NodeID: "node-h1", SubstrateClass: runtimesubstrate.SubstrateClassHardened, TrustBoundary: runtimesubstrate.TrustBoundaryKernelRuntimeLayer},
		AttributionConfidence: runtimesubstrate.RuntimeSubstrateConfidenceHighFidelity,
		FreshnessState:        runtimesubstrate.RuntimeSubstrateFreshnessFresh,
		Reasons:               []string{"exec_event_captured", "signer_linkage_present_without_digest"},
		ObservedAt:            time.Date(2026, 4, 22, 9, 1, 0, 0, time.UTC),
	}
}

func runtimeSubstrateValBHardMismatchExecEvent() runtimesubstrate.RuntimeSubstrateObservedEvent {
	return runtimesubstrate.RuntimeSubstrateObservedEvent{
		EventFamily:           runtimesubstrate.RuntimeSubstrateEventFamilyExecLifecycle,
		CurrentState:          runtimesubstrate.RuntimeSubstrateEventStateObserved,
		Process:               runtimesubstrate.ProcessIdentity{ProcessName: "rogue", ProcessPath: "/usr/local/bin/rogue", BinaryDigest: "sha256:rogue-bin", PID: 3310, CgroupID: "cg-rogue", NamespaceID: "ns-acme-prod", LineageRef: "init->containerd-shim->rogue"},
		ParentPID:             1024,
		ThreadID:              3310,
		ContainerRuntime:      "containerd",
		NamespaceMode:         "kubernetes_pod_namespace",
		Workload:              runtimesubstrate.WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "rogue", PodUID: "pod-rogue-0", ImageDigest: "sha256:rogue-image"},
		Node:                  runtimesubstrate.NodeIdentity{NodeID: "node-r1", SubstrateClass: runtimesubstrate.SubstrateClassStandard, TrustBoundary: runtimesubstrate.TrustBoundaryKernelRuntimeLayer},
		AttributionConfidence: runtimesubstrate.RuntimeSubstrateConfidenceHighFidelity,
		FreshnessState:        runtimesubstrate.RuntimeSubstrateFreshnessFresh,
		Reasons:               []string{"exec_event_captured", "binary_digest_captured"},
		ObservedAt:            time.Date(2026, 4, 22, 9, 2, 0, 0, time.UTC),
	}
}

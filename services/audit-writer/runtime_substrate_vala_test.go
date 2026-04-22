package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

func TestRuntimeSubstrateValAEntryGateAndEventSchemaHandlers(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/entry-gate?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected entry gate 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var entry runtimeSubstrateEntryGateResponse
	if err := json.NewDecoder(rec.Body).Decode(&entry); err != nil {
		t.Fatalf("decode entry gate: %v", err)
	}
	if entry.CurrentState != runtimesubstrate.RuntimeSubstrateEntryGateStateReady {
		t.Fatalf("expected ready entry gate, got %#v", entry)
	}
	if entry.EntryGate.ImplementationPermission != runtimesubstrate.RuntimeSubstrateImplementationStateRuntimePath {
		t.Fatalf("expected runtime baseline entry permission, got %#v", entry.EntryGate)
	}
	if !containsString(entry.EntryGate.ExplicitExclusions, "absolute_truth_rhetoric") {
		t.Fatalf("expected explicit exclusion, got %#v", entry.EntryGate.ExplicitExclusions)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vala/event-schema?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected event schema 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var schema runtimeSubstrateEventSchemaResponse
	if err := json.NewDecoder(rec.Body).Decode(&schema); err != nil {
		t.Fatalf("decode event schema: %v", err)
	}
	if schema.CurrentState != runtimesubstrate.RuntimeSubstrateValAEventSchemaStateActive {
		t.Fatalf("expected active event schema, got %#v", schema)
	}
	if !containsString(schema.EventSchema.CorrelationStates, runtimesubstrate.RuntimeSubstrateEventStateUnsupported) {
		t.Fatalf("expected unsupported state in schema, got %#v", schema.EventSchema.CorrelationStates)
	}
}

func TestRuntimeSubstrateValASupportMatrixAndEmptyObservability(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vala/support-matrix?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected support matrix 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var support runtimeSubstrateSupportMatrixResponse
	if err := json.NewDecoder(rec.Body).Decode(&support); err != nil {
		t.Fatalf("decode support matrix: %v", err)
	}
	if support.CurrentState != runtimesubstrate.RuntimeSubstrateValASupportMatrixStateActive {
		t.Fatalf("expected active support matrix, got %#v", support)
	}
	if support.ImplementationState != runtimesubstrate.RuntimeSubstrateImplementationStateRuntimePath {
		t.Fatalf("expected runtime baseline support matrix implementation state, got %#v", support)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vala/observability?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected observability 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var observability runtimeSubstrateObservabilityResponse
	if err := json.NewDecoder(rec.Body).Decode(&observability); err != nil {
		t.Fatalf("decode observability: %v", err)
	}
	if observability.CurrentState != runtimesubstrate.RuntimeSubstrateValAObservabilityStateIncomplete {
		t.Fatalf("expected incomplete empty observability, got %#v", observability)
	}
	if observability.ImplementationState != runtimesubstrate.RuntimeSubstrateImplementationStateRuntimePath || observability.RecordKind != runtimesubstrate.RuntimeSubstrateRecordKindObserved || observability.NotRuntimeTruthSurface {
		t.Fatalf("expected runtime-backed observability markers, got %#v", observability)
	}
	if len(observability.Items) != 0 {
		t.Fatalf("expected no observed items before ingest, got %#v", observability.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vala/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs runtimeSubstrateValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if proofs.CurrentState != runtimesubstrate.RuntimeSubstrateValAStateIncomplete {
		t.Fatalf("expected incomplete proofs without runtime observations, got %#v", proofs)
	}
	if proofs.ImplementationState != runtimesubstrate.RuntimeSubstrateImplementationStateRuntimePath {
		t.Fatalf("expected runtime baseline proofs implementation state, got %#v", proofs)
	}
}

func TestRuntimeSubstrateValAObservabilityIngestAndProofs(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	for _, event := range []runtimesubstrate.RuntimeSubstrateObservedEvent{
		runtimeSubstrateExecObservedEvent(),
		runtimeSubstrateProcessStaleEvent(),
		runtimeSubstrateFilePartialEvent(),
		runtimeSubstrateNetworkUnsupportedEvent(),
	} {
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

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vala/observability?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected observability 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var observability runtimeSubstrateObservabilityResponse
	if err := json.NewDecoder(rec.Body).Decode(&observability); err != nil {
		t.Fatalf("decode observability: %v", err)
	}
	if observability.CurrentState != runtimesubstrate.RuntimeSubstrateValAObservabilityStateActive {
		t.Fatalf("expected active observability after ingest, got %#v", observability)
	}
	if len(observability.Items) != 4 {
		t.Fatalf("expected 4 observed items, got %#v", observability.Items)
	}
	if observability.Items[0].Process.PID == 0 || observability.Items[0].Process.ProcessName == "" {
		t.Fatalf("expected stable process identity, got %#v", observability.Items[0])
	}
	if !containsString(observability.Items[0].Reasons, "exec_event_captured") {
		t.Fatalf("expected recorded reasons on exec event, got %#v", observability.Items[0])
	}
	if !containsString(observability.Items[2].UnsupportedFields, "thread_id") {
		t.Fatalf("expected partial file attribution unsupported fields, got %#v", observability.Items[2])
	}
	if observability.Items[3].CurrentState != runtimesubstrate.RuntimeSubstrateEventStateUnsupported {
		t.Fatalf("expected unsupported network item, got %#v", observability.Items[3])
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/vala/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs runtimeSubstrateValAProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if proofs.CurrentState != runtimesubstrate.RuntimeSubstrateValAStateActive {
		t.Fatalf("expected active Val A proofs, got %#v", proofs)
	}
	if len(proofs.ObservedEvents) != 4 {
		t.Fatalf("expected observed events in proofs, got %#v", proofs.ObservedEvents)
	}
}

func TestRuntimeSubstrateValARejectsCrossTenantObservationScope(t *testing.T) {
	store := audit.NewMemoryStore()
	cfg, signer := newOIDCHandlerConfig(t, true, true)
	handler := newHandlerWithAuth(store, "memory", cfg)

	body, err := json.Marshal(runtimeSubstrateValAObservationWriteRequest{
		Event: runtimesubstrate.RuntimeSubstrateObservedEvent{
			EventFamily:           runtimesubstrate.RuntimeSubstrateEventFamilyExecLifecycle,
			CurrentState:          runtimesubstrate.RuntimeSubstrateEventStateObserved,
			Process:               runtimesubstrate.ProcessIdentity{ProcessName: "api", PID: 2048},
			Workload:              runtimesubstrate.WorkloadIdentity{ClusterID: "cluster-a", Namespace: "globex-prod", WorkloadKind: "Deployment", Workload: "api"},
			Node:                  runtimesubstrate.NodeIdentity{NodeID: "node-a", SubstrateClass: runtimesubstrate.SubstrateClassStandard, TrustBoundary: runtimesubstrate.TrustBoundaryKernelRuntimeLayer},
			AttributionConfidence: runtimesubstrate.RuntimeSubstrateConfidenceHighFidelity,
			FreshnessState:        runtimesubstrate.RuntimeSubstrateFreshnessFresh,
			ObservedAt:            time.Date(2026, 4, 22, 10, 0, 0, 0, time.UTC),
		},
	})
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/v1/runtime/substrate-depth/vala/observability?tenant_id=acme&environment=prod", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+signer.token(t, map[string]any{
		"sub":       "operator@example.com",
		"groups":    []string{"changelock-operators"},
		"tenant_id": "acme",
	}))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403 for cross-tenant substrate observation, got %d: %s", rec.Code, rec.Body.String())
	}
}

func runtimeSubstrateExecObservedEvent() runtimesubstrate.RuntimeSubstrateObservedEvent {
	return runtimesubstrate.RuntimeSubstrateObservedEvent{
		EventFamily:           runtimesubstrate.RuntimeSubstrateEventFamilyExecLifecycle,
		CurrentState:          runtimesubstrate.RuntimeSubstrateEventStateObserved,
		Process:               runtimesubstrate.ProcessIdentity{ProcessName: "api", ProcessPath: "/usr/local/bin/api", PID: 2048, CgroupID: "cg-api", NamespaceID: "ns-acme-prod", LineageRef: "init->containerd-shim->api"},
		ParentPID:             1024,
		ThreadID:              2048,
		ContainerRuntime:      "containerd",
		NamespaceMode:         "kubernetes_pod_namespace",
		Workload:              runtimesubstrate.WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "api", PodUID: "pod-api-0", ImageDigest: "sha256:111"},
		Node:                  runtimesubstrate.NodeIdentity{NodeID: "node-a", SubstrateClass: runtimesubstrate.SubstrateClassStandard, TrustBoundary: runtimesubstrate.TrustBoundaryKernelRuntimeLayer},
		AttributionConfidence: runtimesubstrate.RuntimeSubstrateConfidenceHighFidelity,
		FreshnessState:        runtimesubstrate.RuntimeSubstrateFreshnessFresh,
		Reasons:               []string{"exec_event_captured", "process_and_workload_identity_bound"},
		ObservedAt:            time.Date(2026, 4, 22, 9, 0, 0, 0, time.UTC),
	}
}

func runtimeSubstrateProcessStaleEvent() runtimesubstrate.RuntimeSubstrateObservedEvent {
	return runtimesubstrate.RuntimeSubstrateObservedEvent{
		EventFamily:           runtimesubstrate.RuntimeSubstrateEventFamilyProcessLineage,
		CurrentState:          runtimesubstrate.RuntimeSubstrateEventStateStale,
		Process:               runtimesubstrate.ProcessIdentity{ProcessName: "worker", ProcessPath: "/usr/local/bin/worker", PID: 2210, CgroupID: "cg-worker", NamespaceID: "ns-acme-prod", LineageRef: "init->containerd-shim->worker"},
		ParentPID:             1024,
		ThreadID:              2211,
		ContainerRuntime:      "containerd",
		NamespaceMode:         "kubernetes_pod_namespace",
		Workload:              runtimesubstrate.WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "worker", PodUID: "pod-worker-0", ImageDigest: "sha256:222"},
		Node:                  runtimesubstrate.NodeIdentity{NodeID: "node-h1", SubstrateClass: runtimesubstrate.SubstrateClassHardened, TrustBoundary: runtimesubstrate.TrustBoundaryKernelRuntimeLayer},
		AttributionConfidence: runtimesubstrate.RuntimeSubstrateConfidenceLimitedContext,
		FreshnessState:        runtimesubstrate.RuntimeSubstrateFreshnessStale,
		Reasons:               []string{"captured_process_lineage_visible", "freshness_window_expired"},
		ObservedAt:            time.Date(2026, 4, 22, 8, 50, 0, 0, time.UTC),
	}
}

func runtimeSubstrateFilePartialEvent() runtimesubstrate.RuntimeSubstrateObservedEvent {
	return runtimesubstrate.RuntimeSubstrateObservedEvent{
		EventFamily:           runtimesubstrate.RuntimeSubstrateEventFamilyFileActivity,
		CurrentState:          runtimesubstrate.RuntimeSubstrateEventStatePartiallyCorrelated,
		Process:               runtimesubstrate.ProcessIdentity{ProcessName: "batch", ProcessPath: "/usr/local/bin/batch", PID: 3010, CgroupID: "vm-cg-batch", NamespaceID: "vm-namespace", LineageRef: "systemd->batch"},
		ParentPID:             1,
		ContainerRuntime:      "vm_guest_agent",
		NamespaceMode:         "vm_guest_namespace",
		Workload:              runtimesubstrate.WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "batch", PodUID: "pod-batch-0", ImageDigest: "sha256:333"},
		Node:                  runtimesubstrate.NodeIdentity{NodeID: "vm-node-a", SubstrateClass: runtimesubstrate.SubstrateClassStandard, TrustBoundary: runtimesubstrate.TrustBoundaryKernelRuntimeLayer},
		AttributionConfidence: runtimesubstrate.RuntimeSubstrateConfidenceBoundedCorrelation,
		FreshnessState:        runtimesubstrate.RuntimeSubstrateFreshnessFresh,
		UnsupportedFields:     []string{"thread_id", "host_hypervisor_ref"},
		Reasons:               []string{"guest_file_signal_captured", "host_level_hypervisor_context_not_in_scope"},
		ObservedAt:            time.Date(2026, 4, 22, 8, 40, 0, 0, time.UTC),
	}
}

func runtimeSubstrateNetworkUnsupportedEvent() runtimesubstrate.RuntimeSubstrateObservedEvent {
	return runtimesubstrate.RuntimeSubstrateObservedEvent{
		EventFamily:           runtimesubstrate.RuntimeSubstrateEventFamilyNetworkActivity,
		CurrentState:          runtimesubstrate.RuntimeSubstrateEventStateUnsupported,
		Process:               runtimesubstrate.ProcessIdentity{ProcessName: "sync-agent", ProcessPath: "/usr/local/bin/sync-agent", PID: 5100, CgroupID: "cg-sync", NamespaceID: "ns-offline", LineageRef: "systemd->sync-agent"},
		ParentPID:             1,
		ContainerRuntime:      "systemd_service",
		NamespaceMode:         "host_namespace",
		Workload:              runtimesubstrate.WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "sync-agent", PodUID: "pod-sync-0", ImageDigest: "sha256:666"},
		Node:                  runtimesubstrate.NodeIdentity{NodeID: "node-offline-a", SubstrateClass: runtimesubstrate.SubstrateClassHardened, TrustBoundary: runtimesubstrate.TrustBoundaryKernelRuntimeLayer},
		AttributionConfidence: runtimesubstrate.RuntimeSubstrateConfidenceUnsupportedSignal,
		FreshnessState:        runtimesubstrate.RuntimeSubstrateFreshnessUnavailable,
		UnsupportedFields:     []string{"destination_ip", "socket_inode", "remote_cluster_context"},
		Reasons:               []string{"airgapped_network_capture_path_unavailable", "unsupported_is_explicit_not_negative_evidence"},
		ObservedAt:            time.Date(2026, 4, 22, 8, 30, 0, 0, time.UTC),
	}
}

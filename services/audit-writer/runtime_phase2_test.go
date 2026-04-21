package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	runtimesubstrate "github.com/denisgrosek/changelock/internal/runtime"
)

func TestRuntimePhase2SubstrateTruthAndProofs(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandlerWithRuntimesAndSigning(
		store,
		"memory",
		mustStaticAuthConfig(t),
		nil,
		newSyncRuntime(syncConfig{Mode: audit.SyncModeDisabled}),
		newTestSoftwareSigningRuntime(t, "phase2-secret"),
	)

	substrateReq := httptest.NewRequest(http.MethodPost, "/v1/runtime/substrate-truth?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
	  "truth":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "workload":{"cluster_id":"cluster-a","namespace":"acme-prod","workload_kind":"Deployment","workload":"api","image_digest":"sha256:abc"},
	    "process":{"process_name":"api","cgroup_id":"cg-1"},
	    "node":{"node_id":"node-a","substrate_class":"confidential","trust_boundary":"attestation_provider_layer"},
	    "attestation":{"provider":"sgx","quote_type":"sgx_quote","measurement":"m-1","lifecycle_state":"active","observed_state":"verified","credential_release_state":"released"}
	  },
	  "profile_id":"confidential-strict"
	}`))
	substrateReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	substrateReq.Header.Set("Content-Type", "application/json")
	substrateRec := httptest.NewRecorder()
	handler.ServeHTTP(substrateRec, substrateReq)
	if substrateRec.Code != http.StatusCreated {
		t.Fatalf("expected substrate truth 201, got %d: %s", substrateRec.Code, substrateRec.Body.String())
	}

	attestationReq := httptest.NewRequest(http.MethodPost, "/v1/runtime/attestation/verify?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
	  "subject_ref":"cluster-a/acme-prod/Deployment/api",
	  "tenant_id":"acme",
	  "environment":"prod",
	  "provider":"sgx",
	  "quote_type":"sgx_quote",
	  "measurement":"m-1",
	  "lifecycle_state":"active",
	  "substrate_class":"confidential",
	  "trusted_measurements":["m-1"],
	  "require_credential_release":true
	}`))
	attestationReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	attestationReq.Header.Set("Content-Type", "application/json")
	attestationRec := httptest.NewRecorder()
	handler.ServeHTTP(attestationRec, attestationReq)
	if attestationRec.Code != http.StatusOK {
		t.Fatalf("expected attestation verify 200, got %d: %s", attestationRec.Code, attestationRec.Body.String())
	}

	mustSeedRuntimeFinding(t, store)

	simReq := httptest.NewRequest(http.MethodPost, "/v1/runtime/response/simulate?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
	  "subject_ref":"cluster-a/acme-prod/Deployment/api",
	  "action":"apply_network_isolation",
	  "summary":"simulate phase2 response"
	}`))
	simReq.Header.Set("Authorization", "Bearer operator-demo-token")
	simReq.Header.Set("Content-Type", "application/json")
	simRec := httptest.NewRecorder()
	handler.ServeHTTP(simRec, simReq)
	if simRec.Code != http.StatusOK {
		t.Fatalf("expected simulation 200, got %d: %s", simRec.Code, simRec.Body.String())
	}

	rollbackReq := httptest.NewRequest(http.MethodPost, "/v1/runtime/response/rollback-drill?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
	  "subject_ref":"cluster-a/acme-prod/Deployment/api",
	  "target_ref":"gitops:deployments/api@main",
	  "target_digest":"sha256:trusted",
	  "gitops_system":"argocd",
	  "target_verification_state":"verified",
	  "evidence_lock_state":"captured"
	}`))
	rollbackReq.Header.Set("Authorization", "Bearer operator-demo-token")
	rollbackReq.Header.Set("Content-Type", "application/json")
	rollbackRec := httptest.NewRecorder()
	handler.ServeHTTP(rollbackRec, rollbackReq)
	if rollbackRec.Code != http.StatusOK {
		t.Fatalf("expected rollback drill 200, got %d: %s", rollbackRec.Code, rollbackRec.Body.String())
	}

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/phase2/proofs?tenant_id=acme&environment=prod", nil)
	proofsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	proofsRec := httptest.NewRecorder()
	handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}
	var proofs phase2RuntimeProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if proofs.CurrentState != "phase2_core_slice_active" {
		t.Fatalf("expected active phase2 proofs, got %#v", proofs)
	}
}

func TestRuntimePhase2TrustedExecutionProfilesHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))
	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/trusted-execution-profiles?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var response phase2TrustedExecutionProfilesResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(response.Profiles) < 3 || len(response.Adapters) < 3 {
		t.Fatalf("expected profiles and adapters, got %#v", response)
	}
}

func TestRuntimePhase2ProofsRequireAllEvidenceTypes(t *testing.T) {
	store := audit.NewMemoryStore()
	handler := newHandlerWithAuth(store, "memory", mustStaticAuthConfig(t))

	payload, err := canonicalJSON(runtimePhase2EventPayload{
		SchemaVersion: runtimePhase2EventPayloadSchema,
		SubstrateTruth: &runtimesubstrate.SubstrateTruthRecord{
			SchemaVersion: runtimesubstrate.SubstrateTruthSchemaVersion,
			SubjectRef:    "cluster-a/acme-prod/Deployment/api",
			Workload:      runtimesubstrate.WorkloadIdentity{ClusterID: "cluster-a", Namespace: "acme-prod", WorkloadKind: "Deployment", Workload: "api"},
			Node:          runtimesubstrate.NodeIdentity{NodeID: "node-a", SubstrateClass: "confidential"},
			Attestation:   runtimesubstrate.AttestationBinding{Provider: "sgx", ObservedState: "verified"},
			ObservedAt:    mustPhase2Time(),
			CurrentState:  "runtime_truth_bound",
		},
	})
	if err != nil {
		t.Fatalf("canonicalJSON: %v", err)
	}
	if _, err := store.Ingest(context.Background(), audit.Event{
		Component:        runtimePhase2Component,
		EventType:        audit.EventTypeRuntimeSubstrateTruthRecorded,
		TenantID:         "acme",
		Environment:      "prod",
		Namespace:        "acme-prod",
		WorkloadKind:     "Deployment",
		Workload:         "api",
		Decision:         audit.DecisionAllow,
		RuntimeIntegrity: payload,
	}); err != nil {
		t.Fatalf("Ingest substrate truth: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/phase2/proofs?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var response phase2RuntimeProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if response.CurrentState != "phase2_core_incomplete" {
		t.Fatalf("expected incomplete proofs without attestation/simulation/rollback, got %#v", response)
	}
}

func mustSeedRuntimeFinding(t *testing.T, store audit.Store) {
	t.Helper()

	payload, err := canonicalJSON(runtimeIntegrityEventPayload{
		Observation: &runtimeObservationPayload{
			Node:        "node-a",
			Pod:         "api-0",
			ContainerID: "containerd://abc",
			EventType:   "memory_mapping_anomaly",
			Confidence:  runtimeConfidenceHigh,
			EventPayload: map[string]any{
				"unexpected_binary": "/tmp/loader",
			},
		},
	})
	if err != nil {
		t.Fatalf("canonicalJSON observation: %v", err)
	}
	if _, err := store.Ingest(context.Background(), audit.Event{
		Component:        runtimeIntegrityComponent,
		EventType:        audit.EventTypeRuntimeObservationRecorded,
		ClusterID:        "cluster-a",
		TenantID:         "acme",
		Environment:      "prod",
		Namespace:        "acme-prod",
		WorkloadKind:     "Deployment",
		Workload:         "api",
		Digest:           "sha256:abc",
		Decision:         audit.DecisionAllow,
		Reasons:          []string{"executable mapping anomaly observed"},
		RuntimeIntegrity: payload,
	}); err != nil {
		t.Fatalf("Ingest observation: %v", err)
	}
}

func mustPhase2Time() time.Time {
	return time.Date(2026, 4, 21, 11, 0, 0, 0, time.UTC)
}

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
)

func TestExecutionCoverageContracts(t *testing.T) {
	t.Setenv("CHANGELOCK_HANDOFF_SIGNING_SEED", "handoff-seed")

	fixture := forensicsTestFixture(t)
	sealFederationHandoffForTest(t, fixture.handler, incidentAudienceAuditorSafe)
	registerFederationPeerForTest(t, fixture.handler, federationPeerRequest{
		PeerID:            "peer-stale-3b",
		Organization:      "Acme Supplier",
		Region:            "eu-west",
		Cluster:           "supplier-west",
		TrustDomain:       "supplier.example",
		Endpoint:          "https://peer-stale-3b.example.invalid",
		PublicKeys:        []string{"pub-stale-3b"},
		Capabilities:      []string{"sealed_handoff"},
		PolicyRole:        federationPolicyRoleSupplier,
		AcceptedAudiences: []string{incidentAudienceAuditorSafe},
		LastSeen:          mustParseTimeRFC3339(t, time.Now().UTC().Add(-8*time.Hour).Format(time.RFC3339)),
	})
	executeStrictValidationRunForTest(t, fixture.handler, "/v1/validation/execute?tenant_id=acme&environment=prod&service=edge-gateway")

	req := httptest.NewRequest(http.MethodGet, "/v1/execution/coverage?tenant_id=acme&environment=prod&limit=20", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected execution coverage 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response executionCoverageResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode execution coverage: %v", err)
	}
	if response.SchemaVersion != executionCoverageSchemaVersion {
		t.Fatalf("expected execution coverage schema version, got %#v", response)
	}
	if !response.DisconnectedHandoffVerification.OfflineVerificationSupported || len(response.DisconnectedHandoffVerification.Items) == 0 {
		t.Fatalf("expected disconnected handoff verification coverage, got %#v", response.DisconnectedHandoffVerification)
	}
	if !response.OfflineFederation.OfflineModeSupported || len(response.OfflineFederation.LocalTrustAnchors) == 0 {
		t.Fatalf("expected offline federation anchor coverage, got %#v", response.OfflineFederation)
	}
	if len(response.OfflineFederation.StalePeers) == 0 || !response.DegradedMode.Active {
		t.Fatalf("expected degraded federation mode to be visible, got federation=%#v degraded=%#v", response.OfflineFederation, response.DegradedMode)
	}
	if !response.OfflineValidation.OfflineModeSupported || len(response.OfflineValidation.RecentCertificateRefs) == 0 {
		t.Fatalf("expected offline validation coverage, got %#v", response.OfflineValidation)
	}
	delayedSync := findExecutionCoverageCapability(t, response.CapabilityMatrix, "delayed_sync_policy_propagation")
	if delayedSync.CurrentState != "required" || len(response.DelayedSyncSemantics.BlockedClaims) == 0 {
		t.Fatalf("expected delayed sync semantics and capability matrix, got capability=%#v delayed=%#v", delayedSync, response.DelayedSyncSemantics)
	}
	if len(response.DegradedMode.EvidenceSurfaces) == 0 || response.DegradedMode.SummarySemantics == "" {
		t.Fatalf("expected degraded mode evidence summary semantics, got %#v", response.DegradedMode)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/execution/coverage/matrix?tenant_id=acme&environment=prod&limit=20", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected execution coverage matrix 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var matrix executionCoverageMatrixResponse
	if err := json.NewDecoder(rec.Body).Decode(&matrix); err != nil {
		t.Fatalf("decode execution coverage matrix: %v", err)
	}
	if matrix.SchemaVersion != executionCoverageMatrixSchema {
		t.Fatalf("expected execution coverage matrix schema version, got %#v", matrix)
	}
	if len(matrix.CapabilityMatrix) == 0 || matrix.DelayedSyncSemantics.PolicyPropagationModel == "" {
		t.Fatalf("expected explicit coverage matrix semantics, got %#v", matrix)
	}
}

func TestExecutionVMLineage(t *testing.T) {
	fixture := forensicsTestFixture(t)
	now := time.Now().UTC()

	mustIngestExecutionEvent(t, fixture.store, audit.Event{
		RequestID:                "vm-runtime-desired",
		Timestamp:                now.Add(-40 * time.Minute),
		Component:                "runtime-agent",
		EventType:                audit.EventTypeRuntimeDesiredStateRecorded,
		Decision:                 audit.DecisionAllow,
		TenantID:                 "acme",
		ClusterID:                "local",
		Repo:                     "acme/platform-vm",
		Environment:              "prod",
		Namespace:                "acme-prod",
		WorkloadKind:             "VirtualMachine",
		Workload:                 "payments-vm",
		ServiceAccount:           "vm-payments-sa",
		Digest:                   "sha256:payments-vm-v1",
		DesiredStateSourceRef:    "vm-image:payments-vm:v1",
		DesiredStateApprovalID:   "vm-approval-1",
		DesiredStateVerification: "verified",
		Reasons:                  []string{"vm desired state approved"},
		Evidence: &audit.Evidence{
			Artifact: &audit.ArtifactEvidence{
				SignerIdentity:           "vm-signer",
				AttestationPredicate:     "https://slsa.dev/provenance/v1",
				AttestationSubjectDigest: "sha256:payments-vm-v1",
				SBOMHash:                 "sbom-payments-vm-v1",
			},
			Runtime: &audit.RuntimeEvidence{
				ApprovedDigest:         "sha256:payments-vm-v1",
				ExpectedConfigHash:     "cfg-payments-vm-v1",
				ServiceAccountExpected: "vm-payments-sa",
			},
		},
	})
	mustIngestExecutionEvent(t, fixture.store, audit.Event{
		RequestID:                "vm-runtime-active",
		Timestamp:                now.Add(-20 * time.Minute),
		Component:                "runtime-agent",
		EventType:                audit.EventTypeRuntimeActiveStateObserved,
		Decision:                 audit.DecisionAllow,
		TenantID:                 "acme",
		ClusterID:                "local",
		Repo:                     "acme/platform-vm",
		Environment:              "prod",
		Namespace:                "acme-prod",
		WorkloadKind:             "VirtualMachine",
		Workload:                 "payments-vm",
		ServiceAccount:           "vm-payments-sa",
		Digest:                   "sha256:payments-vm-v1",
		ReconciliationStatus:     "in_sync",
		DesiredStateSourceRef:    "vm-image:payments-vm:v1",
		DesiredStateApprovalID:   "vm-approval-1",
		DesiredStateVerification: "verified",
		Reasons:                  []string{"vm runtime observed in sync"},
		Evidence: &audit.Evidence{
			Artifact: &audit.ArtifactEvidence{
				SignerIdentity:           "vm-signer",
				AttestationPredicate:     "https://slsa.dev/provenance/v1",
				AttestationSubjectDigest: "sha256:payments-vm-v1",
				SBOMHash:                 "sbom-payments-vm-v1",
			},
			Runtime: &audit.RuntimeEvidence{
				ApprovedDigest:         "sha256:payments-vm-v1",
				RunningDigest:          "sha256:payments-vm-v1",
				ExpectedConfigHash:     "cfg-payments-vm-v1",
				ActualConfigHash:       "cfg-payments-vm-v1",
				ServiceAccountExpected: "vm-payments-sa",
				ServiceAccountObserved: "vm-payments-sa",
			},
		},
	})
	executeStrictValidationRunForTest(t, fixture.handler, "/v1/validation/execute?tenant_id=acme&environment=prod&service=payments-vm")

	req := httptest.NewRequest(http.MethodGet, "/v1/execution/vm-lineage?tenant_id=acme&environment=prod&limit=20", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected vm lineage 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response executionVMLineageResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode vm lineage: %v", err)
	}
	if response.SchemaVersion != executionVMLineageSchemaVersion {
		t.Fatalf("expected vm lineage schema version, got %#v", response)
	}
	vm := findExecutionVMLineageItem(t, response.Items, "payments-vm")
	if vm.WorkloadKind != "VirtualMachine" || vm.DesiredStateVerification != "verified" {
		t.Fatalf("expected vm workload lineage with verified desired state, got %#v", vm)
	}
	if !containsString(vm.ExpectedSigners, "vm-signer") || !containsString(vm.TrustInputs, "attestation_provenance") {
		t.Fatalf("expected signer and attestation linkage on vm lineage, got %#v", vm)
	}
	if !vm.RuntimePosture.RuntimeModuleReady || vm.RuntimePosture.SchedulingGuidance.Decision != runtimeSchedulingAllowStandard {
		t.Fatalf("expected ready vm runtime posture, got %#v", vm.RuntimePosture)
	}
	if vm.RuntimeState.SBOMVerification.Status != runtimeSBOMStatusVerified || len(vm.ValidationRefs) == 0 {
		t.Fatalf("expected verified vm state and validation linkage, got %#v", vm)
	}
	if len(response.PolicyEvidenceParity.PolicyParity) == 0 || len(response.PolicyEvidenceParity.EvidenceParity) == 0 || len(response.PolicyEvidenceParity.Limitations) == 0 {
		t.Fatalf("expected vm policy/evidence parity contract, got %#v", response.PolicyEvidenceParity)
	}
}

func TestExecutionEphemeralCoverage(t *testing.T) {
	fixture := forensicsTestFixture(t)
	now := time.Now().UTC()

	mustIngestExecutionEvent(t, fixture.store, audit.Event{
		RequestID:                "job-runtime-desired",
		Timestamp:                now.Add(-30 * time.Minute),
		Component:                "runtime-agent",
		EventType:                audit.EventTypeRuntimeDesiredStateRecorded,
		Decision:                 audit.DecisionAllow,
		TenantID:                 "acme",
		ClusterID:                "local",
		Repo:                     "acme/platform-batch",
		Environment:              "prod",
		Namespace:                "acme-prod",
		WorkloadKind:             "Job",
		Workload:                 "nightly-billing-job",
		ServiceAccount:           "job-billing-sa",
		Digest:                   "sha256:job-billing-v1",
		DesiredStateSourceRef:    "job:nightly-billing",
		DesiredStateApprovalID:   "job-approval-1",
		DesiredStateVerification: "verified",
		Reasons:                  []string{"job desired state approved"},
		Evidence: &audit.Evidence{
			Artifact: &audit.ArtifactEvidence{
				SignerIdentity:           "job-signer",
				AttestationPredicate:     "https://slsa.dev/provenance/v1",
				AttestationSubjectDigest: "sha256:job-billing-v1",
				SBOMHash:                 "sbom-job-billing-v1",
			},
			Runtime: &audit.RuntimeEvidence{
				ApprovedDigest:         "sha256:job-billing-v1",
				ExpectedConfigHash:     "cfg-job-billing-v1",
				ServiceAccountExpected: "job-billing-sa",
			},
		},
	})
	mustIngestExecutionEvent(t, fixture.store, audit.Event{
		RequestID:                "job-runtime-active",
		Timestamp:                now.Add(-10 * time.Minute),
		Component:                "runtime-agent",
		EventType:                audit.EventTypeRuntimeActiveStateObserved,
		Decision:                 audit.DecisionAllow,
		TenantID:                 "acme",
		ClusterID:                "local",
		Repo:                     "acme/platform-batch",
		Environment:              "prod",
		Namespace:                "acme-prod",
		WorkloadKind:             "Job",
		Workload:                 "nightly-billing-job",
		ServiceAccount:           "job-billing-sa",
		Digest:                   "sha256:job-billing-v1",
		ReconciliationStatus:     "in_sync",
		DesiredStateSourceRef:    "job:nightly-billing",
		DesiredStateApprovalID:   "job-approval-1",
		DesiredStateVerification: "verified",
		Reasons:                  []string{"job runtime completed"},
		Evidence: &audit.Evidence{
			Runtime: &audit.RuntimeEvidence{
				ApprovedDigest:         "sha256:job-billing-v1",
				RunningDigest:          "sha256:job-billing-v1",
				ExpectedConfigHash:     "cfg-job-billing-v1",
				ActualConfigHash:       "cfg-job-billing-v1",
				ServiceAccountExpected: "job-billing-sa",
				ServiceAccountObserved: "job-billing-sa",
			},
		},
	})
	executeStrictValidationRunForTest(t, fixture.handler, "/v1/validation/execute?tenant_id=acme&environment=prod&service=nightly-billing-job")

	req := httptest.NewRequest(http.MethodGet, "/v1/execution/ephemeral?tenant_id=acme&environment=prod&limit=30", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected ephemeral execution coverage 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response executionEphemeralResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode ephemeral execution response: %v", err)
	}
	if response.SchemaVersion != executionEphemeralSchemaVersion {
		t.Fatalf("expected ephemeral schema version, got %#v", response)
	}
	job := findEphemeralExecutionItem(t, response.Items, executionSubstrateRuntimeJob, "nightly-billing-job")
	if job.SnapshotSemantics == "" || job.RetentionModel == "" || job.Status != "in_sync" {
		t.Fatalf("expected bounded runtime job lineage, got %#v", job)
	}
	validation := findEphemeralValidationItem(t, response.Items)
	if validation.CertificateRef == "" || validation.Mode == "" || validation.SubstrateType != executionSubstrateValidationExec {
		t.Fatalf("expected validation execution lineage, got %#v", validation)
	}
	if response.RetentionContract.SummarySemantics == "" || response.RetentionContract.CorrelationBehavior == "" || len(response.RetentionContract.RetentionSemantics) == 0 {
		t.Fatalf("expected explicit ephemeral retention contract, got %#v", response.RetentionContract)
	}
}

func executeStrictValidationRunForTest(t *testing.T, handler http.Handler, target string) validationExecutionRun {
	t.Helper()
	body := bytes.NewBufferString(`{"mode":"` + validationModeRegression + `","scenario_ids":["` + validationScenarioLatencyBudget + `"]}`)
	req := httptest.NewRequest(http.MethodPost, target, body)
	req.Header.Set("Authorization", "Bearer operator-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected strict validation execute 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var run validationExecutionRun
	if err := json.NewDecoder(rec.Body).Decode(&run); err != nil {
		t.Fatalf("decode strict validation run: %v", err)
	}
	return run
}

func mustIngestExecutionEvent(t *testing.T, store audit.Store, event audit.Event) {
	t.Helper()
	if _, err := store.Ingest(context.Background(), event); err != nil {
		t.Fatalf("ingest execution event: %v", err)
	}
}

func findExecutionVMLineageItem(t *testing.T, items []executionVMLineageItem, workload string) executionVMLineageItem {
	t.Helper()
	for _, item := range items {
		if item.Workload == workload {
			return item
		}
	}
	t.Fatalf("expected vm lineage item for workload %q, got %#v", workload, items)
	return executionVMLineageItem{}
}

func findEphemeralExecutionItem(t *testing.T, items []executionEphemeralItem, substrate, workload string) executionEphemeralItem {
	t.Helper()
	for _, item := range items {
		if item.SubstrateType == substrate && item.Workload == workload {
			return item
		}
	}
	t.Fatalf("expected ephemeral item for substrate=%q workload=%q, got %#v", substrate, workload, items)
	return executionEphemeralItem{}
}

func findEphemeralValidationItem(t *testing.T, items []executionEphemeralItem) executionEphemeralItem {
	t.Helper()
	for _, item := range items {
		if item.SubstrateType == executionSubstrateValidationExec {
			return item
		}
	}
	t.Fatalf("expected validation execution item, got %#v", items)
	return executionEphemeralItem{}
}

func findExecutionCoverageCapability(t *testing.T, items []executionCoverageCapability, capabilityID string) executionCoverageCapability {
	t.Helper()
	for _, item := range items {
		if item.CapabilityID == capabilityID {
			return item
		}
	}
	t.Fatalf("expected capability %q, got %#v", capabilityID, items)
	return executionCoverageCapability{}
}

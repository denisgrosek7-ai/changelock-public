package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/signing"
)

func TestExecutionAmbientReadiness(t *testing.T) {
	t.Setenv("CHANGELOCK_SELF_HEALING_MODE", "quarantine")
	t.Setenv("CHANGELOCK_RUNTIME_QUARANTINE_NETWORK_POLICY_ENABLED", "true")
	t.Setenv("CHANGELOCK_CLOSED_LOOP_REQUIRE_SIGNED_DESIRED_STATE", "true")

	fixture := forensicsTestFixture(t)
	mustIngestExecutionEvent(t, fixture.store, audit.Event{
		RequestID:                "ambient-runtime-active",
		Timestamp:                time.Now().UTC(),
		Component:                "runtime-agent",
		EventType:                audit.EventTypeRuntimeActiveStateObserved,
		Decision:                 audit.DecisionDeny,
		TenantID:                 "acme",
		ClusterID:                "local",
		Repo:                     "acme/platform-edge",
		Environment:              "prod",
		Namespace:                "acme-prod",
		WorkloadKind:             "Deployment",
		Workload:                 "edge-gateway",
		ReconciliationStatus:     "quarantined",
		QuarantineType:           "signer-identity",
		QuarantineReason:         "unauthorized signer identity remains active",
		DesiredStateVerification: "verified",
	})

	req := httptest.NewRequest(http.MethodGet, "/v1/execution/ambient-readiness?tenant_id=acme&environment=prod&limit=20", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected ambient readiness 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response executionAmbientReadinessResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode ambient readiness: %v", err)
	}
	if response.SchemaVersion != executionAmbientReadinessSchema {
		t.Fatalf("expected ambient readiness schema version, got %#v", response)
	}
	if response.CurrentState != "sidecarless_policy_overlay_candidate" {
		t.Fatalf("expected sidecarless policy overlay candidate state, got %#v", response)
	}
	overlay := findAmbientCapability(t, response.CapabilityMatrix, "network_policy_overlay")
	if overlay.CurrentState != "candidate_enabled" || overlay.SidecarRequirement != "not_required" {
		t.Fatalf("expected enabled sidecarless overlay capability, got %#v", overlay)
	}
	if response.ClosedLoopSummary.Quarantined == 0 || response.PostureSummary.TotalSubjects == 0 {
		t.Fatalf("expected closed-loop and posture summary evidence, got closed_loop=%#v posture=%#v", response.ClosedLoopSummary, response.PostureSummary)
	}
}

func TestExecutionConfidentialReadiness(t *testing.T) {
	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/execution/confidential-readiness?tenant_id=acme&environment=prod&limit=20", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected confidential readiness 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response executionConfidentialReadinessResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode confidential readiness: %v", err)
	}
	if response.SchemaVersion != executionConfidentialReadinessSchema {
		t.Fatalf("expected confidential readiness schema version, got %#v", response)
	}
	if response.CurrentState != "metadata_only" {
		t.Fatalf("expected metadata-only confidential readiness by default, got %#v", response)
	}
	if len(response.WorkloadMarkingContract) == 0 || len(response.AttestationLinkage) == 0 {
		t.Fatalf("expected explicit confidential readiness contracts, got %#v", response)
	}
	if response.ScopeSummary.TotalSubjects == 0 || response.ScopeSummary.EvidenceBackedClaims != 0 {
		t.Fatalf("expected bounded confidential summary without substrate evidence, got %#v", response.ScopeSummary)
	}
}

func TestHasConfidentialEvidenceUsesStructuredTokens(t *testing.T) {
	metadataOnly := runtimePostureState{
		SubjectRef:       "edge-service",
		ReadinessSignals: []string{"runtime_findings_present"},
		EvidenceRefs:     []string{"sample://runtime/edge-service"},
	}
	if hasConfidentialEvidence(metadataOnly) {
		t.Fatalf("expected free-form service names not to trigger confidential evidence, got %#v", metadataOnly)
	}

	evidenceBacked := runtimePostureState{
		SubjectRef:       "edge-service",
		ReadinessSignals: []string{"confidential_execution_requested"},
		EvidenceRefs:     []string{"sample://attestation/sev"},
		ActualTrustState: runtimePostureTrustState{AttestationInputs: []string{"sev_attestation"}},
	}
	if !hasConfidentialEvidence(evidenceBacked) {
		t.Fatalf("expected structured confidential evidence markers to be detected, got %#v", evidenceBacked)
	}
}

func TestExecutionComplianceReadiness(t *testing.T) {
	t.Setenv("CHANGELOCK_SIGNER_MODE", signing.ModeVaultTransit)
	t.Setenv("CHANGELOCK_SIGNER_KEY_ID", "changelock-control-plane")
	t.Setenv("CHANGELOCK_SIGNER_VERIFY_ON_READ", "true")
	t.Setenv("CHANGELOCK_VAULT_ADDR", "https://vault.example.com")
	t.Setenv("CHANGELOCK_VAULT_TOKEN", "test-token")
	t.Setenv("CHANGELOCK_VAULT_TRANSIT_KEY", "changelock-control-plane")
	t.Setenv("CHANGELOCK_TRUST_PUBLICATION_MODE", "preview")

	fixture := forensicsTestFixture(t)

	req := httptest.NewRequest(http.MethodGet, "/v1/execution/compliance-readiness?tenant_id=acme&environment=prod&repo=acme/platform-edge", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected compliance readiness 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response executionComplianceReadinessResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode compliance readiness: %v", err)
	}
	if response.SchemaVersion != executionComplianceReadinessSchema {
		t.Fatalf("expected compliance readiness schema version, got %#v", response)
	}
	if response.ComplianceModeMetadata.SignerMode != signing.ModeVaultTransit {
		t.Fatalf("expected vault transit signer mode in compliance metadata, got %#v", response.ComplianceModeMetadata)
	}
	exceptionsBoundary := findCryptoBoundary(t, response.CryptoModuleBoundaries, "exceptions_evidence_signing")
	if !exceptionsBoundary.Enabled || exceptionsBoundary.ProviderMode != signing.ModeVaultTransit {
		t.Fatalf("expected enabled provider-backed crypto boundary, got %#v", exceptionsBoundary)
	}
	if response.FIPSReadiness.State != "provider_backed_candidate" || len(response.StandardsMappings) == 0 {
		t.Fatalf("expected provider-backed crypto readiness and standards mappings, got fips=%#v mappings=%#v", response.FIPSReadiness, response.StandardsMappings)
	}
	if len(response.FormalClaimsExcluded) == 0 {
		t.Fatalf("expected explicit formal-claim exclusions, got %#v", response)
	}
}

func findAmbientCapability(t *testing.T, items []executionAmbientCapability, capabilityID string) executionAmbientCapability {
	t.Helper()
	for _, item := range items {
		if item.CapabilityID == capabilityID {
			return item
		}
	}
	t.Fatalf("expected ambient capability %q, got %#v", capabilityID, items)
	return executionAmbientCapability{}
}

func findCryptoBoundary(t *testing.T, items []executionCryptoModuleBoundary, boundaryID string) executionCryptoModuleBoundary {
	t.Helper()
	for _, item := range items {
		if item.BoundaryID == boundaryID {
			return item
		}
	}
	t.Fatalf("expected crypto boundary %q, got %#v", boundaryID, items)
	return executionCryptoModuleBoundary{}
}

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

func TestRuntimeSubstrateValCEnforcementTaxonomyHandler(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valc/enforcement-taxonomy?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc taxonomy 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var response runtimeSubstrateValCEnforcementTaxonomyResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode valc taxonomy: %v", err)
	}
	if response.CurrentState != runtimesubstrate.RuntimeSubstrateEnforcementTaxonomyStateActive {
		t.Fatalf("expected active taxonomy, got %#v", response)
	}
	if !containsString(response.Taxonomy.EnforcementClasses, runtimesubstrate.RuntimeSubstrateEnforcementClassPrevent) {
		t.Fatalf("expected prevent class, got %#v", response.Taxonomy.EnforcementClasses)
	}
	if !containsString(response.Taxonomy.DecisionModes, runtimesubstrate.RuntimeSubstrateDecisionModeTerminateAndRecover) {
		t.Fatalf("expected terminate-and-recover mode, got %#v", response.Taxonomy.DecisionModes)
	}
}

func TestRuntimeSubstrateValCHandlersAndProofs(t *testing.T) {
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
	postValCDecisionRequest(t, fixture.handler, "/v1/runtime/quarantine?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-VALC-1"}`)
	postValCDecisionRequest(t, fixture.handler, "/v1/runtime/restart-trusted?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-VALC-2"}`)
	postValCDecisionRequest(t, fixture.handler, "/v1/hardening/quarantine?tenant_id=acme&environment=prod", `{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-VALC-3"}`)

	seedRuntimeSubstrateValCPreventExecution(t, fixture.store)

	req := httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valc/action-catalog?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc action catalog 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var catalog runtimeSubstrateValCActionCatalogResponse
	if err := json.NewDecoder(rec.Body).Decode(&catalog); err != nil {
		t.Fatalf("decode valc action catalog: %v", err)
	}
	if catalog.CurrentState != runtimesubstrate.RuntimeSubstrateActionCatalogStateActive {
		t.Fatalf("expected active valc action catalog, got %#v", catalog)
	}
	if !hasValCActionClass(catalog.Items, runtimesubstrate.RuntimeSubstrateEnforcementClassPrevent) {
		t.Fatalf("expected prevent class in action catalog, got %#v", catalog.Items)
	}
	if !hasValCActionClass(catalog.Items, runtimesubstrate.RuntimeSubstrateEnforcementClassTerminate) {
		t.Fatalf("expected terminate class in action catalog, got %#v", catalog.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valc/policy-hook-mapping?tenant_id=acme&environment=prod", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc hook mapping 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var hooks runtimeSubstrateValCPolicyHookMappingResponse
	if err := json.NewDecoder(rec.Body).Decode(&hooks); err != nil {
		t.Fatalf("decode valc hook mapping: %v", err)
	}
	if hooks.CurrentState != runtimesubstrate.RuntimeSubstratePolicyHookMappingStateActive {
		t.Fatalf("expected active hook mapping, got %#v", hooks)
	}
	if !hasValCHookMapping(hooks.Items, "hardening.block_exec_class_next_restart") {
		t.Fatalf("expected hardening next-restart hook mapping, got %#v", hooks.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valc/decision-audit?tenant_id=acme&environment=prod&limit=50", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc decision audit 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var auditResponse runtimeSubstrateValCDecisionAuditResponse
	if err := json.NewDecoder(rec.Body).Decode(&auditResponse); err != nil {
		t.Fatalf("decode valc decision audit: %v", err)
	}
	if auditResponse.CurrentState != runtimesubstrate.RuntimeSubstrateDecisionAuditStateActive {
		t.Fatalf("expected active decision audit, got %#v", auditResponse)
	}
	if !hasValCDecisionAction(auditResponse.Items, "runtime.apply_network_isolation") {
		t.Fatalf("expected runtime containment decision audit item, got %#v", auditResponse.Items)
	}
	if !hasValCDecisionAction(auditResponse.Items, "runtime.restart_from_trusted_image") {
		t.Fatalf("expected runtime terminate decision audit item, got %#v", auditResponse.Items)
	}
	if !hasValCDecisionAction(auditResponse.Items, "hardening.block_exec_class_next_restart") {
		t.Fatalf("expected hardening prevent decision audit item, got %#v", auditResponse.Items)
	}

	req = httptest.NewRequest(http.MethodGet, "/v1/runtime/substrate-depth/valc/proofs?tenant_id=acme&environment=prod&limit=50", nil)
	req.Header.Set("Authorization", "Bearer viewer-demo-token")
	rec = httptest.NewRecorder()
	fixture.handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected valc proofs 200, got %d: %s", rec.Code, rec.Body.String())
	}

	var proofs runtimeSubstrateValCProofsResponse
	if err := json.NewDecoder(rec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode valc proofs: %v", err)
	}
	if proofs.CurrentState != runtimesubstrate.RuntimeSubstrateValCStateActive {
		t.Fatalf("expected active valc proofs, got %#v", proofs)
	}
	if proofs.ValBState != runtimesubstrate.RuntimeSubstrateValBStateActive {
		t.Fatalf("expected active valb dependency, got %#v", proofs)
	}
	if proofs.DecisionAuditState != runtimesubstrate.RuntimeSubstrateDecisionAuditStateActive {
		t.Fatalf("expected active decision audit state, got %#v", proofs)
	}
}

func postValCDecisionRequest(t *testing.T, handler http.Handler, path, body string) {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBufferString(body))
	req.Header.Set("Authorization", "Bearer operator-demo-token")
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 for %s, got %d: %s", path, rec.Code, rec.Body.String())
	}
}

func seedRuntimeSubstrateValCPreventExecution(t *testing.T, store audit.Store) {
	t.Helper()
	executedAt := time.Date(2026, 4, 22, 11, 30, 0, 0, time.UTC)
	payload, err := canonicalJSON(hardeningEventPayload{
		PolicyDecision: &hardeningPolicyDecision{
			SchemaVersion:    hardeningPolicyDecisionSchemaVersion,
			DecisionID:       "hardening-policy-edge-prevent",
			PolicyRef:        "runtime_closed_loop_hardening.v1:profile_deviation:process_hardening",
			ApprovalMode:     recommendationApprovalHumanReview,
			ApprovalRequired: true,
			RollbackRequired: true,
		},
		Execution: &hardeningExecutionRecord{
			SchemaVersion:   hardeningExecutionSchemaVersion,
			ExecutionID:     "hardening-exec-edge-prevent",
			SubjectRef:      runtimeSubjectRef("local", "acme-prod", "Deployment", "edge-gateway"),
			DecisionRef:     "hardening-policy-edge-prevent",
			ExecutedAt:      executedAt,
			ExecutionResult: "process_hardening_staged_for_next_restart",
			ActionsApplied: []hardeningAction{
				{
					SchemaVersion: hardeningActionSchemaVersion,
					ActionID:      "hardening-action-edge-profile",
					ActionType:    hardeningActionTightenRuntimeProfile,
					SubjectRef:    runtimeSubjectRef("local", "acme-prod", "Deployment", "edge-gateway"),
					Scope:         "workload_only",
					Parameters:    map[string]any{"enforcement_timing": "next_restart"},
					IsImmediate:   false,
					IsReversible:  true,
				},
				{
					SchemaVersion: hardeningActionSchemaVersion,
					ActionID:      "hardening-action-edge-exec-class",
					ActionType:    hardeningActionBlockExecClass,
					SubjectRef:    runtimeSubjectRef("local", "acme-prod", "Deployment", "edge-gateway"),
					Scope:         "workload_only",
					Parameters:    map[string]any{"blocked_exec_class": "unknown_or_unsigned", "enforcement_timing": "next_restart"},
					IsImmediate:   false,
					IsReversible:  true,
				},
			},
			RollbackPlan: []string{"rollback staged process hardening after clean verification"},
		},
	})
	if err != nil {
		t.Fatalf("canonical hardening payload: %v", err)
	}
	_, err = store.Ingest(context.Background(), audit.Event{
		RequestID:        "hardening-exec-edge-prevent",
		Timestamp:        executedAt,
		Component:        hardeningComponent,
		EventType:        audit.EventTypeHardeningActionApplied,
		TenantID:         "acme",
		Environment:      "prod",
		ClusterID:        "local",
		Namespace:        "acme-prod",
		WorkloadKind:     "Deployment",
		Workload:         "edge-gateway",
		Decision:         audit.DecisionAllow,
		DriftResult:      hardeningModeProcessHardening,
		DriftSeverity:    "high",
		RuntimeIntegrity: payload,
	})
	if err != nil {
		t.Fatalf("seed valc prevent execution: %v", err)
	}
}

func hasValCActionClass(items []runtimesubstrate.RuntimeSubstrateEnforcementActionCatalogItem, class string) bool {
	for _, item := range items {
		if item.GuaranteeClass == class {
			return true
		}
	}
	return false
}

func hasValCHookMapping(items []runtimesubstrate.RuntimeSubstratePolicyHookMapping, actionID string) bool {
	for _, item := range items {
		if item.ActionID == actionID {
			return true
		}
	}
	return false
}

func hasValCDecisionAction(items []runtimesubstrate.RuntimeSubstrateDecisionAuditRecord, actionID string) bool {
	for _, item := range items {
		if item.ActionID == actionID {
			return true
		}
	}
	return false
}

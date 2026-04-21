package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRuntimeResponsePolicySurfaceAndEnforcementMetadata(t *testing.T) {
	fixture := forensicsTestFixture(t)

	policyReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/response-policy?tenant_id=acme&environment=prod", nil)
	policyReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	policyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(policyRec, policyReq)
	if policyRec.Code != http.StatusOK {
		t.Fatalf("expected runtime response policy 200, got %d: %s", policyRec.Code, policyRec.Body.String())
	}

	var policy runtimeResponsePolicyResponse
	if err := json.NewDecoder(policyRec.Body).Decode(&policy); err != nil {
		t.Fatalf("decode runtime response policy: %v", err)
	}
	if policy.SchemaVersion != runtimeResponsePolicySchemaVersion || !policy.LeastInvasiveFirst {
		t.Fatalf("expected runtime response policy contract, got %#v", policy)
	}
	if !containsString(policy.AutonomousActions, runtimeActionCaptureForensics) || !containsString(policy.ApprovalRequiredActions, runtimeActionApplyNetworkIsolation) {
		t.Fatalf("expected explicit autonomy split in policy surface, got %#v", policy)
	}
	if len(policy.ForensicFirstPolicy.SnapshotFirstRequiredFor) == 0 || len(policy.BlastRadiusSafetyLimits) == 0 {
		t.Fatalf("expected forensic-first and blast-radius policy sections, got %#v", policy)
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
	if evaluated.ResponseMode != runtimeResponseModeApprovalGated || !evaluated.ApprovalRequired {
		t.Fatalf("expected approval-gated runtime decision metadata, got %#v", evaluated)
	}
	if evaluated.ConfidenceLevel == "" || !evaluated.ForensicFirst || !evaluated.RollbackRequired || evaluated.TTL == "" {
		t.Fatalf("expected confidence, forensic-first, rollback, and ttl metadata, got %#v", evaluated)
	}
	if evaluated.LeastInvasiveRank <= 0 || evaluated.SafetyLimitRef == "" {
		t.Fatalf("expected least-invasive rank and safety limit ref, got %#v", evaluated)
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
	var forensicDecision runtimeEnforcementDecision
	if err := json.NewDecoder(forensicRec.Body).Decode(&forensicDecision); err != nil {
		t.Fatalf("decode runtime forensic snapshot: %v", err)
	}
	if forensicDecision.ApprovalRequired || forensicDecision.ResponseMode != runtimeResponseModeBoundedAutonomous {
		t.Fatalf("expected autonomous forensic-first snapshot path, got %#v", forensicDecision)
	}
	if forensicDecision.LeastInvasiveRank >= evaluated.LeastInvasiveRank {
		t.Fatalf("expected forensic snapshot to stay ahead of network isolation in least-invasive ordering, got forensic=%#v evaluate=%#v", forensicDecision, evaluated)
	}

	quarantineReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/runtime/quarantine?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`","approval_ref":"APR-3A-1"}`),
	)
	quarantineReq.Header.Set("Authorization", "Bearer operator-demo-token")
	quarantineReq.Header.Set("Content-Type", "application/json")
	quarantineRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(quarantineRec, quarantineReq)
	if quarantineRec.Code != http.StatusOK {
		t.Fatalf("expected runtime quarantine 200, got %d: %s", quarantineRec.Code, quarantineRec.Body.String())
	}

	historyReq := httptest.NewRequest(http.MethodGet, "/v1/runtime/enforcement?tenant_id=acme&environment=prod&limit=20", nil)
	historyReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	historyRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(historyRec, historyReq)
	if historyRec.Code != http.StatusOK {
		t.Fatalf("expected runtime enforcement history 200, got %d: %s", historyRec.Code, historyRec.Body.String())
	}
	var history runtimeEnforcementListResponse
	if err := json.NewDecoder(historyRec.Body).Decode(&history); err != nil {
		t.Fatalf("decode runtime enforcement history: %v", err)
	}
	item := findRuntimeEnforcementDecision(t, history.Items, runtimeActionApplyNetworkIsolation, "network_isolation_applied")
	if item.TTL == "" || !item.RollbackRequired || item.SafetyLimitRef == "" {
		t.Fatalf("expected runtime enforcement history to retain response metadata, got %#v", item)
	}
}

func TestHardeningDecisionCarriesResponseMetadata(t *testing.T) {
	fixture := forensicsTestFixture(t)

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

	evaluateReq := httptest.NewRequest(
		http.MethodPost,
		"/v1/hardening/evaluate?tenant_id=acme&environment=prod",
		bytes.NewBufferString(`{"finding_id":"`+binaryFinding.FindingID+`"}`),
	)
	evaluateReq.Header.Set("Authorization", "Bearer operator-demo-token")
	evaluateReq.Header.Set("Content-Type", "application/json")
	evaluateRec := httptest.NewRecorder()
	fixture.handler.ServeHTTP(evaluateRec, evaluateReq)
	if evaluateRec.Code != http.StatusOK {
		t.Fatalf("expected hardening evaluate 200, got %d: %s", evaluateRec.Code, evaluateRec.Body.String())
	}

	var evaluated hardeningEvaluationResponse
	if err := json.NewDecoder(evaluateRec.Body).Decode(&evaluated); err != nil {
		t.Fatalf("decode hardening evaluate: %v", err)
	}
	if evaluated.PolicyDecision.ResponseMode != runtimeResponseModeApprovalGated || !evaluated.PolicyDecision.ApprovalRequired {
		t.Fatalf("expected hardening decision approval metadata, got %#v", evaluated.PolicyDecision)
	}
	if evaluated.PolicyDecision.ConfidenceLevel == "" || !evaluated.PolicyDecision.ForensicFirst {
		t.Fatalf("expected hardening confidence and forensic-first metadata, got %#v", evaluated.PolicyDecision)
	}
	if !evaluated.PolicyDecision.RollbackRequired || evaluated.PolicyDecision.TTL == "" || evaluated.PolicyDecision.LeastInvasiveRank <= 0 || evaluated.PolicyDecision.SafetyLimitRef == "" {
		t.Fatalf("expected ttl, rollback, least-invasive, and safety metadata on hardening decision, got %#v", evaluated.PolicyDecision)
	}
	if len(evaluated.Actions) < 2 || evaluated.Actions[0].ActionType != hardeningActionRequestForensics {
		t.Fatalf("expected least-invasive hardening ordering to stay forensic-first, got %#v", evaluated.Actions)
	}
}

func findRuntimeEnforcementDecision(t *testing.T, items []runtimeEnforcementDecision, action string, result string) runtimeEnforcementDecision {
	t.Helper()
	for _, item := range items {
		if item.Action == action && item.ExecutionResult == result {
			return item
		}
	}
	t.Fatalf("expected runtime enforcement action %q result %q, got %#v", action, result, items)
	return runtimeEnforcementDecision{}
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/denisgrosek/changelock/internal/audit"
)

func TestPhase4EnterpriseFlowAndProofs(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	workflowReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/workflow/lifecycle?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "workflow_id":"wf-100",
	    "artifact_type":"finding",
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "severity":"critical",
	    "requested_state":"resolved",
	    "validation_required":true,
	    "validation_state":"verified",
	    "owners":{"finding_owner":"team-api","remediation_owner":"team-api","approver":"security-admin","compliance_owner":"compliance-lead"},
	    "evidence_refs":["event://deploy-gate/1","report:/v1/runtime/phase2/proofs"]
	  }
	}`))
	workflowReq.Header.Set("Authorization", "Bearer operator-demo-token")
	workflowReq.Header.Set("Content-Type", "application/json")
	workflowRec := httptest.NewRecorder()
	handler.ServeHTTP(workflowRec, workflowReq)
	if workflowRec.Code != http.StatusCreated {
		t.Fatalf("expected workflow 201, got %d: %s", workflowRec.Code, workflowRec.Body.String())
	}

	connectorReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/workflow/connectors/reconcile?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "workflow_id":"wf-100",
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "connector_system":"jira",
	    "connector_ref":"JIRA-100",
	    "object_type":"ticket",
	    "internal_state":"resolved",
	    "external_state":"closed",
	    "validation_state":"verified",
	    "health":{"current_state":"healthy"},
	    "evidence_refs":["ticket://JIRA-100"]
	  }
	}`))
	connectorReq.Header.Set("Authorization", "Bearer operator-demo-token")
	connectorReq.Header.Set("Content-Type", "application/json")
	connectorRec := httptest.NewRecorder()
	handler.ServeHTTP(connectorRec, connectorReq)
	if connectorRec.Code != http.StatusCreated {
		t.Fatalf("expected reconciliation 201, got %d: %s", connectorRec.Code, connectorRec.Body.String())
	}

	partnerReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/partner-trust/intake?tenant_id=acme&environment=prod&repo=github.com/acme/api&partner_id=vendor-a", bytes.NewBufferString(`{
	  "input":{
	    "partner_id":"vendor-a",
	    "organization":"Vendor A",
	    "trust_domain":"suppliers.acme",
	    "handoff_ref":"handoff-100",
	    "verification_status":"verified",
	    "freshness_state":"fresh",
	    "policy_compatibility":"compatible",
	    "incident_disclosure_status":"shared",
	    "partner_visible_evidence":["sealed://proof/100"],
	    "evidence_refs":["handoff://100"]
	  }
	}`))
	partnerReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	partnerReq.Header.Set("Content-Type", "application/json")
	partnerRec := httptest.NewRecorder()
	handler.ServeHTTP(partnerRec, partnerReq)
	if partnerRec.Code != http.StatusCreated {
		t.Fatalf("expected partner intake 201, got %d: %s", partnerRec.Code, partnerRec.Body.String())
	}

	complianceReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/governance/compliance-mapping?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "control_family":"soc2.cc7",
	    "control_id":"CC7.2",
	    "coverage_state":"full",
	    "freshness_state":"fresh",
	    "evidence_refs":["event://deploy-gate/1"],
	    "technical_event_refs":["deploy_gate_decision","runtime_phase2_proofs"]
	  }
	}`))
	complianceReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	complianceReq.Header.Set("Content-Type", "application/json")
	complianceRec := httptest.NewRecorder()
	handler.ServeHTTP(complianceRec, complianceReq)
	if complianceRec.Code != http.StatusCreated {
		t.Fatalf("expected compliance mapping 201, got %d: %s", complianceRec.Code, complianceRec.Body.String())
	}

	driftReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/governance/policy-drift?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "input":{
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "actor":"security-admin",
	    "previous_mode":"deny",
	    "current_mode":"exception",
	    "change_reason":"temporary maintenance window",
	    "exception_id":"exc-100",
	    "impacted_controls":["CC7.2"],
	    "evidence_refs":["exception://exc-100"]
	  }
	}`))
	driftReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	driftReq.Header.Set("Content-Type", "application/json")
	driftRec := httptest.NewRecorder()
	handler.ServeHTTP(driftRec, driftReq)
	if driftRec.Code != http.StatusCreated {
		t.Fatalf("expected policy drift 201, got %d: %s", driftRec.Code, driftRec.Body.String())
	}

	reportReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/governance/executive-report?tenant_id=acme&environment=prod&repo=github.com/acme/api", bytes.NewBufferString(`{
	  "scope_ref":"tenant:acme"
	}`))
	reportReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	reportReq.Header.Set("Content-Type", "application/json")
	reportRec := httptest.NewRecorder()
	handler.ServeHTTP(reportRec, reportReq)
	if reportRec.Code != http.StatusOK {
		t.Fatalf("expected executive report 200, got %d: %s", reportRec.Code, reportRec.Body.String())
	}

	dashboardReq := httptest.NewRequest(http.MethodGet, "/v1/enterprise/partner-trust/dashboard?tenant_id=acme&environment=prod&repo=github.com/acme/api&partner_id=vendor-a", nil)
	dashboardReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	dashboardRec := httptest.NewRecorder()
	handler.ServeHTTP(dashboardRec, dashboardReq)
	if dashboardRec.Code != http.StatusOK {
		t.Fatalf("expected partner dashboard 200, got %d: %s", dashboardRec.Code, dashboardRec.Body.String())
	}
	var dashboard phase4PartnerDashboardResponse
	if err := json.NewDecoder(dashboardRec.Body).Decode(&dashboard); err != nil {
		t.Fatalf("decode dashboard: %v", err)
	}
	if len(dashboard.Items) != 1 || !dashboard.Items[0].SensitiveSignalsRedacted {
		t.Fatalf("expected bounded partner dashboard, got %#v", dashboard)
	}

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/enterprise/phase4/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&subject_ref=cluster-a/acme-prod/Deployment/api&partner_id=vendor-a", nil)
	proofsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	proofsRec := httptest.NewRecorder()
	handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected phase4 proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}
	var proofs phase4ProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if proofs.CurrentState != phase4ProofStateActive {
		t.Fatalf("expected active phase4 proofs, got %#v", proofs)
	}
}

func TestPhase4ProofsRequireAllEvidenceTypes(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	workflowReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/workflow/lifecycle?tenant_id=acme&environment=prod", bytes.NewBufferString(`{
	  "input":{
	    "workflow_id":"wf-100",
	    "artifact_type":"finding",
	    "subject_ref":"cluster-a/acme-prod/Deployment/api",
	    "severity":"high",
	    "requested_state":"resolved",
	    "validation_required":true,
	    "validation_state":"verified",
	    "owners":{"finding_owner":"team-api"},
	    "evidence_refs":["event://1"]
	  }
	}`))
	workflowReq.Header.Set("Authorization", "Bearer operator-demo-token")
	workflowReq.Header.Set("Content-Type", "application/json")
	workflowRec := httptest.NewRecorder()
	handler.ServeHTTP(workflowRec, workflowReq)
	if workflowRec.Code != http.StatusCreated {
		t.Fatalf("expected workflow 201, got %d: %s", workflowRec.Code, workflowRec.Body.String())
	}

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/enterprise/phase4/proofs?tenant_id=acme&environment=prod&subject_ref=cluster-a/acme-prod/Deployment/api", nil)
	proofsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	proofsRec := httptest.NewRecorder()
	handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected phase4 proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}
	var proofs phase4ProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if proofs.CurrentState != phase4ProofStateIncomplete {
		t.Fatalf("expected incomplete proofs without connector/partner/compliance/drift/executive artifacts, got %#v", proofs)
	}
}

func TestPhase4PartnerIntakeUsesQueryPartnerIDForPostAndFilters(t *testing.T) {
	handler := newHandlerWithAuth(audit.NewMemoryStore(), "memory", mustStaticAuthConfig(t))

	partnerReq := httptest.NewRequest(http.MethodPost, "/v1/enterprise/partner-trust/intake?tenant_id=acme&environment=prod&repo=github.com/acme/api&partner_id=vendor-a", bytes.NewBufferString(`{
	  "input":{
	    "organization":"Vendor A",
	    "trust_domain":"suppliers.acme",
	    "handoff_ref":"handoff-200",
	    "verification_status":"verified",
	    "freshness_state":"fresh",
	    "policy_compatibility":"compatible",
	    "incident_disclosure_status":"shared",
	    "partner_visible_evidence":["sealed://proof/200"],
	    "evidence_refs":["handoff://200"]
	  }
	}`))
	partnerReq.Header.Set("Authorization", "Bearer security-admin-demo-token")
	partnerReq.Header.Set("Content-Type", "application/json")
	partnerRec := httptest.NewRecorder()
	handler.ServeHTTP(partnerRec, partnerReq)
	if partnerRec.Code != http.StatusCreated {
		t.Fatalf("expected partner intake 201, got %d: %s", partnerRec.Code, partnerRec.Body.String())
	}

	intakeReq := httptest.NewRequest(http.MethodGet, "/v1/enterprise/partner-trust/intake?tenant_id=acme&environment=prod&repo=github.com/acme/api&partner_id=vendor-a", nil)
	intakeReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	intakeRec := httptest.NewRecorder()
	handler.ServeHTTP(intakeRec, intakeReq)
	if intakeRec.Code != http.StatusOK {
		t.Fatalf("expected intake list 200, got %d: %s", intakeRec.Code, intakeRec.Body.String())
	}
	var intake phase4PartnerIntakeListResponse
	if err := json.NewDecoder(intakeRec.Body).Decode(&intake); err != nil {
		t.Fatalf("decode intake list: %v", err)
	}
	if len(intake.Items) != 1 || intake.Items[0].PartnerID != "vendor-a" {
		t.Fatalf("expected partner intake to inherit query partner_id, got %#v", intake)
	}

	dashboardReq := httptest.NewRequest(http.MethodGet, "/v1/enterprise/partner-trust/dashboard?tenant_id=acme&environment=prod&repo=github.com/acme/api&partner_id=vendor-a", nil)
	dashboardReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	dashboardRec := httptest.NewRecorder()
	handler.ServeHTTP(dashboardRec, dashboardReq)
	if dashboardRec.Code != http.StatusOK {
		t.Fatalf("expected partner dashboard 200, got %d: %s", dashboardRec.Code, dashboardRec.Body.String())
	}
	var dashboard phase4PartnerDashboardResponse
	if err := json.NewDecoder(dashboardRec.Body).Decode(&dashboard); err != nil {
		t.Fatalf("decode dashboard: %v", err)
	}
	if len(dashboard.Items) != 1 || dashboard.Items[0].PartnerID != "vendor-a" {
		t.Fatalf("expected partner dashboard to filter by inherited query partner_id, got %#v", dashboard)
	}

	proofsReq := httptest.NewRequest(http.MethodGet, "/v1/enterprise/phase4/proofs?tenant_id=acme&environment=prod&repo=github.com/acme/api&partner_id=vendor-a", nil)
	proofsReq.Header.Set("Authorization", "Bearer viewer-demo-token")
	proofsRec := httptest.NewRecorder()
	handler.ServeHTTP(proofsRec, proofsReq)
	if proofsRec.Code != http.StatusOK {
		t.Fatalf("expected phase4 proofs 200, got %d: %s", proofsRec.Code, proofsRec.Body.String())
	}
	var proofs phase4ProofsResponse
	if err := json.NewDecoder(proofsRec.Body).Decode(&proofs); err != nil {
		t.Fatalf("decode proofs: %v", err)
	}
	if len(proofs.PartnerArtifacts) != 1 || proofs.PartnerArtifacts[0].PartnerID != "vendor-a" {
		t.Fatalf("expected phase4 proofs to expose partner artifact under query partner_id, got %#v", proofs)
	}
}

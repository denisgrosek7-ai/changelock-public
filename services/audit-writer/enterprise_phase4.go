package main

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/denisgrosek/changelock/internal/audit"
	"github.com/denisgrosek/changelock/internal/auth"
	"github.com/denisgrosek/changelock/internal/compliance"
	"github.com/denisgrosek/changelock/internal/connectors"
	"github.com/denisgrosek/changelock/internal/handoff"
	"github.com/denisgrosek/changelock/internal/httpjson"
	"github.com/denisgrosek/changelock/internal/workflow"
)

const (
	phase4EnterpriseComponent        = "phase4-enterprise-manager"
	phase4EnterprisePayloadSchema    = "4.enterprise_event_payload.v1"
	phase4WorkflowListSchema         = "4.enterprise_workflow_list.v1"
	phase4ConnectorListSchema        = "4.enterprise_connector_reconciliation_list.v1"
	phase4PartnerIntakeListSchema    = "4.enterprise_partner_intake_list.v1"
	phase4PartnerDashboardListSchema = "4.enterprise_partner_dashboard_list.v1"
	phase4ComplianceListSchema       = "4.enterprise_compliance_list.v1"
	phase4PolicyDriftListSchema      = "4.enterprise_policy_drift_list.v1"
	phase4ExecutiveReportListSchema  = "4.enterprise_executive_report_list.v1"
	phase4ProofsSchema               = "4.enterprise_phase4_proofs.v1"
	phase4ProofStateIncomplete       = "phase4_core_incomplete"
	phase4ProofStateActive           = "phase4_core_slice_active"
	phase4WorkflowStateActive        = "workflow_orchestration_active"
	phase4ConnectorStateActive       = "connector_reconciliation_active"
	phase4PartnerStateActive         = "partner_trust_active"
	phase4ComplianceStateActive      = "governance_compliance_active"
	phase4PolicyDriftStateActive     = "policy_drift_governance_active"
	phase4ExecutiveStateActive       = "executive_reporting_active"
)

type phase4EnterprisePayload struct {
	SchemaVersion  string                              `json:"schema_version"`
	Workflow       *workflow.LifecycleRecord           `json:"workflow,omitempty"`
	Reconciliation *connectors.ReconciliationRecord    `json:"reconciliation,omitempty"`
	PartnerIntake  *handoff.IntakeRecord               `json:"partner_intake,omitempty"`
	Compliance     *compliance.ComplianceMappingRecord `json:"compliance,omitempty"`
	PolicyDrift    *compliance.PolicyDriftRecord       `json:"policy_drift,omitempty"`
	Executive      *compliance.ExecutiveReport         `json:"executive,omitempty"`
}

type phase4EnterpriseFilter struct {
	ClusterID   string
	TenantID    string
	Environment string
	Repo        string
	SubjectRef  string
	WorkflowID  string
	PartnerID   string
	Limit       int
}

type phase4WorkflowRequest struct {
	Input workflow.LifecycleInput `json:"input"`
}

type phase4ConnectorRequest struct {
	Input connectors.ReconciliationInput `json:"input"`
}

type phase4PartnerIntakeRequest struct {
	Input handoff.IntakeInput `json:"input"`
}

type phase4ComplianceRequest struct {
	Input compliance.MappingInput `json:"input"`
}

type phase4PolicyDriftRequest struct {
	Input compliance.DriftInput `json:"input"`
}

type phase4ExecutiveReportRequest struct {
	ScopeRef string `json:"scope_ref,omitempty"`
}

type phase4WorkflowListResponse struct {
	SchemaVersion string                     `json:"schema_version"`
	CurrentState  string                     `json:"current_state"`
	Items         []workflow.LifecycleRecord `json:"items"`
	Limitations   []string                   `json:"limitations,omitempty"`
}

type phase4WorkflowResponse struct {
	Status   string                   `json:"status"`
	Workflow workflow.LifecycleRecord `json:"workflow"`
}

type phase4ConnectorListResponse struct {
	SchemaVersion string                            `json:"schema_version"`
	CurrentState  string                            `json:"current_state"`
	Items         []connectors.ReconciliationRecord `json:"items"`
	Limitations   []string                          `json:"limitations,omitempty"`
}

type phase4ConnectorResponse struct {
	Status         string                          `json:"status"`
	Reconciliation connectors.ReconciliationRecord `json:"reconciliation"`
}

type phase4PartnerIntakeListResponse struct {
	SchemaVersion string                 `json:"schema_version"`
	CurrentState  string                 `json:"current_state"`
	Items         []handoff.IntakeRecord `json:"items"`
	Limitations   []string               `json:"limitations,omitempty"`
}

type phase4PartnerIntakeResponse struct {
	Status string               `json:"status"`
	Intake handoff.IntakeRecord `json:"intake"`
}

type phase4PartnerDashboardResponse struct {
	SchemaVersion string                        `json:"schema_version"`
	CurrentState  string                        `json:"current_state"`
	Items         []handoff.DashboardProjection `json:"items"`
	Limitations   []string                      `json:"limitations,omitempty"`
}

type phase4ComplianceListResponse struct {
	SchemaVersion string                               `json:"schema_version"`
	CurrentState  string                               `json:"current_state"`
	Items         []compliance.ComplianceMappingRecord `json:"items"`
	Limitations   []string                             `json:"limitations,omitempty"`
}

type phase4ComplianceResponse struct {
	Status     string                             `json:"status"`
	Compliance compliance.ComplianceMappingRecord `json:"compliance"`
}

type phase4PolicyDriftListResponse struct {
	SchemaVersion string                         `json:"schema_version"`
	CurrentState  string                         `json:"current_state"`
	Items         []compliance.PolicyDriftRecord `json:"items"`
	Limitations   []string                       `json:"limitations,omitempty"`
}

type phase4PolicyDriftResponse struct {
	Status      string                       `json:"status"`
	PolicyDrift compliance.PolicyDriftRecord `json:"policy_drift"`
}

type phase4ExecutiveReportListResponse struct {
	SchemaVersion string                       `json:"schema_version"`
	CurrentState  string                       `json:"current_state"`
	Items         []compliance.ExecutiveReport `json:"items"`
	Limitations   []string                     `json:"limitations,omitempty"`
}

type phase4ExecutiveReportResponse struct {
	Status string                     `json:"status"`
	Report compliance.ExecutiveReport `json:"report"`
}

type phase4ProofsResponse struct {
	SchemaVersion             string                               `json:"schema_version"`
	CurrentState              string                               `json:"current_state"`
	WorkflowArtifacts         []workflow.LifecycleRecord           `json:"workflow_artifacts,omitempty"`
	ReconciliationArtifacts   []connectors.ReconciliationRecord    `json:"reconciliation_artifacts,omitempty"`
	PartnerArtifacts          []handoff.IntakeRecord               `json:"partner_artifacts,omitempty"`
	PartnerDashboardArtifacts []handoff.DashboardProjection        `json:"partner_dashboard_artifacts,omitempty"`
	ComplianceArtifacts       []compliance.ComplianceMappingRecord `json:"compliance_artifacts,omitempty"`
	PolicyDriftArtifacts      []compliance.PolicyDriftRecord       `json:"policy_drift_artifacts,omitempty"`
	ExecutiveArtifacts        []compliance.ExecutiveReport         `json:"executive_artifacts,omitempty"`
	Limitations               []string                             `json:"limitations,omitempty"`
}

func (s server) enterpriseWorkflowLifecycleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		items, err := s.listPhase4WorkflowArtifacts(ctx, filter)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase4WorkflowListResponse{
			SchemaVersion: phase4WorkflowListSchema,
			CurrentState:  map[bool]string{true: phase4WorkflowStateActive, false: "workflow_orchestration_empty"}[len(items) > 0],
			Items:         items,
			Limitations: []string{
				"Workflow lifecycle remains canonical only when linked validation and approval discipline are satisfied.",
			},
		})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		var request phase4WorkflowRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request.Input.SubjectRef = firstNonEmpty(normalizePhase4SubjectRef(request.Input.SubjectRef), filter.SubjectRef)
		record := workflow.EvaluateLifecycle(request.Input, time.Now)
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		if err := s.persistPhase4Event(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeEnterpriseWorkflowRecorded, filter, &record, nil, nil, nil, nil, nil); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusCreated, phase4WorkflowResponse{
			Status:   "recorded",
			Workflow: record,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) enterpriseConnectorReconciliationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		items, err := s.listPhase4ConnectorArtifacts(ctx, filter)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase4ConnectorListResponse{
			SchemaVersion: phase4ConnectorListSchema,
			CurrentState:  map[bool]string{true: phase4ConnectorStateActive, false: "connector_reconciliation_empty"}[len(items) > 0],
			Items:         items,
			Limitations: []string{
				"Connector reconciliation remains a projection layer and never replaces the canonical workflow state or validation truth.",
			},
		})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		var request phase4ConnectorRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request.Input.SubjectRef = firstNonEmpty(normalizePhase4SubjectRef(request.Input.SubjectRef), filter.SubjectRef)
		record := connectors.Reconcile(request.Input, time.Now)
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		if err := s.persistPhase4Event(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeEnterpriseConnectorReconciliationRecorded, filter, nil, &record, nil, nil, nil, nil); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusCreated, phase4ConnectorResponse{
			Status:         "recorded",
			Reconciliation: record,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) enterprisePartnerIntakeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		items, err := s.listPhase4PartnerArtifacts(ctx, filter)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase4PartnerIntakeListResponse{
			SchemaVersion: phase4PartnerIntakeListSchema,
			CurrentState:  map[bool]string{true: phase4PartnerStateActive, false: "partner_trust_empty"}[len(items) > 0],
			Items:         items,
			Limitations: []string{
				"Partner intake remains tied to local verifier evidence and local policy compatibility review.",
			},
		})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		var request phase4PartnerIntakeRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		record := handoff.EvaluateIntake(request.Input, time.Now)
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		if err := s.persistPhase4Event(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeEnterprisePartnerTrustRecorded, filter, nil, nil, &record, nil, nil, nil); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusCreated, phase4PartnerIntakeResponse{
			Status: "recorded",
			Intake: record,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) enterprisePartnerDashboardHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parsePhase4EnterpriseFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	items, err := s.listPhase4PartnerArtifacts(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	dashboards := make([]handoff.DashboardProjection, 0, len(items))
	for _, item := range items {
		dashboards = append(dashboards, item.Dashboard)
	}
	httpjson.Write(w, http.StatusOK, phase4PartnerDashboardResponse{
		SchemaVersion: phase4PartnerDashboardListSchema,
		CurrentState:  map[bool]string{true: phase4PartnerStateActive, false: "partner_dashboard_empty"}[len(dashboards) > 0],
		Items:         dashboards,
		Limitations: []string{
			"Partner dashboards remain bounded projections and intentionally exclude internal-only runtime and investigation context.",
		},
	})
}

func (s server) enterpriseComplianceMappingHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		items, err := s.listPhase4ComplianceArtifacts(ctx, filter)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase4ComplianceListResponse{
			SchemaVersion: phase4ComplianceListSchema,
			CurrentState:  map[bool]string{true: phase4ComplianceStateActive, false: "compliance_mapping_empty"}[len(items) > 0],
			Items:         items,
			Limitations: []string{
				"Compliance mappings remain bounded by evidence coverage and explicitly distinguish full, partial, inferred, and missing support.",
			},
		})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		var request phase4ComplianceRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request.Input.SubjectRef = firstNonEmpty(normalizePhase4SubjectRef(request.Input.SubjectRef), filter.SubjectRef)
		record := compliance.EvaluateComplianceMapping(request.Input, time.Now)
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		if err := s.persistPhase4Event(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeEnterpriseComplianceMappingRecorded, filter, nil, nil, nil, &record, nil, nil); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusCreated, phase4ComplianceResponse{
			Status:     "recorded",
			Compliance: record,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) enterprisePolicyDriftHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		items, err := s.listPhase4DriftArtifacts(ctx, filter)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase4PolicyDriftListResponse{
			SchemaVersion: phase4PolicyDriftListSchema,
			CurrentState:  map[bool]string{true: phase4PolicyDriftStateActive, false: "policy_drift_empty"}[len(items) > 0],
			Items:         items,
			Limitations: []string{
				"Policy drift reporting keeps identity trail and evidence linkage; it does not replace bundle history or approval boundaries.",
			},
		})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		var request phase4PolicyDriftRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		request.Input.SubjectRef = firstNonEmpty(normalizePhase4SubjectRef(request.Input.SubjectRef), filter.SubjectRef)
		record := compliance.EvaluatePolicyDrift(request.Input, time.Now)
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		if err := s.persistPhase4Event(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeEnterprisePolicyDriftRecorded, filter, nil, nil, nil, nil, &record, nil); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusCreated, phase4PolicyDriftResponse{
			Status:      "recorded",
			PolicyDrift: record,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) enterpriseExecutiveReportHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		items, err := s.listPhase4ExecutiveArtifacts(ctx, filter)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase4ExecutiveReportListResponse{
			SchemaVersion: phase4ExecutiveReportListSchema,
			CurrentState:  map[bool]string{true: phase4ExecutiveStateActive, false: "executive_reporting_empty"}[len(items) > 0],
			Items:         items,
			Limitations: []string{
				"Executive reporting is a bounded governance summary and remains traceable back to workflow, partner, compliance, and drift evidence.",
			},
		})
	case http.MethodPost:
		principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleOperator, auth.RoleSecurityAdmin)
		if !ok {
			return
		}
		r = authorizedRequest
		var err error
		r, err = applyPrincipalTenantToRequest(principal, r)
		if err != nil {
			httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
			return
		}
		filter, err := parsePhase4EnterpriseFilter(r)
		if err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		var request phase4ExecutiveReportRequest
		if err := httpjson.Decode(r, &request); err != nil {
			httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
		defer cancel()
		report, err := s.buildPhase4ExecutiveReport(ctx, filter, request.ScopeRef)
		if err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		if err := s.persistPhase4Event(ctx, requestIDFromHeader(r), principal.Subject, audit.EventTypeEnterpriseExecutiveReportRecorded, filter, nil, nil, nil, nil, nil, &report); err != nil {
			httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		httpjson.Write(w, http.StatusOK, phase4ExecutiveReportResponse{
			Status: "recorded",
			Report: report,
		})
	default:
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (s server) enterprisePhase4ProofsHandler(w http.ResponseWriter, r *http.Request) {
	principal, authorizedRequest, ok := s.authorize(w, r, auth.RoleViewer, auth.RoleOperator, auth.RoleSecurityAdmin)
	if !ok {
		return
	}
	r = authorizedRequest
	var err error
	r, err = applyPrincipalTenantToRequest(principal, r)
	if err != nil {
		httpjson.Write(w, auth.StatusCode(err), map[string]string{"error": err.Error()})
		return
	}
	if r.Method != http.MethodGet {
		httpjson.Write(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	filter, err := parsePhase4EnterpriseFilter(r)
	if err != nil {
		httpjson.Write(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), s.requestTimeout)
	defer cancel()
	workflows, err := s.listPhase4WorkflowArtifacts(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	reconciliations, err := s.listPhase4ConnectorArtifacts(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	partners, err := s.listPhase4PartnerArtifacts(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	complianceItems, err := s.listPhase4ComplianceArtifacts(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	drifts, err := s.listPhase4DriftArtifacts(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	executive, err := s.listPhase4ExecutiveArtifacts(ctx, filter)
	if err != nil {
		httpjson.Write(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	dashboards := make([]handoff.DashboardProjection, 0, len(partners))
	for _, item := range partners {
		dashboards = append(dashboards, item.Dashboard)
	}
	currentState := phase4ProofStateIncomplete
	if hasPhase4ValidatedWorkflow(workflows) && hasPhase4ReconciledConnector(reconciliations) && hasPhase4AcceptedPartner(partners) && len(complianceItems) > 0 && len(drifts) > 0 && len(executive) > 0 {
		currentState = phase4ProofStateActive
	}
	httpjson.Write(w, http.StatusOK, phase4ProofsResponse{
		SchemaVersion:             phase4ProofsSchema,
		CurrentState:              currentState,
		WorkflowArtifacts:         takePhase4WorkflowArtifacts(workflows, 5),
		ReconciliationArtifacts:   takePhase4ConnectorArtifacts(reconciliations, 5),
		PartnerArtifacts:          takePhase4PartnerArtifacts(partners, 5),
		PartnerDashboardArtifacts: takePhase4DashboardArtifacts(dashboards, 5),
		ComplianceArtifacts:       takePhase4ComplianceArtifacts(complianceItems, 5),
		PolicyDriftArtifacts:      takePhase4DriftArtifacts(drifts, 5),
		ExecutiveArtifacts:        takePhase4ExecutiveArtifacts(executive, 5),
		Limitations: []string{
			"Phase 4 proofs expose bounded enterprise workflow, partner trust, governance, and executive reporting artifacts tied to the canonical audit spine.",
		},
	})
}

func parsePhase4EnterpriseFilter(r *http.Request) (phase4EnterpriseFilter, error) {
	filter := phase4EnterpriseFilter{
		TenantID:    strings.TrimSpace(r.URL.Query().Get("tenant_id")),
		Environment: strings.TrimSpace(r.URL.Query().Get("environment")),
		Repo:        strings.TrimSpace(r.URL.Query().Get("repo")),
		SubjectRef:  normalizePhase4SubjectRef(strings.TrimSpace(r.URL.Query().Get("subject_ref"))),
		WorkflowID:  strings.TrimSpace(r.URL.Query().Get("workflow_id")),
		PartnerID:   strings.TrimSpace(r.URL.Query().Get("partner_id")),
		Limit:       runtimeLimit(r),
	}
	if filter.SubjectRef != "" {
		clusterID, _, _, _, err := parseRuntimeSubjectRef(filter.SubjectRef)
		if err != nil {
			return filter, audit.ErrInvalidFilter
		}
		filter.ClusterID = clusterID
	}
	return filter, nil
}

func (s server) buildPhase4ExecutiveReport(ctx context.Context, filter phase4EnterpriseFilter, scopeRef string) (compliance.ExecutiveReport, error) {
	workflows, err := s.listPhase4WorkflowArtifacts(ctx, filter)
	if err != nil {
		return compliance.ExecutiveReport{}, err
	}
	reconciliations, err := s.listPhase4ConnectorArtifacts(ctx, filter)
	if err != nil {
		return compliance.ExecutiveReport{}, err
	}
	partners, err := s.listPhase4PartnerArtifacts(ctx, filter)
	if err != nil {
		return compliance.ExecutiveReport{}, err
	}
	complianceItems, err := s.listPhase4ComplianceArtifacts(ctx, filter)
	if err != nil {
		return compliance.ExecutiveReport{}, err
	}
	drifts, err := s.listPhase4DriftArtifacts(ctx, filter)
	if err != nil {
		return compliance.ExecutiveReport{}, err
	}
	return compliance.BuildExecutiveReport(compliance.ExecutiveReportInput{
		ScopeRef:                firstNonEmpty(strings.TrimSpace(scopeRef), filter.SubjectRef, "tenant:"+filter.TenantID),
		WorkflowArtifacts:       workflows,
		ReconciliationArtifacts: reconciliations,
		PartnerArtifacts:        partners,
		ComplianceArtifacts:     complianceItems,
		DriftArtifacts:          drifts,
	}, time.Now), nil
}

func (s server) persistPhase4Event(ctx context.Context, requestID, actor, eventType string, filter phase4EnterpriseFilter, workflowRecord *workflow.LifecycleRecord, reconciliation *connectors.ReconciliationRecord, partner *handoff.IntakeRecord, complianceRecord *compliance.ComplianceMappingRecord, drift *compliance.PolicyDriftRecord, executive *compliance.ExecutiveReport) error {
	payload, err := canonicalJSON(phase4EnterprisePayload{
		SchemaVersion:  phase4EnterprisePayloadSchema,
		Workflow:       workflowRecord,
		Reconciliation: reconciliation,
		PartnerIntake:  partner,
		Compliance:     complianceRecord,
		PolicyDrift:    drift,
		Executive:      executive,
	})
	if err != nil {
		return err
	}
	subjectRef := ""
	digest := ""
	reasons := []string{"phase4_enterprise_recorded"}
	decision := audit.DecisionAllow
	if workflowRecord != nil {
		subjectRef = firstNonEmpty(subjectRef, workflowRecord.SubjectRef)
		reasons = append(reasons, workflowRecord.CurrentState)
		reasons = append(reasons, workflowRecord.ReasonCodes...)
		if workflowRecord.CurrentState == workflow.StateRejected {
			decision = audit.DecisionDeny
		}
	}
	if reconciliation != nil {
		subjectRef = firstNonEmpty(subjectRef, reconciliation.SubjectRef)
		reasons = append(reasons, reconciliation.CurrentState)
		reasons = append(reasons, reconciliation.ReasonCodes...)
	}
	if partner != nil {
		reasons = append(reasons, partner.CurrentState, partner.PartnerID)
		reasons = append(reasons, partner.ReasonCodes...)
		if partner.CurrentState == handoff.IntakeStateRejected {
			decision = audit.DecisionDeny
		}
	}
	if complianceRecord != nil {
		subjectRef = firstNonEmpty(subjectRef, complianceRecord.SubjectRef)
		reasons = append(reasons, complianceRecord.CurrentState, complianceRecord.CoverageState)
		reasons = append(reasons, complianceRecord.ReasonCodes...)
	}
	if drift != nil {
		subjectRef = firstNonEmpty(subjectRef, drift.SubjectRef)
		reasons = append(reasons, drift.CurrentState)
		reasons = append(reasons, drift.ImpactSummary...)
	}
	if executive != nil {
		reasons = append(reasons, executive.CurrentState)
		reasons = append(reasons, executive.Highlights...)
	}
	clusterID := filter.ClusterID
	namespace := ""
	workloadKind := ""
	workload := ""
	if parsedCluster, parsedNamespace, parsedWorkloadKind, parsedWorkload, err := parseRuntimeSubjectRef(subjectRef); err == nil {
		clusterID = firstNonEmpty(clusterID, parsedCluster)
		namespace = parsedNamespace
		workloadKind = parsedWorkloadKind
		workload = parsedWorkload
	}
	_, err = s.store.Ingest(ctx, audit.Event{
		RequestID:    firstNonEmpty(strings.TrimSpace(requestID), audit.NewRequestID()),
		Component:    phase4EnterpriseComponent,
		EventType:    eventType,
		Actor:        strings.TrimSpace(actor),
		ClusterID:    clusterID,
		TenantID:     firstNonEmpty(filter.TenantID, audit.TenantFromNamespace(namespace)),
		Environment:  firstNonEmpty(filter.Environment, audit.EnvironmentFromNamespace(namespace)),
		Namespace:    namespace,
		WorkloadKind: workloadKind,
		Workload:     workload,
		Repo:         filter.Repo,
		Digest:       digest,
		Decision:     decision,
		Reasons:      uniquePhase4Strings(reasons),
		Enterprise:   payload,
	})
	return err
}

func (s server) listPhase4WorkflowArtifacts(ctx context.Context, filter phase4EnterpriseFilter) ([]workflow.LifecycleRecord, error) {
	events, err := s.listPhase4Events(ctx, filter, audit.EventTypeEnterpriseWorkflowRecorded)
	if err != nil {
		return nil, err
	}
	items := []workflow.LifecycleRecord{}
	for _, item := range events {
		payload := parsePhase4EnterprisePayload(item.Enterprise)
		if payload.Workflow == nil {
			continue
		}
		if !matchesPhase4Subject(filter.SubjectRef, payload.Workflow.SubjectRef) {
			continue
		}
		if filter.WorkflowID != "" && payload.Workflow.WorkflowID != filter.WorkflowID {
			continue
		}
		items = append(items, *payload.Workflow)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ObservedAt.After(items[j].ObservedAt) })
	return items, nil
}

func (s server) listPhase4ConnectorArtifacts(ctx context.Context, filter phase4EnterpriseFilter) ([]connectors.ReconciliationRecord, error) {
	events, err := s.listPhase4Events(ctx, filter, audit.EventTypeEnterpriseConnectorReconciliationRecorded)
	if err != nil {
		return nil, err
	}
	items := []connectors.ReconciliationRecord{}
	for _, item := range events {
		payload := parsePhase4EnterprisePayload(item.Enterprise)
		if payload.Reconciliation == nil {
			continue
		}
		if !matchesPhase4Subject(filter.SubjectRef, payload.Reconciliation.SubjectRef) {
			continue
		}
		if filter.WorkflowID != "" && payload.Reconciliation.WorkflowID != filter.WorkflowID {
			continue
		}
		items = append(items, *payload.Reconciliation)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ObservedAt.After(items[j].ObservedAt) })
	return items, nil
}

func (s server) listPhase4PartnerArtifacts(ctx context.Context, filter phase4EnterpriseFilter) ([]handoff.IntakeRecord, error) {
	events, err := s.listPhase4Events(ctx, filter, audit.EventTypeEnterprisePartnerTrustRecorded)
	if err != nil {
		return nil, err
	}
	items := []handoff.IntakeRecord{}
	for _, item := range events {
		payload := parsePhase4EnterprisePayload(item.Enterprise)
		if payload.PartnerIntake == nil {
			continue
		}
		if filter.PartnerID != "" && payload.PartnerIntake.PartnerID != filter.PartnerID {
			continue
		}
		items = append(items, *payload.PartnerIntake)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ObservedAt.After(items[j].ObservedAt) })
	return items, nil
}

func (s server) listPhase4ComplianceArtifacts(ctx context.Context, filter phase4EnterpriseFilter) ([]compliance.ComplianceMappingRecord, error) {
	events, err := s.listPhase4Events(ctx, filter, audit.EventTypeEnterpriseComplianceMappingRecorded)
	if err != nil {
		return nil, err
	}
	items := []compliance.ComplianceMappingRecord{}
	for _, item := range events {
		payload := parsePhase4EnterprisePayload(item.Enterprise)
		if payload.Compliance == nil {
			continue
		}
		if !matchesPhase4Subject(filter.SubjectRef, payload.Compliance.SubjectRef) {
			continue
		}
		items = append(items, *payload.Compliance)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ObservedAt.After(items[j].ObservedAt) })
	return items, nil
}

func (s server) listPhase4DriftArtifacts(ctx context.Context, filter phase4EnterpriseFilter) ([]compliance.PolicyDriftRecord, error) {
	events, err := s.listPhase4Events(ctx, filter, audit.EventTypeEnterprisePolicyDriftRecorded)
	if err != nil {
		return nil, err
	}
	items := []compliance.PolicyDriftRecord{}
	for _, item := range events {
		payload := parsePhase4EnterprisePayload(item.Enterprise)
		if payload.PolicyDrift == nil {
			continue
		}
		if !matchesPhase4Subject(filter.SubjectRef, payload.PolicyDrift.SubjectRef) {
			continue
		}
		items = append(items, *payload.PolicyDrift)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ObservedAt.After(items[j].ObservedAt) })
	return items, nil
}

func (s server) listPhase4ExecutiveArtifacts(ctx context.Context, filter phase4EnterpriseFilter) ([]compliance.ExecutiveReport, error) {
	events, err := s.listPhase4Events(ctx, filter, audit.EventTypeEnterpriseExecutiveReportRecorded)
	if err != nil {
		return nil, err
	}
	items := []compliance.ExecutiveReport{}
	for _, item := range events {
		payload := parsePhase4EnterprisePayload(item.Enterprise)
		if payload.Executive == nil {
			continue
		}
		items = append(items, *payload.Executive)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].GeneratedAt.After(items[j].GeneratedAt) })
	return items, nil
}

func (s server) listPhase4Events(ctx context.Context, filter phase4EnterpriseFilter, eventType string) ([]audit.StoredEvent, error) {
	return s.store.ListEvents(ctx, audit.EventFilter{
		EventType:   eventType,
		Component:   phase4EnterpriseComponent,
		TenantID:    filter.TenantID,
		Environment: filter.Environment,
		Repo:        filter.Repo,
		Limit:       max(filter.Limit, 200),
	})
}

func parsePhase4EnterprisePayload(value json.RawMessage) phase4EnterprisePayload {
	if len(value) == 0 || string(value) == "null" {
		return phase4EnterprisePayload{}
	}
	var payload phase4EnterprisePayload
	if err := json.Unmarshal(value, &payload); err != nil {
		return phase4EnterprisePayload{}
	}
	return payload
}

func hasPhase4ValidatedWorkflow(items []workflow.LifecycleRecord) bool {
	for _, item := range items {
		if item.ClosureReady || item.ValidationState == workflow.ValidationStateVerified || item.CurrentState == workflow.StateValidated {
			return true
		}
	}
	return false
}

func hasPhase4ReconciledConnector(items []connectors.ReconciliationRecord) bool {
	for _, item := range items {
		if item.SafeToAutoClose || item.CurrentState == connectors.StateSynced {
			return true
		}
	}
	return false
}

func hasPhase4AcceptedPartner(items []handoff.IntakeRecord) bool {
	for _, item := range items {
		if item.CurrentState == handoff.IntakeStateAccepted {
			return true
		}
	}
	return false
}

func takePhase4WorkflowArtifacts(items []workflow.LifecycleRecord, limit int) []workflow.LifecycleRecord {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takePhase4ConnectorArtifacts(items []connectors.ReconciliationRecord, limit int) []connectors.ReconciliationRecord {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takePhase4PartnerArtifacts(items []handoff.IntakeRecord, limit int) []handoff.IntakeRecord {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takePhase4DashboardArtifacts(items []handoff.DashboardProjection, limit int) []handoff.DashboardProjection {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takePhase4ComplianceArtifacts(items []compliance.ComplianceMappingRecord, limit int) []compliance.ComplianceMappingRecord {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takePhase4DriftArtifacts(items []compliance.PolicyDriftRecord, limit int) []compliance.PolicyDriftRecord {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func takePhase4ExecutiveArtifacts(items []compliance.ExecutiveReport, limit int) []compliance.ExecutiveReport {
	if len(items) > limit {
		return items[:limit]
	}
	return items
}

func matchesPhase4Subject(filterSubject, itemSubject string) bool {
	filterSubject = normalizePhase4SubjectRef(filterSubject)
	itemSubject = normalizePhase4SubjectRef(itemSubject)
	return filterSubject == "" || filterSubject == itemSubject
}

func normalizePhase4SubjectRef(value string) string {
	value = strings.TrimSpace(value)
	if value == "" || strings.Contains(value, "|") {
		return value
	}
	parts := strings.Split(value, "/")
	if len(parts) != 4 {
		return value
	}
	return runtimeSubjectRef(parts[0], parts[1], parts[2], parts[3])
}

func uniquePhase4Strings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	seen := map[string]struct{}{}
	items := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		items = append(items, value)
	}
	if len(items) == 0 {
		return nil
	}
	return items
}
